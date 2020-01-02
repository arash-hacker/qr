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
	DIV := GPG1.Divide(GPG2)

	GPG1.SetGroupBlock(
		utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t),
		utils.GetGroupBlock(v, e))

	GPG1.SetGroupBlockECC(
		utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t),
		utils.GetGroupBlock(v, e),
		utils.GetECCWTotalNumberOfCodeWord(v, e),
		utils.GetECCWCount(v, e),
	)
	s:=GPG1.Serialize(utils.GetGroupBlock(v,e)["GROUP1"][0]+utils.GetGroupBlock(v,e)["GROUP2"][0])
	//fmt.Println(m.ParseBinary())
	//fmt.Println(utils.BreakUp8Bit(m.Data, m.ParseBinary(), v, e, t), )
	DIV.ToAntilog()

	fmt.Println(s)
}
