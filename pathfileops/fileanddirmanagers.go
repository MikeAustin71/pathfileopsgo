package pathfileops

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
	"time"
)

/*
	This source code file contains types 'DirMgr' and
	'FileMgr'.

	The Source Repository for this source code file is :
		https://github.com/MikeAustin71/pathfilego.git

	Dependencies:
	-------------

	Types 'DirMgr' and 'FileMgr' depend on type, 'FileHellper'
	which is contained in source code file, 'filehelper.go' located
	in this directory.

	To use type 'DirMgr' or type 'FileMgr' you will need the following
	two source code files:

		1. pathfilego/003_filehelper/common/filehelper.go

		2. pathfilego/003_filehelper/common/fileanddirmanagers.go


*/

type DirMgrCollection struct {
	DirMgrs []DirMgr
}

// AddDirMgr - Adds a DirMgr object to the collection
func (dMgrs *DirMgrCollection) AddDirMgr(dMgr DirMgr) {
	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr.CopyOut())
}

func (dMgrs *DirMgrCollection) AddDirMgrByPathFile(pathFileName string) error {
	ePrefix := "DirMgrCollection.AddDirMgrByPathFile() "

	dMgr, err := DirMgr{}.New(pathFileName)

	if err != nil {
		s := ePrefix +
			"Error returned from DirMgr{}.New(pathFileName). pathFileName='%v' Error='%v'"
		return fmt.Errorf(s, pathFileName, err.Error())
	}

	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr)

	return nil
}

// AddFileMgrByFileInfo - Adds a File Manager object to the collection based on input from
// a directory path string and a os.FileInfo object.
func (dMgrs *DirMgrCollection) AddFileInfo(pathFile string, info os.FileInfo) error {

	ePrefix := "DirMgrCollection) AddFileMgrByFileInfo() "

	dMgr, err := DirMgr{}.NewFromFileInfo(pathFile, info)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error retrned from DirMgr{}."+
			"NewFromFileInfo(pathFile, info). pathFile='%v' info.Name()='%v'  Error='%v'",
			pathFile, info.Name(), err.Error())
	}

	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr)

	return nil
}

// AddDirMgrCollection - Adds another collection of File Manager (DirMgr)
// objects to the current collection.
func (dMgrs *DirMgrCollection) AddDirMgrCollection(dMgrs2 *DirMgrCollection) {

	lOmc2 := len(dMgrs2.DirMgrs)

	if lOmc2 == 0 {
		return
	}

	for i := 0; i < lOmc2; i++ {
		dMgrs.AddDirMgr(dMgrs2.DirMgrs[i].CopyOut())
	}

	return
}

// CopyOut - Returns an DirMgrCollection which is an
// exact duplicate of the current DirMgrCollection
func (dMgrs *DirMgrCollection) CopyOut() (DirMgrCollection, error) {

	ePrefix := "DirMgrCollection.CopyOut() "

	dMgrs2 := DirMgrCollection{}

	lOmc := len(dMgrs.DirMgrs)

	if lOmc == 0 {
		return DirMgrCollection{},
			errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	for i := 0; i < lOmc; i++ {
		dMgrs2.AddDirMgr(dMgrs.DirMgrs[i].CopyOut())
	}

	return dMgrs2, nil
}

// FindDirectories - searches through the DirMgrCollection to find
// DirMgr objects matching specified search criteria.
func (dMgrs *DirMgrCollection) FindDirectories(
	fileSelectionCriteria FileSelectionCriteria) (DirMgrCollection, error) {

	ePrefix := "DirMgrCollection.FindDirectories() "

	lDirCol := len(dMgrs.DirMgrs)

	if lDirCol == 0 {
		return DirMgrCollection{}, nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	dMgrs2 := DirMgrCollection{}

	for i := 0; i < lDirCol; i++ {
		dMgr := dMgrs.DirMgrs[i]

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

// GetArrayLength - returns the array length of the
// DirMgrCollection File Managers (FMgrs) array.
func (dMgrs *DirMgrCollection) GetArrayLength() int {
	return len(dMgrs.DirMgrs)
}

// PopLastDirMgr - Removes the last File Manager (DirMgr) object
// from the collections array, and returns it to the calling method.
func (dMgrs *DirMgrCollection) PopLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopLastDirMgr() "

	l1 := len(dMgrs.DirMgrs)

	if l1 == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	fmgr := dMgrs.DirMgrs[l1-1].CopyOut()

	dMgrs.DirMgrs = dMgrs.DirMgrs[0 : l1-1]

	return fmgr, nil
}

// PopFirstDirMgr - Removes the first OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (dMgrs *DirMgrCollection) PopFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopFirstDirMgr() "

	l1 := len(dMgrs.DirMgrs)

	if l1 == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	om := dMgrs.DirMgrs[0].CopyOut()

	dMgrs.DirMgrs = dMgrs.DirMgrs[1:l1]

	return om, nil
}

// PopDirMgrAtIndex - Returns a copy of the File Manager (DirMgr) object located
// at index, 'idx', in the DirMgrCollection array. As a 'Pop' method, the original
// DirMgr object is deleted from the DirMgrCollection array.
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(idx int) (DirMgr, error) {
	ePrefix := "DirMgrCollection.PopDirMgrAtIndex() "

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix+"Error: Input Parameter is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	lDirMgrs := len(dMgrs.DirMgrs)

	if idx >= lDirMgrs {
		return DirMgr{}, fmt.Errorf(ePrefix+"Error: Input Parameter is greater than the "+
			"length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, lDirMgrs)
	}

	if idx == 0 {
		return dMgrs.PopFirstDirMgr()
	}

	if idx == lDirMgrs-1 {
		return dMgrs.PopLastDirMgr()
	}

	dMgr := dMgrs.DirMgrs[idx].CopyOut()

	dirCol2 := DirMgrCollection{}

	for i := 0; i < lDirMgrs; i++ {

		if i != idx {
			dirCol2.DirMgrs = append(dirCol2.DirMgrs, dMgrs.DirMgrs[i].CopyOut())
		}

	}

	dMgrs.DirMgrs = dirCol2.DirMgrs

	return dMgr, nil
}

// PeekFirstDirMgr - Returns the first element from the
// DirMgrCollection, but does NOT remove
// it from the OpsMessages array.
func (dMgrs *DirMgrCollection) PeekFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekFirstDirMgr() "

	if len(dMgrs.DirMgrs) == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	return dMgrs.DirMgrs[0].CopyOut(), nil
}

// PeekLastDirMgr - Returns the last element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (dMgrs *DirMgrCollection) PeekLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekLastDirMgr()"

	l1 := len(dMgrs.DirMgrs)

	if l1 == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	return dMgrs.DirMgrs[l1-1].CopyOut(), nil
}

// PeekDirMgrAtIndex - Returns a copy of the File Manager (DirMgr) object located
// at array index 'idx' in the DirMgrCollection. This is a 'Peek' method and the
// original DirMgr object is not deleted from the DirMgrCollection array.
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(idx int) (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekDirMgrAtIndex() "

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= len(dMgrs.DirMgrs) {
		return DirMgr{}, fmt.Errorf(ePrefix+"Error: Input Parameter is greater than the "+
			"length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ",
			idx, len(dMgrs.DirMgrs))

	}

	return dMgrs.DirMgrs[idx].CopyOut(), nil
}

// DirMgr - This structure and associated methods
// are used to manage a specific directory.
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

// DeleteDirFiles - deletes files in the current directory based
// on input parameter, 'formatSpec'. This specification will be
// used to match files in the current directory. If a match is
// found that file will be deleted.
//
// Examples:
//
//   formatSpec
//     *.*				Delete all files in directory
//     *.txt			Delete all files in directory with
//     							file extension = '.txt'
//     wilson.*		Delete all files in directory with
//                  file name = 'wilson'
//
func (dMgr *DirMgr) DeleteDirFiles(formatSpec string) error {

	ePrefix := "DirMgr.DeleteDirFiles() "
	target := dMgr.GetAbsolutePathWithSeparator() + formatSpec

	err := os.Remove(target)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by os.Remove(target). "+
			"target='%v' Error='%v' ", target, err.Error())
	}

	return nil
}

// DeleteDirContents - Deletes all the files in the current
// directory and ONLY the current directory.
//
// Files in subdirectories are NOT deleted.
//
// Reference:
// https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
//
func (dMgr *DirMgr) DeleteDirContents() error {

	ePrefix := "DirMgr.DeleteDirContents() "

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

	names, err := dir.Readdirnames(-1)

	if err != nil {
		_ = dir.Close()
		return fmt.Errorf(ePrefix+
			"Error returned by dir.Readdirnames(-1). "+
			"dMgr.absolutePath='%v' Error='%v' ",
			dMgr.absolutePath, err.Error())
	}

	for _, name := range names {

		err = os.RemoveAll(fp.Join(dMgr.absolutePath, name))

		if err != nil {
			_ = dir.Close()
			return fmt.Errorf(ePrefix+
				"Error returned by dir.Readdirnames(-1). "+
				"dMgr.absolutePath='%v' fileName='%v' Error='%v' ",
				dMgr.absolutePath, name, err.Error())
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
//									DirMgrs          	[]DirMgr
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
		dMgr.actualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(dMgr.absolutePath, info)
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
		dMgr.actualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(dMgr.path, info)
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

// EqualPaths - Compares two DirMgr objects to determine
// if their paths are equal.
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

	if dMgr.isInitialized != dMgr2.isInitialized {
		return false
	}

	if dMgr.absolutePath == dMgr2.absolutePath &&
		dMgr.path == dMgr2.path {
		return true
	}

	return false
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
//
// Return Value:
// =============
//
//	DirectoryTreeInfo structure	-
//					type DirectoryTreeInfo struct {
//						StartPath            	string								// The starting path or directory for the file search
//						DirMgrs          	[]DirMgr									// DirMgrs found during directory tree search
//						FoundFiles           	[]FileWalkInfo				// Found Files matching file selection criteria
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
	dMgr.actualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(dMgr.absolutePath, info)

	return dMgr.actualDirFileInfo.CopyOut(), nil
}

// GetOriginalPath - Returns the original path used to initialize
// this Directory Manager instance.
//
func (dMgr *DirMgr) GetOriginalPath() string {
	return dMgr.originalPath
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

// New - Returns a new DirMgr object and populates the
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

	ePrefix := "DirMgr.New() "

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

	dMgr.actualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(pathStr, info)
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
			"Error returned from DirMgr{}.New(newPath). "+
			"newPath='%v'  Error='%v'", newPath, err2.Error())
		return
	}

	err = nil
	return
}

// FileMgrCollection - A collection of FileMgr objects
type FileMgrCollection struct {
	FMgrs []FileMgr
}

// AddFileMgr - Adds a FileMgr object to the collection
func (fMgrs *FileMgrCollection) AddFileMgr(fMgr FileMgr) {
	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr.CopyOut())
}

// AddFileMgrByDirFileNameExt - Add a new File Manager using
// input parameters 'directory' and 'pathFileNameExt'.
func (fMgrs *FileMgrCollection) AddFileMgrByDirFileNameExt(
	directory DirMgr,
	pathFileNameExt string) error {

	ePrefix := "FileMgrCollection.AddFileMgrByDirFileNameExt() "

	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(directory, pathFileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr)

	return nil
}

// AddFileMgrByPathFile - Add a new File Manager based on
// input parameter pathFileName
func (fMgrs *FileMgrCollection) AddFileMgrByPathFile(
	pathFileName string) error {

	ePrefix := "FileMgrCollection.AddFileMgrByPathFile() "

	fMgr, err := FileMgr{}.New(pathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from FileMgr{}.New(pathFileName). "+
			"pathFileName='%v' Error='%v'", pathFileName, err.Error())
	}

	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr)

	return nil
}

// AddFileMgrByFileInfo - Adds a File Manager object to the collection based on input from
// a directory path string and a os.FileInfo object.
func (fMgrs *FileMgrCollection) AddFileMgrByFileInfo(pathFile string, info os.FileInfo) error {

	ePrefix := "FileMgrCollection) AddFileMgrByFileInfo() "

	fMgr, err := FileMgr{}.NewFromFileInfo(pathFile, info)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error retrned from FileMgr{}.NewFromFileInfo(pathFile, info). "+
			"pathFile='%v' info.Name()='%v'  Error='%v'", pathFile, info.Name(), err.Error())
	}

	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr)

	return nil
}

// AddFileMgrCollection - Adds another collection of File Manager (FileMgr)
// objects to the current collection.
func (fMgrs *FileMgrCollection) AddFileMgrCollection(fMgrs2 *FileMgrCollection) {

	lOmc2 := len(fMgrs2.FMgrs)

	if lOmc2 == 0 {
		return
	}

	for i := 0; i < lOmc2; i++ {
		fMgrs.AddFileMgr(fMgrs2.FMgrs[i].CopyOut())
	}

	return
}

// CopyFilesToDir - Copies all the files in the File Manager Collection to
// the specified target directory.
//
func (fMgrs *FileMgrCollection) CopyFilesToDir(targetDirectory DirMgr) error {

	ePrefix := "FileMgrCollection.CopyFilesToDir() "
	maxLen := len(fMgrs.FMgrs)

	if maxLen == 0 {
		return errors.New(ePrefix + "ERROR - Collection contains ZERO File Managers!")
	}

	for i := 0; i < maxLen; i++ {
		err := fMgrs.FMgrs[i].CopyFileToDir(targetDirectory)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Copy Failure on index='%v' file='%v'. Error='%v'",
				i, fMgrs.FMgrs[i].absolutePathFileName, err.Error())
		}

	}

	return nil
}

// CopyOut - Returns an FileMgrCollection which is an
// exact duplicate of the current FileMgrCollection
func (fMgrs *FileMgrCollection) CopyOut() (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.CopyOut() "

	fMgrs2 := FileMgrCollection{}

	lOmc := len(fMgrs.FMgrs)

	if lOmc == 0 {
		return FileMgrCollection{},
			errors.New(ePrefix +
				"Error: Empty FileMgrCollection. No messages available!")
	}

	for i := 0; i < lOmc; i++ {
		fMgrs2.AddFileMgr(fMgrs.FMgrs[i].CopyOut())
	}

	return fMgrs2, nil
}

// FindFiles - Searches the current FileMgrCollection and returns a new
// FileMgrCollection containing FileMgr objects which match the specified
// search criteria.
//
func (fMgrs *FileMgrCollection) FindFiles(
	fileSelectionCriteria FileSelectionCriteria) (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.FindFiles() "

	lDirCol := len(fMgrs.FMgrs)

	if lDirCol == 0 {
		return FileMgrCollection{}, nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	fMgrs2 := FileMgrCollection{}

	for i := 0; i < lDirCol; i++ {
		fMgr := fMgrs.FMgrs[i]

		if fMgr.actualFileInfo.IsFInfoInitialized {

			isMatchedFile, err = fh.FilterFileName(fMgr.actualFileInfo, fileSelectionCriteria)

			if err != nil {
				return FileMgrCollection{},
					fmt.Errorf(ePrefix+
						"Error returned by "+
						"fh.FilterFileName(fMgr.actualFileInfo, fileSelectionCriteria) "+
						"fMgr.actualFileInfo.Name()='%v'  Error='%v'",
						fMgr.actualFileInfo.Name(), err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(fMgr.fileNameExt)

			isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)

			if err != nil {
				return FileMgrCollection{}, fmt.Errorf(ePrefix+
					"Error returned by fh.FilterFileName(fip, fileSelectionCriteria) "+
					"fip.Name()='%v'  Error='%v'", fip.Name(), err.Error())
			}

		}

		if isMatchedFile && err == nil {
			fMgrs2.AddFileMgr(fMgr)
		}

	}

	return fMgrs2, nil
}

// GetArrayLength - returns the array length of the
// FileMgrCollection File Managers (FMgrs) array.
func (fMgrs *FileMgrCollection) GetArrayLength() int {
	return len(fMgrs.FMgrs)
}

// Returns a deep copy of the File Manager at index 'idx'.
//
func (fMgrs *FileMgrCollection) GetFileMgrAtIdx(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.GetFileMgrAtIdx() "

	if idx < 0 {
		return FileMgr{},
			errors.New(ePrefix + "ERROR: 'idx' is LESS THAN ZERO!")
	}

	if idx > len(fMgrs.FMgrs)-1 {
		return FileMgr{},
			fmt.Errorf(ePrefix+"ERROR: 'idx' exceeds FileMgrCollection Array Boundary. \n"+
				"Maximum Index='%v' Input parameter 'idx' = '%v' ", len(fMgrs.FMgrs)-1, idx)
	}

	return fMgrs.FMgrs[idx].CopyOut(), nil
}

// PopLastFMgr - Removes the last File Manager (FileMgr) object
// from the collections array, and returns it to the calling method.
func (fMgrs *FileMgrCollection) PopLastFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopLastFMgr() "

	l1 := len(fMgrs.FMgrs)

	if l1 == 0 {
		return FileMgr{}, errors.New(ePrefix +
			"Error: Empty FileMgrCollection. No messages available!")
	}

	fmgr := fMgrs.FMgrs[l1-1].CopyOut()

	fMgrs.FMgrs = fMgrs.FMgrs[0 : l1-1]

	return fmgr, nil
}

// PopFirstFMgr - Removes the first OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (fMgrs *FileMgrCollection) PopFirstFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFirstFMgr() "

	l1 := len(fMgrs.FMgrs)

	if l1 == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: Empty FileMgrCollection. No messages available!")
	}

	om := fMgrs.FMgrs[0].CopyOut()

	fMgrs.FMgrs = fMgrs.FMgrs[1:l1]

	return om, nil
}

// PopFMgrAtIndex - Returns a copy of the File Manager (FileMgr) object located
// at index, 'idx', in the FileMgrCollection array. As a 'Pop' method, the original
// FileMgr object is deleted from the FileMgrCollection array.
func (fMgrs *FileMgrCollection) PopFMgrAtIndex(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFMgrAtIndex() "

	if idx < 0 {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	lFMgrs := len(fMgrs.FMgrs)

	if idx >= lFMgrs {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is greater than the length of the collection index. "+
			"Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, lFMgrs)
	}

	if idx == 0 {
		return fMgrs.PopFirstFMgr()
	}

	if idx == lFMgrs-1 {
		return fMgrs.PopLastFMgr()
	}

	fmgr := fMgrs.FMgrs[idx].CopyOut()

	fCol := FileMgrCollection{}

	for i := 0; i < lFMgrs; i++ {

		if i != idx {
			fCol.FMgrs = append(fCol.FMgrs, fMgrs.FMgrs[i].CopyOut())
		}

	}

	fMgrs.FMgrs = fCol.FMgrs

	return fmgr, nil
}

// PeekFirstFMgr - Returns the first element from the
// FileMgrCollection, but does NOT remove
// it from the OpsMessages array.
func (fMgrs *FileMgrCollection) PeekFirstFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekFirstFMgr() "

	if len(fMgrs.FMgrs) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: Empty FileMgrCollection. No messages available!")
	}

	return fMgrs.FMgrs[0].CopyOut(), nil
}

// PeekLastFMgr - Returns the last element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (fMgrs *FileMgrCollection) PeekLastFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekLastFMgr()"

	l1 := len(fMgrs.FMgrs)

	if l1 == 0 {
		return FileMgr{}, errors.New(ePrefix +
			"Error: Empty FileMgrCollection. No messages available!")
	}

	return fMgrs.FMgrs[l1-1].CopyOut(), nil
}

// PeekFMgrAtIndex - Returns a copy of the File Manager (FileMgr) object located
// at array index 'idx' in the FileMgrCollection. This is a 'Peek' method and the
// original FileMgr object is not deleted from the FileMgrCollection array.
func (fMgrs *FileMgrCollection) PeekFMgrAtIndex(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekFMgrAtIndex() "

	if idx < 0 {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= len(fMgrs.FMgrs) {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is greater than the length of the collection index. "+
			"Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, len(fMgrs.FMgrs))
	}

	return fMgrs.FMgrs[idx].CopyOut(), nil
}

// FileMgr - This structure and associated methods
// are used to manage a specific file.
type FileMgr struct {
	isInitialized                   bool
	originalPathFileName            string
	dMgr                            DirMgr
	absolutePathFileName            string
	isAbsolutePathFileNamePopulated bool
	doesAbsolutePathFileNameExist   bool
	fileName                        string
	isFileNamePopulated             bool
	fileExt                         string
	isFileExtPopulated              bool
	fileNameExt                     string
	isFileNameExtPopulated          bool
	filePtr                         *os.File
	isFilePtrOpen                   bool
	actualFileInfo                  FileInfoPlus
}

// CopyFileToDir - Copies the file identified by the current File Manager
// (FileMgr) object to another directory specified by input parameter
// 'dir', a 'DirMgr' object.
func (fMgr *FileMgr) CopyFileToDir(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDir() "
	err := dir.IsDirMgrValid("")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: Input parmater dir is INVALID! Error='%v'",
			err.Error())
	}

	newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
			"fMgr.fileNameExt) dir.absolutePath='%v'  fMgr.fileNameExt='%v'  Error='%v'",
			dir.absolutePath, fMgr.fileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgr(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgr(&newFMgr) "+
			"newFMgr.absolutePathFileName='%v'  Error='%v'",
			newFMgr.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileMgr - Copies the file represented by the current File Manager object
// to a location specified by an input parameter File Manager.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
func (fMgr *FileMgr) CopyFileMgr(fMgrDest *FileMgr) error {

	ePrefix := "FileMgr.CopyFileMgr() "

	err := fMgr.IsFileMgrValid("")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: This FileMgr object is INVALID! Error='%v'", err.Error())
	}

	if fMgrDest == nil {
		return errors.New(ePrefix +
			"Error: Input parameter fMgrDest is a nil pointer!")
	}

	err = fMgrDest.IsFileMgrValid("")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: The Destination FileMgr object is INVALID! Error='%v'",
			err.Error())
	}

	if !fMgr.doesAbsolutePathFileNameExist {
		return fmt.Errorf(ePrefix+"Error: Source file-  %v  DOES NOT EXIST!",
			fMgr.absolutePathFileName)
	}

	if !fMgr.actualFileInfo.Mode().IsRegular() {
		return fmt.Errorf(ePrefix+
			"Error: Source file is a Non-Regular "+
			"File and cannot be compied. File='%v'", fMgr.absolutePathFileName)
	}

	err = fMgrDest.dMgr.MakeDir()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Atempted creation of destination directory FAILED! Error= '%v'",
			err.Error())
	}

	destFileExists, err := fMgrDest.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fMgrDest.DoesThisFileExist(). Error= '%v'", err.Error())
	}

	if destFileExists && !fMgrDest.actualFileInfo.Mode().IsRegular() {
		return fmt.Errorf(ePrefix+
			"Error: Destination file exists and it is NOT a 'regular' file. "+
			"Copy operation aborted! Destination File='%v' ", fMgrDest.absolutePathFileName)
	}

	// Source and destination files are the same file.
	// No need to copy over source. Just return - Copy
	// operation complete.
	if destFileExists && os.SameFile(fMgr.actualFileInfo, fMgrDest.actualFileInfo) {
		return nil
	}

	err = fMgrDest.dMgr.MakeDir()

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	// See Reference:
	// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

	fh := FileHelper{}

	err = fh.CopyFileByIoByLink(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.CopyFileByLinkByIo(fMgr.absolutePathFileName, "+
			"fMgrDest.absolutePathFileName) fMgr.absolutePathFileName='%v' "+
			"fMgrDest.absolutePathFileName='%v' Error='%v'",
			fMgr.absolutePathFileName, fMgrDest.absolutePathFileName, err.Error())
	}

	destFileExists, err = fMgrDest.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgrDest.DoesThisFileExist(). "+
			"fMgrDest.absolutePathFileName='%v'  Error='%v'",
			fMgrDest.absolutePathFileName, err.Error())
	}

	if !destFileExists {
		return fmt.Errorf(ePrefix+
			"Error: After attempted file copy to destination file. Destination "+
			"file does NOT exist! fMgrDest.absolutePathFileName='%v'",
			fMgrDest.absolutePathFileName)
	}

	return nil
}

// CopyFileStr - Copies file from fMgr.absolutePathFileName to
// to destination path & File Name.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileStr(dstPathFileName string) error {

	ePrefix := "FileMgr.CopyFileStr "

	fMgrDest, err := FileMgr{}.New(dstPathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by FileMgr{}.New(dstPathFileName). "+
			"dstPathFileName='%v' Error='%v'", dstPathFileName, err.Error())
	}

	err = fMgr.CopyFileMgr(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from fMgr.CopyFileMgr(&fMgrDest) "+
			"fMgrDest.absolutePathFileName='%v'  Error='%v'", fMgrDest.absolutePathFileName, err.Error())
	}

	return nil
}

// CloseFile - This method will call the Close()
// method on the current file pointer, FileHelper.filePtr
func (fMgr *FileMgr) CloseFile() error {

	ePrefix := "FileMgr.CloseFile() "
	var err error

	if fMgr.filePtr == nil {
		fMgr.isFilePtrOpen = false
		return nil
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.FlushBytesToDisk().  "+
			"Error='%v'", err.Error())

	}

	err = fMgr.filePtr.Close()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Received Error from fMgr.filePtr.Close(). "+
			"fMgr.absolutePathFileName= '%v' ", fMgr.absolutePathFileName)
	}

	fMgr.isFilePtrOpen = false
	fMgr.filePtr = nil

	return nil
}

// CopyIn - Copies data from an incoming FileMgr object
// into the current FileMgr object.
func (fMgr *FileMgr) CopyIn(fmgr2 *FileMgr) {

	if fmgr2 == nil {
		panic("FileMgr.CopyIn() - Input parameter is a nil pointer!")
	}

	fMgr.isInitialized = fmgr2.isInitialized
	fMgr.dMgr.CopyIn(&fmgr2.dMgr)
	fMgr.originalPathFileName = fmgr2.originalPathFileName
	fMgr.absolutePathFileName = fmgr2.absolutePathFileName
	fMgr.isAbsolutePathFileNamePopulated = fmgr2.isAbsolutePathFileNamePopulated
	fMgr.doesAbsolutePathFileNameExist = fmgr2.doesAbsolutePathFileNameExist
	fMgr.fileName = fmgr2.fileName
	fMgr.isFileNamePopulated = fmgr2.isFileNamePopulated
	fMgr.fileExt = fmgr2.fileExt
	fMgr.isFileExtPopulated = fmgr2.isFileExtPopulated
	fMgr.fileNameExt = fmgr2.fileNameExt
	fMgr.isFileNameExtPopulated = fmgr2.isFileNameExtPopulated
	fMgr.filePtr = fmgr2.filePtr
	fMgr.isFilePtrOpen = fmgr2.isFilePtrOpen
	fMgr.actualFileInfo = fmgr2.actualFileInfo.CopyOut()

	return
}

// CopyOut - Duplicates the file information in the current
// FileMgr object and returns it as a new FileMgr object.
func (fMgr *FileMgr) CopyOut() FileMgr {

	fmgr2 := FileMgr{}

	fmgr2.isInitialized = fMgr.isInitialized
	fmgr2.dMgr = fMgr.dMgr.CopyOut()
	fmgr2.originalPathFileName = fMgr.originalPathFileName
	fmgr2.absolutePathFileName = fMgr.absolutePathFileName
	fmgr2.isAbsolutePathFileNamePopulated = fMgr.isAbsolutePathFileNamePopulated
	fmgr2.doesAbsolutePathFileNameExist = fMgr.doesAbsolutePathFileNameExist
	fmgr2.fileName = fMgr.fileName
	fmgr2.isFileNamePopulated = fMgr.isFileNamePopulated
	fmgr2.fileExt = fMgr.fileExt
	fmgr2.isFileExtPopulated = fMgr.isFileExtPopulated
	fmgr2.fileNameExt = fMgr.fileNameExt
	fmgr2.isFileNameExtPopulated = fMgr.isFileNameExtPopulated
	fmgr2.filePtr = fMgr.filePtr
	fmgr2.isFilePtrOpen = fMgr.isFilePtrOpen
	fmgr2.actualFileInfo = fMgr.actualFileInfo.CopyOut()

	return fmgr2
}

// Equal - Compares a second FileHelper data structure
// to the current FileHelper data structure and returns
// a boolean value indicating whether they are equal
// in all respects.
func (fMgr *FileMgr) Equal(fmgr2 *FileMgr) bool {

	if fMgr.isInitialized != fmgr2.isInitialized ||
		fMgr.originalPathFileName != fmgr2.originalPathFileName ||
		fMgr.isAbsolutePathFileNamePopulated != fmgr2.isAbsolutePathFileNamePopulated ||
		fMgr.doesAbsolutePathFileNameExist != fmgr2.doesAbsolutePathFileNameExist ||
		fMgr.absolutePathFileName != fmgr2.absolutePathFileName ||
		fMgr.fileName != fmgr2.fileName ||
		fMgr.isFileNamePopulated != fmgr2.isFileNamePopulated ||
		fMgr.fileExt != fmgr2.fileExt ||
		fMgr.isFileExtPopulated != fmgr2.isFileExtPopulated ||
		fMgr.fileNameExt != fmgr2.fileNameExt ||
		fMgr.isFileNameExtPopulated != fmgr2.isFileNameExtPopulated ||
		fMgr.filePtr != fmgr2.filePtr ||
		fMgr.isFilePtrOpen != fmgr2.isFilePtrOpen {

		return false
	}

	if !fMgr.dMgr.Equal(&fmgr2.dMgr) {
		return false
	}

	if !fMgr.actualFileInfo.Equal(&fmgr2.actualFileInfo) {
		return false
	}

	return true
}

// Empty - resets all data fields in the FileMgr structure to
// their uninitialized or zero state.
func (fMgr *FileMgr) Empty() {
	fMgr.isInitialized = false
	fMgr.dMgr = DirMgr{}
	fMgr.originalPathFileName = ""
	fMgr.absolutePathFileName = ""
	fMgr.isAbsolutePathFileNamePopulated = false
	fMgr.doesAbsolutePathFileNameExist = false
	fMgr.fileName = ""
	fMgr.isFileNamePopulated = false
	fMgr.fileExt = ""
	fMgr.isFileExtPopulated = false
	fMgr.fileNameExt = ""
	fMgr.isFileNameExtPopulated = false
	fMgr.filePtr = nil
	fMgr.isFilePtrOpen = false
	fMgr.actualFileInfo = FileInfoPlus{}

}

// CreateDirAndFile - Performs two operations:
// This is a Wrapper function for os.Create - Create a file.
//
//  If the home directory does not currently exist, this method
//  will first create the directory tree, before creating the new
//	file.
//
// If the file previously exists, it will be truncated.
//
// Note that if this method successfully creates the file, a
// File Pointer (*File) will be stored in the FileMgr field
// fMgr.filePtr. Be sure to close the File Pointer when
// finished with it. See FileMgr.CloseFile().
//
func (fMgr *FileMgr) CreateDirAndFile() error {

	ePrefix := "FileMgr:CreateDirAndFile() "
	var err error

	err = fMgr.IsFileMgrValid(ePrefix)

	if err != nil {
		return err
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(fMgr.dMgr.absolutePath) {
		// Directory does NOT exist, create it!

		err := fh.MakeDirAll(fMgr.dMgr.absolutePath)

		if err != nil {
			return fmt.Errorf(ePrefix+"Errors from FileHelper:"+
				"MakeDirAll(fMgr.dMgr.absolutePath). fMgr.dMgr.absolutePath='%v'  Error='%v' ",
				fMgr.dMgr.absolutePath, err.Error())
		}

		fMgr.dMgr.doesAbsolutePathExist = true

	} else {

		fMgr.dMgr.doesAbsolutePathExist = true

	}

	fMgr.filePtr, err = os.Create(fMgr.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error creating File. Error returned from os.Create("+
			"fMgr.absolutePathFileName). fMgr.absolutePathFileName='%v' Error='%v' ",
			fMgr.absolutePathFileName, err.Error())
	}

	fMgr.isFilePtrOpen = true

	return nil

}

// CreateFile - Creates the File identified by FileMgr.absolutePathFileName.
// If the directory in the path file name designation does not exist, this
// method will throw an error.
//
// See Method CreateDirAndFile() which will create both the directory and the file
// as required.
//
// Note that if the file is actually created, the returned file pointer (*File)
// is stored in the FileMgr field, fMgr.filePtr. Be sure to 'close' the File Pointer
// when finished with it. See FileMgr.CloseFile()
//
func (fMgr *FileMgr) CreateFile() error {

	ePrefix := "FileMgr:CreateFile() Error - "

	fh := FileHelper{}

	if !fMgr.isInitialized {
		return errors.New(ePrefix + " FileMgr is NOT Initialized!")
	}

	if !fMgr.dMgr.isAbsolutePathPopulated {
		return errors.New(ePrefix + " FileMgrDMgr.isAbsolutePathPopulated is NOT populated!")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return errors.New(ePrefix + " FileMgr.absolutePathFileName is EMPTY!")
	}

	if !fh.DoesFileExist(fMgr.dMgr.absolutePath) {
		fMgr.dMgr.doesAbsolutePathExist = false
		return fmt.Errorf(ePrefix+
			"FileMgr.dMgr.absolutePath Does NOT exist! Create the path. "+
			"FileMgr.dMgr.absolutePath='%v'", fMgr.dMgr.absolutePath)
	}

	var err error

	fMgr.filePtr, err = os.Create(fMgr.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error creating File. Error returned from "+
			"os.Create(fMgr.absolutePathFileName). fMgr.absolutePathFileName='%v' Error='%v' ",
			fMgr.absolutePathFileName, err.Error())
	}

	fMgr.isFilePtrOpen = true

	return nil

}

// DeleteThisFile - Deletes the file identified by FileMgr.absolutePathFileName
// in the current FileHelper structure.
func (fMgr *FileMgr) DeleteThisFile() error {

	ePrefix := "FileMgr.DeleteThisFile() "

	err := fMgr.IsFileMgrValid("")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: This FileMgr object is INVALID!  Error='%v'", err.Error())
	}

	if fMgr.filePtr != nil {

		err = fMgr.filePtr.Close()

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error from fMgr.filePtr.Close()!  Error='%v'", err.Error())
		}
	}

	fMgr.isFilePtrOpen = false

	err = os.Remove(fMgr.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"- os.Remove(fMgr.absolutePathFileName) "+
			"returned an error. absolutePathFileName='%v'   Error='%v'",
			fMgr.absolutePathFileName, err.Error())
	}

	fileExists, err := fMgr.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fMgr.DoesThisFileExist() "+
			"fMgr.absolutePathFileName='%v'  Error='%v'", fMgr.absolutePathFileName, err.Error())
	}

	if fileExists {
		return fmt.Errorf(ePrefix+
			"Error: Attempted file deletion FAILED!. "+
			"File still exists. fMgr.absolutePathFileName='%v'", fMgr.absolutePathFileName)
	}

	fMgr.actualFileInfo = FileInfoPlus{}

	return nil
}

// DoesThisFileExist - Returns a boolean value
// designated whether the file specified by the
// current FileMgr.absolutePathFileName field
// exists.
func (fMgr *FileMgr) DoesThisFileExist() (bool, error) {

	ePrefix := "FileMgr.DoesThisFileExist() "

	if !fMgr.isInitialized {
		return false,
			errors.New(ePrefix +
				"Error: The File Manager data structure has NOT been initialized.")
	}

	if fMgr.isAbsolutePathFileNamePopulated == false {
		return false,
			errors.New(ePrefix + " Error: absolutePathFileName is NOT POPULATED!")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return false, errors.New(ePrefix + " Error: absolutePathFileName is EMPTY!")
	}

	info, err := os.Stat(fMgr.absolutePathFileName)

	if err != nil {
		fMgr.actualFileInfo = FileInfoPlus{}
		fMgr.doesAbsolutePathFileNameExist = false
	} else {
		fMgr.isAbsolutePathFileNamePopulated = true
		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.actualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.absolutePathFileName, info)
	}

	fMgr.dMgr.DoesDirMgrPathExist()
	fMgr.dMgr.DoesDirMgrPathExist()

	return fMgr.doesAbsolutePathFileNameExist, nil
}

// FlushBytesToDisk - After Writing bytes to a file, use this
// method to commit the contents of the current file to
// stable storage.
func (fMgr *FileMgr) FlushBytesToDisk() error {

	ePrefix := "FileMgr.FlushBytesToDisk() "

	var err error

	if fMgr.isFilePtrOpen && fMgr.filePtr != nil {

		err = fMgr.filePtr.Sync()

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from fMgr.filePtr.Sync()")
		}

	}

	return nil
}

// GetAbsolutePathFileName - Returns the absolute path,
// file name and file extension for the current File Manager
// instance.
//
func (fMgr *FileMgr) GetAbsolutePathFileName() string {
	return fMgr.absolutePathFileName
}

// GetDirMgr - returns a deep copy of the Directory
// Manager for this FileMgr instance.
func (fMgr *FileMgr) GetDirMgr() DirMgr {
	return fMgr.dMgr.CopyOut()
}

// GetFileExt() - returns a string containing the
// File Extension for this File Manager instance.
func (fMgr *FileMgr) GetFileExt() string {
	return fMgr.fileExt
}

// GetFileInfoPlus - Returns a FileInfoPlus instance containing
// os.FileInfo and other data on the current FileManager instance.
//
func (fMgr *FileMgr) GetFileInfoPlus() (FileInfoPlus, error) {
	ePrefix := "FileMgr.GetFileInfoPlus() "

	err := fMgr.IsFileMgrValid(ePrefix)

	if err != nil {
		return FileInfoPlus{}, err
	}

	if !fMgr.isInitialized {
		return FileInfoPlus{},
			errors.New(ePrefix +
				"Error: This data structure is NOT initialized.")
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return FileInfoPlus{},
			errors.New(ePrefix +
				"Error: absolutePathFileName is NOT populated/initialized.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return FileInfoPlus{},
			errors.New(ePrefix +
				"Error: absolutePathFileName is EMPTY!")
	}

	info, err := os.Stat(fMgr.absolutePathFileName)

	if err != nil {
		return FileInfoPlus{},
			fmt.Errorf(ePrefix+"Error returned by "+
				"os.Stat(fMgr.absolutePathFileName). fMgr.absolutePathFileName='%v'  Error='%v'",
				fMgr.absolutePathFileName, err.Error())
	}

	fMgr.actualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.dMgr.absolutePath, info)

	return fMgr.actualFileInfo.CopyOut(), nil
}

// GetFileInfo - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on the specific file identified
// by FileMgr.absolutePathFileName. If the file does NOT exist,
// an error will be triggered.
//
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }
func (fMgr *FileMgr) GetFileInfo() (os.FileInfo, error) {

	ePrefix := "FileMgr.DeleteThisFile() "

	if !fMgr.isInitialized {
		return nil,
			errors.New(ePrefix +
				"Error: This data structure is NOT initialized.")
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return nil,
			errors.New(ePrefix +
				"Error: absolutePathFileName is NOT populated/initialized.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return nil,
			errors.New(ePrefix +
				"Error: absolutePathFileName is EMPTY!")
	}

	info, err := os.Stat(fMgr.absolutePathFileName)

	if err != nil {
		return nil,
			fmt.Errorf(ePrefix+"Error returned by "+
				"os.Stat(fMgr.absolutePathFileName). fMgr.absolutePathFileName='%v'  Error='%v'",
				fMgr.absolutePathFileName, err.Error())
	}

	fMgr.actualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.dMgr.absolutePath, info)

	return info, nil
}

// GetFileName - returns the file name for this
// File Manager.
//
func (fMgr *FileMgr) GetFileName() string {
	return fMgr.fileName
}

// GetFileNameExt - Returns a string containing the
// combination of file name and file extension configured
// for this File Manager instance
//
func (fMgr *FileMgr) GetFileNameExt() string {
	return fMgr.fileNameExt
}

// GetOriginalPathFileName - Returns the path and file name
// used originally to configure this File Manager object.
//
func (fMgr *FileMgr) GetOriginalPathFileName() string {

	return fMgr.originalPathFileName

}

// GetFilePtr - will return the internal *os.File pointer
// for this File Manager instance. Depending on circumstances,
// this pointer may be nil.
//
func (fMgr *FileMgr) GetFilePtr() *os.File {
	return fMgr.filePtr
}

// IsAbsolutePathFileNamePopulated - Returns a boolean value
// indicating whether absolute path and file name is
// initialized and populated.
//
func (fMgr *FileMgr) IsAbsolutePathFileNamePopulated() bool {

	if len(fMgr.absolutePathFileName) == 0 {
		fMgr.isAbsolutePathFileNamePopulated = false
	} else {
		fMgr.isAbsolutePathFileNamePopulated = true
	}

	return fMgr.isAbsolutePathFileNamePopulated
}

// IsFileExtPopulated - Returns a boolean value indicating
// whether the File Extension for this File Manager instance
// is populated.
//
func (fMgr *FileMgr) IsFileExtPopulated() bool {

	if len(fMgr.fileExt) == 0 {
		fMgr.isFileExtPopulated = false
	} else {
		fMgr.isFileExtPopulated = true
	}

	return fMgr.isFileExtPopulated
}

// IsFileMgrValid - Analyzes the current FileMgr object. If the
// current FileMgr object is INVALID, an error is returned.
//
// If the current FileMgr is VALID, this method returns 'nil'
//
func (fMgr *FileMgr) IsFileMgrValid(errorPrefixStr string) error {

	ePrefix := strings.TrimRight(errorPrefixStr, " ") + "FileMgr.IsFileMgrValid()"

	if !fMgr.isInitialized {
		return errors.New(ePrefix + " Error: This data structure is NOT initialized.")
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return errors.New(ePrefix + " Error: absolutePathFileName is NOT populated/initialized.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return errors.New(ePrefix + " Error: absolutePathFileName is EMPTY!")
	}

	_, _ = fMgr.DoesThisFileExist()

	err := fMgr.dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return fmt.Errorf("FileMgr Directory Manager INVALID - %v", err.Error())
	}

	return nil
}

// IsFileNamePopulated - returns a boolean value
// indicating whether the file name for this File
// Manager object is populated.
//
func (fMgr *FileMgr) IsFileNamePopulated() bool {

	if len(fMgr.fileName) == 0 {
		fMgr.isFileNamePopulated = false
	} else {
		fMgr.isFileNamePopulated = true
	}

	return fMgr.isFileNamePopulated
}

// IsFileNameExtPopulated - Returns a boolean value indicating
// whether the File Name and Extension for this File Manager
// instance has been populated.
func (fMgr *FileMgr) IsFileNameExtPopulated() bool {

	if len(fMgr.fileNameExt) == 0 {
		fMgr.isFileNameExtPopulated = false
	} else {
		fMgr.isFileNameExtPopulated = true
	}

	return fMgr.isFileNameExtPopulated
}

// IsFilePointerOpen - Returns a boolean value indicating
// whether the File Pointer (*os.File) for this File Manager
// instance is open, or not.
//
func (fMgr *FileMgr) IsFilePointerOpen() bool {
	return fMgr.isFilePtrOpen
}

// isInitialized - Returns a boolean indicating whether the FileMgr
// object is properly initialized.
//
func (fMgr *FileMgr) IsInitialized() bool {
	return fMgr.isInitialized
}

// MoveFileToNewDirMgr - This method will move the file identified
// by the current FileMgr to a new path contained in the input parameter
// 'dMgr'.
//
// IMPORTANT:
// ==========
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
func (fMgr *FileMgr) MoveFileToNewDirMgr(dMgr DirMgr) (newFMgr FileMgr, err error) {
	ePrefix := "FileMgr.MoveFileToNewDirMgr() "
	newFMgr = FileMgr{}
	err = nil

	if !fMgr.isInitialized {
		err = errors.New(ePrefix +
			"Error: The current FileMgr object is NOT initialized. It is Empty!")
		return
	}

	err2 := fMgr.IsFileMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Error: Current FileMgr object is INVALID!. Error='%v'",
			err2.Error())
		return
	}

	if !dMgr.isInitialized {
		err = errors.New(ePrefix +
			"Error: Input parameter 'dMgr' is NOT initialized. It is Empty!")
		return
	}

	err2 = dMgr.IsDirMgrValid("")

	if err2 != nil {
		err = errors.New(ePrefix + "Error: Input parameter 'dMgr' reports as INVALID!")
		return
	}

	if !dMgr.doesAbsolutePathExist && !dMgr.doesPathExist {
		err = fmt.Errorf(ePrefix+
			"Error: Destination path DOES NOT EXIST!. "+
			"For this DirMgr object, both absolutePath and path DO NOT EXIST! "+
			"dMgr.absolutePath='%v'  dMgr.path='%v'", dMgr.absolutePath, dMgr.path)
		return
	}

	srcFileExists, err2 := fMgr.DoesThisFileExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by fMgr.DoesThisFileExist(). "+
			"fMgr.absolutePathFileName='%v'  Error='%v'", fMgr.absolutePathFileName, err2.Error())
		return
	}

	if !srcFileExists {
		err = fmt.Errorf(ePrefix+
			"Error: The source file identified by the current FileMgr object DOES NOT EXIST! "+
			"fMgr.absolutePathFileName='%v'", fMgr.absolutePathFileName)
		return
	}

	var destPathFileName string

	if dMgr.doesAbsolutePathExist {
		destPathFileName = dMgr.GetAbsolutePathWithSeparator() + fMgr.fileNameExt
	} else {
		destPathFileName = dMgr.GetPathWithSeparator() + fMgr.fileNameExt
	}

	fh := FileHelper{}
	_, err = fh.MoveFile(fMgr.absolutePathFileName, destPathFileName)

	if err != nil {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error returned from "+
				"fh.MoveFile(fMgr.absolutePathFileName, destPathFileName). "+
				"fMgr.absolutePathFileName='%v' pathFile='%v' Error='%v'",
				fMgr.absolutePathFileName, destPathFileName, err.Error())
	}

	newFMgr, err2 = FileMgr{}.New(destPathFileName)

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.New(destPathFileName). destPathFileName='%v' "+
			"Error='%v'", destPathFileName, err2.Error())
		return
	}

	doesFileExist, err2 := fMgr.DoesThisFileExist()

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix+"Error: Old file still exists! Error='%v'",
			err2.Error())
		return
	}

	if doesFileExist == true {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix+"Error: Old file still exists! Old File Name: %v ",
			fMgr.absolutePathFileName)
		return
	}

	err = nil

	return newFMgr, err

}

// MoveFileToNewDir - This method will move the current file
// identified by this FileMgr object to a new path designated
// by input parameter string, 'dirPath'.
//
// IMPORTANT:
// ==========
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
func (fMgr *FileMgr) MoveFileToNewDir(dirPath string) (newFMgr FileMgr, err error) {

	newFMgr = FileMgr{}
	err = nil

	ePrefix := "FileMgr.MoveFileToNewDir() "

	lPath := len(dirPath)

	if lPath == 0 {
		err = errors.New(ePrefix +
			"Error: Input parameter 'dirPath' is a Zero length string!")
		return
	}

	if !fMgr.isInitialized {
		err = errors.New(ePrefix +
			"Error: The current FileMgr object is NOT Initialized! It is EMPTY!")
		return
	}

	err2 := fMgr.IsFileMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error: Current FileMgr object is INVALID!. Error='%v'",
			err2.Error())
		return
	}

	if !fMgr.doesAbsolutePathFileNameExist {
		err = fmt.Errorf(ePrefix+
			"Error: The source files does NOT exist. srcFile='%v' ",
			fMgr.absolutePathFileName)
		return
	}

	dMgr, err2 := DirMgr{}.New(dirPath)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from DirMgr{}.New(dirPath). dirPath='%v'  Error='%v'",
			dirPath, err2.Error())
		return
	}

	if !dMgr.isInitialized {
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'dirPath' "+
			"generated an empty DirMgr object. dirPath='%v'", dirPath)
		return
	}

	pathExists, absPathExists := dMgr.DoesDirectoryExist()

	if !pathExists && !absPathExists {
		err = fmt.Errorf(ePrefix+
			"Error: Target Destination path DOES NOT EXIST! dirPath='%v'",
			dirPath)
		return
	}

	newFMgr, err2 = fMgr.MoveFileToNewDirMgr(dMgr)

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix+
			"Error returned by fMgr.MoveFileToNewDirMgr(dMgr). "+
			"dMgr.path='%v' Error='%v'", dMgr.path, err2.Error())
		return
	}

	return newFMgr, nil
}

// New - Creates a new FileMgr object. Input parameter parses out the
// path, file name and file extension. The file data is returned in
// the data fields of the new FileMgr object.
//
// Input Parameter
// ===============
//
// pathFileNameExt string - Must consist of a valid path, file name
// 													and file extension. The file need not exist.
//													Failure to provide a properly formatted path
//													path, file name will result in an error.
//
// Example Usage:
// fmgr := FileMgr{}.New("../common/fileName.ext")
//
func (fMgr FileMgr) New(pathFileNameExt string) (FileMgr, error) {

	ePrefix := "FileMgr.New() "

	if pathFileNameExt == "" {
		return FileMgr{}, errors.New(ePrefix + "-Error: pathFileNameExt is Empty!")
	}

	fMgrOut := FileMgr{}

	isEmpty, err := fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error returned from fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt) "+
			"pathFileNameExt='%v'  Error='%v'", pathFileNameExt, err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Returned FileMgr is Empty! pathFileNameExt='%v' ", pathFileNameExt)
	}

	return fMgrOut, nil
}

// NewFromFileInfo - Creates and returns a new FileMgr object based on input from a
// directory path string and an os.FileInfo object.
func (fMgr FileMgr) NewFromFileInfo(pathStr string, info os.FileInfo) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromFInfo() "
	var err error

	fh := FileHelper{}

	pathStr, err = fh.AddPathSeparatorToEndOfPathStr(pathStr)

	if err != nil {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error returned from fh.AddPathSeparatorToEndOfPathStr(pathStr). "+
				"pathStr='%v' Error='%v'", pathStr, err.Error())
	}

	pathFileName := pathStr + info.Name()

	fmgr2 := FileMgr{}

	isEmpty, err := fmgr2.SetFileMgrFromPathFileName(pathFileName)

	if err != nil {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error returned from fmgr2.SetFileMgrFromPathFileName(pathFileName). "+
				"pathFileName='%v' Error='%v'", pathFileName, err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error returned FileMgr is Empty! pathFileName='%v'", pathFileName)
	}

	fmgr2.actualFileInfo = FileInfoPlus{}.NewPathFileInfo(pathStr, info)

	return fmgr2, nil
}

// NewFromDirMgrFileNameExt - this method is designed to create a new File Manager (FileMgr)
// object from two input parameters.
//
// Input parameter 'dirMgr' is a valid and correctly populated Directory Manager object containing
// the file path.
//
// Input parameter 'fileNameExt' is a string containing the file name and the file extension.
//
// 'dirMgr' and 'fileNameExt' will be combined to create a new File Manager (FileMgr) with a
// properly configured path, file name and file extension.
//
func (fMgr FileMgr) NewFromDirMgrFileNameExt(
	dirMgr DirMgr,
	fileNameExt string) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromDirMgrFileNameExt() "

	if len(fileNameExt) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: Input Parameter fileNameExt is a zero length string!")
	}

	if dirMgr.isInitialized {
		err := dirMgr.IsDirMgrValid("")

		if err != nil {
			return FileMgr{},
				errors.New(ePrefix +
					"Error: DirMgr object, 'dirMgr', passed as input parameter is INVALID!")
		}

	}

	fmgr2 := FileMgr{}

	isEmpty, err := fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error returned by fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt). "+
			"Error='%v'", err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Empty FileMgr object returned by fmgr2.SetFileMgrFromDirMgrFileName("+
			"dirMgr, fileNameExt) dirMgr.path='%v' fileNameExt='%v'", dirMgr.path, fileNameExt)
	}

	return fmgr2, nil
}

// NewFromDirStrFileNameStr - Creates a new file manager object (FileMgr) from a directory
// string and a File Name and Extension string passed as input parameters.
func (fMgr FileMgr) NewFromDirStrFileNameStr(
	dirStr,
	fileNameExtStr string) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromDirStrFileNameStr() "

	if len(dirStr) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: Input parameter 'dirStr' is a Zero Length String!")
	}

	if len(fileNameExtStr) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: Input parameter 'fileNameExtStr' is a Zero Length String!")
	}

	dMgr, err := DirMgr{}.New(dirStr)

	if err != nil {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error returned by DirMgr{}.New(dirStr). dirStr='%v'  Error='%v'",
				dirStr, err.Error())
	}

	fMgr2, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExtStr)

	if err != nil {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExtStr). "+
				"fileNameExtStr='%v'  Error='%v'", fileNameExtStr, err.Error())
	}

	return fMgr2, nil

}

// OpenThisFileReadOnly - Opens the file identified by the current
// FileMgr object as a 'Read-Only' File. Subsequent operations may
// read from this file but may NOT write to this file.
//
//
func (fMgr *FileMgr) OpenThisFileReadOnly() error {
	ePrefix := "FileMgr.OpenThisFileReadWrite() "
	var err error

	if !fMgr.isInitialized {
		return errors.New(ePrefix +
			"Error: The File Manager data structure has NOT been initialized.")
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return errors.New(ePrefix +
			"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return errors.New(ePrefix +
			"Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if fMgr.isFilePtrOpen {
		err = fMgr.CloseFile()
		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by fMgr.CloseFile(). absolutePathFileName='%v'  Error='%v'",
				fMgr.absolutePathFileName, err.Error())
		}
	}

	doesFileExist, err := fMgr.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.DoesThisFileExist() - %v", err.Error())
	}

	if !doesFileExist {

		err = fMgr.CreateDirAndFile()

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error from fMgr.CreateDirAndFile(fMgr.absolutePathFileName). "+
				"absolutePathFileName='%v'. Error='%v'", fMgr.absolutePathFileName, err.Error())
		}

		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.isAbsolutePathFileNamePopulated = true
	}

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error creating File. Error returned from os.Create(fMgr.absolutePathFileName). "+
			"fMgr.absolutePathFileName='%v' Error='%v' ", fMgr.absolutePathFileName, err.Error())
	}

	fMgr.filePtr, err = os.OpenFile(fMgr.absolutePathFileName, os.O_RDONLY, 0666)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error opening file: '%v' Error= '%v'", fMgr.absolutePathFileName, err.Error())
	}

	fMgr.isFilePtrOpen = true

	return nil

}

// OpenThisFileReadWrite - Opens a file using file data in the
// current FileHelper data fields. If successful, this method
// will use FileHelper.absolutePathFileName to open an *os.File
// or File Pointer.
//
// As the method's name implies, the 'FileHelper.absolutePathFileName'
// will be opened for reading and writing. If FileHelper.absolutePathFileName
// does not exist, it will be created. The FileMode is set to'rwxrwxrwx' and
// the permission Mode= '0666'
//
func (fMgr *FileMgr) OpenThisFileReadWrite() error {
	var err error

	ePrefix := "FileMgr.OpenThisFileReadWrite() "

	if !fMgr.isInitialized {
		return errors.New(ePrefix +
			"Error: The File Manager data structure has NOT been initialized.")
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return errors.New(ePrefix +
			"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return errors.New(ePrefix +
			"Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if fMgr.isFilePtrOpen || fMgr.filePtr != nil {
		err = fMgr.CloseFile()
		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by fMgr.CloseFile(). absolutePathFileName='%v'  Error='%v'",
				fMgr.absolutePathFileName, err.Error())
		}
	}

	doesFileExist, err := fMgr.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.DoesThisFileExist() - %v", err.Error())
	}

	if !doesFileExist {

		err = fMgr.CreateDirAndFile()

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error from fMgr.CreateDirAndFile(fMgr.absolutePathFileName). "+
				"absolutePathFileName='%v'. Error='%v'",
				fMgr.absolutePathFileName, err.Error())
		}

		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.isAbsolutePathFileNamePopulated = true

		return nil
	}

	fMgr.filePtr, err = os.OpenFile(fMgr.absolutePathFileName, os.O_RDWR, 0666)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error opening file: '%v' Error= '%v'",
			fMgr.absolutePathFileName, err.Error())
	}

	fMgr.isFilePtrOpen = true

	return nil
}

// ReadAllFile - Reads the file identified by the current FileMgr
// and returns the contents in a byte array.
//
// If no errors are encountered the returned 'error' value is
// nil.
func (fMgr *FileMgr) ReadAllFile() ([]byte, error) {
	ePrefix := "FileMgr.ReadAllFile() "
	var err error

	if !fMgr.isInitialized {
		return []byte{}, errors.New(ePrefix +
			"Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return []byte{}, fmt.Errorf(ePrefix+
			"Error - This File Manager is INVALID! Error='%v'", err.Error())
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return []byte{}, errors.New(ePrefix +
			"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return []byte{}, errors.New(ePrefix + "Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if !fMgr.isFilePtrOpen || fMgr.filePtr == nil {

		if fMgr.filePtr != nil {

			err = fMgr.CloseFile()

			if err != nil {
				return []byte{}, fmt.Errorf(ePrefix+
					"Error: Failed to Close '%v'. Error='%v'",
					fMgr.absolutePathFileName, err.Error())
			}
		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return []byte{},
				fmt.Errorf(ePrefix+
					"Error returned from fMgr.OpenThisFileReadWrite() fileNameExt='%v' Error='%v'",
					fMgr.absolutePathFileName, err.Error())
		}

		fMgr.isFilePtrOpen = true
	}

	bytesRead, err := ioutil.ReadAll(fMgr.filePtr)

	if err != nil {
		return []byte{},
			fmt.Errorf(ePrefix+
				"Error returned by ioutil.ReadAll(fMgr.filePtr). "+
				"fileName='%v' Errors='%v'",
				fMgr.absolutePathFileName, err.Error())
	}

	return bytesRead, nil
}

// ReadFileBytes - Reads bytes from the file identified by the current FileMgr
// object. Bytes are stored in 'byteBuff', a byte array passed in as an input
// parameter.
//
// If successful, the returned error value is 'nil'. The returned value 'int'
// contains the number of bytes read from the current file.
func (fMgr *FileMgr) ReadFileBytes(byteBuff []byte) (int, error) {

	ePrefix := "FileMgr.ReadFileBytes() "
	var err error

	if !fMgr.isInitialized {
		return 0, errors.New(ePrefix +
			"Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return 0, fmt.Errorf(ePrefix+
			"Error - This File Manager is INVALID! Error='%v'", err.Error())
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return 0,
			errors.New(ePrefix +
				"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return 0,
			errors.New(ePrefix + " Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if !fMgr.isFilePtrOpen || fMgr.filePtr == nil {

		if fMgr.filePtr != nil {
			err = fMgr.CloseFile()

			if err != nil {
				return 0, fmt.Errorf(ePrefix+
					"Error: Failed to close fMgr.absolutePathFileName='%v'. Error='%v' ",
					fMgr.absolutePathFileName, err.Error())
			}

		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return 0,
				fmt.Errorf(ePrefix+
					"Error returned by fMgr.OpenThisFileReadWrite()  fileName='%v'  Error='%v'",
					fMgr.absolutePathFileName, err.Error())
		}

		fMgr.isFilePtrOpen = true
	}

	bytesRead, err := fMgr.filePtr.Read(byteBuff)

	if err != nil {
		return bytesRead, fmt.Errorf(ePrefix+
			"Error returned by fMgr.filePtr.Read(byteBuff). "+
			"fileName='%v'  Error='%v'", fMgr.absolutePathFileName, err.Error())
	}

	return bytesRead, nil
}

// SetFileMgrFromDirMgrFileName - Sets the data fields of the current FileMgr object
// based on a DirMgr object and a File Name string which are passed as input parameters.
func (fMgr *FileMgr) SetFileMgrFromDirMgrFileName(
	dMgr DirMgr,
	fileNameExt string) (isEmpty bool, err error) {

	ePrefix := "FileMgr.SetFileMgrFromDirMgrFileName() "
	isEmpty = true
	err = nil
	fMgr.Empty()
	fh := FileHelper{}

	if dMgr.isInitialized {

		err2 := dMgr.IsDirMgrValid("")

		if err2 != nil {
			err = fmt.Errorf(ePrefix+
				"Error: Input parameter 'dMgr' is INVALID! dMgr.path='%v'  Error='%v'",
				dMgr.path, err2.Error())
			return
		}

	}

	if len(fileNameExt) == 0 {
		err = errors.New(ePrefix +
			"Error: Input parameter 'fileNameExt' is a Zero length string!")
		return
	}

	adjustedFileNameExt, isFileNameEmpty, err2 := fh.CleanFileNameExtStr(fileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Error returned from fh.CleanFileNameExtStr(fileNameExt). "+
			"fileNameExt='%v' Error='%v'", fileNameExt, err2.Error())
		return
	}

	if isFileNameEmpty {
		err = fmt.Errorf(ePrefix+
			"Error: fileName returned from fh.CleanFileNameExtStr(fileNameExt) "+
			"is a ZERO length string! fileNameExt='%v'", fileNameExt)
		return
	}

	fMgr.dMgr = dMgr.CopyOut()

	s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(adjustedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from fh.GetFileNameWithoutExt(adjustedFileNameExt). "+
			"adjustedFileNameExt='%v'  Error='%v' ", adjustedFileNameExt, err2.Error())
		fMgr.Empty()
		isEmpty = true
		return
	}

	if fNameIsEmpty {
		err = fmt.Errorf(ePrefix+
			"Error: fileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt) "+
			"is Zero length string! adjustedFileNameExt='%v'  ", adjustedFileNameExt)
		fMgr.Empty()
		isEmpty = true
		return

	}

	fMgr.isFileNamePopulated = true
	fMgr.fileName = s

	s, extIsEmpty, err2 := fh.GetFileExtension(adjustedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from fh.GetFileExt(fileNameAndExt). "+
			"fileNameAndExt='%v'  Error='%v' ", adjustedFileNameExt, err2.Error())
		fMgr.Empty()
		isEmpty = true
		return
	}

	if !extIsEmpty {
		fMgr.isFileExtPopulated = true
		fMgr.fileExt = s
	}

	if fMgr.isFileNamePopulated {
		fMgr.isFileNameExtPopulated = true
		fMgr.fileNameExt = fMgr.fileName + fMgr.fileExt
	}

	lPath := len(fMgr.dMgr.absolutePath)
	if lPath == 0 {
		fMgr.absolutePathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.absolutePath[lPath-1] == os.PathSeparator {
		fMgr.absolutePathFileName = fMgr.dMgr.absolutePath + fMgr.fileNameExt

	} else {
		fMgr.absolutePathFileName =
			fMgr.dMgr.absolutePath + string(os.PathSeparator) + fMgr.fileNameExt

	}

	lPath = len(fMgr.dMgr.path)

	if lPath == 0 {
		fMgr.originalPathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.path[lPath-1] == os.PathSeparator {
		fMgr.originalPathFileName = fMgr.dMgr.path + fMgr.fileNameExt

	} else {
		fMgr.originalPathFileName = fMgr.dMgr.path + string(os.PathSeparator) + fMgr.fileNameExt

	}

	fMgr.isAbsolutePathFileNamePopulated = true

	fInfo, err2 := os.Stat(fMgr.absolutePathFileName)

	if err2 == nil {
		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.actualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.dMgr.absolutePath, fInfo)
	} else {
		fMgr.doesAbsolutePathFileNameExist = false
		fMgr.actualFileInfo = FileInfoPlus{}
	}

	fMgr.isInitialized = true

	err = nil
	isEmpty = false

	return

}

// SetFileMgrFromPathFileName - Initializes all the data fields of the
// current FileMgr object based on the path file name string passed to
// this method as an input parameter.
func (fMgr *FileMgr) SetFileMgrFromPathFileName(
	pathFileNameExt string) (isEmpty bool, err error) {

	ePrefix := "FileMgr.SetFileMgrFromPathFileName() "
	isEmpty = true
	err = nil
	fh := FileHelper{}

	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix +
			"Error: Input parameter 'pathFileNameExt' is a zero length or empty string!")
		return
	}

	adjustedPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	adjustedFileNameExt, isEmptyFileName, err2 := fh.CleanFileNameExtStr(adjustedPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt). "+
			"adjustedPathFileNameExt='%v' Error='%v'", adjustedPathFileNameExt, err2.Error())
		return
	}

	if isEmptyFileName {
		err = fmt.Errorf(ePrefix+
			"Error: File Name returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt) "+
			"is a Zero Length String!. pathFileNameExt='%v'", adjustedPathFileNameExt)
	}

	remainingPathStr := strings.TrimSuffix(adjustedPathFileNameExt, adjustedFileNameExt)

	var dMgr DirMgr

	if len(remainingPathStr) == 0 {
		dMgr = DirMgr{}
	} else {

		dMgr, err2 = DirMgr{}.New(remainingPathStr)

		if err2 != nil {
			err = fmt.Errorf(ePrefix+
				"Error returned from DirMgr{}.New(remainingPathStr). "+
				"remainingPathStr='%v'  Error='%v'", remainingPathStr, err2.Error())
			return
		}

	}

	isEmptyFMgr, err2 := fMgr.SetFileMgrFromDirMgrFileName(dMgr, adjustedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, adjustedFileNameExt). "+
			"adjustedFileNameExt='%v'  Error='%v'", adjustedFileNameExt, err2.Error())
		return
	}

	if isEmptyFMgr {
		err = fmt.Errorf(ePrefix+
			"Error: Empty FileMgr returned from fMgr.SetFileMgrFromDirMgrFileName("+
			"dMgr, adjustedFileNameExt). dMgr.path='%v'   adjustedFileNameExt='%v' ",
			dMgr.path, adjustedFileNameExt)
		return
	}

	isEmpty = false
	err = nil

	return
}

// SetFileInfo - Used to initilize the os.FileInfo structure maintained as
// part of the current FileMgr object.
func (fMgr *FileMgr) SetFileInfo(info os.FileInfo) error {
	ePrefix := "FileMgr.SetFileInfo() "

	if info.Name() == "" {
		return errors.New(ePrefix + "Error: info.Name() is an EMPTY string!")
	}

	if info.IsDir() {
		return errors.New(ePrefix + "info.IsDir()=='true'. This is a Directory NOT A FILE!")
	}

	fMgr.actualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)

	if !fMgr.actualFileInfo.IsFInfoInitialized {
		return fmt.Errorf(ePrefix+
			"Error: Failed to initialize fMgr.actualFileInfo object. info.Name()='%v'",
			info.Name())
	}

	return nil
}

// WriteBytesToFileWrites a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
func (fMgr *FileMgr) WriteBytesToFile(bytes []byte) (int, error) {

	ePrefix := "FileMgr.WriteBytesToFile() "
	var err error

	if !fMgr.isInitialized {
		return 0,
			errors.New(ePrefix +
				"Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return 0,
			fmt.Errorf(ePrefix+
				"Error: This File Manger is INVALID! fileNameExt='%v'  "+
				"Error='%v'", fMgr.absolutePathFileName, err.Error())
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return 0,
			errors.New(ePrefix +
				"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return 0, errors.New(ePrefix + "Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if !fMgr.isFilePtrOpen || fMgr.filePtr == nil {

		if fMgr.filePtr != nil {
			err = fMgr.CloseFile()
			if err != nil {
				return 0, fmt.Errorf(ePrefix+"Error: failed to close %v.  Error='%v' ",
					fMgr.absolutePathFileName, err.Error())
			}

		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return 0,
				fmt.Errorf(ePrefix+
					" - fMgr.OpenThisFileReadWrite() returned errors: %v",
					err.Error())
		}

		fMgr.isFilePtrOpen = true
	}

	bytesWritten, err := fMgr.filePtr.Write(bytes)

	if err != nil {
		return bytesWritten,
			fmt.Errorf(ePrefix+
				"Error returned from fMgr.filePtr.Write(str). Output File='%v'. "+
				"Error='%v'", fMgr.absolutePathFileName, err.Error())
	}

	return bytesWritten, nil
}

// WriteStrToFile - Writes a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
func (fMgr *FileMgr) WriteStrToFile(str string) (int, error) {

	ePrefix := "FileMgr.WriteStrToFile() "
	var err error

	if !fMgr.isInitialized {
		return 0, errors.New(ePrefix + "Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return 0,
			fmt.Errorf(ePrefix+
				"Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'",
				fMgr.absolutePathFileName, err.Error())
	}

	if !fMgr.isAbsolutePathFileNamePopulated {
		return 0,
			errors.New(ePrefix +
				"Error: FileMgr.absolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.absolutePathFileName == "" {
		fMgr.isAbsolutePathFileNamePopulated = false
		return 0, errors.New(ePrefix + "Error: FileMgr.absolutePathFileName is EMPTY!")
	}

	if !fMgr.isFilePtrOpen || fMgr.filePtr == nil {

		if fMgr.filePtr != nil {
			err = fMgr.CloseFile()
			if err != nil {
				return 0,
					fmt.Errorf(ePrefix+
						"Error: failed to close %v.  Error='%v' ", fMgr.absolutePathFileName, err.Error())
			}

		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return 0, fmt.Errorf(ePrefix+
				" - fMgr.OpenThisFileReadWrite() returned errors: %v", err.Error())
		}

		fMgr.isFilePtrOpen = true
	}

	bytesWritten, err := fMgr.filePtr.WriteString(str)

	if err != nil {
		return bytesWritten,
			fmt.Errorf(ePrefix+
				"Error returned from fMgr.filePtr.WriteString(str). Output File='%v'. "+
				"Error='%v'", fMgr.absolutePathFileName, err.Error())

	}

	return bytesWritten, nil
}
