package metrics

import (
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
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/success").Mark(c.success())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/bizErr").Mark(c.isBizError())

	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/2xx").Mark(c.is2xx())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/4xx").Mark(c.is4xx())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/5xx").Mark(c.is5xx())

	GetOrRegisterTimer("/rpc/duration/all").UpdateSince(c.start)
	if c.success() {
		GetOrRegisterTimer("/rpc/duration/success").UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		GetOrRegisterTimer("/rpc/duration/2xx").UpdateSince(c.start)
	case c.is4xx():
		GetOrRegisterTimer("/rpc/duration/4xx").UpdateSince(c.start)
	case c.is5xx():
		GetOrRegisterTimer("/rpc/duration/5xx").UpdateSince(c.start)
	}

	module := c.module()
	if len(module) == 0 {
		return
	}

	// RPC rate statistics
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/success/%v", module).Mark(c.success())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/bizErr/%v", module).Mark(c.isBizError())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/2xx/%v", module).Mark(c.is2xx())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/4xx/%v", module).Mark(c.is4xx())
	GetOrRegisterTimeWindowPercentageDefault("/rpc/rate/5xx/%v", module).Mark(c.is5xx())

	GetOrRegisterTimer("/rpc/duration/%v", module).UpdateSince(c.start)
	if c.success() {
		GetOrRegisterTimer("/rpc/duration/success/%v", module).UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		GetOrRegisterTimer("/rpc/duration/2xx/%v", module).UpdateSince(c.start)
	case c.is4xx():
		GetOrRegisterTimer("/rpc/duration/4xx/%v", module).UpdateSince(c.start)
	case c.is5xx():
		GetOrRegisterTimer("/rpc/duration/5xx/%v", module).UpdateSince(c.start)
	}
}

// Monitor metrics
type MonitorMetrics struct{}

func (m *MonitorMetrics) SyncOnceQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return GetOrRegisterTimer("/monitor/sync/once/success")
	}

	return GetOrRegisterTimer("/monitor/sync/once/failure")
}

func (m *MonitorMetrics) ConfirmQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return GetOrRegisterTimer("/monitor/confirm/success")
	}

	return GetOrRegisterTimer("/monitor/confirm/failure")
}

// Store metrics
type StoreMetrics struct{}

func (*StoreMetrics) UpsertBillQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return GetOrRegisterTimer("/store/sqlite/upsertBill/success")
	}

	return GetOrRegisterTimer("/store/sqlite/upsertBill/failure")
}

// Worker metrics
type WorkerMetrics struct{}

func (m *WorkerMetrics) PollOnceQps(err error) metrics.Timer {
	if util.IsInterfaceValNil(err) {
		return GetOrRegisterTimer("/worker/poll/once/success")
	}

	return GetOrRegisterTimer("/worker/poll/once/failure")
}

func (m *WorkerMetrics) PollOnceSize() metrics.Histogram {
	return GetOrRegisterHistogram("/worker/poll/once/size")
}

func (m *WorkerMetrics) ConfirmOnceQps() metrics.Timer {
	return GetOrRegisterTimer("/worker/confirm/once")
}

func (m *WorkerMetrics) UpdateConfirmOnceSize(total, success, retry, reconfirm int) {
	GetOrRegisterHistogram("/worker/confirm/once/size").Update(int64(total))
	GetOrRegisterHistogram("/worker/confirm/once/success/size").Update(int64(success))
	GetOrRegisterHistogram("/worker/confirm/once/retry/size").Update(int64(retry))
	GetOrRegisterHistogram("/worker/confirm/once/reconfirm/size").Update(int64(reconfirm))
}

func (m *WorkerMetrics) SettleOnceQps() metrics.Timer {
	return GetOrRegisterTimer("/worker/settle/once")
}

func (m *WorkerMetrics) UpdateSettleOnceSize(total, success, failure int) {
	GetOrRegisterHistogram("/worker/settle/once/size").Update(int64(total))
	GetOrRegisterHistogram("/worker/settle/once/success/size").Update(int64(success))
	GetOrRegisterHistogram("/worker/settle/once/failure/size").Update(int64(failure))
}
