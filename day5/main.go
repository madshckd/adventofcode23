package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seeds []int     /* seed numbers */
var location_S int  /* smallest location value */

/* map holding list of values of tags mentioned below */
var maps_L = make(map[int]([][]int))

/* tags, used to remove these texts from input file */
var tags = []string{"seed-to-soil map:",
                     "soil-to-fertilizer map:",
                     "fertilizer-to-water map:",
                     "water-to-light map:",
                     "light-to-temperature map:",
                     "temperature-to-humidity map:",
                     "humidity-to-location map:" }

/* function to match seeds with respective required soil, light so on */
func interpret_S(status_E bool) int {
    location_S = 0 

    /* for extended seed number inputs */
    if status_E {
        instance := 0/* to indicate first instance to initiate location_S */

        /* first value is source and following 
           value is range to extend seed input */
        for s := 0; s < len(seeds); s++ {
            if s % 2 == 0 {
                for r := seeds[s]; r < (seeds[s] + seeds[s+1]); r++ {
                    iterate_M(r, instance)
                    instance++
                }
            } else {
                continue
            }
        }
    } else {
        /* for non-expanding seed number inputs */
        for s  := 0 ; s < len(seeds); s++ {
            iterate_M(seeds[s], s)
        }
    }

    /* returning location value */
    return location_S
}

/* iterating map_L to that contains input values */
func iterate_M(key, s int) {
    loc := 0
    
    /* iterating map to find matches
       seed -> soil -> .. -> location */
    for m := 1; m <= len(maps_L); m++ {
        temp := matches_L(key, maps_L[m])
        key = temp

        if m == len(tags) {
            loc = key
        }
    }

    /* after matching location value 
       figuring out if it's the smallest */
    if s == 0 {
        location_S = loc
    } else {
        if loc < location_S {
            location_S = loc
        }
    }
}

/* matches seed->soil->fertilizer->..->location */
func matches_L(key int, list [][]int) int {
    for l := 0; l < len(list); l++ {
        if key >= list[l][1] && key <= (list[l][1] + (list[l][2] - 1)) {
            return (key - list[l][1]) + list[l][0]
        }
    }
    return key
}

/* function to convert string to integer slice */
func change_S(numstr string) int {
    num, _ := strconv.Atoi(numstr)
    return num
}

/* function to clean the input string */
func clean_S(dest int, line string) {
    /* refers to seed inputs */
    if dest == 0 {
        numstr := strings.Split(line, ":")
        val := strings.Split(strings.Trim(numstr[1], " "), " ")
        for pos := 0; pos < len(val); pos++ {
            seeds = append(seeds, change_S(val[pos]))
        }
    } else { /* refers to inputs other than seed numbers */
        temp := []int{}
        val := strings.Split(strings.Trim(line, " "), " ")
        for pos := 0; pos < len(val); pos++ {
            temp = append(temp, change_S(val[pos]))
        }
        /* storing all information in a map 
           refer tags []string for information 
           on headings */
        maps_L[dest] = append(maps_L[dest], temp)
        /* dest depends on tags order < refer above > */
    }

}

/* __main__ */
func main() {
    /* opening file */
    file, err := os.Open("input")

    /* handling file error */
    if err != nil {
        fmt.Println("File Error : ", err)
        return
    }

    /* closing file */
    defer file.Close()

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iterating over contents of file */
    /* both tags used to differentiate seed 
       numbers with other inputs (also with layers) */
    tag_C, tag_D := 0, 0
    for scanner.Scan() {
        /* ignoring newline and also heading texts
            with pre-defined texts -> tags []string */
        if scanner.Text() == tags[tag_C]{
            tag_D++
            if tag_C < len(tags) - 1 {
                tag_C++
            }
            continue
        } else if scanner.Text() == "" {
            continue
        } else {
            /* only accepting valid inputs */
            clean_S(tag_D, scanner.Text())
        }
    }

    /* displaying results */
    /* for both normal seed input and also expanded one */
    fmt.Println("Part one result (smallest location) : ", interpret_S(false))
    fmt.Println("Part two result (smallest location) expanded edition : ", interpret_S(true))

}
