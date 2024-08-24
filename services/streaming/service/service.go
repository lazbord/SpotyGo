package service

import (
	"github.com/lazbord/SpotyGo/common/model"
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

func (a *StreamingService) ServiceGetMusicByID(videoID string) (*model.Music, error) {
	music, err := a.db.DBGetMusicByID(videoID)
	if err != nil {
		return nil, err
	}
	return music, nil
}
