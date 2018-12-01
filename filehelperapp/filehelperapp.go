package main

import (
	pathfileops "../pathfileops"
	"errors"
	"fmt"
)

func main() {

	fh := pathfileops.FileHelper{}

	targetFile, err := fh.MakeAbsolutePath(fh.AdjustPathSlash("..\\logTest\\topTest1.txt"))

	if err != nil {
		panic(errors.New("filehelperapp-main() Error on MakeAbsolutePath()- " + err.Error()))
	}

	fInfo, err := fh.GetFileInfoFromPath(targetFile)

	if err != nil {
		panic(errors.New(fmt.Sprintf("filehelperapp-main() Error on GetFileInfoFromPath(%v)- ", targetFile) + err.Error()))
	}

	fmt.Println("Target File: ", targetFile)
	fmt.Println("FileInfo.IsDir():", fInfo.IsDir())
	dt := pathfileops.DateTimeUtility{}

	tStr := dt.GetDateTimeNanoSecText(fInfo.ModTime())

	fmt.Println("FileInfo.ModTime():", tStr)

	fmt.Println("FileInfo.Name():", fInfo.Name())

	fmt.Println("FileInfo.Mode()", fInfo.Mode())

	fmt.Println("FileInfo.Size():", fInfo.Size())

	fmt.Println("FileInfo.Sys():", fInfo.Sys())

}
