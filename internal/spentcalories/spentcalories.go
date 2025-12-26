package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataTrimmed := strings.TrimSpace(data)
	if len(dataTrimmed) == 0 {
		return 0, "", 0, fmt.Errorf("empty input data")
	}

	dataSplitted := strings.Split(dataTrimmed, ",")
	if len(dataSplitted) != 3 {
		return 0, "", 0, fmt.Errorf("wrong input data")
	}

	steps, err := strconv.Atoi(dataSplitted[0])
	if err != nil {
		return 0, "", 0, err
	}

	walkDuration, err := time.ParseDuration(dataSplitted[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, dataSplitted[1], walkDuration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return stepLength * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)

	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var dist float64
	var speed float64
	var spentCalories float64

	switch training {
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		spentCalories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		spentCalories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %2f ч.\nДистанция: %2f км.\nСкорость: %2f км/ч\nСожгли калорий: %2f\n", training, duration.Hours(), dist, speed, spentCalories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("wrong input data")
	}

	speed := meanSpeed(steps, height, duration)

	return (weight * speed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("wrong input data")
	}

	speed := meanSpeed(steps, height, duration)

	return ((weight * speed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}
