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

	fmt.Println("교정을 모두 완료하였습니다.")
}

func calibrationEC(port serial.Port) {

	var sensorType string
	var lowPoint string
	var highPoint string
	var value string

	fmt.Println("EC센서 교정을 시작합니다.")
	fmt.Println("EC센서 프로프 타입을 입력해주세요. 종류)0.1, 1.0, 10 ex)1.0")

	_, _ = fmt.Scanln(&sensorType)
	fmt.Printf("K%s로 센서가 설정되었습니다.\n")
	switch sensorType {
	case "0.1":
		lowPoint = "1413"
		highPoint = "12880"
		break
	case "1.0":
		lowPoint = "1413"
		highPoint = "12880"
		break
	case "10":
		lowPoint = "12880"
		highPoint = "150000"
	}
	fmt.Printf("Low Point: %s, High Point: %s로 설정되었습니다.\n", lowPoint, highPoint)


	fmt.Println("건조상태 설정을 시작합니다. 센서가 건조한 상태가 아닐 경우, 건조한 상태를 만들어주세요.")
	fmt.Println("건조한 상태일 경우 엔터를 눌러주세요.")
	_, _ = fmt.Scanln(&value)
	result, _ := Write(port, "Cal,dry")
	if strings.Contains(result, "OK") {
		fmt.Println("건조상태에 대한 교정이 완료되었습니다.")
	} else {
		fmt.Println("건조상태 교정 실패.")
		fmt.Println("건조상태 교정 실패.")
	}

	fmt.Println("로우 포인트 교정을 시작합니다.")
	fmt.Printf("%s 용액을 25도의 용액으로 준비한 후 센서를 담그십시오. 완료되면 엔터를 눌러주세요.\n", lowPoint)
	_, _ = fmt.Scanln(&value)
	result, _ = Write(port, "Cal,low," + lowPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("로우 포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("로우 포인트 교정 실패")
		fmt.Println("로우 포인트 교정 실패")
	}

	fmt.Println("추가적으로, 싱글 포인트 교정을 시작시작하시겠습니까? 해당 사항은 옵션입니다. 교정을 원할 경우 Y를 눌러주세요.")
	_, _ = fmt.Scanln(&value)
	if strings.ToUpper(value) == "Y" {
		fmt.Printf("싱글 포인트 교정을 시작합니다. %d 용액에 담겨있는 상태로 유지 후, 엔터를 눌러주세요.")
		_, _ = fmt.Scanln(&value)
		result, _ = Write(port, "Cal," + lowPoint)
		if strings.Contains(result, "OK") {
			fmt.Println("싱글 포인트 교정이 완료되었습니다.")
		} else {
			fmt.Println("싱글 포인트 교정 실패")
			fmt.Println("싱글 포인트 교정 실패")
		}
	}

	fmt.Println("하이 포인트 교정을 시작합니다.")
	fmt.Printf("%s 용액을 25도의 용액으로 준비한 후 센서를 담그십시오. 완료되면 엔터를 눌러주세요.\n", highPoint)
	_, _ = fmt.Scanln(&value)
	result, _ = Write(port, "Cal," + lowPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("하이 포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("하이 포인트 교정 실패")
		fmt.Println("하이 포인트 교정 실패")
	}

}

func calibrationDO(port serial.Port) {

	var value string

	fmt.Println("DO센서 교정을 시작합니다.")

	fmt.Println("DO센서 싱글포인트 교정을 시작합니다. DO센서의 마개를 분리 후 약 30초정도 공기에 노출시킨 후 엔터를 누르세요.")
	_, _ = fmt.Scanln(&value)
	result, _ := Write(port, "Cal")
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
		return
	}

	fmt.Println("듀얼포인트 교정을 시작합니다.")
	fmt.Println("DO Zero 용액에 센서를 담근 후, 1분 30초 뒤 엔터를 누르세요.")
	_, _ = fmt.Scanln(&value)
	result, _ = Write(port, "Cal,0")
	if strings.Contains(result, "OK") {
		fmt.Println("듀얼포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("듀얼포인트 교정 실패")
		fmt.Println("듀얼포인트 교정 실패")
	}

}

func calibrationPH(port serial.Port) {

	var value string

	fmt.Println("pH센서 교정을 시작합니다.")

	fmt.Println("pH센서 싱글(미드)포인트 교정을 시작합니다.")
	fmt.Println("마개를 연 후, 흐르는 물에 pH센서를 세척합니다.")
	fmt.Println("pH 센서 캘리브레이션용 노란색 용기의 용액을 컵에 따른 후, 센서를 담그고 1~2분정도 기다립니다. 시간이 지나면 엔터를 누르십시오.")
	_, _ = fmt.Scanln(&value)
	result, _ := Write(port, "Cal,mid,7.00")
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
	result, _ = Write(port, "Cal,low,4.00")
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
	result, _ = Write(port, "Cal,high,10.00")
	if strings.Contains(result, "OK") {
		fmt.Println("쓰리(하이)포인트 교정이 되었습니다.")
	} else {
		fmt.Println("쓰리(하이)포인트 교정 실패")
		fmt.Println("쓰리(하이)포인트 교정 실패")
	}
}

func calibrationRTD(port serial.Port) {

	var value string

	fmt.Println("온도센서 교정을 시작합니다.")
	fmt.Println("현재 온도센서가 담겨있는 물의 섭시온도는 몇도인가요? ex) 25.5")
	fmt.Print("온도: ")
	fmt.Scanln(&value)
	fmt.Printf("%v℃로 교정시작합니다.\n", value)
	result, _ := Write(port, "Cal,"+value)
	if strings.Contains(result, "OK") {
		println("온도센서 교정이 완료되었습니다.")
	}
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
