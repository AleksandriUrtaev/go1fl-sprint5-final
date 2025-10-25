package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, r := range dataset {

		err := dp.Parse(r)
		if err != nil {
			log.Println("Ошибка:", err)
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка:", err)
			continue
		}
		print(str)
	}
}
