package middlewares

import (
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/utils"
	"net/http"
)

func JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ** Get token from headers
		token := r.Header.Get("Authorization")

		// ** Check if the header has value
		if token == "" {
			core.RespondWithError(w, http.StatusUnauthorized, "Authorization token is missing")
			return
		}

		// ** Check if token is valid
		user, err := utils.ValidateToken(token)
		if err != nil {
			core.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// ** Attach user to context
		ctx := utils.SetUserInContext(r.Context(), user)

		// ** Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
