package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps — структура для дневных шагов.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "678,0h50m"
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат строки для DaySteps")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга количества шагов: %w", err)
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}

	if steps <= 0 {
		return errors.New("шаги должны быть положительным числом")
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительным числом")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo возвращает строку с инфой о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий для прогулки: %w", err)
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
