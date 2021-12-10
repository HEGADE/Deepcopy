package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var possibleDrive = []string{
	"A",
	"B",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}

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

// Getting all subdirecory
func GetAllSubDirectoryPath(homeDrive string) []string {
	subPath := []string{}
	filesBuf, _ := ioutil.ReadDir(homeDrive)
	for _, path := range filesBuf {
		subPath = append(subPath, path.Name())
	}

	return subPath

}

//Getting other drives
func GetTheAllDrive() []string {
	avilableDrive := []string{}
	for i := 0; i < len(possibleDrive); i++ {
		_, err := ioutil.ReadDir(possibleDrive[i] + "://")
		if err == nil {
			avilableDrive = append(avilableDrive, possibleDrive[i])

		}
	}
	return avilableDrive
}
