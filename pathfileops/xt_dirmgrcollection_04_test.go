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

  idx2Path := origPath
  idx2AbsPath := origAbsPath

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

  idx3Path := origPath
  idx3AbsPath := origAbsPath

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

func TestDirMgrCollection_PeekFirstDirMgr_01(t *testing.T) {
  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  dirArray := []string{
    "../logTest",
    "../logTest/CmdrX",
    "../logTest/FileMgmnt",
    "../logTest/FileSrc",
    "../logTest/Level01",
    "../logTest/Level01/Level02"}

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
  }
}

func TestDirMgrCollection_PopLastDirMgr_01(t *testing.T) {

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

func TestDirMgrCollection_PopFirstDirMgr_01(t *testing.T) {

  fh := FileHelper{}
  dMgrs := DirMgrCollection{}

  // # 1
  firstDirPath := fh.AdjustPathSlash("../logTest")

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

func TestDirMgrCollection_PopDirMgrAtIndex_01(t *testing.T) {
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

  idx2Path := origPath
  idx2AbsPath := origAbsPath

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

  idx3Path := origPath
  idx3AbsPath := origAbsPath

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

  _, err := dmgrCol.PopDirMgrAtIndex(-1)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from dmgrCol.PopDirMgrAtIndex(-1)\n" +
      "because the index is less than zero.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestDirMgrCollection_PopDirMgrAtIndex_03(t *testing.T) {

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

