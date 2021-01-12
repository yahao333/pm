package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yahao333/pm/cmd/build"
	"github.com/yahao333/pm/cmd/get"
	"github.com/yahao333/pm/cmd/info"
	"github.com/yahao333/pm/cmd/install"
	"github.com/yahao333/pm/cmd/list"
	"github.com/yahao333/pm/cmd/remove"
	"github.com/yahao333/pm/cmd/verify"
)

func main() {
	log.SetFlags(0)
	switch os.Args[1] {
	case "build":
		build.Run()
	case "verify":
		verify.Run()
	case "install":
		install.Run()
	case "get":
		get.Run()
	case "remove":
		remove.Run()
	case "list":
		list.Run()
	case "info":
		info.Run()
	default:
		fmt.Println("not a command")
	}
}
