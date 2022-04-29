/*
변수의 구성요소
1. 이름
2. 값
3. 메모리 주소
4. 사이즈
5. 끝점? : 주소 + 사이즈

각 변수는 type이 나뉨. (32bit/64bit 컴퓨터에 따라)

int는 32bit에서 4byte, 64bit에서 8byte
그래서 아래와 같은 타입 + 사이즈에 따라 또 나뉜다.

int32 4byte
int64 8byte
int8 1byte => -128 ~ +127
uint8 1byte => 0 ~ +255
int16 2byte

float32 4byte => 숫자 부분 7개까지만 표현 가능, 1개는 10^(-n)을 표현하기 위함
float64 8byte => 숫자 부분 15개까지 표현 가능,


bool size X
string size X
그리고 하나의 문자는 하나의 type rune으로 지정된다.

bigNumber는 하나의 타입이 아니라 아마.. 스트링으로 변환해서? 8byte를 넘어가는 숫자를 표현하는데 쓰임


왜 이런 것들이 필요할까?
메모리는 공간의 한계가 존재함.

그래서 윈도우OS에서는 HDD Swaping을 하는데 이게 굉장히 느림

그리고 사실 메모리는 네트워크 bandwith 때문이 크다.
네트워크에 전송되는 데이터양을 줄이면 줄일 수록 좋다.
그래서 안쓰는 애들을 압축하는 것이 좋다. 그래서 base64 en/decoding을 하는건가?

네트워크 100Mbps?
패킷? 데이터
*/

package main

import "fmt"

func main() {
	var a int
	var b int
	var name string = "jeongseup"
	a = 3
	b = 4

	fmt.Println(a + b)
	fmt.Println(name)
}