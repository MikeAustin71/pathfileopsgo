package appExamples

import (
	appLib "../appLibs"
	pathFileOps "../pathfileops"
	"fmt"
	"time"
)

func WalkDirFindFiles(
	dMgr pathFileOps.DirMgr,
	filePattern string,
	filesOlderThan, filesNewerThan time.Time) error {

	ePrefix := "WalkDirFindFiles() "

	du := appLib.DateTimeUtility{}

	fsc := pathFileOps.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{filePattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOps.ANDFILESELECTCRITERION

	dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		return fmt.Errorf("Error returned from dMgr.FindWalkDirFiles(filePattern, "+
			"filesOlderThan, filesNewerThan) filePattern='%v'  filesOlderThan='%v' "+
			"filesNewerThan='%v' Error='%v' \n",
			filePattern, du.GetDateTimeYMDAbbrvDowNano(filesOlderThan),
			du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
	}

	PrintDirMgrFields(dMgr)
	fmt.Println("----------------------------")
	fmt.Println("     startPath: ", dWalkInfo.StartPath)
	fmt.Println("   filePattern: ", filePattern)
	fmt.Println("filesOlderThan: ", du.GetDateTimeYMDAbbrvDowNano(filesOlderThan))
	fmt.Println("filesNewerThan: ", du.GetDateTimeYMDAbbrvDowNano(filesNewerThan))

	if dWalkInfo.FoundFiles.GetArrayLength() == 0 {
		fmt.Println("No Files Found")
	} else {
		fmt.Println("Files Found: ")
		for i := 0; i < dWalkInfo.FoundFiles.GetArrayLength(); i++ {

			fInfoPlus, err := dWalkInfo.FoundFiles.FMgrs[i].GetFileInfoPlus()

			if err != nil {
				return fmt.Errorf(ePrefix+
					"Error returned from dWalkInfo.FoundFiles.FMgrs[i].GetFileInfoPlus() \n"+
					"i='%v' FileName='%v' Error='%v' \n", i,
					dWalkInfo.FoundFiles.FMgrs[i].GetAbsolutePathFileName(),
					err.Error())
			}

			fmt.Printf("  Name: %v Mod Date: %v path: %v \n",
				dWalkInfo.FoundFiles.FMgrs[i].GetFileNameExt(),
				fInfoPlus.ModTime(), fInfoPlus.DirPath())
		}
		fmt.Println()
	}

	if dWalkInfo.Directories.GetNumOfDirs() > 0 {
		fmt.Println("DirMgrs Found:")

		for k := 0; k < dWalkInfo.Directories.GetNumOfDirs(); k++ {
			fmt.Printf("Dir: %v \n", dWalkInfo.Directories.DirMgrs[k].GetPath())
		}

	}

	if len(dWalkInfo.ErrReturns) > 0 {
		fmt.Println("Errors Found: ")

		for j := 0; j < len(dWalkInfo.ErrReturns); j++ {
			fmt.Printf("Error: %v \n", dWalkInfo.ErrReturns[j])
		}
		fmt.Println("")
	}

	return nil
}

func WalkDirFindFiles2(
	dMgr pathFileOps.DirMgr,
	filePattern string,
	filesOlderThan,
	filesNewerThan time.Time) error {

	ePrefix := "WalkDirFindFiles2() "

	du := appLib.DateTimeUtility{}
	fsc := pathFileOps.FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{filePattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = pathFileOps.ANDFILESELECTCRITERION

	dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from dMgr.FindWalkDirFiles(filePattern, filesOlderThan, "+
			"filesNewerThan) filePattern='%v'  filesOlderThan='%v' filesNewerThan='%v' "+
			"Error='%v' \n",
			filePattern,
			du.GetDateTimeYMDAbbrvDowNano(filesOlderThan),
			du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
	}

	PrintDirMgrFields(dMgr)
	fmt.Println("----------------------------")
	fmt.Println("     startPath: ", dWalkInfo.StartPath)
	fmt.Println("   filePattern: ", filePattern)
	fmt.Println("filesOlderThan: ", du.GetDateTimeYMDAbbrvDowNano(filesOlderThan))
	fmt.Println("filesNewerThan: ", du.GetDateTimeYMDAbbrvDowNano(filesNewerThan))

	if dWalkInfo.FoundFiles.GetArrayLength() == 0 {
		fmt.Println("No Files Found")
	} else {
		fmt.Println("Files Found: ")
		for i := 0; i < dWalkInfo.FoundFiles.GetArrayLength(); i++ {
			fileInfoPlus, err := dWalkInfo.FoundFiles.FMgrs[i].GetFileInfoPlus()

			if err != nil {
				return fmt.Errorf(ePrefix+
					"Error returned by dWalkInfo.FoundFiles.FMgrs[i].GetFileInfoPlus(). "+
					"i='%v' FileName='%v' Error='%v' ",
					i, dWalkInfo.FoundFiles.FMgrs[i].GetAbsolutePathFileName(), err.Error())
			}

			fmt.Printf("  Name: %v SysInfo: %v \n",
				dWalkInfo.FoundFiles.FMgrs[i].GetFileNameExt(),
				fileInfoPlus.Sys())
		}
		fmt.Println()
	}

	if dWalkInfo.Directories.GetNumOfDirs() > 0 {
		fmt.Println("DirMgrs Found:")

		for k := 0; k < dWalkInfo.Directories.GetNumOfDirs(); k++ {
			fmt.Printf("Dir: %v \n", dWalkInfo.Directories.DirMgrs[k].GetPath())
		}

	}

	if len(dWalkInfo.ErrReturns) > 0 {
		fmt.Println("Errors Found: ")

		for j := 0; j < len(dWalkInfo.ErrReturns); j++ {
			fmt.Printf("Error: %v \n", dWalkInfo.ErrReturns[j])
		}
		fmt.Println("")
	}

	return nil
}

func PrintDirMgrFields(dMgr pathFileOps.DirMgr) {

	ePrefix := "PrintDirMgrFields() "

	du := appLib.DateTimeUtility{}
	fmt.Println("-----------------------------------------")
	fmt.Println(" 	DirMgr Fields")
	fmt.Println("-----------------------------------------")

	fmt.Println("                isInitialized: ", dMgr.IsInitialized())
	fmt.Println("                Original path: ", dMgr.GetOriginalPath())
	fmt.Println("                         path: ", dMgr.GetPath())
	fmt.Println("             PathIsPopuslated: ", dMgr.IsPathPopulated())
	fmt.Println("                doesPathExist: ", dMgr.DoesDirMgrPathExist())
	fmt.Println("                   parentPath: ", dMgr.GetParentPath())
	fmt.Println("        isParentPathPopulated: ", dMgr.IsParentPathPopulated())
	fmt.Println("                 relativePath: ", dMgr.GetRelativePath())
	fmt.Println("      isRelativePathPopulated: ", dMgr.IsRelativePathPopulated())
	fmt.Println("                 absolutePath: ", dMgr.GetAbsolutePath())
	fmt.Println("      isAbsolutePathPopulated: ", dMgr.IsAbsolutePathPopulated())
	fmt.Println("isAbsolutePathDifferentFromPath: ", dMgr.IsAbsolutePathDifferentFromPath())
	fmt.Println("        doesAbsolutePathExist: ", dMgr.DoesDirMgrAbsolutePathExist())
	fmt.Println("               Directory Name: ", dMgr.GetDirectoryName())
	fmt.Println("                   volumeName: ", dMgr.GetVolumeName())
	fmt.Println("            isVolumePopulated: ", dMgr.IsVolumeNamePopulated())
	fmt.Println("============== File Info Data ============")

	actualDirFileInfo, err := dMgr.GetFileInfoPlus()

	if err != nil {
		fmt.Printf(ePrefix+
			"Error returned by dMgr.GetFileInfoPlus() Error='%v' \n", err.Error())
		return
	}

	if actualDirFileInfo.IsFInfoInitialized {
		fmt.Println("            File Info IsDir(): ", actualDirFileInfo.IsDir())
		fmt.Println("             File Info Name(): ", actualDirFileInfo.Name())
		fmt.Println("             File Info Size(): ", actualDirFileInfo.Size())
		fmt.Println("          File Info ModTime(): ", du.GetDateTimeYMDAbbrvDowNano(
			actualDirFileInfo.ModTime()))
		fmt.Println("             File Info Mode(): ", actualDirFileInfo.Mode())
		fmt.Println("          File Info     Sys(): ", actualDirFileInfo.Sys())
		if actualDirFileInfo.IsDirPathInitialized {
			fmt.Println("                   Dir path: ", actualDirFileInfo.DirPath())
		}
	} else {
		fmt.Println("File Info Data is NOT Initialized")
	}

	return
}
