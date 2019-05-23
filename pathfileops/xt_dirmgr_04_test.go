
package pathfileops

import (
  "fmt"
  "os"
  "testing"
  "time"
)

func TestDirMgr_FindFilesByNamePattern_01(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_FindFilesByNamePattern_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir2).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.htm")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.htm\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 3 {
    t.Errorf("Test Setup Error: Expected to find 3-htm files in 'testDir'.\n"+
      "Instead, %v-htm files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesByNamePattern_02(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_FindFilesByNamePattern_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir2).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fMgrCollection, err := dMgr.FindFilesByNamePattern("*.txt")

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesByNamePattern(\"*.txt\").\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 10 {
    t.Errorf("Test Setup Error: Expected to find 10-txt files in 'testDir'.\n"+
      "Instead, %v-htm files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesByNamePattern_03(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_FindFilesByNamePattern_03"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir2).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  dMgr.isInitialized = false

  _, err = dMgr.FindFilesByNamePattern("*.txt")

  if err == nil {
    t.Error("Expected an error return from dMgr.FindFilesByNamePattern(\"*.txt\")\n" +
      "because 'dMgr' is INVALID.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  dMgr.isInitialized = true

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesByNamePattern_04(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_FindFilesByNamePattern_04"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir2).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_4_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_01_dir/level_1_4_test.txt")
  // 10 src Files

  // 3 sub dir src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 5 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 10 {
      oldBase = "../filesfortest/levelfilesfortest/level_01_dir"
      newBase = testDir
    } else {

      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  _, err = dMgr.FindFilesByNamePattern("")

  if err == nil {
    t.Error("Expected an error return from dMgr.FindFilesByNamePattern(\"\")\n" +
      "because the input parameter is an EMPTY STRING.\n" +
      "However, NO ERROR WAS RETURNED!")
  }


  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesByNamePattern_05(t *testing.T) {

  testDir := "../checkfiles/TestDirMgr_FindFilesByNamePattern_05"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  _, err = dMgr.FindFilesByNamePattern("*.*")

  if err == nil {
    t.Error("Expected an error return from dMgr.FindFilesByNamePattern(\"*.*\")\n" +
      "because the 'dMgr' path DOES NOT EXIST.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesBySelectCriteria_01(t *testing.T) {

  testDir := "../checkfiles/FindFilesBySelectCriteria_01"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  // 4 txt src Files
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")

  // 3 htm src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  // 3 js src files
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/index.js")
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/paths.js")
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/todoInput.js")

  // 3 md src files
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/CODE_OF_CONDUCT.md")
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/CONTRIBUTION.md")
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/DNCArticle.md")


  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 4 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 7 {
      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    } else if i < 10 {

      oldBase = "../filesfortest/jsFilesForTest"
      newBase = testDir
    } else {
      oldBase = "../filesfortest/mdFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}
  searchPattern1 := "*.htm"
  searchPattern2 := "*.md"
  fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}

  fMgrCollection, err := dMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesBySelectCriteria(fsc).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 6 {
    t.Errorf("Test Setup Error: Expected to find 6-htm and md files in 'testDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  for i := 0;  i < fMgrCollection.GetNumOfFileMgrs(); i++ {

    fmgr, err := fMgrCollection.GetFileMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n" +
        "Error='%v'\n", i, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    ext := fmgr.GetFileExt()

    if ext != ".htm" && ext !=".md" {
      t.Errorf("Error: Wrong file returned by search. Expected returned\n"+
        "to have a file extension of 'htm' or 'md'. Instead, this file had an\n" +
        "extension of '%v'.\nFileName='%v'",
        ext, fmgr.GetAbsolutePath())
    }

  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesBySelectCriteria_02(t *testing.T) {

  testDir := "../checkfiles/FindFilesBySelectCriteria_02"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  srcFiles := make([]string, 0, 50)

  // 4 txt src Files
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_0_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_1_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_2_test.txt")
  srcFiles = append(srcFiles, "../filesfortest/levelfilesfortest/level_0_3_test.txt")

  // 3 htm src files
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006860_sample.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006870_ReadingFiles.htm")
  srcFiles = append(srcFiles, "../filesfortest/htmlFilesForTest/006890_WritingFiles.htm")

  // 3 js src files
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/index.js")
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/paths.js")
  srcFiles = append(srcFiles, "../filesfortest/jsFilesForTest/todoInput.js")

  // 3 md src files
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/CODE_OF_CONDUCT.md")
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/CONTRIBUTION.md")
  srcFiles = append(srcFiles, "../filesfortest/mdFilesForTest/DNCArticle.md")

  // Total of 13-files

  destFile := ""
  oldBase := ""
  newBase := ""

  for i := 0; i < len(srcFiles); i++ {

    if i < 4 {
      oldBase = "../filesfortest/levelfilesfortest"
      newBase = testDir
    } else if i < 7 {
      oldBase = "../filesfortest/htmlFilesForTest"
      newBase = testDir
    } else if i < 10 {

      oldBase = "../filesfortest/jsFilesForTest"
      newBase = testDir
    } else {
      oldBase = "../filesfortest/mdFilesForTest"
      newBase = testDir
    }

    destFile, err = fh.SwapBasePath(oldBase, newBase, srcFiles[i])

    if err != nil {
      t.Errorf("Test File Set Up Error Stage #3 SwapBasePath(oldBase, newBase, srcFiles[%v])\n"+
        "oldBase='%v'\nnewBase='%v'\nError='%v'\n",
        i, oldBase, newBase, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

    err = fh.CopyFileByIo(srcFiles[i], destFile)

    if err != nil {
      t.Errorf("Test Setup Error returned by fh.CopyFileByIo(srcFiles[%v], destFile)\n"+
        "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
        i, srcFiles[i], destFile, err.Error())

      _ = fh.DeleteDirPathAll(testDir)

      return
    }

  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}

  fMgrCollection, err := dMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    t.Errorf("Test Setup Error returned by dMgr.FindFilesBySelectCriteria(fsc).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  if fMgrCollection.GetNumOfFileMgrs() != 13 {
    t.Errorf("Test Setup Error: Expected to find 13-files in 'testDir'.\n"+
      "Instead, %v-files were found.", fMgrCollection.GetNumOfFileMgrs())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesBySelectCriteria_03(t *testing.T) {

  testDir := "../checkfiles/FindFilesBySelectCriteria_03/iDoNotExist"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }


  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}

  _, err = dMgr.FindFilesBySelectCriteria(fsc)

  if err == nil {
    t.Error("Expected an error from dMgr.FindFilesBySelectCriteria(fsc)\n"+
      "because 'dMgr' does NOT EXIST! However, NO ERROR WAS RETURNED!")

  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

func TestDirMgr_FindFilesBySelectCriteria_04(t *testing.T) {

  testDir := "../checkfiles/FindFilesBySelectCriteria_04"

  fh := FileHelper{}

  err := fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = fh.MakeDirAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeDirAll(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  fsc := FileSelectionCriteria{}

  dMgr.isInitialized = false

  _, err = dMgr.FindFilesBySelectCriteria(fsc)

  if err == nil {
    t.Error("Expected an error from dMgr.FindFilesBySelectCriteria(fsc)\n"+
      "because 'dMgr' is INVALID! However, NO ERROR WAS RETURNED!")
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by "+
      "fh.DeleteDirPathAll(testDir)\ntestDir='%v'\n"+
      "Error='%v'\n", testDir, err.Error())
  }

  return
}

/* -------------------------------------------------------------------------------------------------
                        Local Helper Methods
   -------------------------------------------------------------------------------------------------
*/
func dirMgr01TestCreateCheckFiles03DirFiles() (string, error) {
  ePrefix := "TestFile: xt_dirmgr_01_test.go Func: dirMgr01TestCreateCheckFiles03DirFiles() "
  fh := FileHelper{}

  origDir := fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")

  if fh.DoesFileExist(origDir) {

    err := os.RemoveAll(origDir)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
    }

  }

  if fh.DoesFileExist(origDir) {
    return "", fmt.Errorf(ePrefix+"Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
  }

  // origDir does NOT exist!
  var ModePerm os.FileMode = 0777

  err := os.MkdirAll(origDir, ModePerm)

  if err != nil {
    return "", fmt.Errorf(ePrefix+"Error returned from os.MkdirAll(origDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origDir, ModePerm, err.Error())
  }

  if !fh.DoesFileExist(origDir) {
    return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origDir='%v'", origDir)
  }

  fileDir := origDir + string(os.PathSeparator)
  newFile1 := fileDir + "checkFile30001.txt"
  fp1, err := os.Create(newFile1)

  if err != nil {
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile1). newFile1='%v' Error='%v' ", newFile1, err.Error())
  }

  newFile2 := fileDir + "checkFile30002.txt"

  fp2, err := os.Create(newFile2)

  if err != nil {
    _ = fp1.Close()
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile2). newFile2='%v' Error='%v' ", newFile2, err.Error())
  }

  newFile3 := fileDir + "checkFile30003.txt"

  fp3, err := os.Create(newFile3)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile3). newFile3='%v' Error='%v' ", newFile3, err.Error())
  }

  newFile4 := fileDir + "checkFile30004.txt"

  fp4, err := os.Create(newFile4)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    _ = fp3.Close()

    return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile4). newFile4='%v' Error='%v' ", newFile4, err.Error())
  }

  t := time.Now()
  fmtT := t.Format("2006-01-02 Mon 15:04:05.000000000 -0700 MST")
  _, err = fp4.WriteString(fmtT)

  if err != nil {
    _ = fp1.Close()
    _ = fp2.Close()
    _ = fp3.Close()
    return "", fmt.Errorf(ePrefix+"%v", err.Error())
  }

  _ = fp1.Close()
  _ = fp2.Close()
  _ = fp3.Close()
  _ = fp4.Close()

  return origDir, nil
}
