package pathfileops

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	This source code file contains types 'FileMgr' and
	'FileMgrCollection'.

	The Source Repository for this source code file is :
		https://github.com/MikeAustin71/pathfileopsgo.git

	Dependencies:
	-------------

	Types 'FileMgr' and 'FileMgrCollection' depend on type,
	'FileHelper' which is contained in source code file,
	'filehelper.go' located in this directory.

*/

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

	fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(pathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from FileMgr{}.NewFromPathFileNameExtStr(pathFileName). "+
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
		err := fMgrs.FMgrs[i].CopyFileToDirByIoByLink(targetDirectory)

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

// CopyFileToDirByIoByLink - Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by input parameter 'dir',
// an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the file to the destination by creating a new file and
// copying the source file contents to the new destination file using a
// technique known as 'io.Copy'.
//
// If that attempted file copy operation fails, a second attempt will be made
// using a technique known as a 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to create the destination
// file.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileToDirByIoByLink(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDirByIoByLink() "
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

	err = fMgr.CopyFileMgrByIoByLink(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByIoByLink(&newFMgr) "+
			"newFMgr.absolutePathFileName='%v'  Error='%v'",
			newFMgr.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileToDirByLinkByIo - Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by input parameter 'dir',
// an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the source file to the destination directory using a
// technique known as a 'Hard Link'.  This technique will utilize a hard
// symbolic link to the existing source file in order to create the destination
// file.
//
// If the first copy attempt fails, this method will try to copy the file to the
// destination directory by creating a new file and copying the source file contents
// to the new destination file. This technique is known as 'io.Copy'.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileToDirByLinkByIo(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDirByLinkByIo() "
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

	err = fMgr.CopyFileMgrByLinkByIo(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByLinkByIo(&newFMgr) "+
			"newFMgr.absolutePathFileName='%v'  Error='%v'",
			newFMgr.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileToDirByIo - Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by input parameter 'dir',
// an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// One attempt will be made to copy the source file to the specified destination
// directory using a technique known as 'io.Copy'. This technique create a new
// destination file and copies the source file contents to that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileToDirByIo(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDirByIo() "
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

	err = fMgr.CopyFileMgrByIo(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByIo(&newFMgr) "+
			"newFMgr.absolutePathFileName='%v'  Error='%v'",
			newFMgr.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileToDirByLink - Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by input parameter 'dir',
// an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// This method will make one attempt to copy the source file to the specified
// destination directory using a technique known as a 'Hard Link'. This technique
// will utilize a hard symbolic link to the existing source file in order to
// create the destination file.
//
// If the 'Hard Link' copy operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileToDirByLink(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDirByLink() "
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

	err = fMgr.CopyFileMgrByLink(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByLink(&newFMgr) "+
			"newFMgr.absolutePathFileName='%v'  Error='%v'",
			newFMgr.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileMgrByIoByLink - Copies the file represented by the current
// File Manager instance to a location specified by a destination input
// parameter 'fMgrDest', an instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the source file to the destination by creating a new file
// and copying the source file contents to the new destination file using a
// technique known as 'io.Copy'.
//
// If that first file copy operation fails, a second attempt will be made
// using a technique known as a 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to create the destination
// file.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileMgrByIoByLink(fMgrDest *FileMgr) error {

	ePrefix := "FileMgr.CopyFileMgrByIoByLink() "

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
			"File and cannot be copied. File='%v'", fMgr.absolutePathFileName)
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
			"Error returned by fh.CopyFileByIoByLink(fMgr.absolutePathFileName, "+
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

// CopyFileMgrByLinkByIo - Copies the file represented by the current
// File Manager instance to a location specified by a destination input
// parameter 'fMgrDest', an instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the source file to the destination using a technique known
// as a 'Hard Link'.  This technique will utilize a hard symbolic link to the
// existing source file in order to create the destination file.
//
// If the first copy attempt fails, this method will try to copy the file to the
// destination by creating a new file and copying the source file contents to that
// new destination file. This technique is known as 'io.Copy'.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileMgrByLinkByIo(fMgrDest *FileMgr) error {

	ePrefix := "FileMgr.CopyFileMgrByLinkByIo() "

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
			"File and cannot be copied. File='%v'", fMgr.absolutePathFileName)
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

	err = fh.CopyFileByLinkByIo(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

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

// CopyFileMgrByIo - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter 'fMgrDest', an instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// One attempt will be made to copy the source file to the specified destination
// file using a technique known as 'io.Copy'. This technique create a new
// destination file and copies the source file contents to that new destination
// file.
//
// If this attempted 'io.Copy' operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileMgrByIo(fMgrDest *FileMgr) error {

	ePrefix := "FileMgr.CopyFileMgrByIo() "

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
			"File and cannot be copied. File='%v'", fMgr.absolutePathFileName)
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

	err = fh.CopyFileByIo(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.CopyFileByIo(fMgr.absolutePathFileName, "+
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

// CopyFileMgrByLink - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter 'fMgrDest', an instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// This method will make one attempt to copy the source file to the specified
// destination using a technique known as a 'Hard Link'. This technique will
// utilize a hard symbolic link to the existing source file in order to create
// the destination file.
//
// If the 'Hard Link' copy operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileMgrByLink(fMgrDest *FileMgr) error {

	ePrefix := "FileMgr.CopyFileMgrByLink() "

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
			"File and cannot be copied. File='%v'", fMgr.absolutePathFileName)
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

	err = fh.CopyFileByLink(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.CopyFileByLink(fMgr.absolutePathFileName, "+
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

// CopyFileStrByIoByLink - Copies the file represented by the current
// File Manager instance to a location specified by a destination input
// parameter. The destination input parameter, 'dstPathFileNameExt' is
// a string containing the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the file to the destination by creating a new file and
// copying the source file contents to the new destination file. This technique
// is known as 'io.Copy'.
//
// If that attempted file copy operation fails, a second attempt will be made
// using a technique known as a 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to create the destination
// file.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileStrByIoByLink(dstPathFileNameExt string) error {

	ePrefix := "FileMgr.CopyFileStrByIoByLink() "

	fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt). "+
			"dstPathFileNameExt='%v' Error='%v'", dstPathFileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgrByIoByLink(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByIoByLink(&fMgrDest) "+
			"fMgrDest.absolutePathFileName='%v'  Error='%v'", fMgrDest.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileStrByLinkByIo - Copies the file represented by the current
// File Manager instance to a location specified by a destination input
// parameter. The destination input parameter, 'dstPathFileNameExt' is
// a string containing the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// The copy operation will be carried out in two attempts. The first attempt
// will try to copy the source file to the destination using a technique known
// as a 'Hard Link'.  This technique will utilize a hard symbolic link to the
// existing source file in order to create the destination file.
//
// If the first copy attempt fails, this method will try to copy the file to the
// destination by creating a new file and copying the source file contents to the
// new destination file. This technique is known as 'io.Copy'.
//
// If both attempted copy operations fail, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileStrByLinkByIo(dstPathFileNameExt string) error {

	ePrefix := "FileMgr.CopyFileStrByLinkByIo() "

	fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt). "+
			"dstPathFileNameExt='%v' Error='%v'", dstPathFileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgrByLinkByIo(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByLinkByIo(&fMgrDest) "+
			"fMgrDest.absolutePathFileName='%v'  Error='%v'", fMgrDest.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileStrByIo - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter. The destination input parameter, 'dstPathFileNameExt' is
// a string containing the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// One attempt will be made to copy the source file to the specified destination
// using a technique known as 'io.Copy'. This technique create a new destination
// file and copies the source file contents to that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileStrByIo(dstPathFileNameExt string) error {

	ePrefix := "FileMgr.CopyFileStrByIo() "

	fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt). "+
			"dstPathFileNameExt='%v' Error='%v'", dstPathFileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgrByIo(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByIo(&fMgrDest) "+
			"fMgrDest.absolutePathFileName='%v'  Error='%v'", fMgrDest.absolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileStrByLink - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter. The destination input parameter, 'dstPathFileNameExt' is
// a string containing the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// This method will make one attempt to copy the source file to the specified
// destination using a technique known as a 'Hard Link'. This technique will
// utilize a hard symbolic link to the existing source file in order to create
// the destination file.
//
// If 'Hard Link' copy operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgr *FileMgr) CopyFileStrByLink(dstPathFileNameExt string) error {

	ePrefix := "FileMgr.CopyFileStrByLink() "

	fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt). "+
			"dstPathFileNameExt='%v' Error='%v'", dstPathFileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgrByLink(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from fMgr.CopyFileMgrByLink(&fMgrDest) "+
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
		fMgr.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(fMgr.absolutePathFileName, info)
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

	fMgr.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(fMgr.dMgr.absolutePath, info)

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

	fMgr.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(fMgr.dMgr.absolutePath, info)

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

	newFMgr, err2 = FileMgr{}.NewFromPathFileNameExtStr(destPathFileName)

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix+
			"Error returned by FileMgr{}.NewFromPathFileNameExtStr(destPathFileName). destPathFileName='%v' "+
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
			"Error returned from DirMgr{}.NewFromPathFileNameExtStr(dirPath). dirPath='%v'  Error='%v'",
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

// NewFromPathFileNameExtStr - Creates a new FileMgr object. Input parameter parses out the
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
// fmgr := FileMgr{}.NewFromPathFileNameExtStr("../common/fileName.ext")
//
func (fMgr FileMgr) NewFromPathFileNameExtStr(pathFileNameExt string) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromPathFileNameExtStr() "

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

	fmgr2.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathStr, info)

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
				"Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirStr). dirStr='%v'  Error='%v'",
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
		fMgr.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(fMgr.dMgr.absolutePath, fInfo)
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
				"Error returned from DirMgr{}.NewFromPathFileNameExtStr(remainingPathStr). "+
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

type FileOpsCollection struct {
	fileOps []FileOps
}

type FileOps struct {
	isInitialized bool
	source        FileMgr
	destination   FileMgr
	opToExecute   FileOperation
}

// NewByFileMgrs - Creates and returns a new FileOps
// instance based on input parameters 'source' and
// 'destination' File Managers.
//
func (fops FileOps) NewByFileMgrs(
	source,
	destination FileMgr) (FileOps, error) {

	ePrefix := "FileOps.NewByFileMgrs() "

	err := source.IsFileMgrValid(ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf("Source File Manager INVALID! %v", err.Error())
	}

	err = destination.IsFileMgrValid(ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf("Destination File Manager INVALID! %v", err.Error())
	}

	fOpsNew := FileOps{}

	fOpsNew.source = source.CopyOut()
	fOpsNew.destination = destination.CopyOut()

	return fOpsNew, nil
}

// NewByDirMgrFileName - Creates and returns a new FileOps instance
// based on input parameters, source Directory Manger, source file name
// and extension string, destination Directory Manager and destination
// file name and extension string.
//
func (fops FileOps) NewByDirMgrFileName(
	sourceDir DirMgr,
	sourceFileNameExt string,
	destinationDir DirMgr,
	destinationFileNameExt string) (FileOps, error) {

	ePrefix := "FileOps.NewByDirMgrFileName() "

	var err error

	fOpsNew := FileOps{}

	fOpsNew.source, err = FileMgr{}.NewFromDirMgrFileNameExt(sourceDir, sourceFileNameExt)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Source File Error: %v", err.Error())
	}

	fOpsNew.destination, err = FileMgr{}.NewFromDirMgrFileNameExt(destinationDir, destinationFileNameExt)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Destination File Error: %v", err.Error())
	}

	return fOpsNew, nil
}

// NewByPathFileNameExtStrs - Creates and returns a new FileOps instance
// based on two string input parameters. The first represents the path name,
// file name and extension of the source file. The second represents the
// path name, file name and extension of the destination file.
//
func (fops FileOps) NewByPathFileNameExtStrs(
	sourcePathFileNameExt string,
	destinationPathFileNameExt string) (FileOps, error) {

	ePrefix := "FileOps.NewByPathFileNameExtStrs() "

	fOpsNew := FileOps{}

	var err error

	fOpsNew.source, err = FileMgr{}.NewFromPathFileNameExtStr(sourcePathFileNameExt)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Source File Error: %v", err.Error())
	}

	fOpsNew.destination, err = FileMgr{}.NewFromPathFileNameExtStr(destinationPathFileNameExt)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Destination File Error: %v", err.Error())
	}

	return fOpsNew, nil
}
