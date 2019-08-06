package pathfileops

import "testing"


func TestPathValidityStatusCode_EqualOperator_01(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1==status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "However, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_EqualOperator_02(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1==status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_EqualOperator_03(t *testing.T) {

  status1 := PathValidStatus.Invalid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1==status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Invalid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_EqualOperator_04(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1==status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_EqualOperator_05(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Valid()

  result := false

  if status1==status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_EqualOperator_06(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Valid()

  result := false

  if status1==status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}


func TestPathValidityStatusCode_GreaterThanOperator_01(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOperator_02(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOperator_03(t *testing.T) {

  status1 := PathValidStatus.Invalid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Invalid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOperator_04(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1 > status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the greater than operator (status1 > status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOperator_05(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Valid()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOperator_06(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1 > status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}


func TestPathValidityStatusCode_GreaterThanOrEqualOperator_01(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "However, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_02(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1>=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1>=status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_03(t *testing.T) {

  status1 := PathValidStatus.Invalid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Invalid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_04(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_05(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Valid()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_06(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Valid()

  result := false

  if status1>=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the equal operator (status1>=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_GreaterThanOrEqualOperator_07(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_01(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_02(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_03(t *testing.T) {

  status1 := PathValidStatus.Invalid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Invalid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_04(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1<=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1<=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_05(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Valid()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_06(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Valid()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_07(t *testing.T) {

  status1 := PathValidityStatusCode(99)

  status2 := PathValidStatus.Valid()

  result := false

  if status1<=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidityStatusCode(99) and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1<=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_LessThanOrEqualOperator_08(t *testing.T) {

  status1 := PathValidityStatusCode(-99)

  status2 := PathValidStatus.Valid()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidityStatusCode(-99) and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_01(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Unknown()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Unknown()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_02(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_03(t *testing.T) {

  status1 := PathValidStatus.Invalid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Invalid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_04(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Invalid()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Invalid()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_05(t *testing.T) {

  status1 := PathValidStatus.Valid()

  status2 := PathValidStatus.Valid()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathValidStatus.Valid() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathValidityStatusCode_NotEqualOperator_06(t *testing.T) {

  status1 := PathValidStatus.Unknown()

  status2 := PathValidStatus.Valid()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathValidStatus.Unknown() and " +
      "status2=PathValidStatus.Valid()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

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


func TestPathValidityStatusCode_Unknown_01(t *testing.T) {

  status := PathValidStatus.Unknown()

  intStatus := int(status)

  if intStatus != -1 {
    t.Errorf("Error: Expected PathValidStatus.Unknown()=='-1'.\n" +
      "Instead PathValidStatus.Unknown()=='%v'\n", intStatus)
  }
}

func TestPathValidityStatusCode_Unknown_02(t *testing.T) {

  status := PathValidStatus.Unknown()

  statusStr := status.String()

  if statusStr != "Unknown" {
    t.Errorf("Error: For 'PathValidStatus.Unknown()':\n" +
      "Expected PathValidStatus.String()=='Unknown'.\n" +
      "Instead PathValidStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathValidityStatusCode_Unknown_03(t *testing.T) {

  status := PathValidStatus.Unknown()

  statusValue := status.Value()

  if int(statusValue) != -1 {
    t.Errorf("Error: For 'PathValidStatus.Unknown()':\n" +
      "Expected status.Value()=='-1'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathValidityStatusCode_Unknown_04(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("Unknown", true)

  if err != nil {
    t.Errorf("Error returned by PathValidStatusCode(0)." +
      "ParseString(\"Unknown\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != -1 {
    t.Errorf("Error: For 'PathValidStatus.Unknown()':\n" +
      "Expected ParseString()=='-1'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathValidityStatusCode_Unknown_05(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("unknown", false)

  if err != nil {
    t.Errorf("Error returned by PathValidStatusCode(0)." +
      "ParseString(\"unknown\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != -1 {
    t.Errorf("Error: For 'PathValidStatus.Unknown()':\n" +
      "Expected lower case ParseString()=='-1'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathValidityStatusCode_Valid_01(t *testing.T) {

  status := PathValidStatus.Valid()

  intStatus := int(status)

  if intStatus != 1 {
    t.Errorf("Error: Expected PathValidStatus.Valid()=='1'.\n" +
      "Instead PathValidStatus.Valid()=='%v'\n", intStatus)
  }
}

func TestPathValidityStatusCode_Valid_02(t *testing.T) {

  status := PathValidStatus.Valid()

  statusStr := status.String()

  if statusStr != "Valid" {
    t.Errorf("Error: For 'PathValidStatus.Valid()':\n" +
      "Expected PathValidStatus.String()=='Valid'.\n" +
      "Instead PathValidStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathValidityStatusCode_Valid_03(t *testing.T) {

  status := PathValidStatus.Valid()

  statusValue := status.Value()

  if int(statusValue) != 1 {
    t.Errorf("Error: For 'PathValidStatus.Valid()':\n" +
      "Expected status.Value()=='1'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathValidityStatusCode_Valid_04(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("Valid", true)

  if err != nil {
    t.Errorf("Error returned by PathValidityStatusCode(0)." +
      "ParseString(\"Valid\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathValidStatus.Valid()':\n" +
      "Expected ParseString()=='1'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathValidityStatusCode_Valid_05(t *testing.T) {

  statusCode, err :=
    PathValidityStatusCode(0).ParseString("valid", false)

  if err != nil {
    t.Errorf("Error returned by PathValidityStatusCode(0)." +
      "ParseString(\"valid\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathValidStatus.Valid()':\n" +
      "Expected lower case ParseString()=='1'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathValidityStatusCode_Valid_06(t *testing.T) {

  _, err :=
    PathValidityStatusCode(0).ParseString("valid", true)

  if err == nil {
    t.Error("Expected an error return from PathValidityStatusCode(0)." +
      "ParseString(\"valid\", true)\n" +
      "because 'valid' is test as 'case sensitive'." +
      "However, NO ERROR WAS RETURNED!!!!\n")
  }

}
