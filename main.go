package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
  "log"
)

func main() {
  args := os.Args[1:]

  if (len(args) == 0) {
    log.Fatal("Argument error, missing argument number.\nExample: go run main.go 10")
  }

  count, err := strconv.Atoi(args[0])
  if (err != nil) {
    log.Fatal("Argument error, argument number should be integer.\nExample: go run main.go 10")
  }

  sequence := makeSeq(count)
  fmt.Println(strings.Join(sequence, "\n"))
}

func makeSeq(count int) []string {
  arr := []string{"1"}

  i := 0
  for i != count {
    letter := arr[len(arr) - 1]
    arrayLetters := strings.Split(letter, "")

    nextNumber := makeNextSeqNumber(arrayLetters)
    arr = append(arr, nextNumber)

    i = i + 1
  }

  return arr
}

type computedElem struct {
  count int
  elem string
}

// ['1', '1', '1', '2', '2'] => '3122'
func makeNextSeqNumber(arrayStr []string) string {
  var arrayOfElements []computedElem

  for index, elem := range arrayStr {
    if (index == 0) {
      arrayOfElements = append(arrayOfElements, computedElem{1, elem})
      continue
    }

    if (elem == arrayStr[index - 1]) {
      arrayOfElements[len(arrayOfElements) - 1].count++
    } else {
      arrayOfElements = append(arrayOfElements, computedElem{1, elem})
    }
  }

  interArray := []string{}
  for _, computedElem := range arrayOfElements {
    countStr := strconv.Itoa(computedElem.count)
    strArray := []string{countStr, computedElem.elem}
    interArray = append(interArray, strings.Join(strArray, ""))
  }

  return strings.Join(interArray, "")
}

