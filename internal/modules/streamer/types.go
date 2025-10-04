package streamer

import (
	"stream-radar/domain/dto"
)

type RegisterStreamersRequest struct {
	Streamers []dto.StreamerDTO `json:"streamers"`
}

type IStreamerService interface {
	RegisterStreamers(streamers []dto.StreamerDTO) error
}
