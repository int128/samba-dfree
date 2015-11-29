package main

import "os"
import "fmt"
import "log"
import "io/ioutil"
import "path/filepath"
import "syscall"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(wd)
	if err != nil {
		panic(err)
	}

	total, available := uint64(0), uint64(0)
	for _, file := range files {
		path := filepath.Join(wd, file.Name())
		var stat syscall.Statfs_t
		if err := syscall.Statfs(path, &stat); err != nil {
			log.Printf("Warning: ignored path %s (%s)\n", path, err)
		} else {
			total += stat.Blocks * uint64(stat.Bsize) / 1024
			available += stat.Bavail * uint64(stat.Bsize) / 1024
		}
	}

	fmt.Println(total, available)
}
