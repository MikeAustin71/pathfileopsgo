package pathfileops

import (
	"strconv"
	"testing"
)

func TestOsFilePermissionCode_ParseString_01(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeDir()

	actualFileMode, err := OsFilePermCode.ParseString("ModeDir", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modedir", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_02(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeAppend()

	actualFileMode, err := OsFilePermCode.ParseString("ModeAppend", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modeappend", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_03(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeExclusive()

	actualFileMode, err := OsFilePermCode.ParseString("ModeExclusive", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modeexclusive", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_04(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeTemporary()

	actualFileMode, err := OsFilePermCode.ParseString("modetemporary", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("ModeTemporary", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_05(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSymlink()

	actualFileMode, err := OsFilePermCode.ParseString("modesymlink", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_06(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeDevice()

	actualFileMode, err := OsFilePermCode.ParseString("modedevice", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("ModeDevice", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_07(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeNamedPipe()

	actualFileMode, err := OsFilePermCode.ParseString("ModeNamedPipe", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modenamedpipe", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_08(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSocket()

	actualFileMode, err := OsFilePermCode.ParseString("ModeSocket", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modesocket", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_09(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetuid()

	actualFileMode, err := OsFilePermCode.ParseString("ModeSetuid", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modesetuid", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_10(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetgid()

	actualFileMode, err := OsFilePermCode.ParseString("ModeSetgid", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modesetgid", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_11(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeCharDevice()

	actualFileMode, err := OsFilePermCode.ParseString("ModeCharDevice", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modechardevice", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_12(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSticky()

	actualFileMode, err := OsFilePermCode.ParseString("ModeSticky", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modesticky", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_13(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeIrregular()

	actualFileMode, err := OsFilePermCode.ParseString("ModeIrregular", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("modeirregular", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_14(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeNone()

	actualFileMode, err := OsFilePermCode.ParseString("ModeNone", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = OsFilePermCode.ParseString("none", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_15(t *testing.T) {

	_, err := OsFilePermCode.ParseString("jUNO924", false)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"jUNO924\", false) because 'jUNO924' is an invalid text code. " +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_ParseString_16(t *testing.T) {

	_, err := OsFilePermCode.ParseString("jUNO924", true)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"jUNO924\", true) because 'jUNO924' is an invalid text code. " +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_String_01(t *testing.T) {

	expectedStr := "ModeDir"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeDir())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v'. Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}
}

func TestOsFilePermissionCode_String_02(t *testing.T) {

	expectedStr := "ModeNamedPipe"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v'. Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_03(t *testing.T) {

	expectedStr := "ModeNone"

	osPerm := OsFilePermissionCode(0)

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v'. Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_04(t *testing.T) {

	expectedStr := ""

	osPerm := OsFilePermissionCode(999)

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = 'Empty String\"\"'. Instead, osPerm='%v' ",
			osPerm.String())
	}

}

func TestOsFilePermissionCode_String_05(t *testing.T) {

	expectedStr := "ModeAppend"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_06(t *testing.T) {

	expectedStr := "ModeExclusive"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_07(t *testing.T) {

	expectedStr := "ModeTemporary"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeTemporary())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_08(t *testing.T) {

	expectedStr := "ModeSymlink"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeSymlink())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_09(t *testing.T) {

	expectedStr := "ModeDevice"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeDevice())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_10(t *testing.T) {

	expectedStr := "ModeNamedPipe"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_11(t *testing.T) {

	expectedStr := "ModeSocket"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeSocket())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_12(t *testing.T) {

	expectedStr := "ModeSetuid"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_13(t *testing.T) {

	expectedStr := "ModeSetgid"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_14(t *testing.T) {

	expectedStr := "ModeCharDevice"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeCharDevice())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_15(t *testing.T) {

	expectedStr := "ModeSticky"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_16(t *testing.T) {

	expectedStr := "ModeIrregular"

	osPerm := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_Value_01(t *testing.T) {

	expected := OsFilePermissionCode(OsFilePermCode.ModeNone())

	fPerm := OsFilePermCode.ModeNone()

	if expected.Value() != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm), 10))
	}

}

func TestOsFilePermissionCode_Value_02(t *testing.T) {

	x := OsFilePermissionCode(99)

	x = OsFilePermissionCode(OsFilePermCode.ModeDir())

	fPerm := OsFilePermCode.ModeDir()

	if x.Value() != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(fPerm), 10),
			strconv.FormatInt(int64(x), 10))
	}
}

func TestOsFilePermissionCode_Value_03(t *testing.T) {

	x := OsFilePermissionCode(99)

	x = OsFilePermissionCode(OsFilePermCode.ModeAppend())

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	if x != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(x.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}
}
