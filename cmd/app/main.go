package main

import (
	"github.com/douglasfsti/golang-shortener-api/config"
	"github.com/douglasfsti/golang-shortener-api/internal/api"
	Serializer "github.com/douglasfsti/golang-shortener-api/internal/serializer"
	"github.com/douglasfsti/golang-shortener-api/pkg/shortner"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/sony/sonyflake"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

func main() {
	fx.New(
		fx.Provide(
			chi.NewRouter,
			Serializer.NewSerializer,
			NewRedis,
			NewIDGenerator,

			shortner.NewRepository,
			shortner.NewUseCases,
			shortner.NewService,

			config.NewContainer,
		),
		fx.Invoke(func(router *chi.Mux, container config.Container) {
			router = api.GetRouter(router, container)
			log.Fatal(http.ListenAndServe(":8080", router))
		}))
}

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func NewIDGenerator() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now().UTC(),
	})
}
