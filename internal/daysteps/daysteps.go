package daysteps

import (
	"fmt"
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
	dataTrimmed := strings.TrimSpace(data)
	if len(dataTrimmed) == 0 {
		return 0, 0, fmt.Errorf("empty input data")
	}

	dataSplitted := strings.Split(dataTrimmed, ",")
	if len(dataSplitted) != 2 {
		return 0, 0, fmt.Errorf("wrong input data")
	}

	steps, err := strconv.Atoi(dataSplitted[0])
	if err != nil {
		return 0, 0, err
	}

	walkDuration, err := time.ParseDuration(dataSplitted[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkDuration, err := parsePackage(data)
	if steps <= 0 {
		if err != nil {
			fmt.Println(err)
		}
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkDuration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Колчество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, spentCalories)
}
