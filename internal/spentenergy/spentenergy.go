package spentenergy

import (
	"errors"
	"time"
)

const (
	stepLengthCoefficient      = 0.45
	mInKm                      = 1000.0
	minInH                     = 60.0
	walkingCaloriesCoefficient = 0.5
)

// Distance считает дистанцию в километрах по шагам и росту.
func Distance(steps int, height float64) float64 {
	if steps < 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

// MeanSpeed считает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || height <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distance / hours
}

// RunningSpentCalories считает потраченные калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры для расчёта калорий при беге")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// WalkingSpentCalories считает потраченные калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры для расчёта калорий при ходьбе")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
