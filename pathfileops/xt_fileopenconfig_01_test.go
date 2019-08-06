package pathfileops

import (
  "os"
  "strings"
  "testing"
)

func TestFileOpenConfig_CopyIn_01(t *testing.T) {

  expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

  fOpCfg1, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeAppend(), FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by fOpCfg1.New().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFOpenCode, err := fOpCfg1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpCfg1.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v'\n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpCfg2 := FileOpenConfig{}

  fOpCfg2.CopyIn(&fOpCfg1)

  actualFOpenCode2, err := fOpCfg2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpCfg2.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v'\n",
      expectedFOpenCode, actualFOpenCode2)
  }
}

func TestFileOpenConfig_CopyIn_02(t *testing.T) {

  fOpStatus1 := FileOpenConfig{}

  fOpStatus2 := FileOpenConfig{}

  fOpStatus2.CopyIn(&fOpStatus1)

  if !fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Error: Expected fOpStatus1==fOpStatus2.\n" +
      "HOWEVER, THEY ARE NOT EQUAL!\n")
  }
}

func TestFileOpenConfig_CopyOut_01(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.New().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpStatus2 := fOpStatus1.CopyOut()

  actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v'\n",
      expectedFOpenCode, actualFOpenCode2)
  }

}

func TestFileOpenConfig_CopyOut_02(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.New().\n" +
      "Error='%v'\n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v'\n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpStatus2 := fOpStatus1.CopyOut()

  actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v'\n",
      expectedFOpenCode, actualFOpenCode2)
  }

}

func TestFileOpenConfig_CopyOut_03(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpStatus2 := fOpStatus1.CopyOut()

  actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode2)
  }

}

func TestFileOpenConfig_CopyOut_04(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpStatus1.fileOpenModes = make([]FileOpenMode, 0)

  fOpStatus2 := fOpStatus1.CopyOut()

  actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode2)
  }

}

func TestFileOpenConfig_CopyOut_05(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus1, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus1.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus1.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'. Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

  fOpStatus1.fileOpenModes = nil

  fOpStatus2 := fOpStatus1.CopyOut()

  actualFOpenCode2, err := fOpStatus2.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by fOpStatus2.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode2 {
    t.Errorf("Error: Expected File Open Code #2 ='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode2)
  }

}

func TestFileOpenConfig_Equal_01(t *testing.T) {

  fOpStatus1, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus2 := fOpStatus1.CopyOut()

  if !fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2==fOpStatus1.\n" +
      "WRONG: They are NOT Equal!\n")
  }

  if !fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus1==fOpStatus2.\n" +
      "WRONG: They are NOT Equal!\n")
  }

}

func TestFileOpenConfig_Equal_02(t *testing.T) {

  fOpStatus1, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus2, err := FileOpenConfig{}.New(
    FOpenType.TypeWriteOnly(), FOpenMode.ModeAppend(), FOpenMode.ModeExclusive())

  if err != nil {
    t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2!=fOpStatus1.\n" +
      "WRONG: They ARE Equal!\n")
  }

  if fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus1!=fOpStatus2.\nWRONG: They ARE Equal!\n")
  }

}

func TestFileOpenConfig_Equal_03(t *testing.T) {

  fOpStatus1 := FileOpenConfig{}

  fOpStatus2 := FileOpenConfig{}

  if !fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2==fOpStatus1.\nWRONG: They are NOT Equal!\n")
  }

  if !fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus ==fOpStatus2.\nWRONG: They are NOT Equal!\n")
  }

}

func TestFileOpenConfig_Equal_04(t *testing.T) {

  fOpStatus1, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus2 := FileOpenConfig{}

  if fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2!=fOpStatus1.\nWRONG: They ARE Equal!\n")
  }

  if fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus ==fOpStatus2.\nWRONG: They ARE Equal!\n")
  }

}

func TestFileOpenConfig_Equal_05(t *testing.T) {

  fOpStatus1 := FileOpenConfig{}

  fOpStatus2, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2!=fOpStatus1.\nWRONG: They ARE Equal!\n")
  }

  if fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus ==fOpStatus2.\nWRONG: They ARE Equal!\n")
  }

}

func TestFileOpenConfig_Equal_06(t *testing.T) {

  fOpStatus1, err :=
    FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New().\n"+
      "Error='%v'\n", err.Error())
  }

  fOpStatus2, err :=
    FileOpenConfig{}.New(FOpenType.TypeWriteOnly(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2!=fOpStatus1.\nWRONG: They ARE Equal!\n")
  }

  if fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus ==fOpStatus2.\nWRONG: They ARE Equal!\n")
  }

}

func TestFileOpenConfig_Equal_07(t *testing.T) {

  fOpStatus1, err :=
    FileOpenConfig{}.New(
      FOpenType.TypeReadWrite(),
      FOpenMode.ModeAppend(),
      FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by fOpStatus1=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus2, err :=
    FileOpenConfig{}.New(
      FOpenType.TypeReadWrite(),
      FOpenMode.ModeCreate(),
      FOpenMode.ModeExclusive())

  if err != nil {
    t.Errorf("Error returned by fOpStatus2=FileOpenConfig{}.New().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if fOpStatus2.Equal(&fOpStatus1) {
    t.Error("Expected fOpStatus2!=fOpStatus1.\nWRONG: They ARE Equal!\n")
  }

  if fOpStatus1.Equal(&fOpStatus2) {
    t.Error("Expected fOpStatus ==fOpStatus2.\nWRONG: They ARE Equal!\n")
  }

}

func TestFileOpenConfig_New_01(t *testing.T) {

  expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_TRUNC

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),
    FOpenMode.ModeAppend(), FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }
}

func TestFileOpenConfig_New_02(t *testing.T) {

  fOpenType := FileOpenType(-99)

  _, err := FileOpenConfig{}.New(fOpenType, FOpenMode.ModeCreate())

  if err == nil {
    t.Error("Expected Error returned by FileOpenConfig{}.New() " +
      "because of an invalid File Open Type.\n" +
      "However, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileOpenConfig_New_03(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_New_04(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_New_05(t *testing.T) {

  fOpenMode := FileOpenMode(-99)

  _, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), fOpenMode)

  if err == nil {
    t.Error("Expected an error return from FileOpenConfig{}.New()\n" +
      "because the File Open Mode was invalid.\n" +
      "HOWEVER, NO ERROR WAS RETURNED! \n")
  }

}

func TestFileOpenConfig_New_06(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus.fileOpenModes = nil

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_GetCompositeFileOpenCode_01(t *testing.T) {

  fOpCfg := FileOpenConfig{}

  _, err := fOpCfg.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode()\n" +
      "because 'fOpCfg' was NOT initialized.\n" +
      "HOWEVER, NO ERROR WAS RETURNED!\n")

  }

}

func TestFileOpenConfig_GetCompositeFileOpenCode_02(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  _, err = fOpCfg.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode()\n" +
      "because fOpCfg use TypeNone().\n" +
      "However, NO ERROR WAS RETURNED!\n")

  }

}

func TestFileOpenConfig_GetCompositeFileOpenCode_03(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes = nil

  _, err = fOpCfg.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode()\n" +
      "because fOpCfg.fileOpenModes == nil.\n" +
      "However, NO ERROR WAS RETURNED!\n")

  }

}

func TestFileOpenConfig_GetCompositeFileOpenCode_04(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(FOpenType.TypeNone(), FOpenMode.ModeAppend())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenType = FileOpenType(-99)

  _, err = fOpCfg.GetCompositeFileOpenCode()

  if err == nil {
    t.Error("Expected error return from fOpCfg.GetCompositeFileOpenCode()\n" +
      "because fOpCfg.fileOpenType is invalid.\n" +
      "However, NO ERROR WAS RETURNED!\n")

  }

}

func TestFileOpenConfig_GetCompositeFileOpenCode_05(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeReadWrite(), FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpStatus.fileOpenModes = nil

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_GetFileOpenModes_01(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeAppend(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeExclusive())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fileOpenModes := fOpCfg.GetFileOpenModes()

  if len(fileOpenModes) != 3 {
    t.Errorf("Error: Expected fileOpenModes len = 3.\n" +
      "Instead, len='%v'\n",
      len(fileOpenModes))
  }

  hasAppend := 0
  hasCreate := 0
  hasExclusive := 0

  for i := 0; i < len(fileOpenModes); i++ {

    if fileOpenModes[i] == FOpenMode.ModeAppend() {
      hasAppend++
    }

    if fileOpenModes[i] == FOpenMode.ModeCreate() {
      hasCreate++
    }

    if fileOpenModes[i] == FOpenMode.ModeExclusive() {
      hasExclusive++
    }

  }

  if hasAppend != 1 {
    t.Errorf("Error: Could not locate correct number of Appends.\n"+
      "hasAppend='%v'\n", hasAppend)
  }

  if hasCreate != 1 {
    t.Errorf("Error: Could not locate correct number of Creates.\n"+
      "hasCreate='%v'\n", hasCreate)
  }

  if hasExclusive != 1 {
    t.Errorf("Error: Could not locate correct number of Exclusives.\n"+
      "hasExclusive='%v'\n", hasExclusive)
  }

}

func TestFileOpenConfig_GetFileOpenModes_02(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fileOpenModes := fOpCfg.GetFileOpenModes()

  if fileOpenModes == nil {
    t.Error("Error: Returned fileOpenModes is nil!\n")
  }

  if len(fileOpenModes) == 0 {
    t.Error("Error: Returned fileOpenModes has Zero Length!\n")
    return
  }

  if len(fileOpenModes) != 1 {
    t.Errorf("Error: Returned fileOpenModes Length is NOT '1' !\n"+
      "Length='%v'\n", len(fileOpenModes))
    return
  }

  if fileOpenModes[0] != FOpenMode.ModeNone() {
    t.Error("Error: Expected fileOpenModes[0] == FOpenMode.ModeNone().\n" +
      "It is NOT!\n")
  }
}

func TestFileOpenConfig_GetFileOpenModes_03(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fileOpenModes := fOpCfg.GetFileOpenModes()

  if fileOpenModes == nil {
    t.Error("Error: Returned fileOpenModes is nil!\n")
    return
  }

  if len(fileOpenModes) == 0 {
    t.Error("Error: Returned fileOpenModes has Zero Length!\n")
    return
  }

  if len(fileOpenModes) != 1 {
    t.Errorf("Error: Returned fileOpenModes Length is NOT '1' !\n"+
      "Length='%v'\n", len(fileOpenModes))
  }

  if fileOpenModes[0] != FOpenMode.ModeNone() {
    t.Error("Error: Expected fileOpenModes[0] == FOpenMode.ModeNone().\n" +
      "It is NOT!\n")
  }

}

func TestFileOpenConfig_GetFileOpenModes_04(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes = nil

  fileOpenModes := fOpCfg.GetFileOpenModes()

  if fileOpenModes == nil {
    t.Error("Error: Returned fileOpenModes is nil!\n")
    return
  }

  if len(fileOpenModes) == 0 {
    t.Error("Error: Returned fileOpenModes has Zero Length!\n")
    return
  }

  if len(fileOpenModes) != 1 {
    t.Errorf("Error: Returned fileOpenModes Length is NOT '1' !\n"+
      "Length='%v'\n", len(fileOpenModes))
  }

  if fileOpenModes[0] != FOpenMode.ModeNone() {
    t.Error("Error: Expected fileOpenModes[0] == FOpenMode.ModeNone().\n" +
      "It is NOT!\n")
  }

}

func TestFileOpenConfig_GetFileOpenTextString_01(t *testing.T) {

  fOpenCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeExclusive())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  txt := fOpenCfg.GetFileOpenNarrativeText()

  if strings.Index(txt, "ReadWrite") == -1 {
    t.Error("Error: Could not locate 'ReadWrite' in FileOpen Text!\n")
  }

  if strings.Index(txt, "Create") == -1 {
    t.Error("Error: Could not locate 'Create' in FileOpen Text!\n")
  }

  if strings.Index(txt, "Exclusive") == -1 {
    t.Error("Error: Could not locate 'Exclusive' in FileOpen Text!\n")
  }

}

func TestFileOpenConfig_GetFileOpenTextString_02(t *testing.T) {

  fOpenCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpenCfg.fileOpenModes = nil

  txt := fOpenCfg.GetFileOpenNarrativeText()

  if strings.Index(txt, "ReadWrite") == -1 {
    t.Error("Error: Could not locate 'ReadWrite' in FileOpen Text!\n")
  }

  if strings.Index(txt, "None") == -1 {
    t.Error("Error: Could not locate 'None' in FileOpen Text!\n")
  }

}

func TestFileOpenConfig_GetFileOpenType_01(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadWrite(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpenType := fOpCfg.GetFileOpenType()

  if fOpenType != FOpenType.TypeReadWrite() {
    t.Errorf("Error: Expected fOpenType='ReadWrite'.\n" +
      "Instead, fOpenType='%v'\nstring='%s'\n", fOpenType, fOpenType.String())
  }

}

func TestFileOpenConfig_GetFileOpenType_02(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes = nil

  fOpenType := fOpCfg.GetFileOpenType()

  if fOpenType != FOpenType.TypeReadOnly() {
    t.Errorf("Error: Expected fOpenType='ReadOnly'.\n" +
      "Instead, fOpenType='%v'\nstring='%s'\n", fOpenType, fOpenType.String())
  }

}

func TestFileOpenConfig_IsValid_01(t *testing.T) {

  fOpCfg := FileOpenConfig{}

  err := fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg is uninitialized.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_02(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.IsValid()

  if err != nil {
    t.Errorf("Error: Expected no error returned from IsValid().\n" +
      "However, an error was returned!\nError='%v'\n", err.Error())
  }

}

func TestFileOpenConfig_IsValid_03(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeNone(),
    FOpenMode.ModeAppend(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeExclusive())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg File Type=None and multiple Modes.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_04(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenType = FileOpenType(-99)

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg File Type=-99.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_05(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes[0] = FileOpenMode(9999999)

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg File Type=-99.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_06(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes = nil

  err = fOpCfg.IsValid()

  if err != nil {
    t.Errorf("Expected NO ERROR RETURN from IsValid().\n"+
      "However, AN ERROR WAS RETURNED!\nError='%v'\n", err.Error())
  }

}

func TestFileOpenConfig_IsValid_07(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeAppend(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes[0] = FileOpenMode(9999)

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg contained an invalid File Mode.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_08(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeAppend(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes[1] = FOpenMode.ModeNone()

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg has multiple File Modes one of which is 'None'.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_09(t *testing.T) {

  fOpCfg, err := FileOpenConfig{}.New(
    FOpenType.TypeReadOnly(),
    FOpenMode.ModeAppend(),
    FOpenMode.ModeCreate(),
    FOpenMode.ModeTruncate())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  fOpCfg.fileOpenModes[1] = FOpenMode.ModeNone()

  err = fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg has multiple File Modes one of which is 'None'.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_IsValid_10(t *testing.T) {

  fOpCfg := FileOpenConfig{}

  fOpCfg.fileOpenType = FOpenType.TypeNone()

  fOpCfg.fileOpenModes = make([]FileOpenMode, 1)

  fOpCfg.fileOpenModes[0] = FOpenMode.ModeCreate()

  err := fOpCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpCfg.IsValid()\n" +
      "because fOpCfg has File Type='None' and fileOpenModes = ModeCreate.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_SetFileOpenType_01(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpStatus.SetFileOpenType(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by SetFileOpenType{}.New().\n" +
      "Error='%v' \n", err.Error())
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_SetFileOpenType_02(t *testing.T) {

  expectedFOpenCode := os.O_RDWR

  fOpStatus := FileOpenConfig{}

  err := fOpStatus.SetFileOpenType(FOpenType.TypeReadWrite())

  if err != nil {
    t.Errorf("Error returned by SetFileOpenType{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  err = fOpStatus.SetFileOpenModes(FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by fOpStatus." +
      "SetFileOpenModes(FOpenMode.ModeNone()).\n"+
      "Error='%v' \n", err.Error())
    return
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
    return
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_SetFileOpenType_03(t *testing.T) {

  fOpStatus := FileOpenConfig{}

  err := fOpStatus.SetFileOpenType(FileOpenType(-99))

  if err == nil {
    t.Error("Expected an error return from fOpStatus.\n" +
      "SetFileOpenType(FileOpenType(-99)) because FileType== -99.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}

func TestFileOpenConfig_SetFileOpenType_04(t *testing.T) {

  expectedFOpenType := FOpenType.TypeNone()

  fOpCfg := FileOpenConfig{}

  err := fOpCfg.SetFileOpenType(expectedFOpenType)

  if err != nil {
    t.Errorf("Error returned by SetFileOpenType{}.New().\n" +
      "Error='%v' \n", err.Error())
    return
  }

  actualFileOpenType := fOpCfg.GetFileOpenType()

  if expectedFOpenType != actualFileOpenType {
    t.Errorf("Error: Expected File Open Type='%v'.\n"+
      "Instead, actual File Open Type='%v' \n",
      expectedFOpenType.String(), actualFileOpenType.String())
  }

}

func TestFileOpenConfig_SetFileOpenModes_01(t *testing.T) {

  expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_CREATE

  fOpStatus, err := FileOpenConfig{}.New(FOpenType.TypeNone(),
    FOpenMode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.New(FOpenType.TypeNone()," +
      "FOpenMode.ModeNone()).\n" +
      "Error='%v'\n", err.Error())
    return
  }

  err = fOpStatus.SetFileOpenType(FOpenType.TypeWriteOnly())

  if err != nil {
    t.Errorf("Error returned by fOpStatus.SetFileOpenType(FOpenType.TypeWriteOnly()).\n" +
      "Error='%v' \n", err.Error())
  }

  err = fOpStatus.SetFileOpenModes(FOpenMode.ModeAppend(), FOpenMode.ModeCreate())

  if err != nil {
    t.Errorf("Error returned by fOpStatus.SetFileOpenModes(FOpenMode.ModeAppend(), "+
      "FOpenMode.ModeCreate()).\n" +
      "Error='%v' \n",
      err.Error())
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_SetFileOpenModes_02(t *testing.T) {

  expectedFOpenCode := os.O_WRONLY | os.O_APPEND | os.O_CREATE

  fOpStatus := FileOpenConfig{}

  fOpStatus.fileOpenType = FOpenType.TypeWriteOnly()

  fOpStatus.fileOpenModes = nil

  err := fOpStatus.SetFileOpenModes(FOpenMode.ModeAppend(), FOpenMode.ModeCreate())

  if err != nil {
    t.Errorf("Error returned by fOpStatus.SetFileOpenModes(FOpenMode.ModeAppend(), "+
      "FOpenMode.ModeCreate()).\n" +
      "Error='%v' \n",
      err.Error())
  }

  actualFOpenCode, err := fOpStatus.GetCompositeFileOpenCode()

  if err != nil {
    t.Errorf("Error returned by FileOpenConfig{}.GetCompositeFileOpenCode().\n"+
      "Error='%v' \n", err.Error())
  }

  if expectedFOpenCode != actualFOpenCode {
    t.Errorf("Error: Expected File Open Code='%v'.\n" +
      "Instead, actual File Open Code='%v' \n",
      expectedFOpenCode, actualFOpenCode)
  }

}

func TestFileOpenConfig_SetFileOpenModes_03(t *testing.T) {

  fOpnCfg := FileOpenConfig{}

  fOpnCfg.fileOpenType = FOpenType.TypeWriteOnly()

  err := fOpnCfg.SetFileOpenModes()

  if err != nil {
    t.Errorf("Error returned by fOpnCfg.SetFileOpenModes()\n"+
      "Error='%v' \n",
      err.Error())
  }

  err = fOpnCfg.IsValid()

  if err == nil {
    t.Error("Expected an error return from fOpnCfg.IsValid()\n" +
      "because file modes config shows as uninitialized.\n" +
      "However, NO ERROR WAS RETURNED!\n")
  }

}
