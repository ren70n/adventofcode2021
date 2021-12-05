package main

import (
	"fmt"
	"adventofcode2021/helper"
)

func main() {
	dataFile := "data1.txt"

	data := helper.ReadFileInt(dataFile)

	cnt := 0

	for i, _ := range data[:len(data)-3] {
		a,b := 0,0
		for j:=0;j<3;j++{
			a+=data[i+j]
		}
		for j:=1;j<4;j++{
			b+=data[i+j]
		}
		
		if a<b{
			cnt++
		}
	}
	fmt.Println(cnt)
}

