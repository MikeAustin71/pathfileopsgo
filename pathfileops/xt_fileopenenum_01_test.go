package pathfileops

import (
	"os"
	"testing"
)

func TestFileOpenConfig_New_01(t *testing.T) {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.WriteOnly(),
		FOpenMode.Append(), FOpenMode.Truncate())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

}

func TestFileOpenConfig_SetFileOpenType_01(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.None(),
		FOpenMode.None())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	err = fOpStatus.SetFileOpenType(FOpenType.ReadWrite())

	if err != nil {
		t.Errorf("Error returned by SetFileOpenType{}.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

}

func TestFileOpenConfig_SetFileOpenModes_01(t *testing.T) {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.None(),
		FOpenMode.None())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	err = fOpStatus.SetFileOpenType(FOpenType.WriteOnly())

	if err != nil {
		t.Errorf("Error returned by SetFileOpenType{}.New(). Error='%v' \n", err.Error())
	}

	fOpStatus.SetFileOpenModes(FOpenMode.Append(), FOpenMode.Create())

	actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

}

func TestFileOpenType_Value_01(t *testing.T) {

	fot := FOpenType.None()

	expected := -1

	if expected != fot.Value() {
		t.Errorf("Error: Expected File Open Type None=%v. Instead, "+
			"actual None Value ='%v'  ", expected, fot.Value())
	}

}

func TestFileOpenType_Value_02(t *testing.T) {

	fot := FOpenType.ReadOnly()

	expected := os.O_RDONLY

	if expected != fot.Value() {
		t.Errorf("Error: Expected File Open Type Read Only=%v. Instead, "+
			"actual Read Only Value ='%v'  ", expected, fot.Value())
	}

}

func TestFileOpenType_Value_03(t *testing.T) {

	fot := FOpenType.WriteOnly()

	expected := os.O_WRONLY

	if expected != fot.Value() {
		t.Errorf("Error: Expected File Open Type WriteOnly=%v. Instead, "+
			"actual Write Only Value ='%v'  ", expected, fot.Value())
	}

}

func TestFileOpenType_Value_04(t *testing.T) {

	fot := FOpenType.ReadWrite()

	expected := os.O_RDWR

	if expected != fot.Value() {
		t.Errorf("Error: Expected File Open Type Read/Write=%v. Instead, "+
			"actual Read/Write Value ='%v'  ", expected, fot.Value())
	}

}

func TestFileOpenType_String_01(t *testing.T) {

	expected := "None"

	fot := FOpenType.None()

	if expected != fot.String() {
		t.Errorf("Error: Expected File Open Type None String=%v. Instead, "+
			"actual File Open Type None String ='%v'  ", expected, fot.String())
	}

}

func TestFileOpenType_String_02(t *testing.T) {

	expected := "ReadOnly"

	fot := FOpenType.ReadOnly()

	if expected != fot.String() {
		t.Errorf("Error: Expected File Open Type ReadOnly String=%v. Instead, "+
			"actual File Open Type ReadOnly String ='%v'  ", expected, fot.String())
	}

}

func TestFileOpenType_String_03(t *testing.T) {

	expected := "WriteOnly"

	fot := FOpenType.WriteOnly()

	if expected != fot.String() {
		t.Errorf("Error: Expected File Open Type WriteOnly String=%v. Instead, "+
			"actual File Open Type WriteOnly String ='%v'  ", expected, fot.String())
	}

}

func TestFileOpenType_String_04(t *testing.T) {

	expected := "ReadWrite"

	fot := FOpenType.ReadWrite()

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

	fot := FOpenType.ReadWrite()

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

	if FOpenType.ReadWrite() != fot {
		t.Errorf("Error: Expected File Open Type Parse String to generate type "+
			"'ReadWrite'. Instead, it generated type='%v' ", fot.String())
	}

}
