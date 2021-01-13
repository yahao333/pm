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
	"github.com/yahao333/pm/config"
)

const VERSION = "1.0.1"

func main() {
	home := os.Getenv("PM_HOME")
	config.InitConfig(home)

	log.SetFlags(0)
	if len(os.Args) > 1 {
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
	} else {
		fmt.Println("version", VERSION)
	}
}
