package main

import (
	"fmt"
	"time"
)

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
        - 이렇게 고채널은 송신자, 수신자가 서로를 기다리는 속성이 있다.
        - 이를 이용하여 고루틴이 끝날때까지 기다리는 기능을 구현할 수 있다.
    */

    done := make(chan bool)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    }()
    
    // 메인루틴은 고루틴이 끝날 때까지 대기
    // 익명함수 고루틴에서 작업이 끝난 후, done채널에 true를 보내면
    // 수신자 메인루틴은 이를 받고 프로그램을 종료한다.
    <- done


    /*
        2. Go 채널 버퍼링
        - Go 채널은 2가지가 있다.
            1) Unburffered Channel
            2) Buffered Channel
        
        - 위 예제는 Unbuffered Channel로 하나의 수신자가 데이터를 받을 때까지
        - 송신자가 데이터를 보내는 채널에 묶여있게 된다.

        - 하지만 Buffered Channel을 사용하면 비록 수신자가 받을 준비가 안돼도
        - 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있다.
        - 버퍼 채널은 make(chan type, N)함수를 통해 생성한다.
            - N는 사용할 버퍼 개수를 의미한다.
            - make(chan int, 10)는 10개의 정수형을 갖는 버퍼채널이다.
        
        - 고루틴 데드락 예제
        - (fatal error: all goroutines are asleep - deadlock!)
        func main() {
            c := make(chan int)

            // 현재 이곳은 메인루틴이다.
            // 메인루틴에서 채널로 1을 전송
            c <- 1 // 수신 루틴이 없으므로 데드락
            
            fmt.Println(<-c) // 코멘트(주석처리)해도 데드락, 별도의 고루틴이 없기때문
        }
        // 위 코드는 메인루틴에서 채널에 1을 보내면서 상대 수신자를 기다리는데
        // 이 채널을 받는 수신자 고루틴이 없다.

        - 하지만 버퍼채널을 사용하면 당장 수신자가 없더라도
        - 최대 버퍼수까지 데이터를 보낼 수 있다.
    */

    bufferedCh := make(chan int, 1)
    
    // 수신자가 없더라도 보낼 수 있다.
    bufferedCh <- 101
    fmt.Println(<-bufferedCh)


//     done1 := make(chan bool)
//     done2 := make(chan bool)

//     go run1(done1)
//     go run2(done2)

// EXIT:
//     for {
//         select {
//         case <-done1:
//             println("run1 완료")
//         case <-done2:
//             println("run2 완료")
//             break EXIT
//         }
//     }
}

func run1(done chan bool) {
    time.Sleep(3 * time.Second)
    done <- true
}

func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}