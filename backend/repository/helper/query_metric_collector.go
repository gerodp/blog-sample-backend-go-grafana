package helper

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type QueryMetricCollector struct {
	summary *prometheus.HistogramVec
}

func NewQueryCollector(repositoryName string) *QueryMetricCollector {

	s := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    fmt.Sprintf("%s_query_duration", repositoryName),
		Help:    fmt.Sprintf("A histogram of the %s repository query durations in seconds.", repositoryName),
		Buckets: prometheus.ExponentialBuckets(0.0005, 2, 7),
	}, []string{"method"})

	prometheus.MustRegister(*s)

	return &QueryMetricCollector{
		summary: s,
	}
}

func (q *QueryMetricCollector) Start() time.Time {
	return time.Now()
}

func (q *QueryMetricCollector) End(method string, start time.Time) {
	duration := time.Since(start)
	q.summary.With(prometheus.Labels{"method": method}).Observe(duration.Seconds())
}
