package pathfileops

import (
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

func TestFileOpsCollection_GetFileOpsAtIndex_02(t *testing.T) {

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
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i],  err.Error())
      return
    }
  }

  _, err := fOpsCol.GetFileOpsAtIndex(99)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.GetFileOpsAtIndex(99)\n" +
      "because the index, '99', is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

  _, err = fOpsCol.GetFileOpsAtIndex(-1)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.GetFileOpsAtIndex(-1)\n" +
      "because the index, '-1', is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_GetFileOpsAtIndex_03(t *testing.T) {

  fOpsCol := FileOpsCollection{}
  fOpsCol.fileOps = nil

  _, err := fOpsCol.GetFileOpsAtIndex(0)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.GetFileOpsAtIndex(-1)\n" +
      "because the the File Manager Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_GetNumOfFileOps_01(t *testing.T) {

  fOpsCol := FileOpsCollection{}
  fOpsCol.fileOps = nil



  if fOpsCol.GetNumOfFileOps() != 0 {

    t.Errorf("Expected fOpsCol.GetNumOfFileOps() == '0'\n" +
      "Instead, fOpsCol.GetNumOfFileOps()='%v'\n", fOpsCol.GetNumOfFileOps())
  }
}
