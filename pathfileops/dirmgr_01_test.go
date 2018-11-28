package common

import (
	"fmt"
	"testing"
	"strings"
	"os"
	"time"
)

const (
	logDir = "../../003_filehelper/logTest"
	// appDir    = "../../003_filehelper/app"
	commonDir = "../../003_filehelper/common"
)

func TestCleanDir(t *testing.T) {
	var expected, cleanDir, targetDir string

	fh := FileHelper{}
	targetDir = "..///..///003_filehelper//common"

	cleanDir = fh.CleanPathStr(targetDir)
	expected = fh.CleanPathStr(commonDir)
	if cleanDir != expected {
		t.Error(fmt.Sprintf("Expected Clean Version of %v, got: ", commonDir), cleanDir)
	}

}

func TestChangeDir(t *testing.T) {
	var err error
	var startDir, checkDir, targetDir string
	fh := FileHelper{}

	startDir, err = fh.GetAbsCurrDir()

	if err != nil {
		t.Error("GetAnsCurrDir() Failed:", err)
	}

	targetDir, err = fh.MakeAbsolutePath(logDir)

	if err != nil {
		t.Error("MakeAbsolutePath() Failed:", err)
	}

	err = fh.ChangeDir(targetDir)

	if err != nil {
		t.Error("ChangeDir() Failed:", err)
	}

	checkDir, err = fh.GetAbsCurrDir()

	if err != nil {
		t.Error("GetAbsCurrDir() 2 Failed:", err)
	}

	if checkDir != targetDir {
		t.Error("Target Dir != CheckDir")
	}

	err = fh.ChangeDir(startDir)

	if err != nil {
		t.Error("Change To Start Dir Failed:", err)
	}

	checkDir, err = fh.GetAbsCurrDir()

	if err != nil {
		t.Errorf("GetAbsCurrDir() 3 Failed. Error='%v'", err)
	}

	if checkDir != startDir {
		t.Error("Start Dir != CheckDir")
	}
}


func TestDirMgr_CopyIn_01(t *testing.T) {

	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if dMgr2.Path != origDir2 {
		t.Errorf("Expected original dMgr2.Path='%v'. Instead, dMgr2.Path='%v'", origDir2, dMgr2.Path)
	}


	dMgr2.CopyIn(&dMgr)

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyIn(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyIn(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.OriginalPath != dMgr.OriginalPath {
		t.Errorf("After CopyIn(), expected dMgr2.OriginalPath='%v'.  Instead, dMgr2.OriginalPath='%v'.",dMgr.OriginalPath, dMgr2.OriginalPath )
	}

	if dMgr2.Path != dMgr.Path {
		t.Errorf("After CopyIn(), expected dMgr2.Path='%v'.  Instead, dMgr2.Path='%v'.",dMgr.Path, dMgr2.Path )
	}

	if dMgr2.PathIsPopulated != dMgr.PathIsPopulated {
		t.Errorf("After CopyIn(), expected dMgr2.PathIsPopulated='%v'.  Instead, dMgr2.PathIsPopulated='%v'.",dMgr.PathIsPopulated, dMgr2.PathIsPopulated )
	}

	if dMgr2.PathDoesExist != dMgr.PathDoesExist {
		t.Errorf("After CopyIn(), expected dMgr2.PathDoesExist='%v'.  Instead, dMgr2.PathDoesExist='%v'.",dMgr.PathDoesExist, dMgr2.PathDoesExist )
	}

	if dMgr2.ParentPath != dMgr.ParentPath {
		t.Errorf("After CopyIn(), expected dMgr2.ParentPath='%v'.  Instead, dMgr2.ParentPath='%v'.",dMgr.ParentPath, dMgr2.ParentPath )
	}

	if dMgr2.ParentPathIsPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("After CopyIn(), expected dMgr2.ParentPathIsPopulated='%v'.  Instead, dMgr2.ParentPathIsPopulated='%v'.",dMgr.ParentPathIsPopulated, dMgr2.ParentPathIsPopulated )
	}

	if dMgr2.RelativePath != dMgr.RelativePath {
		t.Errorf("After CopyIn(), expected dMgr2.RelativePath='%v'.  Instead, dMgr2.RelativePath='%v'.",dMgr.RelativePath, dMgr2.RelativePath )
	}

	if dMgr2.RelativePathIsPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("After CopyIn(), expected dMgr2.RelativePathIsPopulated='%v'.  Instead, dMgr2.RelativePathIsPopulated='%v'.",dMgr.RelativePathIsPopulated, dMgr2.RelativePathIsPopulated )
	}

	if dMgr2.AbsolutePath != dMgr.AbsolutePath {
		t.Errorf("After CopyIn(), expected dMgr2.AbsolutePath='%v'.  Instead, dMgr2.AbsolutePath='%v'.",dMgr.AbsolutePath, dMgr2.AbsolutePath )
	}

	if dMgr2.AbsolutePathIsPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("After CopyIn(), expected dMgr2.AbsolutePathIsPopulated='%v'.  Instead, dMgr2.AbsolutePathIsPopulated='%v'.",dMgr.AbsolutePathIsPopulated, dMgr2.AbsolutePathIsPopulated )
	}

	if dMgr2.AbsolutePathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("After CopyIn(), expected dMgr2.AbsolutePathDoesExist='%v'.  Instead, dMgr2.AbsolutePathDoesExist='%v'.",dMgr.AbsolutePathDoesExist, dMgr2.AbsolutePathDoesExist )
	}

	if dMgr2.AbsolutePathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("After CopyIn(), expected dMgr2.AbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.AbsolutePathDifferentFromPath='%v'.",dMgr.AbsolutePathDifferentFromPath, dMgr2.AbsolutePathDifferentFromPath )
	}

	if dMgr2.VolumeName != dMgr.VolumeName {
		t.Errorf("After CopyIn(), expected dMgr2.VolumeName='%v'.  Instead, dMgr2.VolumeName='%v'.",dMgr.VolumeName, dMgr2.VolumeName )
	}

	if dMgr2.VolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("After CopyIn(), expected dMgr2.VolumeIsPopulated='%v'.  Instead, dMgr2.VolumeIsPopulated='%v'.",dMgr.VolumeIsPopulated, dMgr2.VolumeIsPopulated )
	}

	if !dMgr2.Equal(&dMgr) {
		t.Error("After CopyIn(), expected dMgr2 to EQUAL dMgr. It did NOT!" )
	}

}

func TestDirMgr_CopyOut_01(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if dMgr2.Path != origDir2 {
		t.Errorf("Expected original dMgr2.Path='%v'. Instead, dMgr2.Path='%v'", origDir2, dMgr2.Path)
	}

	dMgr2 = dMgr.CopyOut()

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.OriginalPath != dMgr.OriginalPath {
		t.Errorf("After CopyOut(), expected dMgr2.OriginalPath='%v'.  Instead, dMgr2.OriginalPath='%v'.",dMgr.OriginalPath, dMgr2.OriginalPath )
	}

	if dMgr2.Path != dMgr.Path {
		t.Errorf("After CopyOut(), expected dMgr2.Path='%v'.  Instead, dMgr2.Path='%v'.",dMgr.Path, dMgr2.Path )
	}

	if dMgr2.PathIsPopulated != dMgr.PathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.PathIsPopulated='%v'.  Instead, dMgr2.PathIsPopulated='%v'.",dMgr.PathIsPopulated, dMgr2.PathIsPopulated )
	}

	if dMgr2.PathDoesExist != dMgr.PathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.PathDoesExist='%v'.  Instead, dMgr2.PathDoesExist='%v'.",dMgr.PathDoesExist, dMgr2.PathDoesExist )
	}

	if dMgr2.ParentPath != dMgr.ParentPath {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPath='%v'.  Instead, dMgr2.ParentPath='%v'.",dMgr.ParentPath, dMgr2.ParentPath )
	}

	if dMgr2.ParentPathIsPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPathIsPopulated='%v'.  Instead, dMgr2.ParentPathIsPopulated='%v'.",dMgr.ParentPathIsPopulated, dMgr2.ParentPathIsPopulated )
	}

	if dMgr2.RelativePath != dMgr.RelativePath {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePath='%v'.  Instead, dMgr2.RelativePath='%v'.",dMgr.RelativePath, dMgr2.RelativePath )
	}

	if dMgr2.RelativePathIsPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePathIsPopulated='%v'.  Instead, dMgr2.RelativePathIsPopulated='%v'.",dMgr.RelativePathIsPopulated, dMgr2.RelativePathIsPopulated )
	}

	if dMgr2.AbsolutePath != dMgr.AbsolutePath {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePath='%v'.  Instead, dMgr2.AbsolutePath='%v'.",dMgr.AbsolutePath, dMgr2.AbsolutePath )
	}

	if dMgr2.AbsolutePathIsPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathIsPopulated='%v'.  Instead, dMgr2.AbsolutePathIsPopulated='%v'.",dMgr.AbsolutePathIsPopulated, dMgr2.AbsolutePathIsPopulated )
	}

	if dMgr2.AbsolutePathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDoesExist='%v'.  Instead, dMgr2.AbsolutePathDoesExist='%v'.",dMgr.AbsolutePathDoesExist, dMgr2.AbsolutePathDoesExist )
	}

	if dMgr2.AbsolutePathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.AbsolutePathDifferentFromPath='%v'.",dMgr.AbsolutePathDifferentFromPath, dMgr2.AbsolutePathDifferentFromPath )
	}

	if dMgr2.VolumeName != dMgr.VolumeName {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeName='%v'.  Instead, dMgr2.VolumeName='%v'.",dMgr.VolumeName, dMgr2.VolumeName )
	}

	if dMgr2.VolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeIsPopulated='%v'.  Instead, dMgr2.VolumeIsPopulated='%v'.",dMgr.VolumeIsPopulated, dMgr2.VolumeIsPopulated )
	}

	if !dMgr2.Equal(&dMgr) {
		t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!" )
	}

}


func TestDirMgr_DeleteAll_01(t *testing.T) {

	fh:=FileHelper{}
	// Set up target directories and files for deletion!
	origDir, err := DirMgr01TestCreateCheckFiles03DirFiles()

	if err != nil {
		t.Errorf("Error returned by DirMgr01TestCreateCheckFiles03DirFiles(). Error='%v'", err.Error())
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origDir). origDir='%v'  Error='%v'",origDir, err.Error())
	}

	err = dMgr.DeleteAll()

	if err !=nil {
		t.Errorf("Error returned by dMgr.DeleteAll(). dMgr.Path='%v'  dMgr.AbsolutePath='%v'  Error='%v'",dMgr.Path, dMgr.AbsolutePath, err.Error())
	}

	if dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePath to be deleted. Instead, it Exists! dMgr.AbsolutePath='%v'", dMgr.AbsolutePath)
	}

	if fh.DoesFileExist(origDir) {
		t.Errorf("Expected origDir to be deleted. Instead, it Exists! origDir='%v'", origDir)
	}

}

func TestDirMgr_DeleteWalkDirFiles_31(t *testing.T) {

	origDir, err :=  DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). Error='%v'", err.Error())
	}

	fh := FileHelper{}

	if ! fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origDir). origDir='%v' Error='%v'", origDir, err.Error())
	}

	searchPattern1 := "*.txt"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr,fOlderThanStr)
	if err!= nil {
		t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern1}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ORFILESELECTCRITERION


	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.", dInfo.DeletedFiles.GetArrayLength())
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

	for i:=0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",newFile3)
	}


	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetArrayLength() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'", dInfo.Directories.GetArrayLength())
	}

}

func TestDirMgr_DeleteWalkDirFiles_32(t *testing.T) {

	origDir, err :=  DirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		t.Errorf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). Error='%v'", err.Error())
	}

	fh := FileHelper{}

	if ! fh.DoesFileExist(origDir) {
		t.Errorf("Error: The target directory does NOT Exist! origDir='%v'", origDir)
	}

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origDir). origDir='%v' Error='%v'", origDir, err.Error())
	}

	searchPattern1 := "*.txt"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 19:54:30.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr,fOlderThanStr)
	if err!= nil {
		t.Errorf("Error returned from time.Parse(fmtstr,fOlderThanStr). fmtstr='%v' fOlderThanStr='%v' Error='%v'", fmtstr, fOlderThanStr, err.Error())
	}

	filesNewerThan := time.Time{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern1}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ORFILESELECTCRITERION


	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		t.Errorf("Expected to find 6-files deleted. Instead, %v-files were deleted.", dInfo.DeletedFiles.GetArrayLength())
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

	for i:=0; i < dInfo.DeletedFiles.GetArrayLength(); i++ {

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile1) {
			oldFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile2) {
			oldFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, oldFile3) {
			oldFile3Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile1) {
			newFile1Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile2) {
			newFile2Found = true
		}

		if strings.Contains(dInfo.DeletedFiles.FMgrs[i].FileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		t.Errorf("Expected deletion of oldFile1='%v'. The file was NOT deleted!",oldFile1)
	}

	if oldFile2Found == false {
		t.Errorf("Expected deletion of oldFile2='%v'. The file was NOT deleted!",oldFile2)
	}

	if oldFile3Found == false {
		t.Errorf("Expected deletion of oldFile3='%v'. The file was NOT deleted!",oldFile3)
	}

	if newFile1Found == false {
		t.Errorf("Expected deletion of newFile1='%v'. The file was NOT deleted!",newFile1)
	}

	if newFile2Found == false {
		t.Errorf("Expected deletion of newFile2='%v'. The file was NOT deleted!",newFile2)
	}

	if newFile3Found == false {
		t.Errorf("Expected deletion of newFile3='%v'. The file was NOT deleted!",newFile3)
	}


	if len(dInfo.ErrReturns) != 0 {
		t.Errorf("Expected zero Error Returns. Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetArrayLength() != 3 {
		t.Errorf("Expected 3-directories to be found. Instead, number of directories found='%v'", dInfo.Directories.GetArrayLength())
	}

}


func TestDirMgr_Equal_01(t *testing.T) {
	
	fh := FileHelper{}
	
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if dMgr2.Path != origDir2 {
		t.Errorf("Expected original dMgr2.Path='%v'. Instead, dMgr2.Path='%v'", origDir2, dMgr2.Path)
	}


	dMgr2 = dMgr.CopyOut()

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.OriginalPath != dMgr.OriginalPath {
		t.Errorf("After CopyOut(), expected dMgr2.OriginalPath='%v'.  Instead, dMgr2.OriginalPath='%v'.",dMgr.OriginalPath, dMgr2.OriginalPath )
	}

	if dMgr2.Path != dMgr.Path {
		t.Errorf("After CopyOut(), expected dMgr2.Path='%v'.  Instead, dMgr2.Path='%v'.",dMgr.Path, dMgr2.Path )
	}

	if dMgr2.PathIsPopulated != dMgr.PathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.PathIsPopulated='%v'.  Instead, dMgr2.PathIsPopulated='%v'.",dMgr.PathIsPopulated, dMgr2.PathIsPopulated )
	}

	if dMgr2.PathDoesExist != dMgr.PathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.PathDoesExist='%v'.  Instead, dMgr2.PathDoesExist='%v'.",dMgr.PathDoesExist, dMgr2.PathDoesExist )
	}

	if dMgr2.ParentPath != dMgr.ParentPath {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPath='%v'.  Instead, dMgr2.ParentPath='%v'.",dMgr.ParentPath, dMgr2.ParentPath )
	}

	if dMgr2.ParentPathIsPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPathIsPopulated='%v'.  Instead, dMgr2.ParentPathIsPopulated='%v'.",dMgr.ParentPathIsPopulated, dMgr2.ParentPathIsPopulated )
	}

	if dMgr2.RelativePath != dMgr.RelativePath {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePath='%v'.  Instead, dMgr2.RelativePath='%v'.",dMgr.RelativePath, dMgr2.RelativePath )
	}

	if dMgr2.RelativePathIsPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePathIsPopulated='%v'.  Instead, dMgr2.RelativePathIsPopulated='%v'.",dMgr.RelativePathIsPopulated, dMgr2.RelativePathIsPopulated )
	}

	if dMgr2.AbsolutePath != dMgr.AbsolutePath {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePath='%v'.  Instead, dMgr2.AbsolutePath='%v'.",dMgr.AbsolutePath, dMgr2.AbsolutePath )
	}

	if dMgr2.AbsolutePathIsPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathIsPopulated='%v'.  Instead, dMgr2.AbsolutePathIsPopulated='%v'.",dMgr.AbsolutePathIsPopulated, dMgr2.AbsolutePathIsPopulated )
	}

	if dMgr2.AbsolutePathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDoesExist='%v'.  Instead, dMgr2.AbsolutePathDoesExist='%v'.",dMgr.AbsolutePathDoesExist, dMgr2.AbsolutePathDoesExist )
	}

	if dMgr2.AbsolutePathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.AbsolutePathDifferentFromPath='%v'.",dMgr.AbsolutePathDifferentFromPath, dMgr2.AbsolutePathDifferentFromPath )
	}

	if dMgr2.VolumeName != dMgr.VolumeName {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeName='%v'.  Instead, dMgr2.VolumeName='%v'.",dMgr.VolumeName, dMgr2.VolumeName )
	}

	if dMgr2.VolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeIsPopulated='%v'.  Instead, dMgr2.VolumeIsPopulated='%v'.",dMgr.VolumeIsPopulated, dMgr2.VolumeIsPopulated )
	}

	if !dMgr2.Equal(&dMgr) {
		t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!" )
	}

	if !dMgr.Equal(&dMgr2) {
		t.Error("After CopyOut(), expected dMgr2 to EQUAL dMgr. It did NOT!" )
	}

}

func TestDirMgr_Equal_02(t *testing.T) {

	fh := FileHelper{}

	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../xxxxfiles/xxxfiles2")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if dMgr2.Path != origDir2 {
		t.Errorf("Expected original dMgr2.Path='%v'. Instead, dMgr2.Path='%v'", origDir2, dMgr2.Path)
	}


	dMgr2 = dMgr.CopyOut()

	// dMgr2 and dMgr are no longer EQUAL
	dMgr2.AbsolutePath = dMgr2.AbsolutePath + "x"

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.IsInitialized != dMgr.IsInitialized {
		t.Errorf("After CopyOut(), expected dMgr2.IsFInfoInitialized='%v'.  Instead, dMgr2.IsFInfoInitialized='%v'.",dMgr.IsInitialized, dMgr2.IsInitialized )
	}

	if dMgr2.OriginalPath != dMgr.OriginalPath {
		t.Errorf("After CopyOut(), expected dMgr2.OriginalPath='%v'.  Instead, dMgr2.OriginalPath='%v'.",dMgr.OriginalPath, dMgr2.OriginalPath )
	}

	if dMgr2.Path != dMgr.Path {
		t.Errorf("After CopyOut(), expected dMgr2.Path='%v'.  Instead, dMgr2.Path='%v'.",dMgr.Path, dMgr2.Path )
	}

	if dMgr2.PathIsPopulated != dMgr.PathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.PathIsPopulated='%v'.  Instead, dMgr2.PathIsPopulated='%v'.",dMgr.PathIsPopulated, dMgr2.PathIsPopulated )
	}

	if dMgr2.PathDoesExist != dMgr.PathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.PathDoesExist='%v'.  Instead, dMgr2.PathDoesExist='%v'.",dMgr.PathDoesExist, dMgr2.PathDoesExist )
	}

	if dMgr2.ParentPath != dMgr.ParentPath {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPath='%v'.  Instead, dMgr2.ParentPath='%v'.",dMgr.ParentPath, dMgr2.ParentPath )
	}

	if dMgr2.ParentPathIsPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.ParentPathIsPopulated='%v'.  Instead, dMgr2.ParentPathIsPopulated='%v'.",dMgr.ParentPathIsPopulated, dMgr2.ParentPathIsPopulated )
	}

	if dMgr2.RelativePath != dMgr.RelativePath {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePath='%v'.  Instead, dMgr2.RelativePath='%v'.",dMgr.RelativePath, dMgr2.RelativePath )
	}

	if dMgr2.RelativePathIsPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.RelativePathIsPopulated='%v'.  Instead, dMgr2.RelativePathIsPopulated='%v'.",dMgr.RelativePathIsPopulated, dMgr2.RelativePathIsPopulated )
	}

	if dMgr2.AbsolutePath == dMgr.AbsolutePath {
		t.Error("After modification, expected dMgr2.AbsolutePath to be different from dMgr.AbsolutePath. ERROR= They ARE EQUAL!")
	}

	if dMgr2.AbsolutePathIsPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathIsPopulated='%v'.  Instead, dMgr2.AbsolutePathIsPopulated='%v'.",dMgr.AbsolutePathIsPopulated, dMgr2.AbsolutePathIsPopulated )
	}

	if dMgr2.AbsolutePathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDoesExist='%v'.  Instead, dMgr2.AbsolutePathDoesExist='%v'.",dMgr.AbsolutePathDoesExist, dMgr2.AbsolutePathDoesExist )
	}

	if dMgr2.AbsolutePathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("After CopyOut(), expected dMgr2.AbsolutePathDifferentFromPath='%v'.  Instead, dMgr2.AbsolutePathDifferentFromPath='%v'.",dMgr.AbsolutePathDifferentFromPath, dMgr2.AbsolutePathDifferentFromPath )
	}

	if dMgr2.VolumeName != dMgr.VolumeName {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeName='%v'.  Instead, dMgr2.VolumeName='%v'.",dMgr.VolumeName, dMgr2.VolumeName )
	}

	if dMgr2.VolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("After CopyOut(), expected dMgr2.VolumeIsPopulated='%v'.  Instead, dMgr2.VolumeIsPopulated='%v'.",dMgr.VolumeIsPopulated, dMgr2.VolumeIsPopulated )
	}

	if dMgr2.Equal(&dMgr) {
		t.Error("After modification, expected dMgr2 to NOT EQUAL to dMgr. Wrong- dMgr2 == dMgr!" )
	}

	if dMgr.Equal(&dMgr2) {
		t.Error("After modification, expected dMgr to NOT EQUAL to dMgr2. Wrong- dMgr == dMgr2!" )
	}

}

func TestDirMgr_EqualPaths_01(t *testing.T) {
	fh := FileHelper{}

	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if !dMgr.EqualPaths(&dMgr2) {
		t.Error("Expected two paths to be EQUAL. Error: They were NOT!")
	}

}

func TestDirMgr_EqualPaths_02(t *testing.T) {
	fh := FileHelper{}

	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles1")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	origDir2 := fh.AdjustPathSlash("../testfiles/testfiles2")

	dMgr2, err := DirMgr{}.New(origDir2)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir2). origDir2=='%v'  Error='%v'", origDir2, err.Error())
	}

	if dMgr.EqualPaths(&dMgr2) {
		t.Error("Expected two paths to be NOT EQUAL. Error: They were EQUAL!")
	}

}

func DirMgr01TestCreateCheckFiles03DirFiles() (string, error) {
	ePrefix := "TestFile: dirmgr_01_test.go Func: DirMgr01TestCreateCheckFiles03DirFiles() "
	fh := FileHelper{}

	origDir :=  fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")

	if fh.DoesFileExist(origDir) {

		err :=	os.RemoveAll(origDir)

		if err != nil {
			return "", fmt.Errorf(ePrefix + "Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
		}

	}

	if fh.DoesFileExist(origDir){
		return "", fmt.Errorf(ePrefix + "Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
	}

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err := os.MkdirAll(origDir, ModePerm)

	if err != nil {
		return "", fmt.Errorf(ePrefix + "Error returned from os.MkdirAll(origDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origDir) {
		return "", fmt.Errorf(ePrefix + "Error: Failed to create directory! origDir='%v'", origDir)
	}

	fileDir := origDir + string(os.PathSeparator)
	newFile1 := fileDir + "checkFile30001.txt"
	fp1, err := os.Create(newFile1)

	if err!= nil{
		return "", fmt.Errorf(ePrefix + "Error returned from os.Create(newFile1). newFile1='%v' Error='%v' ", newFile1, err.Error())
	}

	defer fp1.Close()

	newFile2 := fileDir + "checkFile30002.txt"

	fp2, err := os.Create(newFile2)

	if err!= nil{
		return "", fmt.Errorf(ePrefix + "Error returned from os.Create(newFile2). newFile2='%v' Error='%v' ", newFile2, err.Error())
	}

	defer fp2.Close()

	newFile3 := fileDir + "checkFile30003.txt"

	fp3, err := os.Create(newFile3)

	if err!= nil{
		return "", fmt.Errorf(ePrefix + "Error returned from os.Create(newFile3). newFile3='%v' Error='%v' ", newFile3, err.Error())
	}

	defer fp3.Close()


	newFile4 := fileDir + "checkFile30004.txt"

	fp4, err := os.Create(newFile4)

	if err!= nil{
		return "", fmt.Errorf(ePrefix + "Error returned from os.Create(newFile4). newFile4='%v' Error='%v' ", newFile4, err.Error())
	}

	defer fp4.Close()

	du := DateTimeUtility{}

	fp4.WriteString(du.GetDateTimeYMDAbbrvDowNano(time.Now()))


	return origDir, nil

}