package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type OrganizationRepository interface {
	SaveOrganization(ctx context.Context, org *model.Organization) error
}
