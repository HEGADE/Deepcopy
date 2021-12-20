package main

import (
	"io"
	"os"
	"strings"
	"sync"
)

func Copy(wg *sync.WaitGroup, folder string) {
	defer wg.Done()
	for {
		select {
		case <-dir:

			copiedFileName := strings.Split(<-dir, "\\")
			srcFile, err := os.Open(<-dir)
			check(err)

			destFile, err := os.Create(folder + "/" + copiedFileName[len(copiedFileName)-1])
			check(err)

			_, err = io.Copy(destFile, srcFile)
			check(err)

			check(err)
			srcFile.Close()
			destFile.Close()
		}
	}

}
