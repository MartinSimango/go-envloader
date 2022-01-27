package goenvloader

type EnvironmentVariable struct {
	Name         string
	DefaultValue string
}

type EnvironmentLoader interface {
	LoadStringFromEnv(field string) (string, error)
	LoadIntFromEnv(field string) (int, error)
	LoadFloatFromEnv(field string) (float64, error)
	LoadBoolFromEnv(field string) (bool, error)
}
