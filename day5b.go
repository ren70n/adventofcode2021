package main

import (
	"fmt"
	"strings"
	"strconv"
	// "math"
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
		}else if pA[1] == pB[1]{
			if pA[0]>pB[0]{ pA,pB = pB,pA}

			for i:=pA[0];i<=pB[0];i++{
				respoints[fmt.Sprint(i,pA[1])]++
			}
		}else{
			aa := pA[0]-pA[1]
			bb := pB[0]-pB[1]

			aa*=aa
			bb*=bb

			ab := pA[0]-pB[0]
			ba := pA[1]-pB[1]

			ab*=ab
			ba*=ba

			if aa == bb || ab==ba{
				// lower X always on front
				if pA[0]>pB[0]{pA,pB=pB,pA}

				dir := 1
				// which directions
				if pA[1]>pB[1]{dir = -1}

				y := pA[1]
				for x:=pA[0];x<=pB[0];x++{
					respoints[fmt.Sprint(x,y)]++
					y+=dir
				}
			}
		}
	}
	res := 0
	for _,v:=range respoints{
		if v>1{res++}
	}

	fmt.Println(res)
}