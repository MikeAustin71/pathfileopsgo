package pathfileops

import (
  appLib "MikeAustin71/pathfileopsgo/appLibs"
  "errors"
  "fmt"
  "io"
  "testing"
  "time"
)

func TestFileHelper_MakeAbsolutePath_01(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.MakeAbsolutePath("")

  if err == nil {
    t.Error("Expected an error return from fh.MakeAbsolutePath(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeAbsolutePath_02(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.MakeAbsolutePath("   ")

  if err == nil {
    t.Error("Expected an error return from fh.MakeAbsolutePath(\"\") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeDirAll_01(t *testing.T) {
  fh := FileHelper{}

  err := fh.MakeDirAll("")

  if err == nil {
    t.Error("Expected an error return from fh.MakeDirAll(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeDirAll_02(t *testing.T) {
  fh := FileHelper{}

  err := fh.MakeDirAll("  ")

  if err == nil {
    t.Error("Expected an error return from fh.MakeDirAll(\"    \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_MoveFile_01(t *testing.T) {
  fh := FileHelper{}
  setupFile := fh.AdjustPathSlash("..//logTest//FileMgmnt//TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..//logTest//FileSrc//TestFile003.txt")
  destFile := fh.AdjustPathSlash("..//logTest//TestFile004.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Error(fmt.Sprintf("Error on DeleteDirFile() deleting destination file, '%v'. Error:", destFile), err)
    }

    if fh.DoesFileExist(destFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", destFile))
    }
  }

  err := fh.CopyFileByIo(setupFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to destination file '%v' does NOT Exist. Error='%v'", setupFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Error(fmt.Sprintf("Source File '%v' does NOT EXIST!!", srcFile))
  }

  _, err = fh.MoveFile(srcFile, destFile)

  if err != nil {
    t.Error(fmt.Sprintf("Error on FileHelper:MoveFile() moving src '%v' to destination '%v' ", srcFile, destFile), err)
  }

  if fh.DoesFileExist(srcFile) {
    t.Error(fmt.Sprintf("FileHelper:MoveFile() FAILED! Source File '%v' still exists!!", srcFile))
  }

  if !fh.DoesFileExist(destFile) {
    t.Error(fmt.Sprintf("FileHelper:MoveFile() FAILED! Destination File '%v' DOES NOT EXIST!", destFile))
  }
}

func TestFileHelper_OpenFile_01(t *testing.T) {
  fh := FileHelper{}
  target := fh.AdjustPathSlash(alogtopTest2Text)
  expected := "Top level test file # 2."
  f, err := fh.OpenFileForReading(target)

  if err != nil {
    t.Errorf("Failed to open file: '%v' , got error - '%v'", target, err.Error())
  }

  le := len(expected)
  bRead := make([]byte, le)
  _, err2 := io.ReadAtLeast(f, bRead, 10)

  if err2 != nil {
    t.Errorf("Error Reading Test File: %v. Error = '%v'", target, err.Error())
  }

  s := string(bRead)

  if expected != s {
    t.Errorf("Expected to read string: '%v'. Instead got, '%v'", expected, s)
  }

  _ = f.Close()
}

func TestFileHelper_SwapBasePath_01(t *testing.T) {

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  oldBasePath := "../filesfortest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"

  newPath, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_02(t *testing.T) {

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  oldBasePath := "../filesforTest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"

  newPath, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_03(t *testing.T) {

  targetPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  oldBasePath := "../filesforTest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  _, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.SwapBasePath(...) " +
      "NO ERROR WAS GENERATED!")
  }

}

func createALogTestBottomDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if !fh.DoesFileExist(targetDir) {
    err2 := fh.MakeDirAll(targetDir)

    if err2 != nil {
      return err2
    }
  }

  targetFile := fh.JoinPathsAdjustSeparators(targetDir, alogFile)

  if fh.DoesFileExist(targetFile) {
    err3 := fh.DeleteDirFile(targetFile)
    if err3 != nil {
      return err3
    }
  }

  f, err4 := fh.CreateFile(targetFile)

  if err4 != nil {
    return err4
  }

  nowTime := appLib.DateTimeUtility{}.GetDateTimeNanoSecText(time.Now().Local())

  _, err5 := f.WriteString("Sample Write - " + nowTime + "/n")

  if err5 != nil {
    _ = f.Close()
    return err5
  }

  _, err6 := f.WriteString("File Name: " + targetFile)

  if err6 != nil {
    _ = f.Close()
    return err6
  }

  _ = f.Close()
  return nil
}

func deleteALogTestBottomDirTargetDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if fh.DoesFileExist(targetDir) {
    err2 := fh.DeleteDirPathAll(targetDir)

    if err2 != nil {
      return err2
    }

    if fh.DoesFileExist(targetDir) {
      return errors.New("File still exists:" + targetDir)
    }
  }

  return nil
}
