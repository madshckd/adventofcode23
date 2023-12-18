package main

import (
    "fmt"
    "os"
    "bufio"
)

var pattern [][]string
var notesSummary, pattern_C int

func reflectionLine() {
    possibleRow := 0
    for row := 1; row < len(pattern) - 2; row++ {
        if pattern[row][0] == pattern[row+1][0] {
            possibleRow = row+1
            if pattern[row-1][0] == pattern[row+2][0] {
                notesSummary += (100 * (row + 1))
                return 
            }
        }
    }
    notesSummary += vReflectionLine()
    fmt.Println("possible row : ", possibleRow)
    //fmt.Println(notesSummary)
}

func vReflectionLine() int {
    possibleCol := 0
    for pos := 1; pos < len(pattern[0][0]) - 2; pos++ {
        if iteRow(pos, 1) == len(pattern) {
            possibleCol = pos+1
            if iteRow(pos-1, 2) == len(pattern) {
                return pos+1
            }
        }
    }
    fmt.Println("Pattern and possibleCol: ", pattern_C, possibleCol)
    return 0
}

func iteRow(pos, inc int) int {
    matchCount := 0
    for rows := 0; rows < len(pattern); rows++ {
        if pattern[rows][0][pos] == pattern[rows][0][pos+inc] {
            matchCount++
        }
    }
    return matchCount
}


/* __main__ */
func main() {
    /* opening file */
    file, err := os.Open("input")

    /* handling file error */
    if err != nil {
        fmt.Println("File error : ", err)
        return
    }

    /* closing file */
    defer file.Close()

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iterating over input content */
    for scanner.Scan() {
        if scanner.Text() == "" {
            pattern_C++
            reflectionLine()
            pattern = [][]string{}
            continue
        } else {
            pattern = append(pattern, []string{scanner.Text()})
        }
    }
    pattern_C++
    reflectionLine()
    fmt.Println(notesSummary)
}
