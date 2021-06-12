package main // 작성할 패키지 이름 (main 은 컴파일을 위함)
// main packpage -> main func 를 찾고 시작
import (
	"fmt" // formatting
	"strings"
)

// 인자의 타입을 스킵하면 뒤에있는 타입으로 선언 됨
func multiply(a, b int) int {
	return a * b
}

// 여러 return 을 가질 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println(multiply(2, 2))
	
	// 무조건 모든 return 을 받아야함
	// 무시하려면 _ (컴파일러도 무시)
	// totalLength, _ := lenAndUpper("nico")

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)


	repeatMe("hoon", "kozel", "blang", "good")
}
