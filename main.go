package main

import (
	"fmt"
	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
	"log"
	"strings"
)

func main() {

	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("연결된 시리얼포트를 찾을 수 없습니다.")
	}

	for _, port := range ports {
		fmt.Printf("연결할 수 있는 포트: %v, VID: %v, PID: %v\n", port.Name, port.VID, port.PID)
	}

	var portRTD serial.Port
	var portPh serial.Port
	var portDO serial.Port
	var portEC serial.Port

	for _, portInfo := range ports {

		mode := &serial.Mode{
			BaudRate: 9600,
		}

		if portInfo.PID == "" {
			continue
		}

		port, _ := serial.Open(portInfo.Name, mode)
		fmt.Printf("Opened Port: %v\n", portInfo.Name)

		fmt.Printf("Stop Continious Reading...\n")
		result, _ := Write(port, "C,0\n\r")
		if strings.Contains(result, "ER") {
			fmt.Printf("Error Stop Continious Reading.\n")
		}
		if strings.Contains(result, "OK") {
			fmt.Printf("Success Stop Continious Reading.\n")
		}

		result, _ = Write(port, "i\n\r")
		sensorInfo := strings.Split(result, ",")
		fmt.Printf("%v is %v\n", portInfo.Name, sensorInfo[1])

		switch sensorInfo[1] {
		case "RTD":
			portRTD = port
			println("RTD 맵핑완료")
			break
		case "pH":
			portPh = port
			println("pH 맵핑완료")
			break
		case "DO":
			portDO = port
			println("DO 맵핑완료")
			break
		case "EC":
			portEC = port
			println("EC 맵핑완료")
			break
		}
	}

	if portRTD != nil {
		calibrationRTD(portRTD)
	}

	if portPh != nil {
		calibrationPH(portPh)
	}

	if portDO != nil {
		calibrationDO(portDO)
	}

	if portEC != nil {
		calibrationEC(portEC)
	}

}

func calibrationEC(port serial.Port) {

	fmt.Println("EC센서 교정을 시작합니다.")
}

func calibrationDO(port serial.Port) {

	fmt.Println("DO센서 교정을 시작합니다.")
}

func calibrationPH(port serial.Port) {

	fmt.Println("pH센서 교정을 시작합니다.")

	fmt.Println("pH센서 싱글(미드)포인트 교정을 시작합니다. 시간이 경과된 후 엔터를 입력하세요.")
	var yes string
	fmt.Scanln(&yes)
	fmt.Println("싱글(미드)포인트 교정이 되었습니다.")

	fmt.Println("pH센서 투(로우)포인트 교정을 시작합니다. 시간이 경과된 후 엔터를 입력하세요.")
	fmt.Scanln(&yes)
	fmt.Println("투(로우)포인트 교정이 되었습니다.")

	fmt.Println("pH센서 쓰리(하이)포인트 교정을 시작합니다. 시간이 경과된 후 엔터를 입력하세요.")
	fmt.Scanln(&yes)
	fmt.Println("쓰리(하이)포인트 교정이 되었습니다.")
}

func calibrationRTD(port serial.Port) {

	fmt.Println("온도센서 교정을 시작합니다.")
	fmt.Println("현재 온도센서가 담겨있는 물의 섭시온도는 몇도인가요? ex) 25.5")
	var waterTemp string
	fmt.Print("온도: ")
	fmt.Scanln(&waterTemp)
	fmt.Printf("%v℃로 교정시작합니다.\n", waterTemp)
	//result, _ := Write(port, "Cal,"+waterTemp)
	//if strings.Contains(result, "OK") {
	//	println("온도센서 교정이 완료되었습니다.")
	//}
}

func Read(port serial.Port) (string, error) {
	resultBuf := make([]byte, 0, 3)
	buff := make([]byte, 20)

	for {
		n, err := port.Read(buff)
		if err != nil {
			return "", err
		}

		resultBuf = append(resultBuf, buff[:n]...)

		if buff[n-1] == 13 {
			break
		}
	}

	return string(resultBuf), nil
}

func Write(port serial.Port, data string) (string, error) {

	_, err := port.Write([]byte(data))
	if err != nil {
		return "", err
	}

	result, err := Read(port)
	if err != nil {
		return "", err
	}

	return result, nil
}
