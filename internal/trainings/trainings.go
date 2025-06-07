package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training — данные о тренировке
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("некорректный формат строки тренировки")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга количества шагов: %w", err)
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}

	if steps <= 0 {
		return errors.New("шаги должны быть положительным числом")
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительным числом")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

// ActionInfo возвращает инфу о тренировке (строка и ошибка)
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег", "run", "running":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчета калорий для бега: %w", err)
		}
		return fmt.Sprintf(
			"Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, speed, calories,
		), nil
	case "ходьба", "walk", "walking":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчета калорий для ходьбы: %w", err)
		}
		return fmt.Sprintf(
			"Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, speed, calories,
		), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
у