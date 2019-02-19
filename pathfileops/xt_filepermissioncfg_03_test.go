package pathfileops

import (
	"os"
	"testing"
)

func TestFilePermissionConfig_IsValid_01(t *testing.T) {
	// expectedTextCode := "drwxrwxrwx"
	fh := FileHelper{}

	// drwxrwxrwx   20000000777

	intFMode := fh.ConvertOctalToDecimal(20000000777)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	fPerm.isInitialized = false

	err = fPerm.IsValid()

	if err == nil {
		t.Error("Expected an error to be returned by fPerm.IsValid() because " +
			"fPerm has not been initialized. NO ERROR RETURNED!")
	}

}

func TestFilePermissionConfig_IsValid_02(t *testing.T) {

	// expectedTextCode := "drwxrwxrwx"
	fh := FileHelper{}

	// drwxrwxrwx   20000000777

	intFMode := fh.ConvertOctalToDecimal(20000000777)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	fPerm.fileMode = os.FileMode(01000)

	err = fPerm.IsValid()

	if err == nil {
		t.Error("Expected an error to be returned by fPerm.IsValid() because " +
			"fPerm has an INVALID Entry Type. NO ERROR RETURNED!")
	}

}

func TestFilePermissionConfig_IsValid_03(t *testing.T) {

	// expectedTextCode := "drwxrwxrwx"
	fh := FileHelper{}

	// drwxrwxrwx   20000000777

	intFMode := fh.ConvertOctalToDecimal(20000000777)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	err = fPerm.IsValid()

	if err != nil {
		t.Errorf("Expected no error to be returned by fPerm.IsValid(). "+
			"However, an error was returned. Error='%v' ", err.Error())
	}

}

func TestFilePermissionConfig_NewByComponents_01(t *testing.T) {

	entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "-rwxrwxrwx"
	permissionStr := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
			"permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_NewByComponents_02(t *testing.T) {

	entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeDir()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "drwxrwxrwx"
	permissionStr := "rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
			"permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_NewByComponents_03(t *testing.T) {

	//  ModeSetuid()      os.ModeSetuid         "u" setuid
	entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeSetuid())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeSetuid()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "urw-rw-rw-"
	permissionStr := "rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
			"permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_NewByComponents_04(t *testing.T) {

	// Bad Entry Type Code
	entryType := OsFilePermissionCode(999)

	permissionStr := "rw-rw-rw-"

	_, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

	if err == nil {
		t.Error("Expected error return from bad entry type code 999. " +
			"However, NO ERROR WAS RETURNED! ")
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_01(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "-rwxrwxrwx"
	permissionStr := "-rwxrwxrwx"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_02(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeDir()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "drwxrwxrwx"
	permissionStr := "rwxrwxrwx"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_03(t *testing.T) {

	//  ModeSetuid()      os.ModeSetuid         "u" setuid
	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetuid())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeSetuid()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "urw-rw-rw-"
	permissionStr := "rw-rw-rw-"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_04(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "-rw-rw-rw-"
	permissionStr := "rw-rw-rw-"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_05(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "-r--r--r--"
	permissionStr := "r--r--r--"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_06(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "--w--w--w-"
	permissionStr := "-w--w--w-"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_07(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeDir()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "drw-rw-rw-"
	permissionStr := "rw-rw-rw-"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_08(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeDir()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "dr--r--r--"
	permissionStr := "r--r--r--"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_09(t *testing.T) {

	entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeDir()). "+
			"Error='%v' ", err.Error())
	}

	expectedPermissionTxt := "d-w--w--w-"
	permissionStr := "-w--w--w-"

	fPermCfg := FilePermissionConfig{}

	err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err != nil {
		t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
			"entrType='%s' permissionStr='%s' Error='%v' ",
			entryType.String(), permissionStr, err.Error())
	}

	actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedPermissionTxt != actualPermissionTxt {
		t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
			"Actual Permission Text Code='%v' ",
			expectedPermissionTxt, actualPermissionTxt)
	}

}

func TestFilePermissionConfig_SetFileModeByComponents_10(t *testing.T) {

	// Bad Entry Type Code
	entryType := OsFilePermissionCode(999)
	permissionStr := "rw-rw-rw-"

	fPermCfg := FilePermissionConfig{}

	err := fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err == nil {
		t.Error("Expected error return from bad entry type code 999. " +
			"However, NO ERROR WAS RETURNED! ")
	}
}

func TestFilePermissionConfig_SetFileModeByComponents_11(t *testing.T) {

	entryType := OsFilePermissionCode(OsFilePermCode.ModeNone())
	// Bad Permission String
	permissionStr := "rZ-rz-rz-"

	fPermCfg := FilePermissionConfig{}

	err := fPermCfg.SetFileModeByComponents(entryType, permissionStr)

	if err == nil {
		t.Error("Expected error return from bad permission string. " +
			"However, NO ERROR WAS RETURNED! ")
	}
}

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
