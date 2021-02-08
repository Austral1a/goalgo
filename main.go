package main

import (
    "fmt"
    "sort"
    "strings"
)

/*Is he gonna survive? (8 kyu)*/

func Hero(bullets, dragons int) bool {
    return bullets >= dragons * 2
}

/*Two to one (7 kyu)*/
func TwoToOne(s1 string, s2 string) string {
    m := make(map[string]int)

    for _, a := range s1 {
        for _, b := range s2 {
            m[string(a)]++
            m[string(b)]++
        }
    }

    var res []string

    for k, _ := range m {
        res = append(res, k)
    }

    sort.Strings(res)

    return strings.Join(res, "")
}

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

/*Moves in squared strings (I) (7 kyu)*/
func VertMirror(s string) string {
    var res []string
    for _, v := range strings.Split(s, `\`) {
        temp := strings.Split(v, "")

        for i, j := 0, len(v) - 1; i < j; i, j = i+1, j-1 {
            temp[i], temp[j] = temp[j], temp[i]
        }
        res = append(res, strings.Join(temp, "") + "\\")
    }
    fmt.Println(res)
    return strings.Join(res, " ")
}

/*Multiplies of 3 or 5 (6 kyu)*/
func Multiple3And5(number int) int {
    var res []int
    for i := 0; i < number; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            res = append(res, i)
        } else if i % 3 == 0 {
            res = append(res, i)
        } else if i % 5 == 0 {
            res = append(res, i)
        }
    }

    n := 0
    for _, v := range res {
        n += v
    }

    return n
}

func main() {
    n := Multiple3And5(10)
    fmt.Println(n)
}
