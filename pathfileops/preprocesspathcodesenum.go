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

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of PreProcessPathCode is returned set to
// the value of the associated enumeration.
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
//	PreProcessPathCode      - Upon successful completion, this method will return a new
//	                          instance of PreProcessPathCode set to the value of the
//	                          enumeration matched by the string search performed on
//	                          input parameter,'valueString'.
//
//	error                   - If this method completes successfully, the returned error
//	                          Type is set equal to 'nil'. If an error condition is encountered,
//	                          this method will return an error Type which encapsulates an
//	                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := PreProcessPathCode(0).ParseString("AbsolutePath", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("AbsolutePath()", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("absolutepath", false)
//
//  For all of the cases shown above,
//  t is now equal to PreProcessPathCode(0).AbsolutePath()
//
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


// StatusIsValid - If the value of the current PreProcessPathCode instance
// is 'invalid', this method will return an error.
//
// If the PreProcessPathCode is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (preProcPathCde PreProcessPathCode) StatusIsValid() error {

  _, ok := mPreProcessPathCodeToString[preProcPathCde]

  if !ok {
    ePrefix := "PreProcessPathCode.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current PreProcessPathCode is INVALID! "+
      "PreProcessPathCode Value='%v'", int(preProcPathCde))
  }

  return nil
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


var PreProcPathCode PreProcessPathCode

