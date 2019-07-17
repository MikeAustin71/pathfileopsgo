package pathfileops

import (
  "fmt"
  "strings"
  "testing"
  "time"
)

func TestFileMgrCollection_AddFileMgrCollection(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgrs2 := FileMgrCollection{}

  for i := 0; i < 15; i++ {

    fileNameExt = fmt.Sprintf("testCol2AddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from 2nd run of testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs2.AddFileMgr(fmgr)
  }

  if fMgrs2.GetNumOfFileMgrs() != 15 {
    t.Errorf("Expected fMgrs2 Array Length == 15. Instead fMgrs2.GetNumOfDirs()=='%v'", fMgrs2.GetNumOfFileMgrs())
  }

  fMgrs1.AddFileMgrCollection(&fMgrs2)

  if fMgrs1.GetNumOfFileMgrs() != 25 {
    t.Errorf("Expected augmented fMgrs1 Array Length == 25. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr, err := fMgrs1.PeekLastFileMgr()

  if err != nil {
    t.Errorf("2nd Run: Error returned from fMgrs1.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if fMgr.fileNameExt != "testCol2AddFile_015.txt" {
    t.Errorf("Expected consolidated fMgrs1 to have last fMgr.fileNameExt='testCol2AddFile_015.txt'.\n"+
      "Instead, fMgr.fileNameExt='%v'", fMgr.fileNameExt)
  }

}

func TestFileMgrCollection_AddFileMgr_01(t *testing.T) {
  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  lastFmgr, err := fMgrs.PeekLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekLastDirMgr(). Error='%v'", err)
  }

  if lastFmgr.fileNameExt != "testAddFile_010.txt" {
    t.Errorf("Expected last File Manager to have fileNameExt='testAddFile_010.txt'. Instead fileNameExt='%v'", lastFmgr.fileNameExt)
  }

}

func TestFileMgrCollection_AddFileMgrByPathFile(t *testing.T) {

  var fileNameExt string
  fh := FileHelper{}

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  adjustedPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  fPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
  }

  err = fMgrs.AddFileMgrByPathFileNameExt(fPath)

  if err != nil {
    t.Errorf("Error returned from fMgrs.AddFileMgrByPathFileNameExt(fPath). fPath='%v' Error='%v'", fPath, err.Error())
  }

  fmgr2, err := fMgrs.PeekLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if fmgr2.fileNameExt != "newerFileForTest_01.txt" {
    t.Errorf("Expected Newly Added Fmgr fileNameExt='newerFileForTest_01.txt'.\n"+
      "Instead, fileNameExt='%v'", fmgr2.fileNameExt)
  }

}

func TestFileMgrCollection_AddFileMgrByDirFileNameExt_01(t *testing.T) {

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string
  fileNamesArray := make([]string, 0, 30)

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    fileNamesArray = append(fileNamesArray, strings.ToLower(fileNameExt))

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
    }
  }

  numOfFileManagers := fMgrs.GetNumOfFileMgrs()

  if 10 != numOfFileManagers {
    t.Errorf("ERROR: Expected 10-File Managers.\n"+
      "Instead, the collection actually contains %v-File Managers.\n",
      numOfFileManagers)
  }

  if len(fileNamesArray) != numOfFileManagers {
    t.Errorf("Expected number of elements in 'fileNamesArray'\n"+
      "to equal number of file managers in collection.\n"+
      "They ARE NOT EQUAL!\n"+
      "Length of fileNamesArray='%v'\nNumber of File Managers='%v'\n",
      len(fileNamesArray), numOfFileManagers)
  }

  for k := 0; k < numOfFileManagers; k++ {

    fMgr, err := fMgrs.GetFileMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error retrned by fMgrs.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n",
        k, err.Error())
      return
    }

    fNameExt := strings.ToLower(fMgr.GetFileNameExt())

    if fileNamesArray[k] != fNameExt {
      t.Errorf("Expected File Name Extension='%v'.\n"+
        "Instead, File Name Extension='%v'\n",
        fileNamesArray[k], fNameExt)
    }
  }
}

func TestFileMgrCollection_AddFileMgrByDirFileNameExt_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fileNameExt := ""

  err = fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

  if err == nil {
    t.Error("Expected an error return fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)\n" +
      "because 'fileNameExt' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_AddFileMgrByDirFileNameExt_03(t *testing.T) {

  fMgrs := FileMgrCollection{}

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fileNameExt := "testFile.txt"

  testDMgr.isInitialized = false

  err = fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

  if err == nil {
    t.Error("Expected an error return fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)\n" +
      "because 'testDMgr' is INVALID!.\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_AddFileInfo_01(t *testing.T) {

  var fileNameExt string
  fh := FileHelper{}

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). "+
        "fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }

    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'",
      fMgrs.GetNumOfFileMgrs())
  }

  expectedFileNameExt := "newerFileForTest_01.txt"

  fic := FileInfoPlus{}
  fic.SetName(expectedFileNameExt)
  fic.SetIsDir(false)
  fic.SetSize(123456)
  fic.SetModTime(time.Now().Local())
  fic.SetMode(0666)
  fic.SetSysDataSrc("xyzxyzxyzyzx")
  fic.SetIsFInfoInitialized(true)

  adjustedPath := "../filesfortest/newfilesfortest"

  fPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
  }

  err = fMgrs.AddFileMgrByFileInfo(fPath, fic)

  if err != nil {
    t.Errorf("Error returned from fMgrs.AddFileMgrByFileInfo(fPath, fic). fPath='%v' Error='%v'", fPath, err.Error())

  }

  if fMgrs.GetNumOfFileMgrs() != 11 {
    t.Errorf("Expected fMgrs Array Length == 11.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
    return
  }

  fmgrLast, err := fMgrs.PopLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopLastFileMgr()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fmgrLast.fileNameExt != expectedFileNameExt {
    t.Errorf("Expected fmgrLast.fileNameExt='%v'.\n"+
      "Instead, fmgrLast.fileNameExt='%v'\n",
      expectedFileNameExt, fmgrLast.fileNameExt)
  }

}

func TestFileMgrCollection_CopyFilesToDir_01(t *testing.T) {
  srcPath := "../filesfortest/checkfiles"

  srcDMgr, err := DirMgr{}.New(srcPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcPath)\n"+
      "srcPath='%v'\nError='%v'\n",
      srcPath, err.Error())
    return
  }

  testPath := "../dirmgrtests/TestFileMgrCollection_CopyFilesToDir_01"

  testDMgr, err := DirMgr{}.New(testPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testPath)\n"+
      "testPath='%v'\nError='%v'\n",
      testPath, err.Error())
    return
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Setup Error returned by (1) testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by (1) testDMgr.MakeDir()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  fMgrs, err := srcDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by (1) srcDMgr.FindFilesBySelectCriteria(fsc)\n"+
      "Error='%v'\n", err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  numOfSrcFMgrs := fMgrs.GetNumOfFileMgrs()

  if numOfSrcFMgrs == 0 {
    t.Error("Expected files would be returned from srcDMgr search.\n" +
      "However, Zero Files were returned!\n")
    _ = testDMgr.DeleteAll()
    return
  }

  err = fMgrs.CopyFilesToDir(testDMgr)

  if err != nil {
    t.Errorf("Error returned by err = fMgrs.CopyFilesToDir(testDMgr)\n"+
      "testDMgr='%v'"+
      "Error='%v'\n",
      testDMgr.GetAbsolutePath(),
      err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}

  testDirInfo, err := testDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindFilesBySelectCriteria(fsc)\n"+
      "Error='%v'\n", err)
    _ = testDMgr.DeleteAll()
    return
  }

  if numOfSrcFMgrs != testDirInfo.GetNumOfFileMgrs() {
    t.Errorf("After File Manager Collection Copy Operation,\n"+
      "the number of files copied does NOT match the number of source files.\n"+
      "Expected Number of Files Copied='%v'\n"+
      "Actual Number of Files Copied='%v'\n",
      numOfSrcFMgrs, testDirInfo.GetNumOfFileMgrs())
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Cleanup Error returned by (2) testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
  }
}

func TestFileMgrCollection_CopyFilesToDir_02(t *testing.T) {

  testPath := "../dirmgrtests/TestFileMgrCollection_CopyFilesToDir_02"

  testDMgr, err := DirMgr{}.New(testPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testPath)\n"+
      "testPath='%v'\nError='%v'\n",
      testPath, err.Error())
    return
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Setup Error returned by (1) testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by (1) testDMgr.MakeDir()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fMgrs := FileMgrCollection{}

  err = fMgrs.CopyFilesToDir(testDMgr)

  if err == nil {
    t.Error("Expected an error return by fMgrs.CopyFilesToDir(testDMgr)\n" +
      "because fMgrs is an empty collection.\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Cleanup Error returned by (2) testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
  }
}

func TestFileMgrCollection_CopyOut_01(t *testing.T) {

  fMgrs := FileMgrCollection{}

  _, err := fMgrs.CopyOut()

  if err == nil {
    t.Error("Expected an error return by fMgrs.CopyOut() because\n" +
      "fMgrs is an empty File Manager Collection.\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgrCollection_CopyOut_02(t *testing.T) {

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string
  fileNamesArray := make([]string, 0, 30)

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    filePath := testDMgr.GetAbsolutePathWithSeparator() + fileNameExt

    fileNamesArray = append(fileNamesArray, strings.ToLower(filePath))

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
      return
    }
  }

  fMgrs2, err := fMgrs.CopyOut()

  if err != nil {
    t.Errorf("Error returned by fMgrs.CopyOut()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  origNumOfFMgrs := fMgrs.GetNumOfFileMgrs()

  copiedNumOfFMgrs := fMgrs2.GetNumOfFileMgrs()

  if origNumOfFMgrs != copiedNumOfFMgrs {
    t.Errorf("ERROR: Expected copied number of file managers='%v'\n"+
      "Instead, copied number of file managers='%v'\n",
      origNumOfFMgrs, copiedNumOfFMgrs)
    return
  }

  for k := 0; k < origNumOfFMgrs; k++ {

    origFMgr, err := fMgrs.GetFileMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by fMgrs.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", k, err.Error())
      return
    }

    origPathFileName := strings.ToLower(origFMgr.GetAbsolutePathFileName())

    copiedFMgr, err := fMgrs2.GetFileMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by fMgrs2.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", k, err.Error())
      return
    }

    copiedPathFileName := strings.ToLower(copiedFMgr.GetAbsolutePathFileName())

    if origPathFileName != copiedPathFileName {
      t.Errorf("ERROR: Original File Manager Path File Name NOT EQUAL\n"+
        "to Copied File Manager Path File Name!\n"+
        "Original File Manager Path File Name='%v'\n"+
        "Copied File Manager Path File Name='%v'\n",
        origPathFileName, copiedPathFileName)
    }

  }

}

func TestDirectoryTreeInfo_CopyToDirectoryTree_01(t *testing.T) {

  fh := FileHelper{}
  dir := fh.AdjustPathSlash("../testsrcdir")

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir).\n"+
      "dir='%v'\nError='%v'\n", dir, err.Error())
    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. However, it does NOT exist!\n"+
      "dMgr.path='%v'\ndMgr.AbolutePath='%v'\n",
      dMgr.path, dMgr.absolutePath)
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
      "filesOlderThan, filesNewerThan).\ndir='%v'\nError='%v'\n",
      dir, err.Error())
  }

  baseDir := fh.AdjustPathSlash("../testsrcdir")

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(baseDir) baseDir='%v' Error='%v'", baseDir, err.Error())
  }

  substituteDir := fh.AdjustPathSlash("../testdestdir/destdir")

  substituteDMgr, err := DirMgr{}.New(substituteDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(substituteDir).\n"+
      "substituteDir='%v'  Error='%v'", substituteDir, err.Error())
  }

  newDirTree, err := dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

  if err != nil {
    t.Errorf("Error returned by dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr).\n"+
      "Error='%v'",
      err.Error())
    return
  }

  if len(dirTreeInfo.Directories.dirMgrs) != len(newDirTree.Directories.dirMgrs) {

    t.Errorf("Error: Expected Number of Directories = '%v'.\n"+
      "Instead, Number of NewFromPathFileNameExtStr Directories = '%v'",
      len(dirTreeInfo.Directories.dirMgrs), len(newDirTree.Directories.dirMgrs))
  }

  if len(dirTreeInfo.FoundFiles.fileMgrs) != len(newDirTree.FoundFiles.fileMgrs) {
    t.Errorf("Error: Expected Number of Files = '%v'.\n"+
      "Instead, actual Number of NewFromPathFileNameExtStr Files = '%v'",
      len(dirTreeInfo.FoundFiles.fileMgrs), len(newDirTree.FoundFiles.fileMgrs))
  }

  for i := 0; i < len(newDirTree.FoundFiles.fileMgrs); i++ {
    doesFileExist, err := newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist()

    if err != nil {
      t.Errorf("Error returned by newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist().\n"+
        "i='%v' fileNameExt='%v'  Error='%v'",
        i, newDirTree.FoundFiles.fileMgrs[i].fileNameExt, err.Error())
    }

    if !doesFileExist {
      t.Errorf("Error: Failed to create fileNameExt='%v'.\n"+
        "It does NOT exist in target directory.",
        newDirTree.FoundFiles.fileMgrs[i].fileNameExt)
    }

  }

  err = substituteDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned from substituteDMgr.DeleteAll().\n"+
      "Error='%v'", err.Error())
  }

}

func TestFileMgrCollection_FindFiles(t *testing.T) {

  fmgrCol := FileMgrCollectionTestSetup01()

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.txt"}

  fmgrCol2, err := fmgrCol.FindFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by fmgrCol.FindFiles(fsc).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fmgrCol2.GetNumOfFileMgrs() != 3 {
    t.Errorf("Expected fmgrCol2.GetNumOfDirs()==3 .\n"+
      "Instead, fmgrCol2.GetNumOfDirs()='%v' ",
      fmgrCol2.GetNumOfFileMgrs())
  }

  numOfFoundTextfiles := 0

  for i := 0; i < fmgrCol2.GetNumOfFileMgrs(); i++ {
    if fmgrCol2.fileMgrs[i].fileExt == ".txt" {
      numOfFoundTextfiles++
    }
  }

  if numOfFoundTextfiles != 3 {
    t.Errorf("Expected the number of found text files == 3.\n"+
      "Instead, number of found text files=='%v'", numOfFoundTextfiles)
  }

}

func TestFileMgrCollection_GetFileMgrArray(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt).\n"+
        "fileNameExt='%v'\nError='%v'",
        fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  cntr := 0

  for _, fmgr := range fMgrs.GetFileMgrArray() {

    err := fmgr.IsFileMgrValid("TestFileMgrCollection_GetFileMgrArray Error")

    if err != nil {
      t.Errorf("fmgr is INVALID! file='%v' Error='%v' ",
        fmgr.GetAbsolutePathFileName(), err.Error())
    }

    cntr++
  }

  if cntr != 10 {
    t.Errorf("Error: Expected File Manger Array Count='10'. "+
      "Instead, File Manager Array Count='%v'", cntr)
  }
}

func TestFileMgrCollection_GetFileMgrAtIndex_01(t *testing.T) {

  fm := make([]string, 5, 50)

  fm[0] = "../filesfortest/newfilesfortest/newerFileForTest_02.txt"
  fm[1] = "../filesfortest/newfilesfortest/newerFileForTest_03.txt"
  fm[2] = "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"
  fm[3] = "../filesfortest/oldfilesfortest/006890_WritingFiles.htm"
  fm[4] = "../filesfortest/oldfilesfortest/test.htm"

  fMgrCol := FileMgrCollection{}.New()
  var err error
  fh := FileHelper{}

  for i := 0; i < 5; i++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(fm[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())
    }

    fm[i], err = fh.MakeAbsolutePath(fm[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())

    }

  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  if arrayLen != 5 {
    t.Errorf("Error: Expected Collection array length='5'. "+
      "Instead, array length='%v'. ", arrayLen)
  }

  fMgr, err := fMgrCol.GetFileMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.GetFileMgrAtIndex(2). "+
      "Error='%v' ", err.Error())
    return
  }

  if fm[2] != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected fMgr[2]='%v'. "+
      "Instead, fMgr[2]='%v' ", fm[2], fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_01(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_02(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 0)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_03(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 99)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(10)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_04(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1)

  if err == nil {
    t.Error("Error: Expected an Error to be returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1) " +
      "NO ERROR WAS RETURNED. ")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_05(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 8)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 8) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(8)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(8). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_PopFMgrAtIndex(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrOut, err := fMgrs.PopFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 9 {
    t.Errorf("Expected after Pop Array fMgrs Array Length == 9.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if fmgrOut.fileNameExt != "testAddFile_006.txt" {
    t.Errorf("Expected popped file manger at index=5 to be fileNameExt='testAddFile_006.txt'.\n"+
      "Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
  }

}

func TestFileMgrCollection_PeekFMgrAtIndex(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrOut, err := fMgrs.PeekFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5).\n"+
      "Error='%v'", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected after Peek Array fMgrs Array Length == 10.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if fmgrOut.fileNameExt != "testAddFile_006.txt" {
    t.Errorf("Expected Peek file manger at index=5 to be fileNameExt='testAddFile_006.txt'.\n"+
      "Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
  }

}

func TestFileMgrCollection_PopLastFMgr(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrLast, err := fMgrs.PopLastFileMgr()

  if err != nil {
    t.Errorf("Error returned from fMgrs.PopLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if fmgrLast.fileNameExt != "testAddFile_010.txt" {
    t.Errorf("Expected PopLastDirMgr() to produce fmgrLast.fileNameExt='testAddFile_010.txt'.\n"+
      "Instead, fmgrLast.fileNameExt='%v'", fmgrLast.fileNameExt)
  }

}

// //////////////////////////////////////////////////////////////
// Test Setup Functions
// //////////////////////////////////////////////////////////////
func FileMgrCollectionTestSetup01() FileMgrCollection {

  fh := FileHelper{}
  FMgrs := FileMgrCollection{}

  fPath, _ := fh.MakeAbsolutePath(fh.AdjustPathSlash(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt"))
  fmgr, _ := FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_02.txt"))
  fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_03.txt"))
  fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"))
  fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  fPath, _ = fh.MakeAbsolutePath("../filesfortest/oldfilesfortest/006890_WritingFiles.htm")
  fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  fPath, _ = fh.MakeAbsolutePath("../filesfortest/oldfilesfortest/test.htm")
  fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
  FMgrs.AddFileMgr(fmgr)

  return FMgrs
}

func FileMgrCollectionTestSetupFmgr01(fileNameExt string) (FileMgr, error) {

  ePrefix := "Src File: xt_filemgrcollection_01_test.go  Function: FileMgrCollectionTestSetupFmgr01() "
  fh := FileHelper{}

  pathFileName := "../dirwalktests/dir01/dir02/" + fileNameExt
  adjustedPathFileName := fh.AdjustPathSlash(pathFileName)
  fPath, err := fh.MakeAbsolutePath(adjustedPathFileName)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+"Error return by fh.MakeAbsolutePath(adjustedPathFileName). adjustedPathFileName='%v'  Error='%v'", adjustedPathFileName, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+"Error return by FileMgr{}.NewFromPathFileNameExtStr(fPath). fPath='%v'  Error='%v'", fPath, err.Error())
  }

  return fmgr, nil

}
