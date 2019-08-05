package pathfileops

import "testing"

func TestValidPathStrDto_AbsolutePathDoesExist_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}
  validpathDto.absPathDoesExist = PathExistsStatus.Exists()

  absPathDoesExist := validpathDto.AbsolutePathDoesExist()

  if absPathDoesExist != PathExistsStatus.Exists() {

    t.Errorf("Expected absPathDoesExist==PathExistsStatus.Exists().\n" +
      "Instead, absPathDoesExist==%v\n", absPathDoesExist.String())
  }

}

func TestValidPathStrDto_PathDoesExist_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}
  validpathDto.pathDoesExist = PathExistsStatus.DoesNotExist()

  pathDoesExist := validpathDto.PathDoesExist()

  if pathDoesExist != PathExistsStatus.DoesNotExist() {

    t.Errorf("Expected pathDoesExist==PathExistsStatus.DoesNotExist().\n" +
      "Instead, pathDoesExist==%v\n", pathDoesExist.String())
  }

}

func TestValidPathStrDto_GetPath_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  expectedAbsolutePath,
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto.pathStr = expectedAbsolutePath

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
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  expectedPathStrLen := len(expectedAbsolutePath)

  validpathDto.pathStrLength = expectedPathStrLen

  actualPathStrLen := validpathDto.GetPathStrLen()

  if expectedPathStrLen != actualPathStrLen {

    t.Errorf("Expected validpathDto.GetPathStrLen()=='%v'.\n" +
      "Instead, validpathDto.GetPathStrLen()==%v\n",
      expectedPathStrLen, actualPathStrLen)
  }
}

func TestValidPathStrDto_GetPathFileInfo(t *testing.T) {

  relPath := "../filesfortest/checkfiles/testRead918256.txt"

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

  validpathDto.pathFInfoPlus = fInfo.CopyOut()

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
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  validpathDto.absPathStr = expectedAbsolutePath

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
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  expectedAbsPathStrLen := len(expectedAbsolutePath)

  validpathDto.absPathStrLength = expectedAbsPathStrLen

  actualPathStrLen := validpathDto.GetAbsPathStrLen()

  if expectedAbsPathStrLen != actualPathStrLen {

    t.Errorf("Expected validpathDto.GetAbsPathStrLen()=='%v'.\n" +
      "Instead, validpathDto.GetAbsPathStrLen()==%v\n",
      expectedAbsPathStrLen, actualPathStrLen)
  }
}


func TestValidPathStrDto_GetAbsPathFileInfo(t *testing.T) {

  relPath := "../filesfortest/checkfiles/testRead918256.txt"

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

  validpathDto.absPathFInfoPlus = expectedAbsFInfo.CopyOut()

  actualAbsPathFInfoPlus := validpathDto.GetAbsPathFileInfo()

  if !actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo) {
    t.Error("Error: Expected actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo)=='true'.\n" +
      "Instead actualAbsPathFInfoPlus.Equal(&expectedAbsFInfo)=='false'\n")
  }

}

func TestValidPathStrDto_GetOriginalPathStr_01(t *testing.T) {

  expectedOriginalPath := "../filesfortest/checkfiles/testRead918256.txt"

  validpathDto := ValidPathStrDto{}

  validpathDto.originalPathStr = expectedOriginalPath

  actualOriginalPath := validpathDto.GetOriginalPathStr()

  if expectedOriginalPath != actualOriginalPath {
    t.Errorf("Expected validpathDto.GetOriginalPathStr()=='%v'\n" +
      "Instead, validpathDto.GetOriginalPathStr()=='%v'\n",
      expectedOriginalPath, actualOriginalPath)
  }

}

func TestValidPathStrDto_GetPathVolumeName_01(t *testing.T) {

  validpathDto := ValidPathStrDto{}

  fh := FileHelper{}

  absolutePath,
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
  _,
  expectedVolName := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.pathVolumeName = expectedVolName

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
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  expectedVolumeIndex,
  _,
  _ := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.pathVolumeIndex = expectedVolumeIndex

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
  err := fh.MakeAbsolutePath("../createFilesTest/Level01/Level02")

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(\"../createFilesTest/Level01/Level02\")\n" +
      "Error='%v'\n", err.Error())
    return
  }

  _,
  expectedVolumeStringLength,
  _ := fh.GetVolumeNameIndex(absolutePath)

  validpathDto.pathVolumeStrLength = expectedVolumeStringLength

  actualVolumeStrLength := validpathDto.GetPathVolumeStrLength()

  if expectedVolumeStringLength != actualVolumeStrLength {
    t.Errorf("Expected validpathDto.GetPathVolumeStrLength()=='%v'\n" +
      "Instead, validpathDto.GetPathVolumeStrLength()='%v'\n",
      expectedVolumeStringLength, actualVolumeStrLength)
  }

}
