package pathfileops

import "testing"

func TestFileOpsCollection_AddByPathFileNameExtStrs_01(t *testing.T) {

	sf := make([]string, 5, 10)

	sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
	sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
	sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
	sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
	sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

	df := make([]string, 5, 10)

	df[0] = "../dirmgrtests/level_0_0_test.txt"
	df[1] = "../dirmgrtests/level_0_1_test.txt"
	df[2] = "../dirmgrtests/level_0_2_test.txt"
	df[3] = "../dirmgrtests/level_0_3_test.txt"
	df[4] = "../dirmgrtests/level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
				"i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
			return
		}

		df[i], err = fh.GetAbsPathFromFilePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
			return
		}

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected intial array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	for j := 0; j < arrayLen; j++ {

		fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		srcFmgr := fOps.GetSource()

		if sf[j] != srcFmgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected source file manager[j] = '%v'. "+
				"Instead, source file manager[j]='%v'. j='%v' ",
				sf[j], srcFmgr.GetAbsolutePathFileName(), j)
		}

		destFmgr := fOps.GetDestination()

		if df[j] != destFmgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected destination file manager[j] = '%v'. "+
				"Instead, destination file manager[j]='%v'. j='%v' ",
				df[j], destFmgr.GetAbsolutePathFileName(), j)
		}

	}

}

func TestFileOpsCollection_AddByDirStrsAndFileNameExtStrs_01(t *testing.T) {

	srcDir := "../filesfortest/levelfilesfortest/"

	sf := make([]string, 5, 10)

	sf[0] = "level_0_0_test.txt"
	sf[1] = "level_0_1_test.txt"
	sf[2] = "level_0_2_test.txt"
	sf[3] = "level_0_3_test.txt"
	sf[4] = "level_0_4_test.txt"

	destDir := "../dirmgrtests"
	df := make([]string, 5, 10)

	df[0] = "level_0_0_test.txt"
	df[1] = "level_0_1_test.txt"
	df[2] = "level_0_2_test.txt"
	df[3] = "level_0_3_test.txt"
	df[4] = "level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		err := fOpsCol.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], destDir, df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByDirStrsAndFileNameExtStrs(srcDir, sf[i], "+
				"destDir, df[i]). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		fullSrcPathFileName := fh.JoinPathsAdjustSeparators(srcDir, sf[i])

		sf[i], err = fh.GetAbsPathFromFilePath(fullSrcPathFileName)

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(fullSrcPathFileName). "+
				"i='%v', fullSrcPathFileName='%v' Error='%v' ", i, fullSrcPathFileName, err.Error())
			return
		}

		fullDestPathFileName := fh.JoinPathsAdjustSeparators(destDir, df[i])

		df[i], err = fh.GetAbsPathFromFilePath(fullDestPathFileName)

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(fullDestPathFileName). "+
				"i='%v', fullDestPathFileName='%v' Error='%v' ", i, fullDestPathFileName, err.Error())
			return
		}

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected intial array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	for j := 0; j < arrayLen; j++ {

		fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
			return
		}

		srcFMgr := fOps.GetSource()

		if sf[j] != srcFMgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected source file manager[j] = '%v'. "+
				"Instead, source file manager[j]='%v'. j='%v' ",
				sf[j], srcFMgr.GetAbsolutePathFileName(), j)
		}

		destFMgr := fOps.GetDestination()

		if df[j] != destFMgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected destination file manager[j] = '%v'. "+
				"Instead, destination file manager[j]='%v'. j='%v' ",
				df[j], destFMgr.GetAbsolutePathFileName(), j)
		}

	}

}

func TestFileOpsCollection_AddByFileMgrs_01(t *testing.T) {

	sf := make([]string, 5, 10)

	sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
	sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
	sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
	sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
	sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

	df := make([]string, 5, 10)

	df[0] = "../dirmgrtests/level_0_0_test.txt"
	df[1] = "../dirmgrtests/level_0_1_test.txt"
	df[2] = "../dirmgrtests/level_0_2_test.txt"
	df[3] = "../dirmgrtests/level_0_3_test.txt"
	df[4] = "../dirmgrtests/level_0_4_test.txt"

	fh := FileHelper{}
	fileOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		srcFileMgr, err :=
			FileMgr{}.NewFromPathFileNameExtStr(sf[i])

		if err != nil {
			t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(sf[i]). "+
				"sf[i]='%v' Error='%v' ", sf[i], err.Error())
			return
		}

		destFileMgr, err :=
			FileMgr{}.NewFromPathFileNameExtStr(df[i])

		if err != nil {
			t.Errorf("Error returned by FileMgr{}.NewFromPathFileNameExtStr(df[i]). "+
				"df[i]='%v' Error='%v' ", df[i], err.Error())
			return
		}

		err = fileOpsCol.AddByFileMgrs(srcFileMgr, destFileMgr)

		if err != nil {
			t.Errorf("Error returned by fileOpsCol.AddByFileMgrs(srcFileMgr, destFileMgr). "+
				"srcFileMgr='%v' destFileMgr='%v' Error='%v' ",
				srcFileMgr.GetAbsolutePathFileName(), destFileMgr.GetAbsolutePathFileName(),
				err.Error())
		}

		sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
				"i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
			return
		}

		df[i], err = fh.GetAbsPathFromFilePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
			return
		}

	}

	arrayLen := fileOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Expected final FileOpsCol array length='5'. "+
			"Instead, final array length='%v' ", arrayLen)
	}

	for j := 0; j < arrayLen; j++ {

		fileOps, err := fileOpsCol.PeekFileOpsAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fileOpsCol.PeekFileOpsAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
		}

		srcFileMgr := fileOps.GetSource()

		if sf[j] != srcFileMgr.GetAbsolutePathFileName() {
			t.Errorf("Expected srcFileMgr['%v']='%v'. "+
				"Instead, srcFileMgr='%v' ", j, sf[j],
				srcFileMgr.GetAbsolutePathFileName())
		}

		destFileMgr := fileOps.GetDestination()

		if df[j] != destFileMgr.GetAbsolutePathFileName() {
			t.Errorf("Expected destFileMgr['%v']='%v'. "+
				"Instead, destFileMgr='%v' ", j, df[j],
				destFileMgr.GetAbsolutePathFileName())
		}
	}
}

func TestFileOpsCollection_AddByDirMgrFileName_01(t *testing.T) {

	srcDir := "../filesfortest/levelfilesfortest/"

	srcDirMgr, err := DirMgr{}.New(srcDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(srcDir). "+
			"srcDir='%v' Error='%v' ", srcDir, err.Error())
	}

	sf := make([]string, 5, 10)

	sf[0] = "level_0_0_test.txt"
	sf[1] = "level_0_1_test.txt"
	sf[2] = "level_0_2_test.txt"
	sf[3] = "level_0_3_test.txt"
	sf[4] = "level_0_4_test.txt"

	destDir := "../dirmgrtests"

	destDirMgr, err := DirMgr{}.New(destDir)

	if err != nil {
		t.Errorf("Error returned from DirMgr{}.New(destDir). "+
			"destDir='%v' Error='%v' ", destDir, err.Error())
	}

	df := make([]string, 5, 10)

	df[0] = "level_0_0_test.txt"
	df[1] = "level_0_1_test.txt"
	df[2] = "level_0_2_test.txt"
	df[3] = "level_0_3_test.txt"
	df[4] = "level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {
		err := fOpsCol.AddByDirMgrFileName(
			srcDirMgr.CopyOut(),
			sf[i],
			destDirMgr.CopyOut(),
			df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByDirMgrFileName(...). "+
				"i='%v' Error='%v'  ", i, err.Error())
		}

		sf[i] = fh.JoinPathsAdjustSeparators(srcDirMgr.GetAbsolutePath(), sf[i])

		df[i] = fh.JoinPathsAdjustSeparators(destDirMgr.GetAbsolutePath(), df[i])

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected file fOpsCol array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	for j := 0; j < arrayLen; j++ {

		fileOps, err := fOpsCol.PeekFileOpsAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())
		}

		srcFileMgr := fileOps.GetSource()

		if sf[j] != srcFileMgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected source file manager[j]='%v'. "+
				"Instead, source file manger[j]='%v' ",
				sf[j], srcFileMgr.GetAbsolutePathFileName())
		}

		destFileMgr := fileOps.GetDestination()

		if df[j] != destFileMgr.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected destination file manager[j]='%v'. "+
				"Instead, destination file manger[j]='%v' ",
				df[j], destFileMgr.GetAbsolutePathFileName())
		}

	}
}

func TestFileOpsCollection_GetFileOpsAtIndex_01(t *testing.T) {

	sf := make([]string, 5, 10)

	sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
	sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
	sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
	sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
	sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

	df := make([]string, 5, 10)

	df[0] = "../dirmgrtests/level_0_0_test.txt"
	df[1] = "../dirmgrtests/level_0_1_test.txt"
	df[2] = "../dirmgrtests/level_0_2_test.txt"
	df[3] = "../dirmgrtests/level_0_3_test.txt"
	df[4] = "../dirmgrtests/level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
				"i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
			return
		}

		df[i], err = fh.GetAbsPathFromFilePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
			return
		}

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected intial array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	fOps, err := fOpsCol.GetFileOpsAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.GetFileOpsAtIndex(2). "+
			"Error='%v' ", err.Error())
		return
	}

	srcFile := fOps.GetSource()

	if sf[2] != srcFile.GetAbsolutePathFileName() {
		t.Errorf("Error: Expected source file[2]='%v'. "+
			"Instead, source file[2]='%v' ", sf[2], srcFile.GetAbsolutePathFileName())
	}

	destFile := fOps.GetDestination()

	if df[2] != destFile.GetAbsolutePathFileName() {
		t.Errorf("Error: Expected destination file[2]='%v'. "+
			"Instead, destination file[2]='%v' ", df[2], destFile.GetAbsolutePathFileName())
	}

}

func TestFileOpsCollection_DeleteAtIndex_01(t *testing.T) {

	sf := make([]string, 5, 10)

	sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
	sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
	sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
	sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
	sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

	df := make([]string, 5, 10)

	df[0] = "../dirmgrtests/level_0_0_test.txt"
	df[1] = "../dirmgrtests/level_0_1_test.txt"
	df[2] = "../dirmgrtests/level_0_2_test.txt"
	df[3] = "../dirmgrtests/level_0_3_test.txt"
	df[4] = "../dirmgrtests/level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
				"i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
			return
		}

		df[i], err = fh.GetAbsPathFromFilePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
			return
		}

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected intial array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	err := fOpsCol.DeleteAtIndex(2)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fOpsCol.GetNumOfFileOps()

	if arrayLen != 4 {
		t.Errorf("Expected array length=4 after deletion. "+
			"Instead, array length='%v'", arrayLen)
	}

	for j := 0; j < arrayLen; j++ {

		fOps, err := fOpsCol.PeekFileOpsAtIndex(j)

		if err != nil {
			t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(j). "+
				"j='%v' Error='%v' ", j, err.Error())

		}

		if sf[2] == fOps.source.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected index 2 would be deleted. I was NOT! "+
				"Found source path file name='%v' at index='%v' ",
				fOps.source.GetAbsolutePathFileName(), j)
		}

		if df[2] == fOps.destination.GetAbsolutePathFileName() {
			t.Errorf("Error: Expected index 2 would be deleted. I was NOT! "+
				"Found destination path file name='%v' at index='%v' ",
				fOps.source.GetAbsolutePathFileName(), j)
		}

	}
}

func TestFileOpsCollection_DeleteAtIndex_02(t *testing.T) {

	sf := make([]string, 5, 10)

	sf[0] = "../filesfortest/levelfilesfortest/level_0_0_test.txt"
	sf[1] = "../filesfortest/levelfilesfortest/level_0_1_test.txt"
	sf[2] = "../filesfortest/levelfilesfortest/level_0_2_test.txt"
	sf[3] = "../filesfortest/levelfilesfortest/level_0_3_test.txt"
	sf[4] = "../filesfortest/levelfilesfortest/level_0_4_test.txt"

	df := make([]string, 5, 10)

	df[0] = "../dirmgrtests/level_0_0_test.txt"
	df[1] = "../dirmgrtests/level_0_1_test.txt"
	df[2] = "../dirmgrtests/level_0_2_test.txt"
	df[3] = "../dirmgrtests/level_0_3_test.txt"
	df[4] = "../dirmgrtests/level_0_4_test.txt"

	fh := FileHelper{}
	fOpsCol := FileOpsCollection{}.New()

	for i := 0; i < 5; i++ {

		err := fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i])

		if err != nil {
			t.Errorf("Error returned by fOpsCol.AddByPathFileNameExtStrs(sf[i], df[i]). "+
				"i='%v' Error='%v' ", i, err.Error())
			return
		}

		sf[i], err = fh.GetAbsPathFromFilePath(sf[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(sf[i]). "+
				"i='%v', sf[i]='%v' Error='%v' ", i, sf[i], err.Error())
			return
		}

		df[i], err = fh.GetAbsPathFromFilePath(df[i])

		if err != nil {
			t.Errorf("Error returned by fh.GetAbsPathFromFilePath(df[i]). "+
				"i='%v', df[i]='%v' Error='%v' ", i, df[i], err.Error())
			return
		}

	}

	arrayLen := fOpsCol.GetNumOfFileOps()

	if arrayLen != 5 {
		t.Errorf("Error: Expected intial array length='5'. "+
			"Instead, array length='%v' ", arrayLen)
	}

	err := fOpsCol.DeleteAtIndex(4)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.DeleteAtIndex(4). "+
			"Error='%v' ", err.Error())
	}

	err = fOpsCol.DeleteAtIndex(0)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.DeleteAtIndex(0). "+
			"Error='%v' ", err.Error())
	}

	err = fOpsCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
			"Error='%v' ", err.Error())
	}

	err = fOpsCol.DeleteAtIndex(1)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.DeleteAtIndex(2). "+
			"Error='%v' ", err.Error())
	}

	arrayLen = fOpsCol.GetNumOfFileOps()

	if arrayLen != 1 {
		t.Errorf("Expected array length=1 after deletion. "+
			"Instead, array length='%v'", arrayLen)
	}

	fOps, err := fOpsCol.PeekFileOpsAtIndex(0)

	if err != nil {
		t.Errorf("Error returned by fOpsCol.PeekFileOpsAtIndex(0). "+
			"Error='%v' ", err.Error())
	}

	srcFileMgr := fOps.GetSource()

	if sf[1] != srcFileMgr.GetAbsolutePathFileName() {
		t.Errorf("Error: Expected source path file name='%v'. "+
			"Instead source path file name='%v' ",
			sf[1], srcFileMgr.GetAbsolutePathFileName())
	}

	destinationFileMgr := fOps.GetDestination()

	if df[1] != destinationFileMgr.GetAbsolutePathFileName() {
		t.Errorf("Error: Expected destination path file name='%v'. "+
			"Instead destination path file name='%v' ",
			df[1], destinationFileMgr.GetAbsolutePathFileName())
	}

}
