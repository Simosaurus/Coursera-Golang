package main

import(
	"fmt"
	"strconv"
	"sort"
)

func main(){
	var intext string
	sli := make([]int,0)
	test := "Y"
	for test != "X"{
		fmt.Println("Enter an integer or X to stop")
		fmt.Scanln(&intext)
		test = intext
		if innumber, err := strconv.Atoi(intext); err == nil {
			sli = append(sli,innumber)
			sort.Slice(sli, func(i, j int) bool {
    			return sli[i] < sli[j] })
			fmt.Println(sli)
			} else {
				fmt.Println("Not an int!")
			}
		}
		fmt.Println("X detected, end of program")
		
		}