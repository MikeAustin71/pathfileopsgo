package pathfileops

import "testing"

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
		t.Errorf("Error returned by fOpenCfg = FileOpenConfig{}.New(). " +
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
    t.Errorf("Error returned by fOpenCfg2 = FileOpenConfig{}.New(). " +
      "Error='%v' \n", err.Error())
  }

	fAccess2, err := FileAccessControl{}.New(fOpenCfg2,fPermCfg2)

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
