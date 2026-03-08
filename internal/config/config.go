package config

import (
    "os"
    "time"
)

// Config centralizes application configuration loaded from environment variables.
// Keep this small and explicit so defaults are predictable in production.
type Config struct {
    AppEnv   string
    AppPort  string
    DBDSN    string

    AccessTokenExpires  time.Duration
    RefreshTokenExpires time.Duration

    JWTAccessSecret  string
    JWTRefreshSecret string

    AwsRegion   string
    AwsS3Bucket string

    AdminAPIKey string

    CorsAllowOrigins string
    OtelEndpoint     string
}

// Load reads configuration from environment variables and applies safe defaults.
// It intentionally does not panic; callers may validate required values.
func Load() *Config {
    cfg := &Config{
        AppEnv:          getenv("APP_ENV", "development"),
        AppPort:         getenv("APP_PORT", "8080"),
        DBDSN:           os.Getenv("DB_DSN"),
        JWTAccessSecret: os.Getenv("JWT_ACCESS_SECRET"),
        JWTRefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),

        AwsRegion:   getenv("AWS_REGION", ""),
        AwsS3Bucket: getenv("AWS_S3_BUCKET", ""),

        AdminAPIKey:     os.Getenv("ADMIN_API_KEY"),
        CorsAllowOrigins: getenv("CORS_ALLOW_ORIGINS", "*"),
        OtelEndpoint:     getenv("OTEL_EXPORTER_OTLP_ENDPOINT", ""),
    }

    cfg.AccessTokenExpires = parseDuration(getenv("ACCESS_TOKEN_EXPIRES", "15m"))
    cfg.RefreshTokenExpires = parseDuration(getenv("REFRESH_TOKEN_EXPIRES", "720h"))

    return cfg
}

func getenv(key, def string) string {
    v := os.Getenv(key)
    if v == "" {
        return def
    }
    return v
}

func parseDuration(v string) time.Duration {
    d, err := time.ParseDuration(v)
    if err != nil {
        // Fall back to safe default
        return 15 * time.Minute
    }
    return d
}
