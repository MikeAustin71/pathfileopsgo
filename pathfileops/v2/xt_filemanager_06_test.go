package pathfileops

import (
  "strings"
  "testing"
)

func TestFileMgr_CopyFileToDirByLinkByIo_01(t *testing.T) {

  fileName := "newerFileForTest_01.txt"
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  sourceFile := "../../filesfortest/newfilesfortest/iDoNotExist.txt"

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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
  sourceFile := "../../filesfortest/newfilesfortest/" + fileName

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

  sourceFile := "../../filesfortest/newfilesfortest/iDoNotExist.txt"

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

  rawDestPath := fh.AdjustPathSlash("../../checkfiles/checkfiles02")

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

func TestFileMgr_CopyFromStrings_01(t *testing.T) {

  fh := FileHelper{}

  sourceFile := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  sourceFile = fh.AdjustPathSlash(sourceFile)

  destFile := "../../createFilesTest/TestFileMgr_CopyFromStrings_01.txt"

  destFile = fh.AdjustPathSlash(destFile)

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  fMgrSrc, fMgrDest, err := FileMgr{}.CopyFromStrings(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}."+
      "CopyFromStrings(sourceFile, destFile)\n"+
      "sourceFile='%v'\ndestFile='%v'\nError='%v'",
      sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(destFile)
    return
  }

  absSourcePath, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh."+
      "MakeAbsolutePath(sourceFile)\n"+
      "sourceFile='%v',Error='%v'\n",
      sourceFile, err.Error())
    _ = fh.DeleteDirFile(destFile)
    return
  }

  absSourcePath = strings.ToLower(absSourcePath)

  if absSourcePath != strings.ToLower(fMgrSrc.absolutePathFileName) {
    t.Errorf("Error: Expected source path and file name are NOT EQUAL\n"+
      "to actual source path and file name!\n"+
      "Expected source file='%v'\n"+
      "Actual source file='%v'\n",
      absSourcePath, strings.ToLower(fMgrSrc.absolutePathFileName))
  }

  absDestPath, err := fh.MakeAbsolutePath(destFile)

  if err != nil {
    t.Errorf("Error returned by fh."+
      "MakeAbsolutePath(destFile)\n"+
      "destFile='%v',Error='%v'\n",
      destFile, err.Error())
    _ = fh.DeleteDirFile(destFile)
    return
  }

  absDestPath = strings.ToLower(absDestPath)

  if absDestPath != strings.ToLower(fMgrDest.absolutePathFileName) {
    t.Errorf("Error: Expected destination path and file name are NOT EQUAL\n"+
      "to actual destination path and file name!\n"+
      "Expected destination file='%v'\n"+
      "Actual destination file='%v'\n",
      absDestPath, strings.ToLower(fMgrDest.absolutePathFileName))
  }

  if !fh.DoesFileExist(absSourcePath) {
    t.Errorf("Error: Source File DOES NOT EXIST!\n"+
      "Source File='%v'\n", absSourcePath)
    _ = fh.DeleteDirFile(absDestPath)
    return
  }

  if !fh.DoesFileExist(absDestPath) {
    t.Errorf("Error: After Copy Operation Destination "+
      "File DOES NOT EXIST!\n"+
      "Destination File='%v'\n", absDestPath)
    return
  }

  if !fMgrSrc.DoesFileExist() {
    t.Errorf("Error returned by fMgrSrc.DoesFileExist()\n"+
      "Source File DOES NOT EXIST!\n"+
      "Source File='%v'", fMgrSrc.absolutePathFileName)
  }

  if !fMgrDest.DoesFileExist() {
    t.Errorf("Error returned by fMgrDest.DoesFileExist()\n"+
      "Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'", fMgrDest.absolutePathFileName)
  }

  err = fh.DeleteDirFile(absDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(absDestPath)\n"+
      "absDestPath='%v'\nError='%v'\n",
      absDestPath, err.Error())
  }

}

func TestFileMgr_CopyFromStrings_02(t *testing.T) {

  fh := FileHelper{}

  sourceFile := "../../checkfiles/iDoNotExist.txt"

  sourceFile = fh.AdjustPathSlash(sourceFile)

  destFile := "../../createFilesTest/TestFileMgr_CopyFromStrings_02.txt"

  destFile = fh.AdjustPathSlash(destFile)

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  _, _, err = FileMgr{}.CopyFromStrings(sourceFile, destFile)

  if err == nil {
    t.Error("Expected Error return from FileMgr{}." +
      "CopyFromStrings(sourceFile, destFile)\n" +
      "because 'sourceFile' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }

}

func TestFileMgr_CopyFromStrings_03(t *testing.T) {

  fh := FileHelper{}

  sourceFile := ""

  destFile := "../../createFilesTest/TestFileMgr_CopyFromStrings_03.txt"

  destFile = fh.AdjustPathSlash(destFile)

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  _, _, err = FileMgr{}.CopyFromStrings(sourceFile, destFile)

  if err == nil {
    t.Error("Expected Error return from FileMgr{}." +
      "CopyFromStrings(sourceFile, destFile)\n" +
      "because 'sourceFile' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
  }
}

func TestFileMgr_CopyFromStrings_04(t *testing.T) {

  fh := FileHelper{}

  sourceFile := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  sourceFile = fh.AdjustPathSlash(sourceFile)

  destFile := ""

  _, _, err := FileMgr{}.CopyFromStrings(sourceFile, destFile)

  if err == nil {
    t.Error("Expected Error return from FileMgr{}." +
      "CopyFromStrings(sourceFile, destFile)\n" +
      "because 'destFile' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("Error: Source File DOES NOT EXIST!\n"+
      "sourceFile='%v'\n",
      sourceFile)
  }
}

func TestFileMgr_CreateDir_01(t *testing.T) {

  fh := FileHelper{}

  pathFileName :=
    "../../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateDir_01.txt"

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
    "../../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateDirAndFile_01.txt")

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
    "../../createFilesTest/Level01/Level02/Level03/" + fileName)

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
    "../../createFilesTest/Level01/Level02/Level03/TestFileMgr_CreateThisFile_01.txt")

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
    "../../iDoNotExist/TestFileMgr_CreateThisFile_02.txt")

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