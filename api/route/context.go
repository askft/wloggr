package route

import (
	"net/http"
)

// Context key used to find values in a context.Context
type contextKey int

const (
	_ contextKey = iota
	ckEmail
	ckUserID
)

func ctxString(r *http.Request, key contextKey) string {
	return r.Context().Value(key).(string)
}
