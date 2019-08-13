package pathfileops

import "testing"


func TestDirMgr_CopyDirectoryTree_01(t *testing.T) {

  expectedNumOfDirectories := 5

  srcDir := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirPathAll(targetDir)\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  dtreeCopyStats,
  errs := srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory path DOES NOT EXIST!!\n")

    return
  }

  srcDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if srcDTreeInfo.Directories.GetNumOfDirs() != targetDTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected %v-directories would be created. Instead, %v-directories were created!",
      srcDTreeInfo.Directories.GetNumOfDirs(), targetDTreeInfo.Directories.GetNumOfDirs())

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if srcDTreeInfo.FoundFiles.GetNumOfFileMgrs() != targetDTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("Expected %v-files would be copied. Instead, %v-files were copied!",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), targetDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  if uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs()) != dtreeCopyStats.FilesCopied {
    t.Errorf("Expected %v-files would be copied.\n"+
      "Instead, numberOfFilesCopied='%v'!",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), dtreeCopyStats.FilesCopied)
  }

  if dtreeCopyStats.FilesNotCopied != 0 {
    t.Errorf("Expected that numberOfFilesNotCopied='0'.\n"+
      "Instead, numberOfFilesNotCopied='%v'!",
      dtreeCopyStats.FilesNotCopied)
  }

  if uint64(expectedNumOfDirectories) != dtreeCopyStats.DirsCopied {
    t.Errorf("Expected that %v-directories would be copied.\n"+
      "Instead, %v-directories were copied.",
      expectedNumOfDirectories, dtreeCopyStats.DirsCopied)
  }

  return
}

func TestDirMgr_CopyDirectoryTree_02(t *testing.T) {
  srcDir := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}
  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  srcDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_03(t *testing.T) {
  srcDir := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(targetDir)\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  targetDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'targetDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!\n\n")
  }

  targetDMgr.isInitialized = true

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_04(t *testing.T) {
  srcDir := "../filesfortest/iDoNotExist"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}
  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}
  _,
  errs := srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error from srcDMgr.CopyDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!!\n")
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_05(t *testing.T) {

  srcDir := "../filesfortest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  srcDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc.FileNamePatterns = []string{"*.htm"}

  dtreeCopyStats,
  errs := srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)\n"+
      "targetDMgr='%v'\nErrors:\n", targetDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      t.Errorf("'%v'\n\n", errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

    return
  }

  if targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory DOES EXIST!!\n" +
      "The target directory should NOT have been created because none of the files\n" +
      "is the source directory matched the file selection criteria.\n" +
      "However, the target directory DOES EXIST! ERROR!!!!\n")
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  if uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs()) != dtreeCopyStats.FilesNotCopied {
    t.Errorf("ERROR: Expected numberOfFilesNotCopied='%v'\n"+
      "Instead, numberOfFilesNotCopied='%v'\n",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), dtreeCopyStats.FilesNotCopied)
  }

  err = fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath())

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDMgr.GetAbsolutePath()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopyDirectoryTree_06(t *testing.T) {

  setUpDir1 := "../filesfortest/levelfilesfortest"

  setUpDMgr1, err := DirMgr{}.New(setUpDir1)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setUpDir1)\n"+
      "setUpDir1='%v'\nError='%v'\n",
      setUpDir1, err.Error())
    return
  }

  setupDir2 := "../filesfortest/htmlFilesForTest"

  setUpDMgr2, err := DirMgr{}.New(setupDir2)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir2)\n"+
      "setupDir2='%v'\nError='%v'\n",
      setupDir2, err.Error())
    return
  }

  srcDir := "../createFilesTest/levelfilesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := setUpDMgr1.CopyDirectoryTree(srcDMgr, false, fsc)

  if len(errs) > 0 {
    t.Errorf("Test Setup Errors returned by setUpDMgr1.CopyDirectoryTree(srcDMgr, false, fsc).\n"+
      "srcDMgr='%v'\nErrors Follow:\n%v", srcDMgr.GetAbsolutePath(),
      DirMgr{}.ConsolidateErrors(errs).Error())
    return
  }

  srcHtmlDir := "../createFilesTest/levelfilesfortest/level_01_dir/level_02_dir/htmlFilesForTest"

  srcHtmlDMgr, err := DirMgr{}.New(srcHtmlDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcHtmlDir).\n"+
      "srcHtmlDir='%v'\nError='%v'", srcHtmlDir, err.Error())
    return
  }

  fsc = FileSelectionCriteria{}

  _,
    errs = setUpDMgr2.CopyDirectory(srcHtmlDMgr, fsc, false)

  if len(errs) > 0 {
    t.Errorf("Test Setup Errors returned by setUpDMgr2.CopyDirectory(srcHtmlDMgr, fsc).\n"+
      "srcHtmlDMgr='%v'\nErrors Follow:\n%v\n",
      srcHtmlDMgr.GetAbsolutePath(),
      DirMgr{}.ConsolidateErrors(errs).Error())
    return
  }

  targetDir := "../dirmgrtests/levelfilesfortest"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned from fh.DeleteDirPathAll(targetDir)\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  // Copy '.txt' files only to targetDMgr
  dtreeCopyStats,
  errs := srcDMgr.CopyDirectoryTree(
    targetDMgr,
    false,
    fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n%v",
      targetDMgr.GetAbsolutePath(),
      DirMgr{}.ConsolidateErrors(errs).Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Errorf("ERROR: The target directory path DOES NOT EXIST!!\n"+
      "Number Of FilesCopied='%v'\n", dtreeCopyStats.FilesCopied)

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  srcTextDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by\n"+
      "srcTextDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'",
      srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  expectedNumOfCopiedFiles := srcTextDTreeInfo.FoundFiles.GetNumOfFileMgrs()

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.htm"}

  srcHtmlDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by\n"+
      "srcHtmlDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'",
      srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)
    return
  }

  expectedNumOfFilesNotCopied := srcHtmlDTreeInfo.FoundFiles.GetNumOfFileMgrs()

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)
    return
  }

  expectedNumOfDirectoriesCopied := srcTextDTreeInfo.Directories.GetNumOfDirs() - 1

  if expectedNumOfDirectoriesCopied != targetDTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("Expected %v-directories would be created. Instead, %v-directories were created!\n"+
      "targetDTreeInfo.Directories.GetNumOfDirs()='%v'\n",
      expectedNumOfDirectoriesCopied,
      targetDTreeInfo.Directories.GetNumOfDirs(),
      targetDTreeInfo.Directories.GetNumOfDirs())

  }

  if uint64(expectedNumOfCopiedFiles) != dtreeCopyStats.FilesCopied {
    t.Errorf("Expected %v-files would be copied.\n"+
      "Instead, numberOfFilesCopied-'%v'\n",
      expectedNumOfCopiedFiles, dtreeCopyStats.FilesCopied)
  }

  if uint64(expectedNumOfFilesNotCopied) != dtreeCopyStats.FilesNotCopied {
    t.Errorf("Expected %v-files would NOT be copied.\n"+
      "Instead, numberOfFilesNotCopied='%v'!",
      expectedNumOfFilesNotCopied, dtreeCopyStats.FilesNotCopied)
  }

  if uint64(expectedNumOfDirectoriesCopied) != dtreeCopyStats.DirsCopied {
    t.Errorf("Expected that %v-directories would be copied.\n"+
      "Instead, %v-directories were copied.",
      expectedNumOfDirectoriesCopied, dtreeCopyStats.DirsCopied)
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(fh.DeleteDirPathAll(targetDir)\n"+
      "Target Directory Path='%v'\nError='%v'\n", targetDir, err.Error())
  }

  err = fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(fh.DeleteDirPathAll(srcDir)\n"+
      "Source Directory Path='%v'\nError='%v'\n", srcDir, err.Error())
  }

  return
}

func TestDirMgr_CopyIn_01(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  err := fh.DeleteDirPathAll(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

  err = fh.DeleteDirPathAll(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir2).\n"+
      "origDir2='%v'\nError='%v'\n", origDir2, err.Error())
    return
  }

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir2).\n"+
      "origDir2=='%v'\nError='%v'\n", origDir2, err.Error())

    _ = fh.DeleteDirPathAll(origDir)
    _ = fh.DeleteDirPathAll(origDir2)
    return
  }

  if dMgr2.path != origDir2 {
    t.Errorf("Expected original dMgr2.path='%v'.\n"+
      "Instead, dMgr2.path='%v'\n", origDir2, dMgr2.path)
    _ = fh.DeleteDirPathAll(origDir)
    _ = fh.DeleteDirPathAll(origDir2)
    return
  }

  dMgr2.CopyIn(&dMgr)

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyIn(), expected dMgr2.isFInfoInitialized='%v'.\n"+
      "Instead, dMgr2.isFInfoInitialized='%v'.\n",
      dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyIn(), expected dMgr2.isFInfoInitialized='%v'.\n"+
      "Instead, dMgr2.isFInfoInitialized='%v'.", dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.originalPath != dMgr.originalPath {
    t.Errorf("After CopyIn(), expected dMgr2.originalPath='%v'.\n"+
      "Instead, dMgr2.originalPath='%v'.", dMgr.originalPath, dMgr2.originalPath)
  }

  if dMgr2.path != dMgr.path {
    t.Errorf("After CopyIn(), expected dMgr2.path='%v'.\n"+
      "Instead, dMgr2.path='%v'.", dMgr.path, dMgr2.path)
  }

  if dMgr2.isPathPopulated != dMgr.isPathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isPathPopulated='%v'.\n"+
      "Instead, dMgr2.isPathPopulated='%v'.", dMgr.isPathPopulated, dMgr2.isPathPopulated)
  }

  if dMgr2.doesPathExist != dMgr.doesPathExist {
    t.Errorf("After CopyIn(), expected dMgr2.doesPathExist='%v'.\n"+
      "Instead, dMgr2.doesPathExist='%v'.", dMgr.doesPathExist, dMgr2.doesPathExist)
  }

  if dMgr2.parentPath != dMgr.parentPath {
    t.Errorf("After CopyIn(), expected dMgr2.parentPath='%v'.\n"+
      "Instead, dMgr2.parentPath='%v'.", dMgr.parentPath, dMgr2.parentPath)
  }

  if dMgr2.isParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isParentPathPopulated='%v'.\n"+
      "Instead, dMgr2.isParentPathPopulated='%v'.",
      dMgr.isParentPathPopulated, dMgr2.isParentPathPopulated)
  }

  if dMgr2.absolutePath != dMgr.absolutePath {
    t.Errorf("After CopyIn(), expected dMgr2.absolutePath='%v'.\n"+
      "Instead, dMgr2.absolutePath='%v'.", dMgr.absolutePath, dMgr2.absolutePath)
  }

  if dMgr2.isAbsolutePathPopulated != dMgr.isAbsolutePathPopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isAbsolutePathPopulated='%v'.\n"+
      "Instead, dMgr2.isAbsolutePathPopulated='%v'.",
      dMgr.isAbsolutePathPopulated, dMgr2.isAbsolutePathPopulated)
  }

  if dMgr2.doesAbsolutePathExist != dMgr.doesAbsolutePathExist {
    t.Errorf("After CopyIn(), expected dMgr2.doesAbsolutePathExist='%v'.\n"+
      "Instead, dMgr2.doesAbsolutePathExist='%v'.",
      dMgr.doesAbsolutePathExist, dMgr2.doesAbsolutePathExist)
  }

  if dMgr2.isAbsolutePathDifferentFromPath != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("After CopyIn(), expected dMgr2.isAbsolutePathDifferentFromPath='%v'.\n"+
      "Instead, dMgr2.isAbsolutePathDifferentFromPath='%v'.",
      dMgr.isAbsolutePathDifferentFromPath, dMgr2.isAbsolutePathDifferentFromPath)
  }

  if dMgr2.volumeName != dMgr.volumeName {
    t.Errorf("After CopyIn(), expected dMgr2.volumeName='%v'.\n"+
      "Instead, dMgr2.volumeName='%v'.", dMgr.volumeName, dMgr2.volumeName)
  }

  if dMgr2.isVolumePopulated != dMgr.isVolumePopulated {
    t.Errorf("After CopyIn(), expected dMgr2.isVolumePopulated='%v'.\n"+
      "Instead, dMgr2.isVolumePopulated='%v'.", dMgr.isVolumePopulated, dMgr2.isVolumePopulated)
  }

  if !dMgr2.Equal(&dMgr) {
    t.Error("After CopyIn(), expected dMgr2 to EQUAL dMgr. It did NOT!")
  }

  err = fh.DeleteDirPathAll(origDir)

  if err != nil {
    t.Errorf("Test Clean-up Error returned by fh.DeleteDirPathAll(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
  }

  err = fh.DeleteDirPathAll(origDir2)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(origDir2).\n"+
      "origDir2='%v'\nError='%v'\n", origDir2, err.Error())
  }

  return
}

func TestDirMgr_CopyOut_01(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles1")

  err := fh.DeleteDirPathAll(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\n"+
      "Error='%v'", origDir, err.Error())
    return
  }

  origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

  err = fh.DeleteDirPathAll(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir2).\n"+
      "origDir2='%v'\nError='%v'\n", origDir2, err.Error())
    return
  }

  dMgr2, err := DirMgr{}.New(origDir2)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir2).\n"+
      "origDir2=='%v'\nError='%v'", origDir2, err.Error())

    _ = fh.DeleteDirPathAll(origDir)
    _ = fh.DeleteDirPathAll(origDir2)
    return
  }

  if dMgr2.path != origDir2 {
    t.Errorf("Expected original dMgr2.path='%v'.\n"+
      "Instead, dMgr2.path='%v'",
      origDir2, dMgr2.path)
    _ = fh.DeleteDirPathAll(origDir)
    _ = fh.DeleteDirPathAll(origDir2)
    return
  }

  dMgr2 = dMgr.CopyOut()

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.isFInfoInitialized='%v'.\n"+
      "Instead, dMgr2.isFInfoInitialized='%v'.",
      dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.isInitialized != dMgr.isInitialized {
    t.Errorf("After CopyOut(), expected dMgr2.isFInfoInitialized='%v'.\n"+
      "Instead, dMgr2.isFInfoInitialized='%v'.",
      dMgr.isInitialized, dMgr2.isInitialized)
  }

  if dMgr2.originalPath != dMgr.originalPath {
    t.Errorf("After CopyOut(), expected dMgr2.originalPath='%v'.\n"+
      "Instead, dMgr2.originalPath='%v'.",
      dMgr.originalPath, dMgr2.originalPath)
  }

  if dMgr2.path != dMgr.path {
    t.Errorf("After CopyOut(), expected dMgr2.path='%v'.\n"+
      "Instead, dMgr2.path='%v'.",
      dMgr.path, dMgr2.path)
  }

  if dMgr2.isPathPopulated != dMgr.isPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isPathPopulated='%v'.\n"+
      "Instead, dMgr2.isPathPopulated='%v'.",
      dMgr.isPathPopulated, dMgr2.isPathPopulated)
  }

  if dMgr2.doesPathExist != dMgr.doesPathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesPathExist='%v'.\n"+
      "Instead, dMgr2.doesPathExist='%v'.",
      dMgr.doesPathExist, dMgr2.doesPathExist)
  }

  if dMgr2.parentPath != dMgr.parentPath {
    t.Errorf("After CopyOut(), expected dMgr2.parentPath='%v'.\n"+
      "Instead, dMgr2.parentPath='%v'.", dMgr.parentPath, dMgr2.parentPath)
  }

  if dMgr2.isParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isParentPathPopulated='%v'.\n"+
      "Instead, dMgr2.isParentPathPopulated='%v'.",
      dMgr.isParentPathPopulated, dMgr2.isParentPathPopulated)
  }

  if dMgr2.absolutePath != dMgr.absolutePath {
    t.Errorf("After CopyOut(), expected dMgr2.absolutePath='%v'.\n"+
      "Instead, dMgr2.absolutePath='%v'.", dMgr.absolutePath, dMgr2.absolutePath)
  }

  if dMgr2.isAbsolutePathPopulated != dMgr.isAbsolutePathPopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathPopulated='%v'.\n"+
      "Instead, dMgr2.isAbsolutePathPopulated='%v'.",
      dMgr.isAbsolutePathPopulated, dMgr2.isAbsolutePathPopulated)
  }

  if dMgr2.doesAbsolutePathExist != dMgr.doesAbsolutePathExist {
    t.Errorf("After CopyOut(), expected dMgr2.doesAbsolutePathExist='%v'.\n"+
      "Instead, dMgr2.doesAbsolutePathExist='%v'.",
      dMgr.doesAbsolutePathExist, dMgr2.doesAbsolutePathExist)
  }

  if dMgr2.isAbsolutePathDifferentFromPath != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("After CopyOut(), expected dMgr2.isAbsolutePathDifferentFromPath='%v'.\n"+
      "Instead, dMgr2.isAbsolutePathDifferentFromPath='%v'.",
      dMgr.isAbsolutePathDifferentFromPath, dMgr2.isAbsolutePathDifferentFromPath)
  }

  if dMgr2.volumeName != dMgr.volumeName {
    t.Errorf("After CopyOut(), expected dMgr2.volumeName='%v'.\n"+
      "Instead, dMgr2.volumeName='%v'.",
      dMgr.volumeName, dMgr2.volumeName)
  }

  if dMgr2.isVolumePopulated != dMgr.isVolumePopulated {
    t.Errorf("After CopyOut(), expected dMgr2.isVolumePopulated='%v'.\n"+
      "Instead, dMgr2.isVolumePopulated='%v'.",
      dMgr.isVolumePopulated, dMgr2.isVolumePopulated)
  }

  if !dMgr2.Equal(&dMgr) {
    t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!")
  }

  err = fh.DeleteDirPathAll(origDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir).\n"+
      "origDir='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(origDir2)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(origDir2).\n"+
      "origDir2='%v'\nError='%v'\n", origDir2, err.Error())
    return
  }

}

func TestDirMgr_CopySubDirectoryTree_01(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_01"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  dTreeCopyStats,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = targetDMgr.DeleteAll()
    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory path DOES NOT EXIST!!\n")

    return
  }

  srcDTreeInfo, err := srcDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by srcDMgr.FindWalkSubDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  targetDTreeInfo, err := targetDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkSubDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  if srcDTreeInfo.Directories.GetNumOfDirs() != targetDTreeInfo.Directories.GetNumOfDirs() {
    t.Errorf("ERROR: Expected %v-directories would be created. Instead, %v-directories were created!",
      srcDTreeInfo.Directories.GetNumOfDirs(), targetDTreeInfo.Directories.GetNumOfDirs())

    _ = targetDMgr.DeleteAll()

    return
  }

  if srcDTreeInfo.FoundFiles.GetNumOfFileMgrs() != targetDTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    t.Errorf("ERROR: Expected %v-files would be copied. Instead, %v-files were copied!",
      srcDTreeInfo.FoundFiles.GetNumOfFileMgrs(), targetDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  }

  expectedNumOfDirsCopied := uint64(srcDTreeInfo.Directories.GetNumOfDirs())
  expectedNumOfDirsCreated := expectedNumOfDirsCopied
  expectedTotalDirsScanned := expectedNumOfDirsCopied

  if expectedTotalDirsScanned != dTreeCopyStats.TotalDirsScanned {
    t.Errorf("Error: Expected dTreeCopyStats.TotalDirsScanned='%v'.\n"+
      "Instead, dTreeCopyStats.TotalDirsScanned='%v'\n",
      expectedTotalDirsScanned, dTreeCopyStats.TotalDirsScanned)
  }

  if expectedNumOfDirsCopied != dTreeCopyStats.DirsCopied {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCopied='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCopied='%v'\n",
      expectedNumOfDirsCopied, dTreeCopyStats.DirsCopied)
  }

  if expectedNumOfDirsCreated != dTreeCopyStats.DirsCreated {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCreated='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCreated='%v'\n",
      expectedNumOfDirsCreated, dTreeCopyStats.DirsCreated)
  }

  expectedNumOfFilesCopied := uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  expectedNumOfFileBytesCopied := srcDTreeInfo.FoundFiles.GetTotalFileBytes()
  expectedNumOfFilesNotCopied := uint64(0)
  expectedNumOfFileBytesNotCopied := uint64(0)

  if expectedNumOfFilesCopied != dTreeCopyStats.FilesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCreated='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCreated='%v'\n",
      expectedNumOfFilesCopied, dTreeCopyStats.DirsCreated)
  }

  if expectedNumOfFileBytesCopied != dTreeCopyStats.FileBytesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesCopied='%v'\n",
      expectedNumOfFileBytesCopied, dTreeCopyStats.DirsCreated)
  }

  if expectedNumOfFilesNotCopied != dTreeCopyStats.FilesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FilesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesNotCopied='%v'\n",
      expectedNumOfFilesNotCopied, dTreeCopyStats.FilesNotCopied)
  }

  if expectedNumOfFileBytesNotCopied != dTreeCopyStats.FileBytesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesNotCopied='%v'\n",
      expectedNumOfFileBytesNotCopied, dTreeCopyStats.FileBytesNotCopied)
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_02(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_02"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  srcDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' is INVALID!\n" +
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  srcDMgr.isInitialized = true

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_03(t *testing.T) {

  srcDir := "../filesfortest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_03"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  targetDMgr.isInitialized = false

  _,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'targetDMgr' is INVALID!\n" +
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  targetDMgr.isInitialized = true

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_04(t *testing.T) {

  srcDir := "../filesfortest/iDoNotExist"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_04"

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
  }

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  _,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) == 0 {
    t.Error("Expected Errors to be returned from " +
      "srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n" +
      "because 'srcDMgr' DOES NOT EXIST !\n" +
      "However - NO ERRORS WERE RETURNED!!!\n")

  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_05(t *testing.T) {

  srcDir := "../logTest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_05"

  fh := FileHelper{}

  _ = fh.DeleteDirPathAll(targetDir)

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    _ = fh.DeleteDirPathAll(targetDir)
    return
  }

  fsc := FileSelectionCriteria{}

  dTreeStats,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)

  if len(errs) > 0 {

    t.Errorf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = targetDMgr.DeleteAll()
    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory path DOES NOT EXIST!!\n")

    return
  }

  fsc = FileSelectionCriteria{}

  srcDTreeInfo, err := srcDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by srcDMgr.FindWalkSubDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  fsc = FileSelectionCriteria{}

  targetDTreeInfo, err := targetDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  srcDirs := srcDTreeInfo.Directories.GetNumOfDirs()

  targetDirs := targetDTreeInfo.Directories.GetNumOfDirs()

  if srcDirs != targetDirs {
    t.Errorf("ERROR: Expected %v-directories would be created.\n"+
      "Instead, %v-directories were created!\n",
      srcDirs, targetDirs)

    _ = targetDMgr.DeleteAll()

    return
  }

  tFileInfo, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
    _ = targetDMgr.DeleteAll()

    return
  }

  if tFileInfo.GetNumOfFileMgrs() > 0 {
    t.Errorf("ERROR: Expected ZERO files in top level target directory.\n"+
      "Instead, the top level target directory had %v-files.\nTarget Directory='%v'\n",
      tFileInfo.GetNumOfFileMgrs(), targetDMgr.GetAbsolutePath())
  }

  expectedDirsCopied := uint64(srcDTreeInfo.Directories.GetNumOfDirs())
  expectedDirsCreated := expectedDirsCopied
  expectedTotalDirsScanned := expectedDirsCopied

  if expectedTotalDirsScanned != dTreeStats.TotalDirsScanned {
    t.Errorf("Error: Expected dTreeCopyStats.TotalDirsScanned='%v'.\n"+
      "Instead, dTreeCopyStats.TotalDirsScanned='%v'\n",
      expectedTotalDirsScanned, dTreeStats.TotalDirsScanned)
  }

  if expectedDirsCopied != dTreeStats.DirsCopied {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCopied='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCopied='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCopied)
  }

  if expectedDirsCreated != dTreeStats.DirsCreated {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCreated='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCreated='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCreated)
  }

  expectedFilesCopied := uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  expectedFileBytesCopied := srcDTreeInfo.FoundFiles.GetTotalFileBytes()
  expectedFilesNotCopied := uint64(0)
  expectedFileBytesNotCopied := uint64(0)
  expectedTotalFilesProcessed := expectedFilesCopied

  if expectedFilesCopied != dTreeStats.FilesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FilesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesCopied='%v'\n",
      expectedFilesCopied, dTreeStats.FilesCopied)
  }

  if expectedFileBytesCopied != dTreeStats.FileBytesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesCopied='%v'\n",
      expectedFileBytesCopied, dTreeStats.FileBytesCopied)
  }

  if expectedFilesNotCopied != dTreeStats.FilesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FilesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesNotCopied='%v'\n",
      expectedFilesNotCopied, dTreeStats.FilesNotCopied)
  }

  if expectedFileBytesNotCopied != dTreeStats.FileBytesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesNotCopied='%v'\n",
      expectedFileBytesNotCopied, dTreeStats.FileBytesNotCopied)
  }

  if expectedTotalFilesProcessed != dTreeStats.TotalFilesProcessed {
    t.Errorf("Error: Expected dTreeCopyStats.TotalFilesProcessed='%v'.\n"+
      "Instead, dTreeCopyStats.TotalFilesProcessed='%v'\n",
      expectedTotalFilesProcessed, dTreeStats.TotalFilesProcessed)
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}

func TestDirMgr_CopySubDirectoryTree_06(t *testing.T) {

  srcDir := "../logTest"

  srcDMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "../dirmgrtests/TestDirMgr_CopySubDirectoryTree_06"

  fh := FileHelper{}

  _ = fh.DeleteDirPathAll(targetDir)

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    _ = fh.DeleteDirPathAll(targetDir)
    return
  }

  fsc := FileSelectionCriteria{}

  dTreeStats,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, false, fsc)

  if len(errs) > 0 {
    t.Errorf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = targetDMgr.DeleteAll()

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: The target directory path DOES NOT EXIST!!\n")

    return
  }

  fsc = FileSelectionCriteria{}

  srcDTreeInfo, err := srcDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by srcDMgr.FindWalkSubDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  fsc = FileSelectionCriteria{}

  targetDTreeInfo, err := targetDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    t.Errorf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  srcDirs := srcDTreeInfo.Directories.GetNumOfDirs()
  srcDirs-- // Discount the one empty subdirectory

  targetDirs := targetDTreeInfo.Directories.GetNumOfDirs()

  if srcDirs != targetDirs {
    t.Errorf("ERROR: Expected %v-directories would be created.\n"+
      "Instead, %v-directories were created!\n",
      srcDirs, targetDirs)

    _ = targetDMgr.DeleteAll()

    return
  }

  tFileInfo, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
    _ = targetDMgr.DeleteAll()

    return
  }

  if tFileInfo.GetNumOfFileMgrs() > 0 {
    t.Errorf("ERROR: Expected ZERO files in top level target directory.\n"+
      "Instead, the top level target directory had %v-files.\nTarget Directory='%v'\n",
      tFileInfo.GetNumOfFileMgrs(), targetDMgr.GetAbsolutePath())
  }

  // Subtract 1 to eliminate the empty directory
  expectedDirsCopied := uint64(srcDTreeInfo.Directories.GetNumOfDirs() - 2)
  expectedDirsCreated := uint64(srcDTreeInfo.Directories.GetNumOfDirs() - 2)
  expectedTotalDirsScanned := uint64(srcDTreeInfo.Directories.GetNumOfDirs())

  if expectedTotalDirsScanned != dTreeStats.TotalDirsScanned {
    t.Errorf("Error: Expected dTreeCopyStats.TotalDirsScanned='%v'.\n"+
      "Instead, dTreeCopyStats.TotalDirsScanned='%v'\n",
      expectedTotalDirsScanned, dTreeStats.TotalDirsScanned)
  }

  if expectedDirsCopied != dTreeStats.DirsCopied {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCopied='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCopied='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCopied)

  }

  if expectedDirsCreated != dTreeStats.DirsCreated {
    t.Errorf("Error: Expected dTreeCopyStats.DirsCreated='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCreated='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCreated)
  }

  expectedFilesCopied := uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  expectedFileBytesCopied := srcDTreeInfo.FoundFiles.GetTotalFileBytes()
  expectedFilesNotCopied := uint64(0)
  expectedFileBytesNotCopied := uint64(0)
  expectedTotalFilesProcessed := expectedFilesCopied

  if expectedFilesCopied != dTreeStats.FilesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FilesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesCopied='%v'\n",
      expectedFilesCopied, dTreeStats.FilesCopied)
  }

  if expectedFileBytesCopied != dTreeStats.FileBytesCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesCopied='%v'\n",
      expectedFileBytesCopied, dTreeStats.FileBytesCopied)
  }

  if expectedFilesNotCopied != dTreeStats.FilesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FilesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesNotCopied='%v'\n",
      expectedFilesNotCopied, dTreeStats.FilesNotCopied)
  }

  if expectedFileBytesNotCopied != dTreeStats.FileBytesNotCopied {
    t.Errorf("Error: Expected dTreeCopyStats.FileBytesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesNotCopied='%v'\n",
      expectedFileBytesNotCopied, dTreeStats.FileBytesNotCopied)
  }

  if expectedTotalFilesProcessed != dTreeStats.TotalFilesProcessed {
    t.Errorf("Error: Expected dTreeCopyStats.TotalFilesProcessed='%v'.\n"+
      "Instead, dTreeCopyStats.TotalFilesProcessed='%v'\n",
      expectedTotalFilesProcessed, dTreeStats.TotalFilesProcessed)
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

  return
}
