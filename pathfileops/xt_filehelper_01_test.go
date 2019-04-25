package pathfileops

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "testing"
)

const (
  alogtopTest2Text  = "../logTest/topTest2.txt"
  alogTestBottomDir = "../logTest/CmdrX"
  alogFile          = "CmdrX.log"
  aLoremIpsumTxt    = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum at congue nisi. Fusce viverra non urna et pulvinar. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sodales in nulla at ultricies. Pellentesque nisi mi, efficitur nec magna ac, laoreet efficitur ligula. Phasellus non justo justo. Suspendisse lacus dui, euismod hendrerit dignissim et, pellentesque fermentum ipsum. Duis tempus ex vitae dui commodo, sed sagittis arcu volutpat. Nam imperdiet, enim hendrerit maximus rhoncus, mauris enim convallis orci, non tincidunt leo tortor id lorem. Cras egestas orci non eros venenatis, quis aliquet orci maximus. Duis gravida augue sit amet tristique sagittis. Sed enim risus, suscipit at odio at, pretium facilisis elit. Morbi sit amet vestibulum ipsum. Ut eu turpis arcu."
)

func TestFileHelper_AddPathSeparatorToEndOfPathStr_01(t *testing.T) {
  fh := FileHelper{}

  var notOsSeparator string

  if '\\' == os.PathSeparator {
    notOsSeparator = string('/')
  } else {
    notOsSeparator = string('\\')
  }

  strOsPathSep := string(os.PathSeparator)

  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  testPath = testPath + notOsSeparator

  expectedPath := ".." + strOsPathSep + "filesfortest" + strOsPathSep +
    "newfilesfortest" + strOsPathSep

  actualPath, err := fh.AddPathSeparatorToEndOfPathStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.AddPathSeparatorToEndOfPathStr(testPath). "+
      "Error='%v' ", err.Error())
  }

  if expectedPath != actualPath {
    t.Errorf("Expected Path='%v'. Instead, Path='%v' ",
      expectedPath, actualPath)
  }

}

func TestFileHelper_AddPathSeparatorToEndOfPathStr_02(t *testing.T) {
  fh := FileHelper{}

  strOsPathSep := string(os.PathSeparator)

  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  expectedPath := ".." + strOsPathSep + "filesfortest" + strOsPathSep +
    "newfilesfortest" + strOsPathSep

  actualPath, err := fh.AddPathSeparatorToEndOfPathStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.AddPathSeparatorToEndOfPathStr(testPath). "+
      "Error='%v' ", err.Error())
  }

  if expectedPath != actualPath {
    t.Errorf("Expected Path='%v'. Instead, Path='%v' ",
      expectedPath, actualPath)
  }

}

func TestFileHelper_AddPathSeparatorToEndOfPathStr_03(t *testing.T) {
  fh := FileHelper{}

  strOsPathSep := string(os.PathSeparator)

  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")

  expectedPath := ".." + strOsPathSep + "filesfortest" + strOsPathSep +
    "newfilesfortest" + strOsPathSep

  actualPath, err := fh.AddPathSeparatorToEndOfPathStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.AddPathSeparatorToEndOfPathStr(testPath). "+
      "Error='%v' ", err.Error())
  }

  if expectedPath != actualPath {
    t.Errorf("Expected Path='%v'. Instead, Path='%v' ",
      expectedPath, actualPath)
  }

}

func TestFileHelper_AddPathSeparatorToEndOfPathStr_04(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.AddPathSeparatorToEndOfPathStr("")

  if err == nil {
    t.Error("Expected error return from fh.AddPathSeparatorToEndOfPathStr(\"\") " +
      "because input parameter is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_AddPathSeparatorToEndOfPathStr_05(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.AddPathSeparatorToEndOfPathStr("      ")

  if err == nil {
    t.Error("Expected error return from fh.AddPathSeparatorToEndOfPathStr(\"\") " +
      "because input parameter consists of all space characters. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_AreSameFile_01(t *testing.T) {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  fh := FileHelper{}

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  correctedFile2 := correctedFile1

  relFile1 := fh.AdjustPathSlash(correctedFile1)

  relFile2 := fh.AdjustPathSlash(correctedFile2)

  filesAreSame, err := fh.AreSameFile(relFile1, relFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(relFile1, relFile2). "+
      "relFile1='%v'\nrelFile2='%v'\nError='%v'",
      relFile1, relFile2, err.Error())
  }

  if !filesAreSame {
    t.Error("Error: Expected file comparison='true'. Instead, file comparison='false'.")
  }

}

func TestFileHelper_AreSameFile_02(t *testing.T) {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  fh := FileHelper{}

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_3_test.txt"

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  filesAreSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(correctedFile1, correctedFile2). "+
      "relFile1='%v'\nrelFile2='%v'\nError='%v'",
      correctedFile1, correctedFile2, err.Error())
  }

  if filesAreSame {
    t.Error("Error: Expected file comparison='false'. Instead, file comparison='true'.")
  }

}

func TestFileHelper_AreSameFile_03(t *testing.T) {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  fh := FileHelper{}

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  filesAreSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(correctedFile1, correctedFile2). "+
      "correctedFile1='%v'\ncorrectedFile2='%v'\nError='%v'",
      correctedFile1, correctedFile2, err.Error())
  }

  if !filesAreSame {
    t.Error("Error: Expected file comparison='true'. Instead, file comparison='false'.")
  }

}

func TestFileHelper_AreSameFile_04(t *testing.T) {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  rawFile1 := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir" +
    "/level_04_dir/iDoNotExist1.txt"

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  fh := FileHelper{}

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  if os.PathSeparator == '/' {
    correctedFile1 = strings.ReplaceAll(correctedFile1, "/", "\\")
  }

  if os.PathSeparator == '\\' {
    correctedFile1 = strings.ReplaceAll(correctedFile1, "\\", "/")
  }

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an error return from fh.AreSameFile" +
      "(correctedFile1, correctedFile2) " +
      "because 'relFile1' contained INVALID path separators.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_AreSameFile_05(t *testing.T) {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist2.txt"

  fh := FileHelper{}

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  filesAreSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(correctedFile1, correctedFile2). "+
      "correctedFile1='%v'\ncorrectedFile2='%v'\nError='%v'",
      correctedFile1, correctedFile2, err.Error())
  }

  if filesAreSame {
    t.Error("Error: Expected file comparison='false'. Instead, file comparison='true'.")
  }

}

func TestFileHelper_AreSameFile_06(t *testing.T) {

  correctedFile1 := ""

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist2.txt"

  fh := FileHelper{}

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an err return from fh.AreSameFile(correctedFile1, correctedFile2) " +
      "because correctedFile1 is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_AreSameFile_07(t *testing.T) {

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  correctedFile2 := ""

  fh := FileHelper{}

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an err return from fh.AreSameFile(correctedFile1, correctedFile2) " +
      "because correctedFile2 is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_ChangeWorkingDir_01(t *testing.T) {

  fh := FileHelper{}

  currAbsDir, err := fh.GetAbsCurrDir()

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsCurrDir(). "+
      "Error='%v'", err.Error())
  }

  targetPath := fh.AdjustPathSlash("../filesfortest/iDontExist")

  err = fh.ChangeWorkingDir(targetPath)

  if err == nil {
    t.Error("Expected error return from fh.ChangeWorkingDir(targetPath) " +
      "because targetPath does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  currAbsDir2, err := fh.GetAbsCurrDir()

  if err != nil {
    t.Errorf("Error returned by #2 fh.GetAbsCurrDir(). "+
      "Error='%v'", err.Error())
  }

  currAbsDir2 = strings.ToLower(currAbsDir2)
  currAbsDirX := strings.ToLower(currAbsDir)

  if currAbsDirX != currAbsDir2 {

    err = fh.ChangeWorkingDir(currAbsDir)

    if err != nil {
      t.Errorf("Failed to reset current working directory %v. "+
        "Error='%v' ", currAbsDir, err.Error())
    }

  }

}

func TestFileHelper_ChangeWorkingDir_02(t *testing.T) {

  fh := FileHelper{}

  currAbsDir, err := fh.GetAbsCurrDir()

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsCurrDir(). "+
      "Error='%v'", err.Error())
  }

  targetPath, err := fh.MakeAbsolutePath("../filesfortest/newfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(). "+
      "Error='%v'", err.Error())
  }

  err = fh.ChangeWorkingDir(targetPath)

  if err != nil {
    t.Errorf("Error returned by fh.ChangeWorkingDir(targetPath). "+
      "targetPath='%v' Error='%v'", targetPath, err.Error())
  }

  currAbsDir2, err := fh.GetAbsCurrDir()

  if err != nil {
    t.Errorf("Error returned by #2 fh.GetAbsCurrDir(). "+
      "Error='%v'", err.Error())
  }

  currAbsDir2 = strings.ToLower(currAbsDir2)
  currAbsDirX := strings.ToLower(targetPath)

  if currAbsDirX != currAbsDir2 {

    t.Errorf("Expected new working directory='%v'. Instead, "+
      "new working directory='%v' ",
      currAbsDirX, currAbsDir2)

  }

  err = fh.ChangeWorkingDir(currAbsDir)

  if err != nil {
    t.Errorf("Failed to reset current working directory %v. "+
      "Error='%v' ", currAbsDir, err.Error())
  }

}

func TestFileHelper_ChangeWorkingDir_03(t *testing.T) {

  err := FileHelper{}.ChangeWorkingDir("")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.ChangeWorkingDir(\"\") " +
      "because the input parameter is an empty string.")
  }

}

func TestFileHelper_ChangeWorkingDir_04(t *testing.T) {

  err := FileHelper{}.ChangeWorkingDir("   ")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.ChangeWorkingDir(\"   \") " +
      "because the input parameter consists of all spaces.")
  }

}

func TestFileHelper_CleanDirStr_01(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")
  expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). "+
      "testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_02(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
  expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_03(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_04(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/../dir2/dir3")

  _, _, err := fh.CleanDirStr(testPath)

  if err == nil {
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). testPath='%v'", testPath)
  }

}

func TestFileHelper_CleanDirStr_05(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash(".../filesfortest/newfilesfortest")

  _, _, err := fh.CleanDirStr(testPath)

  if err == nil {
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). It was NOT. testPath='%v'", testPath)
  }

}

func TestFileHelper_CleanDirStr_06(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../..../filesfortest/newfilesfortest")

  _, _, err := fh.CleanDirStr(testPath)

  if err == nil {
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). It was NOT. testPath='%v'", testPath)
  }

}

func TestFileHelper_CleanDirStr_07(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")
  expectedDirName := fh.AdjustPathSlash(".")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_08(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("/xt_dirmgr_01_test.go")
  expectedDirName := fh.AdjustPathSlash("")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if !isDirEmpty {
    t.Error("Expected isDirEmpty='true'. Instead, isDirEmpty='false'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_09(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../../")
  expectedDirName := fh.AdjustPathSlash("../..")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_10(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("D:/")
  expectedDirName := "D:"

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
  }

}

func TestFileHelper_CleanDirStr_11(t *testing.T) {

  _, _, err := FileHelper{}.CleanDirStr("")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CleanDirStr(\"\") " +
      "because the input parameter is an empty string. However, NO ERROR WAS RETURNED! ")
  }
}

func TestFileHelper_CleanDirStr_12(t *testing.T) {

  _, _, err := FileHelper{}.CleanDirStr("      ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CleanDirStr(\"     \") " +
      "because the input parameter consists of all spaces. " +
      "However, NO ERROR WAS RETURNED! ")
  }
}

func TestFileHelper_CleanFileNameExtStr_01(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")
  expectedFileNameExt := "newerFileForTest_01.txt"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_02(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("newerFileForTest_01.txt")
  expectedFileNameExt := "newerFileForTest_01.txt"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_03(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
  _, _, err := fh.CleanFileNameExtStr(testPath)

  if err == nil {
    t.Errorf("Expected error returned by fh.CleanFileNameExtStr(testPath). testPath='%v'. But, no Error was returned. ", testPath)
  }

}

func TestFileHelper_CleanFileNameExtStr_04(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_05(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("     ")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"    \") " +
      "because the input parameter consists of all spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_06(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("...\\")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"    \") " +
      "because the input parameter includes 3-dots ('...'). " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_07(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/.gitignore")
  expectedFileNameExt := ".gitignore"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_08(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01")
  expectedFileNameExt := "newerFileForTest_01"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'",
      testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'",
      expectedFileNameExt, result)
  }

}

func TestFileHelper_ConvertOctalToDecimal_01(t *testing.T) {

  fh := FileHelper{}
  expectedValue := 511

  octalValue := 777

  mode := fh.ConvertOctalToDecimal(octalValue)

  if expectedValue != mode {
    t.Errorf("Error: Expected Value='%v'. Instead, value='%v' ",
      expectedValue, mode)
  }

}

func TestFileHelper_ConvertOctalToDecimal_02(t *testing.T) {

  fh := FileHelper{}
  expectedValue := 438

  octalValue := 666

  mode := fh.ConvertOctalToDecimal(octalValue)

  if expectedValue != mode {
    t.Errorf("Error: Expected Value='%v'. Instead, value='%v' ",
      expectedValue, mode)
  }

}

func TestFileHelper_ConvertDecimalToOctal_01(t *testing.T) {

  fh := FileHelper{}

  expectedOctalValue := 777

  initialDecimalValue := 511

  actualOctalValue := fh.ConvertDecimalToOctal(initialDecimalValue)

  if expectedOctalValue != actualOctalValue {
    t.Errorf("Error: Expected ocatal value='%v'. Instead, actual ocatal value='%v' ",
      expectedOctalValue, actualOctalValue)
  }

}

func TestFileHelper_ConvertDecimalToOctal_02(t *testing.T) {

  fh := FileHelper{}

  expectedOctalValue := 666

  initialDecimalValue := 438

  actualOctalValue := fh.ConvertDecimalToOctal(initialDecimalValue)

  if expectedOctalValue != actualOctalValue {
    t.Errorf("Error: Expected ocatal value='%v'. Instead, actual ocatal value='%v' ",
      expectedOctalValue, actualOctalValue)
  }

}

func TestFileHelper_CopyFileByIo_01(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestCopyFile80179658.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
      "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

  err = FileHelper{}.CopyFileByIo("", destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(\"\", destFile) " +
      "because input parameter source file is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

}

func TestFileHelper_CopyFileByIo_02(t *testing.T) {

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile). "+
        "rawSrcFile='%v' Error='%v' ", rawSrcFile, err.Error())
    }
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "")

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,\"\") " +
      "because input parameter destination file is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CopyFileByIo_03(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestCopyFile2047552.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
        "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
    }
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

  rawSrcFile := "..\\checkfiles\\iDoNOTExist.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile). "+
        "rawSrcFile='%v' Error='%v' ", rawSrcFile, err.Error())
    }
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,destFile) " +
      "because input parameter source file does not exist. " +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

}

func TestFileHelper_CopyFileByIo_04(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
        "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
    }
  }

  srcFile := destFile

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile) " +
      "because input parameter source file is equivalent to destination file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_05(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchDestTestKJ5901375.txt"
  fh := FileHelper{}

  destFile, err := fh.MakeAbsolutePath(rawDestFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
        "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
    }
  }

  err = FileHelper{}.CopyFileByIo("   ", destFile)

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(\"   \", destFile)" +
      "because input parameter source file name consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if fh.DoesFileExist(destFile) {
    err = fh.DeleteDirFile(destFile)
    if err != nil {
      t.Errorf("Error returned from last attempt to delte destFile. "+
        "fh.DeleteDirFile(destFile) destFile='%v' Error='%v' ", destFile, err.Error())
    }
  }
}

func TestFileHelper_CopyFileByIo_06(t *testing.T) {

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  fh := FileHelper{}

  srcFile, err := fh.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile). "+
      "rawSrcFile='%v' Error='%v' ", rawSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "   ")

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(src, \"    \")" +
      "because input parameter destination file consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CopyFileByIo_07(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
        "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
    }
  }

  srcFile := ""

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile) " +
      "because input parameter source file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_08(t *testing.T) {

  destFile := ""

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    if err != nil {
      t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile). "+
        "rawSrcFile='%v' Error='%v' ", rawSrcFile, err.Error())
    }
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile) " +
      "because input parameter destination file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_09(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\Level01\\Level02\\TestFile001.txt")
  if !fh.DoesFileExist(srcFile) {
    fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

    if err != nil {
      t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'. Error='%v'", srcFile, err.Error())
    }

    err = fmgr.CreateDirAndFile()

    if err != nil {
      t.Errorf("Error returned by FileMgr{}.CreateDirAndFile(). srcFile='%v'. Error='%v'", srcFile, err.Error())
    }

    doesFileExist, err := fmgr.DoesThisFileExist()

    if err != nil {
      t.Errorf("Error returned by FileMgr{}.DoesThisFileExist(). srcFile='%v'. Error='%v'", srcFile, err.Error())
    }

    if !doesFileExist {
      t.Errorf("Failed to create Source File == '%v'", srcFile)
    }

  }

  destFile := fh.AdjustPathSlash("..\\logTest\\TestFile002.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Error(fmt.Sprintf("Received Error while deleting destination file '%v', Error:", destFile), err)
    }
  }

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Error(fmt.Sprintf("Error while Copying Source File, '%v' to  Destination File '%v', Error:", srcFile, destFile), err)
  }

  if !fh.DoesFileExist(destFile) {
    t.Error(fmt.Sprintf("Expected destination file: '%v' does NOT Exist.", destFile))
  }
}


func TestFileHelper_CopyFileByIo_10(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestCopyJKC90847211.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
      "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
  }
  _ = FileHelper{}.DeleteDirFile(destFile)

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_2_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
      "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(\"\", destFile) " +
      "because input parameter source file is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

}


func TestFileHelper_CopyFileByLink_01(t *testing.T) {

  testDestFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink("", destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(\"\", destFile) " +
      "because src parameter was an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_02(t *testing.T) {

  testDestFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink("     ", destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(\"    \", destFile) " +
      "because src parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_03(t *testing.T) {

  testSrcFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, "")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, \"\") " +
      "because destination parameter was an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_04(t *testing.T) {

  testSrcFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, "     ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, \"    \") " +
      "because destination parameter consists of all space characters. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_05(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  testDestFile := "../checkfiles/scratchDestTileJJ459821.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, destFile) " +
      "because srcFile does not exist. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if fh.DoesFileExist(destFile) == true {
    err = fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error attempting file deletion of File='%v'. Error='%v' ",
        destFile, err.Error())
    }
  }
}

func TestFileHelper_CopyFileByLink_06(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, srcFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, srcFile) because Source File and Destination File are equivalent.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_07(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  destFile := ""

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Destination File is an EMPTY STRING!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_08(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  destFile := "   "

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Destination File consists" +
      "entirely of blank spaces!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_09(t *testing.T) {

  fh := FileHelper{}

  srcFile := ""

  testDestFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  destFile, err := fh.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Source File is an EMPTY STRING!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_10(t *testing.T) {

  fh := FileHelper{}

  srcFile := "   "

  testDestFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  destFile, err := fh.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Source File consists " +
      "entirely of blank spaces\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_11(t *testing.T) {

  fh := FileHelper{}

  srcFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  destFile := "../createFilesTest/scratchXFg924198.txt"

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  }

  err := fh.CopyFileByLink(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByLink(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
        "Deletion of Destination File Failed. "+
        "It cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  } else {
    t.Errorf("Error: CopyFileByLink Failed. Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
  }

}

func TestFileHelper_CopyFileByLink_12(t *testing.T) {

  fh := FileHelper{}

  srcFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  destFile := "../createFilesTest/scratchWRn877214.txt"

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  }

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by ")
  }

  if !fh.DoesFileExist(destFile) {

    t.Errorf("Error: Setup Failed. Target destination file was NOT created! "+
      "fh.CopyFileByIo(srcFile, destFile) FAILED!\n"+
      "srcFile='%v'\ndestFile='%v'", srcFile, destFile)
    return

  }

  err = fh.CopyFileByLink(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByLink(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
        "Deletion of Destination File Failed. "+
        "It cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  } else {
    t.Errorf("Error: CopyFileByLink Failed. Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
  }

}

func TestFileHelper_CreateFile_01(t *testing.T) {

  err := deleteALogTestBottomDirTargetDir()
  if err != nil {
    t.Error("Failed to delete target Directory:", err)
  }

  err = createALogTestBottomDir()
  if err != nil {
    t.Error("Failed to delete target Directory:", err)
  }
}

func TestFileHelper_CreateFile_02(t *testing.T) {
  // Uses 'Create' to overwrite existing file
  tstFile := "..//logTest//testoverwrite//TestOverwrite001.txt"
  fh := FileHelper{}

  if fh.DoesFileExist(tstFile) {
    err := fh.DeleteDirFile(tstFile)
    if err != nil {
      t.Error(fmt.Sprintf("Error: Deletion Failed On File %v !", tstFile))
    }
  }

  if fh.DoesFileExist(tstFile) {
    t.Error(fmt.Sprintf("Error: Deletion Failed! File %v should not exist!", tstFile))
  }

  f, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Error(fmt.Sprintf("Error: Create File Failed for file: %v", tstFile))
  }

  _, err4 := f.WriteString(aLoremIpsumTxt)

  if err4 != nil {
    _ = f.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err4)
  }

  err = f.Close()

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // Now recreate the original file. It should be
  // truncated with the old contents deleted.
  f2, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Errorf("Error: Re-Creating File %v", tstFile)
  }

  fOvrWriteTxt := "Test Over Write and existing file using Create()"

  _, err5 := f2.WriteString(fOvrWriteTxt)

  if err5 != nil {
    _ = f2.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err5)
  }

  err = f2.Close()

  if err != nil {
    t.Errorf("Error closing f2. %v", err.Error())
  }

  dat, err := ioutil.ReadFile(tstFile)

  if err != nil {
    t.Errorf("Error Reading Re-Written Text for File:'%v' Error='%v'", tstFile, err)
  }

  s := string(dat)

  if s != fOvrWriteTxt {
    t.Errorf("Was expecting to read text: '%v', instead received text: %v", fOvrWriteTxt, s)
  }

}

func TestFileHelper_CreateFile_03(t *testing.T) {

  _, err := FileHelper{}.CreateFile("")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CreateFile(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CreateFile_04(t *testing.T) {

  _, err := FileHelper{}.CreateFile("    ")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CreateFile(\"   \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_DeleteDirFile_01(t *testing.T) {

  err := FileHelper{}.DeleteDirFile("")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirFile(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_DeleteDirFile_02(t *testing.T) {

  err := FileHelper{}.DeleteDirFile("   ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirFile(\"  \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_DeleteDirPathAll_01(t *testing.T) {

  err := FileHelper{}.DeleteDirPathAll("")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirPathAll(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_DeleteDirPathAll_02(t *testing.T) {

  err := FileHelper{}.DeleteDirPathAll("    ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirPathAll(\"      \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_DeleteDirPathAll_03(t *testing.T) {

  targetPath := "../createFilesTest/pathDoesNotExist/fileDoesNotExist.txt"

  err := FileHelper{}.DeleteDirPathAll(targetPath)

  if err != nil {
    t.Errorf("An error was NOT expected from FileHelper{}.DeleteDirPathAll(targetPath) "+
      "because the input parameter targetPath does not exist. "+
      "However, AN ERROR WAS RETURNED! targetPath='%v' Error='%v' ",
      targetPath, err.Error())
  }

}

func TestFileHelper_DoesFileExist_01(t *testing.T) {

  fh := FileHelper{}

  doesFileExist := fh.DoesFileExist("")

  if doesFileExist {
    t.Error("Expected doesFileExist='false' because input parameter " +
      "for fh.DoesFileExist(\"\") is an " +
      "empty string. However, doesFileExist='true'!")
  }

}

func TestFileHelper_DoesFileExist_02(t *testing.T) {

  fh := FileHelper{}

  doesFileExist := fh.DoesFileExist("   ")

  if doesFileExist {
    t.Error("Expected doesFileExist='false' because input parameter " +
      "for fh.DoesFileExist(\"  \") consists entirely of blank spaces. " +
      "However, doesFileExist='true'!")
  }

}

func TestFileHelper_DoesFileInfoExist_01(t *testing.T) {
  fh := FileHelper{}

  doesFileExist, fInfo, err := fh.DoesFileInfoExist("")

  if err == nil {
    t.Error("Expected error from fh.DoesFileInfoExist(\"\") because " +
      "input parameter is an empty string. However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist != false {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") is an empty string. " +
      "However, doesFileExist=='true'!")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") is an empty string. " +
      "However, fInfo is NOT 'nil'!")
  }

}

func TestFileHelper_DoesFileInfoExist_02(t *testing.T) {
  fh := FileHelper{}

  doesFileExist, fInfo, err := fh.DoesFileInfoExist("   ")

  if err == nil {
    t.Error("Expected error from fh.DoesFileInfoExist(\"    \") because " +
      "input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist != false {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") consists entirely of blank spaces. " +
      "However, doesFileExist=='true'!")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(\"    \") consists entirely of blank spaces. " +
      "However, fInfo is NOT 'nil'!")
  }

}

func TestFileHelper_DoesFileInfoExist_03(t *testing.T) {

  fh := FileHelper{}

  testFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  expectedFileName := "level_0_1_test.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DoesFileInfoExist(testFile). "+
      "testFile='%v' Error='%v' ", testFile, err.Error())
  }

  if doesFileExist == false {
    t.Error("Expected doesFileExist=='true' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) actually exists. " +
      "However, doesFileExist=='false' !")
  }

  if fInfo == nil {
    t.Error("Expected fInfo!='nil' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) actually exists. " +
      "However, fInfo IS 'nil'!")
  }

  actualFileName := strings.ToLower(fInfo.Name())

  if expectedFileName != actualFileName {
    t.Errorf("Expected actual file name='%v'. Instead, actual file name='%v'.",
      expectedFileName, actualFileName)
  }

}

func TestFileHelper_DoesFileInfoExist_04(t *testing.T) {

  fh := FileHelper{}

  testFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err == nil {
    t.Error("Expected an error return from fh.DoesFileInfoExist(testFile). " +
      "because 'testFile' does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist == true {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) DOES NOT EXIST. " +
      "However, doesFileExist=='true' !")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) DOES NOT EXIST. " +
      "However, fInfo IS NOT 'nil'!")
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_01(t *testing.T) {
  rawtestStr := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if !doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'false'. Expected a return value of 'true' because testStr ends "+
      "with path separator.  testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_02(t *testing.T) {

  rawtestStr := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'true'. Expected a return value of 'false' because testStr does NOT "+
      "end with a path separator. testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_03(t *testing.T) {

  rawtestStr := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'true'. Expected a return value of 'false' because testStr does NOT "+
      "end with a path separator. testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_04(t *testing.T) {

  fh := FileHelper{}
  doesEndWithSep := fh.DoesStringEndWithPathSeparator("")

  if doesEndWithSep {
    t.Error("Error: fh.DoesStringEndWithPathSeparator(\"\") returned " +
      "'true'. Expected a return value of 'false' because input parameter " +
      "is an empty string. ")
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_05(t *testing.T) {

  fh := FileHelper{}
  doesEndWithSep := fh.DoesStringEndWithPathSeparator("    ")

  if doesEndWithSep {
    t.Error("Error: fh.DoesStringEndWithPathSeparator(\"   \") returned " +
      "'true'. Expected a return value of 'false' because input parameter " +
      "consists entirely of blank spaces. ")
  }

}
