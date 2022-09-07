package metrics

import (
	"fmt"

	mutil "github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/ethereum/go-ethereum/metrics"
)

func GetOrRegisterCounter(nameFormat string, nameArgs ...interface{}) metrics.Counter {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterCounter(name)
}

func GetOrRegisterGauge(nameFormat string, nameArgs ...interface{}) metrics.Gauge {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterGauge(name)
}

func GetOrRegisterGaugeFloat64(nameFormat string, nameArgs ...interface{}) metrics.GaugeFloat64 {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterGaugeFloat64(name)
}

func GetOrRegisterMeter(nameFormat string, nameArgs ...interface{}) metrics.Meter {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterMeter(name)
}

func NewHistogram() metrics.Histogram {
	return metrics.NewHistogram(metrics.NewExpDecaySample(1024, 0.015))
}

func GetOrRegisterHistogram(nameFormat string, nameArgs ...interface{}) metrics.Histogram {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterHistogram(name)
}

func GetOrRegisterTimer(nameFormat string, nameArgs ...interface{}) metrics.Timer {
	name := fmt.Sprintf(nameFormat, nameArgs...)
	return mutil.GetOrRegisterTimer(name)
}
