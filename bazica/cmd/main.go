package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tommitoan/bazica/internal/fourpillars"
	"github.com/tommitoan/bazica/internal/luckpillars"
	"github.com/tommitoan/bazica/internal/ultis"
	"github.com/tommitoan/bazica/model"
)

// Пример использования библиотеки для расчета ба-цзы
func main() {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	gender := 1 // 1 = male, 0 = female

	// Пример использования библиотеки
	fourPillar, passed, remaining, err := fourpillars.GetFourPillars(now, loc, "")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Обработка результата
	luckPillars, err := luckpillars.GetLuckPillars(fourPillar, gender, passed, remaining, now, "")
	if err != nil {
		fmt.Println("Ошибка расчета столпов удачи:", err)
		return
	}

	chart := model.BaziChart{
		FourPillar:  ultis.GetLifeCycleFromFourPillar(fourPillar),
		LuckPillars: luckPillars,
	}

	// Вывод результата в формате JSON
	jsonData, _ := json.MarshalIndent(chart, "", "  ")
	fmt.Println("Рассчитанный ба-цзы график:")
	fmt.Println(string(jsonData))
}
