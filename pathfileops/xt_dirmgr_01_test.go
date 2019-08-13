package pathfileops

import (
  "fmt"
  "strings"
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

func TestDirMgr_ConsolidateErrors_01(t *testing.T) {

  errs := make([]error, 0, 300)

  maxCnt := 9

  for i := 0; i < maxCnt; i++ {

    err := fmt.Errorf("Heder\nError-%v text\n\n", i+1)

    errs = append(errs, err)

  }

  conSolError := DirMgr{}.ConsolidateErrors(errs)

  if conSolError == nil {
    t.Error("Test Error returned from DirMgr{}.ConsolidateErrors(errs) is 'nil'\n")
    return
  }

  errStr := fmt.Sprintf("%v", conSolError.Error())

  if len(errStr) == 0 {
    t.Error("Error string returned from DirMgr{}.ConsolidateErrors(errs) is zero length")
  }

  testCnt := 0

  for j := 0; j < maxCnt; j++ {

    testStr := fmt.Sprintf("Error-%v text", j+1)

    if strings.Contains(errStr, testStr) {
      testCnt++
    }
  }

  if maxCnt != testCnt {
    t.Errorf("ERROR: Expected Error String to contain %v Errors.\n"+
      "Instead, found only %v Errors.",
      maxCnt, testCnt)
  }

}

func TestDirMgr_ConsolidateErrors_02(t *testing.T) {

  testDir := "../checkfiles"
  testDirMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by testDirMgr, err := DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  errs := make([]error, 0, 300)

  maxCnt := 9

  for i := 0; i < maxCnt; i++ {

    err := fmt.Errorf("Heder\nError-%v text\n\n", i+1)

    errs = append(errs, err)

  }

  conSolError := testDirMgr.ConsolidateErrors(errs)

  if conSolError == nil {
    t.Error("Test Error returned from testDirMgr{}.ConsolidateErrors(errs) is 'nil'\n")
    return
  }

  errStr := fmt.Sprintf("%v", conSolError.Error())

  if len(errStr) == 0 {
    t.Error("Error string returned from DirMgr{}.ConsolidateErrors(errs) is zero length")
  }

  testCnt := 0

  for j := 0; j < maxCnt; j++ {

    testStr := fmt.Sprintf("Error-%v text", j+1)

    if strings.Contains(errStr, testStr) {
      testCnt++
    }
  }

  if maxCnt != testCnt {
    t.Errorf("ERROR: Expected Error String to contain %v Errors.\n"+
      "Instead, found only %v Errors.",
      maxCnt, testCnt)
  }

}

func TestDirMgr_ConsolidateErrors_03(t *testing.T) {

  testDir := "../checkfiles"
  testDirMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by testDirMgr, err := DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  errs := make([]error, 0, 300)

  conSolError := testDirMgr.ConsolidateErrors(errs)

  if conSolError != nil {
    t.Error("ERROR: Expected a 'nil' return from  testDirMgr.ConsolidateErrors(errs)\n" +
      "because errs is 'nil'. However, the returned value was NOT 'nil'.")
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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  dirCopyStats,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

  if len(errs) > 0 {

    t.Errorf("Error returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n"+
      "targetDir='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

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
    "level_0_4_test.txt"}

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

  if 5 != dirCopyStats.FilesCopied {
    t.Errorf("Test Setup Error: Expected that dirCopyStats.FilesCopied='5'.\n"+
      "Instead, dirCopyStats.FilesCopied='%v'.\n",
      dirCopyStats.FilesCopied)

    return
  }

  for i := 0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(targetDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k := 0; k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n"+
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName)
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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/iDoNotExist"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  srcDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  targetDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.htm"}

  dirCopyStats,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

  if len(errs) > 0 {

    t.Errorf("Errors returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n"+
      "targetDir='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

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

  if dirCopyStats.FilesCopied != 0 {
    t.Errorf("Expected that dirCopyStats.FilesCopied='0'.\n"+
      "Instead, dirCopyStats.FilesCopied='%v'", dirCopyStats.FilesCopied)
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_CopyDirectory_06(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_06"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  dirCopyStats,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, false)

  if len(errs) > 0 {

    t.Errorf("Error returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n"+
      "targetDir='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

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
    "level_0_4_test.txt"}

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Error: Expected to find 5-files in 'targetDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }

  if 5 != dirCopyStats.FilesCopied {
    t.Errorf("Error: Expected that dirCopyStats.FilesCopied='5'.\n"+
      "Instead, dirCopyStats.FilesCopied='%v'.\n",
      dirCopyStats.FilesCopied)

    return
  }

  if 0 != dirCopyStats.FilesNotCopied {
    t.Errorf("Error: Expected that dirCopyStats.FilesNotCopied='0'.\n"+
      "Instead, dirCopyStats.FilesNotCopied='%v'.\n",
      dirCopyStats.FilesNotCopied)

    return
  }

  if 5 != dirCopyStats.TotalFilesProcessed {
    t.Errorf("Error: Expected that dirCopyStats.TotalFilesProcessed='5'.\n"+
      "Instead, dirCopyStats.TotalFilesProcessed='%v'.\n",
      dirCopyStats.TotalFilesProcessed)

    return
  }

  if 1 != dirCopyStats.DirCreated {
    t.Errorf("Error: Expected that dirCopyStats.DirsCreated='1'.\n"+
      "Instead, dirCopyStats.TotalFilesProcessed='%v'.\n",
      dirCopyStats.DirCreated)

    return
  }

  if dirCopyStats.ComputeError != nil {
    t.Errorf("Error: Expected that dirCopyStats.ComputeError='nil'.\n"+
      "Instead, dirCopyStats.ComputeError='%v'.\n",
      dirCopyStats.ComputeError.Error())

    return
  }

  for i := 0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(targetDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k := 0; k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n"+
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName)
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

func TestDirMgr_CopyDirectory_07(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_CopyFilesToDirectory_07"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "testDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDir1 := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir1).\n"+
      "srcDir1='%v'\nError='%v'\n", srcDir1, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.xxx"}

  dirCopyStats,
  errs := srcDMgr.CopyDirectory(targetDMgr, fsc, true)

  if len(errs) > 0 {

    t.Errorf("Error returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n"+
      "targetDir='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }

  if !targetDMgr.DoesPathExist() {

    t.Errorf("Error: CopyDirectory() was called with parameter, 'copyEmptyDirectory' = 'true'.\n"+
      "Therefore, the target directory should have been created even though no files were\n"+
      "copied to the target director. However, the target directory was NOT created and DOES NOT EXIST!\n"+
      "targetDir='%v'\n",
      targetDMgr.GetAbsolutePath())
    return

  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-files in 'targetDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(targetDir)

    return

  }

  if 0 != dirCopyStats.FilesCopied {
    t.Errorf("Error: Expected that dirCopyStats.FilesCopied='0'.\n"+
      "Instead, dirCopyStats.FilesCopied='%v'.\n",
      dirCopyStats.FilesCopied)

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if 5 != dirCopyStats.TotalFilesProcessed {
    t.Errorf("Error: Expected that dirCopyStats.TotalFilesProcessed='5'.\n"+
      "Instead, dirCopyStats.TotalFilesProcessed='%v'.\n",
      dirCopyStats.TotalFilesProcessed)

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if 5 != dirCopyStats.FilesNotCopied {
    t.Errorf("Error: Expected that dirCopyStats.FilesNotCopied='5'.\n"+
      "Instead, dirCopyStats.FilesNotCopied='%v'.\n",
      dirCopyStats.FilesNotCopied)

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if 1 != dirCopyStats.DirCreated {
    t.Errorf("Error: Expected that dirCopyStats.DirsCreated='1'.\n"+
      "Instead, dirCopyStats.TotalFilesProcessed='%v'.\n",
      dirCopyStats.DirCreated)

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if dirCopyStats.ComputeError != nil {
    t.Errorf("Error: Expected that dirCopyStats.ComputeError='nil'.\n"+
      "Instead, dirCopyStats.ComputeError='%v'.\n",
      dirCopyStats.ComputeError.Error())

    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\targetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}
