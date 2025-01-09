package response

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"strings"
)

type CustomError struct {
	Message    string             `json:"message"`
	Code       string             `json:"code"`
	StatusCode int                `json:"-"`
	StackTrace interface{}        `json:"-"`
	Errors     *[]ValidationError `json:"errors,omitempty"`
}

const (
	ErrCodeAlreadyExists = "ALREADY_EXISTS"
	ErrBadRequest        = "BAD_REQUEST"
	ErrParsingBody       = "PARSING_BODY"
	ErrParsingParams     = "PARSING_PARAMS"
	ErrInternal          = "INTERNAL_ERROR"
	ErrNotFound          = "NOT_FOUND"
	ErrTooManyRequests   = "TOO_MANY_REQUESTS"
)

const PgErrDuplicatedKey string = "23505"

var (
	ErrMsgNotFound         = "The register was not found in the database."
	ErrMsgInternalError    = "Unexpected error"
	ErrMsgBadRequest       = "An error occurred"
	ErrMsgParseError       = "The parameters sent cannot be processed."
	ErrMsgParseParamsError = "The body sent cannot be processed. Please check your request."
	ErrMsgAlreadyExisting  = "The submitted '%s' already exists"
	ErrMsgTooManyRequests  = "Too many requests. Please wait before trying again."
)

func NewCustomError(message, code string, statusCode int, errors *[]ValidationError, stacktrace, field interface{}) *CustomError {
	var formatMessage string

	if field == nil {
		formatMessage = message
	} else {
		formatMessage = fmt.Sprintf(message, field)
	}

	return &CustomError{
		Code:       code,
		Message:    formatMessage,
		StatusCode: statusCode,
		StackTrace: stacktrace,
		Errors:     errors,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewGormError(err error, stacktrace interface{}, entityName ...string) error {
	if IsPostgresRawError(err) {
		var pgError *pgconn.PgError
		ok := errors.As(err, &pgError)
		if !ok {
			return NewCustomError(ErrMsgInternalError, ErrInternal, fiber.StatusInternalServerError, nil, stacktrace, nil)
		}

		if pgError.Code == PgErrDuplicatedKey {
			message := err.Error()
			parts := strings.Split(err.Error(), "\"")
			if len(parts) > 1 {
				indexTableColumn := strings.Split(parts[1], "uni_")
				if len(indexTableColumn) > 1 {
					message = indexTableColumn[1] + " already exists"
				}
			}
			return NewCustomError(message, ErrBadRequest, fiber.StatusBadRequest, nil, nil, nil)
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		message := ""
		if len(entityName) > 0 {
			message += entityName[0]
		}
		message += " not found"
		return NewCustomError(message, ErrNotFound, fiber.StatusNotFound, nil, nil, nil)
	}

	return NewCustomError(ErrMsgInternalError, ErrInternal, fiber.StatusInternalServerError, nil, stacktrace, nil)
}

func IsPostgresRawError(err error) bool {
	var pgError *pgconn.PgError
	ok := errors.As(err, &pgError)
	return ok
}
