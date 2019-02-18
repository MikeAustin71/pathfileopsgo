package pathfileops

import (
	"os"
	"strconv"
	"testing"
)

func TestOsFilePermissionCode_IsValid_01(t *testing.T) {

	fpc := OsFilePermCode.ModeDir()

	if os.ModeDir != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDir. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_02(t *testing.T) {

	fpc := OsFilePermCode.ModeNone()

	fmBase := os.FileMode(0)

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeNone. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_03(t *testing.T) {

	fpc := OsFilePermCode.ModeAppend()

	fmBase := os.ModeAppend

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeAppend. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_04(t *testing.T) {

	fpc := OsFilePermCode.ModeExclusive()

	fmBase := os.ModeExclusive

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeExclusive. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_05(t *testing.T) {

	fpc := OsFilePermCode.ModeTemporary()

	fmBase := os.ModeTemporary

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeTemporary. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_06(t *testing.T) {

	fpc := OsFilePermCode.ModeSymlink()

	fmBase := os.ModeSymlink

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSymlink. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_07(t *testing.T) {

	fpc := OsFilePermCode.ModeDevice()

	fmBase := os.ModeDevice

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDevice. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_08(t *testing.T) {

	fpc := OsFilePermCode.ModeNamedPipe()

	fmBase := os.ModeNamedPipe

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeNamedPipe. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_09(t *testing.T) {

	fpc := OsFilePermCode.ModeSocket()

	fmBase := os.ModeSocket

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSocket. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_10(t *testing.T) {

	fpc := OsFilePermCode.ModeSetuid()

	fmBase := os.ModeSetuid

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSetuid. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_11(t *testing.T) {

	fpc := OsFilePermCode.ModeSetgid()

	fmBase := os.ModeSetgid

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSetgid. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_12(t *testing.T) {

	fpc := OsFilePermCode.ModeCharDevice()

	fmBase := os.ModeCharDevice

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeCharDevice. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_13(t *testing.T) {

	fpc := OsFilePermCode.ModeSticky()

	fmBase := os.ModeSticky

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSticky. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_14(t *testing.T) {

	fpc := OsFilePermCode.ModeIrregular()

	fmBase := os.ModeIrregular

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeIrregular. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_15(t *testing.T) {

	fpc := os.FileMode(999)

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err == nil {
		t.Error("Expected an Error returned from Invalid FileMode code 999. " +
			"NO ERROR WAS RETURNED!!!!")
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

func TestOsFilePermissionCode_GetFileModeLetterCode_01(t *testing.T) {
	expectedLetter := "-"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeNone())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_02(t *testing.T) {
	expectedLetter := "d"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeDir())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_03(t *testing.T) {
	//     ModeAppend            a:      append-only
	expectedLetter := "a"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_04(t *testing.T) {
	//     ModeExclusive         l:      exclusive use

	expectedLetter := "l"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_05(t *testing.T) {

	//     ModeTemporary         T:      temporary file; Plan 9 only

	expectedLetter := "T"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeTemporary())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_06(t *testing.T) {

	//     ModeSymlink           L:      symbolic link

	expectedLetter := "L"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSymlink())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_07(t *testing.T) {

	//     ModeDevice            D:      device file

	expectedLetter := "D"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeDevice())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_08(t *testing.T) {

	//     ModeNamedPipe         p:      named pipe (FIFO)

	expectedLetter := "p"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_09(t *testing.T) {

	//     ModeSocket            S:      Unix domain socket

	expectedLetter := "S"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSocket())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_10(t *testing.T) {

	//     ModeSetuid            u:      setuid

	expectedLetter := "u"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_11(t *testing.T) {

	//     ModeSetgid            g:      setgid

	expectedLetter := "g"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_12(t *testing.T) {

	//     ModeCharDevice        c:      Unix character device, when ModeDevice is set

	expectedLetter := "c"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeCharDevice())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_13(t *testing.T) {

	//     ModeSticky            t:      sticky

	expectedLetter := "t"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_14(t *testing.T) {

	//     ModeIrregular         ?:      non-regular file; nothing else is known about this file

	expectedLetter := "?"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetNewFromLetterCode_01(t *testing.T) {
	//     ModeNone                  -:      is a file

	letterCode := "-"

	expected := OsFilePermissionCode(OsFilePermCode.ModeNone())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_02(t *testing.T) {
	//     ModeDir               d:      is a directory

	letterCode := "d"

	expected := OsFilePermissionCode(OsFilePermCode.ModeDir())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_03(t *testing.T) {

	//     ModeAppend            a:      append-only

	letterCode := "a"

	expected := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_04(t *testing.T) {

	//     ModeExclusive         l:      exclusive use

	letterCode := "l"

	expected := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_05(t *testing.T) {

	//     ModeTemporary         T:      temporary file; Plan 9 only

	letterCode := "T"

	expected := OsFilePermissionCode(OsFilePermCode.ModeTemporary())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_06(t *testing.T) {

	//     ModeSymlink           L:      symbolic link

	letterCode := "L"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSymlink())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_07(t *testing.T) {

	//     ModeDevice            D:      device file

	letterCode := "D"

	expected := OsFilePermissionCode(OsFilePermCode.ModeDevice())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_08(t *testing.T) {

	//     ModeNamedPipe         p:      named pipe (FIFO)

	letterCode := "p"

	expected := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_09(t *testing.T) {

	//     ModeSocket            S:      Unix domain socket

	letterCode := "S"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSocket())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_10(t *testing.T) {

	//     ModeSetuid            u:      setuid

	letterCode := "u"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_11(t *testing.T) {

	//     ModeSetgid            g:      setgid

	letterCode := "g"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_12(t *testing.T) {

	//     ModeCharDevice        c:      Unix character device, when ModeDevice is set

	letterCode := "c"

	expected := OsFilePermissionCode(OsFilePermCode.ModeCharDevice())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_13(t *testing.T) {

	//     ModeSticky            t:      sticky

	letterCode := "t"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_14(t *testing.T) {

	//     ModeIrregular         ?:      non-regular file; nothing else is known about this file

	letterCode := "?"

	expected := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
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
