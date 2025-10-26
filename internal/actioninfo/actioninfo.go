package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, r := range dataset {

		err := dp.Parse(r)
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}
		print(str)
	}
}
