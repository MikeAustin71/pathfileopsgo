package main

import (
	"MikeAustin71/pathfilego/003_filehelper/common"
	"fmt"
	fp "path/filepath"
	"path"
	"os"
	"time"
	"strings"
	"io/ioutil"
)

/*

import (
	"MikeAustin71/pathfilego/003_filehelper/common"
	"fmt"
	"time"
)

const (
	baseAppDir = "../../003_filehelper/app"
	// commonDir        = "../common"
	logTestTopDIR = "../logTest"
	// logTestBottomDir = "../logTest/CmdrX"
	// logFile          = "CmdrX.log"
)

*/

func main() {
	TestWriteFile()
}

func TestWriteFile() {
	fh := common.FileHelper{}

	filePath := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/checkfiles/checkfiles03/testWrite2998.txt")

	fMgr, err := common.FileMgr{}.New(filePath)

	if err != nil {
		fmt.Printf("Error returned from common.FileMgr{}.New(filePath). filePathName='%v'  Error='%v'", filePath, err.Error())
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
		fmt.Printf("Error returned from fMgr.ReadAllFile(). filePathName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
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
		fmt.Printf("Error: After writing string, target file does NOT exist!. FileNameExt='%v'", fMgr.AbsolutePathFileName)
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
		fmt.Printf("Error: Failed to DELETE FileNameExt='%v'", fMgr.AbsolutePathFileName)
		return
	}

	fmt.Println("Successful Completion!")
}

func TestOpenFile() {
	fh := common.FileHelper{}
	ePrefix := "TestOpenFile() "

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/testRead2008.txt")

	fMgr, err := common.FileMgr{}.New(filePath)

	if err != nil {
		fmt.Printf(ePrefix + "Error returned from common.FileMgr{}.New(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	err = fMgr.OpenThisFileReadOnly()

	if err != nil {
		fmt.Printf(ePrefix + "Error returned from fMgr.OpenThisFileReadOnly(). filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	defer fMgr.CloseFile()

	b, err := ioutil.ReadAll(fMgr.FilePtr)

	if err != nil {
		fmt.Printf(ePrefix + "Error returned from ioutil.ReadAll(fMgr.FilePtr) filePath='%v'  Error='%v'", filePath, err.Error())
		return
	}

	str := string(b)

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
	fh := common.FileHelper{}
	substituteDir:= fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/testdestdir/destdir")

	substituteDMgr, err := common.DirMgr{}.New(substituteDir)

	if err != nil {
		fmt.Printf(ePrefix + "Error returned by common.DirMgr{}.New(substituteDir). substituteDir='%v'  Error='%v'\n", substituteDir, err.Error())
		return
	}

	err = substituteDMgr.DeleteAll()

	if err != nil {
		fmt.Printf(ePrefix + "Error returned by substituteDMgr.DeleteAll(). substituteDMgr.Path='%v'  Error='%v'\n", substituteDMgr.Path, err.Error() )
		return
	}

	fmt.Println("Successful Completion")

}

func TestCopyDirectoryTree() {

	ePrefix := "TestCopyDirectoryTree() "
	fh := common.FileHelper{}
	dir := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/testsrcdir")

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	dMgr, err := common.DirMgr{}.New(dir)

	if err!=nil {
		fmt.Printf("Error returned from DirMgr{}.New(dir). dir='%v' Error='%v'\n", dir, err.Error())
		return
	}

	if !dMgr.AbsolutePathDoesExist {
		fmt.Printf("Expected target directory to exist. I does NOT exist. dMgr.Path='%v' dMgr.AbolutePath='%v'\n", dMgr.Path, dMgr.AbsolutePath)
		return
	}

	fsc := common.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = common.ANDFILESELECTCRITERION

	dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err!=nil {
		fmt.Printf("Error returned from dMgr.FindWalkDirFiles(searchPattern, filesOlderThan, filesNewerThan). dir='%v' Error='%v'\n", dir, err.Error())
		return
	}

	baseDir:= fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/testsrcdir")

	baseDMgr, err := common.DirMgr{}.New(baseDir)

	if err != nil {
		fmt.Printf("Error returned by common.DirMgr{}.New(baseDir) baseDir='%v' Error='%v'", baseDir, err.Error())
		return
	}

	substituteDir:= fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/testdestdir/destdir")

	substituteDMgr, err := common.DirMgr{}.New(substituteDir)

	if err != nil {
		fmt.Printf(ePrefix + "Error returned by common.DirMgr{}.New(substituteDir). substituteDir='%v'  Error='%v'", substituteDir, err.Error())
		return
	}

	newDirTree, err := dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

	if err != nil {
		fmt.Printf(ePrefix + "Error returned by dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr). Error='%v'")
		return
	}

	if len(dirTreeInfo.Directories.DirMgrs) != len(newDirTree.Directories.DirMgrs) {

		fmt.Printf(ePrefix + "Error: Expected Number of Directories = '%v'.  Instead, Number of New Directories = '%v'", len(dirTreeInfo.Directories.DirMgrs), len(newDirTree.Directories.DirMgrs))
		return
	}

	if len(dirTreeInfo.FoundFiles.FMgrs) != len(newDirTree.FoundFiles.FMgrs) {
		fmt.Printf(ePrefix + "Error: Expected Number of Files = '%v'.  Instead, actual Number of New Files = '%v'",len(dirTreeInfo.FoundFiles.FMgrs), len(newDirTree.FoundFiles.FMgrs))
		return
	}

	for i:=0; i < len(newDirTree.FoundFiles.FMgrs); i++ {
		doesFileExist, err := newDirTree.FoundFiles.FMgrs[i].DoesThisFileExist()

		if err != nil {
			fmt.Printf(ePrefix + "Error returned by newDirTree.FoundFiles.FMgrs[i].DoesThisFileExist(). i='%v' FileNameExt='%v'  Error='%v'", i, newDirTree.FoundFiles.FMgrs[i].FileNameExt, err.Error())
			return
		}

		if !doesFileExist {
			fmt.Printf(ePrefix + "Error: Failed to create FileNameExt='%v'. It does NOT exist in target directory.",newDirTree.FoundFiles.FMgrs[i].FileNameExt)
			return
		}

	}

	fmt.Println("Successful Completion")
}

func TestMainCleanDirStr(rawPath string) {
	fh := common.FileHelper{}

	dirPath, isEmpty, err := fh.CleanDirStr(rawPath)

	if err != nil {
		fmt.Printf("Error returned by fh.CleanDirStr(rawPath) rawPath='%v'  Error='%v' \n", rawPath, err.Error())
		return
	}

	fmt.Println("=========================================")
	fmt.Println("     Test Clean Directory String")
	fmt.Println("=========================================")
	fmt.Println()
	fmt.Println("Returned Dir Path: ", dirPath)
	fmt.Println("          isEmpty: ", isEmpty)
	fmt.Println("  raw path string: ", rawPath)

}

func TestNewFileMgrFromPathFileNameStr(pathFileNameExt string) {

	fMgr, err := common.FileMgr{}.New(pathFileNameExt)

	if err!=nil {
		fmt.Printf("Error returned from common.FileMgr{}.New(pathFileNameExt) pathFileNameExt='%v'  Error='%v' \n", pathFileNameExt, err.Error())
	}

	common.PrintFileManagerFields(fMgr)

}

func TestNewFileMgrFromDirMgrFileNameExt(rawPath, rawFileNameExt string) {

fh:= common.FileHelper{}
adjustedPath := fh.AdjustPathSlash(rawPath)

	dMgr, err := common.DirMgr{}.New(adjustedPath)

	if err != nil {
		fmt.Printf("Error returned from DirMgr{}.New(adjustedPath). adjustedPath='%v'  Error='%v' \n", adjustedPath, err.Error())
		return
	}

	fMgr, err := common.FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

	if err != nil {
		fmt.Printf("Error returned by FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt). dMgr.Path='%v' rawFileNameExt='%v'  \n", dMgr.Path, rawFileNameExt)
	}

	common.PrintFileManagerFields(fMgr)

}

func TestDirMgr(rawPath string, expectedPath string) {

	fh := common.FileHelper{}


	expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

	if err!= nil {
		fmt.Printf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'\n", expectedPath, err.Error())
		return
	}

	dMgr, err := common.DirMgr{}.New(rawPath)

	if err != nil {
		fmt.Printf("Error returned from DirMgr{}.New(rawPath) rawPath='%v' Error='%v'", rawPath, err.Error())
		return
	}

	common.PrintDirMgrFields(dMgr)

	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("             Expected Results")
	fmt.Println("=========================================")
	fmt.Println("          rawPath: ", rawPath)
	fmt.Println("     expectedPath: ", expectedPath)
	fmt.Println("   expectedAbsDir: ", expectedAbsDir)

}

func TestMainCleanFileNameExt(rawFileNameExt, expectedFileNameExt string) {

	fh := common.FileHelper{}
	adjustedFileNameExt := fh.AdjustPathSlash(rawFileNameExt)
	processedFileNameExt, isFileNameEmpty, err := fh.CleanFileNameExtStr(adjustedFileNameExt)

	if err != nil {
		fmt.Printf("Error returned by fh.CleanFileNameExtStr(adjustedFileNameExt) adjustedFileNameExt='%v' Error='%v'\n", adjustedFileNameExt, err.Error())
		return
	}

	fmt.Println("=========================================")
	fmt.Println("          Clean FileNameExt Tests")
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

	adjustedAbsolutePath, err := 	fp.Abs(adjustedPath)

	if err != nil {
		fmt.Printf("Error returned from fp.Abs(adjustedPath) adjustedPath='%v' Error='%v' \n", adjustedPath, err.Error())
		return
	}

	fpCleanedAbsolutePath := fp.Clean(adjustedAbsolutePath)

	pathCleanedAbsolutePath:= path.Clean(adjustedAbsolutePath)


	fmt.Println("=========================================")
	fmt.Println("           Clean Path Tests")
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

	fh := common.FileHelper{}

	result, isEmpty, err := fh.GetPathFromPathFileName(dir)

	fmt.Println("------------------------------------")
	fmt.Println("     GetPathFromPathFileName")
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println("Original Directory: ", dir)
	fmt.Println("Expected Directory: ", expectedDir)
	fmt.Println("  Result Directory: ", result )
	fmt.Println("           isEmpty: ", isEmpty)
	if err!=nil {
		fmt.Println("               err: ", err.Error())
	}
	fmt.Println("------------------------------------")



}


func TestGetFileName(pathFileName string) {
	fh := common.FileHelper{}

	rawDir := fh.AdjustPathSlash(pathFileName)

	fNameWithExt, isFNameWithExtEmpty, err := fh.GetFileNameWithExt(rawDir)

	expectedFileName := "dirmgr_01_test"

	fNameWithoutExt, isEmpty, err := fh.GetFileNameWithoutExt(rawDir)

	if err != nil {
		fmt.Printf("Error returned from fh.GetFileNameWithoutExt(rawDir) rawDir='%v' Error='%v' \n", rawDir, err.Error() )
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
	fh:= common.FileHelper{}
	origDir :=  fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/logTest")

	dMgr, err := common.DirMgr{}.New(origDir)

	if err!=nil {
		fmt.Printf("Error returned from common.DirMgr{}.New(origDir). origDir='%v' Error='%v'", origDir, err.Error())
		return
	}

	fsc := common.FileSelectionCriteria{}

	findfiles, err :=  dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		fmt.Printf("Error returned from dMgr.FindWalkDirFiles(fsc) dMgr.AbsolutePath='%v'  Error='%v'. \n",dMgr.AbsolutePath, err.Error())
		return
	}

	lDirs := findfiles.Directories.GetArrayLength()

	if lDirs == 0 {
		fmt.Println("Didn't find any directories")
		return
	}

	for i:=0; i < lDirs ; i++ {
		foundDMgr := findfiles.Directories.DirMgrs[i]

		common.PrintDirMgrFields(foundDMgr)
	}

	fmt.Println("Success")
	fmt.Println("AbsolutePath: ", dMgr.AbsolutePath)

}

func TestFilterFile() {

	fia := common.FileInfoPlus{}
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

	fsc := common.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = common.ORFILESELECTCRITERION


	fh := common.FileHelper{}
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


func TestDirMgrWalDirDeleteFiles() {

	origDir, err :=  MainDirMgrTestSetupFileWalkDeleteFiles()

	if err != nil {
		fmt.Printf("Error returned from DirMgr02TestSetupFileWalkDeleteFiles(). Error='%v'\n", err.Error())
		return
	}

	fh := common.FileHelper{}

	if ! fh.DoesFileExist(origDir) {
		fmt.Printf("Error: The target directory does NOT Exist! origDir='%v'\n", origDir)
		return
	}

	dMgr, err := common.DirMgr{}.New(origDir)

	if err != nil {
		fmt.Printf("Error returned by DirMgr{}.New(origDir). origDir='%v' Error='%v'\n", origDir, err.Error())
		return
	}

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	fOlderThanStr := "2017-12-01 00:00:00.000000000 -0600 CST"

	searchPattern := "*.txt"
	filesOlderThan, err := time.Parse(fmtstr, fOlderThanStr)

	if err != nil {
		fmt.Printf("Error returned from time.Parse(fmtstr, fOlderThanStr). fOlderThanStr='%v' Error='%v'\n", fOlderThanStr, err.Error())
		return
	}

	filesNewerThan := time.Time{}

	fsc := common.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = common.ORFILESELECTCRITERION

	dInfo, err := dMgr.DeleteWalkDirFiles(fsc)

	if dInfo.DeletedFiles.GetArrayLength() != 6 {
		fmt.Printf("Expected to find 6-files deleted. Instead, %v-files were deleted.\n", dInfo.DeletedFiles.GetArrayLength())
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
		fmt.Printf("Expected deletion of oldFile1='%v'. The file was NOT deleted!\n",oldFile1)
	}

	if oldFile2Found == false {
		fmt.Printf("Expected deletion of oldFile2='%v'. The file was NOT deleted!\n",oldFile2)
	}

	if oldFile3Found == false {
		fmt.Printf("Expected deletion of oldFile3='%v'. The file was NOT deleted!\n",oldFile3)
	}

	if newFile1Found == false {
		fmt.Printf("Expected deletion of newFile1='%v'. The file was NOT deleted!\n",newFile1)
	}

	if newFile2Found == false {
		fmt.Printf("Expected deletion of newFile2='%v'. The file was NOT deleted!\n",newFile2)
	}

	if newFile3Found == false {
		fmt.Printf("Expected deletion of newFile3='%v'. The file was NOT deleted!\n",newFile3)
	}


	if len(dInfo.ErrReturns) != 0 {
		fmt.Printf("Expected zero Error Returns. Instead number of Error Returns='%v'\n", len(dInfo.ErrReturns))
		return
	}

	if dInfo.Directories.GetArrayLength() != 3 {
		fmt.Printf("Expected 3-directories to be found. Instead, number of directories found='%v'\n", dInfo.Directories.GetArrayLength())
		return
	}

	fmt.Println("Success!")
	fmt.Println("origDir: ", origDir)

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

	fh := common.FileHelper{}
	origDir :=  fh.AdjustPathSlash("../checkfiles/checkfiles02/checkfiles03")
	ePrefix := "main.DeleteDir_01 "
	dMgr, err := common.DirMgr{}.New(origDir)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from DirMgr{}.New(dirToDelete). dirToDelete='%v' Error='%v'", origDir, err.Error())
	}

	err = dMgr.DeleteAll()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from DirMgr{}.New(dirToDelete). dirToDelete='%v' Error='%v'", origDir, err.Error())
	}

	if dMgr.AbsolutePathDoesExist {
		return fmt.Errorf("Directory origDir still exists. dMgr.AbsolutePathDoesExist='%v' origDir='%v' ", dMgr.AbsolutePathDoesExist, origDir)
	}

	if dMgr.PathDoesExist {
		return fmt.Errorf("Directory origDir still exists. dMgr.PathDoesExist='%v' origDir='%v' ", dMgr.PathDoesExist, origDir)
	}

	fmt.Println("Successfully Deleted Directory")
	fmt.Println("origDir: ", origDir)

	return nil
}

func TestMainClean001(srcFile1 string) {

	srcFileCleaned := fp.Clean(srcFile1)
	fh := common.FileHelper{}
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


	fh := common.FileHelper{}
	srcFile := fh.AdjustPathSlash("..\\logTest\\Level01\\Level02\\TestFile001.txt")

	if !fh.DoesFileExist(srcFile) {
		fmgr, err := common.FileMgr{}.New(srcFile)

		if err!=nil {
			fmt.Printf("Error returned by FileMgr{}.New(srcFile). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		err = fmgr.CreateDirAndFile()

		if err!=nil {
			fmt.Printf("Error returned by FileMgr{}.CreateDirAndFile(). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		doesFileExist, err := fmgr.DoesThisFileExist()
		doesFileExist2 := fh.DoesFileExist(fmgr.AbsolutePathFileName)

		if err!=nil {
			fmt.Printf("Error returned by FileMgr{}.DoesThisFileExist(). srcFile='%v'. Error='%v'\n", srcFile, err.Error())
			return
		}

		if !doesFileExist {
			fmt.Printf("Failed to create Source File == '%v'\n", fmgr.AbsolutePathFileName)
			return
		}

		if !doesFileExist2 {
			fmt.Printf("Exist2: Failed to create Source File == '%v'\n", fmgr.AbsolutePathFileName)
			return

		}

	}

	fmt.Println("**** SUCCESS ****")
}

func testMain01CreateCheckFiles03DirFiles() (string, error) {
	ePrefix := "TestFile: dirmgr_01_test.go Func: testDirMgrCreateCheckFiles03DirFiles() "
	fh := common.FileHelper{}

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

	du := common.DateTimeUtility{}

	fp4.WriteString(du.GetDateTimeYMDAbbrvDowNano(time.Now()))


	return origDir, nil
}

func MainDirMgrTestSetupFileWalkDeleteFiles() (string, error) {
	ePrefix := "TestFile: main.go Func: MainDirMgrTestSetupFileWalkDeleteFiles() "

	fh := common.FileHelper{}

	origDir := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/dirwalkdeletetests/dirdelete01")

	if fh.DoesFileExist(origDir) {

		err :=	os.RemoveAll(origDir)

		if err != nil {
			return "", fmt.Errorf(ePrefix + "Error returned by os.RemoveAll(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
		}

	}

	if fh.DoesFileExist(origDir){
		return "", fmt.Errorf(ePrefix + "Error: Attempted to delete origDir='%v'. However, it still Exists!", origDir)
	}

	origFullDir := origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03"

	// origDir does NOT exist!
	var ModePerm os.FileMode = 0777

	err := os.MkdirAll(origFullDir, ModePerm)

	if err != nil {
		return "", fmt.Errorf(ePrefix + "Error returned from os.MkdirAll(origFullDir, ModePerm). origDir='%v' ModePerm='%v'  Error='%v'", origFullDir, ModePerm, err.Error())
	}

	if !fh.DoesFileExist(origFullDir) {
		return "", fmt.Errorf(ePrefix + "Error: Failed to create directory! origFullDir='%v'", origFullDir)
	}

	// Copy Old Files
	dirOldFilesForTest := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/filesfortest/oldfilesfortest")

	if !fh.DoesFileExist(dirOldFilesForTest) {
		return "", fmt.Errorf(ePrefix + "Error: Old Files Directory does NOT exist! dirOldFilesForTest='%v'", dirOldFilesForTest)

	}

	oldFile1 := "test.htm"
	oldFile2 := "006890_WritingFiles.htm"
	oldFile3 := "006870_ReadingFiles.htm"

	srcFile := dirOldFilesForTest + string(os.PathSeparator) + oldFile1
	destFile := origDir + string(os.PathSeparator) + oldFile1

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + oldFile2

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirOldFilesForTest + string(os.PathSeparator) + oldFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + oldFile3

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	// Copy New Files
	newFile1 := "newerFileForTest_01.txt"
	newFile2 := "newerFileForTest_02.txt"
	newFile3 := "newerFileForTest_03.txt"

	dirNewFilesForTest := fh.AdjustPathSlash("D:/go/work/src/MikeAustin71/pathfilego/003_filehelper/filesfortest/newfilesfortest")

	if !fh.DoesFileExist(dirNewFilesForTest) {
		return "", fmt.Errorf(ePrefix + "Error: New Files Directory does NOT exist! dirNewFilesForTest='%v'", dirNewFilesForTest)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile1
	destFile = origDir + string(os.PathSeparator) + newFile1

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile2
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + newFile2

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	srcFile = dirNewFilesForTest + string(os.PathSeparator) + newFile3
	destFile = origDir + string(os.PathSeparator) + "dirdelete02" + string(os.PathSeparator) + "dirdelete03" + string(os.PathSeparator) + newFile3

	err = fh.CopyToNewFile(srcFile, destFile)

	if err != nil {
		return "", fmt.Errorf("Error while Copying Source File, '%v' to  Destination File '%v', Error:'%v'", srcFile, destFile, err)
	}

	return origDir, nil
}
