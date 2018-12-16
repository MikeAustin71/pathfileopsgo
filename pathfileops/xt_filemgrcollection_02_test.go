package pathfileops

import "testing"

func TestFileMgrCollection_DeleteAtIndex_01(t *testing.T) {

	f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
	f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
	f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
	f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

	fMgrCol := FileMgrCollection{}.New()

	err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

	if err != nil {
		t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(f2)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f2). "+
			"f2='%v'  Error='%v' ", f2, err.Error())
	}

	arrayLen := fMgrCol.GetNumOfFileMgrs()

	for i := 0; i < arrayLen; i++ {

		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find file index # 2 on first pass. DID NOT FIND IT!")
		return
	}

	err = fMgrCol.DeleteAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(2) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fMgrCol.GetNumOfFileMgrs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found file at index # 2. IT WAS NOT DELETED!")
	}

}

func TestFileMgrCollection_DeleteAtIndex_02(t *testing.T) {

	f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
	f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
	f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
	f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

	fMgrCol := FileMgrCollection{}.New()

	err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

	if err != nil {
		t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(f1)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f1). "+
			"f1='%v'  Error='%v' ", f1, err.Error())
	}

	arrayLen := fMgrCol.GetNumOfFileMgrs()

	for i := 0; i < arrayLen; i++ {

		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find file index # 1 on first pass. DID NOT FIND IT!")
		return
	}

	err = fMgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(1) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fMgrCol.GetNumOfFileMgrs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found file at index # 1. IT WAS NOT DELETED!")
	}

}

func TestFileMgrCollection_DeleteAtIndex_03(t *testing.T) {

	f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
	f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
	f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
	f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

	fMgrCol := FileMgrCollection{}.New()

	err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

	if err != nil {
		t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(f0)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f0). "+
			"f0='%v'  Error='%v' ", f0, err.Error())
	}

	arrayLen := fMgrCol.GetNumOfFileMgrs()

	for i := 0; i < arrayLen; i++ {

		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find file index # 0 on first pass. DID NOT FIND IT!")
		return
	}

	err = fMgrCol.DeleteAtIndex(0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fMgrCol.GetNumOfFileMgrs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found file at index # 0. IT WAS NOT DELETED!")
	}

}

func TestFileMgrCollection_DeleteAtIndex_04(t *testing.T) {

	f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
	f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
	f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
	f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

	fMgrCol := FileMgrCollection{}.New()

	err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

	if err != nil {
		t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(f3)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(f3). "+
			"f3='%v'  Error='%v' ", f3, err.Error())
	}

	arrayLen := fMgrCol.GetNumOfFileMgrs()

	for i := 0; i < arrayLen; i++ {

		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find file index # 3 on first pass. DID NOT FIND IT!")
		return
	}

	err = fMgrCol.DeleteAtIndex(3)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fMgrCol.GetNumOfFileMgrs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		fileMgr, err := fMgrCol.PeekFileMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fMgrCol.PeekFileMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == fileMgr.GetAbsolutePathFileName() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found file at index # 3. IT WAS NOT DELETED!")
	}

}

func TestFileMgrCollection_DeleteAtIndex_05(t *testing.T) {

	f0 := "..\\dirmgrtests\\dir01\\level_1_1_test.txt"
	f1 := "..\\dirmgrtests\\dir01\\level_1_2_test.txt"
	f2 := "..\\dirmgrtests\\dir01\\level_1_3_test.txt"
	f3 := "..\\dirmgrtests\\dir01\\level_1_4_test.txt"

	fMgrCol := FileMgrCollection{}.New()

	err := fMgrCol.AddFileMgrByPathFileNameExt(f0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.AddFileMgrByPathFileNameExt(f2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = fMgrCol.AddFileMgrByPathFileNameExt(f3)

	if err != nil {
		t.Errorf("Error returned by AddFileMgrByPathFileNameExt(f3). "+
			"Error='%v' ", err.Error())
		return
	}

	arrayLen := fMgrCol.GetNumOfFileMgrs()

	if arrayLen != 4 {
		t.Errorf("Error: Expected intial array length='4'. Instead, array length='%v'",
			arrayLen)
	}

	err = fMgrCol.DeleteAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(2). "+
			"Error='%v' ", err.Error())
	}

	err = fMgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(1). "+
			"Error='%v' ", err.Error())
	}

	err = fMgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by 2nd Pass fMgrCol.DeleteAtIndex(1). "+
			"Error='%v' ", err.Error())
	}

	err = fMgrCol.DeleteAtIndex(0)

	if err != nil {
		t.Errorf("Error returned by fMgrCol.DeleteAtIndex(0). "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fMgrCol.GetNumOfFileMgrs()

	if arrayLen != 0 {
		t.Errorf("Error: Expected final array length='0'.  "+
			"Instead, final array length='%v' ", arrayLen)
	}

}
