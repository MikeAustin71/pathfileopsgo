package pathfileops

import (
  "strconv"
  "testing"
)

func TestFilePermissionConfig_CopyIn_01(t *testing.T) {

  textCode := "drwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := FilePermissionConfig{}

  fpCfg2.CopyIn(&fpCfg)

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }

}

func TestFilePermissionConfig_CopyIn_02(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := FilePermissionConfig{}

  fpCfg2.CopyIn(&fpCfg)

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }

}

func TestFilePermissionConfig_CopyOut_01(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_02(t *testing.T) {

  textCode := "drwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_03(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_04(t *testing.T) {

  textCode := "drw-rw-rw-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_05(t *testing.T) {

  textCode := "-r--r--r--"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_06(t *testing.T) {

  textCode := "dr--r--r--"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_07(t *testing.T) {

  textCode := "--w--w--w-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_08(t *testing.T) {

  textCode := "d-w--w--w-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_09(t *testing.T) {

  textCode := "---xr----x"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_CopyOut_10(t *testing.T) {

  textCode := "d--xr----x"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := fpCfg.CopyOut()

  actualTextCode, err := fpCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg2.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg to EQUAL fpCfg2. THEY ARE NOT EQUAL!")
  }
}

func TestFilePermissionConfig_Empty_01(t *testing.T) {

  testEmpty := FilePermissionConfig{}

  textCode := "drwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  actualTextCode, err := fpCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by #1 fpCfg.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  fpCfg.Empty()

  _, err = fpCfg.GetPermissionTextCode()

  if err == nil {
    t.Error("Expected an Uninitialized Error to be returned by #2 " +
      "fpCfg.GetPermissionTextCode(). NO ERROR WAS RETURNED!!!")
  }

  if !testEmpty.Equal(&fpCfg) {
    t.Error("Error: Expected testEmpty to EQUAL fpCfg. THEY ARE NOT EQUAL!")
  }

  if fpCfg.isInitialized == true {
    t.Error("Expected fpCfg.isInitialized==false. Instead, fpCfg.isInitialized==true")
  }

  if fpCfg.fileMode != 0 {
    t.Errorf("Expected fpCfg.fileMode=='0'. Instead, fpCfg.fileMode octal value =='%v' ",
      strconv.FormatInt(int64(fpCfg.fileMode), 8))
  }

}

func TestFilePermissionConfig_Empty_02(t *testing.T) {

  testEmpty := FilePermissionConfig{}

  textCode := "-rwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  actualTextCode, err := fpCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by #1 fpCfg.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  fpCfg.Empty()

  _, err = fpCfg.GetPermissionTextCode()

  if err == nil {
    t.Error("Expected an Uninitialized Error to be returned by #2 " +
      "fpCfg.GetPermissionTextCode(). NO ERROR WAS RETURNED!!!")
  }

  if !testEmpty.Equal(&fpCfg) {
    t.Error("Error: Expected testEmpty to EQUAL fpCfg. THEY ARE NOT EQUAL!")
  }

  if fpCfg.isInitialized == true {
    t.Error("Expected fpCfg.isInitialized==false. Instead, fpCfg.isInitialized==true")
  }

  if fpCfg.fileMode != 0 {
    t.Errorf("Expected fpCfg.fileMode=='0'. Instead, fpCfg.fileMode octal value =='%v' ",
      strconv.FormatInt(int64(fpCfg.fileMode), 8))
  }

}

func TestFilePermissionConfig_Empty_03(t *testing.T) {

  testEmpty := FilePermissionConfig{}

  textCode := "-rw-rw-rw-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  actualTextCode, err := fpCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by #1 fpCfg.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  fpCfg.Empty()

  _, err = fpCfg.GetPermissionTextCode()

  if err == nil {
    t.Error("Expected an Uninitialized Error to be returned by #2 " +
      "fpCfg.GetPermissionTextCode(). NO ERROR WAS RETURNED!!!")
  }

  if !testEmpty.Equal(&fpCfg) {
    t.Error("Error: Expected testEmpty to EQUAL fpCfg. THEY ARE NOT EQUAL!")
  }

  if fpCfg.isInitialized == true {
    t.Error("Expected fpCfg.isInitialized==false. Instead, fpCfg.isInitialized==true")
  }

  if fpCfg.fileMode != 0 {
    t.Errorf("Expected fpCfg.fileMode=='0'. Instead, fpCfg.fileMode octal value =='%v' ",
      strconv.FormatInt(int64(fpCfg.fileMode), 8))
  }

}

func TestFilePermissionConfig_Empty_04(t *testing.T) {

  testEmpty := FilePermissionConfig{}

  textCode := "d-w--w--w-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  actualTextCode, err := fpCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by #1 fpCfg.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  fpCfg.Empty()

  _, err = fpCfg.GetPermissionTextCode()

  if err == nil {
    t.Error("Expected an Uninitialized Error to be returned by #2 " +
      "fpCfg.GetPermissionTextCode(). NO ERROR WAS RETURNED!!!")
  }

  if !testEmpty.Equal(&fpCfg) {
    t.Error("Error: Expected testEmpty to EQUAL fpCfg. THEY ARE NOT EQUAL!")
  }

  if fpCfg.isInitialized == true {
    t.Error("Expected fpCfg.isInitialized==false. Instead, fpCfg.isInitialized==true")
  }

  if fpCfg.fileMode != 0 {
    t.Errorf("Expected fpCfg.fileMode=='0'. Instead, fpCfg.fileMode octal value =='%v' ",
      strconv.FormatInt(int64(fpCfg.fileMode), 8))
  }

}

func TestFilePermissionConfig_Empty_05(t *testing.T) {

  testEmpty := FilePermissionConfig{}

  textCode := "--w--w--w-"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  actualTextCode, err := fpCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by #1 fpCfg.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if textCode != actualTextCode {
    t.Errorf("Error: Expected text code ='%v'. Instead, text code='%v'. ",
      textCode, actualTextCode)
  }

  fpCfg.Empty()

  _, err = fpCfg.GetPermissionTextCode()

  if err == nil {
    t.Error("Expected an Uninitialized Error to be returned by #2 " +
      "fpCfg.GetPermissionTextCode(). NO ERROR WAS RETURNED!!!")
  }

  if !testEmpty.Equal(&fpCfg) {
    t.Error("Error: Expected testEmpty to EQUAL fpCfg. THEY ARE NOT EQUAL!")
  }

  if fpCfg.isInitialized == true {
    t.Error("Expected fpCfg.isInitialized==false. Instead, fpCfg.isInitialized==true")
  }

  if fpCfg.fileMode != 0 {
    t.Errorf("Expected fpCfg.fileMode=='0'. Instead, fpCfg.fileMode octal value =='%v' ",
      strconv.FormatInt(int64(fpCfg.fileMode), 8))
  }

}

func TestFilePermissionConfig_Equal_01(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg2 = FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  if !fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg == fpCfg2. Wrong, THEY ARE NOT EQUAL")
  }

  if !fpCfg2.Equal(&fpCfg) {
    t.Error("Error: Expected fpCfg2 == fpCfg. Wrong, THEY ARE NOT EQUAL")
  }

  textCode2 := "-rwxrwxrw-"

  fpCfg2, err = FilePermissionConfig{}.New(textCode2)

  if err != nil {
    t.Errorf("Error returned by fpCfg2 = FilePermissionConfig{}.New(textCode2). "+
      "textCode2='%v' Error='%v'", textCode2, err.Error())
  }

  if fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg != fpCfg2. Wrong, THEY ARE EQUAL!")
  }

  if fpCfg2.Equal(&fpCfg) {
    t.Error("Error: Expected fpCfg2 != fpCfg. Wrong, THEY ARE EQUAL!")
  }

}

func TestFilePermissionConfig_Equal_02(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fpCfg, err := FilePermissionConfig{}.New(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg = FilePermissionConfig{}.New(textCode). "+
      "textCode='%v' Error='%v'", textCode, err.Error())
  }

  fpCfg2 := FilePermissionConfig{}

  if fpCfg.Equal(&fpCfg2) {
    t.Error("Error: Expected fpCfg != fpCfg2. Wrong, THEY ARE EQUAL!")
  }

  if fpCfg2.Equal(&fpCfg) {
    t.Error("Error: Expected fpCfg2 != fpCfg. Wrong, THEY ARE EQUAL!")
  }

}
