package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2021/helper"
)

func main(){
	data:= helper.ReadFileString("data6.txt")

	tab := make([]int,0)

	d := strings.Split(data[0],",")

	for _,v:=range d{
		i,_:= strconv.Atoi(v)

		tab = append(tab,i)
	}

	for i:=0;i<80;i++{
		new := 0
		for i,v:=range tab{
			if v>0{
				tab[i]--
			}
			if v==0{
				tab[i]=6
				new++
			}
		}
		for i:=0;i<new;i++{
			tab = append(tab,8)
		}
	}
	fmt.Println(len(tab))
}