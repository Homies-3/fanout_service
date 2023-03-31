package service

import "fanout_service/models"

type IFanoutService interface {
	Publish(p models.Post) error
}
