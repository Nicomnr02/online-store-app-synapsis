package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"online_app_store/model"
	"online_app_store/vars"

	"github.com/golang-jwt/jwt/v5"
)

func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("synapsis")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{Message: http.ErrNoCookie.Error()})
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key
			return vars.JWTKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{Message: http.ErrNoCookie.Error()})
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"]
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			fmt.Println("success! User ID : ", claims["user_id"])

		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{Message: http.ErrNoCookie.Error()})
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		}

	})

}
