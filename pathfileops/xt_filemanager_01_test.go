package pathfileops

import (
  "fmt"
  "testing"
)

func TestFileMgr_ChangePermissionMode_01(t *testing.T) {

  filePath := "../filesfortest/modefilesfortest/modeFileTest_01.txt"

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  basePermission, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rw-rw-rw-\"). "+
      "Error='%v'", err.Error())
  }

  basePermissionText, _ := basePermission.GetPermissionTextCode()

  err = fMgr.ChangePermissionMode(basePermission)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(basePermission). "+
      "basePermission='%v' Error='%v'", basePermissionText, err.Error())
  }

  requestedNewPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\"). "+
      "Error='%v'", err.Error())
  }

  requestedNewPermText, err := requestedNewPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned from requestedNewPerm.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  err = fMgr.ChangePermissionMode(requestedNewPerm)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(requestedNewPerm). "+
      "Error='%v'", err.Error())
  }

  actualNewPermCodeText, err := fMgr.GetFilePermissionTextCodes()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.GetFilePermissionTextCodes(). "+
      "Error='%v'", err.Error())
  }

  if requestedNewPermText != actualNewPermCodeText {
    t.Errorf("Error expected permission='%v'. Instead, permission='%v' ",
      requestedNewPermText, actualNewPermCodeText)
  }

  err = fMgr.ChangePermissionMode(basePermission)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(basePermission). "+
      "basePermission='%v' Error='%v'", basePermissionText, err.Error())
  }

}

func TestFileMgr_ChangePermissionMode_02(t *testing.T) {

  filePath := "../filesfortest/modefilesfortest/iDoNotExist.txt"

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  basePermission, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rw-rw-rw-\"). "+
      "Error='%v'", err.Error())
  }

  err = fMgr.ChangePermissionMode(basePermission)

  if err == nil {
    t.Errorf(" Expected error return from fMgr.ChangePermissionMode(basePermission) " +
      "because file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_ChangePermissionMode_03(t *testing.T) {

  filePath := "../filesfortest/modefilesfortest/modeFileTest_01.txt"

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  basePermission, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rw-rw-rw-\"). "+
      "Error='%v'", err.Error())
  }

  basePermission.isInitialized = false

  err = fMgr.ChangePermissionMode(basePermission)

  if err == nil {
    t.Errorf(" Expected error return from fMgr.ChangePermissionMode(basePermission) " +
      "because file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_ChangePermissionMode_04(t *testing.T) {

  filePath := "../filesfortest/modefilesfortest/modeFileTest_01.txt"

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  basePermission, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-rw-rw-rw-\"). "+
      "Error='%v'", err.Error())
  }

  fMgr.isInitialized = false

  err = fMgr.ChangePermissionMode(basePermission)

  if err == nil {
    t.Errorf(" Expected error return from fMgr.ChangePermissionMode(basePermission) " +
      "because file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CloseThisFile_01(t *testing.T) {
  fh := FileHelper{}

  testFile := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Error("Received Error on GetPathFileNameElements Error:", err)
  }

  fileMgr.isInitialized = false

  err = fileMgr.CloseThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.CloseThisFile() because " +
      "fileMgr is Invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CloseThisFile_02(t *testing.T) {
  fh := FileHelper{}

  testFile := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Error("Received Error on GetPathFileNameElements Error:", err)
  }

  err = fileMgr.CloseThisFile()

  if err != nil {
    t.Error("Error: File Pointer is 'nil' and NO ERROR should have been returned!")
  }

}

func TestFileMgr_CopyIn_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
  fileName := "xt_dirmgr_01_test"
  fileNameExt := "xt_dirmgr_01_test.go"
  extName := ".go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

  if err != nil {
    t.Error("Received Error on GetPathFileNameElements Error:", err)
  }

  fMgr2 := FileMgr{}

  fMgr2.CopyIn(&fileMgr)

  if fMgr2.fileName != fileName {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileName == '%v', instead got: ", fileName), fMgr2.fileName)
  }

  if fMgr2.fileExt != extName {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileExt == '%v', instead got: ", extName), fMgr2.fileExt)
  }

  if fMgr2.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileNameExt == '%v', instead got: ", fileNameExt), fMgr2.fileNameExt)
  }

  if fMgr2.dMgr.path != expectedDir {
    t.Error(fmt.Sprintf("Expected CopyToThis to return path == '%v', instead got: ", expectedDir), fMgr2.dMgr.path)
  }

  result := fMgr2.Equal(&fileMgr)

  if result != true {
    t.Error("Expected Equal to return 'true' for fMgr2==fileMgr, instead got: ", result)
  }

}

func TestFileMgr_CopyOut_01(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
  fileName := "xt_dirmgr_01_test"
  fileNameExt := "xt_dirmgr_01_test.go"
  extName := ".go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

  if err != nil {
    t.Error("Received Error on GetPathFileNameElements Error:", err)
  }

  fMgr2 := fileMgr.CopyOut()

  if fMgr2.fileName != fileName {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileName == '%v', instead got: ", fileName), fMgr2.fileName)
  }

  if fMgr2.fileExt != extName {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileExt == '%v', instead got: ", extName), fMgr2.fileExt)
  }

  if fMgr2.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected CopyToThis to return fileNameExt == '%v', instead got: ", fileNameExt), fMgr2.fileNameExt)
  }

  if fMgr2.dMgr.path != expectedDir {
    t.Error(fmt.Sprintf("Expected CopyToThis to return path == '%v', instead got: ", expectedDir), fMgr2.dMgr.path)
  }

  result := fMgr2.Equal(&fileMgr)

  if result != true {
    t.Error("Expected Equal to return 'true' for fMgr2==fileMgr, instead got: ", result)
  }

}

func TestFileMgr_CopyFileMgrByIo_01(t *testing.T) {

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

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). "+
      "destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByIo(&destFMgr). "+
      "srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'",
      srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) "+
      "destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
  }

  if fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. "+
      "Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

}

func TestFileMgr_CopyFileMgrByIo_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIo_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  destFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr) because " +
      "destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIo(nil)

  if err == nil {
    t.Error("Expected error return from CopyFileMgrByIo(nil) because " +
      "nil was passed to method. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIo_05(t *testing.T) {

  expectedFileNameExt := "iDoNotExist.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    t.Error("Expected error return from CopyFileMgrByIo(&destFMgr) because " +
      "source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIo_06(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  destFMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    t.Error("Expected error return from CopyFileMgrByIo(&destFMgr) because " +
      "source file is equivalent to destination file. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByIoByLink(&destFMgr). srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'", srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) "+
      "destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
  }

  if fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. "+
      "Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByIoByLink(&destFMgr). srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'", srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
  }

  if fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  destFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_05(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIoByLink(nil)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "destFMgr is 'nil'. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_06(t *testing.T) {

  expectedFileNameExt := "iDoNotExist.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
  }

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "srcFMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_07(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  destFMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "srcFMgr is equivalent to destFMgr. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLink_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByLink(&newFileMgr). "+
      "newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileMgrByLink_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

  }
}

func TestFileMgr_CopyFileMgrByLink_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  newFileMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
      "because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

  }
}

func TestFileMgr_CopyFileMgrByLink_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  err = srcFMgr.CopyFileMgrByLink(nil)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(nil) " +
      "because nil was passed to this method. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLink_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(nil) " +
      "because srcFMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLink_06(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  destFileMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByLink(&destFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&destFileMgr) " +
      "because srcFMgr and destFileMgr are equivalent. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLinkByIo_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr). "+
      "newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  newFileMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  err = srcFMgr.CopyFileMgrByLinkByIo(nil)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(nil) " +
      "because nil was passed to the method. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByIo_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByIo(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
      "Destination File='%v'",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
      "Destination File='%v' ", rawAbsDestPath)
  }

}

func TestFileMgr_CopyFileStrByIo_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByIo_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByIo("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByIo_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist := fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}
