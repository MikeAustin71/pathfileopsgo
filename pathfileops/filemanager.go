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

  filePathDoesExist,
    err := fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: This file does NOT exist!\n"+
      "File Name:'%v' ", fMgr.absolutePathFileName)
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

  err = os.Chmod(fMgr.absolutePathFileName, fileMode)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by os.Chmod(fMgr.absolutePathFileName, fileMode).\n"+
      "fileMode='%v'\nError='%v'\n",
      mode.GetPermissionFileModeValueText(), err.Error())
  }

  _,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "Verify fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  return nil
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

  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    fMgr.filePtr = nil
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()
    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0
    return err
  }

  if fMgr.filePtr == nil {
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()
    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0
    return nil
  }

  fileOpenType, err := fMgr.fileAccessStatus.GetFileOpenType()

  if err == nil {

    if fileOpenType == FOpenType.TypeWriteOnly() ||
      fileOpenType == FOpenType.TypeReadWrite() {

      err = fMgr.FlushBytesToDisk()

      if err != nil {

        _ = fMgr.filePtr.Close()

        fMgr.isFilePtrOpen = false
        fMgr.fileAccessStatus.Empty()
        fMgr.fileBufRdr = nil
        fMgr.fileBufWriter = nil
        fMgr.fileBytesWritten = 0
        fMgr.buffBytesWritten = 0

        return fmt.Errorf(ePrefix+
          "Error returned from fMgr.FlushBytesToDisk().  "+
          "Error='%v'", err.Error())
      }
    }
  }

  fMgr.dataMutex.Lock()

  err = fMgr.filePtr.Close()

  fMgr.dataMutex.Unlock()

  fMgr.isFilePtrOpen = false
  fMgr.filePtr = nil
  fMgr.fileAccessStatus.Empty()
  fMgr.fileBufRdr = nil
  fMgr.fileBufWriter = nil
  fMgr.fileBytesWritten = 0
  fMgr.buffBytesWritten = 0

  if err != nil {

    return fmt.Errorf(ePrefix+
      "Received Error from fMgr.filePtr.Close().\n"+
      "fMgr.absolutePathFileName= '%v'\n",
      fMgr.absolutePathFileName)
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

  fMgrHlpr := fileMgrHelper{}

  filePathDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: This file does NOT exist!\n"+
      "File Name:'%v' ", fMgr.absolutePathFileName)
  }

  if fMgrDest == nil {
    return errors.New(ePrefix +
      "Error: Input parameter fMgrDest is a nil pointer!\n")
  }

  err = fMgrDest.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: The Destination FileMgr object is INVALID!\n"+
      "Error='%v'\n",
      err.Error())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(fMgrDest,
    PreProcPathCode.None(),
    ePrefix,
    "fMgrDest.absolutePathFileName")

  if err != nil {
    return err
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied.\nFile='%v'\n",
      fMgr.absolutePathFileName)
  }

  err = fMgrDest.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Atempted creation of destination directory FAILED! Error= '%v'",
      err.Error())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(fMgrDest,
    PreProcPathCode.None(),
    ePrefix,
    "fMgrDest.absolutePathFileName")

  if err != nil {
    return err
  }

  if filePathDoesExist && !fMgrDest.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Destination file exists and it is NOT a 'regular' file.\n"+
      "Copy operation aborted!\nDestination File='%v'\n",
      fMgrDest.absolutePathFileName)
  }

  // See Reference:
  // https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

  fMgr.dataMutex.Lock()

  err = FileHelper{}.CopyFileByIo(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

  fMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.CopyFileByIo(fMgr.absolutePathFileName, "+
      "fMgrDest.absolutePathFileName) fMgr.absolutePathFileName='%v' "+
      "fMgrDest.absolutePathFileName='%v' Error='%v'",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName, err.Error())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgrDest,
    PreProcPathCode.None(),
    ePrefix,
    "fMgrDest.absolutePathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: After attempted file copy to destination file.\n"+
      "Destination file does NOT exist!\n"+
      "fMgrDest='%v'\n",
      fMgrDest.absolutePathFileName)
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

  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return err
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

  fileDoesExist, nonPathError := fMgr.DoesThisFileExist()

  if nonPathError != nil {
    return fmt.Errorf(ePrefix+"%v", nonPathError)
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+"Error: Source file DOES NOT EXIST!\n"+
      "Source File (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    return fmt.Errorf(ePrefix+"Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied.\n"+
      "File='%v'\n", fMgr.absolutePathFileName)
  }

  err = fMgrDest.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Atempted creation of destination directory FAILED!\n"+
      "Error='%v'\n",
      err.Error())
  }

  fileDoesExist, nonPathError = fMgrDest.DoesThisFileExist()

  if nonPathError != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fMgrDest.DoesThisFileExist().\nError= '%v'\n",
      nonPathError.Error())
  }

  if fileDoesExist && !fMgrDest.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Destination file exists and it is NOT a 'regular' file.\n"+
      "Copy operation aborted! Destination File='%v'\n",
      fMgrDest.absolutePathFileName)
  }

  // See Reference:
  // https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

  fMgr.dataMutex.Lock()

  fh := FileHelper{}

  err = fh.CopyFileByIoByLink(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

  fMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.CopyFileByIoByLink(fMgr.absolutePathFileName, "+
      "fMgrDest.absolutePathFileName)\nMgr.absolutePathFileName='%v'\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName, err.Error())
  }

  fileDoesExist, nonPathError = fMgrDest.DoesThisFileExist()

  if nonPathError != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgrDest.DoesThisFileExist().\n"+
      "fMgrDest='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, nonPathError.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: After attempted file copy to destination file.\n"+
      "Destination file does NOT exist!\n"+
      "fMgrDest='%v'\n",
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
      "Error: This FileMgr object is INVALID!\nError='%v'\n",
      err.Error())
  }

  if fMgrDest == nil {
    return errors.New(ePrefix +
      "Error: Input parameter fMgrDest is a nil pointer!\n")
  }

  err = fMgrDest.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: The Destination FileMgr object is INVALID!\nError='%v'\n",
      err.Error())
  }

  if !fMgr.doesAbsolutePathFileNameExist {
    return fmt.Errorf(ePrefix+"Error: Source file DOES NOT EXIST!\n"+
      "Source File='%v'\n",
      fMgr.absolutePathFileName)
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied.\n"+
      "File='%v'\n",
      fMgr.absolutePathFileName)
  }

  err = fMgrDest.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Atempted creation of destination directory FAILED!\n"+
      "Error= '%v'\n",
      err.Error())
  }

  destFileExists, err := fMgrDest.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fMgrDest.DoesThisFileExist().\n"+
      "Error= '%v'\n",
      err.Error())
  }

  if destFileExists && !fMgrDest.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Destination file exists and it is NOT a 'regular' file.\n"+
      "Copy operation aborted!\nDestination File='%v'\n",
      fMgrDest.absolutePathFileName)
  }

  err = fMgrDest.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  // See Reference:
  // https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

  fMgr.dataMutex.Lock()

  err = FileHelper{}.CopyFileByLink(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

  fMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.CopyFileByLink(fMgr.absolutePathFileName, "+
      "fMgrDest.absolutePathFileName)\nfMgr.absolutePathFileName='%v'\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName, err.Error())
  }

  destFileExists, err = fMgrDest.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgrDest.DoesThisFileExist().\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, err.Error())
  }

  if !destFileExists {
    return fmt.Errorf(ePrefix+
      "Error: After attempted file copy to destination file. Destination "+
      "file does NOT exist!\n"+
      "fMgrDest.absolutePathFileName='%v'\n",
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
      "Error: This FileMgr object is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

  if fMgrDest == nil {
    return errors.New(ePrefix +
      "Error: Input parameter fMgrDest is a nil pointer!\n")
  }

  err = fMgrDest.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: The Destination FileMgr object (input parameter 'fMgrDest') "+
      "is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: Source File DOES NOT EXIST!\n"+
      "Source File='%v'\n",
      fMgr.absolutePathFileName)
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    return fmt.Errorf(ePrefix+"Error: Source and Destination File are the same!\n"+
      "Source File='%v'\n"+
      "Destination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied. File='%v'", fMgr.absolutePathFileName)
  }

  err = fMgrDest.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Atempted creation of destination directory FAILED!\n"+
      "Error= '%v'\n",
      err.Error())
  }

  destFileExists, err := fMgrDest.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fMgrDest.DoesThisFileExist().\n"+
      "Error= '%v'\n", err.Error())
  }

  if destFileExists && !fMgrDest.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: Destination file exists and it is NOT a 'regular' file. "+
      "Copy operation aborted!\n"+
      "Destination File='%v' ", fMgrDest.absolutePathFileName)
  }

  // See Reference:
  // https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

  fMgr.dataMutex.Lock()

  err = FileHelper{}.CopyFileByLinkByIo(fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

  fMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.CopyFileByLinkByIo(fMgr.absolutePathFileName, "+
      "fMgrDest.absolutePathFileName)\nfMgr.absolutePathFileName='%v'\n"+
      "fMgrDest.absolutePathFileName='%v'\n"+
      "Error='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName, err.Error())
  }

  destFileExists, err = fMgrDest.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgrDest.DoesThisFileExist().\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, err.Error())
  }

  if !destFileExists {
    return fmt.Errorf(ePrefix+
      "Error: After attempted file copy to destination file. Destination "+
      "file does NOT exist!\nfMgrDest.absolutePathFileName='%v'\n",
      fMgrDest.absolutePathFileName)
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+"This File Manager instance is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+"This File Manager file DOES NOT EXIST!\n"+
      "File Name='%v'\n", fMgr.absolutePathFileName)
  }

  fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt).\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&fMgrDest) {
    return fmt.Errorf(ePrefix+"Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByIo(&fMgrDest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByIo(&fMgrDest) "+
      "fMgrDest.absolutePathFileName='%v'  Error='%v'", fMgrDest.absolutePathFileName, err.Error())
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "This File Manager instance is INVALID!\n"+
      "Error='%v' ", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+"Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v' ", fMgr.absolutePathFileName)
  }

  fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt).\n"+
      "dstPathFileNameExt='%v'\nError='%v'",
      dstPathFileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&fMgrDest) {
    return fmt.Errorf(ePrefix+"Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByIoByLink(&fMgrDest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByIoByLink(&fMgrDest)\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, err.Error())
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "This File Manager instance is INVALID!\nError='%v'\n",
      err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n", fMgr.absolutePathFileName)
  }

  fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt).\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&fMgrDest) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByLink(&fMgrDest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByLink(&fMgrDest)\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, err.Error())
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+"This File Manager instance is INVALID! Error='%v' ", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n",
      fMgr.absolutePathFileName)
  }

  fMgrDest, err := FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}.NewFromPathFileNameExtStr("+
      "dstPathFileNameExt).\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&fMgrDest) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByLinkByIo(&fMgrDest)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByLinkByIo(&fMgrDest)\n"+
      "fMgrDest.absolutePathFileName='%v'\nError='%v'\n",
      fMgrDest.absolutePathFileName, err.Error())
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "This File Manager instance is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n", fMgr.absolutePathFileName)
  }

  err = dir.IsDirMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'\n",
      err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir='%v'\nfMgr.fileNameExt='%v'\nError='%v'",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&newFMgr) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, newFMgr.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByIo(&newFMgr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByIo(&newFMgr)\n"+
      "newFMgr.absolutePathFileName='%v'\nError='%v'\n",
      newFMgr.absolutePathFileName, err.Error())
  }

  return nil
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+"This File Manager instance is INVALID! Error='%v' ", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n",
      fMgr.absolutePathFileName)
  }

  err = dir.IsDirMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir.absolutePath='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&newFMgr) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, newFMgr.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByIoByLink(&newFMgr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByIoByLink(&newFMgr)\n"+
      "newFMgr='%v'\nError='%v'\n",
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+"This File Manager instance is INVALID! Error='%v' ", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n", fMgr.absolutePathFileName)
  }

  err = dir.IsDirMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir.absolutePath='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&newFMgr) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\n Destination File='%v'\n",
      fMgr.absolutePathFileName, newFMgr.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByLink(&newFMgr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByLink(&newFMgr)\n"+
      "newFMgr.absolutePathFileName='%v'\nError='%v'\n",
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

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "This File Manager instance is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "FileName='%v'\n",
      fMgr.absolutePathFileName)
  }

  err = dir.IsDirMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir.absolutePath='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err.Error())
  }

  if fMgr.EqualAbsPaths(&newFMgr) {
    return fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\nDestination File='%v'\n",
      fMgr.absolutePathFileName, newFMgr.absolutePathFileName)
  }

  err = fMgr.CopyFileMgrByLinkByIo(&newFMgr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from fMgr.CopyFileMgrByLinkByIo(&newFMgr)\n"+
      "newFMgr='%v'\nError='%v'\n",
      newFMgr.absolutePathFileName, err.Error())
  }

  return nil
}

// CopyIn - Copies data from an incoming FileMgr object
// into the current FileMgr object.
//
// Note: Internal file pointer will not be copied.
//
func (fMgr *FileMgr) CopyIn(fmgr2 *FileMgr) {

  if fmgr2 == nil {
    panic("FileMgr.CopyIn() - Input parameter is a nil pointer!")
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

  err = fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return err
  }

  doesDirExist, err := fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n",
      err.Error())
  }

  if !doesDirExist {
    // Directory path does NOT exist. Create it!
    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

  }

  return nil
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

  err = fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return err
  }

  doesDirExist, err := fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n",
      err.Error())
  }

  if !doesDirExist {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+"- Directory Create Failed! %v\n", err.Error())
    }

  }

  err = fMgr.CreateThisFile()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error creating File. "+
      "fMgr.absolutePathFileName='%v' Error='%v' ",
      fMgr.absolutePathFileName, err.Error())
  }

  return nil
}

// CreateThisFile - Creates the File identified by FileMgr.absolutePathFileName.
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

  var err error

  err = fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return err
  }

  doesDirExist, err := fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n",
      err.Error())
  }

  if !doesDirExist {

    return fmt.Errorf(ePrefix+
      "Error: Directory Path DOES NOT EXIST!\n"+
      "DirPath='%v'\n", fMgr.dMgr.GetAbsolutePath())
  }

  //  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
  fOpenCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    _ = fMgr.filePtr.Close()
    fMgr.isFilePtrOpen = false
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    _ = fMgr.filePtr.Close()
    fMgr.isFilePtrOpen = false
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCfg)

  if err != nil {

    return fmt.Errorf(ePrefix+
      "Error opening file from fMgr.OpenThisFile(fileAccessCfg).\n"+
      "File Name='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  return nil
}

// DeleteThisFile - Deletes the file identified by FileMgr.absolutePathFileName
// in the current FileHelper structure.
func (fMgr *FileMgr) DeleteThisFile() error {

  ePrefix := "FileMgr.DeleteThisFile() "

  err := fMgr.IsFileMgrValid("")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: This FileMgr object is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
  }

  // If file does not exist, return with no error.
  if !fileDoesExist {
    return nil
  }

  if fMgr.filePtr != nil {

    err = fMgr.CloseThisFile()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "Error from fMgr.filePtr.Close()!  Error='%v'", err.Error())
    }
  }

  fMgr.isFilePtrOpen = false

  fMgr.dataMutex.Lock()

  err = os.Remove(fMgr.absolutePathFileName)

  fMgr.dataMutex.Unlock()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "- os.Remove(fMgr.absolutePathFileName) "+
      "returned an error.\n"+
      "absolutePathFileName='%v'\nError='%v'",
      fMgr.absolutePathFileName, err.Error())
  }

  fileDoesExist, err = fMgr.DoesThisFileExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fMgr.DoesThisFileExist()\n"+
      "fMgr.absolutePathFileName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  if fileDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: Attempted file deletion FAILED!. "+
      "File still exists.\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  fMgr.actualFileInfo = FileInfoPlus{}

  return nil
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

  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return false
  }

  return fMgr.doesAbsolutePathFileNameExist
}

// DoesThisFileExist - Returns a boolean value
// designating whether the file specified by the
// current FileMgr.absolutePathFileName field
// exists.
func (fMgr *FileMgr) DoesThisFileExist() (fileDoesExist bool, nonPathError error) {

  fMgr.dataMutex.Lock()

  ePrefix := "FileMgr.DoesThisFileExist() "
  fileDoesExist = false
  nonPathError = nil
  fMgrHelpr := fileMgrHelper{}

  fileDoesExist,
    nonPathError = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if nonPathError != nil {
    fileDoesExist = false
    fMgr.dataMutex.Unlock()
    return fileDoesExist, nonPathError
  }

  fMgr.dataMutex.Unlock()
  return fileDoesExist, nonPathError
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
    fMgr.isFilePtrOpen != fmgr2.isFilePtrOpen ||
    fMgr.fileRdrBufSize != fmgr2.fileRdrBufSize ||
    fMgr.fileWriterBufSize != fmgr2.fileWriterBufSize {

    return false
  }

  if !fMgr.fileAccessStatus.Equal(&fmgr2.fileAccessStatus) {
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

  f1 := strings.ToLower(fMgr.GetFileName())

  f2 := strings.ToLower(fmgr2.GetFileName())

  if f1 != f2 {
    return false
  }

  f1 = strings.ToLower(fMgr.GetFileExt())
  f2 = strings.ToLower(fmgr2.GetFileExt())

  if f1 != f2 {
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

  f1 := strings.ToLower(fMgr.GetAbsolutePathFileName())
  f2 := strings.ToLower(fmgr2.GetAbsolutePathFileName())

  if f1 != f2 {
    return false
  }

  return true
}

// Empty - resets all data fields in the FileMgr structure to
// their uninitialized or zero state.
func (fMgr *FileMgr) Empty() {

  fMgr.dataMutex.Lock()

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
  fMgr.fileAccessStatus.Empty()
  fMgr.actualFileInfo = FileInfoPlus{}
  fMgr.fileBufRdr = nil
  fMgr.fileBufWriter = nil
  fMgr.fileBytesWritten = 0
  fMgr.buffBytesWritten = 0
  fMgr.fileRdrBufSize = 0
  fMgr.fileWriterBufSize = 0

  fMgr.dataMutex.Unlock()
}

// FlushBytesToDisk - After Writing bytes to a file, use this
// method to commit the contents of the current file to
// stable storage.
func (fMgr *FileMgr) FlushBytesToDisk() error {

  ePrefix := "FileMgr.FlushBytesToDisk() "

  var err error

  if fMgr.filePtr != nil &&
    fMgr.fileBufWriter != nil &&
    fMgr.buffBytesWritten > 0 {

    fMgr.dataMutex.Lock()

    err = fMgr.fileBufWriter.Flush()

    fMgr.dataMutex.Unlock()

    if err != nil {
      return fmt.Errorf(ePrefix+"Error returned from fMgr.fileBufWriter.Flush(). Error='%v' ", err.Error())
    }

  }

  if fMgr.filePtr != nil &&
    fMgr.fileBytesWritten > 0 ||
    fMgr.buffBytesWritten > 0 {

    fMgr.dataMutex.Lock()

    err = fMgr.filePtr.Sync()

    fMgr.dataMutex.Unlock()

    if err != nil {
      return fmt.Errorf(ePrefix+"Error returned from fMgr.filePtr.Sync() Error='%v'", err.Error())
    }

  }

  return nil
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
func (fMgr *FileMgr) GetAbsolutePathFileName() string {
  return fMgr.absolutePathFileName
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

  return fMgr.fileBufRdr
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
  return fMgr.fileBufWriter
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
  return fMgr.buffBytesWritten + fMgr.fileBytesWritten
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
  return fMgr.fileExt
}

// GetFileInfo - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on the specific file identified
// by FileMgr.absolutePathFileName.
//
// An error will be triggered if the file path does NOT exist!
//
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }
func (fMgr *FileMgr) GetFileInfo() (fInfo os.FileInfo, err error) {

  fMgr.dataMutex.Lock()

  fInfo = nil
  err = nil
  filePathDoesExist := false
  ePrefix := "FileMgr.GetFileInfo() "

  fMgrHelpr := fileMgrHelper{}

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

  fMgr.dataMutex.Lock()

  fInfo = FileInfoPlus{}
  err = nil
  ePrefix := "FileMgr.GetFileInfoPlus() "

  fMgrHelpr := fileMgrHelper{}
  filePathDoesExist := false

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

  err := fMgr.ResetFileInfo()

  if err != nil {
    return time.Time{},
      fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return fMgr.actualFileInfo.ModTime(), nil
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
//                  'timeFormat'. If 'timeFormat' is empty or if 'timeFormat' is an
//                  an invalid format, the default format "2019-03-12 21:49:00:00"
//                  will be substituted.
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//
//                  Note: an error will be returned if the file identified by the current
//                  File Manager instance does NOT exist.
//
func (fMgr *FileMgr) GetFileModTimeStr(timeFormat string) (string, error) {

  ePrefix := "FileMgr.GetFileModTimeStr() "

  err := fMgr.ResetFileInfo()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  defaultFmt := "2006-01-02 15:04:05 -0700 MST"

  if timeFormat == "" {
    timeFormat = defaultFmt
  }

  t := fMgr.actualFileInfo.ModTime()

  tStr := t.Format(timeFormat)

  if tStr == timeFormat {
    tStr = t.Format(defaultFmt)
  }

  return tStr, nil
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
  return fMgr.fileName
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
  return fMgr.fileNameExt
}

// GetFilePermissionConfig - Returns a FilePermissionConfig instance encapsulating
// all of the permission information associated with this file.
//
// If the file does NOT exist, this method will return an error.
//
func (fMgr *FileMgr) GetFilePermissionConfig() (FilePermissionConfig, error) {

  ePrefix := "FileMgr.GetFilePermissionConfig() "

  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+" %v\n", err.Error())
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
        "fMgr='%v'\nError='%v'\n",
        fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return FilePermissionConfig{},
      errors.New(ePrefix +
        "The current (FileMgr) file DOES NOT EXIST!\n")
  }

  fPerm, err := FilePermissionConfig{}.NewByFileMode(fMgr.actualFileInfo.Mode())

  if err != nil {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "%v", err.Error())
  }

  return fPerm, nil
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

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Non-Path Error returned by fMgr.DoesThisFileExist()\n"+
        "fMgr='%v'\nError='%v'\n",
        fMgr.GetAbsolutePathFileName(), err.Error())
  }

  if !fileDoesExist {
    return "",
      errors.New(ePrefix + "The current (FileMgr) file DOES NOT EXIST!")
  }

  fPerm, err := FilePermissionConfig{}.NewByFileMode(fMgr.actualFileInfo.Mode())

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "%v", err.Error())

  }

  permissionText, err := fPerm.GetPermissionTextCode()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "%v", err.Error())

  }

  return permissionText, nil
}

// GetFilePtr - will return the internal *os.File pointer
// for this File Manager instance. Depending on circumstances,
// this pointer may be nil.
//
func (fMgr *FileMgr) GetFilePtr() *os.File {
  return fMgr.filePtr
}

// GetFileSize() - Returns os.FileInfo.Size() length in bytes for regular files;
// system-dependent for others.
//
// If the File Manager file does NOT exist, or if there is a file-path error, the
// value returned is -1.
//
func (fMgr *FileMgr) GetFileSize() int64 {

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    return -1
  }

  if !fileDoesExist {
    return -1
  }

  return fMgr.actualFileInfo.Size()
}

// GetOriginalPathFileName - Returns the path and file name
// used originally to configure this File Manager object.
//
// Note: The original path and file name will be adjusted
// to reflect the path operators used in the operating system
// for the host computer.
//
func (fMgr *FileMgr) GetOriginalPathFileName() string {

  return fMgr.originalPathFileName
}

// GetReaderBufferSize() - Returns the size for the internal
// Bufio Reader's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Read' Operation.
//
func (fMgr *FileMgr) GetReaderBufferSize() int {

  if fMgr.fileBufRdr != nil {
    fMgr.fileRdrBufSize = fMgr.fileBufRdr.Size()
  }

  return fMgr.fileRdrBufSize
}

// GetWriterBufferSize() - Returns the size for the internal
// Bufio Writers's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Write' Operation.
//
func (fMgr *FileMgr) GetWriterBufferSize() int {

  if fMgr.fileBufWriter != nil {
    fMgr.fileWriterBufSize = fMgr.fileBufWriter.Size()
  }

  return fMgr.fileWriterBufSize
}

// IsAbsolutePathFileNamePopulated - Returns a boolean value
// indicating whether absolute path and file name is
// initialized and populated.
//
func (fMgr *FileMgr) IsAbsolutePathFileNamePopulated() bool {

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil || len(fMgr.absolutePathFileName) == 0 {
    fMgr.isAbsolutePathFileNamePopulated = false
  } else {
    fMgr.isAbsolutePathFileNamePopulated = fileDoesExist
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

  if len(fMgr.fileExt) > 0 &&
    len(fMgr.fileName) > 0 {

    fMgr.isFileNameExtPopulated = true

  } else {

    fMgr.isFileNameExtPopulated = false
  }

  return fMgr.isFileNameExtPopulated
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

// IsFilePointerOpen - Returns a boolean value indicating
// whether the File Pointer (*os.File) for this File Manager
// instance is open, or not.
//
func (fMgr *FileMgr) IsFilePointerOpen() bool {
  if fMgr.filePtr == nil {
    fMgr.isFilePtrOpen = false
  } else {
    fMgr.isFilePtrOpen = true
  }

  return fMgr.isFilePtrOpen
}

// isInitialized - Returns a boolean indicating whether the FileMgr
// object is properly initialized.
//
func (fMgr *FileMgr) IsInitialized() bool {
  return fMgr.isInitialized
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
// currently exist, this method will return an error.
//
func (fMgr *FileMgr) MoveFileToNewDir(dirPath string) (newFMgr FileMgr, err error) {

  newFMgr = FileMgr{}
  err = nil

  ePrefix := "FileMgr.MoveFileToNewDir() "
  errCode := 0

  errCode, _, dirPath =
    FileHelper{}.isStringEmptyOrBlank(dirPath)

  if errCode < 0 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'dirPath' is a Zero length string!")
    return newFMgr, err
  }

  fMgr.dataMutex.Lock()

  fMgrHelpr := fileMgrHelper{}
  filePathDoesExist := false

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return newFMgr, err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: The source files does NOT exist.\n"+
      "srcFile (FileMgr)='%v' ",
      fMgr.absolutePathFileName)
    return newFMgr, err
  }

  dMgr, err2 := DirMgr{}.New(dirPath)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from DirMgr{}.NewFromPathFileNameExtStr(dirPath). dirPath='%v'  Error='%v'",
      dirPath, err2.Error())
    return newFMgr, err
  }

  err2 = dMgr.IsDirMgrValid("")

  if err2 != nil {

    err = fmt.Errorf(ePrefix+
      "Error: Input parameter 'dirPath' "+
      "generated an INVALID DirMgr.\n"+
      "dirPath='%v'\nError='%v'\n",
      dirPath, err2)

    return newFMgr, err
  }

  pathExists, err2 := dMgr.DoesThisDirectoryExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: dMgr.DoesThisDirectoryExist() "+
      "returned a non-path error.\n"+
      "dirPath='%v'\nError='%v'\n",
      dirPath, err2)

    return newFMgr, err
  }

  if !pathExists {
    err = fmt.Errorf(ePrefix+
      "Error: Target Destination path DOES NOT EXIST!\n"+
      "dirPath='%v'",
      dirPath)
    return newFMgr, err
  }

  targetFile := dMgr.GetAbsolutePathWithSeparator() +
    fMgr.fileNameExt

  fMgr.dataMutex.Lock()

  err2 = FileHelper{}.MoveFile(
    fMgr.absolutePathFileName, targetFile)

  if err2 != nil {
    newFMgr = FileMgr{}
    err = fmt.Errorf(ePrefix+
      "Error returned by FileHelper{}.MoveFile("+
      "fMgr.absolutePathFileName, newFMgr.absolutePathFileName)\n"+
      "fMgr.absolutePathFileName='%v'\n"+
      "newFMgr.absolutePathFileName='%v'\nError='%v'",
      fMgr.absolutePathFileName, newFMgr.absolutePathFileName, err2.Error())

  } else {

    newFMgr, err2 = FileMgr{}.New(targetFile)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error creating new File Manager!\n"+
        "File was NOT Moved!\n"+
        "Error returned by FileMgr{}.New(targetFile)\n"+
        "targetFile='%v'\nError='%v'",
        targetFile, err2.Error())

      newFMgr = FileMgr{}
    } else {
      err = nil
    }
  }

  fMgr.dataMutex.Unlock()
  return newFMgr, err
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
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
// If the input parameter 'dMgr' does not contain a path which currently
// exists, an error will be returned.
//
func (fMgr *FileMgr) MoveFileToNewDirMgr(dMgr DirMgr) (newFMgr FileMgr, err error) {
  ePrefix := "FileMgr.MoveFileToNewDirMgr() "
  newFMgr = FileMgr{}
  err = nil

  err2 := dMgr.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error: Input parameter 'dMgr' reports as INVALID! "+
      "Error='%v'", err2.Error())
    return newFMgr, err
  }

  fMgr.dataMutex.Lock()

  fMgrHelpr := fileMgrHelper{}
  filePathDoesExist := false

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return newFMgr, err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: The source file identified by the "+
      "current FileMgr object DOES NOT EXIST!\n"+
      "source file (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)
    return newFMgr, err
  }

  newFMgr = FileMgr{}

  filePathDoesExist, err = dMgr.DoesThisDirectoryExist()

  if err != nil {
    return newFMgr, err
  }

  if !filePathDoesExist {

    err = fmt.Errorf(ePrefix+
      "Error: Destination (dMgr) path DOES NOT EXIST!.\n"+
      "dMgr='%v'\n", dMgr.absolutePath)
    return newFMgr, err
  }

  destPathFileName := dMgr.GetAbsolutePathWithSeparator() + fMgr.fileNameExt

  fMgr.dataMutex.Lock()

  fh := FileHelper{}

  err2 = fh.MoveFile(fMgr.absolutePathFileName, destPathFileName)

  if err2 != nil {
    newFMgr = FileMgr{}
    err = fmt.Errorf(ePrefix+
      "Error returned from "+
      "fh.MoveFile(fMgr.absolutePathFileName, destPathFileName).\n"+
      "fMgr.absolutePathFileName='%v'\ndestPathFileName='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, destPathFileName, err2.Error())

  } else {

    newFMgr, err2 = FileMgr{}.NewFromPathFileNameExtStr(destPathFileName)

    if err2 != nil {
      newFMgr = FileMgr{}
      err = fmt.Errorf(ePrefix+
        "Error returned by FileMgr{}.NewFromPathFileNameExtStr(destPathFileName). destPathFileName='%v' "+
        "Error='%v'", destPathFileName, err2.Error())
    } else {
      err = nil
    }
  }

  fMgr.dataMutex.Unlock()

  return newFMgr, err
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

  fMgrOut := FileMgr{}

  isEmpty, err := fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error returned from fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt) "+
      "pathFileNameExt='%v'  Error='%v'\n", pathFileNameExt, err.Error())
  }

  if isEmpty {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Returned FileMgr is Empty! pathFileNameExt='%v' ", pathFileNameExt)
  }

  err = fMgrOut.IsFileMgrValid("")

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "New File Manager is INVALID! pathFileNameExt='%v' Error='%v' ", pathFileNameExt, err.Error())
  }

  return fMgrOut, nil
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

  if len(fileNameExt) == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: Input Parameter fileNameExt is a zero length string!")
  }

  err := dirMgr.IsDirMgrValid("")

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error: Input parameter 'dirMgr' is INVALID! Error='%v' ", err.Error())
  }

  fmgr2 := FileMgr{}

  isEmpty, err := fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error returned by fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt). "+
      "Error='%v'\n", err.Error())
  }

  if isEmpty {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Empty FileMgr object returned by fmgr2.SetFileMgrFromDirMgrFileName("+
      "dirMgr, fileNameExt) dirMgr.path='%v' fileNameExt='%v'", dirMgr.path, fileNameExt)
  }

  err = fmgr2.IsFileMgrValid("")

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "The new File Manager is INVALID! Error='%v' ", err.Error())
  }

  return fmgr2, nil
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

  dirMgr, err := DirMgr{}.New(dirStr)

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirMgr). dirStr='%v'  Error='%v'",
        dirStr, err.Error())
  }

  fmgr2 := FileMgr{}

  isEmpty, err := fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExtStr)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error returned by fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExtStr). "+
      "Directory Path='%v' File Name='%v' Error='%v'\n",
      dirMgr.GetAbsolutePath(), fileNameExtStr, err.Error())
  }

  if isEmpty {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Empty FileMgr object returned by fmgr2.SetFileMgrFromDirMgrFileName("+
      "dirMgr, fileNameExtStr) dirMgr.path='%v' fileNameExt='%v'",
      dirMgr.GetAbsolutePath(), fileNameExtStr)
  }

  err = fmgr2.IsFileMgrValid("")

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "The new File Manager is INVALID! Error='%v' ", err.Error())
  }

  return fmgr2, nil

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
  var err error

  if info == nil {
    return FileMgr{},
      errors.New(ePrefix + "Error: Input parameter 'info' is 'nil' and INVALID!")
  }

  fileName := info.Name()

  if len(fileName) == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: info.Name() is an empty string and therefore INVALID!")
  }

  if len(dirPathStr) == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: input parameter 'dirPathStr' is an empty string and therefore INVALID!.")
  }

  dMgr, err := DirMgr{}.New(dirPathStr)

  if err != nil {
    return FileMgr{},
      fmt.Errorf("Input Parameter 'dirPathStr' is INVALID! DirMgr{}.New(dirPathStr) "+
        "Error='%v' ", err.Error())
  }

  fmgr2 := FileMgr{}

  isEmpty, err := fmgr2.SetFileMgrFromDirMgrFileName(dMgr, fileName)

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error returned from fmgr2.SetFileMgrFromDirMgrFileName(dMgr, fileName). "+
        "dMgr='%v', fileName='%v' Error='%v'\n",
        dMgr.GetAbsolutePath(), fileName, err.Error())
  }

  if isEmpty {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error returned FileMgr is Empty! directory path='%v' fileName='%v'",
      dMgr.GetAbsolutePath(), fileName)
  }

  fmgr2.actualFileInfo = FileInfoPlus{}.NewFromPathFileInfo(dMgr.GetAbsolutePath(), info)

  err = fmgr2.IsFileMgrValid("")

  if err != nil {
    return FileMgr{},
      fmt.Errorf(ePrefix+"The New File Manager is INVALID! Error='%v' ", err.Error())
  }

  return fmgr2, nil
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

  if pathFileNameExt == "" {
    return FileMgr{}, errors.New(ePrefix + "-Error: pathFileNameExt is Empty!")
  }

  fMgrOut := FileMgr{}

  isEmpty, err := fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt)

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error returned from fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt) "+
      "pathFileNameExt='%v'  Error='%v' \n", pathFileNameExt, err.Error())
  }

  if isEmpty {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Returned FileMgr is Empty! pathFileNameExt='%v' ", pathFileNameExt)
  }

  err = fMgrOut.IsFileMgrValid("")

  if err != nil {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "New File Manager is INVALID! pathFileNameExt='%v' Error='%v'", pathFileNameExt, err.Error())
  }

  return fMgrOut, nil
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

  fMgr.dataMutex.Lock()

  fMgrHelpr := fileMgrHelper{}
  _,
    err = fMgrHelpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  fMgr.dataMutex.Unlock()

  if err != nil {
    return err
  }

  if fMgr.filePtr != nil {
    _ = fMgr.CloseThisFile()
  }

  fMgr.filePtr = nil
  fMgr.isFilePtrOpen = false
  fMgr.fileAccessStatus.Empty()
  fMgr.fileBytesWritten = 0
  fMgr.buffBytesWritten = 0
  fMgr.fileBufRdr = nil
  fMgr.fileBufWriter = nil

  err = fileAccessCtrl.IsValid()

  if err != nil {

    return fmt.Errorf(ePrefix+"Input parameter 'fileAccessCtrl' is INVALID!\n"+
      "%v\n", err.Error())
  }

  filePathDoesExist, err := fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Non-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
      "dMgr='%v'\nError='%v'\n",
      fMgr.dMgr.GetAbsolutePath(), err.Error())
  }

  if !filePathDoesExist {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

  }

  fMgr.fileAccessStatus = fileAccessCtrl.CopyOut()

  fOpenParm, fPermParm, err := fMgr.fileAccessStatus.GetFileOpenAndPermissionCodes()

  if err != nil {
    fMgr.fileAccessStatus.Empty()
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fMgr.dataMutex.Lock()

  fMgr.filePtr, err = os.OpenFile(fMgr.absolutePathFileName, fOpenParm, fPermParm)

  fMgr.dataMutex.Unlock()

  if err != nil {
    fMgr.filePtr = nil
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0
    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil

    return fmt.Errorf(ePrefix+
      "Error opening file from os.OpenFile(): '%v' Error= '%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  fMgr.isFilePtrOpen = true

  return nil
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
      "The file to be opened as 'Read-Only' does not Exist!\n"+
      "(FileMgr) FileName: %v\n",
      fMgr.absolutePathFileName)

  }

  if fMgr.filePtr != nil {
    _ = fMgr.CloseThisFile()
  }

  fileOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(), FOpenMode.ModeNone())

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  return nil
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
  var filePathDoesExist bool

  ePrefix := "FileMgr.OpenThisFileWriteOnly() "

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

  if fMgr.filePtr != nil {
    _ = fMgr.CloseThisFile()
  }

  if !filePathDoesExist {

    err = fMgr.CreateDirAndFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

    err = fMgr.CloseThisFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }
  }

  fileOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("--w--w--w-")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  return nil
}

// OpenThisFileWriteOnlyAppend - Opens the current file for 'Write Only'
// with an 'Append' mode. All bytes written to the file will be written
// at the end of the current file and none of the file's original content
// will be overwritten.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will will create them.
//
func (fMgr *FileMgr) OpenThisFileWriteOnlyAppend() error {
  var err error
  var filePathDoesExist bool

  ePrefix := "FileMgr.OpenThisFileWriteOnlyAppend() "

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

  if fMgr.filePtr != nil {
    _ = fMgr.CloseThisFile()
  }

  if !filePathDoesExist {

    err = fMgr.CreateDirAndFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

    err = fMgr.CloseThisFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

  }

  fileOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeAppend())

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("--w--w--w-")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  return nil
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
  var filePathDoesExist bool

  ePrefix := "FileMgr.OpenThisFileReadWrite() "

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

  if fMgr.filePtr != nil {
    _ = fMgr.CloseThisFile()
  }

  if !filePathDoesExist {

    err = fMgr.CreateDirAndFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

    err = fMgr.CloseThisFile()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v", err.Error())
    }

  }

  fileOpenCfg, err :=
    FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  filePermCfg, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fMgr.OpenThisFile(fileAccessCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  return nil
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
    return bytesRead, err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+"ERROR: The File Manager (FileMgr) file "+
      "DOES NOT EXIST!\n"+
      "FileMgr='%v'\n", fMgr.absolutePathFileName)
    return bytesRead, err
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeReadOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return bytesRead, err
    }

  }

  err = nil

  fMgr.dataMutex.Lock()

  bytesRead, err2 = ioutil.ReadAll(fMgr.filePtr)

  fMgr.dataMutex.Unlock()

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned by ioutil.ReadAll(fMgr.filePtr). "+
        "fileName='%v' Errors='%v'",
        fMgr.absolutePathFileName, err2.Error())

  }

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
    return bytesRead, err
  }

  if !filePathDoesExist {

    err = fmt.Errorf(ePrefix+
      "The file to be opened Reading does not Exist!\n"+
      "(FileMgr) FileName: %v\n",
      fMgr.absolutePathFileName)

    return bytesRead, err
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeReadOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return bytesRead, err
    }

  }

  err = nil

  fMgr.dataMutex.Lock()

  if fMgr.fileBufRdr == nil {
    if fMgr.fileRdrBufSize > 0 {
      fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
    } else {
      fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
    }
  }

  bytesRead, err2 = fMgr.fileBufRdr.Read(byteBuff)

  fMgr.dataMutex.Unlock()

  if err2 != nil &&
    err2 == io.EOF {

    err = err2

  } else if err2 != nil {

    err = fmt.Errorf("Error returned by fMgr.fileBufRdr.Read(byteBuff). "+
      "File='%v' Error='%v' ", fMgr.absolutePathFileName, err2.Error())
  }

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
func (fMgr *FileMgr) ReadFileLine(delim byte) (bytesRead []byte, err error) {

  ePrefix := "FileMgr.ReadBuffBytes() "
  bytesRead = []byte{}
  err = nil

  var err2 error

  err2 = fMgr.IsFileMgrValid("")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'",
        fMgr.absolutePathFileName, err2.Error())

    return bytesRead, err
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeReadOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return bytesRead, err
    }

  }

  err = nil

  fMgr.dataMutex.Lock()

  if fMgr.fileBufRdr == nil {
    if fMgr.fileRdrBufSize > 0 {
      fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
    } else {
      fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
    }
  }

  bytesRead, err2 = fMgr.fileBufRdr.ReadBytes(delim)

  fMgr.dataMutex.Unlock()

  if err2 != nil &&
    err2 == io.EOF {
    err = err2

  } else if err2 != nil {

    err = fmt.Errorf("Error returned from fMgr.fileBufRdr.ReadBytes(delim). "+
      "Error='%v' ", err2.Error())
  }

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

  err2 = fMgr.IsFileMgrValid("")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'",
        fMgr.absolutePathFileName, err2.Error())

    return stringRead, err
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeReadOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return stringRead, err
    }

  }

  err = nil

  fMgr.dataMutex.Lock()

  if fMgr.fileBufRdr == nil {
    if fMgr.fileRdrBufSize > 0 {
      fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
    } else {
      fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
    }
  }

  stringRead, err2 = fMgr.fileBufRdr.ReadString(delim)

  fMgr.dataMutex.Unlock()

  if err2 != nil &&
    err2 == io.EOF {
    err = err2

  } else if err2 != nil {

    err = fmt.Errorf("Error returned from fMgr.fileBufRdr.ReadString(delim). "+
      "Error='%v' ", err2.Error())
  }

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

  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return err
  }

  if !fMgr.doesAbsolutePathFileNameExist {
    return fmt.Errorf(ePrefix + "Current FileMgr file DOES NOT EXIST!")
  }

  return nil
}

// SetReaderBufferSize - Sets the Read Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
//
func (fMgr *FileMgr) SetReaderBufferSize(readBuffSize int) {
  fMgr.fileRdrBufSize = readBuffSize
}

// SetWriterBufferSize - Sets the Write Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
//
func (fMgr *FileMgr) SetWriterBufferSize(writeBuffSize int) {
  fMgr.fileWriterBufSize = writeBuffSize
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
    return errors.New(ePrefix + "Error: Input parameter 'info' is 'nil' and INVALID!")
  }

  if info.Name() == "" {
    return errors.New(ePrefix + "Error: info.Name() is an EMPTY string!")
  }

  if info.IsDir() {
    return errors.New(ePrefix + "info.IsDir()=='true'. This is a Directory NOT A FILE!")
  }

  if strings.ToLower(info.Name()) != strings.ToLower(fMgr.fileNameExt) {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'info' does NOT match current FileMgr file name. "+
      "FileMgr File Name='%v' info File Name='%v' ", fMgr.fileNameExt, info.Name())
  }

  fMgr.dataMutex.Lock()

  fMgr.actualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)

  fMgr.dataMutex.Unlock()

  if !fMgr.actualFileInfo.IsFInfoInitialized {
    return fmt.Errorf(ePrefix+
      "Error: Failed to initialize fMgr.actualFileInfo object. info.Name()='%v'",
      info.Name())
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
//                  values are set to valid legitimate values.
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
  isEmpty = true
  err = nil
  fh := FileHelper{}

  err2 := dMgr.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: Input parameter 'dMgr' is INVALID! dMgr.path='%v'  Error='%v'",
      dMgr.path, err2.Error())
    return
  }

  errCode, _, fileNameExt := fh.isStringEmptyOrBlank(fileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'fileNameExt' is a Zero length string!")
    return
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'fileNameExt' consists entirely of blank spaces!")
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

  fMgr.Empty()

  fMgr.dataMutex.Lock()

  fMgr.dMgr = dMgr.CopyOut()

  s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(adjustedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFileNameWithoutExt(adjustedFileNameExt). "+
      "adjustedFileNameExt='%v'  Error='%v' ", adjustedFileNameExt, err2.Error())
    fMgr.dataMutex.Unlock()
    fMgr.Empty()
    isEmpty = true
    return
  }

  if fNameIsEmpty {
    err = fmt.Errorf(ePrefix+
      "Error: fileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt) "+
      "is Zero length string! adjustedFileNameExt='%v'  ", adjustedFileNameExt)
    fMgr.dataMutex.Unlock()
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
    fMgr.dataMutex.Unlock()
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

  _,
    filePathDoesExist,
    fInfoPlus,
    nonPathError :=
    FileHelper{}.doesPathFileExist(
      fMgr.absolutePathFileName,
      PreProcPathCode.None(), // Do NOT perform pre-processing on path
      ePrefix,
      "fMgr.absolutePathFileName")

  if filePathDoesExist && nonPathError == nil {
    fMgr.doesAbsolutePathFileNameExist = true
    fMgr.actualFileInfo = fInfoPlus.CopyOut()
  } else {
    fMgr.doesAbsolutePathFileNameExist = false
    fMgr.actualFileInfo = FileInfoPlus{}
  }

  fMgr.isInitialized = true

  err = nil
  isEmpty = false

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
  fh := FileHelper{}

  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'pathFileNameExt' is a zero length or empty string!")
    return
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'pathFileNameExt' consists entirely of blank spaces!")
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

// WriteBytesToFile a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
//
// If the number of bytes written to the file is less than len(bytes),
// an error will be returned.
//
// This method uses the 'bufio' package.
//
func (fMgr *FileMgr) WriteBytesToFile(bytes []byte) (numBytesWritten int, err error) {

  ePrefix := "FileMgr.WriteBytesToFile() "
  err = nil
  numBytesWritten = 0

  var err2 error

  err2 = fMgr.IsFileMgrValid("")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'",
        fMgr.absolutePathFileName, err2.Error())

    return numBytesWritten, err
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeWriteOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return numBytesWritten, err
    }

  }

  err = nil

  fMgr.dataMutex.Lock()

  if fMgr.fileBufWriter == nil {
    if fMgr.fileWriterBufSize > 0 {
      fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
    } else {
      fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
    }
  }

  numBytesWritten, err2 = fMgr.fileBufWriter.Write(bytes)

  fMgr.dataMutex.Unlock()

  fMgr.buffBytesWritten += uint64(numBytesWritten)

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned from fMgr.fileBufWriter.Write(bytes). Output File='%v'. "+
        "Error='%v'", fMgr.absolutePathFileName, err2.Error())

  }

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

  err2 := fMgr.IsFileMgrValid("")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'",
        fMgr.absolutePathFileName, err2.Error())

    return numBytesWritten, err
  }

  fMgr.dataMutex.Lock()

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {

    if fOpenType == FOpenType.TypeWriteOnly() ||
      fOpenType == FOpenType.TypeReadWrite() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    err2 != nil ||
    invalidAccessType {

    fMgr.dataMutex.Unlock()
    // If the path and file name do not exist, this method will
    // attempt to create said path and file name.
    err2 = fMgr.OpenThisFileReadWrite()

    if err2 != nil {
      err =
        fmt.Errorf(ePrefix+
          " - fMgr.OpenThisFileReadWrite() returned errors: %v",
          err2.Error())

      return numBytesWritten, err
    }

    fMgr.dataMutex.Lock()
  }

  err = nil

  if fMgr.fileBufWriter == nil {
    if fMgr.fileWriterBufSize > 0 {
      fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
    } else {
      fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
    }
  }

  numBytesWritten, err2 = fMgr.fileBufWriter.WriteString(str)

  fMgr.buffBytesWritten += uint64(numBytesWritten)

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned from fMgr.filePtr.WriteString(str). Output File='%v'. "+
        "Error='%v'", fMgr.absolutePathFileName, err2.Error())

  }

  fMgr.dataMutex.Unlock()
  return numBytesWritten, err
}
