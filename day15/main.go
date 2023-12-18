package main

import (
	"bufio"
	"fmt"
	"os"
)

/* global : hashresult of string & sum of hashresult */
var sum, hashResult int

/* hashing algorithm */
/* 
    first step : find ascii value of given character and add to hashresult
    second step : multiply itself with 17
    third step : find remainder for hashresult by dividing it with 256
    fourth step : add the hashresult to the sum
*/
func findHash(input byte) {
    hashResult += int(input)
    hashResult *= 17
    hashResult %= 256
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

    /* buffer to read each character from input file */
    buf := bufio.NewReader(file)

    for {
        if byteInput, err := buf.ReadByte(); err != nil {
                break
        } else {
            /* discarding comma and newline characters */
            if byteInput == byte(',') || byteInput == byte('\n') {
                sum+= hashResult
                hashResult = 0
                continue
            } else {
                findHash(byteInput)
            }
        }
    }

    /* part one : result */
    fmt.Println("Part one (sum of hashes) : ", sum)

}
