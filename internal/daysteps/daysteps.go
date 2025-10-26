package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("ошибка преобразования")
	}

	stepStr := parts[0]
	//if strings.TrimSpace(stepStr) != stepStr {
	//	return errors.New("Неверный формат шагов")
	//}

	ds.Steps, err = strconv.Atoi(stepStr)
	if err != nil {
		return errors.New("ошибка преобразования шагов")
	}
	if ds.Steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}

	ds.Duration, err = time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return errors.New("ошибка преобразования длительности")
	}

	if ds.Duration <= 0 {
		return errors.New("длительность должна быть больше 0")
	}

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println("ошибка расчёта калорий:", err)
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, cal,
	), nil

}
