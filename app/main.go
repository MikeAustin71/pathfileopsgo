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

	testDirStr := "../dirmgrtests"

	fh := pathFileOps.FileHelper{}

	dirPath, err := fh.GetAbsPathFromFilePath(testDirStr)

	if err != nil {
		fmt.Printf("Error from fh.GetAbsPathFromFilePath(testDirStr). "+
			"testDirStr='%v' Error='%v'n", testDirStr, err.Error())
		return
	}

	dMgr, err := pathFileOps.DirMgr{}.New(dirPath)

	if err != nil {
		fmt.Printf("Error from DirMgr{}.New(dirPath). "+
			"dirPath='%v' Error='%v'n", dirPath, err.Error())
		return

	}

	dMgrCol, err := dMgr.GetThisDirectoryTree()
	maxDirs := dMgrCol.GetNumOfDirs()
	fmt.Println("Returned dMgrCol Length", maxDirs)

	fmt.Println("main()")

	foundDir, err := dMgrCol.PopDirMgrAtIndex(2)

	if err != nil {
		fmt.Printf("Error from dMgrCol.PopDirMgrAtIndex(2). "+
			"Error='%v'n", err.Error())
		return
	}

	fmt.Println("Expected \\dirmgrtests\\dir01\\dir02 - Found, ", foundDir.GetAbsolutePath())

	maxDirs = dMgrCol.GetNumOfDirs()
	fmt.Println("New dMgrCol Length", maxDirs)

	for i := 0; i < maxDirs; i++ {

		foundDir, err := dMgrCol.PopFirstDirMgr()

		if err != nil {
			fmt.Printf("Error from dMgrCol.PopFirstDirMgr(). "+
				"i='%v' Error='%v'n", i, err.Error())
			return
		}

		fmt.Println(i, "  FoundDir: ", foundDir.GetAbsolutePath())
	}

	fmt.Println("Num Of Dirs In Collection: ", dMgrCol.GetNumOfDirs())
	return

}
