package pathfileops

import (
  "fmt"
  "os"
  "time"
)

type dirMgrHelper struct {
  dMgr DirMgr
}

func (dMgrHlpr *dirMgrHelper) deleteDirectoryAll(
  dMgr *DirMgr,
  ePrefix string) error {

  ePrefixCurrMethod := "dirMgrHelper.deleteDirectoryAll() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod + "\n"
  }

  dirPathDoesExist,
    _,
    err :=
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      ePrefix,
      "dMgr")

  if err != nil {
    return err
  }

  if !dirPathDoesExist {
    return nil
  }

  var err2 error

  for i := 0; i < 3; i++ {

    err2 = os.RemoveAll(dMgr.absolutePath)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error returned by os.RemoveAll(dMgr.absolutePath) "+
        "returned error.\n"+
        "dMgr.absolutePath='%v'\nError='%v'\n",
        dMgr.absolutePath, err2.Error())
    } else {
      err = nil
      break
    }

    time.Sleep(50 * time.Millisecond)
  }

  if err != nil {
    return err
  }

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      ePrefix,
      "dMgr")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "ERROR: After attempted directory deletion, a non-path error was returned.\n"+
      "Error='%v'\n", err.Error())
  }

  if dirPathDoesExist {
    return fmt.Errorf(ePrefix+
      "Error: FAILED TO DELETE DIRECTORY!!\n"+
      "Directory Path still exists!\n"+
      "Directory Path='%v'\n", dMgr.absolutePath)
  }

  return nil
}
func (dMgrHlpr *dirMgrHelper) doesDirectoryExist(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (dirPathDoesExist bool, fInfo FileInfoPlus, err error) {

  ePrefixCurrMethod := "dirMgrHelper.doesDirectoryExist() "

  dirPathDoesExist = false
  fInfo = FileInfoPlus{}
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if !dMgr.isInitialized {
    err = fmt.Errorf(ePrefix +
      "Error: DirMgr is NOT Initialized.\n")
    return dirPathDoesExist, fInfo, err
  }

  var absFInfo os.FileInfo
  _,
    dMgr.doesAbsolutePathExist,
    absFInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.absolutePath,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr.absolutePath")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.doesAbsolutePathExist {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = nil
    return dirPathDoesExist, fInfo, err
  }

  if !absFInfo.Mode().IsDir() {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    err = fmt.Errorf(ePrefix+
      "Error: Directory path exists, but "+
      "it is a file - NOT A DIRECTORY!\n"+
      "DirMgr='%v'\n",
      dMgr.absolutePath)
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  _,
    dMgr.doesPathExist,
    _,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.path,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr.path")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.doesPathExist {
    err = fmt.Errorf(ePrefix+
      "Error: Directory absolute path exists, "+
      "but Directory original 'path' DOES NOT "+
      "EXIST!\n"+
      "dMgr.absolutePath='%v'\n"+
      "dMgr.path='%v'\n",
      dMgr.absolutePath, dMgr.path)
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  // both dMgr.path and dMgr.doesAbsolutePathExist
  // exist. And, there are no errors

  dMgr.doesAbsolutePathExist = true
  dMgr.doesPathExist = true
  dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromFileInfo(absFInfo)
  fInfo = dMgr.actualDirFileInfo.CopyOut()
  dirPathDoesExist = false
  err = nil
  return dirPathDoesExist, fInfo, err
}

/*
func (dMgrHlpr *dirMgrHelper) doesDirectoryExist(
  dMgr *DirMgr,
  preProcessCode PreProcessPathCode,
  ePrefix string,
  dMgrLabel string) ( dirPathDoesExist bool,
                      fInfo FileInfoPlus,
                      err error) {

  ePrefixCurrMethod := "dirMgrHelper.doesDirectoryExist() "

  dirPathDoesExist = false
  fInfo = FileInfoPlus{}
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod + "\n"
  }

  if !dMgr.isInitialized {
    err = fmt.Errorf(ePrefix +
      "Error: DirMgr is NOT Initialized.\n")
    return dirPathDoesExist, fInfo, err
  }

  dMgr.isAbsolutePathPopulated = false

  if dMgr.absolutePath == "" {
    err = fmt.Errorf(
      ePrefix +
        "Error: DirMgr.absolutePath is EMPTY!.\n")
    return dirPathDoesExist, fInfo, err
  }

  dMgr.isAbsolutePathPopulated = true

  dMgr.isPathPopulated = false

  if dMgr.path == "" {
    err = fmt.Errorf(ePrefix + "Error: DirMgr.absolutePath is EMPTY!.")
    return dirPathDoesExist, fInfo, err
  }

  dMgr.isPathPopulated = true

  var absFInfo os.FileInfo
  _,
    dMgr.doesAbsolutePathExist,
    absFInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.absolutePath,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr.absolutePath")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.doesAbsolutePathExist {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = nil
    return dirPathDoesExist, fInfo, err
  }

  if !absFInfo.Mode().IsDir() {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    err = fmt.Errorf(ePrefix+
      "Error: Directory path exists, but "+
      "it is a file - NOT A DIRECTORY!\n"+
      "DirMgr='%v'\n",
      dMgr.absolutePath)
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  _,
    dMgr.doesPathExist,
    _,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.path,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr.path")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.doesPathExist {
    err = fmt.Errorf(ePrefix+
      "Error: Directory absolute path exists, "+
      "but Directory original 'path' DOES NOT "+
      "EXIST!\n"+
      "dMgr.absolutePath='%v'\n"+
      "dMgr.path='%v'\n",
      dMgr.absolutePath, dMgr.path)
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  // both dMgr.path and dMgr.doesAbsolutePathExist
  // exist. And, there are no errors

  dMgr.doesAbsolutePathExist = true
  dMgr.doesPathExist = true
  dMgr.actualDirFileInfo = FileInfoPlus{}.NewFromFileInfo(absFInfo)
  fInfo = dMgr.actualDirFileInfo.CopyOut()
  dirPathDoesExist = false
  err = nil
  return dirPathDoesExist, fInfo, err
}
*/

func (dMgrHlpr *dirMgrHelper) lowLevelDoesDirectoryExist(
  dirPath string,
  preProcessCode PreProcessPathCode,
  ePrefix,
  filePathTitle string) (absolutePath string,
  dirPathDoesExist bool,
  fInfo FileInfoPlus,
  err error) {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelDoesDirectoryExist() "

  absolutePath = ""
  dirPathDoesExist = false
  fInfo = FileInfoPlus{}
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if len(filePathTitle) == 0 {
    filePathTitle = "DirMgr Path"
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, absolutePath =
    fh.isStringEmptyOrBlank(dirPath)

  if errCode == -1 {
    err = fmt.Errorf(ePrefix+
      "Error: Input parameter '%v' is an empty string!\n", filePathTitle)
    return absolutePath, dirPathDoesExist, fInfo, err
  }

  if errCode == -2 {
    err = fmt.Errorf(ePrefix+
      "Error: Input parameter '%v' consists of blank spaces!\n",
      filePathTitle)
    return absolutePath, dirPathDoesExist, fInfo, err
  }

  var err2 error

  if preProcessCode == PreProcPathCode.PathSeparator() {

    absolutePath = fh.AdjustPathSlash(dirPath)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    absolutePath, err2 = fh.MakeAbsolutePath(dirPath)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+"fh.MakeAbsolutePath() FAILED!\n"+
        "%v", err2.Error())
      return absolutePath, dirPathDoesExist, fInfo, err
    }

  } else {
    absolutePath = dirPath
  }

  var info os.FileInfo

  for i := 0; i < 3; i++ {

    dirPathDoesExist = false
    fInfo = FileInfoPlus{}
    err = nil

    info, err2 = os.Stat(absolutePath)

    if err2 != nil {

      if os.IsNotExist(err2) {

        dirPathDoesExist = false
        fInfo = FileInfoPlus{}
        err = nil
        return absolutePath, dirPathDoesExist, fInfo, err
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be test
      // up to 3-times before it is returned.
      err = fmt.Errorf(ePrefix+"Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n",
        filePathTitle, filePathTitle, absolutePath, err2.Error())
      fInfo = FileInfoPlus{}
      dirPathDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      dirPathDoesExist = true
      err = nil
      fInfo = FileInfoPlus{}.NewFromFileInfo(info)
      return absolutePath, dirPathDoesExist, fInfo, err
    }

    time.Sleep(30 * time.Millisecond)
  }

  return absolutePath, dirPathDoesExist, fInfo, err
}
