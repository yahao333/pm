package info

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yahao333/pm/metadata"

	"github.com/yahao333/pm/config"
)

func Run() {
	packageName := os.Args[2]
	f, err := os.Open(fmt.Sprintf(filepath.Join(config.Conf.BaseDir, "/.pm/%s/metadata.json"), packageName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	m, err := metadata.New(f)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
