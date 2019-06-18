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
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg = FileOpenConfig{}.New(). "+
      "Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  textCode2 := "--w--w--w-"

  fPermCfg2, err := FilePermissionConfig{}.New(textCode2)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg2, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeCreate(), FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg2 = FileOpenConfig{}.New(). "+
      "Error='%v' \n", err.Error())
    return
  }

  fAccess2, err := FileAccessControl{}.New(fOpenCfg2, fPermCfg2)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  if fAccess2.Equal(&fAccess1) {
    t.Error("Error: Expected fAccess2!=fAccess1. However, THEY ARE EQUAL!")
  }

  if fAccess1.Equal(&fAccess2) {
    t.Error("Error: Expected fAccess2!=fAccess1. However, THEY ARE EQUAL!")
  }

}

func TestFileAccessControl_Equal_02(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  textCode2 := "-rw-rw-rw-"

  fPermCfg2, err := FilePermissionConfig{}.New(textCode2)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode2). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fAccess2, err := FileAccessControl{}.New(fOpenCfg, fPermCfg2)

  if err != nil {
    t.Errorf("Error returned by fAccess2 = FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg2). Error='%v' \n", err.Error())
    return
  }

  if fAccess1.Equal(&fAccess2) {
    t.Error("Expected that fAccess1 != fAccess2 because permissions are different. " +
      "Instead, they are EQUAL!")
  }

}

func TestFileAccessControl_GetCompositeFileOpenCode_01(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  accessFileOpenCode, err := fAccess1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetCompositeFileOpenCode() "+
      "Error='%v' \n", err.Error())
    return
  }

  originalFileOpenCode, err := fOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.GetCompositeFileOpenCode() "+
      "Error='%v' \n", err.Error())
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  accessFileOpenCode, err := fAccess1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetCompositeFileOpenCode() "+
      "Error='%v' \n", err.Error())
    return
  }

  originalFileOpenCode, err := fOpenCfg.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.GetCompositeFileOpenCode() "+
      "Error='%v' \n", err.Error())
    return
  }

  if originalFileOpenCode != accessFileOpenCode {

    t.Errorf("Error: Expected originalFileOpenCode to Equal accessFileOpenCode. "+
      "THEY ARE NOT EQUAL! originalFileOpenCode='%s' accessFileOpenCode='%s' ",
      fOpenCfg.GetFileOpenNarrativeText(), fAccess1.fileOpenCodes.GetFileOpenNarrativeText())
  }

}

func TestFileAccessControl_GetCompositeFileOpenCode_03(t *testing.T) {

  fAccess1 := FileAccessControl{}

  _, err := fAccess1.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error from fAccess1.GetCompositeFileOpenCode() " +
      "because fAccess1 is uninitialized. However, NO ERROR WAS RETURN! \n")
  }

}

func TestFileAccessControl_GetCompositeFileOpenCode_04(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.fileOpenCodes.fileOpenType = FileOpenType(-99)

  _, err = fAccess1.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error return from fAccess1.GetCompositeFileOpenCode() " +
      "because File Open Type Code is invalid. However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetCompositePermissionCode01(t *testing.T) {

  textCode := "-rw-rw-rw-"
  expectedFMode := os.FileMode(0666)

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fMode, err := fAccess1.GetCompositePermissionMode()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetCompositePermissionMode(). "+
      "Error='%v' \n", err.Error())
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fMode, err := fAccess1.GetCompositePermissionMode()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetCompositePermissionMode(). "+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFMode != fMode {
    t.Error("Expected File Mode == 0666. Actual File Mode is different")
  }

}

func TestFileAccessControl_GetCompositePermissionCode03(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions = FilePermissionConfig{}

  _, err = fAccess1.GetCompositePermissionMode()

  if err == nil {
    t.Error("Expected an error return from fAccess1.GetCompositePermissionMode() " +
      "because fAcess1.permissions was uninitialized. However, NO ERROR WAS RETURNED! \n")
    return
  }

}

func TestFileAccessControl_GetCompositePermissionCode04(t *testing.T) {

  fAccess1 := FileAccessControl{}

  _, err := fAccess1.GetCompositePermissionMode()

  if err == nil {
    t.Error("Expected error return from fAccess1.GetCompositePermissionMode() " +
      "because it is uninitialized. However, NO ERROR WAS RETURNED!\n")
    return
  }

}

func TestFileAccessControl_GetCompositePermissionCode05(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions.isInitialized = false

  _, err = fAccess1.GetCompositePermissionMode()

  if err == nil {
    t.Error("Expected an error returned from fAccess1.GetCompositePermissionMode() " +
      "because fAccess1.permissions.isInitialized = false " +
      "However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetCompositePermissionModeText_01(t *testing.T) {

  textCode := "-rw-rw-rw-"
  expectedFMode := "0666"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
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
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fModeText := fAccess1.GetCompositePermissionModeText()

  if expectedFMode != fModeText {
    t.Errorf("Expected File Mode == '%v'. Actual File Mode Text == '%v'. ",
      expectedFMode, fModeText)
  }

}

func TestFileAccessControl_GetCompositePermissionModeText_03(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions = FilePermissionConfig{}

  fModeText := fAccess1.GetCompositePermissionModeText()

  if strings.Index(strings.ToLower(fModeText), "invalid") == -1 {
    t.Error("Expected fModeText of contain 'invalid' because fAccess1.permissions" +
      " is empty. However, no error was detected!")
  }

}

func TestFileAccessControl_GetCompositePermissionModeText_04(t *testing.T) {

  fAccess1 := FileAccessControl{}

  fModeText := fAccess1.GetCompositePermissionModeText()

  if strings.Index(strings.ToLower(fModeText), "invalid") == -1 {
    t.Error("Expected error message containing 'invalid. No such error " +
      "message was received.")
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_01(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fOpenCode, fPermCode, err := fAccess1.GetFileOpenAndPermissionCodes()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetFileOpenAndPermissionCodes() "+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpenCode != int(FOpenType.TypeReadWrite()) {
    t.Error("Expected fOpenCode to contain FOpenType.TypeReadWrite(). It did Not!")
  }

  if textCode != fPermCode.String() {
    t.Errorf("Expected fPermCode.String() == '%v'. Actual File Mode Text == '%v'. ",
      textCode, fPermCode.String())
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_02(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fOpenCode, fPermCode, err := fAccess1.GetFileOpenAndPermissionCodes()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetFileOpenAndPermissionCodes() "+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpenCode != int(FOpenType.TypeReadOnly()) {
    t.Error("Expected fOpenCode to contain FOpenType.TypeReadWrite(). It did Not!")
  }

  if textCode != fPermCode.String() {
    t.Errorf("Expected fPermCode.String() == '%v'. Actual File Mode Text == '%v'. ",
      textCode, fPermCode.String())
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_03(t *testing.T) {

  textCode := "drwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fOpenCode, fPermCode, err := fAccess1.GetFileOpenAndPermissionCodes()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetFileOpenAndPermissionCodes() "+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpenCode != int(FOpenType.TypeWriteOnly()) {
    t.Error("Expected fOpenCode to contain FOpenType.TypeReadWrite(). It did Not!")
  }

  if textCode != fPermCode.String() {
    t.Errorf("Expected fPermCode.String() == '%v'. Actual File Mode Text == '%v'. ",
      textCode, fPermCode.String())
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_04(t *testing.T) {

  textCode := "drwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.fileOpenCodes = FileOpenConfig{}

  _, _, err = fAccess1.GetFileOpenAndPermissionCodes()

  if err == nil {
    t.Errorf("Expected an error from fAccess1.GetFileOpenAndPermissionCodes() " +
      "because fAcess1.fileOpenCodes are uninitialized. However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_05(t *testing.T) {

  textCode := "drwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions = FilePermissionConfig{}

  _, _, err = fAccess1.GetFileOpenAndPermissionCodes()

  if err == nil {
    t.Errorf("Expected an error from fAccess1.GetFileOpenAndPermissionCodes() " +
      "because fAcess1.permissions are uninitialized. However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_06(t *testing.T) {

  fAccess1 := FileAccessControl{}

  _, _, err := fAccess1.GetFileOpenAndPermissionCodes()

  if err == nil {
    t.Errorf("Expected an error from fAccess1.GetFileOpenAndPermissionCodes() " +
      "because fAcess1 was uninitialized. However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetFileOpenAndPermissionCodes_07(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.fileOpenCodes.fileOpenType = FileOpenType(-999)

  _, _, err = fAccess1.GetFileOpenAndPermissionCodes()

  if err == nil {
    t.Error("Expected an error return from fAccess1.GetFileOpenAndPermissionCodes() " +
      "because fAccess1.fileOpenCodes.fileOpenType is invalid. " +
      "However, NO ERROR WAS RETURNED! \n")
  }

}
func TestFileAccessControl_GetFileOpenAndPermissionCodes_08(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions.fileMode = os.FileMode(999999)

  _, _, err = fAccess1.GetFileOpenAndPermissionCodes()

  if err == nil {
    t.Error("Expected an error return from fAccess1.GetFileOpenAndPermissionCodes() " +
      "because fAccess1.fileOpenCodes.fileOpenType is invalid. " +
      "However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileAccessControl_GetFileOpenConfig_01(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1FileOpenCfg, err := fAccess1.GetFileOpenConfig()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetFileOpenConfig() "+
      "Error='%v' \n", err.Error())
    return
  }

  if !fOpenCfg.Equal(&fAccess1FileOpenCfg) {
    t.Error("Expected original fOpenCfg to equal returned file open configuration. " +
      "The two are NOT equal!")
  }

  if !fAccess1FileOpenCfg.Equal(&fOpenCfg) {
    t.Error("Expected returned fAccess1FileOpenCfg to equal the original file open " +
      "configuration. The two are NOT equal!")
  }

}

func TestFileAccessControl_GetFileOpenConfig_02(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions = FilePermissionConfig{}

  _, err = fAccess1.GetFileOpenConfig()

  if err == nil {
    t.Error("Expected error from fAccess1.GetFileOpenConfig() " +
      "because fAccess1.permissions is uninitialized. " +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileAccessControl_GetFileOpenConfig_03(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.fileOpenCodes = FileOpenConfig{}

  _, err = fAccess1.GetFileOpenConfig()

  if err == nil {
    t.Error("Expected error from fAccess1.GetFileOpenConfig() " +
      "because fAccess1.fileOpenCodes is uninitialized. " +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileAccessControl_GetFileOpenConfig_04(t *testing.T) {

  fAccess1 := FileAccessControl{}

  _, err := fAccess1.GetFileOpenConfig()

  if err == nil {
    t.Error("Expected error from fAccess1.GetFileOpenConfig() " +
      "because fAccess1 is uninitialized. " +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileAccessControl_GetFilePermissionConfig_01(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fActualPermCfg, err := fAccess1.GetFilePermissionConfig()

  if err != nil {
    t.Errorf("Error returned by fAccess1.GetFilePermissionConfig() "+
      "Error='%v' \n", err.Error())
    return
  }

  if !fPermCfg.Equal(&fActualPermCfg) {
    t.Error("Expected the original file permission config to equal the " +
      "returned file permission config. They ARE NOT EQUAL!")
  }

  if !fActualPermCfg.Equal(&fPermCfg) {
    t.Error("Expected the returned file permission config to equal the " +
      "original file permission config. They ARE NOT EQUAL!")
  }

}

func TestFileAccessControl_GetFilePermissionConfig_02(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.fileOpenCodes = FileOpenConfig{}

  _, err = fAccess1.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected error return from fAccess1.GetFilePermissionConfig() " +
      "because fAccess1.fileOpenCodes are uninitialized. \n")
  }

}

func TestFileAccessControl_GetFilePermissionConfig_03(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
    return
  }

  fOpenCfg, err := FileOpenConfig{}.New(FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpenCfg.New(). Error='%v' \n", err.Error())
    return
  }

  fAccess1, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

  if err != nil {
    t.Errorf("Error returned by FileAccessControl{}.New("+
      "fOpenCfg, fPermCfg). Error='%v' \n", err.Error())
    return
  }

  fAccess1.permissions = FilePermissionConfig{}

  _, err = fAccess1.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected error return from fAccess1.GetFilePermissionConfig() " +
      "because fAccess1.permissions are uninitialized. \n")
  }

}

func TestFileAccessControl_GetFilePermissionConfig_04(t *testing.T) {

  fAccess1 := FileAccessControl{}

  _, err := fAccess1.GetFilePermissionConfig()

  if err == nil {
    t.Error("Expected error return from fAccess1.GetFilePermissionConfig() " +
      "because fAccess1 is uninitialized. \n")
  }

}
