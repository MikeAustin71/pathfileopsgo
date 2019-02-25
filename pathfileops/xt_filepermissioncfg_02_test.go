package pathfileops

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestFilePermissionConfig_GetCompositePermissionMode_01(t *testing.T) {

	expectedFileMode := os.FileMode(0666)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(expectedFileMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode"+
			"(os.FileMode(0666)). Error='%v' ", err.Error())
	}

	fMode, err := fPerm.GetCompositePermissionMode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetCompositePermissionMode() "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != fMode {
		t.Errorf("Error: Expected File Mode octal value='%v'. Instead, "+
			"File Mode octal value ='%v' ",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(fMode), 8))
	}
}

func TestFilePermissionConfig_GetCompositePermissionMode_02(t *testing.T) {

	expectedFileMode := os.FileMode(0666)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(expectedFileMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode"+
			"(os.FileMode(0666)). Error='%v' ", err.Error())
	}

	fPerm.isInitialized = false

	_, err = fPerm.GetCompositePermissionMode()

	if err == nil {
		t.Errorf("Expected error to be returned by fPerm.GetCompositePermissionMode() " +
			"because isInitialized flag is 'false'. However, NO ERROR WAS RETURNED!!")
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_01(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg := FilePermissionConfig{}

	err := fpCfg.SetFileModeByTextCode(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	actualEntryType, err := fpCfg.GetEntryTypeComponent()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetEntryTypeComponent(). "+
			"Error='%v' ", err.Error())
	}

	expectedEntryType := OsFilePermissionCode(os.ModeDir)

	if expectedEntryType != actualEntryType {
		t.Errorf("Error: Expected Entry Type Component = %s. Instead, "+
			"Entry Type Component= %s",
			expectedEntryType.String(),
			actualEntryType.String())
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_02(t *testing.T) {

	textCode := "rwxrwxrwx"

	fpCfg := FilePermissionConfig{}

	osPerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetgid())

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	err = fpCfg.SetFileModeByComponents(osPerm, textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	actualEntryType, err := fpCfg.GetEntryTypeComponent()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetEntryTypeComponent(). "+
			"Error='%v' ", err.Error())
	}

	expectedEntryType := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	if actualEntryType != expectedEntryType {
		t.Errorf("Error: Expected Entry Type Component (ModeSetgid) = %s. Instead, "+
			"Entry Type Component= %s",
			expectedEntryType.String(),
			actualEntryType.String())
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_03(t *testing.T) {

	textCode := "rw-rw-rw-"

	fpCfg := FilePermissionConfig{}

	osPerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetuid())

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	err = fpCfg.SetFileModeByComponents(osPerm, textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	actualEntryType, err := fpCfg.GetEntryTypeComponent()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetEntryTypeComponent(). "+
			"Error='%v' ", err.Error())
	}

	expectedEntryType := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	if actualEntryType != expectedEntryType {
		t.Errorf("Error: Expected Entry Type Component (ModeSetgid) = %s. Instead, "+
			"Entry Type Component= %s",
			expectedEntryType.String(),
			actualEntryType.String())
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_04(t *testing.T) {

	textCode := "-w--w--w-"

	fpCfg := FilePermissionConfig{}

	osPerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	err = fpCfg.SetFileModeByComponents(osPerm, textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	actualEntryType, err := fpCfg.GetEntryTypeComponent()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetEntryTypeComponent(). "+
			"Error='%v' ", err.Error())
	}

	expectedEntryType := OsFilePermissionCode(OsFilePermCode.ModeNone())

	if actualEntryType != expectedEntryType {
		t.Errorf("Error: Expected Entry Type Component (ModeSetgid) = %s. Instead, "+
			"Entry Type Component= %s",
			expectedEntryType.String(),
			actualEntryType.String())
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_05(t *testing.T) {

	textCode := "r--r--r--"

	fpCfg := FilePermissionConfig{}

	osPerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNamedPipe())

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	err = fpCfg.SetFileModeByComponents(osPerm, textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
			"Error='%v' ", err.Error())
	}

	actualEntryType, err := fpCfg.GetEntryTypeComponent()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetEntryTypeComponent(). "+
			"Error='%v' ", err.Error())
	}

	expectedEntryType := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	if actualEntryType != expectedEntryType {
		t.Errorf("Error: Expected Entry Type Component (ModeSetgid) = %s. Instead, "+
			"Entry Type Component= %s",
			expectedEntryType.String(),
			actualEntryType.String())
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_06(t *testing.T) {

	fpCfg := FilePermissionConfig{}

	_, err := fpCfg.GetEntryTypeComponent()

	if err == nil {
		t.Error("Expected an error from fpCfg.GetEntryTypeComponent() because " +
			"fpCfg was not initialized. NO ERROR WAS RETURNED!")
	}

}

func TestFilePermissionConfig_GetEntryTypeComponent_07(t *testing.T) {

	fpCfg := FilePermissionConfig{}

	fpCfg.isInitialized = true

	fpCfg.fileMode = os.FileMode(01000)

	_, err := fpCfg.GetEntryTypeComponent()

	if err == nil {
		t.Error("Expected an error from fpCfg.GetEntryTypeComponent() because " +
			"fpCfg.fileMode has an invalid code. NO ERROR WAS RETURNED!")
	}

}

func TestFilePermissionConfig_GetIsDir_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	isDir, err := fpCfg.GetIsDir()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetIsDir(). Error='%v' ", err.Error())
	}

	if isDir {
		t.Error("Error: Expected fpCfg.GetIsDir()=='false'. Instead, it returned 'true'.")
	}

}

func TestFilePermissionConfig_GetIsDir_02(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	isDir, err := fpCfg.GetIsDir()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetIsDir(). Error='%v' ", err.Error())
	}

	if !isDir {
		t.Error("Error: Expected fpCfg.GetIsDir()=='false'. Instead, it returned 'true'.")
	}

}

func TestFilePermissionConfig_GetIsDir_03(t *testing.T) {

	fpCfg := FilePermissionConfig{}

	_, err := fpCfg.GetIsDir()

	if err == nil {
		t.Error("Expected fpCfg.GetIsDir() to return an error because " +
			"fpCfg was not initialized. NO ERROR WAS RETURNED!")
	}

}

func TestFilePermissionConfig_GetIsRegular_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	isRegular, err := fpCfg.GetIsRegular()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetIsRegular(). Error='%v' ", err.Error())
	}

	if !isRegular {
		t.Error("Error: Expected fpCfg.GetIsRegular()=='true'. Instead, it returned 'false'.")
	}

}

func TestFilePermissionConfig_GetIsRegular_02(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	isRegular, err := fpCfg.GetIsRegular()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetIsRegular(). Error='%v' ", err.Error())
	}

	if isRegular {
		t.Error("Error: Expected fpCfg.GetIsRegular()=='false'. Instead, it returned 'true'.")
	}

}

func TestFilePermissionConfig_GetIsRegular_03(t *testing.T) {

	fpCfg := FilePermissionConfig{}

	_, err := fpCfg.GetIsRegular()

	if err == nil {
		t.Error("Expected an error to be returned by fpCfg.GetIsRegular() because " +
			"fpCfg was NOT initialized. NO ERROR WAS RETURNED.")
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
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

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionBits(). "+
			"Error='%v' ", err.Error())
	}

	fh := FileHelper{}
	intFMode := fh.ConvertDecimalToOctal(int(fMode))

	tenDigitText, err := fpCfg.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, tenDigitText)
	}

	if textCode != tenDigitText {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", textCode, tenDigitText)
	}
}

func TestFilePermissionConfig_GetPermissionBits_15(t *testing.T) {

	fpCfg := FilePermissionConfig{}

	_, err := fpCfg.GetPermissionBits()

	if err == nil {
		t.Error("Expected an error return from fpCfg.GetPermissionBits() " +
			"because fpCfg was not initialized. NO ERROR WAS RETURNED!")
	}

}

func TestFilePermissionConfig_GetPermissionComponents_01(t *testing.T) {
	expectedTextCode := "-rw-rw-rw-"

	expectedPermissionFileMode := os.FileMode(0666)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_02(t *testing.T) {
	expectedTextCode := "drw-rw-rw-"

	expectedPermissionFileMode := os.FileMode(0666)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents() "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_03(t *testing.T) {
	expectedTextCode := "-r--r--r--"

	expectedPermissionFileMode := os.FileMode(0444)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_04(t *testing.T) {
	expectedTextCode := "dr--r--r--"

	expectedPermissionFileMode := os.FileMode(0444)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_05(t *testing.T) {
	expectedTextCode := "--w--w--w-"

	expectedPermissionFileMode := os.FileMode(0222)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_06(t *testing.T) {
	expectedTextCode := "d-w--w--w-"

	expectedPermissionFileMode := os.FileMode(0222)

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	expectedOsFilePerm, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	actualOsFilePerm, actualPermFMode, err := fPerm.GetPermissionComponents()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetPermissionComponents(). "+
			"Error='%v' ", err.Error())
	}

	if expectedOsFilePerm != actualOsFilePerm {
		t.Errorf("Error: Expected OsFilePerm='%v'. Instead, OsFilePerm='%v' ",
			expectedOsFilePerm.String(), actualOsFilePerm.String())
	}

	if expectedPermissionFileMode != actualPermFMode {
		t.Errorf("Error: Expected Permission File Mode Octal Value='%s'. Instead, "+
			"Permission File Mode Octal Value='%s'",
			strconv.FormatInt(int64(expectedPermissionFileMode), 8),
			strconv.FormatInt(int64(actualPermFMode), 8))
	}

}

func TestFilePermissionConfig_GetPermissionComponents_07(t *testing.T) {

	expectedTextCode := "--w--w--w-"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	_, err = OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode"+
			"(OsFilePermCode.ModeNone()).  Error='%v' ", err.Error())
	}

	fPerm.fileMode = os.FileMode(01000)

	_, _, err = fPerm.GetPermissionComponents()

	if err == nil {
		t.Error("Expected an error return by fPerm.GetPermissionComponents() " +
			"because fPerm.fileMode has an invalid EntryType")
	}

}

func TestFilePermissionConfig_GetPermissionComponents_08(t *testing.T) {

	expectedTextCode := "--w--w--w-"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	fPerm.isInitialized = false

	_, _, err = fPerm.GetPermissionComponents()

	if err == nil {
		t.Error("Expected an error return from fPerm.GetPermissionComponents() " +
			"because fPerm was not initialized. NO ERROR WAS RETURNED!!")
	}

}

func TestFilePermissionConfig_GetPermissionComponents_09(t *testing.T) {

	fPerm := FilePermissionConfig{}

	_, _, err := fPerm.GetPermissionComponents()

	if err == nil {
		t.Error("Expected an error return from fPerm.GetPermissionComponents() " +
			"because fPerm was not initialized. NO ERROR WAS RETURNED!!")
	}

}

func TestFilePermissionConfig_GetPermissionFileModeValueText_01(t *testing.T) {

	expectedTextCode := "-rwxrwxrwx"

	expectedValueCodeText := "777"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	actualFileModeValueText := fPerm.GetPermissionFileModeValueText()

	if strings.Index(actualFileModeValueText, expectedValueCodeText) == -1 {
		t.Errorf("Expected File Mode Value Text='%v'. Instead, File Mode Value Text='%v' ",
			expectedValueCodeText, actualFileModeValueText)
	}

}

func TestFilePermissionConfig_GetPermissionFileModeValueText_02(t *testing.T) {

	expectedTextCode := "-rw-rw-rw-"

	expectedValueCodeText := "666"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	actualFileModeValueText := fPerm.GetPermissionFileModeValueText()

	if strings.Index(actualFileModeValueText, expectedValueCodeText) == -1 {
		t.Errorf("Expected File Mode Value Text='%v'. Instead, File Mode Value Text='%v' ",
			expectedValueCodeText, actualFileModeValueText)
	}

}

func TestFilePermissionConfig_GetPermissionFileModeValueText_03(t *testing.T) {

	fPerm := FilePermissionConfig{}

	actualFileModeValueText := fPerm.GetPermissionFileModeValueText()

	if strings.Index(strings.ToLower(actualFileModeValueText), "invalid") == -1 {
		t.Error("Expected error stirng containing 'invalid'. No such error string was received.")
	}

}

func TestFilePermissionConfig_GetPermissionNarrativeText_01(t *testing.T) {

	expectedTextCode := "-rwxrwxrwx"
	expectedFileModeVal := "0777"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	narTxt := fPerm.GetPermissionNarrativeText()

	if strings.Index(narTxt, "ModeFile") == -1 {
		t.Error("Error: Expected narrative text to include 'ModeFile'. It did NOT!")
	}

	if strings.Index(narTxt, expectedTextCode) == -1 {
		t.Errorf("Error: Expected narrative text to include '%v'. It did NOT!", expectedTextCode)
	}

	if strings.Index(narTxt, expectedFileModeVal) == -1 {
		t.Errorf("Error: Expected narrative text to include '%v'. It did NOT!", expectedFileModeVal)
	}

}

func TestFilePermissionConfig_GetPermissionNarrativeText_02(t *testing.T) {

	expectedTextCode := "drwxrwxrwx"
	expectedFileModeVal := "020000000777"

	fPerm, err := FilePermissionConfig{}.New(expectedTextCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(expectedTextCode). "+
			"expectedTextCode='%v' Error='%v' ", expectedTextCode, err.Error())
	}

	narTxt := fPerm.GetPermissionNarrativeText()

	if strings.Index(narTxt, "ModeDir") == -1 {
		t.Error("Error: Expected narrative text to include 'ModeFile'. It did NOT!")
	}

	if strings.Index(narTxt, expectedTextCode) == -1 {
		t.Errorf("Error: Expected narrative text to include '%v'. It did NOT!", expectedTextCode)
	}

	if strings.Index(narTxt, expectedFileModeVal) == -1 {
		t.Errorf("Error: Expected narrative text to include '%v'. It did NOT!", expectedFileModeVal)
	}

}

func TestFilePermissionConfig_GetPermissionNarrativeText_03(t *testing.T) {

	fPerm := FilePermissionConfig{}

	x := fPerm.GetPermissionNarrativeText()

	if strings.Index(strings.ToLower(x), "invalid") == -1 {
		t.Error("Expected error string containing 'invalid'. No such error string was received.")
	}

}

func TestFilePermissionConfig_GetPermissionTextCode_01(t *testing.T) {

	expectedTextCode := "drwxrwxrwx"
	fh := FileHelper{}

	// drwxrwxrwx   20000000777

	intFMode := fh.ConvertOctalToDecimal(20000000777)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	textCode, err := fPerm.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedTextCode != textCode {
		t.Errorf("Error: Expected permission text code='%v'. Instead, "+
			"permission text code='%v'",
			expectedTextCode, textCode)
	}
}

func TestFilePermissionConfig_GetPermissionTextCode_02(t *testing.T) {

	// -r--r--r--   0444
	expectedTextCode := "-r--r--r--"
	fh := FileHelper{}

	intFMode := fh.ConvertOctalToDecimal(444)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	textCode, err := fPerm.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedTextCode != textCode {
		t.Errorf("Error: Expected permission text code='%v'. Instead, "+
			"permission text code='%v'",
			expectedTextCode, textCode)
	}
}

func TestFilePermissionConfig_GetPermissionTextCode_03(t *testing.T) {

	// --w--w--w-   0222
	expectedTextCode := "--w--w--w-"
	fh := FileHelper{}

	intFMode := fh.ConvertOctalToDecimal(222)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	textCode, err := fPerm.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedTextCode != textCode {
		t.Errorf("Error: Expected permission text code='%v'. Instead, "+
			"permission text code='%v'",
			expectedTextCode, textCode)
	}
}

func TestFilePermissionConfig_GetPermissionTextCode_04(t *testing.T) {

	// -rw-rw-rw-   0666
	expectedTextCode := "-rw-rw-rw-"
	fh := FileHelper{}

	intFMode := fh.ConvertOctalToDecimal(666)

	osFMode := os.FileMode(intFMode)

	fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	textCode, err := fPerm.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedTextCode != textCode {
		t.Errorf("Error: Expected permission text code='%v'. Instead, "+
			"permission text code='%v'",
			expectedTextCode, textCode)
	}
}
