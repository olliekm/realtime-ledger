package api

import (
	"log"
	"net/http"
)

type Auth struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (a *Auth) Populate() {
	a.tokenUsers["00000000"] = "user0"
	a.tokenUsers["aaaaaaaa"] = "userA"
	a.tokenUsers["05f717e5"] = "randomUser"
	a.tokenUsers["deadbeef"] = "user0"
}

func (a *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if user, ok := a.tokenUsers[token]; ok {
			// You can set the user in the context if needed
			log.Printf("Authenticated user: %s", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
