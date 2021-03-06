package pathfileops

import (
  "errors"
  "fmt"
  "os"
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

// ConsolidateErrors - Consolidates an array of errors into a
// single error.
func (dMgr DirMgr) ConsolidateErrors(errors []error) error {

  return FileHelper{}.ConsolidateErrors(errors)
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
// Input Parameters (Continued):
//
//  copyEmptyDirectory bool - If set to 'true' the target directory will be
//                            created regardless of whether any files are
//                            copied to that directory. Remember that files
//                            are only copied to the target directory if
//                            they meet file selection criteria specified
//                            by input parameter 'fileSelectCriteria'.
//
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  dirCopyStats DirectoryCopyStats - This type is used to return data on the
//                                    copy operation.
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
  fileSelectCriteria FileSelectionCriteria,
  copyEmptyDirectory bool) (dirCopyStats DirectoryCopyStats, errs []error) {

  ePrefix := "DirMgr.CopyDirectory() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirCopyStats,
    errs = dMgrHlpr.copyDirectory(
    dMgr,
    &targetDMgr,
    fileSelectCriteria,
    copyEmptyDirectory,
    ePrefix,
    "dMgr",
    "targetDMgr")

  dMgr.dataMutex.Unlock()

  return dirCopyStats, errs
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
//  targetDMgr         DirMgr - An instance of 'DirMgr' initialized with the directory
//                              path of the target directory to which selected files
//                              will be copied. If the target directory does not exist,
//                              this method will attempt to create it.
//
//  copyEmptyDirectories bool - If a target directory tree path does not previously exist,
//                              the default behavior is to create that directory ONLY if
//                              files matching the file selection criteria are identified
//                              for that directory. If no files match the file selection
//                              criteria, the default is to NOT create the target directory
//                              path.
//
//                              If the parameter 'copyEmptyDirectories' is set to 'true' all
//                              target directory tree paths will be created regardless of
//                              whether files are copied to those directories.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//
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
  fileSelectCriteria FileSelectionCriteria) (dTreeCopyStats DirTreeCopyStats, errs []error) {

  ePrefix := "DirMgr.CopyDirectoryTree() "
  errs = nil

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dTreeCopyStats,
    errs = dMgrHlpr.copyDirectoryTree(
    dMgr,
    &targetDMgr,
    copyEmptyDirectories,
    false, // skipTopLevelDirectory
    fileSelectCriteria,
    ePrefix,
    "dMgr",
    "targetDMgr")

  dMgr.dataMutex.Unlock()

  return dTreeCopyStats,
    errs
}

// CopyIn - Receives a pointer to an incoming DirMgr object as
// an input parameter and copies the values from the incoming
// object to the current DirMgr object. When the copy operation
// is completed, the current DirMgr object is a duplicate of the
// incoming DirMgr object.
func (dMgr *DirMgr) CopyIn(dMgrIn *DirMgr) {

  dMgrHlpr := dirMgrHelper{}

  dMgrIn.dataMutex.Lock()
  dMgr.dataMutex.Lock()

  dMgrHlpr.copyIn(dMgr, dMgrIn)

  dMgrIn.dataMutex.Unlock()
  dMgr.dataMutex.Unlock()
}

// CopyOut - Makes a duplicate copy of the current DirMgr values and
// returns them in a new DirMgr object.
func (dMgr *DirMgr) CopyOut() DirMgr {

  dOut := DirMgr{}
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dOut = dMgrHlpr.copyOut(dMgr)

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
  fileSelectCriteria FileSelectionCriteria) (dTreeCopyStats DirTreeCopyStats, errs []error) {

  ePrefix := "DirMgr.CopySubDirectoryTree() "
  errs = nil
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dTreeCopyStats,
    errs = dMgrHlpr.copyDirectoryTree(
    dMgr,
    &targetDMgr,
    copyEmptyDirectories,
    true, // skipTopLevelDirectory
    fileSelectCriteria,
    ePrefix,
    "dMgr",
    "targetDMgr")

  dMgr.dataMutex.Unlock()

  return dTreeCopyStats, errs
}

// DeleteAll - BE CAREFUL!!! - This method will remove the directory identified by
// this DirMgr instance. It will also delete all child directories and files in the
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

// DeleteAllFilesInDir - Deletes all the files in the current directory.
// ONLY files in the current directory are deleted. Sub-directories are
// NOT deleted and files in sub-directories are NOT deleted.
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
func (dMgr *DirMgr) DeleteAllFilesInDir() (deleteDirStats DeleteDirFilesStats, errs []error) {

  ePrefix := "DirMgr.DeleteAllFilesInDir() "
  dMgrHlpr := dirMgrHelper{}
  errs = nil

  dMgr.dataMutex.Lock()

  deleteDirStats,
    errs = dMgrHlpr.deleteAllFilesInDirectory(
    dMgr,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return deleteDirStats, errs
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
  ePrefix := "DirMgr.CopyDirectory() "

  errs = make([]error, 0, 300)

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  errs = dMgrHlpr.deleteAllSubDirectories(
    dMgr,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return errs
}

// DeleteDirectoryTreeFiles - Deletes files in the directory tree. The parent
// directory for this tree is the directory specified by the current 'DirMgr'
// instance. Files eligible for deletion must match the file selection criteria
// specified by input parameter 'deleteFileSelectionCriteria'. The file deletion
// operation will search the parent directory ('DirMgr') and all sub-directories
// screening for files which match the file selection criteria.
//
// The file deletion operation is conducted in three steps:
//    1. The criteria for selecting files to be deleted is created using
//       input parameter 'deleteFileSelectionCriteria'.
//
//    2. A file search is conducted which includes the DirMgr parent directory
//       and all sub-directories in the tree.
//
//    3. Files processed during the directory tree search are compared to the
//       file selection criteria specified by 'deleteFileSelectionCriteria'.
//       Those files which match the selection criteria are then deleted.
//
// This method is similar to method 'DirMgr.DeleteWalkDirFiles()'. However, this
// method returns less data and is designed to work with very large numbers of
// files and directories.
//
// Note: As a result of this operation, files within directory tree folders may be
// deleted, but the folders or directory elements will NEVER be deleted.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//  deleteDirStats DeleteDirFilesStats - Statistics generated by the delete operation
//                                       performed on the current directory identified
//                                       by DirMgr.
//
//  errs                      []error  - An array of errors is returned. If the method
//                                       completes successfully with no errors, a
//                                       ZERO-length array is returned.
//
//                                       If errors are encountered they are stored in
//                                       the error array and returned to the caller.
//
func (dMgr *DirMgr) DeleteDirectoryTreeFiles(
  deleteFileSelectionCriteria FileSelectionCriteria) (deleteDirStats DeleteDirFilesStats, errs []error) {
  ePrefix := "DirMgr.DeleteDirectoryTreeFiles() "

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  deleteDirStats,
    errs = dMgrHlpr.deleteDirectoryTreeStats(
    dMgr,
    false, // skip top level (parent) directory
    true,  // scan sub-directories
    deleteFileSelectionCriteria,
    ePrefix,
    "dMgr",
    "deleteFileSelectionCriteria")

  dMgr.dataMutex.Unlock()

  return deleteDirStats, errs
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
func (dMgr *DirMgr) DeleteFilesByNamePattern(
  fileSearchPattern string) (deleteDirStats DeleteDirFilesStats, errs []error) {

  ePrefix := "DirMgr.DeleteFilesByNamePattern() "

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  deleteDirStats,
    errs = dMgrHlpr.deleteFilesByNamePattern(
    dMgr,
    fileSearchPattern,
    ePrefix,
    "dMgr",
    "fileSearchPattern")

  dMgr.dataMutex.Unlock()

  return deleteDirStats, errs
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
  deleteFileSelectionCriteria FileSelectionCriteria) (deleteDirStats DeleteDirFilesStats, errs []error) {

  ePrefix := "DirMgr.DeleteDirectoryTreeFiles() "

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  deleteDirStats,
    errs = dMgrHlpr.deleteDirectoryTreeStats(
    dMgr,
    false, // skip top level (parent) directory
    false, //scan sub-directories
    deleteFileSelectionCriteria,
    ePrefix,
    "dMgr",
    "deleteFileSelectionCriteria")

  dMgr.dataMutex.Unlock()

  return deleteDirStats, errs
}

// DeleteSubDirectoryTreeFiles - Deletes sub-directory files. For this operation, the
// current 'DirMgr' is classified as the top level or parent directory. Files in this
// parent directory will NEVER BE DELETED.
//
// !!! BE CAREFUL !!! This method deletes files in sub-directories!
//
// Files eligible for deletion must match the file selection criteria specified by input
// parameter 'deleteFileSelectionCriteria'. The file deletion operation will exclude the
// parent directory ('DirMgr') and confine the file search to the sub-directories underneath
// the parent directory. The file search will screen for files which match the file selection
// criteria in the sub-directory tree.
//
// The file deletion operation is conducted in three steps:
//
//    1. The criteria for selecting files to be deleted is created using input parameter
//       'deleteFileSelectionCriteria'.
//
//    2. A file search is conducted which excludes the DirMgr parent directory and focuses
//       exclusively on all sub-directories in the tree.
//
//    3. Files processed during the sub-directory tree search are compared to the file
//       selection criteria specified by 'deleteFileSelectionCriteria'. Those files which
//       match the selection criteria are then deleted.
//
// Note: As a result of this operation, files within sub-directory tree folders may be
// deleted, but the folders or directory elements themselves will NEVER be deleted.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//  deleteDirStats DeleteDirFilesStats - Statistics generated by the delete operation
//                                       performed on the current directory identified
//                                       by DirMgr.
//
//  errs                      []error  - An array of errors is returned. If the method
//                                       completes successfully with no errors, a
//                                       ZERO-length array is returned.
//
//                                       If errors are encountered they are stored in
//                                       the error array and returned to the caller.
//
func (dMgr *DirMgr) DeleteSubDirectoryTreeFiles(
  deleteFileSelectionCriteria FileSelectionCriteria) (deleteDirStats DeleteDirFilesStats, errs []error) {
  ePrefix := "DirMgr.DeleteSubDirectoryTreeFiles() "

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  deleteDirStats,
    errs = dMgrHlpr.deleteDirectoryTreeStats(
    dMgr,
    true, // skip top level (parent) directory
    true, // scan sub-directories
    deleteFileSelectionCriteria,
    ePrefix,
    "dMgr",
    "deleteFileSelectionCriteria")

  dMgr.dataMutex.Unlock()

  return deleteDirStats, errs
}

// DeleteWalkDirFiles - !!! BE CAREFUL !!! This method deletes files
// in a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the current DirMgr object which is treated as the
// parent directory. This method 'walks the directory tree' locating
// all files in the directory tree which match the file selection
// criteria submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// This method will delete files in the entire directory tree including
// the parent directory and its sub-directory tree.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
  dMgrHlpr := dirMgrHelper{}
  var err error
  var errs []error

  dMgr.dataMutex.Lock()

  deleteFilesInfo,
    errs =
    dMgrHlpr.deleteDirectoryTreeInfo(
      dMgr,
      deleteFileSelectionCriteria,
      false, // skip top level directory
      true,  // scan sub-directories
      ePrefix,
      "dMgr",
      "deleteFileSelectionCriteria")

  if len(errs) > 0 {
    err = dMgr.ConsolidateErrors(errs)
  }

  dMgr.dataMutex.Unlock()

  return deleteFilesInfo, err
}

/*
func (dMgr *DirMgr) DeleteWalkDirFiles(
  deleteFileSelectionCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {

  ePrefix := "DirMgr.DeleteWalkDirFiles() "
  deleteFilesInfo := DirectoryDeleteFileInfo{}
  dMgrHlpr := dirMgrHelper{}
  var dirPathDoesExist bool
  var err error

  dMgr.dataMutex.Lock()

  dirPathDoesExist,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr")

  if err != nil {
    dMgr.dataMutex.Unlock()
    return deleteFilesInfo, err
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: 'dMgr' Directory Path DOES NOT EXIST!\n"+
      "dMgr='%v'\n",
      dMgr.absolutePath)
    dMgr.dataMutex.Unlock()
    return deleteFilesInfo, err
  }

  deleteFilesInfo.StartPath = dMgr.absolutePath

  deleteFilesInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

  fh := FileHelper{}

  err2 := fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

  if err2 != nil {

    deleteFilesInfo = DirectoryDeleteFileInfo{}

    err = fmt.Errorf(ePrefix+"Error returned by FileHelper."+
      "makeFileHelperWalkDirDeleteFilesFunc(&dWalkInfo). "+
      "dWalkInfo.StartPath='%v' Error='%v' ", deleteFilesInfo.StartPath, err2.Error())
  }

  return deleteFilesInfo, err
}
*/

// DoesAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.absolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
//
func (dMgr *DirMgr) DoesAbsolutePathExist() bool {

  dMgrHlpr := dirMgrHelper{}
  dirPathDoesExist := false
  var err error

  dMgr.dataMutex.Lock()

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      "",
      "dMgr")

  if err != nil {
    dirPathDoesExist = false
  }

  dMgr.dataMutex.Unlock()

  return dirPathDoesExist
}

// DoesDirectoryExist - Returns two boolean values indicating whether or not the
// Directory path exists and whether or not the Directory absolute path exists.
//
func (dMgr *DirMgr) DoesDirectoryExist() (doesPathExist, doesAbsolutePathExist bool) {

  dMgrHlpr := dirMgrHelper{}
  dirPathDoesExist := false
  var err error

  dMgr.dataMutex.Lock()

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      "",
      "dMgr")

  if err != nil {
    dirPathDoesExist = false
  }

  dMgr.dataMutex.Unlock()

  return dirPathDoesExist, dirPathDoesExist
}

// DoesPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition it also
// updates the DirMgr field DirMgr.doesPathExist field.
//
func (dMgr *DirMgr) DoesPathExist() bool {

  dMgrHlpr := dirMgrHelper{}
  dirPathDoesExist := false
  var err error

  dMgr.dataMutex.Lock()

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      "",
      "dMgr")

  if err != nil {
    dirPathDoesExist = false
  }

  dMgr.dataMutex.Unlock()

  return dirPathDoesExist
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

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  directoryDoesExist,
    _,
    nonPathError =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr")

  dMgr.dataMutex.Unlock()

  return directoryDoesExist, nonPathError
}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _ = dMgrHlpr.empty(dMgr, "DirMgr.Empty() ", "dMgr")

  dMgr.dataMutex.Unlock()

}

// Equal - Compares two DirMgr objects to determine if
// they are equal.
func (dMgr *DirMgr) Equal(dMgr2 *DirMgr) bool {

  dMgrHlpr := dirMgrHelper{}
  isEqual := false

  dMgr2.dataMutex.Lock()
  dMgr.dataMutex.Lock()

  isEqual = dMgrHlpr.equal(dMgr, dMgr2)

  dMgr2.dataMutex.Unlock()
  dMgr.dataMutex.Unlock()

  return isEqual
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

  dMgrHlpr := dirMgrHelper{}
  isEqual := false

  dMgr.dataMutex.Lock()
  dMgr2.dataMutex.Lock()
  isEqual = dMgrHlpr.equalAbsolutePaths(
    dMgr,
    dMgr2)
  dMgr2.dataMutex.Unlock()
  dMgr.dataMutex.Unlock()

  return isEqual
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

  dMgrHlpr := dirMgrHelper{}
  isEqual := false

  dMgr.dataMutex.Lock()
  dMgr2.dataMutex.Lock()
  isEqual = dMgrHlpr.equalPaths(
    dMgr,
    dMgr2)
  dMgr2.dataMutex.Unlock()
  dMgr.dataMutex.Unlock()

  return isEqual
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//          FileOperationCode(0).MoveSourceFileToDestinationFile() FileOperationCode = iota
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
//    [] error -  This function will return an array of errors containing error messages
//                generated during the performance of specified File Operations on the
//                designated directory. If the error array returned is empty or has
//                a zero length, it signals that no errors were encountered and all operations
//                completed successfully.
//
func (dMgr *DirMgr) ExecuteDirectoryFileOps(
  fileSelectCriteria FileSelectionCriteria,
  fileOps []FileOperationCode,
  targetBaseDir DirMgr) (errs []error) {

  ePrefix := "DirMgr.ExecuteDirectoryFileOps() "

  dMgrHlpr := dirMgrHelper{}

  errs = dMgrHlpr.executeDirectoryFileOps(
    dMgr,
    fileSelectCriteria,
    fileOps,
    &targetBaseDir,
    ePrefix,
    "dMgr",
    "targetBaseDir",
    "fileSelectCriteria",
    "fileOps")

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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//    FileOperationCode(0).MoveSourceFileToDestinationFile()
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
//  []error -  This function will return an array of strings containing error messages
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

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  errs = dMgrHlpr.executeDirectoryTreeOps(
    dMgr,
    fileSelectCriteria,
    fileOps,
    &targetBaseDir,
    ePrefix,
    "dMgr",
    "targetBaseDir",
    "fileOps")

  dMgr.dataMutex.Unlock()

  return errs
}

// FindDirectoryTreeFiles - This method returns file information on files residing in a
// specific directory tree identified by the current DirMgr instance. The directory
// identified by 'DirMgr' is treated as the the parent directory for the search.
//
// In addition to file information, this method also returns data on the directory tree
// being searched including the parent directory and all sub-directories in the tree.
//
// This method 'walks the directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// All directories including the top level parent directory ('DirMgr') are searched. This
// differs from method 'DirMgr.FindWalkSubDirFiles()' which only searches the sub-directory
// tree.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to zero
// values or 'Inactive', then ALL FILES in the directory are selected and returned in
// the field, 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the returned collection
// 'DirectoryTreeInfo.Directories'. This returned 'DirectoryTreeInfo.Directories'
// collection will always include the top level parent directory identified by 'DirMgr'.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//  dTreeInfo - DirectoryTreeInfo structure	-
//            type DirectoryTreeInfo struct {
//              StartPath             string                // The starting path or directory for the file
//                                                          // search
//
//              Directories           DirMgrCollection      // Directory Managers found during directory tree
//                                                          // search.
//
//                                                          // This collection will ALWAYS return the parent
//                                                          // directory ('DirMgr') as the first entry in the
//                                                          // collection.
//
//              FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//              ErrReturns            []error               // Internal System errors encountered
//              FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//                                                          // input parameter to this method.
//             }
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
func (dMgr *DirMgr) FindDirectoryTreeFiles(
  fileSelectionCriteria FileSelectionCriteria) (dTreeInfo DirectoryTreeInfo, errs []error) {

  ePrefix := "DirMgr.FindDirectoryTreeFiles() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dTreeInfo,
    errs = dMgrHlpr.findDirectoryTreeFiles(
    dMgr,
    fileSelectionCriteria,
    false,
    true,
    ePrefix,
    "dMgr",
    "fileSelectionCriteria")

  dMgr.dataMutex.Unlock()

  return dTreeInfo, errs

}

// FindFilesByNamePattern - searches the current directory using a name pattern file
// search criteria.
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

  dMgrHlpr := dirMgrHelper{}
  fileMgrCol := FileMgrCollection{}.New()
  var err error

  dMgr.dataMutex.Lock()

  fileMgrCol,
    err = dMgrHlpr.findFilesByNamePattern(
    dMgr,
    fileSearchPattern,
    ePrefix,
    "dMgr",
    "fileSearchPattern")

  dMgr.dataMutex.Unlock()

  return fileMgrCol, err
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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

  dMgrHlpr := dirMgrHelper{}

  dTreeInfo := DirectoryTreeInfo{}
  var err error

  dMgr.dataMutex.Lock()

  dTreeInfo,
  errs := dMgrHlpr.findDirectoryTreeFiles(
    dMgr,
    fileSelectCriteria,
    false, // skip top level directory
    false, // scan sub-directories
    ePrefix,
    "dMgr",
    "fileSelectCriteria")

  if len(errs) > 0 {
    err = dMgr.ConsolidateErrors(errs)
  }

  dMgr.dataMutex.Unlock()

  if err != nil {
    return FileMgrCollection{}, err
  }

  return dTreeInfo.FoundFiles, err
}

// FindWalkDirFiles - This method returns file information on files residing in a
// specific directory tree identified by the current DirMgr instance. The directory
// identified by 'DirMgr' is treated as the the parent directory for the search.
//
// In addition to file information, this method also returns data on the directory tree
// being searched including the parent directory and all sub-directories in the tree.
//
// This method 'walks the directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// All directories including the top level parent directory ('DirMgr') are searched. This
// differs from method 'DirMgr.FindWalkSubDirFiles()' which only searches the sub-directory
// tree.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to zero
// values or 'Inactive', then ALL FILES in the directory are selected and returned in
// the field, 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the returned collection
// 'DirectoryTreeInfo.Directories'. This returned 'DirectoryTreeInfo.Directories'
// collection will always include the top level parent directory identified by 'DirMgr'.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//            Directories           DirMgrCollection      // Directory Managers found during directory tree search.
//                                                        // This collection will ALWAYS return the parent directory
//                                                        // ('DirMgr') as the first entry in the collection.
//            FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//            ErrReturns            []error               // Internal System errors encountered
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

  ePrefix := "DirMgr.FindWalkDirFiles() "
  dTreeInfo := DirectoryTreeInfo{}
  var err error
  var errs []error
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dTreeInfo,
    errs = dMgrHlpr.findDirectoryTreeFiles(
    dMgr,
    fileSelectCriteria,
    false, // skip top level directory
    true,  // scan sub-directories
    ePrefix,
    "dMgr",
    "fileSelectCriteria")

  if len(errs) > 0 {
    err = dMgr.ConsolidateErrors(errs)
  }

  dMgr.dataMutex.Unlock()

  return dTreeInfo, err
}

// FindWalkSubDirFiles - This method returns file information on files residing in a
// sub-directory tree identified by the current DirMgr instance. As such, this method
// will NOT search the top level directory, parent directory identified by the current
// DirMgr instance. However, all directories subsidiary to the parent directory ('DirMgr')
// will be searched.
//
// This method 'walks the directory tree' locating all files in the sub-directory tree which
// match the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to zero
// values or 'Inactive', then ALL FILES in the suc-directory tree are selected and returned
// in the field, 'DirectoryTreeInfo.FoundFiles'.
//
//
// All directories searched will be included in the returned collection 'DirectoryTreeInfo.Directories'.
// If the parent directory has NO sub-directories, this returned collection will be empty.
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//            Directories           DirMgrCollection      // Directory Managers found during directory tree
//                                                        // search. Note: The top level parent directory will
//                                                        // NEVER be included.
//            FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//            ErrReturns            []error               // Internal System errors encountered
//            FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//                                                        // input parameter to this method.
//          }
//
//          If successful, files matching the file selection criteria input
//          parameter shown above will be returned in a 'DirectoryTreeInfo'
//          object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//          on all the files in the specified sub-directory tree which match the file
//          selection criteria.
//
//          Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//                to determine if any internal system errors were encountered while processing
//                the directory tree.
//
//  error - If a program execution error is encountered during processing, it will
//          be returned as an 'error' type. Also, see the comment on
//          'DirectoryTreeInfo.ErrReturns', above.
//
func (dMgr *DirMgr) FindWalkSubDirFiles(
  fileSelectCriteria FileSelectionCriteria) (dTreeInfo DirectoryTreeInfo, err error) {

  ePrefix := "DirMgr.FindWalkSubDirFiles() "
  dMgrHlpr := dirMgrHelper{}
  var errs []error

  dMgr.dataMutex.Lock()

  dTreeInfo,
    errs = dMgrHlpr.findDirectoryTreeFiles(
    dMgr,
    fileSelectCriteria,
    true, // skip top level directory
    true, // scan sub-directories
    ePrefix,
    "dMgr",
    "fileSelectCriteria")

  if len(errs) > 0 {
    err = dMgr.ConsolidateErrors(errs)
  }

  dMgr.dataMutex.Unlock()

  /*
     dTreeInfo,
       err = dMgrHlpr.findFilesWalkDirectoryTree(
       dMgr,
       fileSelectCriteria,
       true, // skipTopLevelDirectory
       ePrefix,
       "dMgr")
  */
  return dTreeInfo, err
}

// GetAbsolutePath - Returns a string containing the
// absolute path for the current Directory Manager
// instance. This string returned by this method
// will NOT have a trailing path separator.
//
// See companion method GetAbsolutePathLc() to
// acquire a lower case version of absolute path.
//
func (dMgr *DirMgr) GetAbsolutePath() string {

  dMgrHlpr := dirMgrHelper{}
  absolutePath := ""

  dMgr.dataMutex.Lock()

  _,
  _,
  err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  if err != nil {
    absolutePath = ""
  } else {
    absolutePath = dMgr.absolutePath
  }

  dMgr.dataMutex.Unlock()

  return absolutePath
}


// GetAbsolutePath - Returns a string containing the
// low case version of the absolute path for the current
// Directory Manager instance.

// This string returned by this method will NOT have a
// trailing path separator. It will consist of all lower
// case characters.
//
// See the companion method GetAbsolutePath() to return
// an absolute path string with upper and lower case
// characters.
//
func (dMgr *DirMgr) GetAbsolutePathLc() string {
  dMgrHlpr := dirMgrHelper{}
  absolutePath := ""

  dMgr.dataMutex.Lock()
  _,
  _,
  err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  if err != nil {
    absolutePath = ""
  } else {
    absolutePath = strings.ToLower(dMgr.absolutePath)
  }

  dMgr.dataMutex.Unlock()

  return absolutePath
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
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  pathElements, _ = dMgrHlpr.getAbsolutePathElements(dMgr, "", "")

  dMgr.dataMutex.Unlock()

  return pathElements
}

// GetAbsolutePathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character. The path string may consist of upper and
// lower case characters.
//
// See the companion method GetAbsolutePathWithSeparatorLc()
// which returns a path string consisting of all lower case
// characters.
//
func (dMgr *DirMgr) GetAbsolutePathWithSeparator() string {

  dMgrHlpr := dirMgrHelper{}

  absolutePath := ""

  dMgr.dataMutex.Lock()

  _,
  _,
  err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  if err != nil {
    absolutePath = ""
  } else {
    absolutePath = dMgr.absolutePath
  }

  dMgr.dataMutex.Unlock()

  lPath := len(absolutePath)

  if lPath == 0 {
    return ""
  }

  if absolutePath[lPath-1] != os.PathSeparator {
    return absolutePath + string(os.PathSeparator)
  }

  return absolutePath
}

// GetAbsolutePathWithSeparatorLc - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character. The path string will consists of all lower
// case characters.
//
// See the companion method GetAbsolutePathWithSeparator()
// which returns a path string consisting of upper and lower
// case characters.
//
func (dMgr *DirMgr) GetAbsolutePathWithSeparatorLc() string {

  dMgrHlpr := dirMgrHelper{}

  absolutePath := ""

  dMgr.dataMutex.Lock()

  _,
  _,
  err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  if err != nil {
    absolutePath = ""
  } else {
    absolutePath = strings.ToLower(dMgr.absolutePath)
  }

  dMgr.dataMutex.Unlock()

  lPath := len(absolutePath)

  if lPath == 0 {
    return ""
  }

  if absolutePath[lPath-1] != os.PathSeparator {
    return absolutePath + string(os.PathSeparator)
  }

  return absolutePath
}

// GetDirectoryBytes - Returns the number of bytes in the current directory identified
// by the 'DirMgr' instance. This method only returns bytes in the current directory
//
func (dMgr *DirMgr) GetDirectoryStats() (dirStats DirectoryStatsDto, errs []error) {

  ePrefix := "DirMgr.GetDirectoryStats() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirStats,
    errs = dMgrHlpr.findDirectoryTreeStats(
    dMgr,
    false,
    true,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return dirStats, errs
}

// GetDirectoryTreeBytes - Returns all the bytes in a directory tree.
// The parent directory for the search is identified by the current
// DirMgr instance.
//
func (dMgr *DirMgr) GetDirectoryTreeStats() (dirStats DirectoryStatsDto, errs []error) {

  ePrefix := "DirMgr.GetDirectoryTreeStats() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirStats,
    errs = dMgrHlpr.findDirectoryTreeStats(
    dMgr,
    false,
    true,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return dirStats, errs
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

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirMgrs, errs =
    dMgrHlpr.getDirectoryTree(dMgr, ePrefix, "dMgr")

  dMgr.dataMutex.Unlock()

  return dirMgrs, errs
}

// GetDirectoryName - Returns a string containing the name
// of the directory without out the parent path.
//
func (dMgr *DirMgr) GetDirectoryName() string {

  directoryName := ""
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  if !dMgr.isInitialized {

    directoryName = ""

  } else {

    directoryName = dMgr.directoryName

  }

  dMgr.dataMutex.Unlock()

  return directoryName
}

// GetFileInfoPlus - Returns a FileInfoPlus instance detailing file
// system information on the directory identified by the current
// Directory Manager instance.
//
func (dMgr *DirMgr) GetFileInfoPlus() (FileInfoPlus, error) {

  ePrefix := "DirMgr.GetFileInfoPlus() "
  fileInfoPlus := FileInfoPlus{}
  var err error
  var dirDoesExist bool
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirDoesExist,
    fileInfoPlus,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr")

  if err == nil && !dirDoesExist {
    fileInfoPlus = FileInfoPlus{}
    err = fmt.Errorf(ePrefix+"DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'\n", dMgr.absolutePath)
  }

  dMgr.dataMutex.Unlock()

  return fileInfoPlus, err
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
  fileInfoPlus := FileInfoPlus{}
  var err error
  var dirDoesExist bool
  dMgrHlpr := dirMgrHelper{}
  fPermCfg := FilePermissionConfig{}

  dMgr.dataMutex.Lock()

  dirDoesExist,
    fileInfoPlus,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr")

  if err == nil && !dirDoesExist {

    err = fmt.Errorf(ePrefix+
      "DirMgr Path DOES NOT EXIST!\n"+
      "DirMgr Path='%v'\n",
      dMgr.absolutePath)

  } else if err == nil && dirDoesExist {

    fPermCfg, err = FilePermissionConfig{}.NewByFileMode(fileInfoPlus.Mode())
  }

  dMgr.dataMutex.Unlock()

  return fPermCfg, err
}

// GetNumberOfAbsPathElements - Returns the number of elements
// or path components in the absolute path of the current
// Directory Manager instance.
func (dMgr *DirMgr) GetNumberOfAbsPathElements() int {

  pathElements := make([]string, 0, 50)
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  pathElements, _ = dMgrHlpr.getAbsolutePathElements(dMgr, "", "")

  dMgr.dataMutex.Unlock()

  return len(pathElements)
}

// GetOriginalPath - Returns the original path used to initialize
// this Directory Manager instance.
//
func (dMgr *DirMgr) GetOriginalPath() string {
  originalPath := ""

  dMgr.dataMutex.Lock()

  if !dMgr.isInitialized {

    originalPath = ""

  } else {

    originalPath = dMgr.originalPath

  }

  dMgr.dataMutex.Unlock()

  return originalPath
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
func (dMgr *DirMgr) GetParentDirMgr() (dirMgrOut DirMgr, hasParent bool, err error) {

  ePrefix := "DirMgr.GetParentDirMgr() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirMgrOut,
    hasParent,
    err = dMgrHlpr.getParentDirMgr(
    dMgr,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return dirMgrOut, hasParent, err
}

// GetParentPath - Returns a string containing the
// parent path for the current Directory Manager
// instance. The Parent Path string will NOT contain
// a trailing path separator.
//
func (dMgr *DirMgr) GetParentPath() string {

  ePrefix := "DirMgr.GetParentDirMgr() "
  dMgrHlpr := dirMgrHelper{}
  dirMgrOut := DirMgr{}
  parentPath := ""
  var err error

  dMgr.dataMutex.Lock()

  dirMgrOut,
    _,
    err = dMgrHlpr.getParentDirMgr(
    dMgr,
    ePrefix,
    "dMgr")

  if err != nil ||
    !dMgr.isInitialized {
    parentPath = ""
  } else {
    parentPath = dirMgrOut.absolutePath
  }

  dMgr.dataMutex.Unlock()

  return parentPath
}

// GetPath - Returns the path used to configure this
// Directory Manager Instance. It will NOT contain a
// trailing path separator. It may or may not be an
// absolute path.
//
func (dMgr *DirMgr) GetPath() string {

  dMgrHlpr := dirMgrHelper{}
  dPath := ""

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  if len(dMgr.path) == 0 ||
    !dMgr.isInitialized {
    dPath = ""
  } else {
    dPath = dMgr.path
  }

  dMgr.dataMutex.Unlock()

  return dPath
}

// GetPathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetPathWithSeparator() string {
  dMgrHlpr := dirMgrHelper{}
  dPath := ""

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  if len(dMgr.path) == 0 ||
    !dMgr.isInitialized {

    dPath = ""

  } else {

    dPath = dMgr.path
  }

  dMgr.dataMutex.Unlock()

  lPath := len(dPath)

  if lPath == 0 {
    return ""
  }

  if dPath[lPath-1] != os.PathSeparator {
    return dPath + string(os.PathSeparator)
  }

  return dPath
}

// GetVolumeName - Returns a string containing the volume name
// of the directory identified by the current Directory Manager
// instance.
//
func (dMgr *DirMgr) GetVolumeName() string {

  volumeName := ""
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  if len(dMgr.volumeName) == 0 ||
    !dMgr.isInitialized {

    volumeName = ""

  } else {

    volumeName = dMgr.volumeName

  }

  dMgr.dataMutex.Unlock()

  return volumeName
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

  isInitialized := false

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  isInitialized = dMgr.isInitialized

  dMgr.dataMutex.Unlock()

  return isInitialized
}

// IsParentPathPopulated - Returns a boolean value
// indicating whether the parent path for this Directory
// Manager instance is populated.
//
func (dMgr *DirMgr) IsParentPathPopulated() bool {

  dMgrHlpr := dirMgrHelper{}
  isParentPathPopulated := false

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.getParentDirMgr(
    dMgr,
    "",
    "dMgr")

  if len(dMgr.parentPath) == 0 ||
    !dMgr.isInitialized {
    isParentPathPopulated = false
  } else {
    isParentPathPopulated = true
  }

  dMgr.dataMutex.Unlock()

  return isParentPathPopulated
}

// IsPathPopulated - Returns a boolean value indicating
// whether the current Directory Manager path string is
// populated.
//
func (dMgr *DirMgr) IsPathPopulated() bool {

  dMgrHlpr := dirMgrHelper{}
  isDMgrPathPopulated := false

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.getParentDirMgr(
    dMgr,
    "",
    "dMgr")

  if len(dMgr.path) == 0 ||
    !dMgr.isInitialized {
    isDMgrPathPopulated = false
  } else {
    isDMgrPathPopulated = true
  }

  dMgr.dataMutex.Unlock()

  return isDMgrPathPopulated
}

func (dMgr *DirMgr) ParseValidPathStr(pathStr string) (ValidPathStrDto, error) {

  dMgrHlpr := dirMgrHelper{}

  validPathDto,
  err := dMgrHlpr.getValidPathStr(
    pathStr,
    "DirMgr.ParseValidPathStr() ",
    "pathStr")

  return validPathDto, err
}

// IsVolumeNamePopulated - Returns a boolean value indicating
// whether the Volume Name for the current Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsVolumeNamePopulated() bool {

  dMgrHlpr := dirMgrHelper{}
  isDMgrVolumePopulated := false

  dMgr.dataMutex.Lock()

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "dMgr")

  if len(dMgr.volumeName) == 0 ||
    !dMgr.isInitialized {

    isDMgrVolumePopulated = false

  } else {

    isDMgrVolumePopulated = true

  }

  dMgr.dataMutex.Unlock()

  return isDMgrVolumePopulated
}

// MakeDir - If the directory path identified by the current DirMgr
// object does not exist, this method will create that directory path.
// The path will be created using permission specifications passed through
// input parameter 'fPermCfg'.
//
func (dMgr *DirMgr) MakeDirWithPermission(fPermCfg FilePermissionConfig) error {

  ePrefix := "DirMgr.MakeDirWithPermission() "
  var err error
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    err = dMgrHlpr.lowLevelMakeDirWithPermission(
    dMgr,
    fPermCfg,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  return err
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
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  _,
    err = dMgrHlpr.lowLevelMakeDir(
    dMgr,
    ePrefix,
    "dMgr")

  dMgr.dataMutex.Unlock()

  // No errors - directory created.
  return err
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
//                                         err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                         err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
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
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
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
//  numOfSrcFilesMoved     int - Contains the number of files actually moved to the target
//                               directory.
//
//  numOfSrcFilesRemaining int - Returns the number of source files remaining in the current
//                               directory after the move operation was completed.
//
//  numOfSubDirectories    int - Returns the number of sub-directories which exist in the
//                               DirMgr parent directory.
//
//  dMgrDirWasDeleted     bool - Returns 'true' if the current directory ('DirMgr') was deleted.
//
//  errs              []error  - An array of errors is returned. If the method completes
//                               successfully with no errors, a ZERO-length array is
//                               is returned.
//
//                               If errors are encountered they are stored in the error
//                               array and returned to the caller.
//
func (dMgr *DirMgr) MoveDirectory(
  targetDMgr DirMgr,
  fileSelectCriteria FileSelectionCriteria) (dirMoveStats DirectoryMoveStats,
  errs []error) {

  ePrefix := "DirMgr.MoveDirectory() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirMoveStats,
    errs = dMgrHlpr.moveDirectory(
    dMgr,
    &targetDMgr,
    fileSelectCriteria,
    ePrefix,
    "dMgr",
    "targetDMgr",
    "fileSelectCriteria")

  dMgr.dataMutex.Unlock()

  return dirMoveStats, errs
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
//
func (dMgr *DirMgr) MoveDirectoryTree(
  targetDMgr DirMgr) (
  dirMoveStats DirectoryMoveStats,
  errs []error) {

  ePrefix := "DirMgr.MoveDirectoryTree() "

  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirMoveStats,
    errs = dMgrHlpr.moveDirectoryTree(
    dMgr,
    &targetDMgr,
    ePrefix,
    "dMgr",
    "targetDMgr")

  dMgr.dataMutex.Unlock()

  return dirMoveStats, errs
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
func (dMgr *DirMgr) MoveSubDirectoryTree(targetDMgr DirMgr) (
  dirMoveStats DirectoryMoveStats, errs []error) {

  ePrefix := "DirMgr.MoveSubDirectoryTree() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  dirMoveStats,
    errs =
    dMgrHlpr.moveSubDirectoryTree(
      dMgr,
      &targetDMgr,
      ePrefix,
      "dMgr",
      "destinationDMgr")

  dMgr.dataMutex.Unlock()

  return dirMoveStats, errs
}

// New - Returns a new DirMgr object and populates the
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

  ePrefix := "DirMgr.New() "
  dMgrHlpr := dirMgrHelper{}

  newDirMgr := DirMgr{}

  isEmpty,
  err := dMgrHlpr.setDirMgr(
    &newDirMgr,
    pathStr,
    ePrefix,
    "newDirMgr",
    "pathStr")

  if err != nil {
    return DirMgr{}, err
  }

  if isEmpty {
    return DirMgr{}, fmt.Errorf(ePrefix+
      "\nERROR: dMgrHlpr.SetDirMgr(pathStr) returned an EMPTY DirMgr\n"+
      "pathStr='%v'\n",
      pathStr)
  }

  return newDirMgr, nil
}

// NewFromFileInfo - Returns a new DirMgr object based on two input parameters:
// 		- A parent directory path string
//		- An os.FileInfo object containing the directory name.
//
func (dMgr DirMgr) NewFromFileInfo(
  parentDirectoryPath string, info os.FileInfo) (DirMgr, error) {

  ePrefix := "DirMgr.NewFromFileInfo() "
  dMgrHlpr := dirMgrHelper{}
  newDirMgr := DirMgr{}

  if info == nil {
    return DirMgr{},
      errors.New(ePrefix +
        "ERROR: Input parameter 'info' is 'nil' and INVALID!\n")
  }

  isEmpty, err := dMgrHlpr.setDirMgrWithPathDirectoryName(
    &newDirMgr,
    parentDirectoryPath,
    info.Name(),
    ePrefix,
    "newDirMgr",
    "parentDirectoryPath",
    "FileInfo.Name()")

  if err != nil {
    return DirMgr{}, err
  }

  if isEmpty {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "Newly generated 'DirMgr' is Empty!\n"+
        "dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
        "parentDirectoryPath='%v'\n"+
        "FileInfo.Name()='%v'\n",
        parentDirectoryPath,
        info.Name())
  }

  return newDirMgr, nil
}

// NewFromDirMgrFileInfo - Configures and returns a new 'DirMgr' instance based on
// two input parameters, 'directory' and 'fileNameExt'.
//
// Input parameter 'directory' is of type 'DirMgr' and is treated as the parent directory.
// The final directory name is provided by the input parameter 'fileInfo' of type
// 'os.FileInfo'.
//
func (dMgr DirMgr) NewFromDirMgrFileInfo(
  parentDirectory DirMgr, directoryFileInfo os.FileInfo) (DirMgr, error) {

  ePrefix := "DirMgr.NewFromDirMgrFileInfo() "

  if directoryFileInfo == nil {
    return DirMgr{},
      errors.New(ePrefix +
        "\nERROR: Input parameter 'directoryFileInfo' is 'nil' and therefore invalid!\n")
  }

  var err error

  err = parentDirectory.IsDirMgrValid("")

  if err != nil {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "\nInput parameter 'parentDirectory' is invalid!\n"+
        "%v", err.Error())
  }

  dMgrHlpr := dirMgrHelper{}

  newDirMgr := DirMgr{}

  isEmpty := false

  isEmpty,
    err = dMgrHlpr.setDirMgrFromKnownPathDirName(
    &newDirMgr,
    parentDirectory.GetAbsolutePath(),
    directoryFileInfo.Name(),
    ePrefix,
    "newDirMgr",
    "parentDirectory",
    "directoryFileInfo.Name()")

  if err == nil && isEmpty {
    err = fmt.Errorf(ePrefix+
      "\nERROR: The DirMgr instance generated is empty and contains no data!\n"+
      "parentDirectory='%v'\n"+
      "directory='%v'\n", parentDirectory.GetAbsolutePath(), directoryFileInfo.Name())
  }

  if err != nil {
    return DirMgr{}, err
  }

  return newDirMgr, nil
}

// NewFromFileMgr - Configures and returns a new 'DirMgr' instance based
// on input parameter 'fileMgr' which is of type 'FileMgr'.
//
//
func (dMgr DirMgr) NewFromFileMgr(fileMgr FileMgr) (DirMgr, error) {

  ePrefix := "DirMgr.NewFromFileMgr() "

  err := fileMgr.IsFileMgrValid("")

  if err != nil {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "\nERROR: Input parameter 'fileMgr' is invalid!\n"+
        "%v", err.Error())
  }

  return fileMgr.GetDirMgr(), nil
}

// NewFromKnownPathDirectoryName - Configures and returns
// a new 'DirMgr' instance using a parent path name and
// directory name. The parent path and directory name are
// combined to form the full path for the new 'DirMgr'
// instance.
//
// This method will populate all internal field values
// with new values based on input parameters 'parentPathName'
// and 'directoryName'.
//
// This method differs from similar methods in that it assumes
// the input parameters are known values and do not require
// the usual analysis and validation screening applied by
// other methods.
//
func (dMgr DirMgr) NewFromKnownPathDirectoryName(
  parentPathName string, directoryName string) (DirMgr, error) {

  ePrefix := "DirMgr.NewFromKnownPathDirectoryName() "

  newDirMgr := DirMgr{}

  dMgrHlpr := dirMgrHelper{}

  var isEmpty bool
  var err error

  isEmpty,
    err = dMgrHlpr.setDirMgrFromKnownPathDirName(
    &newDirMgr,
    parentPathName,
    directoryName,
    ePrefix,
    "newDirMgr",
    "parentPathName",
    "directoryName")

  if err != nil {
    return DirMgr{}, err
  }

  if isEmpty {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "Newly generated 'DirMgr' is Empty!\n"+
        "dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
        "parentPathName='%v'\n"+
        "directoryName='%v'\n",
        parentPathName,
        directoryName)
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
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  isEmpty     bool - If the outcome of setting new values for DirMgr is an
//                     'empty' DirMgr instance, this value will be set to 'true'.
//
//  err        error - If a program execution error is encountered during
//                     processing, it will be returned as an 'error' type.
//
func (dMgr *DirMgr) SetDirMgr(pathStr string) (isEmpty bool, err error) {

  ePrefix := "DirMgr.SetDirMgr() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  isEmpty,
    err = dMgrHlpr.setDirMgr(
    dMgr,
    pathStr,
    ePrefix,
    "dMgr",
    "pathStr")

  dMgr.dataMutex.Unlock()

  return isEmpty, err
}

// SetDirMgrFromKnownPathDirectoryName - Configures the internal
// field values for the current DirMgr instance using a parent
// path name and a directory name. The parent path and directory
// name are combined to form the full path for the current 'DirMgr'
// instance.
//
// This method will replace all previous field values with new
// values based on input parameters 'parentPathName' and
// 'directoryName'.
//
// This method differs from other "Set" methods in that it
// assumes the input parameters are known values and do not
// require the usual analysis and validation screening applied
// by similar methods.
//
// If more rigours input parameter validation is required,
// consider using method, DirMgr.SetDirMgr().
//
func (dMgr *DirMgr) SetDirMgrFromKnownPathDirName(
  parentPathName, directoryName string) (isEmpty bool, err error) {

  ePrefix := "DirMgr.setDirMgrFromKnownPathDirName() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  isEmpty,
    err = dMgrHlpr.setDirMgrFromKnownPathDirName(
    dMgr,
    parentPathName,
    directoryName,
    ePrefix,
    "dMgr",
    "parentPathName",
    "directoryName")

  dMgr.dataMutex.Unlock()

  return isEmpty, err

}

// SetDirMgrWithFileInfo - Sets the DirMgr fields and path strings for the current
// DirMgr object based on an input 'parentDirectoryPath' parameter and an
// os.FileInfo input parameter, 'info'.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//  parentDirectoryPath string - The parent directory path.
//
//  info           os.FileInfo - The os.FileInfo containing the directory name.
//
// ---------------------------------------------------------------------------
//
// Return Value:
//
//  err        error - If a program execution error is encountered during
//                     processing, it will be returned as an 'error' type.
//
func (dMgr *DirMgr) SetDirMgrWithFileInfo(
  parentDirectoryPath string,
  info os.FileInfo) (err error) {

  ePrefix := "DirMgr.SetDirMgrWithFileInfo() "

  dMgrHlpr := dirMgrHelper{}
  isEmpty := true

  if info == nil {
    return errors.New(ePrefix +
      "ERROR: Input parameter 'info' is 'nil' and INVALID!\n")
  }

  dMgr.dataMutex.Lock()

  isEmpty,
    err = dMgrHlpr.setDirMgrWithPathDirectoryName(
    dMgr,
    parentDirectoryPath,
    info.Name(),
    ePrefix,
    "dMgr",
    "parentDirectoryPath",
    "FileInfo.Name()")

  if err == nil && isEmpty {
    err = fmt.Errorf(ePrefix+
      "Newly generated 'DirMgr' is Empty!\n"+
      "dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
      "parentDirectoryPath='%v'\n"+
      "FileInfo.Name()='%v'\n",
      parentDirectoryPath,
      info.Name())
  }

  dMgr.dataMutex.Unlock()

  return err
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
func (dMgr *DirMgr) SetPermissions(permissionConfig FilePermissionConfig) (err error) {

  ePrefix := "DirMgr.SetPermissions() "
  err = nil
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  err = dMgrHlpr.setPermissions(
    dMgr,
    permissionConfig,
    ePrefix,
    "dMgr",
    "permissionConfig")

  dMgr.dataMutex.Unlock()

  return err
}

// SubstituteBaseDir - Substitute 'baseDir' segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
//
func (dMgr *DirMgr) SubstituteBaseDir(
  baseDir DirMgr,
  substituteBaseDir DirMgr) (newDMgr DirMgr, err error) {

  ePrefix := "DirMgr.SubstituteBaseDir() "
  dMgrHlpr := dirMgrHelper{}

  dMgr.dataMutex.Lock()

  newDMgr,
    err = dMgrHlpr.substituteBaseDir(
    dMgr,
    &baseDir,
    &substituteBaseDir,
    ePrefix,
    "DirMgr",
    "baseDir",
    "substituteBaseDir")

  dMgr.dataMutex.Unlock()

  return newDMgr, err
}
