package pathfileops

import (
  "fmt"
  "io"
  "os"
  pf "path/filepath"
  "time"
)

type dirMgrHelper struct {
  dMgr DirMgr
}

// copyDirectory - Helper method used by DirMgr. This method copies
// files from the directory identified by by DirMgr to a target
// directory. The files to be copied are selected according to
// file selection criteria specified by input parameter,
// 'fileSelectCriteria'.
//
func (dMgrHlpr *dirMgrHelper) copyDirectory(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (errs []error) {

  errs = make([]error, 0, 300)
  ePrefixCurrMethod := "dirMgrHelper.copyDirectory() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2, err3 error
  var dirPathDoesExist, targetPathDoesExist bool

  dirPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {

    errs = append(errs, err)
    return errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nThe current DirMgr path DOES NOT EXIST!\n"+
      "%v.absolutePath='%v'\n",
      dMgrLabel, dMgr.absolutePath)

    errs = append(errs, err)
    return errs
  }

  targetPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      targetDMgr,
      PreProcPathCode.None(),
      ePrefix,
      targetDMgrLabel)

  if err != nil {

    errs = append(errs, err)
    return errs
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)

    return errs
  }

  osPathSeparatorStr := string(os.PathSeparator)

  var src, target string
  var isMatch bool
  var nameFileInfos []os.FileInfo
  err3 = nil

  fh := FileHelper{}

  for err3 != io.EOF {

    nameFileInfos, err3 = dirPtr.Readdir(1000)

    if err3 != nil && err3 != io.EOF {
      _ = dirPtr.Close()
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Readdirnames(1000).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
      return errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        // We don't care about sub-directories
        continue

      }

      // This is not a directory. It is a file.
      // Determine if it matches the find file criteria.
      isMatch, err =
        fh.FilterFileName(nameFInfo, fileSelectCriteria)

      if err != nil {

        err2 =
          fmt.Errorf("\n"+ePrefix+
            "\nError returned by fh.FilterFileName(nameFInfo, fileSelectCriteria).\n"+
            "%v directorySearched='%v'\nfileName='%v'\nError='%v'\n",
            dMgrLabel, dMgr.absolutePath, nameFInfo.Name(), err.Error())

        errs = append(errs, err2)

        continue
      }

      if !isMatch {

        continue

      } else {

        // We have a match

        // Create Directory if needed
        if !targetPathDoesExist {

          err = dMgrHlpr.lowLevelMakeDir(
            targetDMgr,
            ePrefix,
            "targetDMgr")

          if err != nil {
            err2 = fmt.Errorf("\n"+ePrefix+
              "\nError creating target directory!\n"+
              "%v Directory='%v'\nError='%v'\n",
              targetDMgrLabel,
              targetDMgr.absolutePath, err.Error())

            errs = append(errs, err2)
            err3 = io.EOF
            break
          }

          targetPathDoesExist = true
        }

        src = dMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        target = targetDMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        err = fh.CopyFileByIo(src, target)

        if err != nil {
          err2 = fmt.Errorf("\n"+ePrefix+
            "\nERROR: fh.CopyFileByIo(src, target)\n"+
            "src='%v'\ntarget='%v'\nError='%v'\n\n",
            src, target, err.Error())

          errs = append(errs, err2)
        }
      }
    }
  }

  if dirPtr != nil {

    err = dirPtr.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by %v dirPtr.Close().\n"+
        "%v='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
    }
  }

  return errs
}

// copyDirectoryTree - Helper method for 'DirMgr'. This method is
// designed to copy entire directory trees.
func (dMgrHlpr *dirMgrHelper) copyDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  copyEmptyDirectories bool,
  skipTopLevelDirectory bool,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.copyDirectoryTree() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dMgrPathDoesExist,
    _,
    err :=
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: %v directory path DOES NOT EXIST!\n",
      dMgrLabel)

    errs = append(errs, err)
    return errs
  }

  var err2, err3 error

  _,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      targetDMgr,
      PreProcPathCode.None(),
      ePrefix,
      targetDMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  fh := FileHelper{}

  baseDirLen := len(dMgr.absolutePath)
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dirPtr *os.File
  var nextTargetDMgr DirMgr
  var isMatch, isTopLevelDir bool
  var srcFile, targetFile string
  var nextDir DirMgr

  nextDir, err = DirMgr{}.New(dMgr.absolutePath)

  if err != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by DirMgr{}.New(%v.absolutePath).\n"+
      "%v.absolutePath='%v'",
      dMgrLabel, dMgrLabel, dMgr.absolutePath)

    errs = append(errs, err)
    return errs
  }

  for nextDir.isInitialized {

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {
        err2 = fmt.Errorf(ePrefix+
          "\nError return from #1 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    nextTargetDMgr, err = DirMgr{}.New(
      targetDMgr.GetAbsolutePath() +
        nextDir.absolutePath[baseDirLen:])

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError return by DirMgr{}.New(%v.GetAbsolutePath() + "+
        "nextDir.absolutePath[baseDirLen:])\n"+
        "%v.GetAbsolutePath()='%v'\n"+
        "nextDir.absolutePath[baseDirLen:]='%v'\n"+
        "Error='%v'\n\n",
        targetDMgrLabel, targetDMgrLabel,
        targetDMgr.absolutePath,
        nextDir.absolutePath[baseDirLen:],
        err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {

        err2 = fmt.Errorf(ePrefix+
          "\nError return from #2 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    if baseDirLen == len(nextDir.GetAbsolutePath()) {
      isTopLevelDir = true
    } else {
      isTopLevelDir = false
    }

    if isTopLevelDir &&
      !skipTopLevelDirectory &&
      copyEmptyDirectories {

      err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "1-nextTargetDMgr")

    } else if !isTopLevelDir && copyEmptyDirectories {

      err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "2-nextTargetDMgr")

    } else {
      err = nil
    }

    if err != nil {

      err2 = fmt.Errorf("\n"+ePrefix+
        "\nError creating target directory!\n"+
        "Target Next Directory='%v'\nError='%v'\n",
        nextTargetDMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {

        err2 = fmt.Errorf(ePrefix+
          "\nError return from #3 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n", err.Error())
        errs = append(errs, err2)
        return

      } else if err != nil {

        break
      }

      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dirPtr.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "\nError returned by dirPtr.Readdirnames(-1).\n"+
          "dMgr.absolutePath='%v'\nError='%v'\n",
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          err = dirs.AddDirMgrByPathNameStr(
            nextDir.absolutePath +
              osPathSepStr +
              nameFInfo.Name())

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "\nError returned by dirs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

        } else if isTopLevelDir && skipTopLevelDirectory {
          // Do NOT process files for top level directory
          // when parameter skipTopLevelDirectory = 'true'

          continue

        } else {
          // This is a file which is eligible for processing

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, fileSelectCriteria)

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "\nError returned by fh.FilterFileName(nameFInfo, fileSelectCriteria).\n"+
                "%v directorySearched='%v'\nfileName='%v'\nError='%v'\n\n",
                dMgrLabel,
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match

            dMgrPathDoesExist,
              _,
              err =
              dMgrHlpr.doesDirectoryExist(
                &nextTargetDMgr,
                PreProcPathCode.None(),
                ePrefix,
                "nextTargetDMgr")

            if err != nil {
              errs = append(errs, err)
              err3 = io.EOF
              break
            }

            // Create Directory if needed
            if !dMgrPathDoesExist {

              err = dMgrHlpr.lowLevelMakeDir(
                &nextTargetDMgr,
                ePrefix,
                "3-nextTargetDMgr")

              if err != nil {
                err2 = fmt.Errorf("\n"+ePrefix+
                  "\nError creating targetFile directory!\n"+
                  "Target Directory='%v'\nError='%v'\n\n",
                  nextTargetDMgr.absolutePath, err.Error())

                errs = append(errs, err2)
                err3 = io.EOF
                break
              }
            }

            srcFile = nextDir.absolutePath +
              osPathSepStr + nameFInfo.Name()

            targetFile = nextTargetDMgr.absolutePath +
              osPathSepStr + nameFInfo.Name()

            err = fh.CopyFileByIo(srcFile, targetFile)

            if err != nil {
              err2 = fmt.Errorf("\n"+ePrefix+
                "\nERROR: fh.CopyFileByIo(srcFile, targetFile)\n"+
                "srcFile='%v'\ntargetFile='%v'\nError='%v'\n\n",
                srcFile, targetFile, err.Error())

              errs = append(errs, err2)
            }
          }
        }
      }
    }

    if dirPtr != nil {
      err = dirPtr.Close()

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close().\n"+
          "dirPtr='%v'\nError='%v'\n",
          dMgr.absolutePath, err.Error())

        errs = append(errs, err2)
      }
    }

    nextDir, err = dirs.PopFirstDirMgr()
    if err != nil && err != io.EOF {

      err2 = fmt.Errorf(ePrefix+
        "\nError return from #4 dirs.PopFirstDirMgr()\n"+
        "Error='%v'\n", err.Error())

      errs = append(errs, err2)
      return

    } else if err != nil {

      break
    }

  }

  return errs
}

// deleteAllFilesInDirectory - Helper method used by DirMgr. This
// method deletes all the files in the current directory. ONLY files
// are deleted NOT sub-directories.
//
func (dMgrHlpr *dirMgrHelper) deleteAllFilesInDirectory(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.deleteAllFilesInDirectory() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  errs = make([]error, 0, 300)
  var err2, err3 error
  osPathSepStr := string(os.PathSeparator)

  dirPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    "dMgr")

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dirPathDoesExist {
    err =
      fmt.Errorf(ePrefix+
        "\nERROR: %v Path DOES NOT EXIST!\n"+
        "%v Path='%v'\n",
        dMgrLabel,
        dMgrLabel,
        dMgr.absolutePath)

    errs = append(errs, err)
    return errs
  }

  dMgrPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  err3 = nil
  var nameFileInfos []os.FileInfo

  for err3 != io.EOF {

    nameFileInfos, err3 = dMgrPtr.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      if dMgrPtr != nil {
        _ = dMgrPtr.Close()
      }

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dMgrPtr.Readdirnames(-1).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath,
        err3.Error())

      errs = append(errs, err2)
      return errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        continue

      } else {

        if !nameFInfo.Mode().IsRegular() {
          err2 = fmt.Errorf(ePrefix+
            "\nError: fileName is NOT classified as a 'Regular' File!\n"+
            "fileName='%v'",
            dMgr.absolutePath+osPathSepStr+nameFInfo.Name())
          errs = append(errs, err2)
        }

        // This is a file
        err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFInfo.Name())

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "\nError returned by os.Remove(fileName).\n"+
            "An attempt to delete 'fileName' as Failed!\n"+
            "%v.absolutePath='%v'\nfileName='%v'\nError='%v'\n\n",
            dMgrLabel,
            dMgr.absolutePath,
            dMgr.absolutePath+osPathSepStr+nameFInfo.Name(),
            err.Error())

          errs = append(errs, err2)
        }
      }
    }
  }

  if dMgrPtr != nil {

    err = dMgrPtr.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dMgrPtr.Close().\n"+
        "An attempt to close the os.File pointer to the current\n"+
        "%v path has FAILED!\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())
      errs = append(errs, err2)
    }
  }

  return errs
}

// deleteAllSubDirectories - The directory identified by the input
// parameter 'dMgr' instance is treated as the parent directory.
// This method will then proceed to delete all directories and files
// which are are subsidiary to this parent directory. Essentially,
// all sub-directories which are subordinate to the the 'dMgr'
// directory will be deleted along with their constituent files.
//
func (dMgrHlpr *dirMgrHelper) deleteAllSubDirectories(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.doesDirectoryExist() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dirPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)
    errs = append(errs, err)
    return errs
  }

  var err2, err3 error

  dirMgrPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error return by os.Open(dMgr.absolutePath). "+
      "dMgr.absolutePath='%v' Error='%v' ",
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)

    return errs
  }

  var nameFileInfos []os.FileInfo
  err3 = nil
  osPathSeparatorStr := string(os.PathSeparator)

  for err3 != io.EOF {

    nameFileInfos, err3 = dirMgrPtr.Readdir(1000)

    if err3 != nil && err3 != io.EOF {
      _ = dirMgrPtr.Close()
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirMgrPtr.Readdirnames(1000).\n"+
        "dMgr.absolutePath='%v'\nError='%v'\n\n",
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        err = os.RemoveAll(dMgr.absolutePath + osPathSeparatorStr + nameFInfo.Name())

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "\nError returned by os.RemoveAll(subDir)\n"+
            "subDir='%v'\nError='%v'\n\n",
            dMgr.absolutePath+osPathSeparatorStr+nameFInfo.Name(), err.Error())

          errs = append(errs, err2)

          continue
        }
      }
    }
  }

  if dirMgrPtr != nil {

    err = dirMgrPtr.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by %vPtr.Close().\n"+
        "%v='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
    }
  }

  return errs
}

// deleteDirectoryAll - This method will remove the directory identified by the
// input parameter 'dMgr'. It will also delete all child directories and files
// in the directory tree.
//
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

// deleteDirectoryTreeFiles - Helper method designed to delete files using
// file selection criteria. Scope of scans and file deletions is controlled
// by input parameter 'scanSubDirectories'. If set to 'true' files may be
// deleted in the entire directory tree. If set to 'false' the file deletions
// are limited solely to the directory identified by the current 'DirMgr'
// instance.
//
func (dMgrHlpr *dirMgrHelper) deleteDirectoryTreeFiles(
  dMgr *DirMgr,
  scanSubDirectories bool,
  deleteFileSelectionCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  deleteSelectionLabel string) (numOfSubDirectories,
  numOfRemainingFiles,
  numOfDeletedFiles int,
  errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.deleteDirectoryTreeFiles() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dirPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  var err2, err3 error

  osPathSepStr := string(os.PathSeparator)
  var xNumOfSubDirectories, xNumOfTotalFiles, xNumOfDeletedFiles int

  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dirPtr *os.File
  var isMatch bool
  var nextDir DirMgr

  nextDir, err = DirMgr{}.New(dMgr.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nError returned by DirMgr{}.New(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())
    errs = append(errs, err)

    return numOfSubDirectories,
      numOfRemainingFiles,
      numOfDeletedFiles,
      errs
  }

  fh := FileHelper{}

  for nextDir.isInitialized {

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath). "+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)

      nextDir, err = dirs.PopFirstDirMgr()

      if err != nil && err != io.EOF {
        err2 = fmt.Errorf(ePrefix+
          "\nError returned by #1 dirs.PopFirstDirMgr()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        return numOfSubDirectories,
          numOfRemainingFiles,
          numOfDeletedFiles,
          errs

      } else if err != nil {

        break
      }

      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dirPtr.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdirnames(-1).\n"+
          "%v.absolutePath='%v'\nError='%v'\n\n",
          dMgrLabel,
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          xNumOfSubDirectories++

          if !scanSubDirectories {
            continue
          }

          err = dirs.AddDirMgrByPathNameStr(nextDir.absolutePath + osPathSepStr + nameFInfo.Name())

          if err != nil {

            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by dirs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

        } else {
          // This is a file which is eligible for processing

          xNumOfTotalFiles++

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, deleteFileSelectionCriteria)

          if err != nil {

            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by fh.FilterFileName(nameFInfo, %v).\n"+
                "%v directory searched='%v'\nfileName='%v'\nError='%v'\n\n",
                deleteSelectionLabel, dMgrLabel,
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match, delete the file

            err = os.Remove(nextDir.absolutePath + osPathSepStr + nameFInfo.Name())

            if err != nil {
              err2 = fmt.Errorf(ePrefix+
                "\nERROR returned by os.Remove(pathFileName)\n"+
                "pathFileName='%v'\nError='%v'\n\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(),
                err.Error())

              errs = append(errs, err2)

            }

            xNumOfDeletedFiles++
          }
        }
      }
    }

    if dirPtr != nil {
      err = dirPtr.Close()

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "Error returned by dirPtr.Close(). "+
          "dirPtr='%v' Error='%v' ",
          dMgr.absolutePath, err.Error())

        errs = append(errs, err2)
      }

      dirPtr = nil
    }

    nextDir, err = dirs.PopFirstDirMgr()

    if err != nil && err != io.EOF {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by #2 dirs.PopFirstDirMgr()\n"+
        "Error='%v'\n", err.Error())

      errs = append(errs, err2)

      return numOfSubDirectories,
        numOfRemainingFiles,
        numOfDeletedFiles,
        errs

    } else if err != nil {

      break
    }

  }

  numOfSubDirectories = xNumOfSubDirectories
  numOfRemainingFiles = xNumOfTotalFiles - xNumOfDeletedFiles
  numOfDeletedFiles = xNumOfDeletedFiles

  return numOfSubDirectories,
    numOfRemainingFiles,
    numOfDeletedFiles,
    errs
}

// deleteFilesByNamePattern - Receives a string defining a pattern to
// use in searching file names for all files in the directory identified
// by the input parameter 'dMgr'.
//
// If a file name matches the pattern specified by input parameter,
// 'fileSearchPattern', it will be deleted.
//
func (dMgrHlpr *dirMgrHelper) deleteFilesByNamePattern(
  dMgr *DirMgr,
  fileSearchPattern string,
  ePrefix string,
  dMgrLabel string,
  fileSearchLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.deleteFilesByNamePattern() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dirPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return errs
  }

  var err2, err3 error

  fh := FileHelper{}

  errCode := 0

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode == -1 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' is an empty string!\n",
      fileSearchLabel)

    errs = append(errs, err2)
    return errs
  }

  if errCode == -2 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' consists of blank spaces!\n",
      fileSearchLabel)

    errs = append(errs, err2)
    return errs
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  err3 = nil
  var nameFileInfos []os.FileInfo
  osPathSepStr := string(os.PathSeparator)
  var isMatch bool

  for err3 != io.EOF {

    nameFileInfos, err3 = dirPtr.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      if dirPtr != nil {
        _ = dirPtr.Close()
      }

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Readdirnames(-1).\n"+
        "#v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
      return errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        continue

      } else {

        isMatch, err = pf.Match(fileSearchPattern, nameFInfo.Name())

        if err != nil {

          err2 = fmt.Errorf(ePrefix+
            "\nError returned by (path/filepath) pf.Match(%v, fileName).\n"+
            "\n%v Directory Searched='%v'\n%v='%v' fileName='%v'\nError='%v'\n\n",
            fileSearchLabel,
            dMgrLabel,
            dMgr.absolutePath,
            fileSearchLabel,
            fileSearchPattern,
            nameFInfo.Name(),
            err.Error())

          errs = append(errs, err2)
          continue
        }

        if !isMatch {

          continue

        } else {

          err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFInfo.Name())

          if err != nil {
            err2 = fmt.Errorf(ePrefix+
              "\nError returned by os.Remove(pathFileName).\n"+
              "pathFileName='%v'\nError='%v'\n\n",
              dMgr.absolutePath+osPathSepStr+nameFInfo.Name(),
              err.Error())

            errs = append(errs, err2)
            continue
          }
        }
      }
    }
  }

  if dirPtr != nil {

    err = dirPtr.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Close().\n"+
        "%v='%v'\nError='%v'\n\n",
        dMgrLabel,
        dMgr.absolutePath, err.Error())
      errs = append(errs, err2)
    }
  }

  return errs
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

func (dMgrHlpr *dirMgrHelper) empty(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) error {

  var err error

  err = nil

  dMgr.isInitialized = false
  dMgr.originalPath = ""
  dMgr.path = ""
  dMgr.isPathPopulated = false
  dMgr.doesPathExist = false
  dMgr.parentPath = ""
  dMgr.isParentPathPopulated = false
  dMgr.absolutePath = ""
  dMgr.isAbsolutePathPopulated = false
  dMgr.doesAbsolutePathExist = false
  dMgr.isAbsolutePathDifferentFromPath = false
  dMgr.directoryName = ""
  dMgr.volumeName = ""
  dMgr.isVolumePopulated = false
  dMgr.actualDirFileInfo = FileInfoPlus{}

  return err
}

// executeDirectoryFileOps - Performs a a file operation on specified 'selected' files
// in the current directory ONLY. This function does NOT perform operations on the
// sub directories (a.k.a. the directory tree).
//
func (dMgrHlpr *dirMgrHelper) executeDirectoryFileOps(
  dMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  fileOps []FileOperationCode,
  targetBaseDir *DirMgr,
  ePrefix string,
  dMgrLabel string,
  targetDirLabel string,
  fileSelectLabel string,
  fileOpsLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.executeDirectoryFileOps() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dMgrPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return errs
  }

  _,
    _,
    err2 := dMgrHlpr.doesDirectoryExist(
    targetBaseDir,
    PreProcPathCode.None(),
    ePrefix,
    targetDirLabel)

  if err2 != nil {
    errs = append(errs, err2)
    return errs
  }

  if len(fileOps) == 0 {

    err2 = fmt.Errorf(ePrefix+
      "\nError: The input parameter '%v' is a ZERO LENGTH ARRAY!\n",
      fileOpsLabel)

    errs = append(errs, err2)
    return errs
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {
    err2 := fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  nameFileInfos, err := dirPtr.Readdir(-1)

  if err != nil {
    if dirPtr != nil {
      _ = dirPtr.Close()
    }

    err2 = fmt.Errorf(ePrefix+
      "\nError returned by dirPtr.Readdirnames(-1).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return errs
  }

  fh := FileHelper{}
  var isMatch bool
  var fileOp FileOps
  srcFileNameExt := ""

  for _, nameFInfo := range nameFileInfos {

    if nameFInfo.IsDir() {
      continue
    }

    // Must be a file - process it!

    // This is not a directory. It is a file.
    // Determine if it matches the find file criteria.
    isMatch, err = fh.FilterFileName(nameFInfo, fileSelectCriteria)

    if err != nil {

      if dirPtr != nil {
        _ = dirPtr.Close()
      }

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by FileHelper{}.FilterFileName(nameFInfo, %v).\n"+
        "%v Directory Searched='%v'\nfileName='%v'\nError='%v'\n",
        fileSelectLabel,
        dMgrLabel,
        dMgr.absolutePath, nameFInfo.Name(), err.Error())

      errs = append(errs, err2)
      return errs
    }

    if !isMatch {

      continue

    }

    // Must be a match - this is a 'selected' file!
    srcFileNameExt = nameFInfo.Name()

    fileOp, err = FileOps{}.NewByDirStrsAndFileNameExtStrs(
      dMgr.GetAbsolutePath(),
      srcFileNameExt,
      targetBaseDir.GetAbsolutePath(),
      srcFileNameExt)

    if err != nil {

      if dirPtr != nil {
        _ = dirPtr.Close()
      }

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by FileOps{}.NewByDirStrsAndFileNameExtStrs()\n"+
        "%v Source Path='%v'\nsrcFileNameExt='%v'\n"+
        "%v Destination Directory='%v'\nDestination File='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.GetAbsolutePath(),
        srcFileNameExt,
        targetDirLabel,
        targetBaseDir.GetAbsolutePath(),
        srcFileNameExt,
        err.Error())

      errs = append(errs, err2)
      return errs
    }

    for i := 0; i < len(fileOps); i++ {

      err = fileOp.ExecuteFileOperation(fileOps[i])

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "\nError returned by fileOp.ExecuteFileOperation(fileOps[%v]). "+
          "FileOps='%v'\nError='%v'\n\n",
          i, fileOps[i].String(), err.Error())

        // Store the error and continue processing
        // file operations.
        errs = append(errs, err2)
      }
    }

    // finished applying file operations to this file.
    // Get another one and continue...
  }

  if dirPtr != nil {

    err = dirPtr.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Close().\n"+
        "Error='%v'\n", err.Error())

      errs = append(errs, err2)
    }
  }

  return errs
}

// executeDirectoryTreeOps - Performs File Operations on 'selected'
// files in the directory tree identified by the input parameter,
// 'dMgr'.
//
// The 'dMgr' path therefore serves as the parent directory for
// file operations performed on the directory tree. Designated
// file operations will therefore be performed on all files in
// the parent directory as well as all files in all sub-directories.
//
func (dMgrHlpr *dirMgrHelper) executeDirectoryTreeOps(
  dMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  fileOps []FileOperationCode,
  targetBaseDir *DirMgr,
  ePrefix string,
  dMgrLabel string,
  targetDirLabel string,
  fileOpsLabel string) (errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.executeDirectoryTreeOps() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dMgrPathDoesExist,
    _,
    err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    errs = append(errs, err)
    return errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return errs
  }

  _,
    _,
    err2 := dMgrHlpr.doesDirectoryExist(
    targetBaseDir,
    PreProcPathCode.None(),
    ePrefix,
    targetDirLabel)

  if err2 != nil {
    errs = append(errs, err2)
    return errs
  }

  if len(fileOps) == 0 {

    err2 = fmt.Errorf(ePrefix+
      "\nError: The input parameter '%v' is a ZERO LENGTH ARRAY!\n",
      fileOpsLabel)

    errs = append(errs, err2)
    return errs
  }

  dirOp := DirTreeOp{}.New()
  dirOp.CallingFunc = ePrefix + "\n"
  dirOp.FileOps = append(dirOp.FileOps, fileOps...)

  dirOp.TargetBaseDir, err = DirMgr{}.New(targetBaseDir.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nError returned by dirOp.TargetBaseDir = DirMgr{}.New(%v.absolutePath)\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      targetDirLabel,
      targetDirLabel,
      targetBaseDir.absolutePath,
      err.Error())

    errs = append(errs, err2)
    return errs
  }

  dirOp.SourceBaseDir, err = DirMgr{}.New(dMgr.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nError returned by dirOp.SourceBaseDir = DirMgr{}.New(%v.absolutePath)\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel,
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())

    errs = append(errs, err2)
    return errs
  }

  dirOp.FileSelectCriteria = fileSelectCriteria

  err = pf.Walk(dMgr.absolutePath, dMgrHlpr.executeFileOpsOnFoundFiles(&dirOp))

  if err != nil {
    err2 = fmt.Errorf("\n"+ePrefix+
      "\nError returned by (path/filepath) pf.Walk("+
      "%v.absolutePath, dMgrHlpr.executeFileOpsOnFoundFiles(&dirOp)).\n"+
      "%v.absolutePath='%v'\nError='%v'\n\n",
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())

    errs = append(errs, dirOp.ErrReturns...)
    errs = append(errs, err2)
    return errs
  }

  return dirOp.ErrReturns
}

// executeFileOpsOnFoundFiles - This function is designed to work in conjunction
// with a walk directory function like FindWalkDirFiles. It will process
// files extracted from a 'Directory Walk' operation initiated by the
// 'filepath.Walk' method.
//
// Thereafter, file operations will be performed on files in the directory
// tree as specified by the 'dirOp' parameter.
//
func (dMgrHlpr *dirMgrHelper) executeFileOpsOnFoundFiles(dirOp *DirTreeOp) func(string, os.FileInfo, error) error {
  return func(pathFile string, info os.FileInfo, erIn error) error {

    ePrefix := "\ndirMgrHelper.executeFileOpsOnFoundFiles() "
    var err2 error

    if erIn != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned from directory walk function.\n"+
        "pathFile='%v'\nError='%v'\n",
        pathFile, erIn.Error())
      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    if info.IsDir() {
      return nil
    }

    fh := FileHelper{}

    // This is not a directory. It is a file.
    // Determine if it matches the find file criteria.
    isFoundFile, err := fh.FilterFileName(info, dirOp.FileSelectCriteria)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria)\n"+
        "\npathFile='%v'\ninfo.Name()='%v'\nError='%v'\n",
        pathFile, info.Name(), err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    if !isFoundFile {
      return nil
    }

    srcFileNameExt := info.Name()

    destDir, err := fh.SwapBasePath(
      dirOp.SourceBaseDir.absolutePath,
      dirOp.TargetBaseDir.absolutePath,
      pathFile)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by fh.SwapBasePath(dirOp.SourceBaseDir, "+
        "dirOp.TargetBaseDir, pathFile).\n"+
        "dirOp.SourceBaseDir='%v'\n"+
        "dirOp.TargetBaseDir='%v'\n"+
        "pathFile='%v'\n"+
        "Error='%v'\n",
        dirOp.SourceBaseDir.absolutePath,
        dirOp.TargetBaseDir.absolutePath,
        pathFile,
        err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    fileOp, err := FileOps{}.NewByDirStrsAndFileNameExtStrs(
      pathFile, srcFileNameExt, destDir, srcFileNameExt)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by FileOps{}.NewByDirStrsAndFileNameExtStrs(pathFile, "+
        "srcFileNameExt, destDir, srcFileNameExt)\n"+
        "pathFile='%v'\n"+
        "srcFileNameExt='%v'\n"+
        "destDir='%v'\n"+
        "Error='%v'\n",
        pathFile,
        srcFileNameExt,
        destDir,
        err.Error())

      dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
      return nil
    }

    for i := 0; i < len(dirOp.FileOps); i++ {

      err = fileOp.ExecuteFileOperation(dirOp.FileOps[i])

      if err != nil {
        err2 = fmt.Errorf(ePrefix+
          "\nError returned by fileOp.ExecuteFileOperation(dirOp.FileOps[i]).\n"+
          "i='%v'\nFileOps='%v'\nError='%v'\n",
          i, dirOp.FileOps[i].String(), err.Error())

        dirOp.ErrReturns = append(dirOp.ErrReturns, err2)

      }
    }

    return nil
  }
}

// lowLevelDeleteDirectoryAll - Helper method designed for use by DirMgr.
// This method will delete the designated directory ('dMgr') and all
// subsidiary directories and files.
//
// This method will not perform validation services.
//
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

// lowLevelMakeDir - Helper Method used by 'DirMgr'. This method will create
// the directory path including parent directories for the path specified by
// 'dMgr'.
func (dMgrHlpr *dirMgrHelper) lowLevelMakeDir(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelMakeDir() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dMgrPathDoesExist,
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

  if dMgrPathDoesExist {
    // The directory exists
    // Nothing to do.
    return nil
  }

  fPermCfg, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
  }

  modePerm, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by fPermCfg.GetCompositePermissionMode().\n"+
      "Error='%v\n", err.Error())
  }

  err = os.MkdirAll(dMgr.absolutePath, modePerm)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by os.MkdirAll(%v.absolutePath, modePerm).\n"+
      "%v.absolutePath='%v'\nmodePerm=\"drwxrwxrwx\"\n"+
      "Error='%v'\n", err.Error())
  }

  dMgrPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {
    return fmt.Errorf("Error: After attempted directory creation, "+
      "a non-path error was generated!\n"+
      "%v.absolutePath='%v'\n"+
      "Error='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())
  }

  if !dMgrPathDoesExist {
    return fmt.Errorf("Error: After attempted directory creation,\n"+
      "the directory DOES NOT EXIST!\n"+
      "%v=%v\n", dMgrLabel, dMgr.absolutePath)
  }

  return nil
}

// lowLevelMakeDirWithPermission - Helper Method used by 'DirMgr'. This method
// will create the directory path including parent directories for the path
// specified by 'dMgr'. The permission used to create the directory path is
// specified by input parameter
//
func (dMgrHlpr *dirMgrHelper) lowLevelMakeDirWithPermission(
  dMgr *DirMgr,
  fPermCfg FilePermissionConfig,
  ePrefix string,
  dMgrLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelMakeDir() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  dMgrPathDoesExist,
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

  if dMgrPathDoesExist {
    // The directory exists
    // Nothing to do.
    return nil
  }

  err = fPermCfg.IsValid()

  if err != nil {
    return fmt.Errorf("Input Parameter 'fPermCfg' is INVALID!\n"+
      "Error returned by fPermCfg.IsValid().\n"+
      "Error='%v'\n", err.Error())
  }

  modePerm, err := fPermCfg.GetCompositePermissionMode()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by fPermCfg.GetCompositePermissionMode().\n"+
      "Error='%v\n", err.Error())
  }

  err = os.MkdirAll(dMgr.absolutePath, modePerm)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError returned by os.MkdirAll(%v.absolutePath, modePerm).\n"+
      "%v.absolutePath='%v'\nmodePerm=\"drwxrwxrwx\"\n"+
      "Error='%v'\n",
      dMgrLabel, dMgr.absolutePath, err.Error())
  }

  dMgrPathDoesExist,
    _,
    err =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err != nil {
    return fmt.Errorf("Error: After attempted directory creation, "+
      "a non-path error was generated!\n"+
      "%v.absolutePath='%v'\n"+
      "Error='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())
  }

  if !dMgrPathDoesExist {
    return fmt.Errorf("Error: After attempted directory creation,\n"+
      "the directory DOES NOT EXIST!\n"+
      "%v=%v\n", dMgrLabel, dMgr.absolutePath)
  }

  return nil
}
