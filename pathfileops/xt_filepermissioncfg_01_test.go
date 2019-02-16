package pathfileops

import (
	"strings"
	"testing"
)

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_02(t *testing.T) {

	textCode := "-rwx------"
	expectedIntPermissionBits := 700

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_03(t *testing.T) {

	textCode := "-rwxrwx---"
	expectedIntPermissionBits := 770

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_04(t *testing.T) {

	textCode := "-rwxrwxrwx"
	expectedIntPermissionBits := 777

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_05(t *testing.T) {

	textCode := "---x--x--x"
	expectedIntPermissionBits := 111

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_06(t *testing.T) {

	textCode := "--w--w--w-"
	expectedIntPermissionBits := 222

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_07(t *testing.T) {

	textCode := "--wx-wx-wx"
	expectedIntPermissionBits := 333

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_08(t *testing.T) {

	textCode := "-r--r--r--"
	expectedIntPermissionBits := 444

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_09(t *testing.T) {

	textCode := "-r-xr-xr-x"
	expectedIntPermissionBits := 555

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_10(t *testing.T) {

	textCode := "-rw-rw-rw-"
	expectedIntPermissionBits := 666

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_11(t *testing.T) {

	textCode := "-rwxr-----"
	expectedIntPermissionBits := 740

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_12(t *testing.T) {

	textCode := "drw-rw-rw-"
	expectedIntPermissionBits := 666

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_13(t *testing.T) {

	textCode := "drwxrwxrwx"
	expectedIntPermissionBits := 777

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}

}

func TestFilePermissionConfig_GetPermissionBits_14(t *testing.T) {

	textCode := "d---------"

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

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != textBits {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"permission text string='%v' ", expectedPermissionTxt, textBits)
	}
}
