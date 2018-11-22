package domain

import "time"

type Category int

const (
	CategoryUnknown Category = iota
	CategoryAnimation
	CategoryMusic
	CategoryVariety
	CategoryDrama
	CategoryCinema
	CategorySports
	CategoryInformation
	CategoryNews
	CategoryDocumentary
	CategoryTheater
	CategoryHobby
	CategoryWelfare
	CategoryEtc
)

type PublishStatus int

const (
	PublishStatusUnknown PublishStatus = iota
	PublishStatusUnpublished
	PublishStatusPublished
	PublishStatusArchived
	PublishStatusTemporaryRestored
)

type Video struct {
	ID            string
	SeriesID      string
	Title         string
	Description   string
	Chapter       int
	Category      Category
	PublishStatus PublishStatus
	Duration      time.Duration
	Sha1Sum       string
	FileSize      uint64
	Width         uint
	Height        uint
	RecordedAt    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
