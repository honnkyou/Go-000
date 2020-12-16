package biz

type Homework struct {
	id     int64
	title  string
	author string
	detail string
}

type HWRepo interface {
	Take(*Homework) int64
}

type HWUseCase struct {
	repo HWRepo
}

func NewHWUseCase(repo HWRepo) *HWUseCase {
	return &HWUseCase{repo: repo}
}

func (u *HWUseCase) TakeHomework(h *Homework) {
	id := u.repo.Take(h)
	h.id = id
}
