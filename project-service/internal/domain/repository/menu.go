package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
)

type MenuRepo interface {
	GetAll(ctx context.Context) ([]*model.ProjectMenu, error)
}
