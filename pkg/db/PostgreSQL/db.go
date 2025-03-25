package db

//
//import (
//	"context"
//	gormext "github.com/h-varmazyar/gopack/gorm"
//	log "github.com/sirupsen/logrus"
//	"go.uber.org/fx"
//	"gorm.io/gorm"
//)
//
//type DB struct {
//	PostgresDB *gorm.DB
//}
//
//type Params struct {
//	fx.In
//
//	Context context.Context
//	Log     *log.Logger
//}
//
//type Result struct {
//}
//
//func NewDatabase(p Params) (*DB, error) {
//	db := new(DB)
//	if configs.DbType == gormext.PostgreSQL {
//		postgres, err := newPostgres(p.Context, *configs)
//		if err != nil {
//			log.WithError(err).Error("failed to create new postgres")
//			return nil, err
//		}
//		db.PostgresDB = postgres
//
//		err = createMigrateTable(p.Context, db)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		return nil, ErrInvalidDB
//	}
//
//	return db, nil
//}
//
//func newPostgres(_ context.Context, configs gormext.Configs) (*gorm.DB, error) {
//	db, err := gormext.Open(configs)
//	if err != nil {
//		log.WithError(err).Error("failed to open database")
//		return nil, err
//	}
//
//	if err = db.Transaction(func(tx *gorm.DB) error {
//		if err = gormext.EnableExtensions(tx,
//			gormext.UUIDExtension,
//		); err != nil {
//			return err
//		}
//		return nil
//	}); err != nil {
//		log.WithError(err).Error("failed to add extensions to database")
//		return nil, err
//	}
//
//	return db, nil
//}
//
//func createMigrateTable(_ context.Context, db *DB) error {
//	err := db.PostgresDB.AutoMigrate(new(Migration))
//	if err != nil {
//		return err
//	}
//	return nil
//}
