package pathfileops

import "testing"

func TestFileMgr_CopyFileStrByIoByLink_01(t *testing.T) {

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
    t.Errorf("Non-Path Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByIoByLink_01.txt"

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
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
        "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

      return
    }
  }

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Clean-Up Error returned by fh.DoesThisFileExist(rawAbsDestPath)\n"+
      "rawAbsDestPath='%v'\nError='%v'\n", rawAbsDestPath, err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
  }

}

func TestFileMgr_CopyFileStrByIoByLink_02(t *testing.T) {

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
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST!\n"+
      "Source File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByIoByLink_02.txt"

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

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_03(t *testing.T) {

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
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist().\n"+
      "srcFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByIoByLink("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(\"\")\n" +
      "because destination file path is empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

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

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByIoByLink_04.txt"

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
        "rawAbsDestPath='%v'\nError='%v'\n",
        rawAbsDestPath, err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)\n" +
      "because the source file (srcFMgr) does NOT exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_05(t *testing.T) {

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

  destFilePath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByIoByLink(destFilePath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)\n" +
      "because the source file is equivalent to destination file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileStrByLink_01(t *testing.T) {

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

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLink_01.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath).\n"+
      "rawRelDestPath='%v'\nError='%v'\n",
      rawRelDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Error returned by fh.DoesThisFileExist"+
      "(rawAbsDestPath)\nrawAbsDestPath='%v'\nError='%v'",
      rawAbsDestPath, err.Error())
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

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByLink(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Clean-Up Error returned by fh.DoesThisFileExist(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByLink_02(t *testing.T) {

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
      "Error='%v' ", err.Error())
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
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLink_02.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
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

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileStrByLink_03(t *testing.T) {

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
      "Error='%v' ", err.Error())
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

  err = srcFMgr.CopyFileStrByLink("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath)\n" +
      "because destination file path is empty string.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileStrByLink_04(t *testing.T) {

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
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "absoluteSourceFile='%v'\nError='%v'\n",
      absoluteSourceFile, err.Error())
    return
  }

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLink_04.txt"

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
      t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath).\n"+
        "rawAbsDestPath='%v'\nError='%v'\n",
        rawAbsDestPath, err.Error())
      return
    }
  }

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath)\n" +
      "because source file does NOT exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileStrByLink_05(t *testing.T) {

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

  rawAbsDestPath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath)\n" +
      "because source file is equivalent to destination file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_CopyFileStrByLinkByIo_01(t *testing.T) {

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

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLinkByIo_01.txt"

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

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByLink(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  doesFileExist, err = fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Clean-Up Error returned by fh.DoesThisFileExist(rawAbsDestPath).\n"+
      "rawAbsDestPath='%v'\nError='%v'\n",
      rawAbsDestPath, err.Error())
    return
  }

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_02(t *testing.T) {

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

  rawRelDestPath := "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLinkByIo_02.txt"

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

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)\n" +
      "because srcFMgr.isInitialized == false.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_03(t *testing.T) {

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

  err = srcFMgr.CopyFileStrByLinkByIo("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)\n" +
      "because destination file path is empty string.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_04(t *testing.T) {

  sourceFile :=
    "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

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

  rawRelDestPath :=
    "../checkfiles/checkfiles02/TestFileMgr_CopyFileStrByLinkByIo_04.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist, err := fh.DoesThisFileExist(rawAbsDestPath)

  if err != nil {
    t.Errorf("Non-Path Error returned by fh.DoesThisFileExist"+
      "(rawAbsDestPath)\nrawAbsDestPath='%v'\nError='%v'",
      rawAbsDestPath, err.Error())
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

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)\n" +
      "because source file does NOT exist.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation!\n"+
      "Destination File='%v'\n", rawAbsDestPath)
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_05(t *testing.T) {

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

  rawAbsDestPath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)\n" +
      "because destination file path is empty string.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}
