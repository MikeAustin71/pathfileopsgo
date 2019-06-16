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

  ePrefix := "fileMgrHelper.doesDirPathExist() "

  filePathDoesExist = false
  nonPathError = nil

  if len(errorPrefix) > 0 {
    ePrefix = errorPrefix
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
    nonPathError = fmt.Errorf(ePrefix+
      "Error: '%v' is an empty string!", filePathTitle)
    return filePathDoesExist, nonPathError
  }

  if errCode == -2 {
    fileMgr.isAbsolutePathFileNamePopulated = false
    nonPathError = fmt.Errorf(ePrefix+
      "Error: '%v' consists of blank spaces!", filePathTitle)
    return filePathDoesExist, nonPathError
  }

  if !fileMgr.isInitialized {
    nonPathError = errors.New(ePrefix +
      "Error: This data structure is NOT initialized.\n" +
      "fileMgr.isInitialized='false'\n")
    return filePathDoesExist, nonPathError
  }

  var err error

  err = fileMgr.dMgr.IsDirMgrValid(ePrefix)

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
      nonPathError = fmt.Errorf(ePrefix+"FileHelper{}.MakeAbsolutePath() FAILED!\n"+
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
        filePathDoesExist = false
        nonPathError = nil
        _ = fileMgr.dMgr.DoesPathExist()
        _ = fileMgr.dMgr.DoesAbsolutePathExist()
        return filePathDoesExist, nonPathError
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be tested
      // up to 3-times before it is returned.
      nonPathError = fmt.Errorf(ePrefix+"Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n",
        filePathTitle, filePathTitle, fileMgr.absolutePathFileName, err.Error())
      fileMgr.actualFileInfo = FileInfoPlus{}
      filePathDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      filePathDoesExist = true
      nonPathError = nil
      fileMgr.doesAbsolutePathFileNameExist = true
      fileMgr.actualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)
      _ = fileMgr.dMgr.DoesPathExist()
      _ = fileMgr.dMgr.DoesAbsolutePathExist()

      return filePathDoesExist, nonPathError
    }

    time.Sleep(30 * time.Millisecond)
  }

  _ = fileMgr.dMgr.DoesPathExist()
  _ = fileMgr.dMgr.DoesAbsolutePathExist()

  return filePathDoesExist, nonPathError
}

// copyFileToDirSetup - Helper method used by FileMgr Copy
// to Directory routines.
//
func (fMgrHlpr *fileMgrHelper) copyFileToDirSetup(
  fMgr *FileMgr,
  dir DirMgr,
  ePrefix string) (fMgrDest FileMgr, err error) {

  err = nil
  fMgrDest = FileMgr{}
  ePrefixExtra := "fileMgrHelper.copyFileToDirSetup "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixExtra
  } else {
    ePrefix = ePrefix + "- " + ePrefixExtra
  }

  filePathDoesExist,
    err2 := fMgrHlpr.doesFileMgrPathFileExist(
    fMgr,
    PreProcPathCode.None(),
    ePrefix,
    "fMgr.absolutePathFileName")

  if err2 != nil {
    err = err2
    return fMgrDest, err
  }

  if !filePathDoesExist {
    err = fmt.Errorf(ePrefix+
      "This File Manager file DOES NOT EXIST!\n"+
      "(FileMgr) FileName='%v'\n",
      fMgr.absolutePathFileName)

    return fMgrDest, err
  }

  if !fMgr.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "Error: Source file is a Non-Regular "+
      "File and cannot be copied.\nFile='%v'\n",
      fMgr.absolutePathFileName)

    return fMgrDest, err
  }

  err2 = dir.IsDirMgrValid("")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: Input parmater dir is INVALID!\n"+
      "Error='%v'", err2.Error())
    return fMgrDest, err
  }

  fMgrDest, err2 = FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, "+
      "fMgr.fileNameExt)\n"+
      "dir.absolutePath='%v'\nfMgr.fileNameExt='%v'\nError='%v'\n",
      dir.absolutePath, fMgr.fileNameExt, err2.Error())
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  if fMgr.EqualAbsPaths(&fMgrDest) {

    err = fmt.Errorf(ePrefix+
      "Error: Source and Destination File are the same!\n"+
      "Source File='%v'\n Destination File='%v'\n",
      fMgr.absolutePathFileName, fMgrDest.absolutePathFileName)
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  err2 = fMgrDest.dMgr.MakeDir()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Atempted creation of destination directory FAILED!\n"+
      "dMgr='%v'\n"+
      "Error= '%v'\n",
      fMgrDest.dMgr.absolutePath, err2.Error())
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  filePathDoesExist,
    err2 = fMgrDest.DoesThisFileExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Non-Path Error Returned by fMgrDest.DoesThisFileExist().\n"+
      "Error='%v'\n", err2.Error())
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  if filePathDoesExist && !fMgrDest.actualFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "Error: Destination file exists and it is NOT a 'regular' file.\n"+
      "Copy operation aborted!\nDestination File='%v'\n",
      fMgrDest.absolutePathFileName)
    fMgrDest = FileMgr{}
    return fMgrDest, err
  }

  err = nil
  return fMgrDest, err
}

// copyFileToFMgrCleanUp - Helper method used to perform
// clean up on Copy File Methods.
func (fMgrHlpr *fileMgrHelper) copyFileToFMgrCleanUp(
  fMgrDest *FileMgr,
  ePrefix string) (err error) {

  err = nil
  ePrefixExtra := "fileMgrHelper.copyFileToFMgrCleanUp "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixExtra
  } else {
    ePrefix = ePrefix + "- " + ePrefixExtra
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
      "Error: After attempted file copy to destination file,\n"+
      "Destination file does NOT exist!\n"+
      "fMgrDest='%v'\n",
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

  err = nil
  return err
}
