package utils

import "time"

// NowUTC returns current time in UTC.
func NowUTC() time.Time {
    return time.Now().UTC()
}

// ParseRFC3339 parses RFC3339 timestamps safely.
func ParseRFC3339(v string) (time.Time, error) {
    return time.Parse(time.RFC3339, v)
}

// StartOfDayUTC returns midnight UTC for a given time.
func StartOfDayUTC(t time.Time) time.Time {
    u := t.UTC()
    return time.Date(u.Year(), u.Month(), u.Day(), 0, 0, 0, 0, time.UTC)
}
