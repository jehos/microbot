package microbot

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsController() http.Handler {
	return promhttp.Handler()
}
