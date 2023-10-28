package repository

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "link"

type postgresRepository struct {
	*db.DB
	log *log.Logger
}

func NewPostgresRepository(ctx context.Context, logger *log.Logger, db *db.DB) (Repository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return &postgresRepository{
		DB:  db,
		log: logger,
	}, nil
}

func migration(_ context.Context, dbInstance *db.DB) error {
	var err error
	migrations := make(map[string]struct{})
	tags := make([]string, 0)
	err = dbInstance.PostgresDB.Table(db.MigrationTable).Where("table_name = ?", tableName).Select("tag").Find(&tags).Error
	if err != nil {
		return err
	}

	for _, tag := range tags {
		migrations[tag] = struct{}{}
	}

	newMigrations := make([]*db.Migration, 0)
	err = dbInstance.PostgresDB.Transaction(func(tx *gorm.DB) error {
		if _, ok := migrations["v1.0.0"]; !ok {
			err = tx.AutoMigrate(new(entities.Link))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				TableName:   tableName,
				Tag:         "v1.0.0",
				Description: "create link table",
			})
		}
		err = tx.Model(new(db.Migration)).CreateInBatches(&newMigrations, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) Create(ctx context.Context, link *entities.Link) error {
	err := r.PostgresDB.Save(link).Error
	if err != nil {
		return ErrFailedToCreateLink.AddOriginalError(err)
	}
	return nil
}

func (r *postgresRepository) ReturnByKey(ctx context.Context, key string) (*entities.Link, error) {
	var link *entities.Link
	err := r.PostgresDB.Model(new(entities.Link)).Where("key = ?", key).First(link).Error
	if err != nil {
		return nil, ErrLinkNotFound.AddOriginalError(err)
	}
	return link, nil
}

func (r *postgresRepository) Visit(_ context.Context, id uint) error {
	err := r.PostgresDB.
		Model(new(entities.Link)).
		Where("id = ?", id).
		Update("total_visit", gorm.Expr("total_visit + 1")).Error

	if err != nil {
		return ErrIncreaseVisitFailed.AddOriginalError(err)
	}
	return nil
}
