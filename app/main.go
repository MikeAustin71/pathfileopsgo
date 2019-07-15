package main

import (
  pf "../pathfileops"
  "fmt"
)

/*


import (
  pf "../pathfileops"
  "fmt"
)


*/

func main() {

  mainTest88CopySubDirectoryTree()

}

func mainTest88CopySubDirectoryTree() {

  srcDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\logTest"

  srcDMgr, err := pf.DirMgr{}.New(srcDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  targetDir := "D:\\T06\\TestDirMgr_CopySubDirectoryTree_06"

  fh := pf.FileHelper{}

  _ = fh.DeleteDirPathAll(targetDir)

  targetDMgr, err := pf.DirMgr{}.New(targetDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    _ = fh.DeleteDirPathAll(targetDir)
    return
  }

  fsc := pf.FileSelectionCriteria{}

  var copyEmptyDirectories bool

  copyEmptyDirectories = false

  dTreeStats,
    errs := srcDMgr.CopySubDirectoryTree(targetDMgr, copyEmptyDirectories, fsc)

  if len(errs) > 0 {
    fmt.Printf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n\n%v",
      targetDMgr.GetAbsolutePath(),
      targetDMgr.ConsolidateErrors(errs))

    _ = targetDMgr.DeleteAll()

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    fmt.Println("ERROR: The target directory path DOES NOT EXIST!!")

    return
  }

  fsc = pf.FileSelectionCriteria{}

  srcDTreeInfo, err := srcDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    fmt.Printf("Test Verification Error returned by srcDMgr.FindWalkSubDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'", srcDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  fsc = pf.FileSelectionCriteria{}

  targetDTreeInfo, err := targetDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    fmt.Printf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = targetDMgr.DeleteAll()

    return
  }

  srcDirs := srcDTreeInfo.Directories.GetNumOfDirs()
  srcDirs-- // Discount the one empty subdirectory

  targetDirs := targetDTreeInfo.Directories.GetNumOfDirs()

  if srcDirs != targetDirs {
    fmt.Printf("ERROR: Expected %v-directories would be created.\n"+
      "Instead, %v-directories were created!\n",
      srcDirs, targetDirs)

    _ = targetDMgr.DeleteAll()

    return
  }

  tFileInfo, err := targetDMgr.FindFilesBySelectCriteria(fsc)

  if err != nil {
    fmt.Printf("Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
      "targetDMgr='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
    _ = targetDMgr.DeleteAll()

    return
  }

  if tFileInfo.GetNumOfFileMgrs() > 0 {
    fmt.Printf("ERROR: Expected ZERO files in top level target directory.\n"+
      "Instead, the top level target directory had %v-files.\nTarget Directory='%v'\n",
      tFileInfo.GetNumOfFileMgrs(), targetDMgr.GetAbsolutePath())
  }

  // Subtract 1 to eliminate the empty directory
  expectedDirsCopied := uint64(srcDTreeInfo.Directories.GetNumOfDirs() - 2)
  expectedDirsCreated := uint64(srcDTreeInfo.Directories.GetNumOfDirs() - 2)
  expectedTotalDirsProcessed := uint64(srcDTreeInfo.Directories.GetNumOfDirs())

  if expectedTotalDirsProcessed != dTreeStats.TotalDirsScanned {
    fmt.Printf("Error: Expected dTreeCopyStats.TotalDirsScanned='%v'.\n"+
      "Instead, dTreeCopyStats.TotalDirsScanned='%v'\n",
      expectedTotalDirsProcessed, dTreeStats.TotalDirsScanned)
  }

  if expectedDirsCopied != dTreeStats.DirsCopied {
    fmt.Printf("Error: Expected dTreeCopyStats.DirsCopied='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCopied='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCopied)

  }

  if expectedDirsCreated != dTreeStats.DirsCreated {
    fmt.Printf("Error: Expected dTreeCopyStats.DirsCreated='%v'.\n"+
      "Instead, dTreeCopyStats.DirsCreated='%v'\n",
      expectedDirsCopied, dTreeStats.DirsCreated)
  }

  expectedFilesCopied := uint64(srcDTreeInfo.FoundFiles.GetNumOfFileMgrs())
  expectedFileBytesCopied := srcDTreeInfo.FoundFiles.GetTotalFileBytes()
  expectedFilesNotCopied := uint64(0)
  expectedFileBytesNotCopied := uint64(0)
  expectedTotalFilesProcessed := expectedFilesCopied

  if expectedFilesCopied != dTreeStats.FilesCopied {
    fmt.Printf("Error: Expected dTreeCopyStats.FilesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesCopied='%v'\n",
      expectedFilesCopied, dTreeStats.FilesCopied)
  }

  if expectedFileBytesCopied != dTreeStats.FileBytesCopied {
    fmt.Printf("Error: Expected dTreeCopyStats.FileBytesCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesCopied='%v'\n",
      expectedFileBytesCopied, dTreeStats.FileBytesCopied)
  }

  if expectedFilesNotCopied != dTreeStats.FilesNotCopied {
    fmt.Printf("Error: Expected dTreeCopyStats.FilesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FilesNotCopied='%v'\n",
      expectedFilesNotCopied, dTreeStats.FilesNotCopied)
  }

  if expectedFileBytesNotCopied != dTreeStats.FileBytesNotCopied {
    fmt.Printf("Error: Expected dTreeCopyStats.FileBytesNotCopied='%v'.\n"+
      "Instead, dTreeCopyStats.FileBytesNotCopied='%v'\n",
      expectedFileBytesNotCopied, dTreeStats.FileBytesNotCopied)
  }

  if expectedTotalFilesProcessed != dTreeStats.TotalFilesProcessed {
    fmt.Printf("Error: Expected dTreeCopyStats.TotalFilesProcessed='%v'.\n"+
      "Instead, dTreeCopyStats.TotalFilesProcessed='%v'\n",
      expectedTotalFilesProcessed, dTreeStats.TotalFilesProcessed)
  }

  fmt.Println("             mainTest84CopyDirTree()                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Total Directories Processed: ", dTreeStats.TotalDirsScanned)
  fmt.Println("         Directories Copied: ", dTreeStats.DirsCopied)
  fmt.Println("        Directories Created: ", dTreeStats.DirsCreated)
  fmt.Println("      Total Files Processed: ", dTreeStats.TotalFilesProcessed)
  fmt.Println("               Files Copied: ", dTreeStats.FilesCopied)
  fmt.Println("           Files Not Copied: ", dTreeStats.FilesNotCopied)
  fmt.Println("Copy Empty Directories Flag: ", copyEmptyDirectories)

  /*
    err = targetDMgr.DeleteAll()

    if err != nil {
      fmt.Printf("Test Clean-Up Error returned by targetDMgr.DeleteAll()\n"+
        "Target Directory Absolute Path='%v'\nError='%v'\n", targetDMgr.GetAbsolutePath(), err.Error())
    }
  */

  return
}

/*

func mainTest87CopyDirectoryTree() {

  setUpDir1 := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest"

  setUpDMgr1, err := pf.DirMgr{}.New(setUpDir1)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(setUpDir1)\n"+
      "setUpDir1='%v'\nError='%v'\n",
      setUpDir1, err.Error())
    return
  }

  setupDir2 := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\htmlFilesForTest"

  setUpDMgr2, err := pf.DirMgr{}.New(setupDir2)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(setupDir2)\n"+
      "setupDir2='%v'\nError='%v'\n",
      setupDir2, err.Error())
    return
  }

  srcDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\createFilesTest\\levelfilesfortest"

  srcDMgr, err := pf.DirMgr{}.New(srcDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}

  _,
  errs := setUpDMgr1.CopyDirectoryTree(srcDMgr, false, fsc)

  if len(errs) > 0 {
    fmt.Printf("Test Setup Errors returned by setUpDMgr1.CopyDirectoryTree(srcDMgr, false, fsc).\n"+
      "srcDMgr='%v'\nErrors Follow:\n%v", srcDMgr.GetAbsolutePath(),
      pf.DirMgr{}.ConsolidateErrors(errs).Error())
    return
  }

  srcHtmlDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\createFilesTest\\" +
    "levelfilesfortest\\level_01_dir\\level_02_dir\\htmlFilesForTest"

  srcHtmlDMgr, err := pf.DirMgr{}.New(srcHtmlDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(srcHtmlDir).\n"+
      "srcHtmlDir='%v'\nError='%v'", srcHtmlDir, err.Error())
    return
  }

  fsc = pf.FileSelectionCriteria{}

  _,
    errs = setUpDMgr2.CopyDirectory(srcHtmlDMgr, fsc, false)

  if len(errs) > 0 {
    fmt.Printf("Test Setup Errors returned by setUpDMgr2.CopyDirectory(srcHtmlDMgr, fsc).\n"+
      "srcHtmlDMgr='%v'\nErrors Follow:\n%v\n",
      srcHtmlDMgr.GetAbsolutePath(),
      pf.DirMgr{}.ConsolidateErrors(errs).Error())
    return
  }

  targetDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\dirmgrtests\\levelfilesfortest"

  fh := pf.FileHelper{}

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    fmt.Printf("Error returned from fh.DeleteDirPathAll(targetDir)\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  targetDMgr, err := pf.DirMgr{}.New(targetDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'", targetDir, err.Error())
    return
  }

  fsc = pf.FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  // Copy '.txt' files only to targetDMgr
  dtreeCopyStats,
  errs := srcDMgr.CopyDirectoryTree(
    targetDMgr,
    false,
    fsc)

  if len(errs) > 0 {
    fmt.Printf("Errors returned by srcDMgr.CopyDirectoryTree(targetDMgr, false, fsc)\n"+
      "targetDMgr='%v'\nErrors Follow:\n%v",
      targetDMgr.GetAbsolutePath(),
      pf.DirMgr{}.ConsolidateErrors(errs).Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  if !targetDMgr.DoesAbsolutePathExist() {
    fmt.Printf("ERROR: The target directory path DOES NOT EXIST!!\n"+
      "Number Of FilesCopied='%v'\n", dtreeCopyStats.FilesCopied)

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  fsc = pf.FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.txt"}

  srcTextDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    fmt.Printf("Test Verification Error returned by\n"+
      "srcTextDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'",
      srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)

    return
  }

  expectedNumOfCopiedFiles := srcTextDTreeInfo.FoundFiles.GetNumOfFileMgrs()

  fsc = pf.FileSelectionCriteria{}
  fsc.FileNamePatterns = []string{"*.htm"}

  srcHtmlDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    fmt.Printf("Test Verification Error returned by\n"+
      "srcHtmlDTreeInfo, err := srcDMgr.FindWalkDirFiles(fsc).\n"+
      "source directory='%v'\nError='%v'",
      srcDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)
    return
  }

  expectedNumOfFilesNotCopied := srcHtmlDTreeInfo.FoundFiles.GetNumOfFileMgrs()

  targetDTreeInfo, err := targetDMgr.FindWalkDirFiles(fsc)

  if err != nil {
    fmt.Printf("Test Verification Error returned by targetDMgr.FindWalkDirFiles(fsc).\n"+
      "target directory='%v'\nError='%v'", targetDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(targetDir)
    _ = fh.DeleteDirPathAll(srcDir)
    return
  }

  expectedNumOfDirectoriesCopied := srcTextDTreeInfo.Directories.GetNumOfDirs() - 1

  if expectedNumOfDirectoriesCopied != targetDTreeInfo.Directories.GetNumOfDirs() {
    fmt.Printf("Expected %v-directories would be created. Instead, %v-directories were created!\n"+
      "targetDTreeInfo.Directories.GetNumOfDirs()='%v'\n",
      expectedNumOfDirectoriesCopied,
      targetDTreeInfo.Directories.GetNumOfDirs(),
      targetDTreeInfo.Directories.GetNumOfDirs())

  }

  if uint64(expectedNumOfCopiedFiles) != dtreeCopyStats.FilesCopied {
    fmt.Printf("Expected %v-files would be copied.\n"+
      "Instead, numberOfFilesCopied-'%v'\n",
      expectedNumOfCopiedFiles, dtreeCopyStats.FilesCopied)
  }

  if uint64(expectedNumOfFilesNotCopied) != dtreeCopyStats.FilesNotCopied {
    fmt.Printf("Expected %v-files would NOT be copied.\n"+
      "Instead, numberOfFilesNotCopied='%v'!",
      expectedNumOfFilesNotCopied, dtreeCopyStats.FilesNotCopied)
  }

  if uint64(expectedNumOfDirectoriesCopied) != dtreeCopyStats.DirsCopied {
    fmt.Printf("Expected that %v-directories would be copied.\n"+
      "Instead, %v-directories were copied.",
      expectedNumOfDirectoriesCopied, dtreeCopyStats.DirsCopied)
  }

  err = fh.DeleteDirPathAll(targetDir)

  if err != nil {
    fmt.Printf("Test Clean-Up Error returned by fh.DeleteDirPathAll(fh.DeleteDirPathAll(targetDir)\n"+
      "Target Directory Path='%v'\nError='%v'\n", targetDir, err.Error())
  }

  err = fh.DeleteDirPathAll(srcDir)

  if err != nil {
    fmt.Printf("Test Clean-Up Error returned by fh.DeleteDirPathAll(fh.DeleteDirPathAll(srcDir)\n"+
      "Source Directory Path='%v'\nError='%v'\n", srcDir, err.Error())
  }

  return
}

func mainTest86CopySubDirTree() {

  src := "D:\\T05\\levelfilesfortest"

  dst := "D:\\T06\\levelfilesfortest"

  srcDMgr, err := pf.DirMgr{}.New(src)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(src)\n"+
      "src='%v'\nError='%v'\n", src, err.Error())
    return
  }

  targetDMgr, err := pf.DirMgr{}.New(dst)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(dst)\n"+
      "dst='%v'\nError='%v'\n", dst, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}
  // fsc.FileNamePatterns = []string{"*.htm"}
  var copyEmptyDirectories bool
  copyEmptyDirectories = false

  dtreeStats,
  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, copyEmptyDirectories, fsc)

  if len(errs) > 0 {
    fmt.Printf("Errors returned by srcDMgr.CopyDirectoryTree("+
      "targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\n"+
      "Errors Follow:\n%v",
      targetDMgr.GetAbsolutePath(),
      srcDMgr.ConsolidateErrors(errs))
    return
  }

  if dtreeStats.ComputeError != nil {
    fmt.Printf("Error returned by dtreeStats.ComputeError\n"+
      "Error='%v'\n", dtreeStats.ComputeError.Error())
  }

  fmt.Println("             mainTest84CopyDirTree()                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Total Directories Processed: ", dtreeStats.TotalDirsScanned)
  fmt.Println("         Directories Copied: ", dtreeStats.DirsCopied)
  fmt.Println("        Directories Created: ", dtreeStats.DirsCreated)
  fmt.Println("      Total Files Processed: ", dtreeStats.TotalFilesProcessed)
  fmt.Println("               Files Copied: ", dtreeStats.FilesCopied)
  fmt.Println("           Files Not Copied: ", dtreeStats.FilesNotCopied)
  fmt.Println("Copy Empty Directories Flag: ", copyEmptyDirectories)

}

func mainTest85FindWalkSubDirFiles() {

  testDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\logTest"

  testDMgr, err := pf.DirMgr{}.New(testDir)

  if err != nil {
    fmt.Printf("Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n",
      testDir, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}

  dTreeInfo, err := testDMgr.FindWalkSubDirFiles(fsc)

  if err != nil {
    fmt.Printf("Error returned by testDMgr.FindWalkSubDirFiles(fsc)\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(),
      err.Error())
    return
  }

  expectedNumOfDirs := 7
  expectedNumOfFiles := 5

  if len(dTreeInfo.ErrReturns) > 0 {
    fmt.Printf("dTreeInfo Returned Errors:\n\n%v",
      testDMgr.ConsolidateErrors(dTreeInfo.ErrReturns))
    return
  }

  if expectedNumOfFiles != dTreeInfo.FoundFiles.GetNumOfFileMgrs() {
    fmt.Printf("Error: Expected dTreeInfo.FoundFiles.GetNumOfFileMgrs()='%v'.\n"+
      "Instead, dTreeInfo.FoundFiles.GetNumOfFileMgrs()='%v'\n",
      expectedNumOfFiles, dTreeInfo.FoundFiles.GetNumOfFileMgrs())
    return
  }

  if expectedNumOfDirs != dTreeInfo.Directories.GetNumOfDirs() {
    fmt.Printf("Error: Expected dTreeInfo.Directories.GetNumOfDirs()='%v'\n"+
      "Instead, dTreeInfo.Directories.GetNumOfDirs()='%v'\n",
      expectedNumOfDirs, dTreeInfo.Directories.GetNumOfDirs())
    return
  }

  fmt.Println("            mainTest85FindWalkSubDirFiles               ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Number Of Directories Found: ", dTreeInfo.Directories.GetNumOfDirs())
  fmt.Println("      Number Of Files Found: ", dTreeInfo.FoundFiles.GetNumOfFileMgrs())

}


func mainTest84CopyDirTree() {

  src := "D:\\T05\\levelfilesfortest"

  dst := "D:\\T06\\levelfilesfortest"

  srcDMgr, err := pf.DirMgr{}.New(src)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(src)\n"+
      "src='%v'\nError='%v'\n", src, err.Error())
    return
  }

  targetDMgr, err := pf.DirMgr{}.New(dst)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(dst)\n"+
      "dst='%v'\nError='%v'\n", dst, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}
  //fsc.FileNamePatterns = []string{"*.txt"}
  var copyEmptyDirectories bool
  copyEmptyDirectories = false

  dtreeStats,
    errs := srcDMgr.CopyDirectoryTree(targetDMgr, copyEmptyDirectories, fsc)

  if len(errs) > 0 {
    fmt.Printf("Errors returned by srcDMgr.CopyDirectoryTree("+
      "targetDMgr, true, fsc)\n"+
      "targetDMgr='%v'\n"+
      "Errors Follow:\n%v",
      targetDMgr.GetAbsolutePath(),
      srcDMgr.ConsolidateErrors(errs))
    return
  }

  if dtreeStats.ComputeError != nil {
    fmt.Printf("Error returned by dtreeStats.ComputeError\n"+
      "Error='%v'\n", dtreeStats.ComputeError.Error())
  }

  fmt.Println("             mainTest84CopyDirTree()                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Total Directories Processed: ", dtreeStats.TotalDirsScanned)
  fmt.Println("         Directories Copied: ", dtreeStats.DirsCopied)
  fmt.Println("      Total Files Processed: ", dtreeStats.TotalFilesProcessed)
  fmt.Println("               Files Copied: ", dtreeStats.FilesCopied)
  fmt.Println("           Files Not Copied: ", dtreeStats.FilesNotCopied)
  fmt.Println("Copy Empty Directories Flag: ", copyEmptyDirectories)

}


func mainTest83DmgrDeleteDirAll() {

  //srcDir := "D:\\T04\\checkfiles\\checkfiles03\\dir01\\dir02\\dir03"
  srcDir := "D:\\T04\\checkfiles\\checkfiles03\\dir01"

  dMgr, err := pf.DirMgr{}.New(srcDir)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(srcDir)\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  err = dMgr.DeleteAll()

  if err != nil {
    fmt.Printf("Error returned by dMgr.DeleteAll()\n"+
      "dMgr='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), err.Error())
    return
  }

  fmt.Println("           mainTest83DmgrDeleteDirAll()                 ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()

}

func mainTest82CopyByIO() {

  fh := pf.FileHelper{}

  // setupFileName := "testRead918256.txt"

  //sourceFile := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_3_test.txt"

  sourceFile := "D:\\T03\\ppc_6800_gsg.pdf"

  //		sourceFile := fh.AdjustPathSlash(
  //		"D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\checkfiles\\" + setupFileName)

  sourceFile = fh.AdjustPathSlash(sourceFile)

  destFile := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\ppc_6800_gsg.pdf")

  fileDoesExist, err := fh.DoesThisFileExist(sourceFile)

  if err != nil {
    fmt.Printf("Error returned by fh.DoesThisFileExist(sourceFile)\n"+
      "sourceFile='%v'\nError='%v'\n", sourceFile, err.Error())
    return
  }

  if !fileDoesExist {
    fmt.Printf("Test Setup Error: Source File DOES NOT EXIST!\n"+
      "sourceFile='%v'\n", sourceFile)
    return
  }

  sourceFMgr, err := pf.FileMgr{}.New(sourceFile)

  if err != nil {
    fmt.Printf("Error returned by pf.FileMgr{}.New(sourceFile).\n"+
      "sourceFile='%v'\nError='%v'\n",
      sourceFile, err.Error())
    return
  }

  destFMgr, err := pf.FileMgr{}.New(destFile)

  if err != nil {
    fmt.Printf("Error returned by pf.FileMgr{}.New(destFile).\n"+
      "destFile='%v'\nError='%v'\n",
      destFile, err.Error())
    return
  }

  err = sourceFMgr.CopyFileMgrByIoWithBuffer(&destFMgr, 0)

  if err != nil {
    fmt.Printf("Error returned by fh.CopyFileByIo(sourceFile, destFile)\n"+
      "sourceFile='%v'\ndestFile='%v'\nError='%v'\n",
      sourceFile, destFile, err.Error())
    return
  }

  fileDoesExist, err = fh.DoesThisFileExist(destFile)

  if err != nil {
    fmt.Printf("Error returned by fh.DoesThisFileExist(destFile)\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  if !fileDoesExist {
    fmt.Printf("Error: After the copy operation, the Destination File\n"+
      "DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
    return
  }

  fmt.Println("               mainTest82CopyByIO()                     ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println(" Copied Source File: ", sourceFile)
  fmt.Println()
  fmt.Println("To Destination File: ", destFile)

}


func mainTest81ReadFileLine() {
  // TestFileMgr_ReadFileLine_03
  // xt_filemanager_07_test.go

  fh := pf.FileHelper{}

  setupFileName := "testRead918256.txt"

  setupFile := fh.AdjustPathSlash(
    "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\checkfiles\\" + setupFileName)

  filePath := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\" + setupFileName)

  absBaseFilePath, err := fh.MakeAbsolutePath(
    "D:\\T04\\checkfiles\\checkfiles03")

  if err != nil {
    fmt.Printf("Test Setup Error: Error returned by fh.MakeAbsolutePath"+
      "(\"../checkfiles/checkfiles03/checkfiles03_02\").\n"+
      "Error='%v'\n", err.Error())
    return
  }

  err = fh.MakeDirAll(absBaseFilePath)

  if err != nil {
    fmt.Printf("Test Setup Error: Error returned by fh.MakeDirAll(absBaseFilePath).\n"+
      "absBaseFilePath='%v'\nError='%v'\n", absBaseFilePath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    fmt.Printf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    fmt.Printf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n"+
      "setupFile='%v'\nfilePath='%v'\nError='%v'\n",
      setupFile, filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from common.FileMgr{}."+
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())

    return
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 50)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      fmt.Printf("Error returned by fMgr.ReadFileLine(delim) on "+
        "Line#1.\n"+
        "fMgr='%v'\nError='%v'\n",
        fMgr.GetAbsolutePathFileName(), err.Error())
      _ = fMgr.CloseThisFile()
      return
    }

    fmt.Printf("Line-%v: %v\n", i, string(bytes))

  }

  fmt.Println()

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\n Error='%v'",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  if fMgr.GetFilePtr() != nil {
    fmt.Println("ERROR: After fMgr.CloseThisFile(), expected " +
      "fMgr.filePtr==nil.\n" +
      "However, fMgr.filePtr IS NOT EQUAL TO NIL!")
    _ = fMgr.CloseThisFile()
    return
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  isErr := false

  if "Thank you, for your support." != actualStr {
    fmt.Printf("Expected line #4 = 'Thank you, for your support.'\n"+
      "Instead, line #4 = '%v'\n", actualStr)
    isErr = true
  }

  if !isErrEOF {
    fmt.Println("ERROR: Expected the last error return from fMgr.ReadFileLine(delim)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!")
    isErr = true
  }

  _ = fMgr.CloseThisFile()

  _ = fMgr.DeleteThisFile()

  if isErr {
    return
  }

  fmt.Println("             mainTest81ReadFileLine()                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}

func mainTest80FileAccessCtrlDetection() {

  fileAccessCtrl, err2 := pf.FileAccessControl{}.NewWriteOnlyAccess()

  if err2 != nil {
    fmt.Printf(
      "Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
        "Error='%v'\n", err2.Error())
    return
  }

  fNewOpenType, err2 := fileAccessCtrl.GetFileOpenType()

  if err2 != nil {
    fmt.Printf("Error returned by fileAccessCtrl.GetFileOpenType()!\n"+
      "Error='%v'\n", err2.Error())
    return
  }

  if fNewOpenType != pf.FOpenType.TypeReadWrite() &&
    fNewOpenType != pf.FOpenType.TypeWriteOnly() {

    fmt.Printf("fNewOpenType error!\n"+
      "fNewOpenType=='%v'\n", fNewOpenType.String())
    return
  }

  fmt.Println("         maintTest80FileAccessCtrlDetection()           ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println("fNewOpenType: ", fNewOpenType.String())
}

func maintTest79WriteBytes() {

  fh := pf.FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := "D:\\T04\\checkfiles\\checkfiles03\\testWriteXX241289.txt"

  absFilePath, err := fh.MakeAbsolutePath(filePath)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath(filePath)\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(absFilePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    fmt.Printf("Error returned by fMgr.FlushBytesToDisk(). Error='%v' ",
      err.Error())
    return
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().")
    return
  }

  if verifyBytesWritten != uint64(numBytesWritten) {
    fmt.Printf("verifyBytesWritten != numBytesWritten\n" +
      "verifyBytesWritten='%v'\nnumBytesWritten='%v'\n",
      verifyBytesWritten, uint64(numBytesWritten))
    return
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    fmt.Printf("Error returned by fMgr.ReadFileBytes(bytesRead). Error='%v'",
      err.Error())
    return
  }

  if numBytesRead == 0 {
    fmt.Printf("Number of bytes read returned by fMgr.ReadFileBytes() is ZERO!\n" +
      "fMgr='%v'\n",
      fMgr.GetAbsolutePath())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.CloseThisFile().")
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
    return
  }

  fmt.Println("               maintTest79WriteBytes                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}

func maintTest78WriteBytes() {

  fh := pf.FileHelper{}

  testText := "Now is the time for all good men to come to the aid of their country."

  lenTestText := len(testText)

  filePath := "D:\\T04\\checkfiles\\checkfiles03\\testWriteXX241289.txt"

  absFilePath, err := fh.MakeAbsolutePath(filePath)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath(filePath)\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(absFilePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.FlushBytesToDisk()

  if err != nil {
    fmt.Printf("Error returned by fMgr.FlushBytesToDisk(). Error='%v' ",
      err.Error())
    return
  }

  verifyBytesWritten := fMgr.GetFileBytesWritten()

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().")
    return
  }

  bytesRead := make([]byte, lenTestText+5)

  numBytesRead, err := fMgr.ReadFileBytes(bytesRead)

  if err != nil {
    fmt.Printf("Error returned by fMgr.ReadFileBytes(bytesRead). Error='%v'",
      err.Error())
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.CloseThisFile().")
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED! Error='%v'", err.Error())
    return
  }

  stringRead := string(bytesRead)

  stringRead = stringRead[:len(bytesRead)-5]

  if testText != stringRead {
    fmt.Printf("Error: Expected stringRead='%v'. Instead, stringRead='%v' ",
      testText, stringRead)
    return
  }

  if verifyBytesWritten != uint64(lenTestText) {
    fmt.Printf("Error: verifyBytesWritten != lenTestText. verifyBytesWritten='%v' "+
      "lenTestText='%v' ", verifyBytesWritten, lenTestText)
  }

  if numBytesRead != lenTestText {
    fmt.Printf("Error: numBytesRead != lenTestText. numBytesRead='%v' "+
      "lenTestText='%v' ", numBytesRead, lenTestText)
  }

  if numBytesRead != numBytesWritten {
    fmt.Printf("Error: numBytesRead != numBytesWritten. numBytesRead='%v' "+
      "numBytesWritten='%v' ", numBytesRead, numBytesWritten)
  }

  fmt.Println("               maintTest78WriteBytes                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}

func maintTest77OpenThisFileWriteOnlyAppend() {

  fh := pf.FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  // testText2 := "Damn the torpedoes, full speed ahead!\n"

  filePath := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\TestFileMgr_OpenThisFileWriteOnlyAppend_01.txt")

  basePath := fh.AdjustPathSlash("D:\\T04\\checkfiles")

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    fmt.Printf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateDirAndFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)
  bytesWritten := 0
  // fMgr.isFilePtrOpen = false
  bytesWritten, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fmt.Println("bytesWritten: ", bytesWritten)

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED!\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  fmt.Println("      maintTest77OpenThisFileWriteOnlyAppend            ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}

func mainTest76OpenThisFileWriteOnlyAppend() {

  fh := pf.FileHelper{}

  testText1 := "Now is the time for all good men to come to the aid of their country.\n"

  testText2 := "Damn the torpedoes, full speed ahead!\n"

  filePath := fh.AdjustPathSlash(
    "D:\\T04\\checkfiles\\checkfiles03\\TestFileMgr_OpenThisFileWriteOnlyAppend_01.txt")

  basePath := fh.AdjustPathSlash("D:\\T04\\checkfiles")

  err := fh.DeleteDirPathAll(basePath)

  if err != nil {
    fmt.Printf("Error returned by fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr"+
      "(filePath).\nfilePathName='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fMgr.CreateDirAndFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CreateThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite := []byte(testText1)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnlyAppend().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesToWrite = []byte(testText2)

  _, err = fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.WriteBytesToFile(bytesToWrite).\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  bytesRead1, err := fMgr.ReadFileLine('\n')

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.ReadFileLine(newline).\n"+
      "Error='%v'\n\n", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  bytesRead2, err := fMgr.ReadFileLine('\n')

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.ReadFileLine(newline).\n"+
      "Error='%v'\n", err.Error())
    _ = fMgr.CloseThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #2 fMgr.CloseThisFile().\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("fMgr.DeleteThisFile() FAILED!\n"+
      "fMgr='%v'\nError='%v'\n",
      fMgr.GetAbsolutePathFileName(), err.Error())
    return
  }

  stringRead := string(bytesRead1)

  stringRead = stringRead[:len(stringRead)-1]

  stringRead1 := stringRead

  testText1 = testText1[:len(testText1)-1]

  setSuccess := true

  if testText1 != stringRead {
    fmt.Printf("Error: Expected #1 stringRead='%v'.\n"+
      "Instead, #1 stringRead='%v'\n",
      testText1, stringRead)
    setSuccess = false
  }

  stringRead = string(bytesRead2)

  stringRead = strings.Replace(stringRead, "\r\n", "", -1)

  testText2 = strings.Replace(testText2, "\r\n", "", -1)

  if testText2 != stringRead {
    fmt.Printf("Error: Expected #2 stringRead='%v'.\n"+
      "Instead, #2 stringRead='%v'\n",
      testText2, stringRead)
    setSuccess = false
  }

  if !setSuccess {
    return
  }

  fmt.Println("       mainTest76OpenThisFileWriteOnlyAppend            ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println("stringRead1: ", stringRead1)
  fmt.Println("  testText1: ", testText1)

  return
}


func maintTest75FileMgrGetTimeVal() {

  filePath :=
    "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_2_test.txt"

  fMgr, err := pf.FileMgr{}.New(filePath)

  if err != nil {
    fmt.Printf("Error returned  by pf.FileMgr{}.New(filePath)\n"+
      "filePath='%v'\n"+
      "Error='%v'\n", filePath, err.Error())
    return
  }

  fileModTime, err := fMgr.GetFileModTime()

  if err != nil {
    fmt.Printf("Error returned by fMgr.GetFileModTime()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  fmt.Println("          maintTest75FileMgrGetTimeVal                 ")
  fmt.Println("********************************************************")
  fmt.Println("    fileModTime: ", fileModTime.Format(timeFormatSpec))

}

func mainTest73FileHelperFileExist() {

  filePath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\checkfiles"

  dMgr, err := pf.DirMgr{}.New(filePath)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(filePath)\n"+
      "Error='%v'", err.Error())
    return
  }

  dirDoesExist, err := dMgr.DoesThisDirectoryExist()

  fmt.Println("          mainTest73FileHelperFileExist                 ")
  fmt.Println("********************************************************")
  fmt.Println("    filePath: ", filePath)
  fmt.Println("dirDoesExist: ", dirDoesExist)

}

func mainTest72OpenReadOnlyFile() {
  fh := pf.FileHelper{}

  rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\checkfiles\\TestFileMgr_OpenThisFileReadOnly_03.txt"
  filePath, err := fh.MakeAbsolutePath(rawPath)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath("+
      "rawPath)\n"+
      "rawPath='%v'\n"+
      "Error='%v'\n", rawPath, err.Error())
    return
  }

  err = fh.DeleteDirFile(filePath)

  if err != nil {
    fmt.Printf("Error returned from fh.DeleteDirFile(filePath)\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n", filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  fileDoesExist, err := fMgr.DoesThisFileExist()

  if err != nil {
    fmt.Printf("Non-Path Error returned from #1 fMgr.DoesThisFileExist().\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fh.DeleteDirFile(filePath)
    return
  }

  if fileDoesExist {
    fmt.Printf("ERROR: Test file should NOT exist!.\n"+
      "However, test file DOES EXIST!\n"+
      "test file='%v'", filePath)
    _ = fh.DeleteDirFile(filePath)
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.CreateThisFile().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())
    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.OpenThisFileReadOnly()

  if err != nil {
    fmt.Printf("Error returned from fMgr.OpenThisFileReadOnly().\n"+
      "filePath='%v'\nError='%v'\n",
      fMgr.GetAbsolutePath(), err.Error())

    _ = fMgr.DeleteThisFile()
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.CloseThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }

  err = fMgr.DeleteThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.DeleteThisFile().\n"+
      "Error='%v'\n",
      err.Error())
  }

}

func mainTest71IsPathFileString() {
  fh := pf.FileHelper{}

  testPath := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/" +
    "level_02_dir/iDoNotExist")

  pathFileType, absolutePath, err := fh.IsPathFileString(testPath)

  if err != nil {
    fmt.Printf("Error returned from fh.IsPathFileString(testPath). "+
      "pathFile='%v' Error='%v' ", testPath, err.Error())
    return
  }

  fmt.Println("              mainTest70AdjustPathStr                   ")
  fmt.Println("********************************************************")
  fmt.Println("    testPath: ", testPath)
  fmt.Println("absolutePath: ", absolutePath)
  fmt.Println("pathFileType: ", pathFileType.String())

}

func mainTest70AdjustPathStr() {

  fh := pf.FileHelper{}

  testPath := "../../../"

  adjustedPath := fh.AdjustPathSlash(testPath)

  fmt.Println("              mainTest70AdjustPathStr                   ")
  fmt.Println("********************************************************")
  fmt.Println("    testPath: ", testPath)
  fmt.Println("adjustedPath: ", adjustedPath)
}

func mainTest69CleanDirStr() {

  fh := pf.FileHelper{}

  //   testPathFile := "/d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops/" +
  //     "levelfilesfortest/level_0_0_test.txt"

  //    testPathFile := "d:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\pathfileops" +
  //      "\\levelfilesfortest\\level_0_0_test.txt"


  testPathFile := "../filesfortest//levelfilesfortest/level_01_dir/level_1_1_test.txt"

  absFilePath, err := fh.MakeAbsolutePath(testPathFile)

  if err != nil {

  }

  fmt.Println("              mainTest67AreFilesSame                   ")
  fmt.Println("********************************************************")

  volName := fp.VolumeName(testPathFile)

  cleanFilePath, isEmpty, err := fh.CleanDirStr(testPathFile)

  if err != nil {
    fmt.Printf("Error returned by fh.CleanDirStr(testPathFile)\n"+
      "testPathFile='%v'\nError='%v'\n",
      testPathFile, err.Error())
    return
  }

  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println("testPathFile: ", testPathFile)
  fmt.Println("--------------------------------------------------------")
  fmt.Println("fh.CleanDirStr() Results:")
  fmt.Println("--------------------------------------------------------")
  fmt.Println("      isEmpty: ", isEmpty)
  fmt.Println("    cleanPath: ", cleanFilePath)
  fmt.Println("  Volume Name: ", volName)
  fmt.Println("Absolute Path: ", absFilePath)
  fmt.Println()
}

*/
