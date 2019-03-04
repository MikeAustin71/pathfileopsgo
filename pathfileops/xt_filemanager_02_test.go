package pathfileops

import (
  "io"
  "io/ioutil"
  "os"
  "strings"
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

    t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

}

func TestFileMgr_ReadAllFile_01(t *testing.T) {

  expectedBytes := int(8819)

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead857268.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile(). Error='%v' ",
      err.Error())
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'. Instead, " +
      "the number of bytes read='%v' ", expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("Error: Expected fMgr.filePtr == nil. fMgr.filePtr IS NOT NIL!")
  }

}

func TestFileMgr_ReadFileLine_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  bytes, err := fMgr.ReadFileLine(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileLine(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Now is the time for all good men" != actualStr {
    t.Errorf("Expected line #1 = 'Now is the time for all good men'. Instead, "+
      "line #1 = '%v'", actualStr)
  }

}

func TestFileMgr_ReadFileLine_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  fMgr.isInitialized = false

  _, err = fMgr.ReadFileLine(delim)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileLine(delim) on Line#1 " +
      "because fMgr.isInitialized = false. However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileMgr_ReadFileLine_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileLine(delim) on Line#1. "+
        "Error='%v'", err.Error())
    }
  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4 = 'Thank you, for your support.'. Instead, "+
      "line #4 = '%v'", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim) " +
      "to be io.EOF. Instead, error WAS NOT equal to io.EOF!")
  }

}

func TestFileMgr_ReadFileLine_04(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). Error='%v'",
      err.Error())
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileLine(delim) on Line#1. "+
        "Error='%v'", err.Error())
    }
  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4 = 'Thank you, for your support.'. Instead, "+
      "line #4 = '%v'", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim) " +
      "to be io.EOF. Instead, error WAS NOT equal to io.EOF!")
  }

}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  byteBuff := make([]byte, 2048, 2048)

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
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

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile() Error='%v'", err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

}

func TestFileMgr_ReadFileBytes_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr.isInitialized = false

  byteBuff := make([]byte, 2048, 2048)

  _, err = fMgr.ReadFileBytes(byteBuff)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileBytes(byteBuff) " +
      "because fMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_ReadFileString_01(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  if strings.Index(actualStr,"\r\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-2]
  } else {
    actualStr = actualStr[0:len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

}

func TestFileMgr_ReadFileString_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). Error='%v'",
      err.Error())
  }

  delim := byte('\n')

  var actualStr string

  for i := 0; i < 4; i++ {

    actualStr, err = fMgr.ReadFileString(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileString(delim) on Line#1. "+
        "Error='%v'", err.Error())
    }
  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  if strings.Index(actualStr,"\r\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-2]
  } else if strings.Index(actualStr,"\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-1]
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4 = 'Thank you, for your support.'. Instead, "+
      "line #4 = '%v'", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim) " +
      "to be io.EOF. Instead, error WAS NOT equal to io.EOF!")
  }

}

func TestFileMgr_ReadFileString_03(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). " +
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  if strings.Index(actualStr,"\r\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-2]
  } else {
    actualStr = actualStr[0:len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

}

func TestFileMgr_ReadFileString_04(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). " +
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  if strings.Index(actualStr,"\r\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-2]
  } else {
    actualStr = actualStr[0:len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

}

func TestFileMgr_ReadFileString_05(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadWrite(). " +
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  fMgr.fileAccessStatus.Empty()

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil. " +
      "fMgr.filePtr IS NOT EQUAL TO NIL!")
  }

  if strings.Index(actualStr,"\r\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-2]
  } else if strings.Index(actualStr,"\n") > -1 {
    actualStr = actualStr[0:len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

}

func TestFileMgr_ReadFileString_06(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  fMgr.isInitialized = false

  _, err = fMgr.ReadFileString(delim)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileLine(delim) on Line#1 " +
      "because fMgr.isInitialized = false. However, NO ERROR WAS RETURNED! ")
  }

}

func TestFileMgr_ResetFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). " +
      "filePath='%v'  Error='%v'", filePath, err.Error())
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
