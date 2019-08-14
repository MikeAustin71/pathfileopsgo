package pathfileops

import "testing"

func TestFileMgr_DeleteThisFile_01(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash(
    "../../createFilesTest/Level01/Level02/Level03/TestFileMgr_DeleteThisFile_01.txt")

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(testFile)\n"+
      "Attempted deletion of 'testFile' FAILED!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  err = fileMgr.DeleteThisFile()

  if err == nil {
    t.Error("Expected error return from fileMgr.DeleteThisFile()\n" +
      "because the 'fileMgr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileMgr_DeleteThisFile_02(t *testing.T) {
  fh := FileHelper{}
  testFile := fh.AdjustPathSlash("../../checkfiles/TestFileMgr_DeleteThisFile_02.txt")

  err := fh.DeleteDirFile(testFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(testFile)\n"+
      "Attempted deletion of 'testFile' FAILED!\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error thrown on FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  err = fileMgr.CreateDirAndFile()

  if err != nil {
    t.Errorf("Error returned by CreateDirAndFile().\n"+
      "File='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePathFileName(), err.Error())
    _ = fh.DeleteDirFile(testFile)
    return
  }

  err = fileMgr.DeleteThisFile()

  if err != nil {
    t.Errorf("Error returned by fileMgr.DeleteThisFile().\n"+
      "Attempted deletion of 'fileMgr' FAILED!\n"+
      "File='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  doesThisFileExist, err := fh.DoesThisFileExist(fileMgr.GetAbsolutePathFileName())

  if err != nil {
    t.Errorf("Test Clean-Up Non-Path Error returned from path!\n"+
      "Final Deletion of 'fileMgr' FAILED!\n"+
      "fileMgr='%v'\nError='%v'\n",
      fileMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirFile(fileMgr.GetAbsolutePath())
    return
  }

  if doesThisFileExist {
    t.Errorf("ERROR: Final Deletion of fileMgr FAILED!\n"+
      "File='%v'\n",
      fileMgr.GetAbsolutePath())

    _ = fh.DeleteDirFile(fileMgr.GetAbsolutePath())
  }

}

func TestFileMgr_DoesFileExist_01(t *testing.T) {

  testFile := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr"+
      "(testFile)\ntestFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  if fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='false'\n" +
      "because the fileMgr is invalid.\n" +
      "However, the return value was 'true'!!\n")
  }
}

func TestFileMgr_DoesFileExist_02(t *testing.T) {

  testFile := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile)\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  if !fileMgr.DoesFileExist() {
    t.Error("Expected return value fileMgr.DoesFileExist()='true'\n" +
      "because the 'FileMgr' file does exist.\n" +
      "However, the return value was 'false'!\n")
  }
}

func TestFileMgr_DoesThisFileExist_01(t *testing.T) {

  testFile := "../../filesfortest/newfilesfortest/newerFileForTest_01.txt"

  fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(testFile).\n"+
      "testFile='%v'\nError='%v'\n",
      testFile, err.Error())
    return
  }

  fileMgr.isInitialized = false

  _, err = fileMgr.DoesThisFileExist()

  if err == nil {
    t.Error("Expected error return from fileMgr.DoesThisFileExist()\n" +
      "because the fileMgr is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileMgr_Empty_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  fileMgr1.Empty()

  fileMgr2 := FileMgr{}
  fileMgr2.Empty()

  if !fileMgr1.Equal(&fileMgr2) {
    t.Error("Error: Expected empty fileMgr1 to equal empty fileMgr2.\n" +
      "However, THEY ARE NOT EQUAL!\n")
  }

}

func TestFileMgr_Equal_01(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  fileMgr2 := fileMgr1.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != true {
    t.Error("Expected Equal to return 'true' for fileMgr1==fileMgr1.\n" +
      "Instead, fileMgr1==fileMgr1 returned 'false'.\n")
  }

}

func TestFileMgr_Equal_02(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\FileMgmnt\\TestFile003.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("ERROR: Expected fileMgr1==fileMgr2 to return 'false'.\n" +
      "Instead, fileMgr1==fileMgr2 returned 'true'\n")
  }

}

func TestFileMgr_Equal_03(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "New(\"-rwxrwxrwx\").\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}."+
      "New(FOpenType.TypeReadWrite()).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  fAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}."+
      "New(fOpenCfg, fPermCfg).\n"+
      "Error='%v'\n", err.Error())

    return
  }

  fileMgr2.fileAccessStatus = fAccessCfg.CopyOut()

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead fileMgr1==fileMgr2 returned 'true'.\n" +
      "The fileAccessStatus values are different.\n")
  }

}

func TestFileMgr_Equal_04(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error from FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fileMgr2.dMgr = DirMgr{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead, 'true' was returned for fileMgr1==fileMgr2.\n" +
      "Directory Managers are different.\n")
  }
}

func TestFileMgr_Equal_05(t *testing.T) {
  fh := FileHelper{}

  relPath1 := "..\\..\\logTest\\CmdrX\\CmdrX.log"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v', Error='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\..\\logTest\\CmdrX\\CmdrX.log"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  fileMgr2.actualFileInfo = FileInfoPlus{}

  if fileMgr2.Equal(&fileMgr1) != false {
    t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2.\n" +
      "Instead, 'true' was returned for fileMgr1==fileMgr2.\n" +
      "actualFileInfo's are different.\n")
  }
}

func TestFileMgr_EqualAbsPaths_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_1_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path to EQUAL "+
      "fileMgr2 absolute path.\n"+
      "However, Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualAbsPaths_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 := "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 := "..\\..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\level_1_1_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path to EQUAL "+
      "fileMgr2 absolute path.\n"+
      "However, Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}

func TestFileMgr_EqualAbsPaths_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n", relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())

    return
  }

  relPath2 := "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_2_0_test.txt"
  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualAbsPaths(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path to NOT "+
      "EQUAL fileMgr2 absolute path.\n"+
      "However, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
    return
  }

}

func TestFileMgr_EqualFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_02_dir\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 file name ext to EQUAL "+
      "fileMgr2 file name ext.\n"+
      "However, they ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_2_2_xray.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT "+
      "EQUAL fileMgr2 file name ext.\n"+
      "However, they ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.jag"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 file name ext to NOT "+
      "EQUAL fileMgr2 file name ext.\n"+
      "However, they ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
  }

}

func TestFileMgr_EqualPathFileNameExt_01(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path file "+
      "name ext to EQUAL\n"+
      "fileMgr2 absolute path file name ext.\n"+
      "However, the Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\n fileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_02(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\LEVEL_1_0_TEST.TXT"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: Expected fileMgr1 absolute path file name "+
      "ext to EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "However, the Paths ARE NOT EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_03(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_02_dir\\level_1_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}

func TestFileMgr_EqualPathFileNameExt_04(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_X_0_test.txt"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT EQUAL fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }

}

func TestFileMgr_EqualPathFileNameExt_05(t *testing.T) {

  fh := FileHelper{}

  relPath1 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"

  filePath1, err := fh.MakeAbsolutePath(relPath1)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1).\n"+
      "relPath1='%v'\nError='%v'\n",
      relPath1, err.Error())
    return
  }

  fileMgr1, err := FileMgr{}.New(filePath1)

  if err != nil {
    t.Errorf("Received Error on FileMgr{}.New(filePath1).\n"+
      "filePath1='%v'\nError='%v'\n",
      filePath1, err.Error())
    return
  }

  relPath2 :=
    "..\\..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.log"

  filePath2, err := fh.MakeAbsolutePath(relPath2)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2).\n"+
      "relPath2='%v'\nError='%v'\n",
      relPath2, err.Error())
    return
  }

  fileMgr2, err := FileMgr{}.New(filePath2)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.New(filePath2).\n"+
      "filePath2='%v'\nError='%v'\n",
      filePath2, err.Error())
    return
  }

  if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
    t.Errorf("Error: expected fileMgr1 absolute path file "+
      "name ext to NOT equal fileMgr2\n"+
      "absolute path file name ext.\n"+
      "Instead, Paths ARE EQUAL!\n"+
      "fileMgr1='%v'\nfileMgr2='%v'\n",
      fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
  }
}
