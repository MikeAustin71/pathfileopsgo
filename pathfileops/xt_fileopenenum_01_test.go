package pathfileops

import (
	"os"
	"testing"
)

func TestFileOpenConfig_CopyIn_01(t *testing.T) {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
		FOpenMode.ModeAppend(), FOpenMode.ModeTruncate())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus2 := FileOpenConfig{}

	fOpStatus2.CopyIn(&fOpStatus1)

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}
}

func TestFileOpenConfig_CopyIn_02(t *testing.T) {

	fOpStatus1 := FileOpenConfig{}

	fOpStatus2 := FileOpenConfig{}

	fOpStatus2.CopyIn(&fOpStatus1)

	if !fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Error: Expected fOpStatus1==fOpStatus2. THEY ARE NOT EQUAL!")
	}
}

func TestFileOpenConfig_CopyOut_01(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus2 := fOpStatus1.CopyOut()

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}

}

func TestFileOpenConfig_CopyOut_02(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus2 := fOpStatus1.CopyOut()

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}

}

func TestFileOpenConfig_CopyOut_03(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus2 := fOpStatus1.CopyOut()

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}

}

func TestFileOpenConfig_CopyOut_04(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus1.fileOpenModes = make([]FileOpenMode, 0)

	fOpStatus2 := fOpStatus1.CopyOut()

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}

}

func TestFileOpenConfig_CopyOut_05(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.New(). Error='%v' \n", err.Error())
	}

	actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode {
		t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode)
	}

	fOpStatus1.fileOpenModes = nil

	fOpStatus2 := fOpStatus1.CopyOut()

	actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFOpenCode != actualFOpenCode2 {
		t.Errorf("Error: Expected File Open Code #2 ='%v'. Instead, "+
			"actual File Open Code='%v' \n",
			expectedFOpenCode, actualFOpenCode2)
	}

}

func TestFileOpenConfig_Equal_01(t *testing.T) {

	fOpStatus1, err :=
		FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fOpStatus2 := fOpStatus1.CopyOut()

	if !fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2==fOpStatus1. WRONG: They are NOT Equal!")
	}

	if !fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus1==fOpStatus2. WRONG: They are NOT Equal!")
	}

}

func TestFileOpenConfig_Equal_02(t *testing.T) {

	fOpStatus1, err :=
		FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fOpStatus2, err := FileOpenConfig{}.New(
		FOpenType.TypeWriteOnly(), FOpenMode.ModeAppend(), FOpenMode.ModeExclusive())

	if err != nil {
		t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	if fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2!=fOpStatus1. WRONG: They ARE Equal!")
	}

	if fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus1!=fOpStatus2. WRONG: They ARE Equal!")
	}

}

func TestFileOpenConfig_Equal_03(t *testing.T) {

	fOpStatus1 := FileOpenConfig{}

	fOpStatus2 := FileOpenConfig{}

	if !fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2==fOpStatus1. WRONG: They are NOT Equal!")
	}

	if !fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus ==fOpStatus2. WRONG: They are NOT Equal!")
	}

}

func TestFileOpenConfig_Equal_04(t *testing.T) {

	fOpStatus1, err :=
		FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fOpStatus2 := FileOpenConfig{}

	if fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2!=fOpStatus1. WRONG: They ARE Equal!")
	}

	if fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus ==fOpStatus2. WRONG: They ARE Equal!")
	}

}

func TestFileOpenConfig_Equal_05(t *testing.T) {

	fOpStatus1 := FileOpenConfig{}

	fOpStatus2, err :=
		FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	if fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2!=fOpStatus1. WRONG: They ARE Equal!")
	}

	if fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus ==fOpStatus2. WRONG: They ARE Equal!")
	}

}

func TestFileOpenConfig_Equal_06(t *testing.T) {

	fOpStatus1, err :=
		FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fOpStatus2, err :=
		FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	if fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2!=fOpStatus1. WRONG: They ARE Equal!")
	}

	if fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus ==fOpStatus2. WRONG: They ARE Equal!")
	}

}

func TestFileOpenConfig_Equal_07(t *testing.T) {

	fOpStatus1, err :=
		FileOpenConfig{}.New(
			FOpenType.TypeReadWrite(),
			FOpenMode.ModeAppend(),
			FOpenMode.ModeTruncate())

	if err != nil {
		t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fOpStatus2, err :=
		FileOpenConfig{}.New(
			FOpenType.TypeReadWrite(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeExclusive())

	if err != nil {
		t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	if fOpStatus2.Equal(&fOpStatus1) {
		t.Error("Expected fOpStatus2!=fOpStatus1. WRONG: They ARE Equal!")
	}

	if fOpStatus1.Equal(&fOpStatus2) {
		t.Error("Expected fOpStatus ==fOpStatus2. WRONG: They ARE Equal!")
	}

}

func TestFileOpenConfig_New_01(t *testing.T) {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
		FOpenMode.ModeAppend(), FOpenMode.ModeTruncate())

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

func TestFileOpenConfig_New_02(t *testing.T) {

	fOpenType := FileOpenType(-99)

	_, err := FileOpenConfig{}.New(fOpenType, FOpenMode.ModeCreate())

	if err == nil {
		t.Error("Expected Error returned by FileOpenConfig{}.New()" + "" +
			"because of an invalid File Open Type. However, NO ERROR WAS RETURNED! \n")
	}

}

func TestFileOpenConfig_New_03(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

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

func TestFileOpenConfig_New_04(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

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

func TestFileOpenConfig_New_05(t *testing.T) {

	fOpenMode := FileOpenMode(-99)

	_, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), fOpenMode)

	if err == nil {
		t.Error("Expected an error return from FileOpenConfig{}.New() " +
			"because the File Open Mode was invalid. NO ERROR RETURNED! \n")
	}

}

func TestFileOpenConfig_GetCompositeFileOpenCode_01(t *testing.T) {

	fOpCfg := FileOpenConfig{}

	_, err := fOpCfg.GetCompositeFileOpenCode()

	if err == nil {
		t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode() " +
			"because fOpCfg was NOT initialized. NO ERROR WAS RETURNED!")

	}

}

func TestFileOpenConfig_GetCompositeFileOpenCode_02(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	_, err = fOpCfg.GetCompositeFileOpenCode()

	if err == nil {
		t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode() " +
			"because fOpCfg use TypeNone(). However, NO ERROR WAS RETURNED!")

	}

}

func TestFileOpenConfig_GetCompositeFileOpenCode_03(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fOpCfg.fileOpenModes = nil

	_, err = fOpCfg.GetCompositeFileOpenCode()

	if err == nil {
		t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode() " +
			"because fOpCfg.fileOpenModes == nil . However, NO ERROR WAS RETURNED!")

	}

}

func TestFileOpenConfig_GetCompositeFileOpenCode_04(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fOpCfg.fileOpenType = FileOpenType(-99)

	_, err = fOpCfg.GetCompositeFileOpenCode()

	if err == nil {
		t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode() " +
			"because fOpCfg.fileOpenType is invalid. However, NO ERROR WAS RETURNED!")

	}

}

func TestFileOpenConfig_GetCompositeFileOpenCode_05(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fOpCfg.fileOpenModes = make([]FileOpenMode, 0)

	_, err = fOpCfg.GetCompositeFileOpenCode()

	if err == nil {
		t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode() " +
			"because fOpCfg.fileOpenModes has Zero Length. However, NO ERROR WAS RETURNED!")

	}

}

func TestFileOpenConfig_GetFileOpenModes_01(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(
		FOpenType.TypeReadWrite(),
		FOpenMode.ModeAppend(),
		FOpenMode.ModeCreate(),
		FOpenMode.ModeExclusive())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fileOpenModes := fOpCfg.GetFileOpenModes()

	if len(fileOpenModes) != 3 {
		t.Errorf("Error: Expected fileOpenModes len = 3. Instead, len='%v'",
			len(fileOpenModes))
	}

	hasAppend := 0
	hasCreate := 0
	hasExclusive := 0

	for i := 0; i < len(fileOpenModes); i++ {

		if fileOpenModes[i] == FOpenMode.ModeAppend() {
			hasAppend++
		}

		if fileOpenModes[i] == FOpenMode.ModeCreate() {
			hasCreate++
		}

		if fileOpenModes[i] == FOpenMode.ModeExclusive() {
			hasExclusive++
		}

	}

	if hasAppend != 1 {
		t.Errorf("Error: Could not locate correct number of Appends. "+
			"hasAppend='%v'", hasAppend)
	}

	if hasCreate != 1 {
		t.Errorf("Error: Could not locate correct number of Creates. "+
			"hasCreate='%v'", hasCreate)
	}

	if hasExclusive != 1 {
		t.Errorf("Error: Could not locate correct number of Exclusives. "+
			"hasExclusive='%v'", hasExclusive)
	}

}

func TestFileOpenConfig_GetFileOpenModes_02(t *testing.T) {

	fOpCfg, err := FileOpenConfig{}.New(
		FOpenType.TypeReadWrite())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	fileOpenModes := fOpCfg.GetFileOpenModes()

	if fileOpenModes == nil {
		t.Error("Error: Returned fileOpenModes is nil!")
	}

	if len(fileOpenModes) == 0 {
		t.Error("Error: Returned fileOpenModes has Zero Length!")
	}

	if len(fileOpenModes) != 1 {
		t.Errorf("Error: Returned fileOpenModes Length is NOT '1' ! "+
			"Length='%v' ", len(fileOpenModes))
	}

	if fileOpenModes[0] != FOpenMode.ModeNone() {
		t.Error("Error: Expected fileOpenModes[0] == FOpenMode.ModeNone(). " +
			"It is NOT!")
	}

}

func TestFileOpenConfig_SetFileOpenType_01(t *testing.T) {

	expectedFOpenCode := os.O_RDWR

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	err = fOpStatus.SetFileOpenType(FOpenType.TypeReadWrite())

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

	fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
	}

	err = fOpStatus.SetFileOpenType(FOpenType.TypeWriteOnly())

	if err != nil {
		t.Errorf("Error returned by SetFileOpenType{}.New(). Error='%v' \n", err.Error())
	}

	fOpStatus.SetFileOpenModes(FOpenMode.ModeAppend(), FOpenMode.ModeCreate())

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
