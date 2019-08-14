package pathfileops

import "testing"

func TestDirMgr_DeleteAll_01(t *testing.T) {

  fh := FileHelper{}
  // Set up target directories and files for deletion!
  origDir, err := dirMgr01TestCreateCheckFiles03DirFiles()

  if err != nil {
    t.Errorf("Error returned by dirMgr01TestCreateCheckFiles03DirFiles(). Error='%v'", err.Error())
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
    return
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

  testDir := "../../checkfiles/TestDirMgr_DeleteAll_02"

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

  testDir := "../../checkfiles/DeleteAllFilesInDir_01"

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

  _,
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

  testDir := "../../checkfiles/iDoNotExist"

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

  _,
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

  testDir := "../../checkfiles/TestDirMgr_DeleteFilesByNamePattern_03"
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

  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../../filesfortest/htmlFilesForTest"
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

  deleteDirStats,
  errArray := dMgr.DeleteAllFilesInDir()

  if len(errArray) > 0 {

    t.Errorf("Errors returned by dMgr.DeleteAllFilesInDir().\n"+
      "testDir='%v'\nErrors Follow:\n\n%v",
      testDir,
      dMgr.ConsolidateErrors(errArray))

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

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.FilesDeleted != 10 {
    t.Errorf("Error: Expected deleteDirStats.FilesDeleted='10'.\n"+
      "Instead, deleteDirStats.FilesDeleted='%v'",
      deleteDirStats.FilesDeleted)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.FilesRemaining != 0 {
    t.Errorf("Error: Expected deleteDirStats.FilesRemaining == '0'\n"+
      "Instead, deleteDirStats.FilesRemaining == '%v'",
      deleteDirStats.FilesRemaining)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.TotalSubDirectories != 1 {
    t.Errorf("Error: Expected deleteDirStats.TotalSubDirectories == '1'\n"+
      "Instead, deleteDirStats.TotalSubDirectories == '%v'\n",
      deleteDirStats.TotalSubDirectories)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.DirectoriesDeleted != 0 {
    t.Errorf("Error: Expected deleteDirStats.DirectoriesDeleted == '0'\n"+
      "Instead, deleteDirStats.DirectoriesDeleted == '%v'\n",
      deleteDirStats.DirectoriesDeleted)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.TotalDirsScanned != 1 {
    t.Errorf("Error: Expected deleteDirStats.TotalDirsScanned == '1'\n"+
      "Instead, deleteDirStats.TotalDirsScanned == '%v'\n",
      deleteDirStats.TotalDirsScanned)

    _ = fh.DeleteDirPathAll(testDir)
    return
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

func TestDirMgr_DeleteAllSubDirectories_01(t *testing.T) {

  expectedNumOfDirectories := 5

  testDir := "../../dirmgrtests/TestDirMgr_DeleteAllSubDirectories_01"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

  setupDir := "../../filesfortest/levelfilesfortest"

  setupDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(setupDir).\n"+
      "setupDir='%v'\nError='%v'", setupDir, err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc := FileSelectionCriteria{}

  dtreeCopyStats,
  errs := setupDMgr.CopyDirectoryTree(testDMgr, true, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by setupDMgr.CopyDirectoryTree(testDMgr,true, fsc)\n"+
      "testDMgr='%v'\nErrors Follow:'\n%v",
      testDMgr.GetAbsolutePath(), testDMgr.ConsolidateErrors(errs))

    _ = testDMgr.DeleteAll()

    return
  }

  setupDInfo, err := setupDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by setupDMgr.FindWalkDirFiles(fsc).\n"+
      "setupDMgr='%v'\nError='%v'\n", setupDMgr.GetAbsolutePath(), err.Error())

    _ = testDMgr.DeleteAll()

    return

  }

  testDInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkDirFiles(fsc).\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = testDMgr.DeleteAll()

    return

  }

  if testDInfo.Directories.GetNumOfDirs() != setupDInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected the number of directories in 'testDir' to = %v.\n"+
      "Intead, the number of directories in 'testDir' = %v.\n"+
      "testDir='%v'.\n", testDInfo.Directories.GetNumOfDirs(), setupDInfo.Directories.GetNumOfDirs(),
      testDMgr.GetAbsolutePath())

    _ = testDMgr.DeleteAll()

    return
  }

  if uint64(expectedNumOfDirectories) != dtreeCopyStats.DirsCopied {
    t.Errorf("ERROR: Expected %v-directories would be copied.\n"+
      "Instead, dtreeCopyStats.DirsCopied='%v'",
      expectedNumOfDirectories, dtreeCopyStats.DirsCopied)
    _ = testDMgr.DeleteAll()

    return
  }

  if uint64(testDInfo.Directories.GetNumOfDirs()) != dtreeCopyStats.DirsCopied {
    t.Errorf("ERROR: testDInfo.Directories.GetNumOfDirs() != dtreeCopyStats.DirsCopied\n"+
      "testDInfo.Directories.GetNumOfDirs()='%v'\n"+
      "dtreeCopyStats.DirsCopied='%v'\n",
      testDInfo.Directories.GetNumOfDirs(),
      dtreeCopyStats.DirsCopied)
  }

  if dtreeCopyStats.FilesCopied != uint64(testDInfo.FoundFiles.GetNumOfFileMgrs()) {
    t.Errorf("ERROR: dtreeCopyStats.FilesCopied='%v'.\n"+
      "However, testDInfo.FoundFiles.GetNumOfFileMgrs()='%v'\n",
      dtreeCopyStats.FilesCopied, testDInfo.FoundFiles.GetNumOfFileMgrs())
  }

  if dtreeCopyStats.FilesNotCopied != 0 {
    t.Errorf("Expected that dtreeCopyStats.FilesNotCopied='0'.\n"+
      "Instead, dtreeCopyStats.FilesNotCopied='%v'!",
      dtreeCopyStats.FilesNotCopied)
  }

  errs = testDMgr.DeleteAllSubDirectories()

  if len(errs) > 0 {
    t.Errorf("Errors returned by testDMgr.DeleteAllSubDirectories()\n"+
      "testDMgr='%v'\nErrors:'\n\n", testDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("%v\n\n", errs[i].Error())
    }

    _ = testDMgr.DeleteAll()

    return
  }

  testDInfo, err = testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by 2nd run of testDMgr.FindWalkDirFiles(fsc).\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = testDMgr.DeleteAll()

    return

  }

  if testDInfo.Directories.GetNumOfDirs() != 1 {
    t.Errorf("After deletion of sub-directories it was expected that\n"+
      "the number of directories in 'testDir' would equal '1'.\n"+
      "Instead, the number of directories in 'testDir' equals '%v'.\n"+
      "testDir='%v'\n", testDInfo.Directories.GetNumOfDirs(), testDMgr.GetAbsolutePath())
  }

  err = testDMgr.DeleteAll()

  if err != nil {

    t.Errorf("Test Clean-Up Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_DeleteAllSubDirectories_02(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteAllSubDirectories_02"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = testDMgr.DeleteAll()

    return
  }

  testDMgr.isInitialized = false

  errs := testDMgr.DeleteAllSubDirectories()

  if len(errs) == 0 {

    t.Error("Expected an error return from testDMgr.DeleteAllSubDirectories()\n" +
      "because 'testDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!\n")
  }

  testDMgr.isInitialized = true

  _ = testDMgr.DeleteAll()

  return

}

func TestDirMgr_DeleteAllSubDirectories_03(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteAllSubDirectories_03"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  errs := testDMgr.DeleteAllSubDirectories()

  if len(errs) == 0 {

    t.Error("Expected an error return from testDMgr.DeleteAllSubDirectories()\n" +
      "because 'testDMgr' DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!\n")
  }

  _ = testDMgr.DeleteAll()

  return
}

func TestDirMgr_DeleteDirectoryTreeFiles_01(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteDirectoryTreeFiles_01"

  sourceDir1 := "../../filesfortest/levelfilesfortest"

  sourceDir2 := "../../filesfortest/htmlFilesForTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    return
  }

  sourceDMgr1, err := DirMgr{}.New(sourceDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir1)\n"+
      "sourceDir1='%v'\nError='%v'\n", sourceDir1, err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  sourceDMgr2, err := DirMgr{}.New(sourceDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir2)\n"+
      "sourceDir2='%v'\nError='%v'\n", sourceDir2, err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := sourceDMgr1.CopyDirectoryTree(testDMgr, true, fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr1.CopyDirectoryTree(testDMgr, true, fsc)\n"+
      "sourceDMgr1='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr1.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))
    _ = testDMgr.DeleteAll()
    return
  }

  _,
    errs = sourceDMgr2.CopyDirectory(testDMgr, fsc, false)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr2.CopyDirectoryTree(testDMgr, true, fsc)\n"+
      "sourceDMgr2='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr2.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))
  }

  testDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  expectedNumOfDirectories := testDtreeInfo.Directories.GetNumOfDirs() - 1

  expectedNumOfDeletedFiles := testDtreeInfo.FoundFiles.GetNumOfFileMgrs()

  expectedNumOfDeletedFileBytes := testDtreeInfo.FoundFiles.GetTotalFileBytes()

  fsc = FileSelectionCriteria{}

  deleteDirStats,
  errs := testDMgr.DeleteDirectoryTreeFiles(fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by testDMgr.DeleteDirectoryTreeFiles(fsc)\n"+
      "sourceDMgr2='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr2.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = testDMgr.DeleteAll()
    return
  }

  if uint64(expectedNumOfDirectories) != deleteDirStats.TotalSubDirectories {
    t.Errorf("Expected numOfSubDirectories='%v'\nInstead, numOfSubDirectories='%v'\n",
      expectedNumOfDirectories, deleteDirStats.TotalSubDirectories)
  }

  if uint64(expectedNumOfDeletedFiles) != deleteDirStats.FilesDeleted {
    t.Errorf("Expected numOfDeletedFiles='%v'\nInstead, numOfDeletedFiles='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeleted)
  }

  if expectedNumOfDeletedFileBytes != deleteDirStats.FilesDeletedBytes {
    t.Errorf("Expected deleteDirStats.FilesDeletedBytes='%v'\n"+
      "Instead, deleteDirStats.FilesDeletedBytes='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeletedBytes)
  }

  if deleteDirStats.FilesRemaining != 0 {
    t.Errorf("Expected numOfRemainingFiles='0'.\nInstead, numOfRemainingFiles='%v'\n",
      deleteDirStats.FilesRemaining)
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_DeleteDirectoryTreeFiles_02(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteDirectoryTreeFiles_02"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := testDMgr.DeleteDirectoryTreeFiles(fsc)

  if len(errs) == 0 {
    t.Error("ERROR: Expected an error return from testDMgr.DeleteDirectoryTreeFiles(fsc)\n" +
      "because testDmgr directory DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgr_DeleteDirectoryTreeFiles_03(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteDirectoryTreeFiles_03"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
  }

  fsc := FileSelectionCriteria{}

  testDMgr.isInitialized = false

  _,
  errs := testDMgr.DeleteDirectoryTreeFiles(fsc)

  if len(errs) == 0 {
    t.Error("ERROR: Expected an error return from testDMgr.DeleteDirectoryTreeFiles(fsc)\n" +
      "because testDmgr is  INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

  testDMgr.isInitialized = true

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_DeleteDirectoryTreeFiles_04(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_DeleteDirectoryTreeFiles_04"

  sourceDir1 := "../../filesfortest/levelfilesfortest"

  sourceDir2 := "../../filesfortest/htmlFilesForTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    return
  }

  sourceDMgr1, err := DirMgr{}.New(sourceDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir1)\n"+
      "sourceDir1='%v'\nError='%v'\n", sourceDir1, err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  sourceDMgr2, err := DirMgr{}.New(sourceDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir2)\n"+
      "sourceDir2='%v'\nError='%v'\n", sourceDir2, err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := sourceDMgr1.CopyDirectoryTree(testDMgr, true, fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr1.CopyDirectoryTree(testDMgr, true, fsc)\n"+
      "sourceDMgr1='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr1.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))
  }

  _,
    errs = sourceDMgr2.CopyDirectory(testDMgr, fsc, false)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr2.CopyDirectoryTree(testDMgr, true, fsc)\n"+
      "sourceDMgr2='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr2.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  testDtreeTxtInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDtreeTxtInfo, err := testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.htm"}

  testDtreeHtmInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDtreeHtmInfo, err :=testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}
  testDtreeAllInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDtreeAllInfo, err :=testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  expectedNumOfDirsScanned := uint64(testDtreeAllInfo.Directories.GetNumOfDirs())

  expectedNumOfSubdirectories := expectedNumOfDirsScanned - 1

  expectedNumOfDeletedFiles := uint64(testDtreeTxtInfo.FoundFiles.GetNumOfFileMgrs())

  expectedNumOfDeletedFileBytes := testDtreeTxtInfo.FoundFiles.GetTotalFileBytes()

  expectedNumOfRemainingFiles := uint64(testDtreeHtmInfo.FoundFiles.GetNumOfFileMgrs())

  expectedNumOfRemainingFileBytes := testDtreeHtmInfo.FoundFiles.GetTotalFileBytes()

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  deleteDirStats,
  errs := testDMgr.DeleteDirectoryTreeFiles(fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by testDMgr.DeleteDirectoryTreeFiles(fsc) fsc='*.txt'\n"+
      "sourceDMgr2='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr2.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))
  }

  if expectedNumOfDirsScanned != deleteDirStats.TotalDirsScanned {
    t.Errorf("Expected deleteDirStats.TotalDirsScanned='%v'\n"+
      "Instead, deleteDirStats.TotalDirsScanned='%v'\n",
      expectedNumOfDirsScanned, deleteDirStats.TotalDirsScanned)
  }

  if 5 == deleteDirStats.DirectoriesDeleted {
    t.Errorf("Expected deleteDirStats.DirectoriesDeleted='5'\n"+
      "Instead, deleteDirStats.DirectoriesDeleted='%v'\n",
      deleteDirStats.TotalDirsScanned)
  }

  if expectedNumOfSubdirectories != deleteDirStats.TotalSubDirectories {
    t.Errorf("Expected deleteDirStats.TotalSubDirectories='%v'\n"+
      "Instead, deleteDirStats.TotalSubDirectories='%v'\n",
      expectedNumOfSubdirectories, deleteDirStats.TotalSubDirectories)
  }

  if expectedNumOfDeletedFiles != deleteDirStats.FilesDeleted {
    t.Errorf("Expected deleteDirStats.FilesDeleted='%v'\n"+
      "Instead, deleteDirStats.FilesDeleted='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeleted)
  }

  if expectedNumOfDeletedFileBytes != deleteDirStats.FilesDeletedBytes {
    t.Errorf("Expected deleteDirStats.FilesDeletedBytes='%v'\n"+
      "Instead, deleteDirStats.FilesDeletedBytes='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeletedBytes)
  }

  if expectedNumOfRemainingFiles != deleteDirStats.FilesRemaining {
    t.Errorf("Expected numOfRemainingFiles='0'.\nInstead, numOfRemainingFiles='%v'\n",
      deleteDirStats.FilesRemaining)
  }

  if expectedNumOfRemainingFileBytes != deleteDirStats.FilesRemainingBytes {
    t.Errorf("Expected deleteDirStats.FilesRemainingBytes='%v'.\n"+
      "Instead, deleteDirStats.FilesRemainingBytes='%v'\n",
      expectedNumOfRemainingFileBytes, deleteDirStats.FilesRemaining)
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by testDMgr.DeleteAll()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
  }

}
