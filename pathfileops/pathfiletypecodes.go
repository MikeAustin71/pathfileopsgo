package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

var mPathFileTypeStringToCode = map[string]PathFileTypeCode{
  "None":        PathFileTypeCode(0).None(),
  "Path":        PathFileTypeCode(0).Path(),
  "PathFile":    PathFileTypeCode(0).PathFile(),
  "File":        PathFileTypeCode(0).File(),
  "Volume":      PathFileTypeCode(0).Volume(),
  "VolumeName":  PathFileTypeCode(0).Volume(),
  "Volume Name": PathFileTypeCode(0).Volume(),
}

var mPathFileTypeLwrCaseStringToCode = map[string]PathFileTypeCode{
  "none":        PathFileTypeCode(0).None(),
  "path":        PathFileTypeCode(0).Path(),
  "pathfile":    PathFileTypeCode(0).PathFile(),
  "file":        PathFileTypeCode(0).File(),
  "volume":      PathFileTypeCode(0).Volume(),
  "volumename":  PathFileTypeCode(0).Volume(),
  "volume name": PathFileTypeCode(0).Volume(),
}

var mPathFileTypeCodeToString = map[PathFileTypeCode]string{
  PathFileTypeCode(0).None():     "None",
  PathFileTypeCode(0).Path():     "Path",
  PathFileTypeCode(0).PathFile(): "PathFile",
  PathFileTypeCode(0).File():     "File",
  PathFileTypeCode(0).Volume():   "Volume",
}

type PathFileTypeCode int

func (pfTyp PathFileTypeCode) None() PathFileTypeCode { return 0 }

func (pfTyp PathFileTypeCode) Path() PathFileTypeCode { return 1 }

func (pfTyp PathFileTypeCode) PathFile() PathFileTypeCode { return 2 }

func (pfTyp PathFileTypeCode) File() PathFileTypeCode { return 3 }

func (pfTyp PathFileTypeCode) Volume() PathFileTypeCode { return 4 }

func (pfTyp PathFileTypeCode) ParseString(
  valueString string,
  caseSensitive bool) (PathFileTypeCode, error) {

  ePrefix := "OsFilePermissionCode.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return PathFileTypeCode(0).None(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters "+
        "valueString='%v' ", valueString)
  }

  var ok bool

  var pathFileTypeCode PathFileTypeCode

  if caseSensitive {

    pathFileTypeCode, ok = mPathFileTypeStringToCode[valueString]

    if !ok {
      return PathFileType.None(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    pathFileTypeCode, ok = mPathFileTypeLwrCaseStringToCode[valueString]

    if !ok {
      return PathFileType.None(),
        errors.New(ePrefix + "Invalid Permission Code!")
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

// Value - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (pfTyp PathFileTypeCode) Value() PathFileTypeCode {

  return pfTyp
}

var PathFileType PathFileTypeCode
