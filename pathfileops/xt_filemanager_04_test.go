package pathfileops

import (
  "fmt"
  "strings"
  "testing"
)

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

  if bufReader == nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioReader(). " +
      "Pointer IS NIL!")
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

  if bufReader == nil {
    t.Error("Error: Expected pointer return from srcFMgr.GetBufioWriter(). " +
      "Pointer IS NIL!")
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
      "because the file does NOT exist. However, NO ERROR WAS RETURNED!" )
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
      "because the file manager (srcFmgr) is invalid. However, NO ERROR WAS RETURNED!" )
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
