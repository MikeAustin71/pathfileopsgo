package pathfileops

import (
  "fmt"
  "reflect"
  "strings"
)

// mFileOperationCodeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperationCode.
var mFileOperationCodeIntToString = map[int]string{}

// mFileOperationCodeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperationCode.
var mFileOperationCodeStringToInt = map[string]int{}

// mFileOperationCodeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperationCode.
// This map is used for case insensitive look ups.
var mFileOperationCodeLwrCaseStringToInt = map[string]int{}

// FileOperationCode - Integer enumeration. Signals
// the type of operation to be performed on a file.
//
//
// Usage:
//
// ----------------------------------------------------
//
// To designate a File Operation Code, use the form:
//
// 	FileOperationCode(0).MoveSourceFileToDestinationFile()
//
// To access the File Operation Code using strictly dot
// notation, use either the private variable:
//
// 	fileOpCode.MoveSourceFileToDestinationFile()
//
// OR the public global variable:
//
//  FileOpCode.MoveSourceFileToDestinationFile()
//
//
// Listing Of File Operation Codes:
//
// ----------------------------------------------------
//
// FileOperationCode(0).None()
// FileOperationCode(0).MoveSourceFileToDestinationFile()
// FileOperationCode(0).MoveSourceFileToDestinationDir()
// FileOperationCode(0).DeleteDestinationFile()
// FileOperationCode(0).DeleteSourceFile()
// FileOperationCode(0).DeleteSourceAndDestinationFiles()
// FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
// FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
// FileOperationCode(0).CopySourceToDestinationByHardLink()
// FileOperationCode(0).CopySourceToDestinationByIo()
// FileOperationCode(0).CreateSourceDir()
// FileOperationCode(0).CreateSourceDirAndFile()
// FileOperationCode(0).CreateSourceFile()
// FileOperationCode(0).CreateDestinationDir()
// FileOperationCode(0).CreateDestinationDirAndFile()
// FileOperationCode(0).CreateDestinationFile()
//
// FileOperationCode has been adapted to function as an enumeration of valid
// File Operation Code values. Since Go does not directly support enumerations,
// the 'FileOperationCode' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type FileOperationCode int

// None - No operation (NOOP) No File Operation is performed.
//
// Usage:
//
//	FileOperationCode(0).None()
//
func (fop FileOperationCode) None() FileOperationCode { return FileOperationCode(0) }

// MoveSourceFileToDestinationFile - Moves the source file to the destination
// file and then deletes the original source file.
//
// Usage:
//
//	FileOperationCode(0).MoveSourceFileToDestinationFile()
func (fop FileOperationCode) MoveSourceFileToDestinationFile() FileOperationCode {
  return FileOperationCode(1)
}

// MoveSourceFileToDestinationDir - Moves the source file to the destination
// directory and then deletes the original source file. The destination file
// name is the same as the source file name.
//
// Usage:
//
//	FileOperationCode(0).MoveSourceFileToDestinationDir()
func (fop FileOperationCode) MoveSourceFileToDestinationDir() FileOperationCode {
  return FileOperationCode(2)
}

// DeleteDestinationFile() - Deletes the Destination file if it exists
//
// Usage:
//
//	FileOperationCode(0).DeleteDestinationFile()
func (fop FileOperationCode) DeleteDestinationFile() FileOperationCode { return FileOperationCode(3) }

// DeleteSourceFile() - Deletes the Source file if it exists
//
// Usage:
//
//	FileOperationCode(0).DeleteSourceFile()
func (fop FileOperationCode) DeleteSourceFile() FileOperationCode { return FileOperationCode(4) }

// DeleteSourceAndDestinationFiles - Deletes both the Source and Destination files
// if they exist.
//
// Usage:
//
//	FileOperationCode(0).DeleteSourceAndDestinationFiles()
func (fop FileOperationCode) DeleteSourceAndDestinationFiles() FileOperationCode {
  return FileOperationCode(5)
}

// FileOperationCode(0).CopySourceToDestinationByHardLinkByIo() - Copies the Source File to the
// Destination using two copy attempts. The first copy is by Hard Link.
// If the first copy attempt fails, a second copy attempt is initiated
// by creating a new file and copying the contents by 'io.Copy'.
//
// An error is returned only if both copy attempts fail. The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()()
//
func (fop FileOperationCode) CopySourceToDestinationByHardLinkByIo() FileOperationCode {
  return FileOperationCode(6)
}

// FileOperationCode(0).CopySourceToDestinationByIoByHardLink() - Copies the Source File to the Destination
// using two copy attempts. The first copy is by 'io.Copy' which creates a new file
// and copies the contents to the new file. If the first attempt fails, a second
// copy attempt is initiated using 'copy by hard link'.
//
// An error is returned only if both copy attempts fail. The source file is
// unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//
func (fop FileOperationCode) CopySourceToDestinationByIoByHardLink() FileOperationCode {
  return FileOperationCode(7)
}

// CopySourceToDestinationByHardLink - Copies the Source File to the Destination
// using one copy mode. The only copy attempt utilizes 'Copy by Hard Link'. If
// this attempted copy fails, an error is returned.  The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByHardLink()
func (fop FileOperationCode) CopySourceToDestinationByHardLink() FileOperationCode {
  return FileOperationCode(8)
}

// CopySourceToDestinationByIo - Copies the Source File to the Destination
// using only one copy mode. The only copy attempt is initiated using 'Copy by IO'
// or 'io.Copy'.  If this fails an error is returned. The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByIo()
func (fop FileOperationCode) CopySourceToDestinationByIo() FileOperationCode {
  return FileOperationCode(9)
}

// CreateSourceDir - Creates the Source Directory
//
// Usage:
//
//	FileOperationCode(0).CreateSourceDir()
// formerly FileOperationCode(0).CreateSourceDir()
func (fop FileOperationCode) CreateSourceDir() FileOperationCode { return FileOperationCode(10) }

// CreateSourceDirAndFile - Creates the Source Directory and File
//
// Usage:
//
//	FileOperationCode(0).CreateSourceDirAndFile()
func (fop FileOperationCode) CreateSourceDirAndFile() FileOperationCode { return FileOperationCode(11) }

// CreateSourceFile - Creates the Source File
//
// Usage:
//
//	FileOperationCode(0).CreateSourceFile()
func (fop FileOperationCode) CreateSourceFile() FileOperationCode { return FileOperationCode(12) }

// CreateDestinationDir - Creates the Destination Directory
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationDir()
func (fop FileOperationCode) CreateDestinationDir() FileOperationCode {
  return FileOperationCode(13)
}

// CreateDestinationDirAndFile - Creates the Destination Directory and File
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationDirAndFile()
func (fop FileOperationCode) CreateDestinationDirAndFile() FileOperationCode {
  return FileOperationCode(14)
}

// CreateDestinationFile - Creates the Destination File
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationFile()
func (fop FileOperationCode) CreateDestinationFile() FileOperationCode {
  return FileOperationCode(15)
}

// IsValid - If the value of the current FileOperationCode is 'invalid',
// this method will return an error. If the FileOperationCode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fop FileOperationCode) IsValid() error {

  fop.checkInitializeMaps(false)

  _, ok := mFileOperationCodeIntToString[int(fop)]

  if !ok {
    ePrefix := "FileOperationCode.IsValid() "
    return fmt.Errorf(ePrefix+"Error: File Operation Code INVALID!. Unknown Value='%v' ", fop.Value())
  }

  return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOperationCode is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
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
//	FileOperationCode - Upon successful completion, this method will return a new
//	               instance of FileOperationCode set to the value of the enumeration
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
//	t, err := FileOperationCode(0).ParseString("MoveSourceFileToDestinationFile", true)
//                                  OR
//	t, err := FileOperationCode(0).ParseString("movesourcefiletodestination", false)
//
//	For all of the cases shown above,
//	t is now equal to FileOperationCode(0).MoveSourceFileToDestinationFile()
//
func (fop FileOperationCode) ParseString(
  valueString string,
  caseSensitive bool) (FileOperationCode, error) {

  ePrefix := "FileOperationCode.ParseString() "

  fop.checkInitializeMaps(false)

  result := FileOperationCode(0)
  if len(valueString) < 3 {
    return result,
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
  }

  var ok bool
  var idx int

  if caseSensitive {

    idx, ok = mFileOperationCodeStringToInt[valueString]

    if !ok {
      return FileOperationCode(0),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileOperationCode. valueString='%v' ", valueString)
    }

    result = FileOperationCode(idx)

  } else {

    idx, ok = mFileOperationCodeLwrCaseStringToInt[strings.ToLower(valueString)]

    if !ok {
      return FileOperationCode(0),
        fmt.Errorf(ePrefix+
          "'valueString' did NOT MATCH a FileOperationCode. valueString='%v' ", valueString)
    }

    result =
      FileOperationCode(idx)
  }

  return result, nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOperationCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the FileOperationCode value is invalid, this method
//           will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOperationCode(0).MoveSourceFileToDestinationFile()
//	str := t.String()
//	    str is now equal to "MoveSourceFileToDestinationFile"
//
func (fop FileOperationCode) String() string {

  fop.checkInitializeMaps(false)

  str, ok := mFileOperationCodeIntToString[int(fop)]

  if !ok {
    return ""
  }

  return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOperationCode
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fop FileOperationCode) Value() int {
  return int(fop)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time FileOperationCode.String() or
// FileOperationCode.ParseString() a call is made to this method to determine if
// these maps have been initialized. If the maps and look up data have been
// properly initialized and indexed, this method returns without taking action.
//
// On the other hand, if the maps have not yet been initialized, this method will
// initialize all associated map slices.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	reInitialize     bool - If 'true', this will force initialization of
//	                        all associated maps.
//
func (fop FileOperationCode) checkInitializeMaps(reInitialize bool) {

  if !reInitialize &&
    mFileOperationCodeIntToString != nil &&
    len(mFileOperationCodeIntToString) > 12 &&
    mFileOperationCodeStringToInt != nil &&
    len(mFileOperationCodeStringToInt) > 12 &&
    mFileOperationCodeLwrCaseStringToInt != nil &&
    len(mFileOperationCodeLwrCaseStringToInt) > 12 {
    return
  }

  var t = FileOperationCode(0).MoveSourceFileToDestinationFile()

  mFileOperationCodeIntToString = make(map[int]string, 0)
  mFileOperationCodeStringToInt = make(map[string]int, 0)
  mFileOperationCodeLwrCaseStringToInt = make(map[string]int, 0)

  s := reflect.TypeOf(t)

  r := reflect.TypeOf(0)
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
    mFileOperationCodeIntToString[x] = f
    mFileOperationCodeStringToInt[f] = x
    mFileOperationCodeLwrCaseStringToInt[strings.ToLower(f)] = x
  }

}

// fileOpCode - Internal or private global variable of type FileOperationCode.
//
// Usage:
//
//	fileOpCode.None()
//	fileOpCode.CopySourceToDestinationByHardLink()
//
var fileOpCode = FileOperationCode(0)

// FileOp - global variable of type FileOperationCode.
// Provides alternate, easier access to FileOperationCode enumeration values.
//
// Usage:
//
//	FileOpCode.None()
//	FileOpCode.CopySourceToDestinationByIo()
//
var FileOpCode = FileOperationCode(0)
