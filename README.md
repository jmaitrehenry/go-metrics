# Go Metrics

Go Metrics is an utility package for creating prometheus collectors
and have a default set of GO and host metrics.

## Usage

```go
package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/jmaitrehenry/gometrics

	"log"
	"net/http"
	"time"
)

func main() {
	metrics, promMetrics := NewAppCollector()
	prometheus.MustRegister(promMetrics)
	
	// every x time, we refresh our internal gauge.
	go func() {
		ticker := time.Tick(1 * time.Second)
		for {
			select {
			case <-ticker:
				metrics.Collect(m)
			}
		}
	}()
	
	http.Handle("/metrics", promhttp.Handler())
	
	log.Printf("Starting metric server on %s\n", ":9090")
	http.ListenAndServe(":9090", nil)
}

func NewAppCollector() (*AppMetrics, *gometrics.PrometheusMetricsCollector) {
	ns := "something"
	system := "api"

	metrics := NewAppMetrics(ns, system)
	c := gometrics.Collector{
		Namespace: ns,
		Subsystem: system,
		EnableCPU: true,
		EnableMem: true,
	}

	return metrics, gometrics.NewPrometheusMetrics(c, metrics)
}


type AppMetrics struct {
	RequestCount prometheus.Gauge
	ResponseCodeCount *prometheus.GaugeVec
	
	metrics map[string]gometrics.MetricInfo
	labels  []string
}

// New AppMetrics initialize the metrics and also create metric definitions
func NewAppMetrics(ns, subsystem string) *AppMetrics {
	return &AppMetrics{
		RequestCount:  prometheus.NewGauge(prometheus.GaugeOpts{
            Namespace: ns,
            Subsystem: subsystem,
            Name:      "requests_count",
        }),
        ResponseCodeCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Subsystem: subsystem,
			Name:      "response_code_count",
		}, []string{"code"}),
		metrics: map[string]gometrics.MetricInfo{
			"requests_count": gometrics.NewMetric(ns, subsystem, "requests_count", "", prometheus.CounterValue, nil, []string{}),
			"response_code_count": gometrics.NewMetric(ns, subsystem, "response_code_count", "", prometheus.GaugeValue, nil, []string{"code"}),
		},
	}
}

// GetMetrics is an internal function use by gometrics
func (c AppMetrics) GetMetrics() map[string]gometrics.MetricInfo {
	return c.metrics
}

// CollectStats is call by the metric endpoint and it's use by the prometheus package to show the actual metrics with values
func (c AppMetrics) CollectStats() map[string]map[string]gometrics.MetricValue {
	stats := make(map[string]map[string]gometrics.MetricValue)

	stats["requests_count"] = map[string]gometrics.MetricValue{"default": gometrics.MetricValue{Collector: requestCounter.Collect}}
	for _, label := range c.labels {
		stats["response_code_count"][label] = gometrics.MetricValue{Collector: c.ResponseCodeCount.WithLabelValues(label).Collect}
	}
	return stats
}

// Collect will update the internal metric with the new value
func (m *AppMetrics) Collect() {
	m.RequestCount.Set(float64(64))
	m.labels = []string{"200", "500"}
	m.ResponseCodeCount.WithLabelValues("200").Set(float64(3))
	m.ResponseCodeCount.WithLabelValues("500").Set(float64(56))
}
```
