package service

import (
	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/store/memdb"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

type Factory struct {
	Blockchain *BlockchainService
	Billing    *BillingService
}

func MustNewFactory(
	w3c *web3go.Client, store *sqlite.SqliteStore, memStore *memdb.MemStore, provider *blockchain.Provider) *Factory {
	blockchainSvc, err := NewBlockchainService(store, memStore, provider)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create blockchain service")
	}

	return &Factory{
		Blockchain: blockchainSvc,
		Billing:    NewBillingService(store, blockchainSvc),
	}
}
