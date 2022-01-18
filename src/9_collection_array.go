package main

import "fmt"

func main() {
	/*
		1. 배열(Array)
		- 배열은 동일 타입 데이터를 연속적인 메모리 공간에 순서적으로 저장하는 자료구조이다.
		- zero based array: 인덱스가 0부터 시작한다.
		- var 변수명 [배열크기]데이터타입
		- 배열크기는 데이텉입으 구성하는 한 요소이다.
			- [3]int와 [5]int는 서로 다른 타입이다.
	*/
	var a [3]int // 0으로 초기화 된다.
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a, a[2])

	/*
		2. 배열의 초기화
		- [배열크기]데이터타입 {초기값, ...}
		- 배열크기는 ...으로 생략할 수 있다.
		- [...]으로 배열크기 생략하면 초기화 요소 숫자만큼 배열크기가 정해진다.
	*/
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3} // 배열 크기 자동 설정

	println(len(a1))
	println(len(a2))

	/*
		3. 다차원 배열
		- 배열크기 부분을 복수개로 설정하여 선언한다.
	*/
	var multiArrary [3][4][5]int // 3 * 4 * 5
	multiArrary[0][1][2] = 10
	fmt.Println(multiArrary)
}
