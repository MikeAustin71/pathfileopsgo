package pathfileops

import (
	"fmt"
	"reflect"
	"strings"
)

// mFileSelectCriterionTypeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperationCode.
var mFileSelectCriterionTypeIntToString = map[int]string{}

// mFileSelectCriterionTypeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperationCode.
var mFileSelectCriterionTypeStringToInt = map[string]int{}

// mFileSelectCriterionTypeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperationCode.
// This map is used for case insensitive look ups.
var mFileSelectCriterionTypeLwrCaseStringToInt = map[string]int{}

// FileSelectCriterionMode - An enumeration which serve as parameters for file selections.
// File Selection criteria can either be And'ed or Or'ed together. The FileSelectionCriteriaMode
// determines which operation will be applied to file selection criteria.
//
type FileSelectCriterionMode int

// None - Signals that no selection is present. Same as NOOP or No Selection Criterion
func (fSel FileSelectCriterionMode) None() FileSelectCriterionMode { return FileSelectCriterionMode(0) }

// ANDSelect - File Selection Criterion are And'ed
// together. If there are three file selection criterion then
// all three must be satisfied before a file is selected.
//
func (fSel FileSelectCriterionMode) ANDSelect() FileSelectCriterionMode {
	return FileSelectCriterionMode(1)
}

// ORSelect() - File Selection Criterion are Or'd together.
// If there are three file selection criterion then satisfying any
// one of the three criterion will cause the file to be selected.
// FileSelectMode.ORSelect()
func (fSel FileSelectCriterionMode) ORSelect() FileSelectCriterionMode {
	return FileSelectCriterionMode(2)
}

// IsValid - If the value of the current FileSelectCriterionMode is 'invalid',
// this method will return an error. If the FileSelectCriterionMode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fSel FileSelectCriterionMode) IsValid() error {

	fSel.checkInitializeMaps(false)

	_, ok := mFileSelectCriterionTypeIntToString[int(fSel)]

	if !ok {
		ePrefix := "FileSelectCriterionMode.IsValid() "
		return fmt.Errorf(ePrefix+"Error: This FileSelectCriterionMode is INVALID! Unknown Code='%v' ", int(fSel))
	}

	return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileSelectCriterionMode'. This is a standard utility method
// and is not part of the valid enumerations for this type.
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

	fSel.checkInitializeMaps(false)

	str, ok := mFileSelectCriterionTypeIntToString[int(fSel)]

	if !ok {
		return ""
	}

	return str
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileSelectCriterionMode is returned set to the value of the
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
//	t, err := FileSelectCriterionMode(0).ParseString("MoveSourceFileToDestination")
//
//	    t is now equal to FileSelectCriterionMode(0).MoveSourceFileToDestination()
//
func (fSel FileSelectCriterionMode) ParseString(
	valueString string,
	caseSensitive bool) (FileSelectCriterionMode, error) {

	ePrefix := "FileSelectCriterionMode.ParseString() "

	fSel.checkInitializeMaps(false)

	result := FileSelectCriterionMode(0)
	if len(valueString) < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mFileSelectCriterionTypeStringToInt[valueString]

		if !ok {
			return FileSelectCriterionMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileSelectCriterionMode. valueString='%v' ", valueString)
		}

		result = FileSelectCriterionMode(idx)

	} else {

		idx, ok = mFileSelectCriterionTypeLwrCaseStringToInt[strings.ToLower(valueString)]

		if !ok {
			return FileSelectCriterionMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileSelectCriterionMode. valueString='%v' ", valueString)
		}

		result =
			FileSelectCriterionMode(idx)
	}

	return result, nil
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileSelectCriterionMode
// instance.
//
func (fSel FileSelectCriterionMode) Value() int {
	return int(fSel)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time FileSelectCriterionMode.String() or
// FileSelectCriterionMode.ParseString() a call is made to this method to determine if
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
func (fSel FileSelectCriterionMode) checkInitializeMaps(reInitialize bool) {

	if !reInitialize &&
		mFileSelectCriterionTypeIntToString != nil &&
		len(mFileSelectCriterionTypeIntToString) > 2 &&
		mFileSelectCriterionTypeStringToInt != nil &&
		len(mFileSelectCriterionTypeStringToInt) > 2 &&
		mFileSelectCriterionTypeLwrCaseStringToInt != nil &&
		len(mFileSelectCriterionTypeLwrCaseStringToInt) > 2 {
		return
	}

	var t = FileSelectCriterionMode(0).None()

	mFileSelectCriterionTypeIntToString = make(map[int]string, 0)
	mFileSelectCriterionTypeStringToInt = make(map[string]int, 0)
	mFileSelectCriterionTypeLwrCaseStringToInt = make(map[string]int, 0)

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(int(0))
	args := [1]reflect.Value{reflect.Zero(s)}

	for i := 0; i < s.NumMethod(); i++ {

		f := s.Method(i).Name

		if f == "String" ||
			f == "ParseString" ||
			f == "Value" ||
			f == "checkInitializeMaps" ||
			f == "IsValid" {
			continue
		}

		value := s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
		x := int(value)
		mFileSelectCriterionTypeIntToString[x] = f
		mFileSelectCriterionTypeStringToInt[f] = x
		mFileSelectCriterionTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

}

var fileSelectMode FileSelectCriterionMode

var FileSelectMode FileSelectCriterionMode
