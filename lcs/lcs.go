package lcs

import (
	"fmt"
	"math"
	"strings"
)

func Reverse( orig string ) string {
	var c []string = strings.Split(orig, "");

	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	return strings.Join( c, "" );
}

func Lcs(){
	var string_a,string_b string
	var DP [][]int
	fmt.Scan(&string_a)
	fmt.Scan(&string_b)
	for i:=0;i<len(string_a)+1;i++{
		DP = append(DP,make([]int,len(string_b)+1))
	}
	//creating the DP grid
	for i := 1;i<len(string_a)+1;i++{
		for j := 1;j<len(string_b)+1;j++{
			if string_a[i-1] == string_b[j-1]{
				DP[i][j] = DP[i-1][j-1] + 1
			}else{
				DP[i][j] = int(math.Max(float64(DP[i-1][j]),float64(DP[i][j-1])))
			}
		}
	}
	//searching the DP grid
	var answer string
	answer_length := DP[len(string_a)][len(string_b)]
	for i := len(string_a)-1;i>0;i--{
		for index,value := range(DP[i][1:]){
			if value == answer_length && value !=0{
				answer += string_b[index:index+1]
				answer_length--
				break
			}
		}
	}
	fmt.Println(Reverse(answer))
}
