package nulls

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer 구조체는 동일하지만 *Int를 사용한다는 점만 다름
type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

// PointerEncoding 함수는 nil/omitted 값을 처리하는 방법을 보여줌
func PointerEncoding() error {
	e := ExamplePointer{}
	// no age가 0 age라는 점에 주의?
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, no age %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, with age = 0: %+v\n", e)
	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with age = 0:", string(value))

	return nil
}
