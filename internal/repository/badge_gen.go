package repository

import (
	"io/ioutil"
	"os"
	"strings"

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
		if Exists(c.RepositoryConfig.Path + "/" + f.Name()) {
			continue
		}
		Gen(c, strings.Split(f.Name(), ".")[0])
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
