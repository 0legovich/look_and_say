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
  i := 0
  addNextElement, result := elementGenerator()
  for i != count {
    addNextElement()
    i = i + 1
  }

  return result()
}

type computedElem struct {
  count int
  elem string
}

func elementGenerator() (func(), func() []string) {
  resultArray := make([]string, 1)
  resultArray[0] = "1"

  addNextElement := func() {
    var arrayOfElements []computedElem
    letter := resultArray[len(resultArray) - 1]
    arrayStr := strings.Split(letter, "")

    for index, elem := range arrayStr {
      if (index != 0 && elem == arrayStr[index - 1]) {
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
    resultArray = append(resultArray, strings.Join(interArray, ""))
  }

  result := func() []string {
    return resultArray
  }

  return addNextElement, result
}
