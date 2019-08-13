package pathfileops

import (
  "fmt"
  "testing"
)

func TestValidPathStrDto_AbsolutePathDoesExist_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  err := validpathDto.SetAbsPathDoesExistStatus(PathExistsStatus.Exists())

  if err != nil {
    t.Errorf("Error returned by validpathDto." +
      "SetAbsPathDoesExistStatus(PathExistsStatus.Exists())\n" +
      "Error='%v'\n", err.Error())
  }

  absPathDoesExist := validpathDto.AbsolutePathDoesExist()

  if absPathDoesExist != PathExistsStatus.Exists() {

    t.Errorf("Expected absPathDoesExist==PathExistsStatus.Exists().\n" +
      "Instead, absPathDoesExist==%v\n", absPathDoesExist.String())
  }
}

func TestValidPathStrDto_GetPath_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  expectedAbsolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto.SetPath(expectedAbsolutePath)

  actualPath := validpathDto.GetPath()

  if expectedAbsolutePath != actualPath {

    t.Errorf("Expected validpathDto.GetPath()=='%v'.\n" +
      "Instead, validpathDto.GetPath()==%v\n",
      expectedAbsolutePath, actualPath)
  }
}

func TestValidPathStrDto_GetPathStrLen_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  expectedAbsolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto.SetPath(expectedAbsolutePath)

  expectedPathStrLen := len(expectedAbsolutePath)

  actualPathStrLen := validpathDto.GetPathStrLen()

  if expectedPathStrLen != actualPathStrLen {

    t.Errorf("Expected validpathDto.GetPathStrLen()=='%v'.\n" +
      "Instead, validpathDto.GetPathStrLen()==%v\n",
      expectedPathStrLen, actualPathStrLen)
  }
}

func TestValidPathStrDto_GetPathFileInfo(t *testing.T) {

  relPath := "../../filesfortest/checkfiles/testRead918256.txt"

  fMgr, err := FileMgr{}.New(relPath)

  if err != nil {
    t.Errorf("Test setup error returned by FileMgr{}.New(relPath).\n" +
      "relPath='%v'\n" +
      "Error='%v'\n", relPath, err.Error())
    return
  }

  fInfo, err := fMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Test setup error returned by fMgr.GetFileInfoPlus()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto := ValidPathStrDto{}

  validpathDto.SetPathFileInfo(fInfo)

  actualPathFInfoPlus := validpathDto.GetPathFileInfo()

  if !actualPathFInfoPlus.Equal(&fInfo) {
    t.Error("Error: Expected actualPathFInfoPlus.Equal(&fInfo)=='true'.\n" +
      "Instead actualPathFInfoPlus.Equal(&fInfo)=='false'\n")
  }
}

func TestValidPathStrDto_GetAbsPath_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  expectedAbsolutePath,
  err := fh.MakeAbsolutePath("../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto.SetAbsPath(expectedAbsolutePath)

  actualPath := validpathDto.GetAbsPath()

  if expectedAbsolutePath != actualPath {
    t.Errorf("Expected validpathDto.GetAbsPath()=='%v'.\n" +
      "Instead, validpathDto.GetAbsPath()==%v\n",
      expectedAbsolutePath, actualPath)
  }
}

func TestValidPathStrDto_GetAbsPathStrLen_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  expectedAbsolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  expectedAbsPathStrLen := len(expectedAbsolutePath)

  validpathDto.SetAbsPath(expectedAbsolutePath)

  actualPathStrLen := validpathDto.GetAbsPathStrLen()

  if expectedAbsPathStrLen != actualPathStrLen {

    t.Errorf("Expected validpathDto.GetAbsPathStrLen()=='%v'.\n" +
      "Instead, validpathDto.GetAbsPathStrLen()==%v\n",
      expectedAbsPathStrLen, actualPathStrLen)
  }
}


func TestValidPathStrDto_GetAbsPathFileInfo(t *testing.T) {

  relPath := "../../filesfortest/checkfiles/testRead918256.txt"

  fMgr, err := FileMgr{}.New(relPath)

  if err != nil {
    t.Errorf("Test setup error returned by FileMgr{}.New(relPath).\n" +
      "relPath='%v'\n" +
      "Error='%v'\n", relPath, err.Error())
    return
  }

  expectedAbsFInfo, err := fMgr.GetFileInfoPlus()

  if err != nil {
    t.Errorf("Test setup error returned by fMgr.GetFileInfoPlus()\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto := ValidPathStrDto{}

  validpathDto.SetAbsPathFileInfo(expectedAbsFInfo)

  actualAbsPathFInfoPlus := validpathDto.GetAbsPathFileInfo()

  if !actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo) {
    t.Error("Error: Expected actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo)=='true'.\n" +
      "Instead actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo)=='false'\n")
  }

}

func TestValidPathStrDto_GetError(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  expectedErr := fmt.Errorf("Error Test!")

  validpathDto.SetError(expectedErr)

  actualErr := validpathDto.GetError()

  if actualErr == nil {
    t.Errorf("Expected an error return from validpathDto.GetError().\n" +
      "Instead, the return value was 'nil'.\n")
    return
  }

  actualErrStr := actualErr.Error()
  expectedErrStr := expectedErr.Error()

  if actualErrStr != expectedErrStr {
    t.Errorf("Expected actualErrStr='%v'\n" +
      "Instead, actualErrStr='%v'\n",
      actualErrStr, expectedErrStr)
    return
  }
}

func TestValidPathStrDto_GetOriginalPathStr_01(t *testing.T) {

  expectedOriginalPath := "../../../filesfortest/checkfiles/testRead918256.txt"

  validpathDto := ValidPathStrDto{}

  validpathDto.SetOriginalPathStr(expectedOriginalPath)

  actualOriginalPath := validpathDto.GetOriginalPathStr()

  if expectedOriginalPath != actualOriginalPath {
    t.Errorf("Expected validpathDto.GetOriginalPathStr()=='%v'\n" +
      "Instead, validpathDto.GetOriginalPathStr()=='%v'\n",
      expectedOriginalPath, actualOriginalPath)
  }

}

func TestValidPathStrDto_GetPathType(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  expectedPathType := PathFileType.Path()

  validpathDto.SetPathType(expectedPathType)

  actualPathType := validpathDto.GetPathType()

  if expectedPathType != actualPathType {
    t.Errorf("Error: Expected actualPathType='%v'\n" +
      "Instead, actualPathType='%v'\n",
      expectedPathType, actualPathType)
  }

}

func TestValidPathStrDto_GetPathVolumeName_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  absolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
  _,
  expectedVolName := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.SetPathVolumeName(expectedVolName)

  actualVolumeName := validpathDto.GetPathVolumeName()

  if expectedVolName != actualVolumeName {
    t.Errorf("Expected validpathDto.GetPathVolumeName()=='%v'\n" +
      "Instead, validpathDto.GetPathVolumeName()='%v'\n",
      expectedVolName, actualVolumeName)
  }

}

func TestValidPathStrDto_GetPathVolumeIndex_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  absolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  expectedVolumeIndex,
  _,
  _ := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.SetPathVolumeIndex(expectedVolumeIndex)

  actualVolumeIndex := validpathDto.GetPathVolumeIndex()

  if expectedVolumeIndex != actualVolumeIndex {
    t.Errorf("Expected validpathDto.GetPathVolumeIndex()=='%v'\n" +
      "Instead, validpathDto.GetPathVolumeIndex()='%v'\n",
      expectedVolumeIndex, actualVolumeIndex)
  }
}

func TestValidPathStrDto_GetPathVolumeStrLength_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  absolutePath,
  err := fh.MakeAbsolutePath("../../../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../../../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
  expectedVolumeStringLength,
  expectedVolumeStr := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.SetPathVolumeName(expectedVolumeStr)

  actualVolumeStrLength := validpathDto.GetPathVolumeStrLength()

  if expectedVolumeStringLength != actualVolumeStrLength {
    t.Errorf("Expected validpathDto.GetPathVolumeStrLength()=='%v'\n" +
      "Instead, validpathDto.GetPathVolumeStrLength()='%v'\n",
      expectedVolumeStringLength, actualVolumeStrLength)
  }
}

func TestValidPathStrDto_IsInitialized(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  validpathDto.SetIsInitialized(true)

  isInitialized := validpathDto.IsInitialized()

  if isInitialized != true {
    t.Error("Expected isInitialized=='true'\n" +
      "Instead, isInitialized='false'\n")
  }
}

func TestValidPathStrDto_IsDtoValid_01(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  err = validPathDto.IsDtoValid("" )

  if err != nil {
    t.Errorf("Expected no error return from validPathDto.IsDtoValid(\"\")\n" +
      "Instead, an error was returned.\n" +
      "Error='%v'\n", err.Error())
  }

}

func TestValidPathStrDto_IsDtoValid_02(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.isInitialized = false

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because inInitialized='false'.\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_IsDtoValid_03(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.pathIsValid = PathValidStatus.Unknown()

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because validPathDto.pathIsValid = PathValidStatus.Unknown().\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_IsDtoValid_04(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.pathStr = ""

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because validPathDto.pathStr = \"\".\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_IsDtoValid_05(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.absPathStr = ""

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because validPathDto.absPathStr = \"\".\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_IsDtoValid_06(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.pathDoesExist  = PathExistsStatusCode(-99)

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because validPathDto.pathDoesExist is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_IsDtoValid_07(t *testing.T) {

  testPath := "../../../createFilesTest/Level01/Level02"

  dMgr := DirMgr{}

  validPathDto,
  err := dMgr.ParseValidPathStr(testPath)

  if err != nil {
    fmt.Printf("Error returned by dMgr.ParseValidPathStr(testPath)\n"+
      "testPath='%v'\n"+
      "Error='%v'\n",
      testPath, err.Error())
    return
  }

  validPathDto.absPathDoesExist  = PathExistsStatusCode(-99)

  err = validPathDto.IsDtoValid("" )

  if err == nil {
    t.Error("Expected an error return from validPathDto.IsDtoValid(\"\")\n" +
      "because validPathDto.absPathDoesExist is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!'\n")
  }

}

func TestValidPathStrDto_PathDoesExist_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  err := validpathDto.SetPathDoesExist(PathExistsStatus.DoesNotExist())

  if err != nil {
    t.Errorf("Error returned by validpathDto.SetPathDoesExist" +
      "(PathExistsStatus.DoesNotExist())\n" +
      "Error='%v'\n", err.Error())
  }

  pathDoesExist := validpathDto.PathDoesExist()

  if pathDoesExist != PathExistsStatus.DoesNotExist() {

    t.Errorf("Expected pathDoesExist==PathExistsStatus.DoesNotExist().\n" +
      "Instead, pathDoesExist==%v\n", pathDoesExist.String())
  }
}

func TestValidPathStrDto_PathIsValid(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  expectedPathIsValidValue := PathValidStatus.Invalid()

  err := validpathDto.SetPathIsValid(expectedPathIsValidValue)

  if err != nil {
    t.Errorf("Error returned by SetPathIsValid(expectedPathIsValidValue)\n" +
      "expectedPathIsValidValue='%v'\n" +
      "Error='%v'\n", expectedPathIsValidValue.String(), err.Error())
    return
  }

  actualPathIsValid := validpathDto.GetPathIsValid()

  if expectedPathIsValidValue != actualPathIsValid {
    t.Errorf("Expected that actualPathIsValid='%v'\n" +
      "Instead, actualPathIsValid='%v'\n",
      expectedPathIsValidValue, actualPathIsValid)
  }
}

func TestValidPathStrDto_SetAbsPathDoesExistStatus(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  invalidPathExistsStatus := PathExistsStatusCode(-99)

  err := validpathDto.SetAbsPathDoesExistStatus(invalidPathExistsStatus)

  if err == nil {
    t.Error("Expected an error return validpathDto." +
      "SetAbsPathDoesExistStatus(invalidPathExistsStatus)\n" +
      "because the PathExistStatusCode is invalid.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }


}
