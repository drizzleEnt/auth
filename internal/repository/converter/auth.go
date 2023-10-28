package converter

import (
	"github.com/drizzleent/auth/internal/model"
	repoModel "github.com/drizzleent/auth/internal/repository/data_model"
)

func ToModleUserFromRepo(rUser *repoModel.User) *model.User {

	return &model.User{
		ID:        rUser.ID,
		Name:      rUser.Name,
		Email:     rUser.Name,
		Password:  rUser.Password,
		Role:      rUser.Role,
		CreatedAt: rUser.CreatedAt,
		UpdatedAt: rUser.UpdatedAt,
	}
}
