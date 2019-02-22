package microbot

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "microbot_request_duration_seconds",
			Help: "Histogram of all request duration.",
		},
		[]string{"handler", "status", "method"},
	)

	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "microbot_requests_total",
			Help: "Total number of all requests.",
		},
		[]string{"handler", "status", "method"},
	)

	panics = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "microbot_panic_total",
			Help: "Total number of panic.",
		})

	accessibility = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "microbot_db_ping_duration_seconds",
		Help:    "Histogram of DB ping request duration.",
		Buckets: []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
)

func init() {
	prometheus.MustRegister(duration)
	prometheus.MustRegister(requests)
	prometheus.MustRegister(panics)
	prometheus.MustRegister(accessibility)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				results := PingDB()
				for _, r := range results {
					accessibility.Observe(r.duration)
				}
			}
		}
	}()
}
