package pathfileops

import (
  "errors"
  "fmt"
  "strings"
)

// FileOpsCollection - A collection of files and file operations which are designed
// to perform specific actions on disk files.
//
type FileOpsCollection struct {
  fileOps []FileOps
}

// AddByFileOps - Adds a FileOps object to the existing collection
// based on the 'FileOps' Input parameter.
func (fOpsCol *FileOpsCollection) AddByFileOps(fileOp FileOps) error {

  ePrefix := "FileOpsCollection.AddByFileOps() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  if !fileOp.IsInitialized() {

    return errors.New(ePrefix +
      "ERROR: Input parameter 'fileOp' is NOT initialized!\n")
  }

  fOpsCol.fileOps = append(fOpsCol.fileOps, fileOp.CopyOut())

  return nil
}

// AddByFileMgrs - Adds another FileOps object to the collection based source
// and destination input parameters of type 'FileMgr'.
//
func (fOpsCol *FileOpsCollection) AddByFileMgrs(
  sourceFileMgr,
  destinationFileMgr FileMgr) error {

  ePrefix := "FileOpsCollection.AddByFileMgrs() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  newFileOps, err := FileOps{}.NewByFileMgrs(sourceFileMgr, destinationFileMgr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileOps{}.NewByFileMgrs(sourceFileMgr, destinationFileMgr). "+
      "Error='%v' ", err.Error())
  }

  fOpsCol.fileOps = append(fOpsCol.fileOps, newFileOps)

  return nil
}

// AddByDirMgrFileName - Creates and Adds another FileOps object to the
// collection based on input parameters consisting of a pair of DirMgr
// and file name extension strings for source and destination.
//
func (fOpsCol *FileOpsCollection) AddByDirMgrFileName(
  sourceDirMgr DirMgr,
  sourceFileNameExt string,
  destinationDirMgr DirMgr,
  destinationFileNameExt string) error {

  ePrefix := "FileOpsCollection.AddByDirMgrFileName() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  newFileOps, err := FileOps{}.NewByDirMgrFileName(
    sourceDirMgr,
    sourceFileNameExt,
    destinationDirMgr,
    destinationFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileOps{}.NewByDirMgrFileName(...). "+
      "Error='%v' ", err.Error())
  }

  fOpsCol.fileOps = append(fOpsCol.fileOps, newFileOps)

  return nil
}

// AddByDirStrsAndFileNameExtStrs - Creates and adds another File Operations
// object to the collection based on two pairs of directory name and file name
// extension strings for both source and destination respectively.
//
func (fOpsCol *FileOpsCollection) AddByDirStrsAndFileNameExtStrs(
  sourceDirStr,
  sourceFileNameExtStr,
  destinationDirStr,
  destinationFileNameExtStr string) error {

  ePrefix := "FileOpsCollection.AddByDirStrsAndFileNameExtStrs() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  newFileOps, err :=
    FileOps{}.NewByDirStrsAndFileNameExtStrs(
      sourceDirStr,
      sourceFileNameExtStr,
      destinationDirStr,
      destinationFileNameExtStr)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs(...) "+
      "Error='%v' ", err.Error())
  }

  fOpsCol.fileOps = append(fOpsCol.fileOps, newFileOps)

  return nil
}

// AddByPathFileNameExtStrs - Creates and adds another File Operations
// object to the collection based on two input strings which contain the
// full path name, file name and file extension for the source and
// destination respectively.
//
func (fOpsCol *FileOpsCollection) AddByPathFileNameExtStrs(
  sourcePathFileNameExt,
  destinationPathFileNameExt string) error {

  ePrefix := "FileOpsCollection.AddByPathFileNameExtStrs() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  newFileOps, err :=
    FileOps{}.NewByPathFileNameExtStrs(
      sourcePathFileNameExt, destinationPathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by FileOps{}.NewByPathFileNameExtStrs(...) "+
      "sourcePathFileNameExt='%v' destinationPathFileNameExt='%v' Error='%v' ",
      sourcePathFileNameExt, destinationPathFileNameExt, err.Error())
  }

  fOpsCol.fileOps = append(fOpsCol.fileOps, newFileOps)

  return nil
}

// CopyOut - Returns an FileMgrCollection which is an
// exact duplicate of the current FileMgrCollection.
// The copy is operation is a 'deep copy'.
func (fOpsCol *FileOpsCollection) CopyOut() (FileOpsCollection, error) {

  ePrefix := "FileOpsCollection.CopyOut() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  fOpsCol2 := FileOpsCollection{}

  fOpsCol2.fileOps = make([]FileOps, 0, 100)

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return FileOpsCollection{},
      errors.New(ePrefix +
        "Error: This File Operations Collection ('FileOpsCollection') is EMPTY! ")
  }

  for i := 0; i < arrayLen; i++ {

    err := fOpsCol2.AddByFileOps(fOpsCol.fileOps[i].CopyOut())

    if err != nil {
      return FileOpsCollection{},
        fmt.Errorf(ePrefix +
          "Error returned by fOpsCol2.AddByFileOps(fOp)\n" +
          "Index='%v'\nError='%v'\n",
          i, err.Error())
    }
  }

  return fOpsCol2, nil
}

// DeleteAtIndex - Deletes a member File Operations element
// from the collection at the index specified by input
// parameter, 'idx'.
//
// If successful, at the completion of this method, the File
// Operations Collection array will have a length which is one
// less than the starting array length.
//
func (fOpsCol *FileOpsCollection) DeleteAtIndex(idx int) error {

  ePrefix := "FileOpsCollection.DeleteAtIndex() "

  if idx < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is less than zero. "+
      "Index Out-Of-Range! idx='%v'", idx)
  }

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return errors.New(ePrefix +
      "Error: The File Operations Collection, 'FileOpsCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is greater than the "+
      "length of the collection index. Index Out-Of-Range! "+
      "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if arrayLen == 1 {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  } else if idx == 0 {
    // arrayLen > 1 and requested idx = 0
    fOpsCol.fileOps = fOpsCol.fileOps[1:]
  } else if idx == arrayLen-1 {
    // arrayLen > 1 and requested idx = last element index
    fOpsCol.fileOps = fOpsCol.fileOps[0 : arrayLen-1]
  } else {
    // arrayLen > 1 and idx is in between
    // first and last elements
    fOpsCol.fileOps =
      append(fOpsCol.fileOps[0:idx], fOpsCol.fileOps[idx+1:]...)
  }

  return nil
}

// Equal - Compares the input parameter FileOpsCollection to the current
// FileOpsCollection instance. If they are equal, this method returns
// true.
//
func (fOpsCol *FileOpsCollection) Equal(fOpsCol2 *FileOpsCollection) bool {

  if fOpsCol2 == nil {
    return false
  }

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  if fOpsCol2.fileOps == nil {
    fOpsCol2.fileOps = make([]FileOps, 0, 50)
  }

  if len(fOpsCol.fileOps) != len(fOpsCol2.fileOps) {
    return false
  }

  for i:=0; i < len(fOpsCol.fileOps); i++ {

    if !fOpsCol.fileOps[i].source.Equal(&fOpsCol2.fileOps[i].source) {
      return false
    }

    if !fOpsCol.fileOps[i].destination.Equal(&fOpsCol2.fileOps[i].destination) {
      return false
    }

    if int(fOpsCol.fileOps[i].opToExecute) != int(fOpsCol2.fileOps[i].opToExecute) {
      return false
    }

  }

  return true
}

// ExecuteFileOperations - Executes a file operation on
// each member of the File Operations Collection. Any
// errors are collected and returned in an error array.
//
// The type of file operation performed is specified by
// input parameter, 'fileOp'. 'fileOp' is of type
// 'FileOperationCode'.
//
func (fOpsCol *FileOpsCollection) ExecuteFileOperations(
  fileOp FileOperationCode) error {

  ePrefix := "FileOpsCollection.ExecuteFileOperation() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return errors.New(ePrefix +
      "Error: This File Operations Collection ('FileOpsCollection') is EMPTY! ")
  }

  var b strings.Builder

  _, err := fmt.Fprintf(&b, "%s Errors Returned by ExecuteFileOperations()", ePrefix)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by initial fmt.Fprint(). %v", err.Error())
  }

  errNo := 0

  for i := 0; i < arrayLen; i++ {

    err = fOpsCol.fileOps[i].ExecuteFileOperation(fileOp)

    if err != nil {

      errNo++

      _, err2 := fmt.Fprintf(&b, "%d. %v  ", errNo, err.Error())

      if err2 != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by fmt.Fprint(). %s", err2.Error())
      }

    }
  }

  if errNo > 0 {
    return errors.New(b.String())
  }

  return nil
}

// GetFileOpsAtIndex - If successful, this method returns a pointer to
// the FileOps instance at the array index specified. The 'Peek' and 'Pop'
// methods below return FileOps objects using a 'deep' copy and therefore
// offer better protection against data corruption.
//
func (fOpsCol *FileOpsCollection) GetFileOpsAtIndex(idx int) (*FileOps, error) {

  ePrefix := "FileOpsCollection.GetFileOpsAtIndex() "

  emptyFileOps := FileOps{}

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return &emptyFileOps,
      fmt.Errorf(ePrefix +
        "Error: This File Operations Collection ('FileOpsCollection') is EMPTY!")
  }

  if idx < 0 || idx >= arrayLen {

    return &emptyFileOps,
      fmt.Errorf(ePrefix+
        "Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
        "The minimum index is '0'. "+
        "The maximum index is '%v'. ", idx, arrayLen-1)

  }

  return &fOpsCol.fileOps[idx], nil
}

// GetNumOfFileOps - Returns the number of File Operations objects
// in the collection.  Effectively, this is the array length of
// internal field FileOpsCollection.fileOps.
//
func (fOpsCol *FileOpsCollection) GetNumOfFileOps() int {

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  return len(fOpsCol.fileOps)

}

// InsertFileOpsAtIndex - Inserts a new File Operations type ('FileOps') into
// the collection at array 'index'. The new File Operations instance is passed
// as input parameter 'fOps'.
//
// If input parameter 'index' is less than zero, an error will be returned. If
// 'index' exceeds the value of the last index in the collection, 'fOps' will be
// added to the end of the collection at the next legal index.
//
func (fOpsCol *FileOpsCollection) InsertFileOpsAtIndex(fOps FileOps, index int) error {

  ePrefix := "FileMgrCollection.InsertFileOpsAtIndex() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  if index < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'index' is LESS THAN ZERO! "+
      "index='%v' ", index)
  }

  lenfMgrs := len(fOpsCol.fileOps)

  if index >= lenfMgrs {
    fOpsCol.fileOps = append(fOpsCol.fileOps, fOps.CopyOut())
    return nil
  }

  newFileMgrs := make([]FileOps, 0, 100)

  if index == 0 {
    newFileMgrs = append(newFileMgrs, fOps.CopyOut())
    fOpsCol.fileOps = append(newFileMgrs, fOpsCol.fileOps...)
    return nil
  }

  newFileMgrs = append(newFileMgrs, fOpsCol.fileOps[index:]...)

  fOpsCol.fileOps = append(fOpsCol.fileOps[:index])
  fOpsCol.fileOps = append(fOpsCol.fileOps, fOps.CopyOut())
  fOpsCol.fileOps = append(fOpsCol.fileOps, newFileMgrs...)

  return nil
}

// New - Creates and returns a new, properly initialized
// instance of 'FileOpsCollection'
func (fOpsCol FileOpsCollection) New() FileOpsCollection {

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  newFileOpsCol := FileOpsCollection{}

  newFileOpsCol.fileOps = make([]FileOps, 0, 100)

  return newFileOpsCol
}

// NewFromFileMgrCollection - Creates and returns a new
// File Operations Collection ('FileOpsCollection')
// generated from an existing File Manger Collection
// ('FileMgrCollection') and a target base directory.
//
// The source files for the new File Operations Collection
// are taken from the input parameter 'fMgrCol', the
// incoming File Manager Collection.
//
// The destination files for the new File Operations Collection
// are created from the source file names. The destination file
// directories are created by substituting the target base
// directory ('targetBaseDir') for the source base directory
// ('sourceBaseDir') in the source directory tree.
//
// This substitution is helpful when copying one directory tree
// to another directory tree.
//
func (fOpsCol FileOpsCollection) NewFromFileMgrCollection(
  fMgrCol *FileMgrCollection,
  sourceBaseDir,
  targetBaseDir *DirMgr) (FileOpsCollection, error) {

  ePrefix := "FileOpsCollection.NewFromFileMgrCollection() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  if fMgrCol.fileMgrs == nil {
    fMgrCol.fileMgrs = make([]FileMgr, 0, 100)
  }

  srcBaseDir := strings.ToLower(sourceBaseDir.GetAbsolutePath())

  targBaseDir := targetBaseDir.GetAbsolutePath()

  srcBaseDirLen := len(srcBaseDir)

  arrayLen := fMgrCol.GetNumOfFileMgrs()

  newFileOpsCol := FileOpsCollection{}.New()

  newFileOpsCol.fileOps = make([]FileOps, 0, arrayLen+10)

  for i := 0; i < arrayLen; i++ {

    srcFMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

    if err != nil {
      return FileOpsCollection{},
        fmt.Errorf(ePrefix+
          "Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
          "i='%v' Error='%v' ", i, err.Error())
    }

    srcPathFileName := srcFMgr.GetAbsolutePathFileName()

    idx := strings.Index(strings.ToLower(srcPathFileName), srcBaseDir)

    if idx < 0 {
      return FileOpsCollection{},
        fmt.Errorf(ePrefix+
          "Error: Could not locate source base directory in source file path! "+
          "Source Base Directory:='%v' Source Path File Name='%v'",
          srcBaseDir, srcPathFileName)
    }

    // targetDir + pathFile[lenSrcBaseDir:]
    targetPathFileName := targBaseDir + srcPathFileName[srcBaseDirLen:]

    destFMgr, err := FileMgr{}.NewFromPathFileNameExtStr(targetPathFileName)

    if err != nil {
      return FileOpsCollection{},
        fmt.Errorf(ePrefix+
          "Error returned by FileMgr{}.NewFromPathFileNameExtStr(targetPathFileName). "+
          "targetPathFileName='%v' Error='%v' ", targetPathFileName, err.Error())
    }

    err = newFileOpsCol.AddByFileMgrs(srcFMgr, destFMgr)

    if err != nil {
      return FileOpsCollection{},
        fmt.Errorf(ePrefix+
          "Error returned by newFileOpsCol.AddByFileMgrs(srcFMgr, destFMgr). "+
          "srcFMgr='%v' destFMgr='%v' Error='%v' ",
          srcFMgr.GetAbsolutePathFileName(), destFMgr.GetAbsolutePathFileName(),
          err.Error())
    }
  }

  return newFileOpsCol, nil
}

// PopFileOpsAtIndex - Returns a copy of the File Operations (FileOps)
// object located at index, 'idx', in the File Operations Collection
// ('FileOpsCollection') array. As a 'Pop' method, the original File
// Operations ('FileOps') object is deleted from the File Operations
// Collection ('FileOpsCollection') array.
//
// Therefore at the completion of this method, the File Operations
// Collection array has a length which is one less than the starting
// array length.
//
func (fOpsCol *FileOpsCollection) PopFileOpsAtIndex(idx int) (FileOps, error) {

  ePrefix := "FileOpsCollection.PopFileOpsAtIndex() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  if idx < 0 {
    return FileOps{}, fmt.Errorf(
      ePrefix+
        "Error: Input Parameter is less than zero. "+
        "Index Out-Of-Range! idx='%v'", idx)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection, 'FileOpsCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return FileOps{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter is greater than the "+
        "length of the collection index. Index Out-Of-Range! "+
        "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if idx == 0 {
    return fOpsCol.PopFirstFileOps()
  }

  if idx == arrayLen-1 {
    return fOpsCol.PopLastFileOps()
  }

  fileOps := fOpsCol.fileOps[idx].CopyOut()

  fOpsCol.fileOps =
    append(fOpsCol.fileOps[0:idx], fOpsCol.fileOps[idx+1:]...)

  return fileOps, nil
}

// PopFirstFileOps  - Returns a deep copy of the first File Operations
// ('FileOps') object in the File Operations Collection array. As a
// 'Pop' method, the original File Operations ('FileOps') object is
// deleted from the File Operations Collection ('FileOpsCollection')
// array.
//
// Therefore at the completion of this method, the File Operations
// Collection array has a length which is one less than the starting
// array length.
//
func (fOpsCol *FileOpsCollection) PopFirstFileOps() (FileOps, error) {

  ePrefix := "DirMgrCollection.PopFirstDirMgr() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  if len(fOpsCol.fileOps) == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection is EMPTY!")
  }

  fileOps := fOpsCol.fileOps[0].CopyOut()

  fOpsCol.fileOps = fOpsCol.fileOps[1:]

  return fileOps, nil
}

// PopLastFileOps - Returns a deep copy of the last File Operations
// ('FileOps') object in the File Operations Collection array. As a
// 'Pop' method, the original File Operations ('FileOps') object is
// deleted from the File Operations Collection ('FileOpsCollection')
// array.
//
// Therefore, at the completion of this method, the File Operations
// Collection array has a length which is one less than the starting
// array length.
//
func (fOpsCol *FileOpsCollection) PopLastFileOps() (FileOps, error) {

  ePrefix := "FileOpsCollection.PopLastFileOps() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 100)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection, 'FileOpsCollection', is EMPTY!")
  }

  fileOps := fOpsCol.fileOps[arrayLen-1].CopyOut()

  fOpsCol.fileOps = fOpsCol.fileOps[0 : arrayLen-1]

  return fileOps, nil

}

// PeekFileOpsAtIndex - Returns a deep copy of the File Operations
// ('FileOps') object located at array index 'idx' in the File
// Operations Collection ('FileOpsCollection'). This is a 'Peek'
// method and therefore the original File Operations ('FileOps')
// object is NOT deleted from the File Operations Collection
// ('FileOpsCollection') array.
//
// At the completion of this method, the length of the File
// Operations Collection ('FileOpsCollection') array will remain
// unchanged.
//
func (fOpsCol *FileOpsCollection) PeekFileOpsAtIndex(idx int) (FileOps, error) {

  ePrefix := "FileOpsCollection.PeekFileOpsAtIndex() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  arrayLen := len(fOpsCol.fileOps)

  if arrayLen == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection, 'FileOpsCollection' is EMPTY!")
  }

  if idx < 0 {
    return FileOps{}, fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is less than zero. "+
      "Index Out-Of-Range! idx='%v'", idx)
  }

  if idx >= arrayLen {
    return FileOps{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter 'idx' is greater than the "+
        "length of the collection array. Index Out-Of-Range! "+
        "idx='%v' Array Length='%v' ",
        idx, arrayLen)

  }

  return fOpsCol.fileOps[idx].CopyOut(), nil
}

// PeekFirstFileOps - Returns a deep copy of the first File
// Operations ('FileOps') object in the File Operations Collection
// ('FileOpsCollection'). This is a 'Peek' method and therefore
// the original File Operations ('FileOps') object is NOT
// deleted from the File Operations Collection ('FileOpsCollection')
// array.
//
// At the completion of this method, the length of the File
// Operations Collection ('FileOpsCollection') array will remain
// unchanged.
//
func (fOpsCol *FileOpsCollection) PeekFirstFileOps() (FileOps, error) {

  ePrefix := "FileOpsCollection.PeekFirstFileOps() "

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  if len(fOpsCol.fileOps) == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection ('FileOpsCollection') is EMPTY!")
  }

  return fOpsCol.fileOps[0].CopyOut(), nil
}

// PeekLastFileOps - Returns a deep copy of the last File
// Operations ('FileOps') object in the File Operations
// Collection ('FileOpsCollection').
//
// This is a 'Peek' method and therefore the original File
// Operations ('FileOps') object is NOT deleted from the
// File Operations Collection ('FileOpsCollection') array.
//
// At the completion of this method, the length of the File
// Operations Collection ('FileOpsCollection') array will
// remain unchanged.
//
func (fOpsCol *FileOpsCollection) PeekLastFileOps() (FileOps, error) {

  ePrefix := "FileOpsCollection.PeekLastFileOps()"

  if fOpsCol.fileOps == nil {
    fOpsCol.fileOps = make([]FileOps, 0, 50)
  }

  arrayLen := len(fOpsCol.fileOps)
  if arrayLen == 0 {
    return FileOps{},
      errors.New(ePrefix +
        "Error: The File Operations Collection, 'FileOpsCollection' is EMPTY!")
  }

  return fOpsCol.fileOps[arrayLen-1].CopyOut(), nil
}

