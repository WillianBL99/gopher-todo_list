package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	usecase "github.com/willianbl99/todo-list_api/internal/application/usecase/task"
	"github.com/willianbl99/todo-list_api/internal/infra/http/dto"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type TaskController struct {
	Providers struct {
		SaveTask         usecase.SaveTask
		GetTaskById      usecase.GetTaskById
		UpdateTask       usecase.UpdateTask
		GetAllTasks      usecase.GetAllTasks
		GetTasksByStatus usecase.GetTasksByStatus
		DeleteTask       usecase.DeleteTask
		MoveTask         usecase.MoveTask
	}
	Repository repository.TaskRepository
}

func NewTaskController(r repository.TaskRepository) *TaskController {
	tc := &TaskController{
		Repository: r,
	}
	tc.Providers.SaveTask = usecase.SaveTask{TaskRepository: r}
	tc.Providers.GetTaskById = usecase.GetTaskById{TaskRepository: r}
	tc.Providers.UpdateTask = usecase.UpdateTask{TaskRepository: r}
	tc.Providers.DeleteTask = usecase.DeleteTask{TaskRepository: r}
	tc.Providers.GetAllTasks = usecase.GetAllTasks{TaskRepository: r}
	tc.Providers.GetTasksByStatus = usecase.GetTasksByStatus{TaskRepository: r}
	tc.Providers.MoveTask = usecase.MoveTask{TaskRepository: r}

	return tc
}

func (tc *TaskController) SaveTask(w http.ResponseWriter, r *http.Request) {
	st := dto.SaveTaskBodyRequest{}

	if err := dto.ToDTO(r, &st); err != nil {
		e.New().InfraHttpErr(e.BadBodyRequest, err.Error()).ToHttp(w)
		return
	}

	uid := r.Context().Value(dto.UserId).(string)
	tk, err := tc.Providers.SaveTask.Execute(st.List, st.Title, st.Description, uid)
	if err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tk)
}

func (tc *TaskController) GetTaskById(w http.ResponseWriter, r *http.Request) {
	tctx := r.Context().Value(dto.TaskCtx).(*dto.TaskValue)
	tkid := tctx.Get(dto.TaskId)
	tk, err := tc.Providers.GetTaskById.Execute(tkid)
	if err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tk)
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ut := dto.UpdateTaskBodyRequest{}

	if err := dto.ToDTO(r, &ut); err != nil {
		e.New().InfraHttpErr(e.BadBodyRequest, err.Error()).ToHttp(w)
		return
	}

	tctx := r.Context().Value(dto.TaskCtx).(*dto.TaskValue)

	if err := tc.Providers.UpdateTask.Execute(tctx.Get(dto.TaskId), ut.List, ut.Title, ut.Description); err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "Task updated successfully")
}

func (tc *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(dto.UserId).(string)
	status := r.URL.Query().Get("status")
	var tasks []entity.Task
	var err *e.Error

	if status != "" {
		tasks, err = tc.Providers.GetTasksByStatus.Execute(uid, status)
	} else {
		tasks, err = tc.Providers.GetAllTasks.Execute(uid)
	}

	if err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	tctx := r.Context().Value(dto.TaskCtx).(*dto.TaskValue)
	tid := tctx.Get(dto.TaskId)

	if err := tc.Providers.DeleteTask.Execute(tid); err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Task deleted successfully")
}

func (tc *TaskController) MoveTask(st entity.Status) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context().Value(dto.TaskCtx).(*dto.TaskValue)
		tid := ctx.Get(dto.TaskId)

		if err := tc.Providers.MoveTask.Execute(tid, string(st)); err != nil {
			e.New().InfraHttpErrRec(err).ToHttp(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "Task moved successfully")
	}
}
