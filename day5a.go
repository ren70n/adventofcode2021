package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2021/helper"
)

func main(){
	data:= helper.ReadFileString("data5.txt")

	respoints := make(map[string] int)

	for _, row:=range data{
		points := strings.Split(row," -> ")

		pA,pB := []int{0,0},[]int{0,0}

		for i,a := range strings.Split(points[0],","){
			pA[i],_ = strconv.Atoi(a)
		}

		for i,b:=range strings.Split(points[1],","){
			pB[i],_ = strconv.Atoi(b)
		}
		// is line correct
		if pA[0] == pB[0] {
			if pA[1]>pB[1]{ pA,pB = pB,pA}

			for i:=pA[1];i<=pB[1];i++{
				respoints[fmt.Sprint(pA[0],i)]++
			}
		}

		if pA[1] == pB[1]{
			if pA[0]>pB[0]{ pA,pB = pB,pA}

			for i:=pA[0];i<=pB[0];i++{
				respoints[fmt.Sprint(i,pA[1])]++
			}
		}
	}
	res := 0
	for _,v:=range respoints{
		if v>1{res++}
	}
	fmt.Println(res)
}