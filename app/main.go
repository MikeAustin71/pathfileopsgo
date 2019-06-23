package main

import (
  pf "../pathfileops"
  "fmt"
  "strings"
)

/*


import (
  pf "../pathfileops"
  "fmt"
)


*/

func main() {

  mainTest76OpenThisFileWriteOnlyAppend()

}

func maintTest77OpenThisFileWriteOnlyAppend() {

  fh := pf.FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  // testText2 := "Damn the torpedoes, full speed ahead!\n"

  filePath := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\TestFileMgr_OpenThisFileWriteOnlyAppend_01.txt")

  basePath := fh.AdjustPathSlash("D:\\T04\\checkfiles")

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    fmt.Printf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateDirAndFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)
  bytesWritten := 0
  // fMgr.isFilePtrOpen = false
  bytesWritten, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fmt.Println("bytesWritten: ", bytesWritten)

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED!\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fmt.Println("      maintTest77OpenThisFileWriteOnlyAppend            ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}

func mainTest76OpenThisFileWriteOnlyAppend() {

  fh := pf.FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  testText2 := "Damn the torpedoes, full speed ahead!\n"

  filePath := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\TestFileMgr_OpenThisFileWriteOnlyAppend_01.txt")

  basePath := fh.AdjustPathSlash("D:\\T04\\checkfiles")

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    fmt.Printf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateDirAndFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnlyAppend().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite = []byte(testText2)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead1, err := fMgr.ReadFileLine('\n')

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.ReadFileLine(newline).\n"+
      "Error='%v'\n\n", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  bytesRead2, err := fMgr.ReadFileLine('\n')

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.ReadFileLine(newline).\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED!\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  stringRead := string(bytesRead1)

  stringRead = stringRead[:len(stringRead)-1]

  stringRead1 := stringRead

  testText1 = testText1[:len(testText1)-1]

  setSuccess := true

  if testText1 != stringRead {
    fmt.Printf("Error: Expected #1 stringRead='%v'.\n"+
      "Instead, #1 stringRead='%v'\n",
      testText1, stringRead)
    setSuccess = false
  }

  stringRead = string(bytesRead2)

  stringRead = strings.Replace(stringRead, "\r\n", "", -1)

  testText2 = strings.Replace(testText2, "\r\n", "", -1)

  if testText2 != stringRead {
    fmt.Printf("Error: Expected #2 stringRead='%v'.\n"+
      "Instead, #2 stringRead='%v'\n",
      testText2, stringRead)
    setSuccess = false
  }

  if !setSuccess {
    return
  }

  fmt.Println("       mainTest76OpenThisFileWriteOnlyAppend            ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println("stringRead1: ", stringRead1)
  fmt.Println("  testText1: ", testText1)

  return
}

/*

func maintTest75FileMgrGetTimeVal() {

  filePath :=
    "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_2_test.txt"

  fMgr, err := pf.FileMgr{}.New(filePath)

  if err != nil {
    fmt.Printf("Error returned  by pf.FileMgr{}.New(filePath)\n"+
      "filePath='%v'\n"+
      "Error='%v'\n", filePath, err.Error())
    return
  }

  fileModTime, err := fMgr.GetFileModTime()

  if err != nil {
    fmt.Printf("Error returned by fMgr.GetFileModTime()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  fmt.Println("          maintTest75FileMgrGetTimeVal                 ")
  fmt.Println("********************************************************")
  fmt.Println("    fileModTime: ", fileModTime.Format(timeFormatSpec))

}

func mainTest73FileHelperFileExist() {

  filePath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\checkfiles"

  dMgr, err := pf.DirMgr{}.New(filePath)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(filePath)\n"+
      "Error='%v'", err.Error())
    return
  }

  dirDoesExist, err := dMgr.DoesThisDirectoryExist()

  fmt.Println("          mainTest73FileHelperFileExist                 ")
  fmt.Println("********************************************************")
  fmt.Println("    filePath: ", filePath)
  fmt.Println("dirDoesExist: ", dirDoesExist)

}

func mainTest72OpenReadOnlyFile() {
  fh := pf.FileHelper{}

  rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\checkfiles\\TestFileMgr_OpenThisFileReadOnly_03.txt"
  filePath, err := fh.MakeAbsolutePath(rawPath)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath("+
      "rawPath)\n"+
      "rawPath='%v'\n"+
      "Error='%v'\n", rawPath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    fmt.Printf("Error returned from fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    fmt.Printf("Non-Path Error returned from #1 fMgr.DoesThisFileExist().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  if fileDoesExist {
    fmt.Printf("ERROR: Test file should NOT exist!.\n"+
      "However, test file DOES EXIST!\n"+
      "test file='%v'", filePath)
    _ = fh.DeleteDirFile(filePath)
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.CreateThisFile().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    fmt.Printf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())

    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.CloseThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.DeleteThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }

}

func mainTest71IsPathFileString() {
  fh := pf.FileHelper{}

  testPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  pathFileType, absolutePath, err := fh.IsPathFileString(testPath)

  if err != nil {
    fmt.Printf("Error returned from fh.IsPathFileString(testPath). "+
      "pathFile='%v' Error='%v' ", testPath, err.Error())
    return
  }

  fmt.Println("              mainTest70AdjustPathStr                   ")
  fmt.Println("********************************************************")
  fmt.Println("    testPath: ", testPath)
  fmt.Println("absolutePath: ", absolutePath)
  fmt.Println("pathFileType: ", pathFileType.String())

}

func mainTest70AdjustPathStr() {

  fh := pf.FileHelper{}

  testPath := "../../../"

  adjustedPath := fh.AdjustPathSlash(testPath)

  fmt.Println("              mainTest70AdjustPathStr                   ")
  fmt.Println("********************************************************")
  fmt.Println("    testPath: ", testPath)
  fmt.Println("adjustedPath: ", adjustedPath)
}

func mainTest69CleanDirStr() {

  fh := pf.FileHelper{}

  //   testPathFile := "/d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops/" +
  //     "levelfilesfortest/level_0_0_test.txt"

  //    testPathFile := "d:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\pathfileops" +
  //      "\\levelfilesfortest\\level_0_0_test.txt"


  testPathFile := "../filesfortest//levelfilesfortest/level_01_dir/level_1_1_test.txt"

  absFilePath, err := fh.MakeAbsolutePath(testPathFile)

  if err != nil {

  }

  fmt.Println("              mainTest67AreFilesSame                   ")
  fmt.Println("********************************************************")

  volName := fp.VolumeName(testPathFile)

  cleanFilePath, isEmpty, err := fh.CleanDirStr(testPathFile)

  if err != nil {
    fmt.Printf("Error returned by fh.CleanDirStr(testPathFile)\n"+
      "testPathFile='%v'\nError='%v'\n",
      testPathFile, err.Error())
    return
  }

  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println("testPathFile: ", testPathFile)
  fmt.Println("--------------------------------------------------------")
  fmt.Println("fh.CleanDirStr() Results:")
  fmt.Println("--------------------------------------------------------")
  fmt.Println("      isEmpty: ", isEmpty)
  fmt.Println("    cleanPath: ", cleanFilePath)
  fmt.Println("  Volume Name: ", volName)
  fmt.Println("Absolute Path: ", absFilePath)
  fmt.Println()
}

*/
