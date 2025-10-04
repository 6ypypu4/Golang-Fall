package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// Auth middleware checks X-API-Key and logs requests
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("used %s on %s", r.Method, r.URL.Path)
		//log.Printf("%s %s", r.Header, r.Body) //test

		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			log.Println(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
