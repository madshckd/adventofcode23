package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var row, col, rowMax int
var outline [][]int

func toInt(numstr string) int {
    num, err := strconv.Atoi(numstr)

    if err != nil {
        fmt.Println("Integer conversion : ", err)
    }

    return num
}


func dig(direction string, steps int) {
    irow := row
    switch direction {
    case "L" :
        col -= steps
        break
    case "R" :
        col += steps
        break
    case "U" :
        i := row-1
        row -= steps
        for i > row {
            outline = append(outline, []int{i, col})
            i--
        }
        break
    case "D" :
        j := row+1
        row += steps
        for j < row {
            outline = append(outline, []int{j, col})
            j++
        }
        break
    }

    if irow > rowMax {
        rowMax = irow
    }

    outline = append(outline, []int{row, col})

}

func digInterior() {
    var lavaVolume int
    inline := make(map[int]([]int))

    for _, ele := range outline {
        inline[ele[0]] = append(inline[ele[0]], ele[1])
    }

    for _, ele := range inline {
        slices.Sort(ele)
        lavaVolume += ( ele[ (len(ele) - 1) ] - ele[0] ) + 1
    }
    fmt.Println("Part one (result) : ", lavaVolume)
}

func parseInput(diggerInput string) {
    digInfo := strings.Split(diggerInput, " ")
    dig(digInfo[0], toInt(digInfo[1]))
}

/* __main__ */
func main() {
    file, err := os.Open("sample")

    if err != nil {
        fmt.Println("File Operation Error : ", err)
        return
    }

    buf := bufio.NewScanner(file)

    for buf.Scan() {
        parseInput(strings.Trim((strings.Split(buf.Text(), "("))[0], " "))
    }

    digInterior()
}
