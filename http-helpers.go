package blogaggregator

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  res, err := json.Marshal(payload)
  if err != nil {
    code = http.StatusInternalServerError
    res = []byte(err.Error())
  }

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(res)
}

func ResponseWithError(w http.ResponseWriter, code int, err error) {
  RespondWithJSON(w, code, err)
}
