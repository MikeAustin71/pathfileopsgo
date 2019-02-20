package pathfileops

import "fmt"

type FileAccessControl struct {
	isInitialized bool
	permissions   FilePermissionConfig
	fileOpenCodes FileOpenConfig
}

// New - Creates and returns a new instance of type FileAccessControl.
//
func (fAccess FileAccessControl) New(
	openCodes FileOpenConfig,
	permissions FilePermissionConfig) (FileAccessControl, error) {

	ePrefix := "FileAccessControl.New() "

	err := openCodes.IsValid()

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf(ePrefix+"Input parameter 'openCodes' is INVALID! "+
				"Error='%v' ", err.Error())
	}

	err = permissions.IsValid()

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf(ePrefix+"Input parameter 'permissions' is INVALID! "+
				"Error='%v' ", err.Error())
	}

	fA2 := FileAccessControl{}

	fA2.fileOpenCodes = openCodes.CopyOut()

	fA2.permissions = permissions.CopyOut()

	fA2.isInitialized = true

	return fA2, nil
}

// CopyIn - Receives a FileAccessControl instance and copies all the data
// fields to the current FileAccessControl instance. When complete, both
// the incoming and current FileAccessControl instances will be identical.
//
// The type of copy operation performed is a 'deep copy'.
//
func (fAccess *FileAccessControl) CopyIn(fA2 FileAccessControl) {

	fAccess.isInitialized = fA2.isInitialized

	fAccess.fileOpenCodes.CopyIn(&fA2.fileOpenCodes)

	fAccess.permissions.CopyIn(&fA2.permissions)

}

// CopyOut - Creates and returns a deep copy of the current
// FileAccessControl instance.
func (fAccess *FileAccessControl) CopyOut() FileAccessControl {

	fA2 := FileAccessControl{}

	fA2.isInitialized = fAccess.isInitialized
	fA2.fileOpenCodes = fAccess.fileOpenCodes.CopyOut()
	fA2.permissions = fAccess.permissions.CopyOut()

	return fA2
}

// Empty - ReInitializes the current FileAccessControl instance to
// empty or zero values.
//
func (fAccess *FileAccessControl) Empty() {
	fAccess.fileOpenCodes.Empty()
	fAccess.permissions.Empty()
	fAccess.isInitialized = false
}

// Equal - Returns 'true' if the incoming FileAccessControl instance
// is equal in all respects to the current FileAccessControl instance.
//
func (fAccess *FileAccessControl) Equal(fA2 FileAccessControl) bool {

	if fAccess.isInitialized != fA2.isInitialized {
		return false
	}

	if !fAccess.fileOpenCodes.Equal(&fA2.fileOpenCodes) {
		return false
	}

	if !fAccess.permissions.Equal(fA2.permissions) {
		return false
	}

	return true
}
