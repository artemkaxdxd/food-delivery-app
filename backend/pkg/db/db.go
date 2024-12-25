package db

import (
	"backend/config"
	"backend/pkg/logger"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type (
	Database interface {
		Instance() *gorm.DB
		Close() error
		Ping()
	}

	Postgres struct {
		DB *gorm.DB
		l  logger.Logger
	}
)

func NewPostgres(cfg config.Postgres, l logger.Logger) (*Postgres, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DB, cfg.Port)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return &Postgres{DB: db, l: l}, nil
}

func (p *Postgres) Instance() *gorm.DB {
	return p.DB
}

func (p *Postgres) Close() error {
	if p.DB == nil {
		return errors.New("db connection is already closed")
	}
	db, err := p.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (p *Postgres) Ping() {
	sqlDB, err := p.DB.DB()
	if err != nil {
		text := fmt.Sprintf("postgres ping error: %s", err.Error())
		p.l.Error(text, err)
	}

	err = sqlDB.Ping()
	if err != nil {
		text := fmt.Sprintf("postgres ping error: %s", err.Error())
		p.l.Error(text, err)
	}
}
