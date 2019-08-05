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

  absPathStrLength int            // Length of the absolute path string

  pathIsValid int                 // -1 - don't know
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
  newValPathDto.pathIsValid = -1
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

func (vpDto *ValidPathStrDto) PathDoesExist() PathExistsStatusCode {
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

