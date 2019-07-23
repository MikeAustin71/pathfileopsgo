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
  ErrReturns               []error
  DeleteFileSelectCriteria FileSelectionCriteria
  DeletedFiles             FileMgrCollection
}

// ValidPathStrDto - Used to transfer valid
// path string attributes and associated errors
//
type ValidPathStrDto struct {
  isInitialized bool

  originalPathStr string  // The original, unformatted path string

  pathStr string    // The path string which may or may not be
  // the absolute path

  pathFInfoPlus FileInfoPlus // Only populated if absValidPath
  // exists on disk.

  pathDoesExist int // -1 = don't know, has not been tested
  //  0 - No it doesnt exist
  //  1 - Yes, it does exist

  pathStrLength int // -1  = unknown
  // >-1 = known correct length

  absPathStr string // The absolute path version of 'path'

  absPathFInfoPlus 	FileInfoPlus // Only populated if absValidPath
                                 // exists on disk.

  absPathDoesExist int // -1 = don't know, has not been tested
  //  0 - No it doesnt exist
  //  1 - Yes, it does exist

  absPathStrLength int // -1  = unknown
  // >-1 = known correct length

  pathIsValid int // -1 - don't know
  //  0 - No path is NOT valid
  //  1 - Yes, path is valid

  pathVolumeName string

  pathVolumeIndex int

  pathVolumeStrLength int

  err error // If no error is encountered
  // this value is nil
}

func (vpDto ValidPathStrDto) New() ValidPathStrDto {
  newValPathDto := ValidPathStrDto{}
  newValPathDto.pathStr = ""
  newValPathDto.pathStrLength = -1
  newValPathDto.absPathStr = ""
  newValPathDto.absPathStrLength = -1
  newValPathDto.pathDoesExist = -1
  newValPathDto.absPathDoesExist = -1
  newValPathDto.pathIsValid = -1
  newValPathDto.isInitialized = false
  newValPathDto.pathVolumeName = ""
  newValPathDto.pathVolumeIndex = -1
  newValPathDto.pathVolumeStrLength = 0
  newValPathDto.err = nil

  return newValPathDto
}

func (vpDto *ValidPathStrDto) AbsolutePathDoesExist() int {
  return vpDto.absPathDoesExist
}

func (vpDto *ValidPathStrDto) PathDoesExist() int {
  return vpDto.pathDoesExist
}

func (vpDto *ValidPathStrDto) GetPath() string {
  return vpDto.pathStr
}

func (vpDto *ValidPathStrDto) GetPathStrLen() int {
  return vpDto.pathStrLength
}

func (vpDto *ValidPathStrDto) GetPathFileInfo() FileInfoPlus {
  return vpDto.pathFInfoPlus
}

func (vpDto *ValidPathStrDto) GetAbsPath() string {
  return vpDto.absPathStr
}

func (vpDto *ValidPathStrDto) GetAbsPathStrLen() int {
  return vpDto.absPathStrLength
}

func (vpDto *ValidPathStrDto) GetAbsPathFileInfo() FileInfoPlus {
  return vpDto.absPathFInfoPlus
}

func (vpDto *ValidPathStrDto) GetOriginalPathStr() string {
  return vpDto.originalPathStr
}

func (vpDto *ValidPathStrDto) GetPathVolumeName() string {
  return vpDto.pathVolumeName
}

func (vpDto *ValidPathStrDto) GetPathVolumeIndex() int {
  return vpDto.pathVolumeIndex
}

func (vpDto *ValidPathStrDto) GetPathVolumeStrLength() int {
  return vpDto.pathVolumeStrLength
}

func (vpDto *ValidPathStrDto) GetError() error {
  return vpDto.err
}

func (vpDto *ValidPathStrDto) PathIsValid() int {
  return vpDto.pathIsValid
}

func (vpDto *ValidPathStrDto) IsDtoValid(ePrefix string) error {

  if len(ePrefix) == 0 {
    ePrefix = "ValidPathStrDto.IsDtoValid() "
  } else {
    ePrefix = ePrefix + "- ValidPathStrDto.IsDtoValid()\n"
  }

  if !vpDto.isInitialized {
    return errors.New(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "This ValidPathStrDto instance is NOT initialized!\n" +
      "vpDto.isInitialized='false'\n")
  }

  if vpDto.pathIsValid != 1 {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "The ValidPathStrDto 'Path Is Valid flag' is Invalid!\n" +
      "vpDto.pathIsValid=%v'\n", vpDto.pathIsValid)
  }

  if len(vpDto.pathStr) == 0 {
    return errors.New(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "The ValidPathStrDto 'pathStr' is EMPTY!\n")
  }

  if len(vpDto.absPathStr) == 0 {
    return errors.New(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "The ValidPathStrDto absolute path string is EMPTY!\n")
  }

  if vpDto.pathDoesExist < -1 || vpDto.pathDoesExist > 1 {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "ValidPathStrDto.pathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.pathDoesExist='%v'\n", vpDto.pathDoesExist)
  }

  if vpDto.absPathDoesExist < -1 || vpDto.absPathDoesExist > 1 {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "ValidPathStrDto.absPathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.absPathDoesExist='%v'\n", vpDto.absPathDoesExist)
  }

  return nil
}

func (vpDto *ValidPathStrDto) IsPathExistenceTestValid(ePrefix string) error {

  if len(ePrefix) == 0 {
    ePrefix = "ValidPathStrDto.IsPathExistenceTestValid() "
  } else {
    ePrefix = ePrefix + "- ValidPathStrDto.IsPathExistenceTestValid()\n"
  }

  if vpDto.pathDoesExist < 0 || vpDto.pathDoesExist > 1 {
    return fmt.Errorf(ePrefix +
      "ERROR: The ValidPathStrDto Path Existence Test is INVALID!\n" +
      "ValidPathStrDto.pathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.pathDoesExist='%v'\n", vpDto.pathDoesExist)
  }

  if vpDto.absPathDoesExist < 0 || vpDto.absPathDoesExist > 1 {
    return fmt.Errorf(ePrefix +
      "ERROR: The ValidPathStrDto Absolute Path Existence Test is INVALID!\n" +
      "ValidPathStrDto.absPathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.absPathDoesExist='%v'\n", vpDto.absPathDoesExist)
  }

  return nil
}

func (vpDto *ValidPathStrDto) IsInitialized() bool {
  return vpDto.isInitialized
}
