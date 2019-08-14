package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

var mPathFileTypeStringToCode = map[string]PathFileTypeCode{
  "None":          PathFileTypeCode(0).None(),
  "Path":          PathFileTypeCode(0).Path(),
  "PathFile":      PathFileTypeCode(0).PathFile(),
  "File":          PathFileTypeCode(0).File(),
  "Volume":        PathFileTypeCode(0).Volume(),
  "VolumeName":    PathFileTypeCode(0).Volume(),
  "Volume Name":   PathFileTypeCode(0).Volume(),
  "Indeterminate": PathFileTypeCode(0).Indeterminate(),
  "Unknown":       PathFileTypeCode(0).Indeterminate(),
}

var mPathFileTypeLwrCaseStringToCode = map[string]PathFileTypeCode{
  "none":          PathFileTypeCode(0).None(),
  "path":          PathFileTypeCode(0).Path(),
  "pathfile":      PathFileTypeCode(0).PathFile(),
  "file":          PathFileTypeCode(0).File(),
  "volume":        PathFileTypeCode(0).Volume(),
  "volumename":    PathFileTypeCode(0).Volume(),
  "volume name":   PathFileTypeCode(0).Volume(),
  "indeterminate": PathFileTypeCode(0).Indeterminate(),
  "unknown":       PathFileTypeCode(0).Indeterminate(),
}

var mPathFileTypeCodeToString = map[PathFileTypeCode]string{
  PathFileTypeCode(0).None():          "None",
  PathFileTypeCode(0).Path():          "Path",
  PathFileTypeCode(0).PathFile():      "PathFile",
  PathFileTypeCode(0).File():          "File",
  PathFileTypeCode(0).Volume():        "Volume",
  PathFileTypeCode(0).Indeterminate(): "Indeterminate",
}


// PathFileTypeCode - This type is an enumeration describing the status
// of a path or path/file name.
//
//
//                      Path File
//   Method             Type Code
//    Name              Constant        Description
//  ______________________________________________________________________
//  None()                  0           Path/file name type has NOT been
//                                        tested and its status not known.
//
//  Path()                  1           Tests have established that the
//                                        Path/file name string is a
//                                        directory path which does NOT
//                                        contain a file name.
//
//  PathFile()              2           Tests have established that the
//                                        Path/file name string includes
//                                        both a directory path AND a file
//                                        name.
//
//  File()                  3           Tests have established that the
//                                        Path/file name string consists
//                                        solely of a file name and does
//                                        NOT include a directory path.
//
//  Volume()                4           Tests have established that the
//                                        Path/file name string consists
//                                        solely of a volume name and does
//                                        NOT include a directory path or
//                                        file name.
//
//  Indeterminate()         5           Tests have been conducted on the
//                                        Path/file name string, but the
//                                        string cannot be classified and
//                                        its status cannot be determined
//                                        with certainty.
//
// 'PathFileTypeCode' has been adapted to function as an enumeration
// describing the type of a path/file name string. Since Go does not
// directly support enumerations, the 'PathFileTypeCode' type has been
// configured to function in a manner similar to classic enumerations found
// in other languages like C#. For additional information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type PathFileTypeCode int

// None - 'None' signals that the PathFileTypeCode has not been set or initialized.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) None() PathFileTypeCode { return PathFileTypeCode(0) }

// Path - This code signals that the subject path string has been classified as a
// directory path.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) Path() PathFileTypeCode { return PathFileTypeCode(1) }

// PathFile - This code signals that subject path string has been classified as a
// file path meaning that the path string consists of both a directory path and a
// file name.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) PathFile() PathFileTypeCode { return PathFileTypeCode(2) }

// File - This code signals that the subject path string consists only of a file name
// and does NOT include a directory path.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) File() PathFileTypeCode { return PathFileTypeCode(3) }

// Volume - The code signals that the entire path string consists solely of a volume
// and does NOT include a directory path or file name.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) Volume() PathFileTypeCode { return PathFileTypeCode(4) }

// Indeterminate - This code signals that after tests, the status of the subject path
// string cannot be established. This code covers edge cases where it cannot be determined
// whether the string is a Path and File or a directory path.
//
// This method is part of the standard enumeration.
//
func (pfTyp PathFileTypeCode) Indeterminate() PathFileTypeCode { return PathFileTypeCode(5) }

// ParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// PathFileTypeCode is returned set to the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'valid' will NOT
//	                       match the enumeration name, 'Valid'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'valid' will match the
//	                       enumeration name 'Valid'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	PathFileTypeCode    - Upon successful completion, this method will return a new
//	                      instance of PathExistsStatusCode set to the value of the
//	                      enumeration matched by the string search performed on
//	                      input parameter,'valueString'.
//
//	error               - If this method completes successfully, the returned error
//	                      Type is set equal to 'nil'. If an error condition is encountered,
//	                      this method will return an error Type which encapsulates an
//	                      appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	t, err := PathFileTypeCode(0).ParseString("File", true)
//                            OR
//	t, err := PathValidityStatusCode(0).ParseString("File()", true)
//                            OR
//	t, err := PathValidityStatusCode(0).ParseString("file", false)
//
//	For all of the cases shown above,
//  t is now equal to PathValidityStatusCode(0).File()
//
func (pfTyp PathFileTypeCode) ParseString(
  valueString string,
  caseSensitive bool) (PathFileTypeCode, error) {

  ePrefix := "PathFileTypeCode.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return PathFileTypeCode(0).None(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
        "valueString='%v'\n", valueString)
  }

  var ok bool

  var pathFileTypeCode PathFileTypeCode

  if caseSensitive {

    pathFileTypeCode, ok = mPathFileTypeStringToCode[valueString]

    if !ok {
      return PathFileType.None(),
        errors.New(ePrefix + "Invalid PathFileTypeCode Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    pathFileTypeCode, ok = mPathFileTypeLwrCaseStringToCode[valueString]

    if !ok {
      return PathFileType.None(),
        errors.New(ePrefix + "Invalid PathFileTypeCode Code!")
    }

  }

  return pathFileTypeCode, nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'PathFileTypeCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the PathFileTypeCode value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= PathFileTypeCode(0).PathFile()
//	str := t.String()
//	    str is now equal to "PathFile"
//
func (pfTyp PathFileTypeCode) String() string {

  label, ok := mPathFileTypeCodeToString[pfTyp]

  if !ok {
    return ""
  }

  return label
}

// StatusValue - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (pfTyp PathFileTypeCode) StatusValue() PathFileTypeCode {

  return pfTyp
}

// PathFileType - public global variable of
// type PathFileTypeCode.
//
// This variable serves as an easier, short hand
// technique for accessing PathFileTypeCode values.
//
// Usage:
//  PathFileTypeCode.None()
//  PathFileTypeCode.Path()
//  PathFileTypeCode.File()
//  PathFileTypeCode.Volume()
//  PathFileTypeCode.Indeterminate()
//
var PathFileType PathFileTypeCode
