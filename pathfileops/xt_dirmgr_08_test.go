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
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
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

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.", true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.", true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.", expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.", expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.", true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_02(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("./")
  expectedPath := fh.AdjustPathSlash(".")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
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
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'", true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.", true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.", true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.", true, dMgr.isParentPathPopulated)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.", true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_03(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../")
  expectedPath := fh.AdjustPathSlash("..")
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
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
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'",
      origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
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

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_04(t *testing.T) {
  fh := FileHelper{}
  xDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
  }

  volName := fp.VolumeName(xDir)

  origDir := fh.AdjustPathSlash(volName)
  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir := origDir

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
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
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      false, dMgr.isParentPathPopulated)
  }

  if false != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      false, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_05(t *testing.T) {
  fh := FileHelper{}
  origDir, err := fh.GetCurrentDir()

  if err != nil {
    t.Errorf("Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
  }

  expectedPath := fh.AdjustPathSlash(origDir)
  expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
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
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'", expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.", expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.", expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.", expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if true != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      true, dMgr.isParentPathPopulated)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_06(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/test2007.txt")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

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

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.", expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_07(t *testing.T) {
  fh := FileHelper{}
  origDir := fh.AdjustPathSlash("../testfiles/testfiles2/")
  expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
  expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
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

  expectedRelativePath := "testfiles2"

  dMgr, err := DirMgr{}.New(origDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. Instead, "+
      "isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
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

  expectedRelativePath := ""

  dMgr, err := DirMgr{}.New(rawDir)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawDir). rawDir=='%v' Error='%v'", rawDir, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.",
      expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsDir != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsDir, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if false != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. "+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      false, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if false != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      false, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if false != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      false, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_New_09(t *testing.T) {
  fh := FileHelper{}
  rawPath := "../filesfortest/newfilesfortest"
  expectedPath := fh.AdjustPathSlash(rawPath)
  expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expectedPath). expectedPath=='%v'  Error='%v'", expectedPath, err.Error())
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

  expectedParentPath := strings.TrimSuffix(expectedAbsPath, fh.AdjustPathSlash("/newfilesfortest"))
  expectedIsParentPathPopulated := false

  if expectedParentPath != "" {
    expectedIsParentPathPopulated = true
  }

  expectedRelativePath := "newfilesfortest"

  dMgr, err := DirMgr{}.New(rawPath)

  if err != nil {
    t.Errorf("Error returned from DirMgr{}.NewFromPathFileNameExtStr(rawPath). rawPath=='%v' Error='%v'",
      rawPath, err.Error())
  }

  if true != dMgr.isInitialized {
    t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
      true, dMgr.isInitialized)
  }

  if true != dMgr.isPathPopulated {
    t.Errorf("Expected DirMgr.isPathPopulated=='%v'. Instead, DirMgr.isPathPopulated=='%v'",
      true, dMgr.isPathPopulated)
  }

  if expectedPath != dMgr.path {
    t.Errorf("Expected path=='%v'. Instead, path=='%v'.", expectedPath, dMgr.path)
  }

  if expectedPathDoesExist != dMgr.doesPathExist {
    t.Errorf("Expected doesPathExist=='%v'. Instead, doesPathExist=='%v'.",
      expectedPathDoesExist, dMgr.doesPathExist)
  }

  if true != dMgr.isAbsolutePathPopulated {
    t.Errorf("Expected isAbsolutePathPopulated=='%v'. Instead, isAbsolutePathPopulated=='%v'.",
      true, dMgr.isAbsolutePathPopulated)
  }

  if expectedAbsPath != dMgr.absolutePath {
    t.Errorf("Expected absolutePath=='%v'. Instead, absolutePath=='%v'",
      expectedAbsPath, dMgr.absolutePath)
  }

  if expectedAbsPathDoesExist != dMgr.doesAbsolutePathExist {
    t.Errorf("Expected doesAbsolutePathExist=='%v'. Instead, doesAbsolutePathExist=='%v'.",
      expectedAbsPathDoesExist, dMgr.doesAbsolutePathExist)
  }

  if true != dMgr.isAbsolutePathDifferentFromPath {
    t.Errorf("Expected isAbsolutePathDifferentFromPath=='%v'. "+
      "Instead, isAbsolutePathDifferentFromPath=='%v'.",
      true, dMgr.isAbsolutePathDifferentFromPath)
  }

  if expectedVolumeName != dMgr.volumeName {
    t.Errorf("Expected volumeName=='%v'. Instead, volumeName=='%v'.",
      expectedVolumeName, dMgr.volumeName)
  }

  if expectedVolumeIsPopulated != dMgr.isVolumePopulated {
    t.Errorf("Expected isVolumePopulated=='%v'. Instead, isVolumePopulated=='%v'.",
      expectedVolumeIsPopulated, dMgr.isVolumePopulated)
  }

  if expectedParentPath != dMgr.parentPath {
    t.Errorf("Expected parentPath=='%v'. Instead, parentPath=='%v'.",
      expectedParentPath, dMgr.parentPath)
  }

  if expectedIsParentPathPopulated != dMgr.isParentPathPopulated {
    t.Errorf("Expected isParentPathPopulated=='%v'. Instead, isParentPathPopulated=='%v'.",
      expectedIsParentPathPopulated, dMgr.isParentPathPopulated)
  }

  if expectedRelativePath != dMgr.relativePath {
    t.Errorf("Expected relativePath=='%v'. Instead, relativePath=='%v'.",
      expectedRelativePath, dMgr.relativePath)
  }

  if true != dMgr.isRelativePathPopulated {
    t.Errorf("Expected isRelativePathPopulated=='%v'. Instead, isRelativePathPopulated=='%v'.",
      true, dMgr.isRelativePathPopulated)
  }

}

func TestDirMgr_SubstituteBaseDir_01(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01")

  substitutePath := fh.AdjustPathSlash("../checkfiles")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
      "dMgrSubstitute).  Error='%v'",
      err.Error())
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.  Instead substituted "+
      "path = '%v' ",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.  Instead "+
      "substituted absolute path = '%v' ",
      expectedAbsPath, dMgrResult.absolutePath)
  }

}

func TestDirMgr_SubstituteBaseDir_02(t *testing.T) {

  fh := FileHelper{}

  rawOrigPath := fh.AdjustPathSlash("../dirwalktests/dir01/dir02/dir03/")

  rawBasePath := fh.AdjustPathSlash("../dirwalktests/dir01/")

  substitutePath := fh.AdjustPathSlash("../checkfiles/")

  expectedPath := fh.AdjustPathSlash("../checkfiles/dir02/dir03")

  expectedAbsPath, err := fh.GetAbsPathFromFilePath(expectedPath)

  if err != nil {
    t.Errorf("Error returned by fh.GetAbsPathFromFilePath(expectedPath). expectedPath='%v'  Error='%v'", expectedPath, err.Error())
  }

  expectedAbsPath = strings.ToLower(expectedAbsPath)

  dMgrOrig, err := DirMgr{}.New(rawOrigPath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
  }

  dMgrBase, err := DirMgr{}.New(rawBasePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
  }

  dMgrSubstitute, err := DirMgr{}.New(substitutePath)

  if err != nil {
    t.Errorf("Error returned by DirMgr{}.NewFromPathFileNameExtStr(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
  }

  dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

  if err != nil {
    t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute).  Error='%v'", err.Error())
  }

  if expectedAbsPath != dMgrResult.path {
    t.Errorf("Expected final substituted path = '%v'.  Instead substituted path = '%v' ",
      expectedAbsPath, dMgrResult.path)
  }

  if expectedAbsPath != dMgrResult.absolutePath {
    t.Errorf("Expected final substituted absolute path = '%v'.  Instead substituted absolute path = '%v' ",
      expectedAbsPath, dMgrResult.absolutePath)
  }

}

