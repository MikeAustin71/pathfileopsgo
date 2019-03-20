package pathfileops

import (
  "testing"
)

func TestFileMgr_CopyFileStrByIoByLink_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
      "Destination File='%v'",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
      "Destination File='%v' ", rawAbsDestPath)
  }

}

func TestFileMgr_CopyFileStrByIoByLink_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByIoByLink("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist := fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByIoByLink_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  destFilePath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByIoByLink(destFilePath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
      "because source file is equivalent to destination file. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByLink_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByLink(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
      "Destination File='%v'",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
      "Destination File='%v' ", rawAbsDestPath)
  }

}

func TestFileMgr_CopyFileStrByLink_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByLink_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByLink("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByLink_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist := fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByLink_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawAbsDestPath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
      "because source file is equivalent to destination file. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByLinkByIo_01(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err != nil {
    t.Errorf("Error returned by srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if !doesFileExist {
    t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
      "Destination File='%v'",
      rawAbsDestPath)
    return
  }

  err = fh.DeleteDirFile(rawAbsDestPath)

  if err != nil {
    t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
      "Destination File='%v' ", rawAbsDestPath)
  }

}

func TestFileMgr_CopyFileStrByLinkByIo_02(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist = fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
      "because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileStrByLinkByIo_03(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  doesFileExist, err := srcFMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
      "Error='%v' ", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
      srcFMgr.GetAbsolutePathFileName())
    return
  }

  err = srcFMgr.CopyFileStrByLinkByIo("")

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_04(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

  rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
      "rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
  }

  doesFileExist := fh.DoesFileExist(rawAbsDestPath)

  if doesFileExist {
    err = fh.DeleteDirFile(rawAbsDestPath)
    t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
      "rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

  }

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
      "because source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileStrByLinkByIo_05(t *testing.T) {

  sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fh := FileHelper{}

  absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
      "Error='%v' ", err.Error())
  }

  srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
      "Error='%v' ", err.Error())
  }

  rawAbsDestPath := srcFMgr.GetAbsolutePathFileName()

  err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

  if err == nil {
    t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
      "because destination file path is empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileMgr_CopyFileToDirByIo_01(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByIo(destDMgr). "+
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

func TestFileMgr_CopyFileToDirByIo_02(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIo_03(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIo_04(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
      "destination directory is equivalent to source directory. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIo_05(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
      "srcFMgr is invalid. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIo_06(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIo(destDMgr)

  if err == nil {
    t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
      "source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_01(t *testing.T) {

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
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if doesFileExist {

    err = newFileMgr.DeleteThisFile()

    if err != nil {
      t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
        newFileMgr.GetAbsolutePathFileName())
    }
  }

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err != nil {
    t.Errorf("Error returned from srcFMgr.CopyFileToDirByIoByLink(destDMgr). "+
      "destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
  }

  doesFileExist, err = newFileMgr.DoesThisFileExist()

  if err != nil {
    t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
      "Error='%v'", err.Error())
  }

  if !doesFileExist {
    t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
      srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
    return
  }

  err = newFileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Expected that newly copied file would be deleted. "+
      "Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_02(t *testing.T) {

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

  srcFMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_03(t *testing.T) {

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

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
      "srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_04(t *testing.T) {

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

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
      "srcFMgr directory is equivalent to destination directory. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileMgr_CopyFileToDirByIoByLink_05(t *testing.T) {

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

  destDMgr.isInitialized = false

  err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

  if err == nil {
    t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
      "the source file does NOT exist. However, NO ERROR WAS RETURNED!")
  }

}
