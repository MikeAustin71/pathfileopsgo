package pathfileops

import (
  "fmt"
  "os"
  "testing"
)

func TestFileHelper_JoinPathsAdjustSeparators_01(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }
}

func TestFileHelper_JoinPathsAdjustSeparators_02(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_03(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common")
  file1 := "/xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_04(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_05(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common//")
  file1 := "xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_06(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common//")
  file1 := "//xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

  result1 := fh.JoinPathsAdjustSeparators(path1, file1)

  if result1 != expected1 {
    t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_07(t *testing.T) {
  fh := FileHelper{}
  path1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/")
  path12, err := fh.GetAbsPathFromFilePath(path1)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(path1) path1='%v'  Error='%v'", path1, err.Error())
  }

  file1 := "//xt_dirmgr_01_test.go"
  expected1 := fh.AdjustPathSlash("../../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")
  expected12, err := fh.GetAbsPathFromFilePath(expected1)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expected1) expected1='%v'  Error='%v'", expected1, err.Error())
  }

  result1 := fh.JoinPathsAdjustSeparators(path12, file1)

  if result1 != expected12 {
    t.Errorf("Joined path and file name. Expected result '%v'. Instead result='%v'",
      expected12, result1)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_08(t *testing.T) {
  fh := FileHelper{}
  path1 := ""
  path2 := ""

  result := fh.JoinPathsAdjustSeparators(path1, path2)

  if result != path1 {
    t.Errorf("Expected result empty string. Instead result='%v'",
      result)
  }

}

func TestFileHelper_JoinPathsAdjustSeparators_09(t *testing.T) {
  fh := FileHelper{}
  path1 := "   "
  path2 := "   "

  result := fh.JoinPathsAdjustSeparators(path1, path2)

  if result != "" {
    t.Errorf("Expected result empty string. Instead result='%v'",
      result)
  }

}

func TestFileHelper_JoinPaths_01(t *testing.T) {
  fh := FileHelper{}
  path1 := "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir"
  file1 := "level_3_1_test.txt"
  expected1 := fh.AdjustPathSlash(
    "../../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result1 := fh.JoinPaths(path1, file1)

  if result1 != expected1 {
    t.Errorf("Joined path and file name. Expected result '%v', instead got: %v", expected1, result1)
  }

}

func TestFileHelper_JoinPaths_02(t *testing.T) {
  fh := FileHelper{}
  path1 := ""
  path2 := ""

  result1 := fh.JoinPaths(path1, path2)

  if result1 != "" {
    t.Errorf("Expected result = empty string, instead got result='%v'", result1)
  }

}

func TestFileHelper_JoinPaths_03(t *testing.T) {
  fh := FileHelper{}
  path1 := "   "
  path2 := "   "

  result1 := fh.JoinPaths(path1, path2)

  if result1 != "" {
    t.Errorf("Expected result = empty string, instead got result='%v'", result1)
  }

}

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
  baseDirPath := fh.AdjustPathSlash("../../checkfiles/TestFileHelper_MakeDirAll_03")
  dirPath := fh.AdjustPathSlash("../../checkfiles/TestFileHelper_MakeDirAll_03/tdir01/tdir02/tdir03")

  _, err := os.Stat(baseDirPath)

  if err == nil {
    err = fh.DeleteDirPathAll(baseDirPath)

    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test setup.\n"+
        "Attempted deletion of base directory FAILED!\nbaseDirPath='%v'\n",
        baseDirPath)
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: Attempted deletion of baseDirPath during test setup FAILED!\n"+
        "baseDirPath still exists!\nbaseDirPath='%v'\n", baseDirPath)
    }
  }

  err = fh.MakeDirAll(dirPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeDirAll(dirPath).\n"+
      "dirPath='%v'\nError='%v'", dirPath, err.Error())
    return
  }

  _, err = os.Stat(dirPath)

  if err != nil {
    t.Errorf("Error: fh.MakeDirAll(dirPath) FAILED!\n"+
      "os.Stat() confirms that dirPath DOES NOT EXIST!\n"+
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.DeleteDirPathAll(baseDirPath)

  if err != nil {
    t.Errorf("Error returned during clean-up by fh.DeleteDirPathAll(baseDirPath).\n"+
      "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
  }

  _, err = os.Stat(baseDirPath)

  if err == nil {
    t.Errorf("ERROR: baseDirPath still EXISTS! Attempted deletion FAILED!\n"+
      "baseDirPath='%v'\n", baseDirPath)
  }

}

func TestFileHelper_MakeDirAllPerm01(t *testing.T) {

  fh := FileHelper{}
  baseDirPath := fh.AdjustPathSlash("../../checkfiles/TestFileHelper_MakeDirAllPerm01")
  dirPath := fh.AdjustPathSlash(
    "../../checkfiles/TestFileHelper_MakeDirAllPerm01/tDir1/tDir2/tDir3/tDir4")

  _, err := os.Stat(baseDirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(baseDirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test startup.\n"+
        "This means that the baseDirPath and all subsidiary directories could NOT be deleted!\n"+
        "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: During test startup attempts to delete the test baseDirPath FAILED!\n"+
        "baseDirPath still EXISTS!\nbaseDirPath='%v'\n", baseDirPath)
      return
    }

  }

  permissionCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  permissionCfg.isInitialized = false

  err = fh.MakeDirAllPerm(dirPath, permissionCfg)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from fh.MakeDirAllPerm(dirPath, permissionCfg)\n"+
      "because 'permissionCfg' is Invalid. However, NO ERROR WAS RETURNED!\n"+
      "dirPath='%v'\n", dirPath)
  }

  _, err = os.Stat(baseDirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(baseDirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(baseDirPath) during test clean-up.\n"+
        "This means that the baseDirPath and all subsidiary directories could NOT be deleted!\n"+
        "baseDirPath='%v'\nError='%v'\n", baseDirPath, err.Error())
      return
    }

    _, err = os.Stat(baseDirPath)

    if err == nil {
      t.Errorf("ERROR: During test clean-up attempts to delete the test baseDirPath FAILED!\n"+
        "baseDirPath still EXISTS!\nbaseDirPath='%v'\n", baseDirPath)
    }

  }

}

func TestFileHelper_MakeDirPerm_01(t *testing.T) {

  fh := FileHelper{}
  dirPath := fh.AdjustPathSlash(
    "../../checkfiles/TestFileHelper_MakeDirPerm_01")

  _, err := os.Stat(dirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(dirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(dirPath) during test startup.\n"+
        "This means that the dirPath could NOT be deleted!\n"+
        "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
      return
    }

    _, err = os.Stat(dirPath)

    if err == nil {
      t.Errorf("ERROR: During test startup attempts to delete the test dirPath FAILED!\n"+
        "dirPath still EXISTS!\nbaseDirPath='%v'\n", dirPath)
      return
    }

  }

  permissionCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    return
  }

  permissionCfg.isInitialized = false

  err = fh.MakeDirPerm(dirPath, permissionCfg)

  if err == nil {
    t.Errorf("ERROR: Expected an error return from fh.MakeDirPerm(dirPath, permissionCfg)\n"+
      "because 'permissionCfg' is Invalid. However, NO ERROR WAS RETURNED!\n"+
      "dirPath='%v'\n", dirPath)
  }

  _, err = os.Stat(dirPath)

  if err == nil {

    err = fh.DeleteDirPathAll(dirPath)
    if err != nil {
      t.Errorf("Error returned by fh.DeleteDirPathAll(dirPath) during test clean-up.\n"+
        "This means that the dirPath could NOT be deleted!\n"+
        "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
      return
    }

    _, err = os.Stat(dirPath)

    if err == nil {
      t.Errorf("ERROR: During test clean-up attempts to delete the test dirPath FAILED!\n"+
        "dirPath still EXISTS!\ndirPath='%v'\n", dirPath)
    }

  }

}

func TestFileHelper_MakeDir01(t *testing.T) {
  fh := FileHelper{}
  dirPath := fh.AdjustPathSlash("../../checkfiles/TestFileHelper_MakeDir01")

  _, err := os.Stat(dirPath)

  if err == nil {

    err = fh.DeleteDirFile(dirPath)

    if err != nil {
      t.Errorf("Error during test setup. Directory already exists!\n"+
        "Attempted directory deletion FAILED!\ndirPath='%v'\n", dirPath)
      return
    }

  }

  _, err = os.Stat(dirPath)

  if err == nil {
    t.Errorf("ERROR: Setup tests directory still exists!\n"+
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.MakeDir(dirPath)

  if err != nil {
    t.Errorf("Error returned by fh.MakeDir(dirPath).\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
    return
  }

  _, err = os.Stat(dirPath)

  if err != nil {
    t.Errorf("ERROR: MakeDir(dirPath) failed to create test directory.\n"+
      "dirPath='%v'\n", dirPath)
    return
  }

  err = fh.DeleteDirFile(dirPath)

  if err != nil {
    t.Errorf("Error returned by cleanup fh.DeleteDirFile(dirPath).\n"+
      "dirPath='%v'\n", dirPath)
    return
  }

  _, err = os.Stat(dirPath)

  if err == nil {
    t.Errorf("Error: Clean-up FAILED! Test directory still exists!\n"+
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
  setupFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/scratchTestFileHelper_MoveFile_01.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error during test setup deleting destination file.\n"+
        "Destination File='%v'\nError:'%v'\n",
        destFile, err)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error on test setup: destination file, STILL EXISTS!\n"+
        "Destination File='%v'", destFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error during test setup deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error on test setup: source file, STILL EXISTS!\n"+
        "Source File='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying 'setupFile' to 'srcFile'.\n"+
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n"+
      "setupFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupFile, srcFile, err.Error())
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup FAILED! Source File does NOT EXIST!!\n"+
      "srcFile='%v'", srcFile)
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returend by fh.MoveFile(srcFile, destFile)\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
    return
  }

  if fh.DoesFileExist(srcFile) {
    t.Errorf("FileHelper:MoveFile() FAILED! Source File still exists!!\n"+
      "Source File='%v'\n", srcFile)
  }

  if !fh.DoesFileExist(destFile) {
    _ = fh.DeleteDirFile(srcFile)
    _ = fh.DeleteDirFile(destFile)
    t.Errorf("FileHelper:MoveFile() FAILED! Destination File DOES NOT EXIST!\n"+
      "Destination File='%v'\n", destFile)
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Attempted deletion of destination "+
      "file FAILED!\nDestination File still exists!\nDestination File='%v'",
      destFile)
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Deletion of source file "+
      "FAILED!\nSource File='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

}

func TestFileHelper_MoveFile_02(t *testing.T) {
  fh := FileHelper{}
  srcFile := ""
  destFile := fh.AdjustPathSlash("../../logTest/scratchTestFileHelper_MoveFile_02.txt")

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test start-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because srcFile is an empty string. However, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
  }
}

func TestFileHelper_MoveFile_03(t *testing.T) {
  fh := FileHelper{}
  srcFile := "   "
  destFile := fh.AdjustPathSlash("../../logTest/scratchTestFileHelper_MoveFile_03.txt")

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test start-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because srcFile consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
  }

}

func TestFileHelper_MoveFile_04(t *testing.T) {

  fh := FileHelper{}
  setupFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := "    "

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error during test setup deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error on test setup: source file, STILL EXISTS!\n"+
        "Source File='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying 'setupFile' to 'srcFile'.\n"+
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n"+
      "setupFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupFile, srcFile, err.Error())
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because destFile consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  fileDoesExist, err := fh.DoesThisFileExist(srcFile)

  if err != nil {
    t.Errorf("Non-Path Error fh.DoesThisFileExist(srcFile)\n"+
      "Error='%v'\n", err)
    _ = fh.DeleteDirFile(srcFile)
    return
  }

  if !fileDoesExist {
    t.Errorf("Error: Source file should NOT have been deleted "+
      "during the move operation\n"+
      "because there was an error (dest is an empty string).\n"+
      "Instead the source file WAS DELETED!\n"+
      "srcFile='%v'", srcFile)
  }

  _ = fh.DeleteDirFile(srcFile)

}

func TestFileHelper_MoveFile_05(t *testing.T) {

  fh := FileHelper{}
  setupFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")
  destFile := ""

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error during test setup deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error on test setup: source file, STILL EXISTS!\n"+
        "Source File='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupFile, srcFile)

  if err != nil {
    t.Errorf("Received error copying 'setupFile' to 'srcFile'.\n"+
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n"+
      "setupFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupFile, srcFile, err.Error())
    _ = fh.DeleteDirFile(srcFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup FAILED! Source File does NOT EXIST!!\n"+
      "srcFile='%v'", srcFile)
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because 'destFile' is an empty string.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Deletion of source file "+
      "FAILED!\nSource File='%v'\nError='%v'\n",
      srcFile, err.Error())
  }

}

func TestFileHelper_MoveFile_06(t *testing.T) {

  fh := FileHelper{}
  srcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/iDoNotExist.txt")
  destFile := "../../logTest/FileMgmnt/scratchTestFileHelper_MoveFile_06.txt"

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test set-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because 'srcFile' DOES NOT EXIST!.\nHowever, NO ERROR WAS RETURNED!\n")
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error during test clean-up: Deletion of destination file "+
      "FAILED!\nDestination File='%v'\nError='%v'\n",
      destFile, err.Error())
  }

}

func TestFileHelper_MoveFile_07(t *testing.T) {
  fh := FileHelper{}
  setupDestFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_0_0_test.txt")
  setupSrcFile := fh.AdjustPathSlash("../../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/setuplevel_0_3_test.txt")
  destFile := fh.AdjustPathSlash("../../createFilesTest/scratchTestFileHelper_MoveFile_07.txt")

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error during test setup deleting destination file.\n"+
        "Destination File='%v'\nError:'%v'\n",
        destFile, err)
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Error on test setup: destination file, STILL EXISTS!\n"+
        "Destination File='%v'", destFile)
      return
    }
  }

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error during test setup deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error on test setup: srcFile file, STILL EXISTS!\n"+
        "Source File='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupDestFile, destFile)

  if err != nil {
    t.Errorf("Received error copying 'setupDestFile' to 'srcFile'.\n"+
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n"+
      "setupDestFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupDestFile, srcFile, err.Error())
    _ = os.Remove(destFile)
    return
  }

  err = fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Error returned from fh.CopyFileByIo(setupSrcFile, srcFile).\n"+
      "Test Setup for source file FAILED!\n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    _ = os.Remove(destFile)
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup FAILED! Source File does NOT EXIST!!\n"+
      "Source File='%v'", srcFile)
    _ = os.Remove(destFile)
    return
  }

  srcFileInfo, err := os.Stat(srcFile)

  if err != nil {
    t.Errorf("Unexpected Error returned from os.Stat(srcFile).\n"+
      "Source File does NOT exist!\n"+
      "srcFile='%v'\nError='%v'\n",
      srcFile, err.Error())
    _ = os.Remove(destFile)
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.MoveFile(setupSrcFile, destFile)\n"+
      "setupSrcFile='%v'\ndestFile='%v'\nError='%v'\n",
      setupSrcFile, destFile, err.Error())
    _ = os.Remove(srcFile)
    _ = os.Remove(destFile)
    return
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: MoveFile() did NOT create the destFile!\n"+
      "Destination File='%v'\n", destFile)
    _ = os.Remove(srcFile)
    _ = os.Remove(destFile)
    return
  }

  destFileInfo, err := os.Stat(destFile)

  if err != nil {
    t.Errorf("Unexpected Error returned by os.Stat(destFile).\n"+
      "Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    _ = os.Remove(srcFile)
    _ = os.Remove(destFile)
    return
  }

  if srcFileInfo.Size() != destFileInfo.Size() {
    t.Errorf("Error: The destination file size in bytes does not match the\n"+
      "original source file size in bytes!\nSource File Size='%v', "+
      "Destination File Size='%v'\n",
      srcFileInfo.Size(), destFileInfo.Size())
  }

  if fh.DoesFileExist(srcFile) {
    t.Errorf("Error: MoveFile() did NOT delete the source file!\n"+
      "Source File='%v'\n", srcFile)
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error deleting destFile.\n"+
      "Error returned by fh.DeleteDirFile(destFile)\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Error deleting srcFile.\n"+
      "Error returned by fh.DeleteDirFile(srcFile)\n"+
      "srcFile='%v'\nError='%v'\n", srcFile, err.Error())
  }

}

func TestFileHelper_MoveFile_08(t *testing.T) {
  fh := FileHelper{}

  newSrcBaseDir := "../../TestFileHelper_MoveFile_08_Source"

  newSrcDir := newSrcBaseDir + "/dir1"

  newSrcBaseDir = fh.AdjustPathSlash(newSrcBaseDir)
  newSrcDir = fh.AdjustPathSlash(newSrcDir)

  err := fh.MakeDirAll(newSrcDir)

  if err != nil {
    t.Errorf("Test Setup ERROR returned by fh.MakeDirAll(newSrcDir).\n"+
      "newSrcDir='%v'\nError='%v'\n", newSrcDir, err.Error())
    return
  }

  newBaseTargDir := "../../TestFileHelper_MoveFile_08Dest"
  newTargDir := newBaseTargDir + "/dirX1"

  newBaseTargDir = fh.AdjustPathSlash(newBaseTargDir)
  newTargDir = fh.AdjustPathSlash(newTargDir)

  err = fh.MoveFile(newSrcDir, newTargDir)

  if err == nil {
    t.Errorf("Expected an error return from fh.MoveFile(newSrcDir, newTargDir) because\n" +
      "both parameter, 'newSrcDir' and 'newTargDir' are directories.\n" +
      "However, NO ERROR WAS RETURNED.")
  }

  err = fh.DeleteDirPathAll(newSrcBaseDir)

  if err != nil {
    t.Errorf("Test File Clean-Up ERROR: fh.DeleteDirPathAll(newSrcBaseDir).\n"+
      "newSrcBaseDir='%v'\nError='%v'\n", newSrcBaseDir, err.Error())
  }

  err = fh.DeleteDirPathAll(newBaseTargDir)

  if err != nil {
    t.Errorf("Test File Clean-Up ERROR: fh.DeleteDirPathAll(newBaseTargDir).\n"+
      "newBaseTargDir='%v'\nError='%v'\n", newBaseTargDir, err.Error())
  }

}

func TestFileHelper_MoveFile_09(t *testing.T) {
  fh := FileHelper{}
  setupSrcFile := fh.AdjustPathSlash("../../logTest/FileMgmnt/TestFile003.txt")
  srcFile := fh.AdjustPathSlash("../../checkfiles/TestFile003.txt")

  destDir := "../../createFilesTest/TestFileHelper_MoveFile_09"
  destFile := fh.AdjustPathSlash(destDir + "/scratchTestFileHelper_MoveFile_09.txt")

  if fh.DoesFileExist(srcFile) {

    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Error during test setup deleting source file.\n"+
        "Source File='%v'\nError:'%v'\n",
        srcFile, err)
      return
    }

    if fh.DoesFileExist(srcFile) {
      t.Errorf("Error on test setup: source file, STILL EXISTS!\n"+
        "Source File='%v'", srcFile)
      return
    }
  }

  err := fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Setup Error deleting destination directory.\n"+
      "fh.DeleteDirPathAll(destDir)"+
      "destDir='%v'\nError:'%v'\n",
      destDir, err)
    return
  }

  err = fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Test Setup Error copying 'setupSrcFile' to 'srcFile'.\n"+
      "Test Setup FAILED! 'srcFile' does NOT Exist. \n"+
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    return
  }

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("Test Setup Error: Source File does NOT exist!\n"+
      "Source File='%v'\n", srcFile)
    return
  }

  err = fh.MoveFile(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.MoveFile(srcFile, destFile)\n" +
      "because destFile directory DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(destDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error deleting destination directory.\n"+
      "fh.DeleteDirPathAll(destDir)"+
      "destDir='%v'\nError:'%v'\n",
      destDir, err)
    return
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error: Attempted deletion of source file "+
      "file FAILED!\nfh.DeleteDirFile(srcFile)\n"+
      "Source File='%v'", srcFile)
  }

}

func TestFileHelper_OpenDirectory_01(t *testing.T) {

  directoryPath := ""

  fh := FileHelper{}

  fPtr, err := fh.OpenDirectory(directoryPath, true)

  if err == nil {
    t.Error("Expected fh.OpenDirectory(directoryPath, true) to return an error because\n" +
      "'directoryPath' is an empty string. However, NO ERROR WAS RETURNED!\n")
  }

  if fPtr != nil {
    t.Error("Expected file pointer returned by fh.OpenDirectory(...) to be nil.\n" +
      "ERROR: Returned file pointer is NOT nil!")
    _ = fPtr.Close()
  }

}

func TestFileHelper_OpenDirectory_02(t *testing.T) {

  directoryPath := "         "

  fh := FileHelper{}

  fPtr, err := fh.OpenDirectory(directoryPath, true)

  if err == nil {
    t.Error("Expected fh.OpenDirectory(directoryPath, true) to return an error because\n" +
      "'directoryPath' consists entirely of blank spaces. However, NO ERROR WAS RETURNED!\n")
  }

  if fPtr != nil {
    t.Error("Expected file pointer returned by fh.OpenDirectory(...) to be nil.\n" +
      "ERROR: Returned file pointer is NOT nil!")
    _ = fPtr.Close()
  }

}

func TestFileHelper_OpenDirectory_03(t *testing.T) {

  fh := FileHelper{}

  basePath := "./createFilesTest/TestFileHelper_OpenDirectoryPerm_03"
  basePath = fh.AdjustPathSlash(basePath)
  directoryPath := "../../createFilesTest/TestFileHelper_OpenDirectoryPerm_03/x03"

  directoryPath = fh.AdjustPathSlash(directoryPath)

  if fh.DoesFileExist(basePath) {

    err := fh.DeleteDirPathAll(basePath)

    if err != nil {
      t.Errorf("Test Setup ERROR: fh.DeleteDirPathAll(basePath) failed!\n"+
        "Returned Error='%v'\n", err.Error())
      return
    }

    if fh.DoesFileExist(basePath) {
      t.Errorf("Test Setup ERROR: 'basePath' still exists. Attempted deletion FAILED!\n"+
        "basePath='%v'", basePath)
      return
    }
  }

  fPtr, err := fh.OpenDirectory(directoryPath, false)

  if err == nil {
    t.Error("Expected fh.OpenDirectory((directoryPath, false)\n" +
      "to return an error because 'directoryPath' doesn't exist and createDir=false.\n." +
      "However, NO ERROR WAS RETURNED!\n")
  }

  if fPtr != nil {
    t.Error("Expected file pointer returned by fh.OpenDirectory(...) to be nil.\n" +
      "ERROR: Returned file pointer is NOT nil!")
    _ = fPtr.Close()
  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(basePath).\n"+
      "basePath='%v'\nError='%v'", basePath, err.Error())
  }

}

func TestFileHelper_OpenDirectory_04(t *testing.T) {

  fh := FileHelper{}

  basePath := "../../createFilesTest/TestFileHelper_OpenDirectory_04"
  directoryPath := basePath + "/x03"

  basePath = fh.AdjustPathSlash(basePath)
  directoryPath = fh.AdjustPathSlash(directoryPath)

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Setup ERROR: fh.DeleteDirPathAll(basePath) failed!\n"+
      "Returned Error='%v'\n", err.Error())
    return
  }

  fPtr, err := fh.OpenDirectory(directoryPath, true)

  if err != nil {
    t.Errorf("Error returned by fh.OpenDirectory((directoryPath, true)\n"+
      "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
  }

  if fPtr == nil {
    t.Error("Expected valid file pointer to be returned by fh.OpenDirectory(...).\n" +
      "However, the file pointer is nil!")

  } else {

    err = fPtr.Close()

    if err != nil {
      t.Errorf("Test Clean-Up ERROR: Error returned by fPtr.Close().\n"+
        "Error='%v'\n", err.Error())
    }

  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(basePath).\n"+
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
  }
}

func TestFileHelper_OpenDirectory_05(t *testing.T) {

  fh := FileHelper{}

  basePath := "../../createFilesTest/TestFileHelper_OpenDirectory_05"
  directoryPath := basePath + "/x03"

  basePath = fh.AdjustPathSlash(basePath)
  directoryPath = fh.AdjustPathSlash(directoryPath)

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Setup ERROR: fh.DeleteDirPathAll(basePath) failed!\n"+
      "Returned Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(directoryPath)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(directoryPath)\n"+
      "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
    return
  }

  fPtr, err := fh.OpenDirectory(directoryPath, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenDirectory((directoryPath, true)\n"+
      "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
  }

  if fPtr == nil {

    t.Error("Expected valid file pointer to be returned by fh.OpenDirectory(...).\n" +
      "However, the file pointer is nil!")

  } else {

    err = fPtr.Close()

    if err != nil {
      t.Errorf("Test Clean-Up ERROR: Error returned by fPtr.Close().\n"+
        "Error='%v'\n", err.Error())
    }

  }

  err = fh.DeleteDirPathAll(basePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(basePath).\n"+
      "basePath='%v'\nError='%v'\n", basePath, err.Error())
  }
}

func TestFileHelper_OpenFile_01(t *testing.T) {

  targetFile := "../../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n"+
      "targetFile='%v', Error='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Error("fh.OpenFile(targetFile, fOpCfg, fPermCfg) returned a nil pointer.")
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close()\n"+
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
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'targetFile' is an empty string!\n" +
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
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'targetFile' consists entirely of blank spaces!\n" +
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_04(t *testing.T) {

  targetFile := "../../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
  }

  fOpCfg.isInitialized = false

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'fOpCfg' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_05(t *testing.T) {

  targetFile := "../../filesfortest/levelfilesfortest/level_0_0_test.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg.isInitialized = false

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error from return fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because parameter 'fPermCfg' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!\n")

    if fPtr != nil {
      _ = fPtr.Close()
    }

    return
  }

}

func TestFileHelper_OpenFile_06(t *testing.T) {

  targetFile := "../../filesfortest/levelfilesfortest/iDoNotExist.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenType(FOpenType.TypeReadOnly()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpCfg.SetFileOpenModes(FOpenMode.ModeAppend())\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fPtr, err := fh.OpenFile(targetFile, fOpCfg, fPermCfg)

  if err == nil {
    t.Error("Expected an error return from fh.OpenFile(targetFile, fOpCfg, fPermCfg)\n" +
      "because targetFile does NOT exist.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  if fPtr != nil {
    _ = fPtr.Close()
  }

  return
}
