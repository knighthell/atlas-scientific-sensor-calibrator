package sensor

import (
	"fmt"
	"strings"
)

type RTDSensor struct {
	Sensor
}

func (s *RTDSensor) Calibration() {
	var value string

	fmt.Println("온도센서 교정을 시작합니다.")
	fmt.Println("현재 온도센서가 담겨있는 물의 섭시온도는 몇도인가요? ex) 25.5")
	fmt.Print("온도: ")
	_, _ = fmt.Scanln(&value)
	fmt.Printf("%v℃로 교정시작합니다.\n", value)
	result, _ := write(s.Port, "Cal,"+value)
	if strings.Contains(result, "OK") {
		println("온도센서 교정이 완료되었습니다.")
	}
}
