package probe

import (
	"io/ioutil"
	"os"
)

var (
	fileName = "/tmp/probe"
	file     *os.File
)

func Create() error {
	_, err := os.Stat("/tmp")
	if os.IsNotExist(err) {
		err = os.Mkdir("/tmp", 0777)
	}
	if err != nil {
		return err
	}
	file, err = ioutil.TempFile("/tmp", "probe")
	if err != nil {
		return err
	}
	fileName = file.Name()
	return err
}

func Remove() error {
	return os.Remove(fileName)
}

func Exists() bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}
