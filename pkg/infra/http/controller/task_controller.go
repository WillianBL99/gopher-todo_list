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
		SaveTask usecase.SaveTask
	}
	Repository repository.TaskRepository
}

func NewTaskController(r repository.TaskRepository) *TaskController {
	tc := &TaskController{}
	tc.Providers.SaveTask = usecase.SaveTask{r}

	return tc
}

func (tc *TaskController) SaveTask(w http.ResponseWriter, r *http.Request) {
	st := dto.SaveTask{}

	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := st.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uid := r.Context().Value("userId").(string)
	
	err := tc.Providers.SaveTask.Execute(st.Title, st.Description, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Task created successfully")
}