package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	usecase "github.com/willianbl99/todo-list_api/pkg/application/usecase/task"
	"github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
)

type TaskController struct {
	Providers struct {
		SaveTask         usecase.SaveTask
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
		herr.BadBodyRequest(w, err)
		return
	}

	uid := r.Context().Value(dto.UserId).(string)

	err := tc.Providers.SaveTask.Execute(st.Title, st.Description, uid)
	if err != nil {
		herr.AppToHttp(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Task created successfully")
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ut := dto.UpdateTaskBodyRequest{}

	if err := dto.ToDTO(r, &ut); err != nil {
		herr.BadBodyRequest(w, err)
		return
	}

	tctx := r.Context().Value(dto.TaskCtx).(*dto.TaskValue)

	if tctx.Get(dto.TaskTitle) == "" {
		tctx.Set(dto.TaskTitle, ut.Title)
	}

	if tctx.Get(dto.TaskDescription) == "" {
		tctx.Set(dto.TaskDescription, ut.Description)
	}

	if err := tc.Providers.UpdateTask.Execute(
		tctx.Get(dto.TaskId),
		tctx.Get(dto.TaskTitle),
		tctx.Get(dto.TaskDescription),
	); err != nil {
		herr.AppToHttp(w, err)
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
	var err error

	if status != "" {
		tasks, err = tc.Providers.GetTasksByStatus.Execute(uid, status)
	} else {
		tasks, err = tc.Providers.GetAllTasks.Execute(uid)
	}

	if err != nil {
		herr.AppToHttp(w, err)
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
		herr.AppToHttp(w, err)
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
			herr.AppToHttp(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "Task moved successfully")
	}
}
