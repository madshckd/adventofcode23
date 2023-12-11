package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var status_J bool           /* applying joker card rule (part two) */
var hands_G [][]string      /* collection of all hands */ 
var bets_G [][]int          /* collection of bets for all hands */
var types_M = map[string]([]int){
    "five_K" : []int{},
    "four_K" : []int{},
    "full_H" : []int{},
    "three_K": []int{},
    "two_P"  : []int{},
    "one_P"  : []int{},
    "high_C" : []int{},
}

/* cards priority level only applicable to part one 
   for part two it's almost same except j which will
   be treated as weakest card */
var cards_P = map[byte]int {
    '2' : 1, '3' : 2, '4' : 3, '5' : 4, '6' : 5, '7' : 6, '8' : 7,
    '9' : 8, 'T' : 9, 'J' : 10, 'Q' : 11, 'K' : 12, 'A' : 13,
}

/* function to arrange hands in specfic types
   like arranging lowest to highest in a specific types */
/* just a linear sorting */
func arrange_H(types string, hands []int) {
    for l := 0; l < len(hands); l++{
        for r := l+1; r < len(hands); r++ {
            if sort_T(hands[l], hands[r]) {
                temp := types_M[types][l]
                types_M[types][l] = types_M[types][r]
                types_M[types][r] = temp
            }
        }
    }
}

/* function to sort to find lowest and highest hand to arrange them */
/* works differently for part two -> since J is the weakes */
func sort_T(hand_L, hand_R int) bool {
    for i := 0; i< len(hands_G[hand_L][0]); i++ {
        if cards_P[hands_G[hand_L][0][i]] > cards_P[hands_G[hand_R][0][i]] {
            if status_J {
                if hands_G[hand_L][0][i] == byte('J') {
                    return false
                }
            }
            return true
        } else if cards_P[hands_G[hand_L][0][i]] < cards_P[hands_G[hand_R][0][i]] {
            if status_J {
                if hands_G[hand_R][0][i] == byte('J') {
                    return true
                }
            }
            return false
        } else {
            continue
        }
    }
    return false
}

/* function to find total winning points, iterating from lowest hand types 
   and iterating it's elements and rank variable from lowest to highest 
   starting from 1 to n (total number of hands */
func total_W() {
    types := []string{
        "high_C", "one_P", "two_P", "three_K", "full_H", "four_K", "five_K",
    }

    total_P, rank := 0, 1
    for t := 0; t < len(types); t++ {
        for ele := 0; ele < len(types_M[types[t]]); ele++{
            total_P += bets_G[types_M[types[t]][ele]][0] * rank 
            rank++
        }
    }

    fmt.Println(total_P)
}

/* splitting each hands to know it's types */
/* for part two , joker rule is applicable
    makeing it as a powerful wild card */
func split_T() {
    cur_H := len(hands_G) - 1
    sum_M, matches_J := 0, 0 
    for  i, ele := range hands_G[cur_H][0]{
        if byte(ele) == byte('J') {
            matches_J++
        }
        matches := 0
        for j := i+1; j < len(hands_G[cur_H][0]); j++{
            if byte(ele) == hands_G[cur_H][0][j]{
                matches++
            }
        }
        sum_M += matches 
    }

    /* depending on status_J (joker rule application) */
    if status_J {
        append_S(Joker_R(sum_M, matches_J), cur_H)
    } else {
        append_S(sum_M, cur_H)
    }
}

/* function to find it's type applying joker rule for part two */
func Joker_R(sum_M, matches_J int) int {
    add_V := 0
    if matches_J > 0 {
        if sum_M == 1 {
            if matches_J == 1 || matches_J == 2{
                add_V = 2
            } 
        } else if sum_M == 2 {
            if matches_J == 1 {
                add_V = 2
            } else if matches_J == 2 {
                add_V = 4
            } else {
                add_V = 0
            }
        } else if sum_M == 3 {
            if matches_J == 1 || matches_J == 3 {
                add_V = 3
            } else {
                add_V = 0
            }
        } else if sum_M == 4 {
            if matches_J == 2 || matches_J == 3 {
                add_V = 6
            }else {
                add_V = 0
            }
        } else if sum_M == 6 {
            if matches_J == 1 || matches_J == 4{
                add_V = 4
            }
        } else {
            if sum_M == 0 {
                add_V = 1
            } else {
                    add_V = 0
            }
        }
    }
    return (sum_M + add_V)
}


/* appending hand's value to it's appropriate types */
/* append switch, depending on it's type it gets appended */
func append_S(sum_M, cur_H int) {
    switch sum_M {
        case 0 : 
            append_T("high_C", cur_H)
            break
        case 1 : 
            append_T("one_P", cur_H)
            break
        case 2 :    
            append_T("two_P", cur_H)
            break
        case 3 :
            append_T("three_K", cur_H)
            break
        case 4 :
            append_T("full_H", cur_H)
            break
        case 6 :
            append_T("four_K", cur_H)
            break
        case 10 :
            append_T("five_K", cur_H)
            break
        default :
            break
    }
}

/* appending elements to it's types */
func append_T(type_H string, cur_H int) {
    types_M[type_H] = append(types_M[type_H], cur_H)
}


/* converting string to integer */
func change_S(numstr string) int {
    num, _ := strconv.Atoi(numstr)
    return num
}

/* __main__ */
func main() {
    var input_U int
    /* opening file */
    file, err := os.Open("input")

    /* handling file error */
    if err != nil {
        fmt.Println("File Error : ", err)
        return 
    }

    /* closing file */
    defer file.Close()

    /* user input */
    fmt.Printf("Enter choice \n1.Part One \n2.Part Two with joker rule applied \n>>> ")
    fmt.Scanf("%d", &input_U)

    /* depending on user input, result is determined */
    if input_U == 2 {
        status_J = true
    } else {
        status_J = false
    }

    /* buffer scanner */
    scanner := bufio.NewScanner(file)

    /* iterating over input file contents */
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")

        hands_G = append(hands_G, []string{strings.Trim(line[0], " ")})
        bets_G = append(bets_G, []int{change_S(strings.Trim(line[1], " "))})

        split_T()
    }

    /* iteratings hands collection based 
       on its types to display the results */
    for types, hands := range types_M {
        arrange_H(types, hands)
    }
    fmt.Printf("Part %d result => ", input_U)
    total_W()

}
