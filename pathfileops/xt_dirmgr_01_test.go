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

func TestDirMgr_CopyDirectoryTree_02(t *testing.T) {
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

  srcDMgr.isInitialized = false

  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_03(t *testing.T) {
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

  targetDMgr.isInitialized = false

  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'targetDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!\n\n")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_04(t *testing.T) {
  srcDir := "../filesfortest/iDoNotExist"

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

  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!!\n")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_05(t *testing.T) {

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
  fsc.FileNamePatterns = []string{"*.htm"}

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

  if targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory DOES EXIST!!\n" +
      "The target directory should NOT have been created because none of the files\n" +
      "is the source directory matched the file selection criteria.\n" +
      "However, the target directory DOES EXIST! ERROR!!!!\n")
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

func TestDirMgr_CopySubDirectoryTree_01(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_01"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }


  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)


  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "targetDMgr='%v'\nErrors:\n",  targetDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("'%v'\n\n", errs[i].Error())
    }

    _ = targetDMgr.DeleteAll()

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

    _ = targetDMgr.DeleteAll()

    return
  }

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n" +
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  if srcDTreeInfo.Directories.GetNumOfDirs() != targetDTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("ERROR: Expected %v-directories would be created. Instead, %v-directories were created!",
      srcDTreeInfo.Directories.GetNumOfDirs(), targetDTreeInfo.Directories.GetNumOfDirs())

    _ = targetDMgr.DeleteAll()

    return
  }

  if srcDTreeInfo.FoundFiles.GetNumOfFileMgrs() != targetDTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("ERROR: Expected %v-files would be copied. Instead, %v-files were copied!",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), targetDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_02(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_02"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }


  fsc := FileSelectionCriteria{}

  srcDMgr.isInitialized = false

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' is INVALID!\n"+
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  srcDMgr.isInitialized = true

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_03(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_03"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }


  fsc := FileSelectionCriteria{}

  targetDMgr.isInitialized = false

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'targetDMgr' is INVALID!\n"+
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  targetDMgr.isInitialized = true

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}


func TestDirMgr_CopySubDirectoryTree_04(t *testing.T) {

  srcDir := "../filesfortest/iDoNotExist"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_04"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }


  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' DOES NOT EXIST !\n"+
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_05(t *testing.T) {

  srcDir := "../logTest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_05"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
  }


  fsc := FileSelectionCriteria{}

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, false, fsc)


  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "targetDMgr='%v'\nErrors:\n",  targetDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("'%v'\n\n", errs[i].Error())
    }

    _ = targetDMgr.DeleteAll()

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

    _ = targetDMgr.DeleteAll()

    return
  }

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n" +
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  srcDirs := srcDTreeInfo.Directories.GetNumOfDirs()
  srcDirs--

  targetDirs := targetDTreeInfo.Directories.GetNumOfDirs()

  if srcDirs != targetDirs   {
    t.Errorf("ERROR: Expected %v-directories would be created. Instead, %v-directories were created!",
      srcDirs, targetDirs)

    _ = targetDMgr.DeleteAll()

    return
  }

  tFileInfo, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if tFileInfo.GetNumOfFileMgrs() > 0 {
    t.Errorf("ERROR: Expected ZERO files in top level target directory.\n" +
      "Instead, the top level target directory had %v-files.\nTarget Directory='%v'\n",
      tFileInfo.GetNumOfFileMgrs(), targetDMgr.GetAbsolutePath())
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}
