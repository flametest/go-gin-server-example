package mysql

type Config struct {
	DSN     string `yaml:"dsn"`
	MaxOpen int    `yaml:"maxOpen"`
	MaxIdle int    `yaml:"maxIdle"`
	Debug   bool   `yaml:"debug"`
}
