package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cards_W, cards_P [][]int
var total_P int

/* function to convert string to integer */
func change_S(str string) int {
    num, _ := strconv.Atoi(str)
    return num
}

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
    points := 0
    line := (len(cards_W) - 1);
    for pos := 0; pos < len(cards_W[line]); pos++ {
        if found_C(line, pos) {
            if points > 0 {
                points *= 2
            } else {
                points++
            }
        }
    }

    total_P += points
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

    slices.Sort(temp)

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

    arrange_C((strings.Split(strings.Trim(cards[0], " "), " ")), true)
    arrange_C((strings.Split(strings.Trim(cards[1], " "), " ")), false)

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

    /*total points for part one */
    fmt.Println("Part One Result : ", total_P)
}
