package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/bootstrapper"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/config"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func main() {
	//new validator
	v := validator.New()

	//parent context
	ctx := context.Background()

	//custom logger
	log := logger.NewSlogAdapter(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	//load config
	cfg := config.LoadConfig()

	//validate .env config
	cfg.Validate(v)

	DB_PORT, err := strconv.Atoi(cfg.DBPort)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	dbConfig := &database.DBConfig{
		Host:     cfg.DBHost,
		Port:     DB_PORT,
		Username: cfg.DBUsername,
		Password: cfg.DBPassword,
		Database: cfg.DBDatabase,
		Timezone: cfg.DBTimezone,
		Retry: config.RetryConfig{
			Max:   cfg.RetryConfig.Max,
			Delay: cfg.RetryConfig.Delay,
		},
	}

	db, err := dbConfig.InitConnectionDB(ctx, cfg)

	if err != nil {
		log.Fatal(ctx, fmt.Sprintf("[DB] Database Connection Failed ,%v", err))

	}

	echo := echo.New()

	apiServer := bootstrapper.NewServer(echo, cfg, db)

	apiServer.InitAPI()

	listen := net.JoinHostPort(cfg.AppHost, cfg.AppPort)

	srv := &http.Server{
		Addr:    listen,
		Handler: echo,
	}

	//blocking operation
	log.Info(ctx, fmt.Sprintf("[Server] Running on : http://%v:%v", cfg.AppHost, cfg.AppPort))

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(ctx, fmt.Sprintf("[Server] Error running server : %v", err))
	}

}
