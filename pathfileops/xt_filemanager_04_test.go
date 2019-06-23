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
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  fMgr1DMgr := fileMgr1.GetDirMgr()

  expectedAbsPath := fMgr1DMgr.GetAbsolutePath()

  actualAbsPath := fileMgr1.GetAbsolutePath()

  if expectedAbsPath != actualAbsPath {
    t.Errorf("Expected absolute path='%v'.\n"+
      "Instead, absolute path='%v'\n",
      expectedAbsPath, actualAbsPath)
  }

}

func TestFileMgr_GetAbsolutePathFileName_01(t *testing.T) {
  fh := FileHelper{}
  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  absPathFileName := fileMgr1.GetAbsolutePathFileName()
  absPathFileName = strings.ToLower(absPathFileName)
  filePath1 = strings.ToLower(filePath1)

  if filePath1 != absPathFileName {
    t.Errorf("Error: Expected absPathFileName='%v'.\n"+
      "Instead, absPathFileName='%v'\n",
      filePath1, absPathFileName)
  }

}

func TestFileMgr_GetBufioReader_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly().\n"+
      "FileName:'%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  bufReader := srcFMgr.GetBufioReader()

  if bufReader == nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioReader() to be populated " +
      "and NOT 'nil'.\n" +
      "However, POINTER IS 'nil' and NOT populated!\n")
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CloseThisFile().\n"+
      "Error='%v'\n", err.Error())
  }
}

func TestFileMgr_GetBufioWriter_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by srcFMgrOpenThisFileWriteOnlyAppend().\n"+
      "FileName:'%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  bufReader := srcFMgr.GetBufioWriter()

  if bufReader == nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioWriter()\n" +
      "to be poulated and NOT 'nil.\nHowever, POINTER IS NIL and NOT POPULATED!\n")
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CloseThisFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_GetDirMgr_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  expectedDirMgrPath := strings.ToLower(dMgr.GetAbsolutePath())

  srcDMgr := srcFMgr.GetDirMgr()

  actualDirMgrPath := strings.ToLower(srcDMgr.GetAbsolutePath())

  if expectedDirMgrPath != actualDirMgrPath {
    t.Errorf("Error: Expected returned directory path='%v'.\n"+
      "Instead, returned directory path='%v'\n",
      expectedDirMgrPath, actualDirMgrPath)
  }

}

func TestFileMgr_GetFileExt(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  expectedFileExt := ".txt"

  actualFileExt := srcFMgr.GetFileExt()

  if expectedFileExt != actualFileExt {
    t.Errorf("Error: Expected returned file extension='%v'.\n"+
      "Instead returned file extension='%v'\n",
      expectedFileExt, actualFileExt)
  }

}

func TestFileMgr_GetFileInfo_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  fInfo, err := srcFMgr.GetFileInfo()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileInfo()\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

  actualFileNameExt := strings.ToLower(fInfo.Name())

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name:='%v'.\n"+
      "Instead, File Name='%v'\n",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFileInfo_02(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  _, err = srcFMgr.GetFileInfo()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfo()\n" +
      "because file does not exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFileInfo_03(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileInfo()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfo()\n" +
      "because srcFMgr is invalid.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFileInfoPlus_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  fInfoPlus, err := srcFMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileInfoPlus().\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

  actualFileNameExt := strings.ToLower(fInfoPlus.Name())

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name:='%v'.\nInstead, File Name='%v'\n",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFileInfoPlus_02(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  _, err = srcFMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfoPlus()\n" +
      "because file does not exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_GetFileInfoPlus_03(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("Error expected error return from srcFMgr.GetFileInfoPlus()\n" +
      "because srcFMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFileModTime_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime().\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  expectedTimeStr := modTime.Format(timeFormatSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr("")

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\").\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'.\n"+
      "Instead, Time String='%v'.\n",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime().\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  timeFmtSpec := "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

  expectedTimeStr := modTime.Format(timeFmtSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr(timeFmtSpec)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\").\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'.\n"+
      "Instead, Time String='%v'.\n",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTime().\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  expectedTimeStr := modTime.Format(timeFormatSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr("xx-xx-xxxx xx:xx:xx")

  if err != nil {
    t.Errorf("Error returned from srcFMgr.GetFileModTimeStr(\"\").\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  if expectedTimeStr != modTimeStr {
    t.Errorf("Expected Time String='%v'.\n"+
      "Instead, Time String='%v'.\n",
      expectedTimeStr, modTimeStr)
  }

}

func TestFileMgr_GetFileModTime_04(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFileModTimeStr("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFileModTimeStr(\"\")\n" +
      "because the target file DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFileModTime_05(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/levelfilesfortest/level_01_dir/iDoNotExist.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  _, err = srcFMgr.GetFileModTimeStr("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFileModTimeStr(\"\")\n" +
      "because the target file DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFileName_01(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  expectedFileName := "newerFileForTest_01"

  actualFileName := srcFMgr.GetFileName()

  if expectedFileName != actualFileName {
    t.Errorf("Error: Expected File Name='%v'.\n"+
      "Instead, actual File Name='%v'\n",
      expectedFileName, actualFileName)
  }

}

func TestFileMgr_GetFileNameExt_01(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  absPath, err := fh.MakeAbsolutePath(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absPath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  expectedFileNameExt := "newerFileForTest_01.txt"

  actualFileNameExt := srcFMgr.GetFileNameExt()

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name Ext='%v'.\n"+
      "Instead, actual File Name Ext='%v'\n",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFileNameExt_02(t *testing.T) {
  fh := FileHelper{}

  expectedFileNameExt := "basefilenoext"

  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    fmt.Printf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirStrFileNameStr(absPath, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  actualFileNameExt := srcFMgr.GetFileNameExt()

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("Error: Expected File Name Ext='%v'.\n"+
      "Instead, actual File Name Ext='%v'\n",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgr_GetFilePermissionConfig_01(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadWrite()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadWrite().\n"+
      "srcFMgr='%v'\nError='%v' ",
      srcFMgr.GetAbsolutePathFileName(), err.Error())

    return
  }

  expectedPermissionCodes := "-rw-rw-rw-"

  actualPermissionTextCfg, err := srcFMgr.GetFilePermissionConfig()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.GetFilePermissionConfig().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())

    return
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 srcFMgr.CloseThisFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  actualPermissionTextCodes, err := actualPermissionTextCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned from actualPermissionTextCfg.GetPermissionTextCode().\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  if expectedPermissionCodes != actualPermissionTextCodes {
    t.Errorf("Error: Expected Permission Code='%v'.\n"+
      "Instead, Permission Code='%v'\n",
      expectedPermissionCodes, actualPermissionTextCodes)
  }

}

func TestFileMgr_GetFilePermissionConfig_02(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDoNotExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  _, err = srcFMgr.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected an error return from srcFMgr.GetFilePermissionConfig()\n" +
      "because the file does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFilePermissionConfig_03(t *testing.T) {
  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected an error return from srcFMgr.GetFilePermissionConfig()\n" +
      "because the file manager (srcFMgr) is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFilePermissionTextCodes_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadWrite()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadWrite().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())

    return
  }

  expectedPermissionCodes := "-rw-rw-rw-"

  actualPermissionTextCodes, err := srcFMgr.GetFilePermissionTextCodes()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.GetFilePermissionTextCodes().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())

    return
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 srcFMgr.CloseThisFile(). "+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())

    return
  }

  if expectedPermissionCodes != actualPermissionTextCodes {
    t.Errorf("Error: Expected Permission Code='%v'.\n"+
      "Instead, Permission Code='%v'\n",
      expectedPermissionCodes, actualPermissionTextCodes)
  }
}

func TestFileMgr_GetFilePermissionTextCodes_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDoNotExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  _, err = srcFMgr.GetFilePermissionTextCodes()

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFilePermissionTextCodes()\n" +
      "because file does not exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFilePermissionTextCodes_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  _, err = srcFMgr.GetFilePermissionTextCodes()

  if err == nil {
    t.Error("Expected error return from srcFMgr.GetFilePermissionTextCodes()\n" +
      "because srcFMgr is invalid.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_GetFilePtr_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {

    _ = srcFMgr.CloseThisFile()

    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  fPtr := srcFMgr.GetFilePtr()

  if fPtr == nil {
    t.Errorf("Error: Expected a populated file pointer.\n"+
      "However, the file pointer is nil!\n"+
      "srcFMgr='%v'\n", srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by final srcFMgr.CloseThisFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_GetFileSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(29)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='29'.\nInstead, file size='%v'\n"+
      "File='%v'",
      actualFileSize, srcFMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_GetFileSize_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDontExist_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(-1)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='-1'.\nInstead, file size='%v'\n"+
      "File='%v'\n",
      actualFileSize, srcFMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_GetFileSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.isInitialized = false

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(-1)

  if expectedFileSize != actualFileSize {
    t.Errorf("Expected file size='%v'.\nInstead, file size='%v'\n"+
      "File='%v'\n",
      expectedFileSize, actualFileSize, srcFMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_GetOriginalPathFileName_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  originalPathFileName := srcFMgr.GetOriginalPathFileName()

  if targetFile != originalPathFileName {
    t.Errorf("Error: Expected Original Path and File Name='%v'.\n"+
      "Instead, Original Path and File Name='%v'\n",
      targetFile, originalPathFileName)
  }

}

func TestFileMgr_GetReaderBufferSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  readBufSize := srcFMgr.GetReaderBufferSize()

  if readBufSize != 0 {
    t.Errorf("Error: Expected Bufio Read Buffer Size='0'.\n"+
      "Instead, Read Buffer Size ='%v'\n", readBufSize)
  }

}

func TestFileMgr_GetReaderBufferSize_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  expectedBufSize := 20000 // int

  srcFMgr.SetReaderBufferSize(expectedBufSize)

  readBufSize := srcFMgr.GetReaderBufferSize()

  if expectedBufSize != readBufSize {
    t.Errorf("Error: Expected Bufio Read Buffer Size='%v'.\n"+
      "Instead, Read Buffer Size ='%v'\n",
      expectedBufSize, readBufSize)
  }

}

func TestFileMgr_GetReaderBufferSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly().\n"+
      "srcFMgr='%v'\nError='%v'",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  bytes := make([]byte, 10000, 20000)

  _, err = srcFMgr.ReadFileBytes(bytes)

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.ReadFileBytes(bytes)."+
      "srcFMgr='%v'\nError='%v'",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  readBufSize := srcFMgr.GetReaderBufferSize()

  _ = srcFMgr.CloseThisFile()

  if readBufSize < 10 {
    t.Errorf("Error: Expected Bufio Read Buffer Size >10.\n"+
      "Instead, Read Buffer Size ='%v' ", readBufSize)
  }

}

func TestFileMgr_GetWriterBufferSize_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  if actualWriterBuffSize != 0 {
    t.Errorf("Error: Expected Bufio Write Buffer Size='0'.\n"+
      "Instead, Write Buffer Size ='%v'\n", actualWriterBuffSize)
    return
  }

}

func TestFileMgr_GetWriterBufferSize_02(t *testing.T) {

  expectedWriteBufSize := 20000 // int

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.SetWriterBufferSize(expectedWriteBufSize)

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  if actualWriterBuffSize != expectedWriteBufSize {
    t.Errorf("Error: Expected Bufio Write Buffer Size='%v'.\n"+
      "Instead, Write Buffer Size ='%v'\n",
      expectedWriteBufSize, actualWriterBuffSize)
  }

}

func TestFileMgr_GetWriterBufferSize_03(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../checkfiles/TestFileMgr_GetWriterBufferSize_03.txt")

  err := fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\n", targetFile)
    return
  }

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.CreateDirAndFile()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.CreateDirAndFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = srcFMgr.OpenThisFileWriteOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.OpenThisFileWriteOnly().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  _, err = srcFMgr.WriteStrToFile("Hello World!")

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned by srcFMgr.WriteStrToFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  actualWriterBuffSize := srcFMgr.GetWriterBufferSize()

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CloseThisFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    _ = srcFMgr.DeleteThisFile()
    return
  }

  err = srcFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DeleteThisFile().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if actualWriterBuffSize < 10 {
    t.Errorf("Error: Expected Bufio Write Buffer Size > 10.\n"+
      "Instead, Write Buffer Size ='%v'\n",
      actualWriterBuffSize)
  }
}

func TestFileMgr_IsAbsolutePathFileNamePopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

  if !isAbsPathFileName {
    t.Error("Error: Expected Absolute Path File Name to be populated.\n" +
      "It was NOT!\n")
  }

}

func TestFileMgr_IsAbsolutePathFileNamePopulated_02(t *testing.T) {

  srcFMgr := FileMgr{}

  isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

  if isAbsPathFileName {
    t.Error("Error: Expected Absolute Path File Name NOT populated.\n" +
      "WRONG - It IS populated!\n")
  }

}

func TestFileMgr_IsFileExtPopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  isFileExtPopulated := srcFMgr.IsFileExtPopulated()

  if !isFileExtPopulated {
    t.Error("Expected srcFMgr.IsFileExtPopulated() == 'true'.\n" +
      "Instead, it is 'false'\n")
  }

}

func TestFileMgr_IsFileExtPopulated_02(t *testing.T) {

  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n",
      targetDir, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile).\n"+
      "DirMgr='%v'\ntargetFile='%v'\nError='%v'\n",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
    return
  }

  isFileExtPopulated := srcFMgr.IsFileExtPopulated()

  if isFileExtPopulated {
    t.Error("Expected srcFMgr.IsFileExtPopulated() == 'false'.\n" +
      "Instead, it is 'true'\n")
  }

}

func TestFileMgr_IsFileNameExtPopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  isFileNameExtPopulated := srcFMgr.IsFileNameExtPopulated()

  if !isFileNameExtPopulated {
    t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'true'.\n" +
      "Instead, it is 'false'\n")
  }
}

func TestFileMgr_IsFileNameExtPopulated_02(t *testing.T) {
  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n",
      targetDir, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile).\n"+
      "DirMgr='%v'\ntargetFile='%v'\nError='%v'\n",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
    return
  }

  isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

  if isFileNamePopulated {

    t.Errorf("Expected srcFMgr.IsFileNameExtPopulated() == 'false'.\n"+
      "Instead, it is 'true'.\n"+
      "FileName='%v'\nFile Extension='%v'\nLen File Ext='%v'\n",
      srcFMgr.GetFileName(), srcFMgr.GetFileExt(), len(srcFMgr.GetFileExt()))
  }

}

func TestFileMgr_IsFileNameExtPopulated_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

  if isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'false'.\n" +
      "Instead, it is 'true'\n")
  }

}

func TestFileMgr_IsFileNamePopulated_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if !isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'.\n" +
      "Instead, it is 'false'\n")
  }
}

func TestFileMgr_IsFileNamePopulated_02(t *testing.T) {

  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n",
      targetDir, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile).\n"+
      "DirMgr='%v'\ntargetFile='%v'\nError='%v'\n",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
    return
  }

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if !isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'.\nInstead, it is 'false'")
  }

}

func TestFileMgr_IsFileNamePopulated_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFileNamePopulated := srcFMgr.IsFileNamePopulated()

  if isFileNamePopulated {
    t.Error("Expected srcFMgr.IsFileNamePopulated() == 'false'.\n" +
      "Instead, it is 'true'\n")
  }

}

func TestFileMgr_IsFilePointerOpen_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  err = srcFMgr.OpenThisFileReadOnly()

  if err != nil {
    _ = srcFMgr.CloseThisFile()
    t.Errorf("Error returned from srcFMgr.OpenThisFileReadOnly().\n"+
      "srcFMgr='%v'\nError='%v'",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  err = srcFMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from final srcFMgr.CloseThisFile().\n"+
      "srcFMgr='%v'\nError='%v'",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if !isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'true'.\n" +
      "Instead, it is FALSE!\n")
  }
}

func TestFileMgr_IsFilePointerOpen_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  if isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'false'.\n" +
      "Instead, it is TRUE!\n")
  }
}

func TestFileMgr_IsFilePointerOpen_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isFilePointerOpen := srcFMgr.IsFilePointerOpen()

  if isFilePointerOpen {
    t.Error("Expected isFilePointerOpen = 'false'.\n" +
      "Instead, it is TRUE!\n")
  }

}

func TestFileMgr_IsInitialized_01(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  isInitialized := srcFMgr.IsInitialized()

  if !isInitialized {
    t.Error("Expected isInitialized = 'true'.\n" +
      "Instead, it is FALSE!\n")
  }
}

func TestFileMgr_IsInitialized_02(t *testing.T) {

  fh := FileHelper{}

  targetFile := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := FileMgr{}.New(targetFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
    return
  }

  srcFMgr.Empty()

  isInitialized := srcFMgr.IsInitialized()

  if isInitialized {
    t.Error("Expected isInitialized = 'false'.\n" +
      "Instead, it is TRUE!\n")
  }
}

func TestFileMgr_IsInitialized_03(t *testing.T) {

  srcFMgr := FileMgr{}

  isInitialized := srcFMgr.IsInitialized()

  if isInitialized {
    t.Error("Expected isInitialized = 'false'.\n" +
      "Instead, it is TRUE!\n")
  }
}
