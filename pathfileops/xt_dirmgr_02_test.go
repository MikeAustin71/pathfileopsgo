package pathfileops

import (
	appLib "MikeAustin71/pathfileopsgo/appLibs"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDirMgr_FindWalkDirFiles_01(t *testing.T) {

	err := DirMgr02SetupDirWalkTests()

	if err != nil {
		t.Errorf("Error return from DirMgr02SetupDirWalkTests(). "+
			"Error='%v' ", err.Error())
		return
	}

	fh := FileHelper{}

	baseDirPath := fh.AdjustPathSlash("../dirwalktests/dir01")

	dir, err := fh.MakeAbsolutePath(baseDirPath)

	if err != nil {
		t.Errorf("fh.MakeAbsolutePath(baseDirPath) baseDirPath='%v' Error='%v'",
			baseDirPath, err.Error())
		return
	}

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	dMgr, err := DirMgr{}.New(dir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'",
			dir, err.Error())
	}

	if !dMgr.doesAbsolutePathExist {
		t.Errorf("Expected target directory to exist. I does NOT exist. "+
			"dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
			"filesOlderThan, filesNewerThan). dir='%v' Error='%v'",
			dir, err.Error())
	}

	if dirTreeInfo.FoundFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files in directory tree. Instead, found %v files.",
			dirTreeInfo.FoundFiles.GetArrayLength())
	}

	foundnewTestFile3 := false
	foundOldReadingFile := false

	for i := 0; i < dirTreeInfo.FoundFiles.GetArrayLength(); i++ {
		if dirTreeInfo.FoundFiles.FMgrs[i].fileNameExt == "newerFileForTest_03.txt" {
			foundnewTestFile3 = true
		}

		if dirTreeInfo.FoundFiles.FMgrs[i].fileNameExt == "006870_ReadingFiles.htm" {
			foundOldReadingFile = true
		}
	}

	if !foundnewTestFile3 {
		t.Error("Expected FoundFiles would include newerFileForTest_03.txt. It did NOT!")
	}

	if !foundOldReadingFile {
		t.Error("Expected FoundFiles would include 006870_ReadingFiles.htm. It did NOT!")
	}

	if dirTreeInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected the number of directories found would equal 3. It did NOT! "+
			"Number of directories= '%v'", dirTreeInfo.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

	foundDir3 := false

	for j := 0; j < dirTreeInfo.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dirTreeInfo.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dirTreeInfo.DirMgrs. "+
			"This directory was NOT found!", dir3)
	}

}

func TestDirMgr_FindWalkDirFiles_02(t *testing.T) {

	err := DirMgr02SetupDirWalkTests()

	if err != nil {
		t.Errorf("Error returned by DirMgr02SetupDirWalkTests(). "+
			"Error='%v' ", err.Error())
		return
	}

	fh := FileHelper{}
	dir := fh.AdjustPathSlash("../dirwalktests/dir01")

	searchPattern := "*Files.htm"
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	dMgr, err := DirMgr{}.New(dir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'",
			dir, err.Error())
	}

	if !dMgr.doesAbsolutePathExist {
		t.Errorf("Expected target directory to exist. I does NOT exist. "+
			"dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dWalker, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
			"filesOlderThan, filesNewerThan). dir='%v' Error='%v'", dir, err.Error())
	}

	if dWalker.FoundFiles.GetArrayLength() != 2 {
		t.Errorf("Expected to find 2-files in directory tree. Instead, "+
			"found %v files.", dWalker.FoundFiles.GetArrayLength())
	}

	foundWritingFiles := false
	foundOldReadingFile := false

	for i := 0; i < dWalker.FoundFiles.GetArrayLength(); i++ {
		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "006890_WritingFiles.htm" {
			foundWritingFiles = true
		}

		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "006870_ReadingFiles.htm" {
			foundOldReadingFile = true
		}
	}

	if !foundWritingFiles {
		t.Error("Expected FoundFiles would include 006890_WritingFiles.htm. " +
			"It did NOT!")
	}

	if !foundOldReadingFile {
		t.Error("Expected FoundFiles would include 006870_ReadingFiles.htm. " +
			"It did NOT!")
	}

	if dWalker.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected the number of directories found would equal 3. "+
			"It did NOT! Number of directories= '%v'",
			dWalker.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

	foundDir3 := false

	for j := 0; j < dWalker.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dWalker.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dWalker.DirMgrs. "+
			"This directory was NOT found!", dir3)
	}

}

func TestDirMgr_FindWalkDirFiles_03(t *testing.T) {

	err := DirMgr02SetupDirWalkTests()

	if err != nil {
		t.Errorf("Error returned by DirMgr02SetupDirWalkTests(). "+
			"Error='%v' ", err.Error())
		return
	}

	fh := FileHelper{}
	dir := fh.AdjustPathSlash("../dirwalktests/dir01")
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

	searchPattern := ""
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		t.Errorf("Failed to parse FilesOlderThan time value. "+
			"fmtstr='%v' fOlderThanStr='%v'  Error='%v'",
			fmtstr, fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}
	filesNewerThan.AddDate(0, 1, 0)

	dMgr, err := DirMgr{}.New(dir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). "+
			"dir='%v' Error='%v'", dir, err.Error())
	}

	if !dMgr.doesAbsolutePathExist {
		t.Errorf("Expected target directory to exist. I does NOT exist. "+
			"dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)
	}
	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dWalker, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
			"filesOlderThan, filesNewerThan). dir='%v' Error='%v'", dir, err.Error())
	}

	if dWalker.FoundFiles.GetArrayLength() != 3 {
		t.Errorf("Expected to find 3-files in directory tree. Instead, "+
			"found %v files.", dWalker.FoundFiles.GetArrayLength())
	}

	foundOldWritingFiles := false
	foundOldReadingFile := false
	foundOldTestHtmFile := false

	for i := 0; i < dWalker.FoundFiles.GetArrayLength(); i++ {
		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "006890_WritingFiles.htm" {
			foundOldWritingFiles = true
		}

		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "006870_ReadingFiles.htm" {
			foundOldReadingFile = true
		}

		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "test.htm" {
			foundOldTestHtmFile = true
		}

	}

	if !foundOldWritingFiles {
		t.Error("Expected FoundFiles would include 006890_WritingFiles.htm. " +
			"It did NOT!")
	}

	if !foundOldReadingFile {
		t.Error("Expected FoundFiles would include 006870_ReadingFiles.htm. " +
			"It did NOT!")
	}

	if !foundOldTestHtmFile {
		t.Error("Expected FoundFiles would include test.htm. " +
			"It did NOT!")
	}

	if dWalker.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected the number of directories found would equal 3. "+
			"It did NOT! Number of directories= '%v'",
			dWalker.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

	foundDir3 := false

	for j := 0; j < dWalker.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dWalker.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dWalker.DirMgrs. "+
			"This directory was NOT found!", dir3)
	}

}

func TestDirMgr_FindWalkDirFiles_04(t *testing.T) {

	err := DirMgr02SetupDirWalkTests()

	if err != nil {
		t.Errorf("Error returned by DirMgr02SetupDirWalkTests(). "+
			"Error='%v' ", err.Error())
		return
	}

	fh := FileHelper{}
	dir := fh.AdjustPathSlash("../dirwalktests/dir01")
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fNewerThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

	searchPattern := ""
	filesOlderThan := time.Time{}

	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Failed to parse FilesNewerThan time value. fmtstr='%v' "+
			"fNewerThanStr='%v'  Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	dMgr, err := DirMgr{}.New(dir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'",
			dir, err.Error())
	}

	if !dMgr.doesAbsolutePathExist {
		t.Errorf("Expected target directory to exist. I does NOT exist. "+
			"dMgr.path='%v' dMgr.AbolutePath='%v'", dMgr.path, dMgr.absolutePath)
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dWalker, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, "+
			"filesOlderThan, filesNewerThan). dir='%v' Error='%v'", dir, err.Error())
	}

	if dWalker.FoundFiles.GetArrayLength() != 3 {
		t.Errorf("Expected to find 3-files in directory tree. Instead, "+
			"found %v files.", dWalker.FoundFiles.GetArrayLength())
	}

	foundNewerFile1 := false
	foundNewerFile2 := false
	foundNewerFile3 := false

	for i := 0; i < dWalker.FoundFiles.GetArrayLength(); i++ {
		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "newerFileForTest_01.txt" {
			foundNewerFile1 = true
		}

		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "newerFileForTest_02.txt" {
			foundNewerFile2 = true
		}

		if dWalker.FoundFiles.FMgrs[i].fileNameExt == "newerFileForTest_03.txt" {
			foundNewerFile3 = true
		}

	}

	if !foundNewerFile1 {
		t.Error("Expected FoundFiles would include newerFileForTest_01.txt. File was NOT found!")
	}

	if !foundNewerFile2 {
		t.Error("Expected FoundFiles would include newerFileForTest_02.txt. File was NOT found!")
	}

	if !foundNewerFile3 {
		t.Error("Expected FoundFiles would include newerFileForTest_03.txt. File was NOT found!")
	}

	if dWalker.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected the number of directories found would equal 3. It did NOT! "+
			"Number of directories= '%v'", dWalker.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalktests/dir01/dir02/dir03")

	foundDir3 := false

	for j := 0; j < dWalker.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dWalker.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dWalker.DirMgrs. This directory was NOT found!",
			dir3)
	}

}

func TestDirMgr_FindWalkDirFiles_05(t *testing.T) {

	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
	}

	searchPattern := ""
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2016-12-01 19:54:30.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). "+
			"fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
	}

	fNewerThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr,fNewerThanStr). "+
			"fmtstr='%v' fNewerThanStr='%v' Error='%v'", fmtstr, fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ORFILESELECTCRITERION

	dInfo, err := dMgr.FindWalkDirFiles(fsc)

	if dInfo.FoundFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
			dInfo.FoundFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.FoundFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected to find oldFile1='%v'. The file was NOT found!", oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected to find oldFile2='%v'. The file was NOT found!", oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected find oldFile3='%v'. The file was NOT found!", oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected to find newFile1='%v'. The file was NOT found!", newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected to find newFile2='%v'. The file was NOT found!", newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected to find newFile3='%v'. The file was NOT found!", newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. "+
			"Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalkdeletetests/dirdelete01/dirdelete02/dirdelete03")

	foundDir3 := false

	for j := 0; j < dInfo.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dInfo.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dInfo.DirMgrs. "+
			"This directory was NOT found!", dir3)
	}

}

func TestDirMgr_FindWalkDirFiles_06(t *testing.T) {

	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return

	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
			origDir, err.Error())
	}

	searchPattern1 := "*.txt"
	searchPattern2 := "*.htm"

	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ORFILESELECTCRITERION

	dInfo, err := dMgr.FindWalkDirFiles(fsc)

	if dInfo.FoundFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
			dInfo.FoundFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.FoundFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.FoundFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected to find oldFile1='%v'. The file was NOT found!", oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected to find oldFile2='%v'. The file was NOT found!", oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected find oldFile3='%v'. The file was NOT found!", oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected to find newFile1='%v'. The file was NOT found!", newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected to find newFile2='%v'. The file was NOT found!", newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected to find newFile3='%v'. The file was NOT found!", newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
			len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

	dir3 := fh.AdjustPathSlash("/dirwalkdeletetests/dirdelete01/dirdelete02/dirdelete03")

	foundDir3 := false

	for j := 0; j < dInfo.Directories.GetNumOfDirs(); j++ {
		if strings.Contains(dInfo.Directories.DirMgrs[j].path, dir3) {
			foundDir3 = true
		}
	}

	if !foundDir3 {
		t.Errorf("Expected to find Directory %v in dInfo.DirMgrs. "+
			"This directory was NOT found!", dir3)
	}

}

func TestDirMgr_DeleteWalkDirFiles_01(t *testing.T) {

	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
			origDir, err.Error())
	}

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.",
			dInfo.DeletedFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!", oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!", oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!", oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!", newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!", newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
			newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
			len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'",
			dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_02(t *testing.T) {

	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
			origDir, err.Error())
	}

	searchPattern := "*.htm"
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 3 {
		t.Errorf("Expected to find 3-files deleted. Instead, %v-files were deleted.",
			dInfo.DeletedFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
			oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
			oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
			oldFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
			len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_03(t *testing.T) {

	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return

	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
			origDir, err.Error())
	}

	searchPattern := ""
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fOlderThanStr). "+
			"fOlderThanStr='%v' Error='%v'", fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 3 {
		t.Errorf("Expected to find 3-files deleted. Instead, "+
			"%v-files were deleted.", dInfo.DeletedFiles.GetArrayLength())

	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. "+
			"The file was NOT deleted!", oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. "+
			"The file was NOT deleted!", oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. "+
			"The file was NOT deleted!", oldFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of "+
			"Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number "+
			"of directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_04(t *testing.T) {
	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
	}

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"

	searchPattern := "*.txt"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fOlderThanStr). "+
			"fOlderThanStr='%v' Error='%v'", fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 0 {
		t.Errorf("Expected to find 0-files deleted. Instead, %v-files were deleted.",
			dInfo.DeletedFiles.GetArrayLength())
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
			len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of directories "+
			"found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_05(t *testing.T) {
	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'",
			origDir, err.Error())
	}

	searchPattern := ""
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fNewerThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"

	filesOlderThan := time.Time{}

	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fNewerThanStr). "+
			"fNewerThanStr='%v' Error='%v'", fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 3 {
		t.Errorf("Expected to find 3-files deleted. Instead, %v-files "+
			"were deleted.", dInfo.DeletedFiles.GetArrayLength())

	}

	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
			newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
			newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
			newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of "+
			"Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_06(t *testing.T) {

	// origDir = D:\gowork\src\MikeAustin71\pathfileopsgo\dirwalkdeletetests\dirdelete01
	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
		return
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
	}

	searchPattern := ""
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fNewerThanStr := "2016-07-01 00:00:00.000000000 -0500 CDT"
	filesOlderThan := time.Now()
	filesOlderThan.AddDate(0, 1, 0)

	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		t.Errorf("Error returned from time.Parse(fmtstr, fNewerThanStr). "+
			"fNewerThanStr='%v' Error='%v'", fNewerThanStr, err.Error())
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files "+
			"were deleted.", dInfo.DeletedFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!", oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!", oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!", oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!", newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!", newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!", newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'",
			len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())

	}

}

func TestDirMgr_DeleteWalkDirFiles_07(t *testing.T) {
	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
	}

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectByFileMode = 0666
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, "+
			"%v-files were deleted.", dInfo.DeletedFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
			oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
			oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
			oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
			newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
			newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
			newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error "+
			"Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

func TestDirMgr_DeleteWalkDirFiles_08(t *testing.T) {
	origDir, err := DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
	}

	searchPattern1 := "*.txt"
	searchPattern2 := "*.htm"
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern1, searchPattern2}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectByFileMode = 0666
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files "+
			"were deleted.", dInfo.DeletedFiles.GetArrayLength())
	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	oldFile1Found := false
	oldFile2Found := false
	oldFile3Found := false

	newFile1Found := false
	newFile2Found := false
	newFile3Found := false

	for i := 0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",
			oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",
			oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",
			oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",
			newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",
			newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",
			newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of "+
			"Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of "+
			"directories found='%v'", dInfo.Directories.GetNumOfDirs())
	}

}

/*
	************************************************************************************
														Test and Setup Methods
	************************************************************************************
*/

func DirMgr02TestCreateCheckFiles99DirFiles() (string, error) {
	ePrefix := "TestFile: xt_dirmgr_01_test.go Func: testDirMgrCreateCheckFiles03DirFiles() "
	fh := FileHelper{}

	origDir, err := fh.MakeAbsolutePath("../checkfiles/checkfiles99/checkfiles999")

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error returned by fh.MakeAbsolutePath(\"../checkfiles/checkfiles99/checkfiles999\")"+
				"Error='%v' ", err.Error())
	}

	if fh.DoesFileExist(origDir) {

		err := os.RemoveAll(origDir)

		if err != nil {
			return "",
				fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). "+
					"origDir='%v'  Error='%v'", origDir, err.Error())
		}

	}

	if fh.DoesFileExist(origDir) {
		return "",
			fmt.Errorf(ePrefix+"Error: Attempted to delete origDir='%v'. "+
				"However, it still Exists!", origDir)
	}

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err = os.MkdirAll(origDir, ModePerm)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+"Error returned from os.MkdirAll(origDir, ModePerm). "+
				"origDir='%v' ModePerm='%v'  Error='%v'", origDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origDir) {
		return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origDir='%v'", origDir)
	}

	fileDir := origDir + string(os.PathSeparator)
	newFile1 := fileDir + "checkFile99001.txt"
	fp1, err := os.Create(newFile1)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error returned from os.Create(newFile1). newFile1='%v' Error='%v' ",
				newFile1, err.Error())
	}

	newFile2 := fileDir + "checkFile99002.txt"

	fp2, err := os.Create(newFile2)

	if err != nil {
		_ = fp1.Close()
		return "",
			fmt.Errorf(ePrefix+
				"Error returned from os.Create(newFile2). newFile2='%v' Error='%v' ",
				newFile2, err.Error())
	}

	newFile3 := fileDir + "checkFile99003.txt"

	fp3, err := os.Create(newFile3)

	if err != nil {
		_ = fp1.Close()
		_ = fp2.Close()
		return "",
			fmt.Errorf(ePrefix+
				"Error returned from os.Create(newFile3). newFile3='%v' Error='%v' ",
				newFile3, err.Error())
	}

	newFile4 := fileDir + "checkFile99004.txt"

	fp4, err := os.Create(newFile4)

	if err != nil {
		_ = fp1.Close()
		_ = fp2.Close()
		_ = fp3.Close()
		return "",
			fmt.Errorf(ePrefix+
				"Error returned from os.Create(newFile4). newFile4='%v' Error='%v' ",
				newFile4, err.Error())
	}

	du := appLib.DateTimeUtility{}

	_, err = fp4.WriteString(du.GetDateTimeYMDAbbrvDowNano(time.Now()))

	if err != nil {
		_ = fp1.Close()
		_ = fp2.Close()
		_ = fp3.Close()
		_ = fp4.Close()
		return "",
			fmt.Errorf(ePrefix+
				"Error returned from fp4.WriteString(du.GetDateTimeYMDAbbrvDowNano"+
				"(time.Now())). Error='%v' ", err.Error())
	}

	_ = fp1.Close()
	_ = fp2.Close()
	_ = fp3.Close()
	_ = fp4.Close()

	return origDir, nil
}

func DirMgr02TestSetupFileWalkDeleteFiles() (string, error) {

	ePrefix := "xt_dirmgr_02_test.go Func: DirMgr02TestSetupFileWalkDeleteFiles() "

	fh := FileHelper{}

	origDir, err := fh.MakeAbsolutePath("../dirwalkdeletetests/dirdelete01")

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error Return from fh.MakeAbsolutePath(\"../dirwalkdeletetests/dirdelete01\")"+
				"Error='%v'", err.Error())
	}

	err = os.RemoveAll(origDir)

	if err != nil {
		return "", fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). "+
			"origDir='%v'  Error='%v'", origDir, err.Error())
	}

	time.Sleep(100 * time.Millisecond)

	if fh.DoesFileExist(origDir) {

		err = os.RemoveAll(origDir)

		if err != nil {
			return "", fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). "+
				"origDir='%v'  Error='%v'", origDir, err.Error())
		}

		time.Sleep(100 * time.Millisecond)

	}

	if fh.DoesFileExist(origDir) {

		return "", fmt.Errorf(ePrefix+"Error: Attempted to delete origDir='%v'. "+
			"However, it still Exists!", origDir)

	}

	origFullDir := origDir + string(os.PathSeparator) + "dirdelete02" +
		string(os.PathSeparator) + "dirdelete03"

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err = os.MkdirAll(origFullDir, ModePerm)

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"Error returned from os.MkdirAll(origFullDir, ModePerm). "+
			"origDir='%v' ModePerm='%v'  Error='%v'", origFullDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origFullDir) {
		return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origFullDir='%v'", origFullDir)
	}

	dirOldFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/oldfilesfortest")

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"Error returned by fh.MakeAbsolutePath(\"../filesfortest/oldfilesfortest\") "+
			"Error='%v'", err.Error())
	}

	// Copy Old Files

	if !fh.DoesFileExist(dirOldFilesForTest) {
		return "", fmt.Errorf(ePrefix+"Error: Old Files Directory does NOT exist! "+
			"dirOldFilesForTest='%v'", dirOldFilesForTest)

	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"

	srcFile := dirOldFilesForTest + string(os.PathSeparator) + oldFile1
	destFile := origDir + string(os.PathSeparator) + oldFile1

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'",
			srcFile, destFile, err)
	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + oldFile2

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"Error while Copying Source File, '%v' to  Destination File '%v', "+
			"Error:'%v'", srcFile, destFile, err)

	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile3

	destFile = origDir + string(os.PathSeparator) + "dirdelete02" +
		string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + oldFile3

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+"Error while Copying Source File, '%v' to  Destination File '%v', "+
				"Error:'%v'", srcFile, destFile, err)
	}

	// Copy NewFromPathFileNameExtStr Files
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	dirNewFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/newfilesfortest")

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+"Error return from fh.MakeAbsolutePath(dirNewFilesForTest) "+
				"dirNewFilesForTest='%v' Error='%v' ", dirNewFilesForTest, err.Error())
	}

	if !fh.DoesFileExist(dirNewFilesForTest) {
		return "", fmt.Errorf(ePrefix+
			"FATAL ERROR: NewFromPathFileNameExtStr Files Directory does NOT exist! dirNewFilesForTest='%v'",
			dirNewFilesForTest)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile1
	destFile = origDir + string(os.PathSeparator) + newFile1

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'",
				srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) +
		newFile2

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  "+
			"Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) +
		"dirdelete03" + string(os.PathSeparator) + newFile3

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  "+
			"Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	return origDir, nil
}

// Set up Directory Tree, ../dirwalktests for tests!
func DirMgr02SetupDirWalkTests() error {

	ePrefix := "xt_dirmgr02.DirMgr02SetupDirWalkTests() "

	fh := FileHelper{}

	var err error

	dest3, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\"). "+
			"Error='%v' ", err.Error())
	}

	destDirMgr3, err := DirMgr{}.New(dest3)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error from DirMgr{}.NewFromPathFileNameExtStr(dest3) "+
			"dest3='%v' Error='%v' ", dest3, err.Error())
	}

	dest1, err := fh.MakeAbsolutePath("../dirwalktests/dir01")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01\"). "+
			"Error='%v' ", err.Error())
	}

	destDirMgr1, err := DirMgr{}.New(dest1)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error from DirMgr{}.NewFromPathFileNameExtStr(dest1) "+
			"dest1='%v' Error='%v' ", dest1, err.Error())
	}

	dest0, err := fh.MakeAbsolutePath("../dirwalktests")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.MakeAbsolutePath(\"../dirwalktests\"). "+
			"Error='%v' ", err.Error())
	}

	destDirMgr0, err := DirMgr{}.New(dest0)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error from DirMgr{}.NewFromPathFileNameExtStr(dest1) "+
			"dest1='%v' Error='%v' ", dest0, err.Error())
	}

	// If the directory tree ../dirwalktests/dir01/dir02/dir03
	// does not exist, create it.
	if !destDirMgr3.DoesDirMgrAbsolutePathExist() {

		err = destDirMgr3.MakeDir()

		if err != nil {
			return fmt.Errorf(ePrefix+"Error returned by destDirMgr.MakeDir(). "+
				"destDir='%v' Error='%v' ", destDirMgr3.absolutePath, err.Error())
		}

	}

	// Delete all files in the ../dirwalktests
	deleteCriteria := FileSelectionCriteria{}
	deleteCriteria.FileNamePatterns = make([]string, 0, 0)
	deleteCriteria.FilesNewerThan = time.Time{}
	deleteCriteria.FilesOlderThan = time.Time{}
	deleteCriteria.SelectByFileMode = 0

	dirTree0Info, err := destDirMgr0.FindWalkDirFiles(deleteCriteria)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error return from destDirMgr0.FindWalkDirFiles(deleteCriteria). "+
			"Error='%v' ", err.Error())
	}

	// If directory tree ../dirwalktests contains files,
	// delete them.
	if dirTree0Info.FoundFiles.GetArrayLength() > 0 {

		_, err = destDirMgr0.DeleteWalkDirFiles(deleteCriteria)

		if err != nil {
			return fmt.Errorf(ePrefix+"Error returned by destDirMgr0.DeleteWalkDirFiles(deleteCriteria). "+
				"Error='%v'", err.Error())
		}

	}

	dirNewFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/newfilesfortest")

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by fh.MakeAbsolutePath(\"../filesfortest/newfilesfortest\"). "+
			"Error= '%v' ", err.Error())
	}

	srcNewFilesForTest, err := DirMgr{}.New(dirNewFilesForTest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirNewFilesForTest). "+
			"dirNewFilesForTest='%v' Error='%v' ", dirNewFilesForTest, err.Error())
	}

	if !srcNewFilesForTest.DoesDirMgrPathExist() {
		return fmt.Errorf(ePrefix+"FATAL ERROR: Directory %v DOES NOT EXIST", dirNewFilesForTest)
	}

	dirOldFilesForTest, err := fh.MakeAbsolutePath("../filesfortest/oldfilesfortest")

	srcOldFilesForTest, err := DirMgr{}.New(dirOldFilesForTest)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirOldFilesForTest). "+
			"dirOldFilesForTest='%v' Error='%v' ", dirOldFilesForTest, err.Error())
	}

	if !srcOldFilesForTest.DoesDirMgrPathExist() {
		return fmt.Errorf(ePrefix+"FATAL ERROR: Directory %v DOES NOT EXIST", dirOldFilesForTest)
	}

	filesToCopySrc := make([]string, 6, 10)
	filesToCopyDest := make([]string, 6, 10)

	filesToCopySrc[0] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
		"newerFileForTest_01.txt"

	filesToCopyDest[0] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"newerFileForTest_01.txt"

	filesToCopySrc[1] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
		"newerFileForTest_02.txt"

	filesToCopyDest[1] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"newerFileForTest_02.txt"

	filesToCopySrc[2] = srcNewFilesForTest.GetAbsolutePathWithSeparator() +
		"newerFileForTest_03.txt"

	filesToCopyDest[2] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"newerFileForTest_03.txt"

	filesToCopySrc[3] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
		"test.htm"

	filesToCopyDest[3] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"test.htm"

	filesToCopySrc[4] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
		"006890_WritingFiles.htm"

	filesToCopyDest[4] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"006890_WritingFiles.htm"

	filesToCopySrc[5] = srcOldFilesForTest.GetAbsolutePathWithSeparator() +
		"006870_ReadingFiles.htm"

	filesToCopyDest[5] = destDirMgr1.GetAbsolutePathWithSeparator() +
		"006870_ReadingFiles.htm"

	for i := 0; i < len(filesToCopySrc); i++ {

		if !fh.DoesFileExist(filesToCopySrc[i]) {
			return fmt.Errorf(ePrefix+
				"FATAL ERROR: Source File %v DOES NOT EXIST!!!", filesToCopySrc[i])
		}

		err = fh.CopyFileByLinkByIo(filesToCopySrc[i], filesToCopyDest[i])

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by fh.CopyFileByLinkByIo(filesToCopySrc[i], filesToCopyDest[i]). \n"+
				"filesToCopySrc[i]='%v' filesToCopyDest[i]='%v' Error='%v' ",
				filesToCopySrc[i], filesToCopyDest[i], err.Error())
		}

	}

	return nil
}
