package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)


type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`

}

func APIResponse(message string, code int, status string, data interface{}) Response {
	
	fmt.Println(data)

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,

	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {

	var errors []string //untuk buat list yang nampung error

	for _, checkError := range err.(validator.ValidationErrors) {

		errors = append(errors, checkError.Error())
	}
	return errors
}

// func DateValidator(date string)string{

// 	var errors string
// 	re := regexp.MustCompile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")//19 20 mean 2 didepan 19xx dan 20xx

// 	CheckFormat := re.MatchString(date)

// 	if !CheckFormat {

// 		errors = "Please Check Your Date Format Correctly"
// 		return errors
// 	}

// 	GetCurrentDate := time.Now().Local()

// 	if date > GetCurrentDate.String(){

// 		errors = "Your Date Cannot More Than Current Date"
// 		return errors

// 	}

// 	return errors

// }