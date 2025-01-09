package response

import "fmt"

type ValidationError struct {
	Field    string      `json:"field"`
	Message  string      `json:"message"`
	Code     string      `json:"code"`
	Expected interface{} `json:"expected,omitempty"`
	Received interface{} `json:"received,omitempty"`
}

const (
	ErrCodeInvalid       = "INVALID_INPUT"
	ErrCodeInvalidParam  = "INVALID_PARAM"
	ErrCodeRequiredField = "MISSING_FIELD"
	ErrCodeGeneric       = "BAD_REQUEST"
	ErrCodeInvalidLength = "INVALID_LENGTH"
	ErrCodeInvalidFormat = "INVALID_FORMAT"
)

var (
	ErrMsgInvalidField  = "The field '%s' is invalid"
	ErrMsgInvalidParam  = "The parameter '%s' is invalid"
	ErrMsgRequiredField = "The field '%s' is required"
	ErrMsgInvalidLength = "The '%s' provided is not in a valid length"
	ErrMsgInvalidFormat = "The '%s' provided is not in a valid format"
	ErrUniqueFieldInUse = "The field '%s' is unique"
)

func NewValidationError(field, message string, expected, received interface{}) ValidationError {
	var formatMessage string

	if field == "" {
		formatMessage = message
	} else {
		formatMessage = fmt.Sprintf(message, field)
	}

	return ValidationError{
		Code:     getMessageCode(message),
		Field:    field,
		Message:  formatMessage,
		Expected: expected,
		Received: received,
	}
}

func getMessageCode(message string) string {
	var code string
	switch message {
	case ErrMsgInvalidField:
		code = ErrCodeInvalid
		break
	case ErrMsgInvalidParam:
		code = ErrCodeInvalidParam
		break
	case ErrMsgInvalidLength:
		code = ErrCodeInvalidLength
		break
	case ErrMsgInvalidFormat:
		code = ErrCodeInvalidFormat
		break
	case ErrMsgRequiredField:
		code = ErrCodeRequiredField
		break
	default:
		code = ErrCodeGeneric
	}

	return code
}
