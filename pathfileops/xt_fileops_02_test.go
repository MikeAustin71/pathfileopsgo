package pathfileops

import (
  "strings"
  "testing"
)

func TestFileOps_NewByDirMgrFileName_01(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirMgrFileName_01.txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
  err = FileOps{}.NewByDirMgrFileName(sourceDMgr,
                                          sourceFileNameExt,
                                          destDMgr,
                                          destFileNameExt)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByDirMgrFileName(...)\n" +
      "sourceDMgr='%v'\n" +
      "sourceFileNameExt='%v'\n" +
      "destDMgr='%v'\n" +
      "destFileNameExt='%v'\n" +
      "Error='%v'\n",
      sourceDMgr.GetAbsolutePath(),
      sourceFileNameExt,
      destDMgr.GetAbsolutePath(),
      destFileNameExt,
      err.Error())
    return
  }

}

func TestFileOps_NewByDirMgrFileName_02(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirMgrFileName_02.txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  sourceDMgr.isInitialized = false

  _,
  err = FileOps{}.NewByDirMgrFileName(sourceDMgr,
                                          sourceFileNameExt,
                                          destDMgr,
                                          destFileNameExt)

  if err == nil {
    t.Error("ERROR: Expected an error return from FileOps{}.NewByDirMgrFileName(...)\n" +
      "because 'sourceDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }

}

func TestFileOps_NewByDirMgrFileName_03(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirMgrFileName_03.txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDMgr.isInitialized = false

  _,
  err = FileOps{}.NewByDirMgrFileName(sourceDMgr,
                                          sourceFileNameExt,
                                          destDMgr,
                                          destFileNameExt)

  if err == nil {
    t.Error("ERROR: Expected an error return from FileOps{}.NewByDirMgrFileName(...)\n" +
      "because 'destDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }

}

func TestFileOps_NewByDirMgrFileName_04(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirMgrFileName_03.txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  sourceFileNameExt = ""

  _,
  err = FileOps{}.NewByDirMgrFileName(sourceDMgr,
                                          sourceFileNameExt,
                                          destDMgr,
                                          destFileNameExt)

  if err == nil {
    t.Error("ERROR: Expected an error return from FileOps{}.NewByDirMgrFileName(...)\n" +
      "because 'sourceFileNameExt' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }
}

func TestFileOps_NewByDirMgrFileName_05(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_05.txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destFileNameExt = ""

  _,
  err = FileOps{}.NewByDirMgrFileName(sourceDMgr,
                                          sourceFileNameExt,
                                          destDMgr,
                                          destFileNameExt)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByDirMgrFileName()\n" +
      "Error='%v'\n", err.Error())

  }
}

func TestFileOps_NewByDirMgrFileName_06(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "?%**#@!////////.....txt"

  destDMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(destDir)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
    err = FileOps{}.NewByDirMgrFileName(
                      sourceDMgr,
                      sourceFileNameExt,
                      destDMgr,
                      destFileNameExt)

  if err == nil {
    t.Error("Expected an error return from FileOps{}.NewByDirMgrFileName(...)\n" +
      "because 'destFileNameExt' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_01(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_02.txt"

  fOp1, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
                  sourceDir,
                  sourceFileNameExt,
                  destDir,
                  destFileNameExt)

  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  err = fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByHardLinkByIo())

  if err != nil {
    t.Errorf("Error returned by fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByHardLinkByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  fOp2, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
    sourceDir,
    sourceFileNameExt,
    destDir,
    destFileNameExt)

  if err != nil {
    t.Errorf("Error returned by #2 FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  err = fOp2.SetFileOpsCode(FileOpCode.CopySourceToDestinationByHardLinkByIo())

  if err != nil {
    t.Errorf("Error returned by fOp2.SetFileOpsCode(" +
      "FileOpCode.CopySourceToDestinationByHardLinkByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !fOp1.Equal(&fOp2) {
    t.Error("Expected fOp1==fOp2.\n" +
      "Instead, fOp1 is NOT EQUAL to fOp2!\n")
  }

}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_02(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/!@#$$%/////....../level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_02.txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
                  sourceDir,
                  sourceFileNameExt,
                  destDir,
                  destFileNameExt)

  if err == nil {
    t.Error("Expected an error return from FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'sourceDir' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_03(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "!@@#$%/////.......txt"


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_02.txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
                  sourceDir,
                  sourceFileNameExt,
                  destDir,
                  destFileNameExt)

  if err == nil {
    t.Error("Expected an error return from FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'sourceFileNameExt' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_04(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"


  destDir := "../!@#$%//////......./Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_02.txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
                  sourceDir,
                  sourceFileNameExt,
                  destDir,
                  destFileNameExt)

  if err == nil {
    t.Error("Expected an error return from FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'destDir' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_05(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "!@#$///////.......txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
                  sourceDir,
                  sourceFileNameExt,
                  destDir,
                  destFileNameExt)

  if err == nil {
    t.Error("Expected an error return from FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'destFileNameExt' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_06(t *testing.T) {

  // sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceDir := ""
  sourceFileNameExt := "level_2_0_test.txt"


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_06.txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
    sourceDir,
    sourceFileNameExt,
    destDir,
    destFileNameExt)

  if err == nil {
    t.Error("Expected an error return by FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'sourceDir' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_07(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  // sourceFileNameExt := "level_2_0_test.txt"
  sourceFileNameExt := ""


  destDir := "../createFilesTest/Level01/Level02"
  destFileNameExt := "TestFileOps_NewByDirStrsAndFileNameExtStrs_07.txt"

  _, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
    sourceDir,
    sourceFileNameExt,
    destDir,
    destFileNameExt)

  if err == nil {
    t.Error("Expected an error return by FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'sourceFileNameExt' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOps_NewByDirStrsAndFileNameExtStrs_08(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"
  sourceFileNameExt := "level_2_0_test.txt"

  destDir := "../createFilesTest/Level01/Level02"

  destFileNameExt := ""

  expectedAbsDestFile, err := FileHelper{}.MakeAbsolutePath(destDir + "/" + sourceFileNameExt)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}." +
      "MakeAbsolutePath(destDir + \"/\" + sourceFileNameExt)\n" +
      "destDir='%v'\n" +
      "sourceFileNameExt='%v'\n" +
      "Error='%v'\n", destDir, sourceFileNameExt, err.Error())
    return
  }

  expectedAbsDestFile = strings.ToLower(expectedAbsDestFile)

  fOps, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
    sourceDir,
    sourceFileNameExt,
    destDir,
    destFileNameExt)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs(...)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if expectedAbsDestFile != strings.ToLower(fOps.destination.GetAbsolutePathFileName()) {
    t.Errorf("ERROR: Expected destination='%v'\n" +
      "Instead, destination='%v'\n",
      expectedAbsDestFile, strings.ToLower(fOps.destination.GetAbsolutePathFileName()))
  }
}

func TestFileOps_NewByFileMgrs_01(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"
  destFile := "../createFilesTest/Level01/Level02/TestFileOps_NewByFileMgrs_01.txt"

  sourceFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile)\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n",
      sourceFile, err.Error())
    return
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      destFile, err.Error())

    return
  }

  fOp1, err := FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)


  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)\n" +
      "sourceFMgr='%v'\n" +
      "destFileMgr='%v'\n" +
      "Error='%v'\n",
      sourceFMgr.GetAbsolutePath(),
      destFileMgr.GetAbsolutePath(),
      err.Error())
    return
  }

  fOp2, err := FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)

  if err != nil {
    t.Errorf("Error returned by #2 FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)\n" +
      "sourceFMgr='%v'\n" +
      "destFileMgr='%v'\n" +
      "Error='%v'\n",
      sourceFMgr.GetAbsolutePath(),
      destFileMgr.GetAbsolutePath(),
      err.Error())
    return
  }

  if !fOp1.Equal(&fOp2) {
    t.Error("ERROR: fOp1 is NOT EQUAL to fOp2!\n")
  }

}

func TestFileOps_NewByFileMgrs_02(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"
  destFile := "../createFilesTest/Level01/Level02/TestFileOps_NewByFileMgrs_01.txt"

  sourceFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile)\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n",
      sourceFile, err.Error())
    return
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      destFile, err.Error())
    return
  }

  sourceFMgr.isInitialized = false

  _, err = FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)

  if err == nil {
    t.Error("ERROR: Expected an error return from FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)\n" +
      "because 'sourceFMgr' is invalid!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

func TestFileOps_NewByFileMgrs_03(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"
  destFile := "../createFilesTest/Level01/Level02/TestFileOps_NewByFileMgrs_01.txt"

  sourceFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile)\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n",
      sourceFile, err.Error())
    return
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      destFile, err.Error())
    return
  }

  destFileMgr.isInitialized = false

  _, err = FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)

  if err == nil {
    t.Error("ERROR: Expected an error return from FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)\n" +
      "because 'sourceFMgr' is invalid!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOps_NewByPathFileNameExtStrs_01(t *testing.T) {

  sourceFile :=
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"

  destFile :=
    "../createFilesTest/Level01/Level02/TestFileOps_NewByPathFileNameExtStrs_01.txt"

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())

    return
  }

  fOp2, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by #2 FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())

    return
  }

  err = fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())

  if err != nil {
    t.Errorf("Error returned by fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  err = fOp2.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())

  if err != nil {
    t.Errorf("Error returned by fOp2.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !fOp1.Equal(&fOp2) {
    t.Error("Expected that fOp1==fOp2.\n" +
      "Instead, fOp1 is NOT EQUAL to fOp2\n")
  }

}

func TestFileOps_NewByPathFileNameExtStrs_02(t *testing.T) {

  sourceFile :=
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/!/////.......txt"

  destFile :=
    "../createFilesTest/Level01/Level02/TestFileOps_NewByPathFileNameExtStrs_02.txt"

  _, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "because 'sourceFile' is invalid!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOps_NewByPathFileNameExtStrs_03(t *testing.T) {

  sourceFile :=
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"

  destFile :=
    "../createFilesTest/Level01/Level02/?%*!//////.........txt"

  _, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "because 'destFile' is invalid!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOps_SetFileOpsCode_01(t *testing.T) {

  sourceFile :=
    "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"

  destFile :=
    "../createFilesTest/Level01/Level02/TestFileOps_NewByPathFileNameExtStrs_01.txt"

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())

    return
  }

  fOpCode := FileOperationCode(-99)

  err = fOp1.SetFileOpsCode(fOpCode)

  if err == nil {
    t.Error("Expected an error return by fOp1.SetFileOpsCode(fOpCode)\n" +
      "because 'fOpCode' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}
