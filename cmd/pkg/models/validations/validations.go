package validationspackage

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type apiError struct {
	Field string
	Msg   string
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}

// Allows to perform error validation for required fields
func ValidationError(err error, ve validator.ValidationErrors) []apiError {

	if errors.As(err, &ve) {
		out := make([]apiError, len(ve))

		for i, fe := range ve {
			out[i] = apiError{Field: fe.Field(), Msg: msgForTag(fe.Tag())}
		}
		return out
	}
	return nil
}
