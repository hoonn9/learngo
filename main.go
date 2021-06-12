package main // 작성할 패키지 이름 (main 은 컴파일을 위함)
// main packpage -> main func 를 찾고 시작
import "fmt" // formatting


func main() {
	// const name string= "nico"
	// name = "test"

	var name string = "hoon"
	name = "test"

	// 축약형 (변수만 가능, Functiomn 내에서만 사용 가능)
	// 타입 추론
	name2 := "hoon"
	
	fmt.Println(name, name2)
}
