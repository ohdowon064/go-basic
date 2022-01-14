/*
	키워드:

	1. 변수
	- var로 선언
	- var 변수명 변수타입
		- var a int
		- var a int = 10 -> 변수 초기화
	- 선언된 변수가 사용되지 않으면 에러발생
	- var a, b, c int -> 여러개 한번에 지정가능
		- var a, b, c int = 1, 2, 3 -> 여러개 초기화
	- 초기화 안할경우 기본값 설정
		- 숫자형 -> 0
		- bool형 -> false
		- string형 -> ""(빈문자열)
	- 타입추론 제공
		- var i = 1
		- var s = "Hi"
	- := 연산자
		- 함수(func) 내부에서만 사용가능
		- var를 생략하고 변수 선언가능

	2. 상수
	- const 키워드 사용
	- const 상수명 상수타입 = 값
	- 타입추론 가능
	- 여러개 지정 가능
		const (
			Visa = "Visa" // 컴마없음, 그냥 다음줄로
			Master = "MasterCard"
			Amex = "American Express"
		)

		const (
			Apple = iota // iota는 순차부여를 위한 identifier
			Grape // 차례로 0, 1, 2 부여
			Orange
		)
	- 변수는 사용안하면 에러발생, 상수는 사용안해도 에러 발생안함

	3. Go 키워드 25개
	break        default      func         interface    select
	case         defer        go           map          struct
	chan         else         goto         package      switch
	const        fallthrough  if           range        type
	continue     for          import       return       var

*/

package main

func main() {
	var num int = 10
	println(num)

	const MONTH_DAYS = 30

}
