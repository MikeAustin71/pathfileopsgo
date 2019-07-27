package pathfileops

import (
  "errors"
  "fmt"
  "io"
  "math"
  "os"
  "path"
  fp "path/filepath"
  "strings"
  "time"
)

/*
   'filehelper.go' - Contains type 'FileHelper' and related data structures.

    The 'FileHelper' type provides methods used in managing files and
    directories.

   The Source Repository for this source code file is :
     https://github.com/MikeAustin71/pathfilego.git

   'FileHelper' is a dependency of 'DirMgr' and 'FileMgr'. 'FileMgr'
   is located in source file 'filemanager.go'. 'DirMgr' is located
   in 'dirmanager.go'.


*/

// FileHelper - The methods associated with this type provide
// generalized file creation, management and maintenance utilities.
//
// 'FileHelper' is a dependency for types 'DirMgr' and 'FileMgr'.
type FileHelper struct {
  Input  string
  Output string
}

// AddPathSeparatorToEndOfPathStr - Receives a path string as an input
// parameter. If the last character of the path string is not a path
// separator, this method will add a path separator to the end of that
// path string and return it to the calling method.
//
func (fh FileHelper) AddPathSeparatorToEndOfPathStr(pathStr string) (string, error) {

  ePrefix := "FileHelper.AddPathSeparatorToEndOfPathStr() "

  errCode := 0
  lStr := 0
  errCode, lStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    return "", errors.New(ePrefix + "Error: Input parameter 'pathStr' is an empty string!")
  }

  if errCode == -2 {
    return "", errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!")
  }

  if pathStr[lStr-1] == os.PathSeparator {
    return pathStr, nil
  }

  var newPathStr string

  if pathStr[lStr-1] == '/' && '/' != os.PathSeparator {
    newPathStr = pathStr[0 : lStr-1]
    newPathStr += string(os.PathSeparator)
    return newPathStr, nil
  }

  if pathStr[lStr-1] == '\\' && '\\' != os.PathSeparator {
    newPathStr = pathStr[0 : lStr-1]
    newPathStr += string(os.PathSeparator)
    return newPathStr, nil
  }

  newPathStr = pathStr + string(os.PathSeparator)

  return newPathStr, nil
}

// AdjustPathSlash will standardize path
// separators according to operating system
func (fh FileHelper) AdjustPathSlash(path string) string {
  errCode := 0

  errCode, _, path = fh.isStringEmptyOrBlank(path)

  if errCode == -1 {
    return ""
  }

  if errCode == -2 {
    return ""
  }

  if os.PathSeparator != '\\' {
    return strings.ReplaceAll(path, "\\", string(os.PathSeparator))
  }

  if os.PathSeparator != '/' {
    return strings.ReplaceAll(path, "/", string(os.PathSeparator))
  }

  return fp.FromSlash(path)
}

// AreSameFile - Compares two paths or path/file names to determine if they are
// the same and equivalent.
//
// An error will be triggered if one or both of the input parameters, 'pathFile1'
// and 'pathFile2' are empty/blank strings.
//
// If the path file input parameters identify the same file, this method returns
// 'true'.
//
// The two input parameters 'pathFile1' and 'pathFile2' will be converted to their
// absolute paths before comparisons are applied.
//
func (fh FileHelper) AreSameFile(pathFile1, pathFile2 string) (bool, error) {

  ePrefix := "FileHelper.AreSameFile() "

  var pathFile1DoesExist, pathFile2DoesExist bool
  var fInfoPathFile1, fInfoPathFile2 FileInfoPlus
  var err error

  pathFile1,
    pathFile1DoesExist,
    fInfoPathFile1,
    err = fh.doesPathFileExist(pathFile1,
    PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
    ePrefix,
    "pathFile1")

  if err != nil {
    return false, err
  }

  pathFile2,
    pathFile2DoesExist,
    fInfoPathFile2,
    err = fh.doesPathFileExist(pathFile2,
    PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
    ePrefix,
    "pathFile2")

  if err != nil {
    return false, err
  }

  if pathFile1DoesExist && pathFile2DoesExist {

    if os.SameFile(fInfoPathFile1.GetOriginalFileInfo(), fInfoPathFile2.GetOriginalFileInfo()) {
      // pathFile1 and pathFile2 are the same
      // path and file name.

      return true, nil

    }

    return false, nil
  }

  pathFile1 = strings.ToLower(pathFile1)
  pathFile2 = strings.ToLower(pathFile2)

  if pathFile1 == pathFile2 {
    return true, nil
  }

  return false, nil
}

// ChangeFileMode changes the file mode of an existing file designated by input
// parameter 'pathFileName'. The os.FileMode value to which the mode will be changed
// is extracted from input parameter 'filePermission'.
//
// If the file does Not exist, an error is triggered.
//
// If the method succeeds and the file's mode is changed, an error value of 'nil' is
// returned.
//
func (fh FileHelper) ChangeFileMode(pathFileName string, filePermission FilePermissionConfig) error {
  ePrefix := "FileHelper.ChangeFileMode() "
  var err error
  var filePathDoesExist bool

  pathFileName,
    filePathDoesExist,
    _,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf("ERROR: 'pathFileName' DOES NOT EXIST!\n"+
      "pathFileName='%v'\n", pathFileName)
  }

  err = filePermission.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "ERROR: Input parameter 'filePermission' is INVALID!\n"+
      "Error='%v'\n", err.Error())
  }

  newOsFileMode, err := filePermission.GetFileMode()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by filePermission.GetFileMode().\nError='%v'\n",
      err.Error())
  }

  err = os.Chmod(pathFileName, newOsFileMode)

  if err != nil {
    changeModeTxt, _ := filePermission.GetPermissionTextCode()
    changeModeValue := filePermission.GetPermissionFileModeValueText()
    return fmt.Errorf(ePrefix+
      "Error returned by os.Chmod(pathFileName, newOsFileMode).\n"+
      "pathFileName='%v'\nnewOsFileMode Text='%v'\nnewOsFileModeValue='%v'\nError='%v'",
      pathFileName, changeModeTxt, changeModeValue, err.Error())
  }

  return nil
}

// ChangeFileTimes - is a wrapper for os.Chtimes(). This method will set new access and
// modification times for a designated file.
//
// If the path and file name do not exist, this method will return an error.
//
// If successful, the access and modification times for the target file will be changed to
// to those specified by input parameters, 'newAccessTime' and 'newModTime'.
//
func (fh FileHelper) ChangeFileTimes(pathFileName string, newAccessTime, newModTime time.Time) error {

  ePrefix := "FileHelper.ChangeFileTimes() "
  var err error
  var filePathDoesExist bool
  var fInfo FileInfoPlus

  pathFileName,
    filePathDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return err
  }

  if !filePathDoesExist {
    return fmt.Errorf(ePrefix+
      "ERROR: 'pathFileName' DOES NOT EXIST!\n"+
      "pathFileName='%v'\n", pathFileName)
  }

  if fInfo.IsDir() {
    return fmt.Errorf(ePrefix+
      "ERROR: 'pathFileName' is a directory path and NOT a file!\n"+
      "pathFileName='%v'\n", pathFileName)
  }

  tDefault := time.Time{}

  if newAccessTime == tDefault {

    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'newAccessTime' is INVALID!\nnewAccessTime='%v'",
      newAccessTime)

  }

  if newModTime == tDefault {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'newModTime' is INVALID!\nnewModTime='%v'",
      newModTime)
  }

  err = os.Chtimes(pathFileName, newAccessTime, newModTime)

  if err != nil {

    return fmt.Errorf(ePrefix+
      "ERROR returned by os.Chtimes(pathFileName,newAccessTime, newModTime)\n"+
      "newAccessTime='%v'\nnewModTime='%v'\nError='%v'\n",
      newAccessTime.Format("2006-01-02 15:04:05.000000000 -0700 MST"),
      newModTime.Format("2006-01-02 15:04:05.000000000 -0700 MST"),
      err.Error())
  }

  return nil
}

// ChangeWorkingDir - Changes the current working directory to the
// named directory passed in input parameter, 'dirPath'. If there
// is an error, it will be of type *PathError.
func (fh FileHelper) ChangeWorkingDir(dirPath string) error {

  ePrefix := "FileHelper.ChangeWorkingDir() "
  errCode := 0

  errCode, _, dirPath = fh.isStringEmptyOrBlank(dirPath)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'dirPath' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'dirPath' consists of blank spaces!")
  }

  err := os.Chdir(dirPath)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error returned by os.Chdir(dirPath). "+
      "dirPath='%v' Error='%v'", dirPath, err)
  }

  return nil
}

// CleanDirStr - Cleans and formats a directory string.
//
//   Examples:
//
//     dirName = '/someDir/xt_dirmgr_01_test.go' returns "./someDir"
//
//     dirName = '../dir1/dir2/fileName.ext'     returns "../dir1/dir2"
//
//     dirName = 'fileName.ext'                  returns "" isEmpty = true
//
//     dirName = 'xt_dirmgr_01_test.go'          returns "" isEmpty = true
//
//     dirName = '../dir1/dir2/'                 returns "../dir1/dir2"
//
//     dirName = '../../../dir1/'                returns '../../../dir1'
//
//     dirName = '../dir1/dir2/filename.ext'     returns "../dir1/dir2"
//
//     dirName = '../dir1/dir2/.git'             returns "../dir1/dir2"
//
//     dirName = 'somevalidcharacters'           returns "./somevalidchracters"
//
func (fh FileHelper) CleanDirStr(dirNameStr string) (returnedDirName string, isEmpty bool, err error) {

  ePrefix := "FileHelper.CleanDirStr() "
  returnedDirName = ""
  isEmpty = true
  err = nil
  var pathDoesExist bool
  var fInfo FileInfoPlus
  var adjustedDirName string

  adjustedDirName,
    pathDoesExist,
    fInfo,
    err = fh.doesPathFileExist(dirNameStr,
    PreProcPathCode.PathSeparator(), // Convert to os Path Separators
    ePrefix,
    "dirNameStr")

  if err != nil {
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  lAdjustedDirName := len(adjustedDirName)

  osPathSepStr := string(os.PathSeparator)

  if strings.Contains(adjustedDirName, osPathSepStr+osPathSepStr) {
    returnedDirName = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error: Invalid Directory string.\n"+
      "Directory string contains invalid Path Separators.\n"+
      "adjustedDirName='%v'\n",
      adjustedDirName)
    return returnedDirName, isEmpty, err
  }

  if strings.Contains(adjustedDirName, "...") {
    returnedDirName = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error: Invalid Directory string.\n"+
      "Directory string contains invalid dots.\n"+
      "adjustedDirName='%v'\n",
      adjustedDirName)
    return returnedDirName, isEmpty, err
  }

  absPath, err := fh.MakeAbsolutePath(adjustedDirName)

  if err != nil {
    err = fmt.Errorf(ePrefix+"Error occurred while convert path "+
      "to absolute path!\n"+
      "dirPath='%v'\nError='%v'\n",
      adjustedDirName, err.Error())
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  volName := fp.VolumeName(absPath)

  volumeIdx := strings.Index(adjustedDirName, volName)

  if strings.ToLower(volName) == strings.ToLower(adjustedDirName) {

    returnedDirName = adjustedDirName

    if len(returnedDirName) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return returnedDirName, isEmpty, err
  }

  isAbsPath := false

  if strings.ToLower(absPath) == strings.ToLower(adjustedDirName) {
    isAbsPath = true
  }

  dotPlusOsPathSepStr := "." + osPathSepStr

  if pathDoesExist {
    // The path exists

    if fInfo.IsDir() {

      // The path exists and it is a directory
      returnedDirName = adjustedDirName

    } else {
      // The path exists but it is
      // a File Name and NOT a directory name.
      adjustedDirName = adjustedDirName[0 : lAdjustedDirName-len(fInfo.Name())]
      lAdjustedDirName = len(adjustedDirName)

      if lAdjustedDirName < 1 {
        returnedDirName = ""
      } else {
        returnedDirName = adjustedDirName
      }

    }

    lAdjustedDirName = len(returnedDirName)

    if lAdjustedDirName == 0 {
      isEmpty = true
      err = nil
      return returnedDirName, isEmpty, err
    }

    if returnedDirName[lAdjustedDirName-1] == os.PathSeparator {
      // Last character in path is a Path Separator

      if lAdjustedDirName >= 2 &&
        returnedDirName[lAdjustedDirName-2] == '.' {
        // The char immediately preceding the Path Separator
        // in the last char position is a dot.
        // Keep the path separator
        returnedDirName = adjustedDirName
      } else {
        // The char immediately preceding the Path Separator in
        // the last char position is NOT a dot. Delete the
        // trailing path separator
        returnedDirName = returnedDirName[0 : lAdjustedDirName-1]
      }
    }

    lAdjustedDirName = len(returnedDirName)

    if returnedDirName[0] != '.' &&
      returnedDirName[0] != os.PathSeparator &&
      volumeIdx == -1 &&
      !isAbsPath {

      // First letter in path is a character
      // Add a leading dot
      returnedDirName = dotPlusOsPathSepStr + returnedDirName

    } else if returnedDirName[0] == os.PathSeparator &&
      volumeIdx == -1 &&
      !isAbsPath {

      returnedDirName = "." + returnedDirName
    }

    if len(returnedDirName) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return returnedDirName, isEmpty, err

  } // End of pathDoesExist == true

  // The Path DOES NOT EXIST ON DISK
  firstCharIdx, lastCharIdx, err2 :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh."+
      "GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName).\n"+
      "adjustedDirName='%v'\nError='%v'\n",
      adjustedDirName, err2.Error())
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  interiorDotPathIdx := strings.LastIndex(adjustedDirName, "."+string(os.PathSeparator))

  if interiorDotPathIdx > firstCharIdx {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. Invalid interior relative path detected!\n"+
      "adjustedDirName='%v'\n",
      adjustedDirName)
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf("Error returned by fh."+
      "GetPathSeparatorIndexesInPathStr(adjustedDirName).\n"+
      "adjusteDirName='%v'\nError='%v'\n",
      adjustedDirName, err2.Error())
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  lSlashIdxs := len(slashIdxs)

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf("Error returned by fh."+
      "GetDotSeparatorIndexesInPathStr(adjustedDirName).\n"+
      "adjustedDirName='%v'\nError='%v'\n",
      adjustedDirName, err2.Error())
    returnedDirName = ""
    isEmpty = true
    return returnedDirName, isEmpty, err
  }

  lDotIdxs := len(dotIdxs)

  if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // Option # 1
    // There are no valid characters in the string
    returnedDirName = ""

  } else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 2
    // String consists only of eligible alphanumeric characters
    // "sometextstring"
    returnedDirName = dotPlusOsPathSepStr + adjustedDirName

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 3
    // There no dots no characters but string does contain
    // slashes

    if lSlashIdxs > 1 {
      returnedDirName = ""
    } else {
      // lSlashIdxs must be '1'
      returnedDirName = dotPlusOsPathSepStr
    }

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
    // Option # 4
    // strings contains slashes and characters but no dots.

    returnedDirName = adjustedDirName

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // Option # 5
    // dots only. Must be "." or ".."
    // Add trailing path separator
    returnedDirName = adjustedDirName + osPathSepStr

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 6
    // Dots and characters only. No slashes.
    // Maybe 'fileName.ext'
    returnedDirName = ""

  } else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 7
    // Dots and slashes, but no characters.
    returnedDirName = adjustedDirName

  } else {
    // Option # 8
    // MUST BE lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
    // Has dots, slashes and characters

    returnedDirName = adjustedDirName

    // If there is a dot after the last path separator
    // and there is a character following the dot.
    // Therefore, this is a filename extension, NOT a directory
    if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
      dotIdxs[lDotIdxs-1] < lastCharIdx {

      returnedDirName = adjustedDirName[0:slashIdxs[lSlashIdxs-1]]
    }
  }

  lAdjustedDirName = len(returnedDirName)

  if lAdjustedDirName == 0 {
    isEmpty = true
    err = nil
    return returnedDirName, isEmpty, err
  }

  if returnedDirName[lAdjustedDirName-1] == os.PathSeparator {
    // Last character in path is a Path Separator
    if lAdjustedDirName >= 2 &&
      returnedDirName[lAdjustedDirName-2] == '.' {
      // The char immediately preceding the Path Separator
      // is a dot. Keep the path separator
      // do nothing
      lAdjustedDirName = len(returnedDirName)
    } else {
      returnedDirName = returnedDirName[0 : lAdjustedDirName-1]
    }
  }

  lAdjustedDirName = len(returnedDirName)

  if returnedDirName[0] != '.' &&
    returnedDirName[0] != os.PathSeparator &&
    volumeIdx == -1 &&
    !isAbsPath {

    // First letter in path is a character
    // Add a leading dot
    returnedDirName = dotPlusOsPathSepStr + returnedDirName

  } else if returnedDirName[0] == os.PathSeparator &&
    volumeIdx == -1 &&
    !isAbsPath {

    returnedDirName = "." + returnedDirName
  }

  if len(returnedDirName) == 0 {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil
  return returnedDirName, isEmpty, err
}

// CleanFileNameExtStr - Cleans up a file name extension string.
//
//   Example:
//     fileNameExt = '../dir1/dir2/fileName.ext'
//                   returns "fileName.ext" and isEmpty=false
//
//     fileNameExt = 'fileName.ext"
//                    returns "fileName.ext" and isEmpty=false
//
//     fileNameExt = "../filesfortest/newfilesfortest/" (actually exists on disk)
//                   returns "" and isEmpty=true and error = nil
//
//     fileNameExt = '../dir1/dir2/'
//                    returns "" and isEmpty=true
//
//     fileNameExt = '../filesfortest/newfilesfortest/newerFileForTest_01'
//                   returns "newerFileForTest_01" and isEmpty=false
//
//     fileNameExt = '../filesfortest/newfilesfortest/.gitignore'
//                   returns ".gitignore" and isEmpty=false
//
func (fh FileHelper) CleanFileNameExtStr(
  fileNameExtStr string) (returnedFileNameExt string, isEmpty bool, err error) {

  ePrefix := "FileHelper.CleanFileNameExtStr() "
  returnedFileNameExt = ""
  isEmpty = true
  err = nil

  var pathDoesExist bool
  var fInfo FileInfoPlus
  var adjustedFileNameExt string

  adjustedFileNameExt,
    pathDoesExist,
    fInfo,
    err = fh.doesPathFileExist(fileNameExtStr,
    PreProcPathCode.PathSeparator(), // Convert to os Path Separators
    ePrefix,
    "fileNameExtStr")

  if err != nil {
    returnedFileNameExt = ""
    isEmpty = true
    return returnedFileNameExt, isEmpty, err
  }

  osPathSepStr := string(os.PathSeparator)

  if strings.Contains(adjustedFileNameExt, osPathSepStr+osPathSepStr) {
    returnedFileNameExt = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error: Invalid Directory string.\n"+
      "Directory string contains invalid Path Separators.\n"+
      "adjustedFileNameExt='%v' ",
      adjustedFileNameExt)
    return returnedFileNameExt, isEmpty, err
  }

  if strings.Contains(adjustedFileNameExt, "...") {
    err = fmt.Errorf(ePrefix+"Error: Invalid Directory string. Contains invalid dots.\n"+
      "adjustedFileNameExt='%v'\n", adjustedFileNameExt)
    return
  }

  // Find out if the file name extension path
  // actually exists.

  if pathDoesExist {
    // The path exists

    if fInfo.IsDir() {
      // The path exists and it is a directory.
      // There is no File Name present.
      returnedFileNameExt = ""
      isEmpty = true
      err = nil
      return

    } else {
      // The path exists and it is a valid
      // file name.
      returnedFileNameExt = fInfo.Name()
      isEmpty = false

      err = nil
      return
    }
  } // End of if pathDoesExist

  firstCharIdx, lastCharIdx, err :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedFileNameExt)

  if firstCharIdx == -1 || lastCharIdx == -1 {
    err = fmt.Errorf(ePrefix+"File Name Extension string contains no "+
      "valid file name characters!\n"+
      "adjustedFileNameExt='%v'\n",
      adjustedFileNameExt)
    return
  }

  // The file name extension path does not exist

  interiorDotPathIdx := strings.LastIndex(adjustedFileNameExt, "."+string(os.PathSeparator))

  if interiorDotPathIdx > firstCharIdx {
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH. "+
      "Invalid interior relative path detected!\n"+
      "adjustedFileNameExt='%v'\n", adjustedFileNameExt)
    return
  }

  slashIdxs, err := fh.GetPathSeparatorIndexesInPathStr(adjustedFileNameExt)

  if err != nil {
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetPathSeparatorIndexesInPathStr"+
      "(adjustedFileNameExt).\n"+
      "adustedFileNameExt='%v'\nError='%v'\n",
      adjustedFileNameExt, err.Error())
    return
  }

  lSlashIdxs := len(slashIdxs)

  if lSlashIdxs == 0 {
    returnedFileNameExt = adjustedFileNameExt
    isEmpty = false
    err = nil
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(adjustedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf("Error returned by fh."+
      "GetDotSeparatorIndexesInPathStr(adjustedDirName).\n"+
      "adjustedFileNameExt='%v'\nError='%v'\n",
      adjustedFileNameExt, err2.Error())
    returnedFileNameExt = ""
    isEmpty = true
    return returnedFileNameExt, isEmpty, err
  }

  lDotIdxs := len(dotIdxs)

  // Option # 1
  if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // There are no valid characters in the string
    returnedFileNameExt = ""

  } else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 2
    // String consists only of eligible alphanumeric characters
    // "sometextstring"
    returnedFileNameExt = adjustedFileNameExt

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 3
    // There no dots no characters but string does contain
    // slashes
    returnedFileNameExt = ""

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
    // Option # 4
    // strings contains slashes and characters but no dots.

    if lastCharIdx < slashIdxs[lSlashIdxs-1] {
      // Example: ../dir1/dir2/
      returnedFileNameExt = ""
    } else {
      // Must be  lastCharIdx > slashIdxs[lSlashIdxs-1]
      // Example: ../dir1/dir2/filename
      returnedFileNameExt = adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]
    }

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // Option # 5
    // dots only. Must be "." or ".."
    // The is no file name
    returnedFileNameExt = ""

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 6
    // Dots and characters only. No slashes.
    // Maybe 'fileName.ext'
    returnedFileNameExt = adjustedFileNameExt

  } else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 7
    // Dots and slashes, but no characters.
    returnedFileNameExt = ""

  } else {
    // Option # 8
    // MUST BE lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
    // We have dots, slashes and characters

    if lastCharIdx > slashIdxs[lSlashIdxs-1] {
      // The last char comes after the last slash.
      // Take everything after the last slash
      returnedFileNameExt = adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]
    } else {
      // The last character comes before the last
      // slash. There is no file name here.
      returnedFileNameExt = ""
    }

  }

  if len(returnedFileNameExt) == 0 {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil
  return
}

// CleanPathStr - Wrapper Function for filepath.Clean()
// See: https://golang.org/pkg/path/filepath/#Clean
// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
// 1. Replace multiple Separator elements with a single one.
// 2. Eliminate each . path name element (the current directory).
// 3. Eliminate each inner .. path name element (the parent directory)
// 		along with the non-.. element that precedes it.
// 4. Eliminate .. elements that begin a rooted path:
// 		that is, replace "/.." by "/" at the beginning of a path,
// 		assuming Separator is '/'.'
// The returned path ends in a slash only if it represents a root
// directory, such as "/" on Unix or `C:\` on Windows.
// Finally, any occurrences of slash are replaced by Separator.
// If the result of this process is an empty string,
// Clean returns the string ".".

func (fh FileHelper) CleanPathStr(pathStr string) string {

  return fp.Clean(pathStr)
}

// ConsolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (fh FileHelper) ConsolidateErrors(errs []error) error {

  lErrs := len(errs)

  if lErrs == 0 {
    return nil
  }

  errStr := ""

  for i := 0; i < lErrs; i++ {

    if errs[i] == nil {
      continue
    }

    tempStr := fmt.Sprintf("%v", errs[i].Error())

    tempStr = strings.TrimLeft(strings.TrimRight(tempStr, " "), " ")

    strLen := len(tempStr)

    for strings.HasSuffix(tempStr,"\n") &&
      strLen > 1 {

      tempStr = tempStr[0:strLen-1]
      strLen--
    }

    if i == (lErrs - 1) {
      errStr += fmt.Sprintf("%v", tempStr)
    } else if i == 0 {
      errStr = fmt.Sprintf("\n%v\n\n", tempStr)
    } else {
      errStr += fmt.Sprintf("%v\n\n", tempStr)
    }
  }

  return fmt.Errorf("%v", errStr)
}

// ConvertDecimalToOctal - Utility routine to convert a decimal (base 10)
// numeric value to an octal (base 8) numeric value. Useful in
// evaluating 'os.FileMode' values and associated constants.
//
//  Reference:
//   https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  initialDecimalValue := 511
//  expectedOctalValue := 777
//
//  actualOctalValue := ConvertDecimalToOctal(initialDecimalValue)
//
//  'actualOctalValue' is now equal to integer value '777'.
//
// ------------------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an integer with a leading
// zero (e.g. x:= int(0777)), than number ('0777') is treated as an octal value
// and converted to a decimal value. Therefore, x:= int(0777) will mean that 'x'
// is set equal to 511. If you set x:= int(777), x will be set equal to '777'.
//
func (fh FileHelper) ConvertDecimalToOctal(number int) int {

  octal := 0
  counter := 1
  remainder := 0

  for number != 0 {
    remainder = number % 8
    number = number / 8
    octal += remainder * counter
    counter *= 10
  }

  return octal
}

// ConvertOctalToDecimal - Utility routine to convert an octal (base 8)
// numeric value to a decimal (base 10) numeric value. Useful in
// evaluating 'os.FileMode' values and associated constants.
//
//  Reference:
//   https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	expectedDecimalValue := 511
//	initialOctalValue := 777
//	actualDecimalValue := FileHelper{}.ConvertOctalToDecimal(initialOctalValue)
//
//  actualDecimalValue is now equal to integer value, '511'.
//
// ------------------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an integer with a leading
// zero (e.g. x:= int(0777)), than number ('0777') is treated as an octal value
// and converted to a decimal value. Therefore, x:= int(0777) will mean that 'x'
// is set equal to 511. If you set x:= int(777), x will be set equal to '777'.
//
func (fh FileHelper) ConvertOctalToDecimal(number int) int {
  decimal := 0
  counter := 0.0
  remainder := 0

  for number != 0 {
    remainder = number % 10
    decimal += remainder * int(math.Pow(8.0, counter))
    number = number / 10
    counter++
  }
  return decimal
}

// CopyFileByLinkByIo - Copies a file from source to destination
// using one of two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a new destination file and using
// "io.Copy(out, in)" to copy the contents. This is accomplished
// by calling 'FileHelper.CopyFileByIo()'. If  the call to
// 'FileHelper.CopyFileByIo()' fails, this method will attempt
// a second copy method.
//
// The second attempt to to copy the designated file will be
// accomplished by creating a 'hard link' to the source file.
// The second, 'hard link', attempt will call method,
// 'FileHelper.CopyFileByLink()'.
//
// If that 'hard link' operation fails, this method will call
// 'FileHelper.CopyFileByIo()'.
//
// If both attempted file copy operations fail, an error will be
// returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fh FileHelper) CopyFileByIoByLink(src, dst string) (err error) {

  ePrefix := "FileHelper.CopyFileByIoByLink() "

  err = fh.CopyFileByIo(src, dst)

  if err == nil {
    return err
  }

  // fh.CopyFileByIo() failed. Try
  // fh.CopyFileByLink()

  errX := fh.CopyFileByLink(src, dst)

  if errX != nil {
    err = fmt.Errorf(ePrefix+
      "Error: After Copy By IO failed, an error was returned "+
      "by fh.CopyFileByLink(src, dst)\n"+
      "src='%v'\ndst='%v'\nError='%v'\n", src, dst, errX)
    return err
  }

  err = nil

  return err
}

// CopyFileByLinkByIo - Copies a file from source to destination
// using one of two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a 'hard link' to the source file.
// The 'hard link' attempt will call 'FileHelper.CopyFileByLink()'.
//
// If that 'hard link' operation fails, this method will call
// 'FileHelper.CopyFileByIo()'.
//
// CopyFileByIo() will create a new destination file and attempt
// to write the contents of the source file to the new destination
// file using "io.Copy(out, in)".
//
// If both attempted file copy operations fail, an error will be
// returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fh FileHelper) CopyFileByLinkByIo(src, dst string) (err error) {

  ePrefix := "FileHelper.CopyFileByLinkByIo() "

  err = fh.CopyFileByLink(src, dst)

  if err == nil {
    return err
  }

  // Copy by Link Failed. Try CopyFileByIo()
  errX := fh.CopyFileByIo(src, dst)

  if errX != nil {
    err = fmt.Errorf(ePrefix+
      "Error: After Copy By Link failed, an error was returned by fh.CopyFileByIo(src, dst).\n"+
      "src='%v'\ndst='%v'\nError='%v'\n",
      src, dst, errX)
    return err
  }

  err = nil

  return err
}

// CopyFileByLink - Copies a file from source to destination
// by means of creating a 'hard link' to the source file,
// "os.Link(src, dst)".
//
// Note: This method of copying files does not create a new
// destination file and write the contents of the source file
// to destination file. (See CopyFileByIo Below).  Instead, this
// method performs the copy operation by creating a hard symbolic
// link to the source file.
//
// By creating a 'linked' file, changing the contents of one file
// will be reflected in the second. The two linked files are
// 'mirrors' of each other.
//
// Consider using CopyFileByIo() if the 'mirror' feature causes problems.
//
// "os.Link(src, dst)" is the only method employed to copy a
// designated file. If "os.Link(src, dst)" fails, an err is returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// REQUIREMENT: The destination Path must previously exist. The destination file
// need NOT exist as it will be created. If the destination file currently
// exists, it will first be deleted and a new linked file will be crated.
//
func (fh FileHelper) CopyFileByLink(src, dst string) (err error) {

  ePrefix := "FileHelper.CopyFileByLink() "
  err = nil
  var err2 error
  var srcFileDoesExist, dstFileDoesExist bool
  var srcFInfo, dstFInfo FileInfoPlus

  src,
    srcFileDoesExist,
    srcFInfo,
    err = fh.doesPathFileExist(src,
    PreProcPathCode.AbsolutePath(), // Covert to Absolute Path
    ePrefix,
    "src")

  if err != nil {
    return err
  }

  dst,
    dstFileDoesExist,
    dstFInfo,
    err = fh.doesPathFileExist(dst,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "dst")

  if err != nil {
    return err
  }

  areSameFile, err2 := fh.AreSameFile(src, dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error occurred during path file name comparison.\n"+
      "Source File:'%v'\n"+
      "Destination File:'%v'\n"+
      "Error='%v'\n",
      src, dst, err2.Error())
    return err
  }

  if areSameFile {
    err = fmt.Errorf(ePrefix+"Error: The source and destination file"+
      " are the same - equivalent.\n"+
      "Source File:'%v'\nDestination File:'%v'\n",
      src, dst)
    return err
  }

  if !srcFileDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: Input parameter 'src' file DOES NOT EXIST!\n"+
      "src='%v'\n", src)
    return err
  }

  if srcFInfo.IsDir() {
    err = fmt.Errorf(ePrefix+"ERROR: Source File (src) is a 'Directory' NOT A FILE!\n"+
      "Source File (src)='%v'\n", src)
    return err
  }

  if !srcFInfo.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+
      "Error: Non-regular source file.\n"+
      "Source File Name='%v'\n"+
      "Source File Mode='%v'\n",
      srcFInfo.Name(), srcFInfo.Mode().String())
    return err
  }

  // If the destination file does NOT exist - this is not a problem
  // because the destination file will be created later.

  if dstFileDoesExist {
    // The destination file exists. This IS a problem. Link will
    // fail when attempting to create a link to an existing file.

    if dstFInfo.IsDir() {
      err = fmt.Errorf(ePrefix+
        "Error: The destination file ('dst') is NOT A FILE.\n"+
        "It is a DIRECTORY!\n"+
        "Destination File ('dst') = '%v'\n",
        dst)
      return err
    }

    if !dstFInfo.Mode().IsRegular() {
      err = fmt.Errorf(ePrefix+
        "Error: The destination file ('dst') is NOT A REGULAR FILE.\n"+
        "Destination File ('dst') = '%v'\n",
        dst)
      return err
    }

    err2 = os.Remove(dst)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error: The target destination file exists and could NOT be deleted! \n"+
        "destination file='%v'\nError='%v'\n", dst, err2.Error())
      return err
    }

    dst,
      dstFileDoesExist,
      _,
      err = fh.doesPathFileExist(dst,
      PreProcPathCode.None(), // Apply no pre-processing conversion to 'dst'
      ePrefix,
      "dst")

    if err != nil {
      return err
    }

    if dstFileDoesExist {
      err = fmt.Errorf(ePrefix+"Error: Deletion of preexisting "+
        "destination file failed!\n"+
        "The copy link operation cannot proceed!\n"+
        "destination file='%v' ", dst)
      return err
    }
  }

  err2 = os.Link(src, dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"- os.Link(src, dst) FAILED!\n"+
      "src='%v'\ndst='%v'\nError='%v'\n",
      src, dst, err2.Error())
    return err
  }

  dst,
    dstFileDoesExist,
    _,
    err2 = fh.doesPathFileExist(dst,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "dst")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: After Copy By Link Operation, a non-path error was returned on 'dst'.\n"+
      "Error='%v'", err2.Error())
    return err
  }

  if !dstFileDoesExist {
    err = fmt.Errorf(ePrefix+
      "Error: After Copy By Link Operation, the destination file DOES NOT EXIST!\n"+
      "Destination File= dst = %v", dst)
    return err
  }

  err = nil
  return err
}

// CopyFileByIo - Copies file from source path and file name to destination
// path and file name.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this method does
// NOT rely on the creation of symbolic links. Instead, a new destination
// file is created and the contents of the source file are written to
// the new destination file using "io.Copy()".
//
// "io.Copy()" is the only method used to copy the designated source
// file. If this method fails, an error is returned.
//
// If source file is equivalent to the destination file, no action will
// be taken and no error will be returned.
//
// If the destination file does not exist, this method will create.
// However, it will NOT create the destination directory. If the
// destination directory does NOT exist, this method will abort
// the copy operation and return an error.
//
func (fh FileHelper) CopyFileByIo(src, dst string) (err error) {
  ePrefix := "FileHelper.CopyFileByIo() "
  err = nil

  var err2, err3 error
  var srcFileDoesExist, dstFileDoesExist bool
  var srcFInfo, dstFileInfo FileInfoPlus

  src,
    srcFileDoesExist,
    srcFInfo,
    err = fh.doesPathFileExist(src,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "src")

  if err != nil {
    return err
  }

  if !srcFileDoesExist {
    err = fmt.Errorf(ePrefix+"Error: Source File DOES NOT EXIST!\n"+
      "src='%v'\n", src)
    return err
  }

  dst,
    dstFileDoesExist,
    dstFileInfo,
    err = fh.doesPathFileExist(dst,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "dst")

  if err != nil {
    return err
  }

  areSameFile, err2 := fh.AreSameFile(src, dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error occurred during path file name comparison.\n"+
      "Source File:'%v'\nDestination File:'%v'\nError='%v'\n",
      src, dst, err2.Error())
    return err
  }

  if areSameFile {
    err = fmt.Errorf(ePrefix+"Error: The source and destination file "+
      "are the same - equivalent.\n"+
      "Source File:'%v'\nDestination File:'%v'\n",
      src, dst)
    return err
  }

  if srcFInfo.IsDir() {
    err = fmt.Errorf(ePrefix+"Error: Source File is 'Directory' and NOT a file!\n"+
      "Source File='%v'\n", src)
    return err
  }

  if !srcFInfo.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+"Error non-regular source file ='%v'\n"+
      "source file Mode='%v'\n",
      srcFInfo.Name(), srcFInfo.Mode().String())
    return err
  }

  if dstFileDoesExist && dstFileInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "Error: 'dst' is a Directory and NOT a File!\n"+
      "dst='%v'", dst)
    return err
  }

  if dstFileDoesExist && !dstFileInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "Error: 'dst' is NOT a 'Regular' File!\n"+
      "dst='%v'\n", dst)
    return err
  }

  // If the destination file does NOT exist, this is not a problem
  // since it will be created later. If the destination 'Path' does
  // not exist, an error return will be triggered.

  // Create a new destination file and copy source
  // file contents to the destination file.

  // First, open the source file
  inSrcPtr, err2 := os.Open(src)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Open(src) src='%v'  Error='%v'",
      src, err2.Error())
    return err
  }

  // Next, 'Create' the destination file
  // If the destination file previously exists,
  // it will be truncated.
  outDestPtr, err2 := os.Create(dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Create(destinationFile)\n"+
      "destinationFile='%v'\nError='%v'\n",
      dst, err2.Error())

    _ = inSrcPtr.Close()

    return err
  }

  bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

  if err2 != nil {
    _ = inSrcPtr.Close()
    _ = outDestPtr.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from io.Copy(destination, source) \n"+
      "destination='%v'\n"+
      "source='%v'\nError='%v'\n",
      dst, src, err2.Error())
    return err
  }

  errs := make([]error, 0)

  // flush file buffers inSrcPtr memory
  err2 = outDestPtr.Sync()

  if err2 != nil {
    err3 = fmt.Errorf(ePrefix+
      "Error returned from outDestPtr.Sync()\n"+
      "outDestPtr=destination='%v'\nError='%v'\n",
      dst, err2.Error())
    errs = append(errs, err3)
  }

  err2 = inSrcPtr.Close()

  if err2 != nil {
    err3 = fmt.Errorf(ePrefix+
      "Error returned from inSrcPtr.Close()\n"+
      "inSrcPtr=source='%v'\nError='%v'\n",
      src, err2.Error())

    errs = append(errs, err3)
  }

  inSrcPtr = nil

  err2 = outDestPtr.Close()

  if err2 != nil {

    err3 = fmt.Errorf(ePrefix+
      "Error returned from outDestPtr.Close()\noutDestPtr=destination='%v'\nError='%v'\n",
      dst, err2.Error())

    errs = append(errs, err3)
  }

  outDestPtr = nil

  if len(errs) > 0 {
    return fh.ConsolidateErrors(errs)
  }

  _,
    dstFileDoesExist,
    dstFileInfo,
    err2 = fh.doesPathFileExist(dst,
    PreProcPathCode.None(), // Do NOT alter path
    ePrefix,
    "dst")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error: After Copy IO operation, dst "+
      "generated non-path error!\n"+
      "dst='%v'\nError='%v'\n",
      dst, err2.Error())
    return err
  }

  if !dstFileDoesExist {
    err = fmt.Errorf(ePrefix+
      "ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
      "Destination File = 'dst' = '%v'\n", dst)

    return err
  }

  srcFileSize := srcFInfo.Size()

  if bytesCopied != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "Error: Bytes Copied does NOT equal bytes "+
      "in source file!\n"+
      "Source File Bytes='%v'   Bytes Coped='%v'\n"+
      "Source File=src='%v'\nDestination File=dst='%v'\n",
      srcFileSize, bytesCopied,
      src, dst)

    return err
  }

  err = nil

  if dstFileInfo.Size() != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes is source file do NOT equal bytes "+
      "in destination file!\n"+
      "Source File Bytes='%v'   Destination File Bytes='%v'\n"+
      "Source File=src='%v'\nDestination File=dst='%v'\n",
      srcFileSize, dstFileInfo.Size(),
      src, dst)
    return err
  }

  return err
}

// CreateFile - Wrapper function for os.Create. If the path component of input
// parameter 'pathFileName' does not exist, a type *PathError will be returned.
//
// This method will 'create' the file designated by input parameter 'pathFileName'.
// 'pathFileName' should consist of a valid path, file name. The file name may consist
// of a file name and file extension or simply a file name.
//
// If successful, this method will return a valid pointer to a type 'os.File' and
// an error value of 'nil'.
//
func (fh FileHelper) CreateFile(pathFileName string) (*os.File, error) {

  ePrefix := "FileHelper.CreateFile() "
  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return nil, errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
  }

  if errCode == -2 {
    return nil, errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
  }

  var err error

  pathFileName, err = fh.MakeAbsolutePath(pathFileName)

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+
        "Error returned by fh.MakeAbsolutePath(pathFileName)\n"+
        "pathFileName='%v'\nError='%v'\n",
        pathFileName, err.Error())
  }

  filePtr, err := os.Create(pathFileName)

  if err != nil {
    return nil, fmt.Errorf(ePrefix+
      "Error returned from os.Create(pathFileName)\n"+
      "pathFileName='%v'\nError='%v'\n",
      pathFileName, err.Error())
  }

  return filePtr, nil
}

// DeleteDirFile - Wrapper function for Remove.
// Remove removes the named file or directory.
// If there is an error, it will be a Non-Path Error.
//
func (fh FileHelper) DeleteDirFile(pathFile string) error {
  ePrefix := "FileHelper.DeleteDirFile() "
  var fileDoesExist bool
  var err error

  pathFile, fileDoesExist, _, err = fh.doesPathFileExist(
    pathFile,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute File Path
    ePrefix,
    "pathFile")

  if err != nil {
    return err
  }

  if !fileDoesExist {
    // Doesn't exist. Nothing to do.
    return nil
  }

  err = os.Remove(pathFile)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from os.Remove(pathFile).\n"+
      "pathFile='%v'\nError='%v'",
      pathFile, err.Error())
  }

  _, fileDoesExist, _, err = fh.doesPathFileExist(
    pathFile,
    PreProcPathCode.None(), // Apply No Pre-Processing. Take no action
    ePrefix,
    "pathFile")

  if err != nil {
    return fmt.Errorf("After attempted deletion, file error occurred!\n"+
      "pathFile='%v'\n"+
      "%v", pathFile, err.Error())
  }

  if fileDoesExist {
    // File STILL Exists! ERROR!
    return fmt.Errorf("ERROR: After attempted deletion, file still exists!\n"+
      "pathFile='%v'\n", pathFile)
  }

  return nil
}

// DeleteDirPathAll - Wrapper function for RemoveAll. This method removes path
// and any children it contains. It removes everything it can but returns the
// first error it encounters. If the path does not exist, this method takes no
// action and returns nil (no error).
//
// Consider the following Example:
//   1. D:\T08\x294_1\x394_1\x494_1  is a directory path that currently exists and
//      contains files.
//   2. Call DeleteDirPathAll("D:\\T08\\x294_1")
//   3. Upon return from method DeleteDirPathAll():
//      a. Deletion Results:
//          Directory D:\T08\x294_1\x394_1\x494_1 and any files in the 'x494_1' directory are deleted
//          Directory D:\T08\x294_1\x394_1\ and any files in the 'x394_1' directory are deleted
//          Directory D:\T08\x294_1\ and any files in the 'x294_1' directory are deleted
//
//      b. The Parent Path 'D:\T08' and any files in that parent path 'D:\T08'
//         directory are unaffected and continue to exist.
//
func (fh FileHelper) DeleteDirPathAll(pathDir string) error {

  ePrefix := "FileHelper.DeleteDirPathAll() "

  var err, err2 error
  var pathFileDoesExist bool

  pathDir,
    pathFileDoesExist,
    _,
    err = fh.doesPathFileExist(pathDir,
    PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
    ePrefix,
    "pathDir")

  if err != nil {
    return err
  }

  if !pathFileDoesExist {
    // Doesn't exist. Nothing to do.
    return nil
  }

  for i := 0; i < 3; i++ {

    err = nil

    err2 = os.RemoveAll(pathDir)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error returned by os.RemoveAll(pathDir).\n"+
        "pathDir='%v'\nError='%v'",
        pathDir, err2.Error())
    }

    if err == nil {
      break
    }

    time.Sleep(50 * time.Millisecond)
  }

  if err != nil {
    return err
  }

  pathDir,
    pathFileDoesExist,
    _,
    err = fh.doesPathFileExist(pathDir,
    PreProcPathCode.None(), // Apply No Pre-Processing. Take No Action.
    ePrefix,
    "pathDir")

  if err != nil {
    return err
  }

  if pathFileDoesExist {
    // Path still exists. Something is wrong.
    return fmt.Errorf("Delete Failed! 'pathDir' still exists!\n"+
      "pathDir='%v'\n", pathDir)
  }

  return nil
}

// DoesFileExist - Returns a boolean value designating whether the passed
// file name exists.
//
// This method does not differentiate between Path Errors and Non-Path
// Errors returned by os.Stat(). The method only returns a boolean
// value.
//
// If a Non-Path Error is returned by os.Stat(), this method will
// classify the file as "Does NOT Exist" and return a value of
// false.
//
// For a more granular test of whether a file exists, see method
// FileHelper.DoesThisFileExist().
func (fh FileHelper) DoesFileExist(pathFileName string) bool {

  _, pathFileDoesExist, _, nonPathError :=
    fh.doesPathFileExist(
      pathFileName,
      PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
      "",
      "pathFileName")

  if !pathFileDoesExist || nonPathError != nil {
    return false
  }

  return true
}

// DeleteFilesWalkDirectory - This method 'walks' the directory tree searching
// for files which match the file selection criteria specified by input parameter
// 'fileSelectCriteria'. When a file matching said 'fileSelectCriteria' is found,
// that file is deleted.
//
// IMPORTANT: This method deletes files!
//
// This method returns file information on files deleted.
//
// If a file matches the File Selection Criteria ('fileSelectCriteria') it is deleted
// and its file information is recorded in the returned DirectoryDeleteFileInfo instance,
// DirectoryDeleteFileInfo.DeletedFiles.
//
// By the way, if ALL the file selection criterion are set to zero values or 'Inactive',
// then ALL FILES in the directory are selected, deleted and returned in the field,
// 'DirectoryDeleteFileInfo.DeletedFiles'.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  startPath                         string - A string consisting of the starting path or
//                                             or directory from which the file search
//                                             operation will commence.
//
//  fileSelectCriteria FileSelectionCriteria -
//	    This input parameter should be configured with the desired file
//	    selection criteria. Files matching this criteria will be returned as
//	    'Found Files'. If file 'fileSelectCriteria' is uninitialized (FileSelectionCriteria{}).
//      all directories the 'startPath' will be searched and all files within those
//      directories will be deleted.
//
//
//       _______________________________________________________________________________________________
//       type FileSelectionCriteria struct {
//         FileNamePatterns     []string    // An array of strings containing File Name Patterns
//         FilesOlderThan       time.Time   // Match files with older modification date times
//         FilesNewerThan       time.Time   // Match files with newer modification date times
//         SelectByFileMode     FilePermissionConfig // Match file mode (os.FileMode).
//         SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//       }
//
//       The FileSelectionCriteria type allows for configuration of single or multiple file
//       selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//       file must match all, or any one, of the active file selection criterion.
//
//       Elements of the FileSelectionCriteria Type are described below:
//
//       FileNamePatterns []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//       FilesOlderThan  time.Time - This date time type is compared to file
//                                   modification date times in order to determine
//                                   whether the file is older than the 'FilesOlderThan'
//                                   file selection criterion. If the file is older than
//                                   the 'FilesOlderThan' date time, that file is considered
//                                   a 'match'	for this file selection criterion.
//
//                                   If the value of 'FilesOlderThan' is set to time zero,
//                                   the default value for type time.Time{}, then this
//                                   file selection criterion is considered to be 'Inactive'
//                                   or 'Not Set'.
//
//      FilesNewerThan   time.Time - This date time type is compared to the file
//                                   modification date time in order to determine
//                                   whether the file is newer than the 'FilesNewerThan'
//                                   file selection criterion. If the file modification date time
//                                   is newer than the 'FilesNewerThan' date time, that file is
//                                   considered a 'match' for this file selection criterion.
//
//                                   If the value of 'FilesNewerThan' is set to time zero,
//                                   the default value for type time.Time{}, then this
//                                   file selection criterion is considered to be 'Inactive'
//                                   or 'Not Set'.
//
//      SelectByFileMode  FilePermissionConfig -
//                                   Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                   selection criterion allows for the selection of files by File Mode.
//                                   File modes are compared to the value	of 'SelectByFileMode'. If the
//                                   File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                   that file is considered to be a 'match' for this file selection
//                                   criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                        fsc := FileSelectionCriteria{}
//                                          err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                          err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//      SelectCriterionMode FileSelectCriterionMode -
//                                   This parameter selects the manner in which the file selection
//                                   criteria above are applied in determining a 'match' for file
//                                   selection purposes. 'SelectCriterionMode' may be set to one of
//                                   two constant values:
//
//                                   _____________________________________________________________________
//
//                                   FileSelectCriterionMode(0).ANDSelect() -
//                                      File selected if all active selection criteria
//                                      are satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
//
//                                   FileSelectCriterionMode(0).ORSelect() -
//                                      File selected if any active selection criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
//                                   _____________________________________________________________________
//
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
//   If all of the file selection criterion in the FileSelectionCriteria object are
//   'Inactive' or 'Not Set' (set to their zero or default values), then all of
//   the files processed in the directory tree will be deleted and returned in the
//   the file manager collection, DirectoryDeleteFileInfo.DeletedFiles.
//
//     Example:
//        FileNamePatterns  = ZERO Length Array
//        filesOlderThan    = time.Time{}
//        filesNewerThan    = time.Time{}
//
//     In this example, all of the selection criterion are
//     'Inactive' and therefore all of the files encountered
//     in the target directory will be selected and returned
//     as 'Found Files'.
//
//     This same effect can be achieved by simply creating an
//     empty file selection instance:
//
//             FileSelectionCriteria{}
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  DirectoryDeleteFileInfo -
//                      If successful, files matching the file selection criteria input
//                      parameter shown above will be deleted and returned in a
//                      'DirectoryDeleteFileInfo' object. The file manager
//                      'DirectoryDeleteFileInfo.DeletedFiles' contains information on all files
//                      deleted during this operation.
//
//                Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//                      to determine if any internal system errors were encountered while processing
//                      the directory tree.
//
//          __________________________________________________________________________________________________
//
//          type DirectoryDeleteFileInfo struct {
//
//            StartPath             string                // The starting path or directory for the file search
//
//            Directories           DirMgrCollection      // Directory Manager instances found during the
//                                                        //   directory tree search.
//            DeletedFiles          FileMgrCollection     // Contains File Managers for Deleted Files matching
//                                                        //   file selection criteria.
//            ErrReturns            []error               // Internal System errors encountered during the search
//                                                        //   and file deletion operations.
//            FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//                                                        // input parameter to this method.
//          }
//
//          __________________________________________________________________________________________________
//
//  error - If a program execution error is encountered during processing, it will
//          be returned as an 'error' type. Also, see the comment on 'DirectoryDeleteFileInfo.ErrReturns',
//          above.
//
func (fh FileHelper) DeleteFilesWalkDirectory(
  startPath string, fileSelectCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {


  ePrefix := "FileHelper.DeleteFilesWalkDirectory() "

  deleteFilesInfo := DirectoryDeleteFileInfo{}

  errCode := 0

  errCode, _, startPath = fh.isStringEmptyOrBlank(startPath)

  if errCode == -1 {
    return deleteFilesInfo,
      errors.New(ePrefix + "Error: Input parameter 'startPath' is an empty string!")
  }

  if errCode == -2 {
    return deleteFilesInfo,
      errors.New(ePrefix + "Error: Input parameter 'startPath' consists of blank spaces!")
  }

  startPath = fh.AdjustPathSlash(startPath)

  strLen := len(startPath)

  if startPath[strLen-1] == os.PathSeparator {
    startPath = startPath[0:strLen-1]
  }

  var err error

  startPath, err = fh.MakeAbsolutePath(startPath)

  if err != nil {
    return deleteFilesInfo,
      fmt.Errorf(ePrefix+"Error returned by fh.MakeAbsolutePath(startPath). "+
        "startPath='%v' Error='%v' ", startPath, err.Error())
  }

  if !fh.DoesFileExist(startPath) {
    return deleteFilesInfo, fmt.Errorf(ePrefix+
      "Error - startPath DOES NOT EXIST! startPath='%v'", startPath)
  }

  deleteFilesInfo.StartPath = startPath

  deleteFilesInfo.DeleteFileSelectCriteria = fileSelectCriteria

  err = fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

  if err != nil {

    return deleteFilesInfo,
      fmt.Errorf(ePrefix+
        "Error returned from fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc"+
        "(&deleteFilesInfo)). startPath='%v' Error='%v'", startPath, err.Error())
  }

  return deleteFilesInfo, nil

}

// DoesFileInfoExist - returns a boolean value indicating
// whether the path and file name passed to the function
// actually exists.
//
// If the file actually exists, the function will return
// the associated FileInfo structure.
//
// If 'pathFileName' does NOT exist, 'doesFInfoExist' will
// be set to'false', 'fInfo' will be set to 'nil' and the
// returned error value will be 'nil'.
//
func (fh FileHelper) DoesFileInfoExist(
  pathFileName string) (doesFInfoExist bool, fInfo os.FileInfo, err error) {

  ePrefix := "FileHelper.DoesFileInfoExist() "
  doesFInfoExist = false
  fInfo = nil
  fInfoPlus := FileInfoPlus{}

  pathFileName,
    doesFInfoExist,
    fInfoPlus,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    doesFInfoExist = false
    return doesFInfoExist, fInfo, err
  }

  if !doesFInfoExist {
    err = nil
    fInfo = nil
    return doesFInfoExist, fInfo, err
  }

  fInfo = fInfoPlus.GetOriginalFileInfo()

  return doesFInfoExist, fInfo, err
}

// DoesStringEndWithPathSeparator - Returns 'true' if the string ends with a
// valid Path Separator. ('/' or '\' depending on the operating system)
//
func (fh FileHelper) DoesStringEndWithPathSeparator(pathStr string) bool {

  errCode := 0
  lenStr := 0

  errCode, lenStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return false
  }

  if pathStr[lenStr-1] == '\\' || pathStr[lenStr-1] == '/' || pathStr[lenStr-1] == os.PathSeparator {
    return true
  }

  return false
}

// DoesThisFileExist - Returns a boolean value signaling whether the path and file name
// represented by input parameter, 'pathFileName', does in fact exist. Unlike the similar
// method FileHelper.DoesFileExist(), this method returns an error in the case of
// Non-Path errors associated with 'pathFileName'.
//
// Non-Path errors may arise for a variety of reasons, but the most common is associated
// with 'access denied' situations.
//
func (fh FileHelper) DoesThisFileExist(pathFileName string) (pathFileNameDoesExist bool,
  nonPathError error) {

  ePrefix := "FileHelper.DoesThisFileExist() "
  pathFileNameDoesExist = false
  nonPathError = nil

  _, pathFileNameDoesExist, _, nonPathError =
    fh.doesPathFileExist(
      pathFileName,
      PreProcPathCode.AbsolutePath(), // Skip Absolute Path Conversion
      ePrefix,
      "pathFileName")

  if nonPathError != nil {
    pathFileNameDoesExist = false
    return pathFileNameDoesExist, nonPathError
  }

  return pathFileNameDoesExist, nonPathError
}

// FilterFileName - Utility method designed to determine whether a file described by a filePath string
// and an os.FileInfo object meets any one of three criteria: A string pattern match, a modification time
// which is older than the 'findFileOlderThan' parameter or a modification time which is newer than the
// 'findFileNewerThan' parameter.
//
// If the three search criteria are all set the their 'zero' or default values, the no selection filter is
// applied and all files are deemed to be a match for the selection criteria ('isMatchedFile=true').
//
// Three selection criterion are applied to the file name (info.Name()).
//
// If a given selection criterion is set to a zero value, then that criterion is defined as 'not set'
// and therefore not used in determining determining whether a file is a 'match'.
//
// If a given criterion is set to a non-zero value, then that criterion is defined as 'set' and the file
// information must comply with that criterion in order to be judged as a match ('isMatchedFile=true').
//
// If none of the three criterion are 'set', then all files are judged as matched ('isMatchedFile=true').
//
// If one of the three criterion is 'set', then a file must comply with that one criterion in order to
// be judged as matched ('isMatchedFile=true').
//
// If two criteria are 'set', then the file must comply with both of those criterion in order to be judged
// as matched ('isMatchedFile=true').
//
// If three criteria are 'set', then the file must comply with all three criterion in order to be judged
// as matched ('isMatchedFile=true').
//
func (fh *FileHelper) FilterFileName(
  info os.FileInfo,
  fileSelectionCriteria FileSelectionCriteria) (isMatchedFile bool, err error) {

  ePrefix := "FileHelper.FilterFileName() "
  isMatchedFile = false
  err = nil

  if info == nil {
    err = errors.New(ePrefix + "Input parameter 'info' is 'nil' and INVALID!")
    return isMatchedFile, err
  }

  isPatternSet, isPatternMatch, err2 := fh.SearchFilePatternMatch(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.SearchFilePatternMatch(info, fileSelectionCriteria).\n"+
      "info.Name()='%v'\nError='%v'\n", info.Name(), err2.Error())
    isMatchedFile = false
    return
  }

  isFileOlderThanSet, isFileOlderThanMatch, err2 := fh.SearchFileOlderThan(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from dMgr.searchFileOlderThan(info, fileSelectionCriteria)\n"+
      "fileSelectionCriteria.FilesOlderThan='%v' info.Name()='%v'\nError='%v'\n",
      fileSelectionCriteria.FilesOlderThan, info.Name(), err2.Error())
    isMatchedFile = false
    return
  }

  isFileNewerThanSet, isFileNewerThanMatch, err2 := fh.SearchFileNewerThan(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from dMgr.searchFileNewerThan(info, fileSelectionCriteria).\n"+
      "fileSelectionCriteria.FilesNewerThan='%v' info.Name()='%v'\nError='%v'\n",
      fileSelectionCriteria.FilesNewerThan, info.Name(), err2.Error())
    isMatchedFile = false
    return
  }

  isFileModeSearchSet, isFileModeSearchMatch, err2 := fh.SearchFileModeMatch(info, fileSelectionCriteria)

  if err2 != nil {
    fileModeTxt := fileSelectionCriteria.SelectByFileMode.GetPermissionFileModeValueText()
    err = fmt.Errorf(ePrefix+
      "Error returned from dMgr.searchFileModeMatch(info, fileSelectionCriteria).\n"+
      "fileSelectionCriteria.SelectByFileMode='%v'\n"+
      "info.Name()='%v' Error='%v'\n",
      fileModeTxt,
      info.Name(), err2.Error())
    isMatchedFile = false
    return
  }

  // If no file selection criterion are set, then always select the file
  if !isPatternSet && !isFileOlderThanSet && !isFileNewerThanSet && !isFileModeSearchSet {
    isMatchedFile = true
    err = nil
    return
  }

  // If using the AND File Select Criterion Mode, then for criteria that
  // are set and active, they must all be 'matched'.
  if fileSelectionCriteria.SelectCriterionMode == fileSelectMode.ANDSelect() {

    if isPatternSet && !isPatternMatch {
      isMatchedFile = false
      err = nil
      return
    }

    if isFileOlderThanSet && !isFileOlderThanMatch {
      isMatchedFile = false
      err = nil
      return
    }

    if isFileNewerThanSet && !isFileNewerThanMatch {
      isMatchedFile = false
      err = nil
      return
    }

    if isFileModeSearchSet && !isFileModeSearchMatch {
      isMatchedFile = false
      err = nil
      return
    }

    isMatchedFile = true
    err = nil
    return

  } // End of fileSelectMode.ANDSelect()

  // Must be fileSelectMode.ORSelect() Mode
  // If ANY of the section criterion are active and 'matched', then
  // classify the file as matched.

  if isPatternSet && isPatternMatch {
    isMatchedFile = true
    err = nil
    return
  }

  if isFileOlderThanSet && isFileOlderThanMatch {
    isMatchedFile = true
    err = nil
    return
  }

  if isFileNewerThanSet && isFileNewerThanMatch {
    isMatchedFile = true
    err = nil
    return
  }

  if isFileModeSearchSet && isFileModeSearchMatch {
    isMatchedFile = true
    err = nil
    return
  }

  isMatchedFile = false
  err = nil
  return
}

// FindFilesInPath - Will apply a search pattern to files and directories
// in the path designated by input parameter, 'pathName'. If the files
// and or directory names match the input parameter, 'fileSearchPattern'
// they will be returned in an array of strings.
//
// Be Advised!  The names returned in the string array may consist of both
// files and directory names, depending on the specified, 'fileSearchPattern'.
//
// This method uses the "path/filepath" function, 'Glob'. Reference:
//				https://golang.org/pkg/path/filepath/#Glob
//
// The File matching patterns depend on the 'go' "path/filepath" function,
// 'Match'.  Reference
// https://golang.org/pkg/path/filepath/#Match
//
// Note: This method will NOT search sub-directories. It will return the names
// of directories existing in the designated, 'pathName', depending on the
// 'fileSearchPattern' passed as an input parameter.
//
// If Input Parameters 'pathName' or 'fileSearchPattern' are empty strings or consist
// of all space characters, this method will return an error.
//
//   Example 'fileSearchPattern' values:
//         "*"     = Returns all files and directories (everything)
//         "*.*"   = Returns files which have a file extension
//         "*.txt" = Returns only files with a "txt" file extension
//
func (fh FileHelper) FindFilesInPath(pathName, fileSearchPattern string) ([]string, error) {

  ePrefix := "FileHelper.FindFilesInPath() "

  var pathDoesExist bool
  var fInfo FileInfoPlus
  var err error
  var errCode int

  pathName,
    pathDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathName")

  if err != nil {
    return []string{}, err
  }

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode == -1 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'fileSearchPattern' is " +
        "an empty string!\n")
  }

  if errCode == -2 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'fileSearchPattern' consists " +
        "of blank spaces!\n")
  }

  if !pathDoesExist {
    return []string{},
      fmt.Errorf(ePrefix+"Error: Input parameter 'pathName' DOES NOT EXIST!\n"+
        "pathName='%v'\n", pathName)
  }

  if !fInfo.IsDir() {
    return []string{},
      fmt.Errorf(ePrefix+"Error: The path exists, but it NOT a directory!\n"+
        "pathName='%v' ", pathName)
  }

  // fInfo is a Directory.

  searchStr := fh.JoinPathsAdjustSeparators(pathName, fileSearchPattern)

  results, err := fp.Glob(searchStr)

  if err != nil {
    return []string{},
      fmt.Errorf(ePrefix+
        "Error returned by fp.Glob(searchStr).\n"+
        "searchStr='%v'\nError='%v'\n",
        searchStr, err.Error())
  }

  return results, nil
}

// FindFilesWalkDirectory - This method returns file information on files residing in a specified
// directory tree identified by the input parameter, 'startPath'.
//
// This method 'walks the directory tree' locating all files in the directory tree which match
// the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. By the way, if ALL the file selection criterion are set to zero values
// or 'Inactive', then ALL FILES in the directory are selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  startPath                         string - A string consisting of the starting path or
//                                             or directory from which the find files search
//                                             operation will commence.
//
//  fileSelectCriteria FileSelectionCriteria -
//	    This input parameter should be configured with the desired file
//	    selection criteria. Files matching this criteria will be returned as
//	    'Found Files'. If file 'fileSelectCriteria' is uninitialized (FileSelectionCriteria{}).
//      all directories and files will be returned from the 'startPath'
//
//
//       _______________________________________________________________________________________________
//       type FileSelectionCriteria struct {
//         FileNamePatterns     []string    // An array of strings containing File Name Patterns
//         FilesOlderThan       time.Time   // Match files with older modification date times
//         FilesNewerThan       time.Time   // Match files with newer modification date times
//         SelectByFileMode     FilePermissionConfig // Match file mode (os.FileMode).
//         SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//       }
//
//       The FileSelectionCriteria type allows for configuration of single or multiple file
//       selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//       file must match all, or any one, of the active file selection criterion.
//
//       Elements of the FileSelectionCriteria Type are described below:
//
//       FileNamePatterns []string  - An array of strings which may define one or more
//                                    search patterns. If a file name matches any one of the
//                                    search pattern strings, it is deemed to be a 'match'
//                                    for the search pattern criterion.
//
//                                      Example Patterns:
//                                       FileNamePatterns = []string{"*.log"}
//                                       FileNamePatterns = []string{"current*.txt"}
//                                       FileNamePatterns = []string{"*.txt", "*.log"}
//
//                                    If this string array has zero length or if
//                                    all the strings are empty strings, then this
//                                    file search criterion is considered 'Inactive'
//                                    or 'Not Set'.
//
//
//       FilesOlderThan  time.Time - This date time type is compared to file
//                                   modification date times in order to determine
//                                   whether the file is older than the 'FilesOlderThan'
//                                   file selection criterion. If the file is older than
//                                   the 'FilesOlderThan' date time, that file is considered
//                                   a 'match'	for this file selection criterion.
//
//                                   If the value of 'FilesOlderThan' is set to time zero,
//                                   the default value for type time.Time{}, then this
//                                   file selection criterion is considered to be 'Inactive'
//                                   or 'Not Set'.
//
//      FilesNewerThan   time.Time - This date time type is compared to the file
//                                   modification date time in order to determine
//                                   whether the file is newer than the 'FilesNewerThan'
//                                   file selection criterion. If the file modification date time
//                                   is newer than the 'FilesNewerThan' date time, that file is
//                                   considered a 'match' for this file selection criterion.
//
//                                   If the value of 'FilesNewerThan' is set to time zero,
//                                   the default value for type time.Time{}, then this
//                                   file selection criterion is considered to be 'Inactive'
//                                   or 'Not Set'.
//
//      SelectByFileMode  FilePermissionConfig -
//                                   Type FilePermissionConfig encapsulates an os.FileMode. The file
//                                   selection criterion allows for the selection of files by File Mode.
//                                   File modes are compared to the value	of 'SelectByFileMode'. If the
//                                   File Mode for a given file is equal to the value of 'SelectByFileMode',
//                                   that file is considered to be a 'match' for this file selection
//                                   criterion. Examples for setting SelectByFileMode are shown as follows:
//
//                                        fsc := FileSelectionCriteria{}
//                                          err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//                                          err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//      SelectCriterionMode FileSelectCriterionMode -
//                                   This parameter selects the manner in which the file selection
//                                   criteria above are applied in determining a 'match' for file
//                                   selection purposes. 'SelectCriterionMode' may be set to one of
//                                   two constant values:
//
//                                   _____________________________________________________________________
//
//                                   FileSelectCriterionMode(0).ANDSelect() -
//                                      File selected if all active selection criteria
//                                      are satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will not be judged as 'selected' unless all of
//                                      the active selection criterion are satisfied. In other words, if
//                                      three active search criterion are provided for 'FileNamePatterns',
//                                      'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//                                      selected unless it has satisfied all three criterion in this example.
//
//                                   FileSelectCriterionMode(0).ORSelect() -
//                                      File selected if any active selection criterion is satisfied.
//
//                                      If this constant value is specified for the file selection mode,
//                                      then a given file will be selected if any one of the active file
//                                      selection criterion is satisfied. In other words, if three active
//                                      search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//                                      and 'FilesNewerThan', then a file will be selected if it satisfies any
//                                      one of the three criterion in this example.
//
//                                   _____________________________________________________________________
//
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
//   If all of the file selection criterion in the FileSelectionCriteria object are
//   'Inactive' or 'Not Set' (set to their zero or default values), then all of
//   the files processed in the directory tree will be selected and returned as
//   'Found Files'.
//
//     Example:
//        FileNamePatterns  = ZERO Length Array
//        filesOlderThan    = time.Time{}
//        filesNewerThan    = time.Time{}
//
//     In this example, all of the selection criterion are
//     'Inactive' and therefore all of the files encountered
//     in the target directory will be selected and returned
//     as 'Found Files'.
//
//     This same effect can be achieved by simply creating an
//     empty file selection instance:
//
//             FileSelectionCriteria{}
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  DirectoryTreeInfo - If successful, files matching the file selection criteria input
//                      parameter shown above will be returned in a 'DirectoryTreeInfo'
//                      object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//                      on all the files in the specified directory tree which match the file selection
//                      criteria.
//
//                Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//                      to determine if any internal system errors were encountered while processing
//                      the directory tree.
//
//          __________________________________________________________________________________________________
//
//          type DirectoryTreeInfo struct {
//            StartPath             string                // The starting path or directory for the file search
//            dirMgrs               []DirMgr              // dirMgrs found during directory tree search
//            FoundFiles            []FileWalkInfo        // Found Files matching file selection criteria
//            ErrReturns            []string              // Internal System errors encountered
//            FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//                                                        // input parameter to this method.
//          }
//
//          __________________________________________________________________________________________________
//
//  error - If a program execution error is encountered during processing, it will
//          be returned as an 'error' type. Also, see the comment on 'DirectoryTreeInfo.ErrReturns',
//          above.
//
func (fh FileHelper) FindFilesWalkDirectory(
  startPath string,
  fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

  ePrefix := "FileHelper.FindFilesWalkDirectory() "

  findFilesInfo := DirectoryTreeInfo{}

  errCode := 0

  errCode, _, startPath = fh.isStringEmptyOrBlank(startPath)

  if errCode == -1 {
    return findFilesInfo,
      errors.New(ePrefix + "Error: Input parameter 'startPath' is an empty string!")
  }

  if errCode == -2 {
    return findFilesInfo,
      errors.New(ePrefix + "Error: Input parameter 'startPath' consists of blank spaces!")
  }

  startPath = fh.RemovePathSeparatorFromEndOfPathString(startPath)

  var err error

  startPath, err = fh.MakeAbsolutePath(startPath)

  if err != nil {
    return findFilesInfo,
      fmt.Errorf(ePrefix+"Error returned by fh.MakeAbsolutePath(startPath). "+
        "startPath='%v' Error='%v' ", startPath, err.Error())
  }

  if !fh.DoesFileExist(startPath) {
    return findFilesInfo, fmt.Errorf(ePrefix+
      "Error - startPath DOES NOT EXIST! startPath='%v'", startPath)
  }

  findFilesInfo.StartPath = startPath

  findFilesInfo.FileSelectCriteria = fileSelectCriteria

  err = fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

  if err != nil {

    return findFilesInfo,
      fmt.Errorf(ePrefix+
        "Error returned from fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc"+
        "(&findFilesInfo)). startPath='%v' Error='%v'", startPath, err.Error())
  }

  return findFilesInfo, nil
}

// GetAbsCurrDir - Returns the absolute path of the current working
// directory.
//
// The current work directory is determined by a call to os.Getwd().
// 'Getwd()' returns a rooted path name corresponding to the current directory.
// If the current directory can be reached via multiple paths (due to
// symbolic links), 'Getwd()' may return any one of them.
//
func (fh FileHelper) GetAbsCurrDir() (string, error) {
  ePrefix := "FileHelper.GetAbsCurrDir() "

  dir, err := os.Getwd()

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Error returned from os.Getwd().\nError='%v'\n",
      err.Error())
  }

  absDir, err := fh.MakeAbsolutePath(dir)

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(dir).\nError='%v'\n",
      err.Error())

  }

  return absDir, nil
}

// GetAbsPathFromFilePath - Supply a string containing both the path file name and extension.
// This method will then return the absolute value of that path, file name and file extension.
//
func (fh FileHelper) GetAbsPathFromFilePath(filePath string) (string, error) {

  ePrefix := "FileHelper.GetAbsPathFromFilePath() "

  errCode := 0

  errCode, _, filePath = fh.isStringEmptyOrBlank(filePath)

  if errCode == -1 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'filePath' is an empty string!")
  }

  if errCode == -2 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'filePath' consists of blank spaces!")
  }

  testFilePath := fh.AdjustPathSlash(filePath)

  errCode, _, testFilePath = fh.isStringEmptyOrBlank(testFilePath)

  if errCode < 0 {
    return "",
      errors.New(ePrefix +
        "Error: After adjusting path Separators, filePath resolves to an empty string!")
  }

  absPath, err := fh.MakeAbsolutePath(testFilePath)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned from fh.MakeAbsolutePath(testFilePath). "+
        "testFilePath='%v' Error='%v' ",
        testFilePath, err.Error())
  }

  return absPath, nil
}

// GetCurrentDir - Wrapper function for Getwd(). Getwd returns a
// rooted path name corresponding to the current directory.
// If the current directory can be reached via multiple paths
// (due to symbolic links), Getwd may return any one of them.
func (fh FileHelper) GetCurrentDir() (string, error) {

  ePrefix := "FileHelper.GetCurrentDir()"

  currDir, err := os.Getwd()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+"Error returned by os.Getwd(). Error='%v' ",
        err.Error())
  }

  return currDir, nil
}

// GetDotSeparatorIndexesInPathStr - Returns an array of integers representing the
// indexes of dots ('.') located in input parameter 'pathStr'.
//
func (fh FileHelper) GetDotSeparatorIndexesInPathStr(pathStr string) ([]int, error) {

  ePrefix := "FileHelper.GetDotSeparatorIndexesInPathStr() "

  errCode := 0

  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    return []int{},
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' is an empty string!\n")
  }

  if errCode == -2 {
    return []int{},
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' consists of blank spaces!\n")
  }

  var dotIdxs []int

  for i := 0; i < lPathStr; i++ {

    rChar := pathStr[i]

    if rChar == '.' {

      dotIdxs = append(dotIdxs, i)
    }

  }

  return dotIdxs, nil
}

// GetExecutablePathFileName - Gets the path and file name of the
// executable that started the current process.
//
// This executable path and file name is generated by a call to
// os.Executable().
//
// os.Executable() returns the path name for the executable that started
// the current process. There is no guarantee that the path is still
// pointing to the correct executable. If a symlink was used to start
// the process, depending on the operating system, the result might
// be the symlink or the path it pointed to. If a stable result is
// needed, path/filepath.EvalSymlinks might help.
//
// Executable returns an absolute path unless an error occurred.
//
// The main use case is finding resources located relative to an
// executable.
//
// Executable is not supported on nacl.
//
func (fh FileHelper) GetExecutablePathFileName() (string, error) {

  ePrefix := "FileHelper.GetExecutablePathFileName() "

  ex, err := os.Executable()

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+"Error returned by os.Executable(). Error='%v' ",
        err.Error())
  }

  return ex, err
}

// GetFileExt - Returns the File Extension with
// the dot. If there is no File Extension an empty
// string is returned (NO dot included). If the returned
// File Extension is an empty string, the returned
// parameter 'isEmpty' is set equal to 'true'.
//
// When an extension is returned in the 'ext' variable, this
// extension includes a leading dot. Example: '.txt'
//
//    Example:
//
//     Actual File Name Plus Extension: "newerFileForTest_01.txt"
//             Returned File Extension: "txt"
//
//     Actual File Name Plus Extension: "newerFileForTest_01"
//             Returned File Extension: ""
//
//     Actual File Name Plus Extension: ".gitignore"
//             Returned File Extension: ""
//
func (fh FileHelper) GetFileExtension(
  pathFileNameExt string) (ext string, isEmpty bool, err error) {
  ePrefix := "FileHelper.GetFileExt() "

  ext = ""
  isEmpty = true
  err = nil

  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")

    return ext, isEmpty, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")

    return ext, isEmpty, err
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  errCode, _, testPathFileNameExt = fh.isStringEmptyOrBlank(testPathFileNameExt)

  if errCode < 0 {
    err = errors.New(ePrefix +
      "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is an empty string!")

    return ext, isEmpty, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt).\n"+
      "testPathFileNameExt='%v'\nError='%v'\n", testPathFileNameExt, err2)
    return ext, isEmpty, err
  }

  lenDotIdxs := len(dotIdxs)

  // Deal with case where the pathFileNameExt contains
  // no dots.
  if lenDotIdxs == 0 {
    ext = ""
    isEmpty = true
    err = nil
    return ext, isEmpty, err

  }

  firstGoodCharIdx, lastGoodCharIdx, err2 :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt).\n"+
      "testPathFileNameExt='%v'\nError='%v'\n",
      testPathFileNameExt, err2)
    return ext, isEmpty, err
  }

  // Deal with the case where pathFileNameExt contains no
  // valid alpha numeric characters
  if firstGoodCharIdx == -1 || lastGoodCharIdx == -1 {
    ext = ""
    isEmpty = true
    err = nil
    return ext, isEmpty, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
    return ext, isEmpty, err
  }

  lenSlashIdxs := len(slashIdxs)

  if lenSlashIdxs == 0 &&
    lenDotIdxs == 1 &&
    dotIdxs[lenDotIdxs-1] == 0 {
    // deal with the case .gitignore
    ext = ""
    isEmpty = true
    err = nil
    return ext, isEmpty, err
  }

  if lenSlashIdxs == 0 {
    ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
    isEmpty = false
    err = nil
    return ext, isEmpty, err
  }

  // lenDotIdxs and lenSlasIdxs both greater than zero
  if dotIdxs[lenDotIdxs-1] > slashIdxs[lenSlashIdxs-1] &&
    dotIdxs[lenDotIdxs-1] < lastGoodCharIdx {

    ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
    isEmpty = false
    err = nil
    return ext, isEmpty, err

  }

  ext = ""
  isEmpty = true
  err = nil
  return ext, isEmpty, err
}

// GetFileInfo - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on a specific file. If the file
// does NOT exist, an error will be triggered.
//
// This method is similar to FileHelper.DoesFileInfoExist().
//
//  type FileInfo interface {
//    Name()    string       // base name of the file
//    Size()    int64        // length in bytes for regular files; system-dependent for others
//    Mode()    FileMode     // file mode bits
//    ModTime() time.Time    // modification time
//    IsDir()   bool         // abbreviation for Mode().IsDir()
//    Sys()     interface{}  // underlying data source (can return nil)
//  }
//
func (fh FileHelper) GetFileInfo(pathFileName string) (os.FileInfo, error) {

  ePrefix := "FileHelper.GetFileInfo() "
  var pathDoesExist bool
  var fInfo FileInfoPlus
  var err error

  pathFileName,
    pathDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return nil, err
  }

  if !pathDoesExist {
    return nil, fmt.Errorf(ePrefix+
      "Error: Input parameter 'pathFileName' does NOT exist!\n"+
      "pathFileName='%v'\n",
      pathFileName)
  }

  return fInfo.GetOriginalFileInfo(), nil
}

// GetFileLastModificationDate - Returns the last modification'
// date/time on a specific file. If input parameter 'customTimeFmt'
// string is empty, a default time format will be used to format the
// returned time string.
//
// The default date time format is:
//     "2006-01-02 15:04:05.000000000"
//
func (fh FileHelper) GetFileLastModificationDate(
  pathFileName string,
  customTimeFmt string) (time.Time, string, error) {

  ePrefix := "FileHelper.GetFileLastModificationDate() "
  const fmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000 -0700 MST"
  var zeroTime time.Time
  var pathFileNameDoesExist bool
  var fInfo FileInfoPlus
  var err error
  var errCode int

  pathFileName,
    pathFileNameDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathFileName,
    PreProcPathCode.AbsolutePath(), // Skip Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return zeroTime, "", err
  }

  if !pathFileNameDoesExist {
    return zeroTime, "",
      fmt.Errorf(ePrefix+"ERROR: 'pathFileName' DOES NOT EXIST!\n"+
        "pathFileName='%v'\n", pathFileName)
  }

  errCode, _, customTimeFmt = fh.isStringEmptyOrBlank(customTimeFmt)

  fmtStr := customTimeFmt

  if errCode < 0 {
    fmtStr = fmtDateTimeNanoSecondStr
  }

  return fInfo.ModTime(), fInfo.ModTime().Format(fmtStr), nil
}

// GetFileMode - Returns the mode (os.FileMode) for the file designated by
// input parameter 'pathFileName'. If the file does not exist, an error
// is triggered.
//
// The 'os.FileMode' is returned via type 'FilePermissionCfg' which includes
// methods necessary to interpret the 'os.FileMode'.
//
func (fh FileHelper) GetFileMode(pathFileName string) (FilePermissionConfig, error) {

  ePrefix := "FileHelper.GetFileMode() "
  var pathFileNameDoesExist bool
  var fInfo FileInfoPlus
  var err error

  pathFileName,
    pathFileNameDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return FilePermissionConfig{}, err
  }

  if !pathFileNameDoesExist {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "ERROR: 'pathFileName' DOES NOT EXIST!\n"+
        "pathFileName='%v'\n", pathFileName)
  }

  fPermCfg, err := FilePermissionConfig{}.NewByFileMode(fInfo.Mode())

  if err != nil {
    return FilePermissionConfig{},
      fmt.Errorf(ePrefix+
        "Error returned by FilePermissionConfig{}.NewByFileMode(fInfo.Mode()).\n"+
        "fInfo.Mode()='%v'\nError='%v'\n",
        fInfo.Mode(), err.Error())
  }

  return fPermCfg, nil
}

// GetFileNameWithExt - This method expects to receive a valid directory path and file
// name or file name plus extension. It then extracts the File Name and Extension from
// the file path and returns it as a string.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  pathFileNameExt string - This input parameter is expected to contain a properly formatted directory
//                           path and File Name.  The File Name may or may not include a File Extension.
//                           The directory path must include the correct delimiters such as path Separators
//                           ('/' or'\'), dots ('.') and in the case of Windows, a volume designation
//                           (Example: 'F:').
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  fNameExt  string       - If successful, this method will the return value 'fNameExt' equal to the
//                           File Name and File Extension extracted from the input file path, 'pathFileNameExt'.
//                           Example 'fNameExt' return value: 'someFilename.txt'
//
//                           If the File Extension is not present, only the File Name will be returned.
//                           Example return value with no file extension: 'someFilename'.
//
//  isEmpty   bool         - If this method CAN NOT parse out a valid File Name and Extension from
//                           input parameter string 'pathFileNameExt', return value 'fNameExt' will
//                           be set to an empty string and return value 'isEmpty' will be set to 'true'.
//
//  err       error        - If an error is encountered during processing, 'err' will return a properly
//                           formatted 'error' type.
//
//                           Note that if this method cannot parse out a valid	File Name and Extension
//                           due to an improperly formatted directory path and	file name string
//                           (Input Parameter: 'pathFileNameExt'), 'fNameExt' will be set to an empty string,
//                           'isEmpty' will be set to 'true' and 'err' return 'nil'. In this situation, no
//                           error will be returned.
//
func (fh FileHelper) GetFileNameWithExt(
  pathFileNameExt string) (fNameExt string, isEmpty bool, err error) {

  ePrefix := "FileHelper.GetFileNameWithExt "
  fNameExt = ""
  isEmpty = true
  err = nil
  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")
    return fNameExt, isEmpty, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")
    return fNameExt, isEmpty, err
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  volName := fh.GetVolumeName(testPathFileNameExt)

  if volName != "" {
    testPathFileNameExt = strings.TrimPrefix(testPathFileNameExt, volName)
  }

  lTestPathFileNameExt := 0

  errCode, lTestPathFileNameExt, testPathFileNameExt = fh.isStringEmptyOrBlank(testPathFileNameExt)

  if errCode < 0 {
    err = errors.New(ePrefix +
      "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is an empty string!")
    return fNameExt, isEmpty, err
  }

  firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return fNameExt, isEmpty, err
  }

  // There are no alpha numeric characters present.
  // Therefore, there is no file name and extension
  if firstCharIdx == -1 || lastCharIdx == -1 {
    isEmpty = true
    err = nil
    return fNameExt, isEmpty, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return fNameExt, isEmpty, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return fNameExt, isEmpty, err
  }

  lSlashIdxs := len(slashIdxs)
  lDotIdxs := len(dotIdxs)

  if lSlashIdxs > 0 {
    // This string has path separators

    // Last char is a path separator. Therefore,
    // there is no file name and extension.
    if slashIdxs[lSlashIdxs-1] == lTestPathFileNameExt-1 {
      fNameExt = ""
    } else if lastCharIdx > slashIdxs[lSlashIdxs-1] {

      fNameExt = testPathFileNameExt[slashIdxs[lSlashIdxs-1]+1:]

    } else {
      fNameExt = ""
    }

    if len(fNameExt) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return fNameExt, isEmpty, err
  }

  // There are no path separators lSlashIdxs == 0

  if lDotIdxs > 0 {
    // This string has one or more dot separators ('.')

    fNameExt = ""

    if firstCharIdx > dotIdxs[lDotIdxs-1] {
      // Example '.txt' - Valid File name and extension
      // such as '.gitignore'
      fNameExt = testPathFileNameExt[dotIdxs[lDotIdxs-1]:]

    } else if firstCharIdx < dotIdxs[lDotIdxs-1] {
      fNameExt = testPathFileNameExt[firstCharIdx:]
    }

    if len(fNameExt) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return fNameExt, isEmpty, err
  }

  // Must be lSlashIdxs == 0 && lDotIdxs ==  0
  // There are no path Separators and there are
  // no dot separators ('.').

  fNameExt = testPathFileNameExt[firstCharIdx:]

  if len(fNameExt) == 0 {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil

  return fNameExt, isEmpty, err
}

// GetFileNameWithoutExt - returns the file name
// without the path or extension. If the returned
// File Name is an empty string, isEmpty is set to true.
//
//
//    Example:
//
//          Actual Path Plus File Name: = "./pathfilego/003_filehelper/common/xt_dirmgr_01_test.go"
//                  Returned File Name: = "dirmgr_01_test"
//
//     Actual File Name Plus Extension: "newerFileForTest_01.txt"
//                  Returned File Name: "newerFileForTest_01"
//
//     Actual File Name Plus Extension: "newerFileForTest_01"
//                  Returned File Name: "newerFileForTest_01"
//
//     Actual File Name Plus Extension: ".gitignore"
//                  Returned File Name: ".gitignore"
//
//
func (fh FileHelper) GetFileNameWithoutExt(
  pathFileNameExt string) (fName string, isEmpty bool, err error) {

  ePrefix := "FileHelper.GetFileNameWithoutExt() "

  fName = ""
  isEmpty = true
  err = nil
  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")
    return fName, isEmpty, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")
    return fName, isEmpty, err
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  errCode, _, testPathFileNameExt = fh.isStringEmptyOrBlank(testPathFileNameExt)

  if errCode < 0 {
    err = errors.New(ePrefix +
      "Error: Adjusted path version of 'pathFileNameExt', 'testPathFileNameExt' is an empty string!")
    return fName, isEmpty, err
  }

  fileNameExt, isFileNameExtEmpty, err2 := fh.GetFileNameWithExt(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFileNameWithExt(testPathFileNameExt) testPathFileNameExt='%v'  Error='%v'",
      testPathFileNameExt, err2.Error())
    return fName, isEmpty, err
  }

  if isFileNameExtEmpty {
    isEmpty = true
    fName = ""
    err = nil
    return fName, isEmpty, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetDotSeparatorIndexesInPathStr(fileNameExt). fileNameExt='%v'  Error='%v'",
      fileNameExt, err2.Error())
    return fName, isEmpty, err
  }

  lDotIdxs := len(dotIdxs)

  if lDotIdxs == 1 &&
    dotIdxs[lDotIdxs-1] == 0 {
    // Outlier Case: .gitignore
    fName = fileNameExt[0:]

    if fName == "" {
      isEmpty = true
    } else {
      isEmpty = false
    }
    err = nil
    return fName, isEmpty, err
  }

  // Primary Case: filename.ext
  if lDotIdxs > 0 {
    fName = fileNameExt[0:dotIdxs[lDotIdxs-1]]

    if fName == "" {
      isEmpty = true
    } else {
      isEmpty = false
    }
    err = nil
    return fName, isEmpty, err
  }

  // Secondary Case: filename
  fName = fileNameExt

  if fName == "" {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil
  return fName, isEmpty, err
}

// GetFirstLastNonSeparatorCharIndexInPathStr - Basically this method returns
// the first index of the first alpha numeric character in a path string.
//
// Specifically, the character must not be a path Separator ('\', '/') and
// it must not be a dot ('.').
//
// If the first Non-Separator char is found, this method will return
// an integer index which is greater than or equal to zero plus an
// error value of nil.
//
// The first character found will never be part of the volume name.
// Example On Windows: "D:\fDir1\fDir2" - first character index will
// be 3 denoting character 'f'.
//
func (fh FileHelper) GetFirstLastNonSeparatorCharIndexInPathStr(
  pathStr string) (firstIdx, lastIdx int, err error) {

  ePrefix := "FileHelper.GetFirstNonSeparatorCharIndexInPathStr() "
  firstIdx = -1
  lastIdx = -1
  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {

    err = errors.New(ePrefix +
      "Error: Input parameter 'pathStr' is an empty string!\n")

    return firstIdx, lastIdx, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!\n")

    return firstIdx, lastIdx, err
  }

  pathStr = fh.AdjustPathSlash(pathStr)

  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {

    err = fmt.Errorf(ePrefix +
      "Error: After path Separator adjustment, 'pathStr' is an empty string!\n")

    return firstIdx, lastIdx, err
  }

  // skip the volume name. Don't count
  // first characters in the volume name
  volName := fp.VolumeName(pathStr)
  lVolName := len(volName)

  startIdx := 0

  if lVolName > 0 {
    startIdx = lVolName
  }

  var rChar rune

  forbiddenTextChars := []rune{os.PathSeparator,
    '\\',
    '/',
    '|',
    '.',
    '&',
    '!',
    '%',
    '$',
    '#',
    '@',
    '^',
    '*',
    '(',
    ')',
    '-',
    '_',
    '+',
    '=',
    '[',
    '{',
    ']',
    '}',
    '|',
    '<',
    '>',
    ',',
    '~',
    '`',
    ':',
    ';',
    '"',
    '\'',
    '\n',
    '\t',
    '\r'}

  lForbiddenTextChars := len(forbiddenTextChars)

  for i := startIdx; i < lPathStr; i++ {
    rChar = rune(pathStr[i])
    isForbidden := false

    for j := 0; j < lForbiddenTextChars; j++ {
      if rChar == forbiddenTextChars[j] {
        isForbidden = true
      }

    }

    if isForbidden == false {

      if firstIdx == -1 {
        firstIdx = i
      }

      lastIdx = i
    }

  }

  err = nil

  return firstIdx, lastIdx, err
}

// GetLastPathElement - Analyzes a 'pathName' string and returns the last
// element in the path. If 'pathName' ends in a path separator ('/'), this
// method returns an empty string.
//
//  Example:
//
//  pathName = '../dir1/dir2/fileName.ext' will return "fileName.ext"
//  pathName = '../dir1/dir2/' will return ""
//  pathName = 'fileName.ext' will return "fileName.ext"
//  pathName = '../dir1/dir2/dir3' will return "dir3"
//
func (fh FileHelper) GetLastPathElement(pathName string) (string, error) {
  ePrefix := "FileHelper.GetLastPathElement() "
  errCode := 0

  errCode, _, pathName = fh.isStringEmptyOrBlank(pathName)

  if errCode == -1 {
    return "",
      errors.New(ePrefix +
        "Error: Input parameter 'pathName' is an empty string!\n")
  }

  if errCode == -2 {
    return "",
      errors.New(ePrefix +
        "Error: Input parameter 'pathName' consists of blank spaces!\n")
  }

  adjustedPath := fh.AdjustPathSlash(pathName)

  resultAry := strings.Split(adjustedPath, string(os.PathSeparator))

  lResultAry := len(resultAry)

  if lResultAry == 0 {
    return adjustedPath, nil
  }

  return resultAry[lResultAry-1], nil
}

// GetPathAndFileNameExt - Breaks out path and fileName+Ext elements from
// a path string. If both path and fileName are empty strings, this method
// returns an error.
func (fh FileHelper) GetPathAndFileNameExt(
  pathFileNameExt string) (pathDir, fileNameExt string, bothAreEmpty bool, err error) {

  ePrefix := "FileHelper.GetPathAndFileNameExt() "
  pathDir = ""
  fileNameExt = ""
  bothAreEmpty = true
  err = nil
  errCode := 0
  trimmedFileNameExt := ""

  errCode, _, trimmedFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'pathFileName' is an empty string!\n")

    return pathDir, fileNameExt, bothAreEmpty, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'pathFileName' consists of blank spaces!\n")

    return pathDir, fileNameExt, bothAreEmpty, err
  }

  xFnameExt, isEmpty, err2 := fh.GetFileNameWithExt(trimmedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFileNameWithExt(pathFileNameExt).\n"+
      "pathFileNameExt='%v'\nError='%v'\n",
      pathFileNameExt, err2.Error())
    return
  }

  if isEmpty {
    fileNameExt = ""
  } else {
    fileNameExt = xFnameExt
  }

  remainingPathStr := strings.TrimSuffix(trimmedFileNameExt, fileNameExt)

  if len(remainingPathStr) == 0 {
    pathDir = ""

    if pathDir == "" && fileNameExt == "" {
      bothAreEmpty = true
    } else {
      bothAreEmpty = false
    }

    err = nil

    return

  }

  xPath, isEmpty, err2 := fh.GetPathFromPathFileName(remainingPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetPathFromPathFileName(remainingPathStr).\n"+
      "remainingPathStr='%v'\nError='%v'\n",
      remainingPathStr, err2.Error())

    return
  }

  if isEmpty {
    pathDir = ""
  } else {
    pathDir = xPath
  }

  if pathDir == "" && fileNameExt == "" {
    bothAreEmpty = true
  } else {
    bothAreEmpty = false
  }

  err = nil

  return
}

// GetPathFromPathFileName - Returns the path from a path and file name string.
// If the returned path is an empty string, return parameter 'isEmpty' is set to
// 'true'.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileNameExt  string - This is an input parameter. The method expects to
//                            receive a single, properly formatted path and file
//                            name string delimited by dots ('.') and path Separators
//                            ('/' or '\'). On Windows the 'pathFileNameExt' string
//                            valid volume designations (Example: "D:")
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  path             string - This is the directory path extracted from the input parameter
//                            'pathFileNameExt'. If successful, the 'path' string that is returned
//                            by this method WILL NOT include a trailing path separator ('/' or '\'
//                            depending on the os). Example 'path': "./pathfile/003_filehelper"
//
//  isEmpty            bool - If the method determines that it cannot extract a valid directory
//                            path from input parameter 'pathFileNameExt', this boolean value
//                            will be set to 'true'. Failure to extract a valid directory path
//                            will occur if the input parameter 'pathFileNameExt' is not properly
//                            formatted as a valid path and file name.
//
//  err               error - If a processing error is detected, an error will be returned. Note that
//                            in the event that this method fails to extract a valid directory path
//                            'pathFileNameExt' due to the fact that 'pathFileNameExt' was improperly
//                            formatted, 'isEmpty' will be set to 'true', but no error will be returned.
//
//                            If no error occurs, 'err' is set to 'nil'.
//
// ------------------------------------------------------------------------
//
// Examples:
//
//  pathFileNameExt = ""                  returns isEmpty==true  err==nil
//  pathFileNameExt = "D:\"               returns "D:\"
//  pathFileNameExt = "."                 returns ".\"
//  pathFileNameExt = "..\"               returns "..\"
//  pathFileNameExt = "...\"              returns ERROR
//
//  pathFileNameExt = ".\pathfile\003_filehelper\wt_HowToRunTests.md"
//                                        returns ".\pathfile\003_filehelper"
//
//  pathFileNameExt = "someFile.go"       returns ""
//  pathFileNameExt = "..\dir1\dir2\.git" returns "..\dir1\dir2"
//                                         '.git' is assumed to be a file.
//
func (fh FileHelper) GetPathFromPathFileName(
  pathFileNameExt string) (dirPath string, isEmpty bool, err error) {

  ePrefix := "FileHelper.GetPathFromPathFileName() "
  dirPath = ""
  isEmpty = true
  err = nil
  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    err = errors.New(ePrefix +
      "Error: Input parameter 'pathFileNameExt' is an empty string!\n")

    return dirPath, isEmpty, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathFileNameExt' consists of blank spaces!\n")

    return dirPath, isEmpty, err
  }

  testPathStr, isDirEmpty, err2 := fh.CleanDirStr(pathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned by fh.CleanDirStr(pathFileNameExt).\n"+
      "pathFileNameExt='%v'\nError='%v'\n",
      pathFileNameExt, err2.Error())
    return dirPath, isEmpty, err
  }

  if isDirEmpty {
    dirPath = ""
    isEmpty = true
    err = nil
    return dirPath, isEmpty, err
  }

  lTestPathStr := len(testPathStr)

  if lTestPathStr == 0 {
    err = errors.New(ePrefix +
      "Error: AdjustPathSlash was applied to 'pathStr'.\n" +
      "The 'testPathStr' string is a Zero Length string!\n")
    return dirPath, isEmpty, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathStr).\n"+
      "testPathStr='%v'\nError='%v'\n",
      testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  lSlashIdxs := len(slashIdxs)

  firstGoodChar, lastGoodChar, err2 :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr("+
      "testPathStr).\n"+
      "testPathStr='%v'\nError='%v'\n",
      testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathStr).\n"+
      "testPathStr='%v'\nError='%v'\n",
      testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  lDotIdxs := len(dotIdxs)

  var finalPathStr string

  volName := fp.VolumeName(testPathStr)

  if testPathStr == volName {

    finalPathStr = testPathStr

  } else if strings.Contains(testPathStr, "...") {

    err = fmt.Errorf(ePrefix+
      "Error: PATH CONTAINS INVALID Dot Characters!\n"+
      "testPathStr='%v'\n", testPathStr)
    return dirPath, isEmpty, err

  } else if firstGoodChar == -1 || lastGoodChar == -1 {

    absPath, err2 := fh.MakeAbsolutePath(testPathStr)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error returned from fh.MakeAbsolutePath(testPathStr).\n"+
        "testPathStr='%v'\nError='%v'\n",
        testPathStr, err2.Error())

      return dirPath, isEmpty, err
    }

    if absPath == "" {
      err = fmt.Errorf(ePrefix+
        "Error: Could not convert 'testPathStr' to Absolute path!\n"+
        "testPathStr='%v'\n",
        testPathStr)
      return dirPath, isEmpty, err
    }

    finalPathStr = testPathStr

  } else if lSlashIdxs == 0 {
    // No path separators but alpha numeric chars are present
    dirPath = ""
    isEmpty = true
    err = nil
    return dirPath, isEmpty, err

  } else if lDotIdxs == 0 {
    //path separators are present but there are no dots in the string

    if slashIdxs[lSlashIdxs-1] == lTestPathStr-1 {
      // Trailing path separator
      finalPathStr = testPathStr[0:slashIdxs[lSlashIdxs-2]]
    } else {
      finalPathStr = testPathStr
    }

  } else if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] {
    // format: ./dir1/dir2/fileName.ext
    finalPathStr = testPathStr[0:slashIdxs[lSlashIdxs-1]]

  } else if dotIdxs[lDotIdxs-1] < slashIdxs[lSlashIdxs-1] {

    finalPathStr = testPathStr

  } else {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH STRING.\n"+
      "testPathStr='%v'\n", testPathStr)

    return dirPath, isEmpty, err
  }

  if len(finalPathStr) == 0 {
    err = fmt.Errorf(ePrefix + "Error: Processed path is a Zero Length String!\n")

    return dirPath, isEmpty, err
  }

  //Successfully isolated and returned a valid
  // directory path from 'pathFileNameExt'
  dirPath = finalPathStr

  if len(dirPath) == 0 {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil

  return dirPath, isEmpty, err
}

// GetPathSeparatorIndexesInPathStr - Returns an array containing the indexes of
// path Separators (Forward slashes or backward slashes depending on operating
// system).
func (fh FileHelper) GetPathSeparatorIndexesInPathStr(
  pathStr string) ([]int, error) {

  ePrefix := "FileHelper.GetPathSeparatorIndexesInPathStr() "
  errCode := 0
  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    return []int{},
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' is an empty string!\n")
  }

  if errCode == -2 {
    return []int{},
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' consists of blank spaces!\n")
  }

  var slashIdxs []int

  for i := 0; i < lPathStr; i++ {

    rChar := pathStr[i]

    if rChar == os.PathSeparator ||
      rChar == '\\' ||
      rChar == '/' {

      slashIdxs = append(slashIdxs, i)
    }

  }

  return slashIdxs, nil
}

// GetVolumeName - Returns the volume name of associated with
// a given directory path. The method calls the function
// 'path/filepath.VolumeName().
//
// VolumeName() returns the leading volume name if it exists
// in input parameter 'pathStr'.
//
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "c:\foo\bar" it returns "c:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share" on linux
// On other platforms, it returns "".
//
func (fh FileHelper) GetVolumeName(pathStr string) string {

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return ""
  }

  return fp.VolumeName(pathStr)
}

// GetVolumeNameIndex - Analyzes input parameter 'pathStr' to
// determine if it contains a volume name.
// The method calls the function 'path/filepath.VolumeName().
//
// VolumeName() returns the leading volume name if it exists
// in input parameter 'pathStr'.
//
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "c:\foo\bar" it returns "c:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share" on linux
// On other platforms, it returns "".
//
func (fh FileHelper) GetVolumeNameIndex(
  pathStr string) (volNameIndex int, volNameLength int, volNameStr string) {

  volNameIndex = -1
  volNameLength = 0
  volNameStr = ""

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return volNameIndex, volNameLength, volNameStr
  }

  volName := fp.VolumeName(pathStr)

  if len(volName) == 0 {
    return volNameIndex, volNameLength, volNameStr
  }

  volNameIndex = strings.Index(pathStr, volName)

  volNameLength = len(volName)

  if volNameLength > 0 {
    volNameStr = volName
  }

  return volNameIndex, volNameLength, volNameStr
}

// IsAbsolutePath - Compares the input parameter 'pathStr' to
// the absolute path representation for 'pathStr' to determine
// whether 'pathStr' represents an absolute path.
//
func (fh FileHelper) IsAbsolutePath(pathStr string) bool {

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return false
  }

  // Adjust the path separators for the current operating
  // system.
  correctDelimPathStr := strings.ToLower(fh.AdjustPathSlash(pathStr))

  absPath, err := fh.MakeAbsolutePath(pathStr)

  if err != nil {
    return false
  }

  absPath = strings.ToLower(absPath)

  if absPath == correctDelimPathStr {
    return true
  }

  return false
}

// IsPathFileString - Returns 'true' if the it is determined that
// input parameter, 'pathFileStr', represents a directory path,
// file name and optionally, a file extension.
//
// If 'pathFileStr' is judged to be a directory path and file name,
// by definition it cannot be solely a directory path.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileStr            string - The string to be analyzed.
//
// Return Values:
//
//  pathFileType PathFileTypeCode - Path File Type Code indicating whether the input parameter 'pathFileStr'
//                                  is a Path, a Path and File, a File or "Indeterminate". "Indeterminate"
//                                  signals that the nature of 'pathFileStr' cannot be classified as either
//                                  a Path or a Path and File or a File.
//
//                                  --------------------------------------------------------
//                                  PathFileTypeCodes
//                                     0 = None
//                                     1 = Path
//                                     2 = PathFile
//                                     3 = File (with no path)
//                                     4 = Volume
//                                     5 = Indeterminate - Cannot determine whether string is a Path,
//                                         Path & File or File
//
//
//  err                     error - If an error is encountered during processing, it is returned here. If no error
//                                  occurs, 'err' is set to 'nil'. An error will be triggered if the input parameter
//                                  'pathFileStr' cannot no alpa numeric characters.
//
func (fh FileHelper) IsPathFileString(
  pathFileStr string) (pathFileType PathFileTypeCode,
  absolutePathFile string,
  err error) {

  ePrefix := "FileHelper.IsPathFileString() "

  pathFileType = PathFileType.None()
  absolutePathFile = ""
  err = nil

  var pathFileDoesExist bool
  var fInfo FileInfoPlus
  var correctedPathFileStr string

  correctedPathFileStr,
    pathFileDoesExist,
    fInfo,
    err = fh.doesPathFileExist(pathFileStr,
    PreProcPathCode.PathSeparator(), // Convert Path Separators to os default Path Separators
    ePrefix,
    "pathFileStr")

  if err != nil {
    return pathFileType, absolutePathFile, err
  }

  if strings.Contains(correctedPathFileStr, "...") {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH STRING!\n"+
      "pathFileStr='%v'\n", correctedPathFileStr)

    return pathFileType, absolutePathFile, err
  }

  osPathSepStr := string(os.PathSeparator)

  if strings.Contains(correctedPathFileStr, osPathSepStr+osPathSepStr) {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error: Invalid Path File string.\n"+
      "Path File string contains invalid Path Separators.\n"+
      "pathFileStr='%v'\n",
      correctedPathFileStr)

    return pathFileType, absolutePathFile, err
  }

  testAbsPathFileStr, err2 := fh.MakeAbsolutePath(correctedPathFileStr)

  if err2 != nil {
    err = fmt.Errorf("Error converting pathFileStr to absolute path.\n"+
      "pathFileStr='%v'\nError='%v'\n",
      correctedPathFileStr, err2.Error())

    return pathFileType, absolutePathFile, err
  }

  volName := fp.VolumeName(testAbsPathFileStr)

  if strings.ToLower(volName) == strings.ToLower(testAbsPathFileStr) ||
    strings.ToLower(volName) == strings.ToLower(pathFileStr) {
    // This is a volume name not a file Name!
    pathFileType = PathFileType.Volume()
    absolutePathFile = volName
    err = nil
    return pathFileType, absolutePathFile, err
  }

  // See if path actually exists on disk and
  // then examine the File Info object returned.
  if pathFileDoesExist {

    if fInfo.IsDir() {

      pathFileType = PathFileType.Path()

      absolutePathFile = testAbsPathFileStr

      err = nil

    } else {

      pathFileType = PathFileType.PathFile()

      absolutePathFile = testAbsPathFileStr

      err = nil

    }

    return pathFileType, absolutePathFile, err
  } // End of if pathFileDoesExist

  // Ok - We know the testPathFileStr does NOT exist on disk

  firstCharIdx, lastCharIdx, err2 :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr"+
      "(correctedPathFileStr)\n"+
      "correctedPathFileStr='%v'\nError='%v'\n",
      correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  interiorDotPathIdx := strings.LastIndex(correctedPathFileStr, "."+string(os.PathSeparator))

  if interiorDotPathIdx > firstCharIdx {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. Invalid interior relative path detected!\n"+
      "correctedPathFileStr='%v'\n",
      correctedPathFileStr)
    return pathFileType, absolutePathFile, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetPathSeparatorIndexesInPathStr(correctedPathFileStr) returned error.\n"+
      "testAbsPathFileStr='%v'\nError='%v'\n",
      correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetDotSeparatorIndexesInPathStr(correctedPathFileStr) retured error.\n"+
      "correctedPathFileStr='%v'\nError='%v'\n",
      correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  lDotIdxs := len(dotIdxs)

  lSlashIdxs := len(slashIdxs)

  if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // Option # 1
    // There are no valid characters in the string
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "No valid characters in parameter 'pathFileStr'\n"+
      "pathFileStr='%v'\n", correctedPathFileStr)

  } else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 2
    // String consists only of eligible alphanumeric characters
    // "sometextstring"

    // the string has no dots and no path separators.
    absolutePathFile = testAbsPathFileStr

    // Call it a file!
    // Example: "somefilename"
    pathFileType = PathFileType.File()
    err = nil

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 3
    // There no dots no characters but string does contain
    // slashes

    if lSlashIdxs > 1 {
      pathFileType = PathFileType.None()
      absolutePathFile = ""
      err = fmt.Errorf(ePrefix+
        "Invalid parameter 'pathFileStr'!\n"+
        "'pathFileStr' consists entirely of path separators!\n"+
        "pathFileStr='%v'\n",
        correctedPathFileStr)

    } else {
      // lSlashIdxs must be '1'
      absolutePathFile = testAbsPathFileStr

      // Call it a Directory!
      // Example: "/"
      pathFileType = PathFileType.Path()
      err = nil
    }

  } else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
    // Option # 4
    // strings contains slashes and characters but no dots.

    absolutePathFile = testAbsPathFileStr

    if lastCharIdx > slashIdxs[lSlashIdxs-1] {
      // Example D:\dir1\dir2\xray
      // It could be a file or a path. We simply
      // can't tell.
      pathFileType = PathFileType.Indeterminate()
      err = nil

    } else {

      // Call it a Directory!
      // Example: "D:\dir1\dir2\"
      pathFileType = PathFileType.Path()
      err = nil

    }

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
    // Option # 5
    // dots only. Must be "." or ".."

    absolutePathFile = testAbsPathFileStr

    // Call it a Directory!
    // Example: "D:\dir1\dir2\"
    pathFileType = PathFileType.Path()
    err = nil

  } else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
    // Option # 6
    // Dots and characters only. No slashes.
    // Maybe 'fileName.ext'

    absolutePathFile = testAbsPathFileStr

    // Call it a file!
    // Example: "somefilename"
    pathFileType = PathFileType.File()
    err = nil

  } else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
    // Option # 7
    // Dots and slashes, but no characters.

    absolutePathFile = testAbsPathFileStr

    // Call it a Directory!
    // Example: ".\"
    pathFileType = PathFileType.Path()
    err = nil

  } else {
    // Option # 8
    // Must be else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
    // Contains dots slashes and characters

    absolutePathFile = testAbsPathFileStr

    // If there is a dot after the last path separator
    // and there is a character following the dot.
    // Therefore, this is a filename extension, NOT a directory
    if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
      dotIdxs[lDotIdxs-1] < lastCharIdx {

      // Call it a path/file!
      // Example: "D:\dir1\dir2\somefilename.txt"
      pathFileType = PathFileType.PathFile()
      err = nil

    } else if lastCharIdx > slashIdxs[lSlashIdxs-1] &&
      dotIdxs[lDotIdxs-1] < slashIdxs[lSlashIdxs-1] {
      // Example: "..\dir1\dir2\sometext"
      // Could be a file could be a directory
      // Cannot determine.
      pathFileType = PathFileType.Indeterminate()
      err = nil

    } else {

      // Call it a Directory!
      // Example: "D:\dir1\dir2\"
      pathFileType = PathFileType.Path()
      err = nil
    }

  }

  return pathFileType, absolutePathFile, err
}

// IsPathString - Attempts to determine whether a string is a
// path string designating a directory (and not a path file name
// file extension string).
//
// If the path exists on disk, this method will examine the
// associated file information and render a definitive and
// accurate determination as to whether the path string represents
// a directory.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathStr string       - The path string to be analyzed.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  isPathStr       bool - If the input parameter, 'pathStr'
//                         is determined to be a directory
//                         path, this return value is set to
//                         true. Here, a 'directory path' is defined
//                         as a true directory and the path does NOT
//                         contain a file name.
//
//  cannotDetermine bool - If the method cannot determine whether
//                         the input parameter 'pathStr' is or
//                         is NOT a valid directory path, this
//                         this return value will be set to 'true'.
//                         The 'cannotDetermine=true' condition occurs
//                         with path names like 'D:\DirA\common'. The
//                         cannot determine whether 'common' is a file
//                         name or a directory name.
//
//
//  testPathStr   string - Input parameter 'pathStr' is subjected to cleaning routines
//                         designed to exclude extraneous characters from the analysis.
//                         'testPathFileStr' is the actual string on which the analysis was
//                         performed.
//
//  err            error - If an error occurs this return value will
//                         be populated. If no errors occur, this return
//                         value is set to nil.
//
//
func (fh FileHelper) IsPathString(
  pathStr string) (isPathStr bool, cannotDetermine bool, testPathStr string, err error) {

  ePrefix := "FileHelper.IsPathString() "
  testPathStr = ""
  isPathStr = false
  cannotDetermine = false
  err = nil

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' is an empty string!\n")
    return isPathStr, cannotDetermine, testPathStr, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' consists of blank spaces!\n")

    return isPathStr, cannotDetermine, testPathStr, err
  }

  if strings.Contains(pathStr, "...") {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH STRING!\n"+
      "pathStr='%v'\n", pathStr)
    return isPathStr, cannotDetermine, testPathStr, err
  }

  testPathStr = fh.AdjustPathSlash(pathStr)

  pathFileType, _, err2 := fh.IsPathFileString(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return isPathStr, cannotDetermine, testPathStr, err
  }

  if pathFileType == PathFileType.Path() {
    isPathStr = true
    cannotDetermine = false
    err = nil
    return isPathStr, cannotDetermine, testPathStr, err
  }

  if pathFileType == PathFileType.Indeterminate() {
    isPathStr = false
    cannotDetermine = true
    err = nil
    return isPathStr, cannotDetermine, testPathStr, err
  }

  isPathStr = false
  cannotDetermine = false
  err = nil
  return isPathStr, cannotDetermine, testPathStr, err
}

// JoinPathsAdjustSeparators - Joins two
// path strings and standardizes the
// path separators according to the
// current operating system.
func (fh FileHelper) JoinPathsAdjustSeparators(
  p1 string, p2 string) string {

  errCode := 0

  errCode, _, p1 = fh.isStringEmptyOrBlank(p1)

  if errCode < 0 {
    p1 = ""
  }

  errCode, _, p2 = fh.isStringEmptyOrBlank(p2)

  if errCode < 0 {
    p2 = ""
  }

  if p1 == "" &&
    p2 == "" {

    return ""
  }

  ps1 := fh.AdjustPathSlash(fp.Clean(p1))
  ps2 := fh.AdjustPathSlash(fp.Clean(p2))
  return fp.Clean(fh.AdjustPathSlash(path.Join(ps1, ps2)))

}

// JoinPaths - correctly joins 2-paths. Like the method JoinPathsAdjustSeparators()
// this method also converts path separators to the correct path separators for
// the current operating system.
//
func (fh FileHelper) JoinPaths(p1 string, p2 string) string {

  return fh.JoinPathsAdjustSeparators(p1, p2)

}

// MakeAbsolutePath - Supply a relative path or any path
// string and resolve that path to an Absolute path.
// Note: Clean() is called on result by fp.Abs().
func (fh FileHelper) MakeAbsolutePath(relPath string) (string, error) {

  ePrefix := "FileHelper.MakeAbsolutePath() "

  errCode := 0

  errCode, _, relPath = fh.isStringEmptyOrBlank(relPath)

  if errCode == -1 {
    return "",
      errors.New(ePrefix +
        "Error: Input parameter 'relPath' is an empty string!\n")
  }

  if errCode == -2 {
    return "",
      errors.New(ePrefix +
        "Error: Input parameter 'relPath' consists of blank spaces!\n")
  }

  testRelPath := fh.AdjustPathSlash(relPath)

  errCode, _, testRelPath = fh.isStringEmptyOrBlank(testRelPath)

  if errCode < 0 {
    return "", errors.New(ePrefix +
      "Error: Input Parameter 'relPath' adjusted for path Separators is an EMPTY string!\n")
  }

  p, err := fp.Abs(testRelPath)

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Error returned from fp.Abs(testRelPath).\n"+
      "testRelPath='%v'\nError='%v'\n",
      testRelPath, err.Error())
  }

  return p, nil
}

// MakeDirAll - creates a directory named path, along with any necessary
// parent directories. In other words, all directories in the path are
// created.
//
// The permission bits 'drwxrwxrwx' are used for all directories that the
// method creates.
//
// If path is a directory which already exists, this method does nothing
// and returns and error value of 'nil'.
//
// Note that this method calls FileHelper.MakeDirAllPerm()
//
func (fh FileHelper) MakeDirAll(dirPath string) error {
  ePrefix := "FileHelper.MakeDirAll() "
  permission, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  err = fh.MakeDirAllPerm(dirPath, permission)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return nil
}

// MakeDir - Creates a single directory. The method returns an error
// type. If the operation succeeds, the error value is 'nil'. If the
// operation fails the error value is populated with an appropriate
// error message.
//
// This method will fail if the parent directory does not exist.
//
// The permission bits 'drwxrwxrwx' are used for directory creation.
// If path is already a directory, this method does nothing and returns
// an error value of 'nil'.
//
// Note that this method calls FileHelper.MakeDirPerm().
//
func (fh FileHelper) MakeDir(dirPath string) error {

  ePrefix := "FileHelper.MakeDir() "

  permission, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  err = fh.MakeDirPerm(dirPath, permission)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  return nil
}

// MakeDirAllPerm - Creates a directory path along with any necessary
// parent paths.
//
// If the target directory path already exists, this method does nothing
// and returns.
//
// The input parameter 'permission' is of type 'FilePermissionConfig'.
// See method the documentation for method 'FilePermissionConfig.New()'
// for an explanation of permission codes.
//
// If you wish to grant total access to a directory, consider setting
// permission code as follows:
//     FilePermissionConfig{}.New("drwxrwxrwx")
//
// If the parent directories in parameter 'dirPath' do not yet exist, this
// method will create them.
//
func (fh FileHelper) MakeDirAllPerm(dirPath string, permission FilePermissionConfig) error {

  ePrefix := "FileHelper.MakeDirAllPerm() "

  errCode := 0

  errCode, _, dirPath = fh.isStringEmptyOrBlank(dirPath)

  if errCode == -1 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dirPath' is an empty string!\n")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dirPath' consists of blank spaces!\n")
  }

  err2 := permission.IsValid()

  if err2 != nil {
    return fmt.Errorf(ePrefix+"Input parameter 'permission' is INVALID!\n"+
      "Error='%v'\n", err2.Error())
  }

  dirPermCode, err2 := permission.GetCompositePermissionMode()

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "ERROR: INVALID Permission Code\n"+
      "Error='%v'\n", err2.Error())
  }

  dirPath, err2 = fh.MakeAbsolutePath(dirPath)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(dirPath).\n"+
      "dirPath='%v'\nError='%v'\n",
      dirPath, err2.Error())
  }

  err2 = os.MkdirAll(dirPath, dirPermCode)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error return from os.MkdirAll(dirPath, permission).\n"+
      "dirPath='%v'\nError='%v'\n",
      dirPath, err2.Error())
  }

  var pathDoesExist bool

  _,
    pathDoesExist,
    _,
    err2 = fh.doesPathFileExist(dirPath,
    PreProcPathCode.None(), // Take no Pre-Processing Action
    ePrefix,
    "dirPath")

  if err2 != nil {
    return err2
  }

  if !pathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST!\n"+
      "dirPath='%v'\n", dirPath)
  }

  return nil
}

// MakeDirPerm - Creates a single directory using the permission codes passed by input
// parameter 'permission'.
//
// This method will fail if the parent directory does not exist. To create all parent
// directories in the path use method 'FileHelper.MakeDirAllPerm()'.
//
//
// The input parameter 'permission' is of type 'FilePermissionConfig'.
// See method the documentation for method 'FilePermissionConfig.New()'
// for an explanation of permission codes.
//
// If you wish to grant total access to a directory, consider setting
// permission code as follows:
//     FilePermissionConfig{}.New("drwxrwxrwx")
//
// An error will be triggered if the 'dirPath' input parameter represents
// an invalid path or if parent directories in the path do not exist.
//
func (fh FileHelper) MakeDirPerm(dirPath string, permission FilePermissionConfig) error {

  ePrefix := "FileHelper.MakeDirPerm() "

  errCode := 0

  errCode, _, dirPath = fh.isStringEmptyOrBlank(dirPath)

  if errCode == -1 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dirPath' is an empty string!\n")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dirPath' consists of blank spaces!\n")
  }

  err2 := permission.IsValid()

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Input parameter 'permission' is INVALID!\n"+
      "Error='%v'\n", err2.Error())
  }

  dirPermCode, err2 := permission.GetCompositePermissionMode()

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "ERROR: INVALID Permission Code\n"+
      "Error='%v'\n", err2.Error())
  }

  dirPath, err2 = fh.MakeAbsolutePath(dirPath)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(dirPath).\n"+
      "dirPath='%v'\nError='%v'\n", dirPath, err2.Error())
  }

  err2 = os.Mkdir(dirPath, dirPermCode)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error return from os.Mkdir(dirPath, dirPermCode).\n"+
      "dirPath='%v'\nError='%v'\n",
      dirPath, err2.Error())
  }

  var pathDoesExist bool
  _,
    pathDoesExist,
    _,
    err2 = fh.doesPathFileExist(dirPath,
    PreProcPathCode.None(), // Take no Pre-Processing action
    ePrefix,
    "dirPath")

  if err2 != nil {
    return err2
  }

  if !pathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST!\n"+
      "dirPath='%v'\n", dirPath)
  }

  return nil
}

// MoveFile - Copies file from source to destination and, if successful,
// then deletes the original source file.
//
// The copy procedure will carried out using the the 'Copy By Io' technique.
// See FileHelper.CopyFileByIo().
//
// The 'move' operation will create the destination file, but it will NOT
// create the destination directory. If the destination directory does NOT
// exist, an error will be returned.
//
// If this copy operation fails, the method will return an error and it
// will NOT delete the source file.
//
// If an error is encountered during this procedure it will be returned by
// means of the return parameter 'err'.
//
func (fh FileHelper) MoveFile(src, dst string) error {

  ePrefix := "FileHelper.MoveFile() "

  var err error
  var srcFileDoesExist, dstFileDoesExist bool
  var srcFInfo, dstFInfo FileInfoPlus

  src,
    srcFileDoesExist,
    srcFInfo,
    err = fh.doesPathFileExist(src,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "src")

  if err != nil {
    return err
  }

  if !srcFileDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter "+
      "'src' file DOES NOT EXIST!\n"+
      "src='%v'\n", src)
  }

  if srcFInfo.IsDir() {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'src' "+
      "exists and it is directory NOT a file!\n"+
      "src='%v'\n", src)
  }

  dst,
    dstFileDoesExist,
    dstFInfo,
    err = fh.doesPathFileExist(dst,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "dst")

  if err != nil {
    return err
  }

  if dstFileDoesExist && dstFInfo.IsDir() {
    return fmt.Errorf(ePrefix+"Error: Input parameter 'dst' does exist,\n"+
      "but it a 'directory' and NOT a File!\n"+
      "dst='%v'\n", dst)
  }

  // ============================
  // Perform the copy operation!
  // Use Copy By IO Procedure
  // ============================
  err = fh.CopyFileByIo(src, dst)

  if err != nil {
    // Copy Operation Failed. Return an error
    // and DO NOT delete the source file!
    return fmt.Errorf(ePrefix+
      "Error: Copy operation FAILED!\n"+
      "Source File='%v'\n"+
      "Destination File='%v'\nError='%v'\n",
      src, dst, err.Error())
  }

  // CopyFileByIo operation was apparently successful.
  // Now, verify that destination file exists.

  _,
    dstFileDoesExist,
    _,
    err = fh.doesPathFileExist(dst,
    PreProcPathCode.None(), // Take no Pre-Processing action
    ePrefix,
    "dst")

  if err != nil {
    return err
  }

  if !dstFileDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: After Copy Operation, destination file "+
      "DOES NOT EXIST!\n"+
      "Therefore, the copy operation FAILED! Source file was NOT deleted.\n"+
      "destination file='%v'\n", dst)
  }

  // Successful copy operation has been verified.
  // Time to delete the source file.
  err = os.Remove(src)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Copy operation succeeded, but attempted deletion of source file FAILED!\n"+
      "Source File='%v'\n", src)
  }

  _,
    srcFileDoesExist,
    _,
    err = fh.doesPathFileExist(src,
    PreProcPathCode.None(), // Take No Pre-Processing Action
    ePrefix,
    "src")

  if err != nil {
    return err
  }

  if srcFileDoesExist {
    return fmt.Errorf("Verification Error: File 'src' still exists!\n"+
      "src='%v'\n", src)
  }

  // Success, source was copied to destination
  // AND the source file was deleted.

  // Done and we are out of here!
  return nil
}

// OpenDirectory - Opens a directory and returns the associated 'os.File' pointer.
// This method will open a directory designated by input parameter, 'directoryPath'.
//
// The input parameter 'createDir' determines the action taken if 'directoryPath'
// does not exist. If 'createDir' is set to 'true' and 'directoryPath' does not
// currently exist, this method will attempt to create 'directoryPath'. Directories
// created in this manner are configured with Open Type of 'Read-Write' and a
// Permission code of 'drwxrwxrwx'.
//
// Alternatively, if 'createDir' is set to 'false' and 'directoryPath' does NOT exist,
// an error will be returned.
//
// Regardless of whether the target directory path already exists or is created by
// this method, the returned os.File pointer is opened with the 'Read-Only' attribute
// (O_RDONLY) and a permission code of zero ("----------").
//
// Note: The caller is responsible for calling "Close()" on the returned os.File pointer.
//
// --------------------------------------------------------------------------------------------------------
//
// Input Parameters:
//
//
//  directoryPath                  string - A string containing the path name of the directory
//                                          which will be opened.
//
//
//  createDir                        bool - Determines what action will be taken if 'directoryPath'
//                                          does NOT exist. If 'createDir' is set to 'true' and
//                                          'directoryPath' does NOT exist, this method will attempt
//                                          to create 'directoryPath'. Alternatively, if 'createDir'
//                                          is set to false and 'directoryPath' does NOT exist, this
//                                          method will terminate and an error will be returned.
//
//                                          Directories created in this manner will have an Open Type
//                                          of 'Read-Write' and a Permission code of 'drwxrwxrwx'. This
//                                          differs from the Open Type and permission mode represented
//                                          by the returned os.File pointer.
//
// --------------------------------------------------------------------------------------------------------
//
// Return Values:
//
//  *os.File        - If successful, this method returns an os.File pointer
//                    to the directory designated by input parameter 'directoryPath'.
//
//                    If successful, the returned os.File pointer is opened with the
//                    'Read-Only' attribute (O_RDONLY) and a permission code of zero
//                    ("----------").
//
//                    If this method fails, the *os.File return value is 'nil'.
//
//                    Note: The caller is responsible for calling "Close()" on this
//                    os.File pointer.
//
//
//  error           - If the method completes successfully, the error return value
//                    is 'nil'. If the method fails, the error type returned is
//                    populated with an appropriate error message.
//
func (fh FileHelper) OpenDirectory(
  directoryPath string,
  createDir bool) (*os.File, error) {

  ePrefix := "FileHelper.OpenDirectory() "
  var err error
  var directoryPathDoesExist bool
  var dirPathFInfo FileInfoPlus

  directoryPath,
    directoryPathDoesExist,
    dirPathFInfo,
    err = fh.doesPathFileExist(directoryPath,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "directoryPath")

  if err != nil {
    return nil, err
  }

  if directoryPathDoesExist &&
    !dirPathFInfo.IsDir() {

    return nil,
      fmt.Errorf(ePrefix+
        "ERROR: 'directoryPath' does exist, but\n"+
        "IT IS NOT A DIRECTORY!\n"+
        "directoryPath='%v'\n", directoryPath)
  }

  if directoryPathDoesExist && dirPathFInfo.Mode().IsRegular() {
    return nil,
      fmt.Errorf(ePrefix+
        "ERROR: 'directoryPath' does exist, but\n"+
        "it is classifed as a REGULAR File!\n"+
        "directoryPath='%v'\n", directoryPath)
  }

  if !directoryPathDoesExist {

    if !createDir {
      return nil,
        fmt.Errorf(ePrefix+"Error 'directoryPath' DOES NOT EXIST!\n"+
          "directoryPath='%v'\n", directoryPath)
    }

    // Parameter 'createDir' must be 'true'.
    // The error signaled that the path does not exist. So, create the directory path
    err = fh.MakeDirAll(directoryPath)

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix+"ERROR: Attmpted creation of 'directoryPath' FAILED!\n"+
          "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
    }

    // Verify that the directory exists and get
    // the associated file info object.
    _,
      directoryPathDoesExist,
      dirPathFInfo,
      err = fh.doesPathFileExist(directoryPath,
      PreProcPathCode.None(), // Take No Pre-Processing Action
      ePrefix,
      "directoryPath")

    if err != nil {
      return nil, fmt.Errorf(ePrefix+"Error occurred verifying existance of "+
        "newly created 'directoryPath'!\n"+
        "Non-Path error returned by os.Stat(directoryPath)\ndirectoryPath='%v'\nError='%v'\n",
        directoryPath, err.Error())
    }

    if !directoryPathDoesExist {
      return nil, fmt.Errorf(ePrefix+"Error: Verification of newly created "+
        "directoryPath FAILED!\n"+
        "'directoryPath' DOES NOT EXIST!\n"+
        "directoryPath='%v'\n", directoryPath)
    }

    if !dirPathFInfo.IsDir() {
      return nil,
        fmt.Errorf(ePrefix+"ERROR: Input Paramter 'directoryPath' is NOT a directory!\n"+
          "directoryPath='%v'\n", directoryPath)
    }

    if dirPathFInfo.Mode().IsRegular() {
      return nil,
        fmt.Errorf(ePrefix+
          "ERROR: 'directoryPath' does exist, but\n"+
          "it is classifed as a REGULAR File!\n"+
          "directoryPath='%v'\n", directoryPath)
    }
  }

  filePtr, err := os.Open(directoryPath)

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+"File Open Error: %v\n"+
        "directoryPath='%v'\n",
        err.Error(), directoryPath)
  }

  if filePtr == nil {
    return nil, errors.New(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!\n")
  }

  return filePtr, nil
}

// OpenFile - wrapper for os.OpenFile. This method may be used to open or
// create files depending on the File Open and File Permission parameters.
//
// If successful, this method will return a pointer to the os.File object
// associated with the file designated for opening.
//
// The calling routine is responsible for calling "Close()" on this os.File
// pointer.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  pathFileName                   string - A string containing the path and file name
//                                          of the file which will be opened. If a parent
//                                          path component does NOT exist, this method will
//                                          trigger an error.
//
//  fileOpenCfg            FileOpenConfig - This parameter encapsulates the File Open parameters
//                                          which will be used to open subject file. For an
//                                          explanation of File Open parameters, see method
//                                          FileOpenConfig.New().
//
// filePermissionCfg FilePermissionConfig - This parameter encapsulates the File Permission
//                                          parameters which will be used to open the subject
//                                          file. For an explanation of File Permission parameters,
//                                          see method FilePermissionConfig.New().
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  *os.File        - If successful, this method returns an os.File pointer
//                    to the file designated by input parameter 'pathFileName'.
//                    This file pointer can subsequently be used for reading
//                    content from the subject file. It may NOT be used for
//                    writing content to the subject file.
//
//                    If this method fails, the *os.File return value is 'nil'.
//
//                    Note: The caller is responsible for calling "Close()" on this
//                    os.File pointer.
//
//
//  error           - If the method completes successfully, this return value
//                    is 'nil'. If the method fails, the error type returned
//                    is populated with an appropriate error message.
//
func (fh FileHelper) OpenFile(
  pathFileName string,
  fileOpenCfg FileOpenConfig,
  filePermissionCfg FilePermissionConfig) (filePtr *os.File, err error) {

  filePtr = nil
  err = nil
  ePrefix := "FileHelper.OpenFile() "

  var pathFileNameDoesExist bool
  var fInfoPlus FileInfoPlus

  pathFileName,
    pathFileNameDoesExist,
    fInfoPlus,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return nil, err
  }

  if !pathFileNameDoesExist {
    err = fmt.Errorf(ePrefix+
      "ERROR: Input parameter 'pathFileName' DOES NOT EXIST!\n"+
      "pathFileName='%v'\n", pathFileName)

    return filePtr, err
  }

  if fInfoPlus.IsDir() {
    err =
      fmt.Errorf(ePrefix+
        "ERROR: Input parameter 'pathFileName' is "+
        "a 'Directory' - NOT a file!\n"+
        "pathFileName='%v'\n", pathFileName)

    return filePtr, err
  }

  err2 := fileOpenCfg.IsValid()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Input Parameter 'fileOpenCfg' is INVALID!\n"+
      "Error='%v'\n", err2.Error())
    return filePtr, err
  }

  fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return filePtr, err
  }

  err2 = filePermissionCfg.IsValid()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Input Parameter 'filePermissionCfg' is INVALID!\n"+
      "Error='%v'\n", err2.Error())
    return filePtr, err
  }

  fileMode, err2 := filePermissionCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return filePtr, err
  }

  filePtr, err2 = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by os.OpenFile(pathFileName, fOpenCode, fileMode).\n"+
      "pathFileName='%v'\nError='%v'\n", pathFileName, err2.Error())

    return filePtr, err
  }

  if filePtr == nil {
    err = errors.New(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!\n")
    return filePtr, err
  }

  err = nil

  return filePtr, err
}

// OpenFileReadOnly - Opens the designated path file name for reading
// only.
//
// If successful, this method returns a pointer of type *os.File which
// can only be used for reading reading content from the subject file.
// This file pointer is configured for 'Read-Only' operations. You may
// not write to the subject file using this pointer.
//
// If the designated file ('pathFileName') does NOT exist, an error
// will be triggered.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileName        string - A string containing the path and file name
//                               of the file which will be opened in the
//                               'Read-Only' mode. If the path or file does
//                               NOT exist, this method will trigger an error.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  *os.File        - If successful, this method returns an os.File pointer
//                    to the file designated by input parameter 'pathFileName'.
//                    This file pointer can subsequently be used for reading
//                    content from the subject file. It may NOT be used for
//                    writing content to the subject file.
//
//                    If this method fails, the *os.File return value is 'nil'.
//
//                    Note: The caller is responsible for calling "Close()" on this
//                    os.File pointer.
//
//
//  error           - If the method completes successfully, this return value
//                    is 'nil'. If the method fails, the error type returned
//                    is populated with an appropriate error message.
//
func (fh FileHelper) OpenFileReadOnly(pathFileName string) (filePtr *os.File, err error) {

  ePrefix := "FileHelper.OpenFileReadOnly() "

  filePtr = nil
  err = nil
  var err2 error
  var pathFileNameDoesExist bool
  var fInfoPlus FileInfoPlus

  pathFileName,
    pathFileNameDoesExist,
    fInfoPlus,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return filePtr, err
  }

  if !pathFileNameDoesExist {
    err = fmt.Errorf(ePrefix+
      "ERROR: The input parameter 'pathFileName' DOES NOT EXIST!\n"+
      "pathFileName='%v'\n", pathFileName)
    return filePtr, err
  }

  if fInfoPlus.IsDir() {
    err = fmt.Errorf(ePrefix+
      "ERROR: The input parameter 'pathFileName' is a 'Directory' "+
      "and NOT a path file name.\n"+
      "'pathFileName' is therefore INVALID!\n"+
      "pathFileName='%v'\n", pathFileName)

    return filePtr, err
  }

  fileOpenCfg, err2 := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadOnly(),"+
        "FOpenMode.ModeNone()).\n"+
        "Error='%v'\n",
        err2.Error())
    return filePtr, err
  }

  fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error Creating File Open Code.\n"+
      "Error=%v\n",
      err2.Error())
    return filePtr, err
  }

  fPermCfg, err2 := FilePermissionConfig{}.New("-r--r--r--")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix+
        "Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n"+
        "Error='%v'\n", err2.Error())
    return filePtr, err
  }

  fileMode, err2 := fPermCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error Creating File Mode Code.\n"+
      "Error=%v\n",
      err2.Error())

    return filePtr, err
  }

  filePtr, err2 = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"File Open Error: %v\n"+
      "pathFileName='%v'", err2.Error(), pathFileName)
    filePtr = nil
    return filePtr, err
  }

  if filePtr == nil {
    err = fmt.Errorf(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!\n")

    return filePtr, err
  }

  err = nil

  return filePtr, err
}

// OpenFileReadWrite - Opens the file designated by input parameter
// 'fileName' for 'Writing'. The actual permission code used to open
// the file is 'Read/Write'.
//
// If the method is successful, a pointer to the opened file is returned
// along with an error value of 'nil'.
//
// If the file does not exist, this method will attempt to create it.
//
// If the file path does not exist, an error will be triggered.
//
// If the method completes successfully, the caller is responsible for
// call "Close()" on the returned os.File pointer.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileName        string - A string containing the path and file name
//                               of the file which will be opened in the
//                               'Read/Write' mode. If the file does NOT
//                               exist, this method will attempt to create
//                               it. However, if the path component of
//                               'pathFileName' does not exist, an error
//                                will be returned.
//
//  truncateFile          bool - If set to 'true' and the target file will
//                               be truncated to zero bytes in length before
//                               it is opened.
//
//                               If set to 'false', the target file will be
//                               be opened in the 'Append' mode and any bytes
//                               written to the file will be appended to the
//                               end of the file. Under this scenario, the
//                               original file contents are preserved and newly
//                               written bytes are added to the end of the file.
//
//                               If the file designated by input parameter 'pathFileName'
//                               does not exist, this parameter ('truncateFile') is
//                               ignored and the new created file is initialized
//                               containing zero bytes.
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  *os.File        - If successful, this method returns an os.File pointer
//                    to the file opened for 'Read/Write' operations. This
//                    file pointer can be used subsequently for writing
//                    content to, or reading content from, the subject file.
//
//                    If this method fails, this return value is 'nil'.
//
//                    Note: The caller is responsible for calling "Close()" on this
//                    os.File pointer.
//
//
//  error           - If the method completes successfully, this return value
//                    is 'nil'. If the method fails, the error type returned
//                    is populated with an appropriate error message.
//
func (fh FileHelper) OpenFileReadWrite(
  pathFileName string,
  truncateFile bool) (*os.File, error) {

  ePrefix := "FileHelper.OpenFileReadWrite() "

  var fPtr *os.File
  var err error
  var pathFileNameDoesExist bool
  var fInfoPlus FileInfoPlus

  pathFileName,
    pathFileNameDoesExist,
    fInfoPlus,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return nil, err
  }

  var fileOpenCfg FileOpenConfig

  if !pathFileNameDoesExist {
    // pathFileName does NOT exist

    fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
      FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix+
          "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),"+
          "FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\n"+
          "Error='%v'\n",
          err.Error())
    }

  } else {
    // pathFileName does exist

    if fInfoPlus.IsDir() {
      return nil,
        fmt.Errorf(ePrefix+"ERROR: Input parameter 'pathFileName' is "+
          "a 'Directory' NOT a file.\n"+
          "pathFileName='%v'\n", pathFileName)
    }

    if truncateFile {
      // truncateFile == true
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
        FOpenMode.ModeTruncate())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix+
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite(),"+
            "FOpenMode.ModeTruncate()).\n"+
            "Error='%v'\n",
            err.Error())
      }

    } else {
      // truncateFile == false
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
        FOpenMode.ModeAppend())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix+
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite(),"+
            "FOpenMode.ModeAppend()).\n"+
            "Error='%v'\n",
            err.Error())
      }
    }
  }

  fOpenCode, err := fileOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+
        "Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n"+
        "Error='%v'\n", err.Error())
  }

  fileMode, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return nil, fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err != nil {
    return nil, fmt.Errorf(ePrefix+"File Open Error\n"+
      "pathFileName='%v'\nError='%v'\n",
      pathFileName, err.Error())
  }

  if fPtr == nil {
    return nil, fmt.Errorf(ePrefix+
      "ERROR: os.OpenFile() returned a 'nil' file pointer!\n"+
      "pathFileName='%v'\n", pathFileName)
  }

  return fPtr, nil
}

// OpenFileWriteOnly - Opens a file for 'Write-Only' operations. Input parameter
// 'pathFileName' specifies the the path and file name of the file which will be
// opened.
//
// If the path component of 'pathFileName' does not exist, an error will be returned.
//
// If the designated file does not exist, this method will attempt to create the file.
//
// If the method completes successfully, the caller is responsible for calling 'Close()'
// on the returned os.File pointer.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  pathFileName        string - A string containing the path and file name
//                               of the file which will be opened in the
//                               'Write Only' mode. If the file does NOT
//                               exist, this method will attempt to create
//                               it. However, if the path component of
//                               'pathFileName' does not exist, an error
//                                will be returned.
//
//  truncateFile          bool - If set to 'true' and the target file will
//                               be truncated to zero bytes in length before
//                               it is opened.
//
//                               If set to 'false', the target file will be
//                               be opened in the 'Append' mode and any bytes
//                               written to the file will be appended to the
//                               end of the file. Under this scenario, the
//                               original file contents are preserved and newly
//                               written bytes are added to the end of the file.
//
//                               If the file designated by input parameter 'pathFileName'
//                               does not exist, this parameter ('truncateFile') is
//                               ignored and the newly created file is initialized
//                               with zero bytes of content.
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  *os.File        - If successful, this method returns an os.File pointer
//                    to the file opened for 'Write Only' operations. This
//                    file pointer can be used for writing content to the
//                    subject file.
//
//                    If this method fails, this return value is 'nil'.
//
//                    Note: The caller is responsible for calling "Close()" on this
//                    os.File pointer.
//
//
//  error           - If the method completes successfully, this return value
//                    is 'nil'. If the method fails, the error type returned
//                    is populated with an appropriate error message.
//
func (fh FileHelper) OpenFileWriteOnly(
  pathFileName string,
  truncateFile bool) (*os.File, error) {

  ePrefix := "FileHelper.OpenFileWriteOnly() "

  var fPtr *os.File
  var err error
  var pathFileNameDoesExist bool
  var fInfoPlus FileInfoPlus

  pathFileName,
    pathFileNameDoesExist,
    fInfoPlus,
    err = fh.doesPathFileExist(
    pathFileName,
    PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
    ePrefix,
    "pathFileName")

  if err != nil {
    return nil, err
  }

  var fileOpenCfg FileOpenConfig

  if !pathFileNameDoesExist {
    // The pathFileName DOES NOT EXIST!

    fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
      FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix+
          "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),"+
          "FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\nError='%v'\n",
          err.Error())
    }

  } else {
    // The pathFileName DOES EXIST!

    if fInfoPlus.IsDir() {
      return nil,
        fmt.Errorf(ePrefix+"ERROR: Input parameter 'pathFileName' is "+
          "a 'Directory' NOT a file.\n"+
          "pathFileName='%v'\n", pathFileName)
    }

    if truncateFile {
      // truncateFile == true; Set Mode 'Truncate'
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
        FOpenMode.ModeTruncate())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix+
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),"+
            "FOpenMode.ModeTruncate()).\nError='%v'\n",
            err.Error())
      }

    } else {
      // truncateFile == false; Set Mode 'Append'
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
        FOpenMode.ModeAppend())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix+
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),"+
            "FOpenMode.ModeAppend()).\nError='%v'\n",
            err.Error())
      }
    }
  }

  fOpenCode, err := fileOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+
        "Error creating File Open Code.\nError=%v\n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("--wx-wx-wx")

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+
        "Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n"+
        "Error='%v' \n", err.Error())
  }

  fileMode, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return nil, fmt.Errorf(ePrefix+
      "Error creating file mode code.\nError=%v\n", err.Error())
  }

  fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err != nil {
    return nil, fmt.Errorf(ePrefix+
      "Error returned from os.OpenFile().\n"+
      "pathFileName='%v'\nError='%v'\n",
      pathFileName, err.Error())
  }

  if fPtr == nil {
    return nil, fmt.Errorf(ePrefix +
      "ERROR: File pointer returned from os.OpenFile() is 'nil'!\n")
  }

  return fPtr, nil
}

// RemovePathSeparatorFromEndOfPathString - Remove Trailing path Separator from
// a path string - if said trailing path Separator exists.
//
func (fh FileHelper) RemovePathSeparatorFromEndOfPathString(pathStr string) string {
  errCode := 0
  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return ""

  }

  lastChar := rune(pathStr[lPathStr-1])

  if lastChar == os.PathSeparator ||
    lastChar == '\\' ||
    lastChar == '/' {

    return pathStr[0 : lPathStr-1]
  }

  return pathStr
}

// SearchFileModeMatch - This method determines whether the file mode of the file described by input
// parameter, 'info', is match for the File Selection Criteria 'fileSelectCriteria.SelectByFileMode'.
// If the file's FileMode matches the 'fileSelectCriteria.SelectByFileMode' value, the return value,
// 'isFileModeMatch' is set to 'true'.
//
// If 'fileSelectCriteria.SelectByFileMode' is NOT initialized to a valid File Mode value, the
// return value 'isFileModeSet' is set to 'false' signaling the File Mode File Selection Criterion
// is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus'
// object for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func (fh *FileHelper) SearchFileModeMatch(
  info os.FileInfo,
  fileSelectCriteria FileSelectionCriteria) (isFileModeSet, isFileModeMatch bool, err error) {

  if fileSelectCriteria.SelectByFileMode.isInitialized == false {
    isFileModeSet = false
    isFileModeMatch = false
    err = nil
    return isFileModeSet, isFileModeMatch, err
  }

  selectFileMode, err2 := fileSelectCriteria.SelectByFileMode.GetFileMode()

  if err2 != nil {
    ePrefix := "FileHelper.SearchFileModeMatch() "

    err = fmt.Errorf(ePrefix+"SelectByFileMode is INVALID!\n"+
      "Error='%v'\n", err2.Error())
    return isFileModeSet, isFileModeMatch, err
  }

  if selectFileMode == info.Mode() {
    isFileModeSet = true
    isFileModeMatch = true
    err = nil
    return isFileModeSet, isFileModeMatch, err

  }

  isFileModeSet = true
  isFileModeMatch = false
  err = nil
  return isFileModeSet, isFileModeMatch, err
}

// SearchFileNewerThan - This method is called to determine whether the file described by the
// input parameter 'info' is a 'match' for the File Selection Criteria, 'fileSelectCriteria.FilesNewerThan'.
// If the file modification date time occurs after the 'fileSelectCriteria.FilesNewerThan' date time,
// the return value 'isFileNewerThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesNewerThan' is set to time.Time zero ( the default or zero value for this type),
// the return value 'isFileNewerThanSet' is set to 'false' signaling that this search criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func (fh *FileHelper) SearchFileNewerThan(
  info os.FileInfo,
  fileSelectCriteria FileSelectionCriteria) (isFileNewerThanSet, isFileNewerThanMatch bool, err error) {

  isFileNewerThanSet = false
  isFileNewerThanMatch = false
  err = nil

  if fileSelectCriteria.FilesNewerThan.IsZero() {
    isFileNewerThanSet = false
    isFileNewerThanMatch = false
    err = nil
    return
  }

  if fileSelectCriteria.FilesNewerThan.Before(info.ModTime()) {
    isFileNewerThanSet = true
    isFileNewerThanMatch = true
    err = nil
    return

  }

  isFileNewerThanSet = true
  isFileNewerThanMatch = false
  err = nil

  return
}

// SearchFileOlderThan - This method is called to determine whether the file described by the
// input parameter 'info' is a 'match' for the File Selection Criteria, 'fileSelectCriteria.FilesOlderThan'.
// If the file modification date time occurs before the 'fileSelectCriteria.FilesOlderThan' date time,
// the return value 'isFileOlderThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesOlderThan' is set to time.Time zero ( the default or zero value for this type),
// the return value 'isFileOlderThanSet' is set to 'false' signaling that this search criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func (fh *FileHelper) SearchFileOlderThan(
  info os.FileInfo,
  fileSelectCriteria FileSelectionCriteria) (isFileOlderThanSet,
  isFileOlderThanMatch bool,
  err error) {

  if fileSelectCriteria.FilesOlderThan.IsZero() {
    isFileOlderThanSet = false
    isFileOlderThanMatch = false
    err = nil
    return
  }

  if fileSelectCriteria.FilesOlderThan.After(info.ModTime()) {
    isFileOlderThanSet = true
    isFileOlderThanMatch = true
    err = nil
    return
  }

  isFileOlderThanSet = true
  isFileOlderThanMatch = false
  err = nil
  return

}

// SearchFilePatternMatch - used to determine whether a file described by the
// 'info' parameter meets the specified File Selection Criteria and is judged
// to be a match for the fileSelectCriteria FileNamePattern. 'fileSelectCriteria.FileNamePatterns'
// consists of a string array. If the pattern signified by any element in the string array
// is a 'match', the return value 'isPatternMatch' is set to true.
//
// If the 'fileSelectCriteria.FileNamePatterns' array is empty or if it contains only empty strings,
// the return value isPatternSet is set to 'false' signaling that the pattern file search selection
// criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func (fh *FileHelper) SearchFilePatternMatch(
  info os.FileInfo,
  fileSelectCriteria FileSelectionCriteria) (isPatternSet, isPatternMatch bool, err error) {

  ePrefix := "DirMgr.SearchFilePatternMatch()"

  isPatternMatch = false
  isPatternSet = false
  err = nil

  isPatternSet = fileSelectCriteria.ArePatternsActive()

  if !isPatternSet {
    isPatternSet = false
    isPatternMatch = false
    err = nil
    return
  }

  lPats := len(fileSelectCriteria.FileNamePatterns)

  for i := 0; i < lPats; i++ {

    matched, err2 := fp.Match(fileSelectCriteria.FileNamePatterns[i], info.Name())

    if err2 != nil {
      isPatternSet = true
      err = fmt.Errorf(ePrefix+
        "Error returned from filepath.Match(fileSelectCriteria.FileNamePatterns[i], "+
        "info.Name())\nfileSelectCriteria.FileNamePatterns[i]='%v'\ninfo.Name()='%v'\nError='%v'\n",
        fileSelectCriteria.FileNamePatterns[i], info.Name(), err2.Error())
      isPatternMatch = false
      return
    }

    if matched {
      isPatternSet = true
      isPatternMatch = true
      err = nil
      return
    }
  }

  isPatternSet = true
  isPatternMatch = false
  err = nil
  return
}

// SwapBasePath - Searches the 'targetPath' string for the existence of
// 'oldBasePath'. If 'oldBasePath' is found, it is replaced with 'newBasePath'.
//
// If 'oldBasePath' is not found in 'targetPath' an error is returned.
//
// Likewise, if 'oldBasePath' is not located at the beginning of 'targetPath',
// an error will be returned.
//
func (fh FileHelper) SwapBasePath(
  oldBasePath,
  newBasePath,
  targetPath string) (string, error) {

  ePrefix := "FileHelper.SwapBasePath() "

  errCode := 0
  oldBaseLen := 0

  errCode, oldBaseLen, oldBasePath = fh.isStringEmptyOrBlank(oldBasePath)

  if errCode == -1 {
    return "",
      errors.New(ePrefix + "Input parameter 'oldBasePath' is an empty string!")

  }

  if errCode == -2 {
    return "", errors.New(ePrefix +
      "Input parameter 'oldBasePath' consists of all spaces!")
  }

  errCode, _, newBasePath = fh.isStringEmptyOrBlank(newBasePath)

  if errCode == -1 {
    return "",
      errors.New(ePrefix + "Input parameter 'newBasePath' is an empty string!")

  }

  if errCode == -2 {
    return "", errors.New(ePrefix +
      "Input parameter 'newBasePath' consists of all spaces!")
  }

  targetPathLen := 0
  errCode, targetPathLen, targetPath = fh.isStringEmptyOrBlank(targetPath)

  if errCode == -1 {
    return "",
      errors.New(ePrefix + "Input parameter 'targetPath' is an empty string!")

  }

  if errCode == -2 {
    return "", errors.New(ePrefix +
      "Input parameter 'targetPath' consists of all spaces!")
  }

  if oldBaseLen > targetPathLen {
    return "", errors.New(ePrefix +
      "Length of input parameter 'oldBasePath' is greater than length of \n" +
      "input parameter 'targetPath'.\n")
  }

  idx := strings.Index(
    strings.ToLower(targetPath),
    strings.ToLower(oldBasePath))

  if idx < 0 {
    return "",
      fmt.Errorf(ePrefix+
        "Error: Could not locate 'oldBasePath' in 'targetPath'. "+
        "oldBasePath='%v' targetPath='%v' ",
        oldBasePath, targetPath)
  }

  if idx != 0 {
    return "",
      fmt.Errorf(ePrefix+
        "Error: 'oldBasePath' is NOT located at the beginning of 'targetPath'. "+
        "oldBasePath='%v' targetPath='%v' ",
        oldBasePath, targetPath)
  }

  return newBasePath + targetPath[oldBaseLen:], nil
}

/*
  FileHelper private methods
*/

// doesPathFileExist - A helper method which public methods use to determine whether a
// path and file does or does not exist.
//
// This method calls os.Stat(dirPath) which returns an error which is one of two types:
//
//     1. A Non-Path Error - An error which is not path related. It signals some other type
//        of error which makes impossible to determine if the path actually exists. These
//        types of errors generally relate to "access denied" situations, but there may be
//        other reasons behind non-path errors. If a non-path error is returned, no valid
//        existence test can be performed on the file path.
//
//            or
//
//     2. A Path Error - indicates that the path definitely does not exist.
//
// To deal with these types of errors, this method will test path existence up to three times before
// returning a non-path error.
//
func (fh FileHelper) doesPathFileExist(
  filePath string,
  preProcessCode PreProcessPathCode,
  errorPrefix string,
  filePathTitle string) (absFilePath string,
  filePathDoesExist bool,
  fInfo FileInfoPlus,
  nonPathError error) {

  ePrefix := "FileHelper.doesDirPathExist() "

  absFilePath = ""
  filePathDoesExist = false
  fInfo = FileInfoPlus{}
  nonPathError = nil

  if len(errorPrefix) > 0 {
    ePrefix = errorPrefix
  }

  if len(filePathTitle) == 0 {
    filePathTitle = "filePath"
  }

  errCode := 0

  errCode, _, filePath = fh.isStringEmptyOrBlank(filePath)

  if errCode == -1 {
    nonPathError = fmt.Errorf(ePrefix+
      "Error: Input parameter '%v' is an empty string!", filePathTitle)
    return absFilePath, filePathDoesExist, fInfo, nonPathError
  }

  if errCode == -2 {
    nonPathError = fmt.Errorf(ePrefix+
      "Error: Input parameter '%v' consists of blank spaces!", filePathTitle)
    return absFilePath, filePathDoesExist, fInfo, nonPathError
  }

  var err error

  if preProcessCode == PreProcPathCode.PathSeparator() {

    absFilePath = fh.AdjustPathSlash(filePath)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    absFilePath, err = fh.MakeAbsolutePath(filePath)

    if err != nil {
      absFilePath = ""
      nonPathError = fmt.Errorf(ePrefix+"fh.MakeAbsolutePath() FAILED!\n"+
        "%v", err.Error())
      return absFilePath, filePathDoesExist, fInfo, nonPathError
    }
  } else {
    // For any other PreProcPathCode value, apply no pre-processing to
    absFilePath = filePath
  }

  var info os.FileInfo

  for i := 0; i < 3; i++ {

    filePathDoesExist = false
    fInfo = FileInfoPlus{}
    nonPathError = nil

    info, err = os.Stat(absFilePath)

    if err != nil {

      if os.IsNotExist(err) {

        filePathDoesExist = false
        fInfo = FileInfoPlus{}
        nonPathError = nil
        return absFilePath, filePathDoesExist, fInfo, nonPathError
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be test
      // up to 3-times before it is returned.
      nonPathError = fmt.Errorf(ePrefix+"Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n",
        filePathTitle, filePathTitle, filePath, err.Error())
      fInfo = FileInfoPlus{}
      filePathDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      filePathDoesExist = true
      nonPathError = nil
      fInfo = FileInfoPlus{}.NewFromFileInfo(info)
      return absFilePath, filePathDoesExist, fInfo, nonPathError
    }

    time.Sleep(30 * time.Millisecond)
  }

  return absFilePath, filePathDoesExist, fInfo, nonPathError
}

// isStringEmptyOrBlank - Analyzes a string to determine if the string is 'empty' or
// if the string consists of all blanks (spaces).
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//     int - Integer Values Returned:
//           -1 = string is empty
//           -2 = string consists entirely of spaces
//            0 = string contains non-space characters
//
//  string - The original input parameter string ('testStr') from which the
//           leading and trailing spaces have been deleted.
//           Examples:
//
//              1. 'testStr'     = "  a string   "
//                 return string = "a string"
//
//              2. 'testStr'     = "    "
//                 return string = ""
//
//              3. 'testStr'     = ""
//                 return string = ""
//
func (fh FileHelper) isStringEmptyOrBlank(
  testStr string) (errCode int, strLen int, newStr string) {

  errCode = 0
  strLen = 0
  newStr = ""

  if len(testStr) == 0 {
    errCode = -1
    return errCode, strLen, newStr
  }

  newStr = strings.TrimLeft(testStr, " ")

  newStr = strings.TrimRight(newStr, " ")

  strLen = len(newStr)

  if strLen == 0 {
    errCode = -2
    newStr = ""
    return errCode, strLen, newStr
  }

  errCode = 0

  return errCode, strLen, newStr
}

// makeFileHelperWalkDirDeleteFilesFunc - Used in conjunction with DirMgr.DeleteWalDirFiles
// to select and delete files residing the directory tree identified by the current DirMgr
// object.
//
func (fh *FileHelper) makeFileHelperWalkDirDeleteFilesFunc(dInfo *DirectoryDeleteFileInfo) func(string, os.FileInfo, error) error {
  return func(pathFile string, info os.FileInfo, erIn error) error {

    ePrefix := "FileHelper.makeFileHelperWalkDirDeleteFilesFunc()"

    if erIn != nil {
      dInfo.ErrReturns = append(dInfo.ErrReturns, erIn)
      return nil
    }

    if info.IsDir() {

      subDir, err := DirMgr{}.New(pathFile)

      if err != nil {

        ex := fmt.Errorf(ePrefix+
          "Error returned from DirMgr{}.New(pathFile).\n"+
          "pathFile:='%v'\nError='%v'\n", pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

        return nil
      }

      subDir.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)
      dInfo.Directories.AddDirMgr(subDir)

      return nil
    }

    isFoundFile, err := fh.FilterFileName(info, dInfo.DeleteFileSelectCriteria)

    if err != nil {

      ex := fmt.Errorf(ePrefix+
        "Error returned from fh.FilterFileName(info, dInfo.DeleteFileSelectCriteria)\n" +
        "pathFile='%v'\n" +
        "info.Name()='%v'\nError='%v'\n",
        pathFile, info.Name(), err.Error())

      dInfo.ErrReturns = append(dInfo.ErrReturns, ex)
      return nil
    }

    if isFoundFile {

      err := os.Remove(pathFile)

      if err != nil {
        ex := fmt.Errorf(ePrefix+
          "Error returned from os.Remove(pathFile). pathFile='%v' Error='%v'",
          pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex)
        return nil
      }

      err = dInfo.DeletedFiles.AddFileMgrByFileInfo(pathFile, info)

      if err != nil {
        ex := fmt.Errorf(ePrefix+
          "Error returned from dInfo.DeletedFiles.AddFileMgrByFileInfo( pathFile,  info). "+
          "pathFile='%v'  Error='%v'",
          pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex)
        return nil
      }

    }

    return nil
  }
}

// makeFileHelperWalkDirFindFilesFunc - This function is designed to work in conjunction
// with a walk directory function like FindWalkDirFiles. It will process
// files extracted from a 'Directory Walk' operation initiated by the 'filepath.Walk' method.
func (fh *FileHelper) makeFileHelperWalkDirFindFilesFunc(dInfo *DirectoryTreeInfo) func(string, os.FileInfo, error) error {
  return func(pathFile string, info os.FileInfo, erIn error) error {

    ePrefix := "DirMgr.makeFileHelperWalkDirFindFilesFunc() "

    if erIn != nil {
      ex2 := fmt.Errorf(ePrefix+"Error returned from directory walk function. "+
        "pathFile= '%v' Error='%v'", pathFile, erIn.Error())
      dInfo.ErrReturns = append(dInfo.ErrReturns, ex2)
      return nil
    }

    if info.IsDir() {
      subDir, err := DirMgr{}.NewFromFileInfo(pathFile, info)

      if err != nil {

        er2 := fmt.Errorf(ePrefix+"Error returned by DirMgr{}.New(pathFile). "+
          "pathFile='%v' Error='%v'", pathFile, err.Error())
        dInfo.ErrReturns = append(dInfo.ErrReturns, er2)
        return nil
      }

      dInfo.Directories.AddDirMgr(subDir)
      return nil
    }

    fh := FileHelper{}

    // This is not a directory. It is a file.
    // Determine if it matches the find file criteria.
    isFoundFile, err := fh.FilterFileName(info, dInfo.FileSelectCriteria)

    if err != nil {

      er2 := fmt.Errorf(ePrefix+
        "Error returned from dMgr.FilterFileName(info, " +
        "dInfo.FileSelectCriteria)\n" +
        "pathFile='%v'\ninfo.Name()='%v'\nError='%v'\n",
        pathFile, info.Name(), err.Error())

      dInfo.ErrReturns = append(dInfo.ErrReturns, er2)
      return nil
    }

    if isFoundFile {

      fMgr, err2 := FileMgr{}.NewFromPathFileNameExtStr(pathFile)

      if err2 != nil {
        err = fmt.Errorf(ePrefix+
          "Error returned by FileMgr{}.NewFromPathFileNameExtStr(pathFile) "+
          "pathFile='%v' Error='%v' ", pathFile, err2.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, err)

        return nil
      }

      err = dInfo.FoundFiles.AddFileMgrByFileInfo(fMgr.dMgr.GetAbsolutePath(), info)

      if err != nil {
        er2 := fmt.Errorf(ePrefix+
          "Error returned from dInfo.FoundFiles.AddFileMgrByFileInfo(pathFile, info)\n"+
          "pathFile='%v'\ninfo.Name()='%v'\nError='%v'\n",
          pathFile, info.Name(), err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, er2)

        return nil
      }
    }

    return nil
  }
}
