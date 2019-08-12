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
// for the existence of a path and file name. In addition,
// this method performs validation on the 'FileMgr' instance.
//
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

  var  err2, err3 error

  err2 = fileMgr.dMgr.IsDirMgrValid(errorPrefix)

  if err2 != nil {
    nonPathError = fmt.Errorf("\nFileMgr Directory Manager INVALID!\n"+
      "\nError='%v'", err2.Error())
    return filePathDoesExist, nonPathError
  }

  if preProcessCode == PreProcPathCode.PathSeparator() {

    fileMgr.absolutePathFileName = FileHelper{}.AdjustPathSlash(fileMgr.absolutePathFileName)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    fileMgr.absolutePathFileName, err2 = FileHelper{}.MakeAbsolutePath(fileMgr.absolutePathFileName)

    if err2 != nil {
      nonPathError = fmt.Errorf(errorPrefix+
        "\nFileHelper{}.MakeAbsolutePath() FAILED!\n"+
        "%v='%v'"+
        "%v", filePathTitle, fileMgr.absolutePathFileName, err2.Error())
      return filePathDoesExist, nonPathError
    }
  }

  var info FileInfoPlus

  filePathDoesExist,
  info,
  nonPathError =
        fMgrHlpr.lowLevelDoesFileExist(
          fileMgr.absolutePathFileName,
          fileMgr.dMgr.absolutePath,
          errorPrefix,
          "fileMgr.absolutePathFileName",
          "fileMgr.dMgr.absolutePath")

  /*
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
      fileMgr.actualFileInfo, _ = FileInfoPlus{}.NewFromPathFileInfo(fileMgr.dMgr.absolutePath, info)

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

*/

  if nonPathError != nil {
    fileMgr.doesAbsolutePathFileNameExist = false
    fileMgr.actualFileInfo = FileInfoPlus{}
    filePathDoesExist = false
    fileMgr.fileAccessStatus.Empty()
    return filePathDoesExist, nonPathError
  }

  if !filePathDoesExist {
    fileMgr.doesAbsolutePathFileNameExist = false
    fileMgr.actualFileInfo = FileInfoPlus{}
    fileMgr.fileAccessStatus.Empty()
    filePathDoesExist = false
    nonPathError = nil
    _ = fileMgr.dMgr.DoesPathExist()
    _ = fileMgr.dMgr.DoesAbsolutePathExist()
    return filePathDoesExist, nonPathError
  }

  // The path really does exist!
  errs := make([]error, 0, 10)
  filePathDoesExist = true
  nonPathError = nil
  fileMgr.doesAbsolutePathFileNameExist = true

  fileMgr.actualFileInfo, err2 =
    FileInfoPlus{}.NewFromPathFileInfo(fileMgr.dMgr.absolutePath, info)

  if err2 != nil {

    err3 = fmt.Errorf(errorPrefix +
      "\nError returned by FileInfoPlus{}.NewFromPathFileInfo(fileMgr.dMgr.absolutePath, info)\n" +
      "fileMgr.dMgr.absolutePath='%v'\n" +
      "info.Name()='%v'\n" +
      "Error='%v'\n", fileMgr.dMgr.absolutePath, info.Name(), err2.Error())

    errs = append(errs, err3)

  } else {

    permCode, err2 := FilePermissionConfig{}.NewByFileMode(fileMgr.actualFileInfo.Mode())

    if err2 != nil {
      err3 = fmt.Errorf(errorPrefix +
        "\nError returned by FilePermissionConfig{}.NewByFileMode(fileMgr.actualFileInfo.Mode())\n" +
        "Error='%v'\n", err2.Error())

      errs= append(errs, err3)

    } else {

      err2 = fileMgr.fileAccessStatus.SetFilePermissionCodes(permCode)

      if err2 != nil {
        err3 = fmt.Errorf(errorPrefix +
          "\nError returned by fileMgr.fileAccessStatus.SetFilePermissionCodes(permCode)\n" +
          "Error='%v'\n", err2.Error())

        errs= append(errs, err3)
      }

    }
  }

  _ = fileMgr.dMgr.DoesPathExist()

  _ = fileMgr.dMgr.DoesAbsolutePathExist()

  if len(errs) > 0 {
    nonPathError = FileHelper{}.ConsolidateErrors(errs)
  }

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
      err2 = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")
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
  return fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")
}

// ConsolidateErrors - Receives an array of errors and converts them
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

// copyFileSetup - Helper method designed to perform setup and
// validation checks before initiating a copy operation.
//
// This method is designed to be called immediately before method
// fileMgrHelper.lowLevelCopyByIO() and fileMgrHelper.lowLevelCopyByLink().
//
func (fMgrHlpr *fileMgrHelper) copyFileSetup(
  srcFMgr *FileMgr,
  destFMgr *FileMgr,
  createDestinationDir bool,
  deleteExistingDestFile bool,
  ePrefix string,
  srcFMgrLabel string,
  destFMgrLabel string) error {

  ePrefixCurrMethod := "fileMgrHelper.copyFileSetup() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if srcFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      srcFMgrLabel)
  }

  if destFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      destFMgrLabel)
  }

  var err error
  var fMgrFileDoesExist, targetFMgrFileDoesExist bool

  fMgrFileDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    srcFMgr,
    PreProcPathCode.None(),
    ePrefix,
    srcFMgrLabel+".absolutePathFileName")

  if err != nil {
    return err
  }

  if !fMgrFileDoesExist {
    return fmt.Errorf(ePrefix+
      "\nError: Source File ('%v') DOES NOT EXIST!\n"+
      "Source File ('%v')='%v'\n",
      srcFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName)
  }

  err = fMgrHlpr.lowLevelCloseFile(srcFMgr, ePrefix, srcFMgrLabel)

  if err != nil {
    return err
  }

  if !srcFMgr.actualFileInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "\nError: %v file is NOT a 'Regular' File!\n"+
      "%v='%v'\n",
      srcFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName)
  }

  if srcFMgr.actualFileInfo.Mode().IsDir() {
    return fmt.Errorf(ePrefix+
      "\nError: %v is NOT a file! It is a DIRECTORY!\n"+
      "%v='%v'\n",
      srcFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName)
  }

  targetFMgrFileDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    destFMgr,
    PreProcPathCode.None(),
    ePrefix,
    destFMgrLabel+".absolutePathFileName")

  if err != nil {
    return err
  }

  if fMgrFileDoesExist && targetFMgrFileDoesExist {

    if os.SameFile(srcFMgr.actualFileInfo.origFileInfo,
      destFMgr.actualFileInfo.origFileInfo) ||
      strings.ToLower(srcFMgr.absolutePathFileName) ==
          strings.ToLower(destFMgr.absolutePathFileName) {
      return fmt.Errorf(ePrefix+
        "\nError: %v FileInfo is same as %v FileInfo.\n"+
        "These two files are one in the same file!\n"+
        "%v='%v'\n%v='%v'\n",
        srcFMgrLabel, destFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName,
        destFMgrLabel, destFMgr.absolutePathFileName)
    }

  } else if strings.ToLower(srcFMgr.absolutePathFileName) ==
    strings.ToLower(destFMgr.absolutePathFileName) {

    return fmt.Errorf(ePrefix+
      "\nError: %v path file name is same as %v path file name.\n"+
      "These two files are one in the same file!\n"+
      "%v='%v'\n%v='%v'\n",
      srcFMgrLabel, destFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName)
  }

  err = fMgrHlpr.lowLevelCloseFile(destFMgr, ePrefix, destFMgrLabel)

  if err != nil {
    return err
  }

  if targetFMgrFileDoesExist {

    if destFMgr.actualFileInfo.Mode().IsDir() {
      return fmt.Errorf(ePrefix+
        "\nError: %v is NOT A FILE! IT IS A DIRECTORY!!\n"+
        "%v='%v'\n",
        destFMgrLabel, destFMgrLabel, destFMgr.absolutePathFileName)
    }

    if !destFMgr.actualFileInfo.Mode().IsRegular() {

      return fmt.Errorf(ePrefix+
        "\nError: %v exists and is NOT a 'regular' file.\n"+
        "%v='%v'\n",
        destFMgrLabel, destFMgrLabel, destFMgr.absolutePathFileName)

    }

    if deleteExistingDestFile {

      err = fMgrHlpr.lowLevelDeleteFile(destFMgr, ePrefix, destFMgrLabel)

      if err != nil {
        return err
      }
    }

  } else {

    targetFMgrFileDoesExist, err = destFMgr.dMgr.DoesThisDirectoryExist()

    if err != nil {
      return fmt.Errorf(ePrefix+
        "\nError returned by %v.dMgr.DoesThisDirectoryExist()\n"+
        "%v='%v'\nError='%v'\n",
        destFMgrLabel, destFMgrLabel, destFMgr.absolutePathFileName, err.Error())
    }

    if !targetFMgrFileDoesExist && createDestinationDir {

      err = destFMgr.dMgr.MakeDir()

      if err != nil {
        return fmt.Errorf(ePrefix+
          "\nError returned by %v.dMgr.MakeDir()\n"+
          "%v.dMgr='%v'\n"+
          "Error='%v'\n",
          destFMgrLabel, destFMgrLabel, destFMgr.absolutePathFileName, err.Error())
      }

    } else if !targetFMgrFileDoesExist && !createDestinationDir {

      return fmt.Errorf(ePrefix+
        "\nError: %v directory DOES NOT EXIST!\n"+
        "%v='%v'\n",
        destFMgrLabel, destFMgrLabel, destFMgr.absolutePathFileName)
    }
  }

  return nil
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

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

  if err != nil {
    return err
  }

  return fMgrHlpr.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)
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
    _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")
    return err
  }

  if !pathFileNameDoesExist {
    _ = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")
    return nil
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelDeleteFile(fMgr, ePrefix, "fMgr")

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

// lowLevelCopyByIOBuffer - Helper method which implements
// the "Copy IO" operation. Specifically, this method calls
// 'io.CopyBuffer'.
//
// This is a low level method which performs no set up
// or validation services before initiating the "Copy By IO"
// operation.
//
// If the input parameter 'localCopyBuffer' is an array
// of bytes ([]byte) supplied by the calling function.
// The copy operation stages through the provided buffer
// (if one is required) rather than allocating a temporary
// one.
//
// If the length of 'localCopyBuffer' is zero or if
// 'localCopyBuffer' is nil, this method will return an
// error.
//
// The best practice is to call fileMgrHelper.copyFileSetup()
// or similar method first, before calling this method.
//
//   Reference:
//     https://golang.org/pkg/io/#CopyBuffer
//     https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgrHlpr *fileMgrHelper) lowLevelCopyByIOBuffer(srcFMgr *FileMgr,
  destFMgr *FileMgr,
  localCopyBuffer []byte,
  ePrefix string,
  srcFMgrLabel string,
  destFMgrLabel string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelCopyByIOBuffer() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if srcFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      srcFMgrLabel)
  }

  if destFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      destFMgrLabel)
  }

  if localCopyBuffer == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'localCopyBuffer' is nil!\n")
  }

  if len(localCopyBuffer) == 0 {
    return errors.New(ePrefix +
      "\nError: Input parameter 'localCopyBuffer' is a ZERO Length Array!\n")
  }

  var targetFMgrFileDoesExist bool

  // First, open the source file
  srcPtr, err := os.Open(srcFMgr.absolutePathFileName)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by os.Open(%v.absolutePathFileName)\n"+
      "%v.absolutePathFileName='%v'\nError='%v'\n",
      srcFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName)
  }

  destPtr, err := os.Create(destFMgr.absolutePathFileName)

  if err != nil {

    _ = srcPtr.Close()

    return fmt.Errorf(ePrefix+
      "\nError returned by os.Create(%v.absolutePathFileName)\n"+
      "%v.absolutePathFileName='%v'\nError='%v'\n",
      destFMgrLabel, destFMgrLabel,
      destFMgr.absolutePathFileName, err.Error())
  }

  bytesCopied, err := io.CopyBuffer(destPtr, srcPtr, localCopyBuffer)

  if err != nil {
    _ = srcPtr.Close()
    _ = destPtr.Close()

    return fmt.Errorf(ePrefix+
      "\nError returned by io.CopyBuffer(%v, %v)\n"+
      "%v='%v'\n%v='%v'\nError='%v'\n",
      destFMgrLabel, srcFMgrLabel,
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error())
  }

  errs := make([]error, 0, 5)

  err = destPtr.Sync()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by %v.Sync()\n"+
      "%v='%v'\nError='%v'\n",
      destFMgrLabel,
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error()))
  }

  err = srcPtr.Close()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by srcPtr.Close()\n"+
      "srcPtr=%v='%v'\nError='%v'\n",
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      err.Error()))
  }

  srcPtr = nil

  err = destPtr.Close()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by destPtr.Close()\n"+
      "destPtr=%v='%v'\nError='%v'\n",
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error()))
  }

  destPtr = nil

  if len(errs) > 0 {
    return fMgrHlpr.consolidateErrors(errs)
  }

  targetFMgrFileDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    destFMgr,
    PreProcPathCode.None(),
    ePrefix,
    destFMgrLabel+".absolutePathFileName")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nAfter Copy IO operation, %v "+
      "generated non-path error!\n"+
      "%v", destFMgrLabel, err.Error())
  }

  if !targetFMgrFileDoesExist {
    return fmt.Errorf(ePrefix+
      "\nERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
      "Destination File = %v='%v'\n",
      destFMgrLabel, destFMgr.absolutePathFileName)
  }

  srcFileSize := srcFMgr.actualFileInfo.Size()

  if bytesCopied != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes Copied does NOT equal bytes "+
      "in source file!\n"+
      "Source File Bytes='%v'   Bytes Coped='%v'\n"+
      "Source File=%v='%v'\nDestination File=%v='%v'\n",
      srcFileSize, bytesCopied,
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName)
    return err
  }

  if destFMgr.actualFileInfo.Size() != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes is source file do NOT equal bytes "+
      "in destination file!\n"+
      "Source File Bytes='%v'   Destination File Bytes='%v'\n"+
      "Source File=%v='%v'\nDestination File=%v='%v'\n",
      srcFileSize, destFMgr.actualFileInfo.Size(),
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName)
    return err
  }

  return nil
}

// lowLevelCopyByIO - Helper method which implements
// the "Copy IO" operation.
//
// This is a low level method which performs no set up
// validation before initiating the "Copy By IO" operation.
//
// If the input parameter 'localBufferSize' is greater
// than zero, the copy operation will be implemented
// using a local buffer of type []byte. The size of
// this []byte array is therefore specified by
// 'localBufferSize'.
//
// The best practice is to call fileMgrHelper.copyFileSetup()
// or similar method first, before calling this method.
//   Reference:
//     https://golang.org/pkg/io/#CopyBuffer
//     https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fMgrHlpr *fileMgrHelper) lowLevelCopyByIO(
  srcFMgr *FileMgr,
  destFMgr *FileMgr,
  localBufferSize int,
  ePrefix string,
  srcFMgrLabel string,
  destFMgrLabel string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelCopyByIO() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if srcFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      srcFMgrLabel)
  }

  if destFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      destFMgrLabel)
  }

  var targetFMgrFileDoesExist bool

  // First, open the source file
  srcPtr, err := os.Open(srcFMgr.absolutePathFileName)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by os.Open(%v.absolutePathFileName)\n"+
      "%v.absolutePathFileName='%v'\nError='%v'\n",
      srcFMgrLabel, srcFMgrLabel, srcFMgr.absolutePathFileName)
  }

  destPtr, err := os.Create(destFMgr.absolutePathFileName)

  if err != nil {

    _ = srcPtr.Close()

    return fmt.Errorf(ePrefix+
      "\nError returned by os.Create(%v.absolutePathFileName)\n"+
      "%v.absolutePathFileName='%v'\nError='%v'\n",
      destFMgrLabel, destFMgrLabel,
      destFMgr.absolutePathFileName, err.Error())
  }

  var byteBuff []byte

  if localBufferSize > 0 {
    byteBuff = make([]byte, localBufferSize)
  }

  bytesCopied, err := io.CopyBuffer(destPtr, srcPtr, byteBuff)

  if err != nil {
    byteBuff = nil
    _ = srcPtr.Close()
    _ = destPtr.Close()

    return fmt.Errorf(ePrefix+
      "\nError returned by io.CopyBuffer(%v, %v)\n"+
      "%v='%v'\n%v='%v'\nError='%v'\n",
      destFMgrLabel, srcFMgrLabel,
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error())
  }

  errs := make([]error, 0, 5)

  err = destPtr.Sync()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by %v.Sync()\n"+
      "%v='%v'\nError='%v'\n",
      destFMgrLabel,
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error()))
  }

  byteBuff = nil

  err = srcPtr.Close()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by srcPtr.Close()\n"+
      "srcPtr=%v='%v'\nError='%v'\n",
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      err.Error()))
  }

  srcPtr = nil

  err = destPtr.Close()

  if err != nil {
    errs = append(errs, fmt.Errorf(ePrefix+
      "\nError returned by destPtr.Close()\n"+
      "destPtr=%v='%v'\nError='%v'\n",
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error()))
  }

  destPtr = nil

  if len(errs) > 0 {
    return fMgrHlpr.consolidateErrors(errs)
  }

  targetFMgrFileDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    destFMgr,
    PreProcPathCode.None(),
    ePrefix,
    destFMgrLabel+".absolutePathFileName")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nAfter Copy IO operation, %v "+
      "generated non-path error!\n"+
      "%v", destFMgrLabel, err.Error())
  }

  if !targetFMgrFileDoesExist {
    return fmt.Errorf(ePrefix+
      "\nERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
      "Destination File = %v='%v'\n",
      destFMgrLabel, destFMgr.absolutePathFileName)
  }

  srcFileSize := srcFMgr.actualFileInfo.Size()

  if bytesCopied != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes Copied does NOT equal bytes "+
      "in source file!\n"+
      "Source File Bytes='%v'   Bytes Coped='%v'\n"+
      "Source File=%v='%v'\nDestination File=%v='%v'\n",
      srcFileSize, bytesCopied,
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName)
    return err
  }

  if destFMgr.actualFileInfo.Size() != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes is source file do NOT equal bytes "+
      "in destination file!\n"+
      "Source File Bytes='%v'   Destination File Bytes='%v'\n"+
      "Source File=%v='%v'\nDestination File=%v='%v'\n",
      srcFileSize, destFMgr.actualFileInfo.Size(),
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName)
    return err
  }

  return nil
}

// lowLevelCopyByLink - Helper method designed to copy a a file by
// 'link' using method os.Link(src, dst).
//
// "os.Link(src, dst)" is the only method employed to copy a designated
// file. If "os.Link(src, dst)" fails, an error is returned.
//
// By creating a 'linked' file, changing the contents of one file will be
// reflected in the second. The two linked files are 'mirrors' of each
// other.
//
// If the 'mirror' feature causes concerns, use the "Copy By IO' technique
// employed by fileMgrHelper.lowLevelCopyByIO().
//
// Be advised that this method does not parameter validation or setup
// services prior to initiating the 'Copy By Link' operation. It would be
// prudent to call 'fileMgrHelper.copyFileSetup()' first, before calling
// this method.
//
//  Reference
//   https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//   https://golang.org/pkg/os/#Link
//

func (fMgrHlpr *fileMgrHelper) lowLevelCopyByLink(
  srcFMgr *FileMgr,
  destFMgr *FileMgr,
  ePrefix string,
  srcFMgrLabel string,
  destFMgrLabel string) error {
  ePrefixCurrMethod := "fileMgrHelper.lowLevelCloseFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if srcFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      srcFMgrLabel)
  }

  if destFMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      destFMgr)
  }

  err := os.Link(
    srcFMgr.absolutePathFileName,
    destFMgr.absolutePathFileName)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by os.Link("+
      "%v.absolutePathFileName, %v.absolutePathFileName)\n"+
      "%v.absolutePathFileName=%v\n"+
      "%v.absolutePathFileName=%v\n"+
      "Error='%v'",
      srcFMgrLabel, destFMgrLabel,
      srcFMgrLabel, srcFMgr.absolutePathFileName,
      destFMgrLabel, destFMgr.absolutePathFileName,
      err.Error())
  }

  destinationFileDoesExist,
  err := fMgrHlpr.doesFileMgrPathFileExist(
    destFMgr,
    PreProcPathCode.None(),
    ePrefix,
    destFMgrLabel+".absolutePathFileName")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nAfter Copy By Link operation, %v "+
      "generated non-path error!\n"+
      "%v", destFMgrLabel, err.Error())
  }

  if !destinationFileDoesExist {
    return fmt.Errorf(ePrefix+
      "\nERROR: After Copy By Link operation, the destination file DOES NOT EXIST!\n"+
      "Destination File = %v='%v'\n",
      destFMgrLabel, destFMgr.absolutePathFileName)
  }

  return nil
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
  ePrefix string,
  fMgrLabel string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelCloseFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if fMgr == nil {
    return fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is a nil pointer!\n",
      fMgrLabel)
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
        "\nError returned by %v.filePtr.Close()\n"+
        "Error='%v'\n", fMgrLabel, err3.Error())
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

// lowLevelDeleteFile - Helper method which is designed
// to delete the file identified by input parameter, 'fMgr'.
//
// This method performs no validation on 'fMgr' and does
// not verify existence of the 'fMgr' file.
//
func (fMgrHlpr *fileMgrHelper) lowLevelDeleteFile(
  targetFMgr *FileMgr,
  ePrefix string,
  targetFileLabel string) error {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelDeleteFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod

  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if targetFMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'targetFMgr' is a nil pointer!\n")
  }

  var err error

  for i := 0; i < 3; i++ {

    err2 := os.Remove(targetFMgr.absolutePathFileName)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "\nError: os.Remove(%v.absolutePathFileName) "+
        "returned an error.\n"+
        "%v='%v'\nError='%v'",
        targetFileLabel,
        targetFileLabel,
        targetFMgr.absolutePathFileName,
        err2.Error())

    } else {

      return nil
    }

    time.Sleep(50 * time.Millisecond)
  }

  return err
}


// lowLevelDoesFileExist - Tests input parameter
// 'pathFileName' to determine if the file exists
// on disk.
//
// The method makes three calls to os.Stat(pathFileName)
// to ensure that non-path errors are real and not due
// to operating system timing issues.
//
func (fMgrHlpr *fileMgrHelper) lowLevelDoesFileExist(
  pathFileName,
  directoryPath,
  ePrefix,
  pathFileNameLabel,
  directoryPathLabel string) (pathFileDoesExist bool,
                             fInfoPlus FileInfoPlus,
                             err error) {

  ePrefixCurrMethod := "fileMgrHelper.lowLevelDoesFileExist() "

  pathFileDoesExist = false
  fInfoPlus = FileInfoPlus{}
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if len(pathFileNameLabel) == 0 {
    pathFileNameLabel = "pathFileName"
  }

  if len(directoryPathLabel) == 0 {
    directoryPathLabel = "dirPath"
  }

  errCode := 0
  fh := FileHelper{}

  errCode,
  _,
  pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode < 0 {
    err = fmt.Errorf(ePrefix +
      "\nERROR: Input parameter %v is an EMPTY STRING!\n",
      pathFileNameLabel)

    return pathFileDoesExist, fInfoPlus, err
  }

  errCode,
  _,
  directoryPath = fh.isStringEmptyOrBlank(directoryPath)

  var err2 error
  var info os.FileInfo

  for i := 0; i < 3; i++ {

    pathFileDoesExist = false
    fInfoPlus = FileInfoPlus{}
    err = nil

    info, err2 = os.Stat(pathFileName)

    if err2 != nil {

      if os.IsNotExist(err2) {

        pathFileDoesExist = false
        fInfoPlus = FileInfoPlus{}
        err = nil
        return pathFileDoesExist, fInfoPlus, err
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be test
      // up to 3-times before it is returned.
      err = fmt.Errorf(ePrefix+"Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n\n",
        pathFileNameLabel, pathFileNameLabel, err2.Error())
      fInfoPlus = FileInfoPlus{}
      pathFileDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      pathFileDoesExist = true
      err = nil

      if errCode < 0 {
        // If the directoryPath is an empty string,
        // only record the the os.FileInfo data, not the path.
        //
        fInfoPlus = FileInfoPlus{}.NewFromFileInfo(info)

      } else {

        // directoryPath is a string with length greater than zero,
        // record the os.FileInfo data AND the path.
        fInfoPlus, err2 = FileInfoPlus{}.NewFromPathFileInfo(directoryPath, info)

        if err2 != nil {
          err = fmt.Errorf(ePrefix +
            "\nError returned by FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)\n" +
            "Error='%v'\n", err2.Error())
          fInfoPlus = FileInfoPlus{}
        }
      }

      return pathFileDoesExist, fInfoPlus, err
    }

    time.Sleep(30 * time.Millisecond)
  }

  return pathFileDoesExist, fInfoPlus, err
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

  ePrefixCurrMethod := "fileMgrHelper.moveFile() "

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
  var sourceFilePathDoesExist, targetFilePathDoesExist bool

  sourceFilePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

  if err != nil {
    return err
  }

  if !sourceFilePathDoesExist {
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

  targetFilePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    targetFMgr,
    PreProcPathCode.None(),
    ePrefix,
    "targetFMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(targetFMgr, ePrefix, "targetFMgr")

  if err != nil {
    return err
  }

  if !targetFilePathDoesExist {

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

  } else {
    // targetFilePathDoesExist == true
    err = fMgrHlpr.lowLevelDeleteFile(targetFMgr, ePrefix, "targetFMgr")

    if err != nil {
      return err
    }

  }

  // Close the source file
  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

  if err != nil {
    return err
  }

  err = fMgrHlpr.lowLevelCloseFile(targetFMgr, ePrefix, "targetFMgr")

  if err != nil {
    return err
  }

  // Now, open the source file
  srcFilePtr, err := os.Open(fMgr.absolutePathFileName)

  if err != nil {

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

  err = fMgrHlpr.lowLevelDeleteFile(fMgr, ePrefix, "fMgr")

  if err != nil {
    return err
  }

  sourceFilePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "#2 fMgr.absolutePathFileName")

  if err != nil {
    return err
  }

  errs = make([]error, 0, 10)

  if sourceFilePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: Move operation failed. Source File was copied, \n"+
      "but was NOT deleted after copy operation.\n"+
      "Source File 'fMgr'='%v'", fMgr.absolutePathFileName)
    errs = append(errs, err)
  }

  targetFilePathDoesExist,
    err = fMgrHlpr.doesFileMgrPathFileExist(
    targetFMgr,
    PreProcPathCode.None(),
    ePrefix,
    "#2 targetFMgr.absolutePathFileName")

  if err != nil {
    errs = append(errs, err)
    return fMgrHlpr.consolidateErrors(errs)
  }

  if !targetFilePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: The move operation failed!\n"+
      "The target file DOES NOT EXIST!\n"+
      "targetFMgr='%v'\n", targetFMgr.absolutePathFileName)
    errs = append(errs, err)
  }

  return fMgrHlpr.consolidateErrors(errs)
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

  err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err2 = fMgr.actualFileInfo.SetDirectoryPath(dMgr.absolutePath)

    if err2 != nil {
      isEmpty = true
      err = fmt.Errorf(ePrefix +
        "\nError returned by fMgr.actualFileInfo.SetDirectoryPath(dMgr.absolutePath)\n" +
        "dMgr.absolutePath='%v'\n" +
        "%v", dMgr.absolutePath, err2.Error())

      _ = fMgrHlpr.emptyFileMgr(fMgr, ePrefix)
      return isEmpty, err
    }

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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

    err = fMgrHlpr.lowLevelCloseFile(fMgr, ePrefix, "fMgr")

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
