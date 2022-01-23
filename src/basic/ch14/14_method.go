package main

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메서드
func (r Rect) area() int {
	return r.width * r.height
}

// pointer receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	/*
		1. Go method
		- go struct는 필드만을 갖고, 메서드는 별도로 분리하여 정의
		- go method는 다른 형태의 func 함수이다.
			- 일반함수: func 함수이름(파라미터) 반환타입 {}
			- 메서드: func (struct 변수명) 함수이름(파라미터) 반환타입 {}
			- struct 변수는 입력파라미터처럼 사용
		- 메서드는 객체.메서드() 방식으로 호출
	*/
	rect := Rect{10, 20}
	area := rect.area()
	println(area)

	/*
		2. value vs pointer receiver
		- value receiver는 struct 데이터를 복사하여 전달
		- pointer receiver는 struct 포인터만을 전달
		- value는 메서드내에서 struct 필드 변경해도 호출자는 변경안됨
		- pointer는 메서드내에서 필드변경이 호출자에 반영된다.
	*/
	area2 := rect.area2()
	println(area2)
	println(rect.width) // 메서드내 r.width++가 필드에 반영된다.
}
