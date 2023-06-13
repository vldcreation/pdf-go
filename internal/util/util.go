package util

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

func LoadConfig(path string, c interface{}) error {
	var o sync.Once
	o.Do(func() {
		viper.AddConfigPath(path)

		viper.AutomaticEnv()
	})

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

func ParseFile(path string) (filename string, ext string, err error) {
	splitter := "/"
	lastString := ""
	lasStringSplit := []string{}

	strs := strings.Split(path, splitter)

	if len(strs) < 1 {
		lastString = strs[0]
	} else {
		lastString = strs[len(strs)-1]
	}

	lasStringSplit = strings.Split(lastString, ".")

	if len(lasStringSplit) < 2 {
		err = ErrFileInputNotValid
		return
	}

	filename, ext = lasStringSplit[0], lasStringSplit[1]
	return
}
