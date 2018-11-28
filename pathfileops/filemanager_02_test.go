package common

import (
	"testing"
	"io/ioutil"
	"os"
)


func TestFileMgr_OpenThisFileReadOnly_01(t *testing.T) {
	fh := FileHelper{}

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

	fMgr, err := FileMgr{}.New(filePath)

	if err != nil {
		t.Errorf("Error returned from common.FileMgr{}.New(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	err = fMgr.OpenThisFileReadOnly()

	if err != nil {
		t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	defer fMgr.CloseFile()

	b, err := ioutil.ReadAll(fMgr.FilePtr)

	if err != nil {
		t.Errorf("Error returned from ioutil.ReadAll(fMgr.FilePtr) filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	actualStr := string(b)

	expectedStr := "Test Read File. Do NOT alter the contents of this file."

	if expectedStr != actualStr {
		t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
	}

}

func TestFileMgr_OpenThisFileReadWrite_01(t *testing.T) {
	fh := FileHelper{}

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

	fMgr, err := FileMgr{}.New(filePath)

	if err != nil {
		t.Errorf("Error returned from common.FileMgr{}.New(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	err = fMgr.OpenThisFileReadWrite()

	if err != nil {
		t.Errorf("Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	defer fMgr.CloseFile()

	b, err := ioutil.ReadAll(fMgr.FilePtr)

	if err != nil {
		t.Errorf("Error returned from ioutil.ReadAll(fMgr.FilePtr) filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	actualStr := string(b)

	expectedStr := "Test Read File. Do NOT alter the contents of this file."

	if expectedStr != actualStr {
		t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
	}

}

func TestFileMgr_ReadFileBytes_01(t *testing.T) {

	fh := FileHelper{}

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

	fMgr, err := FileMgr{}.New(filePath)

	if err != nil {
		t.Errorf("Error returned from common.FileMgr{}.New(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	byteBuff := make([]byte, 2048, 2048)

	bytesRead, err := fMgr.ReadFileBytes(byteBuff)

	if err != nil {
		t.Errorf("Error returned from fMgr.ReadFileBytes(byteBuff). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	defer fMgr.CloseFile()

	var rStr = make([]rune,0,2048)

	for i:=0; i < len(byteBuff); i++ {

		if byteBuff[i] == 0 {
			break
		}

		rStr = append(rStr, rune(byteBuff[i]))

	}

	expectedStr := "Test Read File. Do NOT alter the contents of this file."
	actualStr := string(rStr)

	if expectedStr != actualStr {
		t.Errorf("Expected Read String='%v'. Instead, Actual Read String='%v'", expectedStr, actualStr)
	}

	expectedBytesRead := len(expectedStr)

	if expectedBytesRead != bytesRead {
		t.Errorf("Expected Bytes Read='%v'.  Instead, Actual Bytes Read='%v'", expectedBytesRead, bytesRead)
	}

}

func TestFileMgr_SetFileInfo(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

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

	fmgr, err := FileMgr{}.New(absPathFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
	}

	err = fmgr.SetFileInfo(info)

	if err != nil {
		t.Errorf("Error returned by fmgr.SetFileInfo(info). info.Name()='%v'  Error='%v'", info.Name(), err.Error())
	}

	if ! fmgr.ActualFileInfo.IsFInfoInitialized {
		t.Error("Error - File Manager FileInfoPlus object is not initialized!")
	}


	if fmgr.ActualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Error = Expected fmgr.ActualFileInfo.Name()='%v'.  Instead, fmgr.ActualFileInfo.Name()='%v'", expectedFileNameExt, fmgr.ActualFileInfo.Name())
	}

}

func TestFileMgr_WriteStrToFile_01(t *testing.T) {

	fh := FileHelper{}

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testWrite2998.txt")

	fMgr, err := FileMgr{}.New(filePath)

	if err != nil {
		t.Errorf("Error returned from common.FileMgr{}.New(filePath). filePathName='%v'  Error='%v'", filePath, err.Error())
	}

	expectedStr := "Test Write File. Do NOT alter the contents of this file."

	lExpectedStr := len(expectedStr)

	bytesWritten, err := fMgr.WriteStrToFile(expectedStr)

	if err != nil {
		t.Errorf("Error returned from fMgr.WriteStrToFile(expectedStr)  expectedStr='%v'  Error='%v'", expectedStr, err.Error())
	}

	err = fMgr.CloseFile()

	if err!=nil {
		t.Errorf("Error returned from fMgr.CloseFile() No 1.  Error='%v'", err.Error())
	}

	bytesRead, err := fMgr.ReadAllFile()

	if err != nil {
		t.Errorf("Error returned from fMgr.ReadAllFile(). filePathName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	if lExpectedStr != bytesWritten {
		t.Errorf("Error: Length of string written NOT equal to Bytes Read! Length of written string='%v'. Actual Bytes Read='%v' ", lExpectedStr, bytesWritten)
	}

	actualStr := string(bytesRead)

	if lExpectedStr != len(actualStr) {
		t.Errorf("Error: Length of actual string read is NOT equal to length of string written. lExpectedStr='%v'  len(actualStr)='%v'", lExpectedStr, len(actualStr))
	}

	if expectedStr != actualStr {
		t.Errorf("Error: expectedStr written='%v'  Actual string read='%v'", expectedStr, actualStr)
	}

	err = fMgr.CloseFile()

	if err!=nil {
		t.Errorf("Error returned by fMgr.CloseFile() No 2. Error='%v'", err.Error())
	}

	err = fMgr.DeleteThisFile()

	if err!=nil {
		t.Errorf("Error returned from fMgr.DeleteThisFile(). Error='%v'", err.Error())
	}

	doesFileExist:= fh.DoesFileExist(filePath)

	if doesFileExist {
		t.Errorf("Error: Failed to DELETE FileNameExt='%v'", fMgr.AbsolutePathFileName)
	}

}
