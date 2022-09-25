package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// TOMLData 구조체는 TOML 구조체 태그를 갖는 일반적인 구조체
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML 함수는 TOMLData 구조체를 TOML 포맷의 bytes.Buffer로 dump(반환)
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}

	return b, nil
}

// Decode 함수는 TOMLData 구조체로 디코딩한다
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
