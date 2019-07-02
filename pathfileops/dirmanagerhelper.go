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
  ePrefix string,
  dMgrLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.deleteDirectoryAll() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dirPathDoesExist,
    _,
    err :=
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {
    return err
  }

  if !dirPathDoesExist {
    return nil
  }

  err = dMgrHlpr.lowLevelDeleteDirectoryAll(
    dMgr,
    ePrefix,
    dMgrLabel)

  if err != nil {
    return err
  }

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nERROR: After attempted directory deletion, a non-path error was returned.\n"+
      "Error='%v'\n", err.Error())
  }

  if dirPathDoesExist {
    return fmt.Errorf(ePrefix+
      "\nError: FAILED TO DELETE DIRECTORY!!\n"+
      "Directory Path still exists!\n"+
      "Directory Path='%v'\n", dMgr.absolutePath)
  }

  return nil
}

// doesDirectoryExist - Helper method used by DirMgr to test for
// existence of directory path. In addition, this method performs
// validation on the 'DirMgr' instance.
//
func (dMgrHlpr *dirMgrHelper) doesDirectoryExist(
  dMgr *DirMgr,
  preProcessCode PreProcessPathCode,
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
      "\nError: DirMgr is NOT Initialized.\n")
    return dirPathDoesExist, fInfo, err
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, dMgr.absolutePath =
    fh.isStringEmptyOrBlank(dMgr.absolutePath)

  if errCode == -1 {
    dMgr.absolutePath = ""
    dMgr.path = ""
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v'.absolutePath is an empty string!\n",
      dMgrLabel)
    return dirPathDoesExist, fInfo, err
  }

  if errCode == -2 {
    dMgr.absolutePath = ""
    dMgr.path = ""
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' consists of blank spaces!\n",
      dMgrLabel)

    return dirPathDoesExist, fInfo, err
  }

  var err2 error

  if preProcessCode == PreProcPathCode.PathSeparator() {

    dMgr.absolutePath = fh.AdjustPathSlash(dMgr.absolutePath)

  } else if preProcessCode == PreProcPathCode.AbsolutePath() {

    dMgr.absolutePath, err2 = fh.MakeAbsolutePath(dMgr.absolutePath)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "\nError: fh.MakeAbsolutePath(%v.absolutePath) FAILED!\n"+
        "%v.absolutePath='%v'\nError='%v'", dMgrLabel, dMgrLabel, err2.Error())
      return dirPathDoesExist, fInfo, err
    }

  }

  var absFInfo, pathFInfo os.FileInfo

  dMgr.doesAbsolutePathExist,
    absFInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.absolutePath,
    ePrefix,
    dMgrLabel+".absolutePath")

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
      "\nError: Directory absolute path exists, but "+
      "it is a file - NOT A DIRECTORY!\n"+
      "%v='%v'\n",
      dMgrLabel,
      dMgr.absolutePath)
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if absFInfo.Mode().IsRegular() {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    err = fmt.Errorf(ePrefix+
      "\nError: Directory absolute path exists, but "+
      "it is classified as as a Regular File!\n"+
      "%v='%v'\n",
      dMgrLabel,
      dMgr.absolutePath)
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  dMgr.doesPathExist,
    pathFInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dMgr.path,
    ePrefix,
    dMgrLabel+".path")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.doesPathExist {
    err = fmt.Errorf(ePrefix+
      "\nError: Directory absolute path exists, "+
      "but original directory 'path' DOES NOT "+
      "EXIST!\n"+
      "%v.absolutePath='%v'\n"+
      "%v.path='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      dMgrLabel,
      dMgr.path)

    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false

    return dirPathDoesExist, fInfo, err
  }

  if !pathFInfo.Mode().IsDir() {
    err = fmt.Errorf(ePrefix+
      "\nError: Directory path absolute path exists, "+
      "but original directory 'path' is NOT A DIRECTORY!!\n"+
      "%v.absolutePath='%v'\n"+
      "%v.path='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      dMgrLabel,
      dMgr.path)

    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    return dirPathDoesExist, fInfo, err
  }

  if pathFInfo.Mode().IsRegular() {
    err = fmt.Errorf(ePrefix+
      "\nError: Directory path exists, "+
      "but original directory 'path' is classified "+
      "as a Regular File!!\n"+
      "%v.absolutePath='%v'\n"+
      "%v.path='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      dMgrLabel,
      dMgr.path)

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
  dirPathDoesExist = true
  err = nil
  return dirPathDoesExist, fInfo, err
}

func (dMgrHlpr *dirMgrHelper) lowLevelDeleteDirectoryAll(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelDeleteDirectoryAll() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2 error

  for i := 0; i < 3; i++ {

    err2 = os.RemoveAll(dMgr.absolutePath)

    if err2 != nil {
      err = fmt.Errorf(ePrefix+
        "Error returned by os.RemoveAll(%v.absolutePath) "+
        "returned error.\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err2.Error())
    } else {
      // err2 == nil
      // Deletion was successful
      return nil
    }

    time.Sleep(50 * time.Millisecond)
  }

  return err
}

func (dMgrHlpr *dirMgrHelper) lowLevelDoesDirectoryExist(
  dirPath string,
  ePrefix,
  dirPathLabel string) (dirPathDoesExist bool,
  fInfo FileInfoPlus,
  err error) {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelDoesDirectoryExist() "

  dirPathDoesExist = false
  fInfo = FileInfoPlus{}
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if len(dirPathLabel) == 0 {
    dirPathLabel = "DirMgr"
  }

  var err2 error
  var info os.FileInfo

  for i := 0; i < 3; i++ {

    dirPathDoesExist = false
    fInfo = FileInfoPlus{}
    err = nil

    info, err2 = os.Stat(dirPath)

    if err2 != nil {

      if os.IsNotExist(err2) {

        dirPathDoesExist = false
        fInfo = FileInfoPlus{}
        err = nil
        return dirPathDoesExist, fInfo, err
      }
      // err == nil and err != os.IsNotExist(err)
      // This is a non-path error. The non-path error will be test
      // up to 3-times before it is returned.
      err = fmt.Errorf(ePrefix+"Non-Path error returned by os.Stat(%v)\n"+
        "%v='%v'\nError='%v'\n",
        dirPathLabel, dirPathLabel, err2.Error())
      fInfo = FileInfoPlus{}
      dirPathDoesExist = false

    } else {
      // err == nil
      // The path really does exist!
      dirPathDoesExist = true
      err = nil
      fInfo = FileInfoPlus{}.NewFromFileInfo(info)
      return dirPathDoesExist, fInfo, err
    }

    time.Sleep(30 * time.Millisecond)
  }

  return dirPathDoesExist, fInfo, err
}
