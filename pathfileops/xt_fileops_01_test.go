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
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
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
    t.Errorf("Error returned by fOp1.ExecuteFileOperation(" +
      "FileOpCode.MoveSourceFileToDestinationFile())\n" +
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
      "destFile='%v'\nError='%v'\n", expectedAbsDestFileName, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
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
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
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
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
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
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_07(t *testing.T) {

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
    _ = fh.DeleteDirFile(sourceFile)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_07.htm"

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }


  if fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After delete operation, the destination\n" +
      "Still Exists!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CopySourceToDestinationByHardLinkByIo())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CopySourceToDestinationByHardLinkByIo())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After copy to destination file,\n" +
      "the source file DOES NOT EXIST!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After copy operation," +
      "the destination file DOES NOT EXIST!!!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_08(t *testing.T) {

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
    _ = fh.DeleteDirFile(sourceFile)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_08.htm"

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }


  if fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After delete operation, the destination\n" +
      "Still Exists!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CopySourceToDestinationByIoByHardLink())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CopySourceToDestinationByIoByHardLink())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After copy to destination file,\n" +
      "the source file DOES NOT EXIST!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After copy operation," +
      "the destination file DOES NOT EXIST!!!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_09(t *testing.T) {

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
    _ = fh.DeleteDirFile(sourceFile)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperation_08.htm"

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("Test Setup Error: After delete operation, the destination\n" +
      "Still Exists!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CopySourceToDestinationByHardLink())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CopySourceToDestinationByHardLink())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirFile(sourceFile)
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After copy to destination file,\n" +
      "the source file DOES NOT EXIST!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After copy operation," +
      "the destination file DOES NOT EXIST!!!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirFile(destFile)
    _ = fh.DeleteDirFile(sourceFile)

    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(destFile)\n" +
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(sourceFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(sourceFile)\n" +
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_10(t *testing.T) {


  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_10/006860_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_10"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }


  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_10/006860_sample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_10"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateSourceDir())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateSourceDir())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if !fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create source directory operation,\n" +
      "the source directory DOES NOT EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create source directory operation," +
      "the destination file suddenly Exists!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\nError='%v'\n", destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_11(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_11/006860_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_11"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_11/006860_sample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_11"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateSourceDirAndFile())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateSourceDirAndFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if !fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create source directory operation,\n" +
      "the source directory DOES NOT EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After create source directory/file operation,\n" +
      "the source file DOES NOT EXIST!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create source directory operation," +
      "the destination file suddenly Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_12(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_12/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_12"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_12/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_12"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fh.MakeDirAll(sourceDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeDirAll(sourceDir).\n" +
      "Attempted creation of soruce directory failed!\n" +
      "sourceDir='%v'\n" +
      "Error='%v\n", sourceDir, err.Error())

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateSourceFile())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateSourceFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After create source file operation,\n" +
      "the source directory DOES NOT EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }


  if fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create source directory operation," +
      "the destination file suddenly Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_13(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_13/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_13"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_13/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_13"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateDestinationDir())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateDestinationDir())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create destination directory operation,\n" +
      "the source directory DOES EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if !fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create destination directory operation," +
      "the destination directory DOES NOT EXIST!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_14(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_14/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_14"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_14/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_14"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateSourceDirAndFile())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateSourceDirAndFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if !fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create source directory/file operation,\n" +
      "the source directory DOES NOT EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if !fh.DoesFileExist(sourceFile) {
    t.Errorf("ERROR: After create source directory/file operation,\n" +
      "the source file DOES NOT EXIST!\n" +
      "sourceFile='%v'\n", sourceFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }



  if fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create source directory/file operation," +
      "the destination directory DOES EXIST!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After create source directory/file operation," +
      "the destination file DOES EXIST!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_15(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_15/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_15"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_15/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_15"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  fOpCode := FileOperationCode(99)

  err = fOp.ExecuteFileOperation(fOpCode)

  if err == nil {
    t.Error("Expected an error return from fOp.ExecuteFileOperation(" +
      "fOpCode)\n" +
      "because fOpCode is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_16(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_16/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_16"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_16/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_16"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fh.MakeDirAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateDestinationFile())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateDestinationFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create destination file operation,\n" +
      "the source directory DOES EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if !fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create destination file operation," +
      "the destination directory DOES NOT EXIST!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }


  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After create destination file operation," +
      "the destination file DOES NOT EXIST!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_ExecuteFileOperation_17(t *testing.T) {

  sourceFile :=
    "../checkfiles/TestFileOps_ExecuteFileOperationSrc_17/src_sample.htm"

  sourceDir := "../checkfiles/TestFileOps_ExecuteFileOperationSrc_17"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    return
  }

  destFile :=
    "../createFilesTest/TestFileOps_ExecuteFileOperationDst_17/destSample.htm"

  destDir := "../createFilesTest/TestFileOps_ExecuteFileOperationDst_17"

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n", destDir, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(destDir) {
    t.Errorf("Test Setup Error: After delete operation, the destination directory\n" +
      "Still Exists!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  fOp, err := FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sourceFile, destFile)\n" +
      "sourceFile='%v'\n" +
      "destFile='%v'\n" +
      "Error='%v'\n", sourceFile, destFile, err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  err = fOp.ExecuteFileOperation(FileOpCode.CreateDestinationDirAndFile())

  if err != nil {
    t.Errorf("Error returned by fOp.ExecuteFileOperation(" +
      "FileOpCode.CreateDestinationDirAndFile())\n" +
      "Error='%v'.\n", err.Error())
    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)
    return
  }

  if fh.DoesFileExist(sourceDir) {
    t.Errorf("ERROR: After create destination directory/file operation,\n" +
      "the source directory DOES EXIST!\n" +
      "sourceDir='%v'\n", sourceDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  if !fh.DoesFileExist(destDir) {
    t.Errorf("ERROR: After create destination directory/file operation," +
      "the destination directory DOES NOT EXIST!\n" +
      "destDir='%v'\n", destDir)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }


  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: After create destination directory/file operation," +
      "the destination file DOES NOT EXIST!\n" +
      "destFile='%v'\n", destFile)

    _ = fh.DeleteDirPathAll(sourceDir)
    _ = fh.DeleteDirPathAll(destDir)

    return
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(destDir)\n" +
      "destDir='%v'\n" +
      "Error='%v'\n",
      destDir, err.Error())
  }

  err = fh.DeleteDirPathAll(sourceDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(sourceDir)\n" +
      "sourceDir='%v'\n" +
      "Error='%v'\n", sourceDir, err.Error())
  }

  return
}

func TestFileOps_IsInitialized_01(t *testing.T) {

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

