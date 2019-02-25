package pathfileops

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// FileAccessControl encapsulates the codes required the open files and
// configure file permissions.
//
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
func (fAccess *FileAccessControl) CopyIn(fA2 *FileAccessControl) {

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
func (fAccess *FileAccessControl) Equal(fA2 *FileAccessControl) bool {

	if fAccess.isInitialized != fA2.isInitialized {
		return false
	}

	if !fAccess.fileOpenCodes.Equal(&fA2.fileOpenCodes) {
		return false
	}

	if !fAccess.permissions.Equal(&fA2.permissions) {
		return false
	}

	return true
}

// GetCompositeFileOpenCode - Returns the composite 'file open' code. This code
// is generated by combining the single FileOpenType value and zero
// or more FileOpenMode values.
//
func (fAccess *FileAccessControl) GetCompositeFileOpenCode() (int, error) {

	ePrefix := "FileAccessControl.GetCompositeFileOpenCode() "

	err := fAccess.IsValid()

	if err != nil {
		return 0, fmt.Errorf(ePrefix+"%v", err.Error())
	}

	fileOpenCodes, err := fAccess.fileOpenCodes.GetCompositeFileOpenCode()

	if err != nil {
		return 0, fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fileOpenCodes, nil
}

// GetCompositePermissionMode - Returns the complete permission code as a type
// os.FileMode.
func (fAccess *FileAccessControl) GetCompositePermissionMode() (os.FileMode, error) {

	ePrefix := "FileAccessControl.GetCompositePermissionMode() "

	err := fAccess.IsValid()

	if err != nil {
		return os.FileMode(9999), fmt.Errorf(ePrefix+"%v", err.Error())
	}

	permissionCode, err := fAccess.permissions.GetCompositePermissionMode()

	if err != nil {
		return os.FileMode(9999), fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return permissionCode, nil
}

// GetCompositePermissionModeText - Returns the composite permission file mode
// numerical value expressed as text.
func (fAccess *FileAccessControl) GetCompositePermissionModeText() string {

	ePrefix := "FileAccessControl.GetCompositePermissionModeText() "

	err := fAccess.IsValid()

	if err != nil {
		return ePrefix + "Current File Access Control Instance is INVALID! " + err.Error()
	}

	return fAccess.permissions.GetPermissionFileModeValueText()
}

// GetFileOpenAndPermissionCodes - Returns both the complete File Open Code
// and complete Permission code.
//
func (fAccess *FileAccessControl) GetFileOpenAndPermissionCodes() (int, os.FileMode, error) {

	ePrefix := "FileAccessControl.GetFileOpenAndPermissionCodes() "

	err := fAccess.IsValid()

	if err != nil {
		return -1, os.FileMode(9999), fmt.Errorf(ePrefix+"%v", err.Error())
	}

	fileOpenCode, err := fAccess.fileOpenCodes.GetCompositeFileOpenCode()

	if err != nil {
		return -1, os.FileMode(9999), fmt.Errorf(ePrefix+"%v", err.Error())
	}

	permissionCode, err := fAccess.permissions.GetCompositePermissionMode()

	if err != nil {
		return -1, os.FileMode(9999), fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fileOpenCode, permissionCode, nil
}

// GetFileOpenConfig - Returns a deep copy of the FileOpenConfig type
// encapsulated by the current FileAccessControl instance.
func (fAccess *FileAccessControl) GetFileOpenConfig() (FileOpenConfig, error) {
	ePrefix := "FileAccessControl.GetFileOpenConfig() "

	err := fAccess.IsValid()

	if err != nil {
		return FileOpenConfig{}, fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fAccess.fileOpenCodes.CopyOut(), nil
}

// GetFilePermissionConfig - Returns a deep copy of the FilePermissionConfig type
// encapsulated by the current FileAccessControl instance.
func (fAccess *FileAccessControl) GetFilePermissionConfig() (FilePermissionConfig, error) {

	ePrefix := "FileAccessControl.GetFilePermissionConfig() "

	err := fAccess.IsValid()

	if err != nil {
		return FilePermissionConfig{}, fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fAccess.permissions.CopyOut(), nil

}

// IsValid - If the current FileAccessControl instance is valid and properly
// initialized, this method returns nil. If the current FileAccessControl
// instance is invalid, this method returns an error.
func (fAccess *FileAccessControl) IsValid() error {

	ePrefix := "FileAccessControl.IsValid() "

	if !fAccess.isInitialized {
		return errors.New(ePrefix +
			"Error: The current FileAccessControl Instance has NOT been initialized!")
	}

	sb := strings.Builder{}
	sb.Grow(300)

	err := fAccess.fileOpenCodes.IsValid()

	if err != nil {
		sb.WriteString(fmt.Sprintf(ePrefix+"File Open codes INVALID! %v\n\n", err.Error()))
	}

	err = fAccess.permissions.IsValid()

	if err != nil {
		sb.WriteString(fmt.Sprintf(ePrefix+"File Permission codes INVALID! %v \n", err.Error()))
	}

	if sb.Len() > 4 {
		return fmt.Errorf("%s", sb.String())
	}

	return nil
}

// SetFileOpenCodes - Assigns 'fileOpenCodes' to internal member variable,
// FileAccessControl.fileOpenCodes
//
func (fAccess *FileAccessControl) SetFileOpenCodes(fileOpenCodes FileOpenConfig) error {

	ePrefix := "FileAccessControl.SetFileOpenCodes() "

	err := fileOpenCodes.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID 'fileOpenCodes'! - %v", err.Error())
	}

	fAccess.fileOpenCodes = fileOpenCodes.CopyOut()

	err = fAccess.permissions.IsValid()

	if err == nil {

		fAccess.isInitialized = true

	}

	return nil
}

// SetFilePermissionCodes - Assigns 'filePermissions' to internal
// member variable FileAccessControl.permissions.
//
func (fAccess *FileAccessControl) SetFilePermissionCodes(
	filePermissions FilePermissionConfig) error {

	ePrefix := "FileAccessControl.SetFilePermissionCodes() "

	err := filePermissions.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: 'filePermissions' INVALID! - %v",
			err.Error())
	}

	fAccess.permissions = filePermissions.CopyOut()

	err = fAccess.fileOpenCodes.IsValid()

	if err == nil {
		fAccess.isInitialized = true
	}

	return nil
}
