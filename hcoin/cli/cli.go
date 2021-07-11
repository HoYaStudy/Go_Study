package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/HoYaStudy/Go_Study/hcoin/explorer"
	"github.com/HoYaStudy/Go_Study/hcoin/rest"
)

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		fmt.Println("Starting Rest API")
		rest.Start(*port)
	case "html":
		fmt.Println("Starting Explorer")
		explorer.Start(*port)
	default:
		usage()
	}
}

func usage() {
	fmt.Println()
	fmt.Println("### Welcome to hCoin ###")
	fmt.Println()
	fmt.Println("Please use the following flags")
	fmt.Println()
	fmt.Println("-port=4000:  Set the PORT of the server")
	fmt.Println("-mode=rest:  Choose between 'html' and 'rest'")
	runtime.Goexit()
}
