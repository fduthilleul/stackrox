package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stackrox/rox/pkg/metrics"
)

var (
	containersStored = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metrics.PrometheusNamespace,
		Subsystem: metrics.SensorSubsystem.String(),
		Name:      "num_containers_in_entity_store",
		Help:      "A gauge to track the number of containers in the entity store",
	})
	historicalContainersStored = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metrics.PrometheusNamespace,
		Subsystem: metrics.SensorSubsystem.String(),
		Name:      "num_historical_containers_in_entity_store",
		Help:      "A gauge to track the number of historical containers in the entity store",
	})
)

// UpdateNumberContainersInEntityStored update number of containers stored
func UpdateNumberContainersInEntityStored(num int) {
	containersStored.Set(float64(num))
}

// UpdateNumberHistoricalContainersInEntityStored update number of containers stored
func UpdateNumberHistoricalContainersInEntityStored(num int) {
	historicalContainersStored.Set(float64(num))
}

func init() {
	prometheus.MustRegister(containersStored)
	prometheus.MustRegister(historicalContainersStored)
}
