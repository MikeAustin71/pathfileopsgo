package pathfileops

import (
  "fmt"
  "os"
  "strings"
  "testing"
  "time"
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

  filesAreSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(correctedFile1, correctedFile2).\n"+
      "relFile1='%v'\nrelFile2='%v'\nError='%v'\n",
      correctedFile1, correctedFile2, err.Error())
    return
  }

  if !filesAreSame {
    t.Error("Error: Expected file comparison='true'.\n" +
      "Instead, file comparison='false'.\n")
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

  filesAreTheSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(correctedFile1, correctedFile2).\n" +
      "correctedFile1='%v'\ncorrectedFile2='%v'\nError='%v'\n",
      correctedFile1, correctedFile2, err.Error())
  }

  if !filesAreTheSame {
    t.Errorf("ERROR: Expected that AreSameFile='true'.\nInstead, AreSameFile='false'\n" +
      "correctedFile1='%v'\ncorrectedFile2='%v'\n",
      correctedFile1, correctedFile2)
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
    t.Error("Expected an error return from fh.AreSameFile(correctedFile1, correctedFile2) " +
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
    t.Error("Expected an error return from fh.AreSameFile(correctedFile1, correctedFile2) " +
      "because correctedFile2 is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_AreSameFile_08(t *testing.T) {

  correctedFile2 := ""

  correctedFile1 := ""

  fh := FileHelper{}

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an error return from fh.AreSameFile(correctedFile1, correctedFile2) " +
      "because both correctedFile1 and correctedFile2 are empty strings.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_AreSameFile_09(t *testing.T) {
  fh := FileHelper{}

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"


  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_3_test.txt"

  if os.PathSeparator == '\\' {
    rawFile1 = strings.ReplaceAll(rawFile1, "\\", "/")
    rawFile2 = strings.ReplaceAll(rawFile2, "\\", "/")
  }

  filesAreSame, err := fh.AreSameFile(rawFile1, rawFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(rawFile1, rawFile2).\n"+
      "rawFile1='%v'\nrawFile2='%v'\nError='%v'\n",
      rawFile1, rawFile2, err.Error())
    return
  }

  if filesAreSame {
    t.Error("Error: Expected file comparison='false'.\n" +
      "Instead, file comparison='true'.\n")
  }

}

func TestFileHelper_AreSameFile_10(t *testing.T) {
  fh := FileHelper{}

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"


  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_3_test.txt"

  if os.PathSeparator == '\\' {
    rawFile1 = strings.ReplaceAll(rawFile1, "\\", "/")
  }

  rawFile2 = fh.AdjustPathSlash(rawFile2)

  filesAreSame, err := fh.AreSameFile(rawFile1, rawFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(rawFile1, rawFile2).\n"+
      "rawFile1='%v'\nrawFile2='%v'\nError='%v'\n",
      rawFile1, rawFile2, err.Error())
    return
  }

  if filesAreSame {
    t.Error("Error: Expected file comparison='false'.\n" +
      "Instead, file comparison='true'.\n")
  }

}

func TestFileHelper_AreSameFile_11(t *testing.T) {
  fh := FileHelper{}

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"

  rawFile1 = fh.AdjustPathSlash(rawFile1)

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_3_test.txt"

  if os.PathSeparator == '\\' {
    rawFile2 = strings.ReplaceAll(rawFile2, "\\", "/")
  }

  filesAreSame, err := fh.AreSameFile(rawFile1, rawFile2)

  if err != nil {
    t.Errorf("Error returned by fh.AreSameFile(rawFile1, rawFile2).\n"+
      "rawFile1='%v'\nrawFile2='%v'\nError='%v'\n",
      rawFile1, rawFile2, err.Error())
    return
  }

  if filesAreSame {
    t.Error("Error: Expected file comparison='false'.\n" +
      "Instead, file comparison='true'.\n")
  }

}

func TestFileHelper_AreSameFile_12(t *testing.T) {

  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist1.txt"

  correctedFile2 := "  "

  fh := FileHelper{}

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an error return from fh.AreSameFile(correctedFile1, correctedFile2) " +
      "because correctedFile2 consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_AreSameFile_13(t *testing.T) {

  correctedFile1 := "   "

  rawFile2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\iDoNotExist2.txt"

  fh := FileHelper{}

  correctedFile2 := fh.AdjustPathSlash(rawFile2)

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an error return from fh.AreSameFile(" +
      "correctedFile1, correctedFile2)\n" +
      "because correctedFile1 consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_AreSameFile_14(t *testing.T) {

  correctedFile1 := "   "

  correctedFile2 := "  "

  fh := FileHelper{}

  _, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err == nil {
    t.Error("Expected an error return from fh.AreSameFile(" +
      "correctedFile1, correctedFile2)\n" +
      "because both correctedFile1 and correctedFile2 consist entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_ChangeFileMode_01(t *testing.T) {

  pathFileName := ""

  filePermission, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error return from FilePermissionConfig{}.New(\"-rwxrwxrwx\").\n" +
      "Error='%v'\n", err.Error())
  }

  fh := FileHelper{}

  err = fh.ChangeFileMode(pathFileName, filePermission)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileMode(pathFileName, filePermission)\n" +
      "because 'pathFileName' is an empty string. However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_ChangeFileMode_02(t *testing.T) {

  pathFileName := "    "

  filePermission, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error return from FilePermissionConfig{}.New(\"-rwxrwxrwx\").\n" +
      "Error='%v'\n", err.Error())
  }

  fh := FileHelper{}

  err = fh.ChangeFileMode(pathFileName, filePermission)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileMode(pathFileName, filePermission)\n" +
      "because 'pathFileName' consists entirely of empty spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_ChangeFileMode_04(t *testing.T) {

  pathFileName := "../createFilesTest/iDoNOTExist.txt"

  filePermission, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error return from FilePermissionConfig{}.New(\"-rwxrwxrwx\").\n" +
      "Error='%v'\n", err.Error())
  }

  fh := FileHelper{}

  err = fh.ChangeFileMode(pathFileName, filePermission)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileMode(pathFileName, filePermission)\n" +
      "because 'pathFileName' does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_ChangeFileMode_05(t *testing.T) {

  basePath := "../createFilesTest/TestFileHelper_ChangeFileMode_05"
  actualPath := basePath + "/level01"

  originalSrc := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  pathFileName := actualPath + "/" +"level_0_3_test.txt"

  fh := FileHelper{}
  var err error

  if fh.DoesFileExist(basePath) {
    err = fh.DeleteDirPathAll(basePath)

    if err != nil {
      t.Errorf("Test Setup Error: Could not delete 'basePath'!\n" +
        "basePath='%v'\nError='%v'\n", basePath, err.Error())
      return
    }
  }

  err = fh.MakeDirAll(actualPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(actualPath).\n" +
      "actualPath='%v'\nError='%v'\n", actualPath, err.Error())
    return
  }

  err = fh.CopyFileByIo(originalSrc, pathFileName)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(originalSrc, pathFileName)." +
      "originalSrc='%v'\npathFileName='%v'\nError='%v'\n",
      originalSrc, pathFileName, err.Error())
    return
  }

  filePermission, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error return from FilePermissionConfig{}.New(\"-rwxrwxrwx\").\n" +
      "Error='%v'\n", err.Error())
    return
  }

  filePermission.isInitialized = false

  err = fh.ChangeFileMode(pathFileName, filePermission)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileMode(pathFileName, filePermission)\n" +
      "because 'filePermission' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error: Could not delete 'basePath'!\n" +
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
    return
  }

}

func TestFileHelper_ChangeFileMode_06(t *testing.T) {

  basePath := "../createFilesTest/TestFileHelper_ChangeFileMode_06"
  actualPath := basePath + "/level01"

  originalSrc := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  pathFileName := actualPath + "/" +"level_0_3_test.txt"

  fh := FileHelper{}
  var err error

  if fh.DoesFileExist(basePath) {
    err = fh.DeleteDirPathAll(basePath)

    if err != nil {
      t.Errorf("Test Setup Error: Could not delete 'basePath'!\n" +
        "basePath='%v'\nError='%v'\n", basePath, err.Error())
      return
    }
  }

  err = fh.MakeDirAll(actualPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(actualPath).\n" +
      "actualPath='%v'\nError='%v'\n", actualPath, err.Error())
    return
  }

  err = fh.CopyFileByIo(originalSrc, pathFileName)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(originalSrc, pathFileName)." +
      "originalSrc='%v'\npathFileName='%v'\nError='%v'\n",
      originalSrc, pathFileName, err.Error())
    return
  }

  originalPermission, err := fh.GetFileMode(pathFileName)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.GetFileMode(pathFileName).\n" +
      "pathFileName='%v'\nError='%v'\n", pathFileName, err.Error())
    return
  }

  filePermission, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error return from FilePermissionConfig{}.New(\"-r--r--r--\").\n" +
      "Error='%v'\n", err.Error())
  }

  err = fh.ChangeFileMode(pathFileName, filePermission)

  if err != nil {
    t.Errorf("Error returned from fh.ChangeFileMode(pathFileName, filePermission)\n" +
      "pathFileName='%v'\nError='%v'\n", pathFileName, err.Error())

    _ = fh.DeleteDirPathAll(basePath)

    return
  }

  newPermission, err := fh.GetFileMode(pathFileName)

  if err != nil {
    t.Errorf("Test Verification Error returned by fh.GetFileMode(pathFileName).\n" +
      "pathFileName='%v'\nError='%v'\n", pathFileName, err.Error())

    _ = fh.DeleteDirPathAll(basePath)

    return
  }

  originalPermissionText, err := originalPermission.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Verification Error returned by originalPermission.GetPermissionTextCode().\n" +
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(basePath)

    return
  }

  originalPermissionValue := originalPermission.GetPermissionFileModeValueText()

  newPermissionText, err := newPermission.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Verification Error returned by newPermission.GetPermissionTextCode().\n" +
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(basePath)

    return
  }

  newPermissionValue := newPermission.GetPermissionFileModeValueText()

  if originalPermission.Equal(&newPermission) {
    t.Errorf("Error: Expected new File Mode to be different from old File Mode.\n" +
      "Instead, they are equal!\nOrigional Permission Text='%v' Orginal Permission Value='%v'\n" +
      "New Permission Text='%v' New Permission Value='%v'",
      originalPermissionText, originalPermissionValue, newPermissionText, newPermissionValue)
  }

  if originalPermissionText == newPermissionText {
    t.Errorf("Error: Expected new File Mode text value to be different from old File " +
      "Mode text value.\nInstead, they are the same!\n" +
      "originalPermissionText='%v' newPermissionText='%v'",
      originalPermissionText, newPermissionText)
  }

  if originalPermissionValue == newPermissionValue {
    t.Errorf("Error: Expected new File Mode numerical value to be different from old File " +
      "Mode numerical value.\nInstead, they are the same!\n" +
      "originalPermissionText='%v' newPermissionText='%v'",
      originalPermissionValue, newPermissionValue)
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error: Could not delete 'basePath'!\n" +
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
    return
  }

}

func TestFileHelper_ChangeFileTimes_01(t *testing.T) {
  originalSrc := "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"

  dest := "../checkfiles/TestFileHelper_ChangeFileTimes_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(dest).\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
    return
  }

  err = fh.MakeDirAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(dest).\n" +
      "dest='%v'\nError='%v'\n",
      dest, err.Error())
    return
  }

  destFile := dest + "/006870_ReadingFiles.htm"

  err = fh.CopyFileByIo(originalSrc, destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(originalSrc, destFile).\n" +
      "originalSrc='%v'\ndestFile='%v'\nError='%v'\n",
      originalSrc, destFile, err.Error())

    _ = fh.DeleteDirPathAll(dest)
    return
  }

  dateFormat := "2006-01-02 15:04:05.000000000 -0700 MST"

  originalModTime, _, err :=
    fh.GetFileLastModificationDate(destFile,dateFormat)

  if err != nil {
    t.Errorf("Test Setup Error returned by #1 fh.GetFileLastModificationDate(destFile,dateFormat).\n" +
      "\ndestFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)

  err = fh.ChangeFileTimes(destFile, newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Error returned by fh.ChangeFileTimes(destFile, newAccessTime, newModTime).\n" +
      "destFile='%v'\nnewAccessTime='%v'\newModTime='%v'\nError='%v'\n",
      destFile,newAccessTime.Format(dateFormat), newModTime.Format(dateFormat), err.Error())
  }

  actualModTime, _, err :=
    fh.GetFileLastModificationDate(destFile,dateFormat)

  if err != nil {
    t.Errorf("Test Setup Error returned by #2 fh.GetFileLastModificationDate(...).\n" +
      "\ndestFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  err = fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(dest)\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
  }

  if originalModTime.Year() == actualModTime.Year() {
    t.Errorf("Error Original Mod Time Year == Actual Mod Time Year!\n" +
      "Original Mod Time Year='%v'\nActual Mod Time Year='%v'\n",
      originalModTime.Year(), actualModTime.Year())
  }

  dateFormat2 := "2006-01-02 15:04:05"

  newModStr := newModTime.Format(dateFormat2)
  actualModStr := actualModTime.Format(dateFormat2)

  if newModStr != actualModStr {
    t.Errorf("ERROR: Expected actual modified time='%v'.\n" +
      "Instead, actual modified time='%v'.\n",newModStr, actualModStr)
  }

}

func TestFileHelper_ChangeFileTimes_02(t *testing.T) {

  fh := FileHelper{}

  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  testFile := ""
  err := fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'testFile' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_ChangeFileTimes_03(t *testing.T) {

  fh := FileHelper{}

  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  testFile := "      "
  err := fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'testFile' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_ChangeFileTimes_04(t *testing.T) {

  fh := FileHelper{}

  testFile := "../checkfiles/iDoNotExist.txt"
  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)

  err := fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'testFile' does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_ChangeFileTimes_05(t *testing.T) {

  fh := FileHelper{}

  testFile := "../checkfiles"
  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)

  err := fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'testFile' is NOT a file - it is a directory.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_ChangeFileTimes_06(t *testing.T) {

  fh := FileHelper{}

  originalSrc := "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"

  dest := "../checkfiles/TestFileHelper_ChangeFileTimes_06"


  err := fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(dest).\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
    return
  }

  err = fh.MakeDirAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(dest).\n" +
      "dest='%v'\nError='%v'\n",
      dest, err.Error())
    return
  }

  testFile := dest + "/006870_ReadingFiles.htm"

  err = fh.CopyFileByIo(originalSrc, testFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(originalSrc, destFile).\n" +
      "originalSrc='%v'\ntestFile='%v'\nError='%v'\n",
      originalSrc, testFile, err.Error())

    _ = fh.DeleteDirPathAll(dest)
    return
  }

  newAccessTime := time.Time{}
  newModTime := time.Date(2006,2,25,10,30,30,1250, time.Local)

  err = fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'newAccessTime' is ZERO.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(dest)\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
  }

}

func TestFileHelper_ChangeFileTimes_07(t *testing.T) {

  fh := FileHelper{}

  originalSrc := "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"

  dest := "../checkfiles/TestFileHelper_ChangeFileTimes_06"

  err := fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(dest).\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
    return
  }

  err = fh.MakeDirAll(dest)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(dest).\n" +
      "dest='%v'\nError='%v'\n",
      dest, err.Error())
    return
  }

  testFile := dest + "/006870_ReadingFiles.htm"

  err = fh.CopyFileByIo(originalSrc, testFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(originalSrc, destFile).\n" +
      "originalSrc='%v'\ntestFile='%v'\nError='%v'\n",
      originalSrc, testFile, err.Error())

    _ = fh.DeleteDirPathAll(dest)
    return
  }

  newAccessTime := time.Date(2006,2,25,10,30,30,1250, time.Local)
  newModTime := time.Time{}

  err = fh.ChangeFileTimes(testFile,newAccessTime, newModTime)

  if err == nil {
    t.Error("Expected an error return from fh.ChangeFileTimes(testFile, " +
      "newAccessTime, newModTime)\nbecause 'newModTime' is ZERO.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(dest)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(dest)\n" +
      "dest='%v'\nError='%v'\n", dest, err.Error())
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
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath).\n" +
      "testPath='%v'\n", testPath)
  }

}

func TestFileHelper_CleanDirStr_05(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash(".../filesfortest/newfilesfortest")

  _, _, err := fh.CleanDirStr(testPath)

  if err == nil {
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath).\n" +
      "It was NOT.\ntestPath='%v'\n", testPath)
  }

}

func TestFileHelper_CleanDirStr_06(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../..../filesfortest/newfilesfortest")

  _, _, err := fh.CleanDirStr(testPath)

  if err == nil {
    t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath).\n" +
      "It was NOT.\ntestPath='%v'\n", testPath)
  }

}

func TestFileHelper_CleanDirStr_07(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")
  expectedDirName := fh.AdjustPathSlash("./")

  cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanDirStr(testPath).\n" +
      "testPath='%v'\nError='%v'", testPath, err.Error())
  }

  if isDirEmpty {
    t.Error("Expected isDirEmpty='false'.\nInstead, isDirEmpty='true'\n")
  }

  if expectedDirName != cleanDirStr {
    t.Errorf("Expected cleanDirStr='%v'.\nInstead cleanDirStr='%v'\n",
      expectedDirName, cleanDirStr)
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
  expectedDirName := fh.AdjustPathSlash("../../")

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
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath).\n" +
      "testPath='%v'\nError='%v'",
      testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'.\n" +
      "Instead, isFileNameEmpty='true'\n")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'.\n" +
      "Instead, it returned '%v'\n",
      expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_03(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
  // testPath is a directory which actually exists
  _, isEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Unexpected Error returned by fh.CleanFileNameExtStr(testPath)\n" +
      "testPath='%v'\nError='%v'\n",
      testPath, err.Error())
    return
  }

  if isEmpty == false {
    t.Error("ERROR: Expected 'isEmpty' == 'true' because input parameter\n" +
      "'testPath' was an actual directory the physically exists on disk.\n" +
      "However, 'isEmpty' return value was 'false'!\n")
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
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "")

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,\"\")\n" +
      "because input parameter destination file is an empty string.\n" +
      "\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_03(t *testing.T) {

  rawDestFile := "..\\checkfiles\\TestFileHelper_CopyFileByIo_03.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

  rawSrcFile := "..\\checkfiles\\iDoNOTExist.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file does not exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)
}

func TestFileHelper_CopyFileByIo_04(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  srcFile := destFile

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file is equivalent to destination file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByIo_05(t *testing.T) {

  rawDestFile := "..\\checkfiles\\TestFileHelper_CopyFileByIo_05.txt"
  fh := FileHelper{}

  destFile, err := fh.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo("   ", destFile)

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(\"   \", destFile)\n" +
      "because input parameter source file name consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  if fh.DoesFileExist(destFile) {
    err = fh.DeleteDirFile(destFile)
    if err != nil {
      t.Errorf("Error returned from last attempt to delete destFile.\n"+
        "fh.DeleteDirFile(destFile)\ndestFile='%v'\nError='%v'\n",
        destFile, err.Error())
    }
  }
}

func TestFileHelper_CopyFileByIo_06(t *testing.T) {

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  fh := FileHelper{}

  srcFile, err := fh.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "   ")

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(src, \"    \")\n" +
      "because input parameter destination file consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_07(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  srcFile := ""

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_08(t *testing.T) {

  destFile := ""

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter destination file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_09(t *testing.T) {

  fh := FileHelper{}
  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir" +
    "\\level_03_dir\\level_3_1_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("ERROR: Setup source file DOES NOT EXIST!\n" +
      "srcFile='%v' \n", srcFile)
    return
  }

  rawDestFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_09.txt"

  destFile := fh.AdjustPathSlash(rawDestFile)

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error returned from fh.DeleteDirFile(destFile).\n" +
        "Attempt to delete prexisting version of destination file FAILED!\n" +
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("ERROR: Prexisting Destination File could NOT be Deleted!\n" +
        "Destination File:'%v'\n", destFile)
      return
    }
  }

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error while Copying Source File to  Destination File!\n" +
      "Source File='%v'\nDestination File='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Error(fmt.Sprintf("After CopyIO Destination File DOES NOT EXIST!\n" +
      "destFile='%v'\n", destFile))
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile).\n" +
    "During clean-up, the attempted deletion of the destination file FAILED!\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of the destination file during " +
      "clean-up FAILED!\ndestFile='%v'", destFile)
  }


}

func TestFileHelper_CopyFileByIo_10(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_10.txt"

  fh := FileHelper{}

  destFile := fh.AdjustPathSlash(rawDestFile)

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error retunred by fh.DeleteDirFile(destFile) during setup.\n" +
      "Attempt deletion of pre-existing version of destination file FAILED!\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Unable to delete pre-existing version of destination file!\n" +
      "destFile='%v'", destFile)
    return
  }

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_2_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  err = fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.CopyFileByIo(srcFile, destFile).\n" +
      "srcFile='%v'\ndestFile='%v\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: CopyFileByIo FAILED! The destination file was NOT created!\n" +
      "destFile='%v'\n", destFile)
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err!=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile).\n" +
      "Attempted deletion of destination file during clean-up FAILED!\n" +
      "destFile='%v'\nError='%v'\n",destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of destination file during clean-up FAILED!\n" +
      "Destination File STILL EXISTS!\n" +
      "Destination File='%v'\n", destFile)
  }

}

func TestFileHelper_CopyFileByIo_11(t *testing.T) {

  fh := FileHelper{}

  destFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_11.txt"

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error retrned by setup fh.DeleteDirFile(destFile).\n" +
        "Attempted deletion of pre-existing destination file FAILED!\n" +
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Attempted deletion of pre-existing destination file FAILED!\n" +
        "destFile='%v'\n", destFile)
      return
    }

  }

  srcFile := "../filesfortest/levelfilesfortest/level_0_2_test.txt"

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by setup fh.CopyFileByIo(srcFile, destFile).\n" +
      "srcFile='%v'\ndestFile='%v\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: Attempted creation of setup destination file FAILED!\n" +
      "destFile='%v'\n", destFile)
    return
  }

  srcFile2 := "../filesfortest/levelfilesfortest/level_0_3_test.txt"

  err = fh.CopyFileByIo(srcFile2, destFile)

  if err != nil {
    t.Errorf("Error returned by 2nd Copy fh.CopyFileByIo(srcFile2, destFile).\n" +
      "srcFile2='%v'\ndestFile='%v\nError='%v'\n",
      srcFile2, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: 2nd Copy of destination file does NOT exist!\n" +
      "Destination File='%v'\n", destFile)
    return
  }

  // 2nd destination file DOES EXIST!

  finfoSrcFile, err := os.Stat(srcFile2)

  if err != nil {
    t.Errorf("Error returned by os.Stat(srcFile2).\n" +
      "srcFile2='%v'\nError='%v'\n", srcFile2, err.Error())
  }

  finfoDestFile, err := os.Stat(destFile)

  if err != nil {
    t.Errorf("Error returned by os.Stat(destFile).\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if finfoSrcFile.Size() != finfoDestFile.Size() {
    t.Errorf("The sizes of the source file and destination file DO NOT MATHCH!\n" +
      "Source File Size='%v'  Destination File Size='%v'.\n",
      finfoSrcFile.Size(), finfoDestFile.Size())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by clean-up fh.DeleteDirFile(destFile).\n" +
      "destFile='%v'\nError='%v' ", destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of destination file FAILED!\n" +
      "Destination File='%v'\n", destFile)
  }

}