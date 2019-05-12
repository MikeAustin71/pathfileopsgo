package appLibs

import (
  "errors"
  "strings"
  "unicode/utf8"
)

// StringUtility - encapsulates a collection of
// methods used to manipulate strings
type StringUtility struct {
  StrIn  string
  StrOut string
}

func (su StringUtility) ReadStrNewLineFromBuffer(buff []byte, previousPartialStr string, startIdx int) (outStr string, nextIdx int, isPartialStr bool) {
  sOut := make([]byte, 0, 512)
  lastIdx := len(buff) - 1
  lPrevPartStr := len(previousPartialStr)
  if lPrevPartStr > 0 {
    for i := 0; i < lPrevPartStr; i++ {
      sOut = append(sOut, previousPartialStr[i])
    }
  }

  for i := startIdx; i <= lastIdx; i++ {

    if buff[i] == '\n' {
      return string(sOut), i + 1, false
    }

    sOut = append(sOut, buff[i])

  }

  return string(sOut), -1, true

}

// StrCenterInStr - returns a string which includes
// a left pad blank string plus the original string.
// The complete string will effectively center the
// original string is a field of specified length.
func (su StringUtility) StrCenterInStr(strToCenter string, fieldLen int) (string, error) {

  pad, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

  if err != nil {
    return "", errors.New("StringUtility:StrCenterInStr() - " + err.Error())
  }

  return pad + strToCenter, nil

}

// StrPadLeftToCenter - Returns a blank string
// which allows centering of the target string
// in a fixed length field.
func (su StringUtility) StrPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

  sLen := su.StrGetRuneCnt(strToCenter)

  if sLen > fieldLen {
    return "", errors.New("StringUtility:StrPadLeftToCenter() - String To Center is longer than Field Length")
  }

  if sLen == fieldLen {
    return "", nil
  }

  margin := (fieldLen - sLen) / 2

  return strings.Repeat(" ", margin), nil
}

// StrGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (su StringUtility) StrGetRuneCnt(targetStr string) int {
  return utf8.RuneCountInString(targetStr)
}

// StrGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (su StringUtility) StrGetCharCnt(targetStr string) int {
  return len([]rune(targetStr))
}

// TrimEndMultiple- Performs the following operations on strings:
// 1. Trims Right and Left for all instances of 'trimChar'
// 2. Within the interior of a string, multiple instances
// 		of 'trimChar' are reduce to a single instance.
func (su StringUtility) TrimEndMultiple(targetStr string, trimChar rune) (rStr string, err error) {

  if targetStr == "" {
    err = errors.New("Empty targetStr")
    return
  }

  fStr := []rune(targetStr)
  lenTargetStr := len(fStr)
  outputStr := make([]rune, lenTargetStr)
  lenTargetStr--
  idx := lenTargetStr
  foundFirstChar := false

  for i := lenTargetStr; i >= 0; i-- {

    if !foundFirstChar && fStr[i] == trimChar {
      continue
    }

    if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
      continue
    }

    if i == 0 && fStr[i] == trimChar {
      continue
    }

    foundFirstChar = true
    outputStr[idx] = fStr[i]
    idx--
  }

  if idx != lenTargetStr {
    idx++
  }

  if outputStr[idx] == trimChar {
    idx++
  }

  result := string(outputStr[idx:])

  return result, nil

}

func (su StringUtility) SwapRune(currentStr string, oldRune rune, newRune rune) (string, error) {

  if currentStr == "" {
    return currentStr, nil
  }

  rStr := []rune(currentStr)

  lrStr := len(rStr)

  for i := 0; i < lrStr; i++ {
    if rStr[i] == oldRune {
      rStr[i] = newRune
    }
  }

  return string(rStr), nil
}
