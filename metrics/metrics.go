package metrics

import (
	metricUtil "github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	RPC     RpcMetrics
	Monitor MonitorMetrics
	Store   StoreMetrics
	Worker  WorkerMetrics
)

// RPC metrics
type RpcMetrics struct{}

func (*RpcMetrics) UpdateWithCollector(c *RpcCollector) {
	// Overall rate statistics
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(100, "/rpc/rate/success").Mark(c.success())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/bizErr").Mark(c.isBizError())

	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/2xx").Mark(c.is2xx())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/4xx").Mark(c.is4xx())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/5xx").Mark(c.is5xx())

	metricUtil.GetOrRegisterTimer("/rpc/duration/all").UpdateSince(c.start)
	if c.success() {
		metricUtil.GetOrRegisterTimer("/rpc/duration/success").UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/2xx").UpdateSince(c.start)
	case c.is4xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/4xx").UpdateSince(c.start)
	case c.is5xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/5xx").UpdateSince(c.start)
	}

	module := c.module()
	if len(module) == 0 {
		return
	}

	// RPC rate statistics
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(100, "/rpc/rate/success/%v", module).Mark(c.success())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/bizErr/%v", module).Mark(c.isBizError())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/2xx/%v", module).Mark(c.is2xx())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/4xx/%v", module).Mark(c.is4xx())
	metricUtil.GetOrRegisterTimeWindowPercentageDefault(0, "/rpc/rate/5xx/%v", module).Mark(c.is5xx())

	metricUtil.GetOrRegisterTimer("/rpc/duration/%v", module).UpdateSince(c.start)
	if c.success() {
		metricUtil.GetOrRegisterTimer("/rpc/duration/success/%v", module).UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/2xx/%v", module).UpdateSince(c.start)
	case c.is4xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/4xx/%v", module).UpdateSince(c.start)
	case c.is5xx():
		metricUtil.GetOrRegisterTimer("/rpc/duration/5xx/%v", module).UpdateSince(c.start)
	}
}

// Monitor metrics
type MonitorMetrics struct{}

func (m *MonitorMetrics) SyncOnceQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return metricUtil.GetOrRegisterTimer("/monitor/sync/once/success")
	}

	return metricUtil.GetOrRegisterTimer("/monitor/sync/once/failure")
}

func (m *MonitorMetrics) ConfirmQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return metricUtil.GetOrRegisterTimer("/monitor/confirm/success")
	}

	return metricUtil.GetOrRegisterTimer("/monitor/confirm/failure")
}

// Store metrics
type StoreMetrics struct{}

func (*StoreMetrics) UpsertBillQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return metricUtil.GetOrRegisterTimer("/store/sqlite/upsertBill/success")
	}

	return metricUtil.GetOrRegisterTimer("/store/sqlite/upsertBill/failure")
}

// Worker metrics
type WorkerMetrics struct{}

func (m *WorkerMetrics) PollOnceQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return metricUtil.GetOrRegisterTimer("/worker/poll/once/success")
	}

	return metricUtil.GetOrRegisterTimer("/worker/poll/once/failure")
}

func (m *WorkerMetrics) PollOnceSize() metrics.Histogram {
	return metricUtil.GetOrRegisterHistogram("/worker/poll/once/size")
}

func (m *WorkerMetrics) ConfirmOnceQps() metrics.Timer {
	return metricUtil.GetOrRegisterTimer("/worker/confirm/once")
}

func (m *WorkerMetrics) UpdateConfirmOnceSize(total, success, retry, reconfirm int) {
	metricUtil.GetOrRegisterHistogram("/worker/confirm/once/size").Update(int64(total))
	metricUtil.GetOrRegisterHistogram("/worker/confirm/once/success/size").Update(int64(success))
	metricUtil.GetOrRegisterHistogram("/worker/confirm/once/retry/size").Update(int64(retry))
	metricUtil.GetOrRegisterHistogram("/worker/confirm/once/reconfirm/size").Update(int64(reconfirm))
}

func (m *WorkerMetrics) SettleOnceQps() metrics.Timer {
	return metricUtil.GetOrRegisterTimer("/worker/settle/once")
}

func (m *WorkerMetrics) UpdateSettleOnceSize(total, success, failure int) {
	metricUtil.GetOrRegisterHistogram("/worker/settle/once/size").Update(int64(total))
	metricUtil.GetOrRegisterHistogram("/worker/settle/once/success/size").Update(int64(success))
	metricUtil.GetOrRegisterHistogram("/worker/settle/once/failure/size").Update(int64(failure))
}
