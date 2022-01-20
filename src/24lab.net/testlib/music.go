/*
	사용자 정의 패키지 생성
	- 사용자 정의 라이브러리 패키지는
	- 일반적으로 폴더를 하나 만들고 그 안에 .go 파일들을 만들어 구성한다.
	- 하나의 서브 폴더안에 있는 .go파일들은 동이랗ㄴ 패키지명을 갖는다.
	- 패키지명은 해당 폴더의 이름과 같게 한다.
	- 즉, 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.

	go install
	- 사이즈가 큰 라이브러리는 go install 명령으로 라이브러리를 컴파일해서 캐시할 수 있다.
	- 이는 다음 빌드시 빌드타임을 크게 줄인다.
	- go 패키지를 빌드하고 /pkg 폴더에 설치하기 위해서 go install 수행
	- 해당 패키지 폴더 내에서 수행한다.
*/
package testlib

import "fmt"

var pop map[string]string

// 패키지 로드시 제일 먼저 자동으로 호출된다.
func init() { // 패키지 로드시 map 초기화
	println("testlib package init called!!")
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

// GetMusic: Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
	return pop[singer]
}

func getKeys() { // 내부에서만 호출 가능
	for _, val := range pop {
		fmt.Println(val)
	}
}
