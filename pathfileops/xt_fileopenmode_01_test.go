package pathfileops

import (
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
