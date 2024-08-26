package prom

import (
	"github.com/dapr-platform/common"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var liveGauge prometheus.Gauge
var alarmCounter *prometheus.CounterVec

func init() {
	register()
	go loopWriteLive()
}

func register() {
	common.Logger.Debugln("register live gauge")
	name := "live"
	liveGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "iot",
		Subsystem: "service",
		Name:      name,
		Help:      "",
	})
	alarmCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "iot",
		Subsystem: "service",
		Name:      "alarm_total",
		Help:      "",
	}, []string{"level"})

	prometheus.MustRegister(liveGauge, alarmCounter)
	common.Logger.Debugln("register live gauge,send counter success")
}
func CountAlarm(level string) {
	alarmCounter.WithLabelValues(level).Inc()
}

func loopWriteLive() {
	for {

		time.Sleep(time.Second * 15)
		liveGauge.Set(1)
	}
}
