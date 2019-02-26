package main

import (
	"./sensor"
	"fmt"
)

func main() {

	rtd, ec, do, ph, err := sensor.GetSensors("a")
	if err != nil {
		fmt.Printf("센서 포트를 설정하는데 오류가 났습니다.\n%v\n", err)
	}

	if rtd != nil {
		rtd.Calibration()
	}

	if ec != nil {
		ec.Calibration()
	}

	if do != nil {
		do.Calibration()
	}

	if ph != nil {
		ph.Calibration()
	}

	fmt.Println("교정을 모두 완료하였습니다.")
}
