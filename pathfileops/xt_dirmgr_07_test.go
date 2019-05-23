package pathfileops


import (
  fp "path/filepath"
  "strings"
  "testing"
)

func TestDirMgr_MakeDir_01(t *testing.T) {

  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../checkfiles/checkfiles99/checkfiles999")

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
  }

  if dMgr.doesAbsolutePathExist {

    err = dMgr.DeleteAll()

    if err != nil {
      t.Errorf("%v", err.Error())
    }

  }

  if dMgr.doesAbsolutePathExist {
    t.Errorf("Error: Attempted to delete dMgr.absolutePath='%v'. Deletion Attempt FAILED. This directory still exists.", dMgr.absolutePath)
  }

  err = dMgr.MakeDir()

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Error: Attempted to create dMgr.absolutePath='%v'. Creation Attempt FAILED. This directory does NOT exist.", dMgr.absolutePath)
  }

}

func TestDirMgr_MoveDirectory_01(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_01"

  srcDir := baseDir +  "/source"

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n" +
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n" +
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr." +
        "CopyDirectory(srcDirMgr, fsc)\n" +
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
      "Instead, %v-files were found.\nSource Dir='%v'\n" ,
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  errs = srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n" +
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

    _ = fh.DeleteDirPathAll(baseDir)
    return
  }

  fileNames := []string{"level_0_0_test.txt",
    "level_0_1_test.txt",
    "level_0_2_test.txt",
    "level_0_3_test.txt",
    "level_0_4_test.txt" }


  fsc = FileSelectionCriteria{}

  fMgrCollection, err = targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 5 {
    t.Errorf("Test Setup Error: Expected to find 5-files in 'targetDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(baseDir)

    return

  }

  for i:=0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fMgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n" +
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(baseDir)

      return
    }

    fileName := fMgr.GetFileNameExt()
    foundFile := false
    for k:=0;k < len(fileNames); k++ {
      if fileNames[k] == fileName {
        foundFile = true
      }
    }

    if foundFile == false {
      t.Errorf("Error: File NOT Found. Expected to find specfic file Name.\n" +
        "However, it WAS NOT FOUND!\nFileName='%v'", fileName )
    }

  }

  fsc = FileSelectionCriteria{}

  fMgrCollection, err = srcDirMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by srcDirMgr.FindFilesBySelectCriteria(fsc).\n"+
      "srcDirMgr='%v'\nError='%v'\n", srcDirMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 0 {
    t.Errorf("Test Setup Error: Expected to find ZERO files in 'sourceDir'.\n"+
      "Instead, %v-files were found.\n" +
      "Source Dir='%v'", fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())
  }

  err = fh.DeleteDirPathAll(baseDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(baseDir)\baseDir='%v'\n"+
      "Error='%v'\n", baseDir, err.Error())
  }

  return
}

func TestDirMgr_MoveDirectory_02(t *testing.T) {

  baseDir := "../checkfiles/TestDirMgr_MoveFilesToDirectory_02"

  srcDir := baseDir +  "/source"

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n" +
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }


  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n" +
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr." +
        "CopyDirectory(srcDirMgr, fsc)\n" +
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
      "Instead, %v-files were found.\nSource Dir='%v'\n" ,
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  srcDirMgr.isInitialized = false

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

  srcDir := baseDir +  "/source"

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n" +
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }


  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n" +
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr." +
        "CopyDirectory(srcDirMgr, fsc)\n" +
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
      "Instead, %v-files were found.\nSource Dir='%v'\n" ,
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(baseDir)

    return
  }

  targetDMgr.isInitialized = false

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n" +
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

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

  srcDir :=   "../checkfiles/sourceTestDirMgr_MoveFilesToDirectory_05"

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
    t.Errorf("Test Setup Error returned from DirMgr{}.New(targetDMgr).\n" +
      "targetDMgr='%v'\nError='%v'\n", targetDMgr, err.Error())
    return
  }

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(srcDir).\n" +
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  origSrcDir := "../filesfortest/levelfilesfortest"

  origSrcDMgr, err := DirMgr{}.New(origSrcDir)

  if err != nil {
    t.Errorf("Test Setup Error returned from DirMgr{}.New(origSrcDir).\n" +
      "origSrcDir='%v'\nError='%v'\n", origSrcDir, err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  errs := origSrcDMgr.CopyDirectory(srcDirMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Test Setup Error returned from origSrcDMgr." +
        "CopyDirectory(srcDirMgr, fsc)\n" +
        "srcDirMgr='%v'\nError='%v'\n\n",
        targetDMgr.GetAbsolutePath(), errs[i].Error())
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
      "Instead, %v-files were found.\nSource Dir='%v'\n" ,
      fMgrCollection.GetNumOfFileMgrs(), srcDirMgr.GetAbsolutePath())

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  fsc.FileNamePatterns = []string{"*.htm"}

  errs = srcDirMgr.MoveDirectory(targetDMgr, fsc)

  if len(errs) > 0 {
    for i:=0; i < len(errs); i++ {
      t.Errorf("Error returned from srcDirMgr.MoveDirectory(targetDMgr, fsc)\n" +
        "targetDir='%v'\nError='%v'\n\n", targetDMgr.GetAbsolutePath(), errs[0].Error())
    }

    _ = fh.DeleteDirPathAll(srcDir)
    _ = fh.DeleteDirPathAll(targetDir)

    return
  }

  if targetDMgr.DoesAbsolutePathExist() {
    t.Error("Expected that 'targetDir' would NOT exist because no files\n" +
      "were selected for the 'move' operation.\nHowever, the 'targetDir' DOES EXIST!!! ERROR!!!!\n")
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

func TestDirMgr_New_01(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.", true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.", true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.", expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.", expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.", true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_02(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("./")
  expectedPath := fh.AdjustPathSlash(".")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.", true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.", true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.", true, dMgr.isParentPathPopulated)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.", true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_03(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../")
  expectedPath := fh.AdjustPathSlash("..")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'",
      origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      true, dMgr.isParentPathPopulated)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_04(t *testing.T) {
  fh := FileHelper{}
  xDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
  }

  volName := fp.VolumeName(xDir)

  origDir := fh.AdjustPathSlash(volName)
  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir := origDir

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      false, dMgr.isParentPathPopulated)
  }

  if false != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      false, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_05(t *testing.T) {
  fh := FileHelper{}
  origDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
  }

  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      true, dMgr.isParentPathPopulated)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_06(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/test2007.txt")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_07(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, "+
      "isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_08(t *testing.T) {
  fh := FileHelper{}
  rawDir := "D:/"
  expectedPath := fh.AdjustPathSlash("D:")
  expectedAbsDir := expectedPath

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := ""

  expectedRelativePath := ""

  dMgr, err := DirMgr{}.New(rawDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDir). rawDir=='%v' Error='%v'", rawDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. "+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      false, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if false != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      false, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_09(t *testing.T) {
  fh := FileHelper{}
  rawPath := "../filesfortest/newfilesfortest"
  expectedPath := fh.AdjustPathSlash(rawPath)
  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expectedPath). expectedPath=='%v'  Error='%v'", expectedPath, err.Error())
  }

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsPath)

  expectedVolumeName := fp.VolumeName(expectedAbsPath)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsPath, fh.AdjustPathSlash("/newfilesfortest"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  expectedRelativePath := "newfilesfortest"

  dMgr, err := DirMgr{}.New(rawPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawPath). rawPath=='%v' Error='%v'",
      rawPath, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsPath != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsPath, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. "+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_SubstituteBaseDir_01(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01")

  substitutePath := fh.AdjustPathSlash("../checkfiles")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).  Error='%v'",
      err.Error())
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.  Instead substituted "+
      "path = '%v' ",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.  Instead "+
      "substituted absolute path = '%v' ",
      expectedAbsPath, dMgrResult.absolutePath)
  }

}

func TestDirMgr_SubstituteBaseDir_02(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute).  Error='%v'", err.Error())
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.  Instead substituted path = '%v' ",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.  Instead substituted absolute path = '%v' ",
      expectedAbsPath, dMgrResult.absolutePath)
  }

}


