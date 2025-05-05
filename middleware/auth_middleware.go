package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/julienschmidt/httprouter"
)

type contextKey string

const userIDKey contextKey = "user_id"

// Middleware dengan pengecualian path tertentu
func NewAuthMiddlewareWithExclusion(next http.Handler, excludedPaths []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cek apakah path termasuk yang dikecualikan
		for _, path := range excludedPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Ambil token dari Authorization header
		tokenString := r.Header.Get("Authorization")

		// Jika tidak ada di header, coba dari cookie
		if tokenString == "" {
			cookie, err := r.Cookie("Authorization")
			if err == nil {
				tokenString = cookie.Value
			}
		}

		// Validasi format token
		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// Bersihkan prefix "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verifikasi token
		userID, err := helper.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Masukkan user_id ke dalam context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Fungsi untuk mengambil user_id dari context
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userIDKey).(int)
	return userID, ok
}
func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := helper.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Masukkan user_id ke dalam context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next(w, r.WithContext(ctx), ps)
	}
}
