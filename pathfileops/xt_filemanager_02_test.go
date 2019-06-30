package pathfileops

import "testing"

func TestFileMgr_CopyFileMgrByLink_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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

  newFileMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "+
      "\"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())

    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
      return
    }
  }

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByLink(&newFileMgr). "+
      "newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
    _ = newFileMgr.DeleteThisFile()
    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileMgrByLink_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

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
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
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

  newFileMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "+
      "\"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePathFileName(), err.Error())
    }

    return
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by newFileMgr.DeleteThisFile().\n"+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePathFileName(), err.Error())
  }

}

func TestFileMgr_CopyFileMgrByLink_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "adjustedSourceFile='%v'\nError='%v' ",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v' ",
      absoluteSourceFile, err.Error())
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

  newFileMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist().\n"+
      "newFileMgr='%v'\nError='%v'",
      newFileMgr.GetAbsolutePath(), err.Error())

    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
      return
    }
  }

  newFileMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
      "because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

  }

  newFileMgr.isInitialized = true

}

func TestFileMgr_CopyFileMgrByLink_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "adjustedSourceFile='%v'\nError='%v' ",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  err = srcFMgr.CopyFileMgrByLink(nil)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(nil) " +
      "because nil was passed to this method. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLink_05(t *testing.T) {

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

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(nil) " +
      "because srcFMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLink_06(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "adjustedSourceFile='%v'\nError='%v'\n",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  destFileMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByLink(&destFileMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&destFileMgr) " +
      "because srcFMgr and destFileMgr are equivalent. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLinkByIo_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "adjustedSourceFile='%v'\nError='%v' ",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "adjustedSourceFile='%v'\nError='%v' ",
      adjustedSourceFile, err.Error())

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

  newFileMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "destDMgr, \"newerFileForTest_01.txt\"). "+
      "destDMgr='%v'\nError='%v'",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "newFileMgr='%v'\nError='%v'",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
      return
    }
  }

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr). "+
      "newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed!\nSrc File='%v'\nDest File='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted!\nNewFile := '%v'\n",
      newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "adjustedSourceFile='%v'\nError='%v'\n",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "adjustedSourceFile='%v'\nError='%v'\n",
      adjustedSourceFile, err.Error())
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

  newFileMgr, err :=
    FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'\n",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile().\n"+
        "newFileMgr='%v'\nError='%v'\n",
        newFileMgr.GetAbsolutePathFileName(), err.Error())
      return
    }
  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_03(t *testing.T) {

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
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "destDMgr, \"newerFileForTest_01.txt\").\n"+
      "destDMgr='%v'\nError='%v'",
      destDMgr.GetAbsolutePath(), err.Error())
    return
  }

  doesFileExist, err := newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "newFileMgr='%v'\nError='%v'\n",
      newFileMgr.GetAbsolutePath(), err.Error())
    return
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile().\n"+
        "Error='%v'\n",
        newFileMgr.GetAbsolutePathFileName())
      return
    }
  }

  newFileMgr.isInitialized = false

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "adjustedSourceFile='%v'\nError='%v'\n",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "absoluteSourceFile='%v'\nError='%v' ",
      absoluteSourceFile, err.Error())
    return
  }

  err = srcFMgr.CopyFileMgrByLinkByIo(nil)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(nil) " +
      "because nil was passed to the method. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrByLinkByIo_05(t *testing.T) {

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

  newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, sourceFile)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
      "Error='%v'", err.Error())
  }

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because srcFMgr does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileMgrByLinkByIo_06(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}
  adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
  absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile).\n"+
      "adjustedSourceFile='%v'\nError='%v'\n",
      adjustedSourceFile, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile).\n"+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  newFileMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
      "because srcFMgr is equivalent to newFileMgr. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileMgrWithBuffer_01(t *testing.T) {

  expectedSrcFileNameExt := "level_0_3_test.txt"
  expectedDestFileNameExt := "TestFileMgr_CopyFileMgrWithBuffer_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}."+
      "NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedSrcFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt"+
      "(destDMgr, expectedDestFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedDestFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedDestFileNameExt, err.Error())
    return
  }

  localBuffer := make([]byte, 5000)

  err = srcFMgr.CopyFileMgrWithBuffer(&destFMgr, localBuffer)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileMgrWithBuffer"+
      "(&destFMgr, localBuffer).\n"+
      "srcFMgr.absolutePathFileName='%v'\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
    return
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true.\n"+
      "Instead it was 'false'!\ndestFMgr.absolutePathFileName='%v'\n",
      destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.\n" +
      "ERROR! Instead, destFMgr.doesAbsolutePathFileNameExist='false'\n")
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName)\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      destFMgr.absolutePathFileName, err.Error())
  }
}

func TestFileMgr_CopyFileMgrWithBuffer_02(t *testing.T) {

  expectedSrcFileNameExt := "iDoNotExist.txt"
  expectedDestFileNameExt := "TestFileMgr_CopyFileMgrWithBuffer_02.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../createFilesTest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, expectedSrcFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedSrcFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n",
      rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(destDMgr, expectedDestFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedDestFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedDestFileNameExt, err.Error())
    return
  }

  localBuffer := make([]byte, 5000)

  err = srcFMgr.CopyFileMgrWithBuffer(&destFMgr, localBuffer)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr, localBuffer)\n" +
      "because 'srcFMgr' file DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName)\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      destFMgr.absolutePathFileName, err.Error())
  }
}

func TestFileMgr_CopyFileMgrWithBuffer_03(t *testing.T) {

  expectedSrcFileNameExt := "level_0_3_test.txt"
  expectedDestFileNameExt := "TestFileMgr_CopyFileMgrWithBuffer_01.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt). "+
      "dMgr.absolutePath='%v' expectedSrcFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt). "+
      "destDMgr.absolutePath='%v'  expectedDestFileNameExt='%v'   Error='%v'",
      destDMgr.absolutePath, expectedDestFileNameExt, err.Error())
    return
  }

  err = srcFMgr.CopyFileMgrWithBuffer(&destFMgr, nil)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr, localBuffer)\n" +
      "because 'localBuffer' is 'nil'!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName)\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      destFMgr.absolutePathFileName, err.Error())
  }
}

func TestFileMgr_CopyFileMgrWithBuffer_04(t *testing.T) {

  expectedSrcFileNameExt := "level_0_3_test.txt"
  expectedDestFileNameExt := "TestFileMgr_CopyFileMgrWithBuffer_04.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt). "+
      "dMgr.absolutePath='%v' expectedSrcFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). "+
      "rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt)

  if err != nil {
    t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt). "+
      "destDMgr.absolutePath='%v'  expectedDestFileNameExt='%v'   Error='%v'",
      destDMgr.absolutePath, expectedDestFileNameExt, err.Error())
    return
  }

  localBuffer := make([]byte, 0)

  err = srcFMgr.CopyFileMgrWithBuffer(&destFMgr, localBuffer)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr, localBuffer)\n" +
      "because 'localBuffer' is is a zero length byte array!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName)\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      destFMgr.absolutePathFileName, err.Error())
  }
}

func TestFileMgr_CopyFileMgrWithBuffer_05(t *testing.T) {

  expectedSrcFileNameExt := "level_0_3_test.txt"

  setupDestFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  expectedDestFileNameExt := "TestFileMgr_CopyFileMgrWithBuffer_05.txt"

  fh := FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest")

  dMgr, err := DirMgr{}.New(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n", adjustedPath, err.Error())
    return
  }

  srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedSrcFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(dMgr, expectedSrcFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedSrcFileNameExt='%v'\nError='%v'\n",
      dMgr.absolutePath, adjustedPath, err.Error())
    return
  }

  rawDestPath := "../checkfiles/checkfiles02"

  destDMgr, err := DirMgr{}.New(rawDestPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath).\n"+
      "rawDestPath='%v'\nError='%v'\n", rawDestPath, err.Error())
    return
  }

  destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedDestFileNameExt)

  if err != nil {
    t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt"+
      "(destDMgr, expectedDestFileNameExt).\n"+
      "destDMgr.absolutePath='%v'\nexpectedDestFileNameExt='%v'\nError='%v'\n",
      destDMgr.absolutePath, expectedDestFileNameExt, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupDestFile, destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error retrned from fh.CopyFileByIo"+
      "(setupDestFile, destFMgr.absolutePathFileName)\n"+
      "setupDestFile='%v'\ndestFMgr.absolutePathFileName='%v'\nError='%v'\n",
      setupDestFile, destFMgr.absolutePathFileName, err.Error())
    return
  }

  localBuffer := make([]byte, 5000)

  err = srcFMgr.CopyFileMgrWithBuffer(&destFMgr, localBuffer)

  if err != nil {
    t.Errorf("Error returned from srcFMgr."+
      "CopyFileMgrWithBuffer(&destFMgr, localBuffer).\n"+
      "srcFMgr.absolutePathFileName='%v'\ndestFMgr.absolutePathFileName='%v'\nError='%v'\n",
      srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
    _ = fh.DeleteDirFile(destFMgr.absolutePathFileName)
    return
  }

  if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
    t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true.\n"+
      "Instead it was 'false' destFMgr.absolutePathFileName='%v'\n",
      destFMgr.absolutePathFileName)
  }

  if !destFMgr.doesAbsolutePathFileNameExist {
    t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.\n" +
      "ERROR! Instead, destFMgr.doesAbsolutePathFileNameExist='false'\n")
  }

  err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName)\n"+
      "destFMgr.absolutePathFileName='%v'\nError='%v'\n",
      destFMgr.absolutePathFileName, err.Error())
  }
}
