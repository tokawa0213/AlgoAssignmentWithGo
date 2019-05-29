// +build ignore

package main

import (
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

func add_up(grid [][]int,i,j int)int{
	return grid[i][j-1]+grid[i][j+1]+grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1]+grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1]
}

func main(){
	const COLUMN = 102
	const ROW = 52
	const N = 10
	generation := make([][]int, ROW)
	next_generation := make([][]int, ROW)
	for i:=0;i< ROW;i++{
		generation[i] = make([]int, COLUMN)
		next_generation[i] = make([]int, COLUMN)
	}
	f,err := os.Open("proc/life_game.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	err2 := termbox.Init()
	if err2 != nil {
		panic(err2)
	}
	defer termbox.Close()
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		row := scanner.Text()
		for j := 0; j < len(row); j++ {
			generation[count+1][j+1], _ = strconv.Atoi(row[j : j+1])
		}
		count++
	}
	MAINLOOP:
	for{
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
				case termbox.KeyEsc:
					break MAINLOOP
				case termbox.KeyEnter:
					for i := 1; i < ROW-1; i++ {
						for j := 1; j < COLUMN-1; j++ {
							if generation[i][j] == 0 {
								if add_up(generation, i, j) == 3 {
									next_generation[i][j] = 1
								}
							} else {
								result := add_up(generation, i, j)
								if result == 2 || result == 3 {
									next_generation[i][j] = 1
								} else {
									next_generation[i][j] = 0
								}
							}
						}
					}
					//deep copy
					generation = append([][]int{}, next_generation...)
					next_generation = make([][]int, ROW)
					for i := 0; i < ROW; i++ {
						next_generation[i] = make([]int, COLUMN)
					}
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					for i := 1; i < ROW-1; i++ {
						for j := 1; j < COLUMN-1; j++ {
							if generation[i][j] == 1{
								termbox.SetCell(i, j, []rune(strconv.Itoa(generation[i][j]))[0], termbox.ColorRed, termbox.ColorDefault)
							}else{
								termbox.SetCell(i, j, []rune(strconv.Itoa(generation[i][j]))[0], termbox.ColorDefault, termbox.ColorDefault)
							}
						}
					}
					termbox.Flush()
				default:
					continue
			}
		}
	}
}