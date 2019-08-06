package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

var mPathValidityStatusCodeToString = map[PathValidityStatusCode]string {
  PathValidityStatusCode(-1): "Unknown",
  PathValidityStatusCode(0) : "Invalid",
  PathValidityStatusCode(1) : "Valid" }


var mPathValidityStatusStringToCode = map[string] PathValidityStatusCode {
  "Unknown" : PathValidityStatusCode(-1),
  "Invalid" : PathValidityStatusCode(0),
  "Valid"   : PathValidityStatusCode(1) }

var mPathValidityStatusLowerCaseStringToCode = map[string] PathValidityStatusCode {
  "unknown" : PathValidityStatusCode(-1),
  "invalid" : PathValidityStatusCode(0),
  "valid"   : PathValidityStatusCode(1) }


// PathValidityStatusCode - This type is an enumeration describing
// the validity of a path or path/file name.
//
//
//                    Path Validity
//   Method            Status Code
//    Name               Constant       Description
//  ______________________________________________________________________
//  Unknown()              -1           Path/file name validity has NOT been
//                                        tested and its status as 'Valid' or
//                                        'invalid' is 'Unknown'.
//
//  Invalid()               0           Tests have verified that the Path/file
//                                        name is 'Invalid'.
//
//  Valid()               +1            Tests have verified that the Path/file
//                                        name is 'Valid'.
//
// 'PathValidityStatusCode' has been adapted to function as an enumeration
// describing the validity of a path/file name string. Since Go does not
// directly support enumerations, the 'PathValidityStatusCode' type has been
// configured to function in a manner similar to classic enumerations found
// in other languages like C#. For additional information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type PathValidityStatusCode int


// Unknown - This status code signals that the validity of the subject
// the file or directory path not yet been established and is therefore
// 'Unknown'. This means that no validity tests have yet been performed
// on the subject file or directory path string.
//
func (pathValid PathValidityStatusCode) Unknown() PathValidityStatusCode {
  return PathValidityStatusCode(-1)
}


// Invalid - This status code signals that a validity test was performed
// on the subject file/directory path string and it was found to be
// 'Invalid'.
//
func (pathValid PathValidityStatusCode) Invalid() PathValidityStatusCode {
  return PathValidityStatusCode(0)
}

// Invalid - This status code signals that a validity test was performed
// on the subject file/directory path string and it was found to be
// a 'Valid' path/file name.
//
func (pathValid PathValidityStatusCode) Valid() PathValidityStatusCode {
  return PathValidityStatusCode(1)
}


// StatusCodesEqual - Compares the current PathValidityStatusCode instance to
// another PathExistsStatusCode instance passed as an input parameter. If the
// two are equal in all respects, this method returns 'true'. Otherwise, this
// method returns 'false'.
//
func (pathValid PathValidityStatusCode) StatusCodesEqual( statusCode PathValidityStatusCode) bool {

  if pathValid == statusCode {
    return true
  }

  return false
}

// StatusIsValid - If the value of the current PathValidityStatusCode instance
// is 'invalid', this method will return an error.
//
// If the PathValidityStatusCode is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (pathValid PathValidityStatusCode) StatusIsValid() error {

  _, ok := mPathValidityStatusCodeToString[pathValid]

  if !ok {
    ePrefix := "PathValidityStatusCode.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current PathValidityStatusCode is INVALID! "+
      "PathValidityStatusCode Value='%v'", int(pathValid))
  }

  return nil
}


// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of mPathValidityStatusCodeToString is returned
// set to the value of the associated enumeration.
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
//	PathExistsStatusCode - Upon successful completion, this method will return a new
//	                       instance of PathExistsStatusCode set to the value of the
//	                       enumeration matched by the string search performed on
//	                       input parameter,'valueString'.
//
//	error                - If this method completes successfully, the returned error
//	                       Type is set equal to 'nil'. If an error condition is encountered,
//	                       this method will return an error Type which encapsulates an
//	                       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	t, err := PathValidityStatusCode(0).ParseString("Valid", true)
//                            OR
//	t, err := PathValidityStatusCode(0).ParseString("Valid()", true)
//                            OR
//	t, err := PathValidityStatusCode(0).ParseString("valid", false)
//
//	For all of the cases shown above,
//  t is now equal to PathValidityStatusCode(0).Valid()
//
//
func (pathValid PathValidityStatusCode) ParseString(
  valueString string, caseSensitive bool) (PathValidityStatusCode, error) {

  ePrefix := "PathValidityStatusCode.ParseString() "

  result := PathValidityStatusCode(-1)

  valueString = strings.TrimLeft(strings.TrimRight(valueString, " "), " ")

  lenValueStr := len(valueString)

  if lenValueStr == 0 {
    return result,
      errors.New(ePrefix +
        "Error: Input parameter 'valueString' is an empty string and therefore INVALID!\n")
  }

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 5 {
    return result,
      fmt.Errorf(ePrefix+
        "Error: Input parameter 'valueString' is INVALID!\n" +
        "Length Less than 5-characters valueString='%v'\n", valueString)
  }

  var ok bool

  if caseSensitive {

    result, ok = mPathValidityStatusStringToCode[valueString]

  } else {
    // Search is NOT Case Sensitive

    valueString = strings.ToLower(valueString)

    result, ok = mPathValidityStatusLowerCaseStringToCode[valueString]

  }

  if !ok {
    result = PathValidityStatusCode(-1)
    return result,
      fmt.Errorf(ePrefix +
        "Error: Invalid PathValidityStatusCode Code!\n" +
        "valueString='%v'\n", valueString)
  }

  return result, nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'PathValidityStatusCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the PathValidityStatusCode value is not found,
//           this method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= PathValidityStatusCode(0).Invalid()
//	str := t.String()
//	    str is now equal to "Invalid"
//
func (pathValid PathValidityStatusCode) String() string {

  label, ok := mPathValidityStatusCodeToString[pathValid]

  if !ok {
    return ""
  }

  return label
}

// Value - Returns the value of the PathExistsStatusCode instance
// as type PathExistsStatusCode.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (pathValid PathValidityStatusCode) StatusValue() PathValidityStatusCode {
  return pathValid
}

// PathValidStatus - public global variable of
// type PathValidityStatusCode.
//
// This variable serves as an easier, short hand
// technique for accessing PathValidityStatusCode
// values.
//
// Usage:
//  PathValidStatus.Unknown()
//  PathValidStatus.Invalid()
//  PathValidStatus.Valid()
//
var PathValidStatus PathValidityStatusCode