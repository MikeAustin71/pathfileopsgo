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

    The Source Repository for this source code file is :
      https://github.com/MikeAustin71/pathfilego.git

    'FileHelper' is a dependency of 'DirMgr' and 'FileMgr'.  'DirMgr' and 'FileMgr'
    are located in source file 'filemanager.go' found in this same
    directory: '003_filehelper/common/filemanager.go'


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
// Before calling this method, it would be wise to ensure that both input parameters
// are using the correct os.PathSeparator character. If the path file strings are
// using the incorrect path separators, an error will be triggered.
//
func (fh FileHelper) AreSameFile(pathFile1, pathFile2 string) (bool, error) {

  ePrefix := "FileHelper.AreSameFile() "

  errCode := 0

  errCode, _, pathFile1 = fh.isStringEmptyOrBlank(pathFile1)

  isEmptyStr1 := false

  if errCode < 0 {
    isEmptyStr1 = true
  }

  isEmptyStr2 := false

  errCode, _, pathFile2 = fh.isStringEmptyOrBlank(pathFile2)

  if errCode < 0 {
    isEmptyStr2 = true
  }

  if isEmptyStr1 && !isEmptyStr2 {
    return false, errors.New(ePrefix + "Error: First Path File String is EMPTY and INVALID!")
  }

  if isEmptyStr2 && !isEmptyStr1 {
    return false, errors.New(ePrefix +
      "Error: Second Path File String is EMPTY and INVALID!")
  }

  if isEmptyStr1 && isEmptyStr2 {
    return false, errors.New(ePrefix +
      "Error: Both Path File Strings are EMPTY and INVALID!")
  }

  correctedPathFile1 := strings.ToLower(fh.AdjustPathSlash(pathFile1))

  correctedPathFile2 := strings.ToLower(fh.AdjustPathSlash(pathFile2))

  lcPathFile1 := strings.ToLower(pathFile1)
  lcPathFile2 := strings.ToLower(pathFile2)

  if correctedPathFile1 != lcPathFile1 &&
    correctedPathFile2 != lcPathFile2 {

    return false, fmt.Errorf(ePrefix+
      "Error: Both input parameters 'pathFile1' and 'pathFile2' contain INVALID "+
      "path separators.\npathFile1='%v'\npathFile2='%v'\n",
      pathFile1, pathFile2)
  }

  if correctedPathFile1 != lcPathFile1 {

    return false, fmt.Errorf(ePrefix+
      "Error: Input parameter 'pathFile1' contains INVALID "+
      "path separators.\npathFile1='%v'\n", pathFile1)
  }

  if correctedPathFile2 != lcPathFile2 {

    return false, fmt.Errorf(ePrefix+
      "Error: Input parameter 'pathFile2' contains INVALID "+
      "path separators.\npathFile2='%v'\n", pathFile2)
  }

  str1Exists := false

  f1, err := os.Stat(pathFile1)

  if err == nil {
    str1Exists = true
  }

  str2Exists := false

  f2, err := os.Stat(pathFile2)

  if err == nil {
    str2Exists = true
  }

  if str1Exists && str2Exists {

    if os.SameFile(f1, f2) {
      // pathFile1 and pathFile2 are the same
      // path and file name.

      return true, nil

    } else {

      return false, nil
    }

  }

  if str1Exists != str2Exists {

    return false, nil
  }

  // Both pathFile1 and pathFile2 do NOT exist

  absPathFile1, err := fh.MakeAbsolutePath(pathFile1)

  if err != nil {
    return false,
      fmt.Errorf(ePrefix+
        "Error: %v", err.Error())
  }

  absPathFile1 = strings.ToLower(absPathFile1)

  absPathFile2, err := fh.MakeAbsolutePath(pathFile2)

  if err != nil {
    return false,
      fmt.Errorf(ePrefix+
        "Error: %v", err.Error())
  }

  absPathFile2 = strings.ToLower(absPathFile2)

  if absPathFile1 == absPathFile2 {
    return true, nil
  }

  return false, nil
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
//     dirName = '../dir1/dir2/fileName.ext' returns "../dir1/dir2"
//
//     dirName = 'fileName.ext' returns "" isEmpty = true
//
//     dirName = '../dir1/dir2/' returns '../dir1/dir2'
//
//     dirName = '../dir1/dir2/filename.ext' returns '../dir1/dir2'
//
func (fh FileHelper) CleanDirStr(dirNameStr string) (dirName string, isEmpty bool, err error) {

  ePrefix := "FileHelper.CleanDirStr() "
  dirName = ""
  isEmpty = true
  err = nil

  errCode := 0

  errCode, _, dirNameStr = fh.isStringEmptyOrBlank(dirNameStr)

  if errCode == -1 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'dirNameStr' is an empty string!")
  }

  if errCode == -2 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'dirNameStr' consists of blank spaces!")
  }

  adjustedDirName := fh.AdjustPathSlash(dirNameStr)

  lAdjustedDirName := len(adjustedDirName)

  if lAdjustedDirName == 0 {
    err = errors.New(ePrefix +
      "Error: After adjusting for path separators, input parameter 'dirNameStr' is an empty string!")
    return
  }

  if strings.Contains(adjustedDirName, "...") {
    err = fmt.Errorf(ePrefix+
      "Error: Invalid Directory string. Contains invalid dots. adjustedDirName='%v' ",
      adjustedDirName)
    return
  }

  volName := fp.VolumeName(adjustedDirName)

  if volName == adjustedDirName {
    dirName = adjustedDirName
    isEmpty = false
    err = nil
    return
  }

  // Find out if the directory path
  // actually exists.
  fInfo, err2 := os.Stat(adjustedDirName)

  if err2 == nil {
    // The path exists

    if fInfo.IsDir() {
      // The path exists and it is a directory
      if adjustedDirName[lAdjustedDirName-1] == os.PathSeparator {
        dirName = adjustedDirName[0 : lAdjustedDirName-1]
      } else {
        dirName = adjustedDirName
      }

      if len(dirName) == 0 {
        isEmpty = true
      } else {
        isEmpty = false
      }

      err = nil
      return

    } else {
      // The path exists but it is
      // a File Name and NOT a directory name.
      adjustedDirName = strings.TrimSuffix(adjustedDirName, fInfo.Name())
      lAdjustedDirName = len(adjustedDirName)

      if lAdjustedDirName < 1 {
        dirName = ""
        isEmpty = true
        err = nil
        return
      }

      if adjustedDirName[lAdjustedDirName-1] == os.PathSeparator {
        dirName = adjustedDirName[0 : lAdjustedDirName-1]
      } else {

        dirName = adjustedDirName
      }

      if len(dirName) == 0 {
        isEmpty = true
      } else {
        isEmpty = false
      }

      err = nil
      return
    }
  }

  firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName). "+
      "adjustedDirName='%v'  Error='%v'", adjustedDirName, err2.Error())
    return
  }

  if firstCharIdx == -1 || lastCharIdx == -1 {
    if adjustedDirName[lAdjustedDirName-1] == os.PathSeparator {
      dirName = adjustedDirName[0 : lAdjustedDirName-1]
    } else {
      dirName = adjustedDirName
    }

    isEmpty = false
    err = nil
    return
  }

  interiorDotPathIdx := strings.LastIndex(adjustedDirName, "."+string(os.PathSeparator))

  if interiorDotPathIdx > firstCharIdx {
    err = fmt.Errorf(ePrefix+
      "Error: INVALID PATH. Invalid interior relative path detected! adjustedDirName='%v'",
      adjustedDirName)
    return
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf("Error returned by fh.GetPathSeparatorIndexesInPathStr(adjustedDirName). "+
      "adjusteDirName='%v'  Error='%v'", adjustedDirName, err2.Error())
    return
  }

  lSlashIdxs := len(slashIdxs)

  if lSlashIdxs == 0 {
    dirName = adjustedDirName
    isEmpty = false
    err = nil
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(adjustedDirName)

  if err2 != nil {
    err = fmt.Errorf("Error returned by fh.GetDotSeparatorIndexesInPathStr(adjustedDirName). "+
      "adjustedDirName='%v'  Error='%v'",
      adjustedDirName, err2.Error())
    return
  }

  lDotIdxs := len(dotIdxs)

  // If a path separator is the last character
  if slashIdxs[lSlashIdxs-1] == lAdjustedDirName-1 {
    dirName = adjustedDirName[0:slashIdxs[lSlashIdxs-1]]
    if len(dirName) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return
  }

  // If there is a dot after the last path separator,
  // this is a filename extension, NOT a directory
  if lDotIdxs > 0 && dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] {

    dirName = adjustedDirName[0:slashIdxs[lSlashIdxs-1]]

    if len(dirName) == 0 {
      isEmpty = true
    } else {
      isEmpty = false
    }

    err = nil
    return
  }

  dirName = adjustedDirName
  isEmpty = false
  err = nil
  return
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
//     fileNameExt = '../dir1/dir2/'
//                    returns "" and isEmpty=true
//
//     fileNameExt = '../filesfortest/newfilesfortest/newerFileForTest_01'
//                   returns "newerFileForTest_01" and isEmpty=false
//
//     fileNameExt = '../filesfortest/newfilesfortest/.gitignore'
//                   returns ".gitignore" and isEmpty=false
//
func (fh FileHelper) CleanFileNameExtStr(fileNameExtStr string) (fileNameExt string, isEmpty bool, err error) {

  ePrefix := "FileHelper.CleanFileNameExtStr() "
  fileNameExt = ""
  isEmpty = true
  err = nil

  errCode := 0

  errCode, _, fileNameExtStr = fh.isStringEmptyOrBlank(fileNameExtStr)

  if errCode == -1 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'fileNameExtStr' is an empty string!")
  }

  if errCode == -2 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'fileNameExtStr' consists of blank spaces!")
  }

  adjustedFileNameExt := fh.AdjustPathSlash(fileNameExtStr)

  if strings.Contains(adjustedFileNameExt, "...") {
    err = fmt.Errorf(ePrefix+"Error: Invalid Directory string. Contains invalid dots. adjustedFileNameExt='%v' ", adjustedFileNameExt)
    return
  }

  // Find out if the file name extension path
  // actually exists.
  fInfo, err2 := os.Stat(adjustedFileNameExt)

  if err2 == nil {
    // The path exists

    if fInfo.IsDir() {
      // The path exists and it is a directory.
      // There is no File Name present.
      fileNameExt = ""
      isEmpty = true
      err = fmt.Errorf(ePrefix+"Error: adjustedFileNameExt exists as a 'Directory' - NOT A FILE NAME! adjustedFileNameExt='%v'", adjustedFileNameExt)
      return
    } else {
      // The path exists and it is a valid
      // file name.
      fileNameExt = fInfo.Name()
      isEmpty = false

      err = nil
      return
    }
  }

  firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedFileNameExt)

  if firstCharIdx == -1 || lastCharIdx == -1 {
    err = fmt.Errorf(ePrefix+"File Name Extension string contains no valid file name characters! adjustedFileNameExt='%v'", adjustedFileNameExt)
    return
  }

  // The file name extension path does not exist

  interiorDotPathIdx := strings.LastIndex(adjustedFileNameExt, "."+string(os.PathSeparator))

  if interiorDotPathIdx > firstCharIdx {
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH. Invalid interior relative path detected! adjustedFileNameExt='%v'", adjustedFileNameExt)
    return
  }

  slashIdxs, err := fh.GetPathSeparatorIndexesInPathStr(adjustedFileNameExt)

  if err != nil {
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetPathSeparatorIndexesInPathStr(adjustedFileNameExt). adustedFileNameExt='%v'  Error='%v'", adjustedFileNameExt, err.Error())
    return
  }

  lSlashIdxs := len(slashIdxs)

  if lSlashIdxs == 0 {
    fileNameExt = adjustedFileNameExt
    isEmpty = false
    err = nil
    return
  }

  if lastCharIdx < slashIdxs[lSlashIdxs-1] {
    // Example: ../dir1/dir2/
    fileNameExt = ""
    isEmpty = true
    err = nil
    return
  }

  result := adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]

  fileNameExt = result

  if len(result) == 0 {
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
    err = fmt.Errorf(ePrefix+"%v", errX)
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
    err = fmt.Errorf(ePrefix+"%v", errX)
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

  errCode := 0

  errCode, _, src = fh.isStringEmptyOrBlank(src)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'src' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'src' consists of blank spaces!")
  }

  errCode, _, dst = fh.isStringEmptyOrBlank(dst)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'dst' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'dst' consists of blank spaces!")
  }

  correctedSrc, err2 := fh.MakeAbsolutePath(src)

  if err2 != nil {
    err = fmt.Errorf(ePrefix +
      "Error from fh.MakeAbsolutePath(src).\nsrc='%v'\nError='%v'",
      src, err2.Error())
    return err
  }

  correctedDest, err2 := fh.MakeAbsolutePath(dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix +
      "Error from fh.MakeAbsolutePath(dst).\nsrc='%v'\nError='%v'",
      dst, err2.Error())
    return err
  }

  areSameFile, err2 := fh.AreSameFile(correctedSrc, correctedDest)

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error occurred during path file name comparison.\n" +
      "Source File:'%v'\nDestination File:'%v'\nError='%v' ",
      correctedSrc, correctedDest, err2.Error())
    return err
  }

  if areSameFile {
    err = fmt.Errorf(ePrefix + "Error: The source and destination file are the same - equivalent.\n" +
      "Source File:'%v'\nDestination File:'%v'\n",
      correctedSrc, correctedDest)
    return err
  }

  sfi, err2 := os.Stat(correctedSrc)

  if err2!=nil {
    err = fmt.Errorf(ePrefix+
      "Error: Input parameter 'src' file DOES NOT EXIST! src='%v'\n" +
      "os.Stat(src) Error='%v'\n", src, err2.Error())
    return err
  }


  if !sfi.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+
      "Error: non-regular source file. Source File Name='%v'  Source File Mode='%v' ",
      sfi.Name(), sfi.Mode().String())
    return err
  }

  _, err2 = os.Stat(correctedDest)

  // If the destination file does NOT exist - this is not a problem
  // because the destination file will be created later.

  if err2 == nil {
    // The destination file exists. This IS a problem. Link will
    // fail when attempting to create a link to an existing file.

    err2 = os.Remove(correctedDest)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error: The target destination file exists and could NOT be deleted! \n"+
        "destination file='%v' Error='%v' ", correctedDest, err2.Error())
      return err
    }

    _, err2 = os.Stat(correctedDest)

    if err2 == nil {
      err = fmt.Errorf(ePrefix+"Error: Deletion of preexisting destination file failed! \n"+
        "The copy link operation cannot proceed! \n"+
        "destination file='%v' ", correctedDest)
      return err
    }

  }

  err2 = os.Link(correctedSrc, correctedDest)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"- os.Link(correctedSrc, correctedDest) FAILED!\n" +
      "src='%v' dst='%v'  Error='%v'", correctedSrc, correctedDest, err2.Error())
    return err
  }

  err = nil

  return err
}

// CopyFileByIo - Copies file from source path and file name
// to destination path and file name.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this method
// does NOT rely on the creation of symbolic links. Instead,
// a new destination file is created and the contents of the source
// file are written to the new destination file using "io.Copy()".
//
// "io.Copy()" is the only method used to copy the designated source
// file. If this method fails, an error is returned.
//
// If source file is equivalent to the destination file, no action will
// be taken and no error will be returned.
//
func (fh FileHelper) CopyFileByIo(src, dst string) (err error) {
  ePrefix := "FileHelper.CopyFileByIo() "
  err = nil
  errCode := 0

  errCode, _, src = fh.isStringEmptyOrBlank(src)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'src' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'src' consists of blank spaces!")
  }

  errCode, _, dst = fh.isStringEmptyOrBlank(dst)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'dst' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'dst' consists of blank spaces!")
  }

  correctedSrc, err2 := fh.MakeAbsolutePath(src)

  if err2 != nil {
    err = fmt.Errorf(ePrefix +
      "Error from fh.MakeAbsolutePath(src).\nsrc='%v'\nError='%v'",
      src, err2.Error())
    return err
  }

  correctedDest, err2 := fh.MakeAbsolutePath(dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix +
      "Error from fh.MakeAbsolutePath(dst).\nsrc='%v'\nError='%v'",
      dst, err2.Error())
    return err
  }

  areSameFile, err2 := fh.AreSameFile(correctedSrc, correctedDest)

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error occurred during path file name comparison.\n" +
      "Source File:'%v'\nDestination File:'%v'\nError='%v'\n",
      correctedSrc, correctedDest, err2.Error())
    return err
  }

  if areSameFile {
    err = fmt.Errorf(ePrefix + "Error: The source and destination file are the same - equivalent.\n" +
      "Source File:'%v'\nDestination File:'%v'\n",
      correctedSrc, correctedDest)
    return err
  }

  sfi, err2 := os.Stat(correctedSrc)

  if err2 != nil {

    err = fmt.Errorf(ePrefix+
      "Error: Source File is NOT Valid! Error returned from os.Stat(src). \n"+
      "src='%v'  Error='%v'\n", correctedSrc, err2.Error())
    return err
  }

  if !sfi.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+"Error non-regular source file ='%v' source file Mode='%v'",
      sfi.Name(), sfi.Mode().String())
    return err
  }

  dfi, err2 := os.Stat(correctedDest)

  // If the destination file does NOT exist, this is not a problem
  // since it will be created later. If the destination Path does
  // not exist, an error return will be triggered.

  if err2 == nil {
    // The destination file already exists!

    if !dfi.Mode().IsRegular() {
      err = fmt.Errorf(ePrefix+
        "Error: non-regular destination file. Cannot Overwrite destination file. "+
        "Destination file='%v' destination file mode='%v'",
        dfi.Name(), dfi.Mode().String())
      return err
    }

  }

  // Create a new destination file and copy source
  // file contents to the destination file.

  in, err2 := os.Open(correctedSrc)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Open(src) src='%v'  Error='%v'",
      correctedSrc, err2.Error())
    return err
  }

  out, err2 := os.Create(correctedDest)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Create(destinationFile)\n"+
      "destinationFile='%v'\nError='%v'\n",
      correctedDest, err2.Error())

    _ = in.Close()

    return err
  }

  if _, err2 = io.Copy(out, in); err2 != nil {
    _ = in.Close()
    _ = out.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from io.Copy(destination, source) \ndestination='%v'\n"+
      "source='%v'\nError='%v'\n",
      correctedDest, correctedSrc, err2.Error())
    return
  }

  // flush file buffers in memory
  err2 = out.Sync()

  if err2 != nil {
    _ = in.Close()
    _ = out.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from out.Sync()\nout=destination='%v'\nError='%v'\n",
      correctedDest, err2.Error())
    return
  }

  err2 = in.Close()

  if err2 != nil {
    _ = out.Close()

    err = fmt.Errorf(ePrefix+
      "Error returned from in.Close()\nin=source='%v'\nError='%v'\n",
      correctedSrc, err2.Error())

    return err
  }

  err2 = out.Close()

  if err2 != nil {

    err = fmt.Errorf(ePrefix+
      "Error returned from out.Close()\nout=destination='%v'\nError='%v'\n",
      correctedDest, err2.Error())

    return err
  }

  err = nil

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

  filePtr, err := os.Create(pathFileName)

  if err != nil {
    return nil, fmt.Errorf(ePrefix+"Error returned from os.Create(pathFileName): '%v' ",
      err.Error())
  }

  return filePtr, nil
}

// DeleteDirFile - Wrapper function for Remove.
// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func (fh FileHelper) DeleteDirFile(pathFile string) error {
  ePrefix := "FileHelper.DeleteDirFile() "

  errCode := 0

  errCode, _, pathFile = fh.isStringEmptyOrBlank(pathFile)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'pathFile' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'pathFile' consists of blank spaces!")
  }

  if !fh.DoesFileExist(pathFile) {
    // Doesn't exist. Nothing to do.
    return nil
  }

  err := os.Remove(pathFile)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error returned from os.Remove(pathFile). pathFile='%v' Error='%v'", pathFile, err.Error())
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
//      a. Paths
//          D:\T08\x294_1\x394_1\x494_1 and any files in the 'x494_1' directory are deleted
//          D:\T08\x294_1\x394_1\ and any files in the 'x394_1' directory are deleted
//          D:\T08\x294_1\ and any files in the 'x294_1' directory are deleted
//
//      b. The Parent Path 'D:\T08' and any files in that parent path 'D:\T08'
//         directory are unaffected and continue to exist.
//
func (fh FileHelper) DeleteDirPathAll(pathDir string) error {

  ePrefix := "FileHelper.DeleteDirPathAll() "

  errCode := 0

  errCode, _, pathDir = fh.isStringEmptyOrBlank(pathDir)

  if errCode == -1 {
    return errors.New(ePrefix + "Error: Input parameter 'pathDir' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix + "Error: Input parameter 'pathDir' consists of blank spaces!")
  }

  // If the path does NOT exist,
  // Do nothing an return.
  _, err := os.Stat(pathDir)

  if err != nil {
    // Doesn't exist. Nothing to do.
    return nil
  }

  err = os.RemoveAll(pathDir)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by os.RemoveAll(pathDir). pathDir='%v'  Error='%v'",
      pathDir, err.Error())
  }

  _, err = os.Stat(pathDir)

  if err == nil {
    // Path still exists. Something is wrong.
    return fmt.Errorf("Delete Failed! 'pathDir' still exists! \n"+
      "pathDir='%v' ", pathDir)
  }

  return nil
}

// DoesFileExist - Returns a boolean value
// designating whether the passed file name
// exists.
func (fh FileHelper) DoesFileExist(pathFileName string) bool {

  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode < 0 {
    return false
  }

  _, err2 := os.Stat(pathFileName)

  if err2 != nil {
    return false
  }

  return true
}

// DoesFileInfoExist - returns a boolean value indicating
// whether the path and file name passed to the function
// actually exists.
//
// If the file actually exists, the function will return
// the associated FileInfo structure.
//
// However, if 'pathFileName' does NOT exist, an error will
// be returned, return value 'doesFInfoExist' will be set to
// 'false' and return value 'fInfo' will be set to nil.
//
func (fh FileHelper) DoesFileInfoExist(
  pathFileName string) (doesFInfoExist bool, fInfo os.FileInfo, err error) {

  ePrefix := "FileHelper.DoesFileInfoExist() "
  doesFInfoExist = false

  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return doesFInfoExist, nil,
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
  }

  if errCode == -2 {
    return doesFInfoExist, nil,
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
  }

  var err2 error

  if fInfo, err2 = os.Stat(pathFileName); os.IsNotExist(err2) {
    err = fmt.Errorf(ePrefix+"Error from os.Stat(pathFileName). "+
      "'pathFileName' does NOT exist! pathfileName='%v' Error='%v' ", pathFileName, err2)
    return doesFInfoExist, fInfo, err
  }

  if err2 != nil {
    err = fmt.Errorf("Error returned by os.Stat(pathFileName). "+
      "pathFileName='%v' Error='%v' ", pathFileName, err.Error())
    return doesFInfoExist, fInfo, err
  }

  doesFInfoExist = true
  err = nil

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
    err = fmt.Errorf(ePrefix+"Error returned from fh.SearchFilePatternMatch(info, fileSelectionCriteria) info.Name()='%v' Error='%v'", info.Name(), err.Error())
    isMatchedFile = false
    return
  }

  isFileOlderThanSet, isFileOlderThanMatch, err2 := fh.SearchFileOlderThan(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned from dMgr.searchFileOlderThan(info, fileSelectionCriteria) fileSelectionCriteria.FilesOlderThan='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.FilesOlderThan, info.Name(), err.Error())
    isMatchedFile = false
    return
  }

  isFileNewerThanSet, isFileNewerThanMatch, err2 := fh.SearchFileNewerThan(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned from dMgr.searchFileNewerThan(info, fileSelectionCriteria) fileSelectionCriteria.FilesNewerThan='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.FilesNewerThan, info.Name(), err.Error())
    isMatchedFile = false
    return
  }

  isFileModeSearchSet, isFileModeSearchMatch, err2 := fh.SearchFileModeMatch(info, fileSelectionCriteria)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned from dMgr.searchFileModeMatch(info, fileSelectionCriteria) fileSelectionCriteria.SelectByFileMode='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.SelectByFileMode, info.Name(), err.Error())
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

  var err error
  errCode := 0

  errCode, _, pathName = fh.isStringEmptyOrBlank(pathName)

  if errCode == -1 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'pathName' is an empty string!")
  }

  if errCode == -2 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'pathName' consists of blank spaces!")
  }

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode == -1 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'fileSearchPattern' is an empty string!")
  }

  if errCode == -2 {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'fileSearchPattern' consists of blank spaces!")
  }

  pathName, err = fh.MakeAbsolutePath(pathName)

  if err != nil {
    return []string{},
      fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fInfo, err := os.Stat(pathName)

  if err != nil && os.IsNotExist(err) {
    return []string{},
      errors.New(ePrefix + "Error: Input parameter 'pathName' DOES NOT EXIST!")
  }

  if err != nil {
    return []string{},
      fmt.Errorf(ePrefix+
        "Error returned by os.Stat(pathName). "+
        "pathName='%v' Error='%v' ", pathName, err.Error())
  }

  if !fInfo.IsDir() {
    return []string{},
      fmt.Errorf(ePrefix+"Error: The path exists, but it NOT a directory! "+
        "pathName='%v' ", pathName)
  }

  // fInfo is a Directory.

  searchStr := fh.JoinPathsAdjustSeparators(pathName, fileSearchPattern)

  results, err := fp.Glob(searchStr)

  if err != nil {
    return []string{},
      fmt.Errorf(ePrefix+
        "Error returned by fp.Glob(searchStr). "+
        "searchStr='%v' Error='%v' ", searchStr, err.Error())
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
//         SelectByFileMode     os.FileMode // Match file mode. Zero if inactive
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
//                                    Example Patterns:
//                                        "*.log"
//                                        "current*.txt"
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
//      SelectByFileMode os.FileMode -
//                                   'os.FileMode' is an uint32 value. This file selection criterion
//                                   allows for the selection of files by File Mode. File Modes
//                                   are compared to the value	of 'SelectByFileMode'. If the File
//                                   Mode for a given file is equal to the value of 'SelectByFileMode',
//                                   that file is considered to be a 'match' for this file selection
//                                   criterion.
//
//                                   If the value of 'SelectByFileMode' is set equal to zero, then
//                                   this file selection criterion is considered 'Inactive' or
//                                   'Not Set'.
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
//        SelectByFileMode  = uint32(0)
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
      "Error returned from os.Getwd(). Error='%v'", err.Error())
  }

  absDir, err := fh.MakeAbsolutePath(dir)

  if err != nil {
    return "", fmt.Errorf(ePrefix+
      "Error returned by fh.MakeAbsolutePath(dir). Error='%v' ",
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
      errors.New(ePrefix + "Error: Input parameter 'pathStr' is an empty string!")
  }

  if errCode == -2 {
    return []int{},
      errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!")
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
      "Error returned from fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt).\n" +
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

// GetFileInfoFromPath - Wrapper function for os.Stat(). This method
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
func (fh FileHelper) GetFileInfoFromPath(pathFileName string) (os.FileInfo, error) {

  ePrefix := "FileHelper.GetFileInfoFromPath() "
  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return nil,
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
  }

  if errCode == -2 {
    return nil,
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
  }

  fileInfo, err := os.Stat(pathFileName)

  if os.IsNotExist(err) {
    return nil, fmt.Errorf(ePrefix+
      "Error: Input parameter 'pathFileName' does NOT exist! "+
      "pathFileName='%v' Error='%v' ",
      pathFileName, err.Error())
  }

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix+
        "Error returned by os.Stat(pathFileName). pathFileName='%v' "+
        "Error='%v' ", pathFileName, err.Error())
  }

  return fileInfo, nil
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
  const fmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"
  var zeroTime time.Time

  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return time.Time{}, "",
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
  }

  if errCode == -2 {
    return time.Time{}, "",
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
  }

  errCode, _, customTimeFmt = fh.isStringEmptyOrBlank(customTimeFmt)

  fmtStr := customTimeFmt

  if errCode < 0 {
    fmtStr = fmtDateTimeNanoSecondStr
  }

  fInfo, err := fh.GetFileInfoFromPath(pathFileName)

  if err != nil {
    return zeroTime, "",
      errors.New(fmt.Sprintf(ePrefix+
        "Error Getting FileInfo on %v Error on GetFileInfoFromPath(): %v",
        pathFileName, err.Error()))
  }

  return fInfo.ModTime(), fInfo.ModTime().Format(fmtStr), nil
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

    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' is an empty string!")

    return firstIdx, lastIdx, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix + "Error: Input parameter 'pathStr' consists of blank spaces!")

    return firstIdx, lastIdx, err
  }

  pathStr = fh.AdjustPathSlash(pathStr)

  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {

    err = fmt.Errorf(ePrefix + "Error: After path Separator adjustment, 'pathStr' is an empty string!")

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
      errors.New(ePrefix + "Error: Input parameter 'pathName' is an empty string!")
  }

  if errCode == -2 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'pathName' consists of blank spaces!")
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
    err = errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
    return pathDir, fileNameExt, bothAreEmpty, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
    return pathDir, fileNameExt, bothAreEmpty, err
  }

  xFnameExt, isEmpty, err2 := fh.GetFileNameWithExt(trimmedFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetFileNameWithExt(pathFileNameExt). pathFileNameExt='%v' Error='%v'", pathFileNameExt, err2.Error())
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
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetPathFromPathFileName(remainingPathStr). remainingPathStr='%v' Error='%v'", remainingPathStr, err2.Error())
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
//														If no error occurs, 'err' is set to 'nil'.
//
// ------------------------------------------------------------------------
//
// Examples:
//
//  pathFileNameExt = ""        returns isEmpty==true  err==nil
//  pathFileNameExt = "D:\"     returns "D:\"
//  pathFileNameExt = "."       returns "."
//  pathFileNameExt = "..\"     returns "..\"
//  pathFileNameExt = "...\"    returns ERROR
//  pathFileNameExt = ".\pathfile\003_filehelper\wt_HowToRunTests.md"  returns ".\pathfile\003_filehelper"
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
    err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")
    return dirPath, isEmpty, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")
    return dirPath, isEmpty, err
  }

  testPathStr, isDirEmpty, err2 := fh.CleanDirStr(pathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned by fh.CleanDirStr(pathFileNameExt). pathFileNameExt='%v'  Error='%v'", pathFileNameExt, err2.Error())
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
      "Error: AdjustPathSlash was applied to 'pathStr'. The 'testPathStr' string is a Zero Length string!")
    return dirPath, isEmpty, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathStr). testPathStr='%v'  Error='%v'",
      testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  lSlashIdxs := len(slashIdxs)

  firstGoodChar, lastGoodChar, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr). "+
      "testPathStr='%v'  Error='%v'",
      testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathStr). "+
      "testPathStr='%v'  Error='%v'", testPathStr, err2.Error())
    return dirPath, isEmpty, err
  }

  lDotIdxs := len(dotIdxs)

  var finalPathStr string

  volName := fp.VolumeName(testPathStr)

  if testPathStr == volName {

    finalPathStr = testPathStr

  } else if strings.Contains(testPathStr, "...") {

    err = fmt.Errorf(ePrefix+"Error: PATH CONTAINS INVALID Dot Characters! testPathStr='%v'", testPathStr)
    return dirPath, isEmpty, err

  } else if firstGoodChar == -1 || lastGoodChar == -1 {

    absPath, err2 := fh.MakeAbsolutePath(testPathStr)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+"Error returned from fh.MakeAbsolutePath(testPathStr). testPathStr='%v' Error='%v'", testPathStr, err2.Error())
      return dirPath, isEmpty, err
    }

    if absPath == "" {
      err = fmt.Errorf(ePrefix+"Error: Could not convert 'testPathStr' to Absolute path! tesPathStr='%v'", testPathStr)
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
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH STRING. testPathStr='%v'", testPathStr)
    return dirPath, isEmpty, err
  }

  if len(finalPathStr) == 0 {
    err = fmt.Errorf(ePrefix + "Error: Processed path is a Zero Length String!")
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
        "Error: Input parameter 'pathStr' is an empty string!")
  }

  if errCode == -2 {
    return []int{},
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' consists of blank spaces!")
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
// VolumeName returns leading volume name.
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share".
// On other platforms it returns "".
//
func (fh FileHelper) GetVolumeName(pathStr string) string {

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return ""
  }

  return fp.VolumeName(pathStr)
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

  errCode := 0

  errCode, _, pathFileStr = fh.isStringEmptyOrBlank(pathFileStr)

  if errCode == -1 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathFileStr' is an empty string!")
    return pathFileType, absolutePathFile, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathFileStr' consists of blank spaces!")

    return pathFileType, absolutePathFile, err
  }

  if strings.Contains(pathFileStr, "...") {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH STRING! pathFileStr='%v'", pathFileStr)
    return pathFileType, absolutePathFile, err
  }

  correctedPathFileStr := fh.AdjustPathSlash(pathFileStr)

  firstCharIdx, lastCharIdx, err2 :=
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr"+
      "(correctedPathFileStr) correctedPathFileStr='%v'  Error='%v'",
      correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetPathSeparatorIndexesInPathStr(correctedPathFileStr) returned error. "+
      "correctedPathFileStr='%v' Error='%v'",
      correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(correctedPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetDotSeparatorIndexesInPathStr(correctedPathFileStr) retured error. "+
      "correctedPathFileStr='%v' Error='%v'", correctedPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  lenDotIdx := len(dotIdxs)

  lenSlashIdx := len(slashIdxs)

  testAbsPathFileStr, err2 := fh.MakeAbsolutePath(correctedPathFileStr)

  if err2 != nil {
    err = fmt.Errorf("Error converting pathFileStr to absolute path. "+
      "pathFileStr='%v' Error='%v' ", pathFileStr, err2.Error())

    return pathFileType, absolutePathFile, err
  }

  if lenDotIdx > 0 &&
    lenSlashIdx == 0 &&
    firstCharIdx > -1 {
    absolutePathFile = testAbsPathFileStr
    pathFileType = PathFileType.File()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  if lenDotIdx == 0 &&
    lenSlashIdx == 0 &&
    firstCharIdx > -1 {

    absolutePathFile = testAbsPathFileStr
    pathFileType = PathFileType.File()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  if firstCharIdx == -1 &&
    lastCharIdx == -1 &&
    lenDotIdx > 0 {

    absolutePathFile = testAbsPathFileStr
    pathFileType = PathFileType.Path()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  lTestAbsPathStr := len(testAbsPathFileStr)

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
  fInfo, err2 := os.Stat(testAbsPathFileStr)

  if err2 == nil {

    if fInfo.IsDir() {

      pathFileType = PathFileType.Path()

      absolutePathFile = testAbsPathFileStr

      err = nil

      return pathFileType, absolutePathFile, err

    } else {

      pathFileType = PathFileType.PathFile()

      absolutePathFile = testAbsPathFileStr

      err = nil

      return pathFileType, absolutePathFile, err
    }

  }

  // Ok - We know the testPathFileStr does NOT exist on disk

  firstCharIdx, lastCharIdx, err2 =
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(testAbsPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr"+
      "(testAbsPathFileStr) testAbsPathFileStr='%v'  Error='%v'",
      testAbsPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  if firstCharIdx == -1 || lastCharIdx == -1 {
    // The path-file-string contains no alpha numeric characters.
    // Therefore, it does NOT contain a file name!
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "testAbsPathFileStr does NOT contain alpha numeric characters "+
      "testAbsPathFileStr='%v'", testAbsPathFileStr)
    return pathFileType, absolutePathFile, err
  }

  if testAbsPathFileStr[lTestAbsPathStr-1] == os.PathSeparator {
    // The last character is a path separator.
    // Example D:\directory1\directory2\
    // So, this must be a path and NOT a path file Name
    pathFileType = PathFileType.Path()

    absolutePathFile = testAbsPathFileStr

    err = nil

    return pathFileType, absolutePathFile, err
  }

  slashIdxs, err2 = fh.GetPathSeparatorIndexesInPathStr(testAbsPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetPathSeparatorIndexesInPathStr(testAbsPathFileStr) returned error. "+
      "testAbsPathFileStr='%v' Error='%v'",
      testAbsPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  dotIdxs, err2 = fh.GetDotSeparatorIndexesInPathStr(testAbsPathFileStr)

  if err2 != nil {
    pathFileType = PathFileType.None()
    absolutePathFile = ""
    err = fmt.Errorf(ePrefix+
      "fh.GetDotSeparatorIndexesInPathStr(testAbsPathFileStr) retured error. "+
      "testAbsPathFileStr='%v' Error='%v'", testAbsPathFileStr, err2.Error())
    return pathFileType, absolutePathFile, err
  }

  lenDotIdx = len(dotIdxs)

  lenSlashIdx = len(slashIdxs)

  if lenSlashIdx == 0 &&
    lenDotIdx > 0 {
    // This is a string of alpha numeric characters which
    // does NOT contain a path separator, but Does contain
    // a dot separator.
    // Example "someFileName.txt"
    // Let's call it a file
    absolutePathFile = testAbsPathFileStr

    // Call it a file!
    // Example: "somefilename"
    pathFileType = PathFileType.File()
    err = nil
    return pathFileType, absolutePathFile, err

  }

  if lenSlashIdx == 0 &&
    lenDotIdx == 0 {

    // the string has no dots and no path separators.
    absolutePathFile = testAbsPathFileStr

    // Call it a file!
    // Example: "somefilename"
    pathFileType = PathFileType.File()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  if lenSlashIdx > 0 &&
    lenDotIdx == 0 {

    // string contains one or more path separators, but no dots (aka periods)
    absolutePathFile = testAbsPathFileStr

    if lastCharIdx > slashIdxs[lenSlashIdx-1] {
      // Example D:\dir1\dir2\xray
      // It could be a file or a path. We simply
      // can't tell.
      pathFileType = PathFileType.Indeterminate()

      err = nil
      return pathFileType, absolutePathFile, err
    }

    // lastCharIdx <= slashIdxs[lenSlashIdx-1]
    // Call it a Path (aka Directory)
    pathFileType = PathFileType.Path()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  // We know that the test string contains both path separators and
  // dot separators ('.')

  if dotIdxs[lenDotIdx-1] > slashIdxs[lenSlashIdx-1] &&
    lastCharIdx > slashIdxs[lenSlashIdx-1] {

    // Example D:\dir1\dir2\xray.txt
    absolutePathFile = testAbsPathFileStr

    // Call this a path file name
    pathFileType = PathFileType.PathFile()
    err = nil
    return pathFileType, absolutePathFile, err
  }

  // Cannot be certain of the result.
  // String doesn't conform to any standard pattern for
  // a file or a path. Don't know for sure what this string is

  absolutePathFile = testAbsPathFileStr

  pathFileType = PathFileType.None()
  err = errors.New(ePrefix + "Unknown string type!")
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

  isPathStr = false
  cannotDetermine = false
  testPathStr = fh.AdjustPathSlash(pathStr)
  err = nil

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' is an empty string!")
    return isPathStr, cannotDetermine, testPathStr, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix +
        "Error: Input parameter 'pathStr' consists of blank spaces!")

    return isPathStr, cannotDetermine, testPathStr, err
  }

  if strings.Contains(pathStr, "...") {
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH STRING! pathStr='%v'", pathStr)
    return isPathStr, cannotDetermine, testPathStr, err
  }

  pathFileType, _, err2 := fh.IsPathFileString(testPathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v", err2.Error())
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
func (fh FileHelper) JoinPathsAdjustSeparators(p1 string, p2 string) string {

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
        "Error: Input parameter 'relPath' is an empty string!")
  }

  if errCode == -2 {
    return "",
      errors.New(ePrefix +
        "Error: Input parameter 'relPath' consists of blank spaces!")
  }

  testRelPath := fh.AdjustPathSlash(relPath)

  errCode, _, testRelPath = fh.isStringEmptyOrBlank(testRelPath)

  if errCode < 0 {
    return "", errors.New(ePrefix +
      "Error: Input Parameter 'relPath' adjusted for path Separators is an EMPTY string!")
  }

  p, err := fp.Abs(testRelPath)

  if err != nil {
    return "Invalid p!", fmt.Errorf(ePrefix+"Error returned from  fp.Abs(testRelPath). testRelPath='%v'  Error='%v'", testRelPath, err.Error())
  }

  return p, err
}

// MakeDirAll - creates a directory named path, along with any necessary
// parent directories.
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
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fh.MakeDirAllPerm(dirPath, permission)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
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
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  err = fh.MakeDirPerm(dirPath, permission)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
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
    return  errors.New(ePrefix +
        "Error: Input parameter 'dirPath' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
        "Error: Input parameter 'dirPath' consists of blank spaces!")
  }

  err2 := permission.IsValid()

  if err2 != nil {
    return fmt.Errorf(ePrefix+"Input parameter 'permission' is INVALID! "+
      "Error='%v' ", err2.Error())
  }

  dirPermCode, err2 := permission.GetCompositePermissionMode()

  if err2 != nil {
    return fmt.Errorf(ePrefix+"INVALID Permission Code "+
      "Error='%v' ", err2.Error())
  }

  err2 = os.MkdirAll(dirPath, dirPermCode)

  if err2 != nil {
    return fmt.Errorf(ePrefix+"Error return from os.MkdirAll(dirPath, permission). "+
      "dirPath='%v' Error='%v' ", dirPath, err2.Error())
  }

  _, err2 = os.Stat(dirPath)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST! \n"+
      "dirPath='%v' \n", dirPath)
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
    return  errors.New(ePrefix +
      "Error: Input parameter 'dirPath' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dirPath' consists of blank spaces!")
  }

  err2 := permission.IsValid()

  if err2 != nil {
    return fmt.Errorf(ePrefix+"Input parameter 'permission' is INVALID! "+
      "Error='%v' ", err2.Error())
  }

  dirPermCode, err2 := permission.GetCompositePermissionMode()

  if err2 != nil {
    return fmt.Errorf(ePrefix+"INVALID Permission Code "+
      "Error='%v' ", err2.Error())
  }

  err2 = os.Mkdir(dirPath, dirPermCode)

  if err2 != nil {
    return fmt.Errorf(ePrefix+"Error return from os.Mkdir(dirPath, dirPermCode). "+
      "dirPath='%v' Error='%v' ", dirPath, err2.Error())
  }

  _, err2 = os.Stat(dirPath)

  if err2 != nil {
    return fmt.Errorf(ePrefix+
      "Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST! \n"+
      "dirPath='%v' \n", dirPath)
  }

  return nil
}

// MoveFile - Copies file from source to destination and, if successful,
// then deletes the original source file.
//
// The copy procedure will carried out using the the 'Copy By Io' technique.
// See FileHelper.CopyFileByIo().
//
// If this copy operation fails, the method will return an error and it
// will NOT delete the source file.
//
// If an error is encountered during this procedure it will be returned by
// means of the return parameter 'err'.
//
func (fh FileHelper) MoveFile(src, dst string) error {

  ePrefix := "FileHelper.MoveFile() "

  errCode := 0

  errCode, _, src = fh.isStringEmptyOrBlank(src)

  if errCode == -1 {
    return errors.New(ePrefix +
      "Error: Input parameter 'src' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
      "Error: Input parameter 'src' consists of blank spaces!")
  }

  errCode, _, dst = fh.isStringEmptyOrBlank(dst)

  if errCode == -1 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dst' is an empty string!")
  }

  if errCode == -2 {
    return errors.New(ePrefix +
      "Error: Input parameter 'dst' consists of blank spaces!")
  }

  _, err := os.Stat(src)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
  }


  // ============================
  // Perform the copy operation!
  // ============================
  err = fh.CopyFileByIo(src, dst)

  if err != nil {
    // Copy Operation Failed. Return an error
    // and DO NOT delete the source file!
    return fmt.Errorf(ePrefix +
      "Error: Copy operation FAILED!\nSource File='%v'\nDestination File='%v'\nError='%v'\n",
      src, dst, err.Error())
  }

  // CopyFileByIo operation was apparently successful.
  // Now, verify that destination file exists.

  _, err = os.Stat(dst)

  if err != nil {
    return fmt.Errorf(ePrefix + "Error: After Copy Operation, destination file DOES NOT EXIST!\n" +
      "Therefore, the copy operation FAILED! Source file was NOT deleted.\n" +
      "destination file='%v'\n", dst)
  }

  // Successful copy operation has been verified.
  // Time to delete the source file.
  err = os.Remove(src)

  if err != nil {
    return fmt.Errorf(ePrefix +
      "Copy operation succeeded, but attempted deletion of source file FAILED!\n" +
      "Source File='%v'\n", src)
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
  createDir bool,) (*os.File, error) {


  ePrefix := "FileHelper.OpenDirectory() "
  var err error
  errCode := 0

  errCode, _, directoryPath = fh.isStringEmptyOrBlank(directoryPath)

  if errCode == -1 {
    return nil,
      errors.New(ePrefix + "Input parameter 'directoryPath' is an empty string!")

  }

  if errCode == -2 {
    return nil,
      errors.New(ePrefix +
        "Input parameter 'directoryPath' consists of all spaces!")
  }

  directoryPath, err = fh.MakeAbsolutePath(directoryPath)

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix + "Error returned by fh.MakeAbsolutePath(directoryPath).\n" +
        "directoryPath='%v'\nError='%v'", directoryPath, err.Error())
  }

  fInfo, err := os.Stat(directoryPath)

  if err != nil {

    if !createDir {
      return nil,
        fmt.Errorf(ePrefix + "Error returned by os.Stat(directoryPath).\n" +
          "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
    }

    // Parameter 'createDir' must be 'true'.
    // The error signaled that the path does not exist. So, create the directory path
    err = fh.MakeDirAll(directoryPath)

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix + "ERROR: Attmpted creation of 'directoryPath' FAILED!\n" +
          "directoryPath='%v'\nError='%v'\n", directoryPath, err.Error())
    }

    // Verify that the directory exists and get
    // the associated file info object.
    fInfo, err = os.Stat(directoryPath)

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix + "Error occurred verifying existance of newly created 'directoryPath'!\n"+
          "Error returned by os.Stat(directoryPath)\ndirectoryPath='%v'\nError='%v'\n",
          directoryPath, err.Error())
    }

  }

  if !fInfo.IsDir() {
    return nil,
      fmt.Errorf(ePrefix + "ERROR: Input Paramter 'directoryPath' is NOT a directory!\n" +
        "directoryPath='%v'", directoryPath)
  }


  filePtr, err := os.Open(directoryPath)

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix + "File Open Error: %v\n" +
        "directoryPath='%v'", err.Error(), directoryPath)
  }


  if filePtr == nil {
    return nil, errors.New(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!")
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
  errCode := 0
  ePrefix := "FileHelper.OpenFile() "

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    err = errors.New(ePrefix + "Input parameter 'pathFileName' is an empty string!")
    return filePtr, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Input parameter 'pathFileName' consists of all spaces!")
    return filePtr, err
  }

  err2 := fileOpenCfg.IsValid()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Input Parameter 'fileOpenCfg' is INVALID!\n" +
      "Error='%v'", err2.Error())
    return filePtr, err
  }

  fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "%v", err2.Error())
    return filePtr, err
  }

  err2 = filePermissionCfg.IsValid()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Input Parameter 'filePermissionCfg' is INVALID!\n" +
      "Error='%v'\n", err2.Error())
    return filePtr, err
  }

  fileMode, err2 := filePermissionCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "%v", err2.Error())
    return filePtr, err
  }

  filePtr, err2 = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err2 != nil {

    if os.IsNotExist(err2) {
      err = fmt.Errorf(ePrefix+"The 'pathFileName' DOES NOT EXIST! "+
        "pathFileName='%v' Error='%v' ",
        pathFileName, err2.Error())
      filePtr = nil
      return filePtr, err
    }

    err = fmt.Errorf(ePrefix+
      "Error returned by os.OpenFile(pathFileName, fOpenCode, fileMode) "+
      "targetpathFileName='%v' Error='%v' ", pathFileName, err.Error())

    return filePtr, err
  }

  if filePtr == nil {
    err = errors.New(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!")
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
  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    err = errors.New(ePrefix + "Input parameter 'pathFileName' is an empty string!\n")
    return filePtr, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Input parameter 'pathFileName' consists entirely of blank spaces!\n")
    return filePtr, err
  }

  pathFileName, err2 = fh.MakeAbsolutePath(pathFileName)

  if err2 != nil {
    err = fmt.Errorf(ePrefix +
      "Error occurred while converting input parameter file name ('pathFileName') to absolute path.\n" +
      "pathFileName='%v'\nError='%v'\n", pathFileName, err2.Error())

    return filePtr, err
  }


  if !fh.DoesFileExist(pathFileName) {
    err = fmt.Errorf(ePrefix +
      "ERROR: The input parameter 'pathFileName' DOES NOT EXIST!\n" +
      "pathFileName='%v' ", pathFileName)
    return filePtr, err
  }

  fileOpenCfg, err2 := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix +
        "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadOnly()," +
        "FOpenMode.ModeNone()).\nError='%v'\n",
        err2.Error())
    return filePtr, err
  }

  fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error Creating File Open Code.\nError=%v\n",
      err2.Error())
    return filePtr, err
  }

  fPermCfg, err2:= FilePermissionConfig{}.New("-r--r--r--")

  if err2 != nil {
    err =
      fmt.Errorf(ePrefix +
        "Error returned by FilePermissionConfig{}.New(\"-r--r--r--\")\n" +
        "Error='%v' \n", err2.Error())
    return filePtr, err
  }

  fileMode, err2 := fPermCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "Error Creating File Mode Code.\nError=%v\n",
      err2.Error())

    return filePtr, err
  }

  filePtr, err2 = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "File Open Error: %v\n" +
      "pathFileName='%v'", err.Error(), pathFileName)
    filePtr = nil
    return filePtr, err
  }


  if filePtr == nil {
    err = fmt.Errorf(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!")

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

  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return nil,
    errors.New(ePrefix + "Input parameter 'pathFileName' is an empty string!")

  }

  if errCode == -2 {
    return nil, errors.New(ePrefix +
      "Input parameter 'pathFileName' consists of all spaces!")
  }

  pathFileName, err = fh.MakeAbsolutePath(pathFileName)

  if err != nil {
    return nil,
    fmt.Errorf(ePrefix +
      "Error creating absolute path: %v\n" +
      "pathFileName='%v'\n", err.Error(), pathFileName)
  }

  var fileOpenCfg FileOpenConfig

  if !fh.DoesFileExist(pathFileName) {

    fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
      FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix +
          "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly()," +
          "FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\nError='%v'\n",
          err.Error())
    }

  } else {

    if truncateFile {
      // truncateFile == true
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
        FOpenMode.ModeTruncate())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix +
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite()," +
            "FOpenMode.ModeTruncate()).\nError='%v'\n",
            err.Error())
      }

    } else {
      // truncateFile == false
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
        FOpenMode.ModeAppend())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix +
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite()," +
            "FOpenMode.ModeAppend()).\nError='%v'\n",
            err.Error())
      }
    }
  }

  fOpenCode, err := fileOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix + "%v", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("-rwxrwxrwx")

  if err != nil {
    return nil,
    fmt.Errorf(ePrefix +
      "Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n" +
      "Error='%v' \n", err.Error())
  }

  fileMode, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return nil, fmt.Errorf(ePrefix + "%v", err.Error())
  }

  fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err != nil {
    return nil, fmt.Errorf(ePrefix + "File Open Error: %v\n" +
      "pathFileName='%v'", err.Error(), pathFileName)
  }

  if fPtr == nil {
    return nil, fmt.Errorf(ePrefix +
      "ERROR: os.OpenFile() returned a 'nil' file pointer!")
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

  errCode := 0

  errCode, _, pathFileName = fh.isStringEmptyOrBlank(pathFileName)

  if errCode == -1 {
    return nil,
      errors.New(ePrefix + "Input parameter 'pathFileName' is an empty string!")

  }

  if errCode == -2 {
    return nil, errors.New(ePrefix +
      "Input parameter 'pathFileName' consists of all spaces!")
  }

  pathFileName, err = fh.MakeAbsolutePath(pathFileName)

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix +
        "Error creating absolute path: %v\n" +
        "pathFileName='%v'\n", err.Error(), pathFileName)
  }

  var fileOpenCfg FileOpenConfig

  if !fh.DoesFileExist(pathFileName) {
    // The pathFileName DOES NOT EXIST!

    fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
      FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

    if err != nil {
      return nil,
        fmt.Errorf(ePrefix +
          "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly()," +
          "FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\nError='%v'\n",
          err.Error())
    }

  } else {
    // The pathFileName DOES EXIST!

    if truncateFile {
      // truncateFile == true; Set Mode 'Truncate'
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
        FOpenMode.ModeTruncate())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix +
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly()," +
            "FOpenMode.ModeTruncate()).\nError='%v'\n",
            err.Error())
      }

    } else {
      // truncateFile == false; Set Mode 'Append'
      fileOpenCfg, err = FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
        FOpenMode.ModeAppend())

      if err != nil {
        return nil,
          fmt.Errorf(ePrefix +
            "Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly()," +
            "FOpenMode.ModeAppend()).\nError='%v'\n",
            err.Error())
      }
    }
  }

  fOpenCode, err := fileOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix +
        "Error creating File Open Code.\nError=%v\n", err.Error())
  }

  fPermCfg, err := FilePermissionConfig{}.New("--wx-wx-wx")

  if err != nil {
    return nil,
      fmt.Errorf(ePrefix +
        "Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n" +
        "Error='%v' \n", err.Error())
  }

  fileMode, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return nil, fmt.Errorf(ePrefix +
      "Error creating file mode code.\nError=%v\n", err.Error())
  }

  fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

  if err != nil {
    return nil, fmt.Errorf(ePrefix +
      "Error returned from os.OpenFile().\nError='%v'\n" +
      "pathFileName='%v'", err.Error(), pathFileName)
  }

  if fPtr == nil {
    return nil, fmt.Errorf(ePrefix +
      "ERROR: File pointer returned from os.OpenFile() is 'nil'!")
  }

  return fPtr, nil
}

// RemovePathSeparatorFromEndOfPathString - Remove Trailing path Separator from
// a path string - if said trailing path Separator exists.
func (fh FileHelper) RemovePathSeparatorFromEndOfPathString(pathStr string) string {
  lPathStr := len(pathStr)

  if lPathStr == 0 {
    return ""
  }

  lastChar := rune(pathStr[lPathStr-1])

  if lastChar == os.PathSeparator ||
    lastChar == '\\' ||
    lastChar == '/' {

    if lPathStr < 2 {
      return ""
    }

    return pathStr[0 : lPathStr-1]
  }

  return pathStr
}

// SearchFileModeMatch - This method determines whether the file mode of the file described by input
// parameter, 'info', is match for the File Selection Criteria 'fileSelectCriteria.SelectByFileMode'.
// If the file's FileMode matches the 'fileSelectCriteria.SelectByFileMode' value, the return value,
// 'isFileModeMatch' is set to 'true'.
//
// If 'fileSelectCriteria.SelectByFileMode' is set to zero, the return value 'isFileModeSet' set to 'false'
// signaling the File Mode File Selection Criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func (fh *FileHelper) SearchFileModeMatch(info os.FileInfo, fileSelectCriteria FileSelectionCriteria) (isFileModeSet, isFileModeMatch bool, err error) {

  if fileSelectCriteria.SelectByFileMode == 0 {
    isFileModeSet = false
    isFileModeMatch = false
    err = nil
    return
  }

  if fileSelectCriteria.SelectByFileMode == info.Mode() {
    isFileModeSet = true
    isFileModeMatch = true
    err = nil
    return

  }

  isFileModeSet = true
  isFileModeMatch = false
  err = nil
  return
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
func (fh *FileHelper) SearchFileNewerThan(info os.FileInfo, fileSelectCriteria FileSelectionCriteria) (isFileNewerThanSet, isFileNewerThanMatch bool, err error) {

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
func (fh *FileHelper) SearchFileOlderThan(info os.FileInfo, fileSelectCriteria FileSelectionCriteria) (isFileOlderThanSet, isFileOlderThanMatch bool, err error) {

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
func (fh *FileHelper) SearchFilePatternMatch(info os.FileInfo, fileSelectCriteria FileSelectionCriteria) (isPatternSet, isPatternMatch bool, err error) {

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
      err = fmt.Errorf(ePrefix+"Error returned from filepath.Match(fileSelectCriteria.FileNamePatterns[i] , info.Name()) fileSelectCriteria.FileNamePatterns[i]='%v' info.Name()='%v' Error='%v'", fileSelectCriteria.FileNamePatterns[i], info.Name(), err.Error())
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

// SetCurrentWorkingDir - Similar to FileHelper.ChangeWorkingDir().
// However, this method receives a file pointer of type *os.File.
// The input parameter, 'fPtr' must point to a directory. If an
// error is returned, it will be of type *PathError.
func (fh FileHelper) SetCurrentWorkingDir(fPtr *os.File) error {

  err := fPtr.Chdir()

  if err != nil {
    ePrefix := "FileHelper.SetCurrentWorkingDir() "

    return fmt.Errorf(ePrefix+
      "Error returned by fPtr.Chdir(). Error='%v' ", err.Error())
  }

  return nil
}

// SwapBasePath - Searches the 'targetPath' string for
// the existence of 'oldBasePath'. If 'oldBasePath' is
// found, it is replaced with 'newBasePath'.
//
// If 'oldBasePath' is not found in 'targetPath' an
// error is returned.
//
// Likewise, if 'oldBasePath' is not located at the beginning
// of 'targetPath', an error will be returned.
//
func (fh FileHelper) SwapBasePath(
  oldBasePath,
  newBasePath,
  targetPath string) (string, error) {

  ePrefix := "FileHelper.SwapBasePath() "

  oldBaseLen := len(oldBasePath)

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

// WriteBytes - Wrapper for os.File.Write(). Writes an array of bytes
// to an open file pointed to by 'fPtr' (*os.File)
func (fh FileHelper) WriteBytes(b []byte, fPtr *os.File) (int, error) {

  return fPtr.Write(b)

}

// WriteFileStr - Wrapper for *os.File.WriteString. Writes a string
// to an open file pointed to by 'fPtr' (*os.File).
func (fh FileHelper) WriteFileStr(str string, fPtr *os.File) (int, error) {

  return fPtr.WriteString(str)

}

/*
  FileHelper private methods
*/

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
//
func (fh FileHelper) isStringEmptyOrBlank(testStr string) (errCode int, strLen int, newStr string) {

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

    ePrefix := "DirMgr.makeFileHelperWalkDirDeleteFilesFunc"

    if erIn != nil {
      dInfo.ErrReturns = append(dInfo.ErrReturns, erIn.Error())
      return nil
    }

    if info.IsDir() {

      subDir, err := DirMgr{}.New(pathFile)

      if err != nil {
        ex := fmt.Errorf(ePrefix+"Error returned from DirMgr{}.NewFromPathFileNameExtStr(pathFile). pathFile:='%v' Error='%v'", pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())

        if subDir.isInitialized {
          subDir.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)
          dInfo.Directories.AddDirMgr(subDir)
        }

        return nil
      }

      subDir.actualDirFileInfo = FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)
      dInfo.Directories.AddDirMgr(subDir)

      return nil
    }

    fh := FileHelper{}

    isFoundFile, err := fh.FilterFileName(info, dInfo.DeleteFileSelectCriteria)

    if err != nil {

      ex := fmt.Errorf(ePrefix+"Error returned from dMgr.FilterFileName(info, dInfo.DeleteFileSelectCriteria) pathFile='%v' info.Name()='%v' Error='%v' ", pathFile, info.Name(), err.Error())
      dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())
      return nil
    }

    if isFoundFile {

      err := os.Remove(pathFile)

      if err != nil {
        ex := fmt.Errorf(ePrefix+
          "Error returned from os.Remove(pathFile). pathFile='%v' Error='%v'",
          pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())
        return nil
      }

      err = dInfo.DeletedFiles.AddFileMgrByFileInfo(pathFile, info)

      if err != nil {
        ex := fmt.Errorf(ePrefix+
          "Error returned from dInfo.DeletedFiles.AddFileMgrByFileInfo( pathFile,  info). "+
          "pathFile='%v'  Error='%v'",
          pathFile, err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())
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
      dInfo.ErrReturns = append(dInfo.ErrReturns, ex2.Error())
      return nil
    }

    if info.IsDir() {
      subDir, err := DirMgr{}.NewFromFileInfo(pathFile, info)
      if err != nil {

        if subDir.isInitialized {
          dInfo.Directories.AddDirMgr(subDir)
        }

        er2 := fmt.Errorf(ePrefix+"Error returned by DirMgr{}.NewFromPathFileNameExtStr(pathFile). "+
          "pathFile='%v' Error='%v'", pathFile, err.Error())
        dInfo.ErrReturns = append(dInfo.ErrReturns, er2.Error())
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

      er2 := fmt.Errorf(ePrefix+"Error returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria) pathFile='%v' info.Name()='%v' Error='%v' ", pathFile, info.Name(), err.Error())
      dInfo.ErrReturns = append(dInfo.ErrReturns, er2.Error())
      return nil
    }

    if isFoundFile {

      fMgr, err2 := FileMgr{}.NewFromPathFileNameExtStr(pathFile)

      if err2 != nil {
        err = fmt.Errorf(ePrefix+
          "Error returned by FileMgr{}.NewFromPathFileNameExtStr(pathFile) "+
          "pathFile='%v' Error='%v' ", pathFile, err2.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, err.Error())

        return nil
      }

      err = dInfo.FoundFiles.AddFileMgrByFileInfo(fMgr.dMgr.GetAbsolutePath(), info)

      if err != nil {
        er2 := fmt.Errorf(ePrefix+"Error returned from  dInfo.FoundFiles.AddFileMgrByFileInfo( pathFile,  info) "+
          "pathFile='%v' info.Name()='%v' Error='%v' ",
          pathFile, info.Name(), err.Error())

        dInfo.ErrReturns = append(dInfo.ErrReturns, er2.Error())

        return nil
      }
    }

    return nil
  }
}
