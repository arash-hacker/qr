package utils

import (
	"fmt"
)

type Byte struct {
	Data string
}

func (n *Byte) Parse() string {
	s := ""
	for _, data := range n.Data {
		s = s + fmt.Sprintf("%08b", data)
	}
	return s
}
