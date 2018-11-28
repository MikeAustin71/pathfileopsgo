package common

import (
	"fmt"
	"os"
	"testing"
)

func TestFileMgr_CopyOut_01 (t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_01_test.go")
	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
	fileName := "dirmgr_01_test"
	fileNameExt := "dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	fMgr2:= fileMgr.CopyOut()

	if fMgr2.FileName != fileName {
		t.Error(fmt.Sprintf("Expected CopyToThis to return FileName == '%v', instead got: ", fileName), fMgr2.FileName)
	}

	if fMgr2.FileExt != extName {
		t.Error(fmt.Sprintf("Expected CopyToThis to return FileExt == '%v', instead got: ", extName), fMgr2.FileExt)
	}

	if fMgr2.FileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected CopyToThis to return FileNameExt == '%v', instead got: ", fileNameExt), fMgr2.FileNameExt)
	}

	if fMgr2.DMgr.Path != expectedDir {
		t.Error(fmt.Sprintf("Expected CopyToThis to return Path == '%v', instead got: ", expectedDir), fMgr2.DMgr.Path)
	}

	result := fMgr2.Equal(&fileMgr)

	if result != true {
		t.Error("Expected Equal to return 'true' for fMgr2==fileMgr, instead got: ", result)
	}

}

func TestFileMgr_CopyFileMgr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh:= FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.AbsolutePath='%v' expectedFileNameExt='%v'  Error='%v'",dMgr.AbsolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.AbsolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.AbsolutePath, expectedFileNameExt, err.Error())
	}

	err = srcFMgr.CopyFileMgr(&destFMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgr(&destFMgr). srcFMgr.AbsolutePathFileName='%v'  destFMgr.AbsolutePathFileName='%v'  Error='%v'", srcFMgr.AbsolutePathFileName, destFMgr.AbsolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(destFMgr.AbsolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.AbsolutePathFileName)=true. Instead it was 'false' destFMgr.AbsolutePathFileName='%v'", destFMgr.AbsolutePathFileName)
	}

	if !destFMgr.AbsolutePathFileNameDoesExist {
		t.Error("Expected destFMgr.AbsolutePathFileNameDoesExist='true'.  ERROR  destFMgr.AbsolutePathFileNameDoesExist='false'")
	}

	err = fh.DeleteDirFile(destFMgr.AbsolutePathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.AbsolutePathFileName) destFMgr.AbsolutePathFileName='%v' Error='%v'", destFMgr.AbsolutePathFileName, err.Error())
	}

	if fh.DoesFileExist(destFMgr.AbsolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.AbsolutePathFileName)=false. Instead it was 'true' destFMgr.AbsolutePathFileName='%v'", destFMgr.AbsolutePathFileName)
	}

}


func TestFileMgr_CreateDirAndFile_01(t *testing.T) {
	fh := FileHelper{}
	testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
	fileMgr, err := FileMgr{}.New(testFile)

	if err != nil {
		t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
	}

	if fh.DoesFileExist(fileMgr.DMgr.AbsolutePath) {

		err = fh.DeleteDirPathAll(fileMgr.DMgr.AbsolutePath)

		if err != nil {
			t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.DMgr.AbsolutePath). "+
				" fileMgr.DMgr.AbsolutePath='%v'   Error='%v' ",fileMgr.DMgr.AbsolutePath, err.Error())
		}

	}

	if fh.DoesFileExist(fileMgr.DMgr.AbsolutePath) {
		t.Errorf(fmt.Sprintf("Error: Failed to delete existing path '%v'", fileMgr.DMgr.AbsolutePath))
	}

	err = fileMgr.CreateDirAndFile()

	if err != nil {
		t.Errorf("Failed to Create Directory and File '%v', received Error:'%v'", fileMgr.AbsolutePathFileName, err.Error())
	}


	if !fh.DoesFileExist(fileMgr.AbsolutePathFileName) {
		t.Errorf(fmt.Sprintf("File Verfication failed file '%v' DOES NOT EXIST", fileMgr.AbsolutePathFileName))
	}

	s := "Created by File:'filemgr_test.go' Test Method: TestFileHelper_CreateDirAndFile()"

	_, err = fileMgr.WriteStrToFile(s)

	if err != nil {
		t.Errorf("Received error from fileMgr.WriteStrToFile(s). s='%v'  Error='%v' ",s, err.Error())
	}


	err = fileMgr.CloseFile()

	if err != nil {
		t.Errorf("Received error from fileMgr.CloseFile(). fileMgr.AbsolutePathFileName='%v'  Error='%v' ",fileMgr.AbsolutePathFileName, err.Error())
	}

	err = fileMgr.DMgr.DeleteAll()

	if err != nil {
		t.Errorf("Error returned by fileMgr.DMgr.DeleteAll(). Attempted Deletion of %v. Error='%v'", fileMgr.AbsolutePathFileName, err.Error())
	}

}


func TestFileMgr_Equal_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_01_test.go")

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	fileMgr2 := fileMgr.CopyOut()

	// fh is now NOT Equal to fileMgr
	fileMgr2.DMgr.Path = ""

	if fileMgr2.Equal(&fileMgr) == true {
		t.Error("Expected Equal to return 'false' for fh==fileMgr, instead got: ", "false")
	}

}

func TestFileMgr_MoveFileToNewDirMgr_01(t *testing.T) {
	fh := FileHelper{}
	setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
	srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
	destDir := fh.AdjustPathSlash("..\\logTest")
	setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

	if fh.DoesFileExist(setupDestFile) {
		err := fh.DeleteDirFile(setupDestFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
		}

		if fh.DoesFileExist(setupDestFile) {
			t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
		}
	}

	if fh.DoesFileExist(srcFile) {
		err := fh.DeleteDirFile(srcFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
		}

		if fh.DoesFileExist(srcFile) {
			t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
		}
	}

	err := fh.CopyToNewFile(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.New(srcFile)

	if err!=nil {
		t.Errorf("Error returned from FileMgr{}.New(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	dMgr, err := DirMgr{}.New(destDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(destDir). destDir='%v' Error='%v'", destDir, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr). dMgr.Path='%v'  Error='%v'", dMgr.Path, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.AbsolutePathFileName) {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.AbsolutePathFileName='%v'",newFMgr.AbsolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.AbsolutePathFileName='%v' Error='%v'", newFMgr.AbsolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.AbsolutePathFileName='%v'",newFMgr.AbsolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.AbsolutePathFileName='%v'", newFMgr.AbsolutePathFileName)
	}

}

func TestFileMgr_MoveFileToNewDir_01(t *testing.T) {
	fh := FileHelper{}
	setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
	srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
	destDir := fh.AdjustPathSlash("..\\logTest")
	setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

	if fh.DoesFileExist(setupDestFile) {
		err := fh.DeleteDirFile(setupDestFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
		}

		if fh.DoesFileExist(setupDestFile) {
			t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
		}
	}

	if fh.DoesFileExist(srcFile) {
		err := fh.DeleteDirFile(srcFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
		}

		if fh.DoesFileExist(srcFile) {
			t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
		}
	}

	err := fh.CopyToNewFile(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.New(srcFile)

	if err!=nil {
		t.Errorf("Error returned from FileMgr{}.New(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}


	newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.AbsolutePathFileName) {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.AbsolutePathFileName='%v'",newFMgr.AbsolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.AbsolutePathFileName='%v' Error='%v'", newFMgr.AbsolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.AbsolutePathFileName='%v'",newFMgr.AbsolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.AbsolutePathFileName='%v'", newFMgr.AbsolutePathFileName)
	}

}

func TestFileMgr_New_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\dirmgr_01_test.go")
	fileName := "dirmgr_01_test"
	fileNameExt := "dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	if fileMgr.FileName != fileName {
		t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.FileName)
	}

	if fileMgr.FileExt != extName {
		t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.FileExt)
	}

	if fileMgr.FileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.FileNameExt)
	}

	if !fileMgr.IsInitialized {
		t.Error("Expected fileMgr.IsInitialized=='true', got:", fileMgr.IsInitialized)
	}

	if !fileMgr.FileNameIsPopulated {
		t.Error("Expected fileMgr.FileNameIsPopulated=='true', got:", fileMgr.FileNameIsPopulated)
	}

	if !fileMgr.FileNameExtIsPopulated {
		t.Error("Expected fileMgr.FileNameExtIsPopulated=='true', got:", fileMgr.FileNameExtIsPopulated)
	}

	if !fileMgr.FileExtIsPopulated {
		t.Error("Expected fileMgr.FileExtIsPopulated=='true', got:", fileMgr.FileExtIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated=='true', got:", fileMgr.AbsolutePathFileNameIsPopulated)
	}

}

func TestFileMgr_New_02 (t *testing.T) {

	path := ".\\filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err :=  FileMgr{}.New(path)

	if err != nil {
		t.Error("Error returned from FileHelper:GetPathFileNameElements() path== 'dirmgr_01_test.go' Error: ", err)
	}

	if fileMgr.FileNameExt != eFileNameExt {
		t.Error(fmt.Sprintf("Expected extracted FileNameExt == %v, instead got: ", path), fileMgr.FileNameExt)
	}

	if fileMgr.FileName != "filehelperexamples" {
		t.Error("Expected fileMgr.FileName== 'dirmgr_test', instead got: fileMgr.FileName==", fileMgr.FileName)
	}

	if fileMgr.FileExt != ".go" {
		t.Error("Expected fileMgr.FileExt== '.go', instead got: fileMgr.FileExt==", fileMgr.FileName)
	}

	if !fileMgr.DMgr.PathIsPopulated {
		t.Error("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated==", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameDoesExist {
		t.Error("Expected 'fileMgr.AbsolutePathFileNameDoesExist==true', instead got: fileMgr.AbsolutePathFileNameDoesExist==", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.DMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

}

func TestFileMgr_New_03 (t *testing.T) {

	path := "filehelperexamples"

	fileMgr, err := FileMgr{}.New(path)

	if err != nil {
		t.Errorf("Error returned from FileHelper:GetPathFileNameElements() path== '%v' Error: %v", path, err)
	}

	if fileMgr.FileName != path {
		t.Error("Expected fileMgr.FileName=='dirmgr_test', instead got:", fileMgr.FileName)
	}

	if fileMgr.DMgr.PathIsPopulated {
		t.Error("Expected fileMgr.PathIsPopulated==false, instead got:", fileMgr.DMgr.PathIsPopulated)
	}

	if fileMgr.FileExtIsPopulated {
		t.Error("Expected fileMgr.FileExtIsPopulated==false, instead got:", fileMgr.FileExtIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated==true, instead got:", fileMgr.AbsolutePathFileNameIsPopulated)
	}

	if fileMgr.DMgr.AbsolutePathIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathIsPopulated==false, instead got:", fileMgr.DMgr.AbsolutePathIsPopulated)
	}
}

func TestFileMgr_New_04 (t *testing.T) {

	path := "filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err :=  FileMgr{}.New(path)

	if err != nil {
		t.Error("Error returned from FileHelper:GetPathFileNameElements() path== 'dirmgr_01_test.go' Error: ", err)
	}

	if fileMgr.FileNameExt != eFileNameExt {
		t.Error(fmt.Sprintf("Expected extracted FileNameExt == %v, instead got: ", path), fileMgr.FileNameExt)
	}

	if fileMgr.FileName != "filehelperexamples" {
		t.Error("Expected fileMgr.FileName== 'dirmgr_test', instead got: fileMgr.FileName==", fileMgr.FileName)
	}

	if fileMgr.FileExt != ".go" {
		t.Error("Expected fileMgr.FileExt== '.go', instead got: fileMgr.FileExt==", fileMgr.FileName)
	}

	if fileMgr.DMgr.PathIsPopulated {
		t.Error("Expected 'fileMgr.PathIsPopulated==false', instead got: fileMgr.PathIsPopulated==", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameDoesExist {
		t.Error("Expected 'fileMgr.AbsolutePathFileNameDoesExist==true', instead got: fileMgr.AbsolutePathFileNameDoesExist==", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated == 'true'.  Instead, it is 'false'")
	}

	if fileMgr.DMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'false'.  Instead, it is 'true'")
	}

}

func TestFileMgr_NewFromFileInfo_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName:= "newerFileForTest_01"
	expectedExt := ".txt"
	fh:= FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
	}

	fileMgr, err := FileMgr{}.NewFromFileInfo(absPath, info)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info). absPath='%v' info.Name()='%v'  Error='%v'", absPath, info.Name(), err.Error())
	}


	if fileMgr.FileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.FileNameExt == %v, instead fileMgr.FileNameExt='%v' ", expectedFileNameExt, fileMgr.FileNameExt)
	}

	if fileMgr.FileName != expectedFileName {
		t.Errorf("Expected fileMgr.FileName== '%v', instead fileMgr.FileName== '%v'",expectedFileName, fileMgr.FileName)
	}

	if fileMgr.FileExt != expectedExt {
		t.Errorf("Expected fileMgr.FileExt== '%v', instead got: fileMgr.FileExt=='%v'", expectedExt, fileMgr.FileName)
	}

	if !fileMgr.DMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated=='%v'", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameDoesExist {
		t.Errorf("Expected 'fileMgr.AbsolutePathFileNameDoesExist==true', instead got: fileMgr.AbsolutePathFileNameDoesExist=='%v'", fileMgr.AbsolutePathFileNameDoesExist)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.DMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.ActualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.ActualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.ActualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.ActualFileInfo.Name()=='%v'.  Instead, fileMgr.ActualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.ActualFileInfo.Name())
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh:= FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	if expectedAbsPathFileNameExt != fMgr.AbsolutePathFileName {
		t.Errorf("Expected AbsolutePathFileName='%v'.  Instead, AbsolutePathFileName='%v'",expectedAbsPathFileNameExt, fMgr.AbsolutePathFileName)
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_02(t *testing.T) {

	rawFileNameExt := "./newerFileForTest_01.txt"
	expectedFileNameExt := "newerFileForTest_01.txt"

	fh:= FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	absolutePath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'",adjustedPath, err.Error())
	}

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if absolutePath != dMgr.AbsolutePath {
		t.Errorf("Expected dMgr.AbsolutePath='%v'.  Instead, dMgr.AbsolutePath='%v'", absolutePath, dMgr.AbsolutePath)
	}

	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	if expectedAbsPathFileNameExt != fMgr.AbsolutePathFileName {
		t.Errorf("Expected AbsolutePathFileName='%v'.  Instead, AbsolutePathFileName='%v'",expectedAbsPathFileNameExt, fMgr.AbsolutePathFileName)
	}

}

func TestFileMgr_NewFromDirStrFileNameStr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName:= "newerFileForTest_01"
	expectedExt := ".txt"

	fh:= FileHelper{}
	rawPath := "../filesfortest/newfilesfortest"
	expectedPath := fh.AdjustPathSlash(rawPath)
	expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
	}

	fileMgr, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt). rawPath='%v' expectedFileNameExt='%v'  Error='%v'", rawPath, expectedFileNameExt, err.Error())
	}


	if fileMgr.FileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.FileNameExt == %v, instead fileMgr.FileNameExt='%v' ", expectedFileNameExt, fileMgr.FileNameExt)
	}

	if fileMgr.FileName != expectedFileName {
		t.Errorf("Expected fileMgr.FileName== '%v', instead fileMgr.FileName== '%v'",expectedFileName, fileMgr.FileName)
	}

	if fileMgr.FileExt != expectedExt {
		t.Errorf("Expected fileMgr.FileExt== '%v', instead got: fileMgr.FileExt=='%v'", expectedExt, fileMgr.FileName)
	}

	if !fileMgr.DMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated=='%v'", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameDoesExist {
		t.Errorf("Expected 'fileMgr.AbsolutePathFileNameDoesExist==true', instead got: fileMgr.AbsolutePathFileNameDoesExist=='%v'", fileMgr.DMgr.PathIsPopulated)
	}

	if !fileMgr.AbsolutePathFileNameIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathFileNameIsPopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.DMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.ActualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.ActualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.ActualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.ActualFileInfo.Name()=='%v'.  Instead, fileMgr.ActualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.ActualFileInfo.Name())
	}

	if expectedAbsPath !=  fileMgr.DMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath='%v'.  Instead, AbsolutePath='%v' ", expectedAbsPath, fileMgr.DMgr.AbsolutePath )
	}

	if expectedPath != fileMgr.DMgr.Path {
		t.Errorf("Expected Path='%v'.  Instead, Path='%v' ", expectedPath, fileMgr.DMgr.Path )
	}

}
