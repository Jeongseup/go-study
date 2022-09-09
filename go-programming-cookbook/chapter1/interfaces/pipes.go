package interfaces

import (
	"io"
	"os"
)

// PipeExapmle 함수는 io 인터페이스를 사용하는 더 많은 예제를 제공한다.
func PipeExapmle() error {

	r, w := io.Pipe()

	go func() {

		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return nil
	}

	return nil
}
