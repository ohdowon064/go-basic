/*
	키워드: data type, string, type conversion

	1. Go 데이터 타입
		1) 부울린 타입
			bool
		2) 문자열 타입
			string: immutable
		3) 정수형타입
			int, int8, int16, int32, int64
			uint, uint8, uint16, uint32, uint64, uintptr
			uintptr: uint와 크기가 동일하며 포인터를 저장할 때 사용
			int: 64bit -> int64, 32bit -> int32
			uint: 64bit -> int64, 32bit -> int32
		4) Float 및 복소수 타입
			float32 float64 complex64 complex128
		5) 기타
			byte: uint8과 동일하며 바이트코드에 사용
			rune: int32와 동일하며 유니코드 코드포인트에 사용

	2. 문자열
		- 문자열 리터럴은 back quote(``) 또는 double quote("")를 사용
		- ``문자열은 raw stirng으로 복수라인이 가능하며 이스케이프가 무시된다.
		- ""문자열은 interpreted string으로 한 라인만 가능하며 이스케이프가 사용된다.
			- +연산자가 가능하다. +쓰면 다음 줄에 작성가능

	3. type conversion
		- type(value)
		- 100을 flat으로 변환
			float32(100)
		- 문자열을 바이트배열로 변환
			[]byte("ABC")
		- go에서 타입변환은 명시적으로 해야한다.
			var i int = 100
			var u uint = i -> 런타임 에러발생 uint(i) 이렇게 명시적으로 해야함
*/

package main

import (
	. "fmt"       // static import -> 접근연산자없이 바로 사용가능
	ref "reflect" // go alias
)

func main() {
	// Raw String Literal. 복수라인
	rawLiteral := `아리랑\n
아리랑\n
		아라리요`

	// Interpreted String Literal
	interpretedLiteral := "아리랑아리랑\n아라리요" +
		"+사용해서 두 라인 사용가능"

	Println(rawLiteral)
	Println()
	Println(interpretedLiteral)

	// type conversion
	var i int = 100
	// var u uint = i // 에러발생한다. 명시적으로 형변환 해야함
	var u uint = uint(i)
	var f float32 = float32(i)
	Println(f, u) // println(f, u)

	str := "ABC"
	Println(ref.TypeOf(str)) // println(ref.TypeOf(str))
	bytes := []byte(str)
	str2 := string(bytes)
	Println(bytes, str2)

}
