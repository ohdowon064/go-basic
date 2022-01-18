package main

func main() {
	/*
		1. 클로저(Closure)
		- 함수 바깥에 있는 변수를 참조하는 함수값을 말한다.
		- 이때 함수는 바깥의 변수를 함수 안으로 끌어들인 듯이 변수를 읽고 쓸 수 있다.

		예시
		- nextValue() 함수는 "int를 리턴하는 익명함수(func() int)"를 리턴하는 함수
		- go에서 함수는 일급함수이기 때문에 리턴값으로 사용될 수 있다.
		- 그런데 여기서 익명함수가 함수 바깥에 있는 nextValue의 변수 i를 참조한다.
		- 익명함수가 i를 갖는 것이 아니기 때문에 i는 값을 계속해서 유지하고 1씩 증가시킨다.
	*/
	next := nextValue()
	// var next func() int = nextValue()

	for i := 0; i < 10; i++ {
		println(next())
	}

	anotherNext := nextValue() // nextValue()를 새로 생성했으므로 i는 0부터 다시 시작한다.
	println("새로 시작하는 i")
	for i := 0; i < 5; i++ {
		println(anotherNext())
	}
}

func nextValue() func() int { // return type은 int가 아니라 func() int이다.
	i := 0
	return func() int { // 익명함수의 return type이 int이다.
		i++
		return i
	}
}
