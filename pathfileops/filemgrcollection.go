package pathfileops

import (
  "errors"
  "fmt"
  "os"
)

// FileMgrCollection - Manages a collection of FileMgr
// instances.
//
// Dependencies:
// 'FileMgrCollection' depends on type, 'FileHelper'
// which is located in source code file 'filehelper.go'.
type FileMgrCollection struct {
  fileMgrs []FileMgr
}

// AddFileMgr - Adds a FileMgr object to the collection
func (fMgrs *FileMgrCollection) AddFileMgr(fMgr FileMgr) {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
}

// AddFileMgrByDirFileNameExt - Add a new File Manager using
// input parameters 'directory' and 'pathFileNameExt'.
func (fMgrs *FileMgrCollection) AddFileMgrByDirFileNameExt(
  directory DirMgr,
  fileNameExt string) error {

  ePrefix := "FileMgrCollection.AddFileMgrByDirFileNameExt() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(directory, fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"%v", err.Error())
  }

  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

  return nil
}

// AddFileMgrByPathFileNameExt - Add a new File Manager based on
// input parameter 'pathFileNameExt' which includes the full path
// name, file name and file extension.
func (fMgrs *FileMgrCollection) AddFileMgrByPathFileNameExt(
  pathFileNameExt string) error {

  ePrefix := "FileMgrCollection.AddFileMgrByPathFileNameExt() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  fMgr, err := FileMgr{}.NewFromPathFileNameExtStr(pathFileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from FileMgr{}.NewFromPathFileNameExtStr(pathFileNameExt). "+
      "pathFileNameExt='%v' Error='%v'", pathFileNameExt, err.Error())
  }

  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

  return nil
}

// AddFileMgrByDirStrFileNameStr - Adds a FileMgr object to the
// collection based on input parameter strings, 'pathName' and
// 'fileNameExt'.
//
func (fMgrs *FileMgrCollection) AddFileMgrByDirStrFileNameStr(
  pathName string,
  fileNameExt string) error {

  ePrefix := "FileMgrCollection.AddFileMgrByDirStrFileNameStr() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  fMgr, err := FileMgr{}.NewFromDirStrFileNameStr(pathName, fileNameExt)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error creating FileMgr: %v", err.Error())
  }

  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

  return nil

}

// AddFileMgrByFileInfo - Adds a File Manager object to the collection based on input from
// a directory path string and a os.FileInfo object.
func (fMgrs *FileMgrCollection) AddFileMgrByFileInfo(pathFile string, info os.FileInfo) error {

  ePrefix := "FileMgrCollection) AddFileMgrByFileInfo() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  fMgr, err := FileMgr{}.NewFromFileInfo(pathFile, info)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error retrned from FileMgr{}.NewFromFileInfo(pathFile, info). "+
      "pathFile='%v' info.Name()='%v'  Error='%v'", pathFile, info.Name(), err.Error())
  }

  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

  return nil
}

// AddFileMgrCollection - Adds another collection of File Manager (FileMgr)
// objects to the current collection.
func (fMgrs *FileMgrCollection) AddFileMgrCollection(fMgrs2 *FileMgrCollection) {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if fMgrs2.fileMgrs == nil {
    fMgrs2.fileMgrs = make([]FileMgr, 0, 50)
  }

  lOmc2 := len(fMgrs2.fileMgrs)

  if lOmc2 == 0 {
    return
  }

  for i := 0; i < lOmc2; i++ {
    fMgrs.AddFileMgr(fMgrs2.fileMgrs[i].CopyOut())
  }

  return
}

// CopyFilesToDir - Copies all the files in the File Manager Collection to
// the specified target directory.
//
func (fMgrs *FileMgrCollection) CopyFilesToDir(targetDirectory DirMgr) error {

  ePrefix := "FileMgrCollection.CopyFilesToDir() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  maxLen := len(fMgrs.fileMgrs)

  if maxLen == 0 {
    return errors.New(ePrefix + "ERROR - Collection contains ZERO File Managers!")
  }

  for i := 0; i < maxLen; i++ {
    err := fMgrs.fileMgrs[i].CopyFileToDirByIoByLink(targetDirectory)

    if err != nil {
      return fmt.Errorf(ePrefix+
        "Copy Failure on index='%v' file='%v'. Error='%v'",
        i, fMgrs.fileMgrs[i].absolutePathFileName, err.Error())
    }

  }

  return nil
}

// CopyOut - Returns an FileMgrCollection which is an
// exact duplicate of the current FileMgrCollection
func (fMgrs *FileMgrCollection) CopyOut() (FileMgrCollection, error) {

  ePrefix := "FileMgrCollection.CopyOut() "

  fMgrs2 := FileMgrCollection{}

  fMgrs2.fileMgrs = make([]FileMgr, 0, 50)

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  lOmc := len(fMgrs.fileMgrs)

  if lOmc == 0 {
    return FileMgrCollection{},
      errors.New(ePrefix +
        "Error: This File Manager Collection ('FileMgrCollection') is EMPTY! ")
  }

  for i := 0; i < lOmc; i++ {
    fMgrs2.AddFileMgr(fMgrs.fileMgrs[i].CopyOut())
  }

  return fMgrs2, nil
}

// DeleteAtIndex - Deletes a member File Manager from the
// collection at the index specified by input parameter 'idx'.
//
// If successful, at the completion of this method, the File
// Manager Collection array will have a length which is one
// less than the starting array length.
//
func (fMgrs *FileMgrCollection) DeleteAtIndex(idx int) error {

  ePrefix := "FileMgrCollection.DeleteAtIndex() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if idx < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is less than zero. "+
      "Index Out-Of-Range! idx='%v'", idx)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return errors.New(ePrefix +
      "Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is greater than the "+
      "length of the collection index. Index Out-Of-Range! "+
      "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if arrayLen == 1 {
    fMgrs.fileMgrs = make([]FileMgr, 0, 100)
  } else if idx == 0 {
    // arrayLen > 1 and requested idx = 0
    fMgrs.fileMgrs = fMgrs.fileMgrs[1:]
  } else if idx == arrayLen-1 {
    // arrayLen > 1 and requested idx = last element index
    fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]
  } else {
    // arrayLen > 1 and idx is in between
    // first and last elements
    fMgrs.fileMgrs =
      append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)
  }

  return nil
}

// FindFiles - Searches the current FileMgrCollection and returns a new
// FileMgrCollection containing FileMgr objects which match the specified
// search criteria.
//
func (fMgrs *FileMgrCollection) FindFiles(
  fileSelectionCriteria FileSelectionCriteria) (FileMgrCollection, error) {

  ePrefix := "FileMgrCollection.FindFiles() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  lDirCol := len(fMgrs.fileMgrs)

  if lDirCol == 0 {
    return FileMgrCollection{}, nil
  }

  fh := FileHelper{}

  var isMatchedFile bool
  var err error

  fMgrs2 := FileMgrCollection{}.New()

  for i := 0; i < lDirCol; i++ {
    fMgr := fMgrs.fileMgrs[i]

    if fMgr.actualFileInfo.IsFInfoInitialized {

      isMatchedFile, err = fh.FilterFileName(fMgr.actualFileInfo, fileSelectionCriteria)

      if err != nil {
        return FileMgrCollection{},
          fmt.Errorf(ePrefix+
            "Error returned by "+
            "fh.FilterFileName(fMgr.actualFileInfo, fileSelectionCriteria) "+
            "fMgr.actualFileInfo.Name()='%v'  Error='%v'",
            fMgr.actualFileInfo.Name(), err.Error())
      }

    } else {

      fip := FileInfoPlus{}

      fip.SetName(fMgr.fileNameExt)

      isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)

      if err != nil {
        return FileMgrCollection{}, fmt.Errorf(ePrefix+
          "Error returned by fh.FilterFileName(fip, fileSelectionCriteria) "+
          "fip.Name()='%v'  Error='%v'", fip.Name(), err.Error())
      }

    }

    if isMatchedFile {
      fMgrs2.AddFileMgr(fMgr)
    }

  }

  return fMgrs2, nil
}

// GetFileMgrArray - Returns the entire Directory Manager Array managed
// by this collection.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	[]FileMgr      - The array of of FileMgr instances maintained by this
//	                 collection.
//
func (fMgrs *FileMgrCollection) GetFileMgrArray() []FileMgr {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 10)
  }

  return fMgrs.fileMgrs

}

// GetFileMgrAtIndex - If successful, this method returns a pointer to
// the FileMgr instance at the array index specified. The 'Peek' and 'Pop'
// methods below return FileMgr objects using a 'deep' copy and therefore
// offer better protection against data corruption.
//
func (fMgrs *FileMgrCollection) GetFileMgrAtIndex(idx int) (*FileMgr, error) {

  ePrefix := "FileMgrCollection.GetFileMgrAtIndex() "

  emptyFileMgr := FileMgr{}

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return &emptyFileMgr,
      fmt.Errorf(ePrefix +
        "Error: This File Manager Collection ('FileMgrCollection') is EMPTY!")
  }

  if idx < 0 || idx >= arrayLen {

    return &emptyFileMgr,
      fmt.Errorf(ePrefix+
        "Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
        "The minimum index is '0'. "+
        "The maximum index is '%v'. ", idx, arrayLen-1)

  }

  return &fMgrs.fileMgrs[idx], nil

}

// GetNumOfFileMgrs - returns the array length of the
// of the File Manager Collection, 'FileMgrCollection'.
// Effectively the returned integer is a count of the
// number of File Managers (FileMgr's) in the Collection.
//
func (fMgrs *FileMgrCollection) GetNumOfFileMgrs() int {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  return len(fMgrs.fileMgrs)
}

// GetNumOfFiles - returns the array length of the
// of the File Manager Collection, 'FileMgrCollection'.
// Effectively the returned integer is a count of the
// number of files or File Managers (FileMgr's) in the
// Collection.
//
func (fMgrs *FileMgrCollection) GetNumOfFiles() int {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  return len(fMgrs.fileMgrs)
}

// InsertFileMgrAtIndex - Inserts a new File Manager into the collection at
// array 'index'. The new File Manager is passed as input parameter 'fMgr'.
//
// If input parameter 'index' is less than zero, an error will be returned. If
// 'index' exceeds the value of the last index in the collection, 'fMgr' will be
// added to the end of the collection at the next legal index.
//
func (fMgrs *FileMgrCollection) InsertFileMgrAtIndex(fMgr FileMgr, index int) error {

  ePrefix := "FileMgrCollection.InsertFileMgrAtIndex() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if index < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'index' is LESS THAN ZERO! "+
      "index='%v' ", index)
  }

  lenfMgrs := len(fMgrs.fileMgrs)

  if index >= lenfMgrs {
    fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
    return nil
  }

  newFileMgrs := make([]FileMgr, 0, 100)

  if index == 0 {
    newFileMgrs = append(newFileMgrs, fMgr.CopyOut())
    fMgrs.fileMgrs = append(newFileMgrs, fMgrs.fileMgrs...)
    return nil
  }

  newFileMgrs = append(newFileMgrs, fMgrs.fileMgrs[index:]...)

  fMgrs.fileMgrs = append(fMgrs.fileMgrs[:index])
  fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
  fMgrs.fileMgrs = append(fMgrs.fileMgrs, newFileMgrs...)

  return nil
}

// New - Creates and returns a new, empty and properly initialized
// File Manager Collection ('FileMgrCollection').
func (fMgrs FileMgrCollection) New() FileMgrCollection {

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  newFMgrCol := FileMgrCollection{}

  newFMgrCol.fileMgrs = make([]FileMgr, 0, 100)

  return newFMgrCol
}

// PopFileMgrAtIndex - Returns a deep copy of the File Manager
// ('FileMgr') object located at index, 'idx', in the
// File Manager Collection ('FileMgrCollection') array.
//
// As a 'Pop' method, the original File Manager ('FileMgr')
// object is deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// Therefore, at the completion of this method, the File Manager
// Collection array has a length which is one less than the
// starting array length.
//
func (fMgrs *FileMgrCollection) PopFileMgrAtIndex(idx int) (FileMgr, error) {

  ePrefix := "FileMgrCollection.PopFileMgrAtIndex() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if idx < 0 {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return FileMgr{}, fmt.Errorf(ePrefix+
      "Error: Input Parameter, 'idx' is greater than the length of the "+
      "collection index. Index Out-Of-Range! "+
      "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if idx == 0 {
    return fMgrs.PopFirstFileMgr()
  }

  if idx == arrayLen-1 {
    return fMgrs.PopLastFileMgr()
  }

  fmgr := fMgrs.fileMgrs[idx].CopyOut()

  fMgrs.fileMgrs = append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)

  return fmgr, nil
}

// PopFirstFileMgr - Returns a deep copy of the first File Manager
// ('FileMgr') object in the File Manager Collection array. As a
// 'Pop' method, the original File Manager ('FileMgr') object is
// deleted from the File Manager Collection ('FileMgrCollection')
// array.
//
// Therefore at the completion of this method, the File Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (fMgrs *FileMgrCollection) PopFirstFileMgr() (FileMgr, error) {

  ePrefix := "FileMgrCollection.PopFirstFileMgr() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if len(fMgrs.fileMgrs) == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: The File Manager Collection, 'FileMgrCollection' is EMPTY!")
  }

  fMgr := fMgrs.fileMgrs[0].CopyOut()

  fMgrs.fileMgrs = fMgrs.fileMgrs[1:]

  return fMgr, nil
}

// PopLastFileMgr - Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection array. As a
// 'Pop' method, the original File Manager ('FileMgr') object is
// deleted from the File Manager Collection ('FileMgrCollection')
// array.
//
// Therefore at the completion of this method, the File Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (fMgrs *FileMgrCollection) PopLastFileMgr() (FileMgr, error) {

  ePrefix := "FileMgrCollection.PopLastFileMgr() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return FileMgr{}, errors.New(ePrefix +
      "Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!")
  }

  fmgr := fMgrs.fileMgrs[arrayLen-1].CopyOut()

  fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]

  return fmgr, nil
}

// PeekFileMgrAtIndex - Returns a deep copy of the File Manager
// ('FileMgr') object located at array index 'idx' in the File
// Manager Collection ('FileMgrCollection'). This is a 'Peek'
// method and therefore the original File Manager ('FileMgr')
// object is NOT deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// At the completion of this method, the length of the File
// Manager Collection ('FileMgrCollection') array will remain
// unchanged.
//
func (fMgrs *FileMgrCollection) PeekFileMgrAtIndex(idx int) (FileMgr, error) {

  ePrefix := "FileMgrCollection.PeekFileMgrAtIndex() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: The File Manager Collection, 'FileMgrCollection' is EMPTY!")
  }

  if idx < 0 {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter 'idx' is less than zero. "+
        "Index Out-Of-Range! idx='%v'", idx)
  }

  if idx >= arrayLen {
    return FileMgr{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter 'idx' is greater than the length "+
        "of the collection array. "+
        "Index Out-Of-Range! idx='%v' Array Length='%v' ",
        idx, arrayLen)
  }

  return fMgrs.fileMgrs[idx].CopyOut(), nil
}

// PeekFirstFileMgr - Returns a deep copy of the first File
// Manager ('FileMgr') object in the File Manager Collection
// ('FileMgrCollection'). This is a 'Peek' method and therefore
// the original File Manager ('FileMgr') object is NOT deleted
// from the File Manager Collection ('FileMgrCollection')
// array.
//
// At the completion of this method, the length of the File
// Manager Collection ('FileMgrCollection') array will remain
// unchanged.
//
func (fMgrs *FileMgrCollection) PeekFirstFileMgr() (FileMgr, error) {

  ePrefix := "FileMgrCollection.PeekFirstFileMgr() "

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  if len(fMgrs.fileMgrs) == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: The File Manager Collection ('FileMgrCollection') is EMPTY!")
  }

  return fMgrs.fileMgrs[0].CopyOut(), nil
}

// PeekLastFileMgr - Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection
// ('FileMgrCollection').
//
// This is a 'Peek' method and therefore the original File Manager
// ('FileMgr') object is NOT deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// At the completion of this method, the length of the File Manager
// Collection ('FileMgrCollection') array will remain unchanged.
//
func (fMgrs *FileMgrCollection) PeekLastFileMgr() (FileMgr, error) {

  ePrefix := "FileMgrCollection.PeekLastFileMgr()"

  if fMgrs.fileMgrs == nil {
    fMgrs.fileMgrs = make([]FileMgr, 0, 50)
  }

  arrayLen := len(fMgrs.fileMgrs)

  if arrayLen == 0 {
    return FileMgr{},
      errors.New(ePrefix +
        "Error: The File Manager Collection ('FileMgrCollection') is EMPTY!")
  }

  return fMgrs.fileMgrs[arrayLen-1].CopyOut(), nil
}
