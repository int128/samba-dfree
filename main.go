package main

import "os"
import "fmt"
import "io/ioutil"
import "path/filepath"
import "syscall"

func main() {
	exportdir := "."
	if len(os.Args) > 1 {
		exportdir = os.Args[1]
	}

	absexportdir, err := filepath.Abs(exportdir)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(absexportdir)
	if err != nil {
		panic(err)
	}

	total, available := uint64(0), uint64(0)
	for _, file := range files {
		path := filepath.Join(absexportdir, file.Name())
		var stat syscall.Statfs_t
		if err := syscall.Statfs(path, &stat); err == nil {
			total += stat.Blocks * uint64(stat.Bsize) / 1024
			available += stat.Bavail * uint64(stat.Bsize) / 1024
		}
	}

	fmt.Println(total, available)
}
