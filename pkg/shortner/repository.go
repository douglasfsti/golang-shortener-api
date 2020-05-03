package shortner

import (
	"fmt"
	Serializer "github.com/douglasfsti/golang-shortener-api/internal/serializer"
	"github.com/go-redis/redis"
)

const (
	NoExpiration      = 0
	ErrFailedToSearch = "key: %v err: %s - repository.Redis.Find"
	ErrDataNotFound   = "key: %v not found - repository.Redis.Find"
	ErrFailedToStore  = "key: %v err: %s failed to store - repository.Redis.Store"
)

type Repository interface {
	Find(code interface{}) (*Redirect, error)
	Store(redirect *Redirect) error
}

func NewRepository(client *redis.Client, serializer Serializer.Serializer) Repository {
	return &repository{
		RedisClient: client,
		Serializer:  serializer,
	}
}

type repository struct {
	RedisClient *redis.Client
	Serializer  Serializer.Serializer
}

func (r *repository) Find(code interface{}) (*Redirect, error) {
	data, err := r.RedisClient.Get(r.generateKey(code)).Bytes()
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToSearch, code, err.Error())
	}

	if len(data) == 0 {
		return nil, fmt.Errorf(ErrDataNotFound, code)
	}

	var redirect *Redirect

	return redirect, r.Serializer.Decode(data, &redirect)
}

func (r *repository) Store(redirect *Redirect) error {
	data, err := r.Serializer.Encode(redirect)
	if err != nil {
		return err
	}

	err = r.RedisClient.Set(r.generateKey(redirect.Code), data, NoExpiration).Err()
	if err != nil {
		return fmt.Errorf(ErrFailedToStore, redirect.Code, err.Error())
	}

	return nil
}

func (r *repository) generateKey(code interface{}) string {
	return fmt.Sprintf("redirect/%v", code)
}
