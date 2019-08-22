package main

import (
	"./api"
	"fmt"
)

func main() {
	api := api.New("")
	res, _ := api.Request("mining", "status")
	fmt.Println(string(res))
}
