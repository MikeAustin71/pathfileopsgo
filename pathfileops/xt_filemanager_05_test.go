package pathfileops

import (
  "fmt"
  "os"
  "testing"
)

func TestFileMgr_MoveFileToNewDirMgr_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). destDir='%v' Error='%v'", destDir, err.Error())
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr). dMgr.path='%v'  Error='%v'", dMgr.path, err.Error())
  }

  if !fh.DoesFileExist(newFMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
  }

  if !doesExist {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  err = newFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

}

func TestFileMgr_MoveFileToNewDirMgr_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). "+
      "destDir='%v' Error='%v'", destDir, err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "TestFile003.txt")

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, \"TestFile003.txt\"). "+
      "Error='%v'", err.Error())
  }

  srcFileMgr2 := srcFileMgr.CopyOut()

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because srcFileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = newFMgr.DeleteThisFile()

  if fh.DoesFileExist(newFMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: Deletion of New File failed. File='%v'",
      newFMgr.GetAbsolutePathFileName())
  }

  err = srcFileMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of source file FAILED!. "+
      "Source File ='%v'", srcFileMgr2.GetAbsolutePathFileName())
  }

}

func TestFileMgr_MoveFileToNewDirMgr_03(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(DirMgr{})

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because dMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDirMgr_04(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\iDoNotExist.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  destDir := fh.AdjustPathSlash("..\\logTest")

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "directory='%v'  Error='%v'", destDir, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDirMgr_05(t *testing.T) {

  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", setupSrcFile, err.Error())
  }

  destDir := fh.AdjustPathSlash("..\\xxxIDoNotExist")

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "directory='%v'  Error='%v'", destDir, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr) " +
      "because dMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
  }

  if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
  }

  if !doesExist {
    t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

  err = newFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
  }

}

func TestFileMgr_MoveFileToNewDir_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")
  setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  srcFileMgr2 := srcFileMgr.CopyOut()

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because srcFileMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

  err = srcFileMgr2.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned from srcFileMgr2.DeleteThisFile(). "+
      "srcFile2='%v'  Error='%v'",
      srcFileMgr2.GetAbsolutePathFileName(), err.Error())
  }

  if fh.DoesFileExist(srcFileMgr2.GetAbsolutePathFileName()) {
    t.Errorf("ERROR: Deletion of test file failed!. Test File='%v' ",
      srcFileMgr2.GetAbsolutePathFileName())
  }

}

func TestFileMgr_MoveFileToNewDir_03(t *testing.T) {
  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir("")

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because destDir is an empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_04(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\iDoNotExist.txt")
  destDir := fh.AdjustPathSlash("..\\logTest")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_MoveFileToNewDir_05(t *testing.T) {
  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
  destDir := fh.AdjustPathSlash("..\\iDoNotExit")

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). "+
      "srcFile='%v'  Error='%v'", srcFile, err.Error())
  }

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from  srcFileMgr.MoveFileToNewDir(destDir) " +
      "because destDir does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

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

  expectedFileNameExt := "     "

  rawPath := "../filesfortest/newfilesfortest"

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "expectedFileNameExt consists of blank spaces. " +
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

func TestFileMgr_NewFromDirStrFileNameStr_06(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := "    "

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "raw path consists of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}
