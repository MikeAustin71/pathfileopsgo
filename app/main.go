package main

import (
  pf "../pathfileops"
  "fmt"
  "io"
  "os"
  fp "path/filepath"
  "strings"
  "time"
)

/*


import (
  pf "../pathfileops"
  "fmt"
  "io"
  fp "path/filepath"
  "strings"
)


*/

func main() {

  mainTests{}.mainTest107CopyFileMgrByIo06()

}

type mainTests struct {
  Input  string
  Output string
}

func (mtst mainTests) mainTest107CopyFileMgrByIo06() {


  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := pf.FileHelper{}

  baseDir, err := mtst.getBaseProjectPath(true)

  if err != nil {
    fmt.Printf("Error returned from mtst.getBaseProjectPath(true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  targetDir := baseDir + "filesfortest\\newfilesfortest"
  adjustedPath := fh.AdjustPathSlash(targetDir)

  dMgr, err := pf.DirMgr{}.New(adjustedPath)

  if err != nil {
    fmt.Printf("Error returned from DirMgr{}."+
      "NewFromPathFileNameExtStr(adjustedPath).\n"+
      "adjustedPath='%v'\nError='%v'\n",
      adjustedPath, err.Error())
    return
  }

  srcFMgr, err := pf.FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, "+
      "expectedFileNameExt).\n"+
      "dMgr.absolutePath='%v'\nexpectedFileNameExt='%v'\nError='%v'\n",
      dMgr.GetAbsolutePath(), adjustedPath, err.Error())
    return
  }

  destFMgr := srcFMgr.CopyOut()

  err = srcFMgr.CopyFileMgrByIo(&destFMgr)

  if err == nil {
    fmt.Println("Expected error return from CopyFileMgrByIo(&destFMgr) because " +
      "source file is equivalent to destination file. However, NO ERROR WAS RETURNED!")
  }
  
  
} 

func (mtst mainTests) mainTest106GetFileSize() {

  fh := pf.FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

  srcFMgr, err := pf.FileMgr{}.New(targetFile)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.New(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    return
  }

  actualFileSize := srcFMgr.GetFileSize()

  expectedFileSize := int64(29)

  if expectedFileSize != actualFileSize {
    fmt.Printf("Expected file size='29'.\nInstead, file size='%v'\n"+
      "File='%v'",
      actualFileSize, srcFMgr.GetAbsolutePathFileName())
  }

}

func (mtst mainTests) mainTest105FMgrFInfo() {

  testFile :=
    "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_3_test.txt"

  fMgr, err := pf.FileMgr{}.New(testFile)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.New(testFile)\n" +
      "testFile='%v'\n" +
      "Error='%v'\n", testFile, err.Error())
    return
  }

  fInfoPlus, err := fMgr.GetFileInfoPlus()

  if err != nil {
    fmt.Printf("Error returned from fMgr.GetFileInfoPlus()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  result := fInfoPlus.IsDirectoryPathInitialized()

  if result == false {
    fmt.Printf("ERROR: Expected fInfoPlus.IsDirectoryPathInitialized() would return 'true'\n" +
      "because 'fInfoPlus' is properly initialized.\n" +
      "However, fInfoPlus.IsFileInfoInitialized() returned 'false'\n")
  }

}

func (mtst mainTests) mainTest104FileOpsColEqual() {

  sf := make([]string, 5, 10)

  sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
  sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
  sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

  df := make([]string, 5, 10)

  df[0] = "../dirmgrtests/level_0_0_test.txt"
  df[1] = "../dirmgrtests/level_0_1_test.txt"
  df[2] = "../dirmgrtests/level_0_2_test.txt"
  df[3] = "../dirmgrtests/level_0_3_test.txt"
  df[4] = "../dirmgrtests/level_0_4_test.txt"

  fOpsCol1 := pf.FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    fOp, err := pf.FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

    if err != nil {
      fmt.Printf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i]). "+
        "i='%v'\n" +
        "sf[i]='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v'\n", i,sf[i], df[i], err.Error())
      return
    }

    err = fOp.SetFileOpsCode(pf.FileOpCode.CopySourceToDestinationByHardLink())

    if err != nil {
      fmt.Printf("Error returned by fOp.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
        "Error='%v'\n", err.Error())
      return
    }

    err = fOpsCol1.AddByFileOps(fOp)

    if err != nil {
      fmt.Printf("Error returned by fOpsCol1.AddByFileOps(fOp). "+
        "i='%v'\n" +
        "srcFile='%v'\n" +
        "df[i]='%v'\n" +
        "Error='%v' ", i,sf[i], df[i], err.Error())
      return
    }

  }

  fOpsCol2 := pf.FileOpsCollection{}.New()

  for i := 0; i < 5; i++ {

    if i == 2 {

      fOp, err := pf.FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

      if err != nil {
        fmt.Printf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i]). "+
          "i='%v'\n" +
          "sf[i]='%v'\n" +
          "df[i]='%v'\n" +
          "Error='%v'\n", i,sf[i], df[i], err.Error())
        return
      }

      err = fOp.SetFileOpsCode(pf.FileOpCode.CopySourceToDestinationByIo())

      if err != nil {
        fmt.Printf("Error returned by fOp.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
          "Error='%v'\n", err.Error())
        return
      }

      err = fOpsCol2.AddByFileOps(fOp)

      if err != nil {
        fmt.Printf("Error returned by fOpsCol2.AddByFileOps(fOp). "+
          "i='%v'\n" +
          "srcFile='%v'\n" +
          "df[i]='%v'\n" +
          "Error='%v' ", i,sf[i], df[i], err.Error())
        return
      }

    } else {

      fOp, err := pf.FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i])

      if err != nil {
        fmt.Printf("Error returned by FileOps{}.NewByPathFileNameExtStrs(sf[i], df[i]). "+
          "i='%v'\n" +
          "sf[i]='%v'\n" +
          "df[i]='%v'\n" +
          "Error='%v'\n", i,sf[i], df[i], err.Error())
        return
      }

      err = fOp.SetFileOpsCode(pf.FileOpCode.CopySourceToDestinationByHardLink())

      if err != nil {
        fmt.Printf("Error returned by fOp.SetFileOpsCode(FileOpCode.CopySourceToDestinationByIo())\n" +
          "Error='%v'\n", err.Error())
        return
      }

      err = fOpsCol2.AddByFileOps(fOp)

      if err != nil {
        fmt.Printf("Error returned by fOpsCol2.AddByFileOps(fOp). "+
          "i='%v'\n" +
          "srcFile='%v'\n" +
          "df[i]='%v'\n" +
          "Error='%v' ", i,sf[i], df[i], err.Error())
        return
      }
    }
  }

  if fOpsCol1.Equal(&fOpsCol2) == true {
    fmt.Println("ERROR: Expected that fOpsCol1!=fOpsCol2.\n" +
      "However, THEY ARE EQUAL!!!")
  }

}

func (mtst mainTests) mainTest103FHlprConsolidateErrors() {

  errs := make([]error, 0, 100)

  for i:=0; i < 3; i++ {
    errNo := fmt.Sprintf("Error #%0.3d: Error message.\n\n\n", i)
    err := fmt.Errorf(errNo)

    errs = append(errs, err)
  }

  fh := pf.FileHelper{}

  err := fh.ConsolidateErrors(errs)

  if err == nil {
    fmt.Printf("Error return from fh.ConsolidateErrors(errs) is 'nil'\n")
    return
  }

  errStr := fmt.Sprintf("%v", err.Error())

  errFmtStr := strings.ReplaceAll(errStr, "\n", "@")

  fmt.Println("          mainTest103FHlprConsolidateErrors             ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Decoded Err Str: ", errFmtStr)
  fmt.Println("Error String:")
  fmt.Printf("%vTrailer For Test", err.Error())



}

func (mtst mainTests) mainTest102TestNewFromKnownPath(parentDirectory, subDirectoryName string) {


  dMgr, err :=
    pf.DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    fmt.Printf("%v", err.Error())
    return
  }

  fmt.Println("          mainTest102TestNewFromKnownPath               ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println(" parentDirectory: ", parentDirectory)
  fmt.Println("subDirectoryName: ", subDirectoryName)
  fmt.Println(" DirMgr Abs Path: ", dMgr.GetAbsolutePath())
}

func (mtst mainTests) mainTest101TestBadPathChars(pathStr string) {

  fh := pf.FileHelper{}

  isErrorDetected := false
  doublePathSeparator := string(os.PathSeparator) + string(os.PathSeparator)
  tPathStr := fh.AdjustPathSlash(pathStr)

  if strings.Contains(tPathStr, doublePathSeparator) {
    isErrorDetected = true
  }

  fmt.Println("            mainTest101TestBadPathChars                 ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Original Path String: ", pathStr)
  fmt.Println("      Error Detected: ", isErrorDetected)

}

func (mtst mainTests) mainTest100GetAbsPath(pathStr string ) {

  fh := pf.FileHelper{}
  absOrigDir, err := fh.MakeAbsolutePath(pathStr)

  if err != nil {
    fmt.Printf("Error returned from fh.MakeAbsolutePath(pathStr).\n" +
      "pathStr='%v'\n"+
      "Error='%v'",
      pathStr,
      err.Error())
    return
  }

  fmt.Println("              mainTest100GetAbsPath                     ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Original Path String: ", pathStr)
  fmt.Println("       Absolute Path: ", absOrigDir)

}

func (mtst mainTests) mainTest99GetVolumeName(pathStr string) {

  fh := pf.FileHelper{}

  volIdx,
  volLen,
  volName := fh.GetVolumeNameIndex(pathStr)

  fmt.Println("              mainTest99GetVolumeName                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Original Path String: ", pathStr)
  fmt.Println("        Volume Index: ", volIdx)
  fmt.Println("       Volume Length: ", volLen)
  fmt.Println("         Volume Name: ", volName)

}

func (mtst mainTests) mainTest98ParseValidPathStr(pathStr string) {
  dMgr := pf.DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(pathStr)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(pathStr)\n"+
      "pathStr='%v'\n"+
      "Error='%v'\n",
      pathStr, err.Error())
    return
  }

  err = validPathDto.IsDtoValid("mainTest97ParseValidPathStr() ")

  if err != nil {
    fmt.Printf("%v", err.Error())
    return
  }

  fmt.Println("           mainTest97ParseValidPathStr                  ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("       Original Path String: ", pathStr)
  fmt.Println("Original Path String Length: ", len(pathStr))
  fmt.Println("Original Path String Length: ", validPathDto.GetOriginalPathStr())
  fmt.Println("      Validated Path String: ", validPathDto.GetPath())
  fmt.Println("         Path String Length: ", validPathDto.GetPathStrLen())
  fmt.Println("    Validated Absolute Path: ", validPathDto.GetAbsPath())
  fmt.Println("Absolute Path String Length: ", validPathDto.GetAbsPathStrLen())
  fmt.Println("                Volume Name: ", validPathDto.GetPathVolumeName())
  fmt.Println("               Volume Index: ", validPathDto.GetPathVolumeIndex())
  fmt.Println("       Volume String Length: ", validPathDto.GetPathVolumeStrLength())
  fmt.Println(" Valid Path Dto Initialized: ", validPathDto.IsInitialized())
  fmt.Println(" Valid Path Dto PathIsValid: ", validPathDto.GetPathIsValid())

  if validPathDto.GetPathStrLen() != len(validPathDto.GetPath()) {
    fmt.Printf("Path String Length Error!\n"+
      "Expected String Length='%v'\nActual String Length='%v'\n",
      validPathDto.GetPathStrLen(), len(validPathDto.GetPath()))
  }

  if validPathDto.GetAbsPathStrLen() != len(validPathDto.GetAbsPath()) {
    fmt.Printf("Absolute Path String Length Error!\n"+
      "Expected String Length='%v'\nActual String Length='%v'\n",
      validPathDto.GetAbsPathStrLen(), len(validPathDto.GetAbsPath()))
  }

}

func (mtst mainTests) mainTest97DirNew03(origDir, expectedPath string) {

  fh := pf.FileHelper{}

  rawOrigDir := origDir

  origDir = fh.AdjustPathSlash(origDir)

  expectedPath = fh.AdjustPathSlash(expectedPath)

  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    fmt.Printf("Error returned by err := fh.MakeAbsolutePath(expectedPath)\n"+
      "expectedPath='%v'\n"+
      "Error='%v'\n",
      expectedPath, err.Error())
    return
  }

  dMgr, err := pf.DirMgr{}.New(origDir)

  if err != nil {
    fmt.Printf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n",
      origDir, err.Error())
    return
  }

  if expectedPath != dMgr.GetPath() {
    fmt.Printf("ERROR: Expected path='%v'\n"+
      "Instead, path='%v'\n",
      expectedPath, dMgr.GetPath())
  }

  if expectedAbsPath != dMgr.GetAbsolutePath() {
    fmt.Printf("ERROR: Expected absolute path='%v'\n"+
      "Instead, absolute path='%v'\n",
      expectedAbsPath, dMgr.GetAbsolutePath())
    return
  }

  fmt.Println("                mainTest96DirNew02                      ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("         Original Path: ", rawOrigDir)
  fmt.Println("         Expected Path: ", expectedPath)
  fmt.Println("           Actual Path: ", dMgr.GetPath())
  fmt.Println("Expected Absolute Path: ", expectedAbsPath)
  fmt.Println("  Actual Absolute Path: ", dMgr.GetAbsolutePath())

  return
}

func (mtst mainTests) mainTest96DirNew02() {

  fh := pf.FileHelper{}

  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    fmt.Printf("Error returned by err := fh.MakeAbsolutePath(expectedPath)\n"+
      "expectedPath='%v'\n"+
      "Error='%v'\n",
      expectedPath, err.Error())
    return
  }

  dMgr, err := pf.DirMgr{}.New(origDir)

  if err != nil {
    fmt.Printf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n",
      origDir, err.Error())
    return
  }

  if expectedPath != dMgr.GetPath() {
    fmt.Printf("ERROR: Expected path='%v'\n"+
      "Instead, path='%v'\n",
      expectedPath, dMgr.GetPath())
  }

  if expectedAbsPath != dMgr.GetAbsolutePath() {
    fmt.Printf("ERROR: Expected absolute path='%v'\n"+
      "Instead, absolute path='%v'\n",
      expectedAbsPath, dMgr.GetAbsolutePath())
    return
  }

  fmt.Println("                mainTest96DirNew02                      ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("         Expected Path: ", expectedPath)
  fmt.Println("           Actual Path: ", dMgr.GetPath())
  fmt.Println("Expected Absolute Path: ", expectedAbsPath)
  fmt.Println("  Actual Absolute Path: ", dMgr.GetAbsolutePath())

  return
}

func (mtst mainTests) mainTest95DirNew01() {

  fh := pf.FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/.git")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2/.git")

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    fmt.Printf("Test Startup Error returned by fh.MakeAbsolutePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n", expectedPath, err.Error())
    return
  }

  dMgr, err := pf.DirMgr{}.New(origDir)

  if err != nil {
    fmt.Printf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v' Error='%v'\n",
      origDir, err.Error())
    return
  }

  if true != dMgr.IsInitialized() {
    fmt.Printf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.IsInitialized())
    return
  }

  if expectedAbsDir != dMgr.GetAbsolutePath() {
    fmt.Printf("Expected absolute path does match actual absolute path!\n"+
      "Expected absolute path='%v'\n"+
      "Actual absolute path='%v'\n",
      expectedAbsDir, dMgr.GetAbsolutePath())
  }

  fmt.Println("                mainTest95DirNew01                      ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Expected Absolute Path: ", expectedAbsDir)
  fmt.Println("  Actual Absolute Path: ", dMgr.GetAbsolutePath())
}

func (mtst mainTests) mainTest94Index02() {
  fh := pf.FileHelper{}

  pathStr := fh.AdjustPathSlash("../dir1/dir2/dir3/")
  //dirName := string(os.PathSeparator) +  ".git"
  dirName := string(os.PathSeparator)

  origPathName := pathStr

  lPathStr := len(pathStr)

  if pathStr[lPathStr-1] == os.PathSeparator {
    pathStr = pathStr[0 : lPathStr-1]
  }

  newDirName := dirName

  if newDirName[0] == os.PathSeparator {
    newDirName = newDirName[1:]
  }

  fullPath := pathStr + string(os.PathSeparator) + newDirName

  if len(newDirName) == 0 {
    fullPath = pathStr
  }

  fmt.Println("                mainTest94Index02                       ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("   pathStr: ", origPathName)
  fmt.Println("   dirName: ", dirName)
  fmt.Println("newDirName: ", newDirName)
  fmt.Println("  fullPath: ", fullPath)

  // Print out
  // pathStr:  ..\dir1\dir2\dir3\
  // dirName:  \
  // newDirName:
  // fullPath:  ..\dir1\dir2\dir3
}

func (mtst mainTests) mainTest93Index01() {
  fh := pf.FileHelper{}

  pathStr := fh.AdjustPathSlash("../dir1/dir2/dir3")
  dirName := ".git"

  origPathName := pathStr

  lPathStr := len(pathStr)

  if pathStr[lPathStr-1] != os.PathSeparator {
    pathStr += string(os.PathSeparator)
  }

  fullPath := pathStr + dirName

  fmt.Println("                mainTest93Index01                       ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println(" pathStr: ", origPathName)
  fmt.Println(" dirName: ", dirName)
  fmt.Println("fullPath: ", fullPath)

}

func (mtst mainTests) mainTest92GetDirTreeBytes() {

  var tStart, tEnd time.Time

  targetDir := "D:\\T88"
  // targetDir := "D:\\T88\\pathfileopsgo"
  // targetDir := "D:\\T05\\filesfortest"

  testDMgr, err := pf.DirMgr{}.New(targetDir)

  if err != nil {
    fmt.Printf("Error returned by DirMgr{}.New(targetDir)\n"+
      "targetDir='%v'\n"+
      "Error='%v'\n\n", targetDir, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}

  fmt.Println("Starting Search # 1 ...")
  fmt.Println()

  tStart = time.Now()

  testDInfo1, errs := testDMgr.FindDirectoryTreeFiles(fsc)

  tEnd = time.Now()

  if len(errs) > 0 {
    fmt.Printf("Error returned by testDMgr.FindDirectoryTreeFiles(fsc)\n"+
      "testDMgr='%v'\n"+
      "Errors Follow:\n\n'%v'", targetDir,
      testDMgr.ConsolidateErrors(errs))
    return
  }

  duration1Str := mtst.timer(tStart, tEnd)

  total1FileBytes := testDInfo1.FoundFiles.GetTotalFileBytes()
  total1Files := testDInfo1.FoundFiles.GetNumOfFiles()
  total1Dirs := testDInfo1.Directories.GetNumOfDirs()

  fmt.Println("           mainTest92GetDirTreeBytes()                  ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("         Target Directory: ", targetDir)
  fmt.Println("   Test Directory Manager: ", testDMgr.GetAbsolutePath())
  fmt.Println("      Total Number of Files: ", total1Files)
  fmt.Println("           Total File Bytes: ", total1FileBytes)
  fmt.Println("Total Number of Directories: ", total1Dirs)
  fmt.Println("             Elapsed Time: ", duration1Str)
  fmt.Println()
}

func (mtst mainTests) mainTest91GetDirTreeBytes() {

  var tStart, tEnd time.Time

  /*
    basePath, err := mtst.getBaseProjectPath(true)

    if err != nil {
      fmt.Printf("Error returned by mtst.getBaseProjectPath(false)\n"+
        "Error='%v'\n\n", err.Error())
      return
    }
  */

  // testDir := basePath + "filesfortest"
  testDir := "D:\\T88"

  testDMgr, err := pf.DirMgr{}.New(testDir)

  if err != nil {
    fmt.Printf("Error returned by DirMgr{}.New(testDir)\n"+
      "testDir='%v'\n"+
      "Error='%v'\n\n", testDir, err.Error())
    return
  }

  tStart = time.Now()

  dirStats, errs := testDMgr.GetDirectoryStats()

  tEnd = time.Now()

  if len(errs) > 0 {
    fmt.Printf("Error returned by testDMgr.GetDirectoryStats()\n"+
      "testDMgr='%v'\n"+
      "Errors Follow:\n\n%v",
      testDMgr.GetAbsolutePath(),
      testDMgr.ConsolidateErrors(errs))

    return
  }

  duration1Str := mtst.timer(tStart, tEnd)

  fmt.Println("              mainTest91GetDirTreeBytes()                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("           Test Directory: ", testDir)
  fmt.Println("   Test Directory Manager: ", testDMgr.GetAbsolutePath())
  fmt.Println("Number of Sub-Directories: ", dirStats.NumOfSubDirs())
  fmt.Println("          Number of Files: ", dirStats.NumOfFiles())
  fmt.Println("     Number Of File Bytes: ", dirStats.NumOfBytes())
  fmt.Println("             Elapsed Time: ", duration1Str)
  fmt.Println()

}

func (mtst mainTests) maintTest90GetBaseProject() {

  basePath, err := mtst.getBaseProjectPath(false)

  if err != nil {
    fmt.Printf("Error returned by mtst.getBaseProjectPath(false)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  basePathWithSeparator, err := mtst.getBaseProjectPath(true)

  if err != nil {
    fmt.Printf("Error returned by mtst.getBaseProjectPath(true)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  localPath := "archive"
  fh := pf.FileHelper{}
  absoluteLocalPath := fh.JoinPathsAdjustSeparators(basePath, localPath)

  fmt.Println("          maintTest90GetBaseProject()                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("Project Base Path Without Separator: ", basePath)
  fmt.Println("   Project Base Path With Separator: ", basePathWithSeparator)
  fmt.Println("             Local Constructed Path: ", absoluteLocalPath)

}

func (mtst mainTests) mainTest89MoveSubDirectoryTree() {

  originDir := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\logTest"

  originDMgr, err := pf.DirMgr{}.New(originDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(originDir).\n"+
      "originDir='%v'\nError='%v'", originDir, err.Error())
    return
  }

  srcDir := "D:\\T06\\TestDirMgr_MoveSubDirectoryTree_02"

  srcDMgr, err := pf.DirMgr{}.New(srcDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'", srcDir, err.Error())
    return
  }

  destDir := "D:\\T07\\TestDirMgr_MoveSubDirectoryTree_02"

  destDMgr, err := pf.DirMgr{}.New(destDir)

  if err != nil {
    fmt.Printf("Test Setup Error returned by DirMgr{}.New(destDir).\n"+
      "destDir='%v'\nError='%v'", destDir, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}

  _,
  errs := originDMgr.CopyDirectoryTree(srcDMgr, true, fsc)

  if len(errs) > 0 {
    fmt.Printf("Test Setup Errors returned by originDMgr.CopyDirectoryTree(srcDMgr, true, fsc).\n"+
      "srcDMgr='%v'\nErrors Follow:\n\n%v", srcDMgr.GetAbsolutePath(),
      originDMgr.ConsolidateErrors(errs).Error())
    return
  }

  dirMoveStats,
  errs :=
    srcDMgr.MoveSubDirectoryTree(destDMgr)

  if len(errs) > 0 {
    fmt.Printf("Test Setup Errors returned by srcDMgr.MoveSubDirectoryTree(destDMgr).\n"+
      "destDMgr='%v'\nErrors Follow:\n\n%v", destDMgr.GetAbsolutePath(),
      originDMgr.ConsolidateErrors(errs).Error())
    return
  }

  fmt.Println("          mainTest89MoveSubDirectoryTree()              ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")
  fmt.Println()
  fmt.Println("     dirMoveStats.NumOfSubDirectories: ", dirMoveStats.NumOfSubDirectories)
  fmt.Println("    dirMoveStats.SourceFilesRemaining: ", dirMoveStats.SourceFilesRemaining)
  fmt.Println("dirMoveStats.SourceFileBytesRemaining: ", dirMoveStats.SourceFileBytesRemaining)
  fmt.Println("        dirMoveStats.SourceFilesMoved: ", dirMoveStats.SourceFilesMoved)
  fmt.Println("    dirMoveStats.SourceFileBytesMoved: ", dirMoveStats.SourceFileBytesMoved)
  fmt.Println("  dirMoveStats.TotalSrcFilesProcessed: ", dirMoveStats.TotalSrcFilesProcessed)
  fmt.Println("     dirMoveStats.NumOfSubDirectories: ", dirMoveStats.NumOfSubDirectories)
  fmt.Println("             dirMoveStats.DirsCreated: ", dirMoveStats.DirsCreated)
  fmt.Println("      dirMoveStats.TotalDirsProcessed: ", dirMoveStats.TotalDirsProcessed)

}

func (mtst mainTests) mainTest88CopySubDirectoryTree() {

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

  return
}

func (mtst mainTests) mainTest87CopyDirectoryTree() {

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

func (mtst mainTests) mainTest86CopySubDirTree() {

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

func (mtst mainTests) mainTest85FindWalkSubDirFiles() {

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

func (mtst mainTests) mainTest84CopyDirTree() {

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

func (mtst mainTests) mainTest83DmgrDeleteDirAll() {

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

func (mtst mainTests) mainTest82CopyByIO() {

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

func (mtst mainTests) mainTest81ReadFileLine() {
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

func (mtst mainTests) mainTest80FileAccessCtrlDetection() {

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

func (mtst mainTests) maintTest79WriteBytes() {

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
    fmt.Printf("verifyBytesWritten != numBytesWritten\n"+
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
    fmt.Printf("Number of bytes read returned by fMgr.ReadFileBytes() is ZERO!\n"+
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

func (mtst mainTests) maintTest78WriteBytes() {

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

func (mtst mainTests) maintTest77OpenThisFileWriteOnlyAppend() {

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

func (mtst mainTests) mainTest76OpenThisFileWriteOnlyAppend() {

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

func (mtst mainTests) maintTest75FileMgrGetTimeVal() {

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

func (mtst mainTests) mainTest73FileHelperFileExist() {

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

func (mtst mainTests) mainTest72OpenReadOnlyFile() {
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

func (mtst mainTests) mainTest71IsPathFileString() {
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

func (mtst mainTests) mainTest70AdjustPathStr() {

  fh := pf.FileHelper{}

  testPath := "../../../"

  adjustedPath := fh.AdjustPathSlash(testPath)

  fmt.Println("              mainTest70AdjustPathStr                   ")
  fmt.Println("********************************************************")
  fmt.Println("    testPath: ", testPath)
  fmt.Println("adjustedPath: ", adjustedPath)
}

func (mtst mainTests) mainTest69CleanDirStr() {

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

// getBaseProjectPath - Gets the base path on this machine for the
// 'pathfileopsgo' project.
//
func (mtst mainTests) getBaseProjectPath(
  addTrailingPathSeparator bool) (basePath string, err error) {

  ePrefix := "getBaseProjectPath() "
  fh := pf.FileHelper{}

  basePath = ""
  err = nil
  currDir, err2 := fh.GetAbsCurrDir()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fh.GetAbsCurrDir().\nError='%v'\n", err2.Error())

    return basePath, err
  }

  target := "pathfileopsgo"
  idx := strings.Index(currDir, target)

  if idx < 0 {
    err = fmt.Errorf(ePrefix +
      "Error: Unable to locate \"pathfileopsgo\" in current directory string!\n")

    return basePath, err
  }

  idx += len(target)

  basePath = currDir[0:idx]

  if addTrailingPathSeparator {
    basePath += string(os.PathSeparator)
  }

  return basePath, err
}

func (mtst mainTests) timer(starTime, endTime time.Time) string {

  // MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
  // 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
  MicroSecondNanoseconds := int64(time.Microsecond)

  // MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
  //	 A millisecond is 1/1,000 or 1 one-thousandth of a second
  MilliSecondNanoseconds := int64(time.Millisecond)

  // SecondNanoseconds - Number of Nanoseconds in a Second
  SecondNanoseconds := int64(time.Second)

  // MinuteNanoseconds - Number of Nanoseconds in a minute
  MinuteNanoseconds := int64(time.Minute)

  // HourNanoseconds - Number of Nanoseconds in an hour
  HourNanoseconds := int64(time.Hour)

  t2Dur := endTime.Sub(starTime)

  str := ""

  totalNanoseconds := t2Dur.Nanoseconds()
  numOfHours := int64(0)
  numOfMinutes := int64(0)
  numOfSeconds := int64(0)
  numOfMillisecionds := int64(0)
  numOfMicroseconds := int64(0)
  numOfNanoseconds := int64(0)

  if totalNanoseconds >= HourNanoseconds {
    numOfHours = totalNanoseconds / HourNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfHours * HourNanoseconds)
  }

  if totalNanoseconds >= MinuteNanoseconds {
    numOfMinutes = totalNanoseconds / MinuteNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMinutes * MinuteNanoseconds)
  }

  if totalNanoseconds >= SecondNanoseconds {
    numOfSeconds = totalNanoseconds / SecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
  }

  if totalNanoseconds >= SecondNanoseconds {
    numOfSeconds = totalNanoseconds / SecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
  }

  if totalNanoseconds >= MilliSecondNanoseconds {
    numOfMillisecionds = totalNanoseconds / MilliSecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMillisecionds * MilliSecondNanoseconds)
  }

  if totalNanoseconds >= MicroSecondNanoseconds {
    numOfMicroseconds = totalNanoseconds / MicroSecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMicroseconds * MicroSecondNanoseconds)
  }

  numOfNanoseconds = totalNanoseconds

  if numOfHours > 0 {

    str += fmt.Sprintf("%v-Hours ", numOfHours)

  }

  if numOfMinutes > 0 {

    str += fmt.Sprintf("%v-Minutes ", numOfMinutes)

  }

  if numOfSeconds > 0 || str != "" {

    str += fmt.Sprintf("%v-Seconds ", numOfSeconds)

  }

  if numOfMillisecionds > 0 || str != "" {

    str += fmt.Sprintf("%v-Milliseconds ", numOfMillisecionds)

  }

  if numOfMicroseconds > 0 || str != "" {

    str += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

  }

  str += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

  return str
}
