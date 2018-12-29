package pathfileops

import (
	"testing"
	"time"
)

func TestFileHelper_FilterFileName_01(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := "*.txt"
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_02(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := "*.txt"
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
			"fmtstr='%v' fOlderThanStr='%v' Error='%v'",
			fmtstr, fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_03(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)
	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fOlderThanStr). "+
			"fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("File was NOT found. File should have been found. fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_04(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := "*.txt"
	filesOlderThan := time.Time{}
	fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("File was NOT found. File should have been found. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_05(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	filesOlderThan := time.Time{}
	fNewerThanStr := "2017-09-01 00:00:00.000000000 -0500 CDT"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("File was NOT found. File should have been found. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_06(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := "*.txt"
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if isFound {
		t.Errorf("It was expected that this File would NOT be found. It WAS Found. "+
			"Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_07(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if isFound {
		t.Errorf("It was expected that this file would NOT be Found. Instead, it WAS found. "+
			"Error! fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_08(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := "*.htm"
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if isFound {
		t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_09(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	//filesOlderThan := time.Time{}
	fOlderThanStr := "2017-08-01 00:00:00.000000000 -0500 CDT"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if isFound {
		t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_10(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	filesOlderThan := time.Time{}

	fNewerThanStr := "2017-12-20 00:00:00.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if isFound {
		t.Errorf("Expected that File was NOT found. Instead, File WAS found - Error. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FilterFileName_11(t *testing.T) {

	fia := FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fModTimeStr). "+
			"fmtstr='%v' fModTimeStr='%v' Error='%v'", fmtstr, fModTimeStr, err.Error())
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(true)

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	fh := FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		t.Errorf("Expected that File would be found. However, File WAS NOT found - Error. "+
			"fia.Name()='%v fia.ModTime()='%v'", fia.Name(), fia.ModTime().Format(fmtstr))
	}

}

func TestFileHelper_FindFilesInPath(t *testing.T) {

}

func TestFileHelper_GetFileExtension_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '.go' for valid file extension Instead result='%v' ", result)
	}

}

func TestFileHelper_GetFileExtension_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'.  "+
			"Instead file extension='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("dirmgr_test")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead file extension ='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead, file extension='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
			"Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_07(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_08(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("test.....go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_09(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err == nil {
		t.Errorf("Expected an error to be returned from fh.GetFileExt(commonDir). "+
			"commonDir='%v' NO Error was returned!", commonDir)
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileNameWithExt_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
			false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
			true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_07(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_08(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".go")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err == nil {
		t.Errorf("Expected error returned from fh.GetFileNameWithExt(commonDir). "+
			"Instead, NO ERROR was returned. commonDir='%v'  ", commonDir)
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_09(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err == nil {
		t.Errorf("Error error returned from fh.GetFileNameWithExt(commonDir). Result- commonDir='%v' No Error Returned!", commonDir)
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_01(t *testing.T) {

	rawPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash(rawPath)

	firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if firstCharIdx != 3 {
		t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
	}

	expectedLastIdx := len(adjustedPath) - 1

	if expectedLastIdx != lastCharIdx {
		t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
	}

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_02(t *testing.T) {

	rawPath := "D:/filesfortest/newfilesfortest/newerFileForTest_01.txt"
	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash(rawPath)

	firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if firstCharIdx != 3 {
		t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
	}

	expectedLastIdx := len(adjustedPath) - 1

	if expectedLastIdx != lastCharIdx {
		t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
	}

}
