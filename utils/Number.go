package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

type Number struct {
	Data string
}

func (n *Number) Parse() string {
	r, _ := regexp.Compile("...")
	datas := r.FindAllString(n.Data, -1)

	if len(n.Data)%3 == 0 {

	} else if len(n.Data)%3 == 1 {
		datas = append(datas, n.Data[len(n.Data)-1:])
	} else if len(n.Data)%3 == 2 {
		datas = append(datas, n.Data[len(n.Data)-2:])
	}
	s := ""
	for _, data := range datas {
		if len(data)%3 == 0 {
			i, _ := strconv.ParseInt(data, 10, 64)
			s =s+fmt.Sprintf("%010b", i )
		} else if len(data)%3 == 2 {
			i, _ := strconv.ParseInt(data, 10, 64)
			s =s+fmt.Sprintf("%07b", i)
		} else if len(data)%3 == 1 {
			i, _ := strconv.ParseInt(data, 10, 64)
			s =s+fmt.Sprintf("%04b", i)
		}
	}
	return s
}
