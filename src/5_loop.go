// 키워드: for

package main

func main() {
	/*
		1. for
		- go에서 반복문은 오직 for만 있다.
		- for 초기값; 조건식; 증감 {...}
		- 초기값, 조건, 증감은 생략가능
	*/

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	/*
		2. 조건식만 쓰는 for 루프
		- 초기값, 증감식 생략하고 조건식만 사용가능
		- for를 while처럼 사용가능
	*/
	n := 1
	for n < 100 {
		n *= 2
		if n > 90 {
			println("break!!")
			break
		}
	}
	println(n)

	// 3. for 무한루프
	// 초기값; 조건식; 증감 전부 생략
	// ctrl + c로 종료
	for {
		println("Infinite loop")
		break
	}

	// 4. for range문
	// foreach와 유사하다.
	// 컬렉션에서 요소 한개씩 가져와 차례로 실행한다.
	// for 인덱스, 요소 := range 컬렉션
	names := []string{"홍길동", "이순신", "강감찬"}
	for i, name := range names {
		println(i, name)
	}

	// 5. break, continue, goto
	// break: 경우에 따라 for루프 종료
	// continue: 다음 순회로 이동
	// goto: for 루프와 상관없이 기타 임의의 코드로 이동

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}

	if a == 11 {
		goto GOGOGO
	}
	println(a)

GOGOGO: // 반드시 goto 구문으로 이동안해도 코드실행하면서 있으면 실행시킨다.
	println("Go END", a) // 즉, 단지 GOGOGO label부터 실행시킨다는 뜻
	for a > 10 {
		a--
		goto GOGOGO
	}
	// break label
	i := 0
L1:
	for {
		println("몇번?")
		if i == 0 {
			break L1 // L1 루프 종룐
		}
	}
	println("OK")
}
