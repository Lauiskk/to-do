package company

import (
	"ProjectsGo/internal/database"
	"ProjectsGo/internal/entities/domain"
	"ProjectsGo/pkg/utils/response"
	"ProjectsGo/pkg/utils/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func Create(c *fiber.Ctx) error {
	var toDoObject domain.ToDo

	err := c.BodyParser(&toDoObject)
	if err != nil {
		stackTrace := errors.WithStack(err)
		return response.NewCustomError(response.ErrMsgParseError, response.ErrParsingBody, fiber.StatusBadRequest, nil, stackTrace, nil)
	}

	var errList []response.ValidationError

	err = domain.Validator.Struct(&toDoObject)
	if err != nil {
		fieldsErrors := validators.ValidateFields(err)
		errList = append(errList, *fieldsErrors...)
	}

	if toDoObject.Deadline != nil {
		//TODO validar data de fim
	}

	if len(errList) > 0 {
		return response.NewCustomError(response.ErrMsgBadRequest, response.ErrBadRequest, fiber.StatusBadRequest, &errList, nil, nil)
	}

	err = database.GetDB(c).Create(&toDoObject).Error
	if err != nil {
		stackTrace := errors.WithStack(err)
		return response.NewGormError(err, stackTrace, "company")
	}

	statusCode := fiber.StatusCreated
	return c.Status(statusCode).JSON(&toDoObject)
}
