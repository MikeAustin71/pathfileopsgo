package pathfileops

import (
  "fmt"
  "testing"
)

func TestFileMgr_ChangePermissionMode_01(t *testing.T) {

  filePath := "../../filesfortest/modefilesfortest/modeFileTest_01.txt"

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
    return
  }

  basePermissionText, _ := basePermission.GetPermissionTextCode()

  err = fMgr.ChangePermissionMode(basePermission)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(basePermission). "+
      "basePermission='%v' Error='%v'", basePermissionText, err.Error())
    return
  }

  requestedNewPerm, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\"). "+
      "Error='%v'", err.Error())
    return
  }

  requestedNewPermText, err := requestedNewPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned from requestedNewPerm.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgr.ChangePermissionMode(requestedNewPerm)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(requestedNewPerm). "+
      "Error='%v'", err.Error())
    return
  }

  actualNewPermCodeText, err := fMgr.GetFilePermissionTextCodes()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.GetFilePermissionTextCodes(). "+
      "Error='%v'", err.Error())
    return
  }

  if requestedNewPermText != actualNewPermCodeText {
    t.Errorf("Error expected permission='%v'. Instead, permission='%v' ",
      requestedNewPermText, actualNewPermCodeText)
  }

  err = fMgr.ChangePermissionMode(basePermission)

  if err != nil {
    t.Errorf("Error returned from fMgr.ChangePermissionMode(basePermission). "+
      "basePermission='%v' Error='%v'", basePermissionText, err.Error())
    return
  }

}

func TestFileMgr_ChangePermissionMode_02(t *testing.T) {

  filePath := "../../filesfortest/modefilesfortest/iDoNotExist.txt"

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

  filePath := "../../filesfortest/modefilesfortest/modeFileTest_01.txt"

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
    return
  }

  basePermission.isInitialized = false

  err = fMgr.ChangePermissionMode(basePermission)

  if err == nil {
    t.Errorf(" Expected error return from fMgr.ChangePermissionMode(basePermission) " +
      "because file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  return
}

func TestFileMgr_ChangePermissionMode_04(t *testing.T) {

  filePath := "../../filesfortest/modefilesfortest/iDoNotExist.txt"

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
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

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

  rawDestPath := "../../checkfiles/checkfiles02"

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
    return
  }

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByIo(&destFMgr). "+
      "srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'",
      srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
    return
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true.\n" +
      "Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.\n" +
      "ERROR:  destFMgr.doesAbsolutePathFileNameExist='false'")
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt). dMgr.absolutePath='%v'\n"+
      "expectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "destDMgr, expectedFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr("+
      "rawDestPath).\nrawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "+
      "expectedFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
    return
  }

  destFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr) because " +
      "destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
  }

  destFMgr.isInitialized = true

  err = destFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by destFMgr.DeleteThisFile().\n"+
      "destFMgr='%v'\nError='%v'\n", destFMgr.GetAbsolutePath(), err.Error())
  }
}

func TestFileMgr_CopyFileMgrByIo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\nadjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "+
      "expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\nadjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(rawDestPath).\nrawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(destDMgr, expectedFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}."+
      "NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "+
      "expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../../checkfiles/checkfiles02"

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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../../checkfiles/checkfiles02"

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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../../checkfiles/checkfiles02"

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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'  Error='%v'",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "+
      "expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "+
      "expectedFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
    return
  }

  destFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
  }

  destFMgr.isInitialized = true

  err = destFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by destFMgr.DeleteThisFile().\n"+
      "destFMgr='%v'\nError='%v'\n", destFMgr.GetAbsolutePath(), err.Error())
  }

}

func TestFileMgr_CopyFileMgrByIoByLink_05(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())

    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "+
      "expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt("+
      "destDMgr, expectedFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedFileNameExt, err.Error())
    return
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
  adjustedPath := fh.AdjustPathSlash("../../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  destFMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

  if err == nil {
    t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
      "srcFMgr is equivalent to destFMgr. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByIoWithBuffer_01(t *testing.T) {

  setupFile := "../../filesfortest/levelfilesfortest/level_0_3_test.txt"

  fh := FileHelper{}

  setupFile = fh.AdjustPathSlash(setupFile)

  sourceFile := "../../createFilesTest/level_0_3_test.txt"

  sourceFile = fh.AdjustPathSlash(sourceFile)

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n"+
      "setupFile='%v'\nsourceFile='%v'\nError='%v'\n",
      setupFile, sourceFile, err.Error())
    return
  }

  destFile := "../../createFilesTest/TestFileMgr_CopyFileMgrByIoWithBuffer_01.txt"

  destFile = fh.AdjustPathSlash(destFile)

  srcFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile).\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  err = srcFMgr.CopyFileMgrByIoWithBuffer(&destFMgr, 15000)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileMgrByIoWithBuffer(&destFMgr, 15000)\n"+
      "srcFMgr='%v'\nError='%v'\n", srcFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: After Copy Operation, the destination file DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
  }

  _ = fh.DeleteDirFile(sourceFile)
  _ = fh.DeleteDirFile(destFile)
  return
}

func TestFileMgr_CopyFileMgrByIoWithBuffer_02(t *testing.T) {

  fh := FileHelper{}

  sourceDir := "../../iDoNotExist"

  sourceDir = fh.AdjustPathSlash(sourceDir)

  sourceFile := "../../iDoNotExist/iDoNotExist.txt"

  sourceFile = fh.AdjustPathSlash(sourceFile)

  destFile := "../../createFilesTest/TestFileMgr_CopyFileMgrByIoWithBuffer_02.txt"

  destFile = fh.AdjustPathSlash(destFile)

  srcFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile).\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  err = srcFMgr.CopyFileMgrByIoWithBuffer(&destFMgr, 15000)

  if err == nil {
    t.Errorf("Expected an error returned by srcFMgr.CopyFileMgrByIoWithBuffer(&destFMgr, 15000)\n"+
      "because srcFMgr DOES NOT EXIST!\n"+
      "However, NO ERROR WAS RETURNED!\nsrcFMgr='%v'\n", srcFMgr.absolutePathFileName)
  }

  _ = fh.DeleteDirPathAll(sourceDir)
  _ = fh.DeleteDirFile(destFile)
  return
}

