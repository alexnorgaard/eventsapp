package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func ValidateImage(c echo.Context) error {
	data, err := validate.FromRequest(c.Request())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	v := data.Create()
	v.AddRule("image", "required")
	v.AddRule("image", "isImage")
	if !v.Validate() {
		return errors.New("bad request - invalid image")
	}
	return nil
}
