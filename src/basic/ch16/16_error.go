package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		1. Go 에러
		- Go는 내장타입으로 error라는 interface 타입을 갖는다.
		- Go 에러는 이 error interface를 통해서 주고받는다.
		- error interface는 하나의 메서드를 갖는다.
		- 개발자는 이 인터페이스를 구현하는 커스텀 에러타입을 만들 수 있다.

		type error interface {
			Error() string
		}

		2. Go 에러처리
		- Go 함수가 결과와 에러를 함께 리턴 -> 에러가 nil인지 체크해서 에러 유무 체크가능
		- 예시)
			os.Open() 함수
			func Open(name string) (file *File, err error)
			- 리턴값: File포인터, error 인터페이스
			- 두번째 error체크 후 nil이면 에러없음, nil아니면 err.Error()로 에러확인가능
	*/
	f, err := os.Open("/Users/odowon/tmp/hello.txt")
	if err != nil {
		log.Fatal(err.Error()) // log.Fatal(): 메시지 출력 후 os.Exit(1)을 호출하여 프로그램 종료
	}
	println(f.Name())
	f.Close()

	/*
		- 또 다른 방식으로 error의 Type을 체크해서 에러 타입별로 별도의 에러처리 가능
		- switch에서 변수.(type)으로 타입체크가능
		- default는 에러타입이 nil로 에러가 없는 경우
		- MyError로 커스텀 에러 또는 일반 에러 케이스 처리
		- case error는 모든 에러는 error 인터페이스를 구현하므로 모든 에러에 해당한다.

		_, err := otherFunc()
		switch err.(type) {
		default: // no error
			println("ok")
		case MyError:
			log.Print("Log my error")
		case error:
			log.Fatal(err.Error())
		}
	*/

	f2, err2 := os.Open("/Users/odowon/tmp/hello.txt")
	switch err2.(type) {
	default:
		fmt.Println("ok, there is", f2.Name())
		f2.Close()
	case error:
		log.Fatal(err2.Error())
	}
}
