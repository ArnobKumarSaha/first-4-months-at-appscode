package auth

import (
	"context"
	"fmt"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"github.com/golang-jwt/jwt"
	"net/http"
)


type AuthProductResource struct {
	*schemas.ProductResource
}

// type AuthProductResource schemas.ProductResource

func (rs *AuthProductResource) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of errorhandler, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		fmt.Println("Printing 1 : " , tknStr , c)
		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an errorhandler
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		fmt.Println("Printing 2 : " , tkn, claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}


		//ctx := context.WithValue(r.Context(), "prod_id", chi.URLParam(r, "prod_id"))
		ctx := context.WithValue(r.Context(), "token", tknStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
