package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/uilianlago/API-To-Do-List/application/usecases"
	repositories "github.com/uilianlago/API-To-Do-List/tests/repositories/inmemory"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", MainHandler)

	repository := repositories.InMemoryTaskRepository{}
	fmt.Println("repository created")

	tl1, _ := repository.GetAll()
	fmt.Printf("Length : %d\n", len(tl1))

	st := usecases.SaveTask{Repository: &repository}
	fmt.Println("usecase created")

	st.Execute("Title1", "Describe1", time.Now().AddDate(0, 0, 2))
	st.Execute("Title2", "Describe2", time.Now().AddDate(0, 0, 3))
	fmt.Println("task saved")

	tl2, _ := repository.GetAll()
	fmt.Printf("Length : %d\n", len(tl2))

	showTaskList(tl2)
	fmt.Printf("is zero value: %t\n", tl2[0].CreatedAt.IsZero())

	http.ListenAndServe(":5050", nil)
}

func showTaskList(tl []entity.Task) {
	for _, t := range tl {
		fmt.Println()
		fmt.Printf("ID: %s\n", t.Id.String())
		fmt.Printf("Title: %s\n", t.Title)
		fmt.Printf("Describe: %s\n", t.Describe)
		fmt.Printf("DueDate: %s\n", t.DueDate.Format("2006-01-02 15:04:05"))
		fmt.Printf("Done: %t\n", t.Done)
		fmt.Printf("Deleted at: %s\n", t.DeletedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Created at: %s\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated at: %s\n", t.UpdateAt.Format("2006-01-02 15:04:05"))
		fmt.Println("========================================")
	}
}
