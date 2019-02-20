package pathfileops

import (
	"os"
	"strconv"
	"testing"
)

func TestOsFilePermissionCode_Equal_01(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeNone()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermissionCode(0)."+
			"GetNewFromFileMode(OsFilePermCode.ModeNone()) Error='%v' ",
			err.Error())
	}

	if !osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to Equal osPerm2. THEY ARE NOT EQUAL!")
	}

	if !osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to Equal osPerm1. THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_02(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeDir()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeDir()) Error='%v' ",
			err.Error())
	}

	if !osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to Equal osPerm2. THEY ARE NOT EQUAL!")
	}

	if !osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to Equal osPerm1. THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_03(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNamedPipe())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeNamedPipe()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNamedPipe())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeNamedPipe()) Error='%v' ",
			err.Error())
	}

	if !osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to Equal osPerm2. THEY ARE NOT EQUAL!")
	}

	if !osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to Equal osPerm1. THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_04(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeIrregular())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeIrregular()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeIrregular())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeIrregular()) Error='%v' ",
			err.Error())
	}

	if !osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to Equal osPerm2. THEY ARE NOT EQUAL!")
	}

	if !osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to Equal osPerm1. THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_05(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeNone()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeIrregular())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeIrregular()) Error='%v' ",
			err.Error())
	}

	if osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to NOT Equal osPerm2. However, THEY ARE EQUAL!")
	}

	if osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to NOT Equal osPerm1. However, THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_06(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeDir()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeAppend())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeAppend()) Error='%v' ",
			err.Error())
	}

	if osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to NOT Equal osPerm2. However, THEY ARE EQUAL!")
	}

	if osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to NOT Equal osPerm1. However, THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_07(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNamedPipe())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeNamedPipe()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeIrregular())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeIrregular()) Error='%v' ",
			err.Error())
	}

	if osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to NOT Equal osPerm2. However, THEY ARE EQUAL!")
	}

	if osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to NOT Equal osPerm1. However, THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_Equal_08(t *testing.T) {

	osPerm1, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetuid())

	if err != nil {
		t.Errorf("Error returned by osPerm1=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeSetuid()) Error='%v' ",
			err.Error())
	}

	osPerm2, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetgid())

	if err != nil {
		t.Errorf("Error returned by osPerm2=OsFilePermCode."+
			"GetNewFromFileMode(OsFilePermCode.ModeSetgid()) Error='%v' ",
			err.Error())
	}

	if osPerm1.Equal(osPerm2) {
		t.Error("Expected osPerm1 to NOT Equal osPerm2. However, THEY ARE EQUAL!")
	}

	if osPerm2.Equal(osPerm1) {
		t.Error("Expected osPerm2 to NOT Equal osPerm1. However, THEY ARE NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_IsValid_01(t *testing.T) {

	fpc := OsFilePermCode.ModeDir()

	if os.ModeDir != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDir. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_02(t *testing.T) {

	fpc := OsFilePermCode.ModeNone()

	fmBase := os.FileMode(0)

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeNone. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_03(t *testing.T) {

	fpc := OsFilePermCode.ModeAppend()

	fmBase := os.ModeAppend

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeAppend. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_04(t *testing.T) {

	fpc := OsFilePermCode.ModeExclusive()

	fmBase := os.ModeExclusive

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeExclusive. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_05(t *testing.T) {

	fpc := OsFilePermCode.ModeTemporary()

	fmBase := os.ModeTemporary

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeTemporary. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_06(t *testing.T) {

	fpc := OsFilePermCode.ModeSymlink()

	fmBase := os.ModeSymlink

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSymlink. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_07(t *testing.T) {

	fpc := OsFilePermCode.ModeDevice()

	fmBase := os.ModeDevice

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDevice. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_08(t *testing.T) {

	fpc := OsFilePermCode.ModeNamedPipe()

	fmBase := os.ModeNamedPipe

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeNamedPipe. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_09(t *testing.T) {

	fpc := OsFilePermCode.ModeSocket()

	fmBase := os.ModeSocket

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSocket. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_10(t *testing.T) {

	fpc := OsFilePermCode.ModeSetuid()

	fmBase := os.ModeSetuid

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSetuid. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_11(t *testing.T) {

	fpc := OsFilePermCode.ModeSetgid()

	fmBase := os.ModeSetgid

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSetgid. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_12(t *testing.T) {

	fpc := OsFilePermCode.ModeCharDevice()

	fmBase := os.ModeCharDevice

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeCharDevice. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_13(t *testing.T) {

	fpc := OsFilePermCode.ModeSticky()

	fmBase := os.ModeSticky

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeSticky. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_14(t *testing.T) {

	fpc := OsFilePermCode.ModeIrregular()

	fmBase := os.ModeIrregular

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeIrregular. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err != nil {
		t.Errorf("Error received from osPerm.IsValid(). "+
			"Error='%v'", err.Error())
	}

}

func TestOsFilePermissionCode_IsValid_15(t *testing.T) {

	fpc := os.FileMode(999)

	osPerm := OsFilePermissionCode(fpc)

	err := osPerm.IsValid()

	if err == nil {
		t.Error("Expected an Error returned from Invalid FileMode code 999. " +
			"NO ERROR WAS RETURNED!!!!")
	}

}

func TestOsFilePermissionCode_GetFileModeLetterCode_01(t *testing.T) {
	expectedLetter := "-"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeNone())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_02(t *testing.T) {
	expectedLetter := "d"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeDir())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_03(t *testing.T) {
	//     ModeAppend            a:      append-only
	expectedLetter := "a"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_04(t *testing.T) {
	//     ModeExclusive         l:      exclusive use

	expectedLetter := "l"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_05(t *testing.T) {

	//     ModeTemporary         T:      temporary file; Plan 9 only

	expectedLetter := "T"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeTemporary())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_06(t *testing.T) {

	//     ModeSymlink           L:      symbolic link

	expectedLetter := "L"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSymlink())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_07(t *testing.T) {

	//     ModeDevice            D:      device file

	expectedLetter := "D"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeDevice())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_08(t *testing.T) {

	//     ModeNamedPipe         p:      named pipe (FIFO)

	expectedLetter := "p"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_09(t *testing.T) {

	//     ModeSocket            S:      Unix domain socket

	expectedLetter := "S"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSocket())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_10(t *testing.T) {

	//     ModeSetuid            u:      setuid

	expectedLetter := "u"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_11(t *testing.T) {

	//     ModeSetgid            g:      setgid

	expectedLetter := "g"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_12(t *testing.T) {

	//     ModeCharDevice        c:      Unix character device, when ModeDevice is set

	expectedLetter := "c"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeCharDevice())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_13(t *testing.T) {

	//     ModeSticky            t:      sticky

	expectedLetter := "t"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_14(t *testing.T) {

	//     ModeIrregular         ?:      non-regular file; nothing else is known about this file

	expectedLetter := "?"

	fPerm := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	actualLetter, err := fPerm.GetFileModeLetterCode()

	if err != nil {
		t.Errorf("Error returned by fPerm.GetFileModeLetterCode(). "+
			"Error='%v' ", err.Error())
	}

	if expectedLetter != actualLetter {
		t.Errorf("Expected Letter Code: '%v'. Instead, Actual Letter Code='%v'",
			expectedLetter, actualLetter)
	}
}

func TestOsFilePermissionCode_GetFileModeLetterCode_15(t *testing.T) {

	fPerm := OsFilePermissionCode(12577)

	_, err := fPerm.GetFileModeLetterCode()

	if err == nil {
		t.Errorf("Expected an error return from fPerm.GetFileModeLetterCode() " +
			"because of an invalid OsFilePermissionCode. However, NO ERROR WAS RETURNED!")
	}
}

func TestOsFilePermissionCode_GetNewFromFileMode_01(t *testing.T) {

	expectedFileMode := os.FileMode(0)

	fileMode := OsFilePermCode.ModeNone()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeNone())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_02(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeDir()

	fileMode := OsFilePermCode.ModeDir()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeDir())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_03(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeAppend()

	fileMode := OsFilePermCode.ModeAppend()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_04(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeExclusive()

	fileMode := OsFilePermCode.ModeExclusive()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_05(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetuid()

	fileMode := OsFilePermCode.ModeSetuid()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_06(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSetgid()

	fileMode := OsFilePermCode.ModeSetgid()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_07(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeSticky()

	fileMode := OsFilePermCode.ModeSticky()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_08(t *testing.T) {

	expectedFileMode := OsFilePermCode.ModeIrregular()

	fileMode := OsFilePermCode.ModeIrregular()

	expectedOsPermCode := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	osPermCode, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err != nil {
		t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode(fileMode). "+
			"Error='%v' ", err.Error())
	}

	if expectedFileMode != osPermCode.Value() {
		t.Errorf("Error: Expected File Mode NOT equal to Actual File Mode. "+
			"expected File Mode Octal Value='%s' . Actual File Mode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

	if !expectedOsPermCode.Equal(osPermCode) {
		t.Errorf("Expected expectedOsPermCode==osPermCode. THEY ARE NOT EQUAL"+
			"expected expectedOsPermCode Octal Value='%s' . Actual osPermCode Octal Value= '%s'",
			strconv.FormatInt(int64(expectedFileMode), 8),
			strconv.FormatInt(int64(osPermCode.Value()), 8))
	}

}

func TestOsFilePermissionCode_GetNewFromFileMode_09(t *testing.T) {

	fileMode := os.FileMode(12577)

	_, err := OsFilePermCode.GetNewFromFileMode(fileMode)

	if err == nil {
		t.Error("Expected an error to be returned by OsFilePermCode." +
			"GetNewFromFileMode(fileMode) because fileMode is invalid. " +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_01(t *testing.T) {
	//     ModeNone                  -:      is a file

	letterCode := "-"

	expected := OsFilePermissionCode(OsFilePermCode.ModeNone())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_02(t *testing.T) {
	//     ModeDir               d:      is a directory

	letterCode := "d"

	expected := OsFilePermissionCode(OsFilePermCode.ModeDir())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_03(t *testing.T) {

	//     ModeAppend            a:      append-only

	letterCode := "a"

	expected := OsFilePermissionCode(OsFilePermCode.ModeAppend())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_04(t *testing.T) {

	//     ModeExclusive         l:      exclusive use

	letterCode := "l"

	expected := OsFilePermissionCode(OsFilePermCode.ModeExclusive())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_05(t *testing.T) {

	//     ModeTemporary         T:      temporary file; Plan 9 only

	letterCode := "T"

	expected := OsFilePermissionCode(OsFilePermCode.ModeTemporary())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_06(t *testing.T) {

	//     ModeSymlink           L:      symbolic link

	letterCode := "L"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSymlink())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_07(t *testing.T) {

	//     ModeDevice            D:      device file

	letterCode := "D"

	expected := OsFilePermissionCode(OsFilePermCode.ModeDevice())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_08(t *testing.T) {

	//     ModeNamedPipe         p:      named pipe (FIFO)

	letterCode := "p"

	expected := OsFilePermissionCode(OsFilePermCode.ModeNamedPipe())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_09(t *testing.T) {

	//     ModeSocket            S:      Unix domain socket

	letterCode := "S"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSocket())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_10(t *testing.T) {

	//     ModeSetuid            u:      setuid

	letterCode := "u"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSetuid())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_11(t *testing.T) {

	//     ModeSetgid            g:      setgid

	letterCode := "g"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSetgid())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_12(t *testing.T) {

	//     ModeCharDevice        c:      Unix character device, when ModeDevice is set

	letterCode := "c"

	expected := OsFilePermissionCode(OsFilePermCode.ModeCharDevice())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_13(t *testing.T) {

	//     ModeSticky            t:      sticky

	letterCode := "t"

	expected := OsFilePermissionCode(OsFilePermCode.ModeSticky())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_14(t *testing.T) {

	//     ModeIrregular         ?:      non-regular file; nothing else is known about this file

	letterCode := "?"

	expected := OsFilePermissionCode(OsFilePermCode.ModeIrregular())

	fPerm, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err != nil {
		t.Errorf("Error returned by fPerm.GetNewFromLetterCode(letterCode). "+
			"Error='%v' ", err.Error())
	}

	if !expected.Equal(fPerm) {
		t.Errorf("Error: Expected fPerm='%s'. Instead, fPerm='%s'",
			expected.String(), fPerm.String())
	}

	if expected.Value() != fPerm.Value() {
		t.Errorf("Error: Expected fPerm Value='%s'. Instead, fPerm='%s'",
			strconv.FormatInt(int64(expected.Value()), 10),
			strconv.FormatInt(int64(fPerm.Value()), 10))
	}

}

func TestOsFilePermissionCode_GetNewFromLetterCode_15(t *testing.T) {

	letterCode := "Q"

	_, err := OsFilePermCode.GetNewFromLetterCode(letterCode)

	if err == nil {
		t.Errorf("Expected an error to be returned by fPerm.GetNewFrom" +
			"LetterCode(letterCode) because the letter 'Q' is invalid. " +
			"However, NO ERROR WAS RETURNED!")
	}

}
