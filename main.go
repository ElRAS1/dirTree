package main

import (
	"fmt"
	"io"
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
		levelDir: 0,
		prfile:   printFiles,
	}

	dirArr := make([]string, 0)
	dirFile := make([]string, 0)
	tree(path, &config, &dirArr, &dirFile)
	return nil
}

func tree(path string, config *Config, dirArr *[]string, dirFile *[]string) {

	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, i := range entries {
		if i.IsDir() {
			*dirArr = append(*dirArr, i.Name())
			printDir(dirArr, config)
			config.levelDir++
			*dirArr = make([]string, 0)
			tree(path+"/"+i.Name(), config, dirArr, dirFile)

		} else if config.prfile {
			st, err := os.Stat(path + "/" + i.Name())
			if err != nil {
				return
			}
			size := st.Size()
			if size > 0 {
				*dirFile = append(*dirFile, i.Name()+fmt.Sprintf(" (%db)", size))
			} else {
				*dirFile = append(*dirFile, i.Name()+" (empty)")
			}
			printFile(dirFile, config)
			*dirFile = make([]string, 0)
		}

	}

	config.levelDir--

}

// сделать одну функцию которая будет принимать разделитель
func printDir(dirArr *[]string, config *Config) {

	for indx, i := range *dirArr {
		if indx == len(*dirArr)-1 && config.levelDir > 0 {
			fmt.Printf("|%s└───%s\n", strings.Repeat("    ", config.levelDir), i)
	
		} else {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir), config.tabDir, i)
		}
	}
}

func printFile(dirFile *[]string, config *Config) {

	for indx, i := range *dirFile {
		if indx == len(*dirFile)-1 {
			fmt.Printf("|%s└───%s\n", strings.Repeat("    ", config.levelDir), i)
		} else {
			fmt.Printf("%s%s%s\n", strings.Repeat("    ", config.levelDir), config.tabDir, i)
		}
	}

}
