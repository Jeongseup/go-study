package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	// 이 작업은 설정 옵션을 설정하는 선택적 작업이다
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie
	// 처음에는 헤더를 가져오므로 무시
	_, err := r.Read()
	if err != nil && err != io.EOF {
		log.Fatalf("[ERROR] : %v in Read", err)
		return nil, err
	}

	// CSV를 모두 처리할 때까지 루프를 수행한다
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panicln("Read Error")
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			log.Fatalf("[ERROR : %v in ParseInt", err)
			return nil, err
		}

		// m := Movie{record[0], record[1], int(year)}
		m := Movie{
			Title:    record[0],
			Director: record[1],
			Year:     int(year),
		}

		movies = append(movies, m)
	}

	return movies, nil
}

// AddMoviesFromText 함수는 CSV 파서를 사용한다.
func AddMoviesFromText() error {

	// 이 예제는 csv 패키지를 사용해 문자열을 가져온 후 버퍼로 변환해 읽는 예를 보여준다.
	in := `
- first out headers
movie title;director;year released

- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		log.Fatalf("[ERROR] : %v  in ReadCSV", err)
		// return err
	}

	fmt.Println("Read Buffer String")
	fmt.Printf("%#vn", m)
	return nil
}
