package utils

import (
	"fmt"
	"regexp"
)

type Alpha struct {
	Data string
}

func (n *Alpha) Parse() string {
	r, _ := regexp.Compile("..")
	datas := r.FindAllString(n.Data, -1)

	if len(n.Data)%2 == 0 {

	} else if len(n.Data)%2 == 1 {
		datas = append(datas, n.Data[len(n.Data)-1:])
	}
	s := ""
	for _, data := range datas {
		if len(data)%2 == 0 {
			s = s + fmt.Sprintf("%011b", 45*AlphaNumericValue(data[0])+AlphaNumericValue(data[1]))
		} else {
			s = s + fmt.Sprintf("%06b", AlphaNumericValue(data[0]))
		}
	}
	return s
}
