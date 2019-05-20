package pathfileops

import (
  "fmt"
  "testing"
)

const (
  logDir    = "../logTest"
  commonDir = "../pathfileops"
)

func TestCleanDir(t *testing.T) {
  var expected, cleanDir, targetDir string

  fh := FileHelper{}
  targetDir = "..///pathfileops"

  cleanDir = fh.CleanPathStr(targetDir)
  expected = fh.CleanPathStr(commonDir)
  if cleanDir != expected {
    t.Error(fmt.Sprintf("Expected Clean Version of %v, got: ", commonDir), cleanDir)
  }

}

func TestDirMgr_ChangeWorkingDir_01(t *testing.T) {
  var err error
  var startDir, checkDir, targetDir string
  fh := FileHelper{}

  startDir, err = fh.GetAbsCurrDir()

  if err != nil {
    t.Error("GetAnsCurrDir() Failed:", err)
  }

  targetDir, err = fh.MakeAbsolutePath(logDir)

  if err != nil {
    t.Error("MakeAbsolutePath() Failed:", err)
  }

  err = fh.ChangeWorkingDir(targetDir)

  if err != nil {
    t.Error("ChangeWorkingDir() Failed:", err)
  }

  checkDir, err = fh.GetAbsCurrDir()

  if err != nil {
    t.Error("GetAbsCurrDir() 2 Failed:", err)
  }

  if checkDir != targetDir {
    t.Error("Target Dir != CheckDir")
  }

  err = fh.ChangeWorkingDir(startDir)

  if err != nil {
    t.Error("Change To Start Dir Failed:", err)
  }

  checkDir, err = fh.GetAbsCurrDir()

  if err != nil {
    t.Errorf("GetAbsCurrDir() 3 Failed. Error='%v'", err)
  }

  if checkDir != startDir {
    t.Error("Start Dir != CheckDir")
  }
}

func TestDirMgr_CopyDirectory_01(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n" +
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopyDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n" +
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }
  // 5 txt src Files
  /*
    "../filesfortest/levelfilesfortest/level_0_0_test.txt"
    "../filesfortest/levelfilesfortest/level_0_1_test.txt"
    "../filesfortest/levelfilesfortest/level_0_2_test.txt"
    "../filesfortest/levelfilesfortest/level_0_3_test.txt"
    "../filesfortest/levelfilesfortest/level_0_4_test.txt"
  */

  fileNames := []string{"level_0_0_test.txt",
    "level_0_1_test.txt",
    "level_0_2_test.txt",
    "level_0_3_test.txt",
    "level_0_4_test.txt" }


  fsc = FileSelectionCriteria{}

  fMgrCollection, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Test Setup Error: Expected to find 5-files in 'targetDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }

  for i:=0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n" +
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(targetDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k:=0;k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n" +
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName )
    }

  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectory_02(t *testing.T) {
  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/iDoNotExist"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n" +
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopyDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDMgr.CopyDirectory(targetDMgr, fsc)\n" +
      "because 'srcDMgr' path DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectory_03(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_03"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n" +
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  srcDMgr.isInitialized = false

  errs := srcDMgr.CopyDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDMgr.CopyDirectory(targetDMgr, fsc)\n" +
      "because 'srcDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!")
  }

  srcDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectory_04(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_04"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n" +
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  targetDMgr.isInitialized = false

  errs := srcDMgr.CopyDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDMgr.CopyDirectory(targetDMgr, fsc)\n" +
      "because 'targetDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectory_05(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_05"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n" +
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.htm"}

  errs := srcDMgr.CopyDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n" +
      "targetDir='%v'\n\n", targetDMgr.GetAbsolutePath())

    for i:=0; i < len(errs); i++ {
      t.Errorf("%v", errs[i])
    }

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }

  // 5 txt src Files
  /*
     "../filesfortest/levelfilesfortest/level_0_0_test.txt"
     "../filesfortest/levelfilesfortest/level_0_1_test.txt"
     "../filesfortest/levelfilesfortest/level_0_2_test.txt"
     "../filesfortest/levelfilesfortest/level_0_3_test.txt"
     "../filesfortest/levelfilesfortest/level_0_4_test.txt"
  */

  if targetDMgr.DoesAbsolutePathExist() {
    t.Errorf("Error: Expected that target directory would not exist because\n" +
      "none of the source files matched the search criteria.\n" +
      "However, the target directory DOES EXIST!!!")
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_01(t *testing.T) {

  srcDir := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }

  fh := FileHelper{}

  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)\n" +
      "targetDMgr='%v'\nErrors:\n",  targetDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("'%v'\n\n", errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if ! targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory path DOES NOT EXIST!!\n")

    return
  }

  srcDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by srcDMgr.FindWalkDirFiles(fsc).\n" +
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n" +
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if srcDTreeInfo.Directories.GetNumOfDirs() != targetDTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected %v-directories would be created. Instead, %v-directories were created!",
      srcDTreeInfo.Directories.GetNumOfDirs(), targetDTreeInfo.Directories.GetNumOfDirs())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if srcDTreeInfo.FoundFiles.GetNumOfFileMgrs() != targetDTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("Expected %v-files would be copied. Instead, %v-files were copied!",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), targetDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyIn_01(t *testing.T) {

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

  dMgr2.CopyIn(&dMgr)

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyIn(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyIn(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.originalPath != dMgr.originalPath {
    t.Errorf("After CopyIn(), expected dMgr2.originalPath='%v'.  Instead, dMgr2.originalPath='%v'.", dMgr.originalPath, dMgr2.originalPath)
  }

  if dMgr2.path != dMgr.path {
    t.Errorf("After CopyIn(), expected dMgr2.path='%v'.  Instead, dMgr2.path='%v'.", dMgr.path, dMgr2.path)
  }

  if dMgr2.isPathPopulated != dMgr.isPathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isPathPopulated='%v'.  Instead, dMgr2.isPathPopulated='%v'.", dMgr.isPathPopulated, dMgr2.isPathPopulated)
  }

  if dMgr2.doesPathExist != dMgr.doesPathExist {
    t.Errorf("After CopyIn(), expected dMgr2.doesPathExist='%v'.  Instead, dMgr2.doesPathExist='%v'.", dMgr.doesPathExist, dMgr2.doesPathExist)
  }

  if dMgr2.parentPath != dMgr.parentPath {
    t.Errorf("After CopyIn(), expected dMgr2.parentPath='%v'.  Instead, dMgr2.parentPath='%v'.", dMgr.parentPath, dMgr2.parentPath)
  }

  if dMgr2.isParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isParentPathPopulated='%v'.  Instead, dMgr2.isParentPathPopulated='%v'.", dMgr.isParentPathPopulated, dMgr2.isParentPathPopulated)
  }

  if dMgr2.relativePath != dMgr.relativePath {
    t.Errorf("After CopyIn(), expected dMgr2.relativePath='%v'.  Instead, dMgr2.relativePath='%v'.", dMgr.relativePath, dMgr2.relativePath)
  }

  if dMgr2.isRelativePathPopulated != dMgr.isRelativePathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isRelativePathPopulated='%v'.  Instead, dMgr2.isRelativePathPopulated='%v'.", dMgr.isRelativePathPopulated, dMgr2.isRelativePathPopulated)
  }

  if dMgr2.absolutePath != dMgr.absolutePath {
    t.Errorf("After CopyIn(), expected dMgr2.absolutePath='%v'.  Instead, dMgr2.absolutePath='%v'.", dMgr.absolutePath, dMgr2.absolutePath)
  }

  if dMgr2.isAbsolutePathPopulated != dMgr.isAbsolutePathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isAbsolutePathPopulated='%v'.  Instead, dMgr2.isAbsolutePathPopulated='%v'.", dMgr.isAbsolutePathPopulated, dMgr2.isAbsolutePathPopulated)
  }

  if dMgr2.doesAbsolutePathExist != dMgr.doesAbsolutePathExist {
    t.Errorf("After CopyIn(), expected dMgr2.doesAbsolutePathExist='%v'.  Instead, dMgr2.doesAbsolutePathExist='%v'.", dMgr.doesAbsolutePathExist, dMgr2.doesAbsolutePathExist)
  }

  if dMgr2.isAbsolutePathDifferentFromPath != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("After CopyIn(), expected dMgr2.isAbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.isAbsolutePathDifferentFromPath='%v'.", dMgr.isAbsolutePathDifferentFromPath, dMgr2.isAbsolutePathDifferentFromPath)
  }

  if dMgr2.volumeName != dMgr.volumeName {
    t.Errorf("After CopyIn(), expected dMgr2.volumeName='%v'.  Instead, dMgr2.volumeName='%v'.", dMgr.volumeName, dMgr2.volumeName)
  }

  if dMgr2.isVolumePopulated != dMgr.isVolumePopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isVolumePopulated='%v'.  Instead, dMgr2.isVolumePopulated='%v'.", dMgr.isVolumePopulated, dMgr2.isVolumePopulated)
  }

  if !dMgr2.Equal(&dMgr) {
    t.Error("After CopyIn(), expected dMgr2 to EQUAL dMgr. It did NOT!")
  }

}

func TestDirMgr_CopyOut_01(t *testing.T) {
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

  if dMgr2.relativePath != dMgr.relativePath {
    t.Errorf("After CopyOut(), expected dMgr2.relativePath='%v'.  Instead, dMgr2.relativePath='%v'.", dMgr.relativePath, dMgr2.relativePath)
  }

  if dMgr2.isRelativePathPopulated != dMgr.isRelativePathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isRelativePathPopulated='%v'.  Instead, dMgr2.isRelativePathPopulated='%v'.", dMgr.isRelativePathPopulated, dMgr2.isRelativePathPopulated)
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

}

func TestDirMgr_DeleteAll_01(t *testing.T) {

  fh := FileHelper{}
  // Set up target directories and files for deletion!
  origDir, err := dirMgr01TestCreateCheckFiles03DirFiles()

  if err != nil {
    t.Errorf("Error returned by dirMgr01TestCreateCheckFiles03DirFiles(). Error='%v'", err.Error())
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
  }

  err = dMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteAll(). dMgr.path='%v'  dMgr.absolutePath='%v'  Error='%v'", dMgr.path, dMgr.absolutePath, err.Error())
  }

  if dMgr.doesAbsolutePathExist {
    t.Errorf("Expected absolutePath to be deleted. Instead, it Exists!\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)
  }

  if fh.DoesFileExist(origDir) {
    t.Errorf("Expected origDir to be deleted. Instead, it Exists!\n"+
      "origDir='%v'\n", origDir)
  }

}

func TestDirMgr_DeleteAll_02(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteAll_02"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr.isInitialized = false

  err = dMgr.DeleteAll()

  if err == nil {
    t.Error("Expected dMgr.DeleteAll() to return an error because\n" +
      "'dMgr' is invalid. However, NO ERROR WAS RETURNED!")
  }

  _ = FileHelper{}.DeleteDirPathAll(testDir)

}

func TestDirMgr_DeleteAllFilesInDir_01(t *testing.T) {

  testDir := "../checkfiles/DeleteAllFilesInDir_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  dMgr.isInitialized = false

  errs := dMgr.DeleteAllFilesInDir()

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteAllFilesInDir()\n" +
      "because 'dMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_DeleteAllFilesInDir_02(t *testing.T) {

  testDir := "../checkfiles/iDoNotExist"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errs := dMgr.DeleteAllFilesInDir()

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteAllFilesInDir()\n" +
      "because 'dMgr' path does NOT exist!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_DeleteAllFilesInDir_03(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_03"
  testDir2 := testDir + "/dir2"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir2).\n"+
      "testDir2='%v'\nError='%v'\n", testDir2, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir2
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  dMgrSub, err := DirMgr{}.New(testDir2)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir2).\n"+
      "testDir2='%v'\nError='%v'\n", testDir2, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.*")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 10 {
    t.Errorf("Test Setup Error: Expected to find 10-files in 'testDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errArray := dMgr.DeleteAllFilesInDir()

  if len(errArray) > 0 {

    if len(errArray) == 1 {
      t.Errorf("Error returned by dMgr.DeleteAllFilesInDir().\n"+
        "testDir='%v'\nError='%v'\n", testDir, err.Error())
    } else {
      t.Errorf("Errors returned by dMgr.DeleteAllFilesInDir().\n"+
        "testDir='%v'\n\n", testDir)

      for i := 0; i < len(errArray); i++ {
        if i == len(errArray)-1 {
          t.Errorf("%v\n\n", errArray[i].Error())
        } else {
          t.Errorf("%v\n", errArray[i].Error())
        }
      }
    }

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err = dMgr.FindFilesByNamePattern("*.*")

  if err != nil {
    t.Errorf("Test Setup Error returned by #2 dMgr.FindFilesByNamePattern(\"*.*\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-files in 'testDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

  }

  fMgrCollection, err = dMgrSub.FindFilesByNamePattern("*.*")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgrSub.FindFilesByNamePattern(\"*.*\").\n"+
      "testDir2='%v'\nError='%v'\n", testDir2, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 3 {
    t.Errorf("Error: Expected to find 3-files in 'testDir2'.\n"+
      "Instead, %v-files were found.\ntestDir2='%v'\n",
      fMgrCollection.GetNumOfFileMgrs(), testDir2)

  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_DeleteFilesByNamePattern_01(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFiles := make([]string, 0, 50)
  destFiles = append(destFiles, testDir+"/level_0_0_test.txt")
  destFiles = append(destFiles, testDir+"/level_0_1_test.txt")
  destFiles = append(destFiles, testDir+"/level_0_2_test.txt")
  destFiles = append(destFiles, testDir+"/level_0_3_test.txt")
  destFiles = append(destFiles, testDir+"/level_0_4_test.txt")
  destFiles = append(destFiles, testDir+"/006860_sample.htm")
  destFiles = append(destFiles, testDir+"/006870_ReadingFiles.htm")
  destFiles = append(destFiles, testDir+"/006890_WritingFiles.htm")

  for i := 0; i < len(srcFiles); i++ {

    err = fh.CopyFileByIo(srcFiles[i], destFiles[i])

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFiles[%v])\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, i, srcFiles[i], destFiles[i], err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 3 {
    t.Errorf("Test Setup Error: Expected to find 3-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errArray := dMgr.DeleteFilesByNamePattern("*.htm")

  if len(errArray) > 0 {
    if len(errArray) == 1 {
      t.Errorf("Error returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
        "testDir='%v'\nError='%v'\n", testDir, err.Error())
    } else {
      t.Errorf("Errors returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
        "testDir='%v'\n\n", testDir)

      for i := 0; i < len(errArray); i++ {
        if i == len(errArray)-1 {
          t.Errorf("%v\n\n", errArray[i].Error())
        } else {
          t.Errorf("%v\n", errArray[i].Error())
        }
      }
    }

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err = dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", fMgrCollection.GetNumOfFileMgrs())

  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_DeleteFilesByNamePattern_02(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errs := dMgr.DeleteFilesByNamePattern("")

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteFilesByNamePattern(\"\")\n" +
      "because the file search parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = fh.DeleteDirPathAll(testDir)

}

func TestDirMgr_DeleteFilesByNamePattern_03(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errs := dMgr.DeleteFilesByNamePattern("     ")

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteFilesByNamePattern(\"     \")\n" +
      "because the file search parameter consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = fh.DeleteDirPathAll(testDir)

}

func TestDirMgr_DeleteFilesByNamePattern_04(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  dMgr.isInitialized = false

  errs := dMgr.DeleteFilesByNamePattern("*.*")

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteFilesByNamePattern(\"*.*\")\n" +
      "because the Directory Manager instance (dMgr) is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = fh.DeleteDirPathAll(testDir)

}

func TestDirMgr_DeleteFilesByNamePattern_05(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_05"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  errs := dMgr.DeleteFilesByNamePattern("*.*")

  if len(errs) == 0 {
    t.Error("Expected an error return from dMgr.DeleteFilesByNamePattern(\"*.*\")\n" +
      "because the Directory Manager instance (dMgr) is DOES NOT EXIST.\n" +
      "However, NO ERROR WAS RETURNED!!!!")
  }

  _ = fh.DeleteDirPathAll(testDir)

}

func TestDirMgr_DeleteFilesByNamePattern_06(t *testing.T) {

  baseTestDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_01"

  testDir := baseTestDir + "/xDir01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseTestDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseTestDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")
  // For Sub-Directory files
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")

  destFiles := make([]string, 0, 50)
  destFiles = append(destFiles, baseTestDir+"/level_0_0_test.txt")
  destFiles = append(destFiles, baseTestDir+"/level_0_1_test.txt")
  destFiles = append(destFiles, baseTestDir+"/level_0_2_test.txt")
  destFiles = append(destFiles, baseTestDir+"/level_0_3_test.txt")
  destFiles = append(destFiles, baseTestDir+"/level_0_4_test.txt")
  destFiles = append(destFiles, baseTestDir+"/006860_sample.htm")
  destFiles = append(destFiles, baseTestDir+"/006870_ReadingFiles.htm")
  destFiles = append(destFiles, baseTestDir+"/006890_WritingFiles.htm")
  // Sub Directory files
  destFiles = append(destFiles, testDir+"/level_1_0_test.txt")
  destFiles = append(destFiles, testDir+"/level_1_1_test.txt")
  destFiles = append(destFiles, testDir+"/level_1_2_test.txt")
  destFiles = append(destFiles, testDir+"/level_1_3_test.txt")

  for i := 0; i < len(srcFiles); i++ {

    err = fh.CopyFileByIo(srcFiles[i], destFiles[i])

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFiles[%v])\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, i, srcFiles[i], destFiles[i], err.Error())

      _ = fh.DeleteDirPathAll(baseTestDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(baseTestDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(baseTestDir).\n"+
      "baseTestDir='%v'\nError='%v'\n", baseTestDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  dMgr2Sub, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 3 {
    t.Errorf("Test Setup Error: Expected to find 3-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  errArray := dMgr.DeleteFilesByNamePattern("*.htm")

  if len(errArray) > 0 {
    if len(errArray) == 1 {
      t.Errorf("Error returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
        "testDir='%v'\nError='%v'\n", testDir, err.Error())
    } else {
      t.Errorf("Errors returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
        "testDir='%v'\n\n", testDir)

      for i := 0; i < len(errArray); i++ {
        if i == len(errArray)-1 {
          t.Errorf("%v\n\n", errArray[i].Error())
        } else {
          t.Errorf("%v\n", errArray[i].Error())
        }
      }

    }

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  fMgrCollection, err = dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "baseTestDir='%v'\nError='%v'\n", baseTestDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", fMgrCollection.GetNumOfFileMgrs())

  }

  fMgrCollection, err = dMgr2Sub.FindFilesByNamePattern("*.txt")

  if err != nil {
    t.Errorf("Error returned by dMgr2Sub.FindFilesByNamePattern(\"*.txt\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 4 {
    t.Errorf("Error expected 4-txt files in the sub-directory. However, the number of\n"+
      "found txt file='%v'", fMgrCollection.GetNumOfFileMgrs())
  }

  err = fh.DeleteDirPathAll(baseTestDir)

  if err != nil {
    t.Errorf("Test File Clean-Up error returned by fh.DeleteDirPathAll(baseTestDir).\n"+
      "baseTestDir='%v'\nError='%v'\n", baseTestDir, err.Error())
  }

  return
}

func TestDirMgr_DoesAbsolutePathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.absolutePath = " "

  result := dMgr.DoesAbsolutePathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesAbsolutePathExist()\n" +
      "because dMgr.absolutePath consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesAbsolutePathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesAbsolutePathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesAbsolutePathExist()\n" +
      "because dMgr.absolutePath actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}

func TestDirMgr_DoesPathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.path = " "

  result := dMgr.DoesPathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesPathExist()\n" +
      "because dMgr.path consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesPathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesPathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesPathExist()\n" +
      "because dMgr.path actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}
