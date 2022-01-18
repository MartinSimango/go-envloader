# go-envloader
A simple go library for loading environment variables from template strings

### Install

go get github.com/MartinSimango/go-envloader


### How To Use

``` go

package main

import (
	"fmt"

	goenvloader "github.com/MartinSimango/go-envloader"
)

func main() {
	regexParser := goenvloader.NewCustomerEnvironmentRegexParser("&", ":", &goenvloader.EnclosedType{
		LeftEnclosure:  "<",
		RightEnclosure: ">",
	})
	envLoader := goenvloader.NewEnvironmentLoader(regexParser)

	val, err := envLoader.LoadIntFromEnv("&{ENV_VARIABLE,9000}")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(val)
	}
}

```