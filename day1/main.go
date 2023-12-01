package main 

/*
this code of course provides a valid output
but, the structure that is linear and redundant approach of this code is not so good.
maybe it pretty much argues about my lack of logical thinking
one thing to be feel good about is that no regular expression or third party packages were used.
*/
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

//container to hold integer value
var container_N []int

//map of string and digits (don't know why I used it)
var digits_M = map[string]int{
    "zero"  : 0,
    "one"   : 1,
    "two"   : 2,
    "three" : 3,
    "four"  : 4,
    "five"  : 5,
    "six"   : 6,
    "seven" : 7,
    "eight" : 8,
    "nine"  : 9,
}

//function to convert string type to integer
func check_I(ch string) int {
    conv_N, _ := strconv.Atoi(ch)     /* returning converted value that is int for number, */
    return conv_N                     /* if string it returns 0 as result */
}

//function to pase line 
func parse_L(line string) string {

    container_N = []int{}           /* resetting container_N for every line */

    //iterating each character of a line to gather only numbers
    for pos := 0; pos < len(line); pos++{
        conv_N := check_I(string(line[pos]))
        if conv_N > 0 {
            container_N = append(container_N, conv_N)
        } 
    }

    //returning first and last digit in a line
    return strconv.Itoa(container_N[0])+strconv.Itoa(container_N[len(container_N) - 1])
}

//function to compute part two
/*
after this line, this function will lead to several lines 
not a good way to pack all the computations into single function
basically if a number is presented in the input it just adds to a slice from which two digits are gathered
for text it searches for strings that represents numbers like one, two, three
also it searches for combination of strings like twone which represents two digits which are two and one
*/
func parse_L2(line string) string {

    container_N = []int{}

    pos := 0

    for pos < len(line) {
        conv_N := check_I(string(line[pos]))
        if conv_N > 0 {
            container_N = append(container_N, conv_N)
            pos++
        } else {
            if line[pos] == 'o' {
                if pos + len("one") <= len(line){
                    if line[ pos : pos + len("one")] == "one" {
                        if pos + len("oneight") < len(line) {
                            if line[ pos : pos + len("oneight")] == "oneight" {
                                container_N = append(container_N, digits_M["one"], digits_M["eight"])
                                pos += len("oneight")
                                continue
                            }
                        }
                        container_N = append(container_N, digits_M["one"])
                        pos += len("one")
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 't' {
                if pos + len("two") <= len(line) {
                    if line[ pos : pos + len("two")] == "two" {
                        if pos + len("twone") < len(line) {
                            if line[ pos : pos + len("twone")] == "twone" {
                                container_N = append(container_N, digits_M["two"], digits_M["one"])
                                pos += len("twone")
                                continue
                            }
                        }
                        container_N = append(container_N, digits_M["two"])
                        pos += len("two")
                    } else if pos + len("three") <= len(line) {
                        if line[pos : pos + len("three")] == "three" {
                            if pos + len("threeight") < len(line) {
                                if line[ pos : pos + len("threeight")] == "threeright" {
                                    container_N = append(container_N, digits_M["three"], digits_M["eight"])
                                    pos += len("threeight")
                                    continue
                                }
                            }
                            container_N = append(container_N, digits_M["three"])
                            pos += len("three")
                        } else {
                            pos++
                            continue
                        }
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 'f' {
                if pos + len("four") <= len(line) {
                    if line[ pos : pos + len("four")] == "four" {
                        container_N = append(container_N, digits_M["four"])
                        pos += len("four")
                    } else if line [ pos : pos + len("five")] == "five" {
                        if pos + len("fiveight") < len(line) {
                            if line[ pos : pos + len("fiveight")] == "fiveight" {
                                container_N = append(container_N, digits_M["five"], digits_M["eight"])
                                pos += len("fiveight")
                                continue
                            }
                        }
                        container_N = append(container_N, digits_M["five"])
                        pos += len("five")
                    } else {
                    pos++
                    continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 's' {
                if pos + len("six") <= len(line)  {
                    if line[ pos : pos + len("six")] == "six" {
                        container_N = append(container_N, digits_M["six"])
                        pos += len("six")
                    } else if pos + len("seven") <= len(line) {
                        if line [ pos : pos + len("seven")] == "seven" {
                            if pos + len("sevenine") <= len(line) {
                                if line[pos : pos + len("sevenine")] == "sevenine" {
                                    container_N = append(container_N, digits_M["seven"], digits_M["nine"])
                                    pos += len("sevenine")
                                    continue
                                }
                            }
                            container_N = append(container_N, digits_M["seven"])
                            pos += len("seven")
                        } else {
                            pos++
                            continue
                        }
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 'e' {
                if pos + len("eight") <= len(line) {
                    if line[ pos : pos + len("eight")] == "eight" {
                        if pos + len("eightwo") < len(line) {
                            if line[ pos : pos + len("eightwo")] == "eightwo" {
                                container_N = append(container_N, digits_M["eight"], digits_M["two"])
                                pos += len("eightwo")
                                continue
                            }
                        }
                        if pos + len("eighthree") < len(line) {
                            if line[pos : pos + len("eighthree")] == "eighthree" {
                                container_N = append(container_N, digits_M["eight"], digits_M["three"])
                                pos += len("eighthree")
                                continue
                            }
                        }
                        container_N = append(container_N, digits_M["eight"])
                        pos += len("eight")
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 'n' {
                if pos + len("nine") <= len(line) {
                    if line[ pos : pos + len("nine")] == "nine" {
                        if pos + len("nineight") < len(line) {
                            if line[ pos : pos + len("nineight")] == "nineight" {
                                container_N = append(container_N, digits_M["nine"], digits_M["eight"])
                                pos += len("nineight")
                                continue
                            }
                        }
                        container_N = append(container_N, digits_M["nine"])
                        pos += len("nine")
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else if line[pos] == 'z' {
                if pos + len("zero") <= len(line) {
                    if line[ pos : pos + len("zero")] == "zero" {
                        container_N = append(container_N, digits_M["zero"])
                        pos += len("zero")
                    } else {
                        pos++
                        continue
                    }
                } else {
                    pos++
                    continue
                }
            } else {
                pos++
                continue
            }
        }
    }

    //returing double digit string 
    return strconv.Itoa(container_N[0])+strconv.Itoa(container_N[len(container_N) - 1])
}

//main
func main() {
    //result -> sum of two digit numbers in each line
    var sum1, sum2 int64

    //opening file
    file, err := os.Open("input")

    //for error, display them
    if err != nil {
        fmt.Println(err)
    }

    //closing file
    defer file.Close()

    //scanner to read file
    scanner := bufio.NewScanner(file)

    //reading input file -> line by line
    for scanner.Scan() {
        //calling function to parse line
        digit_T1, _ := strconv.Atoi(parse_L(scanner.Text()))
        digit_T2, _ := strconv.Atoi(parse_L2(scanner.Text()))
        sum1 += int64(digit_T1)
        sum2 += int64(digit_T2)
    }

    //displaying part one result
    fmt.Println("Part one (sum of two digit numbers) := ", sum1)
    fmt.Println("Part two (sum of two digit also considering words) := ", sum2)
}
