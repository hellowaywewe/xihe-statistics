package app

import (
	"errors"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
	"time"
)

func (cmd *UserWithBigModelAddCmd) Validate() error {
	err := errors.New("invalid cmd of add user query big model record")

	b := cmd.UserName != ""

	if !b {
		return err
	}

	_, err = domain.NewBigModel(cmd.BigModel.BigModel())

	if err != nil {
		return err
	}

	return nil
}

type UserWithBigModelAddCmd struct {
	domain.UserWithBigModel
}

type BigModelRecordService interface {
	AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error
	GetBigModelRecordsByType(domain.BigModel) (BigModelDTO, error)
}

func NewBigModelRecordService(
	ub repository.UserWithBigModel,
) BigModelRecordService {
	return bigModelRecordService{
		ub: ub,
	}
}

type bigModelRecordService struct {
	ub repository.UserWithBigModel
}

func (b bigModelRecordService) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error {
	v := new(domain.UserWithBigModel)
	cmd.toBigModel(v)

	err := b.ub.Add(v)
	if err != nil {
		return err
	}
	return nil
}

func (b bigModelRecordService) GetBigModelRecordsByType(d domain.BigModel) (dto BigModelDTO, err error) {
	bm, err := b.ub.Get(d)
	if err != nil {
		return
	}

	users := make([]string, len(bm))
	for j := range bm {
		users[j] = bm[j].UserName
	}
	users = RemoveRepeatedElement(users)

	dto = BigModelDTO{
		BigModel: d.BigModel(),
		Users:    users,
		Counts:   len(users),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05+08:00"),
	}

	return
}

func (cmd *UserWithBigModelAddCmd) toBigModel(r *domain.UserWithBigModel) {
	now := time.Now().Unix()

	*r = domain.UserWithBigModel{
		UserName: cmd.UserName,
		BigModel: cmd.BigModel,
		CreateAt: now,
	}
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

type BigModelDTO struct {
	BigModel string   `json:"bigmodel"`
	Users    []string `json:"user_list"`
	Counts   int      `json:"counts"`
	UpdateAt string   `json:"update_at"`
}

type BigModelListDTO struct {
}
