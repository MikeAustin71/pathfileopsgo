package pathfileops

import (
	"os"
	"strings"
	"testing"
)

func TestFilePermissionConfig_CopyIn_01(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := FilePermissionConfig{}

	fpCfg2.CopyIn(fpCfg)

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}

}

func TestFilePermissionConfig_CopyIn_02(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := FilePermissionConfig{}

	fpCfg2.CopyIn(fpCfg)

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}

}

func TestFilePermissionConfig_CopyOut_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_02(t *testing.T) {

	textCode := "drwxrwxrwx"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_03(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_04(t *testing.T) {

	textCode := "drw-rw-rw-"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_05(t *testing.T) {

	textCode := "-r--r--r--"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_06(t *testing.T) {

	textCode := "dr--r--r--"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_07(t *testing.T) {

	textCode := "--w--w--w-"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_08(t *testing.T) {

	textCode := "d-w--w--w-"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_09(t *testing.T) {

	textCode := "---xr----x"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
	}
}

func TestFilePermissionConfig_CopyOut_10(t *testing.T) {

	textCode := "d--xr----x"

	fpCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fpCfg2 := fpCfg.CopyOut()

	actualTextCode, err := fpCfg2.GetPermissionTextCode()

	if err != nil {
		t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
			"Error='%v'", err.Error())
	}

	if textCode != actualTextCode {
		t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
			textCode, actualTextCode)
	}

	if !fpCfg.Equal(fpCfg2) {
		t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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

	nineDigitText, err := fpCfg.GetPermissionBits()

	if err != nil {
		t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedIntPermissionBits != intFMode {
		t.Errorf("Error: Expected permission bits octal value = '%v' Instead, "+
			"permission bits octal value = '%v' \n"+
			"permission text bits = %v", expectedIntPermissionBits, intFMode, nineDigitText)
	}

	expectedPermissionTxt := strings.Replace(textCode, "d", "-", 1)

	if expectedPermissionTxt != nineDigitText.String() {
		t.Errorf("Error: Expected permission text string='%v'. Instead, "+
			"nine digit permission text string='%v' ", expectedPermissionTxt, nineDigitText.String())
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
