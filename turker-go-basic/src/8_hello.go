/*
컴매을 위한 Go언어 기초 프로그래밍 강좌 8강

작은 묶음 순서
1) library : 책, 지식과 같은 참고 지식같은 코드들로 필요한 기능들을 묶어둔 것.
2) module : 기능을 묶어두긴 했는데. 좀 더 특정 기능에 포커싱된 묶음.
3) package : 그냥 같은 묶음.
4) framework : 기능 묶음인데, 특정 절차까지 포함된 매뉴얼이 포함된 묶음.
5) engin : 종합선물세트? 기능 묶음 + 다른 프로그램들까지 포함된 툴
*/

// 패키지 선언인데
// 메인은 약속된 말로 좀 특이한 의미로 프로그램의 시작점을 의미하는 선언이다.
// exe 실행파일은 원래 HDD안에 있는데 이걸 메모리에 올려서 실행한다.
// 근데 이때 IP(instructor point)가 있어야 거기서 부터 실행할 수 있어서 그 시작점을 main이라 부른다.
// 반면에 라이브러리는 이런 시작점이 없다. 참고용도이기 때문

package main

// fmt 라이브러리를 가져온다.
// golang 표준 패키지 <-> 외부 패키지도 존재(not standard package)
import "fmt"

/*
함수 구조)
fun, 함수정의 / main, 함수이름 / (), input / (), output / {} ..

예시)
func Println(a ...any) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
*/


func main() {
	fmt.Println("Hello 월드!");
}
