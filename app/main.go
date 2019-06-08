package main

import (
  pf "../pathfileops"
  "fmt"
  "io"
)

/*


import (
  pf "../pathfileops"
  "fmt"
)


*/

func main() {

  mainTest66ReadFileBytes()

}

func mainTest66ReadFileBytes() {

  fh := pf.FileHelper{}

  setupFile := fh.AdjustPathSlash(
    "../filesfortest/checkfiles03/testRead2008.txt")

  filePath := fh.AdjustPathSlash(
    "../checkfiles/checkfiles03/testRead2008.txt")

  err := fh.DeleteDirFile(filePath)

  if err != nil {
    fmt.Printf("Test Setup Error returned by fh.DeleteDirFile(filePath)\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  err = fh.CopyFileByIo(setupFile, filePath)

  if err != nil {
    fmt.Printf("Test Setup Error returned by fh.CopyFileByIo(setupFile, filePath)\n" +
      "setupFile='%v'\nfilePath='%v'\n",
      setupFile, filePath)
    return
  }

  fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

  if err != nil {
    fmt.Printf("Error returned from common.FileMgr{}." +
      "NewFromPathFileNameExtStr(filePath).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  byteBuff := make([]byte, 2048, 2048)

  bytesRead, err := fMgr.ReadFileBytes(byteBuff)

  if err != nil &&
    err != io.EOF {
    fmt.Printf("Error returned from fMgr.ReadFileBytes(byteBuff).\n"+
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    _ = fMgr.CloseThisFile()
    return
  } else if err != nil {
    fmt.Printf("Error return from ReadFileBytes:\n" +
      "%v", err.Error())
    return
  } else if err == nil {
    fmt.Printf("Error return from ReadFileBytes is 'nil'!\n\n")
  }

  isErrEOF := false

  if err == io.EOF {
    isErrEOF = true
  }

  if !isErrEOF {
    fmt.Print("ERROR: Expected the last error return from fMgr.ReadFileBytes(byteBuff)\n" +
      "to be io.EOF.\n" +
      "Instead, error WAS NOT equal to io.EOF!\n")
  }

  var rStr = make([]rune, 0, 2048)

  for i := 0; i < len(byteBuff); i++ {

    if byteBuff[i] == 0 {
      break
    }

    rStr = append(rStr, rune(byteBuff[i]))
  }

  expectedStr :=
    "Test Read File. Do NOT alter the contents of this file."
  actualStr := string(rStr)

  if expectedStr != actualStr {
    fmt.Printf("Expected Read String='%v'.\n" +
      "Instead, Actual Read String='%v'\n",
      expectedStr, actualStr)
  }

  expectedBytesRead := len(expectedStr)

  if expectedBytesRead != bytesRead {
    fmt.Printf("Expected Bytes Read='%v'.\n" +
      "Instead, Actual Bytes Read='%v'\n",
      expectedBytesRead, bytesRead)
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    fmt.Printf("Error returned from fMgr.CloseThisFile()\n" +
      "filePath='%v'\nError='%v'\n",
      filePath, err.Error())
    return
  }

  if fMgr.GetFilePtr() != nil {
    fmt.Print("ERROR: After fMgr.CloseThisFile() expected " +
      "fMgr.filePtr==nil.\n" +
      "fMgr.filePtr IS NOT EQUAL TO NIL!\n")
    return
  }

  fmt.Println("              mainTest66ReadFileBytes                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}
