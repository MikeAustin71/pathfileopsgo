package pathfileops

import (
  "fmt"
  "strings"
)


// mFileSelectCriterionTypeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperationCode.
var mFileSelectCriterionTypeIntToString = map[FileSelectCriterionMode]string{
  FileSelectCriterionMode(0) : "None",
  FileSelectCriterionMode(1) : "ANDSelect",
  FileSelectCriterionMode(2) : "ORSelect" }

// mFileSelectCriterionTypeStringToCode - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperationCode.
var mFileSelectCriterionTypeStringToCode = map[string]FileSelectCriterionMode{
  "None"      : FileSelectCriterionMode(0),
  "ANDSelect" : FileSelectCriterionMode(1),
  "ORSelect"  : FileSelectCriterionMode(2) }

// mFileSelectCriterionTypeLwrCaseStringToCode - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperationCode.
// This map is used for case insensitive look ups.
var mFileSelectCriterionTypeLwrCaseStringToCode = map[string]FileSelectCriterionMode{
  "none"      : FileSelectCriterionMode(0),
  "andselect" : FileSelectCriterionMode(1),
  "orselect"  : FileSelectCriterionMode(2) }

// FileSelectCriterionMode - An enumeration which serves as parameters for file selection methods.
// File Selection criteria can either be And'ed or Or'ed together. The FileSelectionCriteriaMode
// determines which operation will be applied to file selection criteria.
//
//
//                    File Select
//   Method          Criterion Mode
//    Name               Constant       Description
//  ______________________________________________________________________
//  None()                  0           Signals that no selection is present.
//                                      Same as NOOP or No Selection Criterion
//
//  ANDSelect()             1           File Selection Criterion are And'ed together.
//                                      If there are three file selection criterion then
//                                      all three must be satisfied before a file is
//                                      selected.
//
//  ORSelect()              2           File Selection Criterion are Or'd together.
//                                      If there are three file selection criterion
//                                      then satisfying any one of the three criterion
//                                      will cause the file to be selected.
//
//
// 'FileSelectCriterionMode' has been adapted to function as an enumeration
// describing the application of criteria used to select files. Since Go does not
// directly support enumerations, the 'FileSelectCriterionMode' type has been
// configured to function in a manner similar to classic enumerations found
// in other languages like C#. For additional information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type FileSelectCriterionMode int

// None - Signals that no selection is present. Same as NOOP or No Selection Criterion
//
// This method is part of the standard enumeration.
//
func (fSel FileSelectCriterionMode) None() FileSelectCriterionMode { return FileSelectCriterionMode(0) }

// ANDSelect - File Selection Criterion are And'ed
// together. If there are three file selection criterion then
// all three must be satisfied before a file is selected.
//
// This method is part of the standard enumeration.
//
func (fSel FileSelectCriterionMode) ANDSelect() FileSelectCriterionMode {
  return FileSelectCriterionMode(1)
}

// ORSelect() - File Selection Criterion are Or'd together.
// If there are three file selection criterion then satisfying any
// one of the three criterion will cause the file to be selected.
// FileSelectMode.ORSelect()
//
// This method is part of the standard enumeration.
//
func (fSel FileSelectCriterionMode) ORSelect() FileSelectCriterionMode {
  return FileSelectCriterionMode(2)
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileSelectCriterionMode is returned set to the value of the
// associated enumeration.
//
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'movesourcefiletodestination' will NOT
//	                       match the enumeration name, 'MoveSourceFileToDestinationFile'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'movesourcefiletodestination'
//	                       will match match enumeration name 'MoveSourceFileToDestinationFile'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	FileSelectCriterionMode - Upon successful completion, this method will return a new
//	               instance of FileSelectCriterionMode set to the value of the enumeration
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
// Usage
//
//  t, err := FileSelectCriterionMode(0).ParseString("MoveSourceFileToDestinationFile", true)
//                                       Or
//  t, err := FileSelectCriterionMode(0).ParseString("movesourcefiletodestination", false)
//
//  For all of the cases shown above,
//	    t is now equal to FileSelectCriterionMode(0).MoveSourceFileToDestinationFile()
//
func (fSel FileSelectCriterionMode) ParseString(
  valueString string,
  caseSensitive bool) (FileSelectCriterionMode, error) {

  ePrefix := "FileSelectCriterionMode.ParseString() "

  result := FileSelectCriterionMode(-1)

  if len(valueString) < 3 {
    return result,
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
  }

  var ok bool

  if caseSensitive {

    result, ok = mFileSelectCriterionTypeStringToCode[valueString]

    if !ok {
      return FileSelectCriterionMode(-1),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileSelectCriterionMode. valueString='%v' ", valueString)
    }

  } else {

    result, ok = mFileSelectCriterionTypeLwrCaseStringToCode[strings.ToLower(valueString)]

    if !ok {
      return FileSelectCriterionMode(-1),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileSelectCriterionMode. valueString='%v' ", valueString)
    }
  }

  return result, nil
}


// StatusIsValid - If the value of the current FileSelectCriterionMode is 'invalid',
// this method will return an error. If the FileSelectCriterionMode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fSel FileSelectCriterionMode) StatusIsValid() error {

  _, ok := mFileSelectCriterionTypeIntToString[fSel]

  if !ok {
    ePrefix := "FileSelectCriterionMode.StatusIsValid() "
    return fmt.Errorf(ePrefix+"Error: This FileSelectCriterionMode is INVALID! Unknown Code='%v' ", int(fSel))
  }

  return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileSelectCriterionMode'. This is a standard utility
// method and is not part of the valid enumerations for this type.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the FileSelectCriterionMode value is invalid,
//           this method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileSelectCriterionMode(0).ORSelect()
//	str := t.String()
//	    str is now equal to 'ORSelect'
//
func (fSel FileSelectCriterionMode) String() string {

  str, ok := mFileSelectCriterionTypeIntToString[fSel]

  if !ok {
    return ""
  }

  return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileSelectCriterionMode
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fSel FileSelectCriterionMode) Value() FileSelectCriterionMode {
  return fSel
}

// PathValidStatus - public global variable of
// type FileSelectCriterionMode.
//
// This variable serves as an easier, short hand
// technique for accessing FileSelectCriterionMode
// values.
//
// Usage:
//  FileSelectMode.None()
//  FileSelectMode.ANDSelect()
//  FileSelectMode.ORSelect()
//
var FileSelectMode FileSelectCriterionMode
