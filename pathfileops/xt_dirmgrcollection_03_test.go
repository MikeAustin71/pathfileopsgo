package pathfileops

import (
  "strings"
  "testing"
)

func TestDirMgrCollection_GetFileMgrAtIndex_01(t *testing.T) {

  df := make([]string, 5, 10)

  df[0] = "..\\dirmgrtests"
  df[1] = "..\\dirmgrtests\\dir01"
  df[2] = "..\\dirmgrtests\\dir01\\dir02"
  df[3] = "..\\dirmgrtests\\dir01\\dir02\\dir03"
  df[4] = "..\\dirmgrtests\\dir01\\dir02\\dir03\\dir04"

  dmgrCol := DirMgrCollection{}.New()

  fh := FileHelper{}

  var err error

  for i := 0; i < 5; i++ {

    err = dmgrCol.AddDirMgrByPathNameStr(df[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
    }

    df[i], err = fh.MakeAbsolutePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
    }

  }

  dirMgr, err := dmgrCol.GetDirMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.GetDirMgrAtIndex(2). "+
      "Error='%v' ", err.Error())
    return
  }

  if df[2] != dirMgr.GetAbsolutePath() {
    t.Errorf("Error: Expected dirMgr[2]='%v'. "+
      "Instead, dirMgr[2]='%v' ", df[2], dirMgr.GetAbsolutePath())
  }

}

func TestDirMgrCollection_GetDirMgrAtIndex_01(t *testing.T) {

  dirStr := []string{
    "..\\dirmgrtests",
    "..\\dirmgrtests\\dir01",
    "..\\dirmgrtests\\dir01\\dir02",
    "..\\dirmgrtests\\dir01\\dir02\\dir03" }


  dmgrCol := DirMgrCollection{}.New()

  for i:=0; i < 4; i++ {

    err := dmgrCol.AddDirMgrByPathNameStr(dirStr[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(dirStr[%v]).\n" +
        "dirStr[%v]='%v'\n"+
        "Error='%v'\n", i, i, dirStr[i], err.Error())
      return
    }

  }

  _, err := dmgrCol.GetDirMgrAtIndex(-1)

  if err == nil {
    t.Error("ERROR: Expected an error return from dmgrCol.GetDirMgrAtIndex(-1)\n" +
      "because the index, '-1' is less than zero!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgrCollection_GetDirMgrAtIndex_02(t *testing.T) {

  dirStr := []string{
    "..\\dirmgrtests",
    "..\\dirmgrtests\\dir01",
    "..\\dirmgrtests\\dir01\\dir02",
    "..\\dirmgrtests\\dir01\\dir02\\dir03" }


  dmgrCol := DirMgrCollection{}.New()

  for i:=0; i < 4; i++ {

    err := dmgrCol.AddDirMgrByPathNameStr(dirStr[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(dirStr[%v]).\n" +
        "dirStr[%v]='%v'\n"+
        "Error='%v'\n", i, i, dirStr[i], err.Error())
      return
    }

  }

  _, err := dmgrCol.GetDirMgrAtIndex(99)

  if err == nil {
    t.Error("ERROR: Expected an error return from dmgrCol.GetDirMgrAtIndex(99)\n" +
      "because the index, '99' exceed the collection's maximum array index.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgrCollection_GetDirMgrAtIndex_03(t *testing.T) {

  dirStr := []string{
    "..\\dirmgrtests",
    "..\\dirmgrtests\\dir01",
    "..\\dirmgrtests\\dir01\\dir02",
    "..\\dirmgrtests\\dir01\\dir02\\dir03" }

  expectedFirstDir, err := FileHelper{}.MakeAbsolutePath(dirStr[0])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirStr[0])\n" +
      "dirStr[0]='%v'\nError='%v'\n",
      dirStr[0], err.Error())
  }

  expectedFirstDir = strings.ToLower(expectedFirstDir)

  dmgrCol := DirMgrCollection{}.New()

  for i:=0; i < 4; i++ {

    err := dmgrCol.AddDirMgrByPathNameStr(dirStr[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(dirStr[%v]).\n" +
        "dirStr[%v]='%v'\n"+
        "Error='%v'\n", i, i, dirStr[i], err.Error())
      return
    }

  }

  dMgrPtr, err := dmgrCol.GetDirMgrAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.GetDirMgrAtIndex(0).\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFirstDir != strings.ToLower(dMgrPtr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected returned dMgr='%v'\n" +
      "Instead, returned dMgr='%v'\n",
      expectedFirstDir, strings.ToLower(dMgrPtr.GetAbsolutePath()))
  }
}

func TestDirMgrCollection_GetDirMgrAtIndex_04(t *testing.T) {

  dirStr := []string{
    "..\\dirmgrtests",
    "..\\dirmgrtests\\dir01",
    "..\\dirmgrtests\\dir01\\dir02",
    "..\\dirmgrtests\\dir01\\dir02\\dir03" }

  expectedLastDir, err := FileHelper{}.MakeAbsolutePath(dirStr[3])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirStr[3])\n" +
      "dirStr[0]='%v'\nError='%v'\n",
      dirStr[0], err.Error())
  }

  expectedLastDir = strings.ToLower(expectedLastDir)

  dMgrCol := DirMgrCollection{}.New()

  for i:=0; i < 4; i++ {

    err := dMgrCol.AddDirMgrByPathNameStr(dirStr[i])

    if err != nil {
      t.Errorf("Error returned by dMgrCol.AddDirMgrByPathNameStr(dirStr[%v]).\n" +
        "dirStr[%v]='%v'\n"+
        "Error='%v'\n", i, i, dirStr[i], err.Error())
      return
    }

  }

  dMgrPtr, err := dMgrCol.GetDirMgrAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by dMgrCol.GetDirMgrAtIndex(3).\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedLastDir != strings.ToLower(dMgrPtr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected returned dMgr='%v'\n" +
      "Instead, returned dMgr='%v'\n",
      expectedLastDir, strings.ToLower(dMgrPtr.GetAbsolutePath()))
  }
}

func TestDirMgrCollection_GetDirMgrAtIndex_05(t *testing.T) {

  dirStr := []string{
    "..\\dirmgrtests",
    "..\\dirmgrtests\\dir01",
    "..\\dirmgrtests\\dir01\\dir02",
    "..\\dirmgrtests\\dir01\\dir02\\dir03" }

  expectedDirStr, err := FileHelper{}.MakeAbsolutePath(dirStr[2])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirStr[3])\n" +
      "dirStr[0]='%v'\nError='%v'\n",
      dirStr[0], err.Error())
  }

  expectedDirStr = strings.ToLower(expectedDirStr)

  dMgrCol := DirMgrCollection{}.New()

  for i:=0; i < 4; i++ {

    err := dMgrCol.AddDirMgrByPathNameStr(dirStr[i])

    if err != nil {
      t.Errorf("Error returned by dMgrCol.AddDirMgrByPathNameStr(dirStr[%v]).\n" +
        "dirStr[%v]='%v'\n"+
        "Error='%v'\n", i, i, dirStr[i], err.Error())
      return
    }

  }

  dMgrPtr, err := dMgrCol.GetDirMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dMgrCol.GetDirMgrAtIndex(2).\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedDirStr != strings.ToLower(dMgrPtr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected returned dMgr='%v'\n" +
      "Instead, returned dMgr='%v'\n",
      expectedDirStr, strings.ToLower(dMgrPtr.GetAbsolutePath()))
  }
}

func TestDirMgrCollection_GetDirMgrAtIndex_06(t *testing.T) {

  dMgrCol := DirMgrCollection{}

  _, err := dMgrCol.GetDirMgrAtIndex(2)

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrCol.GetDirMgrAtIndex(2)\n" +
      "because 'dMgrCol' is empty.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}


func TestDirMgrCollection_GetDirMgrArray_01(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgr, err := dMgrs.PeekLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != origPath {
    t.Errorf("Expected Last DirMgr path='%v'.\n"+
      "Instead, dMgr.path='%v'\n", origPath, dMgr.path)
  }

  if dMgr.absolutePath != origAbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'.\n"+
      "Instead, dMgr.absolutePath='%v'\n",
      origAbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

  cntr := 0

  for _, dirMgr := range dMgrs.GetDirMgrArray() {

    err := dirMgr.IsDirMgrValid("TestDirMgrCollection_GetDirMgrArray_01 Error")

    if err != nil {
      t.Errorf("Invalid DirMgr: Dir='%v'  Error=%v", dirMgr.GetAbsolutePath(), err.Error())
    }

    cntr++
  }

  if cntr != 6 {
    t.Errorf("Expected Diretory Count='6'. Instead Directory Count='%v'", cntr)
  }

}

func TestDirMgrCollection_GetDirMgrArray_02(t *testing.T) {
  dMgrs := DirMgrCollection{}
  dMgrs.dirMgrs = nil

  dMgrArray := dMgrs.GetDirMgrArray()

  if dMgrArray == nil {

    t.Error("ERROR: dMgrArray := dMgrs.GetDirMgrArray().\n" +
      "dMgrArray is nil!!!")

    return
  }

  if len(dMgrArray) != 0 {
    t.Errorf("ERROR: Length of dMgrArray is NOT equal to zero!\n" +
      "dMgrArray := dMgrs.GetDirMgrArray().\n" +
      "len(dMgrArray)='%v'", len(dMgrArray))
  }

}

func TestDirMgrCollection_GetNumOfDirs(t *testing.T) {

  dMgrCol := DirMgrCollection{}

  dMgrCol.dirMgrs = nil

  numOfDirectories := dMgrCol.GetNumOfDirs()

  if numOfDirectories != 0 {
    t.Errorf("ERROR: Expected returned number of directories would equal ZERO\n" +
      "because 'dMgrCol' is empty.\n" +
      "Instead, number of directories='%v'\n", numOfDirectories)
  }
}


func TestDirMgrCollection_InsertDirMgrAtIndex_01(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected DirMgrsCollection Array Length = '6'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  // Insert At Index 3

  origPath = fh.AdjustPathSlash("../testfiles/testfiles2")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (7) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertDMgr, err := DirMgr{}.New(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origAbsPath). "+
      "origAbsPath='%v' Error='%v' ", origAbsPath, err.Error())
  }

  err = dMgrs.InsertDirMgrAtIndex(insertDMgr, 3)

  if err != nil {
    t.Errorf("Error returned by dMgrs.InsertDirMgrAtIndex(insertDMgr, 3). "+
      "insertDMgr='%v' Error='%v' ", insertDMgr, err.Error())
  }

  if dMgrs.GetNumOfDirs() != 7 {
    t.Errorf("Expected DirMgrsCollection Array Length = '7'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  dMgrAtIdx, err := dMgrs.PeekDirMgrAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekDirMgrAtIndex(3). "+
      "Error='%v' ", err.Error())

  }

  if !insertDMgr.Equal(&dMgrAtIdx) {
    t.Error("Error: Expected Directory Manager at index 3 should be equal to " +
      "inserted Directory Manager. THEY ARE NOT EQUAL!")
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_02(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected DirMgrsCollection Array Length = '6'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  // Insert At Index 10 - Exceeds array index boundary

  origPath = fh.AdjustPathSlash("../testfiles/testfiles2")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (7) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertDMgr, err := DirMgr{}.New(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origAbsPath). "+
      "origAbsPath='%v' Error='%v' ", origAbsPath, err.Error())
  }

  err = dMgrs.InsertDirMgrAtIndex(insertDMgr, 10)

  if err != nil {
    t.Errorf("Error returned by dMgrs.InsertDirMgrAtIndex(insertDMgr, 3). "+
      "insertDMgr='%v' Error='%v' ", insertDMgr, err.Error())
  }

  if dMgrs.GetNumOfDirs() != 7 {
    t.Errorf("Expected DirMgrsCollection Array Length = '7'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  dMgrAtIdx, err := dMgrs.PeekDirMgrAtIndex(6)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekDirMgrAtIndex(6). "+
      "Error='%v' ", err.Error())

  }

  if !insertDMgr.Equal(&dMgrAtIdx) {
    t.Error("Error: Expected Directory Manager at index 3 should be equal to " +
      "inserted Directory Manager. THEY ARE NOT EQUAL!")
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_03(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (5) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (6) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected DirMgrsCollection Array Length = '6'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  // Insert At Index 0 - First entry in array

  origPath = fh.AdjustPathSlash("../testfiles/testfiles2")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (7) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertDMgr, err := DirMgr{}.New(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origAbsPath). "+
      "origAbsPath='%v' Error='%v' ", origAbsPath, err.Error())
  }

  err = dMgrs.InsertDirMgrAtIndex(insertDMgr, 0)

  if err != nil {
    t.Errorf("Error returned by dMgrs.InsertDirMgrAtIndex(insertDMgr, 3). "+
      "insertDMgr='%v' Error='%v' ", insertDMgr, err.Error())
  }

  if dMgrs.GetNumOfDirs() != 7 {
    t.Errorf("Expected DirMgrsCollection Array Length = '7'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  dMgrAtIdx, err := dMgrs.PeekDirMgrAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekDirMgrAtIndex(0). "+
      "Error='%v' ", err.Error())

  }

  if !insertDMgr.Equal(&dMgrAtIdx) {
    t.Error("Error: Expected Directory Manager at index 3 should be equal to " +
      "inserted Directory Manager. THEY ARE NOT EQUAL!")
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_04(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected DirMgrsCollection Array Length = '6'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  // Insert At Index -4

  origPath = fh.AdjustPathSlash("../testfiles/testfiles2")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (7) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertDMgr, err := DirMgr{}.New(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origAbsPath). "+
      "origAbsPath='%v' Error='%v' ", origAbsPath, err.Error())
  }

  err = dMgrs.InsertDirMgrAtIndex(insertDMgr, -4)

  if err == nil {
    t.Error("Error: Expected an error return from dMgrs.InsertDirMgrAtIndex(insertDMgr, -4). " +
      "NO ERROR RETURNED!")
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_05(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 1 {
    t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
    t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ",
      origAbsPath, dMgrs.dirMgrs[0].absolutePath)
  }

  // # 2
  origPath = fh.AdjustPathSlash("../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected DirMgrsCollection Array Length = '6'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  // Insert At Index 3

  origPath = fh.AdjustPathSlash("../testfiles/testfiles2")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (7) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'",
      origPath, err.Error())
  }

  insertDMgr, err := DirMgr{}.New(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origAbsPath). "+
      "origAbsPath='%v' Error='%v' ", origAbsPath, err.Error())
  }

  err = dMgrs.InsertDirMgrAtIndex(insertDMgr, 4)

  if err != nil {
    t.Errorf("Error returned by dMgrs.InsertDirMgrAtIndex(insertDMgr, 4). "+
      "insertDMgr='%v' Error='%v' ", insertDMgr, err.Error())
  }

  if dMgrs.GetNumOfDirs() != 7 {
    t.Errorf("Expected DirMgrsCollection Array Length = '7'. Instead, Array Length = '%v'",
      dMgrs.GetNumOfDirs())
  }

  dMgrAtIdx, err := dMgrs.PeekDirMgrAtIndex(4)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekDirMgrAtIndex(3). "+
      "Error='%v' ", err.Error())

  }

  if !insertDMgr.Equal(&dMgrAtIdx) {
    t.Error("Error: Expected Directory Manager at index 3 should be equal to " +
      "inserted Directory Manager. THEY ARE NOT EQUAL!")
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_06(t *testing.T) {
  dMgrCol := DirMgrCollection{}

  dMgr, err := DirMgr{}.New("../logTest/CmdrX")

  if err != nil {
    t.Errorf("Test Setup Error returned by " +
      "DirMgr{}.New(\"../logTest/CmdrX\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  dMgrCol.dirMgrs = nil

  err = dMgrCol.InsertDirMgrAtIndex(dMgr, 0)

  if err != nil {
    t.Errorf("Error returned by dMgrCol.InsertDirMgrAtIndex(dMgr, 0)\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgrCollection_InsertDirMgrAtIndex_07(t *testing.T) {
  dMgrCol := DirMgrCollection{}

  dMgr, err := DirMgr{}.New("../logTest/CmdrX")

  if err != nil {
    t.Errorf("Test Setup Error returned by " +
      "DirMgr{}.New(\"../logTest/CmdrX\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.isInitialized = false

  err = dMgrCol.InsertDirMgrAtIndex(dMgr, 0)

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrCol.InsertDirMgrAtIndex(dMgr, 0)\n" +
      "because 'dMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}
