package pathfileops

import (
  "errors"
  "fmt"
  "io"
  "os"
  fp "path/filepath"
  "strings"
  "sync"
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
  absolutePath                    string
  isAbsolutePathPopulated         bool
  doesAbsolutePathExist           bool
  isAbsolutePathDifferentFromPath bool
  directoryName                   string // Name of directory with out parent path.
  volumeName                      string
  isVolumePopulated               bool
  actualDirFileInfo               FileInfoPlus
  dataMutex                       sync.Mutex // Used internally to ensure thread safe operations
}

// CopyDirectory - Copies files from the directory identified by
// by DirMgr to a target directory. The files to be copied are selected
// according to file selection criteria specified by input parameter,
// 'fileSelectCriteria'.
//
// The selected files are copied by a Copy IO operation. For information
// on the Copy IO procedure see FileHelper{}.CopyFileByIo() method and
// reference:
//   https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// NOTE: This method ONLY copies files from the current directory identified
// by 'DirMgr'. It does NOT copy files from subdirectories.
//
// If the target directory does not exist and files are located matching the
// file selection criteria, this method will attempt to create the target
// directory. However, if no files meet the file selection criteria as
// defined by input parameter,'fileSelectCriteria', this method will NOT
// attempt to create the target directory.
//
// This method is optimized to support the copy of large numbers of files.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  targetDMgr   DirMgr - An instance of 'DirMgr' initialized with the directory
//                        path of the target directory to which selected files
//                        will be copied. If the target directory does not exist,
//                        this method will attempt to create it.
//
//
//  fileSelectCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be copied
//    to the directory identified by input parameter, 'targetDir'.
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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
// the files processed in the directory tree will be selected and copied
// to the target directory.
//
//      Example:
//        fsc := FileSelectCriterionMode{}
//
//        In this example, 'fsc' is NOT initialized. Therefore,
//        all of the selection criterion are 'Inactive'. Consequently,
//        all of the files encountered in the target directory during
//        the search operation will be selected and copied to the
//        target directory.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) CopyDirectory(
  targetDMgr DirMgr,
  fileSelectCriteria FileSelectionCriteria) []error {

  ePrefix := "DirMgr.CopyDirectory() "
  dMgrHlpr := dirMgrHelper{}
  var errs []error

  dMgr.dataMutex.Lock()

  errs = dMgrHlpr.copyDirectory(
    dMgr,
    &targetDMgr,
    fileSelectCriteria,
    ePrefix,
    "dMgr",
    "targetDMgr")

  dMgr.dataMutex.Unlock()

  return errs
}

// CopyDirectoryTree - Copies all selected files in the directory tree to
// a specified target directory tree. If the target directory tree does not
// exist, this method will attempt to create it. See the details of target
// directory tree creation under input parameter 'copyEmptyDirectories'.
//
// If input parameter 'copyEmptyDirectories' is set to 'true', the entire
// directory tree will be created and may contain empty directories. If
// set to false, target directory tree elements will only be created if
// files meet the selection criteria and are subsequently copied to those
// target directory tree paths.
//
// Files eligible for copy to the target directory tree are selected on the
// basis of file selection criteria specified by input parameter,
// 'fileSelectCriteria'.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  targetDMgr        DirMgr - An instance of 'DirMgr' initialized with the directory
//                             path of the target directory to which selected files
//                             will be copied. If the target directory does not exist,
//                             this method will attempt to create it.
//
// copyEmptyDirectories bool - If a target directory tree path does not previously exist,
//                             the default behavior is to create that directory ONLY if
//                             files matching the file selection criteria are identified
//                             for that directory. If no files match the file selection
//                             criteria, the default is to NOT create the target directory
//                             path.
//
//                             If the parameter 'copyEmptyDirectories' is set to 'true' all
//                             target directory tree paths will be created regardless of
//                             whether files are copied to those directories.
//
//  fileSelectCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be copied
//    to the directory identified by input parameter, 'targetDir'.
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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
//        the search operation will be selected and copy to target
//        directory.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) CopyDirectoryTree(
  targetDMgr DirMgr,
  copyEmptyDirectories bool,
  fileSelectCriteria FileSelectionCriteria) (errs []error) {

  ePrefix := "DirMgr.CopyDirectoryTree() "
  errs = nil

  dMgrHlpr := dirMgrHelper{}

  errs = dMgrHlpr.copyDirectoryTree(
    dMgr,
    &targetDMgr,
    true, // createTargetDir
    copyEmptyDirectories,
    false, // skipTopLevelDirectory
    fileSelectCriteria,
    ePrefix,
    "dMgr",
    "targetDMgr")

  return errs
}

// CopyIn - Receives a pointer to a DirMgr object as an
// input parameter and copies the values from the incoming
// object to the current DirMgr object. When the copy operation
// is completed, the current DirMgr object is a duplicate of the
// incoming DirMgr object.
func (dMgr *DirMgr) CopyIn(dmgrIn *DirMgr) {
  dmgrIn.dataMutex.Lock()
  dMgr.dataMutex.Lock()
  dMgr.isInitialized = dmgrIn.isInitialized
  dMgr.originalPath = dmgrIn.originalPath
  dMgr.path = dmgrIn.path
  dMgr.isPathPopulated = dmgrIn.isPathPopulated
  dMgr.doesPathExist = dmgrIn.doesPathExist
  dMgr.parentPath = dmgrIn.parentPath
  dMgr.isParentPathPopulated = dmgrIn.isParentPathPopulated
  dMgr.absolutePath = dmgrIn.absolutePath
  dMgr.isAbsolutePathPopulated = dmgrIn.isAbsolutePathPopulated
  dMgr.doesAbsolutePathExist = dmgrIn.doesAbsolutePathExist
  dMgr.isAbsolutePathDifferentFromPath = dmgrIn.isAbsolutePathDifferentFromPath
  dMgr.directoryName = dmgrIn.directoryName
  dMgr.volumeName = dmgrIn.volumeName
  dMgr.isVolumePopulated = dmgrIn.isVolumePopulated
  dMgr.actualDirFileInfo = dmgrIn.actualDirFileInfo.CopyOut()
  dmgrIn.dataMutex.Unlock()
  dMgr.dataMutex.Unlock()
}

// CopyOut - Makes a duplicate copy of the current DirMgr values and
// returns them in a new DirMgr object.
func (dMgr *DirMgr) CopyOut() DirMgr {

  dOut := DirMgr{}
  dMgr.dataMutex.Lock()

  dOut.isInitialized = dMgr.isInitialized
  dOut.originalPath = dMgr.originalPath
  dOut.path = dMgr.path
  dOut.isPathPopulated = dMgr.isPathPopulated
  dOut.doesPathExist = dMgr.doesPathExist
  dOut.parentPath = dMgr.parentPath
  dOut.isParentPathPopulated = dMgr.isParentPathPopulated
  dOut.absolutePath = dMgr.absolutePath
  dOut.isAbsolutePathPopulated = dMgr.isAbsolutePathPopulated
  dOut.doesAbsolutePathExist = dMgr.doesAbsolutePathExist
  dOut.isAbsolutePathDifferentFromPath = dMgr.isAbsolutePathDifferentFromPath
  dOut.directoryName = dMgr.directoryName
  dOut.volumeName = dMgr.volumeName
  dOut.isVolumePopulated = dMgr.isVolumePopulated
  dOut.actualDirFileInfo = dMgr.actualDirFileInfo.CopyOut()

  dMgr.dataMutex.Unlock()
  return dOut
}

// CopySubDirectoryTree - Treating the directory identified by the current 'DirMgr'
// instance as the parent directory, this method copies all sub-directories and the
// files contained in those sub-directories to the target directory specified by
// input parameter, 'targetDir'. Essentially, the entire directory tree with the
// sole exception of the top level parent directory is copied to target directory.
//
// The files copied must match the selection criteria specified by input parameter
// 'fileSelectCriteria'.
//
// If the target directory does not exist, and valid matching files are identified for
// that directory, this method will attempt to create the target directory. Conversely,
// if no files matching the file selection criteria are found, that corresponding target
// directory will NOT be created.
//
func (dMgr *DirMgr) CopySubDirectoryTree(
  targetDMgr DirMgr,
  copyEmptyDirectories bool,
  fileSelectCriteria FileSelectionCriteria) (errs []error) {

  ePrefix := "DirMgr.CopySubDirectoryTree() "
  var err, err2 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  err = targetDMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    err2 = fmt.Errorf("Input parameter 'targetDMgr' is INVALID!\n"+
      "Error='%v'\n", err.Error())
    errs = append(errs, err2)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {

    err2 = fmt.Errorf(ePrefix+"The current DirMgr path DOES NOT EXIST!\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  // dMgr.absolutePath DOES EXIST!

  return dMgr.copyDirectoryTree(
    targetDMgr,
    copyEmptyDirectories,
    true,
    ePrefix,
    fileSelectCriteria)
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

  dMgrHlpr := dirMgrHelper{}

  var err error

  dMgr.dataMutex.Lock()

  err = dMgrHlpr.deleteDirectoryAll(
    dMgr,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

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
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) DeleteAllFilesInDir() (errs []error) {

  ePrefix := "DirMgr.DeleteAllFilesInDir() "

  errs = make([]error, 0, 300)
  var err, err2, err3 error
  osPathSepStr := string(os.PathSeparator)

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {
    err2 =
      fmt.Errorf(ePrefix+
        "ERROR: DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path='%v'\n", dMgr.absolutePath)
    errs = append(errs, err2)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath).\n"+
      "dMgr.absolutePath='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  err3 = nil
  var nameFileInfos []os.FileInfo

  for err3 != io.EOF {

    nameFileInfos, err3 = dir.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      _ = dir.Close()
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(-1).\n"+
        "dMgr.absolutePath='%v'\nError='%v'\n",
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
      return errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        continue

      } else {

        dMgr.dataMutex.Lock()

        err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFInfo.Name())

        dMgr.dataMutex.Unlock()

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "Error returned by os.Remove(fileName).\n"+
            "dMgr.absolutePath='%v'\nfileName='%v'\nError='%v'\n",
            dMgr.absolutePath,
            dMgr.absolutePath+osPathSepStr+nameFInfo.Name(),
            err.Error())

          errs = append(errs, err2)
        }
      }
    }
  }

  if dir != nil {

    err = dir.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Close(). An attempt to close the os.File pointer to the current\n"+
        "DirMgr path has FAILED!\n"+
        "dMgr.absolutePath='%v'\nError='%v'\n",
        dMgr.absolutePath, err.Error())
      errs = append(errs, err2)
    }
  }

  return errs
}

// DeleteAllSubDirectories - The directory identified by the current
// DirMgr instance is treated as the parent directory. This method
// will then proceed to delete all directories and files which are
// are subsidiary to this parent directory. Essentially, all sub-
// directories which are subordinate to the the DirMgr directory will
// be deleted along with their constituent files.
//
//
//  Example:
//   Parent Directory:
//    DirMgr = d:\parentdirectory
//    files    d:\parentdirectory\file1.txt
//             d:\parentdirectory\file2.txt
//
//   Sub-Directories:
//             d:\parentdirectory\dir01
//             d:\parentdirectory\dir02
//             d:\parentdirectory\dir03
//
//  After Executing DirMgr.DeleteAllSubDirectories() all sub-directories and
//  any files they contain will be deleted. The only directory which remains
//  is the parent directory and any files contained within the parent directory.
//
//    DirMgr = d:\parentdirectory
//    files    d:\parentdirectory\file1.txt
//             d:\parentdirectory\file2.txt
//
func (dMgr *DirMgr) DeleteAllSubDirectories() (errs []error) {

  errs = make([]error, 0, 300)

  ePrefix := "DirMgr.CopyDirectory() "
  var err, err2, err3 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {

    err2 = fmt.Errorf(ePrefix+"The current DirMgr path DOES NOT EXIST!\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)

    return errs
  }

  var nameFileInfos []os.FileInfo
  err3 = nil
  osPathSeparatorStr := string(os.PathSeparator)

  for err3 != io.EOF {

    nameFileInfos, err3 = dir.Readdir(1000)

    if err3 != nil && err3 != io.EOF {
      _ = dir.Close()
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(1000). "+
        "dMgr.absolutePath='%v'\nError='%v'\n\n",
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        dMgr.dataMutex.Lock()

        err = os.RemoveAll(dMgr.absolutePath + osPathSeparatorStr + nameFInfo.Name())

        dMgr.dataMutex.Unlock()

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "Error returned by os.RemoveAll(subDir)\n"+
            "subDir='%v'\nError='%v'\n",
            dMgr.absolutePath+osPathSeparatorStr+nameFInfo.Name(), err.Error())

          errs = append(errs, err2)

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

    errs = append(errs, err2)
  }

  return errs
}

// DeleteDirectoryTreeFiles - Deletes files in the directory tree. The parent
// directory for this tree is the directory specified by the current 'DirMgr'
// instance. The file deletion operation is conducted in three steps:
//    1. The criteria for selecting files to be deleted is created using
//       input parameter 'deleteFileSelectionCriteria'.
//    2. A file search is conducted which includes the DirMgr parent directory
//       and all sub-directories in the tree.
//    3. Files processed during the directory tree search are compared to the
//       file selection criteria specified by 'deleteFileSelectionCriteria'.
//       Those files which match the selection criteria are then deleted.
//
// This method is similar to method 'DirMgr.DeleteWalkDirFiles()'. However this
// method returns less data and is designed to work with very large numbers of
// files and directories.
//
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  deleteFileSelectionCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be deleted.
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file modification
//                                    date times in order to determine whether the file is
//                                    older than the 'FilesOlderThan' file selection criterion.
//                                    If the file modification date time is older than the
//                                    'FilesOlderThan' date time, that file is considered a
//                                    'match'	for this file selection criterion.
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
// the files processed in the directory tree will be selected and deleted.
//
//      Example:
//        fsc := FileSelectCriterionMode{}
//
//        In this example, 'fsc' is NOT initialized. Therefore,
//        all of the selection criterion are 'Inactive'. Consequently,
//        all of the files encountered in the target directory tree
//        during the search operation will be selected and deleted.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) DeleteDirectoryTreeFiles(
  deleteFileSelectionCriteria FileSelectionCriteria) (numOfSubDirectories,
  numOfRemainingFiles,
  numOfDeletedFiles int,
  errs []error) {
  ePrefix := "DirMgr.DeleteDirectoryTreeFiles() "

  numOfSubDirectories = 0
  numOfRemainingFiles = 0
  numOfDeletedFiles = 0
  errs = make([]error, 0, 300)

  nonPathError := dMgr.IsDirMgrValid(ePrefix)

  if nonPathError != nil {
    errs = append(errs, nonPathError)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }
  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {

    errs = append(errs, nonPathError)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  if !dirPathDoesExist {
    nonPathError = fmt.Errorf(ePrefix+
      "Error: Source DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'", dMgr.absolutePath)
    errs = append(errs, nonPathError)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  if !fInfoPlus.IsDir() {
    nonPathError = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, nonPathError)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  numOfSubDirectories,
    numOfRemainingFiles,
    numOfDeletedFiles,
    errs = dMgr.deleteDirectoryTreeFiles(
    true,
    ePrefix,
    deleteFileSelectionCriteria)

  return numOfSubDirectories,
    numOfRemainingFiles,
    numOfDeletedFiles,
    errs
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

  var err2, err, err3 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {
    err2 =
      fmt.Errorf(ePrefix+
        "ERROR: Directory DOES NOT EXIST!\n"+
        "DirMgr Directory='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode == -1 {
    err2 = errors.New(ePrefix +
      "Error: Input parameter 'fileSearchPattern' is an empty string!")

    errs = append(errs, err2)
    return errs
  }

  if errCode == -2 {
    err2 = errors.New(ePrefix +
      "Error: Input parameter 'fileSearchPattern' consists of blank spaces!")
    errs = append(errs, err2)
    return errs
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())
    errs = append(errs, err2)
    return errs
  }

  err3 = nil
  var nameFileInfos []os.FileInfo
  osPathSepStr := string(os.PathSeparator)
  var isMatch bool

  for err3 != io.EOF {

    nameFileInfos, err3 = dir.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      _ = dir.Close()
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(-1). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
      return errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        continue

      } else {

        isMatch, err = fp.Match(fileSearchPattern, nameFInfo.Name())

        if err != nil {

          err2 = fmt.Errorf(ePrefix+
            "Error returned by fp.Match(fileSearchPattern, fileName).\n"+
            "directorySearched='%v'\nfileSearchPattern='%v' fileName='%v'\nError='%v'\n",
            dMgr.absolutePath, fileSearchPattern, nameFInfo.Name(), err.Error())

          errs = append(errs, err2)
          continue
        }

        if !isMatch {

          continue

        } else {

          dMgr.dataMutex.Lock()

          err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFInfo.Name())

          dMgr.dataMutex.Unlock()

          if err != nil {
            err2 = fmt.Errorf(ePrefix+
              "Error returned by os.Remove(pathFileName).\n"+
              "pathFileName='%v'\nError='%v'\n\n",
              dMgr.absolutePath+osPathSepStr+nameFInfo.Name(),
              err.Error())

            errs = append(errs, err2)
            continue
          }
        }
      }
    }
  }

  err = dir.Close()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned by dir.Close().\n"+
      "dir='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())
    errs = append(errs, err2)
  }

  return errs
}

// DeleteFilesBySelectionCriteria - Deletes selected files from the directory
// identified by the current 'DirMgr' instance.
//
// Files in sub-directories are not deleted or altered in any way. ONLY files
// in the directory identified by the current 'DirMgr' instance are deleted.
//
// The file deletion operation consists of three steps:
//
//    1. The criteria for selecting files to be deleted is created using
//       input parameter 'deleteFileSelectionCriteria'.
//
//    2. A file search is conducted which is limited ONLY to the DirMg
//       directory. Files in sub-directory tree ARE NEVER DELETED.
//
//    3. Files processed during the directory search are compared to the
//       file selection criteria specified by 'deleteFileSelectionCriteria'.
//       Those files which match the selection criteria are then deleted.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  deleteFileSelectionCriteria FileSelectionCriteria
//    This input parameter should be configured with the desired file selection
//    criteria. Files in the DirMgr directory, matching this criteria, will be
//    deleted.
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//    FilesOlderThan      time.Time - This date time type is compared to file modification
//                                    date times in order to determine whether the file is
//                                    older than the 'FilesOlderThan' file selection criterion.
//                                    If the file modification date time is older than the
//                                    'FilesOlderThan' date time, that file is considered a
//                                    'match'	for this file selection criterion.
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
// the files processed in the directory tree will be selected and deleted.
//
//      Example:
//        fsc := FileSelectCriterionMode{}
//
//        In this example, 'fsc' is NOT initialized. Therefore,
//        all of the selection criterion are 'Inactive'. Consequently,
//        all of the files encountered in the target directory tree
//        during the search operation will be selected and deleted.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) DeleteFilesBySelectionCriteria(
  deleteFileSelectionCriteria FileSelectionCriteria) (numOfRemainingFiles,
  numOfDeletedFiles int,
  errs []error) {
  ePrefix := "DirMgr.DeleteDirectoryTreeFiles() "

  numOfRemainingFiles = 0
  numOfDeletedFiles = 0
  errs = make([]error, 0, 300)

  nonPathError := dMgr.IsDirMgrValid(ePrefix)

  if nonPathError != nil {
    errs = append(errs, nonPathError)

    return numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)

    return numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  if !dirPathDoesExist {

    nonPathError = fmt.Errorf(ePrefix+
      "Error: Source DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'\n", dMgr.absolutePath)
    errs = append(errs, nonPathError)

    return numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  if !fInfoPlus.IsDir() {
    nonPathError = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, nonPathError)

    return numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  _,
    numOfRemainingFiles,
    numOfDeletedFiles,
    errs = dMgr.deleteDirectoryTreeFiles(
    false,
    ePrefix,
    deleteFileSelectionCriteria)

  return numOfRemainingFiles,
    numOfDeletedFiles,
    errs
}

// DeleteWalkDirFiles - !!! BE CAREFUL !!! This method deletes files
// in a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the current DirMgr object. The method 'walks the
// directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// If a file matches the File Selection Criteria, it is DELETED. By the
// way, if ALL the file selection criterion are set to zero values or
// 'Inactive', then ALL FILES IN THE DIRECTORY ARE DELETED!!!
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return deleteFilesInfo, nonPathError
  }

  if !dirPathDoesExist {
    return deleteFilesInfo,
      fmt.Errorf(ePrefix+
        "ERROR: DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    nonPathError = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    return deleteFilesInfo, nonPathError
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

// DoesAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.absolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
//
func (dMgr *DirMgr) DoesAbsolutePathExist() bool {

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      "",
      "dMgr.absolutePath")

  if nonPathError != nil || !dirPathDoesExist {
    dMgr.doesAbsolutePathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
  } else {
    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = fInfoPlus.CopyOut()
  }

  return dMgr.doesAbsolutePathExist
}

// DoesDirectoryExist - Returns two boolean values indicating whether or not the
// Directory path exists and whether or not the Directory absolute path exists.
//
func (dMgr *DirMgr) DoesDirectoryExist() (doesPathExist, doesAbsolutePathExist bool) {

  doesPathExist = false
  doesAbsolutePathExist = false

  nonPathError := dMgr.IsDirMgrValid("")

  if nonPathError != nil {
    return doesPathExist, doesAbsolutePathExist
  }

  doesPathExist = dMgr.doesPathExist

  doesAbsolutePathExist = dMgr.doesAbsolutePathExist

  return doesPathExist, doesAbsolutePathExist
}

// DoesPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition it also
// updates the DirMgr field DirMgr.doesPathExist field.
//
func (dMgr *DirMgr) DoesPathExist() bool {
  _,
    dirPathDoesExist,
    _,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.path,
      PreProcPathCode.None(),
      "",
      "dMgr.path")

  if nonPathError != nil || !dirPathDoesExist {
    dMgr.doesPathExist = false
  } else {
    dMgr.doesPathExist = true
  }

  return dMgr.doesPathExist
}

// DoesThisDirectoryExist - Returns a boolean value of true if the directory identified
// by the current DirMgr instance does in fact exist.
//
// If, during the process of verifying the existence of the current directory, an error
// is encountered it will be a non-path error. Non-Path errors are most commonly associated
// with 'access-denied' situations. However, there may be other reasons for triggering Non-Path
// errors.
//
// If a Non-Path error is encountered, an appropriate error message is returned along with
// a boolean value of 'false'.
//
func (dMgr *DirMgr) DoesThisDirectoryExist() (directoryDoesExist bool, nonPathError error) {

  ePrefix := "DirMgr.DoesThisDirectoryExist() "

  directoryDoesExist = false
  nonPathError = nil

  fInfoPlus := FileInfoPlus{}
  _,
    directoryDoesExist,
    fInfoPlus,
    nonPathError =
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    directoryDoesExist = false
    return directoryDoesExist, nonPathError
  }

  if !directoryDoesExist {
    dMgr.doesAbsolutePathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
  } else {
    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = fInfoPlus.CopyOut()
  }

  return directoryDoesExist, nonPathError
}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

  dMgr.dataMutex.Lock()

  dMgr.isInitialized = false
  dMgr.originalPath = ""
  dMgr.path = ""
  dMgr.isPathPopulated = false
  dMgr.doesPathExist = false
  dMgr.parentPath = ""
  dMgr.isParentPathPopulated = false
  dMgr.absolutePath = ""
  dMgr.isAbsolutePathPopulated = false
  dMgr.doesAbsolutePathExist = false
  dMgr.isAbsolutePathDifferentFromPath = false
  dMgr.directoryName = ""
  dMgr.volumeName = ""
  dMgr.isVolumePopulated = false
  dMgr.actualDirFileInfo = FileInfoPlus{}

  dMgr.dataMutex.Unlock()
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {

    nonPathError = fmt.Errorf(ePrefix+"The current DirMgr path DOES NOT EXIST!\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)

    errs = append(errs, nonPathError)
    return errs
  }

  if !fInfoPlus.IsDir() {
    nonPathError = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, nonPathError)
    return errs
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

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

      err2 := fmt.Errorf(ePrefix+
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {

    err = fmt.Errorf(ePrefix+"The current DirMgr path DOES NOT EXIST!\n"+
      "dMgr.absolutePath='%v'\n", dMgr.absolutePath)

    errs = append(errs, err)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err)
    return errs
  }

  dirOp.FileOps = append(dirOp.FileOps, fileOps...)
  dirOp.TargetBaseDir = targetBaseDir.CopyOut()
  dirOp.SourceBaseDir = dMgr.CopyOut()
  dirOp.FileSelectCriteria = fileSelectCriteria

  err = fp.Walk(dMgr.GetAbsolutePath(), dMgr.executeFileOpsOnFoundFiles(&dirOp))

  if err != nil {
    err2 := fmt.Errorf(ePrefix+
      "Error returned by fp.Walk(). Error='%v' ", err.Error())
    errs = append(errs, dirOp.ErrReturns...)
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

  nonPathError := dMgr.IsDirMgrValid(ePrefix)

  if nonPathError != nil {
    return FileMgrCollection{}, nonPathError
  }

  errCode := 0

  errCode, _, fileSearchPattern = FileHelper{}.isStringEmptyOrBlank(fileSearchPattern)

  if errCode < 0 {
    return FileMgrCollection{},
      errors.New(ePrefix + "Input parameter 'fileSearchPattern' is INVALID!\n" +
        "'fileSearchPattern' is an EMPTY STRING!\n")
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return FileMgrCollection{}, nonPathError
  }

  if !dirPathDoesExist {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path ='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return FileMgrCollection{}, nonPathError
  }

  if !dirPathDoesExist {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path ='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return FileMgrCollection{},
      fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return findFilesInfo, nonPathError
  }

  if !dirPathDoesExist {
    return findFilesInfo,
      fmt.Errorf(ePrefix+
        "DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path ='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return findFilesInfo,
      fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
  }

  findFilesInfo.StartPath = dMgr.absolutePath

  findFilesInfo.FileSelectCriteria = fileSelectCriteria

  fh := FileHelper{}

  err = fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

  if err != nil {
    return findFilesInfo, fmt.Errorf(ePrefix+
      "Error returned by FileHelper.FindFilesWalkDirectory(&dWalkInfo).\n"+
      "dWalkInfo.StartPath='%v'\nError='%v'\n", findFilesInfo.StartPath, err.Error())
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

// GetDirectoryTree - Returns a DirMgrCollection containing all
// the sub-directories in the path of the parent directory identified
// by the current DirMgr instance.
//
// The returned DirMgrCollection will always contain the parent directory
// at the top of the array (index=0). Therefore, if no errors are encountered,
// the returned DirMgrCollection will always consist of at least one directory.
// If sub-directories are found, then the returned DirMgrCollection will
// contain more than one directory.
//
func (dMgr *DirMgr) GetDirectoryTree() (dirMgrs DirMgrCollection, errs []error) {

  ePrefix := "DirMgr.GetDirectoryTree() "

  dirMgrs = DirMgrCollection{}.New()

  errs = make([]error, 0, 100)

  var err, err2, err3 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return dirMgrs, errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return dirMgrs, errs

  }

  if !dirPathDoesExist {
    err2 =
      fmt.Errorf(ePrefix+
        "ERROR: DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path='%v'\n", dMgr.absolutePath)
    errs = append(errs, err2)
    return dirMgrs, errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return dirMgrs, errs
  }

  dirMgrs.AddDirMgr(dMgr.CopyOut())

  fh := FileHelper{}

  maxLen := dirMgrs.GetNumOfDirs()

  var dir *os.File
  var nameFileInfos []os.FileInfo

  for i := 0; i < maxLen; i++ {

    dir, err = os.Open(dirMgrs.dirMgrs[i].absolutePath)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error return by os.Open(dirMgrs.dirMgrs[%v].absolutePath). "+
        "dMgr.absolutePath='%v'\nError='%v'\n",
        i, dirMgrs.dirMgrs[i].absolutePath, err.Error())
      errs = append(errs, err2)

      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dir.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dir.Readdirnames(-1).\n"+
          "dMgr.absolutePath='%v'\nError='%v'\n",
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          newDirPathFileName :=
            fh.JoinPathsAdjustSeparators(dirMgrs.dirMgrs[i].absolutePath, nameFInfo.Name())

          err = dirMgrs.AddDirMgrByPathNameStr(newDirPathFileName)

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by dirMgrs.AddDirMgrByPathNameStr(newDirPathFileName). "+
                "dir='%v' Error='%v' ",
                newDirPathFileName, err.Error())

            errs = append(errs, err2)
            continue
          }

          maxLen = dirMgrs.GetNumOfDirs()
        }
      }
    }

    if dir != nil {

      err = dir.Close()

      if err != nil {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dir.Close().\n"+
          "dir='%v'\nError='%v'\n",
          dMgr.absolutePath, err.Error())
        errs = append(errs, err2)
      }
    }
  }

  return dirMgrs, errs
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

  var err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return FileInfoPlus{}, err
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return FileInfoPlus{}, nonPathError
  }

  if !dirPathDoesExist {
    return FileInfoPlus{},
      fmt.Errorf(ePrefix+"DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return FileInfoPlus{},
      fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
  }

  dMgr.actualDirFileInfo = fInfoPlus.CopyOut()

  return dMgr.actualDirFileInfo.CopyOut(), nil
}

// GetDirPermissionCodes - If the current directory exists on disk,
// this method will return the Directory Permission Codes encapsulated
// in a type 'FilePermissionConfig'.
//
// If the current Directory does NOT exist, this method will return an
// error.
//
func (dMgr *DirMgr) GetDirPermissionCodes() (FilePermissionConfig, error) {

  ePrefix := "GetDirPermissionCodes() "

  var nonPathError error

  nonPathError = dMgr.IsDirMgrValid(ePrefix)

  if nonPathError != nil {
    return FilePermissionConfig{}, nonPathError
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return FilePermissionConfig{}, nonPathError
  }

  if !dirPathDoesExist {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+"DirMgr Path DOES NOT EXIST!\n"+
        "DirMgr Path='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
  }

  dMgr.actualDirFileInfo = fInfoPlus.CopyOut()

  fPerm, nonPathError := FilePermissionConfig{}.NewByFileMode(dMgr.actualDirFileInfo.Mode())

  if nonPathError != nil {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "Error creating File Permission Configuration\nError='%v'\n",
        nonPathError.Error())
  }

  return fPerm, nil
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

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    hasParent = false
    return dirMgr, hasParent, err
  }

  if len(dMgr.parentPath) == 0 {

    return dMgr.CopyOut(), false, nil

  }

  var err2 error

  dirMgr, err2 = DirMgr{}.New(dMgr.parentPath)

  if err2 != nil {

    err = fmt.Errorf(ePrefix+"%v", err2.Error())
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

// GetVolumeName - Returns a string containing the volume name
// of the directory identified by the current Directory Manager
// instance.
//
func (dMgr *DirMgr) GetVolumeName() string {
  return dMgr.volumeName
}

// IsDirMgrValid - This method examines the current DirMgr object
// to determine whether it has been properly configured.
// If the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error is
// returned.
func (dMgr *DirMgr) IsDirMgrValid(errPrefixStr string) error {

  ePrefix := strings.TrimLeft(strings.TrimRight(errPrefixStr, " "), " ")

  if len(ePrefix) == 0 {
    ePrefix = "DirMgr.IsDirMgrValid() "
  } else {
    ePrefix = ePrefix + "- DirMgr.IsDirMgrValid() "
  }

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return err
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

  _,
    dirPathDoesExist,
    _,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return nonPathError
  }

  if dirPathDoesExist {
    // nothing to do. Exit!
    return nil
  }

  err = fPermCfg.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  modePerm, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  dMgr.dataMutex.Lock()

  err = os.MkdirAll(dMgr.absolutePath, modePerm)

  dMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from os.MkdirAll(dMgr.absolutePath, "+
      "modePerm)\ndMgr.absolutePath='%v'\nmodePerm='%v'\nError='%v'\n",
      dMgr.absolutePath, modePerm.String(), err.Error())
  }

  dMgr.DoesPathExist()
  dMgr.DoesAbsolutePathExist()

  if !dMgr.doesAbsolutePathExist {
    return fmt.Errorf(ePrefix+
      "Error: FAILED TO CREATE DIRECTORY!!\n"+
      "\nDirectory Path='%v'\n", dMgr.absolutePath)
  }

  // No errors - directory created.
  return nil
}

// MakeDir - If the directory path identified by the current DirMgr
// object does not exist, this method will create that directory path.
// The permission specification used to create the directory is
// 'drwxrwxrwx' which is equivalent to octal value, '020000000777'
//
// MakeDir creates a directory named path, along with any necessary
// parent directories.
//
// If the directory creation fails, an error is returned.
//
func (dMgr *DirMgr) MakeDir() error {

  ePrefix := "DirMgr.MakeDir() "
  var err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return err
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  err = dMgr.MakeDirWithPermission(fPermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  // No errors - directory created.
  return nil
}

// MoveDirectory - Moves files from the source directory identified
// by DirMgr to a target directory. The 'move' operation is accomplished
// in three steps. First, the files to be copied are selected according
// to file selection criteria specified by input parameter,'fileSelectCriteria'.
// Second, the selected files are copied to target directory identified
// by the input parameter, 'targetDir'. Finally, after verifying the copy,
// the files are deleted from the source directory (DirMgr).
//
// If, at the conclusion of the 'move' operation, there are no files or
// sub-directories remaining in the source directory (DirMgr), the source
// directory will be delete.
//
// The selected files are copied using Copy IO operation. For information
// on the Copy IO procedure see FileHelper{}.CopyFileByIo() method and
// reference:
//   https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// If the target directory does not previously exist, this method will attempt
// to create the target directory, provided, that files are selected for movement
// to that directory. If no files match the file selection criteria, the target
// directory will NOT be created.
//
// NOTE: This method ONLY moves files from the current directory identified
// by 'DirMgr'. It does NOT move files from subdirectories.
//
// This method is optimized to support the movement of large numbers of files.
//
// ------------------------------------------------------------------------------
//
// IMPORTANT!!!!
// This method will delete files in the current DirMgr path!  If all files have
// been moved out of the directory and there are no sub-Directories remaining,
// the DirMgr directory will likewise be deleted.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  targetDMgr   DirMgr - An instance of 'DirMgr' initialized with the directory
//                        path of the target directory to which selected files
//                        will be moved. If the target directory does not exist,
//                        this method will attempt to create it.
//
//
//  fileSelectCriteria FileSelectionCriteria -
//    This input parameter should be configured with the desired file
//    selection criteria. Files matching this criteria will be moved
//    from the current DirMgr path to the directory identified by input
//    parameter, 'targetDir'.
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
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
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
//                                FileSelectMode.ANDSelect() - File selected if all active selection
//                                      criteria are satisfied.
//
//                                    	If this constant value is specified for the file selection mode,
//                                    	then a given file will not be judged as 'selected' unless all of
//                                    	the active selection criterion are satisfied. In other words, if
//                                    	three active search criterion are provided for 'FileNamePatterns',
//                                    	'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                    	selected unless it has satisfied all three criterion in this example.
//
//                                FileSelectMode.ORSelect() - File selected if any active selection
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
// the files processed in the directory tree will be selected and moved
// to the target directory.
//
//      Example:
//        fsc := FileSelectCriterionMode{}
//
//        In this example, 'fsc' is NOT initialized. Therefore,
//        all of the selection criterion are 'Inactive'. Consequently,
//        all of the files encountered in the target directory during
//        the search operation will be selected and moved
//        to the target directory.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) MoveDirectory(
  targetDMgr DirMgr,
  fileSelectCriteria FileSelectionCriteria) (errs []error) {

  errs = make([]error, 0, 300)

  ePrefix := "DirMgr.MoveDirectory() "
  var err, err2, err3 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  err = targetDMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    err2 = fmt.Errorf("Input parameter 'targetDMgr' is INVALID!\n"+
      "Error='%v'\n", err.Error())
    errs = append(errs, err2)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {
    err2 = fmt.Errorf(ePrefix+
      "Error: Source DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'", dMgr.absolutePath)
    errs = append(errs, nonPathError)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  dMgr.dataMutex.Lock()

  dir, err := os.Open(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)

    return errs
  }

  fh := FileHelper{}
  osPathSeparatorStr := string(os.PathSeparator)

  var src, target string
  var isMatch bool
  var nameFileInfos []os.FileInfo
  numOfSubDirectories := 0
  numOfSrcFiles := 0
  err3 = nil

  for err3 != io.EOF {

    nameFileInfos, err3 = dir.Readdir(1000)

    if err3 != nil && err3 != io.EOF {
      _ = dir.Close()
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Readdirnames(-1). "+
        "dMgr.absolutePath='%v'\nError='%v'\n\n",
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        numOfSubDirectories++
        continue

      }

      numOfSrcFiles++

      // This is not a directory. It is a file.
      // Determine if it matches the find file criteria.
      isMatch, err =
        fh.FilterFileName(nameFInfo, fileSelectCriteria)

      if err != nil {

        err2 =
          fmt.Errorf("\n"+ePrefix+
            "Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
            "directorySearched='%v'\nfileName='%v'\nError='%v'\n",
            dMgr.absolutePath, nameFInfo.Name(), err.Error())

        errs = append(errs, err2)

        continue
      }

      if !isMatch {

        continue

      } else {
        // We have a match

        // Create Directory if needed
        if !targetDMgr.DoesAbsolutePathExist() {

          err = targetDMgr.MakeDir()

          if err != nil {
            err2 = fmt.Errorf("\n"+ePrefix+
              "Error creating target directory!\n"+
              "Target Directory='%v'\nError='%v'\n",
              targetDMgr.absolutePath, err.Error())

            errs = append(errs, err2)

            break
          }
        }

        src = dMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        target = targetDMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        err = fh.MoveFile(src, target)

        if err != nil {
          err2 = fmt.Errorf("\n"+ePrefix+
            "ERROR: fh.MoveFile(src, target)\n"+
            "src='%v'\ntarget='%v'\nError='%v'\n\n",
            src, target, err.Error())

          errs = append(errs, err2)

        }

        numOfSrcFiles--
      }
    }
  }

  err = dir.Close()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned by dir.Close(). "+
      "dir='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
  }

  // If all the source files have been moved and
  // there are no sub-directories, DELETE the
  // directory (dMgr).
  if numOfSrcFiles == 0 && numOfSubDirectories == 0 {
    _ = dMgr.DeleteAll()
  }

  return errs
}

// MoveDirectoryTree - Moves all sub-directories and files plus files in
// the parent DirMgr directory to a target directory tree specified by
// input parameter 'targetDMgr'. If successful, the parent directory
// DirMgr will be deleted along with the entire sub-directory tree.
//
// --------------------------------------------------------------------
//
// !!!! BE CAREFUL !!!! This method will delete the entire directory
// tree identified by DirMgr along with ALL the files in the DirMgr
// directory tree!
//
// --------------------------------------------------------------------
//
// Input Parameters:
//
//  targetDMgr   DirMgr - An instance of 'DirMgr' initialized with the directory
//                        path of the target directory to which all source files
//                        will be moved. If the target directory does not exist,
//                        this method will attempt to create it.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
func (dMgr *DirMgr) MoveDirectoryTree(targetDMgr DirMgr) (errs []error) {

  ePrefix := "DirMgr.MoveDirectoryTree() "
  errs = make([]error, 0, 300)
  var err, err2 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  err = targetDMgr.IsDirMgrValid("targetDMgr ")

  if err != nil {
    err2 = fmt.Errorf(ePrefix+"Input parameter targetDMgr is INVALID!\n"+
      "Error='%v'", err.Error())

    errs = append(errs, err2)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {
    err2 = fmt.Errorf(ePrefix+
      "Error: Source DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'", dMgr.absolutePath)
    errs = append(errs, nonPathError)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  fsc := FileSelectionCriteria{}
  errs2 := dMgr.copyDirectoryTree(
    targetDMgr,
    true,
    false,
    ePrefix,
    fsc)

  if len(errs2) > 0 {
    err2 = fmt.Errorf("Errors occurred while copying directory tree to target directory.\n"+
      "Source Directory='%v'\nTarget Directory='%v'\nErrors Follow:\n\n",
      dMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath())
    errs = append(errs, err2)
    errs = append(errs, errs2...)
    return errs
  }

  // The entire directory tree was copied.
  // Now delete the current directory tree
  // to complete the move operation.

  dMgr.dataMutex.Lock()

  err = os.RemoveAll(dMgr.absolutePath)

  dMgr.dataMutex.Unlock()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+"Files were copied successfuly to target directory.\n"+
      "However, errors occurred while deleting the source directory tree. os.RemoveAll(dMgr.GetAbsolutePath())\n"+
      "dMgr.GetAbsolutePath() (DirMgr)='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), err.Error())

    errs = append(errs, err2)
  }

  return errs
}

// MoveSubDirectoryTree - Moves all sub-directories and their constituent
// files from the source or parent directory 'DirMgr' to a target directory
// tree specified by input parameter 'targetDMgr'. If successful, all
// sub-directories and files in the source directory tree will be deleted.
// The source or parent directory identified by 'DirMgr' and the files
// within 'DirMgr' will NOT be deleted.
//
// --------------------------------------------------------------------
//
// !!!! BE CAREFUL !!!! This method will delete the entire sub-directory
// tree. The source or parent directory 'DirMgr' and its constituent files
// will NOT be deleted.
//
// --------------------------------------------------------------------
//
// Input Parameters:
//
//  targetDMgr   DirMgr - An instance of 'DirMgr' initialized with the directory
//                        path of the target directory to which all source files
//                        will be moved. If the target directory does not exist,
//                        this method will attempt to create it.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  errs     []error  - An array of errors is returned. If the method completes
//                      successfully with no errors, a ZERO-length array is
//                      is returned.
//
//                      If errors are encountered they are stored in the error
//                      array and returned to the caller.
//
func (dMgr *DirMgr) MoveSubDirectoryTree(targetDMgr DirMgr) (errs []error) {

  ePrefix := "DirMgr.MoveSubDirectoryTree() "
  errs = make([]error, 0, 300)
  var err, err2 error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  err = targetDMgr.IsDirMgrValid("targetDMgr ")

  if err != nil {
    err2 = fmt.Errorf(ePrefix+"Input parameter targetDMgr is INVALID!\n"+
      "Error='%v'", err.Error())

    errs = append(errs, err2)
    return errs
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    errs = append(errs, nonPathError)
    return errs
  }

  if !dirPathDoesExist {
    err2 = fmt.Errorf(ePrefix+
      "Error: Source DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'", dMgr.absolutePath)
    errs = append(errs, nonPathError)
    return errs
  }

  if !fInfoPlus.IsDir() {
    err2 = fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)

    errs = append(errs, err2)
    return errs
  }

  fsc := FileSelectionCriteria{}
  errs2 := dMgr.copyDirectoryTree(
    targetDMgr,
    true,
    true,
    ePrefix,
    fsc)

  if len(errs2) > 0 {
    err2 = fmt.Errorf("Errors occurred while copying directory tree to target directory.\n"+
      "Source Directory='%v'\nTarget Directory='%v'\nErrors Follow:\n\n",
      dMgr.GetAbsolutePath(), targetDMgr.GetAbsolutePath())
    errs = append(errs, err2)
    errs = append(errs, errs2...)
    return errs
  }

  // The entire directory tree was copied.
  // Now delete the current directory tree
  // to complete the move operation.

  errs2 = dMgr.DeleteAllSubDirectories()

  if len(errs2) > 0 {
    err2 = fmt.Errorf(ePrefix+"Files were copied successfuly to target directory.\n"+
      "However, errors occurred while deleting the sub-directory tree.\n"+
      "Source Directory (DirMgr)='%v'\nErrors Follow:\n\n",
      dMgr.GetAbsolutePath())

    errs = append(errs, err2)
    errs = append(errs, errs2...)
  }

  return errs
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

  dMgr.dataMutex.Lock()

  fh := FileHelper{}

  err = nil
  isEmpty = true
  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    isEmpty = true
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' is an empty string!")
    dMgr.dataMutex.Unlock()
    return isEmpty, err
  }

  if errCode == -2 {
    isEmpty = true
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!")
    dMgr.dataMutex.Unlock()
    return isEmpty, err
  }

  adjustedTrimmedPathStr := fh.AdjustPathSlash(pathStr)

  finalPathStr, isEmptyPath, err2 := fh.GetPathFromPathFileName(adjustedTrimmedPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. fh.GetPathFromPathFileName(pathStr) "+
      "pathStr='%v'  Error='%v'", pathStr, err2.Error())
    isEmpty = isEmptyPath
    dMgr.dataMutex.Unlock()
    return isEmpty, err
  }

  if isEmptyPath {
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. 'pathStr' generated an Empty path! pathStr='%v' ",
      pathStr)
    dMgr.dataMutex.Unlock()
    return isEmpty, err
  }

  errCode, _, finalPathStr = fh.isStringEmptyOrBlank(finalPathStr)

  if errCode < 0 {
    err = fmt.Errorf(ePrefix+
      "Error: path returned from fh.GetPathFromPathFileName(pathStr) is EMPTY! "+
      "pathStr='%v'", pathStr)
    isEmpty = true
    dMgr.dataMutex.Unlock()
    return isEmpty, err
  }

  dMgr.originalPath = adjustedTrimmedPathStr

  dMgr.path = finalPathStr

  dMgr.isPathPopulated = true

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    fh.doesPathFileExist(
      dMgr.path,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.path")

  if nonPathError != nil {
    dMgr.dataMutex.Unlock()
    dMgr.Empty()
    err = nonPathError
    isEmpty = true
    return isEmpty, err
  }

  if !dirPathDoesExist {
    dMgr.doesPathExist = false

  } else {

    if !fInfoPlus.IsDir() {
      dMgr.dataMutex.Unlock()
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "DMgr='%v'\n", dMgr.absolutePath)
      isEmpty = true
      return isEmpty, err
    }

    dMgr.doesPathExist = true
  }

  // Create absolute path
  if strings.ToLower(dMgr.path) == strings.ToLower(fp.VolumeName(dMgr.path)) {

    dMgr.absolutePath = fh.AdjustPathSlash(dMgr.path)

  } else {

    dMgr.absolutePath, err2 = fh.MakeAbsolutePath(dMgr.path)

    if err2 != nil {
      dMgr.dataMutex.Unlock()
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "- fh.MakeAbsolutePath(dMgr.path) returned error. dMgr.path='%v' Error='%v'",
        dMgr.path, err2.Error())
      isEmpty = true
      return isEmpty, err
    }
  }

  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError =
    fh.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    dMgr.dataMutex.Unlock()
    dMgr.Empty()
    err = nonPathError
    isEmpty = true
    return isEmpty, err
  }

  if dirPathDoesExist {

    if !fInfoPlus.IsDir() {
      dMgr.dataMutex.Unlock()
      dMgr.Empty()
      err = fmt.Errorf(ePrefix+
        "- The Directory Manager absolute path exists and IS NOT A DIRECTORY!.\n"+
        "DirMgr Path='%v'", dMgr.absolutePath)
      isEmpty = true
      return
    }

    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = fInfoPlus.CopyOut()

  } else {
    dMgr.doesAbsolutePathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
  }

  dMgr.isAbsolutePathPopulated = true

  strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
  lStr := len(strAry)
  idxStr := strAry[lStr-1]

  idx := strings.Index(dMgr.absolutePath, idxStr)

  dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

  dMgr.isParentPathPopulated = true

  if dMgr.parentPath == "" {
    dMgr.isParentPathPopulated = false
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

  dMgr.dataMutex.Unlock()

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

// SetPermissions - Sets the read/write and execute permissions for the directory
// identified by the current DirMgr instance. Note the treatment of execute
// permissions may vary by operating system.
//
// The permissions are configured based on input parameter 'permissionConfig' which
// is of type, 'FilePermissionConfig'. For an explanation of permission codes, see
// method 'FilePermissionConfig.New()'.
//
// If the directory identified by the current DirMgr instance does not exist, an
// error will be returned.
//
func (dMgr *DirMgr) SetPermissions(permissionConfig FilePermissionConfig) error {
  ePrefix := "DirMgr.SetPermissions() "
  var err error

  err = dMgr.IsDirMgrValid(ePrefix)

  if err != nil {
    return err
  }

  err = permissionConfig.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+"Input parameter 'permissionConfig' is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }
  _,
    dirPathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if nonPathError != nil {
    return nonPathError
  }

  if !dirPathDoesExist {
    return fmt.Errorf(ePrefix+
      "ERROR: DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'\n", dMgr.absolutePath)
  }

  if !fInfoPlus.IsDir() {
    return fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)
  }

  dMgr.dataMutex.Lock()

  err = FileHelper{}.ChangeFileMode(dMgr.absolutePath, permissionConfig)

  dMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+"Input parameter 'permissionConfig' is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

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
    err = fmt.Errorf(ePrefix+"The base directory was NOT found in the current DirMgr path!\n"+
      "DirMgr Path='%v'\nbaseDir Path='%v'\n",
      thisDirAbsPath, oldBaseAbsPath)

    return newDMgr, err
  }

  if idx != 0 {
    err = fmt.Errorf(ePrefix+"The base directory was NOT found at the beginning of the current DirMgr path!\n"+
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
    err = fmt.Errorf(ePrefix+"ERROR: New generated Directory Path Is Invalid!\n"+
      "newAbsPath='%v'\n", newAbsPath)
    return newDMgr, err
  }

  err = nil
  return newDMgr, err
}

// copyDirectoryTree - This method is a private helper method which is designed
// to be called by other public methods for this type. As such, there is no error
// checking performed on 'dMgr' or 'targetDMgr'.
//
// The calling method is responsible for ensuring that both 'dMgr' and 'targetDMgr'
// are valid DirMgr instances. This is accomplished by calling DirMgr.IsDirMgrValid()
// on both instances. In addition, the calling method is also responsible for verifying
// that the directory path identified by 'dMgr' is valid and currently exists on disk.
//
func (dMgr *DirMgr) copyDirectoryTree(
  targetDMgr DirMgr,
  copyEmptyDirectories bool,
  skipTopLevelDirectory bool,
  errorPrefixLabel string,
  fileSelectCriteria FileSelectionCriteria) (errs []error) {

  ePrefix := "DirMgr.copyDirectoryTree() "

  errs = make([]error, 0, 300)

  fh := FileHelper{}

  errCode := 0

  errCode, _, errorPrefixLabel = fh.isStringEmptyOrBlank(errorPrefixLabel)

  if errCode >= 0 {
    ePrefix = errorPrefixLabel + " "
  }

  var err, err2, err3 error

  baseDirLen := len(dMgr.absolutePath)
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dir *os.File
  var nextTargetDMgr DirMgr
  var isMatch, isTopLevelDir bool
  var srcFile, targetFile string
  var nextDir DirMgr

  nextDir = dMgr.CopyOut()

  if nextDir.isInitialized == false {
    err = errors.New(ePrefix + "ERROR: dMgr.CopyOut() returned a nil value!")
    errs = append(errs, err)
    return errs
  }

  for nextDir.isInitialized {

    dir, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "Error return by os.Open(dMgr.absolutePath). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {
        err2 = fmt.Errorf(ePrefix+
          "Error return from #1 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    nextTargetDMgr, err = DirMgr{}.New(
      targetDMgr.GetAbsolutePath() +
        nextDir.absolutePath[baseDirLen:])

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "Error return by os.Open(dMgr.absolutePath). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {

        err2 = fmt.Errorf(ePrefix+
          "Error return from #2 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    if baseDirLen == len(nextDir.GetAbsolutePath()) {
      isTopLevelDir = true
    } else {
      isTopLevelDir = false
    }

    if isTopLevelDir &&
      !skipTopLevelDirectory &&
      copyEmptyDirectories {

      err = nextTargetDMgr.MakeDir()

    } else if !isTopLevelDir && copyEmptyDirectories {

      err = nextTargetDMgr.MakeDir()

    } else {
      err = nil
    }

    if err != nil {
      err2 = fmt.Errorf("\n"+ePrefix+
        "Error creating target directory!\n"+
        "Target Directory='%v'\nError='%v'\n",
        nextTargetDMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {

        err2 = fmt.Errorf(ePrefix+
          "Error return from #3 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dir.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dir.Readdirnames(-1). "+
          "dMgr.absolutePath='%v'\nError='%v'\n",
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          err = dirs.AddDirMgrByPathNameStr(nextDir.absolutePath + osPathSepStr + nameFInfo.Name())

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by dirs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

        } else if isTopLevelDir && skipTopLevelDirectory {
          // Do NOT process files for top level directory
          // when parameter skipTopLevelDirectory = 'true'

          continue

        } else {
          // This is a file which is eligible for processing

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, fileSelectCriteria)

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
                "directorySearched='%v'\nfileName='%v'\nError='%v'\n",
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match

            // Create Directory if needed
            if !nextTargetDMgr.DoesAbsolutePathExist() {

              err = nextTargetDMgr.MakeDir()

              if err != nil {
                err2 = fmt.Errorf("\n"+ePrefix+
                  "Error creating targetFile directory!\n"+
                  "Target Directory='%v'\nError='%v'\n",
                  nextTargetDMgr.absolutePath, err.Error())

                errs = append(errs, err2)
                err3 = io.EOF
                break
              }
            }

            srcFile = nextDir.absolutePath +
              osPathSepStr + nameFInfo.Name()

            targetFile = nextTargetDMgr.absolutePath +
              osPathSepStr + nameFInfo.Name()

            err = fh.CopyFileByIo(srcFile, targetFile)

            if err != nil {
              err2 = fmt.Errorf("\n"+ePrefix+
                "ERROR: fh.CopyFileByIo(srcFile, targetFile)\n"+
                "srcFile='%v'\ntargetFile='%v'\nError='%v'\n\n",
                srcFile, targetFile, err.Error())

              errs = append(errs, err2)
            }
          }
        }
      }
    }

    if dir != nil {
      err = dir.Close()

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by dir.Close(). "+
          "dir='%v' Error='%v' ",
          dMgr.absolutePath, err.Error())

        errs = append(errs, err2)
      }
    }

    nextDir, err = dirs.PopFirstDirMgr()
    if err != nil && err != io.EOF {

      err2 = fmt.Errorf(ePrefix+
        "Error return from #4 dirs.PopFirstDirMgr()\n"+
        "Error='%v'\n", err.Error())

      errs = append(errs, err2)
      return

    } else if err != nil {

      break
    }

  }

  return errs
}

// deleteDirectoryTreeFiles - Helper method used to delete files using file
// selection criteria. Scope of scans and file deletions is controlled by
// input parameter 'scanSubDirectories'. If set to 'true' files may be deleted
// in the entire directory tree. If set to 'false' the file deletions are limited
// solely to the directory identified by the current 'DirMgr' instance.
//
// Since this is a helper method for public methods associated with this type
// no validation checking is performed on DirMgr or the existence of the DirMgr
// path. The DirMgr path is therefore assumed to exist. Parameter validation
// should therefore be performed by the calling method.
//
func (dMgr *DirMgr) deleteDirectoryTreeFiles(
  scanSubDirectories bool,
  errorPrefixLabel string,
  deleteFileSelectionCriteria FileSelectionCriteria) (numOfSubDirectories,
  numOfRemainingFiles,
  numOfDeletedFiles int,
  errs []error) {

  ePrefix := "DirMgr.deleteDirectoryTreeFiles() "

  if len(errorPrefixLabel) > 0 {
    ePrefix = errorPrefixLabel
  }

  numOfSubDirectories = 0
  numOfRemainingFiles = 0
  numOfDeletedFiles = 0
  errs = make([]error, 0, 300)

  var err2, err, err3 error

  osPathSepStr := string(os.PathSeparator)
  var xNumOfSubDirectories, xNumOfTotalFiles, xNumOfDeletedFiles int

  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dir *os.File
  var isMatch bool
  var nextDir DirMgr
  nextDir = dMgr.CopyOut()

  if !nextDir.isInitialized {
    err2 = errors.New(ePrefix + "Error returned by dMgr.CopyOut().\n" +
      "nextDir.isInitialized='false'.")

    errs = append(errs, err2)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  fh := FileHelper{}

  for nextDir.isInitialized {

    dir, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "Error return by os.Open(dMgr.absolutePath). "+
        "dMgr.absolutePath='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {
        err2 = fmt.Errorf(ePrefix+"Error returned by #1 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())

        errs = append(errs, err2)

        return numOfSubDirectories,
          numOfRemainingFiles,
          numOfDeletedFiles,
          errs

      } else if err != nil {

        break
      }

      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dir.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dir.Readdirnames(-1). "+
          "dMgr.absolutePath='%v'\nError='%v'\n",
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          xNumOfSubDirectories++

          if !scanSubDirectories {
            continue
          }

          err = dirs.AddDirMgrByPathNameStr(nextDir.absolutePath + osPathSepStr + nameFInfo.Name())

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by dirs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

        } else {
          // This is a file which is eligible for processing

          xNumOfTotalFiles++

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, deleteFileSelectionCriteria)

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
                "directorySearched='%v'\nfileName='%v'\nError='%v'\n",
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match, delete the file

            err = os.Remove(nextDir.absolutePath + osPathSepStr + nameFInfo.Name())

            if err != nil {
              err2 = fmt.Errorf("\n"+ePrefix+
                "ERROR returned by os.Remove(pathFileName)\n"+
                "pathFileName='%v'\nError='%v'\n\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(),
                err.Error())

              errs = append(errs, err2)

            }

            xNumOfDeletedFiles++
          }
        }
      }
    }

    if dir != nil {
      err = dir.Close()

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by dir.Close(). "+
          "dir='%v' Error='%v' ",
          dMgr.absolutePath, err.Error())

        errs = append(errs, err2)
      }
    }

    nextDir, err = dirs.PopFirstDirMgr()

    if err != nil && err != io.EOF {
      err2 = fmt.Errorf(ePrefix+"Error returned by #2 dirs.PopFirstDirMgr()\n"+
        "Error='%v'\n", err.Error())

      errs = append(errs, err2)

      return numOfSubDirectories,
        numOfRemainingFiles,
        numOfDeletedFiles,
        errs

    } else if err != nil {

      break
    }

  }

  numOfSubDirectories = xNumOfSubDirectories
  numOfRemainingFiles = xNumOfTotalFiles - xNumOfDeletedFiles
  numOfDeletedFiles = xNumOfDeletedFiles

  return numOfSubDirectories,
    numOfRemainingFiles,
    numOfDeletedFiles,
    errs
}

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
