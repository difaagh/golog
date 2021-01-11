package golog

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	FileName   string
	FolderName string
}

// get the log file save path
func getLogFilePath(c *Config) string {
	return fmt.Sprintf("%s", c.FolderName)
}

// get the name of saved log file
func getLogFileName(c *Config) string {
	return fmt.Sprintf("%s-%s.%s",
		c.FileName,
		time.Now().Format("20060102"),
		"log",
	)
}

// check if the file exists
func isExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// check if the file has permission
func checkPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// create a directory if it does not exist
func createIfNotExist(src string) error {
	if notExist := isExist(src); notExist == true {
		if err := mkdir(src); err != nil {
			return err
		}
	}

	return nil
}

// create a directory
func mkdir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// open for specific mode
func open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func openFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	fmt.Println(dir)
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := checkPermission(src)
	if perm == true {
		return nil, fmt.Errorf("Permission denied src: %s", src)
	}

	err = createIfNotExist(src)
	if err != nil {
		return nil, fmt.Errorf("file is not exist src: %s, err: %v", src, err)
	}

	f, err := open(src+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
