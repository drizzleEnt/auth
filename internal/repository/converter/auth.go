package converter

import (
	"github.com/drizzleent/auth/internal/model"
	repoModel "github.com/drizzleent/auth/internal/repository/data_model"
)

func ToModleUserFromRepo(user repoModel.User) *model.User {

	return &model.User{
		UserCreate: ToUserCFromRepo(user.UserCreate),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func ToUserCFromRepo(user repoModel.UserCreate) model.UserCreate {
	return model.UserCreate{
		UserUpdate: ToUserUFromRepo(user.UserUpdate),
		Password:   user.Password,
	}
}

func ToUserUFromRepo(user repoModel.UserUpdate) model.UserUpdate {
	return model.UserUpdate{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
