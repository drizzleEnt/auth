package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/lib/handler"
	"github.com/drizzleent/auth/internal/lib/repository"
	"github.com/drizzleent/auth/pkg/user_v1"
	_ "github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
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

	if err := gotenv.Load(); err != nil {
		aLog.Fatalf("error loading env vars %v", err)
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     cfg.Host,
		Port:     os.Getenv("PG_PORT"),
		UserName: os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   os.Getenv("PG_DATABASE_NAME"),
		SSLMode:  "disable",
	})
	fmt.Println(cfg.Host, os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE_NAME"))

	if err != nil {
		aLog.Fatalf("cant init db %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	repos := repository.NewRepository(db)

	rpcSrv := handler.NewUserRpcsServer(aLog, repos) //TODO: ADD SERVICE
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
