package main

import (
  pf "../pathfileops"
  "fmt"
)

/*

import (
  "MikeAustin71/pathfilego/003_filehelper/common"
  "fmt"
  "time"
)

*/

func main() {

  mainTest62CopySubDirs()

}

func mainTest62CopySubDirs() {
  srcDir := "D:\\T03\\WebSite_15"

  srcDMgr, err := pf.DirMgr{}.New(srcDir)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(srcDir).\n"+
      "srcDir='%v'\nError='%v'\n", srcDir, err.Error())
    return
  }

  targetDir := "D:\\T04"

  targetDMgr, err := pf.DirMgr{}.New(targetDir)

  if err != nil {
    fmt.Printf("Error returned by pf.DirMgr{}.New(targetDir).\n"+
      "targetDir='%v'\nError='%v'\n", targetDir, err.Error())
    return
  }

  fsc := pf.FileSelectionCriteria{}

  errs := srcDMgr.CopySubDirectoryTree(targetDMgr, false, fsc)

  if len(errs) > 0 {
    fmt.Printf("Errors returned by srcDMgr.CopySubDirectoryTree(targetDMgr,true, fsc).\n"+
      "targetDir='%v'\nErrors:\n\n", targetDMgr.GetAbsolutePath())

    for i := 0; i < len(errs); i++ {
      fmt.Printf("%v\n\n", errs[i].Error())
    }

    fmt.Printf("\n\n")

    return
  }

  fmt.Println("               mainTest62CopySubDirs                    ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}
