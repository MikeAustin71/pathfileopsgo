package main

import (
	pf "../pathfileops"
	"fmt"
	"os"
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

func main() {
 if, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok {
        fmt.Println("File at path", err.path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

*/

func main() {

	mainTest22()

}

func mainTest22() {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus, err := pf.FileOpenStatus{}.New(pf.FOpenType.WriteOnly(),
		pf.FOpenMode.Append(), pf.FOpenMode.Truncate())

	if err != nil {
		fmt.Printf("Error returned by FileOpenStatus{}.New(). Error='%v' \n", err.Error())
		return
	}

	actualFileOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		fmt.Printf("Error returned by FileOpenStatus{}.GetCompositeFileOpenCode(). Error='%v' \n", err.Error())
		return
	}

	if expectedFOpenCode != actualFileOpenCode {
		fmt.Printf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
			expectedFOpenCode, actualFileOpenCode)
		return
	}

	fmt.Printf("Success - File Open Codes Match!")
	return
}

func mainTest21() {

	fmt.Println("--------- Primary Codes ---------")
	fmt.Println("os.O_RDONLY: ", os.O_RDONLY)
	fmt.Println("os.O_WRONLY: ", os.O_WRONLY)
	fmt.Println("os.O_RDWR: ", os.O_RDWR)
	fmt.Println()
	fmt.Println("--------- Control Codes ----------")
	fmt.Println("os.O_APPEND: ", os.O_APPEND)
	fmt.Println("os.O_CREATE: ", os.O_CREATE)
	fmt.Println("os.O_EXCL: ", os.O_EXCL)
	fmt.Println("os.O_SYNC: ", os.O_SYNC)
	fmt.Println("os.O_TRUNC: ", os.O_TRUNC)

}

func mainTest20() {

	fh := pf.FileHelper{}

	relPath := "../testfiles"
	origPath := fh.AdjustPathSlash(relPath)

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		fmt.Printf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
			"origPath= '%v'  Error='%v'", origPath, err.Error())
		return
	}

	testDMgr, err := pf.DirMgr{}.New(origAbsPath)

	if err != nil {
		fmt.Printf("Error returned by pf.DirMgr{}.New(origAbsPath). "+
			"origAbsPath= '%v'  Error='%v'", origAbsPath, err.Error())
		return
	}

	var fileNameExt string

	fMgrs1 := pf.FileMgrCollection{}

	for i := 0; i < 10; i++ {

		fileNameExt = fmt.Sprintf(testDMgr.GetAbsolutePathWithSeparator()+"testAddFile_%03d.txt", i+1)

		fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(fileNameExt)

		if err != nil {
			fmt.Printf("Error returned by pf.FileMgr{}.NewFromPathFileNameExtStr(fileNameExt). "+
				"fileNameExt='%v' Error='%v' ", fileNameExt, err.Error())
			return
		}

		fMgrs1.AddFileMgr(fMgr)
	}

	if fMgrs1.GetNumOfFileMgrs() != 10 {
		fmt.Printf("Expected fMgrs1 Array Length == 10. "+
			"Instead fMgrs1.GetNumOfDirs()=='%v'", fMgrs1.GetNumOfFileMgrs())
		return
	}

	origPath = fh.AdjustPathSlash("../logTest/CmdrX/CmdrX.log")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		fmt.Printf("Error returned by (1) fh.MakeAbsolutePath(origPath). "+
			"origPath= '%v'  Error='%v'\n", origPath, err.Error())
		return
	}

	insertedFMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(origAbsPath)

	if err != nil {
		fmt.Printf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(origAbsPath). \n"+
			"origAbsPath='%v' \nError='%v' \n", origAbsPath, err.Error())
		return
	}

	err = fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5)

	if err != nil {
		fmt.Printf("Error returned by fMgrs1.InsertFileMgrAtIndex(insertedFMgr, 5) "+
			"Error='%v' \n", err.Error())
		return
	}

	numOfFileMgrs := fMgrs1.GetNumOfFileMgrs()

	for i := 0; i < numOfFileMgrs; i++ {

		xFmgr, err := fMgrs1.PeekFileMgrAtIndex(i)

		if err != nil {
			fmt.Printf("Error returned by fMgrs1.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' \n", i, err.Error())
		}

		fmt.Printf("i='%v' xFmgr='%v' \n", i, xFmgr.GetAbsolutePathFileName())
	}

	if fMgrs1.GetNumOfFileMgrs() != 11 {
		fmt.Printf("After insertion, expected fMgrs1 Array Length == 11. "+
			"Instead fMgrs1.GetNumOfDirs()=='%v'\n", fMgrs1.GetNumOfFileMgrs())
		return
	}

	fMgr5, err := fMgrs1.PeekFileMgrAtIndex(5)

	if err != nil {
		fmt.Printf("Error returned by fMgrs1.PeekFileMgrAtIndex(5). Error='%v' \n", err.Error())
		return
	}

	if !insertedFMgr.Equal(&fMgr5) {
		fmt.Printf("Error: Expected insertedFMgr == fMgr5. They WERE NOT EQUAL!\n")
	}

}
