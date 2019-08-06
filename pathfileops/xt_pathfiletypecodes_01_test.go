package pathfileops

import "testing"

func TestPathFileTypeCode_None_01(t *testing.T) {

  status := PathFileType.None()

  intStatus := int(status)

  if intStatus != 0 {
    t.Errorf("Error: Expected PathFileType.None()=='0'.\n" +
      "Instead PathFileType.None()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_None_02(t *testing.T) {

  status := PathFileType.None()

  statusStr := status.String()

  if statusStr != "None" {
    t.Errorf("Error: For 'PathFileType.None()':\n" +
      "Expected PathFileType.String()=='None'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_None_03(t *testing.T) {

  status := PathFileType.None()

  statusValue := status.StatusValue()

  if int(statusValue) != 0 {
    t.Errorf("Error: For 'PathFileType.None()':\n" +
      "Expected status.Value()=='0'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_None_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("None", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"None\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathFileType.None()':\n" +
      "Expected ParseString()=='0'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_None_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("none", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"none\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathFileType.None()':\n" +
      "Expected lower case ParseString()=='0'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

