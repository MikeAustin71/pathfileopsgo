package pathfileops

import (
	appLib "MikeAustin71/pathfileopsgo/appLibs"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"
)

func TestFileHelper_IsAbsolutePath(t *testing.T) {

	fh := FileHelper{}
	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	result := fh.IsAbsolutePath(commonDir)

	if result == true {
		t.Error("IsAbsolutePath result is INVALID. Relative path classified as Absolute Path!")
	}

}

func TestFileHelper_JoinPathsAdjustSeparators_01(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
	file1 := "xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}
}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_02(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/")
	file1 := "/xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_03(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
	file1 := "/xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_04(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common")
	file1 := "xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_05(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common\\")
	file1 := "xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_06(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common//")
	file1 := "//xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")

	result1 := fh.JoinPathsAdjustSeparators(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinMismatchedPathsAdjustSeparators_07(t *testing.T) {
	fh := FileHelper{}
	path1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/")
	path12, err := fh.GetAbsPathFromFilePath(path1)

	if err != nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(path1) path1='%v'  Error='%v'", path1, err.Error())
	}

	file1 := "//xt_dirmgr_01_test.go"
	expected1 := fh.AdjustPathSlash("../../../pathfilego/003_filehelper/common/xt_dirmgr_01_test.go")
	expected12, err := fh.GetAbsPathFromFilePath(expected1)

	if err != nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expected1) expected1='%v'  Error='%v'", expected1, err.Error())
	}

	result1 := fh.JoinPathsAdjustSeparators(path12, file1)

	if result1 != expected12 {
		t.Errorf("Joined path and file name. Expected result '%v'. Instead result='%v'",
			expected12, result1)
	}

}

func TestFileHelper_JoinPaths_03(t *testing.T) {
	fh := FileHelper{}
	path1 := "../../../pathfilego/003_filehelper/common"
	file1 := "xt_dirmgr_01_test.go"
	expected1 := "..\\..\\..\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go"

	result1 := fh.JoinPaths(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_JoinBadPaths_04(t *testing.T) {
	fh := FileHelper{}
	path1 := "../../../pathfilego/003_filehelper/common/"
	file1 := "./xt_dirmgr_01_test.go"
	expected1 := "..\\..\\..\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go"

	result1 := fh.JoinPaths(path1, file1)

	if result1 != expected1 {
		t.Error(fmt.Sprintf("Joined path and file name. Expected result '%v', instead got:", expected1), result1)
	}

}

func TestFileHelper_MoveFile_01(t *testing.T) {
	fh := FileHelper{}
	setupFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
	srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
	destFile := fh.AdjustPathSlash("..\\logTest\\TestFile004.txt")

	if fh.DoesFileExist(destFile) {
		err := fh.DeleteDirFile(destFile)

		if err != nil {
			t.Error(fmt.Sprintf("Error on DeleteDirFile() deleting destination file, '%v'. Error:", destFile), err)
		}

		if fh.DoesFileExist(destFile) {
			t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", destFile))
		}
	}

	err := fh.CopyToNewFile(setupFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to destination file '%v' does NOT Exist. Error='%v'", setupFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Error(fmt.Sprintf("Source File '%v' does NOT EXIST!!", srcFile))
	}

	copyByLink, err := fh.MoveFile(srcFile, destFile)

	if err != nil {
		t.Error(fmt.Sprintf("Error on FileHelper:MoveFile() moving src '%v' to destination '%v' ", srcFile, destFile), err)
	}

	if fh.DoesFileExist(srcFile) {
		t.Error(fmt.Sprintf("FileHelper:MoveFile() FAILED! Source File '%v' still exists!!", srcFile))
	}

	if !fh.DoesFileExist(destFile) {
		t.Error(fmt.Sprintf("FileHelper:MoveFile() FAILED! Destination File '%v' DOES NOT EXIST!", destFile))
	}

	if copyByLink != true {
		t.Error("Expected copyByLink = 'true'. Instead, copyByLink='false'")
	}
}

func TestFileHelper_OpenFile_01(t *testing.T) {
	fh := FileHelper{}
	target := fh.AdjustPathSlash(alogtopTest2Text)
	expected := "Top level test file # 2."
	f, err := fh.OpenFileForReading(target)

	if err != nil {
		t.Errorf("Failed to open file: '%v' , got error - '%v'", target, err.Error())
	}

	le := len(expected)
	bRead := make([]byte, le)
	_, err2 := io.ReadAtLeast(f, bRead, 10)

	if err2 != nil {
		t.Errorf("Error Reading Test File: %v. Error = '%v'", target, err.Error())
	}

	s := string(bRead)

	if expected != s {
		t.Errorf("Expected to read string: '%v'. Instead got, '%v'", expected, s)
	}

	_ = f.Close()
}

func createTargetDir() error {
	fh := FileHelper{}
	targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

	if err1 != nil {
		return err1
	}

	if !fh.DoesFileExist(targetDir) {
		err2 := fh.MakeDirAll(targetDir)

		if err2 != nil {
			return err2
		}
	}

	targetFile := fh.JoinPathsAdjustSeparators(targetDir, alogFile)

	if fh.DoesFileExist(targetFile) {
		err3 := fh.DeleteDirFile(targetFile)
		if err3 != nil {
			return err3
		}
	}

	f, err4 := fh.CreateFile(targetFile)

	if err4 != nil {
		return err4
	}

	nowTime := appLib.DateTimeUtility{}.GetDateTimeNanoSecText(time.Now().Local())

	_, err5 := f.WriteString("Sample Write - " + nowTime + "\n")

	if err5 != nil {
		_ = f.Close()
		return err5
	}

	_, err6 := f.WriteString("File Name: " + targetFile)

	if err6 != nil {
		_ = f.Close()
		return err6
	}

	_ = f.Close()
	return nil
}

func deleteTargetDir() error {
	fh := FileHelper{}
	targetDir, err1 := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogTestBottomDir))

	if err1 != nil {
		return err1
	}

	if fh.DoesFileExist(targetDir) {
		err2 := fh.DeleteDirPathAll(targetDir)

		if err2 != nil {
			return err2
		}

		if fh.DoesFileExist(targetDir) {
			return errors.New("File still exists:" + targetDir)
		}
	}

	return nil
}
