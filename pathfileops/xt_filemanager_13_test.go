package pathfileops

import (
  "io"
  "strings"
  "testing"
)

func TestFileMgr_ReadFileString_01(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) "+
      "on Line#1.\n"+
      "fMgr='%v'\nError='%v'",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileString_02(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  delim := byte('\n')

  var actualStr string

  for i := 0; i < 4; i++ {

    actualStr, err = fMgr.ReadFileString(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileString(delim) "+
        "on Line#1.\n"+
        "fMgr='%v'\nError='%v'\n",
        fMgr.GetAbsolutePathFileName(), err.Error())
      _ = fMgr.CloseThisFile()
      return
    }
  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else if strings.Index(actualStr, "\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  if fMgr.filePtr != nil {
    t.Errorf("ERROR: After fMgr.CloseThisFile(), "+
      "expected fMgr.filePtr==nil.\n"+
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n"+
      "fileMgr='%v'\n",
      fMgr.GetAbsolutePathFileName())
    _ = fMgr.CloseThisFile()
    return
  }

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4= 'Thank you, for your support.'\n"+
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim) " +
      "to be io.EOF. Instead, error WAS NOT equal to io.EOF!")
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileString_03(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString("+
      "delim) on Line#1.\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile(), expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.\n"+
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileString_04(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) "+
      "on Line#1.\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

  _ = fMgr.DeleteThisFile()

}

func TestFileMgr_ReadFileString_05(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadWrite().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  fMgr.fileAccessStatus.Empty()

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString("+
      "delim) on Line#1.\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile(), expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else if strings.Index(actualStr, "\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.\n"+
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileString_06(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  delim := byte('\n')

  fMgr.isInitialized = false

  _, err = fMgr.ReadFileString(delim)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileLine(" +
      "delim) on Line#1\n" +
      "because fMgr.isInitialized = false.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileString_07(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  delim := byte('\n')

  fMgr.fileBufRdr = nil
  fMgr.fileRdrBufSize = 16384

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString("+
      "delim) on Line#1.\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.\n"+
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ResetFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead2008.txt"

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../checkfiles/checkfiles03")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.ResetFileInfo()

  if err != nil {
    t.Errorf("Error returned by fMgr.ResetFileInfo().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  fInfoPlus, err := fMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned by fMgr.GetFileInfoPlus().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  if "testRead2008.txt" != fInfoPlus.fName {
    t.Errorf("Expected file name== 'testRead2008.txt'.\n"+
      "Instead, file name=='%v'\n",
      fInfoPlus.fName)
  }

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}
