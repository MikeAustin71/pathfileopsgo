package pathfileops

import (
  "os"
  "testing"
)

func TestFileMgr_MoveFileToNewDir_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile).\n"+
        "setupDestFile='%v'\nError:'%v'",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error: Attempted Deletion Failed!!\n"+
        "Destination file STILL EXISTS!\n"+
        "setupDestFile='%v'", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error: Attempted Deletion Failed!!\n"+
        "'srcFile' SILL EXISTS!!\n"+
        "srcFile='%v'\n"+
        "Error:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error: Attempted Deletion Failed!\n"+
        "'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Attempt copy operation failed!!\n"+
      "Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
    t.Errorf("Error: Move Operation Failed!\nDestination File DOES NOT EXIST!\n"+
      "Destination File (newFMgr)='%v'", newFMgr.absolutePathFileName)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by newFMgr.DoesThisFileExist().\n"+
      "newFMgr='%v'\nError='%v'\n",
      newFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !doesExist {
    t.Errorf("Error: After Move Operation Destination File DOES NOT EXIST!\n"+
      "Destination File (newFMgr)='%v'",
      newFMgr.absolutePathFileName)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  doesExist, err = fh.DoesThisFileExist(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DoesThisFileExist(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n", srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if doesExist {
    t.Errorf("Error: After move operation, the source file still exists!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(setupDestFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
      "setupDestFile='%v'\n"+
      "Error='%v'\n",
      setupDestFile, err.Error())
  }

}

func TestFileMgr_MoveFileToNewDir_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
        "setupDestFile='%v'\nError='%v'\n",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error: Destination File STILL EXISTS!\n"+
        "setupDestFile='%v'\n", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error deleting source file\n"+
        "Error returned by fh.DeleteDirFile(srcFile)\n"+
        "'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\nError:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error: Attempted Deletion Failed!\n"+
        "'srcFile'  STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\n"+
      "srcFile='%v'\n",
      setupSrcFile,
      srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Copy operation FAILED!\n"+
      "Destination file DOES NOT EXIST!!\n"+
      "Destination file='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDir(destDir)\n" +
      "because srcFileMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(setupDestFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
      "setupDestFile='%v'\n"+
      "Error='%v'\n",
      setupDestFile, err.Error())
  }
}

func TestFileMgr_MoveFileToNewDir_03(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  setupDestFile := fh.AdjustPathSlash("../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    return
  }

  _, err = srcFileMgr.MoveFileToNewDir("")

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDir(\"\") " +
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(setupDestFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
      "setupDestFile='%v'\n"+
      "Error='%v'\n",
      setupDestFile, err.Error())
  }

}

func TestFileMgr_MoveFileToNewDir_04(t *testing.T) {

  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  targetFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir)\n"+
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !targetFMgr.DoesFileExist() {
    t.Errorf("Error: After Move Operation Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n", targetFMgr.GetAbsolutePath())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(setupDestFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
      "setupDestFile='%v'\n"+
      "Error='%v'\n",
      setupDestFile, err.Error())
  }

}

func TestFileMgr_MoveFileToNewDir_05(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../iDoNotExit")
  setupDestFile := fh.AdjustPathSlash("../iDoNotExit/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    return
  }

  targetFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir)\n"+
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if !targetFMgr.DoesFileExist() {
    t.Errorf("Error: After 'move' operation, Target File DOES NOT EXIST!\n"+
      "destDir='%v'\n", destDir)
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(destDir)\n"+
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
  }
}

func TestFileMgr_MoveFileToNewDir_06(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("     ")

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    return
  }

  _, err = srcFileMgr.MoveFileToNewDir(destDir)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDir(destDir)\n" +
      "because the 'destDir' string consists of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

}

func TestFileMgr_MoveFileToNewDir_07(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFile := fh.AdjustPathSlash("../checkfiles/level_0_3_test.txt")
  destDir := fh.AdjustPathSlash("../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../createFilesTest/level_0_3_test.txt")

  if fh.DoesFileExist(setupDestFile) {

    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile).\n"+
        "setupDestFile='%v'\nError:'%v'",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error: Attempted Deletion Failed!!\n"+
        "Destination file STILL EXISTS!\n"+
        "setupDestFile='%v'", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error: Attempted Deletion Failed!!\n"+
        "'srcFile' SILL EXISTS!!\n"+
        "srcFile='%v'\n"+
        "Error:'%v'\n",
        srcFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error: Attempted Deletion Failed!\n"+
        "'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Attempt copy operation failed!!\n"+
      "Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
    t.Errorf("Error: Move Operation Failed!\nDestination File DOES NOT EXIST!\n"+
      "Destination File (newFMgr)='%v'", newFMgr.absolutePathFileName)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  doesExist, err := newFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by newFMgr.DoesThisFileExist().\n"+
      "newFMgr='%v'\nError='%v'\n",
      newFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !doesExist {
    t.Errorf("Error: After Move Operation Destination File DOES NOT EXIST!\n"+
      "Destination File (newFMgr)='%v'",
      newFMgr.absolutePathFileName)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  doesExist, err = fh.DoesThisFileExist(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DoesThisFileExist(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n", srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if doesExist {
    t.Errorf("Error: After move operation, the source file still exists!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(setupDestFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(setupDestFile)\n"+
      "setupDestFile='%v'\n"+
      "Error='%v'\n",
      setupDestFile, err.Error())
  }

}

func TestFileMgr_New_01(t *testing.T) {

  fh := FileHelper{}

  relPath := "..\\logTest\\CmdrX\\CmdrX.log"
  commonDir, err := fh.MakeAbsolutePath(relPath)

  if err != nil {
    t.Errorf("Received Error on fh.MakeAbsolutePath(relPath).\n"+
      "relPath='%v'\nError='%v'\n",
      relPath, err.Error())
    return
  }

  fileName := "CmdrX"
  fileNameExt := "CmdrX.log"
  extName := ".log"

  fileMgr, err := FileMgr{}.New(commonDir)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(commonDir)\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  if fileMgr.fileName != fileName {
    t.Errorf("Expected File Name='%v'\n"+
      "Instead, File Name='%v'\n",
      fileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Errorf("Expected File Extension='%v'\n"+
      "Instead, File Extension='%v'",
      extName, fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Errorf("Expected File Name + Extension='%v'\n"+
      "Instead, File Name + Extension='%v'\n",
      fileNameExt, fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'.\n" +
      "Instead, fileMgr.isInitialized=='false'.\n")
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'.\n" +
      "Instead, fileMgr.isFileNamePopulated='false'\n")
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'\n" +
      "Instead, fileMgr.isFileNameExtPopulated=='false'\n")
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true'.\n" +
      "Instead, fileMgr.isFileExtPopulated='false'.\n")
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, fileMgr.isAbsolutePathFileNamePopulated='false'\n")
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
    t.Errorf("Error returned by FileMgr{}.New(commonDir)\n"+
      "commonDir='%v'\nError='%v'\n",
      commonDir, err.Error())
    return
  }

  if fileMgr.fileName != fileName {
    t.Errorf("Expected File Name='%v'.\n"+
      "Instead, File Name='%v'\n",
      fileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != extName {
    t.Errorf("Expected File Extension='%v'\n"+
      "Instead, File Extension='%v'\n",
      extName, fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Errorf("Expected File Name + Extension='%v'.\n"+
      "Instead, File Name + Extension='%v'.\n",
      fileNameExt, fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'.\n" +
      "Instead, fileMgr.isInitialized=='false'\n")
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'\n" +
      "Instead, fileMgr.isFileNamePopulated='false'\n")
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'\n" +
      "Instead, fileMgr.isFileNameExtPopulated='false'\n")
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true'\n" +
      "Instead, fileMgr.isFileExtPopulated='false'\n")
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'\n" +
      "Instead, fileMgr.isAbsolutePathFileNamePopulated='false'\n")
  }

}

func TestFileMgr_New_03(t *testing.T) {

  _, err := FileMgr{}.New("")

  if err == nil {
    t.Error("Expected error return from FileMgr{}.New(\"\")\n" +
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_New_04(t *testing.T) {

  _, err := FileMgr{}.New("!^%&*()")

  if err == nil {
    t.Error("Expected error return from FileMgr{}.New(\"!^%&*()\")\n" +
      "because the input parameter contains invalid characters.\n" +
      "However, NO ERROR WAS RETURNED!\n")
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
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(commonDir)\n"+
      "commonDir='%v'\nError='%v'\n",
      commonDir, err)
    return
  }

  if fileMgr.fileName != fileName {
    t.Errorf("Expected File Name='%v'.\n"+
      "Instead, File Name='%v'\n",
      fileName, fileMgr.fileName)

  }

  if fileMgr.fileExt != extName {
    t.Errorf("Expected File Extension='%v'\n"+
      "Instead, File Extension='%v'\n", extName, fileMgr.fileExt)
  }

  if fileMgr.fileNameExt != fileNameExt {
    t.Errorf("Expected File Name + Extension='%v'.\n"+
      "Instead, File Name + Extension='%v'\n",
      fileNameExt, fileMgr.fileNameExt)
  }

  if !fileMgr.isInitialized {
    t.Error("Expected fileMgr.isInitialized=='true'.\n" +
      "Instead, fileMgr.isInitialized='false'\n")
  }

  if !fileMgr.isFileNamePopulated {
    t.Error("Expected fileMgr.isFileNamePopulated=='true'.\n" +
      "Instead, fileMgr.isFileNamePopulated='false'\n")
  }

  if !fileMgr.isFileNameExtPopulated {
    t.Error("Expected fileMgr.isFileNameExtPopulated=='true'\n" +
      "Instead, fileMgr.isFileNameExtPopulated='false'\n")
  }

  if !fileMgr.isFileExtPopulated {
    t.Error("Expected fileMgr.isFileExtPopulated=='true'\n" +
      "Instead, fileMgr.isFileExtPopulated='false'\n")
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'\n" +
      "Instead, fileMgr.isAbsolutePathFileNamePopulated=='false'\n")
  }

}

func TestFileMgr_NewFromPathFileNameExtStr_02(t *testing.T) {

  path := "../appExamples/filehelperexamples.go"

  eFileNameExt := "filehelperexamples.go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromPathFileNameExtStr(path)\n"+
      "path=='%v'\nError='%v'",
      path, err)
  }

  if eFileNameExt != fileMgr.fileNameExt {
    t.Errorf("Expected extracted fileNameExt=='%v'.\n"+
      "Instead fileNameExt=='%v'\n",
      eFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != "filehelperexamples" {
    t.Errorf("Expected fileMgr.fileName=='filehelperexamples'.\n"+
      "Instead, fileMgr.fileName== %v\n",
      fileMgr.fileName)
  }

  if fileMgr.fileExt != ".go" {
    t.Errorf("Expected fileMgr.fileExt=='.go'\n"+
      "Instead, fileMgr.fileExt== %v",
      fileMgr.fileExt)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Error("Expected 'fileMgr.isPathPopulated==true'.\n" +
      "Instead, fileMgr.isPathPopulated=='false'\n")
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected 'fileMgr.doesAbsolutePathFileNameExist==true'.\n" +
      "Instead fileMgr.doesAbsolutePathFileNameExist=='false'\n")
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, it is 'false'.")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("Expected fileMgr.doesAbsolutePathExist=='true'.\n" +
      "Instead, it is 'false'.\n")
  }

}

func TestFileMgr_NewFromPathFileNameExtStr_03(t *testing.T) {

  path := "filehelperexamples"

  _, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err == nil {
    t.Error("Expected an error from FileMgr{}." +
      "NewFromPathFileNameExtStr(path)\n" +
      "because path='filehelperexamples'.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_NewFromPathFileNameExtStr_04(t *testing.T) {

  path := "../appExamples/filehelperexamples.go"

  eFileNameExt := "filehelperexamples.go"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path)\n"+
      "path=='%v'\nError: %v\n",
      path, err)
    return
  }

  if eFileNameExt != fileMgr.fileNameExt {
    t.Errorf("Expected extracted fileNameExt=='%v'\n"+
      "Instead fileNameExt=='%v'\n",
      eFileNameExt, fileMgr.fileNameExt)
  }

  if "filehelperexamples" != fileMgr.fileName {
    t.Errorf("Expected fileMgr.fileName== '%v'\n"+
      "Instead fileMgr.fileName== '%v'\n",
      "filehelperexamples", fileMgr.fileName)
  }

  if ".go" != fileMgr.fileExt {
    t.Errorf("Expected fileMgr.fileExt=='.go'\n"+
      "Instead, fileMgr.fileExt== %v\n",
      fileMgr.fileExt)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.dMgr.isPathPopulated==true'\n"+
      "Instead, fileMgr.isPathPopulated==%v\n",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true'\n"+
      "Instead fileMgr.doesAbsolutePathFileNameExist=='%v'\n",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Errorf("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n"+
      "Instead, fileMgr.isAbsolutePathFileNamePopulated=='%v'\n",
      fileMgr.isAbsolutePathFileNamePopulated)
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Errorf("Expected fileMgr.doesAbsolutePathExist=='true'\n"+
      "Instead, it is '%v'\n",
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
    t.Errorf("Error returned from fh.MakeAbsolutePath("+
      "adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  absPathFileNameExt :=
    absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n",
      absPathFileNameExt, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromFileInfo(absPath, info)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info).\n"+
      "absPath='%v'\ninfo.Name()='%v'\nError='%v'\n",
      absPath, info.Name(), err.Error())
    return
  }

  if fileMgr.fileNameExt != expectedFileNameExt {
    t.Errorf("Expected extracted fileMgr.fileNameExt=='%v'\n"+
      "Instead fileMgr.fileNameExt='%v'\n",
      expectedFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != expectedFileName {
    t.Errorf("Expected fileMgr.fileName=='%v'\n"+
      "Instead fileMgr.fileName=='%v'\n",
      expectedFileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != expectedExt {
    t.Errorf("Expected fileMgr.fileExt=='%v'\n"+
      "Instead fileMgr.fileExt=='%v'\n",
      expectedExt, fileMgr.fileName)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.isPathPopulated=='true'\n"+
      "Instead fileMgr.isPathPopulated=='%v'",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist"+
      "==true'\n"+
      "Instead fileMgr.doesAbsolutePathFileNameExist=='%v'",
      fileMgr.doesAbsolutePathFileNameExist)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("ERROR: Expected fileMgr." +
      "isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, it is 'false'\n")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("ERROR: Expected fileMgr.doesAbsolutePathExist==true'.\n" +
      "Instead, it is 'false'\n")
  }

  if !fileMgr.actualFileInfo.IsFInfoInitialized {
    t.Error("ERROR: Expected fileMgr.actualFileInfo.IsFInfoInitialized" +
      "='true'.\nInstead, it is 'false'")
  }

  if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("ERROR: Expected fileMgr.actualFileInfo.Name()=='%v'.\n"+
      "Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
  }

}

func TestFileMgr_NewFromFileInfo_02(t *testing.T) {

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  var info os.FileInfo

  _, err = FileMgr{}.NewFromFileInfo(absPath, info)

  if err == nil {
    t.Errorf("Expected an error from FileMgr{}.NewFromFileInfo(absPath, info)\n" +
      "because input parameter 'info' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_NewFromFileInfo_03(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  absPathFileNameExt :=
    absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n",
      absPathFileNameExt, err.Error())
  }

  absPath = "../iDoNotExist"

  absPath, err = fh.MakeAbsolutePath(absPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(absPath).\n"+
      "absPath='%v'\nError='%v'\n",
      absPath, err.Error())
    return
  }

  expectedFilePath :=
    absPath + string(os.PathSeparator) + expectedFileNameExt

  nFMgr, err := FileMgr{}.NewFromFileInfo(absPath, info)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info).\n"+
      "absPath='%v'\ninfo.Name()='%v'\nError='%v'\n",
      absPath, info.Name(), err.Error())
    return
  }

  if expectedFilePath != nFMgr.GetAbsolutePathFileName() {
    t.Errorf("ERROR: Expected File Path='%v'.\n"+
      "Instead, File Path='%v'\n",
      expectedFilePath, nFMgr.GetAbsolutePathFileName())
  }
}

func TestFileMgr_NewFromFileInfo_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  absPathFileNameExt :=
    absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfo(absPathFileNameExt)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileInfo(absPathFileNameExt).\n"+
      "absPathFileNameExt='%v'\nError='%v'\n",
      absPathFileNameExt, err.Error())
    return
  }

  absPath = ""

  _, err = FileMgr{}.NewFromFileInfo(absPath, info)

  if err == nil {
    t.Error("Expected an error from FileMgr{}.NewFromFileInfo(absPath, info)\n" +
      "because absPath is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_01(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\nadjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  fMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dMgr, expectedFileNameExt).\n"+
      "dMgr='%v'\nrawFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), expectedFileNameExt, err.Error())
    return
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
    t.Errorf("Expected absolutePathFileName='%v'.\n"+
      "Instead, absolutePathFileName='%v'\n",
      expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
  }
}

func TestFileMgr_NewFromDirMgrFileNameExt_02(t *testing.T) {

  rawFileNameExt := "./newerFileForTest_01.txt"
  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  fMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, rawFileNameExt).\n"+
      "dMgr='%v'\nrawFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), rawFileNameExt, err.Error())
  }

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
  }

  expectedAbsPathFileNameExt :=
    absPath + string(os.PathSeparator) + expectedFileNameExt

  if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
    t.Errorf("Expected absolutePathFileName='%v'.\n"+
      "Instead, absolutePathFileName='%v'\n",
      expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  rawPath := "../filesfortest/newfilesfortest"

  absolutePath, err := fh.MakeAbsolutePath(rawPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawPath).\n"+
      "rawPath='%v'\nError='%v'\n",
      rawPath, err.Error())
  }

  dMgr, err := DirMgr{}.New(absolutePath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(absolutePath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      absolutePath, err.Error())
    return
  }

  _, err = FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "")

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}." +
      "NewFromDirMgrFileNameExt(dMgr, \"\")\n" +
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_04(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr"+
      "(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  dMgr.isInitialized = false

  _, err =
    FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}" +
      "NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)\n" +
      "because the dMgr is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_NewFromDirMgrFileNameExt_05(t *testing.T) {

  expectedFileNameExt := "$%!*().#+_"

  fh := FileHelper{}

  adjustedPath := fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}."+
      "NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  _, err =
    FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err == nil {
    t.Errorf("Expected error return from FileMgr{}" +
      "NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)\n" +
      "because the expectedFileNameExt contains invalid characters.\n" +
      "However, NO ERROR WAS RETURNED!\n")
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
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  fileMgr, err :=
    FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirStrFileNameStr("+
      "rawPath, expectedFileNameExt).\n"+
      "rawPath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      rawPath, expectedFileNameExt, err.Error())
    return
  }

  if fileMgr.fileNameExt != expectedFileNameExt {
    t.Errorf("Expected extracted fileMgr.fileNameExt==%v.\n"+
      "Instead fileMgr.fileNameExt='%v'\n",
      expectedFileNameExt, fileMgr.fileNameExt)
  }

  if fileMgr.fileName != expectedFileName {
    t.Errorf("Expected fileMgr.fileName== '%v'\n"+
      "Instead fileMgr.fileName== '%v'\n",
      expectedFileName, fileMgr.fileName)
  }

  if fileMgr.fileExt != expectedExt {
    t.Errorf("Expected fileMgr.fileExt== '%v'\n"+
      "Instead got: fileMgr.fileExt=='%v'\n",
      expectedExt, fileMgr.fileName)
  }

  if !fileMgr.dMgr.isPathPopulated {
    t.Errorf("Expected 'fileMgr.isPathPopulated==true'.\n"+
      "Instead fileMgr.isPathPopulated=='%v'\n",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.doesAbsolutePathFileNameExist {
    t.Errorf("Expected fileMgr.doesAbsolutePathFileNameExist"+
      "=='true'\nInstead fileMgr.doesAbsolutePathFileNameExist=='%v'\n",
      fileMgr.dMgr.isPathPopulated)
  }

  if !fileMgr.isAbsolutePathFileNamePopulated {
    t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true'.\n" +
      "Instead, it is 'false'.\n")
  }

  if !fileMgr.dMgr.doesAbsolutePathExist {
    t.Error("Expected fileMgr.doesAbsolutePathExist=='true'.\n" +
      "Instead, it is 'false'.\n")
  }

  if !fileMgr.actualFileInfo.IsFInfoInitialized {
    t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.\n" +
      "Error, it is 'false'.\n")
  }

  if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
    t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.\n"+
      "Instead fileMgr.actualFileInfo.Name()=='%v'.\n",
      expectedFileNameExt, fileMgr.actualFileInfo.Name())
  }

  if expectedAbsPath != fileMgr.dMgr.absolutePath {
    t.Errorf("Expected absolutePath='%v'.\n"+
      "Instead, absolutePath='%v'\n",
      expectedAbsPath, fileMgr.dMgr.absolutePath)
  }

  if expectedPath != fileMgr.dMgr.path {
    t.Errorf("Expected path='%v'.\n"+
      "Instead, path='%v'\n",
      expectedPath, fileMgr.dMgr.path)
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_02(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := ""

  _, err :=
    FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)\n" +
      "because rawPath is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_NewFromDirStrFileNameStr_03(t *testing.T) {

  expectedFileNameExt := ""

  rawPath := "../filesfortest/newfilesfortest"

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)\n" +
      "because expectedFileNameExt is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_NewFromDirStrFileNameStr_04(t *testing.T) {

  expectedFileNameExt := "     "

  rawPath := "../filesfortest/newfilesfortest"

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)\n" +
      "because expectedFileNameExt consists of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_NewFromDirStrFileNameStr_05(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := ""

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)\n" +
      "because raw path is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_NewFromDirStrFileNameStr_06(t *testing.T) {

  expectedFileNameExt := "newerFileForTest_01.txt"

  rawPath := "    "

  _, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    t.Error("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)\n" +
      "because raw path consists of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}
