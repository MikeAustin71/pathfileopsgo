package pathfileops

import (
  "errors"
  "fmt"
  "os"
  "time"
)

// FileInfoPlus - Conforms to the os.FileInfo interface. This structure will store
// os.FileInfo information plus additional information related to a file or directory.
//
type FileInfoPlus struct {

  // isFInfoInitialized - Not part of FileInfo interface.
  // 'true' = structure fields have been properly initialized
  isFInfoInitialized bool

  // isDirPathInitialized - Not part of FileInfo interface.
  //   'true' = structure field 'dirPath' has been successfully initialized
  isDirPathInitialized bool

  // CreateTimeStamp - Not part of FileInfo interface.
  // Date time at which this instance of Type 'FileInfoPlus' was initialized
  CreateTimeStamp time.Time

  dirPath      string      // Not part of FileInfo interface. Directory path associated with file name
  fName        string      // FileInfo.Name() base name of the file
  fSize        int64       // FileInfo.Size() length in bytes for regular files; system-dependent for others
  fMode        os.FileMode // FileInfo.Mode() file mode bits
  fModTime     time.Time   // FileInfo.ModTime() file modification time
  isDir        bool        // FileInfo.IsDir() 'true'= this is a directory not a file
  dataSrc      interface{} // FileInfo.Sys() underlying data source (can return nil)
  origFileInfo os.FileInfo
}

// Name - base name of the file.
//  Example:
//              Complete File Name: "newerFileForTest_01.txt"
//    Base Name returned by Name(): "newerFileForTest_01.txt"
//
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
// to another as a portable. Not all bits apply to all systems.
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

// SysAsString - underlying data source. If Sys is
// 'nil', this method will return an empty string.
//
func (fip FileInfoPlus) SysAsString() string {
  if fip.dataSrc == nil {
    return ""
  }

  str := fmt.Sprintf("%v", fip.dataSrc)

  return str
}

// CopyOut - Creates a deep copy of the current FileInfoPlus
// instance and returns it.
//
// This method is NOT part of the FileInfo interface.
//
func (fip *FileInfoPlus) CopyOut() FileInfoPlus {
  newInfo := FileInfoPlus{}

  newInfo.SetName(fip.Name())
  newInfo.SetSize(fip.Size())
  newInfo.SetMode(fip.Mode())
  newInfo.SetModTime(fip.ModTime())
  newInfo.SetIsDir(fip.IsDir())
  newInfo.SetSysDataSrc(fip.Sys())
  _ = newInfo.SetDirectoryPath(fip.DirPath())
  newInfo.isFInfoInitialized = fip.isFInfoInitialized
  newInfo.CreateTimeStamp = fip.CreateTimeStamp
  newInfo.origFileInfo = fip.origFileInfo

  return newInfo
}

// DirPath - Returns the directory path. This field, FileInfoPlus.dirPath,
// is not part of the standard FileInfo interface.
//
// This method is NOT part of the FileInfo interface.
//
func (fip *FileInfoPlus) DirPath() string {
  return fip.dirPath
}

// Equal - Compares two FileInfoPlus objects to determine
// if they are equal.
//
// This method is NOT part of the FileInfo interface.
//
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

// GetOriginalFileInfo - If the FileInfoPlus instance was initialized
// with an os.FileInfo value, this method will return that original
// os.FileInfo value. This is useful for passing parameters to some
// low level go routines such as os.SameFile().
//
// This method is NOT part of the FileInfo interface.
//
func (fip *FileInfoPlus) GetOriginalFileInfo() os.FileInfo {
  return fip.origFileInfo
}

// IsFileInfoInitialized - Returns a boolean value signaling whether
// this instance of FileInfoPlus has been initialized.
//
// A FileInfoPlus instance is properly initialized only if one of the
// following three methods is called:
//
// 1. FileInfoPlus.NewFromFileInfo()
// 2. FileInfoPlus.NewFromPathFileInfo()
// 3. FileInfoPlus.SetIsFInfoInitialized()
//
// This method is NOT part of the FileInfo interface.
//
func (fip *FileInfoPlus) IsFileInfoInitialized() bool {
  return fip.isFInfoInitialized
}

// IsDirectoryPathInitialized - Returns a boolean value signaling whether
// the directory path has been initialized for this instance of the
// FileInfoPlus instance. FYI, the fields FileInfoPlus.isDirPathInitialized
// and FileInfoPlus.dirPath do NOT exist in a standard os.FileInfo object.
//
// A FileInfoPlus directory path is properly initialized only if one of
// the following two methods is called:
//
// 1. FileInfoPlus.NewFromPathFileInfo()
// 2. FileInfoPlus.SetDirectoryPath
//
// This method is NOT part of the FileInfo interface.
//
func (fip *FileInfoPlus) IsDirectoryPathInitialized() bool {
  return fip.isDirPathInitialized
}


// NewFromDirMgrFileInfo - Creates and returns a new FileInfoPlus object
// populated with a Directory Manager (DirMgr) and File Info data (os.FileInfo)
// received from the input parameters 'dMgr' and 'info'.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//  fip, err := FileInfoPlus{}.NewFromDirMgrFileInfo(dMgr, info)
//  fip is now configured as a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromDirMgrFileInfo(
  dMgr DirMgr,
  info os.FileInfo) (FileInfoPlus, error) {

  ePrefix := "FileInfoPlus.NewFromDirMgrFileInfo() "

  err := dMgr.IsDirMgrValid("")

  if err != nil {
    return FileInfoPlus{},
      fmt.Errorf(ePrefix + "ERROR: Input Parameter 'dMgr' is INVALID!\n" +
        "%v", err.Error())
  }

  if info == nil {
    return FileInfoPlus{},
      errors.New(ePrefix + "ERROR: Input Parameter 'info' is nil !\n")
  }

  newInfo := FileInfoPlus{}.NewFromFileInfo(info)

  newInfo.dirPath = dMgr.GetAbsolutePath()

  newInfo.isDirPathInitialized = true

  return newInfo, nil
}

// NewFromFileInfo - Creates and returns a new FileInfoPlus object
// populated with FileInfo data received from the input parameter.
// Notice that this version of the 'NewFromPathFileNameExtStr' method does NOT set the
// Directory path.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//  fip := FileInfoPlus{}.NewFromFileInfo(info)
//  fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromFileInfo(info os.FileInfo) FileInfoPlus {

  if info == nil {
    return FileInfoPlus{}
  }

  newInfo := FileInfoPlus{}

  newInfo.SetName(info.Name())
  newInfo.SetSize(info.Size())
  newInfo.SetMode(info.Mode())
  newInfo.SetModTime(info.ModTime())
  newInfo.SetIsDir(info.IsDir())
  newInfo.SetSysDataSrc(info.Sys())
  newInfo.SetIsFInfoInitialized(true)
  newInfo.origFileInfo = info
  return newInfo
}

// NewFromPathFileInfo - Creates and returns a new FileInfoPlus object
// populated with directory path and FileInfo data received from
// the input parameters.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//  fip, err := FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)
//  fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromPathFileInfo(
  dirPath string,
  info os.FileInfo) (FileInfoPlus, error) {

  ePrefix := "FileInfoPlus.NewFromPathFileInfo() "

  errCode := 0

  errCode,
    _,
    dirPath =  FileHelper{}.isStringEmptyOrBlank(dirPath)

  if errCode < 0 {
    return FileInfoPlus{},
      fmt.Errorf(ePrefix +
        "\nError: Input parameter 'dirPath' is an EMPTY String!\n")
  }

  if info == nil {
    return FileInfoPlus{},
      errors.New(ePrefix + "ERROR: Input Parameter 'info' is nil !\n")
  }

  newInfo := FileInfoPlus{}.NewFromFileInfo(info)

  newInfo.dirPath = dirPath

  newInfo.isDirPathInitialized = true

  return newInfo, nil
}

// SetDirectoryPath - Sets the dirPath field. This
// field is not part of the standard FileInfo data structure.
func (fip *FileInfoPlus) SetDirectoryPath(dirPath string) error {

  fh := FileHelper{}
  errCode := 0

  errCode,
  _,
  dirPath =  fh.isStringEmptyOrBlank(dirPath)

  if errCode < 0 {
    return fmt.Errorf("FileInfoPlus.SetDirectoryPath()\n" +
      "Error: Input parameter 'dirPath' is an EMPTY String!\n")
  }

  dirPath = fh.RemovePathSeparatorFromEndOfPathString(dirPath)

  fip.dirPath = dirPath

  fip.isDirPathInitialized = true

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
    fip.isFInfoInitialized = false
    fip.CreateTimeStamp = time.Time{}
    return
  }

  fip.isFInfoInitialized = true
  fip.CreateTimeStamp = time.Now().Local()
  return
}
