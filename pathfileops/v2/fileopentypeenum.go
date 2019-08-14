package pathfileops

import (
  "fmt"
  "os"
  "reflect"
  "strings"
)

// mFileOpenTypeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOpenType.
var mFileOpenTypeIntToString = map[int]string{}

// mFileOpenTypeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOpenType.
var mFileOpenTypeStringToInt = map[string]int{}

// mFileOpenTypeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOpenType.
// This map is used for case insensitive look ups.
var mFileOpenTypeLwrCaseStringToInt = map[string]int{}

// FileOpenType - In order to open a file, exactly one of the
// following File Open Codes MUST be specified:
//
//  FileOpenType(0).TypeReadOnly()
//  FileOpenType(0).TypeWriteOnly()
//  FileOpenType(0).TypeReadWrite()
//
// In addition, one of the three previous codes may be or'd with
// zero or more of the following File Open Modes (Type: 'FileOpenMode')
// to better control file open behavior.
//
//  FileOpenMode(0).ModeAppend()
//  FileOpenMode(0).ModeCreate()
//  FileOpenMode(0).ModeExclusive()
//  FileOpenMode(0).ModeSync()
//  FileOpenMode(0).ModeTruncate()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.
//
// FileOpenType has been adapted to function as an enumeration of valid
// File Open Type values. Since Go does not directly support enumerations,
// the 'FileOpenType' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type FileOpenType int

// None - No File Open Type specified
func (fOpenType FileOpenType) TypeNone() FileOpenType { return -1 }

// ReadOnly - File opened for 'Read Only' access
func (fOpenType FileOpenType) TypeReadOnly() FileOpenType { return FileOpenType(os.O_RDONLY) }

// WriteOnly - File opened for 'Write Only' access
func (fOpenType FileOpenType) TypeWriteOnly() FileOpenType { return FileOpenType(os.O_WRONLY) }

// ReadWrite - File opened for 'Read and Write' access
func (fOpenType FileOpenType) TypeReadWrite() FileOpenType { return FileOpenType(os.O_RDWR) }

// IsValid - If the value of the current FileOpenType is 'invalid',
// this method will return an error. If the FileOpenType is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fOpenType FileOpenType) IsValid() error {

  fOpenType.checkInitializeMaps(false)

  _, ok := mFileOpenTypeIntToString[int(fOpenType)]

  if !ok {
    ePrefix := "FileOpenType.IsValid() "
    return fmt.Errorf(ePrefix+
      "Error: Invalid FileOpenType! Current FileOpenType='%v'",
      fOpenType)
  }

  return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOpenType is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  valueString   string - A string which will be matched against the
//                         enumeration string values. If 'valueString'
//                         is equal to one of the enumeration names, this
//                         method will proceed to successful completion.
//
//                         You can prefix the string with "Type" or not.
//                         Examples: "ReadOnly" or "TypeReadOnly"
//                         Either string will produce the correct result.
//
//  caseSensitive   bool - If 'true' the search for enumeration names
//                         will be case sensitive and will require an
//                         exact match. Therefore, 'readonly' will NOT
//                         match the enumeration name, 'ReadOnly'.
//
//                         If 'false' a case insensitive search is conducted
//                         for the enumeration name. In this case, 'readonly'
//                         will match match enumeration name 'ReadOnly'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FileOpenType - Upon successful completion, this method will return a new
//	               instance of FileOpenType set to the value of the enumeration
//	               matched by the string search performed on input parameter,
//	               'valueString'.
//
//	error        - If this method completes successfully, the returned error
//	               Type is set equal to 'nil'. If an error condition is encountered,
//	               this method will return an error Type which encapsulates an
//	               appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//   t, err := FileOpenType(0).ParseString("ReadOnly", true)
//                          Or
//   t, err := FileOpenType(0).ParseString("TypeReadOnly", true)
//                          Or
//   t, err := FileOpenType(0).ParseString("TypeReadOnly()", true)
//                          Or
//   t, err := FileOpenType(0).ParseString("ReadOnly()", true)
//                          Or
//   t, err := FileOpenType(0).ParseString("readonly", false)
//
//   In of the cases shown above, t is now equal to FileOpenType(0).ReadOnly()
//
func (fOpenType FileOpenType) ParseString(
  valueString string,
  caseSensitive bool) (FileOpenType, error) {

  ePrefix := "FileOpenType.ParseString() "

  fOpenType.checkInitializeMaps(false)

  result := FileOpenType(0)

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return result,
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
  }

  var ok bool
  var idx int

  if caseSensitive {

    if !strings.HasPrefix(valueString, "Type") {
      valueString = "Type" + valueString
    }

    idx, ok = mFileOpenTypeStringToInt[valueString]

    if !ok {
      return FileOpenType(0),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileOpenType. valueString='%v' ", valueString)
    }

    result = FileOpenType(idx)

  } else {

    valueString = strings.ToLower(valueString)

    if !strings.HasPrefix(valueString, "type") {
      valueString = "type" + valueString
    }

    idx, ok = mFileOpenTypeLwrCaseStringToInt[valueString]

    if !ok {
      return FileOpenType(0),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileOpenType. valueString='%v' ", valueString)
    }

    result =
      FileOpenType(idx)
  }

  return result, nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOpenType'. This is a standard utility method
// and is not part of the valid enumerations for this type.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the FileOpenType value is invalid, this method will
//           return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//   t   := FileOpenType(0).TypeReadWrite()
//   str := t.String()
//     str is now equal to "TypeReadWrite"
//
func (fOpenType FileOpenType) String() string {

  fOpenType.checkInitializeMaps(false)

  str, ok := mFileOpenTypeIntToString[int(fOpenType)]

  if !ok {
    return ""
  }

  return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOpenType
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (fOpenType FileOpenType) Value() int {
  return int(fOpenType)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time FileOpenType.String() or
// FileOpenType.ParseString() a call is made to this method to determine if
// these maps have been initialized. If the maps and look up data have been
// properly initialized and indexed, this method returns without taking action.
//
// On the other hand, if the maps have not yet been initialized, this method will
// initialize all associated map slices.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	reInitialize     bool - If 'true', this will force initialization of
//	                        all associated maps.
//
func (fOpenType FileOpenType) checkInitializeMaps(reInitialize bool) {

  if !reInitialize &&
    mFileOpenTypeIntToString != nil &&
    len(mFileOpenTypeIntToString) > 3 &&
    mFileOpenTypeStringToInt != nil &&
    len(mFileOpenTypeStringToInt) > 3 &&
    mFileOpenTypeLwrCaseStringToInt != nil &&
    len(mFileOpenTypeLwrCaseStringToInt) > 3 {
    return
  }

  var t = FileOpenType(0).TypeReadOnly()

  mFileOpenTypeIntToString = make(map[int]string, 0)
  mFileOpenTypeStringToInt = make(map[string]int, 0)
  mFileOpenTypeLwrCaseStringToInt = make(map[string]int, 0)

  s := reflect.TypeOf(t)

  intZero := 0

  r := reflect.TypeOf(intZero)
  args := [1]reflect.Value{reflect.Zero(s)}

  for i := 0; i < s.NumMethod(); i++ {

    f := s.Method(i).Name

    if f == "String" ||
      f == "ParseString" ||
      f == "Value" ||
      f == "IsValid" ||
      f == "checkInitializeMaps" {
      continue
    }

    value := s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
    x := int(value)
    mFileOpenTypeIntToString[x] = f
    mFileOpenTypeStringToInt[f] = x
    mFileOpenTypeLwrCaseStringToInt[strings.ToLower(f)] = x
  }

}

// FOpenType - This public global variable allows
// easy access to the enumerations of the FileOpenType
// using the dot operator.
//
//  Example:
//
//     FOpenType.TypeReadOnly()
//     FOpenType.TypeWriteOnly()
//     FOpenType.TypeReadWrite()
//
var FOpenType = FileOpenType(0)
