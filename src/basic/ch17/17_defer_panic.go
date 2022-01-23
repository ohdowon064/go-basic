package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		1. 지연실행 defer
		- defer 키워드는 특정 문장, 함수를 나중에 실행하게 한다.
			-> 나중: defer를 호출하는 함수가 리턴하기 직전에
		- 일반적으로 defer는 C#, java에서 finally블럭처럼 마지막 clean-up작업을 위해 사용
		- 아래 예제는 차후 어떤 에러가 나도 항상 파일을 Close할 수 있게 한다.
	*/
	f, err := os.Open("/Users/odowon/tmp/hello.txt")
	if err != nil {
		panic(err)
	}

	// 여기서 defer를 호출하는 함수 -> main
	// 즉, main이 마지막에 파일 Close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))

	/*
		2. panic 함수
		- 내장함수 panic()은 현재 함수를 즉시 멈추고 defer를 모두 실행 후 즉시 리턴한다.
		- panic 모드 실행방식은 다시 상위함수에도 똑같이 적용되고,
		- 콜스택을 타고 올라가며 적용된다.
		- 마지막에 프로그램이 에러를 내고 종료하게 된다.
	*/

	// // 잘못된 파일명을 넣음
	// openFile("Invalid.txt")

	// // openFile() 안에서 panic이 실행되면
	// // 아래 println 문장은 실행안됨
	// println("Done")

	/*
		3. recover 함수
		- 내장함수 recover()는 panic에 의한 패닉상태를 다시 정상으로 되돌리는 함수
		- 위 예제는 println()이 호출되지 못하고 프로그램이 crash
		- recover를 사용하면 panic을 제거하고 openFileWithRecover() 다음인 println()을 호출한다.
	*/
	openFileWithRecover("Invalid.txt") // panic 발생하지만 recover

	// recover에 의해 정상으로 돌아오고 다음 실행
	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func openFileWithRecover(fn string) {
	// defer 함수. panic 호출 시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
			fmt.Println("BUT RECOVER!")
		}
	}() // ()는 호출했다는 의미, 정의와 호출을 동시에 한 것

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
