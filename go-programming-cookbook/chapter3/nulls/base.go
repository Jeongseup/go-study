package nulls

import (
	"encoding/json"
	"fmt"
)

// name 값은 있지만 age값은 없는 json
const (
	jsonBlob     = `{"name":"Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// Example 구조체는 age와 name 필드를 갖는 기본 구조체
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

func BaseEncoding() error {
	e := Example{}
	// no age가 0 age라는 점에 주의?
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age %+v\n", e)
	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)
	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}
