package pathfileops

import (
  "testing"
  "time"
)

func TestFileHelper_FindFilesInPath_01(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../../filesfortest/levelfilesfortest")

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

    _ = targetDir.DeleteAll()

    return
  }

  if targetDir.DoesAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
      return
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOperationCode(0).CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errArray) > 0 {
    for i := 0; i < len(errArray); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryFileOps-Error: %v", errArray[i].Error())
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

  targetDirStr, err := fh.MakeAbsolutePath("../../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../../filesfortest/levelfilesfortest")

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

    _ = targetDir.DeleteAll()
    return
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

  fileOps[0] = FileOperationCode(0).CopySourceToDestinationByIo()

  errArray := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errArray) > 0 {
    for i := 0; i < len(errArray); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryTreeOps-Error: %v", errArray[i])
    }

    _ = targetDir.DeleteAll()

    return
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

  pathFileName := "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

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

  pathFileName := "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

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

  pathFileName := "../../filesfortest/levelfilesfortest/level_01_dir/iDoNotExistDir"

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
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr).\n"+
      "fmtstr='%v'  fModTimeStr='%v'\nError='%v'\n",
      fmtstr, fModTimeStr, err.Error())
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
    t.Errorf("File was NOT found. File should have been found.\n"+
      "fia.Name()='%v fia.ModTime()='%v'\n",
      fia.Name(), fia.ModTime().Format(fmtstr))
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
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr).\n"+
      "fmtstr='%v' fModTimeStr='%v'\nError='%v'", fmtstr, fModTimeStr, err.Error())
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
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr).\n"+
      "fmtstr='%v' fModTimeStr='%v'\nError='%v'\n",
      fmtstr, fModTimeStr, err.Error())
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
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr).\n"+
      "fmtstr='%v' fModTimeStr='%v'\nError='%v'\n",
      fmtstr, fModTimeStr, err.Error())
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

