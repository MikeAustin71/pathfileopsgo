package pathfileops

import (
  "testing"
  "time"
)

func TestDirectoryTreeInfo_CopyToDirectoryTree_01(t *testing.T) {

  fh := FileHelper{}
  dir := fh.AdjustPathSlash("../testsrcdir")

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir).\n"+
      "dir='%v'\nError='%v'\n", dir, err.Error())
    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. However, it does NOT exist!\n"+
      "dMgr.path='%v'\ndMgr.AbolutePath='%v'\n",
      dMgr.path, dMgr.absolutePath)
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
      "filesOlderThan, filesNewerThan).\ndir='%v'\nError='%v'\n",
      dir, err.Error())
  }

  baseDir := fh.AdjustPathSlash("../testsrcdir")

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}." +
      "NewFromPathFileNameExtStr(baseDir)\n" +
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
  }

  substituteDir := fh.AdjustPathSlash("../testdestdir/destdir")

  substituteDMgr, err := DirMgr{}.New(substituteDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(substituteDir).\n"+
      "substituteDir='%v'\nError='%v'\n", substituteDir, err.Error())
    return
  }

  newDirTree, err := dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

  if err != nil {
    t.Errorf("Error returned by dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr).\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  if len(dirTreeInfo.Directories.dirMgrs) != len(newDirTree.Directories.dirMgrs) {

    t.Errorf("Error: Expected Number of Directories = '%v'.\n"+
      "Instead, Number of NewFromPathFileNameExtStr Directories = '%v'\n",
      len(dirTreeInfo.Directories.dirMgrs), len(newDirTree.Directories.dirMgrs))
  }

  if len(dirTreeInfo.FoundFiles.fileMgrs) != len(newDirTree.FoundFiles.fileMgrs) {
    t.Errorf("Error: Expected Number of Files = '%v'.\n"+
      "Instead, actual Number of NewFromPathFileNameExtStr Files = '%v'\n",
      len(dirTreeInfo.FoundFiles.fileMgrs), len(newDirTree.FoundFiles.fileMgrs))
  }

  for i := 0; i < len(newDirTree.FoundFiles.fileMgrs); i++ {
    doesFileExist, err := newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist()

    if err != nil {
      t.Errorf("Error returned by newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist().\n"+
        "i='%v'\nfileNameExt='%v'\nError='%v'\n",
        i, newDirTree.FoundFiles.fileMgrs[i].fileNameExt, err.Error())
    }

    if !doesFileExist {
      t.Errorf("Error: Failed to create fileNameExt='%v'.\n"+
        "It does NOT exist in target directory.\n",
        newDirTree.FoundFiles.fileMgrs[i].fileNameExt)
    }

  }

  err = substituteDMgr.DeleteAll()

  if err != nil {
    t.Errorf("Error returned from substituteDMgr.DeleteAll().\n"+
      "Error='%v'\n", err.Error())
  }

}

func TestDirectoryTreeInfo_CopyToDirectoryTree_02(t *testing.T) {

  fh := FileHelper{}
  dir := fh.AdjustPathSlash("../testsrcdir")

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir).\n"+
      "dir='%v'\nError='%v'\n", dir, err.Error())
    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. However, it does NOT exist!\n"+
      "dMgr.path='%v'\ndMgr.AbolutePath='%v'\n",
      dMgr.path, dMgr.absolutePath)
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
      "filesOlderThan, filesNewerThan).\ndir='%v'\nError='%v'\n",
      dir, err.Error())
  }

  baseDir := fh.AdjustPathSlash("../testsrcdir")

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(baseDir)\n" +
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
    return
  }

  substituteDir := fh.AdjustPathSlash("../testdestdir/destdir")

  substituteDMgr, err := DirMgr{}.New(substituteDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(substituteDir).\n"+
      "substituteDir='%v'\nError='%v'\n", substituteDir, err.Error())
    return
  }

  baseDMgr.isInitialized = false

  _, err = dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

  if err == nil {
    t.Error("Expected an error return from dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)\n" +
      "because 'baseDMgr' is INVALID!\n")
  }
}

func TestDirectoryTreeInfo_CopyToDirectoryTree_03(t *testing.T) {

  fh := FileHelper{}
  dir := fh.AdjustPathSlash("../testsrcdir")

  searchPattern := ""
  filesOlderThan := time.Time{}
  filesNewerThan := time.Time{}

  dMgr, err := DirMgr{}.New(dir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir).\n"+
      "dir='%v'\nError='%v'\n", dir, err.Error())
    return
  }

  if !dMgr.doesAbsolutePathExist {
    t.Errorf("Expected target directory to exist. However, it does NOT exist!\n"+
      "dMgr.path='%v'\ndMgr.AbolutePath='%v'\n",
      dMgr.path, dMgr.absolutePath)
    return
  }

  fsc := FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{searchPattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = FileSelectMode.ANDSelect()

  dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
      "filesOlderThan, filesNewerThan).\ndir='%v'\nError='%v'\n",
      dir, err.Error())
  }

  baseDir := fh.AdjustPathSlash("../testsrcdir")

  baseDMgr, err := DirMgr{}.New(baseDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(baseDir)\n" +
      "baseDir='%v'\nError='%v'\n", baseDir, err.Error())
  }

  substituteDir := fh.AdjustPathSlash("../testdestdir/destdir")

  substituteDMgr, err := DirMgr{}.New(substituteDir)

  if err != nil {
    t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(substituteDir).\n"+
      "substituteDir='%v'\nError='%v'\n", substituteDir, err.Error())
    return
  }

  substituteDMgr.isInitialized = false

  _, err = dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

  if err == nil {
    t.Error("Expected an error return from dirTreeInfo." +
      "CopyToDirectoryTree(baseDMgr, substituteDMgr)\n" +
      "because 'baseDMgr' is INVALID!\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

