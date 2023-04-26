package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lovelyrrg51/go_backend/app/config"
	"github.com/lovelyrrg51/go_backend/app/logger"
	"github.com/lovelyrrg51/go_backend/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	lgmql "gorm.io/gorm/logger"
)

var DB *gorm.DB

func getDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", config.Cfg.DbUser, config.Cfg.DbPass, config.Cfg.DbHost, config.Cfg.DbPort, config.Cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: lgmql.Default.LogMode(lgmql.Info)})

	if err != nil {
		logger.Error("Error when connect to database " + err.Error())
		os.Exit(1)
	}

	sqlDB, errPool := db.DB()
	if errPool != nil {
		logger.Error("Error when setting option for database " + errPool.Error())
		os.Exit(1)
	}

	maxIdle, _ := strconv.Atoi(config.Cfg.DbMaxIdle)
	maxOpenCon, _ := strconv.Atoi(config.Cfg.DbMaxOpenCon)
	maxLifeTime, _ := time.ParseDuration(config.Cfg.DbMaxLifetime)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpenCon)
	sqlDB.SetConnMaxLifetime(maxLifeTime)

	logger.Info("Connect to database successfully")

	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		logger.Error("Error when auto migrate database " + err.Error())
		os.Exit(1)
	}
	logger.Info("Migrated database")

	return db
}

func init() {
	DB = getDatabase()
}
