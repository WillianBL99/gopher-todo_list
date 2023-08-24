package herr

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Layer string
type ErrorBody struct {
	Title       string
	Description string
	Status      int
}

const (
	InfraHttp   Layer = "infra-http"
	InfraDb     Layer = "infra-db"
	InfraTest   Layer = "infra-test"
	Application Layer = "application"
)

var (
	BadRequest             ErrorBody = ErrorBody{"Bad Request", "Your request is invalid, please check your request and try again.", http.StatusBadRequest}
	BadBodyRequest         ErrorBody = ErrorBody{"Bad Body Request", "Your body request is invalid, please check your request and try again.", http.StatusBadRequest}
	UserNotFound           ErrorBody = ErrorBody{"User Not Found", "User not found, please check your request and try again.", http.StatusNotFound}
	TaskNotFound           ErrorBody = ErrorBody{"Task Not Found", "Task not found, please check your request and try again.", http.StatusNotFound}
	InternalServerError    ErrorBody = ErrorBody{"Internal Server Error", "Something went wrong, please try again later.", http.StatusInternalServerError}
	Conflict               ErrorBody = ErrorBody{"Conflict", "Your request is conflict, please check your request and try again.", http.StatusConflict}
	Unauthorized           ErrorBody = ErrorBody{"Unauthorized", "You are not authorized to access this resource.", http.StatusUnauthorized}
	EmailAlreadyExists     ErrorBody = ErrorBody{"Email Already Exists", "Email already exists, please try again with another email.", http.StatusConflict}
	InvalidToken           ErrorBody = ErrorBody{"Invalid Token", "Invalid token, please check your token and try again.", http.StatusUnauthorized}
	ExpiredToken           ErrorBody = ErrorBody{"Expired Token", "Expired token, please check your token and try again.", http.StatusUnauthorized}
	InvalidId              ErrorBody = ErrorBody{"Invalid Id", "Invalid id, please check your request and try again.", http.StatusBadRequest}
	InvalidStatus          ErrorBody = ErrorBody{"Invalid Status", "Invalid status, please check your request and try again.", http.StatusBadRequest}
	EmailOrPasswordInvalid ErrorBody = ErrorBody{"Email Or Password Invalid", "Email or password invalid, please check your request and try again.", http.StatusUnauthorized}
	EmptyField             ErrorBody = ErrorBody{"Empty Field", "Empty field, please check your request and try again.", http.StatusBadRequest}
	ErrorOnConnectDb       ErrorBody = ErrorBody{"Error On Connect Db", "Error on connect db, please check your request and try again.", http.StatusInternalServerError}
	ErrorOnStartServer     ErrorBody = ErrorBody{"Error On Start Server", "Error on start server, please check your request and try again.", http.StatusInternalServerError}
)

type Error struct {
	Layer       Layer
	Status      int
	Title       string
	Description string
	SubErrors   []string
}

func New() *Error {
	return &Error{}
}

func (e *Error) Error() *Error {
	return e
}

func (e *Error) SetLayer(layer Layer) *Error {
	e.Layer = layer
	return e
}

func (e *Error) CustomError(err ErrorBody, subErrors ...string) *Error {
	if e.Layer == "" {
		e.Layer = "Unset"
	}
	return &Error{
		Layer:       e.Layer,
		Status:      err.Status,
		Title:       err.Title,
		Description: err.Description,
		SubErrors:   subErrors,
	}
}

func (e *Error) NewErr(layer Layer, title, description string, subErrors ...error) *Error {
	return &Error{
		Layer:       layer,
		Status:      http.StatusInternalServerError,
		Title:       title,
		Description: description,
		SubErrors:   e.SubErrors,
	}
}

func (e *Error) ToHttp(w http.ResponseWriter) {
	if e == nil {
		panic("Error is nil")
	}
	fmt.Println(toLog(e))

	w.Header().Set("Content-Type", "application/json")
	st := e.Status
	if st == 0 {
		st = http.StatusHTTPVersionNotSupported
	}
	w.WriteHeader(st)
	message := fmt.Sprintf(`{"title": "%s", "description": "%s"}`, e.Title, e.Description)
	if e.Title == BadBodyRequest.Title {
		message = fmt.Sprintf(`{"title": "%s", "description": "%s"}`, e.Title, strings.Join(e.SubErrors, ","))
	}
	fmt.Fprintf(w, fmt.Sprintf(`Error: %s`, message))
}

func toLog(e *Error) string {
	tieFormat := "2006-01-02 15:04:05"
	currentTime := time.Now().Format(tieFormat)
	return fmt.Sprintf(
		"[%s] - Error:\n> Layer: %s, Status: %d, Title: %s\n  Description: %s\n  SubErrors: %v",
		currentTime, e.Layer, e.Status, e.Title, e.Description, e.SubErrors,
	)
}

func (e *Error) ToSubErr() []string {
	errs := []string{}
	errs = append(errs, fmt.Sprintf(`{"title": "%s", "description": "%s"}`, e.Title, e.Description))
	errs = append(errs, e.SubErrors...)
	return errs
}

func (e *Error) Fatal() {
	panic(toLog(e))
}

func (e *Error) AppErr(err ErrorBody, subErr ...string) *Error {
	e = e.CustomError(err, subErr...)
	return e.SetLayer(Application)
}

func (e *Error) InfraHttpErr(err ErrorBody, subErr ...string) *Error {
	e = e.CustomError(err, subErr...)
	return e.SetLayer(InfraHttp)
}

func (e *Error) InfraHttpErrRec(err *Error) *Error {
	e = err
	return e.SetLayer(InfraHttp)
}

func (e *Error) InfraDbErr(err ErrorBody, subErr ...string) *Error {
	e = e.CustomError(err, subErr...)
	return e.SetLayer(InfraDb)
}
