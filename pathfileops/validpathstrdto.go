package pathfileops

import (
  "errors"
  "fmt"
)

// ValidPathStrDto - Used to transfer file/path string attributes and
// associated errors.
//
type ValidPathStrDto struct {
  isInitialized bool          // signals whether the current ValidPathStrDto instance
                              //  has been properly initialized.

  originalPathStr string      // The original, unformatted path string

  pathStr string              // The path string which may or may not be
                              //  the absolute path

  pathFInfoPlus FileInfoPlus  // Only populated if absValidPath exists on disk.

  pathDoesExist PathExistsStatusCode  // -1 = don't know, file/path existence has not been tested
                                      //  0 - No, tests show the file/path doesn't exist on disk.
                                      //  1 - Yes, tests show the file/path does exist on disk.

  pathStrLength int               // Length of the path string

  absPathStr string               // The absolute path version of 'path'

  absPathFInfoPlus 	FileInfoPlus  // Only populated if absValidPath
                                  // exists on disk.

  absPathDoesExist PathExistsStatusCode // -1 = don't know, has not been tested
                                        //  0 - No, tests shown path doesn't exist
                                        //  1 - Yes, tests show path does exist

  absPathStrLength int                  // Length of the absolute path string

  pathType PathFileTypeCode             // The path type. Path File, Path Directory

  pathIsValid PathValidityStatusCode    // -1 - don't know
                                        //  0 - No path is NOT valid
                                        //  1 - Yes, path is valid

  pathVolumeName string           // Volume name associated with current path

  pathVolumeIndex int             // Index of the starting character of Volume Name
                                  // in the path string.

  pathVolumeStrLength int         // Length of the Volume name in the path string.

  err error                       // If no error is encountered
                                  // this value is nil
}

func (vpDto ValidPathStrDto) New() ValidPathStrDto {
  newValPathDto := ValidPathStrDto{}
  newValPathDto.pathStr = ""
  newValPathDto.pathStrLength = -1
  newValPathDto.absPathStr = ""
  newValPathDto.absPathStrLength = -1
  newValPathDto.pathDoesExist = PathExistsStatus.Unknown()
  newValPathDto.absPathDoesExist = PathExistsStatus.Unknown()
  newValPathDto.pathIsValid = PathValidStatus.Unknown()
  newValPathDto.isInitialized = false
  newValPathDto.pathVolumeName = ""
  newValPathDto.pathVolumeIndex = -1
  newValPathDto.pathVolumeStrLength = 0
  newValPathDto.err = nil

  return newValPathDto
}

func (vpDto *ValidPathStrDto) AbsolutePathDoesExist() PathExistsStatusCode {
  return vpDto.absPathDoesExist
}

// GetAbsPath - "getter" method for internal field,
// ValidPathStrDto.absPathStr .
//
func (vpDto *ValidPathStrDto) GetAbsPath() string {
  return vpDto.absPathStr
}

// GetAbsPathStrLen - "getter" method for internal field,
// ValidPathStrDto.absPathStrLength .
//
func (vpDto *ValidPathStrDto) GetAbsPathStrLen() int {
  return vpDto.absPathStrLength
}

// GetAbsPathFileInfo - "getter" method for internal field,
// ValidPathStrDto.absPathFInfoPlus .
//
func (vpDto *ValidPathStrDto) GetAbsPathFileInfo() FileInfoPlus {
  return vpDto.absPathFInfoPlus
}

// GetOriginalPathStr - "getter" method for internal field,
// ValidPathStrDto.originalPathStr .
//
func (vpDto *ValidPathStrDto) GetOriginalPathStr() string {
  return vpDto.originalPathStr
}

// GetPath - "getter" method for internal field,
// ValidPathStrDto.pathStr .
//
func (vpDto *ValidPathStrDto) GetPath() string {
  return vpDto.pathStr
}

// GetPathFileInfo - "getter" method for internal field,
// ValidPathStrDto.pathFInfoPlus .
//
func (vpDto *ValidPathStrDto) GetPathFileInfo() FileInfoPlus {
  return vpDto.pathFInfoPlus
}

// GetPathIsValid - "getter" method for internal field,
// ValidPathStrDto.pathIsValid .
//
func (vpDto *ValidPathStrDto) GetPathIsValid() PathValidityStatusCode {
  return vpDto.pathIsValid
}

// GetPathStrLen - "getter" method for internal field,
// ValidPathStrDto.pathStrLength .
//
func (vpDto *ValidPathStrDto) GetPathStrLen() int {
  return vpDto.pathStrLength
}

// GetPathType - "getter" method for internal field,
// ValidPathStrDto.pathType .
//
func (vpDto *ValidPathStrDto) GetPathType() PathFileTypeCode {
  return vpDto.pathType
}


// GetPathVolumeName - "getter" method for internal field,
// ValidPathStrDto.pathVolumeName .
//
func (vpDto *ValidPathStrDto) GetPathVolumeName() string {
  return vpDto.pathVolumeName
}

// GetPathVolumeIndex - "getter" method for internal field,
// ValidPathStrDto.pathVolumeIndex .
//
func (vpDto *ValidPathStrDto) GetPathVolumeIndex() int {
  return vpDto.pathVolumeIndex
}

// GetPathVolumeStrLength - "getter" method for internal field,
// ValidPathStrDto.pathVolumeStrLength .
//
func (vpDto *ValidPathStrDto) GetPathVolumeStrLength() int {
  return vpDto.pathVolumeStrLength
}

// GetError - "getter" method for internal field,
// ValidPathStrDto.err . Process errors occurring during
// operations performed on the current ValidPathStrDto
// instance are recorded here.
//
func (vpDto *ValidPathStrDto) GetError() error {
  return vpDto.err
}

// IsDtoValid - Performs an analysis of the internal data structures for
// the current ValidPathStrDto instance to determine its validity.
//
// If the current ValidPathStrDto is invalid, this method will return
// an error. If the instance is valid, this method will return 'nil'.
//
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

  if vpDto.pathIsValid != PathValidStatus.Valid() {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "The ValidPathStrDto 'Path Is Valid flag' is Invalid!\n" +
      "vpDto.pathIsValid=%v'\n", vpDto.pathIsValid.String())
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

  if vpDto.pathDoesExist < PathExistsStatus.Unknown() ||
    vpDto.pathDoesExist > PathExistsStatus.Exists() {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "ValidPathStrDto.pathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.pathDoesExist='%v'\n", vpDto.pathDoesExist)
  }

  if vpDto.absPathDoesExist < PathExistsStatus.Unknown() ||
    vpDto.absPathDoesExist > PathExistsStatus.Exists() {
    return fmt.Errorf(ePrefix +
      "ERROR: This ValidPathStrDto is INVALID!\n" +
      "ValidPathStrDto.absPathDoesExist holds an invalid value.\n" +
      "ValidPathStrDto.absPathDoesExist='%v'\n", vpDto.absPathDoesExist)
  }

  return nil
}

// IsInitialized - Returns a boolean value indicating whether the current instance
// of ValidPathStrDto has been initialized.
//
func (vpDto *ValidPathStrDto) IsInitialized() bool {
  return vpDto.isInitialized
}

// PathDoesExist - Returns a boolean value indicating whether the path
// defined by this ValidPathStrDto instance actually exists on disk.
//
//                                        Int
//     Enumeration Value                 Value         Definition
//  ----------------------------------------------------------------------------
//  PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//                                              tested and status is 'Unknown'.
//
//  PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//                                              and path file name does NOT exist on
//                                              disk.
//
//  PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//                                              and path file name DOES exist on
//                                              disk.
//
func (vpDto *ValidPathStrDto) PathDoesExist() PathExistsStatusCode {
  return vpDto.pathDoesExist
}

// PathIsValid - Returns an PathValidityStatusCode enumeration indicating
// whether the path encapsulated by this ValidPathStrDto instance is
// valid.
//
//                                        Int
//     Enumeration Value                 Value         Definition
//  ----------------------------------------------------------------------------
//  PathValidityStatusCode(0).Unknown()   -1  Path/file name validity has NOT been
//                                            tested and its status as 'Valid' or
//                                            'Invalid' is Unknown.
//
//  PathValidityStatusCode(0).Invalid()    0  Tests have verified that the Path/file
//                                            name is 'Invalid'.
//
//  PathValidityStatusCode(0).Valid()     +1  Tests have verified that the Path/file
//                                            name is 'Valid'.
//
func (vpDto *ValidPathStrDto) PathIsValid() PathValidityStatusCode {
  return vpDto.pathIsValid
}

// SetIsInitialized - "setter" method for internal data field
// vpDto.isInitialized .
//
func (vpDto *ValidPathStrDto) SetIsInitialized( isInitialized bool) {
  vpDto.isInitialized = isInitialized
}

// SetPath - "setter" method for internal data fields
// vpDto.pathStr and  vpDto.pathStrLength.
//
func (vpDto *ValidPathStrDto) SetPath(pathStr string) {
  vpDto.pathStr = pathStr

  vpDto.pathStrLength = len(pathStr)
}

// SetPathFileInfo - "setter" method for internal data field
// vpDto.pathFInfoPlus .
//
func (vpDto *ValidPathStrDto) SetPathFileInfo(fInfPlus FileInfoPlus) {
  vpDto.pathFInfoPlus = fInfPlus.CopyOut()
}

// SetAbsPath - "setter" method for internal data fields
// vpDto.absPathStr and vpDto.absPathStrLength.
//
func (vpDto *ValidPathStrDto) SetAbsPath(absPathStr string) {
  vpDto.absPathStr = absPathStr
  vpDto.absPathStrLength = len(absPathStr)
}

// SetAbsPathDoesExistStatus - Input parameter 'absPathExistStatus' is used to set
// the internal field vpDto.absPathDoesExist which is of type PathExistsStatusCode.
//
//                                        Int
//     Enumeration Value                 Value         Definition
//  ----------------------------------------------------------------------------
//  PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//                                              tested and status is 'Unknown'.
//
//  PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//                                              and path file name does NOT exist on
//                                              disk.
//
//  PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//                                              and path file name DOES exist on
//                                              disk.
//
func (vpDto *ValidPathStrDto) SetAbsPathDoesExistStatus(
  absPathExistStatus PathExistsStatusCode) error {

  err := absPathExistStatus.StatusIsValid()

  if err != nil {
    return fmt.Errorf("ValidPathStrDto.SetAbsPathDoesExistStatus()" +
      " %v\n", err.Error())
  }

  vpDto.absPathDoesExist = absPathExistStatus

  return nil
}

// SetAbsPath - "setter" method for internal data field
// vpDto.absPathFInfoPlus .
//
func (vpDto *ValidPathStrDto) SetAbsPathFileInfo(absFInfo FileInfoPlus) {
  vpDto.absPathFInfoPlus = absFInfo.CopyOut()
}

// SetOriginalPathStr - "setter" method for internal data field
// vpDto.originalPathStr .
//
func (vpDto *ValidPathStrDto) SetOriginalPathStr(originalPathStr string) {
  vpDto.originalPathStr = originalPathStr
}

// SetOriginalPathStr - "setter" method for internal data field
// vpDto.pathDoesExist which is of type PathExistsStatusCode.
//
//                                        Int
//     Enumeration Value                 Value         Definition
//  ----------------------------------------------------------------------------
//  PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//                                              tested and status is 'Unknown'.
//
//  PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//                                              and path file name does NOT exist on
//                                              disk.
//
//  PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//                                              and path file name DOES exist on
//                                              disk.
//
func (vpDto *ValidPathStrDto) SetPathDoesExist(
  pathDoesExist PathExistsStatusCode) error {

  err := pathDoesExist.StatusIsValid()

  if err != nil {
    return fmt.Errorf("ValidPathStrDto.SetPathDoesExist() %v",
      err.Error())
  }

  vpDto.pathDoesExist = pathDoesExist

  return nil
}

// SetPathIsValid - "setter" method for internal data field vpDto.pathIsValid
// which is of type PathValidityStatusCode.
//
//                    Path Validity
//   Method            Status Code
//    Name               Constant       Description
//  ______________________________________________________________________
//
//  PathValidityStatusCode(0).Unknown()     -1  Path/file name validity has NOT been
//                                              tested and its status as 'Valid' or
//                                              'invalid' is 'Unknown'.
//
//  PathValidityStatusCode(0).Invalid()      0  Tests have verified that the Path/file
//                                              name is 'Invalid'.
//
//  PathValidityStatusCode(0).Valid()       +1  Tests have verified that the Path/file
//                                              name is 'Valid'.
//
func (vpDto *ValidPathStrDto) SetPathIsValid(
  pathIsValid PathValidityStatusCode) error {

  err := pathIsValid.StatusIsValid()

  if err != nil {
    return fmt.Errorf("ValidPathStrDto.SetPathIsValid() %v",
      err.Error())
  }

  vpDto.pathIsValid = pathIsValid

  return nil
}

// SetOriginalPathStr - "setter" method for internal data field vpDto.pathDoesExist
// which is of type PathFileTypeCode.
//
//                                 Path File
//           Method                Type Code
//            Name                 Constant            Description
//  _______________________________________________________________________________
//  PathFileTypeCode(0).None()        0           Path/file name type has NOT been
//                                                tested and its status not known.
//
//  PathFileTypeCode.Path()           1           Tests have established that the
//                                                Path/file name string is a
//                                                directory path which does NOT
//                                                contain a file name.
//
//  PathFileTypeCode.PathFile()       2           Tests have established that the
//                                                Path/file name string includes
//                                                both a directory path AND a file
//                                                name.
//
//  PathFileTypeCode.File()           3           Tests have established that the
//                                                Path/file name string consists
//                                                solely of a file name and does
//                                                NOT include a directory path.
//
//  PathFileTypeCode.Volume()         4           Tests have established that the
//                                                Path/file name string consists
//                                                solely of a volume name and does
//                                                NOT include a directory path or
//                                                file name.
//
//  PathFileTypeCode.Indeterminate()  5           Tests have been conducted on the
//                                                Path/file name string, but the
//                                                string cannot be classified and
//                                                its status cannot be determined
//                                                with certainty.
//
func (vpDto *ValidPathStrDto) SetPathType(pathType PathFileTypeCode) {
  vpDto.pathType = pathType
}

// SetPathVolumeName - "setter" method for internal data fields
// vpDto.pathVolumeName and vpDto.pathVolumeStrLength .
//
func (vpDto *ValidPathStrDto) SetPathVolumeName(volumeNameStr string) {
  vpDto.pathVolumeName = volumeNameStr
  vpDto.pathVolumeStrLength = len(volumeNameStr)
}

// SetPathVolumeIndex - "setter" method for internal data field
// vpDto.pathVolumeIndex .
//
func (vpDto *ValidPathStrDto) SetPathVolumeIndex(volIndex int) {
  vpDto.pathVolumeIndex = volIndex
}

// SetError - "setter" method for internal data field
// vpDto.err .
//
func (vpDto *ValidPathStrDto) SetError(err error) {
  vpDto.err = fmt.Errorf("%v", err.Error())
}

