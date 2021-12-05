package main

import (
	"fmt"
	"strconv"
	"adventofcode2021/helper"
)

func main() {
	data:= helper.ReadFileString("data3.txt")


	gamma := getGamma(data)
	epsilon := getEpsilen(data)

	g,_ := strconv.ParseInt(gamma,2,64)
	e,_ := strconv.ParseInt(epsilon,2,64)

	fmt.Println(g*e)
}

func getGamma(data []string)string{
	if len(data)==1{
		return string(data[0])
	}

	z,o := getZeroOne(data)

	toSave := "0"
	if o>=z{
		toSave = "1"
	}

	nData := make([]string,0)

	for _,v:=range(data){
		if string(v[0])==toSave{
			nData = append(nData,v[1:])
		}
	}

	return toSave + getGamma(nData)
}
func getEpsilen(data []string)string{
	if len(data)==1{
		return string(data[0])
	}

	z,o := getZeroOne(data)

	toSave := "1"
	if o>=z{
		toSave = "0"
	}

	nData := make([]string,0)

	for _,v:=range(data){
		if string(v[0])==toSave{
			nData = append(nData,v[1:])
		}
	}

	return toSave + getEpsilen(nData)
}

func getZeroOne(data []string)(int,int){
	zero,one:=0,0
	for _,d:=range data{
		if d[0]=='1'{one++}else{zero++}
	}
	return zero,one
}