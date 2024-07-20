package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load()
  port := os.Getenv("PORT")

  mux := http.NewServeMux()

  server := http.Server {
    Addr: port,
    Handler: mux,
  }
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  res, err := json.Marshal(payload)
  if err != nil {
    code = http.StatusInternalServerError
    res = []byte(err.Error())
  }

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(res)
}

func responseWithError(w http.ResponseWriter, code int, err error) {
  respondWithJSON(w, code, err)
}
