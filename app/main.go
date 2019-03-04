package main

import (
	pf "../pathfileops"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	mainTest35()
}

func mainTest35() {

	expectedStr := "Now is the time for all good men\n"

	fh := pf.FileHelper{}

	filePath := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/testRead918256.txt")

	fMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(filePath)

	if err != nil {
		fmt.Printf("Error returned from common.FileMgr{}.NewFromPathFileNameExtStr(filePath). filePath='%v'  Error='%v'", filePath, err.Error())
	}

	delim := byte('\n')

	bytes, err := fMgr.ReadFileLine(delim)

	if err != nil {
		fmt.Printf("Error returned by fMgr.ReadFileLine(delim) on Line#1. " +
			"Error='%v'", err.Error())
	}

	actualStr := string(bytes)

	err = fMgr.CloseThisFile()

	if err != nil {
		fmt.Printf("Error returned by fMgr.CloseThisFile(). Error='%v'",
			err.Error())
	}


	if  expectedStr != actualStr {

		fmt.Println("Length Expected String: ", len(expectedStr))
		fmt.Println("  Length Actual String: ", len(actualStr))
		expectedStr2 := "!" + expectedStr + "!"
		expectedStr2 = strings.Replace(expectedStr2,"\r", "*", -1)
		expectedStr2 = strings.Replace(expectedStr2,"\n", "%", -1)

		actualStr2 := "!" + actualStr + "!"
		actualStr2 = strings.Replace(actualStr2,"\r", "*", -1)
		actualStr2 = strings.Replace(actualStr2,"\n", "%", -1)

		fmt.Println("        Expected String: ", expectedStr2)
		fmt.Println("          Actual String: ", actualStr2)


	} else {
		fmt.Println("************ Success **************")
	}


}

func mainTest34() {

	textCode := "-rwxrwxrwx"

	fpCfg, err := pf.FilePermissionConfig{}.New(textCode)

	if err != nil {
		fmt.Printf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
		return
	}

	octalCode := fpCfg.GetPermissionFileModeValueText()

	fmt.Println("       Permission Codes")
	fmt.Println("----------------------------------")
	fmt.Println("Original Text Code: ", textCode)
	fmt.Printf("Octal Code: %s", octalCode)

}

func mainTest33() {

	fOpenCfg, err := pf.FileOpenConfig{}.New(
		pf.FOpenType.TypeReadWrite(),
		pf.FOpenMode.ModeCreate(),
		pf.FOpenMode.ModeExclusive())

	if err != nil {
		fmt.Printf("Error returned by FileOpenConfig{}.New(). Error='%v' \n", err.Error())
		return
	}

	openCodes := fOpenCfg.GetFileOpenNarrativeText()

	fmt.Println()
	fmt.Println("FileOpenConfig{}.GetFileOpenNarrativeText()")
	fmt.Println("Open Codes: ", openCodes)
}

func mainTest32() {

	textCode := "-rwxrwxrwx"

	fpCfg, err := pf.FilePermissionConfig{}.New(textCode)

	if err != nil {
		fmt.Printf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
			"textCode='%v' Error='%v'", textCode, err.Error())
		return
	}

	narrativeCode := fpCfg.GetPermissionNarrativeText()

	fmt.Println()
	fmt.Println("Test FilePermissionConfig{}.GetPermissionNarrativeText()")
	fmt.Println("          textCode: ", textCode)
	fmt.Println("narrativeText Code: ", narrativeCode)
}

func mainTest31() {

	fOpStatus1 := pf.FileOpenConfig{}

	fOpStatus2 := pf.FileOpenConfig{}

	fOpStatus2.CopyIn(&fOpStatus1)

	if !fOpStatus1.Equal(&fOpStatus2) {
		fmt.Println("Error: Expected fOpStatus1==fOpStatus2. THEY ARE NOT EQUAL!")
	} else {
		fmt.Println("Successful Completion!")
	}
}

func mainTest30() {

	rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest"

	dMgr, err := pf.DirMgr{}.New(rawPath)

	if err != nil {
		fmt.Printf("Error returned by DirMgr{}.New(rawPath). "+
			"rawPath='%v' Error='%v' \n\n", rawPath, err.Error())
		return
	}

	permissionText, err := dMgr.GetDirPermissionTextCodes()

	if err != nil {
		fmt.Printf("Error returned by dMgr.GetDirPermissionTextCodes(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("  Directory: ", dMgr.GetAbsolutePath())
	fmt.Println("Permissions: ", permissionText)

	fInfoPlus, err := dMgr.GetFileInfoPlus()

	if err != nil {
		fmt.Printf("Error returned by dMgr.GetFileInfoPlus(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("    FInfo Mode Str: ", fInfoPlus.Mode().String())
	fmt.Println("Finfo Mode IsDir():", fInfoPlus.Mode().IsDir())
}

func mainTest29() {

	rawPath := "D:\\gowork\\src\\MikeAustin71\\pathfileopsgo\\filesfortest\\levelfilesfortest\\level_0_3_test.txt"

	fileMgr, err := pf.FileMgr{}.NewFromPathFileNameExtStr(rawPath)

	if err != nil {
		fmt.Printf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(rawPath). "+
			"rawPath='%v' Error='%v' \n\n", rawPath, err.Error())
		return
	}

	permissionText, err := fileMgr.GetFilePermissionTextCodes()

	if err != nil {
		fmt.Printf("Error returned by fileMgr.GetFilePermissionTextCodes(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("File: ", fileMgr.GetAbsolutePathFileName())
	fmt.Println("Permissions: ", permissionText)
}

func mainTest28() {

	unixPermissions := os.FileMode(0777)

	fmt.Printf("Unix Permissions Decimal Value: %s\n",
		strconv.FormatInt(int64(unixPermissions), 10))

	fmt.Printf("Unix Permissions Octal Value: %s\n",
		strconv.FormatInt(int64(unixPermissions), 8))

	intUnixPermissions := int(unixPermissions)

	fmt.Printf("intUnixPermissions Decimal Value: %s\n",
		strconv.FormatInt(int64(intUnixPermissions), 10))

	fmt.Printf("intUnixPermissions Octal Value: %s\n",
		strconv.FormatInt(int64(intUnixPermissions), 8))

	dir := os.ModeDir | os.FileMode(0333)

	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Printf("dir Decimal Value: %s\n",
		strconv.FormatInt(int64(dir), 10))

	fmt.Printf("dir Octal Value: %s\n",
		strconv.FormatInt(int64(dir), 8))

	baseDir := dir &^ os.FileMode(0777)

	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Printf("baseDir Decimal Value: %s\n",
		strconv.FormatInt(int64(baseDir), 10))

	fmt.Printf("baseDir Octal Value: %s\n",
		strconv.FormatInt(int64(baseDir), 8))

}

func mainTest27() {

	unixPermissions := os.FileMode(0777)

	shift := uint(9)

	xPermissions := unixPermissions >> shift

	fmt.Printf("Unix Permissions Decimal Value: %s\n",
		strconv.FormatInt(int64(unixPermissions), 10))

	fmt.Printf("Unix Permissions Octal Value: %s\n",
		strconv.FormatInt(int64(unixPermissions), 8))

	fmt.Printf("shift= %d \n", shift)

	fmt.Printf("xPermissions Decimal Value: %s\n",
		strconv.FormatInt(int64(xPermissions), 10))

	fmt.Printf("xPermissions Octal Value: %s\n",
		strconv.FormatInt(int64(xPermissions), 8))

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

	fmt.Printf("                        OS MODE CONSTANTS \n")

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeDir Decimal Value: %s  \n  Octal Value : %s  \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeDir), 10),
		strconv.FormatInt(int64(os.ModeDir), 8),
		strconv.FormatInt(int64(os.ModeDir), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeAppend Decimal Value: %s  \n  Octal Value: %s  \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeAppend), 10),
		strconv.FormatInt(int64(os.ModeAppend), 8),
		strconv.FormatInt(int64(os.ModeAppend), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeExclusive Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeExclusive), 10),
		strconv.FormatInt(int64(os.ModeExclusive), 8),
		strconv.FormatInt(int64(os.ModeExclusive), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeTemporary Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeTemporary), 10),
		strconv.FormatInt(int64(os.ModeTemporary), 8),
		strconv.FormatInt(int64(os.ModeTemporary), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeSymlink Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeSymlink), 10),
		strconv.FormatInt(int64(os.ModeSymlink), 8),
		strconv.FormatInt(int64(os.ModeSymlink), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeDevice Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeDevice), 10),
		strconv.FormatInt(int64(os.ModeDevice), 8),
		strconv.FormatInt(int64(os.ModeDevice), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeNamedPipe Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeNamedPipe), 10),
		strconv.FormatInt(int64(os.ModeNamedPipe), 8),
		strconv.FormatInt(int64(os.ModeNamedPipe), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeSocket Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeSocket), 10),
		strconv.FormatInt(int64(os.ModeSocket), 8),
		strconv.FormatInt(int64(os.ModeSocket), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeSetuid Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeSetuid), 10),
		strconv.FormatInt(int64(os.ModeSetuid), 8),
		strconv.FormatInt(int64(os.ModeSetuid), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeSetgid Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeSetgid), 10),
		strconv.FormatInt(int64(os.ModeSetgid), 8),
		strconv.FormatInt(int64(os.ModeSetgid), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeCharDevice Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeCharDevice), 10),
		strconv.FormatInt(int64(os.ModeCharDevice), 8),
		strconv.FormatInt(int64(os.ModeCharDevice), 2))

	fmt.Printf("ModeSticky Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeSticky), 10),
		strconv.FormatInt(int64(os.ModeSticky), 8),
		strconv.FormatInt(int64(os.ModeSticky), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

	fmt.Printf("ModeIrregular Decimal Value: %s   \n  Octal Value: %s \n  Binary Value: %s \n\n",
		strconv.FormatInt(int64(os.ModeIrregular), 10),
		strconv.FormatInt(int64(os.ModeIrregular), 8),
		strconv.FormatInt(int64(os.ModeIrregular), 2))

	fmt.Printf("-----------------------------------------------------------------------------------\n\n")

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

	fMode, err := fPerm.GetCompositePermissionMode()

	if err != nil {
		fmt.Printf("Error returned by fPerm.GetCompositePermissionMode(). "+
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

	fOpStatus, err := pf.FileOpenConfig{}.New(pf.FOpenType.TypeWriteOnly(),
		pf.FOpenMode.ModeAppend(), pf.FOpenMode.ModeTruncate())

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
