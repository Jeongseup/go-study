package confformat

import "fmt"

const (
	exampleTOML = `name="Example1"
age=99
    `

	exampleJSON = `{"name":"Example2","age":98}`

	exampleYAML = `name: Example3
age: 97
    `
)

// UnmarshalAll 함수는 다양한 포맷의 데이터를 가져와 구조체로 변환
func UnmarshalAll() error {
	t := TOMLData{}
	j := JSONData{}
	y := YAMLData{}

	// toml unmarshaling
	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	// json unmarshaling
	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	// yaml unmarshaling
	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("YAML Unmarshal =", y)

	return nil
}
