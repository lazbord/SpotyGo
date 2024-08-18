package service

import (
	"github.com/lazbord/SpotyGo/services/streaming/database"
)

type StreamingService struct {
	db *database.Adapter
}

func NewStreamingService(db *database.Adapter) *StreamingService {
	return &StreamingService{
		db: db,
	}
}
