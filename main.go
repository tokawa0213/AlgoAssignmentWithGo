package main

import (
	"atcoder_seisen/proc/lcs"
	"atcoder_seisen/proc/life_game"
	"atcoder_seisen/proc/linear_regression"
	"atcoder_seisen/proc/tsp"
	"fmt"
)

func main(){
	var choice string
	fmt.Scan(&choice)
	if choice == "tsp"{
		tsp.Tsp()
	}else if choice == "lcs"{
		lcs.Lcs()
	}else if choice == "linear_regression"{
		linear_regression.Linear_regression()
	}else if choice == "life_game"{
		life_game.Life_game()
	}else{
		panic("ERROR")
	}
}
