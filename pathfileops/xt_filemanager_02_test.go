package pathfileops

import (
	"fmt"
	"os"
	"testing"
)

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

	srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	dMgr, err := DirMgr{}.New(destDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). destDir='%v' Error='%v'", destDir, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr). dMgr.path='%v'  Error='%v'", dMgr.path, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
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

	srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

}

func TestFileMgr_New_01(t *testing.T) {

	fh := FileHelper{}

	relPath := "..\\logTest\\CmdrX\\CmdrX.log"
	commonDir, err := fh.MakeAbsolutePath(relPath)

	if err != nil {
		t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
			"relPath='%v'  Error='%v'", relPath, err.Error())
	}

	fileName := "CmdrX"
	fileNameExt := "CmdrX.log"
	extName := ".log"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
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

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("..\\logTest\\CmdrX\\CmdrX.log")

	fileName := "CmdrX"
	fileNameExt := "CmdrX.log"
	extName := ".log"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
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

func TestFileMgr_NewFromPathFileNameExtStr_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")
	fileName := "xt_dirmgr_01_test"
	fileNameExt := "xt_dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

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

func TestFileMgr_NewFromPathFileNameExtStr_02(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
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

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated==%v",
			fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist==%v", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
	}

}

func TestFileMgr_NewFromPathFileNameExtStr_03(t *testing.T) {

	path := "filehelperexamples"

	_, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err == nil {
		t.Error("Expected an error from FileMgr{}.NewFromPathFileNameExtStr(path) " +
			"because path='filehelperexamples'. However, NO ERROR WAS RETURNED! ")
	}
}

func TestFileMgr_NewFromPathFileNameExtStr_04(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
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

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.dMgr.isPathPopulated==true', instead got: "+
			"fileMgr.isPathPopulated==%v",
			fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist== %v", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Errorf("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, "+
			"fileMgr.isAbsolutePathFileNamePopulated == '%v' ", fileMgr.isAbsolutePathFileNamePopulated)
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Errorf("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is '%v'",
			fileMgr.dMgr.doesAbsolutePathExist)
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

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
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
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
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
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if absolutePath != dMgr.absolutePath {
		t.Errorf("Expected dMgr.absolutePath='%v'.  Instead, dMgr.absolutePath='%v'", absolutePath, dMgr.absolutePath)
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

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.actualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
	}

	if expectedAbsPath != fileMgr.dMgr.absolutePath {
		t.Errorf("Expected absolutePath='%v'.  Instead, absolutePath='%v' ", expectedAbsPath, fileMgr.dMgr.absolutePath)
	}

	if expectedPath != fileMgr.dMgr.path {
		t.Errorf("Expected path='%v'.  Instead, path='%v' ", expectedPath, fileMgr.dMgr.path)
	}

}
