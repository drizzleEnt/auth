package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/drizzleent/auth/internal/api/auth"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/service"
	serviceMock "github.com/drizzleent/auth/internal/service/mocks"
	desc "github.com/drizzleent/auth/pkg/user_v2"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type AuthServiceMock func(mc *minimock.Controller) service.AuthService

	type agrs struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		serviceErr = fmt.Errorf("service error")

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Animal()

		req = &desc.CreateRequest{
			Info: &desc.UserCreate{
				UserUpdate: &desc.UserUpdate{
					Id:    id,
					Name:  name,
					Email: email,
				},
				Password: password,
			},
		}

		info = &model.UserCreate{
			UserUpdate: model.UserUpdate{
				ID:    id,
				Name:  name,
				Email: email,
			},
			Password: password,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		agrs            agrs
		want            *desc.CreateResponse
		err             error
		authServiceMock AuthServiceMock
	}{
		{
			name: "Seccses case",
			agrs: agrs{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "Error case",
			agrs: agrs{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			res, err := api.Create(tt.agrs.ctx, tt.agrs.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}
