package main

import (
	"fmt"
	"net/http"
  BA "github.com/Couches/blog-aggregator"
  Endpoints "github.com/Couches/blog-aggregator/internal/endpoints"

	"github.com/joho/godotenv"
  _ "github.com/lib/pq"
)

var configPath string = "settings.yaml"

func main() {
  godotenv.Load()

  config, err := BA.LoadAppConfiguration(configPath)
  if err != nil {
    fmt.Println("Failed to load app configurations at \"%v\" with error: %v", configPath, err)
  }

  mux := http.NewServeMux()

  Endpoints.ConfigureEndpoints(mux, config)

  server := http.Server {
    Addr: config.Server.Address,
    Handler: mux,
  }

  server.ListenAndServe()
}
