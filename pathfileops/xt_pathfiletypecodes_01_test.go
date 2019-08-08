package pathfileops

import "testing"

func TestPathFileTypeCode_File_01(t *testing.T) {

  status := PathFileType.File()

  intStatus := int(status)

  if intStatus != 3 {
    t.Errorf("Error: Expected PathFileType.File()=='3'.\n" +
      "Instead PathFileType.File()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_File_02(t *testing.T) {

  status := PathFileType.File()

  statusStr := status.String()

  if statusStr != "File" {
    t.Errorf("Error: For 'PathFileType.File()':\n" +
      "Expected PathFileType.String()=='File'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_File_03(t *testing.T) {

  status := PathFileType.File()

  statusValue := status.StatusValue()

  if int(statusValue) != 3 {
    t.Errorf("Error: For 'PathFileType.File()':\n" +
      "Expected status.Value()=='3'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_File_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("File", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"File\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 3 {
    t.Errorf("Error: For 'PathFileType.File()':\n" +
      "Expected ParseString()=='3'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_File_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("file", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"file\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 3 {
    t.Errorf("Error: For 'PathFileType.File()':\n" +
      "Expected lower case ParseString()=='3'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_Indeterminate_01(t *testing.T) {

  status := PathFileType.Indeterminate()

  intStatus := int(status)

  if intStatus != 5 {
    t.Errorf("Error: Expected PathFileType.Indeterminate()=='5'.\n" +
      "Instead PathFileType.Indeterminate()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_Indeterminate_02(t *testing.T) {

  status := PathFileType.Indeterminate()

  statusStr := status.String()

  if statusStr != "Indeterminate" {
    t.Errorf("Error: For 'PathFileType.Indeterminate()':\n" +
      "Expected PathFileType.String()=='Indeterminate'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_Indeterminate_03(t *testing.T) {

  status := PathFileType.Indeterminate()

  statusValue := status.StatusValue()

  if int(statusValue) != 5 {
    t.Errorf("Error: For 'PathFileType.Indeterminate()':\n" +
      "Expected status.Value()=='5'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_Indeterminate_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("Indeterminate", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"Indeterminate\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 5 {
    t.Errorf("Error: For 'PathFileType.Indeterminate()':\n" +
      "Expected ParseString()=='5'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_Indeterminate_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("indeterminate", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"indeterminate\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 5 {
    t.Errorf("Error: For 'PathFileType.Indeterminate()':\n" +
      "Expected lower case ParseString()=='5'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

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

func TestPathFileTypeCode_ParseString_01(t *testing.T) {

  statusStr, err := PathFileTypeCode(0).ParseString("Volume()", true)

  if err != nil {
    t.Errorf("Error returned by ParseString(\"Volume()\", true)\n" +
      "Error='%v'\n", statusStr)
  }

}

func TestPathFileTypeCode_ParseString_02(t *testing.T) {

  _, err := PathFileTypeCode(0).ParseString("xy", true)

  if err == nil {
    t.Error("Expected an error return from ParseString(\"xy\", true)\n" +
      "because \"xy\" is an invalid string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestPathFileTypeCode_ParseString_03(t *testing.T) {

  _, err := PathFileTypeCode(0).ParseString("xy", false)

  if err == nil {
    t.Error("Expected an error return from ParseString(\"xy\", false)\n" +
      "because \"xy\" is an invalid string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }
}

func TestPathFileTypeCode_Path_01(t *testing.T) {

  status := PathFileType.Path()

  intStatus := int(status)

  if intStatus != 1 {
    t.Errorf("Error: Expected PathFileType.Path()=='1'.\n" +
      "Instead PathFileType.Path()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_Path_02(t *testing.T) {

  status := PathFileType.Path()

  statusStr := status.String()

  if statusStr != "Path" {
    t.Errorf("Error: For 'PathFileType.Path()':\n" +
      "Expected PathFileType.String()=='Path'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_Path_03(t *testing.T) {

  status := PathFileType.Path()

  statusValue := status.StatusValue()

  if int(statusValue) != 1 {
    t.Errorf("Error: For 'PathFileType.Path()':\n" +
      "Expected status.Value()=='1'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_Path_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("Path", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"Path\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathFileType.Path()':\n" +
      "Expected ParseString()=='1'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_Path_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("path", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"path\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathFileType.Path()':\n" +
      "Expected lower case ParseString()=='1'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_PathFile_01(t *testing.T) {

  status := PathFileType.PathFile()

  intStatus := int(status)

  if intStatus != 2 {
    t.Errorf("Error: Expected PathFileType.PathFile()=='2'.\n" +
      "Instead PathFileType.PathFile()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_PathFile_02(t *testing.T) {

  status := PathFileType.PathFile()

  statusStr := status.String()

  if statusStr != "PathFile" {
    t.Errorf("Error: For 'PathFileType.PathFile()':\n" +
      "Expected PathFileType.String()=='PathFile'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_PathFile_03(t *testing.T) {

  status := PathFileType.PathFile()

  statusValue := status.StatusValue()

  if int(statusValue) != 2 {
    t.Errorf("Error: For 'PathFileType.PathFile()':\n" +
      "Expected status.Value()=='2'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_PathFile_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("PathFile", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"PathFile\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 2 {
    t.Errorf("Error: For 'PathFileType.PathFile()':\n" +
      "Expected ParseString()=='2'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_PathFile_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("pathfile", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"pathfile\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 2 {
    t.Errorf("Error: For 'PathFileType.PathFile()':\n" +
      "Expected lower case ParseString()=='2'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_String(t *testing.T) {

  status := PathFileTypeCode(-99)

  statusStr := status.String()

  if statusStr != "" {
    t.Errorf("Expected an empty string return from status.String()\n" +
      "because 'status' is invalid.\n" +
      "Instead, status.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_Volume_01(t *testing.T) {

  status := PathFileType.Volume()

  intStatus := int(status)

  if intStatus != 4 {
    t.Errorf("Error: Expected PathFileType.Volume()=='4'.\n" +
      "Instead PathFileType.Volume()=='%v'\n", intStatus)
  }
}

func TestPathFileTypeCode_Volume_02(t *testing.T) {

  status := PathFileType.Volume()

  statusStr := status.String()

  if statusStr != "Volume" {
    t.Errorf("Error: For 'PathFileType.Volume()':\n" +
      "Expected PathFileType.String()=='Volume'.\n" +
      "Instead PathFileType.String()=='%v'\n", statusStr)
  }
}

func TestPathFileTypeCode_Volume_03(t *testing.T) {

  status := PathFileType.Volume()

  statusValue := status.StatusValue()

  if int(statusValue) != 4 {
    t.Errorf("Error: For 'PathFileType.Volume()':\n" +
      "Expected status.Value()=='4'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathFileTypeCode_Volume_04(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("Volume", true)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"Volume\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 4 {
    t.Errorf("Error: For 'PathFileType.Volume()':\n" +
      "Expected ParseString()=='4'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathFileTypeCode_Volume_05(t *testing.T) {

  statusCode, err :=
    PathFileTypeCode(0).ParseString("volume", false)

  if err != nil {
    t.Errorf("Error returned by PathFileTypeCode(0)." +
      "ParseString(\"volume\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 4 {
    t.Errorf("Error: For 'PathFileType.Volume()':\n" +
      "Expected lower case ParseString()=='4'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}
