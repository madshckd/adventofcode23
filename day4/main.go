package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/* global variables 

cards_W => collection of winning cards
cards_P => collection of playable cards
total_P => total points (part one result)
total_I => total instances (part two result)
points_C => holds points for each cards

*/

var cards_W, cards_P [][]int
var total_P, total_I int
var points_C []int

/* function to convert string to integer */
func change_S(str string) int {
    num, _ := strconv.Atoi(str)
    return num
}

/* 
function to find total instances
based on matches found (winning cards)
*/
    
func instances_A() {
    instances_C := make([]int, len(cards_W))

    /* iterating through points from each cards
        to add instances of card copies in 
        following cards */

    for pos := 0; pos < len(instances_C); pos++ {
        instances_C[pos]++
        for iter := 0; iter < instances_C[pos]; iter++ {
            for ele := pos + 1; ele <= (pos + (points_C[pos])); ele++ {
                instances_C[ele]++
            }
        }

        /* displaying part two result */
        total_I += instances_C[pos]
    }
}

/* function to found match in playable card against winning cards
    binary search implementation */
func found_C(line, pos int) bool {
    target := cards_W[line][pos]

    low := 0
    high := len(cards_P[line]) - 1

    for low <= high {
        mid := (high + low) / 2
        if target < cards_P[line][mid] {
            high = mid - 1
        } else if target > cards_P[line][mid] {
            low = mid + 1
        } else {
            return true
        }
    }
    return false
}

/* function to iterate winning cards over playable cards
    to find points */
func points_W() {
    matches_C := 0
    points := 0
    /* finding matches to calculate points */
    line := (len(cards_W) - 1);
    for pos := 0; pos < len(cards_W[line]); pos++ {
        if found_C(line, pos) {
            if points > 0 {
                points *= 2
            } else {
                points++
            }
            matches_C++
        }
    }

    total_P += points
    /* following is used for part two */
    /* to know winning matches to produce card copies */
    points_C = append(points_C, matches_C)
}

/* function to convert string slice to integer slice and sort them*/
func arrange_C(cards []string, side_D bool) {
    var temp = []int{}
    for pos := 0; pos < len(cards); pos ++ {
        val := change_S(cards[pos])
        if val == 0 {
            continue
        } else {
            temp = append(temp, val)
        }
    }

    /* sorting */
    slices.Sort(temp)

    /* default is winning card allocation for boolean value true */
    if side_D {
        cards_W = append(cards_W, temp)
    } else {
        cards_P = append(cards_P, temp)
    }
}

/* function to pasrse a iine */
func parse_L(line string) {
    temp := strings.Split(line, ":")
    cards := strings.Split(temp[1], "|")

    /* just stripping all other useless contents */

    arrange_C((strings.Split(strings.Trim(cards[0], " "), " ")), true)
    arrange_C((strings.Split(strings.Trim(cards[1], " "), " ")), false)

    /* calculate points for every matches of winning cards */
    points_W()
}

/* __main__ */
func main() {
    /* opening file */
    file, err := os.Open("input")

    /* handling error */
    if err != nil {
        fmt.Println("File Error : ", err)
        return
    }

    /* closing file */
    defer file.Close()

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iteration over input file */
    for scanner.Scan() {
        parse_L(scanner.Text())
    }

    /* for part two -> to calculate number of instances */
    instances_A()

    /* displaying results */
    fmt.Println("Part One Result : ", total_P)
    fmt.Println("Part Two Result : ", total_I)

}
