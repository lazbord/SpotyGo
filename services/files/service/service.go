package service

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/kkdai/youtube/v2"
	"github.com/lazbord/SpotyGo/model"
	"github.com/lazbord/SpotyGo/services/files/database"
)

type FilesService struct {
	db *database.Adapter
}

func NewFilesService(db *database.Adapter) *FilesService {
	return &FilesService{
		db: db,
	}
}

func (a *FilesService) DownloadVideo(videoID string) error {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return err
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return err
	}
	defer stream.Close()

	videoName := videoID + ".mp4"
	mp3FileName := videoID + ".mp3"

	file, err := os.Create(videoName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}

	err = ConvertToMP3(videoName, mp3FileName)
	if err != nil {
		log.Fatalf("Failed to convert MP4 to MP3: %s\n", err)
		return err
	}

	mp3Data, err := os.ReadFile(mp3FileName)
	if err != nil {
		return err
	}

	music := model.Music{
		ID:       videoID,
		Name:     video.Title,
		Artist:   video.Author,
		Duration: formatDuration(video.Duration),
		Data:     mp3Data,
		Filename: mp3FileName,
	}

	_, err = a.db.DBAddMusic(music)
	if err != nil {
		return err
	}

	os.Remove(videoName)
	os.Remove(mp3FileName)

	return nil
}

func ConvertToMP3(inputFile, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, outputFile, "-y")
	return cmd.Run()
}

func formatDuration(d time.Duration) string {
	return d.String()
}
