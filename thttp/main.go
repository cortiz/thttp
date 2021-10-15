package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Flag []string
}

func main() {
	ctx := kong.Parse(&CLI)
	fmt.Println(ctx.Command())
	for i := range CLI.Flag {
		fmt.Printf("%d %s \n", i, CLI.Flag[i])
	}
	// request, err := request.NewRequest("../test-requests/simpleRequest.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%s %s\n", request.Method, request.URI)
	// for key, value := range request.Headers {
	// 	fmt.Printf("%s:%s\n", key, value)
	// }

}
