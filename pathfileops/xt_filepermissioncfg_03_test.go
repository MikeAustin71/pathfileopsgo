package pathfileops

import (
  "os"
  "strconv"
  "testing"
)

func TestFilePermissionConfig_IsValid_01(t *testing.T) {
  // expectedTextCode := "drwxrwxrwx"
  fh := FileHelper{}

  // drwxrwxrwx   20000000777

  intFMode := fh.ConvertOctalToDecimal(20000000777)

  osFMode := os.FileMode(intFMode)

  fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
      "Error='%v' ", err.Error())
  }

  fPerm.isInitialized = false

  err = fPerm.IsValid()

  if err == nil {
    t.Error("Expected an error to be returned by fPerm.IsValid() because " +
      "fPerm has not been initialized. NO ERROR RETURNED!")
  }

}

func TestFilePermissionConfig_IsValid_02(t *testing.T) {

  // expectedTextCode := "drwxrwxrwx"
  fh := FileHelper{}

  // drwxrwxrwx   20000000777

  intFMode := fh.ConvertOctalToDecimal(20000000777)

  osFMode := os.FileMode(intFMode)

  fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
      "Error='%v' ", err.Error())
  }

  fPerm.fileMode = os.FileMode(01000)

  err = fPerm.IsValid()

  if err == nil {
    t.Error("Expected an error to be returned by fPerm.IsValid() because " +
      "fPerm has an INVALID Entry Type. NO ERROR RETURNED!")
  }

}

func TestFilePermissionConfig_IsValid_03(t *testing.T) {

  // expectedTextCode := "drwxrwxrwx"
  fh := FileHelper{}

  // drwxrwxrwx   20000000777

  intFMode := fh.ConvertOctalToDecimal(20000000777)

  osFMode := os.FileMode(intFMode)

  fPerm, err := FilePermissionConfig{}.NewByFileMode(osFMode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode(osFMode). "+
      "Error='%v' ", err.Error())
  }

  err = fPerm.IsValid()

  if err != nil {
    t.Errorf("Expected no error to be returned by fPerm.IsValid(). "+
      "However, an error was returned. Error='%v' ", err.Error())
  }

}

func TestFilePermissionConfig_New_01(t *testing.T) {

  permissionStr := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.New(permissionStr)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.New(permissionStr) "+
      "Error='%v' ", err.Error())
  }

  actualTextCode, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPermCfg.GetPermissionTextCode() "+
      "Error='%v' ", err.Error())
  }

  if permissionStr != actualTextCode {
    t.Errorf("Error: Expected actual text code='%v' .Instead, "+
      "actual text code='%v'",
      permissionStr, actualTextCode)
  }

}

func TestFilePermissionConfig_New_02(t *testing.T) {

  permissionStr := "xvumnoqade"

  _, err := FilePermissionConfig{}.New(permissionStr)

  if err == nil {
    t.Error("Expected error return from FilePermissionConfig{}.New(permissionStr) " +
      "because of invalid permissionStr. NO ERROR WAS RETURNED!")
  }

}

func TestFilePermissionConfig_NewByComponents_01(t *testing.T) {

  entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeNone()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "-rwxrwxrwx"
  permissionStr := "-rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
      "permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_NewByComponents_02(t *testing.T) {

  entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "drwxrwxrwx"
  permissionStr := "rwxrwxrwx"

  fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
      "permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_NewByComponents_03(t *testing.T) {

  //  ModeSetuid()      os.ModeSetuid         "u" setuid
  entryType, err := OsFilePermissionCode(0).GetNewFromFileMode(OsFilePermCode.ModeSetuid())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeSetuid()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "urw-rw-rw-"
  permissionStr := "rw-rw-rw-"

  fPermCfg, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByComponents(entryType, "+
      "permissionStr). entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_NewByComponents_04(t *testing.T) {

  // Bad Entry Type Code
  entryType := OsFilePermissionCode(999)

  permissionStr := "rw-rw-rw-"

  _, err := FilePermissionConfig{}.NewByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected error return from bad entry type code 999. " +
      "However, NO ERROR WAS RETURNED! ")
  }

}

func TestFilePermissionConfig_NewByFileMode_01(t *testing.T) {

  expectedFileMode := os.FileMode(0666)

  fPerm, err := FilePermissionConfig{}.NewByFileMode(expectedFileMode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}.NewByFileMode"+
      "(os.FileMode(0666)). Error='%v' ", err.Error())
  }

  actualFileMode, err := fPerm.GetCompositePermissionMode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetCompositePermissionMode()"+
      "Error='%v' ", err.Error())
  }

  if expectedFileMode != actualFileMode {
    t.Errorf("Error: Expected actual file mode octal value = '%s' Instead, "+
      "actual file mode octal value= '%s' ",
      strconv.FormatInt(int64(expectedFileMode), 8),
      strconv.FormatInt(int64(actualFileMode), 8))
  }

}

func TestFilePermissionConfig_NewByFileMode_02(t *testing.T) {

  expectedFileMode := os.FileMode(9236)

  _, err := FilePermissionConfig{}.NewByFileMode(expectedFileMode)

  if err == nil {
    t.Error("Expected error return from FilePermissionConfig{}.NewByFileMode" +
      "(expectedFileMode) because of invalid FileMode. NO ERROR WAS RETURNED!!! ")
  }

}

func TestFilePermissionConfig_NewByOctalDigits_01(t *testing.T) {

  expectedTextCode := "-rw-rw-rw-"
  octalCode := int(666)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_02(t *testing.T) {

  expectedTextCode := "drw-rw-rw-"
  octalCode := int(20000000666)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_03(t *testing.T) {

  expectedTextCode := "--w--w--w-"
  octalCode := int(222)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_04(t *testing.T) {

  expectedTextCode := "-r--r--r--"
  octalCode := int(444)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_05(t *testing.T) {

  expectedTextCode := "dr--r--r--"
  octalCode := int(20000000444)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_06(t *testing.T) {

  expectedTextCode := "d-w--w--w-"
  octalCode := int(20000000222)

  fPerm, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by FilePermissionConfig{}."+
      "NewByOctalDigits(octalCode) Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_NewByOctalDigits_07(t *testing.T) {

  octalCode := int(12577)

  _, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err == nil {
    t.Error("Expected an error to be returned by FilePermissionConfig{}." +
      "NewByOctalDigits(octalCode) because of invalid octalCode.  " +
      "NO ERROR WAS RETURNED!")
  }

}

func TestFilePermissionConfig_NewByOctalDigits_08(t *testing.T) {

  octalCode := int(12577)

  _, err := FilePermissionConfig{}.NewByOctalDigits(octalCode)

  if err == nil {
    t.Error("Expected an error to be returned by FilePermissionConfig{}." +
      "NewByOctalDigits(octalCode) because of invalid octalCode.  " +
      "NO ERROR WAS RETURNED!")
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_01(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeNone()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "-rwxrwxrwx"
  permissionStr := "-rwxrwxrwx"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_02(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "drwxrwxrwx"
  permissionStr := "rwxrwxrwx"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_03(t *testing.T) {

  //  ModeSetuid()      os.ModeSetuid         "u" setuid
  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeSetuid())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeSetuid()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "urw-rw-rw-"
  permissionStr := "rw-rw-rw-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_04(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeNone()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "-rw-rw-rw-"
  permissionStr := "rw-rw-rw-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_05(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeNone()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "-r--r--r--"
  permissionStr := "r--r--r--"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_06(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeNone())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeNone()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "--w--w--w-"
  permissionStr := "-w--w--w-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_07(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "drw-rw-rw-"
  permissionStr := "rw-rw-rw-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_08(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "dr--r--r--"
  permissionStr := "r--r--r--"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_09(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  expectedPermissionTxt := "d-w--w--w-"
  permissionStr := "-w--w--w-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err != nil {
    t.Errorf("Error returned by fPermCfg.SetFileModeByComponents(entryType, permissionStr). "+
      "entrType='%s' permissionStr='%s' Error='%v' ",
      entryType.String(), permissionStr, err.Error())
  }

  actualPermissionTxt, err := fPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fpCfg.GetPermissionTextCode(). "+
      "Error='%v' ", err.Error())
  }

  if expectedPermissionTxt != actualPermissionTxt {
    t.Errorf("Error: Expected Permission Text Code='%v'. Instead, "+
      "Actual Permission Text Code='%v' ",
      expectedPermissionTxt, actualPermissionTxt)
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_10(t *testing.T) {

  // Bad Entry Type Code
  entryType := OsFilePermissionCode(999)
  permissionStr := "rw-rw-rw-"

  fPermCfg := FilePermissionConfig{}

  err := fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected error return from bad entry type code 999. " +
      "However, NO ERROR WAS RETURNED! ")
  }
}

func TestFilePermissionConfig_SetFileModeByComponents_11(t *testing.T) {

  entryType := OsFilePermissionCode(OsFilePermCode.ModeNone())
  // Bad Permission String
  permissionStr := "rZ-rz-rz-"

  fPermCfg := FilePermissionConfig{}

  err := fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected error return from bad permission string. " +
      "However, NO ERROR WAS RETURNED! ")
  }
}

func TestFilePermissionConfig_SetFileModeByComponents_12(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  permissionStr := "-w--w--w--w--w--w--w--w--"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected an error from fPermCfg.SetFileModeByComponents(entryType, " +
      "permissionStr). because permission string was longer than 10-characters. " +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFilePermissionConfig_SetFileModeByComponents_13(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }
  permissionStr := "-w-zzz-w-"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected an error from fPermCfg.SetFileModeByComponents(entryType, " +
      "permissionStr). because permission string invalid group characters. " +
      "However, NO ERROR WAS RETURNED!!")
  }
}

func TestFilePermissionConfig_SetFileModeByComponents_14(t *testing.T) {

  entryType, err := OsFilePermCode.GetNewFromFileMode(OsFilePermCode.ModeDir())

  if err != nil {
    t.Errorf("Error returned by OsFilePermCode.GetNewFromFileMode("+
      "OsFilePermCode.ModeDir()). "+
      "Error='%v' ", err.Error())
  }

  permissionStr := "-w--w-ZZZ"

  fPermCfg := FilePermissionConfig{}

  err = fPermCfg.SetFileModeByComponents(entryType, permissionStr)

  if err == nil {
    t.Error("Expected an error from fPermCfg.SetFileModeByComponents(entryType, " +
      "permissionStr). because permission string invalid 'other' group characters. " +
      "However, NO ERROR WAS RETURNED!!")
  }
}

func TestFilePermissionConfig_SetFileModeByOctalDigits_01(t *testing.T) {

  expectedTextCode := "-rw-rw-rw-"
  octalCode := int(666)

  fPerm := FilePermissionConfig{}

  err := fPerm.SetFileModeByOctalDigits(octalCode)

  if err != nil {
    t.Errorf("Error returned by fPerm.SetFileModeByOctalDigits(octalCode). "+
      "Error='%v'", err.Error())
  }

  actualTextCode, err := fPerm.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by fPerm.GetPermissionTextCode(). "+
      "Error='%v'", err.Error())
  }

  if expectedTextCode != actualTextCode {
    t.Errorf("Error: Expected permission text code='%v'. Instead, "+
      "permission text code='%v' ", expectedTextCode, actualTextCode)
  }

}

func TestFilePermissionConfig_SetFileModeByOctalDigits_02(t *testing.T) {

  octalCode := int(12666)

  fPerm := FilePermissionConfig{}

  err := fPerm.SetFileModeByOctalDigits(octalCode)

  if err == nil {
    t.Error("Expected error to be returned by fPerm.SetFileMode" +
      "ByOctalDigits(octalCode) because octal code was invalid! " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFilePermissionConfig_SetFileModeByTextCode_01(t *testing.T) {

  textCode := "-rwxrwxrwx"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_02(t *testing.T) {

  textCode := "drwxrwxrwx"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_03(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_04(t *testing.T) {

  textCode := "drw-rw-rw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_05(t *testing.T) {

  textCode := "-rwx------"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_06(t *testing.T) {

  textCode := "-rwxrwx---"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_07(t *testing.T) {

  textCode := "---x--x--x"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_08(t *testing.T) {

  textCode := "--w--w--w-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_09(t *testing.T) {

  textCode := "--wx-wx-wx"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_10(t *testing.T) {

  textCode := "-r--r--r--"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_11(t *testing.T) {

  textCode := "-r-xr-xr-x"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_12(t *testing.T) {

  textCode := "-rw-rw-rw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_13(t *testing.T) {

  textCode := "-rwxr-----"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_14(t *testing.T) {

  textCode := "drw-rw-rw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_15(t *testing.T) {

  textCode := "----------"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_16(t *testing.T) {

  textCode := "d---------"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err != nil {
    t.Errorf("Error returned by fpCfg.SetFileModeByTextCode(textCode). "+
      "Error='%v' ", err.Error())
  }

  fileMode, err := fpCfg.GetCompositePermissionMode()

  if textCode != fileMode.String() {
    t.Errorf("Error: Expected File Mode text = '%v'. Instead, text = '%v' .",
      textCode, fileMode.String())
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_17(t *testing.T) {

  textCode := "-rwxrwxrwxrwxrwx"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err == nil {
    t.Error("Expected error to be returned by fpCfg.SetFileModeBy" +
      "TextCode(textCode) because input text was longer than 10-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_18(t *testing.T) {

  textCode := "-rwx"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err == nil {
    t.Error("Expected error to be returned by fpCfg.SetFileModeBy" +
      "TextCode(textCode) because input text was less than 10-characters. " +
      "However, NO ERROR WAS RETURNED!")
  }
}

func TestFilePermissionConfig_SetFileModeByTextCode_19(t *testing.T) {

  textCode := "-ZZZrw-rw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err == nil {
    t.Error("Expected error to be returned by " +
      "fpCfg.SetFileModeByTextCode(textCode). because owner characters are invalid. " +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_20(t *testing.T) {

  textCode := "-rw-ZZZrw-"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err == nil {
    t.Error("Expected error to be returned by " +
      "fpCfg.SetFileModeByTextCode(textCode). because group characters are invalid. " +
      "However, NO ERROR WAS RETURNED!!")
  }

}

func TestFilePermissionConfig_SetFileModeByTextCode_21(t *testing.T) {

  textCode := "-rw-rw-zZZ"

  fpCfg := FilePermissionConfig{}

  err := fpCfg.SetFileModeByTextCode(textCode)

  if err == nil {
    t.Error("Expected error to be returned by " +
      "fpCfg.SetFileModeByTextCode(textCode). because 'other' characters are invalid. " +
      "However, NO ERROR WAS RETURNED!!")
  }

}
