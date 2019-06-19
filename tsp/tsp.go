package tsp

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MAX = 10000000

func distance(p1 []float64,p2 []float64)int{
	return int(math.Round(math.Sqrt((p1[0]-p2[0])*(p1[0]-p2[0])+(p1[1]-p2[1])*(p1[1]-p2[1]))))
}

func fill_distance_grid(i,j int, distance_grid *[][]int,point_array [][]float64){
	if i != j{
		(*distance_grid)[i][j] = distance(point_array[i],point_array[j])
	}
}

func calculate_distance(route []int,distance_array [][]int)int{
	answer := 0
	for i:=0;i<len(route)-1;i++{
		answer += distance_array[route[i]][route[i+1]]
	}
	return answer
}

func search_in_array(array []int,find int)bool{
	for _,value := range array{
		if value == find{
			return true
		}
	}
	return false
}

func Tsp() {
	var input_number int
	var x, y float64
	var point_array [][]float64
	var distance_array [][]int
	var distance_array_piece []int
	fmt.Scan(&input_number)

	for i := 0; i < input_number; i++ {
		distance_array_piece = append(distance_array_piece, MAX)
	}

	for i := 0; i < input_number; i++ {
		fmt.Scan(&x, &y)
		point := []float64{x, y}
		point_array = append(point_array, point)
		distance_array = append(distance_array, append([]int{}, distance_array_piece...))
	}

	start := time.Now()
	for i := 0; i < input_number; i++ {
		for j := 0; j < input_number; j++ {
			fill_distance_grid(i, j, &distance_array, point_array)
		}
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

	//Search
	visited := []int{0}
	arr := []int{}
	start_point := 0
	new_start_point := 0
	for{
		min := MAX
		for i:=0;i<input_number;i++{
			if min>distance_array[start_point][i] && !search_in_array(visited,i){
				min = distance_array[start_point][i]
				new_start_point = i
			}
		}
		if min == MAX{break}
		arr = append(arr,distance_array[start_point][new_start_point])
		visited = append(visited,new_start_point)
		start_point = new_start_point
	}

	score := calculate_distance(visited,distance_array)
	fmt.Println(score)

	fmt.Println(visited)

	//2-opt
	for i:=0;i<1000000;i++{
		first := rand.Intn(len(visited))
		second := rand.Intn(len(visited))
		if first == second{continue}
		visited[first],visited[second] = visited[second],visited[first]
		new_score := calculate_distance(visited,distance_array)
		if score > new_score{
			score = new_score
		}else{
			visited[first],visited[second] = visited[second],visited[first]
		}
		fmt.Println(score)
	}
	fmt.Println(visited)
}