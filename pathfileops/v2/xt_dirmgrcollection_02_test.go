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

func TestDirMgrCollection_FindDirectories_01(t *testing.T) {
  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../../logTest")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  origDirMgr, err := DirMgr{}.New(origPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origPath).\n"+
      "origPath='%v'\nError='%v'\n", origPath, err.Error())
    return
  }

  if origDirMgr.absolutePath != origAbsPath {
    t.Errorf("Expected origDirMgr.absolutePath='%v'. Instead, origDirMgr.absolutePath='%v'", origAbsPath, origDirMgr.absolutePath)
  }

  fsc := FileSelectionCriteria{}
  dWlkr, err := origDirMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error retured from origDirMgr.FindWalkDirFiles(fsc).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*Level*"}

  dCol, err := dWlkr.Directories.FindDirectories(fsc)

  if err != nil {
    t.Errorf("Error returned by dCol, err := dWlkr.Directories.FindDirectories(fsc)\n"+
      "Error='%v'\n", err.Error())
    return
  }

  isLevel02Found := false
  isLevel01Found := false

  for i := 0; i < dCol.GetNumOfDirs(); i++ {
    if strings.Contains(dCol.dirMgrs[i].directoryName, "Level01") {
      isLevel01Found = true
    }

    if strings.Contains(dCol.dirMgrs[i].directoryName, "Level02") {
      isLevel02Found = true
    }
  }

  if !isLevel01Found {
    t.Error("Expected to find a directory 'Level01'. It was NOT found!")
  }

  if !isLevel02Found {
    t.Error("Expected to find a directory 'Level02'. It was NOT found!")
  }

}

func TestDirMgrCollection_FindDirectories_02(t *testing.T){

  dMgrCol := DirMgrCollection{}
  dMgrCol.dirMgrs = nil

  fsc := FileSelectionCriteria{}

  _, err := dMgrCol.FindDirectories(fsc)

  if err != nil {
    t.Errorf("ERROR: Expected NO error return from dMgrCol.FindDirectories(fsc)\n" +
      "because 'dMgrCol' is empty.\n" +
      "However, an error was returned!\n" +
      "Error='%v'\n", err.Error())
  }

}
