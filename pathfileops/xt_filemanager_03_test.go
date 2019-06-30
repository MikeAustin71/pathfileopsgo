package pathfileops

import "testing"

func TestFileMgr_CopyFileStrByIo_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {

    err = fh.DeleteDirFile(rawAbsDestPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath).\n"+
        "rawAbsDestPath='%v'\nError='%v'\n",
        rawAbsDestPath, err.Error())
      return
    }

  }

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByIo(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v' ", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByIo_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v' ",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source Test File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByIo_02.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Error returned by fh.DoesThisFileExist(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n", rawAbsDestPath, err.Error())
    return
  }

  if doesFileExist {

    err = fh.DeleteDirFile(rawAbsDestPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath).\n"+
        "rawAbsDestPath='%v'\nError='%v'\n",
        rawAbsDestPath, err.Error())
      return
    }
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
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by srcFMgr.DoesThisFileExist().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByIo("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(\"\") " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIo_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile). "+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByIo_04.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  doesFileExist := fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {

    err = fh.DeleteDirFile(rawAbsDestPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
        "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIo_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  rawAbsDestPath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
      "because source file is equivalent to destination file. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIo_06(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {

    err = fh.DeleteDirFile(rawAbsDestPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath).\n"+
        "rawAbsDestPath='%v'\nError='%v'\n",
        rawAbsDestPath, err.Error())
      return
    }

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) \n" +
      "because srcFMgr is INVALID!\n " +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v' ", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByIo_07(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/TestFileMgr_CopyFileStrByIo_07.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.DoesThisFileExist(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  if doesFileExist {
    t.Errorf("Test Setup Error: After deletion attempt, 'rawAbsDestPath' "+
      "still exists!\nrawAbsDestPath='%v'\n", rawAbsDestPath)
    return
  }

  setupFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"

  setupFile, err = fh.MakeAbsolutePath(setupFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupFile)\n"+
      "setupFile='%v'\nError='%v'\n", setupFile, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, rawAbsDestPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, rawAbsDestPath)\n"+
      "setupFile='%v'\nrawAbsDestPath='%v'\nError='%v'\n", setupFile, rawAbsDestPath, err.Error())
    _ = fh.DeleteDirFile(rawAbsDestPath)
    return
  }

  err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByIo(rawAbsDestPath)\n"+
      "rawAbsDestPath='%v'\nError='%v'\n", rawAbsDestPath, err.Error())
    _ = fh.DeleteDirFile(rawAbsDestPath)
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("After Copy Operation - Error returned by fh.DoesThisFileExist(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    _ = fh.DeleteDirFile(rawAbsDestPath)
    return
  }

  if !doesFileExist {
    t.Errorf("After Copy Operation DESTINATION FILE DOES NOT EXIST!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v' ", rawAbsDestPath)
  }
}
