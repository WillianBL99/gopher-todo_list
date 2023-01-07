package herr

import (
	"fmt"
	"net/http"
	"strings"
)

type errorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

const (
	User_Not_Found            = "User Not Found"
	Task_Not_Found            = "Task Not Found"
	Internal_Server_Error     = "Internal Server Error"
	Conflict                  = "Conflict"
	Bad_Request               = "Bad Request"
	Unauthorized              = "Unauthorized"
	Email_Already_Exists      = "Email Already Exists"
	Invalid_Token             = "Invalid Token"
	Invalid_Task_Id           = "Invalid Task Id"
	Invalid_Task_Status       = "Invalid Task Status"
	Invalid_User_Id           = "Invalid User Id"
	Email_Or_Password_Invalid = "Email Or Password Invalid"
)

func AppToHttp(w http.ResponseWriter, err error) {
	hte := NewHttp()
	te := strings.Split(err.Error(), "$:")[0]

	switch te {
	case User_Not_Found:
		hte.UserNotFound(w)
		return
	case Task_Not_Found:
		hte.TaskNotFound(w)
		return
	case Internal_Server_Error:
		hte.InternalServerError(w)
		return
	case Conflict:
		hte.Conflict(w)
		return
	case Bad_Request:
		hte.BadRequest(w)
		return
	case Unauthorized:
		hte.Unauthorized(w)
		return
	case Email_Already_Exists:
		hte.EmailAlreadyExists(w)
		return
	default:
		GenHttpError(
			w,
			http.StatusInternalServerError,
			Internal_Server_Error,
			fmt.Sprintf("Error not mapped: %s - err: %s", err.Error(), te),
		)
		return
	}
}
