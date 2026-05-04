package db

import (
    "log"
    "os"
    "time"

    "user-profile-system-backend-go/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectMySQL() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN environment variable is required")
    }

    conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatalf("PostgreSQL connection error: %v", err)
    }

    sqlDB, _ := conn.DB()
    sqlDB.SetMaxOpenConns(25)
    sqlDB.SetMaxIdleConns(25)
    sqlDB.SetConnMaxLifetime(5 * time.Minute)

    DB = conn

    migrate()
}

func migrate() {
    DB.AutoMigrate(
        &models.User{},
        &models.RefreshToken{},
        &models.UserProfile{},
        &models.UserSettings{},
        &models.UserFavorite{},
        &models.UserDownload{},
        &models.UserHistory{},
        &models.ActivityLog{},
        &models.ListeningEvent{},
    )
}
