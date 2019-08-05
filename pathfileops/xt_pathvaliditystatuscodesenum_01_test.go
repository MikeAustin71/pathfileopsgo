package pathfileops

import "testing"

func TestPathValidityStatusCode_Invalid_01(t *testing.T) {

  status := PathValidStatus.Invalid()

  intStatus := int(status)

  if intStatus != 0 {
    t.Errorf("Error: Expected PathValidStatus.Invalid()=='0'.\n" +
      "Instead PathValidStatus.Invalid()=='%v'\n", intStatus)
  }
}

func TestPathValidityStatusCode_Invalid_02(t *testing.T) {

  status := PathValidStatus.Invalid()

  statusStr := status.String()

  if statusStr != "Invalid" {
    t.Errorf("Error: For 'PathValidStatus.Invalid()':\n" +
      "Expected PathExistsStatus.String()=='Invalid'.\n" +
      "Instead PathExistsStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathValidityStatusCode_Invalid_03(t *testing.T) {

  status := PathValidStatus.Invalid()

  statusValue := status.Value()

  if int(statusValue) != 0 {
    t.Errorf("Error: For 'PathValidStatus.Invalid()':\n" +
      "Expected status.Value()=='0'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathValidityStatusCode_Invalid_04(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("Invalid", true)

  if err != nil {
    t.Errorf("Error returned by PathValidityStatusCode(0)." +
      "ParseString(\"Invalid\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathValidStatus.Invalid()':\n" +
      "Expected ParseString()=='0'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathValidityStatusCode_Invalid_05(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("invalid", false)

  if err != nil {
    t.Errorf("Error returned by PathValidityStatusCode(0)." +
      "ParseString(\"invalid\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathValidStatus.Invalid()':\n" +
      "Expected lower case ParseString()=='0'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

