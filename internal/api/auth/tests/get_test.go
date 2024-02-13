package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/drizzleent/auth/internal/api/auth"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/service"
	serviceMock "github.com/drizzleent/auth/internal/service/mocks"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type AuthServiceMock func(mc *minimock.Controller) service.AuthService

	type agrs struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		password  = gofakeit.Animal()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		info = &model.User{
			UserCreate: model.UserCreate{
				UserUpdate: model.UserUpdate{
					ID:    id,
					Name:  name,
					Email: email,
				},
				Password: password,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		res = &desc.GetResponse{
			User: &desc.User{
				UserCreate: &desc.UserCreate{
					UserUpdate: &desc.UserUpdate{
						Id:    id,
						Name:  name,
						Email: email,
					},
					Password: password,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            agrs
		want            *desc.GetResponse
		err             error
		AuthServiceMock AuthServiceMock
	}{
		{
			name: "Secsses case",
			args: agrs{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			AuthServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.GetMock.Expect(ctx, id).Return(info, nil)
				return mock
			},
		},
		{
			name: "Error case",
			args: agrs{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			AuthServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			AuthServiceMock := tt.AuthServiceMock(mc)
			api := auth.NewImplementation(AuthServiceMock)

			get, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, get)
		})

	}
}
