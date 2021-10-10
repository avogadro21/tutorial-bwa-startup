package global

import "github.com/go-playground/validator/v10"

type ResponseData struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(status string, code int, message string, data interface{}) ResponseData {
	response := ResponseData{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	return response
}

func ValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
