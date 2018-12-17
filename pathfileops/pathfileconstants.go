package pathfileops

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type FileOperation int

const (
	// Moves the source file to the destination file and
	// then deletes the original source file
	MOVESOURCETODESTINATION FileOperation = iota

	// Deletes the Destination file if it exists
	//
	DELETEDESTINATIONFILE

	// Deletes the Source file if it exists
	DELETESOURCEFILE

	// Deletes both the Source and Destination files
	// if they exist.
	DELETESOURCEandDESTINATIONFILES

	// Copies the Source File to the Destination
	// using two copy attempts. The first copy is
	// by Hard Link. If the first copy attempt fails,
	// a second copy attempt is initiated/ by creating
	// a new file and copying the contents by 'io.Copy'.
	// An error is returned only if both copy attempts
	// fail. The source file is unaffected.
	//
	// // See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
	//
	COPYSOURCETODESTINATIONByHardLinkByIo

	// Copies the Source File to the Destination
	// using two copy attempts. The first copy is
	// by 'io.Copy' which creates a new file and copies
	// the contents to the new file. If the first attempt
	// fails, a second copy attempt is initiated using
	// 'copy by hard link'. An error is returned only
	// if both copy attempts fail. The source file is
	// unaffected.
	//
	// // See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
	//
	COPYSOURCETODESTINATIONByIoByHardLink

	// Copies the Source File to the Destination
	// using one copy mode. The only copy attempt
	// utilizes 'Copy by Hard Link'. If this fails
	// an error is returned.  The source file is
	// unaffected.
	COPYSOURCETODESTINATIONByHardLink

	// Copies the Source File to the Destination
	// using only one copy mode. The only copy
	// attempt is initiated using 'Copy by IO' or
	// 'io.Copy'.  If this fails an error is returned.
	// The source file is unaffected.
	//
	COPYSOURCETODESTINATIONByIo

	// Creates the Source Directory
	CREATE_SOURCE_DIR

	// Creates the Source Directory and File
	CREATE_SOURCE_DIR_AND_FILE

	// Creates the Source File
	CREATE_SOURCE_FILE

	// Creates the Destination Directory
	CREATE_DESTINATION_DIR

	// Creates the Destination Directory and File
	CREATE_DESTINATION_DIR_AND_FILE

	// Creates the Destination File
	CREATE_DESTINATION_FILE
)

func (pathFileOp FileOperation) String() string {
	return PathFileOperationNames[pathFileOp]
}

var PathFileOperationNames = [...]string{"Delete Destination",
	"Delete Source",
	"Delete Source and Destination",
	"Copy Source to Destination ByHardLink ByIo",
	"Copy Source to Destination ByIo ByHardLink",
	"Copy Source to Destination By HardLink",
	"Copy Source to Destination By Io",
}

// FileSelectCriterionMode - Used in conjunction with the
// FileSelectionCriteria structure, below
type FileSelectCriterionMode int

// String - Method used to display the text
// name of an Operations Message Type.
func (fSelectMode FileSelectCriterionMode) String() string {
	return FileSelectCriterionModeNames[fSelectMode]
}

const (

	// ANDFILESELECTCRITERION - 0 File Selection Criterion are And'ed
	// together. If there are three file selection criterion then
	// all three must be satisfied before a file is selected.
	ANDFILESELECTCRITERION FileSelectCriterionMode = iota

	// ORFILESELECTCRITERION - 1 File Selection Criterion are Or'd together.
	// If there are three file selection criterion then satisfying any
	// one of the three criterion will cause the file to be selected.
	ORFILESELECTCRITERION
)

// FileSelectCriterionModeNames - String Array holding File Select Criteria  names.
var FileSelectCriterionModeNames = [...]string{"AND File Select Criterion", "OR File Select Criterion"}

// FileSelectionCriteria - Used is selecting file names. These
// data fields specify the criterion used to determine if a
// file should be selected for some type of operation.
// Example: find files or delete files operations
type FileSelectionCriteria struct {
	FileNamePatterns []string // a string array containing one or more file matching
	// patterns. Example '*.txt' '*.log' 'common*.*'
	FilesOlderThan   time.Time   // Used to select files with a modification less than this date time
	FilesNewerThan   time.Time   // Used to select files with a modification greater than this date time
	SelectByFileMode os.FileMode // Used to select files with equivalent FileMode values
	//   Note: os.FileMode is an uint32 type
	SelectCriterionMode FileSelectCriterionMode // Can be one of two values:
	// ANDFILESELECTCRITERION or ORFILESELECTCRITERION
	//
	// ANDFILESELECTCRITERION = select a file only if ALL
	//										      the selection criterion
	//                          are satisfied.
	//
	// ORFILESELECTCRITERION  = select a file if only ONE
	//													of the selection criterion
	//													are satisfied.
}

// ArePatternsActive - surveys the FileNamePatterns string
// array to determine if there currently any active search
// file pattern string.
//
// A search file pattern is considered active if the string
// length of the pattern string is greater than zero.
func (fsc *FileSelectionCriteria) ArePatternsActive() bool {

	lPats := len(fsc.FileNamePatterns)

	if lPats == 0 {
		return false
	}

	isActive := false

	for i := 0; i < lPats; i++ {
		fsc.FileNamePatterns[i] = strings.TrimRight(strings.TrimLeft(fsc.FileNamePatterns[i], " "), " ")
		if fsc.FileNamePatterns[i] != "" {
			isActive = true
		}

	}

	return isActive
}

// FileInfoPlus - Conforms to the os.FileInfo interface. This structure will store
// FileInfo information plus additional information related to a file or directory.
type FileInfoPlus struct {
	IsFInfoInitialized bool // Not part of FileInfo interface.
	// 'true' = structure fields have been properly initialized
	IsDirPathInitialized bool // Not part of FileInfo interface.
	// 'true' = structure field 'dirPath' has been successfully initialized
	CreateTimeStamp time.Time // Not part of FileInfo interface.
	// Date time at which this instance was initialized
	dirPath  string      // Not part of FileInfo interface. Directory path associated with file name
	fName    string      // FileInfo.Name() base name of the file
	fSize    int64       // FileInfo.Size() length in bytes for regular files; system-dependent for others
	fMode    os.FileMode // FileInfo.Mode() file mode bits
	fModTime time.Time   // FileInfo.ModTime() file modification time
	isDir    bool        // FileInfo.IsDir() 'true'= this is a directory not a file
	dataSrc  interface{} // FileInfo.Sys() underlying data source (can return nil)
}

// Name - base name of the file
func (fip FileInfoPlus) Name() string {

	return fip.fName
}

//Size - file length in bytes for regular files; system-dependent for others
func (fip FileInfoPlus) Size() int64 {
	return fip.fSize
}

// Mode - file mode bits. See os.FileMode
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
//
// type FileMode uint32
//
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
// const (
//  // The single letters are the abbreviations
//  // used by the String method's formatting.
// 	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
// 	ModeAppend                                     // a: append-only
// 	ModeExclusive                                  // l: exclusive use
// 	ModeTemporary                                  // T: temporary file; Plan 9 only
// 	ModeSymlink                                    // L: symbolic link
// 	ModeDevice                                     // D: device file
// 	ModeNamedPipe                                  // p: named pipe (FIFO)
// 	ModeSocket                                     // S: Unix domain socket
// 	ModeSetuid                                     // u: setuid
// 	ModeSetgid                                     // g: setgid
// 	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
// 	ModeSticky                                     // t: sticky
//
// 	// Mask for the type bits. For regular files, none will be set.
// 	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//
// 	ModePerm FileMode = 0777 // Unix permission bits
// )
//
func (fip FileInfoPlus) Mode() os.FileMode {
	return fip.fMode
}

// ModTime - file modification time
func (fip FileInfoPlus) ModTime() time.Time {
	return fip.fModTime
}

// IsDir - 'true' = this is a directory,
// not a file.
//
// abbreviation for Mode().IsDir()
//
func (fip FileInfoPlus) IsDir() bool {
	return fip.isDir
}

// Sys - underlying data source (can return nil)
func (fip FileInfoPlus) Sys() interface{} {
	return fip.dataSrc
}

// Sys - underlying data source (can return nil)
func (fip FileInfoPlus) SysAsString() string {
	if fip.dataSrc == nil {
		return ""
	}

	str := fmt.Sprintf("%v", fip.dataSrc)

	return str
}

// CopyOut - Creates a copy of the current FileInfoPlus
// instance and returns it.
func (fip *FileInfoPlus) CopyOut() FileInfoPlus {
	newInfo := FileInfoPlus{}

	newInfo.SetName(fip.Name())
	newInfo.SetSize(fip.Size())
	newInfo.SetMode(fip.Mode())
	newInfo.SetModTime(fip.ModTime())
	newInfo.SetIsDir(fip.IsDir())
	newInfo.SetSysDataSrc(fip.Sys())
	_ = newInfo.SetDirectoryPath(fip.DirPath())
	newInfo.IsFInfoInitialized = fip.IsFInfoInitialized
	newInfo.IsDirPathInitialized = fip.IsDirPathInitialized
	newInfo.CreateTimeStamp = fip.CreateTimeStamp
	return newInfo
}

// DirPath - Returns the directory path. This field, FileInfoPlus.dirPath,
// is not part of the standard FileInfo interface.
func (fip *FileInfoPlus) DirPath() string {
	return fip.dirPath
}

// Equal - Compares two FileInfoPlus objects to determine
// if they are equal.
func (fip *FileInfoPlus) Equal(fip2 *FileInfoPlus) bool {

	if fip.Name() != fip2.Name() ||
		fip.Size() != fip2.Size() ||
		fip.Mode() != fip2.Mode() ||
		fip.ModTime() != fip2.ModTime() ||
		fip.IsDir() != fip2.IsDir() {

		return false
	}

	if fip.DirPath() != fip2.DirPath() {
		return false
	}

	if fip.Sys() == nil && fip2.Sys() == nil {
		return true
	}

	if fip.Sys() == nil && fip2.Sys() != nil {
		return false
	}

	if fip.Sys() != nil && fip2.Sys() == nil {
		return false
	}

	strFipSys := fmt.Sprintf("%v", fip.Sys())
	strFip2Sys := fmt.Sprintf("%v", fip2.Sys())

	if strFipSys != strFip2Sys {

		return false
	}

	return true

}

// NewFromFileInfo - Creates and returns a new FileInfoPlus object
// populated with FileInfo data received from the input parameter.
// Notice that this version of the 'NewFromPathFileNameExtStr' method does NOT set the
// Directory path. This method is NOT part of the FileInfo interface.
//
// Example Usage:
//	fip := FileInfoPlus{}.NewFromFileInfo(info)
//  -- fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromFileInfo(info os.FileInfo) FileInfoPlus {
	newInfo := FileInfoPlus{}

	newInfo.SetName(info.Name())
	newInfo.SetSize(info.Size())
	newInfo.SetMode(info.Mode())
	newInfo.SetModTime(info.ModTime())
	newInfo.SetIsDir(info.IsDir())
	newInfo.SetSysDataSrc(info.Sys())
	newInfo.SetIsFInfoInitialized(true)
	return newInfo
}

// NewFromPathFileInfo - Creates and returns a new FileInfoPlus object
// populated with directory path and FileInfo data received from
// the input parameters.
//
// Example Usage:
//	fip := FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)
//  -- fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromPathFileInfo(
	dirPath string,
	info os.FileInfo) FileInfoPlus {

	newInfo := FileInfoPlus{}.NewFromFileInfo(info)
	_ = newInfo.SetDirectoryPath(dirPath)
	return newInfo
}

// SetDirectoryPath - Sets the dirPath field. This
// field is not part of the standard FileInfo data structure.
func (fip *FileInfoPlus) SetDirectoryPath(dirPath string) error {
	fh := FileHelper{}
	dirPath = strings.TrimLeft(strings.TrimRight(dirPath, " "), " ")

	if len(dirPath) == 0 {
		return fmt.Errorf("FileInfoPlus.SetDirectoryPath() Error: 'dirPath' is a Zero Length String!")
	}

	dirPath = fh.RemovePathSeparatorFromEndOfPathString(dirPath)
	fip.dirPath = dirPath
	fip.IsDirPathInitialized = true
	return nil
}

// SetName - Sets the file name field.
func (fip *FileInfoPlus) SetName(name string) {
	fip.fName = name
}

// SetSize - Sets the file size field
func (fip *FileInfoPlus) SetSize(fileSize int64) {
	fip.fSize = fileSize
}

// SetMode - Sets the file Mode
func (fip *FileInfoPlus) SetMode(fileMode os.FileMode) {
	fip.fMode = fileMode
}

// SetModTime - Sets the file modification time
func (fip *FileInfoPlus) SetModTime(fileModTime time.Time) {
	fip.fModTime = fileModTime
}

// SetIsDir - Sets is directory field.
func (fip *FileInfoPlus) SetIsDir(isDir bool) {
	fip.isDir = isDir
}

// SetSysDataSrc - Sets the dataSrc field
func (fip *FileInfoPlus) SetSysDataSrc(sysDataSrc interface{}) {
	fip.dataSrc = sysDataSrc
}

// SetIsFInfoInitialized - Sets the flag for 'Is File Info Initialized'
// If set to 'true' it means that all of the File Info fields have
// been initialized.
func (fip *FileInfoPlus) SetIsFInfoInitialized(isInitialized bool) {
	if !isInitialized {
		fip.IsFInfoInitialized = false
		fip.CreateTimeStamp = time.Time{}
		return
	}

	fip.IsFInfoInitialized = true
	fip.CreateTimeStamp = time.Now().Local()
	return
}

// DirectoryTreeInfo - structure used
// to 'Find' files in a directory specified
// by 'StartPath'. The file search will be
// filtered by a 'FileSelectCriteria' object.
//
// 'FileSelectCriteria' is a FileSelectionCriteria type
// which contains FileNamePatterns strings and
//'FilesOlderThan' or 'FilesNewerThan' date time
// parameters which can be used as a selection
// criteria.
//
type DirectoryTreeInfo struct {
	StartPath          string
	Directories        DirMgrCollection
	FoundFiles         FileMgrCollection
	ErrReturns         []string
	FileSelectCriteria FileSelectionCriteria
}

// CopyToDirectoryTree - Copies an entire directory tree to an alternate location.
// The copy operation includes all files and all directories in the designated directory
// tree.
func (dirTree *DirectoryTreeInfo) CopyToDirectoryTree(baseDir, newBaseDir DirMgr) (DirectoryTreeInfo, error) {

	ePrefix := "DirectoryTreeInfo.CopyToDirectoryTree() "

	newDirTree := DirectoryTreeInfo{}

	if !baseDir.isInitialized {
		return newDirTree, errors.New(ePrefix + "Error: Input parameter 'baseDir' is NOT initialized. It is EMPTY!")
	}

	err2 := baseDir.IsDirMgrValid("")

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix+"Error: Input Parameter 'baseDir' is INVALID! Error='%v'", err2.Error())
	}

	if !newBaseDir.isInitialized {
		return newDirTree, errors.New(ePrefix + "Error: Input parameter 'newBaseDir' is NOT initialized. It is EMPTY!")

	}

	err2 = newBaseDir.IsDirMgrValid("")

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix+"Error: Input Parameter 'newBaseDir' is INVALID! Error='%v'", err2.Error())
	}

	err2 = newBaseDir.MakeDir()

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix+"Error returned from  newBaseDir.MakeDir(). newBaseDir.absolutePath='%v'  Error='%v'", newBaseDir.absolutePath, err2.Error())
	}

	lAry := len(dirTree.Directories.dirMgrs)

	// Make the new Directory Tree
	for i := 0; i < lAry; i++ {

		newDMgr, err2 := dirTree.Directories.dirMgrs[i].SubstituteBaseDir(baseDir, newBaseDir)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned from SubstituteBaseDir(baseDir, newBaseDir). i='%v' Error='%v'", i, err2.Error())
		}

		err2 = newDMgr.MakeDir()

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned fromnewDMgr.MakeDir()  Error='%v'", err2.Error())

		}

		newDirTree.Directories.AddDirMgr(newDMgr)

	}

	lAry = len(dirTree.FoundFiles.fileMgrs)

	for j := 0; j < lAry; j++ {

		fileDMgr, err2 := dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir). Error='%v'", err2.Error())
		}

		newFileMgr, err2 := FileMgr{}.NewFromDirMgrFileNameExt(fileDMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by FileMgr{}.NewFromDirMgrFileNameExt(dMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt) dirTree.FoundFiles.fileMgrs[j].fileNameExt='%v' j='%v' Error='%v'", dirTree.FoundFiles.fileMgrs[j].fileNameExt, j, err2.Error())
		}

		err2 = dirTree.FoundFiles.fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix+"Error returned by fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr) SrcFileName:'%v'  DestFileName:'%v' Error='%v'", dirTree.FoundFiles.fileMgrs[j].fileNameExt, newFileMgr.fileNameExt, err2.Error())

		}

		newDirTree.FoundFiles.AddFileMgr(newFileMgr)
	}

	return newDirTree, nil
}

// DirectoryDeleteFileInfo - structure used
// to delete files in a directory specified
// by 'StartPath'. Deleted files will be selected
// based on 'DeleteFileSelectCriteria' value.
//
// 'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
// type which contains  FileNamePatterns strings and the
// FilesOlderThan or FilesNewerThan date time parameters
// which can be used as file selection criteria.
type DirectoryDeleteFileInfo struct {
	StartPath                string
	Directories              DirMgrCollection
	ErrReturns               []string
	DeleteFileSelectCriteria FileSelectionCriteria
	DeletedFiles             FileMgrCollection
}
