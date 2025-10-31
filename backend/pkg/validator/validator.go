package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
  v *validator.Validate
}

func NewValidator() *Validator {
  return &Validator{v: validator.New()}
}

func (val *Validator) Validate(i interface{}) error {
  return val.v.Struct(i)
}
