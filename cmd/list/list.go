package list

import (
	"log"
	"path/filepath"
	"regexp"

	"github.com/yahao333/pm/config"
)

func Run() {
	pattern := regexp.MustCompile(filepath.Join(config.Conf.BaseDir, "/.pm/(.*)/metadata.json"))
	paths, err := filepath.Glob(filepath.Join(config.Conf.BaseDir, "/.pm/*/metadata.json"))
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range paths {
		log.Println(pattern.FindStringSubmatch(p)[1])
	}
}
