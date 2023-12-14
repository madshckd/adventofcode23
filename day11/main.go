package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var galaxyL [][]uint            /* slice to store positions of each galaxies */
var empty_R, empty_C []int      /* slice to store positions of empty rows and columns */

func calculateLength() {
    /* to hold differences in x-axis, y-axis and again sum of lengths */
    var diff_Y, diff_X, sumL uint

    /* iterating over all pairs : for galaxies that are found */
    for galaxy := 0; galaxy < len(galaxyL); galaxy++ {
        for Ogalaxy := galaxy + 1; Ogalaxy < len(galaxyL); Ogalaxy++ {
            /* differnce in x axis between the given pair of galaxy */
            diff_X = ((galaxyL[Ogalaxy][0]) - (galaxyL[galaxy][0]) )

            /* for y axis, theres a chance that preceding element in the given pair
               to have greater values, so we compare them */
            if galaxyL[Ogalaxy][1] > galaxyL[galaxy][1] {
                diff_Y = ((galaxyL[Ogalaxy][1]) - (galaxyL[galaxy][1]) )
            } else {
                diff_Y = ((galaxyL[galaxy][1]) - (galaxyL[Ogalaxy][1]) ) 
            }

            
            /* sum of lengths of all pairs */
            if diff_X == 0 {
                sumL += diff_Y
            } else {
                sumL += diff_X + diff_Y
            }
        }
    }

    /* part one : displaying sum of lengths of all pairs of galaxies */
    fmt.Println("Result (sum of shortest lengths of all pairs of galaxies : ", sumL)
}

/* function to expand empty rows based on exapansion rate*/
func expansion_R(rateX int) {
    for _, row_G := range galaxyL {
        expandR := 0
        for _, rowE := range empty_R {
            if row_G[0] > uint(rowE) {
                expandR++
            }
        }
        /* expanding rows based on rateX */
        row_G[0] += uint( (expandR * rateX) - expandR )
    }
}

/* fucntion to expand empty columns based on expansion rate */
func expansion_C(lineL, rateX int) {
    /* iterating galaxy column values to find empty columns */
    for col := 0; col < lineL; col++ {
        count := 0
        for _, col_G := range galaxyL {
            if col_G[1] == uint(col) {
                count++
            }
        }
        if count == 0 {
            empty_C = append(empty_C, col)
        }
    }

    for _, col_G := range galaxyL {
        expandC := 0
        for _, colE := range empty_C {
            if col_G[1] > uint(colE) {
                expandC++
            }
        }

        /* expanding columns based on rateX */
        col_G[1] += uint( (expandC * rateX) - expandC )
    }

}

/* __main__ */
func main() {

    var line uint           /* points current line of input */
    var lineL, rateX int    /* length of input line & expansion rate */

    /* user prompt */
    fmt.Print(`>>> Hello, enter the expansion rate for empty rows and columns
    For Part One, it is twice so 2
    For Part Two, it is a million so 1000000
    >>> `)
    fmt.Scanf("%d", &rateX)

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

    /* regular expression to indicate # character
       which refers to galaxy in the given input 
       and . character to find empty rows */
    regexp_G := regexp.MustCompile(`[#]{1}`)
    regexp_E := regexp.MustCompile(`[^#]`)

    /* iterating over input contents */
    for scanner.Scan() {
        lineL = len(scanner.Text())
        for _, ele := range (regexp_G.FindAllIndex([]byte(scanner.Text()), -1)) {
            /* finding each appearance of # character, storing it's position */
            galaxyL = append(galaxyL, []uint{line, uint(ele[0])})
        }
        if len(scanner.Text()) == len(regexp_E.FindAllIndex([]byte(scanner.Text()), -1)) {
            /* finding empty rows using regular expression
               line without # character */
            empty_R = append(empty_R, int(line))
        }
        line++
    }

    /* universe expansion where there is no galaxies */
    expansion_C(lineL, rateX)
    expansion_R(rateX)

    /* to calculate sum of lengths between all pairs of given galaxies */
    calculateLength()
}
