package pathfileops

import (
  "fmt"
  "os"
  "strings"
  "testing"
  "time"
)

func TestDirMgr_FindWalkDirFiles_01(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := fh.AdjustPathSlash("../dirwalktests/dir01")

  dir, err := fh.MakeAbsolutePath(baseDirPath)

  if err != nil {
    t.Errorf("fh.MakeAbsolutePath(baseDirPath) baseDirPath='%v' Error='%v'",
      baseDirPath, err.Error())
    return
  }

  err = dirMgr02SetupDirWalkTests()

  if err != nil {
    t.Errorf("Error return from dirMgr02SetupDirWalkTests(). "+
      "Error='%v' ", err.Error())

    _ = fh.DeleteDirPathAll(dir)

    return
  }

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'",
      dir, err.Error())

    _ = fh.DeleteDirPathAll(dir)

    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. Instead, it does NOT exist. "+
      "dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)

    _ = fh.DeleteDirPathAll(dir)

    return

  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
      "filesOlderThan, filesNewerThan). dir='%v' Error='%v'",
      dir, err.Error())

    _ = fh.DeleteDirPathAll(dir)

    return
  }

  if dirTreeInfo.FoundFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files in directory tree. Instead, found %v files.",
      dirTreeInfo.FoundFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(dir)

    return
  }

  foundnewTestFile3 := false
  foundOldReadingFile := false

  for i := 0; i < dirTreeInfo.FoundFiles.GetNumOfFileMgrs(); i++ {
    if dirTreeInfo.FoundFiles.fileMgrs[i].fileNameExt == "newerFileForTest_03.txt" {
      foundnewTestFile3 = true
    }

    if dirTreeInfo.FoundFiles.fileMgrs[i].fileNameExt == "006870_ReadingFiles.htm" {
      foundOldReadingFile = true
    }
  }

  if !foundnewTestFile3 {
    t.Error("Expected FoundFiles would include newerFileForTest_03.txt. It did NOT!")
  }

  if !foundOldReadingFile {
    t.Error("Expected FoundFiles would include 006870_ReadingFiles.htm. It did NOT!")
  }

  if dirTreeInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected the number of directories found would equal 3. It did NOT! "+
      "Number of directories= '%v'", dirTreeInfo.Directories.GetNumOfDirs())
  }

  dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

  foundDir3 := false

  for j := 0; j < dirTreeInfo.Directories.GetNumOfDirs(); j++ {
    if strings.Contains(dirTreeInfo.Directories.dirMgrs[j].path, dir3) {
      foundDir3 = true
    }
  }

  if !foundDir3 {
    t.Errorf("Expected to find Directory %v in dirTreeInfo.dirMgrs. "+
      "This directory was NOT found!", dir3)
  }

  _ = fh.DeleteDirPathAll(dir)

  return

}

func TestDirMgr_FindWalkDirFiles_02(t *testing.T) {

  fh := FileHelper{}
  dir := fh.AdjustPathSlash("../dirwalktests/dir01")
  absDir, err := fh.MakeAbsolutePath(dir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(dir)\n"+
      "dir='%v'\nError='%v'\n",
      dir, err.Error())

    _ = fh.DeleteDirPathAll(dir)

    return
  }

  err = dirMgr02SetupDirWalkTests()

  if err != nil {
    t.Errorf("Error returned by dirMgr02SetupDirWalkTests(). "+
      "Error='%v' ", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  searchPattern := "*Files.htm"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'",
      dir, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. Instead, it does NOT exist. "+
      "dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dWalker, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
      "filesOlderThan, filesNewerThan). dir='%v' Error='%v'", dir, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dWalker.FoundFiles.GetNumOfFileMgrs() != 2 {
    t.Errorf("Expected to find 2-files in directory tree. Instead, "+
      "found %v files.", dWalker.FoundFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  foundWritingFiles := false
  foundOldReadingFile := false

  for i := 0; i < dWalker.FoundFiles.GetNumOfFileMgrs(); i++ {
    if dWalker.FoundFiles.fileMgrs[i].fileNameExt == "006890_WritingFiles.htm" {
      foundWritingFiles = true
    }

    if dWalker.FoundFiles.fileMgrs[i].fileNameExt == "006870_ReadingFiles.htm" {
      foundOldReadingFile = true
    }
  }

  if !foundWritingFiles {
    t.Error("Expected FoundFiles would include 006890_WritingFiles.htm. " +
      "It did NOT!")
  }

  if !foundOldReadingFile {
    t.Error("Expected FoundFiles would include 006870_ReadingFiles.htm. " +
      "It did NOT!")
  }

  if dWalker.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected the number of directories found would equal 3. "+
      "It did NOT! Number of directories= '%v'",
      dWalker.Directories.GetNumOfDirs())
  }

  dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

  foundDir3 := false

  for j := 0; j < dWalker.Directories.GetNumOfDirs(); j++ {
    if strings.Contains(dWalker.Directories.dirMgrs[j].path, dir3) {
      foundDir3 = true
    }
  }

  if !foundDir3 {
    t.Errorf("Expected to find Directory %v in dWalker.dirMgrs. "+
      "This directory was NOT found!", dir3)
  }

  _ = fh.DeleteDirPathAll(absDir)

  return
}

func TestDirMgr_FindWalkDirFiles_03(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../dirwalkdeletetests/dirdelete01"

  absDir, err := fh.MakeAbsolutePath(baseDirPath)

  if err != nil {
    t.Errorf("fh.MakeAbsolutePath(baseDirPath) baseDirPath='%v' Error='%v'",
      baseDirPath, err.Error())

    _ = fh.DeleteDirPathAll(baseDirPath)

    return
  }

  origDir, err := dirMgr02TestSetupFileWalkDeleteFiles()

  if err != nil {
    t.Errorf("Error returned from dirMgr02TestSetupFileWalkDeleteFiles(). "+
      "Error='%v'", err.Error())
    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
      "Error='%v'", origDir, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  searchPattern := ""
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fOlderThanStr := "2016-12-01 19:54:30.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). "+
      "fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fNewerThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr,fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ORSelect()

  dInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.FindWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.FoundFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
      dInfo.FoundFiles.GetNumOfFileMgrs())
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

  for i := 0; i < dInfo.FoundFiles.GetNumOfFileMgrs(); i++ {

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile1) {
      oldFile1Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile2) {
      oldFile2Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile3) {
      oldFile3Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile1) {
      newFile1Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile2) {
      newFile2Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile3) {
      newFile3Found = true
    }

  }

  if oldFile1Found == false {
    t.Errorf("Expected to find oldFile1='%v'. The file was NOT found!", oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected to find oldFile2='%v'. The file was NOT found!", oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected find oldFile3='%v'. The file was NOT found!", oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected to find newFile1='%v'. The file was NOT found!", newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected to find newFile2='%v'. The file was NOT found!", newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected to find newFile3='%v'. The file was NOT found!", newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. "+
      "Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  dir3 := fh.AdjustPathSlash("/dirwalkdeletetests/dirdelete01/dirdelete02/dirdelete03")

  foundDir3 := false

  for j := 0; j < dInfo.Directories.GetNumOfDirs(); j++ {
    if strings.Contains(dInfo.Directories.dirMgrs[j].path, dir3) {
      foundDir3 = true
    }
  }

  if !foundDir3 {
    t.Errorf("Expected to find Directory %v in dInfo.dirMgrs. "+
      "This directory was NOT found!", dir3)
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(absDir)\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

  return

}

func TestDirMgr_FindWalkDirFiles_04(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../dirwalkdeletetests/dirdelete01"

  absDir, err := fh.MakeAbsolutePath(baseDirPath)

  if err != nil {
    t.Errorf("fh.MakeAbsolutePath(baseDirPath) baseDirPath='%v' Error='%v'",
      baseDirPath, err.Error())

    _ = fh.DeleteDirPathAll(baseDirPath)

    return
  }

  origDir, err := dirMgr02TestSetupFileWalkDeleteFiles()

  if err != nil {
    t.Errorf("Error returned from dirMgr02TestSetupFileWalkDeleteFiles(). "+
      "Error='%v'", err.Error())
    return

  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
      origDir, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  searchPattern1 := "*.txt"
  searchPattern2 := "*.htm"

  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ORSelect()

  dInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.FindWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.FoundFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
      dInfo.FoundFiles.GetNumOfFileMgrs())
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

  for i := 0; i < dInfo.FoundFiles.GetNumOfFileMgrs(); i++ {

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile1) {
      oldFile1Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile2) {
      oldFile2Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, oldFile3) {
      oldFile3Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile1) {
      newFile1Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile2) {
      newFile2Found = true
    }

    if strings.Contains(dInfo.FoundFiles.fileMgrs[i].fileNameExt, newFile3) {
      newFile3Found = true
    }

  }

  if oldFile1Found == false {
    t.Errorf("Expected to find oldFile1='%v'. The file was NOT found!", oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected to find oldFile2='%v'. The file was NOT found!", oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected find oldFile3='%v'. The file was NOT found!", oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected to find newFile1='%v'. The file was NOT found!", newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected to find newFile2='%v'. The file was NOT found!", newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected to find newFile3='%v'. The file was NOT found!", newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
      len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  dir3 := fh.AdjustPathSlash("/dirwalkdeletetests/dirdelete01/dirdelete02/dirdelete03")

  foundDir3 := false

  for j := 0; j < dInfo.Directories.GetNumOfDirs(); j++ {
    if strings.Contains(dInfo.Directories.dirMgrs[j].path, dir3) {
      foundDir3 = true
    }
  }

  if !foundDir3 {
    t.Errorf("Expected to find Directory %v in dInfo.dirMgrs. "+
      "This directory was NOT found!", dir3)
  }

  _ = fh.DeleteDirPathAll(absDir)

  return
}

func TestDirMgr_FindWalkDirFiles_05(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  sourceDMgr.isInitialized = false

  _, err = sourceDMgr.FindWalkDirFiles(fsc)

  if err == nil {
    t.Error("Expected an error return from sourceDMgr.FindWalkDirFiles(fsc)\n" +
      "because sourceDMgr is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_FindWalkDirFiles_06(t *testing.T) {

  sourceDir := "../filesfortest/iDoNotExist"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _, err = sourceDMgr.FindWalkDirFiles(fsc)

  if err == nil {
    t.Error("Expected an error return from sourceDMgr.FindWalkDirFiles(fsc)\n" +
      "because sourceDMgr directory DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_FindWalkSubDirFiles_01(t *testing.T) {
  testDir := "../logTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  dTreeInfo, err := testDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkSubDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(),
      err.Error())
    return
  }

  expectedNumOfDirs := 7
  expectedNumOfFiles := 5

  if expectedNumOfFiles != dTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("Error: Expected dTreeInfo.FoundFiles.GetNumOfFileMgrs()='%v'.\n"+
      "Instead, dTreeInfo.FoundFiles.GetNumOfFileMgrs()='%v'\n",
      expectedNumOfFiles, dTreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  if expectedNumOfDirs != dTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Error: Expected dTreeInfo.Directories.GetNumOfDirs()='%v'\n"+
      "Instead, dTreeInfo.Directories.GetNumOfDirs()='%v'\n",
      expectedNumOfDirs, dTreeInfo.Directories.GetNumOfDirs())
  }

  if len(dTreeInfo.ErrReturns) > 0 {
    t.Errorf("dTreeInfo Returned Errors:\n\n%v",
      testDMgr.ConsolidateErrors(dTreeInfo.ErrReturns))
  }

}

func TestDirMgr_FindWalkSubDirFiles_02(t *testing.T) {
  testDir := "../IDoNotExist"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _, err = testDMgr.FindWalkSubDirFiles(fsc)

  if err == nil {
    t.Errorf("Expected an error return from " +
      "testDMgr.FindWalkSubDirFiles(fsc)\n" +
      "because the testDMgr Directory DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!!")
  }
}

/*
  ************************************************************************************
                            Test and Setup Methods
  ************************************************************************************
*/

func dirMgr02TestSetupFileWalkDeleteFiles() (string, error) {

  ePrefix := "xt_dirmgr_08_test.go Func: dirMgr02TestSetupFileWalkDeleteFiles() "

  fh := FileHelper{}

  origDir, err := fh.MakeAbsolutePath("../dirwalkdeletetests/dirdelete01")

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by fh.MakeAbsolutePath(\"../dirwalkdeletetests/dirdelete01\").\n"+
        "Error='%v'\n", err.Error())
  }

  err = fh.DeleteDirPathAll(origDir)

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Test Setup Error returned by fh.DeleteDirPathAll(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
  }

  origFullDir := origDir + string(os.PathSeparator) + "dirdelete02" +
    string(os.PathSeparator) + "dirdelete03"

  // origDir does NOT exist!

  err = fh.MakeDirAll(origFullDir)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned from fh.MakeDirAll(origFullDir).\n"+
        "origDir='%v'\nError='%v'", origFullDir, err.Error())
  }

  dirOldFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/oldfilesfortest")

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(\"../filesfortest/oldfilesfortest\").\n"+
      "Error='%v'\n", err.Error())
  }

  // Copy Old Files

  if !fh.DoesFileExist(dirOldFilesForTest) {
    return "", fmt.Errorf(ePrefix+"Error: Old Files Directory does NOT exist!\n"+
      "dirOldFilesForTest='%v'", dirOldFilesForTest)

  }

  oldFiles := make([]string, 3, 10)

  oldFiles[0] = "test.htm"
  oldFiles[1] = "006890_WritingFiles.htm"
  oldFiles[2] = "006870_ReadingFiles.htm"

  for i := 0; i < 3; i++ {

    srcFile := dirOldFilesForTest + string(os.PathSeparator) + oldFiles[i]
    destFile := origDir + string(os.PathSeparator) + oldFiles[i]

    err = fh.CopyFileByIo(srcFile, destFile)

    if err != nil {
      return "", fmt.Errorf(ePrefix+
        "Error while Copying Source File to  Destination File: fh.CopyFileByIo(srcFile, destFile)\n"+
        "srcFile='%v', destFile='%v'\nError:'%v'",
        srcFile, destFile, err)
    }

  }

  // Copy NewFromPathFileNameExtStr Files
  newFiles := make([]string, 3, 10)

  newFiles[0] = "newerFileForTest_01.txt"
  newFiles[1] = "newerFileForTest_02.txt"
  newFiles[2] = "newerFileForTest_03.txt"

  dirNewFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/newfilesfortest")

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+"Error return from fh.MakeAbsolutePath(dirNewFilesForTest)\n"+
        "dirNewFilesForTest='%v'\nError='%v' ", dirNewFilesForTest, err.Error())
  }

  for i := 0; i < 3; i++ {

    srcFile := dirNewFilesForTest + string(os.PathSeparator) + newFiles[i]
    destFile := origDir + string(os.PathSeparator) + newFiles[i]

    err = fh.CopyFileByIo(srcFile, destFile)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+
          "Error while Copying Source File to  Destination File\n"+
          "fh.CopyFileByIo(srcFile, destFile)\n"+
          "srcFile='%v'\ndestFile='%v'\nError:'%v'",
          srcFile, destFile, err)
    }

  }

  return origDir, nil
}

// Set up Directory Tree, ../dirwalktests for tests!
func dirMgr02SetupDirWalkTests() error {

  ePrefix := "xt_dirmgr08.dirMgr02SetupDirWalkTests() "

  fh := FileHelper{}

  var err error

  err = fh.DeleteDirPathAll("../dirwalktests")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.DeleteDirPathAll(\"../dirwalktests\"). "+
      "Error='%v' ", err.Error())

  }

  dest3, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\"). "+
      "Error='%v' ", err.Error())
  }

  destDirMgr3, err := DirMgr{}.New(dest3)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error from DirMgr{}.New(dest3)) "+
      "dest3='%v' Error='%v' ", dest3, err.Error())
  }

  dest1, err := fh.MakeAbsolutePath("../dirwalktests/dir01")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01\"). "+
      "Error='%v' ", err.Error())
  }

  destDirMgr1, err := DirMgr{}.New(dest1)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error from DirMgr{}.NewFromPathFileNameExtStr(dest1) "+
      "dest1='%v' Error='%v' ", dest1, err.Error())
  }

  dest0, err := fh.MakeAbsolutePath("../dirwalktests")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(\"../dirwalktests\"). "+
      "Error='%v' ", err.Error())
  }

  destDirMgr0, err := DirMgr{}.New(dest0)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error from DirMgr{}.NewFromPathFileNameExtStr(dest1) "+
      "dest1='%v' Error='%v' ", dest0, err.Error())
  }

  // If the directory tree ../dirwalktests/dir01/dir02/dir03
  // does not exist, create it.
  if !destDirMgr3.DoesAbsolutePathExist() {

    err = destDirMgr3.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+"Error returned by destDirMgr.MakeDir(). "+
        "destDir='%v' Error='%v' ", destDirMgr3.absolutePath, err.Error())
    }

  }

  // Delete all files in the ../dirwalktests
  deleteCriteria := FileSelectionCriteria{}
  deleteCriteria.FileNamePatterns = make([]string, 0, 0)
  deleteCriteria.FilesNewerThan = time.Time{}
  deleteCriteria.FilesOlderThan = time.Time{}

  dirTree0Info, err := destDirMgr0.FindWalkDirFiles(deleteCriteria)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error return from destDirMgr0.FindWalkDirFiles(deleteCriteria). "+
      "Error='%v' ", err.Error())
  }

  // If directory tree ../dirwalktests contains files,
  // delete them.
  if dirTree0Info.FoundFiles.GetNumOfFileMgrs() > 0 {

    _, err = destDirMgr0.DeleteWalkDirFiles(deleteCriteria)

    if err != nil {
      return fmt.Errorf(ePrefix+"Error returned by destDirMgr0.DeleteWalkDirFiles(deleteCriteria). "+
        "Error='%v'", err.Error())
    }

  }

  dirNewFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/newfilesfortest")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(\"../filesfortest/newfilesfortest\"). "+
      "Error= '%v' ", err.Error())
  }

  srcNewFilesForTest, err := DirMgr{}.New(dirNewFilesForTest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirNewFilesForTest). "+
      "dirNewFilesForTest='%v' Error='%v' ", dirNewFilesForTest, err.Error())
  }

  if !srcNewFilesForTest.DoesPathExist() {
    return fmt.Errorf(ePrefix+"FATAL ERROR: Directory %v DOES NOT EXIST", dirNewFilesForTest)
  }

  dirOldFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/oldfilesfortest")

  srcOldFilesForTest, err := DirMgr{}.New(dirOldFilesForTest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirOldFilesForTest). "+
      "dirOldFilesForTest='%v' Error='%v' ", dirOldFilesForTest, err.Error())
  }

  if !srcOldFilesForTest.DoesPathExist() {
    return fmt.Errorf(ePrefix+"FATAL ERROR: Directory %v DOES NOT EXIST", dirOldFilesForTest)
  }

  filesToCopySrc := make([]string, 6, 10)
  filesToCopyDest := make([]string, 6, 10)

  filesToCopySrc[0] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
    "newerFileForTest_01.txt"

  filesToCopyDest[0] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "newerFileForTest_01.txt"

  filesToCopySrc[1] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
    "newerFileForTest_02.txt"

  filesToCopyDest[1] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "newerFileForTest_02.txt"

  filesToCopySrc[2] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
    "newerFileForTest_03.txt"

  filesToCopyDest[2] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "newerFileForTest_03.txt"

  filesToCopySrc[3] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
    "test.htm"

  filesToCopyDest[3] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "test.htm"

  filesToCopySrc[4] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
    "006890_WritingFiles.htm"

  filesToCopyDest[4] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "006890_WritingFiles.htm"

  filesToCopySrc[5] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
    "006870_ReadingFiles.htm"

  filesToCopyDest[5] = destDirMgr1.GetAbsolutePathWithSeparator() +
    "006870_ReadingFiles.htm"

  for i := 0; i < len(filesToCopySrc); i++ {

    if !fh.DoesFileExist(filesToCopySrc[i]) {
      return fmt.Errorf(ePrefix+
        "FATAL ERROR: Source File %v DOES NOT EXIST!!!", filesToCopySrc[i])
    }

    err = fh.CopyFileByIoByLink(filesToCopySrc[i], filesToCopyDest[i])

    if err != nil {
      return fmt.Errorf(ePrefix+
        "Error returned by fh.CopyFileByLinkByIo(filesToCopySrc[i], filesToCopyDest[i]). \n"+
        "filesToCopySrc[i]='%v' filesToCopyDest[i]='%v' Error='%v' ",
        filesToCopySrc[i], filesToCopyDest[i], err.Error())
    }

  }

  return nil
}
