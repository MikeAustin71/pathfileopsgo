package main

import (
  fp "path/filepath"
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

  mainTest71IsPathFileString()

}

func mainTest71IsPathFileString() {
  fh := pf.FileHelper{}

  testPath :=  fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
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
/*
  testPathFile := "/d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops/" +
    "levelfilesfortest/level_0_0_test.txt"

   testPathFile := "d:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\pathfileops" +
     "\\levelfilesfortest\\level_0_0_test.txt"

*/

  testPathFile := "../filesfortest//levelfilesfortest/level_01_dir/level_1_1_test.txt"

  absFilePath, err := fh.MakeAbsolutePath(testPathFile)

  if err != nil {

  }

  fmt.Println("              mainTest67AreFilesSame                   ")
  fmt.Println("********************************************************")

  volName :=  fp.VolumeName(testPathFile)

  cleanFilePath, isEmpty, err := fh.CleanDirStr(testPathFile)

  if err != nil {
    fmt.Printf("Error returned by fh.CleanDirStr(testPathFile)\n" +
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
