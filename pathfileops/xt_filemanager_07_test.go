package pathfileops

import (
  "fmt"
  "os"
  "testing"
)

func TestFileMgr_SetFileInfo_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  err = fmgr.SetFileInfo(info)

  if err != nil {
    t.Errorf("Error returned by fmgr.SetFileInfo(info). info.Name()='%v'  Error='%v'", info.Name(), err.Error())
  }

  if !fmgr.actualFileInfo.IsFInfoInitialized {
    t.Error("Error - File Manager FileInfoPlus object is not initialized!")
  }

  if fmgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("Error = Expected fmgr.actualFileInfo.Name()='%v'.  "+
      "Instead, fmgr.actualFileInfo.Name()='%v'",
      expectedFileNameExt, fmgr.actualFileInfo.Name())
  }

}

func TestFileMgr_SetFileInfo_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  err = fmgr.SetFileInfo(nil)

  if err == nil {
    t.Error("Expected error return from fmgr.SetFileInfo(nil) because " +
      "input parameter fileInfo is 'nil'. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileInfo_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  err = fmgr.SetFileInfo(info)

  if err == nil {
    t.Error("Expected an error from fmgr.SetFileInfo(info) because input parameter 'info' " +
      "is a directory an not a file. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileInfo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  alternativeFileNameExt := "newerFileForTest_03.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  absAlternatePathFileNameExt := absPath + string(os.PathSeparator) + alternativeFileNameExt

  info, err := fh.GetFileInfo(absAlternatePathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  err = fmgr.SetFileInfo(info)

  if err == nil {
    t.Error("Expected an error from fmgr.SetFileInfo(info) because input parameter 'info' " +
      "is a file name which is different from fmgr file name. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt) " +
      "isEmpty=='false'. Instead, isEmpty=='true'.")
  }

  if err != nil {
    t.Errorf("Error returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt). "+
      "dMgr='%v' expectedFileNameExt='%v'  Error='%v'",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_02(t *testing.T) {

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  _, err = fMgr.SetFileMgrFromDirMgrFileName(dMgr, "")

  if err == nil {
    t.Error("Expected error return from fMgr.SetFileMgrFromDirMgrFileName(dMgr, \"\") " +
      "because input parameter fileNameExt is an empty string! " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_03(t *testing.T) {

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  _, err = fMgr.SetFileMgrFromDirMgrFileName(dMgr, "    ")

  if err == nil {
    t.Error("Expected error return from fMgr.SetFileMgrFromDirMgrFileName(dMgr, \"  \") " +
      "because input parameter fileNameExt consists of blank spaces! " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  dMgr.isInitialized = false

  _, err = fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from fMgr.SetFileMgrFromDirMgrFileName(dMgr, \"  \") " +
      "because input parameter dMgr is invalid! " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_05(t *testing.T) {

  expectedFileNameExt := "basefilenoext"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt) " +
      "isEmpty=='false'. Instead, isEmpty=='true'.")
  }

  if err != nil {
    t.Errorf("Error returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt). "+
      "dMgr='%v' expectedFileNameExt='%v'  Error='%v'",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_06(t *testing.T) {

  expectedFileNameExt := ".xgitignore"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt) " +
      "isEmpty=='false'. Instead, isEmpty=='true'.")
  }

  if err != nil {
    t.Errorf("Error returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt). "+
      "dMgr='%v' expectedFileNameExt='%v'  Error='%v'",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_01(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\logTest\\CmdrX\\CmdrX.log"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
      "relPath='%v'  Error='%v'", relPath, err.Error())
  }

  fileName := "CmdrX"
  fileNameExt := "CmdrX.log"
  extName := ".log"

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)  "+
      "pathFileNameExt='%v' Error='%v'",
      pathFileNameExt, err.Error())
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt) expected " +
      "isEmpty='false'. Instead, isEmpty='false'. ")
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_02(t *testing.T) {

  fileMgr := FileMgr{}

  _, err := fileMgr.SetFileMgrFromPathFileName("")

  if err == nil {
    t.Error("Expected error return from fileMgr.SetFileMgrFromPathFileName(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_03(t *testing.T) {

  fileMgr := FileMgr{}

  _, err := fileMgr.SetFileMgrFromPathFileName("      ")

  if err == nil {
    t.Error("Expected error return from fileMgr.SetFileMgrFromPathFileName(\"    \") " +
      "because the input parameter consists of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_04(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\filesfortest\\basefilesfortest\\basefilenoext"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
      "relPath='%v'  Error='%v'", relPath, err.Error())
  }

  fileName := "basefilenoext"
  fileNameExt := "basefilenoext"
  extName := ""

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)  "+
      "pathFileNameExt='%v' Error='%v'",
      pathFileNameExt, err.Error())
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt) expected " +
      "isEmpty='false'. Instead, isEmpty='false'. ")
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='false', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_05(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\filesfortest\\basefilesfortest\\.xgitignore"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
      "relPath='%v'  Error='%v'", relPath, err.Error())
  }

  fileName := ".xgitignore"
  fileNameExt := ".xgitignore"
  extName := ""

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)  "+
      "pathFileNameExt='%v' Error='%v'",
      pathFileNameExt, err.Error())
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt) expected " +
      "isEmpty='false'. Instead, isEmpty='false'. ")
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='false', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_WriteBytesToFile_01(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWriteXX241289.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk(). Error='%v' ",
      err.Error())
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead). Error='%v'",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().")
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'. Instead, stringRead='%v' ",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText. verifyBytesWritten='%v' "+
      "lenTestText='%v' ", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText. numBytesRead='%v' "+
      "lenTestText='%v' ", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten. numBytesRead='%v' "+
      "numBytesWritten='%v' ", numBytesRead, numBytesWritten)

  }

}

func TestFileMgr_WriteBytesToFile_02(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWriteXX241289.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  bytesToWrite := []byte(testText)

  fMgr.isInitialized = false

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err == nil {
    t.Error("Expected an error from fMgr.WriteBytesToFile(bytesToWrite) " +
      "because fMgr.isInitialized == false. Instead, NO ERROR WAS RETURNED!")
  }

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr2.DeleteThisFile(). Error='%v' ",
      err.Error())
  }

}

func TestFileMgr_WriteBytesToFile_03(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := fh.AdjustPathSlash("../checkfiles/scratchTestWriteVV6431271.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile(). "+
      "filePathName='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). "+
      "Error='%v' ", err.Error())
  }

  if !fMgr.DoesFileExist() {
    t.Errorf("Error: Failed to create File Name:'%v'. Terminating test.",
      fMgr.GetAbsolutePathFileName())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk(). Error='%v' ",
      err.Error())
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().")
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead). Error='%v'",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #3 fMgr.CloseThisFile().")
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'. Instead, stringRead='%v' ",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText. verifyBytesWritten='%v' "+
      "lenTestText='%v' ", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText. numBytesRead='%v' "+
      "lenTestText='%v' ", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten. numBytesRead='%v' "+
      "numBytesWritten='%v' ", numBytesRead, numBytesWritten)

  }

}

func TestFileMgr_WriteStrToFile_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWrite2998.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  expectedStr := "Test Write File. Do NOT alter the contents of this file."

  lExpectedStr := len(expectedStr)

  bytesWritten, err := fMgr.WriteStrToFile(expectedStr)

  if err != nil {
    t.Errorf("Error returned from fMgr.WriteStrToFile(expectedStr)  expectedStr='%v'  Error='%v'", expectedStr, err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile() No 1.  Error='%v'", err.Error())
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadAllFile(). filePathName='%v'  Error='%v'", fMgr.absolutePathFileName, err.Error())
  }

  if lExpectedStr != bytesWritten {
    t.Errorf("Error: Length of string written NOT equal to Bytes Read! Length of written string='%v'. Actual Bytes Read='%v' ", lExpectedStr, bytesWritten)
  }

  actualStr := string(bytesRead)

  if lExpectedStr != len(actualStr) {
    t.Errorf("Error: Length of actual string read is NOT equal to length of string written. lExpectedStr='%v'  len(actualStr)='%v'", lExpectedStr, len(actualStr))
  }

  if expectedStr != actualStr {
    t.Errorf("Error: expectedStr written='%v'  Actual string read='%v'", expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile() No 2. Error='%v'", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile(). Error='%v'", err.Error())
  }

  doesFileExist := fh.DoesFileExist(filePath)

  if doesFileExist {
    t.Errorf("Error: Failed to DELETE fileNameExt='%v'", fMgr.absolutePathFileName)
  }

}

func TestFileMgr_WriteStrToFile_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWrite2998.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile(). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile() No 1.  Error='%v'", err.Error())
  }

  expectedStr := "Test Write File. Do NOT alter the contents of this file."

  fMgr.isInitialized = false

  _, err = fMgr.WriteStrToFile(expectedStr)

  if err == nil {
    t.Error("Expected an error to be returned from fMgr.WriteStrToFile(expectedStr) " +
      "because fMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error Deleting File: %v. Error returned by fMgr2.DeleteThisFile(). "+
      "Error='%v'", fMgr2.GetAbsolutePathFileName(), err.Error())
  }

}
