package pathfileops

import "testing"

func TestFileAccessControl_New_01(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg := FileOpenConfig{}

	_, err = FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err == nil {
		t.Error("Expected an error returned from FileAccessControl{}.New(" +
			"fOpenCfg, fPermCfg) because fOpenCfg is uninitialized. " +
			"However, NO ERROR WAS RETURNED! \n")
	}

}

func TestFileAccessControl_New_02(t *testing.T) {

	fPermCfg := FilePermissionConfig{}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	_, err = FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err == nil {
		t.Error("Expected an error returned from FileAccessControl{}.New(" +
			"fOpenCfg, fPermCfg) because fPermCfg is uninitialized. " +
			"However, NO ERROR WAS RETURNED! \n")
	}

}

func TestFileAccessControl_SetFileOpenCodes_01(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fOpenCfg2, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	err = fAccess1.SetFileOpenCodes(fOpenCfg2)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFileOpenCodes(fOpenCfg2). "+
			"Error='%v' \n", err.Error())
	}

	actualFOpen2Cfg, err := fAccess1.GetFileOpenConfig()

	if err != nil {
		t.Errorf("Error returned by actualFOpen2Cfg = fAccess1.GetFileOpenConfig(). "+
			"Error='%v' \n", err.Error())
	}

	if !fOpenCfg2.Equal(&actualFOpen2Cfg) {
		t.Error("Expected fOpenCfg2 to equal actualFOpen2Cfg. " +
			"THEY ARE NOT EQUAL!")
	}

	if !actualFOpen2Cfg.Equal(&fOpenCfg2) {
		t.Error("Expected actualFOpen2Cfg to equal fOpenCfg2. " +
			"THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_SetFileOpenCodes_02(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1 := FileAccessControl{}

	err = fAccess1.SetFileOpenCodes(fOpenCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFileOpenCodes(fOpenCfg). "+
			"Error='%v' \n", err.Error())
	}

	err = fAccess1.SetFilePermissionCodes(fPermCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFilePermissionCodes(fPermCfg). "+
			"Error='%v' \n", err.Error())
	}

	if fAccess1.isInitialized == false {
		t.Error("Expected fAccess1.isInitialized == 'true'. Instead, it is 'false'. ")
	}

	actualFOpenCfg, err := fAccess1.GetFileOpenConfig()

	if err != nil {
		t.Errorf("Error returned by actualFOpenCfg = fAccess1.GetFileOpenConfig(). "+
			"Error='%v' \n", err.Error())
	}

	if !fOpenCfg.Equal(&actualFOpenCfg) {
		t.Error("Expected fOpenCfg to equal actualFOpenCfg. " +
			"THEY ARE NOT EQUAL!")
	}

	if !actualFOpenCfg.Equal(&fOpenCfg) {
		t.Error("Expected actualFOpenCfg to equal fOpenCfg. " +
			"THEY ARE NOT EQUAL!")
	}
}

func TestFileAccessControl_SetFileOpenCodes_03(t *testing.T) {

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1 := FileAccessControl{}

	err = fAccess1.SetFileOpenCodes(fOpenCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFileOpenCodes(fOpenCfg). "+
			"Error='%v' \n", err.Error())
	}

	if fAccess1.isInitialized == true {
		t.Error("Expected fAccess1.isInitialized to equal 'false'. " +
			"Instead, it is 'true'.")
	}
}

func TestFileAccessControl_SetFileOpenCodes_04(t *testing.T) {

	fOpenCfg := FileOpenConfig{}

	fAccess1 := FileAccessControl{}

	err := fAccess1.SetFileOpenCodes(fOpenCfg)

	if err == nil {
		t.Error("Expected error return from fAccess1.SetFileOpenCodes(fOpenCfg) " +
			"because fOpenCfg is uninitialized. However, NO ERROR WAS RETURNED!\n")
	}

}

func TestFileAccessControl_SetFilePermissionCodes_01(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	textCode2 := "drwxrwxrwx"

	fPermCfg2, err := FilePermissionConfig{}.New(textCode2)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	err = fAccess1.SetFilePermissionCodes(fPermCfg2)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFilePermissionCodes(fPermCfg2). "+
			"Error='%v' \n", err.Error())
	}

	actualFPermCfg2, err := fAccess1.GetFilePermissionConfig()

	if err != nil {
		t.Errorf("Error returned by actualFPermCfg2 = fAccess1.GetFilePermissionConfig(). "+
			"Error='%v' \n", err.Error())
	}

	if !fPermCfg2.Equal(&actualFPermCfg2) {
		t.Error("Expected fPermCfg2 to equal actualFPermCfg2. " +
			"THEY ARE NOT EQUAL!")
	}

	if !actualFPermCfg2.Equal(&fPermCfg2) {
		t.Error("Expected actualFOpen2Cfg to equal fOpenCfg2. " +
			"THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_SetFilePermissionCodes_02(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1 := FileAccessControl{}

	err = fAccess1.SetFileOpenCodes(fOpenCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFileOpenCodes(fOpenCfg). "+
			"Error='%v' \n", err.Error())
	}

	err = fAccess1.SetFilePermissionCodes(fPermCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFilePermissionCodes(fPermCfg). "+
			"Error='%v' \n", err.Error())
	}

	if fAccess1.isInitialized == false {
		t.Error("Expected fAccess1.isInitialized == 'true'. Instead, it is 'false'. ")
	}

	actualFPermCfg, err := fAccess1.GetFilePermissionConfig()

	if err != nil {
		t.Errorf("Error returned by actualFPermCfg = fAccess1.GetFilePermissionConfig(). "+
			"Error='%v' \n", err.Error())
	}

	if !fPermCfg.Equal(&actualFPermCfg) {
		t.Error("Expected fPermCfg to equal actualFPermCfg. " +
			"THEY ARE NOT EQUAL!")
	}

	if !actualFPermCfg.Equal(&fPermCfg) {
		t.Error("Expected actualFPermCfg to equal fPermCfg. " +
			"THEY ARE NOT EQUAL!")
	}
}

func TestFileAccessControl_SetFilePermissionCodes_03(t *testing.T) {

	textCode := "-rw-rw-rw-"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fAccess1 := FileAccessControl{}

	err = fAccess1.SetFilePermissionCodes(fPermCfg)

	if err != nil {
		t.Errorf("Error returned by fAccess1.SetFilePermissionCodes(fPermCfg). "+
			"Error='%v' \n", err.Error())
	}

	if fAccess1.isInitialized == true {
		t.Error("Expected fAccess1.isInitialized to equal 'false'. " +
			"Instead, it is 'true'.")
	}
}

func TestFileAccessControl_SetFilePermissionCodes_04(t *testing.T) {

	fPermCfg := FilePermissionConfig{}

	fAccess1 := FileAccessControl{}

	err := fAccess1.SetFilePermissionCodes(fPermCfg)

	if err == nil {
		t.Error("Expected error return from fAccess1.SetFilePermissionCodes(fPermCfg) " +
			"because fPermCfg is uninitialized. However, NO ERROR WAS RETURNED!\n")
	}

}
