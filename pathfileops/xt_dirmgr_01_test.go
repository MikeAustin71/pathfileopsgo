package pathfileops

import (
  "fmt"
  "os"
  "testing"
  "time"
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
  origDir, err := DirMgr01TestCreateCheckFiles03DirFiles()

  if err != nil {
    t.Errorf("Error returned by DirMgr01TestCreateCheckFiles03DirFiles(). Error='%v'", err.Error())
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

func TestDirMgr_DoesDirMgrAbsolutePathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.absolutePath = " "

  result := dMgr.DoesDirMgrAbsolutePathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesDirMgrAbsolutePathExist()\n" +
      "because dMgr.absolutePath consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesDirMgrAbsolutePathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesDirMgrAbsolutePathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesDirMgrAbsolutePathExist()\n" +
      "because dMgr.absolutePath actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}

func TestDirMgr_DoesDirMgrPathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.path = " "

  result := dMgr.DoesDirMgrPathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesDirMgrPathExist()\n" +
      "because dMgr.path consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesDirMgrPathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesDirMgrPathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesDirMgrPathExist()\n" +
      "because dMgr.path actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}

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

  if dMgr2.relativePath != dMgr.relativePath {
    t.Errorf("After CopyOut(), expected dMgr2.relativePath='%v'.  Instead, dMgr2.relativePath='%v'.", dMgr.relativePath, dMgr2.relativePath)
  }

  if dMgr2.isRelativePathPopulated != dMgr.isRelativePathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isRelativePathPopulated='%v'.  Instead, dMgr2.isRelativePathPopulated='%v'.", dMgr.isRelativePathPopulated, dMgr2.isRelativePathPopulated)
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

  if targetDir.DoesDirMgrAbsolutePathExist() {

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

  errStrs := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errStrs) > 0 {
    for i := 0; i < len(errStrs); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryFileOps-Error: %v", errStrs[i])
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
    t.Errorf("Test Clean-Up error returend by err = targetDir.DeleteAll().\n"+
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

  if targetDir.DoesDirMgrAbsolutePathExist() {

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

  errStrs := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errStrs) == 0 {
    t.Error("Expected an error from sourceDir.ExecuteDirectoryFileOps(fileSelect, " +
      "fileOps, targetDir)\nbecause 'sourceDir' is INVALID.\n" +
      "Instead, NO ERROR WAS RETURNED!!!\n")

  }

  err = targetDir.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up error returend by err = targetDir.DeleteAll().\n"+
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

  if targetDir.DoesDirMgrAbsolutePathExist() {

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

  errStrs := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errStrs) > 0 {
    for i := 0; i < len(errStrs); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryTreeOps-Error: %v", errStrs[i])
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

  _ = targetDir.DeleteAll()

}

func TestDirMgr_GetAbsolutePathElements(t *testing.T) {

  testDir := "D:\\Adir\\Bdir\\Cdir\\Ddir\\Edir"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir). Error='%v' ",
      err.Error())
  }

  elementsArray := dMgr.GetAbsolutePathElements()

  if len(elementsArray) != 6 {
    t.Errorf("Error: Expected length of Elements Array='6'. Instead, "+
      "Elements Array length='%v'", len(elementsArray))
  }

  if "D:" != elementsArray[0] {
    t.Errorf("Error. Expected elementsArray[0]=\"D:\". Instead, "+
      "elementsArray[0]=\"%v\"", elementsArray[0])
  }

  if "Adir" != elementsArray[1] {
    t.Errorf("Error. Expected elementsArray[1]=\"Adir\". Instead, "+
      "elementsArray[1]=\"%v\"", elementsArray[1])
  }

  if "Bdir" != elementsArray[2] {
    t.Errorf("Error. Expected elementsArray[2]=\"Bdir\". Instead, "+
      "elementsArray[2]=\"%v\"", elementsArray[2])
  }

  if "Cdir" != elementsArray[3] {
    t.Errorf("Error. Expected elementsArray[3]=\"Cdir\". Instead, "+
      "elementsArray[3]=\"%v\"", elementsArray[3])
  }

  if "Ddir" != elementsArray[4] {
    t.Errorf("Error. Expected elementsArray[4]=\"Ddir\". Instead, "+
      "elementsArray[4]=\"%v\"", elementsArray[4])
  }

  if "Edir" != elementsArray[5] {
    t.Errorf("Error. Expected elementsArray[4]=\"Edir\". Instead, "+
      "elementsArray[4]=\"%v\"", elementsArray[4])
  }

}

func DirMgr01TestCreateCheckFiles03DirFiles() (string, error) {
  ePrefix := "TestFile: xt_dirmgr_01_test.go Func: DirMgr01TestCreateCheckFiles03DirFiles() "
  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")

  if fh.DoesFileExist(origDir) {

    err := os.RemoveAll(origDir)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
    }

  }

  if fh.DoesFileExist(origDir) {
    return "", fmt.Errorf(ePrefix+"Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
  }

  // origDir does NOT exist!
  var ModePerm os.FileMode = 0777

  err := os.MkdirAll(origDir, ModePerm)

  if err != nil {
    return "", fmt.Errorf(ePrefix+"Error returned from os.MkdirAll(origDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origDir, ModePerm, err.Error())
  }

  if !fh.DoesFileExist(origDir) {
    return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origDir='%v'", origDir)
  }

  fileDir := origDir + string(os.PathSeparator)
  newFile1 := fileDir + "checkFile30001.txt"
  fp1, err := os.Create(newFile1)

  if err != nil {
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile1). newFile1='%v' Error='%v' ", newFile1, err.Error())
  }

  newFile2 := fileDir + "checkFile30002.txt"

  fp2, err := os.Create(newFile2)

  if err != nil {
    _ = fp1.Close()
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile2). newFile2='%v' Error='%v' ", newFile2, err.Error())
  }

  newFile3 := fileDir + "checkFile30003.txt"

  fp3, err := os.Create(newFile3)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile3). newFile3='%v' Error='%v' ", newFile3, err.Error())
  }

  newFile4 := fileDir + "checkFile30004.txt"

  fp4, err := os.Create(newFile4)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    _ = fp3.Close()

    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile4). newFile4='%v' Error='%v' ", newFile4, err.Error())
  }

  t := time.Now()
  fmtT := t.Format("2006-01-02 Mon 15:04:05.000000000 -0700 MST")
  _, err = fp4.WriteString(fmtT)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    _ = fp3.Close()
    return "", fmt.Errorf(ePrefix+"%v", err.Error())
  }

  _ = fp1.Close()
  _ = fp2.Close()
  _ = fp3.Close()
  _ = fp4.Close()

  return origDir, nil
}
