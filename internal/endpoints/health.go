package endpoints

import "net/http"
import BA "github.com/Couches/blog-aggregator"

func HealthCheck(w http.ResponseWriter, r *http.Request, config BA.AppConfiguration) {
  res := struct {
    Status string `json:"status"`
  } {
    Status: "ok",
  }
  BA.RespondWithJSON(w, http.StatusOK, res)
}
