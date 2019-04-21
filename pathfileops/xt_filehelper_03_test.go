package pathfileops

import (
  "strings"
  "testing"
)

func TestFileHelper_GetFileNameWithoutExt_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedFileName := "xt_dirmgr_01_test"

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty == true {
    t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_01_test")
  expectedFileName := "dirmgr_01_test"

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty == true {
    t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v'", isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileNameWithoutExt to return result == '%v' for valid file name, instead got '%v' ", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")
  expectedFileName := ""

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). "+
      "commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v', instead got %v ",
      true, isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
      "instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
  expectedFileName := "xt_dirmgr_01_test"

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty == true {
    t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")
  expectedFileName := ""

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v'. Instead isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
      "instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
  expectedFileName := "common"

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v'. Instead isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
      "instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := ""
  expectedFileName := ""

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err == nil {
    t.Error("Expected error return from fh.GetFileNameWithoutExt(commonDir) because " +
      "'commonDir' is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if true != isEmpty {
    t.Error("Expected GetFileNameWithoutExt isEmpty='true'. Instead isEmpty='false' ")
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
      "instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFileNameWithoutExt_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := "   "
  expectedFileName := ""

  result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

  if err == nil {
    t.Error("Expected error return from fh.GetFileNameWithoutExt(commonDir) because " +
      "'commonDir' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if true != isEmpty {
    t.Error("Expected GetFileNameWithoutExt isEmpty='true'. Instead isEmpty='false' ")
  }

  if result != expectedFileName {
    t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
      "instead got: %v", expectedFileName, result)
  }

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_01(t *testing.T) {

  rawPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(rawPath)

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  if firstCharIdx != 3 {
    t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
  }

  expectedLastIdx := len(adjustedPath) - 1

  if expectedLastIdx != lastCharIdx {
    t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
  }

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_02(t *testing.T) {

  rawPath := "D:/filesfortest/newfilesfortest/newerFileForTest_01.txt"
  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(rawPath)

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  if firstCharIdx != 3 {
    t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
  }

  expectedLastIdx := len(adjustedPath) - 1

  if expectedLastIdx != lastCharIdx {
    t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
  }

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_03(t *testing.T) {

  fh := FileHelper{}

  adjustedPath := ""

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

  if err == nil {
    t.Error("Expected error return from fh.GetFirstLastNonSeparatorCharIndexInPathStr" +
      "(adjustedPath) because 'adjustedPath' is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if firstCharIdx != -1 {
    t.Errorf("Expected first char index= '-1'.  Instead, first char index= '%v'", firstCharIdx)
  }

  if lastCharIdx != -1 {
    t.Errorf("Expected last index = '-1'.  Instead, last index = '%v'", lastCharIdx)
  }
}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_04(t *testing.T) {

  fh := FileHelper{}

  adjustedPath := "     "

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

  if err == nil {
    t.Error("Expected error return from fh.GetFirstLastNonSeparatorCharIndexInPathStr" +
      "(adjustedPath) because 'adjustedPath' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if firstCharIdx != -1 {
    t.Errorf("Expected first char index= '-1'.  Instead, first char index= '%v'", firstCharIdx)
  }

  if lastCharIdx != -1 {
    t.Errorf("Expected last index = '-1'.  Instead, last index = '%v'", lastCharIdx)
  }
}

func TestFileHelper_GetLastPathElement_01(t *testing.T) {

  fh := FileHelper{}
  expectedLastPathElement := "level_4_0_test.txt"

  testPathFileName := fh.AdjustPathSlash(
    "filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_04_dir/" +
      expectedLastPathElement)

  lastPathElement, err := fh.GetLastPathElement(testPathFileName)

  if err != nil {
    t.Errorf("Error returned by fh.GetLastPathElement(testPathFileName). "+
      "testPathFileName='%v' Error='%v' ", testPathFileName, err.Error())
  }

  if expectedLastPathElement != lastPathElement {
    t.Errorf("Error: Expected lastPathElement='%v'. Instead, lastPathElement='%v' ",
      expectedLastPathElement, lastPathElement)
  }

}

func TestFileHelper_GetLastPathElement_02(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.GetLastPathElement("")

  if err == nil {
    t.Error("Expected an error return from fh.GetLastPathElement(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetLastPathElement_03(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.GetLastPathElement("")

  if err == nil {
    t.Error("Expected an error return from fh.GetLastPathElement(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetLastPathElement_04(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.GetLastPathElement("  ")

  if err == nil {
    t.Error("Expected an error return from fh.GetLastPathElement(\"   \") " +
      "because the input parameter consists entirely of empty spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetLastPathElement_05(t *testing.T) {

  fh := FileHelper{}

  expectedLastPathElement := "level_04_dir"

  testPathFileName := fh.AdjustPathSlash(
    "filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/" +
      expectedLastPathElement)

  lastPathElement, err := fh.GetLastPathElement(testPathFileName)

  if err != nil {
    t.Errorf("Error returned by fh.GetLastPathElement(testPathFileName). "+
      "testPathFileName='%v' Error='%v' ", testPathFileName, err.Error())
  }

  if expectedLastPathElement != lastPathElement {
    t.Errorf("Error: Expected lastPathElement='%v'. Instead, lastPathElement='%v' ",
      expectedLastPathElement, lastPathElement)
  }

}

func TestFileHelper_GetLastPathElement_06(t *testing.T) {

  fh := FileHelper{}
  expectedLastPathElement := "level_4_0_test.txt"

  lastPathElement, err := fh.GetLastPathElement(expectedLastPathElement)

  if err != nil {
    t.Errorf("Error returned by fh.GetLastPathElement(expectedLastPathElement). "+
      "expectedLastPathElement='%v' Error='%v' ", expectedLastPathElement, err.Error())
  }

  if expectedLastPathElement != lastPathElement {
    t.Errorf("Error: Expected lastPathElement='%v'. Instead, lastPathElement='%v' ",
      expectedLastPathElement, lastPathElement)
  }

}

func TestFileHelper_GetPathFromPathFileName_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("..\\..\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("..\\..\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid file name. Instead path='%v'", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name. "+
      "Instead path=='%v' ", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Expected no error returned from fh.GetPathFromPathFileName(commonDir). "+
      "Instead an error WAS Returned. commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name. "+
      "Instead path=='%v' ", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). "+
      "commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return "+
      "'false', instead isEmpty='%v' ", isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file "+
      "name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("./")

  expectedDir := "."

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). "+
      "commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file "+
      "name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".")

  expectedDir := fh.AdjustPathSlash(".")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' "+
      "Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid "+
      "path/file name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_09(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("..")

  expectedDir := fh.AdjustPathSlash("..")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file "+
      "name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_10(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("")

  expectedDir := fh.AdjustPathSlash("")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err == nil {
    t.Errorf("Expected error to be returned from fh.GetPathFromPathFileName(commonDir). "+
      "commonDir='%v' No Error Returned!", commonDir)
  }

  if true != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file"+
      "name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_11(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("../../../")

  expectedDir := fh.AdjustPathSlash("../../..")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file "+
      "name, instead got: %v", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_12(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash(".")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Expected no error returned from fh.GetPathFromPathFileName(commonDir). "+
      "Instead an error WAS Returned. commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file"+
      "name. Instead path=='%v' ", expectedDir, result)
  }

}

func TestFileHelper_GetPathFromPathFileName_13(t *testing.T) {
  fh := FileHelper{}

  result, isEmpty, err := fh.GetPathFromPathFileName("     ")

  if err == nil {
    t.Error("Expected an error return from fh.GetPathFromPathFileName(\"   \") " +
      "because the input parameter consists entirely of spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isEmpty == false {
    t.Error("Expected isEmpty='true', instead isEmpty='false' ")
  }

  if result != "" {
    t.Errorf("Expected GetPathFromPathFileName to return path == 'empty string'.  "+
      "Instead path=='%v' ", result)
  }

}

func TestFileHelper_GetPathAndFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  expectedFileNameExt := "xt_dirmgr_01_test.go"

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v' "+
      "Error='%v'", commonDir, err.Error())
  }

  if false != bothAreEmpty {
    t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
      false, bothAreEmpty)
  }

  if pathDir != expectedDir {
    t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. "+
      "Instead, path== '%v' ", expectedDir, pathDir)
  }

  if fileNameExt != expectedFileNameExt {
    t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
      "fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
  }

}

func TestFileHelper_GetPathAndFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  expectedFileNameExt := ""

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if false != bothAreEmpty {
    t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
      false, bothAreEmpty)
  }

  if pathDir != expectedDir {
    t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ",
      expectedDir, pathDir)
  }

  if fileNameExt != expectedFileNameExt {
    t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
      "fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
  }

}

func TestFileHelper_GetPathAndFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  expectedFileNameExt := "dirmgr_test"

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  "+
      "Error='%v'", commonDir, err.Error())
  }

  if false != bothAreEmpty {
    t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
      false, bothAreEmpty)
  }

  if pathDir != expectedDir {
    t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ",
      expectedDir, pathDir)
  }

  if fileNameExt != expectedFileNameExt {
    t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
      "fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
  }

}

func TestFileHelper_GetPathAndFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("")

  expectedFileNameExt := "xt_dirmgr_01_test.go"

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if false != bothAreEmpty {
    t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ", false, bothAreEmpty)
  }

  if pathDir != expectedDir {
    t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ", expectedDir, pathDir)
  }

  if fileNameExt != expectedFileNameExt {
    t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
  }

}

func TestFileHelper_GetPathAndFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt("")

  if err == nil {
    t.Error("Expected error return from fh.GetPathAndFileNameExt(\"\") because " +
      "the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if pathDir != "" {
    t.Errorf("Expected pathDir would be an empty string. Instead, pathDir='%v'",
      pathDir)
  }

  if fileNameExt != "" {
    t.Errorf("Expected fileNameExt would be an empty string. Instead, pathDir='%v'",
      fileNameExt)
  }

  if bothAreEmpty == false {
    t.Error("Expected bothAreEmpty='true'. Instead, bothArEmpty='false'. ")
  }

}

func TestFileHelper_GetPathAndFileNameExt_06(t *testing.T) {

  fh := FileHelper{}

  pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt("    ")

  if err == nil {
    t.Error("Expected error return from fh.GetPathAndFileNameExt(\"   \") because " +
      "the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if pathDir != "" {
    t.Errorf("Expected pathDir would be an empty string. Instead, pathDir='%v'",
      pathDir)
  }

  if fileNameExt != "" {
    t.Errorf("Expected fileNameExt would be an empty string. Instead, pathDir='%v'",
      fileNameExt)
  }

  if bothAreEmpty == false {
    t.Error("Expected bothAreEmpty='true'. Instead, bothArEmpty='false'. ")
  }

}

func TestFileHelper_GetPathSeparatorIndexesInPathStr_01(t *testing.T) {

  fh := FileHelper{}

  idxs, err := fh.GetPathSeparatorIndexesInPathStr("")

  if err == nil {
    t.Error("Expected error return from fh.GetPathSeparatorIndexesInPathStr(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if len(idxs) != 0 {
    t.Errorf("Expected length of indexes='0'. Instead length of indexes='%v' ",
      len(idxs))
  }

}

func TestFileHelper_GetPathSeparatorIndexesInPathStr_02(t *testing.T) {

  fh := FileHelper{}

  idxs, err := fh.GetPathSeparatorIndexesInPathStr("     ")

  if err == nil {
    t.Error("Expected error return from fh.GetPathSeparatorIndexesInPathStr(\"     \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if len(idxs) != 0 {
    t.Errorf("Expected length of indexes='0'. Instead length of indexes='%v' ",
      len(idxs))
  }

}

// /d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops
// D:\gowork\src\MikeAustin71\pathfileopsgo\pathfileops
func TestFileHelper_GetVolumeName_01(t *testing.T) {

  fh := FileHelper{}

  volumeName := fh.GetVolumeName("")

  if volumeName != "" {
    t.Errorf("Expected an empty string return from fh.GetVolumeName(\"\") because "+
      "the input parameter is an empty string. Instead, the return value='%v' ", volumeName)
  }
}

func TestFileHelper_GetVolumeName_02(t *testing.T) {

  fh := FileHelper{}

  volumeName := fh.GetVolumeName("  ")

  if volumeName != "" {
    t.Errorf("Expected an empty string return from fh.GetVolumeName(\"\") because "+
      "the input parameter consists of blank spaces. Instead, the return value='%v' ", volumeName)
  }
}

func TestFileHelper_GetVolumeName_03(t *testing.T) {

  fh := FileHelper{}

  testVolStr := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\pathfileops"

  expectedVolName := strings.ToLower("D:")

  volumeName := fh.GetVolumeName(testVolStr)

  if expectedVolName != strings.ToLower(volumeName) {
    t.Errorf("Expected volumeName='%v'. Instead, volName='%v' ",
      expectedVolName, strings.ToLower(volumeName))
  }

}

func TestFileHelper_GetVolumeName_04(t *testing.T) {

  fh := FileHelper{}

  testVolStr := "D:\\"

  expectedVolName := strings.ToLower("D:")

  volumeName := fh.GetVolumeName(testVolStr)

  if expectedVolName != strings.ToLower(volumeName) {
    t.Errorf("Expected volumeName='%v'. Instead, volName='%v' ",
      expectedVolName, strings.ToLower(volumeName))
  }

}

func TestFileHelper_GetVolumeName_05(t *testing.T) {

  fh := FileHelper{}

  testVolStr := "D:"

  expectedVolName := strings.ToLower("D:")

  volumeName := fh.GetVolumeName(testVolStr)

  if expectedVolName != strings.ToLower(volumeName) {
    t.Errorf("Expected volumeName='%v'. Instead, volName='%v' ",
      expectedVolName, strings.ToLower(volumeName))
  }

}
