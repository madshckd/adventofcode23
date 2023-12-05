package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
)

var numbers_G, symbols_G [][][]int
var numbers_M = make(map[string]int)
var sum_PN int

func parse_PN() {
    line := 0 

    for line < len(numbers_G) {
        for num := 0; num < len(numbers_G[line]); num++ {
            if line > 0 && line != len(numbers_G) - 1{
                if check_P(line, num) || check_C(line, num) || check_N(line, num) {
                    sum_PN += (numbers_M[strconv.Itoa(line) + strconv.Itoa(numbers_G[line][num][0]) + strconv.Itoa(numbers_G[line][num][1])])
                }
            } else if line == len(numbers_G) - 1{
                if check_P(line, num) || check_C(line, num) {
                    sum_PN += (numbers_M[strconv.Itoa(line) + strconv.Itoa(numbers_G[line][num][0]) + strconv.Itoa(numbers_G[line][num][1])])
                }
            } else {
                if check_N(line, num) || check_C(line, num) {
                    sum_PN += (numbers_M[strconv.Itoa(line) + strconv.Itoa(numbers_G[line][num][0]) + strconv.Itoa(numbers_G[line][num][1])])
                }
            }

        }
        line++
    }
}

func check_P(line, num int) bool {
    prev_L := numbers_G[line][num][0] - 1
    prev_R := numbers_G[line][num][1] + 1

    for i := 0; i < len(symbols_G[line - 1]); i++ {
        for j := prev_L; j < prev_R; j++ {
            if symbols_G[line - 1][i][0] == j {
                return true
            }
        }
    }
    return false
}

func check_C(line, num int) bool {
    cur_L := numbers_G[line][num][0] - 1
    cur_R := numbers_G[line][num][1]

    for i := 0; i < len(symbols_G[line]); i++ {
        if cur_L == symbols_G[line][i][0] || cur_R == symbols_G[line][i][0] {
            return true
        }
    }
    return false
}

func check_N(line, num int) bool {
    next_L := numbers_G[line][num][0] - 1
    next_R := numbers_G[line][num][1] + 1

    for i := 0; i < len(symbols_G[line + 1]); i++ {
        for j := next_L; j < next_R; j++ {
            if symbols_G[line + 1][i][0] == j {
                return true
            }
        }
    }
    return false
}

func change_S(numstr string) int {
    num, _ := strconv.Atoi(numstr)
    return num
}

func main() {
    var line_C int

    file, err := os.Open("input")

    if err != nil {
        fmt.Println("File Loading Error : ", err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    numbers_regex := regexp.MustCompile(`\d+`)
    symbols_regex := regexp.MustCompile(`[^a-zA-z0-9_.]`)

    for scanner.Scan(){
        numbers_G = append(numbers_G, (numbers_regex.FindAllIndex([]byte(scanner.Text()), -1)))

        for i := 0; i < len(numbers_G[line_C]); i++ {
            numbers_M[strconv.Itoa(line_C) + strconv.Itoa(numbers_G[line_C][i][0]) + strconv.Itoa(numbers_G[line_C][i][1])] = 
                change_S(scanner.Text()[numbers_G[line_C][i][0] : numbers_G[line_C][i][1]])

        }

        symbols_G = append(symbols_G, (symbols_regex.FindAllIndex([]byte(scanner.Text()), -1)))
        line_C++
    }
    parse_PN()

    fmt.Println("Part One Resut : ",  sum_PN)
}
