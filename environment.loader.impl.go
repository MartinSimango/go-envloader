package goenvloader

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvironmentLoaderImpl struct {
	RegexParser EnvironmentRegexParser
}

//Check we implement interface
var _ EnvironmentLoader = &EnvironmentLoaderImpl{}

func NewEnvironmentLoader(regexParser EnvironmentRegexParser) *EnvironmentLoaderImpl {
	return &EnvironmentLoaderImpl{
		RegexParser: regexParser,
	}
}

func NewDefaultEnvironmentLoader() *EnvironmentLoaderImpl {
	regexParser := NewCustomerEnvironmentRegexParser("", "", nil)
	return NewEnvironmentLoader(regexParser)
}

func (ecfp *EnvironmentLoaderImpl) LoadIntFromEnv(field string) (int, error) {
	envValue, err := ecfp.RegexParser.GetEnv(field)

	if err != nil {
		return 0, err
	}

	value := ecfp.getEnv(*envValue)

	if value == "" {
		return 0, fmt.Errorf("error: could not find environment variable '%s'", envValue.Name)

	}

	intFieldValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("error: failed to convert string to int: %w", err)
	}
	return intFieldValue, nil
}

func (ecfp *EnvironmentLoaderImpl) LoadStringFromEnv(field string) (string, error) {
	envValue, err := ecfp.RegexParser.GetEnv(field)

	if err != nil {
		return "", err
	}

	value := ecfp.getEnv(*envValue)

	if value == "" {
		return "", fmt.Errorf("could not find environment variable '%s'", envValue.Name)
	}
	return value, nil
}

func (ecfp *EnvironmentLoaderImpl) getEnv(envVariable EnvironmentVariable) string {
	env := strings.TrimSpace(os.Getenv(envVariable.Name))
	defaultValue := strings.TrimSpace(envVariable.DefaultValue)

	if env == "" {
		return defaultValue
	}
	return env
}
