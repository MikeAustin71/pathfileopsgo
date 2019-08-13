package pathfileops

import (
  "io"
  "strings"
  "testing"
)

func TestFileMgr_ReadAllFile_01(t *testing.T) {

  expectedBytes := 8819 // int

  fh := FileHelper{}

  setupFileName := "testRead857268.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by " +
      "fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by " +
      "fh.MakeDirAll(absBaseFilePath).\n"+
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

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'.\n"+
      "Instead, the number of bytes read='%v'\n",
      expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("Error: Expected fMgr.filePtr == nil.\n" +
      "fMgr.filePtr IS NOT NIL!\n")
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadAllFile_02(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead857268.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  fMgr.isInitialized = false

  _, err = fMgr.ReadAllFile()

  if err == nil {
    t.Error("Expected an error return from fMgr.ReadAllFile()\n" +
      "because 'fMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadAllFile_03(t *testing.T) {

  expectedBytes := 155 // int

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'.\n"+
      "Instead, the number of bytes read='%v'\n",
      expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileLine_01(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  bytes, err := fMgr.ReadFileLine(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr."+
      "ReadFileLine(delim) on Line#1.\n"+
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
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "However, fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Now is the time for all good men" != actualStr {
    t.Errorf("Expected line #1 = 'Now is the time for "+
      "all good men'.\n"+
      "Instead, line #1 = '%v'\n",
      actualStr)
  }

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileLine_02(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  _, err = fMgr.ReadFileLine(delim)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileLine(delim) on Line#1\n" +
      "because fMgr.isInitialized = false.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileLine_03(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileLine(delim) on "+
        "Line#1.\n"+
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

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\n Error='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile(), expected " +
      "fMgr.filePtr==nil.\n" +
      "However, fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    _ = fMgr.CloseThisFile()
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4 = 'Thank you, for your support.'\n"+
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileLine_04(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileLine"+
        "(delim) on Line#1.\n"+
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

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected fMgr.filePtr==nil.\n" +
      "However, fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4 = 'Thank you, for your support.'\n"+
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileLine_05(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03/checkfiles03_02")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  bytes, err := fMgr.ReadFileLine(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileLine(delim) "+
      "on Line#1.\n"+
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
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Now is the time for all good men" != actualStr {
    t.Errorf("Expected line #1= 'Now is the time for all good men'.\n"+
      "Instead, line #1 = '%v'\n", actualStr)
    return
  }

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

  fh := FileHelper{}

  setupFileName := "testRead2008.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03\").\n"+
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

  byteBuff := make([]byte, 2048, 2048)

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil &&
    err != io.EOF {
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  var rStr = make([]rune, 0, 2048)

  for i := 0; i < len(byteBuff); i++ {

    if byteBuff[i] == 0 {
      break
    }

    rStr = append(rStr, rune(byteBuff[i]))
  }

  expectedStr :=
    "Test Read File. Do NOT alter the contents of this file."
  actualStr := string(rStr)

  if expectedStr != actualStr {
    t.Errorf("Expected Read String='%v'.\n"+
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.\n"+
      "Instead, Actual Bytes Read='%v'\n",
      expectedBytesRead, bytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile()\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileBytes_02(t *testing.T) {

  fh := FileHelper{}
  setupFileName := "testRead2008.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  fMgr.isInitialized = false

  byteBuff := make([]byte, 2048, 2048)

  _, err = fMgr.ReadFileBytes(byteBuff)

  if err == nil {
    t.Error("Expected error return from fMgr.ReadFileBytes(" +
      "byteBuff)\n" +
      "because fMgr.isInitialized = false.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadFileBytes_03(t *testing.T) {

  fh := FileHelper{}
  setupFileName := "testRead2008.txt"

  setupFile := fh.AdjustPathSlash(
    "../../filesfortest/checkfiles/" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "../../checkfiles/checkfiles03/checkfiles03_02/" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "../../checkfiles/checkfiles03")

  if err != nil {
    t.Errorf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../../checkfiles/checkfiles03/checkfiles03_02\").\n"+
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

  byteBuff := make([]byte, 2048, 2048)

  fMgr.fileBufRdr = nil
  fMgr.fileRdrBufSize = 16384

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil &&
    err != io.EOF {
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff).\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
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
    t.Errorf("Expected Read String='%v'.\n"+
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.\n"+
      "Instead, Actual Bytes Read='%v'\n",
      expectedBytesRead, bytesRead)
  }

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()
}
