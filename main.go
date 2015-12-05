package main

import "os"
import "fmt"
import "io/ioutil"
import "path/filepath"
import "syscall"

func queryBlocks(path string) (uint64, uint64) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, 0
	} else {
		total := stat.Blocks * uint64(stat.Bsize) / 1024
		available := stat.Bavail * uint64(stat.Bsize) / 1024
		return total, available
	}
}

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

	totals, availables := queryBlocks(absexportdir)
	for _, file := range files {
		path := filepath.Join(absexportdir, file.Name())
		total, available := queryBlocks(path)
		totals += total
		availables += available
	}

	fmt.Println(totals, availables)
}
