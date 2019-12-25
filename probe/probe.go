package probe

import (
	"io/ioutil"
	"os"
	"time"
)

var (
	fileName = "/tmp/probe"
)

func Create() error {
	_, err := os.Stat("/tmp")
	if os.IsNotExist(err) {
		err = os.Mkdir("/tmp", 0777)
	}
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName,[]byte(time.Now().String()),0777)
	if err != nil {
		return err
	}
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
