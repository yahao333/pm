package get

import (
	"log"
	"os"

	"github.com/yahao333/pm/cmd"
	"github.com/yahao333/pm/openpgp"
)

func Run() {
	packagePath, err := cmd.DownloadPackage(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if !openpgp.Verify(packagePath) {
		log.Fatal("not valid")
	}
	log.Println("ok")
}
