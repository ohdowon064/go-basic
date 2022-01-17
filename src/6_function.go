package main

func main() {
	/*
		1. 함수
		- func funcName(param Type, param2 Type2, ...) returnType {
		- 호출되는 함수가 호출하는 함수보다 반드시 앞에 있을 필요없다.
	*/
	msg := "Hello"
	say(msg)

	/*
		2. Pass By Reference
		- 파라미터 전달방식은 Pass By Value vs Pass By Reference
			1) Pass By Value
			- 값이 파라미터에 복사되어 전다로딘다.
			- 함수내에서 변경해도 호출함수에서의 변수는 변함없다.
			2) Pass By Reference
			- 변수앞에 &를 붙이면 변수 주소를 표시하게 된다. -> 포인터
			- 변수는 값이 아닌 메모리 영역의 주소를 갖는다.
			- 값 할당은 *변수 = 값 이렇게 하는데 이를 Dereferencing(역참조)이라 한다.
			- reference(참조): 변수의 주소값을 갖는 것
			- dereference(역참조): 메모리 주소에 있는 값에 접근하는 것
	*/
	say_by_reference(&msg)
	println(msg)

	/*
		3. variadic function - 가변인자함수
		- 고정된 수가 아닌 여러개 파라미터를 전달 -> 가변 파라미터
		- 파이썬의 *와 유사하다.
		- ... 3개의 마침표를 타입앞에 붙여 사용한다.
			ex) ...string, ...int
	*/

	say_by_variadic_function("Hi")
	say_by_variadic_function("My", "name", "is", "dowon")

	/*
		4. 함수 리턴값
		- go는 함수 리턴값이 0, 1, 2개이상일 수 있다.
		- Named Return Parameter 기능을 제공한다.
			- 리턴되는 값들을 함수에 정의된 리턴 파라미터에 할당하는 기능이다.
			- 이때는 아무 값없는 return만 적는다.
				- 실제 return에는 아무것도 없다.
				- 하지만 리턴되는 값이 있을경우에는 빈 return을 적는다.
				- 생략하면 에러
	*/
	println(sum(1, 7, 3, 5, 9))
	println(sum_with_count(1, 7, 3, 5, 9))
	println(sum_with_named_return_parameter(1, 7, 3, 5, 9))
}

func say(msg string) string {
	println(msg)
	return msg
}

func say_by_reference(msg *string) *string {
	println(*msg)
	*msg = "Changed" // 메시지 변경, 역참조
	return msg
}

func say_by_variadic_function(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum_with_count(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 개수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum_with_named_return_parameter(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	// return count, total
	return // 그냥 리턴만 적어도 된다.
	// return 100, -999 // 타입만 맞춰서 아무값을 리턴할 수도 있다.
}
