package pathfileops

import (
  "bufio"
  "errors"
  "fmt"
  "io"
  "os"
  "strings"
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
    errorPrefix = errorPrefix + "- " + ePrefixCurrMethod
  }

  if len(filePathTitle) == 0 {
    filePathTitle = "filePath"
  }

  if fileMgr == nil {
    nonPathError = fmt.Errorf(errorPrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return filePathDoesExist, nonPathError
  }

  errCode := 0

  errCode,
    _, fileMgr.absolutePathFileName =
    FileHelper{}.isStringEmptyOrBlank(fileMgr.absolutePathFileName)

  if errCode == -1 {
    fileMgr.isAbsolutePathFileNamePopulated = false
    nonPathError = fmt.Errorf(errorPrefix+
      "\nError: '%v' is an empty string!\n",
      filePathTitle)

    return filePathDoesExist, nonPathError
  }

  if errCode == -2 {
    fileMgr.isAbsolutePathFileNamePopulated = false
    nonPathError = fmt.Errorf(errorPrefix+
      "\nError: '%v' consists of blank spaces!\n",
      filePathTitle)

    return filePathDoesExist, nonPathError
  }

  if !fileMgr.isInitialized {
    nonPathError = errors.New(errorPrefix +
      "\nError: This data structure is NOT initialized.\n" +
      "fileMgr.isInitialized='false'\n")

    return filePathDoesExist, nonPathError
  }

  var err error

  err = fileMgr.dMgr.IsDirMgrValid(errorPrefix)

  if err != nil {
    nonPathError = fmt.Errorf("\nFileMgr Directory Manager INVALID!\n"+
      "\nError='%v'", err.Error())
    return filePathDoesExist, nonPathError
  }

  if preProcessCode == PreProcPathCode.PathSeparator() {

    fileMgr.absolutePathFileName = FileHelper{}.AdjustPathSlash(fileMgr.absolutePathFileName)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    fileMgr.absolutePathFileName, err = FileHelper{}.MakeAbsolutePath(fileMgr.absolutePathFileName)

    if err != nil {
      nonPathError = fmt.Errorf(errorPrefix+
        "\nFileHelper{}.MakeAbsolutePath() FAILED!\n"+
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
        "\nNon-Path error returned by os.Stat(%v)\n"+
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
// file. The version of the 'file close' operation
// first checks to verify whether the file exists
// on disk.
//
func (fMgrHlpr *fileMgrHelper) closeFile(
  fMgr *FileMgr, ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.closeFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  errs := make([]error, 0)

  var err2 error
  fileDoesExist := false

  fileDoesExist,
    err2 = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err2 != nil {

    errs = append(errs, err2)

    err2 = nil

    if fMgr.filePtr != nil {
      err2 = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
    }

    if err2 != nil {
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
    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return fMgrHlpr.consolidateErrors(errs)
  }

  // fileDoesExist == true
  return fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
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

  if fMgr == nil {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return err
  }

  if fMgrDest == nil {
    err = errors.New(ePrefix +
      "\nError: Destination File Manager (fMgrDest) is a nil pointer!\n")
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
      "\nError: This File Manager file does NOT exist!\n"+
      "(FileMgr) File Name:'%v' ", fMgr.absolutePathFileName)
    return err
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "\nError: Source file is a Non-Regular "+
      "File and cannot be copied.\n"+
      "Copy Operation Aborted.\n"+
      "Source File (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)
    return err
  }

  if fMgr.actualFileInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "\nError: Source File File Manger exists, but "+
      "it is classified as a Directory!\n"+
      "Source File (FileMgr)='%v'\n",
      fMgr.absolutePathFileName)

    return err
  }

  err2 = fMgrDest.IsFileMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError: The Destination FileMgr object is INVALID!\n"+
      "Error='%v'\n",
      err2.Error())
    return err
  }

  filePathDoesExist, err2 =
    fMgrDest.dMgr.DoesThisDirectoryExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nDestination Directory Error: Non-Path Error returned by\n"+
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
        "\nError: Attempted creation of destination directory FAILED!\n"+
        "Destination Directory (fMgrDest.dMgr)='%v'\n"+
        "Error= '%v'",
        fMgrDest.dMgr.GetAbsolutePath(), err2.Error())

      return err
    }

    filePathDoesExist,
      err2 = fMgrDest.dMgr.DoesThisDirectoryExist()

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "\nNon-Path Error returned by "+
        "fMgrDest.dMgr.DoesThisDirectoryExist()\n"+
        "fMgrDest.dMgr='%v'\n"+
        "Error='%v'\n",
        fMgrDest.dMgr.absolutePath, err2.Error())

      return err
    }

    if !filePathDoesExist {
      err = fmt.Errorf(ePrefix+
        "\nError: Attempted verification of destination directory "+
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
      "\nNon-Path Error returned by fMgrDest.DoesThisFileExist().\n"+
      "Error='%v'", err2.Error())

    return err
  }

  if filePathDoesExist &&
    fMgrDest.actualFileInfo.Mode().IsDir() {

    err = fmt.Errorf(ePrefix+
      "\nError: Destination File (fMgrDest) exists, but "+
      "it is classified as a Directory!\n"+
      "Destination File (fMgrDest)='%v'\n",
      fMgrDest.absolutePathFileName)

    return err
  }

  if filePathDoesExist &&
    !fMgrDest.actualFileInfo.Mode().IsRegular() {

    err = fmt.Errorf(ePrefix+
      "\nError: Destination file is a Non-Regular "+
      "File and cannot be the target of a copy operations.\n"+
      "Copy Operation Aborted.\n"+
      "Destination File (fMgrDest)='%v'\n",
      fMgrDest.absolutePathFileName)

    return err
  }

  if fMgr.EqualAbsPaths(fMgrDest) {
    err = fmt.Errorf(ePrefix+
      "\nError: Source and Destination File are the same!\n"+
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

  if fMgr == nil {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return fMgrDest, err
  }

  err2 := dir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError: Input parmater dir is INVALID!\n"+
      "Error='%v'", err2.Error())
    return fMgrDest, err
  }

  fMgrDest, err2 = FileMgr{}.NewFromDirMgrFileNameExt(
    dir, fMgr.fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
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

  if fMgr == nil {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return fMgrDest, err
  }

  var errCode int

  errCode,
    _,
    dstPathFileNameExt =
    FileHelper{}.isStringEmptyOrBlank(dstPathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "\nError: 'dstPathFileNameExt' is an Empty or " +
      "Zero Length String!\n")
    return fMgrDest, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "\nError: 'dstPathFileNameExt' consists entirely of " +
      "blank spaces!\n")
    return fMgrDest, err
  }

  var err2 error

  fMgrDest, err2 =
    FileMgr{}.NewFromPathFileNameExtStr(dstPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by FileMgr{}."+
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

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
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
    return fmt.Errorf(ePrefix+"\n%v\n",
      err.Error())
  }

  return nil
}

// createFile - FileMgr helper method. Creates a file
// identified by input parameter 'fMgr'.
//
// The file open operation uses create/truncate open
// codes.
//
func (fMgrHlpr *fileMgrHelper) createFile(
  fMgr *FileMgr,
  createTheDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.createFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  createTruncateAccessCtrl, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  err = createTruncateAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by createTruncateAccessCtrl.IsValid()\n"+
      "Error='%v'", err.Error())
  }

  var directoryPathDoesExist bool

  _,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  directoryPathDoesExist, err =
    fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
      "fMgr.dMgr='%v'\nError='%v'\n",
      fMgr.dMgr.absolutePath, err.Error())
  }

  if !directoryPathDoesExist && createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by fMgr.dMgr.MakeDir().\n"+
        "fMgr.dMgr='%v'\nError='%v'\n",
        fMgr.dMgr.absolutePath, err.Error())
    }

  } else if !directoryPathDoesExist && !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "\nError: File Manager Directory Path DOES NOT EXIST!\n"+
      "fMgr.dMgr='%v'\n", fMgr.dMgr.absolutePath)
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  return fMgrHlpr.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)
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

  if fMgrDest == nil {
    err = fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgrDest' is a nil pointer!\n")
    return err
  }

  filePathDoesExist,
    err2 := fMgrDest.DoesThisFileExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nAfter attempted file copy, Non-Path Error returned by "+
      "fMgrDest.DoesThisFileExist().\n"+
      "Error='%v'\n", err2.Error())

    return err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: After attempted %v to destination file,\n"+
      "Destination file does NOT exist!\n"+
      "fMgrDest='%v'\n",
      copyTypeLabel,
      fMgrDest.absolutePathFileName)

    return err
  }

  if !fMgrDest.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "\nError: Destination File was copied, but it is"+
      "classified as a Non-Regular File!\n"+
      "fMgrDest='%v'\n",
      fMgrDest.absolutePathFileName)
    return err
  }

  if fMgrDest.actualFileInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "\nError: Destination File was copied using procedure %v.\n"+
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

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  pathFileNameDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
    return err
  }

  if !pathFileNameDoesExist {
    _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
    return nil
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  /*
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

      time.Sleep(50 * time.Millisecond)
    }

    if err != nil {
      return err
    }
  */

  err = os.Remove(fMgr.absolutePathFileName)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nos.Remove(fMgr.absolutePathFileName) "+
      "returned an error.\n"+
      "absolutePathFileName='%v'\nError='%v'",
      fMgr.absolutePathFileName, err.Error())
  }

  pathFileNameDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path error returned after file deletion.\n"+
      "%v\n", err.Error())
  }

  if pathFileNameDoesExist {

    return fmt.Errorf(ePrefix+
      "\nError: Attempted file deletion FAILED!. "+
      "File still exists.\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  return nil
}

// emptyFileMgr - Helper method designed to "empty" or
// set the data fields of FileMgr to their zero or initialized
// values.
func (fMgrHlpr *fileMgrHelper) emptyFileMgr(
  fMgr *FileMgr,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.emptyFileMgr() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nInput parameter 'fMgr' is a nil pointer!\n")
  }

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

  return nil
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

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  var errs = make([]error, 0)

  var err2, err3 error

  if fMgr.filePtr != nil &&
    fMgr.fileBufWriter != nil {

    err3 = fMgr.fileBufWriter.Flush()

    if err3 != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned from fMgr.fileBufWriter."+
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
        "\nError returned from fMgr.filePtr.Sync()\n"+
        "Error='%v'\n", err3.Error())
      errs = append(errs, err2)
    }
  }

  return fMgrHlpr.consolidateErrors(errs)
}

// lowLevelCloseFile - Helper method for the File
// Manager Type.
//
// This method Closes the file identified by 'fMgr'.
//
// This method does check to verify that the file
// does exist.
//
// This method will call "fMgrHlpr.flushBytesToDisk()"
//
func (fMgrHlpr *fileMgrHelper) lowLevelCloseFile(
  fMgr *FileMgr,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelCloseFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  errs := make([]error, 0)
  var err2, err3 error

  if fMgr.filePtr == nil {

    fMgr.isFilePtrOpen = false
    fMgr.fileAccessStatus.Empty()

    fMgr.fileBufRdr = nil
    fMgr.fileBufWriter = nil
    fMgr.fileBytesWritten = 0
    fMgr.buffBytesWritten = 0

    return nil
  }

  // fMgr.filePtr != nil
  err2 = fMgrHlpr.flushBytesToDisk(fMgr, ePrefix)

  if err2 != nil {
    errs = append(errs, err2)
  }

  if fMgr.filePtr != nil {

    err3 = fMgr.filePtr.Close()

    if err3 != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by fMgr.filePtr.Close()\n"+
        "Error='%v'\n", err3.Error())
      errs = append(errs, err2)
    }
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
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  err := fileAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nParameter 'fileAccessCtrl' is INVALID!\n"+
      "%v\n", err.Error())
  }

  fOpenParm, fPermParm, err :=
    fileAccessCtrl.GetFileOpenAndPermissionCodes()

  if err != nil {
    fMgr.fileAccessStatus.Empty()
    return fmt.Errorf(ePrefix+
      "\n%v\n", err.Error())
  }

  fMgr.fileAccessStatus = fileAccessCtrl.CopyOut()

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "\n%v\n", err2.Error())
  }

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
      "\nError opening file from os.OpenFile():\n"+
      "'%v'\nError= '%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  if fMgr.filePtr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: os.OpenFile() returned a 'nil' file pointer!\n"+
      "FileMgr='%v'\n",
      fMgr.absolutePathFileName)
  }

  fMgr.isFilePtrOpen = true

  if fOpenType == FOpenType.TypeReadWrite() {

    if fMgr.fileBufWriter == nil {
      if fMgr.fileWriterBufSize > 0 {
        fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
      } else {
        fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
      }
    }

    if fMgr.fileBufRdr == nil {
      if fMgr.fileRdrBufSize > 0 {
        fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
      } else {
        fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
      }
    }
  }

  if fOpenType == FOpenType.TypeWriteOnly() {

    if fMgr.fileBufWriter == nil {
      if fMgr.fileWriterBufSize > 0 {
        fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
      } else {
        fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
      }
    }
  }

  if fOpenType == FOpenType.TypeReadOnly() {

    if fMgr.fileBufRdr == nil {
      if fMgr.fileRdrBufSize > 0 {
        fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
      } else {
        fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
      }
    }
  }

  return nil
}

// moveFile - Helper method designed to move a
// file to a new destination. The operation is
// performed in two steps. First, the source
// file is copied to a destination specified
// by parameter, 'targetFMgr'. Second, if the
// the copy operation is successful, the source
// file ('fMgr') will be deleted.
//
func (fMgrHlpr *fileMgrHelper) moveFile(
  fMgr *FileMgr,
  targetFMgr *FileMgr,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  if targetFMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'targetFMgr' is a nil pointer!\n")
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
    return fmt.Errorf(ePrefix+
      "\nError: Source File Manager 'fMgr' DOES NOT EXIST!\n"+
      "fMgr='%v'\n",
      fMgr.absolutePathFileName)
  }

  if fMgr.actualFileInfo.Mode().IsDir() {
    return fmt.Errorf(ePrefix+
      "\nError: Source File 'fMgr' is NOT A FILE!\n"+
      "IT IS A DIRECTORY!!\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "\nError: Source File Invalid. "+
      "The Source File 'fMgr' IS NOT A REGULAR FILE!\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    targetFMgr,
    PreProcPathCode.None(),
    ePrefix,
    "targetFMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    err = targetFMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by targetFMgr."+
        "dMgr.MakeDir()\n"+
        "targetFMgr.dMgr='%v'\n"+
        "Error='%v'\n",
        targetFMgr.dMgr.absolutePath,
        err.Error())
    }
  }

  if filePathDoesExist {

    err = fMgrHlpr.lowLevelCloseFile(targetFMgr, ePrefix)

    if err != nil {
      return err
    }

    err = os.Remove(targetFMgr.absolutePathFileName)

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nos.Remove(targetFMgr.absolutePathFileName) "+
        "returned an error.\n"+
        "targetFMgr='%v'\nError='%v'",
        targetFMgr.absolutePathFileName, err.Error())
    }
  }

  // Close the source file
  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(targetFMgr, ePrefix)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError closing targetFMgr\n"+
      "%v\n", err.Error())
  }

  // Now, open the source file
  srcFilePtr, err := os.Open(fMgr.absolutePathFileName)

  if err != nil {

    _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    return fmt.Errorf(ePrefix+
      "Error returned from os.Open(fMgr.absolutePathFileName).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.absolutePathFileName, err.Error())
  }

  // Next, 'Create' the destination file
  // If the destination file previously exists,
  // it will be truncated.
  outFilePtr, err := os.Create(targetFMgr.absolutePathFileName)

  if err != nil {

    _ = srcFilePtr.Close()

    return fmt.Errorf(ePrefix+
      "\nError returned from os.Create("+
      "targetFMgr.absolutePathFileName)\n"+
      "targetFMgr='%v'\nError='%v'\n",
      targetFMgr.absolutePathFileName, err.Error())
  }

  bytesToRead := fMgr.actualFileInfo.Size()

  bytesCopied, err := io.Copy(outFilePtr, srcFilePtr)

  if err != nil {
    _ = srcFilePtr.Close()
    _ = outFilePtr.Close()
    return fmt.Errorf(ePrefix+
      "\nError returned from io.Copy("+
      "outFilePtr, srcFilePtr).\n"+
      "outFile='%v'\n"+
      "srcFile='%v'\n"+
      "Error='%v'\n",
      targetFMgr.absolutePathFileName,
      fMgr.absolutePathFileName,
      err.Error())
  }

  errs := make([]error, 0, 10)

  err = outFilePtr.Sync()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError flushing buffers!\n"+
      "\nError returned from outFilePtr.Sync()\n"+
      "outFilePtr=targetFMgr='%v'\nError='%v'\n",
      targetFMgr.absolutePathFileName, err.Error()))
  }

  err = srcFilePtr.Close()

  if err != nil {
    errs = append(errs, err)
  }

  err = outFilePtr.Close()

  if err != nil {
    errs = append(errs, err)
  }

  if bytesToRead != bytesCopied {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes copied not equal bytes "+
      "in source file!\n"+
      "Bytes To Read='%v'   Bytes Copied='%v'\n"+
      "Source File 'fMgr'='%v'\n"+
      "Target File targetFMgr='%v'\n",
      bytesToRead, bytesCopied,
      fMgr.absolutePathFileName,
      targetFMgr.absolutePathFileName)
    errs = append(errs, err)
  }

  if len(errs) > 0 {
    return fMgrHlpr.consolidateErrors(errs)
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "#2 fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  errs = make([]error, 0, 10)

  if filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: Move operation failed. Source File was copied, \n"+
      "but was NOT deleted after copy operation.\n"+
      "Source File 'fMgr'='%v'", fMgr.absolutePathFileName)
    errs = append(errs, err)
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    targetFMgr,
    PreProcPathCode.None(),
    ePrefix,
    "#2 targetFMgr.absolutePathFileName")

  if err != nil {
    errs = append(errs, err)
    return fMgrHlpr.consolidateErrors(errs)
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "\nError: The move operation failed!\n"+
      "The target file DOES NOT EXIST!\n"+
      "targetFMgr='%v'\n", targetFMgr.absolutePathFileName)
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
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
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

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\n%v\n",
        err.Error())
    }

    var createTruncateAccessCtrl FileAccessControl

    createTruncateAccessCtrl, err = FileAccessControl{}.NewReadWriteCreateTruncateAccess()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\n%v\n", err.Error())
    }

    err = fMgrHlpr.lowLevelOpenFile(
      fMgr,
      createTruncateAccessCtrl,
      ePrefix)

    if err != nil {
      return err
    }

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }
  }

  return fMgrHlpr.lowLevelOpenFile(fMgr, fileAccessCtrl, ePrefix)
}

// openFile - Helper method used to open files specified
// by FileMgr.
//
// If the directory does not previously exist, it will be
// created provided that input parameter, 'createTheDirectory',
// is set equal to 'true'.
//
// If the file does not previously exist, it will be created
// provided that input parameter, 'createTheFile', is set
// equal to 'true'.
//
// If the file does previously exist, it will be closed before
// being re-opened using the input parameter, 'openFileAccessCtrl',
// as a file open access specification.
//
func (fMgrHlpr *fileMgrHelper) openFile(
  fMgr *FileMgr,
  openFileAccessCtrl FileAccessControl,
  createTheDirectory bool,
  createTheFile bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFile() "
  var filePathDoesExist, directoryPathDoesExist bool
  var err error

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  err = openFileAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nInput Parameter 'openFileAccessCtrl' is INVALID!\n"+
      "Error='%v'", err.Error())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  directoryPathDoesExist, err =
    fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
      "fMgr.dMgr='%v'\nError='%v'\n",
      fMgr.dMgr.absolutePath, err.Error())
  }

  if !directoryPathDoesExist && createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by fMgr.dMgr.MakeDir().\n"+
        "fMgr.dMgr='%v'\nError='%v'\n",
        fMgr.dMgr.absolutePath, err.Error())
    }

  } else if !directoryPathDoesExist && !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "\nError: File Manager Directory Path DOES NOT EXIST!\n"+
      "fMgr.dMgr='%v'\n", fMgr.dMgr.absolutePath)
  }

  if !filePathDoesExist && !createTheFile {

    return fmt.Errorf(ePrefix+
      "\nError: The File Manager File (fMgr) DOES NOT EXIST!\n"+
      "File (fMgr)='%v'\n", fMgr.absolutePathFileName)

  } else if !filePathDoesExist && createTheFile {

    createTruncateAccessCtrl, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

    if err != nil {
      return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
    }

    err = fMgrHlpr.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)

    if err != nil {
      return err
    }
  } // if !filePathDoesExist && !createTheFile

  // At this point, the file exists on disk. Close it!

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  // Now, open the file with the specified access configuration
  return fMgrHlpr.lowLevelOpenFile(
    fMgr,
    openFileAccessCtrl,
    ePrefix)
}

// readFileSetup - Helper method designed to provide
// standard setup for methods writing data to the file
// identified by FileMgr.
//
func (fMgrHlpr *fileMgrHelper) readFileSetup(
  fMgr *FileMgr,
  readAccessCtrl FileAccessControl,
  createTheDirectory bool,
  ePrefix string) error {

  var err, err2 error
  var filePathDoesExist, dirPathDoesExist bool
  ePrefixCurrMethod := "fileMgrHelper.readFileSetup() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  err = readAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by readAccessCtrl.IsValid()\n"+
      "Error='%v'\n", err.Error())
  }

  fNewOpenType, err := readAccessCtrl.GetFileOpenType()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by readAccessCtrl.GetFileOpenType()\n"+
      "Error='%v'\n", err.Error())
  }

  if fNewOpenType != FOpenType.TypeReadWrite() &&
    fNewOpenType != FOpenType.TypeReadOnly() {

    return fmt.Errorf(ePrefix+
      "\nError: Input Parameter readAccessCtrl is NOT an "+
      "open file 'READ' Type.\n"+
      "readAccessCtrl File Open Type='%v'\n",
      fNewOpenType.String())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  dirPathDoesExist, err =
    fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
      "fMgr.dMgr='%v'\nError='%v'\n",
      fMgr.dMgr.absolutePath, err.Error())
  }

  if !dirPathDoesExist && createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by fMgr.dMgr.MakeDir().\n"+
        "fMgr.dMgr='%v'\nError='%v'",
        fMgr.dMgr.absolutePath, err.Error())
    }

  } else if !dirPathDoesExist && !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "\nError: File Manager (fMgr) Directory Path DOES NOT EXIST!\n"+
      "Directory Path (fMgr.dMgr)='%v'\n", fMgr.dMgr.absolutePath)
  }

  if !filePathDoesExist {

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }

    createTruncateAccessCtrl, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

    if err != nil {
      return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
    }

    err = fMgrHlpr.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)

    if err != nil {
      return err
    }

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {
    if fOpenType == FOpenType.TypeReadWrite() ||
      fOpenType == FOpenType.TypeReadOnly() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    invalidAccessType ||
    err2 != nil {

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }

    // The file exists, just open it for read/write
    // access.
    err = fMgrHlpr.lowLevelOpenFile(fMgr, readAccessCtrl, ePrefix)

    if err != nil {
      return err
    }
  }

  if fMgr.fileBufRdr == nil {
    err = fmt.Errorf(ePrefix +
      "\nError: fMgr.fileBufRdr == nil\n")
    return err
  }

  return nil
}

// setFileMgrDirMgrFileName - Helper method which configures a
// a FileMgr instance based on input parameters 'dMgr' and
// 'fileNameExt'.
//
func (fMgrHlpr *fileMgrHelper) setFileMgrDirMgrFileName(
  fMgr *FileMgr,
  dMgr DirMgr,
  fileNameExt string,
  ePrefix string) (isEmpty bool, err error) {

  ePrefixCurrMethod := "fileMgrHelper.setFileMgrDirMgrFileName() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  isEmpty = true
  err = nil

  if fMgr == nil {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return isEmpty, err
  }

  err2 := dMgr.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError: Input parameter 'dMgr' is INVALID!\n"+
      "dMgr.absolutePath='%v'\nError='%v'\n",
      dMgr.absolutePath, err2.Error())
    return isEmpty, err
  }

  fh := FileHelper{}

  errCode, _, fileNameExt := fh.isStringEmptyOrBlank(fileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fileNameExt' is a Zero length string!\n")
    return isEmpty, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fileNameExt' consists entirely of blank spaces!\n")
    return isEmpty, err
  }

  adjustedFileNameExt,
    isFileNameEmpty,
    err2 :=
    fh.CleanFileNameExtStr(fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned from fh.CleanFileNameExtStr(fileNameExt).\n"+
      "fileNameExt='%v'\nError='%v'\n",
      fileNameExt, err2.Error())
    return isEmpty, err
  }

  if isFileNameEmpty {
    err = fmt.Errorf(ePrefix+
      "\nError: fileName returned from fh.CleanFileNameExtStr(fileNameExt) "+
      "is a ZERO length string!\nfileNameExt='%v'\n", fileNameExt)
    return isEmpty, err
  }

  err = fMgrHlpr.emptyFileMgr(fMgr, ePrefix)

  if err != nil {
    return isEmpty, err
  }

  fMgr.dMgr = dMgr.CopyOut()

  s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(adjustedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned from fh.GetFileNameWithoutExt(adjustedFileNameExt).\n"+
      "adjustedFileNameExt='%v'\nError='%v'\n ",
      adjustedFileNameExt, err2.Error())
    isEmpty = true
    _ = fMgrHlpr.emptyFileMgr(fMgr, ePrefix)
    return isEmpty, err
  }

  if fNameIsEmpty {
    err = fmt.Errorf(ePrefix+
      "Error: fileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt)\n"+
      "is Zero length string!\n"+
      "adjustedFileNameExt='%v'\n", adjustedFileNameExt)
    _ = fMgrHlpr.emptyFileMgr(fMgr, ePrefix)
    isEmpty = true
    return isEmpty, err
  }

  fMgr.isFileNamePopulated = true
  fMgr.fileName = s

  s, extIsEmpty, err2 := fh.GetFileExtension(adjustedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned from fh.GetFileExt(fileNameAndExt).\n"+
      "fileNameAndExt='%v'\nError='%v'\n",
      adjustedFileNameExt, err2.Error())

    isEmpty = true

    _ = fMgrHlpr.emptyFileMgr(fMgr, ePrefix)
    return isEmpty, err
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
    fh.doesPathFileExist(
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

  return isEmpty, err
}

// setFileMgrPathFileName - Helper method which configures a
// a FileMgr instance based on input parameter 'pathFileNameExt'.
//
func (fMgrHlpr *fileMgrHelper) setFileMgrPathFileName(
  fMgr *FileMgr,
  pathFileNameExt string,
  ePrefix string) (isEmpty bool, err error) {

  ePrefixCurrMethod := "fileMgrHelper.setFileMgrPathFileName() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  isEmpty = true
  err = nil

  if fMgr == nil {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
    return isEmpty, err
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'pathFileNameExt' is a zero length or empty string!\n")
    return isEmpty, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "\nError: Input parameter 'pathFileNameExt' consists entirely of blank spaces!\n")
    return isEmpty, err
  }

  adjustedPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  adjustedFileNameExt, isEmptyFileName, err2 := fh.CleanFileNameExtStr(adjustedPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt).\n"+
      "adjustedPathFileNameExt='%v'\nError='%v'\n",
      adjustedPathFileNameExt, err2.Error())
    return isEmpty, err
  }

  if isEmptyFileName {
    err = fmt.Errorf(ePrefix+
      "\nError: File Name returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt)\n"+
      "is a Zero Length String!.\n"+
      "pathFileNameExt='%v'\n",
      adjustedPathFileNameExt)
    return isEmpty, err
  }

  remainingPathStr := strings.TrimSuffix(adjustedPathFileNameExt, adjustedFileNameExt)

  var dMgr DirMgr

  errCode, _, remainingPathStr = fh.isStringEmptyOrBlank(remainingPathStr)

  if errCode < 0 {
    dMgr = DirMgr{}

  } else {

    dMgr, err2 = DirMgr{}.New(remainingPathStr)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "\nError returned from DirMgr{}.NewFromPathFileNameExtStr("+
        "remainingPathStr).\n"+
        "remainingPathStr='%v'\nError='%v'\n",
        remainingPathStr, err2.Error())
      return isEmpty, err
    }
  }

  isEmpty, err =
    fMgrHlpr.setFileMgrDirMgrFileName(fMgr, dMgr, adjustedFileNameExt, ePrefix)

  return isEmpty, err
}

// writeFileSetup - Helper method designed to provide
// standard setup for methods writing data to the file
// identified by FileMgr.
//
func (fMgrHlpr *fileMgrHelper) writeFileSetup(
  fMgr *FileMgr,
  writeAccessCtrl FileAccessControl,
  createTheDirectory bool,
  ePrefix string) error {

  var err, err2 error
  var filePathDoesExist, dirPathDoesExist bool
  ePrefixCurrMethod := "fileMgrHelper.writeFileSetup() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  err = writeAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by writeAccessCtrl.IsValid()\n"+
      "Error='%v'\n", err.Error())
  }

  fNewOpenType, err := writeAccessCtrl.GetFileOpenType()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by writeAccessCtrl.GetFileOpenType()\n"+
      "Error='%v'\n", err.Error())
  }

  if fNewOpenType != FOpenType.TypeReadWrite() &&
    fNewOpenType != FOpenType.TypeWriteOnly() {

    return fmt.Errorf(ePrefix+
      "\nError: Input Parameter writeAccessCtrl is NOT an "+
      "open file 'WRITE' Type.\n"+
      "readAccessCtrl File Open Type='%v'\n",
      fNewOpenType.String())
  }

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  dirPathDoesExist, err =
    fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
      "fMgr.dMgr='%v'\nError='%v'\n",
      fMgr.dMgr.absolutePath, err.Error())
  }

  if !dirPathDoesExist && createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by fMgr.dMgr.MakeDir().\n"+
        "fMgr.dMgr='%v'\nError='%v'",
        fMgr.dMgr.absolutePath, err.Error())
    }

  } else if !dirPathDoesExist && !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "\nError: File Manager (fMgr) Directory Path DOES NOT EXIST!\n"+
      "Directory Path (fMgr.dMgr)='%v'\n", fMgr.dMgr.absolutePath)
  }

  if !filePathDoesExist {

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }

    createTruncateAccessCtrl, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

    if err != nil {
      return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
    }

    err = fMgrHlpr.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)

    if err != nil {
      return err
    }

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }
  }

  invalidAccessType := true

  fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType()

  if err2 == nil {
    if fOpenType == fOpenType.TypeReadWrite() ||
      fOpenType == fOpenType.TypeWriteOnly() {

      invalidAccessType = false
    }
  }

  if !fMgr.isFilePtrOpen ||
    fMgr.filePtr == nil ||
    invalidAccessType ||
    err2 != nil {

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }

    // The file exists, just open it for read/write
    // access.
    err = fMgrHlpr.lowLevelOpenFile(fMgr, writeAccessCtrl, ePrefix)

    if err != nil {
      return err
    }
  }

  if fMgr.fileBufWriter == nil {
    err = fmt.Errorf(ePrefix +
      "\nError: fMgr.fileBufWriter == nil\n")
    return err
  }

  return nil
}
