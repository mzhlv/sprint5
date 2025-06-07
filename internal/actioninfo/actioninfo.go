package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для тренировок и прогулок.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info — выводит инфу по всем строкам dataset через dp.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга строки '%s': %v", data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации по строке '%s': %v", data, err)
			continue
		}
		fmt.Println(info)
	}
}
