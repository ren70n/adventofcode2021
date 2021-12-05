package main

import (
	"fmt"
	"strconv"
	"adventofcode2021/helper"
)

func main() {
	data:= helper.ReadFileString("data3.txt")

	m := make(map[int]int)
	for _,s := range data{

		for i,v:=range s{
			if v =='1'{
				m[i]++	
			}
		}
	}
	l := len(data)
	gamStr,epsStr := "",""
	for i:=0;i<len(data[0]);i++{
		if l/2 > m[i]{
			gamStr+="0"
			epsStr+="1"
		}else{
			gamStr+="1"
			epsStr+="0"
		}
	}
	gamma,_ := strconv.ParseInt(gamStr,2,32)
	epsilon,_ := strconv.ParseInt(epsStr,2,32)
	fmt.Println(gamma* epsilon)
}