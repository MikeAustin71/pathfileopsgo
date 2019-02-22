package pathfileops

import (
	"strconv"
	"strings"
	"testing"
)

func TestOsFilePermissionCode_ParseString_01(t *testing.T) {

	inputStrs := []string {"ModeDir", "ModeDir()", "Dir", "Dir()" }

	expectedFileMode := OsFilePermCode.ModeDir()
	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_02(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeAppend()
	inputStrs := []string {"ModeAppend", "ModeAppend()", "Append", "Append()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_03(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeExclusive()
	inputStrs := []string {"ModeExclusive", "ModeExclusive()", "Exclusive", "Exclusive()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_04(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeTemporary()
	inputStrs := []string {"ModeTemporary", "ModeTemporary()", "Temporary", "Temporary()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_05(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSymlink()
	inputStrs := []string {"ModeSymlink", "ModeSymlink()", "Symlink", "Symlink()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_06(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeDevice()
	inputStrs := []string {"ModeDevice", "ModeDevice()", "Device", "Device()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_07(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeNamedPipe()
	inputStrs := []string {"ModeNamedPipe", "ModeNamedPipe()", "NamedPipe", "NamedPipe()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_08(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSocket()
	inputStrs := []string {"ModeSocket", "ModeSocket()", "Socket", "Socket()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}
}

func TestOsFilePermissionCode_ParseString_09(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetuid()
	inputStrs := []string {"ModeSetuid", "ModeSetuid()", "Setuid", "Setuid()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}
}

func TestOsFilePermissionCode_ParseString_10(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetgid()
	inputStrs := []string {"ModeSetgid", "ModeSetgid()", "Setgid", "Setgid()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_11(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeCharDevice()
	inputStrs := []string {"ModeCharDevice", "ModeCharDevice()", "CharDevice", "CharDevice()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_12(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSticky()
	inputStrs := []string {"ModeSticky", "ModeSticky()", "Sticky", "Sticky()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}

}

func TestOsFilePermissionCode_ParseString_13(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeIrregular()
	inputStrs := []string {"ModeIrregular", "ModeIrregular()", "Irregular", "Irregular()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}
}

func TestOsFilePermissionCode_ParseString_14(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeNone()
	inputStrs := []string {"ModeNone", "ModeNone()", "None", "None()" }

	actualFileMode:= OsFilePermissionCode(0)
	var err error

	for i:=0; i<len(inputStrs); i++ {

		actualFileMode, err = OsFilePermCode.ParseString(inputStrs[i], true)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(inputStrs[i], true). " +
				"index='%v' inputStr='%v' Error='%v' ", i, inputStrs[i], err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' inputStrs[%v]='%v' ", strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, inputStrs[i])
		}

		lwrCase := strings.ToLower(inputStrs[i])

		actualFileMode, err = OsFilePermCode.ParseString(lwrCase, false)

		if err != nil {
			t.Errorf("Error returned by OsFilePermCode.ParseString(lwrCase, false). " +
				"index='%v' lwrCase='%v' Error='%v' ", i, lwrCase, err.Error())
		}

		if expectedFileMode != actualFileMode.Value() {
			t.Errorf("Error: Expected File Mode ocatal value='%s'. Instead, Actual File Mode "+
				"ocatal value='%s' index='%v' lwrCase='%v' ",
				strconv.FormatInt(int64(expectedFileMode), 8),
				strconv.FormatInt(int64(actualFileMode), 8), i, lwrCase)
		}

	}
}

func TestOsFilePermissionCode_ParseString_15(t *testing.T) {

	_, err := OsFilePermCode.ParseString("Jun924", true)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"Jun924\", true) because 'Jun924' is an invalid text code. " +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_ParseString_16(t *testing.T) {

	_, err := OsFilePermCode.ParseString("Ju", true)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"Ju\", true) because 'Ju' is less than 3-characters and is " +
			"therefore invalid. However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_ParseString_17(t *testing.T) {

	_, err := OsFilePermCode.ParseString("jun924", false)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"jun924\", false) because 'jun924' is an invalid text code. " +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_ParseString_18(t *testing.T) {

	_, err := OsFilePermCode.ParseString("ju", false)

	if err == nil {
		t.Error("Expected an error to be returned from  OsFilePermCode." +
			"ParseString(\"ju\", false) because 'ju' is less than 3-characters and is " +
			"therefore invalid. However, NO ERROR WAS RETURNED!")
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
