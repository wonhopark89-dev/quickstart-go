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
- defer 는 함수가 끝나고 난 후에 실행된다
```
func sample() {
    defer fmt.PrintLn("sample done")
    return
}
```
- range 는 인덱스 와 요소값을 반환
```
for index, number := range numbers {
		fmt.Println(index, number)
}
```

- if, switch 각 분기문안에서만 사용할 수 있는 변수 선언이 가능
```
func sample(param1 int) bool {
   
  if param1 < 10 {
    return false
  }   
  return true


  if param2 := param1 + 5; param2 < 20 {
    // param2 는 if 문 안에서만 사용할 수 있음
    return false
  }  
  return true
}
```
- pointer
```
a := 1
b := a // 값 참조
b := &a // 주소 참조
```

- array and slice
```
names := [3]string{"a","b","c"} // array

names := []string{"a","b","c"} // slice
names = append(names, "d") // 추가 가능
```