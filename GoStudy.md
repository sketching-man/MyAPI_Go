# Go에 대해서 공부한 부분 적어보자.

## Go 기본 키워드

Go 프로그래밍 언어는 다음과 같은 25개의 예약 키워드들을 갖는다.  
이들 키워드들은 변수명, 상수명, 함수명 등의 이름으로 사용할 수 없다.

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Go 변수

변수는 키워드 `var` 를 사용하여 선언한다.  
`var` 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다.  

```
var a int  
var i, j, k int
var i, j, k int = 1, 2, 3
```

만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다.  
따라서 사용되지 않는 변수는 프로그램에서 삭제한다.  

Short Assignment Statement ( := ) 를 사용할 수 있다.  
즉, `var i = 1` 을 쓰는 대신 `i := 1` 이라고 `var` 를 생략하고 사용할 수 있다.  
하지만 이러한 표현은 func 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다.

## Go 상수

상수는 키워드 `const` 를 사용하여 선언한다.  
`const` 키워드 뒤에 상수명을 적고, 그 뒤에 상수타입, 그리고 상수값을 할당한다.

```
const c int = 10
const s string = "Hi"

const c = 10            // 타입 적지 않으면 자동으로 타입을 추론
const s = "Hi"          // 타입 적지 않으면 자동으로 타입을 추론

const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
```

Enum 처럼 사용이 가능하다. 이 때 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다.  
아래의 경우에 iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.

```
const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
)
```

## Go 타입

1. 부울린 타입: bool
2. 문자열 타입: string, 한번 생성되면 수정될 수 없는 Immutable 타입
3. 정수형 타입: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
4. Float 및 복소수 타입: float32 float64 complex64 complex128
5. 기타 타입: byte(uint8, 바이트 코드) rune(int32, 유니코드 코드포인트)

### string에 대하여

\` \` 로 싸인 문자열은 Raw하게 인식되며, \n 등의 기호를 해석하지 않는다.  
"" 로 싸인 문자열은 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석한다.

## Go 조건문

```
if k == 1 {
    println("One")
} else if k == 2 {  //같은 라인
    println("Two")
} else {   //같은 라인
    println("Other")
}
```

```
switch category {
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
```

단, Go는 `case`문에 expression을 정의할 수 있다(type 별 조건, 수식 결과 별 조건으로 만들기 가능).  
`break`가 필요 없다. `case` 이후 아래로 내려가게 하려면 `fallthrough` 키워드를 사용하자.

## Go 반복문

```
// 횟수 반복
for i := 1; i <= 100; i++ {
    sum += i
}
```

```
// 조건 반복 (while)
for n < 100 {
    n *= 2      
    //if n > 90 {
    //   break 
    //}
}
```

```
// 무한 루프
for {
    println("Infinite loop")
}
```

```
// foreach
names := []string{"홍길동", "이순신", "강감찬"}
for index, name := range names {
    println(index, name)
}
```

이 외에도 continue, break를 다른 언어와 똑같이 사용할 수 있다.

## Go 함수

To Be Written...

## Go 포인터

To Be Written...

### new vs make

# Go Tips

* 프로젝트 루트 디렉토리 기준으로 `mod init` 동작을 통해 `go.mod` 파일을 만들었을 때,  
타 디렉토리의 패키지들은 전체 모듈 이름을 다 적고 나서 마지막에 폴더 이름을 적어야 정상적으로 import 할 수 있음.
* 현재 패키지가 아닌 다른 패키지의 함수는 앞 글자가 대문자인 경우에만 export 됨.

# 부록. Go 관련 공부 사이트

* http://golang.site/go/basics
* https://yoongrammer.tistory.com/category/%EC%96%B8%EC%96%B4/Go%20%EC%96%B8%EC%96%B4
* https://dev-yakuza.posstree.com/ko/golang/
