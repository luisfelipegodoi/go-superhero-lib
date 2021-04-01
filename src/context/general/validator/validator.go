package validator

import "github.com/go-playground/validator/v10"

// validatorImpl modelo para a validação do bind dos requests
type validatorImpl struct {
	v *validator.Validate
}

// Validate metodo que implementa a interface do validator para execução da validação
func (cv *validatorImpl) Validate(i interface{}) error {
	return cv.v.Struct(i)
}

// Validator interface do componente de validação
type Validator interface {
	Validate(i interface{}) error
}

// New cria uma nova implementação da interface Validator
func New() Validator {
	return &validatorImpl{
		v: validator.New(),
	}
}
