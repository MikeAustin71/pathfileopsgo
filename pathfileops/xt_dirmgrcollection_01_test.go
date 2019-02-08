package pathfileops

import (
	"testing"
)

func TestDirMgrCollection_AddDirMgr_01(t *testing.T) {

	fh := FileHelper{}
	dMgrs := DirMgrCollection{}

	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	if dMgrs.GetNumOfDirs() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
	}

	if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	dMgr, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.path != origPath {
		t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
	}

	if dMgr.absolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
	}

	if dMgrs.GetNumOfDirs() != 6 {
		t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
	}

}

func TestDirMgrCollection_AddDirMgrCollection(t *testing.T) {

	fh := FileHelper{}
	dMgrs := DirMgrCollection{}

	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	if dMgrs.GetNumOfDirs() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
	}

	if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	dMgr, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.path != origPath {
		t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
	}

	if dMgr.absolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
	}

	if dMgrs.GetNumOfDirs() != 6 {
		t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
	}

	dMgrs2 := DirMgrCollection{}

	// # Phase 2-2
	origPath = fh.AdjustPathSlash("../filesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-1 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs2.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # Phase 2-2
	origPath = fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-2 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs2.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # Phase 2-3
	origPath = fh.AdjustPathSlash("../filesfortest/oldfilesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-3 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs2.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	dMgrs.AddDirMgrCollection(&dMgrs2)

	if dMgrs.GetNumOfDirs() != 9 {
		t.Errorf("Expected after addition - final dMgrs.GetNumOfDirs() == 9.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
	}

	dMgr2, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by Phase 2 dMgrs.PeekLastDirMgr().  Error='%v'", err.Error())
	}

	if dMgr2.path != origPath {
		t.Errorf("Expected Last DirMgr 2 path='%v'. Instead, dMgr2.path='%v'", origPath, dMgr2.path)
	}

	if dMgr2.absolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr 2 absolutePath='%v'. Instead, dMgr2.absolutePath='%v'", origAbsPath, dMgr2.absolutePath)
	}

}

func TestDirMgrCollection_GetDirMgrArray_01(t *testing.T) {

	fh := FileHelper{}
	dMgrs := DirMgrCollection{}

	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	if dMgrs.GetNumOfDirs() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetNumOfDirs())
	}

	if dMgrs.dirMgrs[0].absolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 absolutePath='%v'. Instead, absolutePath='%v' ", origAbsPath, dMgrs.dirMgrs[0].absolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	err = dMgrs.AddDirMgrByPathNameStr(origPath)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	dMgr, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.path != origPath {
		t.Errorf("Expected Last DirMgr path='%v'. Instead, dMgr.path='%v'", origPath, dMgr.path)
	}

	if dMgr.absolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr absolutePath='%v'. Instead, dMgr.absolutePath='%v'", origAbsPath, dMgr.absolutePath)
	}

	if dMgrs.GetNumOfDirs() != 6 {
		t.Errorf("Expected final dMgrs.GetNumOfDirs() == 6.  Instead, dMgrs.GetNumOfDirs()=='%v'", dMgrs.GetNumOfDirs())
	}

	cntr := 0

	for _, dirMgr := range dMgrs.GetDirMgrArray() {

		err := dirMgr.IsDirMgrValid("TestDirMgrCollection_GetDirMgrArray_01 Error")

		if err != nil {
			t.Errorf("Invalid DirMgr: Dir='%v'  Error=%v", dirMgr.GetAbsolutePath(), err.Error())
		}

		cntr++
	}

	if cntr != 6 {
		t.Errorf("Expected Diretory Count='6'. Instead Directory Count='%v'", cntr)
	}

}
