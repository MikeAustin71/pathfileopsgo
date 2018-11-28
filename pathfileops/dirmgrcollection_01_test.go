package common

import (
	"testing"
	"strings"
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

	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != origPath {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", origPath, dMgr.Path)
	}

	if dMgr.AbsolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", origAbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 6 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 6.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
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

	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != origPath {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", origPath, dMgr.Path)
	}

	if dMgr.AbsolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", origAbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 6 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 6.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}


	dMgrs2 := DirMgrCollection{}

	// # Phase 2-2
	origPath = fh.AdjustPathSlash("../filesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-1 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs2.AddDirMgrByPathFile(origPath)

	// # Phase 2-2
	origPath = fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-2 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs2.AddDirMgrByPathFile(origPath)

	// # Phase 2-3
	origPath = fh.AdjustPathSlash("../filesfortest/oldfilesfortest")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by Phase2-3 fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs2.AddDirMgrByPathFile(origPath)

	dMgrs.AddDirMgrCollection(&dMgrs2)

	if dMgrs.GetArrayLength() != 9 {
		t.Errorf("Expected after addition - final dMgrs.GetArrayLength() == 9.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}

	dMgr2, err := dMgrs.PeekLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by Phase 2 dMgrs.PeekLastDirMgr().  Error='%v'", err.Error())
	}

	if dMgr2.Path != origPath {
		t.Errorf("Expected Last DirMgr 2 Path='%v'. Instead, dMgr2.Path='%v'", origPath, dMgr2.Path)
	}

	if dMgr2.AbsolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr 2 AbsolutePath='%v'. Instead, dMgr2.AbsolutePath='%v'", origAbsPath, dMgr2.AbsolutePath)
	}

}

func TestDirMgrCollection_PopLastDirMgr_01(t *testing.T) {

	fh := FileHelper{}
	dMgrs := DirMgrCollection{}

	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PopLastDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != origPath {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", origPath, dMgr.Path)
	}

	if dMgr.AbsolutePath != origAbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", origAbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 5 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 5.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}


}

func TestDirMgrCollection_PopFirstDirMgr_01(t *testing.T) {

	fh := FileHelper{}
	dMgrs := DirMgrCollection{}


	// # 1
	firstDirPath := fh.AdjustPathSlash("../logTest")

	origPath := firstDirPath

	origAbsPath, err := fh.MakeAbsolutePath(origPath)


	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	firstAbsDirPath := origAbsPath

	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PopFirstDirMgr()

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != firstDirPath {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", firstDirPath, dMgr.Path)
	}

	if dMgr.AbsolutePath != firstAbsDirPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", firstAbsDirPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 5 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 5.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}

}

func TestDirMgrCollection_PopDirMgrAtIndex(t *testing.T) {
	fh := FileHelper{}
	dMgrs := DirMgrCollection{}


	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)


	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}


	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	idx2Path := origPath
	idx2AbsPath := origAbsPath

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	idx3Path := origPath
	idx3AbsPath := origAbsPath

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PopDirMgrAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != idx2Path {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", idx2Path, dMgr.Path)
	}

	if dMgr.AbsolutePath != idx2AbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", idx2AbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 5 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 5.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}


	dMgr, err = dMgrs.PopDirMgrAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != idx3Path {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", idx3Path, dMgr.Path)
	}

	if dMgr.AbsolutePath != idx3AbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", idx3AbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 4 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 4.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}

}

func TestDirMgrCollection_PeekDirMgrAtIndex_01(t *testing.T) {
	fh := FileHelper{}
	dMgrs := DirMgrCollection{}


	// # 1
	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)


	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}


	dMgrs.AddDirMgrByPathFile(origPath)

	if dMgrs.GetArrayLength() != 1 {
		t.Errorf("Expected DirMgrsCollection Array Length = '1'. Instead, Array Length = '%v'", dMgrs.GetArrayLength())
	}

	if dMgrs.DirMgrs[0].AbsolutePath != origAbsPath {
		t.Errorf("Expected Addition #1 AbsolutePath='%v'. Instead, AbsolutePath='%v' ", origAbsPath, dMgrs.DirMgrs[0].AbsolutePath)
	}

	// # 2
	origPath = fh.AdjustPathSlash("../logTest/CmdrX")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 3
	origPath = fh.AdjustPathSlash("../logTest/FileMgmnt")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	idx2Path := origPath
	idx2AbsPath := origAbsPath

	dMgrs.AddDirMgrByPathFile(origPath)

	// #4
	origPath = fh.AdjustPathSlash("../logTest/FileSrc")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (3) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	idx3Path := origPath
	idx3AbsPath := origAbsPath

	dMgrs.AddDirMgrByPathFile(origPath)

	// #5
	origPath = fh.AdjustPathSlash("../logTest/Level01")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	// # 6
	origPath = fh.AdjustPathSlash("../logTest/Level01/Level02")

	origAbsPath, err = fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (4) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	dMgrs.AddDirMgrByPathFile(origPath)

	dMgr, err := dMgrs.PeekDirMgrAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != idx2Path {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", idx2Path, dMgr.Path)
	}

	if dMgr.AbsolutePath != idx2AbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", idx2AbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 6 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 6.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}


	dMgr, err = dMgrs.PeekDirMgrAtIndex(3)

	if err != nil {
		t.Errorf("Error returned by dMgrs.PeekLastDirMgr(). Error='%v'", err.Error())
	}

	if dMgr.Path != idx3Path {
		t.Errorf("Expected Last DirMgr Path='%v'. Instead, dMgr.Path='%v'", idx3Path, dMgr.Path)
	}

	if dMgr.AbsolutePath != idx3AbsPath {
		t.Errorf("Expected Last DirMgr AbsolutePath='%v'. Instead, dMgr.AbsolutePath='%v'", idx3AbsPath, dMgr.AbsolutePath)
	}

	if dMgrs.GetArrayLength() != 6 {
		t.Errorf("Expected final dMgrs.GetArrayLength() == 6.  Instead, dMgrs.GetArrayLength()=='%v'", dMgrs.GetArrayLength())
	}

}

func TestDirMgrCollection_FindDirectories_01(t *testing.T) {
	fh := FileHelper{}

	origPath := fh.AdjustPathSlash("../logTest")

	origAbsPath, err := fh.MakeAbsolutePath(origPath)

	if err != nil {
		t.Errorf("Error returned by (2) fh.MakeAbsolutePath(origPath). origPath= '%v'  Error='%v'", origPath, err.Error())
	}

	origDirMgr, err := DirMgr{}.New(origPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origPath). origPath='%v'  Error='%v'", origPath, err.Error())
	}

	if origDirMgr.AbsolutePath != origAbsPath {
		t.Errorf("Expected origDirMgr.AbsolutePath='%v'. Instead, origDirMgr.AbsolutePath='%v'", origAbsPath, origDirMgr.AbsolutePath)
	}

	fsc := FileSelectionCriteria{}
	dWlkr, err := origDirMgr.FindWalkDirFiles(fsc)

	if err != nil {
		t.Errorf("Error retured from origDirMgr.FindWalkDirFiles(fsc).  Error='%v'", err.Error())
	}

	fsc = FileSelectionCriteria{}
	fsc.FileNamePatterns = []string{"*Level*"}

	dCol, err := dWlkr.Directories.FindDirectories(fsc)

	isLevel02Found := false
	isLevel01Found := false

	for i:=0; i < dCol.GetArrayLength(); i++ {
		if strings.Contains(dCol.DirMgrs[i].DirectoryName, "Level01") {
			isLevel01Found = true
		}

		if strings.Contains(dCol.DirMgrs[i].DirectoryName, "Level02") {
			isLevel02Found = true
		}
	}

	if !isLevel01Found {
		t.Error("Expected to find a directory 'Level01'. It was NOT found!")
	}

	if !isLevel02Found {
		t.Error("Expected to find a directory 'Level02'. It was NOT found!")
	}

}