package models

import (
	"context"
	"phalanx/vender"
	"time"
)

type Tag struct {
	vender.BaseModel `bun:"table:tags"`
	ID               int64  `bun:",pk,autoincrement"`
	Name             string `bun:",unique,notnull"`
}

type Video struct {
	vender.BaseModel `bun:"table:videos"`
	ID               int64  `bun:",pk,autoincrement"`
	UniqueID         string `bun:",unique,notnull"`
	DisplayTitle     string
	Modtime          time.Time
	FileSize         int64
	Path             string
	Deleted          bool
	Duration         time.Duration
}

type FileMeta interface {
	Id() string
	DisplayTitle() string
	ModTime() time.Time
	Size() int64
	Duration() time.Duration
	SourcePath() string
}

func (video *Video) Set(meta FileMeta) *Video {
	video.UniqueID = meta.Id()
	video.DisplayTitle = meta.DisplayTitle()
	video.Modtime = meta.ModTime()
	video.FileSize = meta.Size()
	video.Duration = meta.Duration()
	video.Path = meta.SourcePath()

	return video
}

type VideoTag struct {
	vender.BaseModel `bun:"table:video_tags"`
	ID               int64 `bun:",pk,autoincrement"`
	VideoID          int64 `bun:",unique,notnull"`
	TagID            int64 `bun:",unique,notnull"`
}

type VideoThumbnail struct {
	vender.BaseModel `bun:"table:video_thumbnail"`
	ID               int64  `bun:",pk,autoincrement"`
	VideoID          int64  `bun:",unique,notnull"`
	ObjectID         string `bun:",unique"`
	Created          time.Time
}

type CrawlLog struct {
	vender.BaseModel `bun:"table:crawl_logs"`
	ID               int64 `bun:",pk,autoincrement"`
	Created          time.Time
}

func Migration(db vender.Database, ctx context.Context) error {
	models := []any{
		(*Tag)(nil),
		(*Video)(nil),
		(*VideoTag)(nil),
		(*VideoThumbnail)(nil),
		(*CrawlLog)(nil),
	}
	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
