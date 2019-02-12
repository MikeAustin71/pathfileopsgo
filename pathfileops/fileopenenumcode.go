package pathfileops

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// FileOpenCode - In order to open a file, exactly one of the
// File Open Codes MUST be specified:
//
//  FileOpenCode(0).ReadOnly()
//  FileOpenCode(0).WriteOnly()
//  FileOpenCode(0).ReadWrite()
//
// In addition, one of the four previous codes may be or'd with
// one or more of the following codes to further control file
// open behavior.
//
//  FileOpenCode(0).Append()
//  FileOpenCode(0).Create()
//  FileOpenCode(0).Exclusive()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.
//
type FileOpenType int

func (fOpenType FileOpenType) None() FileOpenType { return -1 }

func (fOpenType FileOpenType) ReadOnly() FileOpenType { return FileOpenType(os.O_RDONLY) }

func (fOpenType FileOpenType) WriteOnly() FileOpenType { return FileOpenType(os.O_WRONLY) }

func (fOpenType FileOpenType) ReadWrite() FileOpenType { return FileOpenType(os.O_RDWR) }

// mFileOpenModeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperationCode.
var mFileOpenModeIntToString = map[int]string{}

// mFileOpenModeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperationCode.
var mFileOpenModeStringToInt = map[string]int{}

// mFileOpenModeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperationCode.
// This map is used for case insensitive look ups.
var mFileOpenModeLwrCaseStringToInt = map[string]int{}

// FileOpenMode - To further control the file open operation, one
// or more FileOpenMode values may be or'd with a FileOpenType
// code in order to control behavior.
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.
//
type FileOpenMode int

func (fOpenMode FileOpenMode) None() FileOpenMode { return FileOpenMode(-1) }

// Append - append data to the file when writing.
func (fOpenMode FileOpenMode) Append() FileOpenMode { return FileOpenMode(os.O_APPEND) }

// Create - create a new file if none exists.
func (fOpenMode FileOpenMode) Create() FileOpenMode { return FileOpenMode(os.O_CREATE) }

// Exclusive - used with FileOpenControlMode(0).Create(), file must not exist.
func (fOpenMode FileOpenMode) Exclusive() FileOpenMode { return FileOpenMode(os.O_EXCL) }

// Sync - open for synchronous I/O.
func (fOpenCd FileOpenMode) Sync() int { return os.O_SYNC }

// Truncate - if possible, truncate file when opened.
func (fOpenCd FileOpenMode) Truncate() int { return os.O_TRUNC }

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOpenMode'. This is a standard utility method
// and is not part of the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOpenMode(0).MoveSourceFileToDestination()
//	str := t.String()
//	    str is now equal to 'MoveSourceFileToDestination'
//
func (fOpenMode FileOpenMode) String() string {

	fOpenMode.checkInitializeMaps(false)

	return mFileOpenModeIntToString[int(fOpenMode)]
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOpenMode is returned set to the value of the
// associated enumeration.
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
//	                       match the enumeration name, 'MoveSourceFileToDestination'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'movesourcefiletodestination'
//	                       will match match enumeration name 'MoveSourceFileToDestination'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	FileOpenMode - Upon successful completion, this method will return a new
//	               instance of FileOpenMode set to the value of the enumeration
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
//	t, err := FileOpenMode(0).ParseString("MoveSourceFileToDestination")
//
//	    t is now equal to FileOpenMode(0).MoveSourceFileToDestination()
//
func (fOpenMode FileOpenMode) ParseString(
	valueString string,
	caseSensitive bool) (FileOpenMode, error) {

	ePrefix := "FileOpenMode.ParseString() "

	fOpenMode.checkInitializeMaps(false)

	result := FileOpenMode(0)
	if len(valueString) < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mFileOpenModeStringToInt[valueString]

		if !ok {
			return FileOpenMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenMode. valueString='%v' ", valueString)
		}

		result = FileOpenMode(idx)

	} else {

		idx, ok = mFileOpenModeLwrCaseStringToInt[strings.ToLower(valueString)]

		if !ok {
			return FileOpenMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenMode. valueString='%v' ", valueString)
		}

		result =
			FileOpenMode(idx)
	}

	return result, nil
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOpenMode
// instance.
//
func (fOpenMode FileOpenMode) Value() int {
	return int(fOpenMode)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time FileOpenMode.String() or
// FileOpenMode.ParseString() a call is made to this method to determine if
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
func (fOpenMode FileOpenMode) checkInitializeMaps(reInitialize bool) {

	if !reInitialize &&
		mFileOpenModeIntToString != nil &&
		len(mFileOpenModeIntToString) > 5 &&
		mFileOpenModeStringToInt != nil &&
		len(mFileOpenModeStringToInt) > 5 &&
		mFileOpenModeLwrCaseStringToInt != nil &&
		len(mFileOpenModeLwrCaseStringToInt) > 5 {
		return
	}

	var t = FileOpenMode(0).Append()

	mFileOpenModeIntToString = make(map[int]string, 0)
	mFileOpenModeStringToInt = make(map[string]int, 0)
	mFileOpenModeLwrCaseStringToInt = make(map[string]int, 0)

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(int(0))
	args := [1]reflect.Value{reflect.Zero(s)}

	for i := 0; i < s.NumMethod(); i++ {

		f := s.Method(i).Name

		if f == "String" ||
			f == "ParseString" ||
			f == "Value" ||
			f == "checkInitializeMaps" {
			continue
		}

		value := s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
		x := int(value)
		mFileOpenModeIntToString[x] = f
		mFileOpenModeStringToInt[f] = x
		mFileOpenModeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

}

var FOpenType = FileOpenType(0)

var FOpenMode = FileOpenMode(0)
