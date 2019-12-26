package probe

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var (
	fileName = "probeFile"
)

func Create() error {
	err := ioutil.WriteFile(locateFile(),[]byte(time.Now().String()),0644)
	if err != nil {
		return err
	}
	return err
}

func Remove() error {
	return os.Remove(locateFile())
}

func Exists() bool {
	if _, err := os.Stat(locateFile()); err == nil {
		return true
	}
	return false
}

func locateFile() string {
	dir,_ :=  filepath.Abs(filepath.Dir(os.Args[0]))
	return dir+"/"+fileName
}