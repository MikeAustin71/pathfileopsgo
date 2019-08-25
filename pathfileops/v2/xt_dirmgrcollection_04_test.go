package pathfileops

import (
  "io"
  "strings"
  "testing"
)

func TestDirMgrCollection_PeekDirMgrAtIndex_01(t *testing.T) {
  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../../logTest")

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
  origPath = fh.AdjustPathSlash("../../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  idx2Path := origPath
  idx2AbsPath := origAbsPath

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  idx3Path := origPath
  idx3AbsPath := origAbsPath

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgr, err := dMgrs.PeekDirMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != idx2Path {
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", idx2Path, dMgr.path)
  }

  if dMgr.absolutePath != idx2AbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", idx2AbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

  dMgr, err = dMgrs.PeekDirMgrAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != idx3Path {
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", idx3Path, dMgr.path)
  }

  if dMgr.absolutePath != idx3AbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", idx3AbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

}

func TestDirMgrCollection_PeekDirMgrAtIndex_02(t *testing.T) {

  dMgrs := DirMgrCollection{}

  dMgrs.dirMgrs = nil

  _, err := dMgrs.PeekDirMgrAtIndex(2)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dMgrs.PeekDirMgrAtIndex(2)\n" +
      "because 'dMgrs' is empty.\n" +
      "However, NO ERROR WAS RETURNED!!\n")
  } else {

    if err != io.EOF {
      t.Errorf("ERROR: Expected returned error to equal 'io.EOF' because 'dMgrs' is empty.\n"+
        "The returned error DID NOT EQUAL 'io.EOF'!\n"+
        "Error='%v'\n", err.Error())
    }
  }
}

func TestDirMgrCollection_PeekDirMgrAtIndex_03(t *testing.T) {

  dirArray := []string{
    "../../logTest",
    "../../logTest/CmdrX",
    "../../logTest/FileMgmnt",
    "../../logTest/FileSrc",
    "../../logTest/Level01",
    "../../logTest/Level01/Level02"}

  dMgrs := DirMgrCollection{}

  for i:=0; i < len(dirArray); i++ {

    err := dMgrs.AddDirMgrByPathNameStr(dirArray[i])

    if err != nil {
      t.Errorf("Error returned by "+
        "dMgrs.AddDirMgrByPathNameStr(dirArray[%v])\n"+
        "dirArray[%v]='%v'\n" +
        "Error='%v'\n",
        i,
        i,
        dirArray[i],
        err.Error())
      return
    }
  }

  _, err := dMgrs.PeekDirMgrAtIndex(-1)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dMgrs.PeekDirMgrAtIndex(-1)\n" +
      "because the index, '-1', is less than zero.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgrCollection_PeekDirMgrAtIndex_04(t *testing.T) {

  dirArray := []string{
    "../../logTest",
    "../../logTest/CmdrX",
    "../../logTest/FileMgmnt",
    "../../logTest/FileSrc",
    "../../logTest/Level01",
    "../../logTest/Level01/Level02"}

  dMgrs := DirMgrCollection{}

  for i:=0; i < len(dirArray); i++ {

    err := dMgrs.AddDirMgrByPathNameStr(dirArray[i])

    if err != nil {
      t.Errorf("Error returned by "+
        "dMgrs.AddDirMgrByPathNameStr(dirArray[%v])\n"+
        "dirArray[%v]='%v'\n" +
        "Error='%v'\n",
        i,
        i,
        dirArray[i],
        err.Error())
      return
    }
  }

  _, err := dMgrs.PeekDirMgrAtIndex(99)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dMgrs.PeekDirMgrAtIndex(99)\n" +
      "because the index, '99', exceeded the collection's upper array boundary.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgrCollection_PeekFirstDirMgr_01(t *testing.T) {
  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  dirArray := []string{
    "../../logTest",
    "../../logTest/CmdrX",
    "../../logTest/FileMgmnt",
    "../../logTest/FileSrc",
    "../../logTest/Level01",
    "../../logTest/Level01/Level02"}

  absDirArray := make([]string, 0, 30)

  for i := 0; i < len(dirArray); i++ {

    absStr, err := fh.MakeAbsolutePath(dirArray[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(dirArray[%v])\n"+
        "dirArray[%v]='%v'\nError='%v'\n",
        i, i, dirArray[i], err.Error())
      return
    }

    absDirArray = append(absDirArray, strings.ToLower(absStr))

    err = dMgrs.AddDirMgrByPathNameStr(absStr)

    if err != nil {
      t.Errorf("Error returned by "+
        "dMgrs.AddDirMgrByPathNameStr(absStr)\n"+
        "absStr='%v'\nIndex='%v'\nError='%v'\n",
        absStr,
        i,
        err.Error())
      return
    }

  }

  dMgr1, err := dMgrs.PeekFirstDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgr1, err := dMgrs.PeekFirstDirMgr()\n"+
      "Error='%v'", err.Error())
    return
  }

  absPathDMgr1 := strings.ToLower(dMgr1.GetAbsolutePath())

  if absDirArray[0] != absPathDMgr1 {
    t.Errorf("ERROR: Expected First DirMgr Path is \n"+
      "NOT EQUAL to actual First DirMgr Path!\n"+
      "Expected Path='%v'\nActual Path='%v'\n",
      absDirArray[0], absPathDMgr1)
  }

}

func TestDirMgrCollection_PeekFirstDirMgr_02(t *testing.T) {

  dMgrs := DirMgrCollection{}

  _, err := dMgrs.PeekFirstDirMgr()

  if err == nil {
    t.Error("Expected an error return from dMgrs.PeekFirstDirMgr() because\n" +
      "the dMgrs collection is empty. However, NO ERROR WAS RETURNED!")

  } else {

    if err != io.EOF {
      t.Errorf("ERROR: Expected returned error to equal 'io.EOF' because 'dMgrs' is empty.\n"+
        "The returned error DID NOT EQUAL 'io.EOF'!\n"+
        "Error='%v'\n", err.Error())
    }
  }
}

func TestDirMgrCollection_PeekFirstDirMgr_03(t *testing.T) {
  dirArray := []string{
    "../../logTest" }

  expectedPathStr, err := FileHelper{}.MakeAbsolutePath(dirArray[0])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirArray[0])\n" +
      "dirArray[0]='%v'\n" +
      "Error='%v'\n",
      dirArray[0],
      err.Error())
    return
  }

  expectedPathStr = strings.ToLower(expectedPathStr)

  dMgrs := DirMgrCollection{}

  err = dMgrs.AddDirMgrByPathNameStr(dirArray[0])

  if err != nil {
    t.Errorf("Error returned by "+
      "dMgrs.AddDirMgrByPathNameStr(dirArray[0])\n"+
      "dirArray[0]='%v'\n" +
      "Error='%v'\n",
      dirArray[0],
      err.Error())
    return
  }

  dMgr, err := dMgrs.PeekFirstDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekFirstDirMgr()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedPathStr != strings.ToLower(dMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected path string='%v'\n" +
      "Instead, path string='%v'\n",
      expectedPathStr, strings.ToLower(dMgr.GetAbsolutePath()))
  }
}

func TestDirMgrCollection_PeekLastDirMgr_01(t *testing.T) {

  dirArray := []string{
    "../../logTest",
    "../../logTest/CmdrX",
    "../../logTest/FileMgmnt",
    "../../logTest/FileSrc",
    "../../logTest/Level01",
    "../../logTest/Level01/Level02"}

  expectedPathStr, err := FileHelper{}.MakeAbsolutePath(dirArray[5])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirArray[5])\n" +
      "dirArray[5]='%v'\n" +
      "Error='%v'\n",
      dirArray[5],
      err.Error())
    return
  }

  expectedPathStr = strings.ToLower(expectedPathStr)

  dMgrs := DirMgrCollection{}

  for i:=0; i < len(dirArray); i++ {

    err := dMgrs.AddDirMgrByPathNameStr(dirArray[i])

    if err != nil {
      t.Errorf("Error returned by "+
        "dMgrs.AddDirMgrByPathNameStr(dirArray[%v])\n"+
        "dirArray[%v]='%v'\n" +
        "Error='%v'\n",
        i,
        i,
        dirArray[i],
        err.Error())
      return
    }
  }

  dirMgr, err := dMgrs.PeekLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedPathStr != strings.ToLower(dirMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected path string='%v'\n" +
      "Instead, path string='%v'\n",
      expectedPathStr, strings.ToLower(dirMgr.GetAbsolutePath()))
  }

}

func TestDirMgrCollection_PeekLastDirMgr_02(t *testing.T) {

  dirArray := []string{
    "../../logTest" }

  expectedPathStr, err := FileHelper{}.MakeAbsolutePath(dirArray[0])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirArray[0])\n" +
      "dirArray[0]='%v'\n" +
      "Error='%v'\n",
      dirArray[0],
      err.Error())
    return
  }

  expectedPathStr = strings.ToLower(expectedPathStr)

  dMgrs := DirMgrCollection{}

  for i:=0; i < len(dirArray); i++ {

    err := dMgrs.AddDirMgrByPathNameStr(dirArray[i])

    if err != nil {
      t.Errorf("Error returned by "+
        "dMgrs.AddDirMgrByPathNameStr(dirArray[%v])\n"+
        "dirArray[%v]='%v'\n" +
        "Error='%v'\n",
        i,
        i,
        dirArray[i],
        err.Error())
      return
    }
  }

  dirMgr, err := dMgrs.PeekLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedPathStr != strings.ToLower(dirMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected path string='%v'\n" +
      "Instead, path string='%v'\n",
      expectedPathStr, strings.ToLower(dirMgr.GetAbsolutePath()))
  }
}

func TestDirMgrCollection_PeekLastDirMgr_03(t *testing.T) {


  dMgrs := DirMgrCollection{}

  _, err := dMgrs.PeekLastDirMgr()

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dMgrs.PeekLastDirMgr()\n" +
      "because 'dMgrs' is empty.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  } else {
    if err != io.EOF {
      t.Errorf("ERROR: Expected returned error to equal 'io.EOF' because 'dMgrs' is empty.\n" +
        "The returned error DID NOT EQUAL 'io.EOF'!\n" +
        "Error='%v'\n", err.Error())
    }
  }

}

func TestDirMgrCollection_PopFirstDirMgr_01(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  firstDirPath := fh.AdjustPathSlash("../../logTest")

  origPath := firstDirPath

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  firstAbsDirPath := origAbsPath

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
  origPath = fh.AdjustPathSlash("../../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgr, err := dMgrs.PopFirstDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != firstDirPath {
    t.Errorf("Expected Last DirMgr path='%v'.\nInstead, dMgr.path='%v'",
      firstDirPath, dMgr.path)
  }

  if dMgr.absolutePath != firstAbsDirPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'.\n"+
      "Instead, dMgr.absolutePath='%v'", firstAbsDirPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 5 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 5.\n"+
      "Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

}

func TestDirMgrCollection_PopFirstDirMgr_02(t *testing.T) {

  dMgrs := DirMgrCollection{}

  dMgrs.dirMgrs = nil

  _, err := dMgrs.PopFirstDirMgr()

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrs.PopFirstDirMgr()\n" +
      "because 'dMgrs' is empty.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}


func TestDirMgrCollection_PopLastDirMgr_01(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../../logTest")

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
  origPath = fh.AdjustPathSlash("../../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgr, err := dMgrs.PopLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != origPath {
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
  }

  if dMgr.absolutePath != origAbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'.\n"+
      "Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 5 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 5.\n"+
      "Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }
}

func TestDirMgrCollection_PopLastDirMgr_02(t *testing.T) {

  dirStr := []string{
    "../dirmgrtests",
    "../dirmgrtests/dir01",
    "../dirmgrtests/dir01/dir02",
    "../dirmgrtests/dir01/dir02/dir03" }

  expectedPathStr, err := FileHelper{}.MakeAbsolutePath(dirStr[3])

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirStr[3])\n" +
      "Error='%v'\n", err.Error())
  }

  expectedPathStr = strings.ToLower(expectedPathStr)

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

  dMgr, err := dMgrCol.PopLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrCol.PopLastDirMgr().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedPathStr != strings.ToLower(dMgr.GetAbsolutePath()) {

    t.Errorf("ERROR: Expected that last path='%v'.\n" +
      "Instead, last path='%v'\n",
      expectedPathStr, strings.ToLower(dMgr.GetAbsolutePath()))
    return
  }

  if dMgrCol.GetNumOfDirs() != 3 {
    t.Errorf("ERROR: Expected that remaining number of directories='3'.\n" +
      "Instead, the remaining number of directories = '%v'\n", dMgrCol.GetNumOfDirs())
  }

}

func TestDirMgrCollection_PopLastDirMgr_03(t *testing.T) {

  pathStr :=  "../dirmgrtests"

  expectedPathStr, err := FileHelper{}.MakeAbsolutePath(pathStr)

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(dirStr[3])\n" +
      "Error='%v'\n", err.Error())
  }

  expectedPathStr = strings.ToLower(expectedPathStr)

  dMgrCol := DirMgrCollection{}.New()

  err = dMgrCol.AddDirMgrByPathNameStr(pathStr)

  if err != nil {
    t.Errorf("Error returned by dMgrCol.AddDirMgrByPathNameStr(pathStr).\n" +
      "pathStr='%v'\n"+
      "Error='%v'\n", pathStr, err.Error())
    return
  }

  dMgr, err := dMgrCol.PopLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by dMgrCol.PopLastDirMgr().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedPathStr != strings.ToLower(dMgr.GetAbsolutePath()) {

    t.Errorf("ERROR: Expected that last path='%v'.\n" +
      "Instead, last path='%v'\n",
      expectedPathStr, strings.ToLower(dMgr.GetAbsolutePath()))
    return
  }

  if dMgrCol.GetNumOfDirs() != 0 {
    t.Errorf("ERROR: Expected that remaining number of directories='0'.\n" +
      "Instead, the remaining number of directories = '%v'\n", dMgrCol.GetNumOfDirs())
  }

}

func TestDirMgrCollection_PopLastDirMgr_04(t *testing.T) {
  dMgrs := DirMgrCollection{}
  dMgrs.dirMgrs = nil

  _, err := dMgrs.PopLastDirMgr()

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dMgrs.PopLastDirMgr()\n" +
      "because 'dMgrs' is empty.\n" +
      "However, NO ERROR WAS RETURNED!!\n")
  }

}

func TestDirMgrCollection_PopDirMgrAtIndex_01(t *testing.T) {
  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  origPath := fh.AdjustPathSlash("../../logTest")

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
  origPath = fh.AdjustPathSlash("../../logTest/CmdrX")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../logTest/FileMgmnt")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  idx2Path := origPath
  idx2AbsPath := origAbsPath

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #4
  origPath = fh.AdjustPathSlash("../../logTest/FileSrc")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  idx3Path := origPath
  idx3AbsPath := origAbsPath

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // #5
  origPath = fh.AdjustPathSlash("../../logTest/Level01")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 6
  origPath = fh.AdjustPathSlash("../../logTest/Level01/Level02")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgr, err := dMgrs.PopDirMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != idx2Path {
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", idx2Path, dMgr.path)
  }

  if dMgr.absolutePath != idx2AbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'.\n"+
      "Instead, dMgr.absolutePath='%v'", idx2AbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 5 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 5.\n"+
      "Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

  dMgr, err = dMgrs.PopDirMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if dMgr.path != idx3Path {
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", idx3Path, dMgr.path)
  }

  if dMgr.absolutePath != idx3AbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", idx3AbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 4 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 4.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

}

func TestDirMgrCollection_PopDirMgrAtIndex_02(t *testing.T) {

  dirStr := []string{
    "../dirmgrtests",
    "../dirmgrtests/dir01",
    "../dirmgrtests/dir01/dir02",
    "../dirmgrtests/dir01/dir02/dir03" }


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

  _, err := dmgrCol.PopDirMgrAtIndex(-1)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dmgrCol.PopDirMgrAtIndex(-1)\n" +
      "because the index is less than zero.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestDirMgrCollection_PopDirMgrAtIndex_03(t *testing.T) {

  dirStr := []string{
    "../dirmgrtests",
    "../dirmgrtests/dir01",
    "../dirmgrtests/dir01/dir02",
    "../dirmgrtests/dir01/dir02/dir03" }


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

  _, err := dmgrCol.PopDirMgrAtIndex(99)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dmgrCol.PopDirMgrAtIndex(99)\n" +
      "because the index, '99', is less than zero.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestDirMgrCollection_PopDirMgrAtIndex_04(t *testing.T) {

  dMgrCol := DirMgrCollection{}
  dMgrCol.dirMgrs = nil

  _, err := dMgrCol.PopDirMgrAtIndex(5)

  if err == nil {
    t.Error("ERROR: Expected an error return from dMgrCol.PopDirMgrAtIndex(5)\n" +
      "because 'dMgrCol' is empty.\n" +
      "However, NO ERROR WAS RETURNED!\n")
    return
  } else {
    if err != io.EOF {
      t.Errorf("Error returned by dMgrCol.PopDirMgrAtIndex(5).\n" +
        "'dMgrCol' is empty. However, the error returned is NOT 'io.EOF'.\n" +
        "Error='%v'\n", err.Error())
    }
  }

}

func TestDirMgrCollection_SortByAbsPath_01(t *testing.T) {

  df := make([]string, 5, 10)

  df[4] = "../../dirmgrtests"
  df[3] = "../../dirmgrtests/dir01"
  df[2] = "../../dirmgrtests/dir01/dir02"
  df[1] = "../../dirmgrtests/dir01/dir02/dir03"
  df[0] = "../../dirmgrtests/dir01/dir02/dir03/dir04"

  /* sorted
  df[0] = "../../dirmgrtests"
  df[1] = "../../dirmgrtests/dir01"
  df[2] = "../../dirmgrtests/dir01/dir02"
  df[3] = "../../dirmgrtests/dir01/dir02/dir03"
  df[4] = "../../dirmgrtests/dir01/dir02/dir03/dir04"

   */

  dmgrCol := DirMgrCollection{}.New()

  fh := FileHelper{}

  var err error

  for i := 0; i < 5; i++ {

    err = dmgrCol.AddDirMgrByPathNameStr(df[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

    df[i], err = fh.MakeAbsolutePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  dmgrCol.SortByAbsPath(true)

  for k:=0; k < 5; k++ {

    dMgr, err := dmgrCol.PeekDirMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(k)\n" +
        "k='%v'\nError='%v'\n", k, err.Error())
      return
    }

    if dMgr.absolutePath != df[4-k] {
      t.Errorf("Error: Expected '%v'\n" +
        "Instead, received '%v'\n", df[4-k], dMgr.absolutePath)
      return
    }

  }
}

func TestDirMgrCollection_SortByAbsPath_02(t *testing.T) {

  df := make([]string, 10, 10)
  
  df[0] = "../../dirmgrtests"
  df[1] = "../../dirmgrtests/dir01"
  df[2] = "../../dirmgrtests/dir01/dir02"
  df[3] = "../../dirmgrtests/dir01/dir02/dir03"
  df[4] = "../../dirmgrtests/dir01/dir02/dir03/dir04"
  df[5] = "../../Dirmgrtests"
  df[6] = "../../Dirmgrtests/Dir01"
  df[7] = "../../Dirmgrtests/Dir01/Dir02"
  df[8] = "../../Dirmgrtests/Dir01/Dir02/Dir03"
  df[9] = "../../Dirmgrtests/Dir01/Dir02/Dir03/Dir04"

  dfSorted := make([]string, 10, 10)
  dfSorted[0] = "../../Dirmgrtests"
  dfSorted[1] = "../../Dirmgrtests/Dir01"
  dfSorted[2] = "../../Dirmgrtests/Dir01/Dir02"
  dfSorted[3] = "../../Dirmgrtests/Dir01/Dir02/Dir03"
  dfSorted[4] = "../../Dirmgrtests/Dir01/Dir02/Dir03/Dir04"
  dfSorted[5] = "../../dirmgrtests"
  dfSorted[6] = "../../dirmgrtests/dir01"
  dfSorted[7] = "../../dirmgrtests/dir01/dir02"
  dfSorted[8] = "../../dirmgrtests/dir01/dir02/dir03"
  dfSorted[9] = "../../dirmgrtests/dir01/dir02/dir03/dir04"


  dmgrCol := DirMgrCollection{}.New()

  var err error
  fh := FileHelper{}

  for i := 0; i < 10; i++ {

    err = dmgrCol.AddDirMgrByPathNameStr(df[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

    dfSorted[i], err = fh.MakeAbsolutePath(dfSorted[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }
  }

  dmgrCol.SortByAbsPath(false)

  for k:=0; k < 10; k++ {

    dMgr, err := dmgrCol.PeekDirMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(k)\n" +
        "k='%v'\nError='%v'\n", k, err.Error())
      return
    }

    if dMgr.absolutePath != dfSorted[k] {
      t.Errorf("Error: index='%v' Expected '%v'\n" +
        "Instead, received '%v'\n", k, dfSorted[k], dMgr.absolutePath)
    }
  }
}

func TestDirMgrCollection_SortByAbsPath_03(t *testing.T) {

  df := make([]string, 10, 10)

  df[0] = "../../dirmgrtests"
  df[1] = "../../dirmgrtests/dir01"
  df[2] = "../../dirmgrtests/dir01/dir02"
  df[3] = "../../dirmgrtests/dir01/dir02/dir03"
  df[4] = "../../dirmgrtests/dir01/dir02/dir03/dir04"
  df[5] = "../../Dirmgrtests"
  df[6] = "../../Dirmgrtests/Dir01"
  df[7] = "../../Dirmgrtests/Dir01/Dir02"
  df[8] = "../../Dirmgrtests/Dir01/Dir02/Dir03"
  df[9] = "../../Dirmgrtests/Dir01/Dir02/Dir03/Dir04"

  dfSorted := make([]string, 10, 10)
  dfSorted[0] = "../../dirmgrtests"
  dfSorted[1] = "../../Dirmgrtests"
  dfSorted[2] = "../../dirmgrtests/dir01"
  dfSorted[3] = "../../Dirmgrtests/Dir01"
  dfSorted[4] = "../../dirmgrtests/dir01/dir02"
  dfSorted[5] = "../../Dirmgrtests/Dir01/Dir02"
  dfSorted[6] = "../../dirmgrtests/dir01/dir02/dir03"
  dfSorted[7] = "../../Dirmgrtests/Dir01/Dir02/Dir03"
  dfSorted[8] = "../../dirmgrtests/dir01/dir02/dir03/dir04"
  dfSorted[9] = "../../Dirmgrtests/Dir01/Dir02/Dir03/Dir04"


  dmgrCol := DirMgrCollection{}.New()

  var err error
  fh := FileHelper{}

  for i := 0; i < 10; i++ {

    err = dmgrCol.AddDirMgrByPathNameStr(df[i])

    if err != nil {
      t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

    dfSorted[i], err = fh.MakeAbsolutePath(dfSorted[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }
  }

  dmgrCol.SortByAbsPath(true)

  for k:=0; k < 10; k++ {

    dMgr, err := dmgrCol.PeekDirMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(k)\n" +
        "k='%v'\nError='%v'\n", k, err.Error())
      return
    }

    if dMgr.absolutePath != dfSorted[k] {
      t.Errorf("Error: index='%v' Expected '%v'\n" +
        "Instead, received '%v'\n", k, dfSorted[k], dMgr.absolutePath)
    }
  }
}