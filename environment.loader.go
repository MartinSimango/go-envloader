package goenvloader

type EnvironmentVariable struct {
	Name         string
	DefaultValue string
}

type EnvironmentLoader interface {
	LoadStringEnv(field string) (string, error)
	LoadIntEnv(field string) (int, error)
}
