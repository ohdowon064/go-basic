/*
	키워드: 산술, 관계, 논리, bitwise, 할당, 포인터

	1. 산술연산자
		- 사칙연산(+, -, *, /, %), 증감(++, --)
			c = (a + b) / 5;
			i++;

	2. 관계연산자
		- 크기비교, 동일함 체크
			a == b
			a != b
			a >= b
	3. 논리연산자
		- AND, OR, NOT 표현
			a && b
			a || !(c && b)

	4. bitwise 연산자
		- 바이너리 AND, OR, XOR, shift
			c = (a & b) << 5

	5. 할당연산자
		- 할당(=)연산자 이외에 사칙연산, 비트연산 축약할당 연산자
		+=, &=, <<=
		a = 100
		a *= 10
		a >>= 2
		a |= 1 // 1과 or 연산 -> 가장 마지막 비트만 남음, 홀수(1), 짝수(0)

	6. 포인터연산자
		- C++처럼 &, *를 사용해서 참조, 역참조할 때 사용
		- go는 포인터연산자를 제공하지만 포인터 산술기능은 없다. -> 참조/역참조만 가능
		var k int = 10
		var p = &k // k의 주소를 할당
		println(*p) // p가 가리키는 주소의 실제 내용을 출력
*/
package main

import "fmt"

func main() {
	c := 10 + 20
	println(c)
	println(&c)

	k := 10
	p := &k
	fmt.Println(p, *p)

	var bytes []byte = []byte("ABC")
	fmt.Println(bytes, &bytes, &bytes[0], &bytes[1])
}
