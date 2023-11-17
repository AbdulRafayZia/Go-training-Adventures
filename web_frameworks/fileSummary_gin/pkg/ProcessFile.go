package pkg

import (
    "fmt"
    "io/ioutil"
    "log"
    "main.go/summary"
)
func ProcessFile(filePath string, routines int) {
    channal := make(chan summary.Summary)
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fileData := string(content)
    chunk := len(fileData) / routines
    startIndex := 0
    endIndex := chunk
    for iterations := 0; iterations < routines; iterations++ {
        go summary.Counts(fileData[startIndex:endIndex], channal)
        startIndex = endIndex
        endIndex += chunk
    }
    for iterations := 0; iterations < routines; iterations++ {
        counts := <-channal
        fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
        fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
        fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
        fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
    }
}