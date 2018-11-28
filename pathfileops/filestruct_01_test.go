package common

import (
	"os"
	"testing"
	"time"
)

func TestFileInfoPlus_CopyOut_01(t *testing.T) {
	fip := FileInfoPlus{}
	fip.SetName("newerFileForTest_01.txt")

	err :=
		fip.SetDirectoryPath("D:/gowork/src/MikeAustin71/pathfilego/003_filehelper/filesfortest/newfilesfortest")

	if err != nil {
		t.Errorf("Error returned from fip.SetDirectoryPath(). " +
			"Error='%v'",  err.Error())
	}


	fip.SetMode(0777)
	fip.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fip.SetModTime(fModTime)
	fip.SetIsDir(false)
	fip.SetSysDataSrc(nil)
	fip.SetIsFInfoInitialized(true)
	
	fip2 := fip.CopyOut()
	
	if fip.Name() != fip2.Name() {
		t.Errorf("Error CopyOut. Names do not match. fip.Name()= '%v'  fip2.Name()= '%v' ", fip.Name(), fip2.Name())
	}

	if fip.Size() != fip2.Size() {
		t.Errorf("Error CopyOut. File Sizes do not match. fip.Size()= '%v'  fip2.Size()= '%v' ", fip.Size(), fip2.Size())
	}

	if fip.Mode() != fip2.Mode() {
		t.Errorf("Error CopyOut. Modes do not match. fip.Mode()= '%v'  fip2.Mode()= '%v' ", fip.Mode(), fip2.Mode())
	}

	if fip.ModTime() != fip2.ModTime() {
		t.Errorf("Error CopyOut. ModTimes do not match. fip.ModTime()= '%v'  fip2.ModTime()= '%v' ", fip.ModTime(), fip2.ModTime())
	}

	if fip.Sys() != fip2.Sys() {
		t.Errorf("Error CopyOut. Names do not match. fip.Sys()= '%v'  fip2.Sys()= '%v' ", fip.Sys(), fip2.Sys())
	}

	if fip.IsDir() != fip2.IsDir() {
		t.Errorf("Error CopyOut. Names do not match. fip.IsDir()= '%v'  fip2.IsDir()= '%v' ", fip.IsDir(), fip2.IsDir())
	}

	if fip.DirPath() != fip2.DirPath() {
		t.Errorf("Error CopyOut. Names do not match. fip.DirPath()= '%v'  fip2.DirPath()= '%v' ", fip.DirPath(), fip2.DirPath())
	}

	if fip.IsFInfoInitialized != fip2.IsFInfoInitialized {
		t.Errorf("Error CopyOut. IsFInfoInitialized values do not match. fip.IsFInfoInitialized= '%v'  fip2.IsFInfoInitialized= '%v' ", fip.IsFInfoInitialized, fip2.IsFInfoInitialized)
	}

	if fip.IsDirPathInitialized != fip2.IsDirPathInitialized {
		t.Errorf("Error CopyOut. IsDirPathInitialized values do not match. fip.IsDirPathInitialized= '%v'  fip2.IsDirPathInitialized= '%v' ", fip.IsDirPathInitialized, fip2.IsDirPathInitialized)
	}

	if fip.CreateTimeStamp != fip2.CreateTimeStamp {
		t.Errorf("Error CopyOut. CreateTimeStamp values do not match. fip.CreateTimeStamp= '%v'  fip2.CreateTimeStamp= '%v' ", fip.CreateTimeStamp, fip2.CreateTimeStamp)
	}

}

func TestFileInfoPlus_NewFromFileInfo_01(t *testing.T) {
	fh:= FileHelper{}

	baseFileName := "newerFileForTest_01.txt"
	baseDirPath := "../filesfortest/newfilesfortest"

	absBaseDirPath, err := fh.MakeAbsolutePath(baseDirPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(baseDirPath). baseDirPath='%v' Error='%v'", baseDirPath, err.Error())
	}

	absPathFileName, _ := fh.AddPathSeparatorToEndOfPathStr(absBaseDirPath)
	absPathFileName = absPathFileName + baseFileName

	fInfo, err := fh.GetFileInfoFromPath(absPathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileName). absPathFileName='%v' Error='%v'", absPathFileName, err.Error())
	}

	fip := FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Name() != baseFileName {
		t.Errorf("Expected fip.Name()='%v'. Instead, fip.Name()='%v'.", baseFileName, fip.Name())
	}

	if fip.IsDir() == true {
		t.Error("Expected  fip.IsDir()=false. Instead, fip.IsDir()=true")
	}

	if fip.IsFInfoInitialized == false {
		t.Error("Expected fip.IsFInfoInitialized='true'.  Instead, fip.IsFInfoInitialized='false'")
	}

	if fip.CreateTimeStamp.IsZero() {
		t.Error("Expected fip.CreateTimeStamp to be a non-zero value. Instead fip.CreateTimeStamp=time.Zero.")
	}

	if fip.ModTime().IsZero() {
		t.Error("Expected fip.ModTime() to be non-zero.  Instead, fip.ModTime() is time.Zero.")
	}

	var testMode os.FileMode

	testMode = 0666 // rwxrwxrwx

	if fip.Mode() != testMode {
		t.Errorf("Expected fip.Mode()=777. Instead, fip.Mode()='%v'", fip.Mode())
	}

}

func TestFileInfoPlus_Equal_01(t *testing.T) {
	fh:= FileHelper{}

	baseFileName := "newerFileForTest_01.txt"

	baseDirPath := "D:/gowork/src/MikeAustin71/pathfilego/003_filehelper/filesfortest/newfilesfortest"

	absBaseDirPath, err := fh.MakeAbsolutePath(baseDirPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(baseDirPath). baseDirPath='%v' Error='%v'", baseDirPath, err.Error())
	}

	absPathFileName, _ := fh.AddPathSeparatorToEndOfPathStr(absBaseDirPath)
	absPathFileName = absPathFileName + baseFileName

	fInfo, err := fh.GetFileInfoFromPath(absPathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileName). absPathFileName='%v' Error='%v'", absPathFileName, err.Error())
	}

	fip := FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Name() != baseFileName {
		t.Errorf("Expected fip.Name()='%v'. Instead, fip.Name()='%v'.", baseFileName, fip.Name())
	}


	fip2 := FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Equal(&fip2) == false {
		t.Error("Expected  fip to EQUAL fip2. It DID NOT!")
	}

}

func TestFileInfoPlus_Equal_02(t *testing.T) {
	fh:= FileHelper{}

	baseFileName := "newerFileForTest_01.txt"
	baseDirPath := "D:/gowork/src/MikeAustin71/pathfilego/003_filehelper/filesfortest/newfilesfortest"

	absBaseDirPath, err := fh.MakeAbsolutePath(baseDirPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(baseDirPath). baseDirPath='%v' Error='%v'", baseDirPath, err.Error())
	}

	absPathFileName, _ := fh.AddPathSeparatorToEndOfPathStr(absBaseDirPath)
	absPathFileName = absPathFileName + baseFileName

	fInfo, err := fh.GetFileInfoFromPath(absPathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileName). absPathFileName='%v' Error='%v'", absPathFileName, err.Error())
	}

	fip := FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Name() != baseFileName {
		t.Errorf("Expected fip.Name()='%v'. Instead, fip.Name()='%v'.", baseFileName, fip.Name())
	}


	fip2 := FileInfoPlus{}.NewFromFileInfo(fInfo)

	err = fip2.SetDirectoryPath("XXX")

	if err != nil {
		t.Errorf("Error returned from fip2.SetDirectoryPath(\"XXX\").  Error='%v'",
			err.Error())
	}


	if fip.Equal(&fip2) == true {
		t.Error("Expected fip to NOT EQUAL fip2. Error- fip==fip2")
	}

}

func TestFileSelectionCriteria_01(t *testing.T) {

	fsc := FileSelectionCriteria{}

	// DEFAULT SHOULD BE ANDFILESELECTCRITERION
	if fsc.SelectCriterionMode != ANDFILESELECTCRITERION {
		t.Errorf("Expected default FileSelectionCriteria.SelectCriterionMode=ANDFILESELECTCRITERION. Instead, FileSelectionCriteria.SelectCriterionMode='%v'", fsc.SelectCriterionMode)
	}

}

func TestFileSelectCriterionMode_Text_01(t *testing.T) {
	var r FileSelectCriterionMode

	r = ANDFILESELECTCRITERION

	var s string

	s = r.String()

	if s != "AND File Select Criterion" {
		t.Errorf("Expected string 'AND File Select Criterion'. Instead, got %v", s)
	}
}

func TestFileSelectCriterionMode_Text_02(t *testing.T) {
	var r FileSelectCriterionMode

	r = ORFILESELECTCRITERION

	var s string

	s = r.String()

	if s != "OR File Select Criterion" {
		t.Errorf("Expected string 'OR File Select Criterion'. Instead, got %v", s)
	}
}

func TestFileSelectCriterionMode_Value_01(t *testing.T) {
	var r FileSelectCriterionMode

	var i int

	r = ANDFILESELECTCRITERION

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'ANDFILESELECTCRITERIA' value = 0. Instead, got %v", i)
	}
}

func TestFileSelectCriterionMode_Value_02(t *testing.T) {
	var r FileSelectCriterionMode

	var i int

	r = ORFILESELECTCRITERION

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'ORFILESELECTCRITERION' value = 1. Instead, got %v", i)
	}
}

