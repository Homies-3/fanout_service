package service

import (
	"context"
	"encoding/json"
	"fanout_service/models"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type fanoutService struct {
	l     *log.Logger
	cache *redis.Client
}

func NewFanoutService(l *log.Logger, cache *redis.Client) IFanoutService {
	return fanoutService{
		l:     l,
		cache: cache,
	}
}

func (fS fanoutService) Publish(p models.Post) error {
	b, err := json.Marshal(&p)
	if err != nil {
		return err
	}

	rCmd := fS.cache.Set(context.Background(), p.UserId.String(), b, time.Hour*1)
	if err = rCmd.Err(); err != nil {
		fS.l.Println(err)
		return err
	}

	return nil
}
