package pathfileops

import (
  "fmt"
  "testing"
)

func TestFileMgr_CopyFileToDirByLink_01(t *testing.T) {

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
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
        "Error='%v'", err.Error())
    }

  }

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByLink(destDMgr). "+
      "destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
  }

  fileExists, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if !fileExists {
    t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileToDirByLink_02(t *testing.T) {

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
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
        "Error='%v'", err.Error())
    }

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByLink(destDMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByLink_03(t *testing.T) {

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
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
        "Error='%v'", err.Error())
    }

  }

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr) because " +
      "destDMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByLink_04(t *testing.T) {

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

  destDMgr := srcFMgr.GetDirMgr()

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr) because " +
      "source directory equals destination directory. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByLink_05(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr) because " +
      "source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CreateDir_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {

    err = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    if err != nil {
      t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath). "+
        " fileMgr.dMgr.absolutePath='%v'   Error='%v' ", fileMgr.dMgr.absolutePath, err.Error())
    }

  }

  if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {
    t.Errorf("Error: Failed to delete existing path '%v'",
      fileMgr.dMgr.absolutePath)
  }

  err = fileMgr.CreateDir()

  if err != nil {
    t.Errorf("Error returned from fileMgr.CreateDir(). Error='%v'",
      err.Error())
  }

  dirMgr := fileMgr.GetDirMgr()

  if !fh.DoesFileExist(dirMgr.GetAbsolutePath()) {
    t.Errorf("Error: Failed to create directory path '%v'. "+
      "PATH DOES NOT EXIST!",
      dirMgr.GetAbsolutePath())
  } else {

    err = dirMgr.DeleteAll()

    if err != nil {
      t.Errorf("Error returned from dirMgr.DeleteAll(). Error='%v'",
        err.Error())
    }

    if fh.DoesFileExist(dirMgr.GetAbsolutePath()) {
      t.Errorf("ERROR: Final Deletion of Directory Path FAILED! "+
        "File Manager Directory Path='%v' ", dirMgr.GetAbsolutePath())
    }
  }
}

func TestFileMgr_CreateDir_02(t *testing.T) {

  fileMgr := FileMgr{}

  err := fileMgr.CreateDir()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateDir() because " +
      "File Manager was NOT initialized. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CreateDirAndFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {

    err = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    if err != nil {
      t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath). "+
        " fileMgr.dMgr.absolutePath='%v'   Error='%v' ", fileMgr.dMgr.absolutePath, err.Error())
    }

  }

  if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {
    t.Errorf("Error: Failed to delete existing path '%v'",
      fileMgr.dMgr.absolutePath)
  }

  err = fileMgr.CreateDirAndFile()

  if err != nil {
    t.Errorf("Failed to Create Directory and File '%v', received Error:'%v'",
      fileMgr.absolutePathFileName, err.Error())
  }

  if !fh.DoesFileExist(fileMgr.absolutePathFileName) {
    t.Errorf(fmt.Sprintf("File Verfication failed file '%v' DOES NOT EXIST", fileMgr.absolutePathFileName))
  }

  s := "Created by File:'filemgr_test.go' Test Method: TestFileHelper_CreateDirAndFile()"

  _, err = fileMgr.WriteStrToFile(s)

  if err != nil {
    t.Errorf("Received error from fileMgr.WriteStrToFile(s). s='%v'  Error='%v' ", s, err.Error())
  }

  err = fileMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Received error from fileMgr.CloseThisFile(). fileMgr.absolutePathFileName='%v'  Error='%v' ", fileMgr.absolutePathFileName, err.Error())
  }

  err = fileMgr.dMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by fileMgr.dMgr.DeleteAll(). Attempted Deletion of %v. Error='%v'", fileMgr.absolutePathFileName, err.Error())
  }

}

func TestFileMgr_CreateDirAndFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  fileMgr.isInitialized = false

  err = fileMgr.CreateDirAndFile()

  if err == nil {
    t.Error("Expected an error return from fileMgr.CreateDirAndFile() because " +
      "fileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CreateThisFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  fileMgr.isInitialized = false

  err = fileMgr.CreateThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateThisFile() because " +
      "fileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CreateThisFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../iDoNotExist/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  dirMgr := fileMgr.GetDirMgr()

  _ = dirMgr.DeleteAll()

  err = fileMgr.CreateThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateThisFile() because " +
      "the fileMgr directory does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  _ = dirMgr.DeleteAll()

}

func TestFileMgr_DeleteThisFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile). "+
      "testFile='%v' Error='%v'", testFile, err.Error())
  }

  fileMgr.isInitialized = false

  err = fileMgr.DeleteThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.DeleteThisFile() because " +
      "the fileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_DeleteThisFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../checkfiles/scratchTestFile995681.txt")

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPathFileNameExtStr(testFile). "+
      "testFile='%v' Error='%v'", testFile, err.Error())
  }

  err = fileMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.CreateThisFile(). "+
      "File='%v' Error='%v'", fileMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.CreateThisFile(). "+
      "File='%v' Error='%v'", fileMgr.GetAbsolutePathFileName(), err.Error())
  }

  if fh.DoesFileExist(fileMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: Test file WAS NOT DELETED. File='%v' ",
      fileMgr.GetAbsolutePathFileName())

    err = fh.DeleteDirFile(fileMgr.GetAbsolutePathFileName())

    if err != nil {
      t.Errorf("Error: 2nd Attempted deletion of test file FAILED! "+
        "File='%v' ", fileMgr.GetAbsolutePathFileName())
    }

  }
}

func TestFileMgr_DoesFileExist_01(t *testing.T) {

  testFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  fileMgr.isInitialized = false

  if fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='false' because" +
      "the fileMgr is invalid. However, the return value was 'true'!")
  }

}

func TestFileMgr_DoesFileExist_02(t *testing.T) {

  testFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  if !fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='true' because" +
      "the 'FileMgr' file does exist. However, the return value was 'false'!")
  }

}

func TestFileMgr_DoesThisFileExist_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
  }

  fileMgr.isInitialized = false

  _, err = fileMgr.DoesThisFileExist()

  if err == nil {
    t.Error("Expected error return from fileMgr.DoesThisFileExist() because " +
      "the fileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_Empty_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by filePath1, err := fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.NewFromPathFileNameExtStr(filePath1). "+
      "filePath1='%v' Error='%v'", filePath1, err.Error())
  }

  fileMgr1.Empty()

  fileMgr2 := FileMgr{}
  fileMgr2.Empty()

  if !fileMgr1.Equal(&fileMgr2) {
    t.Error("Error: Expected empty fileMgr1 to equal empty fileMgr2. " +
      "However, THEY ARE NOT EQUAL!")
  }

}

func TestFileMgr_Equal_01(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by filePath1, err := fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.NewFromPathFileNameExtStr(filePath1). "+
      "filePath1='%v' Error='%v'", filePath1, err.Error())
  }

  fileMgr2 := fileMgr1.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != true {
    t.Error("Expected Equal to return 'true' for fileMgr1==fileMgr1, instead got: ", "false")
  }

}

func TestFileMgr_Equal_02(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  relPath2 := "..\\logTest\\FileMgmnt\\TestFile003.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2, instead got: 'true'")
  }

}

func TestFileMgr_Equal_03(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\"). "+
      "Error='%v'", err.Error())
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite()). "+
      "Error='%v'", err.Error())
  }

  fAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New(fOpenCfg, fPermCfg). "+
      "Error='%v'", err.Error())
  }

  fileMgr2.fileAccessStatus = fAccessCfg.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2. " +
      "Instead got: 'true'. fileAccessStatus are different.")
  }

}

func TestFileMgr_Equal_04(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  fileMgr2.dMgr = DirMgr{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2. " +
      "Instead got: 'true'. Directory Managers are different")
  }

}

func TestFileMgr_Equal_05(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  fileMgr2.actualFileInfo = FileInfoPlus{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2. " +
      "Instead got: 'true'. acutalFileInfo's are different")
  }

}

func TestFileMgr_EqualAbsPaths_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_1_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path to equal fileMgr2 absolute path. "+
      "Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualAbsPaths_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\level_1_1_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path to equal fileMgr2 absolute path. "+
      "Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualAbsPaths_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_2_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path to NOT fileMgr2 absolute path. "+
      "Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
      "They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\LEVEL_1_0_TEST.TXT"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
      "They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_02_dir\\LEVEL_1_0_TEST.TXT"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
      "They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_2_2_xray.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT equal fileMgr2 file name ext. "+
      "However, they ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.jag"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT equal fileMgr2 file name ext. "+
      "However, they ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualPathFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file name ext to equal "+
      "fileMgr2 absolute path file name ext. "+
      "Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\LEVEL_1_0_TEST.TXT"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file name ext to equal fileMgr2 "+
      "absolute path file name ext. Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_02_dir\\level_1_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
      "absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_X_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
      "absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
      "relPath1='%v' Error='%v'", relPath1, err.Error())
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
      "filePath1='%v' Error='%v' ", filePath1, err.Error())
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
      "relPath2='%v' Error='%v'", relPath2, err.Error())
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
      "filePath2='%v' Error='%v'", filePath2, err.Error())
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
      "absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}
