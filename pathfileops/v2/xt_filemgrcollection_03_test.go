package pathfileops

import (
  "fmt"
  "sort"
  "strings"
  "testing"
)

func TestFileMgrCollection_SortByAbsPathFileName_01(t *testing.T) {

  testDir1 := "../../dirmgrtests/dir01/dir02"

  runelc := 'a'

  const aryLen = 12

  expectedAry := make([]string, aryLen)

  fh := FileHelper{}

  fMgrCol := FileMgrCollection{}.New()

  var err error

  for i:=0 ; i < aryLen; i++ {


    strChar := string(runelc)

    if (i+1) % 2 == 0 {

      strChar =  strings.ToUpper(strChar)

    }

    fileName := fmt.Sprintf("fileName_%v_%03d.txt", strChar, i+1)

    testFile := testDir1 + "/" + fileName

    testFile, err = fh.MakeAbsolutePath(testFile)

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(testFile)\n" +
        "testFile='%v'\nError='%v'\n", testFile, err.Error())
      return
    }

    runelc++

    expectedAry[i] = testFile

  }

  for j:=0; j < aryLen; j++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(expectedAry[(aryLen-1)-j])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(expectedAry[(aryLen-1)-j])\n" +
        "(aryLen-1)-j=%v'" +
        "expectedAry[(aryLen-1)-j]='%v'\nError='%v'\n", (aryLen-1)-j, expectedAry[(aryLen-1)-j], err.Error())
      return
    }

  }

  var fMgr FileMgr

  fMgrCol.SortByAbsPathFileName(true)

  for m:=0; m < aryLen; m++ {

    fMgr, err = fMgrCol.PeekFileMgrAtIndex(m)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(index)\n" +
        "index='%v'\nError='%v\n", m, err.Error())
      return
    }

    if expectedAry[m] != fMgr.absolutePathFileName {
      t.Errorf("For index='%v'.\n" +
        "Expected='%v'\n" +
        "  Actual='%v'\n", m, expectedAry[m], fMgr.absolutePathFileName)
    }

  }

}

func TestFileMgrCollection_SortByAbsPathFileName_02(t *testing.T) {

  testDir1 := "../../dirmgrtests/dir01/dir02"

  runelc := 'a'

  const aryLen = 12

  expectedAry := make([]string, aryLen)

  fh := FileHelper{}

  fMgrCol := FileMgrCollection{}.New()

  var err error

  for i:=0 ; i < aryLen; i++ {


    strChar := string(runelc)

    if (i+1) % 2 == 0 {

      strChar =  strings.ToUpper(strChar)

    }

    fileName := fmt.Sprintf("fileName_%v_%03d.txt", strChar, i+1)

    testFile := testDir1 + "/" + fileName

    testFile, err = fh.MakeAbsolutePath(testFile)

    if err != nil {
      t.Errorf("Error returned by fh.MakeAbsolutePath(testFile)\n" +
        "testFile='%v'\nError='%v'\n", testFile, err.Error())
      return
    }

    runelc++

    expectedAry[i] = testFile

  }

  for j:=0; j < aryLen; j++ {

    err = fMgrCol.AddFileMgrByPathFileNameExt(expectedAry[(aryLen-1)-j])

    if err != nil {
      t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(expectedAry[(aryLen-1)-j])\n" +
        "(aryLen-1)-j=%v'" +
        "expectedAry[(aryLen-1)-j]='%v'\nError='%v'\n", (aryLen-1)-j, expectedAry[(aryLen-1)-j], err.Error())
      return
    }

  }

  var fMgr FileMgr

  fMgrCol.SortByAbsPathFileName(false)

  sort.Strings(expectedAry)

  for m:=0; m < aryLen; m++ {

    fMgr, err = fMgrCol.PeekFileMgrAtIndex(m)

    if err != nil {
      t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(index)\n" +
        "index='%v'\nError='%v\n", m, err.Error())
      return
    }

    if expectedAry[m] != fMgr.absolutePathFileName {
      t.Errorf("For index='%v'.\n" +
        "Expected='%v'\n" +
        "  Actual='%v'\n", m, expectedAry[m], fMgr.absolutePathFileName)
    }

  }

}

