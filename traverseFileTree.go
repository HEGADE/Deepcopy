package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sync"
)

func ThrowDirectoryName(homeDrive string, wg *sync.WaitGroup, fileExt string) {
	defer wg.Done()
	fileErr := filepath.Walk(homeDrive, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == fileExt {
			memSafe.Lock()
			dir <- path
			memSafe.Unlock()
		}
		return nil
	})
	if fileErr != nil {

		fmt.Println(fileErr)
	}

}
func GetAllSubDirectoryPath(homeDrive string) []string {
	subPath := []string{}
	filesBuf, _ := ioutil.ReadDir(homeDrive)
	for _, path := range filesBuf {
		subPath = append(subPath, path.Name())
	}

	return subPath

}
