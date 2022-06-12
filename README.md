# go-study
고랭을 스터디할 고랭

GOROOT 는 GO의 tools, standard library, compiler 가 있다.
GOPATH 가 중요한데 바로 workspace 경로를 가리키는 변수다. 만약 내가 default 가 아닌 곳에서 workspace를 만든다면 변경해야한다.
GOBIN 'go install' 해서 생성되는 실행파일의 저장 경로 /bin


출처: https://litaro.tistory.com/entry/Go-Packages?category=794245 [litaro's blog]
ㅁㅇㅇㅁㅇ


unitl 1:33:35

```go

a := []int{1,2,3} // slice
a := [...]int{1,2,3} // array

// slice
a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// copy a
// NOT reference
b = a
// But slice a like array[:] (which means sliced)
// all data will be refercens
// that is both in slice and in array 
b := a[:]



append function is posssible only slice! NOT array.
because array dont need to that function, just insert data into index where you want 
```



- [golang-development study](https://youtu.be/W5b64DXeP0o)


- [Learn Go Programming by Building 11 Projects](https://youtu.be/jFfo23yIWac)
