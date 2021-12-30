package main

import (
	"fmt"
	"strings"
	// "strconv"
	// "math"
	"adventofcode2021/helper"
)

func main(){
	data:= helper.ReadFileString("data8.txt")

	ret := 0

	for _, row:=range data{
		a := row[strings.Index(row," | ")+3:]
		r := strings.Split(a," ")
		
		for _,v:=range r{
			l := len(v)

			if l==4 || l==3 || l==7 || l==2{
				ret++
			}
		}
	}
	fmt.Println(ret)
}