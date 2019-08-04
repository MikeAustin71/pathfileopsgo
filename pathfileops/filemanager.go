package pathfileops

import (
  "bufio"
  "errors"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "strings"
  "sync"
  "time"
)

/*
  This source code file contains type 'FileMgr'.

  The Source Repository for this source code file is :
    https://github.com/MikeAustin71/pathfileopsgo.git

  Dependencies:
  -------------

  Type 'FileMgr' depends on type,'FileHelper'
  which is contained in source code file,
  'filehelper.go' located in this directory.

*/

// FileMgr - This type and its associated methods are used to manage
// organize and control disk files and file permissions.
//
// Dependencies:
//
// Type 'FileMgr' depends on type,'FileHelper'
// which is contained in source code file,
// 'filehelper.go' located in this directory.
//
// To create an instance of type 'FileMgr' use one of the
// 'FileMgr.New' methods.
//
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
  fileBufRdr                      *bufio.Reader
  fileRdrBufSize                  int
  fileBufWriter                   *bufio.Writer
  fileWriterBufSize               int
  isFilePtrOpen                   bool
  fileAccessStatus                FileAccessControl
  actualFileInfo                  FileInfoPlus
  fileBytesWritten                uint64
  buffBytesWritten                uint64
  dataMutex                       sync.Mutex // Used internally to ensure thread safe operations
}

// ChangePermissionMode - This method is a wrapper for os.Chmod().
//
// ChangePermissionMode changes the permissions mode of the named file to mode. If the file is a symbolic
// link, it changes the mode of the link's target. If there is an error, it will be of type
// *PathError.
//
// A different subset of the mode bits are used, depending on the operating system.
//
// On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and ModeSticky are used.
//
// On Windows, the mode must be non-zero but otherwise only the 0200 bit (owner writable) of mode
// is used; it controls whether the file's read-only attribute is set or cleared. attribute. The
// other bits are currently unused. Use mode 0400 for a read-only file and 0600 for a
// readable+writable file.
//
//   For Windows - Eligible permission codes are:
//
//     'modeStr'      Octal
//      Symbolic      Mode Value     File Access
//
//      -r--r--r--     0444          File - read only
//      -rw-rw-rw-     0666          File - read & write
//
// On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive, and ModeTemporary are used.
//
func (fMgr *FileMgr) ChangePermissionMode(mode FilePermissionConfig) error {

  ePrefix := "FileMgr.ChangePermissionMode() "

  fMgrHelpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  filePathDoesExist,
  err := fMgrHelpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: This file does NOT exist!\n"+
      "(FileMgr) File Name:'%v' ",
      fMgr.absolutePathFileName)
  }

  err = mode.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'mode' is invalid!\n"+
      "Error='%v'\n", err.Error())
  }

  fileMode, err := mode.GetCompositePermissionMode()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fMgr.dataMutex.Lock()

  err = os.Chmod(fMgr.absolutePathFileName, fileMode)

  if err != nil {

    fMgr.dataMutex.Unlock()

    return fmt.Errorf(ePrefix+
      "Error returned by os.Chmod(fMgr.absolutePathFileName, fileMode).\n"+
      "fileMode='%v'\nError='%v'\n",
      mode.GetPermissionFileModeValueText(), err.Error())
  }

  err = nil
  _,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "Verify fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()
  return err
}

// CloseThisFile - This method will call the Close()
// method on the current file pointer, FileMgr.filePtr.
//
// In addition, if the file has been opened for Write or
// Read-Write, this method will automatically flush the
// file buffers.
//
func (fMgr *FileMgr) CloseThisFile() error {

  ePrefix := "FileMgr.CloseThisFile() "

  var err error

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.closeFile(fMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// CopyFileMgrByIo - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter 'fMgrDest', an instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method
// will attempt to create it.
//
// One attempt will be made to copy the source file to the specified
// destination file using a technique known as 'io.Copy'. This technique
// will create a new destination file and copy the source file contents
// to that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error will be
// returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
// https://golang.org/pkg/io/#Copy
//
// ----------------------------------------------------------------------------------
//
// Input Parameter:
//
//  fMgrDest  FileMgr - This File Manager type specifies the path and file name
//                      of the destination file to which the source file identified
//                      by the current File Manger will be copied.
//
//                      If the directory path associated with 'fMgrDest' this
//                      method will attempt to create it.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error         - If this method completes successfully, the returned 'error'
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an 'error' Type which contains an
//	                appropriate error message.
//
//	                Note: an error will be returned if the file identified by the
//	                current source File Manager instance does NOT exist.
//
func (fMgr *FileMgr) CopyFileMgrByIo(fMgrDest *FileMgr) error {

  ePrefix := "FileMgr.CopyFileMgrByIo() "

  fMgrHlpr := fileMgrHelper{}

  var err error
  sourceFMgrLabel := "fMgrSource"
  destFMgrLabel := "fMgrDest"

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true,  // create target/destination directory if it does not exist
    false, // delete existing target/destination file
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err == nil {
    err = fMgrHlpr.lowLevelCopyByIO(
      fMgr,
      fMgrDest,
      0,
      ePrefix,
      sourceFMgrLabel,
      destFMgrLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
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
  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err := fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true, //create destination directory
    true,
    ePrefix,
    "fMgrSource",
    "fMgrDest")

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  // See Reference:
  // https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    fMgrDest,
    0, // local buffer size - take default
    ePrefix,
    "fMgrSource",
    "fMgrDest")

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByLink(
      fMgr,
      fMgrDest,
      ePrefix,
      "fMgrSource",
      "fMgrDest")
  }

  fMgr.dataMutex.Unlock()

  return err
}

// CopyFileMgrByIo - Copies the file represented by the current File
// Manager instance to a location specified by a destination input
// parameter 'fMgrDest', another instance of type FileMgr.
//
// Note that if the destination directory does not exist, this method
// will attempt to create it.
//
// One attempt will be made to copy the source file to the specified
// destination file using a technique known as 'io.Copy'. This technique
// will create a new
// destination file and copies the source file contents to that new destination
// file.
//
// If this attempted 'io.Copy' operation fails, and error will be returned.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//
// Input Parameter:
//
//  fMgrDest  FileMgr - This File Manager type specifies the path and file name
//                      of the destination file to which the source file identified
//                      by the current File Manger will be copied.
//
//                      If the directory path associated with 'fMgrDest' this
//                      method will attempt to create it.
//
//  bufferSize    int - The size in bytes of a local buffer which will be used
//                      to copy the source File Manager (fMgr) to the destination
//                      File Manager (fMgrDest). This is useful when copying large
//                      files.
//
//                      If this value is set equal or less than zero, the default
//                      default internal buffer will be used. Reference:
//                        https://golang.org/pkg/io/#CopyBuffer
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error         - If this method completes successfully, the returned 'error'
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an 'error' Type which contains an
//	                appropriate error message.
//
//	                Note: an error will be returned if the file identified by the
//	                current source File Manager instance does NOT exist.
//
func (fMgr *FileMgr) CopyFileMgrByIoWithBuffer(
  fMgrDest *FileMgr, bufferSize int) error {

  ePrefix := "FileMgr.CopyFileMgrByIo() "

  sourceFMgrLabel := "fMgrSource"

  destFMgrLabel := "fMgrDest"

  fMgrHlpr := fileMgrHelper{}

  var err error

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true,  // create target/destination directory if it does not exist
    false, // delete target/destination file if it exists
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    fMgrDest,
    bufferSize,
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  fMgr.dataMutex.Unlock()

  return err
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
// By creating a 'linked' file, changing the contents of one file will be
// reflected in the second. The two linked files are 'mirrors' of each
// other.
//
// Consider using FileMgr.CopyFileMgrByIo() if the 'mirror' feature causes
// problems.
//
// If the 'Hard Link' copy operation fails, and error will be returned.
//
//   Reference:
//    https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//    https://golang.org/pkg/os/#Link
//
func (fMgr *FileMgr) CopyFileMgrByLink(fMgrDest *FileMgr) error {

  ePrefix := "FileMgr.CopyFileMgrByLink() "

  fMgrHlpr := fileMgrHelper{}

  sourceFMgrLabel := "fMgrSource"

  destFMgrLabel := "fMgrDest"

  fMgr.dataMutex.Lock()

  err := fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true, // create target/destination directory
    true, // delete existing target/destination file
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    fMgrDest,
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  fMgr.dataMutex.Unlock()

  return err
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

  fMgrHlpr := fileMgrHelper{}

  sourceFMgrLabel := "fMgrSource"
  destFMgrLabel := "fMgrDestDir"

  fMgr.dataMutex.Lock()

  err := fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true,
    true,
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    fMgrDest,
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByIO(
      fMgr,
      fMgrDest,
      0,
      ePrefix,
      sourceFMgrLabel,
      destFMgrLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
}

// CopyFileMgrWithBuffer - Performs a Copy By IO operation which copies
// the file identified by the current File Manager to the file specified
// by the input parameter 'fMgrDest'. 'fMgrDest' is another instance of
// 'FileMgr'.
//
// This method differs from a similar method, 'FileMgr.CopyFileMgrByIo()'
// in that the calling function must supply a local copy buffer of type
// []byte. The copy operation stages through the provided buffer (if one
// is required) rather than allocating a temporary internal buffer. This
// approach can be useful when copying a series of large files.
//
// Internally, this method calls io.CopyBuffer().
//
//   Reference:
//     https://golang.org/pkg/io/#CopyBuffer
//     https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------------------------
//
// Input Parameter:
//
//  fMgrDest       FileMgr - An instance of FileMgr which identifies the destination
//                           file to which the source file identified by the current
//                           FileMgr will be copied.
//
//  localCopyBuffer []byte - A byte array which will be used as an internal buffer through
//                           which to stage the copy operation. If this buffer is 'nil' or
//                           if this is a zero length byte array, an error will be returned.
//
// ------------------------------------------------------------------------------------------
//
// Return Value:
//
//  error - If this method completes successfully, the returned 'error'
//          Type is set equal to 'nil'. If an error condition is encountered,
//          this method will return an 'error' Type which contains an
//          appropriate error message.
//
//          Note: an error will be returned if the source file identified by
//          the current FileMgr instance does NOT exist. Likewise, an error
//          will also be triggered if input parameter 'localCopyBuffer' is
//          'nil' or if it is a zero length byte array.
//
func (fMgr *FileMgr) CopyFileMgrWithBuffer(
  fMgrDest *FileMgr,
  localCopyBuffer []byte) error {

  ePrefix := "FileMgr.CopyFileMgrWithBuffer() "

  fMgrHlpr := fileMgrHelper{}

  var err error
  sourceFMgrLabel := "fMgrSource"
  destFMgrLabel := "fMgrDest"

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    fMgrDest,
    true,  // create target/destination directory if it does not exist
    false, // delete existing target/destination file
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err == nil {
    err = fMgrHlpr.lowLevelCopyByIOBuffer(
      fMgr,
      fMgrDest,
      localCopyBuffer,
      ePrefix,
      sourceFMgrLabel,
      destFMgrLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
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
// using a technique known as 'Copy By IO'. This technique create a new destination
// file and copies the source file contents to that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error will be returned.
//
//  Reference:
//   https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//   https://golang.org/pkg/io/#Copy
//
func (fMgr *FileMgr) CopyFileStrByIo(dstPathFileNameExt string) error {

  ePrefix := "FileMgr.CopyFileStrByIo() "

  fMgrDest, err := FileMgr{}.New(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"Invalid input parameter 'dstPathFileNameExt'\n"+
      "Error returned by FileMgr{}.New(dstPathFileNameExt)\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  fMgrSourceLabel := "fMgrSource"

  fMgrDestLabel := "fMgrDest"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true,  // create destination directory
    false, // delete destination file if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    &fMgrDest,
    0, // Local Buffer Size - Use Default
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  fMgr.dataMutex.Unlock()

  return err
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

  fMgrDest, err := FileMgr{}.New(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"Invalid input parameter 'dstPathFileNameExt'\n"+
      "Error returned by FileMgr{}.New(dstPathFileNameExt)\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  fMgrSourceLabel := "fMgrSource"

  fMgrDestLabel := "fMgrDest"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // Create Destination Directory
    true, // Delete Destination File if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    &fMgrDest,
    0,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByLink(
      fMgr,
      &fMgrDest,
      ePrefix,
      fMgrSourceLabel,
      fMgrDestLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
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

  fMgrDest, err := FileMgr{}.New(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"Invalid input parameter 'dstPathFileNameExt'\n"+
      "Error returned by FileMgr{}.New(dstPathFileNameExt)\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  fMgrSourceLabel := "fMgrSource"

  fMgrDestLabel := "fMgrDest"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // create destination directory
    true, // delete destination file if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    &fMgrDest,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  fMgr.dataMutex.Unlock()

  return err
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

  fMgrDest, err := FileMgr{}.New(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"Invalid input parameter 'dstPathFileNameExt'\n"+
      "Error returned by FileMgr{}.New(dstPathFileNameExt)\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  fMgrSourceLabel := "fMgrSource"

  fMgrDestLabel := "fMgrDest"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // Create Destination Directory
    true, // Delete Destination File if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    &fMgrDest,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByIO(
      fMgr,
      &fMgrDest,
      0,
      ePrefix,
      fMgrSourceLabel,
      fMgrDestLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
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

  fMgrHlpr := fileMgrHelper{}

  err := dir.IsDirMgrValid(ePrefix + "Input Parameter 'dir' ")

  if err != nil {
    return err
  }

  sourceFMgrLabel := "fMgrSource"
  destFMgrLabel := "fMgrDestDir"

  fMgr.dataMutex.Lock()

  fMgrDest, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dir, fMgr.fileNameExt)\n"+
      "dir='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true,  // create target/destination directory if it does not exist
    false, // delete existing target/destination file
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    &fMgrDest,
    0, // Local Buffer Size = default
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  fMgr.dataMutex.Unlock()

  return err
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
// technique known as 'io.Copy' or 'Copy by IO'.
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
  fMgrHlpr := fileMgrHelper{}

  err := dir.IsDirMgrValid(ePrefix + "Input Parameter 'dir' ")

  if err != nil {
    return err
  }

  fMgrSourceLabel := "fMgrSource"
  fMgrDestLabel := "fMgrDestDir"

  fMgr.dataMutex.Lock()

  fMgrDest, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dir, fMgr.fileNameExt)\n"+
      "dir='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // Create Destination Directory
    true, // Delete Destination File if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByIO(
    fMgr,
    &fMgrDest,
    0,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByLink(
      fMgr,
      &fMgrDest,
      ePrefix,
      fMgrSourceLabel,
      fMgrDestLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
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

  err := dir.IsDirMgrValid(ePrefix + "Input Parameter 'dir' ")

  if err != nil {
    return err
  }

  fMgrSourceLabel := "fMgrSource"
  fMgrDestLabel := "fMgrDestDir"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  fMgrDest, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dir, fMgr.fileNameExt)\n"+
      "dir='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // Create Destination Directory
    true, // Delete Destination File if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    &fMgrDest,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  fMgr.dataMutex.Unlock()

  return err
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
  err := dir.IsDirMgrValid(ePrefix + "Input Parameter 'dir' ")

  if err != nil {
    return err
  }

  fMgrSourceLabel := "fMgrSource"
  fMgrDestLabel := "fMgrDestDir"

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  fMgrDest, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt("+
      "dir, fMgr.fileNameExt)\n"+
      "dir='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  err = fMgrHlpr.copyFileSetup(
    fMgr,
    &fMgrDest,
    true, // Create Destination Directory
    true, // Delete Destination File if it exists
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return err
  }

  err = fMgrHlpr.lowLevelCopyByLink(
    fMgr,
    &fMgrDest,
    ePrefix,
    fMgrSourceLabel,
    fMgrDestLabel)

  if err != nil {
    err = fMgrHlpr.lowLevelCopyByIO(
      fMgr,
      &fMgrDest,
      0,
      ePrefix,
      fMgrSourceLabel,
      fMgrDestLabel)
  }

  fMgr.dataMutex.Unlock()

  return err
}

// CopyFromStrings - Copies a source file to a destination file. The source and
// destination files are identified by two strings specifying path, file name,
// and file extensions. The copy operation employs the 'Copy By IO' technique
// (a.k.a. 'io.Copy').
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
// https://golang.org/pkg/io/#Copy
//
// ----------------------------------------------------------------------------------
//
// Input Parameter:
//
// sourcePathFileNameExt  string - A text string which identifies the path, file name
//                                 and file extension of the source file which will be
//                                 used in the copy operation.
//
// destPathFileNameExt    string - A text string which identifies the path, file name
//                                 and file extension of the destination file which will
//                                 be used in the copy operation.
//
//
// ------------------------------------------------------------------------------------------
//
// Return Values:
//
//  srcFMgr  FileMgr - If successful, this method returns a File Manager instance which
//                     identifies the source file used in the copy operation.
//
//  destFMgr FileMgr - If successful, this method returns a File Manager instance which
//                     identifies the destination file used in the copy operation.
//
//  err        error - If this method completes successfully, the returned 'error'
//                     Type is set equal to 'nil'. If an error condition is encountered,
//                     this method will return an 'error' Type which contains an
//                     appropriate error message.
//
//                     Note: an error will be returned if the file identified by the
//                     input parameter 'sourcePathFileNameExt' does NOT exist.
//
// ------------------------------------------------------------------------------------------
//
// Usage:
//
//     srcFileMgr, destFileMgr, err := FileMgr{}.CopyFromStrings(
//                                     "../Dir1/srcFile",
//                                     "../Dir2/destFile")
//
func (fMgr FileMgr) CopyFromStrings(
  sourcePathFileNameExt,
  destPathFileNameExt string) (fMgrSrc, fMgrDest FileMgr, err error) {

  ePrefix := "FileMgr.CopyFromStrings() "
  fMgrDest = FileMgr{}
  var err2 error

  fMgrSrc, err2 = FileMgr{}.New(sourcePathFileNameExt)

  if err2 != nil {
    fMgrSrc = FileMgr{}

    err = fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.New(sourcePathFileNameExt)\n"+
      "sourcePathFileNameExt='%v'\nError='%v'\n",
      sourcePathFileNameExt, err2.Error())

    return fMgrSrc, fMgrDest, err
  }

  fMgrDest, err2 = FileMgr{}.New(destPathFileNameExt)

  if err2 != nil {
    fMgrSrc = FileMgr{}
    fMgrDest = FileMgr{}
    err = fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.New(destPathFileNameExt)\n"+
      "destPathFileNameExt='%v'\nError='%v'\n",
      destPathFileNameExt, err2.Error())
    return fMgrSrc, fMgrDest, err
  }

  fMgrHlpr := fileMgrHelper{}

  sourceFMgrLabel := "fMgrSource"
  destFMgrLabel := "fMgrDest"

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.copyFileSetup(
    &fMgrSrc,
    &fMgrDest,
    true,  // create target/destination directory if it does not exist
    false, // delete existing target/destination file
    ePrefix,
    sourceFMgrLabel,
    destFMgrLabel)

  if err == nil {
    err = fMgrHlpr.lowLevelCopyByIO(
      &fMgrSrc,
      &fMgrDest,
      0,
      ePrefix,
      sourceFMgrLabel,
      destFMgrLabel)
  }

  if err != nil {
    fMgrSrc = FileMgr{}
    fMgrDest = FileMgr{}
  }

  fMgr.dataMutex.Unlock()

  return fMgrSrc, fMgrDest, err
}

// CopyIn - Copies data from an incoming FileMgr object
// into the current FileMgr object.
//
// Note: Internal file pointer will not be copied.
//
func (fMgr *FileMgr) CopyIn(fmgr2 *FileMgr) {

  if fmgr2 == nil {
    fmgr2 = &FileMgr{}
  }

  fMgr.dataMutex.Lock()
  fmgr2.dataMutex.Lock()


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
  fMgr.filePtr = nil
  fMgr.isFilePtrOpen = false
  fMgr.fileAccessStatus = fmgr2.fileAccessStatus.CopyOut()
  fMgr.actualFileInfo = fmgr2.actualFileInfo.CopyOut()
  fMgr.fileBytesWritten = 0
  fMgr.buffBytesWritten = 0
  fMgr.fileRdrBufSize = fmgr2.fileRdrBufSize
  fMgr.fileWriterBufSize = fmgr2.fileWriterBufSize

  fmgr2.dataMutex.Unlock()
  fMgr.dataMutex.Unlock()

  return
}

// CopyOut - Duplicates the file information in the current
// FileMgr object and returns it as a new FileMgr object.
//
// Note: Internal File Pointer will not be copied.
//
func (fMgr *FileMgr) CopyOut() FileMgr {

  fMgr.dataMutex.Lock()

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
  fmgr2.filePtr = nil
  fmgr2.isFilePtrOpen = false
  fmgr2.fileAccessStatus = fMgr.fileAccessStatus.CopyOut()
  fmgr2.actualFileInfo = fMgr.actualFileInfo.CopyOut()
  fmgr2.fileBytesWritten = 0
  fmgr2.buffBytesWritten = 0
  fmgr2.fileRdrBufSize = fMgr.fileRdrBufSize
  fmgr2.fileWriterBufSize = fMgr.fileWriterBufSize

  fMgr.dataMutex.Unlock()

  return fmgr2
}

// CreateDir - Creates the directory previously configured
// for this file manager instance.
//
func (fMgr *FileMgr) CreateDir() error {

  ePrefix := "FileMgr.CreateDir() "
  var err error

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.createDirectory(fMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// CreateDirAndFile - Performs two operations:
//
//  If the home directory does not currently exist, this method
//  will first create the directory tree containing the file.
//
// Next, the file will be created. If the file previously exists,
// it will be truncated.
//
// Note that if this method successfully creates the file, a
// File Pointer (*File) will be stored in the FileMgr field
// fMgr.filePtr. Be sure to close the File Pointer when
// finished with it. See FileMgr.CloseThisFile().
//
func (fMgr *FileMgr) CreateDirAndFile() error {

  ePrefix := "FileMgr:CreateDirAndFile() "
  var err error

  fMgr.dataMutex.Lock()

  fMgrHlpr := fileMgrHelper{}

  err = fMgrHlpr.createFile(
    fMgr,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// CreateThisFile - Creates the File identified by FileMgr
// (FileMgr.absolutePathFileName).
//
// If the directory in the path file name designation does not exist, this
// method will throw an error.
//
// See Method CreateDirAndFile() which will create both the directory and the file
// as required.
//
// Note that if the file is actually created, the returned file pointer (*File)
// is stored in the FileMgr field, fMgr.filePtr. Be sure to 'close' the File Pointer
// when finished with it. See FileMgr.CloseThisFile()
//
func (fMgr *FileMgr) CreateThisFile() error {

  ePrefix := "FileMgr:CreateThisFile() "

  fMgrHlpr := fileMgrHelper{}
  var err error

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.createFile(
    fMgr,
    false,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// DeleteThisFile - Deletes the file identified by FileMgr.absolutePathFileName
// in the current FileHelper structure.

func (fMgr *FileMgr) DeleteThisFile() error {

  ePrefix := "FileMgr.DeleteThisFile() "
  var err error

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.deleteFile(fMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// DoesFileExist - returns 'true' if the subject FileMgr file does
// in fact exist. If the file does NOT exist, a boolean value of
// 'false' is returned.
//
// This method uses os.Stat() to test for the existence of a file
// path. If a non-path error is returned by os.Stat(), it is
// ignored and the file path is classified as 'Does Not Exist'.
//
//
// This is very similar to 'FileMgr.DoesThisFileExist()'.
// However, unlike this method, 'FileMgr.DoesThisFileExist()'
// will return a non-path error. This method only returns
// a boolean value signaling whether the file path exists.
//
// If this method encounters a non-path error or if the current
// FileMgr instance is invalid, a boolean value of 'false'
// will be returned.
//
func (fMgr *FileMgr) DoesFileExist() bool {

  ePrefix := "FileMgr.DoesFileExist() "

  var err error
  pathFileDoesExist := false

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  pathFileDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return false
  }

  return pathFileDoesExist
}

// DoesThisFileExist - Returns a boolean value
// designating whether the file specified by the
// current FileMgr.absolutePathFileName field
// exists.
func (fMgr *FileMgr) DoesThisFileExist() (fileDoesExist bool, nonPathError error) {

  ePrefix := "FileMgr.DoesThisFileExist() "
  fileDoesExist = false
  nonPathError = nil
  fMgrHelpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  fileDoesExist,
    nonPathError = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if nonPathError != nil {
    fileDoesExist = false
  }

  fMgr.dataMutex.Unlock()

  return fileDoesExist, nonPathError
}

// Equal - Compares a second FileHelper data structure
// to the current FileHelper data structure and returns
// a boolean value indicating whether they are equal
// in all respects.
func (fMgr *FileMgr) Equal(fmgr2 *FileMgr) bool {

  if fmgr2 == nil {
    fmgr2 = &FileMgr{}
  }

  result := true

  fMgr.dataMutex.Lock()
  fmgr2.dataMutex.Lock()

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
    fMgr.isFilePtrOpen != fmgr2.isFilePtrOpen ||
    fMgr.fileRdrBufSize != fmgr2.fileRdrBufSize ||
    fMgr.fileWriterBufSize != fmgr2.fileWriterBufSize {

    result = false
  }

  if result==true &&
    !fMgr.fileAccessStatus.Equal(&fmgr2.fileAccessStatus) {
    result = false
  }

  if result ==true &&
    !fMgr.dMgr.Equal(&fmgr2.dMgr) {
    result = false
  }

  if result == true &&
    !fMgr.actualFileInfo.Equal(&fmgr2.actualFileInfo) {
    result = false
  }


  fmgr2.dataMutex.Unlock()
  fMgr.dataMutex.Unlock()

  return result
}

// EqualAbsPaths - Returns 'true' if both the current File Manager
// and the input File Manager ('fmgr2') have the same file paths.
//
// In other words, this method answers the question, 'Do Both Files
// have the same directory?'.
//
// The path comparisons are case insensitive. This means that both
// paths will be converted to lower case before making the comparison.
//
// Also, the path comparison will be performed on the absolute paths
// associated with the two File Managers.
//
// If the file paths are NOT equal, this method returns 'false.
//
// NOTE: This method will NOT test the equality of file names and
// extensions. ONLY the file paths (directories) will be compared.
//
func (fMgr *FileMgr) EqualAbsPaths(fmgr2 *FileMgr) bool {

  if fmgr2 == nil {
    fmgr2 = &FileMgr{}
  }

  fDirMgr := fMgr.GetDirMgr()

  fDirMgr2 := fmgr2.GetDirMgr()

  return fDirMgr.EqualAbsPaths(&fDirMgr2)
}

// EqualFileNameExt - Returns 'true' if both the current File Manager
// and the input File Manager ('fmgr2') have the same file name and file
// extension.
//
// The File Name and File Extension comparisons are case insensitive. This
// means that both file name and file extension will be converted to lower
// case before making the comparision.
//
//	Example: xray.txt is considered equal to XRAY.TXT
//
// If the either the File Name or File Extension are NOT equal, this method
// will return 'false'.
//
// NOTE: This method will NOT test the equality of the file paths or directories.
// ONLY the file name and file extension are tested for equaltiy.
//
func (fMgr *FileMgr) EqualFileNameExt(fmgr2 *FileMgr) bool {

  if fmgr2 == nil {
    fmgr2 = &FileMgr{}
  }

  f1FileName := ""
  f2FileName := ""
  f1FileExt := ""
  f2FileExt := ""

  fMgr.dataMutex.Lock()
  fmgr2.dataMutex.Lock()

  f1FileName = fMgr.fileName
  f2FileName = fmgr2.fileName
  f1FileExt = fMgr.fileExt
  f2FileExt = fmgr2.fileExt

  fmgr2.dataMutex.Unlock()
  fMgr.dataMutex.Unlock()

  f1FileName = strings.ToLower(f1FileName)

  f2FileName = strings.ToLower(f2FileName)

  if f1FileName != f2FileName {
    return false
  }

  f1FileExt = strings.ToLower(f1FileExt)
  f2FileExt = strings.ToLower(f2FileExt)

  if f1FileExt != f2FileExt {
    return false
  }

  return true
}

// EqualPathFileNameExt - Returns 'true' if both the current File Manager
// and the input File Manager ('fmgr2') have the same absolute path,
// file name and file extension.
//
// The string comparisons are case insensitive. This means that the paths,
// file names and file extensions will all be converted to lower case
// before making the comparision.
//
//	Example: d:\dir1\xray.txt is considered equal to D:\DIR1\XRAY.TXT
//
// If the path, file name or file extensions are NOT equal, this method
// will return 'false'.
//
func (fMgr *FileMgr) EqualPathFileNameExt(fmgr2 *FileMgr) bool {

  if fmgr2 == nil {
    fmgr2 = &FileMgr{}
  }

  f1AbsolutePathFileName := ""
  f2AbsolutePathFileName := ""

  fMgr.dataMutex.Lock()
  fmgr2.dataMutex.Lock()

  f1AbsolutePathFileName = fMgr.absolutePathFileName

  f2AbsolutePathFileName = fmgr2.absolutePathFileName

  fmgr2.dataMutex.Unlock()
  fMgr.dataMutex.Unlock()

  f1AbsolutePathFileName = strings.ToLower(f1AbsolutePathFileName)
  f2AbsolutePathFileName = strings.ToLower(f2AbsolutePathFileName)

  if f1AbsolutePathFileName != f2AbsolutePathFileName {
    return false
  }

  return true
}

// Empty - resets all data fields in the FileMgr structure to
// their uninitialized or zero state.
func (fMgr *FileMgr) Empty() {

  fMgrHlpr := fileMgrHelper{}
  var err error

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.emptyFileMgr(fMgr, "FileMgr.Empty() ")

  fMgr.dataMutex.Unlock()

  if err != nil {
    panic(err.Error())
  }

  return
}

// FlushBytesToDisk - After Writing bytes to a file, use this
// method to commit the contents of the current file to
// stable storage.
func (fMgr *FileMgr) FlushBytesToDisk() error {

  ePrefix := "FileMgr.FlushBytesToDisk() "

  var err error
  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.flushBytesToDisk(fMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// GetAbsolutePath - Returns the absolute path
// for the current File Manager instance.
//
// Note: The file name and file extension are NOT included.
// Only the absolute path is returned as a 'string'.
//
func (fMgr *FileMgr) GetAbsolutePath() string {

  return fMgr.dMgr.GetAbsolutePath()
}

// GetAbsolutePathFileName - Returns the absolute path,
// file name and file extension for the current File Manager
// instance.
//
// The returned path/file name string may consist of both
// upper and lower case characters.
//
// See the companion method GetAbsolutePathFileNameLc() which
// returns a path/file name string consisting entirely of
// lower case characters.
//
func (fMgr *FileMgr) GetAbsolutePathFileName() string {

  var absolutePathFileName string

  fMgr.dataMutex.Lock()

  if !fMgr.isInitialized {

    absolutePathFileName = ""

  } else{

    absolutePathFileName = fMgr.absolutePathFileName
  }

  fMgr.dataMutex.Unlock()

  return absolutePathFileName
}

// GetAbsolutePathFileNameLc - Returns the absolute path,
// file name and file extension for the current File Manager
// instance.
//
// The returned path/file name string will consist entirely
// of lower case characters.
//
// See companion method GetAbsolutePathFileName() which
// returns the same path/file name string with upper and
// lower case characters.
//
func (fMgr *FileMgr) GetAbsolutePathFileNameLc() string {

  var absolutePathFileName string

  fMgr.dataMutex.Lock()

  if !fMgr.isInitialized {
    absolutePathFileName = ""
  } else{

    absolutePathFileName = strings.ToLower(fMgr.absolutePathFileName)
  }

  fMgr.dataMutex.Unlock()

  return absolutePathFileName
}

// GetBufioReader - Returns a pointer to the internal bufio.Reader,
// FileMgr.fileBufRdr. This pointer is initialized when the file is
// opened for Read or Read-Write operations.
//
// Be advised that if the file identified by the current FileMgr
// instance has not been opened the returned pointer to
// 'fMgr.fileBufRdr' may be nil.
//
func (fMgr *FileMgr) GetBufioReader() *bufio.Reader {

  var fileBufRdr *bufio.Reader

  fMgr.dataMutex.Lock()

  fileBufRdr = fMgr.fileBufRdr

  fMgr.dataMutex.Unlock()

  return fileBufRdr
}

// GetBufioWriter - Returns a pointer to the internal bufio.Writer,
// 'FileMgr.fileBufWriter'. This pointer is initialized when the file
// is opened for Write or Read-Write operations.
//
// Be advised that if the file identified by the current FileMgr
// instance has not been opened, the returned pointer to
// 'fMgr.fileBufWriter' may be nil.
//
func (fMgr *FileMgr) GetBufioWriter() *bufio.Writer {

  var bufWriter *bufio.Writer

  fMgr.dataMutex.Lock()

  bufWriter = fMgr.fileBufWriter

  fMgr.dataMutex.Unlock()

  return bufWriter
}

// GetDirMgr - returns a deep copy of the Directory
// Manager for this FileMgr instance.
func (fMgr *FileMgr) GetDirMgr() DirMgr {
  return fMgr.dMgr.CopyOut()
}

// GetFileBytesWritten - Returns the sum of private member variables,
// 'FileMgr.buffBytesWritten' + 'FileMgr.fileBytesWritten'.
//
// These variables records the number of bytes written to the FileMgr's
// target file since it was opened with 'Write' or 'Read-Write' permissions.
//
func (fMgr *FileMgr) GetFileBytesWritten() uint64 {
  var bytesWritten uint64

  fMgr.dataMutex.Lock()

  bytesWritten = fMgr.buffBytesWritten + fMgr.fileBytesWritten

  fMgr.dataMutex.Unlock()

  return bytesWritten
}

// GetFileExt() - returns a string containing the File Extension for this
// File Manager instance.
//
// IMPORTANT:
// The returned file extension will contain the preceding dot separator.
//
//    Example:
//            File Name Plus Extension: "newerFileForTest_01.txt"
//             Returned File Extension: ".txt"
//
//            File Name Plus Extension: "newerFileForTest_01"
//             Returned File Extension: ""
//
func (fMgr *FileMgr) GetFileExt() string {
  fileExt := ""

  fMgr.dataMutex.Lock()

  if fMgr.isInitialized == false {
    fileExt = ""
  } else {
    fileExt = fMgr.fileExt
  }

  fMgr.dataMutex.Unlock()

  return fileExt
}

// GetFileInfo - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on the specific file identified
// by FileMgr.absolutePathFileName.
//
// An error will be triggered if the file path does NOT exist!
//
//  type FileInfo interface {
// 	 Name() string       // base name of the file
// 	 Size() int64        // length in bytes for regular files; system-dependent for others
// 	 Mode() FileMode     // file mode bits
// 	 ModTime() time.Time // modification time
// 	 IsDir() bool        // abbreviation for Mode().IsDir()
// 	 Sys() interface{}   // underlying data source (can return nil)
//  }
func (fMgr *FileMgr) GetFileInfo() (fInfo os.FileInfo, err error) {

  fInfo = nil
  err = nil
  filePathDoesExist := false
  ePrefix := "FileMgr.GetFileInfo() "

  fMgrHelpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    fInfo = nil
  } else if filePathDoesExist {
    err = nil
    fInfo = fMgr.actualFileInfo.GetOriginalFileInfo()
  } else {
    fInfo = nil
    err =
      fmt.Errorf(ePrefix+
        "The current FileMgr file DOES NOT EXIST!\n"+
        "File='%v'\n", fMgr.absolutePathFileName)
  }

  fMgr.dataMutex.Unlock()

  return fInfo, err
}

// GetFileInfoPlus - Returns a FileInfoPlus instance containing
// os.FileInfo and other data on the current FileManager instance.
//
// An error will be triggered if the file path does NOT exist!
//
func (fMgr *FileMgr) GetFileInfoPlus() (fInfo FileInfoPlus, err error) {


  fInfo = FileInfoPlus{}
  err = nil
  ePrefix := "FileMgr.GetFileInfoPlus() "

  fMgrHelpr := fileMgrHelper{}
  filePathDoesExist := false

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    fInfo = FileInfoPlus{}
  } else if filePathDoesExist {
    fInfo = fMgr.actualFileInfo.CopyOut()
    err = nil
  } else {
    fInfo = FileInfoPlus{}
    err = fmt.Errorf(ePrefix+
      "Error: File Manager file DOES NOT EXIST!\n"+
      "File='%v'\n", fMgr.absolutePathFileName)
  }

  fMgr.dataMutex.Unlock()

  return fInfo, err
}

// GetFileModTime - Returns the time of the last file modification as a
// type, 'time.Time'.
//
// If the file does NOT exist, an error is returned.
//
func (fMgr *FileMgr) GetFileModTime() (time.Time, error) {

  ePrefix := "FileMgr.GetFileModTime() "

  fMgrHelpr := fileMgrHelper{}
  var err error
  filePathDoesExist := false
  modTime := time.Time{}

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err == nil && filePathDoesExist {
    modTime = fMgr.actualFileInfo.ModTime()
  }

  fMgr.dataMutex.Unlock()

  if err == nil && !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Current FileMgr file DOES NOT EXIST!\n"+
      "FileMgr='%v'\n", fMgr.absolutePathFileName)
  }

  return modTime, err
}

// GetFileModTimeStr - Returns the time of the last file modification as
// a string. If the file does NOT exist, an error is returned.
//
//
// ----------------------------------------------------------------------------------
//
// Input Parameters:
//
//  timeFormat  string - A format string used to format the last modification time for
//                       the file identified by the current File Manager (FileMgr) instance.
//
//                       If the string is empty ("") or if the time format is invalid, the
//                       method will automatically format the time using the default format,
//                       "2019-03-12 21:49:00:00".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	string        - The time at which the current file was last modified formatted
//	                as a string. The time format is determined by input parameter
//	                'timeFormat'. If 'timeFormat' is empty or if 'timeFormat' is an
//	                invalid format, the default format "2019-03-12 21:49:00:00" will
//	                be substituted.
//
//	error         - If this method completes successfully, the returned error Type is
//	                set equal to 'nil'. If an error condition is encountered, this
//	                method will return an error Type which contains an appropriate
//	                error message.
//
//	                Note: an error will be returned if the file identified by the current
//	                File Manager instance does NOT exist.
//
func (fMgr *FileMgr) GetFileModTimeStr(timeFormat string) (string, error) {

  ePrefix := "FileMgr.GetFileModTimeStr() "

  fMgrHelpr := fileMgrHelper{}
  var err error
  filePathDoesExist := false
  modTime := time.Time{}

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err == nil && filePathDoesExist {
    modTime = fMgr.actualFileInfo.ModTime()
  }

  fMgr.dataMutex.Unlock()

  defaultFmt := "2006-01-02 15:04:05 -0700 MST"

  if err == nil && !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Current FileMgr file DOES NOT EXIST!\n"+
      "FileMgr='%v'\n", fMgr.absolutePathFileName)
  }

  if timeFormat == "" {
    timeFormat = defaultFmt
  }

  tStr := modTime.Format(timeFormat)

  if tStr == timeFormat {
    tStr = modTime.Format(defaultFmt)
  }

  return tStr, err
}

// GetFileName - returns the file name for this
// File Manager.
//
//    Example:
//
//            File Name Plus Extension: "newerFileForTest_01.txt"
//                  Returned File Name: "newerFileForTest_01"
//
//            File Name Plus Extension: "newerFileForTest_01"
//                  Returned File Name: "newerFileForTest_01"
//
func (fMgr *FileMgr) GetFileName() string {

  fileName := ""

  fMgr.dataMutex.Lock()

  if fMgr.isInitialized == false {
    fileName = ""
  } else {
    fileName = fMgr.fileName
  }

  fMgr.dataMutex.Unlock()

  return fileName
}

// GetFileNameExt - Returns a string containing the
// combination of file name and file extension configured
// for this File Manager instance
//
//    Example:
//
//            File Name Plus Extension: "newerFileForTest_01.txt"
//   Returned File Name Plus Extension: "newerFileForTest_01.txt"
//
//            File Name Plus Extension: "newerFileForTest_01"
//   Returned File Name Plus Extension: "newerFileForTest_01"
//
func (fMgr *FileMgr) GetFileNameExt() string {

  fileNameExt := ""

  fMgr.dataMutex.Lock()

  if fMgr.isInitialized == false {
    fileNameExt = ""

  } else {
    fileNameExt = fMgr.fileNameExt
  }

  fMgr.dataMutex.Unlock()

  return fileNameExt
}

// GetFilePermissionConfig - Returns a FilePermissionConfig instance encapsulating
// all of the permission information associated with this file.
//
// If the file does NOT exist, this method will return an error.
//
func (fMgr *FileMgr) GetFilePermissionConfig() (fPerm FilePermissionConfig, err error) {

  ePrefix := "FileMgr.GetFilePermissionConfig() "

  fPerm = FilePermissionConfig{}
  err = nil
  var filePathDoesExist bool
  var err2 error
  fMgrHelpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "FileMgr.absolutePathFileName")

  if err==nil &&
      !filePathDoesExist {
      err =
      errors.New(ePrefix +
        "The current FileMgr file DOES NOT EXIST!\n")
  }

  if err == nil {

    fPerm, err2 = FilePermissionConfig{}.NewByFileMode(fMgr.actualFileInfo.Mode())

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          "\n%v\n", err2.Error())
    }
  }

  fMgr.dataMutex.Unlock()

  return fPerm, err
}

// GetFilePermissionTextCodes - If the current file exists on disk,
// this method will return the File Permission Codes, otherwise known
// as the unix permission bits, in the form of a 10-character string.
//
// If the file does NOT exist, this method will return an error.
//
// Examples:
//
//   Permission
//     Text        Octal            File Access
//     Codes      Notation          Permission Descriptions
//
//   ----------     0000             File - no permissions
//   -rwx------     0700             File - read, write, & execute only for owner
//   -rwxrwx---     0770             File - read, write, & execute for owner and group
//   -rwxrwxrwx     0777             File - read, write, & execute for owner, group and others
//   ---x--x--x     0111             File - execute
//   --w--w--w-     0222             File - write only
//   --wx-wx-wx     0333             File - write & execute
//   -r--r--r--     0444             File - read only
//   -r-xr-xr-x     0555             File - read & execute
//   -rw-rw-rw-     0666             File - read & write
//   -rwxr-----     0740             File - Owner can read, write, & execute. Group can only read;
//                                   File - others have no permissions
//   drwxrwxrwx     20000000777      File - Directory - read, write, & execute for owner, group and others
//
func (fMgr *FileMgr) GetFilePermissionTextCodes() (string, error) {

  ePrefix := "FileMgr.GetFilePermissionTextCodes() "

  fMgrHlpr := fileMgrHelper{}
  filePathDoesExist := false
  var err, err2 error
  fPerm := 	FilePermissionConfig{}
  permissionText := ""

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "FileMgr")

  if err == nil && !filePathDoesExist {
    err =	fmt.Errorf(ePrefix +
      "Error: The current (FileMgr) path/file name DOES NOT EXIST!\n" +
      "Therefore no file permissions are available.\n" +
      "path/file name='%v'\n", fMgr.absolutePathFileName)
  }

  if err == nil {
    fPerm,
      err2 = FilePermissionConfig{}.NewByFileMode(fMgr.actualFileInfo.Mode())

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "%v", err2.Error())

    } else {

      permissionText, err2 = fPerm.GetPermissionTextCode()

      if err2 != nil {
        err =
          fmt.Errorf(ePrefix+
            "%v", err2.Error())
      }
    }
  }

  fMgr.dataMutex.Unlock()

  return permissionText, err
}

// GetFilePtr - will return the internal *os.File pointer
// for this File Manager instance. Depending on circumstances,
// this pointer may be nil.
//
func (fMgr *FileMgr) GetFilePtr() *os.File {

  var filePtr *os.File

  fMgr.dataMutex.Lock()

  filePtr = fMgr.filePtr

  fMgr.dataMutex.Unlock()

  return filePtr
}

// GetFileSize() - Returns os.FileInfo.Size() length in bytes for regular files;
// system-dependent for others.
//
// If the File Manager file does NOT exist, or if there is a file-path error, the
// value returned is -1.
//
func (fMgr *FileMgr) GetFileSize() int64 {

  fMgrHlpr := fileMgrHelper{}
  filePathDoesExist := false
  var err error
  var fileSize int64

  fMgr.dataMutex.Lock()

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    "",
    "FileMgr")

  if err == nil && !filePathDoesExist {
    err = fmt.Errorf("DoesNOTExist")
  }

  if err == nil {
    fileSize = fMgr.actualFileInfo.Size()
  }	else {
    fileSize = -1
  }

  fMgr.dataMutex.Unlock()

  return fileSize
}

// GetOriginalPathFileName - Returns the path and file name
// used originally to configure this File Manager object.
//
// Note: The original path and file name will be adjusted
// to reflect the path operators used in the operating system
// for the host computer.
//
func (fMgr *FileMgr) GetOriginalPathFileName() string {

  originalPathName := ""

  fMgr.dataMutex.Lock()

  originalPathName = fMgr.originalPathFileName

  fMgr.dataMutex.Unlock()

  return originalPathName
}

// GetReaderBufferSize() - Returns the size for the internal
// Bufio Reader's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Read' Operation.
//
func (fMgr *FileMgr) GetReaderBufferSize() int {

  readerBuffSize := 0

  fMgr.dataMutex.Lock()

  if fMgr.fileBufRdr != nil {
    fMgr.fileRdrBufSize = fMgr.fileBufRdr.Size()
    readerBuffSize = fMgr.fileRdrBufSize
  } else {
    readerBuffSize = fMgr.fileRdrBufSize
  }

  fMgr.dataMutex.Unlock()

  return readerBuffSize
}

// GetWriterBufferSize() - Returns the size for the internal
// Bufio Writers's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Write' Operation.
//
func (fMgr *FileMgr) GetWriterBufferSize() int {

  writerBuffSize := 0

  fMgr.dataMutex.Lock()

  if fMgr.fileBufWriter != nil {
    fMgr.fileWriterBufSize = fMgr.fileBufWriter.Size()
    writerBuffSize = fMgr.fileWriterBufSize
  } else {
    writerBuffSize = fMgr.fileWriterBufSize
  }

  fMgr.dataMutex.Unlock()

  return writerBuffSize
}

// IsAbsolutePathFileNamePopulated - Returns a boolean value
// indicating whether absolute path and file name is
// initialized and populated.
//
func (fMgr *FileMgr) IsAbsolutePathFileNamePopulated() bool {

  fMgrHlpr := fileMgrHelper{}
  var err error
  isAbsolutePathFileNamePopulated := false

  fMgr.dataMutex.Lock()
  _,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    "",
    "FileMgr")

  if err != nil || len(fMgr.absolutePathFileName) == 0 {
    fMgr.isAbsolutePathFileNamePopulated = false
    isAbsolutePathFileNamePopulated = false
  } else {
    fMgr.isAbsolutePathFileNamePopulated = true
    isAbsolutePathFileNamePopulated = true
  }

  fMgr.dataMutex.Unlock()

  return isAbsolutePathFileNamePopulated
}

// IsFileExtPopulated - Returns a boolean value indicating
// whether the File Extension for this File Manager instance
// is populated.
//
func (fMgr *FileMgr) IsFileExtPopulated() bool {

  fMgrHlpr := fileMgrHelper{}
  var err error
  isFileExtPopulated := false

  fMgr.dataMutex.Lock()

  _,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    "",
    "FileMgr.absolutePathFileName")

  if err != nil ||
    len(fMgr.fileExt) == 0 {
    fMgr.isFileExtPopulated = false
    isFileExtPopulated = false
  } else {
    fMgr.isFileExtPopulated = true
    isFileExtPopulated = true
  }

  fMgr.dataMutex.Unlock()

  return isFileExtPopulated
}

// IsFileMgrValid - Analyzes the current FileMgr object. If the
// current FileMgr object is INVALID, an error is returned.
//
// If the current FileMgr is VALID, this method returns 'nil'
//
func (fMgr *FileMgr) IsFileMgrValid(errorPrefixStr string) (err error) {

  fMgr.dataMutex.Lock()

  err = nil
  ePrefix := strings.TrimRight(errorPrefixStr, " ") + "FileMgr.IsFileMgrValid()"

  if !fMgr.isInitialized {
    err = errors.New(ePrefix + " Error: This data structure is NOT initialized.")
  } else if fMgr.absolutePathFileName == "" {
    fMgr.isAbsolutePathFileNamePopulated = false
    err = errors.New(ePrefix + " Error: absolutePathFileName is EMPTY!")
  } else {
    err2 := fMgr.dMgr.IsDirMgrValid(ePrefix)

    if err2 != nil {
      err = fmt.Errorf("FileMgr Directory Manager INVALID\n"+
        "Error='%v'\n",
        err2.Error())
    }

  }

  if err == nil {
    fMgrHelpr := fileMgrHelper{}
    _,
      err = fMgrHelpr.doesFileMgrPathFileExist(
      fMgr,
      PreProcPathCode.None(),
      ePrefix,
      "fMgr.absolutePathFileName")

    _ = fMgr.dMgr.DoesPathExist()
    _ = fMgr.dMgr.DoesAbsolutePathExist()
  }

  fMgr.dataMutex.Unlock()

  return err
}

// IsFileNameExtPopulated - Returns a boolean value indicating whether both the
// File Name and Extension for this File Manager instance have been populated.
//
// If either the File Name or the File Extension is blank (empty), this method
// returns false.
//
// Both the File Name AND the File Extension must be populated before this method
// returns 'true'.
//
func (fMgr *FileMgr) IsFileNameExtPopulated() bool {

  fMgrHlpr := fileMgrHelper{}
  var err error
  isFileNameExtPopulated := false

  fMgr.dataMutex.Lock()

  _,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    "",
    "FileMgr.absolutePathFileName")

  if err==nil &&
    len(fMgr.fileExt) > 0 &&
    len(fMgr.fileName) > 0 {

    fMgr.isFileNameExtPopulated = true
    isFileNameExtPopulated = true
  } else {

    fMgr.isFileNameExtPopulated = false
    isFileNameExtPopulated = false
  }

  fMgr.dataMutex.Unlock()

  return isFileNameExtPopulated
}

// IsFileNamePopulated - returns a boolean value
// indicating whether the file name for this File
// Manager object is populated.
//
func (fMgr *FileMgr) IsFileNamePopulated() bool {

  fMgrHlpr := fileMgrHelper{}
  var err error
  isFileNamePopulated := false

  fMgr.dataMutex.Lock()

  _,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    "",
    "FileMgr.absolutePathFileName")


  if err != nil || len(fMgr.fileName) == 0 {
    fMgr.isFileNamePopulated = false
    isFileNamePopulated = false
  } else {
    fMgr.isFileNamePopulated = true
    isFileNamePopulated = true
  }

  fMgr.dataMutex.Unlock()

  return isFileNamePopulated
}

// IsFilePointerOpen - Returns a boolean value indicating
// whether the File Pointer (*os.File) for this File Manager
// instance is open, or not.
//
func (fMgr *FileMgr) IsFilePointerOpen() bool {

  var isFilePtrOpen bool

  fMgr.dataMutex.Lock()

  if fMgr.filePtr == nil {
    fMgr.isFilePtrOpen = false
    isFilePtrOpen = false
  } else {
    fMgr.isFilePtrOpen = true
    isFilePtrOpen = true
  }

  fMgr.dataMutex.Unlock()

  return isFilePtrOpen
}

// isInitialized - Returns a boolean indicating whether the FileMgr
// object is properly initialized.
//
func (fMgr *FileMgr) IsInitialized() bool {

  isInitialized := false

  fMgr.dataMutex.Lock()

  isInitialized = fMgr.isInitialized

  fMgr.dataMutex.Unlock()

  return isInitialized
}

// MoveFileToFileMgr - This method will move the file identified
// by the current FileMgr to a new path and file name specified
// by the input parameter 'destinationFMgr' which is of type
// 'FileMgr'.
//
// IMPORTANT:
//
// The file identified by the current FileMgr object will be
// DELETED!
//
// The new file is defined by input parameter 'destinationFMgr'.
//
// If the input parameter 'destinationFMgr' contains a directory path
// which currently does NOT exist, it will be created.
//
func (fMgr *FileMgr) MoveFileToFileMgr(destinationFMgr FileMgr) error {

  ePrefix := "FileMgr.MoveFileToFileMgr() "
  var err error

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.moveFile(fMgr, &destinationFMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// MoveFileToFileStr - This method will move the file identified
// by the current FileMgr to a new path and file name contained
// in the input parameter 'pathFileNameExt' which is of type
// string.
//
// IMPORTANT:
//
// The file identified by the current FileMgr object will be
// DELETED!
//
// The new file is defined by input parameter 'pathFileNameExt'.
//
// If the input parameter 'pathFileNameExt' contains a directory path
// which currently does NOT exist, it will be created.
//
func (fMgr *FileMgr) MoveFileToFileStr(pathFileNameExt string) (FileMgr, error) {

  ePrefix := "FileMgr.MoveFileToFileStr() "
  var err error

  targetFMgr, err := FileMgr{}.New(pathFileNameExt)

  if err != nil {
    return FileMgr{},
      fmt.Errorf("Error returned by FileMgr{}.New(pathFileNameExt)\n"+
        "pathFileName='%v'\nError='%v'\n", pathFileNameExt, err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.moveFile(fMgr, &targetFMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return targetFMgr, err
}

// MoveFileToNewDir - This method will move the current file
// identified by this FileMgr object to a new path designated
// by input parameter string, 'dirPath'.
//
// IMPORTANT:
//
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
// If input parameter 'dirPth' contains a directory path which does not
// currently exist, it will be created.
//
func (fMgr *FileMgr) MoveFileToNewDir(dirPath string) (FileMgr, error) {

  ePrefix := "FileMgr.MoveFileToNewDir() "

  dirMgr, err := DirMgr{}.New(dirPath)

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+"Error returned by DirMgr{}.New(dirPath).\n"+
        "dirPath='%v'\nError='%v'\n", dirPath, err.Error())
  }

  targetFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, fMgr.fileNameExt)

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+"Error returned by FileMgr{}.NewFromDirMgrFileNameExt"+
        "(dirMgr, fMgr.fileNameExt)\n"+
        "dirMgr='%v'\n"+
        "fMgr.fileNameExt='%v'\nError='%v'\n",
        dirMgr.absolutePath, fMgr.fileNameExt, err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.moveFile(fMgr, &targetFMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return targetFMgr, err
}

// MoveFileToNewDirMgr - This method will move the file identified
// by the current FileMgr to a new path contained in the input parameter
// 'dMgr'.
//
// IMPORTANT:
//
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned as a type
// 'FileMgr'.
//
// If the input parameter 'dMgr' does not contain a path which currently
// exists, it will be created.
//
func (fMgr *FileMgr) MoveFileToNewDirMgr(dMgr DirMgr) (FileMgr, error) {

  ePrefix := "FileMgr.MoveFileToNewDirMgr() "

  var err error
  var targetFMgr FileMgr

  err2 := dMgr.IsDirMgrValid("")

  if err2 != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+"Error: Input parameter 'dMgr' is INVALID! "+
      "Error='%v'", err2.Error())
  }

  fMgr.dataMutex.Lock()

  targetFile := dMgr.GetAbsolutePathWithSeparator() + fMgr.fileNameExt

  targetFMgr, err2 = FileMgr{}.New(targetFile)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.New(targetFile)\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err2.Error())

    fMgr.dataMutex.Unlock()
    return FileMgr{}, err
  }

  fMgrHlpr := fileMgrHelper{}

  err = fMgrHlpr.moveFile(fMgr, &targetFMgr, ePrefix)

  fMgr.dataMutex.Unlock()

  return targetFMgr, err
}

// New - Creates a new File Manager ('FileMgr') instance. This method receives an
// input parameter of type string and parses out the path, file name and file
// extension.
//
// The file data is returned in the data fields of the new FileMgr object.
//
// This method is identical to method 'NewFromPathFileNameExtStr()'.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//	pathFileNameExt string - Must consist of a valid path, file name
//	                         and file extension. The file need not exist.
//	                         Failure to provide a properly formatted path
//	                         path, file name will result in an error.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgr       - If the method completes successfully, a valid FileMgr instance
//	                is returned containing information on the file specified in the
//	                two input parameter 'pathFileNameExt'.
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	fmgr := FileMgr{}.New("../common/fileName.ext")
//
func (fMgr FileMgr) New(pathFileNameExt string) (FileMgr, error) {

  ePrefix := "FileMgr.New() "

  fMgr2 := FileMgr{}
  fMgrHlpr := fileMgrHelper{}

  isEmpty, err := fMgrHlpr.setFileMgrPathFileName(&fMgr2, pathFileNameExt, ePrefix)

  if isEmpty {

    fMgr2 = FileMgr{}

    if err == nil {
      return FileMgr{}, fmt.Errorf(ePrefix+
        "Error: The FileMgr instance generated by input parameter 'pathFileNameExt' is Empty!\n"+
        "pathFileNameExt='%v'\n",
        pathFileNameExt)
    }
  }

  return fMgr2, err
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
// ----------------------------------------------------------------------------------
//
// Input Parameters:
//
//  dirMgr      DirMgr - A valid and properly initialized Directory Manager (type: DirMgr).
//                       This object contains the file path which will be combined with the
//                       the second input parameter, 'fileNameExt' to create the new File
//                       Manager (type: FileMgr)
//
//  fileNameExt string - A string containing the file name and file extension which will be
//                       used to create the new File Manager (type: FileMgr).
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgr       - If the method completes successfully, a valid FileMgr instance
//	                is returned containing information on the file specified in the
//	                two input parameters 'dirMgr' and 'fileNameExt'.
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  dMgr, err := DirMgr{}.New("D:\\TestDir")
//
//  if err != nil {
//   fmt.Printf("Error='%v' \n", err.Error())
//   return
//  }
//
//  fileNameExt := "yourFile.txt"
//
//  fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExt)
//
//
func (fMgr FileMgr) NewFromDirMgrFileNameExt(
  dirMgr DirMgr,
  fileNameExt string) (FileMgr, error) {

  ePrefix := "FileMgr.NewFromDirMgrFileNameExt() "

  fMgr2 := FileMgr{}

  fMgrHlpr := fileMgrHelper{}

  isEmpty, err := fMgrHlpr.setFileMgrDirMgrFileName(&fMgr2, dirMgr, fileNameExt, ePrefix)

  if isEmpty {

    fMgr2 = FileMgr{}

    if err == nil {
      return FileMgr{}, fmt.Errorf(ePrefix+
        "Error: The FileMgr instance generated by input parameters 'dirMgr' and "+
        "'fileNameExt' is Empty!\n"+
        "dirMgr='%v'\nfileNameExt='%v'\n",
        dirMgr.absolutePath, fileNameExt)
    }
  }

  return fMgr2, err
}

// NewFromDirStrFileNameStr - Creates a new file manager object (FileMgr) from a directory
// string and a File Name and Extension string passed as input parameters.
//
// ----------------------------------------------------------------------------------
//
// Input Parameters:
//
//  dirStr         string - A string containing the directory or file path which will be
//                          combined with the second input parameter, 'fileNameExt' to
//                          create the new File Manager (type: FileMgr).
//
//  fileNameExtStr string - A string containing the file name and file extension which will be
//                          used to create the new File Manager (type: FileMgr).
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgr       - If the method completes successfully, a valid FileMgr instance
//	                is returned containing information on the file specified in the
//	                two input parameters 'dirStr' and 'fileNameExtStr'.
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  dirStr := "D:\\TestDir"
//
//  fileNameExtStr := "yourFile.txt"
//
//  fMgr, err := FileMgr{}.NewFromDirStrFileNameStr(dirStr, fileNameExtStr)
//
func (fMgr FileMgr) NewFromDirStrFileNameStr(
  dirStr,
  fileNameExtStr string) (FileMgr, error) {

  ePrefix := "FileMgr.NewFromDirStrFileNameStr() "

  dirMgr, err := DirMgr{}.New(dirStr)

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirMgr).\n"+
        "dirStr='%v'\nError='%v'\n",
        dirStr, err.Error())
  }

  fMgr2 := FileMgr{}

  fMgrHlpr := fileMgrHelper{}

  isEmpty, err := fMgrHlpr.setFileMgrDirMgrFileName(&fMgr2, dirMgr, fileNameExtStr, ePrefix)

  if isEmpty {

    fMgr2 = FileMgr{}

    if err == nil {
      return FileMgr{}, fmt.Errorf(ePrefix+
        "Error: The FileMgr instance generated by input parameters 'dirStr' and "+
        "'fileNameExtStr' is Empty!\n"+
        "dirStr='%v'\nfileNameExtStr='%v'\n",
        dirStr, fileNameExtStr)
    }
  }

  return fMgr2, err
}

// NewFromFileInfo - Creates and returns a new FileMgr object based on input from a
// directory path string and an os.FileInfo object.
//
// ----------------------------------------------------------------------------------
//
// Input Parameters:
//
//  dirPathStr      string - The directory path. NOTE: This does NOT contain the
//	                         file name.
//
//	info       os.FileInfo - A valid and populated FileInfo structure containing the
//	                         file name.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgr       - If the method completes successfully, a valid FileMgr instance
//	                is returned containing information on the file specified in the
//	                input parameter 'info'.
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
func (fMgr FileMgr) NewFromFileInfo(dirPathStr string, info os.FileInfo) (FileMgr, error) {

  ePrefix := "FileMgr.NewFromFileInfo() "

  if info == nil {
    return FileMgr{},
      errors.New(ePrefix + "Error: Input parameter 'info' is 'nil' and INVALID!")
  }

  fileName := info.Name()

  dirMgr, err := DirMgr{}.New(dirPathStr)

  if err != nil {
    return FileMgr{},
      fmt.Errorf("Input Parameter 'dirPathStr' is INVALID! DirMgr{}.New(dirPathStr) "+
        "Error='%v' ", err.Error())
  }

  fMgr2 := FileMgr{}

  fMgrHlpr := fileMgrHelper{}

  isEmpty, err := fMgrHlpr.setFileMgrDirMgrFileName(&fMgr2, dirMgr, fileName, ePrefix)

  if isEmpty {

    fMgr2 = FileMgr{}

    if err == nil {
      return FileMgr{}, fmt.Errorf(ePrefix+
        "Error: The FileMgr instance generated by input parameters 'dirPathStr' and "+
        "'info' is Empty!\n"+
        "dirPathStr='%v'\ninfo='%v'\n",
        dirPathStr, fileName)
    }
  }

  return fMgr2, err
}

// NewFromPathFileNameExtStr - Creates a new File Manager ('FileMgr') instance. This
// method receives an input parameter of type string and parses out the path, file
// name and file extension.
//
// The file data is returned in the data fields of the new FileMgr object.
//
// This method is identical to method 'New()'.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileNameExt string - Must consist of a valid path, file name
//                           and file extension. The file need not exist.
//                           Failure to provide a properly formatted path
//                           path and file name will result in an error.
//                           The file extension is optional.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	fmgr := FileMgr{}.NewFromPathFileNameExtStr("../somedirectory/fileName.ext")
//
func (fMgr FileMgr) NewFromPathFileNameExtStr(pathFileNameExt string) (FileMgr, error) {

  ePrefix := "FileMgr.NewFromPathFileNameExtStr() "

  fMgr2 := FileMgr{}

  fMgrHlpr := fileMgrHelper{}

  isEmpty, err := fMgrHlpr.setFileMgrPathFileName(&fMgr2, pathFileNameExt, ePrefix)

  if isEmpty {

    fMgr2 = FileMgr{}

    if err == nil {
      return FileMgr{}, fmt.Errorf(ePrefix+
        "Error: The FileMgr instance generated by input parameter 'pathFileNameExt' is Empty!\n"+
        "pathFileNameExt='%v'\n",
        pathFileNameExt)
    }
  }

  return fMgr2, err
}

// OpenThisFile - Opens the file identified by the current FileMgr object
// using the file open parameters and file permission parameters contained
// in the 'FileAccessControl' instance passed as 'fileAccessCtrl'.
//
// Note: If the FileMgr directory path does not exist, this method will
// create that directory path.
//
func (fMgr *FileMgr) OpenThisFile(fileAccessCtrl FileAccessControl) error {
  ePrefix := "FileMgr.OpenThisFile() "
  var err error

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    fileAccessCtrl,
    true,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// OpenThisFileReadOnly - Opens the file identified by the current
// FileMgr object as a 'Read-Only' File. Subsequent operations may
// read from this file but may NOT write to this file.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for reading only.
//
// If FileMgr.absolutePathFileName does not exist, an error will be
// returned.
//
// If successful, the FileMode is set to "-r--r--r--" and the permission
// Mode is set to '0444'.
//
// Note: If the 'FileMgr' directory path or file do not exist, this
// method will return an error.
//
func (fMgr *FileMgr) OpenThisFileReadOnly() error {

  ePrefix := "FileMgr.OpenThisFileReadOnly() "

  readOnlyAccessCtrl, err := FileAccessControl{}.NewReadOnlyAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    readOnlyAccessCtrl,
    false,
    false,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// OpenThisFileWriteOnly - Opens the current file for 'WriteOnly'
// operations.  If successful, this method will use FileMgr.absolutePathFileName
// to open an *os.File or File Pointer.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for writing only. If FileMgr.absolutePathFileName
// does not exist, it will be created. The FileMode is set to "--w--w--w-" and
// the permission Mode is set to '0222'.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will will create them.
//
func (fMgr *FileMgr) OpenThisFileWriteOnly() error {
  var err error
  ePrefix := "FileMgr.OpenThisFileWriteOnly() "
  fMgrHlpr := fileMgrHelper{}

  writeOnlyAccessCtrl, err := FileAccessControl{}.NewWriteOnlyAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    writeOnlyAccessCtrl,
    true,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()
  return err
}

// OpenThisFileWriteOnlyAppend - Opens the current file for 'Write Only'
// with an 'Append' mode. All bytes written to the file will be written
// at the end of the current file and none of the file's original content
// will be overwritten.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will will create both the path and file.
//
func (fMgr *FileMgr) OpenThisFileWriteOnlyAppend() error {

  ePrefix := "FileMgr.OpenThisFileWriteOnlyAppend() "

  writeOnlyFileAccessCfg, err := FileAccessControl{}.NewWriteOnlyAppendAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    writeOnlyFileAccessCfg,
    true,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// OpenThisFileWriteOnlyAppend - Opens the current file for 'Write Only'
// with an 'Truncate' mode. This means that if the file previously exists,
// all of the existing file content will be deleted before the 'write'
// operation will begin.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will will create both the path and file.
//
func (fMgr *FileMgr) OpenThisFileWriteOnlyTruncate() error {

  ePrefix := "FileMgr.OpenThisFileWriteOnlyTruncate() "

  writeOnlyTruncateAccessCfg, err := FileAccessControl{}.NewWriteOnlyTruncateAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    writeOnlyTruncateAccessCfg,
    true,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// OpenThisFileReadWrite - Opens the file identified by the current
// FileMgr object. If successful, this method will use
// FileMgr.absolutePathFileName to open an *os.File or File Pointer.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for reading and writing. If FileMgr.absolutePathFileName
// does not exist, it will be created. The FileMode is set to'-rw-rw-rw-' and
// the permission Mode= '0666'.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will will create them.
//
func (fMgr *FileMgr) OpenThisFileReadWrite() error {
  var err error

  ePrefix := "FileMgr.OpenThisFileReadWrite() "

  readWriteAccessCtrl, err := FileAccessControl{}.NewReadWriteAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.openFile(
    fMgr,
    readWriteAccessCtrl,
    true,
    true,
    ePrefix)

  fMgr.dataMutex.Unlock()

  return err
}

// ReadAllFile - Reads the file identified by the current FileMgr
// and returns the contents in a byte array.
//
// If no errors are encountered the returned 'error' value is
// nil. This method does not return io.EOF as an error. This is
// is because it reads the entire file and therefor no End Of File
// flag is required.
//
func (fMgr *FileMgr) ReadAllFile() (bytesRead []byte, err error) {

  ePrefix := "FileMgr.ReadAllFile() "
  bytesRead = []byte{}
  err = nil
  var err2 error

  readWriteAccessCtrl, err2 := FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return bytesRead, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.readFileSetup(
    fMgr,
    readWriteAccessCtrl,
    false,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return bytesRead, err
  }

  bytesRead, err2 = ioutil.ReadAll(fMgr.filePtr)

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned by ioutil.ReadAll(fMgr.filePtr).\n"+
        "fileName='%v'\nErrors='%v'\n",
        fMgr.absolutePathFileName, err2.Error())
  }

  fMgr.dataMutex.Unlock()
  return bytesRead, err
}

// ReadFileBytes - Reads bytes from the file identified by the current FileMgr
// object. Bytes are stored in 'byteBuff', a byte array passed in as an input
// parameter.
//
// If successful, the returned error value is 'nil'. The returned value 'int'
// contains the number of bytes read from the current file.
//
// At End of File (EOF), the byte count will be zero and err will be equal to
// 'io.EOF'.
//
func (fMgr *FileMgr) ReadFileBytes(byteBuff []byte) (bytesRead int, err error) {

  ePrefix := "FileMgr.ReadFileBytes() "
  bytesRead = 0
  err = nil
  var err2 error

  readWriteAccessCtrl, err2 := FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
      "Error='%v'\n", err2.Error())
    return bytesRead, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.readFileSetup(
    fMgr,
    readWriteAccessCtrl,
    false,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return bytesRead, err
  }

  bytesRead, err2 = fMgr.fileBufRdr.Read(byteBuff)

  if err2 != nil &&
    err2 == io.EOF {

    err = err2

  } else if err2 != nil {

    err = fmt.Errorf("Error returned by fMgr.fileBufRdr.Read(byteBuff). "+
      "File='%v' Error='%v' ", fMgr.absolutePathFileName, err2.Error())
  }

  fMgr.dataMutex.Unlock()
  return bytesRead, err
}

// ReadFileLine - Effectively, this method reads a file one line
// at a time and returns the line as an array of bytes. The delimiter
// for lines read is specified by input parameter 'delim' of type byte.
//
// This method uses the 'bufio' package.
//
// If End Of File (EOF) is reached, this method returns the 'bytesRead' and
// an error which is equal to 'io.EOF'.
//
// Note that the delimiter is returned as the last character in bytes read.
//
func (fMgr *FileMgr) ReadFileLine(delim byte) (bytesRead []byte, err error) {

  ePrefix := "FileMgr.ReadFileLine() "
  bytesRead = []byte{}
  err = nil
  var err2 error

  readWriteAccessCtrl, err2 := FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
      "Error='%v'\n", err2.Error())
    return bytesRead, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.readFileSetup(
    fMgr,
    readWriteAccessCtrl,
    false,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return bytesRead, err
  }

  bytesRead, err2 = fMgr.fileBufRdr.ReadBytes(delim)

  if err2 != nil &&
    err2 == io.EOF {
    err = err2

  } else if err2 != nil {

    err = fmt.Errorf(ePrefix+
      "\nError returned from fMgr.fileBufRdr.ReadBytes(delim).\n"+
      "Error='%v'\n", err2.Error())
  }

  fMgr.dataMutex.Unlock()

  return bytesRead, err
}

// ReadFileString - Wrapper for bufio.ReadString. ReadFileString reads until the
// first occurrence of 'delim' in the input, returning a string containing the
// data up to and including the delimiter. If ReadFileString encounters an error
// before finding a delimiter, it returns the data read before the error and
// the error itself (often io.EOF). ReadFileString returns err != nil if and
// only if the returned data does not end in 'delim'.
//
func (fMgr *FileMgr) ReadFileString(delim byte) (stringRead string, err error) {

  ePrefix := "FileMgr.ReadFileString() "
  stringRead = ""
  err = nil
  var err2 error

  readWriteAccessCtrl, err2 := FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
      "Error='%v'\n", err2.Error())
    return stringRead, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.readFileSetup(
    fMgr,
    readWriteAccessCtrl,
    false,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return stringRead, err
  }

  stringRead, err2 = fMgr.fileBufRdr.ReadString(delim)

  if err2 != nil &&
    err2 == io.EOF {
    err = err2

  } else if err2 != nil {

    err = fmt.Errorf("Error returned from fMgr.fileBufRdr.ReadString(delim). "+
      "Error='%v' ", err2.Error())
  }

  fMgr.dataMutex.Unlock()
  return stringRead, err
}

// ResetFileInfo - Acquires the current os.FileInfo
// data associated with the file identified by the
// current FileMgr instance.
//
// An error will be triggered if the file path does NOT exist!
//
func (fMgr *FileMgr) ResetFileInfo() error {

  ePrefix := "ResetFileInfo() "

  var err error
  var filePathDoesExist bool

  fMgr.dataMutex.Lock()

  fMgrHelpr := fileMgrHelper{}
  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "Current FileMgr file DOES NOT EXIST!\n"+
      "FMgr='%v'", fMgr.absolutePathFileName)
  }

  return nil
}

// SetReaderBufferSize - Sets the Read Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
//
func (fMgr *FileMgr) SetReaderBufferSize(readBuffSize int) {
  fMgr.dataMutex.Lock()
  fMgr.fileRdrBufSize = readBuffSize
  fMgr.dataMutex.Unlock()
}

// SetWriterBufferSize - Sets the Write Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
//
func (fMgr *FileMgr) SetWriterBufferSize(writeBuffSize int) {
  fMgr.dataMutex.Lock()
  fMgr.fileWriterBufSize = writeBuffSize
  fMgr.dataMutex.Unlock()
}

// SetFileInfo - Used to initialize the os.FileInfo structure maintained as
// part of the current FileMgr object.
//
// Be careful: Failure of the input parameter to match the current FileMgr file name
// will trigger an error.
//
func (fMgr *FileMgr) SetFileInfo(info os.FileInfo) error {

  ePrefix := "FileMgr.SetFileInfo() "

  if info == nil {
    return errors.New(ePrefix +
      "Error: Input parameter 'info' is 'nil' and INVALID!\n")
  }

  if info.Name() == "" {
    return errors.New(ePrefix +
      "Error: info.Name() is an EMPTY string!\n")
  }

  if info.IsDir() {
    return errors.New(ePrefix +
      "Input parameter info.IsDir()=='true'. This is a Directory NOT A FILE!\n")
  }

  if strings.ToLower(info.Name()) != strings.ToLower(fMgr.fileNameExt) {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'info' does NOT match current FileMgr file name.\n"+
      "FileMgr File Name='%v'\ninfo File Name='%v'\n",
      fMgr.fileNameExt, info.Name())
  }

  fMgr.dataMutex.Lock()

  fMgr.actualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)

  fMgr.dataMutex.Unlock()

  if !fMgr.actualFileInfo.isFInfoInitialized {
    return fmt.Errorf(ePrefix+
      "Error: Failed to initialize fMgr.actualFileInfo object.\n"+
      "info.Name()='%v'", info.Name())
  }

  return nil
}

// SetFileMgrFromDirMgrFileName - Sets the data fields of the current FileMgr object
// based on a DirMgr object and a File Name string which are passed as input parameters.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	isEmpty       - This value is set to 'false' if, and only if, all internal
//	                values are set to valid legitimate values.
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//
func (fMgr *FileMgr) SetFileMgrFromDirMgrFileName(
  dMgr DirMgr,
  fileNameExt string) (isEmpty bool, err error) {

  ePrefix := "FileMgr.SetFileMgrFromDirMgrFileName() "

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()
  isEmpty, err = fMgrHlpr.setFileMgrDirMgrFileName(fMgr, dMgr, fileNameExt, ePrefix)
  fMgr.dataMutex.Unlock()

  return isEmpty, err
}

// SetFileMgrFromPathFileName - Initializes all the data fields of the
// current FileMgr object based on the path file name string passed to
// this method as an input parameter.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	isEmpty       - This value is set to 'false' if, and only if, all internal
//                  values are set of valid legitimate values.
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//

func (fMgr *FileMgr) SetFileMgrFromPathFileName(
  pathFileNameExt string) (isEmpty bool, err error) {

  ePrefix := "FileMgr.SetFileMgrFromPathFileName() "
  isEmpty = true
  err = nil

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  isEmpty, err = fMgrHlpr.setFileMgrPathFileName(fMgr, pathFileNameExt, ePrefix)

  fMgr.dataMutex.Unlock()

  return isEmpty, err
}

// WriteBytesToFile a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
//
// If the number of bytes written to the file is less than len(bytes),
// an error will be returned.
//
// This method uses the 'bufio' package.
//
func (fMgr *FileMgr) WriteBytesToFile(
  bytes []byte) (numBytesWritten int, err error) {

  ePrefix := "FileMgr.WriteBytesToFile() "
  err = nil
  numBytesWritten = 0

  readWriteAccessCtrl, err2 :=
    FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileAccessControl{}."+
      "NewReadWriteAccess()\n."+
      "Error='%v'\n", err2.Error())

    return numBytesWritten, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.writeFileSetup(
    fMgr,
    readWriteAccessCtrl,
    true,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return numBytesWritten, err
  }

  numBytesWritten, err2 =
    fMgr.fileBufWriter.Write(bytes)

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned from fMgr.fileBufWriter.Write(bytes). Output File='%v'. "+
        "Error='%v'",
        fMgr.absolutePathFileName, err2.Error())

  } else {

    fMgr.buffBytesWritten += uint64(numBytesWritten)

  }

  fMgr.dataMutex.Unlock()

  return numBytesWritten, err
}

// WriteStrToFile - Writes a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
//
// If the number of bytes written to the file is less than len(str),
// an error will be returned.
//
// This method uses the 'bufio' package.
//
func (fMgr *FileMgr) WriteStrToFile(str string) (numBytesWritten int, err error) {

  ePrefix := "FileMgr.WriteStrToFile() "

  numBytesWritten = 0
  err = nil
  var err2 error
  var readWriteAccessCtrl FileAccessControl

  readWriteAccessCtrl, err2 =
    FileAccessControl{}.NewReadWriteAccess()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileAccessControl{}."+
      "NewReadWriteAccess()\n."+
      "Error='%v'\n", err2.Error())

    return numBytesWritten, err
  }

  fMgrHlpr := fileMgrHelper{}

  fMgr.dataMutex.Lock()

  err = fMgrHlpr.writeFileSetup(
    fMgr,
    readWriteAccessCtrl,
    true,
    ePrefix)

  if err != nil {
    fMgr.dataMutex.Unlock()
    return numBytesWritten, err
  }

  numBytesWritten, err2 = fMgr.fileBufWriter.WriteString(str)

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned from fMgr.filePtr.WriteString(str). Output File='%v'. "+
        "Error='%v'", fMgr.absolutePathFileName, err2.Error())

  } else {
    fMgr.buffBytesWritten += uint64(numBytesWritten)
  }

  fMgr.dataMutex.Unlock()
  return numBytesWritten, err
}
