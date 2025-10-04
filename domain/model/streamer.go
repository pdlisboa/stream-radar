package model

import (
	"encoding/json"
	"gorm.io/datatypes"
	"stream-radar/domain/dto"
)

type Streamer struct {
	Id      uint           `gorm:"primaryKey" json:"id"`
	Name    string         `gorm:"not null" json:"name"`
	Aliases datatypes.JSON `gorm:"not null" json:"aliases"`
}

func (streamer Streamer) ToDto() dto.StreamerDTO {
	var aliases []string
	err := streamer.Aliases.Scan(&aliases)

	if err != nil {
		return dto.StreamerDTO{
			Name:    streamer.Name,
			Aliases: aliases,
		}
	}

	return dto.StreamerDTO{}
}

func (Streamer) FromDto(dto dto.StreamerDTO) Streamer {
	aliasesByte, err := json.Marshal(dto.Aliases)
	if err != nil {
		panic(err)
	}

	return Streamer{
		Name:    dto.Name,
		Aliases: aliasesByte,
	}
}
