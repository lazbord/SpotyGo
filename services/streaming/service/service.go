package service

import (
	"github.com/lazbord/SpotyGo/model"
	"github.com/lazbord/SpotyGo/services/streaming/database"
	"github.com/pkg/errors"
)

type StreamingService struct {
	db *database.Adapter
}

func NewStreamingService(db *database.Adapter) *StreamingService {
	return &StreamingService{
		db: db,
	}
}

func (a *StreamingService) ServiceGetMusicByID(id string) (*model.Music, error) {
	music, err := a.db.DBGetMusicByID(id)
	if err != nil {
		return nil, errors.New("No music found with this id")
	}

	return music, nil
}

func (a *StreamingService) ServiceAddMusic() {
	music := model.Music{
		ID:       "123456789",
		Name:     "Beat it",
		Artist:   "Michaël Jakson",
		Duration: "1:00",
	}

	a.db.DBAddMusic(music)
}
