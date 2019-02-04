package main

import(
	pf "../pathfileops"
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

	testDir := "D:\\test1"

	testDirMgr, err := pf.DirMgr{}.New(testDir)

	if err != nil {
		fmt.Printf("Error returned by pf.DirMgr{}.New(testDir). " +
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("testDirMgr Path: ", testDirMgr.GetAbsolutePath())

	parentDirMgr, hasParent, err := testDirMgr.GetParentDirMgr()

	if err != nil {
		fmt.Printf("Error returned by testDirMgr.GetParentDirMgr(). " +
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("parentDirMgr Path: ", parentDirMgr.GetAbsolutePath())
	fmt.Println("  hasParent Value: ", hasParent)

	return
}
