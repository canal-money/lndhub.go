package models

import (
	"time"
	"github.com/uptrace/bun"
)

// Currency : Currency Model

type Currency struct {
	ID          int64          `bun:",pk,autoincrement"`
	Name        string         `bun:",unique,notnull"`
	Code		string         `bun:",unique,notnull"`
	Granularity int64          `bun:",notnull"`
	CreatedAt   time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   bun.NullTime
}