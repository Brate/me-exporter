package helpers

import (
	"errors"
	"fmt"
)

const (
	ErrorInternal     = "InternalError"
	ErrorBadRequest   = "BadRequest"
	ErrorCompareCrypt = "CompareCryptError"
	ErrorValidation   = "ValidationError"
)

func ErrorMessage(msg string, value interface{}) error {
	errMsg := fmt.Sprintf(msg, value)
	return errors.New(errMsg)
}
