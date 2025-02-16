package app

import (
	"avito-shop-service/internal/bootstrap"
	"avito-shop-service/internal/config"
	"avito-shop-service/internal/database/repository"
	"avito-shop-service/internal/database/txmanager"
	"avito-shop-service/internal/handler"
	"avito-shop-service/internal/http"
	"avito-shop-service/internal/service/crypto"
	"avito-shop-service/internal/service/jwtservice"
	"avito-shop-service/internal/usecase"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {
	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	if err = bootstrap.ApplyMigrations(cfg, db); err != nil {
		return err
	}

	txm := txmanager.New(db)

	jwt := jwtservice.New(cfg)

	crypt := crypto.New()

	repos := repository.New(db)

	useCases := usecase.New(jwt, crypt, txm, repos)

	handlers := handler.New(useCases)

	server := http.NewServer(cfg, handlers)

	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Close(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")

	return nil
}
