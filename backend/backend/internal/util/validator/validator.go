package validator

import "github.com/asaskevich/govalidator"

type Validator interface {
	ValidateStruct(interface{}) error
}

type GoValidator struct{}

func NewGoValidator() *GoValidator {
	return &GoValidator{}
}

func (g *GoValidator) ValidateStruct(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}
