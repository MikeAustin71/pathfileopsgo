package pathfileops

import "testing"

func TestDirMgr_GetAbsolutePathElements(t *testing.T) {

  testDir := "D:\\Adir\\Bdir\\Cdir\\Ddir\\Edir"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir). Error='%v' ",
      err.Error())
  }

  elementsArray := dMgr.GetAbsolutePathElements()

  if len(elementsArray) != 6 {
    t.Errorf("Error: Expected length of Elements Array='6'. Instead, "+
      "Elements Array length='%v'", len(elementsArray))
  }

  if "D:" != elementsArray[0] {
    t.Errorf("Error. Expected elementsArray[0]=\"D:\". Instead, "+
      "elementsArray[0]=\"%v\"", elementsArray[0])
  }

  if "Adir" != elementsArray[1] {
    t.Errorf("Error. Expected elementsArray[1]=\"Adir\". Instead, "+
      "elementsArray[1]=\"%v\"", elementsArray[1])
  }

  if "Bdir" != elementsArray[2] {
    t.Errorf("Error. Expected elementsArray[2]=\"Bdir\". Instead, "+
      "elementsArray[2]=\"%v\"", elementsArray[2])
  }

  if "Cdir" != elementsArray[3] {
    t.Errorf("Error. Expected elementsArray[3]=\"Cdir\". Instead, "+
      "elementsArray[3]=\"%v\"", elementsArray[3])
  }

  if "Ddir" != elementsArray[4] {
    t.Errorf("Error. Expected elementsArray[4]=\"Ddir\". Instead, "+
      "elementsArray[4]=\"%v\"", elementsArray[4])
  }

  if "Edir" != elementsArray[5] {
    t.Errorf("Error. Expected elementsArray[4]=\"Edir\". Instead, "+
      "elementsArray[4]=\"%v\"", elementsArray[4])
  }

}

func TestDirMgr_GetDirectoryTree_01(t *testing.T) {

  testDir := "../filesfortest/levelfilesfortest"
  // filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_04_dir

  dirNames := []string{
              "levelfilesfortest",
              "level_01_dir",
              "level_02_dir",
              "level_03_dir",
              "level_04_dir"}

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dirs, errs := dMgr.GetDirectoryTree()

  if len(errs) > 0 {
    t.Error("Errors returned by dMgr.GetDirectoryTree():\n")
    for i:=0; i < len(errs); i++ {
      t.Errorf("Error='%v'\n", errs[i].Error())
    }

    return
  }

  if dirs.GetNumOfDirs() != 5 {
    t.Errorf("Expected returned number of directories would equal '5'.\n" +
      "Instead, the returned number of directories equal '%v'",
      dirs.GetNumOfDirs())
  }

  for i:=0; i < dirs.GetNumOfDirs(); i++ {
    dMgr, err := dirs.GetDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dirs.GetDirMgrAtIndex(%v).\n" +
        "Error='%v'\n\n", i, err.Error())
    }

    dName := dMgr.GetDirectoryName()
    foundName := false

    for k:=0; k < len(dirNames); k++ {
      if dName == dirNames[k] {
        foundName = true
        break
      }
    }

    if !foundName {
      t.Errorf("\nDirectory name was not among expected names.\nDir Name='%v'\n",
        dName)
    }

  }
}

func TestDirMgr_GetDirectoryTree_02(t *testing.T) {

  testDir := "../filesfortest/levelfilesfortest"
  // filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_04_dir

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr.isInitialized = false

  _, errs := dMgr.GetDirectoryTree()

  if len(errs) == 0 {
    t.Error("Expected errors to be returned by dMgr.GetDirectoryTree()\n" +
      "because 'dMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_GetDirectoryTree_03(t *testing.T) {

  testDir := "../filesfortest/iDoNotExist"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n" +
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  _, errs := dMgr.GetDirectoryTree()

  if len(errs) == 0 {
    t.Error("Expected errors to be returned by dMgr.GetDirectoryTree()\n" +
      "because 'dMgr' path DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_GetParentDirMgr_01(t *testing.T) {
  fh := FileHelper{}

  origBaseAbsPath, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\") " +
      "Error='%v' ", err.Error())
  }

  origParentPath , err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02\") " +
      "Error='%v' ", err.Error())
  }

  baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) " +
      "Error='%v' ", err.Error())
  }

  parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

  if err != nil {
    t.Errorf("Error returned by baseDMgr.GetParentDirMgr() " +
      "Error='%v' ", err.Error())
  }

  if origParentPath != parentDMgr.GetAbsolutePath() {
    t.Errorf("Error: Expected parentPath='%v'. Instead, parentPath='%v'. ",
      origParentPath, parentDMgr.GetAbsolutePath())
  }

  if true != hasParent {
    t.Errorf("Error: Expected hasParent='true'. Instead, hasParent='%v'.",
      hasParent)
  }

}

func TestDirMgr_GetParentDirMgr_02(t *testing.T) {

  origBaseAbsPath := "D:\\"


  baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) " +
      "Error='%v' ", err.Error())
  }

  parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

  if err != nil {
    t.Errorf("Error returned by baseDMgr.GetParentDirMgr() " +
      "Error='%v' ", err.Error())
  }


  isEqual := baseDMgr.Equal(&parentDMgr)

  if isEqual != true {
    t.Error("Error: Expected baseDMgr==parentDMgr. THEY ARE NOT EQUAL!")
  }

  if false != hasParent {
    t.Errorf("Error: Expected hasParent='false'. Instead, hasParent='%v'.",
      hasParent)
  }

}

func TestDirMgr_GetNumberOfAbsPathElements_01(t *testing.T) {

  origBaseAbsPath := "D:\\dir01\\dir02\\dir03\\dir04"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
      "Error='%v' ", err.Error())
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 5 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
      "number of elements='%v' ", 5, numOfElements)
  }

}

func TestDirMgr_GetNumberOfAbsPathElements_02(t *testing.T) {

  origBaseAbsPath := "D:\\"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
      "Error='%v' ", err.Error())
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 1 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
      "number of elements='%v' ", 1, numOfElements)
  }

}

func TestDirMgr_GetNumberOfAbsPathElements_03(t *testing.T) {

  origBaseAbsPath := "D:\\test01"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
      "Error='%v' ", err.Error())
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 2 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
      "number of elements='%v' ", 2, numOfElements)
  }

}