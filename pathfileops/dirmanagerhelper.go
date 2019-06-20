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

  _,
    dirPathDoesExist,
    fInfoPlus,
    err :=
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if err != nil {
    return err
  }

  if !dirPathDoesExist {
    return nil
  }

  if dirPathDoesExist && !fInfoPlus.IsDir() {
    return fmt.Errorf(ePrefix+
      "ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
      "DMgr='%v'\n", dMgr.absolutePath)
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

  _,
    dirPathDoesExist,
    fInfoPlus,
    err =
    FileHelper{}.doesPathFileExist(
      dMgr.absolutePath,
      PreProcPathCode.None(),
      ePrefix,
      "dMgr.absolutePath")

  if err != nil {
    dMgr.doesAbsolutePathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dMgr.doesPathExist = false
    return fmt.Errorf(ePrefix+
      "ERROR: After attempted directory deletion, a non-path error was returned.\n"+
      "Directory Path='%v'\nError='%v'\n",
      dMgr.absolutePath, err.Error())
  }

  if dirPathDoesExist {
    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = fInfoPlus.CopyOut()
    _ = dMgr.DoesPathExist()
    return fmt.Errorf(ePrefix+
      "Error: FAILED TO DELETE DIRECTORY!!\n"+
      "Directory Path='%v'\n", dMgr.absolutePath)
  }

  return nil
}
