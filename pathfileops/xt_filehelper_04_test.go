package pathfileops

import (
  "fmt"
  "strings"
  "testing"
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

  expectedFileType := PathFileType.Indeterminate()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

  expectedFileType := PathFileType.Volume()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
  }

  absExpectedPathFile := strings.ToLower(expectedPathFile)

  absolutePath = strings.ToLower(absolutePath)

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }
}

func TestFileHelper_IsPathFileString_06(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("D:\\")

  expectedPathFile := fh.AdjustPathSlash("D:\\")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.Path()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
  }

  absExpectedPathFile := strings.ToLower(expectedPathFile)

  absolutePath = strings.ToLower(absolutePath)

  if absExpectedPathFile != absolutePath {
    t.Errorf("Error: Expected 'absolutePath'='%v'. Instead, 'absolutePath='%v'.",
      absExpectedPathFile, absolutePath)
  }
}

func TestFileHelper_IsPathFileString_07(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("fileIDoNotExist.txt")

  expectedPathFile := fh.AdjustPathSlash("fileIDoNotExist.txt")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.File()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathFileString_08(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("fileIDoNotExist")

  expectedPathFile := fh.AdjustPathSlash("fileIDoNotExist")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.File()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathFileString_09(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("..")

  expectedPathFile := fh.AdjustPathSlash("..")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.Path()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathFileString_10(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash(".")

  expectedPathFile := fh.AdjustPathSlash(".")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.Path()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathFileString_11(t *testing.T) {

  fh := FileHelper{}
  pathFile := ""

  _, _, err := fh.IsPathFileString(pathFile)

  if err == nil {
    t.Error("Expected an error return from fh.IsPathFileString(pathFile) " +
      "because 'pathFile' is an empty string. However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileHelper_IsPathFileString_12(t *testing.T) {

  fh := FileHelper{}
  pathFile := "   "

  _, _, err := fh.IsPathFileString(pathFile)

  if err == nil {
    t.Error("Expected an error return from fh.IsPathFileString(pathFile) " +
      "because 'pathFile' consists of blank spaces. However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileHelper_IsPathFileString_13(t *testing.T) {

  fh := FileHelper{}
  pathFile := "..\\...\\"

  _, _, err := fh.IsPathFileString(pathFile)

  if err == nil {
    t.Error("Expected an error return from fh.IsPathFileString(pathFile) " +
      "because 'pathFile' 3-dots ('...'). However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileHelper_IsPathFileString_14(t *testing.T) {

  fh := FileHelper{}
  pathFile := "....\\"

  _, _, err := fh.IsPathFileString(pathFile)

  if err == nil {
    t.Error("Expected an error return from fh.IsPathFileString(pathFile) " +
      "because 'pathFile' 4-dots ('....'). However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileHelper_IsPathFileString_15(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash(".\\")

  expectedPathFile := fh.AdjustPathSlash(".\\")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.Path()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathFileString_16(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("..\\..\\..\\")

  expectedPathFile := fh.AdjustPathSlash("..\\..\\..\\")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  expectedFileType := PathFileType.Path()

  if expectedFileType != pathFileType {
    t.Errorf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "testFilePathStr='%v' ", expectedFileType.String(), pathFileType.String(), absolutePath)
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

func TestFileHelper_IsPathString_01(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("..\\..\\..\\")
  expectedPathStr := fh.AdjustPathSlash("..\\..\\..\\")

  isPath, cannotDetermine, testPathStr, err := fh.IsPathString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if true != isPath {
    t.Errorf("Expected isPath='%v'. Instead, isPath='%v' "+
      "testPathStr='%v' ", true, isPath, testPathStr)
  }

  if expectedPathStr != testPathStr {
    t.Errorf("Error: Expected 'expectedPathStr'='%v'. Instead, 'expectedPathStr='%v'.",
      expectedPathStr, testPathStr)
  }

  if false != cannotDetermine {
    t.Errorf("Error: Expected 'cannotDetermine'='%v'. Instead, 'cannotDetermine'='%v' ",
      false, cannotDetermine)
  }

}

func TestFileHelper_IsPathString_02(t *testing.T) {

  fh := FileHelper{}

  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir")

  expectedPathStr := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir")

  isPath, cannotDetermine, testPathStr, err := fh.IsPathString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if true != isPath {
    t.Errorf("Expected isPath='%v'. Instead, isPath='%v' "+
      "testPathStr='%v' ", true, isPath, testPathStr)
  }

  if expectedPathStr != testPathStr {
    t.Errorf("Error: Expected 'expectedPathStr'='%v'. Instead, 'expectedPathStr='%v'.",
      expectedPathStr, testPathStr)
  }

  if false != cannotDetermine {
    t.Errorf("Error: Expected 'cannotDetermine'='%v'. Instead, 'cannotDetermine'='%v' ",
      false, cannotDetermine)
  }

}

func TestFileHelper_IsPathString_03(t *testing.T) {

  fh := FileHelper{}

  pathFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  expectedPathStr := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\iDoNotExist")

  isPath, cannotDetermine, testPathStr, err := fh.IsPathString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
  }

  if false != isPath {
    t.Errorf("Expected isPath='%v'. Instead, isPath='%v' "+
      "testPathStr='%v' ", false, isPath, testPathStr)
  }

  if expectedPathStr != testPathStr {
    t.Errorf("Error: Expected 'expectedPathStr'='%v'. Instead, 'expectedPathStr='%v'.",
      expectedPathStr, testPathStr)
  }

  if true != cannotDetermine {
    t.Errorf("Error: Expected 'cannotDetermine'='%v'. Instead, 'cannotDetermine'='%v' ",
      true, cannotDetermine)
  }

}

func TestFileHelper_IsPathString_04(t *testing.T) {

  fh := FileHelper{}

  pathFile := fh.AdjustPathSlash("")

  _, _, _, err := fh.IsPathString(pathFile)

  if err == nil {
    t.Errorf("Expected an error return from fh.IsPathString(pathFile) " +
      "because 'pathFile' is an empty string. " +
      "However, NO ERROR WAS RETURNE!")
  }

}

func TestFileHelper_IsPathString_05(t *testing.T) {

  fh := FileHelper{}

  pathFile := fh.AdjustPathSlash("      ")

  _, _, _, err := fh.IsPathString(pathFile)

  if err == nil {
    t.Errorf("Expected an error return from fh.IsPathString(pathFile) " +
      "because 'pathFile' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNE!")
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

func TestFileHelper_JoinPathsAdjustSeparators_02(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_03(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_04(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_05(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common//")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_06(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common//")
  file1 := "//xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_07(t *testing.T) {
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

func TestFileHelper_JoinPathsAdjustSeparators_08(t *testing.T) {
  fh := FileHelper{}
  path1 := ""
  path2 := ""

  result := fh.JoinPathsAdjustSeparators(path1, path2)

  if result != path1 {
    t.Errorf("Expected result empty string. Instead result='%v'",
      result)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_09(t *testing.T) {
  fh := FileHelper{}
  path1 := "   "
  path2 := "   "

  result := fh.JoinPathsAdjustSeparators(path1, path2)

  if result != "" {
    t.Errorf("Expected result empty string. Instead result='%v'",
      result)
  }

}

func TestFileHelper_JoinPaths_01(t *testing.T) {
  fh := FileHelper{}
  path1 := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"
  file1 := "level_3_1_test.txt"
  expected1 := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result1 := fh.JoinPaths(path1, file1)

  if result1 != expected1 {
    t.Errorf("Joined path and file name. Expected result '%v', instead got: %v", expected1, result1)
  }

}

func TestFileHelper_JoinPaths_02(t *testing.T) {
  fh := FileHelper{}
  path1 := ""
  path2 := ""

  result1 := fh.JoinPaths(path1, path2)

  if result1 != "" {
    t.Errorf("Expected result = empty string, instead got result='%v'", result1)
  }

}

func TestFileHelper_JoinPaths_03(t *testing.T) {
  fh := FileHelper{}
  path1 := "   "
  path2 := "   "

  result1 := fh.JoinPaths(path1, path2)

  if result1 != "" {
    t.Errorf("Expected result = empty string, instead got result='%v'", result1)
  }

}
