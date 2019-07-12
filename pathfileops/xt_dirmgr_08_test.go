package pathfileops

import (
  fp "path/filepath"
  "strings"
  "testing"
)

func TestDirMgr_New_01(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")

  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir).\n"+
      "origDir=='%v'  Error='%v'", origDir, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'. Instead, DirMgr.isFInfoInitialized=='%v'", true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'.\n"+
      "Instead, parentPath=='%v'.", expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

}

func TestDirMgr_New_02(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("./")
  expectedPath := fh.AdjustPathSlash("./")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir).\n"+
      "origDir=='%v'  Error='%v'", origDir, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v' Error='%v'", origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.",
      true, dMgr.isParentPathPopulated)
  }

}

func TestDirMgr_New_03(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../")
  expectedPath := fh.AdjustPathSlash("../")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir).\n"+
      "origDir=='%v' Error='%v'", origDir, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'",
      origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      true, dMgr.isParentPathPopulated)
  }
}

func TestDirMgr_New_04(t *testing.T) {
  fh := FileHelper{}
  xDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir().\n"+
      "Error='%v'", err.Error())
    return
  }

  volName := fp.VolumeName(xDir)

  origDir := fh.AdjustPathSlash(volName)
  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir := origDir

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v' Error='%v'", origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.\n",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.\n"+
      "Parent Path='%v'\n",
      false, dMgr.isParentPathPopulated, dMgr.parentPath)
  }
}

func TestDirMgr_New_05(t *testing.T) {
  fh := FileHelper{}
  origDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir().\n"+
      "Error='%v'\n", err.Error())
    return
  }

  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(origDir).\n"+
      "origDir=='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(origDir)

  expectedAbsPathDoesExist := fh.DoesFileExist(origDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n", origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'\n",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.\n", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.\n",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.\n",
      true, dMgr.isParentPathPopulated)
  }
}

func TestDirMgr_New_06(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/test2007.txt")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.MakeAbsolutePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n", expectedPath, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v' Error='%v'\n",
      origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'\n",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.\n",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.\n",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'.\n"+
      "Instead, parentPath=='%v'.\n",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.\n",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }
}

func TestDirMgr_New_07(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.MakeAbsolutePath(origDir).\n"+
      "origDir=='%v'\nError='%v'\n",
      origDir, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath :=
    strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/testfiles2"))

  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(origDir).\n"+
      "origDir=='%v'\nError='%v'\n",
      origDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'\n",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'.\n"+
      "Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }
}

func TestDirMgr_New_08(t *testing.T) {
  fh := FileHelper{}
  rawDir := "D:/"
  expectedPath := fh.AdjustPathSlash("D:")
  expectedAbsDir := expectedPath

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)

  expectedVolumeName := fp.VolumeName(expectedAbsDir)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath := ""

  dMgr, err := DirMgr{}.New(rawDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawDir).\n"+
      "rawDir=='%v'\nError='%v'\n",
      rawDir, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'\n",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.\n",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.\n",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'.\n"+
      "Instead, parentPath=='%v'.\n",
      expectedParentPath, dMgr.parentPath)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.\n"+
      "Parent Path='%v'\n",
      false, dMgr.isParentPathPopulated, dMgr.parentPath)
  }
}

func TestDirMgr_New_09(t *testing.T) {
  fh := FileHelper{}
  rawPath := "../filesfortest/newfilesfortest"
  expectedPath := fh.AdjustPathSlash(rawPath)
  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath=='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedPathDoesExist := fh.DoesFileExist(expectedPath)

  expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsPath)

  expectedVolumeName := fp.VolumeName(expectedAbsPath)
  var expectedVolumeIsPopulated bool

  if expectedVolumeName != "" {
    expectedVolumeIsPopulated = true
  } else {
    expectedVolumeIsPopulated = false
  }

  expectedParentPath :=
    strings.TrimSuffix(expectedAbsPath, fh.AdjustPathSlash("/newfilesfortest"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  dMgr, err := DirMgr{}.New(rawPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.New(rawPath).\n"+
      "rawPath=='%v'\nError='%v'\n",
      rawPath, err.Error())
    return
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.isFInfoInitialized=='%v'.\n"+
      "Instead, DirMgr.isFInfoInitialized=='%v'\n",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'.\n"+
      "Instead, DirMgr.isPathPopulated=='%v'\n",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'.\n"+
      "Instead, path=='%v'.\n",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'.\n"+
      "Instead, doesPathExist=='%v'.\n",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'.\n"+
      "Instead, isAbsolutePathPopulated=='%v'.\n",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsPath != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'.\n"+
      "Instead, absolutePath=='%v'\n",
      expectedAbsPath, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'.\n"+
      "Instead, doesAbsolutePathExist=='%v'.\n",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'.\n"+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.\n",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'.\n"+
      "Instead, volumeName=='%v'.\n",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'.\n"+
      "Instead, isVolumePopulated=='%v'.\n",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'.\n"+
      "Instead, parentPath=='%v'.\n",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'.\n"+
      "Instead, isParentPathPopulated=='%v'.\n",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }
}

func TestDirMgr_SetDirMgr_01(t *testing.T) {

  firstDir := "../checkfiles"

  dMgr, err := DirMgr{}.New(firstDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(firstDir).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDir := ""

  _, err = dMgr.SetDirMgr(testDir)

  if err == nil {
    t.Error("Expected an error return from dMgr.SetDirMgr(testDir) because\n" +
      "'testDir' is an empty string.\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

}

func TestDirMgr_SetDirMgr_02(t *testing.T) {

  firstDir := "../checkfiles"

  dMgr, err := DirMgr{}.New(firstDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(firstDir).\n"+
      "Error='%v'\n", err.Error())
    return
  }

  testDir := "      "

  _, err = dMgr.SetDirMgr(testDir)

  if err == nil {
    t.Error("Expected an error return from dMgr.SetDirMgr(testDir) because\n" +
      "'testDir' consists entirely of blank spaces.\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }
}

func TestDirMgr_SetPermissions_01(t *testing.T) {

  testDir := "../dirmgrtests/TestDirMgr_SetPermissions_01"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  fh := FileHelper{}

  permissionsCfg2, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Test Setup Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permission2Txt, err := permissionsCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Setup Error returned by permissionsCfg2.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  testDMgr.isInitialized = false

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because testDMgr is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  testDMgr.isInitialized = true

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), permissionsCfg2)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), permission2Txt, err.Error())
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_02(t *testing.T) {

  testDir := "../dirmgrtests/TestDirMgr_SetPermissions_02"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  fh := FileHelper{}

  permissionsCfg2, err := FilePermissionConfig{}.New("drwxrwxrwx")

  if err != nil {
    t.Errorf("Test Setup Error returned by FilePermissionConfig{}.New(\"drwxrwxrwx\")\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permission2Txt, err := permissionsCfg2.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Test Setup Error returned by permissionsCfg2.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(testDir)
    return
  }

  permissionsCfg.isInitialized = false

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because permissionsCfg is INVALID!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), permissionsCfg2)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), permission2Txt, err.Error())
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_03(t *testing.T) {

  testDir := "../dirmgrtests/TestDirMgr_SetPermissions_01"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  permissionsCfg, err := FilePermissionConfig{}.New("dr--r--r--")

  err = testDMgr.SetPermissions(permissionsCfg)

  if err == nil {
    t.Error("Expected an error returned by testDMgr.SetPermissions(permissionsCfg)\n" +
      "because testDMgr directory DOES NOT EXIST!\nHowever, NO ERROR WAS RETURNED!!!\n")
  }

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }

}

func TestDirMgr_SetPermissions_04(t *testing.T) {

  testDir := "../dirmgrtests/TestDirMgr_SetPermissions_04"

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by DirMgr{}.New(testDir).\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  fh := FileHelper{}

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
    return
  }

  err = testDMgr.MakeDir()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.MakeDir().\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  originalPermCfg, err := testDMgr.GetDirPermissionCodes()

  if err != nil {
    t.Errorf("Test Setup Error returned by testDMgr.GetDirPermissionCodes()\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  originalPermCfgStr, err := originalPermCfg.GetPermissionTextCode()

  if err != nil {

    t.Errorf("Test Setup Error returned by originalPermCfg.GetPermissionTextCode()\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  newPermissionsCfgStr := "dr--r--r--"
  newPermissionsCfg, err := FilePermissionConfig{}.New(newPermissionsCfgStr)

  err = testDMgr.SetPermissions(newPermissionsCfg)

  if err != nil {
    t.Errorf("Error returned by testDMgr.SetPermissions(newPermissionsCfg).\n"+
      "testDMgr='%v'\nnewPermissionsCfg='%v'\n",
      testDMgr.GetAbsolutePath(), newPermissionsCfgStr)
  }

  actualPermCfg, err := testDMgr.GetDirPermissionCodes()

  if err != nil {
    t.Errorf("Error returning actual permission configuration by "+
      "testDMgr.GetDirPermissionCodes()\n"+
      "testDMgr='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  actualPermCfgStr, err := actualPermCfg.GetPermissionTextCode()

  if err != nil {
    t.Errorf("Error returned by actualPermCfg.GetPermissionTextCode().\n"+
      "Error='%v'\n", err.Error())

    _ = fh.DeleteDirPathAll(testDir)

    return
  }

  err = fh.ChangeFileMode(testDMgr.GetAbsolutePath(), originalPermCfg)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.ChangeFileMode(testDMgr."+
      "GetAbsolutePath(), permissionsCfg2).\n"+
      "testDMgr='%v'\npermissionsCfg2='%v'\nError='%v'\n",
      testDMgr.GetAbsolutePath(), originalPermCfgStr, err.Error())
  }

  if actualPermCfgStr == originalPermCfgStr {
    t.Errorf("ERROR: Actual Permission Codes equals Original Permission Codes\n"+
      "Actual Permission Codes='%v'\nOriginal Permission Codes='%v'\n",
      actualPermCfgStr, originalPermCfgStr)
  }

  actualPermCfgRunes := []rune(actualPermCfgStr)
  cntOfRs := 0
  for i := 0; i < len(actualPermCfgRunes); i++ {
    if actualPermCfgRunes[i] == 'r' {
      cntOfRs++
    }
  }

  if cntOfRs != 3 {
    t.Errorf("Expected the Actual Permissions Codes to contain 3-r's or read-only codes.\n"+
      "It did NOT! Therefore the operation to change Permissions Codes FAILED!\n"+
      "Expected Permission Codes='%v'\nActual Permission Codes='%v'\n",
      newPermissionsCfgStr, actualPermCfgStr)
  }

  err = fh.DeleteDirPathAll(testDir)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(testDir)\n"+
      "testDir='%v'\nError='%v'\n", testDir, err.Error())
  }
}

func TestDirMgr_SubstituteBaseDir_01(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01")

  substitutePath := fh.AdjustPathSlash("../checkfiles")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).\nError='%v'\n",
      err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.\n"+
      "Instead substituted path = '%v'\n",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.\n"+
      "Instead substituted absolute path = '%v'\n",
      expectedAbsPath, dMgrResult.absolutePath)
  }

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}

func TestDirMgr_SubstituteBaseDir_02(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return

  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).\nError='%v'\n", err.Error())
    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.\n"+
      "Instead substituted path = '%v'\n",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.\n"+
      "Instead substituted absolute path = '%v'\n",
      expectedAbsPath, dMgrResult.absolutePath)
  }

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return

}

func TestDirMgr_SubstituteBaseDir_03(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\n"+
      "Error='%v'", rawBasePath, err.Error())
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())
  }

  dMgrOrig.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrOrig is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrOrig.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}

func TestDirMgr_SubstituteBaseDir_04(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n", rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrBase is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrBase.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

}

func TestDirMgr_SubstituteBaseDir_05(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedBasePath := fh.AdjustPathSlash("../checkfiles/dir02")

  err := fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
    return
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Startup Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
    return
  }

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath).\n"+
      "expectedPath='%v'\nError='%v'\n",
      expectedPath, err.Error())
    return
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawOrigPath).\n"+
      "rawOrigPath='%v'\nError='%v'\n",
      rawOrigPath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(rawBasePath).\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(substitutePath).\n"+
      "substitutePath='%v'\nError='%v'\n",
      substitutePath, err.Error())

    _ = fh.DeleteDirPathAll(rawBasePath)
    _ = fh.DeleteDirPathAll(expectedBasePath)
    return
  }

  dMgrSubstitute.isInitialized = false

  _, err = dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err == nil {
    t.Error("Expected an error return from dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)\n" +
      "because dMgrSubstitute is INVALID!\nHowever, NO ERROR WAS RETURNED!!!!\n")
  }

  dMgrSubstitute.isInitialized = true

  err = fh.DeleteDirPathAll(rawBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(rawBasePath)\n"+
      "rawBasePath='%v'\nError='%v'\n",
      rawBasePath, err.Error())
  }

  err = fh.DeleteDirPathAll(expectedBasePath)

  if err != nil {
    t.Errorf("Test Clean-Up Error returned by fh.DeleteDirPathAll(expectedBasePath)\n"+
      "expectedBasePath='%v'\nError='%v'\n",
      expectedBasePath, err.Error())
  }

  return
}
