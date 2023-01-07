package herr

import (
	"encoding/json"
	"net/http"
)

type funcHttpError func(w http.ResponseWriter)

type HttpError struct {
	BadRequest             funcHttpError
	UserNotFound           funcHttpError
	TaskNotFound           funcHttpError
	InternalServerError    funcHttpError
	Conflict               funcHttpError
	Unauthorized           funcHttpError
	EmailAlreadyExists     funcHttpError
	InvalidToken           funcHttpError
	InvalidTaskId          funcHttpError
	InvalidTaskStatus      funcHttpError
	InvalidUserId          funcHttpError
	EmailOrPasswordInvalid funcHttpError
}

func NewHttp() *HttpError {
	return &HttpError{
		BadRequest:             defaultHttpError(http.StatusBadRequest, Bad_Request, "Your request is invalid, please check your request and try again."),
		UserNotFound:           defaultHttpError(http.StatusNotFound, User_Not_Found, "User not found, please check your request and try again."),
		TaskNotFound:           defaultHttpError(http.StatusNotFound, Task_Not_Found, "Task not found, please check your request and try again."),
		InternalServerError:    defaultHttpError(http.StatusInternalServerError, Internal_Server_Error, "Something went wrong, please try again later."),
		Conflict:               defaultHttpError(http.StatusConflict, Conflict, "Your request is conflict, please check your request and try again."),
		Unauthorized:           defaultHttpError(http.StatusUnauthorized, Unauthorized, "You are not authorized to access this resource."),
		EmailAlreadyExists:     defaultHttpError(http.StatusConflict, Email_Already_Exists, "Email already exists, please try again with another email."),
		InvalidToken:           defaultHttpError(http.StatusUnauthorized, Invalid_Token, "Invalid token, please check your token and try again."),
		InvalidTaskId:          defaultHttpError(http.StatusBadRequest, Invalid_Task_Id, "Invalid task id, please check your request and try again."),
		InvalidTaskStatus:      defaultHttpError(http.StatusBadRequest, Invalid_Task_Status, "Invalid task status, please check your request and try again."),
		InvalidUserId:          defaultHttpError(http.StatusBadRequest, Invalid_User_Id, "Invalid user id, please check your request and try again."),
		EmailOrPasswordInvalid: defaultHttpError(http.StatusUnauthorized, Email_Or_Password_Invalid, "Email or password invalid, please check your request and try again."),
	}
}

func GenHttpError(w http.ResponseWriter, stc int, ms string, dcp string) {
	w.Header().Set("Content-Type", "application/json")
	em := errorResponse{
		Message:     ms,
		Description: dcp,
	}

	w.WriteHeader(stc)
	err := json.NewEncoder(w).Encode(em)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func BadBodyRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	GenHttpError(w, http.StatusBadRequest, Bad_Request, err.Error())
}

func defaultHttpError(stc int, ms string, dcp string) func(w http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		w.WriteHeader(stc)
		GenHttpError(w, stc, ms, dcp)
	}
}
