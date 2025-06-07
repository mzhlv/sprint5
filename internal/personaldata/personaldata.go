package personaldata

import "fmt"

// Personal — структура с личными данными пользователя.
type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес (кг)
	Height float64 // Рост (м)
}

// Print — выводит данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
