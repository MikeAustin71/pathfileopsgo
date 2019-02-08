package pathfileops

import (
	"fmt"
	"reflect"
	"strings"
)

// mFileOperationTypeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperation.
var mFileOperationTypeIntToString = map[int]string{}

// mFileOperationTypeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperation.
var mFileOperationTypeStringToInt = map[string]int{}

// mFileOperationTypeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperation.
// This map is used for case insensitive look ups.
var mFileOperationTypeLwrCaseStringToInt = map[string]int{}

// FileOperation - Integer enumeration. Signals
// the type of operation to be performed on a file.
type FileOperation int

// None - No operation (NOOP) No File Operation is performed.
//
// Usage:
//
//	FileOperation(0).None()
//
func (fop FileOperation) None() FileOperation { return FileOperation(0) }

// MoveSourceFileToDestination - Moves the source file to the destination
// file and then deletes the original source file.
//
// Usage:
//
//	FileOperation(0).MoveSourceFileToDestination()
func (fop FileOperation) MoveSourceFileToDestination() FileOperation { return FileOperation(1) }

// DeleteDestinationFile() - Deletes the Destination file if it exists
//
// Usage:
//
//	FileOperation(0).DeleteDestinationFile()
func (fop FileOperation) DeleteDestinationFile() FileOperation { return FileOperation(2) }

// DeleteSourceFile() - Deletes the Source file if it exists
//
// Usage:
//
//	FileOperation(0).DeleteSourceFile()
func (fop FileOperation) DeleteSourceFile() FileOperation { return FileOperation(3) }

// DeleteSourceAndDestinationFiles - Deletes both the Source and Destination files
// if they exist.
//
// Usage:
//
//	FileOperation(0).DeleteSourceAndDestinationFiles()
func (fop FileOperation) DeleteSourceAndDestinationFiles() FileOperation { return FileOperation(4) }

// FileOperation(0).CopySourceToDestinationByHardLinkByIo() - Copies the Source File to the
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
//	FileOperation(0).FileOperation(0).CopySourceToDestinationByHardLinkByIo()()
//
func (fop FileOperation) CopySourceToDestinationByHardLinkByIo() FileOperation {
	return FileOperation(5)
}

// FileOperation(0).CopySourceToDestinationByIoByHardLink() - Copies the Source File to the Destination
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
//	FileOperation(0).CopySourceToDestinationByIoByHardLink()
//
func (fop FileOperation) CopySourceToDestinationByIoByHardLink() FileOperation {
	return FileOperation(6)
}

// CopySourceToDestinationByHardLink - Copies the Source File to the Destination
// using one copy mode. The only copy attempt utilizes 'Copy by Hard Link'. If
// this attempted copy fails, an error is returned.  The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperation(0).CopySourceToDestinationByHardLink()
func (fop FileOperation) CopySourceToDestinationByHardLink() FileOperation { return FileOperation(7) }

// CopySourceToDestinationByIo - Copies the Source File to the Destination
// using only one copy mode. The only copy attempt is initiated using 'Copy by IO'
// or 'io.Copy'.  If this fails an error is returned. The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperation(0).CopySourceToDestinationByIo()
func (fop FileOperation) CopySourceToDestinationByIo() FileOperation { return FileOperation(8) }

// CreateSourceDir - Creates the Source Directory
//
// Usage:
//
//	FileOperation(0).CreateSourceDir()
// formerly FileOperation(0).CreateSourceDir()
func (fop FileOperation) CreateSourceDir() FileOperation { return FileOperation(9) }

// CreateSourceDirAndFile - Creates the Source Directory and File
//
// Usage:
//
//	FileOperation(0).CreateSourceDirAndFile()
func (fop FileOperation) CreateSourceDirAndFile() FileOperation { return FileOperation(10) }

// CreateSourceFile - Creates the Source File
//
// Usage:
//
//	FileOperation(0).CreateSourceFile()
func (fop FileOperation) CreateSourceFile() FileOperation { return FileOperation(11) }

// CreateDestinationDir - Creates the Destination Directory
//
// Usage:
//
//	FileOperation(0).CreateDestinationDir()
func (fop FileOperation) CreateDestinationDir() FileOperation { return FileOperation(12) }

// CreateDestinationDirAndFile - Creates the Destination Directory and File
//
// Usage:
//
//	FileOperation(0).CreateDestinationDirAndFile()
func (fop FileOperation) CreateDestinationDirAndFile() FileOperation { return FileOperation(13) }

// CreateDestinationFile - Creates the Destination File
//
// Usage:
//
//	FileOperation(0).CreateDestinationFile()
// formerly FileOperation(0).CreateDestinationFile()
func (fop FileOperation) CreateDestinationFile() FileOperation { return FileOperation(14) }

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOperation'. This is a standard utility method
// and is not part of the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOperation(0).MoveSourceFileToDestination()
//	str := t.String()
//	    str is now equal to 'MoveSourceFileToDestination'
//
func (fop FileOperation) String() string {

	fop.checkInitializeMaps(false)

	return mFileOperationTypeIntToString[int(fop)]
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOperation is returned set to the value of the
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
//	FileOperation - Upon successful completion, this method will return a new
//	               instance of FileOperation set to the value of the enumeration
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
//	t, err := FileOperation(0).ParseString("MoveSourceFileToDestination")
//
//	    t is now equal to FileOperation(0).MoveSourceFileToDestination()
//
func (fop FileOperation) ParseString(
	valueString string,
	caseSensitive bool) (FileOperation, error) {

	ePrefix := "FileOperation.ParseString() "

	fop.checkInitializeMaps(false)

	result := FileOperation(0)
	if len(valueString) < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mFileOperationTypeStringToInt[valueString]

		if !ok {
			return FileOperation(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOperation. valueString='%v' ", valueString)
		}

		if !ok {
			return FileOperation(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOperation. valueString='%v' ", valueString)
		}

		result = FileOperation(idx)

	} else {

		idx, ok = mFileOperationTypeLwrCaseStringToInt[strings.ToLower(valueString)]

		if !ok {
			return FileOperation(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOperation. valueString='%v' ", valueString)
		}

		result =
			FileOperation(idx)
	}

	return result, nil
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOperation
// instance.
//
func (fop FileOperation) Value() int {
	return int(fop)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time FileOperation.String() or
// FileOperation.ParseString() a call is made to this method to determine if
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
func (fop FileOperation) checkInitializeMaps(reInitialize bool) {

	if !reInitialize &&
		mFileOperationTypeIntToString != nil &&
		len(mFileOperationTypeIntToString) > 12 &&
		mFileOperationTypeStringToInt != nil &&
		len(mFileOperationTypeStringToInt) > 12 &&
		mFileOperationTypeLwrCaseStringToInt != nil &&
		len(mFileOperationTypeLwrCaseStringToInt) > 12 {
		return
	}

	var t = FileOperation(0).MoveSourceFileToDestination()

	mFileOperationTypeIntToString = make(map[int]string, 0)
	mFileOperationTypeStringToInt = make(map[string]int, 0)
	mFileOperationTypeLwrCaseStringToInt = make(map[string]int, 0)

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
		mFileOperationTypeIntToString[x] = f
		mFileOperationTypeStringToInt[f] = x
		mFileOperationTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

}

var fileOp = FileOperation(0)

// FileOp - global variable of type FileOperation.
// Provides alternate, easier access to FileOperation enumeration values.
//
// Example:
//	FileOp.CopySourceToDestinationByIo()
var FileOp = FileOperation(0)
