package pathfileops

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const (
	alogtopTest2Text  = "../logTest/topTest2.txt"
	alogTestBottomDir = "../logTest/CmdrX"
	alogFile          = "CmdrX.log"
	aLoremIpsumTxt    = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum at congue nisi. Fusce viverra non urna et pulvinar. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sodales in nulla at ultricies. Pellentesque nisi mi, efficitur nec magna ac, laoreet efficitur ligula. Phasellus non justo justo. Suspendisse lacus dui, euismod hendrerit dignissim et, pellentesque fermentum ipsum. Duis tempus ex vitae dui commodo, sed sagittis arcu volutpat. Nam imperdiet, enim hendrerit maximus rhoncus, mauris enim convallis orci, non tincidunt leo tortor id lorem. Cras egestas orci non eros venenatis, quis aliquet orci maximus. Duis gravida augue sit amet tristique sagittis. Sed enim risus, suscipit at odio at, pretium facilisis elit. Morbi sit amet vestibulum ipsum. Ut eu turpis arcu."
)

func TestFileHelper_CleanDirStr_01(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")
	expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_02(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
	expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_03(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
	expectedDirName := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_04(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/../dir2/dir3")

	_, _, err := fh.CleanDirStr(testPath)

	if err == nil {
		t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). testPath='%v'", testPath)
	}

}

func TestFileHelper_CleanDirStr_05(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash(".../filesfortest/newfilesfortest")

	_, _, err := fh.CleanDirStr(testPath)

	if err == nil {
		t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). It was NOT. testPath='%v'", testPath)
	}

}

func TestFileHelper_CleanDirStr_06(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../..../filesfortest/newfilesfortest")

	_, _, err := fh.CleanDirStr(testPath)

	if err == nil {
		t.Errorf("Expected error to be returned by fh.CleanDirStr(testPath). It was NOT. testPath='%v'", testPath)
	}

}

func TestFileHelper_CleanDirStr_07(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")
	expectedDirName := fh.AdjustPathSlash(".")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_08(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("/xt_dirmgr_01_test.go")
	expectedDirName := fh.AdjustPathSlash("")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if !isDirEmpty {
		t.Error("Expected isDirEmpty='true'. Instead, isDirEmpty='false'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_09(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../../")
	expectedDirName := fh.AdjustPathSlash("../..")

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanDirStr_10(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("D:/")
	expectedDirName := "D:"

	cleanDirStr, isDirEmpty, err := fh.CleanDirStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanDirStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isDirEmpty {
		t.Error("Expected isDirEmpty='false'. Instead, isDirEmpty='true'")
	}

	if expectedDirName != cleanDirStr {
		t.Errorf("Expected cleanDirStr='%v'. Instead cleanDirStr='%v'", expectedDirName, cleanDirStr)
	}

}

func TestFileHelper_CleanFileNameExtStr_01(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")
	expectedFileNameExt := "newerFileForTest_01.txt"
	result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isFileNameEmpty {
		t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
	}

	if expectedFileNameExt != result {
		t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
	}

}

func TestFileHelper_CleanFileNameExtStr_02(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("newerFileForTest_01.txt")
	expectedFileNameExt := "newerFileForTest_01.txt"
	result, isFileNameEmpty, err := fh.CleanFileNameExtStr(testPath)

	if err != nil {
		t.Errorf("Error returned by fh.CleanFileNameExtStr(testPath). testPath='%v' Error='%v'", testPath, err.Error())
	}

	if isFileNameEmpty {
		t.Error("Expected isFileNameEmpty='false'. Instead, isFileNameEmpty='true'")
	}

	if expectedFileNameExt != result {
		t.Errorf("Expected fh.CleanFileNameExtStr to return '%v'. Instead, it returned '%v'", expectedFileNameExt, result)
	}

}

func TestFileHelper_CleanFileNameExtStr_03(t *testing.T) {
	fh := FileHelper{}
	testPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest/")
	_, _, err := fh.CleanFileNameExtStr(testPath)

	if err == nil {
		t.Errorf("Expected error returned by fh.CleanFileNameExtStr(testPath). testPath='%v'. But, no Error was returned. ", testPath)
	}

}

func TestFileHelper_CopyFile_01(t *testing.T) {

	fh := FileHelper{}
	srcFile := fh.AdjustPathSlash("..\\logTest\\Level01\\Level02\\TestFile001.txt")
	if !fh.DoesFileExist(srcFile) {
		fmgr, err := FileMgr{}.New(srcFile)

		if err != nil {
			t.Errorf("Error returned by FileMgr{}.New(srcFile). srcFile='%v'. Error='%v'", srcFile, err.Error())
		}

		err = fmgr.CreateDirAndFile()

		if err != nil {
			t.Errorf("Error returned by FileMgr{}.CreateDirAndFile(). srcFile='%v'. Error='%v'", srcFile, err.Error())
		}

		doesFileExist, err := fmgr.DoesThisFileExist()

		if err != nil {
			t.Errorf("Error returned by FileMgr{}.DoesThisFileExist(). srcFile='%v'. Error='%v'", srcFile, err.Error())
		}

		if !doesFileExist {
			t.Errorf("Failed to create Source File == '%v'", srcFile)
		}

	}

	destFile := fh.AdjustPathSlash("..\\logTest\\TestFile002.txt")

	if fh.DoesFileExist(destFile) {
		err := fh.DeleteDirFile(destFile)

		if err != nil {
			t.Error(fmt.Sprintf("Received Error while deleting destination file '%v', Error:", destFile), err)
		}
	}

	err := fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		t.Error(fmt.Sprintf("Error while Copying Source File, '%v' to  Destination File '%v', Error:", srcFile, destFile), err)
	}

	if !fh.DoesFileExist(destFile) {
		t.Error(fmt.Sprintf("Expected destination file: '%v' does NOT Exist.", destFile))
	}
}

func TestFileHelper_CreateFile_01(t *testing.T) {

	err := deleteTargetDir()
	if err != nil {
		t.Error("Failed to delete target Directory:", err)
	}

	err = createTargetDir()
	if err != nil {
		t.Error("Failed to delete target Directory:", err)
	}
}

func TestFileHelper_CreateFile_02(t *testing.T) {
	// Uses 'Create' to overwrite existing file
	tstFile := "..//logTest//testoverwrite//TestOverwrite001.txt"
	fh := FileHelper{}

	if fh.DoesFileExist(tstFile) {
		err := fh.DeleteDirFile(tstFile)
		if err != nil {
			t.Error(fmt.Sprintf("Error: Deletion Failed On File %v !", tstFile))
		}
	}

	if fh.DoesFileExist(tstFile) {
		t.Error(fmt.Sprintf("Error: Deletion Failed! File %v should not exist!", tstFile))
	}

	f, err := fh.CreateFile(tstFile)

	if err != nil {
		t.Error(fmt.Sprintf("Error: Create File Failed for file: %v", tstFile))
	}

	_, err4 := f.WriteString(aLoremIpsumTxt)

	if err4 != nil {
		_ = f.Close()
		t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err4)
	}

	err = f.Close()

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// Now recreate the original file. It should be
	// truncated with the old contents deleted.
	f2, err := fh.CreateFile(tstFile)

	if err != nil {
		t.Errorf("Error: Re-Creating File %v", tstFile)
	}

	fOvrWriteTxt := "Test Over Write and existing file using Create()"

	_, err5 := f2.WriteString(fOvrWriteTxt)

	if err5 != nil {
		_ = f2.Close()
		t.Error(fmt.Sprintf("Error Re-Writing to File: %v, Error: ", tstFile), err5)
	}

	err = f2.Close()

	if err != nil {
		t.Errorf("Error closing f2. %v", err.Error())
	}

	dat, err := ioutil.ReadFile(tstFile)

	if err != nil {
		t.Errorf("Error Reading Re-Written Text for File:'%v' Error='%v'", tstFile, err)
	}

	s := string(dat)

	if s != fOvrWriteTxt {
		t.Errorf("Was expecting to read text: '%v', instead received text: %v", fOvrWriteTxt, s)
	}

}

func TestFileHelper_ExtractBlankFileExt(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\dirmgr_test")

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned by fh.GetFileExtension(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty GetFileExtension for absent file extension to return 'true'. Instead, isEmpty='%v' ", isEmpty)
	}

	if result != "" {
		t.Errorf("Expected GetFileExtension to return empty result for absent file extension. Instead file extension='%v' ", result)
	}

}

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
