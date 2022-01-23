package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func main() {
	/*
		1. Go Interface
		- 구조체 -> 필드들의 집합체
		- 인터페이스 -> 메서드들의 집합체
		- interfacesms type이 구현해야하는 메서드 원형(prototype)을 정의한다.

		2. interface 구현
		- 하나의 사용자 정의 타입이 interface를 구현하기 위해서는
			-> 단순히 해당 인터페이스가 갖는 모든 메서드를 구현하면 된다.
		- interface도 struct처럼 type문으로 정의한다.

		3. interface 사용
		1) 함수가 파라미터로 인터페이스를 받는 경우
			-> 어떤 타입이든 해당 인터페이스를 구현하면 모두 입력파라미터가 될 수 있다.
			-> 인터페이스 구현 -> 인터페이스 내 모든 메서드 구현
	*/
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)

	/*
		4. 인터페이스 타입
		- go 프로그래밍을 하다보념 빈 인터페이스(empty interface)를 자주 본다.
		- 이것은 인터페이스 타입(interface type)으로 불린다.
		- 빈 인터페이스는 interface{}와 같이 표현한다.

		- 예시) 표준패키지 함수 prototype
		func Marshal(v interface{}) ([]byte, error);
		func Println(a ...interface{}) (n int, err error);

		- empty interface는 메서드가 전혀없는 인터페이스이다.
			-> 모든 타입은 적어도 0개 메서드를 구현한다. -> 당연한 말
			-> 따라서 Go에서 모든 type을 나타내기 위해 empty interface를 사용한다.
			-> 즉, empty interface는 모든 타입을 담을 수 있는 컨테이너이다.
				-> Dynamic Type
			-> C#, java의 object, C/C++의 void*
	*/
	var x interface{} // empty interface로 모든 타입가능
	x = 1
	x = "Tom"
	printIt(x)

	/*
		5. type assertion
		- x.(T)
			- interface type x와 어떤 Type T에 대해서,
			- x가 nil이 아니며, x는 T타입에 속한다는 것을 확인(assert)하는 것
		- x가 nil이거나 x의 타입이 T가 아니라면 런타임에러 발생
		- x가 T타입인 경우 T타입의 x를 리턴한다.
	*/
	var a interface{} = 1
	println(a)
	thisValueAndType(a)

	i := a     // a와 i는 dynamic type(interface {}, emtpy interface), 값은 1
	println(i) // 포인터주소 출력
	thisValueAndType(i)

	j := a.(int) // a는 현재 int타입이므로 int타입의 x를 반환 -> j는 int타입
	println(j)   // 1 출력
	thisValueAndType(j)

	// k := a.(string) // a는 string타입이 아니므로 런타임에러발생
	// print(k)
	// thisValueAndType(k)
	k, ok := a.(string) // flag값을 받으면 panic없이 핸들링가능
	fmt.Println(k, ok)
}

// Rect struct 정의
type Rect struct {
	width, height float64
}

// Circle struct 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // 인터페이스 메서드 호출
		fmt.Println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v) // Tom
}

func thisValueAndType(this interface{}) {
	fmt.Println(this, reflect.TypeOf(this))
}
