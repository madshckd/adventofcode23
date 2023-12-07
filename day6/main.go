package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

/* grid to hold time - distance values */
var grid_TD, grid_STD [][]int


/* function to count number of ways
   for each race and multiplying
   results of each race */
func count_W(input [][]int) int {
    product := 1
    for t := 0; t < len(input[0]); t++ {
        ways := 0
        /* button_H represents amount of time in ms */
        for button_H := 0; button_H <= input[0][t]; button_H++ {
            if ((input[0][t] - button_H) * button_H) > input[1][t] {
                /* with beating recorded distance -> ways gets incremented */
                ways++
            }
        }
        /* product of number of ways available for winning each races */
        product *= ways
    }
    return product
}

/* function to  convert string -> number */
func change_S(str string) int {
    num, _ := strconv.Atoi(str)
    return num
}

/* function to clean strings */
func clean_S(line string) {
    str := ""
    vals := []int{}

    /* striping and cleaning unwanted and to accept valid values */
    numstr := (strings.Split(strings.Trim(line, " "), " "))

    for pos := 0; pos < len(numstr); pos++ {
        num := change_S(numstr[pos])
        if num != 0 {
            vals = append(vals, num)
            str += strconv.Itoa(num)
        }
    }
    /* grid_TD holds normal time - distance values
       grid_STD holds concatenated or real values */
    grid_TD = append(grid_TD, vals)
    grid_STD = append(grid_STD, []int{change_S(str)})
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
    for scanner.Scan() {
        line  := strings.Split(scanner.Text(), ":")
        clean_S(line[1])
    }

    /* finding out different ways to win each races */
    /* displaying results */
    fmt.Println("Part One Result (product) : ", count_W(grid_TD))
    fmt.Println("Part Two Result (product) : ", count_W(grid_STD))

}
