package pathfileops

import (
	"fmt"
	"testing"
	"time"
)

func TestFileMgrCollection_AddFileMgrCollection(t *testing.T) {

	var fileNameExt string

	fMgrs1 := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs1.AddFileMgr(fmgr)
	}

	if fMgrs1.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs1 Array Length == 10. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
	}

	fMgrs2 := FileMgrCollection{}

	for i := 0; i < 15; i++ {

		fileNameExt = fmt.Sprintf("testCol2AddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from 2nd run of testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs2.AddFileMgr(fmgr)
	}

	if fMgrs2.GetNumOfFileMgrs() != 15 {
		t.Errorf("Expected fMgrs2 Array Length == 15. Instead fMgrs2.GetNumOfDirs()=='%v'", fMgrs2.GetNumOfFileMgrs())
	}

	fMgrs1.AddFileMgrCollection(&fMgrs2)

	if fMgrs1.GetNumOfFileMgrs() != 25 {
		t.Errorf("Expected augmented fMgrs1 Array Length == 25. Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
	}

	fMgr, err := fMgrs1.PeekLastFileMgr()

	if err != nil {
		t.Errorf("2nd Run: Error returned from fMgrs1.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if fMgr.fileNameExt != "testCol2AddFile_015.txt" {
		t.Errorf("Expected consolidated fMgrs1 to have last fMgr.fileNameExt='testCol2AddFile_015.txt'. Instead, fMgr.fileNameExt='%v'", fMgr.fileNameExt)
	}

}

func TestFileMgrCollection_AddFileMgr_01(t *testing.T) {
	var fileNameExt string

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	lastFmgr, err := fMgrs.PeekLastFileMgr()

	if err != nil {
		t.Errorf("Error returned by fMgrs.PeekLastDirMgr(). Error='%v'", err)
	}

	if lastFmgr.fileNameExt != "testAddFile_010.txt" {
		t.Errorf("Expected last File Manager to have fileNameExt='testAddFile_010.txt'. Instead fileNameExt='%v'", lastFmgr.fileNameExt)
	}

}

func TestFileMgrCollection_AddFileMgrByPathFile(t *testing.T) {

	var fileNameExt string
	fh := FileHelper{}

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	adjustedPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	fPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
	}

	err = fMgrs.AddFileMgrByPathFile(fPath)

	if err != nil {
		t.Errorf("Error returned from fMgrs.AddFileMgrByPathFile(fPath). fPath='%v' Error='%v'", fPath, err.Error())
	}

	fmgr2, err := fMgrs.PeekLastFileMgr()

	if err != nil {
		t.Errorf("Error returned by fMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if fmgr2.fileNameExt != "newerFileForTest_01.txt" {
		t.Errorf("Expected Newly Added Fmgr fileNameExt='newerFileForTest_01.txt'. Instead, fileNameExt='%v'", fmgr2.fileNameExt)
	}

}

func TestFileMgrCollection_AddFileInfo_01(t *testing.T) {

	var fileNameExt string
	fh := FileHelper{}

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). "+
				"fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}

		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'",
			fMgrs.GetNumOfFileMgrs())
	}

	expectedFileNameExt := "newerFileForTest_01.txt"

	fic := FileInfoPlus{}
	fic.SetName(expectedFileNameExt)
	fic.SetIsDir(false)
	fic.SetSize(123456)
	fic.SetModTime(time.Now().Local())
	fic.SetMode(0666)
	fic.SetSysDataSrc("xyzxyzxyzyzx")
	fic.SetIsFInfoInitialized(true)

	adjustedPath := "../filesfortest/newfilesfortest"

	fPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
	}

	err = fMgrs.AddFileMgrByFileInfo(fPath, fic)

	if err != nil {
		t.Errorf("Error returned from fMgrs.AddFileMgrByFileInfo(fPath, fic). fPath='%v' Error='%v'", fPath, err.Error())

	}

	if fMgrs.GetNumOfFileMgrs() != 11 {
		t.Errorf("Expected fMgrs Array Length == 11. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	fmgrLast, err := fMgrs.PopLastFileMgr()

	if fmgrLast.fileNameExt != expectedFileNameExt {
		t.Errorf("Expected fmgrLast.fileNameExt='%v'.  Instead, fmgrLast.fileNameExt='%v'", expectedFileNameExt, fmgrLast.fileNameExt)
	}

}

func TestFileMgrCollection_PopFMgrAtIndex(t *testing.T) {

	var fileNameExt string

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	fmgrOut, err := fMgrs.PopFileMgrAtIndex(5)

	if err != nil {
		t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5). Error='%v'", err.Error())
	}

	if fMgrs.GetNumOfFileMgrs() != 9 {
		t.Errorf("Expected after Pop Array fMgrs Array Length == 9. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	if fmgrOut.fileNameExt != "testAddFile_006.txt" {
		t.Errorf("Expected popped file manger at index=5 to be fileNameExt='testAddFile_006.txt'. Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
	}

}

func TestFileMgrCollection_PeekFMgrAtIndex(t *testing.T) {

	var fileNameExt string

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	fmgrOut, err := fMgrs.PeekFileMgrAtIndex(5)

	if err != nil {
		t.Errorf("Error returned by fMgrs.PopDirMgrAtIndex(5). Error='%v'", err.Error())
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected after Peek Array fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	if fmgrOut.fileNameExt != "testAddFile_006.txt" {
		t.Errorf("Expected Peek file manger at index=5 to be fileNameExt='testAddFile_006.txt'. Instead, fileNameExt='%v'", fmgrOut.fileNameExt)
	}

}

func TestFileMgrCollection_PopLastFMgr(t *testing.T) {

	var fileNameExt string

	fMgrs := FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf("testAddFile_%03d.txt", i+1)
		fmgr, err := FileMgrCollectionTestSetupFmgr01(fileNameExt)
		if err != nil {
			t.Errorf("Error returned from testFileMgrCollection_SetupFmgr_01(fileNameExt). fileNameExt='%v'  Error='%v'", fileNameExt, err.Error())
		}
		fMgrs.AddFileMgr(fmgr)
	}

	if fMgrs.GetNumOfFileMgrs() != 10 {
		t.Errorf("Expected fMgrs Array Length == 10. Instead fMgrs.GetNumOfDirs()=='%v'", fMgrs.GetNumOfFileMgrs())
	}

	fmgrLast, err := fMgrs.PopLastFileMgr()

	if err != nil {
		t.Errorf("Error returned from fMgrs.PopLastDirMgr(). Error='%v'", err.Error())
	}

	if fmgrLast.fileNameExt != "testAddFile_010.txt" {
		t.Errorf("Expected PopLastDirMgr() to produce fmgrLast.fileNameExt='testAddFile_010.txt'. Instead, fmgrLast.fileNameExt='%v'", fmgrLast.fileNameExt)
	}

}

func TestFileMgrCollection_FindFiles(t *testing.T) {

	fmgrCol := FileMgrCollectionTestSetup01()

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{"*.txt"}

	fmgrCol2, err := fmgrCol.FindFiles(fsc)

	if err != nil {
		t.Errorf("Error returned by fmgrCol.FindFiles(fsc). Error='%v'", err.Error())
	}

	if fmgrCol2.GetNumOfFileMgrs() != 3 {
		t.Errorf("Expected fmgrCol2.GetNumOfDirs()==3 . Instead, fmgrCol2.GetNumOfDirs()='%v' ",
			fmgrCol2.GetNumOfFileMgrs())
	}

	numOfFoundTextfiles := 0

	for i := 0; i < fmgrCol2.GetNumOfFileMgrs(); i++ {
		if fmgrCol2.fileMgrs[i].fileExt == ".txt" {
			numOfFoundTextfiles++
		}
	}

	if numOfFoundTextfiles != 3 {
		t.Errorf("Expected the number of found text files == 3. Instead, number of found text files=='%v'", numOfFoundTextfiles)
	}

}

// //////////////////////////////////////////////////////////////
// Test Setup Functions
// //////////////////////////////////////////////////////////////
func FileMgrCollectionTestSetup01() FileMgrCollection {

	fh := FileHelper{}
	FMgrs := FileMgrCollection{}

	fPath, _ := fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt"))
	fmgr, _ := FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_02.txt"))
	fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_03.txt"))
	fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	fPath, _ = fh.MakeAbsolutePath(fh.AdjustPathSlash("../filesfortest/oldfilesfortest/006870_ReadingFiles.htm"))
	fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	fPath, _ = fh.MakeAbsolutePath("../filesfortest/oldfilesfortest/006890_WritingFiles.htm")
	fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	fPath, _ = fh.MakeAbsolutePath("../filesfortest/oldfilesfortest/test.htm")
	fmgr, _ = FileMgr{}.NewFromPathFileNameExtStr(fPath)
	FMgrs.AddFileMgr(fmgr)

	return FMgrs
}

func FileMgrCollectionTestSetupFmgr01(fileNameExt string) (FileMgr, error) {

	ePrefix := "Src File: xt_filemgrcollection_01_test.go  Function: FileMgrCollectionTestSetupFmgr01() "
	fh := FileHelper{}

	pathFileName := "../dirwalktests/dir01/dir02/" + fileNameExt
	adjustedPathFileName := fh.AdjustPathSlash(pathFileName)
	fPath, err := fh.MakeAbsolutePath(adjustedPathFileName)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix+"Error return by fh.MakeAbsolutePath(adjustedPathFileName). adjustedPathFileName='%v'  Error='%v'", adjustedPathFileName, err.Error())
	}

	fmgr, err := FileMgr{}.NewFromPathFileNameExtStr(fPath)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix+"Error return by FileMgr{}.NewFromPathFileNameExtStr(fPath). fPath='%v'  Error='%v'", fPath, err.Error())
	}

	return fmgr, nil

}

func TestDirectoryTreeInfo_CopyToDirectoryTree_01(t *testing.T) {

	fh := FileHelper{}
	dir := fh.AdjustPathSlash("../testsrcdir")

	searchPattern := ""
	filesOlderThan := time.Time{}
	filesNewerThan := time.Time{}

	dMgr, err := DirMgr{}.New(dir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(dir). dir='%v' Error='%v'\n", dir, err.Error())
	}

	if !dMgr.doesAbsolutePathExist {
		t.Errorf("Expected target directory to exist. I does NOT exist. dMgr.path='%v' dMgr.AbolutePath='%v'\n", dMgr.path, dMgr.absolutePath)
		return
	}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{searchPattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dirTreeInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error returned from dMgr.FindWalkDirFiles(searchPattern, filesOlderThan, filesNewerThan). dir='%v' Error='%v'\n", dir, err.Error())
	}

	baseDir := fh.AdjustPathSlash("../testsrcdir")

	baseDMgr, err := DirMgr{}.New(baseDir)

	if err != nil {
		t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(baseDir) baseDir='%v' Error='%v'", baseDir, err.Error())
	}

	substituteDir := fh.AdjustPathSlash("../testdestdir/destdir")

	substituteDMgr, err := DirMgr{}.New(substituteDir)

	if err != nil {
		t.Errorf("Error returned by common.DirMgr{}.NewFromPathFileNameExtStr(substituteDir). substituteDir='%v'  Error='%v'", substituteDir, err.Error())
	}

	newDirTree, err := dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr)

	if err != nil {
		t.Errorf("Error returned by dirTreeInfo.CopyToDirectoryTree(baseDMgr, substituteDMgr). Error='%v'",
			err.Error())
	}

	if len(dirTreeInfo.Directories.DirMgrs) != len(newDirTree.Directories.DirMgrs) {

		t.Errorf("Error: Expected Number of Directories = '%v'.  Instead, Number of NewFromPathFileNameExtStr Directories = '%v'", len(dirTreeInfo.Directories.DirMgrs), len(newDirTree.Directories.DirMgrs))
	}

	if len(dirTreeInfo.FoundFiles.fileMgrs) != len(newDirTree.FoundFiles.fileMgrs) {
		t.Errorf("Error: Expected Number of Files = '%v'.  Instead, actual Number of NewFromPathFileNameExtStr Files = '%v'", len(dirTreeInfo.FoundFiles.fileMgrs), len(newDirTree.FoundFiles.fileMgrs))
	}

	for i := 0; i < len(newDirTree.FoundFiles.fileMgrs); i++ {
		doesFileExist, err := newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist()

		if err != nil {
			t.Errorf("Error returned by newDirTree.FoundFiles.fileMgrs[i].DoesThisFileExist(). i='%v' fileNameExt='%v'  Error='%v'", i, newDirTree.FoundFiles.fileMgrs[i].fileNameExt, err.Error())
		}

		if !doesFileExist {
			t.Errorf("Error: Failed to create fileNameExt='%v'. It does NOT exist in target directory.", newDirTree.FoundFiles.fileMgrs[i].fileNameExt)
		}

	}

	err = substituteDMgr.DeleteAll()

	if err != nil {
		t.Errorf("Error returned from substituteDMgr.DeleteAll(). Error='%v'", err.Error())
	}

}
