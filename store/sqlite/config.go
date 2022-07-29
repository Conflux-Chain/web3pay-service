package sqlite

import (
	stdLog "log"
	"os"
	"time"

	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Config represents the sqlite configurations to open a database instance.
type Config struct {
	Database string `default:"file::memory:?cache=shared"`
	// For shared memory sqlite instance, never expire the connection otherwise
	// the memory db may be released once no connection existed.
	ConnMaxLifetime time.Duration `default:"0"`
	MaxOpenConns    int           `default:"10"`
	MaxIdleConns    int           `default:"10"`
}

// MustNewConfigFromViper creates an instance of Config from Viper.
func MustNewConfigFromViper() Config {
	var config Config
	viperutil.MustUnmarshalKey("store.sqlite", &config)

	return config
}

// MustOpenOrCreate creates an instance of store or exits on any error.
func (config *Config) MustOpenOrCreate(tables ...interface{}) *SqliteStore {
	db := config.MustNewDB(config.Database)

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		db = db.Debug()
	}

	var currentTables []string
	res := db.Select("name").Table("sqlite_schema").
		Where("type='table' AND name NOT LIKE 'sqlite_%'").Find(&currentTables)
	if res.Error != nil {
		logrus.WithError(res.Error).Fatal("Failed to query tables")
	}

	if len(currentTables) == 0 && len(tables) > 0 {
		if err := db.Migrator().CreateTable(tables...); err != nil {
			logrus.WithError(err).Fatal("Failed to create database tables")
		}
	}

	if sqlDb, err := db.DB(); err != nil {
		logrus.WithError(err).Fatal("Failed to init sqlite db")
	} else {
		sqlDb.SetConnMaxLifetime(config.ConnMaxLifetime)
		sqlDb.SetMaxOpenConns(config.MaxOpenConns)
		sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	}

	return NewSqliteStore(db)
}

func (config *Config) MustNewDB(database string) *gorm.DB {
	gLogLevel := gormLogger.Error
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		gLogLevel = gormLogger.Info
	} else if logrus.IsLevelEnabled(logrus.WarnLevel) {
		gLogLevel = gormLogger.Warn
	}

	// create gorm logger by customizing the default logger
	gLogger := gormLogger.New(
		stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
		gormLogger.Config{
			// slow SQL threshold (200ms)
			SlowThreshold: time.Millisecond * 200,
			// log level
			LogLevel: gLogLevel,
			// never logging on ErrRecordNotFound error, otherwise logs may grow exploded
			IgnoreRecordNotFoundError: true,
			// use colorful print
			Colorful: true,
		},
	)

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{Logger: gLogger})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open sqlite")
	}

	return db
}
