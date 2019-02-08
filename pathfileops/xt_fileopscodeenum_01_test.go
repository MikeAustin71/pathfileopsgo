package pathfileops

import (
	"strings"
	"testing"
)

func TestFileOperationCode_01(t *testing.T) {

	fopFirst := FileOpCode.None()

	if int(fopFirst) != 0 {
		t.Errorf("Error: Expected first File Operations Code = 0.  Instead, first "+
			"File Operation Code = '%v' ", int(fopFirst))
	}

	if 0 != fopFirst.Value() {
		t.Errorf("Error: Expected first File Operations Code Value = 0.  Instead, first "+
			"File Operation Code Value = '%v' ", fopFirst.Value())
	}

	fopLast := FileOpCode.CreateDestinationFile()

	if int(fopLast) != 14 {
		t.Errorf("Error: Expected FileOpCode.CreateDestinationFile() = 14.  Instead, "+
			"FileOpCode.CreateDestinationFile() = '%v' ", int(fopLast))
	}

	if 14 != fopLast.Value() {
		t.Errorf("Error: Expected FileOpCode.CreateDestinationFile() Value = 14.  Instead, "+
			"FileOpCode.CreateDestinationFile() = '%v' ", fopLast.Value())
	}

}

func TestFileOperationCode_02(t *testing.T) {

	fopNone := FileOpCode.None()

	strValue := fopNone.String()

	if "None" != strValue {
		t.Errorf("Error: Expected string value of FileOpCode.None() = 'None' .  Instead, "+
			"string value of FileOpCode.None() = '%v' ", strValue)
	}

	fopLast := FileOpCode.CreateDestinationFile()

	strValue = fopLast.String()

	if "CreateDestinationFile" != strValue {
		t.Errorf("Error: Expected string value of FileOpCode.CreateDestinationFile() = 'CreateDestinationFile' . "+
			" Instead, string value of FileOpCode.CreateDestinationFile() = '%v' ", strValue)

	}

}

func TestFileOperationCode_03(t *testing.T) {

	strValue := "None"

	fopNone, err := FileOpCode.ParseString(strValue, true)

	if err != nil {
		t.Errorf("Error returned by FileOpCode.ParseString(strValue, true). "+
			"strValue='%v' Error='%v' ", strValue, err.Error())
	}

	if fopNone != FileOperationCode(0).None() {
		t.Errorf("Error: Expected fopNone = FileOperationCode(0).None() .  Instead, "+
			"string value of fopNone = '%v' - int Value of fopNone='%v' ",
			fopNone.String(), int(fopNone))
	}

	strValue = "CreateDestinationFile"

	fopLast, err := FileOpCode.ParseString(strValue, true)

	if err != nil {
		t.Errorf("Error returned by (2) FileOpCode.ParseString(strValue, true). "+
			"strValue='%v' Error='%v' ", strValue, err.Error())
	}

	if fopLast != FileOperationCode(0).CreateDestinationFile() {
		t.Errorf("Error: Expected string value of fopLast = FileOperationCode(0)."+
			"CreateDestinationFile(). Instead, string value of fopLast = '%v' - "+
			"int value of fopLas = '%v' ", fopLast.String(), int(fopLast))

	}

}

func TestFileOperationCode_04(t *testing.T) {

	opsAry := make([]FileOperationCode, 15)

	opsAry[0] = FileOperationCode(0).None()
	opsAry[1] = FileOpCode.MoveSourceFileToDestination()
	opsAry[2] = fileOpCode.DeleteDestinationFile()
	opsAry[3] = FileOperationCode(0).DeleteSourceFile()
	opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
	opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
	opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
	opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
	opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
	opsAry[9] = FileOperationCode(0).CreateSourceDir()
	opsAry[10] = FileOpCode.CreateSourceDirAndFile()
	opsAry[11] = FileOpCode.CreateSourceFile()
	opsAry[12] = fileOpCode.CreateDestinationDir()
	opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
	opsAry[14] = FileOpCode.CreateDestinationFile()

	opsStrings := make([]string, 15)
	opsStrings[0] = "None"
	opsStrings[1] = "MoveSourceFileToDestination"
	opsStrings[2] = "DeleteDestinationFile"
	opsStrings[3] = "DeleteSourceFile"
	opsStrings[4] = "DeleteSourceAndDestinationFiles"
	opsStrings[5] = "CopySourceToDestinationByHardLinkByIo"
	opsStrings[6] = "CopySourceToDestinationByIoByHardLink"
	opsStrings[7] = "CopySourceToDestinationByHardLink"
	opsStrings[8] = "CopySourceToDestinationByIo"
	opsStrings[9] = "CreateSourceDir"
	opsStrings[10] = "CreateSourceDirAndFile"
	opsStrings[11] = "CreateSourceFile"
	opsStrings[12] = "CreateDestinationDir"
	opsStrings[13] = "CreateDestinationDirAndFile"
	opsStrings[14] = "CreateDestinationFile"

	for i := 0; i < len(opsAry); i++ {

		if opsAry[i].String() != opsStrings[i] {
			t.Errorf("Error: opsAry[i].String() != opsStrings[i]. "+
				"opsAry[%v].String()='%v' opsStrings[%v]='%v'", i, opsAry[i].String(), i, opsStrings[i])
		}

	}

}

func TestFileOperationCode_05(t *testing.T) {

	opsAry := make([]FileOperationCode, 15)

	opsAry[0] = FileOperationCode(0).None()
	opsAry[1] = FileOpCode.MoveSourceFileToDestination()
	opsAry[2] = fileOpCode.DeleteDestinationFile()
	opsAry[3] = FileOperationCode(0).DeleteSourceFile()
	opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
	opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
	opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
	opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
	opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
	opsAry[9] = FileOperationCode(0).CreateSourceDir()
	opsAry[10] = FileOpCode.CreateSourceDirAndFile()
	opsAry[11] = FileOpCode.CreateSourceFile()
	opsAry[12] = fileOpCode.CreateDestinationDir()
	opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
	opsAry[14] = FileOpCode.CreateDestinationFile()

	opsStrings := make([]string, 15)
	opsStrings[0] = "None"
	opsStrings[1] = "MoveSourceFileToDestination"
	opsStrings[2] = "DeleteDestinationFile"
	opsStrings[3] = "DeleteSourceFile"
	opsStrings[4] = "DeleteSourceAndDestinationFiles"
	opsStrings[5] = "CopySourceToDestinationByHardLinkByIo"
	opsStrings[6] = "CopySourceToDestinationByIoByHardLink"
	opsStrings[7] = "CopySourceToDestinationByHardLink"
	opsStrings[8] = "CopySourceToDestinationByIo"
	opsStrings[9] = "CreateSourceDir"
	opsStrings[10] = "CreateSourceDirAndFile"
	opsStrings[11] = "CreateSourceFile"
	opsStrings[12] = "CreateDestinationDir"
	opsStrings[13] = "CreateDestinationDirAndFile"
	opsStrings[14] = "CreateDestinationFile"

	for i := 0; i < len(opsAry); i++ {

		fop, err := FileOperationCode(0).ParseString(opsStrings[i], true)

		if err != nil {
			t.Errorf("Error returned by FileOperationCode(0).ParseString(opsStrings[i], true). "+
				"i='%v' opsStrings[%v]='%v' Error='%v' ", i, i, opsStrings[i], err.Error())
		}

		if fop != opsAry[i] {
			t.Errorf("Error: fop != opsAry[i]. "+
				"fop.String() ='%v' opsAry[%v]='%v'", fop.String(), i, opsAry[i].String())
		}

	}
}

func TestFileOperationCode_06(t *testing.T) {

	opsAry := make([]FileOperationCode, 15)

	opsAry[0] = FileOperationCode(0).None()
	opsAry[1] = FileOpCode.MoveSourceFileToDestination()
	opsAry[2] = fileOpCode.DeleteDestinationFile()
	opsAry[3] = FileOperationCode(0).DeleteSourceFile()
	opsAry[4] = FileOperationCode(0).DeleteSourceAndDestinationFiles()
	opsAry[5] = FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
	opsAry[6] = FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
	opsAry[7] = FileOperationCode(0).CopySourceToDestinationByHardLink()
	opsAry[8] = FileOperationCode(0).CopySourceToDestinationByIo()
	opsAry[9] = FileOperationCode(0).CreateSourceDir()
	opsAry[10] = FileOpCode.CreateSourceDirAndFile()
	opsAry[11] = FileOpCode.CreateSourceFile()
	opsAry[12] = fileOpCode.CreateDestinationDir()
	opsAry[13] = fileOpCode.CreateDestinationDirAndFile()
	opsAry[14] = FileOpCode.CreateDestinationFile()

	opsStrings := make([]string, 15)
	opsStrings[0] = strings.ToLower("None")
	opsStrings[1] = strings.ToLower("MoveSourceFileToDestination")
	opsStrings[2] = strings.ToLower("DeleteDestinationFile")
	opsStrings[3] = strings.ToLower("DeleteSourceFile")
	opsStrings[4] = strings.ToLower("DeleteSourceAndDestinationFiles")
	opsStrings[5] = strings.ToLower("CopySourceToDestinationByHardLinkByIo")
	opsStrings[6] = strings.ToLower("CopySourceToDestinationByIoByHardLink")
	opsStrings[7] = strings.ToLower("CopySourceToDestinationByHardLink")
	opsStrings[8] = strings.ToLower("CopySourceToDestinationByIo")
	opsStrings[9] = strings.ToLower("CreateSourceDir")
	opsStrings[10] = strings.ToLower("CreateSourceDirAndFile")
	opsStrings[11] = strings.ToLower("CreateSourceFile")
	opsStrings[12] = strings.ToLower("CreateDestinationDir")
	opsStrings[13] = strings.ToLower("CreateDestinationDirAndFile")
	opsStrings[14] = strings.ToLower("CreateDestinationFile")

	for i := 0; i < len(opsAry); i++ {

		fop, err := FileOperationCode(0).ParseString(opsStrings[i], false)

		if err != nil {
			t.Errorf("Error returned by FileOperationCode(0).ParseString(opsStrings[i], true). "+
				"i='%v' opsStrings[%v]='%v' Error='%v' ", i, i, opsStrings[i], err.Error())
		}

		if fop != opsAry[i] {
			t.Errorf("Error: fop != opsAry[i]. "+
				"fop.String() ='%v' opsAry[%v]='%v'", fop.String(), i, opsAry[i].String())
		}

	}
}
