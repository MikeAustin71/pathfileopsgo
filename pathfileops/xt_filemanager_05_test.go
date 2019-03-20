package pathfileops

import (
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "strings"
  "testing"
)

func TestFileMgr_New_01(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\logTest\\CmdrX\\CmdrX.log"
  commonDir, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
      "relPath='%v'  Error='%v'", relPath, err.Error())
  }

  fileName := "CmdrX"
  fileNameExt := "CmdrX.log"
  extName := ".log"

  fileMgr, err := FileMgr{}.New(commonDir)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_New_02(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("..\\logTest\\CmdrX\\CmdrX.log")

  fileName := "CmdrX"
  fileNameExt := "CmdrX.log"
  extName := ".log"

  fileMgr, err := FileMgr{}.New(commonDir)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_New_03(t *testing.T) {

  _, err := FileMgr{}.New("")

  if err == nil {
    t.Error("Expected error return from FileMgr{}.New(\"\") because " +
      "the input parameter is an empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_New_04(t *testing.T) {

  _, err := FileMgr{}.New("!^%&*()")

  if err == nil {
    t.Error("Expected error return from FileMgr{}.New(\"!^%&*()\") because " +
      "the input parameter contains invalid characters. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_NewFromPathFileNameExtStr_01(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")
  fileName := "xt_dirmgr_01_test"
  fileNameExt := "xt_dirmgr_01_test.go"
  extName := ".go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

  if err != nil {
    t.Error("Received Error on GetPathFileNameElements Error:", err)
  }

  if fileMgr.fileName != fileName {
    t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
  }

}

func TestFileMgr_NewFromPathFileNameExtStr_02(t *testing.T) {

  path := "../appExamples/filehelperexamples.go"

  eFileNameExt := "filehelperexamples.go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
      "path== '%v' Error: %v", path, err)
  }

  if eFileNameExt != fileMgr.fileNameExt {
    t.Errorf("Expected extracted fileNameExt == %v, instead got: %v",
      eFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != "filehelperexamples" {
    t.Errorf("Expected fileMgr.fileName== 'filehelperexamples', "+
      "instead got: fileMgr.fileName== %v", fileMgr.fileName)
  }

  if fileMgr.fileExt != ".go" {
    t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
      fileMgr.fileExt)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated==%v",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
      "fileMgr.doesAbsolutePathFileNameExist==%v", fileMgr.doesAbsolutePathFileNameExist)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
  }

}

func TestFileMgr_NewFromPathFileNameExtStr_03(t *testing.T) {

  path := "filehelperexamples"

  _, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err == nil {
    t.Error("Expected an error from FileMgr{}.NewFromPathFileNameExtStr(path) " +
      "because path='filehelperexamples'. However, NO ERROR WAS RETURNED! ")
  }
}

func TestFileMgr_NewFromPathFileNameExtStr_04(t *testing.T) {

  path := "../appExamples/filehelperexamples.go"

  eFileNameExt := "filehelperexamples.go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
      "path=='%v' Error: %v ", path, err)
  }

  if eFileNameExt != fileMgr.fileNameExt {
    t.Errorf("Expected extracted fileNameExt == '%v', instead got: '%v' ",
      eFileNameExt, fileMgr.fileNameExt)
  }

  if "filehelperexamples" != fileMgr.fileName {
    t.Errorf("Expected fileMgr.fileName== '%v', instead got: fileMgr.fileName== '%v'",
      "filehelperexamples", fileMgr.fileName)
  }

  if ".go" != fileMgr.fileExt {
    t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
      fileMgr.fileExt)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.dMgr.isPathPopulated==true', instead got: "+
      "fileMgr.isPathPopulated==%v",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
      "fileMgr.doesAbsolutePathFileNameExist== %v", fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Errorf("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, "+
      "fileMgr.isAbsolutePathFileNamePopulated == '%v' ", fileMgr.isAbsolutePathFileNamePopulated)
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Errorf("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is '%v'",
      fileMgr.dMgr.doesAbsolutePathExist)
  }

}

func TestFileMgr_NewFromFileInfo_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"
  expectedFileName := "newerFileForTest_01"
  expectedExt := ".txt"
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

  fileMgr, err := FileMgr{}.NewFromFileInfo(absPath, info)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info). absPath='%v' info.Name()='%v'  Error='%v'", absPath, info.Name(), err.Error())
  }

  if fileMgr.fileNameExt != expectedFileNameExt {
    t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != expectedFileName {
    t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != expectedExt {
    t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.doesAbsolutePathFileNameExist)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
  }

  if !fileMgr.actualFileInfo.IsFInfoInitialized {
    t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
  }

  if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
  }

}

func TestFileMgr_NewFromFileInfo_02(t *testing.T) {

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  var info os.FileInfo

  _, err = FileMgr{}.NewFromFileInfo(absPath, info)

  if err == nil {
    t.Errorf("Expected an error from FileMgr{}.NewFromFileInfo(absPath, info) because " +
      "input parameter 'info' is INVALID!  However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromFileInfo_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). "+
      "absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  absPath = "../iDoNotExist"

  _, err = FileMgr{}.NewFromFileInfo(absPath, info)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info). "+
      "absPath='%v' info.Name()='%v'  Error='%v'", absPath, info.Name(), err.Error())
  }

}

func TestFileMgr_NewFromFileInfo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
  }

  absPath = ""

  _, err = FileMgr{}.NewFromFileInfo(absPath, info)

  if err == nil {
    t.Error("Expected an error from FileMgr{}.NewFromFileInfo(absPath, info) because " +
      "absPath is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
      "dMgr='%v' rawFileNameExt='%v' Error='%v' ",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
    t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_02(t *testing.T) {

  rawFileNameExt := "./newerFileForTest_01.txt"
  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt). "+
      "dMgr='%v' rawFileNameExt='%v' Error='%v' ",
      dMgr.GetAbsolutePath(), rawFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
    t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  rawPath := "../filesfortest/newfilesfortest"

  absolutePath, err := fh.MakeAbsolutePath(rawPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawPath). rawPath='%v' Error='%v'",
      rawPath, err.Error())
  }

  dMgr, err := DirMgr{}.New(absolutePath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(absolutePath). "+
      "adjustedPath='%v'  Error='%v'", absolutePath, err.Error())
  }

  _, err = FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "")

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, \"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  dMgr.isInitialized = false

  _, err = FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}" +
      "NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt) " +
      "because the dMgr is INVALID. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_05(t *testing.T) {

  expectedFileNameExt := "$%!*().#+_"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  _, err = FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}" +
      "NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt) " +
      "because the expectedFileNameExt contains invalid characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"
  expectedFileName := "newerFileForTest_01"
  expectedExt := ".txt"

  fh := FileHelper{}
  rawPath := "../filesfortest/newfilesfortest"
  expectedPath := fh.AdjustPathSlash(rawPath)
  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
  }

  fileMgr, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt). rawPath='%v' expectedFileNameExt='%v'  Error='%v'", rawPath, expectedFileNameExt, err.Error())
  }

  if fileMgr.fileNameExt != expectedFileNameExt {
    t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != expectedFileName {
    t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != expectedExt {
    t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
  }

  if !fileMgr.actualFileInfo.IsFInfoInitialized {
    t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
  }

  if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
  }

  if expectedAbsPath != fileMgr.dMgr.absolutePath {
    t.Errorf("Expected absolutePath='%v'.  Instead, absolutePath='%v' ", expectedAbsPath, fileMgr.dMgr.absolutePath)
  }

  if expectedPath != fileMgr.dMgr.path {
    t.Errorf("Expected path='%v'.  Instead, path='%v' ", expectedPath, fileMgr.dMgr.path)
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := ""

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "rawPath is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_03(t *testing.T) {

  expectedFileNameExt := ""

  rawPath := "../filesfortest/newfilesfortest"

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "expectedFileNameExt is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_04(t *testing.T) {

  expectedFileNameExt := "$!#%^&*()_=.t%t"

  rawPath := "../filesfortest/newfilesfortest"

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "expectedFileNameExt consists of invalid characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_05(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := ""

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "raw path is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

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
    t.Errorf("Error = Expected fmgr.actualFileInfo.Name()='%v'.  "+
      "Instead, fmgr.actualFileInfo.Name()='%v'",
      expectedFileNameExt, fmgr.actualFileInfo.Name())
  }

}

func TestFileMgr_WriteBytesToFile_01(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

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

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    t.Errorf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    t.Errorf("Error returned by fMgr.FlushBytesToDisk(). Error='%v' ",
      err.Error())
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    t.Errorf("Error returned by fMgr.ReadFileBytes(bytesRead). Error='%v'",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #2 fMgr.CloseThisFile().")
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    t.Errorf("Error: Expected stringRead='%v'. Instead, stringRead='%v' ",
      testText, stringRead)
  }

  if verifyBytesWritten != uint64(lenTestText) {
    t.Errorf("Error: verifyBytesWritten != lenTestText. verifyBytesWritten='%v' "+
      "lenTestText='%v' ", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    t.Errorf("Error: numBytesRead != lenTestText. numBytesRead='%v' "+
      "lenTestText='%v' ", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    t.Errorf("Error: numBytesRead != numBytesWritten. numBytesRead='%v' "+
      "numBytesWritten='%v' ", numBytesRead, numBytesWritten)

  }

}

func TestFileMgr_WriteBytesToFile_02(t *testing.T) {

  fh := FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWriteXX241289.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by #1 fMgr.CloseThisFile().")
  }

  bytesToWrite := []byte(testText)

  fMgr.isInitialized = false

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err == nil {
    t.Error("Expected an error from fMgr.WriteBytesToFile(bytesToWrite) " +
      "because fMgr.isInitialized == false. Instead, NO ERROR WAS RETURNED!")
  }

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fMgr2.DeleteThisFile(). Error='%v' ",
      err.Error())
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

func TestFileMgr_WriteStrToFile_02(t *testing.T) {
  fh := FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWrite2998.txt")

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  fMgr2 := fMgr.CopyOut()

  err = fMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CreateThisFile(). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned from fMgr.CloseThisFile() No 1.  Error='%v'", err.Error())
  }

  expectedStr := "Test Write File. Do NOT alter the contents of this file."

  fMgr.isInitialized = false

  _, err = fMgr.WriteStrToFile(expectedStr)

  if err == nil {
    t.Error("Expected an error to be returned from fMgr.WriteStrToFile(expectedStr) " +
      "because fMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

  err = fMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error Deleting File: %v. Error returned by fMgr2.DeleteThisFile(). "+
      "Error='%v'", fMgr2.GetAbsolutePathFileName(), err.Error())
  }

}
