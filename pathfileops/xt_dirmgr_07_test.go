package pathfileops

import (
  "testing"
)

func TestDirMgr_MakeDir_01(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../checkfiles/checkfiles99/TestDirMgr_MakeDir_01")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir='%v'\nError='%v'\n",
      origDir, err.Error())
    return
  }

  if dMgr.doesAbsolutePathExist {

    err = dMgr.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by dMgr.DeleteAll()\n"+
        "Error='%v'", err.Error())
      return
    }

  }

  if dMgr.doesAbsolutePathExist {
    t.Errorf("Error: Attempted to delete dMgr.absolutePath='%v'.\n"+
      "Deletion Attempt FAILED. This directory still exists.\n",
      dMgr.absolutePath)

    return
  }

  err = dMgr.MakeDir()

  if err != nil {
    t.Errorf("Error returned by dMgr.MakeDir()\n"+
      "Error='%v'", err.Error())
    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Error: Attempted to create dMgr.absolutePath='%v'.\n"+
      "Creation Attempt FAILED. This directory does NOT exist.\n",
      dMgr.absolutePath)
  }

  err = dMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by dMgr.DeleteAll().\n"+
      "Error='%v'\n", err.Error())
  }

}

func TestDirMgr_MakeDir_02(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../checkfiles/checkfiles99/TestDirMgr_MakeDir_02")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir='%v'\nError='%v'\n",
      origDir, err.Error())

    return
  }

  if dMgr.doesAbsolutePathExist {

    err = dMgr.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by dMgr.DeleteAll()\n"+
        "Error=%v", err.Error())
    }

    return
  }

  if dMgr.doesAbsolutePathExist {
    t.Errorf("Error: Attempted to delete dMgr.absolutePath.\n"+
      "Deletion Attempt FAILED. This directory still exists.\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)

    return
  }

  dMgr.isInitialized = false

  err = dMgr.MakeDir()

  if err == nil {
    t.Error("Expected an error to be returned by dMgr.MakeDir()\n" +
      "because dMgr is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
    return
  }

  dMgr.isInitialized = true

  err = dMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by dMgr.DeleteAll().\n"+
      "Error='%v'\n", err.Error())
  }
}

func TestDirMgr_MakeDirWithPermission_01(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_MakeDirWithPermission_01"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())
  }

  err = targetDMgr.MakeDirWithPermission(fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Errorf("ERROR: Target Directory Path DOES NOT EXIST!\n"+
      "Target Directory='%v'\n", targetDMgr.GetAbsolutePath())
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_MakeDirWithPermission_02(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_MakeDirWithPermission_02"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())
  }

  targetDMgr.isInitialized = false

  err = targetDMgr.MakeDirWithPermission(fPermCfg)

  if err == nil {
    t.Error("Expected an error return from targetDMgr.MakeDirWithPermission(fPermCfg)\n" +
      "because targetDMgr is INVALID!\nHowever NO ERROR WAS RETURNED!!!\n")
  }

  targetDMgr.isInitialized = true

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }
}

func TestDirMgr_MakeDirWithPermission_03(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_MakeDirWithPermission_03"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fPermCfg.isInitialized = false

  err = targetDMgr.MakeDirWithPermission(fPermCfg)

  if err == nil {
    t.Error("Expected an error return from targetDMgr.MakeDirWithPermission(fPermCfg)\n" +
      "because fPermCfg is INVALID!\nHowever NO ERROR WAS RETURNED!!!\n")
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }
}

func TestDirMgr_MakeDirWithPermission_04(t *testing.T) {

  targetDir := "../checkfiles/TestDirMgr_MakeDirWithPermission_04"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  err = targetDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.MakeDir().\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())
  }

  err = targetDMgr.MakeDirWithPermission(fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\").\n"+
      "Error='%v'\n", err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    t.Errorf("ERROR: Target Directory Path DOES NOT EXIST!\n"+
      "Target Directory='%v'\n", targetDMgr.GetAbsolutePath())
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
      "Target Directory='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
  }

}

func TestDirMgr_MoveDirectory_01(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_01"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fileNames := []string{"level_0_0_test.txt",
    "level_0_1_test.txt",
    "level_0_2_test.txt",
    "level_0_3_test.txt",
    "level_0_4_test.txt"}

  expectedNumOfSrcDirFiles := len(fileNames)

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

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != expectedNumOfSrcDirFiles {
    t.Errorf("Test Setup Error: Expected to find %v-files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\nSource Dir='%v'\n",
      expectedNumOfSrcDirFiles,
      fMgrCollection.GetNumOfFileMgrs(),
      srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  numOfSrcFilesMoved,
    numOfSrcFilesRemaining,
    numOfSubDirectories,
    dMgrDirWasDeleted,
    errs := srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n"+
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)
    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err = targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != expectedNumOfSrcDirFiles {
    t.Errorf("Test Setup Error: Expected to find %v-files in 'targetDir'.\n"+
      "Instead, %v-files were found.",
      expectedNumOfSrcDirFiles,
      fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseDir)

    return

  }

  for i := 0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(baseDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k := 0; k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n"+
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName)
    }

  }

  if srcDirMgr.DoesAbsolutePathExist() {
    t.Errorf("Error: That 'sourceDir' would be deleted since all the files were mvoed.\n"+
      "Instead, the source directory still exists.\n"+
      "Source Dir='%v'", srcDirMgr.GetAbsolutePath())
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  if numOfSrcFilesRemaining != 0 {
    t.Errorf("Error: Expected that number of source files remaining would be zero!\n"+
      "Instead, number of source files remaining='%v'\n", numOfSrcFilesRemaining)
  }

  if numOfSrcFilesMoved != expectedNumOfSrcDirFiles {
    t.Errorf("Error: Expected that the number of source files moved\n"+
      "to target directory would equal '%v'.\n"+
      "Instead, the number of source files moved='%v'\n",
      expectedNumOfSrcDirFiles, numOfSrcFilesMoved)
  }

  if numOfSubDirectories != 0 {
    t.Errorf("Error: Expected that the number of sub-directories counted would equal zero.\n"+
      "Instead, the number of sub-directories='%v' .\n",
      numOfSubDirectories)
  }

  if !dMgrDirWasDeleted {
    t.Error("Error: Expected dMgrDirWasDeleted='true'.\n" +
      "Instead, dMgrDirWasDeleted='false'!\n")
  }

  return
}

func TestDirMgr_MoveDirectory_02(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_02"

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

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Test Setup Error: Expected to find 5-files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\nSource Dir='%v'\n",
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

  _,
    _,
    _,
    _,
    errs = srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n" +
      "because srcDirMgr is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectory_03(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_03"

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

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        targetDMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Test Setup Error: Expected to find 5-files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\nSource Dir='%v'\n",
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

  _,
    _,
    _,
    _,
    errs = srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n" +
      "because targetDMgr is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectory_04(t *testing.T) {

  srcDir := "../checkfiles/MoveFilesToDirectory_04"

  targetDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_04"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
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

  fsc := FileSelectionCriteria{}

  _,
    _,
    _,
    _,
    errs := srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) == 0 {
    t.Error("Expected an error return from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n" +
      "because srcDirMgr is DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

  err = fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectory_05(t *testing.T) {

  srcDir := "../checkfiles/sourceTestDirMgr_MoveFilesToDirectory_05"

  targetDir := "../checkfiles/targetTestDirMgr_MoveFilesToDirectory_05"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
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

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Test Setup Error: Expected to find 5-files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\nSource Dir='%v'\n",
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  fsc.FileNamePatterns = []string{"*.htm"}

  _,
    _,
    _,
    _,
    errs = srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n"+
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if targetDMgr.DoesAbsolutePathExist() {
    t.Error("ERROR: Expected that 'targetDir' would NOT exist because no files\n" +
      "were selected for the 'move' operation.\nHowever, the 'targetDir' DOES EXIST!!! ERROR!!!!\n")
  }

  if !srcDirMgr.DoesAbsolutePathExist() {
    t.Errorf("ERROR: Expected that the source directory would NOT be deleted.\n"+
      "However, the source directory has been delted and DOES NOT EXIST!\n"+
      "Source Directory='%v'\n", srcDirMgr.GetAbsolutePath())
  }

  err = fh.DeleteDirPathAll(srcDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(srcDir)\nsrcDir='%v'\n"+
      "Error='%v'\n", srcDir, err.Error())
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(targetDir)\ntargetDir='%v'\n"+
      "Error='%v'\n", targetDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectory_06(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_06"

  srcDir := baseDir + "/source"

  targetDir := baseDir + "/target"

  fileNames := []string{"level_0_0_test.txt",
    "level_0_1_test.txt",
    "level_0_2_test.txt",
    "level_0_3_test.txt",
    "level_0_4_test.txt"}

  expectedNumOfSrcDirFilesMoved := len(fileNames)
  expectedTotalNumOfSrcDirFiles := 8
  expectedNumOfSrcDirFilesRemaining := 3

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

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        targetDMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  origSrcDir = "../filesfortest/htmlFilesForTest"

  origSrcDMgr, err = DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Html File Setup Error returned from DirMgr{}.New(origSrcDir2).\n"+
      "origSrcDir2='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc = FileSelectionCriteria{}

  errs = origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr2 'html' files."+
        "CopyDirectory(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err := srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != expectedTotalNumOfSrcDirFiles {
    t.Errorf("Test Setup Error: Expected to find %v-files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\nSource Dir='%v'\n",
      expectedTotalNumOfSrcDirFiles,
      fMgrCollection.GetNumOfFileMgrs(),
      srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  fsc = FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  numOfSrcFilesMoved,
    numOfSrcFilesRemaining,
    numOfSubDirectories,
    dMgrDirWasDeleted,
    errs := srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n"+
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)
    return
  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err = targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != expectedNumOfSrcDirFilesMoved {
    t.Errorf("Test Setup Error: Expected to find %v-files in 'targetDir'.\n"+
      "Instead, %v-files were found.",
      expectedNumOfSrcDirFilesMoved,
      fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseDir)

    return

  }

  for i := 0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n"+
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(baseDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k := 0; k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n"+
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName)
    }

  }

  if !srcDirMgr.DoesAbsolutePathExist() {
    t.Errorf("Error: Expected that 'sourceDir' would be still exist since not all\n"+
      "the files were moved.\nInstead, the source directory DOES NOT EXIST.\n"+
      "Source Dir='%v'", srcDirMgr.GetAbsolutePath())
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }
  if numOfSrcFilesRemaining != expectedNumOfSrcDirFilesRemaining {
    t.Errorf("Error: Expected that number of source files remaining would be %v.\n"+
      "Instead, number of source files remaining='%v'\n",
      expectedNumOfSrcDirFilesRemaining,
      numOfSrcFilesRemaining)
  }

  if numOfSrcFilesMoved != expectedNumOfSrcDirFilesMoved {
    t.Errorf("Error: Expected that the number of source files moved\n"+
      "to target directory would equal '%v'.\n"+
      "Instead, the number of source files moved='%v'\n",
      expectedNumOfSrcDirFilesMoved, numOfSrcFilesMoved)
  }

  if numOfSubDirectories != 0 {
    t.Errorf("Error: Expected that the number of sub-directories counted would equal zero.\n"+
      "Instead, the number of sub-directories='%v' .\n",
      numOfSubDirectories)
  }

  if dMgrDirWasDeleted {
    t.Error("Error: Expected dMgrDirWasDeleted='false'.\n" +
      "Instead, dMgrDirWasDeleted='true'!\n")
  }

  return
}

func TestDirMgr_MoveDirectoryTree_01(t *testing.T) {

  baseDir := "../dirmgrtests/TestDirMgr_MoveDirectoryTree_01"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

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

  errs = srcDirMgr.MoveDirectoryTree(targetDMgr)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectoryTree(targetDMgr)\n"+
        "srcDirMgr='%v'\ntargetDir='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

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

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectoryTree_02(t *testing.T) {

  baseDir := "../dirmgrtests/TestDirMgr_MoveDirectoryTree_02"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

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

  baseDir := "../dirmgrtests/TestDirMgr_MoveDirectoryTree_03"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

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

  baseDir := "../dirmgrtests/TestDirMgr_MoveDirectoryTree_04"

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

  baseDir := "../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_01"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

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

  errs = srcDirMgr.MoveSubDirectoryTree(targetDMgr)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveSubDirectoryTree(targetDMgr)\n"+
        "srcDirMgr='%v'\ntargetDir='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

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
  origDtreeFiles -= 4

  if origDtreeFiles != targetDtreeInfo.FoundFiles.GetNumOfFileMgrs() {
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

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveSubDirectoryTree_02(t *testing.T) {

  baseDir := "../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_02"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

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

  baseDir := "../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_03"

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

  origSrcDir := "../logTest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n"+
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectoryTree(srcDirMgr, true, fsc)

  if len(errs) > 0 {
    for i := 0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr."+
        "CopyDirectoryTree(srcDirMgr, fsc)\n"+
        "srcDirMgr='%v'\nError='%v'\n\n",
        srcDirMgr.GetAbsolutePath(), errs[i].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

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

  baseDir := "../dirmgrtests/TestDirMgr_MoveSubDirectoryTree_03"

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
