package pathfileops

import (
	"os"
	"strconv"
	"testing"
)

func TestFileOpenMode_IsValid_01(t *testing.T) {

	mode := FOpenMode.ModeAppend()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeAppend() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_02(t *testing.T) {

	mode := FOpenMode.ModeTruncate()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeTruncate() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_03(t *testing.T) {

	mode := FOpenMode.ModeCreate()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeCreate() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_04(t *testing.T) {

	mode := FOpenMode.ModeExclusive()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeExclusive() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_05(t *testing.T) {

	mode := FOpenMode.ModeSync()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeSync() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_06(t *testing.T) {

	mode := FOpenMode.ModeNone()

	err := mode.IsValid()

	if err != nil {
		t.Errorf("Error: Expected FOpenMode.ModeNone() would be treated as VALID! "+
			"Instead, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFileOpenMode_IsValid_07(t *testing.T) {

	mode := FileOpenMode(-999)

	err := mode.IsValid()

	if err == nil {
		t.Error("Expected an error return from mode.IsValid() " +
			"because mode = FileOpenMode(-999). Instead, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_01(t *testing.T) {

	inputStr := "ModeNone"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_02(t *testing.T) {

	inputStr := "None"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_03(t *testing.T) {

	inputStr := "ModeNone()"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_04(t *testing.T) {

	inputStr := "None()"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_05(t *testing.T) {

	inputStr := "modenone"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_06(t *testing.T) {

	inputStr := "none"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_07(t *testing.T) {

	inputStr := "modenone()"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_08(t *testing.T) {

	inputStr := "none()"
	expectedModeVal := FOpenMode.ModeNone()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_09(t *testing.T) {

	inputStr := "ModeSync"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_10(t *testing.T) {

	inputStr := "Sync"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_11(t *testing.T) {

	inputStr := "ModeSync()"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_12(t *testing.T) {

	inputStr := "Sync()"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_13(t *testing.T) {

	inputStr := "modesync"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_14(t *testing.T) {

	inputStr := "sync"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_15(t *testing.T) {

	inputStr := "modesync()"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_16(t *testing.T) {

	inputStr := "sync()"
	expectedModeVal := FOpenMode.ModeSync()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_17(t *testing.T) {

	inputStr := "ModeExclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_18(t *testing.T) {

	inputStr := "Exclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_19(t *testing.T) {

	inputStr := "ModeExclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_20(t *testing.T) {

	inputStr := "Exclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_21(t *testing.T) {

	inputStr := "modeexclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_22(t *testing.T) {

	inputStr := "exclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_23(t *testing.T) {

	inputStr := "modeexclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_24(t *testing.T) {

	inputStr := "exclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_25(t *testing.T) {

	inputStr := "ModeExclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_26(t *testing.T) {

	inputStr := "Exclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_27(t *testing.T) {

	inputStr := "ModeExclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_28(t *testing.T) {

	inputStr := "Exclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_29(t *testing.T) {

	inputStr := "modeexclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_30(t *testing.T) {

	inputStr := "exclusive"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_31(t *testing.T) {

	inputStr := "modeexclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_32(t *testing.T) {

	inputStr := "exclusive()"
	expectedModeVal := FOpenMode.ModeExclusive()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_33(t *testing.T) {

	inputStr := "ModeCreate"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_34(t *testing.T) {

	inputStr := "Create"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_35(t *testing.T) {

	inputStr := "ModeCreate()"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_36(t *testing.T) {

	inputStr := "Create()"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_37(t *testing.T) {

	inputStr := "modecreate"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_38(t *testing.T) {

	inputStr := "create"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_39(t *testing.T) {

	inputStr := "modecreate()"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_40(t *testing.T) {

	inputStr := "create()"
	expectedModeVal := FOpenMode.ModeCreate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_41(t *testing.T) {

	inputStr := "ModeTruncate"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_42(t *testing.T) {

	inputStr := "Truncate"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_43(t *testing.T) {

	inputStr := "ModeTruncate()"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_44(t *testing.T) {

	inputStr := "Truncate()"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_45(t *testing.T) {

	inputStr := "modetruncate"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_46(t *testing.T) {

	inputStr := "truncate"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_47(t *testing.T) {

	inputStr := "modetruncate()"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_48(t *testing.T) {

	inputStr := "truncate()"
	expectedModeVal := FOpenMode.ModeTruncate()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_49(t *testing.T) {

	inputStr := "ModeAppend"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_50(t *testing.T) {

	inputStr := "Append"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_51(t *testing.T) {

	inputStr := "ModeAppend()"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_52(t *testing.T) {

	inputStr := "Append()"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, true)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, true). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_53(t *testing.T) {

	inputStr := "modeappend"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_54(t *testing.T) {

	inputStr := "append"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_55(t *testing.T) {

	inputStr := "modeappend()"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_56(t *testing.T) {

	inputStr := "append()"
	expectedModeVal := FOpenMode.ModeAppend()

	actualModeVal, err := FOpenMode.ParseString(inputStr, false)

	if err != nil {
		t.Errorf("Error returned by FOpenMode.ParseString(inputStr, false). "+
			"inputStr='%v' Error='%v' ", inputStr, err.Error())
	}

	if expectedModeVal != actualModeVal {
		t.Errorf("Expected returned mode value='%v'. Instead, actual mode octal value='%s' ",
			expectedModeVal.String(), strconv.FormatInt(int64(actualModeVal), 8))
	}

}

func TestFileOpenMode_ParseString_57(t *testing.T) {

	inputStr := "Ap"

	_, err := FOpenMode.ParseString(inputStr, true)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, true). " +
			"inputStr='ap' Less than 3-chars. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_58(t *testing.T) {

	inputStr := "ap"

	_, err := FOpenMode.ParseString(inputStr, false)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, false). " +
			"inputStr='ap' Less than 3-chars. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_59(t *testing.T) {

	inputStr := "Ap"

	_, err := FOpenMode.ParseString(inputStr, true)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, true). " +
			"inputStr='ap' Less than 3-chars. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_60(t *testing.T) {

	inputStr := "ap"

	_, err := FOpenMode.ParseString(inputStr, false)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, false). " +
			"inputStr='ap' Less than 3-chars. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_61(t *testing.T) {

	inputStr := "Apxyzu"

	_, err := FOpenMode.ParseString(inputStr, true)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, true) " +
			"because 'inputStr' is invalid. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_ParseString_62(t *testing.T) {

	inputStr := "apxyzu"

	_, err := FOpenMode.ParseString(inputStr, false)

	if err == nil {
		t.Error("Expected an error return from FOpenMode.ParseString(inputStr, false) " +
			"because 'inputStr' is invalid. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileOpenMode_String_01(t *testing.T) {

	fileOpenMode := FOpenMode.ModeTruncate()
	expectedStr := "ModeTruncate"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_02(t *testing.T) {

	fileOpenMode := FOpenMode.ModeCreate()
	expectedStr := "ModeCreate"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_03(t *testing.T) {

	fileOpenMode := FOpenMode.ModeExclusive()
	expectedStr := "ModeExclusive"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_04(t *testing.T) {

	fileOpenMode := FOpenMode.ModeSync()
	expectedStr := "ModeSync"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_05(t *testing.T) {

	fileOpenMode := FOpenMode.ModeAppend()
	expectedStr := "ModeAppend"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_06(t *testing.T) {

	fileOpenMode := FOpenMode.ModeNone()
	expectedStr := "ModeNone"

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_String_07(t *testing.T) {

	fileOpenMode := FileOpenMode(999)
	expectedStr := ""

	actualStr := fileOpenMode.String()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected String()='%v'. Instead, String()='%v'",
			expectedStr, actualStr)
	}

}

func TestFileOpenMode_Value_01(t *testing.T) {

	fileOpenMode := FOpenMode.ModeNone()

	expectedValue := -1

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_02(t *testing.T) {

	fileOpenMode := FOpenMode.ModeAppend()

	expectedValue := os.O_APPEND

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_03(t *testing.T) {

	fileOpenMode := FOpenMode.ModeSync()

	expectedValue := os.O_SYNC

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_04(t *testing.T) {

	fileOpenMode := FOpenMode.ModeExclusive()

	expectedValue := os.O_EXCL

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_05(t *testing.T) {

	fileOpenMode := FOpenMode.ModeCreate()

	expectedValue := os.O_CREATE

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_06(t *testing.T) {

	fileOpenMode := FOpenMode.ModeTruncate()

	expectedValue := os.O_TRUNC

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}

func TestFileOpenMode_Value_07(t *testing.T) {

	expectedValue := 9999

	fileOpenMode := FileOpenMode(expectedValue)

	if expectedValue != fileOpenMode.Value() {
		t.Errorf("Expected ModeNone value='%v'. Instead, value='%v'",
			expectedValue, fileOpenMode.Value())
	}

}
