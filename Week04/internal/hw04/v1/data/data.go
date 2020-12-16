package data

import (
	"Week04/internal/hw04/v1/biz"
	"time"
)

func NewHWRepo() biz.HWRepo {
	return &homeworkRepo{}
}

type homeworkRepo struct{}

func (repo *homeworkRepo) Take(homework *biz.Homework) int64 {
	homework.detail = "test"
	
	homeworkId := time.Now().Unix()

	return homeworkId
}
