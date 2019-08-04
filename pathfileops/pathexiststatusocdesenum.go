package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

var  mPathExistsStatusCodeToString = map[PathExistsStatusCode]string {
  PathExistsStatusCode(-1): "Unknown",
  PathExistsStatusCode(0) : "DoesNotExist",
  PathExistsStatusCode(1) : "Exists" }

var mPathExistsStatusStringToCode = map[string] PathExistsStatusCode {
  "Unknown" : PathExistsStatusCode(-1),
  "DoesNotExist" : PathExistsStatusCode(0),
  "Exists" : PathExistsStatusCode(1) }

var mPathExistsStatusLowerCaseStringToCode = map[string] PathExistsStatusCode {
  "unknown" : PathExistsStatusCode(-1),
  "doesnotexist" : PathExistsStatusCode(0),
  "exists" : PathExistsStatusCode(1) }




// PathExistsStatusCode - This type is an enumeration of file path
// existence status codes. These status codes describe the state
// of a file or directory on disk. The existence of a path or
// directory is usually determined by the method os.Stat() which
// establishes whether or not a file or directory currently exists
// on disk.
//
//
//                    Path Existence
//   Method            Status Code
//    Name               Constant       Description
//  ______________________________________________________________________
//  Unknown()              -1           Path file existence has NOT been
//                                        tested and status is 'Unknown'.
//
//  DoesNotExist()          0           Path file existence HAS been tested
//                                        and path file name does NOT exist on
//                                        disk.
//
//  Exists()               +1           Path file existence HAD been tested
//                                        and path file name DOES exist on
//                                        disk.
//
// PathExistsStatusCode has been adapted to function as an enumeration
// of valid Path Existence Status Code values. Since Go does not directly
// support enumerations, the 'PathExistsStatusCode' type has been configured
// to function in a manner similar to classic enumerations found in other
// languages like C#. For additional information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type PathExistsStatusCode int

// Unknown - This file path status code signals that the existence of
// the file or directory on disk has not yet been established. That is,
// the test for existence has not yet been performed.
//
func (pathExist PathExistsStatusCode) Unknown() PathExistsStatusCode {
  return PathExistsStatusCode(-1)
}

// DoesNotExist - This Path Existence Status code signals that the test
// for existence on disk HAS been performed and that the subject file
// or directory DOES NOT EXIST on disk.
//
func (pathExist PathExistsStatusCode) DoesNotExist() PathExistsStatusCode {
  return PathExistsStatusCode(0)
}

// Exists - This Path Existence Status code signals that the test
// for existence on disk HAS been performed and that the subject file
// or directory DOES EXIST on disk.
//
func (pathExist PathExistsStatusCode) Exists() PathExistsStatusCode {
  return PathExistsStatusCode(1)
}

// StatusCodesEqual - Compares the current PathExistsStatusCode instance to another
// PathExistsStatusCode instance passed as an input parameter. If the two are
// equal in all respects, this method returns 'true'. Otherwise, this method
// returns 'false'.
//
func (pathExist PathExistsStatusCode) StatusCodesEqual( statusCode PathExistsStatusCode) bool {

  if pathExist == statusCode {
    return true
  }

  return false
}

// IsValid - If the value of the current PathExistsStatusCode instance is
// 'invalid', this method will return an error.
//
// If the PathExistsStatusCode is instance is 'valid', this method will return
// a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (pathExist PathExistsStatusCode) StatusIsValid() error {

  _, ok := mPathExistsStatusCodeToString[pathExist]

  if !ok {
    ePrefix := "PathExistsStatusCode.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current PathExistsStatusCode is INVALID! "+
      "PathExistsStatusCode Value='%v'", int(pathExist))
  }

  return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful, a new
// instance of PathExistsStatusCode is returned set to the value of the
// associated enumeration.
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
//	                       exact match. Therefore, 'exists' will NOT
//	                       match the enumeration name, 'Exists'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'exists'
//	                       will match match enumeration name 'Exists'.
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
//	t, err := PathExistsStatusCode(0).ParseString("Exists", true)
//                            OR
//	t, err := OsFilePermissionCode(0).ParseString("Exists()", true)
//                            OR
//	t, err := OsFilePermissionCode(0).ParseString("exists", false)
//
//	For all of the cases shown above,
//  t is now equal to PathExistsStatusCode(0).Exists()
//
//
func (pathExist PathExistsStatusCode) ParseString(
  valueString string, caseSensitive bool) (PathExistsStatusCode, error) {

    ePrefix := "PathExistsStatusCode.StatusCodeParseString() "

  result := PathExistsStatusCode(-1)

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


  if lenValueStr < 6 {
    return result,
      fmt.Errorf(ePrefix+
        "Error: Input parameter 'valueString' is INVALID!\n" +
        "Length Less than 6-characters valueString='%v'\n", valueString)
  }

  var ok bool

  if caseSensitive {

    result, ok = mPathExistsStatusStringToCode[valueString]

  } else {
    // Search is NOT Case Sensitive

    valueString = strings.ToLower(valueString)

    result, ok = mPathExistsStatusLowerCaseStringToCode[valueString]

  }

  if !ok {
    result = PathExistsStatusCode(-1)
    return result,
      fmt.Errorf(ePrefix +
        "Error: Invalid Permission Code!\n" +
        "valueString='%v'\n", valueString)
  }

  return result, nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'PathExistsStatusCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the PathExistsStatusCode value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= PathExistsStatusCode(0).DoesNotExist()
//	str := t.String()
//	    str is now equal to "DoesNotExist"
//
func (pathExist PathExistsStatusCode) String() string {

  label, ok := mPathExistsStatusCodeToString[pathExist]

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
func (pathExist PathExistsStatusCode) Value() PathExistsStatusCode {
  return pathExist
}


var PathExistsStatus PathExistsStatusCode