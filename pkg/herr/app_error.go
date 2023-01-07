package herr

import (
	"errors"
	"fmt"
)

type funcAppError error

type AppError struct {
	UserNotFound           funcAppError
	TaskNotFound           funcAppError
	InternalServerError    funcAppError
	Conflict               funcAppError
	BadRequest             funcAppError
	EmailAlreadyExists     funcAppError
	InvalidToken					 funcAppError
	InvalidTaskId          funcAppError
	InvalidTaskStatus      funcAppError
	InvalidUserId          funcAppError
	EmailOrPasswordInvalid funcAppError
}

func NewApp() *AppError {
	return &AppError{
		UserNotFound:           defaultAppError(User_Not_Found, "User not found, please check the params and try again."),
		TaskNotFound:           defaultAppError(Task_Not_Found, "Task not found, please check the params and try again."),
		InternalServerError:    defaultAppError(Internal_Server_Error, "Something went wrong, please try again later."),
		Conflict:               defaultAppError(Conflict, "Your request is conflict, please check your request and try again."),
		BadRequest:             defaultAppError(Bad_Request, "Your request is invalid, please check your request and try again."),
		EmailAlreadyExists:     defaultAppError(Email_Already_Exists, "Email already exists, please try again with another email."),
		InvalidToken:           defaultAppError(Invalid_Token, "Invalid token, please check the params and try again."),
		InvalidTaskId:          defaultAppError(Invalid_Task_Id, "Invalid task id, please check the params and try again."),
		InvalidTaskStatus:      defaultAppError(Invalid_Task_Status, "Invalid task status, please check the params and try again."),
		InvalidUserId:          defaultAppError(Invalid_User_Id, "Invalid user id, please check the params and try again."),
		EmailOrPasswordInvalid: defaultAppError(Email_Or_Password_Invalid, "Email or password invalid, please check the params and try again."),
	}
}

func GenAppError(ms string, dcp string) error {
	return defaultAppError(ms, dcp)
}

func defaultAppError(ms string, dcp string) error {
	return errors.New(fmt.Sprintf("%s$:%s$:", ms, dcp))
}
