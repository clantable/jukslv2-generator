package usecase

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/clantable/jukslv2-generator/config"
	"github.com/clantable/jukslv2-generator/util"
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
		if Exists(c.UsecaseConfig.Path + "/" + f.Name()) {
			continue
		}
		if strings.Contains(f.Name(), "_test.go") {
			continue
		}
		if util.Contains(c.DomainConfig.ExceptFiles, f.Name()) {
			continue

		}
		Gen(c, strings.Split(f.Name(), ".")[0])
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
