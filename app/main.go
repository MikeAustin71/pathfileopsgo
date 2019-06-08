package main

import (
  pf "../pathfileops"
  "fmt"
)

/*


import (
  pf "../pathfileops"
  "fmt"
)


*/

func main() {

  mainTest67AreFilesSame()

}

func mainTest67AreFilesSame() {
  // ..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir\\level_04_dir\\level_4_2_test.txt

  fh := pf.FileHelper{}

  /*
  rawFile1 := "..\\filesfortest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir" +
    "\\level_04_dir\\level_4_2_test.txt"
  */

  rawFile1 := "D:/gowork/src/MikeAustin71/pathfileopsgo/filesfortest/levelfilesfortest" +
    "/level_01_dir/level_02_dir/level_03_dir/level_04_dir/level_4_2_test.txt"

  correctedFile1 := fh.AdjustPathSlash(rawFile1)

  correctedFile2 := correctedFile1

  filesAreSame, err := fh.AreSameFile(correctedFile1, correctedFile2)

  if err != nil {
    fmt.Printf("Error returned by fh.AreSameFile(relFile1, relFile2). "+
      "relFile1='%v'\nrelFile2='%v'\nError='%v'",
      correctedFile1, correctedFile2, err.Error())
    return
  }

  if !filesAreSame {
    fmt.Println ("Error: Expected file comparison='true'. Instead, file comparison='false'.")
    return
  }

  fmt.Println("              mainTest67AreFilesSame                   ")
  fmt.Println("********************************************************")
  fmt.Println("                    SUCCESS!!!                          ")
  fmt.Println("********************************************************")

}
