package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var Conf *Config

type Config struct {
	BaseDir      string
	KeyPrintPath string
}

func InitConfig(home string) {
	if home == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		home = usr.HomeDir
	}

	Conf = &Config{
		BaseDir:      home,
		KeyPrintPath: filepath.Join(home, ".keyprint"),
	}
	CheckDirectory(home)
}

func CheckDirectory(homeDir string) {

	path := filepath.Join(homeDir, "/.pm")
	_ = os.Mkdir(path, 0755)
	_ = os.Mkdir(filepath.Join(homeDir, "/.pm/cache"), 0755)
	/*
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, mode)
		}
	*/
}
