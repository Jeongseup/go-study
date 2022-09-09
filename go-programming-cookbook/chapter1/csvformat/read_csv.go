package csvformat

import (
	"encoding/csv"
	"io"
	"strconv"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// 처음에는 헤더를 가져오므로 무시
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{
			Title:    record[0],
			Director: record[1],
			Year:     int(year),
		}

		movies = append(movies, m)
	}

	return movies, nil
}
