package sensor

import (
	"fmt"
	"strings"
)

type ECSensor struct {
	Sensor
	ProbeType string
	LowPoint  string
	HighPoint string
}

func (s *ECSensor) Calibration() {

	var value string

	fmt.Println("EC센서 교정을 시작합니다.")
	fmt.Println("EC센서 프로프 타입을 입력해주세요. 종류)0.1, 1.0, 10 ex)1.0")

	_, _ = fmt.Scanln(&s.ProbeType)
	fmt.Printf("K%s로 센서가 설정되었습니다.\n")
	switch s.ProbeType {
	case "0.1":
		s.LowPoint = "1413"
		s.HighPoint = "12880"
		break
	case "1.0":
		s.LowPoint = "1413"
		s.HighPoint = "12880"
		break
	case "10":
		s.LowPoint = "12880"
		s.HighPoint = "150000"
	}
	fmt.Printf("Low Point: %s, High Point: %s로 설정되었습니다.\n", s.LowPoint, s.HighPoint)

	fmt.Println("건조상태 설정을 시작합니다. 센서가 건조한 상태가 아닐 경우, 건조한 상태를 만들어주세요.")
	fmt.Println("건조한 상태일 경우 엔터를 눌러주세요.")
	_, _ = fmt.Scanln(&value)
	result, _ := write(s.Port, "Cal,dry")
	if strings.Contains(result, "OK") {
		fmt.Println("건조상태에 대한 교정이 완료되었습니다.")
	} else {
		fmt.Println("건조상태 교정 실패.")
		fmt.Println("건조상태 교정 실패.")
	}

	fmt.Println("로우 포인트 교정을 시작합니다.")
	fmt.Printf("%s 용액을 25도의 용액으로 준비한 후 센서를 담그십시오. 완료되면 엔터를 눌러주세요.\n", s.LowPoint)
	_, _ = fmt.Scanln(&value)
	result, _ = write(s.Port, "Cal,low,"+s.LowPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("로우 포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("로우 포인트 교정 실패")
		fmt.Println("로우 포인트 교정 실패")
	}

	fmt.Println("추가적으로, 싱글 포인트 교정을 시작시작하시겠습니까? 해당 사항은 옵션입니다. 교정을 원할 경우 Y를 눌러주세요.")
	_, _ = fmt.Scanln(&value)
	if strings.ToUpper(value) == "Y" {
		fmt.Printf("싱글 포인트 교정을 시작합니다. %d 용액에 담겨있는 상태로 유지 후, 엔터를 눌러주세요.\n", s.LowPoint)
		_, _ = fmt.Scanln(&value)
		result, _ = write(s.Port, "Cal,high,"+s.LowPoint)
		if strings.Contains(result, "OK") {
			fmt.Println("싱글 포인트 교정이 완료되었습니다.")
		} else {
			fmt.Println("싱글 포인트 교정 실패")
			fmt.Println("싱글 포인트 교정 실패")
		}
	}

	fmt.Println("하이 포인트 교정을 시작합니다.")
	fmt.Printf("%s 용액을 25도의 용액으로 준비한 후 센서를 담그십시오. 완료되면 엔터를 눌러주세요.\n", s.HighPoint)
	_, _ = fmt.Scanln(&value)
	result, _ = write(s.Port, "Cal,high,"+s.HighPoint)
	if strings.Contains(result, "OK") {
		fmt.Println("하이 포인트 교정이 완료되었습니다.")
	} else {
		fmt.Println("하이 포인트 교정 실패")
		fmt.Println("하이 포인트 교정 실패")
	}
}
