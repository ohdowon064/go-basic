package main

import "24lab.net/testlib"

func main() {
	/*
		1. Go 패키지
		- 고는 패키지를 통해 모듈화, 재사용 기능을 제공한다.
		- 패키지를 사용해서 작은 단윙의 컴포넌트를 작성하고,
		- 이런 작은 패키지를 사용해서 프로그램을 개발할 것을 권장한다.

		- 고는 개발에 필요한 패키지를 표준 라이브러리로 제공한다.
			- GOROOT/pkg 안에 존재한다.
		- GOROOT 환경변수는 Go 설치 디렉토리

		2. main 패키지
		- 일반적인 패키지는 라이브러리로 사용된다.
		- 하지만 "main" 패키지는 Go Compiler에 의해 특별하게 인식된다.
			- 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.
			- 이 main 패키지 안의 main() 함수가 프로그램의 시작점(entry point)가 된다.
		- 따라서 패키지를 공유 라이브러리로 만들때는 main 패키지나 main 함수를 사용해서는 안된다.

		3. 패키지 import
		- 다른 패키지를 사용하기 위해서는 import 해야한다.
			- import "fmt"
		- 패키지를 import 할 때 고컴파일러는 GOROOT 또는 GOPATH 환경변수를 검색한다.
		- 검색기준
			- 표준패키지 -> GOROOT/pkg에서 검색
			- 사용자패키지/3rd Party 패키지 -> GOPATH/pkg에서 검색
		- GOROOT는 Go 설치지 자동 설정
		- GOPATH는 사용자가 지정해야한다.

		4. 패키지 scope
		- 패키지 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재한다.
		- 이름(identifier)이 대문자로 시작하면 public -> 패키지 외부에서 호출가능
		- 이름이 소문자로 시작하면 non-public -> 패키지 내부에서만 사용가능

		5. 패키지 init 함수와 alias
		- init()함수는 패키지 실행시 처음으로 호출되는 함수이다.
		- 패키지가 로드되면 실행되는 함수로 호출없이 자동으로 호출된다.

		- import 하면서 단지 해당 패키지 안의 init()함수만을 호출하고자 하는 경우
		- 패키지 import 시 _ alias를 지정한다.
		예시. other/xlib 패키지를 호출하면서 _ alias를 지정
		package main
		import _ "other/xlib"

		- 만약 패키지 이름이 동일하지만 다른 버전, 다른 위치에서 로딩할 경우
		- 패키지 alias로 구분할 수 있다.
		import (
			mongo "other/mongo/db"
			mysql "other/mysql/db"
		)

		func main() {
			mongodb := mongo.Get()
			mydb := mysql.Get()
		}
	*/

	song := testlib.GetMusic("Alicia Keys")
	println(song)
}
