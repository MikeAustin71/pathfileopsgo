package pathfileops

import (
  "fmt"
  "testing"
  "time"
)


func TestFileHelper_FindFilesWalkDirectory_01(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../../filesfortest/levelfilesfortest/level_01_dir"

  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
    return
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles == 0 {
    t.Error("Error: Expected found files to be greater than zero. " +
      "ZERO FILES FOUND!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_02(t *testing.T) {

  fh := FileHelper{}

  fsc := FileSelectionCriteria{}

  startPath := "../../filesfortest/levelfilesfortest/level_01_dir"
  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
    return
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles == 0 {
    t.Error("Error: Expected found files to be greater than zero. " +
      "ZERO FILES FOUND!")
  }
}

func TestFileHelper_FindFilesWalkDirectory_03(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../../filesfortest/iDoNotExist/childDoesNotExist"

  _, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err == nil {
    t.Error("ERROR: Expected an error returned from fh.FindFilesWalkDirectory(startPath, fsc)\n" +
      "because 'startPath' does not exist on disk.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileHelper_FindFilesWalkDirectory_04(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err := fh.FindFilesWalkDirectory("", fsc)

  if err == nil {
    t.Error("Expected error return from fh.FindFilesWalkDirectory(\"\", fsc) " +
      "because first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_05(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err := fh.FindFilesWalkDirectory("    ", fsc)

  if err == nil {
    t.Error("Expected error return from fh.FindFilesWalkDirectory(\"    \", fsc) " +
      "because first input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_06(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../../testdestdir"
  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
    return
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles != 0 {
    t.Errorf("Expected zero found files. Instead, files were found. Error!"+
      "numOfFiles='%v' ", numOfFiles)
  }

}

func TestFileHelper_FindFilesWalkDirectory_07(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "....////&!////@$...///&*()"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err := fh.FindFilesWalkDirectory("    ", fsc)

  if err == nil {
    t.Error("Expected error return from fh.FindFilesWalkDirectory(\"....////&!////@$...///&*()\", fsc) " +
      "because first input parameter is an invalid directory. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetAbsPathFromFilePath_01(t *testing.T) {

  _, err := FileHelper{}.GetAbsPathFromFilePath("")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}." +
      "GetAbsPathFromFilePath(\"\") because the input parameter is an " +
      "empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetAbsPathFromFilePath_02(t *testing.T) {

  _, err := FileHelper{}.GetAbsPathFromFilePath("    ")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}." +
      "GetAbsPathFromFilePath(\"     \") because the input parameter consists " +
      "entirely of blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetDotSeparatorIndexesInPathStr_01(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetDotSeparatorIndexesInPathStr("")

  if err == nil {
    t.Error("Expected an error return from fh." +
      "GetDotSeparatorIndexesInPathStr(\"\") because the input parameter is an " +
      "empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetDotSeparatorIndexesInPathStr_02(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetDotSeparatorIndexesInPathStr("   ")

  if err == nil {
    t.Error("Expected an error return from fh." +
      "GetDotSeparatorIndexesInPathStr(\"  \") because the input parameter consists " +
      "entirely of blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetExecutablePathFileName_01(t *testing.T) {

  fh := FileHelper{}

  exePathFileName, err := fh.GetExecutablePathFileName()

  if err != nil {
    t.Errorf("Error returned by fh.GetExecutablePathFileName(). "+
      "Error='%v' ", err.Error())
  }

  if len(exePathFileName) == 0 {
    t.Error("Error: The Executable path an file name returned by " +
      "fh.GetExecutablePathFileName() is a zero length string!")
  }
}

func TestFileHelper_GetFileExtension_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '.go' for valid file extension Instead result='%v' ", result)
  }

}

func TestFileHelper_GetFileExtension_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'.  "+
      "Instead file extension='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("dirmgr_test")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead file extension ='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead, file extension='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
      "Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".gitignore")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). "+
      "commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
      "Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("test.....go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_09(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err == nil {
    t.Errorf("Expected an error to be returned from fh.GetFileExt(commonDir). "+
      "commonDir='%v' NO Error was returned!", commonDir)
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_10(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\dirmgr_test")

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty GetFileExt for absent file extension to return 'true'. Instead, isEmpty='%v' ", isEmpty)
  }

  if result != "" {
    t.Errorf("Expected GetFileExt to return empty result for absent file extension. Instead file extension='%v' ", result)
  }

}

func TestFileHelper_GetFileExtension_11(t *testing.T) {
  fh := FileHelper{}

  result, isEmpty, err := fh.GetFileExtension("   ")

  if err == nil {
    t.Error("Expected an error to be returned from fh.GetFileExt(\"    \") because " +
      "the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != "" {
    t.Errorf("Expected GetFileExt to return and empty string. Instead, result='%v' ",
      result)
  }
}

func TestFileHelper_GetFileInfoFromPath_01(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetFileInfo("")

  if err == nil {
    t.Error("Expected error from fh.GetFileInfo(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileInfoFromPath_02(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetFileInfo("    ")

  if err == nil {
    t.Error("Expected error from fh.GetFileInfo(\"   \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileInfoFromPath_03(t *testing.T) {
  fh := FileHelper{}

  rawPath := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

  _, err := fh.GetFileInfo(rawPath)

  if err == nil {
    t.Error("Expected error from fh.GetFileInfo(rawPath) " +
      "because the input parameter 'rawPath' does NOT exist. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_01(t *testing.T) {

  fh := FileHelper{}
  target, err := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogtopTest2Text))

  if err != nil {
    t.Error("Error from FileHelper:MakeAbsolutePath(): ", err.Error())
  }

  tStrFmt := "2006-01-02 15:04:05.000000000"

  fileTime, tStr, err := fh.GetFileLastModificationDate(target, tStrFmt)

  if err != nil {
    t.Errorf("Error returned by FileHelper:GetFileLastModificationDate().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  fInfo, err := fh.GetFileInfo(target)

  if err != nil {
    t.Errorf("Error returned by FileHelper:GetFileInfo().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFileTime := fInfo.ModTime()

  expected := actualFileTime.Format(tStrFmt)

  if tStr != expected {
    t.Errorf("Expected Time String for file %v == %v.\n" +
      "Instead, received time string= %v", target, expected, tStr)
  }

  if !actualFileTime.Equal(fileTime) {
    t.Errorf("Expected Time Value='%v'. Instead, Time Value='%v'\n",
      actualFileTime, fileTime)
  }
}

func TestFileHelper_GetFileLastModificationDate_02(t *testing.T) {

  fh := FileHelper{}
  tStrFmt := "2006-01-02 15:04:05.000000000"

  _, _, err := fh.GetFileLastModificationDate("", tStrFmt)

  if err == nil {
    t.Error("Expected error return from fh.GetFileLastModificationDate" +
      "(\"\", tStrFmt) because the first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_03(t *testing.T) {

  fh := FileHelper{}
  tStrFmt := "2006-01-02 15:04:05.000000000"

  _, _, err := fh.GetFileLastModificationDate("   ", tStrFmt)

  if err == nil {
    t.Error("Expected error return from fh.GetFileLastModificationDate" +
      "(\"    \", tStrFmt) because the first input parameter consists entirely of " +
      "blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_04(t *testing.T) {

  fh := FileHelper{}
  target, err := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogtopTest2Text))

  if err != nil {
    t.Error("Error from FileHelper:MakeAbsolutePath(): ", err.Error())
  }

  tStrFmt := ""

  fileTime, tStr, err := fh.GetFileLastModificationDate(target, tStrFmt)

  if err != nil {
    t.Error("Error from FileHelper:GetFileLastModificationDate():", err.Error())
  }

  fInfo, err := fh.GetFileInfo(target)

  if err != nil {
    t.Errorf("Error returned by FileHelper:GetFileInfo()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFileTime := fInfo.ModTime()

  expected := actualFileTime.Format("2006-01-02 15:04:05.000000000 -0700 MST")

  if tStr != expected {
    t.Errorf("Expected Time String for file %v == %v, received time string: %v", target, expected, tStr)
  }

  if !actualFileTime.Equal(fileTime) {
    t.Error(fmt.Sprintf("Expected Time value %v, instead got:", actualFileTime), fileTime)
  }
}

func TestFileHelper_GetFileLastModificationDate_05(t *testing.T) {

  fh := FileHelper{}
  tStrFmt := "2006-01-02 15:04:05.000000000"
  targetPath :=
    "../../createFilesTest/Level01/Level02/TestFileHelper_GetFileLastModificationDate_05.txt"

  _, _, err := fh.GetFileLastModificationDate(targetPath, tStrFmt)

  if err == nil {
    t.Error("Expected error return from fh.GetFileLastModificationDate" +
      "(targetPath, tStrFmt) because 'targetPath' does NOT exist!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}


func TestFileHelper_GetFileMode_01(t *testing.T) {

  pathFileName := ""

  fh := FileHelper{}

  _, err := fh.GetFileMode(pathFileName)

  if err == nil {
    t.Error("Expected an error return from fh.GetFileMode(pathFileName) because\n" +
      "parameter 'pathFileName' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetFileMode_02(t *testing.T) {

  pathFileName := "       "

  fh := FileHelper{}

  _, err := fh.GetFileMode(pathFileName)

  if err == nil {
    t.Error("Expected an error return from fh.GetFileMode(pathFileName) because\n" +
      "parameter 'pathFileName' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetFileMode_03(t *testing.T) {

  pathFileName := "../../createFilesTest/iDoNOTExist.txt"

  fh := FileHelper{}

  pathFileName = fh.AdjustPathSlash(pathFileName)

  _, err := fh.GetFileMode(pathFileName)

  if err == nil {
    t.Error("Expected an error return from fh.GetFileMode(pathFileName) because\n" +
      "parameter 'pathFileName' does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetFileMode_04(t *testing.T) {

  pathFileName := "../../filesfortest/levelfilesfortest/level_0_3_test.txt"

  fh := FileHelper{}

  pathFileName = fh.AdjustPathSlash(pathFileName)

  filePermission, err := fh.GetFileMode(pathFileName)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileMode(pathFileName)\n" +
      "pathFileName='%v'\nError='%v'", pathFileName, err.Error())
    return
  }

  err = filePermission.IsValid()

  if err != nil {
    t.Errorf("Error returned by filePermission.IsValid().\n" +
      "There is something wrong with FilePermissionCfg object retunred by\n" +
      "fh.GetFileMode(pathFileName).\npathFileName='%v'\nError='%v'",
      pathFileName, err.Error())
  }

  return
}

func TestFileHelper_GetFileNameWithExt_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
      false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\")
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
      true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".gitignore")
  expectedFNameExt := ".gitignore"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). "+
      "commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_09(t *testing.T) {
  fh := FileHelper{}

  commonDir := ""
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err == nil {
    t.Errorf("Expected error return from fh.GetFileNameWithExt(commonDir) " +
      "because 'commonDir' is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_10(t *testing.T) {
  fh := FileHelper{}

  commonDir := "   "
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err == nil {
    t.Errorf("Expected error return from fh.GetFileNameWithExt(commonDir) " +
      "because 'commonDir' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }
}

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

  rawPath := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(rawPath)

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  if firstCharIdx != 6 {
    t.Errorf("Expected first char index= '6'.  Instead, first char index= '%v'", firstCharIdx)
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
