# go-envloader
A simple go library for loading environment variables with default values.
### Install
```
go get github.com/MartinSimango/go-envloader
```


### Example

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

	val, err := envLoader.LoadIntFromEnv("&<ENV_VARIABLE:9000>")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(val)
	}
}

```
The regex parser defines the format of how the environment variable and it's default will be. For the above example the format of environment variable is: \
&{ENVIRONMENT_VARIABLE:default_value}
