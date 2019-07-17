package appExamples

import (
  appLib "../appLibs"
  pathFileOps "../pathfileops"
  "fmt"
  "time"
)

type DirMgrExamples struct {
  Input  string
  Output string
}

func (dMgrEx DirMgrExamples) WalkDirFindFiles(
  dMgr pathFileOps.DirMgr,
  filePattern string,
  filesOlderThan, filesNewerThan time.Time) error {

  ePrefix := "WalkDirFindFiles() "

  du := appLib.DateTimeUtility{}

  fsc := pathFileOps.FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{filePattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = pathFileOps.FileSelectMode.ANDSelect()

  dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    return fmt.Errorf("Error returned from dMgr.FindWalkDirFiles(filePattern, "+
      "filesOlderThan, filesNewerThan) filePattern='%v'  filesOlderThan='%v' "+
      "filesNewerThan='%v' Error='%v' \n",
      filePattern, du.GetDateTimeYMDAbbrvDowNano(filesOlderThan),
      du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
  }

  dMgrEx.PrintDirMgrFields(dMgr)
  fmt.Println("----------------------------")
  fmt.Println("     startPath: ", dWalkInfo.StartPath)
  fmt.Println("   filePattern: ", filePattern)
  fmt.Println("filesOlderThan: ", du.GetDateTimeYMDAbbrvDowNano(filesOlderThan))
  fmt.Println("filesNewerThan: ", du.GetDateTimeYMDAbbrvDowNano(filesNewerThan))

  if dWalkInfo.FoundFiles.GetNumOfFileMgrs() == 0 {
    fmt.Println("No Files Found")
  } else {
    fmt.Println("Files Found: ")
    for i := 0; i < dWalkInfo.FoundFiles.GetNumOfFileMgrs(); i++ {

      fileMgr, err := dWalkInfo.FoundFiles.PeekFileMgrAtIndex(i)

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by dWalkInfo.FoundFiles.PeekFileMgrAtIndex(i). "+
          "i='%v' Error='%v' ", i, err.Error())
      }

      fInfoPlus, err := fileMgr.GetFileInfoPlus()

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned from fileMgr.GetFileInfoPlus() \n"+
          "i='%v' FileName='%v' Error='%v' \n", i,
          fileMgr.GetAbsolutePathFileName(),
          err.Error())
      }

      fmt.Printf("  Name: %v Mod Date: %v path: %v \n",
        fileMgr.GetFileNameExt(),
        fInfoPlus.ModTime(), fInfoPlus.DirPath())
    }
    fmt.Println()
  }

  if dWalkInfo.Directories.GetNumOfDirs() > 0 {
    fmt.Println("dirMgrs Found:")

    for k := 0; k < dWalkInfo.Directories.GetNumOfDirs(); k++ {

      dirMgr, err := dWalkInfo.Directories.PeekDirMgrAtIndex(k)

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by dWalkInfo.Directories.PeekDirMgrAtIndex(k). "+
          "k='%v' Error='%v' ", k, err.Error())
      }

      fmt.Printf("Dir: %v \n", dirMgr.GetPath())
    }

  }

  if len(dWalkInfo.ErrReturns) > 0 {
    fmt.Println("Errors Found: ")

    for j := 0; j < len(dWalkInfo.ErrReturns); j++ {
      fmt.Printf("Error Returns: %v \n", dWalkInfo.ErrReturns[j])
    }

    fmt.Println("")
  }

  return nil
}

func (dMgrEx DirMgrExamples) WalkDirFindFiles2(
  dMgr pathFileOps.DirMgr,
  filePattern string,
  filesOlderThan,
  filesNewerThan time.Time) error {

  ePrefix := "WalkDirFindFiles2() "

  du := appLib.DateTimeUtility{}
  fsc := pathFileOps.FileSelectionCriteria{}

  fsc.FileNamePatterns = []string{filePattern}
  fsc.FilesOlderThan = filesOlderThan
  fsc.FilesNewerThan = filesNewerThan
  fsc.SelectCriterionMode = pathFileOps.FileSelectMode.ANDSelect()

  dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from dMgr.FindWalkDirFiles(filePattern, filesOlderThan, "+
      "filesNewerThan) filePattern='%v'  filesOlderThan='%v' filesNewerThan='%v' "+
      "Error='%v' \n",
      filePattern,
      du.GetDateTimeYMDAbbrvDowNano(filesOlderThan),
      du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
  }

  dMgrEx.PrintDirMgrFields(dMgr)
  fmt.Println("----------------------------")
  fmt.Println("     startPath: ", dWalkInfo.StartPath)
  fmt.Println("   filePattern: ", filePattern)
  fmt.Println("filesOlderThan: ", du.GetDateTimeYMDAbbrvDowNano(filesOlderThan))
  fmt.Println("filesNewerThan: ", du.GetDateTimeYMDAbbrvDowNano(filesNewerThan))

  if dWalkInfo.FoundFiles.GetNumOfFileMgrs() == 0 {
    fmt.Println("No Files Found")
  } else {

    fmt.Println("Files Found: ")

    for i := 0; i < dWalkInfo.FoundFiles.GetNumOfFileMgrs(); i++ {

      fileMgr, err := dWalkInfo.FoundFiles.PeekFileMgrAtIndex(i)

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by dWalkInfo.FoundFiles.PeekFileMgrAtIndex(i). "+
          "i='%v' Error='%v' ", i, err.Error())
      }

      fileInfoPlus, err := fileMgr.GetFileInfoPlus()

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by fileMgr.GetFileInfoPlus(). "+
          "i='%v' FileName='%v' Error='%v' ",
          i, fileMgr.GetAbsolutePathFileName(), err.Error())
      }

      fmt.Printf("  Name: %v SysInfo: %v \n",
        fileMgr.GetFileNameExt(),
        fileInfoPlus.Sys())
    }

    fmt.Println()
  }

  if dWalkInfo.Directories.GetNumOfDirs() > 0 {
    fmt.Println("dirMgrs Found:")

    for k := 0; k < dWalkInfo.Directories.GetNumOfDirs(); k++ {

      dirMgr, err := dWalkInfo.Directories.PeekDirMgrAtIndex(k)

      if err != nil {
        return fmt.Errorf(ePrefix+
          "Error returned by dWalkInfo.Directories.PeekDirMgrAtIndex(k). "+
          "k='%v' Error='%v' ", k, err.Error())

      }

      fmt.Printf("Dir: %v \n", dirMgr.GetPath())
    }

  }

  if len(dWalkInfo.ErrReturns) > 0 {
    fmt.Println("Errors Found: ")

    for j := 0; j < len(dWalkInfo.ErrReturns); j++ {
      fmt.Printf("Error: %v \n", dWalkInfo.ErrReturns[j])
    }
    fmt.Println("")
  }

  return nil
}

func (dMgrEx DirMgrExamples) PrintDirMgrFields(dMgr pathFileOps.DirMgr) {

  ePrefix := "PrintDirMgrFields() "

  du := appLib.DateTimeUtility{}
  fmt.Println("-----------------------------------------")
  fmt.Println(" 	DirMgr Fields")
  fmt.Println("-----------------------------------------")

  fmt.Println("                isInitialized: ", dMgr.IsInitialized())
  fmt.Println("                Original path: ", dMgr.GetOriginalPath())
  fmt.Println("                         path: ", dMgr.GetPath())
  fmt.Println("             PathIsPopuslated: ", dMgr.IsPathPopulated())
  fmt.Println("                doesPathExist: ", dMgr.DoesPathExist())
  fmt.Println("                   parentPath: ", dMgr.GetParentPath())
  fmt.Println("        isParentPathPopulated: ", dMgr.IsParentPathPopulated())
  fmt.Println("                 absolutePath: ", dMgr.GetAbsolutePath())
  fmt.Println("        doesAbsolutePathExist: ", dMgr.DoesAbsolutePathExist())
  fmt.Println("               Directory Name: ", dMgr.GetDirectoryName())
  fmt.Println("                   volumeName: ", dMgr.GetVolumeName())
  fmt.Println("            isVolumePopulated: ", dMgr.IsVolumeNamePopulated())
  fmt.Println("============== File Info Data ============")

  actualDirFileInfo, err := dMgr.GetFileInfoPlus()

  if err != nil {
    fmt.Printf(ePrefix+
      "Error returned by dMgr.GetFileInfoPlus() Error='%v' \n", err.Error())
    return
  }

  if actualDirFileInfo.IsFileInfoInitialized() {
    fmt.Println("            File Info IsDir(): ", actualDirFileInfo.IsDir())
    fmt.Println("             File Info Name(): ", actualDirFileInfo.Name())
    fmt.Println("             File Info Size(): ", actualDirFileInfo.Size())
    fmt.Println("          File Info ModTime(): ", du.GetDateTimeYMDAbbrvDowNano(
      actualDirFileInfo.ModTime()))
    fmt.Println("             File Info Mode(): ", actualDirFileInfo.Mode())
    fmt.Println("          File Info     Sys(): ", actualDirFileInfo.Sys())
    if actualDirFileInfo.IsDirectoryPathInitialized() {
      fmt.Println("                   Dir path: ", actualDirFileInfo.DirPath())
    }
  } else {
    fmt.Println("File Info Data is NOT Initialized")
  }

  return
}
