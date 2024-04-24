package apis

import (
    "net/http"
    "context"

    "github.com/LOTaher/softbase/core"
)

// Middleware for passing the database store across handlers

//////////////////////// ////////////////////// ////////////////////// //////////////////////
type StoreContextKey struct {}

func DatabaseMiddleware(store *core.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), StoreContextKey{}, store)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
//////////////////////// ////////////////////// ////////////////////// //////////////////////
