package worker

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	utilviper "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestSDKSettlementConfigDefaults(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	configFile := filepath.Join(t.TempDir(), "config.yaml")
	require.NoError(t, os.WriteFile(configFile, []byte("worker:\n  pollingInterval: 2s\n  pollingBatchSize: 4\nstore:\n  sqlite:\n    database: 'file:config-test?mode=memory&cache=shared'\n"), 0600))
	t.Setenv("WEB3PAY_WORKER_MAXSETTLEMENTRETRIES", "0")
	t.Setenv("WEB3PAY_WORKER_MINCONFIRMEDBLOCKS", "0")
	t.Setenv("WEB3PAY_STORE_SQLITE_MAXIDLECONNS", "0")
	utilviper.MustInit("web3pay", configFile)
	worker := MustNewBlockchainWorkerFromViper(nil, nil, nil, nil)
	require.Equal(t, 2*time.Second, worker.PollingInterval)
	require.Equal(t, 4, worker.PollingBatchSize)
	require.Equal(t, uint(5), worker.MaxSettlementRetries)
	require.Equal(t, uint64(30), worker.MinConfirmedBlocks)
	require.Equal(t, 30*time.Second, worker.MaxPendingAwaitDuration)
	require.False(t, worker.PersistOnShutdown)
	storeConfig := sqlite.MustNewConfigFromViper()
	require.Equal(t, "file:config-test?mode=memory&cache=shared", storeConfig.Database)
	require.Equal(t, 10, storeConfig.MaxIdleConns)
	require.Equal(t, 10, storeConfig.MaxOpenConns)
	require.Zero(t, storeConfig.ConnMaxLifetime)
}
