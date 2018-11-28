package main

import (
	"MikeAustin71/pathfilego/003_filehelper/common"
	"errors"
	"fmt"
)

func main()  {

	fh := common.FileHelper{}

	targetFile, err := fh.MakeAbsolutePath(fh.AdjustPathSlash( "..\\logTest\\topTest1.txt"))

	if err != nil {
		panic (errors.New("filehelperapp-main() Error on MakeAbsolutePath()- "+ err.Error()) )
	}

	fInfo, err := fh.GetFileInfoFromPath(targetFile)

	if err != nil {
		panic (errors.New(fmt.Sprintf("filehelperapp-main() Error on GetFileInfoFromPath(%v)- ",targetFile)+ err.Error()) )
	}

	fmt.Println("Target File: ", targetFile)
	fmt.Println("FileInfo.IsDir():", fInfo.IsDir())
	dt := common.DateTimeUtility{}

	tStr := dt.GetDateTimeNanoSecText(fInfo.ModTime())

	fmt.Println("FileInfo.ModTime():", tStr)

	fmt.Println("FileInfo.Name():", fInfo.Name())

	fmt.Println("FileInfo.Mode()", fInfo.Mode())

	fmt.Println("FileInfo.Size():" ,fInfo.Size())

	fmt.Println("FileInfo.Sys():", fInfo.Sys())

}
