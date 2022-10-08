package mycontext

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initalize 함수는 설정을 위해 세 개의 함수를 호출
// 그런 다음 종료 전에 로그 기록
func Initalize() {
	// 기본 로그 설정
	log.SetHandler(text.New(os.Stdout))
	// context 초기화
	ctx := context.Background()
	// 로거 생성 후 ctx 연결
	ctx, e := FromContext(ctx, log.Log)

	// 필드 설정
	ctx = WithField(ctx, "id", "123")
	ctx = WithField(ctx, "id2", "124")
	ctx = WithField(ctx, "id3", "125")

	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")

}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
	_ = ctx
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
	_ = ctx
}
