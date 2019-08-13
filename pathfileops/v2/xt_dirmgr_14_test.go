package pathfileops

import (
  "os"
  "strings"
  "testing"
)

func TestDirMgr_SetDirMgr_01(t *testing.T) {

  firstDir := "../../checkfiles"

  dMgr, err := DirMgr{}.New(firstDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(firstDir).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDir := ""

  _, err = dMgr.SetDirMgr(testDir)

  if err == nil {
    t.Error("Expected an error return from dMgr.SetDirMgr(testDir) because\n" +
      "'testDir' is an empty string.\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

}

func TestDirMgr_SetDirMgr_02(t *testing.T) {

  firstDir := "../../checkfiles"

  dMgr, err := DirMgr{}.New(firstDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(firstDir).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDir := "      "

  _, err = dMgr.SetDirMgr(testDir)

  if err == nil {
    t.Error("Expected an error return from dMgr.SetDirMgr(testDir) because\n" +
      "'testDir' consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_SetDirMgrFromKnownPathDirName_01(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  isEmpty, err := targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "baseDir='%v'\nsubDir='%v'\nError='%v'\n",baseDir, subDir, err.Error())
    return
  }

  if isEmpty {
    t.Error("ERROR: targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "returned 'isEmpty' = 'true'!")
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected absolute path does NOT match actual absolute path!\n" +
      "Expected absolute path='%v'\n" +
      "Actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.absolutePath))
  }

}

func TestDirMgr_SetDirMgrFromKnownPathDirName_02(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "./htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  isEmpty, err := targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "baseDir='%v'\nsubDir='%v'\nError='%v'\n",baseDir, subDir, err.Error())
    return
  }

  if isEmpty {
    t.Error("ERROR: targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "returned 'isEmpty' = 'true'!")
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected absolute path does NOT match actual absolute path!\n" +
      "Expected absolute path='%v'\n" +
      "Actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.absolutePath))
  }

}

func TestDirMgr_SetDirMgrFromKnownPathDirName_03(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "../../htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  isEmpty, err := targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "baseDir='%v'\nsubDir='%v'\nError='%v'\n",baseDir, subDir, err.Error())
    return
  }

  if isEmpty {
    t.Error("ERROR: targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "returned 'isEmpty' = 'true'!")
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected absolute path does NOT match actual absolute path!\n" +
      "Expected absolute path='%v'\n" +
      "Actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.absolutePath))
  }

}

func TestDirMgr_SetDirMgrFromKnownPathDirName_04(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := ".htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  isEmpty, err := targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "baseDir='%v'\nsubDir='%v'\nError='%v'\n",baseDir, subDir, err.Error())
    return
  }

  if isEmpty {
    t.Error("ERROR: targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "returned 'isEmpty' = 'true'!")
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected absolute path does NOT match actual absolute path!\n" +
      "Expected absolute path='%v'\n" +
      "Actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.absolutePath))
  }

}

func TestDirMgr_SetDirMgrFromKnownPathDirName_05(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "..htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  isEmpty, err := targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err != nil {
    t.Errorf("Error returned by targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "baseDir='%v'\nsubDir='%v'\nError='%v'\n",baseDir, subDir, err.Error())
    return
  }

  if isEmpty {
    t.Error("ERROR: targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "returned 'isEmpty' = 'true'!")
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected absolute path does NOT match actual absolute path!\n" +
      "Expected absolute path='%v'\n" +
      "Actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.absolutePath))
  }

}

func TestDirMgr_SetDirMgrFromKnownPathDirName_06(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  subDir = ""

  _, err = targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err == nil {

    t.Error("Expected an error return from targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "because subDir is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

    return
  }
}

func TestDirMgr_SetDirMgrFromKnownPathDirName_07(t *testing.T) {

  fh := FileHelper{}
  setupDir := fh.AdjustPathSlash("../../dirmgrtests/dir01/dir02/dir03")
  baseDir := fh.AdjustPathSlash("../../filesfortest")
  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")
  subDir := "htmlFilesForTest"

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(targetDir)\n"+
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  expectedAbsSetupDir,
  err := fh.MakeAbsolutePath(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(setupDir)\n"+
      "setupDir='%v'\nError:='%v'\n",
      setupDir, err.Error())
    return

  }

  expectedAbsSetupDir = strings.ToLower(expectedAbsSetupDir)

  targetDMgr, err := DirMgr{}.New(setupDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(setupDir)\n" +
      "setupDir='%v'\nError='%v'\n", setupDir, err.Error())
    return
  }

  if expectedAbsSetupDir != strings.ToLower(targetDMgr.absolutePath) {
    t.Errorf("ERROR: Expected setup absolute path does " +
      "NOT match setup actual absolute path!\n" +
      "Expected setup absolute path='%v'\n" +
      "Actual setup absolute path='%v'",
      expectedAbsSetupDir, strings.ToLower(targetDMgr.absolutePath))
    return
  }

  baseDir = ""

  _, err = targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)

  if err == nil {

    t.Error("Expected an error return from targetDMgr.SetDirMgrFromKnownPathDirName(baseDir, subDir)\n" +
      "because subDir is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

    return
  }
}

func TestDirMgr_SetDirMgrWithFileInfo_01(t *testing.T) {

  fh := FileHelper{}

  testDir1 := "../../createFilesTest/Level01/Level02"
  testDir2 := "../../checkfiles/checkfiles03/checkfiles03_02"
  testDir3 := "../../checkfiles/checkfiles03/checkfiles03_02/htmlFilesForTest"

  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  absTestDir1, err := fh.MakeAbsolutePath(testDir1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  absTestDir1 = strings.ToLower(absTestDir1)

  absTestDir3, err := fh.MakeAbsolutePath(testDir3)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir3)\n" +
      "testDir3='%v'\n" +
      "Error='%v'\n",
      testDir3, err.Error())
    return
  }

  absTestDir3 = strings.ToLower(absTestDir3)

  testDMgr, err := DirMgr{}.New(testDir1)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  if absTestDir1 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Test Setup Error:\n" +
      "Expected that setup absolute path='%v'\n" +
      "Instead, absolute path='%v'\n",
      absTestDir1, strings.ToLower(testDMgr.GetAbsolutePath()))
    return
  }

  err = testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo)

  if err != nil {
    t.Errorf("Error returned by testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo).\n" +
      "testDir2='%v'\ntargetFileInfo='%v'\nError='%v'\n",
      testDir2, targetFileInfo.Name(), err.Error())
    return
  }

  if absTestDir3 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Error:\n" +
      "Expected that final absolute path='%v'\n" +
      "Instead, final absolute path='%v'\n",
      absTestDir3, testDMgr.GetAbsolutePath())
    return
  }
}

func TestDirMgr_SetDirMgrWithFileInfo_02(t *testing.T) {

  fh := FileHelper{}

  testDir1 := "../../createFilesTest/Level01/Level02"
  testDir3 := "../../checkfiles/checkfiles03/checkfiles03_02/htmlFilesForTest"

  targetDir := fh.AdjustPathSlash("../../filesfortest/htmlFilesForTest")

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  absTestDir1, err := fh.MakeAbsolutePath(testDir1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  absTestDir1 = strings.ToLower(absTestDir1)

  absTestDir3, err := fh.MakeAbsolutePath(testDir3)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir3)\n" +
      "testDir3='%v'\n" +
      "Error='%v'\n",
      testDir3, err.Error())
    return
  }

  absTestDir3 = strings.ToLower(absTestDir3)

  testDMgr, err := DirMgr{}.New(testDir1)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  if absTestDir1 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Test Setup Error:\n" +
      "Expected that setup absolute path='%v'\n" +
      "Instead, absolute path='%v'\n",
      absTestDir1, strings.ToLower(testDMgr.GetAbsolutePath()))
    return
  }

  testDir2 := ""

  err = testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo)

  if err == nil {
    t.Error("Expected an error return from testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo)\n" +
      "because testDir2 is an empty string.\nHowever, NO ERROR WAS RETURNED!!\n")
    return
  }

  if absTestDir1 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Test Cleanup Error:\n" +
      "The testDMgr should remain unchanged.\n" +
      "Expected absolute path='%v'\n" +
      "Instead, absolute path='%v'\n",
      absTestDir1, strings.ToLower(testDMgr.GetAbsolutePath()))
    return
  }

  err = testDMgr.IsDirMgrValid("")

  if err != nil {
    t.Errorf("Expected that 'testDMgr' would be valid.\n" +
      "However, testDMgr.IsDirMgrValid(\"\") returned an error.\n" +
      "Error='%v'\n", err.Error())
  }
}

func TestDirMgr_SetDirMgrWithFileInfo_03(t *testing.T) {

  fh := FileHelper{}

  testDir1 := "../../createFilesTest/Level01/Level02"
  testDir2 := "../../checkfiles/checkfiles03/checkfiles03_02"
  testDir3 := "../../checkfiles/checkfiles03/checkfiles03_02/htmlFilesForTest"

  absTestDir1, err := fh.MakeAbsolutePath(testDir1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  absTestDir1 = strings.ToLower(absTestDir1)

  absTestDir3, err := fh.MakeAbsolutePath(testDir3)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(testDir3)\n" +
      "testDir3='%v'\n" +
      "Error='%v'\n",
      testDir3, err.Error())
    return
  }

  absTestDir3 = strings.ToLower(absTestDir3)

  testDMgr, err := DirMgr{}.New(testDir1)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir1)\n" +
      "testDir1='%v'\n" +
      "Error='%v'\n",
      testDir1, err.Error())
    return
  }

  if absTestDir1 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Test Setup Error:\n" +
      "Expected that setup absolute path='%v'\n" +
      "Instead, absolute path='%v'\n",
      absTestDir1, strings.ToLower(testDMgr.GetAbsolutePath()))
    return
  }

  var targetFileInfo os.FileInfo

  err = testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo)

  if err == nil {
    t.Error("Expected an error return from testDMgr.SetDirMgrWithFileInfo(testDir2, targetFileInfo)\n" +
      "because targetFileInfo is 'nil'\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

  if absTestDir1 != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("Test Cleanup Error:\n" +
      "The testDMgr should remain unchanged.\n" +
      "Expected absolute path='%v'\n" +
      "Instead, absolute path='%v'\n",
      absTestDir1, strings.ToLower(testDMgr.GetAbsolutePath()))
    return
  }

  err = testDMgr.IsDirMgrValid("")

  if err != nil {
    t.Errorf("Expected that 'testDMgr' would be valid.\n" +
      "However, testDMgr.IsDirMgrValid(\"\") returned an error.\n" +
      "Error='%v'\n", err.Error())
  }

}


func TestDirMgr_SetPermissions_01(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_SetPermissions_01"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  fh := FileHelper{}

  permissionsCfg2, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Test Setup Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permission2Txt, err := permissionsCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Setup Error returned by permissionsCfg2.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  testDMgr.isInitialized = false

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because testDMgr is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  testDMgr.isInitialized = true

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), permissionsCfg2)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), permission2Txt, err.Error())
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_02(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_SetPermissions_02"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  fh := FileHelper{}

  permissionsCfg2, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Test Setup Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permission2Txt, err := permissionsCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Setup Error returned by permissionsCfg2.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permissionsCfg.isInitialized = false

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because permissionsCfg is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), permissionsCfg2)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), permission2Txt, err.Error())
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_03(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_SetPermissions_01"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because testDMgr directory DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_04(t *testing.T) {

  testDir := "../../dirmgrtests/TestDirMgr_SetPermissions_04"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  originalPermCfg, err := testDMgr.GetDirPermissionCodes()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.GetDirPermissionCodes()\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  originalPermCfgStr, err := originalPermCfg.GetPermissionTextCode()

  if err != nil {

    t.Errorf("Test Setup Error returned by originalPermCfg.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  newPermissionsCfgStr := "dr--r--r--"
  newPermissionsCfg, err := FilePermissionConfig{}.New(newPermissionsCfgStr)

  err = testDMgr.SetPermissions(newPermissionsCfg)

  if err != nil {
    t.Errorf("Error returned by testDMgr.SetPermissions(newPermissionsCfg).\n"+
      "testDMgr='%v'\nnewPermissionsCfg='%v'\n",
      testDMgr.GetAbsolutePath(), newPermissionsCfgStr)
  }

  actualPermCfg, err := testDMgr.GetDirPermissionCodes()

  if err != nil {
    t.Errorf("Error returning actual permission configuration by "+
      "testDMgr.GetDirPermissionCodes()\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  actualPermCfgStr, err := actualPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by actualPermCfg.GetPermissionTextCode().\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), originalPermCfg)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), originalPermCfgStr, err.Error())
  }

  if actualPermCfgStr == originalPermCfgStr {
    t.Errorf("ERROR: Actual Permission Codes equals Original Permission Codes\n"+
      "Actual Permission Codes='%v'\nOriginal Permission Codes='%v'\n",
      actualPermCfgStr, originalPermCfgStr)
  }

  actualPermCfgRunes := []rune(actualPermCfgStr)
  cntOfRs := 0
  for i := 0; i < len(actualPermCfgRunes); i++ {
    if actualPermCfgRunes[i] == 'r' {
      cntOfRs++
    }
  }

  if cntOfRs != 3 {
    t.Errorf("Expected the Actual Permissions Codes to contain 3-r's or read-only codes.\n"+
      "It did NOT! Therefore the operation to change Permissions Codes FAILED!\n"+
      "Expected Permission Codes='%v'\nActual Permission Codes='%v'\n",
      newPermissionsCfgStr, actualPermCfgStr)
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_SubstituteBaseDir_01(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../../dirwalktests/dir01/dir02/dir03")

  rawBasePath := fh.AdjustPathSlash("../../dirwalktests/dir01")

  substitutePath := fh.AdjustPathSlash("../../checkfiles")

  expectedPath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).\nError='%v'\n",
      err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.\n"+
      "Instead substituted path = '%v'\n",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.\n"+
      "Instead substituted absolute path = '%v'\n",
      expectedAbsPath, dMgrResult.absolutePath)
  }

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}

func TestDirMgr_SubstituteBaseDir_02(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return

  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).\nError='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.\n"+
      "Instead substituted path = '%v'\n",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.\n"+
      "Instead substituted absolute path = '%v'\n",
      expectedAbsPath, dMgrResult.absolutePath)
  }

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return

}

func TestDirMgr_SubstituteBaseDir_03(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\n"+
      "Error='%v'", rawBasePath, err.Error())
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())
  }

  dMgrOrig.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrOrig is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrOrig.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}

func TestDirMgr_SubstituteBaseDir_04(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n", rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrBase is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrBase.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

}

func TestDirMgr_SubstituteBaseDir_05(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrSubstitute is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrSubstitute.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}

