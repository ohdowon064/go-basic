/*
	키워드: if, switch

	1. if
	- 조건이 참이면 {}블럭 안의 내용을 실행
	- 조건식을 ()로 감싸지 않아도 된다.
	- 시작 브레이스 {는 반드시 if문과 같은 라인에 두어야한다. -> 안지키면 에러발생
		- else if, else를 쓸 경우 종료 브레이스 }와 같은 라인에 두어야한다.
	- 조건식은 반드시 Boolena식으로 표현되어야한다.
		- 조건식에 1, 0과 같은 숫자를 쓸 수 없다.
	- 조건식 시작전 간단한 문장(optional statement) 실행가능
		- optional statement와 conditional statement는 ;로 구분
		- 이때 선언한 변수는 if-else scope안에서만 사용가능
		- 선언은 := 로만 가능

	2. switch
	- 여러 값 비교, 다수 조건식 체크에 사용가능
	- switch-case로 사용가능
	- case에 여러 값이 있을 경우 ,로 구분 가능
	- 특징: 코드로 알아보자
*/

package main

import . "fmt"

func main() {
	// a := 10
	// if a { -> non-boolean condition in if statement Error
	// 	println(a)
	// }

	// if 1 == 1{
	// 	println("abc")
	// }
	// else if 2 == 2{ -> }와 else if는 같은 라인이어야한다.
	// 	println("abc")
	// }

	i := 10
	if k := i * 2; k < 100 {
		println(k)
	}

	var name string
	var category = 1
	switch category {
	case 1:
		name = "Paper Book" // break을 안써도 다음 case로 가지않는다.
	case 2:
		name = "eBook"
	case 3, 4: // 복수 case 가능
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	var v interface{} = 100
	switch v.(type) { // switch-case type 분기처리, 타입으로 분기처리 가능
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}
	Println(v)

	check(1)
}

func grade(score int) {
	switch { // switch 뒤에 expression이 없어도 된다. true로 가정하고 case문 검사
	case score >= 90:
		println("A")
	case score >= 90:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

func check(val int) {
	switch val {
	case 1:
		Println("1 이하")
		fallthrough // 자동 break을 안하려면 fallthrough 키워드 사용
	case 2:
		Println("2 이하")
		fallthrough
	case 3:
		Println("3 이하")
		fallthrough
	default:
		Println("default 도달")
	}
}
