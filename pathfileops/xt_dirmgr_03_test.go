package pathfileops

import (
	fp "path/filepath"
	"strings"
	"testing"
)

func TestDirMgr_MakeDir_01(t *testing.T) {

	fh := FileHelper{}

	origDir := fh.AdjustPathSlash("../checkfiles/checkfiles99/checkfiles999")

	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir='%v'  Error='%v'", origDir, err.Error())
	}

	if dMgr.AbsolutePathDoesExist {

		err = dMgr.DeleteAll()

		if err != nil {
			t.Errorf("%v", err.Error())
		}

	}

	if dMgr.AbsolutePathDoesExist {
		t.Errorf("Error: Attempted to delete dMgr.AbsolutePath='%v'. Deletion Attempt FAILED. This directory still exists.", dMgr.AbsolutePath)
	}

	err = dMgr.MakeDir()

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	if !dMgr.AbsolutePathDoesExist {
		t.Errorf("Error: Attempted to create dMgr.AbsolutePath='%v'. Creation Attempt FAILED. This directory does NOT exist.", dMgr.AbsolutePath)
	}

}

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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if expectedParentPath != dMgr.ParentPath {
		t.Errorf("Expected ParentPath=='%v'. Instead, ParentPath=='%v'.", expectedParentPath, dMgr.ParentPath)
	}

	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}

	if expectedRelativePath != dMgr.RelativePath {
		t.Errorf("Expected RelativePath=='%v'. Instead, RelativePath=='%v'.", expectedRelativePath, dMgr.RelativePath)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if true != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", true, dMgr.ParentPathIsPopulated)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'",
			origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.",
			expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.",
			true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.",
			expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.",
			expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if true != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			true, dMgr.ParentPathIsPopulated)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if false != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.",
			false, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if false != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			false, dMgr.ParentPathIsPopulated)
	}

	if false != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			false, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if false != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.",
			false, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if true != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			true, dMgr.ParentPathIsPopulated)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.",
			true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.",
			expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.",
			expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if expectedParentPath != dMgr.ParentPath {
		t.Errorf("Expected ParentPath=='%v'. Instead, ParentPath=='%v'.",
			expectedParentPath, dMgr.ParentPath)
	}

	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}

	if expectedRelativePath != dMgr.RelativePath {
		t.Errorf("Expected RelativePath=='%v'. Instead, RelativePath=='%v'.",
			expectedRelativePath, dMgr.RelativePath)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.",
			expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, "+
			"AbsolutePathDifferentFromPath=='%v'.",
			true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.",
			expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.",
			expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if expectedParentPath != dMgr.ParentPath {
		t.Errorf("Expected ParentPath=='%v'. Instead, ParentPath=='%v'.",
			expectedParentPath, dMgr.ParentPath)
	}

	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}

	if expectedRelativePath != dMgr.RelativePath {
		t.Errorf("Expected RelativePath=='%v'. Instead, RelativePath=='%v'.",
			expectedRelativePath, dMgr.RelativePath)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			true, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(rawDir). rawDir=='%v' Error='%v'", rawDir, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.",
			expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.",
			expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsDir, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if false != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. "+
			"Instead, AbsolutePathDifferentFromPath=='%v'.",
			false, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.",
			expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.",
			expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if expectedParentPath != dMgr.ParentPath {
		t.Errorf("Expected ParentPath=='%v'. Instead, ParentPath=='%v'.",
			expectedParentPath, dMgr.ParentPath)
	}

	if false != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			false, dMgr.ParentPathIsPopulated)
	}

	if expectedRelativePath != dMgr.RelativePath {
		t.Errorf("Expected RelativePath=='%v'. Instead, RelativePath=='%v'.",
			expectedRelativePath, dMgr.RelativePath)
	}

	if false != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			false, dMgr.RelativePathIsPopulated)
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
		t.Errorf("Error returned from DirMgr{}.New(rawPath). rawPath=='%v' Error='%v'",
			rawPath, err.Error())
	}

	if true != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'",
			true, dMgr.IsInitialized)
	}

	if true != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'",
			true, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.",
			expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if true != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.",
			true, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsPath != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'",
			expectedAbsPath, dMgr.AbsolutePath)
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.",
			expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if true != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. "+
			"Instead, AbsolutePathDifferentFromPath=='%v'.",
			true, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.",
			expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.",
			expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}

	if expectedParentPath != dMgr.ParentPath {
		t.Errorf("Expected ParentPath=='%v'. Instead, ParentPath=='%v'.",
			expectedParentPath, dMgr.ParentPath)
	}

	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.",
			expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}

	if expectedRelativePath != dMgr.RelativePath {
		t.Errorf("Expected RelativePath=='%v'. Instead, RelativePath=='%v'.",
			expectedRelativePath, dMgr.RelativePath)
	}

	if true != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.",
			true, dMgr.RelativePathIsPopulated)
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

	dMgrOrig, err := DirMgr{}.New(rawOrigPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
	}

	dMgrBase, err := DirMgr{}.New(rawBasePath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
	}

	dMgrSubstitute, err := DirMgr{}.New(substitutePath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
	}

	dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

	if err != nil {
		t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, "+
			"dMgrSubstitute).  Error='%v'",
			err.Error())
	}

	if expectedPath != dMgrResult.Path {
		t.Errorf("Expected final substituted path = '%v'.  Instead substituted "+
			"path = '%v' ",
			expectedPath, dMgrResult.Path)
	}

	if expectedAbsPath != dMgrResult.AbsolutePath {
		t.Errorf("Expected final substituted absolute path = '%v'.  Instead "+
			"substituted absolute path = '%v' ",
			expectedAbsPath, dMgrResult.AbsolutePath)
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

	dMgrOrig, err := DirMgr{}.New(rawOrigPath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(rawOrigPath). rawOrigPath='%v'  Error='%v'", rawOrigPath, err.Error())
	}

	dMgrBase, err := DirMgr{}.New(rawBasePath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(rawBasePath). rawBasePath='%v'  Error='%v'", rawBasePath, err.Error())
	}

	dMgrSubstitute, err := DirMgr{}.New(substitutePath)

	if err != nil {
		t.Errorf("Error returned by DirMgr{}.New(substitutePath). substitutePath='%v'  Error='%v'", substitutePath, err.Error())
	}

	dMgrResult, err := dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute)

	if err != nil {
		t.Errorf("Error returned by dMgrOrig.SubstituteBaseDir(dMgrBase, dMgrSubstitute).  Error='%v'", err.Error())
	}

	if expectedPath != dMgrResult.Path {
		t.Errorf("Expected final substituted path = '%v'.  Instead substituted path = '%v' ", expectedPath, dMgrResult.Path)
	}

	if expectedAbsPath != dMgrResult.AbsolutePath {
		t.Errorf("Expected final substituted absolute path = '%v'.  Instead substituted absolute path = '%v' ", expectedAbsPath, dMgrResult.AbsolutePath)
	}

}
