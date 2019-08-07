package pathfileops

import (
  "errors"
  "fmt"
  "os"
  "strings"
)

// FileAccessControl encapsulates the codes required the open files and
// configure file permissions. As such this type encapsulates types
// FilePermissionConfig and FileOpenConfig.
//
type FileAccessControl struct {
  isInitialized bool
  permissions   FilePermissionConfig
  fileOpenCodes FileOpenConfig
}

// NewInitialized - Returns a new FileAccessControl instances with all File Open
// Codes and File Permission Codes initialized to 'None'.
func (fAccess FileAccessControl) NewInitialized() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewInitialized() "
  openCodes, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeNone())

  if err != nil {
    return FileAccessControl{},
      fmt.Errorf(ePrefix+"Error returned by FileOpenConfig{}.New("+
        "FOpenType.TypeNone(), FOpenMode.ModeNone())\n"+
        "Error='%v'\n", err.Error())
  }
  entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    return FileAccessControl{},
      fmt.Errorf(ePrefix+
        "Error returned by OsFilePermCode.GetNewFromFileMode("+
        "OsFilePermCode.ModeNone()). "+
        "Error='%v' ", err.Error())
  }

  permissions, err := FilePermissionConfig{}.NewByComponents(
    entryType,
    "---------")

  if err != nil {
    return FileAccessControl{},
      fmt.Errorf(ePrefix+
        "Error returned by FilePermissionConfig{}.NewByComponents("+
        "entryType, \"---------\")\n"+
        "entryType='OsFilePermCode.ModeNone()'\n"+
        "Error='%v'", err.Error())
  }

  fA2 := FileAccessControl{}

  fA2.fileOpenCodes = openCodes.CopyOut()

  fA2.permissions = permissions.CopyOut()

  fA2.isInitialized = true

  return fA2, nil
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

// NewReadWriteAccess - Returns a FileAccessControl instance configured for
// Read/Write access.
func (fAccess FileAccessControl) NewReadWriteAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewReadWriteAccess() "

  fileOpenCfg, err :=
    FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCfg, nil
}

// NewReadWriteCreateTruncateAccess - Returns a FileAccessControl instance
// configured for Read, Write, Create and Truncate access.
//
func (fAccess FileAccessControl) NewReadWriteCreateTruncateAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewReadWriteCreateTruncateAccess() "

  //  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
  fOpenCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCfg, nil
}

// NewReadOnlyAccess - Returns a FileAccessControl instance configured for
// Read-Only access.
//
func (fAccess FileAccessControl) NewReadOnlyAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewReadOnlyAccess() "

  fileOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(), FOpenMode.ModeNone())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCfg, nil
}

// NewWriteOnlyAccess - Returns a FileAccessControl instance configured for
// Write-Only access.
//
func (fAccess FileAccessControl) NewWriteOnlyAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewWriteOnlyAccess() "

  fileOpenCfg, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("--w--w--w-")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCtrl, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCtrl, nil
}

// NewWriteOnlyAppendAccess - Returns a FileAccessControl instance configured for
// Write/Only - Append access.
//
func (fAccess FileAccessControl) NewWriteOnlyAppendAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewWriteOnlyAccess() "

  fileOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeAppend())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("--w--w--w-")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCfg, nil
}

// NewWriteOnlyTruncateAccess - Returns a FileAccessControl instance configured for
// Write/Only - Truncate access.
//
// If the file previously exists, it will be truncated before the writing operation
// commences.
//
func (fAccess FileAccessControl) NewWriteOnlyTruncateAccess() (FileAccessControl, error) {

  ePrefix := "FileAccessControl.NewWriteOnlyTruncateAccess() "

  fileOpenCfg, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeTruncate())

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("--w--w--w-")

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fileAccessCfg, nil
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
//
//  Example:
//
//        -rw-rw-rw- = returned value "0666"
//        drwxrwxrwx = returned value "020000000777"
//
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

// GetFileOpenType - Returns the File Open Type associated with the
// FileOpenConfig type stored as 'FileAccessControl.fileOpenCodes'.
//
func (fAccess *FileAccessControl) GetFileOpenType() (FileOpenType, error) {

  ePrefix := "FileAccessControl.GetFileOpenConfig() "

  err := fAccess.IsValid()

  if err != nil {
    return FileOpenType(99999), fmt.Errorf(ePrefix+"%v", err.Error())
  }

  return fAccess.fileOpenCodes.GetFileOpenType(), nil
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

// GetFilePermissionTextCode - Returns the file mode permissions expressed as
// a text string. The returned string includes the full and complete
// 10-character permission code.
//
//  Example Return Values:
//        -rwxrwxrwx
//        -rw-rw-rw-
//        drwxrwxrwx
//
func (fAccess *FileAccessControl) GetFilePermissionTextCode() (string, error) {

  ePrefix := "FileAccessControl.GetFilePermissionTextCode() "

  permTxtCode, err := fAccess.permissions.GetPermissionTextCode()

  if err != nil {
    return "", fmt.Errorf(ePrefix + "%v\n", err.Error())
  }

  return permTxtCode, nil
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
