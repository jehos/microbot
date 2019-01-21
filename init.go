package microbot

import "github.com/prometheus/client_golang/prometheus"

var (
	duration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "microbot_request_duration_seconds",
		Help:    "Histogram of the /hello request duration.",
		Buckets: []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})

	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "microbot_requests_total",
			Help: "Total number of all requests.",
		},
		[]string{"status"},
	)

	panics = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "microbot_panic_total",
			Help: "Total number of panic.",
		})

	accessibility = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "microbot_request_duration_seconds",
		Help:    "Histogram of the /hello request duration.",
		Buckets: []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
)

func init() {
	prometheus.MustRegister(duration)
	prometheus.MustRegister(requests)
	prometheus.MustRegister(panics)
}
