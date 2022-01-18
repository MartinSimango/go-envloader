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

func NewEnvironmentConfigFileParserImpl(regexParser EnvironmentRegexParser) *EnvironmentLoaderImpl {
	return &EnvironmentLoaderImpl{
		RegexParser: regexParser,
	}
}

func (ecfp *EnvironmentLoaderImpl) LoadIntEnv(field string) (int, error) {
	envValue, fieldError := ecfp.RegexParser.GetEnv(field)

	if fieldError != nil {
		return 0, fieldError
	}

	value := ecfp.getEnv(*envValue)

	if value == "" {
		return 0, fmt.Errorf("could not find environment variable '%s'", envValue.Name)

	}

	intFieldValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intFieldValue, nil
}

func (ecfp *EnvironmentLoaderImpl) LoadStringEnv(field string) (string, error) {
	envValue, fieldError := ecfp.RegexParser.GetEnv(field)

	if fieldError != nil {
		return "", fieldError
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
