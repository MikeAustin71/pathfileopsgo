package pathfileops

import (
  "fmt"
  "strings"
  "testing"
)

func TestFileMgr_GetAbsolutePath(t *testing.T) {

  fh := FileHelper{}
  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  fMgr1DMgr := fileMgr1.GetDirMgr()

  expectedAbsPath := fMgr1DMgr.GetAbsolutePath()

  actualAbsPath := fileMgr1.GetAbsolutePath()

  if expectedAbsPath != actualAbsPath {
    t.Errorf("Expected absolute path='%v'. Instead, absolute path='%v' ",
      expectedAbsPath, actualAbsPath)
  }

}

func TestFileMgr_GetAbsolutePathFileName_01(t *testing.T) {
  fh := FileHelper{}
  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  absPathFileName := fileMgr1.GetAbsolutePathFileName()
  absPathFileName = strings.ToLower(absPathFileName)
  filePath1 = strings.ToLower(filePath1)

  if filePath1 != absPathFileName {
    t.Errorf("Error: Expected absPathFileName='%v'. Instead, absPathFileName='%v' ",
      filePath1, absPathFileName)
  }

}

func TestFileMgr_GetBufioReader_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
      "dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly(). "+
      "FileName:'%v' Error='%v' ", srcFMgr.GetAbsolutePathFileName(), err.Error())
  }

  bufReader := srcFMgr.GetBufioReader()

  if bufReader != nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioReader() to be 'NIL'. " +
      "However, POINTER IS POPULATED!")
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CloseThisFile(). "+
      "Error='%v' ", err.Error())
  }
}

func TestFileMgr_GetBufioWriter_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
      "dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  err = srcFMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by srcFMgrOpenThisFileWriteOnlyAppend(). "+
      "FileName:'%v' Error='%v' ", srcFMgr.GetAbsolutePathFileName(), err.Error())
  }

  bufReader := srcFMgr.GetBufioWriter()

  if bufReader != nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioWriter() to " +
      "to be 'NIL'.  However, POINTER IS POPULATED!")
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CloseThisFile(). "+
      "Error='%v' ", err.Error())
  }

}

func TestFileMgr_GetDirMgr_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
      "dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  expectedDirMgrPath := strings.ToLower(dMgr.GetAbsolutePath())

  srcDMgr := srcFMgr.GetDirMgr()

  actualDirMgrPath := strings.ToLower(srcDMgr.GetAbsolutePath())

  if expectedDirMgrPath != actualDirMgrPath {
    t.Errorf("Error: Expected returned directory path='%v'. Instead, "+
      "returned directory path='%v' ",
      expectedDirMgrPath, actualDirMgrPath)
  }

}

func TestFileMgr_GetFileExt(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
      "dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  expectedFileExt := ".txt"

  actualFileExt := srcFMgr.GetFileExt()

  if expectedFileExt != actualFileExt {
    t.Errorf("Error: Expected returned file extension='%v'. Instead "+
      "returned file extension='%v' ",
      expectedFileExt, actualFileExt)
  }

}

func TestFileMgr_GetFileInfo_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  fInfo, err := srcFMgr.GetFileInfo()

  expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

  actualFileNameExt := strings.ToLower(fInfo.Name())

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name:='%v'.  Instead, File Name='%v'",
      expectedFileNameExt, actualFileNameExt)
  }
}

func TestFileMgr_GetFileInfo_02(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  _, err = srcFMgr.GetFileInfo()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfo() because " +
      "file does not exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFileInfo_03(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileInfo()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfo() because " +
      "srcFMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFileInfoPlus_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  fInfoPlus, err := srcFMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileInfoPlus(). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

  actualFileNameExt := strings.ToLower(fInfoPlus.Name())

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name:='%v'.  Instead, File Name='%v'",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFileInfoPlus_02(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  _, err = srcFMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfoPlus() because " +
      "file does not exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_GetFileInfoPlus_03(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfoPlus() because " +
      "srcFMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFileModTime_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime(). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  expectedTimeStr := modTime.Format(timeFormatSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr("")

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\"). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'. Instead, Time String='%v'. ",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime(). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  timeFmtSpec := "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

  expectedTimeStr := modTime.Format(timeFmtSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr(timeFmtSpec)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\"). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'. Instead, Time String='%v'. ",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime(). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  expectedTimeStr := modTime.Format(timeFormatSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr("xx-xx-xxxx xx:xx:xx")

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\"). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'. Instead, Time String='%v'. ",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_04(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileModTimeStr("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFileModTimeStr(\"\") " +
      "because the target file DOES NOT EXIST! However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFileModTime_05(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/iDoNotExist.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  _, err = srcFMgr.GetFileModTimeStr("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFileModTimeStr(\"\") " +
      "because the target file DOES NOT EXIST! However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFileName_01(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  expectedFileName := "newerFileForTest_01"

  actualFileName := srcFMgr.GetFileName()

  if expectedFileName != actualFileName {
    t.Errorf("Error: Expected File Name='%v'. Instead, actual File Name='%v'",
      expectedFileName, actualFileName)
  }

}

func TestFileMgr_GetFileNameExt_01(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath). "+
      "absPath='%v' Error='%v'", absPath, err.Error())
  }

  expectedFileNameExt := "newerFileForTest_01.txt"

  actualFileNameExt := srcFMgr.GetFileNameExt()

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name Ext='%v'. Instead, actual File Name Ext='%v'",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFilePermissionConfig_01(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.OpenThisFileReadWrite()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadWrite(). "+
      "Error='%v' ", err.Error())

  }

  expectedPermissionCodes := "-rw-rw-rw-"

  actualPermissionTextCfg, err := srcFMgr.GetFilePermissionConfig()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.GetFilePermissionConfig(). Error='%v'",
      err.Error())
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 srcFMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
  }

  actualPermissionTextCodes, err := actualPermissionTextCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned from actualPermissionTextCfg.GetPermissionTextCode(). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  if expectedPermissionCodes != actualPermissionTextCodes {
    t.Errorf("Error: Expected Permission Code='%v'. Instead, Permission Code='%v'",
      expectedPermissionCodes, actualPermissionTextCodes)
  }

}

func TestFileMgr_GetFilePermissionConfig_02(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDoNotExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  _, err = srcFMgr.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected an error return from srcFMgr.GetFilePermissionConfig() " +
      "because the file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFilePermissionConfig_03(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected an error return from srcFMgr.GetFilePermissionConfig() " +
      "because the file manager (srcFmgr) is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFilePermissionTextCodes_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.OpenThisFileReadWrite()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadWrite(). "+
      "Error='%v' ", err.Error())

  }

  expectedPermissionCodes := "-rw-rw-rw-"

  actualPermissionTextCodes, err := srcFMgr.GetFilePermissionTextCodes()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.GetFilePermissionTextCodes(). Error='%v'",
      err.Error())
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 srcFMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
  }

  if expectedPermissionCodes != actualPermissionTextCodes {
    t.Errorf("Error: Expected Permission Code='%v'. Instead, Permission Code='%v'",
      expectedPermissionCodes, actualPermissionTextCodes)
  }

}

func TestFileMgr_GetFilePermissionTextCodes_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDoNotExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  _, err = srcFMgr.GetFilePermissionTextCodes()

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFilePermissionTextCodes() " +
      "because file does not exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFilePermissionTextCodes_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFilePermissionTextCodes()

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFilePermissionTextCodes() " +
      "because srcFMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_GetFilePtr_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly(). "+
      "Error='%v' ", err.Error())

  }

  fPtr := srcFMgr.GetFilePtr()

  if fPtr == nil {
    t.Error("Error: Expected a populated file pointer. However, the file pointer is nil!")
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by final srcFMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

}

func TestFileMgr_GetFileSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(29)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='29'. Instead, file size='%v'",
      actualFileSize)
  }

}

func TestFileMgr_GetFileSize_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDontExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(-1)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='-1'. Instead, file size='%v'",
      actualFileSize)
  }

}

func TestFileMgr_GetFileSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr.isInitialized = false

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(-1)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='%v'. Instead, file size='%v'",
      expectedFileSize, actualFileSize)
  }

}

func TestFileMgr_GetOriginalPathFileName_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  originalPathFileName := srcFMgr.GetOriginalPathFileName()

  if targetFile != originalPathFileName {
    t.Errorf("Error: Expected Original Path and File Name='%v'. Instead, "+
      "Original Path and File Name='%v'",
      targetFile, originalPathFileName)
  }

}

func TestFileMgr_GetReaderBufferSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  readBufSize := srcFMgr.GetReaderBufferSize()

  if readBufSize != 0 {
    t.Errorf("Error: Expected Bufio Read Buffer Size='0'. Instead, "+
      "Read Buffer Size ='%v' ", readBufSize)
  }

}

func TestFileMgr_GetReaderBufferSize_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  expectedBufSize := int(20000)

  srcFMgr.SetReaderBufferSize(expectedBufSize)

  readBufSize := srcFMgr.GetReaderBufferSize()

  if expectedBufSize != readBufSize {
    t.Errorf("Error: Expected Bufio Read Buffer Size='%v'. Instead, "+
      "Read Buffer Size ='%v' ", expectedBufSize, readBufSize)
  }

}

func TestFileMgr_GetReaderBufferSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly(). Error='%v'",
      err.Error())
    return
  }

  bytes := make([]byte, 10000, 20000)

  _, err = srcFMgr.ReadFileBytes(bytes)

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.ReadFileBytes(bytes). Error='%v'",
      err.Error())
    return
  }

  readBufSize := srcFMgr.GetReaderBufferSize()

  _ = srcFMgr.CloseThisFile()

  if readBufSize < 10 {
    t.Errorf("Error: Expected Bufio Read Buffer Size >10. Instead, "+
      "Read Buffer Size ='%v' ", readBufSize)
  }

}

func TestFileMgr_GetWriterBufferSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  if actualWriterBuffSize != 0 {
    t.Errorf("Error: Expected Bufio Write Buffer Size='0'. Instead, "+
      "Write Buffer Size ='%v' ", actualWriterBuffSize)
  }

}

func TestFileMgr_GetWriterBufferSize_02(t *testing.T) {

  expectedWriteBufSize := int(20000)

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr.SetWriterBufferSize(expectedWriteBufSize)

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  if actualWriterBuffSize != expectedWriteBufSize {
    t.Errorf("Error: Expected Bufio Write Buffer Size='%v'. Instead, "+
      "Write Buffer Size ='%v' ", expectedWriteBufSize, actualWriterBuffSize)
  }

}

func TestFileMgr_GetWriterBufferSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/scratchWriteFileTest_5096.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DeleteThisFile(). Error='%v' ",
      err.Error())
  }

  err = srcFMgr.CreateThisFile()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.CreateThisFile(). Error='%v' ",
      err.Error())
    return
  }

  err = srcFMgr.OpenThisFileWriteOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
    return
  }

  _, err = srcFMgr.WriteStrToFile("Hello World!")

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.WriteStrToFile(). Error='%v' ",
      err.Error())
    return
  }

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

  err = srcFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DeleteThisFile(). Error='%v' ",
      err.Error())
  }

  if actualWriterBuffSize < 10 {
    t.Errorf("Error: Expected Bufio Write Buffer Size > 10. Instead, "+
      "Write Buffer Size ='%v' ", actualWriterBuffSize)
  }

}

func TestFileMgr_IsAbsolutePathFileNamePopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

  if !isAbsPathFileName {
    t.Error("Error: Expected Absolute Path File Name to be populated. " +
      "It was NOT!")
  }

}

func TestFileMgr_IsAbsolutePathFileNamePopulated_02(t *testing.T) {

  srcFMgr := FileMgr{}

  isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

  if isAbsPathFileName {
    t.Error("Error: Expected Absolute Path File Name NOT populated. " +
      "WRONG - It IS populated!")
  }

}

func TestFileMgr_IsFileExtPopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isFileExtPopulated := srcFMgr.IsFileExtPopulated()

  if !isFileExtPopulated {
    t.Error("Expected srcFMgr.IsFileExtPopulated() == 'true'. Instead, it is 'false'")
  }

}

func TestFileMgr_IsFileExtPopulated_02(t *testing.T) {

  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
      "targetDir='%v' Error='%v'", targetDir, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
      "DirMgr='%v' targetFile='%v' Error='%v'",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
  }

  isFileExtPopulated := srcFMgr.IsFileExtPopulated()

  if isFileExtPopulated {
    t.Error("Expected srcFMgr.IsFileExtPopulated() == 'false'. Instead, it is 'true'")
  }

}

func TestFileMgr_IsFileNameExtPopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isFileNameExtPopulated := srcFMgr.IsFileNameExtPopulated()

  if !isFileNameExtPopulated {
    t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'true'. Instead, it is 'false'")
  }

}

func TestFileMgr_IsFileNameExtPopulated_02(t *testing.T) {
  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
      "targetDir='%v' Error='%v'", targetDir, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
      "DirMgr='%v' targetFile='%v' Error='%v'",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
  }

  isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

  if isFileNamePopulated {

    t.Errorf("Expected srcFMgr.IsFileNameExtPopulated() == 'false'. Instead, it is 'true'. "+
      "FileName='%v' File Extension='%v' Len File Ext= '%v' ",
      srcFMgr.GetFileName(), srcFMgr.GetFileExt(), len(srcFMgr.GetFileExt()))
  }

}

func TestFileMgr_IsFileNameExtPopulated_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

  if isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'false'. Instead, it is 'true'")
  }

}

func TestFileMgr_IsFileNamePopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if !isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'. Instead, it is 'false'")
  }
}

func TestFileMgr_IsFileNamePopulated_02(t *testing.T) {

  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
      "targetDir='%v' Error='%v'", targetDir, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
      "DirMgr='%v' targetFile='%v' Error='%v'",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
  }

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if !isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'. Instead, it is 'false'")
  }

}

func TestFileMgr_IsFileNamePopulated_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'false'. Instead, it is 'true'")
  }

}

func TestFileMgr_IsFilePointerOpen_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned from srcFMgr.OpenThisFileReadOnly(). "+
      "Error='%v'", err.Error())
  }

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from final srcFMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
  }

  if !isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'true'. Instead, it is FALSE!")
  }

}

func TestFileMgr_IsFilePointerOpen_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  if isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'false'. Instead, it is TRUE!")
  }

}

func TestFileMgr_IsFilePointerOpen_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  if isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'false'. Instead, it is TRUE!")
  }

}

func TestFileMgr_IsInitialized_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  isInitialized := srcFMgr.IsInitialized()

  if !isInitialized {
    t.Error("Expected isInitialized = 'true'. Instead, it is FALSE!")
  }
}

func TestFileMgr_IsInitialized_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'", targetFile, err.Error())
  }

  srcFMgr.Empty()

  isInitialized := srcFMgr.IsInitialized()

  if isInitialized {
    t.Error("Expected isInitialized = 'false'. Instead, it is TRUE!")
  }
}

func TestFileMgr_IsInitialized_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isInitialized := srcFMgr.IsInitialized()

  if isInitialized {
    t.Error("Expected isInitialized = 'false'. Instead, it is TRUE!")
  }
}

func TestFileMgr_MoveFileToNewDirMgr_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). destDir='%v' Error='%v'", destDir, err.Error())
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr). dMgr.path='%v'  Error='%v'", dMgr.path, err.Error())
  }

  if !fh.DoesFileExist(newFMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
  }

  if !doesExist {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  err = newFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

}

func TestFileMgr_MoveFileToNewDirMgr_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). "+
      "destDir='%v' Error='%v'", destDir, err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "TestFile003.txt")

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, \"TestFile003.txt\"). "+
      "Error='%v'", err.Error())
  }

  srcFileMgr2 := srcFileMgr.CopyOut()

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because srcFileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = newFMgr.DeleteThisFile()

  if fh.DoesFileExist(newFMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: Deletion of New File failed. File='%v'",
      newFMgr.GetAbsolutePathFileName())
  }

  err = srcFileMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of source file FAILED!. "+
      "Source File ='%v'", srcFileMgr2.GetAbsolutePathFileName())
  }

}

func TestFileMgr_MoveFileToNewDirMgr_03(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(DirMgr{})

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because dMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDirMgr_04(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\iDoNotExist.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  destDir := fh.AdjustPathSlash("..\\logTest")

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "directory='%v'  Error='%v'", destDir, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDirMgr_05(t *testing.T) {

  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  destDir := fh.AdjustPathSlash("..\\xxxIDoNotExist")

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "directory='%v'  Error='%v'", destDir, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because dMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
  }

  if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
  }

  if !doesExist {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  err = newFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

}

func TestFileMgr_MoveFileToNewDir_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  srcFileMgr2 := srcFileMgr.CopyOut()

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because srcFileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  err = srcFileMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFileMgr2.DeleteThisFile(). "+
      "srcFile2='%v'  Error='%v'",
      srcFileMgr2.GetAbsolutePathFileName(), err.Error())
  }

  if fh.DoesFileExist(srcFileMgr2.GetAbsolutePathFileName()) {
    t.Errorf("ERROR: Deletion of test file failed!. Test File='%v' ",
      srcFileMgr2.GetAbsolutePathFileName())
  }

}

func TestFileMgr_MoveFileToNewDir_03(t *testing.T) {
  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir("")

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because destDir is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_04(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\iDoNotExist.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_05(t *testing.T) {
  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\iDoNotExit")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because destDir does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}
