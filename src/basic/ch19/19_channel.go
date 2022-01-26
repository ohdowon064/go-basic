package main

import "time"

func main() {
    /*
        1. Go 채널
        - 고채널은 데이터를 주고받는 통로이다.
        - 채널은 반드시 make() 함수를 통해 미리 생성되어야 한다.
        - 데이터 전송: 채널<-
        - 데이터 수신: <-채널
        - 채널은 일반적으로 여러 고루틴 사이에 데이터를 주고 받는데 사용된다.
            - 상대편이 준비될 때까지 대기함으로써 별도의 lock을 걸지않고 데이터를 동기화할 수 있다.
        
        예제.
        - 정수형 채널 생성 후 고루틴에서 해당 채널에 123 정수 전송
        - 전송된 데이터를 메인 루틴에서 채널로부터 123데이터를 수신
        - 채널 생성 시 make() 함수에 채널 타입을 정해야한다.
    */
    // 정수형 채널 생성
    ch := make(chan int)

    go func() {
        ch <- 123 // 채널에 123 전송
    }()

    var i int
    i = <- ch // 채널로부터 123 수신
    println(i)
    // 메인루틴은 채널로부터 데이터 수신
    // 상대편 고루틴에서 데이터 전송할 때까지 대기한다.
    // time.Sleep(), fmt.Scanf()가 없어도 고루틴이 끝날 때까지 대기한다.

    /*
    
    */
    done1 := make(chan bool)
    done2 := make(chan bool)

    go run1(done1)
    go run2(done2)

EXIT:
    for {
        select {
        case <-done1:
            println("run1 완료")
        case <-done2:
            println("run2 완료")
            break EXIT
        }
    }
}

func run1(done chan bool) {
    time.Sleep(3 * time.Second)
    done <- true
}

func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}