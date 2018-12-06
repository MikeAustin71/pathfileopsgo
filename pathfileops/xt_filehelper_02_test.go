package pathfileops

import (
	"fmt"
	"testing"
)

func TestFileHelper_GetFileExtension_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '.go' for valid file extension Instead result='%v' ", result)
	}

}

func TestFileHelper_GetFileExtension_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'.  "+
			"Instead file extension='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("dirmgr_test")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead file extension ='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return file extension == '%v'. Instead, file extension='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. "+
			"Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\pathfilego\\003_filehelper\\common\\")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_07(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_08(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("test.....go")

	expectedExt := ".go"

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileExtension_09(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("")

	expectedExt := ""

	result, isEmpty, err := fh.GetFileExtension(commonDir)

	if err == nil {
		t.Errorf("Expected an error to be returned from fh.GetFileExt(commonDir). "+
			"commonDir='%v' NO Error was returned!", commonDir)
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileExt isEmpty=='%v'. Instead, isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedExt {
		t.Errorf("Expected GetFileExt to return result == '%v' for file extension. Instead result='%v' ", expectedExt, result)
	}

}

func TestFileHelper_GetFileNameWithExt_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
			false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("F:\\pathfilego\\003_filehelper\\common\\")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ",
			true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
	expectedFNameExt := "xt_dirmgr_01_test.go"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_07(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("dirmgr_test")
	expectedFNameExt := "dirmgr_test"

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_08(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".go")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err == nil {
		t.Errorf("Expected error returned from fh.GetFileNameWithExt(commonDir). "+
			"Instead, NO ERROR was returned. commonDir='%v'  ", commonDir)
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ",
			expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFileNameWithExt_09(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("")
	expectedFNameExt := ""

	fNameExt, isEmpty, err := fh.GetFileNameWithExt(commonDir)

	if err == nil {
		t.Errorf("Error error returned from fh.GetFileNameWithExt(commonDir). Result- commonDir='%v' No Error Returned!", commonDir)
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
	}

	if expectedFNameExt != fNameExt {
		t.Errorf("Expected GetFileNameWithExt to return fNameExt == '%v'. Istead, fNameExt='%v' ", expectedFNameExt, fNameExt)
	}

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_01(t *testing.T) {

	rawPath := "../filesfortest/newfilesfortest/newerFileForTest_01.txt"
	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash(rawPath)

	firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if firstCharIdx != 3 {
		t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
	}

	expectedLastIdx := len(adjustedPath) - 1

	if expectedLastIdx != lastCharIdx {
		t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
	}

}

func TestFileHelper_GetFirstLastNonSeparatorCharIndexInPathStr_02(t *testing.T) {

	rawPath := "D:/filesfortest/newfilesfortest/newerFileForTest_01.txt"
	fh := FileHelper{}
	adjustedPath := fh.AdjustPathSlash(rawPath)

	firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath)

	if err != nil {
		t.Errorf("Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedPath). adjustedPath='%v'  Error='%v'", adjustedPath, err.Error())
	}

	if firstCharIdx != 3 {
		t.Errorf("Expected first char index= '3'.  Instead, first char index= '%v'", firstCharIdx)
	}

	expectedLastIdx := len(adjustedPath) - 1

	if expectedLastIdx != lastCharIdx {
		t.Errorf("Expected last index = '%v'.  Instead, last index = '%v'", expectedLastIdx, lastCharIdx)
	}

}

func TestFileHelper_GetFileNameWithoutExt_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")
	expectedFileName := "xt_dirmgr_01_test"

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty == true {
		t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, instead got: %v", expectedFileName, result)
	}

}

func TestFileHelper_GetFileNameWithoutExt_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_01_test")
	expectedFileName := "dirmgr_01_test"

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if isEmpty == true {
		t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v'", isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileNameWithoutExt to return result == '%v' for valid file name, instead got '%v' ", expectedFileName, result)
	}

}

func TestFileHelper_GetFileNameWithoutExt_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")
	expectedFileName := ""

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). "+
			"commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v', instead got %v ",
			true, isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
			"instead got: %v", expectedFileName, result)
	}

}

func TestFileHelper_GetFileNameWithoutExt_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")
	expectedFileName := "xt_dirmgr_01_test"

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty == true {
		t.Errorf("Expected isEmpty GetFileNameWithoutExt for valid file extension to return 'false'. Instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, instead got: %v", expectedFileName, result)
	}

}

func TestFileHelper_GetFileNameWithoutExt_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")
	expectedFileName := ""

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if true != isEmpty {
		t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v'. Instead isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
			"instead got: %v", expectedFileName, result)
	}

}

func TestFileHelper_GetFileNameWithoutExt_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")
	expectedFileName := "common"

	result, isEmpty, err := fh.GetFileNameWithoutExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetFileNameWithoutExt(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetFileNameWithoutExt isEmpty='%v'. Instead isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedFileName {
		t.Errorf("Expected GetFileExt to return result == '%v' for valid file name, "+
			"instead got: %v", expectedFileName, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_01(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_02(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("..\\..\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash("..\\..\\pathfilego\\003_filehelper\\common")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_03(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return 'false', instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid file name. Instead path='%v'", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_04(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash("D:\\go\\work\\src\\MikeAustin71\\pathfilego\\003_filehelper\\common")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name. "+
			"Instead path=='%v' ", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_05(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash("")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Expected no error returned from fh.GetPathFromPathFileName(commonDir). "+
			"Instead an error WAS Returned. commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != true {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", true, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file name. "+
			"Instead path=='%v' ", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_06(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")

	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). "+
			"commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty GetPathFromPathFileName for valid file extension to return "+
			"'false', instead isEmpty='%v' ", isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file "+
			"name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_07(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("./")

	expectedDir := "."

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). "+
			"commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file "+
			"name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_08(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".")

	expectedDir := fh.AdjustPathSlash(".")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' "+
			"Error='%v'", commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid "+
			"path/file name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_09(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("..")

	expectedDir := fh.AdjustPathSlash("..")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file "+
			"name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_10(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("")

	expectedDir := fh.AdjustPathSlash("")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err == nil {
		t.Errorf("Expected error to be returned from fh.GetPathFromPathFileName(commonDir). "+
			"commonDir='%v' No Error Returned!", commonDir)
	}

	if true != isEmpty {
		t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
			true, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid pathn/file"+
			"name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_11(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("../../../")

	expectedDir := fh.AdjustPathSlash("../../..")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathFromPathFileName(commonDir). commonDir='%v' Error='%v'",
			commonDir, err.Error())
	}

	if false != isEmpty {
		t.Errorf("Expected GetPathFromPathFileName isEmpty=='%v'. Instead, isEmpty='%v' ",
			false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file "+
			"name, instead got: %v", expectedDir, result)
	}

}

func TestFileHelper_GetPathFromPathFileName_12(t *testing.T) {
	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("./xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash(".")

	result, isEmpty, err := fh.GetPathFromPathFileName(commonDir)

	if err != nil {
		t.Errorf("Expected no error returned from fh.GetPathFromPathFileName(commonDir). "+
			"Instead an error WAS Returned. commonDir='%v' Error='%v'", commonDir, err.Error())
	}

	if isEmpty != false {
		t.Errorf("Expected isEmpty='%v', instead isEmpty='%v' ", false, isEmpty)
	}

	if result != expectedDir {
		t.Errorf("Expected GetPathFromPathFileName to return path == '%v' for valid path/file"+
			"name. Instead path=='%v' ", expectedDir, result)
	}

}

func TestFileHelper_GetPathAndFileNameExt_01(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

	expectedFileNameExt := "xt_dirmgr_01_test.go"

	pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v' "+
			"Error='%v'", commonDir, err.Error())
	}

	if false != bothAreEmpty {
		t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
			false, bothAreEmpty)
	}

	if pathDir != expectedDir {
		t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. "+
			"Instead, path== '%v' ", expectedDir, pathDir)
	}

	if fileNameExt != expectedFileNameExt {
		t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
			"fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
	}

}

func TestFileHelper_GetPathAndFileNameExt_02(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\")

	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

	expectedFileNameExt := ""

	pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  Error='%v'",
			commonDir, err.Error())
	}

	if false != bothAreEmpty {
		t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
			false, bothAreEmpty)
	}

	if pathDir != expectedDir {
		t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ",
			expectedDir, pathDir)
	}

	if fileNameExt != expectedFileNameExt {
		t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
			"fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
	}

}

func TestFileHelper_GetPathAndFileNameExt_03(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common\\dirmgr_test")

	expectedDir := fh.AdjustPathSlash(".\\pathfilego\\003_filehelper\\common")

	expectedFileNameExt := "dirmgr_test"

	pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  "+
			"Error='%v'", commonDir, err.Error())
	}

	if false != bothAreEmpty {
		t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ",
			false, bothAreEmpty)
	}

	if pathDir != expectedDir {
		t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ",
			expectedDir, pathDir)
	}

	if fileNameExt != expectedFileNameExt {
		t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, "+
			"fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
	}

}

func TestFileHelper_GetPathAndFileNameExt_04(t *testing.T) {

	fh := FileHelper{}

	commonDir := fh.AdjustPathSlash("xt_dirmgr_01_test.go")

	expectedDir := fh.AdjustPathSlash("")

	expectedFileNameExt := "xt_dirmgr_01_test.go"

	pathDir, fileNameExt, bothAreEmpty, err := fh.GetPathAndFileNameExt(commonDir)

	if err != nil {
		t.Errorf("Error returned from fh.GetPathAndFileNameExt(commonDir). commonDir='%v'  Error='%v'", commonDir, err.Error())
	}

	if false != bothAreEmpty {
		t.Errorf("Expected GetPathAndFileNameExt bothAreEmpty='%v'. Instead, bothAreEmpty='%v' ", false, bothAreEmpty)
	}

	if pathDir != expectedDir {
		t.Errorf("Expected GetPathAndFileNameExt to return path == '%v'. Instead, path== '%v' ", expectedDir, pathDir)
	}

	if fileNameExt != expectedFileNameExt {
		t.Errorf("Expected GetPathAndFileNameExt to return fileNameExt == '%v'. Instead, fileNameExt == '%v' ", expectedFileNameExt, fileNameExt)
	}

}

func TestFileHelper_GetFileLastModificationDate(t *testing.T) {

	fh := FileHelper{}
	target, err := fh.MakeAbsolutePath(fh.AdjustPathSlash(alogtopTest2Text))

	if err != nil {
		t.Error("Error from FileHelper:MakeAbsolutePath(): ", err.Error())
	}

	tStrFmt := "2006-01-02 15:04:05.000000000"

	fileTime, tStr, err := fh.GetFileLastModificationDate(target, tStrFmt)

	if err != nil {
		t.Error("Error from FileHelper:GetFileLastModificationDate():", err.Error())
	}

	fInfo, err := fh.GetFileInfoFromPath(target)

	if err != nil {
		t.Error("Error from FileHelper:GetFileInfoFromPath():", err.Error())
	}

	actualFileTime := fInfo.ModTime()

	expected := actualFileTime.Format(tStrFmt)

	if tStr != expected {
		t.Error(fmt.Sprintf("Expected Time String for file %v == %v, received time string: ", target, expected), tStr)
	}

	if !actualFileTime.Equal(fileTime) {
		t.Error(fmt.Sprintf("Expected Time value %v, instead got:", actualFileTime), fileTime)
	}
}
