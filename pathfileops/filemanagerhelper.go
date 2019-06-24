package pathfileops

import (
  "bufio"
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
    return fmt.Errorf(ePrefix +
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
    err = fmt.Errorf(ePrefix +
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
    err = fmt.Errorf(ePrefix +
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
    err = fmt.Errorf(ePrefix +
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
    return fmt.Errorf(ePrefix +
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
  createDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.createFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  createTruncateAccessCtrl, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  return fMgrHlpr.openFile(
    fMgr,
    createTruncateAccessCtrl,
    createDirectory,
    ePrefix)

}

// createFileAndClose - Creates the file identified by 'fMgr'. After
// file creation the file is immediately closed.
//
// The file open operation uses create/truncate open
// codes.
//
func (fMgrHlpr *fileMgrHelper) createFileAndClose(
  fMgr *FileMgr,
  createDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.createFileAndClose() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  //  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)

  createFileAccessCfg, err := FileAccessControl{}.NewReadWriteCreateTruncateAccess()

  if err != nil {
    return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
  }

  err = fMgrHlpr.openFile(
    fMgr,
    createFileAccessCfg,
    createDirectory,
    ePrefix)

  if err != nil {
    return err
  }

  return fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
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
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  pathFileNameDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    _ = fMgrHlpr.closeFile(fMgr, ePrefix)
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

      time.Sleep(30 * time.Millisecond)
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
      "%v", err.Error())
  }

  if pathFileNameDoesExist {

    return fmt.Errorf(ePrefix+
      "\nError: Attempted file deletion FAILED!. "+
      "File still exists.\n"+
      "fMgr='%v'\n", fMgr.absolutePathFileName)
  }

  return nil
}

func (fMgrHlpr *fileMgrHelper) emptyFileMgr(
  fMgr *FileMgr) {

  if fMgr == nil {
    panic("fileMgrHelper.emptyFileMgr() - fMgr is nil pointer!\n")
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

  if fMgr == nil {
    return fmt.Errorf(ePrefix +
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
    return fmt.Errorf(ePrefix +
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
    return fmt.Errorf(ePrefix +
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
    return fmt.Errorf(ePrefix +
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

// openFile - Helper method used to open the file
// specified by a File Manager.
func (fMgrHlpr *fileMgrHelper) openFile(
  fMgr *FileMgr,
  fileAccessCtrl FileAccessControl,
  createTheDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  var filePathDoesExist bool

  filePathDoesExist,
    err := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {

    if fMgr.filePtr != nil {
      _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)
    }

    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

  if err != nil {
    return err
  }

  err = fileAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nParameter 'fileAccessCtrl' is INVALID!\n"+
      "%v\n", err.Error())
  }

  filePathDoesExist, err = fMgr.dMgr.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nNon-Path Error returned from fMgr.dMgr.DoesThisDirectoryExist()\n."+
      "fMgr.dMgr Directory Path='%v'\n"+
      "Non-Path Error='%v'\n", err.Error())
  }

  if !filePathDoesExist &&
    createTheDirectory {

    err = fMgr.dMgr.MakeDir()

    if err != nil {
      err = fmt.Errorf(ePrefix+
        "\n%v\n", err.Error())
      return err
    }

  } else if !filePathDoesExist &&
    !createTheDirectory {

    return fmt.Errorf(ePrefix+
      "\nError: Directory Path DOES NOT EXIST!\n"+
      "DirPath (fMgr.dMgr)='%v'\n",
      fMgr.dMgr.GetAbsolutePath())
  }

  return fMgrHlpr.lowLevelOpenFile(
    fMgr,
    fileAccessCtrl,
    ePrefix)
}

func (fMgrHlpr *fileMgrHelper) openFileSetup(
  fMgr *FileMgr,
  createTheDirectory bool,
  ePrefix string) error {

  ePrefixCurrMethod := "fileMgrHelper.openFileSetup() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  var filePathDoesExist, directoryPathDoesExist bool
  var err error

  filePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix)

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

  if !directoryPathDoesExist || !filePathDoesExist {

    err =
      fMgrHlpr.createFileAndClose(
        fMgr,
        createTheDirectory,
        ePrefix)

    if err != nil {
      return err
    }
  }

  return nil
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
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter 'fMgr' is a nil pointer!\n")
  }

  err = writeAccessCtrl.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by writeAccessCtrl.IsValid()\n"+
      "Error='%v'\n", err.Error())
  }

  fMgrHelpr := fileMgrHelper{}

  filePathDoesExist,
    err = fMgrHelpr.doesFileMgrPathFileExist(
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

  if !dirPathDoesExist || !filePathDoesExist {

    err =
      fMgrHelpr.createFileAndClose(
        fMgr,
        createTheDirectory,
        ePrefix)

    if err != nil {
      return err
    }
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
    invalidAccessType ||
    err2 != nil {

    err = fMgrHelpr.lowLevelCloseFile(fMgr, ePrefix)

    if err != nil {
      return err
    }

    // The file exists, just open it for read/write
    // access.
    err = fMgrHelpr.lowLevelOpenFile(fMgr, writeAccessCtrl, ePrefix)

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
