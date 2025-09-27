package main

import (
	"fmt"

	apitest "github.com/hiansh1/go-console-e7.git/weither-api"
)

func main() {
	var name string
	fmt.Println("donner votre nom")
	fmt.Scan(&name)
	apitest.Say_Hello(name)
}
