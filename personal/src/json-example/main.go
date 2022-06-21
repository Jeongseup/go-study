package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type ResponseWithoutJSON struct {
	Page   int
	Fruits []string
}

type ResponseWithJSON struct {
	Page    int      `json:"page"`
	Fruits  []string `json:"fruits"`
	Tag     string   `json:"tag,omitempty"`
	Comment string   `json:"-"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println("origin :", boolB)
	fmt.Println("converted :", string(boolB))
	fmt.Println()

	intB, _ := json.Marshal(1)
	fmt.Println("origin :", intB)
	fmt.Println("converted :", string(intB))
	fmt.Println()

	floatB, _ := json.Marshal(2.34)
	fmt.Println("origin :", floatB)
	fmt.Println("converted :", string(floatB))
	fmt.Println()

	stringB, _ := json.Marshal("gohher")
	fmt.Println("origin :", stringB)
	fmt.Println("converted :", string(stringB))
	fmt.Println()

	sliceD := []string{"apple", "peach", "pear"}
	fmt.Println(sliceD)

	sliceB, _ := json.Marshal(sliceD)
	fmt.Println("origin :", sliceB)
	fmt.Println("converted :", string(sliceB))
	fmt.Println()

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println(mapD, reflect.TypeOf(mapD))

	mapB, _ := json.Marshal(mapD)
	fmt.Println("origin :", mapB, reflect.TypeOf(mapB))
	fmt.Println("converted :", string(mapB), reflect.TypeOf(string(mapB)))
	fmt.Println()

	res1D := &ResponseWithoutJSON{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println(*res1D, res1D)

	res1B, _ := json.Marshal(res1D)
	fmt.Println("origin :", res1B, reflect.TypeOf(res1B))
	fmt.Println("converted :", string(res1B), reflect.TypeOf(string(res1B)))
	fmt.Println()

	res2D := &ResponseWithJSON{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println(res2D)

	res2B, _ := json.Marshal(res2D)
	fmt.Println("origin :", res2B, reflect.TypeOf(res2B))
	fmt.Println("converted :", string(res2B), reflect.TypeOf(string(res2B)))
	fmt.Println()

	// Unmarshal
	// json 타입이 없으니까 string 써주고 타입에는 byte를 줘버리는데?
	// 이제 이걸 map타입으로 형변환시킬 때? Unmarshal를 사용함
	byt := []byte(`{"num":6.13, "str":["a", "b"]}`)
	fmt.Println(byt)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	os.Exit(0)
}
