package pathfileops

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type FileOpenConfig struct {
	isInitialized bool

	fileOpenType FileOpenType

	fileOpenModes []FileOpenMode
}

// CopyIn - Receives a FileOpenConfig instance and copies all the data
// fields to the current FileOpenConfig instance. When complete, both
// the incoming and current FileOpenConfig instances will be identical.
//
func (fOpenStat *FileOpenConfig) CopyIn(fOpStat2 FileOpenConfig) {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpStat2.fileOpenModes == nil {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 0)
	}

	fOpenStat.isInitialized = fOpStat2.isInitialized
	fOpenStat.fileOpenType = fOpStat2.fileOpenType

	lenFOpStat2FOpenModes := len(fOpStat2.fileOpenModes)

	if lenFOpStat2FOpenModes == 0 {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 1)
		fOpStat2.fileOpenModes[0] = FOpenMode.None()
		lenFOpStat2FOpenModes = 1
	}

	fOpenStat.fileOpenModes = make([]FileOpenMode, lenFOpStat2FOpenModes)

	for i := 0; i < lenFOpStat2FOpenModes; i++ {
		fOpenStat.fileOpenModes[i] = fOpStat2.fileOpenModes[i]
	}

}

// CopyOut - Creates and returns a deep copy of the current
// FileOpenConfig instance.
func (fOpenStat *FileOpenConfig) CopyOut() FileOpenConfig {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	fOpStat2 := FileOpenConfig{}
	fOpStat2.isInitialized = fOpenStat.isInitialized
	fOpStat2.fileOpenType = fOpenStat.fileOpenType
	lenFOpenModes := len(fOpenStat.fileOpenModes)

	if lenFOpenModes == 0 {
		fOpenStat.fileOpenModes = append(fOpenStat.fileOpenModes, FOpenMode.None())
		lenFOpenModes = 1
	}

	fOpStat2.fileOpenModes = make([]FileOpenMode, lenFOpenModes)

	for i := 0; i < lenFOpenModes; i++ {
		fOpStat2.fileOpenModes[i] = fOpenStat.fileOpenModes[i]
	}

	return fOpStat2
}

// Empty - ReInitializes the current FileOpenConfig instance to
// empty or zero values.
//
func (fOpenStat *FileOpenConfig) Empty() {

	fOpenStat.isInitialized = false

	fOpenStat.fileOpenType = FOpenType.None()

	fOpenStat.fileOpenModes = make([]FileOpenMode, 0)

	fOpenStat.fileOpenModes = append(fOpenStat.fileOpenModes, FOpenMode.None())

}

// Equal - Returns 'true' if the incoming FileOpenConfig Type is equal in
// all respects to the current FileOpenConfig instance.
//
func (fOpenStat *FileOpenConfig) Equal(fOpStat2 FileOpenConfig) bool {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpStat2.fileOpenModes == nil {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpenStat.isInitialized != fOpStat2.isInitialized {
		return false
	}

	lenFileOpenModes := len(fOpenStat.fileOpenModes)

	if lenFileOpenModes != len(fOpStat2.fileOpenModes) {
		return false
	}

	if fOpenStat.fileOpenType != fOpStat2.fileOpenType {
		return false
	}

	for i := 0; i < lenFileOpenModes; i++ {
		isFound := false

		for j := 0; j < lenFileOpenModes; j++ {
			if fOpStat2.fileOpenModes[j] == fOpenStat.fileOpenModes[i] {
				isFound = true
			}
		}

		if !isFound {
			return false
		}
	}

	return true
}

// New - Creates and returns a fully initialized FileOpenConfig instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  fOpenType FileOpenType - The FileOpenType used to open a file.
//
//  fOpenModes ... FileOpenMode - Zero or more FileOpenMode instances which will be or'd
//                                with the input parameter 'fOpenType' in order to generate
//                                the composite 'file open' code which will be used to open
//                                a file.  If no File Open Modes will be used, the user should
//                                pass 'FileOpenMode(0).None()' or pass nothing for this
//                                parameter.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//
//  FileOpenConfig - If successful, this method will return a new, fully initialized instance
//                   of FileOpenConfig.
//
//	error          - If this method completes successfully, the returned error
//	                 Type is set equal to 'nil'. If an error condition is encountered,
//	                 this method will return an error Type which encapsulates an
//	                 appropriate error message.
//
func (fOpenStat FileOpenConfig) New(fOpenType FileOpenType, fOpenModes ...FileOpenMode) (FileOpenConfig, error) {

	ePrefix := "FileOpenConfig.New() "

	err := fOpenType.IsValid()

	if err != nil {
		return FileOpenConfig{},
			fmt.Errorf(ePrefix+"Error: Input parameter 'fOpenType' is INVALID! fOpenType='%v' ", err.Error())
	}
	resultFOpenStatus := FileOpenConfig{}

	resultFOpenStatus.fileOpenType = fOpenType

	resultFOpenStatus.fileOpenModes = make([]FileOpenMode, 0)

	if len(fOpenModes) == 0 {

		resultFOpenStatus.fileOpenModes = append(resultFOpenStatus.fileOpenModes, FOpenMode.None())

		resultFOpenStatus.isInitialized = true

		return resultFOpenStatus, nil
	}

	for idx, mode := range fOpenModes {

		err = mode.IsValid()

		if err != nil {
			return FileOpenConfig{},
				fmt.Errorf(ePrefix+
					"Error: Input parameter 'fOpenModes' contains an invalid FileOpenMode. Index='%v' ", idx)
		}

		resultFOpenStatus.fileOpenModes = append(resultFOpenStatus.fileOpenModes, mode)

	}

	resultFOpenStatus.isInitialized = true

	return resultFOpenStatus, nil
}

// GetCompositeFileOpenCode - Returns the composite 'file open' code. This code
// is generated by or'ing together the stored single FileOpenType value and zero
// or more FileOpenMode values.
//
func (fOpenStat *FileOpenConfig) GetCompositeFileOpenCode() (int, error) {

	ePrefix := "FileOpenConfig.GetCompositeFileOpenCode() "

	if !fOpenStat.isInitialized {
		return -1,
			errors.New(ePrefix + "Error: The current FileOpenConfig instance is INVALID!")
	}

	if fOpenStat.fileOpenType == FileOpenType(0).None() {
		return -1,
			errors.New(ePrefix + "Error: The stored FileOpenType == 'None'. A valid FileOpenType is required!")
	}

	err := fOpenStat.fileOpenType.IsValid()

	if err != nil {
		return -1,
			fmt.Errorf(ePrefix+
				"Error: The stored FileOpenType is INVALID! FileOpenType='%v' ",
				fOpenStat.fileOpenType.Value())
	}

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	lenFileOpenModes := len(fOpenStat.fileOpenModes)

	if lenFileOpenModes == 0 {
		return fOpenStat.fileOpenType.Value(), nil
	}

	if lenFileOpenModes == 1 &&
		fOpenStat.fileOpenModes[0] == FileOpenMode(0).None() {

		return fOpenStat.fileOpenType.Value(), nil

	}

	fileOpenVal := fOpenStat.fileOpenType.Value()

	for i := 0; i < lenFileOpenModes; i++ {
		fileOpenVal = fileOpenVal | fOpenStat.fileOpenModes[i].Value()
	}

	return fileOpenVal, nil
}

// GetFileOpenModes - Returns a array of stored FileOpenMode values
func (fOpenStat *FileOpenConfig) GetFileOpenModes() []FileOpenMode {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if len(fOpenStat.fileOpenModes) == 0 {
		fOpenStat.fileOpenModes = append(fOpenStat.fileOpenModes, FileOpenMode(0).None())
	}

	resultAry := make([]FileOpenMode, 0)

	resultAry = append(resultAry, fOpenStat.fileOpenModes...)

	return resultAry
}

// GetFileOpenType - Returns the stored FileOpenType value.
func (fOpenStat *FileOpenConfig) GetFileOpenType() FileOpenType {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	return fOpenStat.fileOpenType
}

// IsValid - If the current FileOpenConfig is valid and properly
// initialized, this method returns nil. If the current FileOpenConfig
// instance is invalid, this method returns an error.
func (fOpenStat *FileOpenConfig) IsValid() error {

	ePrefix := "FileOpenConfig.IsValid() "

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if !fOpenStat.isInitialized {
		return errors.New(ePrefix +
			"Error: The current FileOpenConfig instance has NOT been " +
			"properly initialized.")
	}

	err := fOpenStat.fileOpenType.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: The File Open Type is INVALID!. %v", err.Error())
	}

	lenFileOpenModes := len(fOpenStat.fileOpenModes)

	if fOpenStat.fileOpenType == FOpenType.None() &&
		lenFileOpenModes > 1 {
		return errors.New(ePrefix +
			"Error: Current FileOpenConfig has Type='None' and " +
			"multiple File Open Modes!")
	}

	if fOpenStat.fileOpenType == FOpenType.None() &&
		lenFileOpenModes == 1 &&
		fOpenStat.fileOpenModes[0] != FileOpenMode(0).None() {
		return errors.New(ePrefix +
			"Error: Current FileOpenConfig has Type='None' and " +
			"a valid File Open Mode")
	}

	if fOpenStat.fileOpenType != FOpenType.None() {

		for i := 0; i < lenFileOpenModes; i++ {
			if fOpenStat.fileOpenModes[i] == FileOpenMode(0).None() {
				return errors.New(ePrefix + "Error: The File Open Status has multiple File Open Modes " +
					"one of which is 'None'. Resolve this conflict.")
			}
		}

	}

	for i := 0; i < lenFileOpenModes; i++ {

		err := fOpenStat.fileOpenModes[i].IsValid()

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error: A File Open Mode is INVALID! Index='%v' "+
				"Invalid Error='%v' ", i, err.Error())
		}

	}

	return nil
}

// SetFileOpenType - Receives an input parameter 'fOpenType' which is
// used to set the internal stored FileOpenType for the current FileOpenConfig
// instance.
func (fOpenStat *FileOpenConfig) SetFileOpenType(fOpenType FileOpenType) error {

	ePrefix := "FileOpenConfig.SetFileOpenType() "

	err := fOpenType.IsValid()

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if err != nil {
		return fmt.Errorf(ePrefix+"Input parameter 'fOpenType' is INVALID! fOpenType='%v' ",
			fOpenType.Value())
	}

	if fOpenType == FileOpenType(0).None() {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	fOpenStat.fileOpenType = fOpenType

	fOpenStat.isInitialized = true

	return nil
}

// SetFileOpenModes - Receives a series of FileOpenMode instances and
// replaces the internal stored FileOpenMode values for this FileOpenConfig instance.
//
// To clear the current internal FileOpenMode values, pass nothing as an input parameter
// or pass the value FileOpenMode(0).None().
//
func (fOpenStat *FileOpenConfig) SetFileOpenModes(fOpenModes ...FileOpenMode) {

	if fOpenStat.fileOpenModes == nil {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
	}

	if len(fOpenModes) == 0 {
		fOpenStat.fileOpenModes = make([]FileOpenMode, 0)
		return
	}

	fOpenStat.fileOpenModes = make([]FileOpenMode, 0)

	fOpenStat.fileOpenModes = append(fOpenStat.fileOpenModes, fOpenModes...)

	fOpenStat.isInitialized = true

	return
}

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
//  FileOpenType(0).ReadOnly()
//  FileOpenType(0).WriteOnly()
//  FileOpenType(0).ReadWrite()
//
// In addition, one of the three previous codes may be or'd with
// zero or more of the following File Open Modes (Type: 'FileOpenMode')
// to better control file open behavior.
//
//  FileOpenMode(0).Append()
//  FileOpenMode(0).Create()
//  FileOpenMode(0).Exclusive()
//  FileOpenMode(0).Sync()
//  FileOpenMode(0).Truncate()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.
//
type FileOpenType int

// None - No File Open Type specified
func (fOpenType FileOpenType) None() FileOpenType { return -1 }

// ReadOnly - File opened for 'Read Only' access
func (fOpenType FileOpenType) ReadOnly() FileOpenType { return FileOpenType(os.O_RDONLY) }

// WriteOnly - File opened for 'Write Only' access
func (fOpenType FileOpenType) WriteOnly() FileOpenType { return FileOpenType(os.O_WRONLY) }

// ReadWrite - File opened for 'Read and Write' access
func (fOpenType FileOpenType) ReadWrite() FileOpenType { return FileOpenType(os.O_RDWR) }

// IsValid - If the value of the current FileOpenType is 'invalid',
// this method will return an error. If the FileOpenType is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fOpenType FileOpenType) IsValid() error {

	var err error

	switch fOpenType {

	case FileOpenType(0).None():
		err = nil
	case FileOpenType(0).ReadOnly():
		err = nil
	case FileOpenType(0).WriteOnly():
		err = nil
	case FileOpenType(0).ReadWrite():
		err = nil
	default:
		ePrefix := "FileOpenType.IsValid() "
		err = fmt.Errorf(ePrefix+
			"Error: Ivalid FileOpenType! Current FileOpenType='%v'",
			fOpenType)
	}

	return err
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
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'readonly' will NOT
//	                       match the enumeration name, 'ReadOnly'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'readonly'
//	                       will match match enumeration name 'ReadOnly'.
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
//	t, err := FileOpenType(0).ParseString("ReadOnly")
//
//	    t is now equal to FileOpenType(0).ReadOnly()
//
func (fOpenType FileOpenType) ParseString(
	valueString string,
	caseSensitive bool) (FileOpenType, error) {

	ePrefix := "FileOpenType.ParseString() "

	fOpenType.checkInitializeMaps(false)

	result := FileOpenType(0)
	if len(valueString) < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mFileOpenTypeStringToInt[valueString]

		if !ok {
			return FileOpenType(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenType. valueString='%v' ", valueString)
		}

		result = FileOpenType(idx)

	} else {

		idx, ok = mFileOpenTypeLwrCaseStringToInt[strings.ToLower(valueString)]

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
// Usage
//
//	t:= FileOpenType(0).ReadWrite()
//	str := t.String()
//	    str is now equal to "ReadWrite"
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

	var t = FileOpenType(0).ReadOnly()

	mFileOpenTypeIntToString = make(map[int]string, 0)
	mFileOpenTypeStringToInt = make(map[string]int, 0)
	mFileOpenTypeLwrCaseStringToInt = make(map[string]int, 0)

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
		mFileOpenTypeIntToString[x] = f
		mFileOpenTypeStringToInt[f] = x
		mFileOpenTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

}

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
//  FileOpenMode(0).Append()
//  FileOpenMode(0).Create()
//  FileOpenMode(0).Exclusive()
//  FileOpenMode(0).Sync()
//  FileOpenMode(0).Truncate()
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.
//
//  FileOpenType(0).ReadOnly()
//  FileOpenType(0).WriteOnly()
//  FileOpenType(0).ReadWrite()
//
//
//  Reference CONSTANTS: https://golang.org/pkg/os/
//
// This type serves a wrapper for os package constants.

type FileOpenMode int

// None - No File Open Mode is active
func (fOpenMode FileOpenMode) None() FileOpenMode { return FileOpenMode(-1) }

// Append - append data to the file when writing.
func (fOpenMode FileOpenMode) Append() FileOpenMode { return FileOpenMode(os.O_APPEND) }

// Create - create a new file if none exists.
func (fOpenMode FileOpenMode) Create() FileOpenMode { return FileOpenMode(os.O_CREATE) }

// Exclusive - used with FileOpenControlMode(0).Create(), file must not exist.
func (fOpenMode FileOpenMode) Exclusive() FileOpenMode { return FileOpenMode(os.O_EXCL) }

// Sync - open for synchronous I/O.
func (fOpenMode FileOpenMode) Sync() FileOpenMode { return FileOpenMode(os.O_SYNC) }

// Truncate - if possible, truncate file when opened.
func (fOpenMode FileOpenMode) Truncate() FileOpenMode { return FileOpenMode(os.O_TRUNC) }

// IsValid - If the value of the current FileOpenMode is 'invalid',
// this method will return an error. If the FileOpenMode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fOpenMode FileOpenMode) IsValid() error {

	var err error

	switch fOpenMode {

	case FileOpenMode(0).None():
		err = nil
	case FileOpenMode(0).Append():
		err = nil
	case FileOpenMode(0).Create():
		err = nil
	case FileOpenMode(0).Exclusive():
		err = nil
	case FileOpenMode(0).Sync():
		err = nil
	case FileOpenMode(0).Truncate():
		err = nil
	default:
		ePrefix := "FileOpenMode.IsValid() "
		err = fmt.Errorf(ePrefix+
			"Error: Ivalid FileOpenMode! Current FileOpenMode='%v'",
			fOpenMode)
	}

	return err
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
//	t, err := FileOpenMode(0).ParseString("Append")
//
//	    t is now equal to FileOpenMode(0).Append()
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
//	t:= FileOpenMode(0).Append()
//	str := t.String()
//	    str is now equal to 'Append'
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
//     FOpenType.ReadOnly()
//     FOpenType.WriteOnly()
//     FOpenType.ReadWrite()
//
var FOpenType = FileOpenType(0)

// FOpenMode - This public global variable allows
// easy access to the enumerations of the FileOpenMode
// using the dot operator.
//
//  Example:
//
//    FileOpenMode(0).Append()
//    FileOpenMode(0).Create()
//    FileOpenMode(0).Exclusive()
//
var FOpenMode = FileOpenMode(0)
