package confformat

import (
	"bytes"

	"github.com/go-yaml/yaml"
)

// YAMLData 구조체는 YAML 구조체 태그를 갖는 일반적인 구조체
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode will decode into TOMLData
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
