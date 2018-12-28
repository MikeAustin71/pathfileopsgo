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

	sourceDirStr := "D:\\T10\\levelfilesfortest"
	targetDirStr := "D:\\T09\\levelfilesfortest"

	fh := pathFileOps.FileHelper{}

	srcDirPath, err := fh.GetAbsPathFromFilePath(sourceDirStr)

	if err != nil {
		fmt.Printf("Error from fh.GetAbsPathFromFilePath(sourceDirStr). "+
			"sourceDirStr='%v' Error='%v'n", sourceDirStr, err.Error())
		return
	}

	targetDirPath, err := fh.GetAbsPathFromFilePath(targetDirStr)

	if err != nil {
		fmt.Printf("Error from fh.GetAbsPathFromFilePath(targetDirStr). "+
			"targetDirStr='%v' Error='%v'n", targetDirStr, err.Error())
		return
	}

	srcDirMgr, err := pathFileOps.DirMgr{}.New(srcDirPath)

	fileSelect := pathFileOps.FileSelectionCriteria{}

	fileSelect.SelectCriterionMode = pathFileOps.ORFILESELECTCRITERION

	fileOps := make([]pathFileOps.FileOperation, 2, 5)

	fileOps[0] = pathFileOps.COPYSOURCETODESTINATIONByIo
	fileOps[1] = pathFileOps.DELETESOURCEFILE

	errStrs := srcDirMgr.ExecuteDirectoryTreeOp(fileSelect, fileOps, targetDirPath)

	lenErrStrs := len(errStrs)

	if lenErrStrs > 0 {
		fmt.Printf(" %v-Errors from ExecuteDirectoryTreeOp() \n", lenErrStrs)
		for i := 0; i < lenErrStrs; i++ {
			fmt.Printf("%v. %v \n", i, errStrs[i])
		}

		return
	}

	err = srcDirMgr.DeleteAll()

	if err != nil {
		fmt.Printf("Error returned by srcDirMgr.DeleteAll(). Error='%v'  \n", err.Error())
	}

	fmt.Println("Success ExecuteDirectoryTreeOp() Test = NO Errors!")

	return
}
