package response

import (
	"encoding/json"
	"net/http"
)

// Ответ обернутый в json
func Json(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(data)
	if err != nil {
		body = []byte(
			`{"error": "Unknown error in encoding json!"}`)
	}

	w.Write(body)
}
