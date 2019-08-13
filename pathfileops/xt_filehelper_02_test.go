package pathfileops

import (
  "fmt"
  "os"
  "strings"
  "testing"
)

func TestFileHelper_CleanFileNameExtStr_01(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")
  expectedFileNameExt := "newerFileForTest_01.txt"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_02(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("newerFileForTest_01.txt")
  expectedFileNameExt := "newerFileForTest_01.txt"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath).\n"+
      "testPath='%v'\nError='%v'",
      testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'.\n" +
      "Instead, isFileNameEmpty='true'\n")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'.\n"+
      "Instead, it returned '%v'\n",
      expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_03(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
  // testPath is a directory which actually exists
  _, isEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Unexpected Error returned by fh.CleanFileNameExtStr(testPath)\n"+
      "testPath='%v'\nError='%v'\n",
      testPath, err.Error())
    return
  }

  if isEmpty == false {
    t.Error("ERROR: Expected 'isEmpty' == 'true' because input parameter\n" +
      "'testPath' was an actual directory the physically exists on disk.\n" +
      "However, 'isEmpty' return value was 'false'!\n")
  }

}

func TestFileHelper_CleanFileNameExtStr_04(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"\") " +
      "because the input parameter is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_05(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("     ")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"    \") " +
      "because the input parameter consists of all spaces. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_06(t *testing.T) {
  fh := FileHelper{}

  _, _, err := fh.CleanFileNameExtStr("...\\")

  if err == nil {
    t.Error("Expected error return from fh.CleanFileNameExtStr(\"    \") " +
      "because the input parameter includes 3-dots ('...'). " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileHelper_CleanFileNameExtStr_07(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/.gitignore")
  expectedFileNameExt := ".gitignore"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
  }

}

func TestFileHelper_CleanFileNameExtStr_08(t *testing.T) {
  fh := FileHelper{}
  testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01")
  expectedFileNameExt := "newerFileForTest_01"
  result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

  if err != nil {
    t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'",
      testPath, err.Error())
  }

  if isFileNameEmpty {
    t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
  }

  if expectedFileNameExt != result {
    t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'",
      expectedFileNameExt, result)
  }

}

func TestFileHelper_ConsolidateErrors_01(t *testing.T) {

  errs := make([]error, 0, 100)

  for i:=0; i < 3; i++ {
    errNo := fmt.Sprintf("Error #%0.3d: Error message.\n", i)
    err := fmt.Errorf(errNo)

    errs = append(errs, err)
  }

  fh := FileHelper{}

  err := fh.ConsolidateErrors(errs)

  if err == nil {
    t.Errorf("ERROR: fh.ConsolidateErrors(errs) returned 'nil'\n" +
      "instead of the expected error value.\n")
    return
  }

  finalErrStr := fmt.Sprintf("%v", err.Error())

  for k:=0; k < 3; k++ {
    testStr := fmt.Sprintf("Error #%0.3d:", k)

    if !strings.Contains(finalErrStr,testStr) {
      t.Errorf("Error: Expected final error string would contain %v.\n" +
        "It did NOT contain that text!\n" +
        "Final Err Str='%v'\n" +
        "Test text='%v'\n",
        testStr,finalErrStr,testStr)
      return
    }
  }
}

func TestFileHelper_ConsolidateErrors_02(t *testing.T) {
  errs := make([]error, 0, 100)

  fh := FileHelper{}

  err := fh.ConsolidateErrors(errs)

  if err != nil {
    t.Error("ERROR: Expected fh.ConsolidateErrors(errs) to return 'nil'\n" +
      "because 'errs' is an empty array.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

}

func TestFileHelper_ConvertOctalToDecimal_01(t *testing.T) {

  fh := FileHelper{}
  expectedValue := 511

  octalValue := 777

  mode := fh.ConvertOctalToDecimal(octalValue)

  if expectedValue != mode {
    t.Errorf("Error: Expected Value='%v'. Instead, value='%v' ",
      expectedValue, mode)
  }

}

func TestFileHelper_ConvertOctalToDecimal_02(t *testing.T) {

  fh := FileHelper{}
  expectedValue := 438

  octalValue := 666

  mode := fh.ConvertOctalToDecimal(octalValue)

  if expectedValue != mode {
    t.Errorf("Error: Expected Value='%v'. Instead, value='%v' ",
      expectedValue, mode)
  }

}

func TestFileHelper_ConvertDecimalToOctal_01(t *testing.T) {

  fh := FileHelper{}

  expectedOctalValue := 777

  initialDecimalValue := 511

  actualOctalValue := fh.ConvertDecimalToOctal(initialDecimalValue)

  if expectedOctalValue != actualOctalValue {
    t.Errorf("Error: Expected ocatal value='%v'. Instead, actual ocatal value='%v' ",
      expectedOctalValue, actualOctalValue)
  }

}

func TestFileHelper_ConvertDecimalToOctal_02(t *testing.T) {

  fh := FileHelper{}

  expectedOctalValue := 666

  initialDecimalValue := 438

  actualOctalValue := fh.ConvertDecimalToOctal(initialDecimalValue)

  if expectedOctalValue != actualOctalValue {
    t.Errorf("Error: Expected ocatal value='%v'. Instead, actual ocatal value='%v' ",
      expectedOctalValue, actualOctalValue)
  }

}

func TestFileHelper_CopyFileByIo_01(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestCopyFile80179658.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile). "+
      "rawDestFile='%v' Error='%v' ", rawDestFile, err.Error())
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

  err = FileHelper{}.CopyFileByIo("", destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(\"\", destFile) " +
      "because input parameter source file is an empty string. " +
      "However, NO ERROR WAS RETURNED!")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

}

func TestFileHelper_CopyFileByIo_02(t *testing.T) {

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "")

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,\"\")\n" +
      "because input parameter destination file is an empty string.\n" +
      "\nHowever, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_03(t *testing.T) {

  rawDestFile := "..\\checkfiles\\TestFileHelper_CopyFileByIo_03.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  _ = FileHelper{}.DeleteDirFile(destFile)

  rawSrcFile := "..\\checkfiles\\iDoNOTExist.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file does not exist.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  _ = FileHelper{}.DeleteDirFile(destFile)
}

func TestFileHelper_CopyFileByIo_04(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  srcFile := destFile

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file is equivalent to destination file.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }
}

func TestFileHelper_CopyFileByIo_05(t *testing.T) {

  rawDestFile := "..\\checkfiles\\TestFileHelper_CopyFileByIo_05.txt"
  fh := FileHelper{}

  destFile, err := fh.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo("   ", destFile)

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(\"   \", destFile)\n" +
      "because input parameter source file name consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

  if fh.DoesFileExist(destFile) {
    err = fh.DeleteDirFile(destFile)
    if err != nil {
      t.Errorf("Error returned from last attempt to delete destFile.\n"+
        "fh.DeleteDirFile(destFile)\ndestFile='%v'\nError='%v'\n",
        destFile, err.Error())
    }
  }
}

func TestFileHelper_CopyFileByIo_06(t *testing.T) {

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_0_test.txt"
  fh := FileHelper{}

  srcFile, err := fh.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, "   ")

  if err == nil {
    t.Error("Expected an error return from  err = FileHelper{}.CopyFileByIo(src, \"    \")\n" +
      "because input parameter destination file consists entirely of blank spaces.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_07(t *testing.T) {

  rawDestFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  destFile, err := FileHelper{}.MakeAbsolutePath(rawDestFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawDestFile).\n"+
      "rawDestFile='%v'\nError='%v'\n",
      rawDestFile, err.Error())
    return
  }

  srcFile := ""

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter source file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_08(t *testing.T) {

  destFile := ""

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_0_0_test.txt"

  srcFile, err := FileHelper{}.MakeAbsolutePath(rawSrcFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.MakeAbsolutePath(rawSrcFile).\n"+
      "rawSrcFile='%v'\nError='%v'\n",
      rawSrcFile, err.Error())
    return
  }

  err = FileHelper{}.CopyFileByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error from FileHelper{}.CopyFileByIo(srcFile,destFile)\n" +
      "because input parameter destination file is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIo_09(t *testing.T) {

  fh := FileHelper{}
  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir" +
    "\\level_03_dir\\level_3_1_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("ERROR: Setup source file DOES NOT EXIST!\n"+
      "srcFile='%v' \n", srcFile)
    return
  }

  rawDestFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_09.txt"

  destFile := fh.AdjustPathSlash(rawDestFile)

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error returned from fh.DeleteDirFile(destFile).\n"+
        "Attempt to delete prexisting version of destination file FAILED!\n"+
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("ERROR: Prexisting Destination File could NOT be Deleted!\n"+
        "Destination File:'%v'\n", destFile)
      return
    }
  }

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error while Copying Source File to  Destination File!\n"+
      "Source File='%v'\nDestination File='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Error(fmt.Sprintf("After CopyIO Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile))
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile).\n"+
      "During clean-up, the attempted deletion of the destination file FAILED!\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of the destination file during "+
      "clean-up FAILED!\ndestFile='%v'", destFile)
  }

}

func TestFileHelper_CopyFileByIo_10(t *testing.T) {

  rawDestFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_10.txt"

  fh := FileHelper{}

  destFile := fh.AdjustPathSlash(rawDestFile)

  err := fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error retunred by fh.DeleteDirFile(destFile) during setup.\n"+
      "Attempt deletion of pre-existing version of destination file FAILED!\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Unable to delete pre-existing version of destination file!\n"+
      "destFile='%v'", destFile)
    return
  }

  rawSrcFile := "../filesfortest/levelfilesfortest/level_0_2_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  err = fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by FileHelper{}.CopyFileByIo(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: CopyFileByIo FAILED! The destination file was NOT created!\n"+
      "destFile='%v'\n", destFile)
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile).\n"+
      "Attempted deletion of destination file during clean-up FAILED!\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of destination file during clean-up FAILED!\n"+
      "Destination File STILL EXISTS!\n"+
      "Destination File='%v'\n", destFile)
  }

}

func TestFileHelper_CopyFileByIo_11(t *testing.T) {

  fh := FileHelper{}

  destFile := "..\\checkfiles\\scratchTestFileHelper_CopyFileByIo_11.txt"

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error retrned by setup fh.DeleteDirFile(destFile).\n"+
        "Attempted deletion of pre-existing destination file FAILED!\n"+
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("Attempted deletion of pre-existing destination file FAILED!\n"+
        "destFile='%v'\n", destFile)
      return
    }

  }

  srcFile := "../filesfortest/levelfilesfortest/level_0_2_test.txt"

  err := fh.CopyFileByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by setup fh.CopyFileByIo(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: Attempted creation of setup destination file FAILED!\n"+
      "destFile='%v'\n", destFile)
    return
  }

  srcFile2 := "../filesfortest/levelfilesfortest/level_0_3_test.txt"

  err = fh.CopyFileByIo(srcFile2, destFile)

  if err != nil {
    t.Errorf("Error returned by 2nd Copy fh.CopyFileByIo(srcFile2, destFile).\n"+
      "srcFile2='%v'\ndestFile='%v\nError='%v'\n",
      srcFile2, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Errorf("Error: 2nd Copy of destination file does NOT exist!\n"+
      "Destination File='%v'\n", destFile)
    return
  }

  // 2nd destination file DOES EXIST!

  finfoSrcFile, err := os.Stat(srcFile2)

  if err != nil {
    t.Errorf("Error returned by os.Stat(srcFile2).\n"+
      "srcFile2='%v'\nError='%v'\n", srcFile2, err.Error())
  }

  finfoDestFile, err := os.Stat(destFile)

  if err != nil {
    t.Errorf("Error returned by os.Stat(destFile).\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    _ = fh.DeleteDirFile(destFile)
    return
  }

  if finfoSrcFile.Size() != finfoDestFile.Size() {
    t.Errorf("The sizes of the source file and destination file DO NOT MATHCH!\n"+
      "Source File Size='%v'  Destination File Size='%v'.\n",
      finfoSrcFile.Size(), finfoDestFile.Size())
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by clean-up fh.DeleteDirFile(destFile).\n"+
      "destFile='%v'\nError='%v' ", destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of destination file FAILED!\n"+
      "Destination File='%v'\n", destFile)
  }

}

func TestFileHelper_CopyFileByIoByLink_01(t *testing.T) {

  fh := FileHelper{}
  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir" +
    "\\level_03_dir\\level_3_1_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  if !fh.DoesFileExist(srcFile) {
    t.Errorf("ERROR: Setup source file DOES NOT EXIST!\n"+
      "srcFile='%v' \n", srcFile)
    return
  }

  rawDestFile := "..\\checkfiles\\CopyFileByIoByLink_01.txt"

  destFile := fh.AdjustPathSlash(rawDestFile)

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error returned from fh.DeleteDirFile(destFile).\n"+
        "Attempt to delete prexisting version of destination file FAILED!\n"+
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("ERROR: Prexisting Destination File could NOT be Deleted!\n"+
        "Destination File:'%v'\n", destFile)
      return
    }
  }

  err := fh.CopyFileByIoByLink(srcFile, destFile)

  if err != nil {
    t.Errorf("Error while Copying Source File to  Destination File!\n"+
      "Source File='%v'\nDestination File='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if !fh.DoesFileExist(destFile) {
    t.Error(fmt.Sprintf("After CopyIO Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile))
    return
  }

  err = fh.DeleteDirFile(destFile)

  if err != nil {
    t.Errorf("Error returned by fh.DeleteDirFile(destFile).\n"+
      "During clean-up, the attempted deletion of the destination file FAILED!\n"+
      "destFile='%v'\nError='%v'\n", destFile, err.Error())
    return
  }

  if fh.DoesFileExist(destFile) {
    t.Errorf("ERROR: Attempted deletion of the destination file during "+
      "clean-up FAILED!\ndestFile='%v'", destFile)
  }
}

func TestFileHelper_CopyFileByIoByLink_02(t *testing.T) {

  fh := FileHelper{}
  rawSrcFile := "../checkfiles/iDoNotExist.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  rawDestFile := "..\\checkfiles\\TestFileHelper_CopyFileByIoByLink_02.txt"

  destFile := fh.AdjustPathSlash(rawDestFile)

  if fh.DoesFileExist(destFile) {

    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error returned from fh.DeleteDirFile(destFile).\n"+
        "Attempt to delete prexisting version of destination file FAILED!\n"+
        "destFile='%v'\nError='%v'\n", destFile, err.Error())
      return
    }

    if fh.DoesFileExist(destFile) {
      t.Errorf("ERROR: Prexisting Destination File could NOT be Deleted!\n"+
        "Destination File:'%v'\n", destFile)
      return
    }
  }

  err := fh.CopyFileByIoByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.CopyFileByIoByLink(srcFile, destFile)\n" +
      "because 'srcFile' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByIoByLink_03(t *testing.T) {

  fh := FileHelper{}

  rawSrcFile := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir" +
    "\\level_03_dir\\level_3_1_test.txt"

  srcFile := fh.AdjustPathSlash(rawSrcFile)

  rawDestFile := "../checkfiles/checkfiles02"

  destFile := fh.AdjustPathSlash(rawDestFile)

  err := fh.CopyFileByIoByLink(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.CopyFileByIoByLink(srcFile, destFile)\n" +
      "because 'srcFile' DOES NOT EXIST!\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileHelper_CopyFileByLinkByIo_01(t *testing.T) {

  fh := FileHelper{}

  setupSrcFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  srcFile:="../createFilesTest/level_0_1_test.txt"

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Test Setup Error: Setup source file previously exists and "+
        "cannot be deleted!\nsrcFile='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n" +
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    return
  }

  destFile := "../createFilesTest/TestFileHelper_CopyFileByLinkByIo_01.txt"

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }
  }

  err = fh.CopyFileByLinkByIo(srcFile, destFile)

  if err != nil {
    t.Errorf("Error returned by fh.CopyFileByLinkByIo(srcFile, destFile).\n"+
      "srcFile='%v'\ndestFile='%v'\nError='%v'\n",
      srcFile, destFile, err.Error())
  }

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
        "Deletion of Destination File Failed. "+
        "It cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  } else {
    t.Errorf("Error: CopyFileByLink Failed. Destination File DOES NOT EXIST!\n"+
      "destFile='%v'\n", destFile)
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(srcFile)\n" +
      "srcFile='%v'\nError='%v'\n",srcFile, err.Error())
  }

}

func TestFileHelper_CopyFileByLinkByIo_02(t *testing.T) {

  fh := FileHelper{}

  srcFile:=""

  destFile := "../createFilesTest/TestFileHelper_CopyFileByLinkByIo_02.txt"

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error: Target destination file previously exists and "+
        "cannot be deleted!\ndestFile='%v'", destFile)
      return
    }
  }

  err := fh.CopyFileByLinkByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an error return from fh.CopyFileByLinkByIo(srcFile, destFile).\n"+
      "because 'srcFile' is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!\n")
  }

  if fh.DoesFileExist(destFile) {
    err := fh.DeleteDirFile(destFile)

    if err != nil {
      t.Errorf("Error After Copy Destination File Existed. However, the attempted "+
        "Deletion of Destination File Failed. "+
        "It cannot be deleted!\ndestFile='%v'", destFile)
      return
    }

  }

}

func TestFileHelper_CopyFileByLinkByIo_03(t *testing.T) {

  fh := FileHelper{}

  setupSrcFile := "../filesfortest/levelfilesfortest/level_0_1_test.txt"

  srcFile:="../createFilesTest/level_0_1_test.txt"

  if fh.DoesFileExist(srcFile) {
    err := fh.DeleteDirFile(srcFile)

    if err != nil {
      t.Errorf("Test Setup Error: Setup source file previously exists and "+
        "cannot be deleted!\nsrcFile='%v'", srcFile)
      return
    }
  }

  err := fh.CopyFileByIo(setupSrcFile, srcFile)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.CopyFileByIo(setupSrcFile, srcFile)\n" +
      "setupSrcFile='%v'\nsrcFile='%v'\nError='%v'\n",
      setupSrcFile, srcFile, err.Error())
    return
  }

  destFile := ""

  err = fh.CopyFileByLinkByIo(srcFile, destFile)

  if err == nil {
    t.Error("Expected an rror return from fh.CopyFileByLinkByIo(srcFile, destFile).\n"+
      "because destFile is an empty string!\n" +
      "However, NO ERROR WAS RETURNED!!\n")
  }

  err = fh.DeleteDirFile(srcFile)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirFile(srcFile)\n" +
      "srcFile='%v'\nError='%v'\n",srcFile, err.Error())
  }
}
