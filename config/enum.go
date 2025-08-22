package config

type EnumValue struct {
	Name    string
	Display any
	Code    any
}

type Enum struct {
	Name   string
	Type   string
	Values []EnumValue
}

type EnumConfig struct {
	Enums []Enum `yaml:"enums"`
}
