package pathfileops

import (
  "errors"
  "fmt"
  "os"
  fp "path/filepath"
  "strings"
)

/*
  This source code file contains type 'DirMgr' .

  The Source Repository for this source code file is :
    https://github.com/MikeAustin71/pathfileopsgo.git

  Dependencies:
  -------------

  Types 'DirMgr'depend on types, 'FileHelper' and 'FileMgr'
  which are contained in source	code files, 'filehelper.go'
  and 'filemanager.go' located in this directory.

*/

// DirMgr - This type and its associated methods are used to
// manage directories, directory trees and directory permissions.
//
// Dependencies:
//
// Type 'DirMgr' depend on types, 'FileHelper' and 'FileMgr'
// which are contained in source code files, 'filehelper.go'
// and 'filemanager.go' located in this directory.
//
type DirMgr struct {
  isInitialized                   bool
  originalPath                    string
  path                            string // Stored with no trailing path separator
  isPathPopulated                 bool
  doesPathExist                   bool
  parentPath                      string // Stored with no trailing path separator
  isParentPathPopulated           bool
  relativePath                    string // Stored with no preceding path separator
  isRelativePathPopulated         bool
  absolutePath                    string
  isAbsolutePathPopulated         bool
  doesAbsolutePathExist           bool
  isAbsolutePathDifferentFromPath bool
  directoryName                   string // Name of directory with out parent path.
  volumeName                      string
  isVolumePopulated               bool
  actualDirFileInfo               FileInfoPlus
}

// CopyIn - Receives a pointer to a DirMgr object as an
// input parameter and copies the values from the incoming
// object to the current DirMgr object. When the copy operation
// is completed, the current DirMgr object is a duplicate of the
// incoming DirMgr object.
func (dMgr *DirMgr) CopyIn(dmgrIn *DirMgr) {

  dMgr.isInitialized = dmgrIn.isInitialized
  dMgr.originalPath = dmgrIn.originalPath
  dMgr.path = dmgrIn.path
  dMgr.isPathPopulated = dmgrIn.isPathPopulated
  dMgr.doesPathExist = dmgrIn.doesPathExist
  dMgr.parentPath = dmgrIn.parentPath
  dMgr.isParentPathPopulated = dmgrIn.isParentPathPopulated
  dMgr.relativePath = dmgrIn.relativePath
  dMgr.isRelativePathPopulated = dmgrIn.isRelativePathPopulated
  dMgr.absolutePath = dmgrIn.absolutePath
  dMgr.isAbsolutePathPopulated = dmgrIn.isAbsolutePathPopulated
  dMgr.doesAbsolutePathExist = dmgrIn.doesAbsolutePathExist
  dMgr.isAbsolutePathDifferentFromPath = dmgrIn.isAbsolutePathDifferentFromPath
  dMgr.directoryName = dmgrIn.directoryName
  dMgr.volumeName = dmgrIn.volumeName
  dMgr.isVolumePopulated = dmgrIn.isVolumePopulated
  dMgr.actualDirFileInfo = dmgrIn.actualDirFileInfo.CopyOut()
}

// CopyOut - Makes a duplicate copy of the current DirMgr values and
// returns them in a new DirMgr object.
func (dMgr *DirMgr) CopyOut() DirMgr {

  dOut := DirMgr{}

  dOut.isInitialized = dMgr.isInitialized
  dOut.originalPath = dMgr.originalPath
  dOut.path = dMgr.path
  dOut.isPathPopulated = dMgr.isPathPopulated
  dOut.doesPathExist = dMgr.doesPathExist
  dOut.parentPath = dMgr.parentPath
  dOut.isParentPathPopulated = dMgr.isParentPathPopulated
  dOut.relativePath = dMgr.relativePath
  dOut.isRelativePathPopulated = dMgr.isRelativePathPopulated
  dOut.absolutePath = dMgr.absolutePath
  dOut.isAbsolutePathPopulated = dMgr.isAbsolutePathPopulated
  dOut.doesAbsolutePathExist = dMgr.doesAbsolutePathExist
  dOut.isAbsolutePathDifferentFromPath = dMgr.isAbsolutePathDifferentFromPath
  dOut.directoryName = dMgr.directoryName
  dOut.volumeName = dMgr.volumeName
  dOut.isVolumePopulated = dMgr.isVolumePopulated
  dOut.actualDirFileInfo = dMgr.actualDirFileInfo.CopyOut()

  return dOut
}

// DeleteAll - BE CAREFUL!!! - This method will remove the directory identified by
// this DirMgr object. It will also delete all child directories and files in the
// directory tree.
//
// Example:
//
//    Run DeleteAll on Directory: "../pathfilego/003_filehelper/testdestdir/destdir"
//
//    All files and all subdirectories will be deleted.
//
//    Only the parent path will remain: "../pathfilego/003_filehelper/testdestdir"
//
func (dMgr *DirMgr) DeleteAll() error {

  ePrefix := "DirMgr.DeleteAll() "

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return err
  }

  err = nil

  _, err2 := os.Stat(dMgr.absolutePath)

  if err2 == nil {

    err2 = os.RemoveAll(dMgr.absolutePath)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(dMgr.absolutePath) "+
        "returned error. dMgr.absolutePath='%v' Error='%v' ", dMgr.absolutePath, err.Error())
    }

  }

  _ = dMgr.DoesDirMgrPathExist()
  _ = dMgr.DoesDirMgrAbsolutePathExist()

  return err
}

// DeleteAllFilesInDir - Deletes all the files in the current
// directory. ONLY files are deleted NOT directories.
//
// Files in subdirectories are NOT deleted.
//
// Reference:
// https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
//
func (dMgr *DirMgr) DeleteAllFilesInDir() (errs []error) {

  ePrefix := "DirMgr.DeleteAllFilesInDir() "

  errs = make([]error, 0, 300)
  var err, err2 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {

      if os.IsNotExist(err) {
        err2 =
          fmt.Errorf(ePrefix +
            "ERROR: Directory DOES NOT EXIST!\n" +
            "DirMgr Directory='%v'\n",dMgr.absolutePath)
      } else {
        err2 = fmt.Errorf(ePrefix + "Error returned by os.Stat(dMgr.absolutePath).\n" +
          "dMgr.absolutePath='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())
      }

    errs = append(errs, err2)
    return errs
  }

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath).\n"+
      "dMgr.absolutePath='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  fh := FileHelper{}

  nameFileInfos, err := dir.Readdir(-1)

  if err != nil {
    _ = dir.Close()
    err2 = fmt.Errorf(ePrefix+
      "Error returned by dir.Readdirnames(-1).\n"+
      "dMgr.absolutePath='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())
    errs = append(errs, err2)
    return errs

  }

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue

    } else {
      name := fh.JoinPathsAdjustSeparators(dMgr.absolutePath, nameFInfo.Name())

      err = os.Remove(name)

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by os.Remove(name).\n"+
          "dMgr.absolutePath='%v'\nfile name='%v'\nError='%v'\n",
          dMgr.absolutePath,name, err.Error())

        errs = append(errs, err2)
      }
    }
  }

  err = dir.Close()

  if err != nil {
    err2 = fmt.Errorf(ePrefix +
      "Error returned by dir.Close(). An attempt to close the os.File pointer to the current\n" +
      "DirMgr path has FAILED!\n" +
      "dMgr.absolutePath='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())
    errs = append(errs, err2)
    return errs
  }

  return errs
}

// DeleteFilesByNamePattern - Receives a string defining a pattern to use
// in searching file names for all files in the directory identified
// by the current DirMgr instance.
//
// *                        BE CAREFUL!!                         *
// If a file name matches the pattern specified by input parameter,
// 'fileSearchPattern', it will be deleted.
//
// Only files in the directory identified by the current DirMgr instance
// will be subject to deletion. Files in sub-directories or parent directories
// will NOT be deleted or altered in any way.
//
// If the 'fileSearchPattern' is improperly formatted, an error will be returned.
//
// If the directory path identified by the current DirMgr instance does NOT
// exist, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Example 'filePatterns'
//
//   *.*              will match  all files in directory
//   *.html           will match  anyfilename.html
//   a*               will match  appleJack.txt
//   j????row.txt     will match  j1x34row.txt
//   data[0-9]*       will match  data123.csv
//
//   Reference For Matching Details:
//     https://golang.org/pkg/path/filepath/#Match
//
func (dMgr *DirMgr) DeleteFilesByNamePattern(fileSearchPattern string) (errs []error) {

  ePrefix := "DirMgr.DeleteFilesByNamePattern() "

  errs = make([]error, 0, 300)

  var err2, err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs,err)
    return errs
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {

    if os.IsNotExist(err) {
      err2 =
        fmt.Errorf(ePrefix +
          "ERROR: Directory DOES NOT EXIST!\n" +
          "DirMgr Directory='%v'\n",dMgr.absolutePath)
    } else {
      err2 = fmt.Errorf(ePrefix + "Error returned by os.Stat(dMgr.absolutePath).\n" +
        "dMgr.absolutePath='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())
    }

    errs = append(errs,err2)
    return errs
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode == -1 {
    err2 = errors.New(ePrefix +
      "Error: Input parameter 'fileSearchPattern' is an empty string!")

    errs = append(errs,err2)
    return errs
  }

  if errCode == -2 {
    err2 = errors.New(ePrefix +
      "Error: Input parameter 'fileSearchPattern' consists of blank spaces!")
    errs = append(errs,err2)
    return errs
  }

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())
    errs = append(errs,err2)
    return errs
  }

  nameFileInfos, err := dir.Readdir(-1)

  if err != nil {
    _ = dir.Close()
    err2 = fmt.Errorf(ePrefix+
      "Error returned by dir.Readdirnames(-1). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())
    errs = append(errs,err2)
    return errs
  }

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue

    } else {

      fName := nameFInfo.Name()

      isMatch, err := fp.Match(fileSearchPattern, fName)

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by fp.Match(fileSearchPattern, fileName). "+
          "directorySearched='%v' fileSearchPattern='%v' fileName='%v' Error='%v' ",
          dMgr.absolutePath, fileSearchPattern, fName, err.Error())
        errs = append(errs,err2)
        continue
      }

      if !isMatch {
        continue
      } else {

        fullName := fh.JoinPathsAdjustSeparators(dMgr.absolutePath, fName)

        err = os.Remove(fullName)

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "Error returned by os.Remove(fullName).\n"+
            "dMgr.absolutePath='%v'\nfullName='%v'\nError='%v'\n\n",
            dMgr.absolutePath,fullName, err.Error())
          errs = append(errs,err2)
          continue
        }
      }
    }
  }

  err = dir.Close()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned by dir.Close(). "+
      "dir='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())
    errs = append(errs,err2)
  }

  return errs
}

// DeleteWalkDirFiles - !!! BE CAREFUL !!! This method
// deletes files in a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the current DirMgr object. The method 'walks the
// directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// If a file matches the File Selection Criteria, it is DELETED. By the way,
// if ALL the file selection criterion are set to zero values or 'Inactive',
// then ALL FILES IN THE DIRECTORY ARE DELETED!!!
//
// A record of file deletions is included in the returned DirectoryDeleteFileInfo
// structure (DirectoryDeleteFileInfo.DeletedFiles).
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  deleteFileSelectionCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be deleted.
//
//
//    type FileSelectionCriteria struct {
//     FileNamePatterns    []string// An array of strings containing File Name Patterns
//     FilesOlderThan      time.Time// Match files with older modification date times
//     FilesNewerThan      time.Time// Match files with newer modification date times
//     SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//     SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//    FileNamePatterns    []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file
//                                    modification date times in order to determine
//                                    whether the file is older than the 'FilesOlderThan'
//                                    file selection criterion. If the file modification
//                                    date time is older than the 'FilesOlderThan' date time,
//                                    that file is considered a 'match'	for this file selection
//                                    criterion.
//
//                                    If the value of 'FilesOlderThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    FilesNewerThan      time.Time - This date time type is compared to the file
//                                    modification date time in order to determine
//                                    whether the file is newer than the 'FilesNewerThan'
//                                    file selection criterion. If the file modification date time
//                                    is newer than the 'FilesNewerThan' date time, that file is
//                                    considered a 'match' for this file selection criterion.
//
//                                    If the value of 'FilesNewerThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    SelectByFileMode  FilePermissionConfig -
//                                    Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                    selection criterion allows for the selection of files by File Mode.
//                                    File modes are compared to the value	of 'SelectByFileMode'. If the
//                                    File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                    that file is considered to be a 'match' for this file selection
//                                    criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                         fsc := FileSelectionCriteria{}
//                                           err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                           err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//    SelectCriterionMode FileSelectCriterionMode -
//                                    This parameter selects the manner in which the file selection
//                                    criteria above are applied in determining a 'match' for file
//                                    selection purposes. 'SelectCriterionMode' may be set to one of
//                                    two constant values:
//
//                                    FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                    FileSelectMode.ORSelect() - File selected if any active selection
//                                      criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed will be selected and DELETED.
//
//      Example:
//              fsc := FileSelectCriterionMode{}
//
//              In this example, 'fsc' is NOT initialized. Therefore,
//              all of the selection criterion are 'Inactive'. Consequently,
//              all of the files encountered during the search operation
//              will be SELECTED FOR DELETION!
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  DirectoryDeleteFileInfo -
//
//          type DirectoryDeleteFileInfo struct {
//            StartPath                string
//            dirMgrs                  []DirMgr
//            FoundFiles               []FileWalkInfo
//            ErrReturns               []string
//            DeleteFileSelectCriteria FileSelectionCriteria
//            DeletedFiles             []FileWalkInfo
//          }
//
//          If successful, files matching the file selection criteria
//          specified in input parameter 'deleteFileSelectionCriteria'
//          will be DELETED and returned in a 'DirectoryDeleteFileInfo'
//          structure field, 'DirectoryDeleteFileInfo.DeletedFiles.'
//
//          Note: It is a good idea to check the returned field
//                DirectoryDeleteFileInfo.ErrReturns to determine if any
//                system errors were encountered during file processing.
//
//  error - If a program execution error is encountered during processing, it will
//          returned as an 'error' type. Also, see the comment on
//          DirectoryDeleteFileInfo.ErrReturns, above.
//
func (dMgr *DirMgr) DeleteWalkDirFiles(
  deleteFileSelectionCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {

  ePrefix := "DirMgr.DeleteWalkDirFiles() "
  deleteFilesInfo := DirectoryDeleteFileInfo{}

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return deleteFilesInfo, err
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {
    if os.IsNotExist(err) {
      return deleteFilesInfo,
        fmt.Errorf(ePrefix + "ERROR: Directory DOES NOT EXIST!\n" +
          "DirMgr Directory='%v'\n",dMgr.absolutePath)
    }

    return deleteFilesInfo,
      fmt.Errorf(ePrefix + "Error returned by os.Stat(dMgr.absolutePath).\n" +
        "dMgr.absolutePath='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())
  }

  deleteFilesInfo.StartPath = dMgr.absolutePath

  deleteFilesInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

  fh := FileHelper{}

  err = fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

  if err != nil {
    return deleteFilesInfo,
      fmt.Errorf(ePrefix+"Error returned by FileHelper."+
        "makeFileHelperWalkDirDeleteFilesFunc(&dWalkInfo). "+
        "dWalkInfo.StartPath='%v' Error='%v' ", deleteFilesInfo.StartPath, err.Error())
  }

  return deleteFilesInfo, nil

}

// DoesDirectoryExist - Returns two boolean values indicating whether or not the
// Directory path exists and whether or not the Directory absolute path exists.
//
func (dMgr *DirMgr) DoesDirectoryExist() (doesPathExist, doesAbsolutePathExist bool) {

  doesPathExist = dMgr.DoesDirMgrPathExist()

  doesAbsolutePathExist = dMgr.DoesDirMgrAbsolutePathExist()

  return
}

// DoesDirMgrAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.absolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
//
func (dMgr *DirMgr) DoesDirMgrAbsolutePathExist() bool {

  errCode, _ , _ :=FileHelper{}.isStringEmptyOrBlank(dMgr.absolutePath)

  if errCode < 0 {
    dMgr.doesAbsolutePathExist = false
    return dMgr.doesAbsolutePathExist
  }

  info, err := os.Stat(dMgr.absolutePath)

  if err != nil {
    dMgr.doesAbsolutePathExist = false
  } else {
    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.absolutePath, info)
  }

  return dMgr.doesAbsolutePathExist
}

// DoesDirMgrPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition it also
// updates the DirMgr field DirMgr.doesPathExist field.
//
func (dMgr *DirMgr) DoesDirMgrPathExist() bool {

  errCode, _ , _ :=FileHelper{}.isStringEmptyOrBlank(dMgr.path)

  if errCode < 0 {
    dMgr.doesPathExist = false
    return dMgr.doesPathExist
  }

  info, err := os.Stat(dMgr.path)

  if err != nil {
    dMgr.doesPathExist = false
  } else {
    dMgr.doesPathExist = true
    dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.path, info)
  }

  return dMgr.doesPathExist

}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

  dMgr.isInitialized = false
  dMgr.originalPath = ""
  dMgr.path = ""
  dMgr.isPathPopulated = false
  dMgr.doesPathExist = false
  dMgr.parentPath = ""
  dMgr.isParentPathPopulated = false
  dMgr.relativePath = ""
  dMgr.isRelativePathPopulated = false
  dMgr.absolutePath = ""
  dMgr.isAbsolutePathPopulated = false
  dMgr.doesAbsolutePathExist = false
  dMgr.isAbsolutePathDifferentFromPath = false
  dMgr.directoryName = ""
  dMgr.volumeName = ""
  dMgr.isVolumePopulated = false
  dMgr.actualDirFileInfo = FileInfoPlus{}

}

// Equal - Compares two DirMgr objects to determine if
// they are equal.
func (dMgr *DirMgr) Equal(dmgr2 *DirMgr) bool {

  if dMgr.isInitialized != dmgr2.isInitialized ||
    dMgr.originalPath != dmgr2.originalPath ||
    dMgr.path != dmgr2.path ||
    dMgr.isPathPopulated != dmgr2.isPathPopulated ||
    dMgr.doesPathExist != dmgr2.doesPathExist ||
    dMgr.parentPath != dmgr2.parentPath ||
    dMgr.isParentPathPopulated != dmgr2.isParentPathPopulated ||
    dMgr.relativePath != dmgr2.relativePath ||
    dMgr.isRelativePathPopulated != dmgr2.isRelativePathPopulated ||
    dMgr.absolutePath != dmgr2.absolutePath ||
    dMgr.isAbsolutePathPopulated != dmgr2.isAbsolutePathPopulated ||
    dMgr.doesAbsolutePathExist != dmgr2.doesAbsolutePathExist ||
    dMgr.isAbsolutePathDifferentFromPath != dmgr2.isAbsolutePathDifferentFromPath ||
    dMgr.directoryName != dmgr2.directoryName ||
    dMgr.volumeName != dmgr2.volumeName ||
    dMgr.isVolumePopulated != dmgr2.isVolumePopulated {

    return false
  }

  if !dMgr.actualDirFileInfo.Equal(&dmgr2.actualDirFileInfo) {
    return false
  }

  return true
}

// EqualAbsPaths - compares the absolute paths for the current
// directory manager and the input directory manager ('dMgr2').
// If the two absolute paths are equal, the method returns 'true'.
// If the two absolute paths are NOT equal, the method returns 'false'.
// The comparison is NOT case sensitive. In other words, both paths
// are converted to lower case before making the comparision.
//
// If either the current DirMgr ('dMgr') or the input parameter
// 'dMgr2' are uninitialized, a value of 'false' is returned.
//
func (dMgr *DirMgr) EqualAbsPaths(dMgr2 *DirMgr) bool {


  if !dMgr.isInitialized || !dMgr2.isInitialized {
    return false
  }

  lcDMgrPath := strings.ToLower(dMgr.absolutePath)
  lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  return true
}

// EqualPaths - Compares two DirMgr objects to determine
// if their paths are equal. Both Directory Path and
// absolute path must be equivalent.
//
// If the compared paths are equal, the method returns 'true'.
// If the paths are NOT equal, the method returns 'false'.
// The comparisons are NOT case sensitive. In other words, all paths
// are converted to lower case before making the comparisons.
//
// If either the current DirMgr ('dMgr') or the input parameter
// 'dMgr2' are uninitialized, a value of 'false' is returned.
//
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

  if !dMgr.isInitialized || !dMgr2.isInitialized {
    return false
  }

  lcDMgrPath := strings.ToLower(dMgr.absolutePath)
  lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  lcDMgrPath = strings.ToLower(dMgr.path)
  lcDMgr2Path = strings.ToLower(dMgr2.path)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  return true
}

// ExecuteDirectoryFileOps - Performs a a file operation on specified 'selected' files
// in the current directory ONLY. This function does NOT perform operations on the
// sub directories (a.k.a. the directory tree).
//
// To perform file operations on the entire Directory Tree, see Function 'ExecuteDirectoryTreeOps()',
// above.
//
// The types of File Operations performed are generally classified as 'file copy' and
// 'file deletion' operations. The precise file operation applied is defined by the type,
// 'FileOperationCode' which provides a series of constants used to identify the specific file
// operation applied.
//
// Input parameter, 'fileOps' is an array of type 'FileOperationCode' elements. Multiple file
// operations can be applied to a single file. For instance, a 'copy source to destination'
// operation can be followed by a 'delete source file' operation.
//
// The 'selected' files are identified by input parameter 'fileSelectCriteria' of type
// 'FileSelectionCriteria'. This file selection criteria is compared against all files
// in the directory (NOT the Directory Tree) identified by the current 'DirMgr' instance.
// When a match is found, that file is treated as a 'selected' source file and designated
// file operations are performed on that file.
//
// The results or final output from file operations utilizes the final input parameter,
// 'targetBaseDir' of type DirMgr. File operations are applied to selected source files
// and generated output is created in the 'targetBaseDir'.  For example 'copy' or 'move'
// file operations will transfer source files to 'targetBaseDir'.
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// This method performs File Operations ONLY on the directory
// identified by the current DirMgr instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  fileSelectCriteria FileSelectionCriteria
//
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be identified
//    as 'Selected Files'. The specified File Operations (fileOps) will be
//    performed on these selected files.
//
//    type FileSelectionCriteria struct {
//       FileNamePatterns    []string     // An array of strings containing File Name Patterns
//       FilesOlderThan      time.Time    // Match files with older modification date times
//       FilesNewerThan      time.Time    // Match files with newer modification date times
//       SelectByFileMode   FilePermissionConfig  // Match by file mode (os.FileMode).
//       SelectCriterionMode FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//
//    type FileSelectionCriteria struct {
//     FileNamePatterns    []string// An array of strings containing File Name Patterns
//     FilesOlderThan      time.Time// Match files with older modification date times
//     FilesNewerThan      time.Time// Match files with newer modification date times
//     SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//     SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//    FileNamePatterns    []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file
//                                    modification date times in order to determine
//                                    whether the file is older than the 'FilesOlderThan'
//                                    file selection criterion. If the file modification
//                                    date time is older than the 'FilesOlderThan' date time,
//                                    that file is considered a 'match'	for this file selection
//                                    criterion.
//
//                                    If the value of 'FilesOlderThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    FilesNewerThan      time.Time - This date time type is compared to the file
//                                    modification date time in order to determine
//                                    whether the file is newer than the 'FilesNewerThan'
//                                    file selection criterion. If the file modification date time
//                                    is newer than the 'FilesNewerThan' date time, that file is
//                                    considered a 'match' for this file selection criterion.
//
//                                    If the value of 'FilesNewerThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    SelectByFileMode  FilePermissionConfig -
//                                    Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                    selection criterion allows for the selection of files by File Mode.
//                                    File modes are compared to the value	of 'SelectByFileMode'. If the
//                                    File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                    that file is considered to be a 'match' for this file selection
//                                    criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                         fsc := FileSelectionCriteria{}
//                                           err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                           err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//    SelectCriterionMode FileSelectCriterionMode -
//                                    This parameter selects the manner in which the file selection
//                                    criteria above are applied in determining a 'match' for file
//                                    selection purposes. 'SelectCriterionMode' may be set to one of
//                                    two constant values:
//
//                                    FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                    FileSelectMode.ORSelect() - File selected if any active selection
//                                      criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//       Example:
//
//             fsc := FileSelectCriterionMode{}
//
//             In this example, 'fsc' is NOT initialized. Therefore,
//             all of the selection criterion are 'Inactive'. Consequently,
//             all of the files encountered in the target directory during
//             the search operation will be selected and returned as
//             'Found Files'.
//
//
//  fileOps []FileOperationCode - An array of file operations to be performed
//                                on each selected file. Selected files are
//                                identified by matching the file selection
//                                criteria specified by input parameter,
//                                'fileSelectCriteria'. See above.
//
//          The FileOperationCode type consists of the following
//          constants.
//
//          FileOperationCode(0).MoveSourceFileToDestination() FileOperationCode = iota
//            Moves the source file to the destination file and
//            then deletes the original source file
//
//          FileOperationCode(0).DeleteDestinationFile()
//            Deletes the Destination file if it exists
//
//          FileOperationCode(0).DeleteSourceFile()
//            Deletes the Source file if it exists
//
//          FileOperationCode(0).DeleteSourceAndDestinationFiles
//            Deletes both the Source and Destination files
//            if they exist.
//
//          FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//            Copies the Source File to the Destination
//            using two copy attempts. The first copy is
//            by Hard Link. If the first copy attempt fails,
//            a second copy attempt is initiated/ by creating
//            a new file and copying the contents by 'io.Copy'.
//            An error is returned only if both copy attempts
//            fail. The source file is unaffected.
//
//            See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//          FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//            Copies the Source File to the Destination
//            using two copy attempts. The first copy is
//            by 'io.Copy' which creates a new file and copies
//            the contents to the new file. If the first attempt
//            fails, a second copy attempt is initiated using
//            'copy by hard link'. An error is returned only
//            if both copy attempts fail. The source file is
//            unaffected.
//
//            See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//          FileOperationCode(0).CopySourceToDestinationByHardLink()
//            Copies the Source File to the Destination
//            using one copy mode. The only copy attempt
//            utilizes 'Copy by Hard Link'. If this fails
//            an error is returned.  The source file is
//            unaffected.
//
//          FileOperationCode(0).CopySourceToDestinationByIo()
//            Copies the Source File to the Destination
//            using only one copy mode. The only copy
//            attempt is initiated using 'Copy by IO' or
//            'io.Copy'.  If this fails an error is returned.
//            The source file is unaffected.
//
//          FileOperationCode(0).CreateSourceDir()
//            Creates the Source Directory
//
//          FileOperationCode(0).CreateSourceDirAndFile()
//            Creates the Source Directory and File
//
//          FileOperationCode(0).CreateSourceFile()
//            Creates the Source File
//
//          FileOperationCode(0).CreateDestinationDir()
//            Creates the Destination Directory
//
//          FileOperationCode(0).CreateDestinationDirAndFile()
//            Creates the Destination Directory and File
//
//          FileOperationCode(0).CreateDestinationFile()
//            Creates the Destination File
//
//
// ------------------------------------------------------------------------
//
// Input parameters (continued)
//
//  targetBaseDir - The file selection criteria, 'fileSelectCriteria', and
//                  the File Operations, 'fileOps' are applied to files in
//                  the target base directory. This input parameter is of
//                  type 'DirMgr'.
//
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//    [] error - This function will return an array of errors containing error messages
//                generated during the performance of specified File Operations on the
//                designated directory. If the error array returned is empty or has
//                a zero length, it signals that no errors were encountered and all operations
//                completed successfully.
//
func (dMgr *DirMgr) ExecuteDirectoryFileOps(
  fileSelectCriteria FileSelectionCriteria,
  fileOps []FileOperationCode,
  targetBaseDir DirMgr) (errs []error) {

  errs = make([]error, 0, 500)

  ePrefix := "DirMgr.ExecuteDirectoryFileOps() "

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    err2 := fmt.Errorf("%v ", err.Error())
    errs = append(errs, err2)
    return errs
  }

  err = targetBaseDir.IsDirMgrValid("")

  if err != nil {

    err2 := fmt.Errorf(ePrefix+
      "Input parameter 'targetBaseDir' is INVALID!. Error='%v' ",
      err.Error())

    errs = append(errs, err2)

    return errs
  }

  if len(fileOps) == 0 {

    err2 := errors.New(ePrefix +
      "Error: The input parameter 'fileOps' is a ZERO LENGTH ARRAY!")

    errs = append(errs, err2)

    return errs
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {
    var err2 error

    if os.IsNotExist(err) {
      err2 = fmt.Errorf(ePrefix + "ERROR: Source Directory does NOT EXIST!\n" +
        "Source Directory='%v'\n", dMgr.absolutePath)
    } else {
      err2 = fmt.Errorf(ePrefix +
        "Source Directory returned a non-path error from os.Stat().\n" +
        "Source Directory='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())
    }

    errs = append(errs, err2)
    return errs
  }

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {
    err2 := fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  nameFileInfos, err := dir.Readdir(-1)

  if err != nil {
    _ = dir.Close()
    err2 := fmt.Errorf(ePrefix+
      "Error returned by dir.Readdirnames(-1). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  fh := FileHelper{}

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue
    }

    // Must be a file - process it!

    // This is not a directory. It is a file.
    // Determine if it matches the find file criteria.
    isMatch, err := fh.FilterFileName(nameFInfo, fileSelectCriteria)

    if err != nil {

      _ = dir.Close()

      err2 := fmt.Errorf(ePrefix +
        "Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria).\n"+
        "directorySearched='%v'\nfileName='%v'\nError='%v'\n",
        dMgr.absolutePath, nameFInfo.Name(), err.Error())

      errs = append(errs, err2)
      return errs
    }

    if !isMatch {

      continue

    }

    // Must be a match - this is a 'selected' file!

    srcFileNameExt := nameFInfo.Name()

    fileOp, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
      dMgr.GetAbsolutePath(),
      nameFInfo.Name(),
      targetBaseDir.GetAbsolutePath(),
      srcFileNameExt)

    if err != nil {

      _ = dir.Close()

      err2 := fmt.Errorf(ePrefix+
        "Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs()\n"+
        "sourcePath='%v'\nsrcFileNameExt='%v'\ndestDir='%v'\nError='%v'\n",
        dMgr.GetAbsolutePath(), srcFileNameExt, targetBaseDir.GetAbsolutePath(),
        err.Error())
      errs = append(errs, err2)
      return errs
    }

    maxOps := len(fileOps)

    for i := 0; i < maxOps; i++ {

      err = fileOp.ExecuteFileOperation(fileOps[i])

      if err != nil {
        err2 := fmt.Errorf(ePrefix+
          "Error returned by fileOp.ExecuteFileOperation(fileOps[%v]). "+
          "FileOps='%v'\nError='%v'\n\n",
          i, fileOps[i].String(), err.Error())

        // Store the error and continue processing
        // file operations.
        errs = append(errs, err2)
      }
    }

    // finished applying file operations to this file.
    // Get another one and continue...
  }

  _ = dir.Close()

  return errs
}

// ExecuteDirectoryTreeOps - Performs File Operations on specified 'selected'
// files in the directory tree identified by the current 'DirMgr' instance.
// The 'DirMgr' path therefore serves as the parent directory for file operations
// performed on the directory tree.
//
// If you wish to perform File Operations ONLY on the current directory and
// NOT THE ENTIRE DIRECTORY TREE, see Function "ExecuteDirectoryFileOps(), above.
//
// The types of File Operations performed are generally classified as 'file copy'
// and 'file deletion' operations. The precise file operation applied is defined
// by the the type, 'FileOperationCode' which provides a series of constants, or
// enumerations, used to identify the specific file operation applied. Input
// parameter, 'fileOps' is an array of type 'FileOperationCode' elements. Multiple
// file operations can be applied to a single file. For instance, a 'copy source to
// destination' operation can be followed by a 'delete source file' operation.
//
// The 'selected' files are identified by input parameter 'fileSelectCriteria' of
// type 'FileSelectionCriteria'. This file selection criteria is compared against
// all files in the current directory tree identified by the current 'DirMgr'
// instance. When a match is found, that file is treated as a 'selected' file and
// designated file operations are performed on that file.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// This method performs File Operations on THE ENTIRE DIRECTORY
// TREE identified by this DirMgr instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  fileSelectCriteria FileSelectionCriteria
//
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be identified
//    as 'Selected Files'. The specified File Operations (fileOps) will be
//    performed on these selected files.
//
//    type FileSelectionCriteria struct {
//      FileNamePatterns     []string	    // An array of strings containing File Name Patterns
//      FilesOlderThan       time.Time    // Match files with older modification date times
//      FilesNewerThan       time.Time    // Match files with newer modification date times
//      SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//                                        //   is set to 'false'.
//      SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//
//    type FileSelectionCriteria struct {
//     FileNamePatterns    []string// An array of strings containing File Name Patterns
//     FilesOlderThan      time.Time// Match files with older modification date times
//     FilesNewerThan      time.Time// Match files with newer modification date times
//     SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//     SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//    FileNamePatterns    []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file
//                                    modification date times in order to determine
//                                    whether the file is older than the 'FilesOlderThan'
//                                    file selection criterion. If the file modification
//                                    date time is older than the 'FilesOlderThan' date time,
//                                    that file is considered a 'match'	for this file selection
//                                    criterion.
//
//                                    If the value of 'FilesOlderThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    FilesNewerThan      time.Time - This date time type is compared to the file
//                                    modification date time in order to determine
//                                    whether the file is newer than the 'FilesNewerThan'
//                                    file selection criterion. If the file modification date time
//                                    is newer than the 'FilesNewerThan' date time, that file is
//                                    considered a 'match' for this file selection criterion.
//
//                                    If the value of 'FilesNewerThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    SelectByFileMode  FilePermissionConfig -
//                                    Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                    selection criterion allows for the selection of files by File Mode.
//                                    File modes are compared to the value	of 'SelectByFileMode'. If the
//                                    File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                    that file is considered to be a 'match' for this file selection
//                                    criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                         fsc := FileSelectionCriteria{}
//                                           err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                           err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//    SelectCriterionMode FileSelectCriterionMode -
//                                    This parameter selects the manner in which the file selection
//                                    criteria above are applied in determining a 'match' for file
//                                    selection purposes. 'SelectCriterionMode' may be set to one of
//                                    two constant values:
//
//                                    FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                    FileSelectMode.ORSelect() - File selected if any active selection
//                                      criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//      Example:
//           fsc := FileSelectCriterionMode{}
//
//           In this example, 'fsc' is NOT initialized. Therefore,
//           all of the selection criterion are 'Inactive'. Consequently,
//           all of the files encountered in the target directory during
//           the search operaiton will be selected and returned as
//           'Found Files'.
//
// ---------------------------------------------------------------------------
//
//  fileOps []FileOperationCode - An array of file operations to be performed
//                                on each selected file. Selected files are
//                                identified by matching the file selection
//                                criteria specified by input parameter,
//                                'fileSelectCriteria'. See above.
//
//    The FileOperationCode type consists of the following
//    constants.
//
//    FileOperationCode(0).None()
//      No Action
//
//    FileOperationCode(0).MoveSourceFileToDestination()
//      Moves the source file to the destination file and
//      then deletes the original source file
//
//    FileOperationCode(0).DeleteDestinationFile()
//      Deletes the Destination file if it exists
//
//    FileOperationCode(0).DeleteSourceFile()
//      Deletes the Source file if it exists
//
//    FileOperationCode(0).DeleteSourceAndDestinationFiles
//      Deletes both the Source and Destination files
//      if they exist.
//
//    FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//      Copies the Source File to the Destination
//      using two copy attempts. The first copy is
//      by Hard Link. If the first copy attempt fails,
//      a second copy attempt is initiated/ by creating
//      a new file and copying the contents by 'io.Copy'.
//      An error is returned only if both copy attempts
//      fail. The source file is unaffected.
//
//      See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//    FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//      Copies the Source File to the Destination
//      using two copy attempts. The first copy is
//      by 'io.Copy' which creates a new file and copies
//      the contents to the new file. If the first attempt
//      fails, a second copy attempt is initiated using
//      'copy by hard link'. An error is returned only
//      if both copy attempts fail. The source file is
//      unaffected.
//
//      See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//    FileOperationCode(0).CopySourceToDestinationByHardLink()
//      Copies the Source File to the Destination
//      using one copy mode. The only copy attempt
//      utilizes 'Copy by Hard Link'. If this fails
//      an error is returned.  The source file is
//      unaffected.
//
//    FileOperationCode(0).CopySourceToDestinationByIo()
//      Copies the Source File to the Destination
//      using only one copy mode. The only copy
//      attempt is initiated using 'Copy by IO' or
//      'io.Copy'.  If this fails an error is returned.
//      The source file is unaffected.
//
//    FileOperationCode(0).CreateSourceDir()
//      Creates the Source Directory
//
//    FileOperationCode(0).CreateSourceDirAndFile()
//      Creates the Source Directory and File
//
//    FileOperationCode(0).CreateSourceFile()
//      Creates the Source File
//
//    FileOperationCode(0).CreateDestinationDir()
//      Creates the Destination Directory
//
//    FileOperationCode(0).CreateDestinationDirAndFile()
//      Creates the Destination Directory and File
//
//    FileOperationCode(0).CreateDestinationFile()
//      Creates the Destination File
//
// ------------------------------------------------------------------------
//
// Input parameters (continued)
//
//
// targetBaseDir -  The file selection criteria, 'fileSelectCriteria', and
//                  the File Operations, 'fileOps' are applied to files in
//                  the target base directory. This input parameter is of
//                  type 'DirMgr'.
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//  []error - This function will return an array of strings containing error messages
//             generated during the performance of specified File Operations on the
//             designated directory tree. If the string array returned is empty or has
//             a zero length, it signals that no errors were encountered and all operations
//             completed successfully.
//
func (dMgr *DirMgr) ExecuteDirectoryTreeOps(
  fileSelectCriteria FileSelectionCriteria,
  fileOps []FileOperationCode,
  targetBaseDir DirMgr) (errs []error) {

  ePrefix := "DirMgr.ExecuteDirectoryTreeOps() "

  dirOp := DirTreeOp{}.New()
  dirOp.CallingFunc = ePrefix
  errs = make([]error, 0, 300)


  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    err2 := fmt.Errorf("%v ", err.Error())
    errs = append(errs, err2)
    return errs
  }

  err = targetBaseDir.IsDirMgrValid("")

  if err != nil {

    err2 := fmt.Errorf(ePrefix+
      "Input parameter 'targetBaseDir' is INVALID!. Error='%v' ",
      err.Error())

    errs = append(errs, err2)

    return errs
  }

  if len(fileOps) == 0 {

    err2 := errors.New(ePrefix +
      "Error: The input parameter 'fileOps' is a ZERO LENGTH ARRAY!\n")

    errs = append(errs, err2)

    return errs
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {
    var err2 error

    if os.IsNotExist(err) {
      err2 = fmt.Errorf(ePrefix + "DirMgr Source Directory does NOT Exist!\n" +
        "Source Directory='%v'\n", dMgr.absolutePath)
    } else {
      err2 = fmt.Errorf(ePrefix +
        "DirMgr Source Directory returned a non-path error from os.Stat().\n" +
        "Source Directory='%v'\nError='%v'", dMgr.absolutePath, err.Error())
    }

    errs = append(errs, err2)
    return errs

  }

  dirOp.FileOps = append(dirOp.FileOps, fileOps...)
  dirOp.TargetBaseDir = targetBaseDir.CopyOut()
  dirOp.SourceBaseDir = dMgr.CopyOut()
  dirOp.FileSelectCriteria = fileSelectCriteria

  err = fp.Walk(dMgr.GetAbsolutePath(), dMgr.executeFileOpsOnFoundFiles(&dirOp))

  if err != nil {
    err2 := fmt.Errorf(ePrefix +
      "Error returned by fp.Walk(). Error='%v' ", err.Error())
    errs = append(errs, dirOp.ErrReturns ...)
    errs = append(errs, err2)
    return errs
  }

  return dirOp.ErrReturns
}

// FindFilesByNamePattern - Searches files in the current directory ONLY. An attempt
// will be made to match the file name with the specified search pattern string.
// All matched files will be returned in a FileMgrCollection.
//
// Regardless of the search pattern used, this method will never return sub-directories
// of the target search directory.
//
// Again, the file search will always be limited to the directory identified by the
// current DirMgr instance. No sub-directories will be searched.
//
// If the 'fileSearchPattern' is an empty string or improperly formatted, an error
// will be returned.
//
// ------------------------------------------------------------------------
//
// Input parameter
//
//
//  fileSearchPattern  string -  The fileSearchPattern is string containing
//                               parameters used to select target files in
//                               directory identified by the 'DirMgr' type.
//
//                               Example 'fileSearchPattern' strings
//
//                               *.*             will match all files in directory.
//                               *.html          will match  anyfilename.html
//                               a*              will match  appleJack.txt
//                               j????row.txt    will match  j1x34row.txt
//                               data[0-9]*      will match 	data123.csv
//
//                               Reference For File Pattern Matching Details:
//                                 https://golang.org/pkg/path/filepath/#Match
//
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//  FileMgrCollection - If this method completes successfully without error, the
//                      returned FileMgrCollection type will contain an array of
//                      FileMgr types identifying each of the files matched by
//                      input parameter, 'fileSearchPattern'.
//
//  error             - If this method completes successfully, this return value
//                      will be set to 'nil'. Otherwise, a valid error message will
//                      be encapsulated in the returned type 'error'.
//
func (dMgr *DirMgr) FindFilesByNamePattern(fileSearchPattern string) (FileMgrCollection, error) {

  ePrefix := "DirMgr.FindFilesByNamePattern() "

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return FileMgrCollection{}, err
  }

  errCode := 0

  errCode, _, fileSearchPattern = FileHelper{}.isStringEmptyOrBlank(fileSearchPattern)

  if errCode < 0 {
    return FileMgrCollection{},
      errors.New(ePrefix + "Input parameter 'fileSearchPattern' is INVALID!\n" +
        "'fileSearchPattern' is an EMPTY STRING!\n")
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {

    if os.IsNotExist(err) {
      return FileMgrCollection{},
      fmt.Errorf(ePrefix + "DirMgr path does NOT Exist!\n" +
        "DirMgr='%v'\n", dMgr.absolutePath)
    }

    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "DirMgr path returned a non-path error from os.Stat().\n" +
        "DirMgr='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())

  }

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error return by os.Open(dMgr.absolutePath). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  nameFileInfos, err := dir.Readdir(-1)

  if err != nil {
    _ = dir.Close()
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(-1). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  fMgrCol := FileMgrCollection{}

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue

    } else {

      fName := nameFInfo.Name()

      isMatch, err := fp.Match(fileSearchPattern, fName)

      if err != nil {

        _ = dir.Close()

        return FileMgrCollection{},
          fmt.Errorf(ePrefix+
            "Error returned by fp.Match(fileSearchPattern, fileName). "+
            "directorySearched='%v' fileSearchPattern='%v' fileName='%v' Error='%v' ",
            dMgr.absolutePath, fileSearchPattern, fName, err.Error())
      }

      if !isMatch {
        continue
      } else {

        err = fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo)

        if err != nil {
          _ = dir.Close()
          return FileMgrCollection{},
            fmt.Errorf(ePrefix+
              "Error returned by fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo). "+
              "Directory='%v' FileName='%v' Error='%v' ",
              dMgr.absolutePath, fName, err.Error())
        }
      }
    }
  }

  err = dir.Close()

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error returned by dir.Close(). "+
        "dir='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  return fMgrCol, nil
}

// FindFilesBySelectCriteria - Conducts a file search in the directory
// identified by the current DirMgr instance. The file search is limited
// to that directory ONLY. No sub-directories will be searched.
//
// Files matching the "FileSectionCriteria" instance passed as an input
// parameter will be used to screen available files. Any files matching
// the file selection criteria will be returned in a 'FileMgrCollection'.
//
// Only matched files will be returned. No sub-directory names will ever
// be included.
//
// The use of a 'FileSelectionCriteria' structure allows for very flexible
// and granular file searches.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  fileSelectCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be returned as
//    'Found Files'.
//
//    type FileSelectionCriteria struct {
//     FileNamePatterns    []string// An array of strings containing File Name Patterns
//     FilesOlderThan      time.Time// Match files with older modification date times
//     FilesNewerThan      time.Time// Match files with newer modification date times
//     SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//     SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//    FileNamePatterns    []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file
//                                    modification date times in order to determine
//                                    whether the file is older than the 'FilesOlderThan'
//                                    file selection criterion. If the file modification
//                                    date time is older than the 'FilesOlderThan' date time,
//                                    that file is considered a 'match'	for this file selection
//                                    criterion.
//
//                                    If the value of 'FilesOlderThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    FilesNewerThan      time.Time - This date time type is compared to the file
//                                    modification date time in order to determine
//                                    whether the file is newer than the 'FilesNewerThan'
//                                    file selection criterion. If the file modification date time
//                                    is newer than the 'FilesNewerThan' date time, that file is
//                                    considered a 'match' for this file selection criterion.
//
//                                    If the value of 'FilesNewerThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    SelectByFileMode  FilePermissionConfig -
//                                    Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                    selection criterion allows for the selection of files by File Mode.
//                                    File modes are compared to the value	of 'SelectByFileMode'. If the
//                                    File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                    that file is considered to be a 'match' for this file selection
//                                    criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                         fsc := FileSelectionCriteria{}
//                                           err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                           err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//    SelectCriterionMode FileSelectCriterionMode -
//                                    This parameter selects the manner in which the file selection
//                                    criteria above are applied in determining a 'match' for file
//                                    selection purposes. 'SelectCriterionMode' may be set to one of
//                                    two constant values:
//
//                                    FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                    FileSelectMode.ORSelect() - File selected if any active selection
//                                      criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//      Example:
//        fsc := FileSelectCriterionMode{}
//
//        In this example, 'fsc' is NOT initialized. Therefore,
//        all of the selection criterion are 'Inactive'. Consequently,
//        all of the files encountered in the target directory during
//        the search operation will be selected and returned as
//        'Found Files'.
//
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//  FileMgrCollection - If this method completes successfully without error, the
//                      returned FileMgrCollection type will contain an array of
//                      FileMgr types identifying each of the files matched by
//                      input parameter, 'fileSelectCriteria'.
//
//  error             - If this method completes successfully, this return value
//                      will be set to 'nil'. Otherwise, a valid error message will
//                      be encapsulated in the returned type 'error'.
//
func (dMgr *DirMgr) FindFilesBySelectCriteria(
  fileSelectCriteria FileSelectionCriteria) (FileMgrCollection, error) {

  ePrefix := "DirMgr.FindFilesBySelectCriteria() "

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return FileMgrCollection{}, err
  }

  _, err = os.Stat(dMgr.absolutePath)

  if err != nil {

    if os.IsNotExist(err) {
      return FileMgrCollection{},
        fmt.Errorf(ePrefix + "DirMgr path does NOT Exist!\n" +
          "DirMgr='%v'\n", dMgr.absolutePath)
    }

    return FileMgrCollection{},
      fmt.Errorf(ePrefix +
        "DirMgr path returned a non-path error from os.Stat().\n" +
        "DirMgr='%v'\nError='%v'\n", dMgr.absolutePath, err.Error())

  }

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error return by os.Open(dMgr.absolutePath). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  nameFileInfos, err := dir.Readdir(-1)

  if err != nil {
    _ = dir.Close()
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(-1). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  fMgrCol := FileMgrCollection{}
  fh := FileHelper{}

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue

    } else {

      // This is not a directory. It is a file.
      // Determine if it matches the find file criteria.
      isMatch, err := fh.FilterFileName(nameFInfo, fileSelectCriteria)

      if err != nil {

        _ = dir.Close()

        return FileMgrCollection{},
          fmt.Errorf(ePrefix+
            "Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
            "directorySearched='%v'  fileName='%v' Error='%v' ",
            dMgr.absolutePath, nameFInfo.Name(), err.Error())
      }

      if !isMatch {

        continue

      } else {

        err = fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo)

        if err != nil {
          _ = dir.Close()
          return FileMgrCollection{},
            fmt.Errorf(ePrefix+
              "Error returned by fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo). "+
              "Directory='%v' FileName='%v' Error='%v' ",
              dMgr.absolutePath, nameFInfo.Name(), err.Error())
        }
      }
    }
  }

  err = dir.Close()

  if err != nil {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "Error returned by dir.Close(). "+
        "dir='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())
  }

  return fMgrCol, nil
}

// FindWalkDirFiles - This method returns file information on files residing in a specific
// directory tree identified by the current DirMgr object.
//
// This method 'walks the directory tree' locating all files in the directory tree which match
// the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. By the way, if ALL the file selection criterion are set to
// zero values or 'Inactive', then ALL FILES in the directory are selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  fileSelectCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be returned as
//    'Found Files'.
//
//    type FileSelectionCriteria struct {
//      FileNamePatterns     []string     // An array of strings containing File Name Patterns
//      FilesOlderThan       time.Time    // Match files with older modification date times
//      FilesNewerThan       time.Time    // Match files with newer modification date times
//      SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//
//    type FileSelectionCriteria struct {
//     FileNamePatterns    []string// An array of strings containing File Name Patterns
//     FilesOlderThan      time.Time// Match files with older modification date times
//     FilesNewerThan      time.Time// Match files with newer modification date times
//     SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//     SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//    }
//
//    The FileSelectionCriteria type allows for configuration of single or multiple file
//    selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//    file must match all, or any one, of the active file selection criterion.
//
//    Elements of the FileSelectionCriteria are described below:
//
//    FileNamePatterns    []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file
//                                    modification date times in order to determine
//                                    whether the file is older than the 'FilesOlderThan'
//                                    file selection criterion. If the file modification
//                                    date time is older than the 'FilesOlderThan' date time,
//                                    that file is considered a 'match'	for this file selection
//                                    criterion.
//
//                                    If the value of 'FilesOlderThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    FilesNewerThan      time.Time - This date time type is compared to the file
//                                    modification date time in order to determine
//                                    whether the file is newer than the 'FilesNewerThan'
//                                    file selection criterion. If the file modification date time
//                                    is newer than the 'FilesNewerThan' date time, that file is
//                                    considered a 'match' for this file selection criterion.
//
//                                    If the value of 'FilesNewerThan' is set to time zero,
//                                    the default value for type time.Time{}, then this
//                                    file selection criterion is considered to be 'Inactive'
//                                    or 'Not Set'.
//
//    SelectByFileMode  FilePermissionConfig -
//                                    Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                    selection criterion allows for the selection of files by File Mode.
//                                    File modes are compared to the value	of 'SelectByFileMode'. If the
//                                    File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                    that file is considered to be a 'match' for this file selection
//                                    criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                         fsc := FileSelectionCriteria{}
//                                           err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                           err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//    SelectCriterionMode FileSelectCriterionMode -
//                                    This parameter selects the manner in which the file selection
//                                    criteria above are applied in determining a 'match' for file
//                                    selection purposes. 'SelectCriterionMode' may be set to one of
//                                    two constant values:
//
//                                    FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                    FileSelectMode.ORSelect() - File selected if any active selection
//                                      criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//       Example:
//            fsc := FileSelectionCriteria{} // fsc is NOT initialized
//
//            In this example, all of the selection criterion are
//            'Inactive' and therefore all of the files encountered
//            in the target directory will be selected and returned
//            as 'Found Files'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  DirectoryTreeInfo structure	-
//          type DirectoryTreeInfo struct {
//            StartPath             string                // The starting path or directory for the file search
//            Directories           DirMgrCollection      // dirMgrs found during directory tree search
//            FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//            ErrReturns            []string              // Internal System errors encountered
//            FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//                                                        // input parameter to this method.
//          }
//
//          If successful, files matching the file selection criteria input
//          parameter shown above will be returned in a 'DirectoryTreeInfo'
//          object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//          on all the files in the specified directory tree which match the file selection
//          criteria.
//
//          Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//                to determine if any internal system errors were encountered while processing
//                the directory tree.
//
//  error - If a program execution error is encountered during processing, it will
//          be returned as an 'error' type. Also, see the comment on
//          'DirectoryTreeInfo.ErrReturns', above.
//
func (dMgr *DirMgr) FindWalkDirFiles(
  fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

  ePrefix := "DirMgr.GetWalkDirInfo() "
  findFilesInfo := DirectoryTreeInfo{}

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return findFilesInfo, err
  }

  if dMgr.doesAbsolutePathExist {

    findFilesInfo.StartPath = dMgr.absolutePath

  } else if dMgr.DoesDirMgrPathExist() {

    findFilesInfo.StartPath = dMgr.path

  } else {

    return findFilesInfo,
      fmt.Errorf(ePrefix+
        "path and absolutePath - PATH DOES NOT EXIST! "+
        "dMgr.absolutePath='%v' dMgr.path='%v'", dMgr.absolutePath, dMgr.path)
  }

  findFilesInfo.FileSelectCriteria = fileSelectCriteria

  fh := FileHelper{}

  err = fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

  if err != nil {
    return findFilesInfo, fmt.Errorf(ePrefix+
      "Error returned by FileHelper.FindFilesWalkDirectory(&dWalkInfo). "+
      "dWalkInfo.StartPath='%v' Error='%v' ", findFilesInfo.StartPath, err.Error())
  }

  return findFilesInfo, nil
}

// GetAbsolutePath - Returns a string containing the
// absolute path for the current Directory Manager
// instance. This string returned by this method
// will NOT have a trailing path separator.
//
func (dMgr *DirMgr) GetAbsolutePath() string {

  err := dMgr.IsDirMgrValid("")

  if err != nil {
    return ""
  }

  return dMgr.absolutePath
}

// GetAbsolutePathElements - Returns all of the directories and drive
// specifications as an array of strings.
//
// Example
//
// Path = "D:\ADir\BDir\CDir\EDir"
//
// Returned pathElements string array:
//   pathElements[0] = "D:"
//   pathElements[1] = "ADir"
//   pathElements[2] = "BDir"
//   pathElements[3] = "CDir"
//   pathElements[4] = "DDir"
//   pathElements[4] = "EDir"
//
func (dMgr *DirMgr) GetAbsolutePathElements() (pathElements []string) {

  pathElements = make([]string, 0, 50)

  pathStr := dMgr.GetAbsolutePath()

  if len(pathStr) == 0 {
    return pathElements
  }

  pathStr = strings.Replace(pathStr, "\\", "/", -1)

  pathElements = strings.Split(pathStr, "/")

  return pathElements
}

// GetAbsolutePathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetAbsolutePathWithSeparator() string {
  lPath := len(dMgr.absolutePath)

  if lPath == 0 {
    return ""
  }

  if dMgr.absolutePath[lPath-1] != os.PathSeparator {
    return dMgr.absolutePath + string(os.PathSeparator)
  }

  return dMgr.absolutePath
}

// GetDirectoryName - Returns a string containing the name
// of the directory without out the parent path.
//
func (dMgr *DirMgr) GetDirectoryName() string {
  return dMgr.directoryName
}

// GetFileInfoPlus - Returns a FileInfoPlus instance detailing file
// system information on the directory identified by the current
// Directory Manager instance.
//
func (dMgr *DirMgr) GetFileInfoPlus() (FileInfoPlus, error) {

  ePrefix := "DirMgr.GetFileInfoPlus() "

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return FileInfoPlus{}, err
  }

  if dMgr.absolutePath == "" {

    dMgr.doesAbsolutePathExist = false

    return FileInfoPlus{},
      errors.New(ePrefix + "Directory DOES NOT EXIST! FileInfoPlus is NOT available.")
  }

  info, err := os.Stat(dMgr.absolutePath)

  if err != nil {

    return FileInfoPlus{},
      fmt.Errorf(ePrefix+
        "Error returned by os.Stat(dMgr.absolutePath). Error='%v'",
        err.Error())
  }

  dMgr.doesAbsolutePathExist = true
  dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.absolutePath, info)

  return dMgr.actualDirFileInfo.CopyOut(), nil
}

// GetDirPermissionTextCodes -  - If the current directory exists on disk,
// this method will return the Directory Permission Codes, otherwise known
// as the unix permission bits, in the form of a 10-character string.
//
// If the current Directory does NOT exist, this method will return an error.
//
func (dMgr *DirMgr) GetDirPermissionTextCodes() (string, error) {

  ePrefix := "GetDirPermissionTextCodes() "

  if !dMgr.doesAbsolutePathExist {
    return "",
      errors.New(ePrefix +
        "The current directory does NOT exist. Therefore, permission codes " +
        "do NOT exist.")
  }

  if !dMgr.actualDirFileInfo.IsFInfoInitialized {
    return "",
      errors.New(ePrefix +
        "The FileInfo data for this Directory has NOT been initialized.")
  }

  fPerm, err := FilePermissionConfig{}.NewByFileMode(dMgr.actualDirFileInfo.Mode())

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "%v", err.Error())
  }

  permissionText, err := fPerm.GetPermissionTextCode()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "%v", err.Error())
  }

  return permissionText, nil
}

// GetNumberOfAbsPathElements - Returns the number of elements
// or path components in the absolute path of the current
// Directory Manager instance.
func (dMgr *DirMgr) GetNumberOfAbsPathElements() int {

  pathElements := dMgr.GetAbsolutePathElements()

  return len(pathElements)
}

// GetOriginalPath - Returns the original path used to initialize
// this Directory Manager instance.
//
func (dMgr *DirMgr) GetOriginalPath() string {
  return dMgr.originalPath
}

// GetParentDirMgr - Returns a new Directory Manager instance
// which represents the the parent path for the current
// Directory Manager. The current Directory Manager absolute
// path is used in extracting the parent Directory Manager.
//
// Return Values:
//
//	dirMgr   DirMgr - If successful, this methods returns a Directory Manager
//	                  which is a parent of the current Directory Manager.
//
//	hasParent  bool - If 'true', it signals that the current Directory Manager
//	                  has a valid parent directory. If 'false', it signals that
//	                  the current Directory Manager represents a top level directory
//	                  which has no parent directory. In that case a copy of the
//	                  current Directory will be returned.
//
//	err       error - If an error is encountered this error type will be populated
//	                  with an appropriate error message. Otherwise, a value of 'nil'
//	                  will be returned.
//
//	                  If 'hasParent' is 'false', no error will be returned.
//
func (dMgr *DirMgr) GetParentDirMgr() (dirMgr DirMgr, hasParent bool, err error) {

  ePrefix := "DirMgr.GetParentDirMgr() Error: "
  dirMgr = DirMgr{}
  hasParent = true
  err = nil

  if len(dMgr.parentPath) == 0 {

    return dMgr.CopyOut(), false, nil

  }

  var err2 error

  dirMgr, err2 = DirMgr{}.New(dMgr.parentPath)

  if err2 != nil {

    err = fmt.Errorf(ePrefix+"%v", err.Error())
    hasParent = true
    dirMgr = DirMgr{}
    return dirMgr, hasParent, err
  }

  err = nil

  return dirMgr, hasParent, err
}

// GetParentPath - Returns a string containing the
// parent path for the current Directory Manager
// instance. The Parent Path string will NOT contain
// a trailing path separator.
//
func (dMgr *DirMgr) GetParentPath() string {
  return dMgr.parentPath
}

// GetPath - Returns the path used to configure this
// Directory Manager Instance. It will NOT contain a
// trailing path separator. It may or may not be an
// absolute path.
//
func (dMgr *DirMgr) GetPath() string {
  return dMgr.path
}

// GetPathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetPathWithSeparator() string {
  lPath := len(dMgr.path)

  if lPath == 0 {
    return ""
  }

  if dMgr.path[lPath-1] != os.PathSeparator {
    return dMgr.path + string(os.PathSeparator)
  }

  return dMgr.path
}

// GetRelativePath - Returns a string containing the relative path
// for this Directory Manager instance. The relative path is derived
// through comparison to the current Directory Managers parent
// directory. Therefore, the relative path of the current Directory
// Manger is relative to its parent.
//
func (dMgr *DirMgr) GetRelativePath() string {
  return dMgr.relativePath
}

// GetThisDirectoryTree - Returns a DirMgrCollection containing all
// the directories in the path of the parent directory identified by
// the current DirMgr instance.
//
// The returned DirMgrCollection will always contain the parent directory
// and will therefore always consist of at least one directory. If
// sub-directories are found, then the returned DirMgrCollection will
// contain more than one directory.
//
func (dMgr *DirMgr) GetThisDirectoryTree() (DirMgrCollection, error) {

  ePrefix := "DirMgr.GetThisDirectoryTree() "

  dMgrs := DirMgrCollection{}

  err := dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return dMgrs, err
  }

  dMgrs.AddDirMgr(dMgr.CopyOut())

  fh := FileHelper{}

  maxLen := 1

  for i := 0; i < maxLen; i++ {

    dir, err := os.Open(dMgrs.dirMgrs[i].absolutePath)

    if err != nil {
      return DirMgrCollection{},
        fmt.Errorf(ePrefix+
          "Error return by os.Open(dMgrs.dirMgrs[i].absolutePath). "+
          "dMgr.absolutePath='%v' Error='%v' ",
          dMgrs.dirMgrs[i].absolutePath, err.Error())
    }

    nameFileInfos, err := dir.Readdir(-1)

    if err != nil {
      _ = dir.Close()
      return DirMgrCollection{},
        fmt.Errorf(ePrefix+
          "Error returned by dir.Readdirnames(-1). "+
          "dMgr.absolutePath='%v' Error='%v' ",
          dMgr.absolutePath, err.Error())
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        newDirPathFileName :=
          fh.JoinPathsAdjustSeparators(dMgrs.dirMgrs[i].absolutePath, nameFInfo.Name())

        fmt.Println("Next Dir: ", newDirPathFileName)

        // err = dMgrs.AddFileInfo(newDirPathFileName, nameFInfo)
        err = dMgrs.AddDirMgrByPathNameStr(newDirPathFileName)

        if err != nil {
          return DirMgrCollection{},
            fmt.Errorf(ePrefix+
              "Error returned by dMgrs.AddDirMgrByPathNameStr(newDirPathFileName). "+
              "dir='%v' Error='%v' ",
              newDirPathFileName, err.Error())
        }

        fmt.Println("dMgrs Length", dMgrs.GetNumOfDirs())

        maxLen++

      }
    }

    err = dir.Close()

    if err != nil {
      return DirMgrCollection{},
        fmt.Errorf(ePrefix+
          "Error returned by dir.Close(). "+
          "dir='%v' Error='%v' ",
          dMgr.absolutePath, err.Error())
    }

  }

  return dMgrs, nil
}

// GetVolumeName - Returns a string containing the volume name
// of the directory identified by the current Directory Manager
// instance.
//
func (dMgr *DirMgr) GetVolumeName() string {
  return dMgr.volumeName
}

// IsAbsolutePathDifferentFromPath - Returns a boolean value indicating
// whether the absolute path differs from the base path for this
// Directory Manager instance.
//
func (dMgr *DirMgr) IsAbsolutePathDifferentFromPath() bool {

  if dMgr.path != dMgr.absolutePath {
    dMgr.isAbsolutePathDifferentFromPath = true
  } else {
    dMgr.isAbsolutePathDifferentFromPath = false
  }

  return dMgr.isAbsolutePathDifferentFromPath
}

// IsAbsolutePathPopulated - Returns a boolean value indicating
// whether the absolute path for the current Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsAbsolutePathPopulated() bool {

  if len(dMgr.absolutePath) == 0 {
    dMgr.isAbsolutePathPopulated = false
  } else {
    dMgr.isAbsolutePathPopulated = true
  }

  return dMgr.isAbsolutePathPopulated
}

// IsDirMgrValid - This method examines the current DirMgr object
// to determine whether it has been properly configured.
// If the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error is
// returned.
func (dMgr *DirMgr) IsDirMgrValid(errPrefixStr string) error {

  ePrefix := strings.TrimRight(errPrefixStr, " ") + " DirMgr.IsDirMgrValid() "

  if !dMgr.isInitialized {
    return fmt.Errorf(ePrefix + "Error: DirMgr is NOT Initialized.")
  }

  dMgr.isAbsolutePathPopulated = false

  if dMgr.absolutePath == "" {
    return fmt.Errorf(ePrefix + "Error: DirMgr.absolutePath is EMPTY!.")
  }

  dMgr.isAbsolutePathPopulated = true

  dMgr.isPathPopulated = false

  if dMgr.path == "" {
    return fmt.Errorf(ePrefix + "Error: DirMgr.absolutePath is EMPTY!.")
  }

  dMgr.isPathPopulated = true

  _ = dMgr.DoesDirMgrAbsolutePathExist()
  _ = dMgr.DoesDirMgrPathExist()

  return nil
}

// IsInitialized - Returns a boolean value indicating
// whether the Directory Manager instance is initialized.
//
func (dMgr *DirMgr) IsInitialized() bool {

  return dMgr.isInitialized

}

// IsParentPathPopulated - Returns a boolean value
// indicating whether the parent path for this Directory
// Manager instance is populated.
//
func (dMgr *DirMgr) IsParentPathPopulated() bool {

  if len(dMgr.parentPath) == 0 {
    dMgr.isPathPopulated = false
  } else {
    dMgr.isPathPopulated = true
  }

  return dMgr.isPathPopulated
}

// IsPathPopulated - Returns a boolean value indicating
// whether the current Directory Manager path string is
// populated.
//
func (dMgr *DirMgr) IsPathPopulated() bool {

  if len(dMgr.path) == 0 {
    dMgr.isPathPopulated = false
  } else {
    dMgr.isPathPopulated = true
  }

  return dMgr.isPathPopulated
}

// IsRelativePathPopulated - Returns a boolean value indicating
// whether the Relative Path for this Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsRelativePathPopulated() bool {

  if len(dMgr.relativePath) == 0 {
    dMgr.isRelativePathPopulated = false
  } else {
    dMgr.isRelativePathPopulated = true
  }

  return dMgr.isRelativePathPopulated
}

// IsVolumeNamePopulated - Returns a boolean value indicating
// whether the Volume Name for the current Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsVolumeNamePopulated() bool {

  if len(dMgr.volumeName) == 0 {
    dMgr.isVolumePopulated = false
  } else {
    dMgr.isVolumePopulated = true
  }

  return dMgr.isVolumePopulated
}

// MakeDir - If the directory path identified by the current DirMgr
// object does not exist, this method will create that directory path.
// The path will be created using permission specifications passed through
// input parameter 'fPermCfg'.
//
func (dMgr *DirMgr) MakeDirWithPermission(fPermCfg FilePermissionConfig) error {

  ePrefix := "DirMgr.MakeDir() "
  var err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return err
  }

  err = fPermCfg.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  modePerm, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  dMgr.DoesDirMgrPathExist()
  dMgr.DoesDirMgrAbsolutePathExist()

  if dMgr.doesAbsolutePathExist {
    // No need to create directory, it already
    // exists.
    return nil
  }

  err = os.MkdirAll(dMgr.absolutePath, modePerm)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from os.MkdirAll(dMgr.absolutePath, "+
      "modePerm) dMgr.absolutePath='%v' modePerm='%v'  Error='%v'",
      dMgr.absolutePath, modePerm.String(), err.Error())
  }

  dMgr.DoesDirMgrPathExist()
  dMgr.DoesDirMgrAbsolutePathExist()

  // No errors - directory created.
  return nil

}

// MakeDir - If the directory path identified by the current DirMgr
// object does not exist, this method will create that directory path.
// The permission specification used to create the directory is
// 'drwxrwxrwx' which is equivalent to octal value, '020000000777'
//
func (dMgr *DirMgr) MakeDir() error {

  ePrefix := "DirMgr.MakeDir() "
  var err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return nil
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = dMgr.MakeDirWithPermission(fPermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  // No errors - directory created.
  return nil
}

// NewFromPathFileNameExtStr - Returns a new DirMgr object and populates the
// the data fields.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//
//	pathStr string - A path string designating a path or directory.
//	                 To reduce errors, the 'pathStr' should be terminated
//	                 with an appropriate path separator ('/' or '\')
//	                 Example 'pathStr': "C:\dirA\dirB\dirC\"
//
//	Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//
//                      isInitialized:  true
//                       originalPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//                               path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//                    IsPathPopulated:  true
//                      doesPathExist:  true
//                         parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//              isParentPathPopulated:  true
//                       relativePath:  testoverwrite
//            isRelativePathPopulated:  true
//                       absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//            isAbsolutePathPopulated:  true
//    isAbsolutePathDifferentFromPath:  false
//              doesAbsolutePathExist:  true
//                      directoryName:  testoverwrite
//                         volumeName:  D:
//                  isVolumePopulated:  true
//                  actualDirFileInfo:
//                         ========== File Info Data ==========
//                           File Info IsDir():  true
//                            File Info Name():  testoverwrite
//                            File Info Size():  0
//                         File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//                            File Info Mode():  drwxrwxrwx
//                             File Info Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//                                    Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr DirMgr) New(pathStr string) (DirMgr, error) {

  ePrefix := "DirMgr.NewFromPathFileNameExtStr() "

  newDirMgr := DirMgr{}

  _, err := newDirMgr.SetDirMgr(pathStr)

  if err != nil {
    return DirMgr{}, fmt.Errorf(ePrefix+
      "Error returned by newDirMgr.SetDirMgr(pathStr) pathStr='%v' Error='%v'",
      pathStr, err.Error())
  }

  return newDirMgr, nil
}

// NewFromFileInfo - Returns a new DirMgr object based on two input parameters:
// 		- A directory path string
//		- An os.FileInfo object
func (dMgr DirMgr) NewFromFileInfo(pathStr string, info os.FileInfo) (DirMgr, error) {

  ePrefix := "DirMgr) NewFromFileInfo() "

  newDirMgr := DirMgr{}

  err := newDirMgr.SetDirMgrWithFileInfo(pathStr, info)

  if err != nil {
    return DirMgr{},
      fmt.Errorf(ePrefix+"Error returned from '%v' ", err.Error())
  }

  return newDirMgr, nil
}

// SetDirMgr - Sets the DirMgr fields and path strings for the current DirMgr
// instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//
//	pathStr string - A path string designating a path or directory.
//	                 To reduce errors, the 'pathStr' should be terminated
//	                 with an appropriate path separator ('/' or '\')
//	                 Example 'pathStr': "C:\dirA\dirB\dirC\"
//
// Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//
//                      isInitialized:  true
//                       originalPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//                               path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//                    IsPathPopulated:  true
//                      doesPathExist:  true
//                         parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//              isParentPathPopulated:  true
//                       relativePath:  testoverwrite
//            isRelativePathPopulated:  true
//                       absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//            isAbsolutePathPopulated:  true
//    isAbsolutePathDifferentFromPath:  false
//              doesAbsolutePathExist:  true
//                      directoryName:  testoverwrite
//                         volumeName:  D:
//                  isVolumePopulated:  true
//                  actualDirFileInfo:
//                         ========== File Info Data ==========
//                           File Info IsDir():  true
//                            File Info Name():  testoverwrite
//                            File Info Size():  0
//                         File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//                            File Info Mode():  drwxrwxrwx
//                             File Info Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//                                    Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr *DirMgr) SetDirMgr(pathStr string) (isEmpty bool, err error) {

  ePrefix := "DirMgr.SetDirMgr() "

  dMgr.Empty()

  fh := FileHelper{}

  err = nil
  isEmpty = true
  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    isEmpty = true
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' is an empty string!")
    return isEmpty, err
  }

  if errCode == -2 {
    isEmpty = true
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!")
    return isEmpty, err
  }

  adjustedTrimmedPathStr := fh.AdjustPathSlash(pathStr)

  finalPathStr, isEmptyPath, err2 := fh.GetPathFromPathFileName(adjustedTrimmedPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. fh.GetPathFromPathFileName(pathStr) "+
      "pathStr='%v'  Error='%v'", pathStr, err2.Error())
    isEmpty = isEmptyPath
    return isEmpty, err
  }

  if isEmptyPath {
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. 'pathStr' generated an Empty path! pathStr='%v' ",
      pathStr)
    return isEmpty, err
  }

  if len(finalPathStr) == 0 {
    err = fmt.Errorf(ePrefix+
      "Error: path returned from fh.GetPathFromPathFileName(pathStr) is EMPTY! "+
      "pathStr='%v'", pathStr)
    isEmpty = true
    return isEmpty, err
  }

  dMgr.originalPath = adjustedTrimmedPathStr

  dMgr.path = finalPathStr

  dMgr.isPathPopulated = true
  dMgr.DoesDirMgrPathExist()

  if strings.ToLower(dMgr.path) == strings.ToLower(fp.VolumeName(dMgr.path)) {

    dMgr.absolutePath = fh.AdjustPathSlash(dMgr.path)

  } else {

    dMgr.absolutePath, err2 = fh.MakeAbsolutePath(dMgr.path)

    if err2 != nil {
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "- fh.MakeAbsolutePath(dMgr.path) returned error. dMgr.path='%v' Error='%v'",
        dMgr.path, err2.Error())
      isEmpty = true
      return isEmpty, err
    }

  }

  info, err2 := os.Stat(dMgr.absolutePath)

  if err2 == nil {
    if !info.IsDir() {
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "- The Directory Manager absolute path exists and IS NOT A DIRECTORY!.\n" +
        "dMgr.absolutePath='%v' Error='%v'",
        dMgr.absolutePath, err2.Error())
      isEmpty = true
      return
    }
  } else {
    // err2 != nil

    if !os.IsNotExist(err2) {
      // The error returned by os.Stat(dMgr.absolutePath) is NOT
      // a standard PATH DOES NOT EXIST error.
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "Non-Path Error triggered on Directory Manager absoltue path.\n" +
        "os.Stat(dMgr.absolutePath) error.\n" +
        "dMgr.absolutePath='%v' Error='%v'",
        dMgr.absolutePath, err2.Error())
      isEmpty = true
      return isEmpty, err
    }
  }

  dMgr.isAbsolutePathPopulated = true
  dMgr.DoesDirMgrAbsolutePathExist()

  strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
  lStr := len(strAry)
  idxStr := strAry[lStr-1]

  idx := strings.Index(dMgr.absolutePath, idxStr)
  dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

  dMgr.isParentPathPopulated = true

  if dMgr.isAbsolutePathPopulated && dMgr.isParentPathPopulated {

    dMgr.relativePath, err2 = fp.Rel(dMgr.parentPath, dMgr.absolutePath)

    if err2 != nil {
      dMgr.relativePath = ""
      dMgr.isParentPathPopulated = false
    } else {
      dMgr.isParentPathPopulated = true
      dMgr.isRelativePathPopulated = true

    }

  }

  if idxStr != "" {
    dMgr.directoryName = idxStr
  } else {
    dMgr.directoryName = dMgr.absolutePath
  }

  if dMgr.path != dMgr.absolutePath {
    dMgr.isAbsolutePathDifferentFromPath = true
  }

  var vn string
  if dMgr.isAbsolutePathPopulated {
    vn = fp.VolumeName(dMgr.absolutePath)
  } else if dMgr.isPathPopulated {
    vn = fp.VolumeName(dMgr.path)
  }

  if vn != "" {
    dMgr.isVolumePopulated = true
    dMgr.volumeName = vn
  }

  if dMgr.isAbsolutePathPopulated && dMgr.isPathPopulated {
    dMgr.isInitialized = true
    isEmpty = false
  } else {
    isEmpty = true
  }

  err = nil

  return
}

// SetDirMgrWithFileInfo - Sets the DirMgr fields and path strings for the current
// DirMgr object based on an input 'pathStr' parameter and an os.FileInfo input
// parameter ('info').
//
func (dMgr *DirMgr) SetDirMgrWithFileInfo(pathStr string, info os.FileInfo) error {
  ePrefix := "DirMgr.SetDirMgrWithFileInfo() "

  _, err := dMgr.SetDirMgr(pathStr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from dMgr.SetDirMgr(pathStr). "+
      "pathStr='%v'  Error='%v'", pathStr, err.Error())
  }

  dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathStr, info)
  dMgr.directoryName = info.Name()

  return nil
}


// SubstituteBaseDir - Substitute 'baseDir' segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
//
func (dMgr *DirMgr) SubstituteBaseDir(
  baseDir DirMgr,
  substituteBaseDir DirMgr) (newDMgr DirMgr, err error) {

  ePrefix := "DirMgr.SubstituteBaseDir() "

  newDMgr = DirMgr{}
  err = nil

  err2 := baseDir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error: Input parameter 'baseDir' is INVALID!")
    return newDMgr, err
  }

  err2 = substituteBaseDir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error: Input parameter 'substituteBaseDir' is INVALID!")
    return newDMgr, err
  }

  err2 = dMgr.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error: The current DirMgr object is INVALID!")
    return newDMgr, err
  }

  thisDirAbsPath := strings.ToLower(dMgr.absolutePath)

  oldBaseAbsPath := strings.ToLower(baseDir.absolutePath)

  newBaseAbsPath := strings.ToLower(substituteBaseDir.absolutePath)

  idx := strings.Index(thisDirAbsPath, oldBaseAbsPath)

  if idx < 0 {
    err = fmt.Errorf(ePrefix + "The base directory was NOT found in the current DirMgr path!\n" +
      "DirMgr Path='%v'\nbaseDir Path='%v'\n",
      thisDirAbsPath, oldBaseAbsPath)

    return newDMgr, err
  }

  if idx != 0 {
    err = fmt.Errorf(ePrefix + "The base directory was NOT found at the beginning of the current DirMgr path!\n" +
      "DirMgr Path='%v'\nbaseDir Path='%v'\n",
      thisDirAbsPath, oldBaseAbsPath)

    return newDMgr, err
  }

  oldBaseLen := len(oldBaseAbsPath)

  newAbsPath := newBaseAbsPath + thisDirAbsPath[oldBaseLen:]

  isEmpty := false

  isEmpty, err = newDMgr.SetDirMgr(newAbsPath)

  if err != nil {
    newDMgr.Empty()
    return newDMgr, err
  }

  if isEmpty {
    newDMgr.Empty()
    err = fmt.Errorf(ePrefix + "ERROR: New generated Directory Path Is Invalid!\n" +
      "newAbsPath='%v'\n", newAbsPath)
    return newDMgr, err
  }

  err = nil
  return newDMgr, err
}

// SubstituteBaseDir - Substitute baseDir segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
/*
func (dMgr *DirMgr) SubstituteBaseDir(
  baseDir DirMgr,
  substituteBaseDir DirMgr) (newDMgr DirMgr, err error) {

  ePrefix := "DirMgr.SubstituteBaseDir() "

  newDMgr = DirMgr{}
  err = nil

  err2 := baseDir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error: baseDir DirMgr object is INVALID!")
    return
  }

  err2 = substituteBaseDir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error: substituteBaseDir DirMgr object is INVALID!")
    return
  }

  var trimmedRemainingPath string
  var newPath string

  if strings.HasPrefix(dMgr.path, baseDir.path) {

    trimmedRemainingPath = strings.TrimPrefix(dMgr.path, baseDir.path)

    lPath := len(trimmedRemainingPath)

    if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
      trimmedRemainingPath = trimmedRemainingPath[1:]
    }

    newPath = substituteBaseDir.GetPathWithSeparator() + trimmedRemainingPath

  } else if strings.HasPrefix(dMgr.absolutePath, baseDir.absolutePath) {

    trimmedRemainingPath = strings.TrimPrefix(dMgr.absolutePath, baseDir.absolutePath)

    lPath := len(trimmedRemainingPath)

    if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
      trimmedRemainingPath = trimmedRemainingPath[1:]
    }

    newPath = substituteBaseDir.GetAbsolutePathWithSeparator() + trimmedRemainingPath

  } else {
    err = fmt.Errorf(ePrefix+
      "Error: Could not locate baseDir.path or "+
      "baseDir.absolutePath in this dMgr. dMgr.path='%v' dMgr.absolutePath='%v'",
      dMgr.path, dMgr.absolutePath)
    return
  }

  newDMgr, err2 = DirMgr{}.New(newPath)

  if err2 != nil {
    newDMgr = DirMgr{}
    err = fmt.Errorf(ePrefix+
      "Error returned from DirMgr{}.NewFromPathFileNameExtStr(newPath). "+
      "newPath='%v'  Error='%v'", newPath, err2.Error())
    return
  }

  err = nil
  return
}
*/

// executeFileOpsOnFoundFiles - This function is designed to work in conjunction
// with a walk directory function like FindWalkDirFiles. It will process
// files extracted from a 'Directory Walk' operation initiated by the
// 'filepath.Walk' method.
//
// Thereafter, file operations will be performed on files in the directory
// tree as specified by the 'dirOp' parameter.
//
func (dMgr *DirMgr) executeFileOpsOnFoundFiles(dirOp *DirTreeOp) func(string, os.FileInfo, error) error {
  return func(pathFile string, info os.FileInfo, erIn error) error {

    ePrefix := "DirMgr.executeFileOpsOnFoundFiles() "
    var err2 error

    if erIn != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned from directory walk function. "+
        "pathFile= '%v' Error='%v'", pathFile, erIn.Error())
      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    if info.IsDir() {
      return nil
    }

    fh := FileHelper{}

    // This is not a directory. It is a file.
    // Determine if it matches the find file criteria.
    isFoundFile, err := fh.FilterFileName(info, dirOp.FileSelectCriteria)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "Error returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria) "+
        "pathFile='%v' info.Name()='%v' Error='%v' ",
        pathFile, info.Name(), err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    if !isFoundFile {
      return nil
    }

    srcFileNameExt := info.Name()

    destDir, err := fh.SwapBasePath(
      dirOp.SourceBaseDir.GetAbsolutePath(),
      dirOp.TargetBaseDir.GetAbsolutePath(),
      pathFile)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned by fh.SwapBasePath(dirOp.SourceBaseDir, "+
        "dirOp.TargetBaseDir, pathFile). dirOp.SourceBaseDir='%v' "+
        "dirOp.TargetBaseDir='%v' pathFile='%v' Error='%v' ",
        dirOp.SourceBaseDir.GetAbsolutePath(),
        dirOp.TargetBaseDir.GetAbsolutePath(),
        pathFile,
        err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil

    }

    fileOp, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
      pathFile, srcFileNameExt, destDir, srcFileNameExt)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs() "+
        "pathFile='%v' srcFileNameExt='%v' destDir='%v' Error='%v' ",
        pathFile, srcFileNameExt, destDir, err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    maxOps := len(dirOp.FileOps)

    for i := 0; i < maxOps; i++ {

      err = fileOp.ExecuteFileOperation(dirOp.FileOps[i])

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by fileOp.ExecuteFileOperation(dirOp.FileOps[i]). "+
          "i='%v' FileOps='%v' Error='%v' ",
          i, dirOp.FileOps[i].String(), err.Error())

        dirOp.ErrReturns = append(dirOp.ErrReturns, err2)

      }
    }

    return nil
  }

}
