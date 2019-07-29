package pathfileops

import "testing"

func TestFileOps_NewByFileMgrs_01(t *testing.T) {

  sourceFile := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"
  destFile := "../createFilesTest/Level01/Level02/TestFileOps_NewByFileMgrs_01.txt"

  sourceFMgr, err := FileMgr{}.New(sourceFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(sourceFile)\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n",
      sourceFile, err.Error())
  }

  destFileMgr, err := FileMgr{}.New(destFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      destFile, err.Error())
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
