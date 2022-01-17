package main

func main() {
	/*
		1. 익명함수
		- 함수명을 갖지않는 함수
		- 익명함수는 변수에 할당되거나, 파라미터에 직접 정의되어 사용된다.
	*/

	// var sum func(n ...int) int = func(n ...int) int {}도 가능
	// 즉, func(n ...int) int 이것이 sum의 타입이 되는 것이다.
	sum := func(n ...int) int { // 익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) // 익명함수 호출
	println(result)

	/*
		2. 일급함수
		- 함수를 변수처럼 다루는 개념
		- go에서 함수는 일급함수로 기본타입과 동일하게 취급된다.
		- 따라서 함수자체가 파라미터로 전달, 리턴값으로 반환되는데 사용할 수 있다.
	*/

	// add에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// 익명함수를 add 변수에 넣고, add 함수를 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 파라미터에 익명함수를 정의해서 전달
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

	/*
		3. type문을 사용한 함수 원형 정의
		- type은 구조체(struct), 인터페이스(interface)등 CustomType(User Defined Type)을 정의하기 위해 사용
			(Go는 Class 키워드가 없다.)
		- 위 경우 func(x int, y int) int 함수 원형이 반복된다.
		- 이를 type문으로 정의함으로써 간단히 표현하고 재사용성을 향상된다.
		- 함수 원형을 정의하고 함수를 메서드에 파라미터로 전달하고 함수를 리턴받는 기능을 델리게이트(delegate)라 한다.
			- Go는 이러한 delegate기능을 제공한다.
	*/
	ret := calc_with_type(func(x int, y int) int { return x * y }, 10, 50)
	println(ret)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 원형 정의
type calculator func(int, int) int

func calc_with_type(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
