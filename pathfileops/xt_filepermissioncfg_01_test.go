package pathfileops

import "testing"

func TestFilePermissionConfig_SetFileModeByTextCode_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_02(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_03(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_04(t *testing.T) {

	textCode := "drw-rw-rw-"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_05(t *testing.T) {

	textCode := "-rwx------"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_06(t *testing.T) {

	textCode := "-rwxrwx---"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_07(t *testing.T) {

	textCode := "---x--x--x"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_08(t *testing.T) {

	textCode := "--w--w--w-"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_09(t *testing.T) {

	textCode := "--wx-wx-wx"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_10(t *testing.T) {

	textCode := "-r--r--r--"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_11(t *testing.T) {

	textCode := "-r-xr-xr-x"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_12(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_13(t *testing.T) {

	textCode := "-rwxr-----"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_14(t *testing.T) {

	textCode := "drw-rw-rw-"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_15(t *testing.T) {

	textCode := "----------"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_SetFileModeByTextCode_16(t *testing.T) {

	textCode := "d---------"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fileMode, err := fpCfg.GetFileMode()

	if textCode != fileMode.String() {
		t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
			textCode, fileMode.String())
	}

}

func TestFilePermissionConfig_GetPermissionBits_01(t *testing.T) {

	textCode := "----------"
	expectedIntPermissionBits := 0

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	fMode, err := fpCfg.GetPermissionBits()

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	textBits, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, textBits)
	}

}
