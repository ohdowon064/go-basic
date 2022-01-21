package main

import (
	"fmt"
	"reflect"
)

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	/*
		1. 구조체(struct)
		- Go에서 struct는 custom data type을 표현
			- 필드들의 집합체/컨테이너
			- 필드데이터만 가지며, 메서드(행위표현)는 갖지않는다.
		- 전통적인 OOP가 갖는 클래스, 객체, 상속 개념이 없다.
		- struct는 필드만을 가지며, 메서드는 따로 정의한다.

		2. struct 선언
		- struct 정의는 custom type을 정의하는 type 키워드를 사용한다.
		- 예시인 person 구조체를 패키지외부에서 사용하기 위해서는 Person(앞글자 대문자)으로 정의하면 된다.
	*/

	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	/*
		3. struct 객체 생성
		- person{}으로 빈객체 할당 후 필드값 채워넣기
			- 필드 액세스는 .을 사용한다.
		- 초기값으로 할당하기
			- person{name: "bob", age: 20}
			- person{"bob", 20}
			- 필드 생략 시 zero value가 할당된다.
		- 내장함수 new() 사용
			- 모든 필드를 zero value로 초기화하고 객체의 포인터(*person)를 리턴한다.
			- 객체 포인터도 필드 액세스는 .을 사용하는데, 이때 자동으로 dereference가 된다.
	*/
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Tong", age: 50}
	p3 := new(person)
	p3.name = "Oh"
	fmt.Println(p1, p2, p3)
	fmt.Println(reflect.TypeOf(p1), reflect.TypeOf(p2), reflect.TypeOf(p3)) // p3가 객체포인터인 것을 확인할 수 있다.

	/*
		4. 생성자 constructor 함수
		- 언어자체에서 제공하는 생성자는 없다.
		- 구조체 필드가 사용전 초기화되어야하는 경우가 있다.
			- 예를 들어, map 타입필드인 경우 사전 초기화를 하면 매번 초기화할 필요가 없다.
		- 생성자는 struct를 리턴하는 함수로 함수 본문에서 필요한 필드를 초기화한다.
	*/
	d := newDict() // 생성자 호출
	d.data[1] = "A"
	fmt.Println(d)
}

type dict struct {
	data map[int]string
}

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // 포인터 전달
}
