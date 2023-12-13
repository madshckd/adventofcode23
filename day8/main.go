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

/* to store instruction set */
var instruction_S string
/* node map container node = (L_node, R_node) */
var nodes_N = make(map[string]([]string))

/* function to start with nodes that ends with A like __A 
   and iterate it simultaneously to reach end nodes with __Z
   not working out with this implementation
   concurrent might be the right thing to do */
func travel_AN() {
    status_F := false
    steps := 0
    nodes := ends_A()
    count_Z := len(nodes)

    /* iterating till Z nodes are reached */
    for !status_F {
        for dir := 0; dir < len(instruction_S); dir++ {
            z := 0
            temp_N := []string{}
            /* starting with nodes that ends with A */
            for _, ele_N := range nodes {
                if instruction_S[dir] == byte('L') {
                    temp_N = append(temp_N, nodes_N[ele_N][0])
                } else {
                    temp_N = append(temp_N, nodes_N[ele_N][1])
                }
            }
            /* reaching next stop as per instruction to 
               increment the steps value */
            steps++

            /* checking if we reached the destination 
               that is nodes ending with Z */
            for _, ele_N := range temp_N {
                if ele_N[2] == byte('Z') {
                    z++
                }
            }

            if z == count_Z {
                status_F = true
                break
            } else {
                nodes = temp_N
            }
        }

        /* iterating over instruction set if destination is not reached */
        if status_F != true {
            continue
        }
    }
}

/* function to find nodes that ends with __A */
func ends_A() []string {
    nodes_A := []string{}
    for key_N := range nodes_N {
        if key_N[2] == byte('A') {
            nodes_A = append(nodes_A, key_N)
        }
    }
    return nodes_A
}

/* function to iterate given direction instruction
   starting from AAA node to reach ZZZ */
func travel_N() {
    /* initiating steps counter, status_F and node */
    steps := 0
    status_F := false
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
    //travel_AN()
}

