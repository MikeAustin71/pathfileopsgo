package pathfileops

import (
  "testing"
)

func TestFileMgr_CopyFileToDirByLinkByIo_01(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByLinkByIo(destDMgr)\n. "+
      "destPath='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLinkByIo_02(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr) because " +
      "srcFMgr.isInitialized==false. However, NO ERROR WAS RETURNED!")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLinkByIo_03(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr)\n" +
      "because destDMgr.isInitialized==false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLinkByIo_04(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  destDMgr := srcFMgr.GetDirMgr()

  err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr)\n" +
      "because source directory manager equals destination directory manager.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileToDirByLinkByIo_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "Error='%v' ", err.Error())
    return
  }

  rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

  destDMgr, err := DirMgr{}.New(rawDestPath)

  err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr)\n" +
      "because source file does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileToDirByLink_01(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByLink(destDMgr).\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())

    _ = newFileMgr.DeleteThisFile()
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Non-Path Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLink_02(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByLink(destDMgr)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLink_03(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr)\n" +
      "because destDMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr= '%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }
}

func TestFileMgr_CopyFileToDirByLink_04(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../filesfortest/newfilesfortest/" + fileName

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
    t.Errorf("Error: Source Test File DOES NOT EXIST!\nSource File='%v'\n",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  destDMgr := srcFMgr.GetDirMgr()

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr)\n" +
      "because source directory equals destination directory.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileToDirByLink_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "Error='%v'\n", err.Error())
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

  err = srcFMgr.CopyFileToDirByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr)\n" +
      "because source file does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CreateDir_01(t *testing.T) {

  fh := FileHelper{}

  pathFileName :=
    "../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateDir_01.txt"

  testFile := fh.AdjustPathSlash(pathFileName)

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPathFileNameExtStr(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  doesThisFileExist, err :=
    fh.DoesThisFileExist(fileMgr.dMgr.absolutePath)

  if err != nil {
    t.Errorf("Setup Non-Path Error returned by "+
      "fh.DoesThisFileExist(fileMgr.dMgr.absolutePath)\n"+
      "fileMgr.dMgr.absolutePath='%v'\nError='%v'\n",
      fileMgr.dMgr.absolutePath, err.Error())
    return
  }

  if doesThisFileExist {

    err = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    if err != nil {
      t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath).\n"+
        "Attempted Directory Deletion FAILED!\n"+
        "fileMgr.dMgr.absolutePath='%v'\nError='%v'\n",
        fileMgr.dMgr.absolutePath, err.Error())
      return
    }
  }

  err = fileMgr.CreateDir()

  if err != nil {
    t.Errorf("Error returned from fileMgr.CreateDir().\n"+
      "fileMgr='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    return
  }

  dirMgr := fileMgr.GetDirMgr()

  doesThisFileExist, err = fh.DoesThisFileExist(dirMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Non-Path Error returned from path!\n"+
      "fileMgr.CreateDir() FAILED!\n"+
      "Path='%v'\nError='%v'\n",
      dirMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesThisFileExist {
    t.Errorf("Error: Failed to create directory path!\n"+
      "Directory Path='%v'\n",
      dirMgr.GetAbsolutePath())
    return

  } else {

    err = dirMgr.DeleteAll()

    if err != nil {
      t.Errorf("Error returned from dirMgr.DeleteAll().\n"+
        "dirMgr='%v'\nError='%v'\n",
        dirMgr.GetAbsolutePath(), err.Error())
      return
    }

    doesThisFileExist, err = fh.DoesThisFileExist(dirMgr.GetAbsolutePath())

    if err != nil {
      t.Errorf("#2 Non-Path Error returned from path!\n"+
        "Final Deletion of Directory Path FAILED!\n"+
        "Path='%v'\nError='%v'\n",
        dirMgr.GetAbsolutePath(), err.Error())
      return
    }

    if doesThisFileExist {
      t.Errorf("ERROR: Final Deletion of Directory Path FAILED!\n"+
        "File Manager Directory Path='%v'\n",
        dirMgr.GetAbsolutePath())
      return
    }
  }
}

func TestFileMgr_CreateDir_02(t *testing.T) {

  fileMgr := FileMgr{}

  err := fileMgr.CreateDir()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateDir()\n" +
      "because 'fileMgr' (File Manager) was NOT initialized.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CreateDirAndFile_01(t *testing.T) {
  fh := FileHelper{}

  testFile := fh.AdjustPathSlash(
    "../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateDirAndFile_01.txt")

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPathFileNameExtStr(testFile)\n"+
      "testFile='%v', Error='%v'",
      testFile, err.Error())
    return
  }

  doesThisFileExist, err :=
    fh.DoesThisFileExist(fileMgr.dMgr.absolutePath)

  if err != nil {
    t.Errorf("Setup Non-Path Error returned by "+
      "fh.DoesThisFileExist(fileMgr.dMgr.absolutePath)\n"+
      "fileMgr.dMgr.absolutePath='%v'\nError='%v'\n",
      fileMgr.dMgr.absolutePath, err.Error())
    return
  }

  if doesThisFileExist {

    err = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    if err != nil {
      t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath).\n"+
        "Attempted Directory Deletion FAILED!\n"+
        "fileMgr.dMgr.absolutePath='%v'\nError='%v'\n",
        fileMgr.dMgr.absolutePath, err.Error())
      return
    }
  }

  err = fileMgr.CreateDirAndFile()

  if err != nil {
    t.Errorf("Failed to Create Directory and File\n"+
      "Error returned by fileMgr.CreateDirAndFile()\n"+
      "fileMgr='%v'\nError='%v'",
      fileMgr.absolutePathFileName, err.Error())

    _ = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    return
  }

  err = fileMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.CloseThisFile().\n"+
      "Error='%v'\n", err.Error())
  }

  doesThisFileExist, err = fh.DoesThisFileExist(fileMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Non-Path Error retrned by fh.DoesThisFileExist("+
      "fileMgr.absolutePathFileName)\n"+
      "fileMgr.absolutePathFileName='%v'\nError='%v'\n",
      fileMgr.absolutePathFileName, err.Error())

    _ = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    return
  }

  if !doesThisFileExist {
    t.Errorf("File Verfication failed!\n"+
      "File DOES NOT EXIST!"+
      "Path File Name='%v'", fileMgr.absolutePathFileName)

    _ = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

    return
  }

  s := "Created by File:'xt_filemanger_03_test.go' " +
    "Test Method: TestFileHelper_CreateDirAndFile()"

  _, err = fileMgr.WriteStrToFile(s)

  if err != nil {
    t.Errorf("Received error from fileMgr.WriteStrToFile(s).\n"+
      "s='%v'\n\nError='%v'\n", s, err.Error())
  }

  err = fileMgr.CloseThisFile()

  if err != nil {
    t.Errorf("Received error from fileMgr.CloseThisFile().\n"+
      "fileMgr='%v'\nError='%v'\n",
      fileMgr.absolutePathFileName, err.Error())
  }

  err = fileMgr.dMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by fileMgr.dMgr.DeleteAll().\n"+
      "Attempted Directory Deletion Failed!!\n"+
      "Directory=%v\nFileName='%v'\nError='%v'",
      fileMgr.absolutePathFileName,
      fileMgr.GetFileNameExt(),
      err.Error())
  }
}

func TestFileMgr_CreateDirAndFile_02(t *testing.T) {

  fh := FileHelper{}
  fileName := "TestFileMgr_CreateDirAndFile_02"

  testFile := fh.AdjustPathSlash(
    "../createFilesTest/Level01/Level02/Level03/" + fileName)

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Test Startup Error returned by "+
      "fh.DeleteDirFile(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPath"+
      "FileNameExtStr(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())

    return
  }

  fileMgr.isInitialized = false

  err = fileMgr.CreateDirAndFile()

  if err == nil {
    t.Error("Expected an error return from fileMgr.CreateDirAndFile()\n" +
      "because fileMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = fh.DeleteDirFile(testFile)
}

func TestFileMgr_CreateThisFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash(
    "../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateThisFile_01.txt")

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(testFile).\n"+
      "Attempted deletion of 'testFile' FAILED!!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())

    return
  }

  fileMgr.isInitialized = false

  err = fileMgr.CreateThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateThisFile() because " +
      "fileMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(testFile).\n"+
      "Attempted deletion of 'testFile' FAILED!!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

}

func TestFileMgr_CreateThisFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash(
    "../iDoNotExist/TestFileMgr_CreateThisFile_02.txt")

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  dirMgr := fileMgr.GetDirMgr()

  _ = dirMgr.DeleteAll()

  err = fileMgr.CreateThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.CreateThisFile() because\n" +
      "the fileMgr directory does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = dirMgr.DeleteAll()

}

func TestFileMgr_DeleteThisFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash(
    "../createFilesTest/Level01/Level02/Level03/TestFileMgr_DeleteThisFile_01.txt")

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(testFile)\n"+
      "Attempted deletion of 'testFile' FAILED!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  err = fileMgr.DeleteThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.DeleteThisFile()\n" +
      "because the 'fileMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_DeleteThisFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../checkfiles/TestFileMgr_DeleteThisFile_02.txt")

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(testFile)\n"+
      "Attempted deletion of 'testFile' FAILED!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  err = fileMgr.CreateThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.CreateThisFile().\n"+
      "File='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePathFileName(), err.Error())
  }

  err = fileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.DeleteThisFile().\n"+
      "Attempted deletion of 'fileMgr' FAILED!\n"+
      "File='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  doesThisFileExist, err := fh.DoesThisFileExist(fileMgr.GetAbsolutePathFileName())

  if err != nil {
    t.Errorf("Test Clean-Up Non-Path Error returned from path!\n"+
      "Final Deletion of 'fileMgr' FAILED!\n"+
      "fileMgr='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirFile(fileMgr.GetAbsolutePath())
    return
  }

  if doesThisFileExist {
    t.Errorf("ERROR: Final Deletion of fileMgr FAILED!\n"+
      "File='%v'\n",
      fileMgr.GetAbsolutePath())

    _ = fh.DeleteDirFile(fileMgr.GetAbsolutePath())
  }

}

func TestFileMgr_DoesFileExist_01(t *testing.T) {

  testFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr"+
      "(testFile)\ntestFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  if fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='false'\n" +
      "because the fileMgr is invalid.\n" +
      "However, the return value was 'true'!!\n")
  }
}

func TestFileMgr_DoesFileExist_02(t *testing.T) {

  testFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  if !fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='true'\n" +
      "because the 'FileMgr' file does exist.\n" +
      "However, the return value was 'false'!\n")
  }
}

func TestFileMgr_DoesThisFileExist_01(t *testing.T) {

  testFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  _, err = fileMgr.DoesThisFileExist()

  if err == nil {
    t.Error("Expected error return from fileMgr.DoesThisFileExist()\n" +
      "because the fileMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_Empty_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  fileMgr1.Empty()

  fileMgr2 := FileMgr{}
  fileMgr2.Empty()

  if !fileMgr1.Equal(&fileMgr2) {
    t.Error("Error: Expected empty fileMgr1 to equal empty fileMgr2.\n" +
      "However, THEY ARE NOT EQUAL!\n")
  }

}

func TestFileMgr_Equal_01(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  fileMgr2 := fileMgr1.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != true {
    t.Error("Expected Equal to return 'true' for fileMgr1==fileMgr1.\n" +
      "Instead, fileMgr1==fileMgr1 returned 'false'.\n")
  }

}

func TestFileMgr_Equal_02(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\FileMgmnt\\TestFile003.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("ERROR: Expected fileMgr1==fileMgr2 to return 'false'.\n" +
      "Instead, fileMgr1==fileMgr2 returned 'true'\n")
  }

}

func TestFileMgr_Equal_03(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "New(\"-rwxrwxrwx\").\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}."+
      "New(FOpenType.TypeReadWrite()).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}."+
      "New(fOpenCfg, fPermCfg).\n"+
      "Error='%v'\n", err.Error())

    return
  }

  fileMgr2.fileAccessStatus = fAccessCfg.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead fileMgr1==fileMgr2 returned 'true'.\n" +
      "The fileAccessStatus values are different.\n")
  }

}

func TestFileMgr_Equal_04(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error from FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fileMgr2.dMgr = DirMgr{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead, 'true' was returned for fileMgr1==fileMgr2.\n" +
      "Directory Managers are different.\n")
  }
}

func TestFileMgr_Equal_05(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v', Error='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fileMgr2.actualFileInfo = FileInfoPlus{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead, 'true' was returned for fileMgr1==fileMgr2.\n" +
      "acutalFileInfo's are different.\n")
  }
}

func TestFileMgr_EqualAbsPaths_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_1_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path to EQUAL "+
      "fileMgr2 absolute path.\n"+
      "However, Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualAbsPaths_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\level_1_1_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path to EQUAL "+
      "fileMgr2 absolute path.\n"+
      "However, Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}

func TestFileMgr_EqualAbsPaths_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n", relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())

    return
  }

  relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_2_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path to NOT "+
      "EQUAL fileMgr2 absolute path.\n"+
      "However, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
    return
  }

}

func TestFileMgr_EqualFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_02_dir\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_2_2_xray.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT "+
      "EQUAL fileMgr2 file name ext.\n"+
      "However, they ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.jag"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT "+
      "EQUAL fileMgr2 file name ext.\n"+
      "However, they ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualPathFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path file "+
      "name ext to EQUAL\n"+
      "fileMgr2 absolute path file name ext.\n"+
      "However, the Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path file name "+
      "ext to EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "However, the Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_02_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}

func TestFileMgr_EqualPathFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_X_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.log"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT equal fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}
