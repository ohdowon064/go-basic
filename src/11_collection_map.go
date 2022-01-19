package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		1. map
		- map은 key-value 해시테이블을 구현한 자료구조이다.
		- 내장 자료구조
		- map[keyType]valueType
			ex. 키: 정수, 값: 문자열 맵변수 -> var idMap map[int]string

		- 이렇게 생성한 map은 reference 타입이므로 nil값을 갖으며, nil map이라 한다.
			- nil map에는 데이터를 쓸 수 없다.
		- make() 함수로 초기화할 수 있다.
	*/
	var idMap map[int]string // nil map
	if idMap == nil {
		println("idMap은 nil map입니다.")
	}

	idMap = make(map[int]string) // idMap -> map -> runtime.hmap
	// idMap은 "runtime.hmap을 가리키는 map"을 가리키는 맵 변수이다.

	/*
		- make() 함수는 해시테이블을 메모리에 생성하고 해당 메모리를 가리키는 map value를 반환한다.
		- map value는 내부적으로 runtime.hmap 구조체를 가리키는 포인터이다.
		- 따라서 위 변수 idMap은 이 해시테이블을 가리키는 map을 가리키게 된다.

		- map은 make()함수말고 리터럴(literal)을 사용해 초기화가능
		- 리터럴 초기화는 다음과 같다.
	*/
	// 리터럴 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook", // 끝에 , 필수
	}

	fmt.Println(tickers)

	/*
		2. map 사용
		- 처음 map이 make()로 초기화될 때는 아무 데이터가 없다.
		- 데이터 추가는 맵변수[키] = 값 방식으로 추가/수정할 수 있다.
		- 값은 맵변수[키]로 읽을 수 있다.
			- 키값이 없는 경우
			- 값이 레퍼런스타입 -> nil반환
			- 값이 value 타입 -> zeor반환

		- 특정 키 삭제는 delete(맵변수, 키)함수를 사용한다.
	*/

	var m map[int]string
	m = make(map[int]string)
	fmt.Println(m)

	// 추가/갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"
	fmt.Println(m)

	// 값 읽기
	str := m[134]
	println(str)

	noData, exists := m[999]
	println("noData = ", noData) // string의 zero value인 ""(빈 문자열)이 할당되었다.
	println(exists)
	fmt.Println(reflect.TypeOf(noData))

	// 삭제
	delete(m, 777)
	fmt.Println(m)

	/*
		3. map 키체크
		- 어떤 경우 map은 특정 키가 존재하는 체크할 필요가 있다.
		- 맵변수[키]는 2개의 리턴값을 리턴한다.
			1) 키에 상응하는 값
			2) 키가 존재여부의 bool 값
	*/

	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker", val)
	}

	/*
		4. for 루프를 사용한 map 열거
		- for range 루프
		- map의 key, value를 리턴한다.
	*/
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range로 모든 맵 요소 출력
	// map은 unordered 자료구조로 순서는 무작위이다. -> 실행할 때마다 순서바뀜
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
