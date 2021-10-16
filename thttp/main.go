package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

var CLI struct {
	File         []string `short:"f" help:"Files to execute." optional:"true" type:"existingfile"`
	Dir          []string `short:"d" help:"Executes all the files in the given directories" optional:"true" type:"existingdir"`
	Properties   string   `short:"p" help:"Properties file to be use can be yaml, json or property file" optional:"true"`
	Repeat       int      `short:"r" help:"Repeats the requests the given times" optional:"true"`
	HeaderOnly   bool     `short:"I" help:"Print only the headers" optional:"true"`
	WithHeader   bool     `short:"i" help:"Print the response with the headers" optional:"true"`
	Output       string   `short:"o" help:"Write the output to a file" optional:"true"`
	OutputFolder string   `short:"f" help:"Files to execute." optional:"true" type:"existingdir"`
	IgnoreSSL    bool     `short:"k" help:"Ignore SSL Validation" optional:"true"`
	Follow       bool     `short:"l" help:"Follow Redirects" optional:"true"`
	Plugin       []string `help:"Plugins to Load" optional:"true" type:"existingfile"`
}

func main() {
	ctx := kong.Parse(&CLI)
	if len(CLI.Dir) >= 0 && len(CLI.File) >= 0 {
		fmt.Println("Either File or Directory is require")
		err := ctx.PrintUsage(true)
		if err != nil {
			fmt.Println("Error")
		}
		os.Exit(1)
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
