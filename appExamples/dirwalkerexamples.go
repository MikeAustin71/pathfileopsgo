package appExamples

import (
	"fmt"
	"path/filepath"
)

func TestFilePathMatch002() {
	// filename := "D:\\Test03\\start.txt"
	// Note: This doesn't work. \\ generates
	// and invalid response from Match

	filename := "start.txt"

	pattern := "*.txt"

	matched, err := filepath.Match(pattern, filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matched)

}

func TestFilePathGlob003(dirPathAndPattern string) {

	matches, err := filepath.Glob(dirPathAndPattern)

	if err != nil {
		panic(err)
	}

	for i, f := range matches {
		fmt.Println(i, " - ", f)
	}

}
