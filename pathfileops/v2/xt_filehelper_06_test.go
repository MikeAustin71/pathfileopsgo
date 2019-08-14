package pathfileops

import (
  "strings"
  "testing"
)

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

  commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\" +
    "003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\" +
    "003_filehelper\\common")

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

  commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\" +
    "003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\" +
    "003_filehelper\\common")

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

  expectedDir := fh.AdjustPathSlash("./")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). "+
      "commonDir='%v' Error='%v'", commonDir, err.Error())
    return
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v'\n"+
      "for valid path/file name.\nInstead return path == '%v'\n",
      expectedDir, result)
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

  commonDir := fh.AdjustPathSlash("../../../../")

  expectedDir := fh.AdjustPathSlash("../../../../")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir).\n"+
      "commonDir='%v'\nError='%v'\n",
      commonDir, err.Error())
    return
  }

  if false != isEmpty {
    t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'.\n"+
      "Instead, isEmpty='%v'\n",
      false, isEmpty)
  }

  if result != expectedDir {
    t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file "+
      "name\n"+
      "Instead return path == '%v'\n",
      expectedDir, result)
  }
}

func TestFileHelper_GetPathFromPathFileName_12(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")

  expectedDir := fh.AdjustPathSlash("./")

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

func TestFileHelper_GetPathFromPathFileName_14(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\.git")

  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

  result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir).\n"+
      "commonDir='%v'\nError='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension\n"+
      "to return 'false'. Instead isEmpty='%v'\n", isEmpty)
  }

  if result != expectedDir {
    t.Errorf("ERROR: Expected GetPathFromPathFileName to return "+
      "path == '%v' for valid path/file name.\n"+
      "Instead path == %v\n", expectedDir, result)
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

  isLinux := GlobalPathFileOpsSys{}.IsLinuxOperatingSystem()

  if isLinux {
    testVolStr = "/dev/sda1"
  }

  expectedVolName := strings.ToLower(testVolStr)

  volumeName := fh.GetVolumeName(testVolStr)

  if expectedVolName != strings.ToLower(volumeName) {
    t.Errorf("Expected volumeName='%v'. Instead, volName='%v' ",
      expectedVolName, strings.ToLower(volumeName))
  }

}

func TestFileHelper_GetVolumeNameIndex_01(t *testing.T) {
  fh := FileHelper{}

  testVolStr := ""

  _,
  _,
  volumeName := fh.GetVolumeNameIndex(testVolStr)

  if volumeName != "" {
    t.Errorf("Expected returned volumeName to be an EMPTY STRING.\n" +
      "Instead, volume Name='%v'\n",  volumeName)
  }

}

func TestFileHelper_IsAbsolutePath_01(t *testing.T) {

  fh := FileHelper{}
  commonDir := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/" +
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

func TestFileHelper_IsAbsolutePath_03(t *testing.T) {

  fh := FileHelper{}
  absPathDir := ""

  result := fh.IsAbsolutePath(absPathDir)

  if result == true {
    t.Error("Expected a return value of 'false' from fh.IsAbsolutePath(absPathDir) because\n" +
      "'absPathDir' is an empty string. However, the returned value was 'true'. ERROR!\n")
  }

}

func TestFileHelper_IsPathFileString_01(t *testing.T) {

  fh := FileHelper{}
  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir/level_3_1_test.txt")

  expectedPathFile := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir\\level_3_1_test.txt")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
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
  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir/iDoNotExist.txt")

  expectedPathFile := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir\\iDoNotExist.txt")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
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
  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir")

  expectedPathFile := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
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
  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  expectedPathFile := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\iDoNotExist")

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    t.Errorf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
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
    return
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
    return
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
    return
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
    return
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
    return
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
    return
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
    return
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
    return
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

  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir")

  expectedPathStr := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
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

  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  expectedPathStr := fh.AdjustPathSlash("..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
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

  pathFile := ""

  _, _, _, err := fh.IsPathString(pathFile)

  if err == nil {
    t.Errorf("Expected an error return from fh.IsPathString(pathFile) " +
      "because 'pathFile' is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_IsPathString_05(t *testing.T) {

  fh := FileHelper{}

  pathFile := "      "

  _, _, _, err := fh.IsPathString(pathFile)

  if err == nil {
    t.Errorf("Expected an error return from fh.IsPathString(pathFile) " +
      "because 'pathFile' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_IsPathString_06(t *testing.T) {

  fh := FileHelper{}

  pathFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/level_03_dir")

  pathFile = "." + pathFile

  _, _, _, err := fh.IsPathString(pathFile)

  if err == nil {
    t.Errorf("Expected an error return from fh.IsPathString(pathFile) " +
      "because 'pathFile' includes the text '...' . " +
      "However, NO ERROR WAS RETURNED!")
  }

}
