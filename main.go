package main

import (
	"fmt"
	"qr/utils"
)

func main() {
	// fmt.Println("ya");
	// g:=gg.GenPoly{}
	// f:=g.GenGalois(63);
	// fmt.Println(f.GetAll())
	v:=1
	e:='Q'
	t:='A'
	// n := utils.Number{"8675309"}
	// fmt.Println(n.Parse())
	m := utils.Alpha{"HELLO WORLD"}
	//o := utils.Byte{"Hello, world!"}
	//fmt.Println(o.Parse())
	fmt.Println(
		utils.BreakUp8Bit("HELLO WORLD",m.Parse(),v,e,t),
	)
	
	


	//fmt.Println(version,error_code_level)
}
