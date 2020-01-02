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
	t:='N'
	m := utils.Alpha{"876876"}
	//TODO create object polymorphism
	fmt.Println(
		utils.BreakUp8Bit(m.Data,m.Parse(),v,e,t),
	)
	
	


	//fmt.Println(version,error_code_level)
}
