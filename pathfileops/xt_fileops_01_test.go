package pathfileops

import "testing"

func TestFileOps_CopyOut_01(t *testing.T) {
  sourcePath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath := "../dirmgrtests/level_0_0_test.txt"

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourcePath, destPath)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath, destPath)\n" +
      "sourcePath='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath, err.Error())

    return
  }

  fOp1.isInitialized = false

  fOp2 := fOp1.CopyOut()

  fOp3 := FileOps{}

  if !fOp2.Equal(&fOp3) {
    t.Errorf("Expected fOp2 == FileOps{}.\n" +
      "Instead, fOp2 was NOT EQUAL to an empty FileOps instance.\n")
  }
}

func TestFileOps_CopyOut_02(t *testing.T) {

  sourcePath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath := "../dirmgrtests/level_0_0_test.txt"

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourcePath, destPath)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath, destPath)\n" +
      "sourcePath='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath, err.Error())

    return
  }

  fOp2 := fOp1.CopyOut()

  if !fOp2.Equal(&fOp1) {
    t.Errorf("Expected fOp2 == fOp1.\n" +
      "Instead, fOp2 was NOT EQUAL to fOp1.\n")
  }
}

func TestFileOps_EqualPathFileNameExt_01(t *testing.T) {
  sourcePath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath := "../dirmgrtests/level_0_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath, destPath)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath, destPath)\n" +
      "sourcePath='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath, err.Error())

    return
  }


  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath, destPath)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath, destPath)\n" +
      "sourcePath='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath, err.Error())

    return
  }

  if !fOp1.EqualPathFileNameExt(&fOp2) {
    t.Error("ERROR: fOp1 and fOp2 have equivalent path, file names and file extensions.\n" +
      "However, fOp1.EqualPathFileNameExt(&fOp2) returned 'false'!\n")
  }

}

func TestFileOps_EqualPathFileNameExt_02(t *testing.T) {

  sourcePath1 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sourcePath2 := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_2_0_test.txt"

  destPath := "../dirmgrtests/level_0_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath1, destPath)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath1, destPath)\n" +
      "sourcePath1='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath1, destPath, err.Error())

    return
  }


  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath2, destPath)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath2, destPath)\n" +
      "sourcePath2='%v'\n" +
      "destPath='%v'\n" +
      "Error='%v'\n",
      sourcePath2, destPath, err.Error())

    return
  }

  if fOp1.EqualPathFileNameExt(&fOp2) == true {
    t.Error("ERROR: fOp1 and fOp2 have different source path, file names and file extensions.\n" +
      "However, fOp1.EqualPathFileNameExt(&fOp2) returned 'true' signaling equivalency!\n")
  }

}

func TestFileOps_EqualPathFileNameExt_03(t *testing.T) {

  sourcePath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  destPath1 := "../dirmgrtests/level_0_0_test.txt"
  destPath2 := "../dirmgrtests/level_01_dir/level_02_dir/level_2_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath, destPath1)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath, destPath1)\n" +
      "sourcePath='%v'\n" +
      "destPath1='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath1, err.Error())

    return
  }


  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath, destPath2)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath, destPath2)\n" +
      "sourcePath='%v'\n" +
      "destPath2='%v'\n" +
      "Error='%v'\n",
      sourcePath, destPath2, err.Error())

    return
  }

  if fOp1.EqualPathFileNameExt(&fOp2) == true {
    t.Error("ERROR: fOp1 and fOp2 have different destination path, file names and file extensions.\n" +
      "However, fOp1.EqualPathFileNameExt(&fOp2) returned 'true' signaling equivalency!\n")
  }

}

func TestFileOps_IsInitialized(t *testing.T) {


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

  fOp, err := FileOps{}.NewByFileMgrs(sourceFMgr, destFileMgr)


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

  isInitialized := fOp.IsInitialized()

  if isInitialized==false{
    t.Errorf("ERROR: Expected isInitialized=='true'.\n" +
      "However, the actual returned value is 'false'.\n")
  }

}

