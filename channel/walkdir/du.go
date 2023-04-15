package walkdir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//并发目录遍历

// 递归遍历目录，向通道发送文件的大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			s := filepath.Join(dir, entry.Name())
			walkDir(s, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// return dir inner info
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
