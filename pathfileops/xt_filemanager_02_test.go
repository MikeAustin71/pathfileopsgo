package pathfileops

import (
  "io/ioutil"
  "os"
  "testing"
)

func TestFileMgr_OpenThisFileReadOnly_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  b, err := ioutil.ReadAll(fMgr.filePtr)

  if err != nil {
    _ = fMgr.CloseThisFile()
    t.Errorf("Error returned from ioutil.ReadAll(fMgr.filePtr) filePath='%v'  Error='%v'", filePath, err.Error())
    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {
    t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_OpenThisFileReadWrite_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {

    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
    return
  }

  b, err := ioutil.ReadAll(fMgr.filePtr)

  if err != nil {

    _ = fMgr.CloseThisFile()

    t.Errorf("Error returned from ioutil.ReadAll(fMgr.filePtr) filePath='%v'  Error='%v'", filePath, err.Error())

    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {

    _ = fMgr.CloseThisFile()

    t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
    return
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_ResetFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.ResetFileInfo()

  if err != nil {
    t.Errorf("Error returned by fMgr.ResetFileInfo(). Error='%v' ", err.Error())
  }

  fInfoPlus, err := fMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned by fMgr.GetFileInfoPlus(). Error='%v' ", err.Error())
  }

  if "testRead2008.txt" != fInfoPlus.fName {
    t.Errorf("Expected file name== 'testRead2008.txt'. "+
      "Instead, file name=='%v' ", fInfoPlus.fName)
  }

}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  byteBuff := make([]byte, 2048, 2048)

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  var rStr = make([]rune, 0, 2048)

  for i := 0; i < len(byteBuff); i++ {

    if byteBuff[i] == 0 {
      break
    }

    rStr = append(rStr, rune(byteBuff[i]))

  }

  expectedStr := "Test Read File. Do NOT alter the contents of this file."
  actualStr := string(rStr)

  if expectedStr != actualStr {
    t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.  Instead, Actual Bytes Read='%v'", expectedBytesRead, bytesRead)
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_SetFileInfo(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
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
    t.Errorf("Error = Expected fmgr.actualFileInfo.Name()='%v'.  Instead, fmgr.actualFileInfo.Name()='%v'", expectedFileNameExt, fmgr.actualFileInfo.Name())
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
