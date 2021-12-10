package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var dir = make(chan string)
var memSafe = &sync.Mutex{}

func main() {
	existingDrive := GetTheAllDrive()
	fmt.Println(existingDrive)
	fmt.Println("Press ctr+c to exit")
	fmt.Println(`Enter  the file type .pdf .exe .ppt et cetera..
	`)
	homePath, _ := os.UserHomeDir()
	subPath := GetAllSubDirectoryPath(homePath)

	var fileExt string
	fmt.Scanln(&fileExt)
	start := time.Now()
	folder := strings.Split(fileExt, ".")
	wg := &sync.WaitGroup{}
	os.Mkdir(folder[1], 0777)
	homeDrive := homePath
	fmt.Println(homeDrive)

	for p := 0; p < len(existingDrive); p++ {
		subPathDrive := GetAllSubDirectoryPath(existingDrive[p] + ":\\")
		for i := 0; i < len(subPathDrive)-1; i++ {

			wg.Add(2)
			go ThrowDirectoryName(existingDrive[p]+":\\"+subPathDrive[i], wg, fileExt)
			go ThrowDirectoryName(existingDrive[p]+":\\"+subPathDrive[i+1], wg, fileExt)
		}

	}
	for i := 0; i < len(subPath)-1; i++ {
		wg.Add(2)
		go ThrowDirectoryName(homeDrive+"\\"+subPath[i], wg, fileExt)
		go ThrowDirectoryName(homeDrive+"\\"+subPath[i+1], wg, fileExt)
	}

	go Copy(wg, folder[1])
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Copy completed! ", elapsed, "time took")

}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())

	}
}
