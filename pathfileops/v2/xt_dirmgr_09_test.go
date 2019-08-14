package pathfileops

import (
  "os"
  "strings"
  "testing"
  "time"
)

func TestDirMgr_DeleteWalkDirFiles_01(t *testing.T) {

  origDir, err := dirMgr02TestSetupFileWalkDeleteFiles()

  if err != nil {
    t.Errorf("Error returned from dirMgr02TestSetupFileWalkDeleteFiles(). "+
      "Error='%v'", err.Error())
    return
  }

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

  absDir, err := fh.MakeAbsolutePath(baseDirPath)

  if err != nil {
    t.Errorf("fh.MakeAbsolutePath(baseDirPath) baseDirPath='%v' Error='%v'",
      baseDirPath, err.Error())

    _ = fh.DeleteDirPathAll(baseDirPath)

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

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
      dInfo.DeletedFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(absDir)

    return
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
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
      newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
      len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'",
      dInfo.Directories.GetNumOfDirs())
  }

  _ = fh.DeleteDirPathAll(absDir)

  return
}

func TestDirMgr_DeleteWalkDirFiles_02(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

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

  searchPattern := "*.htm"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 3 {
    t.Errorf("Expected to find 3-files deleted. Instead, %v-files were deleted.",
      dInfo.DeletedFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"

  oldFile1Found := false
  oldFile2Found := false
  oldFile3Found := false

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

  }

  if oldFile1Found == false {
    t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
      oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
      oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
      oldFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
      len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  _ = fh.DeleteDirPathAll(absDir)

  return
}

func TestDirMgr_DeleteWalkDirFiles_03(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return

  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Test Setup Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"

  oldFiles := make([]string, 3, 10)

  oldFiles[0] = origDir + string(os.PathSeparator) + oldFile1
  oldFiles[1] = origDir + string(os.PathSeparator) + oldFile2
  oldFiles[2] = origDir + string(os.PathSeparator) + oldFile3

  newAccessTime := time.Date(2017, 2, 25, 10, 30, 30, 1250, time.Local)
  newModTime := time.Date(2017, 2, 25, 10, 30, 30, 1250, time.Local)
  dateFmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

  err = fh.ChangeFileTimes(oldFiles[0], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[0],newAccessTime, newModTime).\n"+
      "oldFiles[0]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[0], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[1], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[1],newAccessTime, newModTime).\n"+
      "oldFiles[1]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[1], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[2], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[2],newAccessTime, newModTime).\n"+
      "oldFiles[2]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[2], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

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

  searchPattern := ""
  fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

  filesOlderThan, err := time.Parse(dateFmtStr, fOlderThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(dateFmtStr, fOlderThanStr). "+
      "fOlderThanStr='%v' Error='%v'", fOlderThanStr, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 3 {
    t.Errorf("Expected to find 3-files deleted. Instead, "+
      "%v-files were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  oldFile1Found := false
  oldFile2Found := false
  oldFile3Found := false

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

  }

  if oldFile1Found == false {
    t.Errorf("Expected deletion of oldFile1='%v'. "+
      "The file was NOT deleted!", oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. "+
      "The file was NOT deleted!", oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. "+
      "The file was NOT deleted!", oldFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of "+
      "Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number "+
      "of directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}

func TestDirMgr_DeleteWalkDirFiles_04(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"

  oldFiles := make([]string, 3, 10)

  oldFiles[0] = origDir + string(os.PathSeparator) + oldFile1
  oldFiles[1] = origDir + string(os.PathSeparator) + oldFile2
  oldFiles[2] = origDir + string(os.PathSeparator) + oldFile3

  newAccessTime := time.Date(2017, 2, 25, 10, 30, 30, 1250, time.Local)
  newModTime := time.Date(2017, 2, 25, 10, 30, 30, 1250, time.Local)
  dateFmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

  err = fh.ChangeFileTimes(oldFiles[0], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[0],newAccessTime, newModTime).\n"+
      "oldFiles[0]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[0], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[1], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[1],newAccessTime, newModTime).\n"+
      "oldFiles[1]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[1], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[2], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[2],newAccessTime, newModTime).\n"+
      "oldFiles[2]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[2], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

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

  fOlderThanStr := "2017-01-30 00:00:00.000000000 -0600 CST"

  searchPattern := "*.htm"
  filesOlderThan, err := time.Parse(dateFmtStr, fOlderThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fOlderThanStr). "+
      "fOlderThanStr='%v' Error='%v'", fOlderThanStr, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 0 {
    t.Errorf("Expected to find 0-files deleted. Instead, %v-files were deleted.",
      dInfo.DeletedFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
      len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of directories "+
      "found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }
}

func TestDirMgr_DeleteWalkDirFiles_05(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  oldFile1 := "test.htm"
  oldFile2 := "006890_WritingFiles.htm"
  oldFile3 := "006870_ReadingFiles.htm"

  oldFiles := make([]string, 3, 10)

  oldFiles[0] = origDir + string(os.PathSeparator) + oldFile1
  oldFiles[1] = origDir + string(os.PathSeparator) + oldFile2
  oldFiles[2] = origDir + string(os.PathSeparator) + oldFile3

  newAccessTime := time.Date(2016, 2, 25, 10, 30, 30, 1250, time.Local)
  newModTime := time.Date(2016, 2, 25, 10, 30, 30, 1250, time.Local)
  dateFmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

  err = fh.ChangeFileTimes(oldFiles[0], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[0],newAccessTime, newModTime).\n"+
      "oldFiles[0]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[0], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 11, 30, 10, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[1], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[1],newAccessTime, newModTime).\n"+
      "oldFiles[1]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[1], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 12, 1, 18, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(oldFiles[2], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(oldFiles[2],newAccessTime, newModTime).\n"+
      "oldFiles[2]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      oldFiles[2], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newFile1 := "newerFileForTest_01.txt"
  newFile2 := "newerFileForTest_02.txt"
  newFile3 := "newerFileForTest_03.txt"

  newFiles := make([]string, 3, 10)

  newFiles[0] = origDir + string(os.PathSeparator) + newFile1
  newFiles[1] = origDir + string(os.PathSeparator) + newFile2
  newFiles[2] = origDir + string(os.PathSeparator) + newFile3

  newAccessTime = time.Date(2018, 1, 25, 10, 30, 30, 1250, time.Local)
  newModTime = time.Date(2018, 1, 25, 10, 30, 30, 1250, time.Local)
  dateFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"

  err = fh.ChangeFileTimes(newFiles[0], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(newFiles[0],newAccessTime, newModTime).\n"+
      "newFiles[0]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      newFiles[0], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 12, 30, 10, 30, 30, 1250, time.Local)
  newModTime = time.Date(2017, 12, 30, 10, 30, 30, 1250, time.Local)

  err = fh.ChangeFileTimes(newFiles[1], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(newFiles[1],newAccessTime, newModTime).\n"+
      "oldFiles[1]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      newFiles[1], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  newAccessTime = time.Date(2017, 12, 1, 19, 54, 31, 1250, time.Local)
  newModTime = time.Date(2017, 12, 1, 19, 54, 31, 1250, time.Local)

  err = fh.ChangeFileTimes(newFiles[2], newAccessTime, newModTime)

  if err != nil {
    t.Errorf("Test Set Error from fh.ChangeFileTimes(newFiles[2],newAccessTime, newModTime).\n"+
      "newFiles[2]='%v'\nnewAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      newFiles[2], newAccessTime.Format(dateFmtStr), newModTime.Format(dateFmtStr), err.Error())

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

  searchPattern := ""

  fNewerThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

  filesOlderThan := time.Time{}

  filesNewerThan, err := time.Parse(dateFmtStr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fNewerThanStr). "+
      "fNewerThanStr='%v' Error='%v'", fNewerThanStr, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 3 {
    t.Errorf("Expected to find 3-files deleted. Instead, %v-files "+
      "were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
  }

  newFile1Found := false
  newFile2Found := false
  newFile3Found := false

  for i := 0; i < dInfo.DeletedFiles.GetNumOfFileMgrs(); i++ {

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

  if newFile1Found == false {
    t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
      newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
      newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
      newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of "+
      "Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}

func TestDirMgr_DeleteWalkDirFiles_06(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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
    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

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
  fNewerThanStr := "2016-07-01 00:00:00.000000000 -0500 CDT"
  filesOlderThan := time.Now()
  filesOlderThan.AddDate(0, 1, 0)

  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fNewerThanStr). "+
      "fNewerThanStr='%v' Error='%v'", fNewerThanStr, err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files "+
      "were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
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
    t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
      len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())

  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}

func TestDirMgr_DeleteWalkDirFiles_07(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

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
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan

  err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))

  if err != nil {
    t.Errorf("Error from fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, "+
      "%v-files were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
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
    t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
      oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
      oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
      oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
      newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
      newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
      newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of Error "+
      "Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}

func TestDirMgr_DeleteWalkDirFiles_08(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

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

  searchPattern1 := "*.txt"
  searchPattern2 := "*.htm"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan

  err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))

  if err != nil {
    t.Errorf("Error returne by fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666)).\n"+
      "Error='%v'\n", err.Error())
  }

  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by dMgr.DeleteWalkDirFiles(fsc)\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
    t.Errorf("Expected to find 6-files deleted. Instead, %v-files "+
      "were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
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
    t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
      oldFile1)
  }

  if oldFile2Found == false {
    t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
      oldFile2)
  }

  if oldFile3Found == false {
    t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
      oldFile3)
  }

  if newFile1Found == false {
    t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
      newFile1)
  }

  if newFile2Found == false {
    t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
      newFile2)
  }

  if newFile3Found == false {
    t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
      newFile3)
  }

  if len(dInfo.ErrReturns) != 0 {
    t.Errorf("Expected zero Error Returns. Instead number of "+
      "Error Returns='%v'", len(dInfo.ErrReturns))
  }

  if dInfo.Directories.GetNumOfDirs() != 3 {
    t.Errorf("Expected 3-directories to be found. Instead, number of "+
      "directories found='%v'", dInfo.Directories.GetNumOfDirs())
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}

func TestDirMgr_DeleteWalkDirFiles_09(t *testing.T) {
  testDir := "../../checkfiles/iDoNotExist"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
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

  err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))

  if err != nil {
    t.Errorf("Error returned by fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666)).\n"+
      "Error='%v'\n", err.Error())
  }

  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err = dMgr.DeleteWalkDirFiles(fsc)

  if err == nil {
    t.Error("Expected an error return from dMgr.DeleteWalkDirFiles(fsc)\n" +
      "because 'dMgr' does NOT exist!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestDirMgr_DeleteWalkDirFiles_10(t *testing.T) {

  fh := FileHelper{}

  baseDirPath := "../../dirwalkdeletetests/dirdelete01"

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

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  if !fh.DoesFileExist(origDir) {
    t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
      origDir)

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origDir)\n"+
      "origDir='%v'\nError='%v'\n",
      origDir, err.Error())
  }

  searchPattern1 := "*.txt"
  searchPattern2 := "*.htm"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan

  err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))

  if err != nil {
    t.Errorf("Error returned by fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666)).\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(absDir)

    return
  }

  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dMgr.isInitialized = false
  _, err = dMgr.DeleteWalkDirFiles(fsc)

  if err == nil {
    t.Error("Expected an error return from dMgr.DeleteWalkDirFiles(fsc)\n" +
      "because 'dMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(absDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error from fh.DeleteDirPathAll(absDir).\n"+
      "absDir='%v'\nError='%v'\n", absDir, err.Error())
  }

}
