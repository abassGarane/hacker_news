package repository

import (
	"context"
	"errors"
	"hacker_news/models"
	"log/slog"

	"github.com/gofiber/fiber/v2/log"
	"github.com/uptrace/bun"
)

type Repository struct {
	DB *bun.DB
}

func New(dsn string) (*Repository, error) {
	db, err := CreateDb(dsn)
	if err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

func (r *Repository) GetAllNews() (items []*models.NewsItem, err error) {
	slog.Info("Get all news was called here")
	err = r.DB.NewSelect().Model(&items).Scan(context.Background())
	log.Debug(items)
	log.Debug(err)
	return
}

func (r *Repository) GetOneById(id int64) (item *models.NewsItem, err error) {
	log.Info("GetItemById reached")
	exists, err := r.DB.NewSelect().Model(&item).Where("id = ?", id).Exists(context.Background())
	if !exists {
		log.Debug(exists)
		return nil, errors.Join(err, errors.New("user does not exist"))
	}
	log.Debug(err)
	log.Debug(item)
	return
}
