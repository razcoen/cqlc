package golang

type Options struct {
	Package  string          `yaml:"package" validate:"required"`
	Out      string          `yaml:"out" validate:"required"`
	Defaults DefaultsOptions `yaml:"defaults"`
}

type DefaultsOptions struct {
	BatchType string `yaml:"batch_type"`
}
