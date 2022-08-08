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

함수는 func 키워드를 사용하여 정의한다.  
func 뒤에 함수명을 적고 괄호 안에 그 함수에 전달하는 패러미터들을 적게 된다.  
함수 패러미터는 0개 이상 사용할 수 있는데, 각 패러미터는 패러미터명 뒤에 int, string 등의 패러미터 타입을 적어서 정의한다.  
함수의 리턴 타입은 패러미터 괄호 뒤에 적어서 정의한다.  
함수는 패키지 안에 정의되며 호출되는 함수가 호출하는 함수의 반드시 앞에 위치해야 할 필요는 없다.

```
func main() {
    msg := "Hello"
    retVal := say(msg)
}
 
func say(msg string) int {
    println(msg)
    return 0
}
```

### Pass By Value와 Pass By Reference

```
package main

func main() {
    msg := "Hello"
    say1(msg)
    println(msg) // Hello

    say2(&msg)
    println(msg) // Changed
}

func say1(msg string) {
    println(msg)
    msg = "Changed" // 메시지 변경 시도
}
 
func say2(msg *string) {
    println(*msg)
    *msg = "Changed" // 메시지 변경 시도
}
```

위의 예제에서 say1()의 경우 msg의 값 "Hello" 문자열을 인자로 넘겼다.  
이 경우 Pass By Value 형태로, 문자열 데이터가 복사되어 함수 say1()에 전달된다.  
즉, 만약 패러미터 msg의 값이 say1() 함수 내에서 변경된다 하더라도 호출함수 main()에서의 msg 변수는 변함이 없다.  
이렇게 Pass By Value의 경우 함수 안에서의 값 변경이 함수 밖에서의 변경과는 무관하다.  

위의 예제에서 say2()의 경우 msg 변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시하게 된다(이는 C 계열 언어와 동일하다).  
포인터라 불리우는 이 용법을 사용하면 함수에 msg 변수의 값을 복사하지 않고 msg 변수의 주소를 전달하게 된다.  
즉, 패러미터 msg의 값을 *msg 형태로 say2() 함수 내에서 변경하면 호출함수 main()에서의 msg 변수도 함께 변경된다.  
이렇게 Pass By Reference의 경우 함수 안에서의 값 변경이 함수 밖에서의 값 변경과 관련이 있다.

### 가변인자 함수

함수에 고정된 수의 패러미터들을 전달하지 않고 다양한 숫자의 패러미터를 전달하고자 할 때 가변 패러미터를 나타내는 ... (3개의 마침표)을 사용한다.  
단, 여러 패러미터의 타입은 동일해야 한다.

```
package main

func main() {   
    say("This", "is", "a", "book")
    say("Hi")
}

func say(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}
```

### 복수 리턴 함수

함수는 리턴값이 없을 수도, 리턴값이 하나 일 수도, 또는 리턴값이 복수 개일 수도 있다.  
C 언어에서 void 혹은 하나의 값만을 리턴하는 것과 대조적으로 Go 언어는 복수개의 값을 리턴할 수 있다.  

```
package main
 
func main() {
    total := sum1(1, 7, 3, 5, 9)
}
 
func sum1(nums ...int) int {
    // ...
    return s
}

func sum2(nums ...int) (int, int) {
    // ...
    return count, s
}

func sum3(nums ...int) (count int, total int) {
    // count 변수와 total 변수 있음
    return
}
```

### 익명 함수

`func` 키워드로 함수 전체를 바로 정의해서 변수에 할당하거나 다른 함수의 패러미터에 직접 정의되어 사용한다.

```
package main
 
func main() {
    sum := func(n ...int) int { //익명함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
 
    result := sum(1, 2, 3, 4, 5) //익명함수 호출
    println(result)
}
```

### 일급(First-class) 함수

함수를 다른 변수와 동일하게 다루는 언어는 일급 함수를 가졌다고 표현한다.  
예를 들어, 일급 함수를 가진 언어에서는 함수를 다른 함수에 인수로 제공하거나, 함수가 함수를 반환할 수 있으며, 변수에도 할당할 수 있다.  
(출처: https://developer.mozilla.org/ko/docs/Glossary/First-class_Function)  

```
package main
 
func main() {
    //변수 add 에 익명 함수 할당
    add := func(i int, j int) int {
        return i + j
    }
 
    // add 함수 전달
    r1 := calc(add, 10, 20)
    println(r1)
 
    // 직접 첫번째 패러미터에 익명함수를 정의함
    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    println(r2)
 
}
 
func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}
```

### 함수 원형 정의 (Delegate)

type 문은 구조체, 인터페이스, 함수 원형 등 User Defined Type을 정의하기 위해 사용된다.  

```
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```

### Closure

Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value)을 말하며, 이 때 함수는 바깥의 변수를 읽거나 쓸 수 있게 된다.  

```
package main
 
func nextValue() func() int {
    i := 0
    return func() int {  // closure 함수
        i++
        return i
    }
}
 
func main() {
    next := nextValue()
 
    println(next())  // 1
    println(next())  // 2
 
    anotherNext := nextValue()
    println(anotherNext()) // 1 다시 시작
    println(anotherNext()) // 2
}
```

예제에서 nextValue() 함수는 int를 리턴하는 익명 함수를 리턴한다.  
이 익명 함수가 그 함수 바깥에 있는 nextValue() 함수 안의 변수 i 를 참조하고 있다. 
next := nextValue() 에서 Closure 함수를 next라는 변수에 할당한 후에, 계속 next()를 2번 호출하는데 이때마다 Clouse 함수내의 변수 i는 계속 증가된 값을 가지고 있게 된다.  
이것은 마치 next 라는 함수값이 변수 i 를 내부에 유지하고 있는 모양새인데, 이를 closure라 칭한다.  
anotherNext := nextValue()와 같이 새로운 Closure 함수 값을 생성할 때, 변수 i과 함께 nextValue 함수 전체가 초기화 되므로 다시 1부터 카운팅을 하게 된다.

## Go 자료구조

### Array

```
var a [3]int  // 정수형 3개 요소를 갖는 배열 a 선언
a[0] = 1
a[1] = 2
a[2] = 3
println(a[1]) // 2 출력

var a1 = [3]int{1, 2, 3} // 값으로 배열 초기화
var a3 = [...]int{1, 2, 3} // 배열크기 자동으로
```

```
var multiArray [3][4][5]int  // 다차원 배열 정의
multiArray[0][1][2] = 10     // 다차원 배열 사용

var multiArray1 = [2][3]int{  // 다차원 배열 값으로 초기화
    {1, 2, 3},
    {4, 5, 6},  // 끝에 콤마 추가
}
```

### Slice

동적으로 크기가 변화하는 1차원 자료구조다.  
다른 언어의 List와 동일한 역할을 수행하며, 배열과 달리 고정된 크기를 미리 지정하지 않을 수 있고, 차후 그 크기를 동적으로 변경할 수도 있고, 또한 부분 배열을 발췌할 수도 있다.  

```
var a []int        // slice 변수 선언
a = []int{1, 2, 3} // slice에 리터럴값 지정
a[1] = 10
fmt.Println(a)     // [1, 10, 3]출력

s := make([]int, 5, 10) // make로 slice 생성
println(len(s), cap(s)) // len 5, cap 10
```

```
s := []int{0, 1, 2, 3, 4, 5}
s = s[2:5]     // 부분 slice 생성, 2, 3, 4
s = s[1:]      // 3, 4
fmt.Println(s) //
```

```
s := []int{0, 1}

// 하나 확장
s = append(s, 2)       // 0, 1, 2
// 복수 요소들 확장
s = append(s, 3, 4, 5) // 0,1,2,3,4,5
```

### Map

### Channel

### new vs make

new()는 메모리를 할당하되 초기화는 하지 않는다.  
메모리를 할당하고 해당 객체에 제로(Zero Value)값을 설정하고 해당 객체에 대한 포인터를 반환한다.

```
m := new(MyType)
fmt.Printf"(%T %#v", m, m) // *main.MyType &amp;main.MyType{}
```

포인터를 반환한다는 의미에서 위 코드와 아래의 코드는 동일하다.

```
m := &MyType{}
fmt.Printf"(%T %#v", m, m) // *main.MyType &amp;main.MyType{}
```

make()의 경우 slice, map, channel 만으로 사용처가 한정되어 있는데, 이 세 자료구조가 내부 데이터 구조를 가지기 때문이다.  
Effective Go 에서는 new()는 초기화를 하지 않는, 즉 제로 값이 되는 것에 대해 helpful 하다고 언급 했다.  
이 세 자료구조는 단순히 제로 값으로 초기화하는 것으로는 사용이 불가하기 때문에, make() 동작으로 내부 이니셜라이즈를 진행해줘야 한다.

# Go Tips

* 프로젝트 루트 디렉토리 기준으로 `mod init` 동작을 통해 `go.mod` 파일을 만들었을 때,  
타 디렉토리의 패키지들은 전체 모듈 이름을 다 적고 나서 마지막에 폴더 이름을 적어야 정상적으로 import 할 수 있음.
* 현재 패키지가 아닌 다른 패키지의 함수는 앞 글자가 대문자인 경우에만 export 됨.

# 부록. Go 관련 공부 사이트

* http://golang.site/go/basics
* https://yoongrammer.tistory.com/category/%EC%96%B8%EC%96%B4/Go%20%EC%96%B8%EC%96%B4
* https://dev-yakuza.posstree.com/ko/golang/
