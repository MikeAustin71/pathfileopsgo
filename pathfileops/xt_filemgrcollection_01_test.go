package pathfileops

import (
  "fmt"
  "strings"
  "testing"
  "time"
)

func TestFileMgrCollection_AddFileMgrCollection_01(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
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
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
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
func TestFileMgrCollection_AddFileMgrCollection_02(t *testing.T) {

  fMgrs1 := FileMgrCollection{}
  fMgrs1.fileMgrs = nil

  fMgrs2 := FileMgrCollection{}
  fMgrs2.fileMgrs = nil

  fMgrs1.AddFileMgrCollection(&fMgrs2)

  if fMgrs1.fileMgrs == nil {
    t.Error("ERROR: Expected fMgrs1.fileMgrs!=nil.\n" +
      "Instead, fMgrs1.fileMgrs==nil!!!")
    return
  }

  if len(fMgrs1.fileMgrs) > 0 {
    t.Errorf("Expected len(fMgrs1.fileMgrs)==0.\n" +
      "Instead, len(fMgrs1.fileMgrs)=='%v'\n",
      len(fMgrs1.fileMgrs))
  }

}



func TestFileMgrCollection_AddFileMgr_01(t *testing.T) {
  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
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
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
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

func TestFileMgrCollection_AddFileMgrByDirStrFileNameStr_01(t *testing.T) {

  fMgrs := FileMgrCollection{}

  testDir := "../createFilesTest"

  absTestDir, err := FileHelper{}.MakeAbsolutePath(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by FileHelper{}.MakeAbsolutePath(testDir)\n" +
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  for i:=0; i < 20; i++ {

    fileNameExt := fmt.Sprintf("FileNameNo%02d.txt", i+1)

    err = fMgrs.AddFileMgrByDirStrFileNameStr(absTestDir, fileNameExt)

    if err != nil {
      t.Errorf("Error returned by fMgrs.AddFileMgrByDirStrFileNameStr(" +
        "absTestDir, fileNameExt)\n" +
        "absTestDir='%v'\nfileNameExt='%v'\nError='%v'\n",
        absTestDir, fileNameExt, err.Error())
      return
    }
  }

  if 20 != fMgrs.GetNumOfFiles() {
    t.Errorf("ERROR: Expected number of file managers='20'.\n" +
      "Instead, number of file managers='%v'\n",
      fMgrs.GetNumOfFiles())
    return
  }

  for k:=0; k < 20; k++ {

    fMgr, err := fMgrs.PeekFileMgrAtIndex(k)

    if err != nil {
      t.Errorf("Error returned by fMgrs.PeekFileMgrAtIndex(%v).\n" +
        "Error='%v'\n",
        k, err.Error())
      return
    }

    err = fMgr.IsFileMgrValid("")

    if err != nil {
      t.Errorf("File Manager #%v is INVALID!\n" +
        "Error='%v'\n", k, err.Error())
      return
    }

    fileNameExt := fMgr.GetFileNameExt()

    expectedFileNameExt := fmt.Sprintf("FileNameNo%02d.txt", k+1)

    if expectedFileNameExt != fileNameExt {
      t.Errorf("ERROR: Expected fileNameExt='%v'.\n" +
        "Instead, fileNameExt='%v'\n",
        expectedFileNameExt, fileNameExt)
      return
    }
  }
}

func TestFileMgrCollection_AddFileInfo_01(t *testing.T) {

  var fileNameExt string
  fh := FileHelper{}

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
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

func TestFileMgrCollection_DeleteAtIndex_01(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(f2)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f2). "+
      "f2='%v'  Error='%v' ", f2, err.Error())
  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  for i := 0; i < arrayLen; i++ {

    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find file index # 2 on first pass. DID NOT FIND IT!")
    return
  }

  err = fMgrCol.DeleteAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(2) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fMgrCol.GetNumOfFileMgrs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found file at index # 2. IT WAS NOT DELETED!")
  }

}

func TestFileMgrCollection_DeleteAtIndex_02(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(f1)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f1). "+
      "f1='%v'  Error='%v' ", f1, err.Error())
  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  for i := 0; i < arrayLen; i++ {

    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find file index # 1 on first pass. DID NOT FIND IT!")
    return
  }

  err = fMgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(1) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fMgrCol.GetNumOfFileMgrs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found file at index # 1. IT WAS NOT DELETED!")
  }

}

func TestFileMgrCollection_DeleteAtIndex_03(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(f0)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f0). "+
      "f0='%v'  Error='%v' ", f0, err.Error())
  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  for i := 0; i < arrayLen; i++ {

    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find file index # 0 on first pass. DID NOT FIND IT!")
    return
  }

  err = fMgrCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fMgrCol.GetNumOfFileMgrs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found file at index # 0. IT WAS NOT DELETED!")
  }

}

func TestFileMgrCollection_DeleteAtIndex_04(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  foundDir := false
  fh := FileHelper{}
  searchStr, err := fh.GetAbsPathFromFilePath(f3)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f3). "+
      "f3='%v'  Error='%v' ", f3, err.Error())
  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  for i := 0; i < arrayLen; i++ {

    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != true {
    t.Error("Expected to find file index # 3 on first pass. DID NOT FIND IT!")
    return
  }

  err = fMgrCol.DeleteAtIndex(3)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0) "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fMgrCol.GetNumOfFileMgrs()

  foundDir = false

  for j := 0; j < arrayLen; j++ {
    fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    if searchStr == fileMgr.GetAbsolutePathFileName() {
      foundDir = true
    }

  }

  if foundDir != false {
    t.Error("Error: Found file at index # 3. IT WAS NOT DELETED!")
  }

}

func TestFileMgrCollection_DeleteAtIndex_05(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  if arrayLen != 4 {
    t.Errorf("Error: Expected intial array length='4'. Instead, array length='%v'",
      arrayLen)
  }

  err = fMgrCol.DeleteAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(2). "+
      "Error='%v' ", err.Error())
  }

  err = fMgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(1). "+
      "Error='%v' ", err.Error())
  }

  err = fMgrCol.DeleteAtIndex(1)

  if err != nil {
    t.Errorf("Error returned by 2nd Pass fMgrCol.DeleteAtIndex(1). "+
      "Error='%v' ", err.Error())
  }

  err = fMgrCol.DeleteAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0). "+
      "Error='%v' ", err.Error())
  }

  arrayLen = fMgrCol.GetNumOfFileMgrs()

  if arrayLen != 0 {
    t.Errorf("Error: Expected final array length='0'.  "+
      "Instead, final array length='%v' ", arrayLen)
  }

}

func TestFileMgrCollection_DeleteAtIndex_06(t *testing.T) {
  fMgrCol := FileMgrCollection{}

  err := fMgrCol.DeleteAtIndex(-3)

  if err == nil {
    t.Error("Expected fMgrCol.DeleteAtIndex(-3) to return an error\n" +
      "because the index is less than zero and invalid!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_DeleteAtIndex_07(t *testing.T) {

  f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
  f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
  f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
  f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

  fMgrCol := FileMgrCollection{}.New()

  err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

  if err != nil {
    t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
      "Error='%v' ", err.Error())
    return
  }

  err = fMgrCol.DeleteAtIndex(19)

  if err == nil {
    t.Error("Expected error return from fMgrCol.DeleteAtIndex(19)\n" +
      "because the index, '19', exceeds the actual number of array elements.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}
