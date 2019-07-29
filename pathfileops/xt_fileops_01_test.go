package pathfileops

import "testing"


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

