package appExamples

import (
	"errors"
	"fmt"
	"time"
)

func ExampleExtractPathElements() {

	fh := FileHelper{}
	commonDir := fh.AdjustPathSlash("..\\..\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		panic(errors.New("ExampleExtractPathElements()- Error returned on fh.GetPathFileNameElements(), Error:" + err.Error()))
	}

	fmgr2 := fileMgr.CopyOut()

	if !fileMgr.Equal(&fmgr2) {
		panic(errors.New("ExampleExtractPathElements() - CopyToThis Equal Analysis Failed!"))
	}

	PrintFileManagerFields(fileMgr)
}

func PathElementsAnalysis(pathFile string) {
	fh := FileHelper{}
	commonDir := fh.AdjustPathSlash(pathFile)

	fMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		panic(errors.New("PathElementsAnalysis()- Error returned on fh.GetPathFileNameElements(), Error:" + err.Error()))
	}

	fMgr2 := fMgr.CopyOut()

	if !fMgr.Equal(&fMgr2) {
		panic(errors.New("PathElementsAnalysis() - CopyOut Equal Analysis Failed!"))
	}

	PrintFileManagerFields(fMgr)

}

func PrintFileManagerFields(fileMgr FileMgr) {

	fmt.Println("======================================")
	fmt.Println("            File Manager")
	fmt.Println("======================================")
	fmt.Println("IsInitialized:", fileMgr.IsInitialized)
	fmt.Println("OriginalPathFileName:", fileMgr.OriginalPathFileName)
	fmt.Println("AbsolutePathFileName:", fileMgr.AbsolutePathFileName)
	fmt.Println("AbsolutePathFileNameIsPopulated:", fileMgr.AbsolutePathFileNameIsPopulated)
	fmt.Println("AbsolutePathFileNameDoesExist:", fileMgr.AbsolutePathFileNameDoesExist)
	fmt.Println("FileName:", fileMgr.FileName)
	fmt.Println("FileNameIsPopulated:", fileMgr.FileNameIsPopulated)
	fmt.Println("FileExt:", fileMgr.FileExt)
	fmt.Println("FileExtIsPopulated:", fileMgr.FileExtIsPopulated)
	fmt.Println("FileNameExt:", fileMgr.FileNameExt)
	fmt.Println("FileNameExtIsPopulated:", fileMgr.FileNameExtIsPopulated)
	fmt.Println("IsFilePtrOpen: ", fileMgr.IsFilePtrOpen)
	PrintFileInfoPlusFields(fileMgr.ActualFileInfo)
	PrintDirMgrFields(fileMgr.DMgr)
}

func PrintFileInfoPlusFields(info FileInfoPlus) {
	fmt.Println("======================================")
	fmt.Println("            File Info Plus")
	fmt.Println("======================================")
	du := DateTimeUtility{}
	fmt.Println("  IsFInfoInitialized: ", info.IsFInfoInitialized)
	fmt.Println("IsDirPathInitialized: ", info.IsDirPathInitialized)
	fmt.Println("     CreateTimeStamp: ", du.GetDateTimeYMDAbbrvDowNano(info.CreateTimeStamp))
	fmt.Println("              Name(): ", info.Name())
	fmt.Println("              Size(): ", info.Size())
	fmt.Println("              Mode(): ", info.Mode())
	fmt.Println("           ModTime(): ", du.GetDateTimeYMDAbbrvDowNano(info.ModTime()))
	fmt.Println("             IsDir(): ", info.IsDir())
	fmt.Println("               Sys(): ", info.Sys())
	fmt.Println("           DirPath(): ", info.DirPath())
}

func CreateFileOnTopOfExistingFile() {
	tstFile := "..//logTest//testoverwrite//TestOverwrite001.txt"
	fMgr, err := FileMgr{}.New(tstFile)

	if err != nil {
		panic(fmt.Errorf("CreateFileOnTopOfExistingFile() - Error: FileMgr{}.New(tstFile) Failed. tstFile='%v' Error='%v'", tstFile, err.Error()))
	}

	if err != nil {
		panic(errors.New(fmt.Sprintf("CreateFileOnTopOfExistingFile() Error Creating File: '%v' Error: %v", tstFile, err.Error())))
	}

	defer fMgr.CloseFile()

	du := DateTimeUtility{}
	str := "Test Over Write Time Stamp: " + du.GetDateTimeEverything(time.Now())
	fMgr.WriteStrToFile(str)

}

func ExampleReadTestFile() {
	tstFile := "../testfiles/TestRead.txt"
	tstOutFile := "../testfiles/Output.txt"
	fh := FileHelper{}
	f, err := fh.OpenFileForReading(tstFile)

	if err != nil {
		fmt.Printf("Error Opening file: %v\n", tstFile)
	}

	defer f.Close()

	fOut, err2 := fh.CreateFile(tstOutFile)

	if err2 != nil {
		fmt.Printf("Error Opening file: %v\n", tstOutFile)
	}

	defer fOut.Close()

	buffer := make([]byte, 50000)
	doRead := true
	su := StringUtility{}
	strCnt := 0
	partialString := ""

	for doRead == true {
		n, err := fh.ReadFileBytes(f, buffer)

		nIdx := 0
		s := ""
		isPartialString := false
		extractStr := true

		for extractStr == true && n > 0 {
			s, nIdx, isPartialString = su.ReadStrNewLineFromBuffer(buffer, partialString, nIdx)

			if !isPartialString {
				strCnt++

				fh.WriteFileStr(fmt.Sprintf("%07d- %s\n", strCnt, s), fOut)

			} else {
				partialString = s
				fh.WriteFileStr(fmt.Sprintf("******* Partial String %07d- %s **********\n", strCnt, s), fOut)
			}

			if nIdx == -1 {
				extractStr = false
			}

		}

		if n < 50000 || err != nil {
			doRead = false
		}

	}

	fmt.Println("Completed File Read and output to output file: ", tstOutFile)
}
