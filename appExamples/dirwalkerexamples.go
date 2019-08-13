package appExamples

import (
  "fmt"
  "path/filepath"
)

type DirWalkerExamples struct {
  Input  string
  Output string
}

func (dWalkEx DirWalkerExamples) TestFilePathMatch002() {
  // filename := "D:\\Test03\\start.txt"
  // Note: This doesn't work. \\ generates
  // and invalid response from Match

  filename := "start.txt"

  pattern := "*.txt"

  matched, err := filepath.Match(pattern, filename)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(matched)

}

func (dWalkEx DirWalkerExamples) TestFilePathGlob003(dirPathAndPattern string) {

  matches, err := filepath.Glob(dirPathAndPattern)

  if err != nil {
    panic(err)
  }

  for i, f := range matches {
    fmt.Println(i, " - ", f)
  }

}
