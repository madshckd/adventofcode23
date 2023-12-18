package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "slices"
)

var cubes, rounded [][]int

func sumLoad() {
    totalLoad := 0
    toSouth := len(rounded)
    for row := 0; row < len(rounded) - 1; row++ {
        totalLoad += toSouth * (len(rounded[row]))
        toSouth--
    }
    fmt.Println(totalLoad)
}

func tilt() {
    for row, rowE := range rounded {
        if row > 0 {
            reachNorth(row, rowE)
        }
    }
    sumLoad()
}

func reachNorth(row int, rowE []int) {
    delEle := []int{}
    for pos := 0; pos < len(rowE); pos++ {
        prevRow := row - 1
        for {
            if slices.Contains(rounded[prevRow], rowE[pos]) ||
                slices.Contains(cubes[prevRow], rowE[pos]) {
                break
            }
            prevRow--
            if prevRow == -1 {
                break
            }
        }
    
        if !(slices.Contains(rounded[prevRow+1], rowE[pos])) {
            rounded[prevRow+1] = append(rounded[prevRow+1], rowE[pos])
            slices.Sort(rounded[prevRow+1])
            delEle = append(delEle, pos)
        }
    }
    deleteEle(row, delEle)

}

func deleteEle(row int, delEle []int) {
    for pos := len(delEle)-1; pos >= 0; pos-- {
        if delEle[pos] == 0 {
            rounded[row] = rounded[row][1:]
        } else if delEle[pos] == (len(rounded[row]) - 1) {
            rounded[row] = rounded[row][: (len(rounded[row]) - 1)]
        } else {
            rounded[row] = append(rounded[row][:(delEle[pos])], rounded[row][(delEle[pos]+1):]...)
        }
    }
}


func cleanIndex(input [][]int, isRound bool) {
    cleanIndex := []int{}
    for _, ele := range input {
        if ele[0]+1 == ele[1] {
            cleanIndex = append(cleanIndex, ele[0])
        } else {
            for i := ele[0]; i < ele[1]; i++ {
                cleanIndex = append(cleanIndex, i)
            }
        }
    }

    if isRound {
        rounded = append(rounded, cleanIndex)
    } else {
        cubes = append(cubes, cleanIndex)
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

    /* regular expression */
    roundedRocks := regexp.MustCompile(`[O]+`)
    cubeRocks := regexp.MustCompile(`[#]+`)

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iterating over input contents */
    for scanner.Scan() {
        cleanIndex(roundedRocks.FindAllIndex([]byte(scanner.Text()), -1), true)
        cleanIndex(cubeRocks.FindAllIndex([]byte(scanner.Text()), -1), false)
    }

    tilt()
}
