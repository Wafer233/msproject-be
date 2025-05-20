package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type MemberRepo interface {
	ExistByEmail(ctx context.Context, email string) (bool, error)
	ExistByAccount(ctx context.Context, account string) (bool, error)
	ExistByMobile(ctx context.Context, mobile string) (bool, error)
	Save(ctx context.Context, member *model.Member) error
	GetByCredentials(ctx context.Context, account string, password string) (*model.Member, error)
}
