package messages

import (
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
)

type Handler struct {
	Log *logrus.Entry

	MaxRetry         int
	TrainingEndpoint string

	BigModel app.BigModelRecordMessageService
	Repo     app.RepoRecordMessageService
	Register app.RegisterRecordMessageService
}

func (h *Handler) AddBigModelRecord(d *domain.UserWithBigModel) error { // implement domain function with app function
	cmd := app.UserWithBigModelAddCmd{
		UserWithBigModel: *d,
	}
	return h.BigModel.AddUserWithBigModel(&cmd)
}

func (h *Handler) AddRepoRecord(d *domain.UserWithRepo) error {
	cmd := app.RepoRecordAddCmd{
		UserWithRepo: *d,
	}
	return h.Repo.Add(&cmd)
}

func (h *Handler) AddRegisterRecord(d *domain.RegisterRecord) error {
	return h.Register.Add(d)
}