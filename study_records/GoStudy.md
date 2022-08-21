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
    sum := func(n ...int) int { //익명 함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
 
    result := sum(1, 2, 3, 4, 5) //익명 함수 호출
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
 
    // 직접 첫번째 패러미터에 익명 함수를 정의함
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

```
// append slice
sliceA := []int{1, 2, 3}
sliceB := []int{4, 5, 6}
sliceA = append(sliceA, sliceB...)  // [1 2 3 4 5 6]

// copy slice
source := []int{0, 1, 2}
target := make([]int, len(source), cap(source)*2)
copy(target, source)
fmt.Println(target)  // [0 1 2 ] 출력
```

### Map

키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조이다.  
"map[Key타입]Value타입" 과 같이 선언할 수 있다.  

```
var idMap map[int]string        // map 변수 선언
idMap = make(map[int]string)    // make로 map 생성

// 리터럴을 사용한 초기화
tickers := map[string]string{
    "GOOG": "Google Inc",
    "MSFT": "Microsoft",
    "FB":   "FaceBook",
}
```

```
m := make(map[int]string)

//추가 혹은 갱신
m[901] = "Apple"
m[134] = "Grape"
m[777] = "Tomato"

// 키에 대한 값 읽기
str := m[134]

noData := m[999] // 값이 없으면 nil 혹은 zero 리턴

// 특정 키 삭제
delete(m, 777)

// map 키 체크
val, exists := tickers["MSFT"]
if !exists {
    println("No MSFT ticker")
}
```

```
// for range 문을 사용하여 모든 맵 요소 출력
// Map은 unordered 이므로 순서는 무작위
for key, val := range myMap {
    fmt.Println(key, val)
}
```

### Channel

Channel은 데이터를 주고 받는 통로라 볼 수 있는데, channel은 make() 함수를 통해 미리 생성되어야 하며, channel 연산자 <- 을 통해 데이터를 보내고 받는다.  
channel은 흔히 routine들 사이 데이터를 주고 받는데 사용되는데, 상대편이 준비될 때까지 channel에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용된다.  

아래 예제에서 메인 함수는 마지막에서 channel로부터 데이터를 받고 있는데, 상대편 routine에서 데이터를 전송할 때까지는 계속 대기하게 된다. 따라서, 이 예제에서는 time.Sleep() 이나 fmt.Scanf() 같이 goroutine 이 끝날 때까지 기다리는 코드를 적지 않았다.

```
// 정수형 channel을 생성
ch := make(chan int)

go func() {
    ch <- 123   // channel에 123을 보냄
}()

var i int
i = <- ch  // channel로부터 123을 받음
println(i)
```

 channel은 수신자와 송신자가 서로를 기다리는 속성이 있어, 이를 통해 routine 대기 시퀀스를 구현할 수 있다.  
 즉, 익명 함수를 사용한 한 Go 루틴에서 어떤 작업이 실행되고 있을 때, 메인루틴은 <-done 에서 계속 수신하며 대기하고 있게 된다.  
 익명 함수 Go 루틴에서 작업이 끝난 후, done channel에 true를 보내면, 수신자 메인루틴은 이를 받고 프로그램을 끝내게 된다.

```
// channel 생성
done := make(chan bool)

// routine 생성 + 시작
go func() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    done <- true
}()

// 위의 routine이 끝날 때까지 대기
<-done
```

2가지의 channel이 있는데, Unbuffered Channel과 Buffered Channel이 있다.  
위의 예제에서의 channel은 Unbuffered Channel로서 이 channel에서는 하나의 수신자가 데이타를 받을 때까지 송신자가 데이타를 보내는 channel에 묶여 있게 된다.  
하지만, Buffered Channel을 사용하면 비록 수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이타를 보내고 계속 다른 일을 수행할 수 있다.  

```
// Unbuffered Channel의 deadlock 상황
func unbuffered() {
  c := make(chan int)
  c <- 1   // 수신자가 없으므로 deadlock
  fmt.Println(<-c) // 여기까지 올 수 없음
}

// Buffered Channel의 deadlock 회피?
func buffered() {
    ch := make(chan int, 1)
    ch <- 101     // 수신자가 없더라도 최대 버퍼 전까지는 넣고 넘어갈 수 있음
    fmt.Println(<-ch)
}
```

channel을 함수의 패러미터도 전달할 때 해당 channel로 송신만 할 것인지 혹은 수신만할 것인지를 지정할 수 있다.  
송신 패러미터는 (p chan<- int)와 같이 chan<- 을 사용하고, 수신 패러미터는 (p <-chan int)와 같이 <-chan 을 사용한다.

```
func main() {
    ch := make(chan string, 1)
    sendChan(ch)
    receiveChan(ch)
}
 
func sendChan(ch chan<- string) {
    ch <- "Data"
    // x := <-ch // 에러발생
}
 
func receiveChan(ch <-chan string) {
    data := <-ch
    fmt.Println(data)
}
```

close()함수를 사용하여 channel을 닫을 수 있다.  
channel이 닫힌 이후에도 계속 수신은 가능하다.

```
ch := make(chan int, 2)
    
// channel에 송신
ch <- 1
ch <- 2
    
// channel을 닫는다
close(ch)

// channel 수신
println(<-ch)
println(<-ch)
    
if _, success := <-ch; !success {
    println("더이상 데이타 없음.")
}
```

for range문을 사용하면 channel로부터 계속 수신하다가 channel이 닫힌 것을 감지하면 for 루프를 종료한다.

```
// 일반 반복문 사용
for {
    if i, success := <-ch; success {
        println(i)
    } else {
        break
    }
}

// range문 사용
for i := range ch {
    println(i)
}
```

Go의 select문은 여러 개의 case문에서 각각 다른 channel을 기다리다가 준비가 된 channel case를 실행한다.  
case channel들이 준비되지 않으면 계속 대기하고, 복수 channel에 신호가 오면 Go 런타임이 랜덤하게 그 중 한 개를 선택한다.  
select문에 default 문이 있으면, case문 channel이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행한다.

```
func main() {
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
    time.Sleep(1 * time.Second)
    done <- true
}
 
func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}
```

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

## Go Routine

Go 루틴(routine)은 Go 런타임이 관리하는 Lightweight 논리적(혹은 가상적) thread다.  
"go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다.  
routine은 비동기적으로(asynchronously) 함수를 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.  

routine은 OS thread보다 훨씬 가볍게 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리한다.  
Go 런타임 상에서 관리되는 작업단위인 여러 routine들은 종종 하나의 OS thread 1개로도 실행되곤 한다.  
즉, routine들은 OS thread와 1 대 1로 대응되지 않고, Multiplexing으로 훨씬 적은 OS thread를 사용한다.  
메모리 측면에서도 OS thread가 1 메가바이트의 스택을 갖는 반면, routine은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다(필요시 동적으로 증가).  
Go 런타임은 routine을 관리하면서 channel을 통해 routine 간의 통신을 쉽게 할 수 있도록 하였다.

```
func say(s string) {
    for i := 0; i < 10; i++ {
        fmt.Println(s, "***", i)
    }
}
 
func main() {
    // 함수를 동기적으로 실행
    say("Sync")
 
    // 함수를 비동기적으로 실행
    go say("Async1")
    go say("Async2")
    go say("Async3")
 
    // 3초 대기
    time.Sleep(time.Second * 3)
}
```

아래 예제에서는 익명 함수를 통해 routine을 실행시킨다.  
여기서 sync.WaitGroup을 사용하고 있는데, 이는 기본적으로 여러 routine들이 끝날 때까지 기다리는 역할을 한다.  
WaitGroup을 사용하기 위해서는 먼저 Add() 메소드에 몇 개의 routine을 기다릴 것인지 지정하고, 각 routine에서 Done() 메서드를 호출한다 (여기서는 defer 를 사용하였다).  
그리고 메인 함수에서는 Wait() 메서드를 호출하여, routine들이 모두 끝나기를 기다린다.

```
func main() {
    // WaitGroup 생성. 2개의 routine을 기다림.
    var wait sync.WaitGroup
    wait.Add(2)
 
    // 익명 함수를 사용한 routine
    go func() {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println("Hello")
    }()
 
    // 익명 함수에 패러미터 전달
    go func(msg string) {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println(msg)
    }("Hi")
 
    wait.Wait() // routine을 두 개 끝날 때까지 대기
}
```

Go는 디폴트로 1개의 CPU를 사용한다.  
즉, 여러 개의 routine을을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다 (Concurrent 처리).  
만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 (Parallel 처리)하게 할 수 있는데, 병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다 (여기서 CPU 수는 Logical CPU 수를 가리킨다).

```
// 4개의 CPU 사용
runtime.GOMAXPROCS(4)
```

## Go 패키지

Go는 패키지를 사용해서 작은 단위의 컴포넌트를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장한다.  
공식 라이브러리도 패키지 형태로 제공한다(표준 라이브러리 패키지).  
특정 패키지의 함수, 구조체, 인터페이스, 메서드를 사용하고 싶다면 해당 패키지의 import가 필요하다.  
현재 패키지의 함수, 구조체, 인터페이스, 메서드를 다른 패키지에서 사용 가능하게 만들고 싶다면 첫문자를 대문자로 시작함으로써 public 으로 사용할 수 있다.  

```
package main    // 현재 패키지 선언
 
import "fmt"    // import 할 패키지 선언
 
func main(){
 fmt.Println("Hello")   // import한 패키지에서 기능 사용
}
```

"main" 이라고 명명된 패키지를 발견하면 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.  
그리고 이 main 패키지 안의 main() 함수가 프로그램의 Entry Point가 된다.  

패키지가 import 되고 최초로 호출될 때 호출되는 init() 함수를 작성할 수 있다.  
즉, init 함수는 패키지가 로드되면서 실행되는 함수로 별도의 호출 없이 자동으로 호출된다.  
패키지를 import 하면서 바로 init() 함수를 호출하고 싶다면 앞에 _ 를 붙인다.  

```
package testlib
 
var pop map[string]string
 
func init() {   // 패키지 사용 시 map 초기화
    pop = make(map[string]string)
}
```

```
package main
import _ "other/xlib"   // 패키지 import 시 init() 호출
```

사용자 정의 라이브러리 패키지는 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성한다.  
하나의 서브 폴더안에 있는 .go 파일들은 동일한 패키지명을 가지며, 패키지명은 해당 폴더의 이름과 같게 한다.  
즉, 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.

## Go 구조체

구조체는 필드의 집합체, 컨테이너이다.  
Go에서 struct는 필드 데이타만을 가지며, 메서드를 갖지 않는다.  
struct를 정의하기 위해서는 Custom Type을 정의하는데 사용하는 type 문을 사용한다.

```
// struct 정의
type person struct {
    name string
    age  int
}

func main() {
    // person 객체 생성
    p := person{}
     
    // 필드값 설정
    p.name = "Lee"
    p.age = 10
}
```

```
// 생성과 동시에 내부 정보 기입
var p1 person 
p1 = person{"Bob", 20}

p2 := person{name: "Sean", age: 50}
```

```
// new로 생성
p := new(person)
p.name = "Lee"  // p가 포인터라도 . 을 사용
```

## Go 메서드

Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.  
메서드는 함수 정의에서 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 되고,  
struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용되는데, 이를 receiver라 칭한다.

```
//Rect struct 정의
type Rect struct {
    width, height int
}
 
//Rect의 area() 메서드 정의
func (r Rect) area() int {      // 어떤 struct의 메서드인지 r Rect 형태로 적음
    return r.width * r.height   // 함수 내부에서 r로 self 객체를 호출
}
 
func main() {
    rect := Rect{10, 20}    // struct 생성
    area := rect.area()     // struct의 메서드 호출
}
```

Receiver는 Value와 Reference 형태로 전달이 가능하다.  
Pass by value일 경우 struct의 데이타를 복사하여 전달하기 때문에 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이타는 변경되지 않는다.  
Pass by reference일 경우 struct의 실 메모리 상 주소를 전달하기 때문에 메서드내에서 그 struct의 필드값이 변경되면 호출자의 데이타가 변경된다.

```
// 포인터 Receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area2()        // 메서드 호출
    println(rect.width, area)   // 11 220 출력
}
```

## Go 인터페이스

struct가 필드들의 집합체라면, interface는 메서드들의 집합체이다.  
interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의한다.  
이후에 해당 interface를 구현하는 메서드에서 내용을 작성한다.  

```
// Shape 인터페이스 정의
type Shape interface {
    area() float64
    perimeter() float64
}

// Rect 정의
type Rect struct {
    width, height float64
}
 
// Circle 정의
type Circle struct {
    radius float64
}
 
//Rect 타입에 대한 Shape 인터페이스 구현 
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
     return 2 * (r.width + r.height)
}
 
//Circle 타입에 대한 Shape 인터페이스 구현 
func (c Circle) area() float64 { 
    return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}
```

함수 파라미터가 interface 인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 모두 입력 파라미터로 사용될 수 있다는 것을 의미한다.  

```
func main() {
    r := Rect{10., 20.}
    c := Circle{10}
 
    showArea(r, c)
}
 
func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area() //인터페이스 메서드 호출
        println(a)
    }
}
```

모든 Type을 나타내기 위해 빈 인터페이스를 사용한다.  
즉, 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며, 여러 다른 언어에서 흔히 일컫는 Dynamic Type 이라고 볼 수 있다.  
(empty interface는 C#, Java 에서 object, C/C++ 에서는 void*)

```
func main() {
    var x interface{}
    x = 1 
    x = "Tom"
 
    printIt(x)
}
 
func printIt(v interface{}) {
    fmt.Println(v) //Tom
}
```

Interface x와 타입 T에 대하여 x.(T)로 표현했을 때, 이는 x가 nil이 아니며, x는 T 타입에 속한다는 점을 확인(assert)하는 것으로 이러한 표현을 "Type Assertion"이라 부른다.

```
var a interface{} = 1
 
i := a       // a와 i 는 dynamic type, 값은 1
j := a.(int) // j는 int 타입, 값은 1

println(i)  // 포인터주소 출력
println(j)  // 1 출력
```

## Go defer, panic, recover

defer 키워드는 특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 한다.  
일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용된다.  

```
func main() {
    f, err := os.Open("1.txt")
    if err != nil {
        panic(err)
    }
 
    // main 마지막에 파일 close 실행
    defer f.Close()
 
    // f를 활용한 작업, 파일 읽기
    bytes := make([]byte, 1024)
    f.Read(bytes)
    println(len(bytes))
}
```

panic() 함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴한다.  
이러한 panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용된다.  
그리고 마지막에는 프로그램이 에러를 내고 종료하게 된다.

```
func main() {
    // 잘못된 파일명을 넣음
    openFile("Invalid.txt")
     
    // openFile() 안에서 panic이 실행되면
    // 아래 println 문장은 실행 안됨
    // panic의 즉시 리턴이 위에서 적용되기 때문
    println("Done")
}
 
func openFile(fn string) {
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
 
    defer f.Close()
}
```

recover()함수는 panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수이다. 

```
func main() {
    // 잘못된 파일명을 넣음
    openFile("Invalid.txt")
 
    // recover에 의해
    // 이 문장 실행됨
    println("Done") 
}
 
func openFile(fn string) {
    // recover 함수를 defer로 걸어둠. panic 호출시 실행
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()
 
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
 
    defer f.Close()
}
```

# Go Tips

* 프로젝트 루트 디렉토리 기준으로 `mod init` 동작을 통해 `go.mod` 파일을 만들었을 때,  
타 디렉토리의 패키지들은 전체 모듈 이름을 다 적고 나서 마지막에 폴더 이름을 적어야 정상적으로 import 할 수 있음.
* 현재 패키지가 아닌 다른 패키지의 함수는 앞 글자가 대문자인 경우에만 export 됨.

# 부록. Go 관련 공부 사이트

* http://golang.site/go/basics
* https://yoongrammer.tistory.com/category/%EC%96%B8%EC%96%B4/Go%20%EC%96%B8%EC%96%B4
* https://dev-yakuza.posstree.com/ko/golang/
