package middleware

import (
	"ProjectsGo/pkg/utils/response"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
	"runtime/debug"
	"strings"
)

func ResponseHandler(c *fiber.Ctx) error {
	err := c.Next()
	statusCode := c.Response().StatusCode()

	var responseBody interface{}

	if err != nil {
		var customErr *response.CustomError
		if errors.As(err, &customErr) {
			if customErr.StackTrace != nil {
				stacktrace := fmt.Sprintf("Stacktrace error: %+v\n", customErr.StackTrace)
				log.Error(stacktrace)
			}

			responseBody = customErr
			statusCode = customErr.StatusCode
		} else {
			responseBody = err.Error()
		}
	} else {
		err = json.Unmarshal(c.Response().Body(), &responseBody)
		if err != nil {
			stacktrace := errors.WithStack(err)

			log.Error(stacktrace)
			customErr := response.NewCustomError(response.ErrMsgParseError, response.ErrBadRequest, fiber.StatusBadRequest, nil, stacktrace, nil)

			responseBody = customErr
			statusCode = customErr.StatusCode
		}
	}

	return c.Status(statusCode).JSON(responseBody)
}

func ErrorHandler(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			stacktrace := GetStacktrace(r)
			log.Error(stacktrace)

			err := response.NewCustomError(response.ErrMsgInternalError, response.ErrInternal, fiber.StatusInternalServerError, nil, nil, nil)
			_ = c.Status(err.StatusCode).JSON(err)
		}
	}()

	return c.Next()
}

func GetStacktrace(message any) string {
	stack := string(debug.Stack())
	stackLines := strings.Split(stack, "\n")

	startIndex := 0
	for i, line := range stackLines {
		if strings.Contains(line, "panic") {
			startIndex = i
			break
		}
	}

	filteredStack := strings.Join(stackLines[startIndex:], "\n")
	return fmt.Sprintf("%v\nStacktrace: %s", message, filteredStack)
}
