package pathfileops

import (
	"fmt"
	"testing"
)

func TestFileMgr_CopyOut_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
	fileName := "xt_dirmgr_01_test"
	fileNameExt := "xt_dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

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

	if fMgr2.dMgr.path != expectedDir {
		t.Error(fmt.Sprintf("Expected CopyToThis to return path == '%v', instead got: ", expectedDir), fMgr2.dMgr.path)
	}

	result := fMgr2.Equal(&fileMgr)

	if result != true {
		t.Error("Expected Equal to return 'true' for fMgr2==fileMgr, instead got: ", result)
	}

}

func TestFileMgr_CopyFileMgrByIo_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
			"adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
			"dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). "+
			"destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'",
			destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	err = srcFMgr.CopyFileMgrByIo(&destFMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgrByIo(&destFMgr). "+
			"srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'",
			srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

	if !destFMgr.doesAbsolutePathFileNameExist {
		t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
	}

	err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) "+
			"destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
	}

	if fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. "+
			"Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

}

func TestFileMgr_CopyFileMgrByIo_02(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByIo(&destFMgr)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByIo_03(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	destFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByIo(&destFMgr)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileMgrByIo(&destFMgr) because " +
			"destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByIo_04(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	err = srcFMgr.CopyFileMgrByIo(nil)

	if err == nil {
		t.Error("Expected error return from CopyFileMgrByIo(nil) because " +
			"nil was passed to method. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByIoByLink_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgrByIoByLink(&destFMgr). srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'", srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=true. Instead it was 'false' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

	if !destFMgr.doesAbsolutePathFileNameExist {
		t.Error("Expected destFMgr.doesAbsolutePathFileNameExist='true'.  ERROR  destFMgr.doesAbsolutePathFileNameExist='false'")
	}

	err = fh.DeleteDirFile(destFMgr.absolutePathFileName)

	if err != nil {
		t.Errorf("Error returned from fh.DeleteDirFile(destFMgr.absolutePathFileName) "+
			"destFMgr.absolutePathFileName='%v' Error='%v'", destFMgr.absolutePathFileName, err.Error())
	}

	if fh.DoesFileExist(destFMgr.absolutePathFileName) {
		t.Errorf("Expected fh.DoesFileExist(destFMgr.absolutePathFileName)=false. "+
			"Instead it was 'true' destFMgr.absolutePathFileName='%v'", destFMgr.absolutePathFileName)
	}

}

func TestFileMgr_CopyFileMgrByIoByLink_02(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgrByIoByLink(&destFMgr). srcFMgr.absolutePathFileName='%v'  destFMgr.absolutePathFileName='%v'  Error='%v'", srcFMgr.absolutePathFileName, destFMgr.absolutePathFileName, err.Error())
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

func TestFileMgr_CopyFileMgrByIoByLink_03(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

	if err == nil {
		t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByIoByLink_04(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	rawDestPath := "../checkfiles/checkfiles02"

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDestPath). rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from  FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, expectedFileNameExt). destDMgr.absolutePath='%v'  expectedFileNameExt='%v'   Error='%v'", destDMgr.absolutePath, expectedFileNameExt, err.Error())
	}

	destFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByIoByLink(&destFMgr)

	if err == nil {
		t.Error("Expected error return from  srcFMgr.CopyFileMgrByIoByLink(&destFMgr) because " +
			"destFMgr.isInitialized = false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByLink_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgrByLink(&newFileMgr). "+
			"newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
	}

	doesFileExist, err = newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
		return
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileMgrByLink_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

	}
}

func TestFileMgr_CopyFileMgrByLink_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	newFileMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByLink(&newFileMgr)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(&newFileMgr) " +
			"because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")

	}
}

func TestFileMgr_CopyFileMgrByLink_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	err = srcFMgr.CopyFileMgrByLink(nil)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileMgrByLink(nil) " +
			"because nil was passed to this method. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileMgrByLinkByIo_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr). "+
			"newFileMgr='%v'  Error='%v'", newFileMgr.GetAbsolutePath(), err.Error())
	}

	doesFileExist, err = newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
		return
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileMgrByLinkByIo_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByLinkByIo_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	newFileMgr.isInitialized = false

	err = srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(&newFileMgr) " +
			"because newFileMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileMgrByLinkByIo_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	err = srcFMgr.CopyFileMgrByLinkByIo(nil)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileMgrByLinkByIo(nil) " +
			"because nil was passed to the method. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByIo_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

	if err != nil {
		t.Errorf("Error returned by srcFMgr.CopyFileStrByIo(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if !doesFileExist {
		t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
			"Destination File='%v'",
			rawAbsDestPath)
		return
	}

	err = fh.DeleteDirFile(rawAbsDestPath)

	if err != nil {
		t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
			"Destination File='%v' ", rawAbsDestPath)
	}

}

func TestFileMgr_CopyFileStrByIo_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByIo_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	err = srcFMgr.CopyFileStrByIo("")

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
			"because destination file path is empty string. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByIo_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist := fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByIo(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIo(rawAbsDestPath) " +
			"because source file does NOT exist. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByIoByLink_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

	if err != nil {
		t.Errorf("Error returned by srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if !doesFileExist {
		t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
			"Destination File='%v'",
			rawAbsDestPath)
		return
	}

	err = fh.DeleteDirFile(rawAbsDestPath)

	if err != nil {
		t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
			"Destination File='%v' ", rawAbsDestPath)
	}

}

func TestFileMgr_CopyFileStrByIoByLink_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileStrByIoByLink_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	err = srcFMgr.CopyFileStrByIoByLink("")

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
			"because destination file path is empty string. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileStrByIoByLink_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist := fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByIoByLink(rawAbsDestPath) " +
			"because source file does NOT exist. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileStrByLink_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

	if err != nil {
		t.Errorf("Error returned by srcFMgr.CopyFileStrByLink(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if !doesFileExist {
		t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
			"Destination File='%v'",
			rawAbsDestPath)
		return
	}

	err = fh.DeleteDirFile(rawAbsDestPath)

	if err != nil {
		t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
			"Destination File='%v' ", rawAbsDestPath)
	}

}

func TestFileMgr_CopyFileStrByLink_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByLink_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	err = srcFMgr.CopyFileStrByLink("")

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
			"because destination file path is empty string. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByLink_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist := fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByLink(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLink(rawAbsDestPath) " +
			"because source file does NOT exist. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByLinkByIo_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

	if err != nil {
		t.Errorf("Error returned by srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if !doesFileExist {
		t.Errorf("Error: Copy Operation FAILED! Destination File DOES NOT EXIST! "+
			"Destination File='%v'",
			rawAbsDestPath)
		return
	}

	err = fh.DeleteDirFile(rawAbsDestPath)

	if err != nil {
		t.Errorf("ERROR: Failed To Delete Test Destination File after copy operation! "+
			"Destination File='%v' ", rawAbsDestPath)
	}

}

func TestFileMgr_CopyFileStrByLinkByIo_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist = fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
			"because srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileStrByLinkByIo_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	doesFileExist, err := srcFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.DoesThisFileExist(). "+
			"Error='%v' ", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: Source Test File DOES NOT EXIST! File='%v'",
			srcFMgr.GetAbsolutePathFileName())
		return
	}

	err = srcFMgr.CopyFileStrByLinkByIo("")

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
			"because destination file path is empty string. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileStrByLinkByIo_04(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"

	fh := FileHelper{}

	absoluteSourceFile, err := fh.MakeAbsolutePath(sourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawRelDestPath := "../checkfiles/checkfiles02/TexLax201521.txt"

	rawAbsDestPath, err := fh.MakeAbsolutePath(rawRelDestPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(rawRelDestPath). "+
			"rawRelDestPath='%v' Error='%v' ", rawRelDestPath, err.Error())
	}

	doesFileExist := fh.DoesFileExist(rawAbsDestPath)

	if doesFileExist {
		err = fh.DeleteDirFile(rawAbsDestPath)
		t.Errorf("Error returned by fh.DeleteDirFile(rawAbsDestPath). "+
			"rawAbsDestPath='%v' Error='%v' ", rawAbsDestPath, err.Error())

	}

	err = srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath)

	if err == nil {
		t.Error("Expected error return from srcFMgr.CopyFileStrByLinkByIo(rawAbsDestPath) " +
			"because source file does NOT exist. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_CopyFileToDirByIoByLink_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Errorr returned by newFileMgr.DeleteThisFile(). Error='%v' ",
				newFileMgr.GetAbsolutePathFileName())
		}
	}

	err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileToDirByIoByLink(destDMgr). "+
			"destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
	}

	doesFileExist, err = newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !doesFileExist {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
		return
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileToDirByIoByLink_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

	if err == nil {
		t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByIoByLink_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	destDMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByIoByLink(destDMgr)

	if err == nil {
		t.Error("Expected an error from srcFMgr.CopyFileToDirByIoByLink(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByLinkByIo_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileToDirByLinkByIo(destDMgr). "+
			"destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
	}

	fileExists, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !fileExists {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileToDirByLinkByIo_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

	if err == nil {
		t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr) because " +
			"srcFMgr.isInitialized==false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByLinkByIo_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	destDMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByLinkByIo(destDMgr)

	if err == nil {
		t.Error("Expected an error from srcFMgr.CopyFileToDirByLinkByIo(destDMgr) because " +
			"destDMgr.isInitialized==false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByIo_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	err = srcFMgr.CopyFileToDirByIo(destDMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileToDirByIo(destDMgr). "+
			"destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
	}

	fileExists, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !fileExists {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileToDirByIo_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByIo(destDMgr)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByIo_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	destDMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByIo(destDMgr)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileToDirByIo(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByLink_01(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	err = srcFMgr.CopyFileToDirByLink(destDMgr)

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CopyFileToDirByLink(destDMgr). "+
			"destPath='%v'  Error='%v'", destDMgr.GetAbsolutePath(), err.Error())
	}

	fileExists, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from newFileMgr.DoesThisFileExist(). "+
			"Error='%v'", err.Error())
	}

	if !fileExists {
		t.Errorf("Error: File Copy Failed! Src File='%v' Dest File='%v'  ",
			srcFMgr.GetAbsolutePathFileName(), newFileMgr.GetAbsolutePathFileName())
	}

	err = newFileMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Expected that newly copied file would be deleted. "+
			"Instead, it was NOT deleted! NewFile := '%v' ", newFileMgr.GetAbsolutePathFileName())
	}

}

func TestFileMgr_CopyFileToDirByLink_02(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	srcFMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByLink(destDMgr)

	if err == nil {
		t.Error("Expected an error return from srcFMgr.CopyFileToDirByLink(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CopyFileToDirByLink_03(t *testing.T) {

	sourceFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedSourceFile := fh.AdjustPathSlash(sourceFile)
	absoluteSourceFile, err := fh.MakeAbsolutePath(adjustedSourceFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedSourceFile). "+
			"Error='%v' ", err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absoluteSourceFile)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(absoluteSourceFile). "+
			"Error='%v' ", err.Error())
	}

	rawDestPath := fh.AdjustPathSlash("../checkfiles/checkfiles02")

	destDMgr, err := DirMgr{}.New(rawDestPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDestPath). "+
			"rawDestPath='%v'  Error='%v'", rawDestPath, err.Error())
	}

	newFileMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(destDMgr, "newerFileForTest_01.txt")

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	doesFileExist, err := newFileMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(). "+
			"Error='%v'", err.Error())
	}

	if doesFileExist {

		err = newFileMgr.DeleteThisFile()

		if err != nil {
			t.Errorf("Error returned from newFileMgr.DeleteThisFile(). "+
				"Error='%v'", err.Error())
		}

	}

	destDMgr.isInitialized = false

	err = srcFMgr.CopyFileToDirByLink(destDMgr)

	if err == nil {
		t.Error("Expected an error return from destDMgr.CopyFileToDirByLink(destDMgr) because " +
			"srcFMgr.isInitialized == false. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_CreateDirAndFile_01(t *testing.T) {
	fh := FileHelper{}
	testFile := fh.AdjustPathSlash("../createFilesTest/Level01/Level02/Level03/TestFile011.txt")
	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(testFile)

	if err != nil {
		t.Errorf("Error thrown on FileHelper:GetPathFileNameElements():'%v'", err.Error())
	}

	if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {

		err = fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath)

		if err != nil {
			t.Errorf("Error thrown on fh.DeleteDirPathAll(fileMgr.dMgr.absolutePath). "+
				" fileMgr.dMgr.absolutePath='%v'   Error='%v' ", fileMgr.dMgr.absolutePath, err.Error())
		}

	}

	if fh.DoesFileExist(fileMgr.dMgr.absolutePath) {
		t.Errorf(fmt.Sprintf("Error: Failed to delete existing path '%v'",
			fileMgr.dMgr.absolutePath))
	}

	err = fileMgr.CreateDirAndFile()

	if err != nil {
		t.Errorf("Failed to Create Directory and File '%v', received Error:'%v'",
			fileMgr.absolutePathFileName, err.Error())
	}

	if !fh.DoesFileExist(fileMgr.absolutePathFileName) {
		t.Errorf(fmt.Sprintf("File Verfication failed file '%v' DOES NOT EXIST", fileMgr.absolutePathFileName))
	}

	s := "Created by File:'filemgr_test.go' Test Method: TestFileHelper_CreateDirAndFile()"

	_, err = fileMgr.WriteStrToFile(s)

	if err != nil {
		t.Errorf("Received error from fileMgr.WriteStrToFile(s). s='%v'  Error='%v' ", s, err.Error())
	}

	err = fileMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Received error from fileMgr.CloseThisFile(). fileMgr.absolutePathFileName='%v'  Error='%v' ", fileMgr.absolutePathFileName, err.Error())
	}

	err = fileMgr.dMgr.DeleteAll()

	if err != nil {
		t.Errorf("Error returned by fileMgr.dMgr.DeleteAll(). Attempted Deletion of %v. Error='%v'", fileMgr.absolutePathFileName, err.Error())
	}

}
