package pathfileops

import (
  "testing"
)

func TestFileMgr_MoveFileToFileMgr_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        destFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", destFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
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
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(destFile).\n"+
      "destFile='%v'\nError='%v'",
      destFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = srcFileMgr.MoveFileToFileMgr(destFileMgr)

  if err != nil {
    t.Errorf("Error returned from "+
      "srcFileMgr.MoveFileToFileMgr(destFileMgr).\n"+
      "destFileMgr='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(destFileMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: After 'move' operation, destination file "+
      "DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      destFileMgr.GetAbsolutePath())
  }

  if fh.DoesFileExist(srcFileMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: After 'move' operation, source file "+
      "still exists!\n"+
      "Source File='%v'\n",
      srcFileMgr.GetAbsolutePath())
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  return
}

func TestFileMgr_MoveFileToFileMgr_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        destFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", destFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
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
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.New(destFile).\n"+
      "destFile='%v'\nError='%v'",
      destFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr.isInitialized = false

  err = srcFileMgr.MoveFileToFileMgr(destFileMgr)

  if err == nil {
    t.Errorf("Expected an error to be returned by "+
      "srcFileMgr.MoveFileToFileMgr(destFileMgr) \n"+
      "because 'srcFileMgr' is INVALID!\n"+
      "However, NO ERROR WAS RETURNED!\n"+
      "destFileMgr='%v'\n",
      srcFile)
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  return
}

func TestFileMgr_MoveFileToFileStr_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        destFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", destFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
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
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  newFMgr, err := srcFileMgr.MoveFileToFileStr(destFile)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToFileStr(destFile).\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: After 'move' operation, destination file "+
      "DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      destFile)
  }

  if fh.DoesFileExist(srcFile) {
    t.Errorf("Error: After 'move' operation, source file "+
      "still exists!\n"+
      "Source File='%v'\n",
      srcFile)
  }

  if !newFMgr.DoesFileExist() {
    t.Errorf("Error: After 'move' operation 'newFMgr' does "+
      "DOES NOT EXIST!\n"+
      "newFMgr='%v'\n", newFMgr.GetAbsolutePathFileName())
  }
  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  return
}

func TestFileMgr_MoveFileToFileStr_02(t *testing.T) {

  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        destFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", destFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
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
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n", srcFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}." +
      "NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToFileStr(destFile)

  if err == nil {
    t.Errorf("Expected an error return from srcFileMgr." +
      "MoveFileToFileStr(destFile) \n"+
      "because srcFileMgr is INVALID!\n"+
      "However, NO ERROR WAS RETURNED!\n"+
      "srcFile='%v'\ndestFile='%v'\n",
      srcFile, destFile)
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  return
}

func TestFileMgr_MoveFileToNewDirMgr_01(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

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
      "srcFile='%v'\nError='%v'",
      srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr).\n"+
      "dMgr.path='%v'\nError='%v'\n",
      dMgr.path, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(newFMgr.GetAbsolutePathFileName()) {
    t.Errorf("Error: newFMgr Destination 'Moved' File DOES NOT EXIST!\n"+
      "srcFileMgr='%v'\nnewFMgr='%v'",
      srcFileMgr.GetAbsolutePathFileName(),
      newFMgr.GetAbsolutePathFileName())
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
    t.Errorf("Error: newFMgr Destination 'Moved' File DOES NOT EXIST!\n"+
      "newFMgr.DoesThisFileExist()=='FALSE'.\n"+
      "newFMgr='%v'", newFMgr.absolutePathFileName)
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

func TestFileMgr_MoveFileToNewDirMgr_02(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../../createFilesTest")
  setupDestFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(setupDestFile) {
    err := fh.DeleteDirFile(setupDestFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file\n"+
        "setupDestFile='%v'\nError:'%v'",
        setupDestFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(setupDestFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "setupDestFile='%v'", setupDestFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(srcFile)\n"+
        "srcFile='%v'\nError='%v'\n",
        srcFile, err.Error())

      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error - Failed to Delete 'srcFile'. 'srcFile'  STILL EXISTS!!\n"+
        "srcFile='%v'\n", srcFile)

      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(setupDestFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByIo(setupSrcFile, srcFile).\n"+
      "setupSrcFile='%v'\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n", setupSrcFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("ERROR: Source File does NOT EXIST!!\n"+
      "srcFile='%v'\n",
      srcFile)
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

  dMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(setupDestFile)
    return
  }

  srcFileMgr.isInitialized = false

  _, err = srcFileMgr.MoveFileToNewDirMgr(dMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(dMgr)\n" +
      "because 'srcFileMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!")
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

func TestFileMgr_MoveFileToNewDirMgr_03(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  setupDestFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

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

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      setupSrcFile, err.Error())

    return
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(DirMgr{})

  if err == nil {
    t.Error("Expected error return from srcFileMgr." +
      "MoveFileToNewDirMgr(DirMgr{})\n" +
      "because the input parameter is invalid.\n" +
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

func TestFileMgr_MoveFileToNewDirMgr_04(t *testing.T) {
  fh := FileHelper{}
  srcDir := fh.AdjustPathSlash("../../iDoNotExist")
  srcFile := fh.AdjustPathSlash("../../iDoNotExist/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../../createFilesTest")
  destFile := fh.AdjustPathSlash("../../createFilesTest/TestFile003.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error on DeleteDirFile() deleting destination file.\n"+
        "Destination File='%v'\n"+
        "Error='%v'\n",
        destFile, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error - destination file, STILL EXISTS!\n"+
        "Destination File='%v'\n", destFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  if fh.DoesFileExist(srcDir) {

    err := fh.DeleteDirPathAll(srcDir)

    if err != nil {
      t.Errorf("Error on DeleteDirPathAll() deleting source directory.\n"+
        "Source Directory='%v'\nError:'%v'\n",
        srcDir, err.Error())
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }

    if fh.DoesFileExist(srcDir) {
      t.Errorf("Error - Failed to Delete 'srcFile' STILL EXISTS!\n"+
        "srcFile='%v'\n", srcFile)
      _ = fh.DeleteDirFile(srcFile)
      _ = fh.DeleteDirFile(destFile)
      return
    }
  }

  srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile).\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    return
  }

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    return
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err == nil {
    t.Error("Expected error return from srcFileMgr.MoveFileToNewDirMgr(destDMgr)\n" +
      "because source file does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(srcDir)\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

}

func TestFileMgr_MoveFileToNewDirMgr_05(t *testing.T) {

  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destDir := fh.AdjustPathSlash("../../xxxIDoNotExist")
  setupDestFile := fh.AdjustPathSlash("../../xxxIDoNotExist/TestFile003.txt")

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
      setupSrcFile, err.Error())
    return
  }

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir).\n"+
      "destDir='%v'\nError='%v'\n",
      destDir, err.Error())
    return
  }

  _, err = srcFileMgr.MoveFileToNewDirMgr(destDMgr)

  if err != nil {
    t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(destDMgr)\n"+
      "destDMgr='%v'\nError='%v'\n", destDMgr.GetAbsolutePath(), err.Error())
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Attempted deletion of setup directory Failed!\n"+
      "Error returned by fh.DeleteDirPathAll(destDir)\n"+
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
  }

}
