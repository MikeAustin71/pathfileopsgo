package main

import (
  pf "../pathfileops"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

/*

import (
  "MikeAustin71/pathfilego/003_filehelper/common"
  "fmt"
  "time"
)

const (
  baseAppDir = "../../003_filehelper/app"
  // commonDir        = "../common"
  logTestTopDIR = "../logTest"
  // logTestBottomDir = "../logTest/CmdrX"
  // logFile          = "CmdrX.log"
)

func main() {
 if, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok {
        fmt.Println("File at path", err.path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

*/

func main() {
  //source := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  source := "D:\\T11\\level01"

  mainTest57GetFileMode(source)

  fPermission, err := pf.FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    fmt.Printf("main()\n" + err.Error() + "\n")
    return
  }

  mainTest58ChangeFileMode(source, fPermission)
  mainTest57GetFileMode(source)

}

func mainTest58ChangeFileMode(pathFileName string, filePermission pf.FilePermissionConfig) {

  fh := pf.FileHelper{}

  fmt.Println("********* mainTest58ChangeFileMode *********")
  fmt.Println()

  err := fh.ChangeFileMode(pathFileName, filePermission)

  changeModeText, err := filePermission.GetPermissionTextCode()

  if err != nil {
    fmt.Printf("mainTest58ChangeFileMode\nError='" + err.Error() + "'\n")
    return
  }

  changeModeValue := filePermission.GetPermissionFileModeValueText()

  fmt.Println("     pathFileName: ", pathFileName)
  fmt.Println()
  fmt.Println(" Change Mode Text: ", changeModeText)
  fmt.Println("Change Mode Value: ", changeModeValue)


}

func mainTest57GetFileMode(pathFileName string) {
  fh := pf.FileHelper{}

  fPermCfg, err := fh.GetFileMode(pathFileName)

  fmt.Println("********* mainTest57GetFileMode *********")
  fmt.Println()
  if err != nil {
    fmt.Printf("mainTest57GetFileMode\nError='" + err.Error() + "'\n")
    return
  }
  fmt.Println("*************** SUCCESS! ****************")

  permCodeText, err := fPermCfg.GetPermissionTextCode()
  permCodeValue := fPermCfg.GetPermissionFileModeValueText()

  fmt.Println(" Path File Name: ", pathFileName)
  fmt.Println()
  fmt.Println(" File Mode Text: ", permCodeText)
  fmt.Println("File Mode Value: ", permCodeValue)

}

func mainTest56CopyFileLink(source, destination string) {

  fh := pf.FileHelper{}

  err := fh.CopyFileByLink(source, destination)

  if err != nil {
    fmt.Printf("Error returned by fh.CopyFileByLink(source, destination). \n" +
      "source='%v'\ndestination='%v'\nError='%v'",
      source, destination, err.Error())
    return
  }

  fmt.Println("********** mainTest56CopyFileLink **********")
  fmt.Println("")
  fmt.Println("           SUCCESS!!!           ")

}

func mainTest55CopyFileIo(source, destination string) {

  fh := pf.FileHelper{}

  err := fh.CopyFileByIo(source, destination)

  if err != nil {
    fmt.Printf("Error returned by fh.CopyFileByIo(source, destination). \n" +
      "source='%v'\ndestination='%v'\nError='%v'",
      source, destination, err.Error())
    return
  }

  fmt.Println("********** mainTest55CopyFileIo **********")
  fmt.Println("")
  fmt.Println("           SUCCESS!!!           ")


}

func mainTest54DeleteAll(targetDir string) {
  // Deletes everything: files and directories.

  fh := pf.FileHelper{}

  err := fh.DeleteDirPathAll(targetDir)

  if err != nil {
    fmt.Printf("Error returned by fh.DeleteDirPathAll(targetDir). \n"+
      "newDir='%v' Error='%v' \n", targetDir, err.Error())
    return
  }

  _, err = os.Stat(targetDir)

  if os.IsExist(err) {
    fmt.Println("os.Stat(targetDir) returned an 'IsExist' type error!")
  }

  if os.IsNotExist(err) {
    fmt.Println("os.Stat(targetDir) returned an 'IsNotExist' type error!")
  }

  if err !=nil {
    fmt.Printf("Expected error reads: %v\n\n", err.Error())
  }

  if err == nil  {
    fmt.Printf("Deletion Failed! \n +"+
      "os.Stat(targetDir) targetDir='%v' \n"+
      "Error='%v' \n", targetDir, err.Error())
    return
  }


  fmt.Println("********** mainTest54DeleteAll **********")
  fmt.Println("")
  fmt.Println("           SUCCESS!!!           ")
  fmt.Println("Deleted Directory: ", targetDir)

}

func mainTest53CreateNewDir(newDir string) {
  // Creates a New Directory including
  // parent directories
  fh := pf.FileHelper{}

  err := fh.MakeDirAll(newDir)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeDirAll(newDir). \n"+
      "newDir='%v' Error='%v' \n", newDir, err.Error())
    return
  }

  _, err = os.Stat(newDir)

  if err != nil {
    fmt.Printf("After creation 'newDir' does NOT Exist! \n"+
      "os.Stat(newDir) newDir='%v' Error='%v'", newDir, err.Error())
    return
  }

  fmt.Println("********** mainTest53CreateNewDir **********")
  fmt.Println("")
  fmt.Println("           SUCCESS!!!           ")
  fmt.Println("Created New Directory: ", newDir)
}

func mainTest52() {
  // /d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops
  // D:\gowork\src\MikeAustin71\pathfileopsgo\pathfileops

  fh := pf.FileHelper{}

  rawPath1 := "/d/gowork/src/MikeAustin71/pathfileopsgo/pathfileops"

  rawPath2 := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\pathfileops"

  rawPath3 := "D:\\"

  vol1 := fh.GetVolumeName(rawPath1)

  fmt.Println("************ mainTest51 ***************")
  fmt.Println("rawPath1: ", rawPath1)
  fmt.Println(" Volume1: ", vol1)
  vol2 := fh.GetVolumeName(rawPath2)
  fmt.Println("----------------------------------------")
  fmt.Println("rawPath2: ", rawPath2)
  fmt.Println(" Volume2: ", vol2)
  fmt.Println("----------------------------------------")
  vol3 := fh.GetVolumeName(rawPath3)
  fmt.Println("rawPath3: ", rawPath3)
  fmt.Println(" Volume2: ", vol3)

}

func mainTest51() {

  fh := pf.FileHelper{}
  rawPathFile := fh.AdjustPathSlash("...")

  pathFile, err := fh.MakeAbsolutePath(rawPathFile)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath(rawPathFile). "+
      "rawPathFile='%v' Error='%v' ", rawPathFile, err.Error())

    return
  }

  expectedPathFile := strings.ToLower(rawPathFile)

  pathFileType, absolutePath, err := fh.IsPathFileString(rawPathFile)

  if err != nil {
    fmt.Printf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
  }

  expectedPathFileCode := pf.PathFileType.Path()

  if expectedPathFileCode != pathFileType {
    fmt.Printf("Expected PathFileTypeCode='%v'. Instead, PathFileTypeCode='%v' "+
      "Expected PathFileType Code: %v  Actual PathFileTypeCode: %v\n"+
      "absolutePath='%v' \n",
      expectedPathFileCode.String(), pathFileType.String(),
      int(expectedPathFileCode), int(pathFileType), absolutePath)
    return
  }

  lwcAbsolutePath := strings.ToLower(absolutePath)

  if expectedPathFile != lwcAbsolutePath {
    fmt.Printf("Error: Expected Absolute Path='%v'\n Actual Absolute Path='%v' ",
      expectedPathFile, lwcAbsolutePath)
  }

  fmt.Println("Raw Path: ", pathFile)
  fmt.Println("Absolute Path: ", absolutePath)
  fmt.Println(" PathFileType: ", pathFileType.String())
}

func mainTest50() {

  fh := pf.FileHelper{}
  rawPathFile := fh.AdjustPathSlash("..\\filesfortest\\levelfilesfortest\\level_01_dir\\" +
    "level_02_dir\\level_03_dir\\level_3_1_test.txt")

  pathFile, err := fh.MakeAbsolutePath(rawPathFile)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath(rawPathFile). "+
      "rawPathFile='%v' Error='%v' ", rawPathFile, err.Error())

    return
  }

  expectedPathFile := strings.ToLower(pathFile)

  pathFileType, absolutePath, err := fh.IsPathFileString(pathFile)

  if err != nil {
    fmt.Printf("Error returned from fh.IsPathFileString(pathFile). "+
      "pathFile='%v' Error='%v' ", pathFile, err.Error())
    return
  }

  if pf.PathFileTypeCode(0).PathFile() != pathFileType {
    fmt.Printf("Expected PathFileTypeCode='PathFile'. Instead, PathFileTypeCode='%v' "+
      "Expected PathFileType Code: %v  Actual PathFileTypeCode: %v\n"+
      "absolutePath='%v' \n",
      pathFileType.String(), int(pf.PathFileTypeCode(0).PathFile()), int(pathFileType), absolutePath)
    return
  }

  absolutePath = strings.ToLower(absolutePath)

  if expectedPathFile != absolutePath {
    fmt.Printf("Error: Expected Absolute Path='%v'\n Actual Absolute Path='%v' ",
      expectedPathFile, absolutePath)
  }
}

func mainTest49() {
  fh := pf.FileHelper{}

  absPathDir := fh.AdjustPathSlash("D:/gowork/src/MikeAustin71/pathfileopsgo/filesfortest/" +
    "levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/level_3_1_test.txt")

  result := fh.IsAbsolutePath(absPathDir)

  fmt.Println("TestPath: ", absPathDir)
  fmt.Println("fh.IsAbsolutePath Result: ", result)

  verifyAbsPath, err := fh.MakeAbsolutePath(absPathDir)

  if err != nil {
    fmt.Printf("Error returned by fh.MakeAbsolutePath(absPathDir). "+
      "absPathDir='%v' Error='%v' ", absPathDir, err.Error())
    return
  }

  fmt.Println("-----------------------------------------------")
  fmt.Println("MakeAbsolutePath(absPathDir) Result: ", verifyAbsPath)

  return
}

func mainTest48() {

  targetDir := "D:\\T03\\WebSite_15\\Vehicle_01"

  searchParm := "*.cs"

  fh := pf.FileHelper{}

  foundfiles, err := fh.FindFilesInPath(targetDir, searchParm)

  if err != nil {
    fmt.Printf("Error returned from fh.FindFilesInPath(targetDir, searchParm). "+
      "Error='%v'\n", err.Error())
    return
  }

  lFoundFiles := len(foundfiles)

  if lFoundFiles == 0 {
    fmt.Printf("FoundFiles is a zero length string array! \n")
    return
  }

  for i := 0; i < lFoundFiles; i++ {
    fmt.Printf("%2d. Found File: %v \n", i+1, foundfiles[i])
  }

  fmt.Println()
  fmt.Println("Successful Completion!")
}

func mainTest47() {
  // D:\gowork\src\MikeAustin71\pathfileopsgo\filesfortest\basefilesfortest\.xgitignore
  expectedFileNameExt := ".xgitignore"

  fh := pf.FileHelper{}
  absPath := fh.AdjustPathSlash("D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\basefilesfortest")

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    fmt.Printf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). "+
      "absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromDirStrFileNameStr(absPath, expectedFileNameExt)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromDirStrFileNameStr(absPath, "+
      "expectedFileNameExt). "+
      "absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
    return
  }

  fmt.Println("****** mainTest47() ******")
  fmt.Println("   actual File Name: ", expectedFileNameExt)
  fmt.Println("        info Name(): ", info.Name())
  fmt.Println("fMgr File Extension: ", fMgr.GetFileExt())
  fmt.Println("     fMgr File Name: ", fMgr.GetFileName())
  fmt.Println(" fMgr File Name Ext: ", fMgr.GetFileNameExt())
  fmt.Println()

}

func mainTest46() {
  expectedFileNameExt := "newerFileForTest_01.txt"

  fh := pf.FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    fmt.Printf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
    return
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    fmt.Printf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). "+
      "absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromDirStrFileNameStr(absPath, expectedFileNameExt)

  fmt.Println("****** mainTest46() ******")
  fmt.Println("   actual File Name: ", expectedFileNameExt)
  fmt.Println("        info Name(): ", info.Name())
  fmt.Println("fMgr File Extension: ", fMgr.GetFileExt())
  fmt.Println("     fMgr File Name: ", fMgr.GetFileName())
  fmt.Println(" fMgr File Name Ext: ", fMgr.GetFileNameExt())
  fmt.Println()
}

func mainTest45() {
  //filesfortest/basefilesfortest/basefilenoext

  expectedFileNameExt := "basefilenoext"

  fh := pf.FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/basefilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    fmt.Printf("Error returned from fh.MakeAbsolutePath(adjustedPath). "+
      "adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
    return
  }

  absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

  info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

  if err != nil {
    fmt.Printf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). "+
      "absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromDirStrFileNameStr(absPath, expectedFileNameExt)
  fmt.Println()
  fmt.Println("****** mainTest45() ******")
  fmt.Println("   actual File Name: ", expectedFileNameExt)
  fmt.Println("        info Name(): ", info.Name())
  fmt.Println("fMgr File Extension: ", fMgr.GetFileExt())
  fmt.Println("     fMgr File Name: ", fMgr.GetFileName())
  fmt.Println(" fMgr File Name Ext: ", fMgr.GetFileNameExt())
  fmt.Println()
}

func mainTest44() {

  expectedFileNameExt := "$!#%^&*()_=.t%t"

  rawPath := "../filesfortest/newfilesfortest"

  _, err := pf.FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

  if err == nil {
    fmt.Println("Expected error return from FileMgr{}." +
      "NewFromDirStrFileNameStr(rawPath, expectedFileNameExt) because " +
      "expectedFileNameExt consists of invalid characters. " +
      "However, NO ERROR WAS RETURNED!")
    return
  }

  fmt.Println()
  fmt.Println("Completed without error.")
  fmt.Println()
}

func mainTest43() {

  fh := pf.FileHelper{}
  adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

  absPath, err := fh.MakeAbsolutePath(adjustedPath)

  if err != nil {
    fmt.Printf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  "+
      "Error='%v'", adjustedPath, err.Error())
    return
  }

  var info os.FileInfo

  _, err = pf.FileMgr{}.NewFromFileInfo(absPath, info)

  if err == nil {
    fmt.Println("Expected an error from FileMgr{}.NewFromFileInfo(absPath, info) because " +
      "input parameter 'info' is INVALID!  However, NO ERROR WAS RETURNED!")

    return
  }

}

func mainTest42() {
  fh := pf.FileHelper{}
  testFile := fh.AdjustPathSlash("../iDoNotExist/TestFile011.txt")
  fileMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(testFile)

  if err != nil {
    fmt.Printf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
    return
  }

  err = fileMgr.CreateThisFile()

  if err == nil {
    fmt.Printf("Expected error return from fileMgr.CreateThisFile() because " +
      "the fileMgr directory does NOT exist. However, NO ERROR WAS RETURNED!")
    return
  }

}

func mainTest41() {

  fh := pf.FileHelper{}

  targetFile := fh.AdjustPathSlash("../filesfortest/levelfilesfortest/level_01_dir/level_1_2_test.txt")

  srcFMgr, err := pf.FileMgr{}.New(targetFile)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.New(targetFile). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  timeFormatSpec := "2006-01-02 15:04:05 -0700 MST"

  modTime, err := srcFMgr.GetFileModTime()

  if err != nil {
    fmt.Printf("Error returned from srcFMgr.GetFileModTime(). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  expectedFileModTimeStr := modTime.Format(timeFormatSpec)

  modTimeStr, err := srcFMgr.GetFileModTimeStr("xx-xx-xxxx xx:xx:xx")

  if err != nil {
    fmt.Printf("Error returned from srcFMgr.GetFileModTimeStr(\"\"). "+
      "targetFile='%v' Error='%v'\n", targetFile, err.Error())
    return
  }

  fmt.Println("  mainTest41() - Testing FileMgr.GetFileModTimeStr()")
  fmt.Println("-------------------------------------------------------")
  fmt.Println("       File Name: ", srcFMgr.GetAbsolutePathFileName())
  fmt.Println("  File Mode Time: ", modTimeStr)
  fmt.Println("Expected ModTime: ", expectedFileModTimeStr)

}

func mainTest40() {
  fh := pf.FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
  targetFile := "newerFileForTest_01"

  dirMgr, err := pf.DirMgr{}.New(targetDir)

  if err != nil {
    fmt.Printf("Error returned from DirMgr{}.New(targetDir). "+
      "targetDir='%v' Error='%v'", targetDir, err.Error())
    return
  }

  srcFMgr, err := pf.FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

  if err != nil {
    fmt.Printf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
      "DirMgr='%v' targetFile='%v' Error='%v'",
      dirMgr.GetAbsolutePath(), targetFile, err.Error())
  }

  isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

  if isFileNamePopulated {

    fmt.Printf("Expected srcFMgr.IsFileNameExtPopulated() == 'false'. Instead, it is 'true'. \n"+
      "FileName='%v' File Extension='%v' Len File Ext= '%v' ",
      srcFMgr.GetFileName(), srcFMgr.GetFileExt(), len(srcFMgr.GetFileExt()))
  }

}

func mainTest39() {

  filePath := "../filesfortest/modefilesfortest/modeFileTest_01.txt"

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  originalPermCode, err := fMgr.GetFilePermissionTextCodes()

  if err != nil {
    fmt.Printf("Error returned from fMgr.GetFilePermissionTextCodes(). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  fmt.Println("**** mainTest39 ****")
  fmt.Println("-------------------------------------------------------------------")
  fmt.Println("                     File Name: ", fMgr.GetAbsolutePathFileName())
  fmt.Println("      Original Permission Code: ", originalPermCode)

  newPerm, err := pf.FilePermissionConfig{}.New("-r--r--r--")

  if err != nil {
    fmt.Printf("Error returned from FilePermissionConfig{}.New(\"-r--r--r--\"). "+
      "Error='%v'", err.Error())
    return
  }

  err = fMgr.ChangePermissionMode(newPerm)

  if err != nil {
    fmt.Printf("Error returned from fMgr.ChangePermissionMode(newPerm). "+
      "Error='%v'", err.Error())
    return
  }

  newActualPermCode, err := fMgr.GetFilePermissionTextCodes()

  if err != nil {
    fmt.Printf("Error returned from #2 fMgr.GetFilePermissionTextCodes(). "+
      "Error='%v'", err.Error())
    return
  }

  newPermText, _ := newPerm.GetPermissionTextCode()

  fmt.Println(" New Requested Permission Code: ", newPermText)
  fmt.Println("    New Actual Permission Code: ", newActualPermCode)

  newPerm, err = pf.FilePermissionConfig{}.New("-rw-rw-rw-")

  err = fMgr.ChangePermissionMode(newPerm)

  if err != nil {
    fmt.Printf("Error returned from #2 fMgr.ChangePermissionMode(newPerm). "+
      "Error='%v'", err.Error())
    return
  }

  newPermText, _ = newPerm.GetPermissionTextCode()

  newActualPermCode, err = fMgr.GetFilePermissionTextCodes()

  if err != nil {
    fmt.Printf("Error returned from #3 fMgr.GetFilePermissionTextCodes(). "+
      "Error='%v'", err.Error())
    return
  }

  fmt.Println("Last Requested Permission Code: ", newPermText)
  fmt.Println("   Last Actual Permission Code: ", newActualPermCode)

  fmt.Println("-------------------------------------------------------------------")
  fmt.Println("                     SUCCESS!!")
}

func mainTest38() {

  testText := "Damn the torpedoes full speed account!\n"
  rawFilePath := "D:\\T08\\testWriteFile801294.txt"
  lenTestText := len(testText)

  fh := pf.FileHelper{}

  filePath := fh.AdjustPathSlash(rawFilePath)

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnlyAppend()

  if err != nil {
    _ = fMgr.CloseThisFile()
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnlyAppend(). Error='%v' ",
      err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    _ = fMgr.CloseThisFile()
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().")
    return
  }

  fmt.Println()
  fmt.Println("    mainTest38()   ")
  fmt.Println("***** SUCCESS *****")
  fmt.Printf("              Test Text: %v", testText)
  fmt.Println("    Length of Test Text: ", lenTestText)
  fmt.Println("Number of Bytes Written: ", numBytesWritten)

}

func mainTest37() {

  testText := "Now is the time for all good men to come to the aid of their country.\n"
  rawFilePath := "D:\\T08\\testWriteFile801294.txt"
  lenTestText := len(testText)

  fh := pf.FileHelper{}

  filePath := fh.AdjustPathSlash(rawFilePath)

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(filePath). "+
      "filePathName='%v'  Error='%v'", filePath, err.Error())
    return
  }

  err = fMgr.CreateThisFile()

  if err != nil {
    _ = fMgr.CloseThisFile()
    fmt.Printf("Error returned by fMgr.CreateThisFile(). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.OpenThisFileWriteOnly()

  if err != nil {
    _ = fMgr.CloseThisFile()
    fmt.Printf("Error returned by fMgr.OpenThisFileWriteOnly(). Error='%v' ",
      err.Error())
    return
  }

  bytesToWrite := []byte(testText)

  numBytesWritten, err := fMgr.WriteBytesToFile(bytesToWrite)

  if err != nil {
    _ = fMgr.CloseThisFile()
    fmt.Printf("Error returned by fMgr.WriteBytesToFile(bytesToWrite). Error='%v' ",
      err.Error())
    return
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by #1 fMgr.CloseThisFile().")
    return
  }

  fmt.Println()
  fmt.Println("    mainTest37()   ")
  fmt.Println("***** SUCCESS *****")
  fmt.Printf("              Test Text: %v", testText)
  fmt.Println("    Length of Test Text: ", lenTestText)
  fmt.Println("Number of Bytes Written: ", numBytesWritten)

}

func mainTest36() {

  expectedStr := "Thank you, for your support."

  fh := pf.FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  bytes := make([]byte, 0, 1000)

  for i := 0; i < 4; i++ {

    bytes, err = fMgr.ReadFileLine(delim)

    if err != nil &&
      err != io.EOF {
      fmt.Printf("Error returned by fMgr.ReadFileLine(delim) on Line#1. "+
        "Error='%v'", err.Error())
    }

  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  actualStr := string(bytes)

  actualStr = strings.Replace(actualStr, "\r\n", "", -1)

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  isExpectedEqualActual := false

  if actualStr == expectedStr {
    isExpectedEqualActual = true
  }

  fmt.Println("           ExpectedStr: ", expectedStr)
  fmt.Println("             ActualStr: ", actualStr)
  fmt.Println("ExpectedStr==ActualStr: ", isExpectedEqualActual)
  fmt.Println("         error==io.EOF: ", isErrEOF)

  if isExpectedEqualActual {
    fmt.Println()
    fmt.Println("********** SUCCESS **********")
  }
}

func mainTest35() {

  expectedStr := "Now is the time for all good men"

  fh := pf.FileHelper{}

  filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
  }

  delim := byte('\n')

  bytes, err := fMgr.ReadFileLine(delim)

  if err != nil {
    fmt.Printf("Error returned by fMgr.ReadFileLine(delim) on Line#1. "+
      "Error='%v'", err.Error())
  }

  actualStr := string(bytes)

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned by fMgr.CloseThisFile(). Error='%v'",
      err.Error())
  }

  fmt.Println("Length Expected String: ", len(expectedStr))
  fmt.Println("  Length Actual String: ", len(actualStr))
  expectedStr2 := "!" + expectedStr + "!"
  expectedStr2 = strings.Replace(expectedStr2, "\r", "*", -1)
  expectedStr2 = strings.Replace(expectedStr2, "\n", "%", -1)

  actualStr2 := "!" + actualStr + "!"
  actualStr2 = strings.Replace(actualStr2, "\r", "*", -1)
  actualStr2 = strings.Replace(actualStr2, "\n", "%", -1)

  fmt.Println("        Expected String: ", expectedStr2)
  fmt.Println("          Actual String: ", actualStr2)

  if expectedStr2 == actualStr2 {
    fmt.Println("************ Success **************")
  }

}

func mainTest34() {

  textCode := "-rwxrwxrwx"

  fpCfg, err := pf.FilePermissionConfig{}.New(textCode)

  if err != nil {
    fmt.Printf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  octalCode := fpCfg.GetPermissionFileModeValueText()

  fmt.Println("       Permission Codes")
  fmt.Println("----------------------------------")
  fmt.Println("Original Text Code: ", textCode)
  fmt.Printf("Octal Code: %s", octalCode)

}

func mainTest33() {

  fOpenCfg, err := pf.FileOpenConfig{}.New(
    pf.FOpenType.TypeReadWrite(),
    pf.FOpenMode.ModeCreate(),
    pf.FOpenMode.ModeExclusive())

  if err != nil {
    fmt.Printf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
    return
  }

  openCodes := fOpenCfg.GetFileOpenNarrativeText()

  fmt.Println()
  fmt.Println("FileOpenConfig{}.GetFileOpenNarrativeText()")
  fmt.Println("Open Codes: ", openCodes)
}

func mainTest32() {

  textCode := "-rwxrwxrwx"

  fpCfg, err := pf.FilePermissionConfig{}.New(textCode)

  if err != nil {
    fmt.Printf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  narrativeCode := fpCfg.GetPermissionNarrativeText()

  fmt.Println()
  fmt.Println("Test FilePermissionConfig{}.GetPermissionNarrativeText()")
  fmt.Println("          textCode: ", textCode)
  fmt.Println("narrativeText Code: ", narrativeCode)
}

func mainTest31() {

  fOpStatus1 := pf.FileOpenConfig{}

  fOpStatus2 := pf.FileOpenConfig{}

  fOpStatus2.CopyIn(&fOpStatus1)

  if !fOpStatus1.Equal(&fOpStatus2) {
    fmt.Println("Error: Expected fOpStatus1==fOpStatus2. THEY ARE NOT EQUAL!")
  } else {
    fmt.Println("Successful Completion!")
  }
}

func mainTest30() {

  rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest"

  dMgr, err := pf.DirMgr{}.New(rawPath)

  if err != nil {
    fmt.Printf("Error returned by DirMgr{}.New(rawPath). "+
      "rawPath='%v' Error='%v' \n\n", rawPath, err.Error())
    return
  }

  permissionText, err := dMgr.GetDirPermissionTextCodes()

  if err != nil {
    fmt.Printf("Error returned by dMgr.GetDirPermissionTextCodes(). "+
      "Error='%v' \n", err.Error())
    return
  }

  fmt.Println("  Directory: ", dMgr.GetAbsolutePath())
  fmt.Println("Permissions: ", permissionText)

  fInfoPlus, err := dMgr.GetFileInfoPlus()

  if err != nil {
    fmt.Printf("Error returned by dMgr.GetFileInfoPlus(). "+
      "Error='%v' \n", err.Error())
    return
  }

  fmt.Println("    FInfo Mode Str: ", fInfoPlus.Mode().String())
  fmt.Println("Finfo Mode IsDir():", fInfoPlus.Mode().IsDir())
}

func mainTest29() {

  rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_3_test.txt"

  fileMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(rawPath)

  if err != nil {
    fmt.Printf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(rawPath). "+
      "rawPath='%v' Error='%v' \n\n", rawPath, err.Error())
    return
  }

  permissionText, err := fileMgr.GetFilePermissionTextCodes()

  if err != nil {
    fmt.Printf("Error returned by fileMgr.GetFilePermissionTextCodes(). "+
      "Error='%v' \n", err.Error())
    return
  }

  fmt.Println("File: ", fileMgr.GetAbsolutePathFileName())
  fmt.Println("Permissions: ", permissionText)
}

func mainTest28() {

  unixPermissions := os.FileMode(0777)

  fmt.Printf("Unix Permissions Decimal Value: %s\n",
    strconv.FormatInt(int64(unixPermissions), 10))

  fmt.Printf("Unix Permissions Octal Value: %s\n",
    strconv.FormatInt(int64(unixPermissions), 8))

  intUnixPermissions := int(unixPermissions)

  fmt.Printf("intUnixPermissions Decimal Value: %s\n",
    strconv.FormatInt(int64(intUnixPermissions), 10))

  fmt.Printf("intUnixPermissions Octal Value: %s\n",
    strconv.FormatInt(int64(intUnixPermissions), 8))

  dir := os.ModeDir | os.FileMode(0333)

  fmt.Println()
  fmt.Println("------------------------------------")
  fmt.Printf("dir Decimal Value: %s\n",
    strconv.FormatInt(int64(dir), 10))

  fmt.Printf("dir Octal Value: %s\n",
    strconv.FormatInt(int64(dir), 8))

  baseDir := dir &^ os.FileMode(0777)

  fmt.Println()
  fmt.Println("------------------------------------")
  fmt.Printf("baseDir Decimal Value: %s\n",
    strconv.FormatInt(int64(baseDir), 10))

  fmt.Printf("baseDir Octal Value: %s\n",
    strconv.FormatInt(int64(baseDir), 8))

}

func mainTest27() {

  unixPermissions := os.FileMode(0777)

  shift := uint(9)

  xPermissions := unixPermissions >> shift

  fmt.Printf("Unix Permissions Decimal Value: %s\n",
    strconv.FormatInt(int64(unixPermissions), 10))

  fmt.Printf("Unix Permissions Octal Value: %s\n",
    strconv.FormatInt(int64(unixPermissions), 8))

  fmt.Printf("shift= %d \n", shift)

  fmt.Printf("xPermissions Decimal Value: %s\n",
    strconv.FormatInt(int64(xPermissions), 10))

  fmt.Printf("xPermissions Octal Value: %s\n",
    strconv.FormatInt(int64(xPermissions), 8))

}

func mainTest26() {

  unixPermissions := os.FileMode(0777)

  fmt.Printf("unixPermissions Decimal: %d  Octal: %s Binary: %s String Value:%x \n",
    unixPermissions, strconv.FormatInt(int64(unixPermissions), 8),
    strconv.FormatInt(int64(unixPermissions), 2), unixPermissions.String())

  fmt.Printf("ModeDir Decimal: %d  Octal: %s Binary: %s String Value: %s \n",
    os.ModeDir, strconv.FormatInt(int64(os.ModeDir), 8),
    strconv.FormatInt(int64(os.ModeDir), 2), os.ModeDir.String())

  directoryPermission := os.ModeDir | unixPermissions

  fmt.Printf("directoryPermission Decimal: %d  Octal: %s Binary: %s \n",
    directoryPermission, strconv.FormatInt(int64(directoryPermission), 8),
    strconv.FormatInt(int64(directoryPermission), 2))
  fmt.Println("-----------------------------------------------------------------------")
  fmt.Printf("   dirctoryPermission string value: %s\n", directoryPermission.String())
  fmt.Printf("    dirctoryPermission IsDir value: %v\n", directoryPermission.IsDir())
  fmt.Printf("dirctoryPermission IsRegular value: %v\n", directoryPermission.IsRegular())

}

func mainTest25() {

  fmt.Printf("                        OS MODE CONSTANTS \n")

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeDir Decimal Value: %s  \n  Octal Value : %s  \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeDir), 10),
    strconv.FormatInt(int64(os.ModeDir), 8),
    strconv.FormatInt(int64(os.ModeDir), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeAppend Decimal Value: %s  \n  Octal Value: %s  \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeAppend), 10),
    strconv.FormatInt(int64(os.ModeAppend), 8),
    strconv.FormatInt(int64(os.ModeAppend), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeExclusive Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeExclusive), 10),
    strconv.FormatInt(int64(os.ModeExclusive), 8),
    strconv.FormatInt(int64(os.ModeExclusive), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeTemporary Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeTemporary), 10),
    strconv.FormatInt(int64(os.ModeTemporary), 8),
    strconv.FormatInt(int64(os.ModeTemporary), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeSymlink Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeSymlink), 10),
    strconv.FormatInt(int64(os.ModeSymlink), 8),
    strconv.FormatInt(int64(os.ModeSymlink), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeDevice Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeDevice), 10),
    strconv.FormatInt(int64(os.ModeDevice), 8),
    strconv.FormatInt(int64(os.ModeDevice), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeNamedPipe Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeNamedPipe), 10),
    strconv.FormatInt(int64(os.ModeNamedPipe), 8),
    strconv.FormatInt(int64(os.ModeNamedPipe), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeSocket Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeSocket), 10),
    strconv.FormatInt(int64(os.ModeSocket), 8),
    strconv.FormatInt(int64(os.ModeSocket), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeSetuid Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeSetuid), 10),
    strconv.FormatInt(int64(os.ModeSetuid), 8),
    strconv.FormatInt(int64(os.ModeSetuid), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeSetgid Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeSetgid), 10),
    strconv.FormatInt(int64(os.ModeSetgid), 8),
    strconv.FormatInt(int64(os.ModeSetgid), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeCharDevice Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeCharDevice), 10),
    strconv.FormatInt(int64(os.ModeCharDevice), 8),
    strconv.FormatInt(int64(os.ModeCharDevice), 2))

  fmt.Printf("ModeSticky Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeSticky), 10),
    strconv.FormatInt(int64(os.ModeSticky), 8),
    strconv.FormatInt(int64(os.ModeSticky), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

  fmt.Printf("ModeIrregular Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
    strconv.FormatInt(int64(os.ModeIrregular), 10),
    strconv.FormatInt(int64(os.ModeIrregular), 8),
    strconv.FormatInt(int64(os.ModeIrregular), 2))

  fmt.Printf("-----------------------------------------------------------------------------------\n\n")

}

func mainTest24() {

  fPerm := pf.FilePermissionConfig{}

  expectedDecimalModeValue := 511
  expectedOctalModeValue :=
    pf.FileHelper{}.ConvertDecimalToOctal(expectedDecimalModeValue)

  modeStr := "-rwxrwxrwx"

  err := fPerm.SetFileModeByTextCode(modeStr)

  if err != nil {
    fmt.Printf("Error returned by fPerm.StringToMode(modeStr). "+
      "modeStr='%v' Error='%v' \n", modeStr, err.Error())
    return
  }

  fMode, err := fPerm.GetCompositePermissionMode()

  if err != nil {
    fmt.Printf("Error returned by fPerm.GetCompositePermissionMode(). "+
      "modeStr='%v' Error='%v' \n", modeStr, err.Error())
    return
  }

  actualDecimalModeValue := int(fMode)

  actualOctalModeValue :=
    pf.FileHelper{}.ConvertDecimalToOctal(actualDecimalModeValue)

  fmt.Println("--- StringToMode Results ----")
  fmt.Println("          Original Mode Str: ", modeStr)
  fmt.Println("            Actual Mode Str: ", fMode.String())
  fmt.Println("Expected Decimal Mode Value: ", expectedDecimalModeValue)
  fmt.Println("  Actual Mode Decimal Value: ", actualDecimalModeValue)
  fmt.Println("  Expected Octal Mode Value: ", expectedOctalModeValue)
  fmt.Println("    Actual Octal Mode Value: ", actualOctalModeValue)
}

func mainTest23() {

  dfm := os.FileMode(os.ModeDir)

  fmt.Println()
  fmt.Println("-------------------------------------------------------------")
  fmt.Printf("File Mode String %s\n", dfm.String())
  tfm := os.FileMode(0777)
  fmt.Printf("4-digit 777 File Mode String %s\n", tfm.String())

  nfm := tfm | dfm

  fmt.Printf("tfm or'd with dfm  %s \n", nfm.String())
  fh := pf.FileHelper{}

  mode := fh.ConvertOctalToDecimal(777)
  fmt.Printf("mode = %d\n", mode)
  decimalEquivalent := fh.ConvertDecimalToOctal(mode)
  tfm = os.FileMode(mode)
  fmt.Printf("3-digit 777 File Mode String %s\n", tfm.String())
  fmt.Printf("Decimal Equivalent %d \n", decimalEquivalent)
}

func mainTest22() {

  expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

  fOpStatus, err := pf.FileOpenConfig{}.New(pf.FOpenType.TypeWriteOnly(),
    pf.FOpenMode.ModeAppend(), pf.FOpenMode.ModeTruncate())

  if err != nil {
    fmt.Printf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
    return
  }

  actualFileOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    fmt.Printf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode(). Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFileOpenCode {
    fmt.Printf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFileOpenCode)
    return
  }

  fmt.Printf("Success - File Open Codes Match!")
  return
}

func mainTest21() {

  fmt.Println("--------- Primary Codes ---------")
  fmt.Println("os.O_RDONLY: ", os.O_RDONLY)
  fmt.Println("os.O_WRONLY: ", os.O_WRONLY)
  fmt.Println("os.O_RDWR: ", os.O_RDWR)
  fmt.Println()
  fmt.Println("--------- Control Codes ----------")
  fmt.Println("os.O_APPEND: ", os.O_APPEND)
  fmt.Println("os.O_CREATE: ", os.O_CREATE)
  fmt.Println("os.O_EXCL: ", os.O_EXCL)
  fmt.Println("os.O_SYNC: ", os.O_SYNC)
  fmt.Println("os.O_TRUNC: ", os.O_TRUNC)

}

func mainTest20() {

  fh := pf.FileHelper{}

  relPath := "../testfiles"
  origPath := fh.AdjustPathSlash(relPath)

  origAbsPath, err := fh.MakeAbsolutePath(origPath)

  if err != nil {
    fmt.Printf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'", origPath, err.Error())
    return
  }

  testDMgr, err := pf.DirMgr{}.New(origAbsPath)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(origAbsPath). "+
      "origAbsPath= '%v'  Error='%v'", origAbsPath, err.Error())
    return
  }

  var fileNameExt string

  fMgrs1 := pf.FileMgrCollection{}

  for i := 0; i < 10; i++ {

    fileNameExt = fmt.Sprintf(testDMgr.GetAbsolutePathWithSeparator()+"testAddFile_%03d.txt", i+1)

    fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(fileNameExt)

    if err != nil {
      fmt.Printf("Error returned by pf.FileMgr{}.NewFromPathFileNameExtStr(fileNameExt). "+
        "fileNameExt='%v' Error='%v' ", fileNameExt, err.Error())
      return
    }

    fMgrs1.AddFileMgr(fMgr)
  }

  if fMgrs1.GetNumOfFileMgrs() != 10 {
    fmt.Printf("Expected fMgrs1 Array Length == 10. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
    return
  }

  origPath = fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

  origAbsPath, err = fh.MakeAbsolutePath(origPath)

  if err != nil {
    fmt.Printf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
      "origPath= '%v'  Error='%v'\n", origPath, err.Error())
    return
  }

  insertedFMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

  if err != nil {
    fmt.Printf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
      "origAbsPath='%v' \nError='%v' \n", origAbsPath, err.Error())
    return
  }

  err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5)

  if err != nil {
    fmt.Printf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
      "Error='%v' \n", err.Error())
    return
  }

  numOfFileMgrs := fMgrs1.GetNumOfFileMgrs()

  for i := 0; i < numOfFileMgrs; i++ {

    xFmgr, err := fMgrs1.PeekFileMgrAtIndex(i)

    if err != nil {
      fmt.Printf("Error returned by fMgrs1.PeekFileMgrAtIndex(i). "+
        "i='%v' Error='%v' \n", i, err.Error())
    }

    fmt.Printf("i='%v' xFmgr='%v' \n", i, xFmgr.GetAbsolutePathFileName())
  }

  if fMgrs1.GetNumOfFileMgrs() != 11 {
    fmt.Printf("After insertion, expected fMgrs1 Array Length == 11. "+
      "Instead fMgrs1.GetNumOfDirs()=='%v'\n", fMgrs1.GetNumOfFileMgrs())
    return
  }

  fMgr5, err := fMgrs1.PeekFileMgrAtIndex(5)

  if err != nil {
    fmt.Printf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' \n", err.Error())
    return
  }

  if !insertedFMgr.Equal(&fMgr5) {
    fmt.Printf("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!\n")
  }

}
