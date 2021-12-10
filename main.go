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

	fmt.Println("Press ctr+c to exit")
	fmt.Println(`Enter  the file type .pdf .exe .ppt et cetera..
	`)
	start := time.Now()

	var fileExt string
	fmt.Scanln(&fileExt)
	folder := strings.Split(fileExt, ".")
	wg := &sync.WaitGroup{}
	os.Mkdir(folder[1], 0777)
	homePath, _ := os.UserHomeDir()
	homeDrive := homePath
	wg.Add(2)
	go ThrowDirectoryName(homeDrive, wg, fileExt)
	go Copy(wg, folder[1])
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Copy completed!", elapsed, "time took")

}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
