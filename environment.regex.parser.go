package goenvloader

type EnvironmentRegexParser interface {
	GetEnv(field string) (*EnvironmentVariable, error)
}
