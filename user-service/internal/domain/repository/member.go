package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type MemberRepository interface {
	FindMemberByAccount(ctx context.Context, account string) (bool, error)
	SaveMember(ctx context.Context, member *model.Member) error
	FindMember(ctx context.Context, account, password string) (*model.Member, error)
	FindMemberById(ctx context.Context, id int64) (*model.Member, error)
}
