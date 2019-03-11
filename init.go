package microbot

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "microbot_http_request_duration_milliseconds",
			Help: "Summary of http request duration in milliseconds.",
		},
		[]string{"handler", "status", "method", "ip_type"},
	)

	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "microbot_http_request_total",
			Help: "Total number of http requests.",
		},
		[]string{"handler", "status", "method", "ip_type"},
	)

	panics = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "microbot_panic_total",
			Help: "Total number of panic.",
		})

	accessibility = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "microbot_db_ping_duration_nanoseconds",
			Help: "Summary of DB ping duration in nanoseconds.",
		},
		[]string{"status"},
	)
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
				for _, r := range PingDB() {
					status := "ok"
					if r.err != nil {
						status = "error"
					}
					accessibility.WithLabelValues(status).Observe(float64(r.duration))
				}
			}
		}
	}()
}
