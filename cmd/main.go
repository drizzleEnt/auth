package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/lib/handler"
	"github.com/drizzleent/auth/pkg/user_v1"
	_ "github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()

	aLog := log.New(os.Stdout, "INFO:", log.Flags())

	url := fmt.Sprintf("%s:%s", cfg.Host, cfg.GrpcPort)

	lis, err := net.Listen("tcp", url)

	if err != nil {
		errStr := fmt.Sprintf("failed to listen server : %v", err)
		aLog.Fatalf(errStr)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	rpcSrv := handler.NewUserRpcsServer(aLog)
	user_v1.RegisterUserV1Server(s, rpcSrv)

	done := make(chan os.Signal, 1)

	go func() {
		if err := s.Serve(lis); err != nil {
			aLog.Fatalf("failed to serve %v", err)
		}
	}()

	aLog.Println("Server Started")
	<-done
	s.GracefulStop()
	aLog.Println("Server stopped")
}
