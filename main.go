package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var dir = make(chan string)
var memSafe = &sync.Mutex{}

func main() {
	fmt.Println(`Enter  the file type 
	.pdf
	.exe
	.ppt	
    et cetera..
	
	`)
	fmt.Println("Press ctr+c to exit")
	var fileExt string
	fmt.Scanln(&fileExt)
	fmt.Println(fileExt)
	folder := strings.Split(fileExt, ".")
	wg := &sync.WaitGroup{}

	dName, _ := os.Getwd()
	os.Mkdir(folder[1], 0777)
	homePath, _ := os.UserHomeDir()
	homeDrive := homePath
	fmt.Println(homeDrive, homePath)
	fmt.Println(dName, homeDrive)
	wg.Add(2)
	go ThrowDirectoryName(homeDrive, wg, fileExt)
	go ThrowDirectoryName(homeDrive, wg, fileExt)
	go Copy(wg, folder[1])
	wg.Wait()
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
