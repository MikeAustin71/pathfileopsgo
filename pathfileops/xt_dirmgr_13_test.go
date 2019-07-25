package pathfileops

import (
  "os"
  fp "path/filepath"
  "runtime"
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

  expectedPath := origDir

  if strings.Contains(strings.ToLower(runtime.GOOS), "windows") {
    expectedPath = expectedPath + string(os.PathSeparator)
  }

  expectedAbsDir := expectedPath

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
  expectedPath := fh.AdjustPathSlash(rawDir)
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

func TestDirMgr_New_10(t *testing.T) {

  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/.git")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2/.git")
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

  expectedParentPath := strings.TrimSuffix(expectedAbsDir, fh.AdjustPathSlash("/.git"))
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

func TestDirMgr_New_11(t *testing.T) {

  testDir := "..\\dirmgrtests\\dir01\\\\dir02\\dir03"

  _, err := DirMgr{}.New(testDir)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.New(testDir)\n" +
      "because testDir contains invalid path separators.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_New_12(t *testing.T) {

  testDir := "..\\dirmgrtests\\dir01...dir02\\dir03"

  _, err := DirMgr{}.New(testDir)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.New(testDir)\n" +
      "because testDir contains invalid dot characters.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_New_13(t *testing.T) {

  testDir := ""

  _, err := DirMgr{}.New(testDir)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.New(testDir)\n" +
      "because testDir is an empty string.\n" +
      "However, NO ERROR WAS RETURNED!!!\n")
  }

}

func TestDirMgr_New_14(t *testing.T) {

  testDir := "../dirmgrtests/dir01/dir02/dir03"

  fh := FileHelper{}

  expectedAbsDir,
  err := fh.MakeAbsolutePath(testDir)

  expectedAbsDir = strings.ToLower(expectedAbsDir)

  testDir = "    ../dirmgrtests/dir01/dir02/dir03   "

  testDMgr, err := DirMgr{}.New(testDir)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.New(testDir)\n" +
      "testDir='%v'\n" +
      "Error:'%v'\n", testDir, err.Error())
    return
  }

  if expectedAbsDir != strings.ToLower(testDMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected absolute path='%v'.\n" +
      "Instead, actual absolute path='%v'\n",
      expectedAbsDir, strings.ToLower(testDMgr.GetAbsolutePath()))
  }
}


func TestDirMgr_New_15(t *testing.T) {

  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/program files/testfiles2")
  expectedPath := fh.AdjustPathSlash("../testfiles/program files/testfiles2")

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

func TestDirMgr_NewFromFileInfo_01(t *testing.T) {
  fh := FileHelper{}
  baseDir := fh.AdjustPathSlash("../filesfortest")
  targetDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  expectedAbsTargetDir,
  err := fh.MakeAbsolutePath(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return

  }

  expectedAbsTargetDir = strings.ToLower(expectedAbsTargetDir)

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)\n" +
      "baseDir='%v'\ntargetFileInfo='%v'\nError='%v'\n",
      baseDir, targetFileInfo.Name(), err.Error())
    return
  }

  if expectedAbsTargetDir != strings.ToLower(targetDMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected absolute path='%v'.\n" +
      "Instead, actual absolute path='%v'\n",
      expectedAbsTargetDir, strings.ToLower(targetDMgr.GetAbsolutePath()))
  }
}

func TestDirMgr_NewFromFileInfo_02(t *testing.T) {
  fh := FileHelper{}

  baseDir := fh.AdjustPathSlash("../testdestdir")

  var targetFileInfo os.FileInfo

  _, err := DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)\n" +
      "because targetFileInfo is empty and uninitialized.\n" +
      "However, NO ERROR WAS RETURNED!!!")
  }

  return
}

func TestDirMgr_NewFromFileInfo_03(t *testing.T) {
  fh := FileHelper{}

  targetDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  baseDir := ""

  _, err = DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)

  if err == nil {

    t.Error("Expected an error return by DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)\n" +
      "because baseDir is an empty string.\nHowever, NO ERROR WAS RETURNED!!!\n")

  }

}

func TestDirMgr_NewFromFileInfo_04(t *testing.T) {
  fh := FileHelper{}
  baseDir := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02")
  finalDir := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/htmlFilesForTest")
  targetDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  expectedAbsFinalDir,
  err := fh.MakeAbsolutePath(finalDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(finalDir)\n" +
      "finalDir='%v'\nError:='%v'\n",
      finalDir, err.Error())
    return

  }

  expectedAbsFinalDir = strings.ToLower(expectedAbsFinalDir)

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  targetDMgr, err := DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromFileInfo(baseDir, targetFileInfo)\n" +
      "baseDir='%v'\ntargetFileInfo='%v'\nError='%v'\n",
      baseDir, targetFileInfo.Name(), err.Error())
    return
  }

  if expectedAbsFinalDir != strings.ToLower(targetDMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected absolute path='%v'.\n" +
      "Instead, actual absolute path='%v'\n",
      expectedAbsFinalDir, strings.ToLower(targetDMgr.GetAbsolutePath()))
  }
}

func TestDirMgr_NewFromFileInfo_05(t *testing.T) {

  fh := FileHelper{}
  baseDir := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02")
  finalDir := fh.AdjustPathSlash("../checkfiles/checkfiles03/checkfiles03_02/htmlFilesForTest")
  targetDir := fh.AdjustPathSlash("../filesfortest/htmlFilesForTest")

  expectedAbsFinalDir,
  err := fh.MakeAbsolutePath(finalDir)

  if err != nil {
    t.Errorf("Error returned by fh.MakeAbsolutePath(finalDir)\n" +
      "finalDir='%v'\nError:='%v'\n",
      finalDir, err.Error())
    return

  }

  expectedAbsFinalDir = strings.ToLower(expectedAbsFinalDir)

  targetFileInfo, err := fh.GetFileInfo(targetDir)

  if err != nil {
    t.Errorf("Error returned by fh.GetFileInfo(targetDir)\n" +
      "targetDir='%v'\nError:='%v'\n",
      targetDir, err.Error())
    return
  }

  targetFileInfoPlus :=
    FileInfoPlus{}.NewFromPathFileInfo(targetDir, targetFileInfo)


  targetDMgr, err := DirMgr{}.NewFromFileInfo(baseDir, targetFileInfoPlus)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromFileInfo(baseDir, targetFileInfoPlus)\n" +
      "baseDir='%v'\ntargetFileInfo='%v'\nError='%v'\n",
      baseDir, targetFileInfo.Name(), err.Error())
    return
  }

  if expectedAbsFinalDir != strings.ToLower(targetDMgr.GetAbsolutePath()) {
    t.Errorf("ERROR: Expected absolute path='%v'.\n" +
      "Instead, actual absolute path='%v'\n",
      expectedAbsFinalDir, strings.ToLower(targetDMgr.GetAbsolutePath()))
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_01(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "TestDirMgr_NewFromKnownPathDirectoryName_01"

  expectedDir := "../checkfiles/checkfiles02/TestDirMgr_NewFromKnownPathDirectoryName_01"

  fh := FileHelper{}

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(expectedDir)\n" +
      "expectedDir='%v'\nError='%v'\n",
      expectedDir, err.Error())
    return
  }

  testDirMgr, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)\n" +
      "parentDirectory='%v'\n subDirectoryName='%v'\nError='%v'\n",
      parentDirectory, subDirectoryName, err.Error())
    return
  }

  if expectedAbsDir != testDirMgr.absolutePath {
    t.Errorf("Expected testDirMgr.absolutePath='%v'\n" +
      "Instead, testDirMgr.absolutePath='%v'\n",
      expectedAbsDir, testDirMgr.absolutePath)
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_02(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "./TestDirMgr_NewFromKnownPathDirectoryName_02"

  expectedDir := "../checkfiles/checkfiles02/TestDirMgr_NewFromKnownPathDirectoryName_02"

  fh := FileHelper{}

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(expectedDir)\n" +
      "expectedDir='%v'\nError='%v'\n",
      expectedDir, err.Error())
    return
  }

  testDirMgr, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)\n" +
      "parentDirectory='%v'\n subDirectoryName='%v'\nError='%v'\n",
      parentDirectory, subDirectoryName, err.Error())
    return
  }

  if expectedAbsDir != testDirMgr.absolutePath {
    t.Errorf("Expected testDirMgr.absolutePath='%v'\n" +
      "Instead, testDirMgr.absolutePath='%v'\n",
      expectedAbsDir, testDirMgr.absolutePath)
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_03(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "../TestDirMgr_NewFromKnownPathDirectoryName_03"

  expectedDir := "../checkfiles/checkfiles02/TestDirMgr_NewFromKnownPathDirectoryName_03"

  fh := FileHelper{}

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(expectedDir)\n" +
      "expectedDir='%v'\nError='%v'\n",
      expectedDir, err.Error())
    return
  }

  testDirMgr, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)\n" +
      "parentDirectory='%v'\n subDirectoryName='%v'\nError='%v'\n",
      parentDirectory, subDirectoryName, err.Error())
    return
  }

  if expectedAbsDir != testDirMgr.absolutePath {
    t.Errorf("Expected testDirMgr.absolutePath='%v'\n" +
      "Instead, testDirMgr.absolutePath='%v'\n",
      expectedAbsDir, testDirMgr.absolutePath)
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_04(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := ".TestDirMgr_NewFromKnownPathDirectoryName_04"

  expectedDir := "../checkfiles/checkfiles02/TestDirMgr_NewFromKnownPathDirectoryName_04"

  fh := FileHelper{}

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(expectedDir)\n" +
      "expectedDir='%v'\nError='%v'\n",
      expectedDir, err.Error())
    return
  }

  testDirMgr, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)\n" +
      "parentDirectory='%v'\n subDirectoryName='%v'\nError='%v'\n",
      parentDirectory, subDirectoryName, err.Error())
    return
  }

  if expectedAbsDir != testDirMgr.absolutePath {
    t.Errorf("Expected testDirMgr.absolutePath='%v'\n" +
      "Instead, testDirMgr.absolutePath='%v'\n",
      expectedAbsDir, testDirMgr.absolutePath)
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_05(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "..TestDirMgr_NewFromKnownPathDirectoryName_05"

  expectedDir := "../checkfiles/checkfiles02/TestDirMgr_NewFromKnownPathDirectoryName_05"

  fh := FileHelper{}

  expectedAbsDir, err := fh.MakeAbsolutePath(expectedDir)

  if err != nil {
    t.Errorf("Test Setup Error returned by fh.MakeAbsolutePath(expectedDir)\n" +
      "expectedDir='%v'\nError='%v'\n",
      expectedDir, err.Error())
    return
  }

  testDirMgr, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)\n" +
      "parentDirectory='%v'\n subDirectoryName='%v'\nError='%v'\n",
      parentDirectory, subDirectoryName, err.Error())
    return
  }

  if expectedAbsDir != testDirMgr.absolutePath {
    t.Errorf("Expected testDirMgr.absolutePath='%v'\n" +
      "Instead, testDirMgr.absolutePath='%v'\n",
      expectedAbsDir, testDirMgr.absolutePath)
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_06(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := ""

  _, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.NewFromKnownPathDirectoryName(" +
      "parentDirectory, subDirectoryName)\n" +
      "because subDirectoryName is an Empty string.\n" +
      "However, NO ERROR WAS RETURNED!\n")
    return
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_07(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "...TestDirMgr_NewFromKnownPathDirectoryName_07"

  _, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.NewFromKnownPathDirectoryName(" +
      "parentDirectory, subDirectoryName)\n" +
      "because subDirectoryName contains '...'\n" +
      "However, NO ERROR WAS RETURNED!\n")
    return
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_08(t *testing.T) {

  parentDirectory := "../checkfiles/checkfiles02"

  subDirectoryName := "TestDirMgr_NewFromKnownPathDirectoryName_08\\\\\\"

  _, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.NewFromKnownPathDirectoryName(" +
      "parentDirectory, subDirectoryName)\n" +
      "because subDirectoryName contains '\\\\'\n" +
      "However, NO ERROR WAS RETURNED!\n")
    return
  }
}

func TestDirMgr_NewFromKnownPathDirectoryName_09(t *testing.T) {

  parentDirectory := "..\\checkfiles\\\\checkfiles02"

  subDirectoryName := "TestDirMgr_NewFromKnownPathDirectoryName_09"

  _, err :=
    DirMgr{}.NewFromKnownPathDirectoryName(parentDirectory, subDirectoryName)

  if err == nil {
    t.Error("Expected an error return from DirMgr{}.NewFromKnownPathDirectoryName(" +
      "parentDirectory, subDirectoryName)\n" +
      "because subDirectoryName contains '\\\\'\n" +
      "However, NO ERROR WAS RETURNED!\n")
    return
  }
}
