package main

import (
	pf "../pathfileops"
	"fmt"
	"os"
	"strconv"
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

	mainTest25()

}

func mainTest26() {

	unixPermissions := os.FileMode(0777)

	fmt.Printf("unixPermissions Decimal: %d  Octal: %s Binary: %s String Value:%x \n",
		unixPermissions, strconv.FormatInt(int64(unixPermissions), 8),
		strconv.FormatInt(int64(unixPermissions), 2), unixPermissions.String())

	fmt.Printf("ModeDir Decimal: %d  Octal: %s Binary: %s String Value: %s \n",
		os.ModeDir, strconv.FormatInt(int64(os.ModeDir), 8),
		strconv.FormatInt(int64(os.ModeDir), 2), os.ModeDir.String())

	directoryPermission := os.ModeDir | unixPermissions

	fmt.Printf("directoryPermission Decimal: %d  Octal: %s Binary: %s \n",
		directoryPermission, strconv.FormatInt(int64(directoryPermission), 8),
		strconv.FormatInt(int64(directoryPermission), 2))
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("   dirctoryPermission string value: %s\n", directoryPermission.String())
	fmt.Printf("    dirctoryPermission IsDir value: %v\n", directoryPermission.IsDir())
	fmt.Printf("dirctoryPermission IsRegular value: %v\n", directoryPermission.IsRegular())

}

func mainTest25() {

	fmt.Printf("ModeDir Decimal: %v  Octal: %s Binary: %s \n",
		os.ModeDir, strconv.FormatInt(int64(os.ModeDir), 8),
		strconv.FormatInt(int64(os.ModeDir), 2))

	fmt.Printf("ModeAppend Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeAppend, os.ModeAppend, os.ModeAppend)

	fmt.Printf("ModeExclusive Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeExclusive, os.ModeExclusive, os.ModeExclusive)

	fmt.Printf("ModeTemporary Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeTemporary, os.ModeTemporary, os.ModeTemporary)

	fmt.Printf("ModeSymlink Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeSymlink, os.ModeSymlink, os.ModeSymlink)

	fmt.Printf("ModeDevice Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeDevice, os.ModeDevice, os.ModeDevice)

	fmt.Printf("ModeNamedPipe Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeNamedPipe, os.ModeNamedPipe, os.ModeNamedPipe)

	fmt.Printf("ModeSocket Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeSocket, os.ModeSocket, os.ModeSocket)

	fmt.Printf("ModeSetuid Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeSetuid, os.ModeSetuid, os.ModeSetuid)

	fmt.Printf("ModeSetgid Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeSetgid, os.ModeSetgid, os.ModeSetgid)

	fmt.Printf("ModeCharDevice Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeCharDevice, os.ModeCharDevice, os.ModeCharDevice)

	fmt.Printf("ModeSticky Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeSticky, os.ModeSticky, os.ModeSticky)

	fmt.Printf("ModeIrregular Decimal: %d  Octal: %o Binary: %b \n",
		os.ModeIrregular, os.ModeIrregular, os.ModeIrregular)

}

func mainTest24() {

	fPerm := pf.FilePermissionConfig{}

	expectedDecimalModeValue := 511
	expectedOctalModeValue :=
		pf.FileHelper{}.ConvertDecimalToOctal(expectedDecimalModeValue)

	modeStr := "-rwxrwxrwx"

	err := fPerm.SetFileModeByTextCode(modeStr)

	if err != nil {
		fmt.Printf("Error returned by fPerm.StringToMode(modeStr). "+
			"modeStr='%v' Error='%v' \n", modeStr, err.Error())
		return
	}

	fMode, err := fPerm.GetFileMode()

	if err != nil {
		fmt.Printf("Error returned by fPerm.GetFileMode(). "+
			"modeStr='%v' Error='%v' \n", modeStr, err.Error())
		return
	}

	actualDecimalModeValue := int(fMode)

	actualOctalModeValue :=
		pf.FileHelper{}.ConvertDecimalToOctal(actualDecimalModeValue)

	fmt.Println("--- StringToMode Results ----")
	fmt.Println("          Original Mode Str: ", modeStr)
	fmt.Println("            Actual Mode Str: ", fMode.String())
	fmt.Println("Expected Decimal Mode Value: ", expectedDecimalModeValue)
	fmt.Println("  Actual Mode Decimal Value: ", actualDecimalModeValue)
	fmt.Println("  Expected Octal Mode Value: ", expectedOctalModeValue)
	fmt.Println("    Actual Octal Mode Value: ", actualOctalModeValue)
}

func mainTest23() {

	dfm := os.FileMode(os.ModeDir)

	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("File Mode String %s\n", dfm.String())
	tfm := os.FileMode(0777)
	fmt.Printf("4-digit 777 File Mode String %s\n", tfm.String())

	nfm := tfm | dfm

	fmt.Printf("tfm or'd with dfm  %s \n", nfm.String())
	fh := pf.FileHelper{}

	mode := fh.ConvertOctalToDecimal(777)
	fmt.Printf("mode = %d\n", mode)
	decimalEquivalent := fh.ConvertDecimalToOctal(mode)
	tfm = os.FileMode(mode)
	fmt.Printf("3-digit 777 File Mode String %s\n", tfm.String())
	fmt.Printf("Decimal Equivalent %d \n", decimalEquivalent)
}

func mainTest22() {

	expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

	fOpStatus, err := pf.FileOpenConfig{}.New(pf.FOpenType.WriteOnly(),
		pf.FOpenMode.Append(), pf.FOpenMode.Truncate())

	if err != nil {
		fmt.Printf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
		return
	}

	actualFileOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

	if err != nil {
		fmt.Printf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode(). Error='%v' \n", err.Error())
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
