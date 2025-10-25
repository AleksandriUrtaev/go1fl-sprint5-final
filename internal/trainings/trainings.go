package trainings

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

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	sl := strings.Split(datastring, ",")
	var parts []string
	for _, v := range sl {
		parts = append(parts, v)
	}
	if len(parts) != 3 {
		return errors.New("Неверное количество аргументов")
	}
	Steps, err := strconv.Atoi(parts[0])
	if err != nil {
		err := errors.New("Ошибка преобразования.")
		return err
	}

	if Steps <= 0 {
		return errors.New("Количество шагов должно быть больше 0.")
	}

	t.Steps = Steps

	t.TrainingType = parts[1]

	td, err := time.ParseDuration(parts[2])
	if err != nil {
		err := errors.New("Ошибка преобразования.")
		return err
	}
	if td <= 0 {
		return errors.New("Продолжительность должна быть больше 0.")
	}
	t.Duration = td

	return nil
}

func (t Training) ActionInfo() (string, error) {

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)

	if dist <= 0 {
		return "", errors.New("Дистанция <=0")
	}

	aspeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if aspeed <= 0 {
		return "", errors.New("Средняя скорость <=0")
	}
	var SpentCal float64
	var err error
	toa := t.TrainingType
	switch toa {

	case "Бег":
		SpentCal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		SpentCal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		log.Println("Ошибка:", err)
		return "", errors.New("ошибка получения значения сожженых калорий.")
	}

	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`,
		toa,
		t.Duration.Hours(),
		dist,
		aspeed,
		SpentCal), nil

}
