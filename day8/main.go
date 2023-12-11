package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

const (
    node_S, node_F string = "AAA", "ZZZ"
)

var status_F bool
var instruction_S string
var nodes_N = make(map[string]([]string))

func travel_AN() {
}

/* function to iterate given direction instruction
   starting from AAA node to reach ZZZ */
func travel_N() {
    /* initiating steps counter and node */
    steps := 0
    node := node_S

    /* iterate over to reach ZZZ node */
    for !status_F {
        /* iterating instruction set */
        for dir := 0; dir < len(instruction_S); dir++ {
            if instruction_S[dir] == byte('L') {
                node = nodes_N[node][0]         /* left instruction */
                steps++
            } else {
                node = nodes_N[node][1]         /* right instruction */
                steps++
            }

            /* check if node is ZZZ */
            if node == node_F {
                status_F = true
                break
            }
        }

        /* again iterating instruction, if ZZZ is not reached */
        if status_F != true {
            continue
        }
    }

    /* part one result -> number of steps to reach ZZZ node */
    fmt.Println("Part One Result : ", steps)
}

/* function to trim string */
func trim_S(input, sep string) string {
    return strings.Trim(input, sep)
}

/* function to construct nodes map */
func construct_N(line string) {
    temp := strings.Split(line, "=")
    key := trim_S(temp[0], " ")

    /* constructing map with coordinate values */
    for _, ele := range (strings.Split(temp[1], ",")) {
        nodes_N[key] = append(nodes_N[key], (trim_S(trim_S(trim_S(ele, " "), "("), ")")))
    }
}

/* __main__ */
func main() {
    var nodes_M bool     /* indicates node map inputs from file */

    /* opening input file */
    file, err := os.Open("input")

    /* handling file reading error */
    if err != nil {
        fmt.Println("File error : ", err)
    }

    /* closing file */
    defer file.Close()

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iterating over contents of input file */
    for scanner.Scan() {
        /* after encountering an empty line
           begin constructing nodes map */
        if scanner.Text() == "" {
            nodes_M = true
            continue
        }

        if nodes_M {
            construct_N(scanner.Text())
        } else {
            instruction_S = scanner.Text()
        }
    }

    /* part one : travel from AAA node to ZZZ node */
    travel_N()

    /* part two : travel from all nodes that ends with A */
    travel_AN()
}

