package main

import (
	"fmt"
	"fr/greytsu/sol_api_products/config"
)

func init() {
	config.LoadEnvironmentVariables()
}

func main() {
	fmt.Print("Hello, world!")

}
