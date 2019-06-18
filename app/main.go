package main

import (
  pf "../pathfileops"
  "fmt"
)

/*


import (
  pf "../pathfileops"
  "fmt"
)


*/

func main() {

  maintTest75FileMgrGetTimeVal()

}

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

/*
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
