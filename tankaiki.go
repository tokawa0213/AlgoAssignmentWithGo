// +build ignore

package main

import (
	"fmt"
)

func solve(a,b,ans0,c,d,ans1 float64) (float64,float64){
	ad_bc := (1/(a*d-b*c))
	a,b,c,d = ad_bc*d,-ad_bc*b,-ad_bc*c,ad_bc*a
	answer_a := a*ans0+b*ans1
	answer_b := c*ans0+d*ans1
	return answer_a,answer_b
}

func main(){
	var input_num int
	var x,y,xi2,xiyi,xi,yi float64
	var point_info [][]float64

	fmt.Scan(&input_num)
	for i:=0;i < input_num;i++{
		var point []float64
		fmt.Scan(&x,&y)
		point = append(point,x,y)
		point_info = append(point_info,point)
	}
	for _,p := range(point_info){
		xi += p[0]
		yi += p[1]
		xi2 += p[0]*p[0]
		xiyi += p[0]*p[1]
	}
	a,b := solve(2*xi2,2*xi,2*xiyi,2*xi,float64(2*input_num),2*yi)
	fmt.Printf("%.4f %.4f",a,b)
}