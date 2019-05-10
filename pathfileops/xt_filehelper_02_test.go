package pathfileops

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "testing"
  "time"
)

func TestFileHelper_CopyFileByLink_01(t *testing.T) {

  testDestFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink("", destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(\"\", destFile) " +
      "because src parameter was an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_02(t *testing.T) {

  testDestFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink("     ", destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(\"    \", destFile) " +
      "because src parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_03(t *testing.T) {

  testSrcFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, "")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, \"\") " +
      "because destination parameter was an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_04(t *testing.T) {

  testSrcFile := "../filesfortest/levelfilesfortest/level_9_9_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, "     ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, \"    \") " +
      "because destination parameter consists of all space characters. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_CopyFileByLink_05(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  testDestFile := "../checkfiles/scratchDestTileJJ459821.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testDestFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected error return from FileHelper{}.CopyFileByLink(srcFile, destFile) " +
      "because srcFile does not exist. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if fh.DoesFileExist(destFile) == true {
    err = fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error attempting file deletion of File='%v'. Error='%v' ",
        destFile, err.Error())
    }
  }
}

func TestFileHelper_CopyFileByLink_06(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, srcFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, srcFile) because Source File and Destination File are equivalent.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_07(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  destFile := ""

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Destination File is an EMPTY STRING!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_08(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  destFile := "   "

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Destination File consists" +
      "entirely of blank spaces!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_09(t *testing.T) {

  fh := FileHelper{}

  srcFile := ""

  testDestFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  destFile, err := fh.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Source File is an EMPTY STRING!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_10(t *testing.T) {

  fh := FileHelper{}

  srcFile := "   "

  testDestFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  destFile, err := fh.MakeAbsolutePath(testDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testDestFile='%v' Error='%v' ", testDestFile, err.Error())
  }

  err = FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Source File consists " +
      "entirely of blank spaces\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_11(t *testing.T) {

  fh := FileHelper{}

  srcFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  destFile := "../createFilesTest/scratchXFg924198.txt"

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  }

  err := fh.CopyFileByLink(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByLink(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
        "Deletion of Destination File Failed. "+
        "It cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  } else {
    t.Errorf("Error: CopyFileByLink Failed. Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
  }

}

func TestFileHelper_CopyFileByLink_12(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_0_3_test.txt")

  srcFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_0_1_test.txt")

  destFile :=
    fh.AdjustPathSlash("../createFilesTest/scratchTestFileHelper_CopyFileByLink_12.txt")

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  }

  err := fh.CopyFileByIo(testSrcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by ")
  }

  if !fh.DoesFileExist(destFile) {

    t.Errorf("Error: Setup Failed. Target destination file was NOT created! "+
      "fh.CopyFileByIo(srcFile, destFile) FAILED!\n"+
      "srcFile='%v'\ndestFile='%v'", srcFile, destFile)
    return

  }

  err = fh.CopyFileByLink(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByLink(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: The copy by link operation failed! " +
      "The destination file does NOT exist.\n"+
      "destFile='%v'\n",
      destFile)
    return
  }

  // Destination file exists
  finfoSrc, err := os.Stat(srcFile)

  if err != nil {
    t.Errorf("Error returned by os.Stat(srcFile).\n" +
      "srcFile='%v'\nError='%v'", srcFile, err.Error())
  }

  finfoDest, err := os.Stat(destFile)

  if err !=nil {
    t.Errorf("Error returned by os.Stat(destFile).\n" +
      "destFile='%v'\nError='%v'", destFile, err.Error())
  }

  if finfoSrc.Size() != finfoDest.Size() {
    t.Errorf("Error: Size of source file does NOT match " +
      "size of destination file.\nSourceFileSize='%v' DestinationFileSize='%v'\n",
      finfoSrc.Size(), finfoDest.Size(),)
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
      "Deletion of Destination File Failed. "+
      "It cannot be deleted!\ndestFile='%v'", destFile)
    return
  }

}

func TestFileHelper_CopyFileByLink_13(t *testing.T) {

  fh := FileHelper{}

  testSrcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"

  srcFile := fh.AdjustPathSlash(testSrcFile)

  err := FileHelper{}.CopyFileByLink(srcFile, srcFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, srcFile) because Source File and Destination File " +
      "are the same.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByLink_14(t *testing.T) {

  fh := FileHelper{}

  srcFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/iDoNotExist.txt")

  destFile :=
    fh.AdjustPathSlash("../createFilesTest/scratchTestFileHelper_CopyFileByLink_14.txt")

  err := FileHelper{}.CopyFileByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CopyFileByLink" +
      "(srcFile, destFile) because Source File DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CreateFile_01(t *testing.T) {

  err := deleteALogTestBottomDirTargetDir()
  if err != nil {
    t.Error("Failed to delete target Directory:", err)
  }

  err = createALogTestBottomDir()
  if err != nil {
    t.Error("Failed to delete target Directory:", err)
  }
}

func TestFileHelper_CreateFile_02(t *testing.T) {
  // Uses 'Create' to overwrite existing file
  tstFile := "..//logTest//testoverwrite//TestOverwrite001.txt"
  fh := FileHelper{}

  if fh.DoesFileExist(tstFile) {
    err := fh.DeleteDirFile(tstFile)
    if err != nil {
      t.Error(fmt.Sprintf("Error: Deletion Failed On File %v !", tstFile))
    }
  }

  if fh.DoesFileExist(tstFile) {
    t.Error(fmt.Sprintf("Error: Deletion Failed! File %v should not exist!", tstFile))
  }

  f, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Error(fmt.Sprintf("Error: Create File Failed for file: %v", tstFile))
  }

  _, err4 := f.WriteString(aLoremIpsumTxt)

  if err4 != nil {
    _ = f.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err4)
  }

  err = f.Close()

  if err != nil {
    t.Errorf("%v", err.Error())
  }

  // Now recreate the original file. It should be
  // truncated with the old contents deleted.
  f2, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Errorf("Error: Re-Creating File %v", tstFile)
  }

  fOvrWriteTxt := "Test Over Write and existing file using Create()"

  _, err5 := f2.WriteString(fOvrWriteTxt)

  if err5 != nil {
    _ = f2.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err5)
  }

  err = f2.Close()

  if err != nil {
    t.Errorf("Error closing f2. %v", err.Error())
  }

  dat, err := ioutil.ReadFile(tstFile)

  if err != nil {
    t.Errorf("Error Reading Re-Written Text for File:'%v' Error='%v'", tstFile, err)
  }

  s := string(dat)

  if s != fOvrWriteTxt {
    t.Errorf("Was expecting to read text: '%v', instead received text: %v", fOvrWriteTxt, s)
  }

}

func TestFileHelper_CreateFile_03(t *testing.T) {

  _, err := FileHelper{}.CreateFile("")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CreateFile(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CreateFile_04(t *testing.T) {

  _, err := FileHelper{}.CreateFile("    ")

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CreateFile(\"   \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CreateFile_05(t *testing.T) {

  testFilePath := "../checkfiles/TestFileHelper_CreateFile_05/dirX1/dirX2/scratchTestFileHelper_CreateFile_05.txt"

  fh := FileHelper{}

  testFilePath = fh.AdjustPathSlash(testFilePath)

  _, err := fh.CreateFile(testFilePath)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.CreateFile(testFilePath) " +
      "because parent directories for testFilePath do NOT exist. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CreateFile_06(t *testing.T) {

  testFilePath := "../checkfiles/scratchTestFileHelper_CreateFile_05.txt"

  fh := FileHelper{}

  testFilePath = fh.AdjustPathSlash(testFilePath)

  fPtr, err := fh.CreateFile(testFilePath)

  if err != nil {
    t.Errorf("Error returned from fh.CreateFile(testFilePath).\n" +
      "testFilePath='%v'\nError='%v'\n", testFilePath, err.Error())

  }

  if fPtr != nil {
    err = fPtr.Close()
    if err != nil {
      t.Errorf("Test Clean-Up Error returned by fPtr.Close().\nError='%v'\n",
        err.Error())
    }
  }

  err = fh.DeleteDirFile(testFilePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(testFilePath).\n" +
      "testFilePath='%v'\nError='%v'\n",testFilePath, err.Error())
  }

}

func TestFileHelper_DeleteDirFile_01(t *testing.T) {

  err := FileHelper{}.DeleteDirFile("")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirFile(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_DeleteDirFile_02(t *testing.T) {

  err := FileHelper{}.DeleteDirFile("   ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirFile(\"  \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_DeleteDirPathAll_01(t *testing.T) {

  err := FileHelper{}.DeleteDirPathAll("")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirPathAll(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_DeleteDirPathAll_02(t *testing.T) {

  err := FileHelper{}.DeleteDirPathAll("    ")

  if err == nil {
    t.Error("Expected error return from FileHelper{}.DeleteDirPathAll(\"      \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_DeleteDirPathAll_03(t *testing.T) {

  targetPath := "../createFilesTest/pathDoesNotExist/fileDoesNotExist.txt"

  err := FileHelper{}.DeleteDirPathAll(targetPath)

  if err != nil {
    t.Errorf("An error was NOT expected from FileHelper{}.DeleteDirPathAll(targetPath) "+
      "because the input parameter targetPath does not exist. "+
      "However, AN ERROR WAS RETURNED! targetPath='%v' Error='%v' ",
      targetPath, err.Error())
  }

}

func TestFileHelper_DeleteDirPathAll_04(t *testing.T) {

  fh := FileHelper{}

  basePath := "../createFilesTest/TestFileHelper_DeleteDirPathAll_04"
  basePath = fh.AdjustPathSlash(basePath)

  targetPath := "../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01"

  targetPath = fh.AdjustPathSlash(targetPath)

  err := os.RemoveAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: os.RemoveAll(targetPath) FAILED!\n" +
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())
    return
  }

  err = fh.MakeDirAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: fh.MakeDirAll(targetPath) FAILED!\n" +
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())

    return
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(basePath).\n" +
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
  }

}

func TestFileHelper_DeleteDirPathAll_05(t *testing.T) {

  fh := FileHelper{}

  basePath := "../createFilesTest/TestFileHelper_DeleteDirPathAll_04"
  basePath = fh.AdjustPathSlash(basePath)

  targetPath := "../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01"
  sourcePath := "../filesfortest/levelfilesfortest"

  testFile := "../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01/level_0_3_test.txt"
  testFile = fh.AdjustPathSlash(testFile)

  srcAry := make([]string, 0, 50)
  destAry := make([]string, 0, 50)

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath + "/" + "level_0_0_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash( targetPath + "/" + "level_0_0_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath + "/" + "level_0_1_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash( targetPath + "/" + "level_0_1_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath + "/" + "level_0_2_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash( targetPath + "/" + "level_0_2_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath + "/" + "level_0_3_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash( targetPath + "/" + "level_0_3_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath + "/" + "level_0_4_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash( targetPath + "/" + "level_0_4_test.txt"))

  targetPath = fh.AdjustPathSlash(targetPath)


  targetPath = fh.AdjustPathSlash(targetPath)

  err := os.RemoveAll(basePath)

  if err != nil {
    t.Errorf("Test Setup Error: os.RemoveAll(basePath) FAILED!\n" +
      "basePath='%v'\nError='%v'\n",
      basePath, err.Error())
    return
  }

  err = fh.MakeDirAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: fh.MakeDirAll(targetPath) FAILED!\n" +
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())

    return
  }


  for i:=0; i < len(srcAry); i++ {

    err := fh.CopyFileByIo(srcAry[i],destAry[i])

    if err != nil {
      t.Errorf("Error returned by fh.CopyFileByIo(srcAry[%v],destAry[%v]). \n" +
        "srcAry='%v'\ndestAry='%v'\nError='%v'", i, i, srcAry[i], destAry[i], err.Error())

      err = os.RemoveAll(basePath)

      t.Errorf("Secondary Clean-Up Error returned by os.RemoveAll(basePath).\n" +
        "basePath='%v'\nError='%v'\n", basePath, err.Error())

      return
    }

  }

  if !fh.DoesFileExist(testFile) {
    t.Errorf("Test Setup FAILED! Test Files DO NOT EXIST!\n" +
      "testfile='%v'", testFile)

    err = os.RemoveAll(basePath)

    t.Errorf("Secondary Clean-Up Error returned by os.RemoveAll(basePath).\n" +
      "basePath='%v'\nError='%v'\n", basePath, err.Error())

    return

  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(basePath).\n" +
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
  }

}

func TestFileHelper_DoesFileExist_01(t *testing.T) {

  fh := FileHelper{}

  doesFileExist := fh.DoesFileExist("")

  if doesFileExist {
    t.Error("Expected doesFileExist='false' because input parameter " +
      "for fh.DoesFileExist(\"\") is an " +
      "empty string. However, doesFileExist='true'!")
  }

}

func TestFileHelper_DoesFileExist_02(t *testing.T) {

  fh := FileHelper{}

  doesFileExist := fh.DoesFileExist("   ")

  if doesFileExist {
    t.Error("Expected doesFileExist='false' because input parameter " +
      "for fh.DoesFileExist(\"  \") consists entirely of blank spaces. " +
      "However, doesFileExist='true'!")
  }

}

func TestFileHelper_DoesFileInfoExist_01(t *testing.T) {
  fh := FileHelper{}

  doesFileExist, fInfo, err := fh.DoesFileInfoExist("")

  if err == nil {
    t.Error("Expected error from fh.DoesFileInfoExist(\"\") because " +
      "input parameter is an empty string. However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist != false {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") is an empty string. " +
      "However, doesFileExist=='true'!")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") is an empty string. " +
      "However, fInfo is NOT 'nil'!")
  }

}

func TestFileHelper_DoesFileInfoExist_02(t *testing.T) {
  fh := FileHelper{}

  doesFileExist, fInfo, err := fh.DoesFileInfoExist("   ")

  if err == nil {
    t.Error("Expected error from fh.DoesFileInfoExist(\"    \") because " +
      "input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist != false {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(\"\") consists entirely of blank spaces. " +
      "However, doesFileExist=='true'!")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(\"    \") consists entirely of blank spaces. " +
      "However, fInfo is NOT 'nil'!")
  }

}

func TestFileHelper_DoesFileInfoExist_03(t *testing.T) {

  fh := FileHelper{}

  testFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  expectedFileName := "level_0_1_test.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DoesFileInfoExist(testFile). "+
      "testFile='%v' Error='%v' ", testFile, err.Error())
  }

  if doesFileExist == false {
    t.Error("Expected doesFileExist=='true' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) actually exists. " +
      "However, doesFileExist=='false' !")
  }

  if fInfo == nil {
    t.Error("Expected fInfo!='nil' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) actually exists. " +
      "However, fInfo IS 'nil'!")
  }

  actualFileName := strings.ToLower(fInfo.Name())

  if expectedFileName != actualFileName {
    t.Errorf("Expected actual file name='%v'. Instead, actual file name='%v'.",
      expectedFileName, actualFileName)
  }

}

func TestFileHelper_DoesFileInfoExist_04(t *testing.T) {

  fh := FileHelper{}

  testFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err == nil {
    t.Error("Expected an error return from fh.DoesFileInfoExist(testFile). " +
      "because 'testFile' does NOT exist. However, NO ERROR WAS RETURNED!")
  }

  if doesFileExist == true {
    t.Error("Expected doesFileExist=='false' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) DOES NOT EXIST. " +
      "However, doesFileExist=='true' !")
  }

  if fInfo != nil {
    t.Error("Expected fInfo=='nil' because input parameter for " +
      "fh.DoesFileInfoExist(testFile) DOES NOT EXIST. " +
      "However, fInfo IS NOT 'nil'!")
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_01(t *testing.T) {
  rawtestStr := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if !doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'false'. Expected a return value of 'true' because testStr ends "+
      "with path separator.  testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_02(t *testing.T) {

  rawtestStr := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'true'. Expected a return value of 'false' because testStr does NOT "+
      "end with a path separator. testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_03(t *testing.T) {

  rawtestStr := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  fh := FileHelper{}

  testStr := fh.AdjustPathSlash(rawtestStr)

  doesEndWithSep := fh.DoesStringEndWithPathSeparator(testStr)

  if doesEndWithSep {
    t.Errorf("Error: fh.DoesStringEndWithPathSeparator(testStr) returned "+
      "'true'. Expected a return value of 'false' because testStr does NOT "+
      "end with a path separator. testStr='%v'", testStr)
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_04(t *testing.T) {

  fh := FileHelper{}
  doesEndWithSep := fh.DoesStringEndWithPathSeparator("")

  if doesEndWithSep {
    t.Error("Error: fh.DoesStringEndWithPathSeparator(\"\") returned " +
      "'true'. Expected a return value of 'false' because input parameter " +
      "is an empty string. ")
  }

}

func TestFileHelper_DoesStringEndWithPathSeparator_05(t *testing.T) {

  fh := FileHelper{}
  doesEndWithSep := fh.DoesStringEndWithPathSeparator("    ")

  if doesEndWithSep {
    t.Error("Error: fh.DoesStringEndWithPathSeparator(\"   \") returned " +
      "'true'. Expected a return value of 'false' because input parameter " +
      "consists entirely of blank spaces. ")
  }

}

func TestFileHelper_FindFilesInPath_01(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesDirMgrAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOperationCode(0).CopySourceToDestinationByIo()

  errStrs := sourceDir.ExecuteDirectoryFileOps(fileSelect, fileOps, targetDir)

  if len(errStrs) > 0 {
    for i := 0; i < len(errStrs); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryFileOps-Error: %v", errStrs[i])
    }
  }

  foundFiles, err := fh.FindFilesInPath(targetDir.GetAbsolutePath(), "*.*")

  lenFoundFiles := len(foundFiles)

  if lenFoundFiles != 5 {
    t.Errorf("Error: Expected to find 5-files. Instead, found %v-files! ",
      lenFoundFiles)
  }

  _ = targetDir.DeleteAll()

}

func TestFileHelper_FindFilesInPath_02(t *testing.T) {

  fh := FileHelper{}

  targetDirStr, err := fh.MakeAbsolutePath("../dirmgrtests/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"../dirmgrtests/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  sourceDirStr, err := fh.MakeAbsolutePath("../filesfortest/levelfilesfortest")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath("+
      "\"..../filesfortest/levelfilesfortest \") "+
      "Error='%v' ", err.Error())
  }

  targetDir, err := DirMgr{}.New(targetDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDirStr) "+
      "targetDirStr='%v' Error='%v' ", targetDirStr, err.Error())
  }

  sourceDir, err := DirMgr{}.New(sourceDirStr)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir) "+
      "sourceDir='%v' Error='%v' ", sourceDir, err.Error())
  }

  if targetDir.DoesDirMgrAbsolutePathExist() {

    err = targetDir.DeleteAll()

    if err != nil {
      t.Errorf("Error returned by targetDir.DeleteAll() "+
        "targetDir='%v' Error='%v' ",
        targetDir.GetAbsolutePath(), err.Error())
    }
  }

  // Target Directory does NOT Exist

  fileSelect := FileSelectionCriteria{}

  fileSelect.SelectCriterionMode = FileSelectMode.ORSelect()

  fileOps := make([]FileOperationCode, 1, 5)

  fileOps[0] = FileOperationCode(0).CopySourceToDestinationByIo()

  errStrs := sourceDir.ExecuteDirectoryTreeOps(fileSelect, fileOps, targetDir)

  if len(errStrs) > 0 {
    for i := 0; i < len(errStrs); i++ {
      t.Errorf("sourceDir.ExecuteDirectoryTreeOps-Error: %v", errStrs[i])
    }
  }

  foundFiles, err := fh.FindFilesInPath(targetDir.GetAbsolutePath(), "*")

  lenFoundFiles := len(foundFiles)

  if lenFoundFiles != 6 {
    t.Errorf("Error: Expected to find 6-files. Instead, found %v-files! ",
      lenFoundFiles)
  }

  _ = targetDir.DeleteAll()

}


func TestFileHelper_FindFilesInPath_03(t *testing.T) {
  fh := FileHelper{}

  foundFiles, err := fh.FindFilesInPath("", "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(\"\", \"*.*\") " +
      "because first input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(\"\", \"*.*\") would be zero length because "+
      "the first input parameter is an empty string."+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_04(t *testing.T) {
  fh := FileHelper{}

  foundFiles, err := fh.FindFilesInPath("   ", "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(\"   \", \"*.*\") " +
      "because first input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(\"    \", \"*.*\") would be zero length because "+
      "the first input parameter consists entirely of empty spaces. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_05(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"\") " +
      "because the second input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"\") would be zero length because "+
      "the second input parameter is an empty string."+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_06(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "    ")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"   \") " +
      "because the second input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"   \") would be zero length because "+
      "the second input parameter consists entirely of empty spaces. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FindFilesInPath_07(t *testing.T) {
  fh := FileHelper{}

  pathFileName := "../filesfortest/levelfilesfortest/level_01_dir/iDoNotExistDir"

  foundFiles, err := fh.FindFilesInPath(pathFileName, "*.*")

  if err == nil {
    t.Error("Expected error return from fh.FindFilesInPath(pathFileName, \"*.*\") " +
      "because input parameter 'pathFileName' DOES NOT EXIST. " +
      "However, NO ERROR WAS RETURNED!")
  }

  lFFiles := len(foundFiles)

  if lFFiles != 0 {
    t.Errorf("Expected that found files array returned from "+
      "fh.FindFilesInPath(pathFileName, \"*.*\") would be zero length because "+
      "the input parameter 'pathFileName' DOES NOT EXIST. "+
      "However, length of found files='%v' ", lFFiles)
  }

}

func TestFileHelper_FilterFileName_01(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_02(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
      "fmtstr='%v' fOlderThanStr='%v' Error='%v'",
      fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_03(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)
  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
      "fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
  }

  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_04(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  filesOlderThan := time.Time{}
  fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_05(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}
  fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("File was NOT found. File should have been found. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_06(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.txt"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("It was expected that this File would NOT be found. It WAS Found. "+
      "Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_07(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("It was expected that this file would NOT be Found. Instead, it WAS found. "+
      "Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_08(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := "*.htm"
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_09(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  //filesOlderThan := time.Time{}
  fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
  filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_10(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}

  fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
  filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

  if err != nil {
    t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
      "fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if isFound {
    t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }

}

func TestFileHelper_FilterFileName_11(t *testing.T) {

  fia := FileInfoPlus{}
  fia.SetName("newerFileForTest_01.txt")
  fia.SetMode(0777)
  fia.SetSize(107633)
  fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
  fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
  fModTime, err := time.Parse(fmtstr, fModTimeStr)

  if err != nil {
    t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
      "fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
  }

  fia.SetModTime(fModTime)
  fia.SetIsDir(false)
  fia.SetSysDataSrc(nil)
  fia.SetIsDir(true)

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  fh := FileHelper{}
  isFound, err := fh.FilterFileName(fia, fsc)

  if !isFound {
    t.Errorf("Expected that File would be found. However, File WAS NOT found - Error. "+
      "fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
  }
}

func TestFileHelper_FilterFileName_12(t *testing.T) {

  fh := FileHelper{}
  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{""}
  fsc.FilesOlderThan = time.Time{}
  fsc.FilesNewerThan = time.Time{}
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  isFound, err := fh.FilterFileName(nil, fsc)

  if err == nil {
    t.Error("Expected an error return from fh.FilterFileName(nil, fsc) because " +
      "the first input parameter is 'nil'. " +
      "However, NO ERROR WAS RETURNED!")
  }

  if isFound {
    t.Error("Expected isFound=='false'. Instead, isFound=='true'. ")
  }
}
