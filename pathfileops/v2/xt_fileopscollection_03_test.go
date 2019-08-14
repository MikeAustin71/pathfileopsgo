package pathfileops

import "testing"

func TestFileOpsCollection_InsertFileOpsAtIndex_01(t *testing.T) {
  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  // # 2
  origPath := fh.AdjustPathSlash("../../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrSrc, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../dirmgrtests/CmdrX.log")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrDst, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  fOp, err := FileOps{}.NewByFileMgrs(fmgrSrc, fmgrDst)

  err = fOpsCol.InsertFileOpsAtIndex(fOp, 2)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.InsertFileOpsAtIndex(fOp, 2). "+
      "Error='%v'", err.Error())
  }

  arrayLen = fOpsCol.GetNumOfFileOps()

  if arrayLen != 6 {
    t.Errorf("Error: Expected after insertion array length='6'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  fOpVerify, err := fOpsCol.PeekFileOpsAtIndex(2)

  if !fOp.Equal(&fOpVerify) {
    t.Error("Expected original file operation to be equal to extracted file operation. " +
      "They are NOT equal!")
  }

}

func TestFileOpsCollection_InsertFileOpsAtIndex_02(t *testing.T) {
  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  // # 2
  origPath := fh.AdjustPathSlash("../../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrSrc, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../dirmgrtests/CmdrX.log")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrDst, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  fOp, err := FileOps{}.NewByFileMgrs(fmgrSrc, fmgrDst)

  err = fOpsCol.InsertFileOpsAtIndex(fOp, 3)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.InsertFileOpsAtIndex(fOp, 3). "+
      "Error='%v'", err.Error())
  }

  arrayLen = fOpsCol.GetNumOfFileOps()

  if arrayLen != 6 {
    t.Errorf("Error: Expected after insertion array length='6'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  fOpVerify, err := fOpsCol.PeekFileOpsAtIndex(3)

  if !fOp.Equal(&fOpVerify) {
    t.Error("Expected original file operation to be equal to extracted file operation. " +
      "They are NOT equal!")
  }

}

func TestFileOpsCollection_InsertFileOpsAtIndex_03(t *testing.T) {
  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  // # 2
  origPath := fh.AdjustPathSlash("../../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrSrc, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../dirmgrtests/CmdrX.log")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrDst, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  fOp, err := FileOps{}.NewByFileMgrs(fmgrSrc, fmgrDst)

  err = fOpsCol.InsertFileOpsAtIndex(fOp, 99)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.InsertFileOpsAtIndex(fOp, 99). "+
      "Error='%v'", err.Error())
  }

  arrayLen = fOpsCol.GetNumOfFileOps()

  if arrayLen != 6 {
    t.Errorf("Error: Expected after insertion array length='6'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  fOpVerify, err := fOpsCol.PeekFileOpsAtIndex(5)

  if !fOp.Equal(&fOpVerify) {
    t.Error("Expected original file operation to be equal to extracted file operation. " +
      "They are NOT equal!")
  }

}

func TestFileOpsCollection_InsertFileOpsAtIndex_04(t *testing.T) {
  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fh := FileHelper{}
  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v' Error='%v' ", i, err.Error())
      return
    }

  }

  arrayLen := fOpsCol.GetNumOfFileOps()

  if arrayLen != 5 {
    t.Errorf("Error: Expected intial array length='5'. "+
      "Instead, array length='%v' ", arrayLen)
  }

  // # 2
  origPath := fh.AdjustPathSlash("../../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrSrc, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // # 3
  origPath = fh.AdjustPathSlash("../../dirmgrtests/CmdrX.log")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  fmgrDst, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("%v\n", err.Error())
  }

  fOp, err := FileOps{}.NewByFileMgrs(fmgrSrc, fmgrDst)

  err = fOpsCol.InsertFileOpsAtIndex(fOp, -3)

  if err == nil {
    t.Error("Error: Expected an error return from err = fOpsCol." +
      "InsertFileOpsAtIndex(fOp, -3). NO ERROR WAS RETURNED!! ")
  }

}

func TestFileOpsCollection_InsertFileOpsAtIndex_05(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  srcFile := "../../filesfortest/levelfilesfortest/level_0_0_test.txt"

  destFile := "../../dirmgrtests/level_0_0_test.txt"

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(srcFile, destFile).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fOpsCol.fileOps = nil

  err = fOpsCol.InsertFileOpsAtIndex(fOp, 0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.InsertFileOpsAtIndex(fOp, 0)\n" +
      "Error='%v'\n", err.Error())
  }
}

func TestFileOpsCollection_InsertFileOpsAtIndex_06(t *testing.T) {
  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]).\n"+
        "i='%v'\nError='%v'\n", i, err.Error())
      return
    }
  }

  srcFile := "../../filesfortest/levelfilesfortest/level_0_0_test.txt"

  destFile := "../../dirmgrtests/level_0_0_test.txt"

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(srcFile, destFile).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fOpsCol.InsertFileOpsAtIndex(fOp, 0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.InsertFileOpsAtIndex(fOp, 0)\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestFileOpsCollection_NewFromFileMgrCollection_01(t *testing.T) {

  sf := make([]string, 10, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"
  sf[5] = "../../filesfortest/levelfilesfortest/level_0_5_test.txt"
  sf[6] = "../../filesfortest/levelfilesfortest/level_0_6_test.txt"
  sf[7] = "../../filesfortest/levelfilesfortest/level_0_7_test.txt"
  sf[8] = "../../filesfortest/levelfilesfortest/level_0_8_test.txt"
  sf[9] = "../../filesfortest/levelfilesfortest/level_0_9_test.txt"

  baseDir :="../../filesfortest/levelfilesfortest"

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(baseDir)\n" +
      "baseDir='%v'\n" +
      "Error='%v'\n", baseDir, err.Error())
  }

  targetDir := "../../dirmgrtests"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\n" +
      "Error='%v'\n", targetDir, err.Error())
  }

  fMgrCol := FileMgrCollection{}

  for i:=0; i < 10; i++ {

    err := fMgrCol.AddFileMgrByPathFileNameExt(sf[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(sf[%v])\n" +
        "sf[%v]='%v'\n" +
        "Error='%v'\n", i, i, sf[i], err.Error())
      return
    }
  }

  fOpsCol, err := FileOpsCollection{}.NewFromFileMgrCollection(&fMgrCol, &baseDMgr, &targetDMgr)

  if err != nil {
    t.Errorf("Error returned by FileOpsCollection{}.NewFromFileMgrCollection(" +
      "&fMgrCol, &baseDMgr, &targetDMgr)\n" +
      "baseDMgr='%v'\n" +
      "targetDMgr='%v'\n" +
      "Error='%v'\n", baseDMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath(), err.Error())
    return
  }

  fOp, err := fOpsCol.PeekFileOpsAtIndex(9)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(9)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualDstFMgr := fOp.GetDestination()

  actualDstDMgr := actualDstFMgr.GetDirMgr()

  if !targetDMgr.EqualAbsPaths(&actualDstDMgr) {
    t.Errorf("ERROR: Expected targetDMgr==actualDstDMgr.\n" +
      "Instead, they ARE NOT EQUAL!\n" +
      "targetDMgr='%v'\n" +
      "actualDstDMgr='%v'\n", targetDMgr.GetAbsolutePath(), actualDstDMgr.GetAbsolutePath())
  }

}

func TestFileOpsCollection_NewFromFileMgrCollection_02(t *testing.T) {

  sf := make([]string, 10, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"
  sf[5] = "../../filesfortest/levelfilesfortest/level_0_5_test.txt"
  sf[6] = "../../filesfortest/levelfilesfortest/level_0_6_test.txt"
  sf[7] = "../../filesfortest/levelfilesfortest/level_0_7_test.txt"
  sf[8] = "../../filesfortest/levelfilesfortest/level_0_8_test.txt"
  sf[9] = "../../filesfortest/levelfilesfortest/level_0_9_test.txt"

  baseDir :="../../filesfortest/levelfilesfortest"

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(baseDir)\n" +
      "baseDir='%v'\n" +
      "Error='%v'\n", baseDir, err.Error())
  }

  targetDir := "../../dirmgrtests"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\n" +
      "Error='%v'\n", targetDir, err.Error())
  }

  fMgrCol := FileMgrCollection{}

  for i:=0; i < 10; i++ {

    err := fMgrCol.AddFileMgrByPathFileNameExt(sf[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(sf[%v])\n" +
        "sf[%v]='%v'\n" +
        "Error='%v'\n", i, i, sf[i], err.Error())
      return
    }
  }

  fOpsCollection := FileOpsCollection{}

  fOpsCollection.fileOps = nil

  fOpsCol, err := FileOpsCollection{}.NewFromFileMgrCollection(&fMgrCol, &baseDMgr, &targetDMgr)

  if err != nil {
    t.Errorf("Error returned by FileOpsCollection{}.NewFromFileMgrCollection(" +
      "&fMgrCol, &baseDMgr, &targetDMgr)\n" +
      "baseDMgr='%v'\n" +
      "targetDMgr='%v'\n" +
      "Error='%v'\n", baseDMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath(), err.Error())
    return
  }

  fOp, err := fOpsCol.PeekFileOpsAtIndex(9)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(9)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualDstFMgr := fOp.GetDestination()

  actualDstDMgr := actualDstFMgr.GetDirMgr()

  if !targetDMgr.EqualAbsPaths(&actualDstDMgr) {
    t.Errorf("ERROR: Expected targetDMgr==actualDstDMgr.\n" +
      "Instead, they ARE NOT EQUAL!\n" +
      "targetDMgr='%v'\n" +
      "actualDstDMgr='%v'\n", targetDMgr.GetAbsolutePath(), actualDstDMgr.GetAbsolutePath())
  }

}

func TestFileOpsCollection_NewFromFileMgrCollection_03(t *testing.T) {


  baseDir :="../../filesfortest/levelfilesfortest"

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(baseDir)\n" +
      "baseDir='%v'\n" +
      "Error='%v'\n", baseDir, err.Error())
  }

  targetDir := "../../dirmgrtests"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\n" +
      "Error='%v'\n", targetDir, err.Error())
  }

  fMgrCol := FileMgrCollection{}
  fMgrCol.fileMgrs = nil

  fOpsCollection := FileOpsCollection{}
   fOpsCollection.fileOps = nil

  _, err = fOpsCollection.NewFromFileMgrCollection(&fMgrCol, &baseDMgr, &targetDMgr)

  if err == nil {
    t.Error("Expected an error return from FileOpsCollection{}." +
      "NewFromFileMgrCollection(&fMgrCol, &baseDMgr, &targetDMgr)\n" +
      "because the 'fMgrCol' is EMPTY!\n")
    return
  }

}

func TestFileOpsCollection_NewFromFileMgrCollection_04(t *testing.T) {

  sf := make([]string, 10, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"
  sf[5] = "../../filesfortest/levelfilesfortest/level_0_5_test.txt"
  sf[6] = "../../filesfortest/levelfilesfortest/level_0_6_test.txt"
  sf[7] = "../../filesfortest/levelfilesfortest/level_0_7_test.txt"
  sf[8] = "../../filesfortest/levelfilesfortest/level_0_8_test.txt"
  sf[9] = "../../filesfortest/levelfilesfortest/level_0_9_test.txt"

  targetDir := "../../dirmgrtests"

  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\n" +
      "Error='%v'\n", targetDir, err.Error())
  }

  fMgrCol := FileMgrCollection{}

  for i:=0; i < 10; i++ {

    err := fMgrCol.AddFileMgrByPathFileNameExt(sf[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(sf[%v])\n" +
        "sf[%v]='%v'\n" +
        "Error='%v'\n", i, i, sf[i], err.Error())
      return
    }
  }

  fOpsCollection := FileOpsCollection{}

  fOpsCollection.fileOps = nil

  _, err = fOpsCollection.NewFromFileMgrCollection(&fMgrCol, nil, &targetDMgr)

  if err == nil {
    t.Error("Expected an error return from fOpsCollection.NewFromFileMgrCollection(" +
      "&fMgrCol, nil, &targetDMgr)\n" +
      "because 'sourceBaseDir' is 'nil'!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

func TestFileOpsCollection_NewFromFileMgrCollection_05(t *testing.T) {

  sf := make([]string, 10, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"
  sf[5] = "../../filesfortest/levelfilesfortest/level_0_5_test.txt"
  sf[6] = "../../filesfortest/levelfilesfortest/level_0_6_test.txt"
  sf[7] = "../../filesfortest/levelfilesfortest/level_0_7_test.txt"
  sf[8] = "../../filesfortest/levelfilesfortest/level_0_8_test.txt"
  sf[9] = "../../filesfortest/levelfilesfortest/level_0_9_test.txt"

  baseDir :="../../filesfortest/levelfilesfortest"

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(baseDir)\n" +
      "baseDir='%v'\n" +
      "Error='%v'\n", baseDir, err.Error())
  }

  fMgrCol := FileMgrCollection{}

  for i:=0; i < 10; i++ {

    err := fMgrCol.AddFileMgrByPathFileNameExt(sf[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(sf[%v])\n" +
        "sf[%v]='%v'\n" +
        "Error='%v'\n", i, i, sf[i], err.Error())
      return
    }
  }

  fOpsCollection := FileOpsCollection{}

  fOpsCollection.fileOps = nil

  _, err = fOpsCollection.NewFromFileMgrCollection(&fMgrCol, &baseDMgr, nil)

  if err == nil {
    t.Error("Expected an error return from fOpsCollection.NewFromFileMgrCollection" +
      "(&fMgrCol, &baseDMgr, nil)\n" +
      "because 'targetBaseDir' is 'nil'!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PopFileOpsAtIndex_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()
  var expectedFOp FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i == 2 {
      expectedFOp = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PopFileOpsAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PopFileOpsAtIndex(2)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFOp.Equal(&actualFOp) {
    t.Errorf("ERROR: Expected expectedFOp == actualFOp.\n" +
      "However, Equal returned expectedFOp != actualFOp.\n" +
      "expectedFOp.source='%v'\n" +
      "expectedFOp.destination='%v'\n" +
      "actualFOp.source='%v'\n" +
      "actualFOp.destination='%v'\n",
      expectedFOp.source.GetAbsolutePathFileName(),
      expectedFOp.destination.GetAbsolutePathFileName(),
      actualFOp.source.GetAbsolutePathFileName(),
      actualFOp.destination.GetAbsolutePathFileName())
  }

  if fOpsCol.GetNumOfFileOps() != 4 {
    t.Errorf("ERROR: Expected that after PopFileOpsAtIndex operation, Number Of FileOps='4'.\n" +
      "Instead, Number Of FileOps='%v'\n", fOpsCol.GetNumOfFileOps() )
  }

}

func TestFileOpsCollection_PopFileOpsAtIndex_02(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()
  var expectedFOp FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i == 0 {
      expectedFOp = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PopFileOpsAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PopFileOpsAtIndex(0)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFOp.Equal(&actualFOp) {
    t.Errorf("ERROR: Expected expectedFOp == actualFOp.\n" +
      "However, Equal returned expectedFOp != actualFOp.\n" +
      "expectedFOp.source='%v'\n" +
      "expectedFOp.destination='%v'\n" +
      "actualFOp.source='%v'\n" +
      "actualFOp.destination='%v'\n",
      expectedFOp.source.GetAbsolutePathFileName(),
      expectedFOp.destination.GetAbsolutePathFileName(),
      actualFOp.source.GetAbsolutePathFileName(),
      actualFOp.destination.GetAbsolutePathFileName())
  }

  if fOpsCol.GetNumOfFileOps() != 4 {
    t.Errorf("ERROR: Expected that after PopFileOpsAtIndex operation, Number Of FileOps='4'.\n" +
      "Instead, Number Of FileOps='%v'\n", fOpsCol.GetNumOfFileOps() )
  }
}

func TestFileOpsCollection_PopFileOpsAtIndex_03(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()
  var expectedFOp FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i == 4 {
      expectedFOp = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PopFileOpsAtIndex(4)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PopFileOpsAtIndex(0)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFOp.Equal(&actualFOp) {
    t.Errorf("ERROR: Expected expectedFOp == actualFOp.\n" +
      "However, Equal returned expectedFOp != actualFOp.\n" +
      "expectedFOp.source='%v'\n" +
      "expectedFOp.destination='%v'\n" +
      "actualFOp.source='%v'\n" +
      "actualFOp.destination='%v'\n",
      expectedFOp.source.GetAbsolutePathFileName(),
      expectedFOp.destination.GetAbsolutePathFileName(),
      actualFOp.source.GetAbsolutePathFileName(),
      actualFOp.destination.GetAbsolutePathFileName())
  }

  if fOpsCol.GetNumOfFileOps() != 4 {
    t.Errorf("ERROR: Expected that after PopFileOpsAtIndex operation, Number Of FileOps='4'.\n" +
      "Instead, Number Of FileOps='%v'\n", fOpsCol.GetNumOfFileOps() )
  }
}

func TestFileOpsCollection_PopFileOpsAtIndex_04(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  _, err := fOpsCol.PopFileOpsAtIndex(-1)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PopFileOpsAtIndex(-1)\n" +
      "because the index '-1' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOpsCollection_PopFileOpsAtIndex_05(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  _, err := fOpsCol.PopFileOpsAtIndex(99)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PopFileOpsAtIndex(99)\n" +
      "because the index '99' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOpsCollection_PopFileOpsAtIndex_06(t *testing.T) {

  fOpsCol := FileOpsCollection{}
  fOpsCol.fileOps = nil

  _, err := fOpsCol.PopFileOpsAtIndex(2)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PopFileOpsAtIndex(2)\n" +
      "because the File Ops Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PopFirstFileOps_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==0 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PopFirstFileOps()

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PopFirstFileOps()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

  if fOpsCol.GetNumOfFileOps() != 4 {
    t.Errorf("ERROR: Expected the remaining File Operations object to be '4'.\n" +
      "Instead, Number Of FileOps='%v'\n", fOpsCol.GetNumOfFileOps())
  }

}

func TestFileOpsCollection_PopFirstFileOps_02(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  _, err := fOpsCol.PopFirstFileOps()

  if err == nil {
    t.Error("ERROR: Expected an error return from fOpsCol.PopFirstFileOps()\n" +
      "because the File Ops Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PopLastFileOps_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==4 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PopLastFileOps()

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PopLastFileOps()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

  if fOpsCol.GetNumOfFileOps() != 4 {
    t.Errorf("ERROR: Expected the remaining File Operations object to be '4'.\n" +
      "Instead, Number Of FileOps='%v'\n", fOpsCol.GetNumOfFileOps())
  }

}

func TestFileOpsCollection_PopLastFileOps_02(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  _, err := fOpsCol.PopLastFileOps()

  if err == nil {
    t.Error("ERROR: Expected an error return from fOpsCol.PopLastFileOps()\n" +
      "because the File Ops Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PeekFileOpsAtIndex_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==2 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PeekFileOpsAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(2)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

}

func TestFileOpsCollection_PeekFileOpsAtIndex_02(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==0 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PeekFileOpsAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(0)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

}

func TestFileOpsCollection_PeekFileOpsAtIndex_03(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==4 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PeekFileOpsAtIndex(4)

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(4)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

}

func TestFileOpsCollection_PeekFileOpsAtIndex_04(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }
  }

  _, err := fOpsCol.PeekFileOpsAtIndex(-1)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PeekFileOpsAtIndex(-1)\n" +
      "because the index, '-1', is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOpsCollection_PeekFileOpsAtIndex_05(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }
  }

  _, err := fOpsCol.PeekFileOpsAtIndex(99)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PeekFileOpsAtIndex(99)\n" +
      "because the index, '99', is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PeekFileOpsAtIndex_06(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}
  fOpsCol.fileOps = nil
  _, err := fOpsCol.PeekFileOpsAtIndex(0)

  if err == nil {
    t.Error("Expected an error return from fOpsCol.PeekFileOpsAtIndex(0)\n" +
      "because the File Operations Collections is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileOpsCollection_PeekFirstFileOps_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==0 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }

  }

  actualFOp, err := fOpsCol.PeekFirstFileOps()

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekFirstFileOps()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

}

func TestFileOpsCollection_PeekFirstFileOps_02(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  _, err := fOpsCol.PeekFirstFileOps()

  if err == nil {
    t.Error("ERROR: Expected an error return from fOpsCol.PeekFirstFileOps()\n" +
      "because the File Ops Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileOpsCollection_PeekLastFileOps_01(t *testing.T) {

  sf := make([]string, 5, 10)

  sf[0] = "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../../dirmgrtests/level_0_0_test.txt"
  df[1] = "../../dirmgrtests/level_0_1_test.txt"
  df[2] = "../../dirmgrtests/level_0_2_test.txt"
  df[3] = "../../dirmgrtests/level_0_3_test.txt"
  df[4] = "../../dirmgrtests/level_0_4_test.txt"

  fOpsCol := FileOpsCollection{}.New()

  var expectedFileOps FileOps

  for i := 0; i < 5; i++ {

    fOp, err := FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err.Error())
      return
    }

    if i==4 {
      expectedFileOps = fOp.CopyOut()
    }

    err = fOpsCol.AddByFileOps(fOp)

    if err != nil {
      t.Errorf("Error returned by fOpsCol.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i, sf[i], df[i], err.Error())
      return
    }
  }

  actualFOp, err := fOpsCol.PeekLastFileOps()

  if err != nil {
    t.Errorf("Error returned by fOpsCol.PeekLastFileOps()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFileOps.Equal(&actualFOp) {
    t.Error("ERROR: Expected expectedFileOps.Equal(&actualFOp)=='true'.\n" +
      "Instead, expectedFileOps.Equal(&actualFOp)=='false' !!\n")
  }

}

func TestFileOpsCollection_PeekLastFileOps_02(t *testing.T) {

  fOpsCol := FileOpsCollection{}

  fOpsCol.fileOps = nil

  _, err := fOpsCol.PeekLastFileOps()

  if err == nil {
    t.Error("ERROR: Expected an error return from fOpsCol.PeekLastFileOps()\n" +
      "because the File Ops Collection is EMPTY!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

