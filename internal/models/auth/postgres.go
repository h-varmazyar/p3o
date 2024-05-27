package auth

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/entities"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "users"

type postgresRepository struct {
	*db.DB
	log *log.Logger
}

func NewPostgresRepository(ctx context.Context, logger *log.Logger, db *db.DB) (Model, error) {
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
				Description: fmt.Sprintf("create %s table", tableName),
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

func (r *postgresRepository) Create(ctx context.Context, link *entities.User) error {
	err := r.PostgresDB.Save(link).Error
	if err != nil {
		return ErrFailedToCreateUser.AddOriginalError(err)
	}
	return nil
}

func (r *postgresRepository) ReturnByMobile(ctx context.Context, mobile string) (*entities.User, error) {
	//err := r.PostgresDB.Save(link).Error
	//if err != nil {
	//	return ErrFailedToCreateUser.AddOriginalError(err)
	//}
	return nil, nil
}

func (r *postgresRepository) ReturnByEmail(ctx context.Context, email string) (*entities.User, error) {
	//err := r.PostgresDB.Save(link).Error
	//if err != nil {
	//	return ErrFailedToCreateUser.AddOriginalError(err)
	//}
	return nil, nil
}
