package pathfileops

import (
  "fmt"
  "io"
  "os"
  pf "path/filepath"
  "strings"
  "time"
)

type dirMgrHelper struct {
  dMgr DirMgr
}

// consolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (dMgrHlpr *dirMgrHelper) consolidateErrors(errs []error) error {

  lErrs := len(errs)

  if lErrs == 0 {
    return nil
  }

  errsStr := ""
  tempStr := ""

  for i := 0; i < lErrs; i++ {

    tempStr = fmt.Sprintf("%v", errs[i].Error())

    tempStr = strings.TrimLeft(strings.TrimRight(tempStr, " "), " ")

    if strings.HasSuffix(tempStr, "\n\n") {

    } else if strings.HasSuffix(tempStr, "\n") {

      tempStr += "\n"

    } else {
      tempStr += "\n\n"
    }

    errsStr += fmt.Sprintf("%v", tempStr)

  }

  return fmt.Errorf("%v", errsStr)
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
  targetDMgrLabel string) (dirCopyStats DirectoryCopyStats,
  errs []error) {

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

        err = fh.CopyFileByIo(src, target)

        if err != nil {
          err2 = fmt.Errorf("\n"+ePrefix+
            "\nERROR: fh.CopyFileByIo(src, target)\n"+
            "src='%v'\ntarget='%v'\nError='%v'\n\n",
            src, target, err.Error())

          errs = append(errs, err2)
        } else {
          dirCopyStats.FilesCopied++
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

func (dMgrHlpr *dirMgrHelper) copyDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  copyEmptyDirectories bool,
  skipTopLevelDirectory bool,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) (dTreeCopyStats DirTreeCopyStats,
  errs []error) {

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

    return dTreeCopyStats,
      errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: %v directory path DOES NOT EXIST!\n",
      dMgrLabel)

    errs = append(errs, err)

    return dTreeCopyStats,
      errs
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

    return dTreeCopyStats,
      errs
  }

  fh := FileHelper{}

  baseDirLen := len(dMgr.absolutePath)
  osPathSepStr := string(os.PathSeparator)
  var nameFileInfos []os.FileInfo
  dirs := DirMgrCollection{}
  var dirPtr *os.File
  var nextTargetDMgr DirMgr
  var isMatch, isTopLevelDir, isNewDir, dirCreated bool
  var srcFile, targetFile string

  nextDir := dMgrHlpr.copyOut(dMgr)

  for nextDir.isInitialized {

    dTreeCopyStats.TotalDirsProcessed++

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

    if dTreeCopyStats.TotalDirsProcessed == 1 {
      isTopLevelDir = true
    } else {
      isTopLevelDir = false
    }

    dirCreated = false

    if isTopLevelDir &&
      !skipTopLevelDirectory &&
      copyEmptyDirectories {

      dTreeCopyStats.DirsCopied++

      dirCreated,
        err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "1-nextTargetDMgr")

    } else if !isTopLevelDir && copyEmptyDirectories {

      dTreeCopyStats.DirsCopied++

      dirCreated,
        err = dMgrHlpr.lowLevelMakeDir(
        &nextTargetDMgr,
        ePrefix,
        "2-nextTargetDMgr")

    } else {
      err = nil
    }

    if err != nil {

      dTreeCopyStats.DirsCopied--

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
    } else if dirCreated {
      dTreeCopyStats.DirsCreated++
      dirCreated = false
    }

    err3 = nil

    isNewDir = true

    for err3 != io.EOF {

      nameFileInfos, err3 = dirPtr.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "\nError returned by dirPtr.Readdirnames(1000).\n"+
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
          dTreeCopyStats.TotalFilesProcessed++

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

            dTreeCopyStats.FilesNotCopied++
            continue

          } else {

            // We have a match
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
              errs = append(errs, err)
              err3 = io.EOF
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
                err3 = io.EOF
                break
              } else if dirCreated {
                dTreeCopyStats.DirsCreated++
              }
            }

            if isNewDir && !copyEmptyDirectories {

              dTreeCopyStats.DirsCopied++
            }

            isNewDir = false

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
            } else {
              dTreeCopyStats.FilesCopied++
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

      err = io.EOF
      break

    } else if err != nil {

      break
    }

  } // end of for nextDir.isInitialized

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

    err2 = fmt.Errorf("ERROR: The copy operation failed to create\n"+
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

  return dTreeCopyStats,
    errs
}

/*
func (dMgrHlpr *dirMgrHelper) copyDirectoryTree(
  dMgr *DirMgr,
  targetDMgr *DirMgr,
  copyEmptyDirectories bool,
  skipTopLevelDirectory bool,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  targetDMgrLabel string) ( dTreeCopyStats DirTreeCopyStats,
                            errs []error) {

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

    return dTreeCopyStats,
      errs
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nError: %v directory path DOES NOT EXIST!\n",
      dMgrLabel)

    errs = append(errs, err)

    return dTreeCopyStats,
      errs
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

    return dTreeCopyStats,
      errs
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

  nextDir := dMgrHlpr.copyOut(dMgr)

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

    if baseDirLen == len(nextDir.absolutePath) {
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

    dTreeCopyStats.TotalDirsProcessed++

    for err3 != io.EOF {

      nameFileInfos, err3 = dirPtr.Readdir(1000)

      if err3 != nil && err3 != io.EOF {

        err2 = fmt.Errorf("\n"+ePrefix+
          "\nError returned by dirPtr.Readdirnames(1000).\n"+
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
          dTreeCopyStats.TotalFilesProcessed++

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

            dTreeCopyStats.FilesNotCopied++
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
            } else {
              dTreeCopyStats.FilesCopied++
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

      err = io.EOF
      break

    } else if err != nil {

      break
    }

  } // end of for nextDir.isInitialized

  if dTreeCopyStats.FilesCopied > 0 {

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

    } else if !dMgrPathDoesExist {
      err2 = fmt.Errorf("ERROR: The copy operation failed to create\n"+
        "the 'targetDMgr' path. 'targetDMgr' path DOES NOT EXIST!\n"+
        "targetDMgr Path='%v'\n\n",
        targetDMgr.absolutePath)
      errs = append(errs, err2)
    }
  }

  if dTreeCopyStats.TotalFilesProcessed !=
    dTreeCopyStats.FilesCopied + dTreeCopyStats.FilesNotCopied {

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

  return dTreeCopyStats,
    errs
}
*/

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
        "\nError returned by dMgrPtr.Readdirnames(1000).\n"+
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
          "\nError returned by dirPtr.Readdirnames(1000).\n"+
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
        "\nError returned by dirPtr.Readdirnames(1000).\n"+
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
    dMgr.isInitialized = false
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
    dMgr.isInitialized = false
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

  return fileMgrCol, dMgrHlpr.consolidateErrors(errs)
}

// findFilesBySelectCriteria - This helper method is designed to conduct
// a file search in the directory identified by the input parameter, 'dMgr.
// The file search is limited to that directory ONLY. No sub-directories
// will be searched.
//
// Files matching the "fileSelectCriteria" input parameter will be used
// to screen available files. Any files matching the file selection
// criteria will be returned in a 'FileMgrCollection'.
//
// Only matched files will be returned. No sub-directory names will ever
// be included.
//
func (dMgrHlpr *dirMgrHelper) findFilesBySelectCriteria(
  dMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string,
  fileSelectLabel string) (FileMgrCollection, error) {

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

  dirPtr, err := os.Open(dMgr.absolutePath)

  if err != nil {
    return fileMgrCol,
      fmt.Errorf(ePrefix+
        "\nError return by os.Open(%v.absolutePath). "+
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
  fh := FileHelper{}

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

        // This is not a directory. It is a file.
        // Determine if it matches the find file criteria.
        isMatch, err = fh.FilterFileName(nameFInfo, fileSelectCriteria)

        if err != nil {

          err2 =
            fmt.Errorf(ePrefix+
              "\nError returned by fh.FilterFileName(nameFInfo, %v).\n"+
              "%v Directory Searched='%v'\nfileName='%v'\nError='%v'\n",
              fileSelectLabel,
              dMgrLabel,
              dMgr.absolutePath,
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
            err2 =
              fmt.Errorf(ePrefix+
                "\nError returned by fileMgrCol.AddFileMgrByFileInfo(%v.absolutePath, nameFInfo). "+
                "%v Directory Path='%v'\nFileName='%v'\nError='%v'\n",
                dMgrLabel,
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
      err2 =
        fmt.Errorf(ePrefix+
          "\nError returned by dirPtr.Close().\n"+
          "%v dirPtr='%v'\nError='%v'\n",
          dMgrLabel,
          dMgr.absolutePath,
          err.Error())
      errs = append(errs, err2)
    }
  }

  return fileMgrCol, dMgrHlpr.consolidateErrors(errs)
}

// findFilesWalkDirectory - This helper method returns file information on
// files residing in a specific directory tree identified by the input
// parameter, 'dMgr'.
//
// This method 'walks the directory tree' locating all files in the directory
// tree which match the file selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to
// zero values or 'Inactive', then ALL FILES in the directory are selected and
// returned in, 'DirectoryTreeInfo.FoundFiles'.
//
func (dMgrHlpr *dirMgrHelper) findFilesWalkDirectory(
  dMgr *DirMgr,
  fileSelectCriteria FileSelectionCriteria,
  ePrefix string,
  dMgrLabel string) (DirectoryTreeInfo, error) {

  var err error
  findFilesInfo := DirectoryTreeInfo{}

  ePrefixCurrMethod := "dirMgrHelper.findFilesWalkDirectory() "

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
    return findFilesInfo, err
  }

  if !dMgrPathDoesExist {
    err = fmt.Errorf(ePrefix+
      "\nERROR: %v Directory Path DOES NOT EXIST!\n"+
      "%v='%v'\n",
      dMgrLabel, dMgrLabel,
      dMgr.absolutePath)

    return findFilesInfo, err
  }

  findFilesInfo.StartPath = dMgr.absolutePath

  findFilesInfo.FileSelectCriteria = fileSelectCriteria

  fh := FileHelper{}

  err = pf.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

  if err != nil {
    return findFilesInfo, fmt.Errorf(ePrefix+
      "Error returned by pf.Walk(findFilesInfo.StartPath, "+
      "fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo)).\n"+
      "dWalkInfo.StartPath='%v'\nError='%v'\n",
      findFilesInfo.StartPath, err.Error())
  }

  return findFilesInfo, nil
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
  fileSelectLabel string) (dirMoveStats DirectoryMoveStats,
  errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.moveDirectory() "

  errs = make([]error, 0, 300)

  if len(ePrefix) == 0 {
    ePrefix = ePrefixCurrMethod
  } else {
    ePrefix = ePrefix + "- " + ePrefixCurrMethod
  }

  var err, err2, err3 error
  var dMgrPathDoesExist, targetDMgrPathDoesExist bool

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
  err3 = nil

  for err3 != io.EOF {

    nameFileInfos, err3 = dir.Readdir(1000)

    if err3 != nil && err3 != io.EOF {

      err2 = fmt.Errorf(ePrefix+
        "\nError returned by dir.Readdirnames(1000).\n"+
        "%v.absolutePath='%v'\nError='%v'\n\n",
        dMgrLabel,
        dMgr.absolutePath,
        err3.Error())

      errs = append(errs, err2)
      err3 = io.EOF
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

            break
          }

          if dirCreated {
            dirMoveStats.DirCreated++
          }

          dirMoveStats.DirCreated++
          targetDMgrPathDoesExist = true
        }

        src = dMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        target = targetDMgr.absolutePath +
          osPathSeparatorStr + nameFInfo.Name()

        err = fh.MoveFile(src, target)

        if err != nil {
          err2 = fmt.Errorf("\n"+ePrefix+
            "ERROR: fh.MoveFile(src, target)\n"+
            "src='%v'\ntarget='%v'\nError='%v'\n\n",
            src, target, err.Error())

          errs = append(errs, err2)
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
  targetDMgrLabel string) (numOfSrcFilesMoved uint64,
  numOfSrcDirectoriesMoved uint64,
  errs []error) {

  ePrefixCurrMethod := "dirMgrHelper.moveDirectoryTree() "

  numOfSrcFilesMoved = 0
  numOfSrcDirectoriesMoved = 0
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

    return numOfSrcFilesMoved,
      numOfSrcDirectoriesMoved,
      errs
  }

  if dTreeCopyStats.FilesNotCopied > 0 {
    err2 = fmt.Errorf(ePrefix+
      "\nError: Some of the files designated to be moved to the target directory, were NOT copied!\n"+
      "Therefore the source directory WILL NOT BE DELETED!\n"+
      "Number of Files NOT Copied='%v'\n",
      "%v Source Directory='%v'\n%v Target Directory='%v'\n\n",
      dTreeCopyStats.FilesNotCopied,
      dMgrLabel, dMgr.absolutePath,
      targetDMgrLabel, targetDMgr.absolutePath)
    errs = append(errs, err2)

    return numOfSrcFilesMoved,
      numOfSrcDirectoriesMoved,
      errs
  }

  numOfSrcFilesMoved = dTreeCopyStats.FilesCopied
  numOfSrcDirectoriesMoved = dTreeCopyStats.DirsCopied

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

  return numOfSrcFilesMoved,
    numOfSrcDirectoriesMoved,
    errs
}
