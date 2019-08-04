package pathfileops

import "testing"

func TestPathExistsStatusCode_DoesNotExist_01(t *testing.T) {
  
  status := PathExistsStatus.DoesNotExist()
  
  intStatus := int(status)
  
  if intStatus != 0 {
    t.Errorf("Error: Expected PathExistsStatus.DoesNotExist()=='0'.\n" +
      "Instead PathExistsStatus.DoesNotExist()=='%v'\n", intStatus)
  }
}

func TestPathExistsStatusCode_DoesNotExist_02(t *testing.T) {
  
  status := PathExistsStatus.DoesNotExist()
  
  statusStr := status.String()
  
  if statusStr != "DoesNotExist" {
    t.Errorf("Error: For 'PathExistsStatus.DoesNotExist()':\n" +
      "Expected PathExistsStatus.String()=='DoesNotExist'.\n" +
      "Instead PathExistsStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathExistsStatusCode_DoesNotExist_03(t *testing.T) {
  
  status := PathExistsStatus.DoesNotExist()
  
  statusValue := status.Value()
  
  if int(statusValue) != 0 {
    t.Errorf("Error: For 'PathExistsStatus.DoesNotExist()':\n" +
      "Expected status.Value()=='0'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathExistsStatusCode_DoesNotExist_04(t *testing.T) {

  statusCode, err := 
    PathExistsStatusCode(0).ParseString("DoesNotExist", true)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"DoesNotExist\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }
  
  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathExistsStatus.DoesNotExist()':\n" +
      "Expected ParseString()=='0'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathExistsStatusCode_DoesNotExist_05(t *testing.T) {

  statusCode, err := 
    PathExistsStatusCode(0).ParseString("doesnotexist", false)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"doesnotexist\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }
  
  if int(statusCode) != 0 {
    t.Errorf("Error: For 'PathExistsStatus.DoesNotExist()':\n" +
      "Expected lower case ParseString()=='0'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathExistsStatusCode_EqualOperator_01(t *testing.T) {

  status1 := PathExistsStatus.Unknown()
  
  status2 := PathExistsStatus.Unknown()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "However, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_EqualOperator_02(t *testing.T) {

  status1 := PathExistsStatus.Unknown()
  
  status2 := PathExistsStatus.DoesNotExist()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_EqualOperator_03(t *testing.T) {

  status1 := PathExistsStatus.DoesNotExist()
  
  status2 := PathExistsStatus.DoesNotExist()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if !result {
    t.Error("Error: For status1=PathExistsStatus.DoesNotExist() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_EqualOperator_04(t *testing.T) {

  status1 := PathExistsStatus.Exists()
  
  status2 := PathExistsStatus.DoesNotExist()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_EqualOperator_05(t *testing.T) {

  status1 := PathExistsStatus.Exists()
  
  status2 := PathExistsStatus.Exists()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_EqualOperator_06(t *testing.T) {

  status1 := PathExistsStatus.Unknown()
  
  status2 := PathExistsStatus.Exists()
  
  result := false
  
  if status1==status2 {
    result = true
  }
  
  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the equal operator (status1==status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_Exists_01(t *testing.T) {
  
  status := PathExistsStatus.Exists()
  
  intStatus := int(status)
  
  if intStatus != 1 {
    t.Errorf("Error: Expected PathExistsStatus.Exists()=='1'.\n" +
      "Instead PathExistsStatus.Exists()=='%v'\n", intStatus)
  }
}

func TestPathExistsStatusCode_Exists_02(t *testing.T) {
  
  status := PathExistsStatus.Exists()
  
  statusStr := status.String()
  
  if statusStr != "Exists" {
    t.Errorf("Error: For 'PathExistsStatus.Exists()':\n" +
      "Expected PathExistsStatus.String()=='Exists'.\n" +
      "Instead PathExistsStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathExistsStatusCode_Exists_03(t *testing.T) {
  
  status := PathExistsStatus.Exists()
  
  statusValue := status.Value()
  
  if int(statusValue) != 1 {
    t.Errorf("Error: For 'PathExistsStatus.Exists()':\n" +
      "Expected status.Value()=='1'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathExistsStatusCode_Exists_04(t *testing.T) {

  statusCode, err := 
    PathExistsStatusCode(0).ParseString("Exists", true)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"Exists\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }
  
  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathExistsStatus.Exists()':\n" +
      "Expected ParseString()=='1'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathExistsStatusCode_Exists_05(t *testing.T) {

  statusCode, err :=
    PathExistsStatusCode(0).ParseString("exists", false)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"exists\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != 1 {
    t.Errorf("Error: For 'PathExistsStatus.Exists()':\n" +
      "Expected lower case ParseString()=='1'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}


func TestPathExistsStatusCode_Exists_06(t *testing.T) {

  _, err :=
    PathExistsStatusCode(0).ParseString("exists", true)

  if err == nil {
    t.Error("Expected an error return from PathExistsStatusCode(0)." +
      "ParseString(\"exists\", true)\n" +
      "because 'exists' is test as 'case sensitive'." +
      "However, NO ERROR WAS RETURNED!!!!\n")
  }

}


func TestPathExistsStatusCode_GreaterThanOperator_01(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOperator_02(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOperator_03(t *testing.T) {

  status1 := PathExistsStatus.DoesNotExist()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.DoesNotExist() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOperator_04(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1 > status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the greater than operator (status1 > status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOperator_05(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1 > status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the greater than operator (status1 > status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOperator_06(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1 > status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the equal operator (status1==status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}


func TestPathExistsStatusCode_GreaterThanOrEqualOperator_01(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "However, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_02(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1>=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1>=status2) to return false.\n" +
      "However, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_03(t *testing.T) {

  status1 := PathExistsStatus.DoesNotExist()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.DoesNotExist() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_04(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_05(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_06(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1>=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the equal operator (status1>=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_GreaterThanOrEqualOperator_07(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1>=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the equal operator (status1>=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_01(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_02(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_03(t *testing.T) {

  status1 := PathExistsStatus.DoesNotExist()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.DoesNotExist() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_04(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1<=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1<=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_05(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_06(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_07(t *testing.T) {

  status1 := PathExistsStatusCode(99)

  status2 := PathExistsStatus.Exists()

  result := false

  if status1<=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatusCode(99) and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1<=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_LessThanOrEqualOperator_08(t *testing.T) {

  status1 := PathExistsStatusCode(-99)

  status2 := PathExistsStatus.Exists()

  result := false

  if status1<=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatusCode(-99) and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1<=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_01(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Unknown()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Unknown()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_02(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_03(t *testing.T) {

  status1 := PathExistsStatus.DoesNotExist()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.DoesNotExist() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_04(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.DoesNotExist()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.DoesNotExist()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_05(t *testing.T) {

  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1!=status2 {
    result = true
  }

  if result {
    t.Error("Error: For status1=PathExistsStatus.Exists() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1!=status2) to return false.\n" +
      "Instead, it returned 'true'!\n")
  }
}

func TestPathExistsStatusCode_NotEqualOperator_06(t *testing.T) {

  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Exists()

  result := false

  if status1!=status2 {
    result = true
  }

  if !result {
    t.Error("Error: For status1=PathExistsStatus.Unknown() and " +
      "status2=PathExistsStatus.Exists()\n" +
      "Expected the not equal operator (status1!=status2) to return true.\n" +
      "Instead, it returned 'false'!\n")
  }
}

func TestPathExistsStatusCode_StatusCodesEqual_01(t *testing.T) {
  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Unknown()

  if status1.StatusCodesEqual(status2) == false {
    t.Error("Error: For status1 == PathExistsStatus.Unknown() and " +
      "status2==PathExistsStatus.Unknown()\n" +
      "Expected status1.StatusCodesEqual(status2)==true.\n" +
      "However, status1.StatusCodesEqual(status2)==false!\n")
  }

}

func TestPathExistsStatusCode_StatusCodesEqual_02(t *testing.T) {
  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Exists()

  if status1.StatusCodesEqual(status2) == true {
    t.Error("Error: For status1 == PathExistsStatus.Unknown() and " +
      "status2==PathExistsStatus.Exists()\n" +
      "Expected status1.StatusCodesEqual(status2)==false.\n" +
      "However, status1.StatusCodesEqual(status2)==true!\n")
  }

}

func TestPathExistsStatusCode_StatusCodesEqual_03(t *testing.T) {
  status1 := PathExistsStatus.DoesNotExist()

  status2 := PathExistsStatus.DoesNotExist()

  if status1.StatusCodesEqual(status2) == false {
    t.Error("Error: For status1 == PathExistsStatus.DoesNotExist() and " +
      "status2==PathExistsStatus.DoesNotExist()\n" +
      "Expected status1.StatusCodesEqual(status2)==true.\n" +
      "However, status1.StatusCodesEqual(status2)==false!\n")
  }

}

func TestPathExistsStatusCode_StatusCodesEqual_04(t *testing.T) {
  status1 := PathExistsStatus.Exists()

  status2 := PathExistsStatus.DoesNotExist()

  if status1.StatusCodesEqual(status2) == true {
    t.Error("Error: For status1 == PathExistsStatus.Exists() and " +
      "status2==PathExistsStatus.DoesNotExist()\n" +
      "Expected status1.StatusCodesEqual(status2)==false.\n" +
      "However, status1.StatusCodesEqual(status2)==true!\n")
  }

}

func TestPathExistsStatusCode_StatusCodesEqual_05(t *testing.T) {
  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.DoesNotExist()

  if status1.StatusCodesEqual(status2) == true {
    t.Error("Error: For status1 == PathExistsStatus.Unknown() and " +
      "status2==PathExistsStatus.DoesNotExist()\n" +
      "Expected status1.StatusCodesEqual(status2)==false.\n" +
      "However, status1.StatusCodesEqual(status2)==true!\n")
  }

}

func TestPathExistsStatusCode_StatusCodesEqual_06(t *testing.T) {
  status1 := PathExistsStatus.Unknown()

  status2 := PathExistsStatus.Exists()

  if status1.StatusCodesEqual(status2) == true {
    t.Error("Error: For status1 == PathExistsStatus.Unknown() and " +
      "status2==PathExistsStatus.Exists()\n" +
      "Expected status1.StatusCodesEqual(status2)==false.\n" +
      "However, status1.StatusCodesEqual(status2)==true!\n")
  }

}

func TestPathExistsStatusCode_StatusIsValid_01(t *testing.T) {

  status := PathExistsStatus.Unknown()

  err := status.StatusIsValid()

  if err != nil {
    t.Errorf("Error returned by status.StatusIsValid()\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestPathExistsStatusCode_StatusIsValid_02(t *testing.T) {

  status := PathExistsStatus.DoesNotExist()

  err := status.StatusIsValid()

  if err != nil {
    t.Errorf("Error returned by status.StatusIsValid()\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestPathExistsStatusCode_StatusIsValid_03(t *testing.T) {

  status := PathExistsStatus.Exists()

  err := status.StatusIsValid()

  if err != nil {
    t.Errorf("Error returned by status.StatusIsValid()\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestPathExistsStatusCode_StatusIsValid_04(t *testing.T) {

  status := PathExistsStatusCode(-2)

  err := status.StatusIsValid()

  if err == nil {
    t.Error("Expected and error return from status.StatusIsValid()\n" +
      "because 'status' == -2 and is INVALID\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestPathExistsStatusCode_StatusIsValid_05(t *testing.T) {

  status := PathExistsStatusCode(2)

  err := status.StatusIsValid()

  if err == nil {
    t.Error("Expected and error return from status.StatusIsValid()\n" +
      "because 'status' == 2 and is INVALID\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestPathExistsStatusCode_StatusIsValid_06(t *testing.T) {

  status := PathExistsStatusCode(-2000)

  err := status.StatusIsValid()

  if err == nil {
    t.Error("Expected and error return from status.StatusIsValid()\n" +
      "because 'status' == -2000 and is INVALID\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestPathExistsStatusCode_StatusIsValid_07(t *testing.T) {

  status := PathExistsStatusCode(9000)

  err := status.StatusIsValid()

  if err == nil {
    t.Error("Expected and error return from status.StatusIsValid()\n" +
      "because 'status' == 9000 and is INVALID\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}


func TestPathExistsStatusCode_Unknown_01(t *testing.T) {

  status := PathExistsStatus.Unknown()

  intStatus := int(status)

  if intStatus != -1 {
    t.Errorf("Error: Expected PathExistsStatus.Unknown()=='-1'.\n" +
      "Instead PathExistsStatus.Exists()=='%v'\n", intStatus)
  }
}

func TestPathExistsStatusCode_Unknown_02(t *testing.T) {

  status := PathExistsStatus.Unknown()

  statusStr := status.String()

  if statusStr != "Unknown" {
    t.Errorf("Error: For 'PathExistsStatus.Unknown()':\n" +
      "Expected PathExistsStatus.String()=='Unknown'.\n" +
      "Instead PathExistsStatus.String()=='%v'\n", statusStr)
  }
}

func TestPathExistsStatusCode_Unknown_03(t *testing.T) {

  status := PathExistsStatus.Unknown()

  statusValue := status.Value()

  if int(statusValue) != -1 {
    t.Errorf("Error: For 'PathExistsStatus.Unknown()':\n" +
      "Expected status.Value()=='-1'.\n" +
      "Instead status.Value()=='%v'\n", int(statusValue))
  }
}

func TestPathExistsStatusCode_Unknown_04(t *testing.T) {

  statusCode, err :=
    PathExistsStatusCode(0).ParseString("Unknown", true)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"Unknown\", true)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != -1 {
    t.Errorf("Error: For 'PathExistsStatus.Unknown()':\n" +
      "Expected ParseString()=='-1'.\n" +
      "Instead ParseString()=='%v'\n", int(statusCode))
  }
}

func TestPathExistsStatusCode_Unknown_05(t *testing.T) {

  statusCode, err :=
    PathExistsStatusCode(0).ParseString("unknown", false)

  if err != nil {
    t.Errorf("Error returned by PathExistsStatusCode(0)." +
      "ParseString(\"unknown\", false)\n" +
      "Error='%v'\n", err.Error())
    return
  }

  if int(statusCode) != -1 {
    t.Errorf("Error: For 'PathExistsStatus.Unknown()':\n" +
      "Expected lower case ParseString()=='-1'.\n" +
      "Instead lower case ParseString()=='%v'\n", int(statusCode))
  }
}
