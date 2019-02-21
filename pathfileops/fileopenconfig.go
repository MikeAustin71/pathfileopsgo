package pathfileops

import (
	"errors"
	"fmt"
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
// The type of copy operation performed is a 'deep copy'.
//
func (fOpenCfg *FileOpenConfig) CopyIn(fOpStat2 *FileOpenConfig) {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpStat2.fileOpenModes == nil {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 0)
	}

	fOpenCfg.isInitialized = fOpStat2.isInitialized
	fOpenCfg.fileOpenType = fOpStat2.fileOpenType

	lenFOpStat2FOpenModes := len(fOpStat2.fileOpenModes)

	if lenFOpStat2FOpenModes == 0 {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 1)
		fOpStat2.fileOpenModes[0] = FOpenMode.ModeNone()
		lenFOpStat2FOpenModes = 1
	}

	fOpenCfg.fileOpenModes = make([]FileOpenMode, lenFOpStat2FOpenModes)

	for i := 0; i < lenFOpStat2FOpenModes; i++ {
		fOpenCfg.fileOpenModes[i] = fOpStat2.fileOpenModes[i]
	}

}

// CopyOut - Creates and returns a deep copy of the current
// FileOpenConfig instance.
func (fOpenCfg *FileOpenConfig) CopyOut() FileOpenConfig {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	fOpStat2 := FileOpenConfig{}
	fOpStat2.isInitialized = fOpenCfg.isInitialized
	fOpStat2.fileOpenType = fOpenCfg.fileOpenType
	lenFOpenModes := len(fOpenCfg.fileOpenModes)

	if lenFOpenModes == 0 {
		fOpenCfg.fileOpenModes = append(fOpenCfg.fileOpenModes, FOpenMode.ModeNone())
		lenFOpenModes = 1
	}

	fOpStat2.fileOpenModes = make([]FileOpenMode, lenFOpenModes)

	for i := 0; i < lenFOpenModes; i++ {
		fOpStat2.fileOpenModes[i] = fOpenCfg.fileOpenModes[i]
	}

	return fOpStat2
}

// Empty - ReInitializes the current FileOpenConfig instance to
// empty or zero values.
//
func (fOpenCfg *FileOpenConfig) Empty() {

	fOpenCfg.isInitialized = false

	fOpenCfg.fileOpenType = FOpenType.TypeNone()

	fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)

	fOpenCfg.fileOpenModes = append(fOpenCfg.fileOpenModes, FOpenMode.ModeNone())

}

// Equal - Returns 'true' if the incoming FileOpenConfig instance
// is equal in all respects to the current FileOpenConfig instance.
//
func (fOpenCfg *FileOpenConfig) Equal(fOpStat2 *FileOpenConfig) bool {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpStat2.fileOpenModes == nil {
		fOpStat2.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpenCfg.isInitialized != fOpStat2.isInitialized {
		return false
	}

	lenfOpStat1 := len(fOpenCfg.fileOpenModes)

	lenfOpStat2 := len(fOpStat2.fileOpenModes)

	if lenfOpStat1 != lenfOpStat2 {
		return false
	}

	if fOpenCfg.fileOpenType != fOpStat2.fileOpenType {
		return false
	}

	for i := 0; i < lenfOpStat1; i++ {
		isFound := false

		for j := 0; j < lenfOpStat1; j++ {
			if fOpStat2.fileOpenModes[j] == fOpenCfg.fileOpenModes[i] {
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
//  error          - If this method completes successfully, the returned error
//                   Type is set equal to 'nil'. If an error condition is encountered,
//                   this method will return an error Type which encapsulates an
//                   appropriate error message.
//
func (fOpenCfg FileOpenConfig) New(fOpenType FileOpenType, fOpenModes ...FileOpenMode) (FileOpenConfig, error) {

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

		resultFOpenStatus.fileOpenModes = append(resultFOpenStatus.fileOpenModes, FOpenMode.ModeNone())

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
func (fOpenCfg *FileOpenConfig) GetCompositeFileOpenCode() (int, error) {

	ePrefix := "FileOpenConfig.GetCompositeFileOpenCode() "

	if !fOpenCfg.isInitialized {
		return -1,
			errors.New(ePrefix + "Error: The current FileOpenConfig instance is INVALID!")
	}

	if fOpenCfg.fileOpenType == FileOpenType(0).TypeNone() {
		return -1,
			errors.New(ePrefix + "Error: The stored FileOpenType == 'None'. A valid FileOpenType is required!")
	}

	err := fOpenCfg.fileOpenType.IsValid()

	if err != nil {
		return -1,
			fmt.Errorf(ePrefix+
				"Error: The stored FileOpenType is INVALID! FileOpenType='%v' ",
				fOpenCfg.fileOpenType.Value())
	}

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	lenFileOpenModes := len(fOpenCfg.fileOpenModes)

	if lenFileOpenModes == 0 {
		return fOpenCfg.fileOpenType.Value(), nil
	}

	if lenFileOpenModes == 1 &&
		fOpenCfg.fileOpenModes[0] == FOpenMode.ModeNone() {

		return fOpenCfg.fileOpenType.Value(), nil

	}

	fileOpenVal := fOpenCfg.fileOpenType.Value()

	for i := 0; i < lenFileOpenModes; i++ {
		fileOpenVal = fileOpenVal | fOpenCfg.fileOpenModes[i].Value()
	}

	return fileOpenVal, nil
}

// GetFileOpenModes - Returns a array of stored FileOpenMode values
func (fOpenCfg *FileOpenConfig) GetFileOpenModes() []FileOpenMode {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if len(fOpenCfg.fileOpenModes) == 0 {
		fOpenCfg.fileOpenModes = append(fOpenCfg.fileOpenModes, FOpenMode.ModeNone())
	}

	resultAry := make([]FileOpenMode, 0)

	resultAry = append(resultAry, fOpenCfg.fileOpenModes...)

	return resultAry
}

// GetFileOpenType - Returns the stored FileOpenType value.
func (fOpenCfg *FileOpenConfig) GetFileOpenType() FileOpenType {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	return fOpenCfg.fileOpenType
}

// IsValid - If the current FileOpenConfig is valid and properly
// initialized, this method returns nil. If the current FileOpenConfig
// instance is invalid, this method returns an error.
func (fOpenCfg *FileOpenConfig) IsValid() error {

	ePrefix := "FileOpenConfig.IsValid() "

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if !fOpenCfg.isInitialized {
		return errors.New(ePrefix +
			"Error: The current FileOpenConfig instance has NOT been " +
			"properly initialized.")
	}

	err := fOpenCfg.fileOpenType.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: The File Open Type is INVALID!. %v", err.Error())
	}

	lenFileOpenModes := len(fOpenCfg.fileOpenModes)

	if fOpenCfg.fileOpenType == FOpenType.TypeNone() &&
		lenFileOpenModes > 1 {
		return errors.New(ePrefix +
			"Error: Current FileOpenConfig has Type='None' and " +
			"multiple File Open Modes!")
	}

	if fOpenCfg.fileOpenType == FOpenType.TypeNone() &&
		lenFileOpenModes == 1 &&
		fOpenCfg.fileOpenModes[0] != FileOpenMode(0).ModeNone() {
		return errors.New(ePrefix +
			"Error: Current FileOpenConfig has Type='None' and " +
			"a valid File Open Mode")
	}

	if fOpenCfg.fileOpenType != FOpenType.TypeNone() &&
		lenFileOpenModes > 1 {

		for i := 0; i < lenFileOpenModes; i++ {
			if fOpenCfg.fileOpenModes[i] == FileOpenMode(0).ModeNone() {
				return errors.New(ePrefix +
					"Error: The File Open Status has multiple File Open Modes " +
					"one of which is 'None'. Resolve this conflict.")
			}
		}

	}

	for i := 0; i < lenFileOpenModes; i++ {

		err := fOpenCfg.fileOpenModes[i].IsValid()

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
func (fOpenCfg *FileOpenConfig) SetFileOpenType(fOpenType FileOpenType) error {

	ePrefix := "FileOpenConfig.SetFileOpenType() "

	err := fOpenType.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"Input parameter 'fOpenType' is INVALID! fOpenType='%v' ",
			fOpenType.Value())
	}

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if fOpenType == FileOpenType(0).TypeNone() {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 1)
		fOpenCfg.fileOpenModes[0] = FOpenMode.ModeNone()
	}

	fOpenCfg.fileOpenType = fOpenType

	fOpenCfg.isInitialized = true

	return nil
}

// SetFileOpenModes - Receives a series of FileOpenMode instances and
// replaces the internal stored FileOpenMode values for this FileOpenConfig instance.
//
// To clear the current internal FileOpenMode values, pass nothing as an input parameter
// or pass the value FileOpenMode(0).None().
//
func (fOpenCfg *FileOpenConfig) SetFileOpenModes(fOpenModes ...FileOpenMode) {

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if len(fOpenModes) == 0 {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
		return
	}

	fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)

	fOpenCfg.fileOpenModes = append(fOpenCfg.fileOpenModes, fOpenModes...)

	fOpenCfg.isInitialized = true

	return
}
