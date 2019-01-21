package microbot

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type (
	// MiddlewareConfig defines the config for Middleware middleware.
	MiddlewareConfig struct {
		// Skipper defines a function to skip middleware.
		// Skipper Skipper

		// Size of the stack to be printed.
		// Optional. Default value 4KB.
		StackSize int `yaml:"stack_size"`

		// DisableStackAll disables formatting stack traces of all other goroutines
		// into buffer after the trace for the current goroutine.
		// Optional. Default value false.
		DisableStackAll bool `yaml:"disable_stack_all"`

		// DisablePrintStack disables printing stack trace.
		// Optional. Default value as false.
		DisablePrintStack bool `yaml:"disable_print_stack"`
	}
)

var (
	// DefaultMiddlewareConfig is the default Middleware middleware config.
	DefaultMiddlewareConfig = MiddlewareConfig{
		// Skipper:           DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
	}
)

// Middleware returns a middleware which recovers from panics anywhere in the chain
// and handles the control to the centralized HTTPErrorHandler.
func Middleware() func(h http.Handler) http.Handler {
	return MiddlewareWithConfig(DefaultMiddlewareConfig)
}

// MiddlewareWithConfig returns a Middleware middleware with config.
// See: `Middleware()`.
func MiddlewareWithConfig(config MiddlewareConfig) func(h http.Handler) http.Handler {
	// Defaults
	// if config.Skipper == nil {
	// 	config.Skipper = DefaultMiddlewareConfig.Skipper
	// }
	if config.StackSize == 0 {
		config.StackSize = DefaultMiddlewareConfig.StackSize
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if config.Skipper(c) {
			// 	return next(c)
			// }

			sw := statusWriter{ResponseWriter: w}
			defer func(begun time.Time) {
				duration.Observe(time.Since(begun).Seconds())

				// hello_requests_total{status="200"} 2385
				requests.With(prometheus.Labels{
					"status": fmt.Sprint(sw.status),
				}).Inc()
			}(time.Now())

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						// c.Logger().Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
						fmt.Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					}
					// c.Error(err)
					fmt.Print(err)

					panics.Inc()
				}
			}()
			h.ServeHTTP(&sw, r)
		})
	}
}
