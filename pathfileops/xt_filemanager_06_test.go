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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFile(FileAccessControl{})

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(FileAccessControl{})\n" +
      "because input parameter FileAccessControl{} is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_OpenThisFile_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../iDoNotExist/iDoNotExist2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()).\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fPerm, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rwxrwxrwx\").\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm).\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(FileAccessControl{})\n" +
      "because input parameter FileAccessControl{} is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = fMgr.CloseThisFile()

  dMgr := fMgr.GetDirMgr()

  _ = dMgr.DeleteAll()
}

func TestFileMgr_OpenThisFile_03(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/TestFileMgr_OpenThisFile_03.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error retruned by fMgr.CreateThisFile()\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()).\n"+
      "Error='%v'\n",
      err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\").\n"+
      "Error='%v'\n",
      err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm).\n"+
      "Error='%v'\n",
      err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFile(fileAccessCtrl).\n"+
      "Error='%v'\n",
      err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile(). "+
      "Error='%v'", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile().\n"+
      "Error='%v'\n",
      err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }
}

func TestFileMgr_OpenThisFile_04(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned from FileOpenConfig{}.New(FOpenType.TypeReadWrite()).\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\").\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fOpenCfg, fPerm)

  if err != nil {
    t.Errorf("Error returned from FileAccessControl{}.New(fOpenCfg, fPerm).\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFile(fileAccessCtrl)

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFile(fileAccessCtrl)\n" +
      "because fMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_OpenThisFileReadOnly_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  b, err := ioutil.ReadAll(fMgr.filePtr)

  if err != nil {
    _ = fMgr.CloseThisFile()
    t.Errorf("Error returned from ioutil.ReadAll(fMgr.filePtr).\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {
    t.Errorf("Expected Read String='%v'.\n" +
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
    return
  }

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_OpenThisFileReadOnly_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileReadOnly()

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFileReadOnly()\n" +
      "because fMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_OpenThisFileReadOnly_03(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/TestFileMgr_OpenThisFileReadOnly_03.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from #1 fMgr.DoesThisFileExist().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  if fileDoesExist {
    t.Errorf("ERROR: Test file should NOT exist!.\n" +
      "However, test file DOES EXIST!\n" +
      "test file='%v'", filePath)
    _ = fh.DeleteDirFile(filePath)
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile().\n" +
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())

    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }
}

func TestFileMgr_OpenThisFileReadOnly_04(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/iDoNotExist.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  _ = fMgr.DeleteThisFile()

  err = fMgr.OpenThisFileReadOnly()

  if err == nil {
    t.Error("Expected an error return from fMgr.OpenThisFileReadOnly()\n" +
      "because the fMgr file does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileReadWrite_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {

    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  b, err := ioutil.ReadAll(fMgr.GetFilePtr())

  if err != nil {

    t.Errorf("Error returned from ioutil.ReadAll(fMgr.GetFilePtr())\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())

    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()

    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {

    t.Errorf("Expected Read String='%v'.\n" +
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from final fMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileReadWrite_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileReadWrite()

  if err == nil {
    t.Error("Expected error return from fMgr.OpenThisFileReadWrite()\n" +
      "because 'fMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_OpenThisFileWriteOnly_01(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/TestFileMgr_OpenThisFileWriteOnly_01.txt")

  err := fh.DeleteDirFile(filePath)

  if err !=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile().\n"+
      "filePath='%v'\nError='%v'",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileWriteOnly()

  if err == nil {
    t.Error("Expected an error returned from fMgr.OpenThisFileWriteOnly()\n" +
      "because fMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileWriteOnly_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/newDir/TestFileMgr_OpenThisFileWriteOnly_02.txt")

  err := fh.DeleteDirFile(filePath)

  if err !=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    t.Errorf("Error returned from fMgr.OpenThisFileWriteOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from fMgr.DoesThisFileExist()\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if !fileDoesExist {
    t.Errorf("Error: The test file should have been created.\n" +
      "However, the test file does NOT exist!\n" +
      "Test File='%v'", fMgr.GetAbsolutePathFileName())

    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  numOfBytesWritten, err := fMgr.WriteStrToFile("Hello world!")

  if err != nil {
    t.Errorf("Error returned from fMgr.WriteStrToFile(\"Hello world!\").\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if numOfBytesWritten < 12 {
    t.Errorf("Expected at least 12-bytes to be written to file.\n" +
      "However, only %v-bytes were written.\n", numOfBytesWritten)
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileWriteOnlyAppend_01(t *testing.T) {

  fh := FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  testText2 := "Damn the torpedoes, full speed ahead!\n"

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/TestFileMgr_OpenThisFileWriteOnlyAppend_01.txt")

  err := fh.DeleteDirFile(filePath)

  if err !=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr" +
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() , err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() ,err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnlyAppend().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite = []byte(testText2)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.WriteBytesToFile(bytesToWrite).\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead, err := fMgr.ReadFileLine('\n')

  bytesRead, err = fMgr.ReadFileLine('\n')

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED!\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  stringRead := string(bytesRead)

  stringRead = strings.Replace(stringRead, "\r\n", "", -1)

  if testText2 != stringRead {
    t.Errorf("Error: Expected stringRead='%v'.\n" +
      "Instead, stringRead='%v'\n",
      testText2, stringRead)
  }

  _ = fMgr.DeleteThisFile()
  return
}

func TestFileMgr_OpenThisFileWriteOnlyAppend_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/TestFileMgr_OpenThisFileWriteOnlyAppend_02.txt")

  err := fh.DeleteDirFile(filePath)

  if err !=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() ,err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by fMgr.DoesThisFileExist()\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if !fileDoesExist {
    t.Errorf("Error: Expected target file to be created.\n" +
      "However, File:'%v' "+
      "does NOT exist.", fMgr.GetAbsolutePathFileName())
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_OpenThisFileWriteOnlyAppend_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/TestFileMgr_OpenThisFileWriteOnlyAppend_03.txt")

  err := fh.DeleteDirFile(filePath)

  if err !=nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  fMgr.isInitialized = false

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err == nil {
    t.Error("ERROR: Expected error return from fMgr.OpenThisFileWriteOnlyAppend()\n" +
      "because 'fMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by fMgr.DoesThisFileExist()\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if fileDoesExist {
    t.Errorf("Error: Expected target file WOULD NOT be created.\n" +
      "However, target file DOES EXIST!\n"+
      "target file='%v'", fMgr.GetAbsolutePathFileName())
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}

func TestFileMgr_ReadAllFile_01(t *testing.T) {

  expectedBytes := int(8819)

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead857268.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadAllFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'.\n" +
      "Instead, the number of bytes read='%v'\n",
      expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if fMgr.filePtr != nil {
    t.Error("Error: Expected fMgr.filePtr == nil.\n" +
      "fMgr.filePtr IS NOT NIL!\n")
  }
}

func TestFileMgr_ReadAllFile_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead857268.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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

}

func TestFileMgr_ReadAllFile_03(t *testing.T) {

  expectedBytes := int(155)

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr" +
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
    t.Errorf("Error returned by fMgr.ReadAllFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() ,err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  lenBytesRead := len(bytesRead)

  if expectedBytes != lenBytesRead {
    t.Errorf("Error: Expected number of bytes read='%v'.\n" +
      "Instead, the number of bytes read='%v'\n",
      expectedBytes, lenBytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_ReadFileLine_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  delim := byte('\n')

  bytes, err := fMgr.ReadFileLine(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr." +
      "ReadFileLine(delim) on Line#1.\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() ,err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'",
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
    t.Errorf("Expected line #1 = 'Now is the time for " +
      "all good men'.\n" +
      "Instead, line #1 = '%v'",
      actualStr)
  }

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_ReadFileLine_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
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
    t.Errorf("Test Clean-Up Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_ReadFileLine_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
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
      t.Errorf("Error returned by fMgr.ReadFileLine(delim) on " +
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
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\n Error='%v'",
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
    t.Errorf("Expected line #4 = 'Thank you, for your support.'\n" +
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_ReadFileLine_04(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName() , err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      t.Errorf("Error returned by fMgr.ReadFileLine" +
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
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'",
     fMgr.GetAbsolutePathFileName() ,err.Error())
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
    t.Errorf("Expected line #4 = 'Thank you, for your support.'\n" +
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
  }
}

func TestFileMgr_ReadFileLine_05(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
    t.Errorf("Error returned by fMgr.ReadFileLine(delim) " +
      "on Line#1.\n"+
      "fMgr='%v'\nError='%v'",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  if "Now is the time for all good men" != actualStr {
    t.Errorf("Expected line #1= 'Now is the time for all good men'.\n" +
      "Instead, line #1 = '%v'\n", actualStr)
    return
  }

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
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

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileBytes(byteBuff)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
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
    t.Errorf("Expected Read String='%v'.\n" +
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.\n" +
      "Instead, Actual Bytes Read='%v'\n",
      expectedBytesRead, bytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile()\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  if fMgr.filePtr != nil {
    t.Error("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
  }

}

func TestFileMgr_ReadFileBytes_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
}

func TestFileMgr_ReadFileBytes_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileBytes(byteBuff)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
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
    t.Errorf("Expected Read String='%v'.\n" +
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    t.Errorf("Expected Bytes Read='%v'.\n" +
      "Instead, Actual Bytes Read='%v'\n",
      expectedBytesRead, bytesRead)
  }

  _ = fMgr.CloseThisFile()

}

func TestFileMgr_ReadFileString_01(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  delim := byte('\n')

  actualStr, err := fMgr.ReadFileString(delim)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileString(delim) " +
      "on Line#1.\n"+
      "fMgr='%v'\nError='%v'",
      fMgr.GetAbsolutePathFileName() ,err.Error())
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

}

func TestFileMgr_ReadFileString_02(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n" +
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
      t.Errorf("Error returned by fMgr.ReadFileString(delim) " +
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
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  if fMgr.filePtr != nil {
    t.Errorf("ERROR: After fMgr.CloseThisFile(), " +
      "expected fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n"+
      "fileMgr='%v'\n",
      fMgr.GetAbsolutePathFileName())
    _ = fMgr.CloseThisFile()
    return
  }

  if "Thank you, for your support." != actualStr {
    t.Errorf("Expected line #4= 'Thank you, for your support.'\n" +
      "Instead, line #4 = '%v'\n", actualStr)
  }

  if !isErrEOF {
    t.Error("ERROR: Expected the last error return from fMgr.ReadFileLine(delim) " +
      "to be io.EOF. Instead, error WAS NOT equal to io.EOF!")
  }

}

func TestFileMgr_ReadFileString_03(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
    t.Errorf("Error returned by fMgr.ReadFileString(" +
      "delim) on Line#1.\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
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
    t.Errorf("Expected line #1 = '%v'.\n" +
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }
}

func TestFileMgr_ReadFileString_04(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
    t.Errorf("Error returned by fMgr.ReadFileString(delim) " +
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

}

func TestFileMgr_ReadFileString_05(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
    t.Errorf("Error returned by fMgr.ReadFileString(" +
      "delim) on Line#1.\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
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
    t.Errorf("Expected line #1 = '%v'.\n" +
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }

}

func TestFileMgr_ReadFileString_06(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}." +
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

}

func TestFileMgr_ReadFileString_07(t *testing.T) {

  expectedStr := "Now is the time for all good men"

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
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
    t.Errorf("Error returned by fMgr.ReadFileString(" +
      "delim) on Line#1.\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n" +
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
    t.Errorf("Expected line #1 = '%v'.\n" +
      "Instead, line #1 = '%v'\n",
      expectedStr, actualStr)
  }

}

func TestFileMgr_ResetFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.ResetFileInfo()

  if err != nil {
    t.Errorf("Error returned by fMgr.ResetFileInfo().\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  fInfoPlus, err := fMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned by fMgr.GetFileInfoPlus().\n" +
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
}
