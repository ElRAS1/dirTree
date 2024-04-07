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
			printTree(&dirArr, config, i, indx, 0)
			tree(path+"/"+dirArr[indx].Name(), config)
		} else {
			printTree(&dirArr, config, i, indx, 1)
		}
	}
	config.levelDir--
}

func printTree(dirArr *[]fs.DirEntry, config *Config, i fs.DirEntry, indx int, shift int) {
	sep := "├───"

	if indx == len(*dirArr)-1 {
		sep = "└───"
	}
	if config.levelDir+shift > 0 {
		fmt.Printf("|")
	}
	fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir+shift), sep, i.Name())

}
