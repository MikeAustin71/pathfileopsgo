package pathfileops

import (
  "io/ioutil"
  "testing"
)

func TestFileMgr_OpenThisFile_01(t *testing.T) {
  fh := FileHelper{}

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\n"+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error retruned by fMgr.CreateThisFile()\n"+
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

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
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
  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_OpenThisFileReadOnly_01(t *testing.T) {
  fh := FileHelper{}

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  b, err := ioutil.ReadAll(fMgr.filePtr)

  if err != nil {
    _ = fMgr.CloseThisFile()
    t.Errorf("Error returned from ioutil.ReadAll(fMgr.filePtr).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {
    t.Errorf("Expected Read String='%v'.\n"+
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
    return
  }

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_OpenThisFileReadOnly_02(t *testing.T) {

  fh := FileHelper{}

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

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

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()
}

func TestFileMgr_OpenThisFileReadOnly_03(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/TestFileMgr_OpenThisFileReadOnly_03.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(filePath)\n"+
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
    t.Errorf("ERROR: Test file should NOT exist!.\n"+
      "However, test file DOES EXIST!\n"+
      "test file='%v'", filePath)
    _ = fh.DeleteDirFile(filePath)
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile().\n"+
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
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr"+
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

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileReadWrite()

  if err != nil {

    t.Errorf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  b, err := ioutil.ReadAll(fMgr.GetFilePtr())

  if err != nil {

    t.Errorf("Error returned from ioutil.ReadAll(fMgr.GetFilePtr())\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())

    _ = fMgr.CloseThisFile()
    return
  }

  actualStr := string(b)

  expectedStr := "Test Read File. Do NOT alter the contents of this file."

  if expectedStr != actualStr {

    t.Errorf("Expected Read String='%v'.\n"+
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from final fMgr.CloseThisFile(). Error='%v' ",
      err.Error())
  }
}

func TestFileMgr_OpenThisFileReadWrite_02(t *testing.T) {
  fh := FileHelper{}

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
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

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n"+
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

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n"+
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
    t.Errorf("Non-Path Error returned from fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if !fileDoesExist {
    t.Errorf("Error: The test file should have been created.\n"+
      "However, the test file does NOT exist!\n"+
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
    t.Errorf("Expected at least 12-bytes to be written to file.\n"+
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

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnlyAppend().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite = []byte(testText2)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead1, err := fMgr.ReadFileLine('\n')

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.ReadFileLine('newline')\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead2, err := fMgr.ReadFileLine('\n')

  if err != nil {
    t.Errorf("Error returned from #2 fMgr.ReadFileLine('newline')\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED!\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  stringRead := string(bytesRead1)

  stringRead = stringRead[:len(stringRead)-1]

  testText1 = testText1[:len(testText1)-1]

  if testText1 != stringRead {
    t.Errorf("Error: Expected #1 stringRead='%v'.\n"+
      "Instead, #1 stringRead='%v'\n",
      testText1, stringRead)
  }

  stringRead = string(bytesRead2)

  stringRead = stringRead[:len(stringRead)-1]

  testText2 = testText2[:len(testText1)-1]

  if testText2 != stringRead {
    t.Errorf("Error: Expected #2 stringRead='%v'.\n"+
      "Instead, #2 stringRead='%v'\n",
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

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
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

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if !fileDoesExist {
    t.Errorf("Error: Expected target file to be created.\n"+
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

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
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
    t.Errorf("Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if fileDoesExist {
    t.Errorf("Error: Expected target file WOULD NOT be created.\n"+
      "However, target file DOES EXIST!\n"+
      "target file='%v'", fMgr.GetAbsolutePathFileName())
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()
}
