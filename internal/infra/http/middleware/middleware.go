package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	"github.com/willianbl99/todo-list_api/internal/infra/http/dto"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
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
		tk, er := handleToken(r)
		if er != nil {
			e.New().InfraHttpErr(e.InvalidToken, er.Error()).ToHttp(w)
			return
		}

		jwt := server.NewJwtService()
		if !jwt.ValidateToken(tk) {
			e.New().InfraHttpErr(e.InvalidToken).ToHttp(w)
			return
		}

		uid, err := jwt.GetTokenId(tk)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			e.New().InfraHttpErr(e.InvalidToken, err.ToSubErr()...).ToHttp(w)
			return
		}

		puid, er := uuid.Parse(uid)
		if er != nil || puid == uuid.Nil {
			e.New().InfraHttpErr(e.InvalidToken, er.Error()).ToHttp(w)
			return
		} else {
			if _, err := m.UserRepository.GetById(puid); err != nil {
				e.New().InfraHttpErr(e.Unauthorized, err.ToSubErr()...).ToHttp(w)
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
				e.New().InfraHttpErr(e.InvalidId, err.Description).ToHttp(w)
				return
			}

			tv := dto.NewTaskValue()
			tv.Set(dto.TaskId, taskId)
			tv.Set(dto.TaskTitle, t.Title)
			tv.Set(dto.TaskDescription, t.Description)

			ctx = context.WithValue(r.Context(), dto.TaskCtx, tv)
		} else {
			e.New().InfraHttpErr(e.InvalidId, err.Error()).ToHttp(w)
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
	split := strings.Split(token, "Bearer ")
	if len(split) != 2 {
		return "", errors.New(e.InvalidToken.Title)
	}

	return split[1], nil
}
