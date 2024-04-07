package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	// "path/filepath"
	// "strings"
)

type Config struct {
	tabFile  string
	tabDir   string
	levelDir int
	// levelFile int
	prfile bool
}

func main() {
	out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	// err := dirTree(out, path, printFiles)

	err := dirTree(out, "testdata", true)

	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	config := Config{
		tabDir:   "├───",
		tabFile:  "└───",
		levelDir: -1,
		prfile:   printFiles,
	}

	// dirArr := make([]fs.DirEntry, 0)
	tree(path, &config)

	return nil
}

func tree(path string, config *Config) {

	entries, err := os.ReadDir(path)
	dirArr := make([]fs.DirEntry, 0)

	if err != nil {
		return
	}
	for _, i := range entries {
		if !i.IsDir() && !config.prfile {
			continue
		}
		dirArr = append(dirArr, i)
	}
	for indx, i := range dirArr {
		if i.IsDir() {
			config.levelDir++
			printTree(&dirArr, config, i, indx)
			tree(path+"/"+dirArr[indx].Name(), config)
		} else {
			printTree(&dirArr, config, i, indx)
		}
	}
	config.levelDir--
}

func printTree(dirArr *[]fs.DirEntry, config *Config, i fs.DirEntry, indx int) {
	if i.IsDir() {
		if indx < len(*dirArr)-1 {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir), config.tabDir, i.Name())

		} else {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir), config.tabFile, i.Name())
		}

	} else {
		if indx < len(*dirArr)-1 {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir+1), config.tabDir, i.Name())

		} else {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir+1), config.tabFile, i.Name())
		}
	}
}
