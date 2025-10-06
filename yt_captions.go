package yt_captions

import (
	"github.com/elbombardi/yt_captions/internal"
	"github.com/elbombardi/yt_captions/models"
)

type CaptionsLoader interface {
	GetAllCaptions(video models.VideoID) ([]*models.Caption, error)
	GetLanguages(video models.VideoID) ([]string, error)
	GetCaptions(video models.VideoID, lang string) (*models.Caption, error)
}

func NewDefaultCaptionsLoader() CaptionsLoader {
	return internal.NewCaptionsParser()
}
