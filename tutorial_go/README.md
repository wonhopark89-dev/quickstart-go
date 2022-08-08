- main.go  
- func main() {} // 기본 진입
```
go run main.go
```

- 함수명 - 대문자 -> export -> import
- 함수명 - 소문자 -> private

----

- go 는 타입 언어
- const 상수 ( 변경 불가 )
- var 변수

```
var name string = "john"
-> name := "john"

func 안에서만 사용가능함, 변수에만 사용가능
축약형은 처음 변수로 부터 타입을 추론

```

- go 는 선언한 변수를 사용하지 않으면 에러 발생
```
{sth} declared but not used
```