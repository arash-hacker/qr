package main;

import (
	"fmt"
	gg "generator_polynomial_galois/gpg"

)

func main(){
	fmt.Println("ya");

	g:=gg.GenPoly{}
	// g.Init(
	// 	map[int]int{
	// 		240:1,
	// 		230:7,
	// 		210:5,
	// 		300:12,
	// 		55:4,
	// 	},
	// )

	// g.Sort()
	// fmt.Println(g.GetAll())

	// g.MultiplyCoesBy(3)
	// fmt.Println(g.GetAll())

	// g.SumExposBy(200)
	// fmt.Println(g.GetAll())
	g.GenGalois(10);
	fmt.Println(g.GetAll())

}