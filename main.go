package main

import (
	"context"
	"fmt"
	"go-filecoin-api-client/api"
)

func main() {
	AddressAPI := api.NewAPI("").Address()
	addrs, err := AddressAPI.Ls(context.Background())
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("addrs:", addrs)
	return
}
