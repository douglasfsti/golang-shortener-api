package config

import (
	Serializer "github.com/douglasfsti/golang-shortener-api/internal/serializer"
	"github.com/douglasfsti/golang-shortener-api/pkg/shortner"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/sony/sonyflake"
	"time"
)

type Container interface {
	// shortner
	GetShortnerService() shortner.Service
	GetShortnerUseCases() shortner.UseCases
	GetShortnerRepository() shortner.Repository
}

func NewContainer() Container {
	return &container{}
}

type container struct {
	// generics
	Serializer Serializer.Serializer
	Router     chi.Router

	// shortner interfaces
	ShortnerService    shortner.Service
	ShortnerUseCases   shortner.UseCases
	ShortnerRepository shortner.Repository

	// shortner third party dependencies
	RedisClient   *redis.Client
	CodeGenerator *sonyflake.Sonyflake
}

// generics
func (c *container) GetSerializer() Serializer.Serializer {
	if c.Serializer == nil {
		c.Serializer = c.NewSerializer()
	}

	return c.Serializer
}

func (c *container) NewSerializer() Serializer.Serializer {
	return Serializer.NewSerializer()
}

// shortner interfaces
func (c *container) GetShortnerService() shortner.Service {
	if c.ShortnerService == nil {
		c.ShortnerService = c.NewShortnerService()
	}

	return c.ShortnerService
}

func (c *container) NewShortnerService() shortner.Service {
	return shortner.NewService(c.GetShortnerRepository(), c.GetShortnerUseCases())
}

func (c *container) GetShortnerUseCases() shortner.UseCases {
	if c.ShortnerUseCases == nil {
		c.ShortnerUseCases = c.NewShortnerUseCases()
	}

	return c.ShortnerUseCases
}

func (c *container) NewShortnerUseCases() shortner.UseCases {
	return shortner.NewUseCases(c.GetCodeGenerator(), c.GetSerializer())
}

func (c *container) GetShortnerRepository() shortner.Repository {
	if c.ShortnerRepository == nil {
		c.ShortnerRepository = c.NewShortnerRepository()
	}

	return c.ShortnerRepository
}

func (c *container) NewShortnerRepository() shortner.Repository {
	return shortner.NewRepository(c.GetRedisClient(), c.GetSerializer())
}

// shortner third party dependencies
func (c *container) GetCodeGenerator() *sonyflake.Sonyflake {
	if c.CodeGenerator == nil {
		c.CodeGenerator = c.NewCodeGenerator()
	}

	return c.CodeGenerator
}

func (c *container) NewCodeGenerator() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now().UTC(),
	})
}

func (c *container) GetRedisClient() *redis.Client {
	if c.RedisClient == nil {
		c.RedisClient = c.NewRedisClient()
	}

	if c.RedisClient.Ping().Err() != nil {
		c.RedisClient = c.NewRedisClient()
	}

	return c.RedisClient
}

func (c *container) NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       RedisDatabase,
	})
}
