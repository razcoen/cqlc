package golang

import "github.com/go-playground/validator/v10"

type Options struct {
	Package string `yaml:"package" validate:"required"`
	Out     string `yaml:"out" validate:"required"`
}

func (o *Options) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
