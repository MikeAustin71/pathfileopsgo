package pathfileops

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "testing"
)

func TestFileHelper_CopyFileByLink_01(t *testing.T) {

  testDestFile := "../../filesfortest/levelfilesfortest/level_9_9_test.txt"

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

  testDestFile := "../../filesfortest/levelfilesfortest/level_9_9_test.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_9_9_test.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_9_9_test.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

  srcFile, err := fh.MakeAbsolutePath(testSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(testSrcFile). "+
      "testSrcFile='%v' Error='%v' ", testSrcFile, err.Error())
  }

  testDestFile := "../../checkfiles/scratchDestTileJJ459821.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_0_4_test.txt"

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

  testDestFile := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

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

  testDestFile := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

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

  srcFile := "../../filesfortest/levelfilesfortest/level_0_1_test.txt"

  destFile := "../../createFilesTest/scratchXFg924198.txt"

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

  testSrcFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_0_3_test.txt")

  srcFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_0_1_test.txt")

  destFile :=
    fh.AdjustPathSlash("../../createFilesTest/scratchTestFileHelper_CopyFileByLink_12.txt")

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
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: The copy by link operation failed! "+
      "The destination file does NOT exist.\n"+
      "destFile='%v'\n",
      destFile)
    return
  }

  // Destination file exists
  finfoSrc, err := os.Stat(srcFile)

  if err != nil {
    t.Errorf("Error returned by os.Stat(srcFile).\n"+
      "srcFile='%v'\nError='%v'", srcFile, err.Error())
  }

  finfoDest, err := os.Stat(destFile)

  if err != nil {
    t.Errorf("Error returned by os.Stat(destFile).\n"+
      "destFile='%v'\nError='%v'", destFile, err.Error())
    return
  }

  if finfoSrc.Size() != finfoDest.Size() {
    t.Errorf("Error: Size of source file does NOT match "+
      "size of destination file.\nSourceFileSize='%v' DestinationFileSize='%v'\n",
      finfoSrc.Size(), finfoDest.Size())
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

  testSrcFile := "../../filesfortest/levelfilesfortest/level_0_3_test.txt"

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

  srcFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/iDoNotExist.txt")

  destFile :=
    fh.AdjustPathSlash("../../createFilesTest/scratchTestFileHelper_CopyFileByLink_14.txt")

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
  fh := FileHelper{}

  tstFile := fh.AdjustPathSlash("../../checkfiles/TestOverwrite001.txt")

  if fh.DoesFileExist(tstFile) {
    err := fh.DeleteDirFile(tstFile)
    if err != nil {
      t.Error(fmt.Sprintf("Error: Deletion Failed On File %v !", tstFile))
    }
  }

  if fh.DoesFileExist(tstFile) {
    t.Error(fmt.Sprintf("Error: Deletion Failed!\n"+
      "File %v should not exist!",
      tstFile))
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  f, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Error(fmt.Sprintf("Error: Create File Failed for file: %v", tstFile))
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  _, err4 := f.WriteString(aLoremIpsumTxt)

  if err4 != nil {
    _ = f.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err4)
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  err = f.Close()

  if err != nil {
    t.Errorf("%v", err.Error())
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  // Now recreate the original file. It should be
  // truncated with the old contents deleted.
  f2, err := fh.CreateFile(tstFile)

  if err != nil {
    t.Errorf("Error: Re-Creating File %v", tstFile)
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  fOvrWriteTxt := "Test Over Write and existing file using Create()"

  _, err5 := f2.WriteString(fOvrWriteTxt)

  if err5 != nil {
    _ = f2.Close()
    t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err5)
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  err = f2.Close()

  if err != nil {
    t.Errorf("Error closing f2. %v", err.Error())
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  dat, err := ioutil.ReadFile(tstFile)

  if err != nil {
    t.Errorf("Error Reading Re-Written Text for File:'%v' Error='%v'", tstFile, err)
    _ = fh.DeleteDirFile(tstFile)
    return
  }

  s := string(dat)

  if s != fOvrWriteTxt {
    t.Errorf("Was expecting to read text: '%v', instead received text: %v", fOvrWriteTxt, s)
  }

  _ = fh.DeleteDirFile(tstFile)

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

  testFilePath := "../../checkfiles/TestFileHelper_CreateFile_05/dirX1/dirX2/scratchTestFileHelper_CreateFile_05.txt"

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

  testFilePath := "../../checkfiles/scratchTestFileHelper_CreateFile_05.txt"

  fh := FileHelper{}

  testFilePath = fh.AdjustPathSlash(testFilePath)

  fPtr, err := fh.CreateFile(testFilePath)

  if err != nil {
    t.Errorf("Error returned from fh.CreateFile(testFilePath).\n"+
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
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(testFilePath).\n"+
      "testFilePath='%v'\nError='%v'\n", testFilePath, err.Error())
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

  targetPath := "../../createFilesTest/pathDoesNotExist/fileDoesNotExist.txt"

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

  basePath := "../../createFilesTest/TestFileHelper_DeleteDirPathAll_04"
  basePath = fh.AdjustPathSlash(basePath)

  targetPath := "../../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01"

  targetPath = fh.AdjustPathSlash(targetPath)

  err := os.RemoveAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: os.RemoveAll(targetPath) FAILED!\n"+
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())
    return
  }

  err = fh.MakeDirAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: fh.MakeDirAll(targetPath) FAILED!\n"+
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())

    return
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(basePath).\n"+
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
  }

}

func TestFileHelper_DeleteDirPathAll_05(t *testing.T) {

  fh := FileHelper{}

  basePath := "../../createFilesTest/TestFileHelper_DeleteDirPathAll_04"
  basePath = fh.AdjustPathSlash(basePath)

  targetPath := "../../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01"
  sourcePath := "../../filesfortest/levelfilesfortest"

  testFile := "../../createFilesTest/TestFileHelper_DeleteDirPathAll_04/dirx01/level_0_3_test.txt"
  testFile = fh.AdjustPathSlash(testFile)

  srcAry := make([]string, 0, 50)
  destAry := make([]string, 0, 50)

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath+"/"+"level_0_0_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash(targetPath+"/"+"level_0_0_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath+"/"+"level_0_1_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash(targetPath+"/"+"level_0_1_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath+"/"+"level_0_2_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash(targetPath+"/"+"level_0_2_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath+"/"+"level_0_3_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash(targetPath+"/"+"level_0_3_test.txt"))

  srcAry = append(srcAry,
    fh.AdjustPathSlash(sourcePath+"/"+"level_0_4_test.txt"))

  destAry = append(destAry,
    fh.AdjustPathSlash(targetPath+"/"+"level_0_4_test.txt"))

  targetPath = fh.AdjustPathSlash(targetPath)

  targetPath = fh.AdjustPathSlash(targetPath)

  err := os.RemoveAll(basePath)

  if err != nil {
    t.Errorf("Test Setup Error: os.RemoveAll(basePath) FAILED!\n"+
      "basePath='%v'\nError='%v'\n",
      basePath, err.Error())
    return
  }

  err = fh.MakeDirAll(targetPath)

  if err != nil {
    t.Errorf("Test Setup Error: fh.MakeDirAll(targetPath) FAILED!\n"+
      "targetPath='%v'\nError='%v'\n",
      targetPath, err.Error())

    return
  }

  for i := 0; i < len(srcAry); i++ {

    err := fh.CopyFileByIo(srcAry[i], destAry[i])

    if err != nil {
      t.Errorf("Error returned by fh.CopyFileByIo(srcAry[%v],destAry[%v]). \n"+
        "srcAry='%v'\ndestAry='%v'\nError='%v'", i, i, srcAry[i], destAry[i], err.Error())

      err = os.RemoveAll(basePath)

      if err != nil {

        t.Errorf("Secondary Clean-Up Error returned by os.RemoveAll(basePath).\n"+
          "basePath='%v'\nError='%v'\n", basePath, err.Error())
      }

      return
    }

  }

  if !fh.DoesFileExist(testFile) {
    t.Errorf("Test Setup FAILED! Test Files DO NOT EXIST!\n"+
      "testfile='%v'", testFile)

    err = os.RemoveAll(basePath)

    if err != nil {
      t.Errorf("Secondary Clean-Up Error returned by os.RemoveAll(basePath).\n"+
        "basePath='%v'\nError='%v'\n", basePath, err.Error())
    }

    return
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirPathAll(basePath).\n"+
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

  testFile := "../../filesfortest/levelfilesfortest/level_0_1_test.txt"

  expectedFileName := "level_0_1_test.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DoesFileInfoExist(testFile). "+
      "testFile='%v' Error='%v' ", testFile, err.Error())
    return
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

  testFile := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

  doesFileExist, fInfo, err := fh.DoesFileInfoExist(testFile)

  if err != nil {
    t.Errorf("Error return from fh.DoesFileInfoExist(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
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

func TestFileHelper_DeleteFilesWalkDirectory_01(t *testing.T) {

  originalSetupDir := "../../filesfortest"

  dMgrOrigSetup, err := DirMgr{}.New(originalSetupDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(originalSetupDir)\n" +
      "originalSetupDir='%v'\nError='%v'\n",
      originalSetupDir, err.Error())
    return
  }

  targetDir := "../../dirwalkdeletetests/filesfortest"
  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by targetDMgr.DeleteAll() ")
  }

  fsc := FileSelectionCriteria{}

  _, errs :=
    dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)

  if len(errs) > 0 {
    t.Errorf("Error returned by dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), FileHelper{}.ConsolidateErrors(errs))
    _ = targetDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.htm"}

  dTreeInfo, errs := targetDMgr.FindDirectoryTreeFiles(fsc)

  if len(errs) > 0 {
    t.Errorf("Error returned by dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), FileHelper{}.ConsolidateErrors(errs))
    _ = targetDMgr.DeleteAll()
    return
  }

  numOfHtmFiles := dTreeInfo.FoundFiles.GetNumOfFiles()

  fsc = FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.htm"}

  deleteInfo, err := FileHelper{}.DeleteFilesWalkDirectory(targetDMgr.GetAbsolutePath(), fsc)

  if err != nil {
   t.Errorf("Error returned by FileHelper{}.DeleteFilesWalkDirectory(targetDMgr.GetAbsolutePath(), fsc).\n" +
     "targetDMgr.GetAbsolutePath()='%v'\nError='%v'\n",
     targetDMgr.GetAbsolutePath(), err.Error())
   _ = targetDMgr.DeleteAll()
   return
  }

  if len(deleteInfo.ErrReturns) > 0 {
    t.Errorf("FileHelper{}.DeleteFilesWalkDirectory yielded deleteInfo.ErrReturns as follows:\n'%v'\n",
      FileHelper{}.ConsolidateErrors(deleteInfo.ErrReturns))
    _ = targetDMgr.DeleteAll()
    return
  }

  numOfDeletedFiles := deleteInfo.DeletedFiles.GetNumOfFiles()

  if numOfHtmFiles != numOfDeletedFiles {
    t.Errorf("Expected number of deleted 'htm' files='%v'\n" +
      "Instead, number of deleted 'htm' files='%v'\n",
      numOfHtmFiles, numOfDeletedFiles)
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), err.Error())
  }

}

func TestFileHelper_DeleteFilesWalkDirectory_02(t *testing.T) {

  originalSetupDir := "../../filesfortest"

  dMgrOrigSetup, err := DirMgr{}.New(originalSetupDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(originalSetupDir)\n" +
      "originalSetupDir='%v'\nError='%v'\n",
      originalSetupDir, err.Error())
    return
  }

  targetDir := "../../dirwalkdeletetests/filesfortest"
  targetDMgr, err := DirMgr{}.New(targetDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(targetDir)\n" +
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned by targetDMgr.DeleteAll() ")
  }

  fsc := FileSelectionCriteria{}

  _, errs :=
    dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)

  if len(errs) > 0 {
    t.Errorf("Error returned by dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), FileHelper{}.ConsolidateErrors(errs))
    _ = targetDMgr.DeleteAll()
    return
  }

  fsc = FileSelectionCriteria{}

  dTreeInfo, errs := targetDMgr.FindDirectoryTreeFiles(fsc)

  if len(errs) > 0 {
    t.Errorf("Error returned by dMgrOrigSetup.CopySubDirectoryTree(targetDMgr,false, fsc)\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), FileHelper{}.ConsolidateErrors(errs))
    _ = targetDMgr.DeleteAll()
    return
  }

  numOfTotalFiles := dTreeInfo.FoundFiles.GetNumOfFiles()

  fsc = FileSelectionCriteria{}

  deleteInfo, err := FileHelper{}.DeleteFilesWalkDirectory(targetDMgr.GetAbsolutePath(), fsc)

  if err != nil {
   t.Errorf("Error returned by FileHelper{}.DeleteFilesWalkDirectory(targetDMgr.GetAbsolutePath(), fsc).\n" +
     "targetDMgr.GetAbsolutePath()='%v'\nError='%v'\n",
     targetDMgr.GetAbsolutePath(), err.Error())
   _ = targetDMgr.DeleteAll()
   return
  }

  if len(deleteInfo.ErrReturns) > 0 {
    t.Errorf("FileHelper{}.DeleteFilesWalkDirectory yielded deleteInfo.ErrReturns as follows:\n'%v'\n",
      FileHelper{}.ConsolidateErrors(deleteInfo.ErrReturns))
    _ = targetDMgr.DeleteAll()
    return
  }

  numOfDeletedFiles := deleteInfo.DeletedFiles.GetNumOfFiles()

  if numOfTotalFiles != numOfDeletedFiles {
    t.Errorf("Expected number of deleted 'htm' files='%v'\n" +
      "Instead, number of deleted 'htm' files='%v'\n",
      numOfTotalFiles, numOfDeletedFiles)
  }

  err = targetDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n" +
      "targetDMgr='%v'\nError='%v'\n",
      targetDMgr.GetAbsolutePath(), err.Error())
  }

}

func TestFileHelper_DeleteFilesWalkDirectory_03(t *testing.T) {

  fsc := FileSelectionCriteria{}

  target := ""

  _, err := FileHelper{}.DeleteFilesWalkDirectory(target, fsc)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}." +
      "DeleteFilesWalkDirectory(target, fsc)\n" +
      "because 'target' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}



func TestFileHelper_DoesStringEndWithPathSeparator_01(t *testing.T) {
  rawtestStr := "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/"

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

  rawtestStr := "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir"

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

  rawtestStr := "../../filesfortest/levelfilesfortest/level_0_1_test.txt"

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

func TestFileHelper_DoesThisFileExist_01(t *testing.T) {
  fh := FileHelper{}

  testDirStr := "../../checkfiles/iDoNotExist.txt"

  pathFileDoesExist, err := fh.DoesThisFileExist(testDirStr)

  if err != nil {
    t.Errorf("Error: Non-Path Error returned from fh.DoesThisFileExist(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  if pathFileDoesExist {
    t.Errorf("Error: Expected result from exitence test = 'File Does NOT Exist!\n"+
      "Instead, existence test= 'File DOES Exist!\n"+
      "testDirStr='%v'", testDirStr)
  }

}

func TestFileHelper_DoesThisFileExist_02(t *testing.T) {
  fh := FileHelper{}

  testDirStr := ""

  _, err := fh.DoesThisFileExist(testDirStr)

  if err == nil {
    t.Error("Expected an error return from fh.DoesThisFileExist(testDirStr)\n" +
      "because input parameter 'testDirStr' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }
}

func TestFileHelper_DoesThisFileExist_03(t *testing.T) {
  fh := FileHelper{}

  testDirStr := "   "

  _, err := fh.DoesThisFileExist(testDirStr)

  if err == nil {
    t.Error("Expected an error return from fh.DoesThisFileExist(testDirStr)\n" +
      "because input parameter 'testDirStr' consists entirely of empty spaces!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }
}

func TestFileHelper_DoesThisFileExist_04(t *testing.T) {
  fh := FileHelper{}

  testDirStr := "../../filesfortest/htmlFilesForTest/006860_sample.htm"

  pathFileDoesExist, err := fh.DoesThisFileExist(testDirStr)

  if err != nil {
    t.Errorf("Error: Error returned from fh.DoesThisFileExist(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  if !pathFileDoesExist {
    t.Errorf("Error: Expected result from exitence test = 'File DOES Exist!\n"+
      "Instead, existence test= 'File Does NOT Exist!\n"+
      "testDirStr='%v'", testDirStr)
  }

}

func TestFileHelper_DoesThisFileExist_05(t *testing.T) {
  fh := FileHelper{}

  testDirStr := "../../filesfortest/htmlFilesForTest"

  pathFileDoesExist, err := fh.DoesThisFileExist(testDirStr)

  if err != nil {
    t.Errorf("Error: Error returned from fh.DoesThisFileExist(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  if !pathFileDoesExist {
    t.Errorf("Error: Expected result from exitence test = 'Directory DOES Exist!\n"+
      "Instead, existence test= 'Directory Does NOT Exist!\n"+
      "testDirStr='%v'", testDirStr)
  }

}

func TestFileHelper_DoesThisFileExist_06(t *testing.T) {
  fh := FileHelper{}

  testDirStr := "filesfortest/iDoNotExist"

  pathFileDoesExist, err := fh.DoesThisFileExist(testDirStr)

  if err != nil {
    t.Errorf("Error: Error returned from fh.DoesThisFileExist(testDirStr)\n"+
      "testDirStr='%v'\nError='%v'\n", testDirStr, err.Error())
    return
  }

  if pathFileDoesExist {
    t.Errorf("Error: Expected result from exitence test = 'Directory DOES NOT Exist!\n"+
      "Instead, existence test= 'Directory DOES Exist!\n"+
      "testDirStr='%v'", testDirStr)
  }
}

