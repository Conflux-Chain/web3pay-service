package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	utilcfg "github.com/Conflux-Chain/go-conflux-util/config"
	utilviper "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestSDKConfigInitialization(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "config.yaml"), []byte("eth:\n  requestTimeout: 9s\n  retry: 2\n  retryInterval: 2s\n"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, ".env"), []byte("WEB3PAY_ETH_REQUESTTIMEOUT=7s\nWEB3PAY_ETH_RETRY=3\nWEB3PAY_ETH_MAXCONNSPERHOST=0\nWEB3PAY_LOG_LEVEL=error\n"), 0600))
	executable, err := os.Executable()
	require.NoError(t, err)
	cmd := exec.Command(executable, "-test.run=^TestSDKConfigProcess$")
	cmd.Dir = dir
	for _, entry := range os.Environ() {
		if !strings.HasPrefix(entry, "WEB3PAY_") {
			cmd.Env = append(cmd.Env, entry)
		}
	}
	cmd.Env = append(cmd.Env, "WEB3PAY_CONFIG_HELPER=1", "WEB3PAY_ETH_REQUESTTIMEOUT=5s")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", output)
}

func TestSDKConfigProcess(t *testing.T) {
	if os.Getenv("WEB3PAY_CONFIG_HELPER") != "1" {
		t.Skip("runs in an isolated process because SDK initialization is startup-only")
	}
	utilcfg.MustInit("web3pay")
	require.Equal(t, logrus.ErrorLevel, logrus.GetLevel(), ".env is loaded before logging initializes")
	var conf ethClientConfig
	require.NoError(t, utilviper.UnmarshalKey("eth", &conf))
	require.Equal(t, 5*time.Second, conf.RequestTimeout, "process environment overrides .env and YAML")
	require.Equal(t, 3, conf.Retry, ".env overrides YAML")
	require.Equal(t, 2*time.Second, conf.RetryInterval, "YAML applies when the environment has no value")
	require.Equal(t, 1024, conf.MaxConnsPerHost, "SDK defaults replace an explicitly configured zero")
	var missing ethClientConfig
	require.NoError(t, utilviper.UnmarshalKey("missing", &missing))
	require.Equal(t, 3*time.Second, missing.RequestTimeout)
	require.Equal(t, time.Second, missing.RetryInterval)
	require.Equal(t, 1024, missing.MaxConnsPerHost)
}
