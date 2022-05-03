# #go lang tutorials for computer binder? 
youtube links : https://youtu.be/JKaJOSweBss
---
## study log

* go lang has NO reperences, but only COPY
 -  그래서 어떤 객체의 값을 변경하기 위해서는 포인터를 사용해야 함

* rune is for UTF8 byte char

* GC(garbage collector)가 필요한 이유
go보다는 C언어에서 생기는데.. 
malloc으로 힙 메모리를 받아뒀는데 free하지 않은 경우

스택 메모리 상에서 함수 안에서 malloc해두고 
free하지 않고, 나와버린 경우 memory leak 발생.

어떻게 그 메모리가 지워도 되는지 확인할까? 
reference counter를 부여해서 이게 0이 되는 경우
날려버림.

* Dynamic Array
복사를 이용해서 추가 메모리 공간이 필요한 경우 복사하고 그 곳을 다시 포인터해서 처리하는데
단, 길이가 늘어날 때 기존 길이의 2배로 늘어나서 여유있게 확보하므로서 매번 복사되게 하지는 않는다.

이제 `make()`로 만든 경우 capacity가 있을 경우 기존 메모리 주소를 사용하게 되는데 
스크립트 언어에서의 referecen같은 느낌이 날 수 있다. 
하지만 헷갈리므로 왠만하면 `b := append(a, 1)` 이렇게 추가하는 경우 
a와 b가 다른 애라고 생각하는 것이 좋다. 

* slice bug
`b := a[:4]`등으로 슬라이스로 데이터를 가져왔다고 해서 복사되는게 아니라
참조하는 것이기 때문에 그냥 아예 새롭게 배열을 하나 make로 만들고
for loop으로 아예 확실하게 복사하는 것이 낫다.


* instance?
객체를 생명체 같이 구체화 된 것.
어떤 struct의 포인터 타입.

```go
	// 그냥 넣으면 안 바뀌고, a의 주소를 대입해줘야 set함수가 실행될 수 있다.
	// 주소값이 복사되기 때문에 참조가 가능해진다.
	// 기능 단위로 볼 것이냐?
	SetNameByFunction(&a, "bbb")
	fmt.Println(a)

	// 관계가 있는 단위로 볼 것이냐?
	// OOP에서 ER(Entity Relationship)이 중요함.
	a.SetNameByMethod("ccc")
	fmt.Println(a)

```