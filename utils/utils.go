package utils

import (
	"anangfirmansyahputra/portfolio-backend/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, response types.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"success":false,"message":"Internal Server Error"}`, http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, types.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}
