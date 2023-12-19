package main

import (
	"bufio"
	"fmt"
    "strconv"
	"os"
)

const (
    BOX_N = 256         /* total boxes */
    MULTIPLY_H = 17     /* multiplied by : used in hash function */
    DIVIDE_H = 256      /* divided by : used in hash function */
)

/* box type a slice of maps { string : int } */
type box [](map[string]int)

/* global : hashresult of string & sum of hashresult */
var sum, hashResult, focusResult int
var boxList [BOX_N]box

/* hashing algorithm */
/* 
    first step : find ascii value of given character and add to hashresult
    second step : multiply itself with 17
    third step : find remainder for hashresult by dividing it with 256
    fourth step : add the hashresult to the sum
*/
func findHash(input byte) {
    hashResult += int(input)
    hashResult *= MULTIPLY_H
    hashResult %= DIVIDE_H
}

/* function to convert string type to integer */
func toInt(numstr string) int {
    num, err := strconv.Atoi(numstr)

    if err != nil {
        fmt.Println("String to Int conversion error : ", err)
    }
    return num
}

/* function to gather boxdetails and proceed with further operations */
func boxDetails(boxInput int, boxLabel string, lenseFocal byte) {
    /* based on lense operation - or = 
       deduction or appending or modifying will take place */
    if byte(boxLabel[(len(boxLabel)-1)]) == byte('-') {
        for _, labels := range boxList[boxInput] {
            if _, isLabelExists := labels[boxLabel[: (len(boxLabel)-1)]]; isLabelExists {
                delete(labels, boxLabel[: (len(boxLabel)-1)])
            }
        }   /* for deduction (-), if label is not present, nothing will happen */
    } else {
        for _, labels := range boxList[boxInput] {
            if _, isLabelExists := labels[boxLabel[: (len(boxLabel)-1)]]; isLabelExists {
                labels[boxLabel[: (len(boxLabel)-1)]] = toInt(string(lenseFocal))
                return
            }
        }

        /* if already present, label's focal is changed (modified)
           or label is appended to the box it belongs to */
        boxList[boxInput] = append(boxList[boxInput], map[string]int{boxLabel[: (len(boxLabel)-1)] : toInt(string(lenseFocal))}) 
    }

}

/* function to find total focus power
   focus power of a label is found by : ( not zero indexed )

    (box) * (position in the given box) * (focal length)

*/
func totalFocus() {
    for list := range boxList {
        pos := 0
        for _, ele := range boxList[list] {
            if len(ele) > 0{        /* empty maps not considered */
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

    /* variables to hold values needed for part two */
    lenseNext := false
    boxInput := 0
    boxLabel := ""

    for {
        byteInput, err := buf.ReadByte() 

        /* sends labels with lenseoperation and boxInput value from hash of labels */
        if lenseNext {
            if byteInput == byte(',') {
                boxDetails(boxInput, boxLabel, 0)
            } else {
                boxDetails(boxInput, boxLabel, byteInput)
            }
            lenseNext = false
        }

        if err != nil {
                break
        } else {
            /* discarding comma and newline characters */
            if byteInput == byte(',') || byteInput == byte('\n') {
                sum+= hashResult        /* part one : after encountering comma */
                boxLabel = ""           /* hash is added to sum */
                hashResult = 0
                continue
            } else {
                /* once lenseoperation is encountered hash value is stored 
                   to find the box it belongs to */
                if byteInput == byte('-') || byteInput == byte('=') {
                    boxInput = hashResult
                    lenseNext = true
                } 

                findHash(byteInput)
                boxLabel += string(byteInput) /* includes both label and lense operation */
            }
        }
    }

    totalFocus()    /* for part two */

    /* part one : result */
    fmt.Println("Part one (sum of hashes) : ", sum)

    /* part two : result */
    fmt.Println("Part two (sum of focal Lengths) : ", focusResult)

}
