package metrics

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
)

const (
	CollectKeyRPCModule  = "ckModule"
	CollectKeyRPCError   = "ckError"
	CollectKeyStatusCode = "ckStatusCode"
)

type Collector interface {
	Collect(key string, val interface{})
	BatchCollect(kvs map[string]interface{})
}

type RpcCollector struct {
	start time.Time
	data  map[string]interface{}
}

func NewRpcCollector() *RpcCollector {
	return &RpcCollector{
		start: time.Now(),
		data:  make(map[string]interface{}),
	}
}

func (c *RpcCollector) Collect(key string, val interface{}) {
	c.data[key] = val
}

func (c *RpcCollector) BatchCollect(kvs map[string]interface{}) {
	for k, v := range kvs {
		c.data[k] = v
	}
}

func (c *RpcCollector) module() string {
	if m, ok := c.data[CollectKeyRPCModule].(string); ok {
		return m
	}

	return ""
}

func (c *RpcCollector) error() error {
	if e, ok := c.data[CollectKeyRPCError].(error); ok {
		return e
	}

	return nil
}

func (c *RpcCollector) statusCode() int {
	if c, ok := c.data[CollectKeyStatusCode].(int); ok {
		return c
	}

	return 0
}

func (c *RpcCollector) success() bool {
	if util.IsInterfaceValNil(c.error()) {
		return true
	}

	if be, ok := model.IsBusinessError(c.error()); ok {
		return be.IsNil()
	}

	return true
}

func (c *RpcCollector) isBizError() bool {
	_, ok := model.IsBusinessError(c.error())
	return ok
}

func (c *RpcCollector) is2xx() bool {
	statusCode := c.statusCode()
	return statusCode == 0 || (statusCode >= 200 && statusCode < 300)
}

func (c *RpcCollector) is4xx() bool {
	statusCode := c.statusCode()
	return statusCode >= 400 && statusCode < 500
}

func (c *RpcCollector) is5xx() bool {
	statusCode := c.statusCode()
	return statusCode >= 500 && statusCode < 600
}
