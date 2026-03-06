package db

import (
    "log"
    "os"
    "time"

    "user-profile-system-backend-go/internal/models"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectMySQL() {
    dsn := os.Getenv("DB_DSN")

    conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatalf("MySQL connection error: %v", err)
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
    )
}
