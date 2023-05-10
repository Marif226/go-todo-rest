package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Marif226/go-todo-rest/internal/handler"
	"github.com/Marif226/go-todo-rest/internal/repository"
	"github.com/Marif226/go-todo-rest/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// set format for logger
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//configure logger
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)

	err := initConfig()
	if err != nil {
		logrus.Fatalf("error during reading the config file, %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error during loading .env variable, %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error during connecting to the database, %s", err.Error())
	}

	r := repository.New(db)
	s := service.New(r)
	h := handler.New(s)

	// srv := server.New()
	e := echo.New()

	// logging middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			res := ctx.Response()

			logrus.Infof("%s %s", req.Method, req.RequestURI)
			err := next(ctx)
			if err != nil {
				ctx.Error(err)
			}
			logrus.Infof("[%d] %s", res.Status, http.StatusText(res.Status))
			return nil
		}
	})

	h.InitRoutes(e)

	// Start server
	go func() {
		err = e.Start(viper.GetString("port"))
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error during running http server, %s", err.Error())
		}
	}()

	// Graceful shutdown

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds. 
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")

	return viper.ReadInConfig()
}