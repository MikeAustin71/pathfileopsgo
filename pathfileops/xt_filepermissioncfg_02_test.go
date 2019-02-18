package pathfileops

import "testing"

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
