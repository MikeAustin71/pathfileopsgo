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
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n" +
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
    return
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
    return
  }

  err = fmgr.SetFileInfo(info)

  if err != nil {
    t.Errorf("Error returned by fmgr.SetFileInfo(info).\n" +
      "info.Name()='%v'\nError='%v'\n", info.Name(), err.Error())
  }

  if !fmgr.actualFileInfo.isFInfoInitialized {
    t.Error("Error - File Manager FileInfoPlus object is not initialized!\n")
  }

  if fmgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("Error = Expected fmgr.actualFileInfo.Name()='%v'.\n"+
      "Instead, fmgr.actualFileInfo.Name()='%v'\n",
      expectedFileNameExt, fmgr.actualFileInfo.Name())
  }

}

func TestFileMgr_SetFileInfo_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n" +
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n",
      absPathFileNameExt, err.Error())
    return
  }

  err = fmgr.SetFileInfo(nil)

  if err == nil {
    t.Error("Expected error return from fmgr.SetFileInfo(nil) because " +
      "input parameter fileInfo is 'nil'.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileInfo_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n" +
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
    return
  }

  err = fmgr.SetFileInfo(info)

  if err == nil {
    t.Error("Expected an error from fmgr.SetFileInfo(info) " +
      "because input parameter 'info'\n" +
      "is a directory an not a file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileInfo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  alternativeFileNameExt := "newerFileForTest_03.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n" +
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  absAlternatePathFileNameExt := absPath + string(os.PathSeparator) + alternativeFileNameExt

  info, err := fh.GetFileInfo(absAlternatePathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n" +
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n", absPathFileNameExt, err.Error())
    return
  }

  err = fmgr.SetFileInfo(info)

  if err == nil {
    t.Error("Expected an error from fmgr.SetFileInfo(info) because " +
      "input parameter 'info'\n" +
      "is a file name which is different from fmgr file name.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr.SetFileMgrFromDirMgrFileName(" +
      "dMgr, expectedFileNameExt) " +
      "isEmpty=='false'.\n" +
      "Instead, isEmpty=='true'.\n")
    return
  }

  if err != nil {
    t.Errorf("Error returned from fMgr.SetFileMgrFromDirMgrFileName(" +
      "dMgr, expectedFileNameExt).\n"+
      "dMgr='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
    return
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.\n" +
      "Instead, absolutePathFileName='%v'\n",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_02(t *testing.T) {

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
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
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  fMgr := FileMgr{}

  _, err = fMgr.SetFileMgrFromDirMgrFileName(dMgr, "    ")

  if err == nil {
    t.Error("Expected error return from fMgr.SetFileMgrFromDirMgrFileName(dMgr, \"  \")\n" +
      "because input parameter fileNameExt consists of blank spaces!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  fMgr := FileMgr{}

  dMgr.isInitialized = false

  _, err = fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from fMgr.SetFileMgrFromDirMgrFileName(dMgr, \"  \")\n" +
      "because input parameter dMgr is invalid!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_05(t *testing.T) {

  expectedFileNameExt := "basefilenoext"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr." +
      "SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)\n" +
      "isEmpty=='false'.\n" +
      "Instead, isEmpty=='true'.\n")
  }

  if err != nil {
    t.Errorf("Error returned from fMgr.SetFileMgrFromDirMgrFileName(" +
      "dMgr, expectedFileNameExt).\n"+
      "dMgr='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.\n" +
      "Instead, absolutePathFileName='%v'\n",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromDirMgrFileName_06(t *testing.T) {

  expectedFileNameExt := ".xgitignore"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  fMgr := FileMgr{}

  isEmpty, err := fMgr.SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt)

  if isEmpty {
    t.Error("Expected that after fMgr." +
      "SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt) " +
      "isEmpty=='false'.\n" +
      "Instead, isEmpty=='true'.\n")
  }

  if err != nil {
    t.Errorf("Error returned from fMgr." +
      "SetFileMgrFromDirMgrFileName(dMgr, expectedFileNameExt).\n"+
      "dMgr='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Expected absolutePathFileName='%v'.\n" +
      "Instead, absolutePathFileName='%v'\n",
      expectedAbsPathFileNameExt, fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_01(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\logTest\\CmdrX\\CmdrX.log"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath).\n"+
      "relPath='%v'\nError='%v'\n", relPath, err.Error())
    return
  }

  fileName := "CmdrX"
  fileNameExt := "CmdrX.log"
  extName := ".log"

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)\n"+
      "pathFileNameExt='%v'\nError='%v'\n",
      pathFileNameExt, err.Error())
    return
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)\n" +
      "expected isEmpty='false'.\n" +
      "Instead, isEmpty='false'.\n")
    return
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v\n." +
      "Instead, got:", fileNameExt),
      fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'.\n" +
      "Instead, got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'.\n" +
      "Instead got: ",
      fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'.\n" +
      "Instead, got:", fileMgr.isFileNameExtPopulated)
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true'\n" +
      "Instead, got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_02(t *testing.T) {

  fileMgr := FileMgr{}

  _, err := fileMgr.SetFileMgrFromPathFileName("")

  if err == nil {
    t.Error("Expected error return from fileMgr.SetFileMgrFromPathFileName(\"\")\n" +
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_03(t *testing.T) {

  fileMgr := FileMgr{}

  _, err := fileMgr.SetFileMgrFromPathFileName("      ")

  if err == nil {
    t.Error("Expected error return from fileMgr.SetFileMgrFromPathFileName(\"    \")\n" +
      "because the input parameter consists of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_04(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\filesfortest\\basefilesfortest\\basefilenoext"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath).\n"+
      "relPath='%v'\nError='%v'\n", relPath, err.Error())
    return
  }

  fileName := "basefilenoext"
  fileNameExt := "basefilenoext"
  extName := ""

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)\n"+
      "pathFileNameExt='%v'\nError='%v'\n",
      pathFileNameExt, err.Error())
    return
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt) expected " +
      "isEmpty='false'.\nInstead, isEmpty='false'.\n")
    return
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name='%v'\nInstead, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v\nInstead, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + " +
      "Extension= '%v'\nInstead, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'\nInstead, got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'\nInstead, got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'\n" +
      "Instead got:", fileMgr.isFileNameExtPopulated)
  }

  if fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='false'\n" +
      "Instead, got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'\n" +
      "Instead, got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_SetFileMgrFromPathFileName_05(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\filesfortest\\basefilesfortest\\.xgitignore"
  pathFileNameExt, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath).\n"+
      "relPath='%v'\nError='%v'\n", relPath, err.Error())
    return
  }

  fileName := ".xgitignore"
  fileNameExt := ".xgitignore"
  extName := ""

  fileMgr := FileMgr{}

  isEmpty, err := fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    t.Errorf("Received Error on fileMgr.SetFileMgrFromPathFileName(pathFileNameExt)\n"+
      "pathFileNameExt='%v'\nError='%v'\n",
      pathFileNameExt, err.Error())
    return
  }

  if isEmpty {
    t.Error("Error: after fileMgr.SetFileMgrFromPathFileName(pathFileNameExt) expected " +
      "isEmpty='false'.\n" +
      "Instead, isEmpty='false'.\n")
    return
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name= '%v'.\n" +
      "Instead, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension='%v'\n" +
      "Instead, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension = '%v'.\n" +
      "Instead, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'\n" +
      "Instead, got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'.\n" +
      "Instead, got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'.\n" +
      "Instead, got:", fileMgr.isFileNameExtPopulated)
  }

  if fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='false'.\n" +
      "Instead, got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, got:",
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
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile().\n" +
      "Error='%v'\n",
      err.Error())
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    t.Errorf("Error returned by fMgr.OpenThisFileWriteOnly().\n" +
      "Error='%v'\n",
      err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n" +
      "Error='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk().\nError='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead).\n" +
      "Error='%v'\n",
      err.Error())

    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED!\n" +
      "Error='%v'", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'.\n" +
      "Instead, stringRead='%v'\n",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText.\n" +
      "verifyBytesWritten='%v'\n"+
      "lenTestText='%v'\n", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText.\n" +
      "numBytesRead='%v'\n"+
      "lenTestText='%v'\n", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten.\n" +
      "numBytesRead='%v'\n"+
      "numBytesWritten='%v'\n", numBytesRead, numBytesWritten)
  }

}

func TestFileMgr_WriteBytesToFile_02(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWriteXX241289.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile().\n" +
      "Error='%v' ",
      err.Error())

    _ = fMgr.CloseThisFile()
    _ = fMgr2.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
  }

  bytesToWrite := []byte(testText)

  fMgr.isInitialized = false

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err == nil {
    t.Error("Expected an error from fMgr.WriteBytesToFile(bytesToWrite) " +
      "because fMgr.isInitialized == false.\n" +
      "Instead, NO ERROR WAS RETURNED!\n")
  }

  _ = fMgr2.CloseThisFile()

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr2.DeleteThisFile().\n" +
      "Error='%v'\n",
      err.Error())
  }
}

func TestFileMgr_WriteBytesToFile_03(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := fh.AdjustPathSlash("../checkfiles/TestFileMgr_WriteBytesToFile_03.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile().\n"+
      "filePathName='%v'\nError='%v'\n", fMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile().\n" +
      "Error='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CloseThisFile().\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if !fMgr.DoesFileExist() {
    t.Errorf("Error: Failed to create File Name:'%v'.\n" +
      "Terminating test.\n",
      fMgr.GetAbsolutePathFileName())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n" +
      "Error='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk().\n" +
      "Error='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead).\n" +
      "Error='%v'\n",
      err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #3 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED!\n" +
      "Error='%v'\n", err.Error())
    return
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'.\n" +
      "Instead, stringRead='%v'\n",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText.\n" +
      "verifyBytesWritten='%v'\n"+
      "lenTestText='%v'\n", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText.\n" +
      "numBytesRead='%v'\n"+
      "lenTestText='%v'\n", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten.\n" +
      "numBytesRead='%v'\n"+
      "numBytesWritten='%v'\n", numBytesRead, numBytesWritten)
  }

}

func TestFileMgr_WriteBytesToFile_04(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := fh.AdjustPathSlash("../checkfiles/TestFileMgr_WriteBytesToFile_04.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Setup Error returned from fMgr.DeleteThisFile().\n"+
      "filePathName='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk().\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead).\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #3 fMgr.CloseThisFile().\n" +
      "Error='%v'\n", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED!\nError='%v'\n",
      err.Error())
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'.\n" +
      "Instead, stringRead='%v'\n",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText.\n" +
      "verifyBytesWritten='%v'\n"+
      "lenTestText='%v'\n", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText.\n" +
      "numBytesRead='%v'\n"+
      "lenTestText='%v'\n", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten.\n" +
      "numBytesRead='%v'\n"+
      "numBytesWritten='%v'\n", numBytesRead, numBytesWritten)

  }

}

func TestFileMgr_WriteStrToFile_01(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/TestFileMgr_WriteStrToFile_01.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v\nError='%v'\n", filePath, err.Error())
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.DeleteThisFile()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  expectedStr := "Test Write File. Do NOT alter the contents of this file."

  lExpectedStr := len(expectedStr)

  bytesWritten, err := fMgr.WriteStrToFile(expectedStr)

  if err != nil {
    t.Errorf("Error returned from fMgr.WriteStrToFile(expectedStr)\n"+
      "expectedStr='%v'\nError='%v'\n", expectedStr, err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile().\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadAllFile().\n"+
      "filePathName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if lExpectedStr != bytesWritten {
    t.Errorf("Error: Length of string written NOT equal to Bytes Read!\n"+
      "Length of written string='%v'.\nActual Bytes Read='%v'\n",
      lExpectedStr, bytesWritten)
  }

  actualStr := string(bytesRead)

  if lExpectedStr != len(actualStr) {
    t.Errorf("Error: Length of actual string read is NOT equal to length "+
      "of string written.\n"+
      "lExpectedStr='%v'\nlen(actualStr)='%v'\n",
      lExpectedStr, len(actualStr))
  }

  if expectedStr != actualStr {
    t.Errorf("Error: expectedStr written='%v'\n"+
      "Actual string read='%v'\n",
      expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile()\n"+
      "Error='%v'\n", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.DeleteThisFile().\n"+
      "Error='%v'\n", err.Error())
  }

  doesFileExist := fh.DoesFileExist(filePath)

  if doesFileExist {
    t.Errorf("Error: Failed to DELETE fileNameExt='%v'\n",
      fMgr.absolutePathFileName)
  }

}

func TestFileMgr_WriteStrToFile_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/TestFileMgr_WriteStrToFile_02.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile().\n"+
      "filePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile()\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  expectedStr := "Test Write File. Do NOT alter the contents of this file."

  fMgr.isInitialized = false

  _, err = fMgr.WriteStrToFile(expectedStr)

  if err == nil {
    t.Error("Expected an error to be returned from fMgr.WriteStrToFile(expectedStr)\n" +
      "because fMgr.isInitialized == false.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  fMgr.isInitialized = true

  _ = fMgr.CloseThisFile()
  _ = fMgr.DeleteThisFile()

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error Deleting File: %v. Error returned by fMgr2.DeleteThisFile().\n"+
      "Error='%v'\n", fMgr2.GetAbsolutePathFileName(), err.Error())
  }

}

func TestFileMgr_WriteStrToFile_03(t *testing.T) {

  fh := FileHelper{}

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/TestFileMgr_WriteStrToFile_03.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePathName='%v\nError='%v'\n", filePath, err.Error())
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.DeleteThisFile()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile().\n"+
      "filePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile()\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  expectedStr := "Damn the torpedoes, full speed ahead!"

  lExpectedStr := len(expectedStr)

  bytesWritten, err := fMgr.WriteStrToFile(expectedStr)

  if err != nil {
    t.Errorf("Error returned from fMgr.WriteStrToFile(expectedStr)\n"+
      "expectedStr='%v'\nError='%v'\n", expectedStr, err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from #1 fMgr.CloseThisFile().\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead, err := fMgr.ReadAllFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.ReadAllFile().\n"+
      "filePathName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, err.Error())
    _ = fMgr.CloseThisFile()
    _ = fMgr.DeleteThisFile()
    return
  }

  if lExpectedStr != bytesWritten {
    t.Errorf("Error: Length of string written NOT equal to Bytes Read!\n"+
      "Length of written string='%v'.\nActual Bytes Read='%v'\n",
      lExpectedStr, bytesWritten)
  }

  actualStr := string(bytesRead)

  if lExpectedStr != len(actualStr) {
    t.Errorf("Error: Length of actual string read is NOT equal to length "+
      "of string written.\n"+
      "lExpectedStr='%v'\nlen(actualStr)='%v'\n",
      lExpectedStr, len(actualStr))
  }

  if expectedStr != actualStr {
    t.Errorf("Error: expectedStr written='%v'\n"+
      "Actual string read='%v'\n",
      expectedStr, actualStr)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile()\n"+
      "Error='%v'\n", err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from #2 fMgr.DeleteThisFile().\n"+
      "Error='%v'\n", err.Error())
  }

  doesFileExist := fh.DoesFileExist(filePath)

  if doesFileExist {
    t.Errorf("Error: Failed to DELETE fileNameExt='%v'\n",
      fMgr.absolutePathFileName)
  }

}
