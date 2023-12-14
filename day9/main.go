package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var oasis_H, difference_L [][]int /* oasis histories */
var sum_N, sum_P int              /* sum of predicted previous and next values of oasis histories */

/* function to construct layers of differences */
func zero_T(oasis_L []int) {
	count_Z := 0
	diff := []int{}

	/* finding difference between two values in oasis history
	   and further proceeding until all differnces become 0 */
	for index := range oasis_L {
		if index != len(oasis_L)-1 {
			diff = append(diff, (oasis_L[index+1] - oasis_L[index]))
			if diff[len(diff)-1] == 0 {
				count_Z++
			}
		}
	}

	/* to store values of differences as layers */
	difference_L = append(difference_L, diff)

	/* until it reaches 0 difference layer it calls this function recursively */
	if count_Z == (len(diff)) {
		next_V()
	} else {
		zero_T(diff)
	}
}

/* finding previous value for each given oasis histories */
func prev_V() {
	value_P := 0
	temp := 0

	/* finding previous value by adding 0 to the last layer
	   and substracting first value of previous layer to first value of current layer
	   iterating it up to find previous value of given oasis history */
	difference_L[len(difference_L)-1] = append(difference_L[len(difference_L)-1], 0)

	for l := len(difference_L) - 1; l > 0; l-- {
		temp = difference_L[l-1][0] - difference_L[l][0]
		difference_L[l-1] = append([]int{temp}, difference_L[l-1]...)
	}

	value_P = (oasis_H[len(oasis_H)-1][0]) - (difference_L[0][0])
	oasis_H[len(oasis_H)-1] = append([]int{value_P}, oasis_H[len(oasis_H)-1]...)

	/* sum of previous values for each oasis histories */
	sum_P += value_P
}

func next_V() {
	value_N := 0
	difference_L[len(difference_L)-1] = append(difference_L[len(difference_L)-1], 0)

	/* finding next values it's pretty much the same as finding previous value
	   except it is done by adding previous layer element and current layer element */
	for l := (len(difference_L) - 1); l > 0; l-- {
		difference_L[l-1] = append(difference_L[l-1],
			((difference_L[l][len(difference_L[l])-1]) +
				(difference_L[l-1][len(difference_L[l-1])-1])))
	}

	value_N = ((difference_L[0][len(difference_L[0])-1]) +
		(oasis_H[len(oasis_H)-1][len(oasis_H[(len(oasis_H)-1)])-1]))

	oasis_H[len(oasis_H)-1] = append(oasis_H[len(oasis_H)-1], value_N)

	/* sum of next values for each oasis histories */
	sum_N += value_N

	/* calling to find previous value for given oasis history */
	prev_V()

	difference_L = [][]int{}
}

/* function to change string to integer */
func change_S(numstr string) int {
	num, _ := strconv.Atoi(numstr)
	return num
}

/* function to append each oasis histories into a slice */
func append_L(line string) {
	temp := strings.Split(line, " ")

	value_H := []int{}

	for _, ele := range temp {
		value_H = append(value_H, change_S(ele))
	}

	oasis_H = append(oasis_H, value_H)

	/* after adding each histories,
	   construction layers of differences which holds
	   differences between all consecutive values are stored */

	zero_T(oasis_H[len(oasis_H)-1])
}

/* __main__ */
func main() {
	/* opening file */
	file, err := os.Open("input")

	/* handling file error */
	if err != nil {
		fmt.Println("File error : ", err)
	}

	/* closing file */
	defer file.Close()

	/* buffer scanner */
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		append_L(scanner.Text())
	}

	/* displaying results */
	fmt.Println("Part One Result : ", sum_N)
	fmt.Println("Part Two Result : ", sum_P)
}
