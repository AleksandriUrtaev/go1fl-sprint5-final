package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if duration <= 0 {
		return 0, errors.New("Продолжительнось duration <=0")
	}

	if steps <= 0 {
		return 0, errors.New("Количество шагов <=0.")
	}

	if weight <= 0 {
		return 0, errors.New("Масса <=0.")
	}

	if height <= 0 {
		return 0, errors.New("Рост <=0.")
	}

	ms := MeanSpeed(steps, height, duration)

	return (weight * ms * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if duration <= 0 {
		return 0, errors.New("Продолжительнось duration <=0")
	}

	if steps <= 0 {
		return 0, errors.New("Количество шагов <=0.")
	}

	if weight <= 0 {
		return 0, errors.New("Масса <=0.")
	}

	if height <= 0 {
		return 0, errors.New("Рост <=0.")
	}

	ms := MeanSpeed(steps, height, duration)

	return (weight * ms * duration.Minutes()) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	aspeed := dist / duration.Hours()

	return aspeed

}

func Distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / float64(mInKm)

}
