package pathfileops

import (
  "errors"
  "fmt"
  "os"
  "time"
)

type fileMgrHelper struct {
  fMgr FileMgr
}

// fMgrDoesPathFileExist - Used by FileMgr type to test
// for the existence of a file path.
func (fMgrHlpr *fileMgrHelper) doesFileMgrPathFileExist(
  fileMgr *FileMgr,
  preProcessCode PreProcessPathCode,
  errorPrefix,
  filePathTitle string) (filePathDoesExist bool,
  nonPathError error) {

  filePathDoesExist = false
  nonPathError = nil

  ePrefixCurrMethod := "fileMgrHelper.doesDirPathExist() "

  if len(errorPrefix) == 0 {
    errorPrefix = ePrefixCurrMethod
  } else {
    errorPrefix = errorPrefix + "- " + ePrefixCurrMethod + "\n"
  }

  if len(filePathTitle) == 0 {
    filePathTitle = "filePath"
  }

  errCode := 0

  errCode,
    _, fileMgr.absolutePathFileName =
    FileHelper{}.isStringEmptyOrBlank(fileMgr.absolutePathFileName)

  if errCode == -1 {
    fileMgr.isAbsolutePathFileNamePopulated = false
    nonPathError = fmt.Errorf(errorPrefix+
      "Error: '%v' is an empty string!", filePathTitle)
    return filePathDoesExist, nonPathError
  }

  if errCode == -2 {
    fileMgr.isAbsolutePathFileNamePopulated = false
    nonPathError = fmt.Errorf(errorPrefix+
      "Error: '%v' consists of blank spaces!", filePathTitle)
    return filePathDoesExist, nonPathError
  }

  if !fileMgr.isInitialized {
    nonPathError = errors.New(errorPrefix +
      "Error: This data structure is NOT initialized.\n" +
      "fileMgr.isInitialized='false'\n")
    return filePathDoesExist, nonPathError
  }

  var err error

  err = fileMgr.dMgr.IsDirMgrValid(errorPrefix)

  if err != nil {
    nonPathError = fmt.Errorf("FileMgr Directory Manager INVALID!\n"+
      "Error='%v'", err.Error())
    return filePathDoesExist, nonPathError
  }

  if preProcessCode == PreProcPathCode.PathSeparator() {

    fileMgr.absolutePathFileName = FileHelper{}.AdjustPathSlash(fileMgr.absolutePathFileName)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    fileMgr.absolutePathFileName, err = FileHelper{}.MakeAbsolutePath(fileMgr.absolutePathFileName)

    if err != nil {
      nonPathError = fmt.Errorf(errorPrefix+
        "FileHelper{}.MakeAbsolutePath() FAILED!\n"+
        "%v='%v'"+
        "%v", filePathTitle, fileMgr.absolutePathFileName, err.Error())
      return filePathDoesExist, nonPathError
    }
  }

  var info os.FileInfo

  for i := 0; i < 3; i++ {

    filePathDoesExist = false
    nonPathError = nil

    info, err = os.Stat(fileMgr.absolutePathFileName)

    if err != nil {

      if os.IsNotExist(err) {
        fileMgr.doesAbsolutePathFileNameExist = false
        fileMgr.actualFileInfo = FileInfoPlus{}
        fileMgr.fileAccessStatus.Empty()
        filePathDoesExist = false
        nonPathError = nil
        _ = fileMgr.dMgr.DoesPathExist()
        _ = fileMgr.dMgr.DoesAbsolutePathExist()
        return filePathDoesExist, nonPathError
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be tested
      // up to 3-times before it is returned.
      nonPathError = fmt.Errorf(errorPrefix+
        "Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n",
        filePathTitle, filePathTitle, fileMgr.absolutePathFileName, err.Error())

      fileMgr.doesAbsolutePathFileNameExist = false
      fileMgr.actualFileInfo = FileInfoPlus{}
      filePathDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      filePathDoesExist = true
      nonPathError = nil
      fileMgr.doesAbsolutePathFileNameExist = true
      fileMgr.actualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)

      permCode, err := FilePermissionConfig{}.NewByFileMode(fileMgr.actualFileInfo.Mode())

      if err == nil {
        _ = fileMgr.fileAccessStatus.SetFilePermissionCodes(permCode)
      }

      _ = fileMgr.dMgr.DoesPathExist()
      _ = fileMgr.dMgr.DoesAbsolutePathExist()

      return filePathDoesExist, nonPathError
    }

    time.Sleep(30 * time.Millisecond)
  }

  if !filePathDoesExist {
    fileMgr.fileAccessStatus.Empty()
  }

  _ = fileMgr.dMgr.DoesPathExist()
  _ = fileMgr.dMgr.DoesAbsolutePathExist()

  return filePathDoesExist, nonPathError
}

// closeFile - Helper method for File Manager
// (FileMgr) Type. It is designed to close a
// file.
//
func (fMgrHlpr *fileMgrHelper) closeFile(
  fMgr *FileMgr, ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.closeFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  errs := make([]error, 0)

  var err2, err3 error
  fileDoesExist := false

  fileDoesExist,
    err2 = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err2 != nil {

    errs = append(errs, err2)

    err2 = fMgrHlpr.flushBytesToDisk(fMgr, ePrefix)

    if err2 != nil {
      errs = append(errs, err2)
    }

    err2 = nil

    if fMgr.filePtr != nil {
      err2 = fMgr.filePtr.Close()
    }

    if err2 != nil {
      errs = append(errs, err2)
    }

    fMgr.isFilePtrOpen = false

    fMgr.fileAccessStatus.Empty()

    if err2 != nil {
      errs = append(errs, err2)
    }

    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return fMgrHlpr.consolidateErrors(errs)
  }

  if !fileDoesExist {

    err2 = nil

    if fMgr.filePtr != nil {
      err2 = fMgr.filePtr.Close()
    }

    if err2 != nil {
      errs = append(errs, err2)
    }

    fMgr.filePtr = nil
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()

    if err2 != nil {
      errs = append(errs, err2)
    }

    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return fMgrHlpr.consolidateErrors(errs)
  }

  if fMgr.filePtr == nil {
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()

    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return fMgrHlpr.consolidateErrors(errs)
  }

  // fMgr.filePtr != nil
  err2 = fMgrHlpr.flushBytesToDisk(fMgr, ePrefix)

  if err2 != nil {
    errs = append(errs, err2)
  }

  err3 = fMgr.filePtr.Close()

  if err3 != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned by fMgr.filePtr.Close()\n"+
      "Error='%v'\n", err3.Error())
    errs = append(errs, err2)
  }

  fMgr.filePtr = nil
  fMgr.isFilePtrOpen = false
  fMgr.fileAccessStatus.Empty()
  fMgr.fileBufRdr = nil
  fMgr.fileBufWriter = nil
  fMgr.fileBytesWritten = 0
  fMgr.buffBytesWritten = 0

  return fMgrHlpr.consolidateErrors(errs)
}

// consolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (fMgrHlpr *fileMgrHelper) consolidateErrors(errs []error) error {

  lErrs := len(errs)

  if lErrs == 0 {
    return nil
  }

  errStr := ""

  for i := 0; i < lErrs; i++ {

    if i == (lErrs - 1) {
      errStr += fmt.Sprintf("%v", errs[i].Error())
    } else {
      errStr += fmt.Sprintf("%v\n", errs[i].Error())
    }
  }

  return fmt.Errorf("%v", errStr)
}

// copyFileToDestFileMgrSetup - Helper method used by FileMgr
// Copy To Destination File Manager routines. This method
// performs standardized setup and error checking functions.
//
func (fMgrHlpr *fileMgrHelper) copyFileToDestFileMgrSetup(
  fMgr,
  fMgrDest *FileMgr,
  ePrefix string) (err error) {

  ePrefixCurrMethod := "fileMgrHelper.copyFileToDestFileMgrSetup() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgrDest == nil {
    err = errors.New(ePrefix +
      "Error: Destination File Manager (fMgrDest) is a nil pointer!\n")
    return err
  }

  filePathDoesExist,
    err2 := fMgrHlpr.doesFileMgrPathFileExist(fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err2 != nil {
    err = err2
    return err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: This File Manager file does NOT exist!\n"+
      "(FileMgr) File Name:'%v' ", fMgr.absolutePathFileName)
    return err
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied.\n"+
      "Copy Operation Aborted.\n"+
      "Source File (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)
    return err
  }

  if fMgr.actualFileInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "Error: Source File File Manger exists, but "+
      "it is classified as a Directory!\n"+
      "Source File (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)

    return err
  }

  err2 = fMgrDest.IsFileMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: The Destination FileMgr object is INVALID!\n"+
      "Error='%v'\n",
      err2.Error())
    return err
  }

  filePathDoesExist, err2 =
    fMgrDest.dMgr.DoesThisDirectoryExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Destination Directory Error: Non-Path Error returned by\n"+
      "fMgrDest.dMgr.DoesThisDirectoryExist().\n"+
      "fMgrDest.dMgr='%v'\n"+
      "Error='%v'\n",
      fMgrDest.dMgr.absolutePath, err2.Error())

    return err
  }

  if !filePathDoesExist {

    err2 = fMgrDest.dMgr.MakeDir()

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error: Attempted creation of destination directory FAILED!\n"+
        "Destination Directory (fMgrDest.dMgr)='%v'\n"+
        "Error= '%v'",
        fMgrDest.dMgr.GetAbsolutePath(), err2.Error())

      return err
    }

    filePathDoesExist,
      err2 = fMgrDest.dMgr.DoesThisDirectoryExist()

    if err2 != nil {
      err = fmt.Errorf("Non-Path Error returned by "+
        "fMgrDest.dMgr.DoesThisDirectoryExist()\n"+
        "fMgrDest.dMgr='%v'\n"+
        "Error='%v'\n",
        fMgrDest.dMgr.absolutePath, err2.Error())

      return err
    }

    if !filePathDoesExist {
      err = fmt.Errorf(ePrefix+
        "Error: Attempted verification of destination directory "+
        "creation FAILED!\n"+
        "Destination Directory DOES NOT EXIST!\n"+
        "Destination Directory (fMgrDest.dMgr)='%v'\n",
        fMgrDest.dMgr.absolutePath)

      return err
    }
  }

  filePathDoesExist,
    err2 = fMgrDest.DoesThisFileExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Non-Path Error returned by fMgrDest.DoesThisFileExist().\n"+
      "Error='%v'", err2.Error())

    return err
  }

  if filePathDoesExist &&
    fMgrDest.actualFileInfo.Mode().IsDir() {

    err = fmt.Errorf(ePrefix+
      "Error: Destination File (fMgrDest) exists, but "+
      "it is classified as a Directory!\n"+
      "Destination File (fMgrDest)='%v'\n",
      fMgrDest.absolutePathFileName)

    return err
  }

  if filePathDoesExist &&
    !fMgrDest.actualFileInfo.Mode().IsRegular() {

    err = fmt.Errorf(ePrefix+
      "Error: Destination file is a Non-Regular "+
      "File and cannot be the target of a copy operations.\n"+
      "Copy Operation Aborted.\n"+
      "Destination File (fMgrDest)='%v'\n",
      fMgrDest.absolutePathFileName)

    return err
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    err = fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File (FileMgr)='%v'\n"+
      "Destination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)

    return err
  }

  err = nil
  return err
}

// copyFileToDirSetup - Helper method used by FileMgr Copy
// to Directory routines for standardized setup and error
// checking.
//
func (fMgrHlpr *fileMgrHelper) copyFileToDirSetup(
  fMgr *FileMgr,
  dir DirMgr,
  ePrefix string) (fMgrDest FileMgr, err error) {

  err = nil
  fMgrDest = FileMgr{}
  ePrefixCurrMethod := "fileMgrHelper.copyFileToDirSetup() "

  originalEPrefix := ePrefix

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
    originalEPrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  err2 := dir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'", err2.Error())
    return fMgrDest, err
  }

  fMgrDest, err2 = FileMgr{}.NewFromDirMgrFileNameExt(
    dir, fMgr.fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir.absolutePath='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err2.Error())
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  err = fMgrHlpr.copyFileToDestFileMgrSetup(fMgr, &fMgrDest, originalEPrefix)

  return fMgrDest, err
}

// copyFileToDestStrSetup - Helper method used by File Manager
// copy strings. These copy operations format the copy destination
// as a string. This method performs standardized setup and error
// checking functions.
//
func (fMgrHlpr *fileMgrHelper) copyFileToDestStrSetup(
  fMgr *FileMgr,
  dstPathFileNameExt,
  ePrefix string) (fMgrDest FileMgr, err error) {

  fMgrDest = FileMgr{}
  err = nil
  ePrefixCurrMethod := "fileMgrHelper.copyFileToDestStrSetup() "

  originalEPrefix := ePrefix

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
    originalEPrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var errCode int

  errCode,
    _,
    dstPathFileNameExt =
    FileHelper{}.isStringEmptyOrBlank(dstPathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "Error: 'dstPathFileNameExt' is an Empty or " +
      "Zero Length String!\n")
    return fMgrDest, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Error: 'dstPathFileNameExt' consists entirely of " +
      "blank spaces!\n")
    return fMgrDest, err
  }

  var err2 error

  fMgrDest, err2 =
    FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by FileMgr{}."+
      "NewFromPathFileNameExtStr(dstPathFileNameExt).\n"+
      "dstPathFileNameExt='%v'\nError='%v'\n",
      dstPathFileNameExt, err2.Error())
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  err = fMgrHlpr.copyFileToDestFileMgrSetup(
    fMgr,
    &fMgrDest,
    originalEPrefix)

  if err != nil {
    fMgrDest = FileMgr{}
  }

  return fMgrDest, err
}

func (fMgrHlpr *fileMgrHelper) createDirectory(
  fMgr *FileMgr,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.createDirectory() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  _,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  err = fMgr.dMgr.MakeDir()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n",
      err.Error())
  }

  return nil
}

// createFile - FileMgr helper method. Creates a file identified by input parameter
// 'fMgr'.
func (fMgrHlpr *fileMgrHelper) createFile(
  fMgr *FileMgr,
  createDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.createFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  //  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
  fOpenCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rw-rw-rw-")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fileAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  err = fMgrHlpr.closeFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  err = fMgrHlpr.openFile(
    fMgr,
    fileAccessCfg,
    createDirectory,
    ePrefix)

  return err
}

// copyFileToFMgrCleanUp - Helper method used to perform
// clean up on Copy File Methods.
func (fMgrHlpr *fileMgrHelper) copyFileToFMgrCleanUp(
  fMgrDest *FileMgr,
  ePrefix,
  copyTypeLabel string) (err error) {

  err = nil
  ePrefixCurrMethod := "fileMgrHelper.copyFileToFMgrCleanUp "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  filePathDoesExist,
    err2 := fMgrDest.DoesThisFileExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "After attempted file copy, Non-Path Error returned by "+
      "fMgrDest.DoesThisFileExist().\n"+
      "Error='%v'\n", err2.Error())

    return err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: After attempted %v to destination file,\n"+
      "Destination file does NOT exist!\n"+
      "fMgrDest='%v'\n",
      copyTypeLabel,
      fMgrDest.absolutePathFileName)

    return err
  }

  if !fMgrDest.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "Error: Destination File was copied, but it is"+
      "classified as a Non-Regular File!\n"+
      "fMgrDest='%v'\n",
      fMgrDest.absolutePathFileName)
    return err
  }

  if fMgrDest.actualFileInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "Error: Destination File was copied using procedure %v.\n"+
      "However, the Destination File is now classified "+
      "as a Directory - NOT A FILE!\n"+
      "fMgrDest='%v'\n",
      copyTypeLabel,
      fMgrDest.absolutePathFileName)
    return err
  }

  err = nil
  return err
}

// deleteFile - Helper method used to delete
// the file identified by 'fMgr'.
//
func (fMgrHlpr *fileMgrHelper) deleteFile(
  fMgr *FileMgr, ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.deleteFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  pathFileNameDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if !pathFileNameDoesExist {
    return nil
  }

  err = fMgrHlpr.closeFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  for i := 0; i < 3; i++ {

    err2 := os.Remove(fMgr.absolutePathFileName)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "- os.Remove(fMgr.absolutePathFileName) "+
        "returned an error.\n"+
        "absolutePathFileName='%v'\nError='%v'",
        fMgr.absolutePathFileName, err2.Error())
    } else {
      err = nil
      break
    }

    time.Sleep(30 * time.Millisecond)
  }

  if err != nil {
    return err
  }

  pathFileNameDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if pathFileNameDoesExist {

    return fmt.Errorf(ePrefix+
      "Error: Attempted file deletion FAILED!. "+
      "File still exists.\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  return nil
}

func (fMgrHlpr *fileMgrHelper) emptyFileMgr(
  fMgr *FileMgr) {

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

  return
}

// flushBytesToDisk - Helper method which is designed
// to flush all buffers and write all data in memory
// to the file.
//
func (fMgrHlpr *fileMgrHelper) flushBytesToDisk(
  fMgr *FileMgr, ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.flushBytesToDisk() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var errs = make([]error, 0)

  var err2, err3 error

  if fMgr.filePtr != nil &&
    fMgr.fileBufWriter != nil {

    err3 = fMgr.fileBufWriter.Flush()

    if err3 != nil {

      err2 = fmt.Errorf(ePrefix+
        "Error returned from fMgr.fileBufWriter."+
        "Flush().\n"+
        "Error='%v'\n",
        err3.Error())

      errs = append(errs, err2)
    }
  }

  if fMgr.filePtr != nil &&
    (fMgr.fileBytesWritten > 0 ||
      fMgr.buffBytesWritten > 0) {

    err3 = fMgr.filePtr.Sync()

    if err3 != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned from fMgr.filePtr.Sync()\n"+
        "Error='%v'\n", err3.Error())
      errs = append(errs, err2)
    }
  }

  return fMgrHlpr.consolidateErrors(errs)
}

// lowLevelOpenFile - Helper method which is designed
// to open a file. Unlike similar methods, this method
// does not check for the existence of the file, does
// not try to close an existing file and will not
// create the directory path.
//
func (fMgrHlpr *fileMgrHelper) lowLevelOpenFile(
  fMgr *FileMgr,
  fileAccessCtrl FileAccessControl,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelOpenFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod + "\n"

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod + "\n"
  }

  err := fileAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+"Parameter 'fileAccessCtrl' is INVALID!\n"+
      "%v\n", err.Error())
  }

  fOpenParm, fPermParm, err :=
    fileAccessCtrl.GetFileOpenAndPermissionCodes()

  if err != nil {
    fMgr.fileAccessStatus.Empty()
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fMgr.fileAccessStatus = fileAccessCtrl.CopyOut()

  fMgr.filePtr, err =
    os.OpenFile(
      fMgr.absolutePathFileName,
      fOpenParm,
      fPermParm)

  if err != nil {

    fMgr.filePtr = nil
    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()
    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return fmt.Errorf(ePrefix+
      "Error opening file from os.OpenFile(): '%v' Error= '%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  return nil
}

// openCreateFile - Helper method designed to open a
// file. If the file does not currently exist, it
// will first be created and then opened with the
// designated access control parameters
// ('fileAccessCtrl').
//
func (fMgrHlpr *fileMgrHelper) openCreateFile(
  fMgr *FileMgr,
  fileAccessCtrl FileAccessControl,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod + "\n"
  }

  var err error
  filePathDoesExist := false

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    err = fMgrHlpr.createFile(fMgr, true, ePrefix)

    if err != nil {
      return err
    }
  }

  return fMgrHlpr.lowLevelOpenFile(fMgr, fileAccessCtrl, ePrefix)
}

// openFile - Helper method used to open the file
// specified by a File Manager.
func (fMgrHlpr *fileMgrHelper) openFile(
  fMgr *FileMgr,
  fileAccessCtrl FileAccessControl,
  createTheDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFile() "

  //originalEPrefix := ePrefix

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
    //originalEPrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  filePathDoesExist := false

  filePathDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  err = fMgrHlpr.closeFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  err = fileAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+"Parameter 'fileAccessCtrl' is INVALID!\n"+
      "%v\n", err.Error())
  }

  filePathDoesExist =
    fMgr.dMgr.doesAbsolutePathExist

  if !filePathDoesExist &&
    createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      err = fmt.Errorf(ePrefix+"%v", err.Error())
      return err
    }

  } else if !filePathDoesExist &&
    !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "Error: Directory Path DOES NOT EXIST!\n"+
      "DirPath (fMgr.dMgr)='%v'\n",
      fMgr.dMgr.GetAbsolutePath())
  }

  return fMgrHlpr.lowLevelOpenFile(
    fMgr,
    fileAccessCtrl,
    ePrefix)
}
