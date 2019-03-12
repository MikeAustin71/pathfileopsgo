package pathfileops

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFileMgr_Empty_01(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by filePath1, err := fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.NewFromPathFileNameExtStr(filePath1). "+
			"filePath1='%v' Error='%v'", filePath1, err.Error())
	}

	fileMgr1.Empty()

	fileMgr2 := FileMgr{}
	fileMgr2.Empty()

	if !fileMgr1.Equal(&fileMgr2) {
		t.Error("Error: Expected empty fileMgr1 to equal empty fileMgr2. " +
			"However, THEY ARE NOT EQUAL!")
	}

}

func TestFileMgr_Equal_01(t *testing.T) {
	fh := FileHelper{}

	relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by filePath1, err := fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.NewFromPathFileNameExtStr(filePath1). "+
			"filePath1='%v' Error='%v'", filePath1, err.Error())
	}

	fileMgr2 := fileMgr1.CopyOut()

	if fileMgr2.Equal(&fileMgr1) != true {
		t.Error("Expected Equal to return 'true' for fileMgr1==fileMgr1, instead got: ", "false")
	}

}

func TestFileMgr_Equal_02(t *testing.T) {
	fh := FileHelper{}

	relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
	}

	relPath2 := "..\\logTest\\FileMgmnt\\TestFile003.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr2.Equal(&fileMgr1) != false {
		t.Error("Expected Equal to return 'false' for fileMgr1==fileMgr2, instead got: 'true'")
	}

}

func TestFileMgr_EqualAbsPaths_01(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_1_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualAbsPaths(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path to equal fileMgr2 absolute path. "+
			"Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualAbsPaths_02(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\level_1_1_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualAbsPaths(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path to equal fileMgr2 absolute path. "+
			"Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualAbsPaths_03(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_2_0_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualAbsPaths(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path to NOT fileMgr2 absolute path. "+
			"Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualFileNameExt_01(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
			"They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
	}

}

func TestFileMgr_EqualFileNameExt_02(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\LEVEL_1_0_TEST.TXT"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
			"They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
	}

}

func TestFileMgr_EqualFileNameExt_03(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_02_dir\\LEVEL_1_0_TEST.TXT"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 file name ext to equal fileMgr2 file name ext. "+
			"They ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
	}

}

func TestFileMgr_EqualFileNameExt_04(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_2_2_xray.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 file name ext to NOT equal fileMgr2 file name ext. "+
			"However, they ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
	}

}

func TestFileMgr_EqualFileNameExt_05(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.jag"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 file name ext to NOT equal fileMgr2 file name ext. "+
			"However, they ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetFileNameExt(), fileMgr2.GetFileNameExt())
	}

}

func TestFileMgr_EqualPathFileNameExt_01(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path file name ext to equal "+
			"fileMgr2 absolute path file name ext. "+
			"Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualPathFileNameExt_02(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\FILESFORTEST\\LEVELFILESFORTEST\\LEVEL_01_DIR\\LEVEL_1_0_TEST.TXT"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if !fileMgr1.EqualPathFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path file name ext to equal fileMgr2 "+
			"absolute path file name ext. Paths ARE NOT EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualPathFileNameExt_03(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_02_dir\\level_1_0_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
			"absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualPathFileNameExt_04(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_X_0_test.txt"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
			"absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_EqualPathFileNameExt_05(t *testing.T) {

	fh := FileHelper{}

	relPath1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.txt"
	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). "+
			"filePath1='%v' Error='%v' ", filePath1, err.Error())
	}

	relPath2 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_1_0_test.log"
	filePath2, err := fh.MakeAbsolutePath(relPath2)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath2). "+
			"relPath2='%v' Error='%v'", relPath2, err.Error())
	}

	fileMgr2, err := FileMgr{}.New(filePath2)

	if err != nil {
		t.Errorf("Error returned by FileMgr{}.New(filePath2). "+
			"filePath2='%v' Error='%v'", filePath2, err.Error())
	}

	if fileMgr1.EqualPathFileNameExt(&fileMgr2) {
		t.Errorf("Error: expected fileMgr1 absolute path file name ext to NOT equal fileMgr2 "+
			"absolute path file name ext. Paths ARE EQUAL! \n fileMgr1='%v' \n fileMgr2='%v'\n",
			fileMgr1.GetAbsolutePath(), fileMgr2.GetAbsolutePath())
	}

}

func TestFileMgr_GetAbsolutePathFileName_01(t *testing.T) {
	fh := FileHelper{}
	relPath1 := "..\\logTest\\CmdrX\\CmdrX.log"

	filePath1, err := fh.MakeAbsolutePath(relPath1)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(relPath1). "+
			"relPath1='%v' Error='%v'", relPath1, err.Error())
	}

	fileMgr1, err := FileMgr{}.New(filePath1)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(filePath1). Error='%v' ", err.Error())
	}

	absPathFileName := fileMgr1.GetAbsolutePathFileName()
	absPathFileName = strings.ToLower(absPathFileName)
	filePath1 = strings.ToLower(filePath1)

	if filePath1 != absPathFileName {
		t.Errorf("Error: Expected absPathFileName='%v'. Instead, absPathFileName='%v' ",
			filePath1, absPathFileName)
	}

}

func TestFileMgr_GetBufioReader_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
			"adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
			"dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	err = srcFMgr.OpenThisFileReadOnly()

	if err != nil {
		t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly(). "+
			"FileName:'%v' Error='%v' ", srcFMgr.GetAbsolutePathFileName(), err.Error())
	}

	bufReader := srcFMgr.GetBufioReader()

	if bufReader == nil {
		t.Error("Error: Expected pointer return from srcFMgr.GetBufioReader(). " +
			"Pointer IS NIL!")
	}

	err = srcFMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CloseThisFile(). "+
			"Error='%v' ", err.Error())
	}
}

func TestFileMgr_GetBufioWriter_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
			"adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
			"dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	err = srcFMgr.OpenThisFileWriteOnlyAppend()

	if err != nil {
		t.Errorf("Error returned by srcFMgrOpenThisFileWriteOnlyAppend(). "+
			"FileName:'%v' Error='%v' ", srcFMgr.GetAbsolutePathFileName(), err.Error())
	}

	bufReader := srcFMgr.GetBufioWriter()

	if bufReader == nil {
		t.Error("Error: Expected pointer return from srcFMgr.GetBufioWriter(). " +
			"Pointer IS NIL!")
	}

	err = srcFMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Error returned from srcFMgr.CloseThisFile(). "+
			"Error='%v' ", err.Error())
	}

}

func TestFileMgr_GetDirMgr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
			"adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
			"dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	expectedDirMgrPath := strings.ToLower(dMgr.GetAbsolutePath())

	srcDMgr := srcFMgr.GetDirMgr()

	actualDirMgrPath := strings.ToLower(srcDMgr.GetAbsolutePath())

	if expectedDirMgrPath != actualDirMgrPath {
		t.Errorf("Error: Expected returned directory path='%v'. Instead, "+
			"returned directory path='%v' ",
			expectedDirMgrPath, actualDirMgrPath)
	}

}

func TestFileMgr_GetFileExt(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). "+
			"adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt). "+
			"dMgr.absolutePath='%v' expectedFileNameExt='%v'  Error='%v'", dMgr.absolutePath, adjustedPath, err.Error())
	}

	expectedFileExt := ".txt"

	actualFileExt := srcFMgr.GetFileExt()

	if expectedFileExt != actualFileExt {
		t.Errorf("Error: Expected returned file extension='%v'. Instead "+
			"returned file extension='%v' ",
			expectedFileExt, actualFileExt)
	}

}

func TestFileMgr_GetFileInfo_01(t *testing.T) {

	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	fInfo, err := srcFMgr.GetFileInfo()

	expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

	actualFileNameExt := strings.ToLower(fInfo.Name())

	if expectedFileNameExt != actualFileNameExt {
		t.Errorf("Error: Expected File Name:='%v'.  Instead, File Name='%v'",
			expectedFileNameExt, actualFileNameExt)
	}
}

func TestFileMgr_GetFileInfo_02(t *testing.T) {

	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	_, err = srcFMgr.GetFileInfo()

	if err == nil {
		t.Error("Error expected error return from srcFMgr.GetFileInfo() because " +
			"file does not exist. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_GetFileInfoPlus_01(t *testing.T) {

	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	fInfoPlus, err := srcFMgr.GetFileInfoPlus()

	expectedFileNameExt := strings.ToLower(srcFMgr.GetFileNameExt())

	actualFileNameExt := strings.ToLower(fInfoPlus.Name())

	if expectedFileNameExt != actualFileNameExt {
		t.Errorf("Error: Expected File Name:='%v'.  Instead, File Name='%v'",
			expectedFileNameExt, actualFileNameExt)
	}

}

func TestFileMgr_GetFileInfoPlus_02(t *testing.T) {

	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/iDoNotExist_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	_, err = srcFMgr.GetFileInfoPlus()

	if err == nil {
		t.Error("Error expected error return from srcFMgr.GetFileInfoPlus() because " +
			"file does not exist. However, NO ERROR WAS RETURNED!")
	}
}

func TestFileMgr_GetFileName_01(t *testing.T) {
	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	expectedFileName := "newerFileForTest_01"

	actualFileName := srcFMgr.GetFileName()

	if expectedFileName != actualFileName {
		t.Errorf("Error: Expected File Name='%v'. Instead, actual File Name='%v'",
			expectedFileName, actualFileName)
	}

}

func TestFileMgr_GetFileNameExt_01(t *testing.T) {
	fh := FileHelper{}
	targetFile := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	absPath, err := fh.MakeAbsolutePath(targetFile)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr, err := FileMgr{}.New(absPath)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(absPath). "+
			"absPath='%v' Error='%v'", absPath, err.Error())
	}

	expectedFileNameExt := "newerFileForTest_01.txt"

	actualFileNameExt := srcFMgr.GetFileNameExt()

	if expectedFileNameExt != actualFileNameExt {
		t.Errorf("Error: Expected File Name Ext='%v'. Instead, actual File Name Ext='%v'",
			expectedFileNameExt, actualFileNameExt)
	}

}

func TestFileMgr_GetFilePermissionTextCodes_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	err = srcFMgr.OpenThisFileReadWrite()

	if err != nil {

		_ = srcFMgr.CloseThisFile()

		t.Errorf("Error returned by srcFMgr.OpenThisFileReadWrite(). "+
			"Error='%v' ", err.Error())

	}

	expectedPermissionCodes := "-rw-rw-rw-"

	actualPermissionTextCodes, err := srcFMgr.GetFilePermissionTextCodes()

	if err != nil {
		_ = srcFMgr.CloseThisFile()
		t.Errorf("Error returned by srcFMgr.GetFilePermissionTextCodes(). Error='%v'",
			err.Error())
	}

	err = srcFMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Error returned by #2 srcFMgr.CloseThisFile(). "+
			"Error='%v'", err.Error())
	}

	if expectedPermissionCodes != actualPermissionTextCodes {
		t.Errorf("Error: Expected Permission Code='%v'. Instead, Permission Code='%v'",
			expectedPermissionCodes, actualPermissionTextCodes)
	}

}

func TestFileMgr_GetFilePermissionTextCodes_02(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDoNotExist_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	_, err = srcFMgr.GetFilePermissionTextCodes()

	if err == nil {
		t.Error("Expected error return from srcFMgr.GetFilePermissionTextCodes() " +
			"because file does not exist. However, NO ERROR WAS RETURNED!")
	}

}

func TestFileMgr_GetFilePtr_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	err = srcFMgr.OpenThisFileReadOnly()

	if err != nil {

		_ = srcFMgr.CloseThisFile()

		t.Errorf("Error returned by srcFMgr.OpenThisFileReadOnly(). "+
			"Error='%v' ", err.Error())

	}

	fPtr := srcFMgr.GetFilePtr()

	if fPtr == nil {
		t.Error("Error: Expected a populated file pointer. However, the file pointer is nil!")
	}

	err = srcFMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Error returned by final srcFMgr.CloseThisFile(). Error='%v' ",
			err.Error())
	}

}

func TestFileMgr_GetFileSize_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	actualFileSize := srcFMgr.GetFileSize()

	expectedFileSize := int64(29)

	if expectedFileSize != actualFileSize {
		t.Errorf("Expected file size='29'. Instead, file size='%v'",
			actualFileSize)
	}

}

func TestFileMgr_GetFileSize_02(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/iDontExist_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	actualFileSize := srcFMgr.GetFileSize()

	expectedFileSize := int64(-1)

	if expectedFileSize != actualFileSize {
		t.Errorf("Expected file size='-1'. Instead, file size='%v'",
			actualFileSize)
	}

}

func TestFileMgr_GetOriginalPathFileName_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	originalPathFileName := srcFMgr.GetOriginalPathFileName()

	if targetFile != originalPathFileName {
		t.Errorf("Error: Expected Original Path and File Name='%v'. Instead, "+
			"Original Path and File Name='%v'",
			targetFile, originalPathFileName)
	}

}

func TestFileMgr_IsAbsolutePathFileNamePopulated_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

	if !isAbsPathFileName {
		t.Error("Error: Expected Absolute Path File Name to be populated. " +
			"It was NOT!")
	}

}

func TestFileMgr_IsAbsolutePathFileNamePopulated_02(t *testing.T) {

	srcFMgr := FileMgr{}

	isAbsPathFileName := srcFMgr.IsAbsolutePathFileNamePopulated()

	if isAbsPathFileName {
		t.Error("Error: Expected Absolute Path File Name NOT populated. " +
			"WRONG - It IS populated!")
	}

}

func TestFileMgr_IsFileExtPopulated_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isFileExtPopulated := srcFMgr.IsFileExtPopulated()

	if !isFileExtPopulated {
		t.Error("Expected srcFMgr.IsFileExtPopulated() == 'true'. Instead, it is 'false'")
	}

}

func TestFileMgr_IsFileExtPopulated_02(t *testing.T) {

	fh := FileHelper{}

	targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
	targetFile := "newerFileForTest_01"

	dirMgr, err := DirMgr{}.New(targetDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
			"targetDir='%v' Error='%v'", targetDir, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

	if err != nil {
		t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
			"DirMgr='%v' targetFile='%v' Error='%v'",
			dirMgr.GetAbsolutePath(), targetFile, err.Error())
	}

	isFileExtPopulated := srcFMgr.IsFileExtPopulated()

	if isFileExtPopulated {
		t.Error("Expected srcFMgr.IsFileExtPopulated() == 'false'. Instead, it is 'true'")
	}

}

func TestFileMgr_IsFileNameExtPopulated_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isFileNameExtPopulated := srcFMgr.IsFileNameExtPopulated()

	if !isFileNameExtPopulated {
		t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'true'. Instead, it is 'false'")
	}

}

func TestFileMgr_IsFileNameExtPopulated_02(t *testing.T) {
	fh := FileHelper{}

	targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
	targetFile := "newerFileForTest_01"

	dirMgr, err := DirMgr{}.New(targetDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
			"targetDir='%v' Error='%v'", targetDir, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

	if err != nil {
		t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
			"DirMgr='%v' targetFile='%v' Error='%v'",
			dirMgr.GetAbsolutePath(), targetFile, err.Error())
	}

	isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

	if isFileNamePopulated {

		t.Errorf("Expected srcFMgr.IsFileNameExtPopulated() == 'false'. Instead, it is 'true'. "+
			"FileName='%v' File Extension='%v' Len File Ext= '%v' ",
			srcFMgr.GetFileName(), srcFMgr.GetFileExt(), len(srcFMgr.GetFileExt()))
	}

}

func TestFileMgr_IsFileNameExtPopulated_03(t *testing.T) {

	srcFMgr := FileMgr{}

	isFileNamePopulated := srcFMgr.IsFileNameExtPopulated()

	if isFileNamePopulated {
		t.Error("Expected srcFMgr.IsFileNameExtPopulated() == 'false'. Instead, it is 'true'")
	}

}

func TestFileMgr_IsFileNamePopulated_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isFileNamePopulated := srcFMgr.IsFileNamePopulated()

	if !isFileNamePopulated {
		t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'. Instead, it is 'false'")
	}
}

func TestFileMgr_IsFileNamePopulated_02(t *testing.T) {

	fh := FileHelper{}

	targetDir := fh.AdjustPathSlash("../filesfortest/newfilesfortest")
	targetFile := "newerFileForTest_01"

	dirMgr, err := DirMgr{}.New(targetDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(targetDir). "+
			"targetDir='%v' Error='%v'", targetDir, err.Error())
	}

	srcFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile)

	if err != nil {
		t.Errorf("FileMgr{}.NewFromDirMgrFileNameExt(dirMgr, targetFile). "+
			"DirMgr='%v' targetFile='%v' Error='%v'",
			dirMgr.GetAbsolutePath(), targetFile, err.Error())
	}

	isFileNamePopulated := srcFMgr.IsFileNamePopulated()

	if !isFileNamePopulated {
		t.Error("Expected srcFMgr.IsFileNamePopulated() == 'true'. Instead, it is 'false'")
	}

}

func TestFileMgr_IsFileNamePopulated_03(t *testing.T) {

	srcFMgr := FileMgr{}

	isFileNamePopulated := srcFMgr.IsFileNamePopulated()

	if isFileNamePopulated {
		t.Error("Expected srcFMgr.IsFileNamePopulated() == 'false'. Instead, it is 'true'")
	}

}

func TestFileMgr_IsFilePointerOpen_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	err = srcFMgr.OpenThisFileReadOnly()

	if err != nil {
		_ = srcFMgr.CloseThisFile()
		t.Errorf("Error returned from srcFMgr.OpenThisFileReadOnly(). "+
			"Error='%v'", err.Error())
	}

	isFilePointerOpen := srcFMgr.IsFilePointerOpen()

	err = srcFMgr.CloseThisFile()

	if err != nil {
		t.Errorf("Error returned from final srcFMgr.CloseThisFile(). "+
			"Error='%v'", err.Error())
	}

	if !isFilePointerOpen {
		t.Error("Expected isFilePointerOpen = 'true'. Instead, it is FALSE!")
	}

}

func TestFileMgr_IsFilePointerOpen_02(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isFilePointerOpen := srcFMgr.IsFilePointerOpen()

	if isFilePointerOpen {
		t.Error("Expected isFilePointerOpen = 'false'. Instead, it is TRUE!")
	}

}

func TestFileMgr_IsFilePointerOpen_03(t *testing.T) {

	srcFMgr := FileMgr{}

	isFilePointerOpen := srcFMgr.IsFilePointerOpen()

	if isFilePointerOpen {
		t.Error("Expected isFilePointerOpen = 'false'. Instead, it is TRUE!")
	}

}

func TestFileMgr_IsInitialized_01(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	isInitialized := srcFMgr.IsInitialized()

	if !isInitialized {
		t.Error("Expected isInitialized = 'true'. Instead, it is FALSE!")
	}
}

func TestFileMgr_IsInitialized_02(t *testing.T) {

	fh := FileHelper{}

	targetFile := fh.AdjustPathSlash("../filesfortest/newfilesfortest/newerFileForTest_01.txt")

	srcFMgr, err := FileMgr{}.New(targetFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.New(targetFile). "+
			"targetFile='%v' Error='%v'", targetFile, err.Error())
	}

	srcFMgr.Empty()

	isInitialized := srcFMgr.IsInitialized()

	if isInitialized {
		t.Error("Expected isInitialized = 'false'. Instead, it is TRUE!")
	}
}

func TestFileMgr_IsInitialized_03(t *testing.T) {

	srcFMgr := FileMgr{}

	isInitialized := srcFMgr.IsInitialized()

	if isInitialized {
		t.Error("Expected isInitialized = 'false'. Instead, it is TRUE!")
	}
}

func TestFileMgr_MoveFileToNewDirMgr_01(t *testing.T) {
	fh := FileHelper{}
	setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
	srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
	destDir := fh.AdjustPathSlash("..\\logTest")
	setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

	if fh.DoesFileExist(setupDestFile) {
		err := fh.DeleteDirFile(setupDestFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
		}

		if fh.DoesFileExist(setupDestFile) {
			t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
		}
	}

	if fh.DoesFileExist(srcFile) {
		err := fh.DeleteDirFile(srcFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
		}

		if fh.DoesFileExist(srcFile) {
			t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
		}
	}

	err := fh.CopyFileByIo(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	dMgr, err := DirMgr{}.New(destDir)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(destDir). destDir='%v' Error='%v'", destDir, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDirMgr(dMgr)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDirMgr(dMgr). dMgr.path='%v'  Error='%v'", dMgr.path, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

}

func TestFileMgr_MoveFileToNewDir_01(t *testing.T) {
	fh := FileHelper{}
	setupSrcFile := fh.AdjustPathSlash("..\\logTest\\FileMgmnt\\TestFile003.txt")
	srcFile := fh.AdjustPathSlash("..\\logTest\\FileSrc\\TestFile003.txt")
	destDir := fh.AdjustPathSlash("..\\logTest")
	setupDestFile := fh.AdjustPathSlash("..\\logTest\\TestFile003.txt")

	if fh.DoesFileExist(setupDestFile) {
		err := fh.DeleteDirFile(setupDestFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting destination file, '%v'. Error:'%v'", setupDestFile, err.Error())
		}

		if fh.DoesFileExist(setupDestFile) {
			t.Error(fmt.Sprintf("Error - destination file, '%v' STILL EXISTS!", setupDestFile))
		}
	}

	if fh.DoesFileExist(srcFile) {
		err := fh.DeleteDirFile(srcFile)

		if err != nil {
			t.Errorf("Error on DeleteDirFile() deleting source file, '%v'. Error:'%v'", srcFile, err.Error())
		}

		if fh.DoesFileExist(srcFile) {
			t.Errorf("Error - Failed to Delete 'srcFile', '%v' STILL EXISTS!", srcFile)
		}
	}

	err := fh.CopyFileByIo(setupSrcFile, srcFile)

	if err != nil {
		t.Errorf("Received error copying setup file '%v' to source file. srcFile '%v' does NOT Exist. Error='%v'", setupSrcFile, srcFile, err.Error())
	}

	if !fh.DoesFileExist(srcFile) {
		t.Errorf("Source File '%v' does NOT EXIST!!", srcFile)
	}

	srcFileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(srcFile)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(srcFile). srcFile='%v'  Error='%v'", srcFile, err.Error())
	}

	newFMgr, err := srcFileMgr.MoveFileToNewDir(destDir)

	if err != nil {
		t.Errorf("Error returned by srcFileMgr.MoveFileToNewDir(destDir). destDir='%v'  Error='%v'", destDir, err.Error())
	}

	if !fh.DoesFileExist(newFMgr.absolutePathFileName) {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	doesExist, err := newFMgr.DoesThisFileExist()

	if err != nil {
		t.Errorf("Error returned by newFMgr.DoesThisFileExist(). newFMgr.absolutePathFileName='%v' Error='%v'", newFMgr.absolutePathFileName, err.Error())
	}

	if !doesExist {
		t.Errorf("Error: NewFromPathFileNameExtStr Destination 'Moved' File DOES NOT EXIST! newFMgr.DoesThisFileExist()=='FALSE' newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

	err = newFMgr.DeleteThisFile()

	if err != nil {
		t.Errorf("Error: Attempted clean-up and deletion of destination file FAILED!. newFMgr.absolutePathFileName='%v'", newFMgr.absolutePathFileName)
	}

}

func TestFileMgr_New_01(t *testing.T) {

	fh := FileHelper{}

	relPath := "..\\logTest\\CmdrX\\CmdrX.log"
	commonDir, err := fh.MakeAbsolutePath(relPath)

	if err != nil {
		t.Errorf("Received Error on fh.MakeAbsolutePath(relPath). "+
			"relPath='%v'  Error='%v'", relPath, err.Error())
	}

	fileName := "CmdrX"
	fileNameExt := "CmdrX.log"
	extName := ".log"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
	}

	if fileMgr.fileName != fileName {
		t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
	}

	if fileMgr.fileExt != extName {
		t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
	}

	if fileMgr.fileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
	}

	if !fileMgr.isInitialized {
		t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
	}

	if !fileMgr.isFileNamePopulated {
		t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
	}

	if !fileMgr.isFileNameExtPopulated {
		t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
	}

	if !fileMgr.isFileExtPopulated {
		t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
	}

}

func TestFileMgr_New_02(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("..\\logTest\\CmdrX\\CmdrX.log")

	fileName := "CmdrX"
	fileNameExt := "CmdrX.log"
	extName := ".log"

	fileMgr, err := FileMgr{}.New(commonDir)

	if err != nil {
		t.Errorf("Received Error on FileMgr{}.New(commonDir)  Error='%v'", err.Error())
	}

	if fileMgr.fileName != fileName {
		t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
	}

	if fileMgr.fileExt != extName {
		t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
	}

	if fileMgr.fileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
	}

	if !fileMgr.isInitialized {
		t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
	}

	if !fileMgr.isFileNamePopulated {
		t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
	}

	if !fileMgr.isFileNameExtPopulated {
		t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
	}

	if !fileMgr.isFileExtPopulated {
		t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
	}

}

func TestFileMgr_NewFromPathFileNameExtStr_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")
	fileName := "xt_dirmgr_01_test"
	fileNameExt := "xt_dirmgr_01_test.go"
	extName := ".go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(commonDir)

	if err != nil {
		t.Error("Received Error on GetPathFileNameElements Error:", err)
	}

	if fileMgr.fileName != fileName {
		t.Error(fmt.Sprintf("Expected File Name, %v, got:", fileName), fileMgr.fileName)
	}

	if fileMgr.fileExt != extName {
		t.Error(fmt.Sprintf("Expected File Extension, %v, got:", extName), fileMgr.fileExt)
	}

	if fileMgr.fileNameExt != fileNameExt {
		t.Error(fmt.Sprintf("Expected File Name + Extension, %v, got:", fileNameExt), fileMgr.fileNameExt)
	}

	if !fileMgr.isInitialized {
		t.Error("Expected fileMgr.isInitialized=='true', got:", fileMgr.isInitialized)
	}

	if !fileMgr.isFileNamePopulated {
		t.Error("Expected fileMgr.isFileNamePopulated=='true', got:", fileMgr.isFileNamePopulated)
	}

	if !fileMgr.isFileNameExtPopulated {
		t.Error("Expected fileMgr.isFileNameExtPopulated=='true', got:", fileMgr.isFileNameExtPopulated)
	}

	if !fileMgr.isFileExtPopulated {
		t.Error("Expected fileMgr.isFileExtPopulated=='true', got:", fileMgr.isFileExtPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated=='true', got:", fileMgr.isAbsolutePathFileNamePopulated)
	}

}

func TestFileMgr_NewFromPathFileNameExtStr_02(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
			"path== '%v' Error: %v", path, err)
	}

	if eFileNameExt != fileMgr.fileNameExt {
		t.Errorf("Expected extracted fileNameExt == %v, instead got: %v",
			eFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != "filehelperexamples" {
		t.Errorf("Expected fileMgr.fileName== 'filehelperexamples', "+
			"instead got: fileMgr.fileName== %v", fileMgr.fileName)
	}

	if fileMgr.fileExt != ".go" {
		t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
			fileMgr.fileExt)
	}

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated==%v",
			fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist==%v", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
	}

}

func TestFileMgr_NewFromPathFileNameExtStr_03(t *testing.T) {

	path := "filehelperexamples"

	_, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err == nil {
		t.Error("Expected an error from FileMgr{}.NewFromPathFileNameExtStr(path) " +
			"because path='filehelperexamples'. However, NO ERROR WAS RETURNED! ")
	}
}

func TestFileMgr_NewFromPathFileNameExtStr_04(t *testing.T) {

	path := "../appExamples/filehelperexamples.go"

	eFileNameExt := "filehelperexamples.go"

	fileMgr, err := FileMgr{}.NewFromPathFileNameExtStr(path)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromPathFileNameExtStr(path) "+
			"path=='%v' Error: %v ", path, err)
	}

	if eFileNameExt != fileMgr.fileNameExt {
		t.Errorf("Expected extracted fileNameExt == '%v', instead got: '%v' ",
			eFileNameExt, fileMgr.fileNameExt)
	}

	if "filehelperexamples" != fileMgr.fileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead got: fileMgr.fileName== '%v'",
			"filehelperexamples", fileMgr.fileName)
	}

	if ".go" != fileMgr.fileExt {
		t.Errorf("Expected fileMgr.fileExt== '.go', instead got: fileMgr.fileExt== %v",
			fileMgr.fileExt)
	}

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.dMgr.isPathPopulated==true', instead got: "+
			"fileMgr.isPathPopulated==%v",
			fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: "+
			"fileMgr.doesAbsolutePathFileNameExist== %v", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Errorf("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, "+
			"fileMgr.isAbsolutePathFileNamePopulated == '%v' ", fileMgr.isAbsolutePathFileNamePopulated)
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Errorf("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is '%v'",
			fileMgr.dMgr.doesAbsolutePathExist)
	}

}

func TestFileMgr_NewFromFileInfo_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName := "newerFileForTest_01"
	expectedExt := ".txt"
	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	absPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	info, err := fh.GetFileInfoFromPath(absPathFileNameExt)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileInfoFromPath(absPathFileNameExt). absPathFileNameExt='%v'  Error='%v'", absPathFileNameExt, err.Error())
	}

	fileMgr, err := FileMgr{}.NewFromFileInfo(absPath, info)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromFileInfo(absPath, info). absPath='%v' info.Name()='%v'  Error='%v'", absPath, info.Name(), err.Error())
	}

	if fileMgr.fileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != expectedFileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
	}

	if fileMgr.fileExt != expectedExt {
		t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
	}

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.doesAbsolutePathFileNameExist)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.actualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, expectedFileNameExt)

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
		t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
	}

}

func TestFileMgr_NewFromDirMgrFileNameExt_02(t *testing.T) {

	rawFileNameExt := "./newerFileForTest_01.txt"
	expectedFileNameExt := "newerFileForTest_01.txt"

	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash("../filesfortest/newfilesfortest")

	absolutePath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v' Error='%v'", adjustedPath, err.Error())
	}

	dMgr, err := DirMgr{}.New(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if absolutePath != dMgr.absolutePath {
		t.Errorf("Expected dMgr.absolutePath='%v'.  Instead, dMgr.absolutePath='%v'", absolutePath, dMgr.absolutePath)
	}

	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, rawFileNameExt)

	absPath, err := fh.MakeAbsolutePath(adjustedPath)

	if err != nil {
		t.Errorf("Error returned from fh.MakeAbsolutePath(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	expectedAbsPathFileNameExt := absPath + string(os.PathSeparator) + expectedFileNameExt

	if expectedAbsPathFileNameExt != fMgr.absolutePathFileName {
		t.Errorf("Expected absolutePathFileName='%v'.  Instead, absolutePathFileName='%v'", expectedAbsPathFileNameExt, fMgr.absolutePathFileName)
	}

}

func TestFileMgr_NewFromDirStrFileNameStr_01(t *testing.T) {

	expectedFileNameExt := "newerFileForTest_01.txt"
	expectedFileName := "newerFileForTest_01"
	expectedExt := ".txt"

	fh := FileHelper{}
	rawPath := "../filesfortest/newfilesfortest"
	expectedPath := fh.AdjustPathSlash(rawPath)
	expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

	if err != nil {
		t.Errorf("Error returned by fh.MakeAbsolutePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
	}

	fileMgr, err := FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt)

	if err != nil {
		t.Errorf("Error returned from FileMgr{}.NewFromDirStrFileNameStr(rawPath, expectedFileNameExt). rawPath='%v' expectedFileNameExt='%v'  Error='%v'", rawPath, expectedFileNameExt, err.Error())
	}

	if fileMgr.fileNameExt != expectedFileNameExt {
		t.Errorf("Expected extracted fileMgr.fileNameExt == %v, instead fileMgr.fileNameExt='%v' ", expectedFileNameExt, fileMgr.fileNameExt)
	}

	if fileMgr.fileName != expectedFileName {
		t.Errorf("Expected fileMgr.fileName== '%v', instead fileMgr.fileName== '%v'", expectedFileName, fileMgr.fileName)
	}

	if fileMgr.fileExt != expectedExt {
		t.Errorf("Expected fileMgr.fileExt== '%v', instead got: fileMgr.fileExt=='%v'", expectedExt, fileMgr.fileName)
	}

	if !fileMgr.dMgr.isPathPopulated {
		t.Errorf("Expected 'fileMgr.isPathPopulated==true', instead got: fileMgr.isPathPopulated=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.doesAbsolutePathFileNameExist {
		t.Errorf("Expected 'fileMgr.doesAbsolutePathFileNameExist==true', instead got: fileMgr.doesAbsolutePathFileNameExist=='%v'", fileMgr.dMgr.isPathPopulated)
	}

	if !fileMgr.isAbsolutePathFileNamePopulated {
		t.Error("Expected fileMgr.isAbsolutePathFileNamePopulated == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.dMgr.doesAbsolutePathExist {
		t.Error("Expected fileMgr.doesAbsolutePathExist == 'true'.  Instead, it is 'false'")
	}

	if !fileMgr.actualFileInfo.IsFInfoInitialized {
		t.Error("Expected fileMgr.actualFileInfo.IsFInfoInitialized='true'.  Error, it is 'false'")
	}

	if fileMgr.actualFileInfo.Name() != expectedFileNameExt {
		t.Errorf("Expected fileMgr.actualFileInfo.Name()=='%v'.  Instead, fileMgr.actualFileInfo.Name()=='%v'.", expectedFileNameExt, fileMgr.actualFileInfo.Name())
	}

	if expectedAbsPath != fileMgr.dMgr.absolutePath {
		t.Errorf("Expected absolutePath='%v'.  Instead, absolutePath='%v' ", expectedAbsPath, fileMgr.dMgr.absolutePath)
	}

	if expectedPath != fileMgr.dMgr.path {
		t.Errorf("Expected path='%v'.  Instead, path='%v' ", expectedPath, fileMgr.dMgr.path)
	}

}
