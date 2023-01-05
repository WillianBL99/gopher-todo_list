package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

func AuthHandler(u repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			at := r.Header.Values("Authorization")
			if len(at) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Should has authorization", http.StatusUnauthorized)
				return
			}
	
			tk, err := extractToken(at[0])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Should has token", http.StatusUnauthorized)
				return
			}
	
			jwt := server.NewJwtService()
			if !jwt.ValidateToken(tk) {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
	
			uid, err := jwt.GetTokenId(tk)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if puid, err := uuid.Parse(uid); err == nil && puid != uuid.Nil {
				_, err := u.GetById(puid)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
	
			ctx := context.WithValue(r.Context(), "userId", uid)
	
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractToken(s string) (string, error) {
	ssplit := strings.Split(s, "Bearer ")
	if len(ssplit) != 2 {
		return "", errors.New("Invalid token")
	}

	return ssplit[1], nil
}
