package main

import (
	"fmt"
	"generator_polynomial_galois/gpg"
	"qr/utils"
)

func main() {
	v := 5
	e := 'Q'
	t := 'A'
	m := utils.Alpha{"HELLO WORLD"}
	//TODO create object polymorphism

	GPG1 := gpg.NewAntiLog(
		utils.ConvertToMessagePoly(
			utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t)),
	)
	GPG2 := gpg.GenGalois(
		utils.GetECCWCount(v, e),
	)
	div:=GPG1.Divide(
		GPG2, 
		utils.GetECCWTotalNumberOfCodeWord(v, e))

	GPG1.SetGroupBlock(
		utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t), 
		utils.GetGroupBlock(v, e))

	//fmt.Println(m.ParseBinary())
	fmt.Println(utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t), )
	fmt.Println(GPG1.GroupBlock)
	fmt.Println(div)
}
