package metrics

import (
	"fmt"

	"github.com/ethereum/go-ethereum/metrics"
)

var Web3PayRegistry = metrics.NewRegistry()

func GetOrRegisterCounter(nameFormat string, nameArgs ...interface{}) metrics.Counter {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return metrics.GetOrRegisterCounter(name, Web3PayRegistry)
}

func GetOrRegisterGauge(nameFormat string, nameArgs ...interface{}) metrics.Gauge {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return metrics.GetOrRegisterGauge(name, Web3PayRegistry)
}

func GetOrRegisterGaugeFloat64(nameFormat string, nameArgs ...interface{}) metrics.GaugeFloat64 {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return metrics.GetOrRegisterGaugeFloat64(name, Web3PayRegistry)
}

func GetOrRegisterMeter(nameFormat string, nameArgs ...interface{}) metrics.Meter {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return metrics.GetOrRegisterMeter(name, Web3PayRegistry)
}

func NewHistogram() metrics.Histogram {
	return metrics.NewHistogram(metrics.NewExpDecaySample(1024, 0.015))
}

func GetOrRegisterHistogram(nameFormat string, nameArgs ...interface{}) metrics.Histogram {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return Web3PayRegistry.GetOrRegister(name, NewHistogram).(metrics.Histogram)
}

func GetOrRegisterTimer(nameFormat string, nameArgs ...interface{}) metrics.Timer {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return metrics.GetOrRegisterTimer(name, Web3PayRegistry)
}
