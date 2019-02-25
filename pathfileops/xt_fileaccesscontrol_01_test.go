package pathfileops

import (
	"os"
	"strings"
	"testing"
)

func TestFileAccessControl_CopyIn_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fAccess2 := FileAccessControl{}

	fAccess2.CopyIn(&fAccess1)

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_CopyIn_02(t *testing.T) {

	fAccess1 := FileAccessControl{}

	fAccess2 := FileAccessControl{}

	fAccess2.CopyIn(&fAccess1)

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_CopyOut_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fAccess2 := fAccess1.CopyOut()

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_CopyOut_02(t *testing.T) {

	fAccess1 := FileAccessControl{}

	fAccess2 := fAccess1.CopyOut()

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

}

func TestFileAccessControl_Empty_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fAccess2 := FileAccessControl{}

	fAccess2.CopyIn(&fAccess1)

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	fAccess2.Empty()

	if fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected after f2Access.Empty() fAccess2!=fAccess1. However, THEY ARE EQUAL!")
	}

	if fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected after f2Access.Empty() fAccess2!=fAccess1. However, THEY ARE EQUAL!")
	}

	fAccess1.Empty()

	if !fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected after both Empty(), fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}

	if !fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected after both Empty() fAccess2==fAccess1. However, THEY ARE NOT EQUAL!")
	}
}

func TestFileAccessControl_Empty_02(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	if !fAccess1.isInitialized {
		t.Error("Error: Expected fAccess1.isInitialized=='true'. However, it is 'false'!")
	}

	if !fAccess1.permissions.isInitialized {
		t.Error("Error: Expected fAccess1.permissions.isInitialized=='true'. However, it is 'false'!")
	}

	if !fAccess1.fileOpenCodes.isInitialized {
		t.Error("Error: Expected fAccess1.fileOpenCodes.isInitialized=='true'. However, it is 'false'!")
	}

	fAccess1.Empty()

	if fAccess1.isInitialized {
		t.Error("Error: Expected fAccess1.isInitialized=='false'. However, it is 'true'!")
	}

	if fAccess1.permissions.isInitialized {
		t.Error("Error: Expected fAccess1.permissions.isInitialized=='false'. However, it is 'true'!")
	}

	if fAccess1.fileOpenCodes.isInitialized {
		t.Error("Error: Expected fAccess1.fileOpenCodes.isInitialized=='false'. However, it is 'true'!")
	}

}

func TestFileAccessControl_Equal_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg = FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	textCode2 := "--w--w--w-"

	fPermCfg2, err := FilePermissionConfig{}.New(textCode2)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg2, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
		FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg2 = FileOpenConfig{}.New(). "+
			"Error='%v' \n", err.Error())
	}

	fAccess2, err := FileAccessControl{}.New(fOpenCfg2, fPermCfg2)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	if fAccess2.Equal(&fAccess1) {
		t.Error("Error: Expected fAccess2!=fAccess1. However, THEY ARE EQUAL!")
	}

	if fAccess1.Equal(&fAccess2) {
		t.Error("Error: Expected fAccess2!=fAccess1. However, THEY ARE EQUAL!")
	}

}

func TestFileAccessControl_GetCompositeFileOpenCode_01(t *testing.T) {

	textCode := "-rwxrwxrwx"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	accessFileOpenCode, err := fAccess1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fAccess1.GetCompositeFileOpenCode() "+
			"Error='%v' \n", err.Error())
	}

	originalFileOpenCode, err := fOpenCfg.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.GetCompositeFileOpenCode() "+
			"Error='%v' \n", err.Error())
	}

	if originalFileOpenCode != accessFileOpenCode {

		t.Errorf("Error: Expected originalFileOpenCode to Equal accessFileOpenCode. "+
			"THEY ARE NOT EQUAL! originalFileOpenCode='%s' accessFileOpenCode='%s' ",
			fOpenCfg.GetFileOpenNarrativeText(), fAccess1.fileOpenCodes.GetFileOpenNarrativeText())
	}

}

func TestFileAccessControl_GetCompositeFileOpenCode_02(t *testing.T) {

	textCode := "dr--r--r--"

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

	accessFileOpenCode, err := fAccess1.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fAccess1.GetCompositeFileOpenCode() "+
			"Error='%v' \n", err.Error())
	}

	originalFileOpenCode, err := fOpenCfg.GetCompositeFileOpenCode()

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.GetCompositeFileOpenCode() "+
			"Error='%v' \n", err.Error())
	}

	if originalFileOpenCode != accessFileOpenCode {

		t.Errorf("Error: Expected originalFileOpenCode to Equal accessFileOpenCode. "+
			"THEY ARE NOT EQUAL! originalFileOpenCode='%s' accessFileOpenCode='%s' ",
			fOpenCfg.GetFileOpenNarrativeText(), fAccess1.fileOpenCodes.GetFileOpenNarrativeText())
	}

}

func TestFileAccessControl_GetCompositePermissionCode01(t *testing.T) {

	textCode := "-rw-rw-rw-"
	expectedFMode := os.FileMode(0666)

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fMode, err := fAccess1.GetCompositePermissionMode()

	if err != nil {
		t.Errorf("Error returned by fAccess1.GetCompositePermissionMode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFMode != fMode {
		t.Error("Expected File Mode == 0666. Actual File Mode is different")
	}

}

func TestFileAccessControl_GetCompositePermissionCode02(t *testing.T) {

	textCode := "-rwxrwxrwx"
	expectedFMode := os.FileMode(0777)

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fMode, err := fAccess1.GetCompositePermissionMode()

	if err != nil {
		t.Errorf("Error returned by fAccess1.GetCompositePermissionMode(). "+
			"Error='%v' \n", err.Error())
	}

	if expectedFMode != fMode {
		t.Error("Expected File Mode == 0666. Actual File Mode is different")
	}

}

func TestFileAccessControl_GetCompositePermissionCode03(t *testing.T) {

	fAccess1 := FileAccessControl{}

	_, err := fAccess1.GetCompositePermissionMode()

	if err == nil {
		t.Error("Expected error return from fAccess1.GetCompositePermissionMode() " +
			"because it is uninitialized. However, NO ERROR WAS RETURNED!\n")
	}

}

func TestFileAccessControl_GetCompositePermissionModeText_01(t *testing.T) {

	textCode := "-rw-rw-rw-"
	expectedFMode := "0666"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fModeText := fAccess1.GetCompositePermissionModeText()

	if expectedFMode != fModeText {
		t.Errorf("Expected File Mode == '%v'. Actual File Mode Text == '%v'. ",
			expectedFMode, fModeText)
	}

}

func TestFileAccessControl_GetCompositePermissionModeText_02(t *testing.T) {

	textCode := "-rwxrwxrwx"
	expectedFMode := "0777"

	fPermCfg, err := FilePermissionConfig{}.New(textCode)

	if err != nil {
		t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
	}

	fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
		FOpenMode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
	}

	fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		t.Errorf("Error returned by FileAccessControl{}.New("+
			"fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
	}

	fModeText := fAccess1.GetCompositePermissionModeText()

	if expectedFMode != fModeText {
		t.Errorf("Expected File Mode == '%v'. Actual File Mode Text == '%v'. ",
			expectedFMode, fModeText)
	}

}

func TestFileAccessControl_GetCompositePermissionModeText_03(t *testing.T) {

	fAccess1 := FileAccessControl{}

	fModeText := fAccess1.GetCompositePermissionModeText()

	if strings.Index(strings.ToLower(fModeText), "invalid") == -1 {
		t.Error("Expected error message containing 'invalid. No such error " +
			"message was received.")
	}

}
