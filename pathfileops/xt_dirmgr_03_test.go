package pathfileops

import "testing"

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

  deleteDirStats,
  errArray := dMgr.DeleteFilesByNamePattern("*.htm")

  if len(errArray) > 0 {

    t.Errorf("Errors returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nErrors Follow:\n\n%v",
      testDir,
      dMgr.ConsolidateErrors(errArray))

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

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.FilesDeleted != 3 {
    t.Errorf("Expected deleteDirStats.FilesDeleted='3'.\n"+
      "Instead, deleteDirStats.FilesDeleted='%v'.",
      deleteDirStats.FilesDeleted)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.FilesRemaining != 5 {
    t.Errorf("Expected deleteDirStats.FilesRemaining='5'.\n"+
      "Instead, deleteDirStats.FilesRemaining='%v'.",
      deleteDirStats.FilesRemaining)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.TotalFilesProcessed != 8 {
    t.Errorf("Expected deleteDirStats.TotalFilesProcessed='8'.\n"+
      "Instead, deleteDirStats.TotalFilesProcessed='%v'.",
      deleteDirStats.TotalFilesProcessed)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.TotalSubDirectories != 0 {
    t.Errorf("Expected deleteDirStats.TotalSubDirectories='0'.\n"+
      "Instead, deleteDirStats.TotalSubDirectories='%v'.",
      deleteDirStats.TotalSubDirectories)

    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  if deleteDirStats.TotalDirsScanned != 1 {
    t.Errorf("Expected deleteDirStats.TotalDirsScanned='1'.\n"+
      "Instead, deleteDirStats.TotalDirsScanned='%v'.",
      deleteDirStats.TotalDirsScanned)

    _ = fh.DeleteDirPathAll(testDir)
    return
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

  _,
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

  _,
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

  _,
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

  _,
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

  fMgrHtmCollection, err := dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by fMgrHtmCollection, err :=\n"+
      "dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  expectedNumOfFilesDeleted := uint64(fMgrHtmCollection.GetNumOfFileMgrs())

  if expectedNumOfFilesDeleted != 3 {
    t.Errorf("Test Setup Error: Expected to find 3-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", expectedNumOfFilesDeleted)

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  expectedNumOfFileBytesDeleted := fMgrHtmCollection.GetTotalFileBytes()

  fMgrTxtCollection, err := dMgr.FindFilesByNamePattern("*.txt")

  if err != nil {
    t.Errorf("Test Setup Error returned by fMgrTxtCollection, err :=\n"+
      "dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  expectedNumOfFilesRemaining := uint64(fMgrTxtCollection.GetNumOfFileMgrs())

  expectedNumOfFileBytesRemaining := fMgrTxtCollection.GetTotalFileBytes()

  expectedNumOfFilesProcessed := expectedNumOfFilesDeleted + expectedNumOfFilesRemaining

  expectedNumOfSubDirectories := uint64(1)

  expectedNumOfDirsScanned := uint64(1)

  deleteDirStats,
  errArray := dMgr.DeleteFilesByNamePattern("*.htm")

  if len(errArray) > 0 {

    t.Errorf("Errors returned by dMgr.DeleteFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nErrors Follow:\n\n%v",
      testDir,
      dMgr.ConsolidateErrors(errArray))

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  fMgrHtmCollection, err = dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "baseTestDir='%v'\nError='%v'\n", baseTestDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if fMgrHtmCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Error: Expected to find 0-html files in 'testDir'.\n"+
      "Instead, %v-html files were found.", fMgrHtmCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfFileBytesDeleted != deleteDirStats.FilesDeletedBytes {
    t.Errorf("Error: Expected deleteDirStats.FilesDeletedBytes=='%v'.\n"+
      "Instead, deleteDirStats.FilesDeletedBytes=='%v'\n",
      expectedNumOfFileBytesDeleted,
      deleteDirStats.FilesDeleted)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfFilesRemaining != deleteDirStats.FilesRemaining {
    t.Errorf("Error: Expected deleteDirStats.FilesRemaining=='%v'.\n"+
      "Instead, deleteDirStats.FilesDeleted=='%v'\n",
      expectedNumOfFilesRemaining,
      deleteDirStats.FilesRemaining)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if deleteDirStats.DirectoriesDeleted != 0 {
    t.Errorf("Error: Expected deleteDirStats.DirectoriesDeleted=='0'.\n"+
      "Instead, deleteDirStats.DirectoriesDeleted=='%v'\n",
      deleteDirStats.DirectoriesDeleted)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfDirsScanned != deleteDirStats.TotalDirsScanned {

    t.Errorf("Error: Expected deleteDirStats.TotalDirsScanned=='%v'.\n"+
      "Instead, deleteDirStats.TotalDirsScanned=='%v'\n",
      expectedNumOfDirsScanned,
      deleteDirStats.TotalDirsScanned)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfSubDirectories != deleteDirStats.TotalSubDirectories {

    t.Errorf("Error: Expected deleteDirStats.TotalSubDirectories=='%v'.\n"+
      "Instead, deleteDirStats.TotalSubDirectories=='%v'\n",
      expectedNumOfSubDirectories,
      deleteDirStats.TotalDirsScanned)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfFileBytesDeleted != deleteDirStats.FilesDeletedBytes {

    t.Errorf("Error: Expected deleteDirStats.FilesDeletedBytes=='%v'.\n"+
      "Instead, deleteDirStats.FilesDeletedBytes=='%v'\n",
      expectedNumOfFileBytesDeleted,
      deleteDirStats.FilesDeletedBytes)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfFileBytesRemaining != deleteDirStats.FilesRemainingBytes {

    t.Errorf("Error: Expected deleteDirStats.FilesDeletedBytes=='%v'.\n"+
      "Instead, deleteDirStats.FilesDeletedBytes=='%v'\n",
      expectedNumOfFileBytesRemaining,
      deleteDirStats.FilesDeletedBytes)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  if expectedNumOfFilesProcessed != deleteDirStats.TotalFilesProcessed {

    t.Errorf("Error: Expected deleteDirStats.TotalFilesProcessed=='%v'.\n"+
      "Instead, deleteDirStats.TotalFilesProcessed=='%v'\n",
      expectedNumOfFilesProcessed,
      deleteDirStats.TotalFilesProcessed)

    _ = fh.DeleteDirPathAll(baseTestDir)
    return
  }

  fMgrHtmCollection, err = dMgr2Sub.FindFilesByNamePattern("*.txt")

  if err != nil {
    t.Errorf("Error returned by dMgr2Sub.FindFilesByNamePattern(\"*.txt\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(baseTestDir)

    return
  }

  if fMgrHtmCollection.GetNumOfFileMgrs() != 4 {
    t.Errorf("Error expected 4-txt files in the sub-directory. However, the number of\n"+
      "found txt file='%v'", fMgrHtmCollection.GetNumOfFileMgrs())
  }

  err = fh.DeleteDirPathAll(baseTestDir)

  if err != nil {
    t.Errorf("Test File Clean-Up error returned by fh.DeleteDirPathAll(baseTestDir).\n"+
      "baseTestDir='%v'\nError='%v'\n", baseTestDir, err.Error())
  }

  return
}

func TestDirMgr_DeleteFilesBySelectionCriteria_01(t *testing.T) {

  fh := FileHelper{}

  testDir := "../dirmgrtests/TestDirMgr_DeleteFilesBySelectionCriteria_01"

  sourceDir1 := "../filesfortest/levelfilesfortest"

  sourceDir2 := "../filesfortest/htmlFilesForTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

  sourceDMgr1, err := DirMgr{}.New(sourceDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir1)\n"+
      "sourceDir1='%v'\nError='%v'\n", sourceDir1, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  sourceDMgr2, err := DirMgr{}.New(sourceDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir2)\n"+
      "sourceDir2='%v'\nError='%v'\n", sourceDir2, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := sourceDMgr1.CopyDirectory(testDMgr, fsc, false)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr1.CopyDirectory(testDMgr, fsc)\n"+
      "sourceDMgr1='%v'\ntestDMgr='%v'\nErrors Follow:\n\n",
      sourceDMgr1.GetAbsolutePath(), testDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("%v\n", errs[i])
    }
    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  _,
    errs = sourceDMgr2.CopyDirectory(testDMgr, fsc, false)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by sourceDMgr2.CopyDirectoryTree(testDMgr, true, fsc)\n"+
      "sourceDMgr2='%v'\ntestDMgr='%v'\nErrors Follow:\n\n",
      sourceDMgr2.GetAbsolutePath(), testDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("%v\n", errs[i])
    }
    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  testFileMgrInfo, err := testDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  expectedNumOfDeletedFiles := uint64(testFileMgrInfo.GetNumOfFileMgrs())
  expectedNumOfFileBytesDeleted := testFileMgrInfo.GetTotalFileBytes()
  expectedNumOfFilesProcessed := expectedNumOfDeletedFiles

  expectedTotalDirsScanned := uint64(1)
  expectedNumOfDirsWhereFilesDeleted := uint64(1)

  deleteDirStats,
  errs := testDMgr.DeleteFilesBySelectionCriteria(fsc)

  if len(errs) != 0 {
    t.Errorf("Errors returned by testDMgr.DeleteFilesBySelectionCriteria(fsc).\n"+
      "testDMgr='%v'\nErrors Follow:\n\n%v",
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if expectedNumOfDeletedFiles != deleteDirStats.FilesDeleted {
    t.Errorf("Error: Expected numOfDeletedFiles='%v'.\nInstead, numOfDeletedFils='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeleted)
  }

  if deleteDirStats.FilesRemaining != 0 {
    t.Errorf("Error: Expected deleteDirStats.FilesRemaining=0.\n"+
      "Instead, deleteDirStats.FilesRemaining='%v'\n",
      deleteDirStats.FilesRemaining)
  }

  if deleteDirStats.FilesRemainingBytes != 0 {
    t.Errorf("Error: Expected deleteDirStats.FilesRemainingBytes=0.\n"+
      "Instead, deleteDirStats.FilesRemainingBytes='%v'\n",
      deleteDirStats.FilesRemainingBytes)
  }

  if expectedNumOfFilesProcessed != deleteDirStats.TotalFilesProcessed {
    t.Errorf("Error: Expected deleteDirStats.TotalFilesProcessed=%v.\n"+
      "Instead, deleteDirStats.TotalFilesProcessed='%v'\n",
      expectedNumOfFilesProcessed,
      deleteDirStats.TotalFilesProcessed)
  }

  if expectedTotalDirsScanned != deleteDirStats.TotalDirsScanned {
    t.Errorf("Error: Expected deleteDirStats.TotalDirsScanned=%v.\n"+
      "Instead, deleteDirStats.TotalDirsScanned='%v'\n",
      expectedTotalDirsScanned,
      deleteDirStats.TotalDirsScanned)
  }

  if expectedNumOfDirsWhereFilesDeleted != deleteDirStats.NumOfDirsWhereFilesDeleted {
    t.Errorf("Error: Expected deleteDirStats.NumOfDirsWhereFilesDeleted=%v.\n"+
      "Instead, deleteDirStats.NumOfDirsWhereFilesDeleted='%v'\n",
      expectedNumOfDirsWhereFilesDeleted,
      deleteDirStats.NumOfDirsWhereFilesDeleted)
  }

  if expectedNumOfFileBytesDeleted != deleteDirStats.FilesDeletedBytes {
    t.Errorf("Error: Expected deleteDirStats.FilesDeletedBytes=%v.\n"+
      "Instead, deleteDirStats.FilesDeletedBytes='%v'\n",
      expectedNumOfFileBytesDeleted,
      deleteDirStats.FilesDeletedBytes)
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_DeleteFilesBySelectionCriteria_02(t *testing.T) {

  fh := FileHelper{}

  testDir := "../dirmgrtests/TestDirMgr_DeleteFilesBySelectionCriteria_02"

  testDir2 := "../dirmgrtests/TestDirMgr_DeleteFilesBySelectionCriteria_02/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_04_dir"

  sourceDir1 := "../filesfortest/levelfilesfortest"

  sourceDir2 := "../filesfortest/htmlFilesForTest"

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

  testDMgr2, err := DirMgr{}.New(testDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir2)\n"+
      "testDir2='%v'\nError='%v'\n", testDir2, err.Error())
  }

  sourceDMgr1, err := DirMgr{}.New(sourceDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir1)\n"+
      "sourceDir1='%v'\nError='%v'\n", sourceDir1, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  sourceDMgr2, err := DirMgr{}.New(sourceDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir2)\n"+
      "sourceDir2='%v'\nError='%v'\n", sourceDir2, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

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

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  _,
    errs = sourceDMgr2.CopyDirectory(testDMgr2, fsc, false)

  if len(errs) != 0 {

    t.Errorf("Setup Errors returned by sourceDMgr2.CopyDirectoryTree(testDMgr2, true, fsc)\n"+
      "sourceDMgr2='%v'\ntestDMgr2='%v'\nErrors Follow:\n\n%v",
      sourceDMgr2.GetAbsolutePath(),
      testDMgr2.GetAbsolutePath(),
      testDMgr2.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.txt"}

  testFileMgrInfo, err := testDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindFilesBySelectCriteria(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  expectedNumOfDeletedFiles := uint64(testFileMgrInfo.GetNumOfFileMgrs())
  expectedNumOfDeletedFileBytes := testFileMgrInfo.GetTotalFileBytes()

  fsc = FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.htm"}

  remainingFileMgrInfo, err := testDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  expectedNumOfRemainingFiles := uint64(remainingFileMgrInfo.GetNumOfFileMgrs())
  expectedNumOfRemainingFileBytes := remainingFileMgrInfo.GetTotalFileBytes()

  fsc = FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.txt"}

  deleteDirStats,
  errs := testDMgr.DeleteFilesBySelectionCriteria(fsc)

  if len(errs) != 0 {
    t.Errorf("Errors returned by testDMgr.DeleteFilesBySelectionCriteria(fsc).\n"+
      "testDMgr='%v'\nErrors Follow:\n\n%v",
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if expectedNumOfDeletedFiles != deleteDirStats.FilesDeleted {
    t.Errorf("Error: Expected deleteDirStats.FilesDeleted='%v'.\n"+
      "Instead, deleteDirStats.FilesDeleted='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeleted)
  }

  if expectedNumOfDeletedFileBytes != deleteDirStats.FilesDeletedBytes {
    t.Errorf("Error: Expected deleteDirStats.FilesDeletedBytes='%v'.\n"+
      "Instead, deleteDirStats.FilesDeletedBytes='%v'\n",
      expectedNumOfDeletedFileBytes,
      deleteDirStats.FilesDeletedBytes)
  }

  if expectedNumOfRemainingFiles != deleteDirStats.FilesRemaining {
    t.Errorf("Error: Expected numOfDeletedFiles='%v'.\n"+
      "Instead, numOfDeletedFils='%v'\n",
      expectedNumOfDeletedFiles, deleteDirStats.FilesDeleted)
  }

  if expectedNumOfRemainingFileBytes != deleteDirStats.FilesRemainingBytes {
    t.Errorf("Error: Expected deleteDirStats.FilesRemainingBytes='%v'.\n"+
      "Instead, deleteDirStats.FilesRemainingBytes='%v'\n",
      expectedNumOfRemainingFileBytes,
      deleteDirStats.FilesRemainingBytes)
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_DeleteFilesBySelectionCriteria_03(t *testing.T) {

  fh := FileHelper{}

  testDir := "../dirmgrtests/TestDirMgr_DeleteFilesBySelectionCriteria_03"

  sourceDir1 := "../filesfortest/levelfilesfortest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  sourceDMgr1, err := DirMgr{}.New(sourceDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir1)\n"+
      "sourceDir1='%v'\nError='%v'\n", sourceDir1, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

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

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  testDMgr.isInitialized = false

  _,
    errs = testDMgr.DeleteFilesBySelectionCriteria(fsc)

  if len(errs) == 0 {
    t.Errorf("Expected an error return from testDMgr.DeleteFilesBySelectionCriteria(fsc)\n" +
      "because 'testDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_DeleteFilesBySelectionCriteria_04(t *testing.T) {

  fh := FileHelper{}

  testDir := "../dirmgrtests/TestDirMgr_DeleteFilesBySelectionCriteria_03"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := testDMgr.DeleteFilesBySelectionCriteria(fsc)

  if len(errs) == 0 {
    t.Errorf("Expected an error return from testDMgr.DeleteFilesBySelectionCriteria(fsc)\n" +
      "because 'testDMgr' directory DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_DeleteSubDirectoryTreeFiles_01(t *testing.T) {

  testDir := "../dirmgrtests/DeleteSubDirectoryTreeFiles_01"

  sourceDir1 := "../logTest"

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

  beforeDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  beforeNumOfFiles := uint64(beforeDtreeInfo.FoundFiles.GetNumOfFileMgrs())
  beforeNumOfFileBytes := beforeDtreeInfo.FoundFiles.GetTotalFileBytes()

  fsc = FileSelectionCriteria{}

  dTreeInfo,
  errs := testDMgr.DeleteSubDirectoryTreeFiles(fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by testDMgr.DeleteDirectoryTreeFiles(fsc)\n"+
      "sourceDMgr1='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr1.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = testDMgr.DeleteAll()
    return
  }

  afterDtreeInfo, err := testDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindFilesBySelectCriteria(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  numOfFilesRemainingInTopDir := uint64(afterDtreeInfo.GetNumOfFileMgrs())

  if numOfFilesRemainingInTopDir != 4 {
    t.Errorf("Error: Expected 4-files would remain in parent directory.\n"+
      "Instead, the number of files in parent directory is %v\n",
      numOfFilesRemainingInTopDir)
    _ = testDMgr.DeleteAll()
    return
  }

  numOfFileBytesRemainingInTopDir := afterDtreeInfo.GetTotalFileBytes()

  afterSubDtreeInfo, err := testDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by testDMgr.FindWalkSubDirFiles(fsc).\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(),
      err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  numFilesRemainingInSubDirs := afterSubDtreeInfo.FoundFiles.GetNumOfFileMgrs()

  if numFilesRemainingInSubDirs != 0 {
    t.Errorf("Error: Expected 0-files would remain in sub-directory tree.\n"+
      "Instead, the number of files in the sub-directory tree is %v\n",
      numFilesRemainingInSubDirs)
  }

  expectedFilesDeleted :=
    beforeNumOfFiles - numOfFilesRemainingInTopDir

  expectedDeletedFileBytes :=
    beforeNumOfFileBytes - numOfFileBytesRemainingInTopDir

  if expectedFilesDeleted != dTreeInfo.FilesDeleted {
    t.Errorf("Error: Expected dTreeInfo.FilesDeleted='%v'\n"+
      "Instead, dTreeInfo.FilesDeleted='%v'\n",
      expectedFilesDeleted, dTreeInfo.FilesDeleted)
  }

  if expectedDeletedFileBytes != dTreeInfo.FilesDeletedBytes {
    t.Errorf("Error: Expected dTreeInfo.FilesDeletedBytes='%v'\n"+
      "Instead, dTreeInfo.FilesDeletedBytes='%v'\n",
      expectedDeletedFileBytes, dTreeInfo.FilesDeletedBytes)
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Clean-Up Error returned by testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgr_DeleteSubDirectoryTreeFiles_02(t *testing.T) {

  testDir := "../dirmgrtests/iDoNotExist"

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

  fsc := FileSelectionCriteria{}

  _,
  errs := testDMgr.DeleteSubDirectoryTreeFiles(fsc)

  if len(errs) == 0 {
    t.Errorf("ERROR: Expected an error return from " +
      "testDMgr.DeleteSubDirectoryTreeFiles(fsc)\n" +
      "because 'testDMgr' identifies a nonexistent direcory.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestDirMgr_DeleteSubDirectoryTreeFiles_03(t *testing.T) {

  testDir := "../dirmgrtests/DeleteSubDirectoryTreeFiles_03"

  sourceDir1 := "../logTest"

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
    errs = testDMgr.DeleteAllFilesInDir()

  if len(errs) > 0 {
    t.Errorf("Errors retuned by testDMgr.DeleteAllFilesInDir().\n"+
      "testDMgr='%v'\nErrors Follow:\n\n%v",
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}

  beforeDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by beforeDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.htm"}

  beforeHtmDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by beforeHtmDtreeInfo, err := testDMgr.FindWalkDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    _ = testDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.htm"}
  dTreeInfo,
  errs := testDMgr.DeleteSubDirectoryTreeFiles(fsc)

  if len(errs) != 0 {
    t.Errorf("Setup Errors returned by testDMgr.DeleteDirectoryTreeFiles(fsc)\n"+
      "sourceDMgr1='%v'\ntestDMgr='%v'\nErrors Follow:\n\n%v",
      sourceDMgr1.GetAbsolutePath(),
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    _ = testDMgr.DeleteAll()
    return
  }

  expectedNumOfNonHtmFiles :=
    uint64(beforeDtreeInfo.FoundFiles.GetNumOfFileMgrs() -
      beforeHtmDtreeInfo.FoundFiles.GetNumOfFileMgrs())

  if expectedNumOfNonHtmFiles != dTreeInfo.FilesRemaining {
    t.Errorf("Error: Expected dTreeInfo.FilesRemaining='%v'.\n"+
      "Instead, dTreeInfo.FilesRemaining='%v'\n",
      expectedNumOfNonHtmFiles,
      dTreeInfo.FilesRemaining)
  }

  expectedNumOfHtmFilesDeleted :=
    uint64(beforeHtmDtreeInfo.FoundFiles.GetNumOfFileMgrs())

  if expectedNumOfHtmFilesDeleted != dTreeInfo.FilesDeleted {
    t.Errorf("Error: Expected dTreeInfo.FilesDeleted='%v'.\n"+
      "Instead, dTreeInfo.FilesDeleted='%v'\n",
      expectedNumOfHtmFilesDeleted,
      dTreeInfo.FilesDeleted)
  }

  expectedNumFileBytesDeleted :=
    beforeHtmDtreeInfo.FoundFiles.GetTotalFileBytes()

  if expectedNumFileBytesDeleted != dTreeInfo.FilesDeletedBytes {
    t.Errorf("Error: Expected dTreeInfo.FilesDeletedBytes='%v'.\n"+
      "Instead, dTreeInfo.FilesDeletedBytes='%v'\n",
      expectedNumFileBytesDeleted,
      dTreeInfo.FilesDeletedBytes)
  }

  err = testDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Clean-Up Error returned by testDMgr.DeleteAll()\n"+
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgr_DoesAbsolutePathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.absolutePath = " "

  result := dMgr.DoesAbsolutePathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesAbsolutePathExist()\n" +
      "because dMgr.absolutePath consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesAbsolutePathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesAbsolutePathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesAbsolutePathExist()\n" +
      "because dMgr.absolutePath actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}

func TestDirMgr_DoesPathExist_01(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dMgr.path = " "
  dMgr.absolutePath = " "

  result := dMgr.DoesPathExist()

  if result == true {
    t.Error("Expected a value of 'false' to be returned from dMgr.DoesPathExist()\n" +
      "because dMgr.path consists entirely of blank spaces.\n" +
      "However, a value of 'true' was returned instead!\n")
  }

}

func TestDirMgr_DoesPathExist_02(t *testing.T) {

  dMgr, err := DirMgr{}.New("../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../checkfiles\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result := dMgr.DoesPathExist()

  if result == false {
    t.Error("Expected a value of 'true' to be returned from dMgr.DoesPathExist()\n" +
      "because dMgr.path actually exists.\n" +
      "However, a value of 'false' was returned instead!")
  }

}

func TestDirMgr_DoesThisDirectoryExist_01(t *testing.T) {
  testDirStr := "../checkfiles"

  testDMgr, err := DirMgr{}.New(testDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  dirPathDoesExist, err := testDMgr.DoesThisDirectoryExist()

  if err != nil {
    t.Errorf("Error returned by testDMgr.DoesThisDirectoryExist()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    return
  }

  if !dirPathDoesExist {
    t.Errorf("Error: Expected that result would be directory 'Does Exist'.\n"+
      "Instead, result was directory 'Does NOT Exist'!\n"+
      "testDMgr='%v'", testDMgr.GetAbsolutePath())
  }

}

func TestDirMgr_DoesThisDirectoryExist_02(t *testing.T) {
  testDirStr := "../checkfiles/iDoNotExist"

  testDMgr, err := DirMgr{}.New(testDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  dirPathDoesExist, err := testDMgr.DoesThisDirectoryExist()

  if err != nil {
    t.Errorf("Error returned by testDMgr.DoesThisDirectoryExist()\n"+
      "testDMgr='%v'\nError='%v'\n", testDMgr.GetAbsolutePath(), err.Error())
    return
  }

  if dirPathDoesExist {
    t.Errorf("Error: Expected that result would be directory 'Does Not Exist'.\n"+
      "Instead, result was directory 'Does Exist'!\n"+
      "testDMgr='%v'", testDMgr.GetAbsolutePath())
  }

}
