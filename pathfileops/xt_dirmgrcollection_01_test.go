package pathfileops

import (
  "strings"
  "testing"
)

func TestDirMgrCollection_AddDirMgr_01(t *testing.T) {

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
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
  }

  if dMgr.absolutePath != origAbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

}

func TestDirMgrCollection_AddDirMgrByKnownPathDirName_01(t *testing.T) {

  dMgrCol := DirMgrCollection{}
  dMgrCol.dirMgrs = nil
  parentPath := "../checkfiles/checkfiles03"

  dirPath := "checkfiles03_02"

  err := dMgrCol.AddDirMgrByKnownPathDirName(parentPath, dirPath)

  if err != nil {
    t.Errorf("Error returned by dMgrCol.AddDirMgrByKnownPathDirName(parentPath, dirPath)\n" +
      "parentPath='%v'\n, dirPath='%v'\nError='%v'\n",
      parentPath, dirPath, err.Error())
    return
  }

  numOfDMgrs := dMgrCol.GetNumOfDirs()

  if numOfDMgrs != 1 {
    t.Errorf("ERROR: Expected dMgrCol.GetNumOfDirs()='1'.\n" +
      "Instead, dMgrCol.GetNumOfDirs()='%v'", dMgrCol.GetNumOfDirs())
  }

}

func TestDirMgrCollection_AddDirMgrByPathNameStr_01(t *testing.T) {

  testDirStr := ""

  dMgrs := DirMgrCollection{}

  err := dMgrs.AddDirMgrByPathNameStr(testDirStr)

  if err == nil {
    t.Error("Expected an error return from dMgrs.AddDirMgrByPathNameStr(testDirStr)\n" +
      "because input parameter 'testDirStr' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgrCollection_AddDirMgrCollection_01(t *testing.T) {

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
    t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
  }

  if dMgr.absolutePath != origAbsPath {
    t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
  }

  if dMgrs.GetNumOfDirs() != 6 {
    t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

  dMgrs2 := DirMgrCollection{}

  // # Phase 2-2
  origPath = fh.AdjustPathSlash("../filesfortest")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by Phase2-1 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs2.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # Phase 2-2
  origPath = fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by Phase2-2 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs2.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # Phase 2-3
  origPath = fh.AdjustPathSlash("../filesfortest/oldfilesfortest")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by Phase2-3 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  err = dMgrs2.AddDirMgrByPathNameStr(origPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  dMgrs.AddDirMgrCollection(&dMgrs2)

  if dMgrs.GetNumOfDirs() != 9 {
    t.Errorf("Expected after addition - final dMgrs.GetNumOfDirs() == 9.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
  }

  dMgr2, err := dMgrs.PeekLastDirMgr()

  if err != nil {
    t.Errorf("Error returned by Phase 2 dMgrs.PeekLastDirMgr().  Error='%v'", err.Error())
    return
  }

  if dMgr2.path != origPath {
    t.Errorf("Expected Last DirMgr 2 path='%v'. Instead, dMgr2.path='%v'", origPath, dMgr2.path)
  }

  if dMgr2.absolutePath != origAbsPath {
    t.Errorf("Expected Last DirMgr 2 absolutePath='%v'.\n"+
      "Instead, dMgr2.absolutePath='%v'", origAbsPath, dMgr2.absolutePath)
  }

}
func TestDirMgrCollection_AddDirMgrCollection_02(t *testing.T) {

  dMgrs1 := DirMgrCollection{}
  dMgrs1.dirMgrs = nil

  dMgrs2 := DirMgrCollection{}
  dMgrs2.dirMgrs = nil

  dMgrs1.AddDirMgrCollection(&dMgrs2)

  if dMgrs1.dirMgrs == nil {
    t.Error("ERROR: dMgrs1.dirMgrs == nil !!!\n")
    return
  }

  if dMgrs2.dirMgrs == nil {
    t.Error("ERROR: dMgrs2.dirMgrs == nil !!!\n")
    return
  }

  if len(dMgrs1.dirMgrs) != 0 {
    t.Errorf("ERROR: Expected len(dMgrs1.dirMgrs) == '0'.\n" +
      "Instead, len(dMgrs1.dirMgrs) == '%v'", len(dMgrs1.dirMgrs))
  }

}

func TestDirMgrCollection_AddFileInfo_01(t *testing.T) {

  testDir := "../logTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fInfo, err := testDMgr.GetFileInfoPlus()

  dMgrs := DirMgrCollection{}

  err = dMgrs.AddFileInfo(testDMgr.absolutePath, fInfo)

  if err != nil {
    t.Errorf("Error returned by dMgrs.AddFileInfo(testDMgr.absolutePath, fInfo)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.absolutePath, err.Error())
  }

  numOfDirs := dMgrs.GetNumOfDirs()

  if numOfDirs != 1 {
    t.Errorf("ERROR: Expected number of directories in collection would equal '1'.\n"+
      "Instead, the number of directories='%v'.\n", numOfDirs)
  }

}

func TestDirMgrCollection_CopyOut_01(t *testing.T) {

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

  dMgrs2, err := dMgrs.CopyOut()

  if err != nil {
    t.Errorf("Error returned by dMgrs.CopyOut()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgrCnt := dMgrs.GetNumOfDirs()

  if len(dirArray) != dMgrCnt {
    t.Errorf("ERROR: Expected Num of dMgrs='%v'\n"+
      "Instead, Num of dMgrs='%v'\nLength of dirArray='%v'\n",
      len(dirArray),
      dMgrCnt,
      len(dirArray))
    return
  }

  dMgr2Cnt := dMgrs2.GetNumOfDirs()

  if dMgrCnt != dMgr2Cnt {
    t.Errorf("ERROR: Expected Num of dMgrs='%v'\n"+
      "Instead, Num of dMgrs='%v'",
      dMgrCnt, dMgr2Cnt)
  }

  for k := 0; k < dMgrCnt; k++ {

    origDir, err := dMgrs.GetDirMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by origDir, err := dMgrs.GetDirMgrAtIndex(%v)\n"+
        "Error='%v'\n",
        k, err.Error())
      return
    }

    origDirAbsPath := strings.ToLower(origDir.GetAbsolutePath())

    copiedDir, err := dMgrs2.GetDirMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by copiedDir, err := dMgrs2.GetDirMgrAtIndex(%v)\n"+
        "Error='%v'\n",
        k, err.Error())
      return
    }

    copiedDirAbsPath := strings.ToLower(copiedDir.GetAbsolutePath())

    if origDirAbsPath != copiedDirAbsPath {
      t.Errorf("ERROR Index %v: Original Path and Copied Path\n"+
        "are NOT EQUAL!\n"+
        "origDirAbsPath='%v'\ncopiedDirAbsPath='%v'\n\n",
        k,
        origDirAbsPath,
        copiedDirAbsPath)
    }

    if absDirArray[k] != copiedDirAbsPath {
      t.Errorf("ERROR: absDirArray[%v] and Copied Path\n"+
        "are NOT EQUAL!\n"+
        "absDirArray[%v]='%v'\ncopiedDirAbsPath='%v'\n\n",
        k,
        k,
        absDirArray[k],
        copiedDirAbsPath)
    }
  }
}

func TestDirMgrCollection_CopyOut_02(t *testing.T) {

  dMgrs := DirMgrCollection{}

  _, err := dMgrs.CopyOut()

  if err == nil {
    t.Errorf("Expected Error return from 'dMgrs' because the collection\n" +
      "has zero members in the collection. However, NO ERROR WAS RETURNED!!!")
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
