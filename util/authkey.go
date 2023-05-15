package util

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"

	"github.com/Conflux-Chain/web3pay-service/types"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
)

const sigAddressCacheSize = 50_000

var (
	stdAuthKeyManager = NewAuthKeyManager()
)

type AuthKeyManager struct {
	lock         sync.Mutex
	msgCache     map[string]string // contract => api auth message
	sigAddrCache *lru.Cache        // sha3(sig) => addr
}

func NewAuthKeyManager() *AuthKeyManager {
	lruCache, _ := lru.New(sigAddressCacheSize)
	return &AuthKeyManager{
		sigAddrCache: lruCache,
		msgCache:     make(map[string]string),
	}
}

type apiAuthMessage struct {
	Domain   string `json:"domain"`   // always be "web3pay"
	Contract string `json:"contract"` // APP contract address
}

func (m *AuthKeyManager) getApiAuthMessage(contract string) (string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if msg, ok := m.msgCache[contract]; ok {
		return msg, nil
	}

	msgb, err := json.Marshal(&apiAuthMessage{
		Domain: "web3pay", Contract: contract,
	})
	if err != nil {
		return "", errors.WithMessage(err, "failed to assemble message")
	}

	msg := string(msgb)
	m.msgCache[contract] = msg

	return msg, nil
}

func (m *AuthKeyManager) GetAddrByApiAuthKey(contract string, apiKey *types.ApiAuthKey) (common.Address, error) {
	cacheKey := crypto.Keccak256Hash([]byte(contract + apiKey.Sig))

	return m.getAddrByCacheKey(cacheKey.String(), func() (common.Address, error) {
		msg, err := m.getApiAuthMessage(contract)
		if err != nil {
			return common.Address{}, err
		}

		addrStr, err := RecoverAddress(msg, apiKey.Sig)
		if err != nil {
			return common.Address{}, err
		}

		return common.HexToAddress(addrStr), nil
	})
}

func (m *AuthKeyManager) GetAddrByBillingAuthKey(billingKey *types.BillingAuthKey) (common.Address, error) {
	cacheKey := crypto.Keccak256Hash([]byte(billingKey.Msg + billingKey.Sig))

	return m.getAddrByCacheKey(cacheKey.String(), func() (common.Address, error) {
		addrStr, err := RecoverAddress(billingKey.Msg, billingKey.Sig)
		if err != nil {
			return common.Address{}, err
		}

		return common.HexToAddress(addrStr), nil
	})
}

func (m *AuthKeyManager) getAddrByCacheKey(
	cacheKey string, genAddr func() (common.Address, error)) (common.Address, error) {

	if val, ok := m.sigAddrCache.Get(cacheKey); ok { // hit in cache
		return val.(common.Address), nil
	}

	lockKey := MutexKey(cacheKey)
	KLock(lockKey)
	defer KUnlock(lockKey)

	if val, ok := m.sigAddrCache.Get(cacheKey); ok { // double check
		return val.(common.Address), nil
	}

	addr, err := genAddr()
	if err != nil {
		return common.Address{}, err
	}

	m.sigAddrCache.Add(cacheKey, addr)
	return addr, nil
}

func getApiAuthMessage(contract string) (string, error) {
	return stdAuthKeyManager.getApiAuthMessage(contract)
}

func GetAddrByApiAuthKey(contract string, apiKey *types.ApiAuthKey) (common.Address, error) {
	return stdAuthKeyManager.GetAddrByApiAuthKey(contract, apiKey)
}

func GetAddrByApiKey(contract, apiKey string) (common.Address, error) {
	key, err := ParseApiKey(apiKey)
	if err != nil {
		return common.Address{}, err
	}

	return stdAuthKeyManager.GetAddrByApiAuthKey(contract, key)
}

func GetAddrByBillingAuthKey(billingKey *types.BillingAuthKey) (common.Address, error) {
	return stdAuthKeyManager.GetAddrByBillingAuthKey(billingKey)
}

func GetAddrByBillingKey(billingKey string) (common.Address, error) {
	key, err := ParseBillingKey(billingKey)
	if err != nil {
		return common.Address{}, err
	}

	return stdAuthKeyManager.GetAddrByBillingAuthKey(key)
}

// BuildApiKey utility function to help build API key with specified APP contract address
// and consumer private key text.
func BuildApiKey(appContract string, consumerPrivateKeyText string) (string, error) {
	apiAuthMessage, err := getApiAuthMessage(appContract)
	if err != nil {
		return "", errors.WithMessage(err, "API auth message error")
	}

	// load private key
	privateKey, err := EcdsaPrivateKeyFromString(consumerPrivateKeyText)
	if err != nil {
		return "", err
	}

	// create signature
	sigstr, _, err := PersonalSign(apiAuthMessage, privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create signature")
	}

	// base58 encoding signature
	sig, _ := hexutil.Decode(sigstr)
	apiKey := base58.Encode(sig)

	return apiKey, nil
}

// alphabet is the modified base58 alphabet used by Bitcoin.
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func isBase58EncodedString(str string) bool {
	for _, rb := range str {
		if !strings.ContainsRune(base58Alphabet, rb) {
			return false
		}
	}

	return true
}

func ParseApiKey(apiKey string) (*types.ApiAuthKey, error) {
	if !isBase58EncodedString(apiKey) {
		return nil, errors.New("malformed encoded key")
	}

	sig := base58.Decode(apiKey)
	if len(sig) < 65 {
		return nil, errors.New("signature bytes too short")
	}

	return &types.ApiAuthKey{Sig: hexutil.Encode(sig)}, nil
}

// BuildBillingKey utility function to help build billing key with specified APP contract address
// and its owner private key text.
func BuildBillingKey(appContract string, ownerPrivateKeyText string) (string, error) {
	// load private key
	privateKey, err := EcdsaPrivateKeyFromString(ownerPrivateKeyText)
	if err != nil {
		return "", err
	}

	// create signature
	sig, _, err := PersonalSign(appContract, privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create signature")
	}

	// json marshal auth key
	authKeyObj, err := json.Marshal(types.BillingAuthKey{
		Msg: appContract, Sig: sig,
	})
	if err != nil {
		return "", errors.WithMessage(err, "failed to json marshal auth key object")
	}

	// base64 encoding auth key json
	billKey := base64.StdEncoding.EncodeToString(authKeyObj)
	return billKey, nil
}

func ParseBillingKey(billingKey string) (*types.BillingAuthKey, error) {
	keyJson, err := base64.StdEncoding.DecodeString(billingKey)
	if err != nil {
		return nil, errors.WithMessage(err, "base64 decode error")
	}

	var key types.BillingAuthKey
	if err := json.Unmarshal(keyJson, &key); err != nil {
		return nil, errors.WithMessage(err, "json decode error")
	}

	if len(key.Msg) == 0 {
		return nil, errors.New("msg not provided")
	}

	if len(key.Sig) == 0 {
		return nil, errors.New("sig not provided")
	}

	// `msg` part must be a valid hex address
	if !common.IsHexAddress(key.Msg) {
		return nil, errors.New("msg not valid hex address")
	}

	return &key, err
}
