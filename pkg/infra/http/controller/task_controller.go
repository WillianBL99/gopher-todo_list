package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	usecase "github.com/willianbl99/todo-list_api/pkg/application/usecase/task"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
)

type TaskController struct {
	Providers struct {
		SaveTask    usecase.SaveTask
		UpdateTask  usecase.UpdateTask
		GetAllTasks usecase.GetAllTasks
	}
	Repository repository.TaskRepository
}

func NewTaskController(r repository.TaskRepository) *TaskController {
	tc := &TaskController{}
	tc.Providers.SaveTask = usecase.SaveTask{r}
	tc.Providers.UpdateTask = usecase.UpdateTask{r}
	tc.Providers.GetAllTasks = usecase.GetAllTasks{r}

	return tc
}

func (tc *TaskController) SaveTask(w http.ResponseWriter, r *http.Request) {
	st := dto.SaveTaskBodyRequest{}

	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := st.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uid := r.Context().Value(dto.UserId).(string)

	err := tc.Providers.SaveTask.Execute(st.Title, st.Description, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Task created successfully")
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ut := dto.UpdateTaskBodyRequest{}

	if err := json.NewDecoder(r.Body).Decode(&ut); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ut.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Task updated successfully")
}

func (tc *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(dto.UserId).(string)

	tasks, err := tc.Providers.GetAllTasks.Execute(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
