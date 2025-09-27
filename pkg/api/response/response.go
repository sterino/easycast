package response

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Base struct {
	Success   bool    `json:"success"`             // Success Успешно ли выполнен запрос
	Error     string  `json:"error,omitempty"`     // Error Сообщение об ошибке
	ErrorFICO *string `json:"errorFICO,omitempty"` // ErrorFICO Сообщение об ошибке из FICO
	ErrorCode int     `json:"errorCode"`           // ErrorCode Код ошибки
}

func ValidationError(errs error) Base {
	var errMsgs []string
	var valErrs validator.ValidationErrors

	if errors.As(errs, &valErrs) {
		for _, err := range valErrs {
			switch err.ActualTag() {
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("%s is not valid", err.Field()))
			}
		}
	}

	return Base{Success: false, Error: strings.Join(errMsgs, ", "), ErrorCode: 0}
}

func Error(msgs ...string) Base {
	return Base{Success: false, Error: strings.Join(msgs, ", "), ErrorCode: 0}
}
