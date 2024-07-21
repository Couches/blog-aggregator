package endpoints

import (
	"fmt"
	"net/http"
  BA "github.com/Couches/blog-aggregator"
)

type Endpoint struct {
  Method string
  Route string
  Callback func(http.ResponseWriter, *http.Request, BA.AppConfiguration)
}

func ConfigureEndpoints(mux *http.ServeMux, config BA.AppConfiguration) {
  for _, endpoint := range endpoints {
    mux.HandleFunc(
      fmt.Sprintf("%v %v", endpoint.Method, endpoint.Route), 
      methodHandler(endpoint, endpoint.Callback, config))
  }
}

func methodHandler(endpoint Endpoint, handlerFunc func(http.ResponseWriter, *http.Request, BA.AppConfiguration), config BA.AppConfiguration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != endpoint.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
			return
		}

		handlerFunc(w, r, config)
	}
}

var endpoints []Endpoint = []Endpoint {
  Endpoint {
    Method: "GET",
    Route: "/v1/healthz",
    Callback: HealthCheck,
  },
  Endpoint {
    Method: "GET",
    Route: "/v1/err",
    Callback: Error,
  },
}

