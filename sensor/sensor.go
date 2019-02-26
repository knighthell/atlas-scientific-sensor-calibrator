package sensor

import (
	"fmt"
	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
	"log"
	"strings"
)

type Sensor struct {
	Port serial.Port
	Type string
}

type Sensors struct {
	RTD RTDSensor
	EC  ECSensor
	DO  DOSensor
	PH  PHSensor
}

func GetSensors(sensorType []string) (*RTDSensor, *ECSensor, *DOSensor, *PHSensor, error) {

	var rtd RTDSensor
	var ph PHSensor
	var do DOSensor
	var ec ECSensor

	fmt.Println("연결된 센서를 찾고있습니다...")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("연결된 시리얼포트를 찾을 수 없습니다.")
	}

	for _, portInfo := range ports {

		mode := &serial.Mode{
			BaudRate: 9600,
		}

		if portInfo.PID == "" {
			continue
		}

		port, _ := serial.Open(portInfo.Name, mode)
		result, _ := write(port, "C,0\n\r")
		if strings.Contains(result, "ER") {
			fmt.Printf("Error Stop Continious Reading.\n")
		}

		result, _ = write(port, "i\n\r")
		sensorInfo := strings.Split(result, ",")
		switch sensorInfo[1] {
		case "RTD":
			rtd.Port = port
			break
		case "pH":
			ph.Port = port
			break
		case "DO":
			do.Port = port
			break
		case "EC":
			ph.Port = port
			break
		}
	}

	return &rtd, &ec, &do, &ph, nil
}

func read(port serial.Port) (string, error) {
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

func write(port serial.Port, data string) (string, error) {

	_, err := port.Write([]byte(data))
	if err != nil {
		return "", err
	}

	result, err := read(port)
	if err != nil {
		return "", err
	}

	return result, nil
}
