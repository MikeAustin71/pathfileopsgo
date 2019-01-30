package main

import (
	pathFileOps "../pathfileops"
	"fmt"
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

	targetDirStr := "D:/TimeZoneDb06/zoneinfo"

	targetDirMgr, err := pathFileOps.DirMgr{}.New(targetDirStr)

	if err != nil {
		fmt.Printf("Error returned by pathFileOps.DirMgr{}.New(targetDirStr). "+
			"targetDirStr='%v' Error='%v' \n", targetDirStr, err.Error())
		return
	}

	fileSelect := pathFileOps.FileSelectionCriteria{}

	//fileSelect.SelectCriterionMode = pathFileOps.ORFILESELECTCRITERION

	dirInfo, err := targetDirMgr.FindWalkDirFiles(fileSelect)

	if err != nil {
		fmt.Printf("Error returned by targetDirMgr.FindWalkDirFiles(fileSelect) " +
			"Error='%v' \n", err.Error())
		return
	}

	numOfFiles := dirInfo.FoundFiles.GetNumOfFileMgrs()

	fmt.Println("                  Found Files")
	fmt.Println("==================================================")
	fmt.Println()
	for i:=0; i < numOfFiles; i++ {

		fMgr, err := dirInfo.FoundFiles.PeekFileMgrAtIndex(i)

		if err != nil {
			fmt.Printf("Found Files Error: Index='%v' Error='%v' \n", i, err.Error())
			return
		}

		fmt.Printf("%03d File: %v \n", i, fMgr.GetAbsolutePathFileName())
	}

	return
}
