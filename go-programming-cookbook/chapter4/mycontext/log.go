package mycontext

import (
	"context"
	"fmt"

	"github.com/apex/log"
)

type key int

/// logFields는 context 로깅에 사용할 키
const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	fmt.Printf("fields : %v, ok: %v\n", fields, ok)
	if !ok {
		fmt.Println("first")
		f := make(log.Fields)

		fields = &f
	}

	return fields
}

// FromContext 함수는 항목과 context를 입력받고
// context 객체로부터 데이터 채워진 항목을 반환
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	// field 초기화
	fields := getFields(ctx)
	// 필드 로그 연결
	e := l.WithFields(fields)

	// 컨텍스트 데이터 채우기
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField 함수는 context에 로그 필드를 추가한다
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields 함수는 context에 많은 로그 필드를 추가한다
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}

	// 컨텍스트에 데이터 채우기
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}
