package pathfileops

import "testing"

func TestFileSelectionCriteria_ANDSelect_01(t *testing.T) {

  fsc := FileSelectionCriteria{}

  // DEFAULT SHOULD BE FileSelectMode.ANDSelect()
  if fsc.SelectCriterionMode != FileSelectMode.None() {
    t.Errorf("Expected default FileSelectionCriteria.SelectCriterionMode="+
      "FileSelectMode.None(). Instead, FileSelectionCriteria.SelectCriterionMode='%v'",
      fsc.SelectCriterionMode)
  }

}

func TestFileSelectCriterionMode_ParseString_01(t *testing.T) {

  _, err := FileSelectCriterionMode(0).ParseString("xy", true)

  if err == nil {
    t.Error("Expected an error return from FileSelectCriterionMode(0)." +
      "ParseString(\"xy\", true)\n" +
      "because 'xy' is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestFileSelectCriterionMode_StatusIsValid_01(t *testing.T) {
  invalidFileSelectCriteria := FileSelectCriterionMode(-99)

  err := invalidFileSelectCriteria.StatusIsValid()

  if err == nil {
    t.Error("Expected an error return from invalidFileSelectCriteria.StatusIsValid()\n" +
      "because the FileSelectCriterionMode is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

}

func TestFileSelectCriterionMode_StatusIsValid_02(t *testing.T) {
  validFileSelectCriteria := FileSelectCriterionMode(0).ANDSelect()

  err := validFileSelectCriteria.StatusIsValid()

  if err != nil {
    t.Errorf("Error returned by validFileSelectCriteria.StatusIsValid()\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestFileSelectCriterionMode_String_01(t *testing.T) {

  fsc := FileSelectCriterionMode(-99)

  fscStr := fsc.String()

  if fscStr != "" {
    t.Errorf("ERROR: Expected that fsc.String() would return and empty string.\n" +
      "because 'fsc' is invalid.\n" +
      "However, fsc.String()='%v'\n", fscStr )
  }

}

func TestFileSelectCriterionMode_Text_01(t *testing.T) {
  var r FileSelectCriterionMode

  r = FileSelectMode.None()

  var s string

  s = r.String()

  if s != "None" {
    t.Errorf("Expected string 'None'. Instead, got %v", s)
  }
}

func TestFileSelectCriterionMode_Text_02(t *testing.T) {
  var r FileSelectCriterionMode

  r = FileSelectMode.ANDSelect()

  var s string

  s = r.String()

  if s != "ANDSelect" {
    t.Errorf("Expected string 'ANDSelect'. Instead, got %v", s)
  }
}

func TestFileSelectCriterionMode_Text_03(t *testing.T) {
  var r FileSelectCriterionMode

  r = FileSelectMode.ORSelect()

  var s string

  s = r.String()

  if s != "ORSelect" {
    t.Errorf("Expected string 'ORSelect'. Instead, got %v", s)
  }
}

func TestFileSelectCriterionMode_Text_04(t *testing.T) {

  s := "None"

  r, err := FileSelectMode.ParseString(s, true)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, false). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if s != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", s, r.String())
  }
}

func TestFileSelectCriterionMode_Text_05(t *testing.T) {

  s := "ANDSelect"

  r, err := FileSelectMode.ParseString(s, true)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, false). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if s != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", s, r.String())
  }
}

func TestFileSelectCriterionMode_Text_06(t *testing.T) {

  s := "ORSelect"

  r, err := FileSelectMode.ParseString(s, true)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, true). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if s != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", s, r.String())
    return
  }
}

func TestFileSelectCriterionMode_Text_07(t *testing.T) {

  s := "XXXXX"

  _, err := FileSelectMode.ParseString(s, true)

  if err == nil {
    t.Error("Error: Expected an error return from ParseString() " +
      "NO ERROR WAS RETURNED!")
  }

}

func TestFileSelectCriterionMode_Text_08(t *testing.T) {

  expectedStr := "None"
  s := "none"

  r, err := FileSelectMode.ParseString(s, false)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, false). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if expectedStr != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", expectedStr, r.String())
    return
  }
}

func TestFileSelectCriterionMode_Text_09(t *testing.T) {

  expectedStr := "ANDSelect"
  s := "andseLect"

  r, err := FileSelectMode.ParseString(s, false)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, false). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if expectedStr != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", expectedStr, r.String())
  }
}

func TestFileSelectCriterionMode_Text_10(t *testing.T) {

  expectedStr := "ORSelect"
  s := "orseleCt"

  r, err := FileSelectMode.ParseString(s, false)

  if err != nil {
    t.Errorf("Error returned by FileSelectMode.ParseString(s, false). "+
      "s='%v' Error='%v' ", s, err.Error())
    return
  }

  if expectedStr != r.String() {
    t.Errorf("Expected string '%v'. Instead, got %v", expectedStr, r.String())
  }
}

func TestFileSelectCriterionMode_Text_11(t *testing.T) {

  s := "ANDSelxct"

  _, err := FileSelectMode.ParseString(s, false)

  if err == nil {
    t.Error("Error: Expected an error return from ParseString() " +
      "NO ERROR WAS RETURNED!")
  }

}

func TestFileSelectCriterionMode_Value_01(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.None()

  i = int(r)

  if i != 0 {
    t.Errorf("Expected 'FileSelectMode.None()' value = 0. Instead, got %v", i)
  }
}

func TestFileSelectCriterionMode_Value_02(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.ANDSelect()

  i = int(r)

  if i != 1 {
    t.Errorf("Expected 'FileSelectMode.ANDSelect()' value = 1. Instead, got %v", i)
  }
}

func TestFileSelectCriterionMode_Value_03(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.ORSelect()

  i = int(r)

  if i != 2 {
    t.Errorf("Expected 'FileSelectMode.ORSelect()' value = 2. Instead, got %v", i)
  }
}

func TestFileSelectCriterionMode_Value_04(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.None()

  i = int(r)

  if i != int(r.Value()) {
    t.Errorf("Expected 'FileSelectMode.None()' value = 0. Instead, got %v", i)
  }
}

func TestFileSelectCriterionMode_Value_05(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.ANDSelect()

  i = int(r)

  if i != int(r.Value()) {
    t.Errorf("Expected 'FileSelectMode.ANDSelect()' value = 1. Instead, got %v", i)
  }
}

func TestFileSelectCriterionMode_Value_06(t *testing.T) {
  var r FileSelectCriterionMode

  var i int

  r = FileSelectMode.ORSelect()

  i = int(r)

  if i != int(r.Value()) {
    t.Errorf("Expected 'FileSelectMode.ORSelect()' value = 2. Instead, got %v", i)
  }
}

