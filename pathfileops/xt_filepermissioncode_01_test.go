package pathfileops

import (
	"os"
	"strconv"
	"testing"
)

func TestOsFilePermissionCode_IsValid_01(t *testing.T) {

	fpc := FilePermCode.ModeDir()

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

	fpc := FilePermCode.None()

	fmBase := os.FileMode(0)

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.None. It WAS NOT EQUAL!"+
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

	fpc := FilePermCode.ModeAppend()

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

	fpc := FilePermCode.ModeExclusive()

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

	fpc := FilePermCode.ModeTemporary()

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

	fpc := FilePermCode.ModeSymlink()

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

	fpc := FilePermCode.ModeDevice()

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

	fpc := FilePermCode.ModeNamedPipe()

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

	fpc := FilePermCode.ModeSocket()

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

	fpc := FilePermCode.ModeSetuid()

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

	fpc := FilePermCode.ModeSetgid()

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

	fpc := FilePermCode.ModeCharDevice()

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

	fpc := FilePermCode.ModeSticky()

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

	fpc := FilePermCode.ModeIrregular()

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

	osPerm := OsFilePermissionCode(FilePermCode.ModeDir())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v'. Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}
}

func TestOsFilePermissionCode_String_02(t *testing.T) {

	expectedStr := "ModeNamedPipe"

	osPerm := OsFilePermissionCode(FilePermCode.ModeNamedPipe())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v'. Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_03(t *testing.T) {

	expectedStr := "None"

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

	osPerm := OsFilePermissionCode(FilePermCode.ModeAppend())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_06(t *testing.T) {

	expectedStr := "ModeExclusive"

	osPerm := OsFilePermissionCode(FilePermCode.ModeExclusive())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_07(t *testing.T) {

	expectedStr := "ModeTemporary"

	osPerm := OsFilePermissionCode(FilePermCode.ModeTemporary())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_08(t *testing.T) {

	expectedStr := "ModeSymlink"

	osPerm := OsFilePermissionCode(FilePermCode.ModeSymlink())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_09(t *testing.T) {

	expectedStr := "ModeDevice"

	osPerm := OsFilePermissionCode(FilePermCode.ModeDevice())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_10(t *testing.T) {

	expectedStr := "ModeNamedPipe"

	osPerm := OsFilePermissionCode(FilePermCode.ModeNamedPipe())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_11(t *testing.T) {

	expectedStr := "ModeSocket"

	osPerm := OsFilePermissionCode(FilePermCode.ModeSocket())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_12(t *testing.T) {

	expectedStr := "ModeSetuid"

	osPerm := OsFilePermissionCode(FilePermCode.ModeSetuid())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_13(t *testing.T) {

	expectedStr := "ModeSetgid"

	osPerm := OsFilePermissionCode(FilePermCode.ModeSetgid())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_14(t *testing.T) {

	expectedStr := "ModeCharDevice"

	osPerm := OsFilePermissionCode(FilePermCode.ModeCharDevice())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_15(t *testing.T) {

	expectedStr := "ModeSticky"

	osPerm := OsFilePermissionCode(FilePermCode.ModeSticky())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_String_16(t *testing.T) {

	expectedStr := "ModeIrregular"

	osPerm := OsFilePermissionCode(FilePermCode.ModeIrregular())

	if expectedStr != osPerm.String() {
		t.Errorf("Error: Expected osPerm = '%v' . Instead, osPerm='%v' ",
			expectedStr, osPerm.String())
	}

}

func TestOsFilePermissionCode_ParseString_01(t *testing.T) {

	expectedFileMode := FilePermCode.ModeDir()

	actualFileMode, err := FilePermCode.ParseString("ModeDir", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modedir", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_02(t *testing.T) {

	expectedFileMode := FilePermCode.ModeAppend()

	actualFileMode, err := FilePermCode.ParseString("ModeAppend", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modeappend", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_03(t *testing.T) {

	expectedFileMode := FilePermCode.ModeExclusive()

	actualFileMode, err := FilePermCode.ParseString("ModeExclusive", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modeexclusive", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_04(t *testing.T) {

	expectedFileMode := FilePermCode.ModeTemporary()

	actualFileMode, err := FilePermCode.ParseString("modetemporary", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("ModeTemporary", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_05(t *testing.T) {

	expectedFileMode := FilePermCode.ModeSymlink()

	actualFileMode, err := FilePermCode.ParseString("modesymlink", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_06(t *testing.T) {

	expectedFileMode := FilePermCode.ModeDevice()

	actualFileMode, err := FilePermCode.ParseString("modedevice", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("ModeDevice", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_07(t *testing.T) {

	expectedFileMode := FilePermCode.ModeNamedPipe()

	actualFileMode, err := FilePermCode.ParseString("ModeNamedPipe", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modenamedpipe", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_08(t *testing.T) {

	expectedFileMode := FilePermCode.ModeSocket()

	actualFileMode, err := FilePermCode.ParseString("ModeSocket", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modesocket", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_09(t *testing.T) {

	expectedFileMode := FilePermCode.ModeSetuid()

	actualFileMode, err := FilePermCode.ParseString("ModeSetuid", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modesetuid", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_10(t *testing.T) {

	expectedFileMode := FilePermCode.ModeSetgid()

	actualFileMode, err := FilePermCode.ParseString("ModeSetgid", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modesetgid", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_11(t *testing.T) {

	expectedFileMode := FilePermCode.ModeCharDevice()

	actualFileMode, err := FilePermCode.ParseString("ModeCharDevice", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modechardevice", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_12(t *testing.T) {

	expectedFileMode := FilePermCode.ModeSticky()

	actualFileMode, err := FilePermCode.ParseString("ModeSticky", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modesticky", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_13(t *testing.T) {

	expectedFileMode := FilePermCode.ModeIrregular()

	actualFileMode, err := FilePermCode.ParseString("ModeIrregular", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("modeirregular", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_ParseString_14(t *testing.T) {

	expectedFileMode := FilePermCode.None()

	actualFileMode, err := FilePermCode.ParseString("None", true)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

	actualFileMode, err = FilePermCode.ParseString("none", false)

	if err != nil {
		t.Errorf("Error: Expected File Mode decimal value='%s'. Instead, Actual File Mode "+
			"decimal value='%s'", strconv.FormatInt(int64(expectedFileMode), 10),
			strconv.FormatInt(int64(actualFileMode), 10))
	}

}

func TestOsFilePermissionCode_GetFileModeLetterCode_01(t *testing.T) {
	expectedLetter := "-"

	fPerm := OsFilePermissionCode(FilePermCode.None())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeDir())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeAppend())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeExclusive())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeTemporary())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeSymlink())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeDevice())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeNamedPipe())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeSocket())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeSetuid())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeSetgid())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeCharDevice())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeSticky())

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

	fPerm := OsFilePermissionCode(FilePermCode.ModeIrregular())

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
	//     None                  -:      is a file

	letterCode := "-"

	expected := OsFilePermissionCode(FilePermCode.None())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeDir())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeAppend())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeExclusive())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeTemporary())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeSymlink())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeDevice())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeNamedPipe())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeSocket())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeSetuid())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeSetgid())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeCharDevice())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeSticky())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.ModeIrregular())

	fPerm, err := FilePermCode.GetNewFromLetterCode(letterCode)

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

	expected := OsFilePermissionCode(FilePermCode.None())

	fPerm := FilePermCode.None()

	if expected.Value() != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm), 10))
	}

}

func TestOsFilePermissionCode_Value_02(t *testing.T) {

	x := OsFilePermissionCode(99)

	x = OsFilePermissionCode(FilePermCode.ModeDir())

	fPerm := FilePermCode.ModeDir()

	if x.Value() != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(fPerm), 10),
			strconv.FormatInt(int64(x), 10))
	}
}

func TestOsFilePermissionCode_Value_03(t *testing.T) {

	x := OsFilePermissionCode(99)

	x = OsFilePermissionCode(FilePermCode.ModeAppend())

	fPerm := OsFilePermissionCode(FilePermCode.ModeAppend())

	if x != fPerm {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(x.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}
}
