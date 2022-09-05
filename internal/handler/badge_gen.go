package handler

import (
	"io/ioutil"
	"os"

	"github.com/clantable/jukslv2-generator/config"
	"github.com/sirupsen/logrus"
)

func BadgeGen(c config.Config) {
	domainFiles, err := ioutil.ReadDir(c.DomainConfig.Path)
	if err != nil {
		logrus.Panic(err)
	}
	for _, f := range domainFiles {
		if f.IsDir() {
			continue
		}
		if Exists(c.HandlerConfig.Path + "/" + f.Name() + ".go") {
			continue
		}
		Gen(c, f.Name())
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
