package pathfileops

import (
	"fmt"
	"os"
	"testing"
)

func TestFileMgr_CopyOut_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
	fileName := "xt_dirmgr_01_test"
	fileNameExt := "xt_dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	fMgr2 := fileMgr.CopyOut()

	if fMgr2.fileName != fileName {
		t.Error(fmt.Sprintf("Expected CopyToThis to return fileName == '%v', instead got: ", fileName), fMgr2.fileName)
	}

	if fMgr2.fileExt != extName {
		t.Error(fmt.Sprintf("Expected CopyToThis to return fileExt == '%v', instead got: ", extName), fMgr2.fileExt)
	}

	if fMgr2.fileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected CopyToThis to return fileNameExt == '%v', instead got: ", fileNameExt), fMgr2.fileNameExt)
	}

	if fMgr2.dMgr.Path != expectedDir {
		t.Error(fmt.Sprintf("Expected CopyToThis to return Path == '%v', instead got: ", expectedDir), fMgr2.dMgr.Path)
	}

	result := fMgr2.Equal(&fileMgr)

	if result != true {
		t.Error("Expected Equal to return 'true' for fMgr2==fileMgr, instead got: ", result)
	}

}

func TestFileMgr_CopyFileMgr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.AbsolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.AbsolutePath, adjustedPath, err.Error())
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
		t.Errorf("Error returned from srcFMgr.CopyFileMgr(&destFMgr). srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'", srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

	if !destFMgr.doesAbsolutePathFileNameExist {
		t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
	}

	err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
	}

	if fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

}

func TestFileMgr_CreateDirAndFile_01(t *testing.T) {
	fh := FileHelper{}
	testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
	fileMgr, err := FileMgr{}.New(testFile)

	if err != nil {
		t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
	}

	if fh.DoesFileExist(fileMgr.dMgr.AbsolutePath) {

		err = fh.DeleteDirPathAll(fileMgr.dMgr.AbsolutePath)

		if err != nil {
			t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.AbsolutePath). "+
				" fileMgr.dMgr.AbsolutePath='%v'   Error='%v' ", fileMgr.dMgr.AbsolutePath, err.Error())
		}

	}

	if fh.DoesFileExist(fileMgr.dMgr.AbsolutePath) {
		t.Errorf(fmt.Sprintf("Error: Failed to delete existing path '%v'", fileMgr.dMgr.AbsolutePath))
	}

	err = fileMgr.CreateDirAndFile()

	if err != nil {
		t.Errorf("Failed to Create Directory and File '%v', received Error:'%v'", fileMgr.absolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(fileMgr.absolutePathFileName) {
		t.Errorf(fmt.Sprintf("File Verfication failed file '%v' DOES NOT EXIST", fileMgr.absolutePathFileName))
	}

	s := "Created by File:'filemgr_test.go' Test Method: TestFileHelper_CreateDirAndFile()"

	_, err = fileMgr.WriteStrToFile(s)

	if err != nil {
		t.Errorf("Received error from fileMgr.WriteStrToFile(s). s='%v'  Error='%v' ", s, err.Error())
	}

	err = fileMgr.CloseFile()

	if err != nil {
		t.Errorf("Received error from fileMgr.CloseFile(). fileMgr.absolutePathFileName='%v'  Error='%v' ", fileMgr.absolutePathFileName, err.Error())
	}

	err = fileMgr.dMgr.DeleteAll()

	if err != nil {
		t.Errorf("Error returned by fileMgr.dMgr.DeleteAll(). Attempted Deletion of %v. Error='%v'", fileMgr.absolutePathFileName, err.Error())
	}

}

func TestFileMgr_Equal_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	fileMgr2 := fileMgr.CopyOut()

	// fh is now NOT Equal to fileMgr
	fileMgr2.dMgr.Path = ""

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

	err := fh.CopyFileByIo(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.New(srcFile)

	if err != nil {
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

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
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

	err := fh.CopyFileByIo(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.New(srcFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: New Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

}

func TestFileMgr_New_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")
	fileName := "xt_dirmgr_01_test"
	fileNameExt := "xt_dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	if fileMgr.fileName != fileName {
		t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
	}

	if fileMgr.fileExt != extName {
		t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
	}

	if fileMgr.fileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
	}

	if !fileMgr.isInitialized {
		t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
	}

	if !fileMgr.isFileNamePopulated {
		t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
	}

	if !fileMgr.isFileNameExtPopulated {
		t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
	}

	if !fileMgr.isFileExtPopulated {
		t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
	}

}

func TestFileMgr_New_02(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.New(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(path) "+
			"path== '%v' Error: %v", path, err)
	}

	if eFileNameExt != fileMgr.fileNameExt {
		t.Errorf("Expected extracted fileNameExt == %v, instead got: %v",
			eFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != "filehelperexamples" {
		t.Errorf("Expected fileMgr.fileName== 'filehelperexamples', "+
			"instead got: fileMgr.fileName== %v", fileMgr.fileName)
	}

	if fileMgr.fileExt != ".go" {
		t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
			fileMgr.fileExt)
	}

	if !fileMgr.dMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated==%v",
			fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist==%v", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

}

func TestFileMgr_New_03(t *testing.T) {

	path := "filehelperexamples"

	fileMgr, err := FileMgr{}.New(path)

	if err != nil {
		t.Errorf("Error returned from FileHelper:GetPathFileNameElements() path== '%v' Error: %v", path, err)
	}

	if fileMgr.fileName != path {
		t.Error("Expected fileMgr.fileName=='dirmgr_test', instead got:", fileMgr.fileName)
	}

	if fileMgr.dMgr.PathIsPopulated {
		t.Error("Expected fileMgr.PathIsPopulated==false, instead got:", fileMgr.dMgr.PathIsPopulated)
	}

	if fileMgr.isFileExtPopulated {
		t.Error("Expected fileMgr.isFileExtPopulated==false, instead got:", fileMgr.isFileExtPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated==true, instead got:", fileMgr.isAbsolutePathFileNamePopulated)
	}

	if fileMgr.dMgr.AbsolutePathIsPopulated {
		t.Error("Expected fileMgr.AbsolutePathIsPopulated==false, instead got:", fileMgr.dMgr.AbsolutePathIsPopulated)
	}
}

func TestFileMgr_New_04(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.New(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(path) "+
			"path=='%v' Error: %v ", path, err)
	}

	if eFileNameExt != fileMgr.fileNameExt {
		t.Errorf("Expected extracted fileNameExt == '%v', instead got: '%v' ",
			eFileNameExt, fileMgr.fileNameExt)
	}

	if "filehelperexamples" != fileMgr.fileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead got: fileMgr.fileName== '%v'",
			"filehelperexamples", fileMgr.fileName)
	}

	if ".go" != fileMgr.fileExt {
		t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
			fileMgr.fileExt)
	}

	if !fileMgr.dMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.dMgr.PathIsPopulated==true', instead got: "+
			"fileMgr.PathIsPopulated==%v",
			fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist== %v", fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Errorf("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, "+
			"fileMgr.isAbsolutePathFileNamePopulated == '%v' ", fileMgr.isAbsolutePathFileNamePopulated)
	}

	if !fileMgr.dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is '%v'",
			fileMgr.dMgr.AbsolutePathDoesExist)
	}

}

func TestFileMgr_NewFromFileInfo_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName := "newerFileForTest_01"
	expectedExt := ".txt"
	fh := FileHelper{}
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

	if fileMgr.fileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != expectedFileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
	}

	if fileMgr.fileExt != expectedExt {
		t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
	}

	if !fileMgr.dMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated=='%v'", fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.actualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
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

	if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
		t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_02(t *testing.T) {

	rawFileNameExt := "./newerFileForTest_01.txt"
	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	absolutePath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
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

	if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
		t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
	}

}

func TestFileMgr_NewFromDirStrFileNameStr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName := "newerFileForTest_01"
	expectedExt := ".txt"

	fh := FileHelper{}
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

	if fileMgr.fileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != expectedFileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
	}

	if fileMgr.fileExt != expectedExt {
		t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
	}

	if !fileMgr.dMgr.PathIsPopulated {
		t.Errorf("Expected 'fileMgr.PathIsPopulated==true', instead got: fileMgr.PathIsPopulated=='%v'", fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.dMgr.PathIsPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.AbsolutePathDoesExist {
		t.Error("Expected fileMgr.AbsolutePathDoesExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.actualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
	}

	if expectedAbsPath != fileMgr.dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath='%v'.  Instead, AbsolutePath='%v' ", expectedAbsPath, fileMgr.dMgr.AbsolutePath)
	}

	if expectedPath != fileMgr.dMgr.Path {
		t.Errorf("Expected Path='%v'.  Instead, Path='%v' ", expectedPath, fileMgr.dMgr.Path)
	}

}
