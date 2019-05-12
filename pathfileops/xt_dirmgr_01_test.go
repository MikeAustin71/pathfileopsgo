package pathfileops

import (
  "fmt"
  "os"
  "testing"
  "time"
)

const (
  logDir = "../logTest"
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
      "dMgr.absolutePath='%v'\n",dMgr.absolutePath)
  }

  if fh.DoesFileExist(origDir) {
    t.Errorf("Expected origDir to be deleted. Instead, it Exists!\n" +
      "origDir='%v'\n", origDir)
  }

}

func TestDirMgr_DeleteAll_02(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteAll_02"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n" +
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

func TestDirMgr_DeleteFilesByNamePattern_01(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string,0, 50)
  
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")
  
  destFiles := make([]string,0,50)
  destFiles = append(destFiles, testDir + "/level_0_0_test.txt")
  destFiles = append(destFiles, testDir + "/level_0_1_test.txt")
  destFiles = append(destFiles, testDir + "/level_0_2_test.txt")
  destFiles = append(destFiles, testDir + "/level_0_3_test.txt")
  destFiles = append(destFiles, testDir + "/level_0_4_test.txt")
  destFiles = append(destFiles, testDir + "/006860_sample.htm")
  destFiles = append(destFiles, testDir + "/006870_ReadingFiles.htm")
  destFiles = append(destFiles, testDir + "/006890_WritingFiles.htm")
  
  
  for i:=0; i < len(srcFiles); i++ {

    err = fh.CopyFileByIo(srcFiles[i], destFiles[i])

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFiles[%v])\n" +
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i,i,srcFiles[i], destFiles[i], err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error() )

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 3 {
    t.Errorf("Test Setup Error: Expected to find 3-html files in 'testDir'.\n" +
      "Instead, %v-html files were found.",fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = dMgr.DeleteFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err = dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error() )

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-html files in 'testDir'.\n" +
      "Instead, %v-html files were found.",fMgrCollection.GetNumOfFileMgrs())

  }

  _ = fh.DeleteDirPathAll(testDir)

  return
}

func TestDirMgr_DeleteFilesByNamePattern_02(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_DeleteFilesByNamePattern_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = dMgr.DeleteFilesByNamePattern("")

  if err == nil {
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
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = dMgr.DeleteFilesByNamePattern("     ")

  if err == nil {
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
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  dMgr.isInitialized = false

  err = dMgr.DeleteFilesByNamePattern("*.*")

  if err == nil {
    t.Error("Expected an error return from dMgr.DeleteFilesByNamePattern(\"*.*\")\n" +
      "because the Directory Manager instance (dMgr) is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = fh.DeleteDirPathAll(testDir)

}

/*
func TestDirMgr_DeleteWalkDirFiles_31(t *testing.T) {

  origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

  if err != nil {
    t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). Error='%v'", err.Error())
  }

  fh := FileHelper{}

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'", origDir, err.Error())
  }

  searchPattern1 := "*.txt"
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)
  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern1}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ORSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"
  newFile1 := "newerFileForTest_01.txt"
  newFile2 := "newerFileForTest_02.txt"
  newFile3 := "newerFileForTest_03.txt"

  oldFile1Found := false
  oldFile2Found := false
  oldFile3Found := false

  newFile1Found := false
  newFile2Found := false
  newFile3Found := false

  for i := 0; i < dInfo.DeletedFiles.GetNumOfFileMgrs(); i++ {

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile1) {
      oldFile1Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile2) {
      oldFile2Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile3) {
      oldFile3Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile1) {
      newFile1Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile2) {
      newFile2Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile3) {
      newFile3Found = true
    }

  }

  if oldFile1Found == false {
    t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!", oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!", oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!", oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!", newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!", newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!", newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

}


func TestDirMgr_DeleteWalkDirFiles_32(t *testing.T) {

  origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

  if err != nil {
    t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). Error='%v'", err.Error())
  }

  fh := FileHelper{}

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'", origDir, err.Error())
  }

  searchPattern1 := "*.txt"
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)
  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern1}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ORSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"
  newFile1 := "newerFileForTest_01.txt"
  newFile2 := "newerFileForTest_02.txt"
  newFile3 := "newerFileForTest_03.txt"

  oldFile1Found := false
  oldFile2Found := false
  oldFile3Found := false

  newFile1Found := false
  newFile2Found := false
  newFile3Found := false

  for i := 0; i < dInfo.DeletedFiles.GetNumOfFileMgrs(); i++ {

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile1) {
      oldFile1Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile2) {
      oldFile2Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, oldFile3) {
      oldFile3Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile1) {
      newFile1Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile2) {
      newFile2Found = true
    }

    if strings.Contains(dInfo.DeletedFiles.fileMgrs[i].fileNameExt, newFile3) {
      newFile3Found = true
    }

  }

  if oldFile1Found == false {
    t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!", oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!", oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!", oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!", newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!", newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!", newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

}
*/

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

  fileOps[0] = FileOperationCode(0).CopySourceToDestinationByIo()

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

  _ = targetDir.DeleteAll()

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
