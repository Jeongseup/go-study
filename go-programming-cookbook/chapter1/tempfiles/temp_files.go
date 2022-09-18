package tempfiles

import (
	"fmt"
	"io/ioutil"
	"log"
)

// WorkWithTemp 함수는 임시 파일 및 디렉터리 작업을 하는 기본적인 패턴을 보여준다
func WorkWithTemp() error {
	// 첫 번째 인자는공백인데, 이는 os.TempDir()이 반환한 위치에 디렉터리를 생성한다는 것을 의미
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		log.Println(err)
		return err
	}
	// 나중에 이 작업을 수행하려는 경우, 함수가 종료될 때
	// 임시 파일의 모든 내용을 삭제
	// defer os.RemoveAll(t)
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(tf.Name())
	return nil
}
