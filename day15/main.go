package main

import (
	"bufio"
	"fmt"
    "strconv"
	"os"
)

type box [](map[string]int)

/* global : hashresult of string & sum of hashresult */
var sum, hashResult, focusResult int
var boxList [256]box

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

func toInt(numstr string) int {
    num, err := strconv.Atoi(numstr)

    if err != nil {
        fmt.Println("String to Int conversion error : ", err)
    }
    return num
}

func boxDetails(boxInput int, boxLabel string, lenseOperation, lenseFocal byte) {
    if lenseOperation == byte('-') {
        for _, labels := range boxList[boxInput] {
            if _, isLabelExists := labels[boxLabel[0:2]]; isLabelExists {
                delete(labels, boxLabel[0:2])
            }
        }
    } else {
        for _, labels := range boxList[boxInput] {
            if _, isLabelExists := labels[boxLabel[0:2]]; isLabelExists {
                labels[boxLabel[0:2]] = toInt(string(lenseFocal))
                return
            }
        }
        boxList[boxInput] = append(boxList[boxInput], map[string]int{boxLabel[0:2] : toInt(string(lenseFocal))}) 
    }
}

func totalFocus() {
    for list := range boxList {
        pos := 0
        for _, ele := range boxList[list] {
            if len(ele) > 0{
                fmt.Println(ele)
                for _, focus := range ele {
                    focusResult += ((pos+1) * focus * (list+1))
                    pos++
                }
            }
        }
    }
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

    lenseNext := false
    lenseOperation := byte('-')
    boxInput := 0
    boxLabel := ""
    for {
        byteInput, err := buf.ReadByte() 

        if lenseNext {
            if byteInput == byte(',') {
                boxDetails(boxInput, boxLabel, lenseOperation, 0)
            } else {
                boxDetails(boxInput, boxLabel, lenseOperation, byteInput)
            }
            lenseNext = false
        }

        if err != nil {
                break
        } else {
            /* discarding comma and newline characters */
            if byteInput == byte(',') || byteInput == byte('\n') {
                sum+= hashResult
                boxLabel = ""
                hashResult = 0
                continue
            } else {
                if byteInput == byte('-') {
                    boxInput = hashResult
                    lenseOperation = byte('-')
                    lenseNext = true
                    
                } else if byteInput == byte('=') {
                    boxInput = hashResult
                    lenseOperation = byte('=')
                    lenseNext = true
                } 
                findHash(byteInput)
                boxLabel += string(byteInput)
            }
        }
    }

    totalFocus()

    /* part one : result */
    fmt.Println("Part one (sum of hashes) : ", sum)

    /* part two : result */
    fmt.Println("Part two (sum of focal Lengths) : ", focusResult)

    for _, ele := range boxList {
        fmt.Println(ele)
    }


}
