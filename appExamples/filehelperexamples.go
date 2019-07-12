package appExamples

import (
  appLib "../appLibs"
  pathFileOps "../pathfileops"
  "errors"
  "fmt"
  "time"
)

func ExampleExtractPathElements() {

  fh := pathFileOps.FileHelper{}
  commonDir := fh.AdjustPathSlash("..\\..\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  fileMgr, err := pathFileOps.FileMgr{}.NewFromPathFileNameExtStr(commonDir)

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
  fh := pathFileOps.FileHelper{}
  commonDir := fh.AdjustPathSlash(pathFile)

  fMgr, err := pathFileOps.FileMgr{}.NewFromPathFileNameExtStr(commonDir)

  if err != nil {
    panic(errors.New("PathElementsAnalysis()- Error returned on fh.GetPathFileNameElements(), Error:" +
      err.Error()))
  }

  fMgr2 := fMgr.CopyOut()

  if !fMgr.Equal(&fMgr2) {
    panic(errors.New("PathElementsAnalysis() - CopyOut Equal Analysis Failed!"))
  }

  PrintFileManagerFields(fMgr)

}

func PrintFileManagerFields(fileMgr pathFileOps.FileMgr) {
  ePrefix := "PrintFileManagerFields() "
  fmt.Println("======================================")
  fmt.Println("            File Manager")
  fmt.Println("======================================")
  fmt.Println("isInitialized:", fileMgr.IsInitialized())
  fmt.Println("originalPathFileName:", fileMgr.GetOriginalPathFileName())
  fmt.Println("absolutePathFileName:", fileMgr.GetAbsolutePathFileName())
  fmt.Println("isAbsolutePathFileNamePopulated:", fileMgr.IsAbsolutePathFileNamePopulated())

  doesFileExist, err := fileMgr.DoesThisFileExist()
  if err != nil {
    fmt.Printf(ePrefix+
      "Error from fileMgr.DoesThisFileExist(). Error='%v' \n", err.Error())
    return
  }

  fmt.Println("DoesThisFileExist():", doesFileExist)
  fmt.Println("fileName:", fileMgr.GetFileName())
  fmt.Println("isFileNamePopulated:", fileMgr.IsFileNamePopulated())
  fmt.Println("fileExt:", fileMgr.GetFileExt())
  fmt.Println("isFileExtPopulated:", fileMgr.IsFileExtPopulated())
  fmt.Println("fileNameExt:", fileMgr.GetFileNameExt())
  fmt.Println("isFileNameExtPopulated:", fileMgr.IsFileNameExtPopulated())
  fmt.Println("isFilePtrOpen: ", fileMgr.IsFilePointerOpen())

  fileInfoPlus, err := fileMgr.GetFileInfoPlus()

  if err != nil {
    fmt.Printf(ePrefix+
      "Error from fileMgr.GetFileInfoPlus(). Error='%v' \n", err.Error())
    return

  }

  PrintFileInfoPlusFields(fileInfoPlus)
  PrintDirMgrFields(fileMgr.GetDirMgr())
}

func PrintFileInfoPlusFields(info pathFileOps.FileInfoPlus) {
  fmt.Println("======================================")
  fmt.Println("            File Info Plus")
  fmt.Println("======================================")
  du := appLib.DateTimeUtility{}
  fmt.Println("  isFInfoInitialized: ", info.IsFileInfoInitialized())
  fmt.Println("isDirPathInitialized: ", info.IsDirectoryPathInitialized())
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
  ePrefix := "CreateFileOnTopOfExistingFile() "

  tstFile := "..//logTest//testoverwrite//TestOverwrite001.txt"
  fMgr, err := pathFileOps.FileMgr{}.NewFromPathFileNameExtStr(tstFile)

  if err != nil {
    _ = fMgr.CloseThisFile()
    panic(fmt.Errorf(ePrefix+
      "- Error: FileMgr{}.NewFromPathFileNameExtStr(tstFile) Failed. tstFile='%v' Error='%v'",
      tstFile, err.Error()))
    return
  }

  du := appLib.DateTimeUtility{}
  str := "Test Over Write Time Stamp: " + du.GetDateTimeEverything(time.Now())
  _, err = fMgr.WriteStrToFile(str)

  if err != nil {
    _ = fMgr.CloseThisFile()
    panic(fmt.Errorf(ePrefix+" %v ", err.Error()))
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    panic(fmt.Errorf(ePrefix+" %v ", err.Error()))
  }

}

func ExampleReadTestFile() {

  ePrefix := "ExampleReadTestFile() "
  tstFile := "../testfiles/TestRead.txt"
  tstOutFile := "../testfiles/Output.txt"
  fh := pathFileOps.FileHelper{}
  f, err := fh.OpenFileReadOnly(tstFile)

  if err != nil {
    fmt.Printf("Error Opening file: %v\n", tstFile)
    return
  }

  fOut, err2 := fh.CreateFile(tstOutFile)

  if err2 != nil {
    _ = f.Close()
    fmt.Printf(ePrefix+"Error Opening file: %v\n", tstOutFile)
    return
  }

  buffer := make([]byte, 50000)
  doRead := true
  su := appLib.StringUtility{}
  strCnt := 0
  partialString := ""

  for doRead == true {
    n, err := f.Read(buffer)

    nIdx := 0
    s := ""
    isPartialString := false
    extractStr := true

    for extractStr == true && n > 0 {
      s, nIdx, isPartialString = su.ReadStrNewLineFromBuffer(buffer, partialString, nIdx)

      if !isPartialString {
        strCnt++

        _, err = fOut.WriteString(fmt.Sprintf("%07d- %s\n", strCnt, s))

        if err != nil {
          _ = f.Close()
          _ = fOut.Close()
          fmt.Printf(ePrefix+"Error Writhing File Str #1: %v\n", err.Error())
          return
        }

      } else {
        partialString = s
        _, err = fOut.WriteString(fmt.Sprintf("******* Partial String %07d- %s **********\n", strCnt, s))

        if err != nil {
          _ = f.Close()
          _ = fOut.Close()
          fmt.Printf(ePrefix+"Error Writing File Str #2: %v\n", err.Error())
          return
        }

      }

      if nIdx == -1 {
        extractStr = false
      }

    }

    if n < 50000 {
      doRead = false
    }

  }

  _ = f.Close()
  _ = fOut.Close()

  fmt.Println("Completed File Read and output to output file: ", tstOutFile)
}
