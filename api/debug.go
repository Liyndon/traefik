package api

import (
	"expvar"
	"fmt"
	"net/http"
	"runtime"

	"github.com/containous/mux"
)

func init() {
	expvar.Publish("Goroutines", expvar.Func(goroutines))
}

func goroutines() interface{} {
	return runtime.NumGoroutine()
}

// DebugHandler expose debug routes
type DebugHandler struct{}

// AddRoutes add debug routes on a router
func (g DebugHandler) AddRoutes(router *mux.Router) {
	router.Methods("GET").Path("/debug/vars").
		HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, "{\n")
			first := true
			expvar.Do(func(kv expvar.KeyValue) {
				if !first {
					fmt.Fprint(w, ",\n")
				}
				first = false
				fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
			})
			fmt.Fprint(w, "\n}\n")
		})
}
