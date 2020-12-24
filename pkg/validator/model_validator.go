package validator

import "github.com/go-playground/validator/v10"

var _theValidator *validator.Validate

// GetValidator Flyweight for getting a structure validator.
func GetValidator() *validator.Validate {
	if _theValidator == nil {
		_theValidator = validator.New()
	}
	return _theValidator
}
