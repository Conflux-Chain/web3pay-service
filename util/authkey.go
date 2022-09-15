package util

import (
	"encoding/base64"
	"encoding/json"
	"sync"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

var (
	stdAuthKeyManager = NewAuthKeyManager()
)

type AuthKeyManager struct {
	lock     sync.Mutex
	msgCache map[string]string // contract => api auth message
}

func NewAuthKeyManager() *AuthKeyManager {
	return &AuthKeyManager{
		msgCache: make(map[string]string),
	}
}

type apiAuthMessage struct {
	Domain   string `json:"domain"`   // always be "web3pay"
	Contract string `json:"contract"` // APP coin contract address
}

func (m *AuthKeyManager) GetApiAuthMessage(contract string) (string, error) {
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

func GetApiAuthMessage(contract string) (string, error) {
	return stdAuthKeyManager.GetApiAuthMessage(contract)
}

// BuildApiKey utility function to help build API key with specified APP coin contract address
// and consumer private key text.
func BuildApiKey(appCoinContract string, consumerPrivateKeyText string) (string, error) {
	apiAuthMessage, err := GetApiAuthMessage(appCoinContract)
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
	apiKey := base58.CheckEncode(sig, 0)

	return apiKey, nil
}

// BuildBillingKey utility function to help build billing key with specified APP coin contract address
// and its owner private key text.
func BuildBillingKey(appCoinContract string, ownerPrivateKeyText string) (string, error) {
	// load private key
	privateKey, err := EcdsaPrivateKeyFromString(ownerPrivateKeyText)
	if err != nil {
		return "", err
	}

	// create signature
	sig, _, err := PersonalSign(appCoinContract, privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create signature")
	}

	// json marshal auth key
	authKeyObj, err := json.Marshal(model.BillingAuthKey{
		Msg: appCoinContract, Sig: sig,
	})
	if err != nil {
		return "", errors.WithMessage(err, "failed to json marshal auth key object")
	}

	// base64 encoding auth key json
	billKey := base64.StdEncoding.EncodeToString(authKeyObj)
	return billKey, nil
}
