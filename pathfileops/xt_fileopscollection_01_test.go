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
