package common

import (
	"fmt"
	"time"
)

func WalkDirFindFiles(dMgr DirMgr, filePattern string, filesOlderThan, filesNewerThan time.Time ) error {

	du := DateTimeUtility{}

	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{filePattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		return fmt.Errorf("Error returned from dMgr.FindWalkDirFiles(filePattern, filesOlderThan, filesNewerThan) filePattern='%v'  filesOlderThan='%v' filesNewerThan='%v' Error='%v' \n",filePattern, du.GetDateTimeYMDAbbrvDowNano(filesOlderThan), du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
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
		for i:=0; i < dWalkInfo.FoundFiles.GetArrayLength(); i ++ {
			fmt.Printf("  Name: %v Mod Date: %v Path: %v \n", dWalkInfo.FoundFiles.FMgrs[i].FileNameExt,dWalkInfo.FoundFiles.FMgrs[i].ActualFileInfo.ModTime(), dWalkInfo.FoundFiles.FMgrs[i].ActualFileInfo.DirPath())
		}
		fmt.Println()
	}

	if dWalkInfo.Directories.GetArrayLength() > 0 {
		fmt.Println("DirMgrs Found:")

		for k:=0; k < dWalkInfo.Directories.GetArrayLength(); k++ {
			fmt.Printf("Dir: %v \n", dWalkInfo.Directories.DirMgrs[k].Path)
		}

	}

	if len(dWalkInfo.ErrReturns) > 0 {
		fmt.Println("Errors Found: ")

		for j:=0; j < len(dWalkInfo.ErrReturns); j++ {
			fmt.Printf("Error: %v \n", dWalkInfo.ErrReturns[j])
		}
		fmt.Println("")
	}

	return nil
}

func WalkDirFindFiles2(dMgr DirMgr, filePattern string, filesOlderThan, filesNewerThan time.Time ) error {

	du := DateTimeUtility{}
	fsc := FileSelectionCriteria{}

	fsc.FileNamePatterns = []string{filePattern}
	fsc.FilesOlderThan = filesOlderThan
	fsc.FilesNewerThan = filesNewerThan
	fsc.SelectCriterionMode = ANDFILESELECTCRITERION

	dWalkInfo, err := dMgr.FindWalkDirFiles(fsc)

	if err != nil {
		return fmt.Errorf("Error returned from dMgr.FindWalkDirFiles(filePattern, filesOlderThan, filesNewerThan) filePattern='%v'  filesOlderThan='%v' filesNewerThan='%v' Error='%v' \n",filePattern, du.GetDateTimeYMDAbbrvDowNano(filesOlderThan), du.GetDateTimeYMDAbbrvDowNano(filesNewerThan), err.Error())
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
		for i:=0; i < dWalkInfo.FoundFiles.GetArrayLength(); i ++ {
			fmt.Printf("  Name: %v SysInfo: %v \n", dWalkInfo.FoundFiles.FMgrs[i].FileNameExt, dWalkInfo.FoundFiles.FMgrs[i].ActualFileInfo.Sys())
		}
		fmt.Println()
	}

	if dWalkInfo.Directories.GetArrayLength() > 0 {
		fmt.Println("DirMgrs Found:")

		for k:=0; k < dWalkInfo.Directories.GetArrayLength(); k++ {
			fmt.Printf("Dir: %v \n", dWalkInfo.Directories.DirMgrs[k].Path)
		}

	}

	if len(dWalkInfo.ErrReturns) > 0 {
		fmt.Println("Errors Found: ")

		for j:=0; j < len(dWalkInfo.ErrReturns); j++ {
			fmt.Printf("Error: %v \n", dWalkInfo.ErrReturns[j])
		}
		fmt.Println("")
	}

	return nil
}

func PrintDirMgrFields(dMgr DirMgr) {
	du := DateTimeUtility{}
	fmt.Println("-----------------------------------------")
	fmt.Println(" 	DirMgr Fields")
	fmt.Println("-----------------------------------------")

	fmt.Println("                IsInitialized: ", dMgr.IsInitialized)
	fmt.Println("                Original Path: ", dMgr.OriginalPath)
	fmt.Println("                         Path: ", dMgr.Path)
	fmt.Println("             PathIsPopuslated: ", dMgr.PathIsPopulated)
	fmt.Println("                PathDoesExist: ", dMgr.PathDoesExist)
	fmt.Println("                   ParentPath: ", dMgr.ParentPath)
	fmt.Println("        ParentPathIsPopulated: ", dMgr.ParentPathIsPopulated)
	fmt.Println("                 RelativePath: ", dMgr.RelativePath)
	fmt.Println("      RelativePathIsPopulated: ", dMgr.RelativePathIsPopulated)
	fmt.Println("                 AbsolutePath: ", dMgr.AbsolutePath)
	fmt.Println("      AbsolutePathIsPopulated: ", dMgr.AbsolutePathIsPopulated)
	fmt.Println("AbsolutePathDifferentFromPath: ", dMgr.AbsolutePathDifferentFromPath)
	fmt.Println("        AbsolutePathDoesExist: ", dMgr.AbsolutePathDoesExist)
	fmt.Println("               Directory Name: ", dMgr.DirectoryName)
	fmt.Println("                   VolumeName: ", dMgr.VolumeName)
	fmt.Println("            VolumeIsPopulated: ", dMgr.VolumeIsPopulated)
	fmt.Println("============== File Info Data ============")

	if dMgr.ActualDirFileInfo.IsFInfoInitialized {
		fmt.Println("            File Info IsDir(): ", dMgr.ActualDirFileInfo.IsDir())
		fmt.Println("             File Info Name(): ", dMgr.ActualDirFileInfo.Name())
		fmt.Println("             File Info Size(): ", dMgr.ActualDirFileInfo.Size())
		fmt.Println("          File Info ModTime(): ", du.GetDateTimeYMDAbbrvDowNano(dMgr.ActualDirFileInfo.ModTime()))
		fmt.Println("             File Info Mode(): ", dMgr.ActualDirFileInfo.Mode())
		fmt.Println("          File Info     Sys(): ", dMgr.ActualDirFileInfo.Sys())
		if dMgr.ActualDirFileInfo.IsDirPathInitialized {
			fmt.Println("                   Dir Path: ", dMgr.ActualDirFileInfo.DirPath())
		}
	} else {
		fmt.Println("File Info Data is NOT Initialized")
	}


	return
}