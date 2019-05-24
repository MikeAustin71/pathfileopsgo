package pathfileops

import (
  "testing"
)

func TestDirMgr_Equal_01(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if dMgr2.path != origDir2 {
    t.Errorf("Expected original dMgr2.path='%v'. Instead, dMgr2.path='%v'", origDir2, dMgr2.path)
  }

  dMgr2 = dMgr.CopyOut()

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.originalPath != dMgr.originalPath {
    t.Errorf("After CopyOut(), expected dMgr2.originalPath='%v'.  Instead, dMgr2.originalPath='%v'.", dMgr.originalPath, dMgr2.originalPath)
  }

  if dMgr2.path != dMgr.path {
    t.Errorf("After CopyOut(), expected dMgr2.path='%v'.  Instead, dMgr2.path='%v'.", dMgr.path, dMgr2.path)
  }

  if dMgr2.isPathPopulated != dMgr.isPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isPathPopulated='%v'.  Instead, dMgr2.isPathPopulated='%v'.", dMgr.isPathPopulated, dMgr2.isPathPopulated)
  }

  if dMgr2.doesPathExist != dMgr.doesPathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesPathExist='%v'.  Instead, dMgr2.doesPathExist='%v'.", dMgr.doesPathExist, dMgr2.doesPathExist)
  }

  if dMgr2.parentPath != dMgr.parentPath {
    t.Errorf("After CopyOut(), expected dMgr2.parentPath='%v'.  Instead, dMgr2.parentPath='%v'.", dMgr.parentPath, dMgr2.parentPath)
  }

  if dMgr2.isParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isParentPathPopulated='%v'.  Instead, dMgr2.isParentPathPopulated='%v'.", dMgr.isParentPathPopulated, dMgr2.isParentPathPopulated)
  }

  if dMgr2.absolutePath != dMgr.absolutePath {
    t.Errorf("After CopyOut(), expected dMgr2.absolutePath='%v'.  Instead, dMgr2.absolutePath='%v'.", dMgr.absolutePath, dMgr2.absolutePath)
  }

  if dMgr2.isAbsolutePathPopulated != dMgr.isAbsolutePathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathPopulated='%v'.  Instead, dMgr2.isAbsolutePathPopulated='%v'.", dMgr.isAbsolutePathPopulated, dMgr2.isAbsolutePathPopulated)
  }

  if dMgr2.doesAbsolutePathExist != dMgr.doesAbsolutePathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesAbsolutePathExist='%v'.  Instead, dMgr2.doesAbsolutePathExist='%v'.", dMgr.doesAbsolutePathExist, dMgr2.doesAbsolutePathExist)
  }

  if dMgr2.isAbsolutePathDifferentFromPath != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.isAbsolutePathDifferentFromPath='%v'.", dMgr.isAbsolutePathDifferentFromPath, dMgr2.isAbsolutePathDifferentFromPath)
  }

  if dMgr2.volumeName != dMgr.volumeName {
    t.Errorf("After CopyOut(), expected dMgr2.volumeName='%v'.  Instead, dMgr2.volumeName='%v'.", dMgr.volumeName, dMgr2.volumeName)
  }

  if dMgr2.isVolumePopulated != dMgr.isVolumePopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isVolumePopulated='%v'.  Instead, dMgr2.isVolumePopulated='%v'.", dMgr.isVolumePopulated, dMgr2.isVolumePopulated)
  }

  if !dMgr2.Equal(&dMgr) {
    t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!")
  }

  if !dMgr.Equal(&dMgr2) {
    t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!")
  }

}

func TestDirMgr_Equal_02(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if dMgr2.path != origDir2 {
    t.Errorf("Expected original dMgr2.path='%v'. Instead, dMgr2.path='%v'", origDir2, dMgr2.path)
  }

  dMgr2 = dMgr.CopyOut()

  // dMgr2 and dMgr are no longer EQUAL
  dMgr2.absolutePath = dMgr2.absolutePath + "x"

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.originalPath != dMgr.originalPath {
    t.Errorf("After CopyOut(), expected dMgr2.originalPath='%v'.  Instead, dMgr2.originalPath='%v'.", dMgr.originalPath, dMgr2.originalPath)
  }

  if dMgr2.path != dMgr.path {
    t.Errorf("After CopyOut(), expected dMgr2.path='%v'.  Instead, dMgr2.path='%v'.", dMgr.path, dMgr2.path)
  }

  if dMgr2.isPathPopulated != dMgr.isPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isPathPopulated='%v'.  Instead, dMgr2.isPathPopulated='%v'.", dMgr.isPathPopulated, dMgr2.isPathPopulated)
  }

  if dMgr2.doesPathExist != dMgr.doesPathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesPathExist='%v'.  Instead, dMgr2.doesPathExist='%v'.", dMgr.doesPathExist, dMgr2.doesPathExist)
  }

  if dMgr2.parentPath != dMgr.parentPath {
    t.Errorf("After CopyOut(), expected dMgr2.parentPath='%v'.  Instead, dMgr2.parentPath='%v'.", dMgr.parentPath, dMgr2.parentPath)
  }

  if dMgr2.isParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isParentPathPopulated='%v'.  Instead, dMgr2.isParentPathPopulated='%v'.", dMgr.isParentPathPopulated, dMgr2.isParentPathPopulated)
  }

  if dMgr2.absolutePath == dMgr.absolutePath {
    t.Error("After modification, expected dMgr2.absolutePath to be different from dMgr.absolutePath. ERROR= They ARE EQUAL!")
  }

  if dMgr2.isAbsolutePathPopulated != dMgr.isAbsolutePathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathPopulated='%v'.  Instead, dMgr2.isAbsolutePathPopulated='%v'.", dMgr.isAbsolutePathPopulated, dMgr2.isAbsolutePathPopulated)
  }

  if dMgr2.doesAbsolutePathExist != dMgr.doesAbsolutePathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesAbsolutePathExist='%v'.  Instead, dMgr2.doesAbsolutePathExist='%v'.", dMgr.doesAbsolutePathExist, dMgr2.doesAbsolutePathExist)
  }

  if dMgr2.isAbsolutePathDifferentFromPath != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.isAbsolutePathDifferentFromPath='%v'.", dMgr.isAbsolutePathDifferentFromPath, dMgr2.isAbsolutePathDifferentFromPath)
  }

  if dMgr2.volumeName != dMgr.volumeName {
    t.Errorf("After CopyOut(), expected dMgr2.volumeName='%v'.  Instead, dMgr2.volumeName='%v'.", dMgr.volumeName, dMgr2.volumeName)
  }

  if dMgr2.isVolumePopulated != dMgr.isVolumePopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isVolumePopulated='%v'.  Instead, dMgr2.isVolumePopulated='%v'.", dMgr.isVolumePopulated, dMgr2.isVolumePopulated)
  }

  if dMgr2.Equal(&dMgr) {
    t.Error("After modification, expected dMgr2 to NOT EQUAL to dMgr. Wrong- dMgr2 == dMgr!")
  }

  if dMgr.Equal(&dMgr2) {
    t.Error("After modification, expected dMgr to NOT EQUAL to dMgr2. Wrong- dMgr == dMgr2!")
  }

}

func TestDirMgr_EqualAbsPaths_01(t *testing.T) {
  fh := FileHelper{}

  origDir := "../testfiles/testfiles2"

  origDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir). "+
      "origDir='%v' Error='%v' ", origDir, err.Error())
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := "../testfiles/testfiles2"

  origDir2, err = fh.MakeAbsolutePath(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir2). "+
      "origDir2='%v' Error='%v' ", origDir2, err.Error())
  }

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). "+
      "origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if !dMgr.EqualAbsPaths(&dMgr2) {
    t.Error("Expected two paths to be EQUAL. Error: They were NOT!")
  }

}

func TestDirMgr_EqualAbsPaths_02(t *testing.T) {
  fh := FileHelper{}

  origDir := "../testfiles/testfiles2"

  origDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir). "+
      "origDir='%v' Error='%v' ", origDir, err.Error())
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := "../testfiles/testfilesx"

  origDir2, err = fh.MakeAbsolutePath(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir2). "+
      "origDir2='%v' Error='%v' ", origDir2, err.Error())
  }

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). "+
      "origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if dMgr.EqualAbsPaths(&dMgr2) {
    t.Error("Expected two paths to be NOT EQUAL. Error: They were EQUAL!")
  }

}

func TestDirMgr_EqualAbsPaths_03(t *testing.T) {
  fh := FileHelper{}

  origDir := "../TESTfiles/TESTfiles2"

  origDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir). "+
      "origDir='%v' Error='%v' ", origDir, err.Error())
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := "../testfiles/testfiles2"

  origDir2, err = fh.MakeAbsolutePath(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(origDir2). "+
      "origDir2='%v' Error='%v' ", origDir2, err.Error())
  }

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). "+
      "origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if !dMgr.EqualAbsPaths(&dMgr2) {
    t.Error("Expected two paths to be EQUAL. Error: They were NOT!")
  }

}

func TestDirMgr_EqualAbsPaths_04(t *testing.T) {

  dirPath := "../checkfiles"

  dMgr1, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr1 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  dMgr2, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr2 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  dMgr1.isInitialized = false

  if dMgr1.EqualAbsPaths(&dMgr2) {
    t.Error("ERROR: Expected the return from dMgr1.EqualAbsPaths(&dMgr2) to be 'false'\n" +
      "because dMgr1 is not initialized. However, the return value was 'true' !")
  }

}

func TestDirMgr_EqualAbsPaths_05(t *testing.T) {

  dirPath := "../checkfiles"

  dMgr1, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr1 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  dMgr2, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr2 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  dMgr2.isInitialized = false

  if dMgr1.EqualAbsPaths(&dMgr2) {
    t.Error("ERROR: Expected the return from dMgr1.EqualAbsPaths(&dMgr2) to be 'false'\n" +
      "because dMgr2 is NOT initialized. However, the return value was 'true' !")
  }

}

func TestDirMgr_EqualAbsPaths_06(t *testing.T) {

  dirPath := "../checkfiles"

  dMgr1, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr1 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  dMgr2 := DirMgr{}

  if dMgr1.EqualAbsPaths(&dMgr2) {
    t.Error("ERROR: Expected the return from dMgr1.EqualAbsPaths(&dMgr2) to be 'false'\n" +
      "because dMgr2 has NOT been initialized. However, the return value was 'true' !")
  }

}

func TestDirMgr_EqualAbsPaths_07(t *testing.T) {

  dirPath := "../checkfiles"

  dMgr1 := DirMgr{}

  dMgr2, err := DirMgr{}.New(dirPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr2 = DirMgr{}.New(dirPath)\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  if dMgr1.EqualAbsPaths(&dMgr2) {
    t.Error("ERROR: Expected the return from dMgr1.EqualAbsPaths(&dMgr2) to be 'false'\n" +
      "because dMgr1 is NOT initialized. However, the return value was 'true' !")
  }

}

func TestDirMgr_EqualPaths_01(t *testing.T) {
  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := fh.AdjustPathSlash("../testfiles/testfiles1")

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if !dMgr.EqualPaths(&dMgr2) {
    t.Error("Expected two paths to be EQUAL. Error: They were NOT!")
  }

}

func TestDirMgr_EqualPaths_02(t *testing.T) {
  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := fh.AdjustPathSlash("../testfiles/testfiles2")

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if dMgr.EqualPaths(&dMgr2) {
    t.Error("Expected two paths to be NOT EQUAL. Error: They were EQUAL!")
  }

}

func TestDirMgr_EqualPaths_03(t *testing.T) {
  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../Testfiles/Testfiles2")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  origDir2 := fh.AdjustPathSlash("../testfiles/testfiles2")

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
  }

  if !dMgr.EqualPaths(&dMgr2) {
    t.Error("Expected two paths to be EQUAL. Error: They were NOT!")
  }

}

func TestDirMgr_EqualPaths_04(t *testing.T) {

  dirPath1 := "../checkfiles/checkfiles02"

  dMgr1, err := DirMgr{}.New(dirPath1)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(dirPath1).\n"+
      "dirPath1='%v'\nError='%v'", dirPath1, err.Error())
  }

  dMgr2, err := DirMgr{}.New(dirPath1)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(dirPath1).\n"+
      "dirPath1='%v'\nError='%v'", dirPath1, err.Error())
  }

  dMgr1.isInitialized = false

  result := dMgr1.EqualPaths(&dMgr2)

  if result == true {
    t.Error("Expected dMgr1.EqualPaths(&dMgr2) to return 'false' \n" +
      "because dMgr1.isInitialized = 'false'.\n" +
      "Instead, result='true'.\n")
  }

}

func TestDirMgr_EqualPaths_05(t *testing.T) {

  dirPath1 := "../checkfiles/checkfiles02"
  dirPath2 := "../createFilesTest/Level01"
  dMgr1, err := DirMgr{}.New(dirPath1)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(dirPath1).\n"+
      "dirPath1='%v'\nError='%v'", dirPath1, err.Error())
  }

  dMgr2, err := DirMgr{}.New(dirPath2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(dirPath2).\n"+
      "dirPath2='%v'\nError='%v'", dirPath2, err.Error())
  }

  dMgr1.absolutePath = dMgr2.absolutePath

  result := dMgr1.EqualPaths(&dMgr2)

  if result == true {
    t.Error("Expected dMgr1.EqualPaths(&dMgr2) to return 'false' \n" +
      "because 'dMgr1.Path' is different from 'dMgr2.Path'.\n" +
      "Instead, result='true'.\n")
  }

}

func TestDirMgr_ExecuteDirectoryFileOps_01(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) > 0 {
    for i := 0; i < len(errArray); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryFileOps-Error: %v", errArray[i].Error())
    }
  }

  dTreeInfo, err := targetDir.FindWalkDirFiles(fileSelect)

  if err != nil {
    t.Errorf("Error returned by targetDir.FindWalkDirFiles(fileSelect) "+
      "targetDir='%v' Error='%v' ",
      targetDir.GetAbsolutePath(), err.Error())
  }

  lenErrs := len(dTreeInfo.ErrReturns)

  if lenErrs > 0 {
    for i := 0; i < len(dTreeInfo.ErrReturns); i++ {
      t.Errorf("targetDir.FindWalkDirFiles-Errors: %v", dTreeInfo.ErrReturns[i])
    }
  }

  lenDirs := dTreeInfo.Directories.GetNumOfDirs()

  if lenDirs != 1 {
    t.Errorf("Error: Expected number of directories found='%v'. "+
      "Instead, number of directories found='%v' ", 1, lenDirs)
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles != 5 {
    t.Errorf("Error: Expected number of found files='%v'. "+
      "Instead, number of found files='%v' ", 5, numOfFiles)
  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error returned by err = targetDir.DeleteAll().\n"+
      "targetDir='%v'\nError='%v'\n", targetDir.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryFileOps_02(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  sourceDir.isInitialized = false

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryFileOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' is INVALID.\n" +
      "Instead, NO ERROR WAS RETURNED!!!\n")

  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error returned by err = targetDir.DeleteAll().\n"+
      "targetDir='%v'\nError='%v'\n", targetDir.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryFileOps_03(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  targetDir.isInitialized = false

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryFileOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' is INVALID.\n" +
      "Instead, NO ERROR WAS RETURNED!!!\n")

  }

  targetDir.isInitialized = true

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error returned by err = targetDir.DeleteAll().\n"+
      "targetDir='%v'\nError='%v'\n", targetDir.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryFileOps_04(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  rawSrcPath := "../filesfortest/iDoNotExist/TestDirMgr_ExecuteDirectoryFileOps_04"

  sourceDirStr, err := fh.MakeAbsolutePath(rawSrcPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "rawSrcPath)\n"+
      "rawSrcPath='%v'\nError='%v'\n", rawSrcPath, err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryFileOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' does NOT EXIST.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error returned by err = targetDir.DeleteAll().\n"+
      "targetDir='%v'\nError='%v'\n", targetDir.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryFileOps_05(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 0, 5)

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryFileOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'fileOps' is a zero length array.\n" +
      "Instead, NO ERROR WAS RETURNED!!!\n")

  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error return3d by err = targetDir.DeleteAll().\n"+
      "targetDir='%v'\nError='%v'\n", targetDir.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryTreeOps_01(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) > 0 {
    for i := 0; i < len(errArray); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryTreeOps-Error: %v", errArray[i])
    }
  }

  dTreeInfo, err := targetDir.FindWalkDirFiles(fileSelect)

  if err != nil {
    t.Errorf("Error returned by targetDir.FindWalkDirFiles(fileSelect) "+
      "targetDir='%v' Error='%v' ",
      targetDir.GetAbsolutePath(), err.Error())
  }

  lenErrs := len(dTreeInfo.ErrReturns)

  if lenErrs > 0 {
    for i := 0; i < len(dTreeInfo.ErrReturns); i++ {
      t.Errorf("targetDir.FindWalkDirFiles-Errors: %v", dTreeInfo.ErrReturns[i])
    }
  }

  lenDirs := dTreeInfo.Directories.GetNumOfDirs()

  if lenDirs != 5 {
    t.Errorf("Error: Expected number of directories found='%v'. "+
      "Instead, number of directories found='%v' ", 5, lenDirs)
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles != 25 {
    t.Errorf("Error: Expected number of found files='%v'. "+
      "Instead, number of found files='%v' ", 25, numOfFiles)
  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDir.DeleteAll().\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryTreeOps_02(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  sourceDir.isInitialized = false

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryTreeOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDir.DeleteAll().\n" +
      "Error='%v'\n", err.Error())
  }


}

func TestDirMgr_ExecuteDirectoryTreeOps_03(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  targetDir.isInitialized = false

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryTreeOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'targetDir' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  targetDir.isInitialized = true

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDir.DeleteAll().\n" +
      "Error='%v'\n", err.Error())
  }


}

func TestDirMgr_ExecuteDirectoryTreeOps_04(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
    return
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/iDoNotExist")

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/iDoNotExist\") "+
      "Error='%v' ", err.Error())
    return
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Test Setup Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
      return
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOpCode.CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryTreeOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' DOES NOT EXIST.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDir.DeleteAll().\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgr_ExecuteDirectoryTreeOps_05(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 0, 5)

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryTreeOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'fileOps' is a zero length array.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDir.DeleteAll().\n" +
      "Error='%v'\n", err.Error())
  }


}
