package common

import (
	fp "path/filepath"
	"testing"
	"strings"
)



func TestDirMgr_MakeDir_01(t *testing.T) {

	fh := FileHelper{}

	origDir :=  fh.AdjustPathSlash("../checkfiles/checkfiles99/checkfiles999")

	dMgr, err := DirMgr{}.New(origDir)


	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir='%v'  Error='%v'",origDir, err.Error())
	}

	if dMgr.AbsolutePathDoesExist {
		dMgr.DeleteAll()
	}

	if dMgr.AbsolutePathDoesExist {
		t.Errorf("Error: Attempted to delete dMgr.AbsolutePath='%v'. Deletion Attempt FAILED. This directory still exists.", dMgr.AbsolutePath)
	}

	dMgr.MakeDir()


	if !dMgr.AbsolutePathDoesExist {
		t.Errorf("Error: Attempted to create dMgr.AbsolutePath='%v'. Creation Attempt FAILED. This directory does NOT exist.", dMgr.AbsolutePath)
	}

}

func TestDirMgr_New_01(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles2")
	expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
	expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(origDir)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(origDir)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedParentPath := strings.TrimSuffix(expectedAbsDir,fh.AdjustPathSlash("/testfiles2"))
	expectedIsParentPathPopulated := false

	if expectedParentPath != "" {
		expectedIsParentPathPopulated = true
	}


	expectedRelativePath := "testfiles2"
	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
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

	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_02(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("./")
	expectedPath := fh.AdjustPathSlash(".")
	expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(origDir)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(origDir)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedIsParentPathPopulated := true

	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}


	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}


	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_03(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../")
	expectedPath := fh.AdjustPathSlash("..")
	expectedAbsDir, err := fh.MakeAbsolutePath(origDir)

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir,err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(origDir)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(origDir)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedIsParentPathPopulated := true

	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}


	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}


	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_04(t *testing.T) {
	fh := FileHelper{}
	xDir, err := fh.GetCurrentDir()

	if err != nil {
		t.Errorf("Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
	}

	volName := fp.VolumeName(xDir)

	origDir :=  fh.AdjustPathSlash(volName )
	expectedPath := fh.AdjustPathSlash(origDir)
	expectedAbsDir := origDir

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(origDir)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(origDir)
	expectedAbsPathDifferentFromPath := false
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedIsParentPathPopulated := false

	expectedIsRelativePathPopulated := false


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}


	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}


	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
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

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(origDir)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(origDir)
	expectedAbsPathDifferentFromPath := false
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedIsParentPathPopulated := true

	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
	}

	if expectedVolumeName != dMgr.VolumeName {
		t.Errorf("Expected VolumeName=='%v'. Instead, VolumeName=='%v'.", expectedVolumeName, dMgr.VolumeName)
	}

	if expectedVolumeIsPopulated != dMgr.VolumeIsPopulated {
		t.Errorf("Expected VolumeIsPopulated=='%v'. Instead, VolumeIsPopulated=='%v'.", expectedVolumeIsPopulated, dMgr.VolumeIsPopulated)
	}


	if expectedIsParentPathPopulated != dMgr.ParentPathIsPopulated {
		t.Errorf("Expected ParentPathIsPopulated=='%v'. Instead, ParentPathIsPopulated=='%v'.", expectedIsParentPathPopulated, dMgr.ParentPathIsPopulated)
	}


	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_06(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles2/test2007.txt")
	expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
	expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)


	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(expectedPath)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedParentPath := strings.TrimSuffix(expectedAbsDir,fh.AdjustPathSlash("/testfiles2"))
	expectedIsParentPathPopulated := false

	if expectedParentPath != "" {
		expectedIsParentPathPopulated = true
	}


	expectedRelativePath := "testfiles2"
	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
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

	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_07(t *testing.T) {
	fh := FileHelper{}
	origDir :=  fh.AdjustPathSlash("../testfiles/testfiles2/")
	expectedPath := fh.AdjustPathSlash("../testfiles/testfiles2")
	expectedAbsDir, err := fh.MakeAbsolutePath(expectedPath)

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(origDir). origDir=='%v'  Error='%v'", origDir, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(expectedPath)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedParentPath := strings.TrimSuffix(expectedAbsDir,fh.AdjustPathSlash("/testfiles2"))
	expectedIsParentPathPopulated := false

	if expectedParentPath != "" {
		expectedIsParentPathPopulated = true
	}


	expectedRelativePath := "testfiles2"
	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(origDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(origDir). origDir=='%v' Error='%v'", origDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
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

	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_08(t *testing.T) {
	fh := FileHelper{}
	rawDir :=  "D:/"
	expectedPath := fh.AdjustPathSlash("D:")
	expectedAbsDir := expectedPath

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(expectedPath)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsDir)
	expectedAbsPathDifferentFromPath := false
	expectedVolumeName := fp.VolumeName(expectedAbsDir)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedParentPath := ""
	expectedIsParentPathPopulated := false

	if expectedParentPath != "" {
		expectedIsParentPathPopulated = true
	}


	expectedRelativePath := ""
	expectedIsRelativePathPopulated := false


	dMgr, err := DirMgr{}.New(rawDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawDir). rawDir=='%v' Error='%v'", rawDir, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsDir != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsDir, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
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

	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_New_09(t *testing.T) {
	fh := FileHelper{}
	rawPath := "../filesfortest/newfilesfortest"
	expectedPath := fh.AdjustPathSlash(rawPath)
	expectedAbsPath, err := fh.MakeAbsolutePath(expectedPath)

	if err!= nil {
		t.Errorf("Error returned from fh.GetAbsPathFromFilePath(expectedPath). expectedPath=='%v'  Error='%v'", expectedPath, err.Error())
	}

	expectedIsInitialized := true
	expectedIsPathPopulated := true
	expectedPathDoesExist := fh.DoesFileExist(expectedPath)
	expectedIsAbsPathPopulated := true
	expectedAbsPathDoesExist := fh.DoesFileExist(expectedAbsPath)
	expectedAbsPathDifferentFromPath := true
	expectedVolumeName := fp.VolumeName(expectedAbsPath)
	var expectedVolumeIsPopulated bool

	if expectedVolumeName != "" {
		expectedVolumeIsPopulated = true
	} else {
		expectedVolumeIsPopulated = false
	}

	expectedParentPath := strings.TrimSuffix(expectedAbsPath,fh.AdjustPathSlash("/newfilesfortest"))
	expectedIsParentPathPopulated := false

	if expectedParentPath != "" {
		expectedIsParentPathPopulated = true
	}


	expectedRelativePath := "newfilesfortest"
	expectedIsRelativePathPopulated := true


	dMgr, err := DirMgr{}.New(rawPath)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(rawPath). rawPath=='%v' Error='%v'", rawPath, err.Error())
	}

	if expectedIsInitialized != dMgr.IsInitialized {
		t.Errorf("Expected DirMgr.IsFInfoInitialized=='%v'. Instead, DirMgr.IsFInfoInitialized=='%v'", expectedIsInitialized, dMgr.IsInitialized)
	}

	if expectedIsPathPopulated != dMgr.PathIsPopulated {
		t.Errorf("Expected DirMgr.PathIsPopulated=='%v'. Instead, DirMgr.PathIsPopulated=='%v'", expectedIsPathPopulated, dMgr.PathIsPopulated)
	}

	if expectedPath != dMgr.Path {
		t.Errorf("Expected Path=='%v'. Instead, Path=='%v'.", expectedPath, dMgr.Path)
	}

	if expectedPathDoesExist != dMgr.PathDoesExist {
		t.Errorf("Expected PathDoesExist=='%v'. Instead, PathDoesExist=='%v'.", expectedPathDoesExist, dMgr.PathDoesExist)
	}

	if expectedIsAbsPathPopulated != dMgr.AbsolutePathIsPopulated {
		t.Errorf("Expected AbsolutePathIsPopulated=='%v'. Instead, AbsolutePathIsPopulated=='%v'.", expectedIsAbsPathPopulated, dMgr.AbsolutePathIsPopulated)
	}

	if expectedAbsPath != dMgr.AbsolutePath {
		t.Errorf("Expected AbsolutePath=='%v'. Instead, AbsolutePath=='%v'", expectedAbsPath, dMgr.AbsolutePath )
	}

	if expectedAbsPathDoesExist != dMgr.AbsolutePathDoesExist {
		t.Errorf("Expected AbsolutePathDoesExist=='%v'. Instead, AbsolutePathDoesExist=='%v'.", expectedAbsPathDoesExist, dMgr.AbsolutePathDoesExist)
	}

	if expectedAbsPathDifferentFromPath != dMgr.AbsolutePathDifferentFromPath {
		t.Errorf("Expected AbsolutePathDifferentFromPath=='%v'. Instead, AbsolutePathDifferentFromPath=='%v'.", expectedAbsPathDifferentFromPath, dMgr.AbsolutePathDifferentFromPath)
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

	if expectedIsRelativePathPopulated != dMgr.RelativePathIsPopulated {
		t.Errorf("Expected RelativePathIsPopulated=='%v'. Instead, RelativePathIsPopulated=='%v'.", expectedIsRelativePathPopulated, dMgr.RelativePathIsPopulated)
	}

}

func TestDirMgr_SubstituteBaseDir_01(t *testing.T) {

	fh := FileHelper{}

	rawOrigPath := "../dirwalktests/dir01/dir02/dir03"

	rawBasePath := "../dirwalktests/dir01"

	substitutePath := "../checkfiles"

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
		t.Errorf("Expected final substituted path = '%v'.  Instead substituted path = '%v' ",expectedPath, dMgrResult.Path)
	}

	if expectedAbsPath != dMgrResult.AbsolutePath {
		t.Errorf("Expected final substituted absolute path = '%v'.  Instead substituted absolute path = '%v' ",expectedAbsPath, dMgrResult.AbsolutePath)
	}

}

func TestDirMgr_SubstituteBaseDir_02(t *testing.T) {

	fh := FileHelper{}

	rawOrigPath := "../dirwalktests/dir01/dir02/dir03/"

	rawBasePath := "../dirwalktests/dir01/"

	substitutePath := "../checkfiles/"

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
		t.Errorf("Expected final substituted path = '%v'.  Instead substituted path = '%v' ",expectedPath, dMgrResult.Path)
	}

	if expectedAbsPath != dMgrResult.AbsolutePath {
		t.Errorf("Expected final substituted absolute path = '%v'.  Instead substituted absolute path = '%v' ",expectedAbsPath, dMgrResult.AbsolutePath)
	}

}



