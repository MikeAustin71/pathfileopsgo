package pathfileops

import (
  "errors"
  "fmt"
  "os"
)


/*
	This source code file contains type 'DirMgrCollection' .

	The Source Repository for this source code file is :
		https://github.com/MikeAustin71/pathfileopsgo.git

	Dependencies:
	-------------

	Type 'DirMgrCollection'depend on types, 'FileHelper',
  'FileMgr' and 'DirMgr' which are contained in source
  code files, 'filehelper.go', 'filemanager.go' and
  'dirmanager.go' located in this directory.

*/

// DirMgrCollection - A collection of Type DirMgr. The collection
// is used to aid in the management of groups of paths and directories.
//
// Dependencies:
//
// Type 'DirMgrCollection' depend on types, 'FileHelper' and
// 'FileMgr' which are contained in source code files: 'filehelper.go'
// and 'filemanager.go' located in this directory.
//
type DirMgrCollection struct {
  dirMgrs []DirMgr
}

// AddDirMgr - Adds a DirMgr object to the collection
func (dMgrs *DirMgrCollection) AddDirMgr(dMgr DirMgr) {

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
}

// AddDirMgrByPathNameStr - Adds a Directory Manager (DirMgr) to the
// collections based on a string input parameter, 'pathName'.
//
func (dMgrs *DirMgrCollection) AddDirMgrByPathNameStr(pathName string) error {
  ePrefix := "DirMgrCollection.AddDirMgrByPathNameStr() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  dMgr, err := DirMgr{}.New(pathName)

  if err != nil {
    s := ePrefix +
      "Error returned from DirMgr{}.New(pathName). pathName='%v' Error='%v'"
    return fmt.Errorf(s, pathName, err.Error())
  }

  dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr)

  return nil
}

// AddFileMgrByFileInfo - Adds a Directory Manager object to the
// collection based on input from a directory path string and an
// os.FileInfo object.
//
func (dMgrs *DirMgrCollection) AddFileInfo(pathFile string, info os.FileInfo) error {

  ePrefix := "DirMgrCollection) AddFileMgrByFileInfo() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  dMgr, err := DirMgr{}.NewFromFileInfo(pathFile, info)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error retrned from DirMgr{}."+
      "NewFromFileInfo(pathFile, info). pathFile='%v' info.Name()='%v'  Error='%v'",
      pathFile, info.Name(), err.Error())
  }

  dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr)

  return nil
}

// AddDirMgrCollection - Adds another collection of File Manager (DirMgr)
// objects to the current collection.
func (dMgrs *DirMgrCollection) AddDirMgrCollection(dMgrs2 *DirMgrCollection) {

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  if dMgrs2.dirMgrs == nil {
    dMgrs2.dirMgrs = make([]DirMgr, 0, 100)
  }

  lOmc2 := len(dMgrs2.dirMgrs)

  if lOmc2 == 0 {
    return
  }

  for i := 0; i < lOmc2; i++ {
    dMgrs.AddDirMgr(dMgrs2.dirMgrs[i].CopyOut())
  }

  return
}

// CopyOut - Returns an DirMgrCollection which is an
// exact duplicate of the current DirMgrCollection
func (dMgrs *DirMgrCollection) CopyOut() (DirMgrCollection, error) {

  ePrefix := "DirMgrCollection.CopyOut() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  dMgrs2 := DirMgrCollection{}.New()

  lOmc := len(dMgrs.dirMgrs)

  if lOmc == 0 {
    return DirMgrCollection{},
      errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
  }

  for i := 0; i < lOmc; i++ {
    dMgrs2.AddDirMgr(dMgrs.dirMgrs[i].CopyOut())
  }

  return dMgrs2, nil
}

// DeleteAtIndex - Deletes a member Directory Manager from
// the collection at the index specified by input parameter 'idx'.
//
// If successful, at the completion of this method, the Directory
// Manager Collection array will have a length which is one less
// than the starting array length.
//
func (dMgrs *DirMgrCollection) DeleteAtIndex(idx int) error {

  ePrefix := "DirMgrCollection.DeleteAtIndex() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  if idx < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is less than zero. "+
      "Index Out-Of-Range! idx='%v'", idx)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return errors.New(ePrefix +
      "Error: The Directory Manager Collection, 'DirMgrCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is greater than the "+
      "length of the collection index. Index Out-Of-Range! "+
      "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if arrayLen == 1 {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  } else if idx == 0 {
    // arrayLen > 1
    dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
  } else if idx == arrayLen-1 {
    // arrayLen > 1
    dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
  } else {
    // arrayLen > 1 and idx is in between
    // first and last elements
    dMgrs.dirMgrs =
      append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)
  }

  return nil
}

// FindDirectories - searches through the DirMgrCollection to find
// DirMgr objects matching specified search criteria.
func (dMgrs *DirMgrCollection) FindDirectories(
  fileSelectionCriteria FileSelectionCriteria) (DirMgrCollection, error) {

  ePrefix := "DirMgrCollection.FindDirectories() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  lDirCol := len(dMgrs.dirMgrs)

  if lDirCol == 0 {
    return DirMgrCollection{}.New(), nil
  }

  fh := FileHelper{}

  var isMatchedFile bool
  var err error

  dMgrs2 := DirMgrCollection{}.New()

  for i := 0; i < lDirCol; i++ {
    dMgr := dMgrs.dirMgrs[i]

    if dMgr.actualDirFileInfo.IsFInfoInitialized {

      isMatchedFile, err = fh.FilterFileName(dMgr.actualDirFileInfo, fileSelectionCriteria)

      if err != nil {
        return DirMgrCollection{}, fmt.Errorf(ePrefix+"Error returned by fh.FilterFileName("+
          "dMgr.actualDirFileInfo, fileSelectionCriteria) dMgr.actualDirFileInfo.Name()='%v' "+
          "Error='%v'", dMgr.actualDirFileInfo.Name(), err.Error())
      }

    } else {

      fip := FileInfoPlus{}

      fip.SetName(dMgr.directoryName)

      isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)

      if err != nil {
        s := ePrefix +
          "Error returned by fh.FilterFileName(fip, fileSelectionCriteria) " +
          "fip.Name()='%v'  Error='%v'"
        return DirMgrCollection{}, fmt.Errorf(s, fip.Name(), err.Error())
      }

    }

    if isMatchedFile && err == nil {
      dMgrs2.AddDirMgr(dMgr)
    }

  }

  return dMgrs2, nil
}

// GetDirMgrArray - Returns the entire Directory Manager Array managed
// by this collection.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]DirMgr      - The array of of DirMgr instances maintained by this
//	                collection.
//
func (dMgrs *DirMgrCollection) GetDirMgrArray() []DirMgr {

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 10)
  }

  return dMgrs.dirMgrs
}

// GetDirMgrAtIndex - If successful, this method returns a pointer to
// the DirMgr instance at the array index specified. The 'Peek' and 'Pop'
// methods below return DirMgr objects using a 'deep' copy and therefore
// offer better protection against data corruption.
//
func (dMgrs *DirMgrCollection) GetDirMgrAtIndex(idx int) (*DirMgr, error) {

  ePrefix := "DirMgrCollection.GetDirMgrAtIndex() "

  emptyDirMgr := DirMgr{}

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return &emptyDirMgr,
      fmt.Errorf(ePrefix +
        "Error: This Directory Manager Collection ('DirMgrCollection') is EMPTY!")
  }

  if idx < 0 || idx >= arrayLen {

    return &emptyDirMgr,
      fmt.Errorf(ePrefix+
        "Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
        "The minimum index is '0'. "+
        "The maximum index is '%v'. ", idx, arrayLen-1)

  }

  return &dMgrs.dirMgrs[idx], nil
}

// GetNumOfDirs - returns the number of directories
// contained in this Directory Manager Collection.
//
func (dMgrs *DirMgrCollection) GetNumOfDirs() int {

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  return len(dMgrs.dirMgrs)
}

// InsertDirMgrAtIndex - Inserts a new Directory Manager into the collection at
// array 'index'. The new Directory Manager is passed as input parameter 'dMgr'.
//
// If input parameter 'index' is less than zero, an error will be returned. If
// 'index' exceeds the value of the last index in the collection, 'dMgr' will be
// added to the end of the collection at the next legal index.
//
func (dMgrs *DirMgrCollection) InsertDirMgrAtIndex(dMgr DirMgr, index int) error {

  ePrefix := "DirMgrCollection.InsertDirMgrAtIndex() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  if index < 0 {
    return fmt.Errorf(ePrefix+
      "Error: Input parameter 'index' is LESS THAN ZERO! "+
      "index='%v' ", index)
  }

  lenDgrs := len(dMgrs.dirMgrs)

  if index >= lenDgrs {
    dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
    return nil
  }

  newDirMgrs := make([]DirMgr, 0, 100)

  if index == 0 {
    newDirMgrs = append(newDirMgrs, dMgr.CopyOut())
    dMgrs.dirMgrs = append(newDirMgrs, dMgrs.dirMgrs...)
    return nil
  }

  newDirMgrs = append(newDirMgrs, dMgrs.dirMgrs[index:]...)

  dMgrs.dirMgrs = append(dMgrs.dirMgrs[:index])
  dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgr.CopyOut())
  dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgrs...)

  return nil
}

// New - Creates and returns a new and properly initialized
// Directory Manager Collection ('DirMgrCollection').
//
func (dMgrs DirMgrCollection) New() DirMgrCollection {

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  newDirMgrCol := DirMgrCollection{}
  newDirMgrCol.dirMgrs = make([]DirMgr, 0, 100)

  return newDirMgrCol
}

// PopDirMgrAtIndex - Returns a copy of the Directory Manager
// ('DirMgr') object located at index, 'idx', in the Directory
// Manager Collection ('DirMgrCollection') array.
//
// As a 'Pop' method, the original Directory Manager ('DirMgr')
// object is deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// Therefore at the completion of this method, the Directory
// Manager Collection array has a length which is one less
// than the starting array length.
//
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(idx int) (DirMgr, error) {

  ePrefix := "DirMgrCollection.PopDirMgrAtIndex() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  if idx < 0 {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter 'idx' is less than zero. "+
        "Index Out-Of-Range! idx='%v'", idx)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Directory Manager Collection, 'DirMgrCollection', is EMPTY!")
  }

  if idx >= arrayLen {
    return DirMgr{}, fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is greater than the "+
      "length of the collection index. Index Out-Of-Range! "+
      "idx='%v' Array Length='%v' ", idx, arrayLen)
  }

  if idx == 0 {
    return dMgrs.PopFirstDirMgr()
  }

  if idx == arrayLen-1 {
    return dMgrs.PopLastDirMgr()
  }

  dMgr := dMgrs.dirMgrs[idx].CopyOut()

  dMgrs.dirMgrs =
    append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)

  return dMgr, nil
}

// PopFirstDirMgr  - Returns a deep copy of the first Directory Manager
// ('DirMgr') object in the Directory Manager Collection array. As a
// 'Pop' method, the original Directory Manager ('DirMgr') object is
// deleted from the Directory Manager Collection ('DirMgrCollection')
// array.
//
// Therefore at the completion of this method, the Directory Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (dMgrs *DirMgrCollection) PopFirstDirMgr() (DirMgr, error) {

  ePrefix := "DirMgrCollection.PopFirstDirMgr() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Directory Manger Collection is EMPTY!")
  }

  dMgr := dMgrs.dirMgrs[0].CopyOut()

  if arrayLen == 1 {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)

  } else {
    // arrayLen > 1
    dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
  }

  return dMgr, nil
}

// PopLastDirMgr - Returns a deep copy of the last Directory Manager
// ('DirMgr') object in the Directory Manager Collection array. As a
// 'Pop' method, the original Directory Manager ('DirMgr') object is
// deleted from the Directory Manager Collection ('DirMgrCollection')
// array.
//
// Therefore at the completion of this method, the Directory Manager
// Collection array has a length which is one less than the starting
// array length.
//
func (dMgrs *DirMgrCollection) PopLastDirMgr() (DirMgr, error) {

  ePrefix := "DirMgrCollection.PopLastDirMgr() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
  }

  dmgr := dMgrs.dirMgrs[arrayLen-1].CopyOut()

  if arrayLen == 1 {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)

  } else {
    // arrayLen > 1
    dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
  }

  return dmgr, nil
}

// PeekDirMgrAtIndex - Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the Directory
// Manager Collection ('DirMgrCollection'). This is a 'Peek' method
// and therefore the original Directory Manager ('DirMgr') object
// is NOT deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(idx int) (DirMgr, error) {

  ePrefix := "DirMgrCollection.PeekDirMgrAtIndex() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
  }

  if idx < 0 {
    return DirMgr{}, fmt.Errorf(ePrefix+
      "Error: Input Parameter 'idx' is less than zero. "+
      "Index Out-Of-Range! idx='%v'", idx)
  }

  if idx >= arrayLen {
    return DirMgr{},
      fmt.Errorf(ePrefix+
        "Error: Input Parameter 'idx' is greater than the "+
        "length of the collection array. Index Out-Of-Range! "+
        "idx='%v' Array Length='%v' ",
        idx, arrayLen)

  }

  return dMgrs.dirMgrs[idx].CopyOut(), nil
}

// PeekFirstDirMgr - Returns a deep copy of the first Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection'). This is a 'Peek' method and therefore
// the original Directory Manager ('DirMgr') object is NOT
// deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekFirstDirMgr() (DirMgr, error) {

  ePrefix := "DirMgrCollection.PeekFirstDirMgr() "

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  if len(dMgrs.dirMgrs) == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Director Manager Collection ('DirMgrCollection') is EMPTY!")
  }

  return dMgrs.dirMgrs[0].CopyOut(), nil
}

// PeekLastDirMgr - Returns a deep copy of the last Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection').
//
// This is a 'Peek' method and therefore the original Directory
// Manager ('DirMgr') object is NOT deleted from the Directory
// Manager Collection ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
//
func (dMgrs *DirMgrCollection) PeekLastDirMgr() (DirMgr, error) {

  ePrefix := "DirMgrCollection.PeekLastDirMgr()"

  if dMgrs.dirMgrs == nil {
    dMgrs.dirMgrs = make([]DirMgr, 0, 100)
  }

  arrayLen := len(dMgrs.dirMgrs)

  if arrayLen == 0 {
    return DirMgr{},
      errors.New(ePrefix +
        "Error: The Directory Manager Collection, 'DirMgrCollection' is EMPTY!")
  }

  return dMgrs.dirMgrs[arrayLen-1].CopyOut(), nil
}
