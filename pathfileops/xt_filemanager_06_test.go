package pathfileops

import (
  "io"
  "io/ioutil"
  "strings"
  "testing"
)

func TestFileMgr_OpenThisFile_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFile(FileAccessControl{})

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(FileAccessControl{}) because " +
      "input parameter FileAccessControl{} is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_OpenThisFile_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../iDoNotExist/iDoNotExist2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()). "+
      "Error='%v'", err.Error())
  }

  fPerm, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rwxrwxrwx\"). "+
      "Error='%v'", err.Error())
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(FileAccessControl{}) because " +
      "input parameter FileAccessControl{} is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = fMgr.CloseThisFile()

  dMgr := fMgr.GetDirMgr()

  _ = dMgr.DeleteAll()

}

func TestFileMgr_OpenThisFile_03(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/scratchTestFile0812.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.CreateThisFile()

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()). "+
      "Error='%v'", err.Error())
  }

  fPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\"). "+
      "Error='%v'", err.Error())
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFile(fileAccessCtrl). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile(). "+
      "Error='%v'", err.Error())
  }

}

func TestFileMgr_OpenThisFile_04(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()). "+
      "Error='%v'", err.Error())
  }

  fPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\"). "+
      "Error='%v'", err.Error())
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm). "+
      "Error='%v'", err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(fileAccessCtrl) because " +
      "fMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = fMgr.CloseThisFile()

}

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

func TestFileMgr_OpenThisFileReadOnly_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileReadOnly()

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFileReadOnly() because " +
      "fMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_OpenThisFileReadOnly_03(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/scratchTestRead067894.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.DeleteThisFile(). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile(). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePath(), err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePath(), err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile(). "+
      "Error='%v'", err.Error())
  }

}

func TestFileMgr_OpenThisFileReadOnly_04(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/iDoNotExist.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePath(), err.Error())
  }

  _ = fMgr.DeleteThisFile()

  err = fMgr.OpenThisFileReadOnly()

  if err == nil {
    t.Error("Expected an error return from fMgr.OpenThisFileReadOnly() because " +
      "the fMgr file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()

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

  if err != nil {
    t.Errorf("Error returned from final fMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

}

func TestFileMgr_OpenThisFileReadWrite_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileReadWrite()

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFileReadWrite() because " +
      "'fMgr' is invalid. However, NO ERROR WAS RETURNED!")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_OpenThisFileWriteOnly_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/scratchTestWrite647182.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile(). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile(). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileWriteOnly()

  if err == nil {
    t.Error("Expected an error returned from fMgr.OpenThisFileWriteOnly() because " +
      "fMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileWriteOnly_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/newDir/scratchTestWrite655349.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    t.Errorf("Error returned from fMgr.OpenThisFileWriteOnly(). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fMgr.DoesFileExist() {
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    t.Errorf("Error: The test file %v should have been created. However, the file "+
      "does NOT exist!", fMgr.GetAbsolutePathFileName())
  }

  numOfBytesWritten, err := fMgr.WriteStrToFile("Hello world!")

  if err != nil {
    t.Errorf("Error returned from fMgr.WriteStrToFile(\"Hello world!\"). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if numOfBytesWritten < 12 {
    t.Errorf("Expected at least 12-bytes to be written to file. However, "+
      "only %v-bytes were written. ", numOfBytesWritten)
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileWriteOnlyAppend_01(t *testing.T) {

  fh := FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  testText2 := "Damn the torpedoes, full speed ahead!\n"

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

  bytesToWrite := []byte(testText1)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnlyAppend(). Error='%v' ",
      err.Error())
  }

  bytesToWrite = []byte(testText2)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  bytesRead, err := fMgr.ReadFileLine('\n')

  bytesRead, err = fMgr.ReadFileLine('\n')

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
  }

  stringRead := string(bytesRead)

  stringRead = strings.Replace(stringRead, "\r\n", "", -1)

  if testText2 != stringRead {
    t.Errorf("Error: Expected stringRead='%v'. Instead, stringRead='%v' ",
      testText2, stringRead)
  }
}

func TestFileMgr_OpenThisFileWriteOnlyAppend_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/scratchTestWriteFX471985.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
  }

  if !fMgr.DoesFileExist() {
    t.Errorf("Error: Expected target file to be created. However, File:'%v' "+
      "does NOT exist.", fMgr.GetAbsolutePathFileName())
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()

}

func TestFileMgr_OpenThisFileWriteOnlyAppend_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/scratchTestWriteJU81294823.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFileWriteOnlyAppend() because " +
      "'fMgr' is invalid. However, NO ERROR WAS RETURNED!")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()

  if fMgr.DoesFileExist() {
    t.Errorf("Error: Expected target file to be deleted. However, File:'%v' "+
      "DOES exist.", fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_ReadAllFile_01(t *testing.T) {

  expectedBytes := int(8819)

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead857268.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile(). Error='%v' ",
      err.Error())
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'. Instead, "+
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

func TestFileMgr_ReadAllFile_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead857268.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr.isInitialized = false

  _, err = fMgr.ReadAllFile()

  if err == nil {
    t.Error("Expected an error return from fMgr.ReadAllFile() because " +
      "'fMgr' is invalid. However, NO ERROR WAS RETURNED!")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_ReadAllFile_03(t *testing.T) {

  expectedBytes := int(155)

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile(). Error='%v' ",
      err.Error())
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'. Instead, "+
      "the number of bytes read='%v' ", expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

}

func TestFileMgr_ReadFileLine_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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

func TestFileMgr_ReadFileLine_05(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  fMgr.fileBufRdr = nil
  fMgr.fileRdrBufSize = 16384

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

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Now is the time for all good men" != actualStr {
    t.Errorf("Expected line #1 = 'Now is the time for all good men'. Instead, "+
      "line #1 = '%v'", actualStr)
  }

}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  byteBuff := make([]byte, 2048, 2048)

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff). "+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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

func TestFileMgr_ReadFileBytes_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  byteBuff := make([]byte, 2048, 2048)

  fMgr.fileBufRdr = nil
  fMgr.fileRdrBufSize = 16384

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil {
    _ = fMgr.CloseThisFile()
    t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff). "+
      "filePath='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
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
    _ = fMgr.CloseThisFile()
    t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.  Instead, Actual Bytes Read='%v'", expectedBytesRead, bytesRead)
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_ReadFileString_01(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else if strings.Index(actualStr, "\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-1]
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
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). "+
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
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
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). "+
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
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
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadWrite(). "+
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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else if strings.Index(actualStr, "\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-1]
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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

func TestFileMgr_ReadFileString_07(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  fMgr.fileBufRdr = nil
  fMgr.fileRdrBufSize = 16384

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

  if strings.Index(actualStr, "\r\n") > -1 {
    actualStr = actualStr[0 : len(actualStr)-2]
  } else {
    actualStr = actualStr[0 : len(actualStr)-1]
  }

  if expectedStr != actualStr {
    t.Errorf("Expected line #1 = '%v'.  Instead, "+
      "line #1 = '%v'", expectedStr, actualStr)
  }

}

func TestFileMgr_ResetFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
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
