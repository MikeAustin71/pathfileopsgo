package pathfileops

import "testing"

func TestFileMgr_CopyFileToDirByIo_01(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByIo(destDMgr).\n"+
      "destPath='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()

    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'",
      newFileMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()
    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      srcFMgr.GetAbsolutePathFileName(),
      newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByIo_02(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByIo_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because\n" +
      "srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByIo_04(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  destDMgr := srcFMgr.GetDirMgr()

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because\n" +
      "destination directory is equivalent to source directory.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileToDirByIo_05(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because\n" +
      "srcFMgr is invalid.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByIo_06(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/iDoNotExist.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr)\n" +
      "because source file does NOT exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
      "Source File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByIoByLink(destDMgr).\n"+
      "destPath='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()

    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'",
      newFileMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()

    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      srcFMgr.GetAbsolutePathFileName(),
      newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_02(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned by srcFMgr.DoesThisFileExist().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_03(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, fileName)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(destDMgr, fileName).\n"+
      "destDMgr='%v'\nfileName='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), fileName, err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByIoByLink_04(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  destDMgr := srcFMgr.GetDirMgr()

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because\n" +
      "srcFMgr directory is equivalent to destination directory.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileToDirByIoByLink_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist.txt"

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
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

  if doesFileExist {
    err = srcFMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Error returned from srcFMgr.DeleteThisFile().\n"+
        "srcFMgr='%v'\nError='%v'\n",
        srcFMgr.GetAbsolutePath(), err.Error())
      return
    }
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because\n" +
      "the source file does NOT exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = srcFMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by srcFMgr.DeleteThisFile().\n"+
      "srcFMgr= '%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), err.Error())
  }

}