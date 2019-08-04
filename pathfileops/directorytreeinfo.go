package pathfileops

import (
  "fmt"
)

// DirectoryTreeInfo - structure used
// to 'Find' files in a directory specified
// by 'StartPath'. The file search will be
// filtered by a 'FileSelectCriteria' object.
//
// 'FileSelectCriteria' is a FileSelectionCriteria type
// which contains FileNamePatterns strings and
//'FilesOlderThan' or 'FilesNewerThan' date time
// parameters which can be used as a selection
// criteria.
//
type DirectoryTreeInfo struct {
  StartPath          string
  Directories        DirMgrCollection
  FoundFiles         FileMgrCollection
  ErrReturns         []error
  FileSelectCriteria FileSelectionCriteria
}

// CopyToDirectoryTree - Copies an entire directory tree to an alternate location.
// The copy operation includes all files and all directories in the designated directory
// tree.
//
func (dirTree *DirectoryTreeInfo) CopyToDirectoryTree(
  baseDir, newBaseDir DirMgr) (DirectoryTreeInfo, error) {

  ePrefix := "DirectoryTreeInfo.CopyToDirectoryTree() "

  newDirTree := DirectoryTreeInfo{}

  err2 := baseDir.IsDirMgrValid("")

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+
      "Error: Input Parameter 'baseDir' is INVALID!\nError='%v'\n", err2.Error())
  }

  err2 = newBaseDir.IsDirMgrValid("")

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+
      "Error: Input Parameter 'newBaseDir' is INVALID!\nError='%v'\n", err2.Error())
  }

  err2 = newBaseDir.MakeDir()

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+
      "Error returned from newBaseDir.MakeDir().\n" +
      "newBaseDir.absolutePath='%v'\nError='%v'\n",
      newBaseDir.absolutePath, err2.Error())
  }

  if dirTree.Directories.dirMgrs == nil {

    dirTree.Directories.dirMgrs = make([]DirMgr, 0, 50 )

  }

  lAry := len(dirTree.Directories.dirMgrs)

  // Make the new Directory Tree
  for i := 0; i < lAry; i++ {

    newDMgr, err2 := dirTree.Directories.dirMgrs[i].SubstituteBaseDir(baseDir, newBaseDir)

    if err2 != nil {
      return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned from SubstituteBaseDir(baseDir, newBaseDir). i='%v' Error='%v'", i, err2.Error())
    }

    err2 = newDMgr.MakeDir()

    if err2 != nil {
      return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned fromnewDMgr.MakeDir()  Error='%v'", err2.Error())

    }

    newDirTree.Directories.AddDirMgr(newDMgr)

  }

  if dirTree.FoundFiles.fileMgrs == nil {
    dirTree.FoundFiles.fileMgrs = make([]FileMgr, 0, 50)
  }

  lAry = len(dirTree.FoundFiles.fileMgrs)

  for j := 0; j < lAry; j++ {

    fileDMgr, err2 := dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir)

    if err2 != nil {
      return DirectoryTreeInfo{},
      fmt.Errorf(ePrefix+
        "Error returned by dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir).\n" +
        "Error='%v'\n", err2.Error())
    }

    newFileMgr, err2 :=
      FileMgr{}.NewFromDirMgrFileNameExt(fileDMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt)

    if err2 != nil {
      return DirectoryTreeInfo{},
      fmt.Errorf(ePrefix+
        "Error returned by FileMgr{}.NewFromDirMgrFileNameExt(" +
        "dMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt)\n" +
        "dirTree.FoundFiles.fileMgrs[j].fileNameExt='%v'\nj='%v'\nError='%v'\n",
        dirTree.FoundFiles.fileMgrs[j].fileNameExt, j, err2.Error())
    }

    err2 = dirTree.FoundFiles.fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr)

    if err2 != nil {
      return DirectoryTreeInfo{},
      fmt.Errorf(ePrefix+
        "Error returned by fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr)\n" +
        "SrcFileName:'%v'\nDestFileName:'%v'\n" +
        "Error='%v'\n",
        dirTree.FoundFiles.fileMgrs[j].fileNameExt, newFileMgr.fileNameExt, err2.Error())

    }

    newDirTree.FoundFiles.AddFileMgr(newFileMgr)
  }

  return newDirTree, nil
}
