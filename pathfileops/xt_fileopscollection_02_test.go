package pathfileops

import (
  "strings"
  "testing"
)

func TestFileOpsCollection_DeleteAtIndex_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
        "i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
      return
    }

    df[i], err = fh.GetAbsPathFromFilePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  err := fOpsCol.DeleteAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fOpsCol.GetNumOfFileOps()

  if arrayLen != 4 {
    t.Errorf("Expected array length=4 after deletion. "+
      "Instead, array length='%v'", arrayLen)
  }

  for j := 0; j < arrayLen; j++ {

    fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())

      return
    }

    if sf[2] == fOps.source.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected index 2 would be deleted. I was NOT! "+
        "Found source path file name='%v' at index='%v' ",
        fOps.source.GetAbsolutePathFileName(), j)
    }

    if df[2] == fOps.destination.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected index 2 would be deleted. I was NOT! "+
        "Found destination path file name='%v' at index='%v' ",
        fOps.source.GetAbsolutePathFileName(), j)
    }

  }
}

func TestFileOpsCollection_DeleteAtIndex_02(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
        "i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
      return
    }

    df[i], err = fh.GetAbsPathFromFilePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  err := fOpsCol.DeleteAtIndex(4)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(4). "+
      "Error='%v' ", err.Error())
  }

  err = fOpsCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(0). "+
      "Error='%v' ", err.Error())
  }

  err = fOpsCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
      "Error='%v' ", err.Error())
  }

  err = fOpsCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fOpsCol.GetNumOfFileOps()

  if arrayLen != 1 {
    t.Errorf("Expected array length=1 after deletion. "+
      "Instead, array length='%v'", arrayLen)
  }

  fOps, err := fOpsCol.PeekFileOpsAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(0). "+
      "Error='%v' ", err.Error())
    return
  }

  srcFileMgr := fOps.GetSource()

  if sf[1] != srcFileMgr.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected source path file name='%v'. "+
      "Instead source path file name='%v' ",
      sf[1], srcFileMgr.GetAbsolutePathFileName())
  }

  destinationFileMgr := fOps.GetDestination()

  if df[1] != destinationFileMgr.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected destination path file name='%v'. "+
      "Instead destination path file name='%v' ",
      df[1], destinationFileMgr.GetAbsolutePathFileName())
  }

}

func TestFileOpsCollection_DeleteAtIndex_03(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }
  err := fOpsCol.DeleteAtIndex(-1)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.DeleteAtIndex(-1) because\n"+
      "the input parameter -1 is an invalid index.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOpsCollection_DeleteAtIndex_04(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }

  err := fOpsCol.DeleteAtIndex(99)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.DeleteAtIndex(-1) because\n"+
      "the input parameter '99' exceeds the internal array's upper boundary.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_DeleteAtIndex_05(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  err := fOpsCol.DeleteAtIndex(2)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.DeleteAtIndex(-1) because\n"+
      "the input parameter '99' exceeds the internal array's upper boundary.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_DeleteAtIndex_06(t *testing.T) {


  srcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"


  destFile := "../dirmgrtests/level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  err := fOpsCol.AddByPathFileNameExtStrs(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(srcFile, destFile).\n" +
      "srcFile='%v'\n" +
      "destFile='%v'\n"+
      "Error='%v'\n", srcFile, destFile, err.Error())
    return
  }

  err = fOpsCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.DeleteAtIndex(0)\n"+
      "the input parameter '99' exceeds the internal array's upper boundary.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_GetFileOpsAtIndex_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
        "i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
      return
    }

    df[i], err = fh.GetAbsPathFromFilePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  fOps, err := fOpsCol.GetFileOpsAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.GetFileOpsAtIndex(2). "+
      "Error='%v' ", err.Error())
    return
  }

  srcFile := fOps.GetSource()

  if sf[2] != srcFile.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected source file[2]='%v'. "+
      "Instead, source file[2]='%v' ", sf[2], srcFile.GetAbsolutePathFileName())
  }

  destFile := fOps.GetDestination()

  if df[2] != destFile.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected destination file[2]='%v'. "+
      "Instead, destination file[2]='%v' ", df[2], destFile.GetAbsolutePathFileName())
  }

}

func TestFileOpsCollection_ExecuteFileOperations_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  destDir := "../dirmgrtests/TestFileOpsCollection_ExecuteFileOperations_01"

  dfFileNameExt := []string{  "level_0_0_test.txt",
                              "level_0_1_test.txt",
                              "level_0_2_test.txt",
                              "level_0_3_test.txt",
                              "level_0_4_test.txt" }

  df := make([]string, 5, 10)

  df[0] =  destDir + "/" + dfFileNameExt[0]
  df[1] =  destDir + "/" + dfFileNameExt[1]
  df[2] =  destDir + "/" + dfFileNameExt[2]
  df[3] =  destDir + "/" + dfFileNameExt[3]
  df[4] =  destDir + "/" + dfFileNameExt[4]


  dstDMgr,
  err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
    return
  }

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      _ = dstDMgr.DeleteAll()
      return
    }
  }

  err = fOpsCol.ExecuteFileOperations(FileOpCode.CopySourceToDestinationByIo())

  if err != nil {
    t.Errorf("Error returned by fOpsCol.ExecuteFileOperations(" +
      "FileOpCode.CopySourceToDestinationByIo())\n" +
      "Error='%v'\n", err.Error())

    _ = dstDMgr.DeleteAll()
    return
  }

  fsc := FileSelectionCriteria{}

  destTreeInfo, err := dstDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by dstDMgr.FindFilesBySelectCriteria(fsc)\n" +
      "Error='%v'\n", err.Error())

    _ = dstDMgr.DeleteAll()
    return
  }

  numOfFiles := destTreeInfo.GetNumOfFiles()

  if numOfFiles != 5 {
    t.Errorf("ERROR: Expected the number of files in the destination directory would equal '5'.\n" +
      "Instead, number of files='%v'\n", numOfFiles)
    _ = dstDMgr.DeleteAll()
    return
  }

  foundFile := false

  for k:=0; k < 5; k++ {

    fMgr, err :=  destTreeInfo.PeekFileMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by destTreeInfo.PeekFileMgrAtIndex(%v)\n" +
        "Error='%v'\n", k, err.Error())
      _ = dstDMgr.DeleteAll()
      return
    }

    fileNameExt := strings.ToLower(fMgr.GetFileNameExt())
    foundFile = false

    for j:=0; j < 5; j++ {

      if fileNameExt == dfFileNameExt[j] {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Copied File NOT Found: %v", fileNameExt)
      _ = dstDMgr.DeleteAll()
      return
    }
  }

  err = dstDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by dstDMgr.DeleteAll()\n" +
      "Error='%v'\n", err.Error())
  }

  return
}

func TestFileOpsCollection_ExecuteFileOperations_02(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  err := fOpsCol.ExecuteFileOperations(FileOpCode.CopySourceToDestinationByIo())

  if err == nil {
    t.Error("Expected an error return from fOpsCol.ExecuteFileOperations(" +
      "FileOpCode.CopySourceToDestinationByIo())\n" +
      "because the File Operations is empty.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

