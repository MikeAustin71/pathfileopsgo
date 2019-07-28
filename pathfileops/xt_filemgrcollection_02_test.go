package pathfileops

import (
  "fmt"
  "strings"
  "testing"
)

func TestFileMgrCollection_FindFiles_01(t *testing.T) {

  fmgrCol, err := fileMgrCollectionTestSetup01()

  if err != nil {
    t.Errorf("Error returned by fileMgrCollectionTestSetup01()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.txt"}

  fmgrCol2, err := fmgrCol.FindFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by fmgrCol.FindFiles(fsc).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fmgrCol2.GetNumOfFileMgrs() != 3 {
    t.Errorf("Expected fmgrCol2.GetNumOfDirs()==3 .\n"+
      "Instead, fmgrCol2.GetNumOfDirs()='%v' ",
      fmgrCol2.GetNumOfFileMgrs())
  }

  numOfFoundTextfiles := 0

  for i := 0; i < fmgrCol2.GetNumOfFileMgrs(); i++ {
    if fmgrCol2.fileMgrs[i].fileExt == ".txt" {
      numOfFoundTextfiles++
    }
  }

  if numOfFoundTextfiles != 3 {
    t.Errorf("Expected the number of found text files == 3.\n"+
      "Instead, number of found text files=='%v'", numOfFoundTextfiles)
  }

}

func TestFileMgrCollection_FindFiles_02(t *testing.T) {

  fmgrCol := FileMgrCollection{}

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{"*.txt"}

  fmgrCol.fileMgrs = nil

  fmgrCol2, err := fmgrCol.FindFiles(fsc)

  if err != nil {
    t.Errorf("Error returned by fmgrCol.FindFiles(fsc).\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if fmgrCol2.fileMgrs == nil {
    t.Error("ERROR: Expected fmgrCol2.fileMgrs != nil.\n" +
      "Instead, fmgrCol2.fileMgrs == nil!!!\n")
    return
  }

  if len(fmgrCol2.fileMgrs) != 0 {
    t.Errorf("ERROR: Expected len(fmgrCol2.fileMgrs)==0.\n" +
      "Instead, len(fmgrCol2.fileMgrs)=='%v'\n",
      len(fmgrCol2.fileMgrs))
  }
}

func TestFileMgrCollection_GetFileMgrArray_01(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt).\n"+
        "fileNameExt='%v'\nError='%v'",
        fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  cntr := 0

  for _, fmgr := range fMgrs.GetFileMgrArray() {

    err := fmgr.IsFileMgrValid("TestFileMgrCollection_GetFileMgrArray Error")

    if err != nil {
      t.Errorf("fmgr is INVALID! file='%v' Error='%v' ",
        fmgr.GetAbsolutePathFileName(), err.Error())
    }

    cntr++
  }

  if cntr != 10 {
    t.Errorf("Error: Expected File Manger Array Count='10'. "+
      "Instead, File Manager Array Count='%v'", cntr)
  }
}

func TestFileMgrCollection_GetFileMgrArray_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  fMgrs.fileMgrs = nil

  fMgrAry := fMgrs.GetFileMgrArray()

  if fMgrAry == nil {
    t.Error("Expected fMgrAry != nil.\n" +
      "Instead, fMgrAry == nil!!\n")
    return
  }

  if len(fMgrAry) > 0 {
    t.Errorf("Expected len(fMgrAry) = zero.\n" +
      "Instead, len(fMgrAry)='%v'", len(fMgrAry))
  }

}

func TestFileMgrCollection_GetFileMgrAtIndex_01(t *testing.T) {

  fm := make([]string, 5, 50)

  fm[0] = "../filesfortest/newfilesfortest/newerFileForTest_02.txt"
  fm[1] = "../filesfortest/newfilesfortest/newerFileForTest_03.txt"
  fm[2] = "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"
  fm[3] = "../filesfortest/oldfilesfortest/006890_WritingFiles.htm"
  fm[4] = "../filesfortest/oldfilesfortest/test.htm"

  fMgrCol := FileMgrCollection{}.New()
  var err error
  fh := FileHelper{}

  for i := 0; i < 5; i++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(fm[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())
    }

    fm[i], err = fh.MakeAbsolutePath(fm[i])

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())

    }

  }

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  if arrayLen != 5 {
    t.Errorf("Error: Expected Collection array length='5'. "+
      "Instead, array length='%v'. ", arrayLen)
  }

  fMgr, err := fMgrCol.GetFileMgrAtIndex(2)

  if err != nil {
    t.Errorf("Error returned by fMgrCol.GetFileMgrAtIndex(2). "+
      "Error='%v' ", err.Error())
    return
  }

  if fm[2] != fMgr.GetAbsolutePathFileName() {
    t.Errorf("Error: Expected fMgr[2]='%v'. "+
      "Instead, fMgr[2]='%v' ", fm[2], fMgr.GetAbsolutePathFileName())
  }

}

func TestFileMgrCollection_GetFileMgrAtIndex_02(t *testing.T) {

  fMgrCol := FileMgrCollection{}.New()
  var err error

  fm := make([]string, 5, 50)

  fm[0] = "../filesfortest/newfilesfortest/newerFileForTest_02.txt"
  fm[1] = "../filesfortest/newfilesfortest/newerFileForTest_03.txt"
  fm[2] = "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"
  fm[3] = "../filesfortest/oldfilesfortest/006890_WritingFiles.htm"
  fm[4] = "../filesfortest/oldfilesfortest/test.htm"

  for i := 0; i < 5; i++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(fm[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())
    }

  }

  _, err = fMgrCol.GetFileMgrAtIndex(99)

  if err == nil {
    t.Error("Expected an error return from fMgrCol.GetFileMgrAtIndex(99)\n" +
      "because the index '99' exceeds the collection's capacity.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_GetFileMgrAtIndex_03(t *testing.T) {

  fMgrCol := FileMgrCollection{}.New()
  var err error

  fm := make([]string, 5, 50)

  fm[0] = "../filesfortest/newfilesfortest/newerFileForTest_02.txt"
  fm[1] = "../filesfortest/newfilesfortest/newerFileForTest_03.txt"
  fm[2] = "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"
  fm[3] = "../filesfortest/oldfilesfortest/006890_WritingFiles.htm"
  fm[4] = "../filesfortest/oldfilesfortest/test.htm"

  for i := 0; i < 5; i++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(fm[i])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(fm[i]). "+
        "i='%v' fm[i]='%v' Error='%v' ", i, fm[i], err.Error())
    }

  }

  _, err = fMgrCol.GetFileMgrAtIndex(-99)

  if err == nil {
    t.Error("Expected an error return from fMgrCol.GetFileMgrAtIndex(-99)\n" +
      "because the index '-99' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_GetFileMgrAtIndex_04(t *testing.T) {

  fMgrCol := FileMgrCollection{}

  fMgrCol.fileMgrs = nil

  _, err := fMgrCol.GetFileMgrAtIndex(5)

  if err == nil {
    t.Error("Expected an error return from fMgrCol.GetFileMgrAtIndex(-99)\n" +
      "because the index '5' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_GetNumOfFileMgrs_01(t *testing.T) {
  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
    }
  }

  numOfFiles := fMgrs.GetNumOfFileMgrs()

  if 10 != numOfFiles {
    t.Errorf("ERROR: Expected 10-File Managers\n" +
      "from fMgrs.GetNumOfFileMgrs().\n"+
      "Instead, the collection actually contains %v-File Managers.\n",
      numOfFiles)
  }

}

func TestFileMgrCollection_GetNumOfFileMgrs_02(t *testing.T) {

  fMgrs := FileMgrCollection{}
  fMgrs.fileMgrs = nil

  numOfFiles := fMgrs.GetNumOfFileMgrs()

  if numOfFiles != 0 {
    t.Errorf("Expected numOfFiles=0 because FileMgrCollection is uninitialized.\n" +
      "However, fMgrs.GetNumOfFileMgrs() returned numOfFiles='%v'\n", numOfFiles)
  }

}

func TestFileMgrCollection_GetNumOfFiles_01(t *testing.T) {

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
    }
  }

  numOfFiles := fMgrs.GetNumOfFiles()

  if 10 != numOfFiles {
    t.Errorf("ERROR: Expected 10-File Managers.\n"+
      "Instead, the collection actually contains %v-File Managers.\n",
      numOfFiles)
  }

}

func TestFileMgrCollection_GetNumOfFiles_02(t *testing.T) {

  fMgrs := FileMgrCollection{}
  fMgrs.fileMgrs = nil

  numOfFiles := fMgrs.GetNumOfFiles()

  if numOfFiles != 0 {
    t.Errorf("Expected numOfFiles=0 because FileMgrCollection is uninitialized.\n" +
      "However, fMgrs.GetNumOfFiles() returned numOfFiles='%v'\n", numOfFiles)
  }

}

func TestFileMgrCollection_GetTotalFileBytes_01(t *testing.T) {

  fMgrs := FileMgrCollection{}
  fMgrs.fileMgrs = nil

  totalFileBytes := fMgrs.GetTotalFileBytes()

  if totalFileBytes != 0 {
    t.Errorf("ERROR: Expected totalFileBytes = zero, because " +
      "fMgrs is uninitialized.\n" +
      "However, actual totalFileBytes='%v'", totalFileBytes)
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_01(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_02(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 0)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_03(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 99)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(10)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_04(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1)

  if err == nil {
    t.Error("Error: Expected an Error to be returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1) " +
      "NO ERROR WAS RETURNED. ")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_05(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 8)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 8) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 12. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(8)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFileMgrAtIndex(8). Error='%v' ", err.Error())
  }

  if !insertedFMgr.Equal(&fMgr5) {
    t.Error("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_06(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", 1)

  fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)

  if err != nil {
    t.Errorf("Error returned from #1 fileMgrCollectionTestSetupFmgr01" +
      "(fileNameExt).\nfileNameExt='%v'\nError='%v'\n",
      fileNameExt, err.Error())
  }

  fMgrs1.AddFileMgr(fmgr)

  fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", 1)

  insertedFMgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)

  if err != nil {
    t.Errorf("Error returned from #2 fileMgrCollectionTestSetupFmgr01" +
      "(fileNameExt).\nfileNameExt='%v'\nError='%v'\n",
      fileNameExt, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1)

  if err == nil {
    t.Error("Expected an error return from fMgrs1.InsertFileMgrAtIndex(insertedFMgr, -1)\n"+
      "because the index was less than zero!\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_07(t *testing.T) {

  var fileNameExt string

  fMgrs1 := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs1.AddFileMgr(fmgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fh := FileHelper{}

  origPath := fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    t.Errorf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
  }

  insertedFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' ", origAbsPath, err.Error())
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 100)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 8) "+
      "Error='%v' ", err.Error())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    t.Errorf("After insertion, expected fMgrs1 Array Length == 11.\n"+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
  }

  fMgrLast, err := fMgrs1.PeekLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekLastFileMgr().\n" +
      "Error='%v'\n", err.Error())
  }

  if !insertedFMgr.Equal(&fMgrLast) {
    t.Error("Error: Expected insertedFMgr == fMgrLast.\n" +
      "However, the two File Managers WERE NOT EQUAL!")
  }

}

func TestFileMgrCollection_InsertFileMgrAtIndex_08(t *testing.T) {

  fMgrs1 := FileMgrCollection{}

  fileNameExt := fmt.Sprintf("testAddFile_%03d.txt", 1)

  fMgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)

  if err != nil {
    t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt).\n" +
      "fileNameExt='%v'\n" +
      "Error='%v'\n",
      fileNameExt, err.Error())
  }

  fMgrs1.fileMgrs = nil

  err = fMgrs1.InsertFileMgrAtIndex(fMgr, 0)

  if err != nil {
    t.Errorf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 0) "+
      "Error='%v' ", err.Error())
  }

  expectedFMgr, err := fMgrs1.PeekFirstFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs1.PeekFirstFileMgr()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedFMgr.Equal(&fMgr) {
    t.Errorf("Expected File Manager='%v'\n" +
      "Instead, File Manager='%v'",
      expectedFMgr.fileNameExt, fMgr.fileNameExt)
  }

}

func TestFileMgrCollection_PeekFirstFileMgr_01(t *testing.T) {

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string
  fileNamesArray := make([]string, 0, 30)

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    fileNamesArray = append(fileNamesArray, strings.ToLower(fileNameExt))

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
    }
  }

  firstFileMgr, err := fMgrs.PeekFirstFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekFirstFileMgr()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  firstFileNameExt := strings.ToLower(firstFileMgr.GetFileNameExt())

  if fileNamesArray[0] != firstFileNameExt {
    t.Errorf("ERROR: Expected First FileNameExt='%v'.\n"+
      "Instead, FileNameExt='%v'\n",
      fileNamesArray[0], firstFileNameExt)
  }
}

func TestFileMgrCollection_PeekFirstFileMgr_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  fMgrs.fileMgrs = nil

  _, err := fMgrs.PeekFirstFileMgr()

  if err == nil {
    t.Error("ERROR: Expected an error return from " +
      "fMgrs.PeekFirstFileMgr()\nbecause the fMgrs collection is empty!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileMgrCollection_PeekFMgrAtIndex_01(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrOut, err := fMgrs.PeekFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5).\n"+
      "Error='%v'", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected after Peek Array fMgrs Array Length == 10.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if fmgrOut.fileNameExt != "testAddFile_006.txt" {
    t.Errorf("Expected Peek file manger at index=5 to be fileNameExt='testAddFile_006.txt'.\n"+
      "Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
  }

}

func TestFileMgrCollection_PeekFMgrAtIndex_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  fMgrs.fileMgrs = nil

  _, err := fMgrs.PeekFileMgrAtIndex(5)

  if err == nil {
    t.Error("Expected an error would be returned by fMgrs.PopDirMgrAtIndex(5).\n"+
      "because fMgrs.fileMgrs is 'nil'.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

func TestFileMgrCollection_PeekFMgrAtIndex_03(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  _, err := fMgrs.PeekFileMgrAtIndex(-5)

  if err == nil {
    t.Error("Expected an error return from fMgrs.PopDirMgrAtIndex(-5).\n"+
      "because the index, '-5' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!")
    return
  }

}

func TestFileMgrCollection_PeekFMgrAtIndex_04(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  _, err := fMgrs.PeekFileMgrAtIndex(100)

  if err == nil {
    t.Error("Expected an error return from fMgrs.PopDirMgrAtIndex(100).\n"+
      "because the index, '100' exceeds the maximum range of the File Managers array.\n" +
      "However, NO ERROR WAS RETURNED!!!")
    return
  }
}

func TestFileMgrCollection_PeekLastFileMgr_01(t *testing.T) {

  testDir := "../createFilesTest"

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string
  fileNamesArray := make([]string, 0, 30)
  expectedFileNameExt := ""

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    if i == 9 {
      expectedFileNameExt = fileNameExt
    }

    fileNamesArray = append(fileNamesArray, strings.ToLower(fileNameExt))

    err := fMgrs.AddFileMgrByDirStrFileNameStr(testDir, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs." +
        "AddFileMgrByDirStrFileNameStr(testDir, fileNameExt).\n"+
        "testDir='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDir,
        fileNameExt,
        err.Error())
    }
  }

  fMgr, err := fMgrs.PeekLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekLastFileMgr()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFileNameExt := fMgr.GetFileNameExt()

  if expectedFileNameExt != actualFileNameExt {
    t.Errorf("ERROR: Expected Last File Name Extension='%v'.\n" +
      "Instead, Last File Name Extension='%v'",
      expectedFileNameExt, actualFileNameExt)
  }

}

func TestFileMgrCollection_PeekLastFileMgr_02(t *testing.T) {

  fMgrs := FileMgrCollection{}.New()

  fMgrs.fileMgrs = nil

  _, err := fMgrs.PeekLastFileMgr()

  if err == nil {
    t.Error("Expected an error return fMgrs.PeekLastFileMgr()\n" +
      "because fMgrs.fileMgrs = nil.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileMgrCollection_PopFirstFileMgr_01(t *testing.T) {

  testDir := "../createFilesTest"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fMgrs := FileMgrCollection{}.New()
  var fileNameExt string
  fileNamesArray := make([]string, 0, 30)

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)

    fileNamesArray = append(fileNamesArray, strings.ToLower(fileNameExt))

    err := fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned from fMgrs.AddFileMgrByDirFileNameExt(testDMgr, fileNameExt).\n"+
        "testDMgr='%v'"+
        "fileNameExt='%v'\nError='%v'",
        testDMgr.GetAbsolutePath(),
        fileNameExt,
        err.Error())
    }
  }

  firstFileMgr, err := fMgrs.PopFirstFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopFirstFileMgr()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  firstFileNameExt := strings.ToLower(firstFileMgr.GetFileNameExt())

  if fileNamesArray[0] != firstFileNameExt {
    t.Errorf("ERROR: Expected First FileNameExt='%v'.\n"+
      "Instead, FileNameExt='%v'\n",
      fileNamesArray[0], firstFileNameExt)
  }
}

func TestFileMgrCollection_PopFirstFileMgr_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  _, err := fMgrs.PopFirstFileMgr()

  if err == nil {
    t.Error("ERROR: Expected an error return from " +
      "fMgrs.PopFirstFileMgr()\nbecause the fMgrs collection is empty!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileMgrCollection_PopFMgrAtIndex_01(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrOut, err := fMgrs.PopFileMgrAtIndex(5)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 9 {
    t.Errorf("Expected after Pop Array fMgrs Array Length == 9.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if fmgrOut.fileNameExt != "testAddFile_006.txt" {
    t.Errorf("Expected popped file manger at index=5 to be fileNameExt='testAddFile_006.txt'.\n"+
      "Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
  }

}

func TestFileMgrCollection_PopFMgrAtIndex_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  _, err := fMgrs.PopFileMgrAtIndex( -1)

  if err == nil {
    t.Error("Expected an error would be returned by fMgrs.PopFileMgrAtIndex( -1)\n" +
      "because the index is less than zero.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_PopFMgrAtIndex_03(t *testing.T) {

  fMgrs := FileMgrCollection{}

  fMgrs.fileMgrs = nil

  _, err := fMgrs.PopFileMgrAtIndex( 1)

  if err == nil {
    t.Error("Expected an error would be returned by fMgrs.PopFileMgrAtIndex(1)\n" +
      "because fMgrs.fileMgrs is 'nil'.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestFileMgrCollection_PopFMgrAtIndex_04(t *testing.T) {


  dirStr := "../checkfiles"


  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt := fmt.Sprintf("testAddFile_%03d.txt", i+1)

    err := fMgrs.AddFileMgrByDirStrFileNameStr(dirStr, fileNameExt)

    if err != nil {
      t.Errorf("Error returned by fMgrs.AddFileMgrByDirStrFileNameStr(dirStr, fileNameExt)\n" +
        "dirStr='%v'\nfileNameExt='%v'\nError='%v'\n",
        dirStr, fileNameExt, err.Error())
      return
    }

  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Test Setup Error: Expected fMgrs Array Length == 10.\n" +
      "Instead fMgrs.GetNumOfDirs()=='%v'\n",
      fMgrs.GetNumOfFileMgrs())
    return
  }

  _, err := fMgrs.PopFileMgrAtIndex(35)

  if err == nil {
    t.Errorf("Expected an error would be returned by fMgrs.PopDirMgrAtIndex(35)\n"+
      "because '35' exceeds the fMgrs array length.\n" +
      "However, NO ERROR WAS RETURNED!!\n")
  }
}

func TestFileMgrCollection_PopFMgrAtIndex_05(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  expectedFMgr, err := fMgrs.PeekFirstFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekFirstFileMgr()\n." +
      "Error='%v'\n", err.Error())
    return
  }

  fMgrOut, err := fMgrs.PopFileMgrAtIndex(0)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(0).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 9 {
    t.Errorf("Expected after Pop Array fMgrs Array Length == 9.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if !fMgrOut.Equal(&expectedFMgr) {
    t.Errorf("Expected popped file manger at index=0 to be FileNameExt: %v.\n"+
      "Instead, FileNameExt='%v'",
      fMgrOut.fileNameExt, expectedFMgr.fileNameExt)
  }

}

func TestFileMgrCollection_PopFMgrAtIndex_06(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  expectedFMgr, err := fMgrs.PeekLastFileMgr()

  if err != nil {
    t.Errorf("Error returned by fMgrs.PeekFirstFileMgr()\n." +
      "Error='%v'\n", err.Error())
    return
  }

  fMgrOut, err := fMgrs.PopFileMgrAtIndex(9)

  if err != nil {
    t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(9).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if fMgrs.GetNumOfFileMgrs() != 9 {
    t.Errorf("Expected after Pop Array fMgrs Array Length == 9.\n"+
      "Instead fMgrs.GetNumOfDirs()=='%v'\n", fMgrs.GetNumOfFileMgrs())
  }

  if !fMgrOut.Equal(&expectedFMgr) {
    t.Errorf("Expected popped file manger at index=0 to be FileNameExt: %v.\n"+
      "Instead, FileNameExt='%v'",
      fMgrOut.fileNameExt, expectedFMgr.fileNameExt)
  }
}


func TestFileMgrCollection_PopLastFMgr_01(t *testing.T) {

  var fileNameExt string

  fMgrs := FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
    fmgr, err := fileMgrCollectionTestSetupFmgr01(fileNameExt)
    if err != nil {
      t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
    }
    fMgrs.AddFileMgr(fmgr)
  }

  if fMgrs.GetNumOfFileMgrs() != 10 {
    t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
  }

  fmgrLast, err := fMgrs.PopLastFileMgr()

  if err != nil {
    t.Errorf("Error returned from fMgrs.PopLastDirMgr(). Error='%v'", err.Error())
    return
  }

  if fmgrLast.fileNameExt != "testAddFile_010.txt" {
    t.Errorf("Expected PopLastDirMgr() to produce fmgrLast.fileNameExt='testAddFile_010.txt'.\n"+
      "Instead, fmgrLast.fileNameExt='%v'", fmgrLast.fileNameExt)
  }

}

func TestFileMgrCollection_PopLastFMgr_02(t *testing.T) {

  fMgrs := FileMgrCollection{}

  fMgrs.fileMgrs = nil

  _, err := fMgrs.PopLastFileMgr()

  if err == nil {
    t.Error("Expected an error would be returned from fMgrs.PopLastDirMgr().\n" +
      "because fMgrs.fileMgrs is 'nil'.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
    return
  }

}

// //////////////////////////////////////////////////////////////
// Test Setup Functions
// //////////////////////////////////////////////////////////////
func fileMgrCollectionTestSetup01() (FileMgrCollection, error) {
  ePrefix :=
    "Src File: xt_filemgrcollection_02_test.go  Function: fileMgrCollectionTestSetup01()\n"

  fh := FileHelper{}
  fMgrs := FileMgrCollection{}

  fPath, err := fh.MakeAbsolutePath(
    "../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/newfilesfortest/newerFileForTest_01.txt\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
    fmt.Errorf(ePrefix +
      "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
      "fPath='%v'\nError='%v'\n",
      fPath, err.Error())
  }

  fMgrs.AddFileMgr(fmgr)

  fPath, err = fh.MakeAbsolutePath(
    "../filesfortest/newfilesfortest/newerFileForTest_02.txt")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/newfilesfortest/newerFileForTest_02.txt\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err = FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
        "fPath='%v'\nError='%v'\n",
        fPath, err.Error())
  }


  fMgrs.AddFileMgr(fmgr)

  fPath, err = fh.MakeAbsolutePath(
    "../filesfortest/newfilesfortest/newerFileForTest_03.txt")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/newfilesfortest/newerFileForTest_03.txt\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err = FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
        "fPath='%v'\nError='%v'\n",
        fPath, err.Error())
  }

  fMgrs.AddFileMgr(fmgr)

  fPath, err = fh.MakeAbsolutePath(
    "../filesfortest/oldfilesfortest/006870_ReadingFiles.htm")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/oldfilesfortest/006870_ReadingFiles.htm\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err = FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
        "fPath='%v'\nError='%v'\n",
        fPath, err.Error())
  }

  fMgrs.AddFileMgr(fmgr)

  fPath, err = fh.MakeAbsolutePath(
    "../filesfortest/oldfilesfortest/006890_WritingFiles.htm")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/oldfilesfortest/006890_WritingFiles.htm\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err = FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
        "fPath='%v'\nError='%v'\n",
        fPath, err.Error())
  }

  fMgrs.AddFileMgr(fmgr)

  fPath, err = fh.MakeAbsolutePath(
    "../filesfortest/oldfilesfortest/test.htm")

  if err != nil {
    return fMgrs, fmt.Errorf(ePrefix +
      "Error returned by fh.MakeAbsolutePath" +
      "(\"../filesfortest/oldfilesfortest/test.htm\")\n" +
      "Error='%v'\n", err.Error())
  }

  fmgr, err = FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(fPath)\n" +
        "fPath='%v'\nError='%v'\n",
        fPath, err.Error())
  }

  fMgrs.AddFileMgr(fmgr)

  return fMgrs, nil
}

func fileMgrCollectionTestSetupFmgr01(fileNameExt string) (FileMgr, error) {

  ePrefix :=
    "Src File: xt_filemgrcollection_02_test.go  Function: fileMgrCollectionTestSetupFmgr01()\n"
  fh := FileHelper{}

  pathFileName := "../dirwalktests/dir01/dir02/" + fileNameExt
  adjustedPathFileName := fh.AdjustPathSlash(pathFileName)
  fPath, err := fh.MakeAbsolutePath(adjustedPathFileName)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error return by fh.MakeAbsolutePath(adjustedPathFileName).\n" +
      "adjustedPathFileName='%v'\nError='%v'\n", adjustedPathFileName, err.Error())
  }

  fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(fPath)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error return by FileMgr{}.NewFromPathFileNameExtStr(fPath).\nfPath='%v'\n" +
      "Error='%v'\n", fPath, err.Error())
  }

  return fmgr, nil
}
