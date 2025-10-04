package streamer

import (
	"fmt"
	"go.uber.org/zap"
	"stream-radar/domain/dto"
	"stream-radar/domain/model"
	"stream-radar/internal/database"
	"stream-radar/internal/logger"
)

type StreamerService struct{}

var service *StreamerService

func init() {
	service = &StreamerService{}
}

func (StreamerService) RegisterStreamers(streamers []dto.StreamerDTO, user model.User) ([]model.Streamer, error) {
	log := logger.GetInstance()

	var streamersData []model.Streamer
	for _, streamer := range streamers {
		streamersData = append(streamersData, model.Streamer{}.FromDto(streamer))
	}

	res := database.DB.CreateInBatches(&streamersData, len(streamersData))

	if res.Error != nil {
		log.Error(fmt.Sprintf("Error creating streamers"), zap.Error(res.Error))
		return nil, res.Error
	}

	err := database.DB.Model(&user).Association("Streamers").Append(&streamersData)

	if err != nil {
		log.Error(fmt.Sprintf("Error associating streamers"), zap.Error(res.Error))
		return nil, err
	}

	return streamersData, nil
}
