package pathfileops

import (
  appLib "MikeAustin71/pathfileopsgo/appLibs"
  "errors"
  "fmt"
  "io"
  "strings"
  "testing"
  "time"
)

func TestFileHelper_IsAbsolutePath_01(t *testing.T) {

  fh := FileHelper{}
  commonDir := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/" +
    "level_03_dir/level_3_1_test.txt")

  result := fh.IsAbsolutePath(commonDir)

  if result == true {
    t.Error("IsAbsolutePath result is INVALID. Relative path classified as Absolute path!")
  }

}

func TestFileHelper_IsAbsolutePath_02(t *testing.T) {

  fh := FileHelper{}
  absPathDir := fh.AdjustPathSlash("D:/gowork/src/MikeAustin71/pathfileopsgo/filesfortest/" +
    "levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result := fh.IsAbsolutePath(absPathDir)

  if result == false {
    t.Error("IsAbsolutePath result is INVALID. Absolute path classified as Relative Path!")
  }

}

func TestFileHelper_IsPathFileString_01(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir/level_3_1_test.txt")

  expectedPathFile := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir\\level_3_1_test.txt")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if pathFileType != PathFileType.PathFile() {
    t.Errorf("Expected PathFileTypeCode='PathFile'. Instead, PathFileTypeCode='%v' ",
      pathFileType.String())
  }

  absExpectedPathFile, err := fh.MakeAbsolutePath(expectedPathFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPathFile). "+
      "expectedPathFile='%v' Error='%v' ", expectedPathFile, err.Error())
  }

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }

}

func TestFileHelper_IsPathFileString_02(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir/iDoNotExist.txt")

  expectedPathFile := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir\\iDoNotExist.txt")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if pathFileType != PathFileType.PathFile() {
    t.Errorf("Expected PathFileTypeCode='PathFile'. Instead, PathFileTypeCode='%v' ",
      pathFileType.String())
  }

  absExpectedPathFile, err := fh.MakeAbsolutePath(expectedPathFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPathFile). "+
      "expectedPathFile='%v' Error='%v' ", expectedPathFile, err.Error())
  }

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }

}

func TestFileHelper_IsPathFileString_03(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir")

  expectedPathFile := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if pathFileType != PathFileType.Path() {
    t.Errorf("Expected PathFileTypeCode='PathFile'. Instead, PathFileTypeCode='%v' ",
      pathFileType.String())
  }

  absExpectedPathFile, err := fh.MakeAbsolutePath(expectedPathFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPathFile). "+
      "expectedPathFile='%v' Error='%v' ", expectedPathFile, err.Error())
  }

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }

}

func TestFileHelper_IsPathFileString_04(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  expectedPathFile := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\iDoNotExist")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if pathFileType != PathFileType.Indeterminate() {
    t.Errorf("Expected PathFileTypeCode='PathFile'. Instead, PathFileTypeCode='%v' ",
      pathFileType.String())
  }

  absExpectedPathFile, err := fh.MakeAbsolutePath(expectedPathFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPathFile). "+
      "expectedPathFile='%v' Error='%v' ", expectedPathFile, err.Error())
  }

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }

}

func TestFileHelper_IsPathFileString_05(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("D:")

  expectedPathFile := fh.AdjustPathSlash("D:")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if pathFileType != PathFileType.Volume() {
    t.Errorf("Expected PathFileTypeCode='Volume'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", pathFileType.String(), absolutePath)
  }

  absExpectedPathFile := strings.ToLower(expectedPathFile)

  absolutePath = strings.ToLower(absolutePath)

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }
}

func TestFileHelper_JoinPathsAdjustSeparators_01(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }
}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_02(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_03(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_04(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_05(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common//")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_06(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common//")
  file1 := "//xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_07(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/")
  path12, err := fh.GetAbsPathFromFilePath(path1)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(path1) path1='%v'  Error='%v'", path1, err.Error())
  }

  file1 := "//xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")
  expected12, err := fh.GetAbsPathFromFilePath(expected1)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expected1) expected1='%v'  Error='%v'", expected1, err.Error())
  }

  result1 := fh.JoinPathsAdjustSeparators(path12, file1)

  if result1 != expected12 {
    t.Errorf("Joined path and file name. Expected result '%v'. Instead result='%v'",
      expected12, result1)
  }

}

func TestFileHelper_JoinPaths_03(t *testing.T) {
  fh := FileHelper{}
  path1 := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"
  file1 := "level_3_1_test.txt"
  expected1 := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result1 := fh.JoinPaths(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinBadPaths_04(t *testing.T) {
  fh := FileHelper{}
  path1 := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"
  file1 := "./level_3_1_test.txt"

  expected1 := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result1 := fh.JoinPaths(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
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
