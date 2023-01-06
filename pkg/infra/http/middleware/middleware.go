package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type Middleware struct {
	UserRepository repository.UserRepository
	TaskRepository repository.TaskRepository
}

func NewMiddleware(ur repository.UserRepository, tr repository.TaskRepository) *Middleware {
	return &Middleware{
		UserRepository: ur,
		TaskRepository: tr,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
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
			_, err := m.UserRepository.GetById(puid)
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

		ctx := context.WithValue(r.Context(), dto.UserId, uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) TaskParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		taskId := chi.URLParam(r, "taskId")
		var ctx context.Context

		if ptaskId, err := uuid.Parse(taskId); err == nil && ptaskId != uuid.Nil {
			t, err := m.TaskRepository.GetById(ptaskId)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}

			tv := dto.NewTaskValue()
			tv.Set(dto.TaskId, taskId)
			tv.Set(dto.TaskTitle, t.Title)
			tv.Set(dto.TaskDescription, t.Description)

			ctx = context.WithValue(r.Context(), dto.TaskCtx, tv)
		} else {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Invalid task id", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractToken(token string) (string, error) {
	splitt := strings.Split(token, "Bearer ")
	if len(splitt) != 2 {
		return "", fmt.Errorf("Invalid token")
	}

	return splitt[1], nil
}
