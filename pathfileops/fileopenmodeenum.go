package pathfileops

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// mFileOpenModeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOpenMode.
var mFileOpenModeIntToString = map[int]string{}

// mFileOpenModeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOpenMode.
var mFileOpenModeStringToInt = map[string]int{}

// mFileOpenModeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOpenMode.
// This map is used for case insensitive look ups.
var mFileOpenModeLwrCaseStringToInt = map[string]int{}

// FileOpenMode - To further control the file open operation, one
// or more FileOpenMode values may be or'd with a FileOpenType
// code in order to control behavior.
//
// In addition, one of the three  codes may be or'd with
// zero or more of the following File Open Modes (Type: 'FileOpenMode')
// to better control file open behavior.
//
// FileOpenMode has been adapted to function as an enumeration of valid
// File Open Mode values. Since Go does not directly support enumerations,
// the 'FileOpenMode' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
//
// These FileOpenMode methods used as enumerators for os mode constants:
//
//  FileOpenMode(0).ModeNone()
//  FileOpenMode(0).ModeAppend()
//  FileOpenMode(0).ModeTypeCreate()
//  FileOpenMode(0).ModeExclusive()
//  FileOpenMode(0).ModeSync()
//  FileOpenMode(0).ModeTruncate()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// The FileOpenType type is used in conjunction with FileOpenMode to specify
// file permissions. Reference the 'FileOpenType' in this 'pathfileops' package.
// The methods used to specify File Open Types are listed as follows:
//
//  FileOpenType(0).TypeReadOnly()
//  FileOpenType(0).TypeWriteOnly()
//  FileOpenType(0).TypeReadWrite()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
type FileOpenMode int

// None - No File Open Mode is active
func (fOpenMode FileOpenMode) ModeNone() FileOpenMode { return FileOpenMode(-1) }

// Append - append data to the file when writing.
func (fOpenMode FileOpenMode) ModeAppend() FileOpenMode { return FileOpenMode(os.O_APPEND) }

// Create - create a new file if none exists.
func (fOpenMode FileOpenMode) ModeCreate() FileOpenMode { return FileOpenMode(os.O_CREATE) }

// Exclusive - used with FileOpenControlMode(0).Create(), file must not exist.
func (fOpenMode FileOpenMode) ModeExclusive() FileOpenMode { return FileOpenMode(os.O_EXCL) }

// Sync - open for synchronous I/O.
func (fOpenMode FileOpenMode) ModeSync() FileOpenMode { return FileOpenMode(os.O_SYNC) }

// Truncate - if possible, truncate file when opened.
func (fOpenMode FileOpenMode) ModeTruncate() FileOpenMode { return FileOpenMode(os.O_TRUNC) }

// IsValid - If the value of the current FileOpenMode is 'invalid',
// this method will return an error. If the FileOpenMode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fOpenMode FileOpenMode) IsValid() error {

	fOpenMode.checkInitializeMaps(false)

	_, ok := mFileOpenModeIntToString[int(fOpenMode)]

	if !ok {
		ePrefix := "FileOpenMode.IsValid() "
		return fmt.Errorf(ePrefix+
			"Error: Ivalid FileOpenMode! Current FileOpenMode='%v'",
			fOpenMode)
	}

	return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOpenMode is returned set to the value of the
// associated enumeration.
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
//	                       exact match. Therefore, 'append' will NOT
//	                       match the enumeration name, 'Append'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'append'
//	                       will match match enumeration name 'Append'.
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
//   t, err := FileOpenMode(0).ParseString("Append")
//
//                        OR
//
//   t, err := FileOpenMode(0).ParseString("ModeAppend")
//
//
//   In either case t is now equal to FileOpenMode(0).Append()
//
func (fOpenMode FileOpenMode) ParseString(
	valueString string,
	caseSensitive bool) (FileOpenMode, error) {

	ePrefix := "FileOpenMode.ParseString() "

	fOpenMode.checkInitializeMaps(false)

	result := FileOpenMode(0)

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

		if !strings.HasPrefix(valueString, "Mode") {
			valueString = "Mode" + valueString
		}

		idx, ok = mFileOpenModeStringToInt[valueString]

		if !ok {
			return FileOpenMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenMode. valueString='%v' ", valueString)
		}

		result = FileOpenMode(idx)

	} else {

		valueString = strings.ToLower(valueString)

		if !strings.HasPrefix(valueString, "mode") {
			valueString = "mode" + valueString
		}

		idx, ok = mFileOpenModeLwrCaseStringToInt[valueString]

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

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOpenMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOpenMode(0).ModeAppend()
//	str := t.String()
//	    str is now equal to 'ModeAppend'
//
func (fOpenMode FileOpenMode) String() string {

	fOpenMode.checkInitializeMaps(false)

	str, ok := mFileOpenModeIntToString[int(fOpenMode)]

	if !ok {
		return "Invalid File Open Mode!"
	}

	return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOpenMode
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
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

	var t = FOpenMode.ModeAppend()

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
			f == "IsValid" ||
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

// FOpenMode - This public global variable allows
// easy access to the enumerations of the FileOpenMode
// using the dot operator.
//
//  Example:
//
//    FileOpenMode(0).ModeAppend()
//    FileOpenMode(0).ModeCreate()
//    FileOpenMode(0).ModeExclusive()
//
var FOpenMode = FileOpenMode(0)
