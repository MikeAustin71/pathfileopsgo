package pathfileops

import (
  "errors"
  "io"
  "os"
  "strings"
  "testing"
  "time"
)

func TestFileHelper_OpenFileReadOnly_01(t *testing.T) {

  fh := FileHelper{}

  source := "../logTest/topTest2.txt"
  source = fh.AdjustPathSlash(alogtopTest2Text)

  target := "../checkfiles/TestFileHelper_OpenFileReadOnly_01.txt"

  target = fh.AdjustPathSlash(target)

  expected := "Top level test file # 2."

  if fh.DoesFileExist(target) {

    err := fh.DeleteDirFile(target)

    if err != nil {
      t.Errorf("Test Setup Error: Attempted deletion of preexisting "+
        "target file FAILED!\ntargetFile='%v'\nError='%v'\n",
        target, err.Error())
      return
    }

    if fh.DoesFileExist(target) {
      t.Errorf("Test Setup Error: Verification of target file deletion FAILED!\n"+
        "Target File still exists after attempted deletion!\ntargetFile='%v'\n",
        target)
      return
    }
  }

  err := fh.CopyFileByIo(source, target)

  if err != nil {
    t.Errorf("Test Setup Error: Copy of source file to target file FAILED!\n"+
      "sourceFile='%v'\ntargetFile='%v'\nError='%v'\n",
      source, target, err.Error())
    return
  }

  f, err := fh.OpenFileReadOnly(target)

  if err != nil {
    t.Errorf("Failed to open file: '%v'\nError='%v'",
      target, err.Error())
    return
  }

  bytes := make([]byte, 500)

  bytesRead, err := f.Read(bytes)

  if err != nil {
    t.Errorf("Error returned from f.Read(bytes).\n"+
      "targetFile='%v'\nError='%v'\n", target, err.Error())
    _ = f.Close()
    _ = fh.DeleteDirFile(target)
    return
  }

  s := string(bytes[0:bytesRead])

  if expected != s {
    t.Errorf("Expected read string='%v'. Instead read string='%v'",
      expected, s)
  }

  _ = f.Close()
  _ = fh.DeleteDirFile(target)
}

func TestFileHelper_OpenFileReadOnly_02(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.OpenFileReadOnly("")

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(\"\") " +
      "because the input parameter is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_OpenFileReadOnly_03(t *testing.T) {

  fh := FileHelper{}

  _, err := fh.OpenFileReadOnly("    ")

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(\"\") " +
      "because the input parameter consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_OpenFileReadOnly_04(t *testing.T) {

  fh := FileHelper{}

  targetFile := "../filesfortest/levelfilesfortest/iDoNotExist.txt"

  targetFile = fh.AdjustPathSlash(targetFile)

  _, err := fh.OpenFileReadOnly(targetFile)

  if err == nil {
    t.Error("Expected an error from fh.OpenFileReadOnly(targetFile) " +
      "because the input parameter 'targetFile' does not exist.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_OpenFileReadOnly_05(t *testing.T) {

  fh := FileHelper{}

  source := "../logTest/topTest2.txt"
  source = fh.AdjustPathSlash(alogtopTest2Text)

  target := "../checkfiles/TestFileHelper_OpenFileReadOnly_01.txt"

  target = fh.AdjustPathSlash(target)

  if fh.DoesFileExist(target) {

    err := fh.DeleteDirFile(target)

    if err != nil {
      t.Errorf("Test Setup Error: Attempted deletion of preexisting "+
        "target file FAILED!\ntargetFile='%v'\nError='%v'\n",
        target, err.Error())
      return
    }

    if fh.DoesFileExist(target) {
      t.Errorf("Test Setup Error: Verification of target file deletion FAILED!\n"+
        "Target File still exists after attempted deletion!\ntargetFile='%v'\n",
        target)
      return
    }
  }

  err := fh.CopyFileByIo(source, target)

  if err != nil {
    t.Errorf("Test Setup Error: Copy of source file to target file FAILED!\n"+
      "sourceFile='%v'\ntargetFile='%v'\nError='%v'\n",
      source, target, err.Error())
    return
  }

  f, err := fh.OpenFileReadOnly(target)

  if err != nil {
    t.Errorf("Failed to open file: '%v'\nError='%v'",
      target, err.Error())
    return
  }

  testText := "Cannot write text to read-only file!"

  _, err = f.WriteString(testText)

  if err == nil {
    t.Errorf("Expected an error return from f.WriteString(testText) " +
      "because\n'f' references a read-only file. However, NO ERROR WAS RETURNED!\n")
  }

  err = f.Close()

  if err != nil {
    t.Errorf("Test Clean-up Error: Error return from f.Close().\n"+
      "File Name='%v'\nError='%v'\n",
      target, err.Error())
  }

  err = fh.DeleteDirFile(target)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error return from fh.DeleteDirFile(target).\n"+
      "target='%v'\nError='%v'", target, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_01(t *testing.T) {

  fh := FileHelper{}
  targetFile := "../checkfiles/scratchTestFileHelper_OpenFileForWriting_01.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  testString := "How now, brown cow!"

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  // truncateFile == false - targetFile does not yet exist!
  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile, false)\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: File Pointer returned by fh.OpenFileReadWrite(targetFile)\n"+
      "is 'nil'!\ntargetFile='%v'", targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(testString)

  if bytesWritten != len(testString) {
    t.Errorf("ERROR: Bytes written to 'targetFile' DO NOT EQUAL the lenth\n"+
      "of 'testString'.\ntargetFile='%v'\nBytesWritten='%v' Length of Test String='%v'\n",
      targetFile, bytesWritten, len(testString))
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Sync()

  if err != nil {
    t.Errorf("Error returned by fPtr.Sync() for 'targetFile'!\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  b := make([]byte, 500)

  bytesRead, err := fPtr.ReadAt(b, 0)

  if err != nil {
    if err != io.EOF {
      t.Errorf("Non-EOF error returned by fPtr.ReadAt(b,0).\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      _ = fPtr.Close()
      _ = fh.DeleteDirFile(targetFile)
      return
    }
  }

  if bytesRead != bytesWritten {
    t.Errorf("ERROR: The bytes written to 'targetFile' do NOT EQUAL the bytes\n"+
      "read from 'targetFile'.\ntargetFile='%v'\nBytes Read='%v'  Bytes Written='%v'\n",
      targetFile, bytesRead, bytesWritten)
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  resultStr := string(b[0:bytesRead])

  if testString != resultStr {
    t.Errorf("ERROR: Expected read string='%v'.\nInstead, read string='%v'.\n",
      testString, resultStr)
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_02(t *testing.T) {

  fh := FileHelper{}
  srcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  srcFile = fh.AdjustPathSlash(srcFile)
  targetFile := "../checkfiles/scratchTestFileHelper_OpenFileForWriting_02.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  testString := "How now, brown cow!"

  err := fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Setup Error returned from fh.DeleteDirFile(targetFile).\n" +
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    return
  }


  fInfo, err := os.Stat(srcFile)

  if err != nil {
    t.Errorf("ERROR: Test Setup Source File DOES NOT EXIST!\n"+
      "Source File='%v'\n", srcFile)
    return
  }

  sourceByteSize := fInfo.Size()

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }
  }

  err = fh.CopyFileByIo(srcFile, targetFile)

  if err != nil {
    t.Errorf("Error returned by test setup op fh.CopyFileByIo(srcFile, targetFile).\n"+
      "srcFile='%v'\ntargetFile='%v'\nError='%v'\n",
      srcFile, targetFile, err.Error())
    return
  }

  if !fh.DoesFileExist(targetFile) {
    t.Errorf("Test Setup Failed! 'targetFile' does NOT EXIST!\n"+
      "targetFile='%v'\n", targetFile)
    return
  }

  // Open file with truncateFile=true
  fPtr, err := fh.OpenFileReadWrite(targetFile, true)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile)\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: File Pointer returned by fh.OpenFileReadWrite(targetFile)\n"+
      "is 'nil'!\ntargetFile='%v'", targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(testString)

  if bytesWritten != len(testString) {
    t.Errorf("ERROR: Bytes written to 'targetFile' DO NOT EQUAL the lenth\n"+
      "of 'testString'.\ntargetFile='%v'\nBytesWritten='%v' Length of Test String='%v'\n",
      targetFile, bytesWritten, len(testString))
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Sync()

  if err != nil {
    t.Errorf("Error returned by fPtr.Sync() for 'targetFile'!\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  b := make([]byte, 500)

  bytesRead, err := fPtr.ReadAt(b, 0)

  if err != nil {
    if err != io.EOF {
      t.Errorf("Non-EOF error returned by fPtr.ReadAt(b,0).\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      _ = fPtr.Close()
      _ = fh.DeleteDirFile(targetFile)
      return
    }
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned after Read Operation on fPtr.Close()!\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  fInfo, err = os.Stat(targetFile)

  if err != nil {
    t.Errorf("ERROR: os.Stat(targetFile) shows targetFile DOES NOT EXIST!\n"+
      "targetFile='%v'\n", targetFile)
    return
  }

  targetFileByteSize := fInfo.Size()

  if sourceByteSize <= targetFileByteSize {
    t.Errorf("ERROR: Orginal Source File Byte Size is less than new "+
      "'targetFile' Byte Size!\nSource File Byte Size='%v'   "+
      "Target File Byte Size='%v'\ntargetFile='%v'\n",
      sourceByteSize, targetFileByteSize, targetFile)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  if bytesRead != bytesWritten {
    t.Errorf("ERROR: The bytes written to 'targetFile' do NOT EQUAL the bytes\n"+
      "read from 'targetFile'.\ntargetFile='%v'\nBytes Read='%v'  Bytes Written='%v'\n",
      targetFile, bytesRead, bytesWritten)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  resultStr := string(b[0:bytesRead])

  if testString != resultStr {
    t.Errorf("ERROR: Expected read string='%v'.\nInstead, read string='%v'.\n",
      testString, resultStr)
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
  }

}

func TestFileHelper_OpenFileReadWrite_03(t *testing.T) {

  targetFile := ""

  fh := FileHelper{}

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because 'targetFile' is an empty string.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n"+
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }

  }

}

func TestFileHelper_OpenFileReadWrite_04(t *testing.T) {

  targetFile := "  "

  fh := FileHelper{}

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because the 'targetFile' parameter consists entirely of blank spaces.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n"+
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }
  }
}

func TestFileHelper_OpenFileReadWrite_05(t *testing.T) {

  targetFile := "../checkfiles/idontexist1/idontexist2/TestFileHelper_OpenFileReadWrite_05.txt"

  fh := FileHelper{}

  targetFile = fh.AdjustPathSlash(targetFile)

  fPtr, err := fh.OpenFileReadWrite(targetFile, false)

  if err == nil {
    t.Error("ERROR: Expected an error return from fh.OpenFileReadWrite" +
      "(targetFile, false)\n" +
      "because the 'targetFile' parameter includes parent directories which DO NOT EXIST.\n" +
      "However NO ERROR WAS RETURNED!!!\n")

    if fPtr != nil {

      err = fPtr.Close()

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fPtr.Close().\n"+
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

      err = fh.DeleteDirFile(targetFile)

      if err != nil {
        t.Errorf("Test Clean-up Error: Error returned from fh.DeleteDirFile(targetFile).\n"+
          "targetFile='%v'\nError='%v'", targetFile, err.Error())
      }

    }

  }
}

func TestFileHelper_OpenFileWriteOnly_01(t *testing.T) {
  fh := FileHelper{}
  srcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  srcFile = fh.AdjustPathSlash(srcFile)
  targetFile := "../checkfiles/TestFileHelper_OpenFileWriteOnly_01.txt"
  targetFile = fh.AdjustPathSlash(targetFile)

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  err := fh.CopyFileByIo(srcFile, targetFile)

  if err != nil {
    t.Errorf("Error returned by test setup op fh.CopyFileByIo(srcFile, targetFile).\n"+
      "srcFile='%v'\ntargetFile='%v'\nError='%v'\n",
      srcFile, targetFile, err.Error())
    return
  }

  if !fh.DoesFileExist(targetFile) {
    t.Errorf("Test Setup Failed! 'targetFile' does NOT EXIST!\n"+
      "targetFile='%v'\n", targetFile)
    return
  }

  fPtr, err := fh.OpenFileWriteOnly(targetFile, false)

  if err != nil {
    t.Errorf("Error returned from fh.OpenFileWriteOnly"+
      "(targetFile,false).\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    if fPtr != nil {
      _ = fPtr.Close()
    }

    err = fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("After OpenFileWriteOnly() an error was returned by "+
        "fh.DeleteDirFile(targetFile)\ntargetFile='%v'\nError='%v'\n",
        targetFile, err.Error())
    }

    return
  }

  bytes := make([]byte, 3000)

  _, err = fPtr.Read(bytes)

  if err == nil {
    t.Errorf("Expected an error retun from fPtr.Read(bytes) "+
      "because\nthe file pointer 'fPtr' was opened as 'Write-Only'!\n"+
      "targetFile='%v'\n", targetFile)
  }

  if fPtr != nil {
    err = fPtr.Close()
    if err != nil {
      t.Errorf("Test Clean-up Error returned by fPtr.Close().\n"+
        "targetFile='%v'\nError='%v'\n",
        targetFile, err.Error())
    }
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error returned by fh.DeleteDirFile("+
      "targetFile)\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())
  }

}

func TestFileHelper_OpenFileWriteOnly_02(t *testing.T) {
  fh := FileHelper{}
  srcFile := "../filesfortest/levelfilesfortest/level_0_3_test.txt"
  srcFile = fh.AdjustPathSlash(srcFile)
  targetFile := "../checkfiles/TestFileHelper_OpenFileWriteOnly_01.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  expectedStr := "How Now Brown Cow!"

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }

  }

  err := fh.CopyFileByIo(srcFile, targetFile)

  if err != nil {
    t.Errorf("Error returned by test setup op fh.CopyFileByIo(srcFile, targetFile).\n"+
      "srcFile='%v'\ntargetFile='%v'\nError='%v'\n",
      srcFile, targetFile, err.Error())
    return
  }

  if !fh.DoesFileExist(targetFile) {
    t.Errorf("Test Setup Failed! 'targetFile' does NOT EXIST!\n"+
      "targetFile='%v'\n", targetFile)
    return
  }

  fPtr, err := fh.OpenFileWriteOnly(targetFile, true)

  if err != nil {
    t.Errorf("Error returned from fh.OpenFileWriteOnly"+
      "(targetFile,false).\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)

    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: fh.OpenFileWriteOnly(targetFile,true)\n"+
      "returned a 'nil' file pointer!\ntargetFile='%v'\n", targetFile)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(expectedStr)

  if err != nil {
    t.Errorf("Error returned by fPtr.WriteString(expectedStr).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after writing bytes to file.\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }


  if bytesWritten != len(expectedStr) {
    t.Errorf("Expected bytes written='%v'. Instead, bytes written='%v'.",
      bytesWritten, len(expectedStr))
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  fPtr, err = fh.OpenFileReadWrite(targetFile, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile, false).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytes := make([]byte, 3000)
  bytesRead, err := fPtr.Read(bytes)

  if err != nil {
    t.Errorf("Error returned by fPtr.Read(bytes).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after bytes read operation.\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  if bytesWritten != bytesRead {
    t.Errorf("Expected bytes read='%v'. Instead, bytes read='%v'\n",
      bytesWritten, bytesRead)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  strRead := string(bytes[0:bytesRead])

  if expectedStr != strRead {
    t.Errorf("Expected read string = '%v'\n"+
      "Instead, read string='%v'\n",
      expectedStr, strRead)
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned by fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
  }

  return
}

func TestFileHelper_OpenFileWriteOnly_03(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../checkfiles/TestFileHelper_OpenFileWriteOnly_03.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  expectedStr := "Now is the time for all good men to come to the aid of their country."

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }
  }

  fPtr, err := fh.OpenFileWriteOnly(targetFile, false)

  if err != nil {
    t.Errorf("Error returned from fh.OpenFileWriteOnly"+
      "(targetFile,false).\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)

    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: fh.OpenFileWriteOnly(targetFile,true)\n"+
      "returned a 'nil' file pointer!\ntargetFile='%v'\n", targetFile)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(expectedStr)

  if err != nil {
    t.Errorf("Error returned by fPtr.WriteString(expectedStr).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after writing bytes to file.\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }


  if bytesWritten != len(expectedStr) {
    t.Errorf("Expected bytes written='%v'. Instead, bytes written='%v'.",
      bytesWritten, len(expectedStr))
    _ = fh.DeleteDirFile(targetFile)
    return
  }


  fPtr, err = fh.OpenFileReadWrite(targetFile, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile, false).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytes := make([]byte, 3000)
  bytesRead, err := fPtr.Read(bytes)

  if err != nil {
    t.Errorf("Error returned by fPtr.Read(bytes).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after bytes read operation.\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  if bytesWritten != bytesRead {
    t.Errorf("Expected bytes read='%v'. Instead, bytes read='%v'\n",
      bytesWritten, bytesRead)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  strRead := string(bytes[0:bytesRead])

  if expectedStr != strRead {
    t.Errorf("Expected read string = '%v'\n"+
      "Instead, read string='%v'\n",
      expectedStr, strRead)
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned by fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
  }

  return
}

func TestFileHelper_OpenFileWriteOnly_04(t *testing.T) {
  fh := FileHelper{}
  targetFile := "../checkfiles/TestFileHelper_OpenFileWriteOnly_03.txt"
  targetFile = fh.AdjustPathSlash(targetFile)
  expectedStr := "The cow jumped over the moon."

  if fh.DoesFileExist(targetFile) {
    err := fh.DeleteDirFile(targetFile)

    if err != nil {
      t.Errorf("ERROR: Test Setup attempted to delete 'targetFile'.\n"+
        "fh.DeleteDirFile(targetFile) returned an error!\n"+
        "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
      return
    }

    if fh.DoesFileExist(targetFile) {
      t.Errorf("ERROR: Test Setup attempted deletion of 'targetFile'.\n"+
        "'targetFile' STILL EXISTS!\n"+
        "targetFile='%v'\n", targetFile)
      return
    }
  }

  fPtr, err := fh.OpenFileWriteOnly(targetFile, false)

  if err != nil {
    t.Errorf("Error returned from fh.OpenFileWriteOnly"+
      "(targetFile,false).\ntargetFile='%v'\nError='%v'\n",
      targetFile, err.Error())

    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)

    return
  }

  if fPtr == nil {
    t.Errorf("ERROR: fh.OpenFileWriteOnly(targetFile,true)\n"+
      "returned a 'nil' file pointer!\ntargetFile='%v'\n", targetFile)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytesWritten, err := fPtr.WriteString(expectedStr)

  if err != nil {
    t.Errorf("Error returned by fPtr.WriteString(expectedStr).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after writing bytes to file.\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }


  if bytesWritten != len(expectedStr) {
    t.Errorf("Expected bytes written='%v'. Instead, bytes written='%v'.",
      bytesWritten, len(expectedStr))
    _ = fh.DeleteDirFile(targetFile)
    return
  }


  fPtr, err = fh.OpenFileReadWrite(targetFile, false)

  if err != nil {
    t.Errorf("Error returned by fh.OpenFileReadWrite(targetFile, false).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    if fPtr != nil {
      _ = fPtr.Close()
    }

    _ = fh.DeleteDirFile(targetFile)
    return
  }

  bytes := make([]byte, 3000)
  bytesRead, err := fPtr.Read(bytes)

  if err != nil {
    t.Errorf("Error returned by fPtr.Read(bytes).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
    _ = fPtr.Close()
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  err = fPtr.Close()

  if err != nil {
    t.Errorf("Error returned by fPtr.Close() after bytes read operation.\n"+
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  if bytesWritten != bytesRead {
    t.Errorf("Expected bytes read='%v'. Instead, bytes read='%v'\n",
      bytesWritten, bytesRead)
    _ = fh.DeleteDirFile(targetFile)
    return
  }

  strRead := string(bytes[0:bytesRead])

  if expectedStr != strRead {
    t.Errorf("Expected read string = '%v'\n"+
      "Instead, read string='%v'\n",
      expectedStr, strRead)
  }

  err = fh.DeleteDirFile(targetFile)

  if err != nil {
    t.Errorf("Test Clean-up Error: Error returned by fh.DeleteDirFile(targetFile).\n"+
      "targetFile='%v'\nError='%v'\n", targetFile, err.Error())
  }

  return
}

func TestFileHelper_OpenFileWriteOnly_05(t *testing.T) {
  fh := FileHelper{}
  targetFile := ""

  _, err := fh.OpenFileWriteOnly(targetFile, false)

  if err == nil {
    t.Error("Expected an error return from fh.OpenFileWriteOnly(targetFile, false)\n" +
      "because parameter 'targetFile' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  return
}

func TestFileHelper_OpenFileWriteOnly_06(t *testing.T) {
  fh := FileHelper{}
  targetFile := "     "

  _, err := fh.OpenFileWriteOnly(targetFile, false)

  if err == nil {
    t.Error("Expected an error return from fh.OpenFileWriteOnly(targetFile, false)\n" +
      "because parameter 'targetFile' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!")
  }

  return
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_01(t *testing.T) {

  fh := FileHelper{}

  pathStr := ""

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if newPathStr != "" {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal an empty string because 'pathStr' is an empty string.\n" +
      "However, a valid string was returned! ERROR!\nresult='%v'", newPathStr)
  }
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_02(t *testing.T) {

  fh := FileHelper{}

  pathStr := "      "

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if newPathStr != "" {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal an empty string because 'pathStr' consists entirely of blank spaces.\n" +
      "However, a valid string was returned! ERROR!\nresult='%v'", newPathStr)
  }
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_03(t *testing.T) {

  fh := FileHelper{}

  pathStrBase := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/" +
    "level_04_dir"

  pathStr := pathStrBase + "/"

  pathStr = fh.AdjustPathSlash(pathStr)

  pathStrBase = fh.AdjustPathSlash(pathStrBase)

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if pathStrBase != newPathStr  {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal\npathStrBase='%v'.\n" +
      "Instead,\nnewPathStr='%v'", pathStrBase, newPathStr)
  }
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_04(t *testing.T) {

  fh := FileHelper{}

  pathStrBase := "../filesfortest/levelfilesfortest/level_01_dir/level_02_dir/level_03_dir/" +
    "level_04_dir"

  pathStr := pathStrBase

  pathStr = fh.AdjustPathSlash(pathStr)

  pathStrBase = fh.AdjustPathSlash(pathStrBase)

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if pathStrBase != newPathStr  {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal\npathStrBase='%v'.\n" +
      "Instead,\nnewPathStr='%v'", pathStrBase, newPathStr)
  }
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_05(t *testing.T) {

  fh := FileHelper{}

  pathStrBase := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\" +
    "level_04_dir"

  pathStr := pathStrBase + "\\"

  if os.PathSeparator == '\\' {
    pathStr = strings.ReplaceAll(pathStr,"\\", "/")
    pathStrBase = strings.ReplaceAll(pathStrBase,"\\", "/")
  }

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if pathStrBase != newPathStr  {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal\npathStrBase='%v'.\n" +
      "Instead,\nnewPathStr='%v'", pathStrBase, newPathStr)
  }
}

func TestFileHelper_RemovePathSeparatorFromEndOfPathString_06(t *testing.T) {

  fh := FileHelper{}

  pathStrBase := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\" +
    "level_04_dir"

  pathStr := pathStrBase

  if os.PathSeparator == '\\' {
    pathStr = strings.ReplaceAll(pathStr,"\\", "/")
    pathStrBase = strings.ReplaceAll(pathStrBase,"\\", "/")
  }

  newPathStr := fh.RemovePathSeparatorFromEndOfPathString(pathStr)

  if pathStrBase != newPathStr  {
    t.Errorf("Expected result from fh.RemovePathSeparatorFromEndOfPathString(pathStr) to\n" +
      "equal\npathStrBase='%v'.\n" +
      "Instead,\nnewPathStr='%v'", pathStrBase, newPathStr)
  }
}

func TestFileHelper_SearchFileModeMatch_01(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt"

  fh := FileHelper{}

  fInfo, err := fh.GetFileInfo(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetFile).\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    return
  }

  fileSelection := FileSelectionCriteria{}

  err = fileSelection.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")

  if err != nil {
    t.Errorf("Error returned by fileSelection.SelectByFileMode.SetFileModeByTextCode" +
      "(\"-r--r--r--\").\nError='%v'\n", err.Error())
    return
  }

  isFileModeSet, isFileModeMatch, err := fh.SearchFileModeMatch(fInfo, fileSelection)

  if err != nil {
    t.Errorf("Error returned by fh.SearchFileModeMatch(fInfo, fileSelection).\n" +
      "Error='%v'\n", err.Error())
  }

  if isFileModeSet == false {
    t.Error("Expected isFileModeSet=='true'. Instead, it is 'false'!")
  }

  if isFileModeMatch == true {
    t.Error("Expected isFileModeMatch=='false'. Instead, it is 'true'!")
  }

}

func TestFileHelper_SearchFileModeMatch_02(t *testing.T) {

  targetFile := "../filesfortest/levelfilesfortest/level_01_dir/level_1_3_test.txt"

  fh := FileHelper{}

  fInfo, err := fh.GetFileInfo(targetFile)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetFile).\n" +
      "targetFile='%v'\nError='%v'", targetFile, err.Error())
    return
  }

  fileSelection := FileSelectionCriteria{}

  err = fileSelection.SelectByFileMode.SetByFileMode(fInfo.Mode())

  if err != nil {
    t.Errorf("Error returned by fileSelection.SelectByFileMode.SetByFileMode"+
      "(fInfo.Mode()).\nError='%v'\n", err.Error())
    return
  }

  isFileModeSet, isFileModeMatch, err := fh.SearchFileModeMatch(fInfo, fileSelection)

  if err != nil {
    t.Errorf("Error returned by fh.SearchFileModeMatch(fInfo, fileSelection).\n" +
      "Error='%v'\n", err.Error())
  }

  if isFileModeSet == false {
    t.Error("Expected isFileModeSet=='true'. Instead, it is 'false'!")
  }

  if isFileModeMatch == false {
    t.Error("Expected isFileModeMatch=='true'. Instead, it is 'false'!")
  }

}

func TestFileHelper_SearchFileModeMatch_03(t *testing.T) {

  fh := FileHelper{}

  fileSelection := FileSelectionCriteria{}

  var fInfo os.FileInfo

  isFileModeSet, isFileModeMatch, err := fh.SearchFileModeMatch(fInfo, fileSelection)

  if err != nil {
    t.Errorf("Error returned by fh.SearchFileModeMatch(fInfo, fileSelection).\n" +
      "Error='%v'\n", err.Error())
  }

  if isFileModeSet == true {
    t.Error("Expected isFileModeSet=='false'. Instead, it is 'true'!")
  }

  if isFileModeMatch == true {
    t.Error("Expected isFileModeMatch=='false'. Instead, it is 'true'!")
  }

}

func TestFileHelper_SwapBasePath_01(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesfortest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"
  expectedTargetPath = fh.AdjustPathSlash(expectedTargetPath)

  newPath, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_02(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  expectedTargetPath := "../dirmgrtests/level_0_0_test.txt"
  expectedTargetPath = fh.AdjustPathSlash(expectedTargetPath)

  newPath, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err != nil {
    t.Errorf("Error returned from FileHelper{}.SwapBasePath(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedTargetPath != newPath {
    t.Errorf("Error: Expected newPath='%v'. Instead, newPath='%v' ",
      expectedTargetPath, newPath)
  }

}

func TestFileHelper_SwapBasePath_03(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from FileHelper{}.SwapBasePath(...) " +
      "NO ERROR WAS GENERATED!")
  }

}

func TestFileHelper_SwapBasePath_04(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := ""

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
    "because parameter 'oldBasePath' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFileHelper_SwapBasePath_05(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "   "

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
    "because parameter 'oldBasePath' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFileHelper_SwapBasePath_06(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "     "

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because parameter 'newBasePath' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }
}

func TestFileHelper_SwapBasePath_07(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := ""

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because parameter 'newBasePath' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }
}


func TestFileHelper_SwapBasePath_08(t *testing.T) {

  fh := FileHelper{}

  targetPath := "     "

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)


  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because parameter 'targetPath' consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }

}


func TestFileHelper_SwapBasePath_09(t *testing.T) {

  fh := FileHelper{}

  targetPath := ""

  oldBasePath := "../filesforTest/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because parameter 'targetPath' is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFileHelper_SwapBasePath_10(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "../filesforTest/levelfilesfortest/dir01/dir02/dir03/dir05/dir06"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because parameter 'oldBasePath' is longer than 'targetBasePath.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }
}


func TestFileHelper_SwapBasePath_11(t *testing.T) {

  fh := FileHelper{}

  targetPath := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  targetPath = fh.AdjustPathSlash(targetPath)

  oldBasePath := "/levelfilesfortest"
  oldBasePath = fh.AdjustPathSlash(oldBasePath)

  newBasePath := "../dirmgrtests"
  newBasePath = fh.AdjustPathSlash(newBasePath)

  _, err := fh.SwapBasePath(
    oldBasePath,
    newBasePath,
    targetPath)

  if err == nil {
    t.Error("Expected an error return from fh.SwapBasePath(oldBasePath,newBasePath,targetPath)\n" +
      "because 'oldBasePath' does NOT begin at beginning of 'targetBasePath.\n" +
      "However, NO ERROR WAS RETURNED!!")
  }

}


func createALogTestBottomDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if !fh.DoesFileExist(targetDir) {
    err2 := fh.MakeDirAll(targetDir)

    if err2 != nil {
      return err2
    }
  }

  targetFile := fh.JoinPathsAdjustSeparators(targetDir, alogFile)

  if fh.DoesFileExist(targetFile) {
    err3 := fh.DeleteDirFile(targetFile)
    if err3 != nil {
      return err3
    }
  }

  f, err4 := fh.CreateFile(targetFile)

  if err4 != nil {
    return err4
  }

  t := time.Now().Local()

  nowTime := t.Format("2006-01-02 15:04:05.000000000")

  _, err5 := f.WriteString("Sample Write - " + nowTime + "/n")

  if err5 != nil {
    _ = f.Close()
    return err5
  }

  _, err6 := f.WriteString("File Name: " + targetFile)

  if err6 != nil {
    _ = f.Close()
    return err6
  }

  _ = f.Close()
  return nil
}

func deleteALogTestBottomDirTargetDir() error {
  fh := FileHelper{}
  targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

  if err1 != nil {
    return err1
  }

  if fh.DoesFileExist(targetDir) {
    err2 := fh.DeleteDirPathAll(targetDir)

    if err2 != nil {
      return err2
    }

    if fh.DoesFileExist(targetDir) {
      return errors.New("File still exists:" + targetDir)
    }
  }

  return nil
}

