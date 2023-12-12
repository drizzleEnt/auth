package converter

import (
	"github.com/drizzleent/auth/internal/model"
	desc "github.com/drizzleent/auth/pkg/login_v1"
)

func ToUserClaimsFromLogin(req *desc.Login) *model.UserClaims {
	return &model.UserClaims{
		Username: req.GetUsername(),
	}
}
