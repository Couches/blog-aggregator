package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	BA "github.com/Couches/blog-aggregator"
	"github.com/Couches/blog-aggregator/internal/database"
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

  db, err := sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
  if err != nil {
    fmt.Println(err.Error())
  }

  dbQueries := database.New(db)
  config.Database.Queries = dbQueries

  fmt.Println(config)

  mux := http.NewServeMux()

  Endpoints.ConfigureEndpoints(mux, config)

  server := http.Server {
    Addr: config.Server.Address,
    Handler: mux,
  }

  server.ListenAndServe()
}
