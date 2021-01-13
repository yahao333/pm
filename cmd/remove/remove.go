package remove

import (
	"github.com/yahao333/pm/config"
	"log"
	"os"
	"path/filepath"
)

func Run() {
	packageName := os.Args[2]
	err := os.RemoveAll(filepath.Join(config.Conf.BaseDir, "/.pm/", packageName))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("~/.pm/" + packageName)
	files, err := filepath.Glob(filepath.Join(config.Conf.BaseDir, "/.pm/pm/cache", packageName+"*"))
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range files {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}
}
