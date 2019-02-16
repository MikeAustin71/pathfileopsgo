package pathfileops

import (
	"os"
	"strconv"
	"testing"
)

func TestOsFilePermissionCode_IsValid_01(t *testing.T) {

	fpc := FilePermCode.ModeDir()

	if os.ModeDir != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDir. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	if os.ModeDir != osPerm.Value() {
		t.Errorf("Error: Expected OsFilePermissionCode instance to equal os.ModeDir. " +
			"The two are NOT EQUAL!")
	}

}

func TestOsFilePermissionCode_IsValid_02(t *testing.T) {

	fpc := FilePermCode.None()

	fmBase := os.FileMode(0)

	if fmBase != fpc {
		t.Errorf("Error: Expected fpc to be equal to os.ModeDir. It WAS NOT EQUAL!"+
			"fpc decimal value = %s octal value = %s",
			strconv.FormatInt(int64(fpc), 10), strconv.FormatInt(int64(fpc), 8))
	}

	osPerm := OsFilePermissionCode(fpc)

	if fmBase != osPerm.Value() {
		t.Errorf("Error: Expected OsFilePermissionCode instance to equal os.ModeDir. " +
			"The two are NOT EQUAL!")
	}

}
