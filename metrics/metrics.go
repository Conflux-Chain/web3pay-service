package metrics

var (
	RPC RpcMetrics
)

// RPC metrics
type RpcMetrics struct{}

func (*RpcMetrics) UpdateWithCollector(c *RpcCollector) {
	// Overall rate statistics
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/success").Mark(c.success())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/bizErr").Mark(c.isBizError())

	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/2xx").Mark(c.is2xx())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/4xx").Mark(c.is4xx())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/5xx").Mark(c.is5xx())

	GetOrRegisterTimer("web3pay/rpc/duration/all").UpdateSince(c.start)
	if c.success() {
		GetOrRegisterTimer("web3pay/rpc/duration/success").UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		GetOrRegisterTimer("web3pay/rpc/duration/2xx").UpdateSince(c.start)
	case c.is4xx():
		GetOrRegisterTimer("web3pay/rpc/duration/4xx").UpdateSince(c.start)
	case c.is5xx():
		GetOrRegisterTimer("web3pay/rpc/duration/5xx").UpdateSince(c.start)
	}

	module := c.module()
	if len(module) == 0 {
		return
	}

	// RPC rate statistics
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/success/%v", module).Mark(c.success())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/bizErr/%v", module).Mark(c.isBizError())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/2xx/%v", module).Mark(c.is2xx())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/4xx/%v", module).Mark(c.is4xx())
	GetOrRegisterTimeWindowPercentageDefault("web3pay/rpc/rate/5xx/%v", module).Mark(c.is5xx())

	GetOrRegisterTimer("web3pay/rpc/duration/%v", module).UpdateSince(c.start)
	if c.success() {
		GetOrRegisterTimer("web3pay/rpc/duration/success/%v", module).UpdateSince(c.start)
	}

	switch {
	case c.is2xx():
		GetOrRegisterTimer("web3pay/rpc/duration/2xx/%v", module).UpdateSince(c.start)
	case c.is4xx():
		GetOrRegisterTimer("web3pay/rpc/duration/4xx/%v", module).UpdateSince(c.start)
	case c.is5xx():
		GetOrRegisterTimer("web3pay/rpc/duration/5xx/%v", module).UpdateSince(c.start)
	}
}
