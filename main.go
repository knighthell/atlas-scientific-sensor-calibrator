package main

import (
	"./sensor"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	var sensingTypes string
	if len(args) == 0 {
		sensingTypes = "ALL"
	} else {
		sensingTypes = strings.ToUpper(strings.Join(args, ", "))
	}

	fmt.Printf("설정할 센서 목록: %s\n", sensingTypes)


	rtd, ec, do, ph, err := sensor.GetSensors(args)
	if err != nil {
		fmt.Printf("센서 포트를 설정하는데 오류가 났습니다.\n%v\n", err)
	}

	if rtd.Port != nil {
		rtd.Calibration()
	}

	if ec.Port != nil {
		ec.Calibration()
	}

	if do.Port != nil {
		do.Calibration()
	}

	if ph.Port != nil {
		ph.Calibration()
	}

	fmt.Println("교정을 모두 완료하였습니다.")
}
