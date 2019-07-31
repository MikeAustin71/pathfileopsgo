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

func TestFileOps_Equal_01(t *testing.T) {

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

  if !fOp1.Equal(&fOp2) {
    t.Error("ERROR: fOp1 and fOp2 have equivalent path, file names and file extensions.\n" +
      "However, fOp1.Equal(&fOp2) returned 'false'!\n")
  }

}

func TestFileOps_Equal_02(t *testing.T) {

  sourcePath1 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath1 := "../dirmgrtests/level_0_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)\n" +
      "sourcePath1='%v'\n" +
      "destPath1='%v'\n" +
      "Error='%v'\n",
      sourcePath1, destPath1, err.Error())

    return
  }

  sourcePath2 := "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  destPath2 := "../dirmgrtests/level_0_0_test.txt"

  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)\n" +
      "sourcePath2='%v'\n" +
      "destPath2='%v'\n" +
      "Error='%v'\n",
      sourcePath2, destPath2, err.Error())

    return
  }

  if fOp1.Equal(&fOp2) {
    t.Error("ERROR: fOp1 and fOp2 different source path, file names and file extensions.\n" +
      "However, fOp1.Equal(&fOp2) returned 'equal'!\n")
  }

}

func TestFileOps_Equal_03(t *testing.T) {

  sourcePath1 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath1 := "../dirmgrtests/level_0_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)\n" +
      "sourcePath1='%v'\n" +
      "destPath1='%v'\n" +
      "Error='%v'\n",
      sourcePath1, destPath1, err.Error())

    return
  }

  sourcePath2 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath2 := "../dirmgrtests/level_0_2_test.txt"

  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)\n" +
      "sourcePath2='%v'\n" +
      "destPath2='%v'\n" +
      "Error='%v'\n",
      sourcePath2, destPath2, err.Error())

    return
  }

  if fOp1.Equal(&fOp2) {
    t.Error("ERROR: fOp1 and fOp2 different destination path, file names and file extensions.\n" +
      "However, fOp1.Equal(&fOp2) returned 'equal'!\n")
  }

}

func TestFileOps_Equal_04(t *testing.T) {

  sourcePath1 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath1 := "../dirmgrtests/level_0_0_test.txt"

  fOps := FileOps{}

  fOp1, err := fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)

  if err != nil {
    t.Errorf("Error returned by #1 fOps.NewByPathFileNameExtStrs(sourcePath1, destPath1)\n" +
      "sourcePath1='%v'\n" +
      "destPath1='%v'\n" +
      "Error='%v'\n",
      sourcePath1, destPath1, err.Error())

    return
  }

  err = fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())

  if err != nil {
    t.Errorf("Error returned by fOp1.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  sourcePath2 := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  destPath2 := "../dirmgrtests/level_0_0_test.txt"

  fOp2, err := fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)

  if err != nil {
    t.Errorf("Error returned by #2 fOps.NewByPathFileNameExtStrs(sourcePath2, destPath2)\n" +
      "sourcePath2='%v'\n" +
      "destPath2='%v'\n" +
      "Error='%v'\n",
      sourcePath2, destPath2, err.Error())

    return
  }

  err = fOp2.SetFileOpsCode(FileOpCode.CopySourceToDestinationByHardLinkByIo())

  if err != nil {
    t.Errorf("Error returned by fOp2.SetFileOpsCode(FileOpCode.CopySourceToDestinationByHardLinkByIo())\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if fOp1.Equal(&fOp2) {
    t.Error("ERROR: fOp1 and fOp2 different file operation codes.\n" +
      "However, fOp1.Equal(&fOp2) returned 'equal'!\n")
  }
}

func TestFileOps_Equal_05(t *testing.T) {

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
    t.Error("ERROR: fOp1 and fOp2 have equivalent path, file names, file extensions and" +
      "File Ops Codes.\n" +
      "However, fOp1.Equal(&fOp2) returned 'false'!\n")
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

func TestFileOps_ExecuteFileOperation_01(t *testing.T) {

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

  err = fOp1.ExecuteFileOperation(FileOpCode.None())

  if err == nil {
    t.Error("Expected an error from fOp1.ExecuteFileOperation(FileOpCode.None())\n" +
      "because the File Operations Code is 'None'.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

}

func TestFileOps_ExecuteFileOperation_02(t *testing.T) {

  setupFile := "../filesfortest/htmlFilesForTest/006860_sample.htm"

  sourceFile :=
    "../checkfiles/checkfiles02/006860_sample.htm"

  fh := FileHelper{}

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n" +
      "setupFile='%v'\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n", setupFile, sourceFile, err.Error())
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_02.htm"

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n" +
      "Destination file already exists!\n" +
      "Error='%v'\n", err.Error())
    return
  }

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    return
  }

  err = fOp1.ExecuteFileOperation(FileOpCode.MoveSourceFileToDestinationFile())

  if err != nil {
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(FileOpCode.MoveSourceFileToDestinationFile())\n" +
      "Error='%v'.\n", err.Error())
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: The destination file does NOT exist!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: Source file still exists!\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n", destFile)
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  return

}

func TestFileOps_ExecuteFileOperation_03(t *testing.T) {

  setupFile := "../filesfortest/htmlFilesForTest/006860_sample.htm"

  sourceFile :=
    "../checkfiles/checkfiles02/006860_sample.htm"

  fh := FileHelper{}

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n" +
      "setupFile='%v'\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n", setupFile, sourceFile, err.Error())
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_02.htm"

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n" +
      "Destination file already exists!\n" +
      "destFile='%v'" +
      "Error='%v'\n",
      destFile,
      err.Error())
    return
  }

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by #1 FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    return
  }

  expectedAbsDestFileName :=
    "../createFilesTest/006860_sample.htm"

  expectedAbsDestFileName,
  err = fh.MakeAbsolutePath(expectedAbsDestFileName)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(expectedAbsDestFileName)\n" +
      "expectedAbsDestFileName='%v'\n" +
      "Error='%v'\n",
      expectedAbsDestFileName,
      err.Error())
  }

  err = fh.DeleteDirFile(expectedAbsDestFileName)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(expectedAbsDestFileName)\n" +
      "Destination file already exists!\n" +
      "expectedAbsDestFileName='%v'" +
      "Error='%v'\n",
      expectedAbsDestFileName, err.Error())
    return
  }

  err = fOp1.ExecuteFileOperation(FileOpCode.MoveSourceFileToDestinationDir())

  if err != nil {
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(FileOpCode.MoveSourceFileToDestinationFile())\n" +
      "Error='%v'.\n", err.Error())
    return
  }

  if !fh.DoesFileExist(expectedAbsDestFileName) {
    t.Errorf("ERROR: The destination file does NOT exist!\n" +
      "destFile='%v'\n", expectedAbsDestFileName)

    _ = fh.DeleteDirFile(expectedAbsDestFileName)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: Source file still exists!\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  err = fh.DeleteDirFile(expectedAbsDestFileName)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n", expectedAbsDestFileName)
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  return
}

func TestFileOps_ExecuteFileOperation_04(t *testing.T) {

  setupFile := "../filesfortest/htmlFilesForTest/006860_sample.htm"

  sourceFile :=
    "../checkfiles/checkfiles02/006860_sample.htm"

  fh := FileHelper{}

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n" +
      "setupFile='%v'\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n", setupFile, sourceFile, err.Error())
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_04.htm"

  err = fh.CopyFileByIo(setupFile, destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, destFile)\n" +
      "setupFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      setupFile,
      destFile,
      err.Error())
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After copy operation, the destination\n" +
      "Does NOT Exist!\n" +
      "destFile='%v'\n", destFile)
    return
  }

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by  FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp1.ExecuteFileOperation(FileOpCode.DeleteDestinationFile())

  if err != nil {
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(FileOpCode.DeleteDestinationFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After deletion of destination file, IT STILL EXISTS!!!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n", destFile)
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  return
}

func TestFileOps_ExecuteFileOperation_05(t *testing.T) {

  setupFile := "../filesfortest/htmlFilesForTest/006860_sample.htm"

  sourceFile :=
    "../checkfiles/checkfiles02/006860_sample.htm"

  fh := FileHelper{}

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n" +
      "setupFile='%v'\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n", setupFile, sourceFile, err.Error())
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_05.htm"

  err = fh.CopyFileByIo(setupFile, destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, destFile)\n" +
      "setupFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      setupFile,
      destFile,
      err.Error())
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After copy operation, the destination\n" +
      "Does NOT Exist!\n" +
      "destFile='%v'\n", destFile)
    return
  }

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by  FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp1.ExecuteFileOperation(FileOpCode.DeleteSourceFile())

  if err != nil {
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(FileOpCode.DeleteSourceFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After deletion of source file, IT STILL EXISTS!!!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n", destFile)
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  return
}

func TestFileOps_ExecuteFileOperation_06(t *testing.T) {

  setupFile := "../filesfortest/htmlFilesForTest/006860_sample.htm"

  sourceFile :=
    "../checkfiles/checkfiles02/006860_sample.htm"

  fh := FileHelper{}

  err := fh.CopyFileByIo(setupFile, sourceFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, sourceFile)\n" +
      "setupFile='%v'\n" +
      "sourceFile='%v'\n" +
      "Error='%v'\n", setupFile, sourceFile, err.Error())
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_06.htm"

  err = fh.CopyFileByIo(setupFile, destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupFile, destFile)\n" +
      "setupFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n",
      setupFile,
      destFile,
      err.Error())
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After copy operation, the destination\n" +
      "Does NOT Exist!\n" +
      "destFile='%v'\n", destFile)
    return
  }

  fOp1, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by  FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp1.ExecuteFileOperation(FileOpCode.DeleteSourceAndDestinationFiles())

  if err != nil {
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(FileOpCode.DeleteSourceFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After deletion of source file, IT STILL EXISTS!!!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After deletion of destination file, IT STILL EXISTS!!!\n" +
      "destFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n", destFile)
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\n", sourceFile)
  }

  return
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

