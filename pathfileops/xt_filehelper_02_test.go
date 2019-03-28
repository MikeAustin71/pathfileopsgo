package pathfileops

import (
  "fmt"
  "testing"
  "time"
)

func TestFileHelper_FindFilesInPath_01(t *testing.T) {

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

  foundFiles, err := fh.FindFilesInPath(targetDir.GetAbsolutePath(), "*.*")

  lenFoundFiles := len(foundFiles)

  if lenFoundFiles != 5 {
    t.Errorf("Error: Expected to find 5-files. Instead, found %v-files! ",
      lenFoundFiles)
  }

  _ = targetDir.DeleteAll()

}

func TestFileHelper_FindFilesInPath_02(t *testing.T) {

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

  errStrs := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errStrs) > 0 {
    for i := 0; i < len(errStrs); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryTreeOps-Error: %v", errStrs[i])
    }
  }

  foundFiles, err := fh.FindFilesInPath(targetDir.GetAbsolutePath(), "*")

  lenFoundFiles := len(foundFiles)

  if lenFoundFiles != 6 {
    t.Errorf("Error: Expected to find 6-files. Instead, found %v-files! ",
      lenFoundFiles)
  }

  _ = targetDir.DeleteAll()

}

func TestFileHelper_FindFilesInPath_03(t *testing.T) {
  fh := FileHelper{}

  foundFiles, err := fh.FindFilesInPath("", "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(\"\", \"*.*\") " +
      "because first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(\"\", \"*.*\") would be zero length because "+
      "the first input parameter is an empty string."+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_04(t *testing.T) {
  fh := FileHelper{}

  foundFiles, err := fh.FindFilesInPath("   ", "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(\"   \", \"*.*\") " +
      "because first input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(\"    \", \"*.*\") would be zero length because "+
      "the first input parameter consists entirely of empty spaces. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_05(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"\") " +
      "because the second input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"\") would be zero length because "+
      "the second input parameter is an empty string."+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_06(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "    ")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"   \") " +
      "because the second input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"   \") would be zero length because "+
      "the second input parameter consists entirely of empty spaces. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_07(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/iDoNotExistDir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"*.*\") " +
      "because input parameter 'pathFileName' DOES NOT EXIST. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"*.*\") would be zero length because "+
      "the input parameter 'pathFileName' DOES NOT EXIST. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FilterFileName_01(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_02(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
      "fmtstr='%v' fOlderThanStr='%v' Error='%v'",
      fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_03(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)
  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
      "fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_04(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  filesOlderThan := time.Time{}
  fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_05(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}
  fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_06(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("It was expected that this File would NOT be found. It WAS Found. "+
      "Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_07(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("It was expected that this file would NOT be Found. Instead, it WAS found. "+
      "Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_08(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.htm"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_09(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_10(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_11(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("Expected that File would be found. However, File WAS NOT found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }
}

func TestFileHelper_FilterFileName_12(t *testing.T) {

  fh := FileHelper{}
  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{""}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  isFound, err := fh.FilterFileName(nil, fsc)

  if err == nil {
    t.Error("Expected an error return from fh.FilterFileName(nil, fsc) because " +
      "the first input parameter is 'nil'. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isFound {
    t.Error("Expected isFound=='false'. Instead, isFound=='true'. ")
  }
}

func TestFileHelper_FindFilesWalkDirectory_01(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../filesfortest/levelfilesfortest/level_01_dir"
  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles == 0 {
    t.Error("Error: Expected found files to be greater than zero. " +
      "ZERO FILES FOUND!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_02(t *testing.T) {

  fh := FileHelper{}

  fsc := FileSelectionCriteria{}

  startPath := "../filesfortest/levelfilesfortest/level_01_dir"
  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles == 0 {
    t.Error("Error: Expected found files to be greater than zero. " +
      "ZERO FILES FOUND!")
  }
}

func TestFileHelper_FindFilesWalkDirectory_03(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../filesfortest/iDoNotExist/childDoesNotExist"

  _, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err == nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
  }

}

func TestFileHelper_FindFilesWalkDirectory_04(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err := fh.FindFilesWalkDirectory("", fsc)

  if err == nil {
    t.Error("Expected error return from fh.FindFilesWalkDirectory(\"\", fsc) " +
      "because first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_05(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  _, err := fh.FindFilesWalkDirectory("    ", fsc)

  if err == nil {
    t.Error("Expected error return from fh.FindFilesWalkDirectory(\"    \", fsc) " +
      "because first input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_FindFilesWalkDirectory_06(t *testing.T) {

  fh := FileHelper{}

  searchPattern := "*.*"

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  startPath := "../testdestdir"
  dTreeInfo, err := fh.FindFilesWalkDirectory(startPath, fsc)

  if err != nil {
    t.Errorf("Error returned by fh.FindFilesWalkDirectory(startPath, fsc). "+
      "startPath='%v' Error='%v' ", startPath, err.Error())
  }

  numOfDirs := dTreeInfo.Directories.GetNumOfDirs()

  if numOfDirs == 0 {
    t.Error("Error: Expected found directories to be greater than zero. " +
      "ZERO DIRECTORIES FOUND!")
  }

  numOfFiles := dTreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numOfFiles != 0 {
    t.Errorf("Expected zero found files. Instead, files were found. Error!"+
      "numOfFiles='%v' ", numOfFiles)
  }

}

func TestFileHelper_GetAbsPathFromFilePath_01(t *testing.T) {

  _, err := FileHelper{}.GetAbsPathFromFilePath("")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}." +
      "GetAbsPathFromFilePath(\"\") because the input parameter is an " +
      "empty string. However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_GetAbsPathFromFilePath_02(t *testing.T) {

  _, err := FileHelper{}.GetAbsPathFromFilePath("    ")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}." +
      "GetAbsPathFromFilePath(\"     \") because the input parameter consists " +
      "entirely of blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetDotSeparatorIndexesInPathStr_01(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetDotSeparatorIndexesInPathStr("")

  if err == nil {
    t.Error("Expected an error return from fh." +
      "GetDotSeparatorIndexesInPathStr(\"\") because the input parameter is an " +
      "empty string. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetDotSeparatorIndexesInPathStr_02(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetDotSeparatorIndexesInPathStr("   ")

  if err == nil {
    t.Error("Expected an error return from fh." +
      "GetDotSeparatorIndexesInPathStr(\"  \") because the input parameter consists " +
      "entirely of blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetExecutablePathFileName_01(t *testing.T) {

  fh := FileHelper{}

  exePathFileName, err := fh.GetExecutablePathFileName()

  if err != nil {
    t.Errorf("Error returned by fh.GetExecutablePathFileName(). "+
      "Error='%v' ", err.Error())
  }

  if len(exePathFileName) == 0 {
    t.Error("Error: The Executable path an file name returned by " +
      "fh.GetExecutablePathFileName() is a zero length string!")
  }
}

func TestFileHelper_GetFileExtension_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '.go' for valid file extension Instead result='%v' ", result)
  }

}

func TestFileHelper_GetFileExtension_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'.  "+
      "Instead file extension='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("dirmgr_test")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead file extension ='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead, file extension='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
      "Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".gitignore")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). "+
      "commonDir='%v' Error='%v'",
      commonDir, err.Error())
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
      "Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("test.....go")

  expectedExt := ".go"

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if false != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_09(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("")

  expectedExt := ""

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err == nil {
    t.Errorf("Expected an error to be returned from fh.GetFileExt(commonDir). "+
      "commonDir='%v' NO Error was returned!", commonDir)
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != expectedExt {
    t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
  }

}

func TestFileHelper_GetFileExtension_10(t *testing.T) {

  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\dirmgr_test")

  result, isEmpty, err := fh.GetFileExtension(commonDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty GetFileExt for absent file extension to return 'true'. Instead, isEmpty='%v' ", isEmpty)
  }

  if result != "" {
    t.Errorf("Expected GetFileExt to return empty result for absent file extension. Instead file extension='%v' ", result)
  }

}

func TestFileHelper_GetFileExtension_11(t *testing.T) {
  fh := FileHelper{}

  result, isEmpty, err := fh.GetFileExtension("   ")

  if err == nil {
    t.Error("Expected an error to be returned from fh.GetFileExt(\"    \") because " +
      "the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if true != isEmpty {
    t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
      true, isEmpty)
  }

  if result != "" {
    t.Errorf("Expected GetFileExt to return and empty string. Instead, result='%v' ",
      result)
  }
}

func TestFileHelper_GetFileInfoFromPath_01(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetFileInfoFromPath("")

  if err == nil {
    t.Error("Expected error from fh.GetFileInfoFromPath(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileInfoFromPath_02(t *testing.T) {
  fh := FileHelper{}

  _, err := fh.GetFileInfoFromPath("    ")

  if err == nil {
    t.Error("Expected error from fh.GetFileInfoFromPath(\"   \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_01(t *testing.T) {

  fh := FileHelper{}
  target, err := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogtopTest2Text))

  if err != nil {
    t.Error("Error from FileHelper:MakeAbsolutePath(): ", err.Error())
  }

  tStrFmt := "2006-01-02 15:04:05.000000000"

  fileTime, tStr, err := fh.GetFileLastModificationDate(target, tStrFmt)

  if err != nil {
    t.Error("Error from FileHelper:GetFileLastModificationDate():", err.Error())
  }

  fInfo, err := fh.GetFileInfoFromPath(target)

  if err != nil {
    t.Error("Error from FileHelper:GetFileInfoFromPath():", err.Error())
  }

  actualFileTime := fInfo.ModTime()

  expected := actualFileTime.Format(tStrFmt)

  if tStr != expected {
    t.Errorf("Expected Time String for file %v == %v, received time string: %v", target, expected, tStr)
  }

  if !actualFileTime.Equal(fileTime) {
    t.Error(fmt.Sprintf("Expected Time value %v, instead got:", actualFileTime), fileTime)
  }
}

func TestFileHelper_GetFileLastModificationDate_02(t *testing.T) {

  fh := FileHelper{}
  tStrFmt := "2006-01-02 15:04:05.000000000"

  _, _, err := fh.GetFileLastModificationDate("", tStrFmt)

  if err == nil {
    t.Error("Expected error return from fh.GetFileLastModificationDate" +
      "(\"\", tStrFmt) because the first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_03(t *testing.T) {

  fh := FileHelper{}
  tStrFmt := "2006-01-02 15:04:05.000000000"

  _, _, err := fh.GetFileLastModificationDate("   ", tStrFmt)

  if err == nil {
    t.Error("Expected error return from fh.GetFileLastModificationDate" +
      "(\"    \", tStrFmt) because the first input parameter consists entirely of " +
      "blank spaces. However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_GetFileLastModificationDate_04(t *testing.T) {

  fh := FileHelper{}
  target, err := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogtopTest2Text))

  if err != nil {
    t.Error("Error from FileHelper:MakeAbsolutePath(): ", err.Error())
  }

  tStrFmt := ""

  fileTime, tStr, err := fh.GetFileLastModificationDate(target, tStrFmt)

  if err != nil {
    t.Error("Error from FileHelper:GetFileLastModificationDate():", err.Error())
  }

  fInfo, err := fh.GetFileInfoFromPath(target)

  if err != nil {
    t.Error("Error from FileHelper:GetFileInfoFromPath():", err.Error())
  }

  actualFileTime := fInfo.ModTime()

  expected := actualFileTime.Format("2006-01-02 15:04:05.000000000")

  if tStr != expected {
    t.Errorf("Expected Time String for file %v == %v, received time string: %v", target, expected, tStr)
  }

  if !actualFileTime.Equal(fileTime) {
    t.Error(fmt.Sprintf("Expected Time value %v, instead got:", actualFileTime), fileTime)
  }
}

func TestFileHelper_GetFileNameWithExt_01(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_02(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_03(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
      false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_04(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_05(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\")
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
      true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_06(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
  expectedFNameExt := "xt_dirmgr_01_test.go"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_07(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash("dirmgr_test")
  expectedFNameExt := "dirmgr_test"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
      commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_08(t *testing.T) {
  fh := FileHelper{}

  commonDir := fh.AdjustPathSlash(".gitignore")
  expectedFNameExt := ".gitignore"

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). "+
      "commonDir='%v'  Error='%v'", commonDir, err.Error())
  }

  if isEmpty != false {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_09(t *testing.T) {
  fh := FileHelper{}

  commonDir := ""
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err == nil {
    t.Errorf("Expected error return from fh.GetFileNameWithExt(commonDir) " +
      "because 'commonDir' is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }

}

func TestFileHelper_GetFileNameWithExt_10(t *testing.T) {
  fh := FileHelper{}

  commonDir := "   "
  expectedFNameExt := ""

  fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

  if err == nil {
    t.Errorf("Expected error return from fh.GetFileNameWithExt(commonDir) " +
      "because 'commonDir' consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isEmpty != true {
    t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
  }

  if expectedFNameExt != fNameExt {
    t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
      expectedFNameExt, fNameExt)
  }
}
