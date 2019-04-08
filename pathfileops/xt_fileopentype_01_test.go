package pathfileops

import (
  "os"
  "testing"
)

func TestFileOpenType_Value_01(t *testing.T) {

  fot := FOpenType.TypeNone()

  expected := -1

  if expected != fot.Value() {
    t.Errorf("Error: Expected File Open Type None=%v. Instead, "+
      "actual None Value ='%v'  ", expected, fot.Value())
  }

}

func TestFileOpenType_Value_02(t *testing.T) {

  fot := FOpenType.TypeReadOnly()

  expected := os.O_RDONLY

  if expected != fot.Value() {
    t.Errorf("Error: Expected File Open Type Read Only=%v. Instead, "+
      "actual Read Only Value ='%v'  ", expected, fot.Value())
  }

}

func TestFileOpenType_Value_03(t *testing.T) {

  fot := FOpenType.TypeWriteOnly()

  expected := os.O_WRONLY

  if expected != fot.Value() {
    t.Errorf("Error: Expected File Open Type WriteOnly=%v. Instead, "+
      "actual Write Only Value ='%v'  ", expected, fot.Value())
  }

}

func TestFileOpenType_Value_04(t *testing.T) {

  fot := FOpenType.TypeReadWrite()

  expected := os.O_RDWR

  if expected != fot.Value() {
    t.Errorf("Error: Expected File Open Type Read/Write=%v. Instead, "+
      "actual Read/Write Value ='%v'  ", expected, fot.Value())
  }

}

func TestFileOpenType_String_01(t *testing.T) {

  expected := "TypeNone"

  fot := FOpenType.TypeNone()

  if expected != fot.String() {
    t.Errorf("Error: Expected File Open Type None String=%v. Instead, "+
      "actual File Open Type None String ='%v'  ", expected, fot.String())
  }

}

func TestFileOpenType_String_02(t *testing.T) {

  expected := "TypeReadOnly"

  fot := FOpenType.TypeReadOnly()

  if expected != fot.String() {
    t.Errorf("Error: Expected File Open Type ReadOnly String=%v. Instead, "+
      "actual File Open Type ReadOnly String ='%v'  ", expected, fot.String())
  }

}

func TestFileOpenType_String_03(t *testing.T) {

  expected := "TypeWriteOnly"

  fot := FOpenType.TypeWriteOnly()

  if expected != fot.String() {
    t.Errorf("Error: Expected File Open Type WriteOnly String=%v. Instead, "+
      "actual File Open Type WriteOnly String ='%v'  ", expected, fot.String())
  }

}

func TestFileOpenType_String_04(t *testing.T) {

  expected := "TypeReadWrite"

  fot := FOpenType.TypeReadWrite()

  if expected != fot.String() {
    t.Errorf("Error: Expected File Open Type ReadWrite String=%v. Instead, "+
      "actual File Open Type ReadWrite String ='%v'  ", expected, fot.String())
  }

}

func TestFileOpenType_IsValid_01(t *testing.T) {

  fot := FileOpenType(-99)

  err := fot.IsValid()

  if err == nil {
    t.Error("Expected Error from IsValid on FileOpenType(-99). NO ERROR WAS RECEIVED!")
  }

}

func TestFileOpenType_IsValid_02(t *testing.T) {

  fot := FOpenType.TypeReadWrite()

  err := fot.IsValid()

  if err != nil {
    t.Error("Expected: IsValid Error returned on VALID FileOpenType 'ReadWrite'")
  }

}

func TestFileOpenType_ParseString_01(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadWrite", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadWrite\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_02(t *testing.T) {

  _, err := FileOpenType(0).ParseString("Rx", true)

  if err == nil {
    t.Error("Expected an error return from FileOpenType(0).ParseString" +
      "(\"Rx\", true). However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOpenType_ParseString_03(t *testing.T) {

  _, err := FileOpenType(0).ParseString("Rxzyjx2v", true)

  if err == nil {
    t.Error("Expected an error return from FileOpenType(0).ParseString" +
      "(\"Rxzyjx2v\", true). However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOpenType_ParseString_04(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeReadWrite", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadWrite\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_05(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeReadOnly", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeReadOnly\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_06(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeReadOnly()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeReadOnly()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_07(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadOnly()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadOnly()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_08(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadOnly", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadOnly\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_09(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("WriteOnly", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"WriteOnly\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_10(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("WriteOnly()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"WriteOnly()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'WriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_11(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeWriteOnly()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeWriteOnly()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_12(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeWriteOnly", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeWriteOnly\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_13(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeNone()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeNone()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeNone() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeNone'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_14(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeNone", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeNone\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeNone() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeNone'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_15(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("None", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"None\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeNone() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'None'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_16(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typEreadwrite()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typEreadwrite()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_17(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("readwrite()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"readwrite()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_18(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("readwrite", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"readwrite\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_19(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typereadwrite", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typereadwrite\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'ReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_20(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typereadonly()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typereadonly()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_21(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("readonly()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"readonly()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_22(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typereadonly", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typereadonly\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_23(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("readonly", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"readonly\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_24(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typewriteonly()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typewriteonly()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_25(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("writeonly()", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"writeonly()\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_26(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("typewriteonly", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"typewriteonly\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_27(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("writeonly", false)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"writeonly\", false). Error='%v' ", err.Error())
  }

  if FOpenType.TypeWriteOnly() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeWriteOnly'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_28(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeReadWrite", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeReadWrite\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_29(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadWrite", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadWrite\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_30(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadWrite", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadWrite\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_31(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("ReadWrite()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"ReadWrite()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_32(t *testing.T) {

  fot, err := FileOpenType(0).ParseString("TypeReadWrite()", true)

  if err != nil {
    t.Errorf("Error returned from FileOpenType(0).ParseString"+
      "(\"TypeReadWrite()\", true). Error='%v' ", err.Error())
  }

  if FOpenType.TypeReadWrite() != fot {
    t.Errorf("Error: Expected File Open Type Parse String to generate type "+
      "'TypeReadWrite'. Instead, it generated type='%v' ", fot.String())
  }

}

func TestFileOpenType_ParseString_33(t *testing.T) {

  _, err := FileOpenType(0).ParseString("XYZ123456()", true)

  if err == nil {
    t.Error("Expected an error return from FileOpenType(0).ParseString" +
      "(\"XYZ123456()\", true). However, NO ERROR WAS RETURNED!")
  }

}

func TestFileOpenType_ParseString_34(t *testing.T) {

  _, err := FileOpenType(0).ParseString("XYZ123456()", false)

  if err == nil {
    t.Error("Expected an error return from FileOpenType(0).ParseString" +
      "(\"XYZ123456()\", false). However, NO ERROR WAS RETURNED!")
  }

}
