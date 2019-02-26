package sensor

import (
	"fmt"
	"strings"
)

type PHSensor struct {
	Sensor
	MidPoint  string
	LowPoint  string
	HighPoint string
}

func (s *PHSensor) Calibration() {
	s.MidPoint = "7.00"
	s.LowPoint = "4.00"
	s.HighPoint = "10.00"

	var value string

	fmt.Println("pH센서 교정을 시작합니다.")

	fmt.Println("pH센서 싱글(미드)포인트 교정을 시작합니다.")
	fmt.Println("마개를 연 후, 흐르는 물에 pH센서를 세척합니다.")
	fmt.Println("pH 센서 캘리브레이션용 노란색 용기의 용액을 컵에 따른 후, 센서를 담그고 1~2분정도 기다립니다. 시간이 지나면 엔터를 누르십시오.")
	_, _ = fmt.Scanln(&value)
	result, _ := write(s.Port, "Cal,mid,"+s.MidPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("싱글(미드)포인트 교정이 되었습니다.")
	} else {
		fmt.Println("싱글(미드)포인트 교정 실패")
		fmt.Println("싱글(미드)포인트 교정 실패")
	}

	fmt.Println("pH센서 투(로우)포인트 교정을 시작합니다")
	fmt.Println("마개를 연 후, 흐르는 물에 pH센서를 세척합니다.")
	fmt.Println("pH 센서 캘리브레이션용 빨간색 용기의 용액을 컵에 따른 후, 센서를 담그고 1~2분정도 기다립니다. 시간이 지나면 엔터를 누르십시오.")
	_, _ = fmt.Scanln(&value)
	result, _ = write(s.Port, "Cal,low,"+s.LowPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("투(로우)포인트 교정이 되었습니다.")
	} else {
		fmt.Println("투(로우)포인트 교정 실패")
		fmt.Println("투(로우)포인트 교정 실패")
	}

	fmt.Println("pH센서 쓰리(하이)포인트 교정을 시작합니다")
	fmt.Println("마개를 연 후, 흐르는 물에 pH센서를 세척합니다.")
	fmt.Println("pH 센서 캘리브레이션용 파란색 용기의 용액을 컵에 따른 후, 센서를 담그고 1~2분정도 기다립니다. 시간이 지나면 엔터를 누르십시오.")
	_, _ = fmt.Scanln(&value)
	result, _ = write(s.Port, "Cal,high,"+s.HighPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("쓰리(하이)포인트 교정이 되었습니다.")
	} else {
		fmt.Println("쓰리(하이)포인트 교정 실패")
		fmt.Println("쓰리(하이)포인트 교정 실패")
	}
}
