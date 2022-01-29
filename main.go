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
var existingDrive []string
var homePath string

func init() {
	fmt.Println("Press ctr+c to Exit!")
	fmt.Println()
	existingDrive = GetTheAllDrive()
	homePath, _ = os.UserHomeDir()
}

func main() {
	fmt.Println("For Entire system copy Enter 1 or For custom path copy Enter 2")
	var choice int
	fmt.Scanln(&choice)
	if choice == 2 {
		fmt.Println("Enter the custom path")
		var customPath string
		fmt.Scanln(&customPath)
		homePath = customPath
		existingDrive = nil

	}
	var fileExt string
	file := File{}

	subPath := file.GetAllSubDirectoryPath(homePath)
	if len(subPath) == 0 {
		fmt.Println("No such Directory found")
		time.Sleep(time.Second * 2)
		return
	}

	fmt.Println(`Enter  the File type .pdf .exe .ppt et cetera..
	`)

	fmt.Scanln(&fileExt)
	start := time.Now()
	folder := strings.Split(fileExt, ".")[1]
	wg := &sync.WaitGroup{}
	os.Mkdir(folder, 0777)
	homeDrive := homePath

	for p := 0; p < len(existingDrive); p++ {
		subPathDrive := file.GetAllSubDirectoryPath(existingDrive[p] + ":\\")
		for i := 0; i < len(subPathDrive)-1; i++ {
			wg.Add(2)
			go file.ThrowDirectoryName(existingDrive[p]+":\\"+subPathDrive[i], wg, fileExt)
			go file.ThrowDirectoryName(existingDrive[p]+":\\"+subPathDrive[i+1], wg, fileExt)
		}

	}
	for i := 0; i < len(subPath)-1; i++ {

		wg.Add(2)
		go file.ThrowDirectoryName(homeDrive+"\\"+subPath[i], wg, fileExt)
		go file.ThrowDirectoryName(homeDrive+"\\"+subPath[i+1], wg, fileExt)
	}

	go Copy(wg, folder)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Copy completed! ", elapsed, "time took")

}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())

	}
}
