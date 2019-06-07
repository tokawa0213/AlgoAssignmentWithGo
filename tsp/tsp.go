package tsp

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type point_object struct {
	x int
	y int
	value int
}

const MAX = 10000000

func distance(p1 []float64,p2 []float64)int{
	return int(math.Round(math.Sqrt((p1[0]-p2[0])*(p1[0]-p2[0])+(p1[1]-p2[1])*(p1[1]-p2[1]))))
}

func fill_distance_grid(i,j int, distance_grid *[][]int,point_array [][]float64){
	if i != j{
		(*distance_grid)[i][j] = distance(point_array[i],point_array[j])
	}
}

func array_min(number ...int)point_object{
	p := point_object{}
	var min = MAX

	for index,value := range number{
		if value < min{
			p.value = value
			p.x = index/int(math.Sqrt(float64(len(number))))
			p.y = index%int(math.Sqrt(float64(len(number))))
			min = value
		}
	}
	return p
}

func flatten(a [][]int) []int{
	var ret []int
	for _,value := range a{
		ret = append(ret,value...)
	}
	return ret
}

func Tsp(){
	var input_number int
	var x,y float64
	var point_array [][]float64
	var distance_array [][]int
	var distance_array_piece []int
	fmt.Scan(&input_number)

	for i:=0;i<input_number;i++{
		distance_array_piece = append(distance_array_piece,MAX)
	}

	for i:=0;i<input_number;i++{
		fmt.Scan(&x,&y)
		point := []float64{x,y}
		point_array = append(point_array,point)
		distance_array = append(distance_array,append([]int{},distance_array_piece...))
	}

	var wg sync.WaitGroup
	c := make(chan bool,50)

	start := time.Now()
	for i:=0;i<input_number;i++{
		for j:=0;j<input_number;j++{
			wg.Add(1)
			c<-true
			go fill_distance_grid(i,j,&distance_array,point_array)
			<-c
			wg.Done()
		}
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println(distance_array)
	/*
	for{
		p := array_min(flatten(distance_array)...)
		if p.value == MAX{
			return
		}else{
			//record the connected dots
			//change the cell to MAX
		}
	}
	*/
}