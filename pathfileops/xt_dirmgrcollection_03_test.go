package pathfileops

import "testing"

func TestDirMgrCollection_DeleteAtIndex_01(t *testing.T) {

	d0 := "..\\dirmgrtests"
	d1 := "..\\dirmgrtests\\dir01"
	d2 := "..\\dirmgrtests\\dir01\\dir02"
	d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

	dmgrCol := DirMgrCollection{}.New()

	err := dmgrCol.AddDirMgrByPathNameStr(d0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir2 := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(d2)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d2). "+
			"Error='%v' ", err.Error())
	}

	arrayLen := dmgrCol.GetNumOfDirs()

	for i := 0; i < arrayLen; i++ {

		dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir2 = true
		}

	}

	if foundDir2 != true {
		t.Error("Expected to find dir # 2 on first pass. DID NOT FIND IT!")
		return
	}

	err = dmgrCol.DeleteAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(2) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = dmgrCol.GetNumOfDirs()

	foundDir2 = false

	for j := 0; j < arrayLen; j++ {
		dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir2 = true
		}

	}

	if foundDir2 != false {
		t.Error("Error: Found dir # 2. IT WAS NOT DELETED!")
	}

}

func TestDirMgrCollection_DeleteAtIndex_02(t *testing.T) {

	d0 := "..\\dirmgrtests"
	d1 := "..\\dirmgrtests\\dir01"
	d2 := "..\\dirmgrtests\\dir01\\dir02"
	d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

	dmgrCol := DirMgrCollection{}.New()

	err := dmgrCol.AddDirMgrByPathNameStr(d0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(d1)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d1). "+
			"d1='%v' Error='%v' ", d1, err.Error())
	}

	arrayLen := dmgrCol.GetNumOfDirs()

	for i := 0; i < arrayLen; i++ {

		dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find dir # 1 on first pass. DID NOT FIND IT!")
		return
	}

	err = dmgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(1) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = dmgrCol.GetNumOfDirs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found dir # 1. IT WAS NOT DELETED!")
	}

}

func TestDirMgrCollection_DeleteAtIndex_03(t *testing.T) {

	d0 := "..\\dirmgrtests"
	d1 := "..\\dirmgrtests\\dir01"
	d2 := "..\\dirmgrtests\\dir01\\dir02"
	d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

	dmgrCol := DirMgrCollection{}.New()

	err := dmgrCol.AddDirMgrByPathNameStr(d0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(d0)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d0). "+
			"d0='%v' Error='%v' ", d0, err.Error())
	}

	arrayLen := dmgrCol.GetNumOfDirs()

	for i := 0; i < arrayLen; i++ {

		dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find dir # 0 on first pass. DID NOT FIND IT!")
		return
	}

	err = dmgrCol.DeleteAtIndex(0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(0) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = dmgrCol.GetNumOfDirs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found dir # 0. IT WAS NOT DELETED!")
	}

}

func TestDirMgrCollection_DeleteAtIndex_04(t *testing.T) {

	d0 := "..\\dirmgrtests"
	d1 := "..\\dirmgrtests\\dir01"
	d2 := "..\\dirmgrtests\\dir01\\dir02"
	d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

	dmgrCol := DirMgrCollection{}.New()

	err := dmgrCol.AddDirMgrByPathNameStr(d0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
			"Error='%v' ", err.Error())
		return
	}

	foundDir := false
	fh := FileHelper{}
	searchStr, err := fh.GetAbsPathFromFilePath(d3)

	if err != nil {
		t.Errorf("Error returned by fh.GetAbsPathFromFilePath(d3). "+
			"d3='%v' Error='%v' ", d3, err.Error())
	}

	arrayLen := dmgrCol.GetNumOfDirs()

	for i := 0; i < arrayLen; i++ {

		dmgr, err := dmgrCol.PeekDirMgrAtIndex(i)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(i). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != true {
		t.Error("Expected to find dir # 3 on first pass. DID NOT FIND IT!")
		return
	}

	err = dmgrCol.DeleteAtIndex(3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(3) "+
			"Error='%v' ", err.Error())
	}

	arrayLen = dmgrCol.GetNumOfDirs()

	foundDir = false

	for j := 0; j < arrayLen; j++ {
		dmgr, err := dmgrCol.PeekDirMgrAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by dmgrCol.PeekDirMgrAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		if searchStr == dmgr.GetAbsolutePath() {
			foundDir = true
		}

	}

	if foundDir != false {
		t.Error("Error: Found dir # 3. IT WAS NOT DELETED!")
	}

}

func TestDirMgrCollection_DeleteAtIndex_05(t *testing.T) {

	d0 := "..\\dirmgrtests"
	d1 := "..\\dirmgrtests\\dir01"
	d2 := "..\\dirmgrtests\\dir01\\dir02"
	d3 := "..\\dirmgrtests\\dir01\\dir02\\dir03"

	dmgrCol := DirMgrCollection{}.New()

	err := dmgrCol.AddDirMgrByPathNameStr(d0)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d0). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d1). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d2). "+
			"Error='%v' ", err.Error())
		return
	}

	err = dmgrCol.AddDirMgrByPathNameStr(d3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(d3). "+
			"Error='%v' ", err.Error())
		return
	}

	arrayLen := dmgrCol.GetNumOfDirs()

	if arrayLen != 4 {
		t.Errorf("Expected array length='4'. Instead, array length='%v'",
			arrayLen)
		return
	}

	err = dmgrCol.DeleteAtIndex(3)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(3). "+
			"Error='%v' ", err.Error())
	}

	err = dmgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.DeleteAtIndex(1). "+
			"Error='%v' ", err.Error())
	}

	err = dmgrCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error: Iteration #2 returned by dmgrCol.DeleteAtIndex(1). "+
			"Error='%v' ", err.Error())
	}

	err = dmgrCol.DeleteAtIndex(0)

	if err != nil {
		t.Errorf("Error: Iteration #2 returned by dmgrCol.DeleteAtIndex(0). "+
			"Error='%v' ", err.Error())
	}

	arrayLen = dmgrCol.GetNumOfDirs()

	if arrayLen != 0 {
		t.Errorf("Error: Expected final array length=0. Instead, array length='%v'",
			arrayLen)
	}

}

func TestDirMgrCollection_GetFileMgrAtIndex_01(t *testing.T) {

	df := make([]string, 5, 10)

	df[0] = "..\\dirmgrtests"
	df[1] = "..\\dirmgrtests\\dir01"
	df[2] = "..\\dirmgrtests\\dir01\\dir02"
	df[3] = "..\\dirmgrtests\\dir01\\dir02\\dir03"
	df[4] = "..\\dirmgrtests\\dir01\\dir02\\dir03\\dir04"

	dmgrCol := DirMgrCollection{}.New()

	fh := FileHelper{}

	var err error

	for i := 0; i < 5; i++ {

		err = dmgrCol.AddDirMgrByPathNameStr(df[i])

		if err != nil {
			t.Errorf("Error returned by dmgrCol.AddDirMgrByPathNameStr(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
		}

		df[i], err = fh.MakeAbsolutePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.MakeAbsolutePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
		}

	}

	dirMgr, err := dmgrCol.GetDirMgrAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by dmgrCol.GetDirMgrAtIndex(2). "+
			"Error='%v' ", err.Error())
	}

	if df[2] != dirMgr.GetAbsolutePath() {
		t.Errorf("Error: Expected dirMgr[2]='%v'. "+
			"Instead, dirMgr[2]='%v' ", df[2], dirMgr.GetAbsolutePath())
	}

}
