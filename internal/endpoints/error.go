package endpoints

import "net/http"
import BA "github.com/Couches/blog-aggregator"

func Error(w http.ResponseWriter, r *http.Request, config BA.AppConfiguration) {
  res := struct {
    Error string `json:"error"`
  } {
    Error: "Internal Server Error",
  }
  BA.RespondWithJSON(w, http.StatusInternalServerError, res)
}
