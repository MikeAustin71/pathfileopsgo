package pathfileops

import (
  "os"
  "testing"
)

func TestDirMgr_GetAbsolutePath_01(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.isInitialized = false

  absPath := sourceDMgr.GetAbsolutePath()

  if absPath != "" {
    t.Errorf("ERROR: Expected an empty string to be returned by sourceDMgr.GetAbsolutePath()\n"+
      "because sourceDMgr is INVALID!\nInstead, the returned string='%v'\n", absPath)
  }

}

func TestDirMgr_GetAbsolutePathElements_01(t *testing.T) {

  testDir := "D:\\Adir\\Bdir\\Cdir\\Ddir\\Edir"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir). Error='%v' ",
      err.Error())
    return
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

func TestDirMgr_GetAbsolutePathElements_02(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.absolutePath = ""

  pathElements := sourceDMgr.GetAbsolutePathElements()

  if len(pathElements) != 0 {
    t.Errorf("ERROR: Expected an empty array to be returned by "+
      "sourceDMgr.GetAbsolutePathElements()\n"+
      "because sourceDMgr.absolutePath is an empty string!\n"+
      "Instead, this method returned pathElements='%v'\n",
      pathElements)
  }

}

func TestDirMgr_GetAbsolutePathWithSeparator_01(t *testing.T) {

  sourceDir := "../filesfortest/levelfilesfortest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.absolutePath = ""

  absPathWithSep := sourceDMgr.GetAbsolutePathWithSeparator()

  if len(absPathWithSep) != 0 {
    t.Errorf("ERROR: Expected an empty string to be returned by "+
      "sourceDMgr.GetAbsolutePathWithSeparator()\n"+
      "because sourceDMgr.absolutePath is an empty string!\nInstead, the method returned "+
      "absPathWithSep='%v'\n",
      absPathWithSep)
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
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dirs, errs := dMgr.GetDirectoryTree()

  if len(errs) > 0 {
    t.Error("Errors returned by dMgr.GetDirectoryTree():\n")
    for i := 0; i < len(errs); i++ {
      t.Errorf("Error='%v'\n", errs[i].Error())
    }

    return
  }

  if dirs.GetNumOfDirs() != 5 {
    t.Errorf("Expected returned number of directories would equal '5'.\n"+
      "Instead, the returned number of directories equal '%v'",
      dirs.GetNumOfDirs())
  }

  for i := 0; i < dirs.GetNumOfDirs(); i++ {
    dMgr, err := dirs.GetDirMgrAtIndex(i)

    if err != nil {
      t.Errorf("Error returned by dirs.GetDirMgrAtIndex(%v).\n"+
        "Error='%v'\n\n", i, err.Error())
      continue
    }

    dName := dMgr.GetDirectoryName()
    foundName := false

    for k := 0; k < len(dirNames); k++ {
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
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
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
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  _, errs := dMgr.GetDirectoryTree()

  if len(errs) == 0 {
    t.Error("Expected errors to be returned by dMgr.GetDirectoryTree()\n" +
      "because 'dMgr' path DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_GetDirectoryTreeStats_01(t *testing.T) {

  testDir := "../filesfortest/levelfilesfortest"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dirStats, errs := dMgr.GetDirectoryTreeStats()

  if len(errs) > 0 {
    t.Error("Errors returned by dMgr.GetDirectoryTreeStats():\n"+
      "Errors Follow:\n\n%v", dMgr.ConsolidateErrors(errs))

    return
  }

  if 4 != dirStats.NumOfSubDirs() {
    t.Errorf("Expected returned number of sub-directories would equal '4'.\n"+
      "Instead, the returned number of directories equal '%v'",
      dirStats.NumOfSubDirs())
  }

  if 25 != dirStats.NumOfFiles() {
    t.Errorf("Expected the number of files='25'.\n"+
      "Instead, the returned number of files='%v'\n",
      dirStats.NumOfFiles())
  }

  if uint64(10017) != dirStats.NumOfBytes() {
    t.Errorf("Expected the number of bytes='10017'.\n"+
      "Instead, number of bytes='%v'\n",
      dirStats.NumOfBytes())
  }
}

func TestDirMgr_GetDirectoryTreeStats_02(t *testing.T) {

  testDir := "../dirmgrtests/TestDirMgr_" +
    "GetDirectoryTreeStats_02/dirIDoNotExist"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  _, errs := dMgr.GetDirectoryTreeStats()

  if len(errs) == 0 {
    t.Error("Expected error return from dMgr.GetDirectoryTreeStats() " +
      "because 'dMgr' Path DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }

}

func TestDirMgr_GetDirectoryTreeStats_03(t *testing.T) {

  testDir := "../filesfortest"

  dMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  dMgr.isInitialized = false

  _, errs := dMgr.GetDirectoryTreeStats()

  if len(errs) == 0 {
    t.Error("Expected error return from dMgr.GetDirectoryTreeStats() " +
      "because dMgr is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")

  }
}

func TestDirMgr_GetDirPermissionTextCodes_01(t *testing.T) {

  sourceDir := "../filesfortest/htmlFilesForTest"

  fh := FileHelper{}

  expectedPerm, err := fh.GetFileMode(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.GetFileMode(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n",
      sourceDir, err.Error())
    return
  }

  expectedPermText, err := expectedPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Setup Error returned by expectedPerm.GetPermissionTextCode().\n"+
      "Error='%v'\n", err.Error())
  }

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  fPerm, err := sourceDMgr.GetDirPermissionCodes()

  if err != nil {
    t.Errorf("Error returned by sourceDMgr.GetDirPermissionCodes().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  permissionCodes, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode().\n"+
      "Error='%v'\n", err.Error())
  }

  if expectedPermText != permissionCodes {
    t.Errorf("ERROR: Expected PermissionCodes='%v'\n"+
      "Instead, permissionCodes='%v'\n", expectedPermText, permissionCodes)
  }

}

func TestDirMgr_GetDirPermissionTextCodes_02(t *testing.T) {

  sourceDir := "../filesfortest/htmlFilesForTest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.isInitialized = false

  _, err = sourceDMgr.GetDirPermissionCodes()

  if err == nil {
    t.Error("Expected an error return from sourceDMgr.GetDirPermissionCodes()\n" +
      "because 'sourceDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgr_GetDirPermissionTextCodes_03(t *testing.T) {

  sourceDir := "../filesfortest/iDoNotExist"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  _, err = sourceDMgr.GetDirPermissionCodes()

  if err == nil {
    t.Error("Expected an error return from sourceDMgr.GetDirPermissionCodes()\n" +
      "because 'sourceDMgr' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgr_GetFileInfoPlus_01(t *testing.T) {

  sourceDir := "../filesfortest/htmlFilesForTest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  fInfoPlus, err := sourceDMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Error returned by sourceDMgr.GetFileInfoPlus()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  dirName := fInfoPlus.Name()

  if "htmlFilesForTest" != dirName {
    t.Errorf("Error: Expected dirName='htmlFilesForTest'.\n"+
      "Instead, dirName='%v'\n", dirName)
  }
}

func TestDirMgr_GetFileInfoPlus_02(t *testing.T) {

  sourceDir := "../filesfortest/htmlFilesForTest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.isInitialized = false

  _, err = sourceDMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("ERROR: Expected an error return from sourceDMgr.GetFileInfoPlus()\n" +
      "because sourceDMgr is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_GetFileInfoPlus_03(t *testing.T) {

  sourceDir := "../filesfortest/iDoNotExist"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  _, err = sourceDMgr.GetFileInfoPlus()

  if err == nil {
    t.Error("ERROR: Expected an error return from sourceDMgr.GetFileInfoPlus()\n" +
      "because sourceDMgr DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_GetNumberOfAbsPathElements_01(t *testing.T) {

  origBaseAbsPath := "D:\\dir01\\dir02\\dir03\\dir04"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). "+
      "Error='%v' ", err.Error())
    return
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 5 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, "+
      "number of elements='%v' ", 5, numOfElements)
  }

}

func TestDirMgr_GetNumberOfAbsPathElements_02(t *testing.T) {

  origBaseAbsPath := "D:\\"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). "+
      "Error='%v' ", err.Error())
    return
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 1 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, "+
      "number of elements='%v' ", 1, numOfElements)
  }

}

func TestDirMgr_GetNumberOfAbsPathElements_03(t *testing.T) {

  origBaseAbsPath := "D:\\test01"

  dMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). "+
      "Error='%v' ", err.Error())
    return
  }

  numOfElements := dMgr.GetNumberOfAbsPathElements()

  if 2 != numOfElements {
    t.Errorf("Expected Number Of directory elements='%v'. Instead, "+
      "number of elements='%v' ", 2, numOfElements)
  }

}

func TestDirMgr_GetOriginalPath_01(t *testing.T) {

  fh := FileHelper{}

  testDir := fh.AdjustPathSlash("../filesfortest/levelfilesfortest")

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  originalPath := testDMgr.GetOriginalPath()

  if testDir != originalPath {
    t.Errorf("Expected Original Path='%v'.\nInstead, Original Path='%v'.\n",
      testDir, originalPath)
  }

}

func TestDirMgr_GetParentDirMgr_01(t *testing.T) {
  fh := FileHelper{}

  origBaseAbsPath, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\") "+
      "Error='%v' ", err.Error())
  }

  origParentPath, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02\") "+
      "Error='%v' ", err.Error())
  }

  baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) "+
      "Error='%v' ", err.Error())
    return
  }

  parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

  if err != nil {
    t.Errorf("Error returned by baseDMgr.GetParentDirMgr() "+
      "Error='%v' ", err.Error())
    return
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
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) "+
      "Error='%v' ", err.Error())
    return
  }

  parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

  if err != nil {
    t.Errorf("Error returned by baseDMgr.GetParentDirMgr() "+
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

func TestDirMgr_GetParentDirMgr_03(t *testing.T) {

  fh := FileHelper{}

  origBaseAbsPath, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\") "+
      "Error='%v' ", err.Error())
    return
  }

  baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) "+
      "Error='%v' ", err.Error())
    return
  }

  baseDMgr.isInitialized = false

  _, _, err = baseDMgr.GetParentDirMgr()

  if err == nil {
    t.Error("Expected an error return from baseDMgr.GetParentDirMgr()\n" +
      "because 'baseDMgr' is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestDirMgr_GetParentPath_01(t *testing.T) {

  testDirParent := "../filesfortest"

  testDir := testDirParent + "/htmlFilesForTest"

  var err error

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

  fh := FileHelper{}

  testDirParent, err = fh.MakeAbsolutePath(testDirParent)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(testDirParent).\n"+
      "testDirParent='%v'\nError='%v'\n", testDirParent, err.Error())
  }

  actualParent := testDMgr.GetParentPath()

  if testDirParent != actualParent {
    t.Errorf("ERROR: Expected parent directory='%v'.\nInstead, parent directory='%v'\n",
      testDirParent, actualParent)
  }

}

func TestDirMgr_GetPath_01(t *testing.T) {

  fh := FileHelper{}

  sourceDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  dMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  dirPath := dMgr.GetPath()

  if sourceDir != dirPath {
    t.Errorf("ERROR: Expected dMgr.GetPath()='%v'.\nInstead, dMgr.GetPath()='%v'.\n",
      sourceDir, dirPath)
  }
}

func TestDirMgr_GetPathWithSeparator_01(t *testing.T) {

  fh := FileHelper{}

  sourceDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  dMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDir = sourceDir + string(os.PathSeparator)

  dirPath := dMgr.GetPathWithSeparator()

  if sourceDir != dirPath {
    t.Errorf("ERROR: Expected dMgr.GetPath()='%v'.\nInstead, dMgr.GetPath()='%v'.\n",
      sourceDir, dirPath)
  }
}

func TestDirMgr_GetPathWithSeparator_02(t *testing.T) {

  fh := FileHelper{}

  sourceDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  dMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDir = sourceDir + string(os.PathSeparator)

  dMgr.path = ""
  dMgr.absolutePath = ""

  dirPath := dMgr.GetPathWithSeparator()

  if dirPath != "" {
    t.Errorf("ERROR: Expected dMgr.GetPath() would return an empty string.\nInstead, dMgr.GetPath()='%v'.\n",
      dirPath)
  }
}

func TestDirMgr_GetPathWithSeparator_03(t *testing.T) {

  fh := FileHelper{}

  sourceDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  dMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDir = sourceDir + string(os.PathSeparator)

  dMgr.isInitialized = false

  dirPath := dMgr.GetPathWithSeparator()

  if dirPath != "" {
    t.Errorf("ERROR: Expected dMgr.GetPath() would return an empty string \n"+
      "becased dMgr is NOT initialized!!\nInstead, dMgr.GetPath()='%v'.\n",
      dirPath)
  }
}

func TestDirMgr_GetVolumeName_01(t *testing.T) {

  fh := FileHelper{}

  sourceDir := "../filesfortest/htmlFilesForTest"

  sourceDir, err := fh.MakeAbsolutePath(sourceDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(sourceDir).\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  expectedVolName := fh.GetVolumeName(sourceDir)

  actualVolName := sourceDMgr.GetVolumeName()

  if expectedVolName != actualVolName {
    t.Errorf("Expected volume name='%v'. Instead, volume name='%v'\n",
      expectedVolName, actualVolName)
  }

}

func TestDirMgr_IsDirMgrValid_01(t *testing.T) {

  sourceDir := "../filesfortest/htmlFilesForTest"

  sourceDMgr, err := DirMgr{}.New(sourceDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(sourceDir)\n"+
      "sourceDir='%v'\nError='%v'\n", sourceDir, err.Error())
    return
  }

  sourceDMgr.absolutePath = ""

  err = sourceDMgr.IsDirMgrValid("")

  if err == nil {
    t.Error("Expected an error return from sourceDMgr.IsDirMgrValid(\"\")\n" +
      "because sourceDMgr.path is an EMPTY STRING!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_IsInitialized_01(t *testing.T) {
  testDMgr := DirMgr{}

  isInitialized := testDMgr.IsInitialized()

  if isInitialized != false {
    t.Error("Error: Expected isInitialized='false'. Instead, isInitialized='true'.\n")
  }

}

func TestDirMgr_IsInitialized_02(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  isInitialized := testDMgr.IsInitialized()

  if isInitialized != true {
    t.Error("Error: Expected isInitialized='true'. Instead, isInitialized='false'.\n")
  }

}

func TestDirMgr_IsParentPathPopulated_01(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  isParentPopulated := testDMgr.IsParentPathPopulated()

  if isParentPopulated != true {
    t.Error("Error: Expected IsParentPathPopulated()='true'. Instead, IsParentPathPopulated()='false'.\n")
  }

}

func TestDirMgr_IsParentPathPopulated_02(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDMgr.isInitialized = false

  isParentPopulated := testDMgr.IsParentPathPopulated()

  if isParentPopulated != false {
    t.Error("Error: Expected testDMgr.IsParentPathPopulated()='false' \n" +
      "because testDMgr is INVALID!\n" +
      "Instead, IsParentPathPopulated()='true'.\n")
  }

}

func TestDirMgr_IsPathPopulated_01(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDMgr.isInitialized = false

  isParentPopulated := testDMgr.IsPathPopulated()

  if isParentPopulated != false {
    t.Error("Error: Expected testDMgr.IsPathPopulated()='false' " +
      "because testDMgr is INVALID!\n" +
      "Instead, IsPathPopulated()='true'.\n")
  }
}

func TestDirMgr_IsPathPopulated_02(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  isParentPopulated := testDMgr.IsPathPopulated()

  if isParentPopulated != true {
    t.Error("Error: Expected IsPathPopulated()='true'. Instead, IsPathPopulated()='false'.\n")
  }
}

func TestDirMgr_IsVolumeNamePopulated_01(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  isVolumeNamePopulated := testDMgr.IsVolumeNamePopulated()

  if isVolumeNamePopulated != true {
    t.Error("Error: Expected IsVolumeNamePopulated()='true'. Instead, IsVolumeNamePopulated()='false'.\n")
  }

}

func TestDirMgr_IsVolumeNamePopulated_02(t *testing.T) {

  testDMgr, err := DirMgr{}.New("../filesfortest/htmlFilesForTest")

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(\"../filesfortest/htmlFilesForTest\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDMgr.isInitialized = false

  isVolumeNamePopulated := testDMgr.IsVolumeNamePopulated()

  if isVolumeNamePopulated != false {
    t.Error("Error: Expected testDMgr.IsVolumeNamePopulated()='false'\n" +
      "because testDMgr is INVALID!\n" +
      "Instead, IsVolumeNamePopulated()='true'.\n")
  }

}
