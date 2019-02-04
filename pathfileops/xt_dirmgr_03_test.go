package pathfileops

import "testing"

func TestDirMgr_GetParentDirMgr_01(t *testing.T) {
	fh := FileHelper{}

	origBaseAbsPath, err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02/dir03")

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02/dir03\") " +
			"Error='%v' ", err.Error())
	}

	origParentPath , err := fh.MakeAbsolutePath("../dirwalktests/dir01/dir02")

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(\"../dirwalktests/dir01/dir02\") " +
			"Error='%v' ", err.Error())
	}

	baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) " +
			"Error='%v' ", err.Error())
	}

	parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

	if err != nil {
		t.Errorf("Error returned by baseDMgr.GetParentDirMgr() " +
			"Error='%v' ", err.Error())
	}

	if origParentPath != parentDMgr.GetAbsolutePath() {
		t.Errorf("Error: Expected parentPath='%v'. Instead, parentPath='%v'. ",
			origParentPath, parentDMgr.GetAbsolutePath())
	}

	if true != hasParent {
		t.Errorf("Error: Expected hasParent='true'. Instead, hasParent='%v'.",
			hasParent)
	}

}

func TestDirMgr_GetParentDirMgr_02(t *testing.T) {

	origBaseAbsPath := "D:\\"


	baseDMgr, err := DirMgr{}.New(origBaseAbsPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath) " +
			"Error='%v' ", err.Error())
	}

	parentDMgr, hasParent, err := baseDMgr.GetParentDirMgr()

	if err != nil {
		t.Errorf("Error returned by baseDMgr.GetParentDirMgr() " +
			"Error='%v' ", err.Error())
	}


	isEqual := baseDMgr.Equal(&parentDMgr)

	if isEqual != true {
		t.Error("Error: Expected baseDMgr==parentDMgr. THEY ARE NOT EQUAL!")
	}

	if false != hasParent {
		t.Errorf("Error: Expected hasParent='false'. Instead, hasParent='%v'.",
			hasParent)
	}

}

func TestDirMgr_GetNumberOfAbsPathElements_01(t *testing.T) {

	origBaseAbsPath := "D:\\dir01\\dir02\\dir03\\dir04"

	dMgr, err := DirMgr{}.New(origBaseAbsPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
			"Error='%v' ", err.Error())
	}

	numOfElements := dMgr.GetNumberOfAbsPathElements()

	if 5 != numOfElements {
		t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
			"number of elements='%v' ", 5, numOfElements)
	}

}

func TestDirMgr_GetNumberOfAbsPathElements_02(t *testing.T) {

	origBaseAbsPath := "D:\\"

	dMgr, err := DirMgr{}.New(origBaseAbsPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
			"Error='%v' ", err.Error())
	}

	numOfElements := dMgr.GetNumberOfAbsPathElements()

	if 1 != numOfElements {
		t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
			"number of elements='%v' ", 1, numOfElements)
	}

}

func TestDirMgr_GetNumberOfAbsPathElements_03(t *testing.T) {

	origBaseAbsPath := "D:\\test01"

	dMgr, err := DirMgr{}.New(origBaseAbsPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(origBaseAbsPath). " +
			"Error='%v' ", err.Error())
	}

	numOfElements := dMgr.GetNumberOfAbsPathElements()

	if 2 != numOfElements {
		t.Errorf("Expected Number Of directory elements='%v'. Instead, " +
			"number of elements='%v' ", 2, numOfElements)
	}

}