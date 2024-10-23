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
)

var requestCounter prometheus.Counter

func main() {
	c := gometrics.Collector{
		Namespace: "myproject",
		Subsystem: "myapp",
		EnableCPU: true,
		EnableMem: true,
	}
	
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "myproject",
		Subsystem: "myapp",
		Name:      "requests_count",
	})
	extraCollector := NewCollector("myproject", "myapp")
	
	prometheus.MustRegister(gometrics.NewPrometheusMetrics(c, extraCollector))
	http.Handle("/metrics", promhttp.Handler())
	
	log.Printf("Starting metric server on %s\n", ":9090")
	http.ListenAndServe(":9090", nil)
}


type Collector struct {
	metrics map[string]gometrics.MetricInfo
}

func NewCollector(namespace, subsystem string) Collector {
	return Collector{
		metrics: map[string]gometrics.MetricInfo{
			"requests_count": gometrics.NewMetric(namespace, subsystem, "requests_count", "", prometheus.CounterValue, nil, []string{}),
		},
	}
}

func (c Collector) GetMetrics() map[string]collector.MetricInfo {
	return c.metrics
}

func (c Collector) CollectStats() map[string]map[string]collector.MetricValue {
	stats := make(map[string]map[string]collector.MetricValue)

	stats["requests_count"] = map[string]collector.MetricValue{"default": collector.MetricValue{Collector: requestCounter.Collect}}

	return stats
}
```
