package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	/*
		1. goroutine
		- 고루틴은 go 런타임이 관리하는 lightweight(경량) 논리적(혹은 가상) 쓰레드
		- go에서 go 키워드를 사용하여 함수를 호출하면 런타임시 새로운 goroutine을 실행한다.
		- 고루틴은 비동기적 함수루틴을 실행한다.
			- 여러 코드를 동시(concurrently)에 실행하는데 사용한다.
		공식문서)
		- 고루틴은 OS쓰레드보다 훨씬 가볍게 비동기 concurrent 처리를 구현하기 위해 만든 것
		- 기본적으로 Go 런타임이 자체 관리한다.
		- 고 런타임에서 관리되는 작업단위인 고루틴들은 하나의 OS 쓰레드로도 실행되곤한다.
			- 즉, 고루틴은 OS 쓰레드와 일대일 대응되지 않는다.
			- multiplxing으로 훨씬 적은 OS 쓰레드를 사용한다.
		- 메모리 측면에서도 OS 쓰레드가 1MB 스택을 갖고, 고루틴은 초기 2KB 스택(필요시 동적으로 증가)을 갖는다.
		- 고 런타임은 고루틴을 관리하면서 Go 채널을 통해 고루틴간의 통신을 쉽게 해준다.
	*/

	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async2")

	// 3초 대기
	time.Sleep(time.Second * 3)
	/*
		1) say("Sync")를 동기적으로 호출
		2) go say("AsyncN")을 비동기적으로 3번 호출

		- say("Sync") 함수가 완전히 끝났을 때, 다음 코드로 이동
		- 다음 go say() 비동기 호출은 별도의 고루틴에서 동작
			- 메인루틴은 계속해서 다음 문장을 실행한다.
			- time.Sleep을 실행
		- 고루틴은 실행순서가 일정하지 않아 프로그램 실행마다 출력결과가 다르다.
	*/

	/*
		2. 익명함수 고루틴
		- 고루틴은 익명함수에 사용할 수 있다.
		- go 키워드 뒤에 익명함수를 정의하여 해당 익명함수를 비동기로 실행
	*/

	// WaitGroupt 생성, 2개의 고루틴을 기다림
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 고루틴
	go func() {
		defer wait.Done() // 익명함수가 끝나면 .Done()을 호출한다.
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // 고루틴이 모두 끝날 때까지 메인루틴 대기
	// 메인루틴은 고루틴인가?
	// https://stackoverflow.com/questions/53388154/is-the-main-function-runs-as-a-goroutine

	/*
		위 예제에서는 sync.WaitGroup을 사용한다.
	*/
}
