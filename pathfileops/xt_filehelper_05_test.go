package pathfileops

import (
  appLib "MikeAustin71/pathfileopsgo/appLibs"
  "errors"
  "io"
  "os"
  "testing"
  "time"
)

func TestFileHelper_MakeAbsolutePath_01(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.MakeAbsolutePath("")

  if err == nil {
    t.Error("Expected an error return from fh.MakeAbsolutePath(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeAbsolutePath_02(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.MakeAbsolutePath("   ")

  if err == nil {
    t.Error("Expected an error return from fh.MakeAbsolutePath(\"\") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeDirAll_01(t *testing.T) {
  fh := FileHelper{}

  err := fh.MakeDirAll("")

  if err == nil {
    t.Error("Expected an error return from fh.MakeDirAll(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_MakeDirAll_02(t *testing.T) {
  fh := FileHelper{}

  err := fh.MakeDirAll("  ")

  if err == nil {
    t.Error("Expected an error return from fh.MakeDirAll(\"    \") " +
      "because the input parameter consists entirely of blank spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFileHelper_MakeDirAll_03(t *testing.T) {

  fh := FileHelper{}
  baseDirPath := fh.AdjustPathSlash("../checkfiles/TestFileHelper_MakeDirAll_03")
  dirPath := fh.AdjustPathSlash("../checkfiles/TestFileHelper_MakeDirAll_03/tdir01/tdir02/tdir03")

  _, err := os.Stat(baseDirPath)

  if err == nil {
    err = fh.DeleteDirPathAll(baseDirPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test setup.\n" +
        "Attempted deletion of base directory FAILED!\nbaseDirPath='%v'\n",
        baseDirPath)
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: Attempted deletion of baseDirPath during test setup FAILED!\n" +
        "baseDirPath still exists!\nbaseDirPath='%v'\n", baseDirPath)
    }
  }

  err = fh.MakeDirAll(dirPath)

  if err !=nil {
    t.Errorf("Error returned by fh.MakeDirAll(dirPath).\n" +
      "dirPath='%v'\nError='%v'", dirPath, err.Error())
    return
  }

  _, err = os.Stat(dirPath)

  if err != nil {
    t.Errorf("Error: fh.MakeDirAll(dirPath) FAILED!\n" +
      "os.Stat() confirms that dirPath DOES NOT EXIST!\n" +
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.DeleteDirPathAll(baseDirPath)

  if err != nil{
    t.Errorf("Error returned during clean-up by fh.DeleteDirPathAll(baseDirPath).\n" +
      "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
  }

  _, err = os.Stat(baseDirPath)

  if err == nil {
    t.Errorf("ERROR: baseDirPath still EXISTS! Attempted deletion FAILED!\n" +
      "baseDirPath='%v'\n", baseDirPath)
  }

}

func TestFileHelper_MakeDirAllPerm01(t *testing.T) {

  fh := FileHelper{}
  baseDirPath := fh.AdjustPathSlash("../checkfiles/TestFileHelper_MakeDirAllPerm01")
  dirPath := fh.AdjustPathSlash(
    "../checkfiles/TestFileHelper_MakeDirAllPerm01/tDir1/tDir2/tDir3/tDir4")

  _, err := os.Stat(baseDirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(baseDirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test startup.\n" +
        "This means that the baseDirPath and all subsidiary directories could NOT be deleted!\n" +
        "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: During test startup attempts to delete the test baseDirPath FAILED!\n" +
        "baseDirPath still EXISTS!\nbaseDirPath='%v'\n", baseDirPath)
      return
    }

  }

  permissionCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil{
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n" +
      "Error='%v'\n", err.Error())
  }

  permissionCfg.isInitialized = false

  err = fh.MakeDirAllPerm(dirPath, permissionCfg)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from fh.MakeDirAllPerm(dirPath, permissionCfg)\n" +
      "because 'permissionCfg' is Invalid. However, NO ERROR WAS RETURNED!\n" +
      "dirPath='%v'\n", dirPath)
  }

  _, err = os.Stat(baseDirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(baseDirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test clean-up.\n" +
        "This means that the baseDirPath and all subsidiary directories could NOT be deleted!\n" +
        "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: During test clean-up attempts to delete the test baseDirPath FAILED!\n" +
        "baseDirPath still EXISTS!\nbaseDirPath='%v'\n", baseDirPath)
    }

  }

}

func TestFileHelper_MakeDirPerm_01(t *testing.T) {

  fh := FileHelper{}
  dirPath := fh.AdjustPathSlash(
    "../checkfiles/TestFileHelper_MakeDirPerm_01")

  _, err := os.Stat(dirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(dirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(dirPath) during test startup.\n" +
        "This means that the dirPath could NOT be deleted!\n" +
        "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
      return
    }

    _, err = os.Stat(dirPath)

    if err == nil {
      t.Errorf("ERROR: During test startup attempts to delete the test dirPath FAILED!\n" +
        "dirPath still EXISTS!\nbaseDirPath='%v'\n", dirPath)
      return
    }

  }

  permissionCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil{
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  permissionCfg.isInitialized = false

  err = fh.MakeDirPerm(dirPath, permissionCfg)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from fh.MakeDirPerm(dirPath, permissionCfg)\n" +
      "because 'permissionCfg' is Invalid. However, NO ERROR WAS RETURNED!\n" +
      "dirPath='%v'\n", dirPath)
  }

  _, err = os.Stat(dirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(dirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(dirPath) during test clean-up.\n" +
        "This means that the dirPath could NOT be deleted!\n" +
        "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
      return
    }

    _, err = os.Stat(dirPath)

    if err == nil {
      t.Errorf("ERROR: During test clean-up attempts to delete the test dirPath FAILED!\n" +
        "dirPath still EXISTS!\ndirPath='%v'\n", dirPath)
    }

  }

}

func TestFileHelper_MakeDir01(t *testing.T) {
  fh := FileHelper{}
  dirPath := fh.AdjustPathSlash("../checkfiles/TestFileHelper_MakeDir01")

  _, err := os.Stat(dirPath)

  if err == nil  {

    err = fh.DeleteDirFile(dirPath)

    if err != nil {
      t.Errorf("Error during test setup. Directory already exists!\n" +
        "Attempted directory deletion FAILED!\ndirPath='%v'\n", dirPath)
      return
    }

  }

  _, err = os.Stat(dirPath)

  if err == nil {
    t.Errorf("ERROR: Setup tests directory still exists!\n" +
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.MakeDir(dirPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeDir(dirPath).\n" +
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  _, err = os.Stat(dirPath)

  if err != nil {
    t.Errorf("ERROR: MakeDir(dirPath) failed to create test directory.\n" +
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.DeleteDirFile(dirPath)

  if err != nil {
    t.Errorf("Error returned by cleanup fh.DeleteDirFile(dirPath).\n" +
      "dirPath='%v'\n", dirPath)
    return
  }

  _, err = os.Stat(dirPath)

  if err == nil {
    t.Errorf("Error: Clean-up FAILED! Test directory still exists!\n" +
      "dirPath='%v'\n", dirPath)
  }

}

func TestFileHelper_MakeDir02(t *testing.T) {

  fh := FileHelper{}
  dirPath := ""
 err := fh.MakeDir(dirPath)

  if err == nil {
    t.Error("Expected an error return from fh.MakeDir(dirPath) because\n" +
      "'dirPath' is an empty string. However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_MakeDir03(t *testing.T) {

  fh := FileHelper{}
  dirPath := "    "

  err := fh.MakeDir(dirPath)

  if err == nil {
    t.Error("Expected an error return from fh.MakeDir(dirPath) because\n" +
      "'dirPath' consists entirely of blank spaces. However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_MoveFile_01(t *testing.T) {
  fh := FileHelper{}
  setupFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../logTest/FileSrc/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../logTest/scratchTestFileHelper_MoveFile_01.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error during test setup deleting destination file.\n" +
        "Destination File='%v'\nError:'%v'\n",
        destFile, err)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error on test setup: destination file, STILL EXISTS!\n" +
        "Destination File='%v'", destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying 'setupFile' to 'srcFile'.\n" +
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n" +
      "setupFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupFile, srcFile, err.Error())
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup FAILED! Source File does NOT EXIST!!\n" +
      "srcFile='%v'", srcFile)
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returend by fh.MoveFile(srcFile, destFile)\n" +
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
    return
  }

  if fh.DoesFileExist(srcFile) {
    t.Errorf("FileHelper:MoveFile() FAILED! Source File still exists!!\n" +
      "Source File='%v'\n", srcFile)
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("FileHelper:MoveFile() FAILED! Destination File DOES NOT EXIST!\n" +
      "Destination File='%v'\n", destFile)
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err !=nil {
    t.Errorf("Error during test clean-up: Attempted deletion of destination " +
      "file FAILED!\nDestination File still exists!\nDestination File='%v'",
      destFile)
  }
}

func TestFileHelper_MoveFile_02(t *testing.T) {
  fh := FileHelper{}
  srcFile := ""
  destFile := fh.AdjustPathSlash("../logTest/scratchTestFileHelper_MoveFile_02.txt")

  err := fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because srcFile is an empty string. However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_MoveFile_03(t *testing.T) {
  fh := FileHelper{}
  srcFile := "   "
  destFile := fh.AdjustPathSlash("../logTest/scratchTestFileHelper_MoveFile_02.txt")

  err := fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because srcFile consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_MoveFile_04(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  destFile := "    "

  err := fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because destFile consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  _, err = os.Stat(srcFile)

  if err != nil {
    t.Errorf("Error: Expected that source file would NOT be deleted with an error\n" +
      "return from MoveFile(). However, the source file WAS DELETED!\nSource File='%v'\n",
      srcFile)
  }

}

func TestFileHelper_MoveFile_05(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/TestFile003.txt")
  destFile := ""

  err := fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because 'destFile' is an empty string.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  _, err = os.Stat(srcFile)

  if err != nil {
    t.Errorf("Error: Expected that source file would NOT be deleted with an error\n" +
      "return from MoveFile(). However, the source file WAS DELETED!\nSource File='%v'\n",
      srcFile)
  }
}

func TestFileHelper_MoveFile_06(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("../logTest/FileMgmnt/iDoNotExist.txt")
  destFile := "../logTest/FileMgmnt/scratchTestFileHelper_MoveFile_06.txt"

  err := fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because 'srcFile' DOES NOT EXIST!.\nHowever, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_MoveFile_07(t *testing.T) {
  fh := FileHelper{}
  setupDestFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_0_3_test.txt")
  setupSrcFile := fh.AdjustPathSlash("../checkfiles/setuplevel_0_3_test.txt")
  destFile := fh.AdjustPathSlash("../logTest//scratchTestFileHelper_MoveFile_07.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error during test setup deleting destination file.\n" +
        "Destination File='%v'\nError:'%v'\n",
        destFile, err)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error on test setup: destination file, STILL EXISTS!\n" +
        "Destination File='%v'", destFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupDestFile, destFile)

  if err != nil {
    t.Errorf("Received error copying 'setupDestFile' to 'srcFile'.\n" +
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n" +
      "setupDestFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupDestFile, srcFile, err.Error())
    return
  }


  err = fh.CopyFileByIo(srcFile, setupSrcFile)

  if err != nil {
    t.Errorf("Error returned from fh.CopyFileByIo(srcFile, setupSrcFile).\n" +
      "Test Setup for source file FAILED!\n" +
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      srcFile, setupSrcFile, err.Error())
    return
  }

  if !fh.DoesFileExist(setupSrcFile) {
    t.Errorf("Test Setup FAILED! Setup Source File does NOT EXIST!!\n" +
      "setupSrcFile='%v'", setupSrcFile)
    return
  }

  err = fh.MoveFile(setupSrcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.MoveFile(setupSrcFile, destFile)\n" +
      "setupSrcFile='%v'\ndestFile='%v'\nError='%v'\n",
      setupSrcFile, destFile, err.Error())
    return
  }

  destFileInfo, err := os.Stat(destFile)

  if err != nil {
    t.Errorf("Error: MoveFile() did NOT create the destFile!\n" +
      "Destination File='%v'\n", destFile)
  }

  srcFileInfo, err := os.Stat(srcFile)

  if srcFileInfo.Size() != destFileInfo.Size() {
    t.Errorf("Error: The destination file size in bytes does not match the\n" +
      "original source file size in bytes!\nSource File Size='%v', " +
      "Destination File Size='%v'\n",
      srcFileInfo.Size(), destFileInfo.Size())
  }

  _ = os.Remove(destFile)
}

func TestFileHelper_OpenFile_01(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "targetFile='%v', Error='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Error("fh.OpenFile(targetFile, fOpCfg, fPermCfg) returned a nil pointer.")
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close()\n" +
      "targetFile='%v', Error='%v'\n", targetFile, err.Error())
  }

}

func TestFileHelper_OpenFile_02(t *testing.T) {

  targetFile := ""

  fh := FileHelper{}

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'targetFile' is an empty string!\n"+
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_03(t *testing.T) {

  targetFile := "     "

  fh := FileHelper{}

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'targetFile' consists entirely of blank spaces!\n"+
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_04(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fOpCfg.isInitialized = false

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'fOpCfg' is INVALID!\n"+
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_05(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg.isInitialized = false

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'fPermCfg' is INVALID!\n"+
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_06(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n" +
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
      "Error='%v' \n", err.Error())
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error return from fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because targetFile does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFileReadOnly_01(t *testing.T) {

  fh := FileHelper{}

  source := "../logTest/topTest2.txt"
  source = fh.AdjustPathSlash(alogtopTest2Text)

  target := "../checkfiles/TestFileHelper_OpenFileReadOnly_01.txt"

  target = fh.AdjustPathSlash(target)

  expected := "Top level test file # 2."

  if fh.DoesFileExist(target) {

    err:= fh.DeleteDirFile(target)

    if err != nil {
      t.Errorf("Test Setup Error: Attempted deletion of preexisting " +
        "target file FAILED!\ntargetFile='%v'\nError='%v'\n",
        target, err.Error())
      return
    }

    if fh.DoesFileExist(target) {
      t.Errorf("Test Setup Error: Verification of target file deletion FAILED!\n" +
        "Target File still exists after attempted deletion!\ntargetFile='%v'\n",
        target)
      return
    }
  }

  err := fh.CopyFileByIo(source, target)

  if err != nil {
    t.Errorf("Test Setup Error: Copy of source file to target file FAILED!\n"+
      "sourceFile='%v'\ntargetFile='%v'\nError='%v'\n",
      source, target, err.Error())
    return
  }

  f, err := fh.OpenFileReadOnly(target)

  if err != nil {
    t.Errorf("Failed to open file: '%v'\nError='%v'",
      target, err.Error())
    return
  }

  bytes := make([]byte, 500)

  bytesRead, err := f.Read(bytes)

  if err != nil {
    t.Errorf("Error returned from f.Read(bytes).\n" +
      "targetFile='%v'\nError='%v'\n",target, err.Error())
    _ = f.Close()
    _ = fh.DeleteDirFile(target)
    return
  }

  s := string(bytes[0:bytesRead])

  if expected != s {
    t.Errorf("Expected read string='%v'. Instead read string='%v'",
      expected, s)
  }

  _ = f.Close()
  _ = fh.DeleteDirFile(target)
}

func TestFileHelper_OpenFileReadOnly_02(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.OpenFileReadOnly("")

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(\"\") "+
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}


func TestFileHelper_OpenFileReadOnly_03(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.OpenFileReadOnly("    ")

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(\"\") "+
      "because the input parameter consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_OpenFileReadOnly_04(t *testing.T) {

  fh := FileHelper{}

  targetFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  targetFile = fh.AdjustPathSlash(targetFile)

  _, err := fh.OpenFileReadOnly(targetFile)

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(targetFile) "+
      "because the input parameter 'targetFile' does not exist.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_OpenFileReadOnly_05(t *testing.T) {

  fh := FileHelper{}

  source := "../logTest/topTest2.txt"
  source = fh.AdjustPathSlash(alogtopTest2Text)

  target := "../checkfiles/TestFileHelper_OpenFileReadOnly_01.txt"

  target = fh.AdjustPathSlash(target)

  if fh.DoesFileExist(target) {

    err:= fh.DeleteDirFile(target)

    if err != nil {
      t.Errorf("Test Setup Error: Attempted deletion of preexisting " +
        "target file FAILED!\ntargetFile='%v'\nError='%v'\n",
        target, err.Error())
      return
    }

    if fh.DoesFileExist(target) {
      t.Errorf("Test Setup Error: Verification of target file deletion FAILED!\n" +
        "Target File still exists after attempted deletion!\ntargetFile='%v'\n",
        target)
      return
    }
  }

  err := fh.CopyFileByIo(source, target)

  if err != nil {
    t.Errorf("Test Setup Error: Copy of source file to target file FAILED!\n"+
      "sourceFile='%v'\ntargetFile='%v'\nError='%v'\n",
      source, target, err.Error())
    return
  }

  f, err := fh.OpenFileReadOnly(target)

  if err != nil {
    t.Errorf("Failed to open file: '%v'\nError='%v'",
      target, err.Error())
    return
  }

  testText := "Cannot write text to read-only file!"

  _, err = f.WriteString(testText)

  if err == nil {
    t.Errorf("Expected an error return from f.WriteString(testText) "+
      "because\n'f' references a read-only file. However, NO ERROR WAS RETURNED!\n")
  }

  err = f.Close()

  if err != nil {
    t.Errorf("Test Clean-up Error: Error return from f.Close().\n" +
      "File Name='%v'\nError='%v'\n",
      target, err.Error())
  }

  err = fh.DeleteDirFile(target)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error return from fh.DeleteDirFile(target).\n" +
      "target='%v'\nError='%v'", target, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../checkfiles/scratchTestFileHelper_OpenFileForWriting_01.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  testString := "How now, brown cow!"

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n" +
        "fh.DeleteDirFile(targetFile) returned an error!\n" +
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n" +
        "'targetFile' STILL EXISTS!\n" +
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  // truncateFile == false - targetFile does not yet exist!
  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile, false)\n" +
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: File Pointer returned by fh.OpenFileReadWrite(targetFile)\n"  +
      "is 'nil'!\ntargetFile='%v'", targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(testString)

  if bytesWritten != len(testString) {
    t.Errorf("ERROR: Bytes written to 'targetFile' DO NOT EQUAL the lenth\n" +
      "of 'testString'.\ntargetFile='%v'\nBytesWritten='%v' Length of Test String='%v'\n",
      targetFile, bytesWritten, len(testString))
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Sync()

  if err != nil {
    t.Errorf("Error returned by fPtr.Sync() for 'targetFile'!\n" +
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  b := make([]byte, 500)

  bytesRead, err := fPtr.ReadAt(b,0)

  if err != nil {
    if err != io.EOF {
      t.Errorf("Non-EOF error returned by fPtr.ReadAt(b,0).\n" +
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      _ = fPtr.Close()
      _ = fh.DeleteDirFile(targetFile)
      return
    }
  }

  if bytesRead != bytesWritten {
    t.Errorf("ERROR: The bytes written to 'targetFile' do NOT EQUAL the bytes\n" +
      "read from 'targetFile'.\ntargetFile='%v'\nBytes Read='%v'  Bytes Written='%v'\n",
    targetFile, bytesRead, bytesWritten)
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  resultStr := string(b[0:bytesRead])

  if testString != resultStr {
    t.Errorf("ERROR: Expected read string='%v'.\nInstead, read string='%v'.\n",
      testString, resultStr)
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }


  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_02(t *testing.T) {

  fh := FileHelper{}
  srcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  srcFile = fh.AdjustPathSlash(srcFile)
  targetFile := "../checkfiles/scratchTestFileHelper_OpenFileForWriting_02.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  testString := "How now, brown cow!"

  fInfo, err := os.Stat(srcFile)

  if err != nil {
    t.Errorf("ERROR: Test Setup Source File DOES NOT EXIST!\n" +
      "Source File='%v'\n", srcFile)
  }

  sourceByteSize := fInfo.Size()

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n" +
        "fh.DeleteDirFile(targetFile) returned an error!\n" +
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n" +
        "'targetFile' STILL EXISTS!\n" +
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  err = fh.CopyFileByIo(srcFile, targetFile)

  if err != nil {
    t.Errorf("Error returned by test setup op fh.CopyFileByIo(srcFile, targetFile).\n" +
      "srcFile='%v'\ntargetFile='%v'\nError='%v'\n",
      srcFile, targetFile, err.Error())
    return
  }

  if !fh.DoesFileExist(targetFile) {
    t.Errorf("Test Setup Failed! 'targetFile' does NOT EXIST!\n" +
      "targetFile='%v'\n", targetFile)
    return
  }

  // Open file with truncateFile=true
  fPtr, err := fh.OpenFileReadWrite(targetFile, true)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile)\n" +
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: File Pointer returned by fh.OpenFileReadWrite(targetFile)\n"  +
      "is 'nil'!\ntargetFile='%v'", targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(testString)

  if bytesWritten != len(testString) {
    t.Errorf("ERROR: Bytes written to 'targetFile' DO NOT EQUAL the lenth\n" +
      "of 'testString'.\ntargetFile='%v'\nBytesWritten='%v' Length of Test String='%v'\n",
      targetFile, bytesWritten, len(testString))
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Sync()

  if err != nil {
    t.Errorf("Error returned by fPtr.Sync() for 'targetFile'!\n" +
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  b := make([]byte, 500)

  bytesRead, err := fPtr.ReadAt(b,0)

  if err != nil {
    if err != io.EOF {
      t.Errorf("Non-EOF error returned by fPtr.ReadAt(b,0).\n" +
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      _ = fPtr.Close()
      _ = fh.DeleteDirFile(targetFile)
      return
    }
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned after Read Operation on fPtr.Close()!\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  fInfo, err = os.Stat(targetFile)

  if err!=nil {
    t.Errorf("ERROR: os.Stat(targetFile) shows targetFile DOES NOT EXIST!\n" +
    "targetFile='%v'\n", targetFile)
    return
  }

  targetFileByteSize := fInfo.Size()

  if sourceByteSize <= targetFileByteSize {
    t.Errorf("ERROR: Orginal Source File Byte Size is less than new " +
      "'targetFile' Byte Size!\nSource File Byte Size='%v'   " +
      "Target File Byte Size='%v'\ntargetFile='%v'\n",
      sourceByteSize, targetFileByteSize, targetFile)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  if bytesRead != bytesWritten {
    t.Errorf("ERROR: The bytes written to 'targetFile' do NOT EQUAL the bytes\n" +
      "read from 'targetFile'.\ntargetFile='%v'\nBytes Read='%v'  Bytes Written='%v'\n",
    targetFile, bytesRead, bytesWritten)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  resultStr := string(b[0:bytesRead])

  if testString != resultStr {
    t.Errorf("ERROR: Expected read string='%v'.\nInstead, read string='%v'.\n",
      testString, resultStr)
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_03(t *testing.T) {

  targetFile := ""

  fh := FileHelper{}

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because 'targetFile' is an empty string.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n" +
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }

  }

}

func TestFileHelper_OpenFileReadWrite_04(t *testing.T) {

  targetFile := "  "

  fh := FileHelper{}

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because the 'targetFile' parameter consists entirely of blank spaces.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n" +
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }
  }
}

func TestFileHelper_OpenFileReadWrite_05(t *testing.T) {

  targetFile := "../checkfiles/idontexist1/idontexist2/TestFileHelper_OpenFileReadWrite_05.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because the 'targetFile' parameter includes parent directories which DO NOT EXIST.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n" +
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

      err = fh.DeleteDirFile(targetFile)

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n" +
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }

  }
}

func TestFileHelper_OpenFileWriteOnly_01(t *testing.T) {
  fh := FileHelper{}
  srcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  srcFile = fh.AdjustPathSlash(srcFile)
  targetFile := "../checkfiles/TestFileHelper_OpenFileWriteOnly_01.txt"
  targetFile = fh.AdjustPathSlash(targetFile)

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n" +
        "fh.DeleteDirFile(targetFile) returned an error!\n" +
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n" +
        "'targetFile' STILL EXISTS!\n" +
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  err := fh.CopyFileByIo(srcFile, targetFile)

  if err != nil {
    t.Errorf("Error returned by test setup op fh.CopyFileByIo(srcFile, targetFile).\n" +
      "srcFile='%v'\ntargetFile='%v'\nError='%v'\n",
      srcFile, targetFile, err.Error())
    return
  }

  if !fh.DoesFileExist(targetFile) {
    t.Errorf("Test Setup Failed! 'targetFile' does NOT EXIST!\n" +
      "targetFile='%v'\n", targetFile)
    return
  }

  fPtr, err := fh.OpenFileWriteOnly(targetFile,false)

  if err != nil {
    t.Errorf("Error returned from fh.OpenFileWriteOnly" +
      "(targetFile,false).\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

  bytes := make([]byte, 3000)

  _, err = fPtr.Read(bytes)

  if err == nil {
    t.Errorf("Expected an error retun from fPtr.Read(bytes) " +
      "because\nthe file pointer 'fPtr' was opened as 'Write-Only'!\n" +
      "targetFile='%v'\n", targetFile)
  }

  if fPtr != nil {
    err = fPtr.Close()
    if err != nil {
      t.Errorf("Test Clean-up Error returned by fPtr.Close().\n" +
        "targetFile='%v'\nError='%v'\n",
        targetFile, err.Error())
    }
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Error("Test Clean-up Error returned by fh.DeleteDirFile(" +
      "targetFile)\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
  }

}

func TestFileHelper_SwapBasePath_01(t *testing.T) {

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  oldBasePath := "../filesfortest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"

  newPath, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_02(t *testing.T) {

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"

  oldBasePath := "../filesforTest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"

  newPath, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_03(t *testing.T) {

  targetPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  oldBasePath := "../filesforTest/levelfilesfortest"

  newBasePath := "../dirmgrtests"

  _, err := FileHelper{}.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.SwapBasePath(...) " +
      "NO ERROR WAS GENERATED!")
  }

}

func createALogTestBottomDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if !fh.DoesFileExist(targetDir) {
    err2 := fh.MakeDirAll(targetDir)

    if err2 != nil {
      return err2
    }
  }

  targetFile := fh.JoinPathsAdjustSeparators(targetDir, alogFile)

  if fh.DoesFileExist(targetFile) {
    err3 := fh.DeleteDirFile(targetFile)
    if err3 != nil {
      return err3
    }
  }

  f, err4 := fh.CreateFile(targetFile)

  if err4 != nil {
    return err4
  }

  nowTime := appLib.DateTimeUtility{}.GetDateTimeNanoSecText(time.Now().Local())

  _, err5 := f.WriteString("Sample Write - " + nowTime + "/n")

  if err5 != nil {
    _ = f.Close()
    return err5
  }

  _, err6 := f.WriteString("File Name: " + targetFile)

  if err6 != nil {
    _ = f.Close()
    return err6
  }

  _ = f.Close()
  return nil
}

func deleteALogTestBottomDirTargetDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if fh.DoesFileExist(targetDir) {
    err2 := fh.DeleteDirPathAll(targetDir)

    if err2 != nil {
      return err2
    }

    if fh.DoesFileExist(targetDir) {
      return errors.New("File still exists:" + targetDir)
    }
  }

  return nil
}
