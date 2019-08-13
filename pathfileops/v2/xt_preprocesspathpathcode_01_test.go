package pathfileops

import "testing"

func TestPreProcessPathCode_None_01(t *testing.T) {

  preProcCode := PreProcPathCode.None()

  if preProcCode != 0 {
    t.Errorf("Error: Expected PreProcPathCode.None() would equal a value of '0'.\n" +
      "Instead, PreProcPathCode.None() equals a value of '%v'\n", int(preProcCode))
  }

}

func TestPreProcessPathCode_PathSeparator_01(t *testing.T) {
  preProcCode := PreProcPathCode.PathSeparator()

  if preProcCode != 1 {
    t.Errorf("Error: Expected PreProcPathCode.PathSeparator() would " +
      "equal a value of '1'.\n" +
      "Instead, PreProcPathCode.None() equals a value of '%v'\n", int(preProcCode))
  }
}

func TestPreProcessPathCode_AbsolutePath_01(t *testing.T) {

  preProcCode := PreProcPathCode.AbsolutePath()

  if preProcCode != 2 {
    t.Errorf("Error: Expected PreProcPathCode.AbsolutePath() would equal a value of '2'.\n" +
      "Instead, PreProcPathCode.None() equals a value of '%v'\n", int(preProcCode))
  }

}

func TestPreProcessPathCode_String_01(t *testing.T) {

  preProcCode := PreProcPathCode.None()

  noneStr := preProcCode.String()

  if noneStr != "None" {
    t.Errorf("Error: Expected PreProcPathCode.None() string value equals 'None'.\n" +
      "Instead, PreProcPathCode.None() string value equals '%v'\n", noneStr)
  }
}

func TestPreProcessPathCode_String_02(t *testing.T) {

  preProcCode := PreProcPathCode.PathSeparator()

  pathSepStr := preProcCode.String()

  if pathSepStr != "PathSeparator" {
    t.Errorf("Error: Expected PreProcPathCode.PathSeparator() string " +
      "value equals 'PathSeparator'.\n" +
      "Instead, PreProcPathCode.PathSeparator() string value equals '%v'\n", pathSepStr)
  }

}

func TestPreProcessPathCode_String_03(t *testing.T) {

  preProcCode := PreProcPathCode.AbsolutePath()

  absolutePathStr := preProcCode.String()

  if absolutePathStr != "AbsolutePath" {
    t.Errorf("Error: Expected PreProcPathCode.AbsolutePath() string " +
      "value equals 'AbsolutePath'.\n" +
      "Instead, PreProcPathCode.AbsolutePath() string value equals '%v'\n", absolutePathStr)
  }
}

func TestPreProcessPathCode_String_04(t *testing.T) {

  status := PreProcessPathCode(-99)

  statusStr := status.String()

  if statusStr != "" {
    t.Errorf("Error: Expected status.String()==\"\"\n" +
      "Instead, status.String()=='%v'\n", status.String())
  }
}

func TestPreProcessPathCode_ParseString_01(t *testing.T) {

  testValue := "none"

  preProcCode, err := PreProcPathCode.ParseString(testValue, false)

  if err != nil {
    t.Errorf("Error returned by PreProcPathCode.ParseString(testValue, false)\n" +
      "testValue='%v'\nError='%v'\n", testValue, err.Error())
    return
  }

  if preProcCode != PreProcPathCode.None() {
    t.Errorf("Error: Expected preProcCode=='None'.\n" +
      "Instead, preProcCode='%v'\n", preProcCode.String())
  }
}

func TestPreProcessPathCode_ParseString_02(t *testing.T) {

  testValue := "pathseparator"

  preProcCode, err := PreProcPathCode.ParseString(testValue, false)

  if err != nil {
    t.Errorf("Error returned by PreProcPathCode.ParseString(testValue, false)\n" +
      "testValue='%v'\nError='%v'\n", testValue, err.Error())
    return
  }

  if preProcCode != PreProcPathCode.PathSeparator() {
    t.Errorf("Error: Expected preProcCode=='PathSeparator'.\n" +
      "Instead, preProcCode='%v'\n", preProcCode.String())
  }
}

func TestPreProcessPathCode_ParseString_03(t *testing.T) {

  testValue := "absolutepath"

  preProcCode, err := PreProcPathCode.ParseString(testValue, false)

  if err != nil {
    t.Errorf("Error returned by PreProcPathCode.ParseString(testValue, false)\n" +
      "testValue='%v'\nError='%v'\n", testValue, err.Error())
    return
  }

  if preProcCode != PreProcPathCode.AbsolutePath() {
    t.Errorf("Error: Expected preProcCode=='AbsolutePath'.\n" +
      "Instead, preProcCode='%v'\n", preProcCode.String())
  }

}

func TestPreProcessPathCode_ParseString_04(t *testing.T) {

  testValue := "None"

  preProcCode, err := PreProcPathCode.ParseString(testValue, true)

  if err != nil {
    t.Errorf("Error returned by PreProcPathCode.ParseString(testValue, false)\n" +
      "testValue='%v'\nError='%v'\n", testValue, err.Error())
    return
  }

  if preProcCode != PreProcPathCode.None() {
    t.Errorf("Error: Expected preProcCode=='None'.\n" +
      "Instead, preProcCode='%v'\n", preProcCode.String())
  }
}

func TestPreProcessPathCode_ParseString_05(t *testing.T) {

  testValue := "PathSeparator"

  preProcCode, err := PreProcPathCode.ParseString(testValue, true)

  if err != nil {
    t.Errorf("Error returned by PreProcPathCode.ParseString(testValue, false)\n" +
      "testValue='%v'\nError='%v'\n", testValue, err.Error())
    return
  }

  if preProcCode != PreProcPathCode.PathSeparator() {
    t.Errorf("Error: Expected preProcCode=='PathSeparator'.\n" +
      "Instead, preProcCode='%v'\n", preProcCode.String())
  }

}

func TestPreProcessPathCode_ParseString_06(t *testing.T) {

  testValue := "xxxxxx"

  _, err := PreProcPathCode.ParseString(testValue, false)

  if err == nil {
    t.Errorf("Expected an error return from PreProcPathCode.ParseString" +
      "(testValue, false)\n" +
      "because testValue='xxxxxx'.\n" +
      "However, NO ERROR WAS RETURNED!!\n",)
  }

}

func TestPreProcessPathCode_ParseString_07(t *testing.T) {

  statusStr := "xyz()"

  _, err := PreProcessPathCode(0).ParseString(statusStr, false)

  if err == nil {
    t.Error("Expected an error return from ParseString(statusStr, false)\n" +
      "because 'statusStr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestPreProcessPathCode_ParseString_08(t *testing.T) {

  statusStr := "absoluxxpath"

  _, err := PreProcessPathCode(0).ParseString(statusStr, false)

  if err == nil {
    t.Error("Expected an error return from ParseString(statusStr, false)\n" +
      "because 'statusStr' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}


func TestPreProcessPathCode_Value_01(t *testing.T) {

  statusCode1 := PreProcPathCode.PathSeparator()

  statusCode2 := statusCode1.Value()

  statusCode3 := PreProcPathCode.PathSeparator()

  if statusCode2 != statusCode3 {

    t.Errorf("Error: Expected statusCode2='%v'.\n" +
      "Instead, statusCode2='%v'\n", statusCode3, statusCode2)
  }

}
