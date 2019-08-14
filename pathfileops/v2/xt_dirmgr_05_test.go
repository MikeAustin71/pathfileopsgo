
package pathfileops

import "testing"


func TestDirMgr_DeleteSubDirectoryTreeFiles_01(t *testing.T) {

  testDir := "../../dirmgrtests/DeleteSubDirectoryTreeFiles_01"

  sourceDir1 := "../../logTest"

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

  testDir := "../../dirmgrtests/iDoNotExist"

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

  testDir := "../../dirmgrtests/DeleteSubDirectoryTreeFiles_03"

  sourceDir1 := "../../logTest"

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

  dMgr, err := DirMgr{}.New("../../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../../checkfiles\")\n"+
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

  dMgr, err := DirMgr{}.New("../../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../../checkfiles\")\n"+
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

  dMgr, err := DirMgr{}.New("../../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../../checkfiles\")\n"+
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

  dMgr, err := DirMgr{}.New("../../checkfiles")

  if err != nil {
    t.Errorf("Test Setup Error returnd by DirMgr{}.New(\"../../checkfiles\")\n"+
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
  testDirStr := "../../checkfiles"

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
  testDirStr := "../../checkfiles/iDoNotExist"

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
