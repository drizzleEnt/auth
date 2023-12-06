package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/drizzleent/auth/internal/api/auth"
	"github.com/drizzleent/auth/internal/service"
	serviceMock "github.com/drizzleent/auth/internal/service/mocks"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type AuthServiceMock func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}

		res any
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
			name: "Secsses case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			AuthServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(t)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			AuthServiceMock := tt.AuthServiceMock
			api := auth.NewImplementation(AuthServiceMock(mc))

			_, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
