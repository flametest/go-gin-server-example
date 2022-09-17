package db

import "time"

type Base struct {
	DomainEvent string    `gorm:"type:varchar(32);column:domain_event"`
	Id          uint64    `gorm:"type:bigint;primaryKey;autoIncrement;column:id"`
	Version     uint64    `gorm:"type:bigint;column:version"`
	CreatedAt   time.Time `gorm:"->;<-:create;type:DATETIME(3);default:CURRENT_TIMESTAMP(3) not null;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(3);default:CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3);column:updated_at"`
}
