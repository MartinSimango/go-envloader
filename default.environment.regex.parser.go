package goenvloader

import (
	"fmt"
	"regexp"
	"strings"
)

type EnclosedType struct {
	LeftEnclosure  string
	RightEnclosure string
}

type EnvironmentRegexParserImpl struct {
	EnvironmentVariableRegExp string
	Seperator                 string
	Prefix                    string
	EncloseType               *EnclosedType
}

// Check if we implement interface
var _ EnvironmentRegexParser = &EnvironmentRegexParserImpl{}

func NewDefaultEnvironmentRegexParser() *EnvironmentRegexParserImpl {
	return &EnvironmentRegexParserImpl{
		EnvironmentVariableRegExp: "^\\${[a-zA-Z_][a-zA-Z_$0-9]*(,..*)?}$",
		Seperator:                 ",",
		Prefix:                    "$",
		EncloseType: &EnclosedType{
			LeftEnclosure:  "{",
			RightEnclosure: "}",
		},
	}
}

func NewCustomerEnvironmentRegexParser(prefix, seperator string, enclosedType *EnclosedType) *EnvironmentRegexParserImpl {
	var leftEnclosure string = ""
	var rightEnclosure string = ""
	if enclosedType != nil {
		leftEnclosure = enclosedType.LeftEnclosure
		rightEnclosure = enclosedType.RightEnclosure
	}
	regex := "^" + regexp.QuoteMeta(prefix) + regexp.QuoteMeta(leftEnclosure) + "[a-zA-Z_][a-zA-Z_$0-9]*(" + regexp.QuoteMeta(seperator) + "..*)?" + regexp.QuoteMeta(rightEnclosure) + "$"
	return &EnvironmentRegexParserImpl{
		EnvironmentVariableRegExp: regex,
		Seperator:                 seperator,
		Prefix:                    prefix,
		EncloseType:               enclosedType,
	}
}

func (rp *EnvironmentRegexParserImpl) GetEnv(field string) (*EnvironmentVariable, error) {
	leftEnclosure := ""
	rightEnclosure := ""
	if rp.EncloseType != nil {
		leftEnclosure = rp.EncloseType.LeftEnclosure
		rightEnclosure = rp.EncloseType.RightEnclosure
	}
	// if strings.TrimSpace(rp.Prefix) != "" {
	if strings.HasPrefix(strings.TrimSpace(field), rp.Prefix+leftEnclosure) {
		matched, err := regexp.MatchString(rp.EnvironmentVariableRegExp, strings.TrimSpace(field))
		if !matched || err != nil {
			if rp.Seperator == "" {
				return nil, fmt.Errorf("error parsing environment variable.\nEnviroment variable must follow format: %s%sENV%s", rp.Prefix, leftEnclosure, rightEnclosure)
			} else {
				return nil, fmt.Errorf("error parsing environment variable.\nEnviroment variable must follow format: %s%sENV%sdefault_value%s or %s%sENV%s", rp.Prefix,
					leftEnclosure, rp.Seperator, rightEnclosure, rp.Prefix, leftEnclosure, rightEnclosure)
			}
		}
		return rp.extractEnvValue(field)
	}
	// }
	return &EnvironmentVariable{
		Name:         "",
		DefaultValue: field,
	}, nil
}

func (rp *EnvironmentRegexParserImpl) extractEnvValue(field string) (*EnvironmentVariable, error) {

	enclosedString, err := rp.getEnclosedString(field)
	if err != nil {
		return nil, err
	}

	if rp.Seperator == "" {
		return &EnvironmentVariable{
			Name:         enclosedString,
			DefaultValue: "",
		}, nil
	}

	envArray := strings.Split(enclosedString, rp.Seperator)
	defaultValue := ""

	if len(envArray) == 2 {
		defaultValue = envArray[1]
	}

	return &EnvironmentVariable{
		Name:         envArray[0],
		DefaultValue: defaultValue,
	}, nil

}

func (rp *EnvironmentRegexParserImpl) getEnclosedString(field string) (string, error) {
	if rp.EncloseType == nil {
		if rp.Prefix == "" {
			return field, nil
		}
		prefixIndex := strings.Index(field, rp.Prefix)
		if prefixIndex == -1 {
			return "", fmt.Errorf("environment variable missing prefix " + rp.Prefix + "'")
		}
		return field[prefixIndex+1:], nil
	}
	if rp.Prefix != "" {
		field = field[1:]
	}
	openingBraceIndex := strings.Index(field, rp.EncloseType.LeftEnclosure)
	closingBraceIndex := strings.LastIndex(field, rp.EncloseType.RightEnclosure)
	if openingBraceIndex == -1 {
		return "", fmt.Errorf("environment variable missing '" + rp.EncloseType.LeftEnclosure + "'")
	}

	if closingBraceIndex == -1 {
		return "", fmt.Errorf("environment variable missing '" + rp.EncloseType.RightEnclosure + "'")
	}

	return field[openingBraceIndex+1 : closingBraceIndex], nil
}
