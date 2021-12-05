package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2021/helper"
)

func main() {
	data:= helper.ReadFileString("data2.txt")

	h,d,aim := 0,0,0
	for _,s := range data{
		row := strings.Split(s," ")
		x,_:=strconv.Atoi(row[1])
		switch row[0]{
		case "forward":
			h+=x
			d+=aim*x
		case "up":
			aim-=x
		case "down":
			aim+=x
		}
	}
	fmt.Println(h*d)
}