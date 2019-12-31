package util

import (
	"encoding/json"
	"net/http"
)

var (
	AUTH_HEADER_KEY = "tokenApi"
	PUBLIC_HEADER_KEY = "publicApi"
)

func RespondJSONError(w http.ResponseWriter, code int, msg interface{}) {
	RespondJSON(w, http.StatusOK, map[string]interface{}{
		"status": code,
		"msg":    msg,
	})
}

func RespondJSONSuccess(w http.ResponseWriter, data interface{}) {

	RespondJSON(w, http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   data,
	})
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
