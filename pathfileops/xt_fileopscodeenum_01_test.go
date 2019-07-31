package pathfileops

import (
  "strings"
  "testing"
)

func TestFileOperationCode_01(t *testing.T) {

  fopFirst := FileOpCode.None()

  if int(fopFirst) != 0 {
    t.Errorf("Error: Expected first File Operations Code = 0.  Instead, first "+
      "File Operation Code = '%v' ", int(fopFirst))
  }

  if 0 != fopFirst.Value() {
    t.Errorf("Error: Expected first File Operations Code Value = 0.  Instead, first "+
      "File Operation Code Value = '%v' ", fopFirst.Value())
  }

  fopLast := FileOpCode.CreateDestinationFile()

  if int(fopLast) != 15 {
    t.Errorf("Error: Expected FileOpCode.CreateDestinationFile() = 15.  Instead, "+
      "FileOpCode.CreateDestinationFile() = '%v' ", int(fopLast))
  }

  if 15 != fopLast.Value() {
    t.Errorf("Error: Expected FileOpCode.CreateDestinationFile() Value = 15.  Instead, "+
      "FileOpCode.CreateDestinationFile() = '%v' ", fopLast.Value())
  }

}

func TestFileOperationCode_02(t *testing.T) {

  fopNone := FileOpCode.None()

  strValue := fopNone.String()

  if "None" != strValue {
    t.Errorf("Error: Expected string value of FileOpCode.None() = 'None' .  Instead, "+
      "string value of FileOpCode.None() = '%v' ", strValue)
  }

  fopLast := FileOpCode.CreateDestinationFile()

  strValue = fopLast.String()

  if "CreateDestinationFile" != strValue {
    t.Errorf("Error: Expected string value of FileOpCode.CreateDestinationFile() = 'CreateDestinationFile' . "+
      " Instead, string value of FileOpCode.CreateDestinationFile() = '%v' ", strValue)

  }

}

func TestFileOperationCode_03(t *testing.T) {

  strValue := "None"

  fopNone, err := FileOpCode.ParseString(strValue, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(strValue, true). "+
      "strValue='%v' Error='%v' ", strValue, err.Error())
    return
  }

  if fopNone != FileOperationCode(0).None() {
    t.Errorf("Error: Expected fopNone = FileOperationCode(0).None() .  Instead, "+
      "string value of fopNone = '%v' - int Value of fopNone='%v' ",
      fopNone.String(), int(fopNone))
  }

  strValue = "CreateDestinationFile"

  fopLast, err := FileOpCode.ParseString(strValue, true)

  if err != nil {
    t.Errorf("Error returned by (2) FileOpCode.ParseString(strValue, true). "+
      "strValue='%v' Error='%v' ", strValue, err.Error())
    return
  }

  if fopLast != FileOperationCode(0).CreateDestinationFile() {
    t.Errorf("Error: Expected string value of fopLast = FileOperationCode(0)."+
      "CreateDestinationFile(). Instead, string value of fopLast = '%v' - "+
      "int value of fopLas = '%v' ", fopLast.String(), int(fopLast))

  }

}

func TestFileOperationCode_04(t *testing.T) {

  opsAry := make([]FileOperationCode, 15)

  opsAry[0] = FileOperationCode(0).None()
  opsAry[1] = FileOpCode.MoveSourceFileToDestinationFile()
  opsAry[2] = fileOpCode.DeleteDestinationFile()
  opsAry[3] = FileOperationCode(0).DeleteSourceFile()
  opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
  opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
  opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
  opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
  opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
  opsAry[9] = FileOperationCode(0).CreateSourceDir()
  opsAry[10] = FileOpCode.CreateSourceDirAndFile()
  opsAry[11] = FileOpCode.CreateSourceFile()
  opsAry[12] = fileOpCode.CreateDestinationDir()
  opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
  opsAry[14] = FileOpCode.CreateDestinationFile()

  opsStrings := make([]string, 15)
  opsStrings[0] = "None"
  opsStrings[1] = "MoveSourceFileToDestinationFile"
  opsStrings[2] = "DeleteDestinationFile"
  opsStrings[3] = "DeleteSourceFile"
  opsStrings[4] = "DeleteSourceAndDestinationFiles"
  opsStrings[5] = "CopySourceToDestinationByHardLinkByIo"
  opsStrings[6] = "CopySourceToDestinationByIoByHardLink"
  opsStrings[7] = "CopySourceToDestinationByHardLink"
  opsStrings[8] = "CopySourceToDestinationByIo"
  opsStrings[9] = "CreateSourceDir"
  opsStrings[10] = "CreateSourceDirAndFile"
  opsStrings[11] = "CreateSourceFile"
  opsStrings[12] = "CreateDestinationDir"
  opsStrings[13] = "CreateDestinationDirAndFile"
  opsStrings[14] = "CreateDestinationFile"

  for i := 0; i < len(opsAry); i++ {

    if opsAry[i].String() != opsStrings[i] {
      t.Errorf("Error: opsAry[i].String() != opsStrings[i]. "+
        "opsAry[%v].String()='%v' opsStrings[%v]='%v'", i, opsAry[i].String(), i, opsStrings[i])
    }

  }

}

func TestFileOperationCode_05(t *testing.T) {

  opsAry := make([]FileOperationCode, 15)

  opsAry[0] = FileOperationCode(0).None()
  opsAry[1] = FileOpCode.MoveSourceFileToDestinationFile()
  opsAry[2] = fileOpCode.DeleteDestinationFile()
  opsAry[3] = FileOperationCode(0).DeleteSourceFile()
  opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
  opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
  opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
  opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
  opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
  opsAry[9] = FileOperationCode(0).CreateSourceDir()
  opsAry[10] = FileOpCode.CreateSourceDirAndFile()
  opsAry[11] = FileOpCode.CreateSourceFile()
  opsAry[12] = fileOpCode.CreateDestinationDir()
  opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
  opsAry[14] = FileOpCode.CreateDestinationFile()

  opsStrings := make([]string, 15)
  opsStrings[0] = "None"
  opsStrings[1] = "MoveSourceFileToDestinationFile"
  opsStrings[2] = "DeleteDestinationFile"
  opsStrings[3] = "DeleteSourceFile"
  opsStrings[4] = "DeleteSourceAndDestinationFiles"
  opsStrings[5] = "CopySourceToDestinationByHardLinkByIo"
  opsStrings[6] = "CopySourceToDestinationByIoByHardLink"
  opsStrings[7] = "CopySourceToDestinationByHardLink"
  opsStrings[8] = "CopySourceToDestinationByIo"
  opsStrings[9] = "CreateSourceDir"
  opsStrings[10] = "CreateSourceDirAndFile"
  opsStrings[11] = "CreateSourceFile"
  opsStrings[12] = "CreateDestinationDir"
  opsStrings[13] = "CreateDestinationDirAndFile"
  opsStrings[14] = "CreateDestinationFile"

  for i := 0; i < len(opsAry); i++ {

    fop, err := FileOperationCode(0).ParseString(opsStrings[i], true)

    if err != nil {
      t.Errorf("Error returned by FileOperationCode(0).ParseString(opsStrings[i], true). "+
        "i='%v' opsStrings[%v]='%v' Error='%v' ", i, i, opsStrings[i], err.Error())
      return
    }

    if fop != opsAry[i] {
      t.Errorf("Error: fop != opsAry[i]. "+
        "fop.String() ='%v' opsAry[%v]='%v'", fop.String(), i, opsAry[i].String())
    }

  }
}

func TestFileOperationCode_06(t *testing.T) {

  opsAry := make([]FileOperationCode, 15)

  opsAry[0] = FileOperationCode(0).None()
  opsAry[1] = FileOpCode.MoveSourceFileToDestinationFile()
  opsAry[2] = fileOpCode.DeleteDestinationFile()
  opsAry[3] = FileOperationCode(0).DeleteSourceFile()
  opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
  opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
  opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
  opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
  opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
  opsAry[9] = FileOperationCode(0).CreateSourceDir()
  opsAry[10] = FileOpCode.CreateSourceDirAndFile()
  opsAry[11] = FileOpCode.CreateSourceFile()
  opsAry[12] = fileOpCode.CreateDestinationDir()
  opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
  opsAry[14] = FileOpCode.CreateDestinationFile()

  opsStrings := make([]string, 15)
  opsStrings[0] = strings.ToLower("None")
  opsStrings[1] = strings.ToLower("MoveSourceFileToDestinationFile")
  opsStrings[2] = strings.ToLower("DeleteDestinationFile")
  opsStrings[3] = strings.ToLower("DeleteSourceFile")
  opsStrings[4] = strings.ToLower("DeleteSourceAndDestinationFiles")
  opsStrings[5] = strings.ToLower("CopySourceToDestinationByHardLinkByIo")
  opsStrings[6] = strings.ToLower("CopySourceToDestinationByIoByHardLink")
  opsStrings[7] = strings.ToLower("CopySourceToDestinationByHardLink")
  opsStrings[8] = strings.ToLower("CopySourceToDestinationByIo")
  opsStrings[9] = strings.ToLower("CreateSourceDir")
  opsStrings[10] = strings.ToLower("CreateSourceDirAndFile")
  opsStrings[11] = strings.ToLower("CreateSourceFile")
  opsStrings[12] = strings.ToLower("CreateDestinationDir")
  opsStrings[13] = strings.ToLower("CreateDestinationDirAndFile")
  opsStrings[14] = strings.ToLower("CreateDestinationFile")

  for i := 0; i < len(opsAry); i++ {

    fop, err := FileOperationCode(0).ParseString(opsStrings[i], false)

    if err != nil {
      t.Errorf("Error returned by FileOperationCode(0).ParseString(opsStrings[i], true). "+
        "i='%v' opsStrings[%v]='%v' Error='%v' ", i, i, opsStrings[i], err.Error())
      return
    }

    if fop != opsAry[i] {
      t.Errorf("Error: fop != opsAry[i]. "+
        "fop.String() ='%v' opsAry[%v]='%v'", fop.String(), i, opsAry[i].String())
    }

  }
}

func TestFileOperationCode_IsValid_01(t *testing.T) {

  c := FileOpCode.CreateDestinationFile()

  err := c.IsValid()

  if err != nil {
    t.Errorf("Error: Expected FileOpCode.CreateDestinationFile() to be Valid. "+
      "It was INVALID!. Error='%v' ", err.Error())
  }

}

func TestFileOperationCode_IsValid_02(t *testing.T) {

  c := FileOperationCode(-99)

  err := c.IsValid()

  if err == nil {
    t.Error("Error: Expected FileOperationCode(-99) to be INVALID. " +
      "Instead, it registered as VALID!.")
  }

}

func TestFileOperationCode_IsValid_03(t *testing.T) {

  c := FileOpCode.CreateSourceDir()

  err := c.IsValid()

  if err != nil {
    t.Errorf("Error: Expected FileOpCode.CreateSourceDir() to be Valid. "+
      "It was INVALID!. Error='%v' ", err.Error())
  }

}

func TestFileOperationCode_IsValid_04(t *testing.T) {

  c := FileOpCode.CopySourceToDestinationByHardLink()

  err := c.IsValid()

  if err != nil {
    t.Errorf("Error: Expected FileOpCode.CopySourceToDestination" +
      "ByHardLink() to be Valid. "+
      "It was INVALID!. Error='%v' ", err.Error())
  }

}

func TestFileOperationCode_IsValid_05(t *testing.T) {

  c := FileOpCode.CopySourceToDestinationByHardLinkByIo()

  err := c.IsValid()

  if err != nil {
    t.Errorf("Error: Expected FileOpCode.CopySourceToDestination" +
      "ByHardLinkByIo() to be Valid. "+
      "It was INVALID!. Error='%v' ", err.Error())
  }

}


func TestFileOperationCode_ParseString_01(t *testing.T) {

  _, err := FileOpCode.ParseString("Mo", true)

  if err == nil {
    t.Error("Expected an error return from FileOpCode." +
      "ParseString(\"Mo\", true) because input is less than 3-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOperationCode_ParseString_02(t *testing.T) {

  _, err := FileOpCode.ParseString("Mxzustvyf", true)

  if err == nil {
    t.Error("Expected an error return from FileOpCode." +
      "ParseString(\"Mo\", true) because input is less than 3-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOperationCode_ParseString03(t *testing.T) {

  inputStr := "None"

  expectedOpCode:= FileOpCode.None()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString04(t *testing.T) {

  inputStr := "MoveSourceFileToDestinationFile"

  expectedOpCode:= FileOpCode.MoveSourceFileToDestinationFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString05(t *testing.T) {

  inputStr := "DeleteDestinationFile"

  expectedOpCode:= FileOpCode.DeleteDestinationFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString06(t *testing.T) {

  inputStr := "DeleteSourceFile"

  expectedOpCode:= FileOpCode.DeleteSourceFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString07(t *testing.T) {

  inputStr := "DeleteSourceAndDestinationFiles"

  expectedOpCode:= FileOpCode.DeleteSourceAndDestinationFiles()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString08(t *testing.T) {

  inputStr := "CopySourceToDestinationByHardLinkByIo"

  expectedOpCode:= FileOpCode.CopySourceToDestinationByHardLinkByIo()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString09(t *testing.T) {

  inputStr := "CopySourceToDestinationByIoByHardLink"

  expectedOpCode:= FileOpCode.CopySourceToDestinationByIoByHardLink()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString10(t *testing.T) {

  inputStr := "CopySourceToDestinationByHardLink"

  expectedOpCode:= FileOpCode.CopySourceToDestinationByHardLink()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString11(t *testing.T) {

  inputStr := "CopySourceToDestinationByIo"

  expectedOpCode:= FileOpCode.CopySourceToDestinationByIo()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString12(t *testing.T) {

  inputStr := "CreateSourceDir"

  expectedOpCode:= FileOpCode.CreateSourceDir()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString13(t *testing.T) {

  inputStr := "CreateSourceDirAndFile"

  expectedOpCode:= FileOpCode.CreateSourceDirAndFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString14(t *testing.T) {

  inputStr := "CreateSourceFile"

  expectedOpCode:= FileOpCode.CreateSourceFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString15(t *testing.T) {

  inputStr := "CreateDestinationDir"

  expectedOpCode:= FileOpCode.CreateDestinationDir()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString16(t *testing.T) {

  inputStr := "CreateDestinationDirAndFile"

  expectedOpCode:= FileOpCode.CreateDestinationDirAndFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}

func TestFileOperationCode_ParseString17(t *testing.T) {

  inputStr := "CreateDestinationFile"

  expectedOpCode:= FileOpCode.CreateDestinationFile()

  actualOpCode, err := FileOpCode.ParseString(inputStr, true)

  if err != nil {
    t.Errorf("Error returned by FileOpCode.ParseString(inputStr, true). " +
      "inputStr='%v' Error='%v' ", inputStr, err.Error())
    return
  }

  if expectedOpCode != actualOpCode {
    t.Errorf("Expected OpCode string='%v' value='%v'. Instead, OpCode value='%v' ",
      expectedOpCode.String(), expectedOpCode, actualOpCode.Value())
  }

}


func TestFileOperationCode_ParseString_17(t *testing.T) {

  _, err := FileOpCode.ParseString("mo", false)

  if err == nil {
    t.Error("Expected an error return from FileOpCode." +
      "ParseString(\"Mo\", true) because input is less than 3-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOperationCode_ParseString_18(t *testing.T) {

  _, err := FileOpCode.ParseString("mxzustvyf", false)

  if err == nil {
    t.Error("Expected an error return from FileOpCode." +
      "ParseString(\"Mo\", true) because input is less than 3-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}


func TestFileOperationCode_String_19(t *testing.T) {

  fOpCode := FileOperationCode(-99)

  str := fOpCode.String()

  if str != "" {
    t.Errorf("Passed Invalid File Operation Code to String(). Expected an " +
      "empty string in return. No Empty String was returned. str='%v' ", str)
  }

}

func TestFileOperationCode_String_02(t *testing.T) {

  fileOpCode := FileOpCode.CreateSourceDir()
  expectedStr := "CreateSourceDir"
  actualStr := fileOpCode.String()

  if expectedStr != actualStr {
    t.Errorf("Expected String()='%v'. Instead, String()='%v' ",
      expectedStr, actualStr)
  }

}

func TestFileOperationCode_String_03(t *testing.T) {

  fileOpCode := FileOpCode.None()
  expectedStr := "None"
  actualStr := fileOpCode.String()

  if expectedStr != actualStr {
    t.Errorf("Expected String()='%v'. Instead, String()='%v' ",
      expectedStr, actualStr)
  }

}

func TestFileOperationCode_String_04(t *testing.T) {

  fileOpCode := FileOpCode.CopySourceToDestinationByHardLink()
  expectedStr := "CopySourceToDestinationByHardLink"
  actualStr := fileOpCode.String()

  if expectedStr != actualStr {
    t.Errorf("Expected String()='%v'. Instead, String()='%v' ",
      expectedStr, actualStr)
  }

}

func TestFileOperationCode_Value_01(t *testing.T) {

  fileOpCode := FileOpCode.CopySourceToDestinationByHardLink()

  expectedCode := 8

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_02(t *testing.T) {

  fileOpCode := FileOperationCode(0)

  expectedCode := 0

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_03(t *testing.T) {

  fileOpCode := FileOpCode.CreateSourceDir()

  expectedCode := 10

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_04(t *testing.T) {

  fileOpCode := FileOpCode.CopySourceToDestinationByHardLinkByIo()

  expectedCode := 6

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_05(t *testing.T) {

  fileOpCode := FileOpCode.CreateDestinationFile()

  expectedCode := 15

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_06(t *testing.T) {

  fileOpCode := FileOpCode.CreateDestinationDirAndFile()

  expectedCode := 14

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}

func TestFileOperationCode_Value_07(t *testing.T) {

  fileOpCode := FileOpCode.CreateDestinationDir()

  expectedCode := 13

  if expectedCode != fileOpCode.Value() {
    t.Errorf("Error: Expected fileOpCode.Value='%v'. Instead, value='%v'",
      expectedCode, fileOpCode.Value())
  }

}
