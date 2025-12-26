package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSplitted := strings.Split(data, ",")
	if len(dataSplitted) != 2 {
		return 0, 0, spentcalories.ErrWrongInputData
	}

	steps, err := strconv.Atoi(dataSplitted[0])
	if err != nil {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(dataSplitted[1])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 || duration <= 0 {
		return 0, 0, spentcalories.ErrWrongInputData
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if steps <= 0 {
		if err != nil {
			log.Println(err)
		}
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, spentCalories)
}
