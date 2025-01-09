package validators

import (
	"ProjectsGo/pkg/utils/response"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func CamelToSnake(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	snake = regexp.MustCompile("_+").ReplaceAllString(snake, "_")
	return strings.ToLower(snake)
}

// TODO:REESTRUTURAR ESTE TRATAMENTO DE ERRO

func ValidateFields(err error) *[]response.ValidationError {
	var errList []response.ValidationError
	if err != nil {
		var message string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				message = response.ErrMsgRequiredField
			case "email", "excludesall=!@#?":
				message = response.ErrMsgInvalidFormat
			case "len", "min", "max":
				message = response.ErrMsgInvalidLength
			default:
				message = response.ErrMsgInvalidField
			}
			fieldName := CamelToSnake(err.Field())
			errList = append(errList, response.NewValidationError(strings.ToLower(fieldName), message, nil, nil))
		}
	}
	return &errList
}
