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

  return fp.FromSlash(path)
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
// Example:
// dirName = '../dir1/dir2/fileName.ext' returns "../dir1/dir2"
// dirName = 'fileName.ext' returns "" isEmpty = true
// dirName = '../dir1/dir2/' returns '../dir1/dir2'
// dirName = '../dir1/dir2/filename.ext' returns '../dir1/dir2'
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
// initialDecimalValue := 511
// expectedOctalValue := 777
//
// actualOctalValue := ConvertDecimalToOctal(initialDecimalValue)
//
// 'actualOctalValue' is now equal to integer value '777'.
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
// "os.Link(src, dst)" is the only method employed to copy a
// designated file. If "os.Link(src, dst)" fails, an err is returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
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

  if !fh.DoesFileExist(src) {
    err = fmt.Errorf(ePrefix+"Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
    return
  }

  sfi, err2 := os.Stat(src)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned from os.Stat(src). src='%v'  Error='%v'", src, err2.Error())
    return
  }

  if !sfi.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+"Error: non-regular source file. Source File Name='%v'  Source File Mode='%v' ", sfi.Name(), sfi.Mode().String())
    return
  }

  dfi, err2 := os.Stat(dst)

  if err2 != nil {

    if !os.IsNotExist(err2) {
      // Must be PathError - path does not exist
      err = fmt.Errorf(ePrefix+"Destination File path Error - path does NOT exist. "+
        "Destination File='%v' Error: %v", dst, err2.Error())
      return
    }

  } else {

    if !(dfi.Mode().IsRegular()) {
      err = fmt.Errorf(ePrefix+"non-regular destination file - Cannot Overwrite "+
        "destination file. Destination File='%v'  Destination File Mode= '%v'",
        dfi.Name(), dfi.Mode().String())
      return
    }

    // Source and destination have the same path
    // and file names. They are one in the same
    // file. Nothing to do.
    if os.SameFile(sfi, dfi) {
      err = nil
      return
    }

  }

  err2 = os.Link(src, dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"- os.Link(src, dst) FAILED! src='%v' dst='%v'  Error='%v'", src, dst, err2.Error())
    return
  }

  err = nil

  return
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

  sfi, err2 := os.Stat(src)

  if err2 != nil {
    if os.IsNotExist(err2) {
      // Must be PathError - source path & file name do not exist!
      err = fmt.Errorf(ePrefix+
        "Source File path Error - path does NOT exist. Source File='%v' Error: %v",
        src, err2.Error())
      return err
    }

    err = fmt.Errorf(ePrefix+
      "Error returned from os.Stat(src). src='%v'  Error='%v'", src, err2.Error())
    return err
  }

  if !sfi.Mode().IsRegular() {
    // cannot copy non-regular files (e.g., directories,
    // symlinks, devices, etc.)
    err = fmt.Errorf(ePrefix+"Error non-regular source file ='%v' source file Mode='%v'",
      sfi.Name(), sfi.Mode().String())
    return err
  }

  dfi, err2 := os.Stat(dst)

  if err2 != nil {

    if !os.IsNotExist(err2) {
      // Must be PathError - path does not exist
      err = fmt.Errorf(ePrefix+"Destination File path Error - path does NOT exist. Destination File='%v' Error: %v", dst, err.Error())
      return err
    }

    // The destination file does not exist

  } else {
    // The destination file already exists!
    if !dfi.Mode().IsRegular() {
      err = fmt.Errorf(ePrefix+"Error: non-regular destination file. Cannot Overwrite destination file. Destination file='%v' destination file mode='%v'", dfi.Name(), dfi.Mode().String())
      return err
    }

    if os.SameFile(sfi, dfi) {
      // Source and destination are the same
      // path and file name.
      err = nil
      return err
    }

  }

  // Create a new destination file and copy source
  // file contents to the destination file.

  in, err2 := os.Open(src)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Open(src) src='%v'  Error='%v'",
      src, err2.Error())
    return err
  }

  out, err2 := os.Create(dst)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from os.Create(destinationFile) "+
      "destinationFile='%v'  Error='%v'",
      dst, err2.Error())

    _ = in.Close()

    return err
  }

  if _, err2 = io.Copy(out, in); err2 != nil {
    _ = in.Close()
    _ = out.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from io.Copy(destination, source) destination='%v' "+
      "source='%v'  Error='%v' ",
      dst, src, err2.Error())
    return
  }

  // flush file buffers in memory
  err2 = out.Sync()

  if err2 != nil {
    _ = in.Close()
    _ = out.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from out.Sync() out=destination='%v' Error='%v'",
      dst, err2.Error())
    return
  }

  err2 = in.Close()

  if err2 != nil {
    _ = out.Close()

    err = fmt.Errorf(ePrefix+
      "Error returned from in.Close() in=source='%v' Error='%v'",
      src, err2.Error())

    return err
  }

  err2 = out.Close()

  if err2 != nil {

    err = fmt.Errorf(ePrefix+
      "Error returned from out.Close() out=destination='%v' Error='%v'",
      dst, err2.Error())

    return err
  }

  err = nil

  return err

}

// CreateFile - Wrapper function for os.Create. If the 'pathFileName' does
// not exist a type *PathError will be returned.
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

// DeleteDirPathAll - Wrapper function for RemoveAll
// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first
// error it encounters. If the path does not exist,
// RemoveAll returns nil (no error).
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
  // 'RemoveAll()' returns 'nil'.
  if !fh.DoesFileExist(pathDir) {
    // Doesn't exist. Nothing to do.
    return nil
  }

  err := os.RemoveAll(pathDir)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(pathDir). pathDir='%v'  Error='%v'", pathDir, err.Error())
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

// GetAbsPathFromFilePath - Supply a string containing both
// the path file name and extension and return the path
// element.
func (fh FileHelper) GetAbsPathFromFilePath(filePath string) (string, error) {
  ePrefix := "FileHelper.GetAbsPathFromFilePath() "

  if len(filePath) == 0 {
    return "", errors.New(ePrefix + "Error: Input parameter 'filePath' is an EMPTY string!")
  }

  testFilePath := fh.AdjustPathSlash(filePath)

  if len(testFilePath) == 0 {
    return "", errors.New(ePrefix + "Error: After adjusting path Separators, filePath resolves to an empty string!")
  }

  absPath, err := fh.MakeAbsolutePath(testFilePath)

  if err != nil {
    return "", fmt.Errorf(ePrefix + "Error returned from ")
  }

  return absPath, nil
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
func (fh FileHelper) GetFileExtension(pathFileNameExt string) (ext string, isEmpty bool, err error) {
  ePrefix := "FileHelper.GetFileExt() "

  ext = ""
  isEmpty = true
  err = nil

  pathFileNameExt = strings.TrimLeft(strings.TrimRight(pathFileNameExt, " "), " ")

  if len(pathFileNameExt) == 0 {
    err = errors.New(ePrefix + "Error: After trimming 'pathFileNameExt'. Input parameter 'pathFileNameExt' is a Zero length string!")
    return
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  lenTestPathFileNameExt := len(testPathFileNameExt)

  if lenTestPathFileNameExt == 0 {
    err = errors.New(ePrefix + "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO length string!")
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
    return
  }

  lenDotIdxs := len(dotIdxs)

  // Deal with case where the pathFileNameExt contains
  // no dots.
  if lenDotIdxs == 0 {
    ext = ""
    isEmpty = true
    err = nil
    return

  }

  firstGoodCharIdx, lastGoodCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
    return
  }

  // Deal with the case where pathFileNameExt contains no
  // valid alpha numeric characters
  if firstGoodCharIdx == -1 || lastGoodCharIdx == -1 {
    ext = ""
    isEmpty = true
    err = nil
    return
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    ext = ""
    isEmpty = true
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
    return
  }

  lenSlashIdxs := len(slashIdxs)

  if lenSlashIdxs == 0 &&
    lenDotIdxs == 1 &&
    dotIdxs[lenDotIdxs-1] == 0 {
    // deal with the case .gitignore
    ext = ""
    isEmpty = true
    err = nil
    return
  }

  if lenSlashIdxs == 0 {
    ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
    isEmpty = false
    err = nil
    return
  }

  // lenDotIdxs and lenSlasIdxs both greater than zero
  if dotIdxs[lenDotIdxs-1] > slashIdxs[lenSlashIdxs-1] &&
    dotIdxs[lenDotIdxs-1] < lastGoodCharIdx {

    ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
    isEmpty = false
    err = nil
    return

  }

  ext = ""
  isEmpty = true
  err = nil
  return
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
// string is empty, default time format will be used to format the
// returned time string.
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
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")
  }

  if errCode == -2 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  volName := fh.GetVolumeName(testPathFileNameExt)

  if volName != "" {
    testPathFileNameExt = strings.TrimPrefix(testPathFileNameExt, volName)
  }

  lTestPathFileNameExt := len(testPathFileNameExt)

  if lTestPathFileNameExt == 0 {
    err = errors.New(ePrefix +
      "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO Length string!")
    return
  }

  firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return
  }

  // There are no alpha numeric characters present.
  // Therefore, there is no file name and extension
  if firstCharIdx == -1 || lastCharIdx == -1 {
    isEmpty = true
    err = nil
    return
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt). "+
      "testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
    return

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
    return
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
    return
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

  return
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

  isEmpty = true
  fName = ""
  err = nil
  errCode := 0

  errCode, _, pathFileNameExt = fh.isStringEmptyOrBlank(pathFileNameExt)

  if errCode == -1 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is an empty string!")
  }

  if errCode == -2 {
    return "", true,
      errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' consists of blank spaces!")
  }

  testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

  if len(testPathFileNameExt) == 0 {
    err = errors.New(ePrefix +
      "Error: Adjusted path version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO Length string!")
    return
  }

  fileNameExt, isFileNameExtEmpty, err2 := fh.GetFileNameWithExt(testPathFileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetFileNameWithExt(testPathFileNameExt) testPathFileNameExt='%v'  Error='%v'",
      testPathFileNameExt, err2.Error())
    return
  }

  if isFileNameExtEmpty {
    isEmpty = true
    fName = ""
    err = nil
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(fileNameExt)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned from fh.GetDotSeparatorIndexesInPathStr(fileNameExt). fileNameExt='%v'  Error='%v'",
      fileNameExt, err2.Error())
    return
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
    return
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
    return
  }

  // Secondary Case: filename
  fName = fileNameExt

  if fName == "" {
    isEmpty = true
  } else {
    isEmpty = false
  }

  err = nil
  return
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
  lPathStr := len(pathStr)
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

  pathStr = fp.FromSlash(pathStr)

  lPathStr = len(pathStr)

  if lPathStr == 0 {

    err = fmt.Errorf(ePrefix + "Error: After path Separator adjustment, 'pathStr' is a Zero length string!")

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

  return
}

// GetLastPathElement - Analyzes a 'pathName' string and returns the last
// element in the path. If 'pathName' ends in a path separator ('/'), this
// method returns an empty string.
//
// Example:
// pathName = '../dir1/dir2/fileName.ext' will return "fileName.ext"
// pathName = '../dir1/dir2/' will return ""
// pathName = 'fileName.ext' will return "fileName.ext"
// pathName = '../dir1/dir2/dir3' will return "dir3"
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
// a given directory path.
func (fh FileHelper) GetVolumeName(pathStr string) string {

  errCode := 0

  errCode, _, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode < 0 {
    return ""
  }

  return fp.VolumeName(pathStr)
}

// GetVolumeSeparatorIdxInPathStr - Returns the index of the
// Windows volume separator from an path string.
//
func (fh FileHelper) GetVolumeSeparatorIdxInPathStr(
  pathStr string) (volIdx int, err error) {

  ePrefix := "FileHelper.GetVolumeSeparatorIdxInPathStr()"

  volIdx = -1
  err = nil
  errCode := 0
  lPathStr := 0

  errCode, lPathStr, pathStr = fh.isStringEmptyOrBlank(pathStr)

  if errCode == -1 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' is an empty string!")
    return volIdx, err
  }

  if errCode == -2 {
    err =
      errors.New(ePrefix + "Error: Input parameter 'pathFileName' consists of blank spaces!")
    return volIdx, err
  }

  for i := 0; i < lPathStr; i++ {

    if rune(pathStr[i]) == ':' {
      volIdx = i
      err = nil
      return volIdx, err
    }
  }

  volIdx = -1
  err = nil
  return volIdx, err
}

// IsAbsolutePath - Wrapper function for path.IsAbs()
// https://golang.org/pkg/path/#IsAbs
// This method reports whether the input parameter is
// an absolute path.
func (fh FileHelper) IsAbsolutePath(pathStr string) bool {
  return path.IsAbs(pathStr)
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
//  pathFileStr   string    - The string to be analyzed.
//
// Return Values:
//
//  isPathFileStr   bool    - A boolean indicating whether the input parameter
//                            'pathFileStr' is in fact both a directory path and file name.
//
//  cannotDetermine bool    - A boolean value indicating whether the method could or
//                            could NOT determine whether input parameter 'pathFileStr'
//                            is a valid directory path and file name.
//
//                            'cannotDetermine' will be set to 'true' if 'pathFileStr'
//                            does not currently exist on disk and 'pathFileStr' is formatted
//                            like the following example:
//                                   "D:\\dirA\\common"
//                            In this example, the method cannot determine if 'common'
//                            is a file name or a directory name.
//
//  testPathFileStr string  - Input parameter 'pathFileStr' is subjected to cleaning routines
//                            designed to exclude extraneous characters from the analysis.
//                            'testPathFileStr' is the actual string on which the analysis was
//                            performed.
//
//
//  err             error   - If an error is encountered during processing, it is returned here.
//                            If no error occurs, 'err' is set to 'nil'.
//
func (fh FileHelper) IsPathFileString(pathFileStr string) (isPathFileStr bool, cannotDetermine bool, testPathFileStr string, err error) {

  ePrefix := "FileHelper.IsPathFileString() "

  if len(pathFileStr) == 0 {
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    testPathFileStr = ""
    err = errors.New(ePrefix + "Error - Zero Length input parameter 'pathFileStr'.")
    return
  }

  testPathFileStr = fp.FromSlash(pathFileStr)
  lTestPathStr := len(testPathFileStr)

  if lTestPathStr == 0 {
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    testPathFileStr = ""
    err = fmt.Errorf(ePrefix+"Error - fp.Clean(fp.FromSlash(pathFileStr)) yielded a Zero Length String. pathFileStr='%v'", pathFileStr)
    return
  }

  // See if path actually exists on disk and
  // then examine the File Info object returned.
  fInfo, err2 := os.Stat(testPathFileStr)

  if err2 == nil {

    if !fInfo.IsDir() {
      isPathFileStr = true
      cannotDetermine = false // High confidence in result
      err = nil
      return

    } else {
      isPathFileStr = false
      cannotDetermine = false // High confidence in result
      err = nil
      return
    }

  }

  // Ok - We know the testPathFileStr does NOT exist on disk

  if strings.Contains(testPathFileStr, "...") {
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    err = fmt.Errorf(ePrefix+"Error: INVALID PATH STRING! testPathFileStr='%v'", testPathFileStr)
    return

  }

  firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileStr)

  if err2 != nil {
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    err = fmt.Errorf(ePrefix+"Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileStr) testPathFileStr='%v'  Error='%v'", testPathFileStr, err2.Error())
    return
  }

  if firstCharIdx == -1 || lastCharIdx == -1 {
    // The pathfilestring contains no alpha numeric characters.
    // Therefore, it does NOT contain a file name!
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  volName := fp.VolumeName(testPathFileStr)

  if volName == testPathFileStr {
    // This is a volume name not a file Name!
    isPathFileStr = false
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileStr)

  if err2 != nil {
    isPathFileStr = false
    cannotDetermine = true
    err = fmt.Errorf(ePrefix+"fh.GetPathSeparatorIndexesInPathStr(testPathFileStr) returned error. testPathFileStr='%v' Error='%v'", testPathFileStr, err2.Error())
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileStr)

  if err2 != nil {
    isPathFileStr = false
    cannotDetermine = true // Uncertain outcome. Cannot determine if this is a path string
    err = fmt.Errorf(ePrefix+"fh.GetDotSeparatorIndexesInPathStr(testPathFileStr) retured error. testPathFileStr='%v' Error='%v'", testPathFileStr, err2.Error())
    return
  }

  lenDotIdx := len(dotIdxs)

  lenSlashIdx := len(slashIdxs)

  if lenSlashIdx == 0 {

    isPathFileStr = false
    cannotDetermine = false // high degree of certainty
    err = nil
    return

  }

  // We know the string contains one or more path separators

  if lenDotIdx == 0 {

    if lastCharIdx > slashIdxs[lenSlashIdx-1] {
      // Example D:\dir1\dir2\xray

      isPathFileStr = true
      cannotDetermine = true // Maybe, really can't tell if xray is a directory or a file!
      err = nil
      return

    }

    isPathFileStr = false
    cannotDetermine = false // high degree of certainty
    err = nil
    return

  }

  // We know that the test string contains both path separators and
  // dot separators ('.')

  if dotIdxs[lenDotIdx-1] > slashIdxs[lenSlashIdx-1] &&
    lastCharIdx > slashIdxs[lenSlashIdx-1] {
    isPathFileStr = true
    cannotDetermine = false // high degree of certainty
    err = nil
    return
  }

  // Check to determine if last character in testPathFileStr is a PathSeparator
  if slashIdxs[lenSlashIdx-1] == lTestPathStr-1 {
    // Yes, last char in testPathFileStr is a PathSeparator. This must be a directory.
    isPathFileStr = false
    cannotDetermine = false // high degree of certainty
    err = nil
    return
  }

  // Cannot be certain of the result.
  // Don't know for sure what this string is
  isPathFileStr = false
  cannotDetermine = true
  err = nil
  return
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

  var fInfo os.FileInfo

  lpathStr := len(pathStr)

  if lpathStr == 0 {
    isPathStr = false
    cannotDetermine = false
    testPathStr = ""
    err = errors.New(ePrefix + "Error - Zero Length input parameter 'pathStr'.")
    return
  }

  testPathStr = fp.FromSlash(pathStr)
  lTestPathStr := len(testPathStr)

  if lTestPathStr == 0 {
    isPathStr = false
    cannotDetermine = false
    testPathStr = ""
    err = fmt.Errorf(ePrefix+"Error - fp.FromSlash(pathStr) yielded a Zero Length String. pathStr='%v'", pathStr)
    return
  }

  // See if path actually exists on disk and
  // then examine the File Info object returned.
  fInfo, err = os.Stat(testPathStr)

  if err == nil {

    if fInfo.IsDir() {
      isPathStr = true
      cannotDetermine = false
      err = nil
      return

    } else {
      isPathStr = false
      cannotDetermine = false
      err = nil
      return
    }

  }

  // Ok - We know the testPathStr does NOT exist on disk

  if strings.Contains(testPathStr, "...") {
    // This is an INVALID path String
    isPathStr = false
    cannotDetermine = false
    err = fmt.Errorf("Error: INVALID PATH String! testPathStr='%v' ", testPathStr)
    return
  }

  volName := fp.VolumeName(testPathStr)

  if testPathStr == volName {
    isPathStr = true
    cannotDetermine = false
    err = nil
    return
  }

  _, checkPathIsEmpty, err2 := fh.GetPathFromPathFileName(testPathStr)

  if err2 != nil {
    isPathStr = false
    cannotDetermine = false
    err = fmt.Errorf(ePrefix+"fh.GetPathFromPathFileName(testPathStr) returned error. testPathStr='%v' Error='%v'", testPathStr, err2.Error())
    return
  }

  if checkPathIsEmpty {
    isPathStr = false
    cannotDetermine = false
    err = nil
    return
  }

  slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    isPathStr = false
    cannotDetermine = false
    err = fmt.Errorf(ePrefix+"fh.GetPathSeparatorIndexesInPathStr(testPathStr) returned error. testPathStr='%v' Error='%v'", testPathStr, err2.Error())
    return
  }

  dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathStr)

  if err2 != nil {
    isPathStr = false
    cannotDetermine = true // Uncertain outcome. Cannot determine if this is a path string
    err = fmt.Errorf(ePrefix+"fh.GetDotSeparatorIndexesInPathStr(testPathStr) retured error. testPathStr='%v' Error='%v'", testPathStr, err2.Error())
    return
  }

  firstNonSepCharIdx, lastNonSepCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr)

  if err2 != nil {
    isPathStr = false
    cannotDetermine = true // Uncertain outcome.
    err = fmt.Errorf(ePrefix+"fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr) retured error. testPathStr='%v' Error='%v'", testPathStr, err2.Error())
    return
  }

  if firstNonSepCharIdx == -1 || lastNonSepCharIdx == -1 {
    // All the characters are separator characters.
    isPathStr = true
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  // *******************************
  // From here on fristNonSepCharIdx
  // and lastNonSepCharIdx MUST be
  // greater than -1
  // *******************************

  lenDotIdx := len(dotIdxs)
  lenSlashIdx := len(slashIdxs)

  // Address case "../common"
  if strings.HasPrefix(testPathStr, "..") {

    if lenDotIdx == 2 {
      isPathStr = true
      cannotDetermine = false // High confidence in result
      err = nil
      return
    }

  }

  if strings.HasPrefix(testPathStr, ".") {

    if lenDotIdx == 1 {
      isPathStr = true
      cannotDetermine = false // High confidence in result
      err = nil
      return
    }

  }

  if lenDotIdx == 0 && lenSlashIdx == 0 {
    // Just text name with no path separators
    // and no dots. This is not a path
    isPathStr = false
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  // Address Case No Slashes only Dots
  if lenDotIdx > 0 && lenSlashIdx == 0 {

    // .common
    if dotIdxs[lenDotIdx-1] < firstNonSepCharIdx {
      isPathStr = true
      cannotDetermine = false // High confidence in result
      err = nil
      return
    }

    // common.
    if dotIdxs[lenDotIdx-1] > firstNonSepCharIdx {
      isPathStr = false
      cannotDetermine = false // High confidence in result
      err = nil
      return
    }

  }

  // Address Case No Dots only slashes
  if lenDotIdx == 0 && lenSlashIdx > 0 {

    isPathStr = true
    cannotDetermine = false // High confidence in result
    err = nil
    return

  }

  // ***********************************
  // Both lenDotIdx and lenSlashIdx are
  // greater than zero. Therefore, both
  // path separators and dot separators
  // are present.
  // ***********************************

  // Address Case: PathSeparator at end of PathStr
  if lTestPathStr-1 == slashIdxs[lenSlashIdx-1] {

    // There is a slash at the end of the path string.
    // This is definitely a path string.
    isPathStr = true
    cannotDetermine = false // High confidence in result
    err = nil
    return

  }

  // Address Case where last dot comes after last path separator
  if dotIdxs[lenDotIdx-1] > firstNonSepCharIdx &&
    dotIdxs[lenDotIdx-1] > slashIdxs[lenSlashIdx-1] {

    // This is a PathFileName string - NOT a PathStr
    isPathStr = false
    cannotDetermine = false // High confidence in result
    err = nil
    return

  }

  // Address case where last path separator comes after
  // the last dot.
  if slashIdxs[lenSlashIdx-1] > dotIdxs[lenDotIdx-1] {
    // This is a PathStr
    isPathStr = true
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  // Address case "/common/xray.txt"
  if lastNonSepCharIdx > dotIdxs[lenDotIdx-1] &&
    lastNonSepCharIdx > slashIdxs[lenDotIdx-1] &&
    dotIdxs[lenDotIdx-1] > slashIdxs[lenDotIdx-1] {

    isPathStr = false
    cannotDetermine = false // High confidence in result
    err = nil
    return
  }

  // Address case "..\dirA\dirB\xray"
  // In this method we will assume that
  // xray is a directory

  if dotIdxs[lenDotIdx-1] < slashIdxs[lenDotIdx-1] &&
    slashIdxs[lenDotIdx-1] < lastNonSepCharIdx {

    isPathStr = true
    cannotDetermine = true // Can't be 100% certain that xray
    // is not a file name.
    err = nil
    return
  }

  // Can't be certain what this string is.
  // could be either directory path or
  // directory path and file name. Let
  // calling method make the call.
  // Example D:\\DirA\\common
  // Is common a file name or a directory name.
  isPathStr = false
  cannotDetermine = true
  err = nil
  return
}

// JoinPathsAdjustSeparators - Joins two
// path strings and standardizes the
// path separators according to the
// current operating system.
func (fh FileHelper) JoinPathsAdjustSeparators(p1 string, p2 string) string {
  ps1 := fp.FromSlash(fp.Clean(p1))
  ps2 := fp.FromSlash(fp.Clean(p2))
  return fp.Clean(fp.FromSlash(path.Join(ps1, ps2)))

}

// JoinPaths - correctly joins 2-paths
func (fh FileHelper) JoinPaths(p1 string, p2 string) string {

  return fp.Clean(path.Join(fp.Clean(p1), fp.Clean(p2)))

}

// MakeAbsolutePath - Supply a relative path or any path
// string and resolve that path to an Absolute path.
// Note: Clean() is called on result by fp.Abs().
func (fh FileHelper) MakeAbsolutePath(relPath string) (string, error) {

  ePrefix := "FileHelper.MakeAbsolutePath() "

  if len(relPath) == 0 {
    return "", errors.New(ePrefix + "Error: Input Parameter 'relPath' is an EMPTY string!")
  }

  testRelPath := fh.AdjustPathSlash(relPath)

  if len(testRelPath) == 0 {
    return "", errors.New(ePrefix +
      "Error: Input Parameter 'relPath' adjusted for path Separators is an EMPTY string!")
  }

  p, err := fp.Abs(testRelPath)

  if err != nil {
    return "Invalid p!", fmt.Errorf(ePrefix+"Error returned from  fp.Abs(testRelPath). testRelPath='%v'  Error='%v'", testRelPath, err.Error())
  }

  return p, err
}

// MakeDirAll - creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error. The permission bits perm
// are used for all directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
func (fh FileHelper) MakeDirAll(dirPath string) error {
  var ModePerm os.FileMode = 0777
  return os.MkdirAll(dirPath, ModePerm)
}

// MakeDir - Makes a directory. Returns
// boolean value of false plus error if
// the operation fails. If successful,
// the function returns true.
func (fh FileHelper) MakeDir(dirPath string) (bool, error) {
  var ModePerm os.FileMode = 0777
  err := os.Mkdir(dirPath, ModePerm)

  if err != nil {
    return false, err
  }

  return true, nil
}

// MoveFile - Copies file from source to destination and, if
// successful, then deletes the original source file.
//
// The copy procedure will first attempt to the the 'Copy By Link' technique.
// See FileHelper.CopyFileByLink().  If this fails, the method will seamlessly
// attempt to copy the file the source file to the destination file by means
// of writing the contents of the source file to a newly created destination
// file. Reference Method FileHelper.CopyFileByIo().
//
// If an error is encountered during this procedure it will be by means of the
// return parameter 'err'.
//
// A boolean value is also returned. If 'copyByLink' is 'true', it signals that
// the move operation was accomplished using the 'CopyFileByLink' technique. If
// the return parameter 'copyByLink' is 'false', it signals that the 'CopyFileByIo'
// technique was used.
//
func (fh FileHelper) MoveFile(src, dst string) (copyByLink bool, err error) {
  ePrefix := "FileHelper.MoveFile() "
  copyByLink = true

  if len(src) == 0 {
    err = errors.New(ePrefix + "Error: Input parameter 'src' is ZERO length string!")
    return
  }

  if len(dst) == 0 {
    err = errors.New(ePrefix + "Error: Input parameter 'dst' is a ZERO length string!")
    return
  }

  if !fh.DoesFileExist(src) {
    err = fmt.Errorf(ePrefix+"Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
    return
  }

  err2 := fh.CopyFileByIo(src, dst)

  if err2 != nil {

    err2 = fh.CopyFileByLink(src, dst)

    if err2 != nil {

      err = fmt.Errorf(ePrefix+"Error returned from fh.CopyFileByLink(src, dst). Error='%v'",
        err2.Error())
    }

    copyByLink = true

    err = nil

    return copyByLink, err
  }

  copyByLink = false

  err2 = fh.DeleteDirFile(src)

  if err2 != nil {
    err = fmt.Errorf("Successfully copied file from source, '%v', to destination '%v'; "+
      "however deletion of source file failed! Error: %v", src, dst, err2.Error())

    return
  }

  err = nil

  return copyByLink, err
}

// OpenFile - wrapper for os.OpenFile. This method may be used to open or
// create files depending on the File Open and File Permission parameters.
//
func (fh FileHelper) OpenFile(
  targetPathFileName string,
  fileOpenCfg FileOpenConfig,
  filePermissionCfg FilePermissionConfig) (filePtr *os.File, err error) {

  filePtr = nil
  err = nil
  errCode := 0
  ePrefix := "FileHelper.OpenFile() "

  errCode, _, targetPathFileName = fh.isStringEmptyOrBlank(targetPathFileName)

  if errCode == -1 {
    err = errors.New(ePrefix + "Input parameter 'targetPathFileName' is an empty string!")
    return filePtr, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Input parameter 'targetPathFileName' consists of all spaces!")
    return filePtr, err
  }

  fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v", err2.Error())
    return filePtr, err
  }

  fileMode, err2 := filePermissionCfg.GetCompositePermissionMode()

  filePtr, err2 = os.OpenFile(targetPathFileName, fOpenCode, fileMode)

  if err2 != nil {

    if os.IsNotExist(err2) {
      err = fmt.Errorf(ePrefix+"The 'targetPathFileName' DOES NOT EXIST! "+
        "targetPathFileName='%v' Error='%v' ",
        targetPathFileName, err2.Error())
      filePtr = nil
      return filePtr, err
    }

    err = fmt.Errorf(ePrefix+
      "Error returned by os.OpenFile(targetPathFileName, fOpenCode, fileMode) "+
      "targetpathFileName='%v' Error='%v' ", targetPathFileName, err.Error())

    return filePtr, err
  }

  err = nil

  return filePtr, err
}

// OpenFileForReading - Wrapper function for os.Open() method which opens
// files on disk. 'Open' opens the named file for reading.
// If successful, methods on the returned file can be used for reading;
// the associated file descriptor has mode O_RDONLY. If there is an error,
// it will be of type *PathError. (See CreateThisFile() above.
func (fh FileHelper) OpenFileForReading(fileName string) (filePtr *os.File, err error) {
  ePrefix := "FileHelper.OpenFileForReading() "

  filePtr = nil
  err = nil
  var err2 error
  errCode := 0

  errCode, _, fileName = fh.isStringEmptyOrBlank(fileName)

  if errCode == -1 {
    err = errors.New(ePrefix + "Input parameter 'fileName' is an empty string!")
    return filePtr, err
  }

  if errCode == -2 {
    err = errors.New(ePrefix +
      "Input parameter 'fileName' consists of all spaces!")
    return filePtr, err
  }

  filePtr, err2 = os.Open(fileName)

  if err2 != nil && os.IsNotExist(err2) {
    err = fmt.Errorf(ePrefix+"Input parameter 'fileName' does NOT exist! "+
      "fileName='%v' os.Open(fileName) Error='%v' ",
      fileName, err2.Error())
    filePtr = nil
    return filePtr, err
  }

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"Error returned by os.Open(fileName). "+
      "fileName='%v' Error='%v' ", fileName, err2.Error())
  }

  err = nil

  return filePtr, err
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

// ReadFileBytes - Read bytes from file into a File Buffer.
func (fh FileHelper) ReadFileBytes(rFile *os.File, byteBuff []byte) (int, error) {
  return rFile.Read(byteBuff)
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
