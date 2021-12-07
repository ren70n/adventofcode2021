package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2021/helper"
)
type board struct{
	numbers [][]int
	shot [][]int

}
func main() {
	data:= helper.ReadFileString("data4.txt")

	bingoStr := strings.Split(data[0],",")
	bingoInt := make([]int,0)
	for _,v:=range bingoStr{
		i,e := strconv.Atoi(v)
		if e==nil{
			bingoInt = append(bingoInt,i)
		}
	}

	boards := fillBoards(data[2:])

	for _, num := range bingoInt{
		boards = markBoards(boards,num)
		b,e := checkLines(boards)
		
		if e == true{
			// fmt.Println(b)
			fmt.Println(countRes(b,num))
			break
		}
	}

}

func countRes(b board,i int)int{
	res := 0
	for y:=0;y<5;y++{
		for x:=0;x<5;x++{
			if b.shot[y][x]==0{
				res+=b.numbers[y][x]
			}
		}
	}

	return res * i
}

func checkLines(boards []board)(board,bool){
	for _, board := range boards{
		for _,row:=range board.shot{
			z:=0
			for _,v:=range row{
				if v==1{z++}
			}
			if z==5 { return board,true}
		}
		for i,_:=range board.shot{
			z:=0
			for x:=0;x<5;x++{
				if board.shot[i][x]==1{z++}
			}
			if z==5 { return board,true}
		}
	}
	return board{}, false
}

func markBoards(boards []board,val int)[]board{
	for _,board := range boards{

		for y, row:= range board.numbers{
			for x, bval := range row{
				if bval == val{
					board.shot[y][x]=1
				} 
			}
		}
	}
	return boards
}
func fillBoards(data []string)[]board{
	boards := make([]board,0)

	tBoard := board{}
	n := make([][]int,0)
	for _,v:=range data{

		if v==""{
			tBoard.shot = make([][]int,5)
			for i,_:= range tBoard.shot{
				tBoard.shot[i] = make([]int,5)
			}
			tBoard.numbers = n
			boards = append(boards,tBoard)
			n = make([][]int,0)
		}else{
			row := make([]int,0)
			num := strings.Split(v," ")
			for _, f:= range num{
				in, e := strconv.Atoi(f)
				if e==nil{
					row = append(row,in)
				}
			}
			n = append(n,row)
		}
	}
	tBoard.shot = make([][]int,5)
	for i,_:= range tBoard.shot{
		// fmt.Println(tBoard.shot[i],make([]int,5))
		tBoard.shot[i] = make([]int,5)
	}
	tBoard.numbers = n
	boards = append(boards,tBoard)
	n = make([][]int,0)

	return boards
}
