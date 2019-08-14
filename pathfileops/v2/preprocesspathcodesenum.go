package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

var mPreProcessPathStringToCode = map[string]PreProcessPathCode{
  "None":          PreProcessPathCode(0).None(),
  "PathSeparator": PreProcessPathCode(0).PathSeparator(),
  "AbsolutePath":  PreProcessPathCode(0).AbsolutePath(),
}

var mPreProcessPathLwrCaseStringToCode = map[string]PreProcessPathCode{
  "none":          PreProcessPathCode(0).None(),
  "pathseparator": PreProcessPathCode(0).PathSeparator(),
  "absolutepath":  PreProcessPathCode(0).AbsolutePath(),
}

var mPreProcessPathCodeToString = map[PreProcessPathCode]string{
  PreProcessPathCode(0).None():          "None",
  PreProcessPathCode(0).PathSeparator():  "PathSeparator",
  PreProcessPathCode(0).AbsolutePath():   "AbsolutePath",
}

type PreProcessPathCode int

// None - Take No Action
func (preProcPathCde PreProcessPathCode) None() PreProcessPathCode {
  return PreProcessPathCode(0)
}

// PathSeparator - Convert path separators to the default value for the
// host operating system.
func (preProcPathCde PreProcessPathCode) PathSeparator() PreProcessPathCode {
  return PreProcessPathCode(1)
}

// AbsolutePath - Convert path string to an absolute path.
func (preProcPathCde PreProcessPathCode) AbsolutePath() PreProcessPathCode {
  return PreProcessPathCode(2)
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'PreProcessPathCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the PreProcessPathCode value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= PreProcessPathCode(0).AbsolutePath()
//	str := t.String()
//	    str is now equal to "AbsolutePath"
//
func (preProcPathCde PreProcessPathCode) String() string {

  label, ok := mPreProcessPathCodeToString[preProcPathCde]

  if !ok {
    return ""
  }

  return label
}

// Value - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (preProcPathCde PreProcessPathCode) Value() PreProcessPathCode {

  return preProcPathCde
}


func (preProcPathCde PreProcessPathCode) ParseString(
  valueString string,
  caseSensitive bool) (PreProcessPathCode, error) {

  ePrefix := "PreProcessPathCode.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return PreProcessPathCode(0).None(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
        "valueString='%v'\n", valueString)
  }

  var ok bool

  var preProcessPathCode PreProcessPathCode

  if caseSensitive {

    preProcessPathCode, ok = mPreProcessPathStringToCode[valueString]

    if !ok {
      return PreProcessPathCode(0).None(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    preProcessPathCode, ok = mPreProcessPathLwrCaseStringToCode[valueString]

    if !ok {
      return PreProcessPathCode(0).None(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  }

  return preProcessPathCode, nil
}


var PreProcPathCode PreProcessPathCode

