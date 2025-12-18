package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Здесь должна быть логика аутентификации
		// Например, проверка JWT токена или сессии
		next.ServeHTTP(w, r)
	})
}
