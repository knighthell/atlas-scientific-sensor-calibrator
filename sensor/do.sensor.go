package sensor

import (
	"fmt"
	"strings"
)

type DOSensor struct {
	Sensor
}

func (s *DOSensor) Calibration() {
	var value string

	fmt.Println("DO센서 교정을 시작합니다.")

	fmt.Println("DO센서 싱글포인트 교정을 시작합니다. DO센서의 마개를 분리 후 약 30초정도 공기에 노출시킨 후 엔터를 누르세요.")
	_, _ = fmt.Scanln(&value)
	result, _ := write(s.Port, "Cal")
	if strings.Contains(result, "OK") {
		fmt.Println("싱글(미드)포인트 교정이 되었습니다.")
	} else {
		fmt.Println("싱글(미드)포인트 교정 실패")
		fmt.Println("싱글(미드)포인트 교정 실패")
	}

	fmt.Println("추가적으로, 영점 조정이 가능한 DO용액(노란케이스 용액)이 있을 경우 듀얼포인트 교정을 진행할 수 있습니다.")
	fmt.Println("듀얼포인트 교정을 진행하시겠습니까? 진행하시겠으면 'Y'를 입력해주세요.")
	_, _ = fmt.Scanln(&value)
	if value != "Y" {
		fmt.Println("DO센서 교정이 완료되었습니다.")
		return
	}

	fmt.Println("듀얼포인트 교정을 시작합니다.")
	fmt.Println("DO Zero 용액에 센서를 담근 후, 1분 30초 뒤 엔터를 누르세요.")
	_, _ = fmt.Scanln(&value)
	result, _ = write(s.Port, "Cal,0")
	if strings.Contains(result, "OK") {
		fmt.Println("듀얼포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("듀얼포인트 교정 실패")
		fmt.Println("듀얼포인트 교정 실패")
	}

	fmt.Println("DO센서 교정이 완료되었습니다.")
}
