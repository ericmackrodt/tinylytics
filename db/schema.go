package db

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	UserIdent    string `gorm:"index"`
	Browser      string `gorm:"index"`
	BrowserMajor string
	BrowserMinor string
	BrowserPatch string
	OS           string `gorm:"index"`
	OSMajor      string
	OSMinor      string
	OSPatch      string
	Country      string `gorm:"index"`
	UserAgent    string
	SessionStart time.Time
	SessionEnd   time.Time
	Events       int64
}

type UserEvent struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Page      string
	EventTime time.Time
	SessionID string
	Session   UserSession `gorm:"foreignKey:SessionID;references:ID"`
}