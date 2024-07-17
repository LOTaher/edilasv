package apis

import (
    "net/http"
    "context"

    "github.com/LOTaher/softbase/core"
)

// middleware for passing the database store across handlers
type storeContextKey struct {}

var StoreContextKey = storeContextKey{}

func DatabaseMiddleware(store *core.Store) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := context.WithValue(r.Context(), StoreContextKey, store)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// middleware for validating the API key
func KeyMiddleware(validKey string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            key := r.Header.Get("X-API-Key")
            if key == "" || key != validKey {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
