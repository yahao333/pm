package verify

import (
	"log"
	"os"
	"strings"

	"github.com/yahao333/pm/cmd"
	"github.com/yahao333/pm/openpgp"
)

func Run() {
	packageName := os.Args[2]
	if !strings.HasSuffix(packageName, "tar.gz") {
		parts := strings.Split(packageName, "-")
		packageName = cmd.FormatPackageName(parts[0], parts[1], "", "")
	}
	if !strings.HasSuffix(packageName, ".tar.gz") {
		log.Fatal("invalid package name: ", packageName)
	}
	if !openpgp.Verify(packageName) {
		log.Fatal("not valid")
	}
	log.Println("ok")
}
