package pathfileops

import (
	"errors"
	"fmt"
	"os"
	fp "path/filepath"
	"strings"
	"time"
)

/*
	This source code file contains types 'DirMgr' and
	'DirMgrCollection'.

	The Source Repository for this source code file is :
		https://github.com/MikeAustin71/pathfileopsgo.git

	Dependencies:
	-------------

	Types 'DirMgr' and 'DirMgrCollection' depend on types,
	'FileHelper' and 'FileMgr' which are contained in source
	code files, 'filehelper.go' and 'filemanagers.go' located
	in this directory.

*/

// DirTreeOp - Contains parameters used in Directory Operations
type DirTreeOp struct {
	CallingFunc        string
	FileOps            []FileOperation
	FileSelectCriteria FileSelectionCriteria
	SourceBaseDir      DirMgr
	TargetBaseDir      DirMgr
	ErrReturns         []string
}

// New - Creates, initializes and returns a new
// DirTreeOp instance.
func (dTreeOp DirTreeOp) New() DirTreeOp {
	newDTreeOp := DirTreeOp{}
	newDTreeOp.ErrReturns = make([]string, 0, 100)
	return newDTreeOp
}

// DirMgrCollection - Holds a collection of Type DirMgr.
//
// The Source Repository for this source code file is :
// https://github.com/MikeAustin71/pathfileopsgo.git
//
// Dependencies:
//
// Type 'DirMgrCollection' depend on types, 'FileHelper' and
// 'FileMgr' which are contained in source code files: 'filehelper.go'
// and 'filemanagers.go' located in this directory.
//
type DirMgrCollection struct {
	dirMgrs []DirMgr
}

// AddDirMgr - Adds a DirMgr object to the collection
func (dMgrs *DirMgrCollection) AddDirMgr(dMgr DirMgr) {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
}

// AddDirMgrByPathNameStr - Adds a Directory Manager (DirMgr) to the
// collections based on a string input parameter, 'pathName'.
//
func (dMgrs *DirMgrCollection) AddDirMgrByPathNameStr(pathName string) error {
	ePrefix := "DirMgrCollection.AddDirMgrByPathNameStr() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	dMgr, err := DirMgr{}.New(pathName)

	if err != nil {
		s := ePrefix +
			"Error returned from DirMgr{}.New(pathName). pathName='%v' Error='%v'"
		return fmt.Errorf(s, pathName, err.Error())
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr)

	return nil
}

// AddFileMgrByFileInfo - Adds a Directory Manager object to the
// collection based on input from a directory path string and an
// os.FileInfo object.
//
func (dMgrs *DirMgrCollection) AddFileInfo(pathFile string, info os.FileInfo) error {

	ePrefix := "DirMgrCollection) AddFileMgrByFileInfo() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	dMgr, err := DirMgr{}.NewFromFileInfo(pathFile, info)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error retrned from DirMgr{}."+
			"NewFromFileInfo(pathFile, info). pathFile='%v' info.Name()='%v'  Error='%v'",
			pathFile, info.Name(), err.Error())
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr)

	return nil
}

// AddDirMgrCollection - Adds another collection of File Manager (DirMgr)
// objects to the current collection.
func (dMgrs *DirMgrCollection) AddDirMgrCollection(dMgrs2 *DirMgrCollection) {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if dMgrs2.dirMgrs == nil {
		dMgrs2.dirMgrs = make([]DirMgr, 0, 100)
	}

	lOmc2 := len(dMgrs2.dirMgrs)

	if lOmc2 == 0 {
		return
	}

	for i := 0; i < lOmc2; i++ {
		dMgrs.AddDirMgr(dMgrs2.dirMgrs[i].CopyOut())
	}

	return
}

// CopyOut - Returns an DirMgrCollection which is an
// exact duplicate of the current DirMgrCollection
func (dMgrs *DirMgrCollection) CopyOut() (DirMgrCollection, error) {

	ePrefix := "DirMgrCollection.CopyOut() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	dMgrs2 := DirMgrCollection{}.New()

	lOmc := len(dMgrs.dirMgrs)

	if lOmc == 0 {
		return DirMgrCollection{},
			errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	for i := 0; i < lOmc; i++ {
		dMgrs2.AddDirMgr(dMgrs.dirMgrs[i].CopyOut())
	}

	return dMgrs2, nil
}

// DeleteAtIndex - Deletes a member Directory Manager from
// the collection at the index specified by input parameter 'idx'.
//
// If successful, at the completion of this method, the Directory
// Manager Collection array will have a length which is one less
// than the starting array length.
//
func (dMgrs *DirMgrCollection) DeleteAtIndex(idx int) error {

	ePrefix := "DirMgrCollection.DeleteAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if idx < 0 {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return errors.New(ePrefix +
			"Error: The Directory Manager Collection, 'DirMgrCollection', is EMPTY!")
	}

	if idx >= arrayLen {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is greater than the "+
			"length of the collection index. Index Out-Of-Range! "+
			"idx='%v' Array Length='%v' ", idx, arrayLen)
	}

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	} else if idx == 0 {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
	} else if idx == arrayLen-1 {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		dMgrs.dirMgrs =
			append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)
	}

	return nil
}

// FindDirectories - searches through the DirMgrCollection to find
// DirMgr objects matching specified search criteria.
func (dMgrs *DirMgrCollection) FindDirectories(
	fileSelectionCriteria FileSelectionCriteria) (DirMgrCollection, error) {

	ePrefix := "DirMgrCollection.FindDirectories() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	lDirCol := len(dMgrs.dirMgrs)

	if lDirCol == 0 {
		return DirMgrCollection{}.New(), nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	dMgrs2 := DirMgrCollection{}.New()

	for i := 0; i < lDirCol; i++ {
		dMgr := dMgrs.dirMgrs[i]

		if dMgr.actualDirFileInfo.IsFInfoInitialized {

			isMatchedFile, err = fh.FilterFileName(dMgr.actualDirFileInfo, fileSelectionCriteria)

			if err != nil {
				return DirMgrCollection{}, fmt.Errorf(ePrefix+"Error returned by fh.FilterFileName("+
					"dMgr.actualDirFileInfo, fileSelectionCriteria) dMgr.actualDirFileInfo.Name()='%v' "+
					"Error='%v'", dMgr.actualDirFileInfo.Name(), err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(dMgr.directoryName)

			isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)

			if err != nil {
				s := ePrefix +
					"Error returned by fh.FilterFileName(fip, fileSelectionCriteria) " +
					"fip.Name()='%v'  Error='%v'"
				return DirMgrCollection{}, fmt.Errorf(s, fip.Name(), err.Error())
			}

		}

		if isMatchedFile && err == nil {
			dMgrs2.AddDirMgr(dMgr)
		}

	}

	return dMgrs2, nil
}

// GetDirMgrArray - Returns the entire Directory Manager Array managed
// by this collection.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]DirMgr      - The array of of DirMgr instances maintained by this
//	                collection.
//
func (dMgrs *DirMgrCollection) GetDirMgrArray() []DirMgr {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 10)
	}

	return dMgrs.dirMgrs
}

// GetDirMgrAtIndex - If successful, this method returns a pointer to
// the DirMgr instance at the array index specified. The 'Peek' and 'Pop'
// methods below return DirMgr objects using a 'deep' copy and therefore
// offer better protection against data corruption.
//
func (dMgrs *DirMgrCollection) GetDirMgrAtIndex(idx int) (*DirMgr, error) {

	ePrefix := "DirMgrCollection.GetDirMgrAtIndex() "

	emptyDirMgr := DirMgr{}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return &emptyDirMgr,
			fmt.Errorf(ePrefix +
				"Error: This Directory Manager Collection ('DirMgrCollection') is EMPTY!")
	}

	if idx < 0 || idx >= arrayLen {

		return &emptyDirMgr,
			fmt.Errorf(ePrefix+
				"Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
				"The minimum index is '0'. "+
				"The maximum index is '%v'. ", idx, arrayLen-1)

	}

	return &dMgrs.dirMgrs[idx], nil
}

// GetNumOfDirs - returns the number of directories
// contained in this Directory Manager Collection.
//
func (dMgrs *DirMgrCollection) GetNumOfDirs() int {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	return len(dMgrs.dirMgrs)
}

// InsertDirMgrAtIndex - Inserts a new Directory Manager into the collection at
// array 'index'. The new Directory Manager is passed as input parameter 'dMgr'.
//
// If input parameter 'index' is less than zero, an error will be returned. If
// 'index' exceeds the value of the last index in the collection, 'dMgr' will be
// added to the end of the collection at the next legal index.
func (dMgrs *DirMgrCollection) InsertDirMgrAtIndex(dMgr DirMgr, index int) error {

	ePrefix := "DirMgrCollection.InsertDirMgrAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if index < 0 {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'index' is LESS THAN ZERO! "+
			"index='%v' ", index)
	}

	lenDgrs := len(dMgrs.dirMgrs)

	if index >= lenDgrs {
		dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
		return nil
	}

	newDirMgrs := make([]DirMgr, 1, 100)

	if index == 0 {
		newDirMgrs[0] = dMgr.CopyOut()
		dMgrs.dirMgrs = append(newDirMgrs, dMgrs.dirMgrs...)
		return nil
	}

	newDirMgrs = append(newDirMgrs, dMgrs.dirMgrs[index:]...)

	dMgrs.dirMgrs = append(dMgrs.dirMgrs[:index])
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgrs...)

	return nil
}

// New - Creates and returns a new and properly initialized
// Directory Manager Collection ('DirMgrCollection').
//
func (dMgrs DirMgrCollection) New() DirMgrCollection {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	newDirMgrCol := DirMgrCollection{}
	newDirMgrCol.dirMgrs = make([]DirMgr, 0, 100)

	return newDirMgrCol
}

// PopDirMgrAtIndex - Returns a copy of the Directory Manager
// ('DirMgr') object located at index, 'idx', in the Directory
// Manager Collection ('DirMgrCollection') array.
//
// As a 'Pop' method, the original Directory Manager ('DirMgr')
// object is deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// Therefore at the completion of this method, the Directory
// Manager Collection array has a length which is one less
// than the starting array length.
//
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(idx int) (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopDirMgrAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if idx < 0 {
		return DirMgr{},
			fmt.Errorf(ePrefix+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'", idx)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Directory Manager Collection, 'DirMgrCollection', is EMPTY!")
	}

	if idx >= arrayLen {
		return DirMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is greater than the "+
			"length of the collection index. Index Out-Of-Range! "+
			"idx='%v' Array Length='%v' ", idx, arrayLen)
	}

	if idx == 0 {
		return dMgrs.PopFirstDirMgr()
	}

	if idx == arrayLen-1 {
		return dMgrs.PopLastDirMgr()
	}

	dMgr := dMgrs.dirMgrs[idx].CopyOut()

	dMgrs.dirMgrs =
		append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)

	return dMgr, nil
}

// PopFirstDirMgr  - Returns a deep copy of the first Directory Manager
// ('DirMgr') object in the Directory Manager Collection array. As a
// 'Pop' method, the original Directory Manager ('DirMgr') object is
// deleted from the Directory Manager Collection ('DirMgrCollection')
// array.
//
// Therefore at the completion of this method, the Directory Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (dMgrs *DirMgrCollection) PopFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopFirstDirMgr() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Directory Manger Collection is EMPTY!")
	}

	dMgr := dMgrs.dirMgrs[0].CopyOut()

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)

	} else {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
	}

	return dMgr, nil
}

// PopLastDirMgr - Returns a deep copy of the last Directory Manager
// ('DirMgr') object in the Directory Manager Collection array. As a
// 'Pop' method, the original Directory Manager ('DirMgr') object is
// deleted from the Directory Manager Collection ('DirMgrCollection')
// array.
//
// Therefore at the completion of this method, the Directory Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (dMgrs *DirMgrCollection) PopLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopLastDirMgr() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
	}

	dmgr := dMgrs.dirMgrs[arrayLen-1].CopyOut()

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)

	} else {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
	}

	return dmgr, nil
}

// PeekDirMgrAtIndex - Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the Directory
// Manager Collection ('DirMgrCollection'). This is a 'Peek' method
// and therefore the original Directory Manager ('DirMgr') object
// is NOT deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(idx int) (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekDirMgrAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
	}

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= arrayLen {
		return DirMgr{},
			fmt.Errorf(ePrefix+
				"Error: Input Parameter 'idx' is greater than the "+
				"length of the collection array. Index Out-Of-Range! "+
				"idx='%v' Array Length='%v' ",
				idx, arrayLen)

	}

	return dMgrs.dirMgrs[idx].CopyOut(), nil
}

// PeekFirstDirMgr - Returns a deep copy of the first Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection'). This is a 'Peek' method and therefore
// the original Directory Manager ('DirMgr') object is NOT
// deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekFirstDirMgr() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if len(dMgrs.dirMgrs) == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Director Manager Collection ('DirMgrCollection') is EMPTY!")
	}

	return dMgrs.dirMgrs[0].CopyOut(), nil
}

// PeekLastDirMgr - Returns a deep copy of the last Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection').
//
// This is a 'Peek' method and therefore the original Directory
// Manager ('DirMgr') object is NOT deleted from the Directory
// Manager Collection ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekLastDirMgr()"

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
	}

	return dMgrs.dirMgrs[arrayLen-1].CopyOut(), nil
}

// DirMgr - This structure and associated methods are used to
// manage a specific directory.
//
// The Source Repository for this source code file is:
//
//	https://github.com/MikeAustin71/pathfileopsgo.git
//
// Dependencies:
//
// Type 'DirMgr' depend on types, 'FileHelper' and 'FileMgr'
// which are contained in source code files, 'filehelper.go'
// and 'filemanagers.go' located in this directory.
//
type DirMgr struct {
	isInitialized                   bool
	originalPath                    string
	path                            string // Stored with no trailing path separator
	isPathPopulated                 bool
	doesPathExist                   bool
	parentPath                      string // Stored with no trailing path separator
	isParentPathPopulated           bool
	relativePath                    string // Stored with no preceding path separator
	isRelativePathPopulated         bool
	absolutePath                    string
	isAbsolutePathPopulated         bool
	doesAbsolutePathExist           bool
	isAbsolutePathDifferentFromPath bool
	directoryName                   string // Name of directory with out parent path.
	volumeName                      string
	isVolumePopulated               bool
	actualDirFileInfo               FileInfoPlus
}

// CopyIn - Receives a pointer to a DirMgr object as an
// input parameter and copies the values from the incoming
// object to the current DirMgr object. When the copy operation
// is completed, the current DirMgr object is a duplicate of the
// incoming DirMgr object.
func (dMgr *DirMgr) CopyIn(dmgrIn *DirMgr) {

	dMgr.isInitialized = dmgrIn.isInitialized
	dMgr.originalPath = dmgrIn.originalPath
	dMgr.path = dmgrIn.path
	dMgr.isPathPopulated = dmgrIn.isPathPopulated
	dMgr.doesPathExist = dmgrIn.doesPathExist
	dMgr.parentPath = dmgrIn.parentPath
	dMgr.isParentPathPopulated = dmgrIn.isParentPathPopulated
	dMgr.relativePath = dmgrIn.relativePath
	dMgr.isRelativePathPopulated = dmgrIn.isRelativePathPopulated
	dMgr.absolutePath = dmgrIn.absolutePath
	dMgr.isAbsolutePathPopulated = dmgrIn.isAbsolutePathPopulated
	dMgr.doesAbsolutePathExist = dmgrIn.doesAbsolutePathExist
	dMgr.isAbsolutePathDifferentFromPath = dmgrIn.isAbsolutePathDifferentFromPath
	dMgr.directoryName = dmgrIn.directoryName
	dMgr.volumeName = dmgrIn.volumeName
	dMgr.isVolumePopulated = dmgrIn.isVolumePopulated
	dMgr.actualDirFileInfo = dmgrIn.actualDirFileInfo.CopyOut()
}

// CopyOut - Makes a duplicate copy of the current DirMgr values and
// returns them in a new DirMgr object.
func (dMgr *DirMgr) CopyOut() DirMgr {

	dOut := DirMgr{}

	dOut.isInitialized = dMgr.isInitialized
	dOut.originalPath = dMgr.originalPath
	dOut.path = dMgr.path
	dOut.isPathPopulated = dMgr.isPathPopulated
	dOut.doesPathExist = dMgr.doesPathExist
	dOut.parentPath = dMgr.parentPath
	dOut.isParentPathPopulated = dMgr.isParentPathPopulated
	dOut.relativePath = dMgr.relativePath
	dOut.isRelativePathPopulated = dMgr.isRelativePathPopulated
	dOut.absolutePath = dMgr.absolutePath
	dOut.isAbsolutePathPopulated = dMgr.isAbsolutePathPopulated
	dOut.doesAbsolutePathExist = dMgr.doesAbsolutePathExist
	dOut.isAbsolutePathDifferentFromPath = dMgr.isAbsolutePathDifferentFromPath
	dOut.directoryName = dMgr.directoryName
	dOut.volumeName = dMgr.volumeName
	dOut.isVolumePopulated = dMgr.isVolumePopulated
	dOut.actualDirFileInfo = dMgr.actualDirFileInfo.CopyOut()

	return dOut
}

// DeleteAll - BE CAREFUL!!! - This method will remove the directory identified by
// this DirMgr object. It will also delete all child directories and files in the
// directory tree.
//
// Example:
// Run DeleteAll on Directory: "../pathfilego/003_filehelper/testdestdir/destdir"
// All files and all subdirectories will be deleted.
// Only the parent path will remain: "../pathfilego/003_filehelper/testdestdir"
//
func (dMgr *DirMgr) DeleteAll() error {
	ePrefix := "DirMgr.DeleteAll() "
	var err error

	err = dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return err
	}

	if dMgr.doesAbsolutePathExist {

		err = os.RemoveAll(dMgr.absolutePath)

		if err != nil {
			return fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(dMgr.absolutePath) "+
				"returned error. dMgr.absolutePath='%v' Error='%v' ", dMgr.absolutePath, err.Error())
		}

		dMgr.DoesDirMgrPathExist()
		dMgr.DoesDirMgrAbsolutePathExist()

		return nil

	} else if dMgr.doesPathExist {

		time.Sleep(time.Millisecond * 500)

		err = os.RemoveAll(dMgr.path)

		if err != nil {
			return fmt.Errorf(ePrefix+"Error returned by "+
				"os.RemoveAll(dMgr.absolutePath) returned error. "+
				"dMgr.path='%v' Error='%v' ", dMgr.path, err.Error())
		}

		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false

		return nil

	} else {
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		return nil
	}

}

// DeleteFilesInDir - Receives a string defining a pattern to use
// in searching file names for all files in the directory identified
// by the current DirMgr instance. If a file name matches the pattern
// specified by input parameter, 'fileSearchPattern', it will be deleted.
//
// Only files in the directory identified by the current DirMgr instance
// will be subject to deletion. Files in sub-directories will not be
// deleted.
//
// If the 'fileSearchPattern' is improperly formatted, an error will be returned.
//
// Example 'filePatterns'
// =====================
// *.*              will match all files in directory.
// *.html    				will match  anyfilename.html
// a*								will match  appleJack.txt
// j????row.txt     will match  j1x34row.txt
// data[0-9]*				will match 	data123.csv
//
// Reference For Matching Details:
//  https://golang.org/pkg/path/filepath/#Match
//
func (dMgr *DirMgr) DeleteFilesByNamePattern(fileSearchPattern string) error {

	ePrefix := "DirMgr.DeleteFilesInDir() "

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return err
	}

	dir, err := os.Open(dMgr.absolutePath)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error return by os.Open(dMgr.absolutePath). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	nameFileInfos, err := dir.Readdir(-1)

	if err != nil {
		_ = dir.Close()
		return fmt.Errorf(ePrefix+
			"Error returned by dir.Readdirnames(-1). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	fh := FileHelper{}

	for _, nameFInfo := range nameFileInfos {

		if nameFInfo.IsDir() {
			continue

		} else {

			fName := nameFInfo.Name()

			isMatch, err := fp.Match(fileSearchPattern, fName)

			if err != nil {

				_ = dir.Close()

				return fmt.Errorf(ePrefix+
					"Error returned by fp.Match(fileSearchPattern, fileName). "+
					"directorySearched='%v' fileSearchPattern='%v' fileName='%v' Error='%v' ",
					dMgr.absolutePath, fileSearchPattern, fName, err.Error())
			}

			if !isMatch {
				continue
			} else {

				fullName := fh.JoinPathsAdjustSeparators(dMgr.absolutePath, fName)

				err = os.Remove(fullName)

				if err != nil {
					_ = dir.Close()
					return fmt.Errorf(ePrefix+
						"Error returned by os.Remove(fullName). "+
						"fullName='%v' Error='%v' ",
						fullName, err.Error())
				}
			}
		}
	}

	err = dir.Close()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dir.Close(). "+
			"dir='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	return nil
}

// DeleteAllFilesInDir - Deletes all the files in the current
// directory. ONLY files are deleted NOT directories.
//
// Files in subdirectories are NOT deleted.
//
// Reference:
// https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
//
func (dMgr *DirMgr) DeleteAllFilesInDir() error {

	ePrefix := "DirMgr.DeleteAllFilesInDir() "

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return err
	}

	dir, err := os.Open(dMgr.absolutePath)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error return by os.Open(dMgr.absolutePath). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	fh := FileHelper{}

	nameFileInfos, err := dir.Readdir(-1)

	if err != nil {
		_ = dir.Close()
		return fmt.Errorf(ePrefix+
			"Error returned by dir.Readdirnames(-1). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	for _, nameFInfo := range nameFileInfos {

		if nameFInfo.IsDir() {
			continue

		} else {
			name := fh.JoinPathsAdjustSeparators(dMgr.absolutePath, nameFInfo.Name())

			err = os.Remove(name)

			if err != nil {
				_ = dir.Close()
				return fmt.Errorf(ePrefix+
					"Error returned by os.Remove(name). "+
					"name='%v' Error='%v' ",
					name, err.Error())
			}

		}
	}

	_ = dir.Close()

	return nil
}

// DeleteWalkDirFiles - !!! BE CAREFUL !!! This method
// deletes files in a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the current DirMgr object. The method 'walks the
// directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// If a file matches the File Selection Criteria, it is DELETED. By the way,
// if ALL the file selection criterion are set to zero values or 'Inactive',
// then ALL FILES IN THE DIRECTORY ARE DELETED!!!
//
// A record of file deletions is included in the returned DirectoryDeleteFileInfo
// structure (DirectoryDeleteFileInfo.DeletedFiles).
//
// Input Parameter:
// ================
//
// deleteFileSelectionCriteria FileSelectionCriteria
//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be deleted.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string
//					FilesOlderThan							time.Time
//					FilesNewerThan							time.Time
//					SelectByFileMode						os.FileMode
//					SelectCriterionMode					FileSelectCriterionMode // Default is AND Criterion Together
//				}
//
//			The FileSelectionCriteria type allows for configuration of
//			single or multiple file selection criterion. The 'SelectCriterionMode'
//			can be used to specify whether the file must match all
//			or any one of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define
//																		one or more search patterns. If a file
//																		matches any one of the search pattern
//																		strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file is older than
//																		the 'FilesOlderThan' date time, that file is considered
// 																		a 'match'	for this file selection criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is a uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then	a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed will be selected and DELETED.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= 0
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					will be SELECTED FOR DELETION!
//
//
// Return Value:
// =============
//
// 				type DirectoryDeleteFileInfo struct {
//									StartPath            	string
//									dirMgrs          	[]DirMgr
//									FoundFiles           	[]FileWalkInfo
//									ErrReturns           	[]string
//									DeleteFileSelectCriteria FileSelectionCriteria
//									DeletedFiles         	[]FileWalkInfo
//					}
//
//					If successful, files matching the file selection criteria
//					specified in input parameter 'deleteFileSelectionCriteria'
//  				will be DELETED and returned in a 'DirectoryDeleteFileInfo'
//  				structure field, 'DirectoryDeleteFileInfo.DeletedFiles.'
//
//					Note: It is a good idea to check the returned field
// 								DirectoryDeleteFileInfo.ErrReturns to determine if any
// 								system errors were encountered during file processing.
//
//	error	- If a program execution error is encountered during processing, it will
//					returned as an 'error' type. Also, see the comment on
// 					DirectoryDeleteFileInfo.ErrReturns, above.
//
func (dMgr *DirMgr) DeleteWalkDirFiles(
	deleteFileSelectionCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {

	ePrefix := "DirMgr.DeleteWalkDirFiles() "
	deleteFilesInfo := DirectoryDeleteFileInfo{}

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return deleteFilesInfo, err
	}

	if dMgr.doesAbsolutePathExist {

		deleteFilesInfo.StartPath = dMgr.absolutePath

	} else if dMgr.DoesDirMgrPathExist() {

		deleteFilesInfo.StartPath = dMgr.path

	} else {

		return deleteFilesInfo,
			fmt.Errorf(ePrefix+"path and absolutePath - PATH DOES NOT EXIST! "+
				"dMgr.absolutePath='%v' dMgr.path='%v'", dMgr.absolutePath, dMgr.path)
	}

	deleteFilesInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

	fh := FileHelper{}

	err = fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

	if err != nil {
		return deleteFilesInfo,
			fmt.Errorf(ePrefix+"Error returned by FileHelper."+
				"makeFileHelperWalkDirDeleteFilesFunc(&dWalkInfo). "+
				"dWalkInfo.StartPath='%v' Error='%v' ", deleteFilesInfo.StartPath, err.Error())
	}

	return deleteFilesInfo, nil

}

// DoesDirectoryExist - Returns two boolean values indicating whether or not the
// Directory path exists and whether or not the Directory absolute path exists.
//
func (dMgr *DirMgr) DoesDirectoryExist() (doesPathExist, doesAbsolutePathExist bool) {

	doesPathExist = dMgr.DoesDirMgrPathExist()

	doesAbsolutePathExist = dMgr.DoesDirMgrAbsolutePathExist()

	return
}

// DoesDirMgrAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.absolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
//
func (dMgr *DirMgr) DoesDirMgrAbsolutePathExist() bool {

	if dMgr.absolutePath == "" {
		dMgr.doesAbsolutePathExist = false
		return false
	}

	info, err := os.Stat(dMgr.absolutePath)

	if err != nil {
		dMgr.doesAbsolutePathExist = false
	} else {
		dMgr.doesAbsolutePathExist = true
		dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.absolutePath, info)
	}

	return dMgr.doesAbsolutePathExist

}

// DoesDirMgrPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition it also
// updates the DirMgr field DirMgr.doesPathExist field.
//
func (dMgr *DirMgr) DoesDirMgrPathExist() bool {

	if dMgr.path == "" {
		dMgr.isPathPopulated = false
		return false
	}

	info, err := os.Stat(dMgr.path)

	if err != nil {
		dMgr.doesPathExist = false
	} else {
		dMgr.doesPathExist = true
		dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.path, info)
	}

	return dMgr.doesPathExist

}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

	dMgr.isInitialized = false
	dMgr.originalPath = ""
	dMgr.path = ""
	dMgr.isPathPopulated = false
	dMgr.doesPathExist = false
	dMgr.parentPath = ""
	dMgr.isParentPathPopulated = false
	dMgr.relativePath = ""
	dMgr.isRelativePathPopulated = false
	dMgr.absolutePath = ""
	dMgr.isAbsolutePathPopulated = false
	dMgr.doesAbsolutePathExist = false
	dMgr.isAbsolutePathDifferentFromPath = false
	dMgr.directoryName = ""
	dMgr.volumeName = ""
	dMgr.isVolumePopulated = false
	dMgr.actualDirFileInfo = FileInfoPlus{}

}

// Equal - Compares two DirMgr objects to determine if
// they are equal.
func (dMgr *DirMgr) Equal(dmgr2 *DirMgr) bool {

	if dMgr.isInitialized != dmgr2.isInitialized ||
		dMgr.originalPath != dmgr2.originalPath ||
		dMgr.path != dmgr2.path ||
		dMgr.isPathPopulated != dmgr2.isPathPopulated ||
		dMgr.doesPathExist != dmgr2.doesPathExist ||
		dMgr.parentPath != dmgr2.parentPath ||
		dMgr.isParentPathPopulated != dmgr2.isParentPathPopulated ||
		dMgr.relativePath != dmgr2.relativePath ||
		dMgr.isRelativePathPopulated != dmgr2.isRelativePathPopulated ||
		dMgr.absolutePath != dmgr2.absolutePath ||
		dMgr.isAbsolutePathPopulated != dmgr2.isAbsolutePathPopulated ||
		dMgr.doesAbsolutePathExist != dmgr2.doesAbsolutePathExist ||
		dMgr.isAbsolutePathDifferentFromPath != dmgr2.isAbsolutePathDifferentFromPath ||
		dMgr.directoryName != dmgr2.directoryName ||
		dMgr.volumeName != dmgr2.volumeName ||
		dMgr.isVolumePopulated != dmgr2.isVolumePopulated {

		return false
	}

	if !dMgr.actualDirFileInfo.Equal(&dmgr2.actualDirFileInfo) {
		return false
	}

	return true
}

// EqualAbsPaths - compares the absolute paths for the current
// directory manager and the input directory manager ('dMgr2').
// If the two absolute paths are equal, the method returns 'true'.
// If the two absolute paths are NOT equal, the method returns 'false'.
// The comparison is NOT case sensitive. In other words, both paths
// are converted to lower case before making the comparision.
func (dMgr *DirMgr) EqualAbsPaths(dMgr2 *DirMgr) bool {

	if dMgr.isInitialized != dMgr2.isInitialized {
		return false
	}

	lcDMgrPath := strings.ToLower(dMgr.absolutePath)
	lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

	if lcDMgrPath != lcDMgr2Path {
		return false
	}

	return true
}

// EqualPaths - Compares two DirMgr objects to determine
// if their paths are equal. Both Directory Path and
// absolute path must be equivalent.
//
// If the compared paths are equal, the method returns 'true'.
// If the paths are NOT equal, the method returns 'false'.
// The comparisons are NOT case sensitive. In other words, all paths
// are converted to lower case before making the comparisions.
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

	if dMgr.isInitialized != dMgr2.isInitialized {
		return false
	}

	lcDMgrPath := strings.ToLower(dMgr.absolutePath)
	lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

	if lcDMgrPath != lcDMgr2Path {
		return false
	}

	lcDMgrPath = strings.ToLower(dMgr.path)
	lcDMgr2Path = strings.ToLower(dMgr2.path)

	if lcDMgrPath != lcDMgr2Path {
		return false
	}

	return true
}

// ExecuteDirectoryTreeOps - Performs File Operations
// on specified 'selected' files in the directory tree
// identified by the current 'DirMgr' instance.
//
// If you wish to perform File Operations ONLY on the
// current directory and NOT THE ENTIRE DIRECTORY TREE,
// see Function "ExecuteDirectoryFileOps(), below.
//
// The types of File Operations performed are generally
// classified as 'file copy' and 'file deletion' operations.
// The precise file operation applied is defined by the
// the type, 'FileOperation' which provides a series of
// constants used to identify the specific file operation
// applied. Input parameter, 'fileOps' is an array of type
// 'FileOperation' elements. Multiple file operations can
// be applied to a single file. For instance, a 'copy source
// to destination' operation can be followed by a 'delete
// source file' operation.
//
// The 'selected' files are identified by input parameter
// 'fileSelectCriteria' of type 'FileSelectionCriteria'.
// This file selection criteria is compared against all files
// in the directory tree identified by the current 'DirMgr'
// instance. When a match is found, that file is treated as
// a 'selected' file and designated file operations are
// performed on that file.
//
// IMPORTANT:
// ==========
// This method performs File Operations on THE ENTIRE DIRECTORY
// TREE identified by this DirMgr instance.
//
// Input Parameters:
// ================
//
// ---------------------------------------------------------------------------
// fileSelectCriteria FileSelectionCriteria
// ---------------------------------------------------------------------------

//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be identified
// 			as 'Selected Files'. The specified File Operations (fileOps) will be
// 			performed on these selected files.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string		// An array of strings containing File Name Patterns
//					FilesOlderThan							time.Time		// Match files with older modification date times
//					FilesNewerThan							time.Time		// Match files with newer modification date times
//					SelectByFileMode						os.FileMode	// Match file mode. Zero if inactive
//					SelectCriterionMode					FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple file
// 			selection criterion. The 'SelectCriterionMode' can be used to specify whether the
// 			file must match all, or any one, of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define one or more
//																		search patterns. If a file name matches any one of the
// 																		search pattern strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file modification
// 																		date time is older than the 'FilesOlderThan' date time,
// 																		that file is considered a 'match'	for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is an uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File Modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					in the target directory will be selected and returned
//					as 'Found Files'.
//
// ---------------------------------------------------------------------------
// fileOps []FileOperation - An array of file operations to be performed
//                           on each selected file. Selected files are
//                           identified by matching the file selection
//                           criteria specified by input parameter,
//                           'fileSelectCriteria'. See above.
// ---------------------------------------------------------------------------
// The FileOperation type consists of the following
// constants.
//
//	FileOperation(0).None()
//	  No Action
//
//	FileOperation(0).MoveSourceFileToDestination()
//	  Moves the source file to the destination file and
//	  then deletes the original source file
//
// 	FileOperation(0).DeleteDestinationFile()
//	  Deletes the Destination file if it exists
//
// 	FileOperation(0).DeleteSourceFile()
//	  Deletes the Source file if it exists
//
// 	FileOperation(0).DeleteSourceAndDestinationFiles
//	  Deletes both the Source and Destination files
//	  if they exist.
//
// 	FileOperation(0).CopySourceToDestinationByHardLinkByIo()
//	  Copies the Source File to the Destination
//	  using two copy attempts. The first copy is
//	  by Hard Link. If the first copy attempt fails,
//	  a second copy attempt is initiated/ by creating
//	  a new file and copying the contents by 'io.Copy'.
//	  An error is returned only if both copy attempts
//	  fail. The source file is unaffected.
//
//	  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// 	FileOperation(0).CopySourceToDestinationByIoByHardLink()
// 		Copies the Source File to the Destination
// 		using two copy attempts. The first copy is
// 		by 'io.Copy' which creates a new file and copies
// 		the contents to the new file. If the first attempt
// 		fails, a second copy attempt is initiated using
// 		'copy by hard link'. An error is returned only
// 		if both copy attempts fail. The source file is
// 		unaffected.
//
// 		See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// 	FileOperation(0).CopySourceToDestinationByHardLink()
// 		Copies the Source File to the Destination
// 		using one copy mode. The only copy attempt
// 		utilizes 'Copy by Hard Link'. If this fails
// 		an error is returned.  The source file is
// 		unaffected.
//
// 	FileOperation(0).CopySourceToDestinationByIo()
// 		Copies the Source File to the Destination
// 		using only one copy mode. The only copy
// 		attempt is initiated using 'Copy by IO' or
// 		'io.Copy'.  If this fails an error is returned.
// 		The source file is unaffected.
//
// 	FileOperation(0).CreateSourceDir()
// 		Creates the Source Directory
//
// 	FileOperation(0).CreateSourceDirAndFile()
// 		Creates the Source Directory and File
//
// 	FileOperation(0).CreateSourceFile()
// 		Creates the Source File
//
// 	FileOperation(0).CreateDestinationDir()
// 		Creates the Destination Directory
//
// 	FileOperation(0).CreateDestinationDirAndFile()
// 		Creates the Destination Directory and File
//
// 	FileOperation(0).CreateDestinationFile()
// 		Creates the Destination File
//
// ---------------------------------------------------------------------------
// targetBaseDir - 	The file selection criteria, 'fileSelectCriteria', and
// 									the File Operations, 'fileOps' are applied to files in
// 									the target base directory. This input parameter is of
//                  type 'DirMgr'.
// ---------------------------------------------------------------------------
//
// Function Returns:
// =================
//
// This function will return an array of strings containing error messages
// generated during the performance of specified File Operations on the
// designated directory tree. If the string array returned is empty or has
// a zero length, it signals that no errors were encountered and all operations
// completed successfully.
//
func (dMgr *DirMgr) ExecuteDirectoryTreeOps(
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperation,
	targetBaseDir DirMgr) []string {

	ePrefix := "DirMgr.ExecuteDirectoryTreeOps() "

	err := dMgr.IsDirMgrValid(ePrefix)
	dirOp := DirTreeOp{}.New()
	dirOp.CallingFunc = ePrefix

	if err != nil {
		errStr := fmt.Sprintf("%v ", err.Error())
		dirOp.ErrReturns = append(dirOp.ErrReturns, errStr)
		return dirOp.ErrReturns
	}

	err = targetBaseDir.IsDirMgrValid("")

	if err != nil {

		errStr := fmt.Sprintf(ePrefix+
			"Input parameter 'targetBaseDir' is INVALID!. Error='%v' ",
			err.Error())

		dirOp.ErrReturns = append(dirOp.ErrReturns, errStr)

		return dirOp.ErrReturns
	}

	if len(fileOps) == 0 {

		errStr := ePrefix +
			"Error: The input parameter 'fileOps' is a ZERO LENGTH ARRAY!"

		dirOp.ErrReturns = append(dirOp.ErrReturns, errStr)

		return dirOp.ErrReturns
	}

	dirOp.FileOps = append(dirOp.FileOps, fileOps...)
	dirOp.TargetBaseDir = targetBaseDir.CopyOut()
	dirOp.SourceBaseDir = dMgr.CopyOut()
	dirOp.FileSelectCriteria = fileSelectCriteria

	err = fp.Walk(dMgr.GetAbsolutePath(), dMgr.executeFileOpsOnFoundFiles(&dirOp))

	if err != nil {
		errStr := ePrefix +
			fmt.Sprintf("Error returned by fp.Walk(). Error='%v' ", err.Error())
		dirOp.ErrReturns = append(dirOp.ErrReturns, errStr)
	}

	return dirOp.ErrReturns
}

// ExecuteDirectoryFileOps - Performs a a file operation
// on specified 'selected' files in the current directory
// ONLY. This function does NOT perform operations on the
// Directory Tree.
//
// To perform file operations on the entire Directory Tree,
// see Function 'ExecuteDirectoryTreeOps()', above.
//
// The types of File Operations performed are generally
// classified as 'file copy' and 'file deletion' operations.
// The precise file operation applied is defined by the
// the type, 'FileOperation' which provides a series of
// constants used to identify the specific file operation
// applied. Input parameter, 'fileOps' is an array of type
// 'FileOperation' elements. Multiple file operations can
// be applied to a single file. For instance, a 'copy source
// to destination' operation can be followed by a 'delete
// source file' operation.
//
// The 'selected' files are identified by input parameter
// 'fileSelectCriteria' of type 'FileSelectionCriteria'.
// This file selection criteria is compared against all files
// in the directory (NOT the Directory Tree) identified by
// the current 'DirMgr' instance. When a match is found,
// that file is treated as a 'selected' file and designated
// file operations are performed on that file.
//
// IMPORTANT:
// ==========
// This method performs File Operations ONLY on the directory
// identified by the current DirMgr instance.
//
// Input Parameters:
// ================
//
// ---------------------------------------------------------------------------
// fileSelectCriteria FileSelectionCriteria
// ---------------------------------------------------------------------------

//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be identified
// 			as 'Selected Files'. The specified File Operations (fileOps) will be
// 			performed on these selected files.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string		// An array of strings containing File Name Patterns
//					FilesOlderThan							time.Time		// Match files with older modification date times
//					FilesNewerThan							time.Time		// Match files with newer modification date times
//					SelectByFileMode						os.FileMode	// Match file mode. Zero if inactive
//					SelectCriterionMode					FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple file
// 			selection criterion. The 'SelectCriterionMode' can be used to specify whether the
// 			file must match all, or any one, of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define one or more
//																		search patterns. If a file name matches any one of the
// 																		search pattern strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file modification
// 																		date time is older than the 'FilesOlderThan' date time,
// 																		that file is considered a 'match'	for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is an uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File Modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					in the target directory will be selected and returned
//					as 'Found Files'.
//
// ---------------------------------------------------------------------------
// fileOps []FileOperation - An array of file operations to be performed
//                           on each selected file. Selected files are
//                           identified by matching the file selection
//                           criteria specified by input parameter,
//                           'fileSelectCriteria'. See above.
// ---------------------------------------------------------------------------
// The FileOperation type consists of the following
// constants.
//
//	FileOperation(0).MoveSourceFileToDestination() FileOperation = iota
// 		Moves the source file to the destination file and
// 		then deletes the original source file
//
// 	FileOperation(0).DeleteDestinationFile()
// 		Deletes the Destination file if it exists
//
// 	FileOperation(0).DeleteSourceFile()
// 		Deletes the Source file if it exists
//
// 	FileOperation(0).DeleteSourceAndDestinationFiles
// 		Deletes both the Source and Destination files
// 		if they exist.
//
// 	FileOperation(0).CopySourceToDestinationByHardLinkByIo()
// 		Copies the Source File to the Destination
// 		using two copy attempts. The first copy is
// 		by Hard Link. If the first copy attempt fails,
// 		a second copy attempt is initiated/ by creating
// 		a new file and copying the contents by 'io.Copy'.
// 		An error is returned only if both copy attempts
// 		fail. The source file is unaffected.
//
// 		See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// 	FileOperation(0).CopySourceToDestinationByIoByHardLink()
// 		Copies the Source File to the Destination
// 		using two copy attempts. The first copy is
// 		by 'io.Copy' which creates a new file and copies
// 		the contents to the new file. If the first attempt
// 		fails, a second copy attempt is initiated using
// 		'copy by hard link'. An error is returned only
// 		if both copy attempts fail. The source file is
// 		unaffected.
//
// 		See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// 	FileOperation(0).CopySourceToDestinationByHardLink()
// 		Copies the Source File to the Destination
// 		using one copy mode. The only copy attempt
// 		utilizes 'Copy by Hard Link'. If this fails
// 		an error is returned.  The source file is
// 		unaffected.
//
// 	FileOperation(0).CopySourceToDestinationByIo()
// 		Copies the Source File to the Destination
// 		using only one copy mode. The only copy
// 		attempt is initiated using 'Copy by IO' or
// 		'io.Copy'.  If this fails an error is returned.
// 		The source file is unaffected.
//
// 	FileOperation(0).CreateSourceDir()
// 		Creates the Source Directory
//
// 	FileOperation(0).CreateSourceDirAndFile()
// 		Creates the Source Directory and File
//
// 	FileOperation(0).CreateSourceFile()
// 		Creates the Source File
//
// 	FileOperation(0).CreateDestinationDir()
// 		Creates the Destination Directory
//
// 	FileOperation(0).CreateDestinationDirAndFile()
// 		Creates the Destination Directory and File
//
// 	FileOperation(0).CreateDestinationFile()
// 		Creates the Destination File
//
// ---------------------------------------------------------------------------
// targetBaseDir - 	The file selection criteria, 'fileSelectCriteria', and
// 									the File Operations, 'fileOps' are applied to files in
// 									the target base directory. This input parameter is of
//                  type 'DirMgr'.
// ---------------------------------------------------------------------------
//
// Function Returns:
// =================
//
// This function will return an array of strings containing error messages
// generated during the performance of specified File Operations on the
// designated directory tree. If the string array returned is empty or has
// a zero length, it signals that no errors were encountered and all operations
// completed successfully.
//
func (dMgr *DirMgr) ExecuteDirectoryFileOps(
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperation,
	targetBaseDir DirMgr) []string {

	ePrefix := "DirMgr.ExecuteDirectoryFileOps() "
	errStrs := make([]string, 0, 50)
	var errStr string

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		errStr = fmt.Sprintf("%v ", err.Error())
		errStrs = append(errStrs, errStr)
		return errStrs
	}

	err = targetBaseDir.IsDirMgrValid("")

	if err != nil {

		errStr = fmt.Sprintf(ePrefix+
			"Input parameter 'targetBaseDir' is INVALID!. Error='%v' ",
			err.Error())

		errStrs = append(errStrs, errStr)

		return errStrs
	}

	if len(fileOps) == 0 {

		errStr = ePrefix +
			"Error: The input parameter 'fileOps' is a ZERO LENGTH ARRAY!"

		errStrs = append(errStrs, errStr)

		return errStrs
	}

	dir, err := os.Open(dMgr.absolutePath)

	if err != nil {
		errStr = fmt.Sprintf(ePrefix+
			"Error return by os.Open(dMgr.absolutePath). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())

		errStrs = append(errStrs, errStr)
		return errStrs
	}

	nameFileInfos, err := dir.Readdir(-1)

	if err != nil {
		_ = dir.Close()
		errStr = fmt.Sprintf(ePrefix+
			"Error returned by dir.Readdirnames(-1). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())

		errStrs = append(errStrs, errStr)
		return errStrs

	}

	fh := FileHelper{}

	for _, nameFInfo := range nameFileInfos {

		if nameFInfo.IsDir() {
			continue
		}

		// Must be a file - process it!

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.
		isMatch, err := fh.FilterFileName(nameFInfo, fileSelectCriteria)

		if err != nil {

			_ = dir.Close()

			errStr = fmt.Sprintf(ePrefix+
				"Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
				"directorySearched='%v'  fileName='%v' Error='%v' ",
				dMgr.absolutePath, nameFInfo.Name(), err.Error())

			errStrs = append(errStrs, errStr)
			return errStrs
		}

		if !isMatch {

			continue

		}

		// Must be a match - this is a 'selected' file!

		srcFileNameExt := nameFInfo.Name()

		fileOp, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
			dMgr.GetAbsolutePath(),
			nameFInfo.Name(),
			targetBaseDir.GetAbsolutePath(),
			srcFileNameExt)

		if err != nil {

			_ = dir.Close()

			errStr = fmt.Sprintf(ePrefix+
				"Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs() "+
				"sourcePath='%v' srcFileNameExt='%v' destDir='%v' Error='%v' ",
				dMgr.GetAbsolutePath(), srcFileNameExt, targetBaseDir.GetAbsolutePath(),
				err.Error())
			errStrs = append(errStrs, errStr)
			return errStrs
		}

		maxOps := len(fileOps)

		for i := 0; i < maxOps; i++ {

			err = fileOp.ExecuteFileOperation(fileOps[i])

			if err != nil {
				errStr = fmt.Sprintf(ePrefix+
					"Error returned by fileOp.ExecuteFileOperation(fileOps[i]). "+
					"i='%v' FileOps='%v' Error='%v' ",
					i, fileOps[i].String(), err.Error())

				// Store the error and continue processing
				// file operations.
				errStrs = append(errStrs, errStr)
			}
		}

		// finished applying file operations to this file.
		// Get another one and continue...
	}

	_ = dir.Close()

	return errStrs
}

// FindFilesByNamePattern - Searches files in the current directory ONLY. An attempt
// will be made to match the file name with the specified search pattern string.
// All matched files will be returned in a FileMgrCollection.
//
// Regardless of the search pattern used, this method will never return sub-directories
// of the target search directory.
//
// Again, the file search will always be limited to the directory identified by the
// current DirMgr instance. No sub-directories will be searched.
//
// If the 'fileSearchPattern' is improperly formatted, an error will be returned.
//
// Example 'filePatterns'
// =====================
// *.*              will match all files in directory.
// *.html    				will match  anyfilename.html
// a*								will match  appleJack.txt
// j????row.txt     will match  j1x34row.txt
// data[0-9]*				will match 	data123.csv
//
// Reference For Matching Details:
//  https://golang.org/pkg/path/filepath/#Match
//
func (dMgr *DirMgr) FindFilesByNamePattern(fileSearchPattern string) (FileMgrCollection, error) {

	ePrefix := "DirMgr.FindFilesByNamePattern() "

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return FileMgrCollection{}, err
	}

	dir, err := os.Open(dMgr.absolutePath)

	if err != nil {
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error return by os.Open(dMgr.absolutePath). "+
				"dMgr.absolutePath='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	nameFileInfos, err := dir.Readdir(-1)

	if err != nil {
		_ = dir.Close()
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error returned by dir.Readdirnames(-1). "+
				"dMgr.absolutePath='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	fMgrCol := FileMgrCollection{}

	for _, nameFInfo := range nameFileInfos {

		if nameFInfo.IsDir() {
			continue

		} else {

			fName := nameFInfo.Name()

			isMatch, err := fp.Match(fileSearchPattern, fName)

			if err != nil {

				_ = dir.Close()

				return FileMgrCollection{},
					fmt.Errorf(ePrefix+
						"Error returned by fp.Match(fileSearchPattern, fileName). "+
						"directorySearched='%v' fileSearchPattern='%v' fileName='%v' Error='%v' ",
						dMgr.absolutePath, fileSearchPattern, fName, err.Error())
			}

			if !isMatch {
				continue
			} else {

				err = fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo)

				if err != nil {
					_ = dir.Close()
					return FileMgrCollection{},
						fmt.Errorf(ePrefix+
							"Error returned by fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo). "+
							"Directory='%v' FileName='%v' Error='%v' ",
							dMgr.absolutePath, fName, err.Error())
				}
			}
		}
	}

	err = dir.Close()

	if err != nil {
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error returned by dir.Close(). "+
				"dir='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	return fMgrCol, nil
}

// FindFilesBySelectCriteria - Conducts a file search in the directory
// identified by the current DirMgr instance. The file search is limited
// to that directory ONLY. No sub-directories will be searched.
//
// Files matching the "FileSectionCriteria" instance passed as an input
// parameter will be used to screen available files. Any files matching
// the file selection criteria will be returned in a 'FileMgrCollection'.
//
// Only matched files will be returned. No sub-directory names will ever
// be included.
//
// The use of a 'FileSelectionCriteria' structure allows for very flexible
// and granular file searches.
//
// Input Parameter:
// ================
//
// fileSelectCriteria FileSelectionCriteria
//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be returned as
// 			'Found Files'.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string		// An array of strings containing File Name Patterns
//					FilesOlderThan							time.Time		// Match files with older modification date times
//					FilesNewerThan							time.Time		// Match files with newer modification date times
//					SelectByFileMode						os.FileMode	// Match file mode. Zero if inactive
//					SelectCriterionMode					FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple file
// 			selection criterion. The 'SelectCriterionMode' can be used to specify whether the
// 			file must match all, or any one, of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define one or more
//																		search patterns. If a file name matches any one of the
// 																		search pattern strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file modification
// 																		date time is older than the 'FilesOlderThan' date time,
// 																		that file is considered a 'match'	for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is an uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File Modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					in the target directory will be selected and returned
//					as 'Found Files'.
//
func (dMgr *DirMgr) FindFilesBySelectCriteria(
	fileSelectCriteria FileSelectionCriteria) (FileMgrCollection, error) {

	ePrefix := "DirMgr.FindFilesBySelectCriteria() "

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return FileMgrCollection{}, err
	}

	dir, err := os.Open(dMgr.absolutePath)

	if err != nil {
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error return by os.Open(dMgr.absolutePath). "+
				"dMgr.absolutePath='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	nameFileInfos, err := dir.Readdir(-1)

	if err != nil {
		_ = dir.Close()
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error returned by dir.Readdirnames(-1). "+
				"dMgr.absolutePath='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	fMgrCol := FileMgrCollection{}
	fh := FileHelper{}

	for _, nameFInfo := range nameFileInfos {

		if nameFInfo.IsDir() {
			continue

		} else {

			// This is not a directory. It is a file.
			// Determine if it matches the find file criteria.
			isMatch, err := fh.FilterFileName(nameFInfo, fileSelectCriteria)

			if err != nil {

				_ = dir.Close()

				return FileMgrCollection{},
					fmt.Errorf(ePrefix+
						"Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria). "+
						"directorySearched='%v'  fileName='%v' Error='%v' ",
						dMgr.absolutePath, nameFInfo.Name(), err.Error())
			}

			if !isMatch {

				continue

			} else {

				err = fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo)

				if err != nil {
					_ = dir.Close()
					return FileMgrCollection{},
						fmt.Errorf(ePrefix+
							"Error returned by fMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo). "+
							"Directory='%v' FileName='%v' Error='%v' ",
							dMgr.absolutePath, nameFInfo.Name(), err.Error())
				}
			}
		}
	}

	err = dir.Close()

	if err != nil {
		return FileMgrCollection{},
			fmt.Errorf(ePrefix+
				"Error returned by dir.Close(). "+
				"dir='%v' Error='%v' ",
				dMgr.absolutePath, err.Error())
	}

	return fMgrCol, nil
}

// FindWalkDirFiles - This method returns file information on files residing in a specific
// directory tree identified by the current DirMgr object.
//
// This method 'walks the directory tree' locating all files in the directory tree which match
// the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. By the way, if ALL the file selection criterion are set to
// zero values or 'Inactive', then ALL FILES in the directory are selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// Input Parameters:
// ================
//
// fileSelectCriteria FileSelectionCriteria
//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be returned as
// 			'Found Files'.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string		// An array of strings containing File Name Patterns
//					FilesOlderThan							time.Time		// Match files with older modification date times
//					FilesNewerThan							time.Time		// Match files with newer modification date times
//					SelectByFileMode						os.FileMode	// Match file mode. Zero if inactive
//					SelectCriterionMode					FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple file
// 			selection criterion. The 'SelectCriterionMode' can be used to specify whether the
// 			file must match all, or any one, of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define one or more
//																		search patterns. If a file name matches any one of the
// 																		search pattern strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file modification
// 																		date time is older than the 'FilesOlderThan' date time,
// 																		that file is considered a 'match'	for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is an uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File Modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					in the target directory will be selected and returned
//					as 'Found Files'.
//
//
// Return Value:
// =============
//
//	DirectoryTreeInfo structure	-
//					type DirectoryTreeInfo struct {
//						StartPath            	string								// The starting path or directory for the file search
//						Directories      	    DirMgrCollection	   	// dirMgrs found during directory tree search
//						FoundFiles           	FileMgrCollection		 // Found Files matching file selection criteria
//						ErrReturns           	[]string							// Internal System errors encountered
//						FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
// 																												// input parameter to this method.
//					}
//
//					If successful, files matching the file selection criteria input
//  				parameter shown above will be returned in a 'DirectoryTreeInfo'
//  				object. The field 'DirectoryTreeInfo.FoundFiles' contains information
// 					on all the files in the specified directory tree which match the file selection
// 					criteria.
//
//					Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
// 								to determine if any internal system errors were encountered while processing
// 								the directory tree.
//
//	error	- If a program execution error is encountered during processing, it will
//					be returned as an 'error' type. Also, see the comment on
// 					'DirectoryTreeInfo.ErrReturns', above.
//
func (dMgr *DirMgr) FindWalkDirFiles(
	fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

	ePrefix := "DirMgr.GetWalkDirInfo() "
	findFilesInfo := DirectoryTreeInfo{}

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return findFilesInfo, err
	}

	if dMgr.doesAbsolutePathExist {

		findFilesInfo.StartPath = dMgr.absolutePath

	} else if dMgr.DoesDirMgrPathExist() {

		findFilesInfo.StartPath = dMgr.path

	} else {

		return findFilesInfo,
			fmt.Errorf(ePrefix+
				"path and absolutePath - PATH DOES NOT EXIST! "+
				"dMgr.absolutePath='%v' dMgr.path='%v'", dMgr.absolutePath, dMgr.path)
	}

	findFilesInfo.FileSelectCriteria = fileSelectCriteria

	fh := FileHelper{}

	err = fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

	if err != nil {
		return findFilesInfo, fmt.Errorf(ePrefix+
			"Error returned by FileHelper.FindFilesWalkDirectory(&dWalkInfo). "+
			"dWalkInfo.StartPath='%v' Error='%v' ", findFilesInfo.StartPath, err.Error())
	}

	return findFilesInfo, nil
}

// GetAbsolutePath - Returns a string containing the
// absolute path for the current Directory Manager
// instance. This string returned by this method
// will NOT have a trailing path separator.
//
func (dMgr *DirMgr) GetAbsolutePath() string {

	return dMgr.absolutePath
}

// GetAbsolutePathElements - Returns all of the directories and drive
// specifications as an array of strings.
//
// Example
//
// Path = "D:\ADir\BDir\CDir\EDir"
//
// Returned pathElements string array:
//   pathElements[0] = "D:"
//   pathElements[1] = "ADir"
//   pathElements[2] = "BDir"
//   pathElements[3] = "CDir"
//   pathElements[4] = "DDir"
//   pathElements[4] = "EDir"
//
func (dMgr *DirMgr) GetAbsolutePathElements() (pathElements []string) {

	pathElements = make([]string, 0, 50)

	pathStr := dMgr.GetAbsolutePath()

	if len(pathStr) == 0 {
		return pathElements
	}

	pathStr = strings.Replace(pathStr, "\\", "/", -1)

	pathElements = strings.Split(pathStr, "/")

	return pathElements
}

// GetAbsolutePathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetAbsolutePathWithSeparator() string {
	lPath := len(dMgr.absolutePath)

	if lPath == 0 {
		return ""
	}

	if dMgr.absolutePath[lPath-1] != os.PathSeparator {
		return dMgr.absolutePath + string(os.PathSeparator)
	}

	return dMgr.absolutePath
}

// GetDirectoryName - Returns a string containing the name
// of the directory without out the parent path.
//
func (dMgr *DirMgr) GetDirectoryName() string {
	return dMgr.directoryName
}

// GetFileInfoPlus - Returns a FileInfoPlus instance detailing file
// system information on the directory identified by the current
// Directory Manager instance.
//
func (dMgr *DirMgr) GetFileInfoPlus() (FileInfoPlus, error) {

	ePrefix := "DirMgr.GetFileInfoPlus() "

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return FileInfoPlus{}, err
	}

	if dMgr.absolutePath == "" {

		dMgr.doesAbsolutePathExist = false

		return FileInfoPlus{},
			errors.New(ePrefix + "Directory DOES NOT EXIST! FileInfoPlus is NOT available.")
	}

	info, err := os.Stat(dMgr.absolutePath)

	if err != nil {

		return FileInfoPlus{},
			fmt.Errorf(ePrefix+
				"Error returned by os.Stat(dMgr.absolutePath). Error='%v'",
				err.Error())
	}

	dMgr.doesAbsolutePathExist = true
	dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.absolutePath, info)

	return dMgr.actualDirFileInfo.CopyOut(), nil
}

// GetNumberOfAbsPathElements - Returns the number of elements
// or path components in the absolute path of the current
// Directory Manager instance.
func (dMgr *DirMgr) GetNumberOfAbsPathElements() int {

	pathElements := dMgr.GetAbsolutePathElements()

	return len(pathElements)
}

// GetOriginalPath - Returns the original path used to initialize
// this Directory Manager instance.
//
func (dMgr *DirMgr) GetOriginalPath() string {
	return dMgr.originalPath
}

// GetParentDirMgr - Returns a new Directory Manager instance
// which represents the the parent path for the current
// Directory Manager. The current Directory Manager absolute
// path is used in extracting the parent Directory Manager.
//
// Return Values:
//
//	dirMgr   DirMgr - If successful, this methods returns a Directory Manager
//	                  which is a parent of the current Directory Manager.
//
//	hasParent  bool - If 'true', it signals that the current Directory Manager
//	                  has a valid parent directory. If 'false', it signals that
//	                  the current Directory Manager represents a top level directory
//	                  which has no parent directory. In that case a copy of the
//	                  current Directory will be returned.
//
//	err       error - If an error is encountered this error type will be populated
//	                  with an appropriate error message. Otherwise, a value of 'nil'
//	                  will be returned.
//
//	                  If 'hasParent' is 'false', no error will be returned.
//
func (dMgr *DirMgr) GetParentDirMgr() (dirMgr DirMgr, hasParent bool, err error) {

	ePrefix := "DirMgr.GetParentDirMgr() Error: "
	dirMgr = DirMgr{}
	hasParent = true
	err = nil

	if len(dMgr.parentPath) == 0 {

		return dMgr.CopyOut(), false, nil

	}

	var err2 error

	dirMgr, err2 = DirMgr{}.New(dMgr.parentPath)

	if err2 != nil {

		err = fmt.Errorf(ePrefix+"%v", err.Error())
		hasParent = true
		dirMgr = DirMgr{}
		return dirMgr, hasParent, err
	}

	err = nil

	return dirMgr, hasParent, err
}

// GetParentPath - Returns a string containing the
// parent path for the current Directory Manager
// instance. The Parent Path string will NOT contain
// a trailing path separator.
//
func (dMgr *DirMgr) GetParentPath() string {
	return dMgr.parentPath
}

// GetPath - Returns the path used to configure this
// Directory Manager Instance. It will NOT contain a
// trailing path separator. It may or may not be an
// absolute path.
//
func (dMgr *DirMgr) GetPath() string {
	return dMgr.path
}

// GetPathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetPathWithSeparator() string {
	lPath := len(dMgr.path)

	if lPath == 0 {
		return ""
	}

	if dMgr.path[lPath-1] != os.PathSeparator {
		return dMgr.path + string(os.PathSeparator)
	}

	return dMgr.path
}

// GetRelativePath - Returns a string containing the relative path
// for this Directory Manager instance.
//
func (dMgr *DirMgr) GetRelativePath() string {
	return dMgr.relativePath
}

// GetThisDirectoryTree - Returns a DirMgrCollection containing all
// the directories in the path of the parent directory identified by
// the current DirMgr instance.
//
// The returned DirMgrCollection will always contain the parent directory
// and will therefore always consist of at least one directory. If
// sub-directories are found, then the returned DirMgrCollection will
// contain more than one directory.
//
func (dMgr *DirMgr) GetThisDirectoryTree() (DirMgrCollection, error) {

	ePrefix := "DirMgr.GetThisDirectoryTree() "

	dMgrs := DirMgrCollection{}

	err := dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return dMgrs, err
	}

	dMgrs.AddDirMgr(dMgr.CopyOut())

	fh := FileHelper{}

	maxLen := 1

	for i := 0; i < maxLen; i++ {

		dir, err := os.Open(dMgrs.dirMgrs[i].absolutePath)

		if err != nil {
			return DirMgrCollection{},
				fmt.Errorf(ePrefix+
					"Error return by os.Open(dMgrs.dirMgrs[i].absolutePath). "+
					"dMgr.absolutePath='%v' Error='%v' ",
					dMgrs.dirMgrs[i].absolutePath, err.Error())
		}

		nameFileInfos, err := dir.Readdir(-1)

		if err != nil {
			_ = dir.Close()
			return DirMgrCollection{},
				fmt.Errorf(ePrefix+
					"Error returned by dir.Readdirnames(-1). "+
					"dMgr.absolutePath='%v' Error='%v' ",
					dMgr.absolutePath, err.Error())
		}

		for _, nameFInfo := range nameFileInfos {

			if nameFInfo.IsDir() {

				newDirPathFileName :=
					fh.JoinPathsAdjustSeparators(dMgrs.dirMgrs[i].absolutePath, nameFInfo.Name())

				fmt.Println("Next Dir: ", newDirPathFileName)

				// err = dMgrs.AddFileInfo(newDirPathFileName, nameFInfo)
				err = dMgrs.AddDirMgrByPathNameStr(newDirPathFileName)

				if err != nil {
					return DirMgrCollection{},
						fmt.Errorf(ePrefix+
							"Error returned by dMgrs.AddDirMgrByPathNameStr(newDirPathFileName). "+
							"dir='%v' Error='%v' ",
							newDirPathFileName, err.Error())
				}

				fmt.Println("dMgrs Length", dMgrs.GetNumOfDirs())

				maxLen++

			}
		}

		err = dir.Close()

		if err != nil {
			return DirMgrCollection{},
				fmt.Errorf(ePrefix+
					"Error returned by dir.Close(). "+
					"dir='%v' Error='%v' ",
					dMgr.absolutePath, err.Error())
		}

	}

	return dMgrs, nil
}

// GetVolumeName - Returns a string containing the volume name
// of the directory identified by the current Directory Manager
// instance.
//
func (dMgr *DirMgr) GetVolumeName() string {
	return dMgr.volumeName
}

// IsAbsolutePathDifferentFromPath - Returns a boolean value indicating
// whether the absolute path differs from the base path for this
// Directory Manager instance.
//
func (dMgr *DirMgr) IsAbsolutePathDifferentFromPath() bool {

	if dMgr.path != dMgr.absolutePath {
		dMgr.isAbsolutePathDifferentFromPath = true
	} else {
		dMgr.isAbsolutePathDifferentFromPath = false
	}

	return dMgr.isAbsolutePathDifferentFromPath
}

// IsAbsolutePathPopulated - Returns a boolean value indicating
// whether the absolute path for the current Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsAbsolutePathPopulated() bool {

	if len(dMgr.absolutePath) == 0 {
		dMgr.isAbsolutePathPopulated = false
	} else {
		dMgr.isAbsolutePathPopulated = true
	}

	return dMgr.isAbsolutePathPopulated
}

// IsDirMgrValid - This method examines the current DirMgr object
// to determine whether it has been properly configured.
// If the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error is
// returned.
func (dMgr *DirMgr) IsDirMgrValid(errPrefixStr string) error {

	ePrefix := strings.TrimRight(errPrefixStr, " ") + " DirMgr.IsDirMgrValid() "

	if !dMgr.isInitialized {
		return fmt.Errorf(ePrefix + "Error: DirMgr is NOT Initialized.")
	}

	if dMgr.absolutePath == "" {
		return fmt.Errorf(ePrefix + "Error: DirMgr.absolutePath is EMPTY!.")
	}

	dMgr.isAbsolutePathPopulated = true

	if dMgr.path == "" {
		return fmt.Errorf(ePrefix + "Error: DirMgr.absolutePath is EMPTY!.")
	}

	dMgr.isPathPopulated = true

	dMgr.DoesDirMgrAbsolutePathExist()
	dMgr.DoesDirMgrPathExist()

	return nil
}

// IsInitialized - Returns a boolean value indicating
// whether the Directory Manager instance is initialized.
//
func (dMgr *DirMgr) IsInitialized() bool {

	return dMgr.isInitialized

}

// IsParentPathPopulated - Returns a boolean value
// indicating whether the parent path for this Directory
// Manager instance is populated.
//
func (dMgr *DirMgr) IsParentPathPopulated() bool {

	if len(dMgr.parentPath) == 0 {
		dMgr.isPathPopulated = false
	} else {
		dMgr.isPathPopulated = true
	}

	return dMgr.isPathPopulated
}

// IsPathPopulated - Returns a boolean value indicating
// whether the current Directory Manager path string is
// populated.
//
func (dMgr *DirMgr) IsPathPopulated() bool {

	if len(dMgr.path) == 0 {
		dMgr.isPathPopulated = false
	} else {
		dMgr.isPathPopulated = true
	}

	return dMgr.isPathPopulated
}

// IsRelativePathPopulated - Returns a boolean value indicating
// whether the Relative Path for this Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsRelativePathPopulated() bool {

	if len(dMgr.relativePath) == 0 {
		dMgr.isRelativePathPopulated = false
	} else {
		dMgr.isRelativePathPopulated = true
	}

	return dMgr.isRelativePathPopulated
}

// IsVolumeNamePopulated - Returns a boolean value indicating
// whether the Volume Name for the current Directory Manager
// instance is populated.
//
func (dMgr *DirMgr) IsVolumeNamePopulated() bool {

	if len(dMgr.volumeName) == 0 {
		dMgr.isVolumePopulated = false
	} else {
		dMgr.isVolumePopulated = true
	}

	return dMgr.isVolumePopulated
}

// MakeDir - If the directory path identified
// by the current DirMgr object does not exist,
// this method will create that directory path.
func (dMgr *DirMgr) MakeDir() error {

	ePrefix := "DirMgr.MakeDir() "
	var err error

	err = dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return nil
	}

	var ModePerm os.FileMode = 0777

	if dMgr.isAbsolutePathPopulated {

		if dMgr.doesAbsolutePathExist {
			// No need to create directory, it already
			// exists.
			return nil
		}

		err = os.MkdirAll(dMgr.absolutePath, ModePerm)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned from os.MkdirAll(dMgr.absolutePath, "+
				"ModePerm) dMgr.absolutePath='%v' ModePerm='%v'  Error='%v'",
				dMgr.absolutePath, ModePerm, err.Error())
		}

		dMgr.DoesDirMgrPathExist()
		dMgr.DoesDirMgrAbsolutePathExist()

		// No errors - directory created.
		return nil
	}

	// dMgr.isPathPopulated MUST equal 'true'

	if dMgr.doesPathExist {
		// No need to create directory, it already
		// exists.
		return nil
	}

	err = os.MkdirAll(dMgr.path, ModePerm)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from os.MkdirAll(dMgr.path, ModePerm) "+
			"dMgr.path='%v' ModePerm='%v'  Error='%v'",
			dMgr.path, ModePerm, err.Error())
	}

	dMgr.DoesDirMgrPathExist()
	dMgr.DoesDirMgrAbsolutePathExist()

	// No errors - directory created.
	return nil
}

// NewFromPathFileNameExtStr - Returns a new DirMgr object and populates the
// the data fields.
//
// Input Parameters:
// =================
//
// pathStr string 	- A path string designating a path or directory.
// 										To reduce errors, the 'pathStr' should be terminated
//										with an appropriate path separator ('/' or '\')
//										Example 'pathStr': "C:\dirA\dirB\dirC\"
//
// Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//
//		isInitialized:  true
//		Original path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		PathIsPopuslated:  true
//		doesPathExist:  true
//		parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//		isParentPathPopulated:  true
//		relativePath:  testoverwrite
//		isRelativePathPopulated:  true
//		absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		isAbsolutePathPopulated:  true
//		isAbsolutePathDifferentFromPath:  false
//		doesAbsolutePathExist:  true
//		Directory Name:  testoverwrite
//		volumeName:  D:
//		isVolumePopulated:  true
//		========== File Info Data ==========
//		File Info IsDir():  true
//		File Info Name():  testoverwrite
//		File Info Size():  0
//		File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//		File Info Mode():  drwxrwxrwx
//		File Info     Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//		Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr DirMgr) New(pathStr string) (DirMgr, error) {

	ePrefix := "DirMgr.NewFromPathFileNameExtStr() "

	if len(pathStr) == 0 {
		return DirMgr{}, errors.New(ePrefix +
			"Error: Input parameter 'pathStr' is Zero Length!")
	}

	newDirMgr := DirMgr{}

	_, err := newDirMgr.SetDirMgr(pathStr)

	if err != nil {
		return DirMgr{}, fmt.Errorf(ePrefix+
			"Error returned by newDirMgr.SetDirMgr(pathStr) pathStr='%v' Error='%v'",
			pathStr, err.Error())
	}

	return newDirMgr, nil
}

// NewFromFileInfo - Returns a new DirMgr object based on two input parameters:
// 		- A directory path string
//		- An os.FileInfo object
func (dMgr DirMgr) NewFromFileInfo(pathStr string, info os.FileInfo) (DirMgr, error) {

	ePrefix := "DirMgr) NewFromFileInfo() "

	if len(pathStr) == 0 {
		return DirMgr{},
			errors.New(ePrefix +
				"Error: Input parameter 'pathStr' is Zero Length!")
	}

	newDirMgr := DirMgr{}

	err := newDirMgr.SetDirMgrWithFileInfo(pathStr, info)

	if err != nil {
		return DirMgr{},
			fmt.Errorf(ePrefix+"Error returned from '%v' ", err.Error())
	}

	return newDirMgr, nil
}

// SetDirMgr - Sets the DirMgr fields and path strings for the current DirMgr object.
//
// Input Parameters:
// =================
//
// pathStr string 	- A path string designating a path or directory.
// 										To reduce errors, the 'pathStr' should be terminated
//										with an appropriate path separator ('/' or '\')
//										Example 'pathStr': "C:\dirA\dirB\dirC\"
//
// Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//
//		isInitialized:  true
//		Original path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		PathIsPopuslated:  true
//		doesPathExist:  true
//		parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//		isParentPathPopulated:  true
//		relativePath:  testoverwrite
//		isRelativePathPopulated:  true
//		absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		isAbsolutePathPopulated:  true
//		isAbsolutePathDifferentFromPath:  false
//		doesAbsolutePathExist:  true
//		Directory Name:  testoverwrite
//		volumeName:  D:
//		isVolumePopulated:  true
//		========== File Info Data ==========
//		File Info IsDir():  true
//		File Info Name():  testoverwrite
//		File Info Size():  0
//		File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//		File Info Mode():  drwxrwxrwx
//		File Info     Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//		Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr *DirMgr) SetDirMgr(pathStr string) (isEmpty bool, err error) {
	ePrefix := "DirMgr.SetDirMgr() "

	dMgr.Empty()

	fh := FileHelper{}

	err = nil
	isEmpty = true

	if len(pathStr) == 0 {
		isEmpty = true
		err = errors.New(ePrefix +
			"Error: Input parameter 'pathStr' is a Zero length string!")
		return
	}

	trimmedPathStr := strings.TrimLeft(strings.TrimRight(pathStr, " "), " ")

	if len(trimmedPathStr) == 0 {
		err = errors.New(ePrefix +
			"Error: Trimmed Input Parameter 'pathStr' is a Zero length string!")
		isEmpty = true
		return
	}

	adjustedTrimmedPathStr := fh.AdjustPathSlash(trimmedPathStr)

	finalPathStr, isEmptyPath, err2 := fh.GetPathFromPathFileName(adjustedTrimmedPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error: INVALID PATH. fh.GetPathFromPathFileName(pathStr) "+
			"pathStr='%v'  Error='%v'", pathStr, err2.Error())
		isEmpty = isEmptyPath
		return
	}

	if isEmptyPath {
		isEmpty = true
		err = fmt.Errorf(ePrefix+
			"Error: INVALID PATH. 'pathStr' generated an Empty path! pathStr='%v' ",
			pathStr)
		return
	}

	if len(finalPathStr) == 0 {
		err = fmt.Errorf(ePrefix+
			"Error: path returned from fh.GetPathFromPathFileName(pathStr) is EMPTY! "+
			"pathStr='%v'", pathStr)
		isEmpty = true
		return
	}

	dMgr.originalPath = adjustedTrimmedPathStr

	dMgr.path = finalPathStr

	dMgr.isPathPopulated = true
	dMgr.DoesDirMgrPathExist()

	if dMgr.path == fp.VolumeName(dMgr.path) {

		dMgr.absolutePath = dMgr.path

	} else {

		dMgr.absolutePath, err2 = fh.MakeAbsolutePath(dMgr.path)

		if err2 != nil {
			dMgr.Empty()
			err = fmt.Errorf(ePrefix+
				"- fh.MakeAbsolutePath(dMgr.path) returned error. dMgr.path='%v' Error='%v'",
				dMgr.path, err2.Error())
			isEmpty = true
			return
		}

	}

	dMgr.absolutePath = fh.AdjustPathSlash(dMgr.absolutePath)

	dMgr.isAbsolutePathPopulated = true
	dMgr.DoesDirMgrAbsolutePathExist()

	strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
	lStr := len(strAry)
	idxStr := strAry[lStr-1]

	idx := strings.Index(dMgr.absolutePath, idxStr)
	dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

	dMgr.isParentPathPopulated = true

	if dMgr.isAbsolutePathPopulated && dMgr.isParentPathPopulated {

		dMgr.relativePath, err2 = fp.Rel(dMgr.parentPath, dMgr.absolutePath)

		if err2 != nil {
			dMgr.relativePath = ""
			dMgr.isParentPathPopulated = false
		} else {
			dMgr.isParentPathPopulated = true
			dMgr.isRelativePathPopulated = true

		}

	}

	if idxStr != "" {
		dMgr.directoryName = idxStr
	} else {
		dMgr.directoryName = dMgr.absolutePath
	}

	if dMgr.path != dMgr.absolutePath {
		dMgr.isAbsolutePathDifferentFromPath = true
	}

	var vn string
	if dMgr.isAbsolutePathPopulated {
		vn = fp.VolumeName(dMgr.absolutePath)
	} else if dMgr.isPathPopulated {
		vn = fp.VolumeName(dMgr.path)
	}

	if vn != "" {
		dMgr.isVolumePopulated = true
		dMgr.volumeName = vn
	}

	if dMgr.isAbsolutePathPopulated && dMgr.isPathPopulated {
		dMgr.isInitialized = true
		isEmpty = false
	} else {
		isEmpty = true
	}

	err = nil

	return
}

// SetDirMgrWithFileInfo - Sets the DirMgr fields and path strings for the current
// DirMgr object based on an input 'pathStr' parameter and an os.FileInfo input
// parameter ('info').
//
func (dMgr *DirMgr) SetDirMgrWithFileInfo(pathStr string, info os.FileInfo) error {
	ePrefix := "DirMgr.SetDirMgrWithFileInfo() "

	if len(pathStr) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'pathStr' is Zero Length!")
	}

	_, err := dMgr.SetDirMgr(pathStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from dMgr.SetDirMgr(pathStr). "+
			"pathStr='%v'  Error='%v'", pathStr, err.Error())
	}

	dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathStr, info)
	dMgr.directoryName = info.Name()

	return nil
}

// SubstituteBaseDir - Substitute baseDir segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
func (dMgr *DirMgr) SubstituteBaseDir(
	baseDir DirMgr,
	substituteBaseDir DirMgr) (newDMgr DirMgr, err error) {

	ePrefix := "DirMgr.SubstituteBaseDir() "

	newDMgr = DirMgr{}
	err = nil

	if !baseDir.isInitialized {
		err = errors.New(ePrefix + "Error: baseDir DirMgr is NOT Initialized!")
		return
	}

	if !substituteBaseDir.isInitialized {
		err = errors.New(ePrefix + "Error: substituteBaseDir DirMgr is NOT Initialized!")
		return
	}

	err2 := baseDir.IsDirMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error: baseDir DirMgr object is INVALID!")
		return
	}

	err2 = substituteBaseDir.IsDirMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error: substituteBaseDir DirMgr object is INVALID!")
		return
	}

	var trimmedRemainingPath string
	var newPath string

	if strings.HasPrefix(dMgr.path, baseDir.path) {

		trimmedRemainingPath = strings.TrimPrefix(dMgr.path, baseDir.path)

		lPath := len(trimmedRemainingPath)

		if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
			trimmedRemainingPath = trimmedRemainingPath[1:]
		}

		newPath = substituteBaseDir.GetPathWithSeparator() + trimmedRemainingPath

	} else if strings.HasPrefix(dMgr.absolutePath, baseDir.absolutePath) {

		trimmedRemainingPath = strings.TrimPrefix(dMgr.absolutePath, baseDir.absolutePath)

		lPath := len(trimmedRemainingPath)

		if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
			trimmedRemainingPath = trimmedRemainingPath[1:]
		}

		newPath = substituteBaseDir.GetAbsolutePathWithSeparator() + trimmedRemainingPath

	} else {
		err = fmt.Errorf(ePrefix+
			"Error: Could not locate baseDir.path or "+
			"baseDir.absolutePath in this dMgr. dMgr.path='%v' dMgr.absolutePath='%v'",
			dMgr.path, dMgr.absolutePath)
		return
	}

	newDMgr, err2 = DirMgr{}.New(newPath)

	if err2 != nil {
		newDMgr = DirMgr{}
		err = fmt.Errorf(ePrefix+
			"Error returned from DirMgr{}.NewFromPathFileNameExtStr(newPath). "+
			"newPath='%v'  Error='%v'", newPath, err2.Error())
		return
	}

	err = nil
	return
}

// executeFileOpsOnFoundFiles - This function is designed to work in conjunction
// with a walk directory function like FindWalkDirFiles. It will process
// files extracted from a 'Directory Walk' operation initiated by the
// 'filepath.Walk' method.
//
// Thereafter, file operations will be performed on files in the directory
// tree as specified by the 'dirOp' parameter.
//
func (dMgr *DirMgr) executeFileOpsOnFoundFiles(dirOp *DirTreeOp) func(string, os.FileInfo, error) error {
	return func(pathFile string, info os.FileInfo, erIn error) error {

		ePrefix := "DirMgr.executeFileOpsOnFoundFiles() "
		var err2 error

		if erIn != nil {
			err2 = fmt.Errorf(ePrefix+
				"Error returned from directory walk function. "+
				"pathFile= '%v' Error='%v'", pathFile, erIn.Error())
			dirOp.ErrReturns = append(dirOp.ErrReturns, err2.Error())
			return nil
		}

		if info.IsDir() {
			return nil
		}

		fh := FileHelper{}

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.
		isFoundFile, err := fh.FilterFileName(info, dirOp.FileSelectCriteria)

		if err != nil {

			err2 = fmt.Errorf(ePrefix+
				"Error returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria) "+
				"pathFile='%v' info.Name()='%v' Error='%v' ",
				pathFile, info.Name(), err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2.Error())
			return nil
		}

		if !isFoundFile {
			return nil
		}

		srcFileNameExt := info.Name()

		destDir, err := fh.SwapBasePath(
			dirOp.SourceBaseDir.GetAbsolutePath(),
			dirOp.TargetBaseDir.GetAbsolutePath(),
			pathFile)

		if err != nil {
			err2 = fmt.Errorf(ePrefix+
				"Error returned by fh.SwapBasePath(dirOp.SourceBaseDir, "+
				"dirOp.TargetBaseDir, pathFile). dirOp.SourceBaseDir='%v' "+
				"dirOp.TargetBaseDir='%v' pathFile='%v' Error='%v' ",
				dirOp.SourceBaseDir.GetAbsolutePath(),
				dirOp.TargetBaseDir.GetAbsolutePath(),
				pathFile,
				err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2.Error())
			return nil

		}

		fileOp, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
			pathFile, srcFileNameExt, destDir, srcFileNameExt)

		if err != nil {
			err2 = fmt.Errorf(ePrefix+
				"Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs() "+
				"pathFile='%v' srcFileNameExt='%v' destDir='%v' Error='%v' ",
				pathFile, srcFileNameExt, destDir, err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2.Error())
			return nil
		}

		maxOps := len(dirOp.FileOps)

		for i := 0; i < maxOps; i++ {

			err = fileOp.ExecuteFileOperation(dirOp.FileOps[i])

			if err != nil {
				err2 = fmt.Errorf(ePrefix+
					"Error returned by fileOp.ExecuteFileOperation(dirOp.FileOps[i]). "+
					"i='%v' FileOps='%v' Error='%v' ",
					i, dirOp.FileOps[i].String(), err.Error())

				dirOp.ErrReturns = append(dirOp.ErrReturns, err2.Error())

			}
		}

		return nil
	}

}
