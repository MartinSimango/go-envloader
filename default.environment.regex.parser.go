package goenvloader

import (
	"fmt"
	"regexp"
	"strings"
)

type DefaultEnvironmentRegexParser struct {
	EnvironmentVariableRegExp string
}

// Check if we implement interface
var _ EnvironmentRegexParser = &DefaultEnvironmentRegexParser{}

func NewDefaultEnvironmentRegexParser() *DefaultEnvironmentRegexParser {
	return &DefaultEnvironmentRegexParser{
		EnvironmentVariableRegExp: "^\\${[a-zA-Z_][a-zA-Z_$0-9]*(,..*)?}$",
	}
}

func (rp *DefaultEnvironmentRegexParser) GetEnv(field string) (*EnvironmentVariable, error) {
	if strings.HasPrefix(strings.TrimSpace(field), "${") {
		matched, err := regexp.MatchString(rp.EnvironmentVariableRegExp, strings.TrimSpace(field))
		if !matched || err != nil {
			return nil, fmt.Errorf("error parsing environment variable.\nEnviroment variable must follow format: ${ENV, default_value} or ${ENV}")
		}
		return rp.extractEnvValue(field)
	}
	return &EnvironmentVariable{
		Name:         "",
		DefaultValue: field,
	}, nil
}

func (rp *DefaultEnvironmentRegexParser) extractEnvValue(field string) (*EnvironmentVariable, error) {
	openingBraceIndex := strings.Index(field, "{")
	closingBraceIndex := strings.Index(field, "}")
	if openingBraceIndex == -1 {
		return nil, fmt.Errorf("environment variable missing { brace")
	}

	if closingBraceIndex == -1 {
		return nil, fmt.Errorf("environment variable missing { brace")
	}
	enclosedString := field[openingBraceIndex+1 : closingBraceIndex]

	envArray := strings.Split(enclosedString, ",")
	defaultValue := ""
	if len(envArray) == 2 {
		defaultValue = envArray[1]
	}

	return &EnvironmentVariable{
		Name:         envArray[0],
		DefaultValue: defaultValue,
	}, nil

}
