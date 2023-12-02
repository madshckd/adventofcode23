package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var expected_S = [3]int64{12, 13, 14} /* expected values to find possible subset */
var bag_G [3]int64      /* to hold subset color value */
var sum_G int64         /* to hold sum of all gameid that are feasible */
var test_P int64        /* to hold possibility test value which should 
                        be equal to count(subsets) for true */

/* function to convert string to integer */
func toInt(numstr string) int {
    num, _ := strconv.Atoi(numstr)
    return num
}

/* function to parse subsets to find it's possibility */
func possible_S(subsets []string, gameid int) {

    test_P = 0

    /* iterates each subset for the given game id */
    for set := 0; set < len(subsets); set++ {
        
        /* empty bag and dividing subsets to entries to assign values */
        bag_G = [3]int64{}
        cubes := strings.Split(subsets[set], ",")

        /* iterates each colors in a subset */
        for col := 0; col < len(cubes); col++ {
            cube := strings.Split(cubes[col], " ")
        
            /* assigning values based on color given */
            if cube[2] == "red" {
                bag_G[0] = int64(toInt(cube[1]))
            } else if cube [2] == "green" {
                bag_G[1] = int64(toInt(cube[1]))
            } else {
                bag_G[2] = int64(toInt(cube[1]))
            }
        }

        /* checks a subset of a game to test it's possibility to hold expected value */
        if bag_G[0] > expected_S[0] ||
            bag_G[1] > expected_S[1] ||
            bag_G[2] > expected_S[2] {
                continue
            } else {
                test_P++
            }
    }

    /* only if all subsets or bags have possibility to hold expected values,
    game id is added to the sum */
    if test_P == int64(len(subsets)) {
        sum_G += int64(gameid)
    }
}

/* function to parse input values of each game */
func parse_G(line string) {
    values := strings.Split(line, ":")      /* splitting game id and it's subsets */
    gameid := strings.Split(values[0], " ")

    /* passing game id and it's subsets */
    possible_S(strings.Split(values[1], ";"), toInt(gameid[1]))
}

/* __main__ */
func main() {
    /* opening file */
    file, err := os.Open("input")

    /* file opening error, if any, handled */
    if err != nil {
        fmt.Println("error opening file :: ", err)
    }

    defer file.Close()                  /* closing file after reading */

    scanner := bufio.NewScanner(file)   /* scanner to read line by line */

    for scanner.Scan() {
        parse_G(scanner.Text())
    }

    fmt.Println("Part one (result) sum of possible game ids : ", sum_G)
}
