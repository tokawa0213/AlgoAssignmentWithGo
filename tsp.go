// +build ignore

package main

import (
	"fmt"
	"math"
)

func distance(p1 []float64,p2 []float64)int{
	return int(math.Round(math.Sqrt((p1[0]-p2[0])*(p1[0]-p2[0])+(p1[1]-p2[1])*(p1[1]-p2[1]))))
}

func main(){
	var input_number int
	var x,y float64
	var point_array [][]float64
	var answer []int

	fmt.Scan(&input_number)
	for i:=0;i<input_number;i++{
		fmt.Scan(&x,&y)
		point := []float64{x,y}
		point_array = append(point_array,point)
	}
	starting_point := point_array[0]

}