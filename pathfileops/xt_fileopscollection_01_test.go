package pathfileops

import "testing"

func TestFileOpsCollection_AddByDirStrsAndFileNameExtStrs_01(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  sf := make([]string, 5, 10)

  sf[0] = "level_0_0_test.txt"
  sf[1] = "level_0_1_test.txt"
  sf[2] = "level_0_2_test.txt"
  sf[3] = "level_0_3_test.txt"
  sf[4] = "level_0_4_test.txt"

  destDir := "../dirmgrtests"
  df := make([]string, 5, 10)

  df[0] = "level_0_0_test.txt"
  df[1] = "level_0_1_test.txt"
  df[2] = "level_0_2_test.txt"
  df[3] = "level_0_3_test.txt"
  df[4] = "level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], destDir, df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], "+
        "destDir, df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    fullSrcPathFileName := fh.JoinPathsAdjustSeparators(srcDir, sf[i])

    sf[i], err = fh.GetAbsPathFromFilePath(fullSrcPathFileName)

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(fullSrcPathFileName). "+
        "i='%v', fullSrcPathFileName='%v' Error='%v' ", i, fullSrcPathFileName, err.Error())
      return
    }

    fullDestPathFileName := fh.JoinPathsAdjustSeparators(destDir, df[i])

    df[i], err = fh.GetAbsPathFromFilePath(fullDestPathFileName)

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(fullDestPathFileName). "+
        "i='%v', fullDestPathFileName='%v' Error='%v' ", i, fullDestPathFileName, err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  for j := 0; j < arrayLen; j++ {

    fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    srcFMgr := fOps.GetSource()

    if sf[j] != srcFMgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected source file manager[j] = '%v'. "+
        "Instead, source file manager[j]='%v'. j='%v' ",
        sf[j], srcFMgr.GetAbsolutePathFileName(), j)
    }

    destFMgr := fOps.GetDestination()

    if df[j] != destFMgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected destination file manager[j] = '%v'. "+
        "Instead, destination file manager[j]='%v'. j='%v' ",
        df[j], destFMgr.GetAbsolutePathFileName(), j)
    }

  }

}

func TestFileOpsCollection_AddByDirStrsAndFileNameExtStrs_02(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcFile := "level_0_3_test.txt"

  destDir := "../dirmgrtests"

  destFile := "level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  fOpsCol.fileOps = nil

  err :=
    fOpsCol.AddByDirStrsAndFileNameExtStrs(
      srcDir,
      srcFile,
      destDir,
      destFile)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByDirStrsAndFileNameExtStrs(...)\n" +
      "Error='%v'\n", err.Error())
    return
  }

}

func TestFileOpsCollection_AddByDirStrsAndFileNameExtStrs_03(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcFile := "level_0_3_test.txt"

  destDir := "../dirmgrtests"

  destFile := "level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcDir = "!@#$%^&*//////.....txt"

  err :=
    fOpsCol.AddByDirStrsAndFileNameExtStrs(
      srcDir,
      srcFile,
      destDir,
      destFile)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.AddByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'srcDir' is INVALID!\n")
    return
  }

}

func TestFileOpsCollection_AddByDirStrsAndFileNameExtStrs_04(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcFile := "level_0_3_test.txt"

  destDir := "../dirmgrtests"

  destFile := "level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcFile = "!@#$%^&*//////.....txt"

  err :=
    fOpsCol.AddByDirStrsAndFileNameExtStrs(
      srcDir,
      srcFile,
      destDir,
      destFile)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.AddByDirStrsAndFileNameExtStrs(...)\n" +
      "because 'srcFile' is INVALID!\n")
    return
  }

}

func TestFileOpsCollection_AddByDirMgrFileName_01(t *testing.T) {

  srcDir := "../filesfortest/levelfilesfortest/"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  sf := make([]string, 5, 10)

  sf[0] = "level_0_0_test.txt"
  sf[1] = "level_0_1_test.txt"
  sf[2] = "level_0_2_test.txt"
  sf[3] = "level_0_3_test.txt"
  sf[4] = "level_0_4_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  df := make([]string, 5, 10)

  df[0] = "level_0_0_test.txt"
  df[1] = "level_0_1_test.txt"
  df[2] = "level_0_2_test.txt"
  df[3] = "level_0_3_test.txt"
  df[4] = "level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {
    err := fOpsCol.AddByDirMgrFileName(
      srcDirMgr.CopyOut(),
      sf[i],
      destDirMgr.CopyOut(),
      df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByDirMgrFileName(...). "+
        "i='%v' Error='%v'  ", i, err.Error())
    }

    sf[i] = fh.JoinPathsAdjustSeparators(srcDirMgr.GetAbsolutePath(), sf[i])

    df[i] = fh.JoinPathsAdjustSeparators(destDirMgr.GetAbsolutePath(), df[i])

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected file fOpsCol array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  for j := 0; j < arrayLen; j++ {

    fileOps, err := fOpsCol.PeekFileOpsAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    srcFileMgr := fileOps.GetSource()

    if sf[j] != srcFileMgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected source file manager[j]='%v'. "+
        "Instead, source file manger[j]='%v' ",
        sf[j], srcFileMgr.GetAbsolutePathFileName())
    }

    destFileMgr := fileOps.GetDestination()

    if df[j] != destFileMgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected destination file manager[j]='%v'. "+
        "Instead, destination file manger[j]='%v' ",
        df[j], destFileMgr.GetAbsolutePathFileName())
    }

  }
}

func TestFileOpsCollection_AddByDirMgrFileName_02(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest/"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  fOpsCol.fileOps = nil

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr.CopyOut(),
    srcFile,
    destDirMgr.CopyOut(),
    destFile)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByDirMgrFileName(...). "+
      "Error='%v'\n", err.Error())
    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_03(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest/"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcDirMgr.isInitialized = false

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err == nil {
    t.Error("Expected an error return from " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "because 'srcDirMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_04(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest/"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  destDirMgr.isInitialized = false

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err == nil {
    t.Error("Expected an error return from " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "because 'destDirMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_05(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcFile = "!#$%^&///////........txt"

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err == nil {
    t.Error("Expected an error return from " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "because 'srcFile' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_06(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcFile = ""

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err == nil {
    t.Error("Expected an error return from " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "because 'srcFile' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_07(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  destFile = ""

  // No error expected. Source File Name will be
  // substituted for empty destFile string.
  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err != nil {
    t.Errorf("Error returned by " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "Error='%v'\n", err.Error())

    return
  }
}

func TestFileOpsCollection_AddByDirMgrFileName_08(t *testing.T) {

  srcDir := "../checkfiles/levelfilesfortest"

  srcDirMgr, err := DirMgr{}.New(srcDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
      "srcDir='%v' Error='%v' ", srcDir, err.Error())
  }

  srcFile := "level_0_0_test.txt"

  destDir := "../dirmgrtests"

  destDirMgr, err := DirMgr{}.New(destDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(destDir). "+
      "destDir='%v' Error='%v' ", destDir, err.Error())
    return
  }

  destFile := "level_0_0_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  destFile = "!@#$%^///////......txt"

  err = fOpsCol.AddByDirMgrFileName(
    srcDirMgr,
    srcFile,
    destDirMgr,
    destFile)

  if err == nil {
    t.Error("Expected an error return from " +
      "fOpsCol.AddByDirMgrFileName(...) \n"+
      "because 'destFile' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByFileOps_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"


  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}

  for i:=0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[%v], df[%v])\n" +
        "sf[%v]='%v'\ndf[%v]='%v'\nError='%v'\n",
        i, i, i, sf[i], i, df[i], err.Error())
      return
    }

    err = fOp.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())

    if err != nil {
      t.Errorf("Error returned by fOp.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
        "Error='%v'\n", err.Error())
      return
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp)\n" +
        "Error='%v'\n", err.Error())
      return
    }
  }

  numOfFOps := fOpsCol.GetNumOfFileOps()

  if numOfFOps != 5 {
    t.Errorf("ERROR: Expected the number of File Ops='5'.\n" +
      "Instead, File Ops='%v'\n",
      numOfFOps)
  }
}

func TestFileOpsCollection_AddByFileOps_02(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destFile := "../dirmgrtests/level_0_0_test.txt"

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\ndestFile='%v'\nError='%v'\n",
      sourceFile, destFile, err.Error())
    return
  }

  fOpsCol := FileOpsCollection{}

  err = fOpsCol.AddByFileOps(fOp)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp)\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestFileOpsCollection_AddByFileOps_03(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destFile := "../dirmgrtests/level_0_0_test.txt"

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\ndestFile='%v'\nError='%v'\n",
      sourceFile, destFile, err.Error())
    return
  }

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  err = fOpsCol.AddByFileOps(fOp)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp)\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestFileOpsCollection_AddByFileOps_04(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destFile := "../dirmgrtests/level_0_0_test.txt"

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\ndestFile='%v'\nError='%v'\n",
      sourceFile, destFile, err.Error())
    return
  }

  fOpsCol := FileOpsCollection{}

  fOp.isInitialized = false

  err = fOpsCol.AddByFileOps(fOp)

  if err == nil {
    t.Error("Expected an error returned from fOpsCol.AddByFileOps(fOp)\n" +
      "because 'fOp' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByFileMgrs_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fileOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    srcFileMgr, err :=
      FileMgr{}.NewFromPathFileNameExtStr(sf[i])

    if err != nil {
      t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(sf[i]). "+
        "sf[i]='%v' Error='%v' ", sf[i], err.Error())
      return
    }

    destFileMgr, err :=
      FileMgr{}.NewFromPathFileNameExtStr(df[i])

    if err != nil {
      t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(df[i]). "+
        "df[i]='%v' Error='%v' ", df[i], err.Error())
      return
    }

    err = fileOpsCol.AddByFileMgrs(srcFileMgr, destFileMgr)

    if err != nil {
      t.Errorf("Error returned by fileOpsCol.AddByFileMgrs(srcFileMgr, destFileMgr). "+
        "srcFileMgr='%v' destFileMgr='%v' Error='%v' ",
        srcFileMgr.GetAbsolutePathFileName(), destFileMgr.GetAbsolutePathFileName(),
        err.Error())
    }

    sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
        "i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
      return
    }

    df[i], err = fh.GetAbsPathFromFilePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  arrayLen := fileOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Expected final FileOpsCol array length='5'. "+
      "Instead, final array length='%v' ", arrayLen)
  }

  for j := 0; j < arrayLen; j++ {

    fileOps, err := fileOpsCol.PeekFileOpsAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fileOpsCol.PeekFileOpsAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    srcFileMgr := fileOps.GetSource()

    if sf[j] != srcFileMgr.GetAbsolutePathFileName() {
      t.Errorf("Expected srcFileMgr['%v']='%v'. "+
        "Instead, srcFileMgr='%v' ", j, sf[j],
        srcFileMgr.GetAbsolutePathFileName())
    }

    destFileMgr := fileOps.GetDestination()

    if df[j] != destFileMgr.GetAbsolutePathFileName() {
      t.Errorf("Expected destFileMgr['%v']='%v'. "+
        "Instead, destFileMgr='%v' ", j, df[j],
        destFileMgr.GetAbsolutePathFileName())
    }
  }
}

func TestFileOpsCollection_AddByFileMgrs_02(t *testing.T) {

  srcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  srcFMgr, err := FileMgr{}.New(srcFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(srcFile)\n" +
      "srcFile='%v'\nError='%v'\n", srcFile, err.Error())
    return
  }

  destFile := "../dirmgrtests/level_0_0_test.txt"

  destFMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  fOps := FileOpsCollection{}

  fOps.fileOps = nil

  err = fOps.AddByFileMgrs(srcFMgr, destFMgr)

  if err != nil {
    t.Errorf("Error returned by fOps.AddByFileMgrs(srcFMgr, destFMgr)\n" +
      "srcFMgr='%v'\ndestFMgr='%v'\nError='%v'\n",
      srcFMgr.GetAbsolutePathFileName(), destFMgr.GetAbsolutePathFileName(), err.Error())
    return
  }
}

func TestFileOpsCollection_AddByFileMgrs_03(t *testing.T) {

  srcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  srcFMgr, err := FileMgr{}.New(srcFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(srcFile)\n" +
      "srcFile='%v'\nError='%v'\n", srcFile, err.Error())
    return
  }

  destFile := "../dirmgrtests/level_0_0_test.txt"

  destFMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  fOps := FileOpsCollection{}

  srcFMgr.isInitialized = false

  err = fOps.AddByFileMgrs(srcFMgr, destFMgr)

  if err == nil {
    t.Error("Expected an error return from fOps.AddByFileMgrs(srcFMgr, destFMgr)\n" +
      "because srcFMgr is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }
}

func TestFileOpsCollection_AddByPathFileNameExtStrs_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

    sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
        "i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
      return
    }

    df[i], err = fh.GetAbsPathFromFilePath(df[i])

    if err != nil {
      t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
        "i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  for j := 0; j < arrayLen; j++ {

    fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
        "j='%v' Error='%v' ", j, err.Error())
      return
    }

    srcFmgr := fOps.GetSource()

    if sf[j] != srcFmgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected source file manager[j] = '%v'. "+
        "Instead, source file manager[j]='%v'. j='%v' ",
        sf[j], srcFmgr.GetAbsolutePathFileName(), j)
    }

    destFmgr := fOps.GetDestination()

    if df[j] != destFmgr.GetAbsolutePathFileName() {
      t.Errorf("Error: Expected destination file manager[j] = '%v'. "+
        "Instead, destination file manager[j]='%v'. j='%v' ",
        df[j], destFmgr.GetAbsolutePathFileName(), j)
    }

  }

}

func TestFileOpsCollection_AddByPathFileNameExtStrs_02(t *testing.T) {

  srcPathFileName := "../checkfiles/levelfilesfortest/level_0_0_test.txt"

  destPathFileName := "../dirmgrtests/level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  fOpsCol.fileOps = nil

  err := fOpsCol.AddByPathFileNameExtStrs(
    srcPathFileName,
    destPathFileName)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
      "Error='%v' ", err.Error())
    return
  }
}

func TestFileOpsCollection_AddByPathFileNameExtStrs_03(t *testing.T) {

  srcPathFileName := "../checkfiles/levelfilesfortest/level_0_0_test.txt"

  destPathFileName := "../dirmgrtests/level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  fOpsCol.fileOps = nil

  err := fOpsCol.AddByPathFileNameExtStrs(
    srcPathFileName,
    destPathFileName)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
      "Error='%v' ", err.Error())
    return
  }
}

func TestFileOpsCollection_AddByPathFileNameExtStrs_04(t *testing.T) {

  srcPathFileName := "../checkfiles/levelfilesfortest/level_0_0_test.txt"

  destPathFileName := "../dirmgrtests/level_0_1_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  srcPathFileName = "!@#$%^&////////......txt"

  err := fOpsCol.AddByPathFileNameExtStrs(
    srcPathFileName,
    destPathFileName)

  if err == nil {
    t.Error(
      "Expected an error return from fOpsCol." +
        "AddByPathFileNameExtStrs(srcPathFileName,destPathFileName).\n"+
      "because 'srcPathFileName' is INVALID!\n")
    return
  }
}

func TestFileOpsCollection_CopyOut_01(t *testing.T) {


  srcDir := "../checkfiles/levelfilesfortest"

  sf := make([]string, 5, 10)

  sf[0] = "level_0_0_test.txt"
  sf[1] = "level_0_1_test.txt"
  sf[2] = "level_0_2_test.txt"
  sf[3] = "level_0_3_test.txt"
  sf[4] = "level_0_4_test.txt"

  destDir := "../dirmgrtests"
  df := make([]string, 5, 10)

  df[0] = "level_0_0_test.txt"
  df[1] = "level_0_1_test.txt"
  df[2] = "level_0_2_test.txt"
  df[3] = "level_0_3_test.txt"
  df[4] = "level_0_4_test.txt"

  fOpsCol1 := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol1.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], destDir, df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol1.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], "+
        "destDir, df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }
  }

  fOpsCol2, err := fOpsCol1.CopyOut()

  if err != nil {
    t.Errorf("Error returned by fOpsCol1.CopyOut()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if ! fOpsCol1.Equal(&fOpsCol2) {
    t.Error("Error: After 'CopyOut' it was expected that fOpsCol1==fOpsCol2.\n" +
      "However, fOpsCol1 IS NOT EQUAL TO fOpsCol2.\n")
  }

}

func TestFileOpsCollection_CopyOut_02(t *testing.T) {


  fOpsCol1 := FileOpsCollection{}

  fOpsCol2 := FileOpsCollection{}

  fOpsCol1.fileOps = nil
  fOpsCol2.fileOps = nil

  if ! fOpsCol1.Equal(&fOpsCol2) {
    t.Error("Error: After 'CopyOut' it was expected that fOpsCol1==fOpsCol2.\n" +
      "However, fOpsCol1 IS NOT EQUAL TO fOpsCol2.\n")
  }

}
