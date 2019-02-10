package appExamples

import (
	appLib "../appLibs"
	pathFileOp "../pathfileops"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	fp "path/filepath"
	"strings"
	"time"
)

func TestDirMgrWalDirDeleteFiles() {

	ePrefix := "TestDirMgrWalDirDeleteFiles()"

	origDir, err := MainDirMgrTestSetupFileWalkDeleteFiles()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'\n", err.Error())
		return
	}

	fh := pathFileOp.FileHelper{}

	if !fh.DoesFileExist(origDir) {
		fmt.Printf(ePrefix+
			"Error: The target directory does NOT Exist! origDir='%v'\n",
			origDir)
		return
	}

	dMgr, err := pathFileOp.DirMgr{}.New(origDir)

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). "+
			"origDir='%v' Error='%v'\n", origDir, err.Error())
		return
	}

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"

	searchPattern := "*.txt"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned from time.Parse(fmtstr, fOlderThanStr). "+
			"fOlderThanStr='%v' Error='%v'\n", fOlderThanStr, err.Error())
		return
	}

	filesNewerThan := time.Time{}

	fsc := pathFileOp.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOp.FileSelectMode.ORSelect()

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {

		fmt.Printf(ePrefix+
			"Expected to find 6-files deleted. Instead, %v-files were deleted.\n",
			dInfo.DeletedFiles.GetNumOfFileMgrs())
		return
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

	for i := 0; i < dInfo.DeletedFiles.GetNumOfFileMgrs(); i++ {

		fileMgr, err := dInfo.DeletedFiles.PeekFileMgrAtIndex(i)

		if err != nil {
			fmt.Printf(ePrefix+
				"Error returned by dInfo.DeletedFiles.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' \n",
				i, err.Error())
			return

		}

		fileNameExt := fileMgr.GetFileNameExt()

		if strings.Contains(fileNameExt, oldFile1) {
			oldFile1Found = true
		} else if strings.Contains(fileNameExt, oldFile2) {
			oldFile2Found = true
		} else if strings.Contains(fileNameExt, oldFile3) {
			oldFile3Found = true
		} else if strings.Contains(fileNameExt, newFile1) {
			newFile1Found = true
		} else if strings.Contains(fileNameExt, newFile2) {
			newFile2Found = true
		} else if strings.Contains(fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of oldFile1='%v'. The file was NOT deleted!\n", oldFile1)
	}

	if oldFile2Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of oldFile2='%v'. The file was NOT deleted!\n", oldFile2)
	}

	if oldFile3Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of oldFile3='%v'. The file was NOT deleted!\n", oldFile3)
	}

	if newFile1Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of newFile1='%v'. The file was NOT deleted!\n", newFile1)
	}

	if newFile2Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of newFile2='%v'. The file was NOT deleted!\n", newFile2)
	}

	if newFile3Found == false {
		fmt.Printf(ePrefix+
			"Expected deletion of newFile3='%v'. The file was NOT deleted!\n", newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		fmt.Printf(ePrefix+
			"Expected zero Error Returns. Instead number of Error Returns='%v'\n",
			len(dInfo.ErrReturns))
		return
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		fmt.Printf(ePrefix+
			"Expected 3-directories to be found. Instead, number of directories found='%v'\n",
			dInfo.Directories.GetNumOfDirs())
		return
	}

	fmt.Println("Success!")
	fmt.Println("origDir: ", origDir)

}

func GetBaseProjectPath() string {

	ePrefix := "getBaseProjectPath() "
	fh := pathFileOp.FileHelper{}
	currDir, err := fh.GetAbsCurrDir()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by fh.GetAbsCurrDir(). Error='%v' \n", err.Error())
		return "Error"
	}

	target := "pathfileopsgo"
	idx := strings.Index(currDir, target)

	if idx < 0 {
		fmt.Printf(ePrefix +
			"Error: Unable to locate \"pathfileopsgo\" in current directory string! \n")
		return "Error"
	}

	idx += len(target)

	baseDir := currDir[0:idx]

	return baseDir
}

func TestingDirMgrDeleteWalkDirFiles06() {

	ePrefix := "TestingDirMgrDeleteWalkDirFiles06() "

	// origDir = D:\gowork\src\MikeAustin71\pathfileopsgo\dirwalkdeletetests\dirdelete01
	origDir, err := TestingDirMgr02TestSetupFileWalkDeleteFiles()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). "+
			"Error='%v'", err.Error())
		return
	}

	fh := pathFileOp.FileHelper{}

	if !fh.DoesFileExist(origDir) {
		fmt.Printf(ePrefix+
			"Error: The target directory does NOT Exist! origDir='%v'",
			origDir)
		return
	}

	dMgr, err := pathFileOp.DirMgr{}.New(origDir)

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' "+
			"Error='%v'", origDir, err.Error())
		return
	}

	searchPattern := ""
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fNewerThanStr := "2016-07-01 00:00:00.000000000 -0500 CDT"
	//fOlderThanStr := "2018-11-30 00:00:00.000000000 -0600 CST"
	filesOlderThan := time.Now()
	filesOlderThan.AddDate(0, 1, 0)

	/*
		filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

		if err != nil {
			fmt.Printf("Error returned from time.Parse(fmtstr, fOlderThanStr). "+
				"fOlderThanStr='%v' Error='%v'", fOlderThanStr, err.Error())
			return
		}
	*/
	filesNewerThan, err := time.Parse(fmtstr, fNewerThanStr)

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned from time.Parse(fmtstr, fNewerThanStr). "+
			"fNewerThanStr='%v' Error='%v'", fNewerThanStr, err.Error())
	}

	fsc := pathFileOp.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOp.FileSelectMode.ANDSelect()

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetNumOfFileMgrs() != 6 {
		fmt.Printf(ePrefix+
			"Expected to find 6-files deleted. Instead, %v-files "+
			"were deleted.", dInfo.DeletedFiles.GetNumOfFileMgrs())
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

	for i := 0; i < dInfo.DeletedFiles.GetNumOfFileMgrs(); i++ {

		fileMgr, err := dInfo.DeletedFiles.PeekFileMgrAtIndex(i)

		if err != nil {
			fmt.Printf(ePrefix+
				"Error returned by dInfo.DeletedFiles.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		fileNameExt := fileMgr.GetFileNameExt()

		if strings.Contains(fileNameExt, oldFile1) {
			oldFile1Found = true
		} else if strings.Contains(fileNameExt, oldFile2) {
			oldFile2Found = true
		} else if strings.Contains(fileNameExt, oldFile3) {
			oldFile3Found = true
		} else if strings.Contains(fileNameExt, newFile1) {
			newFile1Found = true
		} else if strings.Contains(fileNameExt, newFile2) {
			newFile2Found = true
		} else if strings.Contains(fileNameExt, newFile3) {
			newFile3Found = true
		}

	}

	if oldFile1Found == false {
		fmt.Printf("Expected deletion of oldFile1='%v'. The file was NOT deleted!", oldFile1)
	}

	if oldFile2Found == false {
		fmt.Printf("Expected deletion of oldFile2='%v'. The file was NOT deleted!", oldFile2)
	}

	if oldFile3Found == false {
		fmt.Printf("Expected deletion of oldFile3='%v'. The file was NOT deleted!", oldFile3)
	}

	if newFile1Found == false {
		fmt.Printf("Expected deletion of newFile1='%v'. The file was NOT deleted!", newFile1)
	}

	if newFile2Found == false {
		fmt.Printf("Expected deletion of newFile2='%v'. The file was NOT deleted!", newFile2)
	}

	if newFile3Found == false {
		fmt.Printf("Expected deletion of newFile3='%v'. The file was NOT deleted!", newFile3)
	}

	if len(dInfo.ErrReturns) != 0 {
		fmt.Printf("Expected zero Error Returns. Instead number of Error Returns='%v'", len(dInfo.ErrReturns))
	}

	if dInfo.Directories.GetNumOfDirs() != 3 {
		fmt.Printf("Expected 3-directories to be found. Instead, number of directories found='%v'",
			dInfo.Directories.GetNumOfDirs())
	}

}

func TestingDirMgr02TestSetupFileWalkDeleteFiles() (string, error) {

	ePrefix := "TestFile: xt_dirmgr_02_test.go Func: DirMgr02TestSetupFileWalkDeleteFiles() "

	fh := pathFileOp.FileHelper{}

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
		return "", fmt.Errorf(ePrefix+
			"Error: Old Files Directory does NOT exist! "+
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
			"Error: NewFromPathFileNameExtStr Files Directory does NOT exist! dirNewFilesForTest='%v'",
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
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + newFile2

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + newFile3

	err = fh.CopyFileByLinkByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	return origDir, nil
}

func TestingfileinfoplusEqual01() {
	fh := pathFileOp.FileHelper{}

	baseFileName := "newerFileForTest_01.txt"

	baseDirPath := "../filesfortest/newfilesfortest"

	absCurrPath, err := fh.GetAbsCurrDir()

	if err != nil {
		fmt.Printf("Error returned by fh.GetAbsCurrDir(). %v\n", err.Error())
	}

	fmt.Println("Absolute Current path: ", absCurrPath)

	absBaseDirPath, err := fh.MakeAbsolutePath(baseDirPath)

	if err != nil {
		fmt.Printf("Error returned from fh.MakeAbsolutePath(baseDirPath). "+
			"baseDirPath='%v' Error='%v'", baseDirPath, err.Error())
		return
	}

	fmt.Printf("    Base path: %v\n", baseDirPath)
	fmt.Printf("Absolute path: %v\n", absBaseDirPath)
	fmt.Println()

	absPathFileName, _ := fh.AddPathSeparatorToEndOfPathStr(absBaseDirPath)
	absPathFileName = absPathFileName + baseFileName

	fInfo, err := fh.GetFileInfoFromPath(absPathFileName)

	if err != nil {
		fmt.Printf("Error returned from fh.GetFileInfoFromPath(absPathFileName). "+
			"absPathFileName='%v' Error='%v'", absPathFileName, err.Error())
		return
	}

	fip := pathFileOp.FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Name() != baseFileName {
		fmt.Printf("Expected fip.Name()='%v'. Instead, fip.Name()='%v'.",
			baseFileName, fip.Name())
		return
	}

	fip2 := pathFileOp.FileInfoPlus{}.NewFromFileInfo(fInfo)

	if fip.Equal(&fip2) == false {
		fmt.Println("Expected  fip to EQUAL fip2. It DID NOT!")
		fmt.Println("fip file info")
		PrintFileInfoPlusFields(fip)
		fmt.Println()
		fmt.Println("fip2 file info")
		PrintFileInfoPlusFields(fip2)
	}

}

func TestWriteFile() {
	fh := pathFileOp.FileHelper{}

	filePath := GetBaseProjectPath() +
		fh.AdjustPathSlash("/checkfiles/checkfiles03/testWrite2998.txt")

	fMgr, err := pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(filePath)

	if err != nil {
		fmt.Printf("Error returned from pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePathName='%v'  Error='%v'", filePath, err.Error())
		return
	}

	expectedStr := "Test Write File. Do NOT alter the contents of this file."

	lExpectedStr := len(expectedStr)

	err = fMgr.OpenThisFileReadWrite()

	if err != nil {
		fmt.Printf("Error returned from fMgr.OpenThisFileReadWrite(). Error='%v' \n", err.Error())
	}

	bytesWritten, err := fMgr.WriteStrToFile(expectedStr)

	if err != nil {
		fmt.Printf("Error returned from fMgr.WriteStrToFile(expectedStr)  expectedStr='%v'  Error='%v'\n", expectedStr, err.Error())
		return
	}

	fmt.Println("Expected String Length: ", lExpectedStr)
	fmt.Println("Bytes Written To File: ", bytesWritten)

	_ = fMgr.CloseFile()

	/*
		bytesRead, err := fMgr.ReadAllFile()

		if err != nil {
			fmt.Printf("Error returned from fMgr.ReadAllFile(). filePathName='%v'  Error='%v'", fMgr.GetAbsolutePathFileName(), err.Error())
			return
		}

		if lExpectedStr != bytesWritten {
			fmt.Printf("Error: Length of string written NOT equal to Bytes Read! Length of written string='%v'. Actual Bytes Read='%v' ", lExpectedStr, bytesWritten)
			return
		}

		actualStr := string(bytesRead)

		if lExpectedStr != len(actualStr) {
			fmt.Printf("Error: Legth of actual string read is NOT equal to length of string written. lExpectedStr='%v'  len(actualStr)='%v'", lExpectedStr, len(actualStr))
			return
		}

		if expectedStr != actualStr {
			fmt.Printf("Error: expectedStr written='%v'  Actual string read='%v'", expectedStr, actualStr)
			return
		}

		fMgr.CloseFile()

		doesFileExist := fh.DoesFileExist(filePath)

		if !doesFileExist {
			fmt.Printf("Error: After writing string, target file does NOT exist!. fileNameExt='%v'", fMgr.GetAbsolutePathFileName())
			return
		}

	*/

	err = fMgr.DeleteThisFile()

	if err != nil {
		fmt.Printf("Error returned from fMgr.DeleteThisFile(). Error='%v'", err.Error())
		return
	}

	doesFileExist, err := fMgr.DoesThisFileExist()

	// doesFileExist = fh.DoesFileExist(filePath)

	if doesFileExist {
		fmt.Printf("Error: Failed to DELETE fileNameExt='%v'", fMgr.GetAbsolutePathFileName())
		return
	}

	fmt.Println("Successful Completion!")
}

func TestOpenFile() {
	fh := pathFileOp.FileHelper{}
	ePrefix := "TestOpenFile() "

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

	fMgr, err := pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(filePath)

	if err != nil {
		fmt.Printf(ePrefix+"Error returned from pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	err = fMgr.OpenThisFileReadOnly()

	if err != nil {
		fmt.Printf(ePrefix+"Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	b, err := ioutil.ReadAll(fMgr.GetFilePtr())

	if err != nil {
		fmt.Printf(ePrefix+"Error returned from ioutil.ReadAll(fMgr.filePtr) filePath='%v'  Error='%v'", filePath, err.Error())
		_ = fMgr.CloseFile()
		return
	}

	str := string(b)

	err = fMgr.CloseFile()

	if err != nil {
		fmt.Printf(ePrefix+
			"%v", err.Error())
		return
	}

	fmt.Println("         Test Open File")
	fmt.Println("===================================")
	fmt.Println("Buffer:")
	fmt.Println(str)

	expectedStr := "Test Read File. Do NOT alter the contents of this file."

	if expectedStr == str {
		fmt.Println("Expected String == Actual String!")
	}

}

func TestDeleteDirectoryTree() {

	ePrefix := "TestDeleteDirectoryTree() "
	fh := pathFileOp.FileHelper{}
	substituteDir := GetBaseProjectPath() + fh.AdjustPathSlash("/testdestdir/destdir")

	substituteDMgr, err := pathFileOp.DirMgr{}.New(substituteDir)

	if err != nil {
		fmt.Printf(ePrefix+"Error returned by pathFileOp.DirMgr{}.NewFromPathFileNameExtStr(substituteDir). substituteDir='%v'  Error='%v'\n", substituteDir, err.Error())
		return
	}

	err = substituteDMgr.DeleteAll()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by substituteDMgr.DeleteAll(). substituteDMgr.path='%v'  "+
			"Error='%v'\n", substituteDMgr.GetPath(), err.Error())
		return
	}

	fmt.Println("Successful Completion")

}

func TestCopyDirectoryTree() {

	ePrefix := "TestCopyDirectoryTree() "
	fh := pathFileOp.FileHelper{}
	dir := GetBaseProjectPath() + fh.AdjustPathSlash("/testsrcdir")

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	dMgr, err := pathFileOp.DirMgr{}.New(dir)

	if err != nil {
		fmt.Printf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'\n", dir, err.Error())
		return
	}

	if !dMgr.DoesDirMgrAbsolutePathExist() {
		fmt.Printf("Expected target directory to exist. I does NOT exist. "+
			"dMgr.path='%v' dMgr.AbolutePath='%v'\n",
			dMgr.GetPath(), dMgr.GetAbsolutePath())
		return
	}

	fsc := pathFileOp.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOp.FileSelectMode.ANDSelect()

	dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		fmt.Printf("Error returned from dMgr.FindWalkDirFiles(searchPattern, filesOlderThan, filesNewerThan). dir='%v' Error='%v'\n", dir, err.Error())
		return
	}

	baseDir := GetBaseProjectPath() + fh.AdjustPathSlash("/testsrcdir")

	baseDMgr, err := pathFileOp.DirMgr{}.New(baseDir)

	if err != nil {
		fmt.Printf("Error returned by pathFileOp.DirMgr{}.NewFromPathFileNameExtStr(baseDir) baseDir='%v' Error='%v'", baseDir, err.Error())
		return
	}

	substituteDir := GetBaseProjectPath() + fh.AdjustPathSlash("/testdestdir/destdir")

	substituteDMgr, err := pathFileOp.DirMgr{}.New(substituteDir)

	if err != nil {
		fmt.Printf(ePrefix+"Error returned by pathFileOp.DirMgr{}.NewFromPathFileNameExtStr(substituteDir). substituteDir='%v'  Error='%v'", substituteDir, err.Error())
		return
	}

	newDirTree, err := dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr). Error='%v'", err.Error())
		return
	}

	if dirTreeInfo.Directories.GetNumOfDirs() != newDirTree.Directories.GetNumOfDirs() {

		fmt.Printf(ePrefix+"Error: Expected Number of Directories = '%v'.  "+
			"Instead, Number of NewFromPathFileNameExtStr Directories = '%v'",
			dirTreeInfo.Directories.GetNumOfDirs(), newDirTree.Directories.GetNumOfDirs())
		return
	}

	if dirTreeInfo.FoundFiles.GetNumOfFileMgrs() != newDirTree.FoundFiles.GetNumOfFileMgrs() {
		fmt.Printf(ePrefix+
			"Error: Expected Number of Files = '%v'.  Instead, actual Number of "+
			"NewFromPathFileNameExtStr Files = '%v'",
			dirTreeInfo.FoundFiles.GetNumOfFileMgrs(), newDirTree.FoundFiles.GetNumOfFileMgrs())
		return
	}

	for i := 0; i < newDirTree.FoundFiles.GetNumOfFileMgrs(); i++ {

		fileMgr, err := newDirTree.FoundFiles.PeekFileMgrAtIndex(i)

		if err != nil {
			fmt.Printf(ePrefix+
				"Error returned by newDirTree.FoundFiles.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v'", i, err.Error())
			return
		}

		doesFileExist, err := fileMgr.DoesThisFileExist()

		if err != nil {
			fmt.Printf(ePrefix+
				"Error returned by fileMgr.DoesThisFileExist(). "+
				"i='%v' fileNameExt='%v'  Error='%v'",
				i, fileMgr.GetFileNameExt(), err.Error())
			return
		}

		if !doesFileExist {
			fmt.Printf(ePrefix+
				"Error: Failed to create fileNameExt='%v'. "+
				"It does NOT exist in target directory.",
				fileMgr.GetFileNameExt())
			return
		}

	}

	fmt.Println("Successful Completion")
}

func TestMainCleanDirStr(rawPath string) {
	fh := pathFileOp.FileHelper{}

	dirPath, isEmpty, err := fh.CleanDirStr(rawPath)

	if err != nil {
		fmt.Printf("Error returned by fh.CleanDirStr(rawPath) rawPath='%v'  Error='%v' \n",
			rawPath, err.Error())
		return
	}

	fmt.Println("=========================================")
	fmt.Println("     Test Clean Directory String")
	fmt.Println("=========================================")
	fmt.Println()
	fmt.Println("Returned Dir path: ", dirPath)
	fmt.Println("          isEmpty: ", isEmpty)
	fmt.Println("  raw path string: ", rawPath)

}

func TestNewFileMgrFromPathFileNameStr(pathFileNameExt string) {

	fMgr, err := pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(pathFileNameExt)

	if err != nil {
		fmt.Printf("Error returned from pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(pathFileNameExt) pathFileNameExt='%v'  Error='%v' \n", pathFileNameExt, err.Error())
	}

	PrintFileManagerFields(fMgr)

}

func TestNewFileMgrFromDirMgrFileNameExt(rawPath, rawFileNameExt string) {

	fh := pathFileOp.FileHelper{}
	adjustedPath := fh.AdjustPathSlash(rawPath)

	dMgr, err := pathFileOp.DirMgr{}.New(adjustedPath)

	if err != nil {
		fmt.Printf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v' \n",
			adjustedPath, err.Error())
		return
	}

	fMgr, err := pathFileOp.FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

	if err != nil {
		fmt.Printf("Error returned by FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt). "+
			"dMgr.path='%v' rawFileNameExt='%v'  \n", dMgr.GetPath(), rawFileNameExt)
	}

	PrintFileManagerFields(fMgr)

}

func TestDirMgr(rawPath string, expectedPath string) {

	fh := pathFileOp.FileHelper{}

	expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

	if err != nil {
		fmt.Printf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'\n", expectedPath, err.Error())
		return
	}

	dMgr, err := pathFileOp.DirMgr{}.New(rawPath)

	if err != nil {
		fmt.Printf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawPath) rawPath='%v' Error='%v'", rawPath, err.Error())
		return
	}

	PrintDirMgrFields(dMgr)

	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("             Expected Results")
	fmt.Println("=========================================")
	fmt.Println("          rawPath: ", rawPath)
	fmt.Println("     expectedPath: ", expectedPath)
	fmt.Println("   expectedAbsDir: ", expectedAbsDir)

}

func TestMainCleanFileNameExt(rawFileNameExt, expectedFileNameExt string) {

	fh := pathFileOp.FileHelper{}
	adjustedFileNameExt := fh.AdjustPathSlash(rawFileNameExt)
	processedFileNameExt, isFileNameEmpty, err := fh.CleanFileNameExtStr(adjustedFileNameExt)

	if err != nil {
		fmt.Printf("Error returned by fh.CleanFileNameExtStr(adjustedFileNameExt) adjustedFileNameExt='%v' Error='%v'\n", adjustedFileNameExt, err.Error())
		return
	}

	fmt.Println("=========================================")
	fmt.Println("          Clean fileNameExt Tests")
	fmt.Println("=========================================")
	fmt.Println("      rawFileNameExt: ", rawFileNameExt)
	fmt.Println(" adjustedFileNameExt: ", adjustedFileNameExt)
	fmt.Println("     isFileNameEmpty: ", isFileNameEmpty)
	fmt.Println("processedFileNameExt: ", processedFileNameExt)
	fmt.Println(" expectedFileNameExt: ", expectedFileNameExt)

}

func TestMainCleanPath(rawPath string) {

	adjustedPath := fp.FromSlash(rawPath)

	fpCleanedPath := fp.Clean(adjustedPath)

	pathCleanedPath := path.Clean(adjustedPath)

	adjustedAbsolutePath, err := fp.Abs(adjustedPath)

	if err != nil {
		fmt.Printf("Error returned from fp.Abs(adjustedPath) adjustedPath='%v' Error='%v' \n", adjustedPath, err.Error())
		return
	}

	fpCleanedAbsolutePath := fp.Clean(adjustedAbsolutePath)

	pathCleanedAbsolutePath := path.Clean(adjustedAbsolutePath)

	fmt.Println("=========================================")
	fmt.Println("           Clean path Tests")
	fmt.Println("=========================================")
	fmt.Println("                rawPath: ", rawPath)
	fmt.Println("           adjustedPath: ", adjustedPath)
	fmt.Println("  fp clean adjustedPath: ", fpCleanedPath)
	fmt.Println("path clean adjustedPath: ", pathCleanedPath)
	fmt.Println("------------------------------------------")
	fmt.Println("   adjustedAbsolutePath: ", adjustedAbsolutePath)
	fmt.Println("  fpCleanedAbsolutePath: ", fpCleanedAbsolutePath)
	fmt.Println("pathCleanedAbsolutePath: ", pathCleanedAbsolutePath)

}

func TestMainGetPathFromPathFileName098(dir, expectedDir string) {

	fh := pathFileOp.FileHelper{}

	result, isEmpty, err := fh.GetPathFromPathFileName(dir)

	fmt.Println("------------------------------------")
	fmt.Println("     GetPathFromPathFileName")
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println("Original Directory: ", dir)
	fmt.Println("Expected Directory: ", expectedDir)
	fmt.Println("  Result Directory: ", result)
	fmt.Println("           isEmpty: ", isEmpty)
	if err != nil {
		fmt.Println("               err: ", err.Error())
	}
	fmt.Println("------------------------------------")

}

func TestGetFileName(pathFileName string) {
	fh := pathFileOp.FileHelper{}

	rawDir := fh.AdjustPathSlash(pathFileName)

	fNameWithExt, isFNameWithExtEmpty, err := fh.GetFileNameWithExt(rawDir)

	expectedFileName := "dirmgr_01_test"

	fNameWithoutExt, isEmpty, err := fh.GetFileNameWithoutExt(rawDir)

	if err != nil {
		fmt.Printf("Error returned from fh.GetFileNameWithoutExt(rawDir) rawDir='%v' Error='%v' \n", rawDir, err.Error())
	}

	fmt.Println("=====================================")
	fmt.Println("      File Name With Extension")
	fmt.Println("=====================================")
	fmt.Println("          rawDir: ", rawDir)
	fmt.Println("         isEmpty: ", isFNameWithExtEmpty)
	fmt.Println("    fNameWithExt: ", fNameWithExt)
	fmt.Println()

	fmt.Println("=====================================")
	fmt.Println("    File Name Without Extension")
	fmt.Println("=====================================")
	fmt.Println("          rawDir: ", rawDir)
	fmt.Println("         isEmpty: ", isEmpty)
	fmt.Println(" fNameWithoutExt: ", fNameWithoutExt)
	fmt.Println("----------------------")
	fmt.Println("expectedFileName: ", expectedFileName)

}

func TestDirMgrFileInfo() {
	fh := pathFileOp.FileHelper{}
	origDir := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/logTest")

	dMgr, err := pathFileOp.DirMgr{}.New(origDir)

	if err != nil {
		fmt.Printf("Error returned from pathFileOp.DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir='%v' Error='%v'", origDir, err.Error())
		return
	}

	fsc := pathFileOp.FileSelectionCriteria{}

	findfiles, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		fmt.Printf("Error returned from dMgr.FindWalkDirFiles(fsc) "+
			"dMgr.absolutePath='%v'  Error='%v'. \n",
			dMgr.GetAbsolutePath(), err.Error())
		return
	}

	lDirs := findfiles.Directories.GetNumOfDirs()

	if lDirs == 0 {
		fmt.Println("Didn't find any directories")
		return
	}

	for i := 0; i < lDirs; i++ {

		foundDMgr, err := findfiles.Directories.PeekDirMgrAtIndex(i)

		if err != nil {
			fmt.Printf(
				"Error returned by findfiles.Directories.PeekDirMgrAtIndex(i). "+
					"i='%v' Error='%v' ", i, err.Error())
		}

		PrintDirMgrFields(foundDMgr)
	}

	fmt.Println("Success")
	fmt.Println("absolutePath: ", dMgr.GetAbsolutePath())

}

func TestFilterFile() {

	fia := pathFileOp.FileInfoPlus{}
	fia.SetName("newerFileForTest_01.txt")
	fia.SetMode(0777)
	fia.SetSize(107633)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fModTimeStr := "2017-10-01 00:00:00.000000000 -0500 CDT"
	fModTime, err := time.Parse(fmtstr, fModTimeStr)

	if err != nil {
		fmt.Printf("Error returned from time.Parse(fmtstr, fModTimeStr). fmtstr='%v' fModTimeStr='%v' Error='%v'\n", fmtstr, fModTimeStr, err.Error())
		return
	}

	fia.SetModTime(fModTime)
	fia.SetIsDir(false)
	fia.SetSysDataSrc(nil)
	fia.SetIsDir(false)

	searchPattern := "*.txt"

	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	//filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	fsc := pathFileOp.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOp.FileSelectMode.ORSelect()

	fh := pathFileOp.FileHelper{}
	isFound, err := fh.FilterFileName(fia, fsc)

	if !isFound {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("!!!!! FILE NOT FOUND !!!!!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!")

	} else {

		fmt.Println("-------------------")
		fmt.Println("SUCCESS File Found!")
		fmt.Println("-------------------")

	}

	fmt.Println("       search Pattern: ", searchPattern)
	fmt.Println("search filesOlderThan: ", filesOlderThan.Format(fmtstr))
	fmt.Println("search filesNewerThan: ", filesNewerThan.Format(fmtstr))
	fmt.Println("             fia.Name: ", fia.Name())
	fmt.Println("             fia.Mode: ", fia.Mode())
	fmt.Println("             fia.Size: ", fia.Size())
	fmt.Println("          fia.ModTime: ", fia.ModTime().Format(fmtstr))
	fmt.Println("            fia.IsDir: ", fia.IsDir())
	fmt.Println("              fia.Sys: ", fia.Sys())
}

func DeleteDir01() {

	origDir, err := testMain01CreateCheckFiles03DirFiles()

	if err != nil {
		fmt.Printf("Error returned from testMain01CreateCheckFiles03DirFiles(). Error='%v'", err.Error())
		return
	}

	fmt.Println("origDir Created")
	fmt.Println("origDig = ", origDir)

}

func DeleteDir02() error {

	fh := pathFileOp.FileHelper{}
	origDir := fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")
	ePrefix := "main.DeleteDir_01 "
	dMgr, err := pathFileOp.DirMgr{}.New(origDir)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DirMgr{}.NewFromPathFileNameExtStr(dirToDelete). dirToDelete='%v' Error='%v'", origDir, err.Error())
	}

	err = dMgr.DeleteAll()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DirMgr{}.NewFromPathFileNameExtStr(dirToDelete). dirToDelete='%v' Error='%v'", origDir, err.Error())
	}

	if dMgr.DoesDirMgrAbsolutePathExist() {
		return fmt.Errorf("Directory origDir still exists. dMgr.doesAbsolutePathExist='%v' "+
			"origDir='%v' ", dMgr.DoesDirMgrAbsolutePathExist(), origDir)
	}

	if dMgr.DoesDirMgrPathExist() {
		return fmt.Errorf("Directory origDir still exists. dMgr.doesPathExist='%v' origDir='%v' ",
			dMgr.DoesDirMgrPathExist(), origDir)
	}

	fmt.Println("Successfully Deleted Directory")
	fmt.Println("origDir: ", origDir)

	return nil
}

func TestMainClean001(srcFile1 string) {

	srcFileCleaned := fp.Clean(srcFile1)
	fh := pathFileOp.FileHelper{}
	srcFileAdjusted := fh.AdjustPathSlash(srcFile1)
	dot := path.Ext(srcFile1)
	splitPath, splitFile := fp.Split(srcFile1)
	fmt.Println("       srcFile1:", srcFile1)
	fmt.Println(" srcFileCleaned:", srcFileCleaned)
	fmt.Println("srcFileAdjusted:", srcFileAdjusted)
	fmt.Println("dot from srcFile1:", dot)
	fmt.Println("splitPath srcFile1:", splitPath)
	fmt.Println("splitFile srcFile1:", splitFile)

}
func TestMain901() {

	fh := pathFileOp.FileHelper{}
	srcFile := fh.AdjustPathSlash("..\\logTest\\Level01\\Level02\\TestFile001.txt")

	if !fh.DoesFileExist(srcFile) {
		fMgr, err := pathFileOp.FileMgr{}.NewFromPathFileNameExtStr(srcFile)

		if err != nil {
			fmt.Printf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		err = fMgr.CreateDirAndFile()

		if err != nil {
			fmt.Printf("Error returned by FileMgr{}.CreateDirAndFile(). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		doesFileExist, err := fMgr.DoesThisFileExist()
		doesFileExist2 := fh.DoesFileExist(fMgr.GetAbsolutePathFileName())

		if err != nil {
			fmt.Printf("Error returned by FileMgr{}.DoesThisFileExist(). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		if !doesFileExist {
			fmt.Printf("Failed to create Source File == '%v'\n", fMgr.GetAbsolutePathFileName())
			return
		}

		if !doesFileExist2 {
			fmt.Printf("Exist2: Failed to create Source File == '%v'\n", fMgr.GetAbsolutePathFileName())
			return

		}

	}

	fmt.Println("**** SUCCESS ****")
}

func testMain01CreateCheckFiles03DirFiles() (string, error) {
	ePrefix := "TestFile: dirmgr_01_test.go Func: testDirMgrCreateCheckFiles03DirFiles() "
	fh := pathFileOp.FileHelper{}

	origDir := fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")

	if fh.DoesFileExist(origDir) {

		err := os.RemoveAll(origDir)

		if err != nil {
			return "", fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
		}

	}

	if fh.DoesFileExist(origDir) {
		return "", fmt.Errorf(ePrefix+
			"Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
	}

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err := os.MkdirAll(origDir, ModePerm)

	if err != nil {
		return "", fmt.Errorf(ePrefix+"Error returned from os.MkdirAll(origDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origDir) {
		return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origDir='%v'", origDir)
	}

	fileDir := origDir + string(os.PathSeparator)
	newFile1 := fileDir + "checkFile30001.txt"
	fp1, err := os.Create(newFile1)

	if err != nil {
		return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile1). newFile1='%v' Error='%v' ", newFile1, err.Error())
	}

	newFile2 := fileDir + "checkFile30002.txt"

	fp2, err := os.Create(newFile2)

	if err != nil {
		_ = fp1.Close()

		return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile2). newFile2='%v' Error='%v' ", newFile2, err.Error())
	}

	newFile3 := fileDir + "checkFile30003.txt"

	fp3, err := os.Create(newFile3)

	if err != nil {
		_ = fp1.Close()
		_ = fp2.Close()
		return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile3). newFile3='%v' Error='%v' ", newFile3, err.Error())
	}

	newFile4 := fileDir + "checkFile30004.txt"

	fp4, err := os.Create(newFile4)

	if err != nil {

		_ = fp1.Close()
		_ = fp2.Close()
		_ = fp3.Close()

		return "", fmt.Errorf(ePrefix+"Error returned from os.Create(newFile4). newFile4='%v' Error='%v' ", newFile4, err.Error())
	}

	du := appLib.DateTimeUtility{}

	_, err = fp4.WriteString(du.GetDateTimeYMDAbbrvDowNano(time.Now()))

	if err != nil {
		_ = fp1.Close()
		_ = fp2.Close()
		_ = fp3.Close()
		_ = fp4.Close()

		return "", fmt.Errorf(ePrefix+
			"%v", err.Error())
	}

	_ = fp1.Close()
	_ = fp2.Close()
	_ = fp3.Close()
	_ = fp4.Close()

	return origDir, nil
}

func MainDirMgrTestSetupFileWalkDeleteFiles() (string, error) {
	ePrefix := "appExample.MainDirMgrTestSetupFileWalkDeleteFiles() "

	fh := pathFileOp.FileHelper{}

	origDir := getBaseProjectPath() + fh.AdjustPathSlash("/dirwalkdeletetests/dirdelete01")

	if fh.DoesFileExist(origDir) {

		err := os.RemoveAll(origDir)

		if err != nil {
			return "", fmt.Errorf(ePrefix+"Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
		}

	}

	if fh.DoesFileExist(origDir) {
		return "", fmt.Errorf(ePrefix+"Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
	}

	origFullDir := origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03"

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err := os.MkdirAll(origFullDir, ModePerm)

	if err != nil {
		return "", fmt.Errorf(ePrefix+"Error returned from os.MkdirAll(origFullDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origFullDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origFullDir) {
		return "", fmt.Errorf(ePrefix+"Error: Failed to create directory! origFullDir='%v'", origFullDir)
	}

	// Copy Old Files
	dirOldFilesForTest := getBaseProjectPath() + fh.AdjustPathSlash("/filesfortest/oldfilesfortest")

	if !fh.DoesFileExist(dirOldFilesForTest) {
		return "", fmt.Errorf(ePrefix+"Error: Old Files Directory does NOT exist! dirOldFilesForTest='%v'", dirOldFilesForTest)

	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"

	srcFile := dirOldFilesForTest + string(os.PathSeparator) + oldFile1
	destFile := origDir + string(os.PathSeparator) + oldFile1

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + oldFile2

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + oldFile3

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	// Copy NewFromPathFileNameExtStr Files
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	dirNewFilesForTest := getBaseProjectPath() + fh.AdjustPathSlash("/filesfortest/newfilesfortest")

	if !fh.DoesFileExist(dirNewFilesForTest) {
		return "", fmt.Errorf(ePrefix+"Error: NewFromPathFileNameExtStr Files Directory does NOT exist! dirNewFilesForTest='%v'", dirNewFilesForTest)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile1
	destFile = origDir + string(os.PathSeparator) + newFile1

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + newFile2

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + newFile3

	err = fh.CopyFileByIo(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	return origDir, nil
}

func getBaseProjectPath() string {

	ePrefix := "getBaseProjectPath() "
	fh := pathFileOp.FileHelper{}
	currDir, err := fh.GetAbsCurrDir()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by fh.GetAbsCurrDir(). Error='%v' \n", err.Error())
		return "Error"
	}

	target := "pathfileopsgo"
	idx := strings.Index(currDir, target)

	if idx < 0 {
		fmt.Printf(ePrefix +
			"Error: Unable to locate \"pathfileopsgo\" in current directory string! \n")
		return "Error"
	}

	idx += len(target)

	baseDir := currDir[0:idx]

	return baseDir
}
