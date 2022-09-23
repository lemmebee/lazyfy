package log

import (
	"fmt"
	"os"
	"path"
)

const (
	logLocation = "./.log"
)

func Log(s ...interface{}) error {
	f, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, v := range s {
		_, err := f.WriteString(fmt.Sprintf("%q\n", v))
		if err != nil {
			return err
		}
	}

	return nil
}

func createLogFile(l string) error {
	err := os.MkdirAll(path.Dir(l), os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Stat(l)
	if err != nil && os.IsNotExist(err) {
		_, err = os.Create(l)
		return err
	}

	return err
}
