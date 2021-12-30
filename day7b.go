package main

import (
	"fmt"
	"strings"
	"strconv"
	// "math"
	"adventofcode2021/helper"
)

func main(){
	data:= helper.ReadFileString("data7.txt")

	tmp := strings.Split(data[0],",")

	crabs := make([]int,0) 
	mx:=0
	for _,v:=range tmp{
		i,_:= strconv.Atoi(v)
		crabs = append(crabs,i)
		if i>mx{mx=i}
	}

	prev := 999999999
	for i:=0;i<mx;i++{
		sum:=0
		for _,v:=range crabs{
			l:=v-i
			if l<0{l*=-1}

			for l>0{
				sum+=l
				l--
			}
		}
		if prev>sum{prev = sum}
	}
	fmt.Println(prev)

}