package main

import "fmt"

func main() {
	/*
		http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice
		1. 슬라이스(Slice)  **array보다 slice를 사용하자!
		- go는 배열 크기를 동적으로 증가시키거나, 부분 배열 발췌기능이 없다.
		- size가 정해지지 않은 array를 slice라 한다.
			- 그러나 array와 slice는 메모리 할당 방식이 다르기 때문에 완전히 다른 타입이다.
			- go에서 배열은 거의 안쓴다.
			- 슬라이스에서 배열의 모든 기능이 가능해서 배열은 잘 안 쓴다.
			- effective go에서도 array는 있지만 슬라이스를 사용하라 권한다.
		- 내부적으로 배열에 기초하여 편리하게 사용하기 위해 구현된 자료구조
		- 고정된 크기를 미리 지정하지 않아도 되고, 차후 크기를 동적으로 변경가능
		- 부분 배열 발췌인 파이썬의 슬라이싱과 같은 기능을 지원한다.
		- var v []T : 배열과 달리 크기를 지정하지 않는다. 그냥 []
			ex) 정수형 slice 변수 a: var a []int
	*/
	var a []int        // slice 변수 선언
	a = []int{1, 2, 3} // 슬라이스에 리터럴 값 지정
	a[1] = 10
	fmt.Println(a)

	/*
		2. Slice 생성 방법
		- var v []T 말고도 내장함수 make()를 사용하여 슬라이스 변수 생성 가능
		- make([]Type, length, capacity)
		- 길이(length): 슬라이스의 길이
		- 용량(capacity): 내부 배열의 최대 길이
			- len(), cap()을 통해 확인 가능
		- 이후 모든 요소가 zero value인 슬라이스가 만들어진다.
		- capacity 생략 시 length와 같은 값을 갖는다.
	*/

	s := make([]int, 5, 10)
	fmt.Println("zero value init", s)

	// 길이와 용량을 지정하지 않으면 길이, 용량이 0인 슬라이스를 생성한다.
	// 이를 nil slice라 한다.
	var nilSlice []int
	if nilSlice == nil {
		println("Nil Slice")
	}
	println(len(nilSlice), cap(nilSlice))

	/*
		3. 부분 슬라이스(Sub Slice)
		- 슬리아스[시작인덱스:끝인덱스]
		- 마지막 인덱스는 포함하지 않는다.
			- 시작-inclusive, 끝-exclusive
			- 파이썬과 동일

		- 시작인덱스 생략 -> 처음부터
		- 마지막 인덱스 생략 -> 끝까지
	*/
	sub := []int{0, 1, 2, 3, 4, 5}
	println(len(sub), cap(sub))
	subSlice := sub[2:5]
	fmt.Println(subSlice)

	fmt.Println(sub[1:])
	fmt.Println(sub[:])
	fmt.Println(sub[:3])

	/*
		4. 슬라이스 추가, 병합(append)과 복사(copy)
		- 배열은 고정크기 이상의 데이터 추가 불가능
		- 슬라이스는 자유롭게 요소 추가가능
		- 슬라이스 요소 추가는 go 내장함수 append() 사용
		- append(sliceObj, ...추가요소)
		- 추가요소는 가변인자이다. 여러개 가능
	*/
	println("슬라이스 추가, 병합")
	appendSlice := []int{0, 1}
	println(len(appendSlice), cap(appendSlice))
	fmt.Println(&appendSlice[0])

	// 하나 확장
	appendSlice = append(appendSlice, 2) // 새로운 슬라이스를 생성해서 반환하는 방식 -> 용량 초과시
	println(len(appendSlice), cap(appendSlice))
	fmt.Println(&appendSlice[0]) // 주소값이 다른 것을 확인할 수 있다.

	// 복수 확장
	appendSlice = append(appendSlice, 3, 4, 5)
	fmt.Println(appendSlice)

	/*
		5. append 내부 동작방식
		1) 슬라이스 용량이 남은 경우
			- 슬라이스 길이 변경 후 데이터 추가
		2) 슬라이스 용량이 초과된 경우
			- 현재 용량 2배의 새로운 Underlying array 생성
			- 기존 배열 값 복제 후 다시 슬라이스 할당
	*/
	sliceA := make([]int, 0, 3)
	sliceInfo(sliceA, "sliceA")

	// 하나 씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		sliceInfo(sliceA, "sliceA")
	}

	// 슬라이스 병합
	sliceX := []int{1, 2, 3}
	sliceY := []int{4, 5, 6}
	sliceX = append(sliceX, sliceY...)
	// 합쳐지는 slice는 뒤에 ...을 붙인다.
	// 즉, sliceY가 4, 5, 6으로 치환되는 것이다.
	fmt.Println(sliceX)

	/*
		6. slice 복사 - copy()
		- 슬라이스는 실제 배열을 가리키는 포인터 정보만 갖는다.
		- 복사는 source 슬라이스가 갖는 배열의 데이터를 target 슬라이스가 갖는 배열로 복제하는 것
	*/
	source := []int{1, 2, 3}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source) // source 요소를 target에 복사
	sliceInfo(source, "source")
	sliceInfo(target, "target")

	/*
		7. 슬라이스의 내부구조
		- 슬라이스는 내부적으로 사용되는 배열의 부분영역인 세그먼트에 대한 메타정보를 갖는다.
		- 슬라이스는 3개의 필드로 구성된다.
			1) 내부 배열에 대한 포인터
			2) 세그먼트 길이(length)
			3) 세그먼트 최대 용량(capacity)

		- 처음 슬라이스 생성 시 길이, 용량이 지정된 경우
			- 용량만큼의 배열을 생성하고 슬라이스 첫번째 필드에 해당 배열 처음 메모리 위치를 저장한다.
			- 두번째 길이 필드는 첫 배열 요소로부터의 지정된 길이를 갖는다.
			- 세번째 용량 필드는 전체 배열의 크기를 갖는다.

	*/
	exam := make([]int, 6, 6)
	for i := 0; i < len(exam); i++ {
		exam[i] = i
	}
	sliceInfo(exam, "exam")
	// [배열포인터, len 6, cap 6] 슬라이스 -> [0, 1, 2, 3, 4, 5] 내부배열
	// 배열포인터는 내부배열[0]을 가리킨다.

	slicedExam := exam[2:5]
	sliceInfo(slicedExam, "slicedExam")
	// [배열포인터, len 3, cap 4] 슬라이스 -> [0, 1, [2, 3, 4, 5]] 내부배열
	// 배열포인터는 내부배열[2]를 가리킨다.
	// 16byte 차이나는 것을 알 수 있다.

	slicedExam2 := exam[1:4]
	sliceInfo(slicedExam2, "slicedExam2")
	// 주소 단위는 byte!
	// 8byte차이는 것을 확인할 수 있다.

	// sliceExam의 cap은 4이고
	// sliceExam2의 cap은 5이다. Why????
	// 슬라이싱했을 때 len은 자른만큼이고, cap은 자른 곳부터 끝까지이다.
	// 따라서 [2:5]에서 길이는 2, 3, 4라서 길이3이 되고
	// 용량은 2부터 끝까지. 즉, 기존용량6 - 시작인덱스2인 4가된다.

	// [1:4]는 길이가 1, 2, 3이라서 길이 3이되고,
	// 용량은 기존용량6 - 시작인덱스1인 5가 된다.
}

func sliceInfo(data []int, sliceName string) {
	println()
	fmt.Println("==", sliceName, "slice info==")
	if len(data) > 0 {
		fmt.Println("slice addr: ", &data[0])
	}

	fmt.Println("slice len: ", len(data))
	fmt.Println("slice cap: ", cap(data))
	fmt.Println("slice elem: ", data)
}
