package pathfileops

import (
  "strings"
  "testing"
)

func TestDirMgrCollection_DeleteAtIndex_01(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir2 := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(d2)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d2). "+
      "Error='%v' ", err.Error())
  }

  arrayLen := dmgrCol.GetNumOfDirs()

  for i := 0; i < arrayLen; i++ {

    dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir2 = true
    }

  }

  if foundDir2 != true {
    t.Error("Expected to find dir # 2 on first pass. DID NOT FIND IT!")
    return
  }

  err = dmgrCol.DeleteAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(2) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = dmgrCol.GetNumOfDirs()

  foundDir2 = false

  for j := 0; j < arrayLen; j++ {
    dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir2 = true
    }

  }

  if foundDir2 != false {
    t.Error("Error: Found dir # 2. IT WAS NOT DELETED!")
  }

}

func TestDirMgrCollection_DeleteAtIndex_02(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(d1)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d1). "+
      "d1='%v' Error='%v' ", d1, err.Error())
  }

  arrayLen := dmgrCol.GetNumOfDirs()

  for i := 0; i < arrayLen; i++ {

    dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find dir # 1 on first pass. DID NOT FIND IT!")
    return
  }

  err = dmgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(1) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = dmgrCol.GetNumOfDirs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found dir # 1. IT WAS NOT DELETED!")
  }

}

func TestDirMgrCollection_DeleteAtIndex_03(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(d0)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d0). "+
      "d0='%v' Error='%v' ", d0, err.Error())
  }

  arrayLen := dmgrCol.GetNumOfDirs()

  for i := 0; i < arrayLen; i++ {

    dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find dir # 0 on first pass. DID NOT FIND IT!")
    return
  }

  err = dmgrCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(0) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = dmgrCol.GetNumOfDirs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found dir # 0. IT WAS NOT DELETED!")
  }

}

func TestDirMgrCollection_DeleteAtIndex_04(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(d3)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d3). "+
      "d3='%v' Error='%v' ", d3, err.Error())
  }

  arrayLen := dmgrCol.GetNumOfDirs()

  for i := 0; i < arrayLen; i++ {

    dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find dir # 3 on first pass. DID NOT FIND IT!")
    return
  }

  err = dmgrCol.DeleteAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(3) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = dmgrCol.GetNumOfDirs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == dmgr.GetAbsolutePath() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found dir # 3. IT WAS NOT DELETED!")
  }

}

func TestDirMgrCollection_DeleteAtIndex_05(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  arrayLen := dmgrCol.GetNumOfDirs()

  if arrayLen != 4 {
    t.Errorf("Expected array length='4'. Instead, array length='%v'",
      arrayLen)
    return
  }

  err = dmgrCol.DeleteAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(3). "+
      "Error='%v' ", err.Error())
  }

  err = dmgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.DeleteAtIndex(1). "+
      "Error='%v' ", err.Error())
  }

  err = dmgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error: Iteration #2 returned by dmgrCol.DeleteAtIndex(1). "+
      "Error='%v' ", err.Error())
  }

  err = dmgrCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error: Iteration #2 returned by dmgrCol.DeleteAtIndex(0). "+
      "Error='%v' ", err.Error())
  }

  arrayLen = dmgrCol.GetNumOfDirs()

  if arrayLen != 0 {
    t.Errorf("Error: Expected final array length=0. Instead, array length='%v'",
      arrayLen)
  }

}

func TestDirMgrCollection_DeleteAtIndex_06(t *testing.T) {

  dMgrsCol := DirMgrCollection{}

  dMgrsCol.dirMgrs = nil

  err := dMgrsCol.DeleteAtIndex(-1)

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrsCol.DeleteAtIndex(-1)\n" +
      "because the index was less than zero.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgrCollection_DeleteAtIndex_07(t *testing.T) {

  dMgrsCol := DirMgrCollection{}

  dMgrsCol.dirMgrs = nil

  err := dMgrsCol.DeleteAtIndex(5)

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrsCol.DeleteAtIndex(5)\n" +
      "because the number of array elements in the collection is zero.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgrCollection_DeleteAtIndex_08(t *testing.T) {

  d0 := "..\\dirmgrtests"
  d1 := "..\\dirmgrtests\\dir01"
  d2 := "..\\dirmgrtests\\dir01\\dir02"
  d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

  dmgrCol := DirMgrCollection{}.New()

  err := dmgrCol.AddDirMgrByPathNameStr(d0)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d1)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d2)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = dmgrCol.AddDirMgrByPathNameStr(d3)

  if err != nil {
    t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
      "Error='%v' ", err.Error())
    return
  }

  if dmgrCol.GetNumOfDirs() != 4 {
    t.Errorf("ERROR: Expected Number of Directries='4'.\n" +
      "Instead, Number of Directories='%v'\n", dmgrCol.GetNumOfDirs())
    return
  }

  err = dmgrCol.DeleteAtIndex(99)

  if err == nil {
    t.Error("ERROR: Expected an error return from dmgrCol.DeleteAtIndex(99)\n" +
      "because the index, '99', exceeds the collection's array length.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

}

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

