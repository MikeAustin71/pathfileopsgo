/*
Package 'pathfileops' provides software types and methods used in
the management, organization and control of disk files and directories.

This package incorporates three primary types:

  1. FileHelper - A series of generalized file maintenance utilities

  2. DirMgr     - Directory Manager: Designed for the creation, management
                  and control of directory paths.

  3. FileMgr    - File Manager: Designed for the creation, management and
                  control of disk files.

In addition, the following are used to manage collections of
'DirMgr' and 'FileMgr' types.

  1. DirMgrCollection  - Processes and manages collections of type 'DirMgr'

  2. FileMgrCollection - Processes and manages collections of type 'FileMgr'

  3. FileOpsCollection - Manages collections of operations performed on disk
                         files and directories.


The source code repository for this package is located at:
  https://github.com/MikeAustin71/pathfileopsgo

*/
package pathfileops

import (
  "errors"
  "fmt"
  "strings"
  "time"
)

// FileSelectionCriteria - Used is selecting file names. These
// data fields specify the criterion used to determine if a
// file should be selected for some type of operation.
// Example: find files or delete files operations
type FileSelectionCriteria struct {
  // FileNamePatterns - a string array containing one or more file matching
  // patterns. Example '*.txt' '*.log' 'common*.*'
  FileNamePatterns []string

  // FilesOlderThan - Used to select files with a modification less than this date time
  FilesOlderThan time.Time

  // FilesNewerThan - // Used to select files with a modification greater than this date time
  FilesNewerThan time.Time

  // SelectByFileMode - Used to select files with equivalent os.FileMode values.
  // To select by File Mode, set the FilePermissionCfg type to the desired value
  //  Examples:
  //    fsc := FileSelectionCriteria{}
  //
  //    err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
  //    err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
  //
  // Note: os.FileMode is an uint32 type.
  SelectByFileMode FilePermissionConfig

  // SelectCriterionMode - Can be one of three values:
  //
  // FileSelectMode.None()      = No Operation - No File Select Criterion
  //                                   mode selected
  //
  // FileSelectMode.ANDSelect() = select a file only if ALL
  //                                   the selection criterion are satisfied.
  //
  // FileSelectMode.ORSelect()  = select a file if only ONE
  //                                   of the selection criterion are satisfied.
  //
  // SEE TYPE 'FileSelectCriterionMode'
  SelectCriterionMode FileSelectCriterionMode
}

// ArePatternsActive - surveys the FileNamePatterns string
// array to determine if there currently any active search
// file pattern string.
//
// A search file pattern is considered active if the string
// length of the pattern string is greater than zero.
func (fsc *FileSelectionCriteria) ArePatternsActive() bool {

  lPats := len(fsc.FileNamePatterns)

  if lPats == 0 {
    return false
  }

  isActive := false

  for i := 0; i < lPats; i++ {
    fsc.FileNamePatterns[i] =
      strings.TrimRight(strings.TrimLeft(fsc.FileNamePatterns[i], " "), " ")
    if fsc.FileNamePatterns[i] != "" {
      isActive = true
    }

  }

  return isActive
}

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
func (dirTree *DirectoryTreeInfo) CopyToDirectoryTree(baseDir, newBaseDir DirMgr) (DirectoryTreeInfo, error) {

  ePrefix := "DirectoryTreeInfo.CopyToDirectoryTree() "

  newDirTree := DirectoryTreeInfo{}

  if !baseDir.isInitialized {
    return newDirTree, errors.New(ePrefix + "Error: Input parameter 'baseDir' is NOT initialized. It is EMPTY!")
  }

  err2 := baseDir.IsDirMgrValid("")

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+"Error: Input Parameter 'baseDir' is INVALID! Error='%v'", err2.Error())
  }

  if !newBaseDir.isInitialized {
    return newDirTree, errors.New(ePrefix + "Error: Input parameter 'newBaseDir' is NOT initialized. It is EMPTY!")

  }

  err2 = newBaseDir.IsDirMgrValid("")

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+"Error: Input Parameter 'newBaseDir' is INVALID! Error='%v'", err2.Error())
  }

  err2 = newBaseDir.MakeDir()

  if err2 != nil {
    return newDirTree, fmt.Errorf(ePrefix+"Error returned from  newBaseDir.MakeDir(). newBaseDir.absolutePath='%v'  Error='%v'", newBaseDir.absolutePath, err2.Error())
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

  lAry = len(dirTree.FoundFiles.fileMgrs)

  for j := 0; j < lAry; j++ {

    fileDMgr, err2 := dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir)

    if err2 != nil {
      return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir). Error='%v'", err2.Error())
    }

    newFileMgr, err2 := FileMgr{}.NewFromDirMgrFileNameExt(fileDMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt)

    if err2 != nil {
      return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by FileMgr{}.NewFromDirMgrFileNameExt(dMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt) dirTree.FoundFiles.fileMgrs[j].fileNameExt='%v' j='%v' Error='%v'", dirTree.FoundFiles.fileMgrs[j].fileNameExt, j, err2.Error())
    }

    err2 = dirTree.FoundFiles.fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr)

    if err2 != nil {
      return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr) SrcFileName:'%v'  DestFileName:'%v' Error='%v'", dirTree.FoundFiles.fileMgrs[j].fileNameExt, newFileMgr.fileNameExt, err2.Error())

    }

    newDirTree.FoundFiles.AddFileMgr(newFileMgr)
  }

  return newDirTree, nil
}

// DirectoryDeleteFileInfo - structure used
// to delete files in a directory specified
// by 'StartPath'. Deleted files will be selected
// based on 'DeleteFileSelectCriteria' value.
//
// 'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
// type which contains  FileNamePatterns strings and the
// FilesOlderThan or FilesNewerThan date time parameters
// which can be used as file selection criteria.
type DirectoryDeleteFileInfo struct {
  StartPath                string
  Directories              DirMgrCollection
  ErrReturns               []string
  DeleteFileSelectCriteria FileSelectionCriteria
  DeletedFiles             FileMgrCollection
}
