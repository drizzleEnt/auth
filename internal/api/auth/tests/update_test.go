package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/drizzleent/auth/internal/api/auth"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/service"
	serviceMock "github.com/drizzleent/auth/internal/service/mocks"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type AuthServiceMock func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			Info: &desc.UserUpdate{
				Id:    id,
				Name:  name,
				Email: email,
				Role:  0,
			},
		}

		info = &model.UserUpdate{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  0,
		}

		res interface{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            interface{}
		err             error
		AuthServiceMock AuthServiceMock
	}{
		{
			name: "Secess case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			AuthServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.UpdateMock.Expect(ctx, info).Return(nil)
				return mock
			},
		},
		{
			name: "Error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			AuthServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.UpdateMock.Expect(ctx, info).Return(serviceErr)
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

			_, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
