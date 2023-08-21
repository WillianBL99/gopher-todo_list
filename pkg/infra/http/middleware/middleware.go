package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
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
		tk, err := handleToken(r)
		if err != nil {
			herr.NewHttp().InvalidToken(w)
			return
		}

		jwt := server.NewJwtService()
		if !jwt.ValidateToken(tk) {
			herr.NewHttp().InvalidToken(w)
			return
		}

		uid, err := jwt.GetTokenId(tk)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		puid, err := uuid.Parse(uid)
		if err != nil || puid == uuid.Nil {
			herr.NewHttp().InvalidToken(w)
			return
		} else {
			if _, err := m.UserRepository.GetById(puid); err != nil {
				herr.NewHttp().Unauthorized(w)
				return
			}
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
				herr.AppToHttp(w, err)
				return
			}

			tv := dto.NewTaskValue()
			tv.Set(dto.TaskId, taskId)
			tv.Set(dto.TaskTitle, t.Title)
			tv.Set(dto.TaskDescription, t.Description)

			ctx = context.WithValue(r.Context(), dto.TaskCtx, tv)
		} else {
			herr.NewHttp().InvalidTaskId(w)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleToken(r *http.Request) (string, error) {
	at := r.Header.Get("Authorization")

	tk, err := extractToken(at)
	if err != nil {
		return "", err
	}

	return tk, nil
}

func extractToken(token string) (string, error) {
	splitt := strings.Split(token, "Bearer ")
	if len(splitt) != 2 {
		return "", errors.New(herr.Invalid_Token)
	}

	return splitt[1], nil
}
