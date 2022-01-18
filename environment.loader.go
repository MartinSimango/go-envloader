package goenvloader

type EnvironmentVariable struct {
	Name         string
	DefaultValue string
}

type EnvironmentLoader interface {
	LoadStringFromEnv(field string) (string, error)
	LoadIntFromEnv(field string) (int, error)
}
