package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var path, pattern string
	flag.StringVar(&path, "dir", "examples", "Name of the directory")
	flag.StringVar(&pattern, "pattern", "[0-9].txt", "Pattern name (regex like)")
	flag.Parse()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	total, _ := totalMatches(files, pattern)
	matches := 0
	for _, v := range files {
		if matchPattern(v.Name(), pattern) {
			matches++
			s := fmt.Sprintf("%v (%v of %v)", strings.Split(v.Name(), "_")[0], matches, total)
			os.Rename(
				path+"/"+v.Name(), path+"/"+s)
		}
	}
}

func matchPattern(file, pattern string) bool {
	match, _ := regexp.MatchString(pattern, file)
	return match
}

func totalMatches(list []os.FileInfo, pattern string) (int, int) {
	count, total := 0, 0
	total = len(list)
	for _, v := range list {
		if matchPattern(v.Name(), pattern) {
			count++
		}
	}
	return count, total
}
