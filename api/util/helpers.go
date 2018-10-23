package util

import (
	"encoding/json"
	"net/http"
)

// SendAsJSON attempts to encode `v` and send it through `w` as JSON,
// while remembering and resetting the old Content-Type.
func SendAsJSON(w http.ResponseWriter, v interface{}) error {
	contentType := w.Header().Get("Content-Type")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	w.Header().Set("Content-Type", contentType)
	return err
}
