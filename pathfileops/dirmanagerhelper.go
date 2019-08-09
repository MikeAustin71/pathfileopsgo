package pathfileops

import (
  "errors"
  "fmt"
  "io"
  "os"
  pf "path/filepath"
  "runtime"
  "strings"
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
  copyEmptyDirectory bool,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (dirCopyStats DirectoryCopyStats, errs []error) {

  errs = make([]error, 0, 300)

  ePrefixCurrMethod := "dirMgrHelper.copyDirectory() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2, err3 error
  var dirPathDoesExist, targetPathDoesExist, dirCreated bool

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
    return dirCopyStats, errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nThe current DirMgr path DOES NOT EXIST!\n"+
      "%v.absolutePath='%v'\n",
      dMgrLabel, dMgr.absolutePath)

    errs = append(errs, err)
    return dirCopyStats, errs
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
    return dirCopyStats, errs
  }

  if !targetPathDoesExist && copyEmptyDirectory {

    dirCreated,
      err = dMgrHlpr.lowLevelMakeDir(
      targetDMgr,
      ePrefix,
      "targetDMgr")

    if err != nil {
      errs = append(errs, err)
      return dirCopyStats, errs
    }

    if dirCreated {
      dirCopyStats.DirCreated++
    }

    targetPathDoesExist = true
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)

    return dirCopyStats, errs
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
      return dirCopyStats, errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        // We don't care about sub-directories
        continue

      }

      dirCopyStats.TotalFilesProcessed++

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
        dirCopyStats.FilesNotCopied++
        dirCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())
        continue

      } else {

        // We have a match

        // Create Directory if needed
        if !targetPathDoesExist {

          dirCreated,
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

          if dirCreated {
            dirCopyStats.DirCreated++
          }
        }

        src = dMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        target = targetDMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        err = dMgrHlpr.lowLevelCopyFile(
          src,
          nameFInfo,
          target,
          ePrefix,
          "srcFile",
          "destinationFile")

        if err != nil {
          errs = append(errs, err)
          dirCopyStats.FilesNotCopied++
          dirCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())

        } else {
          dirCopyStats.FilesCopied++
          dirCopyStats.FileBytesCopied += uint64(nameFInfo.Size())
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

  return dirCopyStats, errs
}

// copyDirectoryTree - Helper method for 'DirMgr'. This method is
// designed to copy entire directory trees.
//
func (dMgrHlpr *dirMgrHelper) copyDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  copyEmptyDirectories bool,
  skipTopLevelDirectory bool,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (dTreeCopyStats DirTreeCopyStats, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.copyDirectoryTree() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if targetDMgr == nil {

    err:= fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil'!\n", targetDMgrLabel)

    errs = append(errs, err)

    return dTreeCopyStats, errs

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

    return dTreeCopyStats, errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: %v directory path DOES NOT EXIST!\n"+
      "%v='%v'\n\n",
      dMgrLabel,
      dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return dTreeCopyStats, errs
  }

  _,
  _,
  err2 :=
    dMgrHlpr.doesDirectoryExist(
      targetDMgr,
      PreProcPathCode.None(),
      ePrefix,
      targetDMgrLabel)

  if err2 != nil {

    errs = append(errs, err2)
    return dTreeCopyStats, errs
  }

  baseDirLen := len(dMgr.absolutePath)
  var nextTargetDMgr DirMgr
  var nameFileInfos []os.FileInfo
  osPathSepStr := string(os.PathSeparator)
  dirs := DirMgrCollection{}
  var nextDir DirMgr
  var dirPtr *os.File
  dirPtr = nil
  var srcFile, targetFile string
  fh := FileHelper{}

  dirs.AddDirMgr(dMgrHlpr.copyOut(dMgr))

  if !skipTopLevelDirectory {
    dTreeCopyStats.TotalDirsScanned++
  }

  dirCreated := false
  mainLoopIsDone := false
  file2LoopIsDone := false
  isMatch := false
  isTopLevelDir := true
  isNewDir := false
  isFirstLoop := true

  for !mainLoopIsDone {

    if isFirstLoop {
      isTopLevelDir = true
      isFirstLoop = false
    } else {
      isTopLevelDir = false
    }

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)
      }

      dirPtr = nil
    }

    nextDir, err = dirs.PopFirstDirMgr()

    if err != nil && err == io.EOF {
      mainLoopIsDone = true
      isTopLevelDir = false
      break
    } else if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirs.PopFirstDirMgr().\n"+
        "Error='%v'\n", err.Error())
      errs = append(errs, err2)
      return dTreeCopyStats, errs
    }

    nextTargetDMgr, err = DirMgr{}.New(
      targetDMgr.absolutePath +
        nextDir.absolutePath[baseDirLen:])

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError return by DirMgr{}.New(%v.absolutePath + "+
        "nextDir.absolutePath[baseDirLen:])\n"+
        "%v.absolutePath='%v'\n"+
        "nextDir.absolutePath[baseDirLen:]='%v'\n"+
        "Error='%v'\n\n",
        targetDMgrLabel, targetDMgrLabel,
        targetDMgr.absolutePath,
        nextDir.absolutePath[baseDirLen:],
        err.Error())

      errs = append(errs, err2)
      continue
    }

    dirCreated = false

    if isTopLevelDir &&
      !skipTopLevelDirectory &&
      copyEmptyDirectories {

      dirCreated,
        err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "1-nextTargetDMgr")

    } else if !isTopLevelDir && copyEmptyDirectories {

      dirCreated,
        err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "2-nextTargetDMgr")

    } else {
      err = nil
    }

    if err != nil {

      err2 = fmt.Errorf("\n"+ePrefix+
        "\nError creating target next directory!\n"+
        "Target Next Directory='%v'\nError='%v'\n\n",
        nextTargetDMgr.absolutePath, err.Error())

      errs = append(errs, err2)
      isTopLevelDir = false
      continue

    } else if dirCreated {
      dTreeCopyStats.DirsCreated++
    }

    if !skipTopLevelDirectory && copyEmptyDirectories {
      dTreeCopyStats.DirsCopied++
    } else if skipTopLevelDirectory &&
      copyEmptyDirectories &&
      !isTopLevelDir {

      dTreeCopyStats.DirsCopied++
    }

    isNewDir = true

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(nextDir.absolutePath).\n"+
        "nextDir.absolutePath='%v'\nError='%v'\n\n",
        nextDir.absolutePath, err.Error())

      errs = append(errs, err2)
      continue
    }

    file2LoopIsDone = false

    for !file2LoopIsDone {

      nameFileInfos, err = dirPtr.Readdir(1000)

      if err != nil && err == io.EOF {

        file2LoopIsDone = true

        if len(nameFileInfos) == 0 {

          break
        }

      } else if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdir(1000).\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        file2LoopIsDone = true

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {
          // This is a directory

          err = dirs.AddDirMgrByPathNameStr(
            nextDir.absolutePath +
              osPathSepStr +
              nameFInfo.Name())

          if err != nil {
            err2 = fmt.Errorf(ePrefix+
              "\nError returned by dirs.AddDirMgrByPathNameStr(newDir).\n"+
              "newDir='%v'\nError='%v'\n\n",
              nextDir.absolutePath+osPathSepStr+nameFInfo.Name(),
              err.Error())

            errs = append(errs, err2)
            continue
          }

          // Count Directories Processed
          dTreeCopyStats.TotalDirsScanned++

          continue
        } // End of IsDir()

        // This is a file
        if isTopLevelDir && skipTopLevelDirectory {
          // Skip all files in the
          // parent directory.
          continue
        }

        // This is a file eligible for
        // matching with selection criteria
        dTreeCopyStats.TotalFilesProcessed++

        // Determine if it matches the find file criteria.
        isMatch, err =
          fh.FilterFileName(nameFInfo, fileSelectCriteria)

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "\nError returned by fh.FilterFileName(nameFInfo, fileSelectCriteria).\n"+
            "Directory='%v'\nFile Name='%v'\nError='%v'\n\n",
            nextDir.absolutePath, nameFInfo.Name(), err.Error())

          dTreeCopyStats.TotalFilesProcessed--

          errs = append(errs, err2)

          continue
        }

        if isMatch {
          // This file is a Match!
          dirCreated = false

          dMgrPathDoesExist,
            _,
            err =
            dMgrHlpr.doesDirectoryExist(
              &nextTargetDMgr,
              PreProcPathCode.None(),
              ePrefix,
              "nextTargetDMgr")

          if err != nil {
            file2LoopIsDone = true
            errs = append(errs, err)
            break
          }

          // Create Directory if needed
          if !dMgrPathDoesExist {

            dirCreated,
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
              file2LoopIsDone = true
              break

            } else if dirCreated {
              dTreeCopyStats.DirsCreated++
              dTreeCopyStats.DirsCopied++
            }

          } else if isNewDir && !copyEmptyDirectories {

            dTreeCopyStats.DirsCopied++
          }

          isNewDir = false

          srcFile = nextDir.absolutePath +
            osPathSepStr + nameFInfo.Name()

          targetFile = nextTargetDMgr.absolutePath +
            osPathSepStr + nameFInfo.Name()

          err = dMgrHlpr.lowLevelCopyFile(
            srcFile,
            nameFInfo,
            targetFile,
            ePrefix,
            "srcFile",
            "destinationFile")

          if err != nil {
            errs = append(errs, err)
            dTreeCopyStats.FilesNotCopied++
            dTreeCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())

          } else {
            dTreeCopyStats.FilesCopied++
            dTreeCopyStats.FileBytesCopied += uint64(nameFInfo.Size())
          }

        } else {
          // This file is NOT A Match
          // NOT Selected File
          dTreeCopyStats.FilesNotCopied++
          dTreeCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())

        }
      } // End of range nameFileInfos
    } // End of file 2 Loop

    isTopLevelDir = false
  } // End of main loop

  // Final verification of
  dMgrPathDoesExist,
    _,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    targetDMgr.absolutePath,
    ePrefix,
    "targetDMgr")

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nAfter Copy Operation 'targetDMgr' path returned non-path error!\n"+
      "Error='%v'\n\n", err.Error())

    errs = append(errs, err2)
  }

  if dTreeCopyStats.FilesCopied > 0 &&
    !dMgrPathDoesExist {

    err2 = fmt.Errorf(ePrefix+
      "\nERROR: The copy operation failed to create\n"+
      "the 'targetDMgr' path. 'targetDMgr' path DOES NOT EXIST!\n"+
      "targetDMgr Path='%v'\n\n",
      targetDMgr.absolutePath)
    errs = append(errs, err2)
  }

  if dTreeCopyStats.TotalFilesProcessed !=
    dTreeCopyStats.FilesCopied+dTreeCopyStats.FilesNotCopied {

    err2 = fmt.Errorf(ePrefix+
      "\nFile Counting Error: Number of Total Files Processed\n"+
      "NOT EQUAL to Number of Files Copied Plus Number of Files NOT copied!\n"+
      "Total Number of Files Processed='%v'\n"+
      "         Number of Files Copied='%v'\n"+
      "     Number of Files NOT Copied='%v'\n\n",
      dTreeCopyStats.TotalFilesProcessed,
      dTreeCopyStats.FilesCopied,
      dTreeCopyStats.FilesNotCopied)

    dTreeCopyStats.ComputeError = fmt.Errorf("%v", err2.Error())

    errs = append(errs, err2)
  }

  return dTreeCopyStats, errs
}

// CopyIn - Receives a pointer to a incoming DirMgr object
// ('dMgrIn') as an input parameter and copies the values from
// the incoming object to the input parameter, 'dMgr'.
//
// When the copy operation is the 'dMgr' object is a duplicate
// of the incoming DirMgr object ('dMgrIn').
//
func (dMgrHlpr *dirMgrHelper) copyIn(
  dMgr *DirMgr,
  dMgrIn *DirMgr) {

  if dMgr == nil {
    dMgr = &DirMgr{}
  }

  if dMgrIn == nil {
    dMgrIn = &DirMgr{}
  }

  dMgr.isInitialized = dMgrIn.isInitialized
  dMgr.originalPath = dMgrIn.originalPath
  dMgr.path = dMgrIn.path
  dMgr.isPathPopulated = dMgrIn.isPathPopulated
  dMgr.doesPathExist = dMgrIn.doesPathExist
  dMgr.parentPath = dMgrIn.parentPath
  dMgr.isParentPathPopulated = dMgrIn.isParentPathPopulated
  dMgr.absolutePath = dMgrIn.absolutePath
  dMgr.isAbsolutePathPopulated = dMgrIn.isAbsolutePathPopulated
  dMgr.doesAbsolutePathExist = dMgrIn.doesAbsolutePathExist
  dMgr.isAbsolutePathDifferentFromPath = dMgrIn.isAbsolutePathDifferentFromPath
  dMgr.directoryName = dMgrIn.directoryName
  dMgr.volumeName = dMgrIn.volumeName
  dMgr.isVolumePopulated = dMgrIn.isVolumePopulated
  dMgr.actualDirFileInfo = dMgrIn.actualDirFileInfo.CopyOut()

}

// copyOut - Makes a duplicate copy of input parameter
// 'dMgr' values and returns them as a new DirMgr object.
//
func (dMgrHlpr *dirMgrHelper) copyOut(
  dMgr *DirMgr) DirMgr {

  if dMgr == nil {
    dMgr = &DirMgr{}
  }

  dOut := DirMgr{}

  dOut.isInitialized = dMgr.isInitialized
  dOut.originalPath = dMgr.originalPath
  dOut.path = dMgr.path
  dOut.isPathPopulated = dMgr.isPathPopulated
  dOut.doesPathExist = dMgr.doesPathExist
  dOut.parentPath = dMgr.parentPath
  dOut.isParentPathPopulated = dMgr.isParentPathPopulated
  dOut.absolutePath = dMgr.absolutePath
  dOut.isAbsolutePathPopulated = dMgr.isAbsolutePathPopulated
  dOut.doesAbsolutePathExist = dMgr.doesAbsolutePathExist
  dOut.isAbsolutePathDifferentFromPath = dMgr.isAbsolutePathDifferentFromPath
  dOut.directoryName = dMgr.directoryName
  dOut.volumeName = dMgr.volumeName
  dOut.isVolumePopulated = dMgr.isVolumePopulated
  dOut.actualDirFileInfo = dMgr.actualDirFileInfo.CopyOut()

  return dOut
}

// deleteAllFilesInDirectory - Helper method used by DirMgr. This
// method deletes ALL files in the current directory. ONLY files
// in the current directory are deleted. Files in sub-directories
// are NOT deleted.
//
func (dMgrHlpr *dirMgrHelper) deleteAllFilesInDirectory(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (deleteDirStats DeleteDirFilesStats, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.deleteAllFilesInDirectory() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  errs = make([]error, 0, 300)
  var err2 error
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
    return deleteDirStats, errs
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
    return deleteDirStats, errs
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return deleteDirStats, errs
  }

  deleteDirStats.TotalDirsScanned = 1

  var nameFileInfos []os.FileInfo

  file2LoopIsDone := false

  isNewDir := true

  for !file2LoopIsDone {

    nameFileInfos, err = dirPtr.Readdir(1000)

    if err != nil && err == io.EOF {

      file2LoopIsDone = true

      if len(nameFileInfos) == 0 {

        break
      }

    } else if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Readdirnames(1000).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath,
        err.Error())

      errs = append(errs, err2)

      file2LoopIsDone = true

      break
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {

        deleteDirStats.TotalSubDirectories++

        continue

      } else {

        deleteDirStats.TotalFilesProcessed++

        if !nameFInfo.Mode().IsRegular() {
          err2 = fmt.Errorf(ePrefix+
            "\nError: fileName is NOT classified as a 'Regular' File!\n"+
            "fileName='%v'",
            dMgr.absolutePath+osPathSepStr+nameFInfo.Name())
          errs = append(errs, err2)
          deleteDirStats.FilesRemaining++
          deleteDirStats.FilesRemainingBytes += uint64(nameFInfo.Size())
          continue
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

          deleteDirStats.FilesRemaining++
          deleteDirStats.FilesRemainingBytes += uint64(nameFInfo.Size())

          errs = append(errs, err2)
        } else {

          deleteDirStats.FilesDeleted++
          deleteDirStats.FilesDeletedBytes += uint64(nameFInfo.Size())

          if isNewDir {
            isNewDir = false
            deleteDirStats.NumOfDirsWhereFilesDeleted++
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
        "An attempt to close the os.File pointer to the current\n"+
        "%v path has FAILED!\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())
      errs = append(errs, err2)
    }
  }

  return deleteDirStats, errs
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

// deleteDirectoryTreeInfo - Helper method similar to
// dirMgrHelper.deleteDirectoryTreeStats(). This method
// will optionally delete files in the entire directory
// tree, the top level directory or exclusively in the
// sub-directory tree.
//
// This differs from similar methods in that it returns
// a type DirectoryDeleteFileInfo.
//
func (dMgrHlpr *dirMgrHelper) deleteDirectoryTreeInfo(
  dMgr *DirMgr,
  deleteFileSelectionCriteria FileSelectionCriteria,
  skipTopLevelDirectory bool,
  scanSubDirectories bool,
  ePrefix string,
  dMgrLabel string,
  deleteFileSelectLabel string) (DirectoryDeleteFileInfo, []error) {

  deleteTreeInfo := DirectoryDeleteFileInfo{}
  errs := make([]error, 0, 300)
  ePrefixCurrMethod := "dirMgrHelper.deleteDirectoryTreeInfo() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if skipTopLevelDirectory &&
    !scanSubDirectories {

    err := errors.New(ePrefix +
      "\nERROR: Conflicted Input parameters! skipTopLevelDirectory=true and scanSubDirectories=false.\n" +
      "Impossible combination!!\n")

    errs = append(errs, err)
    return deleteTreeInfo, errs
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
    return deleteTreeInfo, errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)
    return deleteTreeInfo, errs
  }

  var err2 error
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  var dirPtr *os.File
  dirPtr = nil
  fh := FileHelper{}
  var nextDir *DirMgr
  file2LoopIsDone := false
  isMatch := false
  isTopLevelDir := true

  deleteTreeInfo.StartPath = dMgr.absolutePath
  deleteTreeInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria
  deleteTreeInfo.Directories.AddDirMgr(dMgrHlpr.copyOut(dMgr))
  dTreeCnt := 1

  for i := 0; i < dTreeCnt; i++ {

    if i == 0 {
      isTopLevelDir = true
    } else {
      isTopLevelDir = false
    }

    nextDir, err = deleteTreeInfo.Directories.GetDirMgrAtIndex(i)

    if err != nil {
      errs = append(errs, err)
      break
    }

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath). "+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
      dirPtr = nil
      continue
    }

    file2LoopIsDone = false

    for !file2LoopIsDone {

      nameFileInfos, err = dirPtr.Readdir(1000)

      lNameFileInfos := len(nameFileInfos)

      if err != nil && err == io.EOF {

        file2LoopIsDone = true

        if lNameFileInfos == 0 {
          break
        }

      } else if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdir(1000).\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        file2LoopIsDone = true
        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          if !scanSubDirectories {
            continue
          }

          err =
            deleteTreeInfo.Directories.AddDirMgrByKnownPathDirName(
              nextDir.absolutePath,
              nameFInfo.Name())

          if err != nil {
            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by dirs.AddDirMgrByKnownPathDirName(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

          dTreeCnt++

        } else {
          // This is a file which is eligible for processing

          if isTopLevelDir && skipTopLevelDirectory {
            continue
          }

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, deleteFileSelectionCriteria)

          if err != nil {

            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by fh.FilterFileName(nameFInfo, %v).\n"+
                "%v directory searched='%v'\nfileName='%v'\nError='%v'\n\n",
                deleteFileSelectLabel, dMgrLabel,
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match, save file to deleteTreeInfo
            fileToDelete := nextDir.absolutePath + osPathSepStr + nameFInfo.Name()

            err = os.Remove(fileToDelete)

            if err != nil {
              err2 = fmt.Errorf(ePrefix+
                "\nError returned by os.Remove(fileToDelete)\n"+
                "fileToDelete='%v'\nError='%v'\n",
                fileToDelete,
                err.Error())
              errs = append(errs, err2)
              continue
            }

            err = deleteTreeInfo.DeletedFiles.AddFileMgrByDirFileNameExt(nextDir.CopyOut(), nameFInfo.Name())

            if err != nil {
              err2 = fmt.Errorf(ePrefix+
                "\nERROR returned by deleteTreeInfo.DeletedFiles.AddFileMgrByDirFileNameExt(nextDir, fileNameExt)\n"+
                "nextDir='%v'\n"+
                "fileNameExt='%v'"+
                "Error='%v'\n\n",
                nextDir.absolutePath,
                nameFInfo.Name(),
                err.Error())

              errs = append(errs, err2)
            }
          }
        }

      } // End of nameFInfo := range nameFileInfos
    } // End of for !file2LoopIsDone

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)
      }

      dirPtr = nil
    }

  } // End of for i:=0; i < dTreeCnt; i ++

  if len(deleteTreeInfo.Directories.dirMgrs) > 0 && skipTopLevelDirectory {
    _, _ = deleteTreeInfo.Directories.PopFirstDirMgr()
  }

  for i := 0; i < len(errs); i++ {
    err2 = fmt.Errorf("%v", errs[i].Error())
    deleteTreeInfo.ErrReturns =
      append(deleteTreeInfo.ErrReturns, err2)
  }

  return deleteTreeInfo, errs
}

// deleteDirectoryTreeStats - Helper method designed to delete files using
// file selection criteria. Scope of scans and file deletions is controlled
// by input parameter 'scanSubDirectories'. If set to 'true' files may be
// deleted in the entire directory tree. If set to 'false' the file deletions
// are limited solely to the directory identified by the current 'DirMgr'
// instance.
//
func (dMgrHlpr *dirMgrHelper) deleteDirectoryTreeStats(
  dMgr *DirMgr,
  skipTopLevelDirectory,
  scanSubDirectories bool,
  deleteFileSelectionCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  deleteSelectionLabel string) (deleteDirStats DeleteDirFilesStats, errs []error) {
  ePrefixCurrMethod := "dirMgrHelper.deleteDirectoryTreeStats() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if skipTopLevelDirectory &&
    !scanSubDirectories {

    err := errors.New(ePrefix +
      "\nERROR: Conflicted Input parameters! skipTopLevelDirectory=true and scanSubDirectories=false.\n" +
      "Impossible combination!!\n")

    errs = append(errs, err)
    return deleteDirStats, errs
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
    return deleteDirStats, errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return deleteDirStats, errs
  }

  var err2 error
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dirPtr *os.File
  dirPtr = nil
  fh := FileHelper{}
  var nextDir DirMgr
  mainLoopIsDone := false
  file2LoopIsDone := false
  isMatch := false
  isNewDir := false
  isTopLevelDir := true
  isFirstLoop := true

  dirs.AddDirMgr(dMgrHlpr.copyOut(dMgr))

  for !mainLoopIsDone {

    if isFirstLoop {
      isTopLevelDir = true
      isFirstLoop = false
    } else {
      isTopLevelDir = false
    }

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)
      }

      dirPtr = nil
    }

    nextDir, err = dirs.PopFirstDirMgr()

    if err != nil && err == io.EOF {
      mainLoopIsDone = true
      break
    } else if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirs.PopFirstDirMgr().\n"+
        "Error='%v'\n", err.Error())
      errs = append(errs, err2)
      return deleteDirStats, errs
    }

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath). "+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
      continue
    }

    deleteDirStats.TotalDirsScanned++

    isNewDir = true
    file2LoopIsDone = false

    for !file2LoopIsDone {

      nameFileInfos, err = dirPtr.Readdir(1000)

      if err != nil && err == io.EOF {

        file2LoopIsDone = true

        if len(nameFileInfos) == 0 {

          break
        }

      } else if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdir(1000).\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        file2LoopIsDone = true

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          deleteDirStats.TotalSubDirectories++

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

          if isTopLevelDir && skipTopLevelDirectory {
            continue
          }

          deleteDirStats.TotalFilesProcessed++

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

            deleteDirStats.FilesRemaining++
            deleteDirStats.FilesRemainingBytes += uint64(nameFInfo.Size())

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

              deleteDirStats.FilesRemaining++
              deleteDirStats.FilesRemainingBytes += uint64(nameFInfo.Size())

              errs = append(errs, err2)

            } else {
              deleteDirStats.FilesDeleted++
              deleteDirStats.FilesDeletedBytes += uint64(nameFInfo.Size())

              if isNewDir {
                deleteDirStats.NumOfDirsWhereFilesDeleted++
              }

              isNewDir = false

            }
          }
        }

      } // End of nameFInfo := range nameFileInfos
    } // End of for !file2LoopIsDone
  } // End of for !mainLoopIsDone

  return deleteDirStats, errs
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
  fileSearchLabel string) (deleteDirStats DeleteDirFilesStats, errs []error) {

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
    return deleteDirStats, errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return deleteDirStats, errs
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
    return deleteDirStats, errs
  }

  if errCode == -2 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' consists of blank spaces!\n",
      fileSearchLabel)

    errs = append(errs, err2)
    return deleteDirStats, errs
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath, err.Error())

    errs = append(errs, err2)
    return deleteDirStats, errs
  }

  deleteDirStats.TotalDirsScanned++

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
        "\nError returned by dirPtr.Readdirnames(1000).\n"+
        "#v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath, err3.Error())

      errs = append(errs, err2)
      return deleteDirStats, errs
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        deleteDirStats.TotalSubDirectories++
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

        deleteDirStats.TotalFilesProcessed++

        if !isMatch {
          deleteDirStats.FilesRemaining++
          deleteDirStats.FilesRemainingBytes += uint64(nameFInfo.Size())
          continue

        } else {

          err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFInfo.Name())

          if err != nil {
            err2 = fmt.Errorf(ePrefix+
              "\nError returned by os.Remove(pathFileName).\n"+
              "pathFileName='%v'\nError='%v'\n\n",
              dMgr.absolutePath+osPathSepStr+nameFInfo.Name(),
              err.Error())

            deleteDirStats.FilesRemaining++
            deleteDirStats.FilesRemainingBytes -= uint64(nameFInfo.Size())

            errs = append(errs, err2)
            continue
          }

          deleteDirStats.FilesDeleted++
          deleteDirStats.FilesDeletedBytes += uint64(nameFInfo.Size())
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

  return deleteDirStats, errs
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

  if dMgr == nil {
    err = fmt.Errorf(ePrefix +
      "\nError: %v pointer is 'nil'!\n", dMgrLabel)
    return dirPathDoesExist, fInfo, err
  }

  if !dMgr.isInitialized {
    err = errors.New(ePrefix +
      "\nError: DirMgr is NOT Initialized.\n")
    return dirPathDoesExist, fInfo, err
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, dMgr.absolutePath =
    fh.isStringEmptyOrBlank(dMgr.absolutePath)

  if errCode == -1 {
    dMgr.isInitialized = false
    dMgr.absolutePath = ""
    dMgr.path = ""
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v'.absolutePath is an empty string!\n\n",
      dMgrLabel)
    return dirPathDoesExist, fInfo, err
  }

  if errCode == -2 {
    dMgr.isInitialized = false
    dMgr.absolutePath = ""
    dMgr.path = ""
    dMgr.doesAbsolutePathExist = false
    dMgr.doesPathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
    dirPathDoesExist = false
    err = fmt.Errorf(ePrefix+
      "\nError: Input parameter '%v' consists of blank spaces!\n\n",
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
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel, dMgrLabel, err2.Error())

      return dirPathDoesExist, fInfo, err
    }
  }

  errCode, _, dMgr.path =
    fh.isStringEmptyOrBlank(dMgr.path)

  if errCode < 0 {
    dMgr.path = dMgr.absolutePath
  }

  dMgr.isPathPopulated = true

  strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
  lStr := len(strAry)
  idxStr := strAry[lStr-1]

  idx := strings.Index(dMgr.absolutePath, idxStr)

  dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

  dMgr.isParentPathPopulated = true

  if dMgr.parentPath == "" {
    dMgr.isParentPathPopulated = false
  }

  if idxStr != "" {
    dMgr.directoryName = idxStr
  } else {
    dMgr.directoryName = dMgr.absolutePath
  }

  errCode, _, dMgr.path =
    fh.isStringEmptyOrBlank(dMgr.path)

  if dMgr.path != dMgr.absolutePath {
    dMgr.isAbsolutePathDifferentFromPath = true
  }

  var vn string
  if dMgr.isAbsolutePathPopulated {
    vn = pf.VolumeName(dMgr.absolutePath)
  } else if dMgr.isPathPopulated {
    vn = pf.VolumeName(dMgr.path)
  }

  dMgr.isVolumePopulated = false

  if vn != "" {
    dMgr.isVolumePopulated = true
    dMgr.volumeName = vn
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
      "%v='%v'\n\n",
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
      "%v='%v'\n\n",
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
      "%v.path='%v'\n\n",
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
      "%v.path='%v'\n\n",
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
      "%v.path='%v'\n\n",
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
  dMgr.actualDirFileInfo =
    FileInfoPlus{}.NewFromFileInfo(absFInfo)
  fInfo = dMgr.actualDirFileInfo.CopyOut()
  dirPathDoesExist = true
  err = nil
  return dirPathDoesExist, fInfo, err
}

func (dMgrHlpr *dirMgrHelper) empty(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.empty() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if dMgr == nil {
    return errors.New(ePrefix +
      "\nError: Input parameter 'dMgr' pointer is 'nil'!\n")
  }

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

// equal - Compares two DirMgr objects to determine if
// they are equal.
//
func (dMgrHlpr *dirMgrHelper) equal(
  dMgr *DirMgr,
  dMgr2 *DirMgr) bool {

  if dMgr==nil || dMgr2 == nil {
    return false
  }

  if dMgr.isInitialized != dMgr2.isInitialized ||
    dMgr.originalPath != dMgr2.originalPath ||
    dMgr.path != dMgr2.path ||
    dMgr.isPathPopulated != dMgr2.isPathPopulated ||
    dMgr.doesPathExist != dMgr2.doesPathExist ||
    dMgr.parentPath != dMgr2.parentPath ||
    dMgr.isParentPathPopulated != dMgr2.isParentPathPopulated ||
    dMgr.absolutePath != dMgr2.absolutePath ||
    dMgr.isAbsolutePathPopulated != dMgr2.isAbsolutePathPopulated ||
    dMgr.doesAbsolutePathExist != dMgr2.doesAbsolutePathExist ||
    dMgr.isAbsolutePathDifferentFromPath != dMgr2.isAbsolutePathDifferentFromPath ||
    dMgr.directoryName != dMgr2.directoryName ||
    dMgr.volumeName != dMgr2.volumeName ||
    dMgr.isVolumePopulated != dMgr2.isVolumePopulated {

    return false
  }

  if !dMgr.actualDirFileInfo.Equal(&dMgr2.actualDirFileInfo) {
    return false
  }

  return true

}

// EqualAbsPaths - compares the absolute paths for the input
// parameter 'dMgr' and the input parameter ('dMgr2').
//
// If the two absolute paths are equal, the method returns 'true'.
// If the two absolute paths are NOT equal, the method returns 'false'.
// The comparison is NOT case sensitive. In other words, both paths
// are converted to lower case before making the comparision.
//
// If either the input parameter ('dMgr') or the input parameter
// 'dMgr2' are uninitialized, a value of 'false' is returned.
//
func (dMgrHlpr *dirMgrHelper) equalAbsolutePaths(
  dMgr *DirMgr,
  dMgr2 *DirMgr) bool {

  if !dMgr.isInitialized || !dMgr2.isInitialized {
    return false
  }

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  lcDMgrPath := strings.ToLower(dMgr.absolutePath)

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr2,
    PreProcPathCode.None(),
    "",
    "")

  lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  return true
}

// equalPaths - Compares two DirMgr objects to determine
// if their paths are equal. Both Directory Path and
// absolute path must be equivalent.
//
// If the compared paths are equal, the method returns 'true'.
// If the paths are NOT equal, the method returns 'false'.
// The comparisons are NOT case sensitive. In other words, all paths
// are converted to lower case before making the comparisons.
//
// If either the current DirMgr ('dMgr') or the input parameter
// 'dMgr2' are uninitialized, a value of 'false' is returned.
//
func (dMgrHlpr *dirMgrHelper) equalPaths(
  dMgr *DirMgr,
  dMgr2 *DirMgr) bool {

  if !dMgr.isInitialized || !dMgr2.isInitialized {
    return false
  }

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    "",
    "")

  lcDMgrPath := strings.ToLower(dMgr.absolutePath)

  _,
    _,
    _ = dMgrHlpr.doesDirectoryExist(
    dMgr2,
    PreProcPathCode.None(),
    "",
    "")

  lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  lcDMgrPath = strings.ToLower(dMgr.path)
  lcDMgr2Path = strings.ToLower(dMgr2.path)

  if lcDMgrPath != lcDMgr2Path {
    return false
  }

  return true
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
      dMgr.absolutePath,
      srcFileNameExt,
      targetBaseDir.absolutePath,
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
        dMgr.absolutePath,
        srcFileNameExt,
        targetDirLabel,
        targetBaseDir.absolutePath,
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

// findDirectoryTreeFiles - A multifunctional helper method which
// can be used to scan a parent directory or an entire directory
// tree to locate files which match the file selection criteria.
//
func (dMgrHlpr *dirMgrHelper) findDirectoryTreeFiles(
  dMgr *DirMgr,
  fileSelectionCriteria FileSelectionCriteria,
  skipTopLevelDirectory bool,
  scanSubDirectories bool,
  ePrefix string,
  dMgrLabel string,
  fileSelectLabel string) (dTreeInfo DirectoryTreeInfo, errs []error) {

  dTreeInfo = DirectoryTreeInfo{}
  errs = make([]error, 0, 300)

  ePrefixCurrMethod := "dirMgrHelper.findDirectoryTreeFiles() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if skipTopLevelDirectory &&
    !scanSubDirectories {

    err := fmt.Errorf(ePrefix +
      "\nERROR: Conflicted Input parameters! skipTopLevelDirectory=true and scanSubDirectories=false.\n" +
      "Impossible combination!!\n")

    errs = append(errs, err)
    return dTreeInfo, errs
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
    return dTreeInfo, errs
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return dTreeInfo, errs
  }

  var err2 error
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  var dirPtr *os.File
  dirPtr = nil
  fh := FileHelper{}
  var nextDir *DirMgr
  file2LoopIsDone := false
  isMatch := false
  isTopLevelDir := true

  dTreeInfo.Directories.AddDirMgr(dMgrHlpr.copyOut(dMgr))
  dTreeCnt := 1

  for i := 0; i < dTreeCnt; i++ {

    if i == 0 {
      isTopLevelDir = true
    } else {
      isTopLevelDir = false
    }

    nextDir, err = dTreeInfo.Directories.GetDirMgrAtIndex(i)

    if err != nil {
      errs = append(errs, err)
      break
    }

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath). "+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel, dMgrLabel,
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
      dirPtr = nil
      continue
    }

    file2LoopIsDone = false

    for !file2LoopIsDone {

      nameFileInfos, err = dirPtr.Readdir(1000)

      lNameFileInfos := len(nameFileInfos)

      if err != nil && err == io.EOF {

        file2LoopIsDone = true

        if lNameFileInfos == 0 {
          break
        }

      } else if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdir(1000).\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        file2LoopIsDone = true
        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          if !scanSubDirectories {
            continue
          }

          err = dTreeInfo.Directories.AddDirMgrByKnownPathDirName(nextDir.absolutePath, nameFInfo.Name())

          if err != nil {
            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by dirs.AddDirMgrByKnownPathDirName(newDirPathFileName).\n"+
                "newDirPathFileName='%v'\nError='%v'\n\n",
                nextDir.absolutePath+osPathSepStr+nameFInfo.Name(), err.Error())

            errs = append(errs, err2)
            continue
          }

          dTreeCnt++

        } else {
          // This is a file which is eligible for processing

          if isTopLevelDir && skipTopLevelDirectory {
            continue
          }

          // This is not a directory. It is a file.
          // Determine if it matches the find file criteria.
          isMatch, err =
            fh.FilterFileName(nameFInfo, fileSelectionCriteria)

          if err != nil {

            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by fh.FilterFileName(nameFInfo, %v).\n"+
                "%v directory searched='%v'\nfileName='%v'\nError='%v'\n\n",
                fileSelectLabel, dMgrLabel,
                dMgr.absolutePath, nameFInfo.Name(), err.Error())

            errs = append(errs, err2)

            continue
          }

          if !isMatch {

            continue

          } else {

            // We have a match, save file to dTreeInfo

            err = dTreeInfo.FoundFiles.AddFileMgrByDirFileNameExt(nextDir.CopyOut(), nameFInfo.Name())

            if err != nil {
              err2 = fmt.Errorf(ePrefix+
                "\nERROR returned by dTreeInfo.FoundFiles.AddFileMgrByDirFileNameExt(nextDir, fileNameExt)\n"+
                "nextDir='%v'\n"+
                "fileNameExt='%v'"+
                "Error='%v'\n\n",
                nextDir.absolutePath,
                nameFInfo.Name(),
                err.Error())

              errs = append(errs, err2)

            }
          }
        }

      } // End of nameFInfo := range nameFileInfos
    } // End of for !file2LoopIsDone

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)
      }

      dirPtr = nil
    }

  } // End of for !mainLoopIsDone

  if len(dTreeInfo.Directories.dirMgrs) > 0 && skipTopLevelDirectory {
    _, _ = dTreeInfo.Directories.PopFirstDirMgr()
  }

  return dTreeInfo, errs
}

// findDirectoryTreeStats - Scans the parent directory
// or the entire directory tree to calculate and
// return directory information.
//
func (dMgrHlpr *dirMgrHelper) findDirectoryTreeStats(
  dMgr *DirMgr,
  skipTopLevelDirectory bool,
  scanSubDirectories bool,
  ePrefix string,
  dMgrLabel string) (dTreeStats DirectoryStatsDto, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.findDirectoryTreeStats() "

  dTreeStats = DirectoryStatsDto{}
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

    return dTreeStats, errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: %v directory path DOES NOT EXIST!\n"+
      "%v='%v'\n\n",
      dMgrLabel,
      dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return dTreeStats, errs
  }

  var err2 error
  dirs := DirMgrCollection{}
  var nameFileInfos []os.FileInfo
  var nextDir DirMgr
  var dirPtr *os.File
  mainLoopIsDone := false
  isFirstLoop := true
  isTopLevelDir := true
  file2LoopIsDone := false

  dirs.AddDirMgr(dMgrHlpr.copyOut(dMgr))

  for !mainLoopIsDone {

    if isFirstLoop {
      isTopLevelDir = true
      isFirstLoop = false
    } else {
      isTopLevelDir = false
    }

    nextDir, err = dirs.PopFirstDirMgr()

    if err != nil && err == io.EOF {
      mainLoopIsDone = true
      break

    } else if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirs.PopFirstDirMgr().\n"+
        "Error='%v'\n", err.Error())
      errs = append(errs, err2)
      return dTreeStats, errs
    }

    dirPtr, err = os.Open(nextDir.absolutePath)

    if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(nextDir.absolutePath).\n"+
        "nextDir.absolutePath='%v'\nError='%v'\n\n",
        nextDir.absolutePath, err.Error())

      errs = append(errs, err2)
      continue
    }

    file2LoopIsDone = false

    for !file2LoopIsDone {

      nameFileInfos, err = dirPtr.Readdir(1000)

      if err != nil && err == io.EOF {

        file2LoopIsDone = true

        if len(nameFileInfos) == 0 {

          break
        }

      } else if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Readdir(1000).\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        file2LoopIsDone = true

        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {
          // This is a directory
          err = dirs.AddDirMgrByKnownPathDirName(
            nextDir.absolutePath,
            nameFInfo.Name())

          if err != nil {
            errs = append(errs, err2)
            continue
          }

          dTreeStats.numOfSubDirs++

        } else {

          if isTopLevelDir && skipTopLevelDirectory {
            continue
          }

          // This is a file
          dTreeStats.numOfFiles++
          dTreeStats.numOfBytes += uint64(nameFInfo.Size())
        }
      } // for _, nameFInfo := range nameFileInfos
    } // for !file2LoopIsDone

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close()\n"+
          "Error='%v'\n\n", err.Error())

        errs = append(errs, err2)

        mainLoopIsDone = true
        break
      }

      dirPtr = nil
    }

    if isTopLevelDir && !scanSubDirectories {
      mainLoopIsDone = true
      break
    }

  } // for !mainLoopIsDone

  return dTreeStats, errs
}

// findFilesByNamePattern - Searches files in the current directory ONLY. An attempt
// will be made to match the file name with the specified search pattern string.
// All matched files will be returned in a FileMgrCollection.
//
func (dMgrHlpr *dirMgrHelper) findFilesByNamePattern(
  dMgr *DirMgr,
  fileSearchPattern string,
  ePrefix string,
  dMgrLabel string,
  fileSearchLabel string) (FileMgrCollection,
  error) {

  fileMgrCol := FileMgrCollection{}.New()
  var err, err2, err3 error

  ePrefixCurrMethod := "dirMgrHelper.findFilesByNamePattern() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var dMgrPathDoesExist bool

  dMgrPathDoesExist,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    return fileMgrCol, err
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    return fileMgrCol, err
  }

  fh := FileHelper{}

  errCode := 0

  errCode, _, fileSearchPattern = fh.isStringEmptyOrBlank(fileSearchPattern)

  if errCode < 0 {
    return fileMgrCol,
      fmt.Errorf(ePrefix+
        "\nInput parameter '%v' is INVALID!\n"+
        "'%v' is an EMPTY STRING!\n",
        fileSearchLabel,
        fileSearchLabel)
  }

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {
    return fileMgrCol,
      fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgrLabel,
        dMgr.absolutePath,
        err.Error())
  }

  err3 = nil
  var isMatch bool
  var nameFileInfos []os.FileInfo
  errs := make([]error, 0, 300)

  for err3 != io.EOF {

    nameFileInfos, err3 = dirPtr.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dirPtr.Readdirnames(1000).\n"+
        "%v.absolutePath='%v'\nError='%v'\n",
        dMgrLabel,
        dMgr.absolutePath,
        err3.Error())

      errs = append(errs, err2)
      break
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        continue

      } else {
        // This is a file. Check for pattern match.
        isMatch, err = pf.Match(fileSearchPattern, nameFInfo.Name())

        if err != nil {

          err2 = fmt.Errorf(ePrefix+
            "\nError returned by fp.Match(%v, fileName).\n"+
            "directorySearched='%v' %v='%v' fileName='%v' Error='%v' ",
            fileSearchLabel,
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
          // This file is a match. Process it.
          err = fileMgrCol.AddFileMgrByFileInfo(dMgr.absolutePath, nameFInfo)

          if err != nil {

            err2 = fmt.Errorf(ePrefix+
              "\nError returned by fileMgrCol.AddFileMgrByFileInfo(%v.absolutePath, nameFInfo).\n"+
              "Directory='%v'\nFileName='%v'\nError='%v'\n",
              dMgrLabel,
              dMgr.absolutePath,
              nameFInfo.Name(),
              err.Error())

            errs = append(errs, err2)
            err3 = io.EOF
            break
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
        "dirPtr Path='%v'\nError='%v'\n",
        dMgr.absolutePath, err.Error())
      errs = append(errs, err2)
    }
  }

  return fileMgrCol, fh.ConsolidateErrors(errs)
}

// getAbsolutePathElements - Returns all of the directories and drive
// specifications as an array of strings.
//
// Example
//
// Path = "D:\ADir\BDir\CDir\EDir"
//
// Returned pathElements string array:
//   pathElements[0] = "D:"
//   pathElements[1] = "ADir"
//   pathElements[2] = "BDir"
//   pathElements[3] = "CDir"
//   pathElements[4] = "DDir"
//   pathElements[4] = "EDir"
//
func (dMgrHlpr *dirMgrHelper) getAbsolutePathElements(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (pathElements []string, err error) {

  pathElements = make([]string, 0, 50)
  err = nil
  absolutePath := ""

  _,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {

    return pathElements, err

  }

  absolutePath = dMgr.absolutePath

  absolutePath = strings.Replace(absolutePath, "\\", "/", -1)

  pathElements = strings.Split(absolutePath, "/")

  return pathElements, err
}

// getDirectoryTree - Returns a DirMgrCollection containing all
// the sub-directories in the path of the parent directory identified
// by the input parameter 'dMgr'.
//
// The returned DirMgrCollection will always contain the parent directory
// at the top of the array (index=0). Therefore, if no errors are encountered,
// the returned DirMgrCollection will always consist of at least one directory.
// If sub-directories are found, then the returned DirMgrCollection will
// contain more than one directory.
//
func (dMgrHlpr *dirMgrHelper) getDirectoryTree(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (dirMgrs DirMgrCollection, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.getDirectoryTree() "

  dirMgrs = DirMgrCollection{}.New()

  errs = make([]error, 0, 100)

  var err, err2, err3 error

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var dMgrPathDoesExist bool

  dMgrPathDoesExist,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {

    errs = append(errs, err)

    return dirMgrs, errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return dirMgrs, errs
  }

  dirMgrs.AddDirMgr(dMgrHlpr.copyOut(dMgr))

  fh := FileHelper{}

  maxLen := dirMgrs.GetNumOfDirs()

  var dirPtr *os.File
  var nameFileInfos []os.FileInfo

  for i := 0; i < maxLen; i++ {

    dirPtr, err = os.Open(dirMgrs.dirMgrs[i].absolutePath)

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "\nError return by os.Open(dirMgrs.dirMgrs[%v].absolutePath). "+
        "dMgr.absolutePath='%v'\nError='%v'\n\n",
        i, dirMgrs.dirMgrs[i].absolutePath, err.Error())

      errs = append(errs, err2)
      continue
    }

    err3 = nil

    for err3 != io.EOF {

      nameFileInfos, err3 = dirPtr.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dirPtr.Readdirnames(-1).\n"+
          "dMgr.absolutePath='%v'\nError='%v'\n",
          dMgr.absolutePath, err3.Error())

        errs = append(errs, err2)
        break
      }

      for _, nameFInfo := range nameFileInfos {

        if nameFInfo.IsDir() {

          newDirPathFileName :=
            fh.JoinPathsAdjustSeparators(dirMgrs.dirMgrs[i].absolutePath, nameFInfo.Name())

          err = dirMgrs.AddDirMgrByPathNameStr(newDirPathFileName)

          if err != nil {

            err2 =
              fmt.Errorf("\n"+ePrefix+
                "Error returned by dirMgrs.AddDirMgrByPathNameStr(newDirPathFileName). "+
                "dirPtr='%v' Error='%v' ",
                newDirPathFileName, err.Error())

            errs = append(errs, err2)
            continue
          }

          maxLen = dirMgrs.GetNumOfDirs()
        }
      }
    }

    if dirPtr != nil {

      err = dirPtr.Close()

      if err != nil {

        err2 = fmt.Errorf("\n"+ePrefix+
          "Error returned by dirPtr.Close().\n"+
          "dirPtr='%v'\nError='%v'\n",
          dMgr.absolutePath, err.Error())

        errs = append(errs, err2)
      }
    }
  }

  return dirMgrs, errs
}

// getParentDirMgr - Returns a new Directory Manager instance
// which represents the the parent path for the input Directory
// Manager, 'dMgr'. The 'dMgr' absolute path is used in extracting
// the parent Directory Manager.
//
func (dMgrHlpr *dirMgrHelper) getParentDirMgr(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (dirMgrOut DirMgr, hasParent bool, err error) {

  dirMgrOut = DirMgr{}
  hasParent = false
  err = nil
  var err2 error
  ePrefixCurrMethod := "dirMgrHelper.lowLevelDeleteDirectoryAll() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  _,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil && !dMgr.isInitialized {

    dirMgrOut = DirMgr{}
    hasParent = false
    return dirMgrOut, hasParent, err
  }

  err = nil

  if len(dMgr.parentPath) == 0 {

    dirMgrOut = dMgrHlpr.copyOut(dMgr)
    hasParent = false
    err = nil

    return dirMgrOut, hasParent, err

  } else {
    hasParent = true
  }

  dirMgrOut, err2 = DirMgr{}.New(dMgr.parentPath)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by DirMgr{}.New(%v.parentPath).\n"+
      "%v.parentPath=%v\nError='%v'\n",
      dMgrLabel,
      dMgrLabel,
      dMgr.parentPath,
      err2.Error())
    hasParent = true
    dirMgrOut = DirMgr{}
    return dirMgrOut, hasParent, err
  }

  err = nil

  return dirMgrOut, hasParent, err
}

// getValidPathStr - Performs validation on a path string.
// If the string contains a filename and file extension,
// this method will declare an error.
func (dMgrHlpr *dirMgrHelper) getValidPathStr(
  pathStr string,
  ePrefix string,
  pathStrLabel string) (validPathDto ValidPathStrDto, err error) {

  ePrefixCurrMethod := "dirMgrHelper.getValidPathStr() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  fh := FileHelper{}
  validPathDto = ValidPathStrDto{}.New()
  pathSepStr := string(os.PathSeparator)
  dotSeparator := "." + pathSepStr
  doubleDotSeparator := "." + dotSeparator
  doesPathExist := false
  fInfo := FileInfoPlus{}
  var volNameIndex, lSlashIdxs, lDotIdxs,
  strLen, firstCharIdx, lastCharIdx int
  var slashIdxs, dotIdxs []int
  var err2 error
  var volNameStr string

  pathStr,
    strLen,
    err =
    dMgrHlpr.lowLevelScreenPathStrForInvalidChars(
      pathStr,
      ePrefix,
      pathStrLabel)

  if err != nil {
    goto errorExit
  }

  validPathDto.originalPathStr = pathStr

  if strLen > 2 &&
  // Remove trailing slash
    pathStr[strLen-1] == os.PathSeparator &&
    pathStr[strLen-2] != '.' &&
    pathStr[strLen-2] != os.PathSeparator {
    pathStr = pathStr[0 : strLen-1]
    strLen--
  }

  volNameIndex,
    _,
    volNameStr = fh.GetVolumeNameIndex(pathStr)

  slashIdxs, err2 = fh.GetPathSeparatorIndexesInPathStr(pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.GetPathSeparatorIndexesInPathStr(%v).\n"+
      "%v='%v'\nError='%v'\n",
      pathStrLabel,
      pathStrLabel,
      pathStr,
      err2.Error())

    goto errorExit
  }

  lSlashIdxs = len(slashIdxs)

  firstCharIdx, lastCharIdx, err2 =
    fh.GetFirstLastNonSeparatorCharIndexInPathStr(pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr("+
      "%v).\n"+
      "%v='%v'\nError='%v'\n",
      pathStrLabel,
      pathStrLabel,
      pathStr,
      err2.Error())

    goto errorExit
  }

  dotIdxs, err2 = fh.GetDotSeparatorIndexesInPathStr(pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.GetDotSeparatorIndexesInPathStr(%v).\n"+
      "%v='%v'\nError='%v'\n",
      pathStrLabel,
      pathStrLabel,
      pathStr,
      err2.Error())

    goto errorExit
  }

  lDotIdxs = len(dotIdxs)

  // identify obvious valid path strings

  if pathStr == "." {
    validPathDto.pathStr = dotSeparator
    validPathDto.pathIsValid = PathValidStatus.Valid()
    goto successExit
  } else if pathStr == ".." {
    validPathDto.pathStr = doubleDotSeparator
    validPathDto.pathIsValid = PathValidStatus.Valid()
    goto successExit
  } else if pathStr == dotSeparator {
    validPathDto.pathStr = dotSeparator
    validPathDto.pathIsValid = PathValidStatus.Valid()
    goto successExit
  } else if pathStr == doubleDotSeparator {
    validPathDto.pathStr = doubleDotSeparator
    validPathDto.pathIsValid = PathValidStatus.Valid()
    goto successExit
  }

  if volNameIndex == 0 &&
    strings.ToLower(volNameStr) == strings.ToLower(pathStr) {

    if strings.Contains(strings.ToLower(runtime.GOOS), "windows") {
      pathStr += pathSepStr
    }

    validPathDto.pathStr = pathStr
    validPathDto.pathIsValid = PathValidStatus.Valid()
    goto successExit
  }

  // Check conversion to absolute path
  validPathDto.absPathStr,
    err2 =
    fh.MakeAbsolutePath(pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.MakeAbsolutePath("+
      "validPathDto.pathStr)\n"+
      "validPathDto.pathStr='%v'\n"+
      "Error='%v'\n",
      validPathDto.pathStr,
      err2.Error())

    goto errorExit
  }

  if lastCharIdx == -1 &&
    lDotIdxs == 0 &&
    lSlashIdxs == 0 {
    // No characters, no dots and no slashes
    err = fmt.Errorf(ePrefix+
      "\nError: %v is INVALID!\n"+
      "%v contains no valid characters in the string!\n"+
      "%v='%v'\n\n",
      pathStrLabel,
      pathStrLabel,
      pathStrLabel,
      pathStr)

    goto errorExit

  } else if lastCharIdx == -1 &&
    lDotIdxs > 0 &&
    lSlashIdxs == 0 {
    // No characters, no slashes, but Has Dots
    // Note: good dots have already been processed

    err = fmt.Errorf(ePrefix+
      "\nError: %v contains improperly formatted dot characters!\n"+
      "%v='%v'\n\n",
      pathStrLabel,
      pathStrLabel,
      pathStr)

    goto errorExit

  } else if lastCharIdx == -1 &&
    lDotIdxs > 0 &&
    lSlashIdxs > 0 {
    // No characters but Has slashes and
    // has dots.
    validPathDto.pathStr = pathStr
    validPathDto.pathIsValid = PathValidStatus.Valid()

  } else if lastCharIdx == -1 &&
    lDotIdxs == 0 &&
    lSlashIdxs > 0 {
    // No characters, No dots, but Has slashes
    err = fmt.Errorf(ePrefix+
      "\nError: '%v' contains improperly formatted path separator characters!\n"+
      "%v='%v'\n\n",
      pathStrLabel,
      pathStrLabel,
      pathStr)

    goto errorExit

  } else if lastCharIdx > -1 &&
    lDotIdxs == 0 &&
    lSlashIdxs == 0 {
    // Has characters, but No dots and No slashes
    validPathDto.pathStr = pathStr
    validPathDto.pathIsValid = PathValidStatus.Valid()

  } else if lastCharIdx > -1 &&
    lDotIdxs > 0 &&
    lSlashIdxs == 0 {
    // Has characters, Has Dots, but No slashes
    // Example  someFileName.txt

    if lDotIdxs > 1 {
      // To many dots
      err = fmt.Errorf(ePrefix+
        "\nError: %v contains improperly formatted path separator characters!\n"+
        "%v='%v'\n\n",
        pathStrLabel,
        pathStrLabel,
        pathStr)

      goto errorExit

    } else {
      // lDotIdx must equal '1'
      if dotIdxs[0] < firstCharIdx {
        // Example .git = directory
        validPathDto.pathStr = pathStr
        validPathDto.pathIsValid = PathValidStatus.Valid()
      } else {
        err = fmt.Errorf(ePrefix+
          "\nError: %v contains improperly formatted dot characters!\n"+
          "%v='%v'\n\n",
          pathStrLabel,
          pathStrLabel,
          pathStr)

        goto errorExit
      }
    }

  } else if lastCharIdx > -1 &&
    lDotIdxs > 0 &&
    lSlashIdxs > 0 {
    // Has characters, Has slashes, Has dots

    if firstCharIdx < slashIdxs[0] {
      // Example somefile/
      err = fmt.Errorf(ePrefix+
        "\nError: %v contains improperly "+
        "formatted characters and path separators!\n"+
        "%v='%v'\n\n",
        pathStrLabel,
        pathStrLabel,
        pathStr)

      goto errorExit

    } else if dotIdxs[lDotIdxs-1] >
      lastCharIdx {
      // Example somedir.
      err = fmt.Errorf(ePrefix+
        "\nError: %v contains improperly "+
        "formatted dot characters!\n"+
        "%v='%v'\n\n",
        pathStrLabel,
        pathStrLabel,
        pathStr)

      goto errorExit

    } else if dotIdxs[lDotIdxs-1]-slashIdxs[lSlashIdxs-1] == 1 &&
      lastCharIdx > dotIdxs[lDotIdxs-1] {
      // ../dir1/dir2/.git
      validPathDto.pathStr = pathStr
      validPathDto.pathIsValid = PathValidStatus.Valid()

    } else if lastCharIdx > dotIdxs[lDotIdxs-1] &&
      dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
      dotIdxs[lDotIdxs-1]-slashIdxs[lSlashIdxs-1] != 1 {
      // ./dir1/dir2/fileName.ext

      // Trim off trailing file name
      validPathDto.pathStr = pathStr[0:slashIdxs[lSlashIdxs-1]]

      if len(validPathDto.pathStr) == 0 {
        err = fmt.Errorf(ePrefix+
          "\nError: %v contains a "+
          "file name!\n"+
          "Attemp to trim trailing file name failed!\n"+
          "%v='%v'\n\n",
          pathStrLabel,
          pathStrLabel,
          pathStr)

        goto errorExit

      } else {
        validPathDto.pathIsValid = PathValidStatus.Valid()
      }

    } else if lastCharIdx > slashIdxs[lSlashIdxs-1] &&
      slashIdxs[lSlashIdxs-1] > dotIdxs[lDotIdxs-1] {
      // ../dir1/dir2/git
      validPathDto.pathStr = pathStr
      validPathDto.pathIsValid = PathValidStatus.Valid()

    } else {
      // unknown error
      err = fmt.Errorf(ePrefix+
        "\nError: '%v' contains a "+
        "file name!\n"+
        "%v='%v'\n\n",
        pathStrLabel,
        pathStrLabel,
        pathStr)

      goto errorExit
    }

  } else if lastCharIdx > -1 &&
    lDotIdxs == 0 &&
    lSlashIdxs > 0 {
    // Has characters, No Dots, Has slashes

    if slashIdxs[lSlashIdxs-1] > lastCharIdx {

      err = fmt.Errorf(ePrefix+
        "\nError: %v contains improperly "+
        "formatted path separators!\n"+
        "%v='%v'\n\n",
        pathStrLabel,
        pathStrLabel,
        pathStr)

      goto errorExit

    } else {
      validPathDto.pathStr = pathStr
      validPathDto.pathIsValid = PathValidStatus.Valid()
    }

  } else {

    err = fmt.Errorf(ePrefix+
      "\nError: %v is Invalid!\n"+
      "%v='%v'\n\n",
      pathStrLabel,
      pathStrLabel,
      pathStr)

    goto errorExit
  }

successExit:

  // Check conversion to absolute path
  validPathDto.absPathStr,
    err2 =
    fh.MakeAbsolutePath(validPathDto.pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.MakeAbsolutePath("+
      "validPathDto.pathStr)\n"+
      "validPathDto.pathStr='%v'\n"+
      "Error='%v'\n",
      validPathDto.pathStr,
      err2.Error())

    goto errorExit
  }

  validPathDto.pathVolumeIndex,
    validPathDto.pathVolumeStrLength,
    validPathDto.pathVolumeName =
    fh.GetVolumeNameIndex(validPathDto.absPathStr)

  doesPathExist,
    fInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    validPathDto.pathStr,
    ePrefix,
    pathStrLabel)

  if err != nil {
    goto errorExit
  }

  if doesPathExist {
    validPathDto.pathDoesExist = PathExistsStatus.Exists()
    validPathDto.pathFInfoPlus = fInfo.CopyOut()

  } else {
    // doesPathExist = false
    validPathDto.pathDoesExist = PathExistsStatus.DoesNotExist()
  }

  doesPathExist,
    fInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    validPathDto.absPathStr,
    ePrefix,
    pathStrLabel+".absolutePath")

  if err != nil {
    goto errorExit
  }

  if doesPathExist {
    validPathDto.absPathDoesExist = PathExistsStatus.Exists()
    validPathDto.absPathFInfoPlus = fInfo.CopyOut()
  } else {
    // doesPathExist = false
    validPathDto.absPathDoesExist = PathExistsStatus.DoesNotExist()
  }

  if validPathDto.pathDoesExist != validPathDto.absPathDoesExist {

    err = fmt.Errorf(ePrefix+
      "\nERROR: The path and absolute path show different values for "+
      "existence on disk.\n"+
      "validPathDto.pathDoesExist='%v'\n"+
      "validPathDto.absPathDoesExist='%v'\n",
      validPathDto.pathDoesExist.String(),
      validPathDto.absPathDoesExist.String())

    goto errorExit
  }

  if validPathDto.pathDoesExist == PathExistsStatus.Exists() &&
     validPathDto.absPathDoesExist == PathExistsStatus.Exists() {

    if !validPathDto.absPathFInfoPlus.IsDir() {
      err = fmt.Errorf(ePrefix+
        "\nERROR: The '%v' absolute path exists but it is classified "+
        "as a File, NOT a directory!\n"+
        "%v base path='%v'\n"+
        "%v absolute path='%v'\n",
        pathStrLabel,
        pathStrLabel,
        validPathDto.pathStr,
        pathStrLabel,
        validPathDto.absPathStr)

      goto errorExit
    }

    if validPathDto.absPathFInfoPlus.Mode().IsRegular() {
      err = fmt.Errorf(ePrefix+
        "\nERROR: The '%v' absolute path exists but it is classified "+
        "as a 'Regular' File, NOT a directory!\n"+
        "%v base path ='%v'\n"+
        "%v absolute path='%v'\n",
        pathStrLabel,
        pathStrLabel,
        validPathDto.pathStr,
        pathStrLabel,
        validPathDto.absPathStr)

      goto errorExit

    }

    if !validPathDto.pathFInfoPlus.IsDir() {
      err = fmt.Errorf(ePrefix+
        "\nERROR: The '%v' base path exists but it is classified "+
        "as a File, NOT a directory!\n"+
        "%v base path='%v'\n"+
        "%v absolute path='%v'\n",
        pathStrLabel,
        pathStrLabel,
        validPathDto.pathStr,
        pathStrLabel,
        validPathDto.absPathStr)

      goto errorExit
    }

    if validPathDto.pathFInfoPlus.Mode().IsRegular() {
      err = fmt.Errorf(ePrefix+
        "\nERROR: The '%v' base path exists but it is classified "+
        "as a 'Regular' File, NOT a directory!\n"+
        "%v base path='%v'\n"+
        "%v absolute path='%v'\n",
        pathStrLabel,
        pathStrLabel,
        validPathDto.pathStr,
        pathStrLabel,
        validPathDto.absPathStr)

      goto errorExit

    }

    err = nil
  }

errorExit:
  if err != nil {

    validPathDto.pathStr = pathStr
    validPathDto.absPathStr = ""
    validPathDto.pathStrLength =
      len(pathStr)
    validPathDto.pathIsValid = PathValidStatus.Invalid()
    validPathDto.isInitialized = true
    validPathDto.pathIsValid = PathValidStatus.Invalid()
    validPathDto.err = fmt.Errorf("%v", err.Error())
    return validPathDto, err
  }

  validPathDto.pathStrLength =
    len(validPathDto.pathStr)

  validPathDto.absPathStrLength =
    len(validPathDto.absPathStr)

  validPathDto.pathType = PathFileType.Path()
  validPathDto.pathIsValid = PathValidStatus.Valid()
  validPathDto.isInitialized = true

  err = validPathDto.IsDtoValid(ePrefix)

  return validPathDto, err
}

// isPathStringEmptyOrBlank - Determines whether a path string
// is blank.
//
func (dMgrHlpr *dirMgrHelper) isPathStringEmptyOrBlank(
  pathStr string,
  trimTrailingPathSeparator bool,
  ePrefix string,
  pathStrLabel string) (pathFileNameExt string, strLen int, err error) {

  ePrefixCurrMethod := "dirMgrHelper.isPathStringEmptyOrBlank() "

  pathFileNameExt = ""
  err = nil

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  strLen = len(pathStr)

  if strLen == 0 {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v is an empty string!",
      pathStrLabel)

    return pathFileNameExt, strLen, err
  }

  pathFileNameExt = strings.TrimLeft(pathStr, " ")

  pathFileNameExt = strings.TrimRight(pathFileNameExt, " ")

  strLen = len(pathFileNameExt)

  if strLen == 0 {

    err = fmt.Errorf(ePrefix+
      "\nERROR: %v consists entirely of blank spaces!",
      pathStrLabel)

    return pathFileNameExt, strLen, err
  }

  fh := FileHelper{}

  pathFileNameExt = fh.AdjustPathSlash(pathFileNameExt)

  dotPathSeparator := "." + string(os.PathSeparator)

  if strings.HasSuffix(pathFileNameExt, dotPathSeparator) {
    trimTrailingPathSeparator = false
  }

  strLen = len(pathFileNameExt)

  if trimTrailingPathSeparator &&
    pathFileNameExt[strLen-1] == os.PathSeparator {

    pathFileNameExt = pathFileNameExt[0 : strLen-1]
    strLen = len(pathFileNameExt)

  }

  return pathFileNameExt, strLen, err
}

// lowLevelCopyFile - This low level helper method is designed
// to copy files from a source file to a destination file.
//
// No validation or error checking is performed on the input
// parameters.
//
func (dMgrHlpr *dirMgrHelper) lowLevelCopyFile(
  src string,
  srcFInfo os.FileInfo,
  dst,
  ePrefix,
  srcLabel,
  dstLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelCopyFile() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if len(src) == 0 {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter %v is an empty string!\n", srcLabel)
  }

  if len(dst) == 0 {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter %v is an empty string!\n", dstLabel)
  }

  if !srcFInfo.Mode().IsRegular() {
    return fmt.Errorf(ePrefix+
      "Error: %v is a Non-Regular File and cannot be copied!\n"+
      "%v='%v'\n\n",
      srcLabel,
      srcLabel,
      src)
  }

  // First, open the source file
  inSrcPtr, err := os.Open(src)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned from os.Open(src)\n"+
      "%v='%v'\nError='%v'\n\n",
      srcLabel,
      src,
      err.Error())
  }
  // Next, 'Create' the destination file
  // If the destination file previously exists,
  // it will be truncated.
  outDestPtr, err := os.Create(dst)

  if err != nil {

    _ = inSrcPtr.Close()

    return fmt.Errorf(ePrefix+
      "Error returned from os.Create(destinationFile)\n"+
      "%='%v'\nError='%v'\n\n",
      dstLabel,
      dst,
      err.Error())
  }

  bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

  if err2 != nil {
    _ = inSrcPtr.Close()
    _ = outDestPtr.Close()
    err = fmt.Errorf(ePrefix+
      "Error returned from io.Copy(%v, %v) \n"+
      "%v='%v'\n"+
      "%v='%v'\nError='%v'\n\n",
      dstLabel,
      srcLabel,
      dstLabel,
      dst,
      srcLabel,
      src,
      err2.Error())

    return err
  }

  errs := make([]error, 0)

  // flush file buffers inSrcPtr memory
  err = outDestPtr.Sync()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned from outDestPtr.Sync()\n"+
      "%v='%v'\nError='%v'\n\n",
      dstLabel,
      dst,
      err.Error())

    errs = append(errs, err2)
  }

  err = inSrcPtr.Close()

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "Error returned from inSrcPtr.Close()\n"+
      "inSrcPtr=source='%v'\nError='%v'\n\n",
      src, err.Error())

    errs = append(errs, err2)
  }

  inSrcPtr = nil

  err = outDestPtr.Close()

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "Error returned from outDestPtr.Close()\n"+
      "outDestPtr=destination='%v'\nError='%v'\n\n",
      dst, err.Error())

    errs = append(errs, err2)
  }

  outDestPtr = nil

  if len(errs) > 0 {
    return FileHelper{}.ConsolidateErrors(errs)
  }
  var dstFileDoesExist bool
  var dstFileInfo FileInfoPlus

  dstFileDoesExist,
    dstFileInfo,
    err = dMgrHlpr.lowLevelDoesDirectoryExist(
    dst,
    ePrefix,
    dstLabel)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error: After Copy IO operation, %v "+
      "generated non-path error!\n"+
      "%v='%v'\nError='%v'\n\n",
      dstLabel,
      dstLabel,
      dst,
      err.Error())
  }

  if !dstFileDoesExist {
    err = fmt.Errorf(ePrefix+
      "ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
      "Destination File = '%v' = '%v'\n\n",
      dstLabel,
      dst)

    return err
  }

  srcFileSize := srcFInfo.Size()

  if bytesCopied != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "Error: Bytes Copied does NOT equal bytes "+
      "in source file!\n"+
      "Source File Bytes='%v'   Bytes Coped='%v'\n"+
      "Source File=%v='%v'\n"+
      "Destination File=%v='%v'\n\n",
      srcFileSize,
      bytesCopied,
      srcLabel,
      src,
      dstLabel,
      dst)

    return err
  }

  err = nil

  if dstFileInfo.Size() != srcFileSize {
    err = fmt.Errorf(ePrefix+
      "\nError: Bytes is source file do NOT equal bytes "+
      "in destination file!\n"+
      "Source File Bytes='%v'   Destination File Bytes='%v'\n"+
      "Source File=%v='%v'\n"+
      "Destination File=%v='%v'\n\n",
      srcFileSize,
      dstFileInfo.Size(),
      srcLabel,
      src,
      dstLabel,
      dst)
  }

  return err
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

  if dMgr == nil {
    return fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil' !", dMgrLabel)
  }

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
      dMgr.doesAbsolutePathExist = false
      dMgr.doesPathExist = false
      dMgr.actualDirFileInfo = FileInfoPlus{}
      return nil
    }

    time.Sleep(50 * time.Millisecond)
  }

  return err
}

// lowLevelDoesDirectoryExist - This helper method tests for the existence
// of directory path.
//
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

  if len(dirPath) == 0 {
    err = fmt.Errorf(ePrefix +
      "\nError: Input paramter %v is an empty string!\n", dirPathLabel)
    return dirPathDoesExist, fInfo, err
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
        "%v='%v'\nError='%v'\n\n",
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

func (dMgrHlpr *dirMgrHelper) lowLevelDirMgrFieldConfig(
  dMgr *DirMgr,
  validPathDto ValidPathStrDto,
  ePrefix string,
  dMgrLabel string) (isEmpty bool, err error) {

  isEmpty = false
  err = nil

  ePrefixCurrMethod := "dirMgrHelper.lowLevelDirMgrFieldConfig() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if dMgr == nil {
    err = fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil'!\n", dMgrLabel)
    return isEmpty, err
  }

  err = validPathDto.IsDtoValid(ePrefix)

  if err != nil {
    return isEmpty, err
  }

  dMgr.originalPath = validPathDto.originalPathStr

  dMgr.path = validPathDto.pathStr
  dMgr.isPathPopulated = true

  dMgr.absolutePath = validPathDto.absPathStr
  dMgr.isAbsolutePathPopulated = true

  fh := FileHelper{}
  var dirPathDoesExist, dirAbsPathDoesExist bool
  var pathFInfoPlus, absPathFInfoPlus FileInfoPlus

  if validPathDto.pathDoesExist != PathExistsStatus.Unknown() &&
    validPathDto.absPathDoesExist != PathExistsStatus.Unknown() {

    if validPathDto.pathDoesExist == PathExistsStatus.DoesNotExist() {
      dirPathDoesExist = false
    } else {
      dirPathDoesExist = true
      pathFInfoPlus = validPathDto.pathFInfoPlus.CopyOut()
    }

    if validPathDto.absPathDoesExist == PathExistsStatus.DoesNotExist() {
      dirAbsPathDoesExist = false
    } else {
      dirAbsPathDoesExist = true
      absPathFInfoPlus = validPathDto.absPathFInfoPlus.CopyOut()
    }

  } else {
    _,
      dirPathDoesExist,
      pathFInfoPlus,
      err =
      fh.doesPathFileExist(
        dMgr.path,
        PreProcPathCode.None(),
        ePrefix,
        dMgrLabel+".path")

    if err != nil {
      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      isEmpty = true
      return isEmpty, err
    }

    _,
      dirAbsPathDoesExist,
      absPathFInfoPlus,
      err =
      fh.doesPathFileExist(
        dMgr.absolutePath,
        PreProcPathCode.None(),
        ePrefix,
        dMgrLabel+".absolutePath")

    if err != nil {

      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      isEmpty = true
      return isEmpty, err
    }
  }

  if !dirPathDoesExist {
    dMgr.doesPathExist = false

  } else {

    if !pathFInfoPlus.IsDir() {
      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      err = fmt.Errorf(ePrefix+
        "\nERROR: Directory path exists, but it is a File - NOT a directory!\n"+
        "%v='%v'\n",
        dMgrLabel,
        dMgr.path)

      isEmpty = true
      return isEmpty, err
    }

    if pathFInfoPlus.Mode().IsRegular() {

      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      err = fmt.Errorf(ePrefix+
        "\nError: Directory path exists, but "+
        "it is classified as as a Regular File!\n"+
        "%v='%v'\n",
        dMgrLabel,
        dMgr.path)

      isEmpty = true
      return isEmpty, err
    }

    dMgr.doesPathExist = true
  }

  if dirAbsPathDoesExist {

    if !absPathFInfoPlus.IsDir() {
      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      err = fmt.Errorf(ePrefix+
        "\nThe Directory Manager absolute path exists and IS NOT A DIRECTORY!.\n"+
        "%v Path='%v'\n",
        dMgrLabel,
        dMgr.absolutePath)

      isEmpty = true
      return isEmpty, err
    }

    if absPathFInfoPlus.Mode().IsRegular() {

      _ = dMgrHlpr.empty(
        dMgr,
        ePrefix,
        dMgrLabel)

      err = fmt.Errorf(ePrefix+
        "\nError: Directory absolute path exists, but "+
        "it is classified as as a Regular File!\n"+
        "%v='%v'\n",
        dMgrLabel,
        dMgr.absolutePath)

      isEmpty = true
      return isEmpty, err
    }

    dMgr.doesAbsolutePathExist = true
    dMgr.actualDirFileInfo = absPathFInfoPlus.CopyOut()

  } else {
    dMgr.doesAbsolutePathExist = false
    dMgr.actualDirFileInfo = FileInfoPlus{}
  }

  strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
  lStr := len(strAry)
  idxStr := strAry[lStr-1]

  idx := strings.Index(dMgr.absolutePath, idxStr)

  dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

  dMgr.isParentPathPopulated = true

  if dMgr.parentPath == "" {
    dMgr.isParentPathPopulated = false
  }

  if idxStr != "" {
    dMgr.directoryName = idxStr
  } else {
    dMgr.directoryName = dMgr.absolutePath
  }

  if dMgr.path != dMgr.absolutePath {
    dMgr.isAbsolutePathDifferentFromPath = true
  }

  if validPathDto.pathVolumeName != "" {
    dMgr.isVolumePopulated = true
    dMgr.volumeName = validPathDto.pathVolumeName
  }

  if dMgr.isAbsolutePathPopulated && dMgr.isPathPopulated {
    dMgr.isInitialized = true
    isEmpty = false
  } else {
    isEmpty = true
  }

  err = nil

  return isEmpty, err
}

// lowLevelMakeDir - Helper Method used by 'DirMgr'. This method will create
// the directory path including parent directories for the path specified by
// 'dMgr'.
func (dMgrHlpr *dirMgrHelper) lowLevelMakeDir(
  dMgr *DirMgr,
  ePrefix string,
  dMgrLabel string) (dirCreated bool, err error) {

  ePrefixCurrMethod := "dirMgrHelper.lowLevelMakeDir() "
  dirCreated = false
  err = nil

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
    return dirCreated, err
  }

  if dMgrPathDoesExist {
    // The directory exists
    // Nothing to do.
    return dirCreated, err
  }

  fPermCfg, err2 := FilePermissionConfig{}.New("drwxrwxrwx")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err2.Error())

    return dirCreated, err
  }

  modePerm, err2 := fPermCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fPermCfg.GetCompositePermissionMode().\n"+
      "Error='%v\n", err2.Error())
    return dirCreated, err
  }

  err2 = os.MkdirAll(dMgr.absolutePath, modePerm)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by os.MkdirAll(%v.absolutePath, modePerm).\n"+
      "%v.absolutePath='%v'\nmodePerm=\"drwxrwxrwx\"\n"+
      "Error='%v'\n", err2.Error())

    return dirCreated, err
  }

  dMgrPathDoesExist,
    _,
    err2 =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err2 != nil {
    err = fmt.Errorf("Error: After attempted directory creation, "+
      "a non-path error was generated!\n"+
      "%v.absolutePath='%v'\n"+
      "Error='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      err2.Error())
    return dirCreated, err
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf("Error: After attempted directory creation,\n"+
      "the directory DOES NOT EXIST!\n"+
      "%v=%v\n", dMgrLabel, dMgr.absolutePath)

    return dirCreated, err
  }

  dirCreated = true
  err = nil

  return dirCreated, err
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
  dMgrLabel string) (dirCreated bool, err error) {

  dirCreated = false
  err = nil

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
    return dirCreated, err
  }

  if dMgrPathDoesExist {
    // The directory exists
    // Nothing to do.
    return dirCreated, err
  }

  err2 := fPermCfg.IsValid()

  if err2 != nil {
    err = fmt.Errorf("Input Parameter 'fPermCfg' is INVALID!\n"+
      "Error returned by fPermCfg.IsValid().\n"+
      "Error='%v'\n", err2.Error())

    return dirCreated, err
  }

  modePerm, err2 := fPermCfg.GetCompositePermissionMode()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fPermCfg.GetCompositePermissionMode().\n"+
      "Error='%v\n", err2.Error())

    return dirCreated, err
  }

  err2 = os.MkdirAll(dMgr.absolutePath, modePerm)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by os.MkdirAll(%v.absolutePath, modePerm).\n"+
      "%v.absolutePath='%v'\nmodePerm=\"drwxrwxrwx\"\n"+
      "Error='%v'\n",
      dMgrLabel, dMgr.absolutePath, err2.Error())

    return dirCreated, err
  }

  dMgrPathDoesExist,
    _,
    err2 =
    dMgrHlpr.doesDirectoryExist(
      dMgr,
      PreProcPathCode.None(),
      ePrefix,
      dMgrLabel)

  if err2 != nil {
    err = fmt.Errorf("Error: After attempted directory creation, "+
      "a non-path error was generated!\n"+
      "%v.absolutePath='%v'\n"+
      "Error='%v'\n",
      dMgrLabel,
      dMgr.absolutePath,
      err2.Error())
    return dirCreated, err
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf("Error: After attempted directory creation,\n"+
      "the directory DOES NOT EXIST!\n"+
      "%v=%v\n", dMgrLabel, dMgr.absolutePath)
    return dirCreated, err
  }

  dirCreated = true
  err = nil

  return dirCreated, err
}

// lowLevelScreenPathStrForInvalidChars - Examines input parameter 'pathStr'
// to determine if it contains invalid characters.
func (dMgrHlpr *dirMgrHelper) lowLevelScreenPathStrForInvalidChars(
  pathStr string,
  ePrefix string,
  pathStrLabel string) (validPathStr string, validPathStrLength int, err error) {

  validPathStr = ""
  validPathStrLength = 0
  err = nil

  ePrefixCurrMethod := "dirMgrHelper.lowLevelScreenPathStrForInvalidChars() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  strLen := 0

  pathStr,
    strLen,
    err = dMgrHlpr.isPathStringEmptyOrBlank(
    pathStr,
    true, // trim trailing path separator
    ePrefix,
    pathStrLabel)

  if err != nil {

    return validPathStr, validPathStrLength, err
  }

  tripleDotSeparator := "..."
  doublePathSeparator := string(os.PathSeparator) + string(os.PathSeparator)

  if strings.Contains(pathStr, tripleDotSeparator) {

    err = fmt.Errorf(ePrefix+
      "\nERROR: Input parameter '%v' contains invalid dot characters!\n"+
      "%v = %v\n",
      pathStrLabel,
      pathStrLabel,
      pathStr)

    return validPathStr, validPathStrLength, err

  }

  if strings.Contains(pathStr, doublePathSeparator) {

    err = fmt.Errorf(ePrefix+
      "\nERROR: Input parameter '%v' contains invalid path separator characters!\n"+
      "%v = %v\n",
      pathStrLabel,
      pathStrLabel,
      pathStr)

    return validPathStr, validPathStrLength, err
  }

  validPathStr = pathStr
  validPathStrLength = strLen
  err = nil

  return validPathStr, validPathStrLength, err
}

// moveDirectory - Moves files from the source directory identified
// by input parameter 'dMgr' to a target directory identified by input
// parameter 'targetDMgr'. The 'move' operation is accomplished
// in three steps. First, the files to be copied are selected according
// to file selection criteria specified by input parameter,'fileSelectCriteria'.
// Second, the selected files are copied to target directory identified
// by the input parameter, 'targetDMgr'. Finally, after verifying the copy,
// the files are deleted from the source directory ('dMgr').
//
// If, at the conclusion of the 'move' operation, there are no files or
// sub-directories remaining in the source directory (dMgr), the source
// directory will be delete.
//
// The selected files are copied using Copy IO operation. For information
// on the Copy IO procedure see FileHelper{}.CopyFileByIo() method and
// reference:
//   https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// If the target directory ('targetDMgr') does not previously exist, this method
// will attempt to create the target directory, provided, that files are selected
// for movement to that directory. If no files match the file selection criteria,
// the target directory will NOT be created.
//
// NOTE: This method ONLY moves files from the source directory identified by
// 'dMgr'. It does NOT move files from subdirectories.
//
// This method is optimized to support the movement of large numbers of files.
//
// ------------------------------------------------------------------------------
//
// IMPORTANT!!!!
// This method will delete files in the current DirMgr path!  If all files have
// been moved out of the directory and there are no sub-Directories remaining,
// the source directory, 'dMgr', will likewise be deleted.
//
func (dMgrHlpr *dirMgrHelper) moveDirectory(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string,
  fileSelectLabel string) (dirMoveStats DirectoryMoveStats, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.moveDirectory() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2 error
  var dMgrPathDoesExist, targetDMgrPathDoesExist bool

  if targetDMgr == nil {

    err = fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil'!\n", targetDMgrLabel)

    errs = append(errs, err)

    return dirMoveStats, errs

  }

  dMgrPathDoesExist,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {

    errs = append(errs, err)

    return dirMoveStats, errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    errs = append(errs, err)

    return dirMoveStats, errs
  }

  targetDMgrPathDoesExist,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    targetDMgr,
    PreProcPathCode.None(),
    ePrefix,
    targetDMgrLabel)

  if err != nil {

    errs = append(errs, err)

    return dirMoveStats, errs
  }

  fh := FileHelper{}

  dir, err := os.Open(dMgr.absolutePath)

  if err != nil {

    err2 = fmt.Errorf(ePrefix+
      "\nError return by os.Open(%v.absolutePath).\n"+
      "%v.absolutePath='%v'\nError='%v'\n",
      dMgrLabel,
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())

    errs = append(errs, err2)

    return dirMoveStats, errs
  }

  osPathSeparatorStr := string(os.PathSeparator)
  var src, target string
  var isMatch, dirCreated bool
  var nameFileInfos []os.FileInfo

  file2LoopIsDone := false

  for !file2LoopIsDone {

    nameFileInfos, err = dir.Readdir(1000)

    if err != nil && err == io.EOF {
      file2LoopIsDone = true

      if len(nameFileInfos) == 0 {
        break
      }

    } else if err != nil {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dir.Readdirnames(1000).\n"+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel,
        dMgr.absolutePath,
        err.Error())
      errs = append(errs, err2)
      file2LoopIsDone = true
      break
    }

    for _, nameFInfo := range nameFileInfos {

      if nameFInfo.IsDir() {
        dirMoveStats.NumOfSubDirectories++
        continue

      }

      // This is not a directory. It is a file.
      // Determine if it matches the find file criteria.
      dirMoveStats.TotalSrcFilesProcessed++

      isMatch, err =
        fh.FilterFileName(nameFInfo, fileSelectCriteria)

      if err != nil {

        err2 =
          fmt.Errorf(ePrefix+
            "\nError returned by fh.FilterFileName(nameFInfo, %v). "+
            "%v Directory Searched='%v'\nfileName='%v'\nError='%v'\n\n",
            fileSelectLabel,
            dMgrLabel,
            dMgr.absolutePath,
            nameFInfo.Name(),
            err.Error())

        errs = append(errs, err2)

        continue
      }

      if !isMatch {
        dirMoveStats.SourceFilesRemaining++
        continue

      } else {
        // We have a match
        dirCreated = false
        // Create Directory if needed
        if !targetDMgrPathDoesExist {

          dirCreated,
            err = dMgrHlpr.lowLevelMakeDir(
            targetDMgr,
            ePrefix,
            targetDMgrLabel)

          if err != nil {
            err2 = fmt.Errorf(ePrefix+
              "\nError creating target directory!\n"+
              "%v Directory='%v'\nError='%v'\n\n",
              targetDMgrLabel,
              targetDMgr.absolutePath,
              err.Error())

            errs = append(errs, err2)
            file2LoopIsDone = true
            break
          }

          if dirCreated {
            dirMoveStats.DirsCreated++
          }

          dirMoveStats.DirsCreated++
          targetDMgrPathDoesExist = true
        }

        src = dMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        target = targetDMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        err = dMgrHlpr.lowLevelCopyFile(
          src,
          nameFInfo,
          target,
          ePrefix,
          "sourceFile",
          "destinationFile")

        if err != nil {
          errs = append(errs, err)
          dirMoveStats.SourceFilesRemaining++
          continue

        }

        err = os.Remove(src)

        if err != nil {
          err2 = fmt.Errorf(ePrefix+
            "\nError occurred after file copy completed during delete operation!\n"+
            "\nError returned by os.Remove(sourceFile)\n"+
            "sourceFile='%v'\nError='%v'\n\n",
            src, err.Error())

          errs = append(errs, err)
          dirMoveStats.SourceFilesRemaining++
          continue
        }

        dirMoveStats.SourceFilesMoved++
      }
    }
  }

  if dir != nil {
    err = dir.Close()

    if err != nil {
      err2 = fmt.Errorf(ePrefix+
        "Error returned by dir.Close(). "+
        "dir='%v' Error='%v' ",
        dMgr.absolutePath, err.Error())

      errs = append(errs, err2)
    }
  }

  if dirMoveStats.TotalSrcFilesProcessed !=
    dirMoveStats.SourceFilesMoved+dirMoveStats.SourceFilesRemaining {

    err = fmt.Errorf(ePrefix+
      "Counting Error: Total Number of Files processed is NOT EQUAL to\n"+
      "the number of source moved plus the number of source files remaining.\n"+
      "Source Directory= %v.absolutePath='%v'\n"+
      "Total Source Files in %v Directory='%v'\n"+
      "Number of source files moved='%v'\n"+
      "Number of source files remaining='%v'\n\n",
      dMgrLabel,
      dMgr.absolutePath,
      dirMoveStats.TotalSrcFilesProcessed,
      dirMoveStats.SourceFilesMoved,
      dirMoveStats.SourceFilesRemaining)

    errs = append(errs, err)
  }

  // If all the source files have been moved and
  // there are no sub-directories, DELETE the
  // directory (dMgr).
  if dirMoveStats.SourceFilesRemaining == 0 &&
    dirMoveStats.NumOfSubDirectories == 0 {

    err = dMgrHlpr.lowLevelDeleteDirectoryAll(
      dMgr,
      ePrefix,
      dMgrLabel)

    if err != nil {
      errs = append(errs, err)
      dirMoveStats.SourceDirWasDeleted = false
    } else {
      dirMoveStats.SourceDirWasDeleted = true
      dMgr.doesAbsolutePathExist = false
      dMgr.doesPathExist = false
      dMgr.actualDirFileInfo = FileInfoPlus{}
    }
  }

  return dirMoveStats, errs
}

// moveDirectoryTree - Moves all sub-directories and files plus files in
// the parent 'dMgr' directory to a target directory tree specified by
// input parameter 'targetDMgr'. If successful, the parent directory,
// 'dMgr, will be deleted along with the entire sub-directory tree.
//
// --------------------------------------------------------------------
//
// !!!! BE CAREFUL !!!! This method will delete the entire directory
// tree identified by 'dMgr' along with ALL the files in that
// directory tree!
//
// --------------------------------------------------------------------
//
func (dMgrHlpr *dirMgrHelper) moveDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (
  dirMoveStats DirectoryMoveStats, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.moveDirectoryTree() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2 error

  fileSelectCriteria := FileSelectionCriteria{}

  dTreeCopyStats,
  errs2 :=
    dMgrHlpr.copyDirectoryTree(
      dMgr,
      targetDMgr,
      true,
      false,
      fileSelectCriteria,
      ePrefix,
      "dMgr",
      "targetDMgr")

  if len(errs2) > 0 {
    err2 = fmt.Errorf(ePrefix+
      "\nErrors occurred while copying directory tree to target directory.\n"+
      "The source directory WAS NOT DELETED!\n"+
      "%v Source Directory='%v'\n%v Target Directory='%v'\nErrors Follow:\n\n",
      dMgrLabel,
      dMgr.absolutePath,
      targetDMgrLabel,
      targetDMgr.absolutePath)
    errs = append(errs, err2)
    errs = append(errs, errs2...)

    return dirMoveStats, errs
  }

  dirMoveStats.TotalDirsProcessed =
    dTreeCopyStats.TotalDirsScanned

  dirMoveStats.DirsCreated =
    dTreeCopyStats.DirsCreated

  dirMoveStats.NumOfSubDirectories =
    dTreeCopyStats.TotalDirsScanned - 1

  dirMoveStats.SourceFilesRemaining =
    dTreeCopyStats.FilesNotCopied

  dirMoveStats.SourceFileBytesRemaining =
    dTreeCopyStats.FileBytesNotCopied

  if dirMoveStats.SourceFilesRemaining > 0 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Some of the files designated to be moved to the target directory, were NOT copied!\n"+
      "Therefore the source directory WILL NOT BE DELETED!\n"+
      "Number of Files NOT Copied='%v'\n",
      "%v Source Directory='%v'\n%v Target Directory='%v'\n\n",
      dTreeCopyStats.FilesNotCopied,
      dMgrLabel, dMgr.absolutePath,
      targetDMgrLabel, targetDMgr.absolutePath)
    errs = append(errs, err2)

    return dirMoveStats, errs
  }

  dirMoveStats.TotalSrcFilesProcessed =
    dTreeCopyStats.TotalFilesProcessed

  err = dMgrHlpr.lowLevelDeleteDirectoryAll(
    dMgr,
    ePrefix,
    dMgrLabel)

  if err != nil {
    err2 = fmt.Errorf(ePrefix+
      "\nFiles were copied successfuly to target directory.\n"+
      "However, errors occurred while deleting the source directory tree.\n"+
      "%v.absolutePath='%v'\nError='%v'\n\n",
      dMgrLabel,
      dMgr.absolutePath,
      err.Error())

    errs = append(errs, err2)
  }

  dirMoveStats.SourceDirWasDeleted = true
  dirMoveStats.SourceFilesMoved =
    dTreeCopyStats.FilesCopied
  dirMoveStats.SourceFileBytesMoved =
    dTreeCopyStats.FileBytesCopied

  return dirMoveStats, errs
}

// moveSubDirectoryTree - Moves all subdirectories in the 'dMgr'
// tree to the 'targetDMgr' subdirectory tree.
//
func (dMgrHlpr *dirMgrHelper) moveSubDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (
  dirMoveStats DirectoryMoveStats, errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.moveSubDirectoryTree() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err2 error

  fileSelectCriteria := FileSelectionCriteria{}

  dTreeCopyStats,
  errs2 :=
    dMgrHlpr.copyDirectoryTree(
      dMgr,
      targetDMgr,
      true, // copy empty directories
      true, // skip top level directory
      fileSelectCriteria,
      ePrefix,
      "dMgr",
      "targetDMgr")

  if len(errs2) > 0 {
    err2 = fmt.Errorf(ePrefix+
      "\nErrors occurred while copying directory tree to target directory.\n"+
      "The source directory WAS NOT DELETED!\n"+
      "%v Source Directory='%v'\n%v Target Directory='%v'\nErrors Follow:\n\n",
      dMgrLabel,
      dMgr.absolutePath,
      targetDMgrLabel,
      targetDMgr.absolutePath)
    errs = append(errs, err2)
    errs = append(errs, errs2...)

    return dirMoveStats, errs
  }

  dirMoveStats.TotalDirsProcessed =
    dTreeCopyStats.TotalDirsScanned

  dirMoveStats.DirsCreated =
    dTreeCopyStats.DirsCreated

  dirMoveStats.NumOfSubDirectories =
    dTreeCopyStats.TotalDirsScanned

  dirMoveStats.SourceFilesRemaining =
    dTreeCopyStats.FilesNotCopied

  dirMoveStats.SourceFileBytesRemaining =
    dTreeCopyStats.FileBytesNotCopied

  if dirMoveStats.SourceFilesRemaining > 0 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Some of the files designated to be moved to the target directory, were NOT copied!\n"+
      "Therefore the source directory WILL NOT BE DELETED!\n"+
      "Number of Files NOT Copied='%v'\n",
      "%v Source Directory='%v'\n%v Target Directory='%v'\n\n",
      dTreeCopyStats.FilesNotCopied,
      dMgrLabel, dMgr.absolutePath,
      targetDMgrLabel, targetDMgr.absolutePath)

    errs = append(errs, err2)

    return dirMoveStats, errs
  }

  dirMoveStats.TotalSrcFilesProcessed =
    dTreeCopyStats.TotalFilesProcessed

  errs2 = dMgrHlpr.deleteAllSubDirectories(
    dMgr,
    ePrefix,
    "dMgr")

  if len(errs2) > 0 {
    errs = append(errs, errs2...)
    return dirMoveStats, errs
  }

  dirMoveStats.SourceDirWasDeleted = true

  dirMoveStats.SourceFilesMoved =
    dTreeCopyStats.FilesCopied

  dirMoveStats.SourceFileBytesMoved =
    dTreeCopyStats.FileBytesCopied

  return dirMoveStats, errs
}

// setDirMgr - Sets internal values for DirMgr instance based on
// a path or path/file name string passed as an input parameter
//
func (dMgrHlpr *dirMgrHelper) setDirMgr(
  dMgr *DirMgr,
  pathStr string,
  ePrefix string,
  dMgrLabel string,
  pathStrLabel string) (isEmpty bool, err error) {

  err = nil
  isEmpty = true

  ePrefixCurrMethod := "dirMgrHelper.setDirMgr() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if dMgr == nil {

    err = fmt.Errorf(ePrefix +
      "\nInput parameter %v pointer is 'nil'!\n", dMgrLabel)

    return isEmpty, err
  }

  validPathDto := ValidPathStrDto{}.New()

  validPathDto,
    err =
    dMgrHlpr.getValidPathStr(
      pathStr,
      ePrefix,
      "pathStr")

  if err != nil {
    isEmpty = true
    return isEmpty, err
  }

  err = dMgrHlpr.empty(
    dMgr,
    ePrefix,
    dMgrLabel)

  if err != nil {

    return isEmpty, err
  }

  return dMgrHlpr.lowLevelDirMgrFieldConfig(
    dMgr,
    validPathDto,
    ePrefix,
    dMgrLabel)
}

// setDirMgrFromKnownPathDirName - Configures the internal
// field values for the 'dMgr' instance using a parent path
// name and a directory name. The parent path and directory
// name are combined to form the full path for the 'dMgr'
// instance.
//
// This method will replace all previous field values with new
// values based on input parameters 'parentPathName' and
// 'directoryName'.
//
// This method differs from other "Set" methods in that it
// assumes the input parameters are known values and do not
// require the usual analysis and validation screening applied
// by similar methods.
//
func (dMgrHlpr *dirMgrHelper) setDirMgrFromKnownPathDirName(
  dMgr *DirMgr,
  pathStr string,
  dirName string,
  ePrefix string,
  dMgrLabel string,
  pathStrLabel string,
  dirNameLabel string) (isEmpty bool, err error) {

  err = nil
  isEmpty = true

  ePrefixCurrMethod := "dirMgrHelper.setDirMgrFromKnownPathDirName() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if dMgr == nil {

    err = fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil' !\n", dirNameLabel)

    return isEmpty, err
  }

  strLen := 0

  pathStr,
    strLen,
    err = dMgrHlpr.lowLevelScreenPathStrForInvalidChars(
    pathStr,
    ePrefix,
    pathStrLabel)

  if err != nil {

    isEmpty = true
    return isEmpty, err
  }

  lDirName := 0

  dirName,
    lDirName,
    err = dMgrHlpr.lowLevelScreenPathStrForInvalidChars(
    dirName,
    ePrefix,
    dirNameLabel)

  if err != nil {

    isEmpty = true
    return isEmpty, err
  }

  fh := FileHelper{}

  dotSeparator := "." + string(os.PathSeparator)
  doubleDotSeparator := ".." + string(os.PathSeparator)

  if lDirName > 2 &&
    strings.HasPrefix(dirName, doubleDotSeparator) {

    dirName = dirName[3:]

  } else if lDirName > 1 &&
    (strings.HasPrefix(dirName, dotSeparator) ||
      strings.HasPrefix(dirName, "..")) {

    dirName = dirName[2:]

  } else if lDirName > 0 &&
    (dirName[0] == os.PathSeparator ||
      dirName[0] == '.') {

    dirName = dirName[1:]

  }

  finalPathStr := ""

  if pathStr[strLen-1] != os.PathSeparator {
    finalPathStr = pathStr + string(os.PathSeparator) + dirName
  } else {
    finalPathStr = pathStr + dirName
  }

  var err2 error

  validPathDto := ValidPathStrDto{}.New()

  validPathDto.originalPathStr = finalPathStr
  validPathDto.pathStr = finalPathStr

  validPathDto.absPathStr,
    err2 = fh.MakeAbsolutePath(validPathDto.pathStr)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "\nError returned by fh.MakeAbsolutePath(pathStr).\n"+
      "Directory Path='%v'\nError='%v'\n",
      validPathDto.pathStr,
      err2.Error())

    isEmpty = true
    return isEmpty, err
  }

  validPathDto.pathVolumeIndex,
    validPathDto.pathVolumeStrLength,
    validPathDto.pathVolumeName =
    fh.GetVolumeNameIndex(validPathDto.absPathStr)

  validPathDto.pathStrLength = len(pathStr)
  validPathDto.absPathStrLength = len(validPathDto.absPathStr)
  validPathDto.pathDoesExist = PathExistsStatus.Unknown()
  validPathDto.absPathDoesExist = PathExistsStatus.Unknown()
  validPathDto.isInitialized = true
  validPathDto.pathIsValid = PathValidStatus.Valid()

  err = validPathDto.IsDtoValid(ePrefix)

  if err != nil {

    isEmpty = true
    return isEmpty, err
  }

  err = dMgrHlpr.empty(
    dMgr,
    ePrefix,
    dMgrLabel)

  if err != nil {

    return isEmpty, err
  }

  isEmpty,
    err = dMgrHlpr.lowLevelDirMgrFieldConfig(
    dMgr,
    validPathDto,
    ePrefix,
    dMgrLabel)

  return isEmpty, err
}

// setDirMgrWithPathDirectoryName - Configures a Directory Manager
// instance based on 'path' and 'directory name' parameters.
//
func (dMgrHlpr *dirMgrHelper) setDirMgrWithPathDirectoryName(
  dMgr *DirMgr,
  pathStr string,
  directoryName string,
  ePrefix string,
  dMgrLabel string,
  pathStrLabel string,
  directoryNameLabel string) (isEmpty bool, err error) {

  isEmpty = false
  err = nil

  ePrefixCurrMethod := "dirMgrHelper.setDirMgrWithPathDirectoryName() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  if dMgr == nil {

    err = fmt.Errorf(ePrefix +
      "\nError: Input parameter %v pointer is 'nil' !\n", dMgrLabel)

    return isEmpty, err
  }

  strLen := 0

  pathStr,
    strLen,
    err = dMgrHlpr.isPathStringEmptyOrBlank(
    pathStr,
    true,
    ePrefix,
    pathStrLabel)

  if err != nil {
    isEmpty = true
    return isEmpty, err
  }

  directoryName,
    _,
    err = dMgrHlpr.isPathStringEmptyOrBlank(
    directoryName,
    true,
    ePrefix,
    directoryNameLabel)

  if err != nil {
    isEmpty = true
    return isEmpty, err
  }

  if directoryName[0] == os.PathSeparator {
    directoryName = directoryName[1:]
  }

  finalPathStr := ""

  if pathStr[strLen-1] != os.PathSeparator {
    finalPathStr =
      pathStr + string(os.PathSeparator) + directoryName

  } else {
    finalPathStr = pathStr + directoryName
  }

  validPathDto := ValidPathStrDto{}.New()

  validPathDto,
    err =
    dMgrHlpr.getValidPathStr(
      finalPathStr,
      ePrefix,
      "pathStr")

  if err != nil {
    isEmpty = true
    return isEmpty, err
  }

  err = dMgrHlpr.empty(
    dMgr,
    ePrefix,
    dMgrLabel)

  if err != nil {

    return isEmpty, err
  }

  isEmpty,
    err = dMgrHlpr.lowLevelDirMgrFieldConfig(
    dMgr,
    validPathDto,
    ePrefix,
    dMgrLabel)

  return isEmpty, err
}

// setPermissions - Sets the read/write and execute
// permissions for the directory identified by the
// 'dMgr' instance. Note the treatment of 'execute'
// permissions may vary by operating system.
//
func (dMgrHlpr *dirMgrHelper) setPermissions(
  dMgr *DirMgr,
  permissionConfig FilePermissionConfig,
  ePrefix string,
  dMgrLabel string,
  permissionConfigLabel string) error {

  ePrefixCurrMethod := "dirMgrHelper.setPermissions() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  err := permissionConfig.IsValid()

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nInput parameter '%v' is INVALID!\n"+
      "Error returned by %v.IsValid()\n"+
      "%v='%v'\n"+
      "Error='%v'\n\n",
      permissionConfigLabel,
      permissionConfigLabel,
      permissionConfigLabel,
      permissionConfig.GetPermissionNarrativeText(),
      err.Error())
  }

  dirPathDoesExist,
  _,
  err := dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    return err
  }

  if !dirPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    return err
  }

  err = FileHelper{}.ChangeFileMode(
    dMgr.absolutePath,
    permissionConfig)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "\nError retrned by FileHelper{}.ChangeFileMode("+
      "%v.absolutePath, %v)\n"+
      "%v.absolutePath=%v\n"+
      "%v='%v'"+
      "Error='%v'\n\n",
      dMgrLabel,
      permissionConfigLabel,
      dMgrLabel,
      dMgr.absolutePath,
      permissionConfigLabel,
      permissionConfig.GetPermissionNarrativeText(),
      err.Error())
  }

  return nil
}

// substituteBaseDir - Substitute 'baseDir' segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
//
func (dMgrHlpr *dirMgrHelper) substituteBaseDir(
  dMgr *DirMgr,
  baseDir *DirMgr,
  substituteBaseDir *DirMgr,
  ePrefix string,
  dMgrLabel string,
  baseDirLabel string,
  substituteBaseDirLabel string) (newDMgr DirMgr, err error) {

  ePrefixCurrMethod := "dirMgrHelper.substituteBaseDir() "

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  newDMgr = DirMgr{}
  err = nil

  _,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    dMgr,
    PreProcPathCode.None(),
    ePrefix,
    dMgrLabel)

  if err != nil {
    return newDMgr, err
  }

  _,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    baseDir,
    PreProcPathCode.None(),
    ePrefix,
    baseDirLabel)

  if err != nil {
    return newDMgr, err
  }

  _,
    _,
    err = dMgrHlpr.doesDirectoryExist(
    substituteBaseDir,
    PreProcPathCode.None(),
    ePrefix,
    substituteBaseDirLabel)

  if err != nil {
    return newDMgr, err
  }

  thisDirAbsPath := strings.ToLower(dMgr.absolutePath)

  oldBaseAbsPath := strings.ToLower(baseDir.absolutePath)

  newBaseAbsPath := strings.ToLower(substituteBaseDir.absolutePath)

  idx := strings.Index(thisDirAbsPath, oldBaseAbsPath)

  if idx < 0 {
    err = fmt.Errorf(ePrefix+
      "\nThe base directory was NOT found in the current %v path!\n"+
      "%v Path='%v'\n%v Path='%v'\n\n",
      dMgrLabel,
      dMgrLabel,
      thisDirAbsPath,
      baseDirLabel,
      oldBaseAbsPath)

    return newDMgr, err
  }

  if idx != 0 {
    err = fmt.Errorf(ePrefix+
      "\nThe %v directory was NOT found at the beginning of the %v path!\n"+
      "%v Path='%v'\n%v Path='%v'\n\n",
      baseDirLabel,
      dMgrLabel,
      dMgrLabel,
      thisDirAbsPath,
      baseDirLabel,
      oldBaseAbsPath)

    return newDMgr, err
  }

  oldBaseLen := len(oldBaseAbsPath)

  newAbsPath := newBaseAbsPath + thisDirAbsPath[oldBaseLen:]

  isEmpty := false

  isEmpty, err = dMgrHlpr.setDirMgr(
    &newDMgr,
    newAbsPath,
    ePrefix,
    dMgrLabel,
    "newAbsPath")

  if err != nil {

    _ = dMgrHlpr.empty(
      &newDMgr,
      ePrefix,
      dMgrLabel)

    return newDMgr, err
  }

  if isEmpty {

    _ = dMgrHlpr.empty(
      &newDMgr,
      ePrefix,
      dMgrLabel)

    err = fmt.Errorf(ePrefix+
      "\nERROR: New generated Directory Path Is Invalid!\n"+
      "isEmpty='true'\n"+
      "newAbsPath='%v'\n\n", newAbsPath)

    return newDMgr, err
  }

  err = nil
  return newDMgr, err
}
