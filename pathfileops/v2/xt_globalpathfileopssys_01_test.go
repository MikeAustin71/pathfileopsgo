package pathfileops

import (
  "runtime"
  "strings"
  "testing"
)

func TestGlobalPathFileOpsSys_GetOperatingSystem_01(t *testing.T) {

  expectedOperatingSystem := strings.ToLower(runtime.GOOS)

  actualOperatingSystem := GlobalPathFileOpsSys{}.GetOperatingSystem()

  if expectedOperatingSystem != actualOperatingSystem {
    t.Errorf("ERROR: Expected Operating System='%v'\n" +
      "Instead, actual Operating System='%v'\n",
      expectedOperatingSystem, actualOperatingSystem)
  }

}

func TestGlobalPathFileOpsSys_GetArchitecture_01(t *testing.T) {

  expectedArchitecture := strings.ToLower(runtime.GOARCH)

  actualArchitecture := GlobalPathFileOpsSys{}.GetArchitecture()

  if expectedArchitecture != actualArchitecture {
    t.Errorf("ERROR: Expected Operating System='%v'\n" +
      "Instead, actual Operating System='%v'\n",
      expectedArchitecture, actualArchitecture)
  }

}

func TestGlobalPathFileOpsSys_IsWindowsOperatingSystem_01(t *testing.T) {

  expectedOperatingSystem := strings.ToLower(runtime.GOOS)

  expectedIsWindowsOpsSystem := false

  if strings.Contains(expectedOperatingSystem,"windows") {
    expectedIsWindowsOpsSystem = true
  }

  actualIsWindowsOpsSystem := GlobalPathFileOpsSys{}.IsWindowsOperatingSystem()

  if expectedIsWindowsOpsSystem != actualIsWindowsOpsSystem {
    t.Errorf("ERROR: Expected Operating System='%v'\n" +
      "Instead, Operating System='%v'\n",
      expectedIsWindowsOpsSystem, actualIsWindowsOpsSystem)
  }

}

func TestGlobalPathFileOpsSys_IsLinuxOperatingSystem(t *testing.T) {

  expectedOperatingSystem := strings.ToLower(runtime.GOOS)

  expectedIsLinuxOpsSystem := false

  if strings.Contains(expectedOperatingSystem,"linux") {
    expectedIsLinuxOpsSystem = true
  }

  actualIsLinuxOpsSystem := GlobalPathFileOpsSys{}.IsLinuxOperatingSystem()

  if expectedIsLinuxOpsSystem != actualIsLinuxOpsSystem {
    t.Errorf("ERROR: Expected IsLinuxSystem='%v'\n" +
      "Instead, IsLinuxSystem='%v'\n",
      expectedIsLinuxOpsSystem, actualIsLinuxOpsSystem)
  }
}
