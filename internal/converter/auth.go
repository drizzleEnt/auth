package converter

import (
	"database/sql"

	"github.com/drizzleent/auth/internal/model"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updateAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updateAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		UserCreate: ToUserCreateFromService(&user.UserCreate),
		CreatedAt:  timestamppb.New(user.CreatedAt),
		UpdatedAt:  updateAt,
	}
}

func ToUserCreateFromService(user *model.UserCreate) *desc.UserCreate {
	return &desc.UserCreate{
		UserUpdate: ToUserUpdateFromService(&user.UserUpdate),
		Password:   user.Password,
	}
}

func ToUserUpdateFromService(user *model.UserUpdate) *desc.UserUpdate {
	return &desc.UserUpdate{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  desc.Role(user.Role),
	}
}

func ToUserFromDesc(user *desc.User) *model.User {
	var updateAt sql.NullTime

	if user.UpdatedAt.CheckValid() == nil {
		updateAt.Valid = true
		updateAt.Time = user.UpdatedAt.AsTime()
	}

	userCreate := ToUserFromDescCreate(user.GetUserCreate())
	return &model.User{
		UserCreate: *userCreate,
		CreatedAt:  user.CreatedAt.AsTime(),
		UpdatedAt:  updateAt,
	}
}

func ToUserFromDescCreate(user *desc.UserCreate) *model.UserCreate {
	userUpdate := ToUserFromDescUpdate(user.GetUserUpdate())
	return &model.UserCreate{
		UserUpdate: *userUpdate,
		Password:   user.Password,
	}
}

func ToUserFromDescUpdate(user *desc.UserUpdate) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  int(user.GetRole()),
	}
}
