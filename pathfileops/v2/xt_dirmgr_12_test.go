package pathfileops

import "testing"

func TestDirMgr_MoveDirectoryTree_01(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveDirectoryTree_01"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    t.Errorf("Test Setup Error returned from origSrcDMgr."+
      "CopyDirectoryTree(srcDirMgr, fsc)\n"+
      "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
      srcDirMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  origDtreeInfo, err := origSrcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by origSrcDMgr.FindWalkDirFiles(fsc).\n"+
      "origSrcDMgr='%v'\nError='%v'\n", origSrcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  dirMoveStats,
  errs := srcDirMgr.MoveDirectoryTree(targetDMgr)

  if len(errs) > 0 {
    t.Errorf("Error returned from srcDirMgr.MoveDirectoryTree(targetDMgr)\n"+
      "srcDirMgr='%v'\ntargetDir='%v'\nErrors Follow:\n\n%v",
      srcDirMgr.GetAbsolutePath(),
      targetDMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)
    return
  }

  fsc = FileSelectionCriteria{}

  targetDtreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if origDtreeInfo.FoundFiles.GetNumOfFileMgrs() != targetDtreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("Expected the target directory would contain %v-files.\n"+
      "Error: The target directory tree has %v-files.\n",
      origDtreeInfo.FoundFiles.GetNumOfFileMgrs(), targetDtreeInfo.FoundFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if origDtreeInfo.Directories.GetNumOfDirs() != targetDtreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected the target directory would contain %v-directories.\n"+
      "Error: The target directory tree has %v-directories.\n",
      origDtreeInfo.Directories.GetNumOfDirs(), targetDtreeInfo.Directories.GetNumOfDirs())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if srcDirMgr.DoesAbsolutePathExist() {
    t.Errorf("Error: Expected that 'sourceDir' would NOT exist because all files were moved.\n"+
      "Instead, the source directory DOES EXIST!\n"+
      "Source Dir='%v'", srcDirMgr.GetAbsolutePath())
  }

  expectedNumFilesMoved := uint64(origDtreeInfo.FoundFiles.GetNumOfFileMgrs())

  if expectedNumFilesMoved != dirMoveStats.SourceFilesMoved {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFilesMoved='%v'\n"+
      "Instead, dirMoveStats.SourceFilesMoved='%v'\n",
      expectedNumFilesMoved,
      dirMoveStats.SourceFilesMoved)
  }

  expectedNumFileBytesMoved := origDtreeInfo.FoundFiles.GetTotalFileBytes()

  if expectedNumFileBytesMoved != dirMoveStats.SourceFileBytesMoved {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFileBytesMoved='%v'\n"+
      "Instead, dirMoveStats.SourceFileBytesMoved='%v'\n",
      expectedNumFileBytesMoved,
      dirMoveStats.SourceFileBytesMoved)
  }

  if expectedNumFilesMoved != dirMoveStats.TotalSrcFilesProcessed {
    t.Errorf("ERROR: Expected dirMoveStats.TotalSrcFilesProcessed='%v'\n"+
      "Instead, dirMoveStats.TotalSrcFilesProcessed='%v'\n",
      expectedNumFilesMoved,
      dirMoveStats.TotalSrcFilesProcessed)
  }

  expectedNumDirsMoved := uint64(origDtreeInfo.Directories.GetNumOfDirs())

  if expectedNumDirsMoved != dirMoveStats.TotalDirsProcessed {
    t.Errorf("ERROR: Expected dirMoveStats.TotalDirsProcessed='%v'\n"+
      "Instead, dirMoveStats.TotalDirsProcessed='%v'\n",
      expectedNumDirsMoved,
      dirMoveStats.TotalDirsProcessed)
  }

  if expectedNumDirsMoved != dirMoveStats.DirsCreated {
    t.Errorf("ERROR: Expected dirMoveStats.DirsCreated='%v'\n"+
      "Instead, dirMoveStats.DirsCreated='%v'\n",
      expectedNumDirsMoved,
      dirMoveStats.DirsCreated)
  }

  expectedNumDirsMoved--

  if expectedNumDirsMoved != dirMoveStats.NumOfSubDirectories {
    t.Errorf("ERROR: Expected dirMoveStats.NumOfSubDirectories='%v'\n"+
      "Instead, dirMoveStats.NumOfSubDirectories='%v'\n",
      expectedNumDirsMoved,
      dirMoveStats.NumOfSubDirectories)
  }

  if !dirMoveStats.SourceDirWasDeleted {
    t.Error("ERROR: Expected dirMoveStats.SourceDirWasDeleted='true'.\n" +
      "Instead, dirMoveStats.SourceDirWasDeleted='false'\n")
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectoryTree_02(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveDirectoryTree_02"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    t.Errorf("Test Setup Error returned from origSrcDMgr."+
      "CopyDirectoryTree(srcDirMgr, fsc)\n"+
      "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
      srcDirMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

  _,
    errs = srcDirMgr.MoveDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Error("Expected an error return by srcDirMgr.MoveDirectoryTree(targetDMgr) because\n" +
      "'srcDirMgr' is INVALID! However NO ERROR WAS RETURNED!!!\n")
  }

  srcDirMgr.isInitialized = true

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectoryTree_03(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveDirectoryTree_03"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    t.Errorf("Test Setup Error returned from origSrcDMgr."+
      "CopyDirectoryTree(srcDirMgr, fsc)\n"+
      "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
      srcDirMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

  _,
    errs = srcDirMgr.MoveDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Error("Expected an error return by srcDirMgr.MoveDirectoryTree(targetDMgr) because\n" +
      "'targetDMgr' is INVALID! However NO ERROR WAS RETURNED!!!\n")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectoryTree_04(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveDirectoryTree_04"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  _,
  errs := srcDirMgr.MoveDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Error("Expected an error return by srcDirMgr.MoveDirectoryTree(targetDMgr) because\n" +
      "'srcDirMgr' DOES NOT EXIST! However NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveSubDirectoryTree_01(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_01"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {

    t.Errorf("Test Setup Error returned from origSrcDMgr."+
      "CopyDirectoryTree(srcDirMgr, fsc)\n"+
      "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
      srcDirMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  origDtreeInfo, err := origSrcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by origSrcDMgr.FindWalkDirFiles(fsc).\n"+
      "origSrcDMgr='%v'\nError='%v'\n", origSrcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  dirMoveStats,
  errs := srcDirMgr.MoveSubDirectoryTree(targetDMgr)

  if len(errs) > 0 {

    t.Errorf("Error returned from srcDirMgr.MoveSubDirectoryTree(targetDMgr)\n"+
      "srcDirMgr='%v'\ntargetDir='%v'\nErrors Follow:\n\n'%v'",
      srcDirMgr.GetAbsolutePath(),
      targetDMgr.GetAbsolutePath(),
      srcDirMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(baseDir)
    return
  }

  fsc = FileSelectionCriteria{}

  targetDtreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  origDtreeFiles := origDtreeInfo.FoundFiles.GetNumOfFileMgrs()
  expectedNumOfMovedFiles := uint64(origDtreeFiles - 4)

  if expectedNumOfMovedFiles != uint64(targetDtreeInfo.FoundFiles.GetNumOfFileMgrs()) {
    t.Errorf("Expected the target directory would contain %v-files.\n"+
      "Error: The target directory tree has %v-files.\n",
      origDtreeFiles, targetDtreeInfo.FoundFiles.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if origDtreeInfo.Directories.GetNumOfDirs() != targetDtreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected the target directory would contain %v-directories.\n"+
      "Error: The target directory tree has %v-directories.\n",
      origDtreeInfo.Directories.GetNumOfDirs(), targetDtreeInfo.Directories.GetNumOfDirs())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if !srcDirMgr.DoesAbsolutePathExist() {
    t.Errorf("Error: Expected that 'sourceDir' would still exist because only sub-directories\n"+
      "should have been moved.\nInstead, the source directory was deleted and DOES NOT EXIST!\n"+
      "Source Dir='%v'", srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  sourceDtreeInfo, err := srcDirMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by srcDirMgr.FindWalkDirFiles(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if sourceDtreeInfo.Directories.GetNumOfDirs() != 1 {
    t.Errorf("Expected the number of directories remaining in source directories\n"+
      "would equal '1'.\nInstead, the number directories is %v'.",
      sourceDtreeInfo.Directories.GetNumOfDirs())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  origSrcFInfo, err := origSrcDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by origSrcDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "origSrcDMgr='%v'\nError='%v'\n",
      origSrcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if origSrcFInfo.GetNumOfFileMgrs() != sourceDtreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("Expected number of files in source directory='%v'.\n"+
      "Instead, the number of files remaining in source directory='%v'.\n",
      origSrcFInfo.GetNumOfFileMgrs(), sourceDtreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  if expectedNumOfMovedFiles != dirMoveStats.SourceFilesMoved {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFilesMoved='%v'\n"+
      "Instead, dirMoveStats.SourceFilesMoved='%v'\n",
      expectedNumOfMovedFiles,
      dirMoveStats.SourceFilesMoved)
  }

  expectedNumOfMovedFileBytes :=
    targetDtreeInfo.FoundFiles.GetTotalFileBytes()

  if expectedNumOfMovedFileBytes != dirMoveStats.SourceFileBytesMoved {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFileBytesMoved='%v'\n"+
      "Instead, dirMoveStats.SourceFileBytesMoved='%v'\n",
      expectedNumOfMovedFileBytes,
      dirMoveStats.SourceFileBytesMoved)
  }

  if dirMoveStats.SourceFilesRemaining != 0 {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFilesRemaining='0'\n"+
      "Instead, dirMoveStats.SourceFilesRemaining='%v'\n",
      dirMoveStats.SourceFilesRemaining)
  }

  if dirMoveStats.SourceFileBytesRemaining != 0 {
    t.Errorf("ERROR: Expected dirMoveStats.SourceFileBytesRemaining='0'\n"+
      "Instead, dirMoveStats.SourceFileBytesRemaining='%v'\n",
      dirMoveStats.SourceFileBytesRemaining)
  }

  expectedNumOfSubDirs :=
    uint64(targetDtreeInfo.Directories.GetNumOfDirs() - 1)

  if expectedNumOfSubDirs != dirMoveStats.NumOfSubDirectories {
    t.Errorf("ERROR: Expected dirMoveStats.NumOfSubDirectories='%v'\n"+
      "Instead, dirMoveStats.NumOfSubDirectories='%v'\n",
      expectedNumOfSubDirs,
      dirMoveStats.NumOfSubDirectories)
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveSubDirectoryTree_02(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_02"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
        srcDirMgr.GetAbsolutePath(),
        srcDirMgr.ConsolidateErrors(errs))
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

  _,
    errs = srcDirMgr.MoveSubDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Errorf("Expected an error return from srcDirMgr.MoveSubDirectoryTree(targetDMgr)\n" +
      "because 'srcDirMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  srcDirMgr.isInitialized = true

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return

}

func TestDirMgr_MoveSubDirectoryTree_03(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_03"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nErrors Follow:\n\n'%v'",
        srcDirMgr.GetAbsolutePath(),
        srcDirMgr.ConsolidateErrors(errs))
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

  _,
    errs = srcDirMgr.MoveSubDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Errorf("Expected an error return from srcDirMgr.MoveSubDirectoryTree(targetDMgr)\n" +
      "because 'targetDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveSubDirectoryTree_04(t *testing.T) {

  baseDir := "../../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_03"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(baseDir).\n"+
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  _,
  errs := srcDirMgr.MoveSubDirectoryTree(targetDMgr)

  if len(errs) == 0 {
    t.Errorf("Expected an error return from srcDirMgr.MoveSubDirectoryTree(targetDMgr)\n" +
      "because 'srcDirMgr' Does Not Exist!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

